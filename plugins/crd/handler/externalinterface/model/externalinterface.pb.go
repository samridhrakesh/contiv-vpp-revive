// Copyright (c) 2019 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: externalinterface.proto

package model

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

// type of the external interface
type ExternalInterface_Type int32

const (
	ExternalInterface_L2 ExternalInterface_Type = 0
	ExternalInterface_L3 ExternalInterface_Type = 1
)

// Enum value maps for ExternalInterface_Type.
var (
	ExternalInterface_Type_name = map[int32]string{
		0: "L2",
		1: "L3",
	}
	ExternalInterface_Type_value = map[string]int32{
		"L2": 0,
		"L3": 1,
	}
)

func (x ExternalInterface_Type) Enum() *ExternalInterface_Type {
	p := new(ExternalInterface_Type)
	*p = x
	return p
}

func (x ExternalInterface_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalInterface_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_externalinterface_proto_enumTypes[0].Descriptor()
}

func (ExternalInterface_Type) Type() protoreflect.EnumType {
	return &file_externalinterface_proto_enumTypes[0]
}

func (x ExternalInterface_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalInterface_Type.Descriptor instead.
func (ExternalInterface_Type) EnumDescriptor() ([]byte, []int) {
	return file_externalinterface_proto_rawDescGZIP(), []int{0, 0}
}

// ExternalInterface is used to store definition of an external interface defined via CRD.
// It is a logical entity that may mean different physical interfaces on different nodes.
type ExternalInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// logical name of the external interface
	Name string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type ExternalInterface_Type `protobuf:"varint,2,opt,name=type,proto3,enum=model.ExternalInterface_Type" json:"type,omitempty"`
	// Custom network to which this interface belongs.
	// "" or "default" means no specific custom network.
	Network string                             `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Nodes   []*ExternalInterface_NodeInterface `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ExternalInterface) Reset() {
	*x = ExternalInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externalinterface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalInterface) ProtoMessage() {}

func (x *ExternalInterface) ProtoReflect() protoreflect.Message {
	mi := &file_externalinterface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalInterface.ProtoReflect.Descriptor instead.
func (*ExternalInterface) Descriptor() ([]byte, []int) {
	return file_externalinterface_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalInterface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalInterface) GetType() ExternalInterface_Type {
	if x != nil {
		return x.Type
	}
	return ExternalInterface_L2
}

func (x *ExternalInterface) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ExternalInterface) GetNodes() []*ExternalInterface_NodeInterface {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// list of physical interfaces on individual nodes
type ExternalInterface_NodeInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node             string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	VppInterfaceName string `protobuf:"bytes,2,opt,name=vpp_interface_name,json=vppInterfaceName,proto3" json:"vpp_interface_name,omitempty"`
	Ip               string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Vlan             uint32 `protobuf:"varint,4,opt,name=vlan,proto3" json:"vlan,omitempty"`
}

func (x *ExternalInterface_NodeInterface) Reset() {
	*x = ExternalInterface_NodeInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externalinterface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalInterface_NodeInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalInterface_NodeInterface) ProtoMessage() {}

func (x *ExternalInterface_NodeInterface) ProtoReflect() protoreflect.Message {
	mi := &file_externalinterface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalInterface_NodeInterface.ProtoReflect.Descriptor instead.
func (*ExternalInterface_NodeInterface) Descriptor() ([]byte, []int) {
	return file_externalinterface_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExternalInterface_NodeInterface) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *ExternalInterface_NodeInterface) GetVppInterfaceName() string {
	if x != nil {
		return x.VppInterfaceName
	}
	return ""
}

func (x *ExternalInterface_NodeInterface) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ExternalInterface_NodeInterface) GetVlan() uint32 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

var File_externalinterface_proto protoreflect.FileDescriptor

var file_externalinterface_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0xc1, 0x02, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x3c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x75, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6c, 0x61, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x22, 0x16, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x32, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4c, 0x33, 0x10, 0x01, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_externalinterface_proto_rawDescOnce sync.Once
	file_externalinterface_proto_rawDescData = file_externalinterface_proto_rawDesc
)

func file_externalinterface_proto_rawDescGZIP() []byte {
	file_externalinterface_proto_rawDescOnce.Do(func() {
		file_externalinterface_proto_rawDescData = protoimpl.X.CompressGZIP(file_externalinterface_proto_rawDescData)
	})
	return file_externalinterface_proto_rawDescData
}

var file_externalinterface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_externalinterface_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_externalinterface_proto_goTypes = []interface{}{
	(ExternalInterface_Type)(0),             // 0: model.ExternalInterface.Type
	(*ExternalInterface)(nil),               // 1: model.ExternalInterface
	(*ExternalInterface_NodeInterface)(nil), // 2: model.ExternalInterface.NodeInterface
}
var file_externalinterface_proto_depIdxs = []int32{
	0, // 0: model.ExternalInterface.type:type_name -> model.ExternalInterface.Type
	2, // 1: model.ExternalInterface.nodes:type_name -> model.ExternalInterface.NodeInterface
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_externalinterface_proto_init() }
func file_externalinterface_proto_init() {
	if File_externalinterface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_externalinterface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalInterface); i {
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
		file_externalinterface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalInterface_NodeInterface); i {
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
			RawDescriptor: file_externalinterface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_externalinterface_proto_goTypes,
		DependencyIndexes: file_externalinterface_proto_depIdxs,
		EnumInfos:         file_externalinterface_proto_enumTypes,
		MessageInfos:      file_externalinterface_proto_msgTypes,
	}.Build()
	File_externalinterface_proto = out.File
	file_externalinterface_proto_rawDesc = nil
	file_externalinterface_proto_goTypes = nil
	file_externalinterface_proto_depIdxs = nil
}
