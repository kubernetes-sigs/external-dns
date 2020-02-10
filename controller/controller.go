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
)

func init() {
	prometheus.MustRegister(registryErrorsTotal)
	prometheus.MustRegister(sourceErrorsTotal)
	prometheus.MustRegister(sourceEndpointsTotal)
	prometheus.MustRegister(registryEndpointsTotal)
	prometheus.MustRegister(deprecatedRegistryErrors)
	prometheus.MustRegister(deprecatedSourceErrors)
}

// Controller is responsible for orchestrating the different components.
// It works in the following way:
// * Ask the DNS provider for current list of endpoints.
// * Ask the Source for the desired list of endpoints.
// * Take both lists and calculate a Plan to move current towards desired state.
// * Tell the DNS provider to apply the changes calucated by the Plan.
type Controller struct {
	Source   source.Source
	Registry registry.Registry
	// The policy that defines which changes to DNS records are allowed
	Policy plan.Policy
	// The interval between individual synchronizations
	Interval time.Duration
	// The DomainFilter defines which DNS records to keep or exclude
	DomainFilter endpoint.DomainFilter
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {
	records, err := c.Registry.Records(ctx)
	if err != nil {
		registryErrorsTotal.Inc()
		deprecatedRegistryErrors.Inc()
		return err
	}
	registryEndpointsTotal.Set(float64(len(records)))

	ctx = context.WithValue(ctx, provider.RecordsContextKey, records)

	endpoints, err := c.Source.Endpoints()
	if err != nil {
		sourceErrorsTotal.Inc()
		deprecatedSourceErrors.Inc()
		return err
	}
	sourceEndpointsTotal.Set(float64(len(endpoints)))

	plan := &plan.Plan{
		Policies:     []plan.Policy{c.Policy},
		Current:      records,
		Desired:      endpoints,
		DomainFilter: c.DomainFilter,
	}

	plan = plan.Calculate()

	err = c.Registry.ApplyChanges(ctx, plan.Changes)
	if err != nil {
		registryErrorsTotal.Inc()
		deprecatedRegistryErrors.Inc()
		return err
	}
	return nil
}

// Run runs RunOnce in a loop with a delay until stopChan receives a value.
func (c *Controller) Run(ctx context.Context, stopChan <-chan struct{}) {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for {
		err := c.RunOnce(ctx)
		if err != nil {
			log.Error(err)
		}
		select {
		case <-ticker.C:
		case <-stopChan:
			log.Info("Terminating main controller loop")
			return
		}
	}
}
