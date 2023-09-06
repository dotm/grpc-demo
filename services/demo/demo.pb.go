//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative services/demo/demo.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: services/demo/demo.proto

package demo

import (
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

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestMessage string `protobuf:"bytes,1,opt,name=requestMessage,proto3" json:"requestMessage,omitempty"`
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_demo_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_demo_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_services_demo_demo_proto_rawDescGZIP(), []int{0}
}

func (x *DemoRequest) GetRequestMessage() string {
	if x != nil {
		return x.RequestMessage
	}
	return ""
}

type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseMessage string `protobuf:"bytes,1,opt,name=responseMessage,proto3" json:"responseMessage,omitempty"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_demo_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_demo_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_services_demo_demo_proto_rawDescGZIP(), []int{1}
}

func (x *DemoResponse) GetResponseMessage() string {
	if x != nil {
		return x.ResponseMessage
	}
	return ""
}

var File_services_demo_demo_proto protoreflect.FileDescriptor

var file_services_demo_demo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f,
	0x22, 0x35, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x99, 0x02, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x15,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x4b, 0x0a, 0x1c, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x74, 0x6d,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_demo_demo_proto_rawDescOnce sync.Once
	file_services_demo_demo_proto_rawDescData = file_services_demo_demo_proto_rawDesc
)

func file_services_demo_demo_proto_rawDescGZIP() []byte {
	file_services_demo_demo_proto_rawDescOnce.Do(func() {
		file_services_demo_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_demo_demo_proto_rawDescData)
	})
	return file_services_demo_demo_proto_rawDescData
}

var file_services_demo_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_demo_demo_proto_goTypes = []interface{}{
	(*DemoRequest)(nil),  // 0: demo.DemoRequest
	(*DemoResponse)(nil), // 1: demo.DemoResponse
}
var file_services_demo_demo_proto_depIdxs = []int32{
	0, // 0: demo.Demo.SimpleUnaryMethod:input_type -> demo.DemoRequest
	0, // 1: demo.Demo.ServerStreamingMethod:input_type -> demo.DemoRequest
	0, // 2: demo.Demo.ClientStreamingMethod:input_type -> demo.DemoRequest
	0, // 3: demo.Demo.BidirectionalStreamingMethod:input_type -> demo.DemoRequest
	1, // 4: demo.Demo.SimpleUnaryMethod:output_type -> demo.DemoResponse
	1, // 5: demo.Demo.ServerStreamingMethod:output_type -> demo.DemoResponse
	1, // 6: demo.Demo.ClientStreamingMethod:output_type -> demo.DemoResponse
	1, // 7: demo.Demo.BidirectionalStreamingMethod:output_type -> demo.DemoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_demo_demo_proto_init() }
func file_services_demo_demo_proto_init() {
	if File_services_demo_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_demo_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoRequest); i {
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
		file_services_demo_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoResponse); i {
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
			RawDescriptor: file_services_demo_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_demo_demo_proto_goTypes,
		DependencyIndexes: file_services_demo_demo_proto_depIdxs,
		MessageInfos:      file_services_demo_demo_proto_msgTypes,
	}.Build()
	File_services_demo_demo_proto = out.File
	file_services_demo_demo_proto_rawDesc = nil
	file_services_demo_demo_proto_goTypes = nil
	file_services_demo_demo_proto_depIdxs = nil
}