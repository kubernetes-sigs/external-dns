# Setting up Akamai FastDNS

## Prerequisites

Akamai FastDNS provider support was added via [this PR](https://github.com/kubernetes-sigs/external-dns/pull/1384), thus you need to use a release where this pr is included. This should be at least v0.5.18

The Akamai FastDNS provider expects that your zones, you wish to add records to, already exists
and are configured correctly. It does not add, remove or configure new zones in anyway.

To do this pease refer to the [FastDNS documentation](https://learn.akamai.com/en-us/products/web_performance/fast_dns.html).

Additional data you will have to provide:

* Service Consumer Domain
* Access token
* Client token
* Client Secret

Make these available to external DNS somehow. In the following example a secret is used by referencing the secret and its keys in the env section of the deployment.

If you happen to have questions regarding authentification, please refer to the [API Client Authentication documentation](https://developer.akamai.com/legacy/introduction/Client_Auth.html)

## Deployment

Deploying external DNS for Akamai is actually nearly identical to deploying
it for other providers. This is what a sample `deployment.yaml` looks like:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  labels:
    app.kubernetes.io/name: external-dns
    app.kubernetes.io/version: v0.6.0
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app.kubernetes.io/name: external-dns
  template:
    metadata:
      labels:
        app.kubernetes.io/name: external-dns
        app.kubernetes.io/version: v0.6.0
    spec:
      # Only use if you're also using RBAC
      # serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: eu.gcr.io/k8s-artifacts-prod/external-dns/external-dns:v0.6.0
        args:
        - --source=ingress # or service or both
        - --provider=akamai
        - --registry=txt
        - --txt-owner-id={{ owner-id-for-this-external-dns }}
        env:
        - name: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
        - name: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
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
```
## Verify ExternalDNS works (Ingress example)

Create an ingress resource manifest file.

> For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object.

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: foo
  annotations:
    kubernetes.io/ingress.class: "nginx" # use the one that corresponds to your ingress controller.
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - backend:
          serviceName: foo
          servicePort: 80
```

## Verify ExternalDNS works (Service example)

Create the following sample application to test that ExternalDNS works.

> For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

> If you want to give multiple names to service, you can set it to external-dns.alpha.kubernetes.io/hostname with a comma separator.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com
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


**Important!**: Don't run dig, nslookup or similar immediately. You'll get hit by [negative DNS caching](https://tools.ietf.org/html/rfc2308), which is hard to flush.
Wait about 30s-1m (interval for external-dns to kick in)
