package config

import (
	"testing"
)

func TestValidateFlags(t *testing.T) {
	cfg := NewConfig()
	cfg.LogFormat = "test"
	if err := cfg.Validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}
	cfg.LogFormat = ""
	if err := cfg.Validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}
	for _, format := range []string{"text", "json"} {
		cfg.LogFormat = format
		if err := cfg.Validate(); err != nil {
			t.Errorf("supported log format: %s should not fail", format)
		}
	}
}
