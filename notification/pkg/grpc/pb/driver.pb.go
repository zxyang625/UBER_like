// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	PassengerName        string   `protobuf:"bytes,2,opt,name=passenger_name,json=passengerName,proto3" json:"passenger_name,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Origin               string   `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
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

func (m *TakeOrderReply) GetPassengerName() string {
	if m != nil {
		return m.PassengerName
	}
	return ""
}

func (m *TakeOrderReply) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TakeOrderReply) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *TakeOrderReply) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *TakeOrderReply) GetPath() string {
	if m != nil {
		return m.Path
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
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xa6, 0xeb, 0x5a, 0xd7, 0x37, 0x37, 0x46, 0x18, 0xb3, 0x4c, 0xc4, 0x52, 0x10, 0x76, 0xda,
	0x41, 0xf1, 0x07, 0x08, 0x82, 0xec, 0x32, 0x21, 0xec, 0x3e, 0xb2, 0x36, 0xce, 0xe0, 0x9a, 0xd4,
	0x24, 0x55, 0x76, 0xf1, 0xe0, 0xd1, 0xdf, 0xe1, 0x0f, 0x95, 0x24, 0x5d, 0xdd, 0xc6, 0x6e, 0xef,
	0xfb, 0x5e, 0xfb, 0xbe, 0xef, 0x7d, 0x2f, 0x70, 0x9e, 0x4b, 0xf6, 0x41, 0xe5, 0xb4, 0x94, 0x42,
	0x0b, 0xd4, 0x2a, 0x57, 0xe9, 0x1c, 0x86, 0x4f, 0x54, 0x3f, 0x5a, 0x7a, 0xc6, 0x5f, 0x04, 0xa6,
	0xef, 0x15, 0x55, 0x1a, 0x8d, 0xa1, 0x53, 0x29, 0x2a, 0x39, 0x29, 0x68, 0xec, 0x25, 0xde, 0x24,
	0xc2, 0x0d, 0x36, 0xbd, 0x92, 0x28, 0xf5, 0x29, 0x64, 0x1e, 0xb7, 0x5c, 0x6f, 0x87, 0xd3, 0x1f,
	0x0f, 0xd0, 0xd1, 0xc0, 0x72, 0xb3, 0x45, 0x17, 0x70, 0x66, 0x7e, 0x5f, 0xb2, 0xdc, 0x4e, 0xf3,
	0x71, 0x68, 0xe0, 0x2c, 0x47, 0x08, 0xda, 0x56, 0xc3, 0xcd, 0xb1, 0x35, 0x1a, 0x80, 0x4f, 0xd6,
	0x34, 0xf6, 0x13, 0x6f, 0x12, 0x60, 0x53, 0xa2, 0x6b, 0xe8, 0x92, 0x2c, 0x13, 0x15, 0xd7, 0x4b,
	0x5e, 0x15, 0x71, 0xdb, 0x8e, 0x80, 0x9a, 0x9a, 0x57, 0x05, 0x1a, 0x42, 0x40, 0x94, 0xa2, 0x3a,
	0x0e, 0x12, 0x6f, 0xd2, 0xc2, 0x0e, 0xa4, 0x5f, 0x30, 0x58, 0x90, 0x37, 0xfa, 0x2c, 0x73, 0x2a,
	0x77, 0x8b, 0x5d, 0x42, 0xe4, 0x42, 0xf8, 0xf7, 0xd2, 0x71, 0xc4, 0x2c, 0x37, 0x3a, 0x75, 0x73,
	0xcf, 0x14, 0x38, 0x6a, 0x5e, 0xaf, 0xbe, 0x11, 0x19, 0xd1, 0x4c, 0x70, 0xeb, 0x2f, 0xc2, 0x0d,
	0x36, 0xb6, 0x33, 0x22, 0xad, 0xb9, 0x08, 0x9b, 0x32, 0xfd, 0xf5, 0xa0, 0xbf, 0x67, 0xc0, 0x04,
	0x71, 0x03, 0x7d, 0x93, 0x15, 0xe5, 0xeb, 0x43, 0x91, 0x5e, 0xc3, 0x5a, 0x9d, 0x2b, 0x00, 0xa5,
	0x89, 0xd4, 0x4b, 0xcd, 0x0a, 0x97, 0x84, 0x8f, 0x23, 0xcb, 0x2c, 0x58, 0x41, 0xd1, 0x08, 0x42,
	0x21, 0xd9, 0x9a, 0xf1, 0x5a, 0xad, 0x46, 0x28, 0x81, 0x6e, 0x4e, 0x95, 0x66, 0xdc, 0x39, 0x0c,
	0x6c, 0x73, 0x9f, 0x32, 0x79, 0x97, 0x44, 0xbf, 0xc6, 0xa1, 0xcb, 0xdb, 0xd4, 0xb7, 0xdf, 0x1e,
	0x84, 0xee, 0x60, 0xe8, 0x01, 0x7a, 0x07, 0xd7, 0x43, 0xf1, 0xb4, 0x5c, 0x4d, 0x4f, 0xbd, 0x90,
	0xf1, 0xe8, 0x44, 0xc7, 0x6c, 0x78, 0x0f, 0x51, 0xb3, 0x33, 0x1a, 0x9a, 0x8f, 0x8e, 0x6f, 0x30,
	0x46, 0x47, 0x6c, 0xb9, 0xd9, 0xae, 0x42, 0xfb, 0x26, 0xef, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xee, 0x64, 0x59, 0xa3, 0x02, 0x00, 0x00,
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
