# Setting up ExternalDNS for Services on Bizfly Cloud

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using Bizfly Cloud DNS.

Make sure to use **>0.13.5** version of ExternalDNS for this tutorial.

## Creating Bizfly Cloud Credentials

>The Bizfly Cloud API is a RESTful API based on HTTPS requests and JSON responses. If you are registered with Bizfly Cloud, you can create your credentials from [here](https://manage.bizflycloud.vn/account/configuration/credential).

API Token will authentication if `BFC_APP_CREDENTIAL_ID` and `BFC_APP_CREDENTIAL_SECRET` environment variable is set.

If you would like to further restrict the API permissions to a specific zone (or zones), you also need to use the `--zone-id-filter` so that the underlying API requests only access the zones that you explicitly specify, as opposed to accessing all zones.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
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
        image: registry.k8s.io/external-dns/external-dns:v0.13.6
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains
        - --zone-id-filter=1ef149d0-cefa-4477-9161-0e1dff34dc10 # (optional) limit to a specific zone
        - --provider=bizflycloud
        - --bizflycloud-api-page-size=1000 # (optional) configure how many DNS records to fetch per request
        env:
        - name: BFC_APP_CREDENTIAL_ID
          value: "YOUR_BIZFLY_CLOUD_CREDENTIAL_ID"
        - name: BFC_APP_CREDENTIAL_SECRET
          value: "YOUR_BIZFLY_CLOUD_CREDENTIAL_SECRET"
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
  verbs: ["list", "watch"]
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
        image: registry.k8s.io/external-dns/external-dns:v0.13.6
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains
        - --zone-id-filter=1ef149d0-cefa-4477-9161-0e1dff34dc10 # (optional) limit to a specific zone
        - --provider=bizflycloud
        - --bizflycloud-api-page-size=1000 # (optional) configure how many DNS records to fetch per request
        env:
        - name: BFC_APP_CREDENTIAL_ID
          value: "YOUR_BIZFLY_CLOUD_CREDENTIAL_ID"
        - name: BFC_APP_CREDENTIAL_SECRET
          value: "YOUR_BIZFLY_CLOUD_CREDENTIAL_SECRET"
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

Note the annotation on the service; use the same hostname as the Bizfly Cloud DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

By setting the TTL annotation on the service, you have to pass a valid TTL, which must be 5 or above.
This annotation is optional, if you won't set it, it will be default 60.

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the Bizfly Cloud DNS records.

## Verifying Bizfly Cloud DNS records

Select your zone at [Bizfly Cloud dashboard](https://manage.bizflycloud.vn/dns) to view the records for your Bizfly Cloud DNS zone.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Bizfly Cloud DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
