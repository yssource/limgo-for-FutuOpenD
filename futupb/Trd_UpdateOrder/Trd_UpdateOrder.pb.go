// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_UpdateOrder/Trd_UpdateOrder.proto

package Trd_UpdateOrder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "limgo/futupb/Common"
import Trd_Common "limgo/futupb/Trd_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Order                *Trd_Common.Order     `protobuf:"bytes,2,req,name=order" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UpdateOrder_c8ddf54d6a2439b2, []int{0}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrder() *Trd_Common.Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UpdateOrder_c8ddf54d6a2439b2, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "Trd_UpdateOrder.S2C")
	proto.RegisterType((*Response)(nil), "Trd_UpdateOrder.Response")
}

func init() {
	proto.RegisterFile("Trd_UpdateOrder/Trd_UpdateOrder.proto", fileDescriptor_Trd_UpdateOrder_c8ddf54d6a2439b2)
}

var fileDescriptor_Trd_UpdateOrder_c8ddf54d6a2439b2 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xc9, 0x5d, 0xaf, 0xd5, 0xaf, 0x83, 0x18, 0x7f, 0x10, 0x2a, 0xc8, 0x51, 0x50, 0xb3,
	0xf4, 0x5a, 0x82, 0x93, 0x6b, 0x16, 0x17, 0x11, 0xd2, 0x73, 0x14, 0x29, 0xe6, 0x43, 0x97, 0x36,
	0xe1, 0x4b, 0x16, 0x77, 0xff, 0x70, 0xb9, 0xe4, 0x8a, 0xc7, 0x4d, 0xc9, 0xfb, 0x3e, 0x0f, 0xbc,
	0x21, 0x70, 0xd7, 0x92, 0xfd, 0x78, 0xf3, 0x76, 0x17, 0xf1, 0x95, 0x2c, 0xd2, 0x7a, 0x94, 0x1b,
	0x4f, 0x2e, 0x3a, 0x7e, 0x36, 0xaa, 0x17, 0x17, 0xda, 0xed, 0xf7, 0xee, 0xb0, 0xce, 0x47, 0xb6,
	0x16, 0x37, 0x9d, 0xd5, 0x83, 0xff, 0x6b, 0x86, 0xcb, 0x77, 0x28, 0xb7, 0x4a, 0xf3, 0x15, 0x4c,
	0xbf, 0x71, 0x67, 0x91, 0x04, 0xab, 0x0b, 0x39, 0x57, 0x57, 0xcd, 0xc0, 0x6c, 0xc9, 0x3e, 0x27,
	0x68, 0x7a, 0x89, 0x3f, 0x40, 0xe5, 0xba, 0x41, 0x51, 0x24, 0xfb, 0x7c, 0x68, 0xa7, 0x97, 0x98,
	0xcc, 0x97, 0xbf, 0x0c, 0x4e, 0x0c, 0x06, 0xef, 0x0e, 0x01, 0xf9, 0x2d, 0xcc, 0x08, 0x63, 0xfb,
	0xe3, 0x31, 0xad, 0x54, 0x4f, 0x93, 0xd5, 0xe3, 0x66, 0x63, 0x8e, 0x25, 0xbf, 0x86, 0x29, 0x61,
	0x7c, 0x09, 0x5f, 0xa2, 0xa8, 0x99, 0x3c, 0x35, 0x7d, 0xe2, 0x02, 0x66, 0x48, 0xa4, 0x9d, 0x45,
	0x51, 0xd6, 0x4c, 0x56, 0xe6, 0x18, 0xf9, 0x3d, 0x94, 0x41, 0x7d, 0x8a, 0x49, 0xcd, 0xe4, 0x5c,
	0x5d, 0x36, 0xe3, 0x5f, 0xda, 0x2a, 0x6d, 0x3a, 0xe1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x61,
	0xb9, 0x48, 0x50, 0x01, 0x00, 0x00,
}