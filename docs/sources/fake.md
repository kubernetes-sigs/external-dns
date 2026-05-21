---
tags:
  - sources
  - fake
  - testing
---

# Fake Source

The fake source generates synthetic DNS endpoints without requiring a Kubernetes cluster or any real resources.
It produces one endpoint per supported record type per configured domain on every reconciliation cycle,
using documentation-reserved address ranges (RFC 5737 for IPv4, RFC 3849 for IPv6).
`A` and `AAAA` records carry one target per configured domain.

## Use cases

The fake source is a **smoke-test tool**: it answers "does the provider connect and accept records?" quickly,
without any cluster dependency. It is not a substitute for a contract test with known, stable data —
use the [DNSEndpoint CRD source](crd.md) when you need precise control over which record types
are created, how many, and with what values.

**Dry-running a DNS provider**
Validate provider credentials, API connectivity, and record formatting before pointing ExternalDNS at a real cluster:

```console
external-dns --source=fake --provider=aws --domain-filter=example.com --dry-run
```

**Smoke-testing a webhook provider**
Verify that a webhook implementation connects and processes records without needing a Kubernetes cluster:

```console
external-dns --source=fake --provider=webhook --webhook-provider-url=http://localhost:8888
```

For a proper integration test with predictable inputs (specific record types, stable names, exact counts), define your test fixture as `DNSEndpoint` resources and use `--source=crd` instead.

**CI connectivity check**
Spin up ExternalDNS with `--source=fake` in a pipeline to confirm provider connectivity end-to-end without any cluster dependency.

## Generated endpoints

On each reconciliation the fake source emits one endpoint per record type:

| Record type | DNS name                   | Example target                                                | Sources that emit this type |
|:------------|:---------------------------|:--------------------------------------------------------------|:----------------------------|
| `A`         | `<random>.example.com`     | `192.0.2.1` (one per domain)                                  | All sources                 |
| `AAAA`      | `<random>.example.com`     | `2001:db8::1a2b:3c4d` (one per domain)                        | All sources                 |
| `CNAME`     | `<random>.example.com`     | `<random>.example.com`                                        | All sources                 |
| `TXT`       | `<random>.example.com`     | `"heritage=external-dns,external-dns/owner=fake"`             | All sources                 |
| `NS`        | `example.com`              | `<random>.example.com`                                        | Most sources                |
| `MX`        | `example.com`              | `10 <random>.example.com`                                     | Most sources                |
| `SRV`       | `_sip._udp.example.com`    | `10 20 5060 <random>.example.com.`                            | CRD (DNSEndpoint) only      |
| `PTR`       | `<n>.2.0.192.in-addr.arpa` | `<random>.example.com`                                        | CRD (DNSEndpoint) only      |
| `NAPTR`     | `_sip._udp.example.com`    | `100 10 "u" "E2U+sip" "!^.*$!sip:info@example.com!" .`        | CRD (DNSEndpoint) only      |

The NAPTR target fields are: `order preference flags service regexp replacement` (RFC 2915). In the example above: order=100, preference=10, flags=`"u"` (URI result), service=`"E2U+sip"`, regexp=`"!^.*$!sip:info@example.com!"`, replacement=`.` (none).

> **Note:** `SRV`, `PTR`, and `NAPTR` are only reachable in practice via the CRD source (`DNSEndpoint`). The fake source emits them so webhook providers can verify handling of these types, but a passing result does not indicate that any real source will produce them.

IPv4 addresses are drawn from `192.0.2.0/24` and IPv6 from `2001:db8::/32` — both reserved for documentation and examples, so they will never accidentally match real infrastructure.

## Enabling specific record types

By default ExternalDNS only manages `A` and `CNAME` records. Use `--managed-record-types` to opt in to additional types:

```console
external-dns \
  --source=fake \
  --provider=webhook \
  --webhook-provider-url=http://localhost:8888 \
  --managed-record-types=A \
  --managed-record-types=AAAA \
  --managed-record-types=CNAME \
  --managed-record-types=TXT \
  --managed-record-types=SRV \
  --managed-record-types=NS \
  --managed-record-types=MX \
  --managed-record-types=NAPTR
```

To test all record types at once, list every type explicitly. The fake source always generates a full set; `--managed-record-types` controls which ones the provider receives.

## Kubernetes events

When `--emit-events` is configured, the fake source emits Kubernetes events for every DNS change, referencing a synthetic `Pod` object in the `default` namespace. This lets you observe the event stream during testing without real workloads:

```console
external-dns --source=fake --provider=webhook --webhook-provider-url=http://localhost:8888 \
  --emit-events=RecordReady
```

```console
kubectl get events -n default --field-selector reason=RecordReady
```

## Custom domain with `--fqdn-template`

By default the fake source generates endpoints under `example.com`. Use `--fqdn-template` to replace it with your own domain. The template is rendered against a synthetic `Pod` object with `Name=fake` and `Namespace=default`.

```console
# plain domain
external-dns --source=fake --provider=webhook --webhook-provider-url=http://localhost:8888 \
  --fqdn-template=my-company.com

# template expression
external-dns --source=fake --provider=webhook --webhook-provider-url=http://localhost:8888 \
  --fqdn-template={{.Name}}.my-company.com
```

The second example renders to `fake.my-company.com`, so endpoints are generated as `<random>.fake.my-company.com`.

**Multiple domains** can be specified as a comma-separated list. The fake source generates a full set of record types for each domain, and `A`/`AAAA` records carry one target per domain:

```console
external-dns --source=fake --provider=webhook --webhook-provider-url=http://localhost:8888 \
  --fqdn-template=one.example.com,two.example.com,three.example.com
```

This produces `3 × 9 = 27` endpoints in total, with each `A` and `AAAA` record having 3 targets.

> **Note:** Multiple domains add volume but not predictability — names are still random and change on every reconciliation. If your integration test needs a stable, known set of records, use `DNSEndpoint` resources with `--source=crd` instead.

## Limitations

- Endpoints are regenerated with random names on every reconciliation; the fake source does not model resource updates or deletions.
- No Kubernetes cluster is required, but `--emit-events` will fail to post events when kubernetes API server is unreachable.
