# Setting up ExternalDNS for Services on Vultr

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using Vultr DNS.

Make sure to use **>=0.6** version of ExternalDNS for this tutorial.

## Managing DNS with Vultr

If you want to read up on vultr DNS service you can read the following tutorial: 
[Introduction to Vultr DNS](https://www.vultr.com/docs/introduction-to-vultr-dns)

Create a new DNS Zone where you want to create your records in. For the examples we will be using `example.com`

## Creating Vultr Credentials

You will need to create a new API Key which can be found on the [Vultr Dashboard](https://my.vultr.com/settings/#settingsapi).

The environment variable `VULTR_API_KEY` will be needed to run ExternalDNS with Vultr.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.

Begin by creating a Kubernetes secret to securely store your  Akamai Edge DNS Access Tokens. This key will enable ExternalDNS to authenticate with Akamai Edge DNS:

```shell
kubectl create secret generic VULTR_API_KEY --from-literal=VULTR_API_KEY=YOUR_VULTR_API_KEY
```

Ensure to replace YOUR_VULTR_API_KEY, with your actual Vultr API key.


Then apply one of the following manifests file to deploy ExternalDNS.

### Using Helm

reate a values.yaml file to configure ExternalDNS to use Akamai Edge DNS as the DNS provider. This file should include the necessary environment variables:

```shell
provider:
  name: akamai
env:
  - name: VULTR_API_KEY
    valueFrom:
      secretKeyRef:
        name: VULTR_API_KEY
        key: VULTR_API_KEY
```

Finally, install the ExternalDNS chart with Helm using the configuration specified in your values.yaml file:

```shell
helm upgrade --install external-dns external-dns/external-dns --values values.yaml
```

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
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=vultr
        env:
        - name: VULTR_API_KEY
          valueFrom:
            secretKeyRef:
              name: VULTR_API_KEY
              key: VULTR_API_KEY
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
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=vultr
        env:
        - name: VULTR_API_KEY
          valueFrom:
            secretKeyRef:
              name: VULTR_API_KEY
              key: VULTR_API_KEY
```

## Deploying a Nginx Service

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

Note the annotation on the service; use the same hostname as the Vultr DNS zone created above.

ExternalDNS uses this annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```console
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize the Vultr DNS records.

## Verifying Vultr DNS records

Check your [Vultr UI](https://my.vultr.com/dns/) to view the records for your Vultr DNS zone.

Click on the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Vultr DNS records, we can delete the tutorial's example:

```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
```
