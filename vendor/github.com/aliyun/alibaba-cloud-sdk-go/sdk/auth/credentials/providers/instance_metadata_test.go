package providers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInstanceMetadataProvider_Retrieve_Success(t *testing.T) {

	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := "unsupported path: " + r.URL.Path
		status := 500

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 200
		case "/latest/meta-data/ram/security-credentials/ELK":
			body = ` {
			  "AccessKeyId" : "STS.L4aBSCSJVMuKg5U1vFDw",
			  "AccessKeySecret" : "wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pKCNZ9",
			  "Expiration" : "2018-08-20T22:30:02Z",
			  "SecurityToken" : "CAESrAIIARKAAShQquMnLIlbvEcIxO6wCoqJufs8sWwieUxu45hS9AvKNEte8KRUWiJWJ6Y+YHAPgNwi7yfRecMFydL2uPOgBI7LDio0RkbYLmJfIxHM2nGBPdml7kYEOXmJp2aDhbvvwVYIyt/8iES/R6N208wQh0Pk2bu+/9dvalp6wOHF4gkFGhhTVFMuTDRhQlNDU0pWTXVLZzVVMXZGRHciBTQzMjc0KgVhbGljZTCpnJjwySk6BlJzYU1ENUJuCgExGmkKBUFsbG93Eh8KDEFjdGlvbkVxdWFscxIGQWN0aW9uGgcKBW9zczoqEj8KDlJlc291cmNlRXF1YWxzEghSZXNvdXJjZRojCiFhY3M6b3NzOio6NDMyNzQ6c2FtcGxlYm94L2FsaWNlLyo=",
			  "LastUpdated" : "2018-08-20T16:30:01Z",
			  "Code" : "Success"
			}`
			status = 200
		}
		w.Write([]byte(body))
		w.WriteHeader(status)
	}))
	defer ts.Close()

	// Update our securityCredURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
	}()

	credential, err := NewInstanceMetadataProvider().Retrieve()
	if err != nil {
		t.Fatal(err)
	}

	stsTokenCredential, ok := credential.(*credentials.StsTokenCredential)
	if !ok {
		t.Fatal("expected AccessKeyCredential")
	}

	if stsTokenCredential.AccessKeyId != "STS.L4aBSCSJVMuKg5U1vFDw" {
		t.Fatalf("expected AccessKeyId STS.L4aBSCSJVMuKg5U1vFDw but received %s", stsTokenCredential.AccessKeyId)
	}
	if stsTokenCredential.AccessKeySecret != "wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pKCNZ9" {
		t.Fatalf("expected AccessKeySecret wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pKCNZ9 but received %s", stsTokenCredential.AccessKeySecret)
	}
	if !strings.HasPrefix(stsTokenCredential.AccessKeyStsToken, "CAESrAIIARKAA") {
		t.Fatalf("expected AccessKeyStsToken starting with CAESrAIIARKAA but received %s", stsTokenCredential.AccessKeyStsToken)
	}
}
