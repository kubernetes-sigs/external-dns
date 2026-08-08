# Exoscale

## Prerequisites

Exoscale provider support was added via [this PR](https://github.com/kubernetes-sigs/external-dns/pull/625), thus you need to use external-dns v0.5.5.

The Exoscale provider expects that your Exoscale zones, you wish to add records to, already exists
and are configured correctly. It does not add, remove or configure new zones in anyway.

To do this please refer to the [Exoscale DNS documentation](https://community.exoscale.com/documentation/dns/).

Additionally you will have to provide the Exoscale...:

* API Key
* API Secret
* Elastic IP address, to access the workers

## Deployment

Deploying external DNS for Exoscale is actually nearly identical to deploying
it for other providers. This is what a sample `deployment.yaml` looks like:

```yaml
[[% include 'exoscale/extdns.yaml' %]]
```

Optional arguments `--exoscale-apizone` and `--exoscale-apienv` define [Exoscale API Zone](https://community.exoscale.com/documentation/platform/exoscale-datacenter-zones/)
(default `ch-gva-2`) and Exoscale API environment (default `api`, can be used to target non-production API server) respectively.

## RBAC

If your cluster is RBAC enabled, you also need to setup the following, before you can run external-dns:

```yaml
[[% include 'exoscale/rbac.yaml' %]]
```

## Testing and Verification

**Important!**: Remember to change `example.com` with your own domain throughout the following text.

Spin up a simple nginx HTTP server with the following spec (`kubectl apply -f`):

```yaml
[[% include 'exoscale/how-to-test.yaml' %]]
```

**Important!**: Don't run dig, nslookup or similar immediately (until you've
confirmed the record exists). You'll get hit by [negative DNS caching](https://tools.ietf.org/html/rfc2308), which is hard to flush.

Wait about 30s-1m (interval for external-dns to kick in), then check Exoscales [portal](https://portal.exoscale.com/dns/example.com)... via-ingress.example.com should appear as a A and TXT record with your Elastic-IP-address.
