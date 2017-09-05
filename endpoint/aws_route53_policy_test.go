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

func TestRoute53WeightPolicy(t *testing.T) {
	t.Run("AWSRoute53Policy", TestNewAWSRoute53Policy)
}

func TestNewAWSRoute53Policy(t *testing.T) {
	for _, ti := range []struct {
		title         string
		weight        int64
		setIdentifier string
		expectError   bool
	}{
		{
			title:         "invalid weight",
			weight:        -1,
			setIdentifier: "cluster/id",
			expectError:   true,
		},
		{
			title:         "invalid weight",
			weight:        256,
			setIdentifier: "cluster/id",
			expectError:   true,
		},
		{
			title:  "invalid set identifier",
			weight: 1,
			setIdentifier: "11111111111111111111111111111111111111111111111111111111111111111111111111111" +
				"1111111111111111111111111111111111111111111111111111",
			expectError: true,
		},
		{
			title:         "invalid set identifier",
			weight:        1,
			setIdentifier: "",
			expectError:   true,
		},
		{
			title:         "valid weight and set identifier",
			weight:        1,
			setIdentifier: "cluster/id",
			expectError:   false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			_, err := NewAWSRoute53Policy(
				ti.weight,
				ti.setIdentifier,
			)
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAWSRoute53Policy_SetIdentifier(t *testing.T) {
	for _, ti := range []struct {
		title         string
		weight        int64
		setIdentifier string
		expected      string
	}{
		{
			title:         "valid weight ID",
			weight:        1,
			setIdentifier: "cluster/id",
			expected:      "cluster/id",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			policy, _ := NewAWSRoute53Policy(
				ti.weight,
				ti.setIdentifier,
			)
			assert.Equal(t, policy.SetIdentifier, ti.expected)
		})
	}
}

func TestAWSRoute53Policy_Weight(t *testing.T) {
	for _, ti := range []struct {
		title         string
		weight        int64
		setIdentifier string
		expected      int64
	}{
		{
			title:         "valid weight ID",
			weight:        1,
			setIdentifier: "cluster/id",
			expected:      1,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			policy, _ := NewAWSRoute53Policy(
				ti.weight,
				ti.setIdentifier,
			)
			assert.Equal(t, policy.Weight, ti.expected)
		})
	}
}
