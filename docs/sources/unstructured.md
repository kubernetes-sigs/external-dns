---
tags: ["source", "area/fqdn", "area/source", "templating", "unstructured"]
---

# Unstructured Source

The `unstructured` source creates DNS records from any Kubernetes resource using Go templates.
It works with custom resources (CRDs) without requiring typed Go clients.

## Use Cases

Use this source when:

- Your CRD is not supported by a built-in external-dns source
- The resource exposes DNS-relevant data (hostnames, IPs, endpoints) in `.spec` or `.status`
- A built-in source exists but only supports an older API version than you're using
- You want to experiment with custom controllers or meshes but keep external-dns
- Create DNS entries based on any attribute of any Kubernetes object
  - When controller users Labels or Annotations with Json or flat key-value pairs
  - Crossplane managed resources (RDS, ElastiCache, S3, etc.)
  - Support Endpoints or any other native resources
  - Use ConfigMaps as a lightweight DNS registry without needing custom CRDs
- Allows the community to support new CRDs via configuration rather than code changes

> **Note**: Prefer built-in sources when available (e.g., `istio-virtualservice`, `gateway-httproute`) as they provide optimized handling for those resource types.

### Advanced Use Cases

The unstructured source can also be used with:

**Knative Service** - Serverless workloads expose auto-generated URLs in `.status.url`

```yaml
status:
  url: https://hello.default.example.com
```

**Argo Rollouts** - Canary/blue-green deployments with preview services in `.status.canary.stableRS`

```yaml
status:
  canary:
    stableRS: my-app-stable-abc123
```

**Linkerd ServiceProfile** - Service mesh with destination overrides in `.spec.dstOverrides`

```yaml
spec:
  dstOverrides:
  - authority: webapp.default.svc.cluster.local
```

**Crossplane Composition outputs** - Any Crossplane-managed cloud resource (ElastiCache, S3 websites, CloudFront, etc.)

```yaml
status:
  atProvider:
    configurationEndpoint:
      address: my-cache.abc123.cache.amazonaws.com
```

**Cilium BGP PeeringPolicy** - BGP-advertised IPs for LoadBalancer services

```yaml
status:
  conditions:
  - type: Established
    status: "True"
```

**ACK FieldExport** - AWS Controllers for Kubernetes can export resource status (RDS endpoints, S3 bucket URLs) to ConfigMaps via FieldExport, enabling dynamic DNS records

```yaml
# FieldExport copies S3 bucket URL to ConfigMap
apiVersion: services.k8s.aws/v1alpha1
kind: FieldExport
spec:
  from:
    path: ".status.location"
    resource:
      group: s3.services.k8s.aws
      kind: Bucket
      name: my-bucket
  to:
    kind: configmap
    name: bucket-dns
```

## Configuration

| Flag                          | Description                                                                   |
|-------------------------------|-------------------------------------------------------------------------------|
| `--unstructured-resource`     | Resources to watch in `resource.version.group` format (repeatable)            |
| `--fqdn-template`             | Go template for DNS names                                                     |
| `--target-template`      | Go template for DNS targets                                                   |
| `--fqdn-target-template` | Go template returning `host:target` pairs (mutually exclusive with above two) |
| `--label-filter`              | Filter resources by labels                                                    |
| `--annotation-filter`         | Filter resources by annotations                                               |

## Template Syntax

Templates have access to typed-style fields and raw object data:

| Field          | Description          |
|----------------|----------------------|
| `.Name`        | Object name          |
| `.Namespace`   | Object namespace     |
| `.Kind`        | Object kind          |
| `.APIVersion`  | API version          |
| `.Labels`      | Object labels        |
| `.Annotations` | Object annotations   |
| `.Metadata`    | Raw metadata section |
| `.Spec`        | Raw spec section     |
| `.Status`      | Raw status section   |
| `.Object`      | Raw full object      |

## Examples

### ConfigMap DNS Registry

Use ConfigMaps as a lightweight DNS registry without needing custom CRDs. Useful for GitOps workflows where teams manage DNS entries via ConfigMaps in their namespaces.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: api-dns
  namespace: production
  labels:
    external-dns.alpha.kubernetes.io/dns-controller: "dns-controller"
data:
  hostname: api.example.com
  target: 10.0.0.100
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=configmaps.v1 \
  --fqdn-template='{{index .Object.data "hostname"}}' \
  --target-template='{{index .Object.data "target"}}' \
  --label-filter='external-dns.alpha.kubernetes.io/controller=dns-controller'

# Result:
# api.example.com -> 10.0.0.100 (A)
```

### Crossplane RDS Instance

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=rdsinstances.v1alpha1.rds.aws.crossplane.io \
  --fqdn-template='{{.Name}}.db.example.com' \
  --target-template='{{.Status.atProvider.endpoint.address}}'
```

### Multiple Resources

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=virtualmachineinstances.v1.kubevirt.io \
  --unstructured-resource=rdsinstances.v1alpha1.rds.aws.crossplane.io \
  --fqdn-template='{{.Name}}.{{.Kind}}.example.com' \
  --target-template='{{.Status.endpoint}}'
```

### MetalLB IPAddressPool

```yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: production-pool
  namespace: metallb-system
  annotations:
    external-dns.alpha.kubernetes.io/hostname: "lb.example.com"
spec:
  addresses:
  - 192.168.10.11/32
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=ipaddresspools.v1beta1.metallb.io \
  --fqdn-template='{{index .Annotations "external-dns.alpha.kubernetes.io/hostname"}}' \
  --target-template='{{$addr := index .Spec.addresses 0}}{{if contains $addr "/32"}}{{trimSuffix $addr "/32"}}{{else}}{{$addr}}{{end}}'

# Result:
# lb.example.com -> 192.168.10.11 (A)
```

> **Tip**: Use `contains` with `trimSuffix` to extract the IP from `/32` CIDR notation.

### Apache APISIX Route

```yaml
apiVersion: apisix.apache.org/v2
kind: ApisixRoute
metadata:
  name: httpbin
  namespace: ingress-apisix
spec:
  http:
  - name: httpbin
    match:
      hosts:
      - httpbin.example.com
      paths:
      - /ip
    backends:
    - serviceName: httpbin
      servicePort: 80
status:
  apisix:
    gateway: apisix-gateway.ingress-apisix.svc.cluster.local
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=apisixroutes.v2.apisix.apache.org \
  --fqdn-template='{{.Name}}.route.example.com' \
  --target-template='{{.Status.apisix.gateway}}'

# Result:
# httpbin.route.example.com -> apisix-gateway.ingress-apisix.svc.cluster.local (CNAME)
```

### cert-manager Certificate

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: my-app-tls
  namespace: production
  annotations:
    external-dns.alpha.kubernetes.io/target: "10.0.0.50"
spec:
  secretName: my-app-tls-secret
  dnsNames:
  - my-app.example.com
  - www.my-app.example.com
  issuerRef:
    name: letsencrypt-prod
    kind: ClusterIssuer
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=certificates.v1.cert-manager.io \
  --fqdn-template='{{index .Spec.dnsNames 0}}' \
  --target-template='{{index .Annotations "external-dns.alpha.kubernetes.io/target"}}'

# Result:
# my-app.example.com -> 10.0.0.50 (A)
```

### Rancher Node

```yaml
apiVersion: management.cattle.io/v3
kind: Node
metadata:
  name: my-node-1
  namespace: cattle-system
  labels:
    cattle.io/creator: norman
    node-role.kubernetes.io/controlplane: "true"
spec:
  clusterName: c-abcde
  hostname: my-node-1
status:
  nodeName: worker-01
  internalNodeStatus:
    addresses:
    - type: ExternalIP
      address: 203.0.113.10
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=nodes.v3.management.cattle.io \
  --fqdn-template='{{.Spec.hostname}}.nodes.example.com' \
  --target-template='{{(index .Status.internalNodeStatus.addresses 0).address}}' \
  --label-filter='node-role.kubernetes.io/controlplane=true'

# Result:
# my-node-1.nodes.example.com -> 203.0.113.10 (A)
```

### ACK FieldExport with ConfigMap

Use AWS Controllers for Kubernetes (ACK) to dynamically populate ConfigMaps with resource endpoints. FieldExport copies values from ACK-managed resources (RDS, S3, ElastiCache) to ConfigMaps, which external-dns can then use for DNS records.

```yaml
# 1. ACK creates an S3 bucket
apiVersion: s3.services.k8s.aws/v1alpha1
kind: Bucket
metadata:
  name: app-assets
  namespace: default
spec:
  name: my-app-assets-bucket
---
# 2. FieldExport copies the bucket URL to a ConfigMap
apiVersion: services.k8s.aws/v1alpha1
kind: FieldExport
metadata:
  name: export-bucket-url
  namespace: default
spec:
  from:
    path: ".status.location"
    resource:
      group: s3.services.k8s.aws
      kind: Bucket
      name: app-assets
  to:
    kind: configmap
    name: app-assets-dns
    namespace: default
---
# 3. ConfigMap is populated by FieldExport
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-assets-dns
  namespace: default
  labels:
    app.kubernetes.io/managed-by: ack-fieldexport
data:
  default.export-bucket-url: "https://my-app-assets-bucket.s3.amazonaws.com/"
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=configmaps.v1 \
  --fqdn-template='{{if eq .Kind "ConfigMap"}}{{.Name}}.cdn.example.com{{end}}' \
  --target-template='{{if eq .Kind "ConfigMap"}}{{$url := index .Object.data "default.export-bucket-url"}}{{trimSuffix (trimPrefix $url "https://") "/"}}{{end}}' \
  --label-filter='app.kubernetes.io/managed-by=ack-fieldexport'

# Result:
# app-assets-dns.cdn.example.com -> my-app-assets-bucket.s3.amazonaws.com (CNAME)
```

### EndpointSlice for Headless Services

Create per-pod DNS records from EndpointSlice resources for headless services. Each pod gets its own DNS entry pointing to its IP address.

```yaml
apiVersion: discovery.k8s.io/v1
kind: EndpointSlice
metadata:
  name: test-headless-abc12
  namespace: default
  labels:
    endpointslice.kubernetes.io/managed-by: endpointslice-controller.k8s.io
    kubernetes.io/service-name: test-headless
    service.kubernetes.io/headless: ""
addressType: IPv4
endpoints:
- addresses:
  - 10.244.1.2
  conditions:
    ready: true
  nodeName: worker1
  targetRef:
    kind: Pod
    name: app-abc12
    namespace: default
- addresses:
  - 10.244.2.3
  - 10.244.2.4
  conditions:
    ready: true
  nodeName: worker2
  targetRef:
    kind: Pod
    name: app-def34
    namespace: default
ports:
- name: http
  port: 80
  protocol: TCP
```

```bash
external-dns \
  --source=unstructured \
  --unstructured-resource=endpointslices.v1.discovery.k8s.io \
  --fqdn-target-template='{{if and (eq .Kind "EndpointSlice") (hasKey .Labels "service.kubernetes.io/headless")}}{{range $ep := .Object.endpoints}}{{if $ep.conditions.ready}}{{range $ep.addresses}}{{$ep.targetRef.name}}.pod.com:{{.}},{{end}}{{end}}{{end}}{{end}}'

# Result:
# app-abc12.pod.com -> 10.244.1.2 (A)
# app-def34.pod.com -> 10.244.2.3, 10.244.2.4 (A)
```

The `--fqdn-target-template` flag returns `host:target` pairs, enabling 1:1 mapping between hostnames and targets. Useful when a Kubernetes resource contains arrays where each element should produce its own DNS record (e.g., EndpointSlice endpoints, multi-host configurations).

## RBAC

Grant external-dns access to your custom resources:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
  # Add for each resource type
  - apiGroups: ["rds.aws.crossplane.io"]
    resources: ["rdsinstances"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["<your-api-group>"]
    resources: ["<your-resources>"]
    verbs: ["get", "watch", "list"]
```
