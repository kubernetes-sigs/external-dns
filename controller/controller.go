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
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

var (
	registryErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	sourceErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)
	sourceEndpointsTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in all sources",
		},
	)
	registryEndpointsTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in the registry",
		},
	)
	lastSyncTimestamp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "last_sync_timestamp_seconds",
			Help:      "Timestamp of last successful sync with the DNS provider",
		},
	)
	controllerNoChangesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "no_op_runs_total",
			Help:      "Number of reconcile loops ending up with no changes on the DNS provider side.",
		},
	)
	deprecatedRegistryErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	deprecatedSourceErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)
	registryARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "a_records",
			Help:      "Number of Registry A records.",
		},
	)
	registryAAAARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "aaaa_records",
			Help:      "Number of Registry AAAA records.",
		},
	)
	sourceARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "a_records",
			Help:      "Number of Source A records.",
		},
	)
	sourceAAAARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "aaaa_records",
			Help:      "Number of Source AAAA records.",
		},
	)
	verifiedARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_a_records",
			Help:      "Number of DNS A-records that exists both in source and registry.",
		},
	)
	verifiedAAAARecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_aaaa_records",
			Help:      "Number of DNS AAAA-records that exists both in source and registry.",
		},
	)
)

func init() {
	prometheus.MustRegister(registryErrorsTotal)
	prometheus.MustRegister(sourceErrorsTotal)
	prometheus.MustRegister(sourceEndpointsTotal)
	prometheus.MustRegister(registryEndpointsTotal)
	prometheus.MustRegister(lastSyncTimestamp)
	prometheus.MustRegister(deprecatedRegistryErrors)
	prometheus.MustRegister(deprecatedSourceErrors)
	prometheus.MustRegister(controllerNoChangesTotal)
	prometheus.MustRegister(registryARecords)
	prometheus.MustRegister(registryAAAARecords)
	prometheus.MustRegister(sourceARecords)
	prometheus.MustRegister(sourceAAAARecords)
	prometheus.MustRegister(verifiedARecords)
	prometheus.MustRegister(verifiedAAAARecords)
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
	// The nextRunAtMux is for atomic updating of nextRunAt
	nextRunAtMux sync.Mutex
	// DNS record types that will be considered for management
	ManagedRecordTypes []string
	// MinEventSyncInterval is used as window for batching events
	MinEventSyncInterval time.Duration
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {
	records, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Inc()
		deprecatedRegistryErrors.Inc()
		return err
	}

	missingRecords := c.Registry.MissingRecords()

	registryEndpointsTotal.Set(float64(len(records)))
	regARecords, regAAAARecords := countAddressRecords(records)
	registryARecords.Set(float64(regARecords))
	registryAAAARecords.Set(float64(regAAAARecords))
	ctx = context.WithValue(ctx, provider.RecordsContextKey, records)

	endpoints, err := c.Source.Endpoints(ctx)
	if err != nil {
		sourceErrorsTotal.Inc()
		deprecatedSourceErrors.Inc()
		return err
	}
	sourceEndpointsTotal.Set(float64(len(endpoints)))
	srcARecords, srcAAAARecords := countAddressRecords(endpoints)
	sourceARecords.Set(float64(srcARecords))
	sourceAAAARecords.Set(float64(srcAAAARecords))
	vARecords, vAAAARecords := countMatchingAddressRecords(endpoints, records)
	verifiedARecords.Set(float64(vARecords))
	verifiedAAAARecords.Set(float64(vAAAARecords))
	endpoints = c.Registry.AdjustEndpoints(endpoints)

	if len(missingRecords) > 0 {
		// Add missing records before the actual plan is applied.
		// This prevents the problems when the missing TXT record needs to be
		// created and deleted/upserted in the same batch.
		missingRecordsPlan := &plan.Plan{
			Policies:           []plan.Policy{c.Policy},
			Missing:            missingRecords,
			DomainFilter:       endpoint.MatchAllDomainFilters{c.DomainFilter, c.Registry.GetDomainFilter()},
			PropertyComparator: c.Registry.PropertyValuesEqual,
			ManagedRecords:     c.ManagedRecordTypes,
		}
		missingRecordsPlan = missingRecordsPlan.Calculate()
		if missingRecordsPlan.Changes.HasChanges() {
			err = c.Registry.ApplyChanges(ctx, missingRecordsPlan.Changes)
			if err != nil {
				registryErrorsTotal.Inc()
				deprecatedRegistryErrors.Inc()
				return err
			}
			log.Info("All missing records are created")
		}
	}

	plan := &plan.Plan{
		Policies:           []plan.Policy{c.Policy},
		Current:            records,
		Desired:            endpoints,
		DomainFilter:       endpoint.MatchAllDomainFilters{c.DomainFilter, c.Registry.GetDomainFilter()},
		PropertyComparator: c.Registry.PropertyValuesEqual,
		ManagedRecords:     c.ManagedRecordTypes,
	}

	plan = plan.Calculate()

	if plan.Changes.HasChanges() {
		err = c.Registry.ApplyChanges(ctx, plan.Changes)
		if err != nil {
			registryErrorsTotal.Inc()
			deprecatedRegistryErrors.Inc()
			return err
		}
	} else {
		controllerNoChangesTotal.Inc()
		log.Info("All records are already up to date")
	}

	lastSyncTimestamp.SetToCurrentTime()
	return nil
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
	c.nextRunAtMux.Lock()
	defer c.nextRunAtMux.Unlock()
	// schedule only if a reconciliation is not already planned
	// to happen in the following c.MinEventSyncInterval
	if !c.nextRunAt.Before(now.Add(c.MinEventSyncInterval)) {
		c.nextRunAt = now.Add(c.MinEventSyncInterval)
	}
}

func (c *Controller) ShouldRunOnce(now time.Time) bool {
	c.nextRunAtMux.Lock()
	defer c.nextRunAtMux.Unlock()
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
				log.Error(err)
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
