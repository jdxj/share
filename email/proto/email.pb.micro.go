// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/email.proto

package email

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Email service

func NewEmailEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Email service

type EmailService interface {
	Send(ctx context.Context, in *RequestEmail, opts ...client.CallOption) (Email_SendService, error)
}

type emailService struct {
	c    client.Client
	name string
}

func NewEmailService(name string, c client.Client) EmailService {
	return &emailService{
		c:    c,
		name: name,
	}
}

func (c *emailService) Send(ctx context.Context, in *RequestEmail, opts ...client.CallOption) (Email_SendService, error) {
	req := c.c.NewRequest(c.name, "Email.Send", &RequestEmail{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &emailServiceSend{stream}, nil
}

type Email_SendService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ResponseEmail, error)
}

type emailServiceSend struct {
	stream client.Stream
}

func (x *emailServiceSend) Close() error {
	return x.stream.Close()
}

func (x *emailServiceSend) Context() context.Context {
	return x.stream.Context()
}

func (x *emailServiceSend) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *emailServiceSend) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *emailServiceSend) Recv() (*ResponseEmail, error) {
	m := new(ResponseEmail)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Email service

type EmailHandler interface {
	Send(context.Context, *RequestEmail, Email_SendStream) error
}

func RegisterEmailHandler(s server.Server, hdlr EmailHandler, opts ...server.HandlerOption) error {
	type email interface {
		Send(ctx context.Context, stream server.Stream) error
	}
	type Email struct {
		email
	}
	h := &emailHandler{hdlr}
	return s.Handle(s.NewHandler(&Email{h}, opts...))
}

type emailHandler struct {
	EmailHandler
}

func (h *emailHandler) Send(ctx context.Context, stream server.Stream) error {
	m := new(RequestEmail)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.EmailHandler.Send(ctx, m, &emailSendStream{stream})
}

type Email_SendStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ResponseEmail) error
}

type emailSendStream struct {
	stream server.Stream
}

func (x *emailSendStream) Close() error {
	return x.stream.Close()
}

func (x *emailSendStream) Context() context.Context {
	return x.stream.Context()
}

func (x *emailSendStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *emailSendStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *emailSendStream) Send(m *ResponseEmail) error {
	return x.stream.Send(m)
}
