package utils

import "fmt"

func ToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func ConvertStringPointer(in *string) string {
	if in != nil {
		return *in
	}
	return ""
}
