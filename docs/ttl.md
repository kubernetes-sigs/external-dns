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

Provider TTL Defaults
=========

- [AWS (Route53)](#aws-route53)
- [Azure](#azure)
- [DigitalOcean](#digitalocean)
- [DNSimple](#dnsimple)
- [Google](#google)
- [Linode](#linode)
- [Stackpath](#stackpath)
- [TransIP](#transip)
- [UltraDNS](#ultradns)
- [Vultr](#vultr)

PRs welcome!

## Default
When the `external-dns.alpha.kubernetes.io/ttl` annotation is not provided, the TTL will default to 0 seconds and `endpoint.TTL.isConfigured()` will be false.
***
## AWS Route53
The AWS Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the Provider code.
***
## Azure
TTL value should be between 1 and 2,147,483,647 seconds.
By default it will be 300s.
***
## DigitalOcean
The DigitalOcean Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the Provider code.
***
## DNSimple
The DNSimple Provider default TTL is used when the TTL is 0. The default TTL is 3600s.
***
## Google
Previously with the Google Provider, TTL's were hard-coded to 300s.
For safety, the Google Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the Provider code.

For the moment, it is impossible to use a TTL value of 0 with the AWS, DigitalOcean, or Google Providers.
This behavior may change in the future.
***
## Linode
The Linode Provider default TTL is used when the TTL is 0. The default is 24 hours
***
## Stackpath
The Stackpath Provider default TTL is used when the TTL is undefined in the source annotation. The default TTL used is 2 minutes (120 seconds).
***
## TransIP
The TransIP Provider minimal TTL is used when the TTL is 0. The minimal TTL is 60s.
***
## UltraDNS
The UltraDNS Provider minimal TTL is used when the TTL is not provided. The default TTL is account level default TTL, if defined, otherwise 24 hours.
***
## Vultr
The Vultr Provider minimal TTL is used when the TTL is 0. The default is 1 hour.