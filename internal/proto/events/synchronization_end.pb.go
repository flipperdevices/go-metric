// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: events/synchronization_end.proto

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

type SynchronizationEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubghzCount           int32 `protobuf:"varint,1,opt,name=subghz_count,json=subghzCount,proto3" json:"subghz_count,omitempty"`
	RfidCount             int32 `protobuf:"varint,2,opt,name=rfid_count,json=rfidCount,proto3" json:"rfid_count,omitempty"`
	NfcCount              int32 `protobuf:"varint,3,opt,name=nfc_count,json=nfcCount,proto3" json:"nfc_count,omitempty"`
	InfraredCount         int32 `protobuf:"varint,4,opt,name=infrared_count,json=infraredCount,proto3" json:"infrared_count,omitempty"`
	IbuttonCount          int32 `protobuf:"varint,5,opt,name=ibutton_count,json=ibuttonCount,proto3" json:"ibutton_count,omitempty"`
	SynchronizationTimeMs int64 `protobuf:"varint,6,opt,name=synchronization_time_ms,json=synchronizationTimeMs,proto3" json:"synchronization_time_ms,omitempty"`
}

func (x *SynchronizationEnd) Reset() {
	*x = SynchronizationEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_synchronization_end_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynchronizationEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizationEnd) ProtoMessage() {}

func (x *SynchronizationEnd) ProtoReflect() protoreflect.Message {
	mi := &file_events_synchronization_end_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizationEnd.ProtoReflect.Descriptor instead.
func (*SynchronizationEnd) Descriptor() ([]byte, []int) {
	return file_events_synchronization_end_proto_rawDescGZIP(), []int{0}
}

func (x *SynchronizationEnd) GetSubghzCount() int32 {
	if x != nil {
		return x.SubghzCount
	}
	return 0
}

func (x *SynchronizationEnd) GetRfidCount() int32 {
	if x != nil {
		return x.RfidCount
	}
	return 0
}

func (x *SynchronizationEnd) GetNfcCount() int32 {
	if x != nil {
		return x.NfcCount
	}
	return 0
}

func (x *SynchronizationEnd) GetInfraredCount() int32 {
	if x != nil {
		return x.InfraredCount
	}
	return 0
}

func (x *SynchronizationEnd) GetIbuttonCount() int32 {
	if x != nil {
		return x.IbuttonCount
	}
	return 0
}

func (x *SynchronizationEnd) GetSynchronizationTimeMs() int64 {
	if x != nil {
		return x.SynchronizationTimeMs
	}
	return 0
}

var File_events_synchronization_end_proto protoreflect.FileDescriptor

var file_events_synchronization_end_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xf7, 0x01, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x67,
	0x68, 0x7a, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x75, 0x62, 0x67, 0x68, 0x7a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x66, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x66, 0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x66,
	0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e,
	0x66, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x42, 0x24, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_synchronization_end_proto_rawDescOnce sync.Once
	file_events_synchronization_end_proto_rawDescData = file_events_synchronization_end_proto_rawDesc
)

func file_events_synchronization_end_proto_rawDescGZIP() []byte {
	file_events_synchronization_end_proto_rawDescOnce.Do(func() {
		file_events_synchronization_end_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_synchronization_end_proto_rawDescData)
	})
	return file_events_synchronization_end_proto_rawDescData
}

var file_events_synchronization_end_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_synchronization_end_proto_goTypes = []interface{}{
	(*SynchronizationEnd)(nil), // 0: metric.events.SynchronizationEnd
}
var file_events_synchronization_end_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_events_synchronization_end_proto_init() }
func file_events_synchronization_end_proto_init() {
	if File_events_synchronization_end_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_synchronization_end_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynchronizationEnd); i {
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
			RawDescriptor: file_events_synchronization_end_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_synchronization_end_proto_goTypes,
		DependencyIndexes: file_events_synchronization_end_proto_depIdxs,
		MessageInfos:      file_events_synchronization_end_proto_msgTypes,
	}.Build()
	File_events_synchronization_end_proto = out.File
	file_events_synchronization_end_proto_rawDesc = nil
	file_events_synchronization_end_proto_goTypes = nil
	file_events_synchronization_end_proto_depIdxs = nil
}
