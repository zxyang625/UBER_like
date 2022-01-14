// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passenger.proto

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

type GetPassengerInfoRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPassengerInfoRequest) Reset()         { *m = GetPassengerInfoRequest{} }
func (m *GetPassengerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPassengerInfoRequest) ProtoMessage()    {}
func (*GetPassengerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59870f8f4446d21a, []int{0}
}

func (m *GetPassengerInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerInfoRequest.Unmarshal(m, b)
}
func (m *GetPassengerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPassengerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerInfoRequest.Merge(m, src)
}
func (m *GetPassengerInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPassengerInfoRequest.Size(m)
}
func (m *GetPassengerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerInfoRequest proto.InternalMessageInfo

func (m *GetPassengerInfoRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetPassengerInfoRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetPassengerInfoReply struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	AccountNum           int64    `protobuf:"varint,4,opt,name=account_num,json=accountNum,proto3" json:"account_num,omitempty"`
	Asset                float32  `protobuf:"fixed32,5,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPassengerInfoReply) Reset()         { *m = GetPassengerInfoReply{} }
func (m *GetPassengerInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetPassengerInfoReply) ProtoMessage()    {}
func (*GetPassengerInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_59870f8f4446d21a, []int{1}
}

func (m *GetPassengerInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerInfoReply.Unmarshal(m, b)
}
func (m *GetPassengerInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerInfoReply.Marshal(b, m, deterministic)
}
func (m *GetPassengerInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerInfoReply.Merge(m, src)
}
func (m *GetPassengerInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetPassengerInfoReply.Size(m)
}
func (m *GetPassengerInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerInfoReply proto.InternalMessageInfo

func (m *GetPassengerInfoReply) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetPassengerInfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPassengerInfoReply) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetPassengerInfoReply) GetAccountNum() int64 {
	if m != nil {
		return m.AccountNum
	}
	return 0
}

func (m *GetPassengerInfoReply) GetAsset() float32 {
	if m != nil {
		return m.Asset
	}
	return 0
}

type PublishOrderRequest struct {
	PassengerId          int64    `protobuf:"varint,1,opt,name=passenger_id,json=passengerId,proto3" json:"passenger_id,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          string   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	PassengerName        string   `protobuf:"bytes,5,opt,name=passenger_name,json=passengerName,proto3" json:"passenger_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishOrderRequest) Reset()         { *m = PublishOrderRequest{} }
func (m *PublishOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PublishOrderRequest) ProtoMessage()    {}
func (*PublishOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59870f8f4446d21a, []int{2}
}

func (m *PublishOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishOrderRequest.Unmarshal(m, b)
}
func (m *PublishOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishOrderRequest.Marshal(b, m, deterministic)
}
func (m *PublishOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishOrderRequest.Merge(m, src)
}
func (m *PublishOrderRequest) XXX_Size() int {
	return xxx_messageInfo_PublishOrderRequest.Size(m)
}
func (m *PublishOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishOrderRequest proto.InternalMessageInfo

func (m *PublishOrderRequest) GetPassengerId() int64 {
	if m != nil {
		return m.PassengerId
	}
	return 0
}

func (m *PublishOrderRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *PublishOrderRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *PublishOrderRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *PublishOrderRequest) GetPassengerName() string {
	if m != nil {
		return m.PassengerName
	}
	return ""
}

type PublishOrderReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	DriverName           string   `protobuf:"bytes,2,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Car                  string   `protobuf:"bytes,4,opt,name=car,proto3" json:"car,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishOrderReply) Reset()         { *m = PublishOrderReply{} }
func (m *PublishOrderReply) String() string { return proto.CompactTextString(m) }
func (*PublishOrderReply) ProtoMessage()    {}
func (*PublishOrderReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_59870f8f4446d21a, []int{3}
}

func (m *PublishOrderReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishOrderReply.Unmarshal(m, b)
}
func (m *PublishOrderReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishOrderReply.Marshal(b, m, deterministic)
}
func (m *PublishOrderReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishOrderReply.Merge(m, src)
}
func (m *PublishOrderReply) XXX_Size() int {
	return xxx_messageInfo_PublishOrderReply.Size(m)
}
func (m *PublishOrderReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishOrderReply.DiscardUnknown(m)
}

var xxx_messageInfo_PublishOrderReply proto.InternalMessageInfo

func (m *PublishOrderReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *PublishOrderReply) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *PublishOrderReply) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *PublishOrderReply) GetCar() string {
	if m != nil {
		return m.Car
	}
	return ""
}

func (m *PublishOrderReply) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPassengerInfoRequest)(nil), "pb.GetPassengerInfoRequest")
	proto.RegisterType((*GetPassengerInfoReply)(nil), "pb.GetPassengerInfoReply")
	proto.RegisterType((*PublishOrderRequest)(nil), "pb.PublishOrderRequest")
	proto.RegisterType((*PublishOrderReply)(nil), "pb.PublishOrderReply")
}

func init() {
	proto.RegisterFile("passenger.proto", fileDescriptor_59870f8f4446d21a)
}

var fileDescriptor_59870f8f4446d21a = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0xce, 0xd3, 0x30,
	0x10, 0x94, 0x9b, 0x26, 0x34, 0xdb, 0x02, 0xc5, 0xd0, 0x36, 0x14, 0x21, 0x42, 0x24, 0xa4, 0x9e,
	0x7a, 0x80, 0x3b, 0x57, 0x54, 0x0e, 0xa5, 0x58, 0xdc, 0x2b, 0x27, 0x31, 0xad, 0xa5, 0xc4, 0x0e,
	0xb6, 0x03, 0xea, 0x43, 0xc0, 0x95, 0x57, 0xe1, 0xf1, 0x90, 0x9d, 0x9f, 0xf6, 0xfb, 0xda, 0xef,
	0xb6, 0x33, 0x6b, 0xcd, 0xce, 0xee, 0x18, 0x9e, 0x56, 0x54, 0x6b, 0x26, 0x0e, 0x4c, 0xad, 0x2b,
	0x25, 0x8d, 0xc4, 0x83, 0x2a, 0x4d, 0xbe, 0xc2, 0xe2, 0x13, 0x33, 0xbb, 0xae, 0xb3, 0x11, 0xdf,
	0x25, 0x61, 0x3f, 0x6a, 0xa6, 0x0d, 0x5e, 0xc2, 0xa8, 0xd6, 0x4c, 0x09, 0x5a, 0xb2, 0x08, 0xc5,
	0x68, 0x15, 0x92, 0x1e, 0xdb, 0x9e, 0x55, 0xfb, 0x25, 0x55, 0x1e, 0x0d, 0x9a, 0x5e, 0x87, 0x93,
	0xdf, 0x08, 0x66, 0xd7, 0x9a, 0x55, 0x71, 0xc2, 0x0b, 0x78, 0x64, 0x15, 0xf6, 0x3c, 0x77, 0x82,
	0x1e, 0x09, 0x2c, 0xdc, 0xe4, 0x18, 0xc3, 0xd0, 0x8d, 0x69, 0xa4, 0x5c, 0x8d, 0xa7, 0xe0, 0xd1,
	0x03, 0x8b, 0xbc, 0x18, 0xad, 0x7c, 0x62, 0x4b, 0xfc, 0x06, 0xc6, 0x34, 0xcb, 0x64, 0x2d, 0xcc,
	0x5e, 0xd4, 0x65, 0x34, 0x74, 0x12, 0xd0, 0x52, 0xdb, 0xba, 0xc4, 0x2f, 0xc0, 0xb7, 0x43, 0x4d,
	0xe4, 0xc7, 0x68, 0x35, 0x20, 0x0d, 0x48, 0xfe, 0x21, 0x78, 0xbe, 0xab, 0xd3, 0x82, 0xeb, 0xe3,
	0x17, 0x95, 0x33, 0xd5, 0xed, 0xf7, 0x16, 0x26, 0xfd, 0x45, 0xce, 0x96, 0xc6, 0x3d, 0xb7, 0xc9,
	0xf1, 0x6b, 0x00, 0x6d, 0xa8, 0x32, 0x7b, 0xc3, 0x5b, 0x77, 0x1e, 0x09, 0x1d, 0xf3, 0x8d, 0x97,
	0x0c, 0xcf, 0x21, 0x90, 0x8a, 0x1f, 0xb8, 0x70, 0x2e, 0x43, 0xd2, 0x22, 0x1c, 0xc3, 0x38, 0x67,
	0xda, 0x70, 0x41, 0x0d, 0x97, 0xc2, 0x19, 0x0d, 0xc9, 0x25, 0x85, 0xdf, 0xc1, 0x93, 0xf3, 0x6c,
	0xb7, 0xba, 0xef, 0x1e, 0x3d, 0xee, 0xd9, 0x2d, 0x2d, 0x59, 0xf2, 0x07, 0xc1, 0xb3, 0xbb, 0xd6,
	0xed, 0x19, 0xe7, 0x10, 0x68, 0x43, 0x4d, 0xad, 0x9d, 0xe5, 0x11, 0x69, 0x91, 0xbd, 0x4f, 0xae,
	0xf8, 0xcf, 0x4e, 0xb1, 0x39, 0x26, 0x34, 0xd4, 0xb6, 0x4d, 0xad, 0x90, 0x59, 0x63, 0xaa, 0x71,
	0xdc, 0x63, 0x7b, 0xee, 0x8c, 0xaa, 0xd6, 0xab, 0x2d, 0x6d, 0x28, 0x15, 0x35, 0xc7, 0xd6, 0x99,
	0xab, 0xdf, 0xff, 0x45, 0x10, 0xf6, 0xc1, 0xe2, 0xcf, 0x30, 0xbd, 0x1f, 0x34, 0x7e, 0xb5, 0xae,
	0xd2, 0xf5, 0x03, 0x5f, 0x6a, 0xf9, 0xf2, 0x76, 0xd3, 0x2e, 0xf5, 0x11, 0x26, 0x97, 0x9b, 0xe2,
	0x85, 0x7d, 0x7a, 0x23, 0xb6, 0xe5, 0xec, 0xba, 0x51, 0x15, 0xa7, 0x34, 0x70, 0x7f, 0xfa, 0xc3,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x28, 0x03, 0xcd, 0xe6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PassengerClient is the client API for Passenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerClient interface {
	GetPassengerInfo(ctx context.Context, in *GetPassengerInfoRequest, opts ...grpc.CallOption) (*GetPassengerInfoReply, error)
	PublishOrder(ctx context.Context, in *PublishOrderRequest, opts ...grpc.CallOption) (*PublishOrderReply, error)
}

type passengerClient struct {
	cc grpc.ClientConnInterface
}

func NewPassengerClient(cc grpc.ClientConnInterface) PassengerClient {
	return &passengerClient{cc}
}

func (c *passengerClient) GetPassengerInfo(ctx context.Context, in *GetPassengerInfoRequest, opts ...grpc.CallOption) (*GetPassengerInfoReply, error) {
	out := new(GetPassengerInfoReply)
	err := c.cc.Invoke(ctx, "/pb.Passenger/GetPassengerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerClient) PublishOrder(ctx context.Context, in *PublishOrderRequest, opts ...grpc.CallOption) (*PublishOrderReply, error) {
	out := new(PublishOrderReply)
	err := c.cc.Invoke(ctx, "/pb.Passenger/PublishOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerServer is the server API for Passenger service.
type PassengerServer interface {
	GetPassengerInfo(context.Context, *GetPassengerInfoRequest) (*GetPassengerInfoReply, error)
	PublishOrder(context.Context, *PublishOrderRequest) (*PublishOrderReply, error)
}

// UnimplementedPassengerServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerServer struct {
}

func (*UnimplementedPassengerServer) GetPassengerInfo(ctx context.Context, req *GetPassengerInfoRequest) (*GetPassengerInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassengerInfo not implemented")
}
func (*UnimplementedPassengerServer) PublishOrder(ctx context.Context, req *PublishOrderRequest) (*PublishOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishOrder not implemented")
}

func RegisterPassengerServer(s *grpc.Server, srv PassengerServer) {
	s.RegisterService(&_Passenger_serviceDesc, srv)
}

func _Passenger_GetPassengerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPassengerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServer).GetPassengerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passenger/GetPassengerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServer).GetPassengerInfo(ctx, req.(*GetPassengerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passenger_PublishOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServer).PublishOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passenger/PublishOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServer).PublishOrder(ctx, req.(*PublishOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Passenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Passenger",
	HandlerType: (*PassengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPassengerInfo",
			Handler:    _Passenger_GetPassengerInfo_Handler,
		},
		{
			MethodName: "PublishOrder",
			Handler:    _Passenger_PublishOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passenger.proto",
}