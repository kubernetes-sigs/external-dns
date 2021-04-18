/*
Copyright 2020 The Kubernetes Authors.

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

package dynsoap

import (
	"net/http"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// NewDynectClient returns a client with a configured http.Client
// The default settings for the http.client are a timeout of
// 10 seconds and reading proxy variables from http.ProxyFromEnvironment
func NewDynectClient(url string) Dynect {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	soapClient := soap.NewClient(url, soap.WithHTTPClient(client))
	return NewDynect(soapClient)
}

// NewCustomDynectClient returns a client without a configured http.Client
func NewCustomDynectClient(url string, client http.Client) Dynect {
	soapClient := soap.NewClient(url, soap.WithHTTPClient(&client))
	return NewDynect(soapClient)
}
