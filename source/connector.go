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
	"encoding/gob"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	dialTimeout = 30 * time.Second
)

// connectorSource is an implementation of Source that provides endpoints by connecting
// to a remote tcp server. The encoding/decoding is done using encoder/gob package.
type connectorSource struct {
	remoteServer string
}

// NewConnectorSource creates a new connectorSource with the given config.
func NewConnectorSource(remoteServer string) (Source, error) {
	return &connectorSource{
		remoteServer: remoteServer,
	}, nil
}

// Endpoints returns endpoint objects.
func (cs *connectorSource) Endpoints() ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	conn, err := net.DialTimeout("tcp", cs.remoteServer, dialTimeout)
	if err != nil {
		log.Errorf("Connection error: %v", err)
		return nil, err
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	if err := decoder.Decode(&endpoints); err != nil {
		log.Errorf("Decode error: %v", err)
		return nil, err
	}

	log.Debugf("Received endpoints: %#v", endpoints)

	return endpoints, nil
}

func (cs *connectorSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}
