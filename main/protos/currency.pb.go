package protos

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
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

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	// Base is the base currency code for the rate
	Base string `protobuf:"bytes,1,opt,name=Base,json=base,proto3" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination         string   `protobuf:"bytes,2,opt,name=Destination,json=destination,proto3" json:"Destination,omitempty"`
	XxxNounKeyedLiteral struct{} `json:"-"`
	XxxUnrecognized     []byte   `json:"-"`
	XxxSizeCache        int32    `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return filedescriptorD3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XxxUnmarshal(b []byte) error {
	return xxxMessageInfoRateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XxxMarshal(b []byte, deterministic bool) ([]byte, error) {
	return xxxMessageInfoRateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XxxMerge(src proto.Message) {
	xxxMessageInfoRateRequest.Merge(m, src)
}
func (m *RateRequest) XxxSize() int {
	return xxxMessageInfoRateRequest.Size(m)
}
func (m *RateRequest) XxxDiscardUnknown() {
	xxxMessageInfoRateRequest.DiscardUnknown(m)
}

var xxxMessageInfoRateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *RateRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

// RateResponse is the response from a GetRate call, it contains
// rate which is a floating point number and can be used to convert between the
// two currencies specified in the request.
type RateResponse struct {
	Rate                float32  `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XxxNounkeyedliteral struct{} `json:"-"`
	XxxUnrecognized     []byte   `json:"-"`
	XxxSizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return filedescriptorD3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XxxUnmarshal(b []byte) error {
	return xxxMessageinfoRateresponse.Unmarshal(m, b)
}
func (m *RateResponse) XxxMarshal(b []byte, deterministic bool) ([]byte, error) {
	return xxxMessageinfoRateresponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XxxMerge(src proto.Message) {
	xxxMessageinfoRateresponse.Merge(m, src)
}
func (m *RateResponse) XxxSize() int {
	return xxxMessageinfoRateresponse.Size(m)
}
func (m *RateResponse) XxxDiscardUnknown() {
	xxxMessageinfoRateresponse.DiscardUnknown(m)
}

var xxxMessageinfoRateresponse proto.InternalMessageInfo

func (m *RateResponse) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterType((*RateRequest)(nil), "RateRequest")
	proto.RegisterType((*RateResponse)(nil), "RateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", filedescriptorD3dc60ed002193ea)
}

var filedescriptorD3dc60ed002193ea = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x4b, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe6, 0xe2, 0x0e, 0x4a,
	0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x71, 0x4a, 0x2c,
	0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x49, 0x4a, 0x2c, 0x4e, 0x15, 0x52, 0xe0,
	0xe2, 0x76, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x02, 0x4b,
	0x71, 0xa7, 0x20, 0x84, 0x94, 0x94, 0xb8, 0x78, 0x20, 0x86, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x82, 0x4c, 0x29, 0x4a, 0x2c, 0x81, 0x98, 0xc2, 0x14, 0x04, 0x66, 0x1b, 0x19, 0x71, 0x71, 0x38,
	0x43, 0xad, 0x16, 0x52, 0xe3, 0x62, 0x77, 0x4f, 0x2d, 0x01, 0x69, 0x11, 0xe2, 0xd1, 0x43, 0xb2,
	0x5e, 0x8a, 0x57, 0x0f, 0xd9, 0x9c, 0x24, 0x36, 0xb0, 0x1b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x22, 0xcb, 0xaf, 0x3b, 0xb5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func _(cc grpc.ClientConnInterface) Client {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
type Server interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(_ context.Context, _ *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv Server) {
	s.RegisterService(&currencyServicedesc, srv)
}

func getRateHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Server).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var currencyServicedesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    getRateHandler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
