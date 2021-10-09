// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// ProgrammerServiceClient is the client API for ProgrammerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgrammerServiceClient interface {
	GetProgrammerInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type programmerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgrammerServiceClient(cc grpc.ClientConnInterface) ProgrammerServiceClient {
	return &programmerServiceClient{cc}
}

func (c *programmerServiceClient) GetProgrammerInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.ProgrammerService/GetProgrammerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgrammerServiceServer is the server API for ProgrammerService service.
// All implementations must embed UnimplementedProgrammerServiceServer
// for forward compatibility
type ProgrammerServiceServer interface {
	GetProgrammerInfo(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedProgrammerServiceServer()
}

// UnimplementedProgrammerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgrammerServiceServer struct {
}

func (UnimplementedProgrammerServiceServer) GetProgrammerInfo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgrammerInfo not implemented")
}
func (UnimplementedProgrammerServiceServer) mustEmbedUnimplementedProgrammerServiceServer() {}

// UnsafeProgrammerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgrammerServiceServer will
// result in compilation errors.
type UnsafeProgrammerServiceServer interface {
	mustEmbedUnimplementedProgrammerServiceServer()
}

func RegisterProgrammerServiceServer(s grpc.ServiceRegistrar, srv ProgrammerServiceServer) {
	s.RegisterService(&ProgrammerService_ServiceDesc, srv)
}

func _ProgrammerService_GetProgrammerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgrammerServiceServer).GetProgrammerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProgrammerService/GetProgrammerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgrammerServiceServer).GetProgrammerInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgrammerService_ServiceDesc is the grpc.ServiceDesc for ProgrammerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgrammerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProgrammerService",
	HandlerType: (*ProgrammerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProgrammerInfo",
			Handler:    _ProgrammerService_GetProgrammerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "programmer.proto",
}