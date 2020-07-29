# Configuring ExternalDNS to use the Gloo Proxy Source
This tutorial describes how to configure ExternalDNS to use the Gloo Proxy source.
It is meant to supplement the other provider-specific setup tutorials.

### Manifest (for clusters without RBAC enabled)
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=gloo-proxy
        - --gloo-namespace=custom-gloo-system # gloo system namespace. Omit to use the default (gloo-system)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```

### Manifest (for clusters with RBAC enabled)
Could be change if you have mulitple sources

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list","watch"]
- apiGroups: ["gloo.solo.io"]
  resources: ["proxies"]
  verbs: ["get","watch","list"]
- apiGroups: ["gateway.solo.io"]
  resources: ["virtualservices"]
  verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=gloo-proxy
        - --gloo-namespace=custom-gloo-system # gloo system namespace. Omit to use the default (gloo-system)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```

