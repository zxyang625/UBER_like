// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Payment.proto

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

type PayRequest struct {
	BillNum              int64    `protobuf:"varint,1,opt,name=bill_num,json=billNum,proto3" json:"bill_num,omitempty"`
	AccountNum           int64    `protobuf:"varint,2,opt,name=account_num,json=accountNum,proto3" json:"account_num,omitempty"`
	PayPassword          string   `protobuf:"bytes,3,opt,name=pay_password,json=payPassword,proto3" json:"pay_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRequest) Reset()         { *m = PayRequest{} }
func (m *PayRequest) String() string { return proto.CompactTextString(m) }
func (*PayRequest) ProtoMessage()    {}
func (*PayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24d6a3ba68a440e, []int{0}
}

func (m *PayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRequest.Unmarshal(m, b)
}
func (m *PayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRequest.Marshal(b, m, deterministic)
}
func (m *PayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRequest.Merge(m, src)
}
func (m *PayRequest) XXX_Size() int {
	return xxx_messageInfo_PayRequest.Size(m)
}
func (m *PayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayRequest proto.InternalMessageInfo

func (m *PayRequest) GetBillNum() int64 {
	if m != nil {
		return m.BillNum
	}
	return 0
}

func (m *PayRequest) GetAccountNum() int64 {
	if m != nil {
		return m.AccountNum
	}
	return 0
}

func (m *PayRequest) GetPayPassword() string {
	if m != nil {
		return m.PayPassword
	}
	return ""
}

type PayReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayReply) Reset()         { *m = PayReply{} }
func (m *PayReply) String() string { return proto.CompactTextString(m) }
func (*PayReply) ProtoMessage()    {}
func (*PayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f24d6a3ba68a440e, []int{1}
}

func (m *PayReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayReply.Unmarshal(m, b)
}
func (m *PayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayReply.Marshal(b, m, deterministic)
}
func (m *PayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayReply.Merge(m, src)
}
func (m *PayReply) XXX_Size() int {
	return xxx_messageInfo_PayReply.Size(m)
}
func (m *PayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PayReply.DiscardUnknown(m)
}

var xxx_messageInfo_PayReply proto.InternalMessageInfo

func (m *PayReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *PayReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PayRequest)(nil), "pb.PayRequest")
	proto.RegisterType((*PayReply)(nil), "pb.PayReply")
}

func init() {
	proto.RegisterFile("Payment.proto", fileDescriptor_f24d6a3ba68a440e)
}

var fileDescriptor_f24d6a3ba68a440e = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xc6, 0x30,
	0x10, 0x84, 0xe9, 0x1f, 0xf8, 0xff, 0x76, 0x5b, 0x45, 0x72, 0x90, 0xea, 0xc5, 0x5a, 0x2f, 0x3d,
	0xe5, 0xa0, 0x3e, 0x87, 0x84, 0xbc, 0x40, 0xd9, 0xd4, 0x22, 0x62, 0xd2, 0xac, 0x4d, 0x82, 0xe4,
	0xed, 0xa5, 0x21, 0xe8, 0x6d, 0xbf, 0x99, 0x65, 0x67, 0x16, 0xae, 0x24, 0x26, 0xbb, 0x6e, 0x41,
	0xd0, 0xee, 0x82, 0xe3, 0x27, 0xd2, 0xe3, 0x17, 0x80, 0xc4, 0xa4, 0xd6, 0xef, 0xb8, 0xfa, 0xc0,
	0xef, 0xa0, 0xd6, 0x9f, 0xc6, 0xcc, 0x5b, 0xb4, 0x7d, 0x35, 0x54, 0x13, 0x53, 0x97, 0x83, 0xdf,
	0xa2, 0xe5, 0x0f, 0xd0, 0xe2, 0xb2, 0xb8, 0xb8, 0x85, 0xec, 0x9e, 0xb2, 0x0b, 0x45, 0x3a, 0x16,
	0x1e, 0xa1, 0x23, 0x4c, 0x33, 0xa1, 0xf7, 0x3f, 0x6e, 0x7f, 0xef, 0xd9, 0x50, 0x4d, 0x8d, 0x6a,
	0x09, 0x93, 0x2c, 0xd2, 0xf8, 0x0a, 0x75, 0x0e, 0x23, 0x93, 0xf8, 0x2d, 0x9c, 0x7d, 0xc0, 0x10,
	0x7d, 0x0e, 0xaa, 0x55, 0x21, 0x7e, 0x03, 0xcc, 0xfa, 0x8f, 0x7c, 0xbf, 0x51, 0xc7, 0xf8, 0x2c,
	0xe0, 0x52, 0x7a, 0xf3, 0x27, 0x60, 0x12, 0x13, 0xbf, 0x16, 0xa4, 0xc5, 0x7f, 0xed, 0xfb, 0xee,
	0x8f, 0xc9, 0x24, 0x7d, 0xce, 0xdf, 0xbd, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x64, 0x89,
	0x20, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentClient interface {
	Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, "/pb.Payment/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
type PaymentServer interface {
	Pay(context.Context, *PayRequest) (*PayReply, error)
}

// UnimplementedPaymentServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (*UnimplementedPaymentServer) Pay(ctx context.Context, req *PayRequest) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func RegisterPaymentServer(s *grpc.Server, srv PaymentServer) {
	s.RegisterService(&_Payment_serviceDesc, srv)
}

func _Payment_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).Pay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _Payment_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Payment.proto",
}
