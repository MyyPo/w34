// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: dev/v1/dev.proto

package devv1

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

type NewProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Public bool   `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *NewProjectRequest) Reset() {
	*x = NewProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProjectRequest) ProtoMessage() {}

func (x *NewProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProjectRequest.ProtoReflect.Descriptor instead.
func (*NewProjectRequest) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{0}
}

func (x *NewProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewProjectRequest) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

type NewProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewProjectResponse) Reset() {
	*x = NewProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProjectResponse) ProtoMessage() {}

func (x *NewProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProjectResponse.ProtoReflect.Descriptor instead.
func (*NewProjectResponse) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{1}
}

type DeleteProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProjectResponse) Reset() {
	*x = DeleteProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectResponse) ProtoMessage() {}

func (x *DeleteProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteProjectResponse) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{3}
}

type NewLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName  string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	LocationName string `protobuf:"bytes,2,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
}

func (x *NewLocationRequest) Reset() {
	*x = NewLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLocationRequest) ProtoMessage() {}

func (x *NewLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLocationRequest.ProtoReflect.Descriptor instead.
func (*NewLocationRequest) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{4}
}

func (x *NewLocationRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *NewLocationRequest) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

type NewLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewLocationResponse) Reset() {
	*x = NewLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dev_v1_dev_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLocationResponse) ProtoMessage() {}

func (x *NewLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dev_v1_dev_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLocationResponse.ProtoReflect.Descriptor instead.
func (*NewLocationResponse) Descriptor() ([]byte, []int) {
	return file_dev_v1_dev_proto_rawDescGZIP(), []int{5}
}

var File_dev_v1_dev_proto protoreflect.FileDescriptor

var file_dev_v1_dev_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x22, 0x3f, 0x0a, 0x11, 0x4e, 0x65,
	0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x14, 0x0a, 0x12, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x44, 0x65, 0x76, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x79, 0x79, 0x50, 0x6f, 0x2f, 0x77, 0x33, 0x34, 0x2e, 0x47, 0x6f, 0x2f, 0x64,
	0x65, 0x76, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58,
	0x58, 0xaa, 0x02, 0x06, 0x44, 0x65, 0x76, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x44, 0x65, 0x76,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x44, 0x65, 0x76, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x44, 0x65, 0x76, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dev_v1_dev_proto_rawDescOnce sync.Once
	file_dev_v1_dev_proto_rawDescData = file_dev_v1_dev_proto_rawDesc
)

func file_dev_v1_dev_proto_rawDescGZIP() []byte {
	file_dev_v1_dev_proto_rawDescOnce.Do(func() {
		file_dev_v1_dev_proto_rawDescData = protoimpl.X.CompressGZIP(file_dev_v1_dev_proto_rawDescData)
	})
	return file_dev_v1_dev_proto_rawDescData
}

var file_dev_v1_dev_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dev_v1_dev_proto_goTypes = []interface{}{
	(*NewProjectRequest)(nil),     // 0: dev.v1.NewProjectRequest
	(*NewProjectResponse)(nil),    // 1: dev.v1.NewProjectResponse
	(*DeleteProjectRequest)(nil),  // 2: dev.v1.DeleteProjectRequest
	(*DeleteProjectResponse)(nil), // 3: dev.v1.DeleteProjectResponse
	(*NewLocationRequest)(nil),    // 4: dev.v1.NewLocationRequest
	(*NewLocationResponse)(nil),   // 5: dev.v1.NewLocationResponse
}
var file_dev_v1_dev_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dev_v1_dev_proto_init() }
func file_dev_v1_dev_proto_init() {
	if File_dev_v1_dev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dev_v1_dev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProjectRequest); i {
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
		file_dev_v1_dev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProjectResponse); i {
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
		file_dev_v1_dev_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectRequest); i {
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
		file_dev_v1_dev_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectResponse); i {
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
		file_dev_v1_dev_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLocationRequest); i {
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
		file_dev_v1_dev_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLocationResponse); i {
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
			RawDescriptor: file_dev_v1_dev_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dev_v1_dev_proto_goTypes,
		DependencyIndexes: file_dev_v1_dev_proto_depIdxs,
		MessageInfos:      file_dev_v1_dev_proto_msgTypes,
	}.Build()
	File_dev_v1_dev_proto = out.File
	file_dev_v1_dev_proto_rawDesc = nil
	file_dev_v1_dev_proto_goTypes = nil
	file_dev_v1_dev_proto_depIdxs = nil
}
