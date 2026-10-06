/*
Copyright 2026 The Kubernetes Authors.

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

package crd

import "sigs.k8s.io/controller-runtime/pkg/client"

// CRDClients bundles the clients built for the crd source.
type CRDClients struct {
	// reader serves reads from the informer cache.
	reader client.Reader
	// writer bypasses the cache; every request goes straight to the API server.
	writer client.Client
}

// NewCRDClients bundles reader and writer into a CRDClients.
func NewCRDClients(reader client.Reader, writer client.Client) *CRDClients {
	return &CRDClients{reader: reader, writer: writer}
}

// Reader returns the cache-backed client.
func (c *CRDClients) Reader() client.Reader {
	return c.reader
}

// Writer returns the client that bypasses the cache.
func (c *CRDClients) Writer() client.Client {
	return c.writer
}
