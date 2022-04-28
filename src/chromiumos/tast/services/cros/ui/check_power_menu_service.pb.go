// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: check_power_menu_service.proto

package ui

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

type NewChromeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If login is false, a session starts without logging in.
	Login bool `protobuf:"varint,1,opt,name=login,proto3" json:"login,omitempty"`
	// If non-empty, the key is used to load sign-in profile extension.
	// Namely, to show the login screen, but without logging in, login would
	// be set to false, and the key would be supplied with the
	// signinProfileTestExtensionManifestKey.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *NewChromeRequest) Reset() {
	*x = NewChromeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_power_menu_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChromeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChromeRequest) ProtoMessage() {}

func (x *NewChromeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_power_menu_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChromeRequest.ProtoReflect.Descriptor instead.
func (*NewChromeRequest) Descriptor() ([]byte, []int) {
	return file_check_power_menu_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewChromeRequest) GetLogin() bool {
	if x != nil {
		return x.Login
	}
	return false
}

func (x *NewChromeRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PowerMenuPresentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMenuPresent bool `protobuf:"varint,1,opt,name=is_menu_present,json=isMenuPresent,proto3" json:"is_menu_present,omitempty"`
}

func (x *PowerMenuPresentResponse) Reset() {
	*x = PowerMenuPresentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_power_menu_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerMenuPresentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerMenuPresentResponse) ProtoMessage() {}

func (x *PowerMenuPresentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_power_menu_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerMenuPresentResponse.ProtoReflect.Descriptor instead.
func (*PowerMenuPresentResponse) Descriptor() ([]byte, []int) {
	return file_check_power_menu_service_proto_rawDescGZIP(), []int{1}
}

func (x *PowerMenuPresentResponse) GetIsMenuPresent() bool {
	if x != nil {
		return x.IsMenuPresent
	}
	return false
}

type PowerMenuItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuItems []string `protobuf:"bytes,1,rep,name=menu_items,json=menuItems,proto3" json:"menu_items,omitempty"`
}

func (x *PowerMenuItemResponse) Reset() {
	*x = PowerMenuItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_power_menu_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerMenuItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerMenuItemResponse) ProtoMessage() {}

func (x *PowerMenuItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_power_menu_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerMenuItemResponse.ProtoReflect.Descriptor instead.
func (*PowerMenuItemResponse) Descriptor() ([]byte, []int) {
	return file_check_power_menu_service_proto_rawDescGZIP(), []int{2}
}

func (x *PowerMenuItemResponse) GetMenuItems() []string {
	if x != nil {
		return x.MenuItems
	}
	return nil
}

var File_check_power_menu_service_proto protoreflect.FileDescriptor

var file_check_power_menu_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x75, 0x69, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x10, 0x4e,
	0x65, 0x77, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x42, 0x0a, 0x18, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0xc0, 0x02, 0x0a, 0x10, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x6e,
	0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x75, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x10, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x74,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x75, 0x69, 0x2e, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d,
	0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x23, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x75, 0x69, 0x2e, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x75, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_check_power_menu_service_proto_rawDescOnce sync.Once
	file_check_power_menu_service_proto_rawDescData = file_check_power_menu_service_proto_rawDesc
)

func file_check_power_menu_service_proto_rawDescGZIP() []byte {
	file_check_power_menu_service_proto_rawDescOnce.Do(func() {
		file_check_power_menu_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_power_menu_service_proto_rawDescData)
	})
	return file_check_power_menu_service_proto_rawDescData
}

var file_check_power_menu_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_check_power_menu_service_proto_goTypes = []interface{}{
	(*NewChromeRequest)(nil),         // 0: tast.cros.ui.NewChromeRequest
	(*PowerMenuPresentResponse)(nil), // 1: tast.cros.ui.PowerMenuPresentResponse
	(*PowerMenuItemResponse)(nil),    // 2: tast.cros.ui.PowerMenuItemResponse
	(*empty.Empty)(nil),              // 3: google.protobuf.Empty
}
var file_check_power_menu_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.ui.PowerMenuService.NewChrome:input_type -> tast.cros.ui.NewChromeRequest
	3, // 1: tast.cros.ui.PowerMenuService.CloseChrome:input_type -> google.protobuf.Empty
	3, // 2: tast.cros.ui.PowerMenuService.PowerMenuPresent:input_type -> google.protobuf.Empty
	3, // 3: tast.cros.ui.PowerMenuService.PowerMenuItem:input_type -> google.protobuf.Empty
	3, // 4: tast.cros.ui.PowerMenuService.NewChrome:output_type -> google.protobuf.Empty
	3, // 5: tast.cros.ui.PowerMenuService.CloseChrome:output_type -> google.protobuf.Empty
	1, // 6: tast.cros.ui.PowerMenuService.PowerMenuPresent:output_type -> tast.cros.ui.PowerMenuPresentResponse
	2, // 7: tast.cros.ui.PowerMenuService.PowerMenuItem:output_type -> tast.cros.ui.PowerMenuItemResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_power_menu_service_proto_init() }
func file_check_power_menu_service_proto_init() {
	if File_check_power_menu_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_power_menu_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChromeRequest); i {
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
		file_check_power_menu_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerMenuPresentResponse); i {
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
		file_check_power_menu_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerMenuItemResponse); i {
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
			RawDescriptor: file_check_power_menu_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_power_menu_service_proto_goTypes,
		DependencyIndexes: file_check_power_menu_service_proto_depIdxs,
		MessageInfos:      file_check_power_menu_service_proto_msgTypes,
	}.Build()
	File_check_power_menu_service_proto = out.File
	file_check_power_menu_service_proto_rawDesc = nil
	file_check_power_menu_service_proto_goTypes = nil
	file_check_power_menu_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PowerMenuServiceClient is the client API for PowerMenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PowerMenuServiceClient interface {
	// NewChrome starts a Chrome session and processes the sign-in request.
	// CloseChrome must be called later to clean up the associated resources.
	NewChrome(ctx context.Context, in *NewChromeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Close releases the resources obtained by NewChrome.
	CloseChrome(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// PowerMenuPresent returns a bool to indicate whether the presence of a power menu
	// is true. Chrome instance is necessary prior to the deployment. For this reason,
	// NewChrome must be called in prior, but not CloseChrome.
	PowerMenuPresent(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PowerMenuPresentResponse, error)
	// PowerMenuItem returns a slice which contains names of power menu items.
	PowerMenuItem(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PowerMenuItemResponse, error)
}

type powerMenuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPowerMenuServiceClient(cc grpc.ClientConnInterface) PowerMenuServiceClient {
	return &powerMenuServiceClient{cc}
}

func (c *powerMenuServiceClient) NewChrome(ctx context.Context, in *NewChromeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.ui.PowerMenuService/NewChrome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerMenuServiceClient) CloseChrome(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.ui.PowerMenuService/CloseChrome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerMenuServiceClient) PowerMenuPresent(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PowerMenuPresentResponse, error) {
	out := new(PowerMenuPresentResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.ui.PowerMenuService/PowerMenuPresent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerMenuServiceClient) PowerMenuItem(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PowerMenuItemResponse, error) {
	out := new(PowerMenuItemResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.ui.PowerMenuService/PowerMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PowerMenuServiceServer is the server API for PowerMenuService service.
type PowerMenuServiceServer interface {
	// NewChrome starts a Chrome session and processes the sign-in request.
	// CloseChrome must be called later to clean up the associated resources.
	NewChrome(context.Context, *NewChromeRequest) (*empty.Empty, error)
	// Close releases the resources obtained by NewChrome.
	CloseChrome(context.Context, *empty.Empty) (*empty.Empty, error)
	// PowerMenuPresent returns a bool to indicate whether the presence of a power menu
	// is true. Chrome instance is necessary prior to the deployment. For this reason,
	// NewChrome must be called in prior, but not CloseChrome.
	PowerMenuPresent(context.Context, *empty.Empty) (*PowerMenuPresentResponse, error)
	// PowerMenuItem returns a slice which contains names of power menu items.
	PowerMenuItem(context.Context, *empty.Empty) (*PowerMenuItemResponse, error)
}

// UnimplementedPowerMenuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPowerMenuServiceServer struct {
}

func (*UnimplementedPowerMenuServiceServer) NewChrome(context.Context, *NewChromeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChrome not implemented")
}
func (*UnimplementedPowerMenuServiceServer) CloseChrome(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseChrome not implemented")
}
func (*UnimplementedPowerMenuServiceServer) PowerMenuPresent(context.Context, *empty.Empty) (*PowerMenuPresentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerMenuPresent not implemented")
}
func (*UnimplementedPowerMenuServiceServer) PowerMenuItem(context.Context, *empty.Empty) (*PowerMenuItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerMenuItem not implemented")
}

func RegisterPowerMenuServiceServer(s *grpc.Server, srv PowerMenuServiceServer) {
	s.RegisterService(&_PowerMenuService_serviceDesc, srv)
}

func _PowerMenuService_NewChrome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChromeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerMenuServiceServer).NewChrome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.ui.PowerMenuService/NewChrome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerMenuServiceServer).NewChrome(ctx, req.(*NewChromeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerMenuService_CloseChrome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerMenuServiceServer).CloseChrome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.ui.PowerMenuService/CloseChrome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerMenuServiceServer).CloseChrome(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerMenuService_PowerMenuPresent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerMenuServiceServer).PowerMenuPresent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.ui.PowerMenuService/PowerMenuPresent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerMenuServiceServer).PowerMenuPresent(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerMenuService_PowerMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerMenuServiceServer).PowerMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.ui.PowerMenuService/PowerMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerMenuServiceServer).PowerMenuItem(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PowerMenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.ui.PowerMenuService",
	HandlerType: (*PowerMenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewChrome",
			Handler:    _PowerMenuService_NewChrome_Handler,
		},
		{
			MethodName: "CloseChrome",
			Handler:    _PowerMenuService_CloseChrome_Handler,
		},
		{
			MethodName: "PowerMenuPresent",
			Handler:    _PowerMenuService_PowerMenuPresent_Handler,
		},
		{
			MethodName: "PowerMenuItem",
			Handler:    _PowerMenuService_PowerMenuItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "check_power_menu_service.proto",
}
