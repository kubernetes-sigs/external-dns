# Setting up ExternalDNS for Infomaniak

Here we show how to use Infomaniak as a DNS provider for ExternalDNS.

https://www.infomaniak.com/en/domains

Provide a kubernetes cluster and deploy these manifests:

## RBAC

```yaml
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
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
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: kube-system
```

## Secret

Generate a token at https://manager.infomaniak.com/v3/<id>/api/dashboard and create a secret

```shell
kubectl -n kube-system create secret generic external-dns-token \
--from-literal=INFOMANIAK_API_TOKEN=xxx
```

Deployment (replacing XXX by the latest external-dns tag version)

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: kube-system
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
        image: k8s.gcr.io/external-dns/external-dns:XXX
        args:
        - --source=ingress
        - --source=service
        - --provider=infomaniak
        - --log-level=debug # debug only
        - --domain-filter=example.dev
        env:
          - name: INFOMANIAK_API_TOKEN
            valueFrom:
              secretKeyRef:
                name: external-dns-token
                key: INFOMANIAK_API_TOKEN
```
