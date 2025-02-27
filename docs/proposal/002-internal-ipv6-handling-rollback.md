<!-- clone me -->
```yaml
---
title: "Proposal: Rollback IPv6 internal Node IP exposure"
version: if applicable
authors: @ivankatliarchuk, @szuecs, @mloiseleur
creation-date: 2025-01-01
status: approved
---
```

# Introduce Feature Flag for IPv6 Internal Node IP Handling in ''external-dns'' and Change the behavior

## Summary

This proposal aims to introduce a feature flag in 'external-dns' to control the handling of IPv6 internal node IPs.
In the current version, the feature flag will default to the existing behavior. In the next `minor` or `minor+N` version, the default behavior will be reversed, encouraging users to adopt the new behavior while providing a transition period.

## Motivation

The discussion in [issue#4566](https://github.com/kubernetes-sigs/external-dns/issues/4566) and the
subsequent [pr#4574](https://github.com/kubernetes-sigs/external-dns/pull/4574) and [pr#4808](https://github.com/kubernetes-sigs/external-dns/pull/4808) highlighted concerns regarding the treatment of IPv6 internal node IPs.
To address these concerns without causing immediate disruption, a feature flag will allow users to opt-out the current behavior, providing flexibility during the transition.

## Goals

- Introduce feature to toggle the handling of IPv6 internal node IPs

## Non-Goals

- ***Propose/Add an annotation for this specific use case***
  - Provide support for `external-dns.alpha.kubernetes.io/expose-internal-ipv6` in follow-up releases.
  - Managing dual annotation and flag may introduce complexity.

## Proposal

- ***Introduce Feature Flag***
  - Add a feature flag, e.g., `--expose-internal-ipv6=true`, to control the handling of IPv6 internal node IPs.
  - In the current version, this flag will default to `true`, maintaining the existing behavior.

- ***Flip Default Behavior in Next Minor Version***
  - In the subsequent minor release, change the default value of `--expose-internal-ipv6` to `false`, adopting the new behavior by default.
  - Users can still override this behavior by explicitly setting the flag as needed.

Proposed Changes in `source/node.go` file.

```go
// IPv6 addresses are labeled as NodeInternalIP despite being usable externally as well.
if addr.Type == v1.NodeInternalIP && ns.exposeInternalIP && ... {
	pv6Addresses = append(ipv6Addresses, addr.Address)
}
```

## User Stories

- **As a cluster Operator or Administrator**, I want to control the handling of IPv6 internal node IPs to align with defined network topology and configuration.

- **As a SecDevOps**, I want to ensure that `external-dns` does not expose internal IPv6 node addresses via public DNS records, so that I can prevent unintended data leaks and reduce the attack surface of my Kubernetes cluster.

- **As a SecDevOps**, I want to use a feature flag to selectively enable or disable the new IPv6 behavior in `external-dns`, so that I can evaluate its security impact before it becomes the default setting in future releases.

- **As a SecDevOps**, I want to use a feature flag to selectively enable or disable the new IPv6 behavior in `external-dns`, so that I can detect misconfigurations, act on potential security incidents, and ensure compliance with security policies.

## Implementation Steps

- Code Changes:
  - Implement the feature flag in the 'external-dns' codebase to toggle the handling of IPv6 internal node IPs.

- Documentation:
  - Update the 'external-dns' documentation to include information about the feature flag, its purpose, and usage examples.

## Drawbacks

- Introducing a feature flag adds complexity to the configuration and codebase.
- Changing default behavior in a future release may still cause issues for users who are unaware of the change.

## Alternatives

- ***Immediate Behavior Change***
  - Directly change the behavior without a feature flag, which could lead to unexpected issues for users.
- ***No Change***
  - Maintain the current behavior, potentially leaving the concerns unaddressed.
  - Users may not be able to update an `external-dns` version due to security, compliance or any other concerns.
