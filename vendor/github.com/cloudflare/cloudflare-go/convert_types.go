// File contains helper methods for accepting variants (pointers, values,
// slices, etc) of a particular type and returning them in another. A common use
// is pointer to values and back.
//
// _Most_ follow the convention of (where <type> is a Golang type such as Bool):
//
// <type>Ptr: Accepts a value and returns a pointer.
// <type>: Accepts a pointer and returns a value.
// <type>PtrSlice: Accepts a slice of values and returns a slice of pointers.
// <type>Slice: Accepts a slice of pointers and returns a slice of values.
// <type>PtrMap: Accepts a string map of values into a string map of pointers.
// <type>Map: Accepts a string map of pointers into a string map of values.
//
// Not all Golang types are covered here, only those that are commonly used.
package cloudflare

import (
	"reflect"
	"time"
)

// AnyPtr is a helper routine that allocates a new interface value
// to store v and returns a pointer to it.
//
// 	// Usage: var _ *Type = AnyPtr(Type(value) | value).(*Type)
//
// 	var _ *bool = AnyPtr(true).(*bool)
// 	var _ *byte = AnyPtr(byte(1)).(*byte)
// 	var _ *complex64 = AnyPtr(complex64(1.1)).(*complex64)
// 	var _ *complex128 = AnyPtr(complex128(1.1)).(*complex128)
// 	var _ *float32 = AnyPtr(float32(1.1)).(*float32)
// 	var _ *float64 = AnyPtr(float64(1.1)).(*float64)
// 	var _ *int = AnyPtr(int(1)).(*int)
// 	var _ *int8 = AnyPtr(int8(8)).(*int8)
// 	var _ *int16 = AnyPtr(int16(16)).(*int16)
// 	var _ *int32 = AnyPtr(int32(32)).(*int32)
// 	var _ *int64 = AnyPtr(int64(64)).(*int64)
// 	var _ *rune = AnyPtr(rune(1)).(*rune)
// 	var _ *string = AnyPtr("ptr").(*string)
// 	var _ *uint = AnyPtr(uint(1)).(*uint)
// 	var _ *uint8 = AnyPtr(uint8(8)).(*uint8)
// 	var _ *uint16 = AnyPtr(uint16(16)).(*uint16)
// 	var _ *uint32 = AnyPtr(uint32(32)).(*uint32)
// 	var _ *uint64 = AnyPtr(uint64(64)).(*uint64)
func AnyPtr(v interface{}) interface{} {
	r := reflect.New(reflect.TypeOf(v))
	reflect.ValueOf(r.Interface()).Elem().Set(reflect.ValueOf(v))
	return r.Interface()
}

// BytePtr is a helper routine that allocates a new byte value to store v and
// returns a pointer to it.
func BytePtr(v byte) *byte { return &v }

// Complex64Ptr is a helper routine that allocates a new complex64 value to
// store v and returns a pointer to it.
func Complex64Ptr(v complex64) *complex64 { return &v }

// Complex128Ptr is a helper routine that allocates a new complex128 value
// to store v and returns a pointer to it.
func Complex128Ptr(v complex128) *complex128 { return &v }

// RunePtr is a helper routine that allocates a new rune value to store v
// and returns a pointer to it.
func RunePtr(v rune) *rune { return &v }

// TimePtr is a helper routine that allocates a new time.Time value
// to store v and returns a pointer to it.
func TimePtr(v time.Time) *time.Time { return &v }

// DurationPtr is a helper routine that allocates a new time.Duration value
// to store v and returns a pointer to it.
func DurationPtr(v time.Duration) *time.Duration { return &v }

// BoolPtr is a helper routine that allocates a new bool value to store v and
// returns a pointer to it.
func BoolPtr(v bool) *bool { return &v }

// Bool is a helper routine that accepts a bool pointer and returns a value
// to it.
func Bool(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolPtrSlice converts a slice of bool values into a slice of bool pointers.
func BoolPtrSlice(src []bool) []*bool {
	dst := make([]*bool, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// BoolSlice converts a slice of bool pointers into a slice of bool values.
func BoolSlice(src []*bool) []bool {
	dst := make([]bool, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// BoolPtrMap converts a string map of bool values into a string map of bool
// pointers.
func BoolPtrMap(src map[string]bool) map[string]*bool {
	dst := make(map[string]*bool)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// BoolMap converts a string map of bool pointers into a string map of bool
// values.
func BoolMap(src map[string]*bool) map[string]bool {
	dst := make(map[string]bool)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Byte is a helper routine that accepts a byte pointer and returns a
// value to it.
func Byte(v *byte) byte {
	if v != nil {
		return *v
	}
	return byte(0)
}

// Complex64 is a helper routine that accepts a complex64 pointer and
// returns a value to it.
func Complex64(v *complex64) complex64 {
	if v != nil {
		return *v
	}
	return 0
}

// Complex128 is a helper routine that accepts a complex128 pointer and
// returns a value to it.
func Complex128(v *complex128) complex128 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Ptr is a helper routine that allocates a new float32 value to store v
// and returns a pointer to it.
func Float32Ptr(v float32) *float32 { return &v }

// Float32 is a helper routine that accepts a float32 pointer and returns a
// value to it.
func Float32(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32PtrSlice converts a slice of float32 values into a slice of float32
// pointers.
func Float32PtrSlice(src []float32) []*float32 {
	dst := make([]*float32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Float32Slice converts a slice of float32 pointers into a slice of
// float32 values.
func Float32Slice(src []*float32) []float32 {
	dst := make([]float32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float32PtrMap converts a string map of float32 values into a string map of
// float32 pointers.
func Float32PtrMap(src map[string]float32) map[string]*float32 {
	dst := make(map[string]*float32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Float32Map converts a string map of float32 pointers into a string
// map of float32 values.
func Float32Map(src map[string]*float32) map[string]float32 {
	dst := make(map[string]float32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Float64Ptr is a helper routine that allocates a new float64 value to store v
// and returns a pointer to it.
func Float64Ptr(v float64) *float64 { return &v }

// Float64 is a helper routine that accepts a float64 pointer and returns a
// value to it.
func Float64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64PtrSlice converts a slice of float64 values into a slice of float64
// pointers.
func Float64PtrSlice(src []float64) []*float64 {
	dst := make([]*float64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Float64Slice converts a slice of float64 pointers into a slice of
// float64 values.
func Float64Slice(src []*float64) []float64 {
	dst := make([]float64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float64PtrMap converts a string map of float64 values into a string map of
// float64 pointers.
func Float64PtrMap(src map[string]float64) map[string]*float64 {
	dst := make(map[string]*float64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Float64Map converts a string map of float64 pointers into a string
// map of float64 values.
func Float64Map(src map[string]*float64) map[string]float64 {
	dst := make(map[string]float64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// IntPtr is a helper routine that allocates a new int value to store v and
// returns a pointer to it.
func IntPtr(v int) *int { return &v }

// Int is a helper routine that accepts a int pointer and returns a value
// to it.
func Int(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntPtrSlice converts a slice of int values into a slice of int pointers.
func IntPtrSlice(src []int) []*int {
	dst := make([]*int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// IntSlice converts a slice of int pointers into a slice of int values.
func IntSlice(src []*int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// IntPtrMap converts a string map of int values into a string map of int
// pointers.
func IntPtrMap(src map[string]int) map[string]*int {
	dst := make(map[string]*int)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// IntMap converts a string map of int pointers into a string map of int
// values.
func IntMap(src map[string]*int) map[string]int {
	dst := make(map[string]int)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int8Ptr is a helper routine that allocates a new int8 value to store v and
// returns a pointer to it.
func Int8Ptr(v int8) *int8 { return &v }

// Int8 is a helper routine that accepts a int8 pointer and returns a value
// to it.
func Int8(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int8PtrSlice converts a slice of int8 values into a slice of int8 pointers.
func Int8PtrSlice(src []int8) []*int8 {
	dst := make([]*int8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int8Slice converts a slice of int8 pointers into a slice of int8 values.
func Int8Slice(src []*int8) []int8 {
	dst := make([]int8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int8PtrMap converts a string map of int8 values into a string map of int8
// pointers.
func Int8PtrMap(src map[string]int8) map[string]*int8 {
	dst := make(map[string]*int8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int8Map converts a string map of int8 pointers into a string map of int8
// values.
func Int8Map(src map[string]*int8) map[string]int8 {
	dst := make(map[string]int8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int16Ptr is a helper routine that allocates a new int16 value to store v
// and returns a pointer to it.
func Int16Ptr(v int16) *int16 { return &v }

// Int16 is a helper routine that accepts a int16 pointer and returns a
// value to it.
func Int16(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16PtrSlice converts a slice of int16 values into a slice of int16
// pointers.
func Int16PtrSlice(src []int16) []*int16 {
	dst := make([]*int16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int16Slice converts a slice of int16 pointers into a slice of int16
// values.
func Int16Slice(src []*int16) []int16 {
	dst := make([]int16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int16PtrMap converts a string map of int16 values into a string map of int16
// pointers.
func Int16PtrMap(src map[string]int16) map[string]*int16 {
	dst := make(map[string]*int16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int16Map converts a string map of int16 pointers into a string map of
// int16 values.
func Int16Map(src map[string]*int16) map[string]int16 {
	dst := make(map[string]int16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int32Ptr is a helper routine that allocates a new int32 value to store v
// and returns a pointer to it.
func Int32Ptr(v int32) *int32 { return &v }

// Int32 is a helper routine that accepts a int32 pointer and returns a
// value to it.
func Int32(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32PtrSlice converts a slice of int32 values into a slice of int32
// pointers.
func Int32PtrSlice(src []int32) []*int32 {
	dst := make([]*int32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int32Slice converts a slice of int32 pointers into a slice of int32
// values.
func Int32Slice(src []*int32) []int32 {
	dst := make([]int32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int32PtrMap converts a string map of int32 values into a string map of int32
// pointers.
func Int32PtrMap(src map[string]int32) map[string]*int32 {
	dst := make(map[string]*int32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int32Map converts a string map of int32 pointers into a string map of
// int32 values.
func Int32Map(src map[string]*int32) map[string]int32 {
	dst := make(map[string]int32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int64Ptr is a helper routine that allocates a new int64 value to store v
// and returns a pointer to it.
func Int64Ptr(v int64) *int64 { return &v }

// Int64 is a helper routine that accepts a int64 pointer and returns a
// value to it.
func Int64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64PtrSlice converts a slice of int64 values into a slice of int64
// pointers.
func Int64PtrSlice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int64Slice converts a slice of int64 pointers into a slice of int64
// values.
func Int64Slice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int64PtrMap converts a string map of int64 values into a string map of int64
// pointers.
func Int64PtrMap(src map[string]int64) map[string]*int64 {
	dst := make(map[string]*int64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int64Map converts a string map of int64 pointers into a string map of
// int64 values.
func Int64Map(src map[string]*int64) map[string]int64 {
	dst := make(map[string]int64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Rune is a helper routine that accepts a rune pointer and returns a value
// to it.
func Rune(v *rune) rune {
	if v != nil {
		return *v
	}
	return rune(0)
}

// StringPtr is a helper routine that allocates a new string value to store v
// and returns a pointer to it.
func StringPtr(v string) *string { return &v }

// String is a helper routine that accepts a string pointer and returns a
// value to it.
func String(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringPtrSlice converts a slice of string values into a slice of string
// pointers.
func StringPtrSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// StringSlice converts a slice of string pointers into a slice of string
// values.
func StringSlice(src []*string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// StringPtrMap converts a string map of string values into a string map of
// string pointers.
func StringPtrMap(src map[string]string) map[string]*string {
	dst := make(map[string]*string)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// StringMap converts a string map of string pointers into a string map of
// string values.
func StringMap(src map[string]*string) map[string]string {
	dst := make(map[string]string)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// UintPtr is a helper routine that allocates a new uint value to store v
// and returns a pointer to it.
func UintPtr(v uint) *uint { return &v }

// Uint is a helper routine that accepts a uint pointer and returns a value
// to it.
func Uint(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintPtrSlice converts a slice of uint values uinto a slice of uint pointers.
func UintPtrSlice(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// UintSlice converts a slice of uint pointers uinto a slice of uint
// values.
func UintSlice(src []*uint) []uint {
	dst := make([]uint, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// UintPtrMap converts a string map of uint values uinto a string map of uint
// pointers.
func UintPtrMap(src map[string]uint) map[string]*uint {
	dst := make(map[string]*uint)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// UintMap converts a string map of uint pointers uinto a string map of
// uint values.
func UintMap(src map[string]*uint) map[string]uint {
	dst := make(map[string]uint)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint8Ptr is a helper routine that allocates a new uint8 value to store v
// and returns a pointer to it.
func Uint8Ptr(v uint8) *uint8 { return &v }

// Uint8 is a helper routine that accepts a uint8 pointer and returns a
// value to it.
func Uint8(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8PtrSlice converts a slice of uint8 values into a slice of uint8
// pointers.
func Uint8PtrSlice(src []uint8) []*uint8 {
	dst := make([]*uint8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint8Slice converts a slice of uint8 pointers into a slice of uint8
// values.
func Uint8Slice(src []*uint8) []uint8 {
	dst := make([]uint8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint8PtrMap converts a string map of uint8 values into a string map of uint8
// pointers.
func Uint8PtrMap(src map[string]uint8) map[string]*uint8 {
	dst := make(map[string]*uint8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint8Map converts a string map of uint8 pointers into a string
// map of uint8 values.
func Uint8Map(src map[string]*uint8) map[string]uint8 {
	dst := make(map[string]uint8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint16Ptr is a helper routine that allocates a new uint16 value to store v
// and returns a pointer to it.
func Uint16Ptr(v uint16) *uint16 { return &v }

// Uint16 is a helper routine that accepts a uint16 pointer and returns a
// value to it.
func Uint16(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16PtrSlice converts a slice of uint16 values into a slice of uint16
// pointers.
func Uint16PtrSlice(src []uint16) []*uint16 {
	dst := make([]*uint16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint16Slice converts a slice of uint16 pointers into a slice of uint16
// values.
func Uint16Slice(src []*uint16) []uint16 {
	dst := make([]uint16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint16PtrMap converts a string map of uint16 values into a string map of
// uint16 pointers.
func Uint16PtrMap(src map[string]uint16) map[string]*uint16 {
	dst := make(map[string]*uint16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint16Map converts a string map of uint16 pointers into a string map of
// uint16 values.
func Uint16Map(src map[string]*uint16) map[string]uint16 {
	dst := make(map[string]uint16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint32Ptr is a helper routine that allocates a new uint32 value to store v
// and returns a pointer to it.
func Uint32Ptr(v uint32) *uint32 { return &v }

// Uint32 is a helper routine that accepts a uint32 pointer and returns a
// value to it.
func Uint32(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32PtrSlice converts a slice of uint32 values into a slice of uint32
// pointers.
func Uint32PtrSlice(src []uint32) []*uint32 {
	dst := make([]*uint32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint32Slice converts a slice of uint32 pointers into a slice of uint32
// values.
func Uint32Slice(src []*uint32) []uint32 {
	dst := make([]uint32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint32PtrMap converts a string map of uint32 values into a string map of
// uint32 pointers.
func Uint32PtrMap(src map[string]uint32) map[string]*uint32 {
	dst := make(map[string]*uint32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint32Map converts a string map of uint32 pointers into a string
// map of uint32 values.
func Uint32Map(src map[string]*uint32) map[string]uint32 {
	dst := make(map[string]uint32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint64Ptr is a helper routine that allocates a new uint64 value to store v
// and returns a pointer to it.
func Uint64Ptr(v uint64) *uint64 { return &v }

// Uint64 is a helper routine that accepts a uint64 pointer and returns a
// value to it.
func Uint64(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64PtrSlice converts a slice of uint64 values into a slice of uint64
// pointers.
func Uint64PtrSlice(src []uint64) []*uint64 {
	dst := make([]*uint64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint64Slice converts a slice of uint64 pointers into a slice of uint64
// values.
func Uint64Slice(src []*uint64) []uint64 {
	dst := make([]uint64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint64PtrMap converts a string map of uint64 values into a string map of
// uint64 pointers.
func Uint64PtrMap(src map[string]uint64) map[string]*uint64 {
	dst := make(map[string]*uint64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint64Map converts a string map of uint64 pointers into a string map of
// uint64 values.
func Uint64Map(src map[string]*uint64) map[string]uint64 {
	dst := make(map[string]uint64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Time is a helper routine that accepts a time pointer value and returns a
// value to it.
func Time(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// Duration is a helper routine that accepts a time pointer ion value
// and returns a value to it.
// func Duration(v *time.Duration) time.Duration {
// 	if v != nil {
// 		return *v
// 	}
// 	return time.Duration(0)
// }
