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

package api

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBluecatNewGatewayClient(t *testing.T) {
	testCookie := http.Cookie{Name: "testCookie", Value: "exampleCookie"}
	testToken := "exampleToken"
	testgateWayHost := "exampleHost"
	testDNSConfiguration := "exampleDNSConfiguration"
	testDNSServer := "exampleServer"
	testView := "testView"
	testZone := "example.com"
	testVerify := true

	client := NewGatewayClientConfig(testCookie, testToken, testgateWayHost, testDNSConfiguration, testView, testZone, testDNSServer, testVerify)

	if client.Cookie.Value != testCookie.Value || client.Cookie.Name != testCookie.Name || client.Token != testToken || client.Host != testgateWayHost || client.DNSConfiguration != testDNSConfiguration || client.View != testView || client.RootZone != testZone || client.SkipTLSVerify != testVerify {
		t.Fatal("Client values dont match")
	}
}

func TestBluecatExpandZones(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"with subdomain":        {input: "example.com", want: "zones/com/zones/example/zones/"},
		"only top level domain": {input: "com", want: "zones/com/zones/"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := expandZone(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestBluecatValidDeployTypes(t *testing.T) {
	validTypes := []string{"no-deploy", "full-deploy"}
	invalidTypes := []string{"anything-else"}
	for _, i := range validTypes {
		if !IsValidDNSDeployType(i) {
			t.Fatalf("%s should be a valid deploy type", i)
		}
	}
	for _, i := range invalidTypes {
		if IsValidDNSDeployType(i) {
			t.Fatalf("%s should be a invalid deploy type", i)
		}
	}
}

// TODO: Add error checking in case "properties" are not properly formatted
// Example test case... "invalid": {input: "abcde", want: map[string]string{}, err: InvalidProperty},
func TestBluecatSplitProperties(t *testing.T) {
	tests := map[string]struct {
		input string
		want  map[string]string
	}{
		"simple": {input: "ab=cd|ef=gh", want: map[string]string{"ab": "cd", "ef": "gh"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SplitProperties(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
