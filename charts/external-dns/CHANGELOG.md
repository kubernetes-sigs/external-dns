# ExternalDNS Helm Chart Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

<!-- ## [UNRELEASED]
### Added
### Changed
### Deprecated
### Removed -->

## [v1.9.0] - UNRELEASED

### Changed

- Update _ExternalDNS_ version to [v0.11.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.11.0).

## [v1.8.0] - UNRELEASED

### Added

- Add annotations to Deployment. [#2477](https://github.com/kubernetes-sigs/external-dns/pull/2477) from @beastob

### Changed

- Fix RBAC for `istio-virtualservice` source when `istio-gateway` isn't also added. [#2564](https://github.com/kubernetes-sigs/external-dns/pull/2564) from @mcwarman
