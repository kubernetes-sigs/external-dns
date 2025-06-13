# ExternalDNS Helm Chart Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

<!--
### Added - For new features.
### Changed - For changes in existing functionality.
### Deprecated - For soon-to-be removed features.
### Removed - For now removed features.
### Fixed - For any bug fixes.
### Security - In case of vulnerabilities.
-->

## [UNRELEASED]

## [v1.17.0] - 2025-06-04

### Changed

- Allow extraArgs to also be a map enabling overrides of individual values. ([#5293](https://github.com/kubernetes-sigs/external-dns/pull/5293)) _@frittentheke_
- Update CRD. ([#5287](https://github.com/kubernetes-sigs/external-dns/pull/5287)) _@mloiseleur_
- Update CRD. ([#5446](https://github.com/kubernetes-sigs/external-dns/pull/5446)) _@mloiseleur_
- Update _ExternalDNS_ OCI image version to [v0.17.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.17.0). ([#5479](https://github.com/kubernetes-sigs/external-dns/pull/5479)) _@stevehipwell_

### Fixed

- Fix wrong type definitions for webhook probes. ([#5297](https://github.com/kubernetes-sigs/external-dns/pull/5297)) _@semnell_
- Update schema with latest plugin release. ([#5510](https://github.com/kubernetes-sigs/external-dns/pull/5510)) _@mloiseleur

## [v1.16.1] - 2025-04-10

### Changed

- Set defaults for `automountServiceAccountToken` and `serviceAccount.automountServiceAccountToken` to `true` in Helm chart values. ([#5207](https://github.com/kubernetes-sigs/external-dns/pull/5207)) _@t3mi_

### Fixed

- Correctly handle `txtPrefix` and `txtSuffix` arguments when both are provided. ([#5250](https://github.com/kubernetes-sigs/external-dns/pull/5250)) _@ivankatliarchuk_
- Add missing types in the schema for empty values. ([#5228](https://github.com/kubernetes-sigs/external-dns/pull/5228)) _@ivankatliarchuk_
- Add missing types in the schema for empty values. ([#5207](https://github.com/kubernetes-sigs/external-dns/pull/5207)) _@t3mi_

## [v1.16.0] - 2025-03-20

### Added

- Add helm testing framework `helm plugin unittest`. ([#5137](https://github.com/kubernetes-sigs/external-dns/pull/5137)) _@ivankatliarchuk_
- Add ability to generate schema with `helm plugin schema`. ([#5075](https://github.com/kubernetes-sigs/external-dns/pull/5075)) _@ivankatliarchuk_
- Add `docs/contributing/dev-guide.md#helm-values` guide. ([#5075](https://github.com/kubernetes-sigs/external-dns/pull/5075)) _@ivankatliarchuk_

### Changed

- Regenerate JSON schema with `helm-values-schema-json' plugin. ([#5075](https://github.com/kubernetes-sigs/external-dns/pull/5075)) _@ivankatliarchuk_
- Update _ExternalDNS_ OCI image version to [v0.16.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.16.1). ([#5201](https://github.com/kubernetes-sigs/external-dns/pull/5201)) _@stevehipwell_

## [v1.15.2] - 2025-02-14

### Changed

- Added `transportservers` resource to ClusterRole when specifying `f5-transportserver` or `f5-virtualserver` as a source. ([#5066](https://github.com/kubernetes-sigs/external-dns/pull/5066)) _@visokoo_

### Fixed

- Fixed handling of non-string types in `serviceAccount.metadata.annotations` field. ([#5067](https://github.com/kubernetes-sigs/external-dns/pull/5067)) _@hjoshi123_
- Fixed regression where `affinity.nodeAffinity` was being ignored. ([#5046](https://github.com/kubernetes-sigs/external-dns/pull/5046)) _@mkhpalm_

## [v1.15.1] - 2025-01-27

### Added

- Added ability to configure `imagePullSecrets` via helm `global` value. ([#4667](https://github.com/kubernetes-sigs/external-dns/pull/4667)) _@jkroepke_
- Added options to configure `labelFilter` and `managedRecordTypes` via dedicated helm values. ([#4849](https://github.com/kubernetes-sigs/external-dns/pull/4849)) _@abaguas_

### Changed

- Allow templating `serviceaccount.annotations` keys and values, by rendering them using the `tpl` built-in function. ([#4958](https://github.com/kubernetes-sigs/external-dns/pull/4958)) _@fcrespofastly_
- Updated _ExternalDNS_ OCI image version to [v0.15.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.15.1). ([#5028](https://github.com/kubernetes-sigs/external-dns/pull/5028)) _@stevehipwell_

### Fixed

- Fixed automatic addition of pod selector labels to `affinity` and `topologySpreadConstraints` if not defined. ([#4666](https://github.com/kubernetes-sigs/external-dns/pull/4666)) _@pvickery-ParamountCommerce_
- Fixed missing Ingress permissions when using Istio sources. ([#4845](https://github.com/kubernetes-sigs/external-dns/pull/4845)) _@joekhoobyar_

## [v1.15.0] - 2024-09-11

### Changed

- Updated _ExternalDNS_ OCI image version to [v0.15.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.15.0). ([#4735](https://github.com/kubernetes-sigs/external-dns/pull/4735)) _@stevehipwell_

### Fixed

- Fixed `provider.webhook.resources` behavior to correctly leverage resource limits. ([#4560](https://github.com/kubernetes-sigs/external-dns/pull/4560)) _@crutonjohn_
- Fixed `provider.webhook.imagePullPolicy` behavior to correctly leverage pull policy. ([#4643](https://github.com/kubernetes-sigs/external-dns/pull/4643)) _@kimsondrup_
- Fixed to add correct webhook metric port to `Service` and `ServiceMonitor`. ([#4643](https://github.com/kubernetes-sigs/external-dns/pull/4643)) _@kimsondrup_
- Fixed to no longer require the unauthenticated webhook provider port to be exposed for health probes. ([#4691](https://github.com/kubernetes-sigs/external-dns/pull/4691)) _@kimsondrup_ & _@hatrx_

## [v1.14.5] - 2024-06-10

### Added

- Added support for `extraContainers` argument. ([#4432](https://github.com/kubernetes-sigs/external-dns/pull/4432)) _@omerap12_
- Added support for setting `excludeDomains` argument.  ([#4380](https://github.com/kubernetes-sigs/external-dns/pull/4380)) _@bford-evs_

### Changed

- Updated _ExternalDNS_ OCI image version to [v0.14.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.14.2). ([#4541](https://github.com/kubernetes-sigs/external-dns/pull/4541)) _@stevehipwell_
- Updated `DNSEndpoint` CRD. ([#4541](https://github.com/kubernetes-sigs/external-dns/pull/4541)) _@stevehipwell_
- Changed the implementation for `revisionHistoryLimit` to be more generic. ([#4541](https://github.com/kubernetes-sigs/external-dns/pull/4541)) _@stevehipwell_

### Fixed

- Fixed the `ServiceMonitor` job name to correctly use the instance label. ([#4541](https://github.com/kubernetes-sigs/external-dns/pull/4541)) _@stevehipwell_

## [v1.14.4] - 2024-04-05

### Added

- Added support for setting `dnsConfig`. ([#4265](https://github.com/kubernetes-sigs/external-dns/pull/4265)) _@davhdavh_
- Added support for `DNSEndpoint` CRD. ([#4322](https://github.com/kubernetes-sigs/external-dns/pull/4322)) _@onedr0p_

### Changed

- Updated _ExternalDNS_ OCI image version to [v0.14.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.14.1). ([#4357](https://github.com/kubernetes-sigs/external-dns/pull/4357)) _@stevehipwell_

## [v1.14.3] - 2024-01-26

### Fixed

- Fixed args for webhook deployment. ([#4202](https://github.com/kubernetes-sigs/external-dns/pull/4202)) [@webwurst](https://github.com/webwurst)
- Fixed support for `gateway-grpcroute`, `gateway-tlsroute`, `gateway-tcproute` & `gateway-udproute`. ([#4205](https://github.com/kubernetes-sigs/external-dns/pull/4205)) [@orenlevi111](https://github.com/orenlevi111)
- Fixed incorrect implementation for setting the `automountServiceAccountToken`. ([#4208](https://github.com/kubernetes-sigs/external-dns/pull/4208)) [@stevehipwell](https://github.com/stevehipwell)

## [v1.14.2] - 2024-01-22

### Fixed

- Restore template support in `.Values.provider` and `.Values.provider.name`

## [v1.14.1] - 2024-01-12

### Fixed

- Fixed webhook install failure: `"http-webhook-metrics": must be no more than 15 characters`. ([#4173](https://github.com/kubernetes-sigs/external-dns/pull/4173)) [@gabe565](https://github.com/gabe565)

## [v1.14.0] - 2024-01-10

### Added

- Added the option to explicitly enable or disable service account token automounting. ([#3983](https://github.com/kubernetes-sigs/external-dns/pull/3983)) [@gilles-gosuin](https://github.com/gilles-gosuin)
- Added the option to configure revisionHistoryLimit on the K8s Deployment resource. ([#4008](https://github.com/kubernetes-sigs/external-dns/pull/4008)) [@arnisoph](https://github.com/arnisoph)
- Added support for webhook providers, as a sidecar. ([#4032](https://github.com/kubernetes-sigs/external-dns/pull/4032) [@mloiseleur](https://github.com/mloiseleur)
- Added the option to configure ipFamilyPolicy and ipFamilies of external-dns Service.  ([#4153](https://github.com/kubernetes-sigs/external-dns/pull/4153)) [@dongjiang1989](https://github.com/dongjiang1989)

### Changed

- Avoid unnecessary pod restart on each helm chart version. ([#4103](https://github.com/kubernetes-sigs/external-dns/pull/4103)) [@jkroepke](https://github.com/jkroepke)
- Updated _ExternalDNS_ OCI image version to [v0.14.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.14.0). ([#4073](https://github.com/kubernetes-sigs/external-dns/pull/4073)) [@appkins](https://github.com/appkins)

### Deprecated

- The `secretConfiguration` value has been deprecated in favour of creating secrets external to the Helm chart and configuring their use via the `extraVolumes` & `extraVolumeMounts` values. ([#4161](https://github.com/kubernetes-sigs/external-dns/pull/4161)) [@stevehipwell](https://github.com/stevehipwell)

## [v1.13.1] - 2023-09-08

### Added

- Added RBAC for Traefik to ClusterRole. ([#3325](https://github.com/kubernetes-sigs/external-dns/pull/3325)) [@ThomasK33](https://github.com/thomask33)
- Added support for init containers. ([#3325](https://github.com/kubernetes-sigs/external-dns/pull/3838)) [@calvinbui](https://github.com/calvinbui)

### Changed

- Disallowed privilege escalation in container security context and set the seccomp profile type to `RuntimeDefault`. ([#3689](https://github.com/kubernetes-sigs/external-dns/pull/3689)) [@nrvnrvn](https://github.com/nrvnrvn)
- Updated _ExternalDNS_ OCI image version to [v0.13.6](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.13.6). ([#3917](https://github.com/kubernetes-sigs/external-dns/pull/3917)) [@stevehipwell](https://github.com/stevehipwell)

### Removed

- Removed RBAC rule for already removed `contour-ingressroute` source. ([#3764](https://github.com/kubernetes-sigs/external-dns/pull/3764)) [@johngmyers](https://github.com/johngmyers)

## [v1.13.0] - 2023-03-30

### All Changes

- Updated _ExternalDNS_ version to [v0.13.5](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.13.5). ([#3661](https://github.com/kubernetes-sigs/external-dns/pull/3661)) [@GMartinez-Sisti](https://github.com/GMartinez-Sisti)
- Adding missing gateway-httproute cluster role permission. ([#3541](https://github.com/kubernetes-sigs/external-dns/pull/3541)) [@nicon89](https://github.com/nicon89)

## [v1.12.2] - 2023-03-30

### All Changes

- Added support for ServiceMonitor relabelling. ([#3366](https://github.com/kubernetes-sigs/external-dns/pull/3366)) [@jkroepke](https://github.com/jkroepke)
- Updated chart icon path. ([#3492](https://github.com/kubernetes-sigs/external-dns/pull/3494)) [kundan2707](https://github.com/kundan2707)
- Added RBAC for Gateway-API resources to ClusterRole. ([#3499](https://github.com/kubernetes-sigs/external-dns/pull/3499)) [@michaelvl](https://github.com/MichaelVL)
- Added RBAC for F5 VirtualServer to ClusterRole. ([#3503](https://github.com/kubernetes-sigs/external-dns/pull/3503)) [@mikejoh](https://github.com/mikejoh)
- Added support for running ExternalDNS with namespaced scope. ([#3403](https://github.com/kubernetes-sigs/external-dns/pull/3403)) [@jkroepke](https://github.com/jkroepke)
- Updated _ExternalDNS_ version to [v0.13.4](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.13.4). ([#3516](https://github.com/kubernetes-sigs/external-dns/pull/3516)) [@stevehipwell](https://github.com/stevehipwell)

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

<!--
RELEASE LINKS
-->
[UNRELEASED]: https://github.com/kubernetes-sigs/external-dns/tree/master/charts/external-dns
[v1.17.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.17.0
[v1.16.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.16.1
[v1.16.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.16.0
[v1.15.2]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.15.2
[v1.15.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.15.1
[v1.15.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.15.0
[v1.14.5]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.5
[v1.14.4]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.4
[v1.14.3]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.3
[v1.14.2]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.2
[v1.14.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.1
[v1.14.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.14.0
[v1.13.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.13.1
[v1.13.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.13.0
[v1.12.2]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.12.2
[v1.12.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.12.1
[v1.12.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.12.0
[v1.11.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.11.0
[v1.10.1]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.10.1
[v1.10.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.10.0
[v1.9.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.9.0
[v1.8.0]: https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.8.0
