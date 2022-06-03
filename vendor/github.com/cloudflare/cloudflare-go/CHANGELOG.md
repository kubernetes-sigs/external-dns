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

## 0.45.0 (July 20th, 2022)
