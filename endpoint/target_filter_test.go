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

package endpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type targetFilterTest struct {
	stubEndpoint *Endpoint
	targetFilter []string
	exclusions   []string
	targets      []string
	expected     bool
}

var targetFilterTestsStubEndpointNetwork = NewEndpoint("test-a.example.com", RecordTypeA)
var targetFilterTestsStubEndpointCname = NewEndpoint("test-cname.example.com", RecordTypeCNAME)
var targetFilterTests = []targetFilterTest{
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{"10.0.0.0/8"},
		[]string{},
		[]string{"10.1.2.3"},
		true,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{" 10.0.0.0/8 "},
		[]string{},
		[]string{"10.1.2.3"},
		true,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{"0"},
		[]string{},
		[]string{"10.1.2.3"},
		true,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{"10.0.0.0/8"},
		[]string{},
		[]string{"1.1.1.1"},
		false,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{},
		[]string{"10.0.0.0/8"},
		[]string{"1.1.1.1"},
		true,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{},
		[]string{"10.0.0.0/8"},
		[]string{"10.1.2.3"},
		false,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{},
		[]string{"10.0.0.0/8"},
		[]string{"49.13.41.161"},
		true,
	},
	{
		targetFilterTestsStubEndpointNetwork,
		[]string{},
		[]string{"10.0.0.0/8"},
		[]string{"10.0.1.101"},
		false,
	},
	{targetFilterTestsStubEndpointCname,
		[]string{"10.0.0.0/8"},
		[]string{"10.1.0.0/24"},
		[]string{"10.2.254.254", "10.1.1.1", "cname-1.example.com", "random text data"},
		true,
	},
}

func TestTargetFilterWithExclusions(t *testing.T) {
	for i, tt := range targetFilterTests {
		if len(tt.exclusions) == 0 {
			tt.exclusions = append(tt.exclusions, "")
		}
		targetFilter := NewTargetNetFilterWithExclusions(tt.targetFilter, tt.exclusions)
		for _, target := range tt.targets {
			assert.Equal(t, tt.expected, targetFilter.Match(target, tt.stubEndpoint), "should not fail: %v in test-case #%v", target, i)
		}
	}
}

func TestTargetFilterMatchWithEmptyFilter(t *testing.T) {
	for _, tt := range targetFilterTests {
		targetFilter := TargetNetFilter{}
		for i, target := range tt.targets {
			assert.True(t, targetFilter.Match(target, tt.stubEndpoint), "should not fail: %v in test-case #%v", target, i)
		}
	}
}

func TestTargetNetFilter_IsEnabled(t *testing.T) {
	tests := []struct {
		name        string
		filterNets  []string
		excludeNets []string
		want        bool
	}{
		{"both empty", []string{}, []string{}, false},
		{"filterNets non-empty", []string{"10.0.0.0/8"}, []string{}, true},
		{"excludeNets non-empty", []string{}, []string{"10.0.0.0/8"}, true},
		{"both non-empty", []string{"10.0.0.0/8"}, []string{"192.168.0.0/16"}, true},
	}

	for _, tt := range tests {
		tf := NewTargetNetFilterWithExclusions(tt.filterNets, tt.excludeNets)
		assert.Equal(t, tt.want, tf.IsEnabled())
	}
}
