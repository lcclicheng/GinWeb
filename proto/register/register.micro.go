// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/register/register.proto

package go_micro_srv_register

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Register service

type RegisterService interface {
	SmsCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Register(ctx context.Context, in *RegRequset, opts ...client.CallOption) (*RegResponse, error)
	Login(ctx context.Context, in *RegRequset, opts ...client.CallOption) (*RegResponse, error)
}

type registerService struct {
	c    client.Client
	name string
}

func NewRegisterService(name string, c client.Client) RegisterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.register"
	}
	return &registerService{
		c:    c,
		name: name,
	}
}

func (c *registerService) SmsCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Register.SmsCode", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerService) Register(ctx context.Context, in *RegRequset, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "Register.Register", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerService) Login(ctx context.Context, in *RegRequset, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "Register.Login", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterHandler interface {
	SmsCode(context.Context, *Request, *Response) error
	Register(context.Context, *RegRequset, *RegResponse) error
	Login(context.Context, *RegRequset, *RegResponse) error
}

func RegisterRegisterHandler(s server.Server, hdlr RegisterHandler, opts ...server.HandlerOption) error {
	type register interface {
		SmsCode(ctx context.Context, in *Request, out *Response) error
		Register(ctx context.Context, in *RegRequset, out *RegResponse) error
		Login(ctx context.Context, in *RegRequset, out *RegResponse) error
	}
	type Register struct {
		register
	}
	h := &registerHandler{hdlr}
	return s.Handle(s.NewHandler(&Register{h}, opts...))
}

type registerHandler struct {
	RegisterHandler
}

func (h *registerHandler) SmsCode(ctx context.Context, in *Request, out *Response) error {
	return h.RegisterHandler.SmsCode(ctx, in, out)
}

func (h *registerHandler) Register(ctx context.Context, in *RegRequset, out *RegResponse) error {
	return h.RegisterHandler.Register(ctx, in, out)
}

func (h *registerHandler) Login(ctx context.Context, in *RegRequset, out *RegResponse) error {
	return h.RegisterHandler.Login(ctx, in, out)
}
