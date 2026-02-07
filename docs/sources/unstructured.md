# Unstructured FQDN Source

The unstructured-fqdn source creates DNS records from any Kubernetes resource using Go templates.
It works with custom resources (CRDs) without requiring typed Go clients.

## Use Cases

Use this source when:

- Your CRD is not supported by a built-in external-dns source
- The resource exposes DNS-relevant data (hostnames, IPs, endpoints) in `.spec` or `.status`
- A built-in source exists but only supports an older API version than you're using

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

| Field          | Type                     | Description          |
|----------------|--------------------------|----------------------|
| `.Name`        | `string`                 | Object name          |
| `.Namespace`   | `string`                 | Object namespace     |
| `.Kind`        | `string`                 | Object kind          |
| `.APIVersion`  | `string`                 | API version          |
| `.Labels`      | `map[string]string`      | Object labels        |
| `.Annotations` | `map[string]string`      | Object annotations   |
| `.Metadata`    | `map[string]interface{}` | Raw metadata section |
| `.Spec`        | `map[string]interface{}` | Raw spec section     |
| `.Status`      | `map[string]interface{}` | Raw status section   |

### Template Functions

- `toLower`, `contains`, `trimPrefix`, `trimSuffix`, `trim`
- `replace <old> <new> <string>` - Replace substrings
- `index <map/slice> <key/index>...` - Access nested fields/arrays

## Examples

### KubeVirt VirtualMachineInstance

```bash
external-dns \
  --source=unstructured \
  --unstructured-fqdn-resource=virtualmachineinstances.v1.kubevirt.io \
  --fqdn-template='{{.Name}}.{{.Namespace}}.vmi.example.com' \
  --fqdn-target-template='{{index .Status.interfaces 0 "ipAddress"}}'
```

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

### Using Labels

```bash
external-dns \
  --source=unstructured \
  --unstructured-fqdn-resource=applications.v1alpha1.argoproj.io \
  --fqdn-template='{{index .Labels "app.kubernetes.io/instance"}}.apps.example.com' \
  --fqdn-target-template='{{.Status.loadBalancer}}'
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

## RBAC

Grant external-dns access to your custom resources:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
  # Add for each resource type
  - apiGroups: ["kubevirt.io"]
    resources: ["virtualmachineinstances"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["rds.aws.crossplane.io"]
    resources: ["rdsinstances"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["metallb.io"]
    resources: ["ipaddresspools"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["apisix.apache.org"]
    resources: ["apisixroutes"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["cert-manager.io"]
    resources: ["certificates"]
    verbs: ["get", "watch", "list"]
```
