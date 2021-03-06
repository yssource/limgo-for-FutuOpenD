// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_PlaceOrder/Trd_PlaceOrder.proto

package Trd_PlaceOrder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import Common "limgo/futupb/Common"
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

type C2S struct {
	PacketID  *Common.PacketID      `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`
	Header    *Trd_Common.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`
	TrdSide   *int32                `protobuf:"varint,3,req,name=trdSide" json:"trdSide,omitempty"`
	OrderType *int32                `protobuf:"varint,4,req,name=orderType" json:"orderType,omitempty"`
	Code      *string               `protobuf:"bytes,5,req,name=code" json:"code,omitempty"`
	Qty       *float64              `protobuf:"fixed64,6,req,name=qty" json:"qty,omitempty"`
	Price     *float64              `protobuf:"fixed64,7,opt,name=price" json:"price,omitempty"`
	// 以下2个为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice          *bool    `protobuf:"varint,8,opt,name=adjustPrice" json:"adjustPrice,omitempty"`
	AdjustSideAndLimit   *float64 `protobuf:"fixed64,9,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"`
	SecMarket            *int32   `protobuf:"varint,10,opt,name=secMarket" json:"secMarket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetPacketID() *Common.PacketID {
	if m != nil {
		return m.PacketID
	}
	return nil
}

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetTrdSide() int32 {
	if m != nil && m.TrdSide != nil {
		return *m.TrdSide
	}
	return 0
}

func (m *C2S) GetOrderType() int32 {
	if m != nil && m.OrderType != nil {
		return *m.OrderType
	}
	return 0
}

func (m *C2S) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *C2S) GetQty() float64 {
	if m != nil && m.Qty != nil {
		return *m.Qty
	}
	return 0
}

func (m *C2S) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *C2S) GetAdjustPrice() bool {
	if m != nil && m.AdjustPrice != nil {
		return *m.AdjustPrice
	}
	return false
}

func (m *C2S) GetAdjustSideAndLimit() float64 {
	if m != nil && m.AdjustSideAndLimit != nil {
		return *m.AdjustSideAndLimit
	}
	return 0
}

func (m *C2S) GetSecMarket() int32 {
	if m != nil && m.SecMarket != nil {
		return *m.SecMarket
	}
	return 0
}

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderID              *uint64               `protobuf:"varint,2,opt,name=orderID" json:"orderID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2, []int{1}
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

func (m *S2C) GetOrderID() uint64 {
	if m != nil && m.OrderID != nil {
		return *m.OrderID
	}
	return 0
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

// 如果下单返回的retMsg没用描述清楚错误，可再查看errCode了解详情，errCode一些取值和对应的错误描述如下:
// 2: 需要升级到保证金账户
// 3: 需要对交易期权的风险确认才能交易交易期权
// 7: 开户时选择了不希望交易衍生品
// 8: 需要对交易股权的风险确认才能交易交易股权
// 9: 需要对交易低价股的风险确认才能交易交易低价股
// 11: 需要对暗盘交易的风险确认才能进行暗盘交易
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
	return fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Trd_PlaceOrder.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_PlaceOrder.S2C")
	proto.RegisterType((*Request)(nil), "Trd_PlaceOrder.Request")
	proto.RegisterType((*Response)(nil), "Trd_PlaceOrder.Response")
}

func init() {
	proto.RegisterFile("Trd_PlaceOrder/Trd_PlaceOrder.proto", fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2)
}

var fileDescriptor_Trd_PlaceOrder_3bdd20f58d668ce2 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x96, 0xf3, 0xb3, 0x49, 0xa6, 0x12, 0xaa, 0x5c, 0x40, 0x56, 0x41, 0xc8, 0x0a, 0x42, 0xca,
	0x81, 0xa6, 0xab, 0x88, 0x13, 0x37, 0x94, 0x1e, 0xa8, 0x44, 0x61, 0xe5, 0xec, 0x1d, 0x45, 0xf1,
	0x08, 0x42, 0xd9, 0x75, 0x6a, 0xbb, 0x87, 0x9e, 0x79, 0x37, 0x9e, 0x0b, 0xd9, 0x49, 0xba, 0xbb,
	0x68, 0x0f, 0x3d, 0x65, 0xbe, 0x9f, 0x8c, 0xbf, 0x19, 0x1b, 0xde, 0xae, 0xb5, 0xfc, 0xbe, 0xfa,
	0xdd, 0x76, 0xf8, 0x4d, 0x4b, 0xd4, 0x97, 0x87, 0xb0, 0x1c, 0xb4, 0xb2, 0x8a, 0x3e, 0x3b, 0x64,
	0xcf, 0xcf, 0x6a, 0xb5, 0xd9, 0xa8, 0xed, 0xe5, 0xf8, 0x19, 0x4d, 0xe7, 0xaf, 0x9c, 0x69, 0x12,
	0x76, 0xe5, 0x28, 0xe6, 0x7f, 0x03, 0x08, 0xeb, 0xaa, 0xa1, 0xef, 0x21, 0x1d, 0xda, 0xee, 0x16,
	0xed, 0xf5, 0x15, 0x23, 0x3c, 0x28, 0x4e, 0xaa, 0xd3, 0x72, 0x32, 0xae, 0x26, 0x5e, 0x3c, 0x3a,
	0xe8, 0x05, 0x2c, 0x7e, 0x62, 0x2b, 0x51, 0xb3, 0xc0, 0x7b, 0x5f, 0x94, 0x7b, 0x8d, 0xd7, 0x5a,
	0x7e, 0xf6, 0xa2, 0x98, 0x4c, 0x94, 0x41, 0x62, 0xb5, 0x6c, 0x7a, 0x89, 0x2c, 0xe4, 0x41, 0x11,
	0x8b, 0x19, 0xd2, 0xd7, 0x90, 0x29, 0x97, 0x7c, 0xfd, 0x30, 0x20, 0x8b, 0xbc, 0xb6, 0x23, 0x28,
	0x85, 0xa8, 0x53, 0x12, 0x59, 0xcc, 0x83, 0x22, 0x13, 0xbe, 0xa6, 0xa7, 0x10, 0xde, 0xd9, 0x07,
	0xb6, 0xe0, 0x41, 0x41, 0x84, 0x2b, 0xe9, 0x73, 0x88, 0x07, 0xdd, 0x77, 0xc8, 0x12, 0x4e, 0x0a,
	0x22, 0x46, 0x40, 0x39, 0x9c, 0xb4, 0xf2, 0xd7, 0xbd, 0xb1, 0x2b, 0xaf, 0xa5, 0x9c, 0x14, 0xa9,
	0xd8, 0xa7, 0x68, 0x09, 0x74, 0x84, 0x2e, 0xc9, 0xa7, 0xad, 0xfc, 0xd2, 0x6f, 0x7a, 0xcb, 0x32,
	0xdf, 0xe4, 0x88, 0xe2, 0xb2, 0x1a, 0xec, 0x6e, 0x5a, 0x7d, 0x8b, 0x96, 0x01, 0x27, 0x2e, 0xeb,
	0x23, 0x91, 0x7f, 0x85, 0xb0, 0xa9, 0xea, 0xbd, 0xcd, 0x90, 0x27, 0x6e, 0xc6, 0x8f, 0x7b, 0x7d,
	0xc5, 0x02, 0x4e, 0x8a, 0x48, 0xcc, 0x30, 0x5f, 0x42, 0x22, 0xf0, 0xee, 0x1e, 0x8d, 0xa5, 0xef,
	0x20, 0xec, 0x2a, 0x33, 0x35, 0x3c, 0x2b, 0xff, 0x7b, 0x09, 0x75, 0xd5, 0x08, 0xa7, 0xe7, 0x7f,
	0x08, 0xa4, 0x02, 0xcd, 0xa0, 0xb6, 0x06, 0xe9, 0x1b, 0x48, 0x34, 0x5a, 0xbf, 0x56, 0xf7, 0x5f,
	0xfc, 0x31, 0xba, 0xf8, 0xb0, 0x5c, 0x8a, 0x99, 0xa4, 0x2f, 0x61, 0xa1, 0xd1, 0xde, 0x98, 0x1f,
	0xfe, 0xdc, 0x4c, 0x4c, 0xc8, 0x05, 0x42, 0xad, 0x6b, 0xe5, 0xaf, 0xca, 0x8d, 0x38, 0x43, 0x97,
	0xc2, 0x54, 0x1d, 0x8b, 0x38, 0x39, 0x96, 0xa2, 0xa9, 0x6a, 0xe1, 0xf4, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x61, 0xbc, 0x35, 0x91, 0xb8, 0x02, 0x00, 0x00,
}
