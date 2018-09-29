# Change Log

<a name="0.3.0"></a>

## [0.3.0](https://github.com/linode/linodego/compare/v0.2.0...0.3.0) (2018-08-15)

### Breaking Changes

* WaitForVolumeLinodeID return fetch volume for consistency with out WaitFors
* Moved linodego from chiefy to github.com/linode. Thanks [@chiefy](https://github.com/chiefy)!

<a name="v0.2.0"></a>

## [v0.2.0](https://github.com/linode/linodego/compare/v0.1.1...v0.2.0) (2018-08-11)

### Breaking Changes

* WaitFor\* should be client methods
  *use `client.WaitFor...` rather than `linodego.WaitFor(..., client, ...)`*

* remove ListInstanceSnapshots (does not exist in the API)
  *this never worked, so shouldn't cause a problem*

* Changes UpdateOptions and CreateOptions and similar Options parameters to values instead of pointers
  *these were never optional and the function never updated any values in the Options structures*

* fixed various optional/zero Update and Create options
  *some values are now pointers, and vice-versa*

  * Changes InstanceUpdateOptions to use pointers for optional fields Backups and Alerts
  * Changes InstanceClone's Disks and Configs to ints instead of strings

* using new enum string aliased types where appropriate
  *`InstanceSnapshotStatus`, `DiskFilesystem`, `NodeMode`*

### Feature

* add RescueInstance and RescueInstanceOptions
* add CreateImage, UpdateImage, DeleteImage
* add EnableInstanceBackups, CancelInstanceBackups, RestoreInstanceBackup
* add WatchdogEnabled to InstanceUpdateOptions

### Fix

* return Volume from AttachVolume instead of bool
* add more boilerplate to template.go
* nodebalancers and domain records had no pagination support
* NodeBalancer transfer stats are not int

### Tests

* add fixtures and tests for NodeBalancerNodes
* fix nodebalancer tests to handle changes due to random labels
* add tests for nodebalancers and nodebalancer configs
* added tests for Backups flow
* TestListInstanceBackups fixture is hand tweaked because repeated polled events
  appear to get the tests stuck

### Deps

* update all dependencies to latest

<a name="v0.1.1"></a>

## [v0.1.1](https://github.com/linode/linodego/compare/v0.0.1...v0.1.0) (2018-07-30)

Adds more Domain handling

### Fixed

* go-resty doesnt pass errors when content-type is not set
* Domain, DomainRecords, tests and fixtures

### Added

* add CreateDomainRecord, UpdateDomainRecord, and DeleteDomainRecord

<a name="v0.1.0"></a>

## [v0.1.0](https://github.com/linode/linodego/compare/v0.0.1...v0.1.0) (2018-07-23)

Deals with NewClient and context for all http requests

### Breaking Changes

* changed `NewClient(token, *http.RoundTripper)` to `NewClient(*http.Client)`
* changed all `Client` `Get`, `List`, `Create`, `Update`, `Delete`, and `Wait` calls to take context as the first parameter

### Fixed

* fixed docs should now show Examples for more functions

### Added

* added `Client.SetBaseURL(url string)`

<a name="v0.0.1"></a>
## v0.0.1 (2018-07-20)

### Changed

* Initial tagged release
