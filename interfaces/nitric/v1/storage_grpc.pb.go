// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	// Store an item to a bucket
	Put(ctx context.Context, in *StoragePutRequest, opts ...grpc.CallOption) (*StoragePutResponse, error)
	// Retrieve an item from a bucket
	Get(ctx context.Context, in *StorageGetRequest, opts ...grpc.CallOption) (*StorageGetResponse, error)
	// Delete an item from a bucket
	Delete(ctx context.Context, in *StorageDeleteRequest, opts ...grpc.CallOption) (*StorageDeleteResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Put(ctx context.Context, in *StoragePutRequest, opts ...grpc.CallOption) (*StoragePutResponse, error) {
	out := new(StoragePutResponse)
	err := c.cc.Invoke(ctx, "/nitric.v1.storage.Storage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *StorageGetRequest, opts ...grpc.CallOption) (*StorageGetResponse, error) {
	out := new(StorageGetResponse)
	err := c.cc.Invoke(ctx, "/nitric.v1.storage.Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Delete(ctx context.Context, in *StorageDeleteRequest, opts ...grpc.CallOption) (*StorageDeleteResponse, error) {
	out := new(StorageDeleteResponse)
	err := c.cc.Invoke(ctx, "/nitric.v1.storage.Storage/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	// Store an item to a bucket
	Put(context.Context, *StoragePutRequest) (*StoragePutResponse, error)
	// Retrieve an item from a bucket
	Get(context.Context, *StorageGetRequest) (*StorageGetResponse, error)
	// Delete an item from a bucket
	Delete(context.Context, *StorageDeleteRequest) (*StorageDeleteResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Put(context.Context, *StoragePutRequest) (*StoragePutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedStorageServer) Get(context.Context, *StorageGetRequest) (*StorageGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) Delete(context.Context, *StorageDeleteRequest) (*StorageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoragePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.v1.storage.Storage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Put(ctx, req.(*StoragePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.v1.storage.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*StorageGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.v1.storage.Storage/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Delete(ctx, req.(*StorageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nitric.v1.storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Storage_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Storage_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/storage.proto",
}
