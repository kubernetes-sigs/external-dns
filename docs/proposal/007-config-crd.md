```yaml
---
title: "Configuration as a CRD"
version: v1alpha1
authors: "@mloiseleur"
creation-date: 2026-09-26
status: draft
---
```

# Configuration as a CRD

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
    - [Loading](#loading)
    - [Schema Generation](#schema-generation)
    - [Flags and Secrets](#flags-and-secrets)
    - [Configuration Changes](#configuration-changes)
    - [Status and RBAC](#status-and-rbac)
    - [Edge Cases](#edge-cases)
  - [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

Let an ExternalDNS deployment read its configuration from a namespaced `ExternalDNSConfig` resource
instead of CLI flags, gaining apply-time validation and a health `status`. Opt-in and non-breaking:
without `--config-ref`, nothing changes.

The process model is unchanged: **one deployment, one configuration**. This keeps the isolation that
led to rejecting [proposal 006](https://github.com/kubernetes-sigs/external-dns/pull/6512) and only
changes where a deployment's configuration lives.

## Motivation

Configuration is a list of CLI args (or `EXTERNAL_DNS_*` env vars) today:

- **No apply-time validation.** A wrong value or unknown flag surfaces as a `CrashLoopBackOff`.
- **No status.** Configuration, provider auth and sync health are only visible in logs and metrics.
  With one deployment per tenant, there is no `kubectl get` view of which ones are broken.
- **Not machine-readable.** Helm, ArgoCD, operators and policy engines work on an args array, not a
  typed object.
- **Flags do not scale.** Per-source settings
  ([#5397](https://github.com/kubernetes-sigs/external-dns/issues/5397)) would multiply ~180 flags by the
  number of sources; the answer there was "a config file or a CR".

### Goals

- Configure one deployment from one `ExternalDNSConfig`.
- Validate types and enum values at `kubectl apply`.
- Report configuration and sync health in `status.conditions`.
- Generate the CRD schema, CLI flags and `docs/flags.md` from the **same** flag definitions.
- Keep flags working unchanged.

### Non-Goals

- Several configurations per process ([proposal 006](https://github.com/kubernetes-sigs/external-dns/pull/6512)).
- Credential values in the resource; it only references `Secret`s.
- Deprecating flags.
- Per-source overrides ([#5397](https://github.com/kubernetes-sigs/external-dns/issues/5397)). The CRD
  gives them a place later (e.g. `spec.source.service.labelFilter`) without new flags.
- Managing the ExternalDNS `Deployment` (an operator's job).

## Proposal

### User Stories

- **Fail at apply** — `policy: upsert-onyl` is rejected by the API server, with the allowed values,
  before any pod restarts.
- **Fleet health** — a platform team running one ExternalDNS per tenant runs
  `kubectl get externaldnsconfigs -A` and sees which are `Ready=False`, and why.
- **Generated configuration** — tools generate and diff typed objects; policy engines (Kyverno,
  Gatekeeper) enforce rules such as "every config sets `txtOwnerId`".

### API

A namespaced kind in the existing group `externaldns.k8s.io/v1alpha1`. Each flag maps to exactly one
`spec` path, grouped by concern (`--aws-zone-type` → `provider.aws.zoneType`, `--txt-owner-id` →
`registry.txt.ownerId`); global flags stay at the top level:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: ExternalDNSConfig
metadata:
  name: external-dns
  namespace: external-dns
spec:
  provider:
    name: aws
    domainFilter: [example.com]
    aws:
      zoneType: public
  source:
    kinds: [service, ingress]
  registry:
    type: txt
    txt:
      ownerId: my-cluster
  policy: upsert-only
  interval: 1m
status:
  observedGeneration: 3
  lastSuccessfulSyncTime: "2026-09-26T10:00:00Z"
  conditions:
    - type: Ready
      status: "True"
      reason: Synced
      observedGeneration: 3
```

The deployment runs `external-dns --config-ref=external-dns`. The resource is looked up in the pod's
namespace, resolved by the existing `client.CurrentNamespace`.

### Behavior

#### Loading

With `--config-ref` set, at startup:

1. Parse bootstrap flags, build the Kubernetes client, `GET` the resource.
2. Apply `spec` onto the default `Config` through `bindFlags` (see [Schema Generation](#schema-generation)).
3. Run `validation.ValidateConfig`, then start as today.

A missing resource or a failed validation exits non-zero, like an invalid flag today, and sets
`Ready=False` when the resource exists.

#### Schema Generation

179 of the 181 flags are declared in `bindFlags` (`pkg/apis/externaldns/types.go`) through the
`flags.FlagBinder` interface (`internal/flags/binders.go`): name, type, default, help text, allowed enum
values. `docs/flags.md` is already generated from it.

Prerequisites:

- Move `--provider` and `--source` into `bindFlags`; they are declared on kingpin in `App()` to keep
  `Required()`.
- Add a `required` marker to `FlagBinder`, and expose the existing `secure:"yes"` tag on `Config`
  fields (already used to mask the startup log) as a `sensitive` marker.
- Declare a **group** for every flag in `bindFlags` (e.g. `provider.aws`, `registry.txt`, `source`).
  Explicit, because prefixes are ambiguous (`webhook-*` is both a provider and `--webhook-server`).

Path rule: group + flag name without the group's prefix, in camelCase (`--aws-zone-type` in
`provider.aws` → `provider.aws.zoneType`). A flag named like its group gets an explicit leaf
(`--provider` → `provider.name`, `--source` → `source.kinds`, `--registry` → `registry.type`).
One flag, one path, both ways; the generator fails on collisions. `docs/flags.md` gains a path column;
env var names are unchanged.

Two new binders then reuse `bindFlags`:

- **Schema binder** — emits one OpenAPI property per flag at its path (type, `enum`, `required`,
  `description`). Run by `make crd`; CI fails when the CRD is stale.
- **Spec binder** — sets each `Config` field from its `spec` path.

#### Flags and Secrets

With `--config-ref` set, configuration comes **only** from the resource; other flags are rejected at
startup, except **bootstrap flags** needed to read it: `--config-ref`, `--kubeconfig`, `--server`,
Kubernetes client timeouts, `--log-level`, `--log-format`. They are excluded from the schema.

**Sensitive flags**, the `Config` fields tagged `secure:"yes"` (API keys, passwords, TSIG secrets), get a
`<field>SecretRef` instead of a value field:

```yaml
spec:
  provider:
    name: pdns
    pdns:
      server: https://pdns.example.com
      apiKeySecretRef:
        name: pdns-credentials
        key: api-key
```

- **Less exposure** — flags show in the pod spec and `ps`, env vars in `/proc/<pid>/environ` and
  monitoring agents; a ref keeps the value in process memory.
- **Rotation** — a change to a referenced `Secret` restarts ExternalDNS, like a spec change.
- **Same namespace only**, and RBAC scoped with `resourceNames`: otherwise whoever edits the resource
  could point a ref at any `Secret` and send it to a server they also set in `spec`.
- **Env vars still work**; a ref wins when both are set.

Provider SDK credentials (`AWS_*`, `AZURE_*`, `GOOGLE_APPLICATION_CREDENTIALS`, …) are unchanged: env
vars, mounted files, or workload identity.

#### Configuration Changes

ExternalDNS watches its resource. On a `metadata.generation` change it exits cleanly and Kubernetes
restarts it with the new spec: one code path for startup and reconfiguration, and the restart count
records every reload. Cross-field rules (`ValidateConfig`) are checked at that restart.

#### Status and RBAC

`Ready` condition, written by the running ExternalDNS, with reason `InvalidConfig`, `ProviderInitFailed`,
`SyncFailed` or `Synced`; plus `observedGeneration` and `lastSuccessfulSyncTime`. Written on condition
transitions only, `lastSuccessfulSyncTime` at most every 10 minutes, so a large fleet on a 1-minute
interval does not load the API server. Metrics stay the source for per-sync detail.

RBAC: a namespaced `Role` with `get`, `watch` on `externaldnsconfigs`, `patch` on
`externaldnsconfigs/status`, and `get`, `watch` on the referenced `secrets` (`resourceNames`). No new
cluster-scoped permissions.

#### Edge Cases

- **Resource deleted while running** — exit; the pod fails to start until it is recreated. Records are
  untouched.
- **One resource, several deployments** — allowed, but they share a `txtOwnerId`; not recommended.
- **Deprecated flags** — stay in the schema while the flag exists; `resolveDeprecatedFlags` runs after
  loading.

### Drawbacks

- **CRD lifecycle.** Helm installs CRDs on first install only; schema changes need a manual
  `kubectl apply`, as for `DNSEndpoint` today.
- **Group every flag.** A one-time pass over ~180 flags in `bindFlags`, and a group to pick for each new
  flag.
- **Restart on change** costs a container restart and a full initial sync.
- **Two ways to configure**, both to document and support.
- **Secret read access** that env vars do not need, traded against their exposure.

## Alternatives

| Alternative | Why not |
|---|---|
| **Config file** (`--config=path`, from a ConfigMap) | 🟡 Complementary, the spec binder can load it. No apply-time validation, no status. |
| **Hot reload** without restart | ❌ Not now: a second lifecycle for informers and provider clients, partial config on failure. Can be added later without API change. |
| **Flat schema**, one top-level field per flag | ❌ No group to design, but ~170 fields in one `spec` and no place for per-source overrides. |
| **Hand-written schema** with CEL | ❌ Drifts from ~180 flags; conflicts with moving providers out of tree ([#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347)). CEL can be added to the generated schema later. |
| **Status quo**, validation in the Helm values schema | ❌ Helm-only, only for fields the chart models, no status. |
