// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/proto/server/server.proto

package server

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

type FindPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartPageUrl  string `protobuf:"bytes,1,opt,name=start_page_url,json=startPageUrl,proto3" json:"start_page_url,omitempty"`
	FinishPageUrl string `protobuf:"bytes,2,opt,name=finish_page_url,json=finishPageUrl,proto3" json:"finish_page_url,omitempty"`
}

func (x *FindPathRequest) Reset() {
	*x = FindPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_server_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPathRequest) ProtoMessage() {}

func (x *FindPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_server_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPathRequest.ProtoReflect.Descriptor instead.
func (*FindPathRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_server_server_proto_rawDescGZIP(), []int{0}
}

func (x *FindPathRequest) GetStartPageUrl() string {
	if x != nil {
		return x.StartPageUrl
	}
	return ""
}

func (x *FindPathRequest) GetFinishPageUrl() string {
	if x != nil {
		return x.FinishPageUrl
	}
	return ""
}

type FindPathResultId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultId string `protobuf:"bytes,1,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
}

func (x *FindPathResultId) Reset() {
	*x = FindPathResultId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_server_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPathResultId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPathResultId) ProtoMessage() {}

func (x *FindPathResultId) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_server_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPathResultId.ProtoReflect.Descriptor instead.
func (*FindPathResultId) Descriptor() ([]byte, []int) {
	return file_pkg_proto_server_server_proto_rawDescGZIP(), []int{1}
}

func (x *FindPathResultId) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

type FindPathResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathFound bool     `protobuf:"varint,1,opt,name=path_found,json=pathFound,proto3" json:"path_found,omitempty"`
	Path      []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *FindPathResult) Reset() {
	*x = FindPathResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_server_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPathResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPathResult) ProtoMessage() {}

func (x *FindPathResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_server_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPathResult.ProtoReflect.Descriptor instead.
func (*FindPathResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_server_server_proto_rawDescGZIP(), []int{2}
}

func (x *FindPathResult) GetPathFound() bool {
	if x != nil {
		return x.PathFound
	}
	return false
}

func (x *FindPathResult) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_pkg_proto_server_server_proto protoreflect.FileDescriptor

var file_pkg_proto_server_server_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x50, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0e, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0x99,
	0x01, 0x0a, 0x14, 0x57, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68,
	0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x49, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x76, 0x69, 0x63, 0x68, 0x6b, 0x6f, 0x76, 0x2f, 0x77, 0x69, 0x6b, 0x69,
	0x70, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_server_server_proto_rawDescOnce sync.Once
	file_pkg_proto_server_server_proto_rawDescData = file_pkg_proto_server_server_proto_rawDesc
)

func file_pkg_proto_server_server_proto_rawDescGZIP() []byte {
	file_pkg_proto_server_server_proto_rawDescOnce.Do(func() {
		file_pkg_proto_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_server_server_proto_rawDescData)
	})
	return file_pkg_proto_server_server_proto_rawDescData
}

var file_pkg_proto_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_server_server_proto_goTypes = []interface{}{
	(*FindPathRequest)(nil),  // 0: server.FindPathRequest
	(*FindPathResultId)(nil), // 1: server.FindPathResultId
	(*FindPathResult)(nil),   // 2: server.FindPathResult
}
var file_pkg_proto_server_server_proto_depIdxs = []int32{
	0, // 0: server.WikipediaPathfinding.QueueFindPath:input_type -> server.FindPathRequest
	1, // 1: server.WikipediaPathfinding.GetResult:input_type -> server.FindPathResultId
	1, // 2: server.WikipediaPathfinding.QueueFindPath:output_type -> server.FindPathResultId
	2, // 3: server.WikipediaPathfinding.GetResult:output_type -> server.FindPathResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_server_server_proto_init() }
func file_pkg_proto_server_server_proto_init() {
	if File_pkg_proto_server_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_server_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPathRequest); i {
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
		file_pkg_proto_server_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPathResultId); i {
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
		file_pkg_proto_server_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPathResult); i {
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
			RawDescriptor: file_pkg_proto_server_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_server_server_proto_goTypes,
		DependencyIndexes: file_pkg_proto_server_server_proto_depIdxs,
		MessageInfos:      file_pkg_proto_server_server_proto_msgTypes,
	}.Build()
	File_pkg_proto_server_server_proto = out.File
	file_pkg_proto_server_server_proto_rawDesc = nil
	file_pkg_proto_server_server_proto_goTypes = nil
	file_pkg_proto_server_server_proto_depIdxs = nil
}
