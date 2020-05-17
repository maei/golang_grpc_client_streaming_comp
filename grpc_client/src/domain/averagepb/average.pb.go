// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: grpc_client/src/domain/averagepb/average.proto

package averagepb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input int32 `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_client_src_domain_averagepb_average_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_client_src_domain_averagepb_average_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_client_src_domain_averagepb_average_proto_rawDescGZIP(), []int{0}
}

func (x *AverageRequest) GetInput() int32 {
	if x != nil {
		return x.Input
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_client_src_domain_averagepb_average_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_client_src_domain_averagepb_average_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_client_src_domain_averagepb_average_proto_rawDescGZIP(), []int{1}
}

func (x *AverageResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_grpc_client_src_domain_averagepb_average_proto protoreflect.FileDescriptor

var file_grpc_client_src_domain_averagepb_average_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x70, 0x62, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x59, 0x0a, 0x0e,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x2e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_client_src_domain_averagepb_average_proto_rawDescOnce sync.Once
	file_grpc_client_src_domain_averagepb_average_proto_rawDescData = file_grpc_client_src_domain_averagepb_average_proto_rawDesc
)

func file_grpc_client_src_domain_averagepb_average_proto_rawDescGZIP() []byte {
	file_grpc_client_src_domain_averagepb_average_proto_rawDescOnce.Do(func() {
		file_grpc_client_src_domain_averagepb_average_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_client_src_domain_averagepb_average_proto_rawDescData)
	})
	return file_grpc_client_src_domain_averagepb_average_proto_rawDescData
}

var file_grpc_client_src_domain_averagepb_average_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_client_src_domain_averagepb_average_proto_goTypes = []interface{}{
	(*AverageRequest)(nil),  // 0: average.AverageRequest
	(*AverageResponse)(nil), // 1: average.AverageResponse
}
var file_grpc_client_src_domain_averagepb_average_proto_depIdxs = []int32{
	0, // 0: average.AverageService.ComputeAverage:input_type -> average.AverageRequest
	1, // 1: average.AverageService.ComputeAverage:output_type -> average.AverageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_client_src_domain_averagepb_average_proto_init() }
func file_grpc_client_src_domain_averagepb_average_proto_init() {
	if File_grpc_client_src_domain_averagepb_average_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_client_src_domain_averagepb_average_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_grpc_client_src_domain_averagepb_average_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
			RawDescriptor: file_grpc_client_src_domain_averagepb_average_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_client_src_domain_averagepb_average_proto_goTypes,
		DependencyIndexes: file_grpc_client_src_domain_averagepb_average_proto_depIdxs,
		MessageInfos:      file_grpc_client_src_domain_averagepb_average_proto_msgTypes,
	}.Build()
	File_grpc_client_src_domain_averagepb_average_proto = out.File
	file_grpc_client_src_domain_averagepb_average_proto_rawDesc = nil
	file_grpc_client_src_domain_averagepb_average_proto_goTypes = nil
	file_grpc_client_src_domain_averagepb_average_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AverageServiceClient is the client API for AverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AverageServiceClient interface {
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AverageService_ComputeAverageClient, error)
}

type averageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAverageServiceClient(cc grpc.ClientConnInterface) AverageServiceClient {
	return &averageServiceClient{cc}
}

func (c *averageServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AverageService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AverageService_serviceDesc.Streams[0], "/average.AverageService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &averageServiceComputeAverageClient{stream}
	return x, nil
}

type AverageService_ComputeAverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type averageServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *averageServiceComputeAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *averageServiceComputeAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AverageServiceServer is the server API for AverageService service.
type AverageServiceServer interface {
	ComputeAverage(AverageService_ComputeAverageServer) error
}

// UnimplementedAverageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAverageServiceServer struct {
}

func (*UnimplementedAverageServiceServer) ComputeAverage(AverageService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterAverageServiceServer(s *grpc.Server, srv AverageServiceServer) {
	s.RegisterService(&_AverageService_serviceDesc, srv)
}

func _AverageService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AverageServiceServer).ComputeAverage(&averageServiceComputeAverageServer{stream})
}

type AverageService_ComputeAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type averageServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *averageServiceComputeAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *averageServiceComputeAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AverageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "average.AverageService",
	HandlerType: (*AverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeAverage",
			Handler:       _AverageService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_client/src/domain/averagepb/average.proto",
}