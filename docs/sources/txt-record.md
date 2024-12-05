# Creating TXT record with CRD source

You can create and manage TXT records with the help of [CRD source](../contributing/crd-source.md)
and `DNSEndpoint` CRD. Currently, this feature is only supported by `digitalocean` providers.

In order to start managing TXT records you need to set the `--managed-record-types=TXT` flag.

```console
external-dns --source crd --provider {digitalocean} --managed-record-types=A --managed-record-types=CNAME --managed-record-types=TXT
```

Targets within the CRD need to be specified according to the RFC 1035 (section 3.3.14). Below is an example of
`example.com` DNS TXT two records creation.

**NOTE** Current implementation do not support RFC 6763 (section 6).

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: examplemxrecord
spec:
  endpoints:
    - dnsName: example.com
      recordTTL: 180
      recordType: TXT
      targets:
        - SOMETXT
        - ANOTHERTXT
```
