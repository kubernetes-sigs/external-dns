/*
Copyright 2023 The Kubernetes Authors.

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

package tlsutils

import (
	"crypto/tls"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
)

var rsaCertPEM = `-----BEGIN CERTIFICATE-----
MIIB0zCCAX2gAwIBAgIJAI/M7BYjwB+uMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTIwOTEyMjE1MjAyWhcNMTUwOTEyMjE1MjAyWjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANLJ
hPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wok/4xIA+ui35/MmNa
rtNuC+BdZ1tMuVCPFZcCAwEAAaNQME4wHQYDVR0OBBYEFJvKs8RfJaXTH08W+SGv
zQyKn0H8MB8GA1UdIwQYMBaAFJvKs8RfJaXTH08W+SGvzQyKn0H8MAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQEFBQADQQBJlffJHybjDGxRMqaRmDhX0+6v02TUKZsW
r5QuVbpQhH6u+0UgcW0jp9QwpxoPTLTWGXEWBBBurxFwiCBhkQ+V
-----END CERTIFICATE-----
`

var rsaKeyPEM = testingKey(`-----BEGIN RSA TESTING KEY-----
MIIBOwIBAAJBANLJhPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wo
k/4xIA+ui35/MmNartNuC+BdZ1tMuVCPFZcCAwEAAQJAEJ2N+zsR0Xn8/Q6twa4G
6OB1M1WO+k+ztnX/1SvNeWu8D6GImtupLTYgjZcHufykj09jiHmjHx8u8ZZB/o1N
MQIhAPW+eyZo7ay3lMz1V01WVjNKK9QSn1MJlb06h/LuYv9FAiEA25WPedKgVyCW
SmUwbPw8fnTcpqDWE3yTO3vKcebqMSsCIBF3UmVue8YU3jybC3NxuXq3wNm34R8T
xVLHwDXh/6NJAiEAl2oHGGLz64BuAfjKrqwz7qMYr9HCLIe/YsoWq/olzScCIQDi
D2lWusoe2/nEqfDVVWGWlyJ7yOmqaVm/iNUN9B2N2g==
-----END RSA TESTING KEY-----
`)

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func TestCreateTLSConfig(t *testing.T) {

	tests := []struct {
		title         string
		prefix        string
		caFile        string
		certFile      string
		keyFile       string
		isInsecureStr string
		serverName    string
		assertions    func(actual *tls.Config, err error)
	}{
		{
			"Provide only CA returns error",
			"prefix",
			"",
			rsaCertPEM,
			"",
			"",
			"",
			func(actual *tls.Config, err error) {
				assert.Contains(t, err.Error(), "either both cert and key or none must be provided")
			},
		},
		{
			"Invalid cert and key returns error",
			"prefix",
			"",
			"invalid-cert",
			"invalid-key",
			"",
			"",
			func(actual *tls.Config, err error) {
				assert.Contains(t, err.Error(), "could not load TLS cert")
			},
		},
		{
			"Valid cert and key return a valid tls.Config with a certificate",
			"prefix",
			"",
			rsaCertPEM,
			rsaKeyPEM,
			"",
			"server-name",
			func(actual *tls.Config, err error) {
				require.NoError(t, err)
				assert.Equal(t, "server-name", actual.ServerName)
				assert.NotNil(t, actual.Certificates[0])
				assert.False(t, actual.InsecureSkipVerify)
				assert.Equal(t, actual.MinVersion, uint16(defaultMinVersion))
			},
		},
		{
			"Invalid CA file returns error",
			"prefix",
			"invalid-ca-content",
			"",
			"",
			"",
			"",
			func(actual *tls.Config, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "could not read root certs")
			},
		},
		{
			"Invalid CA file path returns error",
			"prefix",
			"ca-path-does-not-exist",
			"",
			"",
			"",
			"server-name",
			func(actual *tls.Config, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "error reading /path/does/not/exist")
			},
		},
		{
			"Complete config with CA, cert, and key returns valid tls.Config",
			"prefix",
			rsaCertPEM,
			rsaCertPEM,
			rsaKeyPEM,
			"",
			"server-name",
			func(actual *tls.Config, err error) {
				require.NoError(t, err)
				assert.Equal(t, "server-name", actual.ServerName)
				assert.NotNil(t, actual.Certificates[0])
				assert.NotNil(t, actual.RootCAs)
				assert.False(t, actual.InsecureSkipVerify)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			// setup
			dir := t.TempDir()

			if tc.caFile != "" {
				path := fmt.Sprintf("%s/caFile", dir)
				utils.WriteToFile(path, tc.caFile)
				t.Setenv(fmt.Sprintf("%s_CA_FILE", tc.prefix), path)
			}

			if tc.caFile == "ca-path-does-not-exist" {
				t.Setenv(fmt.Sprintf("%s_CA_FILE", tc.prefix), "/path/does/not/exist")
			}

			if tc.certFile != "" {
				path := fmt.Sprintf("%s/certFile", dir)
				utils.WriteToFile(path, tc.certFile)
				t.Setenv(fmt.Sprintf("%s_CERT_FILE", tc.prefix), path)
			}

			if tc.keyFile != "" {
				path := fmt.Sprintf("%s/keyFile", dir)
				utils.WriteToFile(path, tc.keyFile)
				t.Setenv(fmt.Sprintf("%s_KEY_FILE", tc.prefix), path)
			}

			if tc.serverName != "" {
				t.Setenv(fmt.Sprintf("%s_TLS_SERVER_NAME", tc.prefix), tc.serverName)
			}

			if tc.isInsecureStr != "" {
				t.Setenv(fmt.Sprintf("%s_INSECURE", tc.prefix), tc.isInsecureStr)
			}

			// test
			actual, err := CreateTLSConfig(tc.prefix)
			tc.assertions(actual, err)
		})
	}

}
