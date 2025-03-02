// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GateClient is the client API for Gate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GateClient interface {
	GetAuthorization(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	GetNotes(ctx context.Context, in *NoteRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	GetTaskStatus(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error)
	DeleteNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error)
}

type gateClient struct {
	cc grpc.ClientConnInterface
}

func NewGateClient(cc grpc.ClientConnInterface) GateClient {
	return &gateClient{cc}
}

func (c *gateClient) GetAuthorization(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/GetAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) GetNotes(ctx context.Context, in *NoteRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/GetNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) GetTaskStatus(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/GetTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) DeleteNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) UpdateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/gate.Gate/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServer is the server API for Gate service.
// All implementations must embed UnimplementedGateServer
// for forward compatibility
type GateServer interface {
	GetAuthorization(context.Context, *AuthRequest) (*TaskResponse, error)
	Register(context.Context, *RegisterRequest) (*TaskResponse, error)
	GetNotes(context.Context, *NoteRequest) (*TaskResponse, error)
	GetTaskStatus(context.Context, *TaskRequest) (*TaskResponse, error)
	CreateNote(context.Context, *Note) (*TaskResponse, error)
	DeleteNote(context.Context, *Note) (*TaskResponse, error)
	UpdateNote(context.Context, *Note) (*TaskResponse, error)
	mustEmbedUnimplementedGateServer()
}

// UnimplementedGateServer must be embedded to have forward compatible implementations.
type UnimplementedGateServer struct {
}

func (UnimplementedGateServer) GetAuthorization(context.Context, *AuthRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorization not implemented")
}
func (UnimplementedGateServer) Register(context.Context, *RegisterRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGateServer) GetNotes(context.Context, *NoteRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotes not implemented")
}
func (UnimplementedGateServer) GetTaskStatus(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskStatus not implemented")
}
func (UnimplementedGateServer) CreateNote(context.Context, *Note) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedGateServer) DeleteNote(context.Context, *Note) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedGateServer) UpdateNote(context.Context, *Note) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedGateServer) mustEmbedUnimplementedGateServer() {}

// UnsafeGateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GateServer will
// result in compilation errors.
type UnsafeGateServer interface {
	mustEmbedUnimplementedGateServer()
}

func RegisterGateServer(s *grpc.Server, srv GateServer) {
	s.RegisterService(&_Gate_serviceDesc, srv)
}

func _Gate_GetAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).GetAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/GetAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).GetAuthorization(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_GetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).GetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/GetNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).GetNotes(ctx, req.(*NoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_GetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).GetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/GetTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).GetTaskStatus(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).CreateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).DeleteNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.Gate/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).UpdateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gate.Gate",
	HandlerType: (*GateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthorization",
			Handler:    _Gate_GetAuthorization_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Gate_Register_Handler,
		},
		{
			MethodName: "GetNotes",
			Handler:    _Gate_GetNotes_Handler,
		},
		{
			MethodName: "GetTaskStatus",
			Handler:    _Gate_GetTaskStatus_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _Gate_CreateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Gate_DeleteNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _Gate_UpdateNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gate/gate.proto",
}
