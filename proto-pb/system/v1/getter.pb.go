// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.7
// source: system/v1/getter.proto

package pb_sys_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_system_v1_getter_proto protoreflect.FileDescriptor

var file_system_v1_getter_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x01,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f,
	0x73, 0x79, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_system_v1_getter_proto_goTypes = []any{
	(*PermissionGet_Request)(nil),   // 0: proto.system.v1.PermissionGet.Request
	(*PermissionList_Request)(nil),  // 1: proto.system.v1.PermissionList.Request
	(*PermissionGet_Response)(nil),  // 2: proto.system.v1.PermissionGet.Response
	(*PermissionList_Response)(nil), // 3: proto.system.v1.PermissionList.Response
}
var file_system_v1_getter_proto_depIdxs = []int32{
	0, // 0: proto.system.v1.Getter.PermissionGet:input_type -> proto.system.v1.PermissionGet.Request
	1, // 1: proto.system.v1.Getter.PermissionList:input_type -> proto.system.v1.PermissionList.Request
	2, // 2: proto.system.v1.Getter.PermissionGet:output_type -> proto.system.v1.PermissionGet.Response
	3, // 3: proto.system.v1.Getter.PermissionList:output_type -> proto.system.v1.PermissionList.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_v1_getter_proto_init() }
func file_system_v1_getter_proto_init() {
	if File_system_v1_getter_proto != nil {
		return
	}
	file_system_v1_message_permission_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_system_v1_getter_proto_rawDesc), len(file_system_v1_getter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_v1_getter_proto_goTypes,
		DependencyIndexes: file_system_v1_getter_proto_depIdxs,
	}.Build()
	File_system_v1_getter_proto = out.File
	file_system_v1_getter_proto_goTypes = nil
	file_system_v1_getter_proto_depIdxs = nil
}
