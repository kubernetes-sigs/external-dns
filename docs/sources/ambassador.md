# Ambassador / Emissary-ingress Host

This tutorial describes how to configure ExternalDNS to use the `ambassador-host` source.
It reads `Host.getambassador.io` resources
([Emissary-ingress](https://www.getambassador.io/docs/emissary), formerly Ambassador)
and creates DNS records for the hostnames they declare.

ExternalDNS uses the `getambassador.io/v3alpha1` CRD version. This requires
Emissary-ingress 3.x (the `datawire/ambassador` v2 CRDs are no longer installed by the
3.10 quickstart). For older deployments still serving only the v2 CRD, stay on
ExternalDNS v0.21.0 or earlier.

## How it works

For each `Host`, the source looks at the `external-dns.ambassador-service` annotation. The
value points to the Emissary-ingress `LoadBalancer` Service whose address is used as the
record target:

- `name` &ndash; service in the same namespace as the `Host`.
- `namespace/name` &ndash; service in an explicit namespace.
- `name.namespace` &ndash; Ambassador's historical cross-namespace syntax.

The hostname comes from `spec.hostname`. A `Host` without the
`external-dns.ambassador-service` annotation is ignored. The target can be overridden with
the standard `external-dns.kubernetes.io/target` annotation; TTL and provider-specific
annotations are honored as well.

## Run it locally with kind

The steps below run end to end on a local [kind](https://kind.sigs.k8s.io/) cluster using
the `inmemory` provider, so no cloud credentials are needed &ndash; the DNS changes
ExternalDNS would apply are printed to its log.

### 1. Create a cluster

```bash
kind create cluster --name external-dns-ambassador
```

### 2. Install Emissary-ingress 3.10

```bash
kubectl apply -f https://app.getambassador.io/yaml/emissary/3.10.0/emissary-crds.yaml
kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system

kubectl create namespace emissary
kubectl apply -f https://app.getambassador.io/yaml/emissary/3.10.0/emissary-emissaryns.yaml
kubectl -n emissary rollout status deployment/emissary-ingress
```

### 3. Deploy ExternalDNS

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["getambassador.io"]
  resources: ["hosts"]
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
  namespace: default
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
        image: registry.k8s.io/external-dns/external-dns:v0.21.0
        args:
        - --source=ambassador-host
        - --policy=upsert-only # prevents ExternalDNS from deleting any records, set --policy=sync to enable full synchronization (including deletions)
        - --provider=inmemory
        - --inmemory-zone=example.com # persist records so updates/deletes are observable
        - --log-level=debug # show the records that would be created
        # for a real provider, replace the two lines above, e.g.:
        # - --provider=xxx
        # - --domain-filter=example.com
        # - --registry=txt
        # - --txt-owner-id=my-identifier
```

Apply it:

```bash
kubectl apply -f external-dns.yaml
```

Or run it on the host from sources (handy for testing local changes to the source). This
uses your current kubeconfig context (the kind cluster), so no in-cluster RBAC is needed:

```bash
go run main.go \
    --source=ambassador-host \
    --provider=inmemory \
    --inmemory-zone=example.com \
    --interval=10s \
    --log-level=debug
```

`--inmemory-zone=example.com` gives the `inmemory` provider a zone to store records in.
Without it the provider keeps no state between reconcile loops, so it re-emits the same
`CREATE` every cycle and you never see an `UPDATE` or `DELETE`. `--interval=10s` shortens
the wait between reconcile loops (default is one minute).

### 4. Create a Host

The `inmemory` provider has no cloud LoadBalancer, so set the target explicitly with the
`external-dns.kubernetes.io/target` annotation. The `external-dns.ambassador-service`
annotation is still required for the `Host` to be processed.

```bash
kubectl apply -f - <<EOF
apiVersion: getambassador.io/v3alpha1
kind: Host
metadata:
  name: my-host
  namespace: default
  annotations:
    external-dns.ambassador-service: emissary/emissary-ingress
    external-dns.kubernetes.io/target: 203.0.113.10
spec:
  hostname: my-host.example.com
  acmeProvider:
    authority: none
EOF
```

### 5. Verify

```bash
kubectl logs -l app=external-dns -f
```

You should see ExternalDNS pick up the `Host` and create an A record for
`my-host.example.com` pointing at `203.0.113.10`:

```text
... level=debug msg="Endpoints generated from Host: default/my-host: [my-host.example.com 0 IN A  203.0.113.10 []]"
... level=info msg="CREATE: my-host.example.com 0 IN A  203.0.113.10 []"
```

With a real provider and a `LoadBalancer` Service, drop the `target` annotation and point
`external-dns.ambassador-service` at the Emissary-ingress Service
(`emissary/emissary-ingress` in the manifests above); ExternalDNS resolves the Service's
external address as the target.

### Cleanup

```bash
kind delete cluster --name external-dns-ambassador
```
