// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SumRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	Sum                  int32    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type PrimeNumberRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberRequest) Reset()         { *m = PrimeNumberRequest{} }
func (m *PrimeNumberRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberRequest) ProtoMessage()    {}
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberRequest.Unmarshal(m, b)
}
func (m *PrimeNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberRequest.Merge(m, src)
}
func (m *PrimeNumberRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberRequest.Size(m)
}
func (m *PrimeNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberRequest proto.InternalMessageInfo

func (m *PrimeNumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberResponse) Reset()         { *m = PrimeNumberResponse{} }
func (m *PrimeNumberResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberResponse) ProtoMessage()    {}
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberResponse.Unmarshal(m, b)
}
func (m *PrimeNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberResponse.Merge(m, src)
}
func (m *PrimeNumberResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberResponse.Size(m)
}
func (m *PrimeNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberResponse proto.InternalMessageInfo

func (m *PrimeNumberResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaxRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxRequest) Reset()         { *m = FindMaxRequest{} }
func (m *FindMaxRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxRequest) ProtoMessage()    {}
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxRequest.Unmarshal(m, b)
}
func (m *FindMaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxRequest.Merge(m, src)
}
func (m *FindMaxRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxRequest.Size(m)
}
func (m *FindMaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxRequest proto.InternalMessageInfo

func (m *FindMaxRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaxResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxResponse) Reset()         { *m = FindMaxResponse{} }
func (m *FindMaxResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxResponse) ProtoMessage()    {}
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxResponse.Unmarshal(m, b)
}
func (m *FindMaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxResponse.Merge(m, src)
}
func (m *FindMaxResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxResponse.Size(m)
}
func (m *FindMaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxResponse proto.InternalMessageInfo

func (m *FindMaxResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberRequest)(nil), "calculator.PrimeNumberRequest")
	proto.RegisterType((*PrimeNumberResponse)(nil), "calculator.PrimeNumberResponse")
	proto.RegisterType((*AverageRequest)(nil), "calculator.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "calculator.AverageResponse")
	proto.RegisterType((*FindMaxRequest)(nil), "calculator.FindMaxRequest")
	proto.RegisterType((*FindMaxResponse)(nil), "calculator.FindMaxResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x7f, 0xd9, 0xf8, 0x6d, 0xf0, 0x6c, 0x6e, 0xfa, 0x08, 0x73, 0x54, 0x70, 0x5a, 0x2f,
	0x15, 0x75, 0x0e, 0xbd, 0x78, 0x55, 0x61, 0x78, 0x51, 0xa4, 0xf3, 0xe4, 0x45, 0xda, 0x2e, 0x4a,
	0x61, 0xfd, 0x63, 0xd2, 0x0c, 0xdf, 0x97, 0x6f, 0x50, 0x9a, 0xa6, 0x6d, 0xe2, 0x36, 0xbd, 0x25,
	0xdf, 0x7c, 0xf8, 0x24, 0x7c, 0x9f, 0x80, 0x13, 0x78, 0x8b, 0x40, 0x2c, 0xbc, 0x2c, 0x61, 0x17,
	0xf5, 0x32, 0xf5, 0xb5, 0xcd, 0x38, 0x65, 0x49, 0x96, 0x20, 0xd4, 0x89, 0xfd, 0x0c, 0x30, 0x13,
	0x91, 0x4b, 0x3f, 0x04, 0xe5, 0x19, 0x1e, 0x41, 0xf7, 0x2d, 0x64, 0x3c, 0x7b, 0x8d, 0x45, 0xe4,
	0x53, 0x36, 0x24, 0x87, 0xc4, 0xf9, 0xef, 0x76, 0x64, 0xf6, 0x28, 0x23, 0x3c, 0x86, 0x2d, 0x4e,
	0x83, 0x24, 0x9e, 0x97, 0x4c, 0x43, 0x32, 0xdd, 0x22, 0x2c, 0x20, 0x7b, 0x04, 0x1d, 0x69, 0xe5,
	0x69, 0x12, 0x73, 0x8a, 0xdb, 0xd0, 0xe4, 0x22, 0x52, 0xb6, 0x7c, 0x69, 0x9f, 0x01, 0x3e, 0xb1,
	0x30, 0xa2, 0x05, 0x5f, 0x5e, 0x3f, 0x80, 0x96, 0x71, 0xb1, 0xda, 0xd9, 0xe7, 0xb0, 0x6b, 0xd0,
	0x4a, 0xbb, 0x09, 0x77, 0xa0, 0x77, 0xb3, 0xa4, 0xcc, 0x7b, 0xa7, 0xeb, 0xc5, 0xcd, 0x8a, 0x3c,
	0x85, 0x7e, 0x45, 0x2a, 0xe9, 0x10, 0xda, 0x5e, 0x11, 0x49, 0x96, 0xb8, 0xe5, 0x36, 0xd7, 0x4e,
	0xc3, 0x78, 0xfe, 0xe0, 0x7d, 0xfe, 0xf5, 0xde, 0x13, 0xe8, 0x57, 0xe4, 0xef, 0x6f, 0xbd, 0xfc,
	0x6a, 0xc0, 0xce, 0x5d, 0x35, 0x8e, 0x19, 0x65, 0xcb, 0x30, 0xa0, 0x78, 0x0d, 0xcd, 0x99, 0x88,
	0x70, 0x30, 0xd6, 0x66, 0x57, 0x8f, 0xc9, 0xda, 0x5b, 0xc9, 0x8b, 0x5b, 0xec, 0x7f, 0xe8, 0x42,
	0x47, 0xab, 0x0a, 0x0f, 0x74, 0x72, 0xb5, 0x71, 0x6b, 0xb4, 0xf1, 0xbc, 0x34, 0x4e, 0x08, 0x4e,
	0xa1, 0xad, 0x5a, 0x42, 0x4b, 0xe7, 0xcd, 0x92, 0xad, 0xfd, 0xb5, 0x67, 0xa5, 0xc7, 0x21, 0x78,
	0x0f, 0x6d, 0x55, 0x8b, 0xe9, 0x31, 0x5b, 0x35, 0x3d, 0x3f, 0x7a, 0xcc, 0x3d, 0x13, 0x72, 0xdb,
	0x7b, 0xe9, 0xea, 0x5f, 0xdc, 0x6f, 0xc9, 0x8f, 0x7d, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0x60, 0x62, 0xa6, 0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server Streaming
	PrimeNumber(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error)
	// Client Streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	// BiDi Streaming
	FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumber(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaxClient{stream}
	return x, nil
}

type CalculatorService_FindMaxClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaxClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server Streaming
	PrimeNumber(*PrimeNumberRequest, CalculatorService_PrimeNumberServer) error
	// Client Streaming
	Average(CalculatorService_AverageServer) error
	// BiDi Streaming
	FindMax(CalculatorService_FindMaxServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumber(req *PrimeNumberRequest, srv CalculatorService_PrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumber not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(srv CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMax(srv CalculatorService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumber(m, &calculatorServicePrimeNumberServer{stream})
}

type CalculatorService_PrimeNumberServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMax(&calculatorServiceFindMaxServer{stream})
}

type CalculatorService_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumber",
			Handler:       _CalculatorService_PrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _CalculatorService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
