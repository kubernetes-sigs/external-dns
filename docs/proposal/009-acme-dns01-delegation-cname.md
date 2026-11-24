```yaml
---
title: "ACME DNS-01 Delegation CNAME Generation"
authors: "@norman-zon"
creation-date: 2026-09-30
status: not-planned
---
```

# ACME DNS-01 Delegation CNAME Generation

## Table of Contents

<!-- toc -->
- [Decision](#decision)
- [Threat Model](#threat-model)
- [Out-of-Tree Shapes](#out-of-tree-shapes)
  - [Provider-Side Credential Scoping](#provider-side-credential-scoping)
  - [Webhook Solver with a Dedicated crd Instance](#webhook-solver-with-a-dedicated-crd-instance)
  - [Webhook Source Sidecar for the Delegation CNAMEs](#webhook-source-sidecar-for-the-delegation-cnames)
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [Scope](#scope)
  - [User Stories](#user-stories)
  - [API](#api)
  - [Behavior](#behavior)
    - [Record Generation](#record-generation)
    - [Wildcards and Deduplication](#wildcards-and-deduplication)
    - [Precedence of Explicit Records](#precedence-of-explicit-records)
    - [Ownership and Cleanup](#ownership-and-cleanup)
    - [Pipeline Placement](#pipeline-placement)
    - [Startup Validation](#startup-validation)
  - [Drawbacks](#drawbacks)
  - [Open Questions](#open-questions)
- [Alternatives](#alternatives)
<!-- /toc -->

## Decision

**Rejected** after review on [#6777](https://github.com/kubernetes-sigs/external-dns/pull/6777):

- ExternalDNS manages records that follow directly from a Kubernetes resource. A record derived
  from one protocol's convention widens that scope, and the reasoning that keeps new providers out
  of tree ([#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347)) applies to
  protocol-specific sources as well.
- The security motivation is narrower than the proposal states. ExternalDNS already holds a
  zone-scoped credential for the production zones, so delegation removes one principal with
  production write access; it does not shrink what ExternalDNS itself can do. Several providers
  also scope credentials below zone level, which covers much of the motivation without any CNAME.
- The CNAMEs can be generated out of tree. The preferred shape is a sidecar behind the
  [webhook source](008-webhook-source.md) that returns one `_acme-challenge` CNAME per hostname.
  Users without the provider-enforced requirement from the threat model can instead run a
  cert-manager webhook solver against a dedicated `crd` instance, which needs no CNAME at all.

The reference implementation [#6548](https://github.com/kubernetes-sigs/external-dns/pull/6548) is
closed. The threat model and the out-of-tree shapes below record the review; the original design
follows them for reference.

## Threat Model

The threat is a compromised ACME client or a leaked ACME DNS credential. With a zone-scoped
credential that means rewriting production A, AAAA and CNAME records in that zone, not only issuing
certificates, which the ACME client can do by design. Delegation reduces the ACME credential to TXT
writes in a dedicated challenge zone, enforced by the DNS provider. It does not change what
ExternalDNS can do; it removes one principal with production write access.

The requirement behind this proposal is provider-enforced: the ACME path, that is the ACME client
and whatever solver it calls, holds credentials for the challenge zone only and can never write into
a production zone. Enforcement inside the cluster, by RBAC or an admission policy, does not satisfy
it, because the production-zone write then still happens, with the credential of whichever component
applies the record.

[RFC 8555 §8.4](https://www.rfc-editor.org/rfc/rfc8555#section-8.4) fixes the validation name at
`_acme-challenge.<host>`, inside the host's zone. The only way to keep the ACME write out of the
production zone is therefore the delegation CNAME into the challenge zone, which ACME CAs follow and
which cert-manager supports with
[`cnameStrategy: Follow`](https://cert-manager.io/docs/configuration/acme/dns01/#delegated-domains-for-dns01).
Under this requirement the per-hostname CNAMEs stay with any solver. What this proposal decided is
only whether generating them from existing sources lives in tree or out of tree.

## Out-of-Tree Shapes

### Provider-Side Credential Scoping

Several providers scope a credential below zone level. A credential limited to TXT records at
`_acme-challenge.*` names covers most of the threat model without any CNAME:

- Route 53: the IAM condition keys `route53:ChangeResourceRecordSetsNormalizedRecordNames`,
  `route53:ChangeResourceRecordSetsRecordTypes` and `route53:ChangeResourceRecordSetsActions`
  ([AWS documentation](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/specifying-rrset-conditions.html)).
- Cloud DNS: IAM bindings per managed zone, with IAM Conditions on individual resource record sets
  ([Google Cloud documentation](https://cloud.google.com/dns/docs/zones/iam-per-resource-zones)).
- NS1: API key and team permissions with `records_allow` entries scoped by zone, domain, subdomain
  inclusion and record type
  ([ns1-go model](https://github.com/ns1/ns1-go/blob/v2/rest/model/account/permissions.go)).

This is provider-specific and not available everywhere. Delegation remains the tightest option and
the only provider-agnostic one.

### Webhook Solver with a Dedicated crd Instance

cert-manager calls two methods on a DNS-01 webhook solver: `Present` writes the challenge TXT
record, `CleanUp` removes it. A solver that creates a `DNSEndpoint` in `Present` and deletes it in
`CleanUp` leaves the DNS write to the `crd` source, and cert-manager holds no DNS credential at all.
Three such solvers exist:
[youniqx/cert-manager-webhook-external-dns](https://github.com/youniqx/cert-manager-webhook-external-dns),
[AbsaOSS/cert-manager-webhook-externaldns](https://github.com/AbsaOSS/cert-manager-webhook-externaldns)
and
[ThatsMrTalbot/acme-webhook-external-dns](https://github.com/ThatsMrTalbot/acme-webhook-external-dns).

On its own this shape does not meet the threat model: the TXT record still lands in the production
zone, written with the credential of ExternalDNS. It is compatible with it when combined with the
delegation CNAMEs: `cnameStrategy: Follow` makes the solver write the delegated name in the
challenge zone, and a dedicated ExternalDNS instance applies it. That instance:

- runs `--source=crd` only, with `--domain-filter` set to the challenge zone and a credential for
  that zone only;
- needs `--managed-record-types=TXT`, because the default set (A, AAAA, CNAME) drops the challenge
  records;
- needs `--policy=sync`, because `upsert-only` never removes the challenge records;
- should run with `--events`, so that a new `DNSEndpoint` triggers a reconcile within
  `--min-event-sync-interval` instead of waiting for `--interval`.

The trust boundary of this instance is cluster RBAC on `DNSEndpoint` in the watched namespace plus
an admission policy that limits `dnsName` to `_acme-challenge.*` and `recordType` to TXT. Anything
that can create a `DNSEndpoint` there can write what the instance's credential allows. In posture
this equals a native cert-manager solver with a zone-scoped credential.

Two defects were found while reviewing this shape. The code was read at
[c76b50e](https://github.com/kubernetes-sigs/external-dns/commit/c76b50eca07644bc1782ea3803213e3ee3d7fa51),
not run against live zones:

- **Concurrent challenges for one name.** An apex name and its wildcard in one `Certificate`
  produce two challenges at the same `_acme-challenge.<name>` with different keys. The plan keys its
  rows by name and set identifier
  ([plan.go#L69](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/plan/plan.go#L69)).
  A solver that names one object per FQDN fails the second `Present` with `AlreadyExists` until the
  first challenge is cleaned up. A solver that gives each challenge its own set identifier makes the
  plan emit two creates for the same name and type; the Google provider submits them as two record
  sets in one change
  ([google.go#L264](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/provider/google/google.go#L264)),
  which the Cloud DNS API rejects. A correct solver merges all keys for one name into one
  `DNSEndpoint` with several targets. None of the three solvers above does.
- **Providers without set-identifier support.** NS1 returns records without an identifier
  ([ns1.go#L165](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/provider/ns1/ns1.go#L165)).
  With a set identifier on the desired side, the reconcile after the record was created plans a
  delete of the old record and a create of the new one; the create runs first, the provider returns
  `ErrRecordExists` as a plain error
  ([ns1.go#L245](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/provider/ns1/ns1.go#L245)),
  and the controller exits
  ([controller.go#L197](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/controller/controller.go#L197)).
  A single challenge is enough to trigger it. This is an ExternalDNS defect in its own right, the
  plan keyed by a set identifier on a provider that cannot store one, and is tracked outside this
  proposal.

### Webhook Source Sidecar for the Delegation CNAMEs

The preferred out-of-tree shape for generating the CNAMEs is a sidecar behind the
[webhook source](008-webhook-source.md). It returns one `_acme-challenge.<host>` CNAME per
hostname, and ExternalDNS treats them like the endpoints of any other source.

- The webhook source is a draft and not implemented as of v0.24.0. Until it lands, the available
  out-of-tree shape is a controller that writes `DNSEndpoint` objects (Alternative 3 below).
- The sidecar has to discover the hostnames itself. It runs on loopback in the ExternalDNS pod and
  shares its service account, so it reads the same `Ingress`, `HTTPRoute` and `Service` objects
  with the RBAC that already exists. Importing the `source` package of this module reuses hostname
  extraction and `--fqdn-template`, which moves the wrapper from #6548 out of tree without changes.
  The package carries no API stability promise.
- The generated records pass through the plan and the registry as usual: they get ownership TXT
  records, are deleted when their hostname disappears, and follow `--policy`. They appear within
  one `--interval`, because the webhook source has no event handler; for a static per-host record
  that is acceptable.
- Wildcard handling, deduplication of a wildcard and its base name, and skipping of names that
  already begin with `_acme-challenge.` stay the sidecar's job.
- The precedence of explicit records from the original design does not carry over. When the `crd`
  source and the sidecar both produce the same challenge name, the plan's conflict resolver
  decides: the smallest target on first creation, the owning resource afterwards
  ([conflict.go](https://github.com/kubernetes-sigs/external-dns/blob/c76b50eca07644bc1782ea3803213e3ee3d7fa51/plan/conflict.go)).

## Summary

Add an opt-in source wrapper that creates an ACME DNS-01 delegation CNAME
(`_acme-challenge.<hostname>`) for every hostname that ExternalDNS discovers from its
sources. The CNAME points into a separate challenge zone, and a Go template controls the
target name:

```text
_acme-challenge.app.example.com CNAME app.example.com.acme.example.net
```

With these CNAMEs in place, ACME clients such as cert-manager can solve DNS-01 challenges
while holding write access only to the challenge zone, not to the production zones.
Nothing changes unless the feature is enabled.

Tracking issue: [#6535](https://github.com/kubernetes-sigs/external-dns/issues/6535).
Reference implementation: [#6548](https://github.com/kubernetes-sigs/external-dns/pull/6548)
(closed with this decision).

## Motivation

DNS-01 is the only ACME challenge type that can issue wildcard certificates. It is also the
usual choice when the endpoints are not reachable from the internet. It requires the ACME
client to publish a TXT record at `_acme-challenge.<hostname>`. By default, the client
therefore needs write access to every zone that contains a hostname it issues certificates
for.

For many organizations that is too much. cert-manager credentials that can write to
`example.com` can also change any production record in that zone. The established way to
reduce this is delegation: a static CNAME at `_acme-challenge.<hostname>` points into a
dedicated zone, and the ACME client may write only to that zone. Both
[cert-manager](https://cert-manager.io/docs/configuration/acme/dns01/#delegated-domains-for-dns01)
and [acme-dns](https://github.com/joohoi/acme-dns) document this pattern. ACME CAs follow
the CNAME when they validate the challenge.

The CNAMEs themselves are the operational problem. Every hostname needs one. On platforms
where application teams create `Ingress`, `HTTPRoute` or `Service` hostnames themselves,
new hostnames appear at any time. The CNAMEs are then managed in one of these ways:

- **Manually or with OpenTofu/Terraform.** This works for static hostnames. Every new
  hostname that a team creates also needs a change in a central repository.
- **One `DNSEndpoint` per application.** This duplicates the hostname, is easy to forget,
  and requires every team to know the delegation convention.
- **A separate controller or policy engine** (for example a Kyverno `generate` rule that
  creates `DNSEndpoint` resources). This adds another component. It also reimplements the
  source parsing, domain filtering and ownership handling that ExternalDNS already has.

ExternalDNS already knows the full set of hostnames and already has write access to the
zones where the delegation CNAMEs have to live. It is the component with the least extra
work and the least extra privilege for this job.

### Goals

- Opt-in generation of `_acme-challenge.<hostname>` CNAMEs for hostnames from any source.
- A configurable mapping from hostname to CNAME target, so that both common delegation
  layouts are supported: a nested name (`app.example.com.acme.example.net`) and a flattened
  name as used by acme-dns (`app-example-com.acme.example.net`).
- Correct wildcard handling, following RFC 8555: `*.app.example.com` is validated at
  `_acme-challenge.app.example.com`.
- The generated records use the existing registry. They get ownership TXT records, and
  they are deleted when their hostname disappears.
- No provider-specific code. The feature works with every in-tree and webhook provider that
  accepts CNAMEs whose names begin with an underscore.
- No change in behavior when the feature is disabled.

### Non-Goals

- Solving ACME challenges or writing challenge TXT records. That remains the job of the
  ACME client.
- Managing the challenge zone (`acme.example.net`) or any records in it.
- Integrating with cert-manager or its CRDs. The feature relies only on DNS conventions
  from RFC 8555.
- Delegation schemes other than a CNAME at `_acme-challenge` (for example NS delegation of
  `_acme-challenge` names).

## Proposal

### Scope

In the issue, the maintainers noted that this would widen the scope of ExternalDNS. Today,
ExternalDNS manages records that follow directly from a Kubernetes resource. This feature
adds records that are **derived** from those resources. Before deciding on the mechanism,
we should answer one question: should ExternalDNS create derived records at all?

ExternalDNS already has precedents for derived records:

- **`--create-ptr`** (source wrapper) creates PTR records that no source declares
  explicitly.
- **NAT64** (source wrapper) creates AAAA records from A-record targets.
- **The TXT registry** writes its own ownership records next to every managed record.

The delegation CNAME fits this pattern. It is created only from hostnames that ExternalDNS
already manages, it lives in the same zones, it goes through the same plan, policy and
registry path, and it adds no new provider permissions. The scope question is therefore
narrower than "should ExternalDNS support ACME". It is: **should there be one more derived
record type, at a well-known name, in a well-known format?**

If maintainers prefer a general mechanism to a feature built for one use case, see
[Alternative 1](#alternative-1-generic-derived-record-templates).

### User Stories

- **Platform team with self-service routes:** "Application teams create `HTTPRoute`s with
  their own hostnames. cert-manager issues certificates through DNS-01. cert-manager must
  not be able to write to our production zones, and teams must not have to create DNS
  records themselves."
- **Wildcard certificates per environment:** "Each preview environment gets a
  `*.pr-123.dev.example.com` certificate. The delegation CNAME must exist as soon as the
  Gateway listener or route exists, and it must disappear when the environment is torn
  down."
- **acme-dns users:** "Our challenge server expects flat labels such as
  `app-example-com.acme.example.net`. We need to control the mapping from hostname to
  target."
- **Partial adoption:** "Only `*.internal.example.com` uses delegated DNS-01. Other zones
  use HTTP-01 or a different setup, and they must not get any extra records."

### API

The reference implementation adds three flags. Each flag can also be set as an
environment variable, like all other flags.

| Flag | Description |
| :--- | :--- |
| `--acme-cname-delegation-target-template` | Go template for the CNAME target. Setting it enables the feature. This follows `--nat64-networks`, where setting a value is the switch and there is no separate enable flag. |
| `--acme-cname-delegation-domain-filter` | Creates CNAMEs only for hostnames that match these domain suffixes. Can be repeated. Default: all hostnames. |
| `--acme-cname-delegation-ttl` | TTL of the generated CNAMEs. Default: `--min-ttl` or the provider default. |

The template receives the following fields. It can use the shared function set of
`--fqdn-template` (`replace`, `trimPrefix`, `toLower`, and so on).

| Field | `app.example.com` | `*.app.example.com` |
| :--- | :--- | :--- |
| `{{ .Hostname }}` | `app.example.com` | `*.app.example.com` |
| `{{ .HostnameWithoutWildcard }}` | `app.example.com` | `app.example.com` |

Example:

```sh
external-dns \
  --source=ingress \
  --source=gateway-httproute \
  --domain-filter=example.com \
  --acme-cname-delegation-target-template='{{ .HostnameWithoutWildcard }}.acme.example.net' \
  --acme-cname-delegation-ttl=300s
```

For an `HTTPRoute` with the hostnames `app.example.com` and `*.dev.example.com`,
ExternalDNS manages the usual records plus these:

```text
_acme-challenge.app.example.com CNAME app.example.com.acme.example.net
_acme-challenge.dev.example.com CNAME dev.example.com.acme.example.net
```

Flattened layout for acme-dns:

```sh
--acme-cname-delegation-target-template='{{ replace "." "-" .HostnameWithoutWildcard }}.acme.example.net'
```

### Behavior

#### Record Generation

- A CNAME is created for every endpoint of type A, AAAA or CNAME whose name matches
  `--acme-cname-delegation-domain-filter`. Other record types (TXT, MX, SRV, and so on) do
  not carry hostnames that need certificates, so they are ignored.
- Endpoints whose names already start with `_acme-challenge.` are skipped. This prevents a
  recursive `_acme-challenge._acme-challenge.…` record.
- If the template fails or renders an empty string for a hostname, that hostname is
  skipped and a warning is logged. The error is soft and does not stop the reconcile loop.

#### Wildcards and Deduplication

- ACME validates `*.app.example.com` at `_acme-challenge.app.example.com`, so the leading
  `*.` is removed from the challenge name.
- Each challenge name gets exactly one CNAME. Duplicates are removed in two cases: when the
  same hostname has both an A and an AAAA record, and when a wildcard exists together with
  its base name (`*.app.example.com` and `app.example.com`). Both names share one challenge
  name.

#### Precedence of Explicit Records

If a source already returns an endpoint named `_acme-challenge.<hostname>` (for example
from a `DNSEndpoint`), no CNAME is generated for that name. A delegation record declared
explicitly always takes precedence. This also gives a way to override individual
hostnames before per-resource annotations exist.

#### Ownership and Cleanup

The generated CNAMEs are regular endpoints. They carry the resource label and object
references of the endpoint they were derived from. The registry therefore creates
ownership TXT records for them as usual, and the plan deletes them when their hostname
disappears. `--policy=upsert-only` and `create-only` apply to them in the same way as to
other records.

#### Pipeline Placement

The wrapper runs after deduplication and the NAT64, target-filter and PTR wrappers, and
before the post-processor. Because it runs before the post-processor, `--min-ttl` also
applies to the generated records.

#### Startup Validation

- ExternalDNS refuses to start if the feature is enabled and CNAME is not in
  `--managed-record-types`.
- `--acme-cname-delegation-domain-filter` and `--acme-cname-delegation-ttl` require
  `--acme-cname-delegation-target-template`. A negative TTL is rejected.
- A warning is logged when `--regex-domain-filter` is also set. The regular expression must
  also match the generated `_acme-challenge.*` names, otherwise the plan filters them out.

### Drawbacks

- **Wider scope.** ExternalDNS gains built-in knowledge of an ACME convention. The feature
  is small (one wrapper of about 165 lines) and isolated, but it is a precedent for further
  derived record types that are tied to one protocol.
- **CNAME exclusivity.** A CNAME at `_acme-challenge.<hostname>` rules out any other record
  at that name. This includes TXT records written by an ACME client that validates
  directly in the production zone. Mixing delegated and direct DNS-01 for the same names
  breaks the direct setup. `--acme-cname-delegation-domain-filter` limits the risk but
  cannot remove it.
- **Migration from manually managed CNAMEs.** Delegation CNAMEs that already exist and have
  no ExternalDNS ownership record are not considered owned. ExternalDNS then tries to
  create a record that already exists. Depending on the provider, this fails or conflicts.
  Operators have to adopt or remove the existing records before they enable the feature.
  This matches today's behavior for any other record that ExternalDNS did not create.
- **Provider support for names that begin with an underscore.** A few providers reject such
  names. The feature does not add provider-specific checks. This is documented as a caveat
  instead.
- **Record count.** Each managed hostname gets up to two more records: the CNAME and its
  ownership TXT record. For large zones this matters for provider quotas and API rate
  limits.

### Open Questions

1. **Generic or specific?** Is a feature built for ACME acceptable, or should it be
   expressed as a generic derived-record mechanism
   ([Alternative 1](#alternative-1-generic-derived-record-templates))?
2. **Per-resource control.** Should resources be able to opt in or out, for example with
   `external-dns.alpha.kubernetes.io/acme-delegation: "false"`? The reference implementation
   leaves this out on purpose. It would need `ProviderSpecific` plumbing through all
   sources, similar to the PTR `record-type` mechanism. The domain filter and the
   precedence of explicit records cover most cases for now.
3. **Restriction by source type.** Is a flag such as
   `--acme-cname-delegation-sources=gateway-httproute,ingress` worth adding, or is the
   domain filter enough?
4. **Flag naming.** `--acme-cname-delegation-*` or a shorter prefix such as
   `--acme-delegation-*`?

## Alternatives

### Alternative 1: Generic derived-record templates

Add a general mechanism that emits extra records from templates. An example is a
repeatable flag `--derived-record='{name}={type}:{target-template}'`, which would express
this use case as
`_acme-challenge.{{ .HostnameWithoutWildcard }}=CNAME:{{ .HostnameWithoutWildcard }}.acme.example.net`.

- **Pro:** ExternalDNS contains no ACME-specific code. The same mechanism covers other
  conventions, such as `_dmarc` or verification records.
- **Con:** It has a much larger API surface and more complex validation, for example of
  record types, name collisions and CNAME exclusivity. The ACME-specific behavior still has
  to be built into the mechanism: wildcard removal, deduplication of wildcard and base
  name, and skipping of challenge names. Without a second concrete use case, this is a
  general design that nobody needs yet.

### Alternative 2: `DNSEndpoint` per application

Teams create a `DNSEndpoint` for their delegation CNAME next to their route.

- **Pro:** No change to ExternalDNS.
- **Con:** It duplicates the hostname, is easy to forget, and puts the delegation
  convention on every team. Wildcard and deduplication rules have to be applied by hand.

### Alternative 3: External controller or policy engine

A Kyverno `generate` rule or a small controller watches `Ingress` and `HTTPRoute` resources
and creates `DNSEndpoint` resources.

- **Pro:** It stays outside ExternalDNS.
- **Con:** It adds another component to run and to secure. It duplicates the hostname
  extraction for each source type (Ingress rules, Gateway listener and route hostname
  intersection, Service annotations), which ExternalDNS already does. It also does not see
  hostnames that come from `--fqdn-template`.

### Alternative 4: Manage delegation CNAMEs with IaC

Keep the CNAMEs in OpenTofu, Terraform or a similar tool.

- **Pro:** No change to any Kubernetes component.
- **Con:** It does not work for hostnames that teams create themselves or that change at
  runtime, such as preview environments.

### Alternative 5: Status quo

The ACME client gets write access to the production zones.

- **Pro:** Nothing to build.
- **Con:** This is the blast radius the delegation pattern exists to avoid.
