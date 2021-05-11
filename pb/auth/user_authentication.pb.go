// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/user_authentication.proto

package auth

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserEmailSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Pin   string `protobuf:"bytes,2,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *UserEmailSignInRequest) Reset() {
	*x = UserEmailSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEmailSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEmailSignInRequest) ProtoMessage() {}

func (x *UserEmailSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEmailSignInRequest.ProtoReflect.Descriptor instead.
func (*UserEmailSignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *UserEmailSignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserEmailSignInRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

type UserPhoneSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Pin   string `protobuf:"bytes,2,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *UserPhoneSignInRequest) Reset() {
	*x = UserPhoneSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPhoneSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPhoneSignInRequest) ProtoMessage() {}

func (x *UserPhoneSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPhoneSignInRequest.ProtoReflect.Descriptor instead.
func (*UserPhoneSignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *UserPhoneSignInRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserPhoneSignInRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

type UserSignInReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserSignInReply) Reset() {
	*x = UserSignInReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInReply) ProtoMessage() {}

func (x *UserSignInReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInReply.ProtoReflect.Descriptor instead.
func (*UserSignInReply) Descriptor() ([]byte, []int) {
	return file_proto_user_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *UserSignInReply) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *UserSignInReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserTokenValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserTokenValidateRequest) Reset() {
	*x = UserTokenValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenValidateRequest) ProtoMessage() {}

func (x *UserTokenValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenValidateRequest.ProtoReflect.Descriptor instead.
func (*UserTokenValidateRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *UserTokenValidateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserTokenValidateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
}

func (x *UserTokenValidateReply) Reset() {
	*x = UserTokenValidateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_authentication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenValidateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenValidateReply) ProtoMessage() {}

func (x *UserTokenValidateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_authentication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenValidateReply.ProtoReflect.Descriptor instead.
func (*UserTokenValidateReply) Descriptor() ([]byte, []int) {
	return file_proto_user_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *UserTokenValidateReply) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

var File_proto_user_authentication_proto protoreflect.FileDescriptor

var file_proto_user_authentication_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x40, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x69, 0x6e, 0x22, 0x40, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x69, 0x6e, 0x22, 0x40, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x30, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x16, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x32, 0xda, 0x01, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x12, 0x17, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x17, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_authentication_proto_rawDescOnce sync.Once
	file_proto_user_authentication_proto_rawDescData = file_proto_user_authentication_proto_rawDesc
)

func file_proto_user_authentication_proto_rawDescGZIP() []byte {
	file_proto_user_authentication_proto_rawDescOnce.Do(func() {
		file_proto_user_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_authentication_proto_rawDescData)
	})
	return file_proto_user_authentication_proto_rawDescData
}

var file_proto_user_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_user_authentication_proto_goTypes = []interface{}{
	(*UserEmailSignInRequest)(nil),   // 0: UserEmailSignInRequest
	(*UserPhoneSignInRequest)(nil),   // 1: UserPhoneSignInRequest
	(*UserSignInReply)(nil),          // 2: UserSignInReply
	(*UserTokenValidateRequest)(nil), // 3: UserTokenValidateRequest
	(*UserTokenValidateReply)(nil),   // 4: UserTokenValidateReply
}
var file_proto_user_authentication_proto_depIdxs = []int32{
	3, // 0: AuthenticationService.AuthenticateToken:input_type -> UserTokenValidateRequest
	0, // 1: AuthenticationService.EmailSignIn:input_type -> UserEmailSignInRequest
	1, // 2: AuthenticationService.PhoneSignIn:input_type -> UserPhoneSignInRequest
	4, // 3: AuthenticationService.AuthenticateToken:output_type -> UserTokenValidateReply
	2, // 4: AuthenticationService.EmailSignIn:output_type -> UserSignInReply
	2, // 5: AuthenticationService.PhoneSignIn:output_type -> UserSignInReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_authentication_proto_init() }
func file_proto_user_authentication_proto_init() {
	if File_proto_user_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEmailSignInRequest); i {
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
		file_proto_user_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPhoneSignInRequest); i {
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
		file_proto_user_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInReply); i {
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
		file_proto_user_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenValidateRequest); i {
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
		file_proto_user_authentication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenValidateReply); i {
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
			RawDescriptor: file_proto_user_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_authentication_proto_goTypes,
		DependencyIndexes: file_proto_user_authentication_proto_depIdxs,
		MessageInfos:      file_proto_user_authentication_proto_msgTypes,
	}.Build()
	File_proto_user_authentication_proto = out.File
	file_proto_user_authentication_proto_rawDesc = nil
	file_proto_user_authentication_proto_goTypes = nil
	file_proto_user_authentication_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	AuthenticateToken(ctx context.Context, in *UserTokenValidateRequest, opts ...grpc.CallOption) (*UserTokenValidateReply, error)
	EmailSignIn(ctx context.Context, in *UserEmailSignInRequest, opts ...grpc.CallOption) (*UserSignInReply, error)
	PhoneSignIn(ctx context.Context, in *UserPhoneSignInRequest, opts ...grpc.CallOption) (*UserSignInReply, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) AuthenticateToken(ctx context.Context, in *UserTokenValidateRequest, opts ...grpc.CallOption) (*UserTokenValidateReply, error) {
	out := new(UserTokenValidateReply)
	err := c.cc.Invoke(ctx, "/AuthenticationService/AuthenticateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) EmailSignIn(ctx context.Context, in *UserEmailSignInRequest, opts ...grpc.CallOption) (*UserSignInReply, error) {
	out := new(UserSignInReply)
	err := c.cc.Invoke(ctx, "/AuthenticationService/EmailSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) PhoneSignIn(ctx context.Context, in *UserPhoneSignInRequest, opts ...grpc.CallOption) (*UserSignInReply, error) {
	out := new(UserSignInReply)
	err := c.cc.Invoke(ctx, "/AuthenticationService/PhoneSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	AuthenticateToken(context.Context, *UserTokenValidateRequest) (*UserTokenValidateReply, error)
	EmailSignIn(context.Context, *UserEmailSignInRequest) (*UserSignInReply, error)
	PhoneSignIn(context.Context, *UserPhoneSignInRequest) (*UserSignInReply, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) AuthenticateToken(context.Context, *UserTokenValidateRequest) (*UserTokenValidateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateToken not implemented")
}
func (*UnimplementedAuthenticationServiceServer) EmailSignIn(context.Context, *UserEmailSignInRequest) (*UserSignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailSignIn not implemented")
}
func (*UnimplementedAuthenticationServiceServer) PhoneSignIn(context.Context, *UserPhoneSignInRequest) (*UserSignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneSignIn not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_AuthenticateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).AuthenticateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthenticationService/AuthenticateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).AuthenticateToken(ctx, req.(*UserTokenValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_EmailSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmailSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).EmailSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthenticationService/EmailSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).EmailSignIn(ctx, req.(*UserEmailSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_PhoneSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPhoneSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).PhoneSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthenticationService/PhoneSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).PhoneSignIn(ctx, req.(*UserPhoneSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateToken",
			Handler:    _AuthenticationService_AuthenticateToken_Handler,
		},
		{
			MethodName: "EmailSignIn",
			Handler:    _AuthenticationService_EmailSignIn_Handler,
		},
		{
			MethodName: "PhoneSignIn",
			Handler:    _AuthenticationService_PhoneSignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user_authentication.proto",
}
