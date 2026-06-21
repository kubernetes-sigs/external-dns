```yaml
---
title: "libdns Provider Adapter for Simple-Zone Providers"
version: v1alpha1
authors: "@mloiseleur"
creation-date: 2026-06-21
status: draft
---
```

# libdns Provider Adapter for Simple-Zone Providers

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
    - [Record Mapping](#record-mapping)
    - [Applying Changes](#applying-changes)
    - [Zone Discovery](#zone-discovery)
    - [Provider Selection](#provider-selection)
    - [Unsupported Endpoint Features](#unsupported-endpoint-features)
  - [Distribution Policy (Open Question)](#distribution-policy-open-question)
  - [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

Add one generic in-tree provider that adapts [libdns](https://github.com/libdns) modules to the
ExternalDNS `Provider` interface. This replaces several bespoke in-tree providers with thin, build-tag
shims over maintained libdns modules — cutting per-provider code and dependency churn — while keeping the
"no new in-tree providers" gate from
[#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347) intact.

It is **not** a reversal of #4347, but a way to make its end state cheaper to maintain and less disruptive
for users of the simpler providers.

## Motivation

Issue #4347 moves providers out of tree, with the webhook mechanism as the replacement. Two frictions recur:

- **User experience.** A webhook provider adds a second container and an unaffiliated third-party image
  into a security-sensitive control loop. See
  [#6491](https://github.com/kubernetes-sigs/external-dns/pull/6491#issuecomment-4706415302): a TransIP
  user on ~10 clusters pushed back on losing in-tree support with only a third-party webhook as the
  alternative.
- **Maintainer cost.** Providers leave because of bespoke code plus frequent SDK dependency bumps — the
  cost #4347 calls out.

libdns is a small, stable interface set (`RecordGetter`, `RecordAppender`, `RecordSetter`,
`RecordDeleter`, `ZoneLister`) with 80+ maintained modules. One adapter serves many providers, and new
DNS vendors integrate by publishing their own libdns module — so the gate stays closed. As of 2026-06,
modules already exist for `transip`, `scaleway`, `linode`, `dnsimple`, `gandi`, `godaddy`, `civo`,
`exoscale`, and `ovh` (no module for `ns1`).

### Goals

- One generic in-tree adapter from libdns modules to the `Provider` interface.
- A libdns-backed replacement path for simple-zone providers before/with their removal.
- Default binary free of any libdns dependency (build-tag gated).
- Keep the #4347 gate: new providers arrive as libdns modules, not in-tree code.

### Non-Goals

- In-tree status for AWS, Azure, Google, or Cloudflare routing — they need the rich endpoint model
  (set identifiers, weighted/latency/geo routing, provider-specific fields) libdns does not model.
- Auto-importing all 80+ libdns modules into the default binary.
- Replacing the webhook mechanism; the two coexist.

## Proposal

A generic provider at `provider/libdns/` implementing the `Provider` interface on top of the libdns
interfaces. It imports no vendor SDK directly — only libdns modules, behind a build tag.

```text
provider/libdns/
  libdns.go     // generic Provider impl (Records, ApplyChanges, AdjustEndpoints)
  registry.go   // name -> factory for the curated set; gated by //go:build libdns
```

### User Stories

- **TransIP user ([#6491](https://github.com/kubernetes-sigs/external-dns/pull/6491)).** After in-tree
  TransIP is removed, an operator keeps a first-party, single-pod setup via the libdns adapter over
  `libdns/transip` — no second container, no third-party image.
- **New DNS vendor.** A vendor publishes a libdns module and is usable through the adapter without adding
  code to this repo or coupling to its release cycle.
- **Maintainer.** Retire a bespoke provider and its SDK by replacing it with a libdns shim behind the
  build tag; the SDK leaves the default binary.

### API

Selection mirrors `--provider`, with a sub-selector and a single JSON config blob (the pattern Caddy uses
for libdns modules), unmarshalled into the concrete provider struct by its shim:

```text
--provider=libdns
--libdns-provider=transip
EXTERNAL_DNS_LIBDNS_CONFIG={ "api_key": "..." }
```

No CRD changes; no changes to the `Source`, `Plan`, or `Registry` layers.

### Behavior

The adapter implements `Records` and `ApplyChanges`. The TXT registry and ownership model are unchanged —
the registry sits above the provider, so TXT markers flow through as ordinary TXT endpoints.

#### Record Mapping

Every libdns record type exposes `RR()` → flat `RR{Name, Type string, TTL time.Duration, Data string}`.
The adapter works only in `RR` terms both directions; the module parses `RR` on write. No per-record-type
switch.

| `endpoint.Endpoint` | libdns | Notes |
|---|---|---|
| `DNSName` (FQDN) | zone + relative `Name` | Names are zone-relative; use `RelativeName`/`AbsoluteName`. Needs zone discovery (below). |
| `Targets` (N) | N `RR` records | Grouped back by `(name, type)` on read. |
| `RecordType` | `RR.Type` | Direct string. |
| `RecordTTL` (seconds) | `time.Duration` | Convert on each boundary. |
| MX/SRV target (`"10 host"`) | `RR.Data` (zone-file value) | Already ExternalDNS's storage form; passes through. |

#### Applying Changes

Group `plan.Changes` by zone, then by `(name, type)` RRset:

- `Create` + `UpdateNew` → `SetRecords` (its "these are the only records for this `(name, type)`"
  semantics match a desired RRset exactly).
- `Delete` → `DeleteRecords`.
- No `RecordSetter`? Emulate via `DeleteRecords` + `AppendRecords`.

#### Zone Discovery

libdns needs the zone as a separate argument, while ExternalDNS hands providers FQDNs. `--domain-filter`
is the primary zone source (works for every module, commonly set already); FQDNs match by longest suffix.
`ZoneLister` is an optional convenience: when implemented, the adapter can auto-discover zones and
`--domain-filter` becomes optional; otherwise `--domain-filter` is required. As of 2026-06, among in-scope
modules only `transip` and `linode` implement `ZoneLister`.

#### Provider Selection

The supported modules are one curated set fixed at build time (no per-provider compile selection), gated
behind a single `libdns` tag:

```go
//go:build libdns

func init() {
    register("transip", func(cfg json.RawMessage) (libdnsClient, error) {
        p := &transip.Provider{}
        return p, json.Unmarshal(cfg, p)
    })
    // ... rest of the curated set
}
```

Without the tag, no libdns module is imported. With `-tags libdns`, the active module is chosen at runtime
via `--libdns-provider`.

#### Unsupported Endpoint Features

`SetIdentifier` drives provider-native routing — multiple records per `(name, type)`, which flat libdns
backends cannot store. It cannot be silently dropped: the plan keys on `(dnsName, setIdentifier)`, but a
flat backend reads back an empty identifier, so desired (non-empty) and current (empty) never match and
the record churns every reconcile under `--update-events`.

As a consequence, the adapter strips it (and warns) in `AdjustEndpoints`.

### Distribution Policy (Open Question)

The `libdns` tag gates the curated set as one unit. The decision for the SIG is whether the official
distribution ships it as a separate image or folds it into the main one — policy, not code.

- **Option A — two images.** Main image unchanged; a separate `external-dns:<ver>-libdns` variant is built
  with `-tags libdns`. Lean default, dependency churn confined to the variant. Cost: an extra image to
  build, scan, release.
- **Option B — single unified image.** Main image built with `-tags libdns` for everyone. Simplest to
  operate, best UX. Cost: the libdns SDKs and their churn land in every user's image.

### Drawbacks

- **Distribution overhead** — Option A adds an image variant; Option B adds the libdns dependency surface
  to every image.
- **Provider quirks still leak** — libdns abstracts the API call, not provider behavior. Quirks like the
  trailing-dot bug in #6491 surface in the module and are fixed upstream, not here.
- **Reduced feature surface** — no routing policies; `SetIdentifier` is stripped in `AdjustEndpoints`,
- **Two integration paths** — adapter vs. webhook; docs must steer users.
- **Coverage gaps** — providers without a module (e.g. `ns1`) are not served.

## Alternatives

### Alternative 1: First-party generic libdns webhook image

- **Pros:** No in-tree libdns dependency; smallest footprint; fully decoupled release cycle.
- **Cons:** Keeps the two-container UX #6491 objected to.
- **Recommendation:** Possible *in addition* to the adapter, not instead.

### Alternative 2: Status quo (webhook only)

Relying on the webhook, including the third-party `orbit-online/external-dns-libdns-webhook`.

- **Pros:** No work; functional today.
- **Cons:** The third-party, low-activity image is the exact concern raised in #6491.
- **Recommendation:** Not recommended as the only option for the simple-zone tier.

### Alternative 3: Keep bespoke in-tree providers

- **Pros:** Best UX; no new abstraction.
- **Cons:** Reintroduces the maintenance and Dependabot burden #4347 removes; does not scale.
- **Recommendation:** Not recommended; contradicts #4347.
