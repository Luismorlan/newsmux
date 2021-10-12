// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: panoptic_config.proto

package protocol

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

type PanopticConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of this config.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Which data collector this config is invoking.
	DataCollectorId PanopticTask_DataCollectorId `protobuf:"varint,2,opt,name=data_collector_id,json=dataCollectorId,proto3,enum=protocol.PanopticTask_DataCollectorId" json:"data_collector_id,omitempty"`
	// parameters that this task should take in.
	TaskParams *TaskParams `protobuf:"bytes,3,opt,name=task_params,json=taskParams,proto3" json:"task_params,omitempty"`
	// TaskSchedule defines the schedule by which this config should run.
	TaskSchedule *TaskSchedule `protobuf:"bytes,4,opt,name=task_schedule,json=taskSchedule,proto3" json:"task_schedule,omitempty"`
}

func (x *PanopticConfig) Reset() {
	*x = PanopticConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanopticConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanopticConfig) ProtoMessage() {}

func (x *PanopticConfig) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanopticConfig.ProtoReflect.Descriptor instead.
func (*PanopticConfig) Descriptor() ([]byte, []int) {
	return file_panoptic_config_proto_rawDescGZIP(), []int{0}
}

func (x *PanopticConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PanopticConfig) GetDataCollectorId() PanopticTask_DataCollectorId {
	if x != nil {
		return x.DataCollectorId
	}
	return PanopticTask_COLLECTOR_UNSPECIFIED
}

func (x *PanopticConfig) GetTaskParams() *TaskParams {
	if x != nil {
		return x.TaskParams
	}
	return nil
}

func (x *PanopticConfig) GetTaskSchedule() *TaskSchedule {
	if x != nil {
		return x.TaskSchedule
	}
	return nil
}

type TaskSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start the task immediately, ignoring the schedule.
	StartImmediatly bool `protobuf:"varint,1,opt,name=start_immediatly,json=startImmediatly,proto3" json:"start_immediatly,omitempty"`
	// There are many schdules that we should be able to support, such as every
	//other duration of time, at certain time of day.
	//
	// Types that are assignable to Schedule:
	//	*TaskSchedule_Routinely
	Schedule isTaskSchedule_Schedule `protobuf_oneof:"schedule"`
}

func (x *TaskSchedule) Reset() {
	*x = TaskSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSchedule) ProtoMessage() {}

func (x *TaskSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSchedule.ProtoReflect.Descriptor instead.
func (*TaskSchedule) Descriptor() ([]byte, []int) {
	return file_panoptic_config_proto_rawDescGZIP(), []int{1}
}

func (x *TaskSchedule) GetStartImmediatly() bool {
	if x != nil {
		return x.StartImmediatly
	}
	return false
}

func (m *TaskSchedule) GetSchedule() isTaskSchedule_Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (x *TaskSchedule) GetRoutinely() *Routinely {
	if x, ok := x.GetSchedule().(*TaskSchedule_Routinely); ok {
		return x.Routinely
	}
	return nil
}

type isTaskSchedule_Schedule interface {
	isTaskSchedule_Schedule()
}

type TaskSchedule_Routinely struct {
	Routinely *Routinely `protobuf:"bytes,2,opt,name=routinely,proto3,oneof"`
}

func (*TaskSchedule_Routinely) isTaskSchedule_Schedule() {}

// Routinely defines a schedule that executes every other duration of time.
type Routinely struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EveryMilliseconds int64 `protobuf:"varint,1,opt,name=every_milliseconds,json=everyMilliseconds,proto3" json:"every_milliseconds,omitempty"`
}

func (x *Routinely) Reset() {
	*x = Routinely{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routinely) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routinely) ProtoMessage() {}

func (x *Routinely) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routinely.ProtoReflect.Descriptor instead.
func (*Routinely) Descriptor() ([]byte, []int) {
	return file_panoptic_config_proto_rawDescGZIP(), []int{2}
}

func (x *Routinely) GetEveryMilliseconds() int64 {
	if x != nil {
		return x.EveryMilliseconds
	}
	return 0
}

// This message is used for the purpose of config push for the scheduler.
type PanopticConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of task configs that's used to configure scheduler.
	Config []*PanopticConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *PanopticConfigs) Reset() {
	*x = PanopticConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanopticConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanopticConfigs) ProtoMessage() {}

func (x *PanopticConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanopticConfigs.ProtoReflect.Descriptor instead.
func (*PanopticConfigs) Descriptor() ([]byte, []int) {
	return file_panoptic_config_proto_rawDescGZIP(), []int{3}
}

func (x *PanopticConfigs) GetConfig() []*PanopticConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_panoptic_config_proto protoreflect.FileDescriptor

var file_panoptic_config_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x1a, 0x0e, 0x70, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50,
	0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0x7a, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x09, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x79,
	0x42, 0x0a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x09,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x50, 0x61, 0x6e, 0x6f,
	0x70, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x69, 0x73,
	0x6d, 0x6f, 0x72, 0x6c, 0x61, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6d, 0x75, 0x78, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_panoptic_config_proto_rawDescOnce sync.Once
	file_panoptic_config_proto_rawDescData = file_panoptic_config_proto_rawDesc
)

func file_panoptic_config_proto_rawDescGZIP() []byte {
	file_panoptic_config_proto_rawDescOnce.Do(func() {
		file_panoptic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_panoptic_config_proto_rawDescData)
	})
	return file_panoptic_config_proto_rawDescData
}

var file_panoptic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_panoptic_config_proto_goTypes = []interface{}{
	(*PanopticConfig)(nil),            // 0: protocol.PanopticConfig
	(*TaskSchedule)(nil),              // 1: protocol.TaskSchedule
	(*Routinely)(nil),                 // 2: protocol.Routinely
	(*PanopticConfigs)(nil),           // 3: protocol.PanopticConfigs
	(PanopticTask_DataCollectorId)(0), // 4: protocol.PanopticTask.DataCollectorId
	(*TaskParams)(nil),                // 5: protocol.TaskParams
}
var file_panoptic_config_proto_depIdxs = []int32{
	4, // 0: protocol.PanopticConfig.data_collector_id:type_name -> protocol.PanopticTask.DataCollectorId
	5, // 1: protocol.PanopticConfig.task_params:type_name -> protocol.TaskParams
	1, // 2: protocol.PanopticConfig.task_schedule:type_name -> protocol.TaskSchedule
	2, // 3: protocol.TaskSchedule.routinely:type_name -> protocol.Routinely
	0, // 4: protocol.PanopticConfigs.config:type_name -> protocol.PanopticConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_panoptic_config_proto_init() }
func file_panoptic_config_proto_init() {
	if File_panoptic_config_proto != nil {
		return
	}
	file_panoptic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_panoptic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanopticConfig); i {
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
		file_panoptic_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSchedule); i {
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
		file_panoptic_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routinely); i {
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
		file_panoptic_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanopticConfigs); i {
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
	file_panoptic_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TaskSchedule_Routinely)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_panoptic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_panoptic_config_proto_goTypes,
		DependencyIndexes: file_panoptic_config_proto_depIdxs,
		MessageInfos:      file_panoptic_config_proto_msgTypes,
	}.Build()
	File_panoptic_config_proto = out.File
	file_panoptic_config_proto_rawDesc = nil
	file_panoptic_config_proto_goTypes = nil
	file_panoptic_config_proto_depIdxs = nil
}
