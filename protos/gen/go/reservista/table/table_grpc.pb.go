// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: reservista/table/table.proto

package proto_table

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

// TableClient is the client API for Table service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TableClient interface {
	GetAllTables(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TableListResponse, error)
	GetTablesByRestId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error)
	GetTable(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableObject, error)
	AddTable(ctx context.Context, in *AddTableRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	UpdateTableById(ctx context.Context, in *UpdateTableRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	DeleteTableById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetAvailableTables(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error)
	GetReservedTables(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error)
}

type tableClient struct {
	cc grpc.ClientConnInterface
}

func NewTableClient(cc grpc.ClientConnInterface) TableClient {
	return &tableClient{cc}
}

func (c *tableClient) GetAllTables(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TableListResponse, error) {
	out := new(TableListResponse)
	err := c.cc.Invoke(ctx, "/table.Table/GetAllTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) GetTablesByRestId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error) {
	out := new(TableListResponse)
	err := c.cc.Invoke(ctx, "/table.Table/GetTablesByRestId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) GetTable(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableObject, error) {
	out := new(TableObject)
	err := c.cc.Invoke(ctx, "/table.Table/GetTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) AddTable(ctx context.Context, in *AddTableRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/table.Table/AddTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) UpdateTableById(ctx context.Context, in *UpdateTableRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/table.Table/UpdateTableById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) DeleteTableById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/table.Table/DeleteTableById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) GetAvailableTables(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error) {
	out := new(TableListResponse)
	err := c.cc.Invoke(ctx, "/table.Table/GetAvailableTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableClient) GetReservedTables(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableListResponse, error) {
	out := new(TableListResponse)
	err := c.cc.Invoke(ctx, "/table.Table/GetReservedTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServer is the server API for Table service.
// All implementations must embed UnimplementedTableServer
// for forward compatibility
type TableServer interface {
	GetAllTables(context.Context, *Empty) (*TableListResponse, error)
	GetTablesByRestId(context.Context, *IDRequest) (*TableListResponse, error)
	GetTable(context.Context, *IDRequest) (*TableObject, error)
	AddTable(context.Context, *AddTableRequest) (*StatusResponse, error)
	UpdateTableById(context.Context, *UpdateTableRequest) (*StatusResponse, error)
	DeleteTableById(context.Context, *IDRequest) (*StatusResponse, error)
	GetAvailableTables(context.Context, *IDRequest) (*TableListResponse, error)
	GetReservedTables(context.Context, *IDRequest) (*TableListResponse, error)
	mustEmbedUnimplementedTableServer()
}

// UnimplementedTableServer must be embedded to have forward compatible implementations.
type UnimplementedTableServer struct {
}

func (UnimplementedTableServer) GetAllTables(context.Context, *Empty) (*TableListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTables not implemented")
}
func (UnimplementedTableServer) GetTablesByRestId(context.Context, *IDRequest) (*TableListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTablesByRestId not implemented")
}
func (UnimplementedTableServer) GetTable(context.Context, *IDRequest) (*TableObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (UnimplementedTableServer) AddTable(context.Context, *AddTableRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTable not implemented")
}
func (UnimplementedTableServer) UpdateTableById(context.Context, *UpdateTableRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTableById not implemented")
}
func (UnimplementedTableServer) DeleteTableById(context.Context, *IDRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTableById not implemented")
}
func (UnimplementedTableServer) GetAvailableTables(context.Context, *IDRequest) (*TableListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableTables not implemented")
}
func (UnimplementedTableServer) GetReservedTables(context.Context, *IDRequest) (*TableListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservedTables not implemented")
}
func (UnimplementedTableServer) mustEmbedUnimplementedTableServer() {}

// UnsafeTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TableServer will
// result in compilation errors.
type UnsafeTableServer interface {
	mustEmbedUnimplementedTableServer()
}

func RegisterTableServer(s grpc.ServiceRegistrar, srv TableServer) {
	s.RegisterService(&Table_ServiceDesc, srv)
}

func _Table_GetAllTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).GetAllTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/GetAllTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).GetAllTables(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_GetTablesByRestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).GetTablesByRestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/GetTablesByRestId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).GetTablesByRestId(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).GetTable(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_AddTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).AddTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/AddTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).AddTable(ctx, req.(*AddTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_UpdateTableById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).UpdateTableById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/UpdateTableById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).UpdateTableById(ctx, req.(*UpdateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_DeleteTableById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).DeleteTableById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/DeleteTableById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).DeleteTableById(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_GetAvailableTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).GetAvailableTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/GetAvailableTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).GetAvailableTables(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Table_GetReservedTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).GetReservedTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/table.Table/GetReservedTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).GetReservedTables(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Table_ServiceDesc is the grpc.ServiceDesc for Table service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Table_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "table.Table",
	HandlerType: (*TableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTables",
			Handler:    _Table_GetAllTables_Handler,
		},
		{
			MethodName: "GetTablesByRestId",
			Handler:    _Table_GetTablesByRestId_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _Table_GetTable_Handler,
		},
		{
			MethodName: "AddTable",
			Handler:    _Table_AddTable_Handler,
		},
		{
			MethodName: "UpdateTableById",
			Handler:    _Table_UpdateTableById_Handler,
		},
		{
			MethodName: "DeleteTableById",
			Handler:    _Table_DeleteTableById_Handler,
		},
		{
			MethodName: "GetAvailableTables",
			Handler:    _Table_GetAvailableTables_Handler,
		},
		{
			MethodName: "GetReservedTables",
			Handler:    _Table_GetReservedTables_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservista/table/table.proto",
}