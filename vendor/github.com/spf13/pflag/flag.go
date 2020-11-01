// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package pflag is a drop-in replacement for Go's flag package, implementing
POSIX/GNU-style --flags.

pflag is compatible with the GNU extensions to the POSIX recommendations
for command-line options. See
http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

Usage:

pflag is a drop-in replacement of Go's native flag package. If you import
pflag under the name "flag" then all code should continue to function
with no changes.

	import flag "github.com/spf13/pflag"

There is one exception to this: if you directly instantiate the Flag struct
there is one more field "Shorthand" that you will need to set.
Most code never instantiates this struct directly, and instead uses
functions such as String(), BoolVar(), and Var(), and is therefore
unaffected.

Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -flagname, stored in the pointer ip, with type *int.
	var ip = flag.Int("flagname", 1234, "help message for flagname")
If you like, you can bind the flag to a variable using the Var() functions.
	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}
Or you can create custom flags that satisfy the Value interface (with
pointer receivers) and couple them to flag parsing by
	flag.Var(&flagVal, "name", "help message for flagname")
For such flags, the default value is just the initial value of the variable.

After all flags are defined, call
	flag.Parse()
to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves,
they are all pointers; if you bind to variables, they're values.
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

After parsing, the arguments after the flag are available as the
slice flag.Args() or individually as flag.Arg(i).
The arguments are indexed from 0 through flag.NArg()-1.

The pflag package also defines some new functions that are not in flag,
that give one-letter shorthands for flags. You can use these by appending
'P' to the name of any function that defines a flag.
	var ip = flag.IntP("flagname", "f", 1234, "help message")
	var flagvar bool
	func init() {
		flag.BoolVarP("boolname", "b", true, "help message")
	}
	flag.VarP(&flagVar, "varname", "v", 1234, "help message")
Shorthand letters can be used with single dashes on the command line.
Boolean shorthand flags can be combined with other shorthand flags.

Command line flag syntax:
	--flag    // boolean flags only
	--flag=x

Unlike the flag package, a single dash before an option means something
different than a double dash. Single dashes signify a series of shorthand
letters for flags. All but the last shorthand letter must be boolean flags.
	// boolean flags
	-f
	-abc
	// non-boolean flags
	-n 1234
	-Ifile
	// mixed
	-abcs "hello"
	-abcn1234

Flag parsing stops after the terminator "--". Unlike the flag package,
flags can be interspersed with arguments anywhere on the command line
before this terminator.

Integer flags accept 1234, 0664, 0x1234 and may be negative.
Boolean flags (in their long form) accept 1, 0, t, f, true, false,
TRUE, FALSE, True, False.
Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by
top-level functions.  The FlagSet type allows one to define
independent sets of flags, such as to implement subcommands
in a command-line interface. The methods of FlagSet are
analogous to the top-level functions for the command-line
flag set.
*/
package pflag

import (
	"bytes"
	"errors"
	goflag "flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// ErrHelp is the error returned if the flag -help is invoked but no such flag is defined.
var ErrHelp = errors.New("pflag: help requested")

// ErrorHandling defines how to handle flag parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota
	// ExitOnError will call os.Exit(2) if an error is found when parsing
	ExitOnError
	// PanicOnError will panic() if an error is found when parsing flags
	PanicOnError
)

// ParseErrorsWhitelist defines the parsing errors that can be ignored
type ParseErrorsWhitelist struct {
	// UnknownFlags will ignore unknown flags errors and continue parsing rest of the flags
	UnknownFlags bool
}

// NormalizedName is a flag name that has been normalized according to rules
// for the FlagSet (e.g. making '-' and '_' equivalent).
type NormalizedName string

// A FlagSet represents a set of defined flags.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage func()

	// SortFlags is used to indicate, if user wants to have sorted flags in
	// help/usage messages.
	SortFlags bool

	// ParseErrorsWhitelist is used to configure a whitelist of errors
	ParseErrorsWhitelist ParseErrorsWhitelist

	name              string
	parsed            bool
	actual            map[NormalizedName]*Flag
	orderedActual     []*Flag
	sortedActual      []*Flag
	formal            map[NormalizedName]*Flag
	orderedFormal     []*Flag
	sortedFormal      []*Flag
	shorthands        map[byte]*Flag
	args              []string // arguments after flags
	argsLenAtDash     int      // len(args) when a '--' was located when parsing, or -1 if no --
	errorHandling     ErrorHandling
	output            io.Writer // nil means stderr; use out() accessor
	interspersed      bool      // allow interspersed option/non-option args
	normalizeNameFunc func(f *FlagSet, name string) NormalizedName

	addedGoFlagSets []*goflag.FlagSet
}

// A Flag represents the state of a flag.
type Flag struct {
	Name                string              // name as it appears on command line
	Shorthand           string              // one-letter abbreviated flag
	Usage               string              // help message
	Value               Value               // value as set
	DefValue            string              // default value (as text); for usage message
	Changed             bool                // If the user set the value (or if left to default)
	NoOptDefVal         string              // default value (as text); if the flag is on the command line without any options
	Deprecated          string              // If this flag is deprecated, this string is the new or now thing to use
	Hidden              bool                // used by cobra.Command to allow flags to be hidden from help/usage text
	ShorthandDeprecated string              // If the shorthand of this flag is deprecated, this string is the new or now thing to use
	Annotations         map[string][]string // used by cobra.Command bash autocomple code
}

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
type Value interface {
	String() string
	Set(string) error
	Type() string
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[NormalizedName]*Flag) []*Flag {
	list := make(sort.StringSlice, len(flags))
	i := 0
	for k := range flags {
		list[i] = string(k)
		i++
	}
	list.Sort()
	result := make([]*Flag, len(list))
	for i, name := range list {
		result[i] = flags[NormalizedName(name)]
	}
	return result
}

// SetNormalizeFunc allows you to add a function which can translate flag names.
// Flags added to the FlagSet will be translated and then when anything tries to
// look up the flag that will also be translated. So it would be possible to create
// a flag named "getURL" and have it translated to "geturl".  A user could then pass
// "--getUrl" which may also be translated to "geturl" and everything will work.
func (f *FlagSet) SetNormalizeFunc(n func(f *FlagSet, name string) NormalizedName) {
	f.normalizeNameFunc = n
	f.sortedFormal = f.sortedFormal[:0]
	for fname, flag := range f.formal {
		nname := f.normalizeFlagName(flag.Name)
		if fname == nname {
			continue
		}
		flag.Name = string(nname)
		delete(f.formal, fname)
		f.formal[nname] = flag
		if _, set := f.actual[fname]; set {
			delete(f.actual, fname)
			f.actual[nname] = flag
		}
	}
}

// GetNormalizeFunc returns the previously set NormalizeFunc of a function which
// does no translation, if not set previously.
func (f *FlagSet) GetNormalizeFunc() func(f *FlagSet, name string) NormalizedName {
	if f.normalizeNameFunc != nil {
		return f.normalizeNameFunc
	}
	return func(f *FlagSet, name string) NormalizedName { return NormalizedName(name) }
}

func (f *FlagSet) normalizeFlagName(name string) NormalizedName {
	n := f.GetNormalizeFunc()
	return n(f, name)
}

func (f *FlagSet) out() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *FlagSet) SetOutput(output io.Writer) {
	f.output = output
}

// VisitAll visits the flags in lexicographical order or
// in primordial order if f.SortFlags is false, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*Flag)) {
	if len(f.formal) == 0 {
		return
	}

	var flags []*Flag
	if f.SortFlags {
		if len(f.formal) != len(f.sortedFormal) {
			f.sortedFormal = sortFlags(f.formal)
		}
		flags = f.sortedFormal
	} else {
		flags = f.orderedFormal
	}

	for _, flag := range flags {
		fn(flag)
	}
}

// HasFlags returns a bool to indicate if the FlagSet has any flags defined.
func (f *FlagSet) HasFlags() bool {
	return len(f.formal) > 0
}

// HasAvailableFlags returns a bool to indicate if the FlagSet has any flags
// that are not hidden.
func (f *FlagSet) HasAvailableFlags() bool {
	for _, flag := range f.formal {
		if !flag.Hidden {
			return true
		}
	}
	return false
}

// VisitAll visits the command-line flags in lexicographical order or
// in primordial order if f.SortFlags is false, calling fn for each.
// It visits all flags, even those not set.
func VisitAll(fn func(*Flag)) {
	CommandLine.VisitAll(fn)
}

// Visit visits the flags in lexicographical order or
// in primordial order if f.SortFlags is false, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*Flag)) {
	if len(f.actual) == 0 {
		return
	}

	var flags []*Flag
	if f.SortFlags {
		if len(f.actual) != len(f.sortedActual) {
			f.sortedActual = sortFlags(f.actual)
		}
		flags = f.sortedActual
	} else {
		flags = f.orderedActual
	}

	for _, flag := range flags {
		fn(flag)
	}
}

// Visit visits the command-line flags in lexicographical order or
// in primordial order if f.SortFlags is false, calling fn for each.
// It visits only those flags that have been set.
func Visit(fn func(*Flag)) {
	CommandLine.Visit(fn)
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
	return f.lookup(f.normalizeFlagName(name))
}

// ShorthandLookup returns the Flag structure of the short handed flag,
// returning nil if none exists.
// It panics, if len(name) > 1.
func (f *FlagSet) ShorthandLookup(name string) *Flag {
	if name == "" {
		return nil
	}
	if len(name) > 1 {
		msg := fmt.Sprintf("can not look up shorthand which is more than one ASCII character: %q", name)
		fmt.Fprintf(f.out(), msg)
		panic(msg)
	}
	c := name[0]
	return f.shorthands[c]
}

// lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) lookup(name NormalizedName) *Flag {
	return f.formal[name]
}

// func to return a given type for a given flag name
func (f *FlagSet) getFlagType(name string, ftype string, convFunc func(sval string) (interface{}, error)) (interface{}, error) {
	flag := f.Lookup(name)
	if flag == nil {
		err := fmt.Errorf("flag accessed but not defined: %s", name)
		return nil, err
	}

	if flag.Value.Type() != ftype {
		err := fmt.Errorf("trying to get %s value of flag of type %s", ftype, flag.Value.Type())
		return nil, err
	}

	sval := flag.Value.String()
	result, err := convFunc(sval)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ArgsLenAtDash will return the length of f.Args at the moment when a -- was
// found during arg parsing. This allows your program to know which args were
// before the -- and which came after.
func (f *FlagSet) ArgsLenAtDash() int {
	return f.argsLenAtDash
}

// MarkDeprecated indicated that a flag is deprecated in your program. It will
// continue to function but will not show up in help or usage messages. Using
// this flag will also print the given usageMessage.
func (f *FlagSet) MarkDeprecated(name string, usageMessage string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	if usageMessage == "" {
		return fmt.Errorf("deprecated message for flag %q must be set", name)
	}
	flag.Deprecated = usageMessage
	flag.Hidden = true
	return nil
}

// MarkShorthandDeprecated will mark the shorthand of a flag deprecated in your
// program. It will continue to function but will not show up in help or usage
// messages. Using this flag will also print the given usageMessage.
func (f *FlagSet) MarkShorthandDeprecated(name string, usageMessage string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	if usageMessage == "" {
		return fmt.Errorf("deprecated message for flag %q must be set", name)
	}
	flag.ShorthandDeprecated = usageMessage
	return nil
}

// MarkHidden sets a flag to 'hidden' in your program. It will continue to
// function but will not show up in help or usage messages.
func (f *FlagSet) MarkHidden(name string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	flag.Hidden = true
	return nil
}

// Lookup returns the Flag structure of the named command-line flag,
// returning nil if none exists.
func Lookup(name string) *Flag {
	return CommandLine.Lookup(name)
}

// ShorthandLookup returns the Flag structure of the short handed flag,
// returning nil if none exists.
func ShorthandLookup(name string) *Flag {
	return CommandLine.ShorthandLookup(name)
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
	normalName := f.normalizeFlagName(name)
	flag, ok := f.formal[normalName]
	if !ok {
		return fmt.Errorf("no such flag -%v", name)
	}

	err := flag.Value.Set(value)
	if err != nil {
		var flagName string
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			flagName = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
		} else {
			flagName = fmt.Sprintf("--%s", flag.Name)
		}
		return fmt.Errorf("invalid argument %q for %q flag: %v", value, flagName, err)
	}

	if !flag.Changed {
		if f.actual == nil {
			f.actual = make(map[NormalizedName]*Flag)
		}
		f.actual[normalName] = flag
		f.orderedActual = append(f.orderedActual, flag)

		flag.Changed = true
	}

	if flag.Deprecated != "" {
		fmt.Fprintf(f.out(), "Flag --%s has been deprecated, %s\n", flag.Name, flag.Deprecated)
	}
	return nil
}

// SetAnnotation allows one to set arbitrary annotations on a flag in the FlagSet.
// This is sometimes used by spf13/cobra programs which want to generate additional
// bash completion information.
func (f *FlagSet) SetAnnotation(name, key string, values []string) error {
	normalName := f.normalizeFlagName(name)
	flag, ok := f.formal[normalName]
	if !ok {
		return fmt.Errorf("no such flag -%v", name)
	}
	if flag.Annotations == nil {
		flag.Annotations = map[string][]string{}
	}
	flag.Annotations[key] = values
	return nil
}

// Changed returns true if the flag was explicitly set during Parse() and false
// otherwise
func (f *FlagSet) Changed(name string) bool {
	flag := f.Lookup(name)
	// If a flag doesn't exist, it wasn't changed....
	if flag == nil {
		return false
	}
	return flag.Changed
}

// Set sets the value of the named command-line flag.
func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined flags in the set.
func (f *FlagSet) PrintDefaults() {
	usages := f.FlagUsages()
	fmt.Fprint(f.out(), usages)
}

// defaultIsZeroValue returns true if the default value for this flag represents
// a zero value.
func (f *Flag) defaultIsZeroValue() bool {
	switch f.Value.(type) {
	case boolFlag:
		return f.DefValue == "false"
	case *durationValue:
		// Beginning in Go 1.7, duration zero values are "0s"
		return f.DefValue == "0" || f.DefValue == "0s"
	case *intValue, *int8Value, *int32Value, *int64Value, *uintValue, *uint8Value, *uint16Value, *uint32Value, *uint64Value, *countValue, *float32Value, *float64Value:
		return f.DefValue == "0"
	case *stringValue:
		return f.DefValue == ""
	case *ipValue, *ipMaskValue, *ipNetValue:
		return f.DefValue == "<nil>"
	case *intSliceValue, *stringSliceValue, *stringArrayValue:
		return f.DefValue == "[]"
	default:
		switch f.Value.String() {
		case "false":
			return true
		case "<nil>":
			return true
		case "":
			return true
		case "0":
			return true
		}
		return false
	}
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for a flag and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the flag's value, or the empty string if the flag is boolean.
func UnquoteUsage(flag *Flag) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = flag.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}

	name = flag.Value.Type()
	switch name {
	case "bool":
		name = ""
	case "float64":
		name = "float"
	case "int64":
		name = "int"
	case "uint64":
		name = "uint"
	case "stringSlice":
		name = "strings"
	case "intSlice":
		name = "ints"
	case "uintSlice":
		name = "uints"
	case "boolSlice":
		name = "bools"
	}

	return
}

// Splits the string `s` on whitespace into an initial substring up to
// `i` runes in length and the remainder. Will go `slop` over `i` if
// that encompasses the entire string (which allows the caller to
// avoid short orphan words on the final line).
func wrapN(i, slop int, s string) (string, string) {
	if i+slop > len(s) {
		return s, ""
	}

	w := strings.LastIndexAny(s[:i], " \t\n")
	if w <= 0 {
		return s, ""
	}
	nlPos := strings.LastIndex(s[:i], "\n")
	if nlPos > 0 && nlPos < w {
		return s[:nlPos], s[nlPos+1:]
	}
	return s[:w], s[w+1:]
}

// Wraps the string `s` to a maximum width `w` with leading indent
// `i`. The first line is not indented (this is assumed to be done by
// caller). Pass `w` == 0 to do no wrapping
func wrap(i, w int, s string) string {
	if w == 0 {
		return strings.Replace(s, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	// space between indent i and end of line width w into which
	// we should wrap the text.
	wrap := w - i

	var r, l string

	// Not enough space for sensible wrapping. Wrap as a block on
	// the next line instead.
	if wrap < 24 {
		i = 16
		wrap = w - i
		r += "\n" + strings.Repeat(" ", i)
	}
	// If still not enough space then don't even try to wrap.
	if wrap < 24 {
		return strings.Replace(s, "\n", r, -1)
	}

	// Try to avoid short orphan words on the final line, by
	// allowing wrapN to go a bit over if that would fit in the
	// remainder of the line.
	slop := 5
	wrap = wrap - slop

	// Handle first line, which is indented by the caller (or the
	// special case above)
	l, s = wrapN(wrap, slop, s)
	r = r + strings.Replace(l, "\n", "\n"+strings.Repeat(" ", i), -1)

	// Now wrap the rest
	for s != "" {
		var t string

		t, s = wrapN(wrap, slop, s)
		r = r + "\n" + strings.Repeat(" ", i) + strings.Replace(t, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	return r

}

// FlagUsagesWrapped returns a string containing the usage information
// for all flags in the FlagSet. Wrapped to `cols` columns (0 for no
// wrapping)
func (f *FlagSet) FlagUsagesWrapped(cols int) string {
	buf := new(bytes.Buffer)

	lines := make([]string, 0, len(f.formal))

	maxlen := 0
	f.VisitAll(func(flag *Flag) {
		if flag.Hidden {
			return
		}

		line := ""
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			line = f