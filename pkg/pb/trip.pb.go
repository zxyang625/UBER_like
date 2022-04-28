// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trip.proto

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

type GenTripRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenTripRequest) Reset()         { *m = GenTripRequest{} }
func (m *GenTripRequest) String() string { return proto.CompactTextString(m) }
func (*GenTripRequest) ProtoMessage()    {}
func (*GenTripRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31f58919dc2e020, []int{0}
}

func (m *GenTripRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTripRequest.Unmarshal(m, b)
}
func (m *GenTripRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTripRequest.Marshal(b, m, deterministic)
}
func (m *GenTripRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTripRequest.Merge(m, src)
}
func (m *GenTripRequest) XXX_Size() int {
	return xxx_messageInfo_GenTripRequest.Size(m)
}
func (m *GenTripRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTripRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenTripRequest proto.InternalMessageInfo

type GenTripReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Trip                 *TripMsg `protobuf:"bytes,3,opt,name=trip,proto3" json:"trip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenTripReply) Reset()         { *m = GenTripReply{} }
func (m *GenTripReply) String() string { return proto.CompactTextString(m) }
func (*GenTripReply) ProtoMessage()    {}
func (*GenTripReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31f58919dc2e020, []int{1}
}

func (m *GenTripReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTripReply.Unmarshal(m, b)
}
func (m *GenTripReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTripReply.Marshal(b, m, deterministic)
}
func (m *GenTripReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTripReply.Merge(m, src)
}
func (m *GenTripReply) XXX_Size() int {
	return xxx_messageInfo_GenTripReply.Size(m)
}
func (m *GenTripReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTripReply.DiscardUnknown(m)
}

var xxx_messageInfo_GenTripReply proto.InternalMessageInfo

func (m *GenTripReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GenTripReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GenTripReply) GetTrip() *TripMsg {
	if m != nil {
		return m.Trip
	}
	return nil
}

type TripMsg struct {
	TripNum              int64    `protobuf:"varint,1,opt,name=trip_num,json=tripNum,proto3" json:"trip_num,omitempty"`
	PassengerId          int64    `protobuf:"varint,2,opt,name=passenger_id,json=passengerId,proto3" json:"passenger_id,omitempty"`
	DriverId             int64    `protobuf:"varint,3,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	PassengerName        string   `protobuf:"bytes,4,opt,name=passenger_name,json=passengerName,proto3" json:"passenger_name,omitempty"`
	DriverName           string   `protobuf:"bytes,5,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	StartTime            int64    `protobuf:"varint,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Origin               string   `protobuf:"bytes,8,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          string   `protobuf:"bytes,9,opt,name=destination,proto3" json:"destination,omitempty"`
	Car                  string   `protobuf:"bytes,10,opt,name=car,proto3" json:"car,omitempty"`
	Path                 string   `protobuf:"bytes,11,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TripMsg) Reset()         { *m = TripMsg{} }
func (m *TripMsg) String() string { return proto.CompactTextString(m) }
func (*TripMsg) ProtoMessage()    {}
func (*TripMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31f58919dc2e020, []int{2}
}

func (m *TripMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TripMsg.Unmarshal(m, b)
}
func (m *TripMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TripMsg.Marshal(b, m, deterministic)
}
func (m *TripMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TripMsg.Merge(m, src)
}
func (m *TripMsg) XXX_Size() int {
	return xxx_messageInfo_TripMsg.Size(m)
}
func (m *TripMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TripMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TripMsg proto.InternalMessageInfo

func (m *TripMsg) GetTripNum() int64 {
	if m != nil {
		return m.TripNum
	}
	return 0
}

func (m *TripMsg) GetPassengerId() int64 {
	if m != nil {
		return m.PassengerId
	}
	return 0
}

func (m *TripMsg) GetDriverId() int64 {
	if m != nil {
		return m.DriverId
	}
	return 0
}

func (m *TripMsg) GetPassengerName() string {
	if m != nil {
		return m.PassengerName
	}
	return ""
}

func (m *TripMsg) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *TripMsg) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TripMsg) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *TripMsg) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *TripMsg) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *TripMsg) GetCar() string {
	if m != nil {
		return m.Car
	}
	return ""
}

func (m *TripMsg) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*GenTripRequest)(nil), "pb.GenTripRequest")
	proto.RegisterType((*GenTripReply)(nil), "pb.GenTripReply")
	proto.RegisterType((*TripMsg)(nil), "pb.TripMsg")
}

func init() {
	proto.RegisterFile("trip.proto", fileDescriptor_a31f58919dc2e020)
}

var fileDescriptor_a31f58919dc2e020 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4d, 0xaa, 0xdb, 0x30,
	0x10, 0xc7, 0xf1, 0x47, 0xfd, 0x31, 0x4e, 0x43, 0x50, 0x21, 0x28, 0x69, 0x4b, 0x5c, 0x43, 0x21,
	0x14, 0x1a, 0x43, 0xba, 0xeb, 0x05, 0x4a, 0x0a, 0xcd, 0xc2, 0x64, 0xd3, 0x55, 0x50, 0x62, 0xe1,
	0x0a, 0x62, 0x59, 0x95, 0xe4, 0x42, 0xb6, 0xbd, 0x42, 0xef, 0xd4, 0x0b, 0xbc, 0x2b, 0xbc, 0x83,
	0x3c, 0x34, 0x36, 0x79, 0x2f, 0xbb, 0x99, 0xdf, 0x7f, 0x3c, 0xe3, 0xf9, 0x8f, 0x00, 0xac, 0x16,
	0x6a, 0xa3, 0x74, 0x67, 0x3b, 0xe2, 0xab, 0xd3, 0xf2, 0x5d, 0xd3, 0x75, 0xcd, 0x85, 0x97, 0x4c,
	0x89, 0x92, 0x49, 0xd9, 0x59, 0x66, 0x45, 0x27, 0xcd, 0x50, 0x51, 0xcc, 0x60, 0xfa, 0x8d, 0xcb,
	0x83, 0x16, 0xaa, 0xe2, 0xbf, 0x7b, 0x6e, 0x6c, 0xf1, 0x13, 0x26, 0x37, 0xa2, 0x2e, 0x57, 0x32,
	0x87, 0xc8, 0x58, 0x66, 0x7b, 0x43, 0xbd, 0xdc, 0x5b, 0x27, 0xd5, 0x98, 0x91, 0x19, 0x04, 0xad,
	0x69, 0xa8, 0x9f, 0x7b, 0xeb, 0xb4, 0x72, 0x21, 0x59, 0x41, 0xe8, 0x66, 0xd3, 0x20, 0xf7, 0xd6,
	0xd9, 0x36, 0xdb, 0xa8, 0xd3, 0xc6, 0xb5, 0xf9, 0x61, 0x9a, 0x0a, 0x85, 0xe2, 0xbf, 0x0f, 0xf1,
	0x48, 0xc8, 0x02, 0x12, 0xc7, 0x8e, 0xb2, 0x6f, 0xb1, 0x71, 0x50, 0xc5, 0x2e, 0xdf, 0xf7, 0x2d,
	0xf9, 0x00, 0x13, 0xc5, 0x8c, 0xe1, 0xb2, 0xe1, 0xfa, 0x28, 0x6a, 0x1c, 0x11, 0x54, 0xd9, 0x8d,
	0xed, 0x6a, 0xf2, 0x16, 0xd2, 0x5a, 0x8b, 0x3f, 0x83, 0x1e, 0xa0, 0x9e, 0x0c, 0x60, 0x57, 0x93,
	0x8f, 0x30, 0x7d, 0xfe, 0x5e, 0xb2, 0x96, 0xd3, 0x10, 0x7f, 0xf2, 0xf5, 0x8d, 0xee, 0x59, 0xcb,
	0xc9, 0x0a, 0xb2, 0xb1, 0x07, 0xd6, 0xbc, 0xc2, 0x1a, 0x18, 0x10, 0x16, 0xbc, 0x07, 0x30, 0x96,
	0x69, 0x7b, 0xb4, 0xa2, 0xe5, 0x34, 0xc2, 0x29, 0x29, 0x92, 0x83, 0x68, 0xb9, 0xdb, 0x80, 0xcb,
	0x7a, 0x10, 0xe3, 0x61, 0x03, 0x2e, 0x6b, 0x94, 0xe6, 0x10, 0x75, 0x5a, 0x34, 0x42, 0xd2, 0x04,
	0xbb, 0x8e, 0x19, 0xc9, 0x21, 0xab, 0xb9, 0xb1, 0x42, 0xe2, 0x0d, 0x68, 0x8a, 0xe2, 0x4b, 0xe4,
	0x5c, 0x3d, 0x33, 0x4d, 0x61, 0x70, 0xf5, 0xcc, 0x34, 0x21, 0x10, 0x2a, 0x66, 0x7f, 0xd1, 0x0c,
	0x11, 0xc6, 0xdb, 0x0a, 0x42, 0xe7, 0x23, 0xf9, 0x0e, 0xf1, 0x78, 0x2b, 0x42, 0x9c, 0xdd, 0xf7,
	0xa7, 0x5c, 0xce, 0xee, 0x98, 0xba, 0x5c, 0x8b, 0xc5, 0xdf, 0x87, 0xc7, 0x7f, 0xfe, 0x9b, 0x62,
	0x5a, 0x3a, 0xb3, 0xcb, 0x86, 0xcb, 0xcf, 0x2e, 0xf8, 0xea, 0x7d, 0x3a, 0x45, 0xf8, 0x20, 0xbe,
	0x3c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xae, 0xc5, 0xe6, 0x40, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TripClient is the client API for Trip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TripClient interface {
	GenTrip(ctx context.Context, in *GenTripRequest, opts ...grpc.CallOption) (*GenTripReply, error)
}

type tripClient struct {
	cc grpc.ClientConnInterface
}

func NewTripClient(cc grpc.ClientConnInterface) TripClient {
	return &tripClient{cc}
}

func (c *tripClient) GenTrip(ctx context.Context, in *GenTripRequest, opts ...grpc.CallOption) (*GenTripReply, error) {
	out := new(GenTripReply)
	err := c.cc.Invoke(ctx, "/pb.Trip/GenTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServer is the server API for Trip service.
type TripServer interface {
	GenTrip(context.Context, *GenTripRequest) (*GenTripReply, error)
}

// UnimplementedTripServer can be embedded to have forward compatible implementations.
type UnimplementedTripServer struct {
}

func (*UnimplementedTripServer) GenTrip(ctx context.Context, req *GenTripRequest) (*GenTripReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenTrip not implemented")
}

func RegisterTripServer(s *grpc.Server, srv TripServer) {
	s.RegisterService(&_Trip_serviceDesc, srv)
}

func _Trip_GenTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServer).GenTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Trip/GenTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServer).GenTrip(ctx, req.(*GenTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Trip",
	HandlerType: (*TripServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenTrip",
			Handler:    _Trip_GenTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trip.proto",
}
