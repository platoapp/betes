
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
	"fmt"
	"net/url"
	"strings"

	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/swag"
)

// BooleanProperty creates a boolean property
func BooleanProperty() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"boolean"}}}
}

// BoolProperty creates a boolean property
func BoolProperty() *Schema { return BooleanProperty() }

// StringProperty creates a string property
func StringProperty() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"string"}}}
}

// CharProperty creates a string property
func CharProperty() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"string"}}}
}

// Float64Property creates a float64/double property
func Float64Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"number"}, Format: "double"}}
}

// Float32Property creates a float32/float property
func Float32Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"number"}, Format: "float"}}
}

// Int8Property creates an int8 property
func Int8Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"integer"}, Format: "int8"}}
}

// Int16Property creates an int16 property
func Int16Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"integer"}, Format: "int16"}}
}

// Int32Property creates an int32 property
func Int32Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"integer"}, Format: "int32"}}
}

// Int64Property creates an int64 property
func Int64Property() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"integer"}, Format: "int64"}}
}

// StrFmtProperty creates a property for the named string format
func StrFmtProperty(format string) *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"string"}, Format: format}}
}

// DateProperty creates a date property
func DateProperty() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"string"}, Format: "date"}}
}

// DateTimeProperty creates a date time property
func DateTimeProperty() *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"string"}, Format: "date-time"}}
}

// MapProperty creates a map property
func MapProperty(property *Schema) *Schema {
	return &Schema{SchemaProps: SchemaProps{Type: []string{"object"}, AdditionalProperties: &SchemaOrBool{Allows: true, Schema: property}}}
}

// RefProperty creates a ref property
func RefProperty(name string) *Schema {
	return &Schema{SchemaProps: SchemaProps{Ref: MustCreateRef(name)}}
}

// RefSchema creates a ref property
func RefSchema(name string) *Schema {
	return &Schema{SchemaProps: SchemaProps{Ref: MustCreateRef(name)}}
}

// ArrayProperty creates an array property
func ArrayProperty(items *Schema) *Schema {
	if items == nil {
		return &Schema{SchemaProps: SchemaProps{Type: []string{"array"}}}
	}
	return &Schema{SchemaProps: SchemaProps{Items: &SchemaOrArray{Schema: items}, Type: []string{"array"}}}
}

// ComposedSchema creates a schema with allOf
func ComposedSchema(schemas ...Schema) *Schema {
	s := new(Schema)
	s.AllOf = schemas
	return s
}

// SchemaURL represents a schema url
type SchemaURL string

// MarshalJSON marshal this to JSON
func (r SchemaURL) MarshalJSON() ([]byte, error) {
	if r == "" {
		return []byte("{}"), nil
	}
	v := map[string]interface{}{"$schema": string(r)}
	return json.Marshal(v)
}

// UnmarshalJSON unmarshal this from JSON
func (r *SchemaURL) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return r.fromMap(v)
}

func (r *SchemaURL) fromMap(v map[string]interface{}) error {
	if v == nil {
		return nil
	}
	if vv, ok := v["$schema"]; ok {
		if str, ok := vv.(string); ok {
			u, err := url.Parse(str)
			if err != nil {
				return err
			}

			*r = SchemaURL(u.String())
		}
	}
	return nil
}

// type ExtraSchemaProps map[string]interface{}

// // JSONSchema represents a structure that is a json schema draft 04
// type JSONSchema struct {
// 	SchemaProps
// 	ExtraSchemaProps
// }

// // MarshalJSON marshal this to JSON
// func (s JSONSchema) MarshalJSON() ([]byte, error) {
// 	b1, err := json.Marshal(s.SchemaProps)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b2, err := s.Ref.MarshalJSON()
// 	if err != nil {
// 		return nil, err
// 	}
// 	b3, err := s.Schema.MarshalJSON()
// 	if err != nil {
// 		return nil, err
// 	}
// 	b4, err := json.Marshal(s.ExtraSchemaProps)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return swag.ConcatJSON(b1, b2, b3, b4), nil
// }