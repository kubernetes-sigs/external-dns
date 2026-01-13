# Informer Sync Timeout Configuration

## Introduction

ExternalDNS interacts with the Kubernetes API server to discover resources (Services, Ingresses, Nodes, etc.)
that require DNS records. During startup and reconciliation, ExternalDNS uses informers to cache Kubernetes resources locally.
The `--informer-sync-timeout` flag allows you to configure the timeout duration for these informer cache synchronization operations.

!!! warning "Investigate Root Causes First"
    Timeout errors during informer sync are typically **symptoms** of underlying issues, not the root cause.
    Before increasing this timeout, investigate:
    
    1. **RBAC Permissions**: Verify the ServiceAccount has proper permissions for all watched resources
    2. **API Server Health**: Check if the Kubernetes API server is overloaded or slow
    3. **Network Issues**: Ensure network connectivity between ExternalDNS and the API server
    4. **Resource Contention**: Check for interference from other controllers (e.g., Gateway API controllers)

## Soft Error Handling

As of this version, ExternalDNS uses **soft error handling** for informer sync failures. Instead of crashing
when cache sync fails or times out, ExternalDNS will:

1. Log a warning with diagnostic information
2. Continue operating with potentially stale or incomplete cache data
3. Retry on the next reconciliation cycle

This approach prevents crash loops that can overwhelm the Kubernetes API server with repeated LIST/WATCH calls,
which can worsen the underlying problem.

## Why Configure Informer Sync Timeout?

In most cases, the default timeout is sufficient. However, you may need to adjust it if:

- You have verified RBAC permissions are correct
- API server health and network connectivity are confirmed
- Your cluster has genuinely slow response times due to size or complexity

Cache synchronization timeout errors look like:

```text
level=warning msg="Cache sync for *v1.Service timed out after 60s: context deadline exceeded. This may indicate RBAC issues, API server latency, or network problems. Continuing with potentially stale data."
```

## Usage

```bash
# New recommended flag
external-dns --informer-sync-timeout=120s

# Deprecated flag (still works for backward compatibility)
external-dns --request-timeout=120s
```

The default value is `60s`. Setting the value to `0s` uses the default timeout.

!!! note "Flag Deprecation"
    The `--request-timeout` flag is deprecated. Use `--informer-sync-timeout` instead, as it more accurately
    describes what this setting controls (informer cache sync, not individual HTTP requests).

## Affected Operations

The `--informer-sync-timeout` setting affects:

1. **Cache Synchronization**: The timeout for waiting for informer caches to sync during startup. This is the primary use case for this flag.

2. **Source Initialization**: All source types (Service, Ingress, Node, Pod, Gateway API, Istio, CRD, etc.) respect this timeout during their initialization phase.

Note: This setting does **not** affect individual Kubernetes API HTTP request timeouts. It only controls
how long ExternalDNS waits for the informer cache to be fully populated.

## Example Scenarios

### After Verifying No RBAC Issues

If you've confirmed permissions are correct but still see timeouts in a large cluster:

```bash
external-dns \
  --source=service \
  --source=ingress \
  --informer-sync-timeout=120s \
  --provider=aws
```

### Slow Network Environments

For environments with confirmed high latency to the API server:

```bash
external-dns \
  --source=service \
  --informer-sync-timeout=90s \
  --provider=cloudflare
```

## Troubleshooting

If you encounter timeout warnings during startup:

1. **Check RBAC Permissions First**: This is the most common cause. Verify that the ExternalDNS ServiceAccount has the necessary permissions:
   ```yaml
   rules:
   - apiGroups: [""]
     resources: ["services", "endpoints", "pods", "nodes"]
     verbs: ["get", "watch", "list"]
   - apiGroups: ["discovery.k8s.io"]
     resources: ["endpointslices"]
     verbs: ["get", "watch", "list"]
   - apiGroups: ["extensions", "networking.k8s.io"]
     resources: ["ingresses"]
     verbs: ["get", "watch", "list"]
   ```

2. **Check API Server Health**: Ensure the Kubernetes API server is responding normally:
   ```bash
   kubectl get --raw /healthz
   ```

3. **Monitor API Server Metrics**: Look for signs of throttling or overload.

4. **Review ExternalDNS Logs**: Look for permission denied (403) errors that might indicate RBAC issues.

5. **Consider Scope Reduction**: Use `--namespace` or label filters to reduce the number of resources watched.

## Related Options

- `--interval`: The interval between synchronization cycles (default: 1m)
- `--min-event-sync-interval`: Minimum interval between event-triggered syncs (default: 5s)
- `--provider-cache-time`: Cache duration for DNS provider records (default: 0s, disabled)
