package yaml

import (
	"io"
)

func yaml_insert_token(parser *yaml_parser_t, pos int, token *yaml_token_t) {
	//fmt.Println("yaml_insert_token", "pos:", pos, "typ:", token.typ, "head:", parser.tokens_head, "len:", len(parser.tokens))

	// Check if we can move the queue at the beginning of the buffer.
	if parser.tokens_head > 0 && len(parser.tokens) == cap(parser.tokens) {
		if parser.tokens_head != len(parser.tokens) {
			copy(parser.tokens, parser.tokens[parser.tokens_head:])
		}
		parser.tokens = parser.tokens[:len(parser.tokens)-parser.tokens_head]
		parser.tokens_head = 0
	}
	parser.tokens = append(parser.tokens, *token)
	if pos < 0 {
		return
	}
	copy(parser.tokens[parser.tokens_head+pos+1:], parser.tokens[parser.tokens_head+pos:])
	parser.tokens[parser.tokens_head+pos] = *token
}

// Create a new parser object.
func yaml_parser_initialize(parser *yaml_parser_t) bool {
	*parser = yaml_parser_t{
		raw_buffer: make([]byte, 0, input_raw_buffer_size),
		buffer:     make([]byte, 0, input_buffer_size),
	}
	return true
}

// Destroy a parser object.
func yaml_parser_delete(parser *yaml_parser_t) {
	*parser = yaml_parser_t{}
}

// String read handler.
func yaml_string_read_handler(parser *yaml_parser_t, buffer []byte) (n int, err error) {
	if parser.input_pos == len(parser.input) {
		return 0, io.EOF
	}
	n = copy(buffer, parser.input[parser.input_pos:])
	parser.input_pos += n
	return n, nil
}

// Reader read handler.
func yaml_reader_read_handler(parser *yaml_parser_t, buffer []byte) (n int, err error) {
	return parser.input_reader.Read(buffer)
}

// Set a string input.
func yaml_parser_set_input_string(parser *yaml_parser_t, input []byte) {
	if parser.read_handler != nil {
		panic("must set the input source only once")
	}
	parser.read_handler = yaml_string_read_handler
	parser.input = input
	parser.input_pos = 0
}

// Set a file input.
func yaml_parser_set_input_reader(parser *yaml_parser_t, r io.Reader) {
	if parser.read_handler != nil {
		panic("must set the input source only once")
	}
	parser.read_handler = yaml_reader_read_handler
	parser.input_reader = r
}

// Set the source encoding.
func yaml_parser_set_encoding(parser *yaml_parser_t, encoding yaml_encoding_t) {
	if parser.encoding != yaml_ANY_ENCODING {
		panic("must set the encoding only once")
	}
	parser.encoding = encoding
}

// Create a new emitter object.
func yaml_emitter_initialize(emitter *yaml_emitter_t) {
	*emitter = yaml_emitter_t{
		buffer:     make([]byte, output_buffer_size),
		raw_buffer: make([]byte, 0, output_raw_buffer_size),
		states:     make([]yaml_emitter_state_t, 0, initial_stack_size),
		events:     make([]yaml_event_t, 0, initial_queue_size),
	}
}

// Destroy an emitter object.
func yaml_emitter_delete(emitter *yaml_emitter_t) {
	*emitter = yaml_emitter_t{}
}

// String write handler.
func yaml_string_write_handler(emitter *yaml_emitter_t, buffer []byte) error {
	*emitter.output_buffer = append(*emitter.output_buffer, buffer...)
	return nil
}

// yaml_writer_write_handler uses emitter.output_writer to write the
// emitted text.
func yaml_writer_write_handler(emitter *yaml_emitter_t, buffer []byte) error {
	_, err := emitter.output_writer.Write(buffer)
	return err
}

// Set a string output.
func yaml_emitter_set_output_string(emitter *yaml_emitter_t, output_buffer *[]byte) {
	if emitter.write_handler != nil {
		panic("must set the output target only once")
	}
	emitter.write_handler = yaml_string_write_handler
	emitter.output_buffer = output_buffer
}

// Set a file output.
func yaml_emitter_set_output_writer(emitter *yaml_emitter_t, w io.Writer) {
	if emitter.write_handler != nil {
		panic("must set the output target only once")
	}
	emitter.write_handler = yaml_writer_write_handler
	emitter.output_writer = w
}

// Set the output encoding.
func yaml_emitter_set_encoding(emitter *yaml_emitter_t, encoding yaml_encoding_t) {
	if emitter.encoding != yaml_ANY_ENCODING {
		panic("must set the output encoding only once")
	}
	emitter.encoding = encoding
}

// Set the canonical output style.
func yaml_emitter_set_canonical(emitter *yaml_emitter_t, canonical bool) {
	emitter.canonical = canonical
}

//// Set the indentation increment.
func yaml_emitter_set_indent(emitter *yaml_emitter_t, indent int) {
	if indent < 2 || indent > 9 {
		indent = 2
	}
	emitter.best_indent = indent
}

// Set the preferred line width.
func yaml_emitter_set_width(emitter *yaml_emitter_t, width int) {
	if width < 0 {
		width = -1
	}
	emitter.best_width = width
}

// Set if unescaped non-ASCII characters are allowed.
func yaml_emitter_set_unicode(emitter *yaml_emitter_t, unicode bool) {
	emitter.unicode = unicode
}

// Set the preferred line break character.
func yaml_emitter_set_break(emitter *yaml_emitter_t, line_break yaml_break_t) {
	emitter.line_break = line_break
}

///*
// * Destroy a token object.
// */
//
//YAML_DECLARE(void)
//yaml_token_delete(yaml_token_t *token)
//{
//    assert(token);  // Non-NULL token object expected.
//
//    switch (token.type)
//    {
//        case YAML_TAG_DIRECTIVE_TOKEN:
//            yaml_free(token.data.tag_directive.handle);
//            yaml_free(token.data.tag_directive.prefix);
//            break;
//
//        case YAML_ALIAS_TOKEN:
//            yaml_free(token.data.alias.value);
//            break;
//
//        case YAML_ANCHOR_TOKEN:
//            yaml_free(token.data.anchor.value);
//            break;
//
//        case YAML_TAG_TOKEN:
//            yaml_free(token.data.tag.handle);
//            yaml_free(token.data.tag.suffix);
//            break;
//
//        case YAML_SCALAR_TOKEN:
//            yaml_free(token.data.scalar.value);
//            break;
//
//        default:
//            break;
//    }
//
//    memset(token, 0, sizeof(yaml_token_t));
//}
//
///*
// * Check if a string is a valid UTF-8 sequence.
// *
// * Check 'reader.c' for more details on UTF-8 encoding.
// */
//
//static int
//yaml_check_utf8(yaml_char_t *start, size_t length)
//{
//    yaml_char_t *end = start+length;
//    yaml_char_t *pointer = start;
//
//    while (pointer < end) {
//        unsigned char octet;
//        unsigned int width;
//        unsigned int value;
//        size_t k;
//
//        octet = pointer[0];
//        width = (octet & 0x80) == 0x00 ? 1 :
//                (octet & 0xE0) == 0xC0 ? 2 :
//                (octet & 0xF0) == 0xE0 ? 3 :
//                (octet & 0xF8) == 0xF0 ? 4 : 0;
//        value = (octet & 0x80) == 0x00 ? octet & 0x7F :
//                (octet & 0xE0) == 0xC0 ? octet & 0x1F :
//                (octet & 0xF0) == 0xE0 ? octet & 0x0F :
//                (octet & 0xF8) == 0xF0 ? octet & 0x07 : 0;
//        if (!width) return 0;
//        if (pointer+width > end) return 0;
//        for (k = 1; k < width; k ++) {
//            octet = pointer[k];
//            if ((octet & 0xC0) != 0x80) return 0;
//            value = (value << 6) + (octet & 0x3F);
//        }
//        if (!((width ==