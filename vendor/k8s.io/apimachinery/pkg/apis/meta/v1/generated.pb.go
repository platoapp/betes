/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/generated.proto

	It has these top-level messages:
		APIGroup
		APIGroupList
		APIResource
		APIResourceList
		APIVersions
		DeleteOptions
		Duration
		ExportOptions
		GetOptions
		GroupKind
		GroupResource
		GroupVersion
		GroupVersionForDiscovery
		GroupVersionKind
		GroupVersionResource
		Initializer
		Initializers
		LabelSelector
		LabelSelectorRequirement
		List
		ListMeta
		ListOptions
		MicroTime
		ObjectMeta
		OwnerReference
		Patch
		Preconditions
		RootPaths
		ServerAddressByClientCIDR
		Status
		StatusCause
		StatusDetails
		Time
		Timestamp
		TypeMeta
		Verbs
		WatchEvent
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_runtime "k8s.io/apimachinery/pkg/runtime"

import time "time"
import k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *APIGroup) Reset()                    { *m = APIGroup{} }
func (*APIGroup) ProtoMessage()               {}
func (*APIGroup) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *APIGroupList) Reset()                    { *m = APIGroupList{} }
func (*APIGroupList) ProtoMessage()               {}
func (*APIGroupList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *APIResource) Reset()                    { *m = APIResource{} }
func (*APIResource) ProtoMessage()               {}
func (*APIResource) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *APIResourceList) Reset()                    { *m = APIResourceList{} }
func (*APIResourceList) ProtoMessage()               {}
func (*APIResourceList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *APIVersions) Reset()                    { *m = APIVersions{} }
func (*APIVersions) ProtoMessage()               {}
func (*APIVersions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *DeleteOptions) Reset()                    { *m = DeleteOptions{} }
func (*DeleteOptions) ProtoMessage()               {}
func (*DeleteOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func (m *Duration) Reset()                    { *m = Duration{} }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{6} }

func (m *ExportOptions) Reset()                    { *m = ExportOptions{} }
func (*ExportOptions) ProtoMessage()               {}
func (*ExportOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{7} }

func (m *GetOptions) Reset()                    { *m = GetOptions{} }
func (*GetOptions) ProtoMessage()               {}
func (*GetOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{8} }

func (m *GroupKind) Reset()                    { *m = GroupKind{} }
func (*GroupKind) ProtoMessage()               {}
func (*GroupKind) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{9} }

func (m *GroupResource) Reset()                    { *m = GroupResource{} }
func (*GroupResource) ProtoMessage()               {}
func (*GroupResource) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{10} }

func (m *GroupVersion) Reset()                    { *m = GroupVersion{} }
func (*GroupVersion) ProtoMessage()               {}
func (*GroupVersion) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{11} }

func (m *GroupVersionForDiscovery) Reset()      { *m = GroupVersionForDiscovery{} }
func (*GroupVersionForDiscovery) ProtoMessage() {}
func (*GroupVersionForDiscovery) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{12}
}

func (m *GroupVersionKind) Reset()                    { *m = GroupVersionKind{} }
func (*GroupVersionKind) ProtoMessage()               {}
func (*GroupVersionKind) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{13} }

func (m *GroupVersionResource) Reset()                    { *m = GroupVersionResource{} }
func (*GroupVersionResource) ProtoMessage()               {}
func (*GroupVersionResource) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{14} }

func (m *Initializer) Reset()                    { *m = Initializer{} }
func (*Initializer) ProtoMessage()               {}
func (*Initializer) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{15} }

func (m *Initializers) Reset()                    { *m = Initializers{} }
func (*Initializers) ProtoMessage()               {}
func (*Initializers) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{16} }

func (m *LabelSelector) Reset()                    { *m = LabelSelector{} }
func (*LabelSelector) ProtoMessage()               {}
func (*LabelSelector) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{17} }

func (m *LabelSelectorRequirement) Reset()      { *m = LabelSelectorRequirement{} }
func (*LabelSelectorRequirement) ProtoMessage() {}
func (*LabelSelectorRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{18}
}

func (m *List) Reset()                    { *m = List{} }
func (*List) ProtoMessage()               {}
func (*List) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{19} }

func (m *ListMeta) Reset()                    { *m = ListMeta{} }
func (*ListMeta) ProtoMessage()               {}
func (*ListMeta) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{20} }

func (m *ListOptions) Reset()                    { *m = ListOptions{} }
func (*ListOptions) ProtoMessage()               {}
func (*ListOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{21} }

func (m *MicroTime) Reset()                    { *m = MicroTime{} }
func (*MicroTime) ProtoMessage()               {}
func (*MicroTime) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{22} }

func (m *ObjectMeta) Reset()                    { *m = ObjectMeta{} }
func (*ObjectMeta) ProtoMessage()               {}
func (*ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{23} }

func (m *OwnerReference) Reset()                    { *m = OwnerReference{} }
func (*OwnerReference) ProtoMessage()               {}
func (*OwnerReference) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{24} }

func (m *Patch) Reset()                    { *m = Patch{} }
func (*Patch) ProtoMessage()               {}
func (*Patch) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{25} }

func (m *Preconditions) Reset()                    { *m = Preconditions{} }
func (*Preconditions) ProtoMessage()               {}
func (*Preconditions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{26} }

func (m *RootPaths) Reset()                    { *m = RootPaths{} }
func (*RootPaths) ProtoMessage()               {}
func (*RootPaths) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{27} }

func (m *ServerAddressByClientCIDR) Reset()      { *m = ServerAddressByClientCIDR{} }
func (*ServerAddressByClientCIDR) ProtoMessage() {}
func (*ServerAddressByClientCIDR) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{28}
}

func (m *Status) Reset()                    { *m = Status{} }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{29} }

func (m *StatusCause) Reset()                    { *m = StatusCause{} }
func (*StatusCause) ProtoMessage()               {}
func (*StatusCause) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{30} }

func (m *StatusDetails) Reset()                    { *m = StatusDetails{} }
func (*StatusDetails) ProtoMessage()               {}
func (*StatusDetails) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{31} }

func (m *Time) Reset()                    { *m = Time{} }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{32} }

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{33} }

func (m *TypeMeta) Reset()                    { *m = TypeMeta{} }
func (*TypeMeta) P