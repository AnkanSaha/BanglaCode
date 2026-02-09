# Changelog

All notable changes to BanglaCode will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- String interpolation support
- Lambda/arrow functions
- Enhanced standard library
- Improved error messages with suggestions

---

## [3.3.0] - 2024-XX-XX

### Added
- Enhanced playground with selectable examples
- Refactored documentation sidebar configuration
- Updated application icons across all platforms
- Improved VSCode extension with better IntelliSense

### Changed
- Documentation website now uses Next.js 16.1.6
- Upgraded React to version 19.2.3
- Improved sidebar navigation in documentation

### Fixed
- Various UI improvements in the playground
- Documentation rendering issues

---

## [3.2.2] - 2024-XX-XX

### Changed
- Version update across all components
- Synchronized version numbers in package.json files

### Fixed
- Minor bug fixes and stability improvements

---

## [3.2.1] - 2024-XX-XX

### Added
- New documentation website with modern design
- Cloudflare Workers deployment support via Wrangler
- Updated versioning script for consistent releases

### Changed
- Improved documentation organization
- Better mobile responsiveness for docs site

---

## [3.2.0] - 2024-XX-XX

### Added
- HTTP server functionality (`server_chalu`)
- JSON response helper (`json_uttor`)
- HTTP client function (`anun`)
- Request/response object handling

### Changed
- Improved HTTP error handling
- Better content-type detection

### Fixed
- Memory leaks in long-running servers
- Connection handling improvements

---

## [3.1.0] - 2024-XX-XX

### Added
- Full JSON support with `json_poro` and `json_banao`
- JSON file import capability
- Better object/map handling

### Changed
- Improved parser for JSON-like object syntax
- Enhanced string escaping in JSON output

### Fixed
- Unicode handling in JSON strings
- Nested object parsing issues

---

## [3.0.0] - 2024-XX-XX

### Added
- **Module System** — Import/export with `ano`/`pathao` keywords
- **Alias imports** — `ano "module.bang" hisabe alias;`
- **Error Handling** — Try/catch/finally with `chesta`/`dhoro_bhul`/`shesh`
- **Throw statement** — `felo` for custom errors
- VSCode extension with full IDE support
- 35+ code snippets
- Hover documentation support

### Changed
- Complete rewrite of the evaluator for better performance
- Improved error messages with line and column numbers
- Better memory management

### Breaking Changes
- Module imports now require explicit exports
- Error objects have new structure

### Fixed
- Stack overflow in deep recursion
- Variable scoping issues in nested functions

---

## [2.1.0] - 2024-XX-XX

### Added
- **Classes and OOP** — `sreni`, `shuru`, `notun`, `ei` keywords
- Method definitions inside classes
- Constructor support
- Instance property access

### Changed
- Enhanced object system for class instances
- Improved `this` (`ei`) binding

### Fixed
- Method resolution in nested classes
- Property assignment issues

---

## [2.0.0] - 2024-XX-XX

### Added
- **For loops** — `ghuriye` keyword
- **Break and continue** — `thamo` and `chharo`
- **Compound assignment** — `+=`, `-=`, `*=`, `/=`
- File I/O functions (`poro`, `lekho`)
- More math functions

### Changed
- Major performance improvements (3-4x faster than v1.x)
- Rewritten lexer for better token handling
- Improved parser error recovery

### Breaking Changes
- Changed some built-in function names for consistency
- Updated operator precedence

### Fixed
- Loop variable scoping
- Operator precedence issues
- Memory usage in large programs

---

## [1.2.0] - 2024-XX-XX

### Added
- Array functions: `dhokao`, `berKoro`, `kato`, `ulto`, `saja`
- String functions: `boroHater`, `chotoHater`, `chhanto`
- `ache` function for membership testing

### Changed
- Improved array performance
- Better string handling

### Fixed
- Array index out of bounds errors
- String concatenation issues

---

## [1.1.0] - 2024-XX-XX

### Added
- **Maps/Objects** — Key-value data structure
- `chabi` function to get map keys
- Dot notation for property access
- Bracket notation for dynamic access

### Changed
- Enhanced parser for object literals
- Improved REPL with multi-line support

### Fixed
- Parser errors with nested structures
- REPL input handling

---

## [1.0.0] - 2024-XX-XX

### Added
- **Initial release of BanglaCode**
- Core language features:
  - Variables with `dhoro`
  - Conditionals with `jodi`/`nahole`
  - While loops with `jotokkhon`
  - Functions with `kaj`/`ferao`
  - Logical operators: `ebong`, `ba`, `na`
  - Boolean values: `sotti`, `mittha`
  - Null value: `khali`
- Basic built-in functions:
  - `dekho` — Print output
  - `dhoron` — Get type
  - `lipi` — Convert to string
  - `sonkha` — Convert to number
  - `dorghyo` — Get length
- Interactive REPL
- Basic error messages
- Example programs

### Notes
- First public release
- Designed for Bengali-speaking students
- Written in Go for performance

---

## Version History Summary

| Version | Release Date | Highlights |
|---------|--------------|------------|
| 3.3.0 | 2024 | Playground enhancements, documentation updates |
| 3.2.x | 2024 | Documentation website, versioning improvements |
| 3.1.0 | 2024 | JSON support |
| 3.0.0 | 2024 | Modules, error handling, VSCode extension |
| 2.1.0 | 2024 | Classes and OOP |
| 2.0.0 | 2024 | For loops, file I/O, major performance boost |
| 1.x.x | 2024 | Arrays, maps, string functions |
| 1.0.0 | 2024 | Initial release |

---

## Upgrade Guides

### Upgrading to 3.x from 2.x

1. **Module exports**: Add `pathao` keyword to functions you want to export
2. **Error handling**: Consider using `chesta`/`dhoro_bhul` for better error management
3. **VSCode**: Install the new extension for better IDE support

### Upgrading to 2.x from 1.x

1. **Built-in functions**: Some function names changed for consistency
2. **Operators**: Review code using compound assignment (now supported)
3. **Loops**: Consider using `ghuriye` for index-based loops

---

## Links

- [GitHub Releases](https://github.com/nexoral/BanglaCode/releases)
- [Migration Guides](https://banglacode.dev/docs/migration)
- [Full Documentation](https://banglacode.dev/docs)

---

[Unreleased]: https://github.com/nexoral/BanglaCode/compare/v3.3.0...HEAD
[3.3.0]: https://github.com/nexoral/BanglaCode/compare/v3.2.2...v3.3.0
[3.2.2]: https://github.com/nexoral/BanglaCode/compare/v3.2.1...v3.2.2
[3.2.1]: https://github.com/nexoral/BanglaCode/compare/v3.2.0...v3.2.1
[3.2.0]: https://github.com/nexoral/BanglaCode/compare/v3.1.0...v3.2.0
[3.1.0]: https://github.com/nexoral/BanglaCode/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/nexoral/BanglaCode/compare/v2.1.0...v3.0.0
[2.1.0]: https://github.com/nexoral/BanglaCode/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/nexoral/BanglaCode/compare/v1.2.0...v2.0.0
[1.2.0]: https://github.com/nexoral/BanglaCode/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/nexoral/BanglaCode/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/nexoral/BanglaCode/releases/tag/v1.0.0
