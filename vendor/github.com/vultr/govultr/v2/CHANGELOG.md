# Change Log

## GoVultr v1 changelog is located [here](https://github.com/vultr/govultr/blob/v1/CHANGELOG.md)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
## [v2.9.0](https://github.com/vultr/govultr/compare/v2.8.1..v2.9.0) (2021-09-27)
### Breaking Change
* Kubernetes : PlanID is now Plan and Count is now NodeQuantity to follow API pattern [161](https://github.com/vultr/govultr/pull/161)

### Enhancement
* Snapshots : Add compressed size field [162](https://github.com/vultr/govultr/pull/162)

## [v2.8.1](https://github.com/vultr/govultr/compare/v2.8.0..v2.8.1) (2021-08-31)
### Enhancement
* Kubernetes : Add support for deletion with resources [159](https://github.com/vultr/govultr/pull/159)
* Kubernetes : Add support for getting available versions[159](https://github.com/vultr/govultr/pull/159)

### Dependency Update
* Bump Go version to 1.16 [158](https://github.com/vultr/govultr/pull/158)

## [v2.8.0](https://github.com/vultr/govultr/compare/v2.7.1..v2.8.0) (2021-08-18)
### Enhancement
* Added support for Vultr Kubernetes Engine [156](https://github.com/vultr/govultr/pull/156)

## [v2.7.1](https://github.com/vultr/govultr/compare/v2.7.0..v2.7.1) (2021-07-15)
### Enhancement
* BareMetal : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)
* Instances : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)

## [v2.7.0](https://github.com/vultr/govultr/compare/v2.6.0..v2.7.0) (2021-07-14)
### Enhancement
* BareMetal : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Instances : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Applications : added support for marketplace applications [150](https://github.com/vultr/govultr/pull/150)

## [v2.6.0](https://github.com/vultr/govultr/compare/v2.5.1..v2.6.0) (2021-07-02)
### Enhancement
* BareMetal : Add support for `persistent_pxe` [148](https://github.com/vultr/govultr/pull/148)

||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
## [v2.9.0](https://github.com/vultr/govultr/compare/v2.8.1..v2.9.0) (2021-09-27)
### Breaking Change
* Kubernetes : PlanID is now Plan and Count is now NodeQuantity to follow API pattern [161](https://github.com/vultr/govultr/pull/161)

### Enhancement
* Snapshots : Add compressed size field [162](https://github.com/vultr/govultr/pull/162)

## [v2.8.1](https://github.com/vultr/govultr/compare/v2.8.0..v2.8.1) (2021-08-31)
### Enhancement
* Kubernetes : Add support for deletion with resources [159](https://github.com/vultr/govultr/pull/159)
* Kubernetes : Add support for getting available versions[159](https://github.com/vultr/govultr/pull/159)

### Dependency Update
* Bump Go version to 1.16 [158](https://github.com/vultr/govultr/pull/158)

## [v2.8.0](https://github.com/vultr/govultr/compare/v2.7.1..v2.8.0) (2021-08-18)
### Enhancement
* Added support for Vultr Kubernetes Engine [156](https://github.com/vultr/govultr/pull/156)

## [v2.7.1](https://github.com/vultr/govultr/compare/v2.7.0..v2.7.1) (2021-07-15)
### Enhancement
* BareMetal : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)
* Instances : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)

## [v2.7.0](https://github.com/vultr/govultr/compare/v2.6.0..v2.7.0) (2021-07-14)
### Enhancement
* BareMetal : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Instances : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Applications : added support for marketplace applications [150](https://github.com/vultr/govultr/pull/150)

## [v2.6.0](https://github.com/vultr/govultr/compare/v2.5.1..v2.6.0) (2021-07-02)
### Enhancement
* BareMetal : Add support for `persistent_pxe` [148](https://github.com/vultr/govultr/pull/148)

>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
=======
## [v2.17.2](https://github.com/vultr/govultr/compare/v2.17.1...v2.17.2) (2022-06-13)

### Enhancement
* Reserved IP: Add support for updating label [227](https://github.com/vultr/govultr/pull/227)

## [v2.17.1](https://github.com/vultr/govultr/compare/v2.17.0...v2.17.1) (2022-06-02)
* Plans: Add GPU specific fields to plan struct [224](https://github.com/vultr/govultr/pull/224)

## [v2.17.0](https://github.com/vultr/govultr/compare/v2.16.0..v2.17.0) (2022-05-17)

### Enhancement
* Kubernetes: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)
* Baremetal: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)
* Instance: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)

### Bug fixes
* Kubernetes: fix data type for `auto_scaler` to avoid sending null values in requests when not set [222](https://github.com/vultr/govultr/pull/222)

### Breaking Change
* Kubernetes: change data type for `Tag` in node pool update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Kubernetes: change data type for `AutoScaler` in node pool update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Baremetal: change data type for `Tag` in update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Instance: change data type for `Tag` in update requirements struct [222](https://github.com/vultr/govultr/pull/222)


## [v2.16.0](https://github.com/vultr/govultr/compare/v2.15.1..v2.16.0) (2022-05-04)

### Enhancement
* Kubernetes: added auto scaler options to node pools [215](https://github.com/vultr/govultr/pull/215)
* Firewall rules: added new field `ip_type` in get/list responses to be consistent with the create calls [216](https://github.com/vultr/govultr/pull/216)
* Kubernetes: Upgrade support [217](https://github.com/vultr/govultr/pull/217)
* Baremetal: Added support for new `tags` field. This field allows multiple string tags to be associated with an instance [218](https://github.com/vultr/govultr/pull/218) 
* Instance: Added support for new `tags` field. This field allows multiple string tags to be associated with an instance [218](https://github.com/vultr/govultr/pull/218)

### Deprecations
* Instance: The `tag` field has been deprecated in favor for `tags` [218](https://github.com/vultr/govultr/pull/218)
* Baremetal: The `tag` field has been deprecated in favor for `tags` [218](https://github.com/vultr/govultr/pull/218)
* Firewall rules: The `type` field has been deprecated in favor for `ip_type` [216](https://github.com/vultr/govultr/pull/216)

### Dependency Update
* Bump github.com/hashicorp/go-retryablehttp from 0.7.0 to 0.7.1 [214](https://github.com/vultr/govultr/pull/214)

## [v2.15.1](https://github.com/vultr/govultr/compare/v2.15.0..v2.15.1) (2022-04-12)
### Bug fixes
* Block : add `omityempty` to `block_type` to prevent deploy issues [212](https://github.com/vultr/govultr/pull/212)

## [v2.15.0](https://github.com/vultr/govultr/compare/v2.14.2..v2.15.0) (2022-04-12)
### Enhancement 
* Block : New optional field `block_type`. This new field is currently optional but may become required at a later release [209](https://github.com/vultr/govultr/pull/209)
* VPC : New API endpoints that will be replacing `network` [210](https://github.com/vultr/govultr/pull/210)
* Updated Go version from 1.16 to 1.17 [208](https://github.com/vultr/govultr/pull/208)

### Deprecations
* Network : The network resource and all related private network fields on structs are deprecated. You should now be using the VPC provided replacements [210](https://github.com/vultr/govultr/pull/210)

## [v2.14.2](https://github.com/vultr/govultr/compare/v2.14.1..v2.14.2) (2022-03-23)
### Bug Fix
* Instances : restore support requestBody [206](https://github.com/vultr/govultr/pull/206) Thanks @andrake81

>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
## [v2.14.1](https://github.com/vultr/govultr/compare/v2.14.0..v2.14.1) (2022-02-02)
### Enhancement
* Improved retry error response [204](https://github.com/vultr/govultr/pull/204)

## [v2.14.0](https://github.com/vultr/govultr/compare/v2.13.0..v2.14.0) (2022-01-21)
### Enhancement
* ListOptions : [Added query param Region](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [200](https://github.com/vultr/govultr/pull/200)
* ListOptions : [Added query param Description](https://www.vultr.com/api/#operation/list-snapshots) that can be used with `Snapshot.List`  [202](https://github.com/vultr/govultr/pull/202)
* Snapshot : `CreateFromURL` has new optional field called `description` which lets you set a custom description [202](https://github.com/vultr/govultr/pull/202)

## [v2.13.0](https://github.com/vultr/govultr/compare/v2.12.0..v2.13.0) (2022-01-05)
### Enhancement
* ListOptions : [Added query params](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [197](https://github.com/vultr/govultr/pull/197)

## [v2.12.0](https://github.com/vultr/govultr/compare/v2.11.1..v2.12.0) (2021-12-01)
### Breaking Changes
* Plans : Changed `MonthlyCost` from `int` to `float32` [192](https://github.com/vultr/govultr/pull/192)

## [v2.11.1](https://github.com/vultr/govultr/compare/v2.11.0..v2.11.1) (2021-11-26)
### Bug fixes
* LoadBalancers : Fixed SSL struct json params to the proper API fields [189](https://github.com/vultr/govultr/pull/189)

## [v2.11.0](https://github.com/vultr/govultr/compare/v2.10.0..v2.11.0) (2021-11-18)
### Breaking Changes
* Instances : Update call will now return `*Instance` in addition to `error` [185](https://github.com/vultr/govultr/pull/185)
* Instances : Reinstall call now allows changing of hostname and also returns `*Instance` in addition to `error` [181](https://github.com/vultr/govultr/pull/181)

### Enhancement
* Instances : The hostname of the instance is now returned in any call that returns Instance data [187](https://github.com/vultr/govultr/pull/187)
* Domains : There is a new field called `dns_sec` which will return `enabled` or `disabled` depending on how your domain is configured [184](https://github.com/vultr/govultr/pull/184)

## [v2.10.0](https://github.com/vultr/govultr/compare/v2.9.2..v2.10.0) (2021-11-03)
### Enhancement
* Billing : Added support for billing [178](https://github.com/vultr/govultr/pull/178)

## [v2.9.2](https://github.com/vultr/govultr/compare/v2.9.1..v2.9.2) (2021-10-20)
### Change
* Iso : Changed `client` field to be unexported [168](https://github.com/vultr/govultr/pull/168)
* Snapshot : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Plans : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Regions : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)

## [v2.9.1](https://github.com/vultr/govultr/compare/v2.9.0..v2.9.1) (2021-10-18)
### Enhancement
* Kubernetes : Added `Tag` support for nodepools [164](https://github.com/vultr/govultr/pull/164)

## [v2.9.0](https://github.com/vultr/govultr/compare/v2.8.1..v2.9.0) (2021-09-27)
### Breaking Change
* Kubernetes : PlanID is now Plan and Count is now NodeQuantity to follow API pattern [161](https://github.com/vultr/govultr/pull/161)

### Enhancement
* Snapshots : Add compressed size field [162](https://github.com/vultr/govultr/pull/162)

## [v2.8.1](https://github.com/vultr/govultr/compare/v2.8.0..v2.8.1) (2021-08-31)
### Enhancement
* Kubernetes : Add support for deletion with resources [159](https://github.com/vultr/govultr/pull/159)
* Kubernetes : Add support for getting available versions[159](https://github.com/vultr/govultr/pull/159)

### Dependency Update
* Bump Go version to 1.16 [158](https://github.com/vultr/govultr/pull/158)

## [v2.8.0](https://github.com/vultr/govultr/compare/v2.7.1..v2.8.0) (2021-08-18)
### Enhancement
* Added support for Vultr Kubernetes Engine [156](https://github.com/vultr/govultr/pull/156)

## [v2.7.1](https://github.com/vultr/govultr/compare/v2.7.0..v2.7.1) (2021-07-15)
### Enhancement
* BareMetal : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)
* Instances : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)

## [v2.7.0](https://github.com/vultr/govultr/compare/v2.6.0..v2.7.0) (2021-07-14)
### Enhancement
* BareMetal : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Instances : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Applications : added support for marketplace applications [150](https://github.com/vultr/govultr/pull/150)

## [v2.6.0](https://github.com/vultr/govultr/compare/v2.5.1..v2.6.0) (2021-07-02)
### Enhancement
* BareMetal : Add support for `persistent_pxe` [148](https://github.com/vultr/govultr/pull/148)

>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
## [v2.14.1](https://github.com/vultr/govultr/compare/v2.14.0..v2.14.1) (2022-02-02)
### Enhancement
* Improved retry error response [204](https://github.com/vultr/govultr/pull/204)

## [v2.14.0](https://github.com/vultr/govultr/compare/v2.13.0..v2.14.0) (2022-01-21)
### Enhancement
* ListOptions : [Added query param Region](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [200](https://github.com/vultr/govultr/pull/200)
* ListOptions : [Added query param Description](https://www.vultr.com/api/#operation/list-snapshots) that can be used with `Snapshot.List`  [202](https://github.com/vultr/govultr/pull/202)
* Snapshot : `CreateFromURL` has new optional field called `description` which lets you set a custom description [202](https://github.com/vultr/govultr/pull/202)

## [v2.13.0](https://github.com/vultr/govultr/compare/v2.12.0..v2.13.0) (2022-01-05)
### Enhancement
* ListOptions : [Added query params](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [197](https://github.com/vultr/govultr/pull/197)

## [v2.12.0](https://github.com/vultr/govultr/compare/v2.11.1..v2.12.0) (2021-12-01)
### Breaking Changes
* Plans : Changed `MonthlyCost` from `int` to `float32` [192](https://github.com/vultr/govultr/pull/192)

## [v2.11.1](https://github.com/vultr/govultr/compare/v2.11.0..v2.11.1) (2021-11-26)
### Bug fixes
* LoadBalancers : Fixed SSL struct json params to the proper API fields [189](https://github.com/vultr/govultr/pull/189)

## [v2.11.0](https://github.com/vultr/govultr/compare/v2.10.0..v2.11.0) (2021-11-18)
### Breaking Changes
* Instances : Update call will now return `*Instance` in addition to `error` [185](https://github.com/vultr/govultr/pull/185)
* Instances : Reinstall call now allows changing of hostname and also returns `*Instance` in addition to `error` [181](https://github.com/vultr/govultr/pull/181)

### Enhancement
* Instances : The hostname of the instance is now returned in any call that returns Instance data [187](https://github.com/vultr/govultr/pull/187)
* Domains : There is a new field called `dns_sec` which will return `enabled` or `disabled` depending on how your domain is configured [184](https://github.com/vultr/govultr/pull/184)

## [v2.10.0](https://github.com/vultr/govultr/compare/v2.9.2..v2.10.0) (2021-11-03)
### Enhancement
* Billing : Added support for billing [178](https://github.com/vultr/govultr/pull/178)

## [v2.9.2](https://github.com/vultr/govultr/compare/v2.9.1..v2.9.2) (2021-10-20)
### Change
* Iso : Changed `client` field to be unexported [168](https://github.com/vultr/govultr/pull/168)
* Snapshot : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Plans : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Regions : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)

## [v2.9.1](https://github.com/vultr/govultr/compare/v2.9.0..v2.9.1) (2021-10-18)
### Enhancement
* Kubernetes : Added `Tag` support for nodepools [164](https://github.com/vultr/govultr/pull/164)

## [v2.9.0](https://github.com/vultr/govultr/compare/v2.8.1..v2.9.0) (2021-09-27)
### Breaking Change
* Kubernetes : PlanID is now Plan and Count is now NodeQuantity to follow API pattern [161](https://github.com/vultr/govultr/pull/161)

### Enhancement
* Snapshots : Add compressed size field [162](https://github.com/vultr/govultr/pull/162)

## [v2.8.1](https://github.com/vultr/govultr/compare/v2.8.0..v2.8.1) (2021-08-31)
### Enhancement
* Kubernetes : Add support for deletion with resources [159](https://github.com/vultr/govultr/pull/159)
* Kubernetes : Add support for getting available versions[159](https://github.com/vultr/govultr/pull/159)

### Dependency Update
* Bump Go version to 1.16 [158](https://github.com/vultr/govultr/pull/158)

## [v2.8.0](https://github.com/vultr/govultr/compare/v2.7.1..v2.8.0) (2021-08-18)
### Enhancement
* Added support for Vultr Kubernetes Engine [156](https://github.com/vultr/govultr/pull/156)

## [v2.7.1](https://github.com/vultr/govultr/compare/v2.7.0..v2.7.1) (2021-07-15)
### Enhancement
* BareMetal : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)
* Instances : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)

## [v2.7.0](https://github.com/vultr/govultr/compare/v2.6.0..v2.7.0) (2021-07-14)
### Enhancement
* BareMetal : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Instances : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Applications : added support for marketplace applications [150](https://github.com/vultr/govultr/pull/150)

## [v2.6.0](https://github.com/vultr/govultr/compare/v2.5.1..v2.6.0) (2021-07-02)
### Enhancement
* BareMetal : Add support for `persistent_pxe` [148](https://github.com/vultr/govultr/pull/148)

>>>>>>> 4d7e5ad26 (update vendored files)
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
<<<<<<< HEAD
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
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
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
* GoVultr v1.0.0 is now on [branch v1](https://github.com/vultr/govultr/tree/v1)
=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
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
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
