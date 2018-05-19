// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Entry
	Entries
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Entry struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Entries struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *Entries) Reset()                    { *m = Entries{} }
func (m *Entries) String() string            { return proto.CompactTextString(m) }
func (*Entries) ProtoMessage()               {}
func (*Entries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entries) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "v1.Entry")
	proto.RegisterType((*Entries)(nil), "v1.Entries")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DNSService service

type DNSServiceClient interface {
	PutRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
	PostRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
	GetRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entries, error)
	DeleteRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
}

type dNSServiceClient struct {
	cc *grpc.ClientConn
}

func NewDNSServiceClient(cc *grpc.ClientConn) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) PutRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/v1.DNSService/PutRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) PostRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/v1.DNSService/PostRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) GetRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entries, error) {
	out := new(Entries)
	err := grpc.Invoke(ctx, "/v1.DNSService/GetRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) DeleteRecord(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := grpc.Invoke(ctx, "/v1.DNSService/DeleteRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DNSService service

type DNSServiceServer interface {
	PutRecord(context.Context, *Entry) (*Entry, error)
	PostRecord(context.Context, *Entry) (*Entry, error)
	GetRecord(context.Context, *Entry) (*Entries, error)
	DeleteRecord(context.Context, *Entry) (*Entry, error)
}

func RegisterDNSServiceServer(s *grpc.Server, srv DNSServiceServer) {
	s.RegisterService(&_DNSService_serviceDesc, srv)
}

func _DNSService_PutRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).PutRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DNSService/PutRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).PutRecord(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_PostRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).PostRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DNSService/PostRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).PostRecord(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DNSService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).GetRecord(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DNSService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).DeleteRecord(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutRecord",
			Handler:    _DNSService_PutRecord_Handler,
		},
		{
			MethodName: "PostRecord",
			Handler:    _DNSService_PostRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _DNSService_GetRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _DNSService_DeleteRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0x86, 0xa8, 0x50, 0x32, 0xe5, 0x62, 0x75, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8,
	0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0xc0, 0xc2, 0x30, 0xae, 0x92, 0x1e,
	0x17, 0x3b, 0x48, 0x5b, 0x66, 0x6a, 0xb1, 0x90, 0x32, 0x17, 0x7b, 0x2a, 0x84, 0x29, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa9, 0x57, 0x66, 0xa8, 0x07, 0x36, 0x34, 0x08, 0x26, 0x63, 0xf4,
	0x85, 0x91, 0x8b, 0xcb, 0xc5, 0x2f, 0x38, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x82,
	0x8b, 0x33, 0xa0, 0xb4, 0x24, 0x28, 0x35, 0x39, 0xbf, 0x28, 0x45, 0x08, 0xa1, 0x5e, 0x0a, 0xc1,
	0x54, 0x12, 0x6d, 0xba, 0xfc, 0x64, 0x32, 0x13, 0xbf, 0x14, 0x97, 0x7e, 0x99, 0xa1, 0x7e, 0x11,
	0x58, 0xa5, 0x15, 0xa3, 0x96, 0x90, 0x25, 0x17, 0x57, 0x40, 0x7e, 0x31, 0x51, 0x5a, 0x95, 0xd0,
	0xb4, 0x9a, 0x73, 0x71, 0xba, 0xa7, 0x62, 0xd1, 0xc9, 0x0d, 0x63, 0x66, 0xa6, 0x16, 0x2b, 0x09,
	0x81, 0xf5, 0xf2, 0x08, 0x21, 0xe9, 0x15, 0xb2, 0xe0, 0xe2, 0x71, 0x49, 0xcd, 0x49, 0x2d, 0x49,
	0xc5, 0x6b, 0x2b, 0x54, 0xa7, 0x16, 0x92, 0xce, 0x24, 0x36, 0x70, 0x20, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0xbf, 0x36, 0xb7, 0x93, 0x01, 0x00, 0x00,
}