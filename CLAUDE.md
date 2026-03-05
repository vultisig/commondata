# commondata

Shared protobuf schemas for cross-platform communication. Source of truth for data contracts between all Vultisig apps.

## Security Tier

HIGH

## Critical Boundaries

- `proto/vultisig/` — Source of truth. Any change here requires regeneration on ALL platforms.
- `Sources/swift/` and `go/vultisig/` — Generated code. Do NOT edit manually.
- Breaking proto changes affect iOS, Android, Windows, and all Go services.

## Key Commands

```bash
# Full workflow via Docker (no local buf needed)
make                    # runs buf format, lint, build, generate in Docker

# Or locally if buf is installed:
buf format --write
buf lint && buf build
buf generate

# Or use the build script directly:
./build.sh              # format + lint + build + generate
```

## Architecture

Protobuf schema repository with code generation.

```
proto/vultisig/         → .proto definitions (SOURCE OF TRUTH)
Sources/swift/vultisig/ → Generated Swift code (do not edit)
go/vultisig/            → Generated Go code (do not edit)
tokens/                 → Token metadata
buf.yaml                → Buf configuration
buf.gen.yaml            → Code generation config
buf.lock                → Buf dependency lock
Package.swift           → Swift package manifest
Package.resolved        → Swift package resolution
go.mod                  → Go module definition
Makefile                → Docker-based buf workflow
build.sh                → buf format + lint + build + generate
```

Update workflow:
1. Edit `.proto` files in `proto/vultisig/`
2. Run `buf format --write && buf lint && buf build`
3. Run `buf generate`
4. Commit generated code alongside proto changes
5. Update submodule refs in iOS, Android repos

## Code Conventions

- Proto3 syntax
- Package: `vultisig.v1`
- Field naming: snake_case
- Enums: UPPER_SNAKE_CASE with 0 = UNSPECIFIED
- Every field change needs backwards compatibility consideration

## Dependencies

- **buf** — Protobuf tooling (required for all operations)
- Used as git submodule in: vultisig-ios, vultisig-android
