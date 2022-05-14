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

// ssState is used for reporting the 