# Multi-Cluster Shared DNS Records

Several ExternalDNS instances — typically one per cluster — may need to publish the same hostname, for failover
clusters, blue/green rollouts, or multi-region deployments.

- [Pattern A](#pattern-a-per-cluster-hostnames-plus-an-aggregating-record) — per-cluster hostnames plus an
  aggregating record. Any provider.
- [Pattern B](#pattern-b-one-owning-cluster-others-follow) — one owning cluster, the others stay out. Any provider.
- [Pattern C](#pattern-c-one-route53-record-set-per-cluster-same-hostname-only-on-aws) — one Route53 record set
  per cluster, all under the same hostname. AWS Route53 only.
- [Pattern D](#pattern-d-per-target-ownership-on-coredns) — every cluster contributes targets to one name.
  CoreDNS only.

Every pattern needs a distinct `--txt-owner-id` per cluster. Sharing one in a shared zone makes each cluster
delete the other's records — see
[OwnerID migration](../registry/txt.md#ownerid-migration-multi-cluster-considerations).

## Why a second instance appears to do nothing

With the [TXT registry](../registry/txt.md), each managed record is claimed by exactly one instance via
`--txt-owner-id`, stored in a companion TXT record. A second instance wanting a record another owner holds
skips it — otherwise both would force their own targets onto it and flip it every sync interval.

The skip is silent at the default log level. `All records are already up to date` is the ordinary zero-change
message, not a mismatch signal. Use `--log-level=debug`:

```text
Skipping endpoint app.example.com 60 IN A  198.51.100.20 [] because owner id does not match (found: "cluster-a", required: "cluster-b")
```

A near-identical message (`... for one or more items to create ...`) covers *creating* a record under a name
another owner holds — cluster B publishing an `AAAA` where cluster A owns the `A`. Only that case increments
the `external_dns_registry_skipped_records_owner_mismatch_per_sync` gauge, see
[Monitoring](../monitoring/index.md).

## Pattern A: per-cluster hostnames plus an aggregating record

Each cluster publishes its own hostname (`app.eu.example.com`, `app.us.example.com`), and the shared
`app.example.com` is a separate record listing every cluster's address, managed with a
[DNSEndpoint CRD](../sources/crd.md) in one designated cluster:

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

The aggregate must be an A record: a CNAME must be the only record for its name
([RFC 1034 §3.6.2](https://datatracker.ietf.org/doc/html/rfc1034#section-3.6.2)), so a multi-target CNAME is
invalid and providers handle it inconsistently.

It also has to be edited whenever a load balancer address changes or a cluster comes and goes. Where addresses
are unstable or health checking is needed, a provider-side traffic-management feature can replace it and point
at the per-cluster hostnames instead.

## Pattern B: one owning cluster, others follow

When the shared record needs a single set of targets and a single manager, designate one cluster and stop the
others from managing that name: `--domain-filter` excludes the name, `--annotation-filter` and `--label-filter`
exclude the Kubernetes object (label selector semantics) when the same manifest ships everywhere and only one
copy carries the marker. Simply not annotating the resource elsewhere works too. `--policy=upsert-only` on the
non-owning instances adds defense in depth, though ownership already stops them deleting what they do not own.

## Pattern C: one Route53 record set per cluster, same hostname (only on AWS)

Route53 calls its unit of configuration a *resource record set*; the console labels it a *record*. Several may
share a name and type as long as each carries a distinct set identifier and a routing policy. ExternalDNS keys
its plan on `(DNS name, set identifier)`, so clusters using different set identifiers never collide.

`external-dns.kubernetes.io/set-identifier` is supported by the AWS provider only — the OCI provider warns and
ignores it, others do not implement it.

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

Ownership TXT records inherit the set identifier and routing policy, so each cluster gets its own TXT record
set and no `--txt-prefix` juggling is needed. TXT names below use the record-type prefix of the current
[TXT registry format](../registry/txt.md#for-version-v018) (v0.18+); older formats carry the set identifier
too, without that prefix:

```text
app.example.com     A    (set-identifier: cluster-a) -> <cluster A LB address>
app.example.com     A    (set-identifier: cluster-b) -> <cluster B LB address>
a-app.example.com   TXT  (set-identifier: cluster-a) -> "heritage=external-dns,external-dns/owner=cluster-a,..."
a-app.example.com   TXT  (set-identifier: cluster-b) -> "heritage=external-dns,external-dns/owner=cluster-b,..."
```

Each cluster manages and deletes only its own record set: removing the Service in cluster B leaves cluster A
untouched. Weights of `90`/`10` drain a cluster progressively, and weight `0` takes it out of rotation without
deleting anything — but Route53 serves a weight `0` record when *every* record in the group is `0`, so never
set them all at once.

Other routing annotations work the same way: `aws-region` for latency-based, `aws-failover` for
active/passive, `aws-multi-value-answer` to answer with several clusters at once. The last is not round-robin
— Route53 returns
[up to eight healthy records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-multivalue.html)
chosen roughly at random, varying per resolver, and returns all of them when there are eight or fewer. See
[AWS routing policies](../tutorials/aws.md#routing-policies).

> Route53 does not allow multivalue answer alias records, and a `type: LoadBalancer` Service produces an alias
> record by default on AWS. `aws-multi-value-answer` therefore only applies to non-alias records — with
> `--aws-prefer-cname`, or when targets are plain IP addresses. Weighted, latency, failover and geolocation
> routing all work with alias records.

## Pattern D: per-target ownership on CoreDNS

`--coredns-strictly-owned` adds an ownership layer underneath the TXT registry: the provider stamps an `owner`
field on every etcd entry it writes — one entry per target, registry TXT entries included — and filters reads
by it. Each instance sees only its own entries, so two clusters never contend for the same name, while CoreDNS
answers with every entry stored under it. Both clusters publishing `app.example.com` therefore produce one
A response carrying both targets, and removing the workload in one cluster takes its entries — the ownership
one included — out of etcd while leaving the other's target in place.

## Not supported: merging targets with the TXT registry

Under the [TXT registry](../registry/txt.md), two instances cannot each contribute targets to one record set —
cluster A adding `1.1.1.1` and cluster B `2.2.2.2` to the same `app.example.com` A record. Its ownership TXT
record is per DNS name and carries exactly one owner, so targets cannot be attributed per cluster, nor removed
when a cluster goes away. On providers without set identifier support, the first cluster to create the record
owns it and the second skips it forever.

Use Pattern C on AWS, Pattern D on CoreDNS, or Pattern A elsewhere. See
[issue #1441](https://github.com/kubernetes-sigs/external-dns/issues/1441) for the discussion.
