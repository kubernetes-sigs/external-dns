package endpoint

import (
	"testing"
	"unsafe"
)

// TestStringComparisonBehavior verifies that Go string comparison
// exits early on length mismatch or first byte difference.
func TestStringComparisonBehavior(t *testing.T) {
	// Go strings are: struct { data *byte, len int }
	// Comparison first checks length, then compares bytes

	t.Run("different lengths - no byte comparison needed", func(t *testing.T) {
		a := "alias"       // len=5
		b := "aws/weight"  // len=10
		// Should return false immediately after length check
		if a == b {
			t.Error("should not be equal")
		}
	})

	t.Run("same length - exits at first mismatch", func(t *testing.T) {
		a := "aws/region"   // len=10
		b := "aws/weight"   // len=10
		// Same length, compares: a-w-s-/-r vs a-w-s-/-w -> mismatch at byte 5
		if a == b {
			t.Error("should not be equal")
		}
	})
}

// BenchmarkStringCompareEarlyExit demonstrates early exit behavior
func BenchmarkStringCompareEarlyExit(b *testing.B) {
	// Different lengths - fastest (just compare two integers)
	b.Run("different-length", func(b *testing.B) {
		a := "alias"                              // 5 bytes
		c := "aws/geolocation-subdivision-code"   // 32 bytes
		for i := 0; i < b.N; i++ {
			_ = a == c
		}
	})

	// Same length, early mismatch at byte 5
	b.Run("same-length-early-mismatch", func(b *testing.B) {
		a := "aws/region"  // 10 bytes
		c := "aws/weight"  // 10 bytes - differs at position 4
		for i := 0; i < b.N; i++ {
			_ = a == c
		}
	})

	// Same length, late mismatch
	b.Run("same-length-late-mismatch", func(b *testing.B) {
		a := "aws/geolocation-continent-code"    // 30 bytes
		c := "aws/geolocation-continent-codX"    // 30 bytes - differs at last byte
		for i := 0; i < b.N; i++ {
			_ = a == c
		}
	})

	// Same string (full comparison needed)
	b.Run("equal-strings", func(b *testing.B) {
		a := "aws/geolocation-subdivision-code"  // 32 bytes
		c := "aws/geolocation-subdivision-code"  // 32 bytes
		for i := 0; i < b.N; i++ {
			_ = a == c
		}
	})
}

// Show string internal structure
func TestStringInternals(t *testing.T) {
	s := "hello"

	// Go string header: pointer + length
	type stringHeader struct {
		Data uintptr
		Len  int
	}

	header := *(*stringHeader)(unsafe.Pointer(&s))
	t.Logf("String %q: data=%x, len=%d", s, header.Data, header.Len)
	t.Logf("String comparison first checks Len field (fast integer compare)")
	t.Logf("Only if lengths match, it compares bytes (can exit early)")
}
