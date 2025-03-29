```yaml
---
title: "Proposal: Defining a path to Beta for DNSEndpoint API"
version: v1alpha1
authors: @ivankatliarchuk, @raffo, @szuecs
creation-date: 2025-02-09
status: approved
---
```

# Proposal: Defining a path to Beta for DNSEndpoint API

## Summary

The `DNSEndpoint` API in Kubernetes SIGs `external-dns` is currently in alpha. To ensure its stability and readiness for production environments, we propose defining and agreeing upon the necessary requirements for its graduation to beta.
By defining clear criteria, we aim to ensure stability, usability, and compatibility with the broader Kubernetes ecosystem. On completions of all this items, we should be in the position to graduate `DNSEndpoint` to `v1beta`.

## Motivation

The DNSEndpoint API is a crucial component of the ExternalDNS project, allowing users to manage DNS records dynamically.
Currently, it remains in the alpha stage, limiting its adoption due to potential instability and lack of guaranteed backward compatibility. By advancing to beta, we can provide users with a more reliable API and encourage wider adoption with confidence in its long-term viability and support.

### Goals

- Define the necessary requirements for `DNSEndpoint` API to reach beta status.
- Improve API stability, usability, and documentation.
- Improve test coverage, automate documentation creation, and validation mechanisms.
- Ensure backward compatibility and migration strategies from alpha to beta.
- Collect and incorporate feedback from existing users to refine the API.
- Address any identified issues or limitations in the current API design.

### Non-Goals

- This proposal does not cover the graduation of ExternalDNS itself to a stable release.
- Making `DNSEndpoint` a stable (GA) API at this stage.
- It does not include implementation details for specific DNS providers.
- It does not introduce new functionality beyond stabilizing the DNSEndpoint API.
- Redesigning the API from scratch.
- Introducing breaking changes that would require significant refactoring for existing users.

## Proposal

The proposal aims to formalize the promotion process by addressing API design, user needs, and implementation details.

To graduate the `DNSEndpoint` API to beta, we propose the following actions:

1. Capture feedback from the community on missing functionality for DNSEndpoint CRD
   - In a form of Github issue, pin the issue to the project
   - Link all CRD related issues to it
2. Refactor `endpoint` folder, move away `api/crd` related stuff to `apis/<apiVersion> folder`
3. Documentation for API to be generated automatically with test coverage, similar to `docs/flags.md`
4. APIs and CRDs discoverable. [doc.crds.dev](https://doc.crds.dev/github.com/kubernetes-sigs/external-dns). Example [crossplane](https://doc.crds.dev/github.com/crossplane/crossplane@v0.10.0)
5. Review and change .status object such that people can debug and monitor DNSEndpoint object behavior.
6. Introduce metrics related to DNSEndpoint CRD
   - Number of CRDs discovered
   - Number of CRDs by status success|fail

Proposed folder structure for `apis`. Examples - [gateway-api](https://github.com/kubernetes-sigs/gateway-api/tree/main/apis)

***Multiple APIs under same version***

```yml
├── apis
│   ├── v1alpha
│   │   ├── util/validation
│   │   ├── doc.go
│   │   └── zz_generated.***.go
│   ├── v1beta  # outside of scope currently, just an example
│   │   ├── util/validation
│   │   ├── doc.go
│   │   └── zz_generated.***.go
│   ├── v1       # outside of scope currently, just an example
│   │   ├── util/validation
│   │   ├── doc.go
│   │   └── zz_generated.***.go
```

Or similar folder structure for `apis`. Examples - [cert-manager](https://github.com/cert-manager/cert-manager/tree/master/pkg/apis)

***APIs versioned independently***

```yml
├── apis
│   ├── dnsendpoint
│   │   ├── v1alpha
│   │   │   ├── util/validation
│   │   │   ├── doc.go
│   │   │   └── zz_generated.***.go
│   │   ├── v1beta  # outside of scope currently, just an example
│   │   │   ├── util/validation
│   │   │   ├── doc.go
│   │   │   └── zz_generated.***.go
│   │   ├── v1       # outside of scope currently, just an example
│   │   │   ├── util/validation
│   │   │   ├── doc.go
│   │   │   └── zz_generated.***.go
│   ├── dnsentry
│   │   ├── v1alpha
```

### User Stories

#### Story 1: Cluster Operator/Admin Managing External DNS

*As a cluster operator or administrator*, I want a stable `DNSEndpoint` API to reliably manage DNS records within Kubernetes so that I can ensure consistent and automated DNS resolution for my services.

#### Story 2: Developers Integrating External DNS

*As a developer*, I want a well-documented `DNSEndpoint` API that allows me to programmatically define and manage DNS records without worrying about breaking changes.

#### Story 3: Cloud-Native Deployments

*As a SRE*, I need a tested and validated `DNSEndpoint` API that integrates seamlessly with cloud-native networking services, ensuring high availability and scalability.

#### Story 4: Platform Engineer

*As a platform engineer*, I want stronger validation and defaulting so that I can reduce misconfigurations and operational overhead.

### API

The DNSEndpoint API should provide a robust Custom Resource Definition (CRD) with well-defined fields and validation.

#### DNSEndpoint

- [ ] DNSEndpoint do not have any changes from v1alpha1.
- [ ] DNSEndpoint to have changes from v1alpha1. `TBD`

```yml
apiVersion: externaldns.k8s.io/v1beta1
kind: DNSEndpoint
metadata:
  name: example-endpoint
spec:
  endpoints:
    - dnsName: "example.com"
      recordType: "A"
      targets:
        - "192.168.1.1"
      ttl: 300
    - dnsName: "www.example.com"
      recordType: "CNAME"
      targets:
        - "example.com"
```

### Behavior

How should the new CRD or feature behave? Are there edge cases?

### Drawbacks

- Transitioning to beta may require deprecating certain alpha features that are deemed unstable.
- Increased maintenance effort to ensure stability and backward compatibility.
- Users of the alpha API may need to adjust their configurations if breaking changes are introduced.
- Additional maintenance and support burden for the `external-dns` maintainers.

## Alternatives

1. **Remain in Alpha**: The DNSEndpoint API could remain in alpha indefinitely, but this would discourage adoption and limit its reliability.

- Pros: No immediate changes or migration concerns.
- Cons: Lack of progress discourages adoption, and users may seek alternative solutions.

2. **Graduate Directly to GA**: Skipping the beta phase could accelerate stability, but it would limit the opportunity for community feedback and refinement.

3. **Introduce a New API Version**: Instead of modifying the existing API, a new version (e.g., `v2alpha1`) could be introduced, allowing gradual migration.

    - Pros: Allowing gradual migration like `v1alpha1` -> `v2alpha1` -> `v1beta`
    - Cons: This approach would require maintaining multiple versions simultaneously.

4. **Redesign the API Before Graduation**

    - Pros: Provides an opportunity to fix any fundamental design flaws before moving to beta.
    - Cons: Increases complexity, delays the beta release, and may introduce unnecessary work for existing users.

5. **Deprecate DNSEndpoint and Rely on External Solutions or Annotations**

    - Pros: Potentially reduces the maintenance burden on the Kubernetes SIG.
    - Cons: Forces users to migrate to third-party solutions or away from CRDs, reducing the cohesion of external-dns within Kubernetes.
