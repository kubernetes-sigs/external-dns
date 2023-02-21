package infoblox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInfobloxProvider(t *testing.T) {
	defaultConfig := StartupConfig{
		// this is the default configuration
		// ... and it is not ready to be used without authentication parameters and
		// ... without defining a host to communicate to.
		Host:          "",
		Port:          443,
		Username:      "",
		Password:      "",
		Version:       "2.3.1",
		SSLVerify:     true,
		ClientCert:    "",
		ClientKey:     "",
		HostRO:        "",
		PortRO:        443,
		UsernameRO:    "",
		PasswordRO:    "",
		VersionRO:     "2.3.1",
		SSLVerifyRO:   true,
		ClientCertRO:  "",
		ClientKeyRO:   "",
		View:          "default",
		MaxResults:    0,
		FQDNRegEx:     "",
		CreatePTR:     false,
		CacheDuration: 0,
	}

	testCfg := defaultConfig
	newProvider, err := NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "either username AND password or certificate AND key MUST be specified", err.Error())

	testCfg = StartupConfig{
		Host:          "localhost",
		Port:          444,
		Username:      "user1",
		Password:      "pass1",
		Version:       "2.11",
		SSLVerify:     true,
		ClientCert:    "",
		ClientKey:     "",
		HostRO:        "127.0.0.1",
		PortRO:        445,
		UsernameRO:    "user2",
		PasswordRO:    "pass2",
		VersionRO:     "2.11.1",
		SSLVerifyRO:   false,
		ClientCertRO:  "",
		ClientKeyRO:   "",
		View:          "non_default",
		MaxResults:    10,
		FQDNRegEx:     ".+\\.test\\.com",
		CreatePTR:     true,
		CacheDuration: 10,
	}
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)
	assert.Equal(t, "non_default", newProvider.view)
	assert.Equal(t, 10, newProvider.cacheDuration)
	assert.Equal(t, true, newProvider.createPTR)
	assert.NotNil(t, newProvider.fqdnRegEx)
	assert.NotNil(t, newProvider.clientRW)
	assert.NotNil(t, newProvider.clientRO)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.Username = "user1"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "either username AND password or certificate AND key MUST be specified", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "either username AND password or certificate AND key MUST be specified", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.PasswordRO = "pass2"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "either username AND password or certificate AND key MUST be specified", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "either username AND password or certificate AND key MUST be specified", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "to use certificate-based authentication you must specify BOTH certificate's and private key's file paths", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientKeyRO = "test-artifacts/ro/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "to use certificate-based authentication you must specify BOTH certificate's and private key's file paths", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "to use certificate-based authentication you must specify BOTH certificate's and private key's file paths", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	testCfg.ClientCertRO = "test-artifacts/ro/client.cert.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "to use certificate-based authentication you must specify BOTH certificate's and private key's file paths", err.Error())

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKeyRO = "test-artifacts/ro/client.key.pem"
	testCfg.ClientCertRO = "test-artifacts/ro/client.cert.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.Nil(t, newProvider)
	assert.NotNil(t, err)
	assert.Equal(t, "to use certificate-based authentication you must specify BOTH certificate's and private key's file paths", err.Error())

	// good cases

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.HostRO = "127.0.0.1"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	testCfg.ClientKeyRO = "test-artifacts/ro/client.key.pem"
	testCfg.ClientCertRO = "test-artifacts/ro/client.cert.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.Username = "user2"
	testCfg.Password = "pass2"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.UsernameRO = "user2"
	testCfg.PasswordRO = "pass2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.UsernameRO = "user2"
	testCfg.ClientCert = "test-artifacts/rw/client.cert.pem"
	testCfg.ClientKey = "test-artifacts/rw/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)

	testCfg = defaultConfig
	testCfg.Host = "localhost"
	testCfg.Username = "user1"
	testCfg.Password = "pass1"
	testCfg.ClientKeyRO = "test-artifacts/rw/client.key.pem"
	newProvider, err = NewInfobloxProvider(testCfg)
	assert.NotNil(t, newProvider)
	assert.Nil(t, err)
}
