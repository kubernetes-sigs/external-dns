Configure DNS record TTL (Time-To-Live)
=======================================

An optional annotation `external-dns.alpha.kubernetes.io/ttl` is available to customize the TTL value of a DNS record.

To configure it, simply annotate a service/ingress, e.g.:

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com.
    external-dns.alpha.kubernetes.io/ttl: "60"
  ...
```

TTL must be a positive integer encoded as string.

Providers
=========

- [x] AWS (Route53)
- [ ] Azure
- [ ] Cloudflare
- [x] DigitalOcean
- [x] Google
- [ ] InMemory

PRs welcome!

Notes
=====
When the `external-dns.alpha.kubernetes.io/ttl` annotation is not provided, the TTL will default to 0 seconds and `enpoint.TTL.isConfigured()` will be false.

### AWS Provider
The AWS Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

### DigitalOcean Provider
The DigitalOcean Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

### Google Provider
Previously with the Google Provider, TTL's were hard-coded to 300s.
For safety, the Google Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

For the moment, it is impossible to use a TTL value of 0 with the AWS, DigitalOcean, or Google Providers.
This behavior may change in the future.
