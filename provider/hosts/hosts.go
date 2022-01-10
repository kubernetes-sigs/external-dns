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

package hosts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "context"

  "sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	log "github.com/sirupsen/logrus"
)

// HostsProvider is an implementation of Provider for hosts file based Provider.
type HostsProvider struct {
  provider.BaseProvider
	DryRun   bool
	hostsFile string
}

// NewHostsProvider initializes a new hosts file based Provider.
func NewHostsProvider(ctx context.Context, hostsFile string, DryRun bool) (*HostsProvider, error) {
	log.Debugf("NewHostsProvider: hostsFile %s", hostsFile)

	p := &HostsProvider{
		DryRun:    DryRun,
		hostsFile: hostsFile,
	}

	return p, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *HostsProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("ApplyChanges")

	f, err := os.Create(p.hostsFile)
	if err != nil {
		log.Fatalf("Cannot open file %s", p.hostsFile)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for _, ep := range changes.Create {
		log.Debugf("Create")
		if ep.RecordType == "A" {
			writeHostsEntry(ep, w)
		}
	}

	for _, ep := range changes.UpdateNew {
		log.Debugf("UpdateNew: DNSName=%s RecordType=%s", ep.DNSName, ep.RecordType)
	}

	for _, ep := range changes.Delete {
		log.Debugf("Delete: DNSName=%s RecordType=%s", ep.DNSName, ep.RecordType)
	}

	return nil
}

// Records returns the list of records in a given hosted zone.
func (p *HostsProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debugf("Records")
	var records []*endpoint.Endpoint
	return records, nil
}

// writeHostsEntry writes hosts file with the host addresses and names
func writeHostsEntry(ep *endpoint.Endpoint, w *bufio.Writer) {
	addresses := make(map[string][]string)

	// get map of hostnames per address
	for _, host := range ep.Targets {
		addresses[host] = append(addresses[host], ep.DNSName)
	}

	// write hosts file
	for address, hosts := range addresses {
		log.Debugf("%s\t%s", address, strings.Join(hosts, " "))
		fmt.Fprintf(w, "%s\t%s\n", address, strings.Join(hosts, " "))
	}

	w.Flush()
}
