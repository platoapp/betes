/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzz

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// fuzzFuncMap is a map from a type to a fuzzFunc that handles that type.
type fuzzFuncMap map[reflect.Type]reflect.Value

// Fuzzer knows how to fill any object with random fields.
type Fuzzer struct {
	fuzzFuncs        fuzzFuncMap
	defaultFuzzFuncs fuzzFuncMap
	r                *rand.Rand
	nilChance        float64
	minElements      int
	maxElements      int
	maxDepth         int
}

// New returns a new Fuzzer. Customize your Fuzzer further by calling Funcs,
// RandSource, NilChance, or NumElements in any order.
func New() *Fuzzer {
	return NewWithSeed(time.Now().UnixNano())
}

func NewWithSeed(seed int64) *Fuzzer {
	f := &Fuzzer{
		defaultFuzzFuncs: fuzzFuncMap{
			reflect.TypeOf(&time.Time{}): reflect.ValueOf(fuzzTime),
		},

		fuzzFuncs:   fuzzFuncMap{},
		r:           rand.New(rand.NewSource(seed)),
		nilChance:   .2,
		minElements: 1,
		maxElements: 10,
		maxDepth:    100,
	}
	return f
}

// Funcs adds each entry in fuzzFuncs as a custom fuzzing function.
//
// Each entry in fuzzFuncs must be a function taking two parameters.
// The first parameter must be a pointer or map. It is the variable that
// function will fill with random data. The second parameter must be a
// fuzz.Continue, which will provide a source of randomness and a way
// to automatically continue fuzzing smaller pieces of the first parameter.
//
// These functions are called sensibly, e.g., if you wanted custom string
// fuzzing, the function `func(s *string, c fuzz.Continue)` would get
// called and passed the address of strings. Maps and pointers will always
// be made/new'd for you, ignoring the NilChange option. For slices, it
// doesn't make much sense to  pre-create them--Fuzzer doesn't know how
// long you want your slice--so take a pointer to a slice, and make it
// yourself. (If you don't want your map/pointer type pre-made, take a
// pointer to it, and make it yourself.) See the examples for a range of
// custom functions.
func (f *Fuzzer) Funcs(fuzzFuncs ...interface{}) *Fuzzer {
	for i := range fuzzFuncs {
		v := reflect.ValueOf(fuzzFuncs[i])
		if v.Kind() != reflect.Func {
			panic("Need only funcs!")
		}
		t := v.Type()
		if t.NumIn() != 2 || t.NumOut() != 0 {
			panic("Need 2 in and 0 out params!")
		}
		argT := t.In(0)
		switch argT.Kind() {
		case reflect.Ptr, reflect.Map:
		default:
			panic("fuzzFunc must take pointer or map type")
		}
		if t.In(1) != reflect.TypeOf(Continue{}) {
			panic("fuzzFunc's second parameter must be type fuzz.Continue")
		}
		f.fuzzFuncs[argT] = v
	}
	return f
}

// RandSource causes f to get values from the given source of randomness.
// Use if you want deterministic fuzzing.
func (f *Fuzzer) RandSource(s rand.Source) *Fuzzer {
	f.r = rand.New(s)
	return f
}

// NilChance sets the probability of creating a nil pointer, map, or slice to
// 'p'. 'p' should be between 0 (no nils) and 1 (all nils), inclusive.
func (f *Fuzzer) NilChance(p float64) *Fuzzer {
	if p < 0 || p > 1 {
		panic("p should be between 0 and 1, inclusive.")
	}
	f.nilChance = p
	return f
}

// NumElements sets the minimum and maximum number of elements that will be
// added to a non-nil map or slice.
func (f *Fuzzer) NumElements(atLeast, atMost int) *Fuzzer {
	if atLeast > atMost {
		panic("atLeast must be <= atMost")
	}
	if atLeast < 0 {
		panic("atLeast must be >= 0")
	}
	f.minElements = atLeast
	f.maxElements = atMost
	return f
}

func (f *Fuzzer) genElementCount() int {
	if f.minElements == f.maxElements {
		return f.minElements
	}
	return f.minElements + f.r.Intn(f.maxElements-f.minElements+1)
}

func (f *Fuzzer) genShouldFill() bool {
	return f.r.Float64() > f.nilChance
}

// MaxDepth sets the maximum number of recursive fuzz calls that will be made
// before stopping.  This includes struct members, pointers, and map and slice
// elements.
func (f *Fuzzer) MaxDepth(d int) *Fuzzer {
	f.maxDepth = d
	return f
}

// Fuzz recursively fills all of obj's fields with something random.  First
// this tries to find a custom fuzz function (see Funcs).  If there is no
// custom function this tests whether the object implements fuzz.Interface and,
// if so, calls Fuzz on it to fuzz itself.  If that fails, this will see if
// there is a default fuzz function provided by this package.  If all of that
// fails, this will generate random values for all primitive fields and then
// recurse for all non-primitives.
//
// This is safe for cyclic or tree-like structs, up to a limit.  Use the
// MaxDepth method to adjust how deep you need it to recurse.
//
// obj must be a pointer. Only exported (public) fields can be set (thanks,
// golang :/ ) Intended for tests, so will panic on bad input or unimplemented
// fields.
func (f *Fuzzer) Fuzz(obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		panic("needed ptr!")
	}
	v = v.Elem()
	f.fuzzWithContext(v, 0)
}

// FuzzNoCustom is just like Fuzz, except that any custom fuzz function for
// obj's type will not be called and obj will not be tested for fuzz.Interface
// conformance.  This applies only to obj and not other instances of obj's
// type.
// Not safe for cyclic or tree-like structs!
// obj must be a pointer. Only exported (public) fields can be set (thanks, golang :/ )
// Intended for tests, so will panic on bad input or unimplemented fields.
func (f *Fuzzer) FuzzNoCustom(obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.K