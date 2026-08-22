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
	"sigs.k8s.io/external-dns/source/annotations"
)

// Validator validates a raw annotation value string, returning an error describing
// why the value is invalid, or nil if the value is acceptable.
type Validator func(value string) error

// Config describes how a single annotation should be validated, documented, and
// what should happen when validation fails.
type Config struct {
	// Validators are run in order against a raw annotation value; the first error wins.
	Validators []Validator
	// WarnMessage explains the fallback behavior used when the value is invalid
	// and strict mode is not in effect for this annotation.
	WarnMessage string
	// StrictMessage explains the consequence of an invalid value when strict mode is
	// in effect for this annotation. An empty StrictMessage means this annotation has
	// no sensible strict behavior, so ModeStrict has no effect on it.
	StrictMessage string
	// Documentation is a human-readable explanation of the annotation, used for documentation.
	Documentation string
}

// Validate runs all of the Config's Validators against value, returning the first error.
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

// warnText returns the message to log when this annotation's value is invalid and
// strict mode does not apply. It falls back to StrictMessage when WarnMessage is unset,
// for annotations whose warn and strict consequences are identical.
func (c Config) warnText() string {
	if c.WarnMessage != "" {
		return c.WarnMessage
	}
	return c.StrictMessage
}

// AnnotationSpec pairs an annotation key with its Config.
type AnnotationSpec struct {
	// Key is the full annotation key, e.g. AccessKey.
	Key string
	// SupportedSources lists the source names (e.g. "service") that consume this annotation.
	// An empty slice means no restriction is enforced.
	SupportedSources []string
	Config
}

// Mode controls how invalid annotation values are handled globally.
type Mode string

const (
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

// ValidateOneOfFold returns a Validator that accepts only the given values, compared
// case-insensitively. Use this for annotations whose extraction code lowercases the
// value before comparing (e.g. via strings.ToLower), so validation matches what the
// source actually accepts.
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

// ValidateHostnames validates a raw hostname/internal-hostname annotation value: a
// comma-separated list of DNS names. Each name must satisfy both the total-length and
// character-format rules in validation.IsDNS1123Subdomain (RFC 1123, max 253 characters),
// and the per-label length rule (RFC 1035 section 2.3.4, max 63 characters), which
// IsDNS1123Subdomain does not check. A name within the 253-character total but with a
// single label over 63 characters is accepted by Kubernetes' own object-name validation,
// yet rejected by endpoint.NewEndpointWithTTL — validating both here catches that case
// before it reaches source extraction.
func ValidateHostnames(value string) error {
	for _, name := range annotations.SplitHostnameAnnotation(value) {
		if err := validateDNSName(name); err != nil {
			return err
		}
	}
	return nil
}

func validateDNSName(name string) error {
	if errs := validation.IsDNS1123Subdomain(name); len(errs) > 0 {
		return fmt.Errorf("%q is not a valid DNS name: %s", name, strings.Join(errs, "; "))
	}
	if label := endpoint.OverlongDNSLabel(name); label != "" {
		return fmt.Errorf("%q is not a valid DNS name: label %q exceeds %d characters", name, label, validation.DNS1123LabelMaxLength)
	}
	return nil
}
