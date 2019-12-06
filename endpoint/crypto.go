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

/*
This code is based on the article https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
written by Nic Raboy.
Gzip compression/decompression has ben added as there as there are issues with txt records longer than
255 characters and encryption + base64 encoding adds some characters to the resulting text record.
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

	log "github.com/sirupsen/logrus"
)

// Decompress data
func DecompressData(data []byte) (resData []byte, err error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return data, err
	}
	defer gz.Close()
	var b bytes.Buffer
	if _, err = b.ReadFrom(gz); err != nil {
		return data, err
	}
	return b.Bytes(), nil
}

// Compress data
func CompressData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return data, nil
	}
	// defer gz.Close()
	if _, err = gz.Write(data); err != nil {
		return data, nil
	}

	if err = gz.Flush(); err != nil {
		return data, nil
	}

	if err = gz.Close(); err != nil {
		return data, nil
	}

	return b.Bytes(), nil
}

// Encrypt data using the supplied AES key
func EncryptText(text string, aesKey []byte) (string, error) {
	block, _ := aes.NewCipher(aesKey)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return text, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return text, err
	}
	data := []byte(text)
	data, err = CompressData(data)
	if err != nil {
		log.Debugf("Failed to compress data based on the text %#v. Got error %#v.", text, err)
	}
	cipherdata := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(cipherdata), nil
}

// Decrypt data using a supplied AES encryption key
func DecryptText(text string, aesKey []byte) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return text, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return text, err
	}
	nonceSize := gcm.NonceSize()
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return text, err
	}
	if len(data) <= nonceSize {
		return text, fmt.Errorf("The encoded data from text %#v is shorter than %#v bytes and can't be decoded", text, nonceSize)
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaindata, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return text, err
	}
	plaindata, err = DecompressData(plaindata)
	if err != nil {
		log.Debugf("Failed to decompress data based on the base64 encoded text %#v. Got error %#v.", text, err)
	}
	return string(plaindata), nil
}
