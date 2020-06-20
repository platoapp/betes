// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// http://github.com/golang/protobuf/
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

package proto

import (
	"reflect"
)

func NewRequiredNotSetError(field string) *RequiredNotSetError {
	return &RequiredNotSetError{field}
}

type Sizer interface {
	Size() int
}

func (o *Buffer) enc_ext_slice_byte(p *Properties, base structPointer) error {
	s := *structPointer_Bytes(base, p.field)
	if s == nil {
		return ErrNil
	}
	o.buf = append(o.buf, s...)
	return nil
}

func size_ext_slice_byte(p *Properties, base structPointer) (n int) {
	s := *structPointer_Bytes(base, p.field)
	if s == nil {
		return 0
	}
	n += len(s)
	return
}

// Encode a reference to bool pointer.
func (o *Buffer) enc_ref_bool(p *Properties, base structPointer) error {
	v := *structPointer_BoolVal(base, p.field)
	x := 0
	if v {
		x = 1
	}
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, uint64(x))
	return nil
}

func size_ref_bool(p *Properties, base structPointer) int {
	return len(p.tagcode) + 1 // each bool takes exactly one byte
}

// Encode a reference to int32 pointer.
func (o *Buffer) enc_ref_int32(p *Properties, base structPointer) error {
	v := structPointer_Word32Val(base, p.field)
	x := int32(word32Val_Get(v))
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, uint64(x))
	return nil
}

func size_ref_int32(p *Properties, base structPointer) (n int) {
	v := structPointer_Word32Val(base, p.field)
	x := int32(word32Val_Get(v))
	n += len(p.tagcode)
	n += p.valSize(uint64(x))
	return
}

func (o *Buffer) enc_ref_uint32(p *Properties, base structPointer) error {
	v := structPointer_Word32Val(base, p.field)
	x := word32Val_Get(v)
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, uint64(x))
	return nil
}

func size_ref_uint32(p *Properties, base structPointer) (n int) {
	v := structPointer_Word32Val(base, p.field)
	x := word32Val_Get(v)
	n += len(p.tagcode)
	n += p.valSize(uint64(x))
	return
}

// Encode a reference to an i