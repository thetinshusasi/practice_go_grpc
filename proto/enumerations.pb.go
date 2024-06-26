// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: enumerations.proto

package proto

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

type Color int32

const (
	Color_NONE  Color = 0
	Color_RED   Color = 1
	Color_GREEN Color = 2
	Color_BLUE  Color = 3
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "NONE",
		1: "RED",
		2: "GREEN",
		3: "BLUE",
	}
	Color_value = map[string]int32{
		"NONE":  0,
		"RED":   1,
		"GREEN": 2,
		"BLUE":  3,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_enumerations_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_enumerations_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_enumerations_proto_rawDescGZIP(), []int{0}
}

type Enumeration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EyeColor Color `protobuf:"varint,1,opt,name=eye_color,json=eyeColor,proto3,enum=proto.enumerations.Color" json:"eye_color,omitempty"`
}

func (x *Enumeration) Reset() {
	*x = Enumeration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enumerations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enumeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enumeration) ProtoMessage() {}

func (x *Enumeration) ProtoReflect() protoreflect.Message {
	mi := &file_enumerations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enumeration.ProtoReflect.Descriptor instead.
func (*Enumeration) Descriptor() ([]byte, []int) {
	return file_enumerations_proto_rawDescGZIP(), []int{0}
}

func (x *Enumeration) GetEyeColor() Color {
	if x != nil {
		return x.EyeColor
	}
	return Color_NONE
}

var File_enumerations_proto protoreflect.FileDescriptor

var file_enumerations_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x09, 0x65, 0x79, 0x65, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x2a,
	0x2f, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x03,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_enumerations_proto_rawDescOnce sync.Once
	file_enumerations_proto_rawDescData = file_enumerations_proto_rawDesc
)

func file_enumerations_proto_rawDescGZIP() []byte {
	file_enumerations_proto_rawDescOnce.Do(func() {
		file_enumerations_proto_rawDescData = protoimpl.X.CompressGZIP(file_enumerations_proto_rawDescData)
	})
	return file_enumerations_proto_rawDescData
}

var file_enumerations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enumerations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enumerations_proto_goTypes = []interface{}{
	(Color)(0),          // 0: proto.enumerations.Color
	(*Enumeration)(nil), // 1: proto.enumerations.Enumeration
}
var file_enumerations_proto_depIdxs = []int32{
	0, // 0: proto.enumerations.Enumeration.eye_color:type_name -> proto.enumerations.Color
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enumerations_proto_init() }
func file_enumerations_proto_init() {
	if File_enumerations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enumerations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enumeration); i {
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
			RawDescriptor: file_enumerations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enumerations_proto_goTypes,
		DependencyIndexes: file_enumerations_proto_depIdxs,
		EnumInfos:         file_enumerations_proto_enumTypes,
		MessageInfos:      file_enumerations_proto_msgTypes,
	}.Build()
	File_enumerations_proto = out.File
	file_enumerations_proto_rawDesc = nil
	file_enumerations_proto_goTypes = nil
	file_enumerations_proto_depIdxs = nil
}
