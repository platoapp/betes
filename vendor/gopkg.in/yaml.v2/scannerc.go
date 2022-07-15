package yaml

import (
	"bytes"
	"fmt"
)

// Introduction
// ************
//
// The following notes assume that you are familiar with the YAML specification
// (http://yaml.org/spec/1.2/spec.html).  We mostly follow it, although in
// some cases we are less restrictive that it requires.
//
// The process of transforming a YAML stream into a sequence of events is
// divided on two steps: Scanning and Parsing.
//
// The Scanner transforms the input stream into a sequence of tokens, while the
// parser transform the sequence of tokens produced by the Scanner into a
// sequence of parsing events.
//
// The Scanner is rather clever and complicated. The Parser, on the contrary,
// is a straightforward implementation of a recursive-descendant parser (or,
// LL(1) parser, as it is usually called).
//
// Actually there are two issues of Scanning that might be called "clever", the
// rest is quite straightforward.  The issues are "block collection start" and
// "simple keys".  Both issues are explained below in details.
//
// Here the Scanning step is explained and implemented.  We start with the list
// of all the tokens produced by the Scanner together with short descriptions.
//
// Now, tokens:
//
//      STREAM-START(encoding)          # The stream start.
//      STREAM-END                      # The stream end.
//      VERSION-DIRECTIVE(major,minor)  # The '%YAML' directive.
//      TAG-DIRECTIVE(handle,prefix)    # The '%TAG' directive.
//      DOCUMENT-START                  # '---'
//      DOCUMENT-END                    # '...'
//      BLOCK-SEQUENCE-START            # Indentation increase denoting a block
//      BLOCK-MAPPING-START             # sequence or a block mapping.
//      BLOCK-END                       # Indentation decrease.
//      FLOW-SEQUENCE-START             # '['
//      FLOW-SEQUENCE-END               # ']'
//      BLOCK-SEQUENCE-START            # '{'
//      BLOCK-SEQUENCE-END              # '}'
//      BLOCK-ENTRY                     # '-'
//      FLOW-ENTRY                      # ','
//      KEY                             # '?' or nothing (simple keys).
//      VALUE                           # ':'
//      ALIAS(anchor)                   # '*anchor'
//      ANCHOR(anchor)                  # '&anchor'
//      TAG(handle,suffix)              # '!handle!suffix'
//      SCALAR(value,style)             # A scalar.
//
// The following two tokens are "virtual" tokens denoting the beginning and the
// end of the stream:
//
//      STREAM-START(encoding)
//      STREAM-END
//
// We pass the information about the input stream encoding with the
// STREAM-START token.
//
// The next two tokens are responsible for tags:
//
//      VERSION-DIRECTIVE(major,minor)
//      TAG-DIRECTIVE(handle,prefix)
//
// Example:
//
//      %YAML   1.1
//      %TAG    !   !foo
//      %TAG    !yaml!  tag:yaml.org,2002:
//      ---
//
// The correspoding sequence of tokens:
//
//      STREAM-START(utf-8)
//      VERSION-DIRECTIVE(1,1)
//      TAG-DIRECTIVE("!","!foo")
//      TAG-DIRECTIVE("!yaml","tag:yaml.org,2002:")
//      DOCUMENT-START
//      STREAM-END
//
// Note that the VERSION-DIRECTIVE and TAG-DIRECTIVE tokens occupy a whole
// line.
//
// The document start and end indicators are represented by:
//
//      DOCUMENT-START
//      DOCUMENT-END
//
// Note that if a YAML stream contains an implicit document (without '---'
// and '...' indicators), no DOCUMENT-START and DOCUMENT-END tokens will be
// produced.
//
// In the following examples, we present whole documents together with the
// produced tokens.
//
//      1. An implicit document:
//
//          'a scalar'
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          SCALAR("a scalar",single-quoted)
//          STREAM-END
//
//      2. An explicit document:
//
//          ---
//          'a scalar'
//          ...
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          DOCUMENT-START
//          SCALAR("a scalar",single-quoted)
//          DOCUMENT-END
//          STREAM-END
//
//      3. Several documents in a stream:
//
//          'a scalar'
//          ---
//          'another scalar'
//          ---
//          'yet another scalar'
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          SCALAR("a scalar",single-quoted)
//          DOCUMENT-START
//          SCALAR("another scalar",single-quoted)
//          DOCUMENT-START
//          SCALAR("yet another scalar",single-quoted)
//          STREAM-END
//
// We have already introduced the SCALAR token above.  The following tokens are
// used to describe aliases, anchors, tag, and scalars:
//
//      ALIAS(anchor)
//      ANCHOR(anchor)
//      TAG(handle,suffix)
//      SCALAR(value,style)
//
// The following series of examples illustrate the usage of these tokens:
//
//      1. A recursive sequence:
//
//          &A [ *A ]
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          ANCHOR("A")
//          FLOW-SEQUENCE-START
//          ALIAS("A")
//          FLOW-SEQUENCE-END
//          STREAM-END
//
//      2. A tagged scalar:
//
//          !!float "3.14"  # A good approximation.
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          TAG("!!","float")
//          SCALAR("3.14",double-quoted)
//          STREAM-END
//
//      3. Various scalar styles:
//
//          --- # Implicit empty plain scalars do not produce tokens.
//          --- a plain scalar
//          --- 'a single-quoted scalar'
//          --- "a double-quoted scalar"
//          --- |-
//            a literal scalar
//          --- >-
//            a folded
//            scalar
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          DOCUMENT-START
//          DOCUMENT-START
//          SCALAR("a plain scalar",plain)
//          DOCUMENT-START
//          SCALAR("a single-quoted scalar",single-quoted)
//          DOCUMENT-START
//          SCALAR("a double-quoted scalar",double-quoted)
//          DOCUMENT-START
//          SCALAR("a literal scalar",literal)
//          DOCUMENT-START
//          SCALAR("a folded scalar",folded)
//          STREAM-END
//
// Now it's time to review collection-related tokens. We will start with
// flow collections:
//
//      FLOW-SEQUENCE-START
//      FLOW-SEQUENCE-END
//      FLOW-MAPPING-START
//      FLOW-MAPPING-END
//      FLOW-ENTRY
//      KEY
//      VALUE
//
// The tokens FLOW-SEQUENCE-START, FLOW-SEQUENCE-END, FLOW-MAPPING-START, and
// FLOW-MAPPING-END represent the indicators '[', ']', '{', and '}'
// correspondingly.  FLOW-ENTRY represent the ',' indicator.  Finally the
// indicators '?' and ':', which are used for denoting mapping keys and values,
// are represented by the KEY and VALUE tokens.
//
// The following examples show flow collections:
//
//      1. A flow sequence:
//
//          [item 1, item 2, item 3]
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          FLOW-SEQUENCE-START
//          SCALAR("item 1",plain)
//          FLOW-ENTRY
//          SCALAR("item 2",plain)
//          FLOW-ENTRY
//          SCALAR("item 3",plain)
//          FLOW-SEQUENCE-END
//          STREAM-END
//
//      2. A flow mapping:
//
//          {
//              a simple key: a value,  # Note that the KEY token is produced.
//              ? a complex key: another value,
//          }
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          FLOW-MAPPING-START
//          KEY
//          SCALAR("a simple key",plain)
//          VALUE
//          SCALAR("a value",plain)
//          FLOW-ENTRY
//          KEY
//          SCALAR("a complex key",plain)
//          VALUE
//          SCALAR("another value",plain)
//          FLOW-ENTRY
//          FLOW-MAPPING-END
//          STREAM-END
//
// A simple key is a key which is not denoted by the '?' indicator.  Note that
// the Scanner still produce the KEY token whenever it encounters a simple key.
//
// For scanning block collections, the following tokens are used (note that we
// repeat KEY and VALUE here):
//
//      BLOCK-SEQUENCE-START
//      BLOCK-MAPPING-START
//      BLOCK-END
//      BLOCK-ENTRY
//      KEY
//      VALUE
//
// The tokens BLOCK-SEQUENCE-START and BLOCK-MAPPING-START denote indentation
// increase that precedes a block collection (cf. the INDENT token in Python).
// The token BLOCK-END denote indentation decrease that ends a block collection
// (cf. the DEDENT token in Python).  However YAML has some syntax pecularities
// that makes detections of these tokens more complex.
//
// The tokens BLOCK-ENTRY, KEY, and VALUE are used to represent the indicators
// '-', '?', and ':' correspondingly.
//
// The following examples show how the tokens BLOCK-SEQUENCE-START,
// BLOCK-MAPPING-START, and BLOCK-END are emitted by the Scanner:
//
//      1. Block sequences:
//
//          - item 1
//          - item 2
//          -
//            - item 3.1
//            - item 3.2
//          -
//            key 1: value 1
//            key 2: value 2
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          SCALAR("item 1",plain)
//          BLOCK-ENTRY
//          SCALAR("item 2",plain)
//          BLOCK-ENTRY
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          SCALAR("item 3.1",plain)
//          BLOCK-ENTRY
//          SCALAR("item 3.2",plain)
//          BLOCK-END
//          BLOCK-ENTRY
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("key 1",plain)
//          VALUE
//          SCALAR("value 1",plain)
//          KEY
//          SCALAR("key 2",plain)
//          VALUE
//          SCALAR("value 2",plain)
//          BLOCK-END
//          BLOCK-END
//          STREAM-END
//
//      2. Block mappings:
//
//          a simple key: a value   # The KEY token is produced here.
//          ? a complex key
//          : another value
//          a mapping:
//            key 1: value 1
//            key 2: value 2
//          a sequence:
//            - item 1
//            - item 2
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("a simple key",plain)
//          VALUE
//          SCALAR("a value",plain)
//          KEY
//          SCALAR("a complex key",plain)
//          VALUE
//          SCALAR("another value",plain)
//          KEY
//          SCALAR("a mapping",plain)
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("key 1",plain)
//          VALUE
//          SCALAR("value 1",plain)
//          KEY
//          SCALAR("key 2",plain)
//          VALUE
//          SCALAR("value 2",plain)
//          BLOCK-END
//          KEY
//          SCALAR("a sequence",plain)
//          VALUE
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          SCALAR("item 1",plain)
//          BLOCK-ENTRY
//          SCALAR("item 2",plain)
//          BLOCK-END
//          BLOCK-END
//          STREAM-END
//
// YAML does not always require to start a new block collection from a new
// line.  If the current line contains only '-', '?', and ':' indicators, a new
// block collection may start at the current line.  The following examples
// illustrate this case:
//
//      1. Collections in a sequence:
//
//          - - item 1
//            - item 2
//          - key 1: value 1
//            key 2: value 2
//          - ? complex key
//            : complex value
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          SCALAR("item 1",plain)
//          BLOCK-ENTRY
//          SCALAR("item 2",plain)
//          BLOCK-END
//          BLOCK-ENTRY
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("key 1",plain)
//          VALUE
//          SCALAR("value 1",plain)
//          KEY
//          SCALAR("key 2",plain)
//          VALUE
//          SCALAR("value 2",plain)
//          BLOCK-END
//          BLOCK-ENTRY
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("complex key")
//          VALUE
//          SCALAR("complex value")
//          BLOCK-END
//          BLOCK-END
//          STREAM-END
//
//      2. Collections in a mapping:
//
//          ? a sequence
//          : - item 1
//            - item 2
//          ? a mapping
//          : key 1: value 1
//            key 2: value 2
//
//      Tokens:
//
//          STREAM-START(utf-8)
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("a sequence",plain)
//          VALUE
//          BLOCK-SEQUENCE-START
//          BLOCK-ENTRY
//          SCALAR("item 1",plain)
//          BLOCK-ENTRY
//          SCALAR("item 2",plain)
//          BLOCK-END
//          KEY
//          SCALAR("a mapping",plain)
//          VALUE
//          BLOCK-MAPPING-START
//          KEY
//          SCALAR("key 1",plain)
//          VALUE
//          SCALAR("value 1",plain)
//          KEY
//          SCALAR("key 2",plain)
//          VALUE
//          SCALAR("value 2",plain)
//          BLOCK-END
//          BLOCK-END
//          STREAM-END
//
// YAML also permits non-indented sequences if they are included into a block
// mapping.  In this case, the token BLOCK-SEQUENCE-START is not produced:
//
//      key:
//      - item 1    # BLOCK-SEQUENCE-START is NOT produced here.
//      - item 2
//
// Tokens:
//
//      STREAM-START(utf-8)
//      BLOCK-MAPPING-START
//      KEY
//      SCALAR("key",plain)
//      VALUE
//      BLOCK-ENTRY
//      SCALAR("