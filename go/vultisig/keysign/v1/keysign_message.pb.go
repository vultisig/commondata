// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vultisig/keysign/v1/keysign_message.proto

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

type KeysignMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId        string          `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ServiceName      string          `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	EncryptionKeyHex string          `protobuf:"bytes,4,opt,name=encryption_key_hex,json=encryptionKeyHex,proto3" json:"encryption_key_hex,omitempty"`
	KeysignPayload   *KeysignPayload `protobuf:"bytes,5,opt,name=keysign_payload,json=keysignPayload,proto3" json:"keysign_payload,omitempty"`
	UseVultisigRelay bool            `protobuf:"varint,6,opt,name=use_vultisig_relay,json=useVultisigRelay,proto3" json:"use_vultisig_relay,omitempty"`
}

func (x *KeysignMessage) Reset() {
	*x = KeysignMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_keysign_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeysignMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeysignMessage) ProtoMessage() {}

func (x *KeysignMessage) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_keysign_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeysignMessage.ProtoReflect.Descriptor instead.
func (*KeysignMessage) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_keysign_message_proto_rawDescGZIP(), []int{0}
}

func (x *KeysignMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *KeysignMessage) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *KeysignMessage) GetEncryptionKeyHex() string {
	if x != nil {
		return x.EncryptionKeyHex
	}
	return ""
}

func (x *KeysignMessage) GetKeysignPayload() *KeysignPayload {
	if x != nil {
		return x.KeysignPayload
	}
	return nil
}

func (x *KeysignMessage) GetUseVultisigRelay() bool {
	if x != nil {
		return x.UseVultisigRelay
	}
	return false
}

type KeysignPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coin      *Coin  `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	ToAddress string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	ToAmount  string `protobuf:"bytes,3,opt,name=to_amount,json=toAmount,proto3" json:"to_amount,omitempty"`
	// Types that are assignable to BlockchainSpecific:
	//
	//	*KeysignPayload_UtxoSpecific
	//	*KeysignPayload_EthereumSpecific
	//	*KeysignPayload_ThorchainSpecific
	//	*KeysignPayload_MayaSpecific
	//	*KeysignPayload_CosmosSpecific
	//	*KeysignPayload_SolanaSpecific
	//	*KeysignPayload_PolkadotSpecific
	//	*KeysignPayload_SuicheSpecific
	BlockchainSpecific isKeysignPayload_BlockchainSpecific `protobuf_oneof:"blockchain_specific"`
	UtxoInfo           []*UtxoInfo                         `protobuf:"bytes,20,rep,name=utxo_info,json=utxoInfo,proto3" json:"utxo_info,omitempty"`
	Memo               *string                             `protobuf:"bytes,21,opt,name=memo,proto3,oneof" json:"memo,omitempty"`
	// Types that are assignable to SwapPayload:
	//
	//	*KeysignPayload_ThorchainSwapPayload
	//	*KeysignPayload_MayachainSwapPayload
	//	*KeysignPayload_OneinchSwapPayload
	SwapPayload         isKeysignPayload_SwapPayload `protobuf_oneof:"swap_payload"`
	Erc20ApprovePayload *Erc20ApprovePayload         `protobuf:"bytes,30,opt,name=erc20_approve_payload,json=erc20ApprovePayload,proto3,oneof" json:"erc20_approve_payload,omitempty"`
	VaultPublicKeyEcdsa string                       `protobuf:"bytes,31,opt,name=vault_public_key_ecdsa,json=vaultPublicKeyEcdsa,proto3" json:"vault_public_key_ecdsa,omitempty"`
	VaultLocalPartyId   string                       `protobuf:"bytes,32,opt,name=vault_local_party_id,json=vaultLocalPartyId,proto3" json:"vault_local_party_id,omitempty"`
}

func (x *KeysignPayload) Reset() {
	*x = KeysignPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vultisig_keysign_v1_keysign_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeysignPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeysignPayload) ProtoMessage() {}

func (x *KeysignPayload) ProtoReflect() protoreflect.Message {
	mi := &file_vultisig_keysign_v1_keysign_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeysignPayload.ProtoReflect.Descriptor instead.
func (*KeysignPayload) Descriptor() ([]byte, []int) {
	return file_vultisig_keysign_v1_keysign_message_proto_rawDescGZIP(), []int{1}
}

func (x *KeysignPayload) GetCoin() *Coin {
	if x != nil {
		return x.Coin
	}
	return nil
}

func (x *KeysignPayload) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *KeysignPayload) GetToAmount() string {
	if x != nil {
		return x.ToAmount
	}
	return ""
}

func (m *KeysignPayload) GetBlockchainSpecific() isKeysignPayload_BlockchainSpecific {
	if m != nil {
		return m.BlockchainSpecific
	}
	return nil
}

func (x *KeysignPayload) GetUtxoSpecific() *UTXOSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_UtxoSpecific); ok {
		return x.UtxoSpecific
	}
	return nil
}

func (x *KeysignPayload) GetEthereumSpecific() *EthereumSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_EthereumSpecific); ok {
		return x.EthereumSpecific
	}
	return nil
}

func (x *KeysignPayload) GetThorchainSpecific() *THORChainSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_ThorchainSpecific); ok {
		return x.ThorchainSpecific
	}
	return nil
}

func (x *KeysignPayload) GetMayaSpecific() *MAYAChainSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_MayaSpecific); ok {
		return x.MayaSpecific
	}
	return nil
}

func (x *KeysignPayload) GetCosmosSpecific() *CosmosSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_CosmosSpecific); ok {
		return x.CosmosSpecific
	}
	return nil
}

func (x *KeysignPayload) GetSolanaSpecific() *SolanaSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_SolanaSpecific); ok {
		return x.SolanaSpecific
	}
	return nil
}

func (x *KeysignPayload) GetPolkadotSpecific() *PolkadotSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_PolkadotSpecific); ok {
		return x.PolkadotSpecific
	}
	return nil
}

func (x *KeysignPayload) GetSuicheSpecific() *SuiSpecific {
	if x, ok := x.GetBlockchainSpecific().(*KeysignPayload_SuicheSpecific); ok {
		return x.SuicheSpecific
	}
	return nil
}

func (x *KeysignPayload) GetUtxoInfo() []*UtxoInfo {
	if x != nil {
		return x.UtxoInfo
	}
	return nil
}

func (x *KeysignPayload) GetMemo() string {
	if x != nil && x.Memo != nil {
		return *x.Memo
	}
	return ""
}

func (m *KeysignPayload) GetSwapPayload() isKeysignPayload_SwapPayload {
	if m != nil {
		return m.SwapPayload
	}
	return nil
}

func (x *KeysignPayload) GetThorchainSwapPayload() *THORChainSwapPayload {
	if x, ok := x.GetSwapPayload().(*KeysignPayload_ThorchainSwapPayload); ok {
		return x.ThorchainSwapPayload
	}
	return nil
}

func (x *KeysignPayload) GetMayachainSwapPayload() *THORChainSwapPayload {
	if x, ok := x.GetSwapPayload().(*KeysignPayload_MayachainSwapPayload); ok {
		return x.MayachainSwapPayload
	}
	return nil
}

func (x *KeysignPayload) GetOneinchSwapPayload() *OneInchSwapPayload {
	if x, ok := x.GetSwapPayload().(*KeysignPayload_OneinchSwapPayload); ok {
		return x.OneinchSwapPayload
	}
	return nil
}

func (x *KeysignPayload) GetErc20ApprovePayload() *Erc20ApprovePayload {
	if x != nil {
		return x.Erc20ApprovePayload
	}
	return nil
}

func (x *KeysignPayload) GetVaultPublicKeyEcdsa() string {
	if x != nil {
		return x.VaultPublicKeyEcdsa
	}
	return ""
}

func (x *KeysignPayload) GetVaultLocalPartyId() string {
	if x != nil {
		return x.VaultLocalPartyId
	}
	return ""
}

type isKeysignPayload_BlockchainSpecific interface {
	isKeysignPayload_BlockchainSpecific()
}

type KeysignPayload_UtxoSpecific struct {
	UtxoSpecific *UTXOSpecific `protobuf:"bytes,4,opt,name=utxo_specific,json=utxoSpecific,proto3,oneof"`
}

type KeysignPayload_EthereumSpecific struct {
	EthereumSpecific *EthereumSpecific `protobuf:"bytes,5,opt,name=ethereum_specific,json=ethereumSpecific,proto3,oneof"`
}

type KeysignPayload_ThorchainSpecific struct {
	ThorchainSpecific *THORChainSpecific `protobuf:"bytes,6,opt,name=thorchain_specific,json=thorchainSpecific,proto3,oneof"`
}

type KeysignPayload_MayaSpecific struct {
	MayaSpecific *MAYAChainSpecific `protobuf:"bytes,7,opt,name=maya_specific,json=mayaSpecific,proto3,oneof"`
}

type KeysignPayload_CosmosSpecific struct {
	CosmosSpecific *CosmosSpecific `protobuf:"bytes,8,opt,name=cosmos_specific,json=cosmosSpecific,proto3,oneof"`
}

type KeysignPayload_SolanaSpecific struct {
	SolanaSpecific *SolanaSpecific `protobuf:"bytes,9,opt,name=solana_specific,json=solanaSpecific,proto3,oneof"`
}

type KeysignPayload_PolkadotSpecific struct {
	PolkadotSpecific *PolkadotSpecific `protobuf:"bytes,10,opt,name=polkadot_specific,json=polkadotSpecific,proto3,oneof"`
}

type KeysignPayload_SuicheSpecific struct {
	SuicheSpecific *SuiSpecific `protobuf:"bytes,11,opt,name=suiche_specific,json=suicheSpecific,proto3,oneof"`
}

func (*KeysignPayload_UtxoSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_EthereumSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_ThorchainSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_MayaSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_CosmosSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_SolanaSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_PolkadotSpecific) isKeysignPayload_BlockchainSpecific() {}

func (*KeysignPayload_SuicheSpecific) isKeysignPayload_BlockchainSpecific() {}

type isKeysignPayload_SwapPayload interface {
	isKeysignPayload_SwapPayload()
}

type KeysignPayload_ThorchainSwapPayload struct {
	ThorchainSwapPayload *THORChainSwapPayload `protobuf:"bytes,22,opt,name=thorchain_swap_payload,json=thorchainSwapPayload,proto3,oneof"`
}

type KeysignPayload_MayachainSwapPayload struct {
	MayachainSwapPayload *THORChainSwapPayload `protobuf:"bytes,23,opt,name=mayachain_swap_payload,json=mayachainSwapPayload,proto3,oneof"`
}

type KeysignPayload_OneinchSwapPayload struct {
	OneinchSwapPayload *OneInchSwapPayload `protobuf:"bytes,24,opt,name=oneinch_swap_payload,json=oneinchSwapPayload,proto3,oneof"`
}

func (*KeysignPayload_ThorchainSwapPayload) isKeysignPayload_SwapPayload() {}

func (*KeysignPayload_MayachainSwapPayload) isKeysignPayload_SwapPayload() {}

func (*KeysignPayload_OneinchSwapPayload) isKeysignPayload_SwapPayload() {}

var File_vultisig_keysign_v1_keysign_message_proto protoreflect.FileDescriptor

var file_vultisig_keysign_v1_keysign_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x76, 0x75, 0x6c,
	0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x2c, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x31, 0x69, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x77, 0x61, 0x70,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x76,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x76,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x77,
	0x61, 0x70, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69,
	0x67, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x48, 0x65, 0x78, 0x12, 0x4c, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x5f, 0x76, 0x75,
	0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x75, 0x73, 0x65, 0x56, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x22, 0x91, 0x0b, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0d, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x75, 0x6c, 0x74,
	0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x54, 0x58, 0x4f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0c,
	0x75, 0x74, 0x78, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x54, 0x0a, 0x11,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73,
	0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00,
	0x52, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x12, 0x57, 0x0a, 0x12, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x48, 0x4f, 0x52, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x11, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x4d, 0x0a, 0x0d, 0x6d,
	0x61, 0x79, 0x61, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x41, 0x59, 0x41, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x61,
	0x79, 0x61, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x4e, 0x0a, 0x0f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x6f,
	0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x54, 0x0a, 0x11, 0x70, 0x6f,
	0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6b,
	0x61, 0x64, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x10,
	0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x75, 0x69, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x75, 0x6c, 0x74,
	0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x69, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x73,
	0x75, 0x69, 0x63, 0x68, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x3a, 0x0a,
	0x09, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73,
	0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x74, 0x78, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x75, 0x74, 0x78, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x88,
	0x01, 0x01, 0x12, 0x61, 0x0a, 0x16, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x73, 0x77, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65,
	0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x48, 0x4f, 0x52, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x01, 0x52,
	0x14, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x61, 0x0a, 0x16, 0x6d, 0x61, 0x79, 0x61, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x48, 0x4f, 0x52,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x01, 0x52, 0x14, 0x6d, 0x61, 0x79, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x77, 0x61,
	0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x5b, 0x0a, 0x14, 0x6f, 0x6e, 0x65, 0x69,
	0x6e, 0x63, 0x68, 0x5f, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69,
	0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x65,
	0x49, 0x6e, 0x63, 0x68, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48,
	0x01, 0x52, 0x12, 0x6f, 0x6e, 0x65, 0x69, 0x6e, 0x63, 0x68, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x61, 0x0a, 0x15, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e,
	0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x63, 0x32, 0x30,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x03,
	0x52, 0x13, 0x65, 0x72, 0x63, 0x32, 0x30, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x16, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x63, 0x64,
	0x73, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x45, 0x63, 0x64, 0x73, 0x61, 0x12, 0x2f, 0x0a,
	0x14, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x42, 0x15,
	0x0a, 0x13, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x42, 0x18,
	0x0a, 0x16, 0x5f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x2f, 0x0a, 0x13, 0x76, 0x75, 0x6c, 0x74,
	0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x5a,
	0x13, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x69, 0x67,
	0x6e, 0x2f, 0x76, 0x31, 0xba, 0x02, 0x02, 0x56, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_vultisig_keysign_v1_keysign_message_proto_rawDescOnce sync.Once
	file_vultisig_keysign_v1_keysign_message_proto_rawDescData = file_vultisig_keysign_v1_keysign_message_proto_rawDesc
)

func file_vultisig_keysign_v1_keysign_message_proto_rawDescGZIP() []byte {
	file_vultisig_keysign_v1_keysign_message_proto_rawDescOnce.Do(func() {
		file_vultisig_keysign_v1_keysign_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_vultisig_keysign_v1_keysign_message_proto_rawDescData)
	})
	return file_vultisig_keysign_v1_keysign_message_proto_rawDescData
}

var file_vultisig_keysign_v1_keysign_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vultisig_keysign_v1_keysign_message_proto_goTypes = []any{
	(*KeysignMessage)(nil),       // 0: vultisig.keysign.v1.KeysignMessage
	(*KeysignPayload)(nil),       // 1: vultisig.keysign.v1.KeysignPayload
	(*Coin)(nil),                 // 2: vultisig.keysign.v1.Coin
	(*UTXOSpecific)(nil),         // 3: vultisig.keysign.v1.UTXOSpecific
	(*EthereumSpecific)(nil),     // 4: vultisig.keysign.v1.EthereumSpecific
	(*THORChainSpecific)(nil),    // 5: vultisig.keysign.v1.THORChainSpecific
	(*MAYAChainSpecific)(nil),    // 6: vultisig.keysign.v1.MAYAChainSpecific
	(*CosmosSpecific)(nil),       // 7: vultisig.keysign.v1.CosmosSpecific
	(*SolanaSpecific)(nil),       // 8: vultisig.keysign.v1.SolanaSpecific
	(*PolkadotSpecific)(nil),     // 9: vultisig.keysign.v1.PolkadotSpecific
	(*SuiSpecific)(nil),          // 10: vultisig.keysign.v1.SuiSpecific
	(*UtxoInfo)(nil),             // 11: vultisig.keysign.v1.UtxoInfo
	(*THORChainSwapPayload)(nil), // 12: vultisig.keysign.v1.THORChainSwapPayload
	(*OneInchSwapPayload)(nil),   // 13: vultisig.keysign.v1.OneInchSwapPayload
	(*Erc20ApprovePayload)(nil),  // 14: vultisig.keysign.v1.Erc20ApprovePayload
}
var file_vultisig_keysign_v1_keysign_message_proto_depIdxs = []int32{
	1,  // 0: vultisig.keysign.v1.KeysignMessage.keysign_payload:type_name -> vultisig.keysign.v1.KeysignPayload
	2,  // 1: vultisig.keysign.v1.KeysignPayload.coin:type_name -> vultisig.keysign.v1.Coin
	3,  // 2: vultisig.keysign.v1.KeysignPayload.utxo_specific:type_name -> vultisig.keysign.v1.UTXOSpecific
	4,  // 3: vultisig.keysign.v1.KeysignPayload.ethereum_specific:type_name -> vultisig.keysign.v1.EthereumSpecific
	5,  // 4: vultisig.keysign.v1.KeysignPayload.thorchain_specific:type_name -> vultisig.keysign.v1.THORChainSpecific
	6,  // 5: vultisig.keysign.v1.KeysignPayload.maya_specific:type_name -> vultisig.keysign.v1.MAYAChainSpecific
	7,  // 6: vultisig.keysign.v1.KeysignPayload.cosmos_specific:type_name -> vultisig.keysign.v1.CosmosSpecific
	8,  // 7: vultisig.keysign.v1.KeysignPayload.solana_specific:type_name -> vultisig.keysign.v1.SolanaSpecific
	9,  // 8: vultisig.keysign.v1.KeysignPayload.polkadot_specific:type_name -> vultisig.keysign.v1.PolkadotSpecific
	10, // 9: vultisig.keysign.v1.KeysignPayload.suiche_specific:type_name -> vultisig.keysign.v1.SuiSpecific
	11, // 10: vultisig.keysign.v1.KeysignPayload.utxo_info:type_name -> vultisig.keysign.v1.UtxoInfo
	12, // 11: vultisig.keysign.v1.KeysignPayload.thorchain_swap_payload:type_name -> vultisig.keysign.v1.THORChainSwapPayload
	12, // 12: vultisig.keysign.v1.KeysignPayload.mayachain_swap_payload:type_name -> vultisig.keysign.v1.THORChainSwapPayload
	13, // 13: vultisig.keysign.v1.KeysignPayload.oneinch_swap_payload:type_name -> vultisig.keysign.v1.OneInchSwapPayload
	14, // 14: vultisig.keysign.v1.KeysignPayload.erc20_approve_payload:type_name -> vultisig.keysign.v1.Erc20ApprovePayload
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_vultisig_keysign_v1_keysign_message_proto_init() }
func file_vultisig_keysign_v1_keysign_message_proto_init() {
	if File_vultisig_keysign_v1_keysign_message_proto != nil {
		return
	}
	file_vultisig_keysign_v1_1inch_swap_payload_proto_init()
	file_vultisig_keysign_v1_blockchain_specific_proto_init()
	file_vultisig_keysign_v1_coin_proto_init()
	file_vultisig_keysign_v1_erc20_approve_payload_proto_init()
	file_vultisig_keysign_v1_thorchain_swap_payload_proto_init()
	file_vultisig_keysign_v1_utxo_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vultisig_keysign_v1_keysign_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KeysignMessage); i {
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
		file_vultisig_keysign_v1_keysign_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeysignPayload); i {
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
	file_vultisig_keysign_v1_keysign_message_proto_msgTypes[1].OneofWrappers = []any{
		(*KeysignPayload_UtxoSpecific)(nil),
		(*KeysignPayload_EthereumSpecific)(nil),
		(*KeysignPayload_ThorchainSpecific)(nil),
		(*KeysignPayload_MayaSpecific)(nil),
		(*KeysignPayload_CosmosSpecific)(nil),
		(*KeysignPayload_SolanaSpecific)(nil),
		(*KeysignPayload_PolkadotSpecific)(nil),
		(*KeysignPayload_SuicheSpecific)(nil),
		(*KeysignPayload_ThorchainSwapPayload)(nil),
		(*KeysignPayload_MayachainSwapPayload)(nil),
		(*KeysignPayload_OneinchSwapPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vultisig_keysign_v1_keysign_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vultisig_keysign_v1_keysign_message_proto_goTypes,
		DependencyIndexes: file_vultisig_keysign_v1_keysign_message_proto_depIdxs,
		MessageInfos:      file_vultisig_keysign_v1_keysign_message_proto_msgTypes,
	}.Build()
	File_vultisig_keysign_v1_keysign_message_proto = out.File
	file_vultisig_keysign_v1_keysign_message_proto_rawDesc = nil
	file_vultisig_keysign_v1_keysign_message_proto_goTypes = nil
	file_vultisig_keysign_v1_keysign_message_proto_depIdxs = nil
}
