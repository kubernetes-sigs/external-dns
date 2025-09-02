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

package azure

import (
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	dns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	privatedns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"

	"github.com/stretchr/testify/assert"
)

func Test_parseMxTarget(t *testing.T) {
	type testCase[T interface {
		dns.MxRecord | privatedns.MxRecord
	}] struct {
		name    string
		args    string
		want    T
		wantErr assert.ErrorAssertionFunc
	}

	tests := []testCase[dns.MxRecord]{
		{
			name: "valid mx target",
			args: "10 example.com",
			want: dns.MxRecord{
				Preference: to.Ptr(int32(10)),
				Exchange:   to.Ptr("example.com"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "valid mx target with a subdomain",
			args: "99 foo-bar.example.com",
			want: dns.MxRecord{
				Preference: to.Ptr(int32(99)),
				Exchange:   to.Ptr("foo-bar.example.com"),
			},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid mx target with misplaced preference and exchange",
			args:    "example.com 10",
			want:    dns.MxRecord{},
			wantErr: assert.Error,
		},
		{
			name:    "invalid mx target without preference",
			args:    "example.com",
			want:    dns.MxRecord{},
			wantErr: assert.Error,
		},
		{
			name:    "invalid mx target with non numeric preference",
			args:    "aa example.com",
			want:    dns.MxRecord{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMxTarget[dns.MxRecord](tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("parseMxTarget(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseMxTarget(%v)", tt.args)
		})
	}
}
