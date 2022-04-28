// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: allowlist_service.proto

package network

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetupFirewallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be a valid port number. Only http/s connection from this port are
	// allowed by the firewall.
	AllowedPort uint32 `protobuf:"varint,1,opt,name=allowed_port,json=allowedPort,proto3" json:"allowed_port,omitempty"`
}

func (x *SetupFirewallRequest) Reset() {
	*x = SetupFirewallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_allowlist_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupFirewallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupFirewallRequest) ProtoMessage() {}

func (x *SetupFirewallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allowlist_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupFirewallRequest.ProtoReflect.Descriptor instead.
func (*SetupFirewallRequest) Descriptor() ([]byte, []int) {
	return file_allowlist_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetupFirewallRequest) GetAllowedPort() uint32 {
	if x != nil {
		return x.AllowedPort
	}
	return 0
}

type GaiaLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Host and port of an HTTP proxy, formatted as "<host>:<port>". The new
	// instance of Chrome will point to the proxy via command line args.
	ProxyHostAndPort string `protobuf:"bytes,3,opt,name=proxy_host_and_port,json=proxyHostAndPort,proto3" json:"proxy_host_and_port,omitempty"`
}

func (x *GaiaLoginRequest) Reset() {
	*x = GaiaLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_allowlist_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaiaLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaiaLoginRequest) ProtoMessage() {}

func (x *GaiaLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allowlist_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaiaLoginRequest.ProtoReflect.Descriptor instead.
func (*GaiaLoginRequest) Descriptor() ([]byte, []int) {
	return file_allowlist_service_proto_rawDescGZIP(), []int{1}
}

func (x *GaiaLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GaiaLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GaiaLoginRequest) GetProxyHostAndPort() string {
	if x != nil {
		return x.ProxyHostAndPort
	}
	return ""
}

type CheckArcAppInstalledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *CheckArcAppInstalledRequest) Reset() {
	*x = CheckArcAppInstalledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_allowlist_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckArcAppInstalledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckArcAppInstalledRequest) ProtoMessage() {}

func (x *CheckArcAppInstalledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allowlist_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckArcAppInstalledRequest.ProtoReflect.Descriptor instead.
func (*CheckArcAppInstalledRequest) Descriptor() ([]byte, []int) {
	return file_allowlist_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckArcAppInstalledRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type CheckExtensionInstalledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionTitle string `protobuf:"bytes,1,opt,name=extension_title,json=extensionTitle,proto3" json:"extension_title,omitempty"`
}

func (x *CheckExtensionInstalledRequest) Reset() {
	*x = CheckExtensionInstalledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_allowlist_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckExtensionInstalledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckExtensionInstalledRequest) ProtoMessage() {}

func (x *CheckExtensionInstalledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allowlist_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckExtensionInstalledRequest.ProtoReflect.Descriptor instead.
func (*CheckExtensionInstalledRequest) Descriptor() ([]byte, []int) {
	return file_allowlist_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckExtensionInstalledRequest) GetExtensionTitle() string {
	if x != nil {
		return x.ExtensionTitle
	}
	return ""
}

var File_allowlist_service_proto protoreflect.FileDescriptor

var file_allowlist_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0x79, 0x0a, 0x10, 0x47, 0x61, 0x69, 0x61, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x2d, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x22,
	0x38, 0x0a, 0x1b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x63, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x1e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x32, 0xfc, 0x02, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x12, 0x27, 0x2e, 0x74, 0x61, 0x73,
	0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x09, 0x47, 0x61, 0x69, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x73,
	0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47,
	0x61, 0x69, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x72, 0x63, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65,
	0x64, 0x12, 0x2e, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x63, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x17, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x31, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72,
	0x6f, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f,
	0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_allowlist_service_proto_rawDescOnce sync.Once
	file_allowlist_service_proto_rawDescData = file_allowlist_service_proto_rawDesc
)

func file_allowlist_service_proto_rawDescGZIP() []byte {
	file_allowlist_service_proto_rawDescOnce.Do(func() {
		file_allowlist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_allowlist_service_proto_rawDescData)
	})
	return file_allowlist_service_proto_rawDescData
}

var file_allowlist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_allowlist_service_proto_goTypes = []interface{}{
	(*SetupFirewallRequest)(nil),           // 0: tast.cros.network.SetupFirewallRequest
	(*GaiaLoginRequest)(nil),               // 1: tast.cros.network.GaiaLoginRequest
	(*CheckArcAppInstalledRequest)(nil),    // 2: tast.cros.network.CheckArcAppInstalledRequest
	(*CheckExtensionInstalledRequest)(nil), // 3: tast.cros.network.CheckExtensionInstalledRequest
	(*empty.Empty)(nil),                    // 4: google.protobuf.Empty
}
var file_allowlist_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.network.AllowlistService.SetupFirewall:input_type -> tast.cros.network.SetupFirewallRequest
	1, // 1: tast.cros.network.AllowlistService.GaiaLogin:input_type -> tast.cros.network.GaiaLoginRequest
	2, // 2: tast.cros.network.AllowlistService.CheckArcAppInstalled:input_type -> tast.cros.network.CheckArcAppInstalledRequest
	3, // 3: tast.cros.network.AllowlistService.CheckExtensionInstalled:input_type -> tast.cros.network.CheckExtensionInstalledRequest
	4, // 4: tast.cros.network.AllowlistService.SetupFirewall:output_type -> google.protobuf.Empty
	4, // 5: tast.cros.network.AllowlistService.GaiaLogin:output_type -> google.protobuf.Empty
	4, // 6: tast.cros.network.AllowlistService.CheckArcAppInstalled:output_type -> google.protobuf.Empty
	4, // 7: tast.cros.network.AllowlistService.CheckExtensionInstalled:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_allowlist_service_proto_init() }
func file_allowlist_service_proto_init() {
	if File_allowlist_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_allowlist_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupFirewallRequest); i {
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
		file_allowlist_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaiaLoginRequest); i {
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
		file_allowlist_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckArcAppInstalledRequest); i {
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
		file_allowlist_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckExtensionInstalledRequest); i {
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
			RawDescriptor: file_allowlist_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_allowlist_service_proto_goTypes,
		DependencyIndexes: file_allowlist_service_proto_depIdxs,
		MessageInfos:      file_allowlist_service_proto_msgTypes,
	}.Build()
	File_allowlist_service_proto = out.File
	file_allowlist_service_proto_rawDesc = nil
	file_allowlist_service_proto_goTypes = nil
	file_allowlist_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AllowlistServiceClient is the client API for AllowlistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllowlistServiceClient interface {
	// SetupFirewall sets up a firewall using `iptables`, blocking https and https
	// connections through the default ports (80,443). Only http/s connections
	// coming from a specified port are allowed.
	SetupFirewall(ctx context.Context, in *SetupFirewallRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GaiaLogin starts a new Chrome instance behind a proxy and performs
	// Chrome OS login using the specified credentials.
	GaiaLogin(ctx context.Context, in *GaiaLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// CheckArcAppInstalled verifies that a specified ARC app is installed.
	CheckArcAppInstalled(ctx context.Context, in *CheckArcAppInstalledRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// CheckExtensionInstalled verifies that specified extension is installed.
	CheckExtensionInstalled(ctx context.Context, in *CheckExtensionInstalledRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type allowlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllowlistServiceClient(cc grpc.ClientConnInterface) AllowlistServiceClient {
	return &allowlistServiceClient{cc}
}

func (c *allowlistServiceClient) SetupFirewall(ctx context.Context, in *SetupFirewallRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.network.AllowlistService/SetupFirewall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allowlistServiceClient) GaiaLogin(ctx context.Context, in *GaiaLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.network.AllowlistService/GaiaLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allowlistServiceClient) CheckArcAppInstalled(ctx context.Context, in *CheckArcAppInstalledRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.network.AllowlistService/CheckArcAppInstalled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allowlistServiceClient) CheckExtensionInstalled(ctx context.Context, in *CheckExtensionInstalledRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.network.AllowlistService/CheckExtensionInstalled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllowlistServiceServer is the server API for AllowlistService service.
type AllowlistServiceServer interface {
	// SetupFirewall sets up a firewall using `iptables`, blocking https and https
	// connections through the default ports (80,443). Only http/s connections
	// coming from a specified port are allowed.
	SetupFirewall(context.Context, *SetupFirewallRequest) (*empty.Empty, error)
	// GaiaLogin starts a new Chrome instance behind a proxy and performs
	// Chrome OS login using the specified credentials.
	GaiaLogin(context.Context, *GaiaLoginRequest) (*empty.Empty, error)
	// CheckArcAppInstalled verifies that a specified ARC app is installed.
	CheckArcAppInstalled(context.Context, *CheckArcAppInstalledRequest) (*empty.Empty, error)
	// CheckExtensionInstalled verifies that specified extension is installed.
	CheckExtensionInstalled(context.Context, *CheckExtensionInstalledRequest) (*empty.Empty, error)
}

// UnimplementedAllowlistServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAllowlistServiceServer struct {
}

func (*UnimplementedAllowlistServiceServer) SetupFirewall(context.Context, *SetupFirewallRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupFirewall not implemented")
}
func (*UnimplementedAllowlistServiceServer) GaiaLogin(context.Context, *GaiaLoginRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GaiaLogin not implemented")
}
func (*UnimplementedAllowlistServiceServer) CheckArcAppInstalled(context.Context, *CheckArcAppInstalledRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckArcAppInstalled not implemented")
}
func (*UnimplementedAllowlistServiceServer) CheckExtensionInstalled(context.Context, *CheckExtensionInstalledRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExtensionInstalled not implemented")
}

func RegisterAllowlistServiceServer(s *grpc.Server, srv AllowlistServiceServer) {
	s.RegisterService(&_AllowlistService_serviceDesc, srv)
}

func _AllowlistService_SetupFirewall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupFirewallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllowlistServiceServer).SetupFirewall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.network.AllowlistService/SetupFirewall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllowlistServiceServer).SetupFirewall(ctx, req.(*SetupFirewallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllowlistService_GaiaLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GaiaLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllowlistServiceServer).GaiaLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.network.AllowlistService/GaiaLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllowlistServiceServer).GaiaLogin(ctx, req.(*GaiaLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllowlistService_CheckArcAppInstalled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckArcAppInstalledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllowlistServiceServer).CheckArcAppInstalled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.network.AllowlistService/CheckArcAppInstalled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllowlistServiceServer).CheckArcAppInstalled(ctx, req.(*CheckArcAppInstalledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllowlistService_CheckExtensionInstalled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExtensionInstalledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllowlistServiceServer).CheckExtensionInstalled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.network.AllowlistService/CheckExtensionInstalled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllowlistServiceServer).CheckExtensionInstalled(ctx, req.(*CheckExtensionInstalledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AllowlistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.network.AllowlistService",
	HandlerType: (*AllowlistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupFirewall",
			Handler:    _AllowlistService_SetupFirewall_Handler,
		},
		{
			MethodName: "GaiaLogin",
			Handler:    _AllowlistService_GaiaLogin_Handler,
		},
		{
			MethodName: "CheckArcAppInstalled",
			Handler:    _AllowlistService_CheckArcAppInstalled_Handler,
		},
		{
			MethodName: "CheckExtensionInstalled",
			Handler:    _AllowlistService_CheckExtensionInstalled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allowlist_service.proto",
}