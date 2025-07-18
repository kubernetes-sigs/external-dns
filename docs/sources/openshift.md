# OpenShift Route Source

This tutorial describes how to configure ExternalDNS to use the OpenShift Route source.
It is meant to supplement the other provider-specific setup tutorials.

## For OCP 4.x

In OCP 4.x, if you have multiple [OpenShift ingress controllers](https://docs.openshift.com/container-platform/4.9/networking/ingress-operator.html) then you must specify an ingress controller name (also called router name), you can get it from the route's `status.ingress[*].routerName` field.
If you don't specify a router name when you have multiple ingress controllers in your cluster then the first router from the route's `status.ingress` will be used. Note that the router must have admitted the route in order to be selected.
Once the router is known, ExternalDNS will use this router's canonical hostname as the target for the CNAME record.

Starting from OCP 4.10 you can use [ExternalDNS Operator](https://github.com/openshift/external-dns-operator) to manage ExternalDNS instances. Example of its custom resource for AWS provider:

```yaml
  apiVersion: externaldns.olm.openshift.io/v1alpha1
  kind: ExternalDNS
  metadata:
    name: sample
  spec:
    provider:
      type: AWS
    source:
      openshiftRouteOptions:
        routerName: default
      type: OpenShiftRoute
    zones:
      - Z05387772BD5723IZFRX3
```

This will create an ExternalDNS POD with the following container args in `external-dns` namespace:

```yaml
spec:
  containers:
  - args:
    - --metrics-address=127.0.0.1:7979
    - --txt-owner-id=external-dns-sample
    - --provider=aws
    - --source=openshift-route
    - --policy=sync
    - --registry=txt
    - --log-level=debug
    - --zone-id-filter=Z05387772BD5723IZFRX3
    - --openshift-router-name=default
    - --txt-prefix=external-dns-
```

## For OCP 3.11 environment

### Prepare ROUTER_CANONICAL_HOSTNAME in default/router deployment

Read and go through [Finding the Host Name of the Router](https://docs.openshift.com/container-platform/3.11/install_config/router/default_haproxy_router.html#finding-router-hostname).
If no ROUTER_CANONICAL_HOSTNAME is set, you must annotate each route with external-dns.alpha.kubernetes.io/target!

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
        image: registry.k8s.io/external-dns/external-dns:v0.18.0
        args:
        - --source=openshift-route
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
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
- apiGroups: ["route.openshift.io"]
  resources: ["routes"]
  verbs: ["get","watch","list"]
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
        image: registry.k8s.io/external-dns/external-dns:v0.18.0
        args:
        - --source=openshift-route
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

### Verify External DNS works (OpenShift Route example)

The following instructions are based on the
[Hello Openshift](https://github.com/openshift/origin/tree/HEAD/examples/hello-openshift).

#### Install a sample service and expose it

```bash
$ oc apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: hello-openshift
  name: hello-openshift
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-openshift
  template:
    metadata:
      labels:
        app: hello-openshift
    spec:
      containers:
      - image: openshift/hello-openshift
        name: hello-openshift
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: hello-openshift
  name: hello-openshift
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: hello-openshift
  sessionAffinity: None
  type: ClusterIP
---
apiVersion: route.openshift.io/v1
kind: Route
metadata:
  name: hello-openshift
spec:
  host: hello-openshift.example.com
  to:
    kind: Service
    name: hello-openshift
    weight: 100
  wildcardPolicy: None
EOF
```

#### Access the sample route using `curl`

```bash
$ curl -i http://hello-openshift.example.com
HTTP/1.1 200 OK
Date: Fri, 10 Apr 2020 09:36:41 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

Hello OpenShift!
```
