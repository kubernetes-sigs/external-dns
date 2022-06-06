/*
Copyright 2020 The Kubernetes Authors.

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

package stackpath

import (
	"context"
	"os"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

func TestNewStackPathProvider(t *testing.T) {
	stackPathConfig := &StackPathConfig{
		Context:      context.Background(),
		DomainFilter: endpoint.NewDomainFilter(nil),
		ZoneIDFilter: provider.NewZoneIDFilter(nil),
		DryRun:       false,
		Testing:      true,
	}

	_, err := NewStackPathProvider(*stackPathConfig)
	if err == nil {
		t.Fatalf("Expected to fail without a valid CLIENT_ID, CLIENT_SECRET, and STACK_ID")
	}

	_ = os.Setenv("STACKPATH_CLIENT_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	_ = os.Setenv("STACKPATH_CLIENT_SECRET", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	_ = os.Setenv("STACKPATH_STACK_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	_, err = NewStackPathProvider(*stackPathConfig)
	if err.Error() != "Error obtaining oauth2 token: 404 Not Found" {
		t.Fatalf("%v", err)
	}
}

// func TestTestTest(t *testing.T) {
// 	provider := &StackPathProvider{
// 		client:       &dns.APIClient{},
// 		context:      context.Background(),
// 		domainFilter: endpoint.DomainFilter{},
// 		zoneIdFilter: provider.ZoneIDFilter{},
// 		stackId:      "",
// 		dryRun:       true,
//      testing:      true
// 	}

// 	a, b, c := provider.getZoneRecords("test.com")

// 	if a != dns.ZoneGetZoneRecordsResponse {
// 	}
// 	{
// 		t.Fatalf("Expected empty response")
// 	}
// 	spew.Dump(b)
// 	spew.Dump(c)

// }

// func TestStackPathRecords(t *testing.T) {
// 	mocked := mockStackPathProvider{}

// 	provider := &StackPathProvider{
// 		client: mocked,
// 	}

// 	_, err := provider.zones()
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}

// }
