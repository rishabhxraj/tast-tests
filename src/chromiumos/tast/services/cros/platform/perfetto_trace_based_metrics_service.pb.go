// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: perfetto_trace_based_metrics_service.proto

package platform

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

type GeneratePerfettoTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GeneratePerfettoTraceRequest) Reset() {
	*x = GeneratePerfettoTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_perfetto_trace_based_metrics_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePerfettoTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePerfettoTraceRequest) ProtoMessage() {}

func (x *GeneratePerfettoTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_perfetto_trace_based_metrics_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePerfettoTraceRequest.ProtoReflect.Descriptor instead.
func (*GeneratePerfettoTraceRequest) Descriptor() ([]byte, []int) {
	return file_perfetto_trace_based_metrics_service_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratePerfettoTraceRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type GeneratePerfettoTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GeneratePerfettoTraceResponse) Reset() {
	*x = GeneratePerfettoTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_perfetto_trace_based_metrics_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePerfettoTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePerfettoTraceResponse) ProtoMessage() {}

func (x *GeneratePerfettoTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_perfetto_trace_based_metrics_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePerfettoTraceResponse.ProtoReflect.Descriptor instead.
func (*GeneratePerfettoTraceResponse) Descriptor() ([]byte, []int) {
	return file_perfetto_trace_based_metrics_service_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratePerfettoTraceResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_perfetto_trace_based_metrics_service_proto protoreflect.FileDescriptor

var file_perfetto_trace_based_metrics_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x65, 0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x61,
	0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x22, 0x36, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66,
	0x65, 0x74, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x37, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xa5, 0x01, 0x0a, 0x20, 0x50, 0x65, 0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x30, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x66, 0x65, 0x74, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_perfetto_trace_based_metrics_service_proto_rawDescOnce sync.Once
	file_perfetto_trace_based_metrics_service_proto_rawDescData = file_perfetto_trace_based_metrics_service_proto_rawDesc
)

func file_perfetto_trace_based_metrics_service_proto_rawDescGZIP() []byte {
	file_perfetto_trace_based_metrics_service_proto_rawDescOnce.Do(func() {
		file_perfetto_trace_based_metrics_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_perfetto_trace_based_metrics_service_proto_rawDescData)
	})
	return file_perfetto_trace_based_metrics_service_proto_rawDescData
}

var file_perfetto_trace_based_metrics_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_perfetto_trace_based_metrics_service_proto_goTypes = []interface{}{
	(*GeneratePerfettoTraceRequest)(nil),  // 0: tast.cros.platform.GeneratePerfettoTraceRequest
	(*GeneratePerfettoTraceResponse)(nil), // 1: tast.cros.platform.GeneratePerfettoTraceResponse
}
var file_perfetto_trace_based_metrics_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.platform.PerfettoTraceBasedMetricsService.GeneratePerfettoTrace:input_type -> tast.cros.platform.GeneratePerfettoTraceRequest
	1, // 1: tast.cros.platform.PerfettoTraceBasedMetricsService.GeneratePerfettoTrace:output_type -> tast.cros.platform.GeneratePerfettoTraceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_perfetto_trace_based_metrics_service_proto_init() }
func file_perfetto_trace_based_metrics_service_proto_init() {
	if File_perfetto_trace_based_metrics_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_perfetto_trace_based_metrics_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePerfettoTraceRequest); i {
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
		file_perfetto_trace_based_metrics_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePerfettoTraceResponse); i {
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
			RawDescriptor: file_perfetto_trace_based_metrics_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_perfetto_trace_based_metrics_service_proto_goTypes,
		DependencyIndexes: file_perfetto_trace_based_metrics_service_proto_depIdxs,
		MessageInfos:      file_perfetto_trace_based_metrics_service_proto_msgTypes,
	}.Build()
	File_perfetto_trace_based_metrics_service_proto = out.File
	file_perfetto_trace_based_metrics_service_proto_rawDesc = nil
	file_perfetto_trace_based_metrics_service_proto_goTypes = nil
	file_perfetto_trace_based_metrics_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PerfettoTraceBasedMetricsServiceClient is the client API for PerfettoTraceBasedMetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PerfettoTraceBasedMetricsServiceClient interface {
	// Use perfetto to generate trace and send back to the host.
	GeneratePerfettoTrace(ctx context.Context, in *GeneratePerfettoTraceRequest, opts ...grpc.CallOption) (PerfettoTraceBasedMetricsService_GeneratePerfettoTraceClient, error)
}

type perfettoTraceBasedMetricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPerfettoTraceBasedMetricsServiceClient(cc grpc.ClientConnInterface) PerfettoTraceBasedMetricsServiceClient {
	return &perfettoTraceBasedMetricsServiceClient{cc}
}

func (c *perfettoTraceBasedMetricsServiceClient) GeneratePerfettoTrace(ctx context.Context, in *GeneratePerfettoTraceRequest, opts ...grpc.CallOption) (PerfettoTraceBasedMetricsService_GeneratePerfettoTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PerfettoTraceBasedMetricsService_serviceDesc.Streams[0], "/tast.cros.platform.PerfettoTraceBasedMetricsService/GeneratePerfettoTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &perfettoTraceBasedMetricsServiceGeneratePerfettoTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PerfettoTraceBasedMetricsService_GeneratePerfettoTraceClient interface {
	Recv() (*GeneratePerfettoTraceResponse, error)
	grpc.ClientStream
}

type perfettoTraceBasedMetricsServiceGeneratePerfettoTraceClient struct {
	grpc.ClientStream
}

func (x *perfettoTraceBasedMetricsServiceGeneratePerfettoTraceClient) Recv() (*GeneratePerfettoTraceResponse, error) {
	m := new(GeneratePerfettoTraceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PerfettoTraceBasedMetricsServiceServer is the server API for PerfettoTraceBasedMetricsService service.
type PerfettoTraceBasedMetricsServiceServer interface {
	// Use perfetto to generate trace and send back to the host.
	GeneratePerfettoTrace(*GeneratePerfettoTraceRequest, PerfettoTraceBasedMetricsService_GeneratePerfettoTraceServer) error
}

// UnimplementedPerfettoTraceBasedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPerfettoTraceBasedMetricsServiceServer struct {
}

func (*UnimplementedPerfettoTraceBasedMetricsServiceServer) GeneratePerfettoTrace(*GeneratePerfettoTraceRequest, PerfettoTraceBasedMetricsService_GeneratePerfettoTraceServer) error {
	return status.Errorf(codes.Unimplemented, "method GeneratePerfettoTrace not implemented")
}

func RegisterPerfettoTraceBasedMetricsServiceServer(s *grpc.Server, srv PerfettoTraceBasedMetricsServiceServer) {
	s.RegisterService(&_PerfettoTraceBasedMetricsService_serviceDesc, srv)
}

func _PerfettoTraceBasedMetricsService_GeneratePerfettoTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GeneratePerfettoTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PerfettoTraceBasedMetricsServiceServer).GeneratePerfettoTrace(m, &perfettoTraceBasedMetricsServiceGeneratePerfettoTraceServer{stream})
}

type PerfettoTraceBasedMetricsService_GeneratePerfettoTraceServer interface {
	Send(*GeneratePerfettoTraceResponse) error
	grpc.ServerStream
}

type perfettoTraceBasedMetricsServiceGeneratePerfettoTraceServer struct {
	grpc.ServerStream
}

func (x *perfettoTraceBasedMetricsServiceGeneratePerfettoTraceServer) Send(m *GeneratePerfettoTraceResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PerfettoTraceBasedMetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.platform.PerfettoTraceBasedMetricsService",
	HandlerType: (*PerfettoTraceBasedMetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GeneratePerfettoTrace",
			Handler:       _PerfettoTraceBasedMetricsService_GeneratePerfettoTrace_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "perfetto_trace_based_metrics_service.proto",
}
