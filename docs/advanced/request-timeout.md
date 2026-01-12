# Request Timeout Configuration

## Introduction

ExternalDNS interacts with the Kubernetes API server to discover resources (Services, Ingresses, Nodes, etc.)
that require DNS records. During startup and reconciliation, ExternalDNS uses informers to cache Kubernetes resources locally.
The `--request-timeout` flag allows you to configure the timeout duration for these Kubernetes API operations.

## Why Configure Request Timeout?

In large Kubernetes clusters with many resources, the initial cache synchronization may take longer than the default timeout. This can result in errors like:

```text
failed to sync <resource>: context deadline exceeded with timeout 60s
```

By configuring a longer request timeout, you can ensure that ExternalDNS has sufficient time to synchronize its cache, especially in clusters with:

- Large numbers of Services, Ingresses, or other watched resources
- High API server latency
- Network conditions that may slow API responses

## Usage

```bash
external-dns --request-timeout=120s
```

The default value is `30s`. Setting the value to `0s` means no timeout (the operation will wait indefinitely).

## Affected Operations

The `--request-timeout` setting affects:

1. **Kubernetes Client Operations**: API calls made to the Kubernetes API server when creating clients for various sources.

2. **Cache Synchronization**: The timeout for waiting for informer caches to sync during startup. This is the primary use case for this flag.

3. **Source Initialization**: All source types (Service, Ingress, Node, Pod, Gateway API, Istio, CRD, etc.) respect this timeout during their initialization phase.

## Example Scenarios

### Large Clusters

For clusters with thousands of resources:

```bash
external-dns \
  --source=service \
  --source=ingress \
  --request-timeout=120s \
  --provider=aws
```

### Slow Network Environments

For environments with high latency to the API server:

```bash
external-dns \
  --source=service \
  --request-timeout=90s \
  --provider=cloudflare
```

### Development and Debugging

To wait indefinitely (useful for debugging):

```bash
external-dns \
  --source=service \
  --request-timeout=0s \
  --provider=inmemory
```

!!! warning
    Setting `--request-timeout=0s` disables the timeout entirely. Use this cautiously in production environments as it may cause ExternalDNS to hang indefinitely if there are connectivity issues.

## Default Behavior

When `--request-timeout` is not specified or set to `0s`, ExternalDNS uses a default timeout of 60 seconds for cache synchronization operations. This ensures backward compatibility while providing a reasonable default for most environments.

## Troubleshooting

If you encounter timeout errors during startup:

1. **Check API Server Health**: Ensure the Kubernetes API server is responding normally.

2. **Verify RBAC Permissions**: Timeout errors can sometimes mask underlying permission issues. Verify that the ExternalDNS service account has the necessary permissions.

3. **Monitor Resource Count**: Large numbers of resources can increase sync time. Consider using `--namespace` or label filters to reduce the scope.

4. **Increase Timeout Gradually**: Start by doubling the timeout value and adjust based on observed sync times.

## Related Options

- `--interval`: The interval between synchronization cycles (default: 1m)
- `--min-event-sync-interval`: Minimum interval between event-triggered syncs (default: 5s)
- `--provider-cache-time`: Cache duration for DNS provider records (default: 0s, disabled)
