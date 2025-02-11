# Service source

The service source creates DNS entries based on `Service` resources.

## Filtering the Services considered

The `--service-type-filter` flag filters Service resources by their `spec.type`.
The flag may be specified multiple times to allow multiple service types.

This source supports the `--label-filter` flag, which filters Service resources
by a set of labels.

## Domain names

The domain names of the DNS entries created from a Service are sourced from the following places:

1. Adds the domain names from any `external-dns.alpha.kubernetes.io/hostname` and/or
`external-dns.alpha.kubernetes.io/internal-hostname` annotation.
This behavior is suppressed if the `--ignore-hostname-annotation` flag was specified.

2. If no DNS entries were produced for a Service by the previous steps
and the `--compatibility` flag was specified, then adds DNS entries per the
selected compatibility mode.

3. If no DNS entries were produced for a Service by the previous steps
or the `--combine-fqdn-annotation` flag was specified, then adds domain names
generated from any`--fqdn-template` flag.

### Domain names for headless service pods

If a headless Service (without an `external-dns.alpha.kubernetes.io/target` annotation) creates DNS entries with targets from
a Pod that has a non-empty `spec.hostname` field, additional DNS entries are created for that Pod, containing the targets from that Pod.
For each domain name created for the Service, the additional DNS entry for the Pod has that domain name prefixed with
the value of the Pod's `spec.hostname` field and a `.`.

## Targets

If the Service has an `external-dns.alpha.kubernetes.io/target` annotation, uses
the values from that. Otherwise, the targets of the DNS entries created from a service are sourced depending
on the Service's `spec.type`:

### LoadBalancer

1. If the hostname came from an `external-dns.alpha.kubernetes.io/internal-hostname` annotation, uses
the Service's `spec.clusterIP` field. If that field has the value `None`, does not generate
any targets for the hostname.

2. Otherwise, if the Service has one or more `spec.externalIPs`, uses the values in that field.

3. Otherwise, iterates over each `status.loadBalancer.ingress`, adding any non-empty `ip` and/or `hostname`.

If the `--resolve-service-load-balancer-hostname` flag was specified, any non-empty `hostname`
is queried through DNS and any resulting IP addresses are added instead.
A DNS query failure results in zero targets being added for that load balancer's ingress hostname.

### ClusterIP (headless)

Iterates over all of the Service's Endpoints's `subsets.addresses`.
If the Service's `spec.publishNotReadyAddresses` is `true` or the `--always-publish-not-ready-addresses` flag is specified,
also iterates over the Endpoints's `subsets.notReadyAddresses`.

1. If an address does not target a `Pod` that matches the Service's `spec.selector`, it is ignored.

2. If the target pod has an `external-dns.alpha.kubernetes.io/target` annotation, uses
the values from that.

3. Otherwise, if the Service has an `external-dns.alpha.kubernetes.io/endpoints-type: NodeExternalIP`
annotation, uses the addresses from the Pod's Node's `status.addresses` that are either of type
`ExternalIP` or IPv6 addresses of type `InternalIP`.

4. Otherwise, if the Service has an `external-dns.alpha.kubernetes.io/endpoints-type: HostIP` annotation
or the `--publish-host-ip` flag was specified, uses the Pod's `status.hostIP` field.

5. Otherwise uses the `ip` field of the address from the Endpoints.

### ClusterIP (not headless)

1. If the hostname came from an `external-dns.alpha.kubernetes.io/internal-hostname` annotation
or the `--publish-internal-services` flag was specified, uses the `spec.ClusterIP`.

2. Otherwise, does not create any targets.

### NodePort

If `spec.ExternalTrafficPolicy` is `Local`, iterates over each Node that both matches the Service's `spec.selector`
and has a `status.phase` of `Running`. Otherwise iterates over all Nodes, of any phase.

Iterates over each relevant Node's `status.addresses`:

1. If there is an `external-dns.alpha.kubernetes.io/access: public` annotation on the Service, uses both addresses with
a `type` of `ExternalIP` and IPv6 addresses with a `type` of `InternalIP`.

2. Otherwise, if there is an `external-dns.alpha.kubernetes.io/access: private` annotation on the Service, uses addresses with
a `type` of `InternalIP`.

3. Otherwise, if there is at least one address with a `type` of `ExternalIP`, uses both addresses with
a `type` of `ExternalIP` and IPv6 addresses with a `type` of `InternalIP`.

4. Otherwise, uses addresses with a `type` of `InternalIP`.

Also iterates over the Service's `spec.ports`, creating a SRV record for each port which has a `nodePort`.
The SRV record has a service of the Service's `name`, a protocol taken from the port's `protocol` field,
a priority of `0` and a weight of `50`.
In order for SRV records to be created, the `--managed-record-types` must have been specified, including `SRV`
as one of the values.

```console
external-dns ... --managed-record-types=A --managed-record-types=CNAME --managed-record-types=SRV
```

### ExternalName

1. If the Service has one or more `spec.externalIPs`, uses the values in that field.
2. Otherwise, creates a target with the value of the Service's `externalName` field.
