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
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newV2Config(t *testing.T) {
	t.Run("should use profile from credentials file", func(t *testing.T) {
		// setup
		dir := t.TempDir()
		credsFile := filepath.Join(dir, "credentials")
		err := os.WriteFile(credsFile, []byte(`
[profile1]
aws_access_key_id=AKID1234
aws_secret_access_key=SECRET1

[profile2]
aws_access_key_id=AKID2345
aws_secret_access_key=SECRET2
`), 0777)
		require.NoError(t, err)

		t.Setenv("AWS_SHARED_CREDENTIALS_FILE", credsFile)

		// when
		cfg, err := newV2Config(AWSSessionConfig{Profile: "profile2"})
		require.NoError(t, err)
		creds, err := cfg.Credentials.Retrieve(context.Background())

		// then
		assert.NoError(t, err)
		assert.Equal(t, "AKID2345", creds.AccessKeyID)
		assert.Equal(t, "SECRET2", creds.SecretAccessKey)
	})

	t.Run("should respect updates to the credentials file", func(t *testing.T) {
		// setup
		dir := t.TempDir()
		credsFile := filepath.Join(dir, "credentials")
		err := os.WriteFile(credsFile, []byte(`
[default]
aws_access_key_id=AKID1234
aws_secret_access_key=SECRET1
`), 0777)
		require.NoError(t, err)

		t.Setenv("AWS_SHARED_CREDENTIALS_FILE", credsFile)

		cfg, err := newV2Config(AWSSessionConfig{})
		require.NoError(t, err)
		creds, err := cfg.Credentials.Retrieve(context.Background())
		require.NoError(t, err)

		assert.Equal(t, "AKID1234", creds.AccessKeyID)
		assert.Equal(t, "SECRET1", creds.SecretAccessKey)

		// given
		err = os.WriteFile(credsFile, []byte(`
[default]
aws_access_key_id=AKID2345
aws_secret_access_key=SECRET2
`), 0777)
		require.NoError(t, err)

		// when
		creds, err = cfg.Credentials.Retrieve(context.Background())

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
		cfg, err := newV2Config(AWSSessionConfig{})
		require.NoError(t, err)
		creds, err := cfg.Credentials.Retrieve(context.Background())

		// then
		assert.NoError(t, err)
		assert.Equal(t, "AKIAIOSFODNN7EXAMPLE", creds.AccessKeyID)
		assert.Equal(t, "topsecret", creds.SecretAccessKey)
	})
}
