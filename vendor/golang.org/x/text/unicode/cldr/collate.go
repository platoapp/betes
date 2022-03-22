// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cldr

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// RuleProcessor can be passed to Collator's Process method, which
// parses the rules and calls the respective method for each rule found.
type RuleProcessor interface {
	Reset(anchor string, before int) error
	Insert(level int, str, context, extend string) error
	Index(id string)
}

const (
	// cldrIndex is a Unicode-reserved sentinel value used to mark the start
	// of a grouping within an index.
	// We ignore any rule that starts with this rune.
	// See http://unicode.org/reports/tr35/#Collation_Elements for details.
	cldrIndex = "\uFDD0"

	// specialAnchor is the format in which to represent logical reset positions,
	// such as "first tertiary ignorable".
	specialAnchor = "<%s/>"
)

// Process parses the rules for the tailorings of this collation
// and calls the respective methods of p for each rule found.
func (c Collation) Process(p RuleProcessor) (err error) {
	if len(c.Cr) > 0 {
		if len(c.Cr) > 1 {
			return fmt.Errorf("multiple cr elements, want 0 or 1")
		}
		return processRules(p, c.Cr[0].Data())
	}
	if c.Rules.Any != nil {
		return c.processXML(p)
	}
	return errors.New("no tailoring data")
}

// processRules parses rules in the Collation Rule Syntax defined in
// http://www.unicode.org/reports/tr35/tr35-collation.html#Collation_Tailorings.
func processRules(p RuleProcessor, s string) (err error) {
	chk := func(s string, e error) string {
		if err == nil {
			err = e
		}
		return s
	}
	i := 0 // Save the line number for use after the loop.
	scanner := bufio.NewScanner(strings.NewReader(s))
	for ; scanner.Scan() && err == nil; i++ {
		for s := skipSpace(scanner.Text()); s != "" && s[0] != '#'; s = skipSpace(s) {
			level := 5
			var ch byte
			switch ch, s = s[0], s[1:]; ch {
			case '&': // followed by <anchor> or '[' <key> ']'
				if s = skipSpace(s); consume(&s, '[') {
					s = chk(parseSpecialAnchor(p, s))
				} else {
					s = chk(parseAnchor(p, 0, s))
				}
			case '<': // sort relation '<'{1,4}, optionally followed by '*'.
				for level = 1; consume(&s, '<'); level++ {
				}
				if level > 4 {
					err = fmt.Errorf("level %d > 4", level)
				}
				fallthrough
			case '=': // identity relation, optionally followed by *.
				if consume(&s, '*') {
					s = chk(parseSequence(p, level, s))
				} else {
					s = chk(parseOrder(p, level, s))
				}
			default:
				chk("", fmt.Errorf("illegal operator %q", ch))
				break
			}
		}
	}
	if chk("", scanner.Err()); err != nil {
		return fmt.Errorf("%d: %v", i, err)
	}
	return nil
}

// parseSpecialAnchor parses the anchor syntax which is either of the form
//    ['before' <level>] <anchor>
// or
//    [<label>]
// The starting should already be consumed.
func parseSpecialAnchor(p RuleProcessor, s string) (tail string, err error) {
	i := strings.IndexByte(s, ']')
	if i == -1 {
		return "", errors.New("unmatched bracket")
	}
	a := strings.TrimSpace(s[:i])
	s = s[i+1:]
	if strings.HasPrefix(a, "before ") {
		l, err := strconv.ParseUint(skipSpace(a[len("before "):]), 10, 3)
		if err != nil {
			return s, err
		}
		return parseAnchor(p, int(l), s)
	}
	return s, p.Reset(fmt.Sprintf(specialAnchor, a), 0)
}

func parseAnchor(p RuleProcessor, level int, s string) (tail string, err error) {
	anchor, s, err := scanString(s)
	if err != nil {
		return s, err
	}
	return s, p.Reset(anchor, level)
}

func parseOrder(p RuleProcessor, level int, s string) (tail string, err error) {
	var value, context, extend st