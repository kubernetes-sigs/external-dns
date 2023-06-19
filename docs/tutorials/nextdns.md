# Setting up ExternalDNS for NextDNS

This tutorial describes how to setup ExternalDNS to sync records with NextDNS rewrites.
NextDNS has a list of rewrites that allow you to set or override the DNS response for any domain or subdomain. This list can contain any number of arbitrary A or CNAME records.
There is an API that ExternalDNS is able to use to manage these records.

## Deploy ExternalDNS

The NextDNS provider requires the Profile ID and your API key provided by NextDNS.
You can find the Profile ID on the Setup tab in the `ID` field.
You can find your API key towards the bottom of the account page.

The API key should be a secret and can be created with this command

```bash
kubectl create secret generic nextdns-api-key \
    --from-literal EXTERNAL_DNS_NEXTDNS_API_KEY=api-key
```

Replacing **"api-key"** with the actual API key from NextDNS.

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
    namespace: external-dns
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
      imagePullSecrets:
        - name: registry-credentials
      containers:
        - name: external-dns
          image: registry.k8s.io/external-dns/external-dns:version-that-supports-this
          envFrom:
            - secretRef:
                name: nextdns-secrets
          args:
            - --source=ingress
            - --provider=nextdns
            - --nextdns-profile-id=profile-id
            - --registry=noop
```

### Arguments

 - `--nextdns-profile-id (env: EXTERNAL_DNS_NEXTDNS_PROFILE_ID)` - The NextDNS Profile ID to update
 - `--nextdns-api-key (env: EXTERNAL_DNS_NEXTDNS_API_KEY)` - Your API key from NextDNS

## Verify ExternalDNS Works

### Ingress Example

Create an Ingress resource. ExternalDNS will use the hostname specified in the Ingress object.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: foo
spec:
  ingressClassName: nginx
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

The below sample application can be used to verify Services work.
For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
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
          name: http
```

You can then look at NextDNS to see if the record was created.

_Change `@45.90.28.0` to the actual address of your NextDNS server_

```bash
$ dig +short @45.90.28.0  nginx.external-dns-test.homelab.com

192.168.100.129
```
