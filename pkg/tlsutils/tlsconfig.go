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

package tlsutils

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const defaultMinVersion = 0

// CreateTLSConfig creates tls.Config instance from TLS parameters passed in environment variables with the given prefix
func CreateTLSConfig(prefix string) (*tls.Config, error) {
	caFile := os.Getenv(fmt.Sprintf("%s_CA_FILE", prefix))
	certFile := os.Getenv(fmt.Sprintf("%s_CERT_FILE", prefix))
	keyFile := os.Getenv(fmt.Sprintf("%s_KEY_FILE", prefix))
	serverName := os.Getenv(fmt.Sprintf("%s_TLS_SERVER_NAME", prefix))
	isInsecureStr := strings.ToLower(os.Getenv(fmt.Sprintf("%s_TLS_INSECURE", prefix)))
	isInsecure := isInsecureStr == "true" || isInsecureStr == "yes" || isInsecureStr == "1"
	tlsConfig, err := NewTLSConfig(certFile, keyFile, caFile, serverName, isInsecure, defaultMinVersion)
	if err != nil {
		return nil, err
	}
	return tlsConfig, nil
}

// NewTLSConfig creates a tls.Config instance from directly-passed parameters, loading the ca, cert, and key from disk
func NewTLSConfig(certPath, keyPath, caPath, serverName string, insecure bool, minVersion uint16) (*tls.Config, error) {
	if certPath != "" && keyPath == "" || certPath == "" && keyPath != "" {
		return nil, errors.New("either both cert and key or none must be provided")
	}
	var certificates []tls.Certificate
	if certPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, fmt.Errorf("could not load TLS cert: %s", err)
		}
		certificates = append(certificates, cert)
	}
	roots, err := loadRoots(caPath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		MinVersion:         minVersion,
		Certificates:       certificates,
		RootCAs:            roots,
		InsecureSkipVerify: insecure,
		ServerName:         serverName,
	}, nil
}

// loads CA cert
func loadRoots(caPath string) (*x509.CertPool, error) {
	if caPath == "" {
		return nil, nil
	}

	roots := x509.NewCertPool()
	pem, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %s", caPath, err)
	}
	ok := roots.AppendCertsFromPEM(pem)
	if !ok {
		return nil, fmt.Errorf("could not read root certs: %s", err)
	}
	return roots, nil
}
