// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: update_test_service.proto

package lacros

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

// BrowserType is to specify the types of supported browsers.
type BrowserType int32

const (
	BrowserType_UNKNOWN         BrowserType = 0
	BrowserType_ASH             BrowserType = 1
	BrowserType_LACROS_ROOTFS   BrowserType = 2
	BrowserType_LACROS_STATEFUL BrowserType = 3
)

// Enum value maps for BrowserType.
var (
	BrowserType_name = map[int32]string{
		0: "UNKNOWN",
		1: "ASH",
		2: "LACROS_ROOTFS",
		3: "LACROS_STATEFUL",
	}
	BrowserType_value = map[string]int32{
		"UNKNOWN":         0,
		"ASH":             1,
		"LACROS_ROOTFS":   2,
		"LACROS_STATEFUL": 3,
	}
)

func (x BrowserType) Enum() *BrowserType {
	p := new(BrowserType)
	*p = x
	return p
}

func (x BrowserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrowserType) Descriptor() protoreflect.EnumDescriptor {
	return file_update_test_service_proto_enumTypes[0].Descriptor()
}

func (BrowserType) Type() protoreflect.EnumType {
	return &file_update_test_service_proto_enumTypes[0]
}

func (x BrowserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrowserType.Descriptor instead.
func (BrowserType) EnumDescriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{0}
}

type TestResult_Status int32

const (
	TestResult_NO_STATUS TestResult_Status = 0
	TestResult_PASSED    TestResult_Status = 1
	TestResult_FAILED    TestResult_Status = 2
)

// Enum value maps for TestResult_Status.
var (
	TestResult_Status_name = map[int32]string{
		0: "NO_STATUS",
		1: "PASSED",
		2: "FAILED",
	}
	TestResult_Status_value = map[string]int32{
		"NO_STATUS": 0,
		"PASSED":    1,
		"FAILED":    2,
	}
)

func (x TestResult_Status) Enum() *TestResult_Status {
	p := new(TestResult_Status)
	*p = x
	return p
}

func (x TestResult_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestResult_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_update_test_service_proto_enumTypes[1].Descriptor()
}

func (TestResult_Status) Type() protoreflect.EnumType {
	return &file_update_test_service_proto_enumTypes[1]
}

func (x TestResult_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestResult_Status.Descriptor instead.
func (TestResult_Status) EnumDescriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{1, 0}
}

// BrowserContext is a shared info to configure or check the browser under test.
type BrowserContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of supported browsers.
	Browser BrowserType `protobuf:"varint,1,opt,name=browser,proto3,enum=tast.cros.lacros.BrowserType" json:"browser,omitempty"`
	// Chrome options used to launch browser.
	Opts []string `protobuf:"bytes,2,rep,name=opts,proto3" json:"opts,omitempty"`
}

func (x *BrowserContext) Reset() {
	*x = BrowserContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserContext) ProtoMessage() {}

func (x *BrowserContext) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserContext.ProtoReflect.Descriptor instead.
func (*BrowserContext) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{0}
}

func (x *BrowserContext) GetBrowser() BrowserType {
	if x != nil {
		return x.Browser
	}
	return BrowserType_UNKNOWN
}

func (x *BrowserContext) GetOpts() []string {
	if x != nil {
		return x.Opts
	}
	return nil
}

// TestResult is detailed test status data for a verification action in a DUT.
type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        TestResult_Status `protobuf:"varint,1,opt,name=status,proto3,enum=tast.cros.lacros.TestResult_Status" json:"status,omitempty"`
	StatusDetails string            `protobuf:"bytes,2,opt,name=status_details,json=statusDetails,proto3" json:"status_details,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{1}
}

func (x *TestResult) GetStatus() TestResult_Status {
	if x != nil {
		return x.Status
	}
	return TestResult_NO_STATUS
}

func (x *TestResult) GetStatusDetails() string {
	if x != nil {
		return x.StatusDetails
	}
	return ""
}

// VerifyUpdateRequest contains the Lacros browser info that is used
// to verify whether the expected Lacros is selected in the given context of
// provisioned browsers and Ash configs.
type VerifyUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AshContext               *BrowserContext   `protobuf:"bytes,1,opt,name=ash_context,json=ashContext,proto3" json:"ash_context,omitempty"`
	ProvisionedLacrosContext []*BrowserContext `protobuf:"bytes,2,rep,name=provisioned_lacros_context,json=provisionedLacrosContext,proto3" json:"provisioned_lacros_context,omitempty"`
	// The following fields describe the Lacros to be selected.
	ExpectedBrowser   BrowserType `protobuf:"varint,3,opt,name=expected_browser,json=expectedBrowser,proto3,enum=tast.cros.lacros.BrowserType" json:"expected_browser,omitempty"` // e.g. LACROS_STATEFUL
	ExpectedVersion   string      `protobuf:"bytes,4,opt,name=expected_version,json=expectedVersion,proto3" json:"expected_version,omitempty"`                                    // e.g. "9999.0.0.1"
	ExpectedComponent string      `protobuf:"bytes,5,opt,name=expected_component,json=expectedComponent,proto3" json:"expected_component,omitempty"`                              // e.g. "lacros-dogfood-dev" for the dev channel
	// Whether to use UI for verification
	UseUi bool `protobuf:"varint,6,opt,name=use_ui,json=useUi,proto3" json:"use_ui,omitempty"`
}

func (x *VerifyUpdateRequest) Reset() {
	*x = VerifyUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUpdateRequest) ProtoMessage() {}

func (x *VerifyUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUpdateRequest.ProtoReflect.Descriptor instead.
func (*VerifyUpdateRequest) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyUpdateRequest) GetAshContext() *BrowserContext {
	if x != nil {
		return x.AshContext
	}
	return nil
}

func (x *VerifyUpdateRequest) GetProvisionedLacrosContext() []*BrowserContext {
	if x != nil {
		return x.ProvisionedLacrosContext
	}
	return nil
}

func (x *VerifyUpdateRequest) GetExpectedBrowser() BrowserType {
	if x != nil {
		return x.ExpectedBrowser
	}
	return BrowserType_UNKNOWN
}

func (x *VerifyUpdateRequest) GetExpectedVersion() string {
	if x != nil {
		return x.ExpectedVersion
	}
	return ""
}

func (x *VerifyUpdateRequest) GetExpectedComponent() string {
	if x != nil {
		return x.ExpectedComponent
	}
	return ""
}

func (x *VerifyUpdateRequest) GetUseUi() bool {
	if x != nil {
		return x.UseUi
	}
	return false
}

// VerifyUpdateResponse contains a test result of version comparison for
// a single action of simulated autoupdate.
type VerifyUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *TestResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *VerifyUpdateResponse) Reset() {
	*x = VerifyUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUpdateResponse) ProtoMessage() {}

func (x *VerifyUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUpdateResponse.ProtoReflect.Descriptor instead.
func (*VerifyUpdateResponse) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyUpdateResponse) GetResult() *TestResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type ClearUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearUpdateRequest) Reset() {
	*x = ClearUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearUpdateRequest) ProtoMessage() {}

func (x *ClearUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearUpdateRequest.ProtoReflect.Descriptor instead.
func (*ClearUpdateRequest) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{4}
}

type ClearUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearUpdateResponse) Reset() {
	*x = ClearUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearUpdateResponse) ProtoMessage() {}

func (x *ClearUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearUpdateResponse.ProtoReflect.Descriptor instead.
func (*ClearUpdateResponse) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{5}
}

type GetBrowserVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Browser BrowserType `protobuf:"varint,1,opt,name=browser,proto3,enum=tast.cros.lacros.BrowserType" json:"browser,omitempty"`
}

func (x *GetBrowserVersionRequest) Reset() {
	*x = GetBrowserVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrowserVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrowserVersionRequest) ProtoMessage() {}

func (x *GetBrowserVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrowserVersionRequest.ProtoReflect.Descriptor instead.
func (*GetBrowserVersionRequest) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetBrowserVersionRequest) GetBrowser() BrowserType {
	if x != nil {
		return x.Browser
	}
	return BrowserType_UNKNOWN
}

type GetBrowserVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []string `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *GetBrowserVersionResponse) Reset() {
	*x = GetBrowserVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_test_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrowserVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrowserVersionResponse) ProtoMessage() {}

func (x *GetBrowserVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_update_test_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrowserVersionResponse.ProtoReflect.Descriptor instead.
func (*GetBrowserVersionResponse) Descriptor() ([]byte, []int) {
	return file_update_test_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetBrowserVersionResponse) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_update_test_service_proto protoreflect.FileDescriptor

var file_update_test_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x61, 0x73,
	0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x22, 0x5d, 0x0a,
	0x0e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x37, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63,
	0x72, 0x6f, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x22, 0xa1, 0x01, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x61,
	0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53, 0x53,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x22, 0xf3, 0x02, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x73, 0x68, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73,
	0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x0a, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x5e, 0x0a, 0x1a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72,
	0x6f, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x18, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x4c, 0x61,
	0x63, 0x72, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x48, 0x0a, 0x10, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x55, 0x69, 0x22, 0x4c, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x53, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2a,
	0x4b, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x53, 0x48, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x41, 0x43, 0x52, 0x4f, 0x53, 0x5f, 0x52,
	0x4f, 0x4f, 0x54, 0x46, 0x53, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x41, 0x43, 0x52, 0x4f,
	0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x46, 0x55, 0x4c, 0x10, 0x03, 0x32, 0xc2, 0x02, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x25, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c,
	0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c,
	0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72,
	0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x6c,
	0x61, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x26, 0x5a, 0x24, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f,
	0x74, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_update_test_service_proto_rawDescOnce sync.Once
	file_update_test_service_proto_rawDescData = file_update_test_service_proto_rawDesc
)

func file_update_test_service_proto_rawDescGZIP() []byte {
	file_update_test_service_proto_rawDescOnce.Do(func() {
		file_update_test_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_test_service_proto_rawDescData)
	})
	return file_update_test_service_proto_rawDescData
}

var file_update_test_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_update_test_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_update_test_service_proto_goTypes = []interface{}{
	(BrowserType)(0),                  // 0: tast.cros.lacros.BrowserType
	(TestResult_Status)(0),            // 1: tast.cros.lacros.TestResult.Status
	(*BrowserContext)(nil),            // 2: tast.cros.lacros.BrowserContext
	(*TestResult)(nil),                // 3: tast.cros.lacros.TestResult
	(*VerifyUpdateRequest)(nil),       // 4: tast.cros.lacros.VerifyUpdateRequest
	(*VerifyUpdateResponse)(nil),      // 5: tast.cros.lacros.VerifyUpdateResponse
	(*ClearUpdateRequest)(nil),        // 6: tast.cros.lacros.ClearUpdateRequest
	(*ClearUpdateResponse)(nil),       // 7: tast.cros.lacros.ClearUpdateResponse
	(*GetBrowserVersionRequest)(nil),  // 8: tast.cros.lacros.GetBrowserVersionRequest
	(*GetBrowserVersionResponse)(nil), // 9: tast.cros.lacros.GetBrowserVersionResponse
}
var file_update_test_service_proto_depIdxs = []int32{
	0,  // 0: tast.cros.lacros.BrowserContext.browser:type_name -> tast.cros.lacros.BrowserType
	1,  // 1: tast.cros.lacros.TestResult.status:type_name -> tast.cros.lacros.TestResult.Status
	2,  // 2: tast.cros.lacros.VerifyUpdateRequest.ash_context:type_name -> tast.cros.lacros.BrowserContext
	2,  // 3: tast.cros.lacros.VerifyUpdateRequest.provisioned_lacros_context:type_name -> tast.cros.lacros.BrowserContext
	0,  // 4: tast.cros.lacros.VerifyUpdateRequest.expected_browser:type_name -> tast.cros.lacros.BrowserType
	3,  // 5: tast.cros.lacros.VerifyUpdateResponse.result:type_name -> tast.cros.lacros.TestResult
	0,  // 6: tast.cros.lacros.GetBrowserVersionRequest.browser:type_name -> tast.cros.lacros.BrowserType
	4,  // 7: tast.cros.lacros.UpdateTestService.VerifyUpdate:input_type -> tast.cros.lacros.VerifyUpdateRequest
	6,  // 8: tast.cros.lacros.UpdateTestService.ClearUpdate:input_type -> tast.cros.lacros.ClearUpdateRequest
	8,  // 9: tast.cros.lacros.UpdateTestService.GetBrowserVersion:input_type -> tast.cros.lacros.GetBrowserVersionRequest
	5,  // 10: tast.cros.lacros.UpdateTestService.VerifyUpdate:output_type -> tast.cros.lacros.VerifyUpdateResponse
	7,  // 11: tast.cros.lacros.UpdateTestService.ClearUpdate:output_type -> tast.cros.lacros.ClearUpdateResponse
	9,  // 12: tast.cros.lacros.UpdateTestService.GetBrowserVersion:output_type -> tast.cros.lacros.GetBrowserVersionResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_update_test_service_proto_init() }
func file_update_test_service_proto_init() {
	if File_update_test_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_update_test_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserContext); i {
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
		file_update_test_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
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
		file_update_test_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUpdateRequest); i {
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
		file_update_test_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUpdateResponse); i {
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
		file_update_test_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearUpdateRequest); i {
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
		file_update_test_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearUpdateResponse); i {
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
		file_update_test_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrowserVersionRequest); i {
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
		file_update_test_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrowserVersionResponse); i {
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
			RawDescriptor: file_update_test_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_update_test_service_proto_goTypes,
		DependencyIndexes: file_update_test_service_proto_depIdxs,
		EnumInfos:         file_update_test_service_proto_enumTypes,
		MessageInfos:      file_update_test_service_proto_msgTypes,
	}.Build()
	File_update_test_service_proto = out.File
	file_update_test_service_proto_rawDesc = nil
	file_update_test_service_proto_goTypes = nil
	file_update_test_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UpdateTestServiceClient is the client API for UpdateTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateTestServiceClient interface {
	// VerifyUpdate sets a DUT with given contexts and checks if the
	// expected version of Lacros is loaded successfully without crash.
	VerifyUpdate(ctx context.Context, in *VerifyUpdateRequest, opts ...grpc.CallOption) (*VerifyUpdateResponse, error)
	// ClearUpdate removes provisioned Lacros in the install path or browser data
	// if needed.
	ClearUpdate(ctx context.Context, in *ClearUpdateRequest, opts ...grpc.CallOption) (*ClearUpdateResponse, error)
	// GetBrowserVersion returns version info of the given browser type.
	// If multiple Lacros browsers are provisioned in the stateful partition,
	// all the versions will be returned.
	GetBrowserVersion(ctx context.Context, in *GetBrowserVersionRequest, opts ...grpc.CallOption) (*GetBrowserVersionResponse, error)
}

type updateTestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateTestServiceClient(cc grpc.ClientConnInterface) UpdateTestServiceClient {
	return &updateTestServiceClient{cc}
}

func (c *updateTestServiceClient) VerifyUpdate(ctx context.Context, in *VerifyUpdateRequest, opts ...grpc.CallOption) (*VerifyUpdateResponse, error) {
	out := new(VerifyUpdateResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.lacros.UpdateTestService/VerifyUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateTestServiceClient) ClearUpdate(ctx context.Context, in *ClearUpdateRequest, opts ...grpc.CallOption) (*ClearUpdateResponse, error) {
	out := new(ClearUpdateResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.lacros.UpdateTestService/ClearUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateTestServiceClient) GetBrowserVersion(ctx context.Context, in *GetBrowserVersionRequest, opts ...grpc.CallOption) (*GetBrowserVersionResponse, error) {
	out := new(GetBrowserVersionResponse)
	err := c.cc.Invoke(ctx, "/tast.cros.lacros.UpdateTestService/GetBrowserVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateTestServiceServer is the server API for UpdateTestService service.
type UpdateTestServiceServer interface {
	// VerifyUpdate sets a DUT with given contexts and checks if the
	// expected version of Lacros is loaded successfully without crash.
	VerifyUpdate(context.Context, *VerifyUpdateRequest) (*VerifyUpdateResponse, error)
	// ClearUpdate removes provisioned Lacros in the install path or browser data
	// if needed.
	ClearUpdate(context.Context, *ClearUpdateRequest) (*ClearUpdateResponse, error)
	// GetBrowserVersion returns version info of the given browser type.
	// If multiple Lacros browsers are provisioned in the stateful partition,
	// all the versions will be returned.
	GetBrowserVersion(context.Context, *GetBrowserVersionRequest) (*GetBrowserVersionResponse, error)
}

// UnimplementedUpdateTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateTestServiceServer struct {
}

func (*UnimplementedUpdateTestServiceServer) VerifyUpdate(context.Context, *VerifyUpdateRequest) (*VerifyUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUpdate not implemented")
}
func (*UnimplementedUpdateTestServiceServer) ClearUpdate(context.Context, *ClearUpdateRequest) (*ClearUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearUpdate not implemented")
}
func (*UnimplementedUpdateTestServiceServer) GetBrowserVersion(context.Context, *GetBrowserVersionRequest) (*GetBrowserVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrowserVersion not implemented")
}

func RegisterUpdateTestServiceServer(s *grpc.Server, srv UpdateTestServiceServer) {
	s.RegisterService(&_UpdateTestService_serviceDesc, srv)
}

func _UpdateTestService_VerifyUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateTestServiceServer).VerifyUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.lacros.UpdateTestService/VerifyUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateTestServiceServer).VerifyUpdate(ctx, req.(*VerifyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateTestService_ClearUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateTestServiceServer).ClearUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.lacros.UpdateTestService/ClearUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateTestServiceServer).ClearUpdate(ctx, req.(*ClearUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateTestService_GetBrowserVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrowserVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateTestServiceServer).GetBrowserVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.lacros.UpdateTestService/GetBrowserVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateTestServiceServer).GetBrowserVersion(ctx, req.(*GetBrowserVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpdateTestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.lacros.UpdateTestService",
	HandlerType: (*UpdateTestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyUpdate",
			Handler:    _UpdateTestService_VerifyUpdate_Handler,
		},
		{
			MethodName: "ClearUpdate",
			Handler:    _UpdateTestService_ClearUpdate_Handler,
		},
		{
			MethodName: "GetBrowserVersion",
			Handler:    _UpdateTestService_GetBrowserVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update_test_service.proto",
}
