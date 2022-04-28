// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: camerabox/alignment_service.proto

package camerabox

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

type ManualAlignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Absolute path for saving data used on DUT.
	DataPath string `protobuf:"bytes,1,opt,name=data_path,json=dataPath,proto3" json:"data_path,omitempty"`
	// Username to login chrome and prepare chrome remote desktop.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Password to login chrome and prepare chrome remote desktop.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// DUT's target camera facing to be aligned.
	Facing Facing `protobuf:"varint,4,opt,name=facing,proto3,enum=tast.cros.camerabox.Facing" json:"facing,omitempty"`
}

func (x *ManualAlignRequest) Reset() {
	*x = ManualAlignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camerabox_alignment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualAlignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualAlignRequest) ProtoMessage() {}

func (x *ManualAlignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_camerabox_alignment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualAlignRequest.ProtoReflect.Descriptor instead.
func (*ManualAlignRequest) Descriptor() ([]byte, []int) {
	return file_camerabox_alignment_service_proto_rawDescGZIP(), []int{0}
}

func (x *ManualAlignRequest) GetDataPath() string {
	if x != nil {
		return x.DataPath
	}
	return ""
}

func (x *ManualAlignRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ManualAlignRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManualAlignRequest) GetFacing() Facing {
	if x != nil {
		return x.Facing
	}
	return Facing_FACING_UNSET
}

type CheckRegressionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Absolute path for saving data used on DUT.
	DataPath string `protobuf:"bytes,1,opt,name=data_path,json=dataPath,proto3" json:"data_path,omitempty"`
	// DUT's target camera facing to be aligned.
	Facing Facing `protobuf:"varint,2,opt,name=facing,proto3,enum=tast.cros.camerabox.Facing" json:"facing,omitempty"`
}

func (x *CheckRegressionRequest) Reset() {
	*x = CheckRegressionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camerabox_alignment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRegressionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRegressionRequest) ProtoMessage() {}

func (x *CheckRegressionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_camerabox_alignment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRegressionRequest.ProtoReflect.Descriptor instead.
func (*CheckRegressionRequest) Descriptor() ([]byte, []int) {
	return file_camerabox_alignment_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRegressionRequest) GetDataPath() string {
	if x != nil {
		return x.DataPath
	}
	return ""
}

func (x *CheckRegressionRequest) GetFacing() Facing {
	if x != nil {
		return x.Facing
	}
	return Facing_FACING_UNSET
}

type CheckRegressionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Check result.
	Result TestResult `protobuf:"varint,1,opt,name=result,proto3,enum=tast.cros.camerabox.TestResult" json:"result,omitempty"`
	// Error message from running check.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CheckRegressionResponse) Reset() {
	*x = CheckRegressionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camerabox_alignment_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRegressionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRegressionResponse) ProtoMessage() {}

func (x *CheckRegressionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_camerabox_alignment_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRegressionResponse.ProtoReflect.Descriptor instead.
func (*CheckRegressionResponse) Descriptor() ([]byte, []int) {
	return file_camerabox_alignment_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRegressionResponse) GetResult() TestResult {
	if x != nil {
		return x.Result
	}
	return TestResult_TEST_RESULT_UNSET
}

func (x *CheckRegressionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_camerabox_alignment_service_proto protoreflect.FileDescriptor

var file_camerabox_alignment_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2f, 0x61, 0x6c, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x63,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x1a, 0x16, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x62, 0x6f, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2e,
	0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x6a,
	0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2e, 0x46, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x68, 0x0a, 0x17, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0xd4, 0x01, 0x0a, 0x10, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x12, 0x27, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2e, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0f, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b,
	0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x62, 0x6f, 0x78, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x61,
	0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f,
	0x78, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_camerabox_alignment_service_proto_rawDescOnce sync.Once
	file_camerabox_alignment_service_proto_rawDescData = file_camerabox_alignment_service_proto_rawDesc
)

func file_camerabox_alignment_service_proto_rawDescGZIP() []byte {
	file_camerabox_alignment_service_proto_rawDescOnce.Do(func() {
		file_camerabox_alignment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_camerabox_alignment_service_proto_rawDescData)
	})
	return file_camerabox_alignment_service_proto_rawDescData
}

var file_camerabox_alignment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_camerabox_alignment_service_proto_goTypes = []interface{}{
	(*ManualAlignRequest)(nil),      // 0: tast.cros.camerabox.ManualAlignRequest
	(*CheckRegressionRequest)(nil),  // 1: tast.cros.camerabox.CheckRegressionRequest
	(*CheckRegressionResponse)(nil), // 2: tast.cros.camerabox.CheckRegressionResponse
	(Facing)(0),                     // 3: tast.cros.camerabox.Facing
	(TestResult)(0),                 // 4: tast.cros.camerabox.TestResult
	(*empty.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_camerabox_alignment_service_proto_depIdxs = []int32{
	3, // 0: tast.cros.camerabox.ManualAlignRequest.facing:type_name -> tast.cros.camerabox.Facing
	3, // 1: tast.cros.camerabox.CheckRegressionRequest.facing:type_name -> tast.cros.camerabox.Facing
	4, // 2: tast.cros.camerabox.CheckRegressionResponse.result:type_name -> tast.cros.camerabox.TestResult
	0, // 3: tast.cros.camerabox.AlignmentService.ManualAlign:input_type -> tast.cros.camerabox.ManualAlignRequest
	1, // 4: tast.cros.camerabox.AlignmentService.CheckRegression:input_type -> tast.cros.camerabox.CheckRegressionRequest
	5, // 5: tast.cros.camerabox.AlignmentService.ManualAlign:output_type -> google.protobuf.Empty
	2, // 6: tast.cros.camerabox.AlignmentService.CheckRegression:output_type -> tast.cros.camerabox.CheckRegressionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_camerabox_alignment_service_proto_init() }
func file_camerabox_alignment_service_proto_init() {
	if File_camerabox_alignment_service_proto != nil {
		return
	}
	file_camerabox_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_camerabox_alignment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManualAlignRequest); i {
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
		file_camerabox_alignment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRegressionRequest); i {
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
		file_camerabox_alignment_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRegressionResponse); i {
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
			RawDescriptor: file_camerabox_alignment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_camerabox_alignment_service_proto_goTypes,
		DependencyIndexes: file_camerabox_alignment_service_proto_depIdxs,
		MessageInfos:      file_camerabox_alignment_service_proto_msgTypes,
	}.Build()
	File_camerabox_alignment_service_proto = out.File
	file_camerabox_alignment_service_proto_rawDesc = nil
	file_camerabox_alignment_service_proto_goTypes = nil
	file_camerabox_alignment_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlignmentServiceClient is the client API for AlignmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlignmentServiceClient interface {
	// ManualAlign opens preview page on DUT and wait until preview is aligned.
	ManualAlign(ctx context.Context, in *ManualAlignRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// CheckRegression opens preview page on DUT and check preview is aligned as
	// regression test.
	CheckRegression(ctx context.Context, in *CheckRegressionRequest, opts ...grpc.CallOption) (*CheckRegressionResponse, error)
}

type alignmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlignmentServiceClient(cc grpc.ClientConnInterface) AlignmentServiceClient {
	return &alignmentServiceClient{cc}
}

func (c *alignmentServiceClient) ManualAlign(ctx context.Context, in *ManualAlignRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.camerabox.AlignmentService/ManualAlign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alignmentServiceClient) CheckRegression(ctx context.Context, in *CheckRegressionRequest, opts ...grpc.CallOption) (*CheckRegressionResponse, error) {
	out := new(CheckRegressionResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.camerabox.AlignmentService/CheckRegression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlignmentServiceServer is the server API for AlignmentService service.
type AlignmentServiceServer interface {
	// ManualAlign opens preview page on DUT and wait until preview is aligned.
	ManualAlign(context.Context, *ManualAlignRequest) (*empty.Empty, error)
	// CheckRegression opens preview page on DUT and check preview is aligned as
	// regression test.
	CheckRegression(context.Context, *CheckRegressionRequest) (*CheckRegressionResponse, error)
}

// UnimplementedAlignmentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlignmentServiceServer struct {
}

func (*UnimplementedAlignmentServiceServer) ManualAlign(context.Context, *ManualAlignRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualAlign not implemented")
}
func (*UnimplementedAlignmentServiceServer) CheckRegression(context.Context, *CheckRegressionRequest) (*CheckRegressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRegression not implemented")
}

func RegisterAlignmentServiceServer(s *grpc.Server, srv AlignmentServiceServer) {
	s.RegisterService(&_AlignmentService_serviceDesc, srv)
}

func _AlignmentService_ManualAlign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualAlignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlignmentServiceServer).ManualAlign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.camerabox.AlignmentService/ManualAlign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlignmentServiceServer).ManualAlign(ctx, req.(*ManualAlignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlignmentService_CheckRegression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRegressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlignmentServiceServer).CheckRegression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.camerabox.AlignmentService/CheckRegression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlignmentServiceServer).CheckRegression(ctx, req.(*CheckRegressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlignmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.camerabox.AlignmentService",
	HandlerType: (*AlignmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManualAlign",
			Handler:    _AlignmentService_ManualAlign_Handler,
		},
		{
			MethodName: "CheckRegression",
			Handler:    _AlignmentService_CheckRegression_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "camerabox/alignment_service.proto",
}
