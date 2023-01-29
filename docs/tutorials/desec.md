# Setting up ExternalDNS for Desec

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using Desec.

Make sure to use **master** version of ExternalDNS for this tutorial.

## Creating Desec Credentials

A secret containing the a Desec API token is needed for this provider. You can get a token for your user [here](https://desec.io/tokens).

To create the API token secret you can run `kubectl create secret generic desec-api-key --from-literal=EXTERNAL_DNS_DESEC_API_KEY=<replace-with-your-access-token>`.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.

Besides the API key, it is mandatory to provide a list of DNS zones you want ExternalDNS to manage. The hosted DNS zones will be provides via the `--domain-filter`.

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
        image: registry.k8s.io/external-dns/external-dns:latest
        args:
        - --source=service # ingress is also possible
        - --provider=desec
        - --domain-filter="example.com"
        env:
        - name: EXTERNAL_DNS_DESEC_API_KEY
          valueFrom:
            secretKeyRef:
              key: EXTERNAL_DNS_DESEC_API_KEY
              name: Desec-api-key

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
        image: registry.k8s.io/external-dns/external-dns:latest
        args:
        - --source=service # ingress is also possible
        - --provider=desec
        - --domain-filter=example.com
        env:
        - name: EXTERNAL_DNS_DESEC_API_KEY
          valueFrom:
            secretKeyRef:
              key: EXTERNAL_DNS_DESEC_API_KEY
              name: Desec-api-key
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
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the Desec DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

By setting the TTL annotation on the service, you have to pass a valid TTL, which must be 120 or above.
This annotation is optional, if you won't set it, it will be 1 (automatic) which is 300.

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the Desec DNS records.

## Verifying Desec DNS records

Check your [Desec domain overview](https://desec.io/domains) to view the domains associated with your Desec account. There you can view the records for each domain.

The records should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Desec DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
