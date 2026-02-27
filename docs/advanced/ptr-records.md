# Automatic PTR (Reverse DNS) Records

> To automatically create PTR records for your A/AAAA endpoints, use the `--create-ptr` flag.
> This feature works with any provider that has authority over reverse DNS zones
> (e.g. `in-addr.arpa` for IPv4 or `ip6.arpa` for IPv6).

## Modes

The `--create-ptr` flag accepts three values:

| Mode         | Behaviour |
|:-------------|:----------|
| `off`        | PTR record creation is disabled. This is the default. |
| `always`     | A PTR record is created for every A/AAAA endpoint. |
| `annotation` | A PTR record is created only when the source resource is annotated with `external-dns.alpha.kubernetes.io/create-ptr: "true"`. |

## Prerequisites

The underlying DNS provider must have authority over the relevant reverse DNS zones.
Include the reverse zone in `--domain-filter` so that ExternalDNS knows it is allowed to manage records there:

```sh
--create-ptr=always --domain-filter=example.com --domain-filter=49.168.192.in-addr.arpa
```

## Usage

### Always mode

In `always` mode every A/AAAA record managed by ExternalDNS automatically gets a corresponding PTR record.
No annotation is needed on individual resources.

```sh
external-dns \
  --provider=pdns \
  --create-ptr=always \
  --domain-filter=example.com \
  --domain-filter=49.168.192.in-addr.arpa
```

### Annotation mode

In `annotation` mode only resources annotated with `external-dns.alpha.kubernetes.io/create-ptr: "true"`
produce PTR records. This gives fine-grained, per-resource control.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: web
  annotations:
    external-dns.alpha.kubernetes.io/hostname: web.example.com
    external-dns.alpha.kubernetes.io/create-ptr: "true"
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

## Behaviour details

- Wildcard records (e.g. `*.example.com`) are skipped — a PTR pointing to a wildcard hostname is not meaningful.
- When multiple A/AAAA records point to the same IP address, a single PTR record is created with all hostnames as targets.
- PTR records whose forward hostname does not match `--domain-filter` are automatically cleaned up.
- PTR records are tracked by the TXT registry like any other record type, so ownership is preserved across restarts.

## Deprecation of `--rfc2136-create-ptr`

The `--rfc2136-create-ptr` flag is deprecated in favor of `--create-ptr`. If `--rfc2136-create-ptr` is set
and `--create-ptr` is not, it is treated as `--create-ptr=always` for backward compatibility.
Migrate to `--create-ptr=always` (or `--create-ptr=annotation`) and remove `--rfc2136-create-ptr` from
your configuration.
