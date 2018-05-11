## Filter-Scenarios and Filter Options

The controller supports various options controlling the
IP ranges for which DNS entries should be created or DNS
domains that should be handled.

### Option `--domain-filter`<domain name>

This option may be given multiple times. It takes a single
DNS argument representing the base domain of a hosted zone.
By default, all hosted zones available for the given
DNS provider and the given credentials are used for potential
target DNS entries.With this optio the supported hosted zones
can be explicitly selected. Despite its name this option
does not filter domains but hosted zones.

A typical scenario for the use of this option is for an
account used by the DNS controller containing more than one
hosted zones. If the actual instance should be limited to a
dedicated set of hosted zones, this option ca be given
to select the hosted zones, that should be used.

### Option `--zone-id-filter`<zone id>

This option may be given multiple times. It takes a single
id argument for a hosted zone. 

This option has a similar effect as the `--domain-filter`
option. But instead of specifying the domain name of the
hosted zone, the id of hosted zone is used.


### Option `--basedomain-filter`<domain name>

This option may be given multiple times. It takes a single
domain name argument. It may be any domain, not only 
base domain names for a hosted zone. It limits the
DNS entry generation to sub domains of the given base domain
set.

A typical scenario for the use of this option are hosted
zones shared among multiple kubernetes clusters, where every
cluster gets its own base domain. Ingresses and services
of one cluster should not use DNS entries interfering with
names used for another cluster.

This can be achieved by using this option with the base
domain name used for the actual cluster.


### Option `--cidr-ignore`<cidr>

This option may be given multiple times. It takes a single
CIDR argument. When collecting the IPs for which DNS entries
should exist, all IPs in those ranges are ignored.

A typical scenario for the use of this option is OpenStack.
The cloud provider here adds multiple IPs for a service of
type load-balancer, if it is configured to add a floating IP
to the load-balancer. The floating IP as well as the
IP of the load-balancer in the subnet it is deployed into
is propagated as IP address. In this case typically only the
floating IP should be used for the DNS entry, because the
load balancer IP is on an internal subnet.

Here this option is given with the IP of the internal
load-balancer IP to omit DNS entries for those IPs.

### Option `--dns-ignore`

This option may be given multiple times. It takes a single
DNS name argument. The DNS name might be a wildcard name.
When collecting target DNS names all entries matching the
given DNS names are ignored.

A typical scenario for the use of this option is
configuring an ingress service of type load-balancer
(for example nginx) with a wildcard DNS entry. In such
a case all subsequent DNS names are already mapped to the
ingress service and it makes no sence to generate explicit
DNS entries for ingresses matching this wildcard domain.

Here this option is given with the wildcard DNS name used
for the ingress service to avoid such explicit DNS entries.



