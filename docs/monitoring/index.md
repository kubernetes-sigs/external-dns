# Monitoring & Observability

Monitoring is a crucial aspect of maintaining the health and performance of your applications.
It involves collecting, analyzing, and using information to ensure that your system is running smoothly and efficiently. Effective monitoring helps in identifying issues early, understanding system behavior, and making informed decisions to improve performance and reliability.

For `external-dns`, all metrics available for scraping are exposed on the `/metrics` endpoint. The metrics are in the Prometheus exposition format, which is widely used for monitoring and alerting.

To access the metrics:

```sh
curl https://localhost:7979/metrics
```

In the metrics output, you'll see the help text, type information, and current value of the `external_dns_registry_endpoints_total` counter:

```yml
# HELP external_dns_registry_endpoints_total Number of Endpoints in the registry
# TYPE external_dns_registry_endpoints_total gauge
external_dns_registry_endpoints_total 11
```

You can configure a locally running [Prometheus instance](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config) to scrape metrics from the application. Here's an example prometheus.yml configuration:

```yml
scrape_configs:
- job_name: external-dns
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:7979
```

For more detailed information on how to instrument application with Prometheus, you can refer to the [Prometheus Go client library documentation](https://prometheus.io/docs/guides/go-application/).

## What metrics can I get from ExternalDNS and what do they mean?

- The project maintain a [metrics page](./metrics.md) with a list of supported custom metrics.
- [Go runtime](https://pkg.go.dev/runtime/metrics#hdr-Supported_metrics) metrics also available for scraping.

ExternalDNS exposes 3 types of metrics: Sources, Registry errors and Cache hits.

`Source`s are mostly Kubernetes API objects. Examples of `source` errors may be connection errors to the Kubernetes API server itself or missing RBAC permissions.
It can also stem from incompatible configuration in the objects itself like invalid characters, processing a broken fqdnTemplate, etc.

`Registry` errors are mostly Provider errors, unless there's some coding flaw in the registry package. Provider errors often arise due to accessing their APIs due to network or missing cloud-provider permissions when reading records.
When applying a changeset, errors will arise if the changeset applied is incompatible with the current state.

In case of an increased error count, you could correlate them with the `http_request_duration_seconds{handler="instrumented_http"}` metric which should show increased numbers for status codes 4xx (permissions, configuration, invalid changeset) or 5xx (apiserver down).

You can use the host label in the metric to figure out if the request was against the Kubernetes API server (Source errors) or the DNS provider API (Registry/Provider errors).

## Owner Mismatch Metrics

The `external_dns_registry_skipped_records_owner_mismatch_per_sync` metric tracks DNS records that were skipped during synchronization because they are owned by a different ExternalDNS instance. This is useful for detecting ownership conflicts in multi-tenant or multi-instance deployments.

The metric includes the following labels:

| Label | Description |
|:------|:------------|
| `record_type` | DNS record type (A, AAAA, CNAME, etc.) |
| `owner` | The owner ID of the current ExternalDNS instance |
| `foreign_owner` | The owner ID found on the existing record |
| `domain` | The naked/apex domain (e.g., "example.com") |

**Note:** The `domain` label uses the naked/apex domain rather than the full FQDN to prevent metric cardinality explosion. With thousands of subdomains under one apex domain, using full FQDNs would create excessive metric series.

## Metrics Best Practices

When scraping ExternalDNS metrics, consider the following best practices:

### Cardinality Management

- **Vector metrics** (those with labels like `record_type`, `domain`) can generate multiple time series. Monitor your Prometheus storage and memory usage accordingly.
- The `domain` label on owner mismatch metrics is intentionally limited to apex domains to bound cardinality.
- Use recording rules to pre-aggregate high-cardinality metrics if you only need totals.

### Recommended Scrape Interval

- A scrape interval of 10-30 seconds is typically sufficient for ExternalDNS metrics.
- Align your scrape interval with ExternalDNS's sync interval (`--interval` flag) for meaningful data.

### Alerting Recommendations

Consider alerting on:

- `external_dns_source_errors_total` or `external_dns_registry_errors_total` increasing - indicates connectivity or permission issues.
- `external_dns_controller_last_sync_timestamp_seconds` not updating - indicates the sync loop may be stuck.
- `external_dns_registry_skipped_records_owner_mismatch_per_sync` non-zero - indicates ownership conflicts that may need investigation.

## Resources

- [Prometheus Instrumentation](https://prometheus.io/docs/practices/instrumentation/)
- [Prometheus Alerting Best Practices](https://prometheus.io/docs/practices/alerting/)
- [Prometheus Recording Rules](https://prometheus.io/docs/practices/rules/)
- [Grafana: How to Manage High Cardinality Metrics](https://grafana.com/blog/2022/02/15/what-are-cardinality-spikes-and-why-do-they-matter/)
