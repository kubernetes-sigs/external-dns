/*
Copyright 2019 The Kubernetes Authors.

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

package source

import (
	"net/url"
	"time"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"sigs.k8s.io/external-dns/endpoint"
)

type cloudfoundrySource struct {
	client *cfclient.Client
}

// NewCloudFoundrySource creates a new cloudfoundrySource with the given config
func NewCloudFoundrySource(cfClient *cfclient.Client) (Source, error) {
	return &cloudfoundrySource{
		client: cfClient,
	}, nil
}

func (rs *cloudfoundrySource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}

// Endpoints returns endpoint objects
func (rs *cloudfoundrySource) Endpoints() ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	u, err := url.Parse(rs.client.Config.ApiAddress)
	if err != nil {
		panic(err)
	}

	domains, _ := rs.client.ListDomains()
	for _, domain := range domains {
		q := url.Values{}
		q.Set("q", "domain_guid:"+domain.Guid)
		routes, _ := rs.client.ListRoutesByQuery(q)
		for _, element := range routes {
			endpoints = append(endpoints,
				endpoint.NewEndpointWithTTL(element.Host+"."+domain.Name, endpoint.RecordTypeCNAME, 300, u.Host))
		}
	}

	return endpoints, nil
}
