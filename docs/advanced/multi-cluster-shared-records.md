# Multi-Cluster Shared DNS Records

Multiple ExternalDNS instances — typically one per cluster — may need to publish the same hostname, for
failover clusters, blue/green rollouts, or multi-region deployments.

- [Pattern A](#pattern-a-per-cluster-hostnames-plus-an-aggregating-record) — per-cluster hostnames plus an
  aggregating record. Any provider.
- [Pattern B](#pattern-b-one-owning-cluster-others-follow) — one owning cluster, the others stay out. Any provider.
- [Pattern C](#pattern-c-one-record-per-cluster-same-hostname-only-on-aws) — one record set per cluster under the
  same hostname. AWS Route53 only.

## Why a second instance appears to do nothing

Every managed record is claimed by exactly one instance, identified by `--txt-owner-id` and stored in a
companion TXT record (see [TXT registry](../registry/txt.md)). When a second instance wants a record another
owner already holds, it skips it and reports `All records are already up to date`. With `--log-level=debug`:

```text
Skipping endpoint app.example.com 60 IN A  198.51.100.20 [] because owner id does not match (found: "cluster-a", required: "cluster-b")
```

This is a safety property: without it, two instances would force their own targets onto the same record set and
flip it back and forth every sync interval.

A near-identical message (`... for one or more items to create ...`) appears when the instance tries to *create* a
record under a name another owner holds — cluster B publishing an `AAAA` where cluster A owns the `A`. Only that
case increments the `external_dns_registry_skipped_records_owner_mismatch_per_sync` gauge — see
[Monitoring](../monitoring/index.md).

Every pattern below requires a distinct `--txt-owner-id` per cluster. Never share one across clusters in a
shared zone: each cluster would then delete the other's records. See
[OwnerID migration: multi-cluster considerations](../registry/txt.md#ownerid-migration-multi-cluster-considerations).

## Pattern A: per-cluster hostnames plus an aggregating record

Works with any provider. Each cluster publishes a hostname unique to itself (`app.eu.example.com`,
`app.us.example.com`), and the shared `app.example.com` is a separate record listing every cluster's address —
managed with a [DNSEndpoint CRD](../sources/crd.md) in one designated cluster:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: app-aggregate
spec:
  endpoints:
    - dnsName: app.example.com
      recordType: A
      recordTTL: 60
      targets:
        - 198.51.100.10 # cluster EU load balancer
        - 203.0.113.10  # cluster US load balancer
```

Note the aggregate is an A record. A CNAME must be the only record for its name
([RFC 1034 §3.6.2](https://datatracker.ietf.org/doc/html/rfc1034#section-3.6.2)), so a multi-target CNAME is
invalid and providers handle it inconsistently — some reject it, some silently keep the first target.

The aggregate must therefore be edited whenever a load balancer address changes, or a cluster is added or
removed. Where addresses are not stable, or health checking is needed, a provider-side traffic-management
feature can replace the aggregate record and point at the per-cluster hostnames ExternalDNS keeps up to date.

## Pattern B: one owning cluster, others follow

If the shared record needs a single set of targets and a single manager, designate one cluster as the owner and
stop the others from managing that name. `--domain-filter` excludes the name itself; `--annotation-filter` and
`--label-filter` exclude the Kubernetes object using label selector semantics — useful when the same manifest
ships to every cluster and only one carries the marker. Not annotating the resource elsewhere works too.
`--policy=upsert-only` on the non-owning instances adds defense in depth, though ownership already prevents
them from deleting records they do not own.

## Pattern C: one record per cluster, same hostname (only on AWS)

Route53 allows several record sets to share a name and type as long as each carries a distinct set identifier
and a routing policy. ExternalDNS keys its plan on `(DNS name, set identifier)`, so clusters using different
set identifiers never collide — each owns its own record set under the same hostname.

`external-dns.kubernetes.io/set-identifier` is supported by the AWS provider only. The OCI provider logs a
warning and ignores it; other providers do not implement it.

Traffic split evenly across two clusters, annotating the same Service in each:

```yaml
# cluster A
external-dns.kubernetes.io/hostname: app.example.com
external-dns.kubernetes.io/set-identifier: cluster-a
external-dns.kubernetes.io/aws-weight: "50"
# cluster B
external-dns.kubernetes.io/hostname: app.example.com
external-dns.kubernetes.io/set-identifier: cluster-b
external-dns.kubernetes.io/aws-weight: "50"
```

Resulting zone contents — the ownership TXT records inherit both the set identifier and the routing policy, so
each cluster gets its own TXT record set and no `--txt-prefix` juggling is required. The TXT names below use the
record-type prefix of the current [TXT registry format](../registry/txt.md#for-version-v018) (v0.18+); older
formats carry the set identifier too, but without that prefix:

```text
app.example.com     A    (set-identifier: cluster-a) -> <cluster A LB address>
app.example.com     A    (set-identifier: cluster-b) -> <cluster B LB address>
a-app.example.com   TXT  (set-identifier: cluster-a) -> "heritage=external-dns,external-dns/owner=cluster-a,..."
a-app.example.com   TXT  (set-identifier: cluster-b) -> "heritage=external-dns,external-dns/owner=cluster-b,..."
```

Each cluster manages and deletes only its own record set: removing the Service in cluster B leaves cluster A
untouched. Shifting weights to `90`/`10` drains a cluster progressively, and weight `0` takes it out of rotation
without deleting anything — Route53 only serves a weight `0` record when *every* record in the group is `0`, so
do not set all of them to `0` at once. Swap the routing annotation for other behaviors — `aws-region` for
latency-based, `aws-failover` for active/passive, `aws-multi-value-answer` for round-robin over all clusters.
See [AWS routing policies](../tutorials/aws.md#routing-policies).

> Route53 does not allow multivalue answer alias records, and a `type: LoadBalancer` Service produces an alias
> record by default on AWS. `aws-multi-value-answer` therefore only applies to non-alias records — with
> `--aws-prefer-cname`, or when targets are plain IP addresses. Weighted, latency, failover and geolocation
> routing all work with alias records.

## Not supported: merging targets from several instances into one record

Two instances cannot each contribute targets to a single record set — cluster A adding `1.1.1.1` and cluster B
adding `2.2.2.2` to the same `app.example.com` A record. On providers without set identifier support, the first
cluster to create the record owns it and the second skips it forever.

This follows from the ownership model: the ownership TXT record is per DNS name and carries exactly one owner,
so individual targets cannot be attributed to individual clusters — and one cluster's target could never be
removed when that cluster goes away.

Use Pattern C on AWS, or Pattern A on any other provider. See
[issue #1441](https://github.com/kubernetes-sigs/external-dns/issues/1441) for the discussion.
