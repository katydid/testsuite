// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topsyturvy.proto

package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TopsyTurvy struct {
	Top                  *Topsy   `protobuf:"bytes,1,opt,name=Top" json:"Top,omitempty"`
	Turf                 *Turvy   `protobuf:"bytes,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopsyTurvy) Reset()         { *m = TopsyTurvy{} }
func (m *TopsyTurvy) String() string { return proto.CompactTextString(m) }
func (*TopsyTurvy) ProtoMessage()    {}
func (*TopsyTurvy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{0}
}
func (m *TopsyTurvy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopsyTurvy.Unmarshal(m, b)
}
func (m *TopsyTurvy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopsyTurvy.Marshal(b, m, deterministic)
}
func (dst *TopsyTurvy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopsyTurvy.Merge(dst, src)
}
func (m *TopsyTurvy) XXX_Size() int {
	return xxx_messageInfo_TopsyTurvy.Size(m)
}
func (m *TopsyTurvy) XXX_DiscardUnknown() {
	xxx_messageInfo_TopsyTurvy.DiscardUnknown(m)
}

var xxx_messageInfo_TopsyTurvy proto.InternalMessageInfo

func (m *TopsyTurvy) GetTop() *Topsy {
	if m != nil {
		return m.Top
	}
	return nil
}

func (m *TopsyTurvy) GetTurf() *Turvy {
	if m != nil {
		return m.Turf
	}
	return nil
}

type Topsy struct {
	Top                  *int64   `protobuf:"varint,1,opt,name=Top" json:"Top,omitempty"`
	Turf                 *int64   `protobuf:"varint,2,opt,name=Turf" json:"Turf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topsy) Reset()         { *m = Topsy{} }
func (m *Topsy) String() string { return proto.CompactTextString(m) }
func (*Topsy) ProtoMessage()    {}
func (*Topsy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{1}
}
func (m *Topsy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topsy.Unmarshal(m, b)
}
func (m *Topsy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topsy.Marshal(b, m, deterministic)
}
func (dst *Topsy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topsy.Merge(dst, src)
}
func (m *Topsy) XXX_Size() int {
	return xxx_messageInfo_Topsy.Size(m)
}
func (m *Topsy) XXX_DiscardUnknown() {
	xxx_messageInfo_Topsy.DiscardUnknown(m)
}

var xxx_messageInfo_Topsy proto.InternalMessageInfo

func (m *Topsy) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func (m *Topsy) GetTurf() int64 {
	if m != nil && m.Turf != nil {
		return *m.Turf
	}
	return 0
}

type Turvy struct {
	Turf                 *int64   `protobuf:"varint,1,opt,name=Turf" json:"Turf,omitempty"`
	Top                  *int64   `protobuf:"varint,2,opt,name=Top" json:"Top,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Turvy) Reset()         { *m = Turvy{} }
func (m *Turvy) String() string { return proto.CompactTextString(m) }
func (*Turvy) ProtoMessage()    {}
func (*Turvy) Descriptor() ([]byte, []int) {
	return fileDescriptor_topsyturvy_7bb8c7da1903a704, []int{2}
}
func (m *Turvy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Turvy.Unmarshal(m, b)
}
func (m *Turvy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Turvy.Marshal(b, m, deterministic)
}
func (dst *Turvy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Turvy.Merge(dst, src)
}
func (m *Turvy) XXX_Size() int {
	return xxx_messageInfo_Turvy.Size(m)
}
func (m *Turvy) XXX_DiscardUnknown() {
	xxx_messageInfo_Turvy.DiscardUnknown(m)
}

var xxx_messageInfo_Turvy proto.InternalMessageInfo

func (m *Turvy) GetTurf() int64 {
	if m != nil && m.Turf != nil {
		return *m.Turf
	}
	return 0
}

func (m *Turvy) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func init() {
	proto.RegisterType((*TopsyTurvy)(nil), "main.TopsyTurvy")
	proto.RegisterType((*Topsy)(nil), "main.Topsy")
	proto.RegisterType((*Turvy)(nil), "main.Turvy")
}
func (this *TopsyTurvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Topsy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func (this *Turvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return TopsyturvyDescription()
}
func TopsyturvyDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3837 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0xff, 0x24, 0xf2, 0x90, 0xa2, 0x46, 0x57, 0xf2, 0x2e, 0x57, 0x8e, 0xbd, 0xbb, 0xb4,
		0x1d, 0xcb, 0xeb, 0x9a, 0x1b, 0xac, 0xbd, 0x6b, 0x2f, 0xb7, 0x89, 0x4b, 0x51, 0x5c, 0x85, 0xae,
		0x24, 0x32, 0x43, 0x2a, 0xfe, 0x09, 0x8a, 0xc1, 0x68, 0x78, 0x49, 0xcd, 0xee, 0x70, 0x66, 0x32,
		0x33, 0xdc, 0xb5, 0x16, 0x7d, 0x70, 0xe1, 0xfe, 0x20, 0x28, 0xfa, 0x5f, 0xa0, 0x89, 0xeb, 0xb8,
		0x4d, 0x8a, 0xd6, 0x6d, 0xfa, 0x97, 0x34, 0x6d, 0xda, 0xa4, 0x2f, 0x7d, 0x49, 0xdb, 0xa7, 0x02,
		0x79, 0xef, 0x43, 0x7f, 0x5c, 0xf4, 0xcf, 0x6d, 0xd2, 0xd6, 0x0f, 0x05, 0x8c, 0x02, 0xc5, 0xfd,
		0x1b, 0xce, 0x0c, 0x29, 0xcd, 0x28, 0x80, 0xed, 0x27, 0x69, 0xce, 0x3d, 0xdf, 0x37, 0xe7, 0x9e,
		0x7b, 0xee, 0x39, 0xe7, 0xde, 0x21, 0x7c, 0xf7, 0x3a, 0x5c, 0x18, 0x59, 0xd6, 0xc8, 0xc0, 0x97,
		0x6d, 0xc7, 0xf2, 0xac, 0x83, 0xc9, 0xf0, 0xf2, 0x00, 0xbb, 0x9a, 0xa3, 0xdb, 0x9e, 0xe5, 0xd4,
		0xa8, 0x0c, 0x2d, 0x33, 0x8d, 0x9a, 0xd0, 0xa8, 0xee, 0xc2, 0xca, 0x4d, 0xdd, 0xc0, 0x5b, 0xbe,
		0x62, 0x0f, 0x7b, 0xe8, 0x59, 0xc8, 0x0e, 0x75, 0x03, 0x57, 0x52, 0x17, 0x32, 0x1b, 0xc5, 0x2b,
		0x8f, 0xd4, 0x22, 0xa0, 0x5a, 0x18, 0xd1, 0x25, 0x62, 0x99, 0x22, 0xaa, 0x6f, 0x67, 0x61, 0x75,
		0xce, 0x28, 0x42, 0x90, 0x35, 0xd5, 0x31, 0x61, 0x4c, 0x6d, 0x14, 0x64, 0xfa, 0x3f, 0xaa, 0xc0,
		0xa2, 0xad, 0x6a, 0xb7, 0xd5, 0x11, 0xae, 0xa4, 0xa9, 0x58, 0x3c, 0xa2, 0x87, 0x00, 0x06, 0xd8,
		0xc6, 0xe6, 0x00, 0x9b, 0xda, 0x51, 0x25, 0x73, 0x21, 0xb3, 0x51, 0x90, 0x03, 0x12, 0xf4, 0x04,
		0xac, 0xd8, 0x93, 0x03, 0x43, 0xd7, 0x94, 0x80, 0x1a, 0x5c, 0xc8, 0x6c, 0xe4, 0x64, 0x89, 0x0d,
		0x6c, 0x4d, 0x95, 0x1f, 0x83, 0xe5, 0xbb, 0x58, 0xbd, 0x1d, 0x54, 0x2d, 0x52, 0xd5, 0x32, 0x11,
		0x07, 0x14, 0x9b, 0x50, 0x1a, 0x63, 0xd7, 0x55, 0x47, 0x58, 0xf1, 0x8e, 0x6c, 0x5c, 0xc9, 0xd2,
		0xd9, 0x5f, 0x98, 0x99, 0x7d, 0x74, 0xe6, 0x45, 0x8e, 0xea, 0x1f, 0xd9, 0x18, 0x35, 0xa0, 0x80,
		0xcd, 0xc9, 0x98, 0x31, 0xe4, 0x8e, 0xf1, 0x5f, 0xcb, 0x9c, 0x8c, 0xa3, 0x2c, 0x79, 0x02, 0xe3,
		0x14, 0x8b, 0x2e, 0x76, 0xee, 0xe8, 0x1a, 0xae, 0x2c, 0x50, 0x82, 0xc7, 0x66, 0x08, 0x7a, 0x6c,
		0x3c, 0xca, 0x21, 0x70, 0xa8, 0x09, 0x05, 0xfc, 0x8a, 0x87, 0x4d, 0x57, 0xb7, 0xcc, 0xca, 0x22,
		0x25, 0x79, 0x74, 0xce, 0x2a, 0x62, 0x63, 0x10, 0xa5, 0x98, 0xe2, 0xd0, 0x35, 0x58, 0xb4, 0x6c,
		0x4f, 0xb7, 0x4c, 0xb7, 0x92, 0xbf, 0x90, 0xda, 0x28, 0x5e, 0xf9, 0xc8, 0xdc, 0x40, 0xe8, 0x30,
		0x1d, 0x59, 0x28, 0xa3, 0x36, 0x48, 0xae, 0x35, 0x71, 0x34, 0xac, 0x68, 0xd6, 0x00, 0x2b, 0xba,
		0x39, 0xb4, 0x2a, 0x05, 0x4a, 0x70, 0x7e, 0x76, 0x22, 0x54, 0xb1, 0x69, 0x0d, 0x70, 0xdb, 0x1c,
		0x5a, 0x72, 0xd9, 0x0d, 0x3d, 0xa3, 0x33, 0xb0, 0xe0, 0x1e, 0x99, 0x9e, 0xfa, 0x4a, 0xa5, 0x44,
		0x23, 0x84, 0x3f, 0x55, 0xbf, 0xb9, 0x00, 0xcb, 0x49, 0x42, 0xec, 0x06, 0xe4, 0x86, 0x64, 0x96,
		0x95, 0xf4, 0x69, 0x7c, 0xc0, 0x30, 0x61, 0x27, 0x2e, 0x7c, 0x9f, 0x4e, 0x6c, 0x40, 0xd1, 0xc4,
		0xae, 0x87, 0x07, 0x2c, 0x22, 0x32, 0x09, 0x63, 0x0a, 0x18, 0x68, 0x36, 0xa4, 0xb2, 0xdf, 0x57,
		0x48, 0xbd, 0x08, 0xcb, 0xbe, 0x49, 0x8a, 0xa3, 0x9a, 0x23, 0x11, 0x9b, 0x97, 0xe3, 0x2c, 0xa9,
		0xb5, 0x04, 0x4e, 0x26, 0x30, 0xb9, 0x8c, 0x43, 0xcf, 0x68, 0x0b, 0xc0, 0x32, 0xb1, 0x35, 0x54,
		0x06, 0x58, 0x33, 0x2a, 0xf9, 0x63, 0xbc, 0xd4, 0x21, 0x2a, 0x33, 0x5e, 0xb2, 0x98, 0x54, 0x33,
		0xd0, 0xf5, 0x69, 0xa8, 0x2d, 0x1e, 0x13, 0x29, 0xbb, 0x6c, 0x93, 0xcd, 0x44, 0xdb, 0x3e, 0x94,
		0x1d, 0x4c, 0xe2, 0x1e, 0x0f, 0xf8, 0xcc, 0x0a, 0xd4, 0x88, 0x5a, 0xec, 0xcc, 0x64, 0x0e, 0x63,
		0x13, 0x5b, 0x72, 0x82, 0x8f, 0xe8, 0x61, 0xf0, 0x05, 0x0a, 0x0d, 0x2b, 0xa0, 0x59, 0xa8, 0x24,
		0x84, 0x7b, 0xea, 0x18, 0xaf, 0xdf, 0x83, 0x72, 0xd8, 0x3d, 0x68, 0x0d, 0x72, 0xae, 0xa7, 0x3a,
		0x1e, 0x8d, 0xc2, 0x9c, 0xcc, 0x1e, 0x90, 0x04, 0x19, 0x6c, 0x0e, 0x68, 0x96, 0xcb, 0xc9, 0xe4,
		0x5f, 0xf4, 0x43, 0xd3, 0x09, 0x67, 0xe8, 0x84, 0x3f, 0x3a, 0xbb, 0xa2, 0x21, 0xe6, 0xe8, 0xbc,
		0xd7, 0x9f, 0x81, 0xa5, 0xd0, 0x04, 0x92, 0xbe, 0xba, 0xfa, 0xa3, 0x70, 0xff, 0x5c, 0x6a, 0xf4,
		0x22, 0xac, 0x4d, 0x4c, 0xdd, 0xf4, 0xb0, 0x63, 0x3b, 0x98, 0x44, 0x2c, 0x7b, 0x55, 0xe5, 0x9f,
		0x17, 0x8f, 0x89, 0xb9, 0xfd, 0xa0, 0x36, 0x63, 0x91, 0x57, 0x27, 0xb3, 0xc2, 0x4b, 0x85, 0xfc,
		0xbf, 0x2c, 0x4a, 0xaf, 0xbe, 0xfa, 0xea, 0xab, 0xe9, 0xea, 0xe7, 0x17, 0x60, 0x6d, 0xde, 0x9e,
		0x99, 0xbb, 0x7d, 0xcf, 0xc0, 0x82, 0x39, 0x19, 0x1f, 0x60, 0x87, 0x3a, 0x29, 0x27, 0xf3, 0x27,
		0xd4, 0x80, 0x9c, 0xa1, 0x1e, 0x60, 0xa3, 0x92, 0xbd, 0x90, 0xda, 0x28, 0x5f, 0x79, 0x22, 0xd1,
		0xae, 0xac, 0xed, 0x10, 0x88, 0xcc, 0x90, 0xe8, 0x13, 0x90, 0xe5, 0x29, 0x9a, 0x30, 0x5c, 0x4a,
		0xc6, 0x40, 0xf6, 0x92, 0x4c, 0x71, 0xe8, 0x01, 0x28, 0x90, 0xbf, 0x2c, 0x36, 0x16, 0xa8, 0xcd,
		0x79, 0x22, 0x20, 0x71, 0x81, 0xd6, 0x21, 0x4f, 0xb7, 0xc9, 0x00, 0x8b, 0xd2, 0xe6, 0x3f, 0x93,
		0xc0, 0x1a, 0xe0, 0xa1, 0x3a, 0x31, 0x3c, 0xe5, 0x8e, 0x6a, 0x4c, 0x30, 0x0d, 0xf8, 0x82, 0x5c,
		0xe2, 0xc2, 0x4f, 0x13, 0x19, 0x3a, 0x0f, 0x45, 0xb6, 0xab, 0x74, 0x73, 0x80, 0x5f, 0xa1, 0xd9,
		0x33, 0x27, 0xb3, 0x8d, 0xd6, 0x26, 0x12, 0xf2, 0xfa, 0x5b, 0xae, 0x65, 0x8a, 0xd0, 0xa4, 0xaf,
		0x20, 0x02, 0xfa, 0xfa, 0x67, 0xa2, 0x89, 0xfb, 0xc1, 0xf9, 0xd3, 0x8b, 0xc6, 0x54, 0xf5, 0x1b,
		0x69, 0xc8, 0xd2, 0x7c, 0xb1, 0x0c, 0xc5, 0xfe, 0x4b, 0xdd, 0x96, 0xb2, 0xd5, 0xd9, 0xdf, 0xdc,
		0x69, 0x49, 0x29, 0x54, 0x06, 0xa0, 0x82, 0x9b, 0x3b, 0x9d, 0x46, 0x5f, 0x4a, 0xfb, 0xcf, 0xed,
		0xbd, 0xfe, 0xb5, 0xa7, 0xa5, 0x8c, 0x0f, 0xd8, 0x67, 0x82, 0x6c, 0x50, 0xe1, 0xa9, 0x2b, 0x52,
		0x0e, 0x49, 0x50, 0x62, 0x04, 0xed, 0x17, 0x5b, 0x5b, 0xd7, 0x9e, 0x96, 0x16, 0xc2, 0x92, 0xa7,
		0xae, 0x48, 0x8b, 0x68, 0x09, 0x0a, 0x54, 0xb2, 0xd9, 0xe9, 0xec, 0x48, 0x79, 0x9f, 0xb3, 0xd7,
		0x97, 0xdb, 0x7b, 0xdb, 0x52, 0xc1, 0xe7, 0xdc, 0x96, 0x3b, 0xfb, 0x5d, 0x09, 0x7c, 0x86, 0xdd,
		0x56, 0xaf, 0xd7, 0xd8, 0x6e, 0x49, 0x45, 0x5f, 0x63, 0xf3, 0xa5, 0x7e, 0xab, 0x27, 0x95, 0x42,
		0x66, 0x3d, 0x75, 0x45, 0x5a, 0xf2, 0x5f, 0xd1, 0xda, 0xdb, 0xdf, 0x95, 0xca, 0x68, 0x05, 0x96,
		0xd8, 0x2b, 0x84, 0x11, 0xcb, 0x11, 0xd1, 0xb5, 0xa7, 0x25, 0x69, 0x6a, 0x08, 0x63, 0x59, 0x09,
		0x09, 0xae, 0x3d, 0x2d, 0xa1, 0x6a, 0x13, 0x72, 0x34, 0xba, 0x10, 0x82, 0xf2, 0x4e, 0x63, 0xb3,
		0xb5, 0xa3, 0x74, 0xba, 0xfd, 0x76, 0x67, 0xaf, 0xb1, 0x23, 0xa5, 0xa6, 0x32, 0xb9, 0xf5, 0xa9,
		0xfd, 0xb6, 0xdc, 0xda, 0x92, 0xd2, 0x41, 0x59, 0xb7, 0xd5, 0xe8, 0xb7, 0xb6, 0xa4, 0x4c, 0x55,
		0x83, 0xb5, 0x79, 0x79, 0x72, 0xee, 0xce, 0x08, 0x2c, 0x71, 0xfa, 0x98, 0x25, 0xa6, 0x5c, 0x33,
		0x4b, 0xfc, 0x8f, 0x69, 0x58, 0x9d, 0x53, 0x2b, 0xe6, 0xbe, 0xe4, 0x39, 0xc8, 0xb1, 0x10, 0x65,
		0xd5, 0xf3, 0xf1, 0xb9, 0x45, 0x87, 0x06, 0xec, 0x4c, 0x05, 0xa5, 0xb8, 0x60, 0x07, 0x91, 0x39,
		0xa6, 0x83, 0x20, 0x14, 0x33, 0x39, 0xfd, 0x47, 0x66, 0x72, 0x3a, 0x2b, 0x7b, 0xd7, 0x92, 0x94,
		0x3d, 0x2a, 0x3b, 0x5d, 0x6e, 0xcf, 0xcd, 0xc9, 0xed, 0x37, 0x60, 0x65, 0x86, 0x28, 0x71, 0x8e,
		0x7d, 0x2d, 0x05, 0x95, 0xe3, 0x9c, 0x13, 0x93, 0xe9, 0xd2, 0xa1, 0x4c, 0x77, 0x23, 0xea, 0xc1,
		0x8b, 0xc7, 0x2f, 0xc2, 0xcc, 0x5a, 0xbf, 0x95, 0x82, 0x33, 0xf3, 0x3b, 0xc5, 0xb9, 0x36, 0x7c,
		0x02, 0x16, 0xc6, 0xd8, 0x3b, 0xb4, 0x44, 0xb7, 0xf4, 0xd1, 0x39, 0x35, 0x98, 0x0c, 0x47, 0x17,
		0x9b, 0xa3, 0x82, 0x45, 0x3c, 0x73, 0x5c, 0xbb, 0xc7, 0xac, 0x99, 0xb1, 0xf4, 0x73, 0x69, 0xb8,
		0x7f, 0x2e, 0xf9, 0x5c, 0x43, 0x1f, 0x04, 0xd0, 0x4d, 0x7b, 0xe2, 0xb1, 0x8e, 0x88, 0x25, 0xd8,
		0x02, 0x95, 0xd0, 0xe4, 0x45, 0x92, 0xe7, 0xc4, 0xf3, 0xc7, 0x33, 0x74, 0x1c, 0x98, 0x88, 0x2a,
		0x3c, 0x3b, 0x35, 0x34, 0x4b, 0x0d, 0x7d, 0xe8, 0x98, 0x99, 0xce, 0x04, 0xe6, 0xc7, 0x40, 0xd2,
		0x0c, 0x1d, 0x9b, 0x9e, 0xe2, 0x7a, 0x0e, 0x56, 0xc7, 0xba, 0x39, 0xa2, 0x15, 0x24, 0x5f, 0xcf,
		0x0d, 0x55, 0xc3, 0xc5, 0xf2, 0x32, 0x1b, 0xee, 0x89, 0x51, 0x82, 0xa0, 0x01, 0xe4, 0x04, 0x10,
		0x0b, 0x21, 0x04, 0x1b, 0xf6, 0x11, 0xd5, 0xaf, 0xe7, 0xa1, 0x18, 0xe8, 0xab, 0xd1, 0x45, 0x28,
		0xdd, 0x52, 0xef, 0xa8, 0x8a, 0x38, 0x2b, 0x31, 0x4f, 0x14, 0x89, 0xac, 0xcb, 0xcf, 0x4b, 0x1f,
		0x83, 0x35, 0xaa, 0x62, 0x4d, 0x3c, 0xec, 0x28, 0x9a, 0xa1, 0xba, 0x2e, 0x75, 0x5a, 0x9e, 0xaa,
		0x22, 0x32, 0xd6, 0x21, 0x43, 0x4d, 0x31, 0x82, 0xae, 0xc2, 0x2a, 0x45, 0x8c, 0x27, 0x86, 0xa7,
		0xdb, 0x06, 0x56, 0xc8, 0xe9, 0xcd, 0xa5, 0x95, 0xc4, 0xb7, 0x6c, 0x85, 0x68, 0xec, 0x72, 0x05,
		0x62, 0x91, 0x8b, 0xb6, 0xe0, 0x41, 0x0a, 0x1b, 0x61, 0x13, 0x3b, 0xaa, 0x87, 0x15, 0xfc, 0xd9,
		0x89, 0x6a, 0xb8, 0x8a, 0x6a, 0x0e, 0x94, 0x43, 0xd5, 0x3d, 0xac, 0xac, 0x11, 0x82, 0xcd, 0x74,
		0x25, 0x25, 0x9f, 0x23, 0x8a, 0xdb, 0x5c, 0xaf, 0x45, 0xd5, 0x1a, 0xe6, 0xe0, 0x93, 0xaa, 0x7b,
		0x88, 0xea, 0x70, 0x86, 0xb2, 0xb8, 0x9e, 0xa3, 0x9b, 0x23, 0x45, 0x3b, 0xc4, 0xda, 0x6d, 0x65,
		0xe2, 0x0d, 0x9f, 0xad, 0x3c, 0x10, 0x7c, 0x3f, 0xb5, 0xb0, 0x47, 0x75, 0x9a, 0x44, 0x65, 0xdf,
		0x1b, 0x3e, 0x8b, 0x7a, 0x50, 0x22, 0x8b, 0x31, 0xd6, 0xef, 0x61, 0x65, 0x68, 0x39, 0xb4, 0x34,
		0x96, 0xe7, 0xa4, 0xa6, 0x80, 0x07, 0x6b, 0x1d, 0x0e, 0xd8, 0xb5, 0x06, 0xb8, 0x9e, 0xeb, 0x75,
		0x5b, 0xad, 0x2d, 0xb9, 0x28, 0x58, 0x6e, 0x5a, 0x0e, 0x09, 0xa8, 0x91, 0xe5, 0x3b, 0xb8, 0xc8,
		0x02, 0x6a, 0x64, 0x09, 0xf7, 0x5e, 0x85, 0x55, 0x4d, 0x63, 0x73, 0xd6, 0x35, 0x85, 0x9f, 0xb1,
		0xdc, 0x8a, 0x14, 0x72, 0x96, 0xa6, 0x6d, 0x33, 0x05, 0x1e, 0xe3, 0x2e, 0xba, 0x0e, 0xf7, 0x4f,
		0x9d, 0x15, 0x04, 0xae, 0xcc, 0xcc, 0x32, 0x0a, 0xbd, 0x0a, 0xab, 0xf6, 0xd1, 0x2c, 0x10, 0x85,
		0xde, 0x68, 0x1f, 0x45, 0x61, 0xcf, 0xc0, 0x9a, 0x7d, 0x68, 0xcf, 0xe2, 0x2e, 0x05, 0x71, 0xc8,
		0x3e, 0xb4, 0xa3, 0xc0, 0x47, 0xe9, 0x81, 0xdb, 0xc1, 0x9a, 0xea, 0xe1, 0x41, 0xe5, 0x6c, 0x50,
		0x3d, 0x30, 0x80, 0x2e, 0x83, 0xa4, 0x69, 0x0a, 0x36, 0xd5, 0x03, 0x03, 0x2b, 0xaa, 0x83, 0x4d,
		0xd5, 0xad, 0x9c, 0x0f, 0x2a, 0x97, 0x35, 0xad, 0x45, 0x47, 0x1b, 0x74, 0x10, 0x5d, 0x82, 0x15,
		0xeb, 0xe0, 0x96, 0xc6, 0x42, 0x52, 0xb1, 0x1d, 0x3c, 0xd4, 0x5f, 0xa9, 0x3c, 0x42, 0xfd, 0xbb,
		0x4c, 0x06, 0x68, 0x40, 0x76, 0xa9, 0x18, 0x3d, 0x0e, 0x92, 0xe6, 0x1e, 0xaa, 0x8e, 0x4d, 0x73,
		0xb2, 0x6b, 0xab, 0x1a, 0xae, 0x3c, 0xca, 0x54, 0x99, 0x7c, 0x4f, 0x88, 0xc9, 0x96, 0x70, 0xef,
		0xea, 0x43, 0x4f, 0x30, 0x3e, 0xc6, 0xb6, 0x04, 0x95, 0x71, 0xb6, 0x0d, 0x90, 0x88, 0x2b, 0x42,
		0x2f, 0xde, 0xa0, 0x6a, 0x65, 0xfb, 0xd0, 0x0e, 0xbe, 0xf7, 0x61, 0x58, 0x22, 0x9a, 0xd3, 0x97,
		0x3e, 0xce, 0x1a, 0x32, 0xfb, 0x30, 0xf0, 0xc6, 0xf7, 0xad, 0x37, 0xae, 0xd6, 0xa1, 0x14, 0x8c,
		0x4f, 0x54, 0x00, 0x16, 0xa1, 0x52, 0x8a, 0x34, 0x2b, 0xcd, 0xce, 0x16, 0x69, 0x33, 0x5e, 0x6e,
		0x49, 0x69, 0xd2, 0xee, 0xec, 0xb4, 0xfb, 0x2d, 0x45, 0xde, 0xdf, 0xeb, 0xb7, 0x77, 0x5b, 0x52,
		0x26, 0xd8, 0x57, 0x7f, 0x3b, 0x0d, 0xe5, 0xf0, 0x11, 0x09, 0xfd, 0x20, 0x9c, 0x15, 0xf7, 0x19,
		0x2e, 0xf6, 0x94, 0xbb, 0xba, 0x43, 0xb7, 0xcc, 0x58, 0x65, 0xe5, 0xcb, 0x5f, 0xb4, 0x35, 0xae,
		0xd5, 0xc3, 0xde, 0x0b, 0xba, 0x43, 0x36, 0xc4, 0x58, 0xf5, 0xd0, 0x0e, 0x9c, 0x37, 0x2d, 0xc5,
		0xf5, 0x54, 0x73, 0xa0, 0x3a, 0x03, 0x65, 0x7a, 0x93, 0xa4, 0xa8, 0x9a, 0x86, 0x5d, 0xd7, 0x62,
		0xa5, 0xca, 0x67, 0xf9, 0x88, 0x69, 0xf5, 0xb8, 0xf2, 0x34, 0x87, 0x37, 0xb8, 0x6a, 0x24, 0xc0,
		0x32, 0xc7, 0x05, 0xd8, 0x03, 0x50, 0x18, 0xab, 0xb6, 0x82, 0x4d, 0xcf, 0x39, 0xa2, 0x8d, 0x71,
		0x5e, 0xce, 0x8f, 0x55, 0xbb, 0x45, 0x9e, 0x3f, 0x98, 0xf3, 0xc9, 0xdf, 0x66, 0xa0, 0x14, 0x6c,
		0x8e, 0xc9, 0x59, 0x43, 0xa3, 0x75, 0x24, 0x45, 0x33, 0xcd, 0xc3, 0x27, 0xb6, 0xd2, 0xb5, 0x26,
		0x29, 0x30, 0xf5, 0x05, 0xd6, 0xb2, 0xca, 0x0c, 0x49, 0x8a, 0x3b, 0xc9, 0x2d, 0x98, 0xb5, 0x08,
		0x79, 0x99, 0x3f, 0xa1, 0x6d, 0x58, 0xb8, 0xe5, 0x52, 0xee, 0x05, 0xca, 0xfd, 0xc8, 0xc9, 0xdc,
		0xcf, 0xf7, 0x28, 0x79, 0xe1, 0xf9, 0x9e, 0xb2, 0xd7, 0x91, 0x77, 0x1b, 0x3b, 0x32, 0x87, 0xa3,
		0x73, 0x90, 0x35, 0xd4, 0x7b, 0x47, 0xe1, 0x52, 0x44, 0x45, 0x49, 0x1d, 0x7f, 0x0e, 0xb2, 0x77,
		0xb1, 0x7a, 0x3b, 0x5c, 0x00, 0xa8, 0xe8, 0x7d, 0x0c, 0xfd, 0xcb, 0x90, 0xa3, 0xfe, 0x42, 0x00,
		0xdc, 0x63, 0xd2, 0x7d, 0x28, 0x0f, 0xd9, 0x66, 0x47, 0x26, 0xe1, 0x2f, 0x41, 0x89, 0x49, 0x95,
		0x6e, 0xbb, 0xd5, 0x6c, 0x49, 0xe9, 0xea, 0x55, 0x58, 0x60, 0x4e, 0x20, 0x5b, 0xc3, 0x77, 0x83,
		0x74, 0x1f, 0x7f, 0xe4, 0x1c, 0x29, 0x31, 0xba, 0xbf, 0xbb, 0xd9, 0x92, 0xa5, 0x74, 0x70, 0x79,
		0x5d, 0x28, 0x05, 0xfb, 0xe2, 0x0f, 0x26, 0xa6, 0xbe, 0x95, 0x82, 0x62, 0xa0, 0xcf, 0x25, 0x0d,
		0x8a, 0x6a, 0x18, 0xd6, 0x5d, 0x45, 0x35, 0x74, 0xd5, 0xe5, 0x41, 0x01, 0x54, 0xd4, 0x20, 0x92,
		0xa4, 0x8b, 0xf6, 0x81, 0x18, 0xff, 0x66, 0x0a, 0xa4, 0x68, 0x8b, 0x19, 0x31, 0x30, 0xf5, 0xa1,
		0x1a, 0xf8, 0x46, 0x0a, 0xca, 0xe1, 0xbe, 0x32, 0x62, 0xde, 0xc5, 0x0f, 0xd5, 0xbc, 0xbf, 0x4b,
		0xc3, 0x52, 0xa8, 0x9b, 0x4c, 0x6a, 0xdd, 0x67, 0x61, 0x45, 0x1f, 0xe0, 0xb1, 0x6d, 0x79, 0xd8,
		0xd4, 0x8e, 0x14, 0x03, 0xdf, 0xc1, 0x46, 0xa5, 0x4a, 0x13, 0xc5, 0xe5, 0x93, 0xfb, 0xd5, 0x5a,
		0x7b, 0x8a, 0xdb, 0x21, 0xb0, 0xfa, 0x6a, 0x7b, 0xab, 0xb5, 0xdb, 0xed, 0xf4, 0x5b, 0x7b, 0xcd,
		0x97, 0x94, 0xfd, 0xbd, 0x1f, 0xde, 0xeb, 0xbc, 0xb0, 0x27, 0x4b, 0x7a, 0x44, 0xed, 0x7d, 0xdc,
		0xea, 0x5d, 0x90, 0xa2, 0x46, 0xa1, 0xb3, 0x30, 0xcf, 0x2c, 0xe9, 0x3e, 0xb4, 0x0a, 0xcb, 0x7b,
		0x1d, 0xa5, 0xd7, 0xde, 0x6a, 0x29, 0xad, 0x9b, 0x37, 0x5b, 0xcd, 0x7e, 0x8f, 0xdd, 0x40, 0xf8,
		0xda, 0xfd, 0xf0, 0xa6, 0x7e, 0x3d, 0x03, 0xab, 0x73, 0x2c, 0x41, 0x0d, 0x7e, 0x76, 0x60, 0xc7,
		0x99, 0x27, 0x93, 0x58, 0x5f, 0x23, 0x25, 0xbf, 0xab, 0x3a, 0x1e, 0x3f, 0x6a, 0x3c, 0x0e, 0xc4,
		0x4b, 0xa6, 0xa7, 0x0f, 0x75, 0xec, 0xf0, 0x0b, 0x1b, 0x76, 0xa0, 0x58, 0x9e, 0xca, 0xd9, 0x9d,
		0xcd, 0x0f, 0x00, 0xb2, 0x2d, 0x57, 0xf7, 0xf4, 0x3b, 0x58, 0xd1, 0x4d, 0x71, 0xbb, 0x43, 0x0e,
		0x18, 0x59, 0x59, 0x12, 0x23, 0x6d, 0xd3, 0xf3, 0xb5, 0x4d, 0x3c, 0x52, 0x23, 0xda, 0x24, 0x81,
		0x67, 0x64, 0x49, 0x8c, 0xf8, 0xda, 0x17, 0xa1, 0x34, 0xb0, 0x26, 0xa4, 0xeb, 0x62, 0x7a, 0xa4,
		0x5e, 0xa4, 0xe4, 0x22, 0x93, 0xf9, 0x2a, 0xbc, 0x9f, 0x9e, 0x5e, 0x2b, 0x95, 0xe4, 0x22, 0x93,
		0x31, 0x95, 0xc7, 0x60, 0x59, 0x1d, 0x8d, 0x1c, 0x42, 0x2e, 0x88, 0xd8, 0x09, 0xa1, 0xec, 0x8b,
		0xa9, 0xe2, 0xfa, 0xf3, 0x90, 0x17, 0x7e, 0x20, 0x25, 0x99, 0x78, 0x42, 0xb1, 0xd9, 0xb1, 0x37,
		0xbd, 0x51, 0x90, 0xf3, 0xa6, 0x18, 0xbc, 0x08, 0x25, 0xdd, 0x55, 0xa6, 0xb7, 0xe4, 0xe9, 0x0b,
		0xe9, 0x8d, 0xbc, 0x5c, 0xd4, 0x5d, 0xff, 0x86, 0xb1, 0xfa, 0x56, 0x1a, 0xca, 0xe1, 0x5b, 0x7e,
		0xb4, 0x05, 0x79, 0xc3, 0xd2, 0x54, 0x1a, 0x5a, 0xec, 0x13, 0xd3, 0x46, 0xcc, 0x87, 0x81, 0xda,
		0x0e, 0xd7, 0x97, 0x7d, 0xe4, 0xfa, 0xdf, 0xa4, 0x20, 0x2f, 0xc4, 0xe8, 0x0c, 0x64, 0x6d, 0xd5,
		0x3b, 0xa4, 0x74, 0xb9, 0xcd, 0xb4, 0x94, 0x92, 0xe9, 0x33, 0x91, 0xbb, 0xb6, 0x6a, 0xd2, 0x10,
		0xe0, 0x72, 0xf2, 0x4c, 0xd6, 0xd5, 0xc0, 0xea, 0x80, 0x1e, 0x3f, 0xac, 0xf1, 0x18, 0x9b, 0x9e,
		0x2b, 0xd6, 0x95, 0xcb, 0x9b, 0x5c, 0x8c, 0x9e, 0x80, 0x15, 0xcf, 0x51, 0x75, 0x23, 0xa4, 0x9b,
		0xa5, 0xba, 0x92, 0x18, 0xf0, 0x95, 0xeb, 0x70, 0x4e, 0xf0, 0x0e, 0xb0, 0xa7, 0x6a, 0x87, 0x78,
		0x30, 0x05, 0x2d, 0xd0, 0x6b, 0x86, 0xb3, 0x5c, 0x61, 0x8b, 0x8f, 0x0b, 0x6c, 0xf5, 0x3b, 0x29,
		0x58, 0x11, 0x07, 0xa6, 0x81, 0xef, 0xac, 0x5d, 0x00, 0xd5, 0x34, 0x2d, 0x2f, 0xe8, 0xae, 0xd9,
		0x50, 0x9e, 0xc1, 0xd5, 0x1a, 0x3e, 0x48, 0x0e, 0x10, 0xac, 0x8f, 0x01, 0xa6, 0x23, 0xc7, 0xba,
		0xed, 0x3c, 0x14, 0xf9, 0x27, 0x1c, 0xfa, 0x1d, 0x90, 0x1d, 0xb1, 0x81, 0x89, 0xc8, 0xc9, 0x0a,
		0xad, 0x41, 0xee, 0x00, 0x8f, 0x74, 0x93, 0x5f, 0xcc, 0xb2, 0x07, 0x71, 0x11, 0x92, 0xf5, 0x2f,
		0x42, 0x36, 0x3f, 0x03, 0xab, 0x9a, 0x35, 0x8e, 0x9a, 0xbb, 0x29, 0x45, 0x8e, 0xf9, 0xee, 0x27,
		0x53, 0x2f, 0xc3, 0xb4, 0xc5, 0xfc, 0xdf, 0x54, 0xea, 0xcb, 0xe9, 0xcc, 0x76, 0x77, 0xf3, 0x2b,
		0xe9, 0xf5, 0x6d, 0x06, 0xed, 0x8a, 0x99, 0xca, 0x78, 0x68, 0x60, 0x8d, 0x58, 0x0f, 0xff, 0x74,
		0x09, 0x9e, 0x1c, 0xe9, 0xde, 0xe1, 0xe4, 0xa0, 0xa6, 0x59, 0xe3, 0xcb, 0x23, 0x6b, 0x64, 0x4d,
		0x3f, 0x7d, 0x92, 0x27, 0xfa, 0x40, 0xff, 0xe3, 0x9f, 0x3f, 0x0b, 0xbe, 0x74, 0x3d, 0xf6, 0x5b,
		0x69, 0x7d, 0x0f, 0x56, 0xb9, 0xb2, 0x42, 0xbf, 0xbf, 0xb0, 0x53, 0x04, 0x3a, 0xf1, 0x0e, 0xab,
		0xf2, 0xb5, 0xb7, 0x69, 0xb9, 0x96, 0x57, 0x38, 0x94, 0x8c, 0xb1, 0x83, 0x46, 0x5d, 0x86, 0xfb,
		0x43, 0x7c, 0x6c, 0x6b, 0x62, 0x27, 0x86, 0xf1, 0xdb, 0x9c, 0x71, 0x35, 0xc0, 0xd8, 0xe3, 0xd0,
		0x7a, 0x13, 0x96, 0x4e, 0xc3, 0xf5, 0x97, 0x9c, 0xab, 0x84, 0x83, 0x24, 0xdb, 0xb0, 0x4c, 0x49,
		0xb4, 0x89, 0xeb, 0x59, 0x63, 0x9a, 0xf7, 0x4e, 0xa6, 0xf9, 0xab, 0xb7, 0xd9, 0x5e, 0x29, 0x13,
		0x58, 0xd3, 0x47, 0xd5, 0xeb, 0x40, 0x3f, 0x39, 0x0d, 0xb0, 0x66, 0xc4, 0x30, 0xfc, 0x35, 0x37,
		0xc4, 0xd7, 0xaf, 0x7f, 0x1a, 0xd6, 0xc8, 0xff, 0x34, 0x2d, 0x05, 0x2d, 0x89, 0xbf, 0xf0, 0xaa,
		0x7c, 0xe7, 0x35, 0xb6, 0x1d, 0x57, 0x7d, 0x82, 0x80, 0x4d, 0x81, 0x55, 0x1c, 0x61, 0xcf, 0xc3,
		0x8e, 0xab, 0xa8, 0xc6, 0x3c, 0xf3, 0x02, 0x37, 0x06, 0x95, 0x2f, 0xbc, 0x13, 0x5e, 0xc5, 0x6d,
		0x86, 0x6c, 0x18, 0x46, 0x7d, 0x1f, 0xce, 0xce, 0x89, 0x8a, 0x04, 0x9c, 0xaf, 0x73, 0xce, 0xb5,
		0x99, 0xc8, 0x20, 0xb4, 0x5d, 0x10, 0x72, 0x7f, 0x2d, 0x13, 0x70, 0xfe, 0x2a, 0xe7, 0x44, 0x1c,
		0x2b, 0x96, 0x94, 0x30, 0x3e, 0x0f, 0x2b, 0x77, 0xb0, 0x73, 0x60, 0xb9, 0xfc, 0x96, 0x26, 0x01,
		0xdd, 0x1b, 0x9c, 0x6e, 0x99, 0x03, 0xe9, 0xb5, 0x0d, 0xe1, 0xba, 0x0e, 0xf9, 0xa1, 0xaa, 0xe1,
		0x04, 0x14, 0x5f, 0xe4, 0x14, 0x8b, 0x44, 0x9f, 0x40, 0x1b, 0x50, 0x1a, 0x59, 0xbc, 0x32, 0xc5,
		0xc3, 0xdf, 0xe4, 0xf0, 0xa2, 0xc0, 0x70, 0x0a, 0xdb, 0xb2, 0x27, 0x06, 0x29, 0x5b, 0xf1, 0x14,
		0xbf, 0x26, 0x28, 0x04, 0x86, 0x53, 0x9c, 0xc2, 0xad, 0xbf, 0x2e, 0x28, 0xdc, 0x80, 0x3f, 0x9f,
		0x83, 0xa2, 0x65, 0x1a, 0x47, 0x96, 0x99, 0xc4, 0x88, 0x2f, 0x71, 0x06, 0xe0, 0x10, 0x42, 0x70,
		0x03, 0x0a, 0x49, 0x17, 0xe2, 0x37, 0xdf, 0x11, 0xdb, 0x43, 0xac, 0xc0, 0x36, 0x2c, 0x8b, 0x04,
		0xa5, 0x5b, 0x66, 0x02, 0x8a, 0xdf, 0xe2, 0x14, 0xe5, 0x00, 0x8c, 0x4f, 0xc3, 0xc3, 0xae, 0x37,
		0xc2, 0x49, 0x48, 0xde, 0x12, 0xd3, 0xe0, 0x10, 0xee, 0xca, 0x03, 0x6c, 0x6a, 0x87, 0xc9, 0x18,
		0x7e, 0x5b, 0xb8, 0x52, 0x60, 0x08, 0x45, 0x13, 0x96, 0xc6, 0xaa, 0xe3, 0x1e, 0xaa, 0x46, 0xa2,
		0xe5, 0xf8, 0x1d, 0xce, 0x51, 0xf2, 0x41, 0xdc, 0x23, 0x13, 0xf3, 0x34, 0x34, 0x5f, 0x11, 0x1e,
		0x09, 0xc0, 0xf8, 0xd6, 0x73, 0x3d, 0x7a, 0xa5, 0x75, 0x1a, 0xb6, 0xdf, 0x15, 0x5b, 0x8f, 0x61,
		0x77, 0x83, 0x8c, 0x37, 0xa0, 0xe0, 0xea, 0xf7, 0x12, 0xd1, 0xfc, 0x9e, 0x58, 0x69, 0x0a, 0x20,
		0xe0, 0x97, 0xe0, 0xdc, 0xdc, 0x32, 0x91, 0x80, 0xec, 0xf7, 0x39, 0xd9, 0x99, 0x39, 0xa5, 0x82,
		0xa7, 0x84, 0xd3, 0x52, 0xfe, 0x81, 0x48, 0x09, 0x38, 0xc2, 0xd5, 0x25, 0x67, 0x05, 0x57, 0x1d,
		0x9e, 0xce, 0x6b, 0x7f, 0x28, 0xbc, 0xc6, 0xb0, 0x21, 0xaf, 0xf5, 0xe1, 0x0c, 0x67, 0x3c, 0xdd,
		0xba, 0x7e, 0x55, 0x24, 0x56, 0x86, 0xde, 0x0f, 0xaf, 0xee, 0x67, 0x60, 0xdd, 0x77, 0xa7, 0x68,
		0x4a, 0x5d, 0x65, 0xac, 0xda, 0x09, 0x98, 0xbf, 0xc6, 0x99, 0x45, 0xc6, 0xf7, 0xbb, 0x5a, 0x77,
		0x57, 0xb5, 0x09, 0xf9, 0x8b, 0x50, 0x11, 0xe4, 0x13, 0xd3, 0xc1, 0x9a, 0x35, 0x32, 0xf5, 0x7b,
		0x78, 0x90, 0x80, 0xfa, 0x8f, 0x22, 0x4b, 0xb5, 0x1f, 0x80, 0x13, 0xe6, 0x36, 0x48, 0x7e, 0xaf,
		0xa2, 0xe8, 0x63, 0xdb, 0x72, 0xbc, 0x18, 0xc6, 0xaf, 0x8b, 0x95, 0xf2, 0x71, 0x6d, 0x0a, 0xab,
		0xb7, 0xa0, 0x4c, 0x1f, 0x93, 0x86, 0xe4, 0x1f, 0x73, 0xa2, 0xa5, 0x29, 0x8a, 0x27, 0x0e, 0xcd,
		0x1a, 0xdb, 0xaa, 0x93, 0x24, 0xff, 0xfd, 0x89, 0x48, 0x1c, 0x1c, 0xc2, 0x13, 0x87, 0x77, 0x64,
		0x63, 0x52, 0xed, 0x13, 0x30, 0x7c, 0x43, 0x24, 0x0e, 0x81, 0xe1, 0x14, 0xa2, 0x61, 0x48, 0x40,
		0xf1, 0xa7, 0x82, 0x42, 0x60, 0x08, 0xc5, 0xa7, 0xa6, 0x85, 0xd6, 0xc1, 0x23, 0xdd, 0xf5, 0x1c,
		0xd6, 0x0a, 0x9f, 0x4c, 0xf5, 0x67, 0xef, 0x84, 0x9b, 0x30, 0x39, 0x00, 0x25, 0x99, 0x88, 0x5f,
		0xa1, 0xd2, 0x93, 0x52, 0xbc, 0x61, 0xdf, 0x14, 0x99, 0x28, 0x00, 0x23, 0xb6, 0x05, 0x3a, 0x44,
		0xe2, 0x76, 0x8d, 0x9c, 0x0f, 0x12, 0xd0, 0x7d, 0x2b, 0x62, 0x5c, 0x4f, 0x60, 0x09, 0x67, 0xa0,
		0xff, 0x99, 0x98, 0xb7, 0xf1, 0x51, 0xa2, 0xe8, 0xfc, 0xf3, 0x48, 0xff, 0xb3, 0xcf, 0x90, 0x2c,
		0x87, 0x2c, 0x47, 0xfa, 0x29, 0x14, 0xf7, 0x63, 0x9d, 0xca, 0x8f, 0xbd, 0xcb, 0xe7, 0x1b, 0x6e,
		0xa7, 0xea, 0x3b, 0x24, 0xc8, 0xc3, 0x4d, 0x4f, 0x3c, 0xd9, 0x6b, 0xef, 0xfa, 0x71, 0x1e, 0xea,
		0x79, 0xea, 0x37, 0x61, 0x29, 0xd4, 0xf0, 0xc4, 0x53, 0xfd, 0x38, 0xa7, 0x2a, 0x05, 0xfb, 0x9d,
		0xfa, 0x55, 0xc8, 0x92, 0xe6, 0x25, 0x1e, 0xfe, 0x13, 0x1c, 0x4e, 0xd5, 0xeb, 0x1f, 0x87, 0xbc,
		0x68, 0x5a, 0xe2, 0xa1, 0x3f, 0xc9, 0xa1, 0x3e, 0x84, 0xc0, 0x45, 0xc3, 0x12, 0x0f, 0xff, 0x29,
		0x01, 0x17, 0x10, 0x02, 0x4f, 0xee, 0xc2, 0xbf, 0xf8, 0xe9, 0x2c, 0x2f, 0x3a, 0xc2, 0x77, 0x37,
		0x60, 0x91, 0x77, 0x2a, 0xf1, 0xe8, 0xcf, 0xf1, 0x97, 0x0b, 0x44, 0xfd, 0x19, 0xc8, 0x25, 0x74,
		0xf8, 0xcf, 0x70, 0x28, 0xd3, 0xaf, 0x37, 0xa1, 0x18, 0xe8, 0x4e, 0xe2, 0xe1, 0x3f, 0xcb, 0xe1,
		0x41, 0x14, 0x31, 0x9d, 0x77, 0x27, 0xf1, 0x04, 0x3f, 0x27, 0x4c, 0xe7, 0x08, 0xe2, 0x36, 0xd1,
		0x98, 0xc4, 0xa3, 0x7f, 0x5e, 0x78, 0x5d, 0x40, 0xea, 0xcf, 0x41, 0xc1, 0x2f, 0x36, 0xf1, 0xf8,
		0x5f, 0xe0, 0xf8, 0x29, 0x86, 0x78, 0x20, 0x50, 0xec, 0xe2, 0x29, 0x7e, 0x51, 0x78, 0x20, 0x80,
		0x22, 0xdb, 0x28, 0xda, 0xc0, 0xc4, 0x33, 0xfd, 0x92, 0xd8, 0x46, 0x91, 0xfe, 0x85, 0xac, 0x26,
		0xcd, 0xf9, 0xf1, 0x14, 0xbf, 0x2c, 0x56, 0x93, 0xea, 0x13, 0x33, 0xa2, 0x1d, 0x41, 0x3c, 0xc7,
		0xaf, 0x08, 0x33, 0x22, 0x0d, 0x41, 0xbd, 0x0b, 0x68, 0xb6, 0x1b, 0x88, 0xe7, 0xfb, 0x3c, 0xe7,
		0x5b, 0x99, 0x69, 0x06, 0xea, 0x2f, 0xc0, 0x99, 0xf9, 0x9d, 0x40, 0x3c, 0xeb, 0x17, 0xde, 0x8d,
		0x9c, 0xdd, 0x82, 0x8d, 0x40, 0xbd, 0x3f, 0x2d, 0x29, 0xc1, 0x2e, 0x20, 0x9e, 0xf6, 0xf5, 0x77,
		0xc3, 0x89, 0x3b, 0xd8, 0x04, 0xd4, 0x1b, 0x00, 0xd3, 0x02, 0x1c, 0xcf, 0xf5, 0x06, 0xe7, 0x0a,
		0x80, 0xc8, 0xd6, 0xe0, 0xf5, 0x37, 0x1e, 0xff, 0x45, 0xb1, 0x35, 0x38, 0x82, 0x6c, 0x0d, 0x51,
		0x7a, 0xe3, 0xd1, 0x6f, 0x8a, 0xad, 0x21, 0x20, 0x24, 0xb2, 0x03, 0xd5, 0x2d, 0x9e, 0xe1, 0x4b,
		0x22, 0xb2, 0x03, 0xa8, 0xfa, 0x1e, 0xac, 0xcc, 0x14, 0xc4, 0x78, 0xaa, 0x2f, 0x73, 0x2a, 0x29,
		0x5a, 0x0f, 0x83, 0xc5, 0x8b, 0x17, 0xc3, 0x78, 0xb6, 0xdf, 0x88, 0x14, 0x2f, 0x5e, 0x0b, 0xeb,
		0x37, 0x20, 0x6f, 0x4e, 0x0c, 0x83, 0x6c, 0x1e, 0x74, 0xf2, 0x0f, 0xec, 0x2a, 0xff, 0xfa, 0x1e,
		0xf7, 0x8e, 0x00, 0xd4, 0xaf, 0x42, 0x0e, 0x8f, 0x0f, 0xf0, 0x20, 0x0e, 0xf9, 0x6f, 0xef, 0x89,
		0x84, 0x49, 0xb4, 0xeb, 0xcf, 0x01, 0xb0, 0xab, 0x11, 0xfa, 0xd9, 0x2f, 0x06, 0xfb, 0xef, 0xef,
		0xf1, 0x9f, 0xbe, 0x4c, 0x21, 0x53, 0x02, 0xf6, 0x43, 0x9a, 0x93, 0x09, 0xde, 0x09, 0x13, 0xd0,
		0x15, 0xb9, 0x0e, 0x8b, 0xb7, 0x5c, 0xcb, 0xf4, 0xd4, 0x51, 0x1c, 0xfa, 0x3f, 0x38, 0x5a, 0xe8,
		0x13, 0x87, 0x8d, 0x2d, 0x07, 0x7b, 0xea, 0xc8, 0x8d, 0xc3, 0xfe, 0x27, 0xc7, 0xfa, 0x00, 0x02,
		0xd6, 0x54, 0xd7, 0x4b, 0x32, 0xef, 0xef, 0x0a, 0xb0, 0x00, 0x10, 0xa3, 0xc9, 0xff, 0xb7, 0xf1,
		0x51, 0x1c, 0xf6, 0x7b, 0xc2, 0x68, 0xae, 0x5f, 0xff, 0x38, 0x14, 0xc8, 0xbf, 0xec, 0xf7, 0x6c,
		0x31, 0xe0, 0xff, 0xe2, 0xe0, 0x29, 0x82, 0xbc, 0xd9, 0xf5, 0x06, 0x9e, 0x1e, 0xef, 0xec, 0xff,
		0xe6, 0x2b, 0x2d, 0xf4, 0xeb, 0x0d, 0x28, 0xba, 0xde, 0x60, 0x30, 0xe1, 0xfd, 0x69, 0x0c, 0xfc,
		0x7f, 0xde, 0xf3, 0xaf, 0x2c, 0x7c, 0xcc, 0x66, 0x6b, 0xfe, 0xed, 0x2b, 0x6c, 0x5b, 0xdb, 0x16,
		0xbb, 0x77, 0x7d, 0xb9, 0x1a, 0x7f, 0x81, 0x0a, 0xff, 0x97, 0x02, 0xc9, 0xb3, 0x6c, 0xf7, 0xc8,
		0x9b, 0x38, 0x77, 0x8e, 0xf8, 0x55, 0x6a, 0x76, 0xac, 0xea, 0xe6, 0xfa, 0xe9, 0xee, 0x5f, 0xab,
		0x3b, 0x00, 0x7d, 0x42, 0xd4, 0x27, 0x44, 0xe8, 0x41, 0xc8, 0xf4, 0x2d, 0x9b, 0x7e, 0x3e, 0x2c,
		0x5e, 0x29, 0xd6, 0x08, 0x61, 0x8d, 0x0e, 0xcb, 0x44, 0x8e, 0xce, 0x43, 0xb6, 0x3f, 0x71, 0x86,
		0xfc, 0x07, 0x8e, 0x62, 0x9c, 0x20, 0x65, 0x3a, 0x50, 0x7d, 0x12, 0x72, 0x54, 0x1d, 0x49, 0x53,
		0xa2, 0x0c, 0xc3, 0xa2, 0x00, 0x36, 0x13, 0x50, 0xa7, 0xef, 0x15, 0x83, 0xa9, 0xe9, 0xa0, 0xa0,
		0x48, 0xfb, 0x14, 0x9b, 0xf9, 0xef, 0xfd, 0xfd, 0x43, 0xa9, 0xaf, 0xfe, 0xc3, 0x43, 0xa9, 0xff,
		0x0f, 0x00, 0x00, 0xff, 0xff, 0x41, 0xf2, 0xcf, 0x21, 0x5a, 0x33, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *TopsyTurvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.TopsyTurvy{")
	if this.Top != nil {
		s = append(s, "Top: "+fmt.Sprintf("%#v", this.Top)+",\n")
	}
	if this.Turf != nil {
		s = append(s, "Turf: "+fmt.Sprintf("%#v", this.Turf)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Topsy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Topsy{")
	if this.Top != nil {
		s = append(s, "Top: "+valueToGoStringTopsyturvy(this.Top, "int64")+",\n")
	}
	if this.Turf != nil {
		s = append(s, "Turf: "+valueToGoStringTopsyturvy(this.Turf, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Turvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Turvy{")
	if this.Turf != nil {
		s = append(s, "Turf: "+valueToGoStringTopsyturvy(this.Turf, "int64")+",\n")
	}
	if this.Top != nil {
		s = append(s, "Top: "+valueToGoStringTopsyturvy(this.Top, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTopsyturvy(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("topsyturvy.proto", fileDescriptor_topsyturvy_7bb8c7da1903a704) }

var fileDescriptor_topsyturvy_7bb8c7da1903a704 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc9, 0x2f, 0x28,
	0xae, 0x2c, 0x29, 0x2d, 0x2a, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0xa4, 0xe4, 0xc3, 0xc5, 0x15, 0x02, 0x32, 0x28, 0x04, 0x64, 0x90, 0x90, 0x2c, 0x17, 0x73, 0x48,
	0x7e, 0x81, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xb7, 0x1e, 0xc8, 0x40, 0x3d, 0xb0, 0x74,
	0x10, 0x48, 0x5c, 0x48, 0x9e, 0x8b, 0x25, 0xa4, 0xb4, 0x28, 0x4d, 0x82, 0x09, 0x45, 0x1e, 0xa4,
	0x33, 0x08, 0x2c, 0xa1, 0xa4, 0xcb, 0xc5, 0x0a, 0x56, 0x2e, 0x24, 0x80, 0x30, 0x88, 0x19, 0xa2,
	0x57, 0x08, 0x49, 0x2f, 0x33, 0x92, 0x72, 0xb0, 0xbd, 0x30, 0x49, 0x46, 0x84, 0x24, 0xcc, 0x08,
	0x26, 0xb8, 0x11, 0x4e, 0x1c, 0x1f, 0x1e, 0xca, 0x31, 0x6e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x03, 0xf9, 0xbd, 0x05, 0xfd, 0x00, 0x00, 0x00,
}
