// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: auth/v1/auth_service.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_v1_auth_service_proto protoreflect.FileDescriptor

var file_auth_v1_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd3, 0x01, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x16,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x84, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42,
	0x10, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x79, 0x79, 0x50, 0x6f, 0x2f, 0x77, 0x33, 0x34, 0x2e, 0x47, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75,
	0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75,
	0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_v1_auth_service_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),         // 0: auth.v1.SignUpRequest
	(*SignInRequest)(nil),         // 1: auth.v1.SignInRequest
	(*RefreshTokensRequest)(nil),  // 2: auth.v1.RefreshTokensRequest
	(*SignUpResponse)(nil),        // 3: auth.v1.SignUpResponse
	(*SignInResponse)(nil),        // 4: auth.v1.SignInResponse
	(*RefreshTokensResponse)(nil), // 5: auth.v1.RefreshTokensResponse
}
var file_auth_v1_auth_service_proto_depIdxs = []int32{
	0, // 0: auth.v1.AuthService.SignUp:input_type -> auth.v1.SignUpRequest
	1, // 1: auth.v1.AuthService.SignIn:input_type -> auth.v1.SignInRequest
	2, // 2: auth.v1.AuthService.RefreshTokens:input_type -> auth.v1.RefreshTokensRequest
	3, // 3: auth.v1.AuthService.SignUp:output_type -> auth.v1.SignUpResponse
	4, // 4: auth.v1.AuthService.SignIn:output_type -> auth.v1.SignInResponse
	5, // 5: auth.v1.AuthService.RefreshTokens:output_type -> auth.v1.RefreshTokensResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_service_proto_init() }
func file_auth_v1_auth_service_proto_init() {
	if File_auth_v1_auth_service_proto != nil {
		return
	}
	file_auth_v1_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_service_proto_depIdxs,
	}.Build()
	File_auth_v1_auth_service_proto = out.File
	file_auth_v1_auth_service_proto_rawDesc = nil
	file_auth_v1_auth_service_proto_goTypes = nil
	file_auth_v1_auth_service_proto_depIdxs = nil
}
