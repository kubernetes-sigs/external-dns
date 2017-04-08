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

package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/digitalocean/godo"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

type mockDigitalOceanInterface interface {
	List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error)
	Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error)
	Delete(context.Context, string) (*godo.Response, error)
}

type mockDigitalOceanClient struct{}

func (m *mockDigitalOceanClient) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanClient) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanClient) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanClient) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanClient) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanClient) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanClient) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanClient) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanClient) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1, Name: "foobar.ext-dns-test.zalando.to."}, {ID: 2}}, nil, nil
}

type mockDigitalOceanListFail struct{}

func (m *mockDigitalOceanListFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{}, nil, fmt.Errorf("Fail to get domains")
}

func (m *mockDigitalOceanListFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanListFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanListFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanListFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanListFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanListFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanListFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanListFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1}, {ID: 2}}, nil, nil
}

type mockDigitalOceanGetFail struct{}

func (m *mockDigitalOceanGetFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanGetFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanGetFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanGetFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanGetFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanGetFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanGetFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to get domain")
}

func (m *mockDigitalOceanGetFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanGetFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1}, {ID: 2}}, nil, nil
}

type mockDigitalOceanCreateFail struct{}

func (m *mockDigitalOceanCreateFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanCreateFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to create domain")
}

func (m *mockDigitalOceanCreateFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanCreateFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanCreateFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanCreateFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanCreateFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanCreateFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanCreateFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1}, {ID: 2}}, nil, nil
}

type mockDigitalOceanDeleteFail struct{}

func (m *mockDigitalOceanDeleteFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, fmt.Errorf("Failed to delete domain")
}
func (m *mockDigitalOceanDeleteFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanDeleteFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1}, {ID: 2}}, nil, nil
}

type mockDigitalOceanRecordsFail struct{}

func (m *mockDigitalOceanRecordsFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanRecordsFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanRecordsFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to get records")
}

func (m *mockDigitalOceanRecordsFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{}, nil, fmt.Errorf("Failed to get records")
}

type mockDigitalOceanUpdateRecordsFail struct{}

func (m *mockDigitalOceanUpdateRecordsFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanUpdateRecordsFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanUpdateRecordsFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanUpdateRecordsFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanUpdateRecordsFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanUpdateRecordsFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to update record")
}

func (m *mockDigitalOceanUpdateRecordsFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanUpdateRecordsFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanUpdateRecordsFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1}, {ID: 2}}, nil, nil
}

type mockDigitalOceanDeleteRecordsFail struct{}

func (m *mockDigitalOceanDeleteRecordsFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanDeleteRecordsFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, fmt.Errorf("Failed to delete record")
}
func (m *mockDigitalOceanDeleteRecordsFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanDeleteRecordsFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1, Name: "foobar.ext-dns-test.zalando.to."}, {ID: 2}}, nil, nil
}

type mockDigitalOceanCreateRecordsFail struct{}

func (m *mockDigitalOceanCreateRecordsFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanCreateRecordsFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanCreateRecordsFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to create record")
}

func (m *mockDigitalOceanCreateRecordsFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanCreateRecordsFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanCreateRecordsFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanCreateRecordsFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanCreateRecordsFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanCreateRecordsFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{{ID: 1, Name: "foobar.ext-dns-test.zalando.to."}, {ID: 2}}, nil, nil
}

func TestDigitalOceanZones(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	_, err := provider.Zones()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	provider.Client = &mockDigitalOceanListFail{}
	_, err = provider.Zones()
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestDigitalOceanZone(t *testing.T) {
	zoneName := "example.com"

	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	domain, err := provider.Zone(zoneName)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	if domain.Name != zoneName {
		t.Errorf("expected %s, got %s", zoneName, domain.Name)
	}

	provider.Client = &mockDigitalOceanGetFail{}
	_, err = provider.Zone(zoneName)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestDigitalOceanCreateZone(t *testing.T) {
	zoneName := "example.com"
	ip := "1.2.3.4"

	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	domain, err := provider.CreateZone(zoneName, ip)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	if domain.Name != zoneName {
		t.Errorf("expected %s, got %s", zoneName, domain.Name)
	}

	provider.Client = &mockDigitalOceanCreateFail{}
	_, err = provider.CreateZone(zoneName, ip)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestDigitalOceanDeleteZone(t *testing.T) {
	zoneName := "example.com"

	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	_, err := provider.DeleteZone(zoneName)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	provider.Client = &mockDigitalOceanDeleteFail{}
	_, err = provider.DeleteZone(zoneName)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}

}

func TestDigitalOceanRecords(t *testing.T) {
	zoneName := "example.com"

	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	_, err := provider.Records(zoneName)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockDigitalOceanRecordsFail{}
	_, err = provider.Records(zoneName)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestDigitalOceanUpdateRecords(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	zone := "ext-dns-test.zalando.to"
	oldCNameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "old-target"}}
	newCNameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "new-target"}}

	oldANameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "8.8.8.8"}}
	newANameEndpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "7.7.7.7"}}
	err := provider.UpdateRecords(zone, newCNameEndpoints, oldCNameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockDigitalOceanUpdateRecordsFail{}
	err = provider.UpdateRecords(zone, newCNameEndpoints, oldCNameEndpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockDigitalOceanRecordsFail{}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockDigitalOceanClient{}
	provider.DryRun = true
	err = provider.UpdateRecords(zone, newCNameEndpoints, oldCNameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	err = provider.UpdateRecords(zone, newANameEndpoints, oldANameEndpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestDigitalOceanDeleteRecords(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	zone := "ext-dns-test.zalando.to"
	endpoints := []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target"}}
	err := provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockDigitalOceanDeleteRecordsFail{}
	err = provider.DeleteRecords(zone, endpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.DryRun = true
	err = provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockDigitalOceanDeleteRecordsFail{}
	provider.DryRun = false
	err = provider.DeleteRecords(zone, endpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	endpoints = []*endpoint.Endpoint{}
	err = provider.DeleteRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestDigitalOceanCreateRecords(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	zone := "ext-dns-test.zalando.to"
	endpoints := []*endpoint.Endpoint{
		{DNSName: "new", Target: "target"},
	}
	err := provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.DryRun = true
	err = provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockDigitalOceanCreateRecordsFail{}
	provider.DryRun = false
	err = provider.CreateRecords(zone, endpoints)
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockDigitalOceanClient{}
	endpoints = []*endpoint.Endpoint{}
	err = provider.CreateRecords(zone, endpoints)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestDigitalOceanApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	zone := "ext-dns-test.zalando.to"
	changes.Create = []*endpoint.Endpoint{{DNSName: "new.ext-dns-test.zalando.to.", Target: "target"}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target"}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-old"}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-new"}}
	err := provider.ApplyChanges(zone, changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestNewDigitalOceanProvider(t *testing.T) {
	_ = os.Setenv("DO_TOKEN", "xxxxxxxxxxxxxxxxx")
	_, err := NewDigitalOceanProvider(true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("DO_TOKEN")
	_, err = NewDigitalOceanProvider(true)
	fmt.Println(err)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestToken(t *testing.T) {
	ts := &TokenSource{AccessToken: "xxxxxxxxxxxxxxxxx"}
	oauthToken, err := ts.Token()
	if err != nil {
		t.Errorf("should not fail")
	}
	if oauthToken.AccessToken != ts.AccessToken {
		t.Errorf("Expected %s, got %s", ts.AccessToken, oauthToken.AccessToken)
	}

}
