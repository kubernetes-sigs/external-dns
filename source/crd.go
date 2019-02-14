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
	"os"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/pkg/crd"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

// crdSource is an implementation of Source for Kubernetes CustomResourceDefinition objects.
type crdSource struct {
	DNSRecordStore cache.Store
	DNSZoneStore   cache.Store
}

// NewCRDSource returns an initilized CRD source
func NewCRDSource(kubeMaster, kubeConfig string) (Source, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(kubeMaster, kubeConfig)
	if err != nil {
		return nil, err
	}

	clientSet, err := crd.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	if err != nil {
		panic(err)
	}

	return crdSource{
		DNSRecordStore: crd.WatchDNSRecordResources(clientSet),
		DNSZoneStore:   crd.WatchDNSZoneResources(clientSet),
	}, nil
}

// Endpoints implements Source and returns endpoint objects from the CRD source DNSRecordStore.
func (s crdSource) Endpoints() ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	list := s.DNSRecordStore.List()

	for _, item := range list {
		for _, record := range item.(*crd.DNSRecord).Spec.Records {
			endpoints = append(endpoints, &endpoint.Endpoint{
				DNSName:          record.Name,
				Targets:          record.Targets,
				RecordType:       record.Type,
				RecordTTL:        record.TTL,
				ProviderSpecific: record.ProviderSpecificOptions,
			})
		}

	}
	return endpoints, nil
}
