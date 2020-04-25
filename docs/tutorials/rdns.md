# Setting up ExternalDNS for RancherDNS(RDNS) with kubernetes
This tutorial describes how to setup ExternalDNS for usage within a kubernetes cluster that makes use of [RDNS](https://github.com/rancher/rdns-server) and [nginx ingress controller](https://github.com/kubernetes/ingress-nginx).  
You need to:
* install RDNS with [etcd](https://github.com/etcd-io/etcd) enabled
* install external-dns with rdns as a provider

## Installing RDNS with etcdv3 backend

### Clone RDNS
```
git clone https://github.com/rancher/rdns-server.git
```

### Installing ETCD
```
cd rdns-server
docker-compose -f deploy/etcdv3/etcd-compose.yaml up -d
```

> ETCD was successfully deployed on `http://172.31.35.77:2379`

### Installing RDNS
```
export ETCD_ENDPOINTS="http://172.31.35.77:2379"
export DOMAIN="lb.rancher.cloud"
./scripts/start etcdv3
```

> RDNS was successfully deployed on `172.31.35.77`

## Installing ExternalDNS
### Install external ExternalDNS
ETCD_URLS is configured to etcd client service address.
RDNS_ROOT_DOMAIN is configured to the same with RDNS DOMAIN environment. e.g. lb.rancher.cloud.

#### Manifest (for clusters without RBAC enabled)
```yaml
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=ingress
        - --provider=rdns
        - --log-level=debug # debug only
        env:
        - name: ETCD_URLS
          value: http://172.31.35.77:2379
        - name: RDNS_ROOT_DOMAIN
          value: lb.rancher.cloud
```

#### Manifest (for clusters with RBAC enabled)
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
- apiGroups: ["extensions"]
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=ingress
        - --provider=rdns
        - --log-level=debug # debug only
        env:
        - name: ETCD_URLS
          value: http://172.31.35.77:2379
        - name: RDNS_ROOT_DOMAIN
          value: lb.rancher.cloud
```

## Testing ingress example
```
$ cat ingress.yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: nginx
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
  - host: nginx.lb.rancher.cloud
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80

$ kubectl apply -f ingress.yaml
ingress.extensions "nginx" created
```

Wait a moment until DNS has the ingress IP. The RDNS IP in this example is "172.31.35.77".
```
$ kubectl get ingress
NAME      HOSTS                    ADDRESS         PORTS     AGE
nginx     nginx.lb.rancher.cloud   172.31.42.211   80        2m

$ kubectl run -it --rm --restart=Never --image=infoblox/dnstools:latest dnstools
If you don't see a command prompt, try pressing enter.
dnstools# dig @172.31.35.77 nginx.lb.rancher.cloud +short
172.31.42.211
dnstools#  
```
