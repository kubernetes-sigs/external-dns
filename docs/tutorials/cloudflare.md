# Setting up ExternalDNS for Services on Cloudflare

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using Cloudflare DNS.

Make sure to use **>=0.4.2** version of ExternalDNS for this tutorial.

## Creating a Cloudflare DNS zone

We highly recommend to read this tutorial if you haven't used Cloudflare before:

[Create a Cloudflare account and add a website](https://support.cloudflare.com/hc/en-us/articles/201720164-Step-2-Create-a-Cloudflare-account-and-add-a-website)

## Creating Cloudflare Credentials

Snippet from [Cloudflare - Getting Started](https://api.cloudflare.com/#getting-started-endpoints):

>Cloudflare's API exposes the entire Cloudflare infrastructure via a standardized programmatic interface. Using Cloudflare's API, you can do just about anything you can do on cloudflare.com via the customer dashboard.

>The Cloudflare API is a RESTful API based on HTTPS requests and JSON responses. If you are registered with Cloudflare, you can obtain your API key from the bottom of the "My Account" page, found here: [Go to My account](https://dash.cloudflare.com/profile).

API Token will be preferred for authentication if `CF_API_TOKEN` environment variable is set. 
Otherwise `CF_API_KEY` and `CF_API_EMAIL` should be set to run ExternalDNS with Cloudflare.

When using API Token authentication the token should be granted Zone `Read` and DNS `Edit` privileges.

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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=cloudflare
        - --cloudflare-proxied # (optional) enable the proxy feature of Cloudflare (DDOS protection, CDN...)
        env:
        - name: CF_API_KEY
          value: "YOUR_CLOUDFLARE_API_KEY"
        - name: CF_API_EMAIL
          value: "YOUR_CLOUDFLARE_EMAIL"
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
        - --provider=cloudflare
        - --cloudflare-proxied # (optional) enable the proxy feature of Cloudflare (DDOS protection, CDN...)
        env:
        - name: CF_API_KEY
          value: "YOUR_CLOUDFLARE_API_KEY"
        - name: CF_API_EMAIL
          value: "YOUR_CLOUDFLARE_EMAIL"
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

Note the annotation on the service; use the same hostname as the Cloudflare DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

By setting the TTL annotation on the service, you have to pass a valid TTL, which must be 120 or above.
This annotation is optional, if you won't set it, it will be 1 (automatic) which is 300.
For Cloudflare proxied entries, set the TTL annotation to 1 (automatic), or do not set it.

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the Cloudflare DNS records.

## Verifying Cloudflare DNS records

Check your [Cloudflare dashboard](https://www.cloudflare.com/a/dns/example.com) to view the records for your Cloudflare DNS zone.

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Cloudflare DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```

## Setting cloudflare-proxied on a per-ingress basis

Using the `external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"` annotation on your ingress, you can specify if the proxy feature of Cloudflare should be enabled for that record. This setting will override the global `--cloudflare-proxied` setting.
