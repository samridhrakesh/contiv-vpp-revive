// Copyright (c) 2017 Cisco and/or its affiliates.
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
// source: sfc.proto

// Package pod defines data model for Kubernetes Pod.

package sfc

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

// Pod is a collection of containers that can run on a host.
// This resource is created by clients and scheduled onto hosts.
type Sfc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the pod unique within the namespace.
	// Cannot be updated.
	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	// Node the pod is deployed into into.
	// Can be updated
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *Sfc) Reset() {
	*x = Sfc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sfc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sfc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sfc) ProtoMessage() {}

func (x *Sfc) ProtoReflect() protoreflect.Message {
	mi := &file_sfc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sfc.ProtoReflect.Descriptor instead.
func (*Sfc) Descriptor() ([]byte, []int) {
	return file_sfc_proto_rawDescGZIP(), []int{0}
}

func (x *Sfc) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *Sfc) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

var File_sfc_proto protoreflect.FileDescriptor

var file_sfc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x66, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x66, 0x63,
	0x22, 0x2b, 0x0a, 0x03, 0x53, 0x66, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x73, 0x66, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sfc_proto_rawDescOnce sync.Once
	file_sfc_proto_rawDescData = file_sfc_proto_rawDesc
)

func file_sfc_proto_rawDescGZIP() []byte {
	file_sfc_proto_rawDescOnce.Do(func() {
		file_sfc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sfc_proto_rawDescData)
	})
	return file_sfc_proto_rawDescData
}

var file_sfc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sfc_proto_goTypes = []interface{}{
	(*Sfc)(nil), // 0: sfc.Sfc
}
var file_sfc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sfc_proto_init() }
func file_sfc_proto_init() {
	if File_sfc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sfc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sfc); i {
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
			RawDescriptor: file_sfc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sfc_proto_goTypes,
		DependencyIndexes: file_sfc_proto_depIdxs,
		MessageInfos:      file_sfc_proto_msgTypes,
	}.Build()
	File_sfc_proto = out.File
	file_sfc_proto_rawDesc = nil
	file_sfc_proto_goTypes = nil
	file_sfc_proto_depIdxs = nil
}
