// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/blog/v1/tag.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TagService_CreateTag_FullMethodName = "/api.blog.v1.TagService/CreateTag"
	TagService_GetTag_FullMethodName    = "/api.blog.v1.TagService/GetTag"
	TagService_UpdateTag_FullMethodName = "/api.blog.v1.TagService/UpdateTag"
	TagService_DeleteTag_FullMethodName = "/api.blog.v1.TagService/DeleteTag"
	TagService_ListTag_FullMethodName   = "/api.blog.v1.TagService/ListTag"
)

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (*ListTagResponse, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_CreateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_GetTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TagService_DeleteTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (*ListTagResponse, error) {
	out := new(ListTagResponse)
	err := c.cc.Invoke(ctx, TagService_ListTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	GetTag(context.Context, *GetTagRequest) (*Tag, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagResponse, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) CreateTag(context.Context, *CreateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagServiceServer) GetTag(context.Context, *GetTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagServiceServer) UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagServiceServer) ListTag(context.Context, *ListTagRequest) (*ListTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTag not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_GetTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_ListTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ListTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_ListTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ListTag(ctx, req.(*ListTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.blog.v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTag",
			Handler:    _TagService_CreateTag_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagService_GetTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagService_DeleteTag_Handler,
		},
		{
			MethodName: "ListTag",
			Handler:    _TagService_ListTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog/v1/tag.proto",
}
