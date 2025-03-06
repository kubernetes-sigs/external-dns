# External DNS Deprecation Policy

This document defines the Deprecation Policy for External DNS.

Kubernetes is a dynamic system driven by APIs, which evolve with each new release. A crucial aspect of any API-driven system is having a well-defined deprecation policy.
This policy informs users about APIs that are slated for removal or modification. Kubernetes follows this principle and periodically refines or upgrades its APIs or capabilities.
Consequently, older features are marked as deprecated and eventually phased out. To avoid breaking existing users, we should follow a simple deprecation policy for behaviors that a slated to be removed.

The features and capabilities either to evolve or need to be removed.

## Deprecation Policy

We follow the [Kubernetes Deprecation Policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/) and [API Versioning Scheme](https://kubernetes.io/docs/reference/using-api/#api-versioning): alpha, beta, GA.
It is therefore important to be aware of deprecation announcements and know when API versions will be removed, to help minimize the effect.

### Scope

* CRDs and API Objects and fields: `.Spec`, `.Status` and `.Status.Conditions[]`
* Annotations objects or it's values
* Controller Configuration: CLI flags & environment variables
* Metrics as defined in the [Kubernetes docs](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-a-metric)
* Revert a specific behavior without an alternative (flag,crd or annotation)

### Non-Scope

Everything not listed in scope is not subject to this deprecation policy and it is subject to breaking changes, updates at any point in time, and deprecation - as long as it follows the Deprecation Process listed below.

This includes, but isn't limited to:

* Any feature/specific behavior not in Scope.
* Source code imports
* Source code refactorings
* Helm Charts
* Release process
* Docker Images (including multi-arch builds)
* Image Signature (including provenance, providers, keys)

## Including features and behaviors to the Deprecation Policy

Any `maintainer` or `contributor` may propose including a feature, component, or behavior out of scope to be in scope of the deprecation policy.

The proposal must clearly outline the rationale for inclusion, the impact on users, stability, long term maintenance plan, and day-to-day activities, if such.

The proposal must be formalized by submitting a `docs/proposal/EDP-XXX.md` document in a Pull Request. Pull request must be labeled with `kind/proposal`.

The proposal template location is [here](docs/proposal/design-template.md). The template is quite complete, one can remove any unnecessary or irrelevant section on a specific proposal.

## Deprecation Process

### Nomination of Deprecation

Any maintainer may propose deprecating a feature, component, or behavior (both in and out of scope). In Scope changes must abide to the Deprecation Policy above.

The proposal must clearly outline the rationale for deprecation, the impact on users, and any alternatives, if such.

The proposal must be formalized by submiting a `design` document as a Pull Request.

### Showcase to Maintainers

The proposing maintainer must present the proposed deprecation to the maintainer group. This can be done synchronously during a community meeting or asynchronously, through a GitHub Pull Request.

### Voting

A majority vote of maintainers is required to approve the deprecation.
Votes may be conducted asynchronously, with a reasonable deadline for responses (e.g., one week). Lazy Consensus applies if the reasonable deadline is extended, with a minimal of at least one other maintainer approving the changes.

### Implementation

Upon approval, the proposing maintainer is responsible for implementing the changes required to mark the feature as deprecated. This includes:

* Updating the codebase with deprecation warnings where applicable.
  * log.Warn("The XXX is on the path of ***DEPRECATION***. We recommend that you use YYY (link to docs)")
* Documenting the deprecation in release notes and relevant documentation.
* Updating APIs, metrics, or behaviors per the Kubernetes Deprecation Policy if in scope.
* If the feature is entirely deprecated, archival of any associated repositories (external provider as example).

### Deprecation Notice in Release

Deprecation must be introduced in the next release. The release must follow semantic versioning:

* If the project is in the 0.x stage, a `minor` version `bump` is required.
* For projects 1.x and beyond, a major version bump is required. For the features completely removed.
  * If it's a flag change/flip, the `minor` version `bump` is acceptable

### Full Deprecation and Removal

The removal must follow standard Kubernetes deprecation timelines if the feature is in scope.
