# ExternalDNS Helm Chart Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

<!-- ## [vX.Y.Z] - UNRELEASED
### Highlights
### All Changes
- Added
- Updated
- Changed
- Fixed
- Deprecated
- Removed -->

## [v1.12.1] - 2023-02-06

### All Changes

- Updated _ExternalDNS_ version to [v0.13.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.13.2). ([#3371](https://github.com/kubernetes-sigs/external-dns/pull/3371)) [@stevehipwell](https://github.com/stevehipwell)
- Added `secretConfiguration.subPath` to mount specific files from secret as a sub-path. ([#3227](https://github.com/kubernetes-sigs/external-dns/pull/3227)) [@jkroepke](https://github.com/jkroepke)
- Changed to use `registry.k8s.io` instead of `k8s.gcr.io`. ([#3261](https://github.com/kubernetes-sigs/external-dns/pull/3261)) [@johngmyers](https://github.com/johngmyers)

## [v1.12.0] - 2022-11-29

### All Changes

- Added ability to provide ExternalDNS with secret configuration via `secretConfiguration`. ([#3144](https://github.com/kubernetes-sigs/external-dns/pull/3144)) [@jkroepke](https://github.com/jkroepke)
- Added the ability to template `provider` & `extraArgs`. ([#3144](https://github.com/kubernetes-sigs/external-dns/pull/3144)) [@jkroepke](https://github.com/jkroepke)
- Added the ability to customise the service account labels. ([#3145](https://github.com/kubernetes-sigs/external-dns/pull/3145)) [@jkroepke](https://github.com/jkroepke)
- Updated _ExternalDNS_ version to [v0.13.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.13.1). ([#3197](https://github.com/kubernetes-sigs/external-dns/pull/3197)) [@stevehipwell](https://github.com/stevehipwell)

## [v1.11.0] - 2022-08-10

### Added

- Added support to configure `dnsPolicy` on the Helm chart deployment. [@michelzanini](https://github.com/michelzanini)
- Added ability to customise the deployment strategy. [mac-chaffee](https://github.com/mac-chaffee)

### Changed

- Updated _ExternalDNS_ version to [v0.12.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.12.2). [@stevehipwell](https://github.com/stevehipwell)
- Changed default deployment strategy to `Recreate`. [mac-chaffee](https://github.com/mac-chaffee)

## [v1.10.1] - 2022-07-11

### Fixed

- Fixed incorrect addition of `namespace` to `ClusterRole` & `ClusterRoleBinding`. [@stevehipwell](https://github.com/stevehipwell)

## [v1.10.0] - 2022-07-08

### Added

- Added `commonLabels` value to allow the addition of labels to all resources. [@stevehipwell](https://github.com/stevehipwell)
- Added support for [Process Namespace Sharing](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/) via the `shareProcessNamespace`
 value. ([#2715](https://github.com/kubernetes-sigs/external-dns/pull/2715)) [@wolffberg](https://github.com/wolffberg)

### Changed

- Update _ExternalDNS_ version to [v0.12.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.12.0). [@vojtechmares](https://github.com/vojtechmares)
- Set resource namespaces to `{{ .Release.Namespace }}` in the templates instead of waiting until apply time for inference. [@stevehipwell](https://github.com/stevehipwell)
- Fixed `rbac.additionalPermissions` default value.([#2796](https://github.com/kubernetes-sigs/external-dns/pull/2796)) [@tamalsaha](https://github.com/tamalsaha)

## [v1.9.0] - 2022-04-19

### Changed

- Update _ExternalDNS_ version to [v0.11.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.11.0). ([#2690](https://github.com/kubernetes-sigs/external-dns/pull/2690)) [@stevehipwell](https://github.com/stevehipwell)

## [v1.8.0] - 2022-04-13

### Added

- Add annotations to Deployment. ([#2477](https://github.com/kubernetes-sigs/external-dns/pull/2477)) [@beastob](https://github.com/beastob)

### Changed

- Fix RBAC for `istio-virtualservice` source when `istio-gateway` isn't also added. ([#2564](https://github.com/kubernetes-sigs/external-dns/pull/2564)) [@mcwarman](https://github.com/mcwarman)
