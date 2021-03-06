// Code generated by protoc-gen-go. DO NOT EDIT.
// source: produk.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	produk.proto

It has these top-level messages:
	AddHutangReq
	ReadHutangByNamaReq
	ReadHutangByNamaResp
	ReadHutangResp
	UpdateHutangReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddHutangReq struct {
	KodeHutang    string `protobuf:"bytes,1,opt,name=KodeHutang" json:"KodeHutang,omitempty"`
	NamaHutang    string `protobuf:"bytes,2,opt,name=NamaHutang" json:"NamaHutang,omitempty"`
	TanggalHutang string `protobuf:"bytes,3,opt,name=TanggalHutang" json:"TanggalHutang,omitempty"`
	TotalHutang   string `protobuf:"bytes,4,opt,name=TotalHutang" json:"TotalHutang,omitempty"`
	Status        int32  `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
	CreateBy      string `protobuf:"bytes,6,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *AddHutangReq) Reset()                    { *m = AddHutangReq{} }
func (m *AddHutangReq) String() string            { return proto.CompactTextString(m) }
func (*AddHutangReq) ProtoMessage()               {}
func (*AddHutangReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddHutangReq) GetKodeHutang() string {
	if m != nil {
		return m.KodeHutang
	}
	return ""
}

func (m *AddHutangReq) GetNamaHutang() string {
	if m != nil {
		return m.NamaHutang
	}
	return ""
}

func (m *AddHutangReq) GetTanggalHutang() string {
	if m != nil {
		return m.TanggalHutang
	}
	return ""
}

func (m *AddHutangReq) GetTotalHutang() string {
	if m != nil {
		return m.TotalHutang
	}
	return ""
}

func (m *AddHutangReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddHutangReq) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadHutangByNamaReq struct {
	NamaHutang string `protobuf:"bytes,1,opt,name=NamaHutang" json:"NamaHutang,omitempty"`
}

func (m *ReadHutangByNamaReq) Reset()                    { *m = ReadHutangByNamaReq{} }
func (m *ReadHutangByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadHutangByNamaReq) ProtoMessage()               {}
func (*ReadHutangByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadHutangByNamaReq) GetNamaHutang() string {
	if m != nil {
		return m.NamaHutang
	}
	return ""
}

type ReadHutangByNamaResp struct {
	KodeHutang    string `protobuf:"bytes,1,opt,name=KodeHutang" json:"KodeHutang,omitempty"`
	NamaHutang    string `protobuf:"bytes,2,opt,name=NamaHutang" json:"NamaHutang,omitempty"`
	TanggalHutang string `protobuf:"bytes,3,opt,name=TanggalHutang" json:"TanggalHutang,omitempty"`
	TotalHutang   string `protobuf:"bytes,4,opt,name=TotalHutang" json:"TotalHutang,omitempty"`
	Status        int32  `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
	CreateBy      string `protobuf:"bytes,6,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *ReadHutangByNamaResp) Reset()                    { *m = ReadHutangByNamaResp{} }
func (m *ReadHutangByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadHutangByNamaResp) ProtoMessage()               {}
func (*ReadHutangByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadHutangByNamaResp) GetKodeHutang() string {
	if m != nil {
		return m.KodeHutang
	}
	return ""
}

func (m *ReadHutangByNamaResp) GetNamaHutang() string {
	if m != nil {
		return m.NamaHutang
	}
	return ""
}

func (m *ReadHutangByNamaResp) GetTanggalHutang() string {
	if m != nil {
		return m.TanggalHutang
	}
	return ""
}

func (m *ReadHutangByNamaResp) GetTotalHutang() string {
	if m != nil {
		return m.TotalHutang
	}
	return ""
}

func (m *ReadHutangByNamaResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadHutangByNamaResp) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadHutangResp struct {
	AllHutang []*ReadHutangByNamaResp `protobuf:"bytes,1,rep,name=allHutang" json:"allHutang,omitempty"`
}

func (m *ReadHutangResp) Reset()                    { *m = ReadHutangResp{} }
func (m *ReadHutangResp) String() string            { return proto.CompactTextString(m) }
func (*ReadHutangResp) ProtoMessage()               {}
func (*ReadHutangResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadHutangResp) GetAllHutang() []*ReadHutangByNamaResp {
	if m != nil {
		return m.AllHutang
	}
	return nil
}

type UpdateHutangReq struct {
	KodeHutang    string `protobuf:"bytes,1,opt,name=KodeHutang" json:"KodeHutang,omitempty"`
	NamaHutang    string `protobuf:"bytes,2,opt,name=NamaHutang" json:"NamaHutang,omitempty"`
	TanggalHutang string `protobuf:"bytes,3,opt,name=TanggalHutang" json:"TanggalHutang,omitempty"`
	TotalHutang   string `protobuf:"bytes,4,opt,name=TotalHutang" json:"TotalHutang,omitempty"`
	Status        int32  `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
	UpdateBy      string `protobuf:"bytes,6,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
}

func (m *UpdateHutangReq) Reset()                    { *m = UpdateHutangReq{} }
func (m *UpdateHutangReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateHutangReq) ProtoMessage()               {}
func (*UpdateHutangReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateHutangReq) GetKodeHutang() string {
	if m != nil {
		return m.KodeHutang
	}
	return ""
}

func (m *UpdateHutangReq) GetNamaHutang() string {
	if m != nil {
		return m.NamaHutang
	}
	return ""
}

func (m *UpdateHutangReq) GetTanggalHutang() string {
	if m != nil {
		return m.TanggalHutang
	}
	return ""
}

func (m *UpdateHutangReq) GetTotalHutang() string {
	if m != nil {
		return m.TotalHutang
	}
	return ""
}

func (m *UpdateHutangReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateHutangReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AddHutangReq)(nil), "grpc.AddHutangReq")
	proto.RegisterType((*ReadHutangByNamaReq)(nil), "grpc.ReadHutangByNamaReq")
	proto.RegisterType((*ReadHutangByNamaResp)(nil), "grpc.ReadHutangByNamaResp")
	proto.RegisterType((*ReadHutangResp)(nil), "grpc.ReadHutangResp")
	proto.RegisterType((*UpdateHutangReq)(nil), "grpc.UpdateHutangReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for HutangService service

type HutangServiceClient interface {
	AddHutang(ctx context.Context, in *AddHutangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadHutangByNama(ctx context.Context, in *ReadHutangByNamaReq, opts ...grpc1.CallOption) (*ReadHutangByNamaResp, error)
	ReadHutang(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadHutangResp, error)
	UpdateHutang(ctx context.Context, in *UpdateHutangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type hutangServiceClient struct {
	cc *grpc1.ClientConn
}

func NewHutangServiceClient(cc *grpc1.ClientConn) HutangServiceClient {
	return &hutangServiceClient{cc}
}

func (c *hutangServiceClient) AddHutang(ctx context.Context, in *AddHutangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.HutangService/AddHutang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hutangServiceClient) ReadHutangByNama(ctx context.Context, in *ReadHutangByNamaReq, opts ...grpc1.CallOption) (*ReadHutangByNamaResp, error) {
	out := new(ReadHutangByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.HutangService/ReadHutangByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hutangServiceClient) ReadHutang(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadHutangResp, error) {
	out := new(ReadHutangResp)
	err := grpc1.Invoke(ctx, "/grpc.HutangService/ReadHutang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hutangServiceClient) UpdateHutang(ctx context.Context, in *UpdateHutangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.HutangService/UpdateHutang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HutangService service

type HutangServiceServer interface {
	AddHutang(context.Context, *AddHutangReq) (*google_protobuf.Empty, error)
	ReadHutangByNama(context.Context, *ReadHutangByNamaReq) (*ReadHutangByNamaResp, error)
	ReadHutang(context.Context, *google_protobuf.Empty) (*ReadHutangResp, error)
	UpdateHutang(context.Context, *UpdateHutangReq) (*google_protobuf.Empty, error)
}

func RegisterHutangServiceServer(s *grpc1.Server, srv HutangServiceServer) {
	s.RegisterService(&_HutangService_serviceDesc, srv)
}

func _HutangService_AddHutang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHutangReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HutangServiceServer).AddHutang(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.HutangService/AddHutang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HutangServiceServer).AddHutang(ctx, req.(*AddHutangReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HutangService_ReadHutangByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHutangByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HutangServiceServer).ReadHutangByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.HutangService/ReadHutangByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HutangServiceServer).ReadHutangByNama(ctx, req.(*ReadHutangByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HutangService_ReadHutang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HutangServiceServer).ReadHutang(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.HutangService/ReadHutang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HutangServiceServer).ReadHutang(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HutangService_UpdateHutang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHutangReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HutangServiceServer).UpdateHutang(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.HutangService/UpdateHutang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HutangServiceServer).UpdateHutang(ctx, req.(*UpdateHutangReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HutangService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.HutangService",
	HandlerType: (*HutangServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddHutang",
			Handler:    _HutangService_AddHutang_Handler,
		},
		{
			MethodName: "ReadHutangByNama",
			Handler:    _HutangService_ReadHutangByNama_Handler,
		},
		{
			MethodName: "ReadHutang",
			Handler:    _HutangService_ReadHutang_Handler,
		},
		{
			MethodName: "UpdateHutang",
			Handler:    _HutangService_UpdateHutang_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "produk.proto",
}

func init() { proto.RegisterFile("produk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0xf6, 0x0b, 0x3b, 0x6d, 0x55, 0xc6, 0x5a, 0x6a, 0x04, 0x29, 0xc1, 0x43, 0x4f, 0x29,
	0x54, 0x04, 0x05, 0x41, 0xac, 0x08, 0x62, 0xc1, 0x43, 0x5a, 0x7f, 0xc0, 0xb6, 0x59, 0x83, 0x98,
	0x76, 0xd3, 0x74, 0x23, 0xf4, 0x8f, 0xf8, 0xa3, 0x04, 0x2f, 0xfe, 0x22, 0xc9, 0xee, 0x36, 0x89,
	0xd1, 0x78, 0xef, 0x71, 0xde, 0xbc, 0x49, 0xe6, 0xbd, 0x9d, 0x07, 0x0d, 0x3f, 0xe0, 0x4e, 0xf8,
	0x6a, 0xf9, 0x01, 0x17, 0x1c, 0xcb, 0x6e, 0xe0, 0xcf, 0x8c, 0x63, 0x97, 0x73, 0xd7, 0x63, 0x7d,
	0x89, 0x4d, 0xc3, 0xe7, 0x3e, 0x9b, 0xfb, 0x62, 0xad, 0x28, 0xe6, 0x07, 0x81, 0xc6, 0x8d, 0xe3,
	0xdc, 0x87, 0x82, 0x2e, 0x5c, 0x9b, 0x2d, 0xf1, 0x04, 0x60, 0xc4, 0x1d, 0xa6, 0x80, 0x0e, 0xe9,
	0x92, 0x5e, 0xcd, 0x4e, 0x21, 0x51, 0xff, 0x91, 0xce, 0xa9, 0xee, 0x17, 0x55, 0x3f, 0x41, 0xf0,
	0x14, 0x9a, 0x13, 0xba, 0x70, 0x5d, 0xea, 0x69, 0x4a, 0x49, 0x52, 0x7e, 0x82, 0xd8, 0x85, 0xfa,
	0x84, 0x8b, 0x98, 0x53, 0x96, 0x9c, 0x34, 0x84, 0x6d, 0xa8, 0x8e, 0x05, 0x15, 0xe1, 0xaa, 0x53,
	0xe9, 0x92, 0x5e, 0xc5, 0xd6, 0x15, 0x1a, 0xb0, 0x73, 0x1b, 0x30, 0x2a, 0xd8, 0x70, 0xdd, 0xa9,
	0xca, 0xb1, 0xb8, 0x36, 0xcf, 0xe1, 0xc0, 0x66, 0x54, 0x8b, 0x19, 0xae, 0xa3, 0xad, 0xb4, 0xa4,
	0xd4, 0xca, 0x24, 0xbb, 0xb2, 0xf9, 0x45, 0xa0, 0xf5, 0x7b, 0x6e, 0xe5, 0x6f, 0xb5, 0x17, 0x0f,
	0xb0, 0x9b, 0x68, 0x92, 0x6a, 0x2e, 0xa0, 0x46, 0x3d, 0x2f, 0x16, 0x53, 0xea, 0xd5, 0x07, 0x86,
	0x15, 0x5d, 0x88, 0xf5, 0x97, 0x78, 0x3b, 0x21, 0x9b, 0x9f, 0x04, 0xf6, 0x9e, 0x7c, 0x87, 0x0a,
	0xb6, 0x55, 0x77, 0xa2, 0x56, 0x4e, 0xbc, 0xd9, 0xd4, 0x83, 0xf7, 0x22, 0x34, 0xd5, 0xf8, 0x98,
	0x05, 0x6f, 0x2f, 0x33, 0x86, 0x97, 0x50, 0x8b, 0x53, 0x80, 0xa8, 0x5c, 0x49, 0xc7, 0xc2, 0x68,
	0x5b, 0x2a, 0x45, 0xd6, 0x26, 0x45, 0xd6, 0x5d, 0x94, 0x22, 0xb3, 0x80, 0x23, 0xd8, 0xcf, 0xfa,
	0x87, 0x47, 0x79, 0xbe, 0x2e, 0x8d, 0x7f, 0x2c, 0x37, 0x0b, 0x78, 0x05, 0x90, 0x74, 0x30, 0xe7,
	0xa7, 0x46, 0x2b, 0xfb, 0x0d, 0x3d, 0x7d, 0x0d, 0x8d, 0xf4, 0x33, 0xe1, 0xa1, 0xe2, 0x65, 0x9e,
	0x2e, 0x5f, 0xcb, 0xb4, 0x2a, 0x91, 0xb3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0xde, 0x63,
	0x0a, 0x47, 0x04, 0x00, 0x00,
}
