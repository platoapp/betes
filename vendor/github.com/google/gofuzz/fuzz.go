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
// pointer 