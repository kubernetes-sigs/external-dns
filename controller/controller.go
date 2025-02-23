/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/metrics"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

var (
	registryErrorsTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	sourceErrorsTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)
	sourceEndpointsTotal = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in all sources",
		},
	)
	registryEndpointsTotal = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in the registry",
		},
	)
	lastSyncTimestamp = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "last_sync_timestamp_seconds",
			Help:      "Timestamp of last successful sync with the DNS provider",
		},
	)
	lastReconcileTimestamp = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "last_reconcile_timestamp_seconds",
			Help:      "Timestamp of last attempted sync with the DNS provider",
		},
	)
	controllerNoChangesTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "no_op_runs_total",
			Help:      "Number of reconcile loops ending up with no changes on the DNS provider side.",
		},
	)
	deprecatedRegistryErrors = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	deprecatedSourceErrors = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)
	registryARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "a_records",
			Help:      "Number of Registry A records.",
		},
	)
	registryAAAARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "aaaa_records",
			Help:      "Number of Registry AAAA records.",
		},
	)
	sourceARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "a_records",
			Help:      "Number of Source A records.",
		},
	)
	sourceAAAARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "aaaa_records",
			Help:      "Number of Source AAAA records.",
		},
	)
	verifiedARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_a_records",
			Help:      "Number of DNS A-records that exists both in source and registry.",
		},
	)
	verifiedAAAARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_aaaa_records",
			Help:      "Number of DNS AAAA-records that exists both in source and registry.",
		},
	)
)

func init() {
	metrics.RegisterMetric.MustRegister(registryErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceEndpointsTotal)
	metrics.RegisterMetric.MustRegister(registryEndpointsTotal)
	metrics.RegisterMetric.MustRegister(lastSyncTimestamp)
	metrics.RegisterMetric.MustRegister(lastReconcileTimestamp)
	metrics.RegisterMetric.MustRegister(deprecatedRegistryErrors)
	metrics.RegisterMetric.MustRegister(deprecatedSourceErrors)
	metrics.RegisterMetric.MustRegister(controllerNoChangesTotal)
	metrics.RegisterMetric.MustRegister(registryARecords)
	metrics.RegisterMetric.MustRegister(registryAAAARecords)
	metrics.RegisterMetric.MustRegister(sourceARecords)
	metrics.RegisterMetric.MustRegister(sourceAAAARecords)
	metrics.RegisterMetric.MustRegister(verifiedARecords)
	metrics.RegisterMetric.MustRegister(verifiedAAAARecords)
}

// Controller is responsible for orchestrating the different components.
// It works in the following way:
// * Ask the DNS provider for current list of endpoints.
// * Ask the Source for the desired list of endpoints.
// * Take both lists and calculate a Plan to move current towards desired state.
// * Tell the DNS provider to apply the changes calculated by the Plan.
type Controller struct {
	Source   source.Source
	Registry registry.Registry
	// The policy that defines which changes to DNS records are allowed
	Policy plan.Policy
	// The interval between individual synchronizations
	Interval time.Duration
	// The DomainFilter defines which DNS records to keep or exclude
	DomainFilter endpoint.DomainFilterInterface
	// The nextRunAt used for throttling and batching reconciliation
	nextRunAt time.Time
	// The runAtMutex is for atomic updating of nextRunAt and lastRunAt
	runAtMutex sync.Mutex
	// The lastRunAt used for throttling and batching reconciliation
	lastRunAt time.Time
	// MangedRecordTypes are DNS record types that will be considered for management.
	ManagedRecordTypes []string
	// ExcludeRecordTypes are DNS record types that will be excluded from management.
	ExcludeRecordTypes []string
	// MinEventSyncInterval is used as window for batching events
	MinEventSyncInterval time.Duration
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {
	lastReconcileTimestamp.Gauge.SetToCurrentTime()

	c.runAtMutex.Lock()
	c.lastRunAt = time.Now()
	c.runAtMutex.Unlock()

	records, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Counter.Inc()
		deprecatedRegistryErrors.Counter.Inc()
		return err
	}

	registryEndpointsTotal.Gauge.Set(float64(len(records)))
	regARecords, regAAAARecords := countAddressRecords(records)
	registryARecords.Gauge.Set(float64(regARecords))
	registryAAAARecords.Gauge.Set(float64(regAAAARecords))
	ctx = context.WithValue(ctx, provider.RecordsContextKey, records)

	endpoints, err := c.Source.Endpoints(ctx)
	if err != nil {
		sourceErrorsTotal.Counter.Inc()
		deprecatedSourceErrors.Counter.Inc()
		return err
	}
	sourceEndpointsTotal.Gauge.Set(float64(len(endpoints)))
	srcARecords, srcAAAARecords := countAddressRecords(endpoints)
	sourceARecords.Gauge.Set(float64(srcARecords))
	sourceAAAARecords.Gauge.Set(float64(srcAAAARecords))
	vARecords, vAAAARecords := countMatchingAddressRecords(endpoints, records)
	verifiedARecords.Gauge.Set(float64(vARecords))
	verifiedAAAARecords.Gauge.Set(float64(vAAAARecords))
	endpoints, err = c.Registry.AdjustEndpoints(endpoints)
	if err != nil {
		return fmt.Errorf("adjusting endpoints: %w", err)
	}
	registryFilter := c.Registry.GetDomainFilter()

	plan := &plan.Plan{
		Policies:       []plan.Policy{c.Policy},
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{c.DomainFilter, registryFilter},
		ManagedRecords: c.ManagedRecordTypes,
		ExcludeRecords: c.ExcludeRecordTypes,
		OwnerID:        c.Registry.OwnerID(),
	}

	plan = plan.Calculate()

	if plan.Changes.HasChanges() {
		err = c.Registry.ApplyChanges(ctx, plan.Changes)
		if err != nil {
			registryErrorsTotal.Counter.Inc()
			deprecatedRegistryErrors.Counter.Inc()
			return err
		}
	} else {
		controllerNoChangesTotal.Counter.Inc()
		log.Info("All records are already up to date")
	}

	lastSyncTimestamp.Gauge.SetToCurrentTime()

	return nil
}

func earliest(r time.Time, times ...time.Time) time.Time {
	for _, t := range times {
		if t.Before(r) {
			r = t
		}
	}
	return r
}

func latest(r time.Time, times ...time.Time) time.Time {
	for _, t := range times {
		if t.After(r) {
			r = t
		}
	}
	return r
}

// Counts the intersections of A and AAAA records in endpoint and registry.
func countMatchingAddressRecords(endpoints []*endpoint.Endpoint, registryRecords []*endpoint.Endpoint) (int, int) {
	recordsMap := make(map[string]map[string]struct{})
	for _, regRecord := range registryRecords {
		if _, found := recordsMap[regRecord.DNSName]; !found {
			recordsMap[regRecord.DNSName] = make(map[string]struct{})
		}
		recordsMap[regRecord.DNSName][regRecord.RecordType] = struct{}{}
	}
	aCount := 0
	aaaaCount := 0
	for _, sourceRecord := range endpoints {
		if _, found := recordsMap[sourceRecord.DNSName]; found {
			if _, found := recordsMap[sourceRecord.DNSName][sourceRecord.RecordType]; found {
				switch sourceRecord.RecordType {
				case endpoint.RecordTypeA:
					aCount++
				case endpoint.RecordTypeAAAA:
					aaaaCount++
				}
			}
		}
	}
	return aCount, aaaaCount
}

func countAddressRecords(endpoints []*endpoint.Endpoint) (int, int) {
	aCount := 0
	aaaaCount := 0
	for _, endPoint := range endpoints {
		switch endPoint.RecordType {
		case endpoint.RecordTypeA:
			aCount++
		case endpoint.RecordTypeAAAA:
			aaaaCount++
		}
	}
	return aCount, aaaaCount
}

// ScheduleRunOnce makes sure execution happens at most once per interval.
func (c *Controller) ScheduleRunOnce(now time.Time) {
	c.runAtMutex.Lock()
	defer c.runAtMutex.Unlock()
	c.nextRunAt = latest(
		c.lastRunAt.Add(c.MinEventSyncInterval),
		earliest(
			now.Add(5*time.Second),
			c.nextRunAt,
		),
	)
}

func (c *Controller) ShouldRunOnce(now time.Time) bool {
	c.runAtMutex.Lock()
	defer c.runAtMutex.Unlock()
	if now.Before(c.nextRunAt) {
		return false
	}
	c.nextRunAt = now.Add(c.Interval)
	return true
}

// Run runs RunOnce in a loop with a delay until context is canceled
func (c *Controller) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if c.ShouldRunOnce(time.Now()) {
			if err := c.RunOnce(ctx); err != nil {
				if errors.Is(err, provider.SoftError) {
					log.Errorf("Failed to do run once: %v", err)
				} else {
					log.Fatalf("Failed to do run once: %v", err)
				}
			}
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			log.Info("Terminating main controller loop")
			return
		}
	}
}
