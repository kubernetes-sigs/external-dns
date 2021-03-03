# Setting up ExternalDNS for Services on OVH

This tutorial describes how to setup ExternalDNS for use within a
Kubernetes cluster using OVH DNS.

Make sure to use **>=0.6** version of ExternalDNS for this tutorial.

## Creating a zone with OVH DNS

If you are new to OVH, we recommend you first read the following
instructions for creating a zone.

[Creating a zone using the OVH manager](https://docs.ovh.com/gb/en/domains/create_a_dns_zone_for_a_domain_which_is_not_registered_at_ovh/)

[Creating a zone using the OVH API](https://api.ovh.com/console/)

## Creating OVH Credentials

You first need to create an OVH application.

Using the [OVH documentation](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/#creation-of-your-application-keys) you will have your `Application key` and `Application secret`

And you will need to generate your consumer key, here the permissions needed :
- GET on `/domain/zone`
- GET on `/domain/zone/*/record`
- GET on `/domain/zone/*/record/*`
- POST on `/domain/zone/*/record`
- DELETE on `/domain/zone/*/record/*`
- POST on `/domain/zone/*/refresh`

You can use the following `curl` request to generate & validated your `Consumer key`

```bash
curl -XPOST -H "X-Ovh-Application: <ApplicationKey>" -H "Content-type: application/json" https://eu.api.ovh.com/1.0/auth/credential -d '{
  "accessRules": [
    {
      "method": "GET",
      "path": "/domain/zone"
    },
    {
      "method": "GET",
      "path": "/domain/zone/*/record"
    },
    {
      "method": "GET",
      "path": "/domain/zone/*/record/*"
    },
    {
      "method": "POST",
      "path": "/domain/zone/*/record"
    },
    {
      "method": "DELETE",
      "path": "/domain/zone/*/record/*"
    },
    {
      "method": "POST",
      "path": "/domain/zone/*/refresh"
    }
  ],
  "redirection":"https://github.com/kubernetes-sigs/external-dns/blob/HEAD/docs/tutorials/ovh.md#creating-ovh-credentials"
}'
```

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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ovh
        env:
        - name: OVH_APPLICATION_KEY
          value: "YOUR_OVH_APPLICATION_KEY"
        - name: OVH_APPLICATION_SECRET
          value: "YOUR_OVH_APPLICATION_SECRET"
        - name: OVH_CONSUMER_KEY
          value: "YOUR_OVH_CONSUMER_KEY_AFTER_VALIDATED_LINK"
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
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["get","watch","list"]
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ovh
        env:
        - name: OVH_APPLICATION_KEY
          value: "YOUR_OVH_APPLICATION_KEY"
        - name: OVH_APPLICATION_SECRET
          value: "YOUR_OVH_APPLICATION_SECRET"
        - name: OVH_CONSUMER_KEY
          value: "YOUR_OVH_CONSUMER_KEY_AFTER_VALIDATED_LINK"
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

Verify that the annotation on the service uses the same hostname as the OVH DNS zone created above. The annotation may also be a subdomain of the DNS zone (e.g. 'www.example.com').

The TTL annotation can be used to configure the TTL on DNS records managed by ExternalDNS and is optional. If this annotation is not set, the TTL on records managed by ExternalDNS will default to 10.

ExternalDNS uses the hostname annotation to determine which services should be registered with DNS. Removing the hostname annotation will cause ExternalDNS to remove the corresponding DNS records.

### Create the deployment and service

```
$ kubectl create -f nginx.yaml
```

Depending on where you run your service, it may take some time for your cloud provider to create an external IP for the service. Once an external IP is assigned, ExternalDNS detects the new service IP address and synchronizes the OVH DNS records.

## Verifying OVH DNS records

Use the OVH manager or API to verify that the A record for your domain shows the external IP address of the services.

## Cleanup

Once you successfully configure and verify record management via ExternalDNS, you can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
