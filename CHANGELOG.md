# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.21] - 2020-01-24

### Fixed
- Missing `encoding/json` import in some cases.

## [0.4.20] - 2020-01-04

### Added
- Fluent field setter template.
- Option to build fluent setters when generating structures from JSON schema.

### Fixed
- Missing arguments/results handling in func definition.
- Duplicated functions in struct.
- Variadic syntax in func argument.

## [0.4.19] - 2019-11-18

### Fixed
- Indentation of pattern properties loop.
- Removed `composer.lock` from exported distribution.

## [0.4.18] - 2019-11-16

### Added
- Optimized memory allocation efficiency in unmarshaling code.

## [0.4.17] - 2019-10-27

### Fixed
- Removed redundant marshaling helper type.

## [0.4.16] - 2019-10-25

### Fixed
- Redundant tabs in `marshalUnion`.

## [0.4.15] - 2019-10-25

### Fixed
- Redundant tabs in `marshalUnion`.

## [0.4.14] - 2019-10-23

### Fixed
- Redundant duplicated early return check.

## [0.4.13] - 2019-10-23

### Changed
- Whitespace formatting for improved readability.

## [0.4.12] - 2019-10-22

### Changed
- Names for logical branches are inherited from reference names where available.

### Added
- Option to `$inheritSchemaFromExamples` when type and properties are missing.

### Fixed
- Invalid property name when stripping parent prefix.

## [0.4.11] - 2019-10-15

### Fixed
- Dereference of `nil` `*json.RawMessage` in additional and pattern properties.

## [0.4.10] - 2019-10-15

### Added
- Support for `x-omitempty` vendor extension.
- Vendor extensions documentation in README.md.

## [0.4.9] - 2019-10-10

### Added
- Option to disable null `additionalProperties` rendering.

## [0.4.8] - 2019-10-09

### Changed
- Constant/enum type names are prefixed with parent property if available.
- Constant/enum item names are prefixed with parent type instead of path. 
- Schema title is used for struct name instead of path where available.

### Added
- Optional support for `x-nullable` (Swagger 2.0), `nullable` (OpenAPI 3.0) properties to enable nullability.

### Fixed
- Type marker stripping from path was affecting regexps too.

## [0.4.7] - 2019-10-05

### Added
- Support for `x-go-type` as object (`go-swagger` format).
- Builder option `$ignoreXGoType` to disregard `x-go-type` hints.
- Builder option `$withZeroValues` to use pointer types to avoid zero-value ambiguity.
- Builder option `$ignoreNullable` to enable `omitempty` for nullable properties.
- Automated API docs.
- Change log.

### Fixed
- Missing processing for `null` `additionalProperties`.
- Processing for nullable (`[null, <other>]`) types.

### Changed
- Multi-type schemas are decomposed into multiple structures.
- Schema title is used for structure name if available.

## [0.4.6] - 2019-10-01

### Added
- More tests.
- CI upgrade.

### Fixed
- Code-style issues.

## [0.4.5] - 2019-07-11

### Added
- Type inference from enum values.

### Changed
- Trivial nesting removed. 

## [0.4.4] - 2019-07-08

### Fixed
- Removed unnecessary regexp dependency, #7.

[0.4.20]: https://github.com/swaggest/go-code-builder/compare/v0.4.19...v0.4.20
[0.4.19]: https://github.com/swaggest/go-code-builder/compare/v0.4.18...v0.4.19
[0.4.18]: https://github.com/swaggest/go-code-builder/compare/v0.4.17...v0.4.18
[0.4.17]: https://github.com/swaggest/go-code-builder/compare/v0.4.16...v0.4.17
[0.4.16]: https://github.com/swaggest/go-code-builder/compare/v0.4.15...v0.4.16
[0.4.15]: https://github.com/swaggest/go-code-builder/compare/v0.4.14...v0.4.15
[0.4.14]: https://github.com/swaggest/go-code-builder/compare/v0.4.13...v0.4.14
[0.4.13]: https://github.com/swaggest/go-code-builder/compare/v0.4.12...v0.4.13
[0.4.12]: https://github.com/swaggest/go-code-builder/compare/v0.4.11...v0.4.12
[0.4.11]: https://github.com/swaggest/go-code-builder/compare/v0.4.10...v0.4.11
[0.4.10]: https://github.com/swaggest/go-code-builder/compare/v0.4.9...v0.4.10
[0.4.9]: https://github.com/swaggest/go-code-builder/compare/v0.4.8...v0.4.9
[0.4.8]: https://github.com/swaggest/go-code-builder/compare/v0.4.7...v0.4.8
[0.4.7]: https://github.com/swaggest/go-code-builder/compare/v0.4.6...v0.4.7
[0.4.6]: https://github.com/swaggest/go-code-builder/compare/v0.4.5...v0.4.6
[0.4.5]: https://github.com/swaggest/go-code-builder/compare/v0.4.4...v0.4.5
[0.4.4]: https://github.com/swaggest/go-code-builder/compare/v0.4.3...v0.4.4
