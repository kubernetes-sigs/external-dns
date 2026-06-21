---
tags: ["tutorial", "libdns", "transip"]
---

# libdns provider adapter

ExternalDNS drives [libdns](https://github.com/libdns) modules through one generic provider
(`--provider=libdns`); the active module is picked at runtime with `--libdns-provider`. It is the
recommended path for the simpler "flat zone" providers moving out of tree
([#4347](https://github.com/kubernetes-sigs/external-dns/issues/4347)) — see
[RFC 005](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/proposal/005-libdns-provider-adapter.md).

This tutorial uses the **TransIP** module (`github.com/libdns/transip`), currently the only one in
the curated set.

## Build with the `libdns` tag

libdns modules pull in vendor SDKs, so they are **not** in the default
`registry.k8s.io/external-dns/external-dns` image. Build with the `libdns` tag:

```bash
# binary
go build -tags libdns -o build/external-dns .

# image (the repo uses `ko`, which honors GOFLAGS)
GOFLAGS=-tags=libdns KO_DOCKER_REPO=my-registry/external-dns ko build --bare .
```

Without the tag, `--provider=libdns --libdns-provider=transip` fails fast with an
`unknown libdns provider "transip"` error.

## Configure

Enable the API and generate a private key in the
[TransIP control panel](https://www.transip.nl/cp/account/api/). The module takes a single JSON blob
via `--libdns-config` (or `EXTERNAL_DNS_LIBDNS_CONFIG`); the relevant
[fields](https://github.com/libdns/transip/blob/master/provider.go):

| Field | Meaning |
|---|---|
| `login` | TransIP username (required) |
| `private_key` | PEM key, **or a path to a file** holding it (required) |
| `token_storage` | `memory`, a file path, or empty (temp dir). Use `memory` on a read-only FS |

Use `--dry-run` to preview changes without writing: the adapter logs what it would set or delete and
applies nothing.

## Deploy ExternalDNS

The config is secret, so keep it in a `Secret`. Mounting the key as a file lets `private_key` point
at a path, avoiding a multi-line PEM inside JSON.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: transip-credentials
type: Opaque
stringData:
  libdns-config: |
    {"login":"YOUR_TRANSIP_LOGIN","private_key":"/etc/transip/private.key","token_storage":"memory"}
  private.key: |
    -----BEGIN PRIVATE KEY-----
    YOUR_TRANSIP_API_PRIVATE_KEY
    -----END PRIVATE KEY-----
---
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
- apiGroups: ["discovery.k8s.io"]
  resources: ["endpointslices"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
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
        image: my-registry/external-dns:libdns # built with -tags libdns
        args:
        - --source=service
        - --domain-filter=example.com # the TransIP zone you own
        - --provider=libdns
        - --libdns-provider=transip
        - --txt-owner-id=external-dns
        env:
        - name: EXTERNAL_DNS_LIBDNS_CONFIG
          valueFrom:
            secretKeyRef:
              name: transip-credentials
              key: libdns-config
        volumeMounts:
        - name: transip-key
          mountPath: /etc/transip
          readOnly: true
      volumes:
      - name: transip-key
        secret:
          secretName: transip-credentials
          items:
          - key: private.key
            path: private.key
```

`--domain-filter` scopes the managed zones. TransIP also implements zone discovery, so it is
optional here, but setting it is recommended.

## Verify

Annotate a Service and check the record after the next reconcile:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.kubernetes.io/hostname: test.example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
```

```bash
dig +short test.example.com @ns0.transip.net
```

## Limitations

The adapter targets flat-zone providers and does **not** support provider-native routing.
`SetIdentifier` and weighted/latency/geo fields are stripped in `AdjustEndpoints` (with a warning) to
keep the reconcile loop convergent. For routing policies, use a provider that models them.
