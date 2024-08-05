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

<<<<<<< HEAD
<<<<<<< HEAD
	pkey, ok := parsed.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key was no RSA key: %T", parsed)
	}
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	pkey := parsed.(*rsa.PrivateKey)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	pkey := parsed.(*rsa.PrivateKey)
=======
	pkey, ok := parsed.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key was no RSA key: %T", parsed)
	}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	digest := sha512.Sum512(body)

	enc, err := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA512, digest[:])
	if err != nil {
		return "", fmt.Errorf("could not sign data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}
