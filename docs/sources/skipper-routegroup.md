# Skipper RouteGroup Source

This tutorial describes how to configure ExternalDNS to use the Skipper [RouteGroup](https://opensource.zalando.com/skipper/kubernetes/routegroup-crd) source.
It is meant to supplement the other provider-specific setup tutorials.

## Manifest (for clusters without RBAC enabled)

```yaml
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
      containers:
      - name: external-dns
        # update this to the desired external-dns version
        image: registry.k8s.io/external-dns/external-dns:v0.22.0
        args:
        - --source=skipper-routegroup
        - --policy=upsert-only # prevents ExternalDNS from deleting any records, set --policy=sync to enable full synchronization (including deletions)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```

## Manifest (for clusters with RBAC enabled)

Could be change if you have multiple sources

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
- apiGroups: [""]
  resources: ["services","pods","nodes"]
  verbs: ["get","list","watch"]
- apiGroups: ["discovery.k8s.io"]
  resources: ["endpointslices"]
  verbs: ["get","list","watch"]
- apiGroups: ["zalando.org"]
  resources: ["routegroups"]
  verbs: ["get","list","watch"]
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
  namespace: default
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
        # update this to the desired external-dns version
        image: registry.k8s.io/external-dns/external-dns:v0.22.0
        args:
        - --source=skipper-routegroup
        - --policy=upsert-only # prevents ExternalDNS from deleting any records, set --policy=sync to enable full synchronization (including deletions)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```
