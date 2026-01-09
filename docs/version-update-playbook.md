# 🧭 External-DNS Version Upgrade Playbook

## Overview

This playbook describes the best practices and steps to safely upgrade **External-DNS** in Kubernetes clusters.

Upgrading External-DNS involves validating configuration compatibility, testing changes, and ensuring no unintended DNS record modifications occur.

> Note; We strongly encourage the community to help the maintainers validate changes before they are merged or released.
> Early validation and feedback are key to ensuring stable upgrades for everyone.

---

## 1. Review Release Notes

- Visit the official [External-DNS Releases](https://github.com/kubernetes-sigs/external-dns/releases).
- Review all versions between your current and target release.
- Pay attention to:
    - **Breaking changes** (flags, CRD fields, provider behaviors). Not all changes could be captured as breaking changes.
    - **Deprecations**
    - **Provider-specific updates**
    - **Bug fixes**

> ⚠️ Breaking CLI flag or annotation changes are common in `0.x` releases.

---

## 2. Review Helm Chart and Configuration

If using Helm:

- Compare your Helm chart version to the version supporting the new app release.
- Check for:
    - `values.yaml` structural changes
    - Default arguments under `extraArgs`
    - Updates to RBAC, ServiceAccounts, or Deployment templates

---

## 3. Check Compatibility

Before upgrading, confirm:

- The new version supports your **Kubernetes version** (e.g., 1.25+).
- The **DNS provider** integration you use is still supported.

> 💡 Watch out for deprecated Kubernetes API versions (e.g., `v1/endpoints` → `discovery.k8s.io/v1/endpointslices`).

---

## 4. Test in Non-Production or with Dry Run flag

Run the new External-DNS version in a **staging cluster**.

- Use `--dry-run` mode to preview intended changes:
  - Validate logs for any unexpected record changes.
  - Ensure `external-dns` correctly identifies and plans updates without actually applying them.
  - **submit a feature request** if `dry-run` is not supported for a specific case

```yaml
args:
  - --dry-run
```

---

5. Backup DNS State

Before applying the upgrade, take a snapshot of your DNS zone(s).

**Example (AWS Route53):**

```sh
aws route53 list-resource-record-sets --hosted-zone-id ZONE_ID > backup.json
```

Use equivalent tooling for your DNS provider (Cloudflare, Google Cloud DNS, etc.).

> Having a backup ensures you can restore records if External-DNS misconfigures entries and you have a solid DR solution.

6. Perform a Controlled Rollout

Instead of upgrading in-place, use a phased rollout across multiple environments or clusters.

Recommended Approaches

a. Multi-Cluster Rollout and Progression

  1. Deploy the new `external-dns` version first in sandbox, then staging, and finally production.
  2. Monitor each environment for correct record syncing and absence of unexpected deletions.
  3. Promote the configuration only after validation in the lower environment.

b. Read-Only Parallel Deployment

  1. Run a second External-DNS instance (e.g., external-dns-readonly) with:

```yaml
args:
  - --dry-run
  - ...other flags
```

  1. Observe logs and planned record updates to confirm behavior.
  2. Observe logs and planned record updates to confirm behavior.

  7. Monitor and Validate

After deploying the new version, continuously observe both application logs and DNS synchronization metrics to ensure External-DNS behaves as expected.

**Logging**

Check logs for anomalies or unexpected record changes:

```yaml
kubectl logs -n external-dns deploy/external-dns --tail=100 -f
```

Look for:

- Creating record or Deleting record entries — validate these match expected changes.
- `WARN` or `ERROR` messages, particularly related to provider authentication or permissions.
- `TXT` registry conflicts (ownership issues between multiple instances).

If using a centralized logging stack (e.g., Loki, Elasticsearch, or CloudWatch Logs):

- Create a temporary dashboard or saved query filtering for "Creating record" OR "Deleting record".
- Correlate `external-dns` logs with DNS provider API logs to detect mismatches.

**Metrics and Observability**

Check metrics exposed by External-DNS (if Prometheus scraping is enabled):

Focus on:

- Error rate (*_errors_total)
- Number of syncs per interval (*_sync_duration_seconds)
- Provider API call spikes

Example PromQL checks:

```promql
rate(external_dns_registry_errors_total[5m]) > 0
rate(external_dns_provider_requests_total{operation="DELETE"}[5m])
```

## External Verification

Ideally, you should have a set of automated tests

Query key DNS records directly:

```sh
dig +short myapp.example.com
nslookup api.staging.example.com
```

Ensure that A, CNAME, and TXT records remain correct and point to expected endpoints.

Additional Tips

- Automate upgrade testing with CI/CD pipelines.
- Maintain clear CHANGELOGs and migration notes for internal users.
- Tag known good versions in Git or Helm values for rollback.
- Avoid skipping multiple minor versions when possible.
