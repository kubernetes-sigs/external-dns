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

import "strings"

// endpointTargetCleaner is essentially a function that cleans a target string for a given record.
// It returns the cleaned target string, which is used to create an endpoint. Different record types have different cleaning rules.
type endpointTargetCleaner interface {
	Clean(target string) string
}

// defaultEndpointTargetCleaner is a function that cleans a target string for a given record, trimming trailing dots which are not allowed
// for most record types.
type defaultEndpointTargetCleaner struct{}

func (v *defaultEndpointTargetCleaner) Clean(target string) string {
	return strings.TrimSuffix(target, ".")
}

// arbitraryTextEndpointTargetCleaner is a function that cleans a target string for a given record, preserving arbitrary text including trailing dots.
// This is used for TXT and NAPTR records.
type arbitraryTextEndpointTargetCleaner struct{}

func (v *arbitraryTextEndpointTargetCleaner) Clean(target string) string {
	return target
}
