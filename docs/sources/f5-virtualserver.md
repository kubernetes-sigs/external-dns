# F5 Networks VirtualServer Source

This tutorial describes how to configure ExternalDNS to use the F5 Networks VirtualServer Source. It is meant to supplement the other provider-specific setup tutorials.

The F5 Networks VirtualServer CRD is part of [this](https://github.com/F5Networks/k8s-bigip-ctlr) project.
See more in-depth info regarding the VirtualServer CRD [in the official documentation](https://github.com/F5Networks/k8s-bigip-ctlr/blob/master/docs/config_examples/customResource/CustomResource.md#virtualserver).

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

## How it works

The F5 VirtualServer source creates DNS records based on the following fields:

- **`spec.host`**: The primary hostname for the virtual server
- **`spec.hostAliases`**: Additional hostnames that should also resolve to the same targets
- **`spec.virtualServerAddress`**: The IP address to use as the target (if no target annotation is set)
- **`status.vsAddress`**: The IP address from the status field (if no spec address or target annotation is set)

### Example VirtualServer with hostAliases

```yaml
apiVersion: cis.f5.com/v1
kind: VirtualServer
metadata:
  name: example-vs
  namespace: default
spec:
  host: www.example.com
  hostAliases:
    - alias1.example.com
    - alias2.example.com
  virtualServerAddress: 192.168.1.100
```

This configuration will create DNS A records for:

- `www.example.com` → `192.168.1.100`
- `alias1.example.com` → `192.168.1.100`
- `alias2.example.com` → `192.168.1.100`

### Target Priority

The source follows this priority order for determining targets:

1. **Target annotation**: `external-dns.alpha.kubernetes.io/target` (highest priority)
2. **Spec address**: `spec.virtualServerAddress`
3. **Status address**: `status.vsAddress`

If none of these are available, the VirtualServer will be skipped.

### TTL Support

You can set a custom TTL using the annotation:

```yaml
annotations:
  external-dns.alpha.kubernetes.io/ttl: "300"
```

### Annotation Filtering

You can filter VirtualServers using the `--annotation-filter` flag to only process those with specific annotations.
