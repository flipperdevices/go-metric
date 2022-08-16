// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: metric.proto

package proto

import (
	events "github.com/flipperdevices/go-metric/internal/proto/events"
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

type MetricReportRequest_Platform int32

const (
	MetricReportRequest_ANDROID       MetricReportRequest_Platform = 0
	MetricReportRequest_ANDROID_DEBUG MetricReportRequest_Platform = 1
	MetricReportRequest_IOS           MetricReportRequest_Platform = 2
	MetricReportRequest_IOS_DEBUG     MetricReportRequest_Platform = 3
)

// Enum value maps for MetricReportRequest_Platform.
var (
	MetricReportRequest_Platform_name = map[int32]string{
		0: "ANDROID",
		1: "ANDROID_DEBUG",
		2: "IOS",
		3: "IOS_DEBUG",
	}
	MetricReportRequest_Platform_value = map[string]int32{
		"ANDROID":       0,
		"ANDROID_DEBUG": 1,
		"IOS":           2,
		"IOS_DEBUG":     3,
	}
)

func (x MetricReportRequest_Platform) Enum() *MetricReportRequest_Platform {
	p := new(MetricReportRequest_Platform)
	*p = x
	return p
}

func (x MetricReportRequest_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricReportRequest_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_metric_proto_enumTypes[0].Descriptor()
}

func (MetricReportRequest_Platform) Type() protoreflect.EnumType {
	return &file_metric_proto_enumTypes[0]
}

func (x MetricReportRequest_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricReportRequest_Platform.Descriptor instead.
func (MetricReportRequest_Platform) EnumDescriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{0, 0}
}

type MetricReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionUuid string                       `protobuf:"bytes,4,opt,name=session_uuid,json=sessionUuid,proto3" json:"session_uuid,omitempty"`
	Version     string                       `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Uuid        string                       `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"` // User uuid
	Platform    MetricReportRequest_Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=metric.MetricReportRequest_Platform" json:"platform,omitempty"`
	Events      []*MetricEventsCollection    `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *MetricReportRequest) Reset() {
	*x = MetricReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricReportRequest) ProtoMessage() {}

func (x *MetricReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricReportRequest.ProtoReflect.Descriptor instead.
func (*MetricReportRequest) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{0}
}

func (x *MetricReportRequest) GetSessionUuid() string {
	if x != nil {
		return x.SessionUuid
	}
	return ""
}

func (x *MetricReportRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MetricReportRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MetricReportRequest) GetPlatform() MetricReportRequest_Platform {
	if x != nil {
		return x.Platform
	}
	return MetricReportRequest_ANDROID
}

func (x *MetricReportRequest) GetEvents() []*MetricEventsCollection {
	if x != nil {
		return x.Events
	}
	return nil
}

type MetricEventsCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*MetricEventsCollection_Open
	//	*MetricEventsCollection_FlipperGattInfo
	//	*MetricEventsCollection_FlipperRpcInfo
	//	*MetricEventsCollection_UpdateFlipperStart
	//	*MetricEventsCollection_UpdateFlipperEnd
	//	*MetricEventsCollection_SynchronizationEnd
	//	*MetricEventsCollection_SubghzProvisioning
	Event isMetricEventsCollection_Event `protobuf_oneof:"event"`
}

func (x *MetricEventsCollection) Reset() {
	*x = MetricEventsCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricEventsCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricEventsCollection) ProtoMessage() {}

func (x *MetricEventsCollection) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricEventsCollection.ProtoReflect.Descriptor instead.
func (*MetricEventsCollection) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{1}
}

func (m *MetricEventsCollection) GetEvent() isMetricEventsCollection_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *MetricEventsCollection) GetOpen() *events.Open {
	if x, ok := x.GetEvent().(*MetricEventsCollection_Open); ok {
		return x.Open
	}
	return nil
}

func (x *MetricEventsCollection) GetFlipperGattInfo() *events.FlipperGattInfo {
	if x, ok := x.GetEvent().(*MetricEventsCollection_FlipperGattInfo); ok {
		return x.FlipperGattInfo
	}
	return nil
}

func (x *MetricEventsCollection) GetFlipperRpcInfo() *events.FlipperRpcInfo {
	if x, ok := x.GetEvent().(*MetricEventsCollection_FlipperRpcInfo); ok {
		return x.FlipperRpcInfo
	}
	return nil
}

func (x *MetricEventsCollection) GetUpdateFlipperStart() *events.UpdateFlipperStart {
	if x, ok := x.GetEvent().(*MetricEventsCollection_UpdateFlipperStart); ok {
		return x.UpdateFlipperStart
	}
	return nil
}

func (x *MetricEventsCollection) GetUpdateFlipperEnd() *events.UpdateFlipperEnd {
	if x, ok := x.GetEvent().(*MetricEventsCollection_UpdateFlipperEnd); ok {
		return x.UpdateFlipperEnd
	}
	return nil
}

func (x *MetricEventsCollection) GetSynchronizationEnd() *events.SynchronizationEnd {
	if x, ok := x.GetEvent().(*MetricEventsCollection_SynchronizationEnd); ok {
		return x.SynchronizationEnd
	}
	return nil
}

func (x *MetricEventsCollection) GetSubghzProvisioning() *events.SubGhzProvisioning {
	if x, ok := x.GetEvent().(*MetricEventsCollection_SubghzProvisioning); ok {
		return x.SubghzProvisioning
	}
	return nil
}

type isMetricEventsCollection_Event interface {
	isMetricEventsCollection_Event()
}

type MetricEventsCollection_Open struct {
	Open *events.Open `protobuf:"bytes,1,opt,name=open,proto3,oneof"`
}

type MetricEventsCollection_FlipperGattInfo struct {
	FlipperGattInfo *events.FlipperGattInfo `protobuf:"bytes,2,opt,name=flipper_gatt_info,json=flipperGattInfo,proto3,oneof"`
}

type MetricEventsCollection_FlipperRpcInfo struct {
	FlipperRpcInfo *events.FlipperRpcInfo `protobuf:"bytes,3,opt,name=flipper_rpc_info,json=flipperRpcInfo,proto3,oneof"`
}

type MetricEventsCollection_UpdateFlipperStart struct {
	UpdateFlipperStart *events.UpdateFlipperStart `protobuf:"bytes,4,opt,name=update_flipper_start,json=updateFlipperStart,proto3,oneof"`
}

type MetricEventsCollection_UpdateFlipperEnd struct {
	UpdateFlipperEnd *events.UpdateFlipperEnd `protobuf:"bytes,5,opt,name=update_flipper_end,json=updateFlipperEnd,proto3,oneof"`
}

type MetricEventsCollection_SynchronizationEnd struct {
	SynchronizationEnd *events.SynchronizationEnd `protobuf:"bytes,12,opt,name=synchronization_end,json=synchronizationEnd,proto3,oneof"`
}

type MetricEventsCollection_SubghzProvisioning struct {
	SubghzProvisioning *events.SubGhzProvisioning `protobuf:"bytes,13,opt,name=subghz_provisioning,json=subghzProvisioning,proto3,oneof"`
}

func (*MetricEventsCollection_Open) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_FlipperGattInfo) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_FlipperRpcInfo) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_UpdateFlipperStart) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_UpdateFlipperEnd) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_SynchronizationEnd) isMetricEventsCollection_Event() {}

func (*MetricEventsCollection_SubghzProvisioning) isMetricEventsCollection_Event() {}

var File_metric_proto protoreflect.FileDescriptor

var file_metric_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x1e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66,
	0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66,
	0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x68, 0x7a, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x36, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4f, 0x53, 0x5f,
	0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x03, 0x22, 0xb9, 0x04, 0x0a, 0x16, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x4c, 0x0a,
	0x11, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x47, 0x61, 0x74, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x6c, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x47, 0x61, 0x74, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x10, 0x66,
	0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x70, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52,
	0x70, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x4f, 0x0a,
	0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x10, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x54,
	0x0a, 0x13, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x48, 0x00,
	0x52, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x64, 0x12, 0x54, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x67, 0x68, 0x7a, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x53, 0x75, 0x62, 0x47, 0x68, 0x7a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x12, 0x73, 0x75, 0x62, 0x67, 0x68, 0x7a, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x1d, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metric_proto_rawDescOnce sync.Once
	file_metric_proto_rawDescData = file_metric_proto_rawDesc
)

func file_metric_proto_rawDescGZIP() []byte {
	file_metric_proto_rawDescOnce.Do(func() {
		file_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_metric_proto_rawDescData)
	})
	return file_metric_proto_rawDescData
}

var file_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_metric_proto_goTypes = []interface{}{
	(MetricReportRequest_Platform)(0), // 0: metric.MetricReportRequest.Platform
	(*MetricReportRequest)(nil),       // 1: metric.MetricReportRequest
	(*MetricEventsCollection)(nil),    // 2: metric.MetricEventsCollection
	(*events.Open)(nil),               // 3: metric.events.Open
	(*events.FlipperGattInfo)(nil),    // 4: metric.events.FlipperGattInfo
	(*events.FlipperRpcInfo)(nil),     // 5: metric.events.FlipperRpcInfo
	(*events.UpdateFlipperStart)(nil), // 6: metric.events.UpdateFlipperStart
	(*events.UpdateFlipperEnd)(nil),   // 7: metric.events.UpdateFlipperEnd
	(*events.SynchronizationEnd)(nil), // 8: metric.events.SynchronizationEnd
	(*events.SubGhzProvisioning)(nil), // 9: metric.events.SubGhzProvisioning
}
var file_metric_proto_depIdxs = []int32{
	0, // 0: metric.MetricReportRequest.platform:type_name -> metric.MetricReportRequest.Platform
	2, // 1: metric.MetricReportRequest.events:type_name -> metric.MetricEventsCollection
	3, // 2: metric.MetricEventsCollection.open:type_name -> metric.events.Open
	4, // 3: metric.MetricEventsCollection.flipper_gatt_info:type_name -> metric.events.FlipperGattInfo
	5, // 4: metric.MetricEventsCollection.flipper_rpc_info:type_name -> metric.events.FlipperRpcInfo
	6, // 5: metric.MetricEventsCollection.update_flipper_start:type_name -> metric.events.UpdateFlipperStart
	7, // 6: metric.MetricEventsCollection.update_flipper_end:type_name -> metric.events.UpdateFlipperEnd
	8, // 7: metric.MetricEventsCollection.synchronization_end:type_name -> metric.events.SynchronizationEnd
	9, // 8: metric.MetricEventsCollection.subghz_provisioning:type_name -> metric.events.SubGhzProvisioning
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_metric_proto_init() }
func file_metric_proto_init() {
	if File_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricReportRequest); i {
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
		file_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricEventsCollection); i {
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
	file_metric_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetricEventsCollection_Open)(nil),
		(*MetricEventsCollection_FlipperGattInfo)(nil),
		(*MetricEventsCollection_FlipperRpcInfo)(nil),
		(*MetricEventsCollection_UpdateFlipperStart)(nil),
		(*MetricEventsCollection_UpdateFlipperEnd)(nil),
		(*MetricEventsCollection_SynchronizationEnd)(nil),
		(*MetricEventsCollection_SubghzProvisioning)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metric_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metric_proto_goTypes,
		DependencyIndexes: file_metric_proto_depIdxs,
		EnumInfos:         file_metric_proto_enumTypes,
		MessageInfos:      file_metric_proto_msgTypes,
	}.Build()
	File_metric_proto = out.File
	file_metric_proto_rawDesc = nil
	file_metric_proto_goTypes = nil
	file_metric_proto_depIdxs = nil
}
