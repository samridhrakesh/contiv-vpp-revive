// Copyright (c) 2018 Cisco and/or its affiliates.
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
// source: stn.proto

// Package stn defines STN (Steal the NIC) GRPC service.

package stn

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

// STNRequest represents a request to steal or release of an interface.
type STNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The interface to be stolen / released.
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// True if DHCP is enabled on the interface.
	DhcpEnabled bool `protobuf:"varint,2,opt,name=dhcp_enabled,json=dhcpEnabled,proto3" json:"dhcp_enabled,omitempty"`
}

func (x *STNRequest) Reset() {
	*x = STNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STNRequest) ProtoMessage() {}

func (x *STNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STNRequest.ProtoReflect.Descriptor instead.
func (*STNRequest) Descriptor() ([]byte, []int) {
	return file_stn_proto_rawDescGZIP(), []int{0}
}

func (x *STNRequest) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *STNRequest) GetDhcpEnabled() bool {
	if x != nil {
		return x.DhcpEnabled
	}
	return false
}

// The reply to the STNRequest. Contains the original config of the stolen interface.
type STNReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result code. 0 = success, non-zero = error.
	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	// Error string in case that result != 0.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// PCI address of the interface.
	PciAddress string `protobuf:"bytes,3,opt,name=pci_address,json=pciAddress,proto3" json:"pci_address,omitempty"`
	// List of IP addresses assigned to the interface.
	IpAddresses []string `protobuf:"bytes,4,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	// List of routes related to the interface.
	Routes []*STNReply_Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	// Kernel driver used by this interface before stealing.
	KernelDriver string `protobuf:"bytes,6,opt,name=kernel_driver,json=kernelDriver,proto3" json:"kernel_driver,omitempty"`
}

func (x *STNReply) Reset() {
	*x = STNReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STNReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STNReply) ProtoMessage() {}

func (x *STNReply) ProtoReflect() protoreflect.Message {
	mi := &file_stn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STNReply.ProtoReflect.Descriptor instead.
func (*STNReply) Descriptor() ([]byte, []int) {
	return file_stn_proto_rawDescGZIP(), []int{1}
}

func (x *STNReply) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *STNReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *STNReply) GetPciAddress() string {
	if x != nil {
		return x.PciAddress
	}
	return ""
}

func (x *STNReply) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *STNReply) GetRoutes() []*STNReply_Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *STNReply) GetKernelDriver() string {
	if x != nil {
		return x.KernelDriver
	}
	return ""
}

// A route related to the interface.
type STNReply_Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Destination subnet prefix.
	DestinationSubnet string `protobuf:"bytes,1,opt,name=destination_subnet,json=destinationSubnet,proto3" json:"destination_subnet,omitempty"`
	// Next hop IP address.
	NextHopIp string `protobuf:"bytes,2,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
}

func (x *STNReply_Route) Reset() {
	*x = STNReply_Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STNReply_Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STNReply_Route) ProtoMessage() {}

func (x *STNReply_Route) ProtoReflect() protoreflect.Message {
	mi := &file_stn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STNReply_Route.ProtoReflect.Descriptor instead.
func (*STNReply_Route) Descriptor() ([]byte, []int) {
	return file_stn_proto_rawDescGZIP(), []int{1, 0}
}

func (x *STNReply_Route) GetDestinationSubnet() string {
	if x != nil {
		return x.DestinationSubnet
	}
	return ""
}

func (x *STNReply_Route) GetNextHopIp() string {
	if x != nil {
		return x.NextHopIp
	}
	return ""
}

var File_stn_proto protoreflect.FileDescriptor

var file_stn_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x74, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x74, 0x6e,
	0x22, 0x56, 0x0a, 0x0a, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x68, 0x63, 0x70, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x68, 0x63,
	0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xa6, 0x02, 0x0a, 0x08, 0x53, 0x54, 0x4e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x63, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x63, 0x69, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6e, 0x2e, 0x53, 0x54,
	0x4e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x56, 0x0a, 0x05, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x49,
	0x70, 0x32, 0xa8, 0x01, 0x0a, 0x03, 0x53, 0x54, 0x4e, 0x12, 0x32, 0x0a, 0x0e, 0x53, 0x74, 0x65,
	0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74,
	0x6e, 0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73,
	0x74, 0x6e, 0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x10, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x6e, 0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x74, 0x6e, 0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x6c, 0x65, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x6e,
	0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x74,
	0x6e, 0x2e, 0x53, 0x54, 0x4e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01,
	0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stn_proto_rawDescOnce sync.Once
	file_stn_proto_rawDescData = file_stn_proto_rawDesc
)

func file_stn_proto_rawDescGZIP() []byte {
	file_stn_proto_rawDescOnce.Do(func() {
		file_stn_proto_rawDescData = protoimpl.X.CompressGZIP(file_stn_proto_rawDescData)
	})
	return file_stn_proto_rawDescData
}

var file_stn_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stn_proto_goTypes = []interface{}{
	(*STNRequest)(nil),     // 0: stn.STNRequest
	(*STNReply)(nil),       // 1: stn.STNReply
	(*STNReply_Route)(nil), // 2: stn.STNReply.Route
}
var file_stn_proto_depIdxs = []int32{
	2, // 0: stn.STNReply.routes:type_name -> stn.STNReply.Route
	0, // 1: stn.STN.StealInterface:input_type -> stn.STNRequest
	0, // 2: stn.STN.ReleaseInterface:input_type -> stn.STNRequest
	0, // 3: stn.STN.StolenInterfaceInfo:input_type -> stn.STNRequest
	1, // 4: stn.STN.StealInterface:output_type -> stn.STNReply
	1, // 5: stn.STN.ReleaseInterface:output_type -> stn.STNReply
	1, // 6: stn.STN.StolenInterfaceInfo:output_type -> stn.STNReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stn_proto_init() }
func file_stn_proto_init() {
	if File_stn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STNRequest); i {
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
		file_stn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STNReply); i {
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
		file_stn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STNReply_Route); i {
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
			RawDescriptor: file_stn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stn_proto_goTypes,
		DependencyIndexes: file_stn_proto_depIdxs,
		MessageInfos:      file_stn_proto_msgTypes,
	}.Build()
	File_stn_proto = out.File
	file_stn_proto_rawDesc = nil
	file_stn_proto_goTypes = nil
	file_stn_proto_depIdxs = nil
}
