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
// source: ipalloc.proto

package ipalloc

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

// CustomPodInterface represents pod IP allocation made for a custom pod interface.
type CustomPodInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                               // interface name
	Network         string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`                                         // network where the interface resides
	IpAddress       string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`                    // allocated IP address
	ServiceEndpoint bool   `protobuf:"varint,4,opt,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"` // true if this interface/IP should be the k8s service endpoint for this pod
}

func (x *CustomPodInterface) Reset() {
	*x = CustomPodInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipalloc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPodInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPodInterface) ProtoMessage() {}

func (x *CustomPodInterface) ProtoReflect() protoreflect.Message {
	mi := &file_ipalloc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPodInterface.ProtoReflect.Descriptor instead.
func (*CustomPodInterface) Descriptor() ([]byte, []int) {
	return file_ipalloc_proto_rawDescGZIP(), []int{0}
}

func (x *CustomPodInterface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomPodInterface) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *CustomPodInterface) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *CustomPodInterface) GetServiceEndpoint() bool {
	if x != nil {
		return x.ServiceEndpoint
	}
	return false
}

// CustomIPAllocation represents pod IP allocation made for custom purposes (other than the main pod IP address).
// It is used for persisting IP allocation of pods with multiple interfaces and multiple IP addresses.
type CustomIPAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName          string                `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	PodNamespace     string                `protobuf:"bytes,2,opt,name=pod_namespace,json=podNamespace,proto3" json:"pod_namespace,omitempty"`
	CustomInterfaces []*CustomPodInterface `protobuf:"bytes,3,rep,name=custom_interfaces,json=customInterfaces,proto3" json:"custom_interfaces,omitempty"`
}

func (x *CustomIPAllocation) Reset() {
	*x = CustomIPAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipalloc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomIPAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomIPAllocation) ProtoMessage() {}

func (x *CustomIPAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_ipalloc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomIPAllocation.ProtoReflect.Descriptor instead.
func (*CustomIPAllocation) Descriptor() ([]byte, []int) {
	return file_ipalloc_proto_rawDescGZIP(), []int{1}
}

func (x *CustomIPAllocation) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *CustomIPAllocation) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

func (x *CustomIPAllocation) GetCustomInterfaces() []*CustomPodInterface {
	if x != nil {
		return x.CustomInterfaces
	}
	return nil
}

var File_ipalloc_proto protoreflect.FileDescriptor

var file_ipalloc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x70, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x69, 0x70, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x49, 0x50, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x70, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x76, 0x2f, 0x76, 0x70,
	0x70, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2f, 0x69,
	0x70, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ipalloc_proto_rawDescOnce sync.Once
	file_ipalloc_proto_rawDescData = file_ipalloc_proto_rawDesc
)

func file_ipalloc_proto_rawDescGZIP() []byte {
	file_ipalloc_proto_rawDescOnce.Do(func() {
		file_ipalloc_proto_rawDescData = protoimpl.X.CompressGZIP(file_ipalloc_proto_rawDescData)
	})
	return file_ipalloc_proto_rawDescData
}

var file_ipalloc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ipalloc_proto_goTypes = []interface{}{
	(*CustomPodInterface)(nil), // 0: ipalloc.CustomPodInterface
	(*CustomIPAllocation)(nil), // 1: ipalloc.CustomIPAllocation
}
var file_ipalloc_proto_depIdxs = []int32{
	0, // 0: ipalloc.CustomIPAllocation.custom_interfaces:type_name -> ipalloc.CustomPodInterface
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ipalloc_proto_init() }
func file_ipalloc_proto_init() {
	if File_ipalloc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ipalloc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPodInterface); i {
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
		file_ipalloc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomIPAllocation); i {
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
			RawDescriptor: file_ipalloc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ipalloc_proto_goTypes,
		DependencyIndexes: file_ipalloc_proto_depIdxs,
		MessageInfos:      file_ipalloc_proto_msgTypes,
	}.Build()
	File_ipalloc_proto = out.File
	file_ipalloc_proto_rawDesc = nil
	file_ipalloc_proto_goTypes = nil
	file_ipalloc_proto_depIdxs = nil
}
