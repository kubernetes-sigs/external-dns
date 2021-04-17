# Setting up ExternalDNS for CoreDNS with minikube
This tutorial describes how to setup ExternalDNS for usage within a [minikube](https://github.com/kubernetes/minikube) cluster that makes use of [CoreDNS](https://github.com/coredns/coredns) and [nginx ingress controller](https://github.com/kubernetes/ingress-nginx).
You need to:
* install CoreDNS with [etcd](https://github.com/etcd-io/etcd) enabled
* install external-dns with coredns as a provider
* enable ingress controller for the minikube cluster


## Creating a cluster
```
minikube start
```

## Installing CoreDNS with etcd enabled
Helm chart is used to install etcd and CoreDNS.
### Initializing helm chart
```
helm init
```
### Installing etcd
[etcd operator](https://github.com/coreos/etcd-operator) is used to manage etcd clusters.
```
helm install stable/etcd-operator --name my-etcd-op
```
etcd cluster is installed with example yaml from etcd operator website.
```
kubectl apply -f https://raw.githubusercontent.com/coreos/etcd-operator/HEAD/example/example-etcd-cluster.yaml
```

### Installing CoreDNS
In order to make CoreDNS work with etcd backend, values.yaml of the chart should be changed with corresponding configurations.
```
wget https://raw.githubusercontent.com/helm/charts/HEAD/stable/coredns/values.yaml
```

You need to edit/patch the file with below diff
```diff
diff --git a/values.yaml b/values.yaml
index 964e72b..e2fa934 100644
--- a/values.yaml
+++ b/values.yaml
@@ -27,12 +27,12 @@ service:

 rbac:
   # If true, create & use RBAC resources
-  create: false
+  create: true
   # Ignored if rbac.create is true
   serviceAccountName: default

 # isClusterService specifies whether chart should be deployed as cluster-service or normal k8s app.
-isClusterService: true
+isClusterService: false

 servers:
 - zones:
@@ -51,6 +51,12 @@ servers:
     parameters: 0.0.0.0:9153
   - name: proxy
     parameters: . /etc/resolv.conf
+  - name: etcd
+    parameters: example.org
+    configBlock: |-
+      stubzones
+      path /skydns
+      endpoint http://10.105.68.165:2379

 # Complete example with all the options:
 # - zones:                 # the `zones` block can be left out entirely, defaults to "."
```
**Note**:
* IP address of etcd's endpoint should be get from etcd client service. It should be "example-etcd-cluster-client" in this example. This IP address is used through this document for etcd endpoint configuration.
```
$ kubectl get svc example-etcd-cluster-client
NAME                          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
example-etcd-cluster-client   ClusterIP   10.105.68.165   <none>        2379/TCP   16m
```
* Parameters should configure your own domain. "example.org" is used in this example.


After configuration done in values.yaml, you can install coredns chart.
```
helm install --name my-coredns --values values.yaml stable/coredns
```

## Installing ExternalDNS
### Install external ExternalDNS
ETCD_URLS is configured to etcd client service address.

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
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=ingress
        - --provider=coredns
        - --log-level=debug # debug only
        env:
        - name: ETCD_URLS
          value: http://10.105.68.165:2379
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=ingress
        - --provider=coredns
        - --log-level=debug # debug only
        env:
        - name: ETCD_URLS
          value: http://10.105.68.165:2379
```

## Enable the ingress controller
You can use the ingress controller in minikube cluster. It needs to enable ingress addon in the cluster.
```
minikube addons enable ingress
```

## Testing ingress example
```
$ cat ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
  - host: nginx.example.org
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80

$ kubectl apply -f ingress.yaml
ingress.extensions "nginx" created
```


Wait a moment until DNS has the ingress IP. The DNS service IP is from CoreDNS service. It is "my-coredns-coredns" in this example.
```
$ kubectl get svc my-coredns-coredns
NAME                 TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
my-coredns-coredns   ClusterIP   10.100.4.143   <none>        53/UDP    12m

$ kubectl get ingress
NAME      HOSTS               ADDRESS     PORTS     AGE
nginx     nginx.example.org   10.0.2.15   80        2m

$ kubectl run -it --rm --restart=Never --image=infoblox/dnstools:latest dnstools
If you don't see a command prompt, try pressing enter.
dnstools# dig @10.100.4.143 nginx.example.org +short
10.0.2.15
dnstools#
```
