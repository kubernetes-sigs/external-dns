# Setting up ExternalDNS for Oracle Cloud Infrastructure (OCI)

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using OCI DNS.

Make sure to use the latest version of ExternalDNS for this tutorial.

## Creating an OCI DNS Zone

Create a DNS zone which will contain the managed DNS records. Let's use `example.com` as an reference here.

For more information about OCI DNS see the documentation [here][1].

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
We first need to create a config file containing the information needed to connect with the OCI API.

Create a new file (oci.yaml) and modify the contents to match the example below. Be sure to adjust the values to match your own credentials:

```yaml
auth:
  region: us-phoenix-1
  tenancy: ocid1.tenancy.oc1...
  user: ocid1.user.oc1...
  key: |
    -----BEGIN RSA PRIVATE KEY-----
    -----END RSA PRIVATE KEY-----
  fingerprint: af:81:71:8e...
compartment: ocid1.compartment.oc1...
```

Create a secret using the config file above:

```shell
$ kubectl create secret generic external-dns-config --from-file=oci.yaml
```

### Manifest (for clusters with RBAC enabled)

Apply the following manifest to deploy ExternalDNS.

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
        - --provider=oci
        - --policy=upsert-only # prevent ExternalDNSfrom deleting any records, omit to enable full synchronization
        - --txt-owner-id=my-identifier
        volumeMounts:
          - name: config
            mountPath: /etc/kubernetes/
      volumes:
      - name: config
        secret:
          secretName: external-dns-config
```

## Verify ExternalDNS works (Service example)

Create the following sample application to test that ExternalDNS works.

> For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: example.com
spec:
  type: LoadBalancer
  ports:
  - port: 80
    name: http
    targetPort: 80
  selector:
    app: nginx
---
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
          name: http
```

Apply the manifest above and wait roughly two minutes and check that a corresponding DNS record for your service was created.

```
$ kubectl apply -f nginx.yaml
```

[1]: https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm
