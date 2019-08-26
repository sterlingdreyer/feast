// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/Feature.proto

package types // import "github.com/gojek/feast/protos/generated/go/feast/types"

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

type Feature struct {
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_Feature_a96c8241b9244062, []int{0}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Feature)(nil), "feast.types.Feature")
}

func init() { proto.RegisterFile("feast/types/Feature.proto", fileDescriptor_Feature_a96c8241b9244062) }

var fileDescriptor_Feature_a96c8241b9244062 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4b, 0x4d, 0x2c,
	0x2e, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x77, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x4b, 0xe9, 0x81, 0xa5, 0xa4, 0xc4, 0x91, 0xd5,
	0x85, 0x25, 0xe6, 0x94, 0x42, 0x55, 0x29, 0xb9, 0x73, 0xb1, 0x43, 0xb5, 0x09, 0x69, 0x70, 0xb1,
	0x96, 0x81, 0x64, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf4, 0x90, 0x0c, 0xd0, 0x03,
	0xeb, 0x09, 0x82, 0x28, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x56, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x82, 0xb9, 0x90, 0x2d, 0x74, 0xe2, 0x81, 0x9a, 0x1a, 0x00,
	0xb2, 0x25, 0xca, 0x2c, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3d,
	0x3f, 0x2b, 0x35, 0x5b, 0x1f, 0xe2, 0x1e, 0xb0, 0x1b, 0x8a, 0xf5, 0xd3, 0x53, 0xf3, 0x52, 0x8b,
	0x12, 0x4b, 0x52, 0x53, 0xf4, 0xd3, 0xf3, 0xf5, 0x91, 0x5c, 0x9a, 0xc4, 0x06, 0x56, 0x60, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x83, 0xa7, 0x49, 0xc1, 0xe7, 0x00, 0x00, 0x00,
}
