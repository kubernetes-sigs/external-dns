# Setting up External DNS with Contour

This tutorial describes how to configure External DNS to use either the Contour `IngressRoute` or `HTTPProxy` source.
The `IngressRoute` CRD is deprecated but still in-use in many clusters however it's recommended that you migrate to the `HTTPProxy` resource.
Using the `HTTPProxy` resource with External DNS requires Contour version 1.5 or greater.

### Example manifests for External DNS
#### Without RBAC
Note that you don't need to enable both of the sources and if you don't enable `contour-ingressroute` you also don't need to configure the `contour-load-balancer` setting.

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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --source=ingress
        - --source=contour-ingressroute # To enable IngressRoute support
        - --source=contour-httpproxy # To enable HTTPProxy support
        - --contour-load-balancer=custom-contour-namespace/custom-contour-lb # For IngressRoute ONLY: load balancer service to be used. Omit to use the default (heptio-contour/contour) 
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

#### With RBAC
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
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
# This section is only for IngressRoute
- apiGroups: ["contour.heptio.com"]
  resources: ["ingressroutes"]
  verbs: ["get","watch","list"]
# This section is only for HTTPProxy
- apiGroups: ["projectcontour.io"]
  resources: ["httpproxies"]
  verbs: ["get","watch","list"]
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --source=ingress
        - --source=contour-ingressroute # To enable IngressRoute support
        - --source=contour-httpproxy # To enable HTTPProxy support
        - --contour-load-balancer=custom-contour-namespace/custom-contour-lb # For IngressRoute ONLY: load balancer service to be used. Omit to use the default (heptio-contour/contour) 
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

### Verify External DNS works
The following instructions are based on the 
[Contour example workload](https://github.com/projectcontour/contour/tree/master/examples/example-workload/httpproxy).

#### Install a sample service
```bash
$ kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: kuard
  name: kuard
spec:
  replicas: 3
  selector:
    matchLabels:
      app: kuard
  template:
    metadata:
      labels:
        app: kuard
    spec:
      containers:
      - image: gcr.io/kuar-demo/kuard-amd64:1
        name: kuard
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: kuard
  name: kuard
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    app: kuard
  sessionAffinity: None
  type: ClusterIP
EOF
```

Then create either a `HTTPProxy` or an `IngressRoute`

#### HTTPProxy
```
$ kubectl apply -f - <<EOF
apiVersion: projectcontour.io/v1
kind: HTTPProxy
metadata:
  labels:
    app: kuard
  name: kuard
  namespace: default
spec:
  virtualhost:
    fqdn: kuard.example.com
  routes:
    - conditions:
      - prefix: /
      services:
        - name: kuard
          port: 80
EOF
```

#### IngressRoute
```
$ kubectl apply -f - <<EOF
apiVersion: contour.heptio.com/v1beta1
kind: IngressRoute
metadata: 
  labels:
    app: kuard
  name: kuard
  namespace: default
spec: 
  virtualhost:
    fqdn: kuard.example.com
  routes: 
  - match: /
    services: 
    - name: kuard
      port: 80
EOF
```

#### Access the sample service using `curl`
```bash
$ curl -i http://kuard.example.com/healthy
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Thu, 27 Jun 2019 19:42:26 GMT
Content-Length: 2

ok
```
