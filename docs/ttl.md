Configure DNS record TTL (Time-To-Live)
=======================================

An optional annotation `external-dns.alpha.kubernetes.io/ttl` is available to customize the TTL value of a DNS record.
TTL is specified as an integer encoded as string representing seconds.

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

TTL can also be specified as a duration value parsable by Golang [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration):

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com.
    external-dns.alpha.kubernetes.io/ttl: "1m"
  ...
```

Both examples result in the same value of 60 seconds TTL.

TTL must be a positive value.

Providers
=========

- [x] AWS (Route53)
- [x] Azure
- [ ] Cloudflare
- [x] DigitalOcean
- [x] DNSimple
- [x] Google
- [ ] InMemory
- [x] Linode
- [x] TransIP
- [x] RFC2136
- [x] Vultr
- [x] UltraDNS

PRs welcome!

Notes
=====
When the `external-dns.alpha.kubernetes.io/ttl` annotation is not provided, the TTL will default to 0 seconds and `endpoint.TTL.isConfigured()` will be false.

### AWS Provider
The AWS Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

## Azure
TTL value should be between 1 and 2,147,483,647 seconds.
By default it will be 300s.

### DigitalOcean Provider
The DigitalOcean Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

### DNSimple Provider
The DNSimple Provider default TTL is used when the TTL is 0. The default TTL is 3600s.

### Google Provider
Previously with the Google Provider, TTL's were hard-coded to 300s.
For safety, the Google Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

For the moment, it is impossible to use a TTL value of 0 with the AWS, DigitalOcean, or Google Providers.
This behavior may change in the future.

### Linode Provider
The Linode Provider default TTL is used when the TTL is 0. The default is 24 hours

### TransIP Provider
The TransIP Provider minimal TTL is used when the TTL is 0. The minimal TTL is 60s.

### Vultr Provider
The Vultr provider minimal TTL is used when the TTL is 0. The default is 1 hour.

### UltraDNS
The UltraDNS provider minimal TTL is used when the TTL is not provided. The default TTL is account level default TTL, if defined, otherwise 24 hours.
