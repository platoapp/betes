// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package norm

import (
	"fmt"
	"unicode/utf8"
)

// MaxSegmentSize is the maximum size of a byte buffer needed to consider any
// sequence of starter and non-starter runes for the purpose of normalization.
const MaxSegmentSize = maxByteBufferSize

// An Iter iterates over a string or byte slice, while normalizing it
// to a given Form.
type Iter struct {
	rb     reorderBuffer
	buf    [maxByteBufferSize]byte
	info   Properties // first character saved from previous iteration
	next   iterFunc   // implementation of next depends on form
	asciiF iterFunc

	p        int    // current position in input source
	multiSeg []byte // remainder of multi-segment decomposition
}

type iterFunc func(*Iter) []byte

// Init initializes i to iterate over src after normalizing it to Form f.
func (i *Iter) Init(f Form, src []byte) {
	i.p = 0
	if len(src) == 0 {
		i.setDone()
		i.rb.nsrc = 0
		return
	}
	i.multiSeg = nil
	i.rb.init(f, src)
	i.next = i.rb.f.nextMain
	i.asciiF = nextASCIIBytes
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.rb.ss.first(i.info)
}

// InitString initializes i to iterate over src after normalizing it to Form f.
func (i *Iter) InitString(f Form, src string) {
	i.p = 0
	if len(src) == 0 {
		i.setDone()
		i.rb.nsrc = 0
		return
	}
	i.multiSeg = nil
	i.rb.initString(f, src)
	i.next = i.rb.f.nextMain
	i.asciiF = nextASCIIString
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.rb.ss.first(i.info)
}

// Seek sets the segment to be returned by the next call to Next to start
// at position p.  It is the responsibility of the caller to set p to the
// start of a segment.
func (i *Iter) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(i.p) + offset
	case 2:
		abs = int64(i.rb.nsrc) + offset
	default:
		return 0, fmt.Errorf("norm: invalid whence")
	}
	if abs < 0 {
		return 0, fmt.Errorf("norm: negative position")
	}
	if int(abs) >= i.rb.nsrc {
		i.setDone()
		return int64(i.p), nil
	}
	i.p = int(abs)
	i.multiSeg = nil
	i.next = i.rb.f.nextMain
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.rb.ss.first(i.info)
	return abs, nil
}

// returnSlice returns a slice of the underlying input type as a byte slice.
// If the underlying is of type []byte, it will simply return a slice.
// If the underlying is of type string, it will copy the slice to the buffer
// and return that.
func (i *Iter) returnSlice(a, b int) []byte {
	if i.rb.src.bytes == nil {
		return i.buf[:copy(i.buf[:], i.rb.src.str[a:b])]
	}
	return i.rb.src.bytes[a:b]
}

// Pos returns the byte position at which the next call to Next will commence processing.
func (i *Iter) Pos() int {
	return i.p
}

func (i *Iter) setDone() {
	i.next = nextDone
	i.p = i.rb.nsrc
}

// Done returns true if there is no more input to process.
func (i *Iter) Done() bool {
	return i.p >= i.rb.nsrc
}

// Next returns f(i.input[i.Pos():n]), where n is a boundary of i.in