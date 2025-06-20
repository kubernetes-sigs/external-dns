# DNS provider API rate limits considerations

## Introduction

By design, external-dns refreshes all the records of a zone using API calls.
This refresh may happen peridically and upon any changed object if the flag `--events` is enabled.

Depending on the size of the zone and the infrastructure deployment, this may lead to external-dns
hitting the DNS provider's rate-limits more easily.

In particular, it has been found that with 200k records in an AWS Route53 zone, each refresh triggers around
70 API calls to retrieve all the records, making it more likely to hit the AWS Route53 API rate limits.

To prevent this problem from happening, external-dns has implemented a cache to reduce the pressure on the DNS
provider APIs.

This cache is optional and systematically invalidated when DNS records have been changed in the cluster
(new or deleted domains or changed target).

## Trade-offs

The major trade-off of this setting relies in the ability to recover from a deleted record on the DNS provider side.
As the DNS records are cached in memory, external-dns will not be made aware of the missing records and will hence
take a longer time to restore the deleted or modified record on the provider side.

This option is enabled using the `--provider-cache-time=15m` command line argument, and turned off when `--provider-cache-time=0m`

## Monitoring

You can evaluate the behaviour of the cache thanks to the built-in metrics

* `external_dns_provider_cache_records_calls`
  * The number of calls to the provider cache Records list.
  * The label `from_cache=true` indicates that the records were retrieved from memory and the DNS provider was not reached
  * The label `from_cache=false` indicates that the cache was not used and the records were retrieved from the provider
* `external_dns_provider_cache_apply_changes_calls`
  * The number of calls to the provider cache ApplyChanges.
  * Each ApplyChange systematically invalidates the cache and makes subsequent Records list to be retrieved from the provider without cache.

## Related options

This global option is available for all providers and can be used in pair with other global
or provider-specific options to fine-tune the behaviour of external-dns
to match the specific needs of your deployments, with the goal to reduce the number of API calls to your DNS provider.

* Google
  * `--google-batch-change-interval=1s` When using the Google provider, set the interval between batch changes. ($EXTERNAL_DNS_GOOGLE_BATCH_CHANGE_INTERVAL)
  * `--google-batch-change-size=1000` When using the Google provider, set the maximum number of changes that will be applied in each batch.
* AWS
  * `--aws-batch-change-interval=1s` When using the AWS provider, set the interval between batch changes.
  * `--aws-batch-change-size=1000` When using the AWS provider, set the maximum number of changes that will be applied in each batch.
  * `--aws-batch-change-size-bytes=32000` When using the AWS provider, set the maximum byte size that will be applied in each batch.
  * `--aws-batch-change-size-values=1000` When using the AWS provider, set the maximum total record values that will be applied in each batch.
  * `--aws-zones-cache-duration=0s` When using the AWS provider, set the zones list cache TTL (0s to disable).
  * `--[no-]aws-zone-match-parent` Expand limit possible target by sub-domains
* Cloudflare
  * `--cloudflare-dns-records-per-page=100` When using the Cloudflare provider, specify how many DNS records listed per page, max possible 5,000 (default: 100)
* OVH
  * `--ovh-api-rate-limit=20` When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)

* Global
  * `--registry=txt` The registry implementation to use to keep track of DNS record ownership.
    * Other registry options such as dynamodb can help mitigate rate limits by storing the registry outside of the DNS hosted zone (default: txt, options: txt, noop, dynamodb, aws-sd)
  * `--txt-cache-interval=0s` The interval between cache synchronizations in duration format (default: disabled)
  * `--interval=1m0s` The interval between two consecutive synchronizations in duration format (default: 1m)
  * `--min-event-sync-interval=5s` The minimum interval between two consecutive synchronizations triggered from kubernetes events in duration format (default: 5s)
  * `--[no-]events` When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)

A general recommendation is to enable `--events` and keep `--min-event-sync-interval` relatively low to have a better responsiveness when records are
created or updated inside the cluster.
This should represent an acceptable propagation time between the creation of your k8s resources and the time they become registered in your DNS server.

On a general manner, the higher the `--provider-cache-time`, the lower the impact on the rate limits, but also, the slower the recovery in case of a deletion.
The `--provider-cache-time` value should hence be set to an acceptable time to automatically recover restore deleted records.

✍️ Note that caching is done within the external-dns controller memory. You can invalidate the cache at any point in time by restarting it (for example doing a rolling update).
