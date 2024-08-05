# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]
## [1.3.7] 
BUG FIXES:
* Resource Rrset:
  - Added "TTL" field in DirPoolProfile to make it compatible with DirPoolDTO.

## [1.3.6] - 2020-06-17
ENHANCEMENTS:
* Resource Zone:
  - Zone DTO now includes Zone Services that enables reading (GET) of zones.
  - Special Characters are now allowed when reading (GET) resources.
  - Integration and Unit testcases have been added.

* Resource Rrset:
  - Now supports the use of “limits” and “offset” for pagination. 
  - Client Name is now used as an identifier during create (POST) calls.
  - Added interfaces in RRSet service as to make it mockable in unit tests.
  - Added additional unit testcases to comply with coding conventions.

* Resource Udnssdk:
  - Integration tests for zone service can now be enabled.

BUG FIXES:
* Resource Rrset 
  - Special characters are now supported when creating (POST) and reading (GET) other resources.
  - In the SBPoolProfile, the weight parameter will not be set by default when left empty.
* Resource Udnssdk
  - API version dependency has been removed allowing for compatibility with latest API.

## [1.3.5] - 2020-04-14
- Added 'availableToServe' to BackupRecord DTO
- Added 'status' to TCPool profile DTO
- Improved the testing steps used latest `go 1.14` and `go mod`

## [1.3.4] - 2018-01-15
### Changed
- Update all references to udnssdk from Ensighten to terra-farm

## [1.3.3] - 2017-12-19
### Changed
- Added 'availableToServe' to SBPool and TCPool profile DTOs.

## [1.3.2] - 2017-03-03
### Changed
- CheckResponse: improve fallthrough error to include full Response Status and properly format Body
- Client.NewRequest: split query to avoid encoding "?" as "%3F" into path

## [1.3.1] - 2017-03-03
### Changed
- Client.NewRequest: shallow-copy BaseURL to avoid retaining modifications

## [1.3.0] - 2017-02-28
### Added
- cmd/udns: add rrset query tool
- DPRDataInfo.Type: add field to support API change

## [1.2.1] - 2016-06-13
### Fixed
* `omitempty` tags fixed for `ProbeInfoDTO.PoolRecord` & `ProbeInfoDTO.ID`
* Check `*http.Response` values for nil before access

## [1.2.0] - 2016-06-09
### Added
* Add probe detail serialization helpers

### Changed
* Flatten udnssdk.Response to mere http.Response
* Extract self-contained passwordcredentials oauth2 TokenSource
* Change ProbeTypes to constants

## [1.1.1] - 2016-05-27
### Fixed
* remove terraform tag for `GeoInfo.Codes`

## [1.1.0] - 2016-05-27
### Added
* Add terraform tags to structs to support mapstructure

### Fixed
* `omitempty` tags fixed for `DirPoolProfile.NoResponse`, `DPRDataInfo.GeoInfo`, `DPRDataInfo.IPInfo`, `IPInfo.Ips` & `GeoInfo.Codes`
* ProbeAlertDataDTO equivalence for times with different locations

### Changed
* Convert RawProfile to use mapstructure and structs instead of round-tripping through json
* CHANGELOG.md: fix link to v1.0.0 commit history

## [1.0.0] - 2016-05-11
### Added
* Support for API endpoints for `RRSets`, `Accounts`,  `DirectionalPools`, Traffic Controller Pool `Probes`, `Events`, `Notifications` & `Alerts`
* `Client` wraps common API access including OAuth, deferred tasks and retries


[1.3.2]: https://github.com/Ensighten/udnssdk/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/Ensighten/udnssdk/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/Ensighten/udnssdk/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/Ensighten/udnssdk/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/Ensighten/udnssdk/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/Ensighten/udnssdk/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/Ensighten/udnssdk/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/Ensighten/udnssdk/compare/v0.0.0...v1.0.0
