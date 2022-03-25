// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cldr

// This file implements the various inheritance constructs defined by LDML.
// See http://www.unicode.org/reports/tr35/#Inheritance_and_Validity
// for more details.

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// fieldIter iterates over fields in a struct. It includes
// fields of embedded structs.
type fieldIter struct {
	v        reflect.Value
	index, n []int
}

func iter(v reflect.Value) fieldIter {
	if v.Kind() != reflect.Struct {
		log.Panicf("value %v must be a struct", v)
	}
	i := fieldIter{
		v:     v,
		index: []int{0},
		n:     []int{v.NumField()},
	}
	i.descent()
	return i
}

func (i *fieldIter) descent() {
	for f := i.field(); f.Anonymous && f.Type.NumField() > 0; f = i.field() {
		i.index = append(i.index, 0)
		i.n = append(i.n, f.Type.NumField())
	}
}

func (i *fieldIter) done() bool {
	return len(i.index) == 1 && i.index[0] >= i.n[0]
}

func skip(f reflect.StructField) bool {
	return !f.Anonymous && (f.Name[0] < 'A' || f.Name[0] > 'Z')
}

func (i *fieldIter