# Automatic PTR (Reverse DNS) Records

> To automatically create PTR records for your A/AAAA endpoints, use the `--create-ptr` flag.
> This feature works with any provider that has authority over reverse DNS zones
> (e.g. `in-addr.arpa` for IPv4 or `ip6.arpa` for IPv6).

## Flag and annotation

The `--create-ptr` flag is a boolean (default: `false`) that sets the global default for PTR record creation.
The `external-dns.alpha.kubernetes.io/record-type` annotation on a resource overrides this default,
following the standard [configuration precedence](configuration-precedence.md):

| Flag    | Annotation             | Result                                            |
|:--------|:-----------------------|:--------------------------------------------------|
| `true`  | (absent)               | PTR record created                                |
| `true`  | `""` (empty)           | PTR record **not** created (annotation opts out)  |
| `false` | (absent)               | PTR record **not** created                        |
| `false` | `"ptr"`                | PTR record created (annotation opts in)           |

## Prerequisites

The underlying DNS provider must have authority over the relevant reverse DNS zones.
Include the reverse zone in `--domain-filter` so that ExternalDNS knows it is allowed to manage records there.

PTR must also be included in `--managed-record-types` so the planner considers PTR records during sync:

```sh
  --create-ptr \
  --managed-record-types=A \
  --managed-record-types=AAAA \
  --managed-record-types=CNAME \
  --managed-record-types=PTR \
  --domain-filter=example.com \
  --domain-filter=49.168.192.in-addr.arpa
```

## Usage

### Global PTR creation

With `--create-ptr` enabled, every A/AAAA record managed by ExternalDNS automatically gets a corresponding PTR record.
Individual resources can opt out by setting the annotation to an empty value.

```sh
external-dns \
  --provider=pdns \
  --create-ptr \
  --domain-filter=example.com \
  --domain-filter=49.168.192.in-addr.arpa
```

### Per-resource PTR creation

Without `--create-ptr`, only resources annotated with `external-dns.alpha.kubernetes.io/record-type: "ptr"`
produce PTR records.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: web
  annotations:
    external-dns.alpha.kubernetes.io/hostname: web.example.com
    external-dns.alpha.kubernetes.io/record-type: "ptr"
spec:
  type: LoadBalancer
  ports:
  - port: 80
  selector:
    app: web
```

If the Service receives the external IP `192.168.49.2`, ExternalDNS creates both
`web.example.com → 192.168.49.2` (A) and `2.49.168.192.in-addr.arpa → web.example.com` (PTR).

Resources without the annotation are not affected.

### DNSEndpoint CRD

The DNSEndpoint CRD source does not process Kubernetes resource annotations.
To enable PTR for individual CRD endpoints, set the `record-type` property
directly in `spec.endpoints[].providerSpecific`:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: web
spec:
  endpoints:
  - dnsName: web.example.com
    recordType: A
    targets:
    - 192.168.49.2
    providerSpecific:
    - name: record-type
      value: ptr
```

When `--create-ptr` is enabled globally, PTR records are generated for all
A/AAAA endpoints regardless of source, including DNSEndpoint.

## Behaviour details

- Wildcard records (e.g. `*.example.com`) are skipped — a PTR pointing to a wildcard hostname is not meaningful.
- When multiple A/AAAA records point to the same IP address, their hostnames are grouped into a single
  ExternalDNS endpoint. In DNS this produces multiple PTR resource records in the same RRSet
  (one per hostname), which is valid per [RFC 1035](https://datatracker.ietf.org/doc/html/rfc1035).
  For example, if both `a.example.com` and `b.example.com` resolve to `192.168.49.2`,
  a `dig -x 192.168.49.2` will return two PTR answers.
  Note that not all providers may handle multi-target PTR records correctly — verify with your
  provider if this applies to your setup.
- PTR records whose forward hostname does not match `--domain-filter` are automatically cleaned up.
- PTR records are tracked by the TXT registry like any other record type, so ownership is preserved across restarts.
