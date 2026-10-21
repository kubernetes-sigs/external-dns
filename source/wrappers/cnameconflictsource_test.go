/*
Copyright 2026 The Kubernetes Authors.

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

package wrappers

import (
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

// The log hook and the metric are process-wide, so this test must not run in parallel.
func TestCNAMEConflictSourceEndpoints(t *testing.T) {
	conflict := func(name, recordType string) []*endpoint.Endpoint {
		return []*endpoint.Endpoint{
			endpoint.NewEndpoint(name, recordType, "a.elb.example.com"),
			endpoint.NewEndpoint(name, recordType, "b.elb.example.com"),
		}
	}
	filter := endpoint.NewDomainFilterWithExclusions([]string{"example.org"}, []string{"internal.example.org"})

	tests := []struct {
		title     string
		filter    endpoint.DomainFilterInterface
		endpoints []*endpoint.Endpoint
		want      float64 // conflicts warned and counted
	}{
		{"conflict inside the domain filter", filter, conflict("app.example.org", endpoint.RecordTypeCNAME), 1},
		{"DNAME conflict inside the domain filter", filter, conflict("app.example.org", endpoint.RecordTypeDNAME), 1},
		{"conflict outside the domain filter", filter, conflict("app.example.com", endpoint.RecordTypeCNAME), 0},
		{"conflict on an excluded name", filter, conflict("app.internal.example.org", endpoint.RecordTypeCNAME), 0},
		{"conflict without a domain filter", nil, conflict("app.example.com", endpoint.RecordTypeCNAME), 1},
		{"same target twice", filter, []*endpoint.Endpoint{
			endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeCNAME, "a.elb.example.com"),
			endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeCNAME, "a.elb.example.com"),
		}, 0},
		{"different set identifiers", filter, []*endpoint.Endpoint{
			endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeCNAME, "a.elb.example.com").WithSetIdentifier("one"),
			endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeCNAME, "b.elb.example.com").WithSetIdentifier("two"),
		}, 0},
		{"conflict between names that differ in case", filter, []*endpoint.Endpoint{
			endpoint.NewEndpoint("App.example.org", endpoint.RecordTypeCNAME, "a.elb.example.com"),
			endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeCNAME, "b.elb.example.com"),
		}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			invalidEndpoints.Reset()
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)
			want := make([]*endpoint.Endpoint, len(tt.endpoints))
			for i, ep := range tt.endpoints {
				want[i] = ep.DeepCopy()
			}
			recordType := tt.endpoints[0].RecordType
			msg := "Only one " + recordType + " per name"

			endpoints, err := NewCNAMEConflictSource(testutils.NewMockSource(tt.endpoints...), tt.filter).Endpoints(t.Context())
			require.NoError(t, err)

			assert.Equal(t, want, endpoints, "endpoints must be returned unchanged")
			testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, tt.want, invalidEndpoints.Gauge, map[string]string{"record_type": strings.ToLower(recordType)})
			if tt.want > 0 {
				logtest.TestHelperLogContainsWithLogLevel(msg, log.WarnLevel, hook, t)
			} else {
				logtest.TestHelperLogNotContains(msg, hook, t)
			}
		})
	}
}
