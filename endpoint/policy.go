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

// Policy is a generic policy that stores provider-specific policies
type Policy struct {
	// AWSRoute53 is the AWS Route53 policy for the endpoint
	AWSRoute53 *AWSRoute53Policy
}

// AttachAWSRoute53Policy attaches a Route53 policy to an endpoint
func (e *Policy) AttachAWSRoute53Policy(awsRoute53Policy *AWSRoute53Policy) {
	e.AWSRoute53 = awsRoute53Policy
}

// HasAWSRoute53Policy checks whether an endpoint has a Route53 policy attached
func (e *Policy) HasAWSRoute53Policy() bool {
	// Return false when no attributes of the AWS Route 53 policy have been set
	if e.AWSRoute53 != nil {
		if e.AWSRoute53.Weight >= 0 && e.AWSRoute53.SetIdentifier != "" {
			return true
		}
	}
	return false
}
