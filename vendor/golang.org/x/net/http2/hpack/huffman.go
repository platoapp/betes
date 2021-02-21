// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpack

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// HuffmanDecode decodes the string in v and writes the expanded
// result to w, returning the number of bytes written to w and the
// Write call's return value. At most one Write call is made.
func HuffmanDecode(w io.Writer, v []byte) (int, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)
	if err := huffmanDecode(buf, 0, v); err != nil {
		return 0, err
	}
	return w.Write(buf.Bytes())
}

// HuffmanDecodeToString decodes the string in v.
func HuffmanDecodeToString(v []byte) (string, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)
	if err := huffmanDecode(buf, 0, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ErrInvalidHuffman is returned for errors found decoding
// Huffman-encoded strings.
var ErrInvalidHuffman = errors.New("hpack: invalid Huffman-encoded data")

// huffmanDecode decodes v to buf.
// If maxLen is greater than 0, attempts to write more to buf than
// maxLen bytes will return ErrStringLength.
func huffmanDecode(buf *bytes.Buffer, maxLen int, v []byte) error {
	n := rootHuffmanNode
	// cur is the bit buffer that has not been fed into n.
	// cbits is the number of low order bits in cur that are valid.
	// sbits is the number of bits of the symbol prefix being decoded.
	cur, cbits, sbits := uint(0), uint8(0), uint8(0)
	for _, b := range v {
		cur = cur<<8 | uint(b)
		cbits += 8
		sbits += 8
		for cbits >= 8 {
			idx := byte(cur >> (cbits - 8))
			n = n.children[idx]
			if n == nil {
				return ErrInvalidHuffman
			}
			if n.children == nil {
				if maxLen != 0 && buf.Len() == maxLen {
					return ErrStringLength
				}
				buf.WriteByte(n.sym)
				cbits -= n.codeLen
				n = rootHuffmanNode
				sbits = cbits
			} else {
				cbits -= 8
			}
		}
	}
	for cbits > 0 {
		n = n.children[byte(cur<<(8-cbits))]
		if n == nil {
			return ErrInvalidHuffman
		}
		if n.children != nil || n.codeLen > cbits {
			break
		}
		if maxLen != 0 && buf.Len() == maxLen {
			return ErrStringLength
		}
		buf.WriteByte(n.sym)
		cbits -= n.codeLen
		n = rootHuffmanNode
		sbits = cbits
	}
	if sbits > 7 {
		// Either there was an incomplete symbol, or overlong padding.
		// Both are decoding errors per RFC 7541 section 5.2.
		