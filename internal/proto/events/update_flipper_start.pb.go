// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: events/update_flipper_start.proto

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

type UpdateFlipperStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateFrom string `protobuf:"bytes,1,opt,name=update_from,json=updateFrom,proto3" json:"update_from,omitempty"`
	UpdateTo   string `protobuf:"bytes,2,opt,name=update_to,json=updateTo,proto3" json:"update_to,omitempty"`
	UpdateId   int64  `protobuf:"varint,3,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
}

func (x *UpdateFlipperStart) Reset() {
	*x = UpdateFlipperStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_update_flipper_start_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlipperStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlipperStart) ProtoMessage() {}

func (x *UpdateFlipperStart) ProtoReflect() protoreflect.Message {
	mi := &file_events_update_flipper_start_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlipperStart.ProtoReflect.Descriptor instead.
func (*UpdateFlipperStart) Descriptor() ([]byte, []int) {
	return file_events_update_flipper_start_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateFlipperStart) GetUpdateFrom() string {
	if x != nil {
		return x.UpdateFrom
	}
	return ""
}

func (x *UpdateFlipperStart) GetUpdateTo() string {
	if x != nil {
		return x.UpdateTo
	}
	return ""
}

func (x *UpdateFlipperStart) GetUpdateId() int64 {
	if x != nil {
		return x.UpdateId
	}
	return 0
}

var File_events_update_flipper_start_proto protoreflect.FileDescriptor

var file_events_update_flipper_start_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x6f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x42, 0x24, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_events_update_flipper_start_proto_rawDescOnce sync.Once
	file_events_update_flipper_start_proto_rawDescData = file_events_update_flipper_start_proto_rawDesc
)

func file_events_update_flipper_start_proto_rawDescGZIP() []byte {
	file_events_update_flipper_start_proto_rawDescOnce.Do(func() {
		file_events_update_flipper_start_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_update_flipper_start_proto_rawDescData)
	})
	return file_events_update_flipper_start_proto_rawDescData
}

var file_events_update_flipper_start_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_update_flipper_start_proto_goTypes = []interface{}{
	(*UpdateFlipperStart)(nil), // 0: metric.events.UpdateFlipperStart
}
var file_events_update_flipper_start_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_events_update_flipper_start_proto_init() }
func file_events_update_flipper_start_proto_init() {
	if File_events_update_flipper_start_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_update_flipper_start_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlipperStart); i {
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
			RawDescriptor: file_events_update_flipper_start_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_update_flipper_start_proto_goTypes,
		DependencyIndexes: file_events_update_flipper_start_proto_depIdxs,
		MessageInfos:      file_events_update_flipper_start_proto_msgTypes,
	}.Build()
	File_events_update_flipper_start_proto = out.File
	file_events_update_flipper_start_proto_rawDesc = nil
	file_events_update_flipper_start_proto_goTypes = nil
	file_events_update_flipper_start_proto_depIdxs = nil
}
