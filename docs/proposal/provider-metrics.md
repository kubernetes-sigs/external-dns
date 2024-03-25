# Exporting Provider-Specific Metrics in External-DNS
*(December 2023)*

## Purpose
- The purpose of this proposal is to introduce the functionality to export provider-specific metrics in the external-dns project.
- This enhancement will enable users to gain insights into provider-specific behavior, leading to better management and troubleshooting of DNS services.

## Motivating Usecases
- **Rate Limit Monitoring** - Many providers implement API rate-limits. Per-provider metrics would enable users to detect and respond when they are approaching these limits:
   - Cloudflare: [#2135](https://github.com/kubernetes-sigs/external-dns/issues/2135)
   - AWS: [#1293](https://github.com/kubernetes-sigs/external-dns/issues/1293)
   - DigitalOcean: [#1429](https://github.com/kubernetes-sigs/external-dns/issues/1429)

- **Performance/Latency Monitoring** - Changes to the volume/content of provider API requests may impact performance of external-dns overall. Per-provider metrics can help identify and debug these regressions.
  - AWS: [#869](https://github.com/kubernetes-sigs/external-dns/issues/869)

- **Debugging** - Per-provider metrics allow for specific failure conditions to be surfaced to users via metrics as opposed to logs. This could expidite investigations.

_This list is non-exhaustive_

## High-Level Design

### Metric Format
Provider-specific metrics will be exported with the format  `external_dns_provider_{metric_name}`. They
should have a `provider` label that indicates the provider the metric is associated with. 

Certain types of metrics (such as request counts, latency) are expected to conform to a common specification to make the experience consistent across providers:

- `external_dns_provider_request_count`
  - Counts the number of requests issued.
  - Type: counter
  - Labels:
     - `method` - The method being invoked (eg. `ListTagsForResourceWithContext`).
     - `error` - Whether the invocation of this method resulted in an error.
     - `error_code` - Error code returned from the method. 
- `external_dns_provider_request_latency_ms`
  - The latency (in ms) of requests issued to each method.
  - Type: Histogram
  - Labels:
     - `method` - The method being invoked (eg. `ListTagsForResourceWithContext`).
- `external_dns_provider_request_rate`
  - Measures the peak per-second request rate using a sliding window. 
     - Ratelimits are often enforced on a more granular time resolution than Prometheus metrics are collected. This metric pre-computes a peak request rate over a sliding window, to give users observability into bursty API requests.
     - This metric has two parameters:
        - The `requestWindowDurationSec` corresponds with the enforcement window of the ratelimit. For instance, AWS applies a 5 request per second limit, so the `requestWindowDurationSec` would be 1s. If ProviderX applied a 100 request per minute limit, the `requestWindowDurationSec` for ProviderX would be 60.
        - The `slidingWindowDurationSec` is the period over which the client-side aggregation finds the maximum. This window should be larger than the Prometheus scrape window, so that spikes are guaranteed to be collected by Prometheus. The default value is 60s.
  - Type: gauge
  - Labels:
     - `function` - The function being invoked (eg. `ListTagsForResourceWithContext`).
     - `requestWindowDurationSec` - The number of seconds in the ratelimit evaluation window (eg. 1s for AWS).
     - `slidingWindowDurationSec` - The number of seconds in the sliding window. 

### Configuration
We will introduce a single flag that can be used to enable flags on a per-provider basis:

```
--metrics-enabled=aws
```

All metrics flags will be disabled by default. The `slidingWindowDurationSec` parameter is also configurable:

```
--metrics-sliding-window-duration=60s
```

## Implementation Plan

Implementation details will be filled-in on a provider-by-provider basis. This document will be updated as POCs commit to add this functionality to current providers:

| Provider | Author | PRs |
|----------|--------|-----|
|          |        |     |
|          |        |     |
|          |        |     |

## Open Questions / FAQ

### How will this work with out-of-tree providers (i.e. webhook providers)?
  - The in-tree Webhook Provider already [exports metrics](https://github.com/kubernetes-sigs/external-dns/blob/master/provider/webhook/webhook.go#L41) about requests issued to the external provider. This metrics will not be impacted by this proposal.
  - If an out-of-tree Webhook implementation wishes to export metrics relating to the requests it is making to external services, it should follow this proposal.
