# Configure DNS record TTL (Time-To-Live)

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

## TTL annotation support

> Note: For TTL annotations to work, the `external-dns.alpha.kubernetes.io/hostname` annotation must be set on the resource and be supported by the provider as well as the source.

### Providers

| Provider       | Supported |
|:---------------|:---------:|
| `Akamai`       |     ✅     |
| `AlibabaCloud` |     ✅     |
| `AWS`          |     ✅     |
| `AWSSD`        |     ✅     |
| `Azure`        |     ✅     |
| `Civo`         |     ❌     |
| `Cloudflare`   |     ✅     |
| `CoreDNS`      |     ❌     |
| `DigitalOcean` |     ✅     |
| `DNSSimple`    |     ✅     |
| `Exoscale`     |     ✅     |
| `Gandi`        |     ✅     |
| `GoDaddy`      |     ✅     |
| `Google GCP`   |     ✅     |
| `InMemory`     |     ❌     |
| `Linode`       |     ❌     |
| `NS1`          |     ❌     |
| `OCI`          |     ✅     |
| `OVH`          |     ❌     |
| `PDNS`         |     ❌     |
| `PiHole`       |     ✅     |
| `Plural`       |     ❌     |
| `RFC2136`      |     ✅     |
| `Scaleway`     |     ✅     |
| `Transip`      |     ✅     |
| `Webhook`      |     ✅     |

### Sources

| Source                 | Supported |
|:-----------------------|:---------:|
| `ambassador-host`      |     ✅     |
| `cloudfoundry`         |     ❌     |
| `connector`            |     ❌     |
| `contour-httpproxy`    |     ✅     |
| `crd`                  |     ❌     |
| `empty`                |     ❌     |
| `f5-transportserver`   |     ✅     |
| `f5-virtualserver`     |     ✅     |
| `fake`                 |     ❌     |
| `gateway-grpcroute`    |     ✅     |
| `gateway-httproute`    |     ✅     |
| `gateway-tcproute`     |     ✅     |
| `gateway-tlsroute`     |     ✅     |
| `gateway-udproute`     |     ✅     |
| `gloo-proxy`           |     ✅     |
| `ingress`              |     ✅     |
| `istio-gateway`        |     ✅     |
| `istio-virtualservice` |     ✅     |
| `kong-tcpingress`      |     ✅     |
| `node`                 |     ✅     |
| `openshift-route`      |     ✅     |
| `pod`                  |     ✅     |
| `service`              |     ✅     |
| `skipper-routegroup`   |     ✅     |
| `traefik-proxy`        |     ✅     |

## Notes

When the `external-dns.alpha.kubernetes.io/ttl` annotation is not provided, the TTL will default to 0 seconds and `endpoint.TTL.isConfigured()` will be false.

### AWS Provider

The AWS Provider overrides the value to 300s when the TTL is 0.
This value is a constant in the provider code.

### Azure

TTL value should be between 1 and 2,147,483,647 seconds.
By default it will be 300s.

### CloudFlare Provider

CloudFlare overrides the value to "auto" when the TTL is 0.

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

## Use Cases for `external-dns.alpha.kubernetes.io/ttl` annotation

The `external-dns.alpha.kubernetes.io/ttl` annotation allows you to set a custom **TTL (Time To Live)** for DNS records managed by `external-dns`.

Use the `external-dns.alpha.kubernetes.io/tt` annotation to fine-tune DNS caching behavior per record, balancing between update frequency and performance.

This is useful in several real-world scenarios depending on how frequently DNS records are expected to change.

---

### Fast Failover for Critical Services

For services that must be highly available—like APIs, databases, or external load balancers—set a **low TTL** (e.g., 30 seconds) so DNS clients quickly update to new IPs during:

- Node failures
- Region failovers
- Blue/green deployments

```yaml
annotations:
  external-dns.alpha.kubernetes.io/ttl: "30s"
```

---

### Long TTL for Static Services

If your service’s IP or endpoint rarely changes (e.g., static websites, internal dashboards), you can set a long TTL (e.g., 86400 seconds = 24 hours) to:

- Reduce DNS query load
- Improve cache performance
- Lower cost with some DNS providers

```yml
annotations:
  external-dns.alpha.kubernetes.io/ttl: "24h"
```

---

### Canary or Experimental Services

Use a short TTL for services under test or experimentation to allow fast DNS propagation when making changes, allowing easy rollback and testing.

---

### Provider-Specific Optimization

Some DNS providers charge per query or have query rate limits. Adjusting the TTL lets you:

- Reduce costs
- Avoid throttling
- Manage DNS traffic load efficiently

---

### Regulatory or Contractual SLAs

Certain environments may require TTL values to align with:

- Regulatory guidelines
- Legacy system compatibility
- Contractual service-level agreements

---

### Autoscaling Node Pools in GCP (or Other Cloud Providers)

In environments like Google Cloud Platform (GCP) using private node IPs for DNS resolution, ExternalDNS may register node IPs with a default TTL of 300 seconds.

During autoscaling events (e.g., node addition/removal or upgrades), DNS records may remain stale for several minutes, causing traffic to be routed to non-existent nodes.

By using the TTL annotation you can:

- Reduce TTL to allow faster DNS propagation
- Ensure quicker routing updates when node IPs change
- Improve resiliency during frequent cluster topology changes
- Fine-grained TTL control helps avoid downtime or misrouting in dynamic, autoscaling environments.
