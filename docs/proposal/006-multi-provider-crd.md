```yaml
---
title: "Multiple Providers and Zones from a Single Deployment via CRDs"
version: v1alpha1
authors: "@mloiseleur"
creation-date: 2026-06-21
status: provisional
---
```

# Multiple Providers and Zones from a Single Deployment via CRDs

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
    - [Pipeline Lifecycle](#pipeline-lifecycle)
    - [Routing Endpoints to a Provider](#routing-endpoints-to-a-provider)
    - [Endpoint Provenance and Merge Ordering](#endpoint-provenance-and-merge-ordering)
    - [Credentials](#credentials)
    - [Ownership Isolation](#ownership-isolation)
    - [Conflict Resolution](#conflict-resolution)
    - [Status and Observability](#status-and-observability)
    - [Edge Cases](#edge-cases)
  - [Drawbacks](#drawbacks)
- [Implementation Plan](#implementation-plan)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

ExternalDNS currently binds **one process to one provider+zone configuration**, supplied as
CLI flags / env vars and resolved **once** at startup
([`controller/execute.go`](https://github.com/kubernetes-sigs/external-dns/blob/master/controller/execute.go)).
Managing several providers, accounts, or zones requires one deployment per configuration.

This proposal moves provider+zone configuration into two new CRDs — `DNSProvider` (namespaced)
and `ClusterDNSProvider` (cluster-scoped) — modeled on cert-manager's `Issuer` / `ClusterIssuer`.

A single ExternalDNS deployment watches these resources and runs one independent reconcile
pipeline per resource, created and torn down dynamically as the CRDs change.

This **replaces** the provider-selection flags (`--provider`, `--domain-filter`, `--txt-owner-id`,
registry flags, credential env, …) rather than living beside them: the CRD becomes the single config
mechanism for providers and zones. Migration is a mechanical near-1:1 flag→field mapping, but a
**breaking change** gated on the next major version. Global flags (`--policy`, `--interval`, source
selection, logging, metrics, leader election) stay on the CLI.

This addresses [#1961](https://github.com/kubernetes-sigs/external-dns/issues/1961)

## Motivation

[#1961](https://github.com/kubernetes-sigs/external-dns/issues/1961) is one of the most requested
features in the project. The recurring pain across the thread:

- **Multiple accounts / subscriptions** — separate AWS accounts or Azure subscriptions, each needing
  its own role/credentials, force one deployment each.
- **Split-horizon DNS** — the same records must be published to an internal zone (e.g. `rfc2136` /
  Active Directory) and a public zone (e.g. Route53).
- **Multi-tenancy** — in shared clusters, each tenant runs its own deployment; there is no clean way
  to say "only namespace X may publish to zone Y".
- **Per-zone credentials** — operators want to revoke a single zone's key without touching others.
- **Operational overhead** — N deployments mean N sets of RBAC, metrics, dashboards, upgrades.

Today's only mitigations — `--domain-filter` repetition or `--regex-domain-filter` — share a
**single** provider and credential set, so they do not cover multi-account, multi-provider, or
multi-tenant cases.

### Goals

- Run **multiple provider+zone configurations from one deployment**, configured via CRDs.
- Configuration changes (add / update / remove a provider) take effect **without restarting** the
  deployment.
- Support **multi-tenant isolation**: namespaced `DNSProvider` for tenant self-service, cluster-scoped
  `ClusterDNSProvider` for platform-wide zones.
- Support **split-horizon**: a single source record may be published by more than one provider.
- **Reuse the existing provider/registry factories** rather than reimplementing provider wiring.
- Make the CRD the **single** provider/zone configuration mechanism, with a mechanical, near-1:1
  migration from the existing provider-selection flags.

### Non-Goals

- **Not** adding new in-tree providers. This is an orchestration layer over existing providers and the
  webhook mechanism; it does not alter the "no new in-tree providers" gate from
  [#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347).
- **Not** a fully-typed per-provider CRD schema for all 20+ providers (see
  [Alternatives](#alternatives)).
- **Not** changing the global controller flags (`--policy`, `--interval`, source selection, logging,
  metrics, leader election); only the provider-selection flags move to the CRD.
- **Not** redesigning the `Source` interface or the `plan` diffing logic.
- **Not** solving cross-deployment leader election beyond what
  [proposal 001](./001-leader-election.md) already covers.

## Proposal

A single ExternalDNS controller becomes a small **operator** that, instead of building one pipeline at
startup, watches `DNSProvider` / `ClusterDNSProvider` resources and maintains a **set of pipelines** —
one per resource. Each pipeline is the existing
`Source → Provider → Registry → Plan → ApplyChanges` flow, parameterized by a per-resource
configuration synthesized from the CRD spec.

```text
                         ┌──────────────────────────────────────────────┐
                         │ ExternalDNS deployment                         │
                         │                                                │
  DNSProvider ───────────►  watch ──► pipeline manager                    │
  ClusterDNSProvider ────►          │   ├─ pipeline A (route53, zoneA)     │
                         │          │   ├─ pipeline B (cloudflare, zoneB)  │
  shared informers ──────►  Source ─┤   └─ pipeline C (rfc2136, internal)  │
  (svc/ingress/routes)   │  store   │        each: own provider+registry   │
                         │          │        own txt-owner-id, own loop    │
                         └──────────────────────────────────────────────┘
```

The Source store (Kubernetes informers for Services / Ingresses / Gateway routes / CRDs) is **shared**:
informers are created once, and each pipeline filters the shared endpoint set down to the records it
owns (see [Routing](#routing-endpoints-to-a-provider)).

### User Stories

#### Story 1: Multiple AWS accounts (#1961, nikolaiderzhak)

*As a platform engineer*, I run workloads spanning three AWS accounts. I create three
`ClusterDNSProvider` resources, each with a Route53 config and a `credentialsRef`/IRSA role for its account.
One ExternalDNS deployment publishes records into the right account based on the zone each hostname
falls under — no more three deployments to upgrade in lockstep.

#### Story 2: Split-horizon (#1961, anthonysomerset, Px-x64)

*As an operator*, `app.example.com` must resolve internally via `rfc2136` (Active Directory) and
externally via Route53. I create two `ClusterDNSProvider` resources covering the same domain and add a
list-valued reference annotation on the Ingress so **both** publish the record.

#### Story 3: Multi-tenant self-service (#1961, sagikazarmark)

*As a cluster admin*, I let team `foo` manage DNS for `foo.example.com` only. I create a namespaced
`DNSProvider` in namespace `foo` scoped to that zone. Resources in `foo` reference it; resources in
other namespaces cannot use it. RBAC on the `DNSProvider` kind controls who may publish where.

#### Story 4: Per-zone credential rotation (#1961, rumstead)

*As an operator*, one zone's API key is leaking quota. I rotate or revoke only that
`DNSProvider`'s referenced Secret; other zones keep running untouched.

### API

Two CRDs in the existing API group `externaldns.k8s.io/v1alpha1`. They are **identical in spec**,
differing only in scope (`Namespaced` vs `Cluster`), mirroring cert-manager's Issuer/ClusterIssuer.

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: ClusterDNSProvider
metadata:
  name: route53-prod
spec:
  # Provider name — resolved by the existing provider factory.
  type: aws
  # Optional: zones this provider is authoritative for. Used for implicit
  # domain-match routing and to scope the provider's Records() calls.
  domainFilter:
    include:
      - prod.example.com
    exclude:
      - internal.prod.example.com
  # Registry / ownership. Defaults: registry "txt", txtOwnerId = metadata.name.
  registry:
    type: txt
    txtOwnerId: route53-prod      # MUST be unique per provider; defaults to name
    txtPrefix: "edns-"
  # Per-provider configuration. Typed common fields above; provider-specific
  # settings passed through as a flag-equivalent map (thin wrapper, see Behavior).
  config:
    aws-zone-type: public
    aws-assume-role: "arn:aws:iam::111122223333:role/external-dns"
  # Credentials, kept distinct in-memory per provider (per-pipeline, not global env).
  credentialsRef:
    name: route53-prod-credentials   # Secret; keys read by this pipeline only
status:
  conditions:
    - type: Ready
      status: "True"
      reason: ProviderInitialized
  observedGeneration: 3
  lastSyncTime: "2026-06-21T10:00:00Z"
  records: 42
```

The namespaced form is identical:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSProvider
metadata:
  name: team-foo-zone
  namespace: foo
spec:
  type: cloudflare
  domainFilter:
    include:
      - foo.example.com
  credentialsRef:
    name: cloudflare-foo-token   # Secret in namespace foo
```

The CRD is the **only** way to configure providers and zones; there is no mode flag. At least one
`DNSProvider` / `ClusterDNSProvider` must exist for the deployment to publish records. The
provider-selection flags (`--provider`, `--domain-filter`, `--txt-owner-id`, `--txt-prefix`, registry
flags, credential env, …) are removed; only genuinely global flags remain (`--policy` default,
`--interval` default, `--log-level`, `--metrics-address`, `--events`, leader election, source selection).

#### Migration

Each removed flag maps near-1:1 to a CRD field; an existing deployment becomes one `ClusterDNSProvider`
plus the unchanged global flags. Breaking, gated on the next major version. An external
`scripts/migrate-flags-to-crd.py` helper renders the CRD (+ Secret stub) from an existing flag set so
migration is `run script → kubectl apply`.

| Removed flag | CRD field |
|---|---|
| `--provider=aws` | `spec.type: aws` |
| `--domain-filter=example.com` | `spec.domainFilter.include: [example.com]` |
| `--exclude-domains=internal.example.com` | `spec.domainFilter.exclude: [internal.example.com]` |
| `--registry=txt` | `spec.registry.type: txt` |
| `--txt-owner-id=me` | `spec.registry.txtOwnerId: me` |
| `--txt-prefix=edns-` | `spec.registry.txtPrefix: edns-` |
| provider-specific flags (`--aws-zone-type`, …) | `spec.config` map entries |
| credential env (`AWS_*`, API tokens) | `spec.credentialsRef` → Secret |

### Behavior

#### Pipeline Lifecycle

A **pipeline manager** watches both CRD kinds:

- **Add** — synthesize a per-provider `externaldns.Config` plus a credential source from the spec
  (see [Credentials](#credentials)), build a `DomainFilter`, then call the **existing**
  [`providerfactory.Select`](https://github.com/kubernetes-sigs/external-dns/blob/master/controller/execute.go#L110)
  and
  [`registryfactory.Select`](https://github.com/kubernetes-sigs/external-dns/blob/master/controller/execute.go#L159),
  and construct a `Controller` bound to the shared source. Start its reconcile loop.
- **Update** (`spec` changed, `observedGeneration < generation`) — rebuild the pipeline; swap atomically.
- **Delete** — stop the loop, drain in-flight work. **Records are not deleted** on CRD removal; ownership
  cleanup remains an explicit operator action (consistent with how stopping a deployment behaves today).

Each pipeline runs its own `RunOnce` on the global `--interval`, and is additionally triggered by shared
source events (`--events`), throttled by `--min-event-sync-interval` exactly as the single pipeline is today.

#### Routing Endpoints to a Provider

The shared source produces the full desired endpoint set once per sync. Each pipeline selects the subset
it should manage using a **two-tier rule** (explicit reference wins, domain match is the fallback):

1. **Explicit reference** — a source resource may name one or more providers via annotation:

   ```yaml
   metadata:
     annotations:
       # namespaced DNSProvider in the resource's namespace
       external-dns.kubernetes.io/provider: "team-foo-zone"
       # cluster-scoped, comma-separated for split-horizon
       external-dns.kubernetes.io/cluster-provider: "route53-prod,rfc2136-internal"
   ```

   When present, **only** the named providers consider the record. A list publishes the **same** record
   through several providers (Story 2). A namespaced resource may only reference `DNSProvider`s in its own
   namespace; it may reference any `ClusterDNSProvider`.

2. **Domain match (fallback)** — when no reference annotation is present, the record is routed to every
   provider whose `spec.domainFilter` matches the record's DNS name. If exactly one matches, it is
   selected; if several match, see [Conflict Resolution](#conflict-resolution).

This keeps **zero-annotation flows working** (records land on the provider that owns their zone) while
giving explicit control where needed.

#### Endpoint Provenance and Merge Ordering

Explicit-reference routing needs each endpoint's source object. Two existing mechanics must be pinned down:

- **Provenance via `RefObjects`.** The reference annotation is read from
  [`Endpoint.RefObjects()`](https://github.com/kubernetes-sigs/external-dns/blob/master/endpoint/endpoint.go),
  not the DNS name — so every routable source must populate `RefObjects` (the gap
  [#6492](https://github.com/kubernetes-sigs/external-dns/pull/6492) closed for ~17 sources). Sources with
  no Kubernetes object (`connector`, static) get **domain-match routing only**.
- **Route before `MergeEndpoints`.**
  [`MergeEndpoints`](https://github.com/kubernetes-sigs/external-dns/blob/master/endpoint/utils.go)
  collapses same-name records into one endpoint with the **union** of their `RefObjects`. Routing runs on
  the **pre-merge** set (each endpoint = one source object, one annotation set); merge then runs within
  each pipeline's subset. This also keeps namespaced-`DNSProvider` enforcement sound.
- **Conflicting references on a merged name** are resolved as in
  [Conflict Resolution](#conflict-resolution): **skip** the contested providers with a `Warning` event +
  metric.

#### Credentials

Each provider's `credentialsRef` points to a Secret. Many providers read credentials from process
**environment** today, so running N providers in one process needs an explicit **per-pipeline credential
boundary** — never a mutated shared `os.Environ()`:

- A per-pipeline key/value source (from `credentialsRef` + `spec.config`) is passed into provider
  construction at the factory boundary; providers read credentials from it, **not** global env.
- Ambient identity (IRSA, Workload Identity) still works when `credentialsRef` is omitted and the
  role/identity is set via `spec.config`.

This is the main factory-layer change; per-provider logic is reused unchanged. **Phase 2 acceptance**:
two same-type pipelines with distinct credentials run concurrently without interference, no global-env
mutation.

#### Ownership Isolation

Each provider gets its **own** `txtOwnerId`, defaulting to `metadata.name`. This keeps TXT registry
markers disjoint across pipelines, so two providers managing overlapping zones do not fight over each
other's ownership records. The validating webhook (below) rejects duplicate `txtOwnerId` values within
the cluster to prevent silent record takeover.

#### Conflict Resolution

When two providers match the same record by **domain** (no explicit reference), this is ambiguous. Rules:

- **Disjoint `domainFilter`s** are the recommended configuration; the validating webhook **warns** on
  overlapping include filters across providers of the same scope.
- For genuine overlap (e.g. split-horizon by design), users **must** use the explicit reference
  annotation; implicit domain match selects **at most one** provider.
- If overlap remains unresolved, the record is **skipped** for the ambiguous providers and a `Warning`
  event + metric is emitted, rather than published twice non-deterministically.

#### Status and Observability

- Each CRD reports `status.conditions` (`Ready`, with reasons such as `ProviderInitialized`,
  `AuthFailed`, `ZoneNotFound`), `observedGeneration`, `lastSyncTime`, and a `records` count.
- Metrics gain a `provider_resource` label (the CRD name) so existing per-sync metrics are reported
  per pipeline. Cardinality is bounded by the number of CRDs.
- Source-resource events reference the provider that published (or skipped) them.

#### Edge Cases

- **No matching provider** — record is unmanaged; surfaced via metric + optional source event.
- **CRD references a missing Secret / zone** — pipeline enters `Ready=False`; other pipelines unaffected.
- **Leader election** — with multiple replicas, [proposal 001](./001-leader-election.md) governs which
  replica is active; pipelines run on the leader. No new cross-pipeline election is introduced.
- **`--once` / dry-run** — apply per pipeline; `--dry-run` is honored globally.
- **Same record, multiple providers** — supported only via explicit list reference, and only with
  distinct `txtOwnerId`s.

### Drawbacks

- **New controller surface** — CRD watching, a pipeline manager with dynamic lifecycle, a validating
  webhook, and status reporting are substantial new code and test surface.
- **Resource usage** — N pipelines means N provider clients and N reconcile loops in one process;
  memory/CPU grow with provider count (still typically cheaper than N deployments).
- **Per-pipeline credential handling at the factory boundary** touches a sensitive, well-exercised code path.
- **Thin-wrapper config is less type-safe** than per-provider schemas: provider-specific keys are a
  passthrough map, so misconfiguration surfaces at pipeline init, not at `kubectl apply` (mitigated by
  status conditions and the validating webhook).
- **Breaking change.** Every deployment — including the common single-provider case — must install the
  CRDs and apply one `ClusterDNSProvider` before publishing. Migration is mechanical but not zero-effort;
  an external `scripts/migrate-flags-to-crd.py` helper renders the CRD (+ Secret stub) from an existing
  flag set (cf. `scripts/aws-cleanup-legacy-txt-records.py`).

## Implementation Plan

Phased. Phases 1–5 land additively behind an alpha CRD path so the work is reviewable incrementally;
phase 6 performs the breaking flag removal at the major-version boundary:

1. **CRD types + scaffolding** — `DNSProvider` / `ClusterDNSProvider` types in `apis/v1alpha1`, `make crd`,
   RBAC, deepcopy. No behavior yet.
2. **Per-pipeline credential/config boundary** — add a per-pipeline credential source at the provider
   factory boundary; synthesize the source from `spec`. *Acceptance*: two same-type pipelines with
   distinct credentials run concurrently; no global-env mutation. Unit-tested without K8s.
3. **Pipeline manager** — watch CRDs, build/teardown pipelines via existing factories, shared source store.
   Domain-match routing only.
4. **Explicit reference annotations + split-horizon** — pre-merge two-tier routing, list references.
   *Acceptance*: routing reads provenance from `RefObjects` before `MergeEndpoints`; conflicting refs on a
   merged name are skipped + warned; namespaced references enforced against the referencing object.
5. **Status, metrics labels, validating webhook** — conditions, duplicate-`txtOwnerId` rejection,
   overlap warnings.
6. **Flag removal + migration** — remove provider-selection flags; ship the flag→field migration guide
   and an external `scripts/migrate-flags-to-crd.py` helper that renders a `ClusterDNSProvider` (+ Secret
   stub) from existing args. Breaking; major version.

## Alternatives

### Alternative 1: Status quo — one deployment per provider

**Description**: Keep the current model; operators deploy ExternalDNS once per provider+zone configuration.

**Pros**:

- Zero new code
- Strong isolation between configurations (separate processes, RBAC, blast radius)
- Simple mental model

**Cons**:

- The exact overhead #1961 is about — N× RBAC, metrics, dashboards, upgrades
- No clean multi-tenant story (no namespace-scoped authorization)

**Recommendation**: ❌ Not recommended — this is the problem being solved.

### Alternative 2: Fully-typed per-provider CRD schema

**Description**: A strongly-typed spec per provider, validated at `kubectl apply`, instead of the
thin-wrapper `config` passthrough map.

**Pros**:

- Best UX and validation (errors surface at apply time, not pipeline init)
- Self-documenting schema per provider

**Cons**:

- Enormous surface across 20+ providers
- Directly conflicts with the out-of-tree direction of
  [#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347)
- Unmaintainable as providers move out of tree

**Recommendation**: ❌ Not recommended as the baseline — the thin-wrapper `config` map can be tightened
to typed fields per provider later if desired.

### Alternative 3: `--domain-filter` / `--regex-domain-filter` only

**Description**: Rely on the existing domain-filter flags to cover multiple zones from one deployment.

**Pros**:

- Exists today, no new code
- Familiar to current users

**Cons**:

- Shares a single provider and a single credential set
- Cannot serve multi-account, multi-provider, or multi-tenant cases

**Recommendation**: 🟡 Complementary, not a substitute — useful within a single provider, orthogonal to
this proposal.

### Alternative 4: Cluster-scoped CRD only (no namespaced kind)

**Description**: Ship only `ClusterDNSProvider`; rely on RBAC over the cluster resource for tenancy.

**Pros**:

- Smaller API surface
- Fewer resources to reason about

**Cons**:

- No tenant self-service
- Multi-tenancy depends entirely on RBAC over a cluster resource
- The issue explicitly asks for namespace-scoped authorization

**Recommendation**: ❌ Not recommended — namespaced + cluster-scoped chosen to match the cert-manager
model the issue references. Cluster scope alone cannot express per-namespace ownership.

### Alternative 5: Additive opt-in mode, keep the flags

**Description**: Keep the provider-selection flags working and add the CRD path behind a mode flag
(e.g. `--provider=crd`), so existing single-provider deployments are unaffected.

**Pros**:

- Non-breaking; gradual adoption
- Existing deployments need no migration

**Cons**:

- Two configuration systems coexist permanently — more code, docs, and support surface
- A mode flag overloaded onto `--provider` collides with existing CRD *source* and CRD *registry*
  terminology
- The dual path becomes the long-term shape, not a transition

**Recommendation**: ❌ Not recommended — the mechanical near-1:1 migration is outweighed by collapsing to
a single configuration mechanism. Retained as the fallback if the breaking change proves too disruptive.
