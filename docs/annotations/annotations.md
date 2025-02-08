# Annotations

ExternalDNS sources support a number of annotations on the Kubernetes resources that they examine.

The following table documents which sources support which annotations:

| Source       | controller | hostname | internal-hostname | target  | ttl     | (provider-specific) |
|--------------|------------|----------|-------------------|---------|---------|---------------------|
| Ambassador   |            |          |                   | Yes     | Yes     | Yes                 |
| Connector    |            |          |                   |         |         |                     |
| Contour      | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| CloudFoundry |            |          |                   |         |         |                     |
| CRD          |            |          |                   |         |         |                     |
| F5           |            |          |                   | Yes     | Yes     |                     |
| Gateway      | Yes        | Yes[^1]  |                   | Yes[^4] | Yes     | Yes                 |
| Gloo         |            |          |                   | Yes     | Yes[^5] | Yes[^5]             |
| Ingress      | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| Istio        | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| Kong         |            | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| Node         | Yes        |          |                   | Yes     | Yes     |                     |
| OpenShift    | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| Pod          |            | Yes      | Yes               | Yes     |         |                     |
| Service      | Yes        | Yes[^1]  | Yes[^1][^2]       | Yes[^3] | Yes     | Yes                 |
| Skipper      | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
| Traefik      |            | Yes[^1]  |                   | Yes     | Yes     | Yes                 |

[^1]: Unless the `--ignore-hostname-annotation` flag is specified.
[^2]: Only behaves differently than `hostname` for `Service`s of type `ClusterIP` or `LoadBalancer`.
[^3]: Also supported on `Pods` referenced from a headless `Service`'s `Endpoints`.
[^4]: The annotation must be on the `Gateway`.
[^5]: The annotation must be on the listener's `VirtualService`.

## external-dns.alpha.kubernetes.io/access

Specifies which set of node IP addresses to use for a `Service` of type `NodePort`.

If the value is `public`, use the Nodes' addresses of type `ExternalIP`, plus IPv6 addresses of type `InternalIP`.

If the value is `private`, use the Nodes' addresses of type `InternalIP`.

If the annotation is not present and there is at least one address of type `ExternalIP`,
behave as if the value were `public`, otherwise behave as if the value were `private`.

## external-dns.alpha.kubernetes.io/controller

If this annotation exists and has a value other than `dns-controller` then the source ignores the resource.

## external-dns.alpha.kubernetes.io/endpoints-type

Specifies which set of addresses to use for a headless `Service`.

If the value is `NodeExternalIP`, use each relevant `Pod`'s `Node`'s address of type `ExternalIP`
plus each IPv6 address of type `InternalIP`.

Otherwise, if the value is `HostIP` or the `--publish-host-ip` flag is specified, use
each relevant `Pod`'s `Status.HostIP`.

Otherwise, use the `IP` of each of the `Service`'s `Endpoints`'s `Addresses`.

## external-dns.alpha.kubernetes.io/hostname

Specifies the domain for the resource's DNS records.

Multiple hostnames can be specified through a comma-separated list, e.g.
`svc.mydomain1.com,svc.mydomain2.com`.

For `Pods`, uses the `Pod`'s `Status.PodIP`, unless they are `hostNetwork: true` in which case the NodeExternalIP is used for IPv4 and NodeInternalIP for IPv6.

## external-dns.alpha.kubernetes.io/ingress-hostname-source

Specifies where to get the domain for an `Ingress` resource.

If the value is `defined-hosts-only`, use only the domains from the `Ingress` spec.

If the value is `annotation-only`, use only the domains from the `Ingress` annotations.

If the annotation is not present, use the domains from both the spec and annotations.

## external-dns.alpha.kubernetes.io/internal-hostname

Specifies the domain for the resource's DNS records that are for use from internal networks.

For `Services` of type `LoadBalancer`, uses the `Service`'s `ClusterIP`.

For `Pods`, uses the `Pod`'s `Status.PodIP`, unless they are `hostNetwork: true` in which case the NodeExternalIP is used for IPv4 and NodeInternalIP for IPv6.

## external-dns.alpha.kubernetes.io/target

Specifies a comma-separated list of values to override the resource's DNS record targets (RDATA).

Targets that parse as IPv4 addresses are published as A records and
targets that parse as IPv6 addresses are published as AAAA records. All other targets
are published as CNAME records.

## external-dns.alpha.kubernetes.io/ttl

Specifies the TTL (time to live) for the resource's DNS records.

The value may be specified as either a duration or an integer number of seconds.
It must be between 1 and 2,147,483,647 seconds.

## Provider-specific annotations

Some providers define their own annotations. Cloud-specific annotations have keys prefixed as follows:

| Cloud      | Annotation prefix                              |
|------------|------------------------------------------------|
| AWS        | `external-dns.alpha.kubernetes.io/aws-`        |
| CloudFlare | `external-dns.alpha.kubernetes.io/cloudflare-` |
| IBM Cloud  | `external-dns.alpha.kubernetes.io/ibmcloud-`   |
| Scaleway   | `external-dns.alpha.kubernetes.io/scw-`        |

Additional annotations that are currently implemented only by AWS are:

### external-dns.alpha.kubernetes.io/alias

If the value of this annotation is `true`, specifies that CNAME records generated by the
resource should instead be alias records.

This annotation is only relevant if the `--aws-prefer-cname` flag is specified.

### external-dns.alpha.kubernetes.io/set-identifier

Specifies the set identifier for DNS records generated by the resource.

A set identifier differentiates among multiple DNS record sets that have the same combination of domain and type.
Which record set or sets are returned to queries is then determined by the configured routing policy.
