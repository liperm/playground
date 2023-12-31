// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUseRequest) Reset()         { *m = GetUseRequest{} }
func (m *GetUseRequest) String() string { return proto.CompactTextString(m) }
func (*GetUseRequest) ProtoMessage()    {}
func (*GetUseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUseRequest.Unmarshal(m, b)
}
func (m *GetUseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUseRequest.Marshal(b, m, deterministic)
}
func (m *GetUseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUseRequest.Merge(m, src)
}
func (m *GetUseRequest) XXX_Size() int {
	return xxx_messageInfo_GetUseRequest.Size(m)
}
func (m *GetUseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUseRequest proto.InternalMessageInfo

func (m *GetUseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUseRequest)(nil), "GetUseRequest")
	proto.RegisterType((*GetUserResponse)(nil), "GetUserResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0x75, 0x4f, 0x2d, 0x09, 0x2d, 0x4e,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0x32, 0xe5, 0xe2, 0x87, 0x28, 0x28, 0x0a, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x45, 0x57, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x1b, 0x99, 0x72, 0xb1, 0x80, 0xf4, 0x08,
	0xe9, 0x72, 0xb1, 0x43, 0xb5, 0x0b, 0xf1, 0xe9, 0xa1, 0xd8, 0x24, 0x25, 0xa0, 0x87, 0x66, 0xb0,
	0x12, 0x83, 0x13, 0x5b, 0x14, 0x8b, 0x9e, 0x7e, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x75, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2d, 0x35, 0x49, 0xab, 0x00, 0x00, 0x00,
}
