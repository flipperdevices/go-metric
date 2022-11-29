// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: events/flipper_rpc_info.proto

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

type FlipperRpcInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SdcardIsAvailable bool  `protobuf:"varint,1,opt,name=sdcard_is_available,json=sdcardIsAvailable,proto3" json:"sdcard_is_available,omitempty"`
	InternalFreeByte  int64 `protobuf:"varint,2,opt,name=internal_free_byte,json=internalFreeByte,proto3" json:"internal_free_byte,omitempty"`
	InternalTotalByte int64 `protobuf:"varint,3,opt,name=internal_total_byte,json=internalTotalByte,proto3" json:"internal_total_byte,omitempty"`
	ExternalFreeByte  int64 `protobuf:"varint,4,opt,name=external_free_byte,json=externalFreeByte,proto3" json:"external_free_byte,omitempty"`
	ExternalTotalByte int64 `protobuf:"varint,5,opt,name=external_total_byte,json=externalTotalByte,proto3" json:"external_total_byte,omitempty"`
}

func (x *FlipperRpcInfo) Reset() {
	*x = FlipperRpcInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_flipper_rpc_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlipperRpcInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlipperRpcInfo) ProtoMessage() {}

func (x *FlipperRpcInfo) ProtoReflect() protoreflect.Message {
	mi := &file_events_flipper_rpc_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlipperRpcInfo.ProtoReflect.Descriptor instead.
func (*FlipperRpcInfo) Descriptor() ([]byte, []int) {
	return file_events_flipper_rpc_info_proto_rawDescGZIP(), []int{0}
}

func (x *FlipperRpcInfo) GetSdcardIsAvailable() bool {
	if x != nil {
		return x.SdcardIsAvailable
	}
	return false
}

func (x *FlipperRpcInfo) GetInternalFreeByte() int64 {
	if x != nil {
		return x.InternalFreeByte
	}
	return 0
}

func (x *FlipperRpcInfo) GetInternalTotalByte() int64 {
	if x != nil {
		return x.InternalTotalByte
	}
	return 0
}

func (x *FlipperRpcInfo) GetExternalFreeByte() int64 {
	if x != nil {
		return x.ExternalFreeByte
	}
	return 0
}

func (x *FlipperRpcInfo) GetExternalTotalByte() int64 {
	if x != nil {
		return x.ExternalTotalByte
	}
	return 0
}

var File_events_flipper_rpc_info_proto protoreflect.FileDescriptor

var file_events_flipper_rpc_info_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xfc,
	0x01, 0x0a, 0x0e, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x70, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x64, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x73, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x73, 0x64, 0x63, 0x61, 0x72, 0x64, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x72,
	0x65, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x72, 0x65, 0x65, 0x42, 0x79, 0x74, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x72, 0x65, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x72, 0x65, 0x65, 0x42, 0x79, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x42, 0x24, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_flipper_rpc_info_proto_rawDescOnce sync.Once
	file_events_flipper_rpc_info_proto_rawDescData = file_events_flipper_rpc_info_proto_rawDesc
)

func file_events_flipper_rpc_info_proto_rawDescGZIP() []byte {
	file_events_flipper_rpc_info_proto_rawDescOnce.Do(func() {
		file_events_flipper_rpc_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_flipper_rpc_info_proto_rawDescData)
	})
	return file_events_flipper_rpc_info_proto_rawDescData
}

var file_events_flipper_rpc_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_flipper_rpc_info_proto_goTypes = []interface{}{
	(*FlipperRpcInfo)(nil), // 0: metric.events.FlipperRpcInfo
}
var file_events_flipper_rpc_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_events_flipper_rpc_info_proto_init() }
func file_events_flipper_rpc_info_proto_init() {
	if File_events_flipper_rpc_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_flipper_rpc_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlipperRpcInfo); i {
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
			RawDescriptor: file_events_flipper_rpc_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_flipper_rpc_info_proto_goTypes,
		DependencyIndexes: file_events_flipper_rpc_info_proto_depIdxs,
		MessageInfos:      file_events_flipper_rpc_info_proto_msgTypes,
	}.Build()
	File_events_flipper_rpc_info_proto = out.File
	file_events_flipper_rpc_info_proto_rawDesc = nil
	file_events_flipper_rpc_info_proto_goTypes = nil
	file_events_flipper_rpc_info_proto_depIdxs = nil
}
