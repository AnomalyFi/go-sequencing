// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sequencing/sequencing.proto

package rollkit

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SubmitRollupTransactionRequest ...
type SubmitRollupTransactionRequest struct {
	// the unhashed rollup id
	RollupId []byte `protobuf:"bytes,1,opt,name=rollup_id,json=rollupId,proto3" json:"rollup_id,omitempty"`
	// the raw data bytes of the rollup transaction
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SubmitRollupTransactionRequest) Reset()         { *m = SubmitRollupTransactionRequest{} }
func (m *SubmitRollupTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitRollupTransactionRequest) ProtoMessage()    {}
func (*SubmitRollupTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b389fdbc59f03c95, []int{0}
}
func (m *SubmitRollupTransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitRollupTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitRollupTransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitRollupTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRollupTransactionRequest.Merge(m, src)
}
func (m *SubmitRollupTransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubmitRollupTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRollupTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRollupTransactionRequest proto.InternalMessageInfo

func (m *SubmitRollupTransactionRequest) GetRollupId() []byte {
	if m != nil {
		return m.RollupId
	}
	return nil
}

func (m *SubmitRollupTransactionRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// SubmitRollupTransactionResponse ...
type SubmitRollupTransactionResponse struct {
}

func (m *SubmitRollupTransactionResponse) Reset()         { *m = SubmitRollupTransactionResponse{} }
func (m *SubmitRollupTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitRollupTransactionResponse) ProtoMessage()    {}
func (*SubmitRollupTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b389fdbc59f03c95, []int{1}
}
func (m *SubmitRollupTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitRollupTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitRollupTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitRollupTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRollupTransactionResponse.Merge(m, src)
}
func (m *SubmitRollupTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *SubmitRollupTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRollupTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRollupTransactionResponse proto.InternalMessageInfo

// Batch holds a list of transactions
type Batch struct {
	Transactions [][]byte `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_b389fdbc59f03c95, []int{2}
}
func (m *Batch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return m.Size()
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetTransactions() [][]byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// VerificationResponse
type VerificationResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *VerificationResponse) Reset()         { *m = VerificationResponse{} }
func (m *VerificationResponse) String() string { return proto.CompactTextString(m) }
func (*VerificationResponse) ProtoMessage()    {}
func (*VerificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b389fdbc59f03c95, []int{3}
}
func (m *VerificationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerificationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationResponse.Merge(m, src)
}
func (m *VerificationResponse) XXX_Size() int {
	return m.Size()
}
func (m *VerificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationResponse proto.InternalMessageInfo

func (m *VerificationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*SubmitRollupTransactionRequest)(nil), "rollkit.SubmitRollupTransactionRequest")
	proto.RegisterType((*SubmitRollupTransactionResponse)(nil), "rollkit.SubmitRollupTransactionResponse")
	proto.RegisterType((*Batch)(nil), "rollkit.Batch")
	proto.RegisterType((*VerificationResponse)(nil), "rollkit.VerificationResponse")
}

func init() { proto.RegisterFile("sequencing/sequencing.proto", fileDescriptor_b389fdbc59f03c95) }

var fileDescriptor_b389fdbc59f03c95 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4e, 0xc2, 0x40,
	0x18, 0xec, 0xfa, 0x07, 0x7e, 0x56, 0x4c, 0x36, 0x26, 0x36, 0x10, 0x57, 0xdc, 0x8b, 0x24, 0x26,
	0x48, 0xf0, 0xee, 0x01, 0x0f, 0x86, 0x8b, 0xc4, 0x62, 0xbc, 0x9a, 0x52, 0x56, 0xdd, 0x58, 0xb7,
	0x75, 0xf7, 0xdb, 0x44, 0x6f, 0x3e, 0x82, 0x8f, 0xe5, 0x91, 0xa3, 0x47, 0x03, 0x2f, 0x62, 0xd8,
	0x22, 0xa8, 0x01, 0xbd, 0xcd, 0xcc, 0xd7, 0x4e, 0x66, 0x26, 0x0b, 0x15, 0x23, 0x1e, 0xad, 0x50,
	0xb1, 0x54, 0xb7, 0x47, 0x33, 0x58, 0xcf, 0x74, 0x8a, 0x29, 0x2d, 0xe8, 0x34, 0x49, 0xee, 0x25,
	0xf2, 0x0b, 0x60, 0x5d, 0xdb, 0x7b, 0x90, 0x18, 0xa6, 0x49, 0x62, 0xb3, 0x4b, 0x1d, 0x29, 0x13,
	0xc5, 0x28, 0x53, 0x15, 0x8e, 0xff, 0x31, 0x48, 0x2b, 0xb0, 0xae, 0xdd, 0xed, 0x5a, 0xf6, 0x03,
	0x52, 0x25, 0x35, 0x3f, 0x2c, 0xe6, 0x42, 0xbb, 0x4f, 0x29, 0xac, 0xf4, 0x23, 0x8c, 0x82, 0x25,
	0xa7, 0x3b, 0xcc, 0xf7, 0x61, 0x6f, 0xa1, 0xa5, 0xc9, 0x52, 0x65, 0x04, 0x3f, 0x84, 0xd5, 0x56,
	0x84, 0xf1, 0x1d, 0xe5, 0xe0, 0xe3, 0xec, 0x6e, 0x02, 0x52, 0x5d, 0xae, 0xf9, 0xe1, 0x0f, 0x8d,
	0x37, 0x60, 0xfb, 0x4a, 0x68, 0x79, 0x23, 0xe3, 0xe8, 0xbb, 0x09, 0x0d, 0xa0, 0x60, 0x6c, 0x1c,
	0x0b, 0x63, 0x5c, 0xac, 0x62, 0xf8, 0x45, 0x9b, 0x2f, 0x04, 0x4a, 0xdd, 0xbc, 0xb2, 0xd0, 0x6d,
	0x95, 0x59, 0xa4, 0x0a, 0x76, 0x16, 0x84, 0xa2, 0x07, 0xf5, 0xc9, 0x18, 0xf5, 0xbf, 0x97, 0x28,
	0xd7, 0xfe, 0xff, 0x70, 0xd2, 0xcf, 0x6b, 0x9e, 0xc2, 0xd6, 0x34, 0x41, 0xc7, 0xe2, 0x38, 0x42,
	0x03, 0xfc, 0x33, 0x81, 0xe7, 0xe2, 0x09, 0xf3, 0xee, 0xa5, 0xa9, 0x9d, 0xe3, 0xe5, 0x5f, 0x9c,
	0x7b, 0xcd, 0x0e, 0x6c, 0x3a, 0x98, 0xd7, 0x17, 0x9a, 0x9e, 0xc0, 0x86, 0xc3, 0xcf, 0xf3, 0x1d,
	0x76, 0xa7, 0x7c, 0xde, 0x60, 0xdc, 0x6b, 0x05, 0x6f, 0x43, 0x46, 0x06, 0x43, 0x46, 0x3e, 0x86,
	0x8c, 0xbc, 0x8e, 0x98, 0x37, 0x18, 0x31, 0xef, 0x7d, 0xc4, 0xbc, 0xde, 0x9a, 0x7b, 0x17, 0xc7,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0xcd, 0xfe, 0xc7, 0x36, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SequencerInputClient is the client API for SequencerInput service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SequencerInputClient interface {
	// SubmitRollupTransaction ...
	SubmitRollupTransaction(ctx context.Context, in *SubmitRollupTransactionRequest, opts ...grpc.CallOption) (*SubmitRollupTransactionResponse, error)
}

type sequencerInputClient struct {
	cc *grpc.ClientConn
}

func NewSequencerInputClient(cc *grpc.ClientConn) SequencerInputClient {
	return &sequencerInputClient{cc}
}

func (c *sequencerInputClient) SubmitRollupTransaction(ctx context.Context, in *SubmitRollupTransactionRequest, opts ...grpc.CallOption) (*SubmitRollupTransactionResponse, error) {
	out := new(SubmitRollupTransactionResponse)
	err := c.cc.Invoke(ctx, "/rollkit.SequencerInput/SubmitRollupTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequencerInputServer is the server API for SequencerInput service.
type SequencerInputServer interface {
	// SubmitRollupTransaction ...
	SubmitRollupTransaction(context.Context, *SubmitRollupTransactionRequest) (*SubmitRollupTransactionResponse, error)
}

// UnimplementedSequencerInputServer can be embedded to have forward compatible implementations.
type UnimplementedSequencerInputServer struct {
}

func (*UnimplementedSequencerInputServer) SubmitRollupTransaction(ctx context.Context, req *SubmitRollupTransactionRequest) (*SubmitRollupTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRollupTransaction not implemented")
}

func RegisterSequencerInputServer(s *grpc.Server, srv SequencerInputServer) {
	s.RegisterService(&_SequencerInput_serviceDesc, srv)
}

func _SequencerInput_SubmitRollupTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRollupTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerInputServer).SubmitRollupTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rollkit.SequencerInput/SubmitRollupTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerInputServer).SubmitRollupTransaction(ctx, req.(*SubmitRollupTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SequencerInput_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rollkit.SequencerInput",
	HandlerType: (*SequencerInputServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitRollupTransaction",
			Handler:    _SequencerInput_SubmitRollupTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencing/sequencing.proto",
}

// SequencerOutputClient is the client API for SequencerOutput service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SequencerOutputClient interface {
	// SubmitRollupTransaction ...
	GetNextBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*Batch, error)
}

type sequencerOutputClient struct {
	cc *grpc.ClientConn
}

func NewSequencerOutputClient(cc *grpc.ClientConn) SequencerOutputClient {
	return &sequencerOutputClient{cc}
}

func (c *sequencerOutputClient) GetNextBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*Batch, error) {
	out := new(Batch)
	err := c.cc.Invoke(ctx, "/rollkit.SequencerOutput/GetNextBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequencerOutputServer is the server API for SequencerOutput service.
type SequencerOutputServer interface {
	// SubmitRollupTransaction ...
	GetNextBatch(context.Context, *Batch) (*Batch, error)
}

// UnimplementedSequencerOutputServer can be embedded to have forward compatible implementations.
type UnimplementedSequencerOutputServer struct {
}

func (*UnimplementedSequencerOutputServer) GetNextBatch(ctx context.Context, req *Batch) (*Batch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextBatch not implemented")
}

func RegisterSequencerOutputServer(s *grpc.Server, srv SequencerOutputServer) {
	s.RegisterService(&_SequencerOutput_serviceDesc, srv)
}

func _SequencerOutput_GetNextBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Batch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerOutputServer).GetNextBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rollkit.SequencerOutput/GetNextBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerOutputServer).GetNextBatch(ctx, req.(*Batch))
	}
	return interceptor(ctx, in, info, handler)
}

var _SequencerOutput_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rollkit.SequencerOutput",
	HandlerType: (*SequencerOutputServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNextBatch",
			Handler:    _SequencerOutput_GetNextBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencing/sequencing.proto",
}

// BatchVerifierClient is the client API for BatchVerifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BatchVerifierClient interface {
	// VerifyBatch ...
	VerifyBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*VerificationResponse, error)
}

type batchVerifierClient struct {
	cc *grpc.ClientConn
}

func NewBatchVerifierClient(cc *grpc.ClientConn) BatchVerifierClient {
	return &batchVerifierClient{cc}
}

func (c *batchVerifierClient) VerifyBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, "/rollkit.BatchVerifier/VerifyBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchVerifierServer is the server API for BatchVerifier service.
type BatchVerifierServer interface {
	// VerifyBatch ...
	VerifyBatch(context.Context, *Batch) (*VerificationResponse, error)
}

// UnimplementedBatchVerifierServer can be embedded to have forward compatible implementations.
type UnimplementedBatchVerifierServer struct {
}

func (*UnimplementedBatchVerifierServer) VerifyBatch(ctx context.Context, req *Batch) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyBatch not implemented")
}

func RegisterBatchVerifierServer(s *grpc.Server, srv BatchVerifierServer) {
	s.RegisterService(&_BatchVerifier_serviceDesc, srv)
}

func _BatchVerifier_VerifyBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Batch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchVerifierServer).VerifyBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rollkit.BatchVerifier/VerifyBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchVerifierServer).VerifyBatch(ctx, req.(*Batch))
	}
	return interceptor(ctx, in, info, handler)
}

var _BatchVerifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rollkit.BatchVerifier",
	HandlerType: (*BatchVerifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyBatch",
			Handler:    _BatchVerifier_VerifyBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencing/sequencing.proto",
}

func (m *SubmitRollupTransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitRollupTransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitRollupTransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintSequencing(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollupId) > 0 {
		i -= len(m.RollupId)
		copy(dAtA[i:], m.RollupId)
		i = encodeVarintSequencing(dAtA, i, uint64(len(m.RollupId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubmitRollupTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitRollupTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitRollupTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Batch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Batch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Batch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transactions[iNdEx])
			copy(dAtA[i:], m.Transactions[iNdEx])
			i = encodeVarintSequencing(dAtA, i, uint64(len(m.Transactions[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VerificationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerificationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerificationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSequencing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSequencing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SubmitRollupTransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollupId)
	if l > 0 {
		n += 1 + l + sovSequencing(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSequencing(uint64(l))
	}
	return n
}

func (m *SubmitRollupTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Batch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Transactions) > 0 {
		for _, b := range m.Transactions {
			l = len(b)
			n += 1 + l + sovSequencing(uint64(l))
		}
	}
	return n
}

func (m *VerificationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovSequencing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSequencing(x uint64) (n int) {
	return sovSequencing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SubmitRollupTransactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmitRollupTransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitRollupTransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollupId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSequencing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollupId = append(m.RollupId[:0], dAtA[iNdEx:postIndex]...)
			if m.RollupId == nil {
				m.RollupId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSequencing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubmitRollupTransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmitRollupTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitRollupTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSequencing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Batch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Batch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Batch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSequencing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, make([]byte, postIndex-iNdEx))
			copy(m.Transactions[len(m.Transactions)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerificationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencing
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerificationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerificationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSequencing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSequencing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencing
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSequencing
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSequencing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSequencing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSequencing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSequencing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSequencing = fmt.Errorf("proto: unexpected end of group")
)
