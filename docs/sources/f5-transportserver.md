# F5 Networks TransportServer Source

This tutorial describes how to configure ExternalDNS to use the F5 Networks TransportServer Source. It is meant to supplement the other provider-specific setup tutorials.

The F5 Networks TransportServer CRD is part of [this](https://github.com/F5Networks/k8s-bigip-ctlr) project. See more in-depth info regarding the TransportServer CRD [here](https://github.com/F5Networks/k8s-bigip-ctlr/tree/master/docs/cis-20.x/config_examples/customResource/TransportServer).

## Start with ExternalDNS with the F5 Networks TransportServer source

1. Make sure that you have the `k8s-bigip-ctlr` installed in your cluster. The needed CRDs are bundled within the controller.

2. In your Helm `values.yaml` add:

```yaml
sources:
  - ...
  - f5-transportserver
  - ...
```

or add it in your `Deployment` if you aren't installing `external-dns` via Helm:

```yaml
args:
- --source=f5-transportserver
```

Note that, in case you're not installing via Helm, you'll need the following in the `ClusterRole` bound to the service account of `external-dns`:

```yaml
- apiGroups:
  - cis.f5.com
  resources:
  - transportservers
  verbs:
  - get
  - list
  - watch
```

### Example TransportServer CR w/ host in spec

```yaml
apiVersion: cis.f5.com/v1
kind: TransportServer
metadata:
  labels:
    f5cr: 'true'
  name: test-ts
  namespace: test-ns
spec:
  bigipRouteDomain: 0
  host: test.example.com
  ipamLabel: vips
  mode: standard
  pool:
    service: test-service
    servicePort: 4222
  virtualServerPort: 4222
```

### Example TransportServer CR w/ target annotation set

If the `external-dns.alpha.kubernetes.io/target` annotation is set, the record created will reflect that and everything else will be ignored.

```yaml
apiVersion: cis.f5.com/v1
kind: TransportServer
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target: 10.172.1.12
  labels:
    f5cr: 'true'
  name: test-ts
  namespace: test-ns
spec:
  bigipRouteDomain: 0
  host: test.example.com
  ipamLabel: vips
  mode: standard
  pool:
    service: test-service
    servicePort: 4222
  virtualServerPort: 4222
```

### Example TransportServer CR w/ VirtualServerAddress set

If `virtualServerAddress` is set, the record created will reflect that. `external-dns.alpha.kubernetes.io/target` will take precedence though.

```yaml
apiVersion: cis.f5.com/v1
kind: TransportServer
metadata:
  labels:
    f5cr: 'true'
  name: test-ts
  namespace: test-ns
spec:
  bigipRouteDomain: 0
  host: test.example.com
  ipamLabel: vips
  mode: standard
  pool:
    service: test-service
    servicePort: 4222
  virtualServerPort: 4222
  virtualServerAddress: 10.172.1.123
```

If there is no target annotation or `virtualServerAddress` field set, then it'll use the `VSAddress` field from the created TransportServer status to create the record.
