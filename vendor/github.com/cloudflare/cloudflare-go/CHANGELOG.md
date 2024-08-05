<<<<<<< HEAD
## 0.51.0 (Unreleased)

## 0.50.0 (September 14, 2022)

ENHANCEMENTS:

* auditlogs: add support for hide_user_logs filter parameter ([#1075](https://github.com/cloudflare/cloudflare-go/issues/1075))

BUG FIXES:

* cloudflare: exiting closer to the source on context timeouts to improve error messaging and better defend from potential edge cases ([#1080](https://github.com/cloudflare/cloudflare-go/issues/1080))
* origin certificate: Fix API auth type used ([#1082](https://github.com/cloudflare/cloudflare-go/issues/1082))

DEPENDENCIES:

* deps: bumps github.com/urfave/cli/v2 from 2.11.2 to 2.14.0 ([#1077](https://github.com/cloudflare/cloudflare-go/issues/1077))
* deps: bumps github.com/urfave/cli/v2 from 2.14.0 to 2.14.1 ([#1081](https://github.com/cloudflare/cloudflare-go/issues/1081))
* deps: bumps github.com/urfave/cli/v2 from 2.14.1 to 2.15.0 ([#1085](https://github.com/cloudflare/cloudflare-go/issues/1085))
* deps: bumps github.com/urfave/cli/v2 from 2.15.0 to 2.16.3 ([#1086](https://github.com/cloudflare/cloudflare-go/issues/1086))

## 0.49.0 (August 31st, 2022)

ENHANCEMENTS:

* access_service_token: add support for refreshing an existing token in place ([#1074](https://github.com/cloudflare/cloudflare-go/issues/1074))
* api: addded context and headers to Raw method ([#1068](https://github.com/cloudflare/cloudflare-go/issues/1068))
* api_shield: add GET/PUT for API Shield Configuration ([#1059](https://github.com/cloudflare/cloudflare-go/issues/1059))
* pages_project: Add `kv_namespaces`, `durable_object_namespaces`, `r2_buckets`, and `d1_databases` bindings to deployment config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* pages_project: Add `preview_deployment_setting`, `preview_branch_includes`, and `preview_branch_excludes` to source config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* pages_project: Add `production_branch` field ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
* teams_account: add support for `os_distro_name` and `os_distro_revision` ([#1073](https://github.com/cloudflare/cloudflare-go/issues/1073))
* url_normalization_settings: Add APIs to get and update URL normalization settings ([#1071](https://github.com/cloudflare/cloudflare-go/issues/1071))
* workers: Support for multipart encoding for DownloadWorker on a module-format Worker script ([#1040](https://github.com/cloudflare/cloudflare-go/issues/1040))

BUG FIXES:

* cloudflare: fix nil dereference error in makeRequestWithAuthTypeAndHeaders ([#1072](https://github.com/cloudflare/cloudflare-go/issues/1072))
* email_routing_rules: Fix response for email routing catch all rule. ([#1070](https://github.com/cloudflare/cloudflare-go/issues/1070))
* email_routing_settings: change enable endpoint from `enabled` to `enable` ([#1060](https://github.com/cloudflare/cloudflare-go/issues/1060))
* stream: Update pctComplete to string from int ([#1066](https://github.com/cloudflare/cloudflare-go/issues/1066))

DEPENDENCIES:

* deps: bumps goreleaser/goreleaser-action from 3.0.0 to 3.1.0 ([#1067](https://github.com/cloudflare/cloudflare-go/issues/1067))

## 0.48.0 (August 22nd, 2022)

ENHANCEMENTS:

* errors: add some error type convenience functions for mocking and inspection ([#1047](https://github.com/cloudflare/cloudflare-go/issues/1047))
* pages_project: Add compatibility date and compatibility_flags to pages deployment configs ([#1051](https://github.com/cloudflare/cloudflare-go/issues/1051))
* teams_account: add support for `suppress_footer` ([#1053](https://github.com/cloudflare/cloudflare-go/issues/1053))

BUG FIXES:

* r2: fix create bucket endpoint ([#1035](https://github.com/cloudflare/cloudflare-go/issues/1035))
* tunnel_configuration: Remove unnecessary double-unmarshalling due to changes in the API ([#1046](https://github.com/cloudflare/cloudflare-go/issues/1046))

## 0.47.1 (August 18th, 2022)

BUG FIXES:

* zonelockdown: add `Priority` to `ZoneLockdownCreateParams` and `ZoneLockdownUpdateParams` ([#1052](https://github.com/cloudflare/cloudflare-go/issues/1052))

## 0.47.0 (August 17th, 2022)

BREAKING CHANGES:

* certificate_packs: deprecate "custom" configuration for ACM everywhere ([#1032](https://github.com/cloudflare/cloudflare-go/issues/1032))

ENHANCEMENTS:

* cloudflare: make it clear when the rate limit retries have been exhausted ([#1043](https://github.com/cloudflare/cloudflare-go/issues/1043))
* email_routing_destination: Adds support for the email routing destination API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* email_routing_rules: Adds support for the email routing rules API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* email_routing_settings: Adds support for the email routing settings API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
* filter: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
* firewall_rule: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
* lockdown: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1017](https://github.com/cloudflare/cloudflare-go/issues/1017))
* r2: Add in support for creating and deleting R2 buckets ([#1028](https://github.com/cloudflare/cloudflare-go/issues/1028))
* rulesets: add support for `http_config_settings` phase and supporting actions ([#1036](https://github.com/cloudflare/cloudflare-go/issues/1036))
* workers-account-settings: Add in support for Workers account settings API ([#1027](https://github.com/cloudflare/cloudflare-go/issues/1027))
* workers-subdomain: Add in support Workers Subdomain API ([#1031](https://github.com/cloudflare/cloudflare-go/issues/1031))
* workers-tail: Add in support for Workers tail API ([#1026](https://github.com/cloudflare/cloudflare-go/issues/1026))
* workers: Add support for attaching a worker to a domain ([#1014](https://github.com/cloudflare/cloudflare-go/issues/1014))
* workers: Add support to upload module workers ([#1010](https://github.com/cloudflare/cloudflare-go/issues/1010))

BUG FIXES:

* email_routing_destination: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* email_routing_rules: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* email_routing_settings: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
* tunnel_routes: Fix not removing route when it contains virtual network ([#1030](https://github.com/cloudflare/cloudflare-go/issues/1030))
* workers_test: Fix incorrect test from PR #1014 ([#1048](https://github.com/cloudflare/cloudflare-go/issues/1048))
* workers_test: Use application/json mime-type in headers ([#1049](https://github.com/cloudflare/cloudflare-go/issues/1049))

DEPENDENCIES:

* deps: bumps golang.org/x/tools/gopls from 0.9.3 to 0.9.4 ([#1044](https://github.com/cloudflare/cloudflare-go/issues/1044))
* deps: bumps github.com/golangci/golangci-lint from 1.47.3 to 1.48.0 ([#1020](https://github.com/cloudflare/cloudflare-go/issues/1020))
* deps: bumps github.com/urfave/cli/v2 from 2.11.1 to 2.11.2 ([#1042](https://github.com/cloudflare/cloudflare-go/issues/1042))
* deps: bumps golang.org/x/tools/gopls from 0.9.1 to 0.9.2 ([#1037](https://github.com/cloudflare/cloudflare-go/issues/1037))
* deps: bumps golang.org/x/tools/gopls from 0.9.2 to 0.9.3 ([#1039](https://github.com/cloudflare/cloudflare-go/issues/1039))

## 0.46.0 (3rd August, 2022)

NOTES:

* docs: add release notes ([#1001](https://github.com/cloudflare/cloudflare-go/issues/1001))

ENHANCEMENTS:

* filter: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* firewall_rule: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
* rulesets: add support for `http_custom_errors` phase ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))
* rulesets: add support for `serve_error` action ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))

BUG FIXES:

* access_application: fix inability to set bool values to false ([#1006](https://github.com/cloudflare/cloudflare-go/issues/1006))
* rulesets: fix sni action parameter ([#1002](https://github.com/cloudflare/cloudflare-go/issues/1002))

DEPENDENCIES:

* provider: bumps github.com/golangci/golangci-lint from 1.47.1 to 1.47.2 ([#1005](https://github.com/cloudflare/cloudflare-go/issues/1005))
* provider: bumps github.com/golangci/golangci-lint from 1.47.2 to 1.47.3 ([#1008](https://github.com/cloudflare/cloudflare-go/issues/1008))
* provider: bumps github.com/urfave/cli/v2 from 2.11.0 to 2.11.1 ([#1003](https://github.com/cloudflare/cloudflare-go/issues/1003))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
## 0.96.0 (Unreleased)

## 0.95.0 (May 8th, 2024)

ENHANCEMENTS:

* access_application: add support for `policies` array ([#1956](https://github.com/cloudflare/cloudflare-go/issues/1956))
* access_application: add support for `scim_config` ([#1921](https://github.com/cloudflare/cloudflare-go/issues/1921))
* access_policy: add support for reusable policies ([#1956](https://github.com/cloudflare/cloudflare-go/issues/1956))
* dlp: add support for zt risk behavior configuration ([#1887](https://github.com/cloudflare/cloudflare-go/issues/1887))

BUG FIXES:

* access_application: fix scim configuration authentication json marshalling ([#1959](https://github.com/cloudflare/cloudflare-go/issues/1959))

DEPENDENCIES:

* deps: bumps dependabot/fetch-metadata from 2.0.0 to 2.1.0 ([#1839](https://github.com/cloudflare/cloudflare-go/issues/1839))
* deps: bumps github.com/urfave/cli/v2 from 2.27.1 to 2.27.2 ([#1861](https://github.com/cloudflare/cloudflare-go/issues/1861))
* deps: bumps golang.org/x/net from 0.24.0 to 0.25.0 ([#1974](https://github.com/cloudflare/cloudflare-go/issues/1974))
* deps: bumps golangci/golangci-lint-action from 4 to 5 ([#1845](https://github.com/cloudflare/cloudflare-go/issues/1845))
* deps: bumps golangci/golangci-lint-action from 5 to 6 ([#1975](https://github.com/cloudflare/cloudflare-go/issues/1975))

## 0.94.0 (April 24th, 2024)

ENHANCEMENTS:

* access_application: support options_preflight_bypass for access_application ([#1790](https://github.com/cloudflare/cloudflare-go/issues/1790))
* gateway: added ecs_support field to teams_location resource ([#1826](https://github.com/cloudflare/cloudflare-go/issues/1826))
* teams_account: adds custom certificate setting to teams account configuration ([#1811](https://github.com/cloudflare/cloudflare-go/issues/1811))
* workers: support deleting namespaced Workers ([#1737](https://github.com/cloudflare/cloudflare-go/issues/1737))

DEPENDENCIES:

* deps: bumps golang.org/x/net from 0.19.0 to 0.23.0 ([#1825](https://github.com/cloudflare/cloudflare-go/issues/1825))

## 0.93.0 (April 10th, 2024)

BREAKING CHANGES:

* dns: Remove "locked" flag which is always false ([#1618](https://github.com/cloudflare/cloudflare-go/issues/1618))

ENHANCEMENTS:

* magic_transit_ipsec_tunnel: Adds support for replay_protection boolean flag ([#1710](https://github.com/cloudflare/cloudflare-go/issues/1710))

DEPENDENCIES:

* deps: bumps golang.org/x/net from 0.22.0 to 0.24.0 ([#1688](https://github.com/cloudflare/cloudflare-go/issues/1688))

## 0.92.0 (March 27th, 2024)

ENHANCEMENTS:

- dlp: Adds support for ocr_enabled boolean flag ([#1600](https://github.com/cloudflare/cloudflare-go/issues/1600))

BUG FIXES:

- teams_rules: add "resolve" to allowable actions ([#1615](https://github.com/cloudflare/cloudflare-go/issues/1615))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.6.0 to 1.7.0 ([#1593](https://github.com/cloudflare/cloudflare-go/issues/1593))
- deps: bumps dependabot/fetch-metadata from 1.7.0 to 2.0.0 ([#1607](https://github.com/cloudflare/cloudflare-go/issues/1607))

## 0.91.0 (March 22nd, 2024)

ENHANCEMENTS:

- access_application: add support for `saml_attribute_transform_jsonata` in saas apps ([#1562](https://github.com/cloudflare/cloudflare-go/issues/1562))
- dlp: Adds support for ocr_enabled boolean flag ([#1600](https://github.com/cloudflare/cloudflare-go/issues/1600))

BUG FIXES:

- teams_rules: add "resolve" to allowable actions ([#1615](https://github.com/cloudflare/cloudflare-go/issues/1615))

DEPENDENCIES:

- deps: bumps actions/checkout from 2 to 4 ([#1573](https://github.com/cloudflare/cloudflare-go/issues/1573))
- deps: bumps dependabot/fetch-metadata from 1.6.0 to 1.7.0 ([#1593](https://github.com/cloudflare/cloudflare-go/issues/1593))
- deps: bumps dependabot/fetch-metadata from 1.7.0 to 2.0.0 ([#1607](https://github.com/cloudflare/cloudflare-go/issues/1607))
- deps: bumps google.golang.org/protobuf from 1.28.0 to 1.33.0 ([#1558](https://github.com/cloudflare/cloudflare-go/issues/1558))

## 0.90.0 (March 13th, 2024)

ENHANCEMENTS:

- access_mutual_tls_certificates: add support for mutual tls hostname settings ([#1516](https://github.com/cloudflare/cloudflare-go/issues/1516))
- device_posture_rule: support last_seen and state for crowdstrike_s2s posture rule ([#1509](https://github.com/cloudflare/cloudflare-go/issues/1509))
- dlp: add support for Context Awareness in DLP profiles ([#1497](https://github.com/cloudflare/cloudflare-go/issues/1497))
- workers: Add Workers for Platforms support for getting a Worker, content and bindings ([#1508](https://github.com/cloudflare/cloudflare-go/issues/1508))
- workers_for_platforms: Add ability to list Workers for Platforms namespaces, get a namespace, create a new namespace or delete a namespace. ([#1508](https://github.com/cloudflare/cloudflare-go/issues/1508))

BUG FIXES:

- dlp: added optional ContextAwareness support ([#1510](https://github.com/cloudflare/cloudflare-go/issues/1510))

DEPENDENCIES:

- deps: bumps github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#1511](https://github.com/cloudflare/cloudflare-go/issues/1511))
- deps: bumps golang.org/x/net from 0.21.0 to 0.22.0 ([#1513](https://github.com/cloudflare/cloudflare-go/issues/1513))

## 0.89.0 (February 28th, 2024)

NOTES:

- zaraz: replace deprecated neoEvents with Actions on Zaraz Config tools schema ([#1490](https://github.com/cloudflare/cloudflare-go/issues/1490))

ENHANCEMENTS:

- magic-transit: Adds IPsec tunnel healthcheck direction & rate parameters ([#1503](https://github.com/cloudflare/cloudflare-go/issues/1503))

BUG FIXES:

- registrar: Fix request method to call domain list endpoint from POST to GET ([#1506](https://github.com/cloudflare/cloudflare-go/issues/1506))

## 0.88.0 (February 14th, 2024)

ENHANCEMENTS:

- access_application: Add support for OIDC SaaS Applications ([#1500](https://github.com/cloudflare/cloudflare-go/issues/1500))
- access_application: Add support for `allow_authenticate_via_warp` ([#1496](https://github.com/cloudflare/cloudflare-go/issues/1496))
- access_application: add support for `name_id_transform_jsonata` in saas apps ([#1505](https://github.com/cloudflare/cloudflare-go/issues/1505))
- access_organization: Add support for `allow_authenticate_via_warp` and `warp_auth_session_duration` ([#1496](https://github.com/cloudflare/cloudflare-go/issues/1496))
- hyperdrive: Add support for hyperdrive CRUD operations ([#1492](https://github.com/cloudflare/cloudflare-go/issues/1492))
- images_variants: Add support for Images Variants CRUD operations ([#1494](https://github.com/cloudflare/cloudflare-go/issues/1494))
- teams_rules: `AntiVirus` settings includes notification settings ([#1499](https://github.com/cloudflare/cloudflare-go/issues/1499))

BUG FIXES:

- hyperdrive: password should be nested in origin ([#1501](https://github.com/cloudflare/cloudflare-go/issues/1501))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.20.0 to 0.21.0 ([#1502](https://github.com/cloudflare/cloudflare-go/issues/1502))
- deps: bumps golangci/golangci-lint-action from 3 to 4 ([#1504](https://github.com/cloudflare/cloudflare-go/issues/1504))

## 0.87.0 (January 31st, 2024)

ENHANCEMENTS:

- access_seats: Add `UpdateAccessUsersSeats` with an array as input for multiple operations ([#1480](https://github.com/cloudflare/cloudflare-go/issues/1480))
- dlp: add support for EDM and CWL datasets ([#1485](https://github.com/cloudflare/cloudflare-go/issues/1485))
- logpush: Add support for Output Options ([#1468](https://github.com/cloudflare/cloudflare-go/issues/1468))
- pages_project: Add `build_caching` attribute ([#1489](https://github.com/cloudflare/cloudflare-go/issues/1489))
- streams: adds support for stream create parameters for tus upload initiate ([#1386](https://github.com/cloudflare/cloudflare-go/issues/1386))
- teams_accounts: add support for extended email matching ([#1486](https://github.com/cloudflare/cloudflare-go/issues/1486))

BUG FIXES:

- access_seats: UpdateAccessUserSeat: fix parameters not being an array when sending to the api. This caused an error when updating a user's seat ([#1480](https://github.com/cloudflare/cloudflare-go/issues/1480))
- access_users: ListAccessUsers was returning wrong values in pointer fields due to variable missused in loop ([#1482](https://github.com/cloudflare/cloudflare-go/issues/1482))
- flarectl: alias zone certs to "ct" instead of duplicating the "c" alias ([#1484](https://github.com/cloudflare/cloudflare-go/issues/1484))

DEPENDENCIES:

- deps: bumps actions/cache from 3 to 4 ([#1483](https://github.com/cloudflare/cloudflare-go/issues/1483))

## 0.86.0 (January 17, 2024)

ENHANCEMENTS:

- access_application: Add support for default_relay_state in saas apps ([#1477](https://github.com/cloudflare/cloudflare-go/issues/1477))
- zaraz: Add support for CRUD APIs ([#1474](https://github.com/cloudflare/cloudflare-go/issues/1474))

DEPENDENCIES:

- deps: bumps github.com/cloudflare/circl from 1.3.3 to 1.3.7 ([#1475](https://github.com/cloudflare/cloudflare-go/issues/1475))
- deps: bumps golang.org/x/net from 0.19.0 to 0.20.0 ([#1476](https://github.com/cloudflare/cloudflare-go/issues/1476))

## 0.85.0 (January 3rd, 2024)

DEPENDENCIES:

- deps: bumps github.com/go-git/go-git/v5 from 5.4.2 to 5.11.0 ([#1470](https://github.com/cloudflare/cloudflare-go/issues/1470))
- deps: bumps github.com/urfave/cli/v2 from 2.26.0 to 2.27.0 ([#1471](https://github.com/cloudflare/cloudflare-go/issues/1471))
- deps: bumps github.com/urfave/cli/v2 from 2.27.0 to 2.27.1 ([#1472](https://github.com/cloudflare/cloudflare-go/issues/1472))

## 0.84.0 (December 20th, 2023)

ENHANCEMENTS:

- access_group: Add support for email lists ([#1445](https://github.com/cloudflare/cloudflare-go/issues/1445))
- device_posture_rules: add support for Access client fields in device posture integrations ([#1464](https://github.com/cloudflare/cloudflare-go/issues/1464))
- page_shield: added support for page shield ([#1459](https://github.com/cloudflare/cloudflare-go/issues/1459))

DEPENDENCIES:

- deps: bumps actions/setup-go from 4 to 5 ([#1460](https://github.com/cloudflare/cloudflare-go/issues/1460))
- deps: bumps github/codeql-action from 2 to 3 ([#1462](https://github.com/cloudflare/cloudflare-go/issues/1462))
- deps: bumps golang.org/x/crypto from 0.14.0 to 0.17.0 ([#1466](https://github.com/cloudflare/cloudflare-go/issues/1466))

## 0.83.0 (December 6th, 2023)

ENHANCEMENTS:

- cloudflare: Add ResultInfo to RawResponse ([#1453](https://github.com/cloudflare/cloudflare-go/issues/1453))
- devices_policy: add fields for Opt-In Split Tunnel Overlapping IPs feature. ([#1454](https://github.com/cloudflare/cloudflare-go/issues/1454))
- stream: Add ScheduledDeletion to StreamCreateVideoParameters ([#1457](https://github.com/cloudflare/cloudflare-go/issues/1457))
- stream: Add ScheduledDeletion to StreamUploadFromURLParameters ([#1457](https://github.com/cloudflare/cloudflare-go/issues/1457))
- stream: Add ScheduledDeletion to StreamVideo ([#1457](https://github.com/cloudflare/cloudflare-go/issues/1457))
- stream: Add ScheduledDeletion to StreamVideoCreate ([#1457](https://github.com/cloudflare/cloudflare-go/issues/1457))
- worker_bindings: Fixing form element name for d1 binding ([#1450](https://github.com/cloudflare/cloudflare-go/issues/1450))
- worker_bindings: add support for `d1` bindings ([#1446](https://github.com/cloudflare/cloudflare-go/issues/1446))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.25.7 to 2.26.0 ([#1456](https://github.com/cloudflare/cloudflare-go/issues/1456))
- deps: bumps golang.org/x/net from 0.18.0 to 0.19.0 ([#1452](https://github.com/cloudflare/cloudflare-go/issues/1452))
- deps: bumps golang.org/x/time from 0.4.0 to 0.5.0 ([#1449](https://github.com/cloudflare/cloudflare-go/issues/1449))

## 0.82.0 (November 22nd, 2023)

ENHANCEMENTS:

- ip_access_rules: Add ListIPAccessRules() to list IP Access Rules ([#1428](https://github.com/cloudflare/cloudflare-go/issues/1428))
- load_balancing: add healthy field to LoadBalancerPool ([#1442](https://github.com/cloudflare/cloudflare-go/issues/1442))

BUG FIXES:

- load_balancing: Add support for virtual network id in origins ([#1441](https://github.com/cloudflare/cloudflare-go/issues/1441))
- per_hostname_tls_setting: use `buildURI` for defining the query parameters when sorting ([#1440](https://github.com/cloudflare/cloudflare-go/issues/1440))

DEPENDENCIES:

- deps: bumps github.com/hashicorp/go-retryablehttp from 0.7.4 to 0.7.5 ([#1438](https://github.com/cloudflare/cloudflare-go/issues/1438))
- deps: bumps golang.org/x/net from 0.17.0 to 0.18.0 ([#1439](https://github.com/cloudflare/cloudflare-go/issues/1439))

## 0.81.0 (November 8th, 2023)

BREAKING CHANGES:

- devices_policy: `CreateDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `DeleteDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `DeviceClientCertificates` is renamed to `DeviceClientCertificates` ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `GetDefaultDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `GetDeviceClientCertificatesZone` is renamed to `GetDeviceClientCertificates` with updated method signatures ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `GetDeviceClientCertificates` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `GetDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `UpdateDefaultDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `UpdateDeviceClientCertificatesZone` is renamed to `UpdateDeviceClientCertificates` with updated method signatures ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- devices_policy: `UpdateDeviceSettingsPolicy` is updated with method signatures matching the library conventions ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))

ENHANCEMENTS:

- access_seats: Add UpdateAccessUserSeat() to list IP Access Rules ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- access_user: Add GetAccessUserActiveSessions() to get all active sessions for a Access/Zero-Trust user. ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- access_user: Add GetAccessUserFailedLogins() to get all failed login attempts for a Access/Zero-Trust user. ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- access_user: Add GetAccessUserLastSeenIdentity() to get last seen identity for a Access/Zero-Trust user. ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- access_user: Add GetAccessUserSingleActiveSession() to get an active session for a Access/Zero-Trust user. ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- access_user: Add ListAccessUsers() to get a list of users for a Access/Zero-Trust account. ([#1427](https://github.com/cloudflare/cloudflare-go/issues/1427))
- devices_policy: Add support for listing device settings policies ([#1433](https://github.com/cloudflare/cloudflare-go/issues/1433))
- teams_rules: Add support for resolver policies ([#1436](https://github.com/cloudflare/cloudflare-go/issues/1436))

DEPENDENCIES:

- deps: bumps golang.org/x/time from 0.3.0 to 0.4.0 ([#1434](https://github.com/cloudflare/cloudflare-go/issues/1434))

## 0.80.0 (October 25th, 2023)

BREAKING CHANGES:

- teams: `BrowserIsolation.UrlBrowserIsolationEnabled` has changed from `bool` to `*bool` to meet the library conventions ([#1424](https://github.com/cloudflare/cloudflare-go/issues/1424))

ENHANCEMENTS:

- access_application: Add support for app launcher customization fields ([#1407](https://github.com/cloudflare/cloudflare-go/issues/1407))
- api_shield_schema: Add support for Get/Update API Shield Operation Schema Validation Settings ([#1422](https://github.com/cloudflare/cloudflare-go/issues/1422))
- api_shield_schema: Add support for Get/Update API Shield Schema Validation Settings ([#1418](https://github.com/cloudflare/cloudflare-go/issues/1418))
- teams: Add support for body_scanning (Enhanced File Detection) in teams account configuration ([#1423](https://github.com/cloudflare/cloudflare-go/issues/1423))
- load_balancing: extend documentation for least_connections steering policy ([#1414](https://github.com/cloudflare/cloudflare-go/issues/1414))
- teams: Add `non_identity_enabled` boolean in browser isolation settings ([#1424](https://github.com/cloudflare/cloudflare-go/issues/1424))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.7.0 to 0.17.0 ([#1421](https://github.com/cloudflare/cloudflare-go/issues/1421))

## 0.79.0 (October 11th, 2023)

ENHANCEMENTS:

- access_organization: Add support for session_duration ([#1415](https://github.com/cloudflare/cloudflare-go/issues/1415))
- access_policy: Add support for session_duration ([#1415](https://github.com/cloudflare/cloudflare-go/issues/1415))

ENHANCEMENTS:

- api_shield_discovery: Add support for Get/Patch API Shield API Discovery Operations ([#1413](https://github.com/cloudflare/cloudflare-go/issues/1413))
- api_shield_schema: Add support for managing schemas for API Shield Schema Validation 2.0 ([#1406](https://github.com/cloudflare/cloudflare-go/issues/1406))
- d1: adds support for d1 ([#1417](https://github.com/cloudflare/cloudflare-go/issues/1417))
- teams: Add `audit_ssh_settings` endpoints ([#1419](https://github.com/cloudflare/cloudflare-go/issues/1419))

BUG FIXES:

- custom_nameservers: change `NSSet` from string to int to match API response ([#1410](https://github.com/cloudflare/cloudflare-go/issues/1410))
- observatory: fix double url encoding ([#1412](https://github.com/cloudflare/cloudflare-go/issues/1412))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.15.0 to 0.16.0 ([#1416](https://github.com/cloudflare/cloudflare-go/issues/1416))
- deps: bumps golang.org/x/net from 0.16.0 to 0.17.0 ([#1420](https://github.com/cloudflare/cloudflare-go/issues/1420))

## 0.78.0 (September 27th, 2023)

BREAKING CHANGES:

- account_role: `AccountRole` has been renamed to `GetAccountRole` to align with the updated method conventions ([#1405](https://github.com/cloudflare/cloudflare-go/issues/1405))
- account_role: `AccountRoles` has been renamed to `ListAccountRoles` to align with the updated method conventions ([#1405](https://github.com/cloudflare/cloudflare-go/issues/1405))

ENHANCEMENTS:

- access_application: Add support for tags ([#1403](https://github.com/cloudflare/cloudflare-go/issues/1403))
- access_tag: Add support for tags ([#1403](https://github.com/cloudflare/cloudflare-go/issues/1403))
- list_item: allow filtering by search term, cursor and per page attributes ([#1409](https://github.com/cloudflare/cloudflare-go/issues/1409))
- observatory: add support for observatory API ([#1401](https://github.com/cloudflare/cloudflare-go/issues/1401))

BUG FIXES:

- account_role: autopaginate all available results instead of a static number ([#1405](https://github.com/cloudflare/cloudflare-go/issues/1405))
- semgrep: Improved IPv4 validation by implementing a new pattern to handle cases where non-IPv4 addresses were previously accepted. ([#1382](https://github.com/cloudflare/cloudflare-go/issues/1382))

DEPENDENCIES:

- deps: bumps codecov/codecov-action from 3 to 4 ([#1402](https://github.com/cloudflare/cloudflare-go/issues/1402))

## 0.77.0 (September 13th, 2023)

ENHANCEMENTS:

- access_identity_provider: add support for email_claim_name and authorization_server_id ([#1390](https://github.com/cloudflare/cloudflare-go/issues/1390))
- access_identity_provider: add support for ping_env_id ([#1391](https://github.com/cloudflare/cloudflare-go/issues/1391))
- dcv_delegation: add GET for DCV Delegation UUID ([#1384](https://github.com/cloudflare/cloudflare-go/issues/1384))
- streams: adds support to initiate tus upload ([#1359](https://github.com/cloudflare/cloudflare-go/issues/1359))
- tunnel: add support for `include_prefix`, `exclude_prefix` in list operations ([#1385](https://github.com/cloudflare/cloudflare-go/issues/1385))

BUG FIXES:

- dns: keep comments when calling UpdateDNSRecord with zero values of UpdateDNSRecordParams ([#1393](https://github.com/cloudflare/cloudflare-go/issues/1393))

DEPENDENCIES:

- deps: bumps actions/checkout from 3 to 4 ([#1387](https://github.com/cloudflare/cloudflare-go/issues/1387))
- deps: bumps golang.org/x/net from 0.14.0 to 0.15.0 ([#1389](https://github.com/cloudflare/cloudflare-go/issues/1389))
- deps: bumps goreleaser/goreleaser-action from 4.4.0 to 4.6.0 ([#1388](https://github.com/cloudflare/cloudflare-go/issues/1388))
- deps: bumps goreleaser/goreleaser-action from 4.6.0 to 5.0.0 ([#1396](https://github.com/cloudflare/cloudflare-go/issues/1396))

## 0.76.0 (August 30th, 2023)

BREAKING CHANGES:

- images: Renamed Image struct "Metadata" field to "Meta" ([#1379](https://github.com/cloudflare/cloudflare-go/issues/1379))

ENHANCEMENTS:

- access_application: added custom_non_identity_deny_url ([#1373](https://github.com/cloudflare/cloudflare-go/issues/1373))
- load_balancer_monitor: add support for `consecutive_up`, `consecutive_down` ([#1380](https://github.com/cloudflare/cloudflare-go/issues/1380))
- workers: Add support for retrieving and uploading only script content. ([#1361](https://github.com/cloudflare/cloudflare-go/issues/1361))
- workers: Add support for retrieving and uploading only script metadata. ([#1361](https://github.com/cloudflare/cloudflare-go/issues/1361))
- workers: allow namespaced scripts to be used as Worker tail consumers ([#1377](https://github.com/cloudflare/cloudflare-go/issues/1377))

BUG FIXES:

- access_application: Use autopaginate flag as expected ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))
- access_ca_certificate: Use autopaginate flag as expected ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))
- access_group: Use autopaginate flag as expected ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))
- access_mutual_tls_certifcate: Use autopaginate flag as expected ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))
- access_policy: Use autopaginate flag as expected ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))
- images: Fix issue parsing Image Details from API due to incorrect struct json field ([#1379](https://github.com/cloudflare/cloudflare-go/issues/1379))
- pagination: Will look at `total_count` and `per_page` to calculate `total_pages` if `total_pages` is zero ([#1372](https://github.com/cloudflare/cloudflare-go/issues/1372))

## 0.75.0 (August 16th, 2023)

BREAKING CHANGES:

- cloudflare: `Raw` method now returns a RawResponse rather than the raw JSON `Result` message ([#1355](https://github.com/cloudflare/cloudflare-go/issues/1355))
- rulesets: Rename `RulesetPhaseRateLimit` to `RulesetPhaseHTTPRatelimit`, to match the phase name ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))
- rulesets: Rename `RulesetPhaseSuperBotFightMode` to `RulesetPhaseHTTPRequestSBFM`, to match the phase name ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))

NOTES:

- rulesets: Remove non-existent `allow` action ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))
- rulesets: Remove non-existent `http_request_main` phase ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))
- rulesets: Remove non-public `http_response_headers_transform_managed` and `http_request_late_transform_managed` phases ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))

ENHANCEMENTS:

- access_group: add auth_context group ruletype ([#1344](https://github.com/cloudflare/cloudflare-go/issues/1344))
- access_identity_provider: add attr conditional_access_enabled ([#1344](https://github.com/cloudflare/cloudflare-go/issues/1344))
- access_identity_provider: add auth context list/put endpoint ([#1344](https://github.com/cloudflare/cloudflare-go/issues/1344))
- access_service_token: add support for managing `Duration` ([#1347](https://github.com/cloudflare/cloudflare-go/issues/1347))
- bot_management: add support for bot_management API ([#1363](https://github.com/cloudflare/cloudflare-go/issues/1363))
- cloudflare: swap `encoding/json` for `github.com/goccy/go-json` ([#1360](https://github.com/cloudflare/cloudflare-go/issues/1360))
- device_posture_rule: support eid_last_seen and risk_level and correct total_score for Tanium posture rule ([#1366](https://github.com/cloudflare/cloudflare-go/issues/1366))
- per_hostname_tls_settings: add support for managing hostname level TLS settings ([#1356](https://github.com/cloudflare/cloudflare-go/issues/1356))
- rulesets: Add the `ddos_mitigation` action ([#1367](https://github.com/cloudflare/cloudflare-go/issues/1367))
- waiting_room: add support for `queueing_status_code` ([#1357](https://github.com/cloudflare/cloudflare-go/issues/1357))
- web_analytics: add support for web_analytics API ([#1348](https://github.com/cloudflare/cloudflare-go/issues/1348))
- workers: add support for tagging Worker scripts ([#1368](https://github.com/cloudflare/cloudflare-go/issues/1368))
- zone_hold: add support for zone hold API ([#1365](https://github.com/cloudflare/cloudflare-go/issues/1365))

BUG FIXES:

- cache_purge: don't escape HTML entity values in URLs for cache keys ([#1360](https://github.com/cloudflare/cloudflare-go/issues/1360))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.12.0 to 0.13.0 ([#1353](https://github.com/cloudflare/cloudflare-go/issues/1353))
- deps: bumps golang.org/x/net from 0.13.0 to 0.14.0 ([#1362](https://github.com/cloudflare/cloudflare-go/issues/1362))
- deps: bumps goreleaser/goreleaser-action from 4.3.0 to 4.4.0 ([#1369](https://github.com/cloudflare/cloudflare-go/issues/1369))

## 0.74.0 (August 2nd, 2023)

ENHANCEMENTS:

- access_application: Add support for custom pages ([#1343](https://github.com/cloudflare/cloudflare-go/issues/1343))
- access_custom_page: Add support for custom pages ([#1343](https://github.com/cloudflare/cloudflare-go/issues/1343))
- access_organization: add support for custom pages ([#1343](https://github.com/cloudflare/cloudflare-go/issues/1343))
- rulesets: Remove internal-only schema kind ([#1346](https://github.com/cloudflare/cloudflare-go/issues/1346))
- rulesets: Remove some request parameters that are not allowed or have no effect ([#1346](https://github.com/cloudflare/cloudflare-go/issues/1346))
- rulesets: Update API reference links ([#1346](https://github.com/cloudflare/cloudflare-go/issues/1346))
- teams-accounts: Adds support for protocol detection ([#1340](https://github.com/cloudflare/cloudflare-go/issues/1340))
- workers: Add `pipeline_hash` field to Workers script response struct. ([#1330](https://github.com/cloudflare/cloudflare-go/issues/1330))
- workers: Add support for declaring arbitrary bindings with UnsafeBinding. ([#1330](https://github.com/cloudflare/cloudflare-go/issues/1330))
- workers: Add support for uploading scripts to a Workers for Platforms namespace. ([#1330](https://github.com/cloudflare/cloudflare-go/issues/1330))
- workers: Add support for uploading workers with Workers for Platforms namespace bindings. ([#1330](https://github.com/cloudflare/cloudflare-go/issues/1330))

BUG FIXES:

- flarectl: allow for create or update to actually create the record ([#1341](https://github.com/cloudflare/cloudflare-go/issues/1341))
- load_balancing: Fix pool creation with MinimumOrigins set to 0 ([#1338](https://github.com/cloudflare/cloudflare-go/issues/1338))
- workers: Fix namespace dispatch upload API path ([#1345](https://github.com/cloudflare/cloudflare-go/issues/1345))

## 0.73.0 (July 19th, 2023)

BREAKING CHANGES:

- pages_deployment: add support for auto pagination ([#1264](https://github.com/cloudflare/cloudflare-go/issues/1264))
- pages_deployment: change DeletePagesDeploymentParams to contain all parameters ([#1264](https://github.com/cloudflare/cloudflare-go/issues/1264))
- pages_project: change to use ResourceContainer for account ID ([#1264](https://github.com/cloudflare/cloudflare-go/issues/1264))
- pages_project: rename PagesProject to GetPagesProject ([#1264](https://github.com/cloudflare/cloudflare-go/issues/1264))
- rulesets: `CreateAccountRuleset` is removed in favour of `CreateRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `CreateZoneRuleset` is removed in favour of `CreateRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `DeleteAccountRuleset` is removed in favour of `DeleteRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `DeleteZoneRuleset` is removed in favour of `DeleteRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `GetAccountRulesetPhase` is removed in favour of `GetEntrypointRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `GetAccountRuleset` is removed in favour of `GetRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `GetZoneRulesetPhase` is removed in favour of `GetEntrypointRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `GetZoneRuleset` is removed in favour of `GetRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `UpdateAccountRulesetPhase` is removed in favour of `UpdateEntrypointRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `UpdateAccountRuleset` is removed in favour of `UpdateRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `UpdateZoneRulesetPhase` is removed in favour of `UpdateEntrypointRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))
- rulesets: `UpdateZoneRuleset` is removed in favour of `UpdateRuleset` ([#1333](https://github.com/cloudflare/cloudflare-go/issues/1333))

ENHANCEMENTS:

- device_posture_rule: support active_threats, network_status, infected, and is_active for sentinelone_s2s posture rule ([#1339](https://github.com/cloudflare/cloudflare-go/issues/1339))
- device_posture_rule: support certificate_id and cn for client_certificate posture rule ([#1339](https://github.com/cloudflare/cloudflare-go/issues/1339))
- images: adds ability to upload image by url ([#1335](https://github.com/cloudflare/cloudflare-go/issues/1335))
- load_balancing: support header session affinity policy ([#1302](https://github.com/cloudflare/cloudflare-go/issues/1302))
- zone: Added `GetRegionalTieredCache` and `UpdateRegionalTieredCache` to allow setting Regional Tiered Cache for a zone. ([#1336](https://github.com/cloudflare/cloudflare-go/issues/1336))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.11.0 to 0.12.0 ([#1328](https://github.com/cloudflare/cloudflare-go/issues/1328))

## 0.72.0 (July 5th, 2023)

BREAKING CHANGES:

- logpush: `CheckAccountLogpushDestinationExists` is removed in favour of `CheckLogpushDestinationExists` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `CheckZoneLogpushDestinationExists` is removed in favour of `CheckLogpushDestinationExists` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `CreateAccountLogpushJob` is removed in favour of `CreateLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `CreateZoneLogpushJob` is removed in favour of `CreateLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `DeleteAccountLogpushJob` is removed in favour of `DeleteLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `DeleteZoneLogpushJob` is removed in favour of `DeleteLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetAccountLogpushFields` is removed in favour of `GetLogpushFields` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetAccountLogpushJob` is removed in favour of `GetLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetAccountLogpushOwnershipChallenge` is removed in favour of `GetLogpushOwnershipChallenge` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetZoneLogpushFields` is removed in favour of `GetLogpushFields` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetZoneLogpushJob` is removed in favour of `GetLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `GetZoneLogpushOwnershipChallenge` is removed in favour of `GetLogpushOwnershipChallenge` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ListAccountLogpushJobsForDataset` is removed in favour of `ListLogpushJobsForDataset` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ListAccountLogpushJobs` is removed in favour of `ListLogpushJobs` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ListZoneLogpushJobsForDataset` is removed in favour of `ListLogpushJobsForDataset` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ListZoneLogpushJobs` is removed in favour of `ListLogpushJobs` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `UpdateAccountLogpushJob` is removed in favour of `UpdateLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `UpdateZoneLogpushJob` is removed in favour of `UpdateLogpushJob` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ValidateAccountLogpushOwnershipChallenge` is removed in favour of `ValidateLogpushOwnershipChallenge` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: `ValidateZoneLogpushOwnershipChallenge` is removed in favour of `ValidateLogpushOwnershipChallenge` with `ResourceContainer` method parameter ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))
- logpush: all methods are updated to use the newer client conventions for method signatures ([#1326](https://github.com/cloudflare/cloudflare-go/issues/1326))

ENHANCEMENTS:

- resource_container: expose `Type` on `*ResourceContainer` to explicitly denote what type of resource it is instead of inferring from `Level`. ([#1325](https://github.com/cloudflare/cloudflare-go/issues/1325))

## 0.71.0 (July 5th, 2023)

BREAKING CHANGES:

- access_application: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_ca_certificate: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_group: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_identity_provider: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_mutual_tls_certificates: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_organization: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_policy: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_service_tokens: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_user_token: refactor methods to use `ResourceContainer` instead of dedicated account/zone methods ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- images: renamed `BaseImage` to `GetBaseImage` to match library conventions ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: renamed `ImageDetails` to `GetImage` to match library conventions ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: renamed `ImagesStats` to `GetImagesStats` to match library conventions ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: updated method signatures of `DeleteImage` to match newer conventions and standards ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: updated method signatures of `ListImages` to match newer conventions and standards ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: updated method signatures of `UpdateImage` to match newer conventions and standards ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- images: updated method signatures of `UploadImage` to match newer conventions and standards ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))

ENHANCEMENTS:

- access_application: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_ca_certificate: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_group: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_identity_provider: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_mutual_tls_certificates: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- access_policy: add support for auto pagination ([#1319](https://github.com/cloudflare/cloudflare-go/issues/1319))
- device_posture_rule: support os_version_extra ([#1316](https://github.com/cloudflare/cloudflare-go/issues/1316))
- images: adds support for v2 when uploading images directly ([#1322](https://github.com/cloudflare/cloudflare-go/issues/1322))
- workers: Add ability to specify tail Workers in script metadata ([#1317](https://github.com/cloudflare/cloudflare-go/issues/1317))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.5.1 to 1.6.0 ([#1320](https://github.com/cloudflare/cloudflare-go/issues/1320))

## 0.70.0 (June 21st, 2023)

BREAKING CHANGES:

- cloudflare: remove `UsingAccount` in favour of resource specific attributes ([#1315](https://github.com/cloudflare/cloudflare-go/issues/1315))
- cloudflare: remove `api.AccountID` from client struct ([#1315](https://github.com/cloudflare/cloudflare-go/issues/1315))
- dns_firewall: modernise method signatures and conventions to align with the experimental client ([#1313](https://github.com/cloudflare/cloudflare-go/issues/1313))
- railgun: remove support for railgun ([#1312](https://github.com/cloudflare/cloudflare-go/issues/1312))
- tunnel: swap `ConnectTimeout`, `TLSTimeout`, `TCPKeepAlive` and `KeepAliveTimeout` to `TunnelDuration` instead of `time.Duration` ([#1303](https://github.com/cloudflare/cloudflare-go/issues/1303))
- virtualdns: remove support in favour of newer DNS firewall methods ([#1313](https://github.com/cloudflare/cloudflare-go/issues/1313))

ENHANCEMENTS:

- custom_nameservers: add support for managing custom nameservers ([#1304](https://github.com/cloudflare/cloudflare-go/issues/1304))
- load_balancing: extend documentation for least_outstanding_requests steering policy ([#1293](https://github.com/cloudflare/cloudflare-go/issues/1293))
- waiting_room: add support for `additional_routes` and `cookie_suffix` ([#1311](https://github.com/cloudflare/cloudflare-go/issues/1311))

DEPENDENCIES:

- deps: bumps github.com/hashicorp/go-retryablehttp from 0.7.3 to 0.7.4 ([#1301](https://github.com/cloudflare/cloudflare-go/issues/1301))
- deps: bumps github.com/urfave/cli/v2 from 2.25.5 to 2.25.6 ([#1305](https://github.com/cloudflare/cloudflare-go/issues/1305))
- deps: bumps github.com/urfave/cli/v2 from 2.25.6 to 2.25.7 ([#1314](https://github.com/cloudflare/cloudflare-go/issues/1314))
- deps: bumps golang.org/x/net from 0.10.0 to 0.11.0 ([#1307](https://github.com/cloudflare/cloudflare-go/issues/1307))
- deps: bumps goreleaser/goreleaser-action from 4.2.0 to 4.3.0 ([#1306](https://github.com/cloudflare/cloudflare-go/issues/1306))

## 0.69.0 (June 7th, 2023)

BREAKING CHANGES:

- stream: StreamVideo.Duration has changed from int to float64. ([#1190](https://github.com/cloudflare/cloudflare-go/issues/1190))

ENHANCEMENTS:

- access: Added `self_hosted_domains` support to access applications ([#1281](https://github.com/cloudflare/cloudflare-go/issues/1281))
- custom_hostname: add support for `bundle_method` TLS configuration ([#1298](https://github.com/cloudflare/cloudflare-go/issues/1298))
- devices_policy: Add missing description field to policy ([#1294](https://github.com/cloudflare/cloudflare-go/issues/1294))
- stream: added metadata support ([#1088](https://github.com/cloudflare/cloudflare-go/issues/1088))

BUG FIXES:

- email_routing_destination: return encountered error, not `ErrMissingAccountID` all the time ([#1297](https://github.com/cloudflare/cloudflare-go/issues/1297))
- stream: Fix a bug that cannot unmarshal video duration number. ([#1190](https://github.com/cloudflare/cloudflare-go/issues/1190))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.5.0 to 1.5.1 ([#1292](https://github.com/cloudflare/cloudflare-go/issues/1292))
- deps: bumps github.com/hashicorp/go-retryablehttp from 0.7.2 to 0.7.3 ([#1300](https://github.com/cloudflare/cloudflare-go/issues/1300))
- deps: bumps github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#1296](https://github.com/cloudflare/cloudflare-go/issues/1296))
- deps: bumps github.com/urfave/cli/v2 from 2.25.3 to 2.25.5 ([#1295](https://github.com/cloudflare/cloudflare-go/issues/1295))

## 0.68.0 (May 24th, 2023)

BREAKING CHANGES:

- r2_bucket: change creation time from string to \*time.Time ([#1265](https://github.com/cloudflare/cloudflare-go/issues/1265))

ENHANCEMENTS:

- adds OriginRequest field to UnvalidatedIngressRule struct. ([#1138](https://github.com/cloudflare/cloudflare-go/issues/1138))
- lists: add support for hostname and ASN lists. ([#1288](https://github.com/cloudflare/cloudflare-go/issues/1288))
- pages: add support for Smart Placement. Added `Placement` in `PagesProjectDeploymentConfigEnvironment`. ([#1279](https://github.com/cloudflare/cloudflare-go/issues/1279))
- r2_bucket: add support for getting a bucket ([#1265](https://github.com/cloudflare/cloudflare-go/issues/1265))
- tunnels: add support for `access` and `http2Origin` keys ([#1291](https://github.com/cloudflare/cloudflare-go/issues/1291))
- workers: add support for Smart Placement. Added `Placement` in `CreateWorkerParams`. ([#1279](https://github.com/cloudflare/cloudflare-go/issues/1279))
- zone: Added `GetCacheReserve` and `UpdateacheReserve` to allow setting Cache Reserve for a zone. ([#1278](https://github.com/cloudflare/cloudflare-go/issues/1278))

BUG FIXES:

- dns: fix MX record priority not set by UpdateDNSRecord ([#1290](https://github.com/cloudflare/cloudflare-go/issues/1290))
- flarectl/dns: ensure MX priority value is dereferenced ([#1289](https://github.com/cloudflare/cloudflare-go/issues/1289))
- turnstile: remove `SiteKey` being sent in rotate secret's request body ([#1285](https://github.com/cloudflare/cloudflare-go/issues/1285))
- turnstile: remove `SiteKey`/`Secret` being sent in update request body ([#1284](https://github.com/cloudflare/cloudflare-go/issues/1284))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.4.0 to 1.5.0 ([#1287](https://github.com/cloudflare/cloudflare-go/issues/1287))
- deps: bumps github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#1286](https://github.com/cloudflare/cloudflare-go/issues/1286))

## 0.67.0 (May 10th, 2023)

NOTES:

- dns_firewall: The `OriginIPs` field has been renamed to `UpstreamIPs`. ([#1246](https://github.com/cloudflare/cloudflare-go/issues/1246))

ENHANCEMENTS:

- device_posture_rule: add input fields tanium, intune and kolide ([#1268](https://github.com/cloudflare/cloudflare-go/issues/1268))
- waiting_room: add support for zone-level settings ([#1276](https://github.com/cloudflare/cloudflare-go/issues/1276))

BUG FIXES:

- rulesets: allow `PreserveQueryString` to be nullable ([#1275](https://github.com/cloudflare/cloudflare-go/issues/1275))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.25.1 to 2.25.3 ([#1274](https://github.com/cloudflare/cloudflare-go/issues/1274))
- deps: bumps golang.org/x/net from 0.9.0 to 0.10.0 ([#1280](https://github.com/cloudflare/cloudflare-go/issues/1280))

## 0.66.0 (26th April, 2023)

ENHANCEMENTS:

- access_application: Add `path_cookie_attribute` app setting ([#1223](https://github.com/cloudflare/cloudflare-go/issues/1223))
- certificate_packs: add `Status` field to indicate the status of certificate pack ([#1271](https://github.com/cloudflare/cloudflare-go/issues/1271))
- data localization: add support for regional hostnames API ([#1270](https://github.com/cloudflare/cloudflare-go/issues/1270))
- dns: add support for importing and exporting DNS records using BIND file configurations ([#1266](https://github.com/cloudflare/cloudflare-go/issues/1266))
- logpush: add support for max upload parameters ([#1272](https://github.com/cloudflare/cloudflare-go/issues/1272))
- turnstile: add support for turnstile ([#1267](https://github.com/cloudflare/cloudflare-go/issues/1267))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.3.6 to 1.4.0 ([#1269](https://github.com/cloudflare/cloudflare-go/issues/1269))

## 0.65.0 (12th April, 2023)

ENHANCEMENTS:

- access: Add `auto_redirect_to_identity` flag to Access organizations ([#1260](https://github.com/cloudflare/cloudflare-go/issues/1260))
- access: Add `isolation_required` flag to Access policies ([#1258](https://github.com/cloudflare/cloudflare-go/issues/1258))
- rulesets: add support for add operation to HTTP header configuration ([#1253](https://github.com/cloudflare/cloudflare-go/issues/1253))
- rulesets: add support for the `compress_response` action ([#1261](https://github.com/cloudflare/cloudflare-go/issues/1261))
- rulesets: add support for the `http_response_compression` phase ([#1261](https://github.com/cloudflare/cloudflare-go/issues/1261))

DEPENDENCIES:

- deps: bumps golang.org/x/net from 0.8.0 to 0.9.0 ([#1263](https://github.com/cloudflare/cloudflare-go/issues/1263))

## 0.64.0 (29th March, 2023)

BREAKING CHANGES:

- dns: Changed Create/UpdateDNSRecord method signatures to return (DNSRecord, error) ([#1243](https://github.com/cloudflare/cloudflare-go/issues/1243))
- zone: `UpdateZoneSingleSetting` has been renamed to `UpdateZoneSetting` and updated method signature inline with our expected conventions ([#1251](https://github.com/cloudflare/cloudflare-go/issues/1251))
- zone: `ZoneSingleSetting` has been renamed to `GetZoneSetting` and updated method signature inline with our expected conventions ([#1251](https://github.com/cloudflare/cloudflare-go/issues/1251))

ENHANCEMENTS:

- access_identity_provider: add `claims` and `scopes` fields ([#1237](https://github.com/cloudflare/cloudflare-go/issues/1237))
- access_identity_provider: add scim_config field ([#1178](https://github.com/cloudflare/cloudflare-go/issues/1178))
- devices_policy: update `Mode` field to use new `ServiceMode` string type with explicit const service mode values ([#1249](https://github.com/cloudflare/cloudflare-go/issues/1249))
- ssl: make `GeoRestrictions` a pointer inside of ZoneCustomSSL ([#1244](https://github.com/cloudflare/cloudflare-go/issues/1244))
- zone: `GetZoneSetting` and `UpdateZoneSetting` now allow configuring the path for where a setting resides instead of assuming `settings` ([#1251](https://github.com/cloudflare/cloudflare-go/issues/1251))

BUG FIXES:

- teams_rules: `AllowChildBypass` changes from a `bool` to `*bool` ([#1242](https://github.com/cloudflare/cloudflare-go/issues/1242))
- teams_rules: `BypassParentRule` changes from a `bool` to `*bool` ([#1242](https://github.com/cloudflare/cloudflare-go/issues/1242))
- tunnel: Fix 'CreateTunnel' for tunnels using config_src ([#1238](https://github.com/cloudflare/cloudflare-go/issues/1238))

DEPENDENCIES:

- deps: bumps actions/setup-go from 3 to 4 ([#1236](https://github.com/cloudflare/cloudflare-go/issues/1236))
- deps: bumps github.com/urfave/cli/v2 from 2.25.0 to 2.25.1 ([#1250](https://github.com/cloudflare/cloudflare-go/issues/1250))

## 0.63.0 (15th March, 2023)

BREAKING CHANGES:

- tunnel: renamed `Tunnel` to `GetTunnel` ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))
- tunnel: renamed `Tunnels` to `ListTunnels` ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))

ENHANCEMENTS:

- access_organization: add ui_read_only_toggle_reason field ([#1181](https://github.com/cloudflare/cloudflare-go/issues/1181))
- added audit_ssh to gateway actions, updated gateway rule settings ([#1226](https://github.com/cloudflare/cloudflare-go/issues/1226))
- addressing: Add `Address Map` support ([#1232](https://github.com/cloudflare/cloudflare-go/issues/1232))
- teams_account: add support for `check_disks` ([#1197](https://github.com/cloudflare/cloudflare-go/issues/1197))
- tunnel: updated parameters to latest API docs ([#1227](https://github.com/cloudflare/cloudflare-go/issues/1227))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.24.4 to 2.25.0 ([#1229](https://github.com/cloudflare/cloudflare-go/issues/1229))
- deps: bumps golang.org/x/net from 0.7.0 to 0.8.0 ([#1228](https://github.com/cloudflare/cloudflare-go/issues/1228))

## 0.62.0 (1st March, 2023)

ENHANCEMENTS:

- dex_test: add CRUD functionality for DEX test configurations ([#1209](https://github.com/cloudflare/cloudflare-go/issues/1209))
- dlp: Adds support for partial payload logging ([#1212](https://github.com/cloudflare/cloudflare-go/issues/1212))
- teams_accounts: Add new root_certificate_installation_enabled field ([#1208](https://github.com/cloudflare/cloudflare-go/issues/1208))
- teams_rules: Add `untrusted_cert` rule setting ([#1214](https://github.com/cloudflare/cloudflare-go/issues/1214))
- tunnels: automatically paginate `ListTunnels` ([#1206](https://github.com/cloudflare/cloudflare-go/issues/1206))

BUG FIXES:

- dex_test: use dex test types and json struct mappings instead of managed networks ([#1213](https://github.com/cloudflare/cloudflare-go/issues/1213))
- dns: dont reuse DNSListResponse when using pagination to avoid Proxied pointer overwrite ([#1222](https://github.com/cloudflare/cloudflare-go/issues/1222))

DEPENDENCIES:

- deps: bumps github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#1220](https://github.com/cloudflare/cloudflare-go/issues/1220))
- deps: bumps github.com/urfave/cli/v2 from 2.24.3 to 2.24.4 ([#1210](https://github.com/cloudflare/cloudflare-go/issues/1210))
- deps: bumps golang.org/x/net from 0.0.0-20220722155237-a158d28d115b to 0.7.0 ([#1218](https://github.com/cloudflare/cloudflare-go/issues/1218))
- deps: bumps golang.org/x/net from 0.0.0-20220722155237-a158d28d115b to 0.7.0 ([#1219](https://github.com/cloudflare/cloudflare-go/issues/1219))
- deps: bumps golang.org/x/text from 0.3.7 to 0.3.8 ([#1215](https://github.com/cloudflare/cloudflare-go/issues/1215))
- deps: bumps golang.org/x/text from 0.3.7 to 0.3.8 ([#1216](https://github.com/cloudflare/cloudflare-go/issues/1216))
- deps: bumps golang.org/x/time from 0.0.0-20220224211638-0e9765cccd65 to 0.3.0 ([#1217](https://github.com/cloudflare/cloudflare-go/issues/1217))

## 0.61.0 (15th February, 2023)

ENHANCEMENTS:

- cloudflare: make it clearer when we hit a server error and to retry later ([#1207](https://github.com/cloudflare/cloudflare-go/issues/1207))
- devices_policy: Add new exclude_office_ips field to policy ([#1205](https://github.com/cloudflare/cloudflare-go/issues/1205))
- dlp_profile: Use int rather than uint for allowed_match_count field ([#1200](https://github.com/cloudflare/cloudflare-go/issues/1200))

BUG FIXES:

- dns: always send `tags` to allow clearing ([#1196](https://github.com/cloudflare/cloudflare-go/issues/1196))
- stream: renamed `RequiredSignedURLs` to `RequireSignedURLs` ([#1202](https://github.com/cloudflare/cloudflare-go/issues/1202))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.24.2 to 2.24.3 ([#1199](https://github.com/cloudflare/cloudflare-go/issues/1199))

## 0.60.0 (1st February, 2023)

BREAKING CHANGES:

- queues: UpdateQueue has been updated to match the API and now correctly updates a Queue's name ([#1188](https://github.com/cloudflare/cloudflare-go/issues/1188))

ENHANCEMENTS:

- dlp_profile: Add new allowed_match_count field to profiles ([#1193](https://github.com/cloudflare/cloudflare-go/issues/1193))
- dns: allow sending empty strings to remove comments ([#1195](https://github.com/cloudflare/cloudflare-go/issues/1195))
- magic_transit_ipsec_tunnel: makes customer endpoint an optional field for ipsec tunnel creation ([#1185](https://github.com/cloudflare/cloudflare-go/issues/1185))
- rulesets: add support for `score_per_period` and `score_response_header_name` ([#1183](https://github.com/cloudflare/cloudflare-go/issues/1183))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.3.5 to 1.3.6 ([#1184](https://github.com/cloudflare/cloudflare-go/issues/1184))
- deps: bumps github.com/urfave/cli/v2 from 2.23.7 to 2.24.1 ([#1180](https://github.com/cloudflare/cloudflare-go/issues/1180))
- deps: bumps github.com/urfave/cli/v2 from 2.24.1 to 2.24.2 ([#1191](https://github.com/cloudflare/cloudflare-go/issues/1191))
- deps: bumps goreleaser/goreleaser-action from 4.1.0 to 4.2.0 ([#1192](https://github.com/cloudflare/cloudflare-go/issues/1192))

## 0.59.0 (January 18th, 2023)

BREAKING CHANGES:

- dns: remove these read-only fields from `UpdateDNSRecordParams`: `CreatedOn`, `ModifiedOn`, `Meta`, `ZoneID`, `ZoneName`, `Proxiable`, and `Locked` ([#1170](https://github.com/cloudflare/cloudflare-go/issues/1170))
- dns: the fields `CreatedOn` and `ModifiedOn` are removed from `ListDNSRecordsParams` ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))

NOTES:

- dns: remove additional lookup from `Update` operations when `Name` or `Type` was omitted ([#1170](https://github.com/cloudflare/cloudflare-go/issues/1170))

ENHANCEMENTS:

- access_organization: add user_seat_expiration_inactive_time field ([#1159](https://github.com/cloudflare/cloudflare-go/issues/1159))
- dns: `GetDNSRecord`, `UpdateDNSRecord` and `DeleteDNSRecord` now return the new, dedicated error `ErrMissingDNSRecordID` when an empty DNS record ID is given. ([#1174](https://github.com/cloudflare/cloudflare-go/issues/1174))
- dns: the URL parameter `tag-match` for listing DNS records is now supported as the field `TagMatch` in `ListDNSRecordsParams` ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))
- dns: update default `per_page` attribute to 100 records ([#1171](https://github.com/cloudflare/cloudflare-go/issues/1171))
- teams_rules: adds support for Egress Policies ([#1142](https://github.com/cloudflare/cloudflare-go/issues/1142))
- workers: Add support for compatibility_date and compatibility_flags when upoading a worker script ([#1177](https://github.com/cloudflare/cloudflare-go/issues/1177))
- workers: script upload now supports Queues bindings ([#1176](https://github.com/cloudflare/cloudflare-go/issues/1176))

BUG FIXES:

- dns: don't send "priority" for list operations as it isn't supported and is only used for internal filtering ([#1167](https://github.com/cloudflare/cloudflare-go/issues/1167))
- dns: the field `Tags` in `ListDNSRecordsParams` was not correctly serialized into URL queries ([#1173](https://github.com/cloudflare/cloudflare-go/issues/1173))
- managednetworks: Update should be PUT ([#1172](https://github.com/cloudflare/cloudflare-go/issues/1172))

## 0.58.1 (January 5th, 2023)

ENHANCEMENTS:

- cloudflare: automatically redact sensitive values from HTTP interactions ([#1164](https://github.com/cloudflare/cloudflare-go/issues/1164))

## 0.58.0 (January 4th, 2023)

BREAKING CHANGES:

- dns: `DNSRecord` has been renamed to `GetDNSRecord` ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
- dns: `DNSRecords` has been renamed to `ListDNSRecords` ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
- dns: method signatures have been updated to align with the upcoming client conventions ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
- origin_ca: renamed to `CreateOriginCertificate` to `CreateOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
- origin_ca: renamed to `OriginCARootCertificate` to `GetOriginCARootCertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
- origin_ca: renamed to `OriginCertificate` to `GetOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
- origin_ca: renamed to `OriginCertificates` to `ListOriginCACertificates` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
- origin_ca: renamed to `RevokeOriginCertificate` to `RevokeOriginCACertificate` ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))

ENHANCEMENTS:

- dns: add support for tags and comments ([#1151](https://github.com/cloudflare/cloudflare-go/issues/1151))
- mtls_certificate: add support for managing mTLS certificates and assocations ([#1150](https://github.com/cloudflare/cloudflare-go/issues/1150))
- origin_ca: add support for using API keys, API tokens or API User service keys for interacting with Origin CA endpoints ([#1161](https://github.com/cloudflare/cloudflare-go/issues/1161))
- workers: Add support for workers logpush enablement on script upload ([#1160](https://github.com/cloudflare/cloudflare-go/issues/1160))

BUG FIXES:

- email_routing_destination: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- email_routing_rules: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- filter: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- firewall_rules: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- lockdown: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- queue: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- teams_list: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))
- workers_kv: use empty reponse struct on each page call ([#1156](https://github.com/cloudflare/cloudflare-go/issues/1156))

DEPENDENCIES:

- deps: bumps github.com/hashicorp/go-retryablehttp from 0.7.1 to 0.7.2 ([#1162](https://github.com/cloudflare/cloudflare-go/issues/1162))

## 0.57.1 (December 23rd, 2022)

ENHANCEMENTS:

- tiered_cache: Add support for Tiered Caching interactions for setting Smart and Generic topologies ([#1149](https://github.com/cloudflare/cloudflare-go/issues/1149))

BUG FIXES:

- workers: correctly set `body` value for non-ES module uploads ([#1155](https://github.com/cloudflare/cloudflare-go/issues/1155))

## 0.57.0 (December 22nd, 2022)

BREAKING CHANGES:

- workers: API operations now target account level resources instead of older zone level resources (these are a 1:1 now) ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_bindings: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_cron_triggers: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_kv: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_routes: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_secrets: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers_tails: method signatures have been updated to align with the upcoming client conventions ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))

NOTES:

- workers: all worker methods have been split into product ownership(-ish) files ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))
- workers: all worker methods now require an explicit `ResourceContainer` for endpoints instead of relying on the globally defined `api.AccountID` ([#1137](https://github.com/cloudflare/cloudflare-go/issues/1137))

ENHANCEMENTS:

- managed_networks: add CRUD functionality for managednetworks ([#1148](https://github.com/cloudflare/cloudflare-go/issues/1148))

DEPENDENCIES:

- deps: bumps goreleaser/goreleaser-action from 3.2.0 to 4.1.0 ([#1146](https://github.com/cloudflare/cloudflare-go/issues/1146))

## 0.56.0 (December 5th, 2022)

BREAKING CHANGES:

- pages: Changed the type of EnvVars in PagesProjectDeploymentConfigEnvironment & PagesProjectDeployment in order to properly support secrets. ([#1136](https://github.com/cloudflare/cloudflare-go/issues/1136))

NOTES:

- pages: removed the v1 logs endpoint for Pages deployments. Please switch to v2: https://developers.cloudflare.com/api/operations/pages-deployment-get-deployment-logs ([#1135](https://github.com/cloudflare/cloudflare-go/issues/1135))

ENHANCEMENTS:

- cache_rules: add ignore option to query string struct ([#1140](https://github.com/cloudflare/cloudflare-go/issues/1140))
- pages: Updates bindings and other Functions related propreties. Service bindings, secrets, fail open/close and usage model are all now supported. ([#1136](https://github.com/cloudflare/cloudflare-go/issues/1136))
- workers: Support for Workers Analytics Engine bindings ([#1133](https://github.com/cloudflare/cloudflare-go/issues/1133))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.23.5 to 2.23.6 ([#1139](https://github.com/cloudflare/cloudflare-go/issues/1139))

## 0.55.0 (November 23th, 2022)

BREAKING CHANGES:

- workers_kv: `CreateWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `DeleteWorkersKVBulk` has been renamed to `DeleteWorkersKVEntries`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `DeleteWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `DeleteWorkersKV` has been renamed to `DeleteWorkersKVEntry`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `ListWorkersKVNamespaces` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `ListWorkersKVsWithOptions` has been removed. Use `ListWorkersKVKeys` instead and pass in the options. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `ListWorkersKVs` has been renamed to `ListWorkersKVKeys`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `ReadWorkersKV` has been renamed to `GetWorkersKV`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `UpdateWorkersKVNamespace` has been updated to match the experimental client method signatures (https://github.com/cloudflare/cloudflare-go/blob/master/docs/experimental.md). ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `WriteWorkersKVBulk` has been renamed to `WriteWorkersKVEntries`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))
- workers_kv: `WriteWorkersKV` has been renamed to `WriteWorkersKVEntry`. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))

ENHANCEMENTS:

- device_posture_rule: add input fields crowdstrike ([#1126](https://github.com/cloudflare/cloudflare-go/issues/1126))
- queue: add support queue API ([#1131](https://github.com/cloudflare/cloudflare-go/issues/1131))
- r2: Add support for listing R2 buckets ([#1063](https://github.com/cloudflare/cloudflare-go/issues/1063))
- workers_domain: add support for workers domain API ([#1130](https://github.com/cloudflare/cloudflare-go/issues/1130))
- workers_kv: `ListWorkersKVNamespaces` automatically paginates all results unless `PerPage` is defined. ([#1115](https://github.com/cloudflare/cloudflare-go/issues/1115))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.23.4 to 2.23.5 ([#1127](https://github.com/cloudflare/cloudflare-go/issues/1127))

## 0.54.0 (November 9th, 2022)

ENHANCEMENTS:

- access: add support for service token rotation ([#1120](https://github.com/cloudflare/cloudflare-go/issues/1120))
- deps: fix import grouping, code formatting and enable goimports linter ([#1121](https://github.com/cloudflare/cloudflare-go/issues/1121))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.3.4 to 1.3.5 ([#1123](https://github.com/cloudflare/cloudflare-go/issues/1123))
- deps: bumps github.com/urfave/cli/v2 from 2.20.3 to 2.23.0 ([#1122](https://github.com/cloudflare/cloudflare-go/issues/1122))
- deps: bumps github.com/urfave/cli/v2 from 2.23.0 to 2.23.2 ([#1124](https://github.com/cloudflare/cloudflare-go/issues/1124))
- deps: bumps github.com/urfave/cli/v2 from 2.23.2 to 2.23.4 ([#1125](https://github.com/cloudflare/cloudflare-go/issues/1125))

## 0.53.0 (October 26th, 2022)

BREAKING CHANGES:

- account_member: `CreateAccountMember` has been updated to accept a `CreateAccountMemberParams` struct instead of multiple parameters ([#1095](https://github.com/cloudflare/cloudflare-go/issues/1095))
- teams_list: updated methods to match the experimental client format ([#1114](https://github.com/cloudflare/cloudflare-go/issues/1114))

ENHANCEMENTS:

- account_member: add support for domain scoped roles ([#1095](https://github.com/cloudflare/cloudflare-go/issues/1095))
- cloudflare: expose `Messages` from the `Response` object ([#1106](https://github.com/cloudflare/cloudflare-go/issues/1106))
- dlp: Adds support for DLP resources ([#1111](https://github.com/cloudflare/cloudflare-go/issues/1111))
- teams_list: `List` operations now automatically paginate ([#1114](https://github.com/cloudflare/cloudflare-go/issues/1114))
- total_tls: adds support for TotalTLS ([#1105](https://github.com/cloudflare/cloudflare-go/issues/1105))
- waiting_room: add support for waiting room rules ([#1102](https://github.com/cloudflare/cloudflare-go/issues/1102))

DEPENDENCIES:

- deps: `ioutil` package is being deprecated in favor of `io` ([#1116](https://github.com/cloudflare/cloudflare-go/issues/1116))
- deps: bumps github.com/stretchr/testify from 1.8.0 to 1.8.1 ([#1119](https://github.com/cloudflare/cloudflare-go/issues/1119))
- deps: bumps github.com/urfave/cli/v2 from 2.19.2 to 2.20.2 ([#1108](https://github.com/cloudflare/cloudflare-go/issues/1108))
- deps: bumps github.com/urfave/cli/v2 from 2.20.2 to 2.20.3 ([#1118](https://github.com/cloudflare/cloudflare-go/issues/1118))
- deps: bumps goreleaser/goreleaser-action from 3.1.0 to 3.2.0 ([#1112](https://github.com/cloudflare/cloudflare-go/issues/1112))
- deps: remove `github.com/pkg/errors` in favor of `errors` ([#1117](https://github.com/cloudflare/cloudflare-go/issues/1117))

## 0.52.0 (October 12th, 2022)

ENHANCEMENTS:

- access: add UI read-only field to organizations ([#1104](https://github.com/cloudflare/cloudflare-go/issues/1104))
- devices_policy: Add support for additional device settings policies ([#1090](https://github.com/cloudflare/cloudflare-go/issues/1090))
- rulesets: add support for `sensitivity_level` to override all rule sensitivity ([#1093](https://github.com/cloudflare/cloudflare-go/issues/1093))

DEPENDENCIES:

- deps: bumps dependabot/fetch-metadata from 1.3.3 to 1.3.4 ([#1097](https://github.com/cloudflare/cloudflare-go/issues/1097))
- deps: bumps github.com/urfave/cli/v2 from 2.16.3 to 2.17.1 ([#1094](https://github.com/cloudflare/cloudflare-go/issues/1094))
- deps: bumps github.com/urfave/cli/v2 from 2.17.1 to 2.19.2 ([#1103](https://github.com/cloudflare/cloudflare-go/issues/1103))

## 0.51.0 (September 28th, 2022)

BREAKING CHANGES:

- load_balancing: update method signatures to match experimental conventions ([#1084](https://github.com/cloudflare/cloudflare-go/issues/1084))

ENHANCEMENTS:

- device_posture_rule: add input fields for linux OS ([#1087](https://github.com/cloudflare/cloudflare-go/issues/1087))
- load_balancing: support adaptive_routing and location_strategy ([#1091](https://github.com/cloudflare/cloudflare-go/issues/1091))

BUG FIXES:

- user-agent-blocking-rules: add missing managed_challenge validation and removed the deprecated whitelist one ([#1089](https://github.com/cloudflare/cloudflare-go/issues/1089))

## 0.50.0 (September 14, 2022)

ENHANCEMENTS:

- auditlogs: add support for hide_user_logs filter parameter ([#1075](https://github.com/cloudflare/cloudflare-go/issues/1075))

BUG FIXES:

- cloudflare: exiting closer to the source on context timeouts to improve error messaging and better defend from potential edge cases ([#1080](https://github.com/cloudflare/cloudflare-go/issues/1080))
- origin certificate: Fix API auth type used ([#1082](https://github.com/cloudflare/cloudflare-go/issues/1082))

DEPENDENCIES:

- deps: bumps github.com/urfave/cli/v2 from 2.11.2 to 2.14.0 ([#1077](https://github.com/cloudflare/cloudflare-go/issues/1077))
- deps: bumps github.com/urfave/cli/v2 from 2.14.0 to 2.14.1 ([#1081](https://github.com/cloudflare/cloudflare-go/issues/1081))
- deps: bumps github.com/urfave/cli/v2 from 2.14.1 to 2.15.0 ([#1085](https://github.com/cloudflare/cloudflare-go/issues/1085))
- deps: bumps github.com/urfave/cli/v2 from 2.15.0 to 2.16.3 ([#1086](https://github.com/cloudflare/cloudflare-go/issues/1086))

## 0.49.0 (August 31st, 2022)

ENHANCEMENTS:

- access_service_token: add support for refreshing an existing token in place ([#1074](https://github.com/cloudflare/cloudflare-go/issues/1074))
- api: addded context and headers to Raw method ([#1068](https://github.com/cloudflare/cloudflare-go/issues/1068))
- api_shield: add GET/PUT for API Shield Configuration ([#1059](https://github.com/cloudflare/cloudflare-go/issues/1059))
- pages_project: Add `kv_namespaces`, `durable_object_namespaces`, `r2_buckets`, and `d1_databases` bindings to deployment config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
- pages_project: Add `preview_deployment_setting`, `preview_branch_includes`, and `preview_branch_excludes` to source config ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
- pages_project: Add `production_branch` field ([#1065](https://github.com/cloudflare/cloudflare-go/issues/1065))
- teams_account: add support for `os_distro_name` and `os_distro_revision` ([#1073](https://github.com/cloudflare/cloudflare-go/issues/1073))
- url_normalization_settings: Add APIs to get and update URL normalization settings ([#1071](https://github.com/cloudflare/cloudflare-go/issues/1071))
- workers: Support for multipart encoding for DownloadWorker on a module-format Worker script ([#1040](https://github.com/cloudflare/cloudflare-go/issues/1040))

BUG FIXES:

- cloudflare: fix nil dereference error in makeRequestWithAuthTypeAndHeaders ([#1072](https://github.com/cloudflare/cloudflare-go/issues/1072))
- email_routing_rules: Fix response for email routing catch all rule. ([#1070](https://github.com/cloudflare/cloudflare-go/issues/1070))
- email_routing_settings: change enable endpoint from `enabled` to `enable` ([#1060](https://github.com/cloudflare/cloudflare-go/issues/1060))
- stream: Update pctComplete to string from int ([#1066](https://github.com/cloudflare/cloudflare-go/issues/1066))

DEPENDENCIES:

- deps: bumps goreleaser/goreleaser-action from 3.0.0 to 3.1.0 ([#1067](https://github.com/cloudflare/cloudflare-go/issues/1067))

## 0.48.0 (August 22nd, 2022)

ENHANCEMENTS:

- errors: add some error type convenience functions for mocking and inspection ([#1047](https://github.com/cloudflare/cloudflare-go/issues/1047))
- pages_project: Add compatibility date and compatibility_flags to pages deployment configs ([#1051](https://github.com/cloudflare/cloudflare-go/issues/1051))
- teams_account: add support for `suppress_footer` ([#1053](https://github.com/cloudflare/cloudflare-go/issues/1053))

BUG FIXES:

- r2: fix create bucket endpoint ([#1035](https://github.com/cloudflare/cloudflare-go/issues/1035))
- tunnel_configuration: Remove unnecessary double-unmarshalling due to changes in the API ([#1046](https://github.com/cloudflare/cloudflare-go/issues/1046))

## 0.47.1 (August 18th, 2022)

BUG FIXES:

- zonelockdown: add `Priority` to `ZoneLockdownCreateParams` and `ZoneLockdownUpdateParams` ([#1052](https://github.com/cloudflare/cloudflare-go/issues/1052))

## 0.47.0 (August 17th, 2022)

BREAKING CHANGES:

- certificate_packs: deprecate "custom" configuration for ACM everywhere ([#1032](https://github.com/cloudflare/cloudflare-go/issues/1032))

ENHANCEMENTS:

- cloudflare: make it clear when the rate limit retries have been exhausted ([#1043](https://github.com/cloudflare/cloudflare-go/issues/1043))
- email_routing_destination: Adds support for the email routing destination API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
- email_routing_rules: Adds support for the email routing rules API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
- email_routing_settings: Adds support for the email routing settings API ([#1034](https://github.com/cloudflare/cloudflare-go/issues/1034))
- filter: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
- firewall_rule: fix double endpoint calls & moving towards common method signature ([#1016](https://github.com/cloudflare/cloudflare-go/issues/1016))
- lockdown: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1017](https://github.com/cloudflare/cloudflare-go/issues/1017))
- r2: Add in support for creating and deleting R2 buckets ([#1028](https://github.com/cloudflare/cloudflare-go/issues/1028))
- rulesets: add support for `http_config_settings` phase and supporting actions ([#1036](https://github.com/cloudflare/cloudflare-go/issues/1036))
- workers-account-settings: Add in support for Workers account settings API ([#1027](https://github.com/cloudflare/cloudflare-go/issues/1027))
- workers-subdomain: Add in support Workers Subdomain API ([#1031](https://github.com/cloudflare/cloudflare-go/issues/1031))
- workers-tail: Add in support for Workers tail API ([#1026](https://github.com/cloudflare/cloudflare-go/issues/1026))
- workers: Add support for attaching a worker to a domain ([#1014](https://github.com/cloudflare/cloudflare-go/issues/1014))
- workers: Add support to upload module workers ([#1010](https://github.com/cloudflare/cloudflare-go/issues/1010))

BUG FIXES:

- email_routing_destination: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
- email_routing_rules: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
- email_routing_settings: Update API reference URLs ([#1038](https://github.com/cloudflare/cloudflare-go/issues/1038))
- tunnel_routes: Fix not removing route when it contains virtual network ([#1030](https://github.com/cloudflare/cloudflare-go/issues/1030))
- workers_test: Fix incorrect test from PR #1014 ([#1048](https://github.com/cloudflare/cloudflare-go/issues/1048))
- workers_test: Use application/json mime-type in headers ([#1049](https://github.com/cloudflare/cloudflare-go/issues/1049))

DEPENDENCIES:

- deps: bumps golang.org/x/tools/gopls from 0.9.3 to 0.9.4 ([#1044](https://github.com/cloudflare/cloudflare-go/issues/1044))
- deps: bumps github.com/golangci/golangci-lint from 1.47.3 to 1.48.0 ([#1020](https://github.com/cloudflare/cloudflare-go/issues/1020))
- deps: bumps github.com/urfave/cli/v2 from 2.11.1 to 2.11.2 ([#1042](https://github.com/cloudflare/cloudflare-go/issues/1042))
- deps: bumps golang.org/x/tools/gopls from 0.9.1 to 0.9.2 ([#1037](https://github.com/cloudflare/cloudflare-go/issues/1037))
- deps: bumps golang.org/x/tools/gopls from 0.9.2 to 0.9.3 ([#1039](https://github.com/cloudflare/cloudflare-go/issues/1039))

## 0.46.0 (3rd August, 2022)

NOTES:

- docs: add release notes ([#1001](https://github.com/cloudflare/cloudflare-go/issues/1001))

ENHANCEMENTS:

- filter: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
- firewall_rule: automatically paginate `List` results unless `Page` and `PerPage` are provided ([#1004](https://github.com/cloudflare/cloudflare-go/issues/1004))
- rulesets: add support for `http_custom_errors` phase ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))
- rulesets: add support for `serve_error` action ([#998](https://github.com/cloudflare/cloudflare-go/issues/998))

BUG FIXES:

- access_application: fix inability to set bool values to false ([#1006](https://github.com/cloudflare/cloudflare-go/issues/1006))
- rulesets: fix sni action parameter ([#1002](https://github.com/cloudflare/cloudflare-go/issues/1002))

DEPENDENCIES:

- provider: bumps github.com/golangci/golangci-lint from 1.47.1 to 1.47.2 ([#1005](https://github.com/cloudflare/cloudflare-go/issues/1005))
- provider: bumps github.com/golangci/golangci-lint from 1.47.2 to 1.47.3 ([#1008](https://github.com/cloudflare/cloudflare-go/issues/1008))
- provider: bumps github.com/urfave/cli/v2 from 2.11.0 to 2.11.1 ([#1003](https://github.com/cloudflare/cloudflare-go/issues/1003))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

## 0.45.0 (July 20th, 2022)
