# F5 Networks TransportServer Source

This tutorial describes how to configure ExternalDNS to use the F5 Networks TransportServer Source. It is meant to supplement the other provider-specific setup tutorials.

The F5 Networks TransportServer CRD is part of [this](https://github.com/F5Networks/k8s-bigip-ctlr) project. See more in-depth info regarding the TransportServer CRD [here](https://github.com/F5Networks/k8s-bigip-ctlr/tree/master/docs/cis-20.x/config_examples/customResource/TransportServer).

## Start with ExternalDNS with the F5 Networks TransportServer source

1. Make sure that you have the `k8s-bigip-ctlr` installed in your cluster. The needed CRDs are bundled within the controller.

2. In your Helm `values.yaml` add:
```
sources:
  - ...
  - f5-transportserver
  - ...
```
or add it in your `Deployment` if you aren't installing `external-dns` via Helm:
```
args:
- --source=f5-transportserver
```

Note that, in case you're not installing via Helm, you'll need the following in the `ClusterRole` bound to the service account of `external-dns`:
```
- apiGroups:
  - cis.f5.com
  resources:
  - transportservers
  verbs:
  - get
  - list
  - watch
```