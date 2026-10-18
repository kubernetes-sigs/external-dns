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

package schema

import (
	"fmt"
	"slices"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/idna"
	"sigs.k8s.io/external-dns/source/annotations"
)

// Validator validates a raw annotation value string, returning an error describing
// why the value is invalid, or nil if the value is acceptable.
type Validator func(value string) error

// Config describes how a single annotation is validated and what happens when it isn't.
type Config struct {
	// Validators run in order; the first error wins.
	Validators  []Validator
	WarnMessage string
	// StrictMessage: empty means ModeStrict has no effect on this annotation.
	StrictMessage string
	Documentation string
}

func (c Config) Validate(value string) error {
	for _, validator := range c.Validators {
		if err := validator(value); err != nil {
			return err
		}
	}
	return nil
}

// IsStrict reports whether an invalid value for this annotation should be treated as
// an error rather than falling back to a default, under the given mode.
func (c Config) IsStrict(mode Mode) bool {
	return mode == ModeStrict && c.StrictMessage != ""
}

func (c Config) warnText() string {
	if c.WarnMessage != "" {
		return c.WarnMessage
	}
	return c.StrictMessage
}

// AnnotationSpec pairs an annotation's key suffix with its Config.
// KeySuffix stays prefix-less so Key() resolves against the current --annotation-prefix,
// not whatever was in effect when Registry was built.
type AnnotationSpec struct {
	KeySuffix        string
	SupportedSources []string
	Config
}

// Key returns the annotation's full key under the currently configured --annotation-prefix.
func (a AnnotationSpec) Key() string {
	return annotations.AnnotationKeyPrefix + a.KeySuffix
}

// Mode controls how invalid annotation values are handled globally.
type Mode string

const (
	// ModeOff disables schema validation; an empty Mode is treated the same.
	ModeOff Mode = "off"
	// ModeWarn logs a warning and falls back to a default value (the historical behavior).
	ModeWarn Mode = "warn"
	// ModeStrict treats an invalid value as an error, for annotations that opt in via StrictMessage.
	ModeStrict Mode = "strict"
)

// ValidateOneOf returns a Validator that accepts only the given values.
func ValidateOneOf(values ...string) Validator {
	return func(value string) error {
		if slices.Contains(values, value) {
			return nil
		}
		return fmt.Errorf("must be one of %v", values)
	}
}

// ValidateOneOfFold is like ValidateOneOf but compares case-insensitively.
func ValidateOneOfFold(values ...string) Validator {
	return func(value string) error {
		for _, v := range values {
			if strings.EqualFold(value, v) {
				return nil
			}
		}
		return fmt.Errorf("must be one of %v (case-insensitive)", values)
	}
}

// ValidateNonEmpty rejects an empty string.
func ValidateNonEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("must not be empty")
	}
	return nil
}

// ValidateTTL validates a raw ttl annotation value using the same parsing and bounds
// checks as annotations.TTLFromAnnotations, without extracting or logging anything itself.
func ValidateTTL(value string) error {
	ttlValue, err := annotations.ParseTTL(value)
	if err != nil {
		return err
	}
	if ttlValue < annotations.TTLMinimum || ttlValue > annotations.TTLMaximum {
		return fmt.Errorf("TTL value %d must be between [%d, %d]", ttlValue, annotations.TTLMinimum, annotations.TTLMaximum)
	}
	return nil
}

// ValidateHostnames validates a comma-separated list of DNS names against what the sources actually accept, not Kubernetes' stricter object-name rules.
func ValidateHostnames(value string) error {
	for _, name := range annotations.SplitHostnameAnnotation(value) {
		if name == "" {
			continue
		}
		if err := validateDNSName(name); err != nil {
			return err
		}
	}
	return nil
}

func validateDNSName(name string) error {
	// Strip the wildcard prefix before IDNA conversion ("*" is not a valid IDNA label),
	// then punycode-convert so DNS1123 (which is ASCII-only) accepts IDN hostnames.
	ascii, err := idna.Profile.ToASCII(strings.TrimPrefix(name, "*."))
	if err != nil {
		return fmt.Errorf("%q is not a valid DNS name: %w", name, err)
	}
	normalized := strings.TrimSuffix(strings.ToLower(ascii), ".")
	if errs := validation.IsDNS1123SubdomainWithUnderscore(normalized); len(errs) > 0 {
		return fmt.Errorf("%q is not a valid DNS name: %s", name, strings.Join(errs, "; "))
	}
	if label := endpoint.OverlongDNSLabel(normalized); label != "" {
		return fmt.Errorf("%q is not a valid DNS name: label %q exceeds %d characters", name, label, validation.DNS1123LabelMaxLength)
	}
	return nil
}
