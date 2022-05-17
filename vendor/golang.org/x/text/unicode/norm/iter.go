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

// Next returns f(i.input[i.Pos():n]), where n is a boundary of i.input.
// For any input a and b for which f(a) == f(b), subsequent calls
// to Next will return the same segments.
// Modifying runes are grouped together with the preceding starter, if such a starter exists.
// Although not guaranteed, n will typically be the smallest possible n.
func (i *Iter) Next() []byte {
	return i.next(i)
}

func nextASCIIBytes(i *Iter) []byte {
	p := i.p + 1
	if p >= i.rb.nsrc {
		i.setDone()
		return i.rb.src.bytes[i.p:p]
	}
	if i.rb.src.bytes[p] < utf8.RuneSelf {
		p0 := i.p
		i.p = p
		return i.rb.src.bytes[p0:p]
	}
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.next = i.rb.f.nextMain
	return i.next(i)
}

func nextASCIIString(i *Iter) []byte {
	p := i.p + 1
	if p >= i.rb.nsrc {
		i.buf[0] = i.rb.src.str[i.p]
		i.setDone()
		return i.buf[:1]
	}
	if i.rb.src.str[p] < utf8.RuneSelf {
		i.buf[0] = i.rb.src.str[i.p]
		i.p = p
		return i.buf[:1]
	}
	i.info = i.rb.f.info(i.rb.src, i.p)
	i.next = i.rb.f.nextMain
	return i.next(i)
}

func nextHangul(i *Iter) []byte {
	p := i.p
	next := p + hangulUTF8Size
	if next >= i.rb.nsrc {
		i.setDone()
	} else if i.rb.src.hangul(next) == 0 {
		i.rb.ss.next(i.info)
		i.info = i.rb.f.info(i.rb.src, i.p)
		i.next = i.rb.f.nextMain
		return i.next(i)
	}
	i.p = next
	return i.buf[:decomposeHangul(i.buf[:], i.rb.src.hangul(p))]
}

func nextDone(i *Iter) []byte {
	return nil
}

// nextMulti is used for iterating over multi-segment decompositions
// for decomposing normal forms.
func nextMulti(i *Iter) []byte {
	j := 0
	d := i.multiSeg
	// skip first rune
	for j = 1; j < len(d) && !utf8.RuneStart(d[j]); j++ {
	}
	for j < len(d) {
		info := i.rb.f.info(input{bytes: d}, j)
		if info.BoundaryBefore() {
			i.multiSeg = d[j:]
			return d[:j]
		}
		j += int(info.size)
	}
	// treat last segment as normal decomposition
	i.next = i.rb.f.nextMain
	return i.next(i)
}

// nextMultiNorm is used for iterating over multi-segment decompositions
// for composing normal forms.
func nextMultiNorm(i *Iter) []byte {
	j := 0
	d := i.multiSeg
	for j < len(d) {
		info := i.rb.f.info(input{bytes: d}, j)
		if info.BoundaryBefore() {
			i.rb.compose()
			seg := i.buf[:i.rb.flushCopy(i.buf[:])]
			i.rb.insertUnsafe(input{bytes: d}, j, info)
			i.multiSeg = d[j+int(info.size):]
			return seg
		}
		i.rb.insertUnsafe(input{bytes: d}, j, info)
		j += int(info.size)
	}
	i.multiSeg = nil
	i.next = nextComposed
	return doNormComposed(i)
}

// nextDecomposed is the implementation of Next for forms NFD and NFKD.
func nextDecomposed(i *Iter) (next []byte) {
	outp := 0
	inCopyStart, outCopyStart := i.p, 0
	for {
		if sz := int(i.info.size); sz <= 1 {
			i.rb.ss = 0
			p := i.p
			i.p++ // ASCII or illegal byte.  Either way, advance by 1.
			if i.p >= i.rb.nsrc {
				i.setDone()
				return i.returnSlice(p, i.p)
			} else if i.rb.src._byte(i.p) < utf8.RuneSelf {
				i.next = i.asciiF
				return i.returnSlice(p, i.p)
			}
			outp++
		} else if d := i.info.Decomposition(); d != nil {
			// Note: If leading CCC != 0, then len(d) == 2 and last is also non-zero.
			// Case 1: there is a leftover to copy.  In this case the decomposition
			// must begin with a modifier and should always be appended.
			// Case 2: no leftover. Simply return d if followed by a ccc == 0 value.
			p := outp + len(d)
			if outp > 0 {
				i.rb.src.copySlice(i.buf[outCopyStart:], inCopyStart, i.p)
				// TODO: this condition should not be possible, but we leave it
				// in for defensive purposes.
				if p > len(i.buf) {
					return i.buf[:outp]
				}
			} else if i.info.m