# External-dns provider

This provider uses the external-dns CRD in an external kubernetes cluster as the provider backend. The idea is that the management cluster is the one that has access to the real provider backend. It may use tools like OpenPolicyAgent Gatekeeper or other admission controllers to centrally govern policies of which DNS entries are allowed and which are not. A workload cluster runs external-dns and only needs a ServiceAccount on the management cluster, but doesn't need credentials to the "real" provider backend.

## Development using Kind

As we have a workload and a management cluster, we create 2 KIND clusters, and install the CRD in both.
In the management cluster we install a external-dns with inmemory provider (strictly this isn't required)
and also set up a ServiceAccount `workload1-external-dns` in namespace `workload1` which we give CRUD
permissions to the DNSEndpoint of name `workload1`.

```
kind create cluster --name kind 
kind create cluster --name mgmt 
kubectl config use-context kind-kind

kubectl --context kind-kind apply -f docs/contributing/crd-source/crd-manifest.yaml
kubectl --context kind-mgmt apply -f docs/contributing/crd-source/crd-manifest.yaml
kubectl --context kind-mgmt apply -f provider/externaldns/mgmt.yaml
```

Next we extract the various values we need to run external-dns:
```
secretName=$(kubectl --context kind-mgmt -n workload1 get sa workload1-external-dns '-ojsonpath={.secrets[0].name}')
kubectl --context kind-mgmt -n workload1 get secret $secretName '-ojsonpath={.data.ca\.crt}' | base64 -D > .ca.crt
kubectl --context kind-mgmt -n workload1 get secret $secretName '-ojsonpath={.data.namespace}' | base64 -D > .namespace
kubectl --context kind-mgmt -n workload1 get secret $secretName '-ojsonpath={.data.token}' | base64 -D > .token

server=$(kubectl --context kind-mgmt config view --minify | yq e '.clusters[0].cluster.server' - | sed 's;https://;;')
k8shost=$(echo $server | awk -F":" '{print $1}')
k8sport=$(echo $server | awk -F":" '{print $2}')
```

Now in one console we run external-dns against the workload cluster, pointing to the mgmt cluster as the target:
```
kubectl config use-context kind-kind
make && build/external-dns --source=crd --provider=externaldns --externaldns-ca-cert-path=.ca.crt --externaldns-namespace-path=.namespace --externaldns-token-path=.token --externaldns-kubernetes-host=$k8shost --externaldns-kubernetes-port=$k8sport --externaldns-resource-name=workload1
```

And in a second console we change the state of workload cluster, and then check that it has been updated in the mgmt cluster:
```
kubectl --context kind-kind apply -f provider/externaldns/workload1.yaml
sleep 60
kubectl --context kind-mgmt -n workload1 get dnsendpoint foo -oyaml

kubectl --context kind-kind apply -f provider/externaldns/workload1b.yaml
sleep 60
kubectl --context kind-mgmt -n workload1 get dnsendpoint foo -oyaml

kubectl --context kind-kind delete -f provider/externaldns/workload1.yaml
sleep 60
kubectl --context kind-mgmt -n workload1 get dnsendpoint foo -oyaml

kubectl --context kind-kind delete -f provider/externaldns/workload1b.yaml
sleep 60
kubectl --context kind-mgmt -n workload1 get dnsendpoint foo -oyaml
```

## Running in Workload Cluster

### One time Management config

To allow for this to work, we need to ensure the SV has the following additional configuration (only admins can do this):
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cl-system-wc-manager-externaldns-role
rules:
- apiGroups: ["externaldns.k8s.io"]
  resources: ["dnsendpoints"]
  verbs: ["create", "update", "delete", "get","watch","list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cl-system-wc-manager-externaldns-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cl-system-wc-manager-externaldns-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: org-system-wc
```

We could also create this per TKC (using `Role` and `RoleBinding`), and then lock down the `resourceNames` in the `Role`.
But it is likely not necessary, as this `ClusterRole` is assigned to the TKG manager, not to the TKC. The per-TKC permissions
are handled below.

### Management config for a Workload cluster

Inside the same namespace as the TKC, create this CR, which 
```
apiVersion: run.tanzu.vmware.com/v1alpha1
kind: ProviderServiceAccount
metadata:
  name: gc-cl-externaldns
  namespace: cl-blog
spec:
  ref:
    name: gc
  rules:
  - apiGroups: ["externaldns.k8s.io"]
    resources: ["dnsendpoints"]
    resourceNames: ["gc-external-dns"]
    verbs: ["create", "update", "delete", "get","watch","list"]
  targetNamespace: external-dns
  targetSecretName: svcreds
```

### Inside the Workload Cluster

See `workload-external-dns.yaml` for a more complete example of how to deploy external-dns, but the key is the `PodSpec`:
```
    spec:
      serviceAccountName: external-dns
      volumes:
      - name: svcreds
        secret:
          secretName: svcreds
      containers:
      - name: external-dns
        # XXX: Adjust this to the right image
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=crd
        - --source=service
        - --provider=externaldns
        - --externaldns-ca-cert-path=/svcreds/ca.crt
        - --externaldns-namespace-path=/svcreds/namespace
        - --externaldns-token-path=/svcreds/token
        - --externaldns-kubernetes-host="supervisor.default.svc"
        - --externaldns-kubernetes-port="6443"
        # XXX: Adjust this to match the name set up in ProviderServiceAccount in SV
        - --externaldns-resource-name=gc-external-dns
        volumeMounts:
        - name: svcreds
          mountPath: "/svcreds"
          readOnly: true
```

