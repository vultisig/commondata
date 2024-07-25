// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vultisig/keysign/v1/blockchain_specific.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UTXOSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByteFee       string `protobuf:"bytes,1,opt,name=byte_fee,json=byteFee,proto3" json:"byte_fee,omitempty"`
	SendMaxAmount bool   `protobuf:"varint,2,opt,name=send_max_amount,json=sendMaxAmount,proto3" json:"send_max_amount,omitempty"`
}

func (x *UTXOSpecific) Reset() {
	*x = UTXOSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UTXOSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UTXOSpecific) ProtoMessage() {}

func (x *UTXOSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UTXOSpecific.ProtoReflect.Descriptor instead.
func (*UTXOSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{0}
}

func (x *UTXOSpecific) GetByteFee() string {
	if x != nil {
		return x.ByteFee
	}
	return ""
}

func (x *UTXOSpecific) GetSendMaxAmount() bool {
	if x != nil {
		return x.SendMaxAmount
	}
	return false
}

type EthereumSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxFeePerGasWei string `protobuf:"bytes,1,opt,name=max_fee_per_gas_wei,json=maxFeePerGasWei,proto3" json:"max_fee_per_gas_wei,omitempty"`
	PriorityFee     string `protobuf:"bytes,2,opt,name=priority_fee,json=priorityFee,proto3" json:"priority_fee,omitempty"`
	Nonce           int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasLimit        string `protobuf:"bytes,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
}

func (x *EthereumSpecific) Reset() {
	*x = EthereumSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumSpecific) ProtoMessage() {}

func (x *EthereumSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumSpecific.ProtoReflect.Descriptor instead.
func (*EthereumSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{1}
}

func (x *EthereumSpecific) GetMaxFeePerGasWei() string {
	if x != nil {
		return x.MaxFeePerGasWei
	}
	return ""
}

func (x *EthereumSpecific) GetPriorityFee() string {
	if x != nil {
		return x.PriorityFee
	}
	return ""
}

func (x *EthereumSpecific) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *EthereumSpecific) GetGasLimit() string {
	if x != nil {
		return x.GasLimit
	}
	return ""
}

type THORChainSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber uint64 `protobuf:"varint,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Fee           uint64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	IsDeposit     bool   `protobuf:"varint,4,opt,name=is_deposit,json=isDeposit,proto3" json:"is_deposit,omitempty"`
}

func (x *THORChainSpecific) Reset() {
	*x = THORChainSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *THORChainSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*THORChainSpecific) ProtoMessage() {}

func (x *THORChainSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use THORChainSpecific.ProtoReflect.Descriptor instead.
func (*THORChainSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{2}
}

func (x *THORChainSpecific) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *THORChainSpecific) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *THORChainSpecific) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *THORChainSpecific) GetIsDeposit() bool {
	if x != nil {
		return x.IsDeposit
	}
	return false
}

type MAYAChainSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber uint64 `protobuf:"varint,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	IsDeposit     bool   `protobuf:"varint,3,opt,name=is_deposit,json=isDeposit,proto3" json:"is_deposit,omitempty"`
}

func (x *MAYAChainSpecific) Reset() {
	*x = MAYAChainSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAYAChainSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAYAChainSpecific) ProtoMessage() {}

func (x *MAYAChainSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAYAChainSpecific.ProtoReflect.Descriptor instead.
func (*MAYAChainSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{3}
}

func (x *MAYAChainSpecific) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *MAYAChainSpecific) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *MAYAChainSpecific) GetIsDeposit() bool {
	if x != nil {
		return x.IsDeposit
	}
	return false
}

type CosmosSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber uint64 `protobuf:"varint,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Gas           uint64 `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	IsDeposit     bool   `protobuf:"varint,4,opt,name=is_deposit,json=isDeposit,proto3" json:"is_deposit,omitempty"`
}

func (x *CosmosSpecific) Reset() {
	*x = CosmosSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosmosSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosmosSpecific) ProtoMessage() {}

func (x *CosmosSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosmosSpecific.ProtoReflect.Descriptor instead.
func (*CosmosSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{4}
}

func (x *CosmosSpecific) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *CosmosSpecific) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *CosmosSpecific) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *CosmosSpecific) GetIsDeposit() bool {
	if x != nil {
		return x.IsDeposit
	}
	return false
}

type SolanaSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecentBlockHash            string  `protobuf:"bytes,1,opt,name=recent_block_hash,json=recentBlockHash,proto3" json:"recent_block_hash,omitempty"`
	PriorityFee                string  `protobuf:"bytes,2,opt,name=priority_fee,json=priorityFee,proto3" json:"priority_fee,omitempty"`
	FromTokenAssociatedAddress *string `protobuf:"bytes,3,opt,name=from_token_associated_address,json=fromTokenAssociatedAddress,proto3,oneof" json:"from_token_associated_address,omitempty"`
	ToTokenAssociatedAddress   *string `protobuf:"bytes,4,opt,name=to_token_associated_address,json=toTokenAssociatedAddress,proto3,oneof" json:"to_token_associated_address,omitempty"`
}

func (x *SolanaSpecific) Reset() {
	*x = SolanaSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolanaSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolanaSpecific) ProtoMessage() {}

func (x *SolanaSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolanaSpecific.ProtoReflect.Descriptor instead.
func (*SolanaSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{5}
}

func (x *SolanaSpecific) GetRecentBlockHash() string {
	if x != nil {
		return x.RecentBlockHash
	}
	return ""
}

func (x *SolanaSpecific) GetPriorityFee() string {
	if x != nil {
		return x.PriorityFee
	}
	return ""
}

func (x *SolanaSpecific) GetFromTokenAssociatedAddress() string {
	if x != nil && x.FromTokenAssociatedAddress != nil {
		return *x.FromTokenAssociatedAddress
	}
	return ""
}

func (x *SolanaSpecific) GetToTokenAssociatedAddress() string {
	if x != nil && x.ToTokenAssociatedAddress != nil {
		return *x.ToTokenAssociatedAddress
	}
	return ""
}

type PolkadotSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecentBlockHash    string `protobuf:"bytes,1,opt,name=recent_block_hash,json=recentBlockHash,proto3" json:"recent_block_hash,omitempty"`
	Nonce              uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	CurrentBlockNumber string `protobuf:"bytes,3,opt,name=current_block_number,json=currentBlockNumber,proto3" json:"current_block_number,omitempty"`
	SpecVersion        uint32 `protobuf:"varint,4,opt,name=spec_version,json=specVersion,proto3" json:"spec_version,omitempty"`
	TransactionVersion uint32 `protobuf:"varint,5,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	GenesisHash        string `protobuf:"bytes,6,opt,name=genesis_hash,json=genesisHash,proto3" json:"genesis_hash,omitempty"`
}

func (x *PolkadotSpecific) Reset() {
	*x = PolkadotSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolkadotSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolkadotSpecific) ProtoMessage() {}

func (x *PolkadotSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolkadotSpecific.ProtoReflect.Descriptor instead.
func (*PolkadotSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{6}
}

func (x *PolkadotSpecific) GetRecentBlockHash() string {
	if x != nil {
		return x.RecentBlockHash
	}
	return ""
}

func (x *PolkadotSpecific) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *PolkadotSpecific) GetCurrentBlockNumber() string {
	if x != nil {
		return x.CurrentBlockNumber
	}
	return ""
}

func (x *PolkadotSpecific) GetSpecVersion() uint32 {
	if x != nil {
		return x.SpecVersion
	}
	return 0
}

func (x *PolkadotSpecific) GetTransactionVersion() uint32 {
	if x != nil {
		return x.TransactionVersion
	}
	return 0
}

func (x *PolkadotSpecific) GetGenesisHash() string {
	if x != nil {
		return x.GenesisHash
	}
	return ""
}

type CoinKeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CoinKeyValuePair) Reset() {
	*x = CoinKeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinKeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinKeyValuePair) ProtoMessage() {}

func (x *CoinKeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinKeyValuePair.ProtoReflect.Descriptor instead.
func (*CoinKeyValuePair) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{7}
}

func (x *CoinKeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CoinKeyValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SuiSpecific struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceGasPrice string              `protobuf:"bytes,1,opt,name=reference_gas_price,json=referenceGasPrice,proto3" json:"reference_gas_price,omitempty"`
	CoinKeyValuePairs []*CoinKeyValuePair `protobuf:"bytes,2,rep,name=coin_key_value_pairs,json=coinKeyValuePairs,proto3" json:"coin_key_value_pairs,omitempty"`
}

func (x *SuiSpecific) Reset() {
	*x = SuiSpecific{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuiSpecific) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuiSpecific) ProtoMessage() {}

func (x *SuiSpecific) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuiSpecific.ProtoReflect.Descriptor instead.
func (*SuiSpecific) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP(), []int{8}
}

func (x *SuiSpecific) GetReferenceGasPrice() string {
	if x != nil {
		return x.ReferenceGasPrice
	}
	return ""
}

func (x *SuiSpecific) GetCoinKeyValuePairs() []*CoinKeyValuePair {
	if x != nil {
		return x.CoinKeyValuePairs
	}
	return nil
}

var File_vultisig_keysign_v1_blockchain_specific_proto protoreflect.FileDescriptor

var file_vultisig_keysign_v1_blockchain_specific_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x51, 0x0a, 0x0c, 0x55, 0x54, 0x58, 0x4f, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x66, 0x65, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x46, 0x65, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x13,
	0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x5f,
	0x77, 0x65, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x46, 0x65,
	0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x57, 0x65, 0x69, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x87, 0x01, 0x0a, 0x11, 0x54, 0x48, 0x4f, 0x52, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x11, 0x4d, 0x41,
	0x59, 0x41, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0xad, 0x02, 0x0a, 0x0e, 0x53, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x12, 0x46, 0x0a, 0x1d, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x1a, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x42, 0x0a, 0x1b, 0x74, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x18, 0x74, 0x6f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x74, 0x6f, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xfd, 0x01, 0x0a, 0x10, 0x50, 0x6f, 0x6c,
	0x6b, 0x61, 0x64, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x2a, 0x0a,
	0x11, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x6e,
	0x65, 0x73, 0x69, 0x73, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3a, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x69, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x47, 0x61, 0x73, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x14, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x73, 0x42, 0x2f, 0x0a, 0x13,
	0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x2e, 0x76, 0x31, 0x5a, 0x13, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x76, 0x31, 0xba, 0x02, 0x02, 0x56, 0x53, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vultisig_keysign_v1_blockchain_specific_proto_rawDescOnce sync.Once
	file_vultisig_keysign_v1_blockchain_specific_proto_rawDescData = file_vultisig_keysign_v1_blockchain_specific_proto_rawDesc
)

func file_vultisig_keysign_v1_blockchain_specific_proto_rawDescGZIP() []byte {
	file_vultisig_keysign_v1_blockchain_specific_proto_rawDescOnce.Do(func() {
		file_vultisig_keysign_v1_blockchain_specific_proto_rawDescData = protoimpl.X.CompressGZIP(file_vultisig_keysign_v1_blockchain_specific_proto_rawDescData)
	})
	return file_vultisig_keysign_v1_blockchain_specific_proto_rawDescData
}

var file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vultisig_keysign_v1_blockchain_specific_proto_goTypes = []any{
	(*UTXOSpecific)(nil),      // 0: vultisig.keysign.v1.UTXOSpecific
	(*EthereumSpecific)(nil),  // 1: vultisig.keysign.v1.EthereumSpecific
	(*THORChainSpecific)(nil), // 2: vultisig.keysign.v1.THORChainSpecific
	(*MAYAChainSpecific)(nil), // 3: vultisig.keysign.v1.MAYAChainSpecific
	(*CosmosSpecific)(nil),    // 4: vultisig.keysign.v1.CosmosSpecific
	(*SolanaSpecific)(nil),    // 5: vultisig.keysign.v1.SolanaSpecific
	(*PolkadotSpecific)(nil),  // 6: vultisig.keysign.v1.PolkadotSpecific
	(*CoinKeyValuePair)(nil),  // 7: vultisig.keysign.v1.CoinKeyValuePair
	(*SuiSpecific)(nil),       // 8: vultisig.keysign.v1.SuiSpecific
}
var file_vultisig_keysign_v1_blockchain_specific_proto_depIdxs = []int32{
	7, // 0: vultisig.keysign.v1.SuiSpecific.coin_key_value_pairs:type_name -> vultisig.keysign.v1.CoinKeyValuePair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vultisig_keysign_v1_blockchain_specific_proto_init() }
func file_vultisig_keysign_v1_blockchain_specific_proto_init() {
	if File_vultisig_keysign_v1_blockchain_specific_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UTXOSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*THORChainSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MAYAChainSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CosmosSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SolanaSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PolkadotSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CoinKeyValuePair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SuiSpecific); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vultisig_keysign_v1_blockchain_specific_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vultisig_keysign_v1_blockchain_specific_proto_goTypes,
		DependencyIndexes: file_vultisig_keysign_v1_blockchain_specific_proto_depIdxs,
		MessageInfos:      file_vultisig_keysign_v1_blockchain_specific_proto_msgTypes,
	}.Build()
	File_vultisig_keysign_v1_blockchain_specific_proto = out.File
	file_vultisig_keysign_v1_blockchain_specific_proto_rawDesc = nil
	file_vultisig_keysign_v1_blockchain_specific_proto_goTypes = nil
	file_vultisig_keysign_v1_blockchain_specific_proto_depIdxs = nil
}
