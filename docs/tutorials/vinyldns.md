# Setting up ExternalDNS for VinylDNS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using VinylDNS.

The environment vars `VINYLDNS_ACCESS_KEY`, `VINYLDNS_SECRET_KEY`, and `VINYLDNS_HOST` will be needed to run ExternalDNS with VinylDNS.

## Create a sample deployment and service for external-dns to use

Run an application and expose it via a Kubernetes Service:

```console
$ kubectl run nginx --image=nginx --replicas=1 --port=80
$ kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
```

Annotate the Service with your desired external DNS name. Make sure to change `example.org` to your domain.

```console
$ kubectl annotate service nginx "external-dns.alpha.kubernetes.io/hostname=nginx.example.org."
```

After the service is up and running, it should get an EXTERNAL-IP. At first this may showing as `<pending>`

```console
$ kubectl get svc
NAME         CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
kubernetes   10.0.0.1     <none>        443/TCP        1h
nginx        10.0.0.115   <pending>     80:30543/TCP   10s
```

Once it's available

```console
% kubectl get svc
NAME         CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
kubernetes   10.0.0.1     <none>        443/TCP        1h
nginx        10.0.0.115   34.x.x.x      80:30543/TCP   2m
```

## Deploy ExternalDNS to Kubernetes

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

**Note for examples below**

When using `registry=txt` option, make sure to also use the `txt-prefix` and `txt-owner-id` options as well. If you try to create a `TXT` record in VinylDNS without a prefix, it will try to create a `TXT` record with the same name as your actual DNS record and fail (creating a stranded record `external-dns` cannot manage).

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
        - --provider=vinyldns
        - --source=service
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --registry=txt
        - --txt-owner-id=grizz
        - --txt-prefix=txt-
        env:
        - name: VINYLDNS_HOST
          value: "YOUR_VINYLDNS_HOST"
        - name: VINYLDNS_ACCESS_KEY
          value: "YOUR_VINYLDNS_ACCESS_KEY"
        - name: VINYLDNS_SECRET_KEY
          value: "YOUR_VINYLDNS_SECRET_KEY"
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
        - --provider=vinyldns
        - --source=service
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --registry=txt
        - --txt-owner-id=grizz
        - --txt-prefix=txt-
        env:
        env:
        - name: VINYLDNS_HOST
          value: "YOUR_VINYLDNS_HOST"
        - name: VINYLDNS_ACCESS_KEY
          value: "YOUR_VINYLDNS_ACCESS_KEY"
        - name: VINYLDNS_SECRET_KEY
          value: "YOUR_VINYLDNS_SECRET_KEYY
```

## Running a locally built version pointed to the above nginx service
Make sure your kubectl is configured correctly. Assuming you have the sources, build and run it like below.

The vinyl access details needs to exported to the environment before running.

```bash
make
# output skipped

export VINYLDNS_HOST=<fqdn of vinyl dns api>
export VINYLDNS_ACCESS_KEY=<access key>
export VINYLDNS_SECRET_KEY=<secret key>

./build/external-dns \
    --provider=vinyldns \
    --source=service \
    --domain-filter=elements.capsps.comcast.net. \
    --zone-id-filter=20e8bfd2-3a70-4e1b-8e11-c9c1948528d3 \
    --registry=txt \
    --txt-owner-id=grizz \
    --txt-prefix=txt- \
    --namespace=default \
    --once \
    --dry-run \
    --log-level debug

INFO[0000] running in dry-run mode. No changes to DNS records will be made.
INFO[0000] Created Kubernetes client https://some-k8s-cluster.example.com
INFO[0001] Zone: [nginx.example.org.]
# output skipped
```

Having `--dry-run=true` and `--log-level=debug` is a great way to see _exactly_ what VinylDNS is doing or is about to do.
