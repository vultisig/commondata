// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: vultisig/vault/v1/vault.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public struct VSVault {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var name: String = String()

  public var publicKeyEcdsa: String = String()

  public var publicKeyEddsa: String = String()

  public var signers: [String] = []

  public var createdAt: SwiftProtobuf.Google_Protobuf_Timestamp {
    get {return _createdAt ?? SwiftProtobuf.Google_Protobuf_Timestamp()}
    set {_createdAt = newValue}
  }
  /// Returns true if `createdAt` has been explicitly set.
  public var hasCreatedAt: Bool {return self._createdAt != nil}
  /// Clears the value of `createdAt`. Subsequent reads from it will return its default value.
  public mutating func clearCreatedAt() {self._createdAt = nil}

  public var hexChainCode: String = String()

  public var keyShares: [VSVault.KeyShare] = []

  public var localPartyID: String = String()

  public var resharePrefix: String = String()

  public var libType: VSLibType = .gg20

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public struct KeyShare {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var publicKey: String = String()

    public var keyshare: String = String()

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
  }

  public init() {}

  fileprivate var _createdAt: SwiftProtobuf.Google_Protobuf_Timestamp? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension VSVault: @unchecked Sendable {}
extension VSVault.KeyShare: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "vultisig.vault.v1"

extension VSVault: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Vault"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .standard(proto: "public_key_ecdsa"),
    3: .standard(proto: "public_key_eddsa"),
    4: .same(proto: "signers"),
    5: .standard(proto: "created_at"),
    6: .standard(proto: "hex_chain_code"),
    7: .standard(proto: "key_shares"),
    8: .standard(proto: "local_party_id"),
    9: .standard(proto: "reshare_prefix"),
    10: .standard(proto: "lib_type"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.publicKeyEcdsa) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.publicKeyEddsa) }()
      case 4: try { try decoder.decodeRepeatedStringField(value: &self.signers) }()
      case 5: try { try decoder.decodeSingularMessageField(value: &self._createdAt) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.hexChainCode) }()
      case 7: try { try decoder.decodeRepeatedMessageField(value: &self.keyShares) }()
      case 8: try { try decoder.decodeSingularStringField(value: &self.localPartyID) }()
      case 9: try { try decoder.decodeSingularStringField(value: &self.resharePrefix) }()
      case 10: try { try decoder.decodeSingularEnumField(value: &self.libType) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.publicKeyEcdsa.isEmpty {
      try visitor.visitSingularStringField(value: self.publicKeyEcdsa, fieldNumber: 2)
    }
    if !self.publicKeyEddsa.isEmpty {
      try visitor.visitSingularStringField(value: self.publicKeyEddsa, fieldNumber: 3)
    }
    if !self.signers.isEmpty {
      try visitor.visitRepeatedStringField(value: self.signers, fieldNumber: 4)
    }
    try { if let v = self._createdAt {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 5)
    } }()
    if !self.hexChainCode.isEmpty {
      try visitor.visitSingularStringField(value: self.hexChainCode, fieldNumber: 6)
    }
    if !self.keyShares.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.keyShares, fieldNumber: 7)
    }
    if !self.localPartyID.isEmpty {
      try visitor.visitSingularStringField(value: self.localPartyID, fieldNumber: 8)
    }
    if !self.resharePrefix.isEmpty {
      try visitor.visitSingularStringField(value: self.resharePrefix, fieldNumber: 9)
    }
    if self.libType != .gg20 {
      try visitor.visitSingularEnumField(value: self.libType, fieldNumber: 10)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSVault, rhs: VSVault) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.publicKeyEcdsa != rhs.publicKeyEcdsa {return false}
    if lhs.publicKeyEddsa != rhs.publicKeyEddsa {return false}
    if lhs.signers != rhs.signers {return false}
    if lhs._createdAt != rhs._createdAt {return false}
    if lhs.hexChainCode != rhs.hexChainCode {return false}
    if lhs.keyShares != rhs.keyShares {return false}
    if lhs.localPartyID != rhs.localPartyID {return false}
    if lhs.resharePrefix != rhs.resharePrefix {return false}
    if lhs.libType != rhs.libType {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSVault.KeyShare: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = VSVault.protoMessageName + ".KeyShare"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "public_key"),
    2: .same(proto: "keyshare"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.publicKey) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.keyshare) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.publicKey.isEmpty {
      try visitor.visitSingularStringField(value: self.publicKey, fieldNumber: 1)
    }
    if !self.keyshare.isEmpty {
      try visitor.visitSingularStringField(value: self.keyshare, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSVault.KeyShare, rhs: VSVault.KeyShare) -> Bool {
    if lhs.publicKey != rhs.publicKey {return false}
    if lhs.keyshare != rhs.keyshare {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
