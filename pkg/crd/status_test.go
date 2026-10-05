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

package crd

import (
	"context"
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source/types"
)

func newStatusWriter(t *testing.T, funcs interceptor.Funcs, objs ...client.Object) *StatusWriter {
	t.Helper()
	scheme := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(scheme))
	c := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(objs...).
		WithInterceptorFuncs(funcs).
		Build()
	return NewStatusWriter(NewCRDClients(c, c))
}

func dnsEndpointRef(name string) *events.ObjectReference {
	return events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "foo", name, "", types.CRD)
}

func TestStatusWriterReportStatus(t *testing.T) {
	mine := &apiv1alpha1.DNSEndpoint{Namespace: "foo", Name: "mine", Generation: 3}

	t.Run("logs only DNSEndpoint objects", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{}, mine)
		foreign := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "foo", "theirs", "", types.Ingress)

		w.ReportStatus(t.Context(), []plan.PlannedObject{
			{Ref: nil, Endpoints: 1},
			{Ref: dnsEndpointRef("mine"), Endpoints: 2, Changed: true},
			{Ref: foreign, Endpoints: 1},
		}, errors.New("provider down"))

		logtest.TestHelperLogContains("DNSEndpoint foo/mine: 2 endpoint(s) planned, apply error: provider down (generation=3, observedGeneration=0)", hook, t)
		logtest.TestHelperLogNotContains("foo/theirs", hook, t)
	})

	t.Run("an apply error skips objects outside the batch", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{}, mine)

		w.ReportStatus(t.Context(), []plan.PlannedObject{
			{Ref: dnsEndpointRef("mine"), Endpoints: 1},
		}, errors.New("provider down"))

		logtest.TestHelperLogContains("DNSEndpoint foo/mine: 1 endpoint(s) planned, apply error: <nil>", hook, t)
	})

	// Deleted between the plan and the apply.
	t.Run("skips missing objects silently", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{})

		w.ReportStatus(t.Context(), []plan.PlannedObject{{Ref: dnsEndpointRef("gone"), Endpoints: 1}}, nil)

		logtest.TestHelperLogNotContains("foo/gone", hook, t)
	})

	t.Run("warns on read errors", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{
			Get: func(context.Context, client.WithWatch, client.ObjectKey, client.Object, ...client.GetOption) error {
				return errors.New("apiserver unavailable")
			},
		}, mine)

		w.ReportStatus(t.Context(), []plan.PlannedObject{{Ref: dnsEndpointRef("mine"), Endpoints: 1}}, nil)

		logtest.TestHelperLogContainsWithLogLevel("Could not get DNSEndpoint foo/mine: apiserver unavailable", log.WarnLevel, hook, t)
	})
}
