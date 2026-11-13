/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// CAATarget is a parsed CAA target (RFC 8659): <flags> <tag> "<value>".
type CAATarget struct {
	flags uint8
	tag   string
	value string
}

// NewCAARecord parses a CAA target; the value is the rest of the line, quoted or not, and must not contain '"' or '\'.
func NewCAARecord(target string) (*CAATarget, error) {
	rest := strings.TrimSpace(target)
	if len(strings.Fields(rest)) < 3 {
		return nil, fmt.Errorf("invalid CAA target %q: expected <flags> <tag> <value>", target)
	}

	flagsField, rest := cutCAAField(rest)
	flags, err := strconv.ParseUint(flagsField, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("invalid CAA target %q: flags %q must be an integer between 0 and 255", target, flagsField)
	}

	tag, rest := cutCAAField(rest)
	if !isCAATag(tag) {
		return nil, fmt.Errorf("invalid CAA target %q: tag %q must be 1 to 255 ASCII letters or digits", target, tag)
	}

	value := rest
	if strings.HasPrefix(value, `"`) {
		if len(value) < 2 || !strings.HasSuffix(value, `"`) {
			return nil, fmt.Errorf("invalid CAA target %q: value has an unterminated quote", target)
		}
		value = value[1 : len(value)-1]
	}
	if strings.ContainsAny(value, `"\`) {
		return nil, fmt.Errorf(`invalid CAA target %q: value must not contain '"' or '\'`, target)
	}

	return &CAATarget{flags: uint8(flags), tag: tag, value: value}, nil
}

// NewCAATargetFromFields builds a CAATarget for providers that return CAA as structured fields.
func NewCAATargetFromFields(flags uint8, tag, value string) *CAATarget {
	return &CAATarget{flags: flags, tag: tag, value: value}
}

func cutCAAField(s string) (string, string) {
	i := strings.IndexFunc(s, unicode.IsSpace)
	return s[:i], strings.TrimLeftFunc(s[i:], unicode.IsSpace)
}

// RFC 8659 section 4.1: tags are ASCII letters and digits, and the tag length is one octet.
func isCAATag(tag string) bool {
	if len(tag) == 0 || len(tag) > 255 {
		return false
	}
	for _, r := range tag {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// GetFlags returns the CAA flags field (RFC 8659 section 4.1).
func (c *CAATarget) GetFlags() uint8 { return c.flags }

// GetTag returns the CAA property tag (RFC 8659 section 4.1).
func (c *CAATarget) GetTag() string { return c.tag }

// GetValue returns the CAA property value without surrounding quotes.
func (c *CAATarget) GetValue() string { return c.value }

// String renders the target in canonical presentation format with a quoted value.
func (c *CAATarget) String() string {
	return fmt.Sprintf(`%d %s "%s"`, c.flags, c.tag, c.value)
}

func (t Targets) validateCAARecord() error {
	for _, target := range t {
		if _, err := NewCAARecord(target); err != nil {
			return err
		}
	}
	return nil
}
