# Change Log

## GoVultr v1 changelog is located [here](https://github.com/vultr/govultr/blob/v1/CHANGELOG.md)

## [v2.5.1](https://github.com/vultr/govultr/compare/v2.5.0..v2.5.1) (2021-05-10)
### Bug fix
* Instances : BackupScheduleReq change DOW + Hour to pointers  [145](https://github.com/vultr/govultr/pull/145)

## [v2.5.0](https://github.com/vultr/govultr/compare/v2.4.2..v2.5.0) (2021-05-06)
### Enhancement
* LoadBalancers : New Features and endpoints [143](https://github.com/vultr/govultr/pull/143)
  * Ability to attach private networks
  * Ability to set firewalls
  * Get Firewall Rules
  * List Firewall Rules 

## [v2.4.2](https://github.com/vultr/govultr/compare/v2.4.1..v2.4.2) (2021-05-03)
### Bug fix
* Instances : ListPrivateNetworks missing paging ability [140](https://github.com/vultr/govultr/pull/140)

## [v2.4.1](https://github.com/vultr/govultr/compare/v2.4.0..v2.4.1) (2021-05-03)
### Dependency updates
* Bump github.com/hashicorp/go-retryablehttp from 0.6.8 to 0.7.0 [138](https://github.com/vultr/govultr/pull/138)
* Bump github.com/google/go-querystring from 1.0.0 to 1.1.0 [137](https://github.com/vultr/govultr/pull/137)

## [v2.4.0](https://github.com/vultr/govultr/compare/v2.3.2..v2.4.0) (2021-02-11)
### Enhancement
* Block Storage - add `mount_id` field to BlockStorage struct [131](https://github.com/vultr/govultr/pull/131)
* Plans - add `disk_count` field to Plan and BareMetalPlan struct [130](https://github.com/vultr/govultr/pull/130)

## [v2.3.2](https://github.com/vultr/govultr/compare/v2.3.1..v2.3.2) (2021-01-06)
### Bug Fix
* Instances - Fixed DetachPrivateNetwork which had wrong URI [122](https://github.com/vultr/govultr/pull/122)

## [v2.3.1](https://github.com/vultr/govultr/compare/v2.3.0..v2.3.1) (2021-01-04)
### Bug Fix
* Domain Record - removed `omitempty` on `name` field in `DomainRecordReq` to allow creation of nameless records. [120](https://github.com/vultr/govultr/pull/120)

## [v2.3.0](https://github.com/vultr/govultr/compare/v2.2.0..v2.3.0) (2020-12-17)
### Enhancement
* Bare Metal - Start call added [118](https://github.com/vultr/govultr/pull/118)

## [v2.2.0](https://github.com/vultr/govultr/compare/v2.1.0..v2.2.0) (2020-12-07)
### Breaking Change
* All bools have been updated to pointers to avoid issues where false values not being sent in request. Thanks @Static-Flow [115](https://github.com/vultr/govultr/pull/115)

## [v2.1.0](https://github.com/vultr/govultr/compare/v2.0.0..v2.1.0) (2020-11-30)
### Bug fixes
* ReservedIP - Attach call creates proper json. [112](https://github.com/vultr/govultr/pull/112)
* User - APIEnabled takes pointer of bool [112](https://github.com/vultr/govultr/pull/112)

## v2.0.0 (2020-11-20)
### Initial Release
* GoVultr v2.0.0 Release - Uses Vultr API v2.
* GoVultr v1.0.0 is now on [branch v1](https://github.com/vultr/govultr/tree/v1)