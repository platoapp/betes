// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package width

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

type foldTransform struct {
	transform.NopResetter
}

func (foldTransform) Span(src []byte, atEOF bool) (n int, err error) {
	for n < len(src) {
		if src[n] < utf8.RuneSelf {
			// ASCII fast path.
			for n++; n < len(src) && src[n] < utf8.RuneSelf; n++ {
			}
			continue
		}
		v, size := trie.lookup(src[n:])
		if size == 0 { // incomplete UTF-8 encoding
			if !atEOF {
				err = transform.ErrShortSrc
			} else {
				n = len(src)
			}
			break
		}
		if elem(v)&tagNeedsFold != 0 {
			err = transform.ErrEndOfSpan
			break
		}
		n += size
	}
	return n, err
}

func (foldTransform) Transform(dst