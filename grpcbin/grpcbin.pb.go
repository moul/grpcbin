// Code generated by protoc-gen-go.
// source: grpcbin.proto
// DO NOT EDIT!

/*
Package grpcbin is a generated protocol buffer package.

It is generated from these files:
	grpcbin.proto

It has these top-level messages:
	EmptyMessage
	DummyMessage
	IndexReply
*/
package grpcbin

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

type DummyMessage_Enum int32

const (
	DummyMessage_ENUM_0 DummyMessage_Enum = 0
	DummyMessage_ENUM_1 DummyMessage_Enum = 1
	DummyMessage_ENUM_2 DummyMessage_Enum = 2
)

var DummyMessage_Enum_name = map[int32]string{
	0: "ENUM_0",
	1: "ENUM_1",
	2: "ENUM_2",
}
var DummyMessage_Enum_value = map[string]int32{
	"ENUM_0": 0,
	"ENUM_1": 1,
	"ENUM_2": 2,
}

func (x DummyMessage_Enum) String() string {
	return proto.EnumName(DummyMessage_Enum_name, int32(x))
}
func (DummyMessage_Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DummyMessage struct {
	FString  string              `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
	FStrings []string            `protobuf:"bytes,2,rep,name=f_strings,json=fStrings" json:"f_strings,omitempty"`
	FInt32   int32               `protobuf:"varint,3,opt,name=f_int32,json=fInt32" json:"f_int32,omitempty"`
	FInt32S  []int32             `protobuf:"varint,4,rep,packed,name=f_int32s,json=fInt32s" json:"f_int32s,omitempty"`
	FEnum    DummyMessage_Enum   `protobuf:"varint,5,opt,name=f_enum,json=fEnum,enum=grpcbin.DummyMessage_Enum" json:"f_enum,omitempty"`
	FEnums   []DummyMessage_Enum `protobuf:"varint,6,rep,packed,name=f_enums,json=fEnums,enum=grpcbin.DummyMessage_Enum" json:"f_enums,omitempty"`
	FSub     *DummyMessage_Sub   `protobuf:"bytes,7,opt,name=f_sub,json=fSub" json:"f_sub,omitempty"`
	FSubs    []*DummyMessage_Sub `protobuf:"bytes,8,rep,name=f_subs,json=fSubs" json:"f_subs,omitempty"`
	FBool    bool                `protobuf:"varint,9,opt,name=f_bool,json=fBool" json:"f_bool,omitempty"`
	FBools   []bool              `protobuf:"varint,10,rep,packed,name=f_bools,json=fBools" json:"f_bools,omitempty"`
	FInt64   int64               `protobuf:"varint,11,opt,name=f_int64,json=fInt64" json:"f_int64,omitempty"`
	FInt64S  []int64             `protobuf:"varint,12,rep,packed,name=f_int64s,json=fInt64s" json:"f_int64s,omitempty"`
	FBytes   []byte              `protobuf:"bytes,13,opt,name=f_bytes,json=fBytes,proto3" json:"f_bytes,omitempty"`
	FBytess  [][]byte            `protobuf:"bytes,14,rep,name=f_bytess,json=fBytess,proto3" json:"f_bytess,omitempty"`
	FFloat   float32             `protobuf:"fixed32,15,opt,name=f_float,json=fFloat" json:"f_float,omitempty"`
	FFloats  []float32           `protobuf:"fixed32,16,rep,packed,name=f_floats,json=fFloats" json:"f_floats,omitempty"`
}

func (m *DummyMessage) Reset()                    { *m = DummyMessage{} }
func (m *DummyMessage) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()               {}
func (*DummyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DummyMessage) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

func (m *DummyMessage) GetFStrings() []string {
	if m != nil {
		return m.FStrings
	}
	return nil
}

func (m *DummyMessage) GetFInt32() int32 {
	if m != nil {
		return m.FInt32
	}
	return 0
}

func (m *DummyMessage) GetFInt32S() []int32 {
	if m != nil {
		return m.FInt32S
	}
	return nil
}

func (m *DummyMessage) GetFEnum() DummyMessage_Enum {
	if m != nil {
		return m.FEnum
	}
	return DummyMessage_ENUM_0
}

func (m *DummyMessage) GetFEnums() []DummyMessage_Enum {
	if m != nil {
		return m.FEnums
	}
	return nil
}

func (m *DummyMessage) GetFSub() *DummyMessage_Sub {
	if m != nil {
		return m.FSub
	}
	return nil
}

func (m *DummyMessage) GetFSubs() []*DummyMessage_Sub {
	if m != nil {
		return m.FSubs
	}
	return nil
}

func (m *DummyMessage) GetFBool() bool {
	if m != nil {
		return m.FBool
	}
	return false
}

func (m *DummyMessage) GetFBools() []bool {
	if m != nil {
		return m.FBools
	}
	return nil
}

func (m *DummyMessage) GetFInt64() int64 {
	if m != nil {
		return m.FInt64
	}
	return 0
}

func (m *DummyMessage) GetFInt64S() []int64 {
	if m != nil {
		return m.FInt64S
	}
	return nil
}

func (m *DummyMessage) GetFBytes() []byte {
	if m != nil {
		return m.FBytes
	}
	return nil
}

func (m *DummyMessage) GetFBytess() [][]byte {
	if m != nil {
		return m.FBytess
	}
	return nil
}

func (m *DummyMessage) GetFFloat() float32 {
	if m != nil {
		return m.FFloat
	}
	return 0
}

func (m *DummyMessage) GetFFloats() []float32 {
	if m != nil {
		return m.FFloats
	}
	return nil
}

type DummyMessage_Sub struct {
	FString string `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
}

func (m *DummyMessage_Sub) Reset()                    { *m = DummyMessage_Sub{} }
func (m *DummyMessage_Sub) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage_Sub) ProtoMessage()               {}
func (*DummyMessage_Sub) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *DummyMessage_Sub) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

type IndexReply struct {
	Description string                 `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Endpoints   []*IndexReply_Endpoint `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *IndexReply) Reset()                    { *m = IndexReply{} }
func (m *IndexReply) String() string            { return proto.CompactTextString(m) }
func (*IndexReply) ProtoMessage()               {}
func (*IndexReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IndexReply) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IndexReply) GetEndpoints() []*IndexReply_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type IndexReply_Endpoint struct {
	Path        string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *IndexReply_Endpoint) Reset()                    { *m = IndexReply_Endpoint{} }
func (m *IndexReply_Endpoint) String() string            { return proto.CompactTextString(m) }
func (*IndexReply_Endpoint) ProtoMessage()               {}
func (*IndexReply_Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *IndexReply_Endpoint) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IndexReply_Endpoint) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "grpcbin.EmptyMessage")
	proto.RegisterType((*DummyMessage)(nil), "grpcbin.DummyMessage")
	proto.RegisterType((*DummyMessage_Sub)(nil), "grpcbin.DummyMessage.Sub")
	proto.RegisterType((*IndexReply)(nil), "grpcbin.IndexReply")
	proto.RegisterType((*IndexReply_Endpoint)(nil), "grpcbin.IndexReply.Endpoint")
	proto.RegisterEnum("grpcbin.DummyMessage_Enum", DummyMessage_Enum_name, DummyMessage_Enum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCBin service

type GRPCBinClient interface {
	Index(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*IndexReply, error)
	DummyUnary(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (*DummyMessage, error)
	DummyServerStream(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (GRPCBin_DummyServerStreamClient, error)
	DummyClientStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyClientStreamClient, error)
	DummyBidirectionalStreamStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyBidirectionalStreamStreamClient, error)
}

type gRPCBinClient struct {
	cc *grpc.ClientConn
}

func NewGRPCBinClient(cc *grpc.ClientConn) GRPCBinClient {
	return &gRPCBinClient{cc}
}

func (c *gRPCBinClient) Index(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*IndexReply, error) {
	out := new(IndexReply)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) DummyUnary(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (*DummyMessage, error) {
	out := new(DummyMessage)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/DummyUnary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) DummyServerStream(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (GRPCBin_DummyServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[0], c.cc, "/grpcbin.GRPCBin/DummyServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPCBin_DummyServerStreamClient interface {
	Recv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyServerStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyServerStreamClient) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCBinClient) DummyClientStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[1], c.cc, "/grpcbin.GRPCBin/DummyClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyClientStreamClient{stream}
	return x, nil
}

type GRPCBin_DummyClientStreamClient interface {
	Send(*DummyMessage) error
	CloseAndRecv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyClientStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyClientStreamClient) Send(m *DummyMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCBinDummyClientStreamClient) CloseAndRecv() (*DummyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCBinClient) DummyBidirectionalStreamStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyBidirectionalStreamStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[2], c.cc, "/grpcbin.GRPCBin/DummyBidirectionalStreamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyBidirectionalStreamStreamClient{stream}
	return x, nil
}

type GRPCBin_DummyBidirectionalStreamStreamClient interface {
	Send(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyBidirectionalStreamStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyBidirectionalStreamStreamClient) Send(m *DummyMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCBinDummyBidirectionalStreamStreamClient) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GRPCBin service

type GRPCBinServer interface {
	Index(context.Context, *EmptyMessage) (*IndexReply, error)
	DummyUnary(context.Context, *DummyMessage) (*DummyMessage, error)
	DummyServerStream(*DummyMessage, GRPCBin_DummyServerStreamServer) error
	DummyClientStream(GRPCBin_DummyClientStreamServer) error
	DummyBidirectionalStreamStream(GRPCBin_DummyBidirectionalStreamStreamServer) error
}

func RegisterGRPCBinServer(s *grpc.Server, srv GRPCBinServer) {
	s.RegisterService(&_GRPCBin_serviceDesc, srv)
}

func _GRPCBin_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).Index(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_DummyUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).DummyUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/DummyUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).DummyUnary(ctx, req.(*DummyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_DummyServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DummyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCBinServer).DummyServerStream(m, &gRPCBinDummyServerStreamServer{stream})
}

type GRPCBin_DummyServerStreamServer interface {
	Send(*DummyMessage) error
	grpc.ServerStream
}

type gRPCBinDummyServerStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyServerStreamServer) Send(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPCBin_DummyClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCBinServer).DummyClientStream(&gRPCBinDummyClientStreamServer{stream})
}

type GRPCBin_DummyClientStreamServer interface {
	SendAndClose(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ServerStream
}

type gRPCBinDummyClientStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyClientStreamServer) SendAndClose(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCBinDummyClientStreamServer) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GRPCBin_DummyBidirectionalStreamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCBinServer).DummyBidirectionalStreamStream(&gRPCBinDummyBidirectionalStreamStreamServer{stream})
}

type GRPCBin_DummyBidirectionalStreamStreamServer interface {
	Send(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ServerStream
}

type gRPCBinDummyBidirectionalStreamStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyBidirectionalStreamStreamServer) Send(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCBinDummyBidirectionalStreamStreamServer) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GRPCBin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcbin.GRPCBin",
	HandlerType: (*GRPCBinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _GRPCBin_Index_Handler,
		},
		{
			MethodName: "DummyUnary",
			Handler:    _GRPCBin_DummyUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DummyServerStream",
			Handler:       _GRPCBin_DummyServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DummyClientStream",
			Handler:       _GRPCBin_DummyClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DummyBidirectionalStreamStream",
			Handler:       _GRPCBin_DummyBidirectionalStreamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcbin.proto",
}

func init() { proto.RegisterFile("grpcbin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x9d, 0xeb, 0x26, 0x69, 0x6f, 0xbb, 0x7d, 0xfd, 0x8c, 0x26, 0xbc, 0x82, 0x90, 0xd5, 0x27,
	0x8b, 0x87, 0xaa, 0xcb, 0x4a, 0x1f, 0x10, 0x0f, 0x68, 0xa3, 0xa0, 0x3d, 0x0c, 0x21, 0x97, 0x3d,
	0x57, 0xc9, 0xea, 0x8c, 0x48, 0x89, 0x13, 0xc5, 0x09, 0xa2, 0xff, 0x89, 0xdf, 0xc4, 0x03, 0xbf,
	0x04, 0xd9, 0x49, 0xb3, 0x0a, 0x06, 0x88, 0x3d, 0xc5, 0xf7, 0x1c, 0x9f, 0x73, 0xaf, 0x4e, 0x6e,
	0x02, 0x87, 0xb7, 0x45, 0x7e, 0x13, 0xc6, 0x6a, 0x9a, 0x17, 0x59, 0x99, 0x11, 0xaf, 0x29, 0x27,
	0x47, 0x30, 0x5c, 0xa6, 0x79, 0xb9, 0xbd, 0x92, 0x5a, 0x07, 0xb7, 0x72, 0xf2, 0xad, 0x0b, 0xc3,
	0x37, 0x55, 0x9a, 0xee, 0x00, 0x72, 0x02, 0xbd, 0x68, 0xad, 0xcb, 0x22, 0x56, 0xb7, 0x14, 0x31,
	0xc4, 0xfb, 0xc2, 0x8b, 0x56, 0xb6, 0x24, 0x4f, 0xa0, 0xbf, 0xa3, 0x34, 0xed, 0x30, 0xcc, 0xfb,
	0xa2, 0xd7, 0x70, 0x9a, 0x3c, 0x06, 0x2f, 0x5a, 0xc7, 0xaa, 0x3c, 0xf3, 0x29, 0x66, 0x88, 0x3b,
	0xc2, 0x8d, 0x2e, 0x4d, 0x55, 0x1b, 0x5a, 0x42, 0xd3, 0x2e, 0xc3, 0xdc, 0x11, 0x5e, 0xcd, 0x68,
	0x72, 0x0a, 0x6e, 0xb4, 0x96, 0xaa, 0x4a, 0xa9, 0xc3, 0x10, 0x3f, 0xf2, 0xc7, 0xd3, 0xdd, 0xd4,
	0xfb, 0x23, 0x4d, 0x97, 0xaa, 0x4a, 0x85, 0x13, 0x99, 0x07, 0x39, 0x33, 0x6d, 0x8c, 0x44, 0x53,
	0x97, 0xe1, 0xbf, 0x68, 0x5c, 0xab, 0xd1, 0x64, 0x0a, 0x4e, 0xb4, 0xd6, 0x55, 0x48, 0x3d, 0x86,
	0xf8, 0xc0, 0x3f, 0xb9, 0x5f, 0xb2, 0xaa, 0x42, 0xd1, 0x8d, 0x56, 0x55, 0x48, 0x66, 0x66, 0x2e,
	0x5d, 0x85, 0x9a, 0xf6, 0x18, 0xfe, 0xb3, 0xc0, 0x31, 0x02, 0x4d, 0x8e, 0x8d, 0x22, 0xcc, 0xb2,
	0x84, 0xf6, 0x19, 0xe2, 0x3d, 0xe1, 0x44, 0xe7, 0x59, 0x96, 0xd4, 0xa1, 0x18, 0x58, 0x53, 0x60,
	0x98, 0xf7, 0x84, 0x6b, 0xf1, 0xbb, 0xb4, 0x16, 0x73, 0x3a, 0x60, 0x88, 0xe3, 0x3a, 0xad, 0xc5,
	0xbc, 0x4d, 0x6b, 0x31, 0xd7, 0x74, 0xc8, 0x30, 0xc7, 0x75, 0x5a, 0x8b, 0x79, 0xa3, 0x09, 0xb7,
	0xa5, 0xd4, 0xf4, 0x90, 0x21, 0x3e, 0x34, 0x66, 0xa6, 0xaa, 0x35, 0x96, 0xd0, 0xf4, 0x88, 0x61,
	0x3e, 0x14, 0x5e, 0xcd, 0x34, 0x9a, 0x28, 0xc9, 0x82, 0x92, 0xfe, 0xc7, 0x10, 0xef, 0x08, 0x37,
	0x7a, 0x6b, 0xaa, 0x5a, 0x63, 0x09, 0x4d, 0x47, 0x0c, 0xf3, 0x8e, 0xf0, 0x6a, 0x46, 0x8f, 0x19,
	0x60, 0x13, 0xc2, 0xef, 0x17, 0x61, 0xf2, 0x1c, 0xba, 0xf6, 0x65, 0x00, 0xb8, 0xcb, 0xf7, 0xd7,
	0x57, 0xeb, 0xd9, 0xe8, 0xa0, 0x3d, 0x9f, 0x8e, 0x50, 0x7b, 0xf6, 0x47, 0x9d, 0xc9, 0x57, 0x04,
	0x70, 0xa9, 0x36, 0xf2, 0x8b, 0x90, 0x79, 0xb2, 0x25, 0x0c, 0x06, 0x1b, 0xa9, 0x6f, 0x8a, 0x38,
	0x2f, 0xe3, 0x4c, 0x35, 0xc6, 0xfb, 0x10, 0x79, 0x09, 0x7d, 0xa9, 0x36, 0x79, 0x16, 0xab, 0xb2,
	0xde, 0xb2, 0x81, 0xff, 0xb4, 0xcd, 0xff, 0xce, 0x69, 0xba, 0x6c, 0x2e, 0x89, 0xbb, 0xeb, 0xe3,
	0xd7, 0xd0, 0xdb, 0xc1, 0x84, 0x40, 0x37, 0x0f, 0xca, 0x4f, 0x4d, 0x0b, 0x7b, 0xfe, 0xb9, 0x7b,
	0xe7, 0x97, 0xee, 0xfe, 0xf7, 0x0e, 0x78, 0xef, 0xc4, 0x87, 0x8b, 0xf3, 0x58, 0x91, 0x17, 0xe0,
	0xd8, 0x7e, 0xe4, 0xb8, 0xed, 0xbf, 0xff, 0xed, 0x8c, 0x1f, 0xdd, 0x33, 0xd6, 0xe4, 0x80, 0xbc,
	0x02, 0xb0, 0x6b, 0x72, 0xad, 0x82, 0x62, 0xbb, 0xa7, 0xdd, 0xdf, 0x9d, 0xf1, 0xfd, 0xf0, 0xe4,
	0x80, 0x2c, 0xe1, 0x7f, 0x8b, 0xac, 0x64, 0xf1, 0x59, 0x16, 0xab, 0xb2, 0x90, 0x41, 0xfa, 0xaf,
	0x26, 0x33, 0xd4, 0xda, 0x5c, 0x24, 0xb1, 0x54, 0xe5, 0xc3, 0x6c, 0x38, 0x22, 0x1f, 0xe1, 0x99,
	0xc5, 0xce, 0xe3, 0x4d, 0x5c, 0xc8, 0x1b, 0x13, 0x51, 0x90, 0xd4, 0x6e, 0x0f, 0xf5, 0x9c, 0xa1,
	0xd0, 0xb5, 0x3f, 0xa5, 0xb3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x55, 0x2f, 0x53, 0xa5,
	0x04, 0x00, 0x00,
}
