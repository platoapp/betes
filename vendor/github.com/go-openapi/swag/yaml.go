// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swag

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"

	yaml "gopkg.in/yaml.v2"
)

// YAMLMatcher matches yaml
func YAMLMatcher(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".yaml" || ext == ".yml"
}

// YAMLToJSON converts YAML unmarshaled data into json compatible data
func YAMLToJSON(data interface{}) (json.RawMessage, error) {
	jm, err := transformData(data)
	if err != nil {
		return nil, err
	}
	b, err := WriteJSON(jm)
	return json.RawMessage(b), err
}

// BytesToYAMLDoc converts a byte slice into a YAML document
func BytesToYAMLDoc(data []byte) (interface{}, error) {
	var canary map[interface{}]interface{} // validate this is an object and not a different type
	if err := yaml.Unmarshal(data, &canary); err != nil {
		return nil, err
	}

	var document yaml.MapSlice // preserve order that is present in the document
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	return document, nil
}

type JSONMapSlice []JSONMapItem

func (s JSONMapSlice) MarshalJSON() ([]byte, error) {
	w := &jwriter.Writer{Flags: jwriter.NilMapAsEmpty | jwriter.NilSliceAsEmpty}
	s.MarshalEasyJSON(w)
	return w.BuildBytes()
}

func (s JSONMapSlice) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	ln := len(s)
	last := ln - 1
	for i := 0; i < ln; i++ {
		s[i].MarshalEasyJSON(w)
		if i != last { // last item
			w.RawByte(',')
		}
	}

	w.RawByte('}')
}

func (s *JSONMapSlice) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	s.UnmarshalEasyJSON(&l)
	return l.Error()
}
func (s *JSONMapSlice) UnmarshalEasyJSON(in *jlexer