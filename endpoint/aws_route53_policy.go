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
	"fmt"
)

// AWSRoute53Policy stores the policy attributes for the Route53 record
type AWSRoute53Policy struct {
	// Weight is the weight of the RecordSet
	Weight int64
	// SetIdentifier for the RecordSet
	SetIdentifier string
}

// NewAWSRoute53Policy does basic validation according to AWS requirements and returns the Route53 policy
func NewAWSRoute53Policy(weight int64, setIdentifier string) (*AWSRoute53Policy, error) {
	if weight < 0 || weight > 255 {
		return &AWSRoute53Policy{}, fmt.Errorf("Weight must be between 0-255. Actual: %d", weight)
	}

	if len(setIdentifier) < 1 || len(setIdentifier) > 128 {
		return &AWSRoute53Policy{}, fmt.Errorf("Set Identifier must be between 1-128 characters. Actual: %s",
			setIdentifier)
	}

	return &AWSRoute53Policy{
		Weight:        weight,
		SetIdentifier: setIdentifier,
	}, nil
}
