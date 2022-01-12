// Code generated by protoc-gen-go. DO NOT EDIT.
// source: billing.proto

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

type BillMsg struct {
	BillNum              int64    `protobuf:"varint,1,opt,name=bill_num,json=billNum,proto3" json:"bill_num,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Origin               string   `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          string   `protobuf:"bytes,6,opt,name=destination,proto3" json:"destination,omitempty"`
	PassengerName        string   `protobuf:"bytes,7,opt,name=passenger_name,json=passengerName,proto3" json:"passenger_name,omitempty"`
	DriverName           string   `protobuf:"bytes,8,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	Payed                bool     `protobuf:"varint,9,opt,name=payed,proto3" json:"payed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillMsg) Reset()         { *m = BillMsg{} }
func (m *BillMsg) String() string { return proto.CompactTextString(m) }
func (*BillMsg) ProtoMessage()    {}
func (*BillMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_958db8ba491a6b57, []int{0}
}

func (m *BillMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillMsg.Unmarshal(m, b)
}
func (m *BillMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillMsg.Marshal(b, m, deterministic)
}
func (m *BillMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillMsg.Merge(m, src)
}
func (m *BillMsg) XXX_Size() int {
	return xxx_messageInfo_BillMsg.Size(m)
}
func (m *BillMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BillMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BillMsg proto.InternalMessageInfo

func (m *BillMsg) GetBillNum() int64 {
	if m != nil {
		return m.BillNum
	}
	return 0
}

func (m *BillMsg) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BillMsg) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *BillMsg) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *BillMsg) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *BillMsg) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *BillMsg) GetPassengerName() string {
	if m != nil {
		return m.PassengerName
	}
	return ""
}

func (m *BillMsg) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *BillMsg) GetPayed() bool {
	if m != nil {
		return m.Payed
	}
	return false
}

type GenBillRequest struct {
	TripMsg              *TripMsg `protobuf:"bytes,1,opt,name=trip_msg,json=tripMsg,proto3" json:"trip_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenBillRequest) Reset()         { *m = GenBillRequest{} }
func (m *GenBillRequest) String() string { return proto.CompactTextString(m) }
func (*GenBillRequest) ProtoMessage()    {}
func (*GenBillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_958db8ba491a6b57, []int{1}
}

func (m *GenBillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenBillRequest.Unmarshal(m, b)
}
func (m *GenBillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenBillRequest.Marshal(b, m, deterministic)
}
func (m *GenBillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenBillRequest.Merge(m, src)
}
func (m *GenBillRequest) XXX_Size() int {
	return xxx_messageInfo_GenBillRequest.Size(m)
}
func (m *GenBillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenBillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenBillRequest proto.InternalMessageInfo

func (m *GenBillRequest) GetTripMsg() *TripMsg {
	if m != nil {
		return m.TripMsg
	}
	return nil
}

type GenBillReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	BillMsg              *BillMsg `protobuf:"bytes,2,opt,name=bill_msg,json=billMsg,proto3" json:"bill_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenBillReply) Reset()         { *m = GenBillReply{} }
func (m *GenBillReply) String() string { return proto.CompactTextString(m) }
func (*GenBillReply) ProtoMessage()    {}
func (*GenBillReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_958db8ba491a6b57, []int{2}
}

func (m *GenBillReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenBillReply.Unmarshal(m, b)
}
func (m *GenBillReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenBillReply.Marshal(b, m, deterministic)
}
func (m *GenBillReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenBillReply.Merge(m, src)
}
func (m *GenBillReply) XXX_Size() int {
	return xxx_messageInfo_GenBillReply.Size(m)
}
func (m *GenBillReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GenBillReply.DiscardUnknown(m)
}

var xxx_messageInfo_GenBillReply proto.InternalMessageInfo

func (m *GenBillReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GenBillReply) GetBillMsg() *BillMsg {
	if m != nil {
		return m.BillMsg
	}
	return nil
}

type GetBillListRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillListRequest) Reset()         { *m = GetBillListRequest{} }
func (m *GetBillListRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillListRequest) ProtoMessage()    {}
func (*GetBillListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_958db8ba491a6b57, []int{3}
}

func (m *GetBillListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillListRequest.Unmarshal(m, b)
}
func (m *GetBillListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillListRequest.Marshal(b, m, deterministic)
}
func (m *GetBillListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillListRequest.Merge(m, src)
}
func (m *GetBillListRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillListRequest.Size(m)
}
func (m *GetBillListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillListRequest proto.InternalMessageInfo

func (m *GetBillListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetBillListReply struct {
	BillMsgList          []*BillMsg `protobuf:"bytes,1,rep,name=bill_msg_list,json=billMsgList,proto3" json:"bill_msg_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetBillListReply) Reset()         { *m = GetBillListReply{} }
func (m *GetBillListReply) String() string { return proto.CompactTextString(m) }
func (*GetBillListReply) ProtoMessage()    {}
func (*GetBillListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_958db8ba491a6b57, []int{4}
}

func (m *GetBillListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillListReply.Unmarshal(m, b)
}
func (m *GetBillListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillListReply.Marshal(b, m, deterministic)
}
func (m *GetBillListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillListReply.Merge(m, src)
}
func (m *GetBillListReply) XXX_Size() int {
	return xxx_messageInfo_GetBillListReply.Size(m)
}
func (m *GetBillListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillListReply proto.InternalMessageInfo

func (m *GetBillListReply) GetBillMsgList() []*BillMsg {
	if m != nil {
		return m.BillMsgList
	}
	return nil
}

func init() {
	proto.RegisterType((*BillMsg)(nil), "pb.BillMsg")
	proto.RegisterType((*GenBillRequest)(nil), "pb.GenBillRequest")
	proto.RegisterType((*GenBillReply)(nil), "pb.GenBillReply")
	proto.RegisterType((*GetBillListRequest)(nil), "pb.GetBillListRequest")
	proto.RegisterType((*GetBillListReply)(nil), "pb.GetBillListReply")
}

func init() {
	proto.RegisterFile("billing.proto", fileDescriptor_958db8ba491a6b57)
}

var fileDescriptor_958db8ba491a6b57 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0x39, 0x7b, 0xc9, 0x4d, 0x6c, 0x29, 0x4b, 0xa9, 0xeb, 0x81, 0x18, 0x02, 0x4a, 0x5e,
	0xbc, 0x42, 0x7d, 0x11, 0x7c, 0xd3, 0x87, 0x22, 0xd8, 0x7b, 0x08, 0x7d, 0x0f, 0x89, 0x19, 0xc2,
	0x40, 0xb2, 0x59, 0x77, 0x37, 0xca, 0xfd, 0x07, 0x7f, 0xb4, 0xcc, 0x6e, 0x5a, 0x7b, 0xe7, 0x5b,
	0xbe, 0x6f, 0xbf, 0xf9, 0x66, 0xe6, 0x9b, 0xc0, 0x79, 0x4b, 0xc3, 0x40, 0xaa, 0xdf, 0x69, 0x33,
	0xb9, 0x49, 0xc4, 0xba, 0xdd, 0x82, 0x33, 0xa4, 0x03, 0x2e, 0xfe, 0xc4, 0x90, 0x7c, 0xa1, 0x61,
	0xb8, 0xb7, 0xbd, 0x78, 0x0d, 0x29, 0x8b, 0x6b, 0x35, 0x8f, 0x32, 0xca, 0xa3, 0x72, 0x55, 0x25,
	0x8c, 0xf7, 0xf3, 0x28, 0xae, 0xe0, 0x4c, 0x1b, 0xfa, 0x81, 0x32, 0xce, 0xa3, 0x32, 0xae, 0x02,
	0x10, 0x6f, 0x00, 0xac, 0x6b, 0x8c, 0xab, 0x1d, 0x8d, 0x28, 0x57, 0xbe, 0x64, 0xe3, 0x99, 0x07,
	0x1a, 0x91, 0xfd, 0x50, 0x75, 0xe1, 0xf1, 0x45, 0xf0, 0x43, 0xd5, 0xf9, 0xa7, 0x6b, 0x58, 0x4f,
	0x86, 0x7a, 0x52, 0xf2, 0x2c, 0x8f, 0xca, 0x4d, 0xb5, 0x20, 0x91, 0x43, 0xd6, 0xa1, 0x75, 0xa4,
	0x1a, 0x47, 0x93, 0x92, 0x6b, 0xff, 0xf8, 0x9c, 0x12, 0xef, 0xe0, 0x42, 0x37, 0xd6, 0xa2, 0xea,
	0xd1, 0xd4, 0xaa, 0x19, 0x51, 0x26, 0x5e, 0x74, 0xfe, 0xc4, 0xee, 0x9b, 0x11, 0xc5, 0x5b, 0xc8,
	0x3a, 0x43, 0xbf, 0x1e, 0x35, 0xa9, 0xd7, 0x40, 0xa0, 0xbc, 0x80, 0x37, 0x6a, 0x0e, 0xd8, 0xc9,
	0x4d, 0x1e, 0x95, 0x69, 0x15, 0x40, 0xf1, 0x09, 0x2e, 0xee, 0x50, 0x71, 0x20, 0x15, 0xfe, 0x9c,
	0xd1, 0x3a, 0xf1, 0x1e, 0x52, 0x8e, 0xab, 0x1e, 0x6d, 0xef, 0x43, 0xc9, 0x6e, 0xb3, 0x9d, 0x6e,
	0x77, 0x0f, 0x86, 0xf4, 0xbd, 0xed, 0xab, 0xc4, 0x85, 0x8f, 0x62, 0x0f, 0x2f, 0x9f, 0x2a, 0xf5,
	0x70, 0xe0, 0x0d, 0xad, 0x6b, 0xdc, 0x6c, 0x7d, 0x55, 0x5a, 0x2d, 0x88, 0xfd, 0x7c, 0xc8, 0xec,
	0x17, 0xff, 0xf3, 0x5b, 0x6e, 0x10, 0x12, 0x67, 0xbf, 0x0f, 0x20, 0xee, 0xd0, 0x31, 0xfd, 0x9d,
	0xac, 0x7b, 0x9c, 0xe6, 0x15, 0x24, 0xb3, 0x45, 0x53, 0x53, 0xb7, 0x5c, 0x68, 0xcd, 0xf0, 0x5b,
	0x57, 0x7c, 0x85, 0xcb, 0x23, 0x39, 0x8f, 0x70, 0x13, 0x8e, 0xcf, 0xad, 0xea, 0x81, 0xac, 0x93,
	0x51, 0xbe, 0x3a, 0xed, 0x97, 0x2d, 0xfd, 0xb8, 0xea, 0xf6, 0x77, 0xf8, 0x17, 0x48, 0xf5, 0xe2,
	0x06, 0x92, 0x65, 0x1d, 0x21, 0x58, 0x7f, 0x9c, 0xca, 0xf6, 0xf2, 0x88, 0xe3, 0x66, 0x9f, 0x21,
	0x7b, 0x36, 0x80, 0xb8, 0x0e, 0x82, 0xd3, 0x05, 0xb6, 0x57, 0xff, 0xf1, 0x7a, 0x38, 0xb4, 0x6b,
	0xff, 0x33, 0x7e, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x85, 0x34, 0x89, 0xad, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BillingClient is the client API for Billing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BillingClient interface {
	GenBill(ctx context.Context, in *GenBillRequest, opts ...grpc.CallOption) (*GenBillReply, error)
	GetBillList(ctx context.Context, in *GetBillListRequest, opts ...grpc.CallOption) (*GetBillListReply, error)
}

type billingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingClient(cc grpc.ClientConnInterface) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) GenBill(ctx context.Context, in *GenBillRequest, opts ...grpc.CallOption) (*GenBillReply, error) {
	out := new(GenBillReply)
	err := c.cc.Invoke(ctx, "/pb.Billing/GenBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetBillList(ctx context.Context, in *GetBillListRequest, opts ...grpc.CallOption) (*GetBillListReply, error) {
	out := new(GetBillListReply)
	err := c.cc.Invoke(ctx, "/pb.Billing/GetBillList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServer is the server API for Billing service.
type BillingServer interface {
	GenBill(context.Context, *GenBillRequest) (*GenBillReply, error)
	GetBillList(context.Context, *GetBillListRequest) (*GetBillListReply, error)
}

// UnimplementedBillingServer can be embedded to have forward compatible implementations.
type UnimplementedBillingServer struct {
}

func (*UnimplementedBillingServer) GenBill(ctx context.Context, req *GenBillRequest) (*GenBillReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenBill not implemented")
}
func (*UnimplementedBillingServer) GetBillList(ctx context.Context, req *GetBillListRequest) (*GetBillListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillList not implemented")
}

func RegisterBillingServer(s *grpc.Server, srv BillingServer) {
	s.RegisterService(&_Billing_serviceDesc, srv)
}

func _Billing_GenBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GenBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Billing/GenBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GenBill(ctx, req.(*GenBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetBillList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetBillList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Billing/GetBillList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetBillList(ctx, req.(*GetBillListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Billing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenBill",
			Handler:    _Billing_GenBill_Handler,
		},
		{
			MethodName: "GetBillList",
			Handler:    _Billing_GetBillList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billing.proto",
}