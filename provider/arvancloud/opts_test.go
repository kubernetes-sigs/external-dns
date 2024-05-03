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

package arvancloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArvanCloud_apply(t *testing.T) {
	tests := []struct {
		do    func() Option
		given *providerOptions
		wants *providerOptions
		name  string
	}{
		{
			name: "should successfully enable cloud proxy option",
			do: func() Option {
				return WithEnableCloudProxy()
			},
			given: &providerOptions{enableCloudProxy: false},
			wants: &providerOptions{enableCloudProxy: true},
		},
		{
			name: "should successfully disable cloud proxy option",
			do: func() Option {
				return WithDisableCloudProxy()
			},
			given: &providerOptions{enableCloudProxy: true},
			wants: &providerOptions{enableCloudProxy: false},
		},
		{
			name: "should successfully set domain per page",
			do: func() Option {
				return WithDomainPerPage(20)
			},
			given: &providerOptions{domainPerPage: 10},
			wants: &providerOptions{domainPerPage: 20},
		},
		{
			name: "should successfully set dns per page",
			do: func() Option {
				return WithDnsPerPage(200)
			},
			given: &providerOptions{dnsPerPage: 100},
			wants: &providerOptions{dnsPerPage: 200},
		},
		{
			name: "should successfully set api version",
			do: func() Option {
				return WithApiVersion("4.0")
			},
			given: &providerOptions{apiVersion: "3.0"},
			wants: &providerOptions{apiVersion: "4.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.do().apply(tt.given)

			assert.Equal(t, tt.wants, tt.given)
		})
	}
}
