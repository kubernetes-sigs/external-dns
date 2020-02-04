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

package source

import (
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// dedupSource is a Source that removes duplicate endpoints from its wrapped source.
type dedupSource struct {
	source Source
}

// NewDedupSource creates a new dedupSource wrapping the provided Source.
func NewDedupSource(source Source) Source {
	return &dedupSource{source: source}
}

// Endpoints collects endpoints from its wrapped source and returns them without duplicates.
func (ms *dedupSource) Endpoints() ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}
	collected := map[string]bool{}

	endpoints, err := ms.source.Endpoints()
	if err != nil {
		return nil, err
	}

	for _, ep := range endpoints {
		identifier := ep.DNSName + " / " + ep.SetIdentifier + " / " + ep.Targets.String()

		if _, ok := collected[identifier]; ok {
			log.Debugf("Removing duplicate endpoint %s", ep)
			continue
		}

		collected[identifier] = true
		result = append(result, ep)
	}

	return result, nil
}

func (ms *dedupSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
	ms.source.AddEventHandler(handler, stopChan, minInterval)
}
