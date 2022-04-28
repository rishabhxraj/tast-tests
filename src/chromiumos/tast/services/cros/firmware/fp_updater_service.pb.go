// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: fp_updater_service.proto

package firmware

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

type ReadFpUpdaterLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestLog   string `protobuf:"bytes,1,opt,name=latest_log,json=latestLog,proto3" json:"latest_log,omitempty"`
	PreviousLog string `protobuf:"bytes,2,opt,name=previous_log,json=previousLog,proto3" json:"previous_log,omitempty"`
}

func (x *ReadFpUpdaterLogsResponse) Reset() {
	*x = ReadFpUpdaterLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fp_updater_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFpUpdaterLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFpUpdaterLogsResponse) ProtoMessage() {}

func (x *ReadFpUpdaterLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fp_updater_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFpUpdaterLogsResponse.ProtoReflect.Descriptor instead.
func (*ReadFpUpdaterLogsResponse) Descriptor() ([]byte, []int) {
	return file_fp_updater_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReadFpUpdaterLogsResponse) GetLatestLog() string {
	if x != nil {
		return x.LatestLog
	}
	return ""
}

func (x *ReadFpUpdaterLogsResponse) GetPreviousLog() string {
	if x != nil {
		return x.PreviousLog
	}
	return ""
}

var File_fp_updater_service_proto protoreflect.FileDescriptor

var file_fp_updater_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x70, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x19, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4c, 0x6f, 0x67, 0x32, 0x6e, 0x0a, 0x10, 0x46, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a,
	0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x74, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fp_updater_service_proto_rawDescOnce sync.Once
	file_fp_updater_service_proto_rawDescData = file_fp_updater_service_proto_rawDesc
)

func file_fp_updater_service_proto_rawDescGZIP() []byte {
	file_fp_updater_service_proto_rawDescOnce.Do(func() {
		file_fp_updater_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_fp_updater_service_proto_rawDescData)
	})
	return file_fp_updater_service_proto_rawDescData
}

var file_fp_updater_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fp_updater_service_proto_goTypes = []interface{}{
	(*ReadFpUpdaterLogsResponse)(nil), // 0: tast.cros.firmware.ReadFpUpdaterLogsResponse
	(*empty.Empty)(nil),               // 1: google.protobuf.Empty
}
var file_fp_updater_service_proto_depIdxs = []int32{
	1, // 0: tast.cros.firmware.FpUpdaterService.ReadUpdaterLogs:input_type -> google.protobuf.Empty
	0, // 1: tast.cros.firmware.FpUpdaterService.ReadUpdaterLogs:output_type -> tast.cros.firmware.ReadFpUpdaterLogsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fp_updater_service_proto_init() }
func file_fp_updater_service_proto_init() {
	if File_fp_updater_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fp_updater_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFpUpdaterLogsResponse); i {
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
			RawDescriptor: file_fp_updater_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fp_updater_service_proto_goTypes,
		DependencyIndexes: file_fp_updater_service_proto_depIdxs,
		MessageInfos:      file_fp_updater_service_proto_msgTypes,
	}.Build()
	File_fp_updater_service_proto = out.File
	file_fp_updater_service_proto_rawDesc = nil
	file_fp_updater_service_proto_goTypes = nil
	file_fp_updater_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FpUpdaterServiceClient is the client API for FpUpdaterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FpUpdaterServiceClient interface {
	// ReadUpdaterLogs reads the latest and previous logs from the fingerprint firmware updater.
	ReadUpdaterLogs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadFpUpdaterLogsResponse, error)
}

type fpUpdaterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFpUpdaterServiceClient(cc grpc.ClientConnInterface) FpUpdaterServiceClient {
	return &fpUpdaterServiceClient{cc}
}

func (c *fpUpdaterServiceClient) ReadUpdaterLogs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadFpUpdaterLogsResponse, error) {
	out := new(ReadFpUpdaterLogsResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.firmware.FpUpdaterService/ReadUpdaterLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FpUpdaterServiceServer is the server API for FpUpdaterService service.
type FpUpdaterServiceServer interface {
	// ReadUpdaterLogs reads the latest and previous logs from the fingerprint firmware updater.
	ReadUpdaterLogs(context.Context, *empty.Empty) (*ReadFpUpdaterLogsResponse, error)
}

// UnimplementedFpUpdaterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFpUpdaterServiceServer struct {
}

func (*UnimplementedFpUpdaterServiceServer) ReadUpdaterLogs(context.Context, *empty.Empty) (*ReadFpUpdaterLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUpdaterLogs not implemented")
}

func RegisterFpUpdaterServiceServer(s *grpc.Server, srv FpUpdaterServiceServer) {
	s.RegisterService(&_FpUpdaterService_serviceDesc, srv)
}

func _FpUpdaterService_ReadUpdaterLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FpUpdaterServiceServer).ReadUpdaterLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.firmware.FpUpdaterService/ReadUpdaterLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FpUpdaterServiceServer).ReadUpdaterLogs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _FpUpdaterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.firmware.FpUpdaterService",
	HandlerType: (*FpUpdaterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadUpdaterLogs",
			Handler:    _FpUpdaterService_ReadUpdaterLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fp_updater_service.proto",
}