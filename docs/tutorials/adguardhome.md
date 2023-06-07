# Setting up ExternalDNS for AdGuard Home

This tutorial describes how to setup ExternalDNS to sync records with AdGuard Home's Custom DNS.
AdGuard Home has an internal list it checks last when resolving requests. This list can contain any number of arbitrary A, AAAA or CNAME records.
There is a pseudo-API exposed that ExternalDNS is able to use to manage these records.

## Minimum Required AdGuard Home Version

Since this provider is built on the [new `/control/rewrite/update` endpoint in AdGuard Home](https://github.com/AdguardTeam/AdGuardHome/commit/0393e4109624395bb97af146f2d0e48ea3d7c37b), it currently requires running an edge version or, once available, a v0.108+ release.

## Deploy ExternalDNS

You can skip to the [manifest](#externaldns-manifest) if authentication is disabled on your AdGuard Home instance or you don't want to use secrets.

If your AdGuard Home server's admin dashboard is protected by a password, you'll likely want to create a secret first containing its value.
This is optional since you _do_ retain the option to pass it as a flag with `--adguardhome-password`.

You can create the secret with:

```bash
kubectl create secret generic adguardhome-password \
    --from-literal EXTERNAL_DNS_ADGUARDHOME_PASSWORD=supersecret
```

Replacing **"supersecret"** with the actual password to your AdGuard Home server.

### ExternalDNS Manifest

Apply the following manifest to deploy ExternalDNS, editing values for your environment accordingly.
Be sure to change the namespace in the `ClusterRoleBinding` if you are using a namespace other than **default**.

```yaml
---
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
  verbs: ["list","watch"]
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
        # If authentication is disabled and/or you didn't create
        # a secret, you can remove this block.
        envFrom:
        - secretRef:
            # Change this if you gave the secret a different name
            name: adguardhome-password
        args:
        - --source=service
        - --source=ingress
        # AdGuard Home only supports A/AAAA/CNAME records so there is no mechanism to track ownership.
        # You don't need to set this flag, but if you leave it unset, you will receive warning
        # logs when ExternalDNS attempts to create TXT records.
        - --registry=noop
        # IMPORTANT: If you have records that you manage manually in AdGuard Home, set
        # the policy to upsert-only so they do not get deleted.
        - --policy=upsert-only
        - --provider=adguardhome
        # Change this to the actual address of your AdGuard Home web server
        - --adguardhome-server=http://adguardhome.dns.svc.cluster.local
      securityContext:
        fsGroup: 65534 # For ExternalDNS to be able to read Kubernetes token files
```

### Arguments

- `--adguardhome-server (env: EXTERNAL_DNS_ADGUARDHOME_SERVER)` - The address of the AdGuard Home web server
- `--adguardhome-username (env: EXTERNAL_DNS_ADGUARDHOME_USERNAME)` - The username to the AdGuard Home web server
- `--adguardhome-password (env: EXTERNAL_DNS_ADGUARDHOME_PASSWORD)` - The password to the AdGuard Home web server

## Verify ExternalDNS Works

### Ingress Example

Create an Ingress resource. ExternalDNS will use the hostname specified in the Ingress object.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: foo
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: foo
            port:
              number: 80
```

### Service Example

The below sample application can be used to verify services work.
For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: my-application
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.homelab.com
spec:
  type: LoadBalancer
  ports:
  - port: 80
    name: http
    targetPort: 80
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-application
spec:
  selector:
    matchLabels:
      app: my-application
  template:
    metadata:
      labels:
        app: my-application
    spec:
      containers:
      - image: my-application
        name: my-application
        ports:
        - containerPort: 80
          name: http
```

You can then query your AdGuard Home to see if the record was created.

_Change `@192.168.100.2` to the actual address of your DNS server_

```bash
$ dig +short @192.168.100.2  nginx.external-dns-test.homelab.com

192.168.100.129
```
