# SRV record with CRD source

You can create and manage SRV records with the help of [CRD source](../contributing/crd-source.md)
and `DNSEndpoint` CRD. The implementation of this feature depends on provider API support, this feature is currently known to be supported supported by `akamai`, `civo`, `cloudflare`, `ibmcloud`, `linode`, `rfc2136` and `pdns` providers.

In order to start managing MX records you need to set the `--managed-record-types SRV` flag.

```console
external-dns --source crd --provider {akamai|civo|cloudflare|ibmcloud|linode|rfc2136|pdns} --managed-record-types A --managed-record-types CNAME --managed-record-types SRV
```

Targets within the CRD need to be specified according to the RFC 2782. Below is an example of
`example.com` DNS SRV record. It specifies a `sip` service of `udp` protocol. It has two targets
of identical priority but with different weights. They point to different backend servers.

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: examplesrvrecord
spec:
  endpoints:
    - dnsName: _sip._udp.example.com
      recordTTL: 180
      recordType: SRV
      targets:
        - 10 10 5060 sip1.example.com
        - 10 20 5060 sip2.example.com
```
