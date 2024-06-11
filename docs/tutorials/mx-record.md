# Creating MX record with CRD source

You can create and manage MX records with the help of [CRD source](../contributing/crd-source.md)
and `DNSEndpoint` CRD. Currently, this feature is only supported by `aws`, `azure`, and `google` providers.

In order to start managing MX records you need to set the `--managed-record-types MX` flag.

```console
external-dns --source crd --provider {aws|azure|google} --managed-record-types A --managed-record-types CNAME --managed-record-types MX
```

Targets within the CRD need to be specified according to the RFC 1034 (section 3.6.1). Below is an example of
`example.com` DNS MX record which specifies two separate targets with distinct priorities.

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: examplemxrecord
spec:
  endpoints:
    - dnsName: example.com
      recordTTL: 180
      recordType: MX
      targets:
        - 10 mailhost1.example.com
        - 20 mailhost2.example.com
```
