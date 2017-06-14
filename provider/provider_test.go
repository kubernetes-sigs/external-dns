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

package provider

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

func TestSuitableType(t *testing.T) {
	for _, tc := range []struct {
		target, recordType, expected string
	}{
		{"8.8.8.8", "", "A"},
		{"foo.example.org", "", "CNAME"},
		{"foo.example.org", "ALIAS", "ALIAS"},
		{"bar.eu-central-1.elb.amazonaws.com", "CNAME", "CNAME"},
	} {
		ep := &endpoint.Endpoint{
			Target:     tc.target,
			RecordType: tc.recordType,
		}

		recordType := suitableType(ep)

		if recordType != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, recordType)
		}
	}
}

func TestEnsureTrailingDot(t *testing.T) {
	for _, tc := range []struct {
		input, expected string
	}{
		{"example.org", "example.org."},
		{"example.org.", "example.org."},
		{"8.8.8.8", "8.8.8.8"},
	} {
		output := ensureTrailingDot(tc.input)

		if output != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, output)
		}
	}
}
