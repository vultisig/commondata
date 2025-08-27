# Vultisig CommonData - Development Guide

## Build Commands
- **Generate all protobuf files**: `make` or `docker run --rm -v "$(PWD)":/workspace -w /workspace --entrypoint /bin/sh bufbuild/buf "./build.sh"`
- **Format proto files**: `buf format -w`
- **Lint proto files**: `buf lint`
- **Build proto files**: `buf build`
- **Generate code**: `buf generate`
- **Test Go code**: `go test -v ./...`
- **Test Swift code**: `swift test`
- **Build Swift package**: `swift build`

## Code Style Guidelines
- **Proto files**: Use snake_case for field names, PascalCase for message/enum names
- **Go imports**: Use protobuf runtime imports from `google.golang.org/protobuf`
- **Swift prefix**: All generated Swift types use `VS` prefix (e.g., `VSKeysignMessage`)
- **Package structure**: Proto files in `proto/vultisig/{domain}/v1/`, generated Go in `go/`, Swift in `Sources/swift/`
- **Proto options**: Always include `go_package`, `java_package`, and `swift_prefix` options
- **Proto syntax**: Use `proto3` syntax exclusively
- **Field numbering**: Reserve fields 1-20 for core properties, 20+ for arrays/optional fields
- **Oneof fields**: Use for blockchain-specific and swap payload variations
- **Error handling**: Proto generation errors fail the build process
- **File naming**: Use snake_case for proto files (e.g., `keysign_message.proto`)
- **Import paths**: Use full package paths for proto imports (e.g., `vultisig/keysign/v1/coin.proto`)
- **Generated code**: Never manually edit generated files in `go/` or `Sources/swift/` directories