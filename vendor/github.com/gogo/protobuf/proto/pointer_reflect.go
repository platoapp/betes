// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2012 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// +build appengine js

// This file contains an implementation of proto field accesses using package reflect.
// It is slower than the code in pointer_unsafe.go but it avoids package unsafe and can
// be used on App Engine.

package proto

import (
	"math"
	"reflect"
)

// A structPointer is a pointer to a struct.
type structPointer struct {
	v reflect.Value
}

// toStructPointer returns a structPointer equivalent to the given reflect value.
// The reflect value must itself be a pointer to a struct.
func toStructPointer(v reflect.Value) structPointer {
	return structPointer{v}
}

// IsNil reports whether p is nil.
func structPointer_IsNil(p structPointer) bool {
	return p.v.IsNil()
}

// Interface returns the struct pointer as an interface value.
func structPointer_Interface(p structPointer, _ reflect.Type) interface{} {
	return p.v.Interface()
}

// A field identifies a field in a struct, accessible from a structPointer.
// In this implementation, a field is identified by the sequence of field indices
// passed to reflect's FieldByIndex.
type field []int

// toField returns a field equivalent to the given reflect field.
func toField(f *reflect.StructField) field {
	return f.Index
}

// invalidField is an invalid field identifier.
var invalidField = field(nil)

// IsValid reports whether the field identifier is valid.
func (f field) IsValid() bool { return f != nil }

// field returns the given field in the struct as a reflect value.
func structPointer_field(p structPointer, f field) reflect.Value {
	// Special case: an extension map entry with a value of type T
	// passes a *T to the struct-handling code with a zero field,
	// expecting that it will be treated as equivalent to *struct{ X T },
	// which has the same memory layout. We have to handle that case
	// specially, because reflect will panic if we call FieldByIndex on a
	// non-struct.
	if f == nil {
		return p.v.Elem()
	}

	return p.v.Elem().FieldByIndex(f)
}

// ifield returns the given field in the struct as an interface value.
func structPointer_ifield(p structPointer, f field) interface{} {
	return structPointer_field(p, f).Addr().Interface()
}

// Bytes returns the address of a []byte field in the struct.
func structPointer_Bytes(p structPointer, f field) *[]byte {
	return structPointer_ifield(p, f).(*[]byte)
}

// BytesSlice returns the address of a [][]byte field in the struct.
func structPointer_BytesSlice(p structPointer, f field) *[][]byte {
	return structPointer_ifield(p, f).(*[][]byte)
}

// Bool returns the address of a *bool field in the struct.
func structPointer_Bool(p structPointer, f field) **bool {
	return structPointer_ifield(p, f).(**bool)
}

// BoolVal returns the address of a bool field in the struct.
func structPointer_BoolVal(p structPointer, f field) *bool {
	return structPointer_ifield(p, f).(*bool)
}

// BoolSlice returns the address of a []bool field in the struct.
func structPointer_BoolSlice(p structPointer, f field) *[]bool {
	return structPointer_ifield(p, f).(*[]bool)
}

// String returns the address of a *string field in the struct.
func structPointer_String(p structPointer, f field) **string {
	return structPointer_ifield(p, f).(**string)
}

// StringVal returns the address of a string field in the struct.
func structPointer_StringVal(p structPointer, f field) *string {
	return structPointer_ifield(p, f).(*string)
}

// StringSlice returns the address of a []string field in the struct.
func structPointer_StringSlice(p structPointer, f field) *[]string {
	return structPointer_ifield(p, f).(*[]string)
}

// Extensions returns the address of an extension map field in the struct.
func structPointer_Extensions(p structPointer, f field) *XXX_InternalExtensions {
	return structPointer_ifield(p, f).(*XXX_InternalExtensions)
}

// ExtMap returns the address of an extension map field in the struct.
func structPointer_ExtMap(p structPointer, f field) *map[int32]Extension {
	return structPointer_ifield(p, f).(*map[int32]Extension)
}

// NewAt returns the reflect.Value for a pointer to a field in the struct.
func structPointer_NewAt(p structPointer, f field, typ reflect.Type) reflect.Value {
	return structPointer_field(p, f).Addr()
}

// SetStructPointer writes a *struct field in the struct.
func structPointer_SetStructPointer(p structPointer, f field, q structPointer) {
	structPointer_field(p, f).Set(q.v)
}

// GetStructPointer reads a *struct field in the struct.
func structPointer_GetStructPointer(p structPointer, f field) structPointer {
	return structPointer{structPointer_field(p, f)}
}

// StructPointerSlice the address of a []*struct field in the struct.
func structPointer_StructPointerSlice(p structPointer, f field) structPointerSlice {
	return structPointerSlice{structPointer_field(p, f)}
}

// A structPointerSlice represents the address of a slice of pointers to structs
// (themselves messages or groups). That is, v.Type() is *[]*struct{...}.
type structPointerSlice struct {
	v reflect.Value
}

func (p structPointerSlice) Len() int                  { return p.v.Len() }
func (p structPointerSlice) Index(i int) structPointer { return structPointer{p.v.Index(i)} }
func (p structPointerSlice) Append(q structPointer) {
	p.v.Set(reflect.Append(p.v, q.v))
}

var (
	int32Type   = reflect.TypeOf(int32(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	float32Type = reflect.TypeOf(float32(0))
	int64Type   = reflect.TypeOf(int64(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	float64Type = reflect.TypeOf(float64(0))
)

// A word32 represents a field of type *int32, *uint32, *float32, or *enum.
// That is, v.Type() is *int32, *uint32, *float32, or *enum and v is assignable.
type word32 struct {
	v reflect.Value
}

// IsNil reports whether p is nil.
func word32_IsNil(p word32) bool {
	return p.v.IsNil()
}

// 