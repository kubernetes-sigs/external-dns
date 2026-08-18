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

package aws

// These tests drive the full reconcile loop, exactly as controller.RunOnce does:
//
//	registry.Records() -> registry.AdjustEndpoints() -> plan.Calculate() -> registry.ApplyChanges()
//
// against the real TXT registry and the real AWS provider, backed by Route53APIStub
// (which models Route53 semantics: CREATE on an existing rrset fails, DELETE of a
// missing rrset fails, UPSERT replaces).
//
// Single-loop provider tests cannot observe two properties that only appear over
// repeated reconciliations, and both are reported symptoms of a load balancer swap:
// whether a steady desired state settles instead of rewriting records forever, and
// whether the previous load balancer address is really gone once it does settle.
//
// provider/factory wraps every provider in factory.AliasNormalizingMiddleware, which is
// not applied here. It is a no-op for these cases: it only rewrites alias=A/AAAA, and
// adjustCNAMERecordAndNewAaaaIfNeeded only produces those from an explicit
// external-dns.alpha.kubernetes.io/aws-alias annotation. A CNAME published by a source
// always yields alias=true.

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/registry/txt"
)

const (
	reconcileZone   = "zone-1.ext-dns-test-2.teapot.zalan.do"
	reconcileZoneID = "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."
	reconcileOwner  = "reconcile-owner"

	// Two load balancers in one region, plus one in another region so that an
	// AliasTarget.HostedZoneId change is covered as well.
	lbOld         = "old-1111.eu-central-1.elb.amazonaws.com"
	lbNew         = "new-2222.eu-central-1.elb.amazonaws.com"
	lbNewOtherReg = "new-3333.us-east-1.elb.amazonaws.com"

	// Targets outside AWS, which stay CNAMEs instead of becoming A ALIAS records.
	lbExternal      = "external.example.net"
	lbExternalOther = "other.example.org"

	// maxReconcileRounds bounds how long a steady desired state may keep producing
	// changes before it is treated as never settling.
	maxReconcileRounds = 8

	// settledRounds is the number of loops a correct reconcile needs: one to apply
	// the change, one to observe that nothing is left to do.
	settledRounds = 2
)

// reconcileLoop runs repeated reconciliations over one provider and registry pair.
type reconcileLoop struct {
	t      *testing.T
	prov   *AWSProvider
	client *Route53APIStub
	reg    registry.Registry
	policy plan.Policy
	cfg    *externaldns.Config
}

func reconcileConfig(prefix, suffix string) *externaldns.Config {
	return &externaldns.Config{
		TXTPrefix:             prefix,
		TXTSuffix:             suffix,
		TXTOwnerID:            reconcileOwner,
		ManagedDNSRecordTypes: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
		ExcludeDNSRecordTypes: []string{},
	}
}

func newReconcileLoop(t *testing.T, cfg *externaldns.Config, policyName string, preferCNAME bool) *reconcileLoop {
	t.Helper()

	prov, client := newAWSProvider(t,
		endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}),
		provider.NewZoneIDFilter([]string{}),
		provider.NewZoneTypeFilter(""),
		defaultEvaluateTargetHealth, preferCNAME, false, nil)
	// Every reconciliation submits a change batch; the production pause between
	// batches would dominate the runtime of these multi-loop tests.
	prov.batchChangeInterval = 0
	// A loop must not be able to settle on a zone Route 53 would refuse.
	client.rejectCNAMEConflicts = true

	reg, err := txt.New(cfg, prov)
	require.NoError(t, err)

	policy, ok := plan.Policies[policyName]
	require.Truef(t, ok, "unknown policy %q", policyName)

	return &reconcileLoop{t: t, prov: prov, client: client, reg: reg, policy: policy, cfg: cfg}
}

// runOnce replays controller.Controller.RunOnce for one desired state.
func (h *reconcileLoop) runOnce(desired []*endpoint.Endpoint) (*plan.Changes, error) {
	ctx := h.t.Context()

	current, err := h.reg.Records(ctx)
	if err != nil {
		return nil, fmt.Errorf("registry.Records: %w", err)
	}
	ctx = context.WithValue(ctx, provider.RecordsContextKey, current)

	// A source rebuilds its endpoints every loop, and AdjustEndpoints mutates them.
	source := make([]*endpoint.Endpoint, 0, len(desired))
	for _, ep := range desired {
		source = append(source, ep.DeepCopy())
	}

	adjusted, err := h.reg.AdjustEndpoints(source)
	if err != nil {
		return nil, fmt.Errorf("registry.AdjustEndpoints: %w", err)
	}

	p := &plan.Plan{
		Policies:       []plan.Policy{h.policy},
		Current:        current,
		Desired:        adjusted,
		DomainFilter:   endpoint.MatchAllDomainFilters{h.reg.GetDomainFilter()},
		ManagedRecords: h.cfg.ManagedDNSRecordTypes,
		ExcludeRecords: h.cfg.ExcludeDNSRecordTypes,
		OwnerID:        h.reg.OwnerID(),
		OldOwnerID:     h.cfg.TXTOwnerOld,
	}
	p = p.Calculate()

	if p.Changes.HasChanges() {
		if err := h.reg.ApplyChanges(ctx, p.Changes); err != nil {
			return p.Changes, fmt.Errorf("registry.ApplyChanges: %w", err)
		}
	}
	return p.Changes, nil
}

// settle reconciles until a loop reports no changes, which is what makes the
// controller log "All records are already up to date", and returns the loop count.
func (h *reconcileLoop) settle(desired []*endpoint.Endpoint) (int, error) {
	for i := 1; i <= maxReconcileRounds; i++ {
		changes, err := h.runOnce(desired)
		if err != nil {
			return i, fmt.Errorf("loop %d: %w", i, err)
		}
		if !changes.HasChanges() {
			return i, nil
		}
	}
	return maxReconcileRounds, fmt.Errorf("desired state still producing changes after %d loops", maxReconcileRounds)
}

// mustSettle asserts the desired state settles within settledRounds loops.
func (h *reconcileLoop) mustSettle(desired []*endpoint.Endpoint) {
	h.t.Helper()
	rounds, err := h.settle(desired)
	require.NoErrorf(h.t, err, "desired state:\n  %s\nzone:\n  %s", describeEndpoints(desired), h.zone())
	require.LessOrEqualf(h.t, rounds, settledRounds,
		"took %d loops to settle; desired state:\n  %s\nzone:\n  %s", rounds, describeEndpoints(desired), h.zone())
}

// zone renders every record the provider holds, as external-dns reads them back.
func (h *reconcileLoop) zone() string {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	lines := make([]string, 0, len(records))
	for _, r := range records {
		alias := ""
		if a := r.GetAliasProperty(); a != endpoint.AliasNone && a != endpoint.AliasFalse {
			alias = " alias=" + string(a)
		}
		lines = append(lines, fmt.Sprintf("%s %s -> %s%s", r.DNSName, r.RecordType, strings.Join(r.Targets, ","), alias))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n  ")
}

// targets returns every non-TXT target published for a hostname, across record types,
// so that assertions compare whole targets instead of matching substrings.
func (h *reconcileLoop) targets(dnsName string) []string {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	var out []string
	for _, r := range records {
		if r.DNSName != dnsName || r.RecordType == endpoint.RecordTypeTXT {
			continue
		}
		out = append(out, r.Targets...)
	}
	sort.Strings(out)
	return out
}

// canonicalRecord renders the parts of a record that the desired state decides. TTL is
// left out because the provider fills in its own default for alias records.
func canonicalRecord(ep *endpoint.Endpoint) string {
	targets := append([]string(nil), ep.Targets...)
	sort.Strings(targets)
	alias := ""
	if a := ep.GetAliasProperty(); a != endpoint.AliasNone && a != endpoint.AliasFalse {
		alias = " alias=" + string(a)
	}
	return fmt.Sprintf("%s %s %q -> %s%s",
		ep.DNSName, ep.RecordType, ep.SetIdentifier, strings.Join(targets, ","), alias)
}

// canonicalZone renders every non-TXT record the provider holds.
func (h *reconcileLoop) canonicalZone() []string {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	out := make([]string, 0, len(records))
	for _, r := range records {
		if r.RecordType == endpoint.RecordTypeTXT {
			continue
		}
		out = append(out, canonicalRecord(r))
	}
	sort.Strings(out)
	return out
}

// canonicalDesired renders what a desired state should leave behind, by running it
// through the same AdjustEndpoints the loop uses.
func (h *reconcileLoop) canonicalDesired(desired []*endpoint.Endpoint) []string {
	h.t.Helper()
	source := make([]*endpoint.Endpoint, 0, len(desired))
	for _, ep := range desired {
		source = append(source, ep.DeepCopy())
	}
	adjusted, err := h.reg.AdjustEndpoints(source)
	require.NoError(h.t, err)

	out := make([]string, 0, len(adjusted))
	for _, r := range adjusted {
		out = append(out, canonicalRecord(r))
	}
	sort.Strings(out)
	return out
}

// aliasHostedZone reads the AliasTarget hosted zone of a record straight from Route 53.
// Records() does not carry that field, so a stale hosted zone would otherwise go unseen.
func (h *reconcileLoop) aliasHostedZone(dnsName, recordType string) string {
	h.t.Helper()
	for _, rrs := range listAWSRecords(h.t, h.client, reconcileZoneID) {
		if rrs.Name == nil || strings.TrimSuffix(*rrs.Name, ".") != dnsName {
			continue
		}
		if string(rrs.Type) != recordType || rrs.AliasTarget == nil {
			continue
		}
		return aws.ToString(rrs.AliasTarget.HostedZoneId)
	}
	return ""
}

// unownedRecords returns the record types published for a hostname that the registry no
// longer attributes to this owner. A record can carry the right target while its
// ownership TXT is missing, and the loop still reports that everything is up to date.
func (h *reconcileLoop) unownedRecords(dnsName string) []string {
	h.t.Helper()
	records, err := h.reg.Records(h.t.Context())
	require.NoError(h.t, err)

	var out []string
	for _, r := range records {
		if r.DNSName != dnsName || r.RecordType == endpoint.RecordTypeTXT {
			continue
		}
		if r.Labels[endpoint.OwnerLabelKey] != reconcileOwner {
			out = append(out, r.RecordType)
		}
	}
	sort.Strings(out)
	return out
}

// findRecord returns the provider's view of one record, or nil when it is absent.
func (h *reconcileLoop) findRecord(dnsName, recordType string) *endpoint.Endpoint {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)
	for _, r := range records {
		if r.DNSName == dnsName && r.RecordType == recordType {
			return r
		}
	}
	return nil
}

// ownsRecord reports whether the registry still attributes a record to this owner.
func (h *reconcileLoop) ownsRecord(dnsName, recordType string) bool {
	h.t.Helper()
	records, err := h.reg.Records(h.t.Context())
	require.NoError(h.t, err)
	for _, r := range records {
		if r.DNSName == dnsName && r.RecordType == recordType {
			return r.Labels[endpoint.OwnerLabelKey] == reconcileOwner
		}
	}
	return false
}

func cnameTo(host, target string) *endpoint.Endpoint {
	return endpoint.NewEndpoint(host+"."+reconcileZone, endpoint.RecordTypeCNAME, target)
}

func describeEndpoints(endpoints []*endpoint.Endpoint) string {
	lines := make([]string, 0, len(endpoints))
	for _, e := range endpoints {
		lines = append(lines, fmt.Sprintf("%s %s -> %s", e.DNSName, e.RecordType, strings.Join(e.Targets, ",")))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n  ")
}

// txtAffixes covers every shape --txt-prefix / --txt-suffix can take, because the
// ownership TXT name is what decides whether a record is still recognised as ours.
var txtAffixes = []struct {
	name           string
	prefix, suffix string
}{
	{"no-affix", "", ""},
	{"prefix", "extdns-", ""},
	{"suffix", "", "-extdns"},
	{"record-type-prefix", "%{record_type}-extdns-", ""},
	{"record-type-suffix", "", "-extdns-%{record_type}"},
}

// TestReconcileLoopReplacesLoadBalancerTargetOfSameRecordType covers a load balancer
// being replaced while the record type stays the same, which is what happens when a
// Service or Ingress is given a new ELB. The old address must be gone once the loop
// settles, under every txt-prefix shape and under both policies that allow updates.
func TestReconcileLoopReplacesLoadBalancerTargetOfSameRecordType(t *testing.T) {
	transitions := []struct {
		name        string
		from, to    string
		preferCNAME bool
	}{
		{name: "elb-to-elb-same-region", from: lbOld, to: lbNew},
		{name: "elb-to-elb-other-region", from: lbOld, to: lbNewOtherReg},
		{name: "external-to-external", from: lbExternal, to: lbExternalOther},
		{name: "elb-to-elb-with-prefer-cname", from: lbOld, to: lbNew, preferCNAME: true},
	}

	for _, affix := range txtAffixes {
		for _, policy := range []string{"sync", "upsert-only"} {
			for _, tr := range transitions {
				t.Run(fmt.Sprintf("%s/%s/%s", affix.name, policy, tr.name), func(t *testing.T) {
					t.Parallel()
					h := newReconcileLoop(t, reconcileConfig(affix.prefix, affix.suffix), policy, tr.preferCNAME)
					host := "app." + reconcileZone

					h.mustSettle([]*endpoint.Endpoint{cnameTo("app", tr.from)})
					h.mustSettle([]*endpoint.Endpoint{cnameTo("app", tr.to)})

					require.NotContainsf(t, h.targets(host), tr.from,
						"previous load balancer %q still resolvable after the loop settled; zone:\n  %s", tr.from, h.zone())
					require.Containsf(t, h.targets(host), tr.to,
						"new load balancer %q missing; zone:\n  %s", tr.to, h.zone())
					require.Emptyf(t, h.unownedRecords(host),
						"swapping the load balancer lost ownership of %v; zone:\n  %s",
						h.unownedRecords(host), h.zone())

					// An alias record also carries the hosted zone of the target load
					// balancer, which changes when the load balancer moves region.
					if hostedZone := canonicalHostedZone(tr.to); hostedZone != "" && !tr.preferCNAME {
						for _, recordType := range []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA} {
							require.Equalf(t, hostedZone, h.aliasHostedZone(host, recordType),
								"%s alias still points at the previous hosted zone; zone:\n  %s", recordType, h.zone())
						}
					}
				})
			}
		}
	}
}

// TestReconcileLoopRecordTypeTransition covers a hostname moving between an AWS load
// balancer, which becomes an A ALIAS, and a non-AWS target, which stays a CNAME.
//
// Route 53 refuses a CNAME that shares a name with any other type, so the two policies
// end very differently. sync deletes the old type first and settles. upsert-only drops
// the delete half, so the create is rejected and the transition never completes.
func TestReconcileLoopRecordTypeTransition(t *testing.T) {
	transitions := []struct {
		name     string
		from, to string
	}{
		{name: "elb-alias-to-external-cname", from: lbOld, to: lbExternal},
		{name: "external-cname-to-elb-alias", from: lbExternal, to: lbNew},
	}

	for _, tr := range transitions {
		t.Run("sync/"+tr.name, func(t *testing.T) {
			t.Parallel()
			h := newReconcileLoop(t, reconcileConfig("", ""), "sync", false)
			host := "app." + reconcileZone

			h.mustSettle([]*endpoint.Endpoint{cnameTo("app", tr.from)})
			desired := []*endpoint.Endpoint{cnameTo("app", tr.to)}
			h.mustSettle(desired)

			require.Equalf(t, h.canonicalDesired(desired), h.canonicalZone(),
				"sync settled on a different zone than the desired state; zone:\n  %s", h.zone())
			require.NotContainsf(t, h.targets(host), tr.from,
				"the previous record type is still present; zone:\n  %s", h.zone())
			require.Emptyf(t, h.unownedRecords(host), "records left without ownership: %v; zone:\n  %s",
				h.unownedRecords(host), h.zone())
		})

		t.Run("upsert-only/"+tr.name, func(t *testing.T) {
			t.Parallel()
			h := newReconcileLoop(t, reconcileConfig("", ""), "upsert-only", false)

			h.mustSettle([]*endpoint.Endpoint{cnameTo("app", tr.from)})

			_, err := h.settle([]*endpoint.Endpoint{cnameTo("app", tr.to)})
			require.Errorf(t, err,
				"upsert-only cannot complete a record type change, Route 53 rejects the batch; zone:\n  %s",
				h.zone())
		})
	}
}

// TestReconcileLoopMigratesLegacyAliasOwnershipTXT starts from a zone written by a
// release that labelled A ALIAS records with a "cname-" ownership TXT, and checks that
// upgrading keeps ownership, settles, and can still swap the load balancer afterwards.
func TestReconcileLoopMigratesLegacyAliasOwnershipTXT(t *testing.T) {
	for _, policy := range []string{"sync", "upsert-only"} {
		for _, affix := range []struct{ name, prefix string }{{"no-affix", ""}, {"prefix", "extdns-"}} {
			t.Run(affix.name+"/"+policy, func(t *testing.T) {
				t.Parallel()
				h := newReconcileLoop(t, reconcileConfig(affix.prefix, ""), policy, false)

				// Rather than hand-encoding the legacy TXT payload, let the registry
				// write a real "cname-" ownership TXT by publishing a plain CNAME, then
				// replace the CNAME with the A ALIAS the older provider stored.
				h.mustSettle([]*endpoint.Endpoint{cnameTo("app", lbExternal)})

				plainCNAME := h.findRecord("app."+reconcileZone, endpoint.RecordTypeCNAME)
				require.NotNil(t, plainCNAME, "setup: expected a plain CNAME")
				require.NotNil(t, h.findRecord(affix.prefix+"cname-app."+reconcileZone, endpoint.RecordTypeTXT),
					"setup: expected a legacy cname- ownership TXT")

				aliasA := endpoint.NewEndpoint("app."+reconcileZone, endpoint.RecordTypeA, lbOld).
					WithProviderSpecific(endpoint.ProviderSpecificAlias, "true").
					WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false")
				require.NoError(t, h.prov.ApplyChanges(t.Context(), &plan.Changes{
					Delete: []*endpoint.Endpoint{plainCNAME},
					Create: []*endpoint.Endpoint{aliasA},
				}))

				desired := []*endpoint.Endpoint{cnameTo("app", lbOld)}

				first, err := h.runOnce(desired)
				require.NoError(t, err)
				require.Emptyf(t, first.Delete, "upgrading deleted live records: %v", first.Delete)

				h.mustSettle(desired)
				require.Truef(t, h.ownsRecord("app."+reconcileZone, endpoint.RecordTypeA),
					"ownership of the A ALIAS lost while migrating cname- to a-; zone:\n  %s", h.zone())

				h.mustSettle([]*endpoint.Endpoint{cnameTo("app", lbNew)})
				require.NotContainsf(t, h.targets("app."+reconcileZone), lbOld,
					"previous load balancer still resolvable after migration; zone:\n  %s", h.zone())
			})
		}
	}
}

// TestReconcileLoopSettlesForRandomisedDesiredStates walks randomised sequences of
// desired states over a few hostnames, mixing alias and non-alias targets, record
// types and hostnames that stop being published. Every step must settle; a step that
// keeps producing changes is the repeated-UPSERT behaviour reported against this issue.
func TestReconcileLoopSettlesForRandomisedDesiredStates(t *testing.T) {
	// Load balancer targets all become A ALIAS records, so swapping between them is
	// always a same-type replacement.
	aliasShapes := []func(host string) *endpoint.Endpoint{
		func(host string) *endpoint.Endpoint { return cnameTo(host, lbOld) },
		func(host string) *endpoint.Endpoint { return cnameTo(host, lbNew) },
		func(host string) *endpoint.Endpoint { return cnameTo(host, lbNewOtherReg) },
	}
	shapes := append([]func(host string) *endpoint.Endpoint{}, aliasShapes...)
	shapes = append(shapes,
		func(host string) *endpoint.Endpoint { return cnameTo(host, lbExternal) },
		func(host string) *endpoint.Endpoint {
			return endpoint.NewEndpoint(host+"."+reconcileZone, endpoint.RecordTypeA, "1.2.3.4")
		},
		func(host string) *endpoint.Endpoint {
			return endpoint.NewEndpoint(host+"."+reconcileZone, endpoint.RecordTypeA, "5.6.7.8")
		},
		func(host string) *endpoint.Endpoint {
			return endpoint.NewEndpoint(host+"."+reconcileZone, endpoint.RecordTypeAAAA, "2001:db8::1")
		},
	)
	hosts := []string{"app", "api", "web"}

	// Every affix shape is covered by the load balancer swap test above. Randomised
	// sequences only need the two the name mapper treats differently: the record type
	// in a label of its own, and the record type inside the affix.
	affixes := []struct {
		name           string
		prefix, suffix string
	}{
		{"no-affix", "", ""},
		{"record-type-prefix", "%{record_type}-extdns-", ""},
	}

	for _, policy := range []string{"sync", "upsert-only"} {
		// upsert-only drops the delete half of a record type transition, and Route 53
		// rejects a CNAME that shares a name with another type. Keep those sequences
		// inside the alias shapes so every step stays a same-type replacement.
		pool := shapes
		if policy != "sync" {
			pool = aliasShapes
		}
		for _, affix := range affixes {
			for seed := range uint64(8) {
				t.Run(fmt.Sprintf("%s/%s/seed-%d", affix.name, policy, seed), func(t *testing.T) {
					t.Parallel()
					rng := rand.New(rand.NewPCG(seed, seed+1))
					h := newReconcileLoop(t, reconcileConfig(affix.prefix, affix.suffix), policy, false)

					for range 10 {
						desired := make([]*endpoint.Endpoint, 0, len(hosts))
						published := make(map[string]bool, len(hosts))
						for _, host := range hosts {
							if rng.IntN(4) == 0 {
								continue // the source stops publishing this hostname
							}
							ep := pool[rng.IntN(len(pool))](host)
							published[ep.DNSName] = true
							desired = append(desired, ep)
						}

						h.mustSettle(desired)

						for _, ep := range desired {
							for _, target := range ep.Targets {
								require.Containsf(t, h.targets(ep.DNSName), target,
									"published target %q missing after settling; desired:\n  %s\nzone:\n  %s",
									target, describeEndpoints(desired), h.zone())
							}
						}

						for dnsName := range published {
							require.Emptyf(t, h.unownedRecords(dnsName),
								"records published for %q without ownership: %v; zone:\n  %s",
								dnsName, h.unownedRecords(dnsName), h.zone())
						}

						if policy != "sync" {
							continue // upsert-only deliberately keeps records the source dropped
						}
						require.Equalf(t, h.canonicalDesired(desired), h.canonicalZone(),
							"sync settled on a different zone than the desired state; desired:\n  %s",
							describeEndpoints(desired))
					}
				})
			}
		}
	}
}
