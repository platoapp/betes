/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package require

import (
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	url "net/url"
	time "time"
)

// Condition uses a Comparison to assert a complex condition.
func Condition(t TestingT, comp assert.Comparison, msgAndArgs ...interface{}) {
	if assert.Condition(t, comp, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Conditionf uses a Comparison to assert a complex condition.
func Conditionf(t TestingT, comp assert.Comparison, msg string, args ...interface{}) {
	if assert.Conditionf(t, comp, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
func Contains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	if assert.Contains(t, s, contains, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) {
	if assert.Containsf(t, s, contains, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExists(t TestingT, path string, msgAndArgs ...interface{}) {
	if assert.DirExists(t, path, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExistsf(t TestingT, path string, msg string, args ...interface{}) {
	if assert.DirExistsf(t, path, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func ElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	if assert.ElementsMatch(t, listA, listB, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	if assert.ElementsMatchf(t, listA, listB, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	if assert.Empty(t, object, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) {
	if assert.Emptyf(t, object, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if assert.Equal(t, expected, actual, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) {
	if assert.EqualError(t, theError, errString, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) {
	if assert.EqualErrorf(t, theError, errString, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValues(t, uint32(123), int32(123))
func EqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if assert.EqualValues(t, expected, actual, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123, "error message %s", "formatted"), int32(123))
func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if assert.EqualValuesf(t, expected, actual, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if assert.Equalf(t, expected, actual, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
// 	   assert.Equal(t, expectedError, err)
//   }
func Error(t TestingT, err error, msgAndArgs ...interface{}) {
	if assert.Error(t, err, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
func Errorf(t TestingT, err error, msg string, args ...interface{}) {
	if assert.Errorf(t, err, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Exactly asserts that two objects are equal in value and type.
//
//    assert.Exactly(t, int32(123), int64(123))
func Exactly(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if assert.Exactly(t, expected, actual, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123, "error message %s", "formatted"), int64(123))
func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if assert.Exactlyf(t, expected, actual, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Fail reports a failure through
func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	if assert.Fail(t, failureMessage, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// FailNow fails test
func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	if assert.FailNow(t, failureMessage, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// FailNowf fails test
func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	if assert.FailNowf(t, failureMessage, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Failf reports a failure through
func Failf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	if assert.Failf(t, failureMessage, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// False asserts that the specified value is false.
//
//    assert.False(t, myBool)
func False(t TestingT, value bool, msgAndArgs ...interface{}) {
	if assert.False(t, value, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Falsef asserts that the specified value is false.
//
//    assert.Falsef(t, myBool, "error message %s", "formatted")
func Falsef(t TestingT, value bool, msg string, args ...interface{}) {
	if assert.Falsef(t, value, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.
func FileExists(t TestingT, path string, msgAndArgs ...interface{}) {
	if assert.FileExists(t, path, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.
func FileExistsf(t TestingT, path string, msg string, args ...interface{}) {
	if assert.FileExistsf(t, path, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if assert.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if assert.HTTPBodyContainsf(t, handler, method, url, values, str, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if assert.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if assert.HTTPBodyNotContainsf(t, handler, method, url, values, str, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if assert.HTTPError(t, handler, method, url, values, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//  assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if assert.HTTPErrorf(t, handler, method, url, values, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if assert.HTTPRedirect(t, handler, method, url, values, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if assert.HTTPRedirectf(t, handler, method, url, values, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns wh