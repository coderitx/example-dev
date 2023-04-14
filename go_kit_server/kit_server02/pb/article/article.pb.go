// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package article

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

type CreateReq struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	CateId               int64    `protobuf:"varint,3,opt,name=CateId,proto3" json:"CateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateReq) GetCateId() int64 {
	if m != nil {
		return m.CateId
	}
	return 0
}

type CreateResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResp) Reset()         { *m = CreateResp{} }
func (m *CreateResp) String() string { return proto.CompactTextString(m) }
func (*CreateResp) ProtoMessage()    {}
func (*CreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{1}
}

func (m *CreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResp.Unmarshal(m, b)
}
func (m *CreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResp.Marshal(b, m, deterministic)
}
func (m *CreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResp.Merge(m, src)
}
func (m *CreateResp) XXX_Size() int {
	return xxx_messageInfo_CreateResp.Size(m)
}
func (m *CreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResp proto.InternalMessageInfo

func (m *CreateResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DetailReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailReq) Reset()         { *m = DetailReq{} }
func (m *DetailReq) String() string { return proto.CompactTextString(m) }
func (*DetailReq) ProtoMessage()    {}
func (*DetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{2}
}

func (m *DetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailReq.Unmarshal(m, b)
}
func (m *DetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailReq.Marshal(b, m, deterministic)
}
func (m *DetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailReq.Merge(m, src)
}
func (m *DetailReq) XXX_Size() int {
	return xxx_messageInfo_DetailReq.Size(m)
}
func (m *DetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_DetailReq proto.InternalMessageInfo

func (m *DetailReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DetailResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	CateId               int64    `protobuf:"varint,4,opt,name=CateId,proto3" json:"CateId,omitempty"`
	UserId               int64    `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailResp) Reset()         { *m = DetailResp{} }
func (m *DetailResp) String() string { return proto.CompactTextString(m) }
func (*DetailResp) ProtoMessage()    {}
func (*DetailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{3}
}

func (m *DetailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResp.Unmarshal(m, b)
}
func (m *DetailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResp.Marshal(b, m, deterministic)
}
func (m *DetailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResp.Merge(m, src)
}
func (m *DetailResp) XXX_Size() int {
	return xxx_messageInfo_DetailResp.Size(m)
}
func (m *DetailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResp.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResp proto.InternalMessageInfo

func (m *DetailResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DetailResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetailResp) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DetailResp) GetCateId() int64 {
	if m != nil {
		return m.CateId
	}
	return 0
}

func (m *DetailResp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateReq)(nil), "CreateReq")
	proto.RegisterType((*CreateResp)(nil), "CreateResp")
	proto.RegisterType((*DetailReq)(nil), "DetailReq")
	proto.RegisterType((*DetailResp)(nil), "DetailResp")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0x95, 0x84, 0x1a, 0xe5, 0x0a, 0x1d, 0x2c, 0x84, 0xac, 0xc2, 0xd0, 0x66, 0xea, 0x64,
	0x24, 0xf8, 0x05, 0x10, 0x16, 0xaf, 0x2e, 0x30, 0xb0, 0x99, 0xfa, 0x0d, 0x96, 0xa2, 0xa6, 0xb8,
	0x4f, 0x4c, 0xfc, 0x78, 0xd4, 0xb8, 0x69, 0x2b, 0xaa, 0x8c, 0x77, 0x4f, 0x3e, 0x7f, 0x77, 0xb8,
	0x76, 0x91, 0xc3, 0xaa, 0x21, 0xbd, 0x89, 0x2d, 0xb7, 0xd5, 0x12, 0x65, 0x1d, 0xc9, 0x31, 0x59,
	0xfa, 0x96, 0x37, 0x18, 0xbd, 0x05, 0x6e, 0x48, 0x65, 0xb3, 0x6c, 0x51, 0xda, 0x24, 0xa4, 0xc2,
	0x65, 0xdd, 0xae, 0x99, 0xd6, 0xac, 0xf2, 0xce, 0xef, 0xa5, 0xbc, 0x85, 0xa8, 0x1d, 0x93, 0xf1,
	0xaa, 0x98, 0x65, 0x8b, 0xc2, 0xee, 0x55, 0x75, 0x0f, 0xf4, 0xa1, 0xdb, 0x8d, 0x9c, 0x20, 0x37,
	0xbe, 0x8b, 0x2c, 0x6c, 0x6e, 0x7c, 0x75, 0x87, 0xf2, 0x95, 0xd8, 0x85, 0x66, 0xf7, 0xe5, 0xff,
	0xe3, 0x2f, 0xd0, 0x1f, 0xcf, 0x9f, 0x1e, 0x01, 0xf3, 0x01, 0xc0, 0x62, 0x08, 0xf0, 0xe2, 0x14,
	0x70, 0xe7, 0xbf, 0x6f, 0x29, 0x1a, 0xaf, 0x46, 0xc9, 0x4f, 0xea, 0xf1, 0x03, 0x93, 0xe7, 0x34,
	0xcf, 0x92, 0xe2, 0x4f, 0x58, 0x91, 0x9c, 0x43, 0xa4, 0x2a, 0x12, 0xfa, 0x30, 0xd4, 0x74, 0xac,
	0x4f, 0xfa, 0xcd, 0x21, 0x12, 0xb2, 0x84, 0x3e, 0x14, 0x9b, 0x8e, 0xf5, 0xb1, 0xc7, 0xcb, 0xd5,
	0x27, 0xb4, 0x7e, 0xd8, 0x2f, 0xff, 0x25, 0xba, 0xe9, 0x9f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0x09, 0x04, 0xe5, 0x8b, 0x01, 0x00, 0x00,
}
