// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package norm

import "unicode/utf8"

const (
	maxNonStarters = 30
	// The maximum number of characters needed for a buffer is
	// maxNonStarters + 1 for the starter + 1 for the GCJ
	maxBufferSize    = maxNonStarters + 2
	maxNFCExpansion  = 3  // NFC(0x1D160)
	maxNFKCExpansion = 18 // NFKC(0xFDFA)

	maxByteBufferSize = utf8.UTFMax * maxBufferSize // 128
)

// ssState is used for reporting the segment state after inserting a rune.
// It is returned by streamSafe.next.
type ssState int

const (
	// Indicates a rune was successfully added to the segment.
	ssSuccess ssState = iota
	// Indicates a rune starts a new segment and should not be added.
	ssStarter
	// Indicates a rune caused a segment overflow and a CGJ should be inserted.
	ssOverflow
)

// streamSafe implements the policy of when a CGJ should be inserted.
type streamSafe uint8

// first inserts the first rune of a segment. It is a faster version of next if
// it is known p represents the first rune in a segment.
func (ss *streamSafe) first(p Properties) {
	*ss = streamSafe(p.nTrailingNonStarters())
}

// insert returns a ssState value to indicate whether a rune represented by p
// can be inserted.
func (ss *streamSafe) next(p Properties) ssState {
	if *ss > maxNonStarters {
		panic("streamSafe was not reset")
	}
	n := p.nLeadingNonStarters()
	if *ss += streamSafe(n); *ss > maxNonStarters {
		*ss = 0
		return ssOverflow
	}
	// The Stream-Safe Text Processing prescribes that the counting can stop
	// as soon as a starter is encountered. However, there are some starters,
	// like Jamo V and T, that can combine with other runes, leaving their
	// successive non-starters appended to the previous, possibly causing an
	// overflow. We will therefore consider any rune with a non-zero nLead to
	// be a non-starter. Note that it always hold that if nLead > 0 then
	// nLead == nTrail.
	if n == 0 {
		*ss = streamSafe(p.nTrailingNonStarters())
		return ssStarter
	}
	return ssSuccess
}

// backwards is used for checking for overflow and segment starts
// when traversing a string backwards. Users do not need to call first
// for the first rune. The state of the streamSafe retains the count of
// the non-starters loaded.
func (ss *streamSafe) backwards(p Properties) ssState {
	if *ss > maxNonStarters {
		panic("streamSafe was not reset")
	}
	c := *ss + streamSafe(p.nTrailingNonStarters())
	if c > maxNonStarters {
		return ssOverflow
	}
	*ss = c
	if p.nLeadingNonStarters() == 0 {
		return ssStarter
	}
	return ssSuccess
}

func (ss streamSafe) isMax() bool {
	return ss == maxNonStarters
}

// GraphemeJoiner is inserted after maxNonStarters non-starter runes.
const GraphemeJoiner = "\u034F"

// reorderBuffer is used to normalize a single segment.  Characters inserted with
// insert are decomposed and reordered based on CCC. The compose method can
// be used to recombine characters.  Note that the byte buffer does not hold
// the UTF-8 characters in order.  Only the rune array is maintained in sorted
// order. flush writes the resulting segment to a byte array.
type reorderBuffer struct {
	rune  [maxBufferSize]Properties // Per character info.
	byte  [maxByteBufferSize]byte   // UTF-8 buffer. Referenced by runeInfo.pos.
	nbyte uint8                     // Number or bytes.
	ss    streamSafe                // For limiting length of non-starter sequence.
	nrune int                       // Number of runeInfos.
	f     formInfo

	src      input
	nsrc     int
	tmpBytes input

	out    []byte
	flushF func(*reorderBuffer) bool
}

func (rb *reorderBuffer) init(f Form, src []byte) {
	rb.f = *formTable[f]
	rb.src.setBytes(src)
	rb.nsrc = len(src)
	rb.ss = 0
}

func (rb *reorderBuffer) initString(f Form, src string) {
	rb.f = *formTable[f]
	rb.src.setString(src)
	rb.nsrc = len(src)
	rb.ss = 0
}

func (rb *reorderBuffer) setFlusher(out []byte, f func(*reorderBuffer) bool) {
	rb.out = out
	rb.flushF = f
}

// reset discards all characters from the buffer.
func (rb *reorderBuffer) reset() {
	rb.nrune = 0
	rb.nbyte = 0
}

func (rb *reorderBuffer) doFlush() bool {
	if rb.f.composing {
		rb.compose()
	}
	res := rb.flushF(rb)
	rb.reset()
	return res
}

// appendFlush appends the normalized segment to rb.out.
func appendFlush(rb *reorderBuffer) bool {
	for i := 0; i < rb.nrune; i++ {
		start := rb.rune[i].pos
		end := start + rb.rune[i].size
		rb.out = append(rb.out, rb.byte[start:end]...)
	}
	return true
}

// flush appends the normalized segment to out and resets rb.
func (rb *reorderBuffer) flush(out []byte) []byte {
	for i := 0; i < rb.nrune; i++ {
		start := rb.rune[i].pos
		end := start + rb.rune[i].size
		out = append(out, rb.byte[start:end]...)
	}
	rb.reset()
	return out
}

// flushCopy copies the normalized segment to buf and resets rb.
// It returns the number of bytes written to buf.
func (rb *reorderBuffer) flushCopy(buf []byte) int {
	p := 0
	for i := 0; i < rb.nrune; i++ {
		runep := rb.rune[i]
		p += copy(buf[p:], rb.byte[runep.pos:runep.pos+runep.size])
	}
	rb.reset()
	return p
}

// insertOrdered inserts a rune in the buffer, ordered by Canonical Combining Class.
// It returns false if the buffer is not large enough to hold the rune.
// It is used internally by insert and insertString only.
func (rb *reorderBuffer) insertOrdered(info Properties) {
	n := rb.nrune
	b := rb.rune[:]
	cc := info.ccc
	if cc > 0 {
		// Find insertion position + move elements to make room.
		for ; n > 0; n-- {
			if b[n-1].ccc <= cc {
				break
			}
			b[n] = b[n-1]
		}
	}
	rb.nrune += 1
	pos := uint8(rb.nbyte)
	rb.nbyte += utf8.UTFMax
	info.pos = pos
	b[n] = info
}

// insertErr is an error code returned by insert. Using this type instead
// of error improves performance up to 20% for many of the benchmarks.
type insertErr int

const (
	iSuccess insertErr = -iota
	iShortDst
	iShortSrc
)

// insertFlush inserts the given rune in the buffer ordered by CCC.
// If a decomposition with multiple segments are encountered, they leading
// ones are flushed.
// It returns a non-zero error code if the rune was not inserted.
func (rb *reorderBuffer) insertFlush(src input, i int, info Properties) insertErr {
	if rune := src.hangul(i); rune != 0 {
		rb.decomposeHangul(rune)
		return iSuccess
	}
	if info.hasDecomposition() {
		return rb.insertDecomposed(info.Decomposition())
	}
	rb.insertSingle(src, i, info)
	return iSuccess
}

// insertUnsafe inserts the given rune in the buffer ordered by CCC.
// It is assumed there is sufficient space to hold the runes. It is the
// responsibility of the caller to ensure this. This can be done by checking
// the state returned by the streamSafe type.
func (rb *reorderBuffer) insertUnsafe(src input, i int, info Properties) {
	if rune := src.hangul(i); rune != 0 {
		rb.decomposeHangul(rune)
	}
	if info.hasDecomposition() {
		// TODO: inline.
		rb.insertDecomposed(info.Decomposition())
	} else {
		rb.insertSingle(src, i, info)
	}
}

// insertDecomposed inserts an entry in to the reorderBuffer for each rune
// in dcomp. dcomp must be a sequence of decomposed UTF-8-encoded runes.
// It flushes the buffer on each new segment start.
func (rb *reorderBuffer) insertDecomposed(dcomp []byte) insertErr {
	rb.tmpBytes.setBytes(dcomp)
	// As the streamSafe accounting already handles the counting for modifiers,
	// we don't ha