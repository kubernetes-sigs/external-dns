# Setting up ExternalDNS for Exoscale

## Prerequisites

Exoscale provider support was added via [this PR](https://github.com/kubernetes-sigs/external-dns/pull/625), thus you need to use external-dns v0.5.5.

The Exoscale provider expects that your Exoscale zones, you wish to add records to, already exists
and are configured correctly. It does not add, remove or configure new zones in anyway.

To do this please refer to the [Exoscale DNS documentation](https://community.exoscale.com/documentation/dns/).

Additionally you will have to provide the Exoscale...:

* API Key
* API Secret
* API Endpoint
* Elastic IP address, to access the workers

## Deployment

Deploying external DNS for Exoscale is actually nearly identical to deploying
it for other providers. This is what a sample `deployment.yaml` looks like:

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
      # Only use if you're also using RBAC
      # serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=ingress # or service or both
        - --provider=exoscale
        - --domain-filter={{ my-domain }}
        - --policy=sync # if you want DNS entries to get deleted as well
        - --txt-owner-id={{ owner-id-for-this-external-dns }}
        - --exoscale-endpoint={{ endpoint }} # usually https://api.exoscale.ch/dns
        - --exoscale-apikey={{ api-key}}
        - --exoscale-apisecret={{ api-secret }}
```

## RBAC

If your cluster is RBAC enabled, you also need to setup the following, before you can run external-dns:

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: default

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
```

## Testing and Verification

**Important!**: Remember to change `example.com` with your own domain throughout the following text.

Spin up a simple nginx HTTP server with the following spec (`kubectl apply -f`):

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx
  annotations:
    kubernetes.io/ingress.class: nginx
    external-dns.alpha.kubernetes.io/target: {{ Elastic-IP-address }}
spec:
  rules:
  - host: via-ingress.example.com
    http:
      paths:
      - backend:
          service:
            name: "nginx"
            port:
              number: 80
        pathType: Prefix

---

apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  ports:
  - port: 80
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
```

**Important!**: Don't run dig, nslookup or similar immediately (until you've
confirmed the record exists). You'll get hit by [negative DNS caching](https://tools.ietf.org/html/rfc2308), which is hard to flush.

Wait about 30s-1m (interval for external-dns to kick in), then check Exoscales [portal](https://portal.exoscale.com/dns/example.com)... via-ingress.example.com should appear as a A and TXT record with your Elastic-IP-address.
