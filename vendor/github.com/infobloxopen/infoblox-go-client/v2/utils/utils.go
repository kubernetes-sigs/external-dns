package utils

// BoolPtr returns a pointer to the bool value passed as argument
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns a pointer to the string value passed as argument
func StringPtr(v string) *string {
	return &v
}

// Uint32Ptr returns a pointer to uint32 value passed as argument
func Uint32Ptr(v uint32) *uint32 {
	return &v
}
