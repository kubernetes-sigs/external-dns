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

package aws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newSession(t *testing.T) {
	t.Run("should use profile from credentials file", func(t *testing.T) {
		// setup
		credsFile, err := prepareCredentialsFile(t)
		defer os.Remove(credsFile.Name())
		require.NoError(t, err)
		os.Setenv("AWS_SHARED_CREDENTIALS_FILE", credsFile.Name())
		defer os.Unsetenv("AWS_SHARED_CREDENTIALS_FILE")

		// when
		s, err := newSession(AWSSessionConfig{Profile: "profile2"})
		require.NoError(t, err)
		creds, err := s.Config.Credentials.Get()

		// then
		assert.NoError(t, err)
		assert.Equal(t, "AKID2345", creds.AccessKeyID)
		assert.Equal(t, "SECRET2", creds.SecretAccessKey)
	})

	t.Run("should respect env variables without profile", func(t *testing.T) {
		// setup
		os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "topsecret")
		defer os.Unsetenv("AWS_ACCESS_KEY_ID")
		defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

		// when
		s, err := newSession(AWSSessionConfig{})
		require.NoError(t, err)
		creds, err := s.Config.Credentials.Get()

		// then
		assert.NoError(t, err)
		assert.Equal(t, "AKIAIOSFODNN7EXAMPLE", creds.AccessKeyID)
		assert.Equal(t, "topsecret", creds.SecretAccessKey)
	})
}

func prepareCredentialsFile(t *testing.T) (*os.File, error) {
	credsFile, err := os.CreateTemp("", "aws-*.creds")
	require.NoError(t, err)
	_, err = credsFile.WriteString("[profile1]\naws_access_key_id=AKID1234\naws_secret_access_key=SECRET1\n\n[profile2]\naws_access_key_id=AKID2345\naws_secret_access_key=SECRET2\n")
	require.NoError(t, err)
	err = credsFile.Close()
	require.NoError(t, err)
	return credsFile, err
}
