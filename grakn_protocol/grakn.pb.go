//
// Copyright (C) 2020 Grakn Labs
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: protobuf/grakn.proto

package grakn_protocol

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_protobuf_grakn_proto protoreflect.FileDescriptor

var file_protobuf_grakn_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xf8, 0x04, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x6b, 0x6e, 0x12, 0x61, 0x0a,
	0x11, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x6b,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x12, 0x5b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x12, 0x52, 0x0a,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x20, 0x2e,
	0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x1a,
	0x20, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x12, 0x5b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x6b,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x12, 0x52,
	0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x20,
	0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x1a, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x12, 0x55, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x61, 0x6b,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3d,
	0x0a, 0x0e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x42, 0x0a, 0x47, 0x72, 0x61, 0x6b, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1f, 0x2e, 0x2f,
	0x67, 0x72, 0x61, 0x6b, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x67,
	0x72, 0x61, 0x6b, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protobuf_grakn_proto_goTypes = []interface{}{
	(*Database_Contains_Req)(nil), // 0: grakn.protocol.Database.Contains.Req
	(*Database_Create_Req)(nil),   // 1: grakn.protocol.Database.Create.Req
	(*Database_All_Req)(nil),      // 2: grakn.protocol.Database.All.Req
	(*Database_Delete_Req)(nil),   // 3: grakn.protocol.Database.Delete.Req
	(*Session_Open_Req)(nil),      // 4: grakn.protocol.Session.Open.Req
	(*Session_Close_Req)(nil),     // 5: grakn.protocol.Session.Close.Req
	(*Transaction_Req)(nil),       // 6: grakn.protocol.Transaction.Req
	(*Database_Contains_Res)(nil), // 7: grakn.protocol.Database.Contains.Res
	(*Database_Create_Res)(nil),   // 8: grakn.protocol.Database.Create.Res
	(*Database_All_Res)(nil),      // 9: grakn.protocol.Database.All.Res
	(*Database_Delete_Res)(nil),   // 10: grakn.protocol.Database.Delete.Res
	(*Session_Open_Res)(nil),      // 11: grakn.protocol.Session.Open.Res
	(*Session_Close_Res)(nil),     // 12: grakn.protocol.Session.Close.Res
	(*Transaction_Res)(nil),       // 13: grakn.protocol.Transaction.Res
}
var file_protobuf_grakn_proto_depIdxs = []int32{
	0,  // 0: grakn.protocol.Grakn.database_contains:input_type -> grakn.protocol.Database.Contains.Req
	1,  // 1: grakn.protocol.Grakn.database_create:input_type -> grakn.protocol.Database.Create.Req
	2,  // 2: grakn.protocol.Grakn.database_all:input_type -> grakn.protocol.Database.All.Req
	3,  // 3: grakn.protocol.Grakn.database_delete:input_type -> grakn.protocol.Database.Delete.Req
	4,  // 4: grakn.protocol.Grakn.session_open:input_type -> grakn.protocol.Session.Open.Req
	5,  // 5: grakn.protocol.Grakn.session_close:input_type -> grakn.protocol.Session.Close.Req
	6,  // 6: grakn.protocol.Grakn.transaction:input_type -> grakn.protocol.Transaction.Req
	7,  // 7: grakn.protocol.Grakn.database_contains:output_type -> grakn.protocol.Database.Contains.Res
	8,  // 8: grakn.protocol.Grakn.database_create:output_type -> grakn.protocol.Database.Create.Res
	9,  // 9: grakn.protocol.Grakn.database_all:output_type -> grakn.protocol.Database.All.Res
	10, // 10: grakn.protocol.Grakn.database_delete:output_type -> grakn.protocol.Database.Delete.Res
	11, // 11: grakn.protocol.Grakn.session_open:output_type -> grakn.protocol.Session.Open.Res
	12, // 12: grakn.protocol.Grakn.session_close:output_type -> grakn.protocol.Session.Close.Res
	13, // 13: grakn.protocol.Grakn.transaction:output_type -> grakn.protocol.Transaction.Res
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_grakn_proto_init() }
func file_protobuf_grakn_proto_init() {
	if File_protobuf_grakn_proto != nil {
		return
	}
	file_protobuf_database_proto_init()
	file_protobuf_session_proto_init()
	file_protobuf_transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_grakn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_grakn_proto_goTypes,
		DependencyIndexes: file_protobuf_grakn_proto_depIdxs,
	}.Build()
	File_protobuf_grakn_proto = out.File
	file_protobuf_grakn_proto_rawDesc = nil
	file_protobuf_grakn_proto_goTypes = nil
	file_protobuf_grakn_proto_depIdxs = nil
}
