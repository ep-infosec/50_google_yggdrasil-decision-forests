//
// Copyright 2022 Google LLC.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: yggdrasil_decision_forests/model/random_forest/random_forest.proto

package proto

import (
	proto1 "github.com/google/yggdrasil-decision-forests/yggdrasil_decision_forests/port/go/metric/proto"
	proto "github.com/google/yggdrasil-decision-forests/yggdrasil_decision_forests/port/go/model/proto"
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

// Header for the random forest model.
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of shards used to store the nodes.
	NumNodeShards *int32 `protobuf:"varint,1,opt,name=num_node_shards,json=numNodeShards" json:"num_node_shards,omitempty"`
	// Number of trees.
	NumTrees *int64 `protobuf:"varint,2,opt,name=num_trees,json=numTrees" json:"num_trees,omitempty"`
	// Whether the vote of individual trees are distributions or winner-take-all.
	WinnerTakeAllInference *bool `protobuf:"varint,3,opt,name=winner_take_all_inference,json=winnerTakeAllInference,def=1" json:"winner_take_all_inference,omitempty"`
	// Evaluation of the model, on the out-of-bag examples, during the training.
	OutOfBagEvaluations []*OutOfBagTrainingEvaluations `protobuf:"bytes,4,rep,name=out_of_bag_evaluations,json=outOfBagEvaluations" json:"out_of_bag_evaluations,omitempty"`
	// Variable importance measures.
	MeanDecreaseInAccuracy []*proto.VariableImportance `protobuf:"bytes,5,rep,name=mean_decrease_in_accuracy,json=meanDecreaseInAccuracy" json:"mean_decrease_in_accuracy,omitempty"`
	MeanIncreaseInRmse     []*proto.VariableImportance `protobuf:"bytes,6,rep,name=mean_increase_in_rmse,json=meanIncreaseInRmse" json:"mean_increase_in_rmse,omitempty"`
	// Container used to store the trees' nodes.
	NodeFormat *string `protobuf:"bytes,7,opt,name=node_format,json=nodeFormat,def=TFE_RECORDIO" json:"node_format,omitempty"`
	// Number of nodes trained and then pruned during the training.
	// The classical random forest learning algorithm does not prune nodes.
	NumPrunedNodes *int64 `protobuf:"varint,8,opt,name=num_pruned_nodes,json=numPrunedNodes" json:"num_pruned_nodes,omitempty"`
}

// Default values for Header fields.
const (
	Default_Header_WinnerTakeAllInference = bool(true)
	Default_Header_NodeFormat             = string("TFE_RECORDIO")
)

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetNumNodeShards() int32 {
	if x != nil && x.NumNodeShards != nil {
		return *x.NumNodeShards
	}
	return 0
}

func (x *Header) GetNumTrees() int64 {
	if x != nil && x.NumTrees != nil {
		return *x.NumTrees
	}
	return 0
}

func (x *Header) GetWinnerTakeAllInference() bool {
	if x != nil && x.WinnerTakeAllInference != nil {
		return *x.WinnerTakeAllInference
	}
	return Default_Header_WinnerTakeAllInference
}

func (x *Header) GetOutOfBagEvaluations() []*OutOfBagTrainingEvaluations {
	if x != nil {
		return x.OutOfBagEvaluations
	}
	return nil
}

func (x *Header) GetMeanDecreaseInAccuracy() []*proto.VariableImportance {
	if x != nil {
		return x.MeanDecreaseInAccuracy
	}
	return nil
}

func (x *Header) GetMeanIncreaseInRmse() []*proto.VariableImportance {
	if x != nil {
		return x.MeanIncreaseInRmse
	}
	return nil
}

func (x *Header) GetNodeFormat() string {
	if x != nil && x.NodeFormat != nil {
		return *x.NodeFormat
	}
	return Default_Header_NodeFormat
}

func (x *Header) GetNumPrunedNodes() int64 {
	if x != nil && x.NumPrunedNodes != nil {
		return *x.NumPrunedNodes
	}
	return 0
}

type OutOfBagTrainingEvaluations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of trees available in the model when evaluated.
	NumberOfTrees *int32                    `protobuf:"varint,1,opt,name=number_of_trees,json=numberOfTrees" json:"number_of_trees,omitempty"`
	Evaluation    *proto1.EvaluationResults `protobuf:"bytes,2,opt,name=evaluation" json:"evaluation,omitempty"`
}

func (x *OutOfBagTrainingEvaluations) Reset() {
	*x = OutOfBagTrainingEvaluations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutOfBagTrainingEvaluations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutOfBagTrainingEvaluations) ProtoMessage() {}

func (x *OutOfBagTrainingEvaluations) ProtoReflect() protoreflect.Message {
	mi := &file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutOfBagTrainingEvaluations.ProtoReflect.Descriptor instead.
func (*OutOfBagTrainingEvaluations) Descriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescGZIP(), []int{1}
}

func (x *OutOfBagTrainingEvaluations) GetNumberOfTrees() int32 {
	if x != nil && x.NumberOfTrees != nil {
		return *x.NumberOfTrees
	}
	return 0
}

func (x *OutOfBagTrainingEvaluations) GetEvaluation() *proto1.EvaluationResults {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

var File_yggdrasil_decision_forests_model_random_forest_random_forest_proto protoreflect.FileDescriptor

var file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDesc = []byte{
	0x0a, 0x42, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74,
	0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x66, 0x6f,
	0x72, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd6, 0x04, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x75, 0x6d, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x65, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x65, 0x65,
	0x73, 0x12, 0x3f, 0x0a, 0x19, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x6b, 0x65,
	0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x16, 0x77, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x16, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x61,
	0x67, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x66, 0x6f,
	0x72, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66,
	0x42, 0x61, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x4f, 0x66, 0x42, 0x61, 0x67,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x75, 0x0a, 0x19, 0x6d,
	0x65, 0x61, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x16, 0x6d, 0x65, 0x61, 0x6e,
	0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61,
	0x63, 0x79, 0x12, 0x6d, 0x0a, 0x15, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x6d, 0x73, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x6d,
	0x65, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x52, 0x6d, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x0c, 0x54, 0x46, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x49, 0x4f, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x64, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x50,
	0x72, 0x75, 0x6e, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x4f,
	0x75, 0x74, 0x4f, 0x66, 0x42, 0x61, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x72, 0x65,
	0x65, 0x73, 0x12, 0x5a, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73,
	0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
}

var (
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescOnce sync.Once
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescData = file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDesc
)

func file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescGZIP() []byte {
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescOnce.Do(func() {
		file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescData = protoimpl.X.CompressGZIP(file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescData)
	})
	return file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDescData
}

var file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_goTypes = []interface{}{
	(*Header)(nil),                      // 0: yggdrasil_decision_forests.model.random_forest.proto.Header
	(*OutOfBagTrainingEvaluations)(nil), // 1: yggdrasil_decision_forests.model.random_forest.proto.OutOfBagTrainingEvaluations
	(*proto.VariableImportance)(nil),    // 2: yggdrasil_decision_forests.model.proto.VariableImportance
	(*proto1.EvaluationResults)(nil),    // 3: yggdrasil_decision_forests.metric.proto.EvaluationResults
}
var file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_depIdxs = []int32{
	1, // 0: yggdrasil_decision_forests.model.random_forest.proto.Header.out_of_bag_evaluations:type_name -> yggdrasil_decision_forests.model.random_forest.proto.OutOfBagTrainingEvaluations
	2, // 1: yggdrasil_decision_forests.model.random_forest.proto.Header.mean_decrease_in_accuracy:type_name -> yggdrasil_decision_forests.model.proto.VariableImportance
	2, // 2: yggdrasil_decision_forests.model.random_forest.proto.Header.mean_increase_in_rmse:type_name -> yggdrasil_decision_forests.model.proto.VariableImportance
	3, // 3: yggdrasil_decision_forests.model.random_forest.proto.OutOfBagTrainingEvaluations.evaluation:type_name -> yggdrasil_decision_forests.metric.proto.EvaluationResults
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_init() }
func file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_init() {
	if File_yggdrasil_decision_forests_model_random_forest_random_forest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutOfBagTrainingEvaluations); i {
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
			RawDescriptor: file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_goTypes,
		DependencyIndexes: file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_depIdxs,
		MessageInfos:      file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_msgTypes,
	}.Build()
	File_yggdrasil_decision_forests_model_random_forest_random_forest_proto = out.File
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_rawDesc = nil
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_goTypes = nil
	file_yggdrasil_decision_forests_model_random_forest_random_forest_proto_depIdxs = nil
}
