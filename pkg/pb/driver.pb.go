// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetDriverInfoRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDriverInfoRequest) Reset()         { *m = GetDriverInfoRequest{} }
func (m *GetDriverInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetDriverInfoRequest) ProtoMessage()    {}
func (*GetDriverInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{0}
}

func (m *GetDriverInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDriverInfoRequest.Unmarshal(m, b)
}
func (m *GetDriverInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDriverInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetDriverInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDriverInfoRequest.Merge(m, src)
}
func (m *GetDriverInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetDriverInfoRequest.Size(m)
}
func (m *GetDriverInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDriverInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDriverInfoRequest proto.InternalMessageInfo

func (m *GetDriverInfoRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetDriverInfoRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetDriverInfoReply struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	AccountNum           int64    `protobuf:"varint,4,opt,name=account_num,json=accountNum,proto3" json:"account_num,omitempty"`
	Asset                float32  `protobuf:"fixed32,5,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDriverInfoReply) Reset()         { *m = GetDriverInfoReply{} }
func (m *GetDriverInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetDriverInfoReply) ProtoMessage()    {}
func (*GetDriverInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{1}
}

func (m *GetDriverInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDriverInfoReply.Unmarshal(m, b)
}
func (m *GetDriverInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDriverInfoReply.Marshal(b, m, deterministic)
}
func (m *GetDriverInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDriverInfoReply.Merge(m, src)
}
func (m *GetDriverInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetDriverInfoReply.Size(m)
}
func (m *GetDriverInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDriverInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDriverInfoReply proto.InternalMessageInfo

func (m *GetDriverInfoReply) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetDriverInfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetDriverInfoReply) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetDriverInfoReply) GetAccountNum() int64 {
	if m != nil {
		return m.AccountNum
	}
	return 0
}

func (m *GetDriverInfoReply) GetAsset() float32 {
	if m != nil {
		return m.Asset
	}
	return 0
}

type TakeOrderRequest struct {
	DriverId             int64    `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	DriverName           string   `protobuf:"bytes,2,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Car                  string   `protobuf:"bytes,4,opt,name=car,proto3" json:"car,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeOrderRequest) Reset()         { *m = TakeOrderRequest{} }
func (m *TakeOrderRequest) String() string { return proto.CompactTextString(m) }
func (*TakeOrderRequest) ProtoMessage()    {}
func (*TakeOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{2}
}

func (m *TakeOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeOrderRequest.Unmarshal(m, b)
}
func (m *TakeOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeOrderRequest.Marshal(b, m, deterministic)
}
func (m *TakeOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeOrderRequest.Merge(m, src)
}
func (m *TakeOrderRequest) XXX_Size() int {
	return xxx_messageInfo_TakeOrderRequest.Size(m)
}
func (m *TakeOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeOrderRequest proto.InternalMessageInfo

func (m *TakeOrderRequest) GetDriverId() int64 {
	if m != nil {
		return m.DriverId
	}
	return 0
}

func (m *TakeOrderRequest) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *TakeOrderRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *TakeOrderRequest) GetCar() string {
	if m != nil {
		return m.Car
	}
	return ""
}

type TakeOrderReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeOrderReply) Reset()         { *m = TakeOrderReply{} }
func (m *TakeOrderReply) String() string { return proto.CompactTextString(m) }
func (*TakeOrderReply) ProtoMessage()    {}
func (*TakeOrderReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{3}
}

func (m *TakeOrderReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeOrderReply.Unmarshal(m, b)
}
func (m *TakeOrderReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeOrderReply.Marshal(b, m, deterministic)
}
func (m *TakeOrderReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeOrderReply.Merge(m, src)
}
func (m *TakeOrderReply) XXX_Size() int {
	return xxx_messageInfo_TakeOrderReply.Size(m)
}
func (m *TakeOrderReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeOrderReply.DiscardUnknown(m)
}

var xxx_messageInfo_TakeOrderReply proto.InternalMessageInfo

func (m *TakeOrderReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *TakeOrderReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDriverInfoRequest)(nil), "pb.GetDriverInfoRequest")
	proto.RegisterType((*GetDriverInfoReply)(nil), "pb.GetDriverInfoReply")
	proto.RegisterType((*TakeOrderRequest)(nil), "pb.TakeOrderRequest")
	proto.RegisterType((*TakeOrderReply)(nil), "pb.TakeOrderReply")
}

func init() {
	proto.RegisterFile("driver.proto", fileDescriptor_521003751d596b5e)
}

var fileDescriptor_521003751d596b5e = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd2, 0x6d, 0x68, 0x86, 0x0f, 0xad, 0x46, 0xd5, 0x6e, 0x14, 0x40, 0x54, 0x39, 0x55,
	0x48, 0x69, 0x24, 0x10, 0x97, 0x3d, 0x23, 0xa1, 0x5e, 0x8a, 0x64, 0xed, 0x7d, 0xe5, 0x26, 0x6e,
	0x14, 0x35, 0xb1, 0x83, 0xed, 0x00, 0x55, 0x55, 0x0e, 0x1c, 0xb9, 0xf2, 0xd3, 0xb8, 0xf0, 0x03,
	0xf8, 0x21, 0xc8, 0x76, 0x52, 0x4a, 0xd5, 0xdb, 0xbc, 0x19, 0xcf, 0xbc, 0x37, 0x6f, 0x0c, 0x4f,
	0x0a, 0x59, 0x7d, 0x66, 0x72, 0xd1, 0x4a, 0xa1, 0x05, 0xfa, 0xed, 0x3a, 0x7e, 0x51, 0x0a, 0x51,
	0xd6, 0x2c, 0xa3, 0x6d, 0x95, 0x51, 0xce, 0x85, 0xa6, 0xba, 0x12, 0x5c, 0xb9, 0x17, 0xc9, 0x0a,
	0xa6, 0x1f, 0x98, 0x7e, 0x6f, 0x9b, 0x96, 0x7c, 0x23, 0x08, 0xfb, 0xd4, 0x31, 0xa5, 0x31, 0x86,
	0x49, 0xa7, 0x98, 0xe4, 0xb4, 0x61, 0x91, 0x37, 0xf3, 0xe6, 0x21, 0x39, 0x62, 0x53, 0x6b, 0xa9,
	0x52, 0x5f, 0x84, 0x2c, 0x22, 0xdf, 0xd5, 0x06, 0x9c, 0xfc, 0xf0, 0x00, 0xcf, 0x06, 0xb6, 0xf5,
	0x0e, 0x6f, 0xe1, 0x91, 0x69, 0x7f, 0xa8, 0x0a, 0x3b, 0x6d, 0x44, 0x02, 0x03, 0x97, 0x05, 0x22,
	0x5c, 0x59, 0x0e, 0x37, 0xc7, 0xc6, 0x78, 0x0d, 0x23, 0x5a, 0xb2, 0x68, 0x34, 0xf3, 0xe6, 0x63,
	0x62, 0x42, 0x7c, 0x05, 0x8f, 0x69, 0x9e, 0x8b, 0x8e, 0xeb, 0x07, 0xde, 0x35, 0xd1, 0x95, 0x1d,
	0x01, 0x7d, 0x6a, 0xd5, 0x35, 0x38, 0x85, 0x31, 0x55, 0x8a, 0xe9, 0x68, 0x3c, 0xf3, 0xe6, 0x3e,
	0x71, 0x20, 0xf9, 0x06, 0xd7, 0xf7, 0x74, 0xcb, 0x3e, 0xca, 0x82, 0xc9, 0x61, 0xb1, 0xe7, 0x10,
	0x3a, 0x8b, 0xfe, 0x69, 0x99, 0xb8, 0xc4, 0xb2, 0x30, 0x3c, 0x7d, 0xf1, 0x44, 0x14, 0xb8, 0xd4,
	0xaa, 0x5f, 0xbd, 0x16, 0xb9, 0x75, 0xd0, 0xea, 0x0b, 0xc9, 0x11, 0x1b, 0xd9, 0x39, 0x95, 0x56,
	0x5c, 0x48, 0x4c, 0x98, 0xdc, 0xc1, 0xb3, 0x13, 0x7e, 0xe3, 0xc3, 0x0d, 0x04, 0x4a, 0x53, 0xdd,
	0x29, 0x4b, 0x3d, 0x21, 0x3d, 0x32, 0xbd, 0x8d, 0x2a, 0x7b, 0x42, 0x13, 0xbe, 0xf9, 0xed, 0x41,
	0xe0, 0x5c, 0xc4, 0xaf, 0xf0, 0xf4, 0x3f, 0x4b, 0x31, 0x5a, 0xb4, 0xeb, 0xc5, 0xa5, 0xb3, 0xc5,
	0x37, 0x17, 0x2a, 0x6d, 0xbd, 0x4b, 0xde, 0x7d, 0xff, 0xf5, 0xe7, 0xa7, 0x9f, 0x61, 0x9a, 0xb9,
	0x65, 0xb2, 0x92, 0xe9, 0xd4, 0x85, 0x69, 0xc5, 0x37, 0x22, 0xdb, 0x0f, 0xd7, 0x3d, 0x64, 0xfb,
	0xe1, 0x98, 0x07, 0xbc, 0x87, 0xf0, 0xb8, 0x00, 0x4e, 0xcd, 0xec, 0x73, 0x3f, 0x63, 0x3c, 0xcb,
	0x1a, 0xb6, 0x97, 0x96, 0xed, 0x36, 0xc1, 0x81, 0x4d, 0xd3, 0x2d, 0x4b, 0x85, 0x79, 0x70, 0xe7,
	0xbd, 0x5e, 0x07, 0xf6, 0xeb, 0xbd, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xe5, 0x40, 0x00,
	0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	GetDriverInfo(ctx context.Context, in *GetDriverInfoRequest, opts ...grpc.CallOption) (*GetDriverInfoReply, error)
	TakeOrder(ctx context.Context, in *TakeOrderRequest, opts ...grpc.CallOption) (*TakeOrderReply, error)
}

type driverClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverClient(cc grpc.ClientConnInterface) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) GetDriverInfo(ctx context.Context, in *GetDriverInfoRequest, opts ...grpc.CallOption) (*GetDriverInfoReply, error) {
	out := new(GetDriverInfoReply)
	err := c.cc.Invoke(ctx, "/pb.Driver/GetDriverInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) TakeOrder(ctx context.Context, in *TakeOrderRequest, opts ...grpc.CallOption) (*TakeOrderReply, error) {
	out := new(TakeOrderReply)
	err := c.cc.Invoke(ctx, "/pb.Driver/TakeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	GetDriverInfo(context.Context, *GetDriverInfoRequest) (*GetDriverInfoReply, error)
	TakeOrder(context.Context, *TakeOrderRequest) (*TakeOrderReply, error)
}

// UnimplementedDriverServer can be embedded to have forward compatible implementations.
type UnimplementedDriverServer struct {
}

func (*UnimplementedDriverServer) GetDriverInfo(ctx context.Context, req *GetDriverInfoRequest) (*GetDriverInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverInfo not implemented")
}
func (*UnimplementedDriverServer) TakeOrder(ctx context.Context, req *TakeOrderRequest) (*TakeOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeOrder not implemented")
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_GetDriverInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDriverInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetDriverInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Driver/GetDriverInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetDriverInfo(ctx, req.(*GetDriverInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_TakeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).TakeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Driver/TakeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).TakeOrder(ctx, req.(*TakeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDriverInfo",
			Handler:    _Driver_GetDriverInfo_Handler,
		},
		{
			MethodName: "TakeOrder",
			Handler:    _Driver_TakeOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}
