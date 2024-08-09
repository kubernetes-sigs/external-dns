# Configuring ExternalDNS to use the Traefik Proxy Source

This tutorial describes how to configure ExternalDNS to use the Traefik Proxy source.
It is meant to supplement the other provider-specific setup tutorials.

# Using Manifest 

## Without RBAC enabled

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
        # update this to the desired external-dns version
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=traefik-proxy
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
        # use this flag when using Traefik helm chart >= v28
        - --traefik-disable-legacy 
```

## With RBAC enabled

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
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list","watch"]
- apiGroups: ["traefik.io"]
  resources: ["ingressroutes", "ingressroutetcps", "ingressrouteudps"]
  verbs: ["get","watch","list"]
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
        # update this to the desired external-dns version
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=traefik-proxy
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=my-identifier
        # use this flag when using Traefik helm chart >= v28
        - --traefik-disable-legacy 
```

# Using Helm chart

## With RBAC enabled

```yaml
rbac:
  create: true
sources:
  - traefik-proxy-v3
extraArgs:
  - traefik-disable-legacy 
```

# Deploying a Traefik IngressRoute

Create an IngressRoute file called 'traefik-ingress.yaml' with the following contents:

```yaml
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: traefik-ingress
  annotations:
    external-dns.alpha.kubernetes.io/target: traefik.example.com
    kubernetes.io/ingress.class: traefik
spec:
  entryPoints:
    - web
    - websecure
  routes:
    - match: Host(`application.example.com`)
      kind: Rule
      services:
        - name: service
          namespace: namespace
          port: port
```

Note the annotation on the IngressRoute (`external-dns.alpha.kubernetes.io/target`); use the same hostname as the Traefik DNS.

ExternalDNS uses this annotation to determine what services should be registered with DNS.

Create the IngressRoute:

```
$ kubectl create -f traefik-ingress.yaml
```

Depending on where you run your IngressRoute it can take a little while for ExternalDNS to synchronize the DNS record.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Traefik DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f traefik-ingress.yaml
$ kubectl delete -f externaldns.yaml
```

## Additional Flags

| Flag | Description |
| --- | --- |
| --traefik-disable-legacy | Disable listeners on Resources under traefik.containo.us |
| --traefik-disable-new | Disable listeners on Resources under traefik.io |

Example:

1. Manifest
```yaml
apiVersion: apps/v1
kind: Deployment
spec:
  ...
  template:
    ...
    spec:
      ...
      args:
        ...
        - --traefik-disable-legacy
```

2. Helm chart
```yaml
sources:
  - traefik-proxy-v3
extraArgs:
  - traefik-disable-legacy 
```

### Disabling Resource Listeners

Traefik has deprecated the legacy API group, traefik.containo.us, in favor of traefik.io. By default the traefik-proxy source will listen for resources under both API groups; however, this may cause timeouts with the following message

```
FATA[0060] failed to sync traefik.io/v1alpha1, Resource=ingressroutes: context deadline exceeded
```

In this case you can disable one or the other API groups with `--traefik-disable-new` or `--traefik-disable-legacy`

