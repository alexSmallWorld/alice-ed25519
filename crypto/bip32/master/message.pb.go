// Copyright © 2021 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/getamis/alice/crypto/bip32/master/message.proto

package master

import (
	circuit "github.com/getamis/alice/crypto/circuit"
	ot "github.com/getamis/alice/crypto/ot"
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

type Type int32

const (
	Type_Initial        Type = 0
	Type_OtReceiver     Type = 1
	Type_OtSendResponse Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Initial",
		1: "OtReceiver",
		2: "OtSendResponse",
	}
	Type_value = map[string]int32{
		"Initial":        0,
		"OtReceiver":     1,
		"OtSendResponse": 2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_github_com_getamis_alice_crypto_bip32_master_message_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=master.Type" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Body:
	//	*Message_Initial
	//	*Message_OtReceiver
	//	*Message_OtSendResponse
	Body isMessage_Body `protobuf_oneof:"body"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Initial
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Message) GetInitial() *BodyInitial {
	if x, ok := x.GetBody().(*Message_Initial); ok {
		return x.Initial
	}
	return nil
}

func (x *Message) GetOtReceiver() *BodyOtReceiver {
	if x, ok := x.GetBody().(*Message_OtReceiver); ok {
		return x.OtReceiver
	}
	return nil
}

func (x *Message) GetOtSendResponse() *BodyOtSendResponse {
	if x, ok := x.GetBody().(*Message_OtSendResponse); ok {
		return x.OtSendResponse
	}
	return nil
}

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Initial struct {
	Initial *BodyInitial `protobuf:"bytes,3,opt,name=initial,proto3,oneof"`
}

type Message_OtReceiver struct {
	OtReceiver *BodyOtReceiver `protobuf:"bytes,4,opt,name=otReceiver,proto3,oneof"`
}

type Message_OtSendResponse struct {
	OtSendResponse *BodyOtSendResponse `protobuf:"bytes,5,opt,name=otSendResponse,proto3,oneof"`
}

func (*Message_Initial) isMessage_Body() {}

func (*Message_OtReceiver) isMessage_Body() {}

func (*Message_OtSendResponse) isMessage_Body() {}

type BodyInitial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtRecMsg      *ot.OtReceiverMessage         `protobuf:"bytes,1,opt,name=otRecMsg,proto3" json:"otRecMsg,omitempty"`
	GarcirMsg     *circuit.GarbleCircuitMessage `protobuf:"bytes,2,opt,name=garcirMsg,proto3" json:"garcirMsg,omitempty"`
	OtherInfoWire [][]byte                      `protobuf:"bytes,3,rep,name=otherInfoWire,proto3" json:"otherInfoWire,omitempty"`
	PubKey        []byte                        `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	N             []byte                        `protobuf:"bytes,5,opt,name=N,proto3" json:"N,omitempty"`
}

func (x *BodyInitial) Reset() {
	*x = BodyInitial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInitial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInitial) ProtoMessage() {}

func (x *BodyInitial) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInitial.ProtoReflect.Descriptor instead.
func (*BodyInitial) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP(), []int{1}
}

func (x *BodyInitial) GetOtRecMsg() *ot.OtReceiverMessage {
	if x != nil {
		return x.OtRecMsg
	}
	return nil
}

func (x *BodyInitial) GetGarcirMsg() *circuit.GarbleCircuitMessage {
	if x != nil {
		return x.GarcirMsg
	}
	return nil
}

func (x *BodyInitial) GetOtherInfoWire() [][]byte {
	if x != nil {
		return x.OtherInfoWire
	}
	return nil
}

func (x *BodyInitial) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *BodyInitial) GetN() []byte {
	if x != nil {
		return x.N
	}
	return nil
}

type BodyOtReceiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtExtReceiveMsg *ot.OtExtReceiveMessage `protobuf:"bytes,1,opt,name=otExtReceiveMsg,proto3" json:"otExtReceiveMsg,omitempty"`
}

func (x *BodyOtReceiver) Reset() {
	*x = BodyOtReceiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyOtReceiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyOtReceiver) ProtoMessage() {}

func (x *BodyOtReceiver) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyOtReceiver.ProtoReflect.Descriptor instead.
func (*BodyOtReceiver) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP(), []int{2}
}

func (x *BodyOtReceiver) GetOtExtReceiveMsg() *ot.OtExtReceiveMessage {
	if x != nil {
		return x.OtExtReceiveMsg
	}
	return nil
}

type BodyOtSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtExtSendResponseMsg *ot.OtExtSendResponseMessage `protobuf:"bytes,1,opt,name=otExtSendResponseMsg,proto3" json:"otExtSendResponseMsg,omitempty"`
}

func (x *BodyOtSendResponse) Reset() {
	*x = BodyOtSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyOtSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyOtSendResponse) ProtoMessage() {}

func (x *BodyOtSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyOtSendResponse.ProtoReflect.Descriptor instead.
func (*BodyOtSendResponse) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP(), []int{3}
}

func (x *BodyOtSendResponse) GetOtExtSendResponseMsg() *ot.OtExtSendResponseMessage {
	if x != nil {
		return x.OtExtSendResponseMsg
	}
	return nil
}

var File_github_com_getamis_alice_crypto_bip32_master_message_proto protoreflect.FileDescriptor

var file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x0a,
	0x6f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0e, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x74,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x08, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x4d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x74, 0x2e, 0x4f, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6f,
	0x74, 0x52, 0x65, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x09, 0x67, 0x61, 0x72, 0x63, 0x69,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x2e, 0x67, 0x61, 0x72, 0x62, 0x6c, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x67, 0x61, 0x72, 0x63, 0x69,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x57, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x69, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x4e,
	0x22, 0x53, 0x0a, 0x0e, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x41, 0x0a, 0x0f, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x74,
	0x2e, 0x4f, 0x74, 0x45, 0x78, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x66, 0x0a, 0x12, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x14, 0x6f,
	0x74, 0x45, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x74, 0x2e, 0x4f,
	0x74, 0x45, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x14, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x2a, 0x37, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescOnce sync.Once
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescData = file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDesc
)

func file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescGZIP() []byte {
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescOnce.Do(func() {
		file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescData)
	})
	return file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDescData
}

var file_github_com_getamis_alice_crypto_bip32_master_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_getamis_alice_crypto_bip32_master_message_proto_goTypes = []interface{}{
	(Type)(0),                            // 0: master.Type
	(*Message)(nil),                      // 1: master.Message
	(*BodyInitial)(nil),                  // 2: master.BodyInitial
	(*BodyOtReceiver)(nil),               // 3: master.BodyOtReceiver
	(*BodyOtSendResponse)(nil),           // 4: master.BodyOtSendResponse
	(*ot.OtReceiverMessage)(nil),         // 5: ot.OtReceiverMessage
	(*circuit.GarbleCircuitMessage)(nil), // 6: circuit.garbleCircuitMessage
	(*ot.OtExtReceiveMessage)(nil),       // 7: ot.OtExtReceiveMessage
	(*ot.OtExtSendResponseMessage)(nil),  // 8: ot.OtExtSendResponseMessage
}
var file_github_com_getamis_alice_crypto_bip32_master_message_proto_depIdxs = []int32{
	0, // 0: master.Message.type:type_name -> master.Type
	2, // 1: master.Message.initial:type_name -> master.BodyInitial
	3, // 2: master.Message.otReceiver:type_name -> master.BodyOtReceiver
	4, // 3: master.Message.otSendResponse:type_name -> master.BodyOtSendResponse
	5, // 4: master.BodyInitial.otRecMsg:type_name -> ot.OtReceiverMessage
	6, // 5: master.BodyInitial.garcirMsg:type_name -> circuit.garbleCircuitMessage
	7, // 6: master.BodyOtReceiver.otExtReceiveMsg:type_name -> ot.OtExtReceiveMessage
	8, // 7: master.BodyOtSendResponse.otExtSendResponseMsg:type_name -> ot.OtExtSendResponseMessage
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_getamis_alice_crypto_bip32_master_message_proto_init() }
func file_github_com_getamis_alice_crypto_bip32_master_message_proto_init() {
	if File_github_com_getamis_alice_crypto_bip32_master_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInitial); i {
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
		file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyOtReceiver); i {
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
		file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyOtSendResponse); i {
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
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Initial)(nil),
		(*Message_OtReceiver)(nil),
		(*Message_OtSendResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_getamis_alice_crypto_bip32_master_message_proto_goTypes,
		DependencyIndexes: file_github_com_getamis_alice_crypto_bip32_master_message_proto_depIdxs,
		EnumInfos:         file_github_com_getamis_alice_crypto_bip32_master_message_proto_enumTypes,
		MessageInfos:      file_github_com_getamis_alice_crypto_bip32_master_message_proto_msgTypes,
	}.Build()
	File_github_com_getamis_alice_crypto_bip32_master_message_proto = out.File
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_rawDesc = nil
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_goTypes = nil
	file_github_com_getamis_alice_crypto_bip32_master_message_proto_depIdxs = nil
}