// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/bank/types/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryBalanceParams defines the params for querying an account balance.
type QueryBalanceRequest struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	Denom   string                                        `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b761440f9b86d1e8, []int{0}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo

func (m *QueryBalanceRequest) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *QueryBalanceRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryBalanceResponse is the response for the QueryBalance rpc method
type QueryBalanceResponse struct {
	Balance *types.Coin `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b761440f9b86d1e8, []int{1}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetBalance() *types.Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}

// QueryAllBalancesParams defines the params for querying all account balances
type QueryAllBalancesRequest struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	Page    *query.PageRequest                            `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (m *QueryAllBalancesRequest) Reset()         { *m = QueryAllBalancesRequest{} }
func (m *QueryAllBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllBalancesRequest) ProtoMessage()    {}
func (*QueryAllBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b761440f9b86d1e8, []int{2}
}
func (m *QueryAllBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllBalancesRequest.Merge(m, src)
}
func (m *QueryAllBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllBalancesRequest proto.InternalMessageInfo

func (m *QueryAllBalancesRequest) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *QueryAllBalancesRequest) GetPage() *query.PageRequest {
	if m != nil {
		return m.Page
	}
	return nil
}

// QueryAllBalancesResponse is the response to the QueryAllBalances rpc method
type QueryAllBalancesResponse struct {
	Balances github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances"`
	Page     *query.PageResponse                      `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (m *QueryAllBalancesResponse) Reset()         { *m = QueryAllBalancesResponse{} }
func (m *QueryAllBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllBalancesResponse) ProtoMessage()    {}
func (*QueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b761440f9b86d1e8, []int{3}
}
func (m *QueryAllBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllBalancesResponse.Merge(m, src)
}
func (m *QueryAllBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllBalancesResponse proto.InternalMessageInfo

func (m *QueryAllBalancesResponse) GetBalances() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *QueryAllBalancesResponse) GetPage() *query.PageResponse {
	if m != nil {
		return m.Page
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryBalanceRequest)(nil), "cosmos_sdk.x.bank.v1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "cosmos_sdk.x.bank.v1.QueryBalanceResponse")
	proto.RegisterType((*QueryAllBalancesRequest)(nil), "cosmos_sdk.x.bank.v1.QueryAllBalancesRequest")
	proto.RegisterType((*QueryAllBalancesResponse)(nil), "cosmos_sdk.x.bank.v1.QueryAllBalancesResponse")
}

func init() { proto.RegisterFile("x/bank/types/query.proto", fileDescriptor_b761440f9b86d1e8) }

var fileDescriptor_b761440f9b86d1e8 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xbf, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe3, 0xf0, 0xa3, 0xe0, 0x76, 0x00, 0x37, 0x12, 0x51, 0x84, 0x2e, 0xe1, 0x06, 0x94,
	0x56, 0x8a, 0xdd, 0x0b, 0x20, 0xe6, 0xa4, 0x62, 0x62, 0x81, 0x63, 0x63, 0xa9, 0x9c, 0x3b, 0xeb,
	0x7a, 0x6a, 0x62, 0x5f, 0xcf, 0xce, 0x29, 0x51, 0xd5, 0x85, 0xbf, 0x00, 0x89, 0x15, 0xb1, 0xb0,
	0x31, 0xf3, 0x47, 0x74, 0xac, 0x84, 0x90, 0x98, 0x0a, 0x4a, 0xf8, 0x2b, 0x98, 0xd0, 0xd9, 0x0e,
	0xb9, 0xd2, 0x70, 0xca, 0xc2, 0x92, 0xdc, 0xd9, 0xdf, 0xf7, 0xde, 0xe7, 0xeb, 0xf7, 0x7c, 0xb0,
	0x3e, 0x21, 0x03, 0xca, 0x8f, 0x88, 0x9a, 0x26, 0x4c, 0x92, 0xe3, 0x31, 0x4b, 0xa7, 0x38, 0x49,
	0x85, 0x12, 0xa8, 0x16, 0x08, 0x39, 0x12, 0xf2, 0x40, 0x86, 0x47, 0x78, 0x82, 0x73, 0x11, 0xce,
	0xbc, 0xc6, 0x43, 0x75, 0x18, 0xa7, 0xe1, 0x41, 0x42, 0x53, 0x35, 0x25, 0x5a, 0x48, 0x22, 0x11,
	0x89, 0xe5, 0x93, 0x89, 0x6e, 0xe0, 0x55, 0x3a, 0x11, 0x0d, 0x19, 0xa1, 0x49, 0x4c, 0x28, 0xe7,
	0x42, 0x51, 0x15, 0x0b, 0x2e, 0xad, 0xfe, 0xae, 0x01, 0xd0, 0xbf, 0x76, 0xe9, 0x7e, 0x81, 0x89,
	0x24, 0x34, 0x8a, 0xb9, 0x8e, 0x30, 0xbb, 0xee, 0x04, 0x6e, 0xbf, 0xcc, 0x77, 0xfa, 0x74, 0x48,
	0x79, 0xc0, 0x7c, 0x76, 0x3c, 0x66, 0x52, 0xa1, 0xe7, 0x70, 0x83, 0x86, 0x61, 0xca, 0xa4, 0xac,
	0x83, 0x16, 0x68, 0x6f, 0xf5, 0xbd, 0x5f, 0x17, 0xcd, 0x4e, 0x14, 0xab, 0xc3, 0xf1, 0x00, 0x07,
	0x62, 0x44, 0x8c, 0x2b, 0xfb, 0xd7, 0x91, 0xa1, 0xb5, 0x8e, 0x7b, 0x41, 0xd0, 0x33, 0x81, 0xfe,
	0x22, 0x03, 0xaa, 0xc1, 0x1b, 0x21, 0xe3, 0x62, 0x54, 0xaf, 0xb6, 0x40, 0xfb, 0xb6, 0x6f, 0x5e,
	0xdc, 0x67, 0xb0, 0x76, 0xb9, 0xb2, 0x4c, 0x04, 0x97, 0x0c, 0x75, 0xe0, 0xc6, 0xc0, 0x2c, 0xe9,
	0xd2, 0x9b, 0xdd, 0x6d, 0x5c, 0x38, 0xc2, 0xcc, 0xc3, 0xfb, 0x22, 0xe6, 0xfe, 0x42, 0xe3, 0xbe,
	0x07, 0xf0, 0x9e, 0xce, 0xd3, 0x1b, 0x0e, 0x6d, 0x2a, 0xf9, 0x5f, 0x5c, 0x3c, 0x86, 0xd7, 0x13,
	0x1a, 0x31, 0x6d, 0x62, 0xb3, 0xdb, 0x2a, 0x42, 0x99, 0x7e, 0x67, 0x1e, 0x7e, 0x41, 0xa3, 0xc5,
	0x11, 0xfa, 0x5a, 0xed, 0x7e, 0x06, 0xb0, 0x7e, 0x15, 0xcf, 0x5a, 0xa5, 0xf0, 0x96, 0xb5, 0x91,
	0x03, 0x5e, 0xfb, 0x87, 0xd7, 0xfe, 0xde, 0xd9, 0x45, 0xb3, 0xf2, 0xe9, 0x7b, 0xb3, 0xbd, 0x06,
	0x79, 0x1e, 0x20, 0xfd, 0x3f, 0x69, 0xd1, 0x93, 0x4b, 0xd4, 0x0f, 0x4a, 0xa8, 0x0d, 0x93, 0xc1,
	0xee, 0x7e, 0xad, 0xc2, 0x2d, 0x8d, 0xfd, 0x8a, 0xa5, 0x59, 0x1c, 0x30, 0xf4, 0x01, 0xd8, 0x05,
	0x6b, 0x02, 0xed, 0xe0, 0x55, 0x83, 0x8d, 0x57, 0x0c, 0x53, 0x63, 0x77, 0x1d, 0xa9, 0x29, 0xef,
	0x3e, 0x7d, 0xf3, 0xe5, 0xe7, 0xbb, 0xaa, 0x87, 0x08, 0x59, 0xc6, 0x10, 0x7b, 0xb9, 0x32, 0x8f,
	0x58, 0x5f, 0xe4, 0xc4, 0x76, 0xe5, 0x94, 0x9c, 0xe8, 0x69, 0x3a, 0x45, 0x1f, 0x01, 0xbc, 0xf3,
	0xf7, 0x41, 0xa3, 0x4e, 0x49, 0xe5, 0xab, 0xf3, 0xd2, 0xc0, 0xeb, 0xca, 0x2d, 0xec, 0x9e, 0x86,
	0xdd, 0x45, 0xed, 0x52, 0x58, 0xb9, 0xa4, 0xed, 0xef, 0x9f, 0xcd, 0x1c, 0x70, 0x3e, 0x73, 0xc0,
	0x8f, 0x99, 0x03, 0xde, 0xce, 0x9d, 0xca, 0xf9, 0xdc, 0xa9, 0x7c, 0x9b, 0x3b, 0x95, 0xd7, 0x3b,
	0xa5, 0xcd, 0x2d, 0x7e, 0x5e, 0x06, 0x37, 0xf5, 0xd5, 0x7d, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0xc7, 0xf5, 0x9c, 0x75, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// QueryBalance queries the balance of a single coin for a single account
	QueryBalance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// QueryBalance queries the balance of all coins for a single account
	QueryAllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryBalance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos_sdk.x.bank.v1.QueryService/QueryBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryAllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error) {
	out := new(QueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos_sdk.x.bank.v1.QueryService/QueryAllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// QueryBalance queries the balance of a single coin for a single account
	QueryBalance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// QueryBalance queries the balance of all coins for a single account
	QueryAllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryBalance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBalance not implemented")
}
func (*UnimplementedQueryServiceServer) QueryAllBalances(ctx context.Context, req *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllBalances not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos_sdk.x.bank.v1.QueryService/QueryBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryBalance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryAllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryAllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos_sdk.x.bank.v1.QueryService/QueryAllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryAllBalances(ctx, req.(*QueryAllBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos_sdk.x.bank.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryBalance",
			Handler:    _QueryService_QueryBalance_Handler,
		},
		{
			MethodName: "QueryAllBalances",
			Handler:    _QueryService_QueryAllBalances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/bank/types/query.proto",
}

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Balance != nil {
		{
			size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Page != nil {
		{
			size, err := m.Page.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Page != nil {
		{
			size, err := m.Page.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Balance != nil {
		l = m.Balance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Page != nil {
		l = m.Page.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Page != nil {
		l = m.Page.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Balance == nil {
				m.Balance = &types.Coin{}
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllBalancesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Page == nil {
				m.Page = &query.PageRequest{}
			}
			if err := m.Page.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllBalancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, types.Coin{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Page == nil {
				m.Page = &query.PageResponse{}
			}
			if err := m.Page.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
