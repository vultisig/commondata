# commondata — Agent Reference

## Overview

Shared protobuf schemas defining the data contract between all Vultisig apps. Changes here affect iOS, Android, Windows, and all Go services.

## Quick Start

```bash
git clone https://github.com/vultisig/commondata.git
cd commondata
make                    # runs buf via Docker (no local buf needed)
# Or locally: ./build.sh (requires buf CLI)
```

## Before You Change Code

1. This is a submodule used by iOS, Android, and Go services — changes ripple everywhere
2. Run `make` (or `./build.sh`) to validate current state
3. Only edit files in `proto/vultisig/` — never touch `Sources/swift/` or `go/vultisig/` (generated)
4. After proto changes: run `buf generate` and commit generated code alongside proto changes
5. Breaking changes affect ALL platforms — deprecate fields, never remove them

## Patterns

- Proto3 syntax, package `vultisig.v1`
- Generated code: Sources/swift/ (Swift), go/vultisig/ (Go) — never edit generated files
- Token metadata in tokens/ directory
- buf.yaml for linting rules, buf.gen.yaml for code generation

## Security Notes

- Proto field removal is a breaking change — deprecate instead
- Generated code must be committed alongside proto changes
- Token metadata affects balance display — validate carefully
