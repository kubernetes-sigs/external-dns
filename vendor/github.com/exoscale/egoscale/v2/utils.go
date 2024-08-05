package v2

import (
	"github.com/gofrs/uuid"
)

func mapValueOrNil(src map[string]string, key string) *string {
	if x, found := src[key]; found {
		return &x
	}

	return nil
}

func IsValidUUID(s string) bool {
	_, err := uuid.FromString(s)
	return err == nil
}
