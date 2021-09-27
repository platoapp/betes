// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collate

import (
	"sort"

	"golang.org/x/text/internal/colltab"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

// newCollator creates a new collator with default options configured.
func newCollator(t colltab.Weighter) *Collator {
	// Initialize a collator with default options.
	c := &Collator{
		options: options{
			ignore: [colltab.NumLevels]bool{
				colltab.Quaternary: true,
				colltab.Identity:   true,
			},
			f: norm.NFD,
			t: t,
		},
	}

	// TODO: store vt in tags or remove.
	c.variableTop = t.Top()

	return c
}

// An Option is used to change the behavior of a Collator. Options override the
// settings passed through the locale identifier.
type Option struct {
	priority int
	f        func(o *options)
}

type prioritizedOptions []Option

func (p prioritizedOptions) Len() int {
	return len(p)
}

func (p prioritizedOptions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p prioritizedOptions) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

type options struct {
	// ignore specifies which levels to ignore.
	ignore [colltab.NumLevels]bool

	// caseLevel is true if there is an addition