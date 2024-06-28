// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: vultisig/keysign/v1/blockchain_specific.proto
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

public struct VSUTXOSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var byteFee: String = String()

  public var sendMaxAmount: Bool = false

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSEthereumSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var maxFeePerGasWei: String = String()

  public var priorityFee: String = String()

  public var nonce: Int64 = 0

  public var gasLimit: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSTHORChainSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var accountNumber: UInt64 = 0

  public var sequence: UInt64 = 0

  public var fee: UInt64 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSMAYAChainSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var accountNumber: UInt64 = 0

  public var sequence: UInt64 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSCosmosSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var accountNumber: UInt64 = 0

  public var sequence: UInt64 = 0

  public var gas: UInt64 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSSolanaSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var recentBlockHash: String = String()

  public var priorityFee: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSPolkadotSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var recentBlockHash: String = String()

  public var nonce: UInt64 = 0

  public var currentBlockNumber: String = String()

  public var specVersion: UInt32 = 0

  public var transactionVersion: UInt32 = 0

  public var genesisHash: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSCoinKeyValuePair {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var key: String = String()

  public var value: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct VSSuiSpecific {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var referenceGasPrice: String = String()

  public var coinKeyValuePairs: [VSCoinKeyValuePair] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension VSUTXOSpecific: @unchecked Sendable {}
extension VSEthereumSpecific: @unchecked Sendable {}
extension VSTHORChainSpecific: @unchecked Sendable {}
extension VSMAYAChainSpecific: @unchecked Sendable {}
extension VSCosmosSpecific: @unchecked Sendable {}
extension VSSolanaSpecific: @unchecked Sendable {}
extension VSPolkadotSpecific: @unchecked Sendable {}
extension VSCoinKeyValuePair: @unchecked Sendable {}
extension VSSuiSpecific: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "vultisig.keysign.v1"

extension VSUTXOSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".UTXOSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "byte_fee"),
    2: .standard(proto: "send_max_amount"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.byteFee) }()
      case 2: try { try decoder.decodeSingularBoolField(value: &self.sendMaxAmount) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.byteFee.isEmpty {
      try visitor.visitSingularStringField(value: self.byteFee, fieldNumber: 1)
    }
    if self.sendMaxAmount != false {
      try visitor.visitSingularBoolField(value: self.sendMaxAmount, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSUTXOSpecific, rhs: VSUTXOSpecific) -> Bool {
    if lhs.byteFee != rhs.byteFee {return false}
    if lhs.sendMaxAmount != rhs.sendMaxAmount {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSEthereumSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".EthereumSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "max_fee_per_gas_wei"),
    2: .standard(proto: "priority_fee"),
    3: .same(proto: "nonce"),
    4: .standard(proto: "gas_limit"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.maxFeePerGasWei) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.priorityFee) }()
      case 3: try { try decoder.decodeSingularInt64Field(value: &self.nonce) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.gasLimit) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.maxFeePerGasWei.isEmpty {
      try visitor.visitSingularStringField(value: self.maxFeePerGasWei, fieldNumber: 1)
    }
    if !self.priorityFee.isEmpty {
      try visitor.visitSingularStringField(value: self.priorityFee, fieldNumber: 2)
    }
    if self.nonce != 0 {
      try visitor.visitSingularInt64Field(value: self.nonce, fieldNumber: 3)
    }
    if !self.gasLimit.isEmpty {
      try visitor.visitSingularStringField(value: self.gasLimit, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSEthereumSpecific, rhs: VSEthereumSpecific) -> Bool {
    if lhs.maxFeePerGasWei != rhs.maxFeePerGasWei {return false}
    if lhs.priorityFee != rhs.priorityFee {return false}
    if lhs.nonce != rhs.nonce {return false}
    if lhs.gasLimit != rhs.gasLimit {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSTHORChainSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".THORChainSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_number"),
    2: .same(proto: "sequence"),
    3: .same(proto: "fee"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularUInt64Field(value: &self.accountNumber) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.sequence) }()
      case 3: try { try decoder.decodeSingularUInt64Field(value: &self.fee) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.accountNumber != 0 {
      try visitor.visitSingularUInt64Field(value: self.accountNumber, fieldNumber: 1)
    }
    if self.sequence != 0 {
      try visitor.visitSingularUInt64Field(value: self.sequence, fieldNumber: 2)
    }
    if self.fee != 0 {
      try visitor.visitSingularUInt64Field(value: self.fee, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSTHORChainSpecific, rhs: VSTHORChainSpecific) -> Bool {
    if lhs.accountNumber != rhs.accountNumber {return false}
    if lhs.sequence != rhs.sequence {return false}
    if lhs.fee != rhs.fee {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSMAYAChainSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".MAYAChainSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_number"),
    2: .same(proto: "sequence"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularUInt64Field(value: &self.accountNumber) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.sequence) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.accountNumber != 0 {
      try visitor.visitSingularUInt64Field(value: self.accountNumber, fieldNumber: 1)
    }
    if self.sequence != 0 {
      try visitor.visitSingularUInt64Field(value: self.sequence, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSMAYAChainSpecific, rhs: VSMAYAChainSpecific) -> Bool {
    if lhs.accountNumber != rhs.accountNumber {return false}
    if lhs.sequence != rhs.sequence {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSCosmosSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CosmosSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_number"),
    2: .same(proto: "sequence"),
    3: .same(proto: "gas"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularUInt64Field(value: &self.accountNumber) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.sequence) }()
      case 3: try { try decoder.decodeSingularUInt64Field(value: &self.gas) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.accountNumber != 0 {
      try visitor.visitSingularUInt64Field(value: self.accountNumber, fieldNumber: 1)
    }
    if self.sequence != 0 {
      try visitor.visitSingularUInt64Field(value: self.sequence, fieldNumber: 2)
    }
    if self.gas != 0 {
      try visitor.visitSingularUInt64Field(value: self.gas, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSCosmosSpecific, rhs: VSCosmosSpecific) -> Bool {
    if lhs.accountNumber != rhs.accountNumber {return false}
    if lhs.sequence != rhs.sequence {return false}
    if lhs.gas != rhs.gas {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSSolanaSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".SolanaSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "recent_block_hash"),
    2: .standard(proto: "priority_fee"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.recentBlockHash) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.priorityFee) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.recentBlockHash.isEmpty {
      try visitor.visitSingularStringField(value: self.recentBlockHash, fieldNumber: 1)
    }
    if !self.priorityFee.isEmpty {
      try visitor.visitSingularStringField(value: self.priorityFee, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSSolanaSpecific, rhs: VSSolanaSpecific) -> Bool {
    if lhs.recentBlockHash != rhs.recentBlockHash {return false}
    if lhs.priorityFee != rhs.priorityFee {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSPolkadotSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".PolkadotSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "recent_block_hash"),
    2: .same(proto: "nonce"),
    3: .standard(proto: "current_block_number"),
    4: .standard(proto: "spec_version"),
    5: .standard(proto: "transaction_version"),
    6: .standard(proto: "genesis_hash"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.recentBlockHash) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.nonce) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.currentBlockNumber) }()
      case 4: try { try decoder.decodeSingularUInt32Field(value: &self.specVersion) }()
      case 5: try { try decoder.decodeSingularUInt32Field(value: &self.transactionVersion) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.genesisHash) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.recentBlockHash.isEmpty {
      try visitor.visitSingularStringField(value: self.recentBlockHash, fieldNumber: 1)
    }
    if self.nonce != 0 {
      try visitor.visitSingularUInt64Field(value: self.nonce, fieldNumber: 2)
    }
    if !self.currentBlockNumber.isEmpty {
      try visitor.visitSingularStringField(value: self.currentBlockNumber, fieldNumber: 3)
    }
    if self.specVersion != 0 {
      try visitor.visitSingularUInt32Field(value: self.specVersion, fieldNumber: 4)
    }
    if self.transactionVersion != 0 {
      try visitor.visitSingularUInt32Field(value: self.transactionVersion, fieldNumber: 5)
    }
    if !self.genesisHash.isEmpty {
      try visitor.visitSingularStringField(value: self.genesisHash, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSPolkadotSpecific, rhs: VSPolkadotSpecific) -> Bool {
    if lhs.recentBlockHash != rhs.recentBlockHash {return false}
    if lhs.nonce != rhs.nonce {return false}
    if lhs.currentBlockNumber != rhs.currentBlockNumber {return false}
    if lhs.specVersion != rhs.specVersion {return false}
    if lhs.transactionVersion != rhs.transactionVersion {return false}
    if lhs.genesisHash != rhs.genesisHash {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSCoinKeyValuePair: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CoinKeyValuePair"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "key"),
    2: .same(proto: "value"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.key) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.value) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.key.isEmpty {
      try visitor.visitSingularStringField(value: self.key, fieldNumber: 1)
    }
    if !self.value.isEmpty {
      try visitor.visitSingularStringField(value: self.value, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSCoinKeyValuePair, rhs: VSCoinKeyValuePair) -> Bool {
    if lhs.key != rhs.key {return false}
    if lhs.value != rhs.value {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension VSSuiSpecific: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".SuiSpecific"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "reference_gas_price"),
    2: .standard(proto: "coin_key_value_pairs"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.referenceGasPrice) }()
      case 2: try { try decoder.decodeRepeatedMessageField(value: &self.coinKeyValuePairs) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.referenceGasPrice.isEmpty {
      try visitor.visitSingularStringField(value: self.referenceGasPrice, fieldNumber: 1)
    }
    if !self.coinKeyValuePairs.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.coinKeyValuePairs, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: VSSuiSpecific, rhs: VSSuiSpecific) -> Bool {
    if lhs.referenceGasPrice != rhs.referenceGasPrice {return false}
    if lhs.coinKeyValuePairs != rhs.coinKeyValuePairs {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}