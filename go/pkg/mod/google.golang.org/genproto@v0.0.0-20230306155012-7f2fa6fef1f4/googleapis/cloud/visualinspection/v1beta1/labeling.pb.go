// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/visualinspection/v1beta1/labeling.proto

package visualinspection

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Describes the state of a labeling job.
type LabelingJob_JobState int32

const (
	// The job state is unspecified.
	LabelingJob_JOB_STATE_UNSPECIFIED LabelingJob_JobState = 0
	// The job has been just created or resumed and processing has not yet
	// begun.
	LabelingJob_JOB_STATE_QUEUED LabelingJob_JobState = 1
	// The service is preparing to run the job.
	LabelingJob_JOB_STATE_PENDING LabelingJob_JobState = 2
	// The job is in progress.
	LabelingJob_JOB_STATE_RUNNING LabelingJob_JobState = 3
	// The job completed successfully.
	LabelingJob_JOB_STATE_SUCCEEDED LabelingJob_JobState = 4
	// The job failed.
	LabelingJob_JOB_STATE_FAILED LabelingJob_JobState = 5
	// The job is being cancelled. From this state the job may only go to
	// either JOB_STATE_SUCCEEDED, JOB_STATE_FAILED or JOB_STATE_CANCELLED.
	LabelingJob_JOB_STATE_CANCELLING LabelingJob_JobState = 6
	// The job has been cancelled.
	LabelingJob_JOB_STATE_CANCELLED LabelingJob_JobState = 7
	// The job has been stopped, and can be resumed.
	LabelingJob_JOB_STATE_PAUSED LabelingJob_JobState = 8
)

// Enum value maps for LabelingJob_JobState.
var (
	LabelingJob_JobState_name = map[int32]string{
		0: "JOB_STATE_UNSPECIFIED",
		1: "JOB_STATE_QUEUED",
		2: "JOB_STATE_PENDING",
		3: "JOB_STATE_RUNNING",
		4: "JOB_STATE_SUCCEEDED",
		5: "JOB_STATE_FAILED",
		6: "JOB_STATE_CANCELLING",
		7: "JOB_STATE_CANCELLED",
		8: "JOB_STATE_PAUSED",
	}
	LabelingJob_JobState_value = map[string]int32{
		"JOB_STATE_UNSPECIFIED": 0,
		"JOB_STATE_QUEUED":      1,
		"JOB_STATE_PENDING":     2,
		"JOB_STATE_RUNNING":     3,
		"JOB_STATE_SUCCEEDED":   4,
		"JOB_STATE_FAILED":      5,
		"JOB_STATE_CANCELLING":  6,
		"JOB_STATE_CANCELLED":   7,
		"JOB_STATE_PAUSED":      8,
	}
)

func (x LabelingJob_JobState) Enum() *LabelingJob_JobState {
	p := new(LabelingJob_JobState)
	*p = x
	return p
}

func (x LabelingJob_JobState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelingJob_JobState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_enumTypes[0].Descriptor()
}

func (LabelingJob_JobState) Type() protoreflect.EnumType {
	return &file_google_cloud_visualinspection_v1beta1_labeling_proto_enumTypes[0]
}

func (x LabelingJob_JobState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelingJob_JobState.Descriptor instead.
func (LabelingJob_JobState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP(), []int{0, 0}
}

// Labeling job to trigger human labeling for adding annotations to images.
type LabelingJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LabelingConfig:
	//	*LabelingJob_AnomalyDetectionLabelingConfig
	LabelingConfig isLabelingJob_LabelingConfig `protobuf_oneof:"labeling_config"`
	// Output only. Resource name of the LabelingJob.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The user-defined name of the LabelingJob.
	// The name can be up to 128 characters long and can be consist of any UTF-8
	// characters.
	// Display name of a LabelingJob.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The resource name of an existing AnnotationSet for all
	// annotations created through this labeling job. The corresponding
	// AnnotationSet needs to match the `labeling_config`, e.g.,
	// AnomalyDetectionLabelingConfig needs AnnotationSet with
	// `classification_label` specified. Format:
	//
	// projects/{project}/locations/{location}/datasets/{dataset}/annotationSets/{annotation_set}.
	OutputAnnotationSet string `protobuf:"bytes,4,opt,name=output_annotation_set,json=outputAnnotationSet,proto3" json:"output_annotation_set,omitempty"`
	// Required. The Google Cloud Storage location of the instruction PDF. This
	// pdf is shared with labelers, and provides detailed description on how to
	// label.
	InstructionUri string `protobuf:"bytes,5,opt,name=instruction_uri,json=instructionUri,proto3" json:"instruction_uri,omitempty"`
	// Required. The full resource name of annotation specs that will be used for
	// labeling. Format
	//
	// `projects/{project}/locations/{location}/datasets/{dataset}/annotationSpecs/{annotation_spec}`.
	AnnotationSpecs []string `protobuf:"bytes,6,rep,name=annotation_specs,json=annotationSpecs,proto3" json:"annotation_specs,omitempty"`
	// Optional. The SpecialistPools' resource names associated with this job.
	SpecialistPools []string `protobuf:"bytes,7,rep,name=specialist_pools,json=specialistPools,proto3" json:"specialist_pools,omitempty"`
	// Optional. The active learning config, e.g., maximum number of items to
	// label.
	ActiveLearningConfig *ActiveLearningConfig `protobuf:"bytes,8,opt,name=active_learning_config,json=activeLearningConfig,proto3" json:"active_learning_config,omitempty"`
	// Output only. The detailed state of the job.
	State LabelingJob_JobState `protobuf:"varint,9,opt,name=state,proto3,enum=google.cloud.visualinspection.v1beta1.LabelingJob_JobState" json:"state,omitempty"`
	// Output only. Current labeling job progress percentage scaled in interval
	// [0, 100], indicating the percentage of DataItems that has been finished.
	LabelingProgress int32 `protobuf:"varint,10,opt,name=labeling_progress,json=labelingProgress,proto3" json:"labeling_progress,omitempty"`
	// Output only. Timestamp when this LabelingJob was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this LabelingJob was updated most recently.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// LabelingJob.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// Label keys must start with a letter.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LabelingJob) Reset() {
	*x = LabelingJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelingJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelingJob) ProtoMessage() {}

func (x *LabelingJob) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelingJob.ProtoReflect.Descriptor instead.
func (*LabelingJob) Descriptor() ([]byte, []int) {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP(), []int{0}
}

func (m *LabelingJob) GetLabelingConfig() isLabelingJob_LabelingConfig {
	if m != nil {
		return m.LabelingConfig
	}
	return nil
}

func (x *LabelingJob) GetAnomalyDetectionLabelingConfig() *AnomalyDetectionLabelingConfig {
	if x, ok := x.GetLabelingConfig().(*LabelingJob_AnomalyDetectionLabelingConfig); ok {
		return x.AnomalyDetectionLabelingConfig
	}
	return nil
}

func (x *LabelingJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelingJob) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *LabelingJob) GetOutputAnnotationSet() string {
	if x != nil {
		return x.OutputAnnotationSet
	}
	return ""
}

func (x *LabelingJob) GetInstructionUri() string {
	if x != nil {
		return x.InstructionUri
	}
	return ""
}

func (x *LabelingJob) GetAnnotationSpecs() []string {
	if x != nil {
		return x.AnnotationSpecs
	}
	return nil
}

func (x *LabelingJob) GetSpecialistPools() []string {
	if x != nil {
		return x.SpecialistPools
	}
	return nil
}

func (x *LabelingJob) GetActiveLearningConfig() *ActiveLearningConfig {
	if x != nil {
		return x.ActiveLearningConfig
	}
	return nil
}

func (x *LabelingJob) GetState() LabelingJob_JobState {
	if x != nil {
		return x.State
	}
	return LabelingJob_JOB_STATE_UNSPECIFIED
}

func (x *LabelingJob) GetLabelingProgress() int32 {
	if x != nil {
		return x.LabelingProgress
	}
	return 0
}

func (x *LabelingJob) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *LabelingJob) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *LabelingJob) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isLabelingJob_LabelingConfig interface {
	isLabelingJob_LabelingConfig()
}

type LabelingJob_AnomalyDetectionLabelingConfig struct {
	// Optional. Labeling Config for `AnomalyDetection` module.
	AnomalyDetectionLabelingConfig *AnomalyDetectionLabelingConfig `protobuf:"bytes,3,opt,name=anomaly_detection_labeling_config,json=anomalyDetectionLabelingConfig,proto3,oneof"`
}

func (*LabelingJob_AnomalyDetectionLabelingConfig) isLabelingJob_LabelingConfig() {}

// Configuration of active learning.
type ActiveLearningConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of human labeled data items.
	MaxItemCount int64 `protobuf:"varint,1,opt,name=max_item_count,json=maxItemCount,proto3" json:"max_item_count,omitempty"`
}

func (x *ActiveLearningConfig) Reset() {
	*x = ActiveLearningConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveLearningConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveLearningConfig) ProtoMessage() {}

func (x *ActiveLearningConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveLearningConfig.ProtoReflect.Descriptor instead.
func (*ActiveLearningConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP(), []int{1}
}

func (x *ActiveLearningConfig) GetMaxItemCount() int64 {
	if x != nil {
		return x.MaxItemCount
	}
	return 0
}

// Labeling configuration of anomaly detection.
type AnomalyDetectionLabelingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The bounding box annotation set for anomaly detection module.
	// Format:
	//
	// projects/{project}/locations/{location}/datasets/{dataset}/annotationSets/{annotation_set}.
	InputBoundingBoxAnnotationSets []string `protobuf:"bytes,1,rep,name=input_bounding_box_annotation_sets,json=inputBoundingBoxAnnotationSets,proto3" json:"input_bounding_box_annotation_sets,omitempty"`
	// A list of AnnotationSpecs to subselect Annotations in
	// `input_bounding_box_annotation_sets`. If specified, only Annotations which
	// have an AnnotationSpec in the list will be used for labeling. Format:
	//
	// `projects/{project}/locations/{location}/datasets/{dataset}/annotationSpecs/{annotation_spec}`.
	AnnotationSpecAllowlist []string `protobuf:"bytes,2,rep,name=annotation_spec_allowlist,json=annotationSpecAllowlist,proto3" json:"annotation_spec_allowlist,omitempty"`
}

func (x *AnomalyDetectionLabelingConfig) Reset() {
	*x = AnomalyDetectionLabelingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyDetectionLabelingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyDetectionLabelingConfig) ProtoMessage() {}

func (x *AnomalyDetectionLabelingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyDetectionLabelingConfig.ProtoReflect.Descriptor instead.
func (*AnomalyDetectionLabelingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP(), []int{2}
}

func (x *AnomalyDetectionLabelingConfig) GetInputBoundingBoxAnnotationSets() []string {
	if x != nil {
		return x.InputBoundingBoxAnnotationSets
	}
	return nil
}

func (x *AnomalyDetectionLabelingConfig) GetAnnotationSpecAllowlist() []string {
	if x != nil {
		return x.AnnotationSpecAllowlist
	}
	return nil
}

// SpecialistPool represents customers' own workforce to work on their data
// labeling jobs. It includes a group of specialist managers who are responsible
// for managing the labelers in this pool as well as customers' data labeling
// jobs associated with this pool.
// Customers create specialist pool as well as start data labeling jobs on
// Cloud, managers and labelers work with the jobs using CrowdCompute console.
type SpecialistPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name for the SpecialistPool.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. A user friendly display name for the Dataset.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The number of Specialists in this SpecialistPool.
	SpecialistManagersCount int32 `protobuf:"varint,3,opt,name=specialist_managers_count,json=specialistManagersCount,proto3" json:"specialist_managers_count,omitempty"`
	// The email addresses of the specialists in the SpecialistPool.
	SpecialistManagerEmails []string `protobuf:"bytes,4,rep,name=specialist_manager_emails,json=specialistManagerEmails,proto3" json:"specialist_manager_emails,omitempty"`
}

func (x *SpecialistPool) Reset() {
	*x = SpecialistPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecialistPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecialistPool) ProtoMessage() {}

func (x *SpecialistPool) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecialistPool.ProtoReflect.Descriptor instead.
func (*SpecialistPool) Descriptor() ([]byte, []int) {
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP(), []int{3}
}

func (x *SpecialistPool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpecialistPool) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SpecialistPool) GetSpecialistManagersCount() int32 {
	if x != nil {
		return x.SpecialistManagersCount
	}
	return 0
}

func (x *SpecialistPool) GetSpecialistManagerEmails() []string {
	if x != nil {
		return x.SpecialistManagerEmails
	}
	return nil
}

var File_google_cloud_visualinspection_v1beta1_labeling_proto protoreflect.FileDescriptor

var file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x0b, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x97, 0x01, 0x0a, 0x21, 0x61, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x48, 0x00, 0x52, 0x1e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x69, 0x0a, 0x15, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x2c,
	0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x61, 0x0a, 0x10,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e,
	0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12,
	0x2e, 0x0a, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12,
	0x76, 0x0a, 0x16, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x56, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x30, 0x0a, 0x11, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x10, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe1, 0x01,
	0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47,
	0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10,
	0x08, 0x3a, 0x75, 0xea, 0x41, 0x72, 0x0a, 0x2b, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x4a, 0x6f, 0x62, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f, 0x62, 0x7d, 0x42, 0x11, 0x0a, 0x0f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3c, 0x0a, 0x14, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x95, 0x02, 0x0a, 0x1e, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x81, 0x01, 0x0a,
	0x22, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x6f, 0x78, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x2f, 0x0a, 0x2d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x52, 0x1e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42,
	0x6f, 0x78, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x73,
	0x12, 0x6f, 0x0a, 0x19, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0xce, 0x02, 0x0a, 0x0e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x19, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x17, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x3a, 0x7e, 0xea, 0x41, 0x7b, 0x0a, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x49, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x7b, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x7d, 0x42, 0xff, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x50, 0x01, 0x5a, 0x55, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x49,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescOnce sync.Once
	file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescData = file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDesc
)

func file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescGZIP() []byte {
	file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescOnce.Do(func() {
		file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescData)
	})
	return file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDescData
}

var file_google_cloud_visualinspection_v1beta1_labeling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_visualinspection_v1beta1_labeling_proto_goTypes = []interface{}{
	(LabelingJob_JobState)(0),              // 0: google.cloud.visualinspection.v1beta1.LabelingJob.JobState
	(*LabelingJob)(nil),                    // 1: google.cloud.visualinspection.v1beta1.LabelingJob
	(*ActiveLearningConfig)(nil),           // 2: google.cloud.visualinspection.v1beta1.ActiveLearningConfig
	(*AnomalyDetectionLabelingConfig)(nil), // 3: google.cloud.visualinspection.v1beta1.AnomalyDetectionLabelingConfig
	(*SpecialistPool)(nil),                 // 4: google.cloud.visualinspection.v1beta1.SpecialistPool
	nil,                                    // 5: google.cloud.visualinspection.v1beta1.LabelingJob.LabelsEntry
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
}
var file_google_cloud_visualinspection_v1beta1_labeling_proto_depIdxs = []int32{
	3, // 0: google.cloud.visualinspection.v1beta1.LabelingJob.anomaly_detection_labeling_config:type_name -> google.cloud.visualinspection.v1beta1.AnomalyDetectionLabelingConfig
	2, // 1: google.cloud.visualinspection.v1beta1.LabelingJob.active_learning_config:type_name -> google.cloud.visualinspection.v1beta1.ActiveLearningConfig
	0, // 2: google.cloud.visualinspection.v1beta1.LabelingJob.state:type_name -> google.cloud.visualinspection.v1beta1.LabelingJob.JobState
	6, // 3: google.cloud.visualinspection.v1beta1.LabelingJob.create_time:type_name -> google.protobuf.Timestamp
	6, // 4: google.cloud.visualinspection.v1beta1.LabelingJob.update_time:type_name -> google.protobuf.Timestamp
	5, // 5: google.cloud.visualinspection.v1beta1.LabelingJob.labels:type_name -> google.cloud.visualinspection.v1beta1.LabelingJob.LabelsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_visualinspection_v1beta1_labeling_proto_init() }
func file_google_cloud_visualinspection_v1beta1_labeling_proto_init() {
	if File_google_cloud_visualinspection_v1beta1_labeling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelingJob); i {
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
		file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveLearningConfig); i {
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
		file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnomalyDetectionLabelingConfig); i {
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
		file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecialistPool); i {
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
	file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LabelingJob_AnomalyDetectionLabelingConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_visualinspection_v1beta1_labeling_proto_goTypes,
		DependencyIndexes: file_google_cloud_visualinspection_v1beta1_labeling_proto_depIdxs,
		EnumInfos:         file_google_cloud_visualinspection_v1beta1_labeling_proto_enumTypes,
		MessageInfos:      file_google_cloud_visualinspection_v1beta1_labeling_proto_msgTypes,
	}.Build()
	File_google_cloud_visualinspection_v1beta1_labeling_proto = out.File
	file_google_cloud_visualinspection_v1beta1_labeling_proto_rawDesc = nil
	file_google_cloud_visualinspection_v1beta1_labeling_proto_goTypes = nil
	file_google_cloud_visualinspection_v1beta1_labeling_proto_depIdxs = nil
}
