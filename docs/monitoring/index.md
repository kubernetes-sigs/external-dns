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
