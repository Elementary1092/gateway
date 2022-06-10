// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: CRUDService.proto

package services

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

// CRUDServiceClient is the client API for CRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDServiceClient interface {
	SavePost(ctx context.Context, in *SavePostDTO, opts ...grpc.CallOption) (*PostDTO, error)
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (CRUDService_GetPostsClient, error)
	DeletePost(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostDTO, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type cRUDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDServiceClient(cc grpc.ClientConnInterface) CRUDServiceClient {
	return &cRUDServiceClient{cc}
}

func (c *cRUDServiceClient) SavePost(ctx context.Context, in *SavePostDTO, opts ...grpc.CallOption) (*PostDTO, error) {
	out := new(PostDTO)
	err := c.cc.Invoke(ctx, "/services.CRUDService/SavePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (CRUDService_GetPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CRUDService_ServiceDesc.Streams[0], "/services.CRUDService/GetPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &cRUDServiceGetPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CRUDService_GetPostsClient interface {
	Recv() (*PostDTO, error)
	grpc.ClientStream
}

type cRUDServiceGetPostsClient struct {
	grpc.ClientStream
}

func (x *cRUDServiceGetPostsClient) Recv() (*PostDTO, error) {
	m := new(PostDTO)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cRUDServiceClient) DeletePost(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/services.CRUDService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) UpdatePost(ctx context.Context, in *UpdatePostDTO, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/services.CRUDService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServiceServer is the server API for CRUDService service.
// All implementations must embed UnimplementedCRUDServiceServer
// for forward compatibility
type CRUDServiceServer interface {
	SavePost(context.Context, *SavePostDTO) (*PostDTO, error)
	GetPosts(*GetPostsRequest, CRUDService_GetPostsServer) error
	DeletePost(context.Context, *DeleteRequest) (*ErrorResponse, error)
	UpdatePost(context.Context, *UpdatePostDTO) (*ErrorResponse, error)
	mustEmbedUnimplementedCRUDServiceServer()
}

// UnimplementedCRUDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServiceServer struct {
}

func (UnimplementedCRUDServiceServer) SavePost(context.Context, *SavePostDTO) (*PostDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
}
func (UnimplementedCRUDServiceServer) GetPosts(*GetPostsRequest, CRUDService_GetPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedCRUDServiceServer) DeletePost(context.Context, *DeleteRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedCRUDServiceServer) UpdatePost(context.Context, *UpdatePostDTO) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedCRUDServiceServer) mustEmbedUnimplementedCRUDServiceServer() {}

// UnsafeCRUDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServiceServer will
// result in compilation errors.
type UnsafeCRUDServiceServer interface {
	mustEmbedUnimplementedCRUDServiceServer()
}

func RegisterCRUDServiceServer(s grpc.ServiceRegistrar, srv CRUDServiceServer) {
	s.RegisterService(&CRUDService_ServiceDesc, srv)
}

func _CRUDService_SavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePostDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).SavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CRUDService/SavePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).SavePost(ctx, req.(*SavePostDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_GetPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CRUDServiceServer).GetPosts(m, &cRUDServiceGetPostsServer{stream})
}

type CRUDService_GetPostsServer interface {
	Send(*PostDTO) error
	grpc.ServerStream
}

type cRUDServiceGetPostsServer struct {
	grpc.ServerStream
}

func (x *cRUDServiceGetPostsServer) Send(m *PostDTO) error {
	return x.ServerStream.SendMsg(m)
}

func _CRUDService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CRUDService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).DeletePost(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CRUDService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).UpdatePost(ctx, req.(*UpdatePostDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUDService_ServiceDesc is the grpc.ServiceDesc for CRUDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.CRUDService",
	HandlerType: (*CRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SavePost",
			Handler:    _CRUDService_SavePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _CRUDService_DeletePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _CRUDService_UpdatePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPosts",
			Handler:       _CRUDService_GetPosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "CRUDService.proto",
}