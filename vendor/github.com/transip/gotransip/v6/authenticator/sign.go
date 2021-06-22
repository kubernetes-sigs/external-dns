package authenticator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	// ErrDecodingPrivateKey will be thrown when an invalid private key has been given
	ErrDecodingPrivateKey = errors.New("could not decode private key")
)

func signWithKey(body []byte, key []byte) (string, error) {
	// prepare key struct
	block, _ := pem.Decode(key)
	if block == nil {
		return "", ErrDecodingPrivateKey
	}
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %w", err)
	}

	pkey := parsed.(*rsa.PrivateKey)
	digest := sha512.Sum512(body)

	enc, err := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA512, digest[:])
	if err != nil {
		return "", fmt.Errorf("could not sign data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}
