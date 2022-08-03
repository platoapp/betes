
/*
Copyright 2014 The Kubernetes Authors.

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

package resource

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/go-openapi/spec"
	inf "gopkg.in/inf.v0"
	openapi "k8s.io/kube-openapi/pkg/common"
)

// Quantity is a fixed-point representation of a number.
// It provides convenient marshaling/unmarshaling in JSON and YAML,
// in addition to String() and Int64() accessors.
//
// The serialization format is:
//
// <quantity>        ::= <signedNumber><suffix>
//   (Note that <suffix> may be empty, from the "" case in <decimalSI>.)
// <digit>           ::= 0 | 1 | ... | 9
// <digits>          ::= <digit> | <digit><digits>
// <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits>
// <sign>            ::= "+" | "-"
// <signedNumber>    ::= <number> | <sign><number>
// <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI>
// <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei
//   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)
// <decimalSI>       ::= m | "" | k | M | G | T | P | E
//   (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)
// <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber>
//
// No matter which of the three exponent forms is used, no quantity may represent
// a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal
// places. Numbers larger or more precise will be capped or rounded up.
// (E.g.: 0.1m will rounded up to 1m.)
// This may be extended in the future if we require larger or smaller quantities.
//
// When a Quantity is parsed from a string, it will remember the type of suffix
// it had, and will use the same type again when it is serialized.
//
// Before serializing, Quantity will be put in "canonical form".
// This means that Exponent/suffix will be adjusted up or down (with a
// corresponding increase or decrease in Mantissa) such that:
//   a. No precision is lost
//   b. No fractional digits will be emitted
//   c. The exponent (or suffix) is as large as possible.
// The sign will be omitted unless the number is negative.
//
// Examples:
//   1.5 will be serialized as "1500m"
//   1.5Gi will be serialized as "1536Mi"
//
// NOTE: We reserve the right to amend this canonical format, perhaps to
//   allow 1.5 to be canonical.
// TODO: Remove above disclaimer after all bikeshedding about format is over,
//   or after March 2015.
//
// Note that the quantity will NEVER be internally represented by a
// floating point number. That is the whole point of this exercise.
//
// Non-canonical values will still parse as long as they are well formed,
// but will be re-emitted in their canonical form. (So always use canonical
// form, or don't diff.)
//
// This format is intended to make it difficult to use these numbers without
// writing some sort of special handling code in the hopes that that will
// cause implementors to also use a fixed point implementation.
//
// +protobuf=true
// +protobuf.embed=string
// +protobuf.options.marshal=false
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type Quantity struct {
	// i is the quantity in int64 scaled form, if d.Dec == nil
	i int64Amount
	// d is the quantity in inf.Dec form if d.Dec != nil
	d infDecAmount
	// s is the generated value of this quantity to avoid recalculation
	s string

	// Change Format at will. See the comment for Canonicalize for
	// more details.
	Format
}

// CanonicalValue allows a quantity amount to be converted to a string.
type CanonicalValue interface {
	// AsCanonicalBytes returns a byte array representing the string representation
	// of the value mantissa and an int32 representing its exponent in base-10. Callers may
	// pass a byte slice to the method to avoid allocations.
	AsCanonicalBytes(out []byte) ([]byte, int32)
	// AsCanonicalBase1024Bytes returns a byte array representing the string representation
	// of the value mantissa and an int32 representing its exponent in base-1024. Callers
	// may pass a byte slice to the method to avoid allocations.
	AsCanonicalBase1024Bytes(out []byte) ([]byte, int32)
}

// Format lists the three possible formattings of a quantity.
type Format string

const (
	DecimalExponent = Format("DecimalExponent") // e.g., 12e6
	BinarySI        = Format("BinarySI")        // e.g., 12Mi (12 * 2^20)
	DecimalSI       = Format("DecimalSI")       // e.g., 12M  (12 * 10^6)
)

// MustParse turns the given string into a quantity or panics; for tests
// or others cases where you know the string is valid.
func MustParse(str string) Quantity {
	q, err := ParseQuantity(str)
	if err != nil {
		panic(fmt.Errorf("cannot parse '%v': %v", str, err))
	}
	return q
}

const (
	// splitREString is used to separate a number from its suffix; as such,
	// this is overly permissive, but that's OK-- it will be checked later.
	splitREString = "^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
)

var (
	// splitRE is used to get the various parts of a number.
	splitRE = regexp.MustCompile(splitREString)

	// Errors that could happen while parsing a string.
	ErrFormatWrong = errors.New("quantities must match the regular expression '" + splitREString + "'")
	ErrNumeric     = errors.New("unable to parse numeric part of quantity")
	ErrSuffix      = errors.New("unable to parse quantity's suffix")
)

// parseQuantityString is a fast scanner for quantity values.
func parseQuantityString(str string) (positive bool, value, num, denom, suffix string, err error) {
	positive = true
	pos := 0
	end := len(str)

	// handle leading sign
	if pos < end {
		switch str[0] {
		case '-':
			positive = false
			pos++
		case '+':
			pos++
		}
	}

	// strip leading zeros
Zeroes:
	for i := pos; ; i++ {
		if i >= end {
			num = "0"
			value = num
			return
		}
		switch str[i] {
		case '0':
			pos++
		default:
			break Zeroes
		}
	}

	// extract the numerator
Num:
	for i := pos; ; i++ {
		if i >= end {
			num = str[pos:end]
			value = str[0:end]
			return
		}
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			num = str[pos:i]
			pos = i
			break Num
		}
	}

	// if we stripped all numerator positions, always return 0
	if len(num) == 0 {
		num = "0"
	}

	// handle a denominator
	if pos < end && str[pos] == '.' {
		pos++
	Denom:
		for i := pos; ; i++ {
			if i >= end {
				denom = str[pos:end]
				value = str[0:end]
				return
			}
			switch str[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				denom = str[pos:i]
				pos = i
				break Denom
			}
		}
		// TODO: we currently allow 1.G, but we may not want to in the future.
		// if len(denom) == 0 {
		// 	err = ErrFormatWrong
		// 	return
		// }
	}
	value = str[0:pos]

	// grab the elements of the suffix
	suffixStart := pos
	for i := pos; ; i++ {
		if i >= end {
			suffix = str[suffixStart:end]
			return
		}
		if !strings.ContainsAny(str[i:i+1], "eEinumkKMGTP") {
			pos = i
			break
		}
	}
	if pos < end {
		switch str[pos] {
		case '-', '+':
			pos++
		}
	}
Suffix:
	for i := pos; ; i++ {
		if i >= end {
			suffix = str[suffixStart:end]
			return
		}
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			break Suffix
		}
	}
	// we encountered a non decimal in the Suffix loop, but the last character
	// was not a valid exponent
	err = ErrFormatWrong
	return
}

// ParseQuantity turns str into a Quantity, or returns an error.
func ParseQuantity(str string) (Quantity, error) {
	if len(str) == 0 {
		return Quantity{}, ErrFormatWrong
	}
	if str == "0" {
		return Quantity{Format: DecimalSI, s: str}, nil
	}

	positive, value, num, denom, suf, err := parseQuantityString(str)
	if err != nil {
		return Quantity{}, err
	}

	base, exponent, format, ok := quantitySuffixer.interpret(suffix(suf))
	if !ok {
		return Quantity{}, ErrSuffix
	}

	precision := int32(0)
	scale := int32(0)
	mantissa := int64(1)
	switch format {
	case DecimalExponent, DecimalSI:
		scale = exponent
		precision = maxInt64Factors - int32(len(num)+len(denom))
	case BinarySI:
		scale = 0
		switch {
		case exponent >= 0 && len(denom) == 0:
			// only handle positive binary numbers with the fast path
			mantissa = int64(int64(mantissa) << uint64(exponent))
			// 1Mi (2^20) has ~6 digits of decimal precision, so exponent*3/10 -1 is roughly the precision
			precision = 15 - int32(len(num)) - int32(float32(exponent)*3/10) - 1
		default:
			precision = -1
		}
	}

	if precision >= 0 {
		// if we have a denominator, shift the entire value to the left by the number of places in the
		// denominator
		scale -= int32(len(denom))
		if scale >= int32(Nano) {
			shifted := num + denom

			var value int64
			value, err := strconv.ParseInt(shifted, 10, 64)
			if err != nil {
				return Quantity{}, ErrNumeric
			}
			if result, ok := int64Multiply(value, int64(mantissa)); ok {
				if !positive {
					result = -result
				}
				// if the number is in canonical form, reuse the string
				switch format {
				case BinarySI:
					if exponent%10 == 0 && (value&0x07 != 0) {
						return Quantity{i: int64Amount{value: result, scale: Scale(scale)}, Format: format, s: str}, nil
					}
				default:
					if scale%3 == 0 && !strings.HasSuffix(shifted, "000") && shifted[0] != '0' {
						return Quantity{i: int64Amount{value: result, scale: Scale(scale)}, Format: format, s: str}, nil
					}
				}
				return Quantity{i: int64Amount{value: result, scale: Scale(scale)}, Format: format}, nil
			}
		}
	}

	amount := new(inf.Dec)
	if _, ok := amount.SetString(value); !ok {
		return Quantity{}, ErrNumeric
	}

	// So that no one but us has to think about suffixes, remove it.
	if base == 10 {
		amount.SetScale(amount.Scale() + Scale(exponent).infScale())
	} else if base == 2 {
		// numericSuffix = 2 ** exponent
		numericSuffix := big.NewInt(1).Lsh(bigOne, uint(exponent))
		ub := amount.UnscaledBig()
		amount.SetUnscaledBig(ub.Mul(ub, numericSuffix))
	}

	// Cap at min/max bounds.
	sign := amount.Sign()
	if sign == -1 {
		amount.Neg(amount)
	}

	// This rounds non-zero values up to the minimum representable value, under the theory that
	// if you want some resources, you should get some resources, even if you asked for way too small
	// of an amount.  Arguably, this should be inf.RoundHalfUp (normal rounding), but that would have
	// the side effect of rounding values < .5n to zero.
	if v, ok := amount.Unscaled(); v != int64(0) || !ok {
		amount.Round(amount, Nano.infScale(), inf.RoundUp)
	}

	// The max is just a simple cap.
	// TODO: this prevents accumulating quantities greater than int64, for instance quota across a cluster
	if format == BinarySI && amount.Cmp(maxAllowed.Dec) > 0 {
		amount.Set(maxAllowed.Dec)
	}

	if format == BinarySI && amount.Cmp(decOne) < 0 && amount.Cmp(decZero) > 0 {
		// This avoids rounding and hopefully confusion, too.
		format = DecimalSI
	}
	if sign == -1 {
		amount.Neg(amount)
	}

	return Quantity{d: infDecAmount{amount}, Format: format}, nil
}

// DeepCopy returns a deep-copy of the Quantity value.  Note that the method
// receiver is a value, so we can mutate it in-place and return it.
func (q Quantity) DeepCopy() Quantity {
	if q.d.Dec != nil {
		tmp := &inf.Dec{}
		q.d.Dec = tmp.Set(q.d.Dec)
	}
	return q
}

// OpenAPIDefinition returns openAPI definition for this type.
func (_ Quantity) OpenAPIDefinition() openapi.OpenAPIDefinition {
	return openapi.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:   []string{"string"},
				Format: "",
			},
		},
	}
}

// CanonicalizeBytes returns the canonical form of q and its suffix (see comment on Quantity).
//
// Note about BinarySI:
// * If q.Format is set to BinarySI and q.Amount represents a non-zero value between
//   -1 and +1, it will be emitted as if q.Format were DecimalSI.
// * Otherwise, if q.Format is set to BinarySI, fractional parts of q.Amount will be
//   rounded up. (1.1i becomes 2i.)
func (q *Quantity) CanonicalizeBytes(out []byte) (result, suffix []byte) {
	if q.IsZero() {
		return zeroBytes, nil
	}

	var rounded CanonicalValue
	format := q.Format
	switch format {
	case DecimalExponent, DecimalSI:
	case BinarySI:
		if q.CmpInt64(-1024) > 0 && q.CmpInt64(1024) < 0 {
			// This avoids rounding and hopefully confusion, too.
			format = DecimalSI
		} else {
			var exact bool
			if rounded, exact = q.AsScale(0); !exact {
				// Don't lose precision-- show as DecimalSI
				format = DecimalSI
			}
		}
	default:
		format = DecimalExponent
	}

	// TODO: If BinarySI formatting is requested but would cause rounding, upgrade to
	// one of the other formats.
	switch format {
	case DecimalExponent, DecimalSI:
		number, exponent := q.AsCanonicalBytes(out)
		suffix, _ := quantitySuffixer.constructBytes(10, exponent, format)
		return number, suffix
	default:
		// format must be BinarySI
		number, exponent := rounded.AsCanonicalBase1024Bytes(out)
		suffix, _ := quantitySuffixer.constructBytes(2, exponent*10, format)
		return number, suffix
	}
}

// AsInt64 returns a representation of the current value as an int64 if a fast conversion
// is possible. If false is returned, callers must use the inf.Dec form of this quantity.
func (q *Quantity) AsInt64() (int64, bool) {