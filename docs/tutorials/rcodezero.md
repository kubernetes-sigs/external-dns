# Setting up ExternalDNS for Services on RcodeZero

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using RcodeZero Anycast DNS.

Make sure to use **>=0.5.0** version of ExternalDNS for this tutorial.

## Creating a RcodeZero DNS zone

After logging into RcodeZero Dashboard add a master domain under [RcodeZero Add Zone](https://my.rcodezero.at/domain/create). Use it throughout this guide (substitute example.com).

## Creating RcodeZero Credentials

> The RcodeZero Anycast-Network is provisioned via web interface or REST-API.

RcodeZero API can be enabled and a key generated on [RcodeZero API](https://my.rcodezero.at/enableapi)

The environment var `RC0_API_KEY` will be needed to run ExternalDNS with RcodeZero.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

### Manifest (for clusters without RBAC enabled)

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
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
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
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
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
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

Create a service file called 'nginx.yaml' with the following contents:

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
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

Note the annotation on the service; use the same hostname as the RcodeZero DNS zone created above. The annotation may also be a subdomain
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
the RcodeZero DNS records.

## Verifying RcodeZero DNS records

Check your [RcodeZero Configured Zones](https://my.rcodezero.at/domain) and select the ExternalDNS managed domain.

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage RcodeZero DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```
