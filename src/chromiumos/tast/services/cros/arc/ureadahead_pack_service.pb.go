// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: ureadahead_pack_service.proto

package arc

import (
	context "context"
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

type UreadaheadPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Credentials to perform login
	Creds string `protobuf:"bytes,5,opt,name=creds,proto3" json:"creds,omitempty"`
}

func (x *UreadaheadPackRequest) Reset() {
	*x = UreadaheadPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ureadahead_pack_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UreadaheadPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UreadaheadPackRequest) ProtoMessage() {}

func (x *UreadaheadPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ureadahead_pack_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UreadaheadPackRequest.ProtoReflect.Descriptor instead.
func (*UreadaheadPackRequest) Descriptor() ([]byte, []int) {
	return file_ureadahead_pack_service_proto_rawDescGZIP(), []int{0}
}

func (x *UreadaheadPackRequest) GetCreds() string {
	if x != nil {
		return x.Creds
	}
	return ""
}

type UreadaheadPackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to generated ureadahead pack.
	PackPath string `protobuf:"bytes,1,opt,name=pack_path,json=packPath,proto3" json:"pack_path,omitempty"`
	// Path to vm generated ureadahead pack.
	VmPackPath string `protobuf:"bytes,2,opt,name=vm_pack_path,json=vmPackPath,proto3" json:"vm_pack_path,omitempty"`
	// Path to log for ureadahead pack generation.
	LogPath string `protobuf:"bytes,3,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
}

func (x *UreadaheadPackResponse) Reset() {
	*x = UreadaheadPackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ureadahead_pack_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UreadaheadPackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UreadaheadPackResponse) ProtoMessage() {}

func (x *UreadaheadPackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ureadahead_pack_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UreadaheadPackResponse.ProtoReflect.Descriptor instead.
func (*UreadaheadPackResponse) Descriptor() ([]byte, []int) {
	return file_ureadahead_pack_service_proto_rawDescGZIP(), []int{1}
}

func (x *UreadaheadPackResponse) GetPackPath() string {
	if x != nil {
		return x.PackPath
	}
	return ""
}

func (x *UreadaheadPackResponse) GetVmPackPath() string {
	if x != nil {
		return x.VmPackPath
	}
	return ""
}

func (x *UreadaheadPackResponse) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

var File_ureadahead_pack_service_proto protoreflect.FileDescriptor

var file_ureadahead_pack_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x72, 0x65, 0x61, 0x64, 0x61, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x63, 0x22, 0x45,
	0x0a, 0x15, 0x55, 0x72, 0x65, 0x61, 0x64, 0x61, 0x68, 0x65, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x72, 0x0a, 0x16, 0x55, 0x72, 0x65, 0x61, 0x64, 0x61, 0x68,
	0x65, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0c,
	0x76, 0x6d, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x32, 0x72, 0x0a, 0x15, 0x55, 0x72, 0x65,
	0x61, 0x64, 0x61, 0x68, 0x65, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x24,
	0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x63, 0x2e, 0x55,
	0x72, 0x65, 0x61, 0x64, 0x61, 0x68, 0x65, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73,
	0x2e, 0x61, 0x72, 0x63, 0x2e, 0x55, 0x72, 0x65, 0x61, 0x64, 0x61, 0x68, 0x65, 0x61, 0x64, 0x50,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a,
	0x21, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x61,
	0x72, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ureadahead_pack_service_proto_rawDescOnce sync.Once
	file_ureadahead_pack_service_proto_rawDescData = file_ureadahead_pack_service_proto_rawDesc
)

func file_ureadahead_pack_service_proto_rawDescGZIP() []byte {
	file_ureadahead_pack_service_proto_rawDescOnce.Do(func() {
		file_ureadahead_pack_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ureadahead_pack_service_proto_rawDescData)
	})
	return file_ureadahead_pack_service_proto_rawDescData
}

var file_ureadahead_pack_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ureadahead_pack_service_proto_goTypes = []interface{}{
	(*UreadaheadPackRequest)(nil),  // 0: tast.cros.arc.UreadaheadPackRequest
	(*UreadaheadPackResponse)(nil), // 1: tast.cros.arc.UreadaheadPackResponse
}
var file_ureadahead_pack_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.arc.UreadaheadPackService.Generate:input_type -> tast.cros.arc.UreadaheadPackRequest
	1, // 1: tast.cros.arc.UreadaheadPackService.Generate:output_type -> tast.cros.arc.UreadaheadPackResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ureadahead_pack_service_proto_init() }
func file_ureadahead_pack_service_proto_init() {
	if File_ureadahead_pack_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ureadahead_pack_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UreadaheadPackRequest); i {
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
		file_ureadahead_pack_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UreadaheadPackResponse); i {
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
			RawDescriptor: file_ureadahead_pack_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ureadahead_pack_service_proto_goTypes,
		DependencyIndexes: file_ureadahead_pack_service_proto_depIdxs,
		MessageInfos:      file_ureadahead_pack_service_proto_msgTypes,
	}.Build()
	File_ureadahead_pack_service_proto = out.File
	file_ureadahead_pack_service_proto_rawDesc = nil
	file_ureadahead_pack_service_proto_goTypes = nil
	file_ureadahead_pack_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UreadaheadPackServiceClient is the client API for UreadaheadPackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UreadaheadPackServiceClient interface {
	// Generate generates ureadahead pack for requested Chrome login mode, initial or provisioned.
	Generate(ctx context.Context, in *UreadaheadPackRequest, opts ...grpc.CallOption) (*UreadaheadPackResponse, error)
}

type ureadaheadPackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUreadaheadPackServiceClient(cc grpc.ClientConnInterface) UreadaheadPackServiceClient {
	return &ureadaheadPackServiceClient{cc}
}

func (c *ureadaheadPackServiceClient) Generate(ctx context.Context, in *UreadaheadPackRequest, opts ...grpc.CallOption) (*UreadaheadPackResponse, error) {
	out := new(UreadaheadPackResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.arc.UreadaheadPackService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UreadaheadPackServiceServer is the server API for UreadaheadPackService service.
type UreadaheadPackServiceServer interface {
	// Generate generates ureadahead pack for requested Chrome login mode, initial or provisioned.
	Generate(context.Context, *UreadaheadPackRequest) (*UreadaheadPackResponse, error)
}

// UnimplementedUreadaheadPackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUreadaheadPackServiceServer struct {
}

func (*UnimplementedUreadaheadPackServiceServer) Generate(context.Context, *UreadaheadPackRequest) (*UreadaheadPackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterUreadaheadPackServiceServer(s *grpc.Server, srv UreadaheadPackServiceServer) {
	s.RegisterService(&_UreadaheadPackService_serviceDesc, srv)
}

func _UreadaheadPackService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UreadaheadPackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UreadaheadPackServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.arc.UreadaheadPackService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UreadaheadPackServiceServer).Generate(ctx, req.(*UreadaheadPackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UreadaheadPackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.arc.UreadaheadPackService",
	HandlerType: (*UreadaheadPackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _UreadaheadPackService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ureadahead_pack_service.proto",
}