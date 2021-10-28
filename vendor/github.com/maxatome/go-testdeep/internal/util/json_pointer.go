// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"strconv"
	"strings"
)

var jsonPointerEsc = strings.NewReplacer("~0", "~", "~1", "/")

const (
	ErrJSONPointerInvalid         = "invalid JSON pointer"
	ErrJSONPointerKeyNotFound     = "key not found"
	ErrJSONPointerArrayNoIndex    = "array but not an index in JSON pointer"
	ErrJSONPointerArrayOutOfRange = "out of array range"
	ErrJSONPointerArrayBadType    = "not a map nor an array"
)

type JSONPointerError struct {
	Type    string
	Pointer string
}

func (e *JSONPointerError) Error() string {
	if e.Pointer == "" {
		return e.Type
	}
	return e.Type + " @" + e.Pointer
}

// JSONPointer returns the value corresponding to JSON pointer
// "pointer" in "v" as RFC 6901 specifies it. To be searched, "v" has
// to contains map[string]interface{} or []interface{} values. All
// other types fail to be searched.
func JSONPointer(v interface{}, pointer string) (interface{}, error) {
	if !strings.HasPrefix(pointer, "/") {
		if pointer == "" {
			return v, nil
		}
		return nil, &JSONPointerError{Type: ErrJSONPointerInvalid}
	}

	pos := 0
	for _, part := range strings.Split(pointer[1:], "/") {
		pos += 1 + len(part)
		part = jsonPointerEsc.Replace(part)

		switch tv := v.(type) {
		case map[string]interface{}:
			var ok bool
			v, ok = tv[part]
			if !ok {
				return nil, &JSONPointerError{
					Type:    ErrJSONPointerKeyNotFound,
					Pointer: pointer[:pos],
				}
			}

		case []interface{}:
			i, err := strconv.Atoi(part)
			if err != nil || i < 0 {
				return nil, &JSONPointerError{
					Type:    ErrJSONPointerArrayNoIndex,
					Pointer: pointer[:pos],
				}
			}
			if i >= len(tv) {
				return nil, &JSONPointerError{
					Type:    ErrJSONPointerArrayOutOfRange,
					Pointer: pointer[:pos],
				}
			}
			v = tv[i]

		default:
			return nil, &JSONPointerError{
				Type:    ErrJSONPointerArrayBadType,
				Pointer: pointer[:pos],
			}
		}
	}

	return v, nil
}
