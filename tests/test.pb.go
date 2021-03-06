// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package tests

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_4837443f5277a657, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_4837443f5277a657, []int{1}
}
func (m *EmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMessage.Unmarshal(m, b)
}
func (m *EmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMessage.Marshal(b, m, deterministic)
}
func (dst *EmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMessage.Merge(dst, src)
}
func (m *EmptyMessage) XXX_Size() int {
	return xxx_messageInfo_EmptyMessage.Size(m)
}
func (m *EmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "service.Message")
	proto.RegisterType((*EmptyMessage)(nil), "service.EmptyMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Throw(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Die(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Info(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/service.Test/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Throw(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/service.Test/Throw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Die(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/service.Test/Die", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Info(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/service.Test/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/service.Test/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Echo(context.Context, *Message) (*Message, error)
	Throw(context.Context, *Message) (*Message, error)
	Die(context.Context, *Message) (*Message, error)
	Info(context.Context, *Message) (*Message, error)
	Ping(context.Context, *EmptyMessage) (*EmptyMessage, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Throw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Throw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/Throw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Throw(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Die_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Die(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/Die",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Die(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Info(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Ping(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Test_Echo_Handler,
		},
		{
			MethodName: "Throw",
			Handler:    _Test_Throw_Handler,
		},
		{
			MethodName: "Die",
			Handler:    _Test_Die_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Test_Info_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Test_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_4837443f5277a657) }

var fileDescriptor_test_4837443f5277a657 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55,
	0x92, 0xe6, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x2d,
	0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0xf8, 0xb8, 0x78, 0x5c, 0x73,
	0x0b, 0x4a, 0x2a, 0xa1, 0x2a, 0x8c, 0xbe, 0x31, 0x72, 0xb1, 0x84, 0xa4, 0x16, 0x97, 0x08, 0xe9,
	0x70, 0xb1, 0xb8, 0x26, 0x67, 0xe4, 0x0b, 0x09, 0xe8, 0x41, 0xcd, 0xd1, 0x83, 0x2a, 0x91, 0xc2,
	0x10, 0x51, 0x62, 0x10, 0xd2, 0xe5, 0x62, 0x0d, 0xc9, 0x28, 0xca, 0x2f, 0x27, 0x52, 0xb9, 0x36,
	0x17, 0xb3, 0x4b, 0x66, 0x2a, 0x91, 0x8a, 0x75, 0xb8, 0x58, 0x3c, 0xf3, 0xd2, 0x88, 0x75, 0x89,
	0x19, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x90, 0x28, 0x5c, 0x0e, 0xd9, 0x7f, 0x52, 0xd8, 0x85,
	0x95, 0x18, 0x92, 0xd8, 0xc0, 0xa1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x27, 0x6e, 0xf5,
	0x51, 0x43, 0x01, 0x00, 0x00,
}
