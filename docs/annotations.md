Annotations
===========

Unless otherwise noted, these annotations can be applied to any supported source type
(e.g. service, ingress, etc.). For the full list, see the [CLI Usage](cli-usage.md) documentation.

| Annotation                                      | Type                             | Description |
| ----------------------------------------------- | -------------------------------- | ----------- |
| external-dns.alpha.kubernetes.io/controller     | string                           | Internal use. Used to determine if a resource is owned by the current controller. |
| external-dns.alpha.kubernetes.io/hostname       | domain name                      | Sets the hostname for which this resource is a target. Overrides computed value. |
| external-dns.alpha.kubernetes.io/access         | enum: { public, private }        | Service resources only. Determines whether a node's internal or external IPs are used as the target. |
| external-dns.alpha.kubernetes.io/endpoints-type |  enum: { NodeExternalIP, HostIP } | Headless services only. Determines whether a pod's host IP, or the node's external IP address is used. |
| external-dns.alpha.kubernetes.io/target         | domain name / IP address         | Overrides the computed target address for the resource. This is useful when using a proxy not managed by Kubernetes. |
| external-dns.alpha.kubernetes.io/ttl            | integer (seconds)                | Sets the TTL of the record, in seconds. |
| external-dns.alpha.kubernetes.io/alias          | boolean                          | Forces the record to use an alias record (in certain cloud environments) instead of a CNAME.    |
| external-dns.alpha.kubernetes.io/ingress-hostname-source | enum: { annotation-only, defined-hosts-only } | Ingress resources only. Filters hostnames to include only the hostname annotation, or the hosts defined on the Ingress. |
| external-dns.alpha.kubernetes.io/internal-hostname | domain name | Service resources only. Similar to hostname, but sets the ClusterIP as the target. If the service is not of type Loadbalancer you need the `--publish-internal-services` flag. |