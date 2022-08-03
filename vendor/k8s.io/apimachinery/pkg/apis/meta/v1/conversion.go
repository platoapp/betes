/*
Copyright 2014 The Kubernetes Authors.

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

package v1

import (
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func AddConversionFuncs(scheme *runtime.Scheme) error {
	return scheme.AddConversionFuncs(
		Convert_v1_TypeMeta_To_v1_TypeMeta,

		Convert_unversioned_ListMeta_To_unversioned_ListMeta,

		Convert_intstr_IntOrString_To_intstr_IntOrString,

		Convert_unversioned_Time_To_unversioned_Time,
		Convert_unversioned_MicroTime_To_unversioned_MicroTime,

		Convert_Pointer_v1_Duration_To_v1_Duration,
		Convert_v1_Duration_To_Pointer_v1_Duration,

		Convert_Slice_string_To_unversioned_Time,

		Convert_resource_Quantity_To_resource_Quantity,

		Convert_string_To_labels_Selector,
		Convert_labels_Selector_To_string,

		Convert_string_To_fields_Selector,
		Convert_fields_Selector_To_string,

		Convert_Pointer_bool_To_bool,
		Convert_bool_To_Pointer_bool,

		Convert_Pointer_string_To_string,
		Convert_string_To_Pointer_string,

		Convert_Pointer_int64_To_int,
		Convert_int_To_Pointer_int64,

		Convert_Pointer_int32_To_int32,
		Convert_int32_To_Pointer_int32,

		Convert_Pointer_float64_To_float64,
		Convert_float64_To_Pointer_float64,

		Conv