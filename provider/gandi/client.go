/*
Copyright 2021 The Kubernetes Authors.
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

package gandi

import (
	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
)

type DomainClientAdapter interface {
	ListDomains() (domains []domain.ListResponse, err error)
}

type domainClient struct {
	Client *domain.Domain
}

func (p *domainClient) ListDomains() (domains []domain.ListResponse, err error) {
	return p.Client.ListDomains()
}

func NewDomainClient(client *domain.Domain) DomainClientAdapter {
	return &domainClient{client}
}

// standardResponse copied from go-gandi/internal/gandi.go
type standardResponse struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	UUID    string          `json:"uuid,omitempty"`
	Object  string          `json:"object,omitempty"`
	Cause   string          `json:"cause,omitempty"`
	Status  string          `json:"status,omitempty"`
	Errors  []standardError `json:"errors,omitempty"`
}

// standardError copied from go-gandi/internal/gandi.go
type standardError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LiveDNSClientAdapter interface {
	GetDomainRecords(fqdn string) (records []livedns.DomainRecord, err error)
	CreateDomainRecord(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error)
	DeleteDomainRecord(fqdn, name, recordtype string) (err error)
	UpdateDomainRecordByNameAndType(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error)
}

type LiveDNSClient struct {
	Client *livedns.LiveDNS
}

func NewLiveDNSClient(client *livedns.LiveDNS) LiveDNSClientAdapter {
	return &LiveDNSClient{client}
}

func (p *LiveDNSClient) GetDomainRecords(fqdn string) (records []livedns.DomainRecord, err error) {
	return p.Client.GetDomainRecords(fqdn)
}

func (p *LiveDNSClient) CreateDomainRecord(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error) {
	res, err := p.Client.CreateDomainRecord(fqdn, name, recordtype, ttl, values)
	if err != nil {
		return standardResponse{}, err
	}

	// response needs to be copied as the Standard* structs are internal
	var errors []standardError
	for _, e := range res.Errors {
		errors = append(errors, standardError(e))
	}
	return standardResponse{
		Code:    res.Code,
		Message: res.Message,
		UUID:    res.UUID,
		Object:  res.Object,
		Cause:   res.Cause,
		Status:  res.Status,
		Errors:  errors,
	}, err
}

func (p *LiveDNSClient) DeleteDomainRecord(fqdn, name, recordtype string) (err error) {
	return p.Client.DeleteDomainRecord(fqdn, name, recordtype)
}

func (p *LiveDNSClient) UpdateDomainRecordByNameAndType(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error) {
	res, err := p.Client.UpdateDomainRecordByNameAndType(fqdn, name, recordtype, ttl, values)
	if err != nil {
		return standardResponse{}, err
	}

	// response needs to be copied as the Standard* structs are internal
	var errors []standardError
	for _, e := range res.Errors {
		errors = append(errors, standardError(e))
	}
	return standardResponse{
		Code:    res.Code,
		Message: res.Message,
		UUID:    res.UUID,
		Object:  res.Object,
		Cause:   res.Cause,
		Status:  res.Status,
		Errors:  errors,
	}, err
}
