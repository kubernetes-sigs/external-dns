package main

import (
	"testing"
)

func TestValidateFlags(t *testing.T) {
	cfg := newConfig()
	cfg.logFormat = "test"
	if err := cfg.validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.logFormat)
	}
	cfg.logFormat = ""
	if err := cfg.validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.logFormat)
	}
	for _, format := range []string{"text", "json"} {
		cfg.logFormat = format
		if err := cfg.validate(); err != nil {
			t.Errorf("supported log format: %s should not fail", format)
		}
	}
}
