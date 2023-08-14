# Setting up ExternalDNS for Services on Domeneshop

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using Domeneshop DNS (also known as dominanameshop.com).

Make sure to use **>=0.13.6** version of ExternalDNS for this tutorial.

## Managing DNS with Domeneshop

If you want to learn about Domeneshop DNS read the following documentation:

- [Domeneshop DNS documentation](https://domainname.shop/dns?currency=USD&lang=en)
- [Domeneshop FAQ](https://domainname.shop/faq?currency=USD&lang=en)

## Creating Domeneshop credentials

Generate a new oauth token by following the instructions at [Access-and-Authentication](https://api.domeneshop.no/docs/#section/Authentication)

The environment variables `DOMENESHOP_API_TOKEN` and `DOMENESHOP_API_SECRET` will be needed to run ExternalDNS with Domeneshop.

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
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=domeneshop
        - --txt-prefix txt # only required if using the txt registry
        env:
        - name: DOMENESHOP_API_TOKEN
          value: "YOUR_DOMENESHOP_API_TOKEN"
        - name: DOMENESHOP_API_SECRET
          value: "YOUR_DOMENESHOP_API_SECRET"
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
        image: registry.k8s.io/external-dns/external-dns:v0.13.5
        args:
          - --source=service # ingress is also possible
          - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
          - --provider=domeneshop
          - --txt-prefix txt # only required if using the txt registry
        env:
          - name: DOMENESHOP_API_TOKEN
            value: "YOUR_DOMENESHOP_API_TOKEN"
          - name: DOMENESHOP_API_SECRET
            value: "YOUR_DOMENESHOP_API_SECRET"
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
    external-dns.alpha.kubernetes.io/hostname: my-app.example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the Domeneshop domain.

ExternalDNS uses this annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```console
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize the Domeneshop DNS records.

## Verifying Domeneshop DNS records

Check the [Domeneshop UI](https://domene.shop/admin?view=domains) to view the records for your Domeneshop domains.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Domeneshop DNS records, we can delete the tutorial's example:

```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
```
