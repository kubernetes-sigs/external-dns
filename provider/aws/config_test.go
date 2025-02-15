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
  "crypto/ecdsa"
  "crypto/elliptic"
  "crypto/rand"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/pem"
  "math/big"
  "os"
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
)

func Test_newV2Config(t *testing.T) {
  t.Run("should use profile from credentials file", func(t *testing.T) {
    // setup
    credsFile, err := prepareCredentialsFile(t)
    defer os.Remove(credsFile.Name())
    require.NoError(t, err)
    os.Setenv("AWS_SHARED_CREDENTIALS_FILE", credsFile.Name())
    defer os.Unsetenv("AWS_SHARED_CREDENTIALS_FILE")

    // when
    cfg, err := newV2Config(AWSSessionConfig{Profile: "profile2"})
    require.NoError(t, err)
    creds, err := cfg.Credentials.Retrieve(context.Background())

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

  t.Run("should use tls config", func(t *testing.T) {
    os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
    os.Setenv("AWS_SECRET_ACCESS_KEY", "topsecret")
    defer os.Unsetenv("AWS_ACCESS_KEY_ID")
    defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

    certPEMMap, err := generateSelfSignedCert(t)
    require.NoError(t, err)

    cfg, err := newV2Config(AWSSessionConfig{
      TLSCertPath:    certPEMMap["clientCertPath"],
      TLSCertKeyPath: certPEMMap["clientKeyPath"],
      TLSCAFilePath:  certPEMMap["rootCAPath"],
    })
    require.NoError(t, err)
    creds, err := cfg.Credentials.Retrieve(context.Background())
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

func certTemplate(keyUsage x509.KeyUsage, extKeyUsage []x509.ExtKeyUsage, isCA bool) (x509.Certificate, error) {
  serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
  serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
  if err != nil {
    return x509.Certificate{}, err
  }

  return x509.Certificate{
    SerialNumber:          serialNumber,
    NotBefore:             time.Now(),
    NotAfter:              time.Now().Add(time.Hour),
    IsCA:                  isCA,
    BasicConstraintsValid: true,
    KeyUsage:              keyUsage,
    ExtKeyUsage:           extKeyUsage,
    Subject: pkix.Name{
      CommonName:   "Test CA",
      Organization: []string{"Test Org"},
    },
  }, nil
}

func generateSelfSignedCert(t *testing.T) (map[string]string, error) {
  t.Helper()
  rootKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil {
    return nil, err
  }

  rootCertTemplate, err := certTemplate(x509.KeyUsageKeyEncipherment|x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, true)
  if err != nil {
    return nil, err
  }

  rootCertDER, err := x509.CreateCertificate(rand.Reader, &rootCertTemplate, &rootCertTemplate, &rootKey.PublicKey, rootKey)
  if err != nil {
    return nil, err
  }

  rootCert, err := x509.ParseCertificate(rootCertDER)
  if err != nil {
    return nil, err
  }

  certTemplate, err := certTemplate(x509.KeyUsageDigitalSignature, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, false)
  if err != nil {
    return nil, err
  }
  certKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil {
    return nil, err
  }

  certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, rootCert, &certKey.PublicKey, rootKey)
  if err != nil {
    return nil, err
  }

  certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

  keyPKCS1, err := x509.MarshalECPrivateKey(certKey)
  if err != nil {
    return nil, err
  }
  keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: keyPKCS1})

  rootCAPath := fileToTempLocation(t, "rootCA*.pem", rootCertDER)
  clientCertPath := fileToTempLocation(t, "clientCert*.pem", certPEM)
  clientKeyPath := fileToTempLocation(t, "clientKey*.pem", keyPEM)

  return map[string]string{
    "rootCAPath":     rootCAPath,
    "clientCertPath": clientCertPath,
    "clientKeyPath":  clientKeyPath,
  }, nil
}

func fileToTempLocation(t *testing.T, pattern string, data []byte) string {
  t.Helper()
  tmpFile, err := os.CreateTemp("", pattern)
  if err != nil {
    t.Fatalf("failed to create temp file: %v", err)
  }

  defer func(tmpFile *os.File) {
    err := tmpFile.Close()
    if err != nil {
      t.Fatalf("failed to close temp file: %v", err)
    }
  }(tmpFile)

  if _, err := tmpFile.Write(data); err != nil {
    t.Fatalf("failed to write data to temp file: %v", err)
  }

  return tmpFile.Name()
}
