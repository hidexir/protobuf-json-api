// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prefectures.proto

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

type Users struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_443aaaf361904ee7, []int{0}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Users) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetUsersRes struct {
	Users                []*Users `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRes) Reset()         { *m = GetUsersRes{} }
func (m *GetUsersRes) String() string { return proto.CompactTextString(m) }
func (*GetUsersRes) ProtoMessage()    {}
func (*GetUsersRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_443aaaf361904ee7, []int{1}
}

func (m *GetUsersRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRes.Unmarshal(m, b)
}
func (m *GetUsersRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRes.Marshal(b, m, deterministic)
}
func (m *GetUsersRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRes.Merge(m, src)
}
func (m *GetUsersRes) XXX_Size() int {
	return xxx_messageInfo_GetUsersRes.Size(m)
}
func (m *GetUsersRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRes proto.InternalMessageInfo

func (m *GetUsersRes) GetUsers() []*Users {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*Users)(nil), "users.Users")
	proto.RegisterType((*GetUsersRes)(nil), "users.GetUsersRes")
}

func init() { proto.RegisterFile("prefectures.proto", fileDescriptor_443aaaf361904ee7) }

var fileDescriptor_443aaaf361904ee7 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0x4a, 0x4d,
	0x4b, 0x4d, 0x2e, 0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d,
	0x2d, 0x4e, 0x2d, 0x2a, 0x56, 0xd2, 0xe6, 0x62, 0x0d, 0x05, 0x31, 0x84, 0xf8, 0xb8, 0x98, 0x32,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2,
	0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x43, 0x2e, 0x6e,
	0xf7, 0xd4, 0x12, 0xb0, 0xfa, 0xa0, 0xd4, 0x62, 0x21, 0x25, 0x2e, 0x88, 0x21, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x60, 0x9e, 0x1e, 0x44, 0x1e, 0x22, 0xe5, 0xa4, 0xcb, 0x25,
	0x95, 0x9c, 0x9f, 0xab, 0x57, 0x92, 0x98, 0x5d, 0x5a, 0x9c, 0x9a, 0x9b, 0x94, 0xa8, 0x97, 0x9e,
	0x5f, 0x9a, 0x92, 0x08, 0x71, 0x43, 0xb1, 0x13, 0x77, 0x00, 0xc2, 0x5d, 0x51, 0x4c, 0x05, 0x49,
	0x49, 0x6c, 0x60, 0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x51, 0x88, 0x6a, 0xb1,
	0x00, 0x00, 0x00,
}