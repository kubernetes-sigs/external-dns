# Setting up ExternalDNS for Services on IONOS

This tutorial describes how to setup ExternalDNS for use within a Kubernetes cluster using IONOS DNS.

Make sure to use **>=0.7** version of ExternalDNS for this tutorial.

## Create an API key

Instructions for creating an API key are [here](https://developer.hosting.ionos.de/docs/getstarted). The API key will be added to the environment variable 'IONOS_API_KEY' via one of the Manifest templates (as described below).

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with. Before applying, modify the Manifest as follows:

- Replace "example.com" with the domain name you want to update.
- Replace "YOUR_IONOS_API_KEY" with the API key created above.

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
          image: k8s.gcr.io/external-dns/external-dns:latest
          args:
            - --source=service # ingress is also possible
            - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
            - --provider=ionos
          env:
            - name: IONOS_API_KEY
              value: "YOUR_IONOS_API_KEY"
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
    resources: ["services", "endpoints", "pods"]
    verbs: ["get", "watch", "list"]
  - apiGroups: ["extensions", "networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get", "watch", "list"]
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
      containers:
        - name: external-dns
          image: k8s.gcr.io/external-dns/external-dns:latest
          args:
            - --source=service # ingress is also possible
            - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
            - --provider=ionos
          env:
            - name: IONOS_API_KEY
              value: "YOUR_IONOS_API_KEY"
```

## Deploying an Nginx Service

After you have deployed ExternalDNS, you can deploy a simple service based on Nginx to test the setup. This is optional, though highly recommended before using ExternalDNS in production.

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

Change the file as follows:

- Replace the annotation of the service; use the same hostname as the DNS zone. The annotation may also be a subdomain of the DNS zone (e.g. 'www.example.com').
- Set the TTL annotation of the service. A valid TTL of 60 or above must be given. This annotation is optional, and defaults to "3600" if no value is given.

These annotations will be used to determine what services should be registered with DNS. Removing these annotations will cause ExternalDNS to remove the corresponding DNS records.

Create the Deployment and Service:

```bash
$ kubectl create -f nginx.yaml
```

Depending on your cloud provider, it might take a while to create an external IP for the service. Once an external IP address is assigned to the service, ExternalDNS will notice the new address and synchronize the IONOS DNS records accordingly.

## Verifying IONOS DNS records

Check the zone records in the IONOS Control Panel. The zone should now contain the external IP address of the service as an A record.

## Cleanup

Once you have verified that ExternalDNS successfully manages IONOS DNS records for external services, you can delete the tutorial example as follows:

```bash
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
