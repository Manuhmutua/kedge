// Code generated by protoc-gen-go.
// source: base/serverstatus.proto
// DO NOT EDIT!

/*
Package base is a generated protocol buffer package.

It is generated from these files:
	base/serverstatus.proto

It has these top-level messages:
	HealthCheckResponse
	FlagzState
	VersionResponse
*/
package base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type HealthCheckResponse struct {
	IsOk bool `protobuf:"varint,1,opt,name=is_ok,json=isOk" json:"is_ok,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthCheckResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

type FlagzState struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Help         string `protobuf:"bytes,2,opt,name=help" json:"help,omitempty"`
	CurrentValue string `protobuf:"bytes,3,opt,name=current_value,json=currentValue" json:"current_value,omitempty"`
	DefaultValue string `protobuf:"bytes,4,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
}

func (m *FlagzState) Reset()                    { *m = FlagzState{} }
func (m *FlagzState) String() string            { return proto.CompactTextString(m) }
func (*FlagzState) ProtoMessage()               {}
func (*FlagzState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlagzState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlagzState) GetHelp() string {
	if m != nil {
		return m.Help
	}
	return ""
}

func (m *FlagzState) GetCurrentValue() string {
	if m != nil {
		return m.CurrentValue
	}
	return ""
}

func (m *FlagzState) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type VersionResponse struct {
	Hash       string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Branchname string `protobuf:"bytes,2,opt,name=branchname" json:"branchname,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
	Go         string `protobuf:"bytes,4,opt,name=go" json:"go,omitempty"`
	Epoch      string `protobuf:"bytes,5,opt,name=epoch" json:"epoch,omitempty"`
	Tag        string `protobuf:"bytes,6,opt,name=tag" json:"tag,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VersionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *VersionResponse) GetBranchname() string {
	if m != nil {
		return m.Branchname
	}
	return ""
}

func (m *VersionResponse) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *VersionResponse) GetGo() string {
	if m != nil {
		return m.Go
	}
	return ""
}

func (m *VersionResponse) GetEpoch() string {
	if m != nil {
		return m.Epoch
	}
	return ""
}

func (m *VersionResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheckResponse)(nil), "base.HealthCheckResponse")
	proto.RegisterType((*FlagzState)(nil), "base.FlagzState")
	proto.RegisterType((*VersionResponse)(nil), "base.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServerStatus service

type ServerStatusClient interface {
	// Serves as a simple Healthcheck of the service.
	HealthCheck(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	// Lists all Flagz states for this server.
	FlagzList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (ServerStatus_FlagzListClient, error)
	// Version returns the git hash and other used when building this server.
	Version(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type serverStatusClient struct {
	cc *grpc.ClientConn
}

func NewServerStatusClient(cc *grpc.ClientConn) ServerStatusClient {
	return &serverStatusClient{cc}
}

func (c *serverStatusClient) HealthCheck(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/base.ServerStatus/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatusClient) FlagzList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (ServerStatus_FlagzListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ServerStatus_serviceDesc.Streams[0], c.cc, "/base.ServerStatus/FlagzList", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStatusFlagzListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStatus_FlagzListClient interface {
	Recv() (*FlagzState, error)
	grpc.ClientStream
}

type serverStatusFlagzListClient struct {
	grpc.ClientStream
}

func (x *serverStatusFlagzListClient) Recv() (*FlagzState, error) {
	m := new(FlagzState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serverStatusClient) Version(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/base.ServerStatus/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServerStatus service

type ServerStatusServer interface {
	// Serves as a simple Healthcheck of the service.
	HealthCheck(context.Context, *google_protobuf.Empty) (*HealthCheckResponse, error)
	// Lists all Flagz states for this server.
	FlagzList(*google_protobuf.Empty, ServerStatus_FlagzListServer) error
	// Version returns the git hash and other used when building this server.
	Version(context.Context, *google_protobuf.Empty) (*VersionResponse, error)
}

func RegisterServerStatusServer(s *grpc.Server, srv ServerStatusServer) {
	s.RegisterService(&_ServerStatus_serviceDesc, srv)
}

func _ServerStatus_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatusServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.ServerStatus/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatusServer).HealthCheck(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatus_FlagzList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStatusServer).FlagzList(m, &serverStatusFlagzListServer{stream})
}

type ServerStatus_FlagzListServer interface {
	Send(*FlagzState) error
	grpc.ServerStream
}

type serverStatusFlagzListServer struct {
	grpc.ServerStream
}

func (x *serverStatusFlagzListServer) Send(m *FlagzState) error {
	return x.ServerStream.SendMsg(m)
}

func _ServerStatus_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatusServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.ServerStatus/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatusServer).Version(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.ServerStatus",
	HandlerType: (*ServerStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ServerStatus_HealthCheck_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _ServerStatus_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FlagzList",
			Handler:       _ServerStatus_FlagzList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "base/serverstatus.proto",
}

func init() { proto.RegisterFile("base/serverstatus.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x9b, 0x36, 0xed, 0x7b, 0xbd, 0xaf, 0xef, 0xbd, 0x32, 0xf5, 0x4f, 0xac, 0x20, 0x12,
	0x37, 0xe2, 0x22, 0x15, 0x5d, 0xe9, 0x52, 0x51, 0x5c, 0x08, 0x42, 0x0b, 0xdd, 0x96, 0x49, 0x7a,
	0x9b, 0x84, 0xa6, 0x99, 0x30, 0x33, 0x29, 0xe8, 0xc2, 0xef, 0xe0, 0xc7, 0xf2, 0x5b, 0xc9, 0xdc,
	0x49, 0xb4, 0x88, 0xdd, 0x9d, 0x39, 0x73, 0xee, 0xcc, 0xe1, 0x77, 0x61, 0x3f, 0xe4, 0x0a, 0x47,
	0x0a, 0xe5, 0x1a, 0xa5, 0xd2, 0x5c, 0x97, 0x2a, 0x28, 0xa4, 0xd0, 0x82, 0xb9, 0xe6, 0x62, 0x78,
	0x18, 0x0b, 0x11, 0x67, 0x38, 0x22, 0x2f, 0x2c, 0x17, 0x23, 0x5c, 0x15, 0xfa, 0xd9, 0x46, 0xfc,
	0x33, 0x18, 0x3c, 0x20, 0xcf, 0x74, 0x72, 0x9b, 0x60, 0xb4, 0x1c, 0xa3, 0x2a, 0x44, 0xae, 0x90,
	0x0d, 0xa0, 0x9d, 0xaa, 0x99, 0x58, 0x7a, 0xce, 0xb1, 0x73, 0xfa, 0x7b, 0xec, 0xa6, 0xea, 0x69,
	0xe9, 0xbf, 0x02, 0xdc, 0x67, 0x3c, 0x7e, 0x99, 0x68, 0xae, 0x91, 0x31, 0x70, 0x73, 0xbe, 0x42,
	0x4a, 0x74, 0xc7, 0xa4, 0x8d, 0x97, 0x60, 0x56, 0x78, 0x4d, 0xeb, 0x19, 0xcd, 0x4e, 0xe0, 0x6f,
	0x54, 0x4a, 0x89, 0xb9, 0x9e, 0xad, 0x79, 0x56, 0xa2, 0xd7, 0xa2, 0xcb, 0x5e, 0x65, 0x4e, 0x8d,
	0x67, 0x42, 0x73, 0x5c, 0xf0, 0x32, 0xab, 0x43, 0xae, 0x0d, 0x55, 0x26, 0x85, 0xfc, 0x37, 0x07,
	0xfe, 0x4f, 0x51, 0xaa, 0x54, 0xe4, 0x9f, 0x45, 0xcd, 0x8f, 0x5c, 0x25, 0x75, 0x0b, 0xa3, 0xd9,
	0x11, 0x40, 0x28, 0x79, 0x1e, 0x25, 0xd4, 0xcf, 0x76, 0xd9, 0x70, 0xcc, 0xcc, 0x9c, 0xeb, 0xba,
	0x08, 0x69, 0xf6, 0x0f, 0x9a, 0xb1, 0xa8, 0x7e, 0x6d, 0xc6, 0x82, 0xed, 0x40, 0x1b, 0x0b, 0x11,
	0x25, 0x5e, 0x9b, 0x2c, 0x7b, 0x60, 0x7d, 0x68, 0x69, 0x1e, 0x7b, 0x1d, 0xf2, 0x8c, 0xbc, 0x78,
	0x77, 0xa0, 0x37, 0x21, 0xf2, 0x13, 0x22, 0xcf, 0x6e, 0xe0, 0xcf, 0x06, 0x50, 0xb6, 0x17, 0x58,
	0xfa, 0x41, 0x4d, 0x3f, 0xb8, 0x33, 0xf4, 0x87, 0x07, 0x81, 0xd9, 0x4d, 0xf0, 0x03, 0x7b, 0xbf,
	0xc1, 0xae, 0xa0, 0x4b, 0xa0, 0x1f, 0x53, 0xa5, 0xb7, 0xbe, 0xd0, 0xb7, 0x2f, 0x7c, 0x6d, 0xc4,
	0x6f, 0x9c, 0x3b, 0xec, 0x1a, 0x7e, 0x55, 0x88, 0xb6, 0x0e, 0xee, 0xda, 0xc1, 0x6f, 0x24, 0xfd,
	0x46, 0xd8, 0xa1, 0xe0, 0xe5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x5f, 0x89, 0xe8, 0x50,
	0x02, 0x00, 0x00,
}
