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
	aesKey := []byte("1WGwUWyQ-vnVIZNpAT8k3i2Vjm4_2xacplj_8FzF7K8=")
	plaintext := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	encryptedtext, err := EncryptText(plaintext, aesKey, nil)
	require.NoError(t, err)
	decryptedtext, _, err := DecryptText(encryptedtext, aesKey)
	require.NoError(t, err)
	if plaintext != decryptedtext {
		t.Errorf("Original plain text %#v differs from the resulting decrypted text %#v", plaintext, decryptedtext)
	}

	// Verify that decrypt returns an error and empty data if wrong AES encryption key is used
	decryptedtext, _, err = DecryptText(encryptedtext, []byte("s'J!jD`].LC?g&Oa11AgTub,j48ts/96"))
	require.Error(t, err)
	if decryptedtext != "" {
		t.Error("Data decryption failed, empty string should be as result")
	}

	// Verify that decrypt returns an error and empty data if unencrypted input is is supplied
	decryptedtext, _, err = DecryptText(plaintext, aesKey)
	require.Error(t, err)
	if decryptedtext != "" {
		t.Errorf("Data decryption failed, empty string should be as result")
	}

	// Verify that a known encrypted text is decrypted to what is expected
	encryptedtext = "AAAAAAAAAAAAAAAAvH4TwB_FqFti68hCmBuj2i5bhrNgwh01-YXi9-VpW2TIpSA0PIjJBDpRcYRxmOo0Ra_KzJ1L8CxX1LxAwp7KKmI13OEO2e-lu9GX4jJlWwSBEkkw3q-vb3mY12r9n58oILbyot_A53vqSVU-chsLqsnF0n0-Ktqt4XAkpnoPVE0YI7fE9cGuMmK0GdwEwZLR_4gsw5jAz5nFqRez3BnKjzq3gwf-n2YoCCKYBeZduVEkhFnR3CUGF7UU30_sV_y6RZJfFTgKLElEAQ_85qe2KL5i9k4g9RUXUwjY019UmhvpOWH-skHDZQ9umbPhL2TB0v9JB48vts9aGCPd81OSBwP1I5Kzb3iKYRJBNVVz5OI1RR79IDt4vw7H4Z0NzvJB5lymr3tO8nVn-gmvF4V_2gIRzk9tcx7B77vrHchJ7xXx4OIUFThR6mFoK1bTzfoaaQsD-EoK6gSOmsvigAf6jG5NXakH6q4PmjhLuMAMoDqG"
	decryptedtext, _, err = DecryptText(encryptedtext, aesKey)
	require.NoError(t, err)
	if decryptedtext != plaintext {
		t.Error("Decryption of text didn't result in expected plaintext result.")
	}
}
