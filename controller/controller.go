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
	"strings"
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
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeA)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeA),
		},
	)
	registryAAAARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeAAAA)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeAAAA),
		},
	)
	registryCnameRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeCNAME)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeCNAME),
		},
	)
	registryTXTRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeTXT)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeTXT),
		},
	)
	registrySRVRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeSRV)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeSRV),
		},
	)
	registryNSRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeNS)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeNS),
		},
	)
	registryPTRRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypePTR)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypePTR),
		},
	)
	registryMXRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeMX)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeMX),
		},
	)
	registryNAPTRRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "registry",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeNAPTR)),
			Help:      fmt.Sprintf("Number of Registry %s records.", endpoint.RecordTypeNAPTR),
		},
	)
	sourceARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeA)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeA),
		},
	)
	sourceAAAARecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeAAAA)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeAAAA),
		},
	)
	sourceCnameRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeCNAME)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeCNAME),
		},
	)
	sourceTXTRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeTXT)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeTXT),
		},
	)
	sourceSRVRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeSRV)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeSRV),
		},
	)
	sourceNSRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeNS)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeNS),
		},
	)
	sourcePTRRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypePTR)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypePTR),
		},
	)
	sourceMXRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeMX)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeMX),
		},
	)
	sourceNAPTRRecords = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "source",
			Name:      fmt.Sprintf("%s_records", strings.ToLower(endpoint.RecordTypeNAPTR)),
			Help:      fmt.Sprintf("Number of Source %s records.", endpoint.RecordTypeNAPTR),
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

	metrics.RegisterMetric.MustRegister(registryARecords)
	metrics.RegisterMetric.MustRegister(registryAAAARecords)
	metrics.RegisterMetric.MustRegister(registryCnameRecords)
	metrics.RegisterMetric.MustRegister(registryTXTRecords)
	metrics.RegisterMetric.MustRegister(registrySRVRecords)
	metrics.RegisterMetric.MustRegister(registryNSRecords)
	metrics.RegisterMetric.MustRegister(registryPTRRecords)
	metrics.RegisterMetric.MustRegister(registryMXRecords)
	metrics.RegisterMetric.MustRegister(registryNAPTRRecords)

	metrics.RegisterMetric.MustRegister(sourceARecords)
	metrics.RegisterMetric.MustRegister(sourceAAAARecords)
	metrics.RegisterMetric.MustRegister(sourceCnameRecords)
	metrics.RegisterMetric.MustRegister(sourceTXTRecords)
	metrics.RegisterMetric.MustRegister(sourceSRVRecords)
	metrics.RegisterMetric.MustRegister(sourceNSRecords)
	metrics.RegisterMetric.MustRegister(sourcePTRRecords)
	metrics.RegisterMetric.MustRegister(sourceMXRecords)
	metrics.RegisterMetric.MustRegister(sourceNAPTRRecords)

	metrics.RegisterMetric.MustRegister(verifiedARecords)
	metrics.RegisterMetric.MustRegister(verifiedAAAARecords)
	metrics.RegisterMetric.MustRegister(consecutiveSoftErrors)
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

	regMetrics := newMetricsRecorder()

	records, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Counter.Inc()
		deprecatedRegistryErrors.Counter.Inc()
		return err
	}

	registryEndpointsTotal.Gauge.Set(float64(len(records)))

	countAddressRecords(regMetrics, records)
	registryARecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeA))
	registryAAAARecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeAAAA))
	registryCnameRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeCNAME))
	registryTXTRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeTXT))
	registrySRVRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeSRV))
	registryNSRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeNS))
	registryPTRRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypePTR))
	registryMXRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeMX))
	registryNAPTRRecords.Gauge.Set(regMetrics.loadFloat64(endpoint.RecordTypeNAPTR))

	ctx = context.WithValue(ctx, provider.RecordsContextKey, records)

	endpoints, err := c.Source.Endpoints(ctx)
	if err != nil {
		sourceErrorsTotal.Counter.Inc()
		deprecatedSourceErrors.Counter.Inc()
		return err
	}
	sourceEndpointsTotal.Gauge.Set(float64(len(endpoints)))

	sourceMetrics := newMetricsRecorder()
	countAddressRecords(sourceMetrics, endpoints)
	sourceARecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeA))
	sourceAAAARecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeAAAA))
	sourceCnameRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeCNAME))
	sourceTXTRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeTXT))
	sourceSRVRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeSRV))
	sourceNSRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeNS))
	sourcePTRRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypePTR))
	sourceMXRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeMX))
	sourceNAPTRRecords.Gauge.Set(sourceMetrics.loadFloat64(endpoint.RecordTypeNAPTR))

	vaMetrics := newMetricsRecorder()
	countMatchingAddressRecords(vaMetrics, endpoints, records)
	verifiedARecords.Gauge.Set(vaMetrics.loadFloat64(endpoint.RecordTypeA))
	verifiedAAAARecords.Gauge.Set(vaMetrics.loadFloat64(endpoint.RecordTypeAAAA))

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

// Counts the intersections of records in endpoint and registry.
func countMatchingAddressRecords(rec *metricsRecorder, endpoints []*endpoint.Endpoint, registryRecords []*endpoint.Endpoint) {
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
}

func countAddressRecords(rec *metricsRecorder, endpoints []*endpoint.Endpoint) {
	for _, endPoint := range endpoints {
		rec.recordEndpointType(endPoint.RecordType)
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
