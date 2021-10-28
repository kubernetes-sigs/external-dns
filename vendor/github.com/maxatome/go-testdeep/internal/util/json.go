// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

var (
	jsonErrPlaceholder   = "<NO_JSON_ERROR!>" // will be overwritten by UnmarshalJSON
	jsonErrCommentPrefix = "<NO_JSON_ERROR!>" // will be overwritten by UnmarshalJSON
	jsonErrorMesgOnce    sync.Once
)

// `{"foo": $bar}`, 8 → `{"foo": "$bar"}`.
func stringifyPlaceholder(buf []byte, dollar int64) ([]byte, error) {
	r, size := utf8.DecodeRune(buf[dollar+1:]) // just after $
	cur := dollar + 1 + int64(size)

	var end int64

	// Numeric placeholder: $1234
	if r >= '0' && r <= '9' {
		for i, c := range buf[cur:] {
			switch c {
			case ' ', '\t', '\r', '\n', ',', '}', ']':
				end = cur + int64(i)
				goto endFound
			default:
				if c < '0' || c > '9' {
					return nil,
						fmt.Errorf(`invalid numeric placeholder at offset %d`, dollar+1)
				}
			}
		}
		end = int64(len(buf))
	endFound:
	} else if shortcut := r == '^'; shortcut || // Operator shortcut, e.g. $^Zero
		unicode.IsLetter(r) || r == '_' { // Named placeholder: $pïpô12
	runes:
		for max := int64(len(buf)); cur < max; cur += int64(size) {
			r, size = utf8.DecodeRune(buf[cur:])
			switch r {
			case '_':
			case ' ', '\t', '\r', '\n', ',', '}', ']':
				break runes
			default:
				if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
					if shortcut {
						return nil,
							fmt.Errorf(`invalid operator shortcut at offset %d`, dollar+1)
					}
					return nil,
						fmt.Errorf(`invalid named placeholder at offset %d`, dollar+1)
				}
			}
		}
		end = cur
	} else {
		return nil, fmt.Errorf(`invalid placeholder at offset %d`, dollar+1)
	}

	// put "" around $éé123, $12345 or $^NotZero
	if cap(buf) == len(buf) {
		// allocate room for 20 extra placeholders
		buf = append(make([]byte, 0, len(buf)+40), buf...)
	}
	buf = append(buf, 0, 0)
	copy(buf[end+2:], buf[end:])
	buf[end+1] = '"'
	copy(buf[dollar+1:], buf[dollar:end])
	buf[dollar] = '"'

	return buf, nil
}

// `{"foo": 123 /* comment */}`, 13 → `{"foo": 123              }`.
func clearComment(buf []byte, slash int64, origErr error) error {
	r, _ := utf8.DecodeRune(buf[slash+1:]) // just after /

	var end int

	switch r {
	case '/': // → // = comment until end of line or buffer
		end = bytes.IndexAny(buf[slash+2:], "\r\n")
		if end < 0 {
			end = len(buf)
		} else {
			end += int(slash) + 2
		}

	case '*': // → /* = comment until */
		end = bytes.Index(buf[slash+2:], []byte(`*/`))
		if end < 0 {
			return fmt.Errorf(`unterminated comment at offset %d`, slash+1)
		}
		end += int(slash) + 2 + 2

	default:
		return origErr
	}

	for i := int(slash); i < end; i++ {
		buf[i] = ' '
	}
	return nil
}

// UnmarshalJSON is a custom json.Unmarshal function allowing to
// handle placeholders not enclosed in strings. It relies on
// json.SyntaxError errors detected before any memory allocation. So
// the performance should not be too bad, avoiding to implement our
// own JSON parser...
func UnmarshalJSON(buf []byte, target interface{}) error {
	jsonErrorMesgOnce.Do(func() {
		var dummy interface{}
		err := json.Unmarshal([]byte(`$x`), &dummy)
		if jerr, ok := err.(*json.SyntaxError); ok {
			jsonErrPlaceholder = jerr.Error()
			jsonErrCommentPrefix = jerr.Error()[:strings.Index(jsonErrPlaceholder, "$")] + "/"
		}
	})

	for {
		err := json.Unmarshal(buf, target)
		if err == nil {
			return nil
		}
		jerr, ok := err.(*json.SyntaxError)
		if ok && jerr.Offset < int64(len(buf)) {
			switch {
			case jerr.Error() == jsonErrPlaceholder:
				buf, err = stringifyPlaceholder(buf, jerr.Offset-1) // "$" pos
				if err == nil {
					continue
				}

			case strings.HasPrefix(jerr.Error(), jsonErrCommentPrefix):
				err = clearComment(buf, jerr.Offset-1, err) // "/" pos
				if err == nil {
					continue
				}
			}
		}
		return err
	}
}
