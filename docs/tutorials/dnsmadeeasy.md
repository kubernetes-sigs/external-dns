# Setting up ExternalDNS for Services on DNSMadeEasy


This tutorial describes how to setup ExternalDNS for usage with DNSMadeEasy.

Make sure to use **>=0.5.0** version of ExternalDNS for this tutorial.

## Retriever your DNSMadeEasy API Access Token

A DNSMadeEasy API key can be acquired by following the [provided documentation from DNSMadeEasy](https://support.dnsmadeeasy.com/support/solutions/articles/47001131906-finding-or-generating-api-keys-in-dns-made-easy)

The following environment variables are used to configure API Authentication:

| Name | Description | Required    | Default  |
| ---- |------------ |-------------|----------|
| `dme_apikey` | DNSMadeEasy API Key | Yes         |          |
| `dme_secretkey` | DNSMadeEasy Secret Key | Yes         |          |
| `dme_insecure` | Ignore SSL Validation Errors | No |  `false` |

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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone you create in DNSMadeEasy, or remove to manage all
        - --provider=dnsmadeeasy
        - --registry=txt
        env:
        - name: dme_apikey
          value: "YOUR_DNSMADEEASY_APIKEY"
        - name: dme_secretkey
          value: "YOUR_DNSMADEEASY_SECRETKEY"
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
      containers:
        - name: external-dns
          image: k8s.gcr.io/external-dns/external-dns:v0.7.6
          args:
            - --source=service
            - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone you create in DNSMadeEasy, or remove to manage all
            - --provider=dnsmadeeasy
            - --registry=txt
          env:
            - name: dme_apikey
              value: "YOUR_DNSMADEEASY_APIKEY"
            - name: dme_secretkey
              value: "YOUR_DNSMADEEASY_SECRETKEY"
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
    external-dns.alpha.kubernetes.io/hostname: validate-external-dns.example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the DNSMadeEasy DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```sh
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service. Check the status by running
`kubectl get services nginx`.  If the `EXTERNAL-IP` field shows an address, the service is ready to be accessed externally.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the DNSMadeEasy DNS records.

## Verifying DNSMadeEasy DNS records

### Looking at the DNSMadeEasy Control Panel

You can view your DNSMadeEasy Control Panel at https://cp.dnsmadeeasy.com.

## Clean up

Now that we have verified that ExternalDNS will automatically manage DNSMadeEasy DNS records, we can delete the tutorial's example:

```sh
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```

### Deleting Created Records

#### Using the DNSMadeEasy Control Panel

You can delete records, in bulk, in the DNSMadeEasy Control Panel at https://cp.dnsmadeeasy.com.