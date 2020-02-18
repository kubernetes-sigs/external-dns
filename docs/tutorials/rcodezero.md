# Setting up ExternalDNS for Services on RcodeZero

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using [RcodeZero Anycast DNS](https://www.rcodezero.at). Make sure to use **>=0.5.0** version of ExternalDNS for this tutorial.

The following steps are required to use RcodeZero with ExternalDNS:

1. Sign up for an RcodeZero account (or use an existing account).
2. Add your zone to the RcodeZero DNS
3. Enable the RcodeZero API, and generate an API key.
4. Deploy ExternalDNS to use the RcodeZero provider.
5. Verify the setup bey deploying a test services (optional)

## Creating a RcodeZero DNS zone

Before records can be added to your domain name automatically, you need to add your domain name to the set of zones managed by RcodeZero. In order to add the zone, perform the following steps:

1. Log in to the RcodeZero Dashboard, and move to the [Add Zone](https://my.rcodezero.at/domain/create) page.
2. Select "MASTER" as domain type, and add your domain name there. Use this domain name instead of "example.com" throughout the rest of this tutorial. 

Note that "SECONDARY" domains cannot be managed by ExternalDNS, because this would not allow modification of records in the zone.

## Enable the API, and create Credentials

> The RcodeZero Anycast-Network is provisioned via web interface or REST-API.

Enable the RcodeZero API to generate an API key on [RcodeZero API](https://my.rcodezero.at/enableapi). The API key will be added to the environment variable 'RC0_API_KEY' via one of the Manifest templates (as described below).

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with. Choose a Manifest from below, depending on whether or not you have RBAC enabled. Before applying it, modify the Manifest as follows:

- Replace "example.com" with the domain name you added to RcodeZero.
- Replace YOUR_RCODEZERO_API_KEY with the API key created above.
- Replace YOUR_ENCRYPTION_KEY_STRING with a string to encrypt the TXT records

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
        - --provider=rcodezero
        - --rc0-enc-txt # (optional) encrypt TXT records; encryption key has to be provided with RC0_ENC_KEY env var.
        env:
        - name: RC0_API_KEY
          value: "YOUR_RCODEZERO_API_KEY"
        - name: RC0_ENC_VAR
          value: "YOUR_ENCRYPTION_KEY_STRING"
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
        - --provider=rcodezero
        - --rc0-enc-txt # (optional) encrypt TXT records; encryption key has to be provided with RC0_ENC_KEY env var.
        env:
        - name: RC0_API_KEY
          value: "YOUR_RCODEZERO_API_KEY"
        - name: RC0_ENC_VAR
          value: "YOUR_ENCRYPTION_KEY_STRING"
```

## Deploying an Nginx Service

After you have deployed ExternalDNS with RcodeZero, you can deploy a simple service based on Nginx to test the setup. This is optional, though highly recommended before using ExternalDNS in production.

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

- Replace the annotation of the service; use the same hostname as the RcodeZero DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').
- Set the TTL annotation of the service. A valid TTL of 120 or above must be given. This annotation is optional, and defaults to "300" if no value is given.

These annotations will be used to determine what services should be registered with DNS. Removing these annotations will cause ExternalDNS to remove the corresponding DNS records.

Create the Deployment and Service:

```bash
$ kubectl create -f nginx.yaml
```

Depending on your cloud provider, it might take a while to create an external IP for the service. Once an external IP address is assigned to the service, ExternalDNS will notice the new address and synchronize the RcodeZero DNS records accordingly.

## Verifying RcodeZero DNS records

Check your [RcodeZero Configured Zones](https://my.rcodezero.at/domain) and select the respective zone name. The zone should now contain the external IP address of the service as an A record.

## Cleanup

Once you have verified that ExternalDNS successfully manages RcodeZero DNS records for external services, you can delete the tutorial example as follows:

```bash
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
