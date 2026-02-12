# Annotations

ExternalDNS sources support a number of annotations on the Kubernetes resources that they examine.

The following table documents which sources support which annotations:

| Source       | controller | hostname | internal-hostname | target  | ttl     | (provider-specific) |
|--------------|------------|----------|-------------------|---------|---------|---------------------|
| Ambassador   |            |          |                   | Yes     | Yes     | Yes                 |
| Connector    |            |          |                   |         |         |                     |
| Contour      | Yes        | Yes[^1]  |                   | Yes     | Yes     | Yes                 |
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
[^4]: For Gateway API sources, annotation placement differs by type. See [Gateway API Annotation Placement](#gateway-api-annotation-placement) for details.
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

Specifies which set of addresses to use for a [`headless Service`](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services).

Supported values:

- `NodeExternalIP`. Required `--service-type-filter=ClusterIP` and `--service-type-filter=Node` or no `--service-type-filter` flag specified.
- `HostIP`.

If the value is `NodeExternalIP`, use each relevant `Pod`'s `Node`'s address of type `ExternalIP`
plus each IPv6 address of type `InternalIP`.

Otherwise, if the value is `HostIP` or the `--publish-host-ip` flag is specified, use
each relevant `Pod`'s `Status.HostIP`.

Otherwise, use the `IP` of each of the `Service`'s `Endpoints`'s `Addresses`.

## external-dns.alpha.kubernetes.io/hostname

Specifies additional domains for the resource's DNS records.

Multiple hostnames can be specified through a comma-separated list, e.g.
`svc.mydomain1.com,svc.mydomain2.com`.

For `Pods`, uses the `Pod`'s `Status.PodIP`, unless they are `hostNetwork: true` in which case the NodeExternalIP is used for IPv4 and NodeInternalIP for IPv6.

Notes:

- This annotation can override or add extra hostnames alongside any automatically derived hostnames (e.g., from Ingress.spec.rules[].host).
- The [`ingress-hostname-source`](#external-dnsalphakubernetesioingress-hostname-source) annotation may be used to specify where to get the domain for an `Ingress` resource.
- Hostnames must match the domain filter set in ExternalDNS (e.g., --domain-filter=example.com).
- This is an alpha annotation — subject to change; newer versions may support alternatives or deprecate it.
- This annotation is helpful for:
  - Services or other resources without native hostname fields.
  - Explicit overrides or multi-host situations.
  - Avoiding reliance on auto-detection or heuristics.

### Use Cases for `external-dns.alpha.kubernetes.io/hostname` annotation

#### Explicit Hostname Mapping for Services

You have a Service (e.g. of type LoadBalancer or ClusterIP) and want to expose it under a custom DNS name:

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  annotations:
    external-dns.alpha.kubernetes.io/hostname: app.example.com
spec:
  type: LoadBalancer
  ...
```

> ExternalDNS will create a A or CNAME record for app.example.com pointing to the external IP or hostname of the service.

#### Multi-Hostname Records

You can assign multiple hostnames by separating them with commas:

```yml
annotations:
  external-dns.alpha.kubernetes.io/hostname: api.example.com,api.internal.example.com
```

> ExternalDNS will create two DNS records for the same service.

#### Static DNS Assignment Without Ingress Rules

When using Ingress, you usually declare hostnames in the spec.rules[].host. But with this annotation, you can manage DNS independently:

```yml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    external-dns.alpha.kubernetes.io/hostname: www.example.com
spec:
  rules:
    - http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: my-service
                port:
                  number: 80
```

> Useful when DNS management is decoupled from routing logic.

## external-dns.alpha.kubernetes.io/ingress-hostname-source

Specifies where to get the domain for an `Ingress` resource.

If the value is `defined-hosts-only`, use only the domains from the `Ingress` spec.

If the value is `annotation-only`, use only the domains from the `Ingress` annotations.

If the annotation is not present, use the domains from both the spec and annotations.

## external-dns.alpha.kubernetes.io/ingress

This annotation allows ExternalDNS to work with Istio & GlooEdge Gateways that don't have a public IP.

It can be used to address a specific architectural pattern, when a Kubernetes Ingress directs all public traffic to an Istio or GlooEdge Gateway:

- **The Challenge**: By default, ExternalDNS sources the public IP address for a DNS record from a Service of type LoadBalancer.
However, in some setups, the Gateway's Service is of type ClusterIP, with all public traffic routed to it via a separate Kubernetes Ingress object. This setup leaves the Gateway without a public IP that ExternalDNS can discover.

- **The Solution**: The annotation on the Istio/GlooEdge Gateway tells ExternalDNS to ignore the Gateway's Service IP. Instead, it directs ExternalDNS to a specified Ingress resource to find the target LoadBalancer IP address.

### Use Cases for `external-dns.alpha.kubernetes.io/ingress` annotation

#### Getting target from Ingress backed Gloo Gateway

```yml
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/ingress: gateway-proxy
  labels:
    app: gloo
  name: gateway-proxy
  namespace: gloo-system
spec:
  bindAddress: '::'
  bindPort: 8080
  options: {}
  proxyNames:
  - gateway-proxy
  ssl: false
  useProxyProto: false
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: gateway-proxy
  namespace: gloo-system
spec:
  ingressClassName: alb
  rules:
  - host: cool-service.example.com
    http:
      paths:
      - backend:
          service:
            name: gateway-proxy
            port:
              name: http
        path: /
        pathType: Prefix
status:
  loadBalancer:
    ingress:
    - hostname: k8s-alb-c4aa37c880-740590208.us-east-1.elb.amazonaws.com
---
# This object is generated by GlooEdge Control Plane from Gateway and VirtualService.
# We have no direct control on this resource
apiVersion: gloo.solo.io/v1
kind: Proxy
metadata:
  labels:
    created_by: gloo-gateway
  name: gateway-proxy
  namespace: gloo-system
spec:
  listeners:
  - bindAddress: '::'
    bindPort: 8080
    httpListener:
      virtualHosts:
      - domains:
        - cool-service.example.com
        metadataStatic:
          sources:
          - observedGeneration: "6652"
            resourceKind: '*v1.VirtualService'
            resourceRef:
              name: cool-service
              namespace: gloo-system
        name: cool-service
        routes:
        - matchers:
          - prefix: /
          metadataStatic:
            sources:
            - observedGeneration: "6652"
              resourceKind: '*v1.VirtualService'
              resourceRef:
                name: cool-service
                namespace: gloo-system
            upgrades:
            - websocket: {}
    metadataStatic:
      sources:
      - observedGeneration: "6111"
        resourceKind: '*v1.Gateway'
        resourceRef:
          name: gateway-proxy
          namespace: gloo-system
    name: listener-::-8080
    useProxyProto: false
```

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
It must be between `1` and `2,147,483,647` seconds.

> Note; setting the value to `0` means, that TTL is not configured and thus use default.

## external-dns.alpha.kubernetes.io/gateway-hostname-source

Specifies where to get the domain for a `Route` resource. This annotation should be present on the actual `Route` resource, not the `Gateway` resource itself.

If the value is `defined-hosts-only`, use only the domains from the `Route` spec.

If the value is `annotation-only`, use only the domains from the `Route` annotations.

If the annotation is not present, use the domains from both the spec and annotations.

## Provider-specific annotations

Some providers define their own annotations. Cloud-specific annotations have keys prefixed as follows:

| Cloud      | Annotation prefix                              |
|------------|------------------------------------------------|
| AWS        | `external-dns.alpha.kubernetes.io/aws-`        |
| CloudFlare | `external-dns.alpha.kubernetes.io/cloudflare-` |
| Scaleway   | `external-dns.alpha.kubernetes.io/scw-`        |

Additional annotations implemented by specific providers:

### external-dns.alpha.kubernetes.io/alias

If the value of this annotation is `true`, specifies that CNAME records generated by the
resource should instead be alias records.

**Supported providers:**

- **AWS**: This annotation is only relevant if the `--aws-prefer-cname` flag is specified.
- **PowerDNS**: When this annotation is set to `true`, CNAME records will be created as ALIAS records.
  This is useful when using PowerDNS with `expand-alias=yes` to resolve CNAME targets to IP addresses
  on the authoritative server side. Alternatively, use the `--prefer-alias` flag to convert all
  CNAME records to ALIAS globally.

### external-dns.alpha.kubernetes.io/set-identifier

Specifies the set identifier for DNS records generated by the resource.

A set identifier differentiates among multiple DNS record sets that have the same combination of domain and type.
Which record set or sets are returned to queries is then determined by the configured routing policy.

## Gateway API Annotation Placement

When using Gateway API sources (`gateway-httproute`, `gateway-grpcroute`, `gateway-tlsroute`, etc.), annotations
are read from different resources: **Gateway resource** reads only `target` annotation, while **Route resources**
(HTTPRoute, GRPCRoute, TLSRoute, etc.) read all other annotations (`hostname`, `ttl`, `controller`, and
provider-specific annotations like `cloudflare-*`, `aws-*`, `scw-*`).

For more details and comprehensive examples, see the
[Gateway API documentation](../sources/gateway-api.md#annotations).
