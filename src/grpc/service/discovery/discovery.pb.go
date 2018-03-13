// Code generated by protoc-gen-gogo.
// source: discovery.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	discovery.proto

It has these top-level messages:
	Node
	Service
	RegistrationRequest
	RegistrationReply
	HeartbeatRequest
	HeartbeatReply
	ListRequest
	ListReply
*/
package discovery

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceType int32

const (
	ServiceType_MasterMaster ServiceType = 0
	ServiceType_MasterSlave  ServiceType = 1
)

var ServiceType_name = map[int32]string{
	0: "MasterMaster",
	1: "MasterSlave",
}
var ServiceType_value = map[string]int32{
	"MasterMaster": 0,
	"MasterSlave":  1,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}
func (ServiceType) EnumDescriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{0} }

type Node struct {
	UUID      string                     `protobuf:"bytes,1,opt,name=UUID,json=uUID,proto3" json:"UUID,omitempty"`
	URI       string                     `protobuf:"bytes,2,opt,name=URI,json=uRI,proto3" json:"URI,omitempty"`
	Leader    bool                       `protobuf:"varint,3,opt,name=leader,proto3" json:"leader,omitempty"`
	Heartbeat *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=heartbeat" json:"heartbeat,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{0} }

func (m *Node) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Node) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *Node) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

func (m *Node) GetHeartbeat() *google_protobuf.Timestamp {
	if m != nil {
		return m.Heartbeat
	}
	return nil
}

func (m *Node) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Service struct {
	Name  string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  ServiceType `protobuf:"varint,2,opt,name=type,proto3,enum=discovery.ServiceType" json:"type,omitempty"`
	Nodes []*Node     `protobuf:"bytes,3,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{1} }

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetType() ServiceType {
	if m != nil {
		return m.Type
	}
	return ServiceType_MasterMaster
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type RegistrationRequest struct {
	UUID   string      `protobuf:"bytes,1,opt,name=UUID,json=uUID,proto3" json:"UUID,omitempty"`
	URI    string      `protobuf:"bytes,2,opt,name=URI,json=uRI,proto3" json:"URI,omitempty"`
	Name   string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type   ServiceType `protobuf:"varint,4,opt,name=type,proto3,enum=discovery.ServiceType" json:"type,omitempty"`
	Leader bool        `protobuf:"varint,5,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (m *RegistrationRequest) Reset()                    { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string            { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()               {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{2} }

func (m *RegistrationRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *RegistrationRequest) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *RegistrationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegistrationRequest) GetType() ServiceType {
	if m != nil {
		return m.Type
	}
	return ServiceType_MasterMaster
}

func (m *RegistrationRequest) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

type RegistrationReply struct {
	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Node    []*Node `protobuf:"bytes,3,rep,name=node" json:"node,omitempty"`
}

func (m *RegistrationReply) Reset()                    { *m = RegistrationReply{} }
func (m *RegistrationReply) String() string            { return proto.CompactTextString(m) }
func (*RegistrationReply) ProtoMessage()               {}
func (*RegistrationReply) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{3} }

func (m *RegistrationReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegistrationReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegistrationReply) GetNode() []*Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type HeartbeatRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Uri  string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (m *HeartbeatRequest) Reset()                    { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()               {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{4} }

func (m *HeartbeatRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *HeartbeatRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type HeartbeatReply struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *HeartbeatReply) Reset()                    { *m = HeartbeatReply{} }
func (m *HeartbeatReply) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatReply) ProtoMessage()               {}
func (*HeartbeatReply) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{5} }

func (m *HeartbeatReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{6} }

func (m *ListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListReply struct {
	Services map[string]*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{7} }

func (m *ListReply) GetServices() map[string]*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "discovery.Node")
	proto.RegisterType((*Service)(nil), "discovery.Service")
	proto.RegisterType((*RegistrationRequest)(nil), "discovery.RegistrationRequest")
	proto.RegisterType((*RegistrationReply)(nil), "discovery.RegistrationReply")
	proto.RegisterType((*HeartbeatRequest)(nil), "discovery.HeartbeatRequest")
	proto.RegisterType((*HeartbeatReply)(nil), "discovery.HeartbeatReply")
	proto.RegisterType((*ListRequest)(nil), "discovery.ListRequest")
	proto.RegisterType((*ListReply)(nil), "discovery.ListReply")
	proto.RegisterEnum("discovery.ServiceType", ServiceType_name, ServiceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Discovery service

type DiscoveryClient interface {
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationReply, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationReply, error) {
	out := new(RegistrationReply)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	out := new(HeartbeatReply)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/discovery.Discovery/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Register(context.Context, *RegistrationRequest) (*RegistrationReply, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Discovery_Heartbeat_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Discovery_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptorDiscovery) }

var fileDescriptorDiscovery = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0xc7, 0x29, 0x2d, 0xbb, 0xf4, 0x54, 0x17, 0x1c, 0x0d, 0xa9, 0x68, 0x14, 0xc7, 0x98, 0x10,
	0x2e, 0x58, 0x83, 0x89, 0x41, 0x2f, 0x4c, 0x4c, 0xd8, 0x44, 0x8c, 0x1f, 0xc9, 0xec, 0x72, 0x6d,
	0x06, 0x7a, 0x16, 0x1b, 0x0b, 0xad, 0xf3, 0x41, 0xd2, 0xe7, 0xf0, 0xd6, 0x17, 0xf1, 0x45, 0x7c,
	0x1e, 0xd3, 0x0e, 0x2d, 0xdd, 0x5d, 0xd6, 0xdd, 0x1b, 0x32, 0x33, 0xe7, 0x7f, 0xce, 0xf9, 0x9d,
	0xff, 0xa1, 0xd0, 0x0a, 0x42, 0xb9, 0x88, 0x37, 0x28, 0xd2, 0x61, 0x22, 0x62, 0x15, 0x13, 0xb7,
	0x7c, 0xe8, 0x3e, 0x5d, 0xc6, 0xf1, 0x32, 0xc2, 0xe3, 0x3c, 0x30, 0xd7, 0xe7, 0xc7, 0x2a, 0x5c,
	0xa1, 0x54, 0x7c, 0x95, 0x18, 0x2d, 0xfd, 0x63, 0x81, 0xf3, 0x25, 0x0e, 0x90, 0x10, 0x70, 0x66,
	0xb3, 0xe9, 0xc4, 0xb7, 0x7a, 0x56, 0xdf, 0x65, 0x8e, 0x9e, 0x4d, 0x27, 0xa4, 0x0d, 0xf6, 0x8c,
	0x4d, 0xfd, 0x7a, 0xfe, 0x64, 0x6b, 0x36, 0x25, 0x1d, 0x38, 0x88, 0x90, 0x07, 0x28, 0x7c, 0xbb,
	0x67, 0xf5, 0x9b, 0x6c, 0x7b, 0x23, 0x63, 0x70, 0xbf, 0x23, 0x17, 0x6a, 0x8e, 0x5c, 0xf9, 0x4e,
	0xcf, 0xea, 0x7b, 0xa3, 0xee, 0xd0, 0xf4, 0x1e, 0x16, 0xbd, 0x87, 0x67, 0x45, 0x6f, 0xb6, 0x13,
	0x93, 0x37, 0x00, 0x3a, 0x09, 0xb8, 0xc2, 0xe0, 0x1b, 0x57, 0x7e, 0xe3, 0xe6, 0xd4, 0xad, 0xfa,
	0xbd, 0xa2, 0x09, 0x1c, 0x9e, 0xa2, 0xd8, 0x84, 0x8b, 0x9c, 0x7e, 0xcd, 0x57, 0x58, 0xd0, 0x67,
	0x67, 0x32, 0x00, 0x47, 0xa5, 0x09, 0xe6, 0xf8, 0x47, 0xa3, 0xce, 0x70, 0x67, 0xd3, 0x36, 0xeb,
	0x2c, 0x4d, 0x90, 0xe5, 0x1a, 0xf2, 0x02, 0x1a, 0xeb, 0x38, 0x40, 0xe9, 0xdb, 0x3d, 0xbb, 0xef,
	0x8d, 0x5a, 0x15, 0x71, 0xe6, 0x0e, 0x33, 0x51, 0xfa, 0xcb, 0x82, 0xfb, 0x0c, 0x97, 0xa1, 0x54,
	0x82, 0xab, 0x30, 0x5e, 0x33, 0xfc, 0xa9, 0x51, 0xaa, 0x5b, 0x9a, 0x57, 0x40, 0xda, 0x7b, 0x20,
	0x9d, 0x5b, 0x40, 0xee, 0xcc, 0x6f, 0x54, 0xcd, 0xa7, 0xe7, 0x70, 0xef, 0x22, 0x54, 0x12, 0xa5,
	0xc4, 0x87, 0x43, 0xa9, 0x17, 0x0b, 0x94, 0x32, 0xa7, 0x6a, 0xb2, 0xe2, 0x5a, 0x62, 0xd4, 0x2b,
	0x18, 0xcf, 0xc1, 0xc9, 0x26, 0xbc, 0x6e, 0xfc, 0x3c, 0x48, 0xc7, 0xd0, 0xfe, 0x50, 0xec, 0xad,
	0x32, 0xb9, 0xd6, 0x61, 0x50, 0x4e, 0xae, 0xc3, 0x20, 0x9b, 0x5c, 0x8b, 0xb0, 0x9c, 0x5c, 0x84,
	0x74, 0x00, 0x47, 0x95, 0xcc, 0xff, 0xe2, 0xd1, 0x67, 0xe0, 0x7d, 0x0a, 0x65, 0xb5, 0xc1, 0xe5,
	0xcd, 0xd2, 0xdf, 0x16, 0xb8, 0x46, 0x93, 0x95, 0x7a, 0x07, 0x4d, 0x69, 0xbc, 0xca, 0x6a, 0x65,
	0xfc, 0xb4, 0xc2, 0x5f, 0xea, 0x0a, 0x43, 0xe5, 0xc9, 0x5a, 0x89, 0x94, 0x95, 0x39, 0xdd, 0xaf,
	0x70, 0xf7, 0x42, 0x28, 0xe3, 0xff, 0x81, 0xe9, 0xb6, 0x63, 0x76, 0x24, 0x7d, 0x68, 0x6c, 0x78,
	0xa4, 0x8d, 0x67, 0xde, 0x88, 0x5c, 0x5d, 0x13, 0x33, 0x82, 0xb7, 0xf5, 0xb1, 0x35, 0x78, 0x09,
	0x5e, 0x65, 0x79, 0xa4, 0x0d, 0x77, 0x3e, 0x73, 0xa9, 0x50, 0x98, 0xdf, 0x76, 0x8d, 0xb4, 0xc0,
	0x33, 0xe7, 0xd3, 0x88, 0x6f, 0xb0, 0x6d, 0x8d, 0xfe, 0x5a, 0xe0, 0x4e, 0x8a, 0x92, 0xe4, 0x23,
	0x34, 0xcd, 0x3e, 0x51, 0x90, 0x27, 0x95, 0x56, 0x7b, 0xfe, 0x79, 0xdd, 0xc7, 0xd7, 0xc6, 0x93,
	0x28, 0xa5, 0x35, 0x72, 0x02, 0x6e, 0xe9, 0x3c, 0x79, 0x54, 0x11, 0x5f, 0xde, 0x64, 0xf7, 0xe1,
	0xfe, 0xa0, 0x29, 0xf3, 0x1a, 0x9c, 0xcc, 0x48, 0xd2, 0xb9, 0xe2, 0xac, 0x49, 0x7e, 0xb0, 0xcf,
	0x71, 0x5a, 0x9b, 0x1f, 0xe4, 0x5f, 0xf0, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x40,
	0xa3, 0xae, 0xa4, 0x04, 0x00, 0x00,
}