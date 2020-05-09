/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Package spew implements a deep pretty printer for Go data structures to aid in
debugging.

A quick overview of the additional features spew provides over the built-in
printing facilities for Go data types are as follows:

	* Pointers are dereferenced and followed
	* Circular data structures are detected and handled properly
	* Custom Stringer/error interfaces are optionally invoked, including
	  on unexported types
	* Custom types which only implement the Stringer/error interfaces via
	  a pointer receiver are optionally invoked when passing non-pointer
	  variables
	* Byte arrays and slices are dumped like the hexdump -C command which
	  includes offsets, byte values in hex, and ASCII output (only when using
	  Dump style)

There are two different approaches spew allows for dumping Go data structures:

	* Dump style which prints with newlines, customizable indentation,
	  and additional debug information such as types and all pointer addresses
	  used to indirect to the final value
	* A custom Formatter interface that integrates cleanly with the standard fmt
	  package and replaces %v, %+v, %#v, and %#+v to provide inline printing
	  similar to the default %v while providing the additional functionality
	  outlined above and passing unsupported format verbs such as %x and %q
	  along to fmt

Quick Start

This section demonstrates how to quickly get started with spew.  See the
sections below for further details on formatting and configuration options.

To dump a variable with full newlines, indentation, type, and pointer
information use Dump, Fdump, or Sdump:
	spew.Dump(myVar1, myVar2, ...)
	spew.Fdump(someWriter, myVar1, myVar2, ...)
	str := spew.Sdump(myVar1, myVar2, ...)

Alternatively, if you would prefer to use format strings with a compacted inline
printing style, use the convenience wrappers Printf, Fprintf, etc with
%v (most compact), %+v (adds pointer addresses), %#v (adds types), or
%#+v (adds types and pointer addresses):
	spew.Printf("myVar1: %v -- myVar2: %+v", myVar1, myVar2)
	spew.Printf("myVar3: %#v -- myVar4: %#+v", myVar3, myVar4)
	spew.Fprintf(someWriter, "myVar1: %v -- myVar2: %+v", myVar1, myVar2)
	spew.Fprintf(someWriter, "myVar3: %#v -- myVar4: %#+v", myVar3, myVar4)

Configuration Options

Configuration of spew is handled by fields in the ConfigState type.  For
convenience, all of the top-level functions use a global state available
via the spew.Config global.

It is also possible to create a ConfigState instance that provides methods
equivalent to the top-level functions.  This allows concurrent configuration
options.  See the ConfigState documentation for more details.

The following configuration options are available:
	* Indent
		String to use for each indentation level for Dump functions.
		It is a single space by default.  A popular alternative is "\t".

	* MaxDepth
		Maximum number of levels to descend into nested data structures.
		There is no limit by default.

	* DisableMethods
		Disables invocation of error and Stringer interface methods.
		Method invocation is enabled by default.

	* DisablePointerMethods
		Disables invocation of error and Stringer interface methods on types
		which only accept pointer receivers from non-pointer variables.
		Pointer method invocation is enabled by default.

	* DisablePointerAddresses
		DisablePointerAddresses specifies whether to disable the printing of
		pointer addresses. This is useful when diffing data structures in tests.

	* DisableCapacities
		DisableCapacities specifies whether to disable the printing of
		capacities for arrays, slices, maps and channels. This is useful when
		diffing data structures in tests.

	* ContinueOnMethod
		Enables recursion into types after invoking error and Stringer interface
		methods. Recursion after method invocation is disabled by default.

	* SortKeys
		Specifies map keys should be sorted before being printed. Use
		this to have a more deterministic, diffable output.  Note that
		only native types (bool, int, uint, floats, uintptr and string)
		and types which implement error or Stringer interfaces are
		supported with other types sorted according to the
		reflect.Value.String() output which guarantees display
		stability.  Natural map order is used by default.

	* SpewKeys
		Specifies that, as a last resort attempt, map keys should be
		spewed to strings and sorted by those strings.  This is only
		considered if SortKeys is true.

Dump Usage

Simply call spew.Dump with a list of variables you want to dump:

	spew.Dump(myVar1, myVar2, ...)

You may also call spew.Fdump if you would prefer to output to an arbitrary
io.Writer.  For example, to dump to standard error:

	spew.Fdump(os.Stderr, myVar1, myVar2, ...)

A third option is t