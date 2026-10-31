```yaml
---
title: "Webhook source"
version: if applicable
authors: "@ivankatliarchuk"
creation-date: 2026-10-25
status: draft
---
```

# Webhook Source

## Table of Contents

<!-- toc -->
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [User Stories](#user-stories)
  - [API](#api)
  - [Exposed Endpoints](#exposed-endpoints)
  - [Specification](#specification)
  - [Behavior](#behavior)
  - [Configuration](#configuration)
  - [Security Considerations](#security-considerations)
  - [Compatibility and Versioning](#compatibility-and-versioning)
  - [Drawbacks](#drawbacks)
  - [Future Work](#future-work)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

ExternalDNS gets its desired state from in-tree `Source` implementations and writes it through a `Provider`.
Providers can already live out of tree through the webhook mechanism (`provider/webhook`).
Sources cannot.

This proposes a `webhook` source: a client, built into ExternalDNS, that fetches `[]*endpoint.Endpoint` from a separate process over HTTP/JSON, mirroring the provider webhook (a negotiated, versioned media type, JSON bodies reusing `endpoint.Endpoint`).
In v1 the webhook must be reachable on loopback only, which keeps the stance identical to the provider webhook: sidecar-only, no client authentication.
The source is polled once per reconcile loop, so its cadence is controlled by the existing `--interval` flag.

## Motivation

Every new source is code in this repository, with its dependencies, its tests, and its maintenance burden.
Requests for vendor-specific sources keep being declined for that reason.
The project has already stopped accepting in-tree providers and is moving existing ones out ([#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347)).
More in-tree sources could still be added, but a webhook interface is the better default for new integrations: it keeps them out of this repository, and adoption of the webhook can then guide whether any further in-tree sources are needed.
The two proposals attack the same problem from different sides: [proposal 007](007-target-from-resource.md) removes the need for vendor-specific sources for data that lives in Kubernetes, and the webhook source covers data that does not.
A recent example is [#6777](https://github.com/kubernetes-sigs/external-dns/pull/6777), a proposal to generate ACME DNS-01 delegation records in-tree.
One review comment there suggests an out-of-tree integration instead, a cert-manager webhook solver that publishes the challenge records, leaving ExternalDNS unchanged.
Today such an integration has to create and delete `DNSEndpoint` objects; a webhook source would give it a way to serve the records directly.

Today the escape hatches for "my records come from somewhere ExternalDNS does not know" are:

- `crd` source: requires something else to write `DNSEndpoint` objects, which means an extra controller, RBAC, and a CRD round trip.
- `unstructured` source: builds records from arbitrary Kubernetes resources using templates and filters, so it covers vendor CRDs without importing them, but only for data that lives in the Kubernetes API.
- `connector` source: gob-encoded over raw TCP. It is Go-specific, unauthenticated, unencrypted, unversioned, and cannot signal changes (`AddEventHandler` is a no-op).

The gap is data that does not live in Kubernetes at all, such as an internal registry, a CMDB, or another control plane.
A webhook source gives those systems a language-agnostic, versioned way to feed ExternalDNS, using a pattern the project already operates and documents.

### Goals

- Define an out-of-process source contract over HTTP/JSON, symmetrical with the provider webhook and reusing the existing JSON serialization of `endpoint.Endpoint`.
- Reuse the existing reconcile cadence: one `Endpoints` call per `--interval`, with no separate polling loop.
- Work unchanged with the existing wrapper pipeline (dedup, NAT64, target filter, post-processor) and alongside other `--source` values.
- Publish the contract as an OpenAPI document so authors in any language can generate servers and clients.
- Fail safely: a webhook outage must never be interpreted as "no records".

### Non-Goals

- Replacing in-tree sources or the `crd` source.
- Deprecating the `connector` source. It stays as is, and any deprecation is a separate decision.
- Forbidding new in-tree sources. The webhook is an additional option, and proposals for new sources are still evaluated on their merits.
- Hosting in-tree sources behind a webhook server (the source equivalent of `--webhook-server`). This can be a follow-up.
- Changing the `Source` interface.
- Publishing Go helper packages for webhook authors. The OpenAPI spec is the contract, and helpers can follow if there is demand.
- Push notifications from the webhook (event streaming). See Future Work.
- Transports other than HTTP/JSON in v1 (gRPC, Unix sockets). See Future Work.
- Reachability off loopback. v1 refuses non-loopback URLs at start-up, and client authentication is a prerequisite for lifting that restriction. See Future Work.

## Proposal

### User Stories

- A platform team keeps hostnames in an internal service registry or CMDB. They run a small sidecar that serves the registry as endpoints, with no CRs or extra controller.
- An infrastructure team publishes DNS names for resources that do not exist in Kubernetes (for example VMs or load balancers managed by another control plane) and wants ExternalDNS to own those records.
- A cert-manager webhook solver holds ACME DNS-01 challenge records and serves them as TXT endpoints,
  instead of creating and deleting `DNSEndpoint` objects (the integration discussed in [#6777](https://github.com/kubernetes-sigs/external-dns/pull/6777)).
  Records appear within one `--interval`, which this use case has to tolerate.
- A team writes its source in Python or Java because the Go-only `connector` is not an option.
- A test suite serves deterministic endpoints from a stub server to exercise a full ExternalDNS pipeline.

### API

The source type is `webhook`, selected with `--source=webhook` and combinable with other sources.
The transport is HTTP/JSON over loopback. `--webhook-source-url` must use scheme `http://` and a loopback host (`localhost`, `127.0.0.0/8`, or `::1`); any other value is rejected at start-up.

The HTTP/JSON transport uses its own media type, distinct from the provider one:

```text
application/external.dns.webhook.source+json;version=1
```

A distinct type lets negotiation fail fast when a source URL points at a provider webhook or the reverse.

| Operation | HTTP Method | Route      | Called by   | Description                               |
|-----------|-------------|------------|-------------|-------------------------------------------|
| Negotiate | GET         | /          | ExternalDNS | Validates the media type                  |
| Endpoints | GET         | /endpoints | ExternalDNS | Returns the desired endpoints             |

Both routes are served on the contract port (default `8889`).

`GET /` performs media-type negotiation only.
The webhook responds with the source media type in `Content-Type`; the body is an empty JSON object (`{}`) today, reserved for additive fields later.
There is no separate readiness route. A webhook that is still warming up returns `503 Service Unavailable` on `GET /endpoints`, which the client treats as a soft error (see Behavior).

`GET /endpoints` returns `200 OK` with a JSON array of `endpoint.Endpoint`, identical to the provider `/records` body:

```json
[
  {
    "dnsName": "app.example.com",
    "targets": ["192.0.2.10"],
    "recordType": "A",
    "recordTTL": 300,
    "providerSpecific": []
  }
]
```

### Exposed Endpoints

These routes are for the webhook container's own probes and observability. ExternalDNS never calls them.

| Purpose      | HTTP Method | Route    | Description                                  |
|--------------|-------------|----------|----------------------------------------------|
| K8s probe    | GET         | /healthz | Used by `livenessProbe` and `readinessProbe` |
| Open Metrics | GET         | /metrics | Optional                                     |

The recommended port is `8081`, bound to all interfaces.
It differs from the provider webhook's `8080` because containers in a pod share a network namespace, so a provider webhook and a source webhook running as sidecars in the same pod cannot both listen on one port.
A single process that implements both webhooks can serve one shared `/healthz` and `/metrics` instead.

The default contract port differs from the provider webhook (`8888`) so both can run as sidecars in one pod.

### Specification

The contract is published as `api/webhook-source.yaml`, an OpenAPI 3.0 document next to `api/webhook.yaml`.
It defines `/` and `/endpoints` with the source media type, and the `200`, `503`, and `5xx` responses described under Behavior.
The `endpoint` schema is referenced from `webhook.yaml` (`webhook.yaml#/components/schemas/endpoint`) rather than copied, so the two HTTP contracts cannot drift apart. In that schema `targets` is an array of strings.

The spec file is normative for webhook authors, and the Go client is the reference consumer.
It is linted in CI (for example Spectral or Redocly), so a malformed or inconsistent contract fails the build.
The OpenAPI document was not linted before, so the rule set starts permissive enough that the existing `webhook.yaml`, which the new document references, passes unchanged.
`/healthz` and `/metrics` are documented in the tutorial only, as in the provider webhook, because ExternalDNS does not call them.

### Behavior

Start-up: validate the URL (scheme `http://`, loopback host) and construct the client, then negotiate with retry and backoff by calling `GET /` with `Accept` set to the source media type and requiring a matching `Content-Type`.
Failure here is a hard error, because start-up misconfiguration should fail fast.

Each call to `Endpoints(ctx)` makes one `GET /endpoints` request.

- `200` with a valid body: return the endpoints. An empty list is valid and means "no desired records".
- `503`: the webhook is warming up. Return `provider.NewSoftErrorf` and keep existing records.
- Any other failure (connection error, timeout, non-`200` status, wrong `Content-Type`, over-limit body, undecodable body): return `provider.NewSoftErrorf`.

Runtime failures are always soft.
`multiSource` fails the whole iteration when any child source errors, and a plain `error` in the reconcile loop exits the process, so a misbehaving sidecar must not be able to stop DNS sync for every other source.
Hard errors are reserved for start-up.

With several webhooks, a soft error from any one of them skips the whole iteration for every source.
This is deliberate: applying a partial result would delete the existing records of the webhook that failed.
The consequence for operators is that an unreliable or slow webhook delays DNS sync for all sources, so each webhook should be as available as the cluster's other sources.

The HTTP client does not follow redirects; any `3xx` is a failure.

An error is never translated into an empty list. With `--policy=sync`, an empty list deletes owned records, so this is the key safety property.
The client cannot tell a legitimately empty list from one produced by a webhook that swallowed an upstream failure, so v1 has no client-side guard against empty responses.
Webhook authors must report upstream failures with an error status, never an empty list, and operators should adopt a new webhook with `--policy=upsert-only`.

Response validation: webhook output is untrusted.
Each endpoint is run through `endpoint.Endpoint.CheckEndpoint` (the same check in-tree sources rely on, which already rejects empty targets, unknown record types, and malformed MX/SRV/PTR targets).
If any entry fails, the whole response is rejected with a soft error and nothing from it is used.
Dropping only the bad entries would turn a malformed entry into the deletion of its existing record at the next plan.
Reusing `CheckEndpoint` keeps the webhook source and in-tree sources on one validation rule set, so the two cannot drift apart.

Labels: the `labels` field is discarded on receipt, and the client sets the `resource` label itself to `webhook/<host:port of its URL>`, as in-tree sources do (for example `crd/<namespace>/<name>`).
This gives conflict resolution and the TXT registry a stable origin, and a webhook cannot impersonate another source's `resource` or influence which source wins a conflict.

Source-level flags such as `--namespace`, `--label-filter`, `--annotation-filter`, and `--fqdn-template` have no effect on the webhook source. Filtering what is published is the webhook's job.
The domain filter and sync policy still apply later in the plan.

Change detection: `AddEventHandler` is a no-op. Changes are picked up on the next `--interval` tick, and `--update-events` does not react to webhook changes.

Wrapper pipeline: the source is built by `source.BuildWithConfig` like any other, so `wrappers.Build` applies dedup, NAT64, target filtering, and the post-processor unchanged.

Reference objects: `Endpoint.RefObjects` is tagged `json:"-"`, so Kubernetes event emission is unavailable for webhook endpoints. This is documented, not worked around.

Metrics: a `webhook_source` subsystem with `endpoints_requests_total` and `endpoints_errors_total`, mirroring `webhook_provider`. They carry a `webhook` label with the `host:port` of the URL so multiple webhooks can be told apart.

Caching: none in v1.

### Configuration

| Flag                             | Default                 | Description                                            |
|----------------------------------|-------------------------|--------------------------------------------------------|
| `--webhook-source-url`           | `http://localhost:8889` | Source webhook base URL, repeatable; loopback host required |
| `--webhook-source-timeout`       | `15s`                   | Round-trip timeout                                     |
| `--webhook-source-max-body-size` | `33554432` (32 MiB)     | Response body cap, `0` disables                        |

`--webhook-source-url` must use scheme `http://` and a loopback host (`localhost`, `127.0.0.0/8`, or `::1`).
Any other scheme or host is rejected at start-up with a clear error message.

The flag can be repeated. `--source=webhook` is given once, and each `--webhook-source-url` value creates one source instance, so several webhooks can feed one ExternalDNS.
The default URL applies only when no URL is given. Duplicate URLs are rejected at start-up.
The timeout and body size settings apply to every instance.
Each webhook in a pod needs its own contract port, in addition to the probe port discussed under Exposed Endpoints.

### Security Considerations

Loopback-only is the whole security story for v1, and it is enforced at start-up.
A compromised webhook source decides which DNS records ExternalDNS creates, and with `--policy=sync` an empty list deletes every owned record it owns. For a cert-manager style use case the webhook can also create DNS-01 TXT records and get certificates issued for any name in the zone.
Server-side TLS does not help here, because nothing authenticates the caller, so lifting the loopback restriction is deferred until client authentication exists.

- Enforced at start-up: `--webhook-source-url` must use `http://` and a loopback host (`localhost`, `127.0.0.0/8`, `::1`); any other value fails fast with a clear error.
- Recommended deployment is a sidecar in the same pod as ExternalDNS, matching the provider webhook stance.
- Operators should keep `--domain-filter` set and use `--policy=upsert-only` while adopting a new webhook, so a misbehaving webhook cannot delete records.
- Response size is capped, and the HTTP client does not follow redirects.

### Compatibility and Versioning

- The HTTP/JSON contract is versioned through the media type (`version=1`). A breaking change would introduce `version=2` alongside the old one.
- Media type matching is parameter-tolerant. The client parses `Content-Type` with `mime.ParseMediaType`, requires the type `application/external.dns.webhook.source+json` and `version=1`, and ignores other parameters such as `charset`.
  A webhook written with a framework that appends `; charset=utf-8` (common in Python and Node HTTP libraries) therefore works without special handling.
  The `Accept` header the client sends stays exact.
- The provider webhook client compares `Content-Type` for exact equality and rejects such replies ([#5176](https://github.com/kubernetes-sigs/external-dns/issues/5176)).
  That fix is tracked separately, and a shared media type parser would keep the two clients consistent.
- The JSON shape of `endpoint.Endpoint` is shared with the provider webhook, so changes to it are already announced in release notes and now cover both.
- Additive negotiation fields are ignored by older clients.
- The `connector` source is untouched. Deprecating it is out of scope here.

### Drawbacks

- A second public, versioned API surface to keep stable, while the provider webhook is itself still described as experimental.
- Network hop and an extra process to operate for what could be a library call.
- Empty-but-valid responses can delete records; the safety rests partly on webhook authors returning an error status correctly.
- No Kubernetes object references, so no per-object events.
- No change notifications, so reaction time is bounded by `--interval`.
- Loopback-only limits deployment topology (one webhook per cluster node, or one per ExternalDNS replica). Shared or centralized webhook services are not possible until client authentication lands.

### Future Work

None of the following is committed by this proposal, but each could be or should be considered.

- gRPC transport, when a concrete need appears (for example server-streaming change notifications, or an in-cluster gRPC service). Adding it today would double the contract (OpenAPI + proto + parity test), pull in protobuf tooling, and fight the binary-size effort, so it waits for real demand.
- `unix://` socket transport on a shared `emptyDir`, where file permissions give real access control without needing network-level authentication.
- Client authentication (bearer token or mTLS) so webhooks reachable off loopback become safe; once that lands the loopback restriction can be relaxed.
- Change notifications, advertised through negotiation, to make `AddEventHandler` meaningful: server-sent events over HTTP (and a server-streaming RPC if gRPC is added).
- Conditional requests (`ETag` and `If-None-Match`) so unchanged state is cheap to poll at short intervals.
- A `--webhook-server` style mode that serves in-tree sources over this contract.
- Webhook-supplied labels, limited to a safe subset (for example a reserved prefix), once there is a concrete need. Accepting arbitrary labels would let one webhook impersonate another source's `resource` and win conflicts, so v1 does not.

## Alternatives

1. **Extend `connector`.** Adds versioning and events to a gob/TCP protocol that is Go-only and plaintext. Diverges from the provider webhook.
2. **gRPC transport in v1.** A better schema and streaming story, but with no streaming use case in v1 and the sidecar-on-loopback deployment, it would only double the contract and dependencies. Deferred to Future Work.
3. **`DNSEndpoint` CRs written by a separate controller.**
   Works today and stays the right choice for Kubernetes-native producers,
   and wins on visibility (`kubectl get`, status) and RBAC-scoped writes.
   The webhook source wins when custom functionality is required that may not be in scope for ExternalDNS to support in-tree.
4. **Generic resource dereferencing ([proposal 007](007-target-from-resource.md)) and the `unstructured` source.** Both solve Kubernetes resource indirection, not non-Kubernetes data sources. Complementary.
5. **Exec or plugin binaries.** Needs a writable runtime and execution trust inside the ExternalDNS image.
