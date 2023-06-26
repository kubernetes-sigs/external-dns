# Registries

A registry persists metadata pertaining to DNS records.

The most important metadata is the owning external-dns deployment.
This is specified using the `--txt-owner-id` flag, specifying a value unique to the
deployment of external-dns and which doesn't change for the lifetime of the deployment.
Deployments in different clusters but sharing a DNS zone need to use different owner IDs.

The registry implementation is specified using the `--registry` flag.

## Supported registries

* [txt](txt.md) (default) - Stores metadata in TXT records in the same provider.
* [dynamodb](dynamodb.md) - Stores metadata in an AWS DynamoDB table.
* noop - Passes metadata directly to the provider. For most providers, this means the metadata is not persisted.
* aws-sd - Stores metadata in AWS Service Discovery. Only usable with the `aws-sd` provider.
