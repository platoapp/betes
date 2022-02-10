// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go gen_common.go -output tables.go
//go:generate go run gen_index.go

package language

// TODO: Remove above NOTE after:
// - verifying that tables are dropped correctly (most notably matcher tables).

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// maxCoreSize is the maximum size of a BCP 47 tag without variants and
	// extensions. Equals max lang (3) + script (4) + max reg (3) + 2 dashes.
	maxCoreSize = 12

	// max99thPercentileSize is a somewhat arbitrary buffer size that presumably
	// is large enough to hold at least 99% of the BCP 47 tags.
	max99thPercentileSize = 32

	// maxSimpleUExtensionSize is the maximum size of a -u extension with one
	// key-type pair. Equals len("-u-") + key (2) + dash + max value (8).
	maxSimpleUExtensionSize = 14
)

// Tag represents a BCP 47 language tag. It is used to specify an instance of a
// specific language or locale. All language tag values are guaranteed to be
// well-formed.
type Tag struct {
	lang   langID
	region regionID
	// TODO: we will soon run out of positions for script. Idea: instead of
	// storing lang, region, and script codes, store only the compact index and
	// have a lookup table from this code to its expansion. This greatly speeds
	// up table lookup, speed up common variant cases.
	// This will also immediately free up 3 extra bytes. Also, the pVariant
	// field can now be moved to the lookup table, as the compact index uniquely
	// determines the offset of a possible variant.
	script   scriptID
	pVariant byte   // offset in str, includes preceding '-'
	pExt     uint16 // offset of first extension, includes preceding '-'

	// str is the string representation of the Tag. It will only be used if the
	// tag has variants or extensions.
	str string
}

// Make is a convenience wrapper for Parse that omits the error.
// In case of an error, a sensible default is returned.
func Make(s string) Tag {
	return Default.Make(s)
}

// Make is a convenience wrapper for c.Parse that omits the error.
// In case of an error, a sensible default is returned.
func (c CanonType) Make(s string) Tag {
	t, _ := c.Parse(s)
	return t
}

// Raw returns the raw base language, script and region, without making an
// attempt to infer their values.
func (t Tag) Raw() (b Base, s Script, r Region) {
	return Base{t.lang}, Script{t.script}, Region{t.region}
}

// equalTags compares language, script and region subtags only.
func (t Tag) equalTags(a Tag) bool {
	return t.lang == a.lang && t.script == a.script && t.region == a.region
}

// IsRoot returns true if t is equal to language "und".
func (t Tag) IsRoot() bool {
	if int(t.pVariant) < len(t.str) {
		return false
	}
	return t.equalTags(und)
}

// private reports whether the Tag consists solely of a private use tag.
func (t Tag) private() bool {
	return t.str != "" && t.pVariant == 0
}

// CanonType can be used to enable or disable various types of canonicalization.
type CanonType int

const (
	// Replace deprecated base languages with their preferred replacements.
	DeprecatedBase CanonType = 1 << iota
	// Replace deprecated scripts with their preferred replacements.
	DeprecatedScript
	// Replace deprecated regions with their preferred replacements.
	DeprecatedRegion
	// Remove redundant scripts.
	SuppressScript
	// Normalize legacy encodings. This includes legacy languages defined in
	// CLDR as well as bibliographic codes defined in ISO-639.
	Legacy
	// Map the dominant language of a macro language group to the macro language
	// subtag. For example cmn -> zh.
	Macro
	// The CLDR flag should be used if full compatibility with CLDR is required.
	// There are a few cases where language.Tag may differ from CLDR. To follow all
	// of CLDR's suggestions, use All|CLDR.
	CLDR

	// Raw can be used to Compose or Parse without Canonicalization.
	Raw CanonType = 0

	// Replace all deprecated tags with their preferred replacements.
	Deprecated = DeprecatedBase | DeprecatedScript | DeprecatedRegion

	// All canonicalizations recommended by BCP 47.
	BCP47 = Deprecated | SuppressScript

	// All canonicalizations.
	All = BCP47 | Legacy | Macro

	// Default is the canonicalization used by Parse, Make and Compose. To
	// preserve as much information as possible, canonicalizations that remove
	// potentially valuable information are not included. The Matcher is
	// designed to recognize similar tags that would be the same if
	// they were canonicalized using All.
	Default = Deprecated | Legacy

	canonLang = DeprecatedBase | Legacy | Macro

	// TODO: LikelyScript, LikelyRegion: suppress similar to ICU.
)

// canonicalize returns the canonicalized equivalent of the tag and
// whether there was any change.
func (t Tag) canonicalize(c CanonType) (Tag, bool) {
	if c == Raw {
		return t, false
	}
	changed := false
	if c&SuppressScript != 0 {
		if t.lang < langNoIndexOffset && uint8(t.script) == suppressScript[t.lang] {
			t.script = 0
			changed = true
		}
	}
	if c&canonLang != 0 {
		for {
			if l, aliasType := normLang(t.lang); l != t.lang {
				switch aliasType {
				case 