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

package crd

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

// WatchDNSZoneResources creates a Store that watches for DNSZone resource changes
func WatchDNSZoneResources(clientSet ExternalDNSV1Alpha1Interface) cache.Store {
	DNSZoneStore, DNSZoneController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.DNSZones("").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.DNSZones("").Watch(lo)
			},
		},
		&DNSZone{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go DNSZoneController.Run(wait.NeverStop)
	return DNSZoneStore
}

// WatchDNSRecordResources creates a Store that watches for DNSRecord resource changes
func WatchDNSRecordResources(clientSet ExternalDNSV1Alpha1Interface) cache.Store {
	DNSRecordStore, DNSRecordController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.DNSRecords("").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.DNSRecords("").Watch(lo)
			},
		},
		&DNSRecord{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)
	go DNSRecordController.Run(wait.NeverStop)
	return DNSRecordStore
}
