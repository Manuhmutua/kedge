// Code generated by protoc-gen-go.
// source: kfe/config/director.proto
// DO NOT EDIT!

package kfe_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import  kfe_config_grpc_routes "github.com/mwitkow/kfe/_protogen/kfe/config/grpc/routes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// / DirectorConfig is the top level configuration message the director.
type DirectorConfig struct {
	Grpc *DirectorConfig_Grpc `protobuf:"bytes,1,opt,name=grpc" json:"grpc,omitempty"`
}

func (m *DirectorConfig) Reset()                    { *m = DirectorConfig{} }
func (m *DirectorConfig) String() string            { return proto.CompactTextString(m) }
func (*DirectorConfig) ProtoMessage()               {}
func (*DirectorConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DirectorConfig) GetGrpc() *DirectorConfig_Grpc {
	if m != nil {
		return m.Grpc
	}
	return nil
}

type DirectorConfig_Grpc struct {
	Routes []*kfe_config_grpc_routes.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *DirectorConfig_Grpc) Reset()                    { *m = DirectorConfig_Grpc{} }
func (m *DirectorConfig_Grpc) String() string            { return proto.CompactTextString(m) }
func (*DirectorConfig_Grpc) ProtoMessage()               {}
func (*DirectorConfig_Grpc) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *DirectorConfig_Grpc) GetRoutes() []*kfe_config_grpc_routes.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*DirectorConfig)(nil), "kfe.config.DirectorConfig")
	proto.RegisterType((*DirectorConfig_Grpc)(nil), "kfe.config.DirectorConfig.Grpc")
}

func init() { proto.RegisterFile("kfe/config/director.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4e, 0x4b, 0xd5,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xc9, 0x2c, 0x4a, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x4e, 0x4b, 0xd5, 0x83, 0x48, 0x49, 0x29, 0x23,
	0x29, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0x86, 0x52, 0x10, 0x0d,
	0x4a, 0x2d, 0x8c, 0x5c, 0x7c, 0x2e, 0x50, 0x33, 0x9c, 0xc1, 0x6a, 0x85, 0x8c, 0xb9, 0x58, 0x40,
	0xca, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xe4, 0xf5, 0x10, 0x46, 0xea, 0xa1, 0xaa, 0xd4,
	0x73, 0x2f, 0x2a, 0x48, 0x0e, 0x02, 0x2b, 0x96, 0xb2, 0xe5, 0x62, 0x01, 0xf1, 0x84, 0x4c, 0xb9,
	0xd8, 0x20, 0xe6, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0x22, 0x6b, 0x07, 0xa9, 0xd4,
	0x83, 0x5a, 0x1f, 0x04, 0xa2, 0x82, 0xa0, 0x8a, 0x93, 0xd8, 0xc0, 0xae, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x68, 0xf1, 0xf7, 0xdb, 0x00, 0x00, 0x00,
}
