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

	registryRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      "records",
			Help:      "Number of registry records partitioned by label name (vector).",
		},
		[]string{"record_type"},
	)

	sourceRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      "records",
			Help:      "Number of source records partitioned by label name (vector).",
		},
		[]string{"record_type"},
	)

	verifiedRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "verified_records",
			Help:      "Number of DNS records that exists both in source and registry (vector).",
		},
		[]string{"record_type"},
	)

	consecutiveSoftErrors = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller",
			Name:      "consecutive_soft_errors",
			Help:      "Number of consecutive soft errors in reconciliation loop.",
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

	metrics.RegisterMetric.MustRegister(registryRecords)
	metrics.RegisterMetric.MustRegister(sourceRecords)
	metrics.RegisterMetric.MustRegister(verifiedRecords)

	metrics.RegisterMetric.MustRegister(consecutiveSoftErrors)
}

// Controller is responsible for orchestrating the different components.
// It works in the following way:
// * Ask the DNS provider for the current list of endpoints.
// * Ask the Source for the desired list of endpoints.
// * Take both lists and calculate a Plan to move current towards the desired state.
// * Tell the DNS provider to apply the changes calculated by the Plan.
type Controller struct {
	Source   source.Source
	Registry registry.Registry
	// The policy that defines which change to DNS records is allowed
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
	// MinEventSyncInterval is used as a window for batching events
	MinEventSyncInterval time.Duration
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {
	lastReconcileTimestamp.Gauge.SetToCurrentTime()

	c.runAtMutex.Lock()
	c.lastRunAt = time.Now()
	c.runAtMutex.Unlock()

	regMetrics := newMetricsRecorder()

	regRecords, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Counter.Inc()
		deprecatedRegistryErrors.Counter.Inc()
		return err
	}

	registryEndpointsTotal.Gauge.Set(float64(len(regRecords)))

	countAddressRecords(regMetrics, regRecords, registryRecords)

	ctx = context.WithValue(ctx, provider.RecordsContextKey, regRecords)

	sourceEndpoints, err := c.Source.Endpoints(ctx)
	if err != nil {
		sourceErrorsTotal.Counter.Inc()
		deprecatedSourceErrors.Counter.Inc()
		return err
	}

	sourceEndpointsTotal.Gauge.Set(float64(len(sourceEndpoints)))

	sourceMetrics := newMetricsRecorder()
	countAddressRecords(sourceMetrics, sourceEndpoints, sourceRecords)

	vaMetrics := newMetricsRecorder()
	countMatchingAddressRecords(vaMetrics, sourceEndpoints, regRecords, verifiedRecords)

	endpoints, err := c.Registry.AdjustEndpoints(sourceEndpoints)
	if err != nil {
		return fmt.Errorf("adjusting endpoints: %w", err)
	}
	registryFilter := c.Registry.GetDomainFilter()

	plan := &plan.Plan{
		Policies:       []plan.Policy{c.Policy},
		Current:        regRecords,
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

// Counts the intersections of records in endpoint and registry.
func countMatchingAddressRecords(rec *metricsRecorder, endpoints []*endpoint.Endpoint, registryRecords []*endpoint.Endpoint, metric metrics.GaugeVecMetric) {
	recordsMap := make(map[string]map[string]struct{})
	for _, regRecord := range registryRecords {
		if _, found := recordsMap[regRecord.DNSName]; !found {
			recordsMap[regRecord.DNSName] = make(map[string]struct{})
		}
		recordsMap[regRecord.DNSName][regRecord.RecordType] = struct{}{}
	}

	for _, sourceRecord := range endpoints {
		if _, found := recordsMap[sourceRecord.DNSName]; found {
			if _, ok := recordsMap[sourceRecord.DNSName][sourceRecord.RecordType]; ok {
				rec.recordEndpointType(sourceRecord.RecordType)
			}
		}
	}

	for _, rt := range endpoint.KnownRecordTypes {
		metric.SetWithLabels(rec.loadFloat64(rt), rt)
	}
}

// countAddressRecords updates the metricsRecorder with the count of each record type
// found in the provided endpoints slice, and sets the corresponding metrics for each
// known DNS record type using the sourceRecords metric.
func countAddressRecords(rec *metricsRecorder, endpoints []*endpoint.Endpoint, metric metrics.GaugeVecMetric) {
	// compute the number of records per type
	for _, endPoint := range endpoints {
		rec.recordEndpointType(endPoint.RecordType)
	}
	// set metrics for each record type
	for _, rt := range endpoint.KnownRecordTypes {
		metric.SetWithLabels(rec.loadFloat64(rt), rt)
	}
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
	var softErrorCount int
	for {
		if c.ShouldRunOnce(time.Now()) {
			if err := c.RunOnce(ctx); err != nil {
				if errors.Is(err, provider.SoftError) {
					softErrorCount++
					consecutiveSoftErrors.Gauge.Set(float64(softErrorCount))
					log.Errorf("Failed to do run once: %v (consecutive soft errors: %d)", err, softErrorCount)
				} else {
					log.Fatalf("Failed to do run once: %v", err)
				}
			} else {
				if softErrorCount > 0 {
					log.Infof("Reconciliation succeeded after %d consecutive soft errors", softErrorCount)
				}
				softErrorCount = 0
				consecutiveSoftErrors.Gauge.Set(0)
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
