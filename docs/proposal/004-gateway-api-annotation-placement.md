```yaml
---
title: "Gateway API Annotation Placement Clarity"
version: v1alpha1
authors: "@lexfrei"
creation-date: 2025-10-23
status: provisional
---
```

# Gateway API Annotation Placement Clarity

## Table of Contents

<!-- toc -->
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [User Stories](#user-stories)
  - [Current Behavior](#current-behavior)
  - [Proposed Solutions](#proposed-solutions)
  - [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
<!-- /toc -->

## Summary

The [annotations documentation](https://kubernetes-sigs.github.io/external-dns/latest/docs/annotations/annotations/)
indicates that Gateway API sources support various annotations, but it does not clearly specify which Kubernetes
resource (Gateway vs HTTPRoute/GRPCRoute/TLSRoute/etc.) these annotations should be placed on. This ambiguity leads
to user confusion and misconfigurations.

This proposal aims to:

1. **Short-term**: Improve documentation to explicitly clarify annotation placement
2. **Long-term**: Consider implementing annotation inheritance from Gateway to Routes

## Motivation

Users frequently misconfigure annotations when using Gateway API sources because the current documentation uses "Gateway" as the source name in the annotation support table, which is ambiguousit refers to gateway-api sources generically, not the Gateway resource specifically.

### Current Implementation Behavior

Based on the source code
([source/gateway.go](https://github.com/kubernetes-sigs/external-dns/blob/master/source/gateway.go)):

**Gateway resource annotations:**

- `external-dns.alpha.kubernetes.io/target` - read from Gateway
  ([line ~380](https://github.com/kubernetes-sigs/external-dns/blob/master/source/gateway.go#L380))

**Route resource annotations (HTTPRoute, GRPCRoute, TLSRoute, TCPRoute, UDPRoute):**

- `external-dns.alpha.kubernetes.io/hostname` - read from Route
- `external-dns.alpha.kubernetes.io/ttl` - read from Route
- `external-dns.alpha.kubernetes.io/controller` - read from Route
- **Provider-specific annotations** (e.g., `cloudflare-proxied`, `aws/*`, `scw/*`, etc.) - read from Route
  ([line ~242](https://github.com/kubernetes-sigs/external-dns/blob/master/source/gateway.go#L242))

This separation aligns with Gateway API architecture:

- **Gateway** = infrastructure layer (IP addresses, listeners, load balancers)
- **Routes** = application layer (DNS records, routing rules, hostnames)

However, users expect provider-specific annotations to work on Gateway (similar to how `target` works), leading to silent failures.

### Goals

- Clarify annotation placement in documentation to prevent user confusion
- Provide practical examples for common providers (Cloudflare, AWS, Scaleway)
- Define a clear, documented contract for where each annotation type should be placed
- Reduce support burden from repeated misconfigurations

### Non-Goals

- This proposal does not address the broader annotation standardization effort discussed in [PR #5080](https://github.com/kubernetes-sigs/external-dns/pull/5080)
- Redesigning the Gateway API source implementation
- Changing behavior for non-Gateway sources (Ingress, Service, etc.)
- Making breaking changes to existing Gateway API functionality

## Proposal

### User Stories

#### Story 1: Platform Engineer with Cloudflare (#5901)

*As a platform engineer*, I set up a Gateway with the `external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"`
annotation, expecting all DNS records for Routes using this Gateway to be proxied through Cloudflare. However, the
annotation is silently ignored, and records are created without proxy, leading to unexpected traffic routing and
security issues.

**Root cause**: Had to dive into source code to discover that provider-specific annotations are only read from
Route resources, not Gateway resources.

**Current workaround**: Must copy the `cloudflare-proxied` annotation to every HTTPRoute manually.

#### Story 2: User Attempting Route-Specific Targets (#4056)

*As a user*, I want to specify different target DNS records for specific hosts while sharing a common Gateway. I
added `external-dns.alpha.kubernetes.io/target` annotation on HTTPRoute to override the Gateway's target for one
specific host, but it doesn't work - the annotation is ignored on HTTPRoute.

**Root cause**: The `target` annotation must be on the Gateway resource, not on Route resources. There's no way to
override targets on a per-Route basis.

**Outcome**: User had to find alternative workarounds to exclude specific hosts or create separate Gateway resources.

### Current Behavior

#### Annotation Placement Matrix

| Annotation Type                                            | Gateway Resource        | Route Resources (HTTPRoute, GRPCRoute, etc.) |
|------------------------------------------------------------|-------------------------|----------------------------------------------|
| `target`                                                   |  **Read from Gateway**  | L Ignored                                    |
| `hostname`                                                 | L Not used              |  **Read from Route**                         |
| `ttl`                                                      | L Not used              |  **Read from Route**                         |
| `controller`                                               | L Not used              |  **Read from Route**                         |
| Provider-specific (`cloudflare-proxied`, `aws/*`, `scw/*`) | L Not used              |  **Read from Route**                         |

#### Code References

```go
// source/gateway.go line ~380
// Target annotation is read from Gateway
override := annotations.TargetsFromTargetAnnotation(gw.gateway.Annotations)

// source/gateway.go line ~242
// Provider-specific annotations are read from Route
providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(annots)
```

Where `annots` is derived from the Route's metadata (`meta.Annotations`), not the Gateway.

### Proposed Solutions

#### Solution 1: Documentation Improvements (Short-term - Quick Win)

**Implementation Status**: Documentation improvements proposed in [PR #5918](https://github.com/kubernetes-sigs/external-dns/pull/5918).

**Note**: If Solution 2 (Annotation Merging) is implemented, the documentation from PR #5918 will require updates
to reflect the new inheritance behavior.

**Changes to `docs/annotations/annotations.md`:**

Expand footnote [^4] or add a new section "Gateway API Annotation Placement" with a detailed table:

```markdown
### Gateway API Annotation Placement

When using Gateway API sources (gateway-httproute, gateway-grpcroute, etc.), annotations must be placed on specific resources:

| Annotation | Placement | Example Resource |
|------------|-----------|------------------|
| `target` | Gateway | `kind: Gateway` |
| `hostname` | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |
| `ttl` | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |
| `controller` | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |
| `cloudflare-proxied` | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |
| `aws-*` (all AWS annotations) | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |
| `scw-*` (all Scaleway annotations) | Route | `kind: HTTPRoute`, `kind: GRPCRoute`, etc. |

**Rationale**: The Gateway resource defines infrastructure (IP addresses, listeners), while Routes define application-level DNS records. Therefore, DNS record properties (TTL, provider settings) are configured on Routes.
```

**Changes to `docs/sources/gateway-api.md`:**

Add a new section after "Hostnames":

```markdown
## Annotations

### Annotation Placement

ExternalDNS reads different annotations from different Gateway API resources:

- **Gateway annotations**: Only `external-dns.alpha.kubernetes.io/target` is read from Gateway resources
- **Route annotations**: All other annotations (hostname, ttl, provider-specific) are read from Route resources

#### Example: Cloudflare Proxied Records

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
  namespace: default
  annotations:
    #  Correct: target annotation on Gateway
    external-dns.alpha.kubernetes.io/target: "203.0.113.1"
spec:
  gatewayClassName: cilium
  listeners:
    - name: https
      hostname: "*.example.com"
      protocol: HTTPS
      port: 443
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: my-route
  annotations:
    #  Correct: provider-specific annotations on HTTPRoute
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
    external-dns.alpha.kubernetes.io/ttl: "300"
spec:
  parentRefs:
    - name: my-gateway
      namespace: default
  hostnames:
    - api.example.com
  rules:
    - backendRefs:
        - name: api-service
          port: 8080
```

#### Example: AWS Route53 with Routing Policies

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: aws-gateway
  annotations:
    #  Correct: target annotation on Gateway
    external-dns.alpha.kubernetes.io/target: "alb-123.us-east-1.elb.amazonaws.com"
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: weighted-route
  annotations:
    #  Correct: AWS-specific annotations on HTTPRoute
    external-dns.alpha.kubernetes.io/aws-weight: "100"
    external-dns.alpha.kubernetes.io/set-identifier: "backend-v1"
spec:
  parentRefs:
    - name: aws-gateway
  hostnames:
    - app.example.com
```

### Common Mistakes

❌ **Incorrect**: Placing provider-specific annotations on Gateway

```yaml
kind: Gateway
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"  # ❌ Ignored
```

❌ **Incorrect**: Placing target annotation on HTTPRoute

```yaml
kind: HTTPRoute
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target: "203.0.113.1"  # ❌ Ignored
```

**Implementation effort**: Low
**Maintenance burden**: Minimal (documentation only)
**User benefit**: Immediate clarity, reduced misconfiguration

#### Solution 2: Annotation Inheritance and Merging (Long-term - Feature Enhancement)

**Reference Implementation**: [PR #5998](https://github.com/kubernetes-sigs/external-dns/pull/5998)

Implement annotation merging logic where:

1. Gateway annotations serve as **defaults** for all Routes attached to that Gateway
2. Route annotations **override** Gateway annotations for specific Routes
3. **All annotations are inheritable**, including `target` — enabling per-Route target overrides

**Proposed implementation** (pseudocode):

```go
// source/gateway.go - proposed changes
func (src *gatewayRouteSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
    // ... existing code ...

    for _, route := range routes {
        // Merge Gateway and Route annotations
        // Route annotations take precedence over Gateway annotations
        gwAnnots := gw.gateway.Annotations
        rtAnnots := route.meta.Annotations
        mergedAnnots := mergeAnnotations(gwAnnots, rtAnnots)

        // Use merged annotations for all annotation processing
        providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(mergedAnnots)
        ttl := annotations.TTLFromAnnotations(mergedAnnots, resource)

        // ... rest of endpoint creation ...
    }
}

// Helper function
func mergeAnnotations(gateway, route map[string]string) map[string]string {
    merged := make(map[string]string, len(gateway)+len(route))

    // Copy Gateway annotations (defaults)
    for k, v := range gateway {
        merged[k] = v
    }

    // Route annotations override Gateway defaults
    for k, v := range route {
        merged[k] = v
    }

    return merged
}
```

**Example use case enabled by this approach:**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: intranet-gateway
  annotations:
    # Default target for internal services
    external-dns.alpha.kubernetes.io/target: "172.16.6.6"
    # Set default for all Routes using this Gateway
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
    external-dns.alpha.kubernetes.io/ttl: "300"
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: internal-api
  # Inherits: target=172.16.6.6, cloudflare-proxied=true, ttl=300 from Gateway
spec:
  parentRefs:
    - name: intranet-gateway
  hostnames:
    - api.internal.example.com
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: public-api
  annotations:
    # Override: expose this route to the public internet
    external-dns.alpha.kubernetes.io/target: "203.0.113.1"
    # Inherits: cloudflare-proxied=true, ttl=300 from Gateway
spec:
  parentRefs:
    - name: intranet-gateway
  hostnames:
    - api.example.com
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: static-assets
  annotations:
    # Override: disable proxying for static content
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "false"
    # Inherits: target=172.16.6.6, ttl=300 from Gateway
spec:
  parentRefs:
    - name: intranet-gateway
  hostnames:
    - static.internal.example.com
```

This example demonstrates a common use case: an intranet Gateway where most services are internal
(`172.16.6.6`), but specific Routes can be exposed publicly (`203.0.113.1`) by overriding the
`target` annotation.

**Benefits:**

- Reduces configuration duplication
- Enables centralized defaults at Gateway level
- Maintains flexibility with Route-level overrides
- Better matches user mental model (infrastructure defaults + application overrides)
- **Solves User Story 2**: Enables per-Route target overrides without creating separate Gateways

**Risks:**

- Backward compatibility concerns (may change behavior for existing users)
- Increased code complexity
- Potential for confusion about precedence rules
- Need for comprehensive testing across all Gateway API Route types

**Mitigation strategies:**

- Feature flag to opt-in to new behavior initially
- Clear documentation of precedence rules
- Extensive test coverage
- Migration guide for users

**Implementation effort**: Medium
**Maintenance burden**: Medium (code + tests + docs)
**User benefit**: Significant reduction in configuration overhead

### Drawbacks

#### Documentation-Only Solution

- Does not address the underlying UX issue (annotation duplication)
- Requires users to manually propagate settings across Routes
- Still allows silent failures if users misplace annotations

#### Annotation Merging Solution

- Adds complexity to the codebase
- Requires careful consideration of precedence rules
- May introduce unexpected behavior changes for existing users
- Needs comprehensive testing for edge cases (multiple Gateways, cross-namespace, etc.)
- Potential performance impact from annotation merging on every reconciliation

## Alternatives

### Alternative 1: Do Nothing (Status Quo)

**Description**: Keep current behavior and documentation as-is.

**Pros**:

- No implementation effort required
- No risk of introducing new bugs
- No breaking changes

**Cons**:

- Users continue to experience confusion and misconfigurations
- Increased support burden on maintainers and community
- Poor user experience compared to other sources (Ingress supports annotations more intuitively)

**Recommendation**: L Not recommended - problem is well-documented and affects user productivity

### Alternative 2: Move All Annotations to Gateway Only

**Description**: Refactor source code to read all annotations from Gateway, not Routes.

**Pros**:

- Simplified mental model (one place for all annotations)
- Centralized configuration

**Cons**:

- **Breaks Gateway API architecture** - Routes define application-layer DNS records, so DNS properties belong on Routes
- Cannot have different settings per Route (e.g., different TTLs for api.example.com vs static.example.com)
- Loses flexibility that Route-level annotations provide
- Requires breaking change to existing implementations

**Recommendation**: L Not recommended - violates Gateway API design principles

### Alternative 3: Support Annotations on Both with Strict Validation

**Description**: Allow annotations on both Gateway and Route, but error/warn if duplicates exist without clear precedence.

**Pros**:

- Provides flexibility
- Catches configuration errors explicitly

**Cons**:

- Confusing for users (two valid places to configure)
- Requires complex validation logic
- Still doesn't solve the "defaults + overrides" use case
- More complex to document and support

**Recommendation**: � Possible but adds complexity without solving core UX issue

### Alternative 4: Create Dedicated GatewayDNSConfig CRD

**Description**: Introduce a new CRD that defines DNS configuration separately from Gateway and Route resources.

**Example**:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: GatewayDNSConfig
metadata:
  name: cloudflare-defaults
spec:
  gatewayRef:
    name: my-gateway
  defaults:
    ttl: 300
    providerSpecific:
      - name: cloudflare-proxied
        value: "true"
---
apiVersion: externaldns.k8s.io/v1alpha1
kind: RouteDNSConfig
metadata:
  name: api-route-dns
spec:
  routeRef:
    kind: HTTPRoute
    name: api-route
  overrides:
    ttl: 60  # Override Gateway default
```

**Pros**:

- Clean separation of concerns
- Clear precedence model
- No annotations needed (type-safe CRDs)
- Aligns with Kubernetes resource composition patterns

**Cons**:

- **Significant implementation effort** (new CRDs, controllers, validation, etc.)
- Adds complexity with additional resources to manage
- Requires migration from annotation-based approach
- Diverges from how other sources work (Ingress, Service use annotations)
- May conflict with future annotation standardization efforts

**Recommendation**: � Potentially valuable long-term, but scope is too large for this specific issue

### Alternative 5: Wait for Annotation Standardization (PR #5080)

**Description**: Defer this work until the broader annotation standardization effort is resolved.

**Pros**:

- Avoids potentially redundant work
- May be addressed as part of larger effort

**Cons**:

- PR #5080 is not yet ready for review and timeline is uncertain
- Users continue to experience issues in the meantime
- Documentation improvements are still valuable regardless of standardization outcome

**Recommendation**: � Partial - implement documentation improvements now (Solution 1), reconsider
annotation merging after standardization is resolved

## Recommendation

**Phased approach**:

1. **Immediate (v0.15.0 or next minor)**: Implement Solution 1 (Documentation Improvements)
   - Low risk, high user value
   - Can be merged quickly
   - Addresses immediate pain points

2. **Near-term**: Review and merge Solution 2 (Annotation Merging)
   - Reference implementation available: [PR #5998](https://github.com/kubernetes-sigs/external-dns/pull/5998)
   - Includes comprehensive test coverage
   - Backward compatible (no breaking changes for existing configurations)
   - Solves User Story 2 (per-Route target overrides)

3. **Future (post-PR #5080 resolution)**: Re-evaluate if additional changes are needed
   - Assess compatibility with annotation standardization outcomes
   - Gather user feedback on the annotation inheritance behavior

This approach provides immediate relief while keeping options open for more comprehensive solutions in the future.
