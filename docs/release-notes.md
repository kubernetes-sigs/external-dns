# Release notes

## Unreleased

### Breaking changes

### Significant changes

* Conflicts between CNAME records and records of other types for the same domain
are now resolved. The CNAME is given lower priority. This fixes a regression in
v0.13.5.

### New features

* The `--exclude-record-types` flag may be used to remove record types from the
set of managed record types.

* The `azure` provider now supports AAAA records.

* The `godaddy` provider now supports domains that weren't registered through GoDaddy itself.

* The `--gloo-namespace` flag may now be specified multiple times.

### Other changes

* The `service` source now avoids publishing addresses for unready nodes for `NodePort` services.

* Fixed a bug in the `txt` registry's handling of uppercase characters in templated prefixes or suffixes.

* Fixed a bug where the `google` provider would attempt to publish records to peering zones.

* Fixed a bug in the `linode` provider that caused unbounded creation of TXT records.

## v0.13.6

### Known issues

* Users upgrading from versions < v0.12.0 with the txt registry (--registry=txt) to this release
should run v0.13.5 at least once (--once) to avoid #3876

### Breaking changes

* The `exoscale` provider no longer supports the `--exoscale-endpoint` flag. It has been replaced 
with the `--exoscale-apienv` and `--exoscale-apizone` flags.

* The `pdns` provider no longer applies "MatchParent" semantics to its domain filters. It
now interprets domain filters the same way as other providers.

### Significant changes

* The image is now built with ko and is based on a distroless base image.

* Updated the AWS IAM policy example to include the `route53:ListTagsForResource` permission.
This is required for the `--aws-zone-tags` flag.

### New features

* Added the [`dynamodb` registry](registry/dynamodb.md).

* Added the `traefik-proxy` source.

* The `istio-gateway` and `istio-virtualserver` sources now support an
`external-dns.alpha.kubernetes.io/ingress` annotation on a `Gateway` to
indicate it is fed from an `Ingress` instead of a `Service`.

* The gloo-proxy source added support for `listener.metadataStatic` VirtualService source objects.

* The `--default-targets` flag now supports creating AAAA and CNAME targets.

* The `aws` provider now supports aliases to the AWS API Gateway. It also added support for
the `ap-southeast-4` region.

* The `azure` providers now support Workload Identity using azidentity.

### Other changes

* The `rfc2136` provider now always uses TCP.

* The `pdns` provider now enables TLS unconditionally.

* Fixed a bug where the txt registry wouldn't delete no-longer-needed records
when the `--txt-prefix` flag value had a template.

* Fixed a bug where some changes to provider-specific properties would not be reconciled.
