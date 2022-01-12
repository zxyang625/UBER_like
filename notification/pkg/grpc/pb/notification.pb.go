// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

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

type NoticeTripRequest struct {
	TripMsg              *TripMsg `protobuf:"bytes,1,opt,name=trip_msg,json=tripMsg,proto3" json:"trip_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoticeTripRequest) Reset()         { *m = NoticeTripRequest{} }
func (m *NoticeTripRequest) String() string { return proto.CompactTextString(m) }
func (*NoticeTripRequest) ProtoMessage()    {}
func (*NoticeTripRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *NoticeTripRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoticeTripRequest.Unmarshal(m, b)
}
func (m *NoticeTripRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoticeTripRequest.Marshal(b, m, deterministic)
}
func (m *NoticeTripRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoticeTripRequest.Merge(m, src)
}
func (m *NoticeTripRequest) XXX_Size() int {
	return xxx_messageInfo_NoticeTripRequest.Size(m)
}
func (m *NoticeTripRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoticeTripRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoticeTripRequest proto.InternalMessageInfo

func (m *NoticeTripRequest) GetTripMsg() *TripMsg {
	if m != nil {
		return m.TripMsg
	}
	return nil
}

type NoticeTripReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoticeTripReply) Reset()         { *m = NoticeTripReply{} }
func (m *NoticeTripReply) String() string { return proto.CompactTextString(m) }
func (*NoticeTripReply) ProtoMessage()    {}
func (*NoticeTripReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *NoticeTripReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoticeTripReply.Unmarshal(m, b)
}
func (m *NoticeTripReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoticeTripReply.Marshal(b, m, deterministic)
}
func (m *NoticeTripReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoticeTripReply.Merge(m, src)
}
func (m *NoticeTripReply) XXX_Size() int {
	return xxx_messageInfo_NoticeTripReply.Size(m)
}
func (m *NoticeTripReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NoticeTripReply.DiscardUnknown(m)
}

var xxx_messageInfo_NoticeTripReply proto.InternalMessageInfo

func (m *NoticeTripReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *NoticeTripReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*NoticeTripRequest)(nil), "pb.NoticeTripRequest")
	proto.RegisterType((*NoticeTripReply)(nil), "pb.NoticeTripReply")
}

func init() {
	proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07)
}

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x92, 0xe2, 0x2a, 0x29, 0xca, 0x2c, 0x80, 0xf0, 0x95, 0xac, 0xb9, 0x04, 0xfd, 0xf2,
	0x4b, 0x32, 0x93, 0x53, 0x43, 0x8a, 0x32, 0x0b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0xd4, 0xb8, 0x38, 0x40, 0x4a, 0xe2, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0xb8, 0xf5, 0x0a, 0x92, 0xf4, 0x40, 0x4a, 0x7c, 0x8b, 0xd3, 0x83, 0xd8, 0x4b, 0x20, 0x0c, 0x25,
	0x6b, 0x2e, 0x7e, 0x64, 0xcd, 0x05, 0x39, 0x95, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25,
	0xa5, 0xc5, 0x60, 0x8d, 0x1c, 0x41, 0x50, 0x9e, 0x90, 0x00, 0x17, 0x33, 0xc8, 0x34, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x10, 0xd3, 0xc8, 0x83, 0x8b, 0xc7, 0x0f, 0xc9, 0x7d, 0x42, 0x16, 0x5c,
	0x5c, 0x08, 0xc3, 0x84, 0x44, 0x41, 0x16, 0x62, 0xb8, 0x4c, 0x4a, 0x18, 0x5d, 0xb8, 0x20, 0xa7,
	0x32, 0x89, 0x0d, 0xec, 0x15, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xaf, 0x68, 0x9c,
	0xf0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	NoticeTrip(ctx context.Context, in *NoticeTripRequest, opts ...grpc.CallOption) (*NoticeTripReply, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) NoticeTrip(ctx context.Context, in *NoticeTripRequest, opts ...grpc.CallOption) (*NoticeTripReply, error) {
	out := new(NoticeTripReply)
	err := c.cc.Invoke(ctx, "/pb.Notification/NoticeTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	NoticeTrip(context.Context, *NoticeTripRequest) (*NoticeTripReply, error)
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) NoticeTrip(ctx context.Context, req *NoticeTripRequest) (*NoticeTripReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeTrip not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_NoticeTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).NoticeTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Notification/NoticeTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).NoticeTrip(ctx, req.(*NoticeTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoticeTrip",
			Handler:    _Notification_NoticeTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
