# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.1] - 2019-10-11

- Adds TTL to RRSetChange (enables support for custom RRSet TTLs)

## [1.1.0] - 2019-01-23

### Added

- Pagination support and ListOption struct
- Encryption support for RRSetService

### Changed

- Client services (Zones, RRSet) to implement interface

## [1.0.0] - 2019-01-15

- Initial stable release

[Unreleased]: https://github.com/nic-at/rc0go/compare/v1.0.0...HEAD
[1.1.1]: https://github.com/nic-at/rc0go/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/nic-at/rc0go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/nic-at/rc0go/compare/v1.0.0...v1.0.0
