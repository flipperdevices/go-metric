// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: events/open.proto

package events

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

type Open_OpenTarget int32

const (
	Open_APP                          Open_OpenTarget = 0
	Open_SAVE_KEY                     Open_OpenTarget = 1
	Open_EMULATE                      Open_OpenTarget = 2
	Open_EDIT                         Open_OpenTarget = 3
	Open_SHARE                        Open_OpenTarget = 4
	Open_EXPERIMENTAL_FM              Open_OpenTarget = 5
	Open_EXPERIMENTAL_SCREENSTREAMING Open_OpenTarget = 6
	Open_SHARE_SHORTLINK              Open_OpenTarget = 7
	Open_SHARE_LONGLINK               Open_OpenTarget = 8
	Open_SHARE_FILE                   Open_OpenTarget = 9
)

// Enum value maps for Open_OpenTarget.
var (
	Open_OpenTarget_name = map[int32]string{
		0: "APP",
		1: "SAVE_KEY",
		2: "EMULATE",
		3: "EDIT",
		4: "SHARE",
		5: "EXPERIMENTAL_FM",
		6: "EXPERIMENTAL_SCREENSTREAMING",
		7: "SHARE_SHORTLINK",
		8: "SHARE_LONGLINK",
		9: "SHARE_FILE",
	}
	Open_OpenTarget_value = map[string]int32{
		"APP":                          0,
		"SAVE_KEY":                     1,
		"EMULATE":                      2,
		"EDIT":                         3,
		"SHARE":                        4,
		"EXPERIMENTAL_FM":              5,
		"EXPERIMENTAL_SCREENSTREAMING": 6,
		"SHARE_SHORTLINK":              7,
		"SHARE_LONGLINK":               8,
		"SHARE_FILE":                   9,
	}
)

func (x Open_OpenTarget) Enum() *Open_OpenTarget {
	p := new(Open_OpenTarget)
	*p = x
	return p
}

func (x Open_OpenTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Open_OpenTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_events_open_proto_enumTypes[0].Descriptor()
}

func (Open_OpenTarget) Type() protoreflect.EnumType {
	return &file_events_open_proto_enumTypes[0]
}

func (x Open_OpenTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Open_OpenTarget.Descriptor instead.
func (Open_OpenTarget) EnumDescriptor() ([]byte, []int) {
	return file_events_open_proto_rawDescGZIP(), []int{0, 0}
}

type Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target Open_OpenTarget `protobuf:"varint,1,opt,name=target,proto3,enum=metric.events.Open_OpenTarget" json:"target,omitempty"`
}

func (x *Open) Reset() {
	*x = Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_open_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Open) ProtoMessage() {}

func (x *Open) ProtoReflect() protoreflect.Message {
	mi := &file_events_open_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Open.ProtoReflect.Descriptor instead.
func (*Open) Descriptor() ([]byte, []int) {
	return file_events_open_proto_rawDescGZIP(), []int{0}
}

func (x *Open) GetTarget() Open_OpenTarget {
	if x != nil {
		return x.Target
	}
	return Open_APP
}

var File_events_open_proto protoreflect.FileDescriptor

var file_events_open_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x41, 0x56, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4d, 0x55,
	0x4c, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x44, 0x49, 0x54, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x5f, 0x46, 0x4d, 0x10, 0x05,
	0x12, 0x20, 0x0a, 0x1c, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x4c,
	0x5f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47,
	0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x53, 0x48, 0x4f, 0x52,
	0x54, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x48, 0x41, 0x52, 0x45,
	0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x09, 0x42, 0x24, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_open_proto_rawDescOnce sync.Once
	file_events_open_proto_rawDescData = file_events_open_proto_rawDesc
)

func file_events_open_proto_rawDescGZIP() []byte {
	file_events_open_proto_rawDescOnce.Do(func() {
		file_events_open_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_open_proto_rawDescData)
	})
	return file_events_open_proto_rawDescData
}

var file_events_open_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_events_open_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_open_proto_goTypes = []interface{}{
	(Open_OpenTarget)(0), // 0: metric.events.Open.OpenTarget
	(*Open)(nil),         // 1: metric.events.Open
}
var file_events_open_proto_depIdxs = []int32{
	0, // 0: metric.events.Open.target:type_name -> metric.events.Open.OpenTarget
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_events_open_proto_init() }
func file_events_open_proto_init() {
	if File_events_open_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_open_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Open); i {
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
			RawDescriptor: file_events_open_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_open_proto_goTypes,
		DependencyIndexes: file_events_open_proto_depIdxs,
		EnumInfos:         file_events_open_proto_enumTypes,
		MessageInfos:      file_events_open_proto_msgTypes,
	}.Build()
	File_events_open_proto = out.File
	file_events_open_proto_rawDesc = nil
	file_events_open_proto_goTypes = nil
	file_events_open_proto_depIdxs = nil
}
