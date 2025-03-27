# Cluster Nodes as Source

This tutorial describes how to configure ExternalDNS to use the cluster nodes as source.
Using nodes (`--source=node`) as source is possible to synchronize a DNS zone with the nodes of a cluster.

The node source adds an `A` record per each node `externalIP` (if not found, any IPv4 `internalIP` is used instead).
It also adds an `AAAA` record per each node IPv6 `internalIP`. Refer to the [IPv6 Behavior](#ipv6-behavior) section for more details.
The TTL of the records can be set with the `external-dns.alpha.kubernetes.io/ttl` node annotation.

Nodes marked as **Unschedulable** as per [core/v1/NodeSpec](https://pkg.go.dev/k8s.io/api@v0.31.1/core/v1#NodeSpec) are excluded.
This avoid exposing Unhealthy, NotReady or SchedulingDisabled (cordon) nodes.

## IPv6 Behavior

By default, ExternalDNS exposes the IPv6 `InternalIP` of the nodes. To prevent this, you can use the `--no-expose-internal-ipv6` flag.
**The default behavior will change in the next minor release.** ExternalDNS will no longer expose the IPv6 `InternalIP` addresses by default.
You can still explicitly expose the internal ipv6 addresses by using the `--expose-internal-ipv6` flag, if needed.

### Example spec (without exposing IPv6 `InternalIP` addresses)

```yaml
spec:
  serviceAccountName: external-dns
  containers:
  - name: external-dns
    image: registry.k8s.io/external-dns/external-dns:v0.16.1 # update this to the desired external-dns version
    args:
    - --source=node # will use nodes as source
    - --provider=aws
    - --zone-name-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
    - --domain-filter=external-dns-test.my-org.com
    - --aws-zone-type=public
    - --registry=txt
    - --fqdn-template={{.Name}}.external-dns-test.my-org.com
    - --txt-owner-id=my-identifier
    - --policy=sync
    - --log-level=debug
    - --no-expose-internal-ipv6
```

## Manifest (for cluster without RBAC enabled)

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
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
        image: registry.k8s.io/external-dns/external-dns:v0.16.1
        args:
        - --source=node # will use nodes as source
        - --provider=aws
        - --zone-name-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --domain-filter=external-dns-test.my-org.com
        - --aws-zone-type=public
        - --registry=txt
        - --fqdn-template={{.Name}}.external-dns-test.my-org.com
        - --txt-owner-id=my-identifier
        - --policy=sync
        - --log-level=debug
```

## Manifest (for cluster with RBAC enabled)

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: ["route.openshift.io"]
  resources: ["routes"]
  verbs: ["get", "watch", "list"]
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns-viewer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
- kind: ServiceAccount
  name: external-dns
  namespace: external-dns
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
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
        image: registry.k8s.io/external-dns/external-dns:v0.16.1
        args:
        - --source=node # will use nodes as source
        - --provider=aws
        - --zone-name-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --domain-filter=external-dns-test.my-org.com
        - --aws-zone-type=public
        - --registry=txt
        - --fqdn-template={{.Name}}.external-dns-test.my-org.com
        - --txt-owner-id=my-identifier
        - --policy=sync
        - --log-level=debug
```
