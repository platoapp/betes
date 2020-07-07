// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: struct.proto

/*
	Package types is a generated protocol buffer package.

	It is generated from these files:
		struct.proto

	It has these top-level messages:
		Struct
		Value
		ListValue
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}
var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (NullValue) EnumDescriptor() ([]byte, []int) { return fileDescriptorStruct, []int{0} }
func (NullValue) XXX_WellKnownType() string       { return "NullValue" }

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Unordered map of dynamically typed values.
	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Struct) Reset()                    { *m = Struct{} }
func (*Struct) ProtoMessage()               {}
func (*Struct) Descriptor() ([]byte, []int) { return fileDescriptorStruct, []int{0} }
func (*Struct) XXX_WellKnownType() string   { return "Struct" }

func (m *Struct) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptorStruct, []int{1} }
func (*Value) XXX_WellKnownType() string   { return "Value" }

type isValue_Kind interface {
	isValue_Kind()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}
type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,proto3,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}
type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,oneof"`
}
type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,6,opt,name=list_value,json=listValue,oneof"`
}

func (*Value_NullValue) isValue_Kind()   {}
func (*Value_NumberValue) isValue_Kind() {}
func (*Value_StringValue) isValue_Kind() {}
func (*Value_BoolValue) isValue_Kind()   {}
func (*Value_StructValue) isValue_Kind() {}
func (*Value_ListValue) isValue_Kind()   {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NULL_VALUE
}

func (m *Value) GetNumberValue() float64 {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetStructValue() *Struct {
	if x, ok := m.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.NullValue))
	case *Value_NumberValue:
		_ = b.EncodeVarint(2<<3 | proto.WireFixed64)
		_ = b.EncodeFi