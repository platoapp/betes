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
// * Path - Used together with [Path Templating](#pathTemplating), where the parameter value is actually part of the operation's URL. This does not include the host or base path of the API. For example, in `/items/{itemId}`, the path parameter is `itemId`.
// * Query - Parameters that are appended to the URL. For example, in `/items?id=###`, the query parameter is `id`.
// * Header - Custom headers that are expected as part of the request.
// * Body - The payload that's appended to the HTTP request. Since there can only be one payload, there can only be *one* body parameter. The name of the body parameter has no effect on the parameter itself and is used for documentation purposes only. Since Form parameters are also in the payload, body and form parameters cannot exist together for the same operation.
// * Form - Used to describe the payload of an HTTP request when either `application/x-www-form-urlencoded` or `multipart/form-data` are used as the content type of the request (in Swagger's definition, the [`consumes`](#operationConsumes) property of an operation). This is the only parameter type that can be used to send files, thus supporting the `file` type. Since form parameters are sent in the payload, they cannot be declared together with a body parameter for the same operation. Form parameters have a different format based on the content-type used (for further details, consult http://www.w3.org/TR/html401/interact/forms.html#h-17.13.4):
//   * `application/x-www-form-urlencoded` - Similar to the format of Query parameters but as a payload. For example, `foo=1&bar=swagger` - both `foo` and `bar` are form parameters. This is normally used for simple parameters that are being transferred.
//   * `multipart/form-data` - each parameter takes a section in the payload with an internal header. For example, for the header `Content-Disposition: form-data; name="submit-name"` the name of the parameter is `submit-name`. This type of form parameters is more commonly used for file transfers.
//
// For more information: http://goo.gl/8us55a#parameterObject
type Parameter struct {
	Refable
	CommonValidations
	SimpleSchema
	VendorExtensible
	ParamProps
}

// JSONLookup look up a value by the json property name
func (p Parameter) JSONLookup(token string) (interface{}, error) {
	if ex, ok := p.Extensions[token]; ok {
		return &ex, nil
	}
	if token == "$ref" {
		return &p.Ref, nil
	}

	r, _, err := jsonpointer.GetForToken(p.CommonValidations, token)
	if err != nil && !strings.HasPrefix(err.Error(), "object has no field") {
		return nil, err
	}
	if r != nil {
		return r, nil
	}
	r, _, err = jsonpointer.GetForToken(p.SimpleSchema, token)
	if err != nil && !strings.HasPrefix(err.Error(), "object has no field") {
		return nil, err
	}
	if r != nil {
		return r, nil
	}
	r, _, err = jsonpointer.GetForToken(p.ParamProps, token)
	return r, err
}

// WithDescription a fluent builder method for the description of the parameter
func (p *Parameter) WithDescription(description string) *Parameter {
	p.Description = description
	return p
}

// Named a fluent builder method to override the name of the parameter
func (p *Parameter) Named(name string) *Parameter {
	p.Name = name
	return p
}

// WithLocation a fluent builder method to override the location of the parameter
func (p *Parameter) WithLocation(in string) *Parameter {
	p.In = in
	return p
}

// Typed a fluent builder method for the type of the parameter value
func (p *Parameter) Typed(tpe, format string) *Parameter {
	p.Type = tpe
	p.Format = format
	return p
}

// CollectionOf a fluent builder method for an array parameter
func (p *Parameter) CollectionOf(items *Items, format string) *Parameter {
	p.Type = "array"
	p.Items = items
	p.CollectionFormat = format
	ret