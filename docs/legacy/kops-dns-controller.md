# Kops dns-controller annotations

Kops includes a dns-controller, and this document describes the existing annotations and their behaviour.  This
document is intended to allow us to see the use-cases identified by kops dns-controller, to ensure the same annotations
can be recognized (perhaps with a `--compatibilty` flag), and to ensure that we have comparable functionality.

## Flags

* `--dns`: `aws-route53,google-clouddns`

The DNS flag lets us choose which DNS provider to use.

* `--watch-ingress`	boolean

Turns ingress functionality on and off.  For AWS at least, we are blocked on switching to a release
from the `kubernetes/ingress` project (instead of one from the `contrib` project).

* `--zones` configures permitted zones, and also disambiguates when domain names are duplicated.  It is a list that matches zones we are allowed to match.

  - `*` and `*/*` are wildcard, and match all zones

  - `example.com` matches zones with name=`example.com`

  - `example.com/1234` matches zones with id=`1234` and name=`example.com`.  This is useful to disambiguate between
multiple zones named `example.com`.

  - `*/1234` matches the zone with id=`1234`.  A zone has a unique name, so this is equivalent to `example.com/1234`,
but a little shorter - and less self-documenting!

* Standard glog flags (--v, --logtostderr etc)

* Standard kubectl_util client flags


## Annotations

We define 2 primary annotations:

* `dns.alpha.kubernetes.io/external` which is used to define a DNS record for accessing the resource publicly (i.e. public IPs)

* `dns.alpha.kubernetes.io/internal` which is used to define a DNS record for accessing the resource from outside the cluster but inside the cloud,
i.e. it will typically use internal IPs for instances.

These annotations may both be comma-separated lists of names.

On a node, we also have a WIP annotation `dns.alpha.kubernetes.io/external-ip`, which configures the external ip
for a node (to work around [#42125](https://github.com/kubernetes/kubernetes/issues/42125)).  That is an annotation
that lets us defined the equivalent of an address with type ExternalIP.

## DNS record mappings

The DNS record mappings try to "do the right thing", but what this means is different for each resource type.

### Ingress

We consult the `Status.LoadBalancer.Ingress` records on the ingress.  For each one, we create a record.
If the record is an IP address, we add an A record.  If the record is a hostname (AWS ELB), we use a CNAME.

We would like to use an ALIAS, but we have not yet done this because of limitations of the DNS provider.

### Pods

For the external annotation, we will map a HostNetwork=true pod to the external IPs of the node.  We create an A record.

For the internal annotation, we will map a HostNetwork=true pod to the internal IPs of the node.  We create an A record.

We ignore pods that are not HostNetwork=true

### Services

* For a Service of Type=LoadBalancer, we look at Status.LoadBalancer.Ingress.  We create CNAMEs to hostnames,
  and A records for IP addresses.  (We should create ALIASes for ELBs).  We do this for both internal & external
  names - there is no difference on GCE or AWS.

* For a Service of Type=NodePort, we create A records for the node's internal/external IP addresses, as appropriate.

(A canonical use for NodePort internal is having a prometheus server running inside EC2 monitoring your kubernetes cluster,
for NodePort external is to expose your service without an ELB).


### Nodes

(We don't currently support annotations on the nodes themselves.  We do set up internal "alias" records,
which is how we do JOINs for e.g. NodePort services)
