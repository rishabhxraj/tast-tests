// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: camerabox/common.proto

package camerabox

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

type Facing int32

const (
	Facing_FACING_UNSET Facing = 0
	// DUT's back camera is facing to chart tablet.
	Facing_FACING_BACK Facing = 1
	// DUT's front camera is facing to chart tablet.
	Facing_FACING_FRONT Facing = 2
)

// Enum value maps for Facing.
var (
	Facing_name = map[int32]string{
		0: "FACING_UNSET",
		1: "FACING_BACK",
		2: "FACING_FRONT",
	}
	Facing_value = map[string]int32{
		"FACING_UNSET": 0,
		"FACING_BACK":  1,
		"FACING_FRONT": 2,
	}
)

func (x Facing) Enum() *Facing {
	p := new(Facing)
	*p = x
	return p
}

func (x Facing) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Facing) Descriptor() protoreflect.EnumDescriptor {
	return file_camerabox_common_proto_enumTypes[0].Descriptor()
}

func (Facing) Type() protoreflect.EnumType {
	return &file_camerabox_common_proto_enumTypes[0]
}

func (x Facing) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Facing.Descriptor instead.
func (Facing) EnumDescriptor() ([]byte, []int) {
	return file_camerabox_common_proto_rawDescGZIP(), []int{0}
}

type TestResult int32

const (
	TestResult_TEST_RESULT_UNSET TestResult = 0
	// Test is passed.
	TestResult_TEST_RESULT_PASSED TestResult = 1
	// Test is failed.
	TestResult_TEST_RESULT_FAILED TestResult = 2
)

// Enum value maps for TestResult.
var (
	TestResult_name = map[int32]string{
		0: "TEST_RESULT_UNSET",
		1: "TEST_RESULT_PASSED",
		2: "TEST_RESULT_FAILED",
	}
	TestResult_value = map[string]int32{
		"TEST_RESULT_UNSET":  0,
		"TEST_RESULT_PASSED": 1,
		"TEST_RESULT_FAILED": 2,
	}
)

func (x TestResult) Enum() *TestResult {
	p := new(TestResult)
	*p = x
	return p
}

func (x TestResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestResult) Descriptor() protoreflect.EnumDescriptor {
	return file_camerabox_common_proto_enumTypes[1].Descriptor()
}

func (TestResult) Type() protoreflect.EnumType {
	return &file_camerabox_common_proto_enumTypes[1]
}

func (x TestResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestResult.Descriptor instead.
func (TestResult) EnumDescriptor() ([]byte, []int) {
	return file_camerabox_common_proto_rawDescGZIP(), []int{1}
}

var File_camerabox_common_proto protoreflect.FileDescriptor

var file_camerabox_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x2a, 0x3d, 0x0a,
	0x06, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x41, 0x43, 0x49, 0x4e,
	0x47, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x43,
	0x49, 0x4e, 0x47, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x41,
	0x43, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x52, 0x4f, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x53, 0x0a, 0x0a,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x53,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x02, 0x42, 0x29, 0x5a, 0x27, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f,
	0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_camerabox_common_proto_rawDescOnce sync.Once
	file_camerabox_common_proto_rawDescData = file_camerabox_common_proto_rawDesc
)

func file_camerabox_common_proto_rawDescGZIP() []byte {
	file_camerabox_common_proto_rawDescOnce.Do(func() {
		file_camerabox_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_camerabox_common_proto_rawDescData)
	})
	return file_camerabox_common_proto_rawDescData
}

var file_camerabox_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_camerabox_common_proto_goTypes = []interface{}{
	(Facing)(0),     // 0: tast.cros.camerabox.Facing
	(TestResult)(0), // 1: tast.cros.camerabox.TestResult
}
var file_camerabox_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_camerabox_common_proto_init() }
func file_camerabox_common_proto_init() {
	if File_camerabox_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_camerabox_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_camerabox_common_proto_goTypes,
		DependencyIndexes: file_camerabox_common_proto_depIdxs,
		EnumInfos:         file_camerabox_common_proto_enumTypes,
	}.Build()
	File_camerabox_common_proto = out.File
	file_camerabox_common_proto_rawDesc = nil
	file_camerabox_common_proto_goTypes = nil
	file_camerabox_common_proto_depIdxs = nil
}