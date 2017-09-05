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
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestEndpointPolicy_AttachAWSRoute53Policy(t *testing.T) {
	p := Policy{}
	w, _ := NewAWSRoute53Policy(1, "cluster/id")
	p.AttachAWSRoute53Policy(w)
	assert.True(t, p.HasAWSRoute53Policy())
}

func TestHasAWSRoute53Policy(t *testing.T) {
	p := Policy{}
	assert.False(t, p.HasAWSRoute53Policy())
}
