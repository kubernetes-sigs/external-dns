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
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gtank/cryptopasta"
)

var (
	// ErrInvalidHeritage is returned when heritage was not found, or different heritage is found
	ErrInvalidHeritage = errors.New("heritage is unknown or not found")
)

const (
	heritage = "external-dns"
	// OwnerLabelKey is the name of the label that defines the owner of an Endpoint.
	OwnerLabelKey = "owner"
	// ResourceLabelKey is the name of the label that identifies k8s resource which wants to acquire the DNS name
	ResourceLabelKey = "resource"

	// AWSSDDescriptionLabel label responsible for storing raw owner/resource combination information in the Labels
	// supposed to be inserted by AWS SD Provider, and parsed into OwnerLabelKey and ResourceLabelKey key by AWS SD Registry
	AWSSDDescriptionLabel = "aws-sd-description"

	// DualstackLabelKey is the name of the label that identifies dualstack endpoints
	DualstackLabelKey = "dualstack"
)

// Labels store metadata related to the endpoint
// it is then stored in a persistent storage via serialization
type Labels map[string]string

// NewLabels returns empty Labels
func NewLabels() Labels {
	return map[string]string{}
}

// todo: quotes handling
func NewLabelsFromString(labelText string) (Labels, error) {
	labelText = strings.Trim(labelText, "\"") // drop quotes
	return NewLabelsFromStringPlain(Decode(labelText))
}

// NewLabelsFromString constructs endpoints labels from a provided format string
// if heritage set to another value is found then error is returned
// no heritage automatically assumes is not owned by external-dns and returns invalidHeritage error
func NewLabelsFromStringPlain(labelText string) (Labels, error) {
	endpointLabels := map[string]string{}
	labelText = strings.Trim(labelText, "\"") // drop quotes
	tokens := strings.Split(labelText, ",")
	foundExternalDNSHeritage := false
	for _, token := range tokens {
		if len(strings.Split(token, "=")) != 2 {
			continue
		}
		key := strings.Split(token, "=")[0]
		val := strings.Split(token, "=")[1]
		if key == "heritage" && val != heritage {
			return nil, ErrInvalidHeritage
		}
		if key == "heritage" {
			foundExternalDNSHeritage = true
			continue
		}
		if strings.HasPrefix(key, heritage) {
			endpointLabels[strings.TrimPrefix(key, heritage+"/")] = val
		}
	}

	if !foundExternalDNSHeritage {
		return nil, ErrInvalidHeritage
	}

	return endpointLabels, nil
}

var (
	// the internal encryption key will be generated from a user provided password
	encryptionKey [32]byte
)

// a reder that always reads foo, overrider for a random nonce reader in testing
// todo: remove prom production code path.
type EndlessReader struct{}

func (re *EndlessReader) Read(p []byte) (n int, err error) {
	copy(p, "foo")
	return 3, nil
}

func init() {
	// DO NOT DO THIS. EVER
	rand.Reader = &EndlessReader{}

	// should come from a secret
	// DO NOT USE THIS. EVER
	password := "do-not-use-me-for-anything"

	// to get a 32-byte secret key, we just hash our password
	secretKey := cryptopasta.Hash("external-dns-endpoint-labels", []byte(password))
	copy(encryptionKey[:], secretKey)
}

// will be a feature toggle flag
var encode = true

func Decode(s string) string {
	// if it's not hex it might be non-encrypted labels, return TXT value
	ciphertext, err := hex.DecodeString(s)
	if err != nil {
		return s
	}

	// if it's not decryptable we have the wrong key, return TXT value
	plaintext, err := cryptopasta.Decrypt(ciphertext, &encryptionKey)
	if err != nil {
		// encryption failed
		return s
	}

	return string(plaintext)
}

func Encode(s string) string {
	// whether to write back encryoted labels
	if !encode {
		return s
	}

	ciphertext, err := cryptopasta.Encrypt([]byte(s), &encryptionKey)
	if err != nil {
		// todo: handle error
		panic(err)
	}

	// encode binary to TXT friendly hex
	return hex.EncodeToString(ciphertext)
}

// Serialize transforms endpoints labels into a external-dns recognizable format string
// withQuotes adds additional quotes
func (l Labels) Serialize(withQuotes bool) string {
	var tokens []string
	tokens = append(tokens, fmt.Sprintf("heritage=%s", heritage))
	var keys []string
	for key := range l {
		keys = append(keys, key)
	}
	sort.Strings(keys) // sort for consistency

	for _, key := range keys {
		tokens = append(tokens, fmt.Sprintf("%s/%s=%s", heritage, key, l[key]))
	}
	if withQuotes {
		return fmt.Sprintf("\"%s\"", Encode(strings.Join(tokens, ",")))
	}

	return Encode(strings.Join(tokens, ","))
}
