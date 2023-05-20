// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: app/user/api/data.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProjectType int32

const (
	ProjectType_UNKNOWN ProjectType = 0
	ProjectType_LANDING ProjectType = 1
	ProjectType_CMS     ProjectType = 2
)

// Enum value maps for ProjectType.
var (
	ProjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LANDING",
		2: "CMS",
	}
	ProjectType_value = map[string]int32{
		"UNKNOWN": 0,
		"LANDING": 1,
		"CMS":     2,
	}
)

func (x ProjectType) Enum() *ProjectType {
	p := new(ProjectType)
	*p = x
	return p
}

func (x ProjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_user_api_data_proto_enumTypes[0].Descriptor()
}

func (ProjectType) Type() protoreflect.EnumType {
	return &file_app_user_api_data_proto_enumTypes[0]
}

func (x ProjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectType.Descriptor instead.
func (ProjectType) EnumDescriptor() ([]byte, []int) {
	return file_app_user_api_data_proto_rawDescGZIP(), []int{0}
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CompressString string `protobuf:"bytes,3,opt,name=compress_string,json=compressString,proto3" json:"compress_string,omitempty"`
	CreatedAt      int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      int64  `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_user_api_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_api_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_app_user_api_data_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetCompressString() string {
	if x != nil {
		return x.CompressString
	}
	return ""
}

func (x *Project) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Project) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_app_user_api_data_proto protoreflect.FileDescriptor

var file_app_user_api_data_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x66, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x30, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x41, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x43, 0x4d, 0x53, 0x10, 0x02, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x4b, 0x31, 0x39, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x66, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_user_api_data_proto_rawDescOnce sync.Once
	file_app_user_api_data_proto_rawDescData = file_app_user_api_data_proto_rawDesc
)

func file_app_user_api_data_proto_rawDescGZIP() []byte {
	file_app_user_api_data_proto_rawDescOnce.Do(func() {
		file_app_user_api_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_user_api_data_proto_rawDescData)
	})
	return file_app_user_api_data_proto_rawDescData
}

var file_app_user_api_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_user_api_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_user_api_data_proto_goTypes = []interface{}{
	(ProjectType)(0), // 0: buildify.app.user.api.ProjectType
	(*Project)(nil),  // 1: buildify.app.user.api.Project
}
var file_app_user_api_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_user_api_data_proto_init() }
func file_app_user_api_data_proto_init() {
	if File_app_user_api_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_user_api_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_app_user_api_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_user_api_data_proto_goTypes,
		DependencyIndexes: file_app_user_api_data_proto_depIdxs,
		EnumInfos:         file_app_user_api_data_proto_enumTypes,
		MessageInfos:      file_app_user_api_data_proto_msgTypes,
	}.Build()
	File_app_user_api_data_proto = out.File
	file_app_user_api_data_proto_rawDesc = nil
	file_app_user_api_data_proto_goTypes = nil
	file_app_user_api_data_proto_depIdxs = nil
}
