---
name: "\U0001F41E Bug report"
about: Report a bug encountered while operating external-dns
title: ''
labels: kind/bug
assignees: ''

---

<!--
Please use this template while reporting a bug and provide as much info as possible. Not doing so may result in your bug not being addressed in a timely manner. Thanks!

Make sure to validate the behavior against latest release https://github.com/kubernetes-sigs/external-dns/releases as we don't support past versions.

Bug Report Guide: https://kubernetes-sigs.github.io/external-dns/latest/docs/contributing/bug-report/
-->

**What happened**:

**What you expected to happen**:

**How to reproduce it (as minimally and precisely as possible)**:

<!--
Please provide as much detail as possible, including Kubernetes manifests with spec.status, ExternalDNS arguments, and logs. A bug that cannot be reproduced won't be fixed.

Provide live objects from the API server — not Helm/Terraform/Flux templates. Include ALL fields. The status section is often critical.

kubectl get <resource> -o yaml   # ingress, service, gateway, dnsendpoint, nodes, …
                                 # before and after the change if reporting a regression
-->

**Anything else we need to know?**:

**Environment**:

- External-DNS version (use `external-dns --version`):
- DNS provider:
- Others:

## Checklist

- [ ] I have searched existing issues and tried to find a fix myself
- [ ] I am using the [latest release](https://github.com/kubernetes-sigs/external-dns/releases),
  or have checked the [staging image](https://kubernetes-sigs.github.io/external-dns/latest/release/#staging-release-cycle) to confirm the bug is still reproducible
- [ ] I have provided the actual process flags (not Helm values)
- [ ] I have provided `kubectl get <resource> -o yaml` output including `status`
- [ ] I have provided full external-dns debug logs
- [ ] I have described what DNS records exist and what I expected
