// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order-service.proto

package order_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type OrderRequest struct {
	Item                 string   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_580508db550b68a7, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

func (m *OrderRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type OrderResponse struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Item                 string               `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Quantity             int32                `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_580508db550b68a7, []int{1}
}

func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (m *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(m, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func (m *OrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderResponse) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

func (m *OrderResponse) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "OrderRequest")
	proto.RegisterType((*OrderResponse)(nil), "OrderResponse")
}

func init() {
	proto.RegisterFile("proto/order-service.proto", fileDescriptor_580508db550b68a7)
}

var fileDescriptor_580508db550b68a7 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x65, 0xd3, 0xa2, 0xf6, 0xa0, 0x1d, 0x3c, 0x05, 0x2f, 0x44, 0x9d, 0xc2, 0x80, 0x23,
	0x85, 0x89, 0x05, 0xa9, 0x4f, 0x80, 0x64, 0xb1, 0x23, 0xb7, 0x3e, 0x2a, 0x4b, 0xa4, 0x4e, 0xec,
	0x0b, 0x12, 0xaf, 0xc0, 0x53, 0x23, 0x9c, 0x3f, 0x22, 0x43, 0xb7, 0xbb, 0xd3, 0x7d, 0x3f, 0x7d,
	0x3f, 0xb8, 0x6b, 0x82, 0x27, 0x5f, 0xfa, 0x60, 0x31, 0x3c, 0x46, 0x0c, 0x5f, 0xee, 0x88, 0x2a,
	0xdd, 0xe4, 0xfd, 0xc9, 0xfb, 0xd3, 0x27, 0x96, 0x69, 0x3b, 0x74, 0x1f, 0x25, 0xb9, 0x1a, 0x23,
	0x99, 0xba, 0xe9, 0x1f, 0x76, 0x2f, 0x70, 0xfb, 0xfa, 0x97, 0xd3, 0xd8, 0x76, 0x18, 0x49, 0x08,
	0x58, 0x38, 0xc2, 0x3a, 0x63, 0x39, 0x2b, 0xd6, 0x3a, 0xcd, 0x42, 0xc2, 0xaa, 0xed, 0xcc, 0x99,
	0x1c, 0x7d, 0x67, 0x3c, 0x67, 0xc5, 0x52, 0x4f, 0xfb, 0xee, 0x87, 0xc1, 0x66, 0x00, 0xc4, 0xc6,
	0x9f, 0x23, 0x8a, 0x2d, 0x70, 0x67, 0x87, 0x3c, 0x77, 0x76, 0x22, 0xf2, 0x0b, 0xc4, 0xab, 0x39,
	0x51, 0x3c, 0x03, 0x1c, 0x03, 0x1a, 0x42, 0xfb, 0x6e, 0x28, 0x5b, 0xe4, 0xac, 0xb8, 0xa9, 0xa4,
	0xea, 0x3d, 0xd4, 0xe8, 0xa1, 0xde, 0x46, 0x0f, 0xbd, 0x1e, 0xbe, 0xf7, 0x54, 0x55, 0xb0, 0x4c,
	0x5d, 0xc4, 0x03, 0xac, 0xf6, 0xd6, 0xf6, 0xf3, 0x46, 0xfd, 0x17, 0x94, 0x5b, 0x35, 0xab, 0x7b,
	0xb8, 0x4e, 0xc8, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x28, 0xa7, 0xf2, 0x45, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	AddOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) AddOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/Order/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	AddOrder(context.Context, *OrderRequest) (*OrderResponse, error)
}

// UnimplementedOrderServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (*UnimplementedOrderServer) AddOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _Order_AddOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order-service.proto",
}