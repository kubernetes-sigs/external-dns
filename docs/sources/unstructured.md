# Unstructured FQDN Source

The `unstructured` source creates DNS records from any Kubernetes resource using Go templates.
It works with custom resources (CRDs) without requiring typed Go clients.

## Use Cases

Use this source when:

- Your CRD is not supported by a built-in external-dns source
- The resource exposes DNS-relevant data (hostnames, IPs, endpoints) in `.spec` or `.status`
- A built-in source exists but only supports an older API version than you're using
- You want to experiment with custom controllers or meshes but keep external-dns

Example CRDs:

- KubeVirt VirtualMachineInstances
- Crossplane managed resources (RDS, ElastiCache, S3, etc.)
- ArgoCD Applications

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

## Configuration

| Flag                           | Description                                                        |
|--------------------------------|--------------------------------------------------------------------|
| `--unstructured-fqdn-resource` | Resources to watch in `resource.version.group` format (repeatable) |
| `--fqdn-template`              | Go template for DNS names                                          |
| `--fqdn-target-template`       | Go template for DNS targets                                        |
| `--label-filter`               | Filter resources by labels                                         |
| `--annotation-filter`          | Filter resources by annotations                                    |

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

## Examples

### Crossplane RDS Instance

```bash
external-dns \
  --source=unstructured \
  --unstructured-fqdn-resource=rdsinstances.v1alpha1.rds.aws.crossplane.io \
  --fqdn-template='{{.Name}}.db.example.com' \
  --fqdn-target-template='{{.Status.atProvider.endpoint.address}}'
```

### Multiple Resources

```bash
external-dns \
  --source=unstructured \
  --unstructured-fqdn-resource=virtualmachineinstances.v1.kubevirt.io \
  --unstructured-fqdn-resource=rdsinstances.v1alpha1.rds.aws.crossplane.io \
  --fqdn-template='{{.Name}}.{{.Kind}}.example.com' \
  --fqdn-target-template='{{.Status.endpoint}}'
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
  --unstructured-fqdn-resource=ipaddresspools.v1beta1.metallb.io \
  --fqdn-template='{{index .Annotations "external-dns.alpha.kubernetes.io/hostname"}}' \
  --fqdn-target-template='{{index .Spec.addresses 0}}'

# Result:
# lb.example.com -> 192.168.10.11/32 (CNAME)
```

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
  --unstructured-fqdn-resource=apisixroutes.v2.apisix.apache.org \
  --fqdn-template='{{.Name}}.route.example.com' \
  --fqdn-target-template='{{.Status.apisix.gateway}}'

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
  --unstructured-fqdn-resource=certificates.v1.cert-manager.io \
  --fqdn-template='{{index .Spec.dnsNames 0}}' \
  --fqdn-target-template='{{index .Annotations "external-dns.alpha.kubernetes.io/target"}}'

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
  --unstructured-fqdn-resource=nodes.v3.management.cattle.io \
  --fqdn-template='{{.Spec.hostname}}.nodes.example.com' \
  --fqdn-target-template='{{(index .Status.internalNodeStatus.addresses 0).address}}' \
  --label-filter='node-role.kubernetes.io/controlplane=true'

# Result:
# my-node-1.nodes.example.com -> 203.0.113.10 (A)
```

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
