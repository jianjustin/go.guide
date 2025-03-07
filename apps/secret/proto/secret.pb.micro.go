// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/secret.proto

package secret

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Api Endpoints for Secret service

func NewSecretEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Secret service

type SecretService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Secret_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Secret_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Secret_BidiStreamService, error)
}

type secretService struct {
	c    client.Client
	name string
}

func NewSecretService(name string, c client.Client) SecretService {
	return &secretService{
		c:    c,
		name: name,
	}
}

func (c *secretService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Secret.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretService) ClientStream(ctx context.Context, opts ...client.CallOption) (Secret_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Secret.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &secretServiceClientStream{stream}, nil
}

type Secret_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type secretServiceClientStream struct {
	stream client.Stream
}

func (x *secretServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *secretServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *secretServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *secretService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Secret_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Secret.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &secretServiceServerStream{stream}, nil
}

type Secret_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type secretServiceServerStream struct {
	stream client.Stream
}

func (x *secretServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *secretServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *secretServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretService) BidiStream(ctx context.Context, opts ...client.CallOption) (Secret_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Secret.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &secretServiceBidiStream{stream}, nil
}

type Secret_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type secretServiceBidiStream struct {
	stream client.Stream
}

func (x *secretServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *secretServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *secretServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *secretServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Secret service

type SecretHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Secret_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Secret_ServerStreamStream) error
	BidiStream(context.Context, Secret_BidiStreamStream) error
}

func RegisterSecretHandler(s server.Server, hdlr SecretHandler, opts ...server.HandlerOption) error {
	type secret interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Secret struct {
		secret
	}
	h := &secretHandler{hdlr}
	return s.Handle(s.NewHandler(&Secret{h}, opts...))
}

type secretHandler struct {
	SecretHandler
}

func (h *secretHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.SecretHandler.Call(ctx, in, out)
}

func (h *secretHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.SecretHandler.ClientStream(ctx, &secretClientStreamStream{stream})
}

type Secret_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type secretClientStreamStream struct {
	stream server.Stream
}

func (x *secretClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *secretClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *secretHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SecretHandler.ServerStream(ctx, m, &secretServerStreamStream{stream})
}

type Secret_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type secretServerStreamStream struct {
	stream server.Stream
}

func (x *secretServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *secretServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *secretHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.SecretHandler.BidiStream(ctx, &secretBidiStreamStream{stream})
}

type Secret_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type secretBidiStreamStream struct {
	stream server.Stream
}

func (x *secretBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *secretBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secretBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secretBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secretBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *secretBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
