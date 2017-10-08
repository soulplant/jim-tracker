// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	User
	Talk
	FetchAllRequest
	FetchAllResponse
	AddUserRequest
	AddUserResponse
	AddTalkRequest
	AddTalkResponse
	ReorderRequest
	ReorderResponse
	GetUsersRequest
	GetUsersResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NextTalk string `protobuf:"bytes,3,opt,name=nextTalk" json:"nextTalk,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNextTalk() string {
	if m != nil {
		return m.NextTalk
	}
	return ""
}

type Talk struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SpeakerId string   `protobuf:"bytes,3,opt,name=speakerId" json:"speakerId,omitempty"`
	Done      bool     `protobuf:"varint,4,opt,name=done" json:"done,omitempty"`
	Url       []string `protobuf:"bytes,5,rep,name=url" json:"url,omitempty"`
}

func (m *Talk) Reset()                    { *m = Talk{} }
func (m *Talk) String() string            { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()               {}
func (*Talk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Talk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Talk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Talk) GetSpeakerId() string {
	if m != nil {
		return m.SpeakerId
	}
	return ""
}

func (m *Talk) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Talk) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

type FetchAllRequest struct {
}

func (m *FetchAllRequest) Reset()                    { *m = FetchAllRequest{} }
func (m *FetchAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchAllRequest) ProtoMessage()               {}
func (*FetchAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type FetchAllResponse struct {
	Version string  `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	User    []*User `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
	Talk    []*Talk `protobuf:"bytes,3,rep,name=talk" json:"talk,omitempty"`
}

func (m *FetchAllResponse) Reset()                    { *m = FetchAllResponse{} }
func (m *FetchAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchAllResponse) ProtoMessage()               {}
func (*FetchAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchAllResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *FetchAllResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *FetchAllResponse) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type AddUserRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddUserResponse struct {
	// The newly added user.
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *AddUserResponse) Reset()                    { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()               {}
func (*AddUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type AddTalkRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *AddTalkRequest) Reset()                    { *m = AddTalkRequest{} }
func (m *AddTalkRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTalkRequest) ProtoMessage()               {}
func (*AddTalkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddTalkRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type AddTalkResponse struct {
	Talk *Talk `protobuf:"bytes,1,opt,name=talk" json:"talk,omitempty"`
}

func (m *AddTalkResponse) Reset()                    { *m = AddTalkResponse{} }
func (m *AddTalkResponse) String() string            { return proto.CompactTextString(m) }
func (*AddTalkResponse) ProtoMessage()               {}
func (*AddTalkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddTalkResponse) GetTalk() *Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type ReorderRequest struct {
	Version      int64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	MoveUserId   string `protobuf:"bytes,2,opt,name=moveUserId" json:"moveUserId,omitempty"`
	AnchorUserId string `protobuf:"bytes,3,opt,name=anchorUserId" json:"anchorUserId,omitempty"`
	Before       bool   `protobuf:"varint,4,opt,name=before" json:"before,omitempty"`
}

func (m *ReorderRequest) Reset()                    { *m = ReorderRequest{} }
func (m *ReorderRequest) String() string            { return proto.CompactTextString(m) }
func (*ReorderRequest) ProtoMessage()               {}
func (*ReorderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReorderRequest) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReorderRequest) GetMoveUserId() string {
	if m != nil {
		return m.MoveUserId
	}
	return ""
}

func (m *ReorderRequest) GetAnchorUserId() string {
	if m != nil {
		return m.AnchorUserId
	}
	return ""
}

func (m *ReorderRequest) GetBefore() bool {
	if m != nil {
		return m.Before
	}
	return false
}

type ReorderResponse struct {
	Accepted bool   `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *ReorderResponse) Reset()                    { *m = ReorderResponse{} }
func (m *ReorderResponse) String() string            { return proto.CompactTextString(m) }
func (*ReorderResponse) ProtoMessage()               {}
func (*ReorderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReorderResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *ReorderResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetUsersRequest struct {
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type GetUsersResponse struct {
	Version int64   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	User    []*User `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
}

func (m *GetUsersResponse) Reset()                    { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()               {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetUsersResponse) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetUsersResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*Talk)(nil), "api.Talk")
	proto.RegisterType((*FetchAllRequest)(nil), "api.FetchAllRequest")
	proto.RegisterType((*FetchAllResponse)(nil), "api.FetchAllResponse")
	proto.RegisterType((*AddUserRequest)(nil), "api.AddUserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "api.AddUserResponse")
	proto.RegisterType((*AddTalkRequest)(nil), "api.AddTalkRequest")
	proto.RegisterType((*AddTalkResponse)(nil), "api.AddTalkResponse")
	proto.RegisterType((*ReorderRequest)(nil), "api.ReorderRequest")
	proto.RegisterType((*ReorderResponse)(nil), "api.ReorderResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "api.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "api.GetUsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApiService service

type ApiServiceClient interface {
	FetchAll(ctx context.Context, in *FetchAllRequest, opts ...grpc.CallOption) (*FetchAllResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	// Change the position of one user in the list of upcoming talks.
	Reorder(ctx context.Context, in *ReorderRequest, opts ...grpc.CallOption) (*ReorderResponse, error)
	AddTalk(ctx context.Context, in *AddTalkRequest, opts ...grpc.CallOption) (*AddTalkResponse, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) FetchAll(ctx context.Context, in *FetchAllRequest, opts ...grpc.CallOption) (*FetchAllResponse, error) {
	out := new(FetchAllResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/FetchAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) Reorder(ctx context.Context, in *ReorderRequest, opts ...grpc.CallOption) (*ReorderResponse, error) {
	out := new(ReorderResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/Reorder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddTalk(ctx context.Context, in *AddTalkRequest, opts ...grpc.CallOption) (*AddTalkResponse, error) {
	out := new(AddTalkResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/AddTalk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiService service

type ApiServiceServer interface {
	FetchAll(context.Context, *FetchAllRequest) (*FetchAllResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	// Change the position of one user in the list of upcoming talks.
	Reorder(context.Context, *ReorderRequest) (*ReorderResponse, error)
	AddTalk(context.Context, *AddTalkRequest) (*AddTalkResponse, error)
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_FetchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).FetchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/FetchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).FetchAll(ctx, req.(*FetchAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_Reorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Reorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/Reorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Reorder(ctx, req.(*ReorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AddTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AddTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/AddTalk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AddTalk(ctx, req.(*AddTalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAll",
			Handler:    _ApiService_FetchAll_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _ApiService_GetUsers_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ApiService_AddUser_Handler,
		},
		{
			MethodName: "Reorder",
			Handler:    _ApiService_Reorder_Handler,
		},
		{
			MethodName: "AddTalk",
			Handler:    _ApiService_AddTalk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xed, 0xd0, 0x38, 0x53, 0x9a, 0x84, 0xed, 0x8f, 0x2c, 0xab, 0xa0, 0x68, 0xc5, 0xc1,
	0xaa, 0x44, 0x03, 0xe5, 0xd6, 0x5b, 0x2e, 0xad, 0x22, 0x10, 0x07, 0x43, 0x1f, 0x60, 0x6b, 0x4f,
	0x5b, 0x53, 0x77, 0xd7, 0xac, 0x9d, 0x88, 0x33, 0x07, 0x2e, 0x1c, 0x79, 0x34, 0x5e, 0x81, 0x07,
	0x41, 0xfb, 0xe3, 0x5f, 0x7a, 0xe8, 0x29, 0x3b, 0xdf, 0x4c, 0xbe, 0xf9, 0x66, 0xe6, 0x93, 0x61,
	0xc2, 0x8a, 0xec, 0xb4, 0x90, 0xa2, 0x12, 0xc4, 0x63, 0x45, 0x16, 0x1e, 0xdf, 0x0a, 0x71, 0x9b,
	0xe3, 0x92, 0x15, 0xd9, 0x92, 0x71, 0x2e, 0x2a, 0x56, 0x65, 0x82, 0x97, 0xa6, 0x84, 0x5e, 0xc0,
	0xe8, 0xaa, 0x44, 0x49, 0xa6, 0xe0, 0x66, 0x69, 0xe0, 0x2c, 0x9c, 0x68, 0x12, 0xbb, 0x59, 0x4a,
	0x08, 0x8c, 0x38, 0x7b, 0xc0, 0xc0, 0xd5, 0x88, 0x7e, 0x93, 0x10, 0x7c, 0x8e, 0xdf, 0xab, 0x2f,
	0x2c, 0xbf, 0x0f, 0x3c, 0x8d, 0x37, 0x31, 0xe5, 0x30, 0x52, 0xbf, 0x4f, 0xe2, 0x39, 0x86, 0x49,
	0x59, 0x20, 0xbb, 0x47, 0xb9, 0x4e, 0x2d, 0x51, 0x0b, 0xa8, 0x7f, 0xa4, 0x82, 0x63, 0x30, 0x5a,
	0x38, 0x91, 0x1f, 0xeb, 0x37, 0x99, 0x83, 0xb7, 0x91, 0x79, 0xf0, 0x6c, 0xe1, 0x45, 0x93, 0x58,
	0x3d, 0xe9, 0x0b, 0x98, 0x5d, 0x60, 0x95, 0xdc, 0xad, 0xf2, 0x3c, 0xc6, 0x6f, 0x1b, 0x2c, 0x2b,
	0xfa, 0x15, 0xe6, 0x2d, 0x54, 0x16, 0x82, 0x97, 0x48, 0x02, 0x18, 0x6f, 0x51, 0x96, 0x99, 0xe0,
	0x56, 0x53, 0x1d, 0x92, 0x97, 0x30, 0xda, 0x94, 0x28, 0x03, 0x77, 0xe1, 0x45, 0xbb, 0x67, 0x93,
	0x53, 0xb5, 0x35, 0xb5, 0x89, 0x58, 0xc3, 0x2a, 0x5d, 0x99, 0x39, 0xdb, 0xb4, 0x1a, 0x30, 0xd6,
	0x30, 0x7d, 0x0d, 0xd3, 0x55, 0x9a, 0xea, 0x7a, 0xd3, 0xbd, 0x19, 0xd4, 0x69, 0x07, 0xa5, 0x6f,
	0x61, 0xd6, 0x54, 0x59, 0x41, 0x75, 0x5b, 0x55, 0xf6, 0x7f, 0x5b, 0x1a, 0x69, 0x5e, 0xdd, 0xc8,
	0xf2, 0x1e, 0xc1, 0x8e, 0xca, 0xac, 0xeb, 0xa5, 0xda, 0xc8, 0x72, 0x9b, 0xca, 0x96, 0x5b, 0x6b,
	0xee, 0x72, 0x77, 0x34, 0xff, 0x74, 0x60, 0x1a, 0xa3, 0x90, 0x69, 0x2b, 0x7a, 0xb0, 0x1e, 0xaf,
	0x5d, 0xcf, 0x2b, 0x80, 0x07, 0xb1, 0xc5, 0x2b, 0xd3, 0xda, 0x5c, 0xaf, 0x83, 0x10, 0x0a, 0xcf,
	0x19, 0x4f, 0xee, 0x84, 0xb4, 0x15, 0xe6, 0x8c, 0x3d, 0x4c, 0x49, 0xbf, 0xc6, 0x1b, 0x21, 0xeb,
	0x5b, 0xda, 0x88, 0x5e, 0xc2, 0xac, 0xd1, 0x61, 0xa5, 0x87, 0xe0, 0xb3, 0x24, 0xc1, 0xa2, 0x42,
	0x33, 0xa7, 0x1f, 0x37, 0x71, 0x57, 0xa4, 0xdb, 0xbb, 0xa1, 0x32, 0xc1, 0x25, 0x56, 0xaa, 0x5b,
	0x59, 0x9b, 0xe0, 0x03, 0xcc, 0x5b, 0xe8, 0x71, 0x13, 0x78, 0x4f, 0x35, 0xc1, 0xd9, 0x2f, 0x0f,
	0x60, 0x55, 0x64, 0x9f, 0x51, 0x6e, 0xb3, 0x04, 0xc9, 0x27, 0xf0, 0x6b, 0x83, 0x91, 0x03, 0x5d,
	0x3b, 0xb0, 0x60, 0x78, 0x38, 0x40, 0x8d, 0x00, 0x7a, 0xf8, 0xe3, 0xcf, 0xdf, 0xdf, 0xee, 0x8c,
	0xec, 0x2d, 0xb7, 0xef, 0x96, 0x37, 0x2a, 0xfb, 0x86, 0xe5, 0x39, 0x59, 0x83, 0x5f, 0x6b, 0xb5,
	0x7c, 0x83, 0x69, 0x2c, 0xdf, 0x70, 0x20, 0x3a, 0xd7, 0x7c, 0x40, 0x7c, 0xc5, 0xa7, 0xed, 0xba,
	0x86, 0xb1, 0x75, 0x1a, 0xd9, 0xd7, 0xff, 0xe9, 0xbb, 0x33, 0x3c, 0xe8, 0x83, 0x96, 0x67, 0x5f,
	0xf3, 0xec, 0xd1, 0x86, 0xe7, 0xdc, 0x39, 0x21, 0x1f, 0x61, 0x6c, 0xaf, 0x63, 0xa9, 0xfa, 0x9e,
	0xb1, 0x54, 0x83, 0x03, 0xd2, 0x23, 0x4d, 0x35, 0xa7, 0xbb, 0x8a, 0x4a, 0x9a, 0xa4, 0x62, 0x33,
	0xc2, 0xf4, 0xa7, 0xa1, 0x11, 0xd6, 0xb1, 0x77, 0x2b, 0xac, 0xeb, 0xe4, 0xbe, 0x30, 0x65, 0xde,
	0x73, 0xe7, 0xe4, 0x7a, 0x47, 0x7f, 0xb1, 0xde, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x22,
	0x8c, 0xf9, 0xe1, 0x04, 0x00, 0x00,
}