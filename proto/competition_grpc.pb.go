// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// CompetitionClient is the client API for Competition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompetitionClient interface {
	GetSolutions(ctx context.Context, in *GetSolutionsRequest, opts ...grpc.CallOption) (Competition_GetSolutionsClient, error)
	SetSolutions(ctx context.Context, in *SetSolutionsRequest, opts ...grpc.CallOption) (*SetSolutionsResponse, error)
}

type competitionClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionClient(cc grpc.ClientConnInterface) CompetitionClient {
	return &competitionClient{cc}
}

func (c *competitionClient) GetSolutions(ctx context.Context, in *GetSolutionsRequest, opts ...grpc.CallOption) (Competition_GetSolutionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Competition_ServiceDesc.Streams[0], "/competition.Competition/GetSolutions", opts...)
	if err != nil {
		return nil, err
	}
	x := &competitionGetSolutionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Competition_GetSolutionsClient interface {
	Recv() (*GetSolutionsResponse, error)
	grpc.ClientStream
}

type competitionGetSolutionsClient struct {
	grpc.ClientStream
}

func (x *competitionGetSolutionsClient) Recv() (*GetSolutionsResponse, error) {
	m := new(GetSolutionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *competitionClient) SetSolutions(ctx context.Context, in *SetSolutionsRequest, opts ...grpc.CallOption) (*SetSolutionsResponse, error) {
	out := new(SetSolutionsResponse)
	err := c.cc.Invoke(ctx, "/competition.Competition/SetSolutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompetitionServer is the server API for Competition service.
// All implementations must embed UnimplementedCompetitionServer
// for forward compatibility
type CompetitionServer interface {
	GetSolutions(*GetSolutionsRequest, Competition_GetSolutionsServer) error
	SetSolutions(context.Context, *SetSolutionsRequest) (*SetSolutionsResponse, error)
	mustEmbedUnimplementedCompetitionServer()
}

// UnimplementedCompetitionServer must be embedded to have forward compatible implementations.
type UnimplementedCompetitionServer struct {
}

func (UnimplementedCompetitionServer) GetSolutions(*GetSolutionsRequest, Competition_GetSolutionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSolutions not implemented")
}
func (UnimplementedCompetitionServer) SetSolutions(context.Context, *SetSolutionsRequest) (*SetSolutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSolutions not implemented")
}
func (UnimplementedCompetitionServer) mustEmbedUnimplementedCompetitionServer() {}

// UnsafeCompetitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompetitionServer will
// result in compilation errors.
type UnsafeCompetitionServer interface {
	mustEmbedUnimplementedCompetitionServer()
}

func RegisterCompetitionServer(s grpc.ServiceRegistrar, srv CompetitionServer) {
	s.RegisterService(&Competition_ServiceDesc, srv)
}

func _Competition_GetSolutions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSolutionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompetitionServer).GetSolutions(m, &competitionGetSolutionsServer{stream})
}

type Competition_GetSolutionsServer interface {
	Send(*GetSolutionsResponse) error
	grpc.ServerStream
}

type competitionGetSolutionsServer struct {
	grpc.ServerStream
}

func (x *competitionGetSolutionsServer) Send(m *GetSolutionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Competition_SetSolutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSolutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).SetSolutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/competition.Competition/SetSolutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).SetSolutions(ctx, req.(*SetSolutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Competition_ServiceDesc is the grpc.ServiceDesc for Competition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Competition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "competition.Competition",
	HandlerType: (*CompetitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSolutions",
			Handler:    _Competition_SetSolutions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSolutions",
			Handler:       _Competition_GetSolutions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/competition.proto",
}
