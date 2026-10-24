# Watching Multiple Namespaces

By default, ExternalDNS watches every namespace. `--namespace` limits the watched
namespaces. Repeat it or pass a comma-separated list:

```sh
--namespace=team-a --namespace=team-b
--namespace=team-a,team-b
```

Sources marked `multiple` in the Namespace column of [Supported Sources](../sources/index.md)
accept several values. With any other source, ExternalDNS refuses to start when given
more than one namespace.

## RBAC

ExternalDNS opens one watch per namespace, so a `ClusterRole` is not required for
namespaced resources. Bind the role in each watched namespace with a `RoleBinding`:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: external-dns
  namespace: team-a # repeat for team-b
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
  - kind: ServiceAccount
    name: external-dns
    namespace: external-dns
```

Cluster-scoped resources still need a `ClusterRoleBinding`, for example nodes for the
`service` source with `NodePort` services.

## Cost

Each source opens one watch per namespace for each resource type it reads. Watching
many namespaces this way costs more API server connections and memory than a single
cluster-wide watch.

When ExternalDNS should read the whole cluster, omit `--namespace` and
narrow the selection with `--label-filter` or `--annotation-filter` instead.
