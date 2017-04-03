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

package storage

import (
	"testing"
	"time"

	"fmt"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
)

func TestPoll(t *testing.T) {
	oldResyncPeriod := resyncPeriod
	stopChan := make(chan struct{}, 1)
	count := 3
	numErr := 1
	resyncPeriod = 1 * time.Second
	defer func() {
		resyncPeriod = oldResyncPeriod
	}()

	syncMock := func() error {
		count--
		if count == 2 {
			return fmt.Errorf("error: dummy")
		}
		if count == 0 {
			stopChan <- struct{}{}
		}
		return nil
	}
	onSyncError := func(err error) {
		numErr--
	}

	checkFunc := func() {
		if count != 0 {
			t.Fatalf("Poll SyncFunc was not called %d times", 3)
		}
		if numErr != 0 {
			t.Fatalf("Error callback was not triggered")
		}
	}
	time.AfterFunc(time.Second*time.Duration(count+1), checkFunc)

	Poll(syncMock, onSyncError, stopChan)
	// <-stopChan
}

func TestUpdatedCache(t *testing.T) {
	for _, ti := range []struct {
		title        string
		records      []endpoint.Endpoint
		cacheRecords []*endpoint.SharedEndpoint
		expected     []*endpoint.SharedEndpoint
	}{
		{
			title:        "all empty",
			records:      []endpoint.Endpoint{},
			cacheRecords: []*endpoint.SharedEndpoint{},
			expected:     []*endpoint.SharedEndpoint{},
		},
		{
			title:   "no records, should produce empty cache",
			records: []endpoint.Endpoint{},
			cacheRecords: []*endpoint.SharedEndpoint{
				{},
			},
			expected: []*endpoint.SharedEndpoint{},
		},
		{
			title: "new records, empty cache",
			records: []endpoint.Endpoint{
				{
					DNSName: "foo.org",
					Target:  "elb.com",
				},
				{
					DNSName: "bar.org",
					Target:  "alb.com",
				},
			},
			cacheRecords: []*endpoint.SharedEndpoint{
				{},
			},
			expected: []*endpoint.SharedEndpoint{
				{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "alb.com",
					},
				},
			},
		},
		{
			title: "new records, non-empty cache",
			records: []endpoint.Endpoint{
				{
					DNSName: "foo.org",
					Target:  "elb.com",
				},
				{
					DNSName: "bar.org",
					Target:  "alb.com",
				},
				{
					DNSName: "owned.org",
					Target:  "8.8.8.8",
				},
			},
			cacheRecords: []*endpoint.SharedEndpoint{
				{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "owned.org",
						Target:  "8.8.8.8",
					},
				},
				{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "to-be-deleted.org",
						Target:  "52.53.54.55",
					},
				},
			},
			expected: []*endpoint.SharedEndpoint{
				{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "alb.com",
					},
				},
				{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "owned.org",
						Target:  "8.8.8.8",
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			if !testutils.SameSharedEndpoints(updatedCache(ti.records, ti.cacheRecords), ti.expected) {
				t.Errorf("incorrect result produced by updatedCache")
			}
		})
	}
}
