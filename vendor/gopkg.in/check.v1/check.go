
// Package check is a rich testing extension for Go's testing package.
//
// For details about the project, see:
//
//     http://labix.org/gocheck
//
package check

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// -----------------------------------------------------------------------
// Internal type which deals with suite method calling.

const (
	fixtureKd = iota
	testKd
)

type funcKind int

const (
	succeededSt = iota
	failedSt
	skippedSt
	panickedSt
	fixturePanickedSt
	missedSt
)

type funcStatus uint32

// A method value can't reach its own Method structure.
type methodType struct {
	reflect.Value
	Info reflect.Method
}

func newMethod(receiver reflect.Value, i int) *methodType {
	return &methodType{receiver.Method(i), receiver.Type().Method(i)}
}

func (method *methodType) PC() uintptr {
	return method.Info.Func.Pointer()
}

func (method *methodType) suiteName() string {
	t := method.Info.Type.In(0)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func (method *methodType) String() string {
	return method.suiteName() + "." + method.Info.Name
}

func (method *methodType) matches(re *regexp.Regexp) bool {
	return (re.MatchString(method.Info.Name) ||
		re.MatchString(method.suiteName()) ||