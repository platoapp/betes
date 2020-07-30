
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wrappers.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	wrappers.proto

It has these top-level messages:
	DoubleValue
	FloatValue
	Int64Value
	UInt64Value
	Int32Value
	UInt32Value
	BoolValue
	StringValue
	BytesValue
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

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

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *DoubleValue) Reset()                    { *m = DoubleValue{} }
func (*DoubleValue) ProtoMessage()               {}
func (*DoubleValue) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{0} }
func (*DoubleValue) XXX_WellKnownType() string   { return "DoubleValue" }

func (m *DoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *FloatValue) Reset()                    { *m = FloatValue{} }
func (*FloatValue) ProtoMessage()               {}
func (*FloatValue) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{1} }
func (*FloatValue) XXX_WellKnownType() string   { return "FloatValue" }

func (m *FloatValue) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Int64Value) Reset()                    { *m = Int64Value{} }
func (*Int64Value) ProtoMessage()               {}
func (*Int64Value) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{2} }
func (*Int64Value) XXX_WellKnownType() string   { return "Int64Value" }

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UInt64Value) Reset()                    { *m = UInt64Value{} }
func (*UInt64Value) ProtoMessage()               {}
func (*UInt64Value) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{3} }
func (*UInt64Value) XXX_WellKnownType() string   { return "UInt64Value" }

func (m *UInt64Value) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Int32Value) Reset()                    { *m = Int32Value{} }
func (*Int32Value) ProtoMessage()               {}
func (*Int32Value) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{4} }
func (*Int32Value) XXX_WellKnownType() string   { return "Int32Value" }

func (m *Int32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UInt32Value) Reset()                    { *m = UInt32Value{} }
func (*UInt32Value) ProtoMessage()               {}
func (*UInt32Value) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{5} }
func (*UInt32Value) XXX_WellKnownType() string   { return "UInt32Value" }

func (m *UInt32Value) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{6} }
func (*BoolValue) XXX_WellKnownType() string   { return "BoolValue" }

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{7} }
func (*StringValue) XXX_WellKnownType() string   { return "StringValue" }

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptorWrappers, []int{8} }
func (*BytesValue) XXX_WellKnownType() string   { return "BytesValue" }

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*DoubleValue)(nil), "google.protobuf.DoubleValue")
	proto.RegisterType((*FloatValue)(nil), "google.protobuf.FloatValue")
	proto.RegisterType((*Int64Value)(nil), "google.protobuf.Int64Value")
	proto.RegisterType((*UInt64Value)(nil), "google.protobuf.UInt64Value")
	proto.RegisterType((*Int32Value)(nil), "google.protobuf.Int32Value")
	proto.RegisterType((*UInt32Value)(nil), "google.protobuf.UInt32Value")
	proto.RegisterType((*BoolValue)(nil), "google.protobuf.BoolValue")
	proto.RegisterType((*StringValue)(nil), "google.protobuf.StringValue")
	proto.RegisterType((*BytesValue)(nil), "google.protobuf.BytesValue")
}
func (this *DoubleValue) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*DoubleValue)
	if !ok {
		that2, ok := that.(DoubleValue)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *FloatValue) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*FloatValue)
	if !ok {
		that2, ok := that.(FloatValue)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *Int64Value) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Int64Value)
	if !ok {
		that2, ok := that.(Int64Value)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *UInt64Value) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*UInt64Value)
	if !ok {
		that2, ok := that.(UInt64Value)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *Int32Value) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Int32Value)
	if !ok {
		that2, ok := that.(Int32Value)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *UInt32Value) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*UInt32Value)
	if !ok {
		that2, ok := that.(UInt32Value)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *BoolValue) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*BoolValue)
	if !ok {
		that2, ok := that.(BoolValue)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if !this.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *StringValue) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*StringValue)
	if !ok {
		that2, ok := that.(StringValue)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Value != that1.Value {
		if this.Value < that1.Value {
			return -1
		}
		return 1
	}
	return 0
}
func (this *BytesValue) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*BytesValue)
	if !ok {
		that2, ok := that.(BytesValue)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Value, that1.Value); c != 0 {
		return c
	}
	return 0
}
func (this *DoubleValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoubleValue)
	if !ok {
		that2, ok := that.(DoubleValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *FloatValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FloatValue)
	if !ok {
		that2, ok := that.(FloatValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *Int64Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int64Value)
	if !ok {
		that2, ok := that.(Int64Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *UInt64Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UInt64Value)
	if !ok {
		that2, ok := that.(UInt64Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *Int32Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int32Value)
	if !ok {
		that2, ok := that.(Int32Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *UInt32Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UInt32Value)
	if !ok {
		that2, ok := that.(UInt32Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *BoolValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BoolValue)
	if !ok {
		that2, ok := that.(BoolValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *StringValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringValue)
	if !ok {
		that2, ok := that.(StringValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *BytesValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BytesValue)
	if !ok {
		that2, ok := that.(BytesValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	return true
}
func (this *DoubleValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.DoubleValue{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FloatValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.FloatValue{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Int64Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.Int64Value{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UInt64Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.UInt64Value{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Int32Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.Int32Value{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UInt32Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.UInt32Value{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BoolValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.BoolValue{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StringValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.StringValue{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BytesValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.BytesValue{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringWrappers(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DoubleValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DoubleValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x9
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	return i, nil
}

func (m *FloatValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FloatValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0xd
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Value))))
		i += 4
	}
	return i, nil
}

func (m *Int64Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Int64Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *UInt64Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UInt64Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *Int32Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Int32Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *UInt32Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UInt32Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *BoolValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoolValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value {
		dAtA[i] = 0x8
		i++
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *StringValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *BytesValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BytesValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWrappers(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeVarintWrappers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedDoubleValue(r randyWrappers, easy bool) *DoubleValue {
	this := &DoubleValue{}
	this.Value = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedFloatValue(r randyWrappers, easy bool) *FloatValue {
	this := &FloatValue{}
	this.Value = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedInt64Value(r randyWrappers, easy bool) *Int64Value {
	this := &Int64Value{}
	this.Value = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedUInt64Value(r randyWrappers, easy bool) *UInt64Value {
	this := &UInt64Value{}
	this.Value = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedInt32Value(r randyWrappers, easy bool) *Int32Value {
	this := &Int32Value{}
	this.Value = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedUInt32Value(r randyWrappers, easy bool) *UInt32Value {
	this := &UInt32Value{}
	this.Value = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedBoolValue(r randyWrappers, easy bool) *BoolValue {
	this := &BoolValue{}
	this.Value = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStringValue(r randyWrappers, easy bool) *StringValue {
	this := &StringValue{}
	this.Value = string(randStringWrappers(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedBytesValue(r randyWrappers, easy bool) *BytesValue {
	this := &BytesValue{}
	v1 := r.Intn(100)
	this.Value = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyWrappers interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneWrappers(r randyWrappers) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringWrappers(r randyWrappers) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneWrappers(r)
	}
	return string(tmps)
}
func randUnrecognizedWrappers(r randyWrappers, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldWrappers(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldWrappers(dAtA []byte, r randyWrappers, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateWrappers(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateWrappers(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *DoubleValue) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *FloatValue) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 5
	}
	return n
}

func (m *Int64Value) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovWrappers(uint64(m.Value))
	}
	return n
}

func (m *UInt64Value) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovWrappers(uint64(m.Value))
	}
	return n
}

func (m *Int32Value) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovWrappers(uint64(m.Value))
	}
	return n
}

func (m *UInt32Value) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovWrappers(uint64(m.Value))
	}
	return n
}

func (m *BoolValue) Size() (n int) {
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func (m *StringValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWrappers(uint64(l))
	}
	return n
}

func (m *BytesValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWrappers(uint64(l))
	}
	return n
}

func sovWrappers(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWrappers(x uint64) (n int) {
	return sovWrappers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DoubleValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DoubleValue{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FloatValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FloatValue{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Int64Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Int64Value{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UInt64Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UInt64Value{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Int32Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Int32Value{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UInt32Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UInt32Value{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BoolValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BoolValue{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StringValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StringValue{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BytesValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BytesValue{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringWrappers(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DoubleValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWrappers
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DoubleValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoubleValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipWrappers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWrappers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FloatValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWrappers
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FloatValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FloatValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Value = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipWrappers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWrappers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF