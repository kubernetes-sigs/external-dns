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

package source

import (
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/types"
)

var _ StatusReporter = &crdSource{}

// The crd source ignores objects from other sources.
func TestCRDSourceReportStatusLogsOnlyItsOwnObjects(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	fakeCache := newFakeCRDCache(t, nil)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
	require.NoError(t, err)

	own := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "foo", "mine", "", types.CRD)
	foreign := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "foo", "theirs", "", types.Ingress)

	cs.ReportStatus(t.Context(), []PlannedObject{
		{Ref: nil, Endpoints: 1},
		{Ref: own, Endpoints: 2},
		{Ref: foreign, Endpoints: 1},
	}, errors.New("provider down"))

	logtest.TestHelperLogContains("dnsendpoint foo/mine: 2 endpoint(s) planned, apply error: provider down", hook, t)
	logtest.TestHelperLogNotContains("foo/theirs", hook, t)
}
