package yaml

import (
	"encoding"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type encoder struct {
	emitter yaml_emitter_t
	event   yaml_event_t
	out     []byte
	flow    bool
	// doneInit holds whether the initial stream_start_event has been
	// emitted.
	doneInit bool
}

func newEncoder() *encoder {
	e := &encoder{}
	yaml_emitter_initialize(&e.emitter)
	yaml_emitter_set_output_string(&e.emitter, &e.out)
	yaml_emitter_set_unicode(&e.emitter, true)
	return e
}

func newEncoderWithWriter(w io.Writer) *encoder {
	e := &encoder{}
	yaml_emitter_initialize(&e.emitter)
	yaml_emitter_set_output_writer(&e.emitter, w)
	yaml_emitter_set_unicode(&e.emitter, true)
	return e
}

func (e *encoder) init() {
	if e.doneInit {
		return
	}
	yaml_stream_start_event_initialize(&e.event, yaml_UTF8_ENCODING)
	e.emit()
	e.doneInit = true
}

func (e *encoder) finish() {
	e.emitter.open_ended = false
	yaml_stream_end_event_initialize(&e.event)
	e.emit()
}

func (e *encoder) destroy() {
	yaml_emitter_delete(&e.emitter)
}

func (e *encoder) emit() {
	// This will internally delete the e.event value.
	e.must(yaml_emitter_emit(&e.emitter, &e.event))
}

func (e *encoder) must(ok bool) {
	if !ok {
		msg := e.emitter.problem
		if msg == "" {
			msg = "unknown problem generating YAML content"
		}
		failf("%s", msg)
	}
}

func (e *encoder) marshalDoc(tag string, in reflect.Value) {
	e.init()
	yaml_document_start_event_initialize(&e.event, nil, nil, true)
	e.emit()
	e.marshal(tag, in)
	yaml_document_end_event_initialize(&e.event, true)
	e.emit()
}

func (e *encoder) marshal(tag string, in reflect.Value) {
	if !in.IsValid() || in.Kind() == reflect.Ptr && in.IsNil() {
		e.nilv()
		return
	}
	iface := in.Interface()
	switch m := iface.(type) {
	case time.Time, *time.Time:
		// Although time.Time implements TextMarshaler,
		// we don't want to treat it as a string for YAML
		// purposes because YAML has special support for
		// timestamps.
	case Marshaler:
		v, err := m.MarshalYAML()
		if err != nil {
			fail(err)
		}
		if v == nil {
			e.nilv()
			return
		}
		in = reflect.ValueOf(v)
	case encoding.TextMarshaler:
		text, err := m.MarshalText()
		if err != nil {
			fail(err)
		}
		in = reflect.ValueOf(string(text))
	case nil:
		e.nilv()
		return
	}
	switch in.Kind() {
	case reflect.Interface:
		e.marshal(tag, in.Elem())
	case reflect.Map:
		e.mapv(tag, in)
	case reflect.Ptr:
		if in.Type() == ptrTimeType {
			e.timev(tag, in.Elem())
		} else {
			e.marshal(tag, in.Elem())
		}
	case reflect.Struct:
		if in.Type() == timeType {
			e.timev(tag, in)
		} else {
			e.structv(tag, in)
		}
	case reflect.Slice, reflect.Array:
		if in.Type().Elem() == mapItemType {
			e.itemsv(tag, in)
		} else {
			e.slicev(tag, in)
		}
	case reflect.String:
		e.stringv(tag, in)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if in.Type() == durationType {
			e.stringv(tag, reflect.ValueOf(iface.(time.Duration).String()))
		} else {
			e.intv(tag, in)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		e.uintv(tag, in)
	case reflect.Float32, reflect.Float64:
		e.floatv(tag, in)
	case reflect.Bool:
		e.boolv(tag, in)
	default:
		panic("cannot marshal type: " + in.Type().String())
	}
}

func (e *encoder) mapv(tag string, in reflect.Value) {
	e.mappingv(tag, func() {
		keys := keyList(in.MapKeys())
		sort.Sort(keys)
		for _, k := range keys {
			e.marshal("", k)
			e.marshal("", in.MapIndex