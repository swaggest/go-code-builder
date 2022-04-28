# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.50] - 2022-04-28

### Fixed

- Exhaustive boolean values elision in const block

## [0.4.49] - 2022-04-28

### Fixed

- Boolean values handling in enum/const block 

## [0.4.48] - 2022-01-02

### Added
- Support for PHP 8.1

## [0.4.47] - 2022-01-02

### Fixed
- Code generation for AsyncAPI 2.1.0 schema

## [0.4.46] - 2021-08-29

### Fixed
- Infinite recursion with JSON Schema draft-07 reference in 3rd party schema.

## [0.4.45] - 2021-07-16

### Fixed
- Enum and const collision with type name.

## [0.4.44] - 2021-04-20

### Fixed
- Missing imports in generated code.

## [0.4.43] - 2021-04-20

### Fixed
- Invalid escaping of single quot.
- Type naming with schema title.

## [0.4.42] - 2020-12-12

### Fixed
- Lint issues

## [0.4.41] - 2020-12-12

### Added
- Properties embedding for `allOf` schemas that only serve plain properties.

## [0.4.40] - 2020-09-30

### Added
- Option to set additional field tags with property name.

### Fixed
- Redundant pointer type `*interface{}` on non-omitted fields.

## [0.4.39] - 2020-08-09

### Fixed
- Redundant `const` check for missing properties.

## [0.4.38] - 2020-04-30

### Fixed
- Rendering of nullable+omitempty slices and maps.

## [0.4.37] - 2020-04-29

### Fixed
- Rendering of required x-nullable fields.

## [0.4.36] - 2020-04-29

### Fixed
- Rendering of required nullable fields.

## [0.4.35] - 2020-04-28

### Added
- Option to disable `required` validation during unmarshal.

## [0.4.34] - 2020-04-12

### Fixed
- Rendering of type `["null", "object"]`.
- Redundant `marshalUnion` call removed.
- Byte slice preallocation with magic number removed in favor of dynamic allocation.

## [0.4.33] - 2020-04-08

### Added
- Field ensurer generation when fluent setters are enabled.

## [0.4.32] - 2020-04-04

### Added
- Support for simplified syntax of `x-go-type: foo-go::foo.Type`.

## [0.4.31] - 2020-04-03

### Added
- Skip generation of properties with `x-generate: false`.
- Option to only generate properties with `x-generate: true`.

## [0.4.30] - 2020-03-30

### Added
- Argument to `MarshalingTestFunc` to avoid default additional properties in tests.

## [0.4.29] - 2020-03-17

### Added
- Option to rename generated symbol names.

## [0.4.28] - 2020-03-10

### Added
- Added support for schema logic resolution based on `oneOf`, `anyOf`, `required`, `not`, `additionalPrperties`.
- Added helper to generate tests for entities.

## [0.4.27] - 2020-02-25

### Added
- `GoBuilder` option to `ignoreRequired` when deciding on pointer type or omitempty.
- Improved receiver naming to use first char of type.

## [0.4.26] - 2020-02-18

### Fixed
- Few code style issues.
- Pointer and omitempty removed for required properties.

## [0.4.25] - 2020-02-04

### Added
- Fluent setter for map item in JSON schema generated structures.

## [0.4.24] - 2020-02-02

### Fixed
- Wrong field referencing of multi-type structures, introduced in `0.4.23`.

## [0.4.23] - 2020-02-02

### Fixed
- Missing fields for multi-type structures.

## [0.4.22] - 2020-01-26

### Added
- Early return in `MarshalJSON` if there are no additional properties to improve performance.
- JSON Schema maker library updated.

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

[0.4.50]: https://github.com/swaggest/go-code-builder/compare/v0.4.49...v0.4.50
[0.4.49]: https://github.com/swaggest/go-code-builder/compare/v0.4.48...v0.4.49
[0.4.48]: https://github.com/swaggest/go-code-builder/compare/v0.4.47...v0.4.48
[0.4.47]: https://github.com/swaggest/go-code-builder/compare/v0.4.46...v0.4.47
[0.4.46]: https://github.com/swaggest/go-code-builder/compare/v0.4.45...v0.4.46
[0.4.45]: https://github.com/swaggest/go-code-builder/compare/v0.4.44...v0.4.45
[0.4.44]: https://github.com/swaggest/go-code-builder/compare/v0.4.43...v0.4.44
[0.4.43]: https://github.com/swaggest/go-code-builder/compare/v0.4.42...v0.4.43
[0.4.42]: https://github.com/swaggest/go-code-builder/compare/v0.4.41...v0.4.42
[0.4.41]: https://github.com/swaggest/go-code-builder/compare/v0.4.40...v0.4.41
[0.4.40]: https://github.com/swaggest/go-code-builder/compare/v0.4.39...v0.4.40
[0.4.39]: https://github.com/swaggest/go-code-builder/compare/v0.4.38...v0.4.39
[0.4.38]: https://github.com/swaggest/go-code-builder/compare/v0.4.37...v0.4.38
[0.4.37]: https://github.com/swaggest/go-code-builder/compare/v0.4.36...v0.4.37
[0.4.36]: https://github.com/swaggest/go-code-builder/compare/v0.4.35...v0.4.36
[0.4.35]: https://github.com/swaggest/go-code-builder/compare/v0.4.34...v0.4.35
[0.4.34]: https://github.com/swaggest/go-code-builder/compare/v0.4.33...v0.4.34
[0.4.33]: https://github.com/swaggest/go-code-builder/compare/v0.4.32...v0.4.33
[0.4.32]: https://github.com/swaggest/go-code-builder/compare/v0.4.31...v0.4.32
[0.4.31]: https://github.com/swaggest/go-code-builder/compare/v0.4.30...v0.4.31
[0.4.30]: https://github.com/swaggest/go-code-builder/compare/v0.4.29...v0.4.30
[0.4.29]: https://github.com/swaggest/go-code-builder/compare/v0.4.28...v0.4.29
[0.4.28]: https://github.com/swaggest/go-code-builder/compare/v0.4.27...v0.4.28
[0.4.27]: https://github.com/swaggest/go-code-builder/compare/v0.4.26...v0.4.27
[0.4.26]: https://github.com/swaggest/go-code-builder/compare/v0.4.25...v0.4.26
[0.4.25]: https://github.com/swaggest/go-code-builder/compare/v0.4.24...v0.4.25
[0.4.24]: https://github.com/swaggest/go-code-builder/compare/v0.4.23...v0.4.24
[0.4.23]: https://github.com/swaggest/go-code-builder/compare/v0.4.22...v0.4.23
[0.4.22]: https://github.com/swaggest/go-code-builder/compare/v0.4.21...v0.4.22
[0.4.21]: https://github.com/swaggest/go-code-builder/compare/v0.4.20...v0.4.21
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
