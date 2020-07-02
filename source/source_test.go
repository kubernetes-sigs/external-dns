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

package source

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestGetTTLFromAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title       string
		annotations map[string]string
		expectedTTL endpoint.TTL
		expectedErr error
	}{
		{
			title:       "TTL annotation not present",
			annotations: map[string]string{"foo": "bar"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is not a number",
			annotations: map[string]string{ttlAnnotationKey: "foo"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("\"foo\" is not a valid TTL value"),
		},
		{
			title:       "TTL annotation value is empty",
			annotations: map[string]string{ttlAnnotationKey: ""},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("\"\" is not a valid TTL value"),
		},
		{
			title:       "TTL annotation value is negative number",
			annotations: map[string]string{ttlAnnotationKey: "-1"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("TTL value must be between [%d, %d]", ttlMinimum, ttlMaximum),
		},
		{
			title:       "TTL annotation value is too high",
			annotations: map[string]string{ttlAnnotationKey: fmt.Sprintf("%d", 1<<32)},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("TTL value must be between [%d, %d]", ttlMinimum, ttlMaximum),
		},
		{
			title:       "TTL annotation value is set correctly using integer",
			annotations: map[string]string{ttlAnnotationKey: "60"},
			expectedTTL: endpoint.TTL(60),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is set correctly using duration (whole)",
			annotations: map[string]string{ttlAnnotationKey: "10m"},
			expectedTTL: endpoint.TTL(600),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is set correcly using duration (fractional)",
			annotations: map[string]string{ttlAnnotationKey: "20.5s"},
			expectedTTL: endpoint.TTL(20),
			expectedErr: nil,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			ttl, err := getTTLFromAnnotations(tc.annotations)
			assert.Equal(t, tc.expectedTTL, ttl)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestSuitableType(t *testing.T) {
	for _, tc := range []struct {
		target, recordType, expected string
	}{
		{"8.8.8.8", "", "A"},
		{"foo.example.org", "", "CNAME"},
		{"bar.eu-central-1.elb.amazonaws.com", "", "CNAME"},
	} {

		recordType := suitableType(tc.target)

		if recordType != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, recordType)
		}
	}
}
