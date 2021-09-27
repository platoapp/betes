// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httplex contains rules around lexical matters of various
// HTTP-related specifications.
//
// This package is shared by the standard library (which vendors it)
// and x/net/http2. It comes with no API stability promise.
package httplex

import (
	"net"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

var isTokenTable = [127]bool{
	'!':  true,
	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,
	'*':  true,
	'+':  true,
	'-':  true,
	'.':  true,
	'0':  true,
	'1':  true,
	'2':  true,
	'3':  true,
	'4':  true,
	'5':  true,
	'6':  true,
	'7':  true,
	'8':  true,
	'9':  true,
	'A':  true,
	'B':  true,
	'C':  true,
	'D':  true,
	'E':  true,
	'F':  true,
	'G':  true,
	'H':  true,
	'I':  true,
	'J':  true,
	'K':  true,
	'L':  true,
	'M':  true,
	'N':  true,
	'O':  true,
	'P':  true,
	'Q':  true,
	'R':  true,
	'S':  true,
	'T':  true,
	'U':  true,
	'W':  true,
	'V':  true,
	'X':  true,
	'Y':  true,
	'Z':  true,
	'^':  true,
	'_':  true,
	'`':  true,
	'a':  true,
	'b':  true,
	'c':  true,
	'd':  true,
	'e':  true,
	'f':  true,
	'g':  true,
	'h':  true,
	'i':  true,
	'j':  true,
	'k':  true,
	'l':  true,
	'm':  true,
	'n':  true,
	'o':  true,
	'p':  true,
	'q':  true,
	'r':  true,
	's':  true,
	't':  true,
	'u':  true,
	'v':  true,
	'w':  true,
	'x':  true,
	'y':  true,
	'z':  true,
	'|':  true,
	'~':  true,
}

func IsTokenRune(r rune) bool {
	i := int(r)
	return i < len(isTokenTable) && isTokenTable[i]
}

func isNotToken(r rune) bool {
	return !IsTokenRune(r)
}

// HeaderValuesContainsToken reports whether any string in values
// contains the provided token, ASCII case-insensitively.
func HeaderValuesContainsToken(values []string, token string) bool {
	for _, v := range values {
		if headerValueContainsToken(v, token) {
			return true
		}
	}
	return false
}

// isOWS reports whether b is an optional whitespace byte, as defined
// by RFC 7230 section 3.2.3.
func isOWS(b byte) bool { return b == ' ' || b == '\t' }

// trimOWS returns x with all optional whitespace removes from the
// beginning and end.
func trimOWS(x string) string {
	// TODO: consider using strings.Trim(x, " \t") instead,
	// if and when it's fast enough. See issue 10292.
	// But this ASCII-only code will probably always beat UTF-8
	// aware code.
	for len(x) > 0 && isOWS(x[0]) {
		x = x[1:]
	}
	for len(x) > 0 && isOWS(x[len(x)-1]) {
		x = x[:len(x)-1]
	}
	return x
}

// headerValueContainsToken reports whether v (assumed to be a
// 0#element, in the ABNF extension described in RFC 7230 section 7)
// contains token amongst its comma-separated tokens, ASCII
// case-insensitively.
func he