/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"fmt"
	"net/http"
	"os"
)

// ClientConfiguration represents the vinyldns client configuration.
type ClientConfiguration struct {
	AccessKey string
	SecretKey string
	Host      string
	UserAgent string
}

// NewConfigFromEnv creates a new ClientConfiguration
// using environment variables.
func NewConfigFromEnv() ClientConfiguration {
	ua := defaultUA()

	if vua := os.Getenv("VINYLDNS_USER_AGENT"); vua != "" {
		ua = vua
	}
	return ClientConfiguration{
		os.Getenv("VINYLDNS_ACCESS_KEY"),
		os.Getenv("VINYLDNS_SECRET_KEY"),
		os.Getenv("VINYLDNS_HOST"),
		ua,
	}
}

// Client is a vinyldns API client.
type Client struct {
	AccessKey  string
	SecretKey  string
	Host       string
	HTTPClient *http.Client
	UserAgent  string
}

// NewClientFromEnv returns a Client configured via
// environment variables.
func NewClientFromEnv() *Client {
	return NewClient(NewConfigFromEnv())
}

// NewClient returns a new vinyldns client using
// the client ClientConfiguration it's passed.
func NewClient(config ClientConfiguration) *Client {
	if config.UserAgent == "" {
		config.UserAgent = defaultUA()
	}

	return &Client{
		config.AccessKey,
		config.SecretKey,
		config.Host,
		&http.Client{},
		config.UserAgent,
	}
}

func defaultUA() string {
	return fmt.Sprintf("go-vinyldns/%s", Version)
}

func logRequests() bool {
	return os.Getenv("VINYLDNS_LOG") != ""
}
