# Configuring ExternalDNS to use the Istio Gateway Source
This tutorial describes how to configure ExternalDNS to use the Istio Gateway source.
It is meant to supplement the other provider-specific setup tutorials.

**Note:** Using the Istio Gateway source requires Istio >=1.0.0.

### Manifest (for clusters without RBAC enabled)
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
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
        - --source=istio-gateway # Create records for hosts specified in a networking.istio.io.Gateway
        - --source=istio-virtual-service # Create records for hosts specified in a networking.istio.io.VirtualService
        - --istio-ingress-gateway=custom-istio-namespace/custom-istio-ingressgateway # load balancer service to be used; can be specified multiple times. Omit to use the default (istio-system/istio-ingressgateway)
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
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
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
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
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
        - --source=istio-gateway # Create records for hosts specified in a networking.istio.io.Gateway
        - --source=istio-virtual-service # Create records for hosts specified in a networking.istio.io.VirtualService
        - --istio-ingress-gateway=custom-istio-namespace/custom-istio-ingressgateway # load balancer service to be used; can be specified multiple times. Omit to use the default (istio-system/istio-ingressgateway)
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
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
    - "httpbin.example.com" # Can be set to "*" if the istio-virtual-service source is enabled
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
