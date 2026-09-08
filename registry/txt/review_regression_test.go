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

package txt

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

func TestTXTRegistryZoneAwareUnmatchedOwnership(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	require.NoError(t, p.CreateZone("example.com"))
	for i, owner := range []string{"one", "two"} {
		require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint(fmt.Sprintf("unrelated-%d.example.com", i), endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: owner}.Serialize(true, false, nil)),
		}}))
	}
	r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true}, zonedProvider{p})
	require.NoError(t, err)
	_, err = r.Records(t.Context())
	require.NoError(t, err)
	stored, err := p.Records(t.Context())
	require.NoError(t, err)
	assert.Len(t, stored, 2)
}

func TestTXTRegistryZoneAwareMultiValueOwnership(t *testing.T) {
	own := endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)
	for _, value := range []string{endpoint.Labels{endpoint.OwnerLabelKey: "other"}.Serialize(true, false, nil), `"verification=unowned"`, `"heritage=invalid"`, own} {
		for _, reverse := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/reverse=%t", value, reverse), func(t *testing.T) {
				p := inmemory.NewInMemoryProvider()
				require.NoError(t, p.CreateZone("example.com"))
				targets := []string{own, value}
				if reverse {
					targets[0], targets[1] = targets[1], targets[0]
				}
				seed := []*endpoint.Endpoint{
					endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1"),
					endpoint.NewEndpoint("a-name.sub-txtsuffix.example.com", endpoint.RecordTypeTXT, targets...),
				}
				require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: seed}))
				before, err := p.Records(t.Context())
				require.NoError(t, err)
				for _, record := range before {
					if record.RecordType == endpoint.RecordTypeTXT {
						require.Equal(t, endpoint.Targets(targets), record.Targets)
					}
				}
				r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA}}, zonedProvider{p})
				require.NoError(t, err)
				records, err := r.Records(t.Context())
				require.ErrorContains(t, err, "multiple targets")
				assert.Nil(t, records)
				after, err := p.Records(t.Context())
				require.NoError(t, err)
				assert.ElementsMatch(t, before, after)
			})
		}
	}
}

func TestTXTRegistryZoneAwareSharedUntypedOwnership(t *testing.T) {
	for _, otherType := range []string{endpoint.RecordTypeMX, endpoint.RecordTypeCNAME} {
		t.Run(otherType, func(t *testing.T) {
			p := inmemory.NewInMemoryProvider()
			require.NoError(t, p.CreateZone("example.com"))
			require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{
				endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1"),
				endpoint.NewEndpoint("name.sub.example.com", otherType, "mail.example.com"),
				endpoint.NewEndpoint("name-txtsuffix.sub.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)),
			}}))
			before, err := p.Records(t.Context())
			require.NoError(t, err)
			r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA, otherType}}, zonedProvider{p})
			require.NoError(t, err)
			records, err := r.Records(t.Context())
			require.ErrorContains(t, err, "shared TXT ownership")
			assert.Nil(t, records)
			after, err := p.Records(t.Context())
			require.NoError(t, err)
			assert.ElementsMatch(t, before, after)
		})
	}
}

func TestTXTRegistryZoneAwareFallbackOwnershipConflict(t *testing.T) {
	for _, alias := range []bool{false, true} {
		for _, fallback := range []string{"name-txtsuffix.sub.example.com", "cname-name-txtsuffix.sub.example.com"} {
			if !alias && strings.HasPrefix(fallback, "cname-") {
				continue
			}
			t.Run(fmt.Sprintf("alias=%t/%s", alias, fallback), func(t *testing.T) {
				p := inmemory.NewInMemoryProvider()
				require.NoError(t, p.CreateZone("example.com"))
				data := endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1")
				if alias {
					data.WithAliasProperty(endpoint.AliasTrue)
				}
				require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{
					data,
					endpoint.NewEndpoint("a-name.sub-txtsuffix.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)),
					endpoint.NewEndpoint(fallback, endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "other"}.Serialize(true, false, nil)),
				}}))
				r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true}, zonedProvider{p})
				require.NoError(t, err)
				records, err := r.Records(t.Context())
				require.ErrorContains(t, err, "conflicting TXT ownership")
				assert.Nil(t, records)
			})
		}
	}
}

func TestTXTRegistryZoneAwareNilStoredTXTLabels(t *testing.T) {
	p := &storedNilLabelsProvider{records: []*endpoint.Endpoint{
		endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1"),
		endpoint.NewEndpoint("a-name.sub-txtsuffix.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)),
	}}
	p.records[1].Labels = nil
	r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA}}, zonedProvider{p})
	require.NoError(t, err)
	records, err := r.Records(t.Context())
	require.NoError(t, err)
	require.NotPanics(t, func() { require.NoError(t, r.ApplyChanges(t.Context(), &plan.Changes{Delete: records})) })
	require.NotNil(t, p.changes)
	require.Len(t, p.changes.Delete, 2)
	assert.Equal(t, "name.sub.example.com", p.changes.Delete[1].Labels[endpoint.OwnedRecordLabelKey])
	assert.Nil(t, p.records[1].Labels)
}

type storedNilLabelsProvider struct {
	provider.Provider
	records []*endpoint.Endpoint
	changes *plan.Changes
}

func (p *storedNilLabelsProvider) Records(context.Context) ([]*endpoint.Endpoint, error) {
	return p.records, nil
}
func (p *storedNilLabelsProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	p.changes = changes
	return nil
}

func TestTXTRegistryZoneAwareIndependentTypedOwnership(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	require.NoError(t, p.CreateZone("example.com"))
	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{
		endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1"),
		endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeMX, "mail.example.com"),
		endpoint.NewEndpoint("a-name.sub-txtsuffix.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)),
		endpoint.NewEndpoint("mx-name.sub-txtsuffix.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "other"}.Serialize(true, false, nil)),
	}}))
	r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA, endpoint.RecordTypeMX}}, zonedProvider{p})
	require.NoError(t, err)
	records, err := r.Records(t.Context())
	require.NoError(t, err)
	require.Len(t, records, 2)
	require.NoError(t, r.ApplyChanges(t.Context(), &plan.Changes{Delete: records}))
	records, err = r.Records(t.Context())
	require.NoError(t, err)
	require.Len(t, records, 1)
	assert.Equal(t, endpoint.RecordTypeMX, records[0].RecordType)
	assert.Equal(t, "other", records[0].Labels[endpoint.OwnerLabelKey])
}

func TestTXTRegistryZoneAwareMatchingFallbackLifecycle(t *testing.T) {
	for _, alias := range []bool{false, true} {
		t.Run(fmt.Sprint(alias), func(t *testing.T) {
			p := inmemory.NewInMemoryProvider()
			require.NoError(t, p.CreateZone("example.com"))
			data := endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1")
			if alias {
				data.WithAliasProperty(endpoint.AliasTrue)
			}
			names := []string{"a-name.sub-txtsuffix.example.com", "name-txtsuffix.sub.example.com"}
			if alias {
				names = append(names, "cname-name-txtsuffix.sub.example.com")
			}
			seed := []*endpoint.Endpoint{data}
			for _, name := range names {
				seed = append(seed, endpoint.NewEndpoint(name, endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)))
			}
			require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: seed}))
			r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA}}, zonedProvider{p})
			require.NoError(t, err)
			records, err := r.Records(t.Context())
			require.NoError(t, err)
			require.Len(t, records, 1)
			updated := records[0].DeepCopy()
			updated.Labels[endpoint.ResourceLabelKey] = "service/default/updated"
			require.NoError(t, r.ApplyChanges(t.Context(), &plan.Changes{UpdateOld: records, UpdateNew: []*endpoint.Endpoint{updated}}))
			stored, err := p.Records(t.Context())
			require.NoError(t, err)
			require.Len(t, stored, len(seed))
			for _, ep := range stored {
				if ep.RecordType == endpoint.RecordTypeTXT {
					assert.Contains(t, names, ep.DNSName)
					assert.Contains(t, ep.Targets[0], "service/default/updated")
				}
			}
			records, err = r.Records(t.Context())
			require.NoError(t, err)
			require.NoError(t, r.ApplyChanges(t.Context(), &plan.Changes{Delete: records}))
			stored, err = p.Records(t.Context())
			require.NoError(t, err)
			assert.Empty(t, stored)
		})
	}
}

func TestTXTRegistryZoneAwareUnrelatedMultiValueTXT(t *testing.T) {
	for _, name := range []string{"unrelated.example.com", "a-name.sub-txtsuffix.example.com"} {
		t.Run(name, func(t *testing.T) {
			p := inmemory.NewInMemoryProvider()
			require.NoError(t, p.CreateZone("example.com"))
			values := []string{`"verification=one"`, `"verification=two"`}
			if name == "unrelated.example.com" {
				values[1] = endpoint.Labels{endpoint.OwnerLabelKey: "other"}.Serialize(true, false, nil)
			}
			require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{endpoint.NewEndpoint(name, endpoint.RecordTypeTXT, values...)}}))
			r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true}, zonedProvider{p})
			require.NoError(t, err)
			_, err = r.Records(t.Context())
			require.NoError(t, err)
		})
	}
}

func TestTXTRegistryZoneAwareSharedAliasOwnership(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	require.NoError(t, p.CreateZone("example.com"))
	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{
		endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "lb.example.net").WithAliasProperty(endpoint.AliasTrue),
		endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeCNAME, "lb.example.net"),
		endpoint.NewEndpoint("cname-name-txtsuffix.sub.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)),
	}}))
	r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true, ManagedDNSRecordTypes: []string{endpoint.RecordTypeA}}, zonedProvider{p})
	require.NoError(t, err)
	records, err := r.Records(t.Context())
	require.ErrorContains(t, err, "shared TXT ownership")
	assert.Nil(t, records)
}

func TestTXTRegistryZoneAwareSetIdentifierIsolation(t *testing.T) {
	p := &storedNilLabelsProvider{}
	for i, owner := range []string{"owner", "other"} {
		id := fmt.Sprint(i)
		p.records = append(p.records,
			endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1").WithSetIdentifier(id),
			endpoint.NewEndpoint("name-txtsuffix.sub.example.com", endpoint.RecordTypeTXT, endpoint.Labels{endpoint.OwnerLabelKey: owner}.Serialize(true, false, nil)).WithSetIdentifier(id),
		)
	}
	r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix", TXTZoneAware: true}, zonedProvider{p})
	require.NoError(t, err)
	records, err := r.Records(t.Context())
	require.NoError(t, err)
	require.Len(t, records, 2)
	assert.Equal(t, "owner", records[0].Labels[endpoint.OwnerLabelKey])
	assert.Equal(t, "other", records[1].Labels[endpoint.OwnerLabelKey])
}

func TestTXTRegistryDefaultOwnershipCompatibility(t *testing.T) {
	for _, multivalue := range []bool{false, true} {
		t.Run(fmt.Sprint(multivalue), func(t *testing.T) {
			own := endpoint.Labels{endpoint.OwnerLabelKey: "owner"}.Serialize(true, false, nil)
			foreign := endpoint.Labels{endpoint.OwnerLabelKey: "other"}.Serialize(true, false, nil)
			marker := endpoint.NewEndpoint("name-txtsuffix.sub.example.com", endpoint.RecordTypeTXT, own)
			if multivalue {
				marker.Targets = append(marker.Targets, foreign)
			}
			p := &storedNilLabelsProvider{records: []*endpoint.Endpoint{
				endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeA, "192.0.2.1"),
				endpoint.NewEndpoint("name.sub.example.com", endpoint.RecordTypeMX, "mail.example.com"), marker,
				endpoint.NewEndpoint("a-name-txtsuffix.sub.example.com", endpoint.RecordTypeTXT, foreign),
			}}
			r, err := New(&externaldns.Config{TXTOwnerID: "owner", TXTSuffix: "-txtsuffix"}, zonedProvider{p})
			require.NoError(t, err)
			records, err := r.Records(t.Context())
			require.NoError(t, err)
			require.Len(t, records, 2)
			assert.Equal(t, "other", records[0].Labels[endpoint.OwnerLabelKey])
			assert.Equal(t, "owner", records[1].Labels[endpoint.OwnerLabelKey])
		})
	}
}
