# kOps dns-controller

kOps includes a dns-controller that is primarily used to bootstrap the cluster, but can also be used for provisioning DNS entries for Services and Ingress.

ExternalDNS can be used as a drop-in replacement for dns-controller if you are running a non-gossip cluster. The flag `--compatibility kops-dns-controller` enables the dns-controller behaviour.

## Annotations

In kops-dns-controller compatibility mode, ExternalDNS supports two additional annotations:

* `dns.alpha.kubernetes.io/external` which is used to define a DNS record for accessing the resource publicly (i.e. public IPs)

* `dns.alpha.kubernetes.io/internal` which is used to define a DNS record for accessing the resource from outside the cluster but inside the cloud,
i.e. it will typically use internal IPs for instances.

These annotations may both be comma-separated lists of names.

## DNS record mappings

The DNS record mappings try to "do the right thing", but what this means is different for each resource type.

### Pods

For the external annotation, ExternalDNS will map a Pod to the external IPs of the Node.

For the internal annotation, ExternalDNS will map a Pod to the internal IPs of the Node.

Annotations added to Pods will always result in an A record being created.

### Services

* For a Service of Type=LoadBalancer, ExternalDNS looks at Status.LoadBalancer.Ingress. It will create CNAMEs to hostnames,
  and A records for IP addresses. It will do this for both internal and external names

* For a Service of Type=NodePort, ExternalDNS will create A records for the Node's internal/external IP addresses, as appropriate.
