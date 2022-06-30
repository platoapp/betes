
// Package inf (type inf.Dec) implements "infinite-precision" decimal
// arithmetic.
// "Infinite precision" describes two characteristics: practically unlimited
// precision for decimal number representation and no support for calculating
// with any specific fixed precision.
// (Although there is no practical limit on precision, inf.Dec can only
// represent finite decimals.)
//
// This package is currently in experimental stage and the API may change.
//
// This package does NOT support:
//  - rounding to specific precisions (as opposed to specific decimal positions)
//  - the notion of context (each rounding must be explicit)
//  - NaN and Inf values, and distinguishing between positive and negative zero
//  - conversions to and from float32/64 types
//
// Features considered for possible addition:
//  + formatting options
//  + Exp method
//  + combined operations such as AddRound/MulAdd etc
//  + exchanging data in decimal32/64/128 formats
//
package inf // import "gopkg.in/inf.v0"

// TODO:
//  - avoid excessive deep copying (quo and rounders)

import (
	"fmt"
	"io"
	"math/big"
	"strings"
)

// A Dec represents a signed arbitrary-precision decimal.
// It is a combination of a sign, an arbitrary-precision integer coefficient
// value, and a signed fixed-precision exponent value.
// The sign and the coefficient value are handled together as a signed value
// and referred to as the unscaled value.
// (Positive and negative zero values are not distinguished.)
// Since the exponent is most commonly non-positive, it is handled in negated
// form and referred to as scale.
//
// The mathematical value of a Dec equals:
//
//  unscaled * 10**(-scale)
//
// Note that different Dec representations may have equal mathematical values.
//
//  unscaled  scale  String()
//  -------------------------
//         0      0    "0"
//         0      2    "0.00"
//         0     -2    "0"
//         1      0    "1"
//       100      2    "1.00"
//        10      0   "10"
//         1     -1   "10"
//
// The zero value for a Dec represents the value 0 with scale 0.
//
// Operations are typically performed through the *Dec type.
// The semantics of the assignment operation "=" for "bare" Dec values is
// undefined and should not be relied on.
//
// Methods are typically of the form:
//
//	func (z *Dec) Op(x, y *Dec) *Dec
//
// and implement operations z = x Op y with the result as receiver; if it
// is one of the operands it may be overwritten (and its memory reused).
// To enable chaining of operations, the result is also returned. Methods
// returning a result other than *Dec take one of the operands as the receiver.
//
// A "bare" Quo method (quotient / division operation) is not provided, as the
// result is not always a finite decimal and thus in general cannot be
// represented as a Dec.
// Instead, in the common case when rounding is (potentially) necessary,
// QuoRound should be used with a Scale and a Rounder.
// QuoExact or QuoRound with RoundExact can be used in the special cases when it
// is known that the result is always a finite decimal.
//
type Dec struct {