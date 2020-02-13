# Running ExternalDNS with limited privileges

You can run ExternalDNS with reduced privileges since `v0.5.6` using the following `SecurityContext`.

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
        image: eu.gcr.io/k8s-artifacts-prod/external-dns/external-dns:v0.6.0 # minimum version is v0.5.6
        args:
        - ... # your arguments here
        securityContext:
          runAsNonRoot: true
          runAsUser: 65534
          readOnlyRootFilesystem: true
          capabilities:
            drop: ["ALL"]
```
