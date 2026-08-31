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

package controller

import (
	"context"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/registry/txt"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

// canonicalizingProvider stores IPv6 targets in their canonical short form, which is
// what Route53, Cloudflare and Azure DNS do with an AAAA target. Everything else is
// delegated to the in-memory provider.
type canonicalizingProvider struct {
	*inmemory.InMemoryProvider
}

func (p *canonicalizingProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	for _, eps := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew, changes.UpdateOld, changes.Delete} {
		for _, ep := range eps {
			for i, t := range ep.Targets {
				if ip, err := netip.ParseAddr(t); err == nil {
					ep.Targets[i] = ip.String()
				}
			}
		}
	}
	return p.InMemoryProvider.ApplyChanges(ctx, changes)
}

func providerAAAATargets(t *testing.T, p provider.Provider, dnsName string) endpoint.Targets {
	t.Helper()
	records, err := p.Records(t.Context())
	require.NoError(t, err)
	for _, r := range records {
		if r.DNSName == dnsName && r.RecordType == endpoint.RecordTypeAAAA {
			return r.Targets
		}
	}
	return nil
}

func ipv6TargetService(target string) *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "app",
			Annotations: map[string]string{
				annotations.HostnameKey: "app.example.com",
				annotations.TargetKey:   target,
			},
		},
		Spec: v1.ServiceSpec{Type: v1.ServiceTypeLoadBalancer},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{Ingress: []v1.LoadBalancerIngress{{IP: "2001:db8::1"}}},
		},
	}
}

func ipv6ServiceSource(t *testing.T, ctx context.Context, kube *fake.Clientset) source.Source {
	t.Helper()
	src, err := source.NewServiceSource(ctx, kube, &source.Config{
		Namespace:      "default",
		LabelFilter:    labels.Everything(),
		TemplateEngine: templatetest.MustEngine(t, "", "", "", false),
	})
	require.NoError(t, err)
	return src
}

// A Service whose target annotation carries an expanded IPv6 alongside a second
// target. The provider stores the first one canonicalized, so on the next sync the
// two forms differ as strings while denoting the same address. Changing only the
// second target must still reach the provider.
func TestReconcileIPv6SecondTargetChangeReachesProvider(t *testing.T) {
	ctx := t.Context()
	kube := fake.NewClientset()

	const expanded = "2001:0db8:0000:0000:0000:0000:0000:0001"
	_, err := kube.CoreV1().Services("default").Create(ctx, ipv6TargetService(expanded+",2001:db8::2"), metav1.CreateOptions{})
	require.NoError(t, err)

	prov := &canonicalizingProvider{inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones([]string{"example.com"}))}
	registry, err := txt.New(&externaldns.Config{
		TXTOwnerID:            "test-cluster",
		ManagedDNSRecordTypes: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	}, prov)
	require.NoError(t, err)

	ctrl := &Controller{
		Source:             ipv6ServiceSource(t, ctx, kube),
		Registry:           registry,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	}

	require.NoError(t, ctrl.RunOnce(ctx))
	require.ElementsMatch(t, []string{"2001:db8::1", "2001:db8::2"}, providerAAAATargets(t, prov, "app.example.com"),
		"first sync should create the record")

	// The user edits the annotation and changes only the second target.
	_, err = kube.CoreV1().Services("default").Update(ctx, ipv6TargetService(expanded+",2001:db8::3"), metav1.UpdateOptions{})
	require.NoError(t, err)
	ctrl.Source = ipv6ServiceSource(t, ctx, kube)

	require.NoError(t, ctrl.RunOnce(ctx))
	assert.ElementsMatch(t, []string{"2001:db8::1", "2001:db8::3"}, providerAAAATargets(t, prov, "app.example.com"),
		"the change to the second target never reached the provider")
}
