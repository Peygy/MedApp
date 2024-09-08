// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/graph_health/graph_health.proto

package graph_health

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

type GetHealthDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetHealthDataRequest) Reset() {
	*x = GetHealthDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_graph_health_graph_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthDataRequest) ProtoMessage() {}

func (x *GetHealthDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_graph_health_graph_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthDataRequest.ProtoReflect.Descriptor instead.
func (*GetHealthDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_graph_health_graph_health_proto_rawDescGZIP(), []int{0}
}

func (x *GetHealthDataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateHealthDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Age      int32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Height   float32 `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight   float32 `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Pulse    int32   `protobuf:"varint,5,opt,name=pulse,proto3" json:"pulse,omitempty"`
	Pressure string  `protobuf:"bytes,6,opt,name=pressure,proto3" json:"pressure,omitempty"`
}

func (x *UpdateHealthDataRequest) Reset() {
	*x = UpdateHealthDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_graph_health_graph_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHealthDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHealthDataRequest) ProtoMessage() {}

func (x *UpdateHealthDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_graph_health_graph_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHealthDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateHealthDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_graph_health_graph_health_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateHealthDataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateHealthDataRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateHealthDataRequest) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UpdateHealthDataRequest) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UpdateHealthDataRequest) GetPulse() int32 {
	if x != nil {
		return x.Pulse
	}
	return 0
}

func (x *UpdateHealthDataRequest) GetPressure() string {
	if x != nil {
		return x.Pressure
	}
	return ""
}

type HealthDataResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age           int32   `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Height        float32 `protobuf:"fixed32,2,opt,name=height,proto3" json:"height,omitempty"`
	Weight        float32 `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Pulse         int32   `protobuf:"varint,4,opt,name=pulse,proto3" json:"pulse,omitempty"`
	Pressure      string  `protobuf:"bytes,5,opt,name=pressure,proto3" json:"pressure,omitempty"`
	DailyWater    float32 `protobuf:"fixed32,6,opt,name=dailyWater,proto3" json:"dailyWater,omitempty"`
	BodyMassIndex float32 `protobuf:"fixed32,7,opt,name=bodyMassIndex,proto3" json:"bodyMassIndex,omitempty"`
}

func (x *HealthDataResponce) Reset() {
	*x = HealthDataResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_graph_health_graph_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthDataResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthDataResponce) ProtoMessage() {}

func (x *HealthDataResponce) ProtoReflect() protoreflect.Message {
	mi := &file_protos_graph_health_graph_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthDataResponce.ProtoReflect.Descriptor instead.
func (*HealthDataResponce) Descriptor() ([]byte, []int) {
	return file_protos_graph_health_graph_health_proto_rawDescGZIP(), []int{2}
}

func (x *HealthDataResponce) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *HealthDataResponce) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *HealthDataResponce) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *HealthDataResponce) GetPulse() int32 {
	if x != nil {
		return x.Pulse
	}
	return 0
}

func (x *HealthDataResponce) GetPressure() string {
	if x != nil {
		return x.Pressure
	}
	return ""
}

func (x *HealthDataResponce) GetDailyWater() float32 {
	if x != nil {
		return x.DailyWater
	}
	return 0
}

func (x *HealthDataResponce) GetBodyMassIndex() float32 {
	if x != nil {
		return x.BodyMassIndex
	}
	return 0
}

var File_protos_graph_health_graph_health_proto protoreflect.FileDescriptor

var file_protos_graph_health_graph_health_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x75, 0x6c, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x75, 0x6c,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x22, 0xce,
	0x01, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x6c, 0x73, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x57, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x57, 0x61, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x6f, 0x64,
	0x79, 0x4d, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x62, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x32,
	0xc3, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_graph_health_graph_health_proto_rawDescOnce sync.Once
	file_protos_graph_health_graph_health_proto_rawDescData = file_protos_graph_health_graph_health_proto_rawDesc
)

func file_protos_graph_health_graph_health_proto_rawDescGZIP() []byte {
	file_protos_graph_health_graph_health_proto_rawDescOnce.Do(func() {
		file_protos_graph_health_graph_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_graph_health_graph_health_proto_rawDescData)
	})
	return file_protos_graph_health_graph_health_proto_rawDescData
}

var file_protos_graph_health_graph_health_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_graph_health_graph_health_proto_goTypes = []any{
	(*GetHealthDataRequest)(nil),    // 0: graph_health.GetHealthDataRequest
	(*UpdateHealthDataRequest)(nil), // 1: graph_health.UpdateHealthDataRequest
	(*HealthDataResponce)(nil),      // 2: graph_health.HealthDataResponce
}
var file_protos_graph_health_graph_health_proto_depIdxs = []int32{
	0, // 0: graph_health.HealthService.GetHealthData:input_type -> graph_health.GetHealthDataRequest
	1, // 1: graph_health.HealthService.UpdateHealthData:input_type -> graph_health.UpdateHealthDataRequest
	2, // 2: graph_health.HealthService.GetHealthData:output_type -> graph_health.HealthDataResponce
	2, // 3: graph_health.HealthService.UpdateHealthData:output_type -> graph_health.HealthDataResponce
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_graph_health_graph_health_proto_init() }
func file_protos_graph_health_graph_health_proto_init() {
	if File_protos_graph_health_graph_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_graph_health_graph_health_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetHealthDataRequest); i {
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
		file_protos_graph_health_graph_health_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateHealthDataRequest); i {
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
		file_protos_graph_health_graph_health_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HealthDataResponce); i {
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
			RawDescriptor: file_protos_graph_health_graph_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_graph_health_graph_health_proto_goTypes,
		DependencyIndexes: file_protos_graph_health_graph_health_proto_depIdxs,
		MessageInfos:      file_protos_graph_health_graph_health_proto_msgTypes,
	}.Build()
	File_protos_graph_health_graph_health_proto = out.File
	file_protos_graph_health_graph_health_proto_rawDesc = nil
	file_protos_graph_health_graph_health_proto_goTypes = nil
	file_protos_graph_health_graph_health_proto_depIdxs = nil
}
