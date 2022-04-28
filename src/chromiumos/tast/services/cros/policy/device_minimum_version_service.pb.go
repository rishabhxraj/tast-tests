// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: policy/device_minimum_version_service.proto

package policy

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_policy_device_minimum_version_service_proto protoreflect.FileDescriptor

var file_policy_device_minimum_version_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x74, 0x0a, 0x1b,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x21, 0x54,
	0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x49, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73,
	0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x72, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_policy_device_minimum_version_service_proto_goTypes = []interface{}{
	(*empty.Empty)(nil), // 0: google.protobuf.Empty
}
var file_policy_device_minimum_version_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.policy.DeviceMinimumVersionService.TestUpdateRequiredScreenIsVisible:input_type -> google.protobuf.Empty
	0, // 1: tast.cros.policy.DeviceMinimumVersionService.TestUpdateRequiredScreenIsVisible:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_policy_device_minimum_version_service_proto_init() }
func file_policy_device_minimum_version_service_proto_init() {
	if File_policy_device_minimum_version_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policy_device_minimum_version_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policy_device_minimum_version_service_proto_goTypes,
		DependencyIndexes: file_policy_device_minimum_version_service_proto_depIdxs,
	}.Build()
	File_policy_device_minimum_version_service_proto = out.File
	file_policy_device_minimum_version_service_proto_rawDesc = nil
	file_policy_device_minimum_version_service_proto_goTypes = nil
	file_policy_device_minimum_version_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceMinimumVersionServiceClient is the client API for DeviceMinimumVersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceMinimumVersionServiceClient interface {
	// Creates a new instance of Chrome using the state from the existing one.
	// Checks that an update required screen with update now button is visible on the login page.
	// Chrome is closed when function exists. This is used by the test policy.DeviceMinimumVersion.
	TestUpdateRequiredScreenIsVisible(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type deviceMinimumVersionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceMinimumVersionServiceClient(cc grpc.ClientConnInterface) DeviceMinimumVersionServiceClient {
	return &deviceMinimumVersionServiceClient{cc}
}

func (c *deviceMinimumVersionServiceClient) TestUpdateRequiredScreenIsVisible(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.policy.DeviceMinimumVersionService/TestUpdateRequiredScreenIsVisible", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceMinimumVersionServiceServer is the server API for DeviceMinimumVersionService service.
type DeviceMinimumVersionServiceServer interface {
	// Creates a new instance of Chrome using the state from the existing one.
	// Checks that an update required screen with update now button is visible on the login page.
	// Chrome is closed when function exists. This is used by the test policy.DeviceMinimumVersion.
	TestUpdateRequiredScreenIsVisible(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedDeviceMinimumVersionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceMinimumVersionServiceServer struct {
}

func (*UnimplementedDeviceMinimumVersionServiceServer) TestUpdateRequiredScreenIsVisible(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestUpdateRequiredScreenIsVisible not implemented")
}

func RegisterDeviceMinimumVersionServiceServer(s *grpc.Server, srv DeviceMinimumVersionServiceServer) {
	s.RegisterService(&_DeviceMinimumVersionService_serviceDesc, srv)
}

func _DeviceMinimumVersionService_TestUpdateRequiredScreenIsVisible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMinimumVersionServiceServer).TestUpdateRequiredScreenIsVisible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.policy.DeviceMinimumVersionService/TestUpdateRequiredScreenIsVisible",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMinimumVersionServiceServer).TestUpdateRequiredScreenIsVisible(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceMinimumVersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.policy.DeviceMinimumVersionService",
	HandlerType: (*DeviceMinimumVersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestUpdateRequiredScreenIsVisible",
			Handler:    _DeviceMinimumVersionService_TestUpdateRequiredScreenIsVisible_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy/device_minimum_version_service.proto",
}
