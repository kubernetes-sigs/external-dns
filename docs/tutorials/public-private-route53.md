# Setting up ExternalDNS using the same domain for public and private Route53 zones

This tutorial describes how to setup ExternalDNS using the same domain for public and private Route53 zones and [nginx-ingress-controller](https://github.com/kubernetes/ingress-nginx). It also outlines how to use [cert-manager](https://github.com/jetstack/cert-manager) to automatically issue SSL certificates from [Let's Encrypt](https://letsencrypt.org/) for both public and private records.

## Deploy public nginx-ingress-controller

Consult [External DNS nginx ingress docs](nginx-ingress.md) for installation guidelines.

Specify `ingress-class` in nginx-ingress-controller container args:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: external-ingress
  name: external-ingress-controller
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-ingress
  template:
    metadata:
      labels:
        app: external-ingress
    spec:
      containers:
      - args:
        - /nginx-ingress-controller
        - --default-backend-service=$(POD_NAMESPACE)/default-http-backend
        - --configmap=$(POD_NAMESPACE)/external-ingress-configuration
        - --tcp-services-configmap=$(POD_NAMESPACE)/external-tcp-services
        - --udp-services-configmap=$(POD_NAMESPACE)/external-udp-services
        - --annotations-prefix=nginx.ingress.kubernetes.io
        - --ingress-class=external-ingress
        - --publish-service=$(POD_NAMESPACE)/external-ingress
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        image: quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.11.0
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: 10254
            scheme: HTTP
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: external-ingress-controller
        ports:
        - containerPort: 80
          name: http
          protocol: TCP
        - containerPort: 443
          name: https
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: 10254
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
```

Set `type: LoadBalancer` in your public nginx-ingress-controller Service definition.

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout: "3600"
    service.beta.kubernetes.io/aws-load-balancer-proxy-protocol: '*'
  labels:
    app: external-ingress
  name: external-ingress
spec:
  externalTrafficPolicy: Cluster
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: http
  - name: https
    port: 443
    protocol: TCP
    targetPort: https
  selector:
    app: external-ingress
  sessionAffinity: None
  type: LoadBalancer
```

## Deploy private nginx-ingress-controller

Consult [External DNS nginx ingress docs](nginx-ingress.md) for installation guidelines.

Make sure to specify `ingress-class` in nginx-ingress-controller container args:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: internal-ingress
  name: internal-ingress-controller
spec:
  replicas: 1
  selector:
    matchLabels:
      app: internal-ingress
  template:
    metadata:
      labels:
        app: internal-ingress
    spec:
      containers:
      - args:
        - /nginx-ingress-controller
        - --default-backend-service=$(POD_NAMESPACE)/default-http-backend
        - --configmap=$(POD_NAMESPACE)/internal-ingress-configuration
        - --tcp-services-configmap=$(POD_NAMESPACE)/internal-tcp-services
        - --udp-services-configmap=$(POD_NAMESPACE)/internal-udp-services
        - --annotations-prefix=nginx.ingress.kubernetes.io
        - --ingress-class=internal-ingress
        - --publish-service=$(POD_NAMESPACE)/internal-ingress
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        image: quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.11.0
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: 10254
            scheme: HTTP
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: internal-ingress-controller
        ports:
        - containerPort: 80
          name: http
          protocol: TCP
        - containerPort: 443
          name: https
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: 10254
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
```

Set additional annotations in your private nginx-ingress-controller Service definition to create an internal load balancer.

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout: "3600"
    service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
    service.beta.kubernetes.io/aws-load-balancer-proxy-protocol: '*'
  labels:
    app: internal-ingress
  name: internal-ingress
spec:
  externalTrafficPolicy: Cluster
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: http
  - name: https
    port: 443
    protocol: TCP
    targetPort: https
  selector:
    app: internal-ingress
  sessionAffinity: None
  type: LoadBalancer
```

## Deploy the public zone ExternalDNS

Consult [AWS ExternalDNS setup docs](aws.md) for installation guidelines.

In ExternalDNS containers args, make sure to specify `annotation-filter` and `aws-zone-type`:

```yaml
apiVersion: apps/v1beta2
kind: Deployment
metadata:
  labels:
    app: external-dns-public
  name: external-dns-public
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns-public
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns-public
    spec:
      containers:
      - args:
        - --source=ingress
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=external-dns
        - --annotation-filter=kubernetes.io/ingress.class in (external-ingress)
        - --aws-zone-type=public
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        name: external-dns-public
```

## Deploy the private zone ExternalDNS

Consult [AWS ExternalDNS setup docs](aws.md) for installation guidelines.

In ExternalDNS containers args, make sure to specify `annotation-filter` and `aws-zone-type`:

```yaml
apiVersion: apps/v1beta2
kind: Deployment
metadata:
  labels:
    app: external-dns-private
  name: external-dns-private
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns-private
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns-private
    spec:
      containers:
      - args:
        - --source=ingress
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=dev.k8s.nexus
        - --annotation-filter=kubernetes.io/ingress.class in (internal-ingress)
        - --aws-zone-type=private
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        name: external-dns-private
```

## Create application Service definitions

For this setup to work, you've to create two Service definitions for your application.

At first, create public Service definition:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: "external-ingress"
  labels:
    app: app
  name: app-public
spec:
  rules:
  - host: app.domain.com
    http:
      paths:
      - backend:
          service:
            name: app
            port:
              number: 80
        pathType: Prefix
```

Then create private Service definition:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: "internal-ingress"
  labels:
    app: app
  name: app-private
spec:
  rules:
  - host: app.domain.com
    http:
      paths:
      - backend:
          service:
            name: app
            port:
              number: 80
        pathType: Prefix
```

Additionally, you may leverage [cert-manager](https://github.com/jetstack/cert-manager) to automatically issue SSL certificates from [Let's Encrypt](https://letsencrypt.org/). To do that, request a certificate in public service definition:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    certmanager.k8s.io/acme-challenge-type: "dns01"
    certmanager.k8s.io/acme-dns01-provider: "route53"
    certmanager.k8s.io/cluster-issuer: "letsencrypt-production"
    kubernetes.io/ingress.class: "external-ingress"
    kubernetes.io/tls-acme: "true"
  labels:
    app: app
  name: app-public
spec:
  rules:
  - host: app.domain.com
    http:
      paths:
      - backend:
          service:
            name: app
            port:
              number: 80
        pathType: Prefix
  tls:
  - hosts:
    - app.domain.com
    secretName: app-tls
```

And reuse the requested certificate in private Service definition:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: "internal-ingress"
  labels:
    app: app
  name: app-private
spec:
  rules:
  - host: app.domain.com
    http:
      paths:
      - backend:
          service:
            name: app
            port:
              number: 80
        pathType: Prefix
  tls:
  - hosts:
    - app.domain.com
    secretName: app-tls
```
