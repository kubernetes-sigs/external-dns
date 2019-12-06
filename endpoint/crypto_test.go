/*
Copyright 2017 The Kubernetes Authors.

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

package endpoint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	// Verify that text encryption and decryption works
	aesKey := []byte("s%zF`.*'5`9.AhI2!B,.~hmbs^.*TL?;")
	plaintext := "Til pigerene havde han skemtsomme ord, han spøkte med byens børn, han svingede sydvesten og sprang ombord; så heiste han fokken, og hjem han for i solskin, den gamle ørn."
	encryptedtext, err := EncryptText(plaintext, aesKey)
	require.NoError(t, err)
	decryptedtext, err := DecryptText(encryptedtext, aesKey)
	require.NoError(t, err)
	if plaintext != decryptedtext {
		t.Errorf("Original plain text %#v differs from the resulting decrypted text %#v", plaintext, decryptedtext)
	}

	// Verify that decrypt returns an error and unmodified data if wrong AES encryption key is used
	decryptedtext, err = DecryptText(encryptedtext, []byte("s'J!jD`].LC?g&Oa11AgTub,j48ts/96"))
	require.Error(t, err)
	if decryptedtext != encryptedtext {
		t.Error("Data decryption failed, but decrypt still didn't return the original input")
	}

	// Verify that decrypt returns an error and unmodified data if unencrypted input is is supplied
	decryptedtext, err = DecryptText(plaintext, aesKey)
	require.Error(t, err)
	if plaintext != decryptedtext {
		t.Errorf("Data decryption failed, but decrypt still didn't return the original input. Original input %#v, returned data %#v", plaintext, decryptedtext)
	}

	// Verify that a known encrypted text is decrypted to what is expected
	encryptedtext = "TIBnEeYWYd+sffKZ6Wk3Js0pR58TbsFxhXUc+6jzgS4AXCb+NSuKjnlRqp6e5imh5b1hrxWmecrI/QNJzC6o/2U5FZfxY5CjSqhIbCV/gWWkgK1Qd1ohzCO4tzkF5fwGCxSW715NLDUImg6xE4D/hI2pY8wcATnvv4qUw6eR78kDzfn4cwO1IeS/HmFn+ASldPvY6Y3JHCSlbvWTekfRrknHJeLHIsPW5yZVPZCq29xGEe+A29Y="
	decryptedtext, err = DecryptText(encryptedtext, aesKey)
	require.NoError(t, err)
	if decryptedtext != plaintext {
		t.Error("Decryption of text didn't result in expected plaintext result.")
	}
}
