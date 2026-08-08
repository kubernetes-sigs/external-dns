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
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const standardGcmNonceSize = 12

// GenerateNonce creates a random base64-encoded nonce of a fixed size.
func GenerateNonce() (string, error) {
	nonce := make([]byte, standardGcmNonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nonce), nil
}

// EncryptText gzips input data and encrypts it using the supplied AES key.
// nonceEncoded must be a base64-encoded nonce of standardGcmNonceSize bytes.
func EncryptText(text string, aesKey []byte, nonceEncoded string) (string, error) {
	if len(nonceEncoded) == 0 {
		return "", fmt.Errorf("nonce must be provided")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, standardGcmNonceSize)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, standardGcmNonceSize)
	if _, err = base64.StdEncoding.Decode(nonce, []byte(nonceEncoded)); err != nil {
		return "", err
	}

	data, err := compressData([]byte(text))
	if err != nil {
		return "", err
	}

	cipherData := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(cipherData), nil
}

// DecryptText decrypts data using the supplied AES encryption key and decompresses it.
// Returns the plaintext, the base64-encoded nonce, and any error.
func DecryptText(text string, aesKey []byte) (string, string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", "", err
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, standardGcmNonceSize)
	if err != nil {
		return "", "", err
	}
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", "", err
	}
	if len(data) <= standardGcmNonceSize {
		return "", "", fmt.Errorf("encrypted data too short: got %d bytes, need more than %d", len(data), standardGcmNonceSize)
	}
	nonce, ciphertext := data[:standardGcmNonceSize], data[standardGcmNonceSize:]
	plaindata, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", "", err
	}
	plaindata, err = decompressData(plaindata)
	if err != nil {
		return "", "", err
	}

	return string(plaindata), base64.StdEncoding.EncodeToString(nonce), nil
}

// decompressData decompresses gzip-compressed data.
func decompressData(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	var b bytes.Buffer
	if _, err = b.ReadFrom(gz); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// compressData compresses data using gzip to minimize storage in the registry.
func compressData(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err = gz.Write(data); err != nil {
		return nil, err
	}

	if err = gz.Flush(); err != nil {
		return nil, err
	}

	if err = gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
