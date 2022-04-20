// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: stock.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId int32 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Stock     int32 `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Lock      int32 `protobuf:"varint,4,opt,name=lock,proto3" json:"lock,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStockRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStockRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *UpdateStockRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *UpdateStockRequest) GetLock() int32 {
	if x != nil {
		return x.Lock
	}
	return 0
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x4c, 0x0a, 0x0c, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stock_proto_goTypes = []interface{}{
	(*UpdateStockRequest)(nil), // 0: UpdateStockRequest
	(*emptypb.Empty)(nil),      // 1: google.protobuf.Empty
}
var file_stock_proto_depIdxs = []int32{
	0, // 0: StockService.UpdateStock:input_type -> UpdateStockRequest
	1, // 1: StockService.UpdateStock:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockServiceClient interface {
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/StockService/UpdateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
type StockServiceServer interface {
	UpdateStock(context.Context, *UpdateStockRequest) (*emptypb.Empty, error)
}

// UnimplementedStockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (*UnimplementedStockServiceServer) UpdateStock(context.Context, *UpdateStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}

func RegisterStockServiceServer(s *grpc.Server, srv StockServiceServer) {
	s.RegisterService(&_StockService_serviceDesc, srv)
}

func _StockService_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StockService/UpdateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStock",
			Handler:    _StockService_UpdateStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}