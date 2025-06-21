# Available Metrics

<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW METRIC ADDED PLEASE RUN 'make generate-metrics-documentation' -->
<!-- markdownlint-disable MD013 -->

All metrics available for scraping are exposed on the `/metrics` endpoint.
The metrics are in the Prometheus exposition format.

To access the metrics:

```sh
curl https://localhost:7979/metrics
```

## Supported Metrics

> Full metric name is constructed as follows:
> `external_dns_<subsystem>_<name>`

| Name                             | Metric Type | Subsystem   |  Help                                                 |
|:---------------------------------|:------------|:------------|:------------------------------------------------------|
| consecutive_soft_errors | Gauge | controller | Number of consecutive soft errors in reconciliation loop. |
| last_reconcile_timestamp_seconds | Gauge | controller | Timestamp of last attempted sync with the DNS provider |
| last_sync_timestamp_seconds | Gauge | controller | Timestamp of last successful sync with the DNS provider |
| no_op_runs_total | Counter | controller | Number of reconcile loops ending up with no changes on the DNS provider side. |
| verified_records | Gauge | controller | Number of DNS records that exists both in source and registry (vector). |
| cache_apply_changes_calls | Counter | provider | Number of calls to the provider cache ApplyChanges. |
| cache_records_calls | Counter | provider | Number of calls to the provider cache Records list. |
| endpoints_total | Gauge | registry | Number of Endpoints in the registry |
| errors_total | Counter | registry | Number of Registry errors. |
| records | Gauge | registry | Number of registry records partitioned by label name (vector). |
| endpoints_total | Gauge | source | Number of Endpoints in all sources |
| errors_total | Counter | source | Number of Source errors. |
| records | Gauge | source | Number of source records partitioned by label name (vector). |
| adjustendpoints_errors_total | Gauge | webhook_provider | Errors with AdjustEndpoints method |
| adjustendpoints_requests_total | Gauge | webhook_provider | Requests with AdjustEndpoints method |
| applychanges_errors_total | Gauge | webhook_provider | Errors with ApplyChanges method |
| applychanges_requests_total | Gauge | webhook_provider | Requests with ApplyChanges method |
| records_errors_total | Gauge | webhook_provider | Errors with Records method |
| records_requests_total | Gauge | webhook_provider | Requests with Records method |

## Available Go Runtime Metrics

> The following Go runtime metrics are available for scraping. Please note that they may change over time and they are OS dependent.

| Name                  |
|:----------------------|
| go_gc_duration_seconds |
| go_gc_gogc_percent |
| go_gc_gomemlimit_bytes |
| go_goroutines |
| go_info |
| go_memstats_alloc_bytes |
| go_memstats_alloc_bytes_total |
| go_memstats_buck_hash_sys_bytes |
| go_memstats_frees_total |
| go_memstats_gc_sys_bytes |
| go_memstats_heap_alloc_bytes |
| go_memstats_heap_idle_bytes |
| go_memstats_heap_inuse_bytes |
| go_memstats_heap_objects |
| go_memstats_heap_released_bytes |
| go_memstats_heap_sys_bytes |
| go_memstats_last_gc_time_seconds |
| go_memstats_mallocs_total |
| go_memstats_mcache_inuse_bytes |
| go_memstats_mcache_sys_bytes |
| go_memstats_mspan_inuse_bytes |
| go_memstats_mspan_sys_bytes |
| go_memstats_next_gc_bytes |
| go_memstats_other_sys_bytes |
| go_memstats_stack_inuse_bytes |
| go_memstats_stack_sys_bytes |
| go_memstats_sys_bytes |
| go_sched_gomaxprocs_threads |
| go_threads |
| http_request_duration_seconds |
| process_cpu_seconds_total |
| process_max_fds |
| process_network_receive_bytes_total |
| process_network_transmit_bytes_total |
| process_open_fds |
| process_resident_memory_bytes |
| process_start_time_seconds |
| process_virtual_memory_bytes |
| process_virtual_memory_max_bytes |
