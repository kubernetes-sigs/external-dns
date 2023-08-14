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

package domeneshop

import "net/http"

// Domain JSON data structure
type Domain struct {
	Name           string   `json:"domain"`
	ID             int      `json:"id"`
	ExpiryDate     string   `json:"expiry_date"`
	Nameservers    []string `json:"nameservers"`
	RegisteredDate string   `json:"registered_date"`
	Registrant     string   `json:"registrant"`
	Renew          bool     `json:"renew"`
	Services       struct {
		DNS       bool   `json:"dns"`
		Email     bool   `json:"email"`
		Registrar bool   `json:"registrar"`
		Webhotel  string `json:"webhotel"`
	} `json:"services"`
	Status string
}

type HttpErrorBody struct {
	Code string `json:"code"`
	Help string `json:"help"`
}

type HttpError struct {
	// TODO error wrap
	Message   error
	Response  *http.Response
	ErrorBody HttpErrorBody
}

func (e *HttpError) Error() string {
	return e.Message.Error()
}

// DNSRecord JSON data structure
type DNSRecord struct {
	Data string `json:"data,omitempty"`
	Host string `json:"host,omitempty"`
	ID   int    `json:"id,omitempty"`
	TTL  int    `json:"ttl,omitempty"`
	Type string `json:"type,omitempty"`
}
