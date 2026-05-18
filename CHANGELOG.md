# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- 

### Changed

- 

### Removed

- 

## [1.0.1] - 2026-05-18

### Added

- Added `HealthInfo` on practice health request.

### Changed

- Rectified URL for practice health endpoint.

### Removed

- Deprecated eventresult health endpoint as it only returns an 404.

## [1.0.0] - 2025-11-19

### Added

- Added this CHANGELOG.md.
- Added the API definitions this library was build against.
- Added missing `practice.Accounts.Activities` endpoint.

### Removed

- Removed `practice.Admin` domain as it is unclear how this should work.
- Removed deprecated functions from `eventresults` domain.

[unreleased]: https://github.com/ysmilda/speedhive-go/compare/v1.0.1...HEAD
[1.0.1]: https://github.com/ysmilda/speedhive-go/releases/tag/v1.0.1
[1.0.0]: https://github.com/ysmilda/speedhive-go/releases/tag/v1.0.0