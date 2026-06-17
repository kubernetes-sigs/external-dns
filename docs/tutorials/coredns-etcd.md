---
tags: ["tutorial", "coredns", "etcd", "kind"]
---

# CoreDNS with etcd backend

## Overview

This tutorial describes how to deploy CoreDNS backed by etcd as a dynamic DNS provider for external-dns.
It shows how to configure external-dns to write DNS records into etcd, which CoreDNS will then serve.

### TL;DR

After completing this lab, you will have a Kubernetes environment running as containers in your local development machine with etcd, coredns and external-dns.

### Notes

- `CoreDNS` and etcd here run inside the cluster for demonstration purposes.
- For real deployments, you can use external etcd or secure etcd with TLS.
- The zone example.org is arbitrary — use your domain.
- `external-dns` automatically maintains records in etcd under `/skydns/<reversed-domain>`.

## Prerequisite

Before you start, ensure you have:

- A running kubernetes cluster.
  - In this tutorial we are going to use [kind](https://kind.sigs.k8s.io/)
- [`kubectl`](https://kubernetes.io/docs/tasks/tools/) and [`helm`](https://helm.sh/)
- `external-dns` source code or [helm chart](https://github.com/kubernetes-sigs/external-dns/tree/master/charts/external-dns)
- `CoreDNS` [helm chart](https://github.com/coredns/helm)
- Optional
  - `dnstools` container for testing
  - `etcdctl` to interact with [etcd](https://etcd.io/docs/v3.4/dev-guide/interacting_v3/)

## Bootstrap Environment

### 1. Create cluster

```sh
kind create cluster --config=docs/snippets/tutorials/coredns/kind.yaml

Creating cluster "coredns-etcd" ...
 ✓ Ensuring node image (kindest/node:v1.33.0) 🖼
 ✓ Preparing nodes 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-coredns-etcd"
You can now use your cluster with:

kubectl cluster-info --context kind-coredns-etcd
```

### 2. Deploy etcd as stateful set

There are multiple options to configure etcd

1. With custom manifest.
2. ETCD [manifest](https://etcd.io/docs/v3.6/op-guide/kubernetes/)
3. ETCD [operator](https://github.com/etcd-io/etcd-operator)

In this tutorial, we'll use the first option.

```sh
# apply custom manifest from external-dns repository
kubectl apply -f docs/snippets/tutorials/coredns/etcd.yaml
# wait until it's ready
kubectl rollout status statefulset etcd

❯❯ partitioned roll out complete: 1 new pods have been updated...
```

Test etcd connectivity:

```sh
kubectl exec -it etcd-0 -- etcdctl member list -wtable

+------------------+---------+--------+------------------------+-------------------------+------------+
|        ID        | STATUS  |  NAME  |       PEER ADDRS       |      CLIENT ADDRS       | IS LEARNER |
+------------------+---------+--------+------------------------+-------------------------+------------+
| 3b3ae05f90cfc535 | started | etcd-0 | http://10.244.1.3:2380 | http://etcd-0.etcd:2379 |      false |
+------------------+---------+--------+------------------------+-------------------------+------------+
```

Test etcd record management:

```sh
kubectl -n default exec -it etcd-0 -- etcdctl put /skydns/org/example/myservice '{"host":"10.0.0.10"}'
❯❯ OK

kubectl -n default exec -it etcd-0 -- etcdctl get /skydns --prefix
❯❯ /skydns/org/example/myservice
❯❯ {"host":"10.0.0.10"}

kubectl -n default exec -it etcd-0 -- etcdctl del /skydns/org/example/myservice
❯❯ 1
```

To access etcd from host:

```sh
etcdctl --endpoints=http://127.0.0.1:32379 member list
❯❯ 3b3ae05f90cfc535, started, etcd-0, http://10.244.1.3:2380, http://etcd-0.etcd:2379, false
```

### 3. Deploy CoreDNS using Helm

- [CoreDNS](https://github.com/coredns/coredns)
- [CoreDNS helm](https://github.com/coredns/helm)

```sh
helm repo add coredns https://coredns.github.io/helm
helm repo update

helm upgrade --install coredns coredns/coredns \
  -f docs/snippets/tutorials/coredns/values-coredns.yaml \
  -n default

❯❯ Release "coredns" does not exist. Installing it now.
```

Validate it's running

```sh
kubectl get pods -l app.kubernetes.io/name=coredns
```

Check the logs for errors

```sh
kubectl logs deploy/coredns -n default -c coredns --tail=50
kubectl logs deploy/coredns -n default -c resolv-check --tail=50
```

Test DNS Resolution

```sh
kubectl run -it --rm dnsutils --image=infoblox/dnstools

❯❯ curl -v http://etcd.default.svc.cluster.local:2379/version
❯❯ dig @coredns.default.svc.cluster.local kubernetes.default.svc.cluster.local
❯❯ dig @coredns.default.svc.cluster.local etcd.default.svc.cluster.local
```

### 3. Configure ExternalDNS

Deploy with helm and minimal configuration.

Add the `external-dns` helm repository and check available versions

```sh
helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
helm repo update
helm search repo external-dns --versions
```

Install with required configuration

```sh
helm upgrade --install external-dns external-dns/external-dns \
  -f docs/snippets/tutorials/coredns/values-extdns-coredns.yaml \
  -n default

❯❯ Release "external-dns" does not exist. Installing it now.
```

Validate pod status and view logs

```sh
kubectl get pods -l app.kubernetes.io/name=external-dns

kubectl logs deploy/external-dns
```

Or run it on the host from sources

```sh
export ETCD_URLS="http://127.0.0.1:32379" # port mapping configured on kind cluster

go run main.go \
    --provider=coredns \
    --source=service \
    --log-level=debug
```

### 3. Configure Test Services

Apply manifest

```sh
kubectl apply -f docs/snippets/tutorials/coredns/fixtures.yaml

kubectl get svc -l svc=test-svc

❯❯ NAME           TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
❯❯ a-g1-record    LoadBalancer   10.96.233.133   <pending>     80:31188/TCP   3m38s
❯❯ aa-g1-record   LoadBalancer   10.96.93.4      <pending>     80:31710/TCP   3m38s
```

Patch services, to manually assign an Ingress IPs. It just makes the Service appear like a real LoadBalancer for tools/tests.

```sh
kubectl patch svc a-g1-record --type=merge \
 -p '{"status":{"loadBalancer":{"ingress":[{"ip":"172.18.0.2"}]}}}' \
  --subresource=status
❯❯ service/a-g1-record patched

kubectl patch svc aa-g1-record --type=merge \
 -p '{"status":{"loadBalancer":{"ingress":[{"ip":"2001:db8::1"}]}}}' \
  --subresource=status
❯❯ service/aa-g1-record patched

kubectl get svc -l svc=test-svc

❯❯ NAME           TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
❯❯ a-g1-record    LoadBalancer   10.96.233.133   172.18.0.2    80:31188/TCP   7m13s
❯❯ aa-g1-record   LoadBalancer   10.96.93.4      2001:db8::1   80:31710/TCP   7m13s
```

### 4. Verify that records are written to etcd

Check `etcd` content. Where you should see keys similar to:

```sh
kubectl exec -it etcd-0 -- etcdctl get /skydns/org/example --prefix --keys-only

❯❯ /skydns/org/example/a-a/1acbad7e
❯❯ /skydns/org/example/a/048b0377
❯❯ /skydns/org/example/aa/2b981607
❯❯ /skydns/org/example/aaaa-aa/1228708f
```

### 5. Test DNS resolution via CoreDNS

Launch a debug pod:

```sh
kubectl run --rm -it dnsutils --image=infoblox/dnstools --restart=Never
```

Run with expected output

```sh
dig +short @coredns.default.svc.cluster.local a.example.org
❯❯ 172.18.0.2

dig +short @coredns.default.svc.cluster.local aa.example.org AAAA
❯❯ 2001:db8::1
```

### 6. PTR records (reverse DNS) — optional

Once forward DNS works, you can optionally enable reverse lookups (PTR records).
ExternalDNS can write PTR records into the same etcd backend — mapping IP addresses back
to hostnames (e.g. `172.18.0.2 → my-service`).

Two things are different from the forward-zone setup:

- ExternalDNS needs the `--managed-record-types=PTR` flag (default only manages A, CNAME, TXT).
- The `--domain-filter` must match the reverse zone you want to manage (e.g. `18.172.in-addr.arpa`).

> **CoreDNS limitation:** the etcd plugin stores exactly **one** PTR record per reverse-DNS
> key. ExternalDNS writes at the exact reverse-DNS path (e.g.
> `/skydns/arpa/in-addr/172/18/0/2`) without any random suffix, so multiple targets for
> the same IP are not supported — the last target wins.

#### Step 6a. Add the reverse zone to CoreDNS

Edit `docs/snippets/tutorials/coredns/values-coredns.yaml` and add a **second** `etcd`
plugin block for the reverse zone, right after the existing `example.org` block:

```yaml
      # existing block (forward zone)
      - name: etcd
        parameters: "example.org"
        configBlock: |
          stubzones
          path /skydns
          endpoint http://etcd.default.svc.cluster.local:2379
          fallthrough
      # ▼ NEW — reverse zone for PTR
      - name: etcd
        parameters: "18.172.in-addr.arpa"
        configBlock: |
          path /skydns
          endpoint http://etcd.default.svc.cluster.local:2379
```

Upgrade CoreDNS and wait for rollout:

```sh
helm upgrade coredns coredns/coredns \
  -f docs/snippets/tutorials/coredns/values-coredns.yaml \
  -n default
kubectl rollout status deploy/coredns
```

#### Step 6b. Reconfigure ExternalDNS for the reverse zone

When running from source, add the two flags:

```sh
export ETCD_URLS="http://127.0.0.1:32379"

go run main.go \
    --provider=coredns \
    --source=service \
    --managed-record-types=PTR \
    --domain-filter=18.172.in-addr.arpa \
    --log-level=debug
```

When running via Helm, update the values (`domainFilters` and `managedRecordTypes`):

```yaml
domainFilters:
  - 18.172.in-addr.arpa      # ← was: example.org

managedRecordTypes:
  - PTR                       # ← add PTR (default: A, CNAME, TXT)
```

```sh
helm upgrade external-dns external-dns/external-dns \
  -f docs/snippets/tutorials/coredns/values-extdns-coredns.yaml \
  -n default
```

#### Step 6c. Create a PTR-enabled service

Annotate a service with the reverse-DNS hostname (`<reversed-ip>.in-addr.arpa`).
The IP in the annotation must match the service's load-balancer IP:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  annotations:
    external-dns.kubernetes.io/hostname: 2.0.18.172.in-addr.arpa
    cluster-name: "cluster1"
spec:
  type: LoadBalancer
  ports:
    - port: 80
  selector:
    app: my-app
```

Patch the status IP so ExternalDNS picks it up:

```sh
kubectl apply -f - <<EOF
<service manifest above>
EOF

kubectl patch svc my-service --type=merge \
  -p '{"status":{"loadBalancer":{"ingress":[{"ip":"172.18.0.2"}]}}}' \
  --subresource=status
```

#### Step 6d. Verify

```sh
# The etcd key sits at the exact reverse-DNS path — no random prefix
kubectl exec -it etcd-0 -- etcdctl get /skydns/arpa/in-addr/172/18/0/2
# /skydns/arpa/in-addr/172/18/0/2
# {"host":"my-service","ttl":0}

# Reverse DNS lookup via CoreDNS
kubectl run --rm -it dnsutils --image=infoblox/dnstools --restart=Never -- \
  dig +short @coredns.default.svc.cluster.local -x 172.18.0.2
# my-service.
```

### 7. Cleanup

```sh
kind delete cluster --name coredns-etcd
```
