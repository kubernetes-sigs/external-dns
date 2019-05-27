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
)

func TestParseIngressGatewayService(t *testing.T) {
	for _, tv := range []struct {
		name              string
		expectedNamespace string
		expectedName      string
		expectedErr       error
	}{
		{
			name:              "istio/ingress-gateway",
			expectedNamespace: "istio",
			expectedName:      "ingress-gateway",
			expectedErr:       nil,
		},
		{
			name:              "no-namespace",
			expectedNamespace: "",
			expectedName:      "",
			expectedErr:       fmt.Errorf("invalid ingress gateway service (namespace/name) found 'no-namespace'"),
		},
	} {
		namespace, name, err := parseIngressGateway(tv.name)

		assert.Equal(t, namespace, tv.expectedNamespace)
		assert.Equal(t, name, tv.expectedName)
		assert.Equal(t, err, tv.expectedErr)
	}
}
