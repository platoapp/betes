// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The trie in this file is used to associate the first full character in an
// UTF-8 string to a collation element. All but the last byte in a UTF-8 byte
// sequence are used to lookup offsets in the index table to be used for the
// next byte. The last byte is used to index into a table of collation elements.
// For a full description, see go.text/collate/build/trie.go.

package colltab

const blockSize = 64

type Trie struct {
	Index0  []uint16 // index for first byte (0xC0-0xFF)
	Values0 []uint32 // index for first byte (0x00-0x7F)
	Index   []uint16
	Values  []uint32
}

const (
	t1 = 0x00 // 0000 0000
	tx = 0x80 // 1000 0000
	t2 = 0xC0 // 1100 0000
	t3 = 0xE0 // 1110 0000
	t4 = 0xF0 // 1111 0000
	t5 = 0xF8 // 1111 1000
	t6 = 0xFC // 1111 1100
	te = 0xFE // 1111 1110
)

func (t *Trie) lookupValue(n uint16, b byte) Elem {
	return Elem(t.Values[int(n)<<6+int(b)])
}

// lookup returns the trie value for the first UTF-8 encoding in s and
// the width in bytes of this encoding. The size will be 0 if s does not
// hold enough bytes to complete the encoding. len(s) must be greater than 0.
func (t *Trie) lookup(s []byte) (v Elem, sz int) {
	c0 := s[0]
	switch {
	case c0 < tx:
		return Elem(t.Values0[c0]), 1
	case c0 