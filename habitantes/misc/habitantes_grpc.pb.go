// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: habitantes.proto

package misc

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

// ServicioHabitantesClient is the client API for ServicioHabitantes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioHabitantesClient interface {
	InicializadorHabitantes(ctx context.Context, in *InicializadorRequest, opts ...grpc.CallOption) (ServicioHabitantes_InicializadorHabitantesClient, error)
	ActualizarEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (ServicioHabitantes_ActualizarEstadoClient, error)
	ConsumirRecurso(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (ServicioHabitantes_ConsumirRecursoClient, error)
}

type servicioHabitantesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioHabitantesClient(cc grpc.ClientConnInterface) ServicioHabitantesClient {
	return &servicioHabitantesClient{cc}
}

func (c *servicioHabitantesClient) InicializadorHabitantes(ctx context.Context, in *InicializadorRequest, opts ...grpc.CallOption) (ServicioHabitantes_InicializadorHabitantesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicioHabitantes_ServiceDesc.Streams[0], "/habitantes.ServicioHabitantes/InicializadorHabitantes", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicioHabitantesInicializadorHabitantesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicioHabitantes_InicializadorHabitantesClient interface {
	Recv() (*InicializadorResponse, error)
	grpc.ClientStream
}

type servicioHabitantesInicializadorHabitantesClient struct {
	grpc.ClientStream
}

func (x *servicioHabitantesInicializadorHabitantesClient) Recv() (*InicializadorResponse, error) {
	m := new(InicializadorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicioHabitantesClient) ActualizarEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (ServicioHabitantes_ActualizarEstadoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicioHabitantes_ServiceDesc.Streams[1], "/habitantes.ServicioHabitantes/ActualizarEstado", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicioHabitantesActualizarEstadoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicioHabitantes_ActualizarEstadoClient interface {
	Recv() (*EstadoResponse, error)
	grpc.ClientStream
}

type servicioHabitantesActualizarEstadoClient struct {
	grpc.ClientStream
}

func (x *servicioHabitantesActualizarEstadoClient) Recv() (*EstadoResponse, error) {
	m := new(EstadoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicioHabitantesClient) ConsumirRecurso(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (ServicioHabitantes_ConsumirRecursoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicioHabitantes_ServiceDesc.Streams[2], "/habitantes.ServicioHabitantes/ConsumirRecurso", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicioHabitantesConsumirRecursoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicioHabitantes_ConsumirRecursoClient interface {
	Recv() (*EstadoResponse, error)
	grpc.ClientStream
}

type servicioHabitantesConsumirRecursoClient struct {
	grpc.ClientStream
}

func (x *servicioHabitantesConsumirRecursoClient) Recv() (*EstadoResponse, error) {
	m := new(EstadoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicioHabitantesServer is the server API for ServicioHabitantes service.
// All implementations must embed UnimplementedServicioHabitantesServer
// for forward compatibility
type ServicioHabitantesServer interface {
	InicializadorHabitantes(*InicializadorRequest, ServicioHabitantes_InicializadorHabitantesServer) error
	ActualizarEstado(*EstadoRequest, ServicioHabitantes_ActualizarEstadoServer) error
	ConsumirRecurso(*EstadoRequest, ServicioHabitantes_ConsumirRecursoServer) error
	mustEmbedUnimplementedServicioHabitantesServer()
}

// UnimplementedServicioHabitantesServer must be embedded to have forward compatible implementations.
type UnimplementedServicioHabitantesServer struct {
}

func (UnimplementedServicioHabitantesServer) InicializadorHabitantes(*InicializadorRequest, ServicioHabitantes_InicializadorHabitantesServer) error {
	return status.Errorf(codes.Unimplemented, "method InicializadorHabitantes not implemented")
}
func (UnimplementedServicioHabitantesServer) ActualizarEstado(*EstadoRequest, ServicioHabitantes_ActualizarEstadoServer) error {
	return status.Errorf(codes.Unimplemented, "method ActualizarEstado not implemented")
}
func (UnimplementedServicioHabitantesServer) ConsumirRecurso(*EstadoRequest, ServicioHabitantes_ConsumirRecursoServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumirRecurso not implemented")
}
func (UnimplementedServicioHabitantesServer) mustEmbedUnimplementedServicioHabitantesServer() {}

// UnsafeServicioHabitantesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioHabitantesServer will
// result in compilation errors.
type UnsafeServicioHabitantesServer interface {
	mustEmbedUnimplementedServicioHabitantesServer()
}

func RegisterServicioHabitantesServer(s grpc.ServiceRegistrar, srv ServicioHabitantesServer) {
	s.RegisterService(&ServicioHabitantes_ServiceDesc, srv)
}

func _ServicioHabitantes_InicializadorHabitantes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InicializadorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicioHabitantesServer).InicializadorHabitantes(m, &servicioHabitantesInicializadorHabitantesServer{stream})
}

type ServicioHabitantes_InicializadorHabitantesServer interface {
	Send(*InicializadorResponse) error
	grpc.ServerStream
}

type servicioHabitantesInicializadorHabitantesServer struct {
	grpc.ServerStream
}

func (x *servicioHabitantesInicializadorHabitantesServer) Send(m *InicializadorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServicioHabitantes_ActualizarEstado_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EstadoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicioHabitantesServer).ActualizarEstado(m, &servicioHabitantesActualizarEstadoServer{stream})
}

type ServicioHabitantes_ActualizarEstadoServer interface {
	Send(*EstadoResponse) error
	grpc.ServerStream
}

type servicioHabitantesActualizarEstadoServer struct {
	grpc.ServerStream
}

func (x *servicioHabitantesActualizarEstadoServer) Send(m *EstadoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServicioHabitantes_ConsumirRecurso_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EstadoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicioHabitantesServer).ConsumirRecurso(m, &servicioHabitantesConsumirRecursoServer{stream})
}

type ServicioHabitantes_ConsumirRecursoServer interface {
	Send(*EstadoResponse) error
	grpc.ServerStream
}

type servicioHabitantesConsumirRecursoServer struct {
	grpc.ServerStream
}

func (x *servicioHabitantesConsumirRecursoServer) Send(m *EstadoResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServicioHabitantes_ServiceDesc is the grpc.ServiceDesc for ServicioHabitantes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioHabitantes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "habitantes.ServicioHabitantes",
	HandlerType: (*ServicioHabitantesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InicializadorHabitantes",
			Handler:       _ServicioHabitantes_InicializadorHabitantes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ActualizarEstado",
			Handler:       _ServicioHabitantes_ActualizarEstado_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConsumirRecurso",
			Handler:       _ServicioHabitantes_ConsumirRecurso_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "habitantes.proto",
}