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

	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/registry"
	"github.com/kubernetes-incubator/external-dns/source"
)

// Controller implementations must have a single method called Run() returning an Error.
type Controller interface {
	Run() error
}

// Config holds the shared configuration options for all Controller instances.
type Config struct {
	// Source defines where endpoints are coming from.
	Source source.Source
	// Registry is where the endpoints will be created and ownership is tracked.
	Registry registry.Registry
	// Policy defines which changes to DNS records are allowed.
	Policy plan.Policy
	// Interval defines the delay between individual synchronizations.
	Interval time.Duration
	// StopChan can be used to tell a controller to stop whatever it's doing.
	StopChan chan struct{}
}
