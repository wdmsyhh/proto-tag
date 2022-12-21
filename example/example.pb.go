// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package staff is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Staff
*/
package staff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Staff struct {
	ID      string       `protobuf:"bytes,1,opt,name=ID"        json:"id,omitempty"     xml:"id,omitempty"`
	Name    string       `protobuf:"bytes,2,opt,name=Name"      json:"name,omitempty"   xml:"name,omitempty"`
	Age     int64        `protobuf:"varint,3,opt,name=Age"      json:"age,omitempty"    xml:"age,omitempty"`
	MyClass *Staff_Class `protobuf:"bytes,4,opt,name=MyClass"   json:"class,omitempty"  xml:"class,omitempty"`
}

func (m *Staff) Reset()                    { *m = Staff{} }
func (m *Staff) String() string            { return proto.CompactTextString(m) }
func (*Staff) ProtoMessage()               {}
func (*Staff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Staff) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Staff) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Staff) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Staff) GetMyClass() *Staff_Class {
	if m != nil {
		return m.MyClass
	}
	return nil
}

type Staff_Class struct {
	ID   string `protobuf:"bytes,1,opt,name=ID"        json:"id,omitempty"     xml:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type"      json:"type,omitempty"   xml:"type,omitempty"`
}

func (m *Staff_Class) Reset()                    { *m = Staff_Class{} }
func (m *Staff_Class) String() string            { return proto.CompactTextString(m) }
func (*Staff_Class) ProtoMessage()               {}
func (*Staff_Class) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Staff_Class) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Staff_Class) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Staff)(nil), "staff.Staff")
	proto.RegisterType((*Staff_Class)(nil), "staff.Staff.Class")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2e, 0x49, 0x4c, 0x4b,
	0x53, 0x9a, 0xc1, 0xc8, 0xc5, 0x1a, 0x0c, 0x62, 0x09, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x79, 0xba, 0x08, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21, 0x01, 0x2e, 0x66, 0xc7, 0xf4, 0x54, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0xe6, 0x20, 0x10, 0x53, 0x48, 0x87, 0x8b, 0xdd, 0xb7, 0xd2, 0x39, 0x27, 0xb1, 0xb8,
	0x58, 0x82, 0x45, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x48, 0x0f, 0x6c, 0xb0, 0x1e, 0xd8, 0x50, 0x3d,
	0xb0, 0x4c, 0x10, 0x4c, 0x89, 0x94, 0x36, 0x17, 0x2b, 0x98, 0x81, 0xcd, 0xb2, 0x90, 0xca, 0x02,
	0xb8, 0x65, 0x20, 0x76, 0x12, 0x1b, 0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef,
	0xfa, 0x73, 0x86, 0xb9, 0x00, 0x00, 0x00,
}