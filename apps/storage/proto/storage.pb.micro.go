// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/storage.proto

package storage

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Storage service

func NewStorageEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Storage service

type StorageService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Storage_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Storage_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Storage_BidiStreamService, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*ConnectResponse, error)
	ConnectCache(ctx context.Context, in *ConnectCacheRequest, opts ...client.CallOption) (*ConnectCacheResponse, error)
}

type storageService struct {
	c    client.Client
	name string
}

func NewStorageService(name string, c client.Client) StorageService {
	return &storageService{
		c:    c,
		name: name,
	}
}

func (c *storageService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) ClientStream(ctx context.Context, opts ...client.CallOption) (Storage_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Storage.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &storageServiceClientStream{stream}, nil
}

type Storage_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type storageServiceClientStream struct {
	stream client.Stream
}

func (x *storageServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *storageServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *storageServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *storageService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Storage_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Storage.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &storageServiceServerStream{stream}, nil
}

type Storage_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type storageServiceServerStream struct {
	stream client.Stream
}

func (x *storageServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *storageServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *storageServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageService) BidiStream(ctx context.Context, opts ...client.CallOption) (Storage_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Storage.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &storageServiceBidiStream{stream}, nil
}

type Storage_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type storageServiceBidiStream struct {
	stream client.Stream
}

func (x *storageServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *storageServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *storageServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *storageServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageService) Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*ConnectResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.Connect", in)
	out := new(ConnectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) ConnectCache(ctx context.Context, in *ConnectCacheRequest, opts ...client.CallOption) (*ConnectCacheResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.ConnectCache", in)
	out := new(ConnectCacheResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Storage service

type StorageHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Storage_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Storage_ServerStreamStream) error
	BidiStream(context.Context, Storage_BidiStreamStream) error
	Connect(context.Context, *ConnectRequest, *ConnectResponse) error
	ConnectCache(context.Context, *ConnectCacheRequest, *ConnectCacheResponse) error
}

func RegisterStorageHandler(s server.Server, hdlr StorageHandler, opts ...server.HandlerOption) error {
	type storage interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
		Connect(ctx context.Context, in *ConnectRequest, out *ConnectResponse) error
		ConnectCache(ctx context.Context, in *ConnectCacheRequest, out *ConnectCacheResponse) error
	}
	type Storage struct {
		storage
	}
	h := &storageHandler{hdlr}
	return s.Handle(s.NewHandler(&Storage{h}, opts...))
}

type storageHandler struct {
	StorageHandler
}

func (h *storageHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.StorageHandler.Call(ctx, in, out)
}

func (h *storageHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.StorageHandler.ClientStream(ctx, &storageClientStreamStream{stream})
}

type Storage_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type storageClientStreamStream struct {
	stream server.Stream
}

func (x *storageClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *storageClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *storageHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StorageHandler.ServerStream(ctx, m, &storageServerStreamStream{stream})
}

type Storage_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type storageServerStreamStream struct {
	stream server.Stream
}

func (x *storageServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *storageServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *storageHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.StorageHandler.BidiStream(ctx, &storageBidiStreamStream{stream})
}

type Storage_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type storageBidiStreamStream struct {
	stream server.Stream
}

func (x *storageBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *storageBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *storageBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *storageHandler) Connect(ctx context.Context, in *ConnectRequest, out *ConnectResponse) error {
	return h.StorageHandler.Connect(ctx, in, out)
}

func (h *storageHandler) ConnectCache(ctx context.Context, in *ConnectCacheRequest, out *ConnectCacheResponse) error {
	return h.StorageHandler.ConnectCache(ctx, in, out)
}
