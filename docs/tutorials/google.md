# Google Cloud DNS

## Annotations

Annotations which are specific to the Google
[CloudDNS](https://cloud.google.com/dns/docs/overview) provider.

### Routing Policy

The [routing policy](https://cloud.google.com/dns/docs/routing-policies-overview)
for resource record sets managed by ExternalDNS may be specified by applying the
`external-dns.alpha.kubernetes.io/google-routing-policy` annotation on any of the
supported [sources](../sources/about.md).

#### Geolocation routing policies

Specifying a value of `geo` for the `external-dns.alpha.kubernetes.io/google-routing-policy`
annotation will enable geolocation routing for associated resource record sets. The
location attributed to resource record sets may be deduced for instances of ExternalDNS
running within the Google Cloud platform or may be specified via the `--google-location`
command-line argument. Alternatively, a location may be explicitly specified via the
`external-dns.alpha.kubernetes.io/google-location` annotation, where the value is one
of Google Cloud's [locations/regions](https://cloud.google.com/docs/geography-and-regions).

For example:
```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: geo-example
  annotations:
    external-dns.alpha.kubernetes.io/google-routing-policy: "geo"
    external-dns.alpha.kubernetes.io/google-location: "us-east1"
```

#### Weighted Round Robin routing policies

Specifying a value of `wrr` for the `external-dns.alpha.kubernetes.io/google-routing-policy`
annotation will enable weighted round-robin routing for associated resource record sets.
The weight to be attributed to resource record sets may be specified via the
`external-dns.alpha.kubernetes.io/google-weight` annotation, where the value is a string
representation of a floating-point number. The `external-dns.alpha.kubernetes.io/set-identifier`
annotation must also be applied providing a string value representation of an index into
the list of potential responses.

For example:
```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: wrr-example
  annotations:
    external-dns.alpha.kubernetes.io/google-routing-policy: "wrr"
    external-dns.alpha.kubernetes.io/google-weight: "100.0"
    external-dns.alpha.kubernetes.io/set-identifier: "0"
```