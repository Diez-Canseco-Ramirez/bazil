// Code generated by protoc-gen-gogo.
// source: bazil.org/bazil/fs/wire/dirent.proto
// DO NOT EDIT!

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		bazil.org/bazil/fs/wire/dirent.proto

	It has these top-level messages:
		Dirent
		Type
		File
		Dir
*/
package wire

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"
import bazil_cas "bazil.org/bazil/cas/wire"

import io1 "io"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Dirent struct {
	Inode            uint64 `protobuf:"varint,1,req,name=inode" json:"inode"`
	Type             `protobuf:"bytes,2,req,name=type,embedded=type" json:"type"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Dirent) Reset()         { *m = Dirent{} }
func (m *Dirent) String() string { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()    {}

func (m *Dirent) GetInode() uint64 {
	if m != nil {
		return m.Inode
	}
	return 0
}

type Type struct {
	File             *File  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Dir              *Dir   `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}

func (m *Type) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Type) GetDir() *Dir {
	if m != nil {
		return m.Dir
	}
	return nil
}

type File struct {
	Manifest         bazil_cas.Manifest `protobuf:"bytes,1,req,name=manifest" json:"manifest"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}

func (m *File) GetManifest() bazil_cas.Manifest {
	if m != nil {
		return m.Manifest
	}
	return bazil_cas.Manifest{}
}

type Dir struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Dir) Reset()         { *m = Dir{} }
func (m *Dir) String() string { return proto.CompactTextString(m) }
func (*Dir) ProtoMessage()    {}

func init() {
}
func (m *Dirent) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto1.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Inode |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto1.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if err := m.Type.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Type) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto1.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.File == nil {
				m.File = &File{}
			}
			if err := m.File.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto1.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.Dir == nil {
				m.Dir = &Dir{}
			}
			if err := m.Dir.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *File) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto1.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if err := m.Manifest.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Dir) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		switch fieldNum {
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *Type) GetValue() interface{} {
	if this.File != nil {
		return this.File
	}
	if this.Dir != nil {
		return this.Dir
	}
	return nil
}

func (this *Type) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *File:
		this.File = vt
	case *Dir:
		this.Dir = vt
	default:
		return false
	}
	return true
}
func (m *Dirent) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDirent(uint64(m.Inode))
	l = m.Type.Size()
	n += 1 + l + sovDirent(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	if m.File != nil {
		l = m.File.Size()
		n += 1 + l + sovDirent(uint64(l))
	}
	if m.Dir != nil {
		l = m.Dir.Size()
		n += 1 + l + sovDirent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *File) Size() (n int) {
	var l int
	_ = l
	l = m.Manifest.Size()
	n += 1 + l + sovDirent(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Dir) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDirent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDirent(x uint64) (n int) {
	return sovDirent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
	return sovDirent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Dirent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Dirent) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintDirent(data, i, uint64(m.Inode))
	data[i] = 0x12
	i++
	i = encodeVarintDirent(data, i, uint64(m.Type.Size()))
	n1, err := m.Type.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Type) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Type) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.File != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDirent(data, i, uint64(m.File.Size()))
		n2, err := m.File.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Dir != nil {
		data[i] = 0x12
		i++
		i = encodeVarintDirent(data, i, uint64(m.Dir.Size()))
		n3, err := m.Dir.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *File) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *File) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDirent(data, i, uint64(m.Manifest.Size()))
	n4, err := m.Manifest.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Dir) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Dir) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Dirent(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Dirent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDirent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
