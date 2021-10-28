package ptr

import "github.com/ukfast/sdk-go/pkg/connection"

// Byte returns a pointer to a byte
func Byte(v byte) *byte {
	return &v
}

// Int returns a pointer to an integer
func Int(v int) *int {
	return &v
}

// Int8 returns a pointer to an 8-bit integer
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to a 16-bit integer
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to a 32-bit integer
func Int32(v int32) *int32 {
	return &v
}

// Int64 returns a pointer to a 64-bit integer
func Int64(v int64) *int64 {
	return &v
}

// UInt returns a pointer to an unsigned integer
func UInt(v uint) *uint {
	return &v
}

// UInt8 returns a pointer to an 8-bit unsigned integer
func UInt8(v uint8) *uint8 {
	return &v
}

// UInt16 returns a pointer to a 16-bit unsigned integer
func UInt16(v uint16) *uint16 {
	return &v
}

// UInt32 returns a pointer to a 32-bit unsigned integer
func UInt32(v uint32) *uint32 {
	return &v
}

// UInt64 returns a pointer to a 64-bit unsigned integer
func UInt64(v uint64) *uint64 {
	return &v
}

// UIntPtr returns a pointer to an unsigned integer
func UIntPtr(v uintptr) *uintptr {
	return &v
}

// Float32 returns a pointer to a 32-bit floating-point number
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to a 64-bit floating-point number
func Float64(v float64) *float64 {
	return &v
}

// Complex64 returns a pointer to a 64-bit complex number
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex128 returns a pointer to a 128-bit complex number
func Complex128(v complex128) *complex128 {
	return &v
}

// String returns a pointer to a string
func String(v string) *string {
	return &v
}

// Bool returns a pointer to a boolean
func Bool(v bool) *bool {
	return &v
}

// Date returns a pointer to a Date
func Date(v connection.Date) *connection.Date {
	return &v
}

// DateTime returns a pointer to a DateTime
func DateTime(v connection.DateTime) *connection.DateTime {
	return &v
}

// IPAddress returns a pointer to a IPAddress
func IPAddress(v connection.IPAddress) *connection.IPAddress {
	return &v
}

// ToIntOrDefault is a helper method for retrieving an integer from a pointer to an
// integer value (or zero value if nil)
func ToIntOrDefault(i *int) int {
	if i != nil {
		return *i
	}

	return 0
}

// ToInt8OrDefault is a helper method for retrieving an 8-bit integer from a pointer to an
// integer value (or zero value if nil)
func ToInt8OrDefault(i *int8) int8 {
	if i != nil {
		return *i
	}

	return 0
}

// ToInt16OrDefault is a helper method for retrieving a 16-bit integer from a pointer to an
// integer value (or zero value if nil)
func ToInt16OrDefault(i *int16) int16 {
	if i != nil {
		return *i
	}

	return 0
}

// ToInt32OrDefault is a helper method for retrieving a 32-bit integer from a pointer to an
// integer value (or zero value if nil)
func ToInt32OrDefault(i *int32) int32 {
	if i != nil {
		return *i
	}

	return 0
}

// ToInt64OrDefault is a helper method for retrieving a 64-bit integer from a pointer to an
// integer value (or zero value if nil)
func ToInt64OrDefault(i *int64) int64 {
	if i != nil {
		return *i
	}

	return 0
}

// ToUIntOrDefault is a helper method for retrieving an unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUIntOrDefault(i *uint) uint {
	if i != nil {
		return *i
	}

	return 0
}

// ToUInt8OrDefault is a helper method for retrieving an 8-bit unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUInt8OrDefault(i *uint8) uint8 {
	if i != nil {
		return *i
	}

	return 0
}

// ToUInt16OrDefault is a helper method for retrieving a 16-bit unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUInt16OrDefault(i *uint16) uint16 {
	if i != nil {
		return *i
	}

	return 0
}

// ToUInt32OrDefault is a helper method for retrieving a 32-bit unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUInt32OrDefault(i *uint32) uint32 {
	if i != nil {
		return *i
	}

	return 0
}

// ToUInt64OrDefault is a helper method for retrieving a 64-bit unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUInt64OrDefault(i *uint64) uint64 {
	if i != nil {
		return *i
	}

	return 0
}

// ToUIntPtrOrDefault is a helper method for retrieving an unsigned integer from a pointer to an
// integer value (or zero value if nil)
func ToUIntPtrOrDefault(i *uintptr) uintptr {
	if i != nil {
		return *i
	}

	return 0
}

// ToFloat32OrDefault is a helper method for retrieving a 32-bit floating-point number from a pointer to an
// floating-point number value (or zero value if nil)
func ToFloat32OrDefault(i *float32) float32 {
	if i != nil {
		return *i
	}

	return 0
}

// ToFloat64OrDefault is a helper method for retrieving a 64-bit floating-point number from a pointer to an
// floating-point number value (or zero value if nil)
func ToFloat64OrDefault(i *float64) float64 {
	if i != nil {
		return *i
	}

	return 0
}

// ToComplex64OrDefault is a helper method for retrieving a 64-bit complex number from a pointer to an
// complex number value (or zero value if nil)
func ToComplex64OrDefault(i *complex64) complex64 {
	if i != nil {
		return *i
	}

	return 0
}

// ToComplex128OrDefault is a helper method for retrieving a 128-bit complex number from a pointer to an
// complex number value (or zero value if nil)
func ToComplex128OrDefault(i *complex128) complex128 {
	if i != nil {
		return *i
	}

	return 0
}

// ToStringOrDefault is a helper method for retrieving a string from a pointer to a
// string value (or zero value if nil)
func ToStringOrDefault(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

// ToBoolOrDefault is a helper method for retrieving a boolean from a pointer to a
// boolean value (or zero value if nil)
func ToBoolOrDefault(b *bool) bool {
	if b != nil {
		return *b
	}

	return false
}

// ToDateOrDefault is a helper method for retrieving a Date from a pointer to a
// Date value (or zero value if nil)
func ToDateOrDefault(d *connection.Date) connection.Date {
	if d != nil {
		return *d
	}

	return ""
}

// ToDateTimeOrDefault is a helper method for retrieving a DateTime from a pointer to a
// DateTime value (or zero value if nil)
func ToDateTimeOrDefault(d *connection.DateTime) connection.DateTime {
	if d != nil {
		return *d
	}

	return ""
}

// ToIPAddressOrDefault is a helper method for retrieving an IPAddress from a pointer to an
// IPAddress value (or zero value if nil)
func ToIPAddressOrDefault(i *connection.IPAddress) connection.IPAddress {
	if i != nil {
		return *i
	}

	return ""
}
