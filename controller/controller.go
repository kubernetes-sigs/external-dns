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
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/registry"
	"github.com/kubernetes-incubator/external-dns/source"
)

var (
	registryErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "registry_errors_total",
			Help: "Number of Registry errors.",
		},
	)
	sourceErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "source_errors_total",
			Help: "Number of Source errors.",
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
)

func init() {
	prometheus.MustRegister(registryErrors)
	prometheus.MustRegister(sourceErrors)
	prometheus.MustRegister(sourceEndpointsTotal)
	prometheus.MustRegister(registryEndpointsTotal)
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
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce() error {
	records, err := c.Registry.Records()
	if err != nil {
		registryErrors.Inc()
		return err
	}
	registryEndpointsTotal.Set(float64(len(records)))

	endpoints, err := c.Source.Endpoints()
	if err != nil {
		sourceErrors.Inc()
		return err
	}
	sourceEndpointsTotal.Set(float64(len(endpoints)))

	plan := &plan.Plan{
		Policies: []plan.Policy{c.Policy},
		Current:  records,
		Desired:  endpoints,
	}

	plan = plan.Calculate()

	err = c.Registry.ApplyChanges(plan.Changes)
	if err != nil {
		registryErrors.Inc()
		return err
	}
	return nil
}

// Run runs RunOnce in a loop with a delay until stopChan receives a value.
func (c *Controller) Run(stopChan <-chan struct{}) {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for {
		err := c.RunOnce()
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
