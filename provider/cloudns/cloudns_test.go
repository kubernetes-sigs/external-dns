package cloudns

import (
	"os"
	"testing"
)

// NewClouDNSProvider creates a new ClouDNSProvider using the specified ClouDNSConfig.
// It authenticates with ClouDNS using the login type specified in the CLOUDNS_LOGIN_TYPE environment variable,
// which can be "user-id", "sub-user", or "sub-user-name". If the CLOUDNS_USER_PASSWORD environment variable is not set,
// an error will be returned. If the CLOUDNS_USER_ID or CLOUDNS_SUB_USER_ID environment variables are not set or are not valid integers,
// an error will be returned. If the CLOUDNS_SUB_USER_NAME environment variable is not set, an error will be returned.
// config is the ClouDNSConfig to be used for creating the ClouDNSProvider.
// It returns the created ClouDNSProvider and a possible error.code
// NewClouDNSProvider creates a new ClouDNSProvider using the specified ClouDNSConfig.
// It authenticates with ClouDNS using the login type specified in the CLOUDNS_LOGIN_TYPE environment variable,
// which can be "user-id", "sub-user", or "sub-user-name". If the CLOUDNS_USER_PASSWORD environment variable is not set,
// an error will be returned. If the CLOUDNS_USER_ID or CLOUDNS_SUB_USER_ID environment variables are not set or are not valid integers,
// an error will be returned. If the CLOUDNS_SUB_USER_NAME environment variable is not set, an error will be returned.
// config is the ClouDNSConfig to be used for creating the ClouDNSProvider.
// It returns the created ClouDNSProvider and a possible error.
func TestNewClouDNSProvider(t *testing.T) {
	tests := []struct {
		name             string
		loginType        string
		userID           string
		subUserID        string
		subUserName      string
		userPassword     string
		expectedError    string
		expectedErrorNil bool
	}{
		{
			name:          "valid user-id login type",
			loginType:     "user-id",
			userID:        "12345",
			userPassword:  "password",
			expectedError: "",
		},
		{
			name:             "invalid user-id login type",
			loginType:        "user-id",
			userID:           "invalid",
			userPassword:     "password",
			expectedError:    "CLOUDNS_USER_ID is not a valid integer",
			expectedErrorNil: false,
		},
		{
			name:          "valid sub-user login type",
			loginType:     "sub-user",
			subUserID:     "12345",
			userPassword:  "password",
			expectedError: "",
		},
		{
			name:             "invalid sub-user login type",
			loginType:        "sub-user",
			subUserID:        "invalid",
			userPassword:     "password",
			expectedError:    "CLOUDNS_SUB_USER_ID is not a valid integer",
			expectedErrorNil: false,
		},
		{
			name:          "valid sub-user-name login type",
			loginType:     "sub-user-name",
			subUserName:   "user",
			userPassword:  "password",
			expectedError: "",
		},
		{
			name:          "invalid login type",
			loginType:     "invalid",
			userPassword:  "password",
			expectedError: "CLOUDNS_LOGIN_TYPE is not valid",
		},
		{
			name:          "missing user password",
			loginType:     "user-id",
			userID:        "12345",
			expectedError: "CLOUDNS_USER_PASSWORD is not set",
		},
		{
			name:          "missing login type",
			userPassword:  "password",
			expectedError: "CLOUDNS_LOGIN_TYPE is not set",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.loginType != "" {
				os.Setenv("CLOUDNS_LOGIN_TYPE", test.loginType)
			} else {
				os.Unsetenv("CLOUDNS_LOGIN_TYPE")
			}
			if test.userID != "" {
				os.Setenv("CLOUDNS_USER_ID", test.userID)
			}
			if test.subUserID != "" {
				os.Setenv("CLOUDNS_SUB_USER_ID", test.subUserID)
			}
			if test.subUserName != "" {
				os.Setenv("CLOUDNS_SUB_USER_NAME", test.subUserName)
			}
			if test.userPassword != "" {
				os.Setenv("CLOUDNS_USER_PASSWORD", test.userPassword)
			} else {
				os.Unsetenv("CLOUDNS_USER_PASSWORD")
			}

			_, err := NewClouDNSProvider(ClouDNSConfig{})
			if err != nil && test.expectedError == "" {
				t.Errorf("got unexpected error: %s", err)
			} else if err == nil && test.expectedError != "" {
				t.Errorf("expected error %q but got nil", test.expectedError)
			} else if err != nil && test.expectedError != "" && err.Error() != test.expectedError {
				t.Errorf("got error %q, want %q", err.Error(), test.expectedError)
			}
			if err == nil && test.expectedErrorNil {
				t.Errorf("expected error but got nil")
			}
		})
	}
}

// rootZone returns the root zone of a domain name. A root zone is the last two parts
// of a domain name, separated by a "." character. If the domain name has less than two parts,
// the domain name is returned as-is.
// domain is the domain name to be checked.
func TestRootZone(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected string
	}{
		{
			name:     "two-part domain",
			domain:   "easy.com",
			expected: "easy.com",
		},
		{
			name:     "three-part domain",
			domain:   "test.this.program.com",
			expected: "program.com",
		},
		{
			name:     "four-part domain",
			domain:   "something.really.long.com",
			expected: "long.com",
		},
		{
			name:     "one-part domain",
			domain:   "root",
			expected: "root",
		},
		{
			name:     "empty domain",
			domain:   "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := rootZone(test.domain)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}

// removeRootZone removes the given root zone and any trailing periods from the domain name.
// If the root zone is not present in the domain, the domain is returned unmodified.
// domain is the domain name to be modified.
// rootZone is the root zone to be removed from the domain name.
func TestRemoveRootZone(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		rootZone string
		expected string
	}{
		{
			name:     "root zone not present",
			domain:   "www.example.com",
			rootZone: "foo",
			expected: "www.example.com",
		},
		{
			name:     "root zone at end of domain",
			domain:   "www.example.com",
			rootZone: "com",
			expected: "www.example",
		},
		{
			name:     "root zone in middle of domain",
			domain:   "www.example.co.uk",
			rootZone: "co.uk",
			expected: "www.example",
		},
		{
			name:     "root zone spans multiple levels",
			domain:   "www.example.co.uk",
			rootZone: "example.co.uk",
			expected: "www",
		},
		{
			name:     "root zone is entire domain",
			domain:   "example.co.uk",
			rootZone: "example.co.uk",
			expected: "",
		},
		{
			name:     "root zone is root",
			domain:   ".",
			rootZone: ".",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeRootZone(test.domain, test.rootZone)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}

// TestRemoveLastOccurrance tests the removeLastOccurrance function.
// It verifies that the function correctly removes the last occurrence of the given
// substring from the input string, and returns the modified string.
// The test cases cover various scenarios including when the substring is not present
// in the input string, when it is at the start or end of the input string, and when
// there are multiple occurrences of the substring in the input string.
func TestRemoveLastOccurrance(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		subStr   string
		expected string
	}{
		{
			name:     "subdomain not present",
			str:      "www.example.com",
			subStr:   "foo",
			expected: "www.example.com",
		},
		{
			name:     "subdomain at start of domain",
			str:      "www.example.com",
			subStr:   "www",
			expected: ".example.com",
		},
		{
			name:     "subdomain at end of domain",
			str:      "www.example.com",
			subStr:   "com",
			expected: "www.example.",
		},
		{
			name:     "subdomain in middle of domain",
			str:      "www.example.com",
			subStr:   ".e",
			expected: "wwwxample.com",
		},
		{
			name:     "multiple occurrences of subdomain",
			str:      "www.example.com.example.com",
			subStr:   "example",
			expected: "www.example.com..com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeLastOccurrance(test.str, test.subStr)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
