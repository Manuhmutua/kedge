// Code generated by protoc-gen-go. DO NOT EDIT.
// source: winch/config/auth.proto

/*
Package winch_config is a generated protocol buffer package.

It is generated from these files:
	winch/config/auth.proto
	winch/config/mapper.proto

It has these top-level messages:
	AuthConfig
	AuthSource
	KubernetesAccess
	OIDCAccess
	DummyAccess
	MapperConfig
	Route
	DirectRoute
	RegexpRoute
*/
package winch_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / AuthConfig is the top level configuration message for a winch auth.
type AuthConfig struct {
	AuthSources []*AuthSource `protobuf:"bytes,1,rep,name=auth_sources,json=authSources" json:"auth_sources,omitempty"`
}

func (m *AuthConfig) Reset()                    { *m = AuthConfig{} }
func (m *AuthConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthConfig) ProtoMessage()               {}
func (*AuthConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthConfig) GetAuthSources() []*AuthSource {
	if m != nil {
		return m.AuthSources
	}
	return nil
}

// / AuthSource specifies the kind of the backend auth we need to inject on winch reqeuest.
type AuthSource struct {
	// name is an ID of auth source. It can be referenced inside winch routing.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*AuthSource_Dummy
	//	*AuthSource_Kube
	//	*AuthSource_Oidc
	Type isAuthSource_Type `protobuf_oneof:"type"`
}

func (m *AuthSource) Reset()                    { *m = AuthSource{} }
func (m *AuthSource) String() string            { return proto.CompactTextString(m) }
func (*AuthSource) ProtoMessage()               {}
func (*AuthSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isAuthSource_Type interface {
	isAuthSource_Type()
}

type AuthSource_Dummy struct {
	Dummy *DummyAccess `protobuf:"bytes,2,opt,name=dummy,oneof"`
}
type AuthSource_Kube struct {
	Kube *KubernetesAccess `protobuf:"bytes,3,opt,name=kube,oneof"`
}
type AuthSource_Oidc struct {
	Oidc *OIDCAccess `protobuf:"bytes,4,opt,name=oidc,oneof"`
}

func (*AuthSource_Dummy) isAuthSource_Type() {}
func (*AuthSource_Kube) isAuthSource_Type()  {}
func (*AuthSource_Oidc) isAuthSource_Type()  {}

func (m *AuthSource) GetType() isAuthSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AuthSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthSource) GetDummy() *DummyAccess {
	if x, ok := m.GetType().(*AuthSource_Dummy); ok {
		return x.Dummy
	}
	return nil
}

func (m *AuthSource) GetKube() *KubernetesAccess {
	if x, ok := m.GetType().(*AuthSource_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *AuthSource) GetOidc() *OIDCAccess {
	if x, ok := m.GetType().(*AuthSource_Oidc); ok {
		return x.Oidc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthSource_OneofMarshaler, _AuthSource_OneofUnmarshaler, _AuthSource_OneofSizer, []interface{}{
		(*AuthSource_Dummy)(nil),
		(*AuthSource_Kube)(nil),
		(*AuthSource_Oidc)(nil),
	}
}

func _AuthSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dummy); err != nil {
			return err
		}
	case *AuthSource_Kube:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *AuthSource_Oidc:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Oidc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthSource.Type has unexpected type %T", x)
	}
	return nil
}

func _AuthSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthSource)
	switch tag {
	case 2: // type.dummy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DummyAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Dummy{msg}
		return true, err
	case 3: // type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KubernetesAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Kube{msg}
		return true, err
	case 4: // type.oidc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OIDCAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Oidc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		s := proto.Size(x.Dummy)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Kube:
		s := proto.Size(x.Kube)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Oidc:
		s := proto.Size(x.Oidc)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / KubernetesAccess is an convenient way of specifying auth for backend. It grabs the data inside already used
// / ~/.kube/config (or any specified config path) and deducts the auth type based on that. NOTE that only these types are
// / supported:
// / - OIDC
type KubernetesAccess struct {
	// User to reference access credentials from.
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// By default ~/.kube/config as usual.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *KubernetesAccess) Reset()                    { *m = KubernetesAccess{} }
func (m *KubernetesAccess) String() string            { return proto.CompactTextString(m) }
func (*KubernetesAccess) ProtoMessage()               {}
func (*KubernetesAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KubernetesAccess) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *KubernetesAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type OIDCAccess struct {
	Provider string   `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Secret   string   `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	Scopes   []string `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	Path     string   `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
}

func (m *OIDCAccess) Reset()                    { *m = OIDCAccess{} }
func (m *OIDCAccess) String() string            { return proto.CompactTextString(m) }
func (*OIDCAccess) ProtoMessage()               {}
func (*OIDCAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OIDCAccess) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *OIDCAccess) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OIDCAccess) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OIDCAccess) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *OIDCAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// DummyAccess just directly passes specified value into auth header. If value is not specified it will return error.
type DummyAccess struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DummyAccess) Reset()                    { *m = DummyAccess{} }
func (m *DummyAccess) String() string            { return proto.CompactTextString(m) }
func (*DummyAccess) ProtoMessage()               {}
func (*DummyAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DummyAccess) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthConfig)(nil), "winch.config.AuthConfig")
	proto.RegisterType((*AuthSource)(nil), "winch.config.AuthSource")
	proto.RegisterType((*KubernetesAccess)(nil), "winch.config.KubernetesAccess")
	proto.RegisterType((*OIDCAccess)(nil), "winch.config.OIDCAccess")
	proto.RegisterType((*DummyAccess)(nil), "winch.config.DummyAccess")
}

func init() { proto.RegisterFile("winch/config/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x6e, 0x9b, 0x40,
	0x14, 0xc7, 0x8b, 0x8d, 0x51, 0x79, 0x78, 0x51, 0x8d, 0xaa, 0x96, 0x7a, 0xe1, 0x22, 0xbc, 0x61,
	0x63, 0x50, 0xdd, 0xaa, 0x9b, 0xae, 0xfc, 0xb1, 0xa8, 0x95, 0x45, 0x24, 0x72, 0x00, 0x0b, 0x86,
	0x89, 0x19, 0xd9, 0x30, 0x88, 0x99, 0xb1, 0xe5, 0xe3, 0xe4, 0x28, 0x39, 0x49, 0xa4, 0x9c, 0x24,
	0x62, 0xc6, 0x01, 0xdb, 0xd9, 0xbd, 0x8f, 0xdf, 0xfb, 0xbf, 0xf9, 0x3f, 0x80, 0xef, 0x47, 0x5a,
	0xe2, 0x3c, 0xc2, 0xac, 0x7c, 0xa4, 0xdb, 0x28, 0x91, 0x22, 0x0f, 0xab, 0x9a, 0x09, 0x86, 0x86,
	0xaa, 0x11, 0xea, 0xc6, 0xe8, 0xef, 0x96, 0x8a, 0x5c, 0xa6, 0x21, 0x66, 0x45, 0x54, 0x1c, 0xa9,
	0xd8, 0xb1, 0x63, 0xb4, 0x65, 0x53, 0x85, 0x4e, 0x0f, 0xc9, 0x9e, 0x66, 0x89, 0x60, 0x35, 0x8f,
	0xda, 0x50, 0xab, 0xf8, 0x6b, 0x80, 0xb9, 0x14, 0xf9, 0x52, 0xa9, 0xa0, 0x7f, 0x30, 0x6c, 0x36,
	0x6c, 0x38, 0x93, 0x35, 0x26, 0xdc, 0x35, 0xbc, 0x7e, 0xe0, 0xcc, 0xdc, 0xf0, 0x72, 0x55, 0xd8,
	0xf0, 0x0f, 0x0a, 0x88, 0x9d, 0xa4, 0x8d, 0xb9, 0xff, 0x6c, 0x68, 0x2d, 0x9d, 0x23, 0x04, 0x66,
	0x99, 0x14, 0xc4, 0x35, 0x3c, 0x23, 0xb0, 0x63, 0x15, 0xa3, 0x5f, 0x30, 0xc8, 0x64, 0x51, 0x9c,
	0xdc, 0x9e, 0x67, 0x04, 0xce, 0xec, 0xc7, 0xb5, 0xf0, 0xaa, 0x69, 0xcd, 0x31, 0x26, 0x9c, 0xff,
	0xff, 0x14, 0x6b, 0x12, 0xfd, 0x01, 0x73, 0x27, 0x53, 0xe2, 0xf6, 0xd5, 0xc4, 0xf8, 0x7a, 0xe2,
	0x4e, 0xa6, 0xa4, 0x2e, 0x89, 0x20, 0xbc, 0x1d, 0x53, 0x34, 0x0a, 0xc1, 0x64, 0x34, 0xc3, 0xae,
	0xa9, 0xa6, 0x6e, 0x0c, 0xdc, 0xaf, 0x57, 0xcb, 0x8e, 0x6f, 0xb8, 0x85, 0x05, 0xa6, 0x38, 0x55,
	0xc4, 0x5f, 0xc0, 0x97, 0x5b, 0x4d, 0x34, 0x02, 0x53, 0x72, 0x52, 0x6b, 0x23, 0x0b, 0xeb, 0xf5,
	0xe5, 0x67, 0xcf, 0x33, 0x62, 0x55, 0x6b, 0x4c, 0x56, 0x89, 0xc8, 0x95, 0x1f, 0x3b, 0x56, 0xb1,
	0xff, 0x64, 0x00, 0x74, 0x2b, 0x90, 0x0f, 0x9f, 0xab, 0x9a, 0x1d, 0x68, 0xf6, 0x41, 0xa2, 0xad,
	0xa3, 0x09, 0xd8, 0x78, 0x4f, 0x49, 0x29, 0x36, 0x34, 0xd3, 0x5a, 0x1d, 0xa4, 0x1b, 0xeb, 0x0c,
	0x8d, 0xc1, 0xe2, 0x04, 0xd7, 0x44, 0xa8, 0x5b, 0x74, 0xc4, 0xb9, 0x8a, 0xbe, 0x81, 0xc5, 0x31,
	0xab, 0x08, 0x77, 0x4d, 0xaf, 0x1f, 0xd8, 0xf1, 0x39, 0x6b, 0xdf, 0x38, 0xb8, 0x78, 0xe3, 0x04,
	0x9c, 0x8b, 0x6b, 0xa3, 0xaf, 0x30, 0x38, 0x24, 0x7b, 0xf9, 0xfe, 0xb1, 0x74, 0x92, 0x5a, 0xea,
	0x17, 0xf9, 0xfd, 0x16, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xea, 0xb7, 0x93, 0x83, 0x02, 0x00, 0x00,
}