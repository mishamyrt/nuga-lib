# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].


## [v0.6.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.6.0) - 2024-03-19
### Features
- add Halo65 keys support


## [v0.5.8](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.8) - 2024-03-01
### Bug Fixes
- use original state write order
- update default Halo75 state
- set correct keys request length

### Refactoring
- make hex printers more testable

### Testing
- add hex


## [v0.5.7](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.7) - 2024-02-29
### Bug Fixes
- use interface for dumps


## [v0.5.6](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.6) - 2024-02-29
### Features
- export handle


## [v0.5.5](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.5) - 2024-02-28
### Bug Fixes
- remove extra logging


## [v0.5.4](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.4) - 2024-02-27
### Features
- add missing codes


## [v0.5.3](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.3) - 2024-02-25
### Bug Fixes
- implement parse in keys simulation


## [v0.5.2](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.2) - 2024-02-25
### Bug Fixes
- add parse to interface


## [v0.5.1](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.1) - 2024-02-25
### Features
- add keys parsing utility


## [v0.5.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.5.0) - 2024-02-25
### Features
- provide dump handler, rework dump format


## [v0.4.2](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.4.2) - 2024-02-24
### Features
- add keys simulation support


## [v0.4.1](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.4.1) - 2024-02-23
### Refactoring
- update keys format, add missing keys

### Testing
- add keymap tests


## [v0.4.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.4.0) - 2024-02-23
### Features
- add keyboard layout map support
- provide keys capability for Halo75
- add raw keys support

### Testing
- add keymap tests


## [v0.3.1](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.3.1) - 2024-01-05
### Bug Fixes
- use correct offset for mac os color change

### Refactoring
- use constants


## [v0.3.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.3.0) - 2024-01-03

## [v0.2.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.2.0) - 2024-01-03
### Features
- add `describe` example
- rework dumper to simulation util (read, dump)

### Refactoring
- simplify imports
- rename git-chglog version variable
- move version handling to `device`
- move simulations to testdata

### Testing
- add tests for simulation and capability
- update templates for new format


## [v0.1.0](https://github.com/mishamyrt/nuga-lib/releases/tag/v0.1.0) - 2023-12-30
### CI
- split lint and test
- add qa workflow

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
[Unreleased]: https://github.com/mishamyrt/nuga-lib/compare/v0.6.0...HEAD
[v0.6.0]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.8...v0.6.0
[v0.5.8]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.7...v0.5.8
[v0.5.7]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.6...v0.5.7
[v0.5.6]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.5...v0.5.6
[v0.5.5]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.4...v0.5.5
[v0.5.4]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.3...v0.5.4
[v0.5.3]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.2...v0.5.3
[v0.5.2]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.1...v0.5.2
[v0.5.1]: https://github.com/mishamyrt/nuga-lib/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/mishamyrt/nuga-lib/compare/v0.4.2...v0.5.0
[v0.4.2]: https://github.com/mishamyrt/nuga-lib/compare/v0.4.1...v0.4.2
[v0.4.1]: https://github.com/mishamyrt/nuga-lib/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/mishamyrt/nuga-lib/compare/v0.3.1...v0.4.0
[v0.3.1]: https://github.com/mishamyrt/nuga-lib/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/mishamyrt/nuga-lib/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/mishamyrt/nuga-lib/compare/v0.1.0...v0.2.0
