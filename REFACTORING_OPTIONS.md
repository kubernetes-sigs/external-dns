# GetRestConfig Refactoring Options

## Current Implementation

The function in `pkg/client/config.go` has this priority:
1. Use explicit `kubeConfig` path if provided
2. Fall back to `~/.kube/config` if it exists
3. Use in-cluster config as last resort

## Refactoring Options

### Option A: Use `clientcmd.NewNonInteractiveDeferredLoadingClientConfig` (Recommended)

This is the idiomatic client-go approach that handles all config loading scenarios:

```go
func GetRestConfig(kubeConfig, apiServerURL string) (*rest.Config, error) {
    loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
    if kubeConfig != "" {
        loadingRules.ExplicitPath = kubeConfig
    }

    configOverrides := &clientcmd.ConfigOverrides{}
    if apiServerURL != "" {
        configOverrides.ClusterInfo.Server = apiServerURL
    }

    return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).ClientConfig()
}
```

**Pros:**
- Respects `KUBECONFIG` environment variable automatically
- Handles multiple kubeconfig files (colon-separated in `KUBECONFIG`)
- Standard pattern used throughout Kubernetes ecosystem
- Less manual file checking
- Eliminates the manual `os.Stat(clientcmd.RecommendedHomeFile)` check

**Cons:**
- Slightly different behavior (respects `KUBECONFIG` env var which current code ignores)

**Test compatibility:** Existing tests that modify `clientcmd.RecommendedHomeFile` still work, but may need to also clear `KUBECONFIG` env var:

```go
t.Cleanup(func() {
    clientcmd.RecommendedHomeFile = prevRecommendedHomeFile
})
t.Setenv("KUBECONFIG", "")  // Ensure env var doesn't override
clientcmd.RecommendedHomeFile = mockKubeCfgPath
```

### Option B: Remove auto-detection entirely

Just don't auto-detect `~/.kube/config`:

```go
func GetRestConfig(kubeConfig, apiServerURL string) (*rest.Config, error) {
    return clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
}
```

**Pros:**
- Simplest implementation
- No magic file detection
- Cleaner for production services running in-cluster

**Cons:**
- Users running locally must explicitly pass `--kubeconfig ~/.kube/config`
- Breaking change in behavior

**Test compatibility:** Tests for `RecommendedHomeFile` fallback become dead code and should be removed.

## Recommendation

**Option A** is recommended because:
1. It follows Kubernetes conventions
2. It's more robust and handles edge cases
3. It maintains backward compatibility with existing behavior
4. Existing tests can be adapted with minimal changes
