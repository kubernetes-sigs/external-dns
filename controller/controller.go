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

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

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
	lastRunAt    time.Time
	EventEmitter events.EventEmitter
	// MangedRecordTypes are DNS record types that will be considered for management.
	ManagedRecordTypes []string
	// ExcludeRecordTypes are DNS record types that will be excluded from management.
	ExcludeRecordTypes []string
	// MinEventSyncInterval is used as a window for batching events
	MinEventSyncInterval time.Duration
	// Old txt-owner value we need to migrate from
	TXTOwnerOld string
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {
	lastReconcileTimestamp.Gauge.SetToCurrentTime()

	c.runAtMutex.Lock()
	c.lastRunAt = time.Now()
	c.runAtMutex.Unlock()

	regRecords, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Counter.Inc()
		deprecatedRegistryErrors.Counter.Inc()
		return err
	}

	registryEndpointsTotal.Gauge.Set(float64(len(regRecords)))

	countAddressRecords(regRecords, registryRecords)

	ctx = context.WithValue(ctx, provider.RecordsContextKey, regRecords)

	sourceEndpoints, err := c.Source.Endpoints(ctx)
	if err != nil {
		sourceErrorsTotal.Counter.Inc()
		deprecatedSourceErrors.Counter.Inc()
		return err
	}

	sourceEndpointsTotal.Gauge.Set(float64(len(sourceEndpoints)))

	countAddressRecords(sourceEndpoints, sourceRecords)
	countMatchingAddressRecords(sourceEndpoints, regRecords, verifiedRecords)

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
		OldOwnerId:     c.TXTOwnerOld,
	}

	plan = plan.Calculate()

	if plan.Changes.HasChanges() {
		err = c.Registry.ApplyChanges(ctx, plan.Changes)
		if err != nil {
			registryErrorsTotal.Counter.Inc()
			deprecatedRegistryErrors.Counter.Inc()
			emitChangeEvent(c.EventEmitter, plan.Changes, events.RecordError)
			return err
		}
		emitChangeEvent(c.EventEmitter, plan.Changes, events.RecordReady)
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
					log.Fatalf("Failed to do run once: %v", err) // nolint: gocritic // exitAfterDefer
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
