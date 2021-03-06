//
// Copyright (C) 2021 Grakn Labs
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
// source: v2/protobuf/answer.proto

package grakn_protocol

import (
	proto "github.com/golang/protobuf/proto"
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

type ConceptMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map     map[string]*Concept `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pattern string              `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *ConceptMap) Reset() {
	*x = ConceptMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptMap) ProtoMessage() {}

func (x *ConceptMap) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptMap.ProtoReflect.Descriptor instead.
func (*ConceptMap) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_answer_proto_rawDescGZIP(), []int{0}
}

func (x *ConceptMap) GetMap() map[string]*Concept {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *ConceptMap) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

type ConceptMapGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       *Concept      `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConceptMaps []*ConceptMap `protobuf:"bytes,2,rep,name=concept_maps,json=conceptMaps,proto3" json:"concept_maps,omitempty"`
}

func (x *ConceptMapGroup) Reset() {
	*x = ConceptMapGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptMapGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptMapGroup) ProtoMessage() {}

func (x *ConceptMapGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptMapGroup.ProtoReflect.Descriptor instead.
func (*ConceptMapGroup) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_answer_proto_rawDescGZIP(), []int{1}
}

func (x *ConceptMapGroup) GetOwner() *Concept {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ConceptMapGroup) GetConceptMaps() []*ConceptMap {
	if x != nil {
		return x.ConceptMaps
	}
	return nil
}

type Numeric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Numeric_LongValue
	//	*Numeric_DoubleValue
	//	*Numeric_Nan
	Value isNumeric_Value `protobuf_oneof:"value"`
}

func (x *Numeric) Reset() {
	*x = Numeric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numeric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numeric) ProtoMessage() {}

func (x *Numeric) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numeric.ProtoReflect.Descriptor instead.
func (*Numeric) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_answer_proto_rawDescGZIP(), []int{2}
}

func (m *Numeric) GetValue() isNumeric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Numeric) GetLongValue() int64 {
	if x, ok := x.GetValue().(*Numeric_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *Numeric) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*Numeric_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Numeric) GetNan() bool {
	if x, ok := x.GetValue().(*Numeric_Nan); ok {
		return x.Nan
	}
	return false
}

type isNumeric_Value interface {
	isNumeric_Value()
}

type Numeric_LongValue struct {
	LongValue int64 `protobuf:"varint,1,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Numeric_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Numeric_Nan struct {
	Nan bool `protobuf:"varint,3,opt,name=nan,proto3,oneof"`
}

func (*Numeric_LongValue) isNumeric_Value() {}

func (*Numeric_DoubleValue) isNumeric_Value() {}

func (*Numeric_Nan) isNumeric_Value() {}

type NumericGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  *Concept `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Number *Numeric `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumericGroup) Reset() {
	*x = NumericGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericGroup) ProtoMessage() {}

func (x *NumericGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumericGroup.ProtoReflect.Descriptor instead.
func (*NumericGroup) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_answer_proto_rawDescGZIP(), []int{3}
}

func (x *NumericGroup) GetOwner() *Concept {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *NumericGroup) GetNumber() *Numeric {
	if x != nil {
		return x.Number
	}
	return nil
}

var File_v2_protobuf_answer_proto protoreflect.FileDescriptor

var file_v2_protobuf_answer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x61, 0x6b,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x19, 0x76, 0x32, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x4d, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x1a, 0x4f, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x4d, 0x61, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x73, 0x22, 0x6c, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x6e, 0x61, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x6e, 0x61, 0x6e, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6e, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x41, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_answer_proto_rawDescOnce sync.Once
	file_v2_protobuf_answer_proto_rawDescData = file_v2_protobuf_answer_proto_rawDesc
)

func file_v2_protobuf_answer_proto_rawDescGZIP() []byte {
	file_v2_protobuf_answer_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_answer_proto_rawDescData)
	})
	return file_v2_protobuf_answer_proto_rawDescData
}

var file_v2_protobuf_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v2_protobuf_answer_proto_goTypes = []interface{}{
	(*ConceptMap)(nil),      // 0: grakn.protocol.ConceptMap
	(*ConceptMapGroup)(nil), // 1: grakn.protocol.ConceptMapGroup
	(*Numeric)(nil),         // 2: grakn.protocol.Numeric
	(*NumericGroup)(nil),    // 3: grakn.protocol.NumericGroup
	nil,                     // 4: grakn.protocol.ConceptMap.MapEntry
	(*Concept)(nil),         // 5: grakn.protocol.Concept
}
var file_v2_protobuf_answer_proto_depIdxs = []int32{
	4, // 0: grakn.protocol.ConceptMap.map:type_name -> grakn.protocol.ConceptMap.MapEntry
	5, // 1: grakn.protocol.ConceptMapGroup.owner:type_name -> grakn.protocol.Concept
	0, // 2: grakn.protocol.ConceptMapGroup.concept_maps:type_name -> grakn.protocol.ConceptMap
	5, // 3: grakn.protocol.NumericGroup.owner:type_name -> grakn.protocol.Concept
	2, // 4: grakn.protocol.NumericGroup.number:type_name -> grakn.protocol.Numeric
	5, // 5: grakn.protocol.ConceptMap.MapEntry.value:type_name -> grakn.protocol.Concept
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v2_protobuf_answer_proto_init() }
func file_v2_protobuf_answer_proto_init() {
	if File_v2_protobuf_answer_proto != nil {
		return
	}
	file_v2_protobuf_concept_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptMap); i {
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
		file_v2_protobuf_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptMapGroup); i {
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
		file_v2_protobuf_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numeric); i {
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
		file_v2_protobuf_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericGroup); i {
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
	file_v2_protobuf_answer_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Numeric_LongValue)(nil),
		(*Numeric_DoubleValue)(nil),
		(*Numeric_Nan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2_protobuf_answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_answer_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_answer_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_answer_proto_msgTypes,
	}.Build()
	File_v2_protobuf_answer_proto = out.File
	file_v2_protobuf_answer_proto_rawDesc = nil
	file_v2_protobuf_answer_proto_goTypes = nil
	file_v2_protobuf_answer_proto_depIdxs = nil
}
