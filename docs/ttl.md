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
- [ ] DigitalOcean
- [ ] Google
- [ ] InMemory

PRs welcome!
