/*
Copyright 2022 The Kubernetes Authors.

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

package plural

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pluralsh/gqlclient"
	"github.com/pluralsh/gqlclient/pkg/utils"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

type Client interface {
	DnsRecords() ([]*DnsRecord, error)
	CreateRecord(record *DnsRecord) (*DnsRecord, error)
	DeleteRecord(name, ttype string) error
}

type Config struct {
	Token    string
	Endpoint string
	Cluster  string
	Provider string
}

type client struct {
	ctx          context.Context
	pluralClient *gqlclient.Client
	config       *Config
}

type DnsRecord struct {
	Type    string
	Name    string
	Records []string
}

func NewClient(conf *Config) (Client, error) {
	base, err := conf.BaseUrl()
	if err != nil {
		return nil, err
	}
	httpClient := http.Client{
		Transport: &authedTransport{
			key:     conf.Token,
			wrapped: http.DefaultTransport,
		},
	}
	endpoint := base + "/gql"
	return &client{
		ctx:          context.Background(),
		pluralClient: gqlclient.NewClient(&httpClient, endpoint),
		config:       conf,
	}, nil
}

func (c *Config) BaseUrl() (string, error) {
	host := "https://app.plural.sh"
	if c.Endpoint != "" {
		host = fmt.Sprintf("https://%s", c.Endpoint)
		if _, err := url.Parse(host); err != nil {
			return "", err
		}
	}
	return host, nil
}

func (client *client) DnsRecords() ([]*DnsRecord, error) {
	resp, err := client.pluralClient.GetDNSRecords(client.ctx, client.config.Cluster, gqlclient.Provider(strings.ToUpper(client.config.Provider)))
	if err != nil {
		return nil, err
	}

	records := make([]*DnsRecord, 0)
	for _, edge := range resp.DNSRecords.Edges {
		if edge.Node != nil {
			record := &DnsRecord{
				Type:    string(edge.Node.Type),
				Name:    edge.Node.Name,
				Records: utils.ConvertStringArrayPointer(edge.Node.Records),
			}
			records = append(records, record)
		}
	}
	return records, nil
}

func (client *client) CreateRecord(record *DnsRecord) (*DnsRecord, error) {
	provider := gqlclient.Provider(strings.ToUpper(client.config.Provider))
	cluster := client.config.Cluster
	attr := gqlclient.DNSRecordAttributes{
		Name:    record.Name,
		Type:    gqlclient.DNSRecordType(record.Type),
		Records: []*string{},
	}

	for _, record := range record.Records {
		attr.Records = append(attr.Records, &record)
	}

	resp, err := client.pluralClient.CreateDNSRecord(client.ctx, cluster, provider, attr)
	if err != nil {
		return nil, err
	}

	return &DnsRecord{
		Type:    string(resp.CreateDNSRecord.Type),
		Name:    resp.CreateDNSRecord.Name,
		Records: utils.ConvertStringArrayPointer(resp.CreateDNSRecord.Records),
	}, nil
}

func (client *client) DeleteRecord(name, ttype string) error {
	if _, err := client.pluralClient.DeleteDNSRecord(client.ctx, name, gqlclient.DNSRecordType(ttype)); err != nil {
		return err
	}

	return nil
}
