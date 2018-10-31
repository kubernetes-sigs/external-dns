# Configure Service Discovery

An optional annotation `external-dns.alpha.kubernetes.io/service` may be added to a service that has an `external-dns.alpha.kubernetes.io/hostname` annotation.

When a service is configured with a hostname e.g. `ldap0.example.com`, a service annotation will cause an SRV record to be created pointing to that hostname.  The SRV record weight, priority and port are all configurable via the annotation.

In order to configure this functionality, an example Service manifest would look like the following:

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: ldap0.example.com
    external-dns.alpha.kubernetes.io/service: _ldaps._tcp.example.com 0 0 636
```

The fields of the `external-dns.alpha.kubernetes.io/service` annotation are separated by white space and are described as follows:

1. SRV record DNS name, including the service name, protocol and domain (string)
2. Priority of the record (integer)
3. Weight of the record (integer)
4. Port of the service (integer)

## Providers

- [ ] AWS (Route53)
- [ ] Azure
- [x] Cloudflare
- [ ] DigitalOcean
- [ ] Google
- [ ] InMemory
- [ ] Linode

## Implementation Details

The DNS planner uses a map internally to aggregate A record targets and select a target, as they can only have a single definition.  With record types such as SRV, MX, NS, etc. records the DNS allows multiple records for the same name.  In order to support this we create a new list based planner which takes into account both name and target to make decisions.  This is used for records that support multiple records for the same name.

The registries need to be aware that multiple records can exist for a single endpoint.  To solve this instead of maintaining a map from name to labels, we maintain a list.  When looking up a matching endpoint we also match on the target (for types that allow multiple records).  The target is encoded as `target` for this purpose.  For example `external-dns/target=0 0 636 ldap0.example.com` can be used to uniquely identify a matching endpoint.

The provisioners likewise must check the entire record for a match and not just the DNS name, as the priority, weight, port and target are relevant.  Additionally when handling registry TXT entries related to managed records they must also be able to uniquely identify the record corresponding to the related resource.
