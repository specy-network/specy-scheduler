// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/proto/specy/request/register.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	math "math"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("relayer/proto/specy/request/register.proto", fileDescriptor_0182d6b526916aa9)
}

var fileDescriptor_0182d6b526916aa9 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x4a, 0xcd, 0x49,
	0xac, 0x4c, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0x48, 0x4d, 0xae, 0xd4,
	0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d,
	0xd2, 0x03, 0x4b, 0x0a, 0xf1, 0x82, 0x65, 0xf5, 0xa0, 0xb2, 0x52, 0xba, 0xf8, 0xb4, 0x86, 0x24,
	0x16, 0x67, 0x07, 0x41, 0xd8, 0x10, 0xdd, 0x52, 0x7a, 0x84, 0x95, 0x17, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x42, 0xd4, 0x1b, 0x45, 0x71, 0x71, 0x04, 0x41, 0xed, 0x17, 0xf2, 0xe3, 0xe2, 0x75, 0x4f,
	0x2d, 0x81, 0x2a, 0x2a, 0xcd, 0x29, 0x11, 0x92, 0xd2, 0x43, 0x71, 0x8b, 0x1e, 0x92, 0x75, 0x52,
	0xd2, 0x58, 0xe5, 0x20, 0x66, 0x2b, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0x89, 0x46, 0x09, 0xc3,
	0x5c, 0x03, 0x71, 0x47, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x66, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0xa1, 0x54, 0xe8, 0x15, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

const (
	Register_GetTaskResult_FullMethodName = "/specy.request.Register/GetTaskResult"
)

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterClient interface {
	GetTaskResult(ctx context.Context, opts ...grpc.CallOption) (Register_GetTaskResultClient, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) GetTaskResult(ctx context.Context, opts ...grpc.CallOption) (Register_GetTaskResultClient, error) {
	stream, err := c.cc.NewStream(ctx, &Register_ServiceDesc.Streams[0], Register_GetTaskResult_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &registerGetTaskResultClient{stream}
	return x, nil
}

type Register_GetTaskResultClient interface {
	Send(*TaskRequest) error
	Recv() (*TaskResponse, error)
	grpc.ClientStream
}

type registerGetTaskResultClient struct {
	grpc.ClientStream
}

func (x *registerGetTaskResultClient) Send(m *TaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registerGetTaskResultClient) Recv() (*TaskResponse, error) {
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterServer is the server API for Register service.
// All implementations must embed UnimplementedRegisterServer
// for forward compatibility
type RegisterServer interface {
	GetTaskResult(Register_GetTaskResultServer) error
	mustEmbedUnimplementedRegisterServer()
}

// UnimplementedRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (UnimplementedRegisterServer) GetTaskResult(Register_GetTaskResultServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaskResult not implemented")
}
func (UnimplementedRegisterServer) mustEmbedUnimplementedRegisterServer() {}

// UnsafeRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServer will
// result in compilation errors.
type UnsafeRegisterServer interface {
	mustEmbedUnimplementedRegisterServer()
}

func RegisterRegisterServer(s grpc.ServiceRegistrar, srv RegisterServer) {
	s.RegisterService(&Register_ServiceDesc, srv)
}

func _Register_GetTaskResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegisterServer).GetTaskResult(&registerGetTaskResultServer{stream})
}

type Register_GetTaskResultServer interface {
	Send(*TaskResponse) error
	Recv() (*TaskRequest, error)
	grpc.ServerStream
}

type registerGetTaskResultServer struct {
	grpc.ServerStream
}

func (x *registerGetTaskResultServer) Send(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registerGetTaskResultServer) Recv() (*TaskRequest, error) {
	m := new(TaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Register_ServiceDesc is the grpc.ServiceDesc for Register service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Register_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "specy.request.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTaskResult",
			Handler:       _Register_GetTaskResult_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "register.proto",
}
