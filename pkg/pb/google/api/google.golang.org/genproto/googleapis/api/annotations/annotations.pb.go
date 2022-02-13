// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("annotations.proto", fileDescriptor_ba12aec1634c43e6)
}

var fileDescriptor_ba12aec1634c43e6 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0x94, 0x12, 0x85, 0xb0, 0xf5, 0x13, 0x0b,
	0x32, 0xf5, 0x33, 0x4a, 0x4a, 0x0a, 0x20, 0x4a, 0xa4, 0x14, 0xa0, 0xc2, 0x60, 0x5e, 0x52, 0x69,
	0x9a, 0x7e, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x7e, 0x11, 0x44, 0x85, 0x95, 0x37,
	0x17, 0x0b, 0x48, 0xbd, 0x90, 0x9c, 0x1e, 0xd4, 0x34, 0x98, 0x52, 0x3d, 0xdf, 0xd4, 0x92, 0x8c,
	0xfc, 0x14, 0xff, 0x02, 0xb0, 0x95, 0x12, 0x1b, 0x4e, 0xed, 0x51, 0x52, 0x60, 0xd4, 0xe0, 0x36,
	0x12, 0xd1, 0x43, 0x58, 0xab, 0xe7, 0x51, 0x52, 0x52, 0x10, 0x54, 0x9a, 0x93, 0x1a, 0x04, 0x36,
	0xc4, 0x29, 0x8f, 0x8b, 0x2f, 0x39, 0x3f, 0x17, 0x49, 0x81, 0x93, 0x80, 0x23, 0xc2, 0xd9, 0x01,
	0x20, 0x93, 0x03, 0x18, 0xa3, 0x1c, 0xa1, 0xf2, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9,
	0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x7b, 0xf5, 0x21, 0x52, 0x89, 0x05, 0x99, 0xc5, 0x60,
	0xaf, 0x20, 0x79, 0xda, 0x1a, 0x89, 0xbd, 0x88, 0x89, 0xc5, 0xdd, 0x31, 0xc0, 0x33, 0x89, 0x0d,
	0xac, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xc3, 0xe1, 0xcd, 0x1d, 0x01, 0x00, 0x00,
}
