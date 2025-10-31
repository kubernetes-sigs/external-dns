# Gloo Proxy Source

This tutorial describes how to configure ExternalDNS to use the Gloo Proxy source.
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
        image: registry.k8s.io/external-dns/external-dns:v0.19.0
        args:
        - --source=gloo-proxy
        - --gloo-namespace=custom-gloo-system # gloo system namespace. Specify multiple times for multiple namespaces. Omit to use the default (gloo-system)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```

## Manifest (for clusters with RBAC enabled)

Could be change if you have mulitple sources

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
  resources: ["services","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["discovery.k8s.io"]
  resources: ["endpointslices"]
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
        image: registry.k8s.io/external-dns/external-dns:v0.19.0
        args:
        - --source=gloo-proxy
        - --gloo-namespace=custom-gloo-system # gloo system namespace. Specify multiple times for multiple namespaces. Omit to use the default (gloo-system)
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
```

## Gateway Annotation

To support setups where an Ingress resource is used provision an external LB you can add the following annotation to your Gateway

**Note:** The Ingress namespace can be omitted if its in the same namespace as the gateway

```bash
$ cat <<EOF | kubectl apply -f -
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
    metadataStatic:
      sources:
      - resourceKind: '*v1.Gateway'
        resourceRef:
          name: gateway-proxy
          namespace: gloo-system
---
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/ingress: "$ingressNamespace/$ingressName"
  labels:
    app: gloo
  name: gateway-proxy
  namespace: gloo-system
spec: {}
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  labels:
    gateway-proxy-id: gateway-proxy
    gloo: gateway-proxy
  name: gateway-proxy
  namespace: gloo-system
spec:
  ingressClassName: alb
EOF
```
