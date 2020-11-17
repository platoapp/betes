/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

// Conditionf uses a Comparison to assert a complex condition.
func Conditionf(t TestingT, comp Comparison, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return Condition(t, comp, append([]interface{}{msg}, args...)...)
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return Contains(t, s, contains, append([]interface{}{msg}, args...)...)
}

// DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExistsf(t TestingT, path string, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return DirExists(t, path, append([]interface{}{msg}, args...)...)
}

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return ElementsMatch(t, listA, listB, append([]interface{}{msg}, args...)...)
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return Empty(t, object, append([]interface{}{msg}, args...)...)
}

// Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return Equal(t, expected, actual, append([]interface{}{msg}, args...)...)
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return EqualError(t, theError, errString, append([]interface{}{msg}, args...)...)
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123, "error message %s", "formatted"), int32(123))
func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return EqualValues(t, expected, actual, append([]interface{}{msg}, args...)...)
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
func Errorf(t TestingT, err error, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return Error(t, err, append([]interface{}{msg}, args...)...)
}

// Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123, "error message %s", "formatted"), int64(123))
func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	if h, ok := t.(tHelp