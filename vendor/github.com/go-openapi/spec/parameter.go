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

// QueryParam creates a query parameter
func QueryParam(name string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "query"}}
}

// HeaderParam creates a header parameter, this is always required by default
func HeaderParam(name string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "header", Required: true}}
}

// PathParam creates a path parameter, this is always required
func PathParam(name string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "path", Required: true}}
}

// BodyParam creates a body parameter
func BodyParam(name string, schema *Schema) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "body", Schema: schema}, SimpleSchema: SimpleSchema{Type: "object"}}
}

// FormDataParam creates a body parameter
func FormDataParam(name string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "formData"}}
}

// FileParam creates a body parameter
func FileParam(name string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name, In: "formData"}, SimpleSchema: SimpleSchema{Type: "file"}}
}

// SimpleArrayParam creates a param for a simple array (string, int, date etc)
func SimpleArrayParam(name, tpe, fmt string) *Parameter {
	return &Parameter{ParamProps: ParamProps{Name: name}, SimpleSchema: SimpleSchema{Type: "array", CollectionFormat: "csv", Items: &Items{SimpleSchema: SimpleSchema{Type: "string", Format: fmt}}}}
}

// ParamRef creates a parameter that's a json reference
func ParamRef(uri string) *Parameter {
	p := new(Parameter)
	p.Ref = MustCreateRef(uri)
	return p
}

type ParamProps struct {
	Description     string  `json:"description,omitempty"`
	Name            string  `json:"name,omitempty"`
	In              string  `json:"in,omitempty"`
	Required        bool    `json:"required,omitempty"`
	Schema          *Schema `json:"schema,omitempty"`          // when in == "body"
	AllowEmptyValue bool    `json:"allowEmptyValue,omitempty"` // when in == "query" || "formData"
}

// Parameter a unique parameter is defined by a combination of a [name](#parameterName) and [location](#parameterIn).
//
// There are five possible parameter types.
// * Path - Used together with [Path Templating](#pathTemplating), where the parameter value is actually part of the operation's URL. This does not include the host or base path of the API. For example, in `/items/{itemId}`, the path parameter is