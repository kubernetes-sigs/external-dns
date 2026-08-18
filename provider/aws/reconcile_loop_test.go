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

// These tests drive the reconciliation core the controller runs once a source has
// produced its endpoints:
//
//	registry.Records() -> registry.AdjustEndpoints() -> plan.Calculate() -> registry.ApplyChanges()
//
// against the real TXT registry and the real AWS provider, backed by Route53APIStub. The
// source itself, the source wrappers, the controller domain filter, the provider factory
// middleware and the provider cache are all outside this file.
//
// The stub applies a change batch all or nothing and, for these tests, refuses a CNAME
// that shares a name with another type, which is what keeps the loop from settling on a
// zone Route 53 would reject. It still accepts a DELETE that does not carry the whole
// existing record set.
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
	"strconv"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/registry/mapper"
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

// recordState renders the parts of a record that the desired state decides. TTL is
// left out because the provider fills in its own default for alias records.
func recordState(ep *endpoint.Endpoint) string {
	targets := append([]string(nil), ep.Targets...)
	sort.Strings(targets)
	alias := ""
	if a := ep.GetAliasProperty(); a != endpoint.AliasNone && a != endpoint.AliasFalse {
		alias = " alias=" + string(a)
	}
	return fmt.Sprintf("%s %s %q -> %s%s",
		ep.DNSName, ep.RecordType, ep.SetIdentifier, strings.Join(targets, ","), alias)
}

// settledState renders every non-TXT record the provider holds.
func (h *reconcileLoop) settledState() []string {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	out := make([]string, 0, len(records))
	for _, r := range records {
		if r.RecordType == endpoint.RecordTypeTXT {
			continue
		}
		out = append(out, recordState(r))
	}
	sort.Strings(out)
	return out
}

// stateFor renders the non-TXT records of the given hostnames only, so a policy which
// keeps what the source dropped can still be held to an exact state for what it publishes.
func (h *reconcileLoop) stateFor(names map[string]bool) []string {
	h.t.Helper()
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	out := make([]string, 0, len(records))
	for _, r := range records {
		if r.RecordType == endpoint.RecordTypeTXT || !names[r.DNSName] {
			continue
		}
		out = append(out, recordState(r))
	}
	sort.Strings(out)
	return out
}

// ownershipTXTNames returns the TXT records which map back to a hostname. A primary
// record can be removed while its ownership TXT stays behind.
func (h *reconcileLoop) ownershipTXTNames(dnsName string) []string {
	h.t.Helper()
	names := mapper.NewAffixNameMapper(h.cfg.TXTPrefix, h.cfg.TXTSuffix, "")
	records, err := h.prov.Records(h.t.Context())
	require.NoError(h.t, err)

	var out []string
	for _, r := range records {
		if r.RecordType != endpoint.RecordTypeTXT {
			continue
		}
		if owned, _ := names.ToEndpointName(r.DNSName); owned == dnsName {
			out = append(out, r.DNSName)
		}
	}
	sort.Strings(out)
	return out
}

// desiredState renders what a desired state should leave behind, by running it
// through the same AdjustEndpoints the loop uses.
func (h *reconcileLoop) desiredState(desired []*endpoint.Endpoint) []string {
	h.t.Helper()
	source := make([]*endpoint.Endpoint, 0, len(desired))
	for _, ep := range desired {
		source = append(source, ep.DeepCopy())
	}
	adjusted, err := h.reg.AdjustEndpoints(source)
	require.NoError(h.t, err)

	out := make([]string, 0, len(adjusted))
	for _, r := range adjusted {
		out = append(out, recordState(r))
	}
	sort.Strings(out)
	return out
}

// rawRecords renders the record sets Route 53 holds, without going through the provider,
// so a test can tell whether a rejected change batch left anything behind.
func (h *reconcileLoop) rawRecords() []string {
	h.t.Helper()
	var out []string
	for _, rrs := range listAWSRecords(h.t, h.client, reconcileZoneID) {
		values := make([]string, 0, len(rrs.ResourceRecords))
		for _, rr := range rrs.ResourceRecords {
			values = append(values, aws.ToString(rr.Value))
		}
		sort.Strings(values)
		out = append(out, fmt.Sprintf("%s %s -> %s", aws.ToString(rrs.Name), rrs.Type, strings.Join(values, ",")))
	}
	sort.Strings(out)
	return out
}

// route53Change builds one change for a plain record.
func route53Change(action route53types.ChangeAction, dnsName string, recordType route53types.RRType, value string) route53types.Change {
	return route53types.Change{
		Action: action,
		ResourceRecordSet: &route53types.ResourceRecordSet{
			Name:            aws.String(dnsName),
			Type:            recordType,
			TTL:             aws.Int64(300),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String(value)}},
		},
	}
}

func (h *reconcileLoop) submit(changes ...route53types.Change) error {
	h.t.Helper()
	_, err := h.client.ChangeResourceRecordSets(h.t.Context(), &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(reconcileZoneID),
		ChangeBatch:  &route53types.ChangeBatch{Changes: changes},
	})
	return err
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

// txtAffixes covers the supported --txt-prefix and --txt-suffix modes, because the
// ownership TXT name is what decides whether a record is still recognised as ours.
// Wildcard replacement and encryption are left to the registry's own tests.
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

			require.Equalf(t, h.desiredState(desired), h.settledState(),
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
			host := "app." + reconcileZone
			before := h.targets(host)

			_, err := h.settle([]*endpoint.Endpoint{cnameTo("app", tr.to)})
			require.ErrorIsf(t, err, provider.SoftError,
				"upsert-only cannot complete a record type change, Route 53 rejects the batch; zone:\n  %s",
				h.zone())

			// What matters is that the records people resolve are still the ones they
			// were. Ownership TXT records go out in their own change batch, so the ones
			// belonging to the record type which could not be created are left behind.
			require.Equalf(t, before, h.targets(host),
				"the rejected transition changed the records being served; zone:\n  %s", h.zone())
		})
	}
}

// TestReconcileLoopMigratesLegacyAliasOwnershipTXT starts from a zone written by a
// release that labelled A ALIAS records with a "cname-" ownership TXT, and checks that
// upgrading keeps ownership, settles, and can still swap the load balancer afterwards.
func TestReconcileLoopMigratesLegacyAliasOwnershipTXT(t *testing.T) {
	for _, policy := range []string{"sync", "upsert-only"} {
		for _, affix := range txtAffixes {
			t.Run(affix.name+"/"+policy, func(t *testing.T) {
				t.Parallel()
				h := newReconcileLoop(t, reconcileConfig(affix.prefix, affix.suffix), policy, false)
				host := "app." + reconcileZone
				only := map[string]bool{host: true}
				names := mapper.NewAffixNameMapper(affix.prefix, affix.suffix, "")

				// Reproduce the zone a release before #6523 would have left. Rather than
				// hand-encoding the TXT payloads, let the registry write real ones: a
				// "cname-" TXT for the A half, which is the one #6523 moves to "a-", and
				// the "aaaa-" TXT the AAAA half has always had. A CNAME cannot share a
				// name with another type, so the two are published one after the other.
				h.mustSettle([]*endpoint.Endpoint{cnameTo("app", lbExternal)})
				plainCNAME := h.findRecord(host, endpoint.RecordTypeCNAME)
				require.NotNil(t, plainCNAME, "setup: expected a plain CNAME")
				legacyTXT := names.ToTXTName(host, endpoint.RecordTypeCNAME)
				require.NotNilf(t, h.findRecord(legacyTXT, endpoint.RecordTypeTXT),
					"setup: expected the legacy ownership TXT %q; zone:\n  %s", legacyTXT, h.zone())
				require.NoError(t, h.prov.ApplyChanges(t.Context(), &plan.Changes{
					Delete: []*endpoint.Endpoint{plainCNAME},
				}))

				h.mustSettle([]*endpoint.Endpoint{
					endpoint.NewEndpoint(host, endpoint.RecordTypeAAAA, "2001:db8::1"),
				})
				plainAAAA := h.findRecord(host, endpoint.RecordTypeAAAA)
				require.NotNil(t, plainAAAA, "setup: expected a plain AAAA")
				require.NotNilf(t, h.findRecord(names.ToTXTName(host, endpoint.RecordTypeAAAA), endpoint.RecordTypeTXT),
					"setup: expected an aaaa- ownership TXT; zone:\n  %s", h.zone())

				alias := func(recordType string) *endpoint.Endpoint {
					return endpoint.NewEndpoint(host, recordType, lbOld).
						WithProviderSpecific(endpoint.ProviderSpecificAlias, "true").
						WithProviderSpecific(providerSpecificEvaluateTargetHealth, strconv.FormatBool(defaultEvaluateTargetHealth))
				}
				require.NoError(t, h.prov.ApplyChanges(t.Context(), &plan.Changes{
					Delete: []*endpoint.Endpoint{plainAAAA},
					Create: []*endpoint.Endpoint{alias(endpoint.RecordTypeA), alias(endpoint.RecordTypeAAAA)},
				}))

				desired := []*endpoint.Endpoint{cnameTo("app", lbOld)}

				first, err := h.runOnce(desired)
				require.NoError(t, err)
				require.Emptyf(t, first.Delete, "upgrading deleted live records: %v", first.Delete)

				// The new ownership TXT has to be there after the very first loop. If it
				// only turned up later, the record would spend a reconciliation unowned.
				require.NotNilf(t,
					h.findRecord(names.ToTXTName(host, endpoint.RecordTypeA), endpoint.RecordTypeTXT),
					"the a- ownership TXT is missing after the first reconcile; zone:\n  %s", h.zone())
				require.NotNilf(t, h.findRecord(legacyTXT, endpoint.RecordTypeTXT),
					"the legacy TXT was removed instead of being left for the cleanup script; zone:\n  %s", h.zone())

				h.mustSettle(desired)
				for _, recordType := range []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA} {
					require.Truef(t, h.ownsRecord(host, recordType),
						"%s ownership lost while moving from cname- to a-; zone:\n  %s", recordType, h.zone())
				}
				require.Equalf(t, h.desiredState(desired), h.stateFor(only),
					"the migrated record is not what the desired state asks for; zone:\n  %s", h.zone())

				// The load balancer still has to be swappable once ownership has moved.
				after := []*endpoint.Endpoint{cnameTo("app", lbNew)}
				h.mustSettle(after)

				require.Containsf(t, h.targets(host), lbNew, "new load balancer missing; zone:\n  %s", h.zone())
				require.NotContainsf(t, h.targets(host), lbOld,
					"previous load balancer still resolvable after migration; zone:\n  %s", h.zone())
				require.Equalf(t, h.desiredState(after), h.stateFor(only),
					"the swapped record is not what the desired state asks for; zone:\n  %s", h.zone())
				for _, recordType := range []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA} {
					require.Equalf(t, canonicalHostedZone(lbNew), h.aliasHostedZone(host, recordType),
						"%s alias still points at the previous hosted zone; zone:\n  %s", recordType, h.zone())
					require.Truef(t, h.ownsRecord(host, recordType),
						"%s ownership lost while swapping the load balancer; zone:\n  %s", recordType, h.zone())
				}
			})
		}
	}
}

// TestReconcileLoopKeepsOwnershipWhenPreferCNAMEIsDropped covers an operator removing
// --aws-prefer-cname, which turns a hostname that was a plain CNAME into an A ALIAS
// without anything in Kubernetes changing. That rewrites the record type, and with it
// the name of the ownership TXT, which is the transition #6368 was reported over.
//
// Only sync is covered. upsert-only keeps the CNAME, so Route 53 refuses the A beside it
// and the transition never completes, which TestReconcileLoopRecordTypeTransition covers.
func TestReconcileLoopKeepsOwnershipWhenPreferCNAMEIsDropped(t *testing.T) {
	for _, affix := range txtAffixes {
		t.Run(affix.name, func(t *testing.T) {
			t.Parallel()
			h := newReconcileLoop(t, reconcileConfig(affix.prefix, affix.suffix), "sync", true)
			host := "app." + reconcileZone
			only := map[string]bool{host: true}
			names := mapper.NewAffixNameMapper(affix.prefix, affix.suffix, "")
			desired := []*endpoint.Endpoint{cnameTo("app", lbOld)}

			h.mustSettle(desired)
			require.Truef(t, h.ownsRecord(host, endpoint.RecordTypeCNAME),
				"the CNAME is not owned before the flag is dropped; zone:\n  %s", h.zone())
			require.Equalf(t, []string{names.ToTXTName(host, endpoint.RecordTypeCNAME)}, h.ownershipTXTNames(host),
				"unexpected ownership before the flag is dropped; zone:\n  %s", h.zone())

			h.prov.preferCNAME = false

			h.mustSettle(desired)

			require.Equalf(t, h.desiredState(desired), h.stateFor(only),
				"the record did not become an alias pair; zone:\n  %s", h.zone())
			for _, recordType := range []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA} {
				require.Truef(t, h.ownsRecord(host, recordType),
					"%s ownership lost when the record type changed; zone:\n  %s", recordType, h.zone())
			}
			require.Emptyf(t, h.unownedRecords(host), "records left without ownership: %v; zone:\n  %s",
				h.unownedRecords(host), h.zone())
			require.Equalf(t,
				[]string{
					names.ToTXTName(host, endpoint.RecordTypeA),
					names.ToTXTName(host, endpoint.RecordTypeAAAA),
				},
				h.ownershipTXTNames(host),
				"the ownership TXT records did not follow the record type; zone:\n  %s", h.zone())
		})
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

	// Each affix mode is covered by the load balancer swap test above. Randomised
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

						// Records the source dropped stay behind under upsert-only, and
						// they have to keep their ownership just as much as the live ones.
						for _, host := range hosts {
							dnsName := host + "." + reconcileZone
							require.Emptyf(t, h.unownedRecords(dnsName),
								"records for %q without ownership: %v; zone:\n  %s",
								dnsName, h.unownedRecords(dnsName), h.zone())
						}

						require.Equalf(t, h.desiredState(desired), h.stateFor(published),
							"a published hostname settled on a different state than desired; desired:\n  %s\nzone:\n  %s",
							describeEndpoints(desired), h.zone())

						if policy != "sync" {
							continue // upsert-only deliberately keeps records the source dropped
						}
						require.Equalf(t, h.desiredState(desired), h.settledState(),
							"sync settled on a different zone than the desired state; desired:\n  %s",
							describeEndpoints(desired))
						for _, host := range hosts {
							dnsName := host + "." + reconcileZone
							if published[dnsName] {
								continue
							}
							require.Emptyf(t, h.ownershipTXTNames(dnsName),
								"sync left ownership TXT records for %q: %v; zone:\n  %s",
								dnsName, h.ownershipTXTNames(dnsName), h.zone())
						}
					}
				})
			}
		}
	}
}

// TestRoute53APIStubAppliesABatchAtomically pins down that a rejected change batch leaves
// the zone exactly as it was. Route 53 validates and applies a batch all or nothing, and
// the loop tests read the zone after a rejected batch to decide whether records survived.
func TestRoute53APIStubAppliesABatchAtomically(t *testing.T) {
	t.Parallel()
	h := newReconcileLoop(t, reconcileConfig("", ""), "sync", false)
	app, api := "app."+reconcileZone, "api."+reconcileZone

	require.NoError(t, h.submit(route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeA, "1.2.3.4")))
	before := h.rawRecords()

	err := h.submit(
		route53Change(route53types.ChangeActionCreate, api, route53types.RRTypeA, "5.6.7.8"),
		route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeA, "9.9.9.9"),
	)
	require.Error(t, err, "the second change duplicates an existing record set")
	require.Equal(t, before, h.rawRecords(), "the rejected batch left the first change behind")
}

// TestRoute53APIStubRefusesACNAMEBesideAnotherType covers the rule the record type
// transition test depends on.
func TestRoute53APIStubRefusesACNAMEBesideAnotherType(t *testing.T) {
	t.Parallel()
	h := newReconcileLoop(t, reconcileConfig("", ""), "sync", false)
	app := "app." + reconcileZone

	require.NoError(t, h.submit(route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeA, "1.2.3.4")))
	before := h.rawRecords()

	err := h.submit(route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeCname, "elsewhere.example.net"))
	require.ErrorContains(t, err, "conflicts with other records")
	require.Equal(t, before, h.rawRecords(), "the rejected batch left the CNAME behind")
}

// TestRoute53APIStubReplacesACNAMEInOneBatch shows the rule is applied to the state the
// batch ends in. The AWS provider puts creates before deletes, so checking each change on
// its own would refuse a replacement Route 53 accepts.
func TestRoute53APIStubReplacesACNAMEInOneBatch(t *testing.T) {
	t.Parallel()
	h := newReconcileLoop(t, reconcileConfig("", ""), "sync", false)
	app := "app." + reconcileZone

	require.NoError(t, h.submit(route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeCname, "old.example.net")))

	require.NoError(t, h.submit(
		route53Change(route53types.ChangeActionCreate, app, route53types.RRTypeA, "1.2.3.4"),
		route53Change(route53types.ChangeActionDelete, app, route53types.RRTypeCname, "old.example.net"),
	))
	require.Equal(t, []string{"app." + reconcileZone + ". A -> 1.2.3.4"}, h.rawRecords())
}
