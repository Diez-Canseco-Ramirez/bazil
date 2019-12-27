// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bazil.org/bazil/db/wire/volume.proto

package wire

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

type VolumeStorage struct {
	Backend              string   `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	SharingKeyName       string   `protobuf:"bytes,2,opt,name=sharingKeyName,proto3" json:"sharingKeyName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeStorage) Reset()         { *m = VolumeStorage{} }
func (m *VolumeStorage) String() string { return proto.CompactTextString(m) }
func (*VolumeStorage) ProtoMessage()    {}
func (*VolumeStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52f12a963a22720, []int{0}
}

func (m *VolumeStorage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeStorage.Unmarshal(m, b)
}
func (m *VolumeStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeStorage.Marshal(b, m, deterministic)
}
func (m *VolumeStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeStorage.Merge(m, src)
}
func (m *VolumeStorage) XXX_Size() int {
	return xxx_messageInfo_VolumeStorage.Size(m)
}
func (m *VolumeStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeStorage.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeStorage proto.InternalMessageInfo

func (m *VolumeStorage) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *VolumeStorage) GetSharingKeyName() string {
	if m != nil {
		return m.SharingKeyName
	}
	return ""
}

func init() {
	proto.RegisterType((*VolumeStorage)(nil), "bazil.db.VolumeStorage")
}

func init() {
	proto.RegisterFile("bazil.org/bazil/db/wire/volume.proto", fileDescriptor_b52f12a963a22720)
}

var fileDescriptor_b52f12a963a22720 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4a, 0xac, 0xca,
	0xcc, 0xd1, 0xcb, 0x2f, 0x4a, 0xd7, 0x07, 0xb3, 0xf4, 0x53, 0x92, 0xf4, 0xcb, 0x33, 0x8b, 0x52,
	0xf5, 0xcb, 0xf2, 0x73, 0x4a, 0x73, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x20,
	0xaa, 0x52, 0x92, 0x94, 0x02, 0xb9, 0x78, 0xc3, 0xc0, 0x32, 0xc1, 0x25, 0xf9, 0x45, 0x89, 0xe9,
	0xa9, 0x42, 0x12, 0x5c, 0xec, 0x49, 0x89, 0xc9, 0xd9, 0xa9, 0x79, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x1a, 0x17, 0x5f, 0x71, 0x46, 0x62, 0x51, 0x66, 0x5e, 0xba,
	0x77, 0x6a, 0xa5, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x01, 0x9a, 0xa8, 0x13, 0x5b, 0x14,
	0x0b, 0xc8, 0xc6, 0x24, 0x36, 0xb0, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x6e,
	0x7e, 0xa7, 0x93, 0x00, 0x00, 0x00,
}
