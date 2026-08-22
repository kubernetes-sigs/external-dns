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
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name       string
		validators []Validator
		value      string
		wantErr    bool
	}{
		{
			name:  "no validators is always valid",
			value: "anything",
		},
		{
			name:       "single validator passes",
			validators: []Validator{ValidateOneOf("a", "b")},
			value:      "a",
		},
		{
			name:       "single validator fails",
			validators: []Validator{ValidateOneOf("a", "b")},
			value:      "c",
			wantErr:    true,
		},
		{
			name: "later validator runs only if earlier ones pass",
			validators: []Validator{
				func(string) error { return nil },
				ValidateOneOf("a", "b"),
			},
			value:   "c",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{Validators: tt.validators}
			err := cfg.Validate(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestConfig_Validate_FirstErrorWins(t *testing.T) {
	cfg := Config{
		Validators: []Validator{
			func(string) error { return errors.New("first error") },
			func(string) error { return errors.New("second error") },
		},
	}
	err := cfg.Validate("value")
	assert.EqualError(t, err, "first error")
}

func TestConfig_IsStrict(t *testing.T) {
	tests := []struct {
		name          string
		mode          Mode
		strictMessage string
		want          bool
	}{
		{name: "warn mode is never strict", mode: ModeWarn, strictMessage: "excluded", want: false},
		{name: "strict mode with StrictMessage is strict", mode: ModeStrict, strictMessage: "excluded", want: true},
		{name: "strict mode without StrictMessage is not strict", mode: ModeStrict, strictMessage: "", want: false},
		{name: "empty mode is never strict", mode: "", strictMessage: "excluded", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{StrictMessage: tt.strictMessage}
			assert.Equal(t, tt.want, cfg.IsStrict(tt.mode))
		})
	}
}

func TestValidateOneOf(t *testing.T) {
	validator := ValidateOneOf("public", "private")

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "first accepted value", value: "public"},
		{name: "second accepted value", value: "private"},
		{name: "unaccepted value", value: "pubic", wantErr: true},
		{name: "empty value", value: "", wantErr: true},
		{name: "case sensitive mismatch", value: "Public", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "must be one of")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateOneOf_NoAcceptedValues(t *testing.T) {
	validator := ValidateOneOf()
	assert.Error(t, validator("anything"))
}

func TestValidateTTL(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid integer seconds", value: "600"},
		{name: "valid duration", value: "10m"},
		{name: "not a number or duration", value: "abc", wantErr: true},
		{name: "negative", value: "-1", wantErr: true},
		{name: "zero is below minimum", value: "0", wantErr: true},
		{name: "over maximum", value: "99999999999", wantErr: true},
		{name: "empty value", value: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTTL(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateHostNames(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "single valid hostname", value: "foo.example.com"},
		{name: "comma-separated valid hostnames", value: "foo.example.com,bar.example.com"},
		{name: "invalid character", value: "foo_bar!.example.com", wantErr: true},
		{name: "one invalid entry in an otherwise valid list", value: "foo.example.com,bad_host,bar.example.com", wantErr: true},
		{
			name:    "label over 63 characters within a 253-character total",
			value:   strings.Repeat("a", 64) + ".example.com",
			wantErr: true,
		},
		{
			name:    "label of exactly 63 characters is valid",
			value:   strings.Repeat("a", 63) + ".example.com",
			wantErr: false,
		},
		{
			name:    "total length over 253 characters",
			value:   strings.Repeat("a", 250) + ".com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateHostnames(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
