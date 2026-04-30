# Gateway API Route Sources

This describes how to configure ExternalDNS to use Gateway API Route sources.
It is meant to supplement the other provider-specific setup tutorials.

## Supported API Versions

ExternalDNS uses Gateway API CRDs, which are distributed at different versions in Standard and/or
Experimental channels as summarized below:

|      Resource      | API Version Used<br> by ExternalDNS | Mininmum Standard<br>Release Channel | Experimental<br>Release Channel |
|--------------------|-------------------------------------|--------------------------------------|---------------------------------|
| Gateway            | v1                                  | v1.0.0                               | v1.0.0                          |
| HTTPRoute          | v1                                  | v1.0.0                               | v1.0.0                          |
| GRPCRoute          | v1                                  | v1.1.0                               | v1.1.0                          |
| ListenerSet        | v1                                  | v1.5.0                               | v1.5.0                          |
| TLSRoute           | v1                                  | v1.5.0                               | v1.0.0                          |
| TCPRoute           | v1alpha2                            | TBD                                  | v1.0.0                          |
| UDPRoute           | v1alpha2                            | TBD                                  | v1.0.0                          |

Gateways and HTTPRoutes were promoted to the Standard channel in Gateway API v1.0.0 and use the
v1 API.

GRPCRoutes were promoted to the Standard channel in Gateway API v1.1.0 and use the
v1 API.

ListenerSets were promoted to the Standard channel in Gateway API v1.5.0.
They use the v1 API and allow attaching additional listeners to an existing Gateway.
Routes that reference a ListenerSet as a parentRef are automatically supported —
ExternalDNS follows the ListenerSet to its parent Gateway to resolve target addresses.
The `external-dns.alpha.kubernetes.io/target` annotation is also supported on ListenerSet
resources. When present, it takes precedence over the parent Gateway's target annotation.
ListenerSet support requires the `--gateway-listener-sets` flag to be enabled.

TLSRoutes were promoted to the Standard channel in Gateway API v1.5.0 but have been
available in the Experimental channel as v1alpha2 since v1.0.0.
ExternalDNS still uses the v1alpha2 API for compatibility with older CRDs but it
has been deprecated and will be removed from future releases, at which point ExternalDNS will
need to migrate to v1. (See [#6247](https://github.com/kubernetes-sigs/external-dns/issues/6247))

TCPRoute and UDPRoute remain experimental and are only available as v1alpha2 in the Experimental channel.

## Hostnames

HTTPRoute and TLSRoute specs, along with their associated Gateway Listeners, contain hostnames that
will be used by ExternalDNS. However, no such hostnames may be specified in TCPRoute or UDPRoute
specs. For TCPRoutes and UDPRoutes, the `external-dns.alpha.kubernetes.io/hostname` annotation
is the recommended way to provide their hostnames to ExternalDNS. This annotation is also supported
for HTTPRoutes and TLSRoutes by ExternalDNS, but it's _strongly_ recommended that they use their
specs to provide all intended hostnames, since the Gateway that ultimately routes their
requests/connections won't recognize additional hostnames from the annotation.

## Annotations

### Annotation Placement

ExternalDNS reads different annotations from different Gateway API resources:

- **Gateway annotations**: Only `external-dns.alpha.kubernetes.io/target` is read from Gateway resources
- **ListenerSet annotations**: The `external-dns.alpha.kubernetes.io/target` annotation is also supported on
  ListenerSet resources. When a Route references a ListenerSet, the ListenerSet target annotation takes
  precedence over the parent Gateway's target annotation. Requires `--gateway-listener-sets`.
- **Route annotations**: All other annotations (hostname, ttl, controller, provider-specific) are read from Route
  resources (HTTPRoute, GRPCRoute, TLSRoute, TCPRoute, UDPRoute)

This separation aligns with Gateway API architecture where Gateway defines infrastructure (IP addresses, listeners)
and Routes define application-level DNS records.

### Examples

#### Example: Cloudflare Proxied Records

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
  namespace: default
  annotations:
    # ✅ Correct: target annotation on Gateway
    external-dns.alpha.kubernetes.io/target: "203.0.113.1"
spec:
  gatewayClassName: cilium
  listeners:
    - name: https
      hostname: "*.example.com"
      protocol: HTTPS
      port: 443
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: my-route
  annotations:
    # ✅ Correct: provider-specific annotations on HTTPRoute
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
    external-dns.alpha.kubernetes.io/ttl: "300"
spec:
  parentRefs:
    - name: my-gateway
      namespace: default
  hostnames:
    - api.example.com
  rules:
    - backendRefs:
        - name: api-service
          port: 8080
```

#### Example: AWS Route53 with Routing Policies

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: aws-gateway
  annotations:
    # ✅ Correct: target annotation on Gateway
    external-dns.alpha.kubernetes.io/target: "alb-123.us-east-1.elb.amazonaws.com"
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: weighted-route
  annotations:
    # ✅ Correct: AWS-specific annotations on HTTPRoute
    external-dns.alpha.kubernetes.io/aws-weight: "100"
    external-dns.alpha.kubernetes.io/set-identifier: "backend-v1"
spec:
  parentRefs:
    - name: aws-gateway
  hostnames:
    - app.example.com
```

### Common Mistakes

❌ **Incorrect**: Placing provider-specific annotations on Gateway

```yaml
kind: Gateway
metadata:
  annotations:
    # ❌ These annotations are ignored on Gateway
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
    external-dns.alpha.kubernetes.io/ttl: "300"
```

❌ **Incorrect**: Placing target annotation on HTTPRoute

```yaml
kind: HTTPRoute
metadata:
  annotations:
    # ❌ This annotation is ignored on Routes
    external-dns.alpha.kubernetes.io/target: "203.0.113.1"
```

### external-dns.alpha.kubernetes.io/gateway-hostname-source

**Why is this needed:**
In certain scenarios, conflicting DNS records can arise when External DNS processes both the hostname annotations and the hostnames defined in the `*Route` spec. For example:

- A CNAME record (`company.public.example.com -> company.private.example.com`) is used to direct traffic to private endpoints (e.g., AWS PrivateLink).
- Some third-party services require traffic to resolve publicly to the Gateway API load balancer, but the hostname (`company.public.example.com`) must remain unchanged to avoid breaking the CNAME setup.
- Without this annotation, External DNS may override the CNAME record with an A record due to conflicting hostname definitions.

**Usage:**
By setting the annotation `external-dns.alpha.kubernetes.io/gateway-hostname-source: annotation-only`, users can instruct External DNS
to ignore hostnames defined in the `HTTPRoute` spec and use only the hostnames specified in annotations. This ensures
compatibility with complex DNS configurations and avoids record conflicts.

**Example:**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/gateway-hostname-source: annotation-only
    external-dns.alpha.kubernetes.io/hostname: company.private.example.com
spec:
  hostnames:
    - company.public.example.com
```

In this example, External DNS will create DNS records only for `company.private.example.com` based on the annotation, ignoring the `hostnames` field in the `HTTPRoute` spec. This prevents conflicts with existing CNAME records while enabling public resolution for specific endpoints.

For a complete list of supported annotations, see the
[annotations documentation](../annotations/annotations.md#gateway-api-annotation-placement).

## Manifest with RBAC

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["get","watch","list"]
- apiGroups: ["gateway.networking.k8s.io"]
  resources: ["gateways","httproutes","grpcroutes","tlsroutes","tcproutes","udproutes","listenersets"]
  verbs: ["get","watch","list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
- kind: ServiceAccount
  name: external-dns
  namespace: default
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: default
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: external-dns
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.21.0
        args:
        # Add desired Gateway API Route sources.
        - --source=gateway-httproute
        - --source=gateway-grpcroute
        - --source=gateway-tlsroute
        - --source=gateway-tcproute
        - --source=gateway-udproute
        # Optionally, limit Routes to those in the given namespace.
        - --namespace=my-route-namespace
        # Optionally, limit Routes to those matching the given label selector.
        - --label-filter=my-route-label==my-route-value
        # Optionally, limit Route endpoints to those Gateways with the given name.
        - --gateway-name=my-gateway-name
        # Optionally, limit Route endpoints to those Gateways in the given namespace.
        - --gateway-namespace=my-gateway-namespace
        # Optionally, limit Route endpoints to those Gateways matching the given label selector.
        - --gateway-label-filter=my-gateway-label==my-gateway-value
        # Optionally, enable ListenerSet support for Routes referencing ListenerSet parentRefs.
        - --gateway-listener-sets
        # Add provider-specific flags...
        - --domain-filter=external-dns-test.my-org.com
        - --provider=google
        - --registry=txt
        - --txt-owner-id=my-identifier
```
