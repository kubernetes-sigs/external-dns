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
[etcd chart by bitnami](https://artifacthub.io/packages/helm/bitnami/etcd) is currently the most popular chart for etcd available on artifacthub. First you will need to add bitnami chart repository to use it
```
helm repo add bitnami https://charts.bitnami.com/bitnami
```

<!-- TODO: Add disclamer for riscs of this approach -->
To make it easier to connect later on, you need to set root password through values.yaml. First let's download them:
```
wget https://raw.githubusercontent.com/bitnami/charts/master/bitnami/etcd/values.yaml
```

Then you need to patch values with this diff
```diff
--- a/values.yaml
+++ b/values.yaml
@@ -101,7 +101,7 @@
     allowNoneAuthentication: true
     ## @param auth.rbac.rootPassword Root user password. The root user is always `root`
     ##
-    rootPassword: ""
+    rootPassword: "NotSecurePassword"
     ## @param auth.rbac.existingSecret Name of the existing secret containing credentials for the root user
     ##
     existingSecret: ""
```

Finally to install etcd with those values run
```
helm install my-etcd --values values.yaml bitnami/etcd
```

### Installing CoreDNS
First, you need to add CoreDNS helm chart repository:
```
helm repo add coredns https://coredns.github.io/helm
```

In order to make CoreDNS work with etcd backend, values.yaml of the chart should be changed with corresponding configurations.
```
wget https://raw.githubusercontent.com/coredns/helm/master/charts/coredns/values.yaml
```

You need to edit/patch the file with below diff
```diff
--- a/values.yaml
+++ b/values.yaml
@@ -81,7 +81,7 @@
   # name:
 
 # isClusterService specifies whether chart should be deployed as cluster-service or normal k8s app.
-isClusterService: true
+isClusterService: false
 
 # Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
 priorityClassName: ""
@@ -90,7 +90,7 @@
 # https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
 servers:
 - zones:
-  - zone: .
+  - zone: example.org
   port: 53
   plugins:
   - name: errors
@@ -100,18 +100,14 @@
       lameduck 5s
   # Serves a /ready endpoint on :8181, required for readinessProbe
   - name: ready
-  # Required to query kubernetes API for data
-  - name: kubernetes
-    parameters: cluster.local in-addr.arpa ip6.arpa
-    configBlock: |-
-      pods insecure
-      fallthrough in-addr.arpa ip6.arpa
-      ttl 30
   # Serves a /metrics endpoint on :9153, required for serviceMonitor
   - name: prometheus
     parameters: 0.0.0.0:9153
-  - name: forward
-    parameters: . /etc/resolv.conf
+  - name: etcd
+    configBlock: |-
+      path /skydns
+      endpoint http://my-etcd.default.svc.cluster.local:2379
+      credentials root NotSecurePassword
   - name: cache
     parameters: 30
   - name: loop
```
**Note**:
* You could either use parameters or zone to configure domain. "example.org" is used in this example.


After configuration done in values.yaml, you can install coredns chart.
```
helm install my-coredns --values values.yaml coredns/coredns
```

## Installing ExternalDNS
### Install external ExternalDNS
Set this environment variables:
ETCD_URLS='http://my-etcd.default.svc.cluster.local:2379'
ETCD_USER='root' 
ETCD_PASS='NotSecurePassword'

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
          value: http://my-etcd.default.svc.cluster.local:2379
        - name: ETCD_USER
          value: root
        - name: ETCD_PASS
          value: NotSecurePassword
```

#### Manifest (for clusters with RBAC enabled)

```yaml
---
apiVersion: rbac.authorization.k8s.io/v1
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
          value: http://my-etcd.default.svc.cluster.local:2379
        - name: ETCD_USER
          value: root
        - name: ETCD_PASS
          value: NotSecurePassword
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
