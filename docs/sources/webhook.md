---
tags:
  - sources
  - webhook
---

# Webhook Source

The webhook source fetches the desired DNS records from a separate process over HTTP/JSON.
It is the source-side counterpart of the [webhook provider](../tutorials/webhook-provider.md):
instead of reading Kubernetes resources, ExternalDNS polls a webhook that returns
`endpoint.Endpoint` objects, once per reconcile loop. The design is recorded in
[proposal 008](../proposal/008-webhook-source.md).

Use it for records that live outside the Kubernetes API, such as an internal service registry,
a CMDB, or records derived by a program that is not worth adding to ExternalDNS in tree.
For records that another controller can express as Kubernetes objects, the
[DNSEndpoint CRD source](crd.md) remains the simpler choice: it gives `kubectl get`, status,
and RBAC-scoped writes for free.

## Deployment

In this version the webhook must listen on loopback. ExternalDNS rejects any
`--webhook-source-url` that does not use the `http` scheme and a loopback host
(`localhost`, `127.0.0.0/8` or `::1`), so the webhook runs as a sidecar container in the
ExternalDNS pod. Nothing authenticates the caller; a webhook reachable from the network would
let any client decide which records ExternalDNS manages.

```yaml
containers:
  - name: external-dns
    image: registry.k8s.io/external-dns/external-dns:<version> # the first release that ships the webhook source
    args:
      - --source=webhook
      - --webhook-source-url=http://localhost:8889
      - --provider=aws
      - --domain-filter=example.com
      - --policy=upsert-only
  - name: source-webhook
    image: example.com/my-source-webhook:1.0.0
    ports:
      - name: http
        containerPort: 8081 # probes and metrics, not the contract port
    livenessProbe:
      httpGet:
        path: /healthz
        port: http
    readinessProbe:
      httpGet:
        path: /healthz
        port: http
```

The source can be combined with other `--source` values, and `--webhook-source-url` can be
repeated to poll several webhooks from one ExternalDNS. Each webhook needs its own contract port,
because containers in a pod share one network namespace.

| Flag                             | Default                 | Description                                                            |
|----------------------------------|-------------------------|------------------------------------------------------------------------|
| `--webhook-source-url`           | `http://localhost:8889` | Base URL of a source webhook; repeatable, `http` and loopback required |
| `--webhook-source-timeout`       | `15s`                   | Round-trip timeout of one request, including reading the body          |
| `--webhook-source-max-body-size` | `33554432` (32 MiB)     | Response body cap in bytes; `0` disables it                            |

Source-level flags such as `--namespace`, `--label-filter`, `--annotation-filter` and
`--fqdn-template` have no effect on this source. Filtering what gets published is the webhook's job.
`--domain-filter` and the sync policy still apply in the plan.

## Behavior

- **Start-up.** ExternalDNS validates every URL, then calls `GET /` on each webhook with the
  `Accept` header set to the source media type and requires a matching `Content-Type`. Connection
  errors and `5xx` responses are retried with backoff; a wrong media type or any other failure is
  a hard error, so a misconfigured pod fails fast.
- **Reconcile.** Each loop makes one `GET /endpoints` request per webhook. A `200` with a JSON array
  returns those endpoints; an empty array is valid and means "no desired records".
- **Failures are soft.** A `503` (webhook still warming up), a connection error, a timeout, any
  other status code, a redirect, a wrong `Content-Type`, a body over the limit, or a body that does
  not decode is reported as a soft error. ExternalDNS logs it, keeps the existing records and retries
  on the next loop. An error is never turned into an empty list, because with `--policy=sync` an
  empty list deletes every record the webhook owns.
- **All or nothing.** With several webhooks, a failure of one skips the whole iteration for every
  source. Applying a partial result would delete the records of the webhook that failed.
- **Validation.** Every endpoint must carry a name without overlong labels, a record type ExternalDNS
  knows, a TTL that is not negative, and pass the same checks the in-tree sources rely on (at least
  one target, valid IP addresses for `A` and `AAAA`, well-formed `MX`, `SRV`, `NAPTR` and `PTR`
  targets). One invalid entry rejects the whole response: dropping only that entry would turn a
  malformed record into the deletion of its existing one at the next plan.
- **Labels.** The `labels` field of the response is discarded. ExternalDNS sets the `resource`
  label to `webhook/<host:port>` itself, so a webhook cannot impersonate another source or decide
  conflicts in its favor.
- **No events.** The contract has no change notifications, so a change shows up on the next
  `--interval` tick; `--events` does not react to webhook changes.
- **No object references.** Kubernetes events are not emitted for records from a webhook, and the
  `DNSEndpoint` status sync does not apply to them.

The source counts its requests in the `external_dns_webhook_source_endpoints_requests_total` and
`external_dns_webhook_source_endpoints_errors_total` metrics, labeled with the webhook's `host:port`.

## Writing a webhook

The media type is `application/external.dns.webhook.source+json;version=1`. It is distinct from the
provider webhook's media type on purpose: a `--webhook-source-url` that points at a provider webhook,
or the other way round, fails negotiation instead of misbehaving later. ExternalDNS parses the
`Content-Type` header as a media type, so a framework that appends `; charset=utf-8` is fine, but the
`version` parameter must match.

| Route        | Method | Called by   | Response                                                         |
|--------------|--------|-------------|------------------------------------------------------------------|
| `/`          | GET    | ExternalDNS | `200` with the media type in `Content-Type`; the body is ignored |
| `/endpoints` | GET    | ExternalDNS | `200` with a JSON array of endpoints, `503` while not ready      |

The endpoint JSON is the same as the provider webhook's `/records` body:

```json
[
  {
    "dnsName": "app.example.com",
    "targets": ["192.0.2.10"],
    "recordType": "A",
    "recordTTL": 300
  },
  {
    "dnsName": "_acme-challenge.app.example.com",
    "targets": ["app.example.com.acme.example.net"],
    "recordType": "CNAME"
  }
]
```

Rules for authors:

- **Report upstream failures with an error status, never with an empty list.** ExternalDNS cannot
  tell a legitimately empty list from one produced by a webhook that swallowed a failure. Operators
  should run a new webhook with `--policy=upsert-only` until they trust it.
- Return `503` while warming up. Any `5xx` is retried at start-up and treated as a soft error later.
- Do not redirect. The client does not follow redirects and treats a `3xx` as a failure.
- Keep the response under the configured body limit and finish writing promptly.
- Serve `/healthz` and, optionally, `/metrics` for Kubernetes probes and monitoring on a separate
  port bound to all interfaces. The recommended port is `8081`; it differs from the provider
  webhook's `8080` so that both kinds of sidecar can run in one pod. The recommended contract port
  is `8889`, for the same reason.
