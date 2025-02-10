# The DynamoDB registry

As opposed to the default TXT registry, the DynamoDB registry stores DNS record metadata in an AWS DynamoDB table instead of in TXT records in a hosted zone.
This following tutorial extends [Setting up ExternalDNS for Services on AWS](../tutorials/aws.md) to use the DynamoDB registry instead.

## IAM permissions

The ExternalDNS [IAM Policy](../tutorials/aws.md#iam-policy) must additionally be granted the following permissions:

```json
    {
      "Effect": "Allow",
      "Action": [
        "DynamoDB:DescribeTable",
        "DynamoDB:PartiQLDelete",
        "DynamoDB:PartiQLInsert",
        "DynamoDB:PartiQLUpdate",
        "DynamoDB:Scan"
      ],
      "Resource": [
        "arn:aws:dynamodb:*:*:table/external-dns"
      ]
    }
```

The region and account ID may be specified explicitly specified instead of using wildcards.

## Create a DynamoDB Table

By default, the DynamoDB registry stores data in the table named `external-dns` and it needs to exist before configuring ExternalDNS to use the DynamoDB registry.
If the DynamoDB table has a different name, it may be specified using the `--dynamodb-table` flag.
If the DynamoDB table is in a different region, it may be specified using the `--dynamodb-region` flag.

The following command creates a DynamoDB table with the name: `external-dns`:

> The table must have a partition (HASH) key named `k` of type string (`S`) and the table must NOT have a sort (RANGE) key.

```bash
aws dynamodb create-table \
  --table-name external-dns \
  --attribute-definitions \
    AttributeName=k,AttributeType=S \
  --key-schema \
    AttributeName=k,KeyType=HASH \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --table-class STANDARD
```

## Set up a hosted zone

Follow [Set up a hosted zone](../tutorials/aws.md#set-up-a-hosted-zone)

## Modify ExternalDNS deployment

The ExternalDNS deployment from [Deploy ExternalDNS](../tutorials/aws.md#deploy-externaldns) needs the following modifications:

* `--registry=txt` should be changed to `--registry=dynamodb`
* Add `--dynamodb-table=external-dns` to specify the name of the DynamoDB table, its value defaults to `external-dns`
* Add `--dynamodb-region=us-east-1` to specify the region of the DynamoDB table

For example:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  labels:
    app.kubernetes.io/name: external-dns
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app.kubernetes.io/name: external-dns
  template:
    metadata:
      labels:
        app.kubernetes.io/name: external-dns
    spec:
      containers:
        - name: external-dns
          image: registry.k8s.io/external-dns/external-dns:v0.15.1
          args:
            - --source=service
            - --source=ingress
            - --domain-filter=example.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
            - --provider=aws
            - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
            - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
            - --registry=dynamodb # previously, --registry=txt
            - --dynamodb-table=external-dns # defaults to external-dns
            - --dynamodb-region=us-east-1 # set to the region the DynamoDB table in
            - --txt-owner-id=my-hostedzone-identifier
          env:
            - name: AWS_DEFAULT_REGION
              value: us-east-1 # change to region where EKS is installed
      # # Uncomment below if using static credentials
      #       - name: AWS_SHARED_CREDENTIALS_FILE
      #        value: /.aws/credentials
      #     volumeMounts:
      #       - name: aws-credentials
      #         mountPath: /.aws
      #         readOnly: true
      # volumes:
      #   - name: aws-credentials
      #     secret:
      #       secretName: external-dns
```

## Validate ExternalDNS works

Create either a [Service](../tutorials/aws.md#verify-externaldns-works-service-example) or an [Ingress](../tutorials/aws.md#verify-externaldns-works-ingress-example) and

After roughly two minutes, check that the corresponding entry was created in the DynamoDB table:

```bash
aws dynamodb scan --table-name external-dns
```

This will show something like:

```json
{
    "Items": [
        {
            "k": {
                "S": "nginx.example.com#A#"
            },
            "o": {
                "S": "my-identifier"
            },
            "l": {
                "M": {
                    "resource": {
                        "S": "service/default/nginx"
                    }
                }
            }
        }
    ],
    "Count": 1,
    "ScannedCount": 1,
    "ConsumedCapacity": null
}
```

## Clean up

In addition to the clean up steps in [Setting up ExternalDNS for Services on AWS](../tutorials/aws.md#clean-up), delete the DynamoDB table that was used as a registry.

```bash
aws dynamodb delete-table \
  --table-name external-dns
```

## Caching

The DynamoDB registry can optionally cache DNS records read from the provider. This can mitigate rate limits imposed by the provider.

Caching is enabled by specifying a cache duration with the `--txt-cache-interval` flag.

## Migration from TXT registry

If any ownership TXT records exist for the configured owner, the DynamoDB registry will migrate
the metadata therein to the DynamoDB table. If any such TXT records exist, any previous values for
`--txt-prefix`, `--txt-suffix`, `--txt-wildcard-replacement`, and `--txt-encrypt-aes-key`
must be supplied.

If TXT records are in the set of managed record types specified by `--managed-record-types`,
it will then delete the ownership TXT records on a subsequent reconciliation.
