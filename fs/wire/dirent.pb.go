// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bazil.org/bazil/fs/wire/dirent.proto

package wire

import (
	wire "bazil.org/bazil/cas/wire"
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

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode,proto3" json:"inode,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Dirent_File
	//	*Dirent_Dir
	//	*Dirent_Tombstone
	Type                 isDirent_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Dirent) Reset()         { *m = Dirent{} }
func (m *Dirent) String() string { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()    {}
func (*Dirent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f819b9c3f7e3499a, []int{0}
}

func (m *Dirent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dirent.Unmarshal(m, b)
}
func (m *Dirent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dirent.Marshal(b, m, deterministic)
}
func (m *Dirent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dirent.Merge(m, src)
}
func (m *Dirent) XXX_Size() int {
	return xxx_messageInfo_Dirent.Size(m)
}
func (m *Dirent) XXX_DiscardUnknown() {
	xxx_messageInfo_Dirent.DiscardUnknown(m)
}

var xxx_messageInfo_Dirent proto.InternalMessageInfo

func (m *Dirent) GetInode() uint64 {
	if m != nil {
		return m.Inode
	}
	return 0
}

type isDirent_Type interface {
	isDirent_Type()
}

type Dirent_File struct {
	File *File `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type Dirent_Dir struct {
	Dir *Dir `protobuf:"bytes,3,opt,name=dir,proto3,oneof"`
}

type Dirent_Tombstone struct {
	Tombstone *Tombstone `protobuf:"bytes,4,opt,name=tombstone,proto3,oneof"`
}

func (*Dirent_File) isDirent_Type() {}

func (*Dirent_Dir) isDirent_Type() {}

func (*Dirent_Tombstone) isDirent_Type() {}

func (m *Dirent) GetType() isDirent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Dirent) GetFile() *File {
	if x, ok := m.GetType().(*Dirent_File); ok {
		return x.File
	}
	return nil
}

func (m *Dirent) GetDir() *Dir {
	if x, ok := m.GetType().(*Dirent_Dir); ok {
		return x.Dir
	}
	return nil
}

func (m *Dirent) GetTombstone() *Tombstone {
	if x, ok := m.GetType().(*Dirent_Tombstone); ok {
		return x.Tombstone
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Dirent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Dirent_File)(nil),
		(*Dirent_Dir)(nil),
		(*Dirent_Tombstone)(nil),
	}
}

type File struct {
	Manifest             *wire.Manifest `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_f819b9c3f7e3499a, []int{1}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetManifest() *wire.Manifest {
	if m != nil {
		return m.Manifest
	}
	return nil
}

// Dir is a directory stored fully in the database, not persisted
// in objects.
type Dir struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dir) Reset()         { *m = Dir{} }
func (m *Dir) String() string { return proto.CompactTextString(m) }
func (*Dir) ProtoMessage()    {}
func (*Dir) Descriptor() ([]byte, []int) {
	return fileDescriptor_f819b9c3f7e3499a, []int{2}
}

func (m *Dir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dir.Unmarshal(m, b)
}
func (m *Dir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dir.Marshal(b, m, deterministic)
}
func (m *Dir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dir.Merge(m, src)
}
func (m *Dir) XXX_Size() int {
	return xxx_messageInfo_Dir.Size(m)
}
func (m *Dir) XXX_DiscardUnknown() {
	xxx_messageInfo_Dir.DiscardUnknown(m)
}

var xxx_messageInfo_Dir proto.InternalMessageInfo

type Tombstone struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tombstone) Reset()         { *m = Tombstone{} }
func (m *Tombstone) String() string { return proto.CompactTextString(m) }
func (*Tombstone) ProtoMessage()    {}
func (*Tombstone) Descriptor() ([]byte, []int) {
	return fileDescriptor_f819b9c3f7e3499a, []int{3}
}

func (m *Tombstone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tombstone.Unmarshal(m, b)
}
func (m *Tombstone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tombstone.Marshal(b, m, deterministic)
}
func (m *Tombstone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tombstone.Merge(m, src)
}
func (m *Tombstone) XXX_Size() int {
	return xxx_messageInfo_Tombstone.Size(m)
}
func (m *Tombstone) XXX_DiscardUnknown() {
	xxx_messageInfo_Tombstone.DiscardUnknown(m)
}

var xxx_messageInfo_Tombstone proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Dirent)(nil), "bazil.db.Dirent")
	proto.RegisterType((*File)(nil), "bazil.db.File")
	proto.RegisterType((*Dir)(nil), "bazil.db.Dir")
	proto.RegisterType((*Tombstone)(nil), "bazil.db.Tombstone")
}

func init() {
	proto.RegisterFile("bazil.org/bazil/fs/wire/dirent.proto", fileDescriptor_f819b9c3f7e3499a)
}

var fileDescriptor_f819b9c3f7e3499a = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x5b, 0x9b, 0x95, 0xed, 0x0d, 0x3d, 0x44, 0x0f, 0xc5, 0xd3, 0x2c, 0x03, 0x77, 0x4a,
	0xc0, 0x1d, 0xbc, 0x4b, 0x91, 0x5e, 0xbc, 0x84, 0x9d, 0xbc, 0xb5, 0xeb, 0xab, 0x3c, 0xe8, 0x92,
	0x91, 0x04, 0x44, 0x3f, 0x8f, 0x1f, 0x54, 0x9a, 0xb5, 0x06, 0xbc, 0x25, 0xfc, 0x7e, 0xef, 0xff,
	0x92, 0x3f, 0x6c, 0xdb, 0xe6, 0x9b, 0x06, 0x61, 0xec, 0x87, 0x0c, 0x27, 0xd9, 0x3b, 0xf9, 0x49,
	0x16, 0x65, 0x47, 0x16, 0xb5, 0x17, 0x67, 0x6b, 0xbc, 0xe1, 0xcb, 0x8b, 0xd5, 0xb5, 0xf7, 0x8f,
	0xff, 0xfd, 0x63, 0x33, 0x0d, 0x9c, 0x1a, 0x4d, 0x3d, 0xba, 0x69, 0xa4, 0xfc, 0x49, 0x21, 0xaf,
	0x42, 0x06, 0xbf, 0x83, 0x05, 0x69, 0xd3, 0x61, 0x91, 0x6e, 0xd2, 0x1d, 0x53, 0x97, 0x0b, 0xdf,
	0x02, 0xeb, 0x69, 0xc0, 0xe2, 0x6a, 0x93, 0xee, 0xd6, 0x4f, 0x37, 0x62, 0x5e, 0x21, 0x5e, 0x69,
	0xc0, 0x3a, 0x51, 0x81, 0xf2, 0x07, 0xc8, 0x3a, 0xb2, 0x45, 0x16, 0xa4, 0xeb, 0x28, 0x55, 0x64,
	0xeb, 0x44, 0x8d, 0x8c, 0xef, 0x61, 0xe5, 0xcd, 0xa9, 0x75, 0xde, 0x68, 0x2c, 0x58, 0x10, 0x6f,
	0xa3, 0x78, 0x98, 0x51, 0x9d, 0xa8, 0xe8, 0xbd, 0xe4, 0xc0, 0xfc, 0xd7, 0x19, 0xcb, 0x67, 0x60,
	0xe3, 0x3e, 0x2e, 0x61, 0x39, 0x7f, 0x20, 0x3c, 0x33, 0x66, 0x1c, 0x1b, 0x27, 0xde, 0x26, 0xa4,
	0xfe, 0xa4, 0x72, 0x01, 0x59, 0x45, 0xb6, 0x5c, 0xc3, 0xea, 0x10, 0x43, 0xdf, 0xd9, 0x58, 0x45,
	0x9b, 0x87, 0x0a, 0xf6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x9c, 0x62, 0xa1, 0x5d, 0x01,
	0x00, 0x00,
}
