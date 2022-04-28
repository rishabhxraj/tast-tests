// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: chrome_service.proto

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

// LoginMode describes the user mode for the login.
type LoginMode int32

const (
	LoginMode_LOGIN_MODE_UNSPECIFIED LoginMode = 0
	LoginMode_LOGIN_MODE_NO_LOGIN    LoginMode = 1 // restart Chrome but don't log in
	LoginMode_LOGIN_MODE_FAKE_LOGIN  LoginMode = 2 // fake login with no authentication
	LoginMode_LOGIN_MODE_GAIA_LOGIN  LoginMode = 3 // real network-based login using GAIA backend
	LoginMode_LOGIN_MODE_GUEST_LOGIN LoginMode = 4 // sign in as ephemeral guest user
)

// Enum value maps for LoginMode.
var (
	LoginMode_name = map[int32]string{
		0: "LOGIN_MODE_UNSPECIFIED",
		1: "LOGIN_MODE_NO_LOGIN",
		2: "LOGIN_MODE_FAKE_LOGIN",
		3: "LOGIN_MODE_GAIA_LOGIN",
		4: "LOGIN_MODE_GUEST_LOGIN",
	}
	LoginMode_value = map[string]int32{
		"LOGIN_MODE_UNSPECIFIED": 0,
		"LOGIN_MODE_NO_LOGIN":    1,
		"LOGIN_MODE_FAKE_LOGIN":  2,
		"LOGIN_MODE_GAIA_LOGIN":  3,
		"LOGIN_MODE_GUEST_LOGIN": 4,
	}
)

func (x LoginMode) Enum() *LoginMode {
	p := new(LoginMode)
	*p = x
	return p
}

func (x LoginMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginMode) Descriptor() protoreflect.EnumDescriptor {
	return file_chrome_service_proto_enumTypes[0].Descriptor()
}

func (LoginMode) Type() protoreflect.EnumType {
	return &file_chrome_service_proto_enumTypes[0]
}

func (x LoginMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginMode.Descriptor instead.
func (LoginMode) EnumDescriptor() ([]byte, []int) {
	return file_chrome_service_proto_rawDescGZIP(), []int{0}
}

// NewRequest to login to Chrome with configurable features, arguments and GAIA
// credentials.
type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginMode LoginMode `protobuf:"varint,1,opt,name=login_mode,json=loginMode,proto3,enum=tast.cros.browser.LoginMode" json:"login_mode,omitempty"`
	// Credentials is login credentials.
	Credentials *NewRequest_Credentials `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// EnableFeatures contains extra Chrome features to enable.
	EnableFeatures []string `protobuf:"bytes,3,rep,name=enable_features,json=enableFeatures,proto3" json:"enable_features,omitempty"`
	// DisableFeatures contains extra Chrome features to disable.
	DisableFeatures []string `protobuf:"bytes,4,rep,name=disable_features,json=disableFeatures,proto3" json:"disable_features,omitempty"`
	// ExtraArgs returns extra arguments to pass to Chrome.
	ExtraArgs []string `protobuf:"bytes,5,rep,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty"`
	// KeepState controls whether to keep existing user profiles.
	KeepState bool `protobuf:"varint,6,opt,name=keep_state,json=keepState,proto3" json:"keep_state,omitempty"`
	// TryReuseSession controls whether to try reusing a current user session.
	TryReuseSession bool `protobuf:"varint,7,opt,name=try_reuse_session,json=tryReuseSession,proto3" json:"try_reuse_session,omitempty"`
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chrome_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chrome_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_chrome_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewRequest) GetLoginMode() LoginMode {
	if x != nil {
		return x.LoginMode
	}
	return LoginMode_LOGIN_MODE_UNSPECIFIED
}

func (x *NewRequest) GetCredentials() *NewRequest_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *NewRequest) GetEnableFeatures() []string {
	if x != nil {
		return x.EnableFeatures
	}
	return nil
}

func (x *NewRequest) GetDisableFeatures() []string {
	if x != nil {
		return x.DisableFeatures
	}
	return nil
}

func (x *NewRequest) GetExtraArgs() []string {
	if x != nil {
		return x.ExtraArgs
	}
	return nil
}

func (x *NewRequest) GetKeepState() bool {
	if x != nil {
		return x.KeepState
	}
	return false
}

func (x *NewRequest) GetTryReuseSession() bool {
	if x != nil {
		return x.TryReuseSession
	}
	return false
}

type NewRequest_Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username is the user name of a user account. It is typically an email
	// address (e.g. example@gmail.com).
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password is the password of a user account.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// GaiaId is a GAIA ID used on fake logins. If it is empty, an ID is
	// generated from the user name. The field is ignored on other type of
	// logins.
	GaiaId string `protobuf:"bytes,3,opt,name=gaia_id,json=gaiaId,proto3" json:"gaia_id,omitempty"`
	// Contact is an email address of a user who owns a test account.
	// When logging in with a test account, its contact user may be
	// notified of a login attempt and asked for approval.
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	// ParentUsername is the user name of a parent account. It is used to
	// approve a login attempt when a child account is supervised by a
	// parent account.
	ParentUsername string `protobuf:"bytes,5,opt,name=parent_username,json=parentUsername,proto3" json:"parent_username,omitempty"`
	// ParentPassword is the pass of a parent account. It is used to approve
	// a login attempt when a child account is supervised by a parent
	// account.
	ParentPassword string `protobuf:"bytes,6,opt,name=parent_password,json=parentPassword,proto3" json:"parent_password,omitempty"`
}

func (x *NewRequest_Credentials) Reset() {
	*x = NewRequest_Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chrome_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest_Credentials) ProtoMessage() {}

func (x *NewRequest_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_chrome_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest_Credentials.ProtoReflect.Descriptor instead.
func (*NewRequest_Credentials) Descriptor() ([]byte, []int) {
	return file_chrome_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NewRequest_Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NewRequest_Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NewRequest_Credentials) GetGaiaId() string {
	if x != nil {
		return x.GaiaId
	}
	return ""
}

func (x *NewRequest_Credentials) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *NewRequest_Credentials) GetParentUsername() string {
	if x != nil {
		return x.ParentUsername
	}
	return ""
}

func (x *NewRequest_Credentials) GetParentPassword() string {
	if x != nil {
		return x.ParentPassword
	}
	return ""
}

var File_chrome_service_proto protoreflect.FileDescriptor

var file_chrome_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x2e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x2e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x73, 0x2e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x75, 0x73, 0x65, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x75, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xca, 0x01,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x69, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x69, 0x61, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2a, 0x92, 0x01, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x4b, 0x45,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x41, 0x49, 0x41, 0x5f, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x04, 0x32,
	0x8a, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x2e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x74, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x75, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chrome_service_proto_rawDescOnce sync.Once
	file_chrome_service_proto_rawDescData = file_chrome_service_proto_rawDesc
)

func file_chrome_service_proto_rawDescGZIP() []byte {
	file_chrome_service_proto_rawDescOnce.Do(func() {
		file_chrome_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chrome_service_proto_rawDescData)
	})
	return file_chrome_service_proto_rawDescData
}

var file_chrome_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chrome_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chrome_service_proto_goTypes = []interface{}{
	(LoginMode)(0),                 // 0: tast.cros.browser.LoginMode
	(*NewRequest)(nil),             // 1: tast.cros.browser.NewRequest
	(*NewRequest_Credentials)(nil), // 2: tast.cros.browser.NewRequest.Credentials
	(*empty.Empty)(nil),            // 3: google.protobuf.Empty
}
var file_chrome_service_proto_depIdxs = []int32{
	0, // 0: tast.cros.browser.NewRequest.login_mode:type_name -> tast.cros.browser.LoginMode
	2, // 1: tast.cros.browser.NewRequest.credentials:type_name -> tast.cros.browser.NewRequest.Credentials
	1, // 2: tast.cros.browser.ChromeService.New:input_type -> tast.cros.browser.NewRequest
	3, // 3: tast.cros.browser.ChromeService.Close:input_type -> google.protobuf.Empty
	3, // 4: tast.cros.browser.ChromeService.New:output_type -> google.protobuf.Empty
	3, // 5: tast.cros.browser.ChromeService.Close:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chrome_service_proto_init() }
func file_chrome_service_proto_init() {
	if File_chrome_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chrome_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_chrome_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest_Credentials); i {
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
			RawDescriptor: file_chrome_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chrome_service_proto_goTypes,
		DependencyIndexes: file_chrome_service_proto_depIdxs,
		EnumInfos:         file_chrome_service_proto_enumTypes,
		MessageInfos:      file_chrome_service_proto_msgTypes,
	}.Build()
	File_chrome_service_proto = out.File
	file_chrome_service_proto_rawDesc = nil
	file_chrome_service_proto_goTypes = nil
	file_chrome_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChromeServiceClient is the client API for ChromeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChromeServiceClient interface {
	// New enables testing for Chrome and logs into a Chrome session.
	// When try_reuse_session is set to true, service tries to reuse existing
	// chrome session if the reuse criteria is met.
	// Close must be called later to clean up the associated resources.
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Close releases the chrome session obtained by New.
	// When there is no chrome session, calling Close returns an error.
	Close(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type chromeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChromeServiceClient(cc grpc.ClientConnInterface) ChromeServiceClient {
	return &chromeServiceClient{cc}
}

func (c *chromeServiceClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.browser.ChromeService/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chromeServiceClient) Close(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tast.cros.browser.ChromeService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChromeServiceServer is the server API for ChromeService service.
type ChromeServiceServer interface {
	// New enables testing for Chrome and logs into a Chrome session.
	// When try_reuse_session is set to true, service tries to reuse existing
	// chrome session if the reuse criteria is met.
	// Close must be called later to clean up the associated resources.
	New(context.Context, *NewRequest) (*empty.Empty, error)
	// Close releases the chrome session obtained by New.
	// When there is no chrome session, calling Close returns an error.
	Close(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedChromeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChromeServiceServer struct {
}

func (*UnimplementedChromeServiceServer) New(context.Context, *NewRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedChromeServiceServer) Close(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterChromeServiceServer(s *grpc.Server, srv ChromeServiceServer) {
	s.RegisterService(&_ChromeService_serviceDesc, srv)
}

func _ChromeService_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChromeServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.browser.ChromeService/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChromeServiceServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChromeService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChromeServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tast.cros.browser.ChromeService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChromeServiceServer).Close(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChromeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tast.cros.browser.ChromeService",
	HandlerType: (*ChromeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _ChromeService_New_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _ChromeService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chrome_service.proto",
}