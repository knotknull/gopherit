// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reporting.proto

/*
Package reporting is a generated protocol buffer package.

It is generated from these files:
	reporting.proto

It has these top-level messages:
	Report
*/
package reporting

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import codes "github.com/gopherguides/training/distributed-systems/grpc/src/import/codes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Report struct {
	Code    codes.Code `protobuf:"varint,1,opt,name=Code,proto3,enum=codes.Code" json:"Code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptorReporting, []int{0} }

func (m *Report) GetCode() codes.Code {
	if m != nil {
		return m.Code
	}
	return codes.Code_OK
}

func (m *Report) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Report)(nil), "reporting.Report")
}

func init() { proto.RegisterFile("reporting.proto", fileDescriptorReporting) }

var fileDescriptorReporting = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8d, 0x31, 0xab, 0x83, 0x40,
	0x0c, 0xc7, 0xf1, 0xf1, 0xf0, 0xe1, 0x3d, 0x68, 0xc1, 0x49, 0xba, 0x54, 0x3a, 0xb9, 0xd4, 0x83,
	0xf6, 0x23, 0x38, 0x77, 0x71, 0xe8, 0xae, 0x77, 0xe1, 0xcc, 0xa0, 0x39, 0x92, 0x38, 0xf4, 0xdb,
	0x17, 0x4f, 0xda, 0x25, 0x90, 0x7f, 0xf2, 0xff, 0xfd, 0xcc, 0x91, 0x21, 0x12, 0x2b, 0x2e, 0xa1,
	0x8d, 0x4c, 0x4a, 0x65, 0xf1, 0x0d, 0x4e, 0xcf, 0x80, 0x3a, 0xad, 0x63, 0xeb, 0x68, 0xb6, 0x81,
	0xe2, 0x04, 0x1c, 0x56, 0xf4, 0x20, 0x56, 0x79, 0xc0, 0x05, 0x97, 0x60, 0x3d, 0x8a, 0x32, 0x8e,
	0xab, 0x82, 0xbf, 0xca, 0x4b, 0x14, 0x66, 0xb1, 0x81, 0xa3, 0xb3, 0xc2, 0xce, 0xe2, 0xbc, 0x81,
	0xac, 0xa3, 0xad, 0x92, 0xe6, 0xae, 0xb8, 0x74, 0x26, 0xef, 0x93, 0xa4, 0x3c, 0x9b, 0xdf, 0x8e,
	0x3c, 0x54, 0x59, 0x9d, 0x35, 0x87, 0xdb, 0x7f, 0xbb, 0x7f, 0x6d, 0x51, 0x9f, 0x0e, 0x65, 0x65,
	0xfe, 0x1e, 0x20, 0x32, 0x04, 0xa8, 0x7e, 0xea, 0xac, 0x29, 0xfa, 0xcf, 0x3a, 0xe6, 0x89, 0x75,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xbc, 0x2f, 0x88, 0xc1, 0x00, 0x00, 0x00,
}
