# Setting up ExternalDNS for ExternalName Services

This tutorial describes how to setup ExternalDNS for usage in conjunction with an ExternalName service.

## Usecases

The main use cases that inspired this feature is the necessity for having a subdomain pointing to an external domain. In this scenario, it makes sense for the subdomain to have a CNAME record pointing to the external domain.

## Setup

### External DNS
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
        - --log-level=debug
        - --source=service
        - --source=ingress
        - --namespace=dev
        - --domain-filter=example.org.
        - --provider=aws
        - --registry=txt
        - --txt-owner-id=dev.example.org
```

### ExternalName Service

```yaml
kind: Service
apiVersion: v1
metadata:
  name: aws-service
  annotations:
    external-dns.alpha.kubernetes.io/hostname: tenant1.example.org,tenant2.example.org
spec:
  type: ExternalName
  externalName: aws.external.com
```

This will create 2 CNAME records pointing to `aws.example.org`:
```
tenant1.example.org
tenant2.example.org
```

### ExternalName Service with an IP address

If `externalName` is an IP address, External DNS will create A records instead of CNAME.

```yaml
kind: Service
apiVersion: v1
metadata:
  name: aws-service
  annotations:
    external-dns.alpha.kubernetes.io/hostname: tenant1.example.org,tenant2.example.org
spec:
  type: ExternalName
  externalName: 111.111.111.111
```

This will create 2 A records pointing to `111.111.111.111`:
```
tenant1.example.org
tenant2.example.org
```
