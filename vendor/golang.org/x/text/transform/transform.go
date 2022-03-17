
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transform provides reader and writer wrappers that transform the
// bytes passing through as well as various transformations. Example
// transformations provided by other packages include normalization and
// conversion between character sets.
package transform // import "golang.org/x/text/transform"

import (
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

var (