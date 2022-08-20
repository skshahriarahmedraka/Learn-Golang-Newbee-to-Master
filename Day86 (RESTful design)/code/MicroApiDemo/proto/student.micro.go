// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: student.proto

package proto

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

// Client API for StudentService service

type StudentService interface {
	GetStudent(ctx context.Context, in *Request, opts ...client.CallOption) (*Student, error)
}

type studentService struct {
	c    client.Client
	name string
}

func NewStudentService(name string, c client.Client) StudentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &studentService{
		c:    c,
		name: name,
	}
}

func (c *studentService) GetStudent(ctx context.Context, in *Request, opts ...client.CallOption) (*Student, error) {
	req := c.c.NewRequest(c.name, "StudentService.GetStudent", in)
	out := new(Student)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StudentService service

type StudentServiceHandler interface {
	GetStudent(context.Context, *Request, *Student) error
}

func RegisterStudentServiceHandler(s server.Server, hdlr StudentServiceHandler, opts ...server.HandlerOption) error {
	type studentService interface {
		GetStudent(ctx context.Context, in *Request, out *Student) error
	}
	type StudentService struct {
		studentService
	}
	h := &studentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StudentService{h}, opts...))
}

type studentServiceHandler struct {
	StudentServiceHandler
}

func (h *studentServiceHandler) GetStudent(ctx context.Context, in *Request, out *Student) error {
	return h.StudentServiceHandler.GetStudent(ctx, in, out)
}