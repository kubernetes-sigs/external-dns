---
tags: ["tutorial", "fips", "security", "compliance"]
---

# Compiling ExternalDNS for FIPS

## Overview

ExternalDNS does not ship pre-built FIPS-compliant images, and there's no current plan to add them. The supported
path today is building your own FIPS-capable binary/image from source, using [Go's native FIPS 140-3
support](https://go.dev/doc/security/fips140) (Go 1.24+): setting `GOFIPS140` at build time compiles in Go's
CMVP-validated cryptographic module and makes `crypto/tls` restrict itself to FIPS-approved algorithms for every TLS
connection ExternalDNS makes - to the Kubernetes API server, to DNS provider APIs, and to webhook providers. It's
pure Go - no BoringCrypto, cgo, or OpenSSL FIPS provider, so there's no separate crypto toolchain to maintain
alongside the regular build.

This gets you a FIPS-*capable* build, not a FIPS-*accredited* deployment - accreditation is an organizational
process (ATO, operational environment review, etc.) layered on top. And capable is not the same as compliant:
auditors under FedRAMP, DoD, HIPAA, or similar frameworks require proof that FIPS enforcement is active at runtime,
not just that a FIPS-capable build exists - verify enforcement rather than relying on a `-fips` tag or vendor claims
alone.

## Build and Verify the Image

```sh
make build.image-fips
```

This builds with `--push=false` and no `--local`, matching the existing `build.image` targets in the Makefile - it
proves the build succeeds without publishing or loading anything locally. Since there's no local artifact yet to
inspect, verify the FIPS module lands in the binary by building it the same way directly:

```sh
GOFIPS140=v1.0.0 CGO_ENABLED=0 go build -o build/external-dns .
go version -m build/external-dns | grep -i -E "fips|GODEBUG"

build	-tags=fips140v1.0
build	DefaultGODEBUG=fips140=on
build	GOFIPS140=v1.0.0-c2097c7c
```

`DefaultGODEBUG=fips140=on` means FIPS mode is the runtime default the moment the binary is built this way - there's
no need to also set `GODEBUG=fips140=on` on the container. A default (non-FIPS) build shows none of these lines.

## Alternative: Vendor-Hardened Images

Rolling your own with `build.image-fips` proves the mechanism works, but the ongoing burden - tracking
module/certificate staleness, rebuild cadence, and producing audit-ready evidence - stays with your team.

Vendors like [Chainguard](https://images.chainguard.dev/directory/image/external-dns-fips/overview) ship a pre-built
`external-dns-fips` image and carry that burden as part of a support relationship: daily rebuilds, SLSA Level 3
provenance, and Sigstore-signed attestations. Their kernel-independence is achieved differently from what's described
above, though - Chainguard relocates the SP 800-90B entropy source into userspace via OpenSSL and the
[Jitter Entropy Library](https://www.chainguard.dev/unchained/kernel-independent-fips-images), rather than Go's native
FIPS 140-3 module. Both approaches avoid depending on a FIPS-enabled kernel, just via different mechanisms - see also
Chainguard's [overview of when/where/why FIPS-validated images are needed](https://www.chainguard.dev/unchained/fips-validated-container-images-when-where-why).

If an auditor wants outside evidence of a continuously maintained compliance boundary, a vendor image is usually the
better fit than self-building.

## Using It with AWS

FIPS image or not, AWS endpoint routing is a separate, already-solved concern - no code change needed. Set these as
env vars (or extra args) on the deployment:

```sh
AWS_USE_FIPS_ENDPOINT=true
AWS_REGION=us-east-1  # must be a region with a FIPS endpoint for the services you use (Route53, Cloud Map, DynamoDB, etc.)
```

The AWS SDK resolves to the FIPS endpoint (e.g. `route53-fips.amazonaws.com`) for that region on its own. Don't
conflate this with the `GOFIPS140` build flag above - one picks which AWS hostname to call, the other governs which
crypto module TLS uses to make that call.

## A Note Before You Deploy

If you're running a FIPS-hardened `external-dns` controller in production - self-built or vendor image - please open
a PR to update this guide with what you learned. This page reflects a first pass; real deployment experience
(verification steps, gotchas, auditor feedback) is the part most worth sharing.

## Additional Resources

- [Go FIPS 140-3 support](https://go.dev/doc/security/fips140)
- [Chainguard: FIPS-validated container images - when, where, why](https://www.chainguard.dev/unchained/fips-validated-container-images-when-where-why)
- [Chainguard: kernel-independent FIPS images](https://www.chainguard.dev/unchained/kernel-independent-fips-images)
- [Chainguard external-dns-fips image](https://images.chainguard.dev/directory/image/external-dns-fips/overview)
