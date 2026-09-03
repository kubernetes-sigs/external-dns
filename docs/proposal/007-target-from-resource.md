```yaml
---
title: "Generic target-from resource dereference annotation"
version: v1alpha1
authors: "@mloiseleur"
creation-date: 2026-07-04
status: draft
---
```

# Generic `target-from` Resource Dereference Annotation

## Table of Contents

<!-- toc -->
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [User Stories](#user-stories)
  - [API](#api)
  - [Behavior](#behavior)
  - [Security Considerations](#security-considerations)
  - [Implementation Steps](#implementation-steps)
  - [Migration and Deprecation](#migration-and-deprecation)
  - [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

ExternalDNS discovers a record's *target* (the IP/hostname it points at) from a fixed set of
per-source locations: a `LoadBalancer` Service, an `Ingress` `status.loadBalancer`, node
addresses, or a literal `target` annotation.

A growing class of setups provisions the public entry point (load balancer, accelerator, gateway)
as a **separate resource** — the address lives in that other resource's `status`.

Each such integration arrives today as its own vendor-specific annotation and code path (`ingress`
for Istio, a proposed `gateway` for Gateway API, a proposed `global-accelerator` for AWS), coupling
a source to a third-party CRD and its API version.

This proposes one generic annotation, `external-dns.kubernetes.io/target-from`, that dereferences
a field on an arbitrary referenced resource and uses the resolved value(s) as the target — one
mechanism instead of the per-product annotations, and **no vendor CRDs** in the sources.

## Motivation

The same request recurs and keeps hitting the same wall — do not couple a generic source to a
specific product's CRD:

- [#6438](https://github.com/kubernetes-sigs/external-dns/pull/6438) — Istio → Gateway API
  `Gateway` (`/hold`: "merging multiple products in a single source").
- [#6190](https://github.com/kubernetes-sigs/external-dns/pull/6190) /
  [#6166](https://github.com/kubernetes-sigs/external-dns/issues/6166) — Ingress/Istio → AWS
  `GlobalAccelerator` (closed: "annotation hack … vendor-specific CRDs baked into the sources").
- [#6457](https://github.com/kubernetes-sigs/external-dns/issues/6457) — Traefik `IngressRoute`
  needing a manual `target`.
- [#4687](https://github.com/kubernetes-sigs/external-dns/issues/4687) — confusion over where a
  Gateway API record's target comes from.

All of these need the same thing: one generic "dereference this CRD field as my target"
annotation, serving any current or future product from any vendor without touching core source
code.

The pattern already exists — the Istio `ingress` annotation dereferences
`Ingress.status.loadBalancer`; `target-from` generalizes that one hard-coded case to any resource
and field, resolved via a dynamic client and the RESTMapper so no source imports a third-party CRD.

### Goals

- One annotation to source a record's target from a field on another resource, addressed by
  group/resource/name.
- Resolve generically, without typed clients or compile-time knowledge of the referenced CRD.
- Reuse the existing Go template engine and `FuncMap` (`source/template`) rather than introducing
  a second expression language.
- Cover the in-flight use cases (Gateway API `Gateway`, AWS Global Accelerator, cross-namespace
  shared LBs) with a single mechanism.
- Preserve target precedence and record-type inference (IP → A/AAAA, hostname → CNAME).
- Offer a deprecation path for the bespoke `ingress` and the proposed per-product annotations.

### Non-Goals

- **Not a replacement for `unstructured`.** `target-from` *augments* a primary source (which
  still supplies hostnames/TTL/provider config); the [`unstructured`
  source](../sources/unstructured.md) *is* the source. The difference is scope, not syntax: both
  use the same templates.
- **Not a new source type.** A cross-cutting annotation resolved by shared code, usable from
  multiple existing sources.
- **Not automatic RBAC.** Granting read access to referenced resources is the operator's job
  (see [Security Considerations](#security-considerations)).

## Proposal

### User Stories

- **Platform operator**: a central Gateway/LB/accelerator lives in a shared namespace. Service
  teams point `target-from` at it instead of hard-coding its address; records follow it when it
  changes.
- **AWS user**: Ingress fronted by a `GlobalAccelerator` from the AWS LB Controller, with no
  external controller copying the value into a `target` annotation.
- **Maintainer**: new "point at product X's address" requests become a documented recipe, not one
  more vendor-CRD coupling.

### API

Key: `external-dns.kubernetes.io/target-from` (subject to `--annotation-prefix`).

Value grammar (`;`-separated for multiple references):

```text
<resource>[.<group>]/[<namespace>/]<name>#<template>
```

- `<resource>` — plural lowercase RBAC name (`gateways`, `services`, `globalaccelerators`).
- `<group>` — API group; omitted for core (`services`).
- `<namespace>` — optional, defaults to the annotated resource's namespace; cross-namespace refs
  are allowed and are the primary use case.
- `<name>` — referenced object's name.
- `<template>` — a Go template evaluated against the referenced object's `unstructured` content,
  using the shared engine and `FuncMap` from `source/template` (`contains`, `trimPrefix`,
  `trimSuffix`, `trim`, `toLower`, `replace`, `isIPv4`, `isIPv6`, `fromJson`). Its output is split
  on `,`, trimmed and de-duplicated — the same convention `--fqdn-template` and `--target-template`
  already use — so one template can yield one target or many. `hasKey` is excluded until its
  signature is generalized (see [Implementation Steps](#implementation-steps)).

Parsing: references are separated by `;`; within a reference, the template is everything after the
**first** `#`. `,` cannot separate references because it is the multi-value separator *inside* a
template's output. The value is tokenized left to right, and a `;` inside a `{{ … }}` action does
not terminate a reference. `<group>` is everything after the first `.` in the first path segment;
no dot means the core group (`services`, `endpoints`).

The API **version** is omitted and resolved via discovery/RESTMapper to the preferred served
version (like `kubectl`), so a CRD graduating `v1beta1` → `v1` does not silently break records.
Explicit-version override: see [Alternatives](#alternatives).

Examples:

```yaml
# Istio Gateway → Gateway API Gateway in a shared namespace
external-dns.kubernetes.io/hostname: my-service.example.com
external-dns.kubernetes.io/target-from: "gateways.gateway.networking.k8s.io/istio-ingress/central-gateway#{{ range .status.addresses }}{{ .value }},{{ end }}"
# Ingress → AWS Global Accelerator (same namespace)
external-dns.kubernetes.io/target-from: "globalaccelerators.aga.k8s.aws/my-accelerator#{{ .status.dnsName }}"
# Shared LoadBalancer Service — hostname and IP from a single reference, IPv4 only
external-dns.kubernetes.io/target-from: "services/lb-ns/shared-lb#{{ range .status.loadBalancer.ingress }}{{ .hostname }},{{ if isIPv4 .ip }}{{ .ip }},{{ end }}{{ end }}"
```

### Behavior

**Scope** — `target-from` resolves *only* the record's targets. Hostnames, TTL, and
provider-specific config still come from the primary resource. The feature is inert unless the
annotation is present, so no feature flag gates it and the upgrade blast radius is limited to
resources that opt in.

**Precedence** — unchanged, with `target-from` where the bespoke `ingress` lookup sits:

1. Literal `target` annotation — highest, unchanged.
2. `target-from` — resolve referenced field(s).
3. Legacy Istio `ingress` — kept as an alias of `target-from` against
   `ingresses.networking.k8s.io/<ref>#{{ range .status.loadBalancer.ingress }}{{ .hostname }},{{ .ip }},{{ end }}`
   during deprecation.
4. Source's native discovery (Service selector, LB status, node IPs) — unchanged.

**Record type** — resolved values reuse `endpoint.SuitableType`: IPv4 → `A`, IPv6 → `AAAA`, else
`CNAME`; mixed IP+hostname produces both, as multi-target LB status does today. Resolved targets
then pass through existing target filters (`--exclude-target-net`, etc.) unchanged.

**Resolution** — a shared resolver keys a dynamic client by the RESTMapper GVR, backed by a
per-GVR informer, so referenced resources are watched and their changes trigger a sync via the
normal event path (no polling). Because cross-namespace references are the primary use case, these
informers watch **all namespaces regardless of `--namespace`** (which scopes only the primary
sources), gated by RBAC. With only `get` granted instead of `list`+`watch`, resolution still works
via direct reads but loses change-driven sync.

**Empty / not-found / missing-field (critical)** — an unresolved reference (object not found,
field absent, empty list) must **not** fall through to native discovery and must **not** delete
records: the endpoint is **skipped for this sync with a warning**, prior state intact. A stale or
empty lookup silently deleting live records is the failure mode this rules out.

Templating needs care here: `text/template` resolves a missing map key to the zero value and
renders `<no value>` rather than failing, which would turn "field absent" into "empty target"
silently. Two guards, both with precedent in-tree:

1. The target-from template is parsed with `Option("missingkey=error")` so an absent key is
a hard template error
2. An empty (or whitespace-only) render is treated as *unresolved*, not as *no targets* — the same "field may not yet be populated" skip the `unstructured` source already does (`source/unstructured_converter.go`).

Neither path may reach the plan as a deletion.

**Edge cases** — multiple refs: de-duplicated union. Template parse/exec error (bad syntax,
missing key, non-scalar value rendered): config error, warn+skip. No RBAC access: warn+skip, never
crash. Unknown GVR (CRD not installed / RESTMapper miss): warn+skip. Cluster-scoped resource:
`<namespace>` must be omitted, supplying one is an error.

### Security Considerations

`target-from` lets a namespaced annotation cause ExternalDNS to read another resource, possibly
cross-namespace. **RBAC is the boundary**: resolution uses ExternalDNS's own ServiceAccount and
reads only what that account is granted, per namespace; no grant means fail closed (warn, skip).

No new default RBAC for arbitrary CRDs ships. Only a resolved field *value* is used, as a DNS
target string — a value a tenant could already set via the literal `target` annotation — so there
is no privilege escalation and no exposure of resource content beyond that field.

The caveat is cluster-wide deployments, where a `ClusterRole` is common and RBAC therefore does
not meaningfully constrain `target-from`.

Residual risk is narrow (deref an unintended CRD field and surface it as a DNS record the tenant
controls), and operators wanting a hard guardrail today can run with namespace-scoped RBAC.

### Implementation Steps

- Shared resolver: parse grammar, RESTMapper lookup, per-GVR dynamic informer, template eval via
  `source/template`. No new dependency, no new expression language.
- One new `source/template` entry point executing against an `unstructured.Unstructured`'s content
  map, reusing the existing `FuncMap` and `execTemplate`'s comma-split/dedupe. Scope
  `Option("missingkey=error")` to the `target-from` template **only**: the existing templates rely
  on today's lenient behavior, so setting it engine-wide would be a breaking change.
- Generalize `hasKey` from `map[string]string` to a map of `any` before advertising it here; as
  written it cannot be called against unstructured content.
- Add `TargetFromKey` to `source/annotations`; wire into `targetsFromX` in sources already honoring
  `target`/`ingress`, starting with Istio + Ingress to cover the open requests.
- Alias the legacy Istio `ingress` annotation onto the resolver, generalizing its `ParseIngress`
  helper into a `ParseNamespacedName` shared by both paths.
- Docs in `docs/annotations/annotations.md`, cross-linked from `unstructured`. Tests mirroring the
  Istio `ingress` coverage plus not-found / empty / RBAC-denied / unknown-GVR / missing key /
  template parse error.

### Migration and Deprecation

- Additive; no existing behavior changes on upgrade.
- Istio `ingress` becomes a documented alias, soft-deprecated (kept working, docs steer to
  `target-from`).
- The per-product annotations (`gateway`, `global-accelerator`) are **not added**; they are
  documented as `target-from` recipes instead.

### Drawbacks

- Dynamic client + per-GVR informer is more machinery than a typed lookup and adds watches.
- A Go template in an annotation is verbose next to a bare field path, and missing keys rendering
  as `<no value>` is a sharp edge; both handled per [Behavior](#behavior), the latter must stay
  covered by tests.
- Under cluster-wide RBAC a tenant could target an unintended resource; see
  [Security Considerations](#security-considerations).
- Version-by-discovery can shift if a CRD's preferred version changes, with the explicit-version
  override as escape hatch.

## Alternatives

- **`unstructured` source (exists).** Reads any CRD field via templates, but *replaces* the
  primary source — you lose its native hostname/TTL/provider handling. `target-from` keeps the
  primary source and only borrows the address; complementary, docs point to whichever fits, and
  both are written in the same template syntax.
- **JSONPath instead of Go templates** (`#status.addresses[*].value`). Terser for the common "one
  field" case, and it errors natively on a missing field.

  It would be a second expression language in a project already standardized on `text/template` —
  `source/template` backs `--fqdn-template`, `--target-template`, `--fqdn-target-template` and
  the `unstructured` source.

  No JSONPath exists in the tree today. It also buys no capability: `[*]` maps onto `range`, multi-value output
  is already handled by the engine's comma-split/dedupe, and templates add conditionals plus the
  existing helpers. Two syntaxes for "read a field off a CRD" is the cost; terseness the only gain.
- **Per-product annotations (`ingress`, `gateway`, `global-accelerator`, …).** The status quo:
  each couples a source to a vendor CRD/version, multiplies code paths, repeatedly rejected
  (#6438, #6190). `target-from` is the generalization that removes the coupling.
- **External controller copies the address into `target`.** No ExternalDNS change, but pushes an
  extra controller, with its own RBAC and failure modes, onto every user who needs this.
- **Explicit-version grammar** (`<resource>.<group>/<version>/<ns>/<name>#<template>`). More
  precise, but pins a version that a CRD will eventually graduate past; offered as an optional
  override on the discovery-based default, not the default.
- **Do nothing.** Each new "point at product X" request keeps arriving as a coupling PR that gets
  held or closed, leaving users on manual `target` or the heavier `unstructured` source.
