# ACME DNS-01 Delegation CNAME Records

> To automatically create ACME DNS-01 delegation CNAME records for your hostnames,
> use the `--acme-cname-delegation-target-template` flag.

## Background

ACME clients such as [cert-manager](https://cert-manager.io/) solve DNS-01 challenges by
publishing a TXT record at `_acme-challenge.<hostname>`. Granting the ACME client write
access to every application DNS zone is often undesirable. A common pattern —
[delegated domains for DNS-01](https://cert-manager.io/docs/configuration/acme/dns01/#delegated-domains-for-dns01) —
is to point the challenge name at a dedicated zone via a CNAME:

```text
_acme-challenge.app.example.com CNAME app.example.com.acme.example.net
```

The ACME client then only needs write permissions for the dedicated challenge zone
(`acme.example.net`), not for the production zone (`example.com`). This shrinks the blast
radius of the ACME client's DNS credentials and lets application teams self-service
Ingresses, Gateway API routes, and Services without DNS write permissions.

ExternalDNS already discovers every hostname that needs regular DNS records, so it can
create the matching delegation CNAMEs as well.

## How it works

When `--acme-cname-delegation-target-template` is set, ExternalDNS additionally creates a
CNAME record `_acme-challenge.<hostname>` for every hostname discovered from the configured
sources (A, AAAA and CNAME records). The CNAME target is rendered from the given Go template.

| Flag | Description |
|:-----|:------------|
| `--acme-cname-delegation-target-template` | Go template for the CNAME target. Setting it enables the feature. |
| `--acme-cname-delegation-domain-filter` | Limit generation to hostnames matching these domain suffixes. Repeatable. Default: all hostnames. |
| `--acme-cname-delegation-ttl` | TTL for the generated CNAMEs, e.g. `300s`. Default: `--min-ttl` or the provider default. |

All flags can also be set as environment variables, e.g.
`EXTERNAL_DNS_ACME_CNAME_DELEGATION_TARGET_TEMPLATE`.

### Template fields

| Field | Value for `app.example.com` | Value for `*.app.example.com` |
|:------|:----------------------------|:------------------------------|
| `{{ .Hostname }}` | `app.example.com` | `*.app.example.com` |
| `{{ .HostnameWithoutWildcard }}` | `app.example.com` | `app.example.com` |

The shared template functions from [FQDN templating](fqdn-templating.md) (`replace`,
`trimPrefix`, `toLower`, …) are available.

## Example

```sh
external-dns \
  --provider=aws \
  --source=ingress \
  --source=gateway-httproute \
  --domain-filter=example.com \
  --acme-cname-delegation-target-template='{{ .HostnameWithoutWildcard }}.acme.example.net' \
  --acme-cname-delegation-ttl=300s
```

Given an `HTTPRoute` with the hostnames `app.example.com` and `*.dev.example.com`,
ExternalDNS manages the regular records plus:

```text
_acme-challenge.app.example.com CNAME app.example.com.acme.example.net
_acme-challenge.dev.example.com CNAME dev.example.com.acme.example.net
```

For [acme-dns](https://github.com/joohoi/acme-dns) style targets, template functions can
flatten the hostname:

```sh
--acme-cname-delegation-target-template='{{ replace "." "-" .HostnameWithoutWildcard }}.acme.example.net'
```

## Behaviour details

- Exactly one delegation CNAME is created per hostname: a hostname with both A and AAAA
  records, or a wildcard plus its base hostname (`*.app.example.com` and `app.example.com`),
  results in a single CNAME.
- Wildcard hostnames map to the challenge name without the wildcard label:
  `*.app.example.com` → `_acme-challenge.app.example.com`. This matches how ACME validates
  wildcard certificates.
- Explicitly defined `_acme-challenge.*` endpoints (for example from a `DNSEndpoint`
  resource) always take precedence: ExternalDNS never generates a delegation CNAME for a
  challenge name that a source already provides.
- The generated CNAMEs live in the same zone as the hostname, so the existing
  `--domain-filter` and registry ownership mechanisms apply. Ownership TXT records are
  created automatically; stale delegation CNAMEs are cleaned up when the hostname disappears.
- ExternalDNS only creates the delegation CNAME. The ACME client must publish the challenge
  TXT record in the target zone (e.g. cert-manager with credentials scoped to
  `acme.example.net`).

## Caveats

- CNAME must be included in `--managed-record-types` (it is by default). ExternalDNS refuses
  to start when the feature is enabled without it.
- When `--regex-domain-filter` is used, the regular expression must also match the generated
  `_acme-challenge.*` names, otherwise they are filtered out at planning time. ExternalDNS
  logs a warning for this combination at startup.
- Some DNS providers do not accept underscore-prefixed record names; verify provider support
  before enabling the feature.
