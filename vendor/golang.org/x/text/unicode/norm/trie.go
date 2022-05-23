// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package norm

type valueRange struct {
	value  uint16 // header: value:stride
	lo, hi byte   // header: lo:n