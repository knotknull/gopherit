// Code generated by protoc-gen-gogo.
// source: service.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Node
	Service
*/
package services

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Service_ServiceType int32

const (
	Service_MASTER_MASTER Service_ServiceType = 0
	Service_MASTER_SLAVE  Service_ServiceType = 1
)

var Service_ServiceType_name = map[int32]string{
	0: "MASTER_MASTER",
	1: "MASTER_SLAVE",
}
var Service_ServiceType_value = map[string]int32{
	"MASTER_MASTER": 0,
	"MASTER_SLAVE":  1,
}

func (x Service_ServiceType) String() string {
	return proto.EnumName(Service_ServiceType_name, int32(x))
}

type Node struct {
	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Uri       string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

type Service struct {
	Name string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type Service_ServiceType `protobuf:"varint,2,opt,name=type,proto3,enum=services.Service_ServiceType" json:"type,omitempty"`
	Node []*Node             `protobuf:"bytes,3,rep,name=node" json:"node,omitempty"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}

func (m *Service) GetNode() []*Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "services.Node")
	proto.RegisterType((*Service)(nil), "services.service")
	proto.RegisterEnum("services.Service_ServiceType", Service_ServiceType_name, Service_ServiceType_value)
}
