/*
Copyright 2018 The Kubernetes Authors.

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
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// httpClientSource is an implementation of Source that provides endpoints
// by requesting them from an HTTP server in a JSON format.
type httpClientSource struct {
	url string
}

var _ Source = (*httpClientSource)(nil)

// NewHTTPClientSource creates a new httpClientSource with the given config.
func NewHTTPClientSource(url string) (Source, error) {
	return &httpClientSource{
		url: url,
	}, nil
}

// Endpoints returns endpoint objects.
func (hcs *httpClientSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	res, err := http.Get(hcs.url)
	if err != nil {
		log.Errorf("Request error: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	endpoints := []*endpoint.Endpoint{}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&endpoints); err != nil {
		log.Errorf("Decode error: %v", err)
		return nil, err
	}

	log.Debugf("Received endpoints: %#v", endpoints)

	return endpoints, nil
}

func (hcs *httpClientSource) AddEventHandler(ctx context.Context, handler func()) {
}
