/*
Copyright 2025 The Kubernetes Authors.

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

package controller

import "sigs.k8s.io/external-dns/endpoint"

type metricsRecorder struct {
	counterPerEndpointType map[string]int
}

func newMetricsRecorder() *metricsRecorder {
	return &metricsRecorder{
		counterPerEndpointType: map[string]int{
			endpoint.RecordTypeA:     0,
			endpoint.RecordTypeAAAA:  0,
			endpoint.RecordTypeCNAME: 0,
			endpoint.RecordTypeTXT:   0,
			endpoint.RecordTypeSRV:   0,
			endpoint.RecordTypeNS:    0,
			endpoint.RecordTypePTR:   0,
			endpoint.RecordTypeMX:    0,
			endpoint.RecordTypeNAPTR: 0,
		},
	}
}

func (m *metricsRecorder) recordEndpointType(endpointType string) {
	m.counterPerEndpointType[endpointType]++
}

func (m *metricsRecorder) getEndpointTypeCount(endpointType string) int {
	if count, ok := m.counterPerEndpointType[endpointType]; ok {
		return count
	}
	return 0
}

func (m *metricsRecorder) loadFloat64(endpointType string) float64 {
	return float64(m.getEndpointTypeCount(endpointType))
}
