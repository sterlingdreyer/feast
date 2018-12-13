// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/FeatureGroupSpec.proto

package specs // import "github.com/gojektech/feast/go-feast-proto/feast/specs"

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

type FeatureGroupSpec struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags                 []string    `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	DataStores           *DataStores `protobuf:"bytes,3,opt,name=dataStores,proto3" json:"dataStores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FeatureGroupSpec) Reset()         { *m = FeatureGroupSpec{} }
func (m *FeatureGroupSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureGroupSpec) ProtoMessage()    {}
func (*FeatureGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureGroupSpec_3b2276702dd15bc5, []int{0}
}
func (m *FeatureGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureGroupSpec.Unmarshal(m, b)
}
func (m *FeatureGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureGroupSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureGroupSpec.Merge(dst, src)
}
func (m *FeatureGroupSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureGroupSpec.Size(m)
}
func (m *FeatureGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureGroupSpec proto.InternalMessageInfo

func (m *FeatureGroupSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FeatureGroupSpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FeatureGroupSpec) GetDataStores() *DataStores {
	if m != nil {
		return m.DataStores
	}
	return nil
}

func init() {
	proto.RegisterType((*FeatureGroupSpec)(nil), "feast.specs.FeatureGroupSpec")
}

func init() {
	proto.RegisterFile("feast/specs/FeatureGroupSpec.proto", fileDescriptor_FeatureGroupSpec_3b2276702dd15bc5)
}

var fileDescriptor_FeatureGroupSpec_3b2276702dd15bc5 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x4b, 0x4d, 0x2c,
	0x2e, 0xd1, 0x2f, 0x2e, 0x48, 0x4d, 0x2e, 0xd6, 0x77, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0x75,
	0x2f, 0xca, 0x2f, 0x2d, 0x08, 0x2e, 0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x06, 0xab, 0xd1, 0x03, 0xab, 0x91, 0x92, 0xc5, 0xa2, 0x01, 0xa1, 0x56, 0x29, 0x9f, 0x4b, 0x00,
	0xdd, 0x14, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6,
	0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xc4, 0xf4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xce,
	0x20, 0x30, 0x5b, 0xc8, 0x9c, 0x8b, 0x2b, 0x25, 0xb1, 0x24, 0x31, 0xb8, 0x24, 0xbf, 0x28, 0xb5,
	0x58, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5c, 0x0f, 0xc9, 0x62, 0x3d, 0x17, 0xb8, 0x74,
	0x10, 0x92, 0x52, 0xa7, 0x68, 0x2e, 0x64, 0xe7, 0x39, 0x89, 0xa2, 0xdb, 0x1e, 0x00, 0x72, 0x56,
	0x94, 0x69, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x56,
	0x6a, 0x76, 0x49, 0x6a, 0x72, 0x86, 0x3e, 0xc4, 0x2b, 0xe9, 0xf9, 0xba, 0x60, 0x86, 0x2e, 0xd8,
	0x07, 0xfa, 0x48, 0xfe, 0x4b, 0x62, 0x03, 0x0b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x2e, 0xdb, 0x62, 0x26, 0x01, 0x00, 0x00,
}