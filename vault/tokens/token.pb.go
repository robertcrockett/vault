// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vault/tokens/token.proto

package tokens

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

// SignedToken
type SignedToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenVersion uint64 `protobuf:"varint,1,opt,name=token_version,json=tokenVersion,proto3" json:"token_version,omitempty"` // always 1 for now
	Hmac         []byte `protobuf:"bytes,2,opt,name=hmac,proto3" json:"hmac,omitempty"`                                      // HMAC of token
	Token        []byte `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                                    // protobuf-marshalled Token message
}

func (x *SignedToken) Reset() {
	*x = SignedToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_tokens_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedToken) ProtoMessage() {}

func (x *SignedToken) ProtoReflect() protoreflect.Message {
	mi := &file_vault_tokens_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedToken.ProtoReflect.Descriptor instead.
func (*SignedToken) Descriptor() ([]byte, []int) {
	return file_vault_tokens_token_proto_rawDescGZIP(), []int{0}
}

func (x *SignedToken) GetTokenVersion() uint64 {
	if x != nil {
		return x.TokenVersion
	}
	return 0
}

func (x *SignedToken) GetHmac() []byte {
	if x != nil {
		return x.Hmac
	}
	return nil
}

func (x *SignedToken) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Random     string `protobuf:"bytes,1,opt,name=random,proto3" json:"random,omitempty"`                            // unencoded equiv of former $randbase62
	LocalIndex uint64 `protobuf:"varint,2,opt,name=local_index,json=localIndex,proto3" json:"local_index,omitempty"` // required storage state to have this token
	IndexEpoch uint32 `protobuf:"varint,3,opt,name=index_epoch,json=indexEpoch,proto3" json:"index_epoch,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_tokens_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_vault_tokens_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_vault_tokens_token_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetRandom() string {
	if x != nil {
		return x.Random
	}
	return ""
}

func (x *Token) GetLocalIndex() uint64 {
	if x != nil {
		return x.LocalIndex
	}
	return 0
}

func (x *Token) GetIndexEpoch() uint32 {
	if x != nil {
		return x.IndexEpoch
	}
	return 0
}

var File_vault_tokens_token_proto protoreflect.FileDescriptor

var file_vault_tokens_token_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x5c, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6d, 0x61, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x6d, 0x61, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x61, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_tokens_token_proto_rawDescOnce sync.Once
	file_vault_tokens_token_proto_rawDescData = file_vault_tokens_token_proto_rawDesc
)

func file_vault_tokens_token_proto_rawDescGZIP() []byte {
	file_vault_tokens_token_proto_rawDescOnce.Do(func() {
		file_vault_tokens_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_tokens_token_proto_rawDescData)
	})
	return file_vault_tokens_token_proto_rawDescData
}

var file_vault_tokens_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vault_tokens_token_proto_goTypes = []any{
	(*SignedToken)(nil), // 0: tokens.SignedToken
	(*Token)(nil),       // 1: tokens.Token
}
var file_vault_tokens_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vault_tokens_token_proto_init() }
func file_vault_tokens_token_proto_init() {
	if File_vault_tokens_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vault_tokens_token_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SignedToken); i {
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
		file_vault_tokens_token_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Token); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vault_tokens_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vault_tokens_token_proto_goTypes,
		DependencyIndexes: file_vault_tokens_token_proto_depIdxs,
		MessageInfos:      file_vault_tokens_token_proto_msgTypes,
	}.Build()
	File_vault_tokens_token_proto = out.File
	file_vault_tokens_token_proto_rawDesc = nil
	file_vault_tokens_token_proto_goTypes = nil
	file_vault_tokens_token_proto_depIdxs = nil
}
