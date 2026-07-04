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

A growing class of setups provisions the public entry point (load balancer, accelerator,
gateway) as a **separate resource** from the one owning the hostnames — the address lives in
that other resource's `status`. Today each such integration arrives as its own vendor-specific
annotation with its own code path baked into the sources (`ingress` for Istio, a proposed
`gateway` for Gateway API, a proposed `global-accelerator` for AWS), each coupling a source to a
third-party CRD and its API version.

This proposes one generic annotation, `external-dns.kubernetes.io/target-from`, that
dereferences a named field on an arbitrary referenced resource and uses the resolved value(s) as
the target. It replaces the per-product annotations with one mechanism and imports **no vendor
CRDs** into the sources.

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

A maintainer sketched the resolution directly in the #6190 review:

> `external-dns.alpha.kubernetes.io/target-from: "aga.k8s.aws/v1beta1/globalaccelerators/namespace/name#status.dnsName"`
> A single generic "dereference this CRD field as my target" annotation. That would serve the
> AWS GA use case, any future GA-like product from any cloud, and be maintainable without
> touching core source code.

This formalizes that: defines the grammar/behavior and specifies generic resolution (dynamic
client + RESTMapper) so no source imports a third-party CRD. The pattern already exists — the
Istio `ingress` annotation dereferences `Ingress.status.loadBalancer`; `target-from` generalizes
that one hard-coded case to any resource and field.

### Goals

- One annotation to source a record's target from a field on another resource, addressed by
  group/resource/name.
- Resolve generically, without typed clients or compile-time knowledge of the referenced CRD.
- Cover the in-flight use cases (Gateway API `Gateway`, AWS Global Accelerator, cross-namespace
  shared LBs) with a single mechanism.
- Preserve target precedence and record-type inference (IP → A/AAAA, hostname → CNAME).
- Offer a deprecation path for the bespoke `ingress` and the proposed per-product annotations.

### Non-Goals

- **Not a templating engine.** Arbitrary transformation is served by the
  [`unstructured` source](../sources/unstructured.md); `target-from` resolves a field path to a
  scalar or list of scalars, nothing more.
- **Not a replacement for `unstructured`.** `target-from` *augments* a primary source (which
  still supplies hostnames/TTL/provider config); `unstructured` *is* the source.
- **Not a new source type.** A cross-cutting annotation resolved by shared code, usable from
  multiple existing sources.
- **Not automatic RBAC.** Granting read access to referenced resources is the operator's job
  (see [Security Considerations](#security-considerations)).

## Proposal

### User Stories

- **Platform operator**: a central Gateway/LB/accelerator lives in a shared namespace. Service
  teams annotate their Istio `Gateway`/`VirtualService`/`Ingress`/route with `target-from`
  pointing at it; records follow that entry point automatically, teams never hard-code its
  address, and records update when it changes.
- **AWS user**: Ingress fronted by a `GlobalAccelerator` from the AWS LB Controller — point
  `target-from` at `globalaccelerators.aga.k8s.aws/.../<name>#status.dnsName`, no external
  controller copying the value into a `target` annotation.
- **Maintainer**: stop accepting one vendor-CRD coupling per product; point new "point at product
  X's address" requests at one documented, product-agnostic annotation.

### API

Key: `external-dns.kubernetes.io/target-from` (subject to `--annotation-prefix`).

Value grammar (comma-separated for multiple references):

```
<resource>[.<group>]/[<namespace>/]<name>#<fieldPath>
```

- `<resource>` — plural lowercase RBAC name (`gateways`, `services`, `globalaccelerators`).
- `<group>` — API group; omitted for core (`services`).
- `<namespace>` — optional, defaults to the annotated resource's namespace; cross-namespace refs
  are allowed and are the primary use case.
- `<name>` — referenced object's name.
- `<fieldPath>` — a [JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath/) resolving
  to a scalar (one target) or list of scalars (many); `[*]` selects across list elements.
  Surrounding `{}` are optional.

Parsing: `<group>` is everything after the first `.` in the first path segment; no dot means the
core group (`services`, `endpoints`).

The API **version** is omitted and resolved via discovery/RESTMapper to the preferred served
version (like `kubectl`), so records don't silently break on a CRD's `v1beta1` → `v1` graduation
(the exact fragility raised against #6190). Explicit-version override: see Alternatives.

Examples:

```yaml
# Istio Gateway → Gateway API Gateway in a shared namespace
external-dns.kubernetes.io/hostname: my-service.example.com
external-dns.kubernetes.io/target-from: "gateways.gateway.networking.k8s.io/istio-ingress/central-gateway#status.addresses[*].value"
# Ingress → AWS Global Accelerator (same namespace)
external-dns.kubernetes.io/target-from: "globalaccelerators.aga.k8s.aws/my-accelerator#status.dnsName"
# Any resource → shared LoadBalancer Service (IP and hostname)
external-dns.kubernetes.io/target-from: "services/lb-ns/shared-lb#status.loadBalancer.ingress[*].hostname,services/lb-ns/shared-lb#status.loadBalancer.ingress[*].ip"
```

### Behavior

**Scope** — `target-from` resolves *only* the record's targets. Hostnames, TTL, and
provider-specific config still come from the primary resource. The feature is inert unless the
annotation is present on a resource; no feature flag gates it, so upgrade blast radius is limited
to resources that opt in.

**Precedence** — unchanged, with `target-from` where the bespoke `ingress` lookup sits:

1. Literal `target` annotation — highest, unchanged.
2. `target-from` — resolve referenced field(s).
3. Legacy Istio `ingress` — kept as an alias of `target-from` against
   `ingresses.networking.k8s.io/<ref>#status.loadBalancer.ingress[*]` during deprecation.
4. Source's native discovery (Service selector, LB status, node IPs) — unchanged.

**Record type** — resolved values reuse `suitableType`: IPv4 → `A`, IPv6 → `AAAA`, else `CNAME`;
mixed IP+hostname produces both, as multi-target LB status does today. Resolved targets then pass
through existing target filters (`--exclude-target-net`, etc.) unchanged.

**Resolution** — a shared resolver keys a dynamic client by the RESTMapper GVR, backed by a
per-GVR informer, so referenced resources are watched/cached and their changes trigger a sync via
the normal event path (no polling). Because cross-namespace references are the primary use case,
these referenced-resource informers watch **all namespaces regardless of `--namespace`** (which
scopes only the primary sources); the watch reads only the referenced GVRs, gated by RBAC. This
requires `list`+`watch` (not just `get`) on those GVRs; with only `get` granted, resolution still
works via direct reads but loses change-driven sync.

**Empty / not-found / missing-field (critical)** — an unresolved reference (object not found,
field absent, empty list) must **not** fall through to native discovery and must **not** delete
records: the endpoint is **skipped for this sync with a warning**, prior state intact. Avoids the
#6438 failure mode (stale/empty lookup silently deleting records) and aligns with its
log-and-continue resilience fix.

**Edge cases** — multiple refs: de-duplicated union. Non-scalar JSONPath result: config error,
warn+skip. No RBAC access: warn+skip, never crash. Unknown GVR (CRD not installed / RESTMapper
miss): warn+skip. Cluster-scoped resource: `<namespace>` must be omitted, supplying one is an
error.

### Security Considerations

`target-from` lets a namespaced annotation cause ExternalDNS to read another resource, possibly
cross-namespace.

**RBAC is the boundary.** Resolution uses ExternalDNS's own ServiceAccount; it reads only what
that account is granted (`get`/`list`/`watch` on the intended `resource.group`, per namespace).
No grant → fail closed (warn, skip). No new default RBAC for arbitrary CRDs ships. RBAC is
strictly finer-grained than any resource-name allowlist (it also scopes by namespace), so an
operator who wants to constrain `target-from` scopes the ServiceAccount's Role accordingly.

Only a resolved field *value* is used, as a DNS target string — a value a tenant could already
set via the literal `target` annotation — so there is no privilege escalation and no exposure of
resource content beyond that one field.

**Caveat for cluster-wide deployments.** ExternalDNS is frequently run with a cluster-wide
`ClusterRole` (sources need services/ingresses/gateways/routes across all namespaces when
`--namespace` is empty). In that setup RBAC does not meaningfully constrain `target-from`. The
residual risk is narrow (deref an unintended CRD field and surface it as a DNS record the tenant
controls). A GVR allowlist flag (`--target-from-allowed-resources`) is deliberately left as a
**future option** rather than shipped in v1alpha1; operators needing a hard guardrail today can
run ExternalDNS with namespace-scoped RBAC. Open question for review: is the allowlist worth
adding, or is RBAC sufficient?

### Implementation Steps

- Shared resolver: parse grammar, RESTMapper lookup, per-GVR dynamic informer, JSONPath eval
  (`k8s.io/client-go/util/jsonpath`).
- Add `TargetFromKey` to `source/annotations`; wire into `targetsFromX` in sources already
  honoring `target`/`ingress` (Istio Gateway, Istio VirtualService, Ingress, Gateway API routes,
  Service, CRD/DNSEndpoint). Start with Istio + Ingress to cover the open requests.
- Alias the legacy Istio `ingress` annotation onto the resolver; keep the `ParseIngress` →
  `ParseNamespacedName` generalization (staged in #6438).
- Docs: new section in `docs/annotations/annotations.md`, cross-linked from `unstructured`.
- Tests mirroring the Istio `ingress` coverage plus not-found / empty / RBAC-denied / unknown-GVR.

### Migration and Deprecation

- Additive; no existing behavior changes on upgrade.
- Istio `ingress` becomes a documented alias, soft-deprecated (kept working, docs steer to
  `target-from`).
- The per-product annotations proposed in #6438 (`gateway`) and #6190 (`global-accelerator`) are
  **not added**; documented as `target-from` recipes instead. #6438 can close keeping only its
  two orthogonal bug fixes; #6190 stays closed.

### Drawbacks

- Dynamic client + per-GVR informer is more machinery than a typed lookup and adds watches.
- JSONPath in an annotation is a small language users can get wrong; mitigated by warn-and-skip
  errors and documented recipes.
- Under cluster-wide RBAC a tenant could target an unintended resource; mitigated by
  namespace-scoped RBAC, with a GVR allowlist flag as a possible future guardrail.
- Version-by-discovery can shift if a CRD's preferred version changes; safer than baking in a
  version, with the explicit-version override as escape hatch.

## Alternatives

- **`unstructured` source (exists).** Reads any CRD field via templates, but *replaces* the
  primary source — you lose its native hostname/TTL/provider handling. `target-from` keeps the
  primary source and only borrows the address; complementary, docs point to whichever fits.
- **Per-product annotations (`ingress`, `gateway`, `global-accelerator`, …).** The status quo:
  each couples a source to a vendor CRD/version, multiplies code paths, repeatedly rejected
  (#6438, #6190). `target-from` is the generalization that removes the coupling.
- **External controller copies the address into `target`.** No ExternalDNS change, but pushes an
  extra controller onto every user; maintainers floated the generic annotation as the better
  answer.
- **Explicit-version grammar** (`<resource>.<group>/<version>/<ns>/<name>#<path>`, the #6190
  form). More precise but reintroduces version-drift fragility; offered as an optional override
  on the discovery-based default, not the default.
- **Do nothing.** Each new "point at product X" request keeps arriving as a coupling PR that gets
  held or closed, leaving users on manual `target` or the heavier `unstructured` source.
