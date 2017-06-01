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

import "github.com/kubernetes-incubator/external-dns/plan"

// BaseController is responsible for orchestrating the different components.
// It works in the following way:
// * Ask the DNS provider for current list of endpoints.
// * Ask the Source for the desired list of endpoints.
// * Take both lists and calculate a Plan to move current towards desired state under a given policy.
// * Tell the DNS provider to apply the changes calculated by the Plan.
type BaseController struct {
	config Config
}

func NewBaseController(cfg Config) Controller {
	return &BaseController{cfg}
}

func (c *BaseController) Run() error {
	records, err := c.config.Registry.Records()
	if err != nil {
		return err
	}

	endpoints, err := c.config.Source.Endpoints()
	if err != nil {
		return err
	}

	plan := &plan.Plan{
		Policies: []plan.Policy{c.config.Policy},
		Current:  records,
		Desired:  endpoints,
	}

	plan = plan.Calculate()

	return c.config.Registry.ApplyChanges(plan.Changes)
}
