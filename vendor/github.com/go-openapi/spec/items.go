
// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"encoding/json"
	"strings"

	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/swag"
)

type SimpleSchema struct {
	Type             string      `json:"type,omitempty"`
	Format           string      `json:"format,omitempty"`
	Items            *Items      `json:"items,omitempty"`
	CollectionFormat string      `json:"collectionFormat,omitempty"`
	Default          interface{} `json:"default,omitempty"`
	Example          interface{} `json:"example,omitempty"`
}

func (s *SimpleSchema) TypeName() string {
	if s.Format != "" {
		return s.Format
	}
	return s.Type
}

func (s *SimpleSchema) ItemsTypeName() string {
	if s.Items == nil {
		return ""
	}
	return s.Items.TypeName()
}

type CommonValidations struct {
	Maximum          *float64      `json:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum,omitempty"`
	MaxLength        *int64        `json:"maxLength,omitempty"`
	MinLength        *int64        `json:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         *int64        `json:"maxItems,omitempty"`
	MinItems         *int64        `json:"minItems,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty"`
	MultipleOf       *float64      `json:"multipleOf,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
}

// Items a limited subset of JSON-Schema's items object.
// It is used by parameter definitions that are not located in "body".
//
// For more information: http://goo.gl/8us55a#items-object
type Items struct {
	Refable
	CommonValidations
	SimpleSchema
	VendorExtensible
}

// NewItems creates a new instance of items
func NewItems() *Items {
	return &Items{}
}

// Typed a fluent builder method for the type of item
func (i *Items) Typed(tpe, format string) *Items {
	i.Type = tpe
	i.Format = format
	return i
}

// CollectionOf a fluent builder method for an array item
func (i *Items) CollectionOf(items *Items, format string) *Items {
	i.Type = "array"
	i.Items = items
	i.CollectionFormat = format
	return i
}

// WithDefault sets the default value on this item
func (i *Items) WithDefault(defaultValue interface{}) *Items {