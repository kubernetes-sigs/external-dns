# Configuring ExternalDNS to use the Istio Gateway Source
This tutorial describes how to configure ExternalDNS to use the Istio Gateway source.
It is meant to supplement the other provider-specific setup tutorials.

**Note:** Using the Istio Gateway source requires Istio >=1.0.0.

* Manifest (for clusters without RBAC enabled)
* Manifest (for clusters with RBAC enabled)
* Update existing Existing DNS Deployment

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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --source=istio-gateway
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

### Manifest (for clusters with RBAC enabled)
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
- apiGroups: ["extensions"] 
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
- apiGroups: ["networking.istio.io"]
  resources: ["gateways"]
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --source=istio-gateway
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

### Update existing Existing DNS Deployment

* For clusters with running `external-dns`, you can just update the deployment.
* With access to the `kube-system` namespace, update the existing `external-dns` deployment.
  * Add a parameter to the arguments of the container to create dns entries with `--source=istio-gateway`.

Execute the following command or update the argument.

```console
kubectl patch deployment external-dns --type='json' \
  -p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/2", "value": "--source=istio-gateway" }]'
```

In case the setup uses a `clusterrole`, just append a new value to the enable the istio group.

```console
kubectl patch clusterrole external-dns --type='json' \
  -p='[{"op": "add", "path": "/rules/4", "value": { "apiGroups": [ "networking.istio.io"], "resources": ["gateways"],"verbs": ["get", "watch", "list" ]} }]'
```

### Verify External DNS works (Gateway example)

Follow the [Istio ingress traffic tutorial](https://istio.io/docs/tasks/traffic-management/ingress/) 
to deploy a sample service that will be exposed outside of the service mesh.
The following are relevant snippets from that tutorial.

#### Install a sample service
With automatic sidecar injection:
```bash
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.0/samples/httpbin/httpbin.yaml
```

Otherwise:
```bash
$ kubectl apply -f <(istioctl kube-inject -f https://raw.githubusercontent.com/istio/istio/release-1.0/samples/httpbin/httpbin.yaml)
```

#### Create an Istio Gateway:
```bash
$ cat <<EOF | kubectl apply -f -
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: httpbin-gateway
spec:
  selector:
    istio: ingressgateway # use Istio default gateway implementation
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "httpbin.example.com"
EOF
```

#### Configure routes for traffic entering via the Gateway:
```bash
$ cat <<EOF | kubectl apply -f -
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: httpbin
spec:
  hosts:
  - "httpbin.example.com"
  gateways:
  - httpbin-gateway
  http:
  - match:
    - uri:
        prefix: /status
    - uri:
        prefix: /delay
    route:
    - destination:
        port:
          number: 8000
        host: httpbin
EOF
```

#### Access the sample service using `curl`
```bash
$ curl -I http://httpbin.example.com/status/200
HTTP/1.1 200 OK
server: envoy
date: Tue, 28 Aug 2018 15:26:47 GMT
content-type: text/html; charset=utf-8
access-control-allow-origin: *
access-control-allow-credentials: true
content-length: 0
x-envoy-upstream-service-time: 5
```

Accessing any other URL that has not been explicitly exposed should return an HTTP 404 error:
```bash
$ curl -I http://httpbin.example.com/headers
HTTP/1.1 404 Not Found
date: Tue, 28 Aug 2018 15:27:48 GMT
server: envoy
transfer-encoding: chunked
```

**Note:** The `-H` flag in the original Istio tutorial is no longer necessary in the `curl` commands.

### Debug External-DNS

* Look for the deployment pod to see the status

```console$ kubectl get pods | grep external-dns
external-dns-6b84999479-4knv9     1/1     Running   0   3h29m
```

* Watch for the logs as follows

```console
$ kubectl logs -f external-dns-6b84999479-4knv9
```

At this point, you can `create` or `update` any `Istio Gateway` object with `hosts` entries array.

> **ATTENTION**: Make sure to specify those whose account is related to the DNS record.

* Successful executions will print the following

```console
time="2020-01-17T06:08:08Z" level=info msg="Desired change: CREATE httpbin.example.com A"
time="2020-01-17T06:08:08Z" level=info msg="Desired change: CREATE httpbin.example.comm TXT"
time="2020-01-17T06:08:08Z" level=info msg="2 record(s) in zone example.comm. were successfully updated"
time="2020-01-17T06:09:08Z" level=info msg="All records are already up to date, there are no changes for the matching hosted zones"
```

* If there's any problem around `clusterrole`, you would see the errors showing wrong permissions:

```console
source \"gateways\" in API group \"networking.istio.io\" at the cluster scope"
time="2020-01-17T06:07:08Z" level=error msg="gateways.networking.istio.io is forbidden: User \"system:serviceaccount:kube-system:external-dns\" cannot list resource \"gateways\" in API group \"networking.istio.io\" at the cluster scope"
```
