# Multi-Namespace Deployments

This guide explains how to configure external-dns to watch multiple namespaces in a multi-tenant cluster.

## Overview

By default, external-dns operates in one of two modes:
- **Single namespace**: `--namespace=production` (watches only one namespace)
- **Cluster-wide**: `--namespace=""` (watches all namespaces)

For multi-tenant clusters, you may want to watch a **selected set of namespaces** without granting cluster-wide access. External-dns now supports **two approaches** for multi-namespace selection.

## Use Cases

### Multi-Tenant Clusters
Each tenant has a set of namespaces identified by labels. You want:
- One external-dns instance per tenant
- Each instance only managing DNS for that tenant's namespaces
- Isolation between tenants

### Gateway + Application Separation
Your Gateway (Istio, Kong, etc.) runs in one namespace, while application services run in separate tenant namespaces. You need external-dns to watch both.

### Environment-Based Deployment
You want to manage DNS for specific environments (prod, staging) without processing development or test namespaces.

---

## Two Approaches for Multi-Namespace Selection

### Approach 1: Explicit Namespace List (Recommended for Small Fixed Sets)

Use `--namespaces` to specify an explicit comma-separated list of namespaces.

**Syntax:**
```bash
--namespaces=namespace1,namespace2,namespace3
```

**Example:**
```bash
external-dns \
  --source=service \
  --namespaces=prod,staging \
  --provider=aws
```

**How it works:**
- Creates per-namespace informer factories (one per namespace)
- Only caches resources from specified namespaces
- Lower memory footprint
- No namespace read permissions needed (unless using label selector)

**When to use:**
- You know exactly which namespaces to watch (fixed list)
- Small to medium number of namespaces (< 10-20)
- You want minimal RBAC permissions
- You want lowest memory usage

---

### Approach 2: Namespace Label Selector (Recommended for Dynamic Discovery)

Use `--namespace-label-selector` to select namespaces by label using standard Kubernetes label selector syntax.

**Syntax:**
```bash
--namespace-label-selector=<label-selector-expression>
```

**Example:**
```bash
external-dns \
  --source=service \
  --namespace-label-selector=tenant=acme \
  --provider=aws
```

**How it works:**
- Creates cluster-wide informers
- Reads namespace labels at runtime
- Filters services based on namespace label matching
- Dynamic discovery (automatically picks up new matching namespaces)

**When to use:**
- Namespaces are identified by custom labels (multi-tenant)
- Dynamic namespace creation/deletion (no restart needed)
- You want label-based tenant isolation
- You prefer standard Kubernetes label selector syntax

---

## Label Selector Syntax

external-dns uses standard Kubernetes label selector syntax:

### Equality-based
```bash
# Single label
--namespace-label-selector=environment=production

# Multiple labels (AND logic)
--namespace-label-selector=tenant=acme,environment=production
```

### Set-based
```bash
# Match multiple values (OR logic)
--namespace-label-selector=kubernetes.io/metadata.name in (prod,staging)

# Exclude namespaces
--namespace-label-selector=tenant=acme,environment notin (test)

# Check label existence
--namespace-label-selector=managed-by-external-dns
```

### Common Patterns

| Scenario | Label Selector |
|----------|----------------|
| Select specific namespaces by name | `kubernetes.io/metadata.name in (ns-a,ns-b)` |
| Select by tenant label | `tenant=acme` |
| Select by environment | `environment in (prod,staging)` |
| Multi-condition | `tenant=acme,environment=production` |

---

## Deployment Examples

### Example 1: Kubernetes Manifest (Explicit Namespaces)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: external-dns
spec:
  replicas: 1
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
        image: registry.k8s.io/external-dns/external-dns:v0.15.0
        args:
        - --source=service
        - --source=ingress
        - --namespaces=prod,staging,istio-gateway  # Explicit list
        - --provider=aws
        - --domain-filter=example.com
        - --log-level=info
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services", "endpoints", "pods"]
  verbs: ["get", "watch", "list"]
- apiGroups: ["extensions", "networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get", "watch", "list"]
# No namespace permissions needed for explicit list
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
  namespace: external-dns
```

---

### Example 2: Kubernetes Manifest (Label Selector)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns-tenant-acme
  namespace: external-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns-tenant-acme
  template:
    metadata:
      labels:
        app: external-dns-tenant-acme
    spec:
      serviceAccountName: external-dns-tenant-acme
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.15.0
        args:
        - --source=service
        - --source=ingress
        - --namespace-label-selector=tenant=acme  # Label-based
        - --provider=aws
        - --domain-filter=acme.example.com
        - --log-level=info
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns-tenant-acme
  namespace: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns-tenant-acme
rules:
- apiGroups: [""]
  resources: ["services", "endpoints", "pods"]
  verbs: ["get", "watch", "list"]
- apiGroups: [""]
  resources: ["namespaces"]  # Required for label selector
  verbs: ["get", "watch", "list"]
- apiGroups: ["extensions", "networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns-tenant-acme
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns-tenant-acme
subjects:
- kind: ServiceAccount
  name: external-dns-tenant-acme
  namespace: external-dns
```

---

## RBAC Requirements

### Explicit Namespace List (`--namespaces`)

**Minimal RBAC** (namespace-specific permissions sufficient):

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns-minimal
rules:
- apiGroups: [""]
  resources: ["services", "endpoints", "pods"]
  verbs: ["get", "watch", "list"]
- apiGroups: ["extensions", "networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get", "watch", "list"]
# No namespace permissions needed
```

### Label Selector (`--namespace-label-selector`)

**⚠️ Important: Namespace Permissions Required**

When using `--namespace-label-selector`, external-dns requires permissions to read namespace objects:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns-label-selector
rules:
- apiGroups: [""]
  resources: ["services", "endpoints", "pods"]
  verbs: ["get", "watch", "list"]
- apiGroups: [""]
  resources: ["namespaces"]  # Required for label-based filtering
  verbs: ["get", "watch", "list"]
- apiGroups: ["extensions", "networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get", "watch", "list"]
```

**Why?** external-dns reads namespace labels to determine which namespaces match the selector.

---

## Multi-Tenant Setup

### Scenario: 3 Tenants, Separate DNS Zones

#### Tenant A (Explicit Namespaces)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns-tenant-a
  namespace: external-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns-tenant-a
  template:
    metadata:
      labels:
        app: external-dns-tenant-a
    spec:
      serviceAccountName: external-dns-tenant-a
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.15.0
        args:
        - --source=service
        - --namespaces=tenant-a-prod,tenant-a-staging
        - --provider=aws
        - --txt-owner-id=tenant-a
        - --domain-filter=tenant-a.example.com
```

#### Tenant B (Label Selector)

**First, label the namespaces:**
```bash
kubectl label namespace tenant-b-prod tenant=tenant-b
kubectl label namespace tenant-b-staging tenant=tenant-b
```

**Then deploy external-dns:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns-tenant-b
  namespace: external-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns-tenant-b
  template:
    metadata:
      labels:
        app: external-dns-tenant-b
    spec:
      serviceAccountName: external-dns-tenant-b
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.15.0
        args:
        - --source=service
        - --namespace-label-selector=tenant=tenant-b
        - --provider=aws
        - --txt-owner-id=tenant-b
        - --domain-filter=tenant-b.example.com
```

**Result**: Each tenant's external-dns instance only manages DNS for its own namespaces.

---

## Choosing the Right Approach

| Factor | Explicit Namespaces (`--namespaces`) | Label Selector (`--namespace-label-selector`) |
|--------|--------------------------------------|----------------------------------------------|
| **Use case** | Fixed namespace list | Dynamic tenant-by-label |
| **Memory usage** | ✅ Lowest (per-namespace caches) | ⚠️ Higher (cluster-wide cache) |
| **RBAC** | ✅ Minimal (no namespace read) | ⚠️ Requires namespace read permissions |
| **Dynamic discovery** | ❌ Restart needed for new namespaces | ✅ Automatic (label-based) |
| **Implementation** | Per-namespace informers | Cluster-wide informer + filter |
| **Best for** | 2-10 known namespaces | Multi-tenant with labeled namespaces |

---

## Validation & Troubleshooting

### Verify Namespace Selection

Enable debug logging to see which namespaces are selected:

```bash
--log-level=debug
```

**Explicit namespace logs:**
```
level=info msg="Creating informer factory for namespace: tenant-a"
level=info msg="Creating informer factory for namespace: tenant-b"
level=debug msg="Found 2 services before namespace filtering"
```

**Label selector logs:**
```
level=info msg="Using namespace label selector: tenant=acme"
level=debug msg="Evaluating service tenant-a/my-svc in namespace tenant-a with labels map[tenant:acme]"
level=debug msg="Including service tenant-a/my-svc: namespace matches selector"
level=debug msg="Skipping service other-ns/other-svc: namespace does not match selector"
```

### Common Issues

#### Issue: "No endpoints generated"
**Cause**: Services in selected namespaces lack annotations or LoadBalancer IPs.

**Solution**: Ensure services have:
```yaml
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: myapp.example.com
spec:
  type: LoadBalancer
status:
  loadBalancer:
    ingress:
    - ip: 1.2.3.4
```

#### Issue: "Invalid namespace label selector"
**Cause**: Incorrect label selector syntax.

**Solution**: Use valid Kubernetes label selector syntax:
- ✅ `tenant=acme`
- ✅ `kubernetes.io/metadata.name in (a,b)`
- ❌ `'tenant=acme'` (remove quotes)
- ❌ `tenant:acme` (use `=`, not `:`)

#### Issue: "Unable to get namespace"
**Cause**: Missing namespace RBAC permissions (label-selector mode only).

**Solution**: Ensure ClusterRole includes:
```yaml
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["get", "watch", "list"]
```

---

## Migration from Single Namespace

### Before (single namespace per instance)
```bash
# Instance 1
--namespace=prod

# Instance 2
--namespace=staging
```

### After (one instance, multiple namespaces)

**Option 1: Explicit list**
```bash
--namespaces=prod,staging
```

**Option 2: Label selector**
```bash
# First, label namespaces
kubectl label namespace prod environment=production
kubectl label namespace staging environment=production

# Single instance
--namespace-label-selector=environment=production
```

---

## Flag Compatibility

### Mutually Exclusive Flags

**You cannot use these flags together:**
- `--namespace` (single)
- `--namespaces` (explicit list)
- `--namespace-label-selector` (label-based)

**Error example:**
```bash
external-dns \
  --namespace=prod \
  --namespaces=staging

# ERROR: only one of --namespace, --namespaces, or --namespace-label-selector can be specified
```

### Precedence (if validation is not enforced)

1. `--namespace-label-selector` (highest priority)
2. `--namespaces`
3. `--namespace` (backwards compatibility, lowest priority)

---

## Performance Considerations

### Memory Usage

| Configuration | Memory Impact | When to Use |
|---------------|---------------|-------------|
| `--namespace=single` | Lowest | Single namespace only |
| `--namespaces=a,b,c` | Low-Medium | 2-20 namespaces |
| `--namespace-label-selector` | Medium-High | Dynamic multi-tenant, labeled namespaces |
| Cluster-wide (no flags) | Highest | Need all namespaces |

### CPU/Watch Overhead

- **Explicit namespaces**: N watch connections (one per namespace)
- **Label selector**: 1 watch connection (cluster-wide) + runtime filtering

---

## Best Practices

1. **Use explicit namespaces** for fixed, small namespace sets (≤ 10 namespaces)
2. **Use label selector** for dynamic multi-tenant environments
3. **Always set `--txt-owner-id`** to differentiate tenant instances
4. **Use `--domain-filter`** to scope DNS zones per tenant
5. **Document your namespace labeling convention** for label-selector approach
6. **Monitor memory usage** when using label-selector or cluster-wide mode

---

## Next Steps

- **Coming soon**: Additional sources (Ingress, Gateway, etc.) will support multi-namespace selection in follow-up PRs
- **Feedback welcome**: This feature is under active development; please report issues or suggestions

---

## Related Documentation

- [Installation Guide](installation.md)
- [RBAC Configuration](rbac.md)
- [Annotation Reference](annotations.md)
- [Flag Reference](flags.md)

---

## Summary Table

| Approach | Flag | Example | RBAC | Memory | Dynamic |
|----------|------|---------|------|--------|---------|
| **Single** | `--namespace` | `--namespace=prod` | Minimal | Lowest | ❌ |
| **Explicit List** | `--namespaces` | `--namespaces=a,b` | Minimal | Low | ❌ |
| **Label Selector** | `--namespace-label-selector` | `--namespace-label-selector=tenant=acme` | Namespace read required | Medium | ✅ |
| **Cluster-wide** | *(none)* | *(none)* | Cluster-wide | Highest | ✅ |
