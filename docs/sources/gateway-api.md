# Gateway API Route Sources

This describes how to configure ExternalDNS to use Gateway API Route sources.
It is meant to supplement the other provider-specific setup tutorials.

## Supported API Versions

ExternalDNS currently supports a mixture of v1alpha2, v1beta1, v1 APIs.

Gateway API has two release channels: Standard and Experimental.
The Experimental channel includes v1alpha2, v1beta2, and v1 APIs.
The Standard channel only includes v1beta2 and v1 APIs, not v1alpha2.

TCPRoutes, TLSRoutes, and UDPRoutes only exist in v1alpha2 and continued support for
these versions is NOT guaranteed. At some time in the future, Gateway API will graduate
these Routes to v1. ExternalDNS will likely follow that upgrade and move to the v1 API,
where they will be available in the Standard release channel. This will be a breaking
change if your Experimental CRDs are not updated to include the new v1 API.

Gateways and HTTPRoutes are available in v1alpha2, v1beta1, and v1 APIs.
However, some notable environments are behind in upgrading their CRDs to include the v1 API.
For compatibility reasons Gateways and HTTPRoutes use the v1beta1 API.

GRPCRoutes are available in v1alpha2 and v1 APIs, not v1beta2.
Therefore, GRPCRoutes use the v1 API which is available in both release channels.
Unfortunately, this means they will not be available in environments with old CRDs.

## Hostnames

HTTPRoute and TLSRoute specs, along with their associated Gateway Listeners, contain hostnames that
will be used by ExternalDNS. However, no such hostnames may be specified in TCPRoute or UDPRoute
specs. For TCPRoutes and UDPRoutes, the `external-dns.alpha.kubernetes.io/hostname` annotation
is the recommended way to provide their hostnames to ExternalDNS. This annotation is also supported
for HTTPRoutes and TLSRoutes by ExternalDNS, but it's _strongly_ recommended that they use their
specs to provide all intended hostnames, since the Gateway that ultimately routes their
requests/connections won't recognize additional hostnames from the annotation.

## Annotations

### Annotation Inheritance

ExternalDNS supports annotation inheritance from Gateway to Route resources. This allows you to:

- Set **default annotations** on Gateway that apply to all attached Routes
- **Override** specific annotations on individual Routes when needed

**Inheritance rules:**

1. Gateway annotations serve as **defaults** for all Routes attached to that Gateway
2. Route annotations **override** Gateway annotations when both specify the same key
3. Each Route is independent — one Route's annotations don't affect other Routes

This approach reduces configuration duplication while maintaining flexibility for per-Route customization.

### Examples

#### Example: Centralized Defaults with Per-Route Overrides

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
  namespace: default
  annotations:
    # Default target for all Routes (intranet)
    external-dns.alpha.kubernetes.io/target: "172.16.6.6"
    # Default TTL for all Routes
    external-dns.alpha.kubernetes.io/ttl: "300"
    # Default Cloudflare proxy setting
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
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
  name: internal-api
  # No annotations — inherits all defaults from Gateway:
  # target=172.16.6.6, ttl=300, cloudflare-proxied=true
spec:
  parentRefs:
    - name: my-gateway
  hostnames:
    - api.internal.example.com
  rules:
    - backendRefs:
        - name: api-service
          port: 8080
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: external-api
  annotations:
    # Override: use external IP instead of intranet
    external-dns.alpha.kubernetes.io/target: "1.2.3.4"
    # Override: shorter TTL for external endpoint
    external-dns.alpha.kubernetes.io/ttl: "60"
    # Inherits cloudflare-proxied=true from Gateway
spec:
  parentRefs:
    - name: my-gateway
  hostnames:
    - api.example.com
  rules:
    - backendRefs:
        - name: api-service
          port: 8080
```

**Result:**

- `api.internal.example.com` → A record `172.16.6.6`, TTL 300, Cloudflare proxied
- `api.example.com` → A record `1.2.3.4`, TTL 60, Cloudflare proxied

#### Example: Different Record Types (A vs CNAME)

Routes can override Gateway's IP target with a hostname to create CNAME records:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
  annotations:
    # Default: A record pointing to load balancer IP
    external-dns.alpha.kubernetes.io/target: "10.0.0.1"
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: cdn-route
  annotations:
    # Override: CNAME record pointing to CDN
    external-dns.alpha.kubernetes.io/target: "cdn.cloudprovider.com"
spec:
  parentRefs:
    - name: my-gateway
  hostnames:
    - static.example.com
```

**Result:** `static.example.com` → CNAME record `cdn.cloudprovider.com`

#### Example: AWS Route53 with Routing Policies

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: aws-gateway
  annotations:
    external-dns.alpha.kubernetes.io/target: "alb-123.us-east-1.elb.amazonaws.com"
    # Default set-identifier for all Routes
    external-dns.alpha.kubernetes.io/set-identifier: "default"
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: weighted-route
  annotations:
    # Override set-identifier for this specific Route
    external-dns.alpha.kubernetes.io/set-identifier: "backend-v1"
    external-dns.alpha.kubernetes.io/aws-weight: "100"
spec:
  parentRefs:
    - name: aws-gateway
  hostnames:
    - app.example.com
```

For a complete list of supported annotations, see the
[annotations documentation](../annotations/annotations.md).

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
  resources: ["gateways","httproutes","grpcroutes","tlsroutes","tcproutes","udproutes"]
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
        image: registry.k8s.io/external-dns/external-dns:v0.20.0
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
        # Add provider-specific flags...
        - --domain-filter=external-dns-test.my-org.com
        - --provider=google
        - --registry=txt
        - --txt-owner-id=my-identifier
```
