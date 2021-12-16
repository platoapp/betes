// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package colltab contains functionality related to collation tables.
// It is only to be used by the collate and search packages.
package colltab // import "golang.org/x/text/internal/colltab"

import (
	"sort"

	"golang.org/x/text/language"
)

// MatchLang finds the index of t in tags, using a matching algorithm used for
// collation and search. tags[0] must be language.Und, the remaining tags should
// be sorted alphabetically.
//
// Language matching for collation and search is different from the matching
// defined by language.Matcher: the (inferred) base language must be an exact
// match for the relevant fields. For example, "gsw" should not match "de".
// Also the parent relation is different, as a parent may have a different
// script. So usually the parent of zh-Hant is und, whereas for MatchLang it is
// zh.
func MatchLang(t language.Tag, tags []language.Tag) int {
	// Canonicalize the values, including collapsing macro languages.
	t, _ = language.All.Canonicalize(t)

	base, conf := t.Base()
	// Estimate the base language, but only use high-confidence values.
	if conf < language.High {
		// The root locale supports "search" and "standard". We assume that any
		// implementation will only use one of both.
		return 0
	}

	// Maximize base and script and normalize the tag.
	if _, s, r := t.Raw(); (r != language.Region{}) {
		p, _ := language.Raw.Compose(base, s, r)
		// Taking the parent forces the script to be maximized.
		p = p.Parent()
		// Add back region and extensions.
		t, _ = language.Raw.Compose(p, r, t.Extensions())
	} else {
		// Set the maximized base language.
		t, _ = language.Raw.Compose(base, s, t.Extensions())
	}

	// Find start index of the language tag.
	start := 1 + sort.Search(len(tags)-1, func(i int) bool {
		b, _, _ := tags[i+1].Raw()
		return base.String() <= b.String()
	})
	if start < len(tags) {
		if b, _, _ := tags[start].Raw(); b != base {
			return 0
		}
	}

	// Besides the base language, script and region, only the collat