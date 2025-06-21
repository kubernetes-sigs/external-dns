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

package registry

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

func TestNewTXTRegistryEncryptionConfig(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	tests := []struct {
		encEnabled      bool
		aesKeyRaw       []byte
		aesKeySanitized []byte
		errorExpected   bool
	}{
		{
			encEnabled:      true,
			aesKeyRaw:       []byte("123456789012345678901234567890asdfasdfasdfasdfa12"),
			aesKeySanitized: []byte{},
			errorExpected:   true,
		},
		{
			encEnabled:      true,
			aesKeyRaw:       []byte("passphrasewhichneedstobe32bytes!"),
			aesKeySanitized: []byte("passphrasewhichneedstobe32bytes!"),
			errorExpected:   false,
		},
		{
			encEnabled:      true,
			aesKeyRaw:       []byte("ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY="),
			aesKeySanitized: []byte{100, 248, 173, 47, 67, 70, 85, 0, 89, 109, 48, 250, 15, 5, 201, 204, 63, 17, 137, 43, 82, 107, 60, 216, 93, 11, 29, 82, 140, 11, 81, 22},
			errorExpected:   false,
		},
	}
	for _, test := range tests {
		actual, err := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "", []string{}, []string{}, test.encEnabled, test.aesKeyRaw)
		if test.errorExpected {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, test.aesKeySanitized, actual.txtEncryptAESKey)
		}
	}
}

func TestGenerateTXTGenerateTextRecordEncryptionWihDecryption(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone(testZone)

	tests := []struct {
		record    *endpoint.Endpoint
		decrypted string
	}{
		{
			record:    newEndpointWithOwner("foo.test-zone.example.org", "new-foo.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2"),
			decrypted: "heritage=external-dns,external-dns/owner=owner-2",
		},
		{
			record:    newEndpointWithOwnerAndLabels("foo.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "owner-1", endpoint.Labels{endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org"}),
			decrypted: "heritage=external-dns,external-dns/ownedRecord=foo.test-zone.example.org,external-dns/owner=owner-1",
		},
		{
			record:    newEndpointWithOwnerAndLabels("bar.test-zone.example.org", "cluster-b", endpoint.RecordTypeCNAME, "owner-1", endpoint.Labels{endpoint.ResourceLabelKey: "ingress/default/foo-127"}),
			decrypted: "heritage=external-dns,external-dns/owner=owner-1,external-dns/resource=ingress/default/foo-127",
		},
		{
			record:    newEndpointWithOwner("dualstack.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, "owner-0"),
			decrypted: "heritage=external-dns,external-dns/owner=owner-0",
		},
	}

	withEncryptionKeys := []string{
		"passphrasewhichneedstobe32bytes!",
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=",
		"01234567890123456789012345678901",
	}

	for _, test := range tests {
		for _, k := range withEncryptionKeys {
			t.Run(fmt.Sprintf("key '%s' with decrypted result '%s'", k, test.decrypted), func(t *testing.T) {
				key := []byte(k)
				r, err := NewTXTRegistry(p, "", "", "owner", time.Minute, "", []string{}, []string{}, true, key)
				assert.NoError(t, err, "Error creating TXT registry")
				txtRecords := r.generateTXTRecord(test.record)
				assert.Len(t, txtRecords, len(test.record.Targets))

				for _, txt := range txtRecords {
					// should return a TXT record with the encryption nonce label. At the moment nonce is not set as label.
					assert.NotContains(t, txt.Labels, "txt-encryption-nonce")

					assert.Len(t, txt.Targets, 1)
					assert.LessOrEqual(t, len(txt.Targets), 1)

					// decrypt targets
					for _, target := range txtRecords[0].Targets {
						encryptedText, errUnquote := strconv.Unquote(target)
						assert.NoError(t, errUnquote, "Error unquoting the encrypted text")

						actual, nonce, errDecrypt := endpoint.DecryptText(encryptedText, r.txtEncryptAESKey)
						assert.NoError(t, errDecrypt, "Error decrypting the encrypted text")

						assert.True(t, strings.HasPrefix(encryptedText, nonce),
							"Nonce '%s' should be a prefix of the encrypted text: '%s'", nonce, encryptedText)
						assert.Equal(t, test.decrypted, actual)
					}
				}
			})
		}
	}
}

func TestApplyRecordsWithEncryption(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone("org")

	key := []byte("ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=")

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, key)

	_ = r.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwner("example.org", "new-loadbalancer-3.org", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("main.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwner("tar.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2"),
			newEndpointWithOwner("thing3.org", "1.2.3.4", endpoint.RecordTypeA, "owner"),
			newEndpointWithOwner("thing4.org", "2001:DB8::2", endpoint.RecordTypeAAAA, "owner"),
		},
	})

	allPlainTextTargetsToAssert := []string{
		"heritage=external-dns,external-dns/",
		"tar.loadbalancer.com",
		"new-loadbalancer-1.lb.com",
		"2001:DB8::2",
		"new-loadbalancer-3.org",
		"1.2.3.4",
	}

	records, _ := p.Records(ctx)
	assert.Len(t, records, 14)
	for _, r := range records {
		if r.RecordType == endpoint.RecordTypeTXT && (strings.HasPrefix(r.DNSName, "cname-") || strings.HasPrefix(r.DNSName, "txt-new-")) {
			assert.NotContains(t, r.Labels, "txt-encryption-nonce")
			// assuming single target, it should be not a plain text
			assert.NotContains(t, r.Targets[0], "heritage=external-dns")
		}
		// All TXT records with new- prefix should have the encryption nonce label and be in plain text
		if r.RecordType == endpoint.RecordTypeTXT && strings.HasPrefix(r.DNSName, "new-") {
			assert.Contains(t, r.Labels, "txt-encryption-nonce")
			// assuming single target, it should be in a plain text
			assert.Contains(t, r.Targets[0], "heritage=external-dns,external-dns/")
		}
		// All CNAME, A and AAAA TXT records should have the encryption nonce label
		if slices.Contains([]string{"CNAME", "A", "AAAA"}, r.RecordType) {
			assert.Contains(t, r.Labels, "txt-encryption-nonce")
			// validate that target is in plain text
			assert.Contains(t, allPlainTextTargetsToAssert, r.Targets[0])
		}
	}
}

func TestApplyRecordsWithEncryptionKeyChanged(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone("org")

	withEncryptionKeys := []string{
		"passphrasewhichneedstobe32bytes!",
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=",
		"01234567890123456789012345678901",
	}

	for _, key := range withEncryptionKeys {
		r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, []byte(key))
		_ = r.ApplyChanges(ctx, &plan.Changes{
			Create: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
				newEndpointWithOwnerAndOwnedRecord("new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
				newEndpointWithOwner("example.org", "new-loadbalancer-3.org", endpoint.RecordTypeCNAME, "owner"),
				newEndpointWithOwnerAndOwnedRecord("main.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
				newEndpointWithOwner("tar.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2"),
				newEndpointWithOwner("thing3.org", "1.2.3.4", endpoint.RecordTypeA, "owner"),
				newEndpointWithOwner("thing4.org", "2001:DB8::2", endpoint.RecordTypeAAAA, "owner"),
			},
		})
	}

	records, _ := p.Records(ctx)
	assert.Len(t, records, 14)
}

func TestApplyRecordsOnEncryptionKeyChangeWithKeyIdLabel(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone("org")

	withEncryptionKeys := []string{
		"passphrasewhichneedstobe32bytes!",
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=",
		"01234567890123456789012345678901",
	}

	for i, key := range withEncryptionKeys {
		r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, []byte(key))
		keyId := fmt.Sprintf("key-id-%d", i)
		changes := []*endpoint.Endpoint{
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("example.org", "new-loadbalancer-3.org", endpoint.RecordTypeCNAME, "owner", "", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("main.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("tar.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2", "", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("thing3.org", "1.2.3.4", endpoint.RecordTypeA, "owner", "", keyId),
			newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel("thing4.org", "2001:DB8::2", endpoint.RecordTypeAAAA, "owner", "", keyId),
		}

		if i == 0 {
			_ = r.ApplyChanges(ctx, &plan.Changes{
				Create: changes,
			})
		} else {
			_ = r.ApplyChanges(context.Background(), &plan.Changes{
				UpdateNew: changes,
			})
		}

	}

	records, _ := p.Records(ctx)
	assert.Len(t, records, 14)

	encryptionNonce := map[string]bool{}

	for _, r := range records {
		if slices.Contains([]string{"A", "AAAA"}, r.RecordType) || (r.RecordType == "CNAME" && strings.HasPrefix(r.DNSName, "new-")) {
			assert.Contains(t, r.Labels, "key-id")
			assert.Equal(t, "key-id-2", r.Labels["key-id"])
			// add encryption nonce to track the number of unique nonce
			encryptionNonce[r.Labels["txt-encryption-nonce"]] = true
		} else if r.RecordType == endpoint.RecordTypeTXT {
			if hasPrefixFromSlice(r.DNSName, []string{"cname-", "txt-new-", "a-", "aaaa-", "txt-"}) {
				assert.NotContains(t, r.Labels, "key-id")
			} else {
				assert.Contains(t, r.Labels, "key-id", r.DNSName)
				assert.Equal(t, "key-id-0", r.Labels["key-id"], r.DNSName)
				// add encryption nonce to track the number of unique nonce
				encryptionNonce[r.Labels["txt-encryption-nonce"]] = true
			}
		}
	}
	assert.LessOrEqual(t, len(encryptionNonce), 5)
}

func hasPrefixFromSlice(str string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

func newEndpointWithOwnerAndOwnedRecordWithKeyIDLabel(dnsName, target, recordType, ownerID string, resource string, keyId string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, recordType, target)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	e.Labels[endpoint.ResourceLabelKey] = resource
	e.Labels["key-id"] = keyId
	return e
}
