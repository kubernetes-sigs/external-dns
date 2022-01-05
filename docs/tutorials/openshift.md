# Configuring ExternalDNS to use the OpenShift Route Source
This tutorial describes how to configure ExternalDNS to use the OpenShift Route source.
It is meant to supplement the other provider-specific setup tutorials.

### For OCP 4.x

In OCP 4.x, if you have multiple ingress controllers then you must specify an ingress controller name or a router name(you can get it from the route's Status.Ingress.RouterName field).
If you don't specify an ingress controller's or router name when you have multiple ingresscontrollers in your environment then the route gets populated with multiple entries of router canonical hostnames which causes external dns to create a CNAME record with multiple router canonical hostnames pointing to the route host which is a violation of RFC 1912 and is not allowed by Cloud Providers which leads to failure of record creation.
Once you specify the ingresscontroller or router name then that will be matched by the external-dns and the router canonical hostname corresponding to this routerName(which is present in route's Status.Ingress.RouterName field) is selected and a CNAME record of this route host pointing to this router canonical hostname is created.

Your externaldns CR shall be created as per the following example.
Replace names in the domain section and zone ID as per your environment.
This is example is for AWS environment.

```yaml

  apiVersion: externaldns.olm.openshift.io/v1alpha1
  kind: ExternalDNS
  metadata:
    name: sample1
  spec:
    domains:
      - filterType: Include
        matchType: Exact
        names: apps.miheer.externaldns
    provider:
      type: AWS
    source:
      hostnameAnnotation: Allow
      openshiftRouteOptions:
        routerName: default
      type: OpenShiftRoute
    zones:
      - Z05387772BD5723IZFRX3

```

This will create an externaldns pod with the following container args under spec in the external-dns namespace where `- --source=openshift-route` and `- --openshift-router-name=default` is added by the external-dns-operator.

```
spec:
  containers:
  - args:
    - --domain-filter=apps.misalunk.externaldns
    - --metrics-address=127.0.0.1:7979
    - --txt-owner-id=external-dns-sample1
    - --provider=aws
    - --source=openshift-route
    - --policy=sync
    - --registry=txt
    - --log-level=debug
    - --zone-id-filter=Z05387772BD5723IZFRX3
    - --openshift-router-name=default
    - --txt-prefix=external-dns-

```

### For OCP 3.11 environment
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
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
  resources: ["services","endpoints","pods"]
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
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
