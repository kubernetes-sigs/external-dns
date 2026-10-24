```yaml
---
title: "ACME DNS-01 Delegation CNAME Generation"
authors: "@norman-zon"
creation-date: 2026-09-30
status: draft
---
```

# ACME DNS-01 Delegation CNAME Generation

## Table of Contents

<!-- toc -->
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
(on hold until this proposal is decided).

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
