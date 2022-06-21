# Setting up external-DNS for use with Stackpath

## Contents
- [Introduction](#introduction)
- [Creating a Zone](#creating-a-zone)
- [Stackpath Credentials](#stackpath-credentials)
- [Deploy external-DNS](#deploy-external-dns)
  - [Manifest without RBAC Enabled](#manifest-without-rbac-enabled)
  - [Manifest with RBAC Enabled](#manifest-with-rbac-enabled)
- [Annotate a Service or Ingress](#annotate-a-service-or-ingress)
  - [Annotation Example](#annotation-example)
  - [Nginx Example](#nginx-example)
- [Logging](#logging)
- [Issues](#issues)

***
## Introduction 

This file describes the basic steps to use external-DNS to advertise your
services on Stackpath DNS.

## Creating a Zone 

Before setting up external-DNS, ensure the site and zone you would like
to create records for exist in Stackpath. [Get Started with Stackpath](https://support.stackpath.com/hc/en-us/articles/360037680972)

## Stackpath Credentials

In order to use external-DNS with Stackpath, you need to supply external-DNS
with your Stackpath Stack ID, Stackpath Client ID, and Stackpath Client Secret.
Your Stackpath Stack ID can be found here: [Stackpath - Stacks](https://control.stackpath.com/stacks), and your Client ID and Client Secret
are located here: [Stackpath - API
Management](https://control.stackpath.com/account/api-management).

## Deploy external-DNS

| Required? | Variable Name | Found At |
| - | - | - |
| Yes | Stackpath Stack ID | https://control.stackpath.com/stacks |
| Yes | Stackpath Client ID | https://control.stackpath.com/account/api-management |
| Yes | Stackpath Client Secret | https://control.stackpath.com/account/api-management |

Copying and modifying the proper YAML manifest below, then modify it with your
Stack ID, Client ID, and Client Secret. Then, use the following command to
install external-DNS on your cluster:
```
$ kubectl apply -f externalDnsManifest.yaml
```

### Manifest without RBAC Enabled

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
        image: k8s.gcr.io/external-dns/external-dns:v0.12.0 #USE LATEST STABLE external-dns RELEASE
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (LINE OPTIONAL) limit to only listed domain
        - --provider=stackpath # Identifies provider to use
        env:
        - name: STACKPATH_CLIENT_ID
          value: "YOUR_STACKPATH_CLIENT_ID"
        - name: STACKPATH_CLIENT_SECRET
          value: "YOUR_STACKPATH_CLIENT_SECRET"
        - name: STACKPATH_STACK_ID
          value: "YOUR_STACKPATH_STACK_ID"
```

### Manifest with RBAC Enabled

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
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["networking","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["endpoints"]
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
        image: k8s.gcr.io/external-dns/external-dns:v0.12.0
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (LINE OPTIONAL) limit to only listed domain
        - --provider=stackpath
        env:
        - name: STACKPATH_CLIENT_ID
          value: "YOUR_STACKPATH_CLIENT_ID"
        - name: STACKPATH_CLIENT_SECRET
          value: "YOUR_STACKPATH_CLIENT_SECRET"
        - name: STACKPATH_STACK_ID
          value: "YOUR_STACKPATH_STACK_ID"
```

## Annotate a Service or Ingress

### Annotation Example

Adding the below annotation to any Service or Ingress will cause external-DNS to
create the appropriate records in Stackpath:
```yaml
metadata:
  name: SERVICE_OR_INGRESS_NAME
  annotations:
    external-dns.alpha.kubernetes.io/hostname: NAME.YOURDOMAIN
    external-dns.alpha.kubernetes.io/ttl: "120" #OPTIONAL LINE
```

### Nginx Example

Create a file named nginx.yaml containing the following:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: NAME.YOURDOMAIN
    external-dns.alpha.kubernetes.io/ttl: "120" #OPTIONAL LINE
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Use the below command to install the nginx service on your cluster and advertise
it to Stackpath using external-DNS:
```
$ kubectl apply -f nginx.yaml
```
In a few minutes, check your Stackpath console or run the `kubectl logs` command
explained in the [next section](#logging). You should see the relevent records have been
created, along with two TXT records each to track ownership of the record.

## Logging
Using the below command, you can open a live stream of the logs from
external-DNS, and read what calls are being made to your Stackpath account from
external-DNS:
```
$ kubectl logs -f <external-dns-pod>
```
**The `-f` flag in the above command allows you to follow the logs live.** Omit this
flag to print the logs up to the time of execution and quit.
## Issues

If you have an issue, find a bug, or want to suggest a feature create a new issue [here](https://github.com/kubernetes-sigs/external-dns/issues/new/choose).