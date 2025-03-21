```yaml
---
title: "Proposal: Annotations graduation to beta"
version: tbd
authors: ivankatliarchuk
creation-date: 2025-feb-9
status: proposal
---
```

# Proposal: Agree on requirements for annotations to graduate to beta

This proposal suggests restructuring the annotation processing mechanism in `external-dns` by adopting a design similar to `ingress-nginx`. The goal is to create a more maintainable, scalable, and user-friendly system for handling annotations, ensuring clear support for different API versions and facilitating automated documentation.

Upon reviewing the current open issues related to annotations in the `external-dns` project, several key challenges have been identified. To address these challenges, the following refinements to the annotation processing proposal are suggested:

- **Deprecation Policy and Migration Path**: Establish a clear deprecation policy for outdated annotations. Implement mechanisms to log warnings when deprecated annotations are used and provide comprehensive migration guides to assist users in transitioning to supported annotations.

- **Conflict Detection and Resolution**: Enhance the annotation processing logic to detect conflicting annotations proactively. Implement validation rules that either prevent conflicts at the time of deployment or resolve them in a predictable manner, ensuring consistent behavior.

**Motivation**

Annotations are a critical mechanism for configuring `external-dns` behavior at the resource level. However, the existing system faces several challenges:

- No automated documentation for available annotations.
- Unclear strategy for supporting different API versions.
- No defined transition path from `external-dns.alpha.kubernetes.io` annotations to a stable format.
- Lack of standardization among annotations.

By adopting a structured approach similar to `ingress-nginx`, we can address these issues and improve the overall functionality and user experience of `external-dns`.

### Goals

- Introduce automated documentation for supported annotations.
- Define a strategy for handling multiple API versions in annotations.
- Ensure backward compatibility where possible.

### Non-Goals

- Establish a migration plan from `external-dns.alpha.kubernetes.io` to `external-dns.beta.kubernetes.io`.
- Deprecating annotations in favor of an alternative configuration method.
- Making `external-dns` rely solely on CRDs for configuration.
- Redesigning the entire `external-dns` architecture.
- Introducing breaking changes that would require significant refactoring for existing users.

## Proposal

1. Structured Annotation Definitions
2. Automated Documentation
3. Versioned Annotations Support
4. Migration Plan
5. TBD
6. TBD

**Structured Annotation Definitions**

    -   Create a dedicated package (e.g., `pkg/annotations`) to house all annotation-related code.
    -   Define each annotation as a separate struct with fields for its name, description, default value, and validation logic.
    -   Implement a central registry to manage all available annotations.

**Automated Documentation**
	-   Introduce a mechanism to automatically generate and publish documentation for annotations.
	-   Provide examples and best practices for usage.
	-   Integrate this annotation docs generation into the build pipeline to ensure documentation is updated with each build and release.

**Versioned Annotations Support**

   - Introduce versioning in annotation definitions to distinguish between `alpha`, `beta`, and stable annotations.
   - Allow processing of both `alpha` and `beta` annotations in parallel.
   - Introduce a mechanism for logging warnings when using deprecated annotations.

 **Migration Plan**
    - Introduce feature flags for annotations to enable strict mode (blocking alpha annotations after a certain point).

A somewhere similar example from [ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/annotations)

```go
var ttlAnnotationGroup = parser.Annotation{
	Group: "conroller|owner|provider-specific|rate-limit|hostname|internal-hostname|target|ttl"
	Annotations: parser.AnnotationFields{
		"external-dns.alpha.kubernetes.io/ttl": {
			Validators:         [parser.ValidateDuration, parser.ValidateRegex, ...],
			Sources:            ["ambassador", "istio", "traefik", "fake", ...],
			Providers:          ["aws", "ibm", "cloudflare", ...],
			ApiVersions:        ["alpha", "beta", ...], // alternative versions
			DeprecatedVersions: true,
			Documentation: `Specifies the TTL (time to live) for the resource's DNS records. The value may be specified as either a duration or an integer number of seconds. It must be between 1 and 2,147,483,647 seconds.`
			UsageExample:  `
				kind: DNSEndpoint
				metadata:
					annotations:
					  external-dns.alpha.kubernetes.io/ttl: "10s"
			    ---
				kind: Service
				metadata:
					annotations:
					  external-dns.alpha.kubernetes.io/ttl: "10s"
			`
		},
		"external-dns.beta.kubernetes.io/ttl": {
			...
		}
	}
}
```

Or example simplified version of annotation processing [project countour](https://github.com/projectcontour/contour/blob/23c1779d25b4737c1d470677d581bb74e310145d/internal/annotation/annotations.go)

### User Stories

1.  **Cluster Administrator**

    -   Wants clear and up-to-date documentation on available annotations to configure `external-dns` effectively.

2.  **Developer**

    -   Seeks a structured and predictable annotation system that simplifies the development and deployment process.
	-   _As a developer_, I want confidence that my annotations will remain supported or have a clear migration path, so I don't introduce breaking changes in my deployments.

3.  **Operator**

    -   Requires a seamless migration path from `alpha` to `beta` annotations to ensure uninterrupted service.

4. **WebHook Maintainer**

	- _As a webhooks maintainer_, I want to implement validation webhooks that check for incorrect or deprecated External-DNS annotations when users create resources
	- _As a webhooks maintainer_, I want to reject resource definitions where multiple conflicting DNS annotations are present,
	- _As a webhooks maintainer_,I want to allow cluster administrators to enable or disable strict annotation enforcement through feature flags,
so that they can gradually adopt new annotation standards without causing sudden disruptions.

5.  **Contributor**

    - _As a contributor_, I want to refactor the annotation processing code to follow a structured and maintainable approach,
so that adding new annotations or deprecating old ones becomes easier and reduces technical debt.
    - _As a contributor_, I want to write unit and integration tests for different annotation use cases, so that annotation-related bugs and regressions are caught early before reaching production users.
    - _As a contributor_, I want to ensure that new annotations can coexist with old ones (alpha), so that users are not forced into immediate migrations and can transition smoothly.
    - _As a contributor_, I want to create a tool that extracts annotation definitions from the source code and generates up-to-date documentation, so that users always have access to accurate information on supported annotations.

5.  **Maintainer**
    - _As a maintainer_, I want to define and communicate a clear lifecycle for annotation versions, so that contributors and users understand when alpha annotations will be deprecated and how to migrate.
	- _As a maintainer_, I want to ensure that annotation behavior is consistent across supported DNS providers (e.g., AWS Route 53, Cloudflare, oogle DNS), so that users do not experience unexpected inconsistencies depending on their provider.
	- _As a maintainer_, I want to establish validation rules that reject conflicting or redundant annotations at runtime, so that users do not face unpredictable behavior due to overlapping DNS rules.
	- _As a maintainer_, I want to collaborate with other Kubernetes SIGs (e.g., SIG-Network, SIG-Auth) to align annotation standards,
so that External-DNS remains compatible with evolving Kubernetes best practices.

### API

Annotations should follow a structured versioning approach, with a clear mapping from `alpha` to `beta`.

Example:

Proposed transition strategy:

-   Maintain support for `alpha` annotations with deprecation warnings.
-   Introduce annotations with improved validation.
-   Provide an automated tool to keep up-to-date documentation.
-   Annotation interface is standartized

### Behavior

-   `external-dns` should recognize both `alpha` and `beta` annotations where applicable.
-   Warnings should be logged when deprecated annotations are used.
-   Warnings should be logged when annotation is not supported by source or provider.
-   Future major versions should drop support for `alpha` annotations after a defined period.
-   Validation logic will ensure that only valid annotations are accepted, providing clear error messages when issues are detected.

### Drawbacks

-   The implementation will require an initial investment of time and resources.
-   Maintaining support for multiple annotation versions may introduce complexity.

## Alternatives

### Alternative 1: Continue with the Current System

-   Pros: No immediate changes required.
-   Cons: Ongoing challenges with documentation, versioning, and user experience.

### Alternative 2: Keep Annotations in Alpha Permanently

-   Pros: No migration burden for users.
-   Cons: Lack of stability signals to users, discouraging adoption.

### Alternative 3: Deprecate Annotations in Favor of CRDs

-   Pros: More structured and Kubernetes-native approach.
-   Cons: Significant changes required for existing users, making adoption harder.