# Bug Report Guide

> **Before filing a bug:** validate the behavior against the [latest release](https://github.com/kubernetes-sigs/external-dns/releases).
> We do not support past versions.
>
> [!WARNING]
>
> The outputs requested in this guide may contain sensitive information such as
> domain names, IP addresses, cloud account IDs, annotation values, or
> credentials. Redact any sensitive values before posting them publicly in a
> GitHub issue.

Bug reports regularly arrive without the information needed to reproduce or debug
them — no process flags, no normalized Kubernetes resources, no logs — forcing
maintainers to ask multiple follow-up rounds before any investigation can start.

A bug that cannot be reproduced will not be fixed. This page explains exactly what
information to collect and how to collect it so that maintainers can reason about
your environment without making assumptions.

---

## Why we need normalized resources

external-dns only reads Kubernetes API objects at runtime. It does not read Helm
values, Terraform state, Flux kustomizations, or AWS Load Balancer Controller
annotations directly — it sees only what those tools produce in the API server.

**Please provide the live Kubernetes objects**, not the templates that generated them.

---

## Reproduce on a local cluster

If your environment is not reproducible or involves proprietary infrastructure,
the fastest path to a fix is reproducing the issue on a local cluster:
[minikube](https://minikube.sigs.k8s.io) or [kind](https://kind.sigs.k8s.io).

---

## Step-by-step: collect the required information

Work through each section below and paste the output into your issue.

### 1 — external-dns info

**Version**

```sh
kubectl get pod -n <namespace> -l app.kubernetes.io/name=external-dns \
  -o jsonpath='{.items[0].spec.containers[0].image}'
```

Or, if you have direct access to the binary:

```sh
external-dns --version
```

**Startup flags**

Helm values and Terraform variables are *not* useful here because they are
transformed before reaching the process. We need the flags that the external-dns
**process** was actually started with.

```sh
kubectl get pod -n <namespace> <pod-name> \
  -o jsonpath='{range .spec.containers[*]}{.args}{end}'
```

Example of the kind of output we need:

```text
--provider=aws
--registry=txt
--txt-owner-id=my-cluster
--source=ingress
--domain-filter=example.com
--log-level=debug
```

**Debug logs**

Enable debug logging before reproducing the issue. If external-dns is already
deployed, patch it:

```sh
kubectl set env deployment/external-dns \
  -n <namespace> \
  EXTERNAL_DNS_LOG_LEVEL=debug
```

Or add `--log-level=debug` to the process args and redeploy.

Once the pod restarts, collect logs **covering the full reconciliation loop**
that should have created or updated the record:

```sh
kubectl logs -n <namespace> \
  -l app.kubernetes.io/name=external-dns \
  --since=10m \
  --prefix=true \
  > extdns-debug.log
```

Paste the full content of `extdns-debug.log` into the issue (or attach the file).

Specifically we look for lines like:

```text
level=debug msg="Desired change: CREATE example.com A [1.2.3.4]"
level=debug msg="No endpoints could be generated from ingress ..."
level=info  msg="All records are already up to date"
```

### 2 — Kubernetes resources

Collect the full YAML — including `status` — for every resource relevant to
your source type. If reporting a regression, include the output **before and
after** the change. The `status.loadBalancer` field is critical for ingress and
service sources.

```sh
kubectl get <resource> -A -o yaml
```

Common examples by source:

```sh
kubectl get ingress,service -A -o yaml          # source=ingress
kubectl get service -A -o yaml                  # source=service
kubectl get gateway,httproute -A -o yaml        # source=gateway-httproute
kubectl get dnsendpoint -A -o yaml              # source=crd
kubectl get nodes -o yaml                       # source=node
```

### 3 — DNS provider: existing vs expected records

Tell us what records **actually exist** in your DNS provider and what you
**expected** to exist.

For Route 53:

```sh
aws route53 list-resource-record-sets \
  --hosted-zone-id <ZONE_ID> \
  --query 'ResourceRecordSets[?Name==`example.com.`]'
```

For other providers, use their CLI or API equivalent, or paste a screenshot from
the console.

Format the answer as:

| Record            | Type  | Value                         | TTL | Expected? |
|-------------------|-------|-------------------------------|-----|-----------|
| `foo.example.com` | `A`   | `1.2.3.4`                     | 300 | yes       |
| `foo.example.com` | `TXT` | `"heritage=external-dns,..."` | 300 | yes       |

### 4 — TXT ownership records

external-dns uses TXT records to track ownership. If records are not being
created or are being deleted unexpectedly, include the TXT records:

```sh
# Route 53 example — look for TXT records with "heritage=external-dns"
aws route53 list-resource-record-sets \
  --hosted-zone-id <ZONE_ID> \
  --query 'ResourceRecordSets[?Type==`TXT`]'
```

---

## Collection scripts

**external-dns info** — version, startup args, and logs:

```sh
[[% include 'snippets/contributing/collect-extdns-info.sh' %]]
```

**Kubernetes resources** — set `RESOURCE` to the resource(s) relevant to your
source (e.g. `ingress`, `"ingress,service"`, `"gateway,httproute"`,
`dnsendpoint`):

```sh
[[% include 'snippets/contributing/collect-resources.sh' %]]
```

---

## Checklist before submitting

- [ ] I have searched existing issues and tried to find a fix myself
- [ ] I am using the [latest release](https://github.com/kubernetes-sigs/external-dns/releases),
  or have checked the [staging image](../release.md#staging-release-cycle) to confirm the bug is still reproducible
- [ ] I have provided the actual process flags (not Helm values)
- [ ] I have provided `kubectl get <resource> -o yaml` output (with `status`)
- [ ] I have provided external-dns debug logs
- [ ] I have described what DNS records exist and what I expected

---

## Notes on third-party controllers

If you are using **AWS Load Balancer Controller**, **Flux**, **Terraform**, or
similar tools alongside external-dns, note that multiple controllers may be
reading and modifying the same Kubernetes objects at runtime. external-dns
maintainers can only reason about what external-dns *sees* in the API server —
please provide normalized Kubernetes objects as described above, rather than the
configuration of the surrounding tooling.

Contributors and maintainers are very unlikely to be running the same stack.
Bug reporters should assume zero shared context — no cluster access, no cloud
account, no Helm values, and no knowledge of any third-party controllers in use. A well-detailed report — see the
[checklist above](#checklist-before-submitting) — minimizes guesswork and
significantly increases the chance of resolution.
