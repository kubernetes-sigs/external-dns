# Setting up ExternalDNS on Google Container Engine

This tutorial describes how to setup ExternalDNS for usage within a GKE cluster. Make sure to use **>=0.4** version of ExternalDNS for this tutorial

## Set up your environment

*If you prefer to try-out ExternalDNS in one of the existing environments you can skip this step*

Setup your environment to work with Google Cloud Platform. Fill in your values as needed, e.g. target project.

```console
$ gcloud config set project "zalando-external-dns-test"
$ gcloud config set compute/region "europe-west1"
$ gcloud config set compute/zone "europe-west1-d"
```

Create a GKE cluster.

```console
$ gcloud container clusters create "external-dns" \
    --num-nodes 1 \
    --scopes "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
```

Create a DNS zone which will contain the managed DNS records.

```console
$ gcloud dns managed-zones create "external-dns-test-gcp-zalan-do" \
    --dns-name "external-dns-test.gcp.zalan.do." \
    --description "Automatically managed zone by kubernetes.io/external-dns"
```

Make a note of the nameservers that were assigned to your new zone.

```console
$ gcloud dns record-sets list \
    --zone "external-dns-test-gcp-zalan-do" \
    --name "external-dns-test.gcp.zalan.do." \
    --type NS
NAME                             TYPE  TTL    DATA
external-dns-test.gcp.zalan.do.  NS    21600  ns-cloud-e1.googledomains.com.,ns-cloud-e2.googledomains.com.,ns-cloud-e3.googledomains.com.,ns-cloud-e4.googledomains.com.
```

In this case it's `ns-cloud-{e1-e4}.googledomains.com.` but your's could slightly differ, e.g. `{a1-a4}`, `{b1-b4}` etc.

Tell the parent zone where to find the DNS records for this zone by adding the corresponding NS records there. Assuming the parent zone is "gcp-zalan-do" and the domain is "gcp.zalan.do" and that it's also hosted at Google we would do the following.

```console
$ gcloud dns record-sets transaction start --zone "gcp-zalan-do"
$ gcloud dns record-sets transaction add ns-cloud-e{1..4}.googledomains.com. \
    --name "external-dns-test.gcp.zalan.do." --ttl 300 --type NS --zone "gcp-zalan-do"
$ gcloud dns record-sets transaction execute --zone "gcp-zalan-do"
```

## Deploy ExternalDNS

### Role-Based Access Control (RBAC)

[RBAC]("https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control") is enabled by default on all Container clusters which are running Kubernetes version 1.6 or higher.

Because of the way Container Engine checks permissions when you create a Role or ClusterRole, you must first create a RoleBinding that grants you all of the permissions included in the role you want to create.

```console
kubectl create clusterrolebinding your-user-cluster-admin-binding --clusterrole=cluster-admin --user=your.google.cloud.email@example.org
```

Connect your `kubectl` client to the cluster you just created.

```console
gcloud container clusters get-credentials "external-dns"
```

Then apply one of the following manifests file to deploy ExternalDNS.

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
        - --domain-filter=external-dns-test.gcp.zalan.do # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=google
#        - --google-project=zalando-external-dns-test # Use this to specify a project different from the one external-dns is running inside
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
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
  verbs: ["get", "watch", "list"]
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
        - --domain-filter=external-dns-test.gcp.zalan.do # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=google
#        - --google-project=zalando-external-dns-test # Use this to specify a project different from the one external-dns is running inside
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --registry=txt
        - --txt-owner-id=my-identifier
```

Use `--dry-run` if you want to be extra careful on the first run. Note, that you will not see any records created when you are running in dry-run mode. You can, however, inspect the logs and watch what would have been done.


## Verify ExternalDNS works

Create the following sample application to test that ExternalDNS works.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.gcp.zalan.do.
spec:
  type: LoadBalancer
  ports:
  - port: 80
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
```

After roughly two minutes check that a corresponding DNS record for your service was created.

```console
$ gcloud dns record-sets list \
    --zone "external-dns-test-gcp-zalan-do" \
    --name "nginx.external-dns-test.gcp.zalan.do."

NAME                                   TYPE  TTL  DATA
nginx.external-dns-test.gcp.zalan.do.  A     300  104.155.60.49
nginx.external-dns-test.gcp.zalan.do.  TXT   300  "heritage=external-dns,external-dns/owner=my-identifier"
```

Note created TXT record alongside A record. TXT record signifies that the corresponding A record is managed by ExternalDNS. This makes ExternalDNS safe for running in environments where there are other records managed via other means.

Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.

```console
$ dig +short @ns-cloud-e1.googledomains.com. nginx.external-dns-test.gcp.zalan.do.
104.155.60.49
```

Given you hooked up your DNS zone with its parent zone you can use `curl` to access your site.

```console
$ curl nginx.external-dns-test.gcp.zalan.do
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
...
</head>
<body>
...
</body>
</html>
```

Let's check that Ingress works as well. Create the following Ingress.

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: nginx
spec:
  rules:
  - host: via-ingress.external-dns-test.gcp.zalan.do
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
```

Again, after roughly two minutes check that a corresponding DNS record for your Ingress was created.

```console
$ gcloud dns record-sets list \
    --zone "external-dns-test-gcp-zalan-do" \
    --name "via-ingress.external-dns-test.gcp.zalan.do." \

NAME                                         TYPE  TTL  DATA
via-ingress.external-dns-test.gcp.zalan.do.  A     300  130.211.46.224
via-ingress.external-dns-test.gcp.zalan.do.  TXT   300  "heritage=external-dns,external-dns/owner=my-identifier"
```

Let's check that we can resolve this DNS name as well.

```console
dig +short @ns-cloud-e1.googledomains.com. via-ingress.external-dns-test.gcp.zalan.do.
130.211.46.224
```

Try with `curl` as well.

```console
$ curl via-ingress.external-dns-test.gcp.zalan.do
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
...
</head>
<body>
...
</body>
</html>
```

## Clean up

Make sure to delete all Service and Ingress objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
$ kubectl delete ingress nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the managed zone and cluster.

```console
$ gcloud dns managed-zones delete "external-dns-test-gcp-zalan-do"
$ gcloud container clusters delete "external-dns"
```

Also delete the NS records for your removed zone from the parent zone.

```console
$ gcloud dns record-sets transaction start --zone "gcp-zalan-do"
$ gcloud dns record-sets transaction remove ns-cloud-e{1..4}.googledomains.com. \
    --name "external-dns-test.gcp.zalan.do." --ttl 300 --type NS --zone "gcp-zalan-do"
$ gcloud dns record-sets transaction execute --zone "gcp-zalan-do"
```

### User Demo How-To Blogs and Examples
* A full demo on GKE Kubernetes + CloudDNS + SA-Permissions [How-to Kubernetes with DNS management (ssl-manager pre-req)](https://medium.com/@jpantjsoha/how-to-kubernetes-with-dns-management-for-gitops-31239ea75d8d)
