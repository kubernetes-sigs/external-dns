# F5 Networks VirtualServer Source

This tutorial describes how to configure ExternalDNS to use the F5 Networks VirtualServer Source. It is meant to supplement the other provider-specific setup tutorials.

The F5 Networks VirtualServer CRD is part of [this](https://github.com/F5Networks/k8s-bigip-ctlr) project.
See more in-depth info regarding the VirtualServer CRD [here](https://github.com/F5Networks/k8s-bigip-ctlr/blob/master/docs/config_examples/customResource/CustomResource.md#virtualserver).

## Start with ExternalDNS with the F5 Networks VirtualServer source

1. Make sure that you have the `k8s-bigip-ctlr` installed in your cluster. The needed CRDs are bundled within the controller.

2. In your Helm `values.yaml` add:

```yaml
sources:
  - ...
  - f5-virtualserver
  - ...
```

or add it in your `Deployment` if you aren't installing `external-dns` via Helm:

```yaml
args:
- --source=f5-virtualserver
```

Note that, in case you're not installing via Helm, you'll need the following in the `ClusterRole` bound to the service account of `external-dns`:

```yaml
- apiGroups:
  - cis.f5.com
  resources:
  - virtualservers
  verbs:
  - get
  - list
  - watch
```
