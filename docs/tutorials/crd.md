# Using CRD Source for DNS Records

This tutorial describes how to use the CRD source with ExternalDNS to manage DNS records. The CRD source allows you to define your desired DNS records declaratively using `DNSEndpoint` custom resources.

## Default Targets and CRD Targets

ExternalDNS has a `--default-targets` flag that can be used to specify a default set of targets for all created DNS records. The behavior of how these default targets interact with targets specified in a `DNSEndpoint` CRD has been refined.

### New Behavior (default)

By default, ExternalDNS now has the following behavior:

- If a `DNSEndpoint` resource has targets specified in its `spec.endpoints[].targets` field, these targets will be used for the DNS record, **overriding** any targets specified via the `--default-targets` flag.
- If a `DNSEndpoint` resource has an **empty** `targets` field, the targets from the `--default-targets` flag will be used. This allows for creating records that point to default load balancers or IPs without explicitly listing them in every `DNSEndpoint` resource.

### Legacy Behavior (`--force-default-targets`)

To maintain backward compatibility and support certain migration scenarios, the `--force-default-targets` flag is available.

- When `--force-default-targets` is used, ExternalDNS will **always** use the targets from `--default-targets`, regardless of whether the `DNSEndpoint` resource has targets specified or not.
This flag allows for a smooth migration path to the new behavior. It allow keeping old CRD resources, allows to start removing targets from one by one resource and then remove the flag.

## Examples

Let's look at how this works in practice. Assume ExternalDNS is running with `--default-targets=1.2.3.4`.

### DNSEndpoint with Targets

Here is a `DNSEndpoint` with a target specified.

```yaml
---
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: targets
  namespace: default
spec:
  endpoints:
  - dnsName: smoke-t.example.com
    recordTTL: 300
    recordType: CNAME
    targets:
      - placeholder
```

- **Without `--force-default-targets` (New Behavior):** A CNAME record for `smoke-t.example.com` will be created pointing to `placeholder`.
- **With `--force-default-targets` (Legacy Behavior):** A CNAME record for `smoke-t.example.com` will be created pointing to `1.2.3.4`. The `placeholder` target will be ignored.

### DNSEndpoint with Empty/No Targets

Here is a `DNSEndpoint` without any targets specified.

```yaml
---
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: no-targets
  namespace: default
spec:
  endpoints:
  - dnsName: smoke-nt.example.com
    recordTTL: 300
    recordType: CNAME
```

- **Without `--force-default-targets` (New Behavior):** A CNAME record for `smoke-nt.example.com` will be created pointing to `1.2.3.4`.
- **With `--force-default-targets` (Legacy Behavior):** A CNAME record for `smoke-nt.example.com` will be created pointing to `1.2.3.4`.

`--force-default-targets` allows migration path to clean CRD resources.
