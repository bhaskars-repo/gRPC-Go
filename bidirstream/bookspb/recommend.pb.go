// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: bookspb/recommend.proto

package bookspb

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

type BookRecommendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *BookRecommendRequest) Reset() {
	*x = BookRecommendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookspb_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRecommendRequest) ProtoMessage() {}

func (x *BookRecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookspb_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRecommendRequest.ProtoReflect.Descriptor instead.
func (*BookRecommendRequest) Descriptor() ([]byte, []int) {
	return file_bookspb_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *BookRecommendRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type BookRecommendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Isbn  string `protobuf:"bytes,3,opt,name=isbn,proto3" json:"isbn,omitempty"`
}

func (x *BookRecommendResponse) Reset() {
	*x = BookRecommendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookspb_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRecommendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRecommendResponse) ProtoMessage() {}

func (x *BookRecommendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookspb_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRecommendResponse.ProtoReflect.Descriptor instead.
func (*BookRecommendResponse) Descriptor() ([]byte, []int) {
	return file_bookspb_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *BookRecommendResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *BookRecommendResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookRecommendResponse) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

var File_bookspb_recommend_proto protoreflect.FileDescriptor

var file_bookspb_recommend_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x22, 0x57, 0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x32, 0x71, 0x0a, 0x14, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3f,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x70, 0x61, 0x72, 0x63,
	0x2e, 0x62, 0x64, 0x73, 0x50, 0x01, 0x5a, 0x27, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x70, 0x61,
	0x72, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x69, 0x64, 0x69,
	0x72, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookspb_recommend_proto_rawDescOnce sync.Once
	file_bookspb_recommend_proto_rawDescData = file_bookspb_recommend_proto_rawDesc
)

func file_bookspb_recommend_proto_rawDescGZIP() []byte {
	file_bookspb_recommend_proto_rawDescOnce.Do(func() {
		file_bookspb_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookspb_recommend_proto_rawDescData)
	})
	return file_bookspb_recommend_proto_rawDescData
}

var file_bookspb_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bookspb_recommend_proto_goTypes = []interface{}{
	(*BookRecommendRequest)(nil),  // 0: recommend.BookRecommendRequest
	(*BookRecommendResponse)(nil), // 1: recommend.BookRecommendResponse
}
var file_bookspb_recommend_proto_depIdxs = []int32{
	0, // 0: recommend.BookRecommendService.recommendedBooks:input_type -> recommend.BookRecommendRequest
	1, // 1: recommend.BookRecommendService.recommendedBooks:output_type -> recommend.BookRecommendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bookspb_recommend_proto_init() }
func file_bookspb_recommend_proto_init() {
	if File_bookspb_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookspb_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRecommendRequest); i {
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
		file_bookspb_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRecommendResponse); i {
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
			RawDescriptor: file_bookspb_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookspb_recommend_proto_goTypes,
		DependencyIndexes: file_bookspb_recommend_proto_depIdxs,
		MessageInfos:      file_bookspb_recommend_proto_msgTypes,
	}.Build()
	File_bookspb_recommend_proto = out.File
	file_bookspb_recommend_proto_rawDesc = nil
	file_bookspb_recommend_proto_goTypes = nil
	file_bookspb_recommend_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookRecommendServiceClient is the client API for BookRecommendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookRecommendServiceClient interface {
	RecommendedBooks(ctx context.Context, opts ...grpc.CallOption) (BookRecommendService_RecommendedBooksClient, error)
}

type bookRecommendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookRecommendServiceClient(cc grpc.ClientConnInterface) BookRecommendServiceClient {
	return &bookRecommendServiceClient{cc}
}

func (c *bookRecommendServiceClient) RecommendedBooks(ctx context.Context, opts ...grpc.CallOption) (BookRecommendService_RecommendedBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BookRecommendService_serviceDesc.Streams[0], "/recommend.BookRecommendService/recommendedBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookRecommendServiceRecommendedBooksClient{stream}
	return x, nil
}

type BookRecommendService_RecommendedBooksClient interface {
	Send(*BookRecommendRequest) error
	Recv() (*BookRecommendResponse, error)
	grpc.ClientStream
}

type bookRecommendServiceRecommendedBooksClient struct {
	grpc.ClientStream
}

func (x *bookRecommendServiceRecommendedBooksClient) Send(m *BookRecommendRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookRecommendServiceRecommendedBooksClient) Recv() (*BookRecommendResponse, error) {
	m := new(BookRecommendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookRecommendServiceServer is the server API for BookRecommendService service.
type BookRecommendServiceServer interface {
	RecommendedBooks(BookRecommendService_RecommendedBooksServer) error
}

// UnimplementedBookRecommendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookRecommendServiceServer struct {
}

func (*UnimplementedBookRecommendServiceServer) RecommendedBooks(BookRecommendService_RecommendedBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method RecommendedBooks not implemented")
}

func RegisterBookRecommendServiceServer(s *grpc.Server, srv BookRecommendServiceServer) {
	s.RegisterService(&_BookRecommendService_serviceDesc, srv)
}

func _BookRecommendService_RecommendedBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookRecommendServiceServer).RecommendedBooks(&bookRecommendServiceRecommendedBooksServer{stream})
}

type BookRecommendService_RecommendedBooksServer interface {
	Send(*BookRecommendResponse) error
	Recv() (*BookRecommendRequest, error)
	grpc.ServerStream
}

type bookRecommendServiceRecommendedBooksServer struct {
	grpc.ServerStream
}

func (x *bookRecommendServiceRecommendedBooksServer) Send(m *BookRecommendResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookRecommendServiceRecommendedBooksServer) Recv() (*BookRecommendRequest, error) {
	m := new(BookRecommendRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BookRecommendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.BookRecommendService",
	HandlerType: (*BookRecommendServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "recommendedBooks",
			Handler:       _BookRecommendService_RecommendedBooks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bookspb/recommend.proto",
}