package oapi

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

var testSeededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func testRandomID(t *testing.T) string {
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("unable to generate a new UUID: %s", err)
	}
	return id.String()
}

func testRandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[testSeededRand.Intn(len(charset))]
	}
	return string(b)
}

func testRandomString(length int) string {
	const defaultCharset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return testRandomStringWithCharset(length, defaultCharset)
}
