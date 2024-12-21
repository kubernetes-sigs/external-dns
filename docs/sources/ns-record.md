# NS record with CRD source

You can create NS records with the help of [CRD source](../contributing/crd-source.md)
and `DNSEndpoint` CRD.

In order to start managing NS records you need to set the `--managed-record-types=NS` flag.

```console
external-dns --source crd --managed-record-types=A --managed-record-types=CNAME --managed-record-types=NS
```

Consider the following example

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: ns-record
spec:
  endpoints:
  - dnsName: zone.example.com
    recordTTL: 300
    recordType: NS
    targets:
    - ns1.example.com
    - ns2.example.com
```

After instantiation of this Custom Resource external-dns will create NS record with the help of configured provider, e.g. `aws`
