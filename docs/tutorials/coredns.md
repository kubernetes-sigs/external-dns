# CoreDNS

- [Documentation](https://coredns.io/)

## Multi cluster support options

The CoreDNS provider allows records from different CoreDNS providers to be separated in a single etcd
by setting the `--coredns-managed-by` flag to a unique identifier. Setting the flag
to `--coredns-strict-managed-by` prevents the inclusion of records that do not have an identifier.

## Specific service annotation options

### Groups

Groups can be used to group set of services together. The main use of this is to limit recursion,
i.e. don't return all records, but only a subset. Let's say we have a configuration like this:

```yaml
[[% include 'tutorials/coredns/coredns-groups.yaml' %]]
```

And we want domain.local to return (127.0.0.1 and 127.0.0.2) and subdom.domain.local to return (127.0.0.3 and 127.0.0.4).
For this the two domains, need to be in different groups. What those groups are does not matter,
as long as a and b belong to the same group which is different from the group c and d belong to.
If a service is found without a group it is always included.
