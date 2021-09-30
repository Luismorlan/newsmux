// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: panoptic.proto

package protocol

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// DataCollectorId defines the data collection logic to execute.
type PanopticTask_DataCollectorId int32

const (
	PanopticTask_COLLECTOR_UNSPECIFIED PanopticTask_DataCollectorId = 0
	PanopticTask_COLLECTOR_JINSHI      PanopticTask_DataCollectorId = 1
	PanopticTask_COLLECTOR_KUAILANSI   PanopticTask_DataCollectorId = 2 // ...
)

// Enum value maps for PanopticTask_DataCollectorId.
var (
	PanopticTask_DataCollectorId_name = map[int32]string{
		0: "COLLECTOR_UNSPECIFIED",
		1: "COLLECTOR_JINSHI",
		2: "COLLECTOR_KUAILANSI",
	}
	PanopticTask_DataCollectorId_value = map[string]int32{
		"COLLECTOR_UNSPECIFIED": 0,
		"COLLECTOR_JINSHI":      1,
		"COLLECTOR_KUAILANSI":   2,
	}
)

func (x PanopticTask_DataCollectorId) Enum() *PanopticTask_DataCollectorId {
	p := new(PanopticTask_DataCollectorId)
	*p = x
	return p
}

func (x PanopticTask_DataCollectorId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PanopticTask_DataCollectorId) Descriptor() protoreflect.EnumDescriptor {
	return file_panoptic_proto_enumTypes[0].Descriptor()
}

func (PanopticTask_DataCollectorId) Type() protoreflect.EnumType {
	return &file_panoptic_proto_enumTypes[0]
}

func (x PanopticTask_DataCollectorId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PanopticTask_DataCollectorId.Descriptor instead.
func (PanopticTask_DataCollectorId) EnumDescriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{4, 0}
}

// example:
// name: 加州分析员
// type: USERS
// external_id:15281511824122
type PanopticSubSource_SubSourceType int32

const (
	PanopticSubSource_UNSPECIFIED PanopticSubSource_SubSourceType = 0
	PanopticSubSource_FLASHNEWS   PanopticSubSource_SubSourceType = 1
	PanopticSubSource_KEYNEWS     PanopticSubSource_SubSourceType = 2
	PanopticSubSource_USERS       PanopticSubSource_SubSourceType = 3
)

// Enum value maps for PanopticSubSource_SubSourceType.
var (
	PanopticSubSource_SubSourceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "FLASHNEWS",
		2: "KEYNEWS",
		3: "USERS",
	}
	PanopticSubSource_SubSourceType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"FLASHNEWS":   1,
		"KEYNEWS":     2,
		"USERS":       3,
	}
)

func (x PanopticSubSource_SubSourceType) Enum() *PanopticSubSource_SubSourceType {
	p := new(PanopticSubSource_SubSourceType)
	*p = x
	return p
}

func (x PanopticSubSource_SubSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PanopticSubSource_SubSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_panoptic_proto_enumTypes[1].Descriptor()
}

func (PanopticSubSource_SubSourceType) Type() protoreflect.EnumType {
	return &file_panoptic_proto_enumTypes[1]
}

func (x PanopticSubSource_SubSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PanopticSubSource_SubSourceType.Descriptor instead.
func (PanopticSubSource_SubSourceType) EnumDescriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{5, 0}
}

type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PanopticJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Job id uniquely identifies a data collection job, which contains multiple
	// heterogeneous tasks collecting data from multiple sources.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Multiple heterogeneous tasks this job contains.
	Tasks []*PanopticTask `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *PanopticJob) Reset() {
	*x = PanopticJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanopticJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanopticJob) ProtoMessage() {}

func (x *PanopticJob) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanopticJob.ProtoReflect.Descriptor instead.
func (*PanopticJob) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{1}
}

func (x *PanopticJob) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *PanopticJob) GetTasks() []*PanopticTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// TaskParams defines the shared/domain specific parameters that customize the
// execution behavior.
type TaskParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If specified, append these header params to the crawler request,
	// overwrite on same key.
	HeaderParams []*KeyValuePair `protobuf:"bytes,1,rep,name=header_params,json=headerParams,proto3" json:"header_params,omitempty"`
	// If specified, overwrite Cookie with the provided cookies.
	Cookies []*KeyValuePair `protobuf:"bytes,2,rep,name=cookies,proto3" json:"cookies,omitempty"`
	// SourceId of this collected CrawlerMessage.
	SourceId   string               `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SubSources []*PanopticSubSource `protobuf:"bytes,4,rep,name=sub_sources,json=subSources,proto3" json:"sub_sources,omitempty"`
	// offsite of domain specific param from 20, <20 is for shared fields
	// Domain specific params that will be passed in to customize the task
	// execution. For example, you'll pass in Weibo/Twitter/ZSXQ user id as part
	// of the task param.
	//
	// Types that are assignable to Params:
	//	*TaskParams_JinshiTaskParams
	Params isTaskParams_Params `protobuf_oneof:"params"`
}

func (x *TaskParams) Reset() {
	*x = TaskParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskParams) ProtoMessage() {}

func (x *TaskParams) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskParams.ProtoReflect.Descriptor instead.
func (*TaskParams) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{2}
}

func (x *TaskParams) GetHeaderParams() []*KeyValuePair {
	if x != nil {
		return x.HeaderParams
	}
	return nil
}

func (x *TaskParams) GetCookies() []*KeyValuePair {
	if x != nil {
		return x.Cookies
	}
	return nil
}

func (x *TaskParams) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *TaskParams) GetSubSources() []*PanopticSubSource {
	if x != nil {
		return x.SubSources
	}
	return nil
}

func (m *TaskParams) GetParams() isTaskParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *TaskParams) GetJinshiTaskParams() *JinshiTaskParams {
	if x, ok := x.GetParams().(*TaskParams_JinshiTaskParams); ok {
		return x.JinshiTaskParams
	}
	return nil
}

type isTaskParams_Params interface {
	isTaskParams_Params()
}

type TaskParams_JinshiTaskParams struct {
	JinshiTaskParams *JinshiTaskParams `protobuf:"bytes,20,opt,name=jinshi_task_params,json=jinshiTaskParams,proto3,oneof"`
}

func (*TaskParams_JinshiTaskParams) isTaskParams_Params() {}

type TaskMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// job_start/end_time describes the execution span of this task.
	TaskStartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=task_start_time,json=taskStartTime,proto3" json:"task_start_time,omitempty"`
	TaskEndTime   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=task_end_time,json=taskEndTime,proto3" json:"task_end_time,omitempty"`
	// How many CrawlerMessage this task collected.
	TotalMessageCollected int32 `protobuf:"varint,3,opt,name=total_message_collected,json=totalMessageCollected,proto3" json:"total_message_collected,omitempty"`
	// How many CrawlerMessage this task failed to collect.
	TotalMessageFailed int32 `protobuf:"varint,4,opt,name=total_message_failed,json=totalMessageFailed,proto3" json:"total_message_failed,omitempty"`
	// Which ip address is this task executing at.
	IpAddr string `protobuf:"bytes,5,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
}

func (x *TaskMetadata) Reset() {
	*x = TaskMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMetadata) ProtoMessage() {}

func (x *TaskMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskMetadata.ProtoReflect.Descriptor instead.
func (*TaskMetadata) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{3}
}

func (x *TaskMetadata) GetTaskStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.TaskStartTime
	}
	return nil
}

func (x *TaskMetadata) GetTaskEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.TaskEndTime
	}
	return nil
}

func (x *TaskMetadata) GetTotalMessageCollected() int32 {
	if x != nil {
		return x.TotalMessageCollected
	}
	return 0
}

func (x *TaskMetadata) GetTotalMessageFailed() int32 {
	if x != nil {
		return x.TotalMessageFailed
	}
	return 0
}

func (x *TaskMetadata) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

// PanopticTask defines a single data collection task for a single source. A
// task is the smallest execution in Lambda.
type PanopticTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID for this task.
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// The mapping from this ID to corresponding collector is hard coded
	DataCollectorId PanopticTask_DataCollectorId `protobuf:"varint,2,opt,name=data_collector_id,json=dataCollectorId,proto3,enum=protocol.PanopticTask_DataCollectorId" json:"data_collector_id,omitempty"`
	// Params that customizes how to collect data from the web.
	TaskParams *TaskParams `protobuf:"bytes,3,opt,name=task_params,json=taskParams,proto3" json:"task_params,omitempty"`
	// Metadata that's mostly for monitoring.
	TaskMetadata *TaskMetadata `protobuf:"bytes,4,opt,name=task_metadata,json=taskMetadata,proto3" json:"task_metadata,omitempty"`
}

func (x *PanopticTask) Reset() {
	*x = PanopticTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanopticTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanopticTask) ProtoMessage() {}

func (x *PanopticTask) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanopticTask.ProtoReflect.Descriptor instead.
func (*PanopticTask) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{4}
}

func (x *PanopticTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *PanopticTask) GetDataCollectorId() PanopticTask_DataCollectorId {
	if x != nil {
		return x.DataCollectorId
	}
	return PanopticTask_COLLECTOR_UNSPECIFIED
}

func (x *PanopticTask) GetTaskParams() *TaskParams {
	if x != nil {
		return x.TaskParams
	}
	return nil
}

func (x *PanopticTask) GetTaskMetadata() *TaskMetadata {
	if x != nil {
		return x.TaskMetadata
	}
	return nil
}

type PanopticSubSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the sub source, this will be written to CrawlerMessage
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type of the sub source, this will be used to infer:
	// 1. for SUBSOURCE_USERS: DataCollector will hard code what url to crawl
	// 1. for news : DataCollector will determine keynews or flashnews and populate specified one
	Type PanopticSubSource_SubSourceType `protobuf:"varint,2,opt,name=type,proto3,enum=protocol.PanopticSubSource_SubSourceType" json:"type,omitempty"`
	// this will be used to construct external request/crawl uri
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *PanopticSubSource) Reset() {
	*x = PanopticSubSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanopticSubSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanopticSubSource) ProtoMessage() {}

func (x *PanopticSubSource) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanopticSubSource.ProtoReflect.Descriptor instead.
func (*PanopticSubSource) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{5}
}

func (x *PanopticSubSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PanopticSubSource) GetType() PanopticSubSource_SubSourceType {
	if x != nil {
		return x.Type
	}
	return PanopticSubSource_UNSPECIFIED
}

func (x *PanopticSubSource) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

// Created empty param here in case we need to pass in additional parameters to
// customize Jinshi's crawler logic.
type JinshiTaskParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JinshiTaskParams) Reset() {
	*x = JinshiTaskParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panoptic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JinshiTaskParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JinshiTaskParams) ProtoMessage() {}

func (x *JinshiTaskParams) ProtoReflect() protoreflect.Message {
	mi := &file_panoptic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JinshiTaskParams.ProtoReflect.Descriptor instead.
func (*JinshiTaskParams) Descriptor() ([]byte, []int) {
	return file_panoptic_proto_rawDescGZIP(), []int{6}
}

var File_panoptic_proto protoreflect.FileDescriptor

var file_panoptic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x52, 0x0a, 0x0b, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x4a,
	0x6f, 0x62, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x07, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x53, 0x75, 0x62, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x4a, 0x0a, 0x12, 0x6a, 0x69, 0x6e, 0x73, 0x68, 0x69, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4a, 0x69, 0x6e, 0x73, 0x68, 0x69, 0x54, 0x61,
	0x73, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x10, 0x6a, 0x69, 0x6e, 0x73,
	0x68, 0x69, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x74, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x22, 0xcc,
	0x02, 0x0a, 0x0c, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50,
	0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x5b, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x4a, 0x49, 0x4e, 0x53,
	0x48, 0x49, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x4f,
	0x52, 0x5f, 0x4b, 0x55, 0x41, 0x49, 0x4c, 0x41, 0x4e, 0x53, 0x49, 0x10, 0x02, 0x22, 0xd0, 0x01,
	0x0a, 0x11, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x53, 0x75, 0x62, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x53, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4c, 0x41,
	0x53, 0x48, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x45, 0x59, 0x4e,
	0x45, 0x57, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x03,
	0x22, 0x12, 0x0a, 0x10, 0x4a, 0x69, 0x6e, 0x73, 0x68, 0x69, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x32, 0x48, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74,
	0x69, 0x63, 0x4a, 0x6f, 0x62, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x50, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x69,
	0x73, 0x6d, 0x6f, 0x72, 0x6c, 0x61, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6d, 0x75, 0x78, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_panoptic_proto_rawDescOnce sync.Once
	file_panoptic_proto_rawDescData = file_panoptic_proto_rawDesc
)

func file_panoptic_proto_rawDescGZIP() []byte {
	file_panoptic_proto_rawDescOnce.Do(func() {
		file_panoptic_proto_rawDescData = protoimpl.X.CompressGZIP(file_panoptic_proto_rawDescData)
	})
	return file_panoptic_proto_rawDescData
}

var file_panoptic_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_panoptic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_panoptic_proto_goTypes = []interface{}{
	(PanopticTask_DataCollectorId)(0),    // 0: protocol.PanopticTask.DataCollectorId
	(PanopticSubSource_SubSourceType)(0), // 1: protocol.PanopticSubSource.SubSourceType
	(*KeyValuePair)(nil),                 // 2: protocol.KeyValuePair
	(*PanopticJob)(nil),                  // 3: protocol.PanopticJob
	(*TaskParams)(nil),                   // 4: protocol.TaskParams
	(*TaskMetadata)(nil),                 // 5: protocol.TaskMetadata
	(*PanopticTask)(nil),                 // 6: protocol.PanopticTask
	(*PanopticSubSource)(nil),            // 7: protocol.PanopticSubSource
	(*JinshiTaskParams)(nil),             // 8: protocol.JinshiTaskParams
	(*timestamp.Timestamp)(nil),          // 9: google.protobuf.Timestamp
}
var file_panoptic_proto_depIdxs = []int32{
	6,  // 0: protocol.PanopticJob.tasks:type_name -> protocol.PanopticTask
	2,  // 1: protocol.TaskParams.header_params:type_name -> protocol.KeyValuePair
	2,  // 2: protocol.TaskParams.cookies:type_name -> protocol.KeyValuePair
	7,  // 3: protocol.TaskParams.sub_sources:type_name -> protocol.PanopticSubSource
	8,  // 4: protocol.TaskParams.jinshi_task_params:type_name -> protocol.JinshiTaskParams
	9,  // 5: protocol.TaskMetadata.task_start_time:type_name -> google.protobuf.Timestamp
	9,  // 6: protocol.TaskMetadata.task_end_time:type_name -> google.protobuf.Timestamp
	0,  // 7: protocol.PanopticTask.data_collector_id:type_name -> protocol.PanopticTask.DataCollectorId
	4,  // 8: protocol.PanopticTask.task_params:type_name -> protocol.TaskParams
	5,  // 9: protocol.PanopticTask.task_metadata:type_name -> protocol.TaskMetadata
	1,  // 10: protocol.PanopticSubSource.type:type_name -> protocol.PanopticSubSource.SubSourceType
	3,  // 11: protocol.DataCollect.Collect:input_type -> protocol.PanopticJob
	3,  // 12: protocol.DataCollect.Collect:output_type -> protocol.PanopticJob
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_panoptic_proto_init() }
func file_panoptic_proto_init() {
	if File_panoptic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_panoptic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValuePair); i {
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
		file_panoptic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanopticJob); i {
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
		file_panoptic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskParams); i {
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
		file_panoptic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskMetadata); i {
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
		file_panoptic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanopticTask); i {
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
		file_panoptic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanopticSubSource); i {
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
		file_panoptic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JinshiTaskParams); i {
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
	file_panoptic_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TaskParams_JinshiTaskParams)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_panoptic_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_panoptic_proto_goTypes,
		DependencyIndexes: file_panoptic_proto_depIdxs,
		EnumInfos:         file_panoptic_proto_enumTypes,
		MessageInfos:      file_panoptic_proto_msgTypes,
	}.Build()
	File_panoptic_proto = out.File
	file_panoptic_proto_rawDesc = nil
	file_panoptic_proto_goTypes = nil
	file_panoptic_proto_depIdxs = nil
}
