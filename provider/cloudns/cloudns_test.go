package cloudns

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	cloudns "github.com/ppmathis/cloudns-go"
	"sigs.k8s.io/external-dns/endpoint"
)

//var mockProvider = &ClouDNSProvider{}

var mockZones = []cloudns.Zone{
	{
		Name:     "test1.com",
		Type:     1,
		Kind:     1,
		IsActive: true,
	},
	{
		Name:     "test2.com",
		Type:     1,
		Kind:     1,
		IsActive: true,
	},
}

var mockRecords = [][]cloudns.Record{
	{
		{
			ID:         1,
			Host:       "",
			Record:     "1.1.1.1",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         2,
			Host:       "sub2",
			Record:     "2.2.2.2",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         3,
			Host:       "sub3",
			Record:     "3.3.3.3",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         4,
			Host:       "",
			Record:     "TextRecord",
			RecordType: "TXT",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         5,
			Host:       "sub5",
			Record:     "SubTextRecord",
			RecordType: "TXT",
			TTL:        60,
			IsActive:   true,
		},
	},
	{
		{
			ID:         6,
			Host:       "",
			Record:     "6.6.6.6",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         7,
			Host:       "sub7",
			Record:     "7.7.7.7",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         8,
			Host:       "sub8",
			Record:     "8.8.8.8",
			RecordType: "A",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         9,
			Host:       "",
			Record:     "TextRecord",
			RecordType: "TXT",
			TTL:        60,
			IsActive:   true,
		},
		{
			ID:         10,
			Host:       "sub5",
			Record:     "SubTextRecord",
			RecordType: "TXT",
			TTL:        60,
			IsActive:   true,
		},
	},
}

var expectedEndpointsOne = []*endpoint.Endpoint{
	// endpoint 1
	endpoint.NewEndpointWithTTL(
		"test1.com",
		"A",
		endpoint.TTL(60),
		"1.1.1.1",
	),
	// endpoint 2
	endpoint.NewEndpointWithTTL(
		"sub2.test1.com",
		"A",
		endpoint.TTL(60),
		"2.2.2.2",
	),
	// endpoint 3
	endpoint.NewEndpointWithTTL(
		"sub3.test1.com",
		"A",
		endpoint.TTL(60),
		"3.3.3.3",
	),
	// endpoint 4
	endpoint.NewEndpointWithTTL(
		"test1.com",
		"TXT",
		endpoint.TTL(60),
		"TextRecord",
	),
	// endpoint 5
	endpoint.NewEndpointWithTTL(
		"sub5.test1.com",
		"TXT",
		endpoint.TTL(60),
		"SubTextRecord",
	),
	// endpoint 6
	endpoint.NewEndpointWithTTL(
		"test2.com",
		"A",
		endpoint.TTL(60),
		"6.6.6.6",
	),
	// endpoint 7
	endpoint.NewEndpointWithTTL(
		"sub7.test2.com",
		"A",
		endpoint.TTL(60),
		"7.7.7.7",
	),
	// endpoint 8
	endpoint.NewEndpointWithTTL(
		"sub8.test2.com",
		"A",
		endpoint.TTL(60),
		"8.8.8.8",
	),
	// endpoint 9
	endpoint.NewEndpointWithTTL(
		"test2.com",
		"TXT",
		endpoint.TTL(60),
		"TextRecord",
	),
	// endpoint 10
	endpoint.NewEndpointWithTTL(
		"sub5.test2.com",
		"TXT",
		endpoint.TTL(60),
		"SubTextRecord",
	),
}

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

//func Test_Records(t *testing.T)

func TestIsValidTTL(t *testing.T) {
	// Test valid TTLs
	if !isValidTTL("60") {
		t.Error("Expected isValidTTL to return true for TTL 60, got false")
	}
	if !isValidTTL("3600") {
		t.Error("Expected isValidTTL to return true for TTL 3600, got false")
	}
	if !isValidTTL("1209600") {
		t.Error("Expected isValidTTL to return true for TTL 1209600, got false")
	}

	// Test invalid TTLs
	if isValidTTL("0") {
		t.Error("Expected isValidTTL to return false for TTL 0, got true")
	}
	if isValidTTL("300000") {
		t.Error("Expected isValidTTL to return false for TTL 300000, got true")
	}
	if isValidTTL("abc") {
		t.Error("Expected isValidTTL to return false for TTL abc, got true")
	}
}

func TestZoneRecordMap(t *testing.T) {

	zoneOneRecordMap := make(cloudns.RecordMap)
	for _, record := range mockRecords[0] {
		zoneOneRecordMap[record.ID] = record
	}

	oneZoneRecordMap := make(map[string]cloudns.RecordMap)
	oneZoneRecordMap["test1.com"] = zoneOneRecordMap

	zoneTwoRecordMap := make(cloudns.RecordMap)
	for _, record := range mockRecords[1] {
		zoneTwoRecordMap[record.ID] = record
	}

	twoZoneRecordMap := make(map[string]cloudns.RecordMap)
	twoZoneRecordMap["test1.com"] = zoneOneRecordMap
	twoZoneRecordMap["test2.com"] = zoneTwoRecordMap

	tests := []struct {
		name           string
		expectedMap    map[string]cloudns.RecordMap
		expectingError bool
		mockFunc       func()
	}{
		{
			name:           "no zones",
			expectedMap:    map[string]cloudns.RecordMap{},
			expectingError: false,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return []cloudns.Zone{}, nil
				}

				listRecords = func(client *cloudns.Client, ctx context.Context, zoneName string) (cloudns.RecordMap, error) {
					return nil, nil
				}
			},
		},
		{
			name:           "no records",
			expectedMap:    map[string]cloudns.RecordMap{},
			expectingError: false,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return mockZones, nil
				}

				listRecords = func(client *cloudns.Client, ctx context.Context, zoneName string) (cloudns.RecordMap, error) {
					return nil, nil
				}
			},
		},
		{
			name:           "list zones error",
			expectedMap:    nil,
			expectingError: true,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return nil, fmt.Errorf("list zones error")
				}
			},
		},
		{
			name:           "list records error",
			expectedMap:    nil,
			expectingError: true,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return mockZones, nil
				}

				listRecords = func(client *cloudns.Client, ctx context.Context, zoneName string) (cloudns.RecordMap, error) {
					return nil, fmt.Errorf("list records error")
				}
			},
		},
		{
			name:           "one zone, five records",
			expectedMap:    oneZoneRecordMap,
			expectingError: false,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return mockZones[0:1], nil
				}

				listRecords = func(client *cloudns.Client, ctx context.Context, zoneName string) (cloudns.RecordMap, error) {
					return zoneOneRecordMap, nil
				}
			},
		},
		{
			name:           "two zones, ten records",
			expectedMap:    twoZoneRecordMap,
			expectingError: false,
			mockFunc: func() {
				listZones = func(client *cloudns.Client, ctx context.Context) ([]cloudns.Zone, error) {
					return mockZones, nil
				}

				listRecords = func(client *cloudns.Client, ctx context.Context, zoneName string) (cloudns.RecordMap, error) {
					if zoneName == "test1.com" {
						return zoneOneRecordMap, nil
					}
					if zoneName == "test2.com" {
						return zoneTwoRecordMap, nil
					}
					return nil, nil
				}
			},
		},
	}

	oriListZones := listZones
	oriListRecords := listRecords

	provider := &ClouDNSProvider{}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			test.mockFunc()
			zoneRecordMap, err := provider.zoneRecordMap(context.Background())

			errExist := err != nil
			if test.expectingError != errExist {
				tt.Errorf("Expected error: %v, got: %v", test.expectingError, errExist)
			}

			if !reflect.DeepEqual(test.expectedMap, zoneRecordMap) {
				tt.Errorf("Error, return value expectation. Want: %+v, got: %+v", test.expectedMap, zoneRecordMap)
			}
		})
	}

	listZones = oriListZones
	listRecords = oriListRecords
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
