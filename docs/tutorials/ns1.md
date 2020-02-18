# Setting up ExternalDNS for Services on NS1

This tutorial describes how to setup ExternalDNS for use within a
Kubernetes cluster using NS1 DNS.

Make sure to use **>=0.5** version of ExternalDNS for this tutorial.

## Creating a zone with NS1 DNS

If you are new to NS1, we recommend you first read the following
instructions for creating a zone.

[Creating a zone using the NS1
portal](https://ns1.com/knowledgebase/creating-a-zone)

[Creating a zone using the NS1
API](https://ns1.com/api#put-create-a-new-dns-zone)

## Creating NS1 Credentials

All NS1 products are API-first, meaning everything that can be done on
the portal---including managing zones and records, data sources and
feeds, and account settings and users---can be done via API.

The NS1 API is a standard REST API with JSON responses. The environment
var `NS1_APIKEY` will be needed to run ExternalDNS with NS1.

### To add or delete an API key

1.  Log into the NS1 portal at [my.nsone.net](http://my.nsone.net).

2.  Click your username in the upper-right corner, and navigate to **Account Settings** \> **Users & Teams**.

3.  Navigate to the _API Keys_ tab, and click **Add Key**.

4.  Enter the name of the application and modify permissions and settings as desired. Once complete, click **Create Key**. The new API key appears in the list.

    Note: Set the permissions for your API keys just as you would for a user or team associated with your organization's NS1 account. For more information, refer to the article [Creating and Managing API Keys](https://help.ns1.com/hc/en-us/articles/360026140094-Creating-managing-users) in the NS1 Knowledge Base.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster with which you want to test ExternalDNS, and then apply one of the following manifest files for deployment:

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
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ns1
        env:
        - name: NS1_APIKEY
          value: "YOUR_NS1_API_KEY"
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
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ns1
        env:
        - name: NS1_APIKEY
          value: "YOUR_NS1_API_KEY"
```

## Deploying an Nginx Service

Create a service file called 'nginx.yaml' with the following contents:

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
    external-dns.alpha.kubernetes.io/hostname: example.com
    external-dns.alpha.kubernetes.io/ttl: "120" #optional
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

**A note about annotations**

Verify that the annotation on the service uses the same hostname as the NS1 DNS zone created above. The annotation may also be a subdomain of the DNS zone (e.g. 'www.example.com').

The TTL annotation can be used to configure the TTL on DNS records managed by ExternalDNS and is optional. If this annotation is not set, the TTL on records managed by ExternalDNS will default to 10.

ExternalDNS uses the hostname annotation to determine which services should be registered with DNS. Removing the hostname annotation will cause ExternalDNS to remove the corresponding DNS records.

### Create the deployment and service

```
$ kubectl create -f nginx.yaml
```

Depending on where you run your service, it may take some time for your cloud provider to create an external IP for the service. Once an external IP is assigned, ExternalDNS detects the new service IP address and synchronizes the NS1 DNS records.

## Verifying NS1 DNS records

Use the NS1 portal or API to verify that the A record for your domain shows the external IP address of the services.

## Cleanup

Once you successfully configure and verify record management via ExternalDNS, you can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
