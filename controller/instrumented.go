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
)

var _ Controller = &InstrumentedController{}

type InstrumentedController struct {
	ctrl Controller
}

var (
	SynchronizationLatencies = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  "externaldns",
			Name:       "synchronization_latencies",
			Help:       "Synchronization latencies summary in seconds for successful and non-successful synchronizations.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(SynchronizationLatencies)
}

func NewInstrumentedController(ctrl Controller) Controller {
	return &InstrumentedController{ctrl: ctrl}
}

func (c *InstrumentedController) Run() error {
	var err error

	defer func(then time.Time) {
		SynchronizationLatencies.
			WithLabelValues(statusLabel(err)).
			Observe(time.Since(then).Seconds())
	}(time.Now())

	// We need to store err to give defer a chance to read it.
	err = c.ctrl.Run()

	return err
}

func statusLabel(err error) string {
	if err != nil {
		return "error"
	}
	return "success"
}
