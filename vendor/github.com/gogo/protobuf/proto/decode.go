// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
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

package proto

/*
 * Routines for decoding protocol buffer data to construct in-memory representations.
 */

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

// errOverflow is returned when an integer is too large to be represented.
var errOverflow = errors.New("proto: integer overflow")

// ErrInternalBadWireType is returned by generated code when an incorrect
// wire type is encountered. It does not get returned to user code.
var ErrInternalBadWireType = errors.New("proto: internal error: bad wiretype for oneof")

// The fundamental decoders that interpret bytes on the wire.
// Those that take integer types all return uint64 and are
// therefore of type valueDecoder.

// DecodeVarint reads a varint-encoded integer from the slice.
// It returns the integer and the number of bytes consumed, or
// zero if there is not enough.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
func DecodeVarint(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		x |= (b & 0x7F) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}

	// The number is too large to represent in a 64-bit value.
	return 0, 0
}

func (p *Buffer) decodeVarintSlow() (x uint64, err error) {
	i := p.index
	l := len(p.buf)

	for shift := uint(0); shift < 64; shift += 7 {
		if i >= l {
			err = io.ErrUnexpectedEOF
			return
		}
		b := p.buf[i]
		i++
		x |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			p.index = i
			return
		}
	}

	// The number is too large to represent in a 64-bit value.
	err = errOverflow
	return
}

// DecodeVarint reads a varint-encoded integer from the Buffer.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
func (p *Buffer) DecodeVarint() (x uint64, err error) {
	i := p.index
	buf := p.buf

	if i >= len(buf) {
		return 0, io.ErrUnexpectedEOF
	} else if buf[i] < 0x80 {
		p.index++
		return uint64(buf[i]), nil
	} else if len(buf)-i < 10 {
		return p.decodeVarintSlow()
	}

	var b uint64
	// we already checked the first byte
	x = uint64(buf[i]) - 0x80
	i++

	b = uint64(buf[i])
	i++
	x += b << 7
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 7

	b = uint64(buf[i])
	i++
	x += b << 14
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 14

	b = uint64(buf[i])
	i++
	x += b << 21
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 21

	b = uint64(buf[i])
	i++
	x += b << 28
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 28

	b = uint64(buf[i])
	i++
	x += b << 35
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 35

	b = uint64(buf[i])
	i++
	x += b << 42
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 42

	b = uint64(buf[i])
	i++
	x += b << 49
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 49

	b = uint64(buf[i])
	i++
	x += b << 56
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 56

	b = uint64(buf[i])
	i++
	x += b << 63
	if b&0x80 == 0 {
		goto done
	}
	// x -= 0x80 << 63 // Always zero.

	return 0, errOverflow

done:
	p.index = i
	return x, nil
}

// DecodeFixed64 reads a 64-bit integer from the Buffer.
// This is the format for the
// fixed64, sfixed64, and double protocol buffer types.
func (p *Buffer) DecodeFixed64() (x uint64, err error) {
	// x, err already 0
	i := p.index + 8
	if i < 0 || i > len(p.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	p.index = i

	x = uint64(p.buf[i-8])
	x |= uint64(p.buf[i-7]) << 8
	x |= uint64(p.buf[i-6]) << 16
	x |= uint64(p.buf[i-5]) << 24
	x |= uint64(p.buf[i-4]) << 32
	x |= uint64(p.buf[i-3]) << 40
	x |= uint64(p.buf[i-2]) << 48
	x |= uint64(p.buf[i-1]) << 56
	return
}

// DecodeFixed32 reads a 32-bit integer from the Buffer.
// This is the format for the
// fixed32, sfixed32, and float protocol buffer types.
func (p *Buffer) DecodeFixed32() (x uint64, err error) {
	// x, err already 0
	i := p.index + 4
	if i < 0 || i > len(p.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	p.index = i

	x = uint64(p.buf[i-4])
	x |= uint64(p.buf[i-3]) << 8
	x |= uint64(p.buf[i-2]) << 16
	x |= uint64(p.buf[i-1]) << 24
	return
}

// DecodeZigzag64 reads a zigzag-encoded 64-bit integer
// from the Buffer.
// This is the format used for the sint64 protocol buffer type.
func (p *Buffer) DecodeZigzag64() (x uint64, err error) {
	x, err = p.DecodeVarint()
	if err != nil {
		return
	}
	x = (x >> 1) ^ uint64((int64(x&1)<<63)>>63)
	return
}

// DecodeZigzag32 reads a zigzag-encoded 32-bit integer
// from  the Buffer.
// This is the format used for the sint32 protocol buffer type.
func (p *Buffer) DecodeZigzag32() (x uint64, err error) {
	x, err = p.DecodeVarint()
	if err != nil {
		return
	}
	x = uint64((uint32(x) >> 1) ^ uint32((int32(x&1)<<31)>>31))
	return
}

// These are not ValueDecoders: they produce an array of bytes or a string.
// bytes, embedded messages

// DecodeRawBytes reads a count-delimited byte buffer from the Buffer.
// This is the format used for the bytes protocol buffer
// type and for embedded messages.
func (p *Buffer) DecodeRawBytes(alloc bool) (buf []byte, err error) {
	n, err := p.DecodeVarint()
	if err != nil {
		return nil, err
	}

	nb := int(n)
	if nb < 0 {
		return nil, fmt.Errorf("proto: bad byte length %d", nb)
	}
	end := p.index + nb
	if end < p.index || end > len(p.buf) {
		return nil, io.ErrUnexpectedEOF
	}

	if !alloc {
		// todo: check if can get more uses of alloc=false
		buf = p.buf[p.index:end]
		p.index += nb
		return
	}

	buf = make([]byte, nb)
	copy(buf, p.buf[p.index:])
	p.index += nb
	return
}

// DecodeStringBytes reads an encoded string from the Buffer.
// This is the format used for the proto2 string type.
func (p *Buffer) DecodeStringBytes() (s string, err error) {
	buf, err := p.DecodeRawBytes(false)
	if err != nil {
		return
	}
	return string(buf), nil
}

// Skip the next item in the buffer. Its wire type is decoded and presented as an argument.
// If the protocol buffer has extensions, and the field matches, add it as an extension.
// Otherwise, if the XXX_unrecognized field exists, append the skipped data there.
func (o *Buffer) skipAndSave(t reflect.Type, tag, wire int, base structPointer, unrecField field) error {
	oi := o.index

	err := o.skip(t, tag, wire)
	if err != nil {
		return err
	}

	if !unrecField.IsValid() {
		return nil
	}

	ptr := structPointer_Bytes(base, unrecField)

	// Add the skipped field to struct field
	obuf := o.buf

	o.buf = *ptr
	o.EncodeVarint(uint64(tag<<3 | wire))
	*ptr = append(o.buf, obuf[oi:o.index]...)

	o.buf = obuf

	return nil
}

// Skip the next item in the buffer. Its wire type is decoded and presented as an argument.
func (o *Buffer) skip(t reflect.Type, tag, wire int) error {

	var u uint64
	var err error

	switch wire {
	case WireVarint:
		_, err = o.DecodeVarint()
	case WireFixed64:
		_, err = o.DecodeFixed64()
	case WireBytes:
		_, err = o.DecodeRawBytes(false)
	case WireFixed32:
		_, err = o.DecodeFixed32()
	case WireStartGroup:
		for {
			u, err = o.DecodeVarint()
			if err != nil {
				break
			}
			fwire := int(u & 0x7)
			if fwire == WireEndGroup {
				break
			}
			ftag := int(u >> 3)
			err = o.skip(t, ftag, fwire)
			if err != nil {
				break
			}
		}
	default:
		err = fmt.Errorf("proto: can't skip unknown wire type %d for %s", wire, t)
	}
	return err
}

// Unmarshaler is the interface representing objects that can
// unmarshal themselves.  The method should reset the receiver before
// decoding starts.  The argument points to data that may be
// overwritten, so implementations should not keep references to the
// buffer.
type Unmarshaler interface {
	Unmarshal([]byte) error
}

// Unmarshal parses the protocol buffer representation in buf and places the
// decoded result in pb.  If the struct underlying pb does not match
// the data in buf, the results can be unpredictable.
//
// Unmarshal resets pb before starting to unmarshal, so any
// existing data in pb is always removed. Use UnmarshalMerge
// to preserve and append to existing data.
func Unmarshal(buf []byte, pb Message) error {
	pb.Reset()
	return UnmarshalMerge(buf, pb)
}

// UnmarshalMerge parses the protocol buffer representation in buf and
// writes the decoded result to pb.  If the struct underlying pb does not match
// the data in buf, the results can be unpredictable.
//
// UnmarshalMerge merges into existing data in pb.
// Most code should use Unmarshal instead.
func UnmarshalMerge(buf []byte, pb Message) error {
	// If the object can unmarshal itself, let it.
	if u, ok := pb.(Unmarshaler); ok {
		return u.Unmarshal(buf)
	}
	return NewBuffer(buf).Unmarshal(pb)
}

// DecodeMessage reads a count-delimited message from the Buffer.
func (p *Buffer) DecodeMessage(pb Message) error {
	enc, err := p.DecodeRawBytes(false)
	if err != nil {
		return err
	}
	return NewBuffer(enc).Unmarshal(pb)
}

// DecodeGroup reads a tag-delimited group from the Buffer.
func (p *Buffer) DecodeGroup(pb Message) error {
	typ, base, err := getbase(pb)
	if err != nil {
		return err
	}
	return p.unmarshalType(typ.Elem(), GetProperties(typ.Elem()), true, base)
}

// Unmarshal parses the protocol buffer representation in the
// Buffer and places the decoded result in pb.  If the struct
// underlying pb does not match the data in the buffer, the results can be
// unpredictable.
//
// Unlike proto.Unmarshal, this does not reset pb before starting to unmarshal.
func (p *Buffer) Unmarshal(pb Message) error {
	// If the object can unmarshal itself, let it.
	if u, ok := pb.(Unmarshaler); ok {
		err := u.Unmarshal(p.buf[p.index:])
		p.index = len(p.buf)
		return err
	}

	typ, base, err := getbase(pb)
	if err != nil {
		return err
	}

	err = p.unmarshalType(typ.Elem(), GetProperties(typ.Elem()), false, base)

	if collectStats {
		stats.Decode++
	}

	return err
}

// unmarshalType does the work of unmarshaling a structure.
func (o *Buffer) unmarshalType(st reflect.Type, prop *StructProperties, is_group bool, base structPointer) error {
	var state errorState
	required, reqFields := prop.reqCount, uint64(0)

	var err error
	for err == nil && o.index < len(o.buf) {
		oi := o.index
		var u uint64
		u, err = o.DecodeVarint()
		if err != nil {
			break
		}
		wire := int(u & 0x7)
		if wire == WireEndGroup {
			if is_group {
				if required > 0 {
					// Not enough information to determine the exact field.
					// (See below.)
					return &RequiredNotSetError{"{Unknown}"}
				}
				return nil // input is satisfied
			}
			return fmt.Errorf("proto: %s: wiretype end group for non-group", st)
		}
		tag := int(u >> 3)
		if tag <= 0 {
			return fmt.Errorf("proto: %s: illegal tag %d (wire type %d)", st, tag, wire)
		}
		fieldnum, ok := prop.decoderTags.get(tag)
		if !ok {
			// Maybe it's an extension?
			if prop.extendable {
				if e, eok := structPointer_Interface(base, st).(extensionsBytes); eok {
					if isExtensionField(e, int32(tag)) {
						if err = o.skip(st, tag, wire); err == nil {
							ext := e.GetExtensions()
							*ext = append(*ext, o.buf[oi:o.index]...)
						}
						continue
					}
				} else if e, _ := extendable(structPointer_Interface(base, st)); isExtensionField(e, int32(tag)) {
					if err = o.skip(st, tag, wire); err == nil {
						extmap := e.extensionsWri