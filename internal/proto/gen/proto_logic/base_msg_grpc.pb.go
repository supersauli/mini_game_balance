// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto_logic/base_msg.proto

package proto_logic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BaseMsgCall_Add_FullMethodName = "/proto_logic.BaseMsgCall/Add"
)

// BaseMsgCallClient is the client API for BaseMsgCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseMsgCallClient interface {
	Add(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*BaseMsg, error)
}

type baseMsgCallClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseMsgCallClient(cc grpc.ClientConnInterface) BaseMsgCallClient {
	return &baseMsgCallClient{cc}
}

func (c *baseMsgCallClient) Add(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*BaseMsg, error) {
	out := new(BaseMsg)
	err := c.cc.Invoke(ctx, BaseMsgCall_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseMsgCallServer is the server API for BaseMsgCall service.
// All implementations must embed UnimplementedBaseMsgCallServer
// for forward compatibility
type BaseMsgCallServer interface {
	Add(context.Context, *BaseMsg) (*BaseMsg, error)
	mustEmbedUnimplementedBaseMsgCallServer()
}

// UnimplementedBaseMsgCallServer must be embedded to have forward compatible implementations.
type UnimplementedBaseMsgCallServer struct {
}

func (UnimplementedBaseMsgCallServer) Add(context.Context, *BaseMsg) (*BaseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedBaseMsgCallServer) mustEmbedUnimplementedBaseMsgCallServer() {}

// UnsafeBaseMsgCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseMsgCallServer will
// result in compilation errors.
type UnsafeBaseMsgCallServer interface {
	mustEmbedUnimplementedBaseMsgCallServer()
}

func RegisterBaseMsgCallServer(s grpc.ServiceRegistrar, srv BaseMsgCallServer) {
	s.RegisterService(&BaseMsgCall_ServiceDesc, srv)
}

func _BaseMsgCall_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseMsgCallServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseMsgCall_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseMsgCallServer).Add(ctx, req.(*BaseMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseMsgCall_ServiceDesc is the grpc.ServiceDesc for BaseMsgCall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseMsgCall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_logic.BaseMsgCall",
	HandlerType: (*BaseMsgCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _BaseMsgCall_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_logic/base_msg.proto",
}