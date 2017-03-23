// Code generated by protoc-gen-go.
// source: kfe/config/grpc/backends/backend.proto
// DO NOT EDIT!

/*
Package kfe_config_grpc_backends is a generated protocol buffer package.

It is generated from these files:
	kfe/config/grpc/backends/backend.proto

It has these top-level messages:
	Backend
	Interceptor
	Security
*/
package kfe_config_grpc_backends

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import  kfe_config_common_resolvers "github.com/mwitkow/kfe/_protogen/kfe/config/common/resolvers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / Balancer chooses which gRPC balancing policy to use.
type Balancer int32

const (
	// ROUND_ROBIN is the simpliest and default load balancing policy
	Balancer_ROUND_ROBIN Balancer = 0
)

var Balancer_name = map[int32]string{
	0: "ROUND_ROBIN",
}
var Balancer_value = map[string]int32{
	"ROUND_ROBIN": 0,
}

func (x Balancer) String() string {
	return proto.EnumName(Balancer_name, int32(x))
}
func (Balancer) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// / Backend is a gRPC ClientConn pool maintained to a single serivce.
type Backend struct {
	// / name is the string identifying the bakcend in all other conifgs.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// / balancer decides which balancing policy to use.
	Balancer Balancer `protobuf:"varint,2,opt,name=balancer,enum=kfe.config.grpc.backends.Balancer" json:"balancer,omitempty"`
	// / disable_conntracking turns off the /debug/events tracing and Prometheus monitoring of the pool sie for this backend.
	DisableConntracking bool `protobuf:"varint,3,opt,name=disable_conntracking,json=disableConntracking" json:"disable_conntracking,omitempty"`
	// / security controls the TLS connection details for the backend. If not present, Insecure (plain text) mode is used.
	Security *Security `protobuf:"bytes,4,opt,name=security" json:"security,omitempty"`
	// / interceptors controls what interceptors will be enabled for this backend.
	Interceptors []*Interceptor `protobuf:"bytes,5,rep,name=interceptors" json:"interceptors,omitempty"`
	// Types that are valid to be assigned to Resolver:
	//	*Backend_Srv
	//	*Backend_K8S
	Resolver isBackend_Resolver `protobuf_oneof:"resolver"`
}

func (m *Backend) Reset()                    { *m = Backend{} }
func (m *Backend) String() string            { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()               {}
func (*Backend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBackend_Resolver interface {
	isBackend_Resolver()
}

type Backend_Srv struct {
	Srv *kfe_config_common_resolvers.SrvResolver `protobuf:"bytes,10,opt,name=srv,oneof"`
}
type Backend_K8S struct {
	K8S *kfe_config_common_resolvers.KubeResolver `protobuf:"bytes,11,opt,name=k8s,oneof"`
}

func (*Backend_Srv) isBackend_Resolver() {}
func (*Backend_K8S) isBackend_Resolver() {}

func (m *Backend) GetResolver() isBackend_Resolver {
	if m != nil {
		return m.Resolver
	}
	return nil
}

func (m *Backend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Backend) GetBalancer() Balancer {
	if m != nil {
		return m.Balancer
	}
	return Balancer_ROUND_ROBIN
}

func (m *Backend) GetDisableConntracking() bool {
	if m != nil {
		return m.DisableConntracking
	}
	return false
}

func (m *Backend) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *Backend) GetInterceptors() []*Interceptor {
	if m != nil {
		return m.Interceptors
	}
	return nil
}

func (m *Backend) GetSrv() *kfe_config_common_resolvers.SrvResolver {
	if x, ok := m.GetResolver().(*Backend_Srv); ok {
		return x.Srv
	}
	return nil
}

func (m *Backend) GetK8S() *kfe_config_common_resolvers.KubeResolver {
	if x, ok := m.GetResolver().(*Backend_K8S); ok {
		return x.K8S
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Backend) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Backend_OneofMarshaler, _Backend_OneofUnmarshaler, _Backend_OneofSizer, []interface{}{
		(*Backend_Srv)(nil),
		(*Backend_K8S)(nil),
	}
}

func _Backend_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Backend)
	// resolver
	switch x := m.Resolver.(type) {
	case *Backend_Srv:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Srv); err != nil {
			return err
		}
	case *Backend_K8S:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.K8S); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Backend.Resolver has unexpected type %T", x)
	}
	return nil
}

func _Backend_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Backend)
	switch tag {
	case 10: // resolver.srv
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kfe_config_common_resolvers.SrvResolver)
		err := b.DecodeMessage(msg)
		m.Resolver = &Backend_Srv{msg}
		return true, err
	case 11: // resolver.k8s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kfe_config_common_resolvers.KubeResolver)
		err := b.DecodeMessage(msg)
		m.Resolver = &Backend_K8S{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Backend_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Backend)
	// resolver
	switch x := m.Resolver.(type) {
	case *Backend_Srv:
		s := proto.Size(x.Srv)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Backend_K8S:
		s := proto.Size(x.K8S)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Interceptor struct {
	// Types that are valid to be assigned to Interceptor:
	//	*Interceptor_Prometheus
	Interceptor isInterceptor_Interceptor `protobuf_oneof:"interceptor"`
}

func (m *Interceptor) Reset()                    { *m = Interceptor{} }
func (m *Interceptor) String() string            { return proto.CompactTextString(m) }
func (*Interceptor) ProtoMessage()               {}
func (*Interceptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isInterceptor_Interceptor interface {
	isInterceptor_Interceptor()
}

type Interceptor_Prometheus struct {
	Prometheus bool `protobuf:"varint,1,opt,name=prometheus,oneof"`
}

func (*Interceptor_Prometheus) isInterceptor_Interceptor() {}

func (m *Interceptor) GetInterceptor() isInterceptor_Interceptor {
	if m != nil {
		return m.Interceptor
	}
	return nil
}

func (m *Interceptor) GetPrometheus() bool {
	if x, ok := m.GetInterceptor().(*Interceptor_Prometheus); ok {
		return x.Prometheus
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Interceptor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Interceptor_OneofMarshaler, _Interceptor_OneofUnmarshaler, _Interceptor_OneofSizer, []interface{}{
		(*Interceptor_Prometheus)(nil),
	}
}

func _Interceptor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Interceptor)
	// interceptor
	switch x := m.Interceptor.(type) {
	case *Interceptor_Prometheus:
		t := uint64(0)
		if x.Prometheus {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Interceptor.Interceptor has unexpected type %T", x)
	}
	return nil
}

func _Interceptor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Interceptor)
	switch tag {
	case 1: // interceptor.prometheus
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Interceptor = &Interceptor_Prometheus{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Interceptor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Interceptor)
	// interceptor
	switch x := m.Interceptor.(type) {
	case *Interceptor_Prometheus:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / Security settings for a backend.
type Security struct {
	// / insecure_skip_verify skips the server certificate verification completely.
	// / No TLS config (for testclient or server) will be used. This should *not* be used in production software.
	InsecureSkipVerify bool `protobuf:"varint,1,opt,name=insecure_skip_verify,json=insecureSkipVerify" json:"insecure_skip_verify,omitempty"`
	// / config_name indicates the TlsServerConfig to be used for this connection.
	ConfigName string `protobuf:"bytes,2,opt,name=config_name,json=configName" json:"config_name,omitempty"`
}

func (m *Security) Reset()                    { *m = Security{} }
func (m *Security) String() string            { return proto.CompactTextString(m) }
func (*Security) ProtoMessage()               {}
func (*Security) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Security) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *Security) GetConfigName() string {
	if m != nil {
		return m.ConfigName
	}
	return ""
}

func init() {
	proto.RegisterType((*Backend)(nil), "kfe.config.grpc.backends.Backend")
	proto.RegisterType((*Interceptor)(nil), "kfe.config.grpc.backends.Interceptor")
	proto.RegisterType((*Security)(nil), "kfe.config.grpc.backends.Security")
	proto.RegisterEnum("kfe.config.grpc.backends.Balancer", Balancer_name, Balancer_value)
}

func init() { proto.RegisterFile("kfe/config/grpc/backends/backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x37, 0x4d, 0x81, 0x30, 0xe1, 0x9f, 0x4c, 0x0f, 0x16, 0x1c, 0x88, 0x22, 0x81, 0x02,
	0x48, 0x09, 0x94, 0x4b, 0x0f, 0xd0, 0x43, 0xe0, 0xd0, 0x15, 0xd2, 0x56, 0xf2, 0x0a, 0x6e, 0x28,
	0x4a, 0xdc, 0xc9, 0x62, 0x79, 0x63, 0x47, 0x76, 0x36, 0x52, 0xdf, 0x87, 0x07, 0x45, 0xeb, 0x24,
	0xdd, 0xf4, 0xd0, 0xde, 0x26, 0x9e, 0xef, 0xf7, 0xcd, 0x64, 0x66, 0xe0, 0x9d, 0xac, 0x31, 0xe3,
	0x5a, 0xd5, 0x62, 0x93, 0x6d, 0x4c, 0xcb, 0xb3, 0xaa, 0xe4, 0x12, 0xd5, 0x95, 0x9d, 0x82, 0xb4,
	0x35, 0xba, 0xd3, 0x84, 0xca, 0x1a, 0xd3, 0x41, 0x97, 0xee, 0x75, 0xe9, 0xa4, 0x7b, 0xf5, 0x71,
	0xe6, 0xc0, 0x75, 0xd3, 0x68, 0x95, 0x19, 0xb4, 0x7a, 0xdb, 0xa3, 0xb1, 0x87, 0x68, 0xb0, 0x89,
	0xff, 0xf9, 0xf0, 0x28, 0x1f, 0x48, 0x42, 0xe0, 0x58, 0x95, 0x0d, 0x52, 0x2f, 0xf2, 0x92, 0xc7,
	0xcc, 0xc5, 0xe4, 0x1c, 0x82, 0xaa, 0xdc, 0x96, 0x8a, 0xa3, 0xa1, 0x47, 0x91, 0x97, 0x3c, 0x3b,
	0x8d, 0xd3, 0xbb, 0x2a, 0xa7, 0xf9, 0xa8, 0x64, 0x37, 0x0c, 0xf9, 0x0c, 0x27, 0x57, 0xc2, 0x96,
	0xd5, 0x16, 0x0b, 0xae, 0x95, 0xea, 0x4c, 0xc9, 0xa5, 0x50, 0x1b, 0xea, 0x47, 0x5e, 0x12, 0xb0,
	0x97, 0x63, 0xee, 0xfb, 0x2c, 0xb5, 0x2f, 0x69, 0x91, 0xef, 0x8c, 0xe8, 0xae, 0xe9, 0x71, 0xe4,
	0x25, 0xe1, 0x7d, 0x25, 0xd7, 0xa3, 0x92, 0xdd, 0x30, 0x64, 0x09, 0x4f, 0x84, 0xea, 0xd0, 0x70,
	0x6c, 0x3b, 0x6d, 0x2c, 0x7d, 0x10, 0xf9, 0x49, 0x78, 0xfa, 0xf6, 0x6e, 0x8f, 0xe5, 0x41, 0xcd,
	0x6e, 0xa1, 0xe4, 0x2b, 0xf8, 0xd6, 0xf4, 0x14, 0x5c, 0x17, 0xc9, 0xdc, 0x61, 0x18, 0x6c, 0x7a,
	0x18, 0xe7, 0xda, 0xf4, 0x6c, 0xfc, 0xb8, 0x58, 0xb0, 0x3d, 0x46, 0xbe, 0x81, 0x2f, 0xcf, 0x2c,
	0x0d, 0x1d, 0xfd, 0xfe, 0x5e, 0xfa, 0xe7, 0xae, 0xc2, 0x39, 0x2e, 0xcf, 0x6c, 0x0e, 0x10, 0x4c,
	0x82, 0xf8, 0x1c, 0xc2, 0x59, 0x97, 0x24, 0x02, 0x68, 0x8d, 0x6e, 0xb0, 0xfb, 0x8b, 0x3b, 0xeb,
	0xf6, 0x15, 0x5c, 0x2c, 0xd8, 0xec, 0x2d, 0x7f, 0x0a, 0xe1, 0xec, 0x4f, 0xe2, 0x3f, 0x10, 0x4c,
	0x93, 0x22, 0x9f, 0xe0, 0x44, 0x28, 0x37, 0x2d, 0x2c, 0xac, 0x14, 0x6d, 0xd1, 0xa3, 0x11, 0xf5,
	0xf5, 0x60, 0xc3, 0xc8, 0x94, 0x5b, 0x4b, 0xd1, 0xfe, 0x76, 0x19, 0xf2, 0x06, 0xc2, 0xa1, 0xf1,
	0xc2, 0xdd, 0xc7, 0x91, 0xbb, 0x0f, 0x18, 0x9e, 0x56, 0x65, 0x83, 0x1f, 0x5e, 0x43, 0x30, 0xed,
	0x9e, 0x3c, 0x87, 0x90, 0x5d, 0xfe, 0x5a, 0xfd, 0x28, 0xd8, 0x65, 0xbe, 0x5c, 0xbd, 0x58, 0x54,
	0x0f, 0xdd, 0xa5, 0x7d, 0xf9, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x50, 0xae, 0x80, 0xda, 0x02,
	0x00, 0x00,
}
