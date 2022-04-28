// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: freeze_fuse_service.proto

package filemanager

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

type TestMountZipAndSuspendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ZipDataPath string `protobuf:"bytes,3,opt,name=zip_data_path,json=zipDataPath,proto3" json:"zip_data_path,omitempty"`
}

func (x *TestMountZipAndSuspendRequest) Reset() {
	*x = TestMountZipAndSuspendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freeze_fuse_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMountZipAndSuspendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMountZipAndSuspendRequest) ProtoMessage() {}

func (x *TestMountZipAndSuspendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_freeze_fuse_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMountZipAndSuspendRequest.ProtoReflect.Descriptor instead.
func (*TestMountZipAndSuspendRequest) Descriptor() ([]byte, []int) {
	return file_freeze_fuse_service_proto_rawDescGZIP(), []int{0}
}

func (x *TestMountZipAndSuspendRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *TestMountZipAndSuspendRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TestMountZipAndSuspendRequest) GetZipDataPath() string {
	if x != nil {
		return x.ZipDataPath
	}
	return ""
}

var File_freeze_fuse_service_proto protoreflect.FileDescriptor

var file_freeze_fuse_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x5f, 0x66, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x61, 0x73,
	0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x1d, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x5a, 0x69, 0x70, 0x41,
	0x6e, 0x64, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x7a, 0x69, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x7a, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x61, 0x74, 0x68, 0x32, 0x7d, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x46, 0x55,
	0x53, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x16, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x5a, 0x69, 0x70, 0x41, 0x6e, 0x64, 0x53, 0x75, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x12, 0x34, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x5a, 0x69, 0x70, 0x41, 0x6e, 0x64, 0x53, 0x75, 0x73, 0x70, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f,
	0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x72, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freeze_fuse_service_proto_rawDescOnce sync.Once
	file_freeze_fuse_service_proto_rawDescData = file_freeze_fuse_service_proto_rawDesc
)

func file_freeze_fuse_service_proto_rawDescGZIP() []byte {
	file_freeze_fuse_service_proto_rawDescOnce.Do(func() {
		file_freeze_fuse_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_freeze_fuse_service_proto_rawDescData)
	})
	return file_freeze_fuse_service_proto_rawDescData
}

var file_freeze_fuse_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_freeze_fuse_service_proto_goTypes = []interface{}{
	(*TestMountZipAndSuspendRequest)(nil), // 0: tast.cros.filemanager.TestMountZipAndSuspendRequest
	(*empty.Empty)(nil),                   // 1: google.protobuf.Empty
}
var file_freeze_fuse_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.filemanager.FreezeFUSEService.TestMountZipAndSuspend:input_type -> tast.cros.filemanager.TestMountZipAndSuspendRequest
	1, // 1: tast.cros.filemanager.FreezeFUSEService.TestMountZipAndSuspend:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_freeze_fuse_service_proto_init() }
func file_freeze_fuse_service_proto_init() {
	if File_freeze_fuse_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freeze_fuse_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMountZipAndSuspendRequest); i {
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
			RawDescriptor: file_freeze_fuse_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_freeze_fuse_service_proto_goTypes,
		DependencyIndexes: file_freeze_fuse_service_proto_depIdxs,
		MessageInfos:      file_freeze_fuse_service_proto_msgTypes,
	}.Build()
	File_freeze_fuse_service_proto = out.File
	file_freeze_fuse_service_proto_rawDesc = nil
	file_freeze_fuse_service_proto_goTypes = nil
	file_freeze_fuse_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FreezeFUSEServiceClient is the client API for FreezeFUSEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FreezeFUSEServiceClient interface {
	TestMountZipAndSuspend(ctx context.Context, in *TestMountZipAndSuspendRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type freezeFUSEServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreezeFUSEServiceClient(cc grpc.ClientConnInterface) FreezeFUSEServiceClient {
	return &freezeFUSEServiceClient{cc}
}

func (c *freezeFUSEServiceClient) TestMountZipAndSuspend(ctx context.Context, in *TestMountZipAndSuspendRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.filemanager.FreezeFUSEService/TestMountZipAndSuspend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreezeFUSEServiceServer is the server API for FreezeFUSEService service.
type FreezeFUSEServiceServer interface {
	TestMountZipAndSuspend(context.Context, *TestMountZipAndSuspendRequest) (*empty.Empty, error)
}

// UnimplementedFreezeFUSEServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFreezeFUSEServiceServer struct {
}

func (*UnimplementedFreezeFUSEServiceServer) TestMountZipAndSuspend(context.Context, *TestMountZipAndSuspendRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMountZipAndSuspend not implemented")
}

func RegisterFreezeFUSEServiceServer(s *grpc.Server, srv FreezeFUSEServiceServer) {
	s.RegisterService(&_FreezeFUSEService_serviceDesc, srv)
}

func _FreezeFUSEService_TestMountZipAndSuspend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMountZipAndSuspendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreezeFUSEServiceServer).TestMountZipAndSuspend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.filemanager.FreezeFUSEService/TestMountZipAndSuspend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreezeFUSEServiceServer).TestMountZipAndSuspend(ctx, req.(*TestMountZipAndSuspendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FreezeFUSEService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.filemanager.FreezeFUSEService",
	HandlerType: (*FreezeFUSEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMountZipAndSuspend",
			Handler:    _FreezeFUSEService_TestMountZipAndSuspend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "freeze_fuse_service.proto",
}
