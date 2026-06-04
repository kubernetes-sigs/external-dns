# The CRD registry

!!! warning "Alpha"
    The CRD registry is at an early, alpha stage. The `DNSRecord` API
    (`externaldns.k8s.io/v1alpha1`) and its behaviour may change in a backward
    incompatible way. It is not yet recommended for production use.

The CRD registry stores DNS record ownership and metadata as `DNSRecord` custom
resources in the Kubernetes cluster, instead of in TXT records (TXT registry) or
an external table (DynamoDB registry).

Each managed endpoint is persisted as a `DNSRecord` object, so the records
ExternalDNS owns can be inspected with plain `kubectl`:

```bash
kubectl get dnsrecords
```

```text
NAME                         DNS NAME             TYPE    SET ID   TARGETS    STATUS
sub-example-com-a-1a2b3c4d   sub.example.com      A                1.2.3.4    Programmed
```

> The CRD registry is a **trustworthy record of what ExternalDNS applied** — it
> is not a mirror of the DNS provider. A `DNSRecord` is written and marked
> `Accepted` before the provider is called, then `Programmed` once the provider
> accepts the change (or `Failed` if it does not). The `STATUS` column shows this
> stage. Only `Programmed` records are treated as current state, so a record left
> un-programmed by a provider failure is re-applied on the next reconcile rather
> than mistaken for one that already exists. Records changed out-of-band directly
> in the provider are not reconciled by this registry.

## Limitations

* Only the **in-cluster** Kubernetes API is currently supported (the cluster
  ExternalDNS runs in). Using another kubeconfig is planned for a follow-up.
* The Helm chart does not yet wire the required RBAC; it must be added manually
  (see below).

## Install the DNSRecord CRD

Apply the `DNSRecord` CustomResourceDefinition before enabling the registry:

```bash
kubectl apply -f config/crd/standard/dnsrecords.externaldns.k8s.io.yaml
```

(The same manifest is published with the Helm chart under
`charts/external-dns/crds/`.)

## RBAC

ExternalDNS needs to read and write `DNSRecord` objects (including their status)
in the namespace where they are stored:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: external-dns-crd-registry
  namespace: external-dns
rules:
  - apiGroups: ["externaldns.k8s.io"]
    resources: ["dnsrecords"]
    verbs: ["get", "list", "watch", "create", "update", "delete"]
  - apiGroups: ["externaldns.k8s.io"]
    resources: ["dnsrecords/status"]
    verbs: ["get", "update"]
```

Bind it to the ExternalDNS `ServiceAccount`:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: external-dns-crd-registry
  namespace: external-dns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: external-dns-crd-registry
subjects:
  - kind: ServiceAccount
    name: external-dns
    namespace: external-dns
```

## Configuration

Enable the registry with the `--registry` flag and identify this ExternalDNS
instance with `--txt-owner-id`:

* `--registry=crd` — use the CRD registry.
* `--txt-owner-id=my-identifier` — a value unique to this ExternalDNS deployment,
  stable for its lifetime. Deployments sharing a DNS zone must use different
  owner IDs. See [Registries](registry.md).
* `--namespace=external-dns` — the namespace `DNSRecord` objects are created in.
  When unset, the registry uses the `default` namespace.

## Status

Each `DNSRecord` carries a status:

* `status.conditions[type=Ready]` reports whether the endpoint is live in the DNS
  provider. Its `reason` captures the lifecycle stage and is surfaced as the
  `STATUS` print column:
  * `Accepted` (`Ready=False`) — ExternalDNS has taken the endpoint into its
      plan but has not programmed it yet.
  * `Programmed` (`Ready=True`) — the endpoint has been applied to the provider.
  * `Failed` (`Ready=False`) — the provider rejected the change. Because the
      provider reports a single batch error that cannot be attributed to
      individual records, every record in a failed batch is marked `Failed`;
      records that were in fact applied are corrected to `Programmed` on the next
      reconcile.
* `status.observedGeneration` records the `.metadata.generation` that was last
  reconciled.

Inspect it with:

```bash
kubectl get dnsrecord sub-example-com-a-1a2b3c4d -o yaml
```

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSRecord
metadata: [...]
spec: [...]
status:
  observedGeneration: 1
  conditions:
    - type: Ready
      status: "True"
      reason: Programmed
      message: Endpoint applied to the DNS provider
      observedGeneration: 1
      lastTransitionTime: "2026-06-04T10:00:00Z"
```
