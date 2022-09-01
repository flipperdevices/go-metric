// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: events/subghz_provisioning.proto

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

type SubGhzProvisioning_RegionSource int32

const (
	SubGhzProvisioning_SIM_NETWORK SubGhzProvisioning_RegionSource = 0
	SubGhzProvisioning_SIM_COUNTRY SubGhzProvisioning_RegionSource = 1
	SubGhzProvisioning_GEO_IP      SubGhzProvisioning_RegionSource = 2
	SubGhzProvisioning_SYSTEM      SubGhzProvisioning_RegionSource = 3
	SubGhzProvisioning_DEFAULT     SubGhzProvisioning_RegionSource = 4
)

// Enum value maps for SubGhzProvisioning_RegionSource.
var (
	SubGhzProvisioning_RegionSource_name = map[int32]string{
		0: "SIM_NETWORK",
		1: "SIM_COUNTRY",
		2: "GEO_IP",
		3: "SYSTEM",
		4: "DEFAULT",
	}
	SubGhzProvisioning_RegionSource_value = map[string]int32{
		"SIM_NETWORK": 0,
		"SIM_COUNTRY": 1,
		"GEO_IP":      2,
		"SYSTEM":      3,
		"DEFAULT":     4,
	}
)

func (x SubGhzProvisioning_RegionSource) Enum() *SubGhzProvisioning_RegionSource {
	p := new(SubGhzProvisioning_RegionSource)
	*p = x
	return p
}

func (x SubGhzProvisioning_RegionSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubGhzProvisioning_RegionSource) Descriptor() protoreflect.EnumDescriptor {
	return file_events_subghz_provisioning_proto_enumTypes[0].Descriptor()
}

func (SubGhzProvisioning_RegionSource) Type() protoreflect.EnumType {
	return &file_events_subghz_provisioning_proto_enumTypes[0]
}

func (x SubGhzProvisioning_RegionSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubGhzProvisioning_RegionSource.Descriptor instead.
func (SubGhzProvisioning_RegionSource) EnumDescriptor() ([]byte, []int) {
	return file_events_subghz_provisioning_proto_rawDescGZIP(), []int{0, 0}
}

type SubGhzProvisioning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionNetwork  string                          `protobuf:"bytes,1,opt,name=region_network,json=regionNetwork,proto3" json:"region_network,omitempty"`
	RegionSim_1    string                          `protobuf:"bytes,2,opt,name=region_sim_1,json=regionSim1,proto3" json:"region_sim_1,omitempty"`
	RegionSim_2    string                          `protobuf:"bytes,3,opt,name=region_sim_2,json=regionSim2,proto3" json:"region_sim_2,omitempty"`
	RegionIp       string                          `protobuf:"bytes,4,opt,name=region_ip,json=regionIp,proto3" json:"region_ip,omitempty"`
	RegionSystem   string                          `protobuf:"bytes,5,opt,name=region_system,json=regionSystem,proto3" json:"region_system,omitempty"`
	RegionProvided string                          `protobuf:"bytes,6,opt,name=region_provided,json=regionProvided,proto3" json:"region_provided,omitempty"`
	IsRoaming      bool                            `protobuf:"varint,7,opt,name=is_roaming,json=isRoaming,proto3" json:"is_roaming,omitempty"`
	RegionSource   SubGhzProvisioning_RegionSource `protobuf:"varint,8,opt,name=region_source,json=regionSource,proto3,enum=metric.events.SubGhzProvisioning_RegionSource" json:"region_source,omitempty"`
}

func (x *SubGhzProvisioning) Reset() {
	*x = SubGhzProvisioning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_subghz_provisioning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubGhzProvisioning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubGhzProvisioning) ProtoMessage() {}

func (x *SubGhzProvisioning) ProtoReflect() protoreflect.Message {
	mi := &file_events_subghz_provisioning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubGhzProvisioning.ProtoReflect.Descriptor instead.
func (*SubGhzProvisioning) Descriptor() ([]byte, []int) {
	return file_events_subghz_provisioning_proto_rawDescGZIP(), []int{0}
}

func (x *SubGhzProvisioning) GetRegionNetwork() string {
	if x != nil {
		return x.RegionNetwork
	}
	return ""
}

func (x *SubGhzProvisioning) GetRegionSim_1() string {
	if x != nil {
		return x.RegionSim_1
	}
	return ""
}

func (x *SubGhzProvisioning) GetRegionSim_2() string {
	if x != nil {
		return x.RegionSim_2
	}
	return ""
}

func (x *SubGhzProvisioning) GetRegionIp() string {
	if x != nil {
		return x.RegionIp
	}
	return ""
}

func (x *SubGhzProvisioning) GetRegionSystem() string {
	if x != nil {
		return x.RegionSystem
	}
	return ""
}

func (x *SubGhzProvisioning) GetRegionProvided() string {
	if x != nil {
		return x.RegionProvided
	}
	return ""
}

func (x *SubGhzProvisioning) GetIsRoaming() bool {
	if x != nil {
		return x.IsRoaming
	}
	return false
}

func (x *SubGhzProvisioning) GetRegionSource() SubGhzProvisioning_RegionSource {
	if x != nil {
		return x.RegionSource
	}
	return SubGhzProvisioning_SIM_NETWORK
}

var File_events_subghz_provisioning_proto protoreflect.FileDescriptor

var file_events_subghz_provisioning_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x68, 0x7a, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xb5, 0x03, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x47, 0x68, 0x7a, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x20, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x5f, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d,
	0x31, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x5f,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x6d, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x70,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x6f, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a,
	0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x47, 0x68, 0x7a, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4d, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4d, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x45, 0x4f, 0x5f, 0x49, 0x50, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x04, 0x42, 0x24, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_subghz_provisioning_proto_rawDescOnce sync.Once
	file_events_subghz_provisioning_proto_rawDescData = file_events_subghz_provisioning_proto_rawDesc
)

func file_events_subghz_provisioning_proto_rawDescGZIP() []byte {
	file_events_subghz_provisioning_proto_rawDescOnce.Do(func() {
		file_events_subghz_provisioning_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_subghz_provisioning_proto_rawDescData)
	})
	return file_events_subghz_provisioning_proto_rawDescData
}

var file_events_subghz_provisioning_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_events_subghz_provisioning_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_subghz_provisioning_proto_goTypes = []interface{}{
	(SubGhzProvisioning_RegionSource)(0), // 0: metric.events.SubGhzProvisioning.RegionSource
	(*SubGhzProvisioning)(nil),           // 1: metric.events.SubGhzProvisioning
}
var file_events_subghz_provisioning_proto_depIdxs = []int32{
	0, // 0: metric.events.SubGhzProvisioning.region_source:type_name -> metric.events.SubGhzProvisioning.RegionSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_events_subghz_provisioning_proto_init() }
func file_events_subghz_provisioning_proto_init() {
	if File_events_subghz_provisioning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_subghz_provisioning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubGhzProvisioning); i {
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
			RawDescriptor: file_events_subghz_provisioning_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_subghz_provisioning_proto_goTypes,
		DependencyIndexes: file_events_subghz_provisioning_proto_depIdxs,
		EnumInfos:         file_events_subghz_provisioning_proto_enumTypes,
		MessageInfos:      file_events_subghz_provisioning_proto_msgTypes,
	}.Build()
	File_events_subghz_provisioning_proto = out.File
	file_events_subghz_provisioning_proto_rawDesc = nil
	file_events_subghz_provisioning_proto_goTypes = nil
	file_events_subghz_provisioning_proto_depIdxs = nil
}