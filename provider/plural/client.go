package plural

import (
	"context"
	"fmt"
	"strings"

	"github.com/michaeljguarino/graphql"
)

type Config struct {
	Token    string
	Endpoint string
	Cluster  string
	Provider string
}

type Client struct {
	gqlClient *graphql.Client
	config    *Config
}

type DnsRecord struct {
	Type    string
	Name    string
	Records []string
}

const DnsRecordFragment = `
	fragment DnsRecord on DnsRecord {
		type
		name
		records
	}
`

var dnsRecordsQuery = fmt.Sprintf(`
	query DnsRecords($cluster: String!, $provider: Provider!) {
		dnsRecords(cluster: $cluster, provider: $provider, first: 500) {
			edges { node { ...DnsRecord } }
		}
	}
	%s
`, DnsRecordFragment)

var createDnsRecord = fmt.Sprintf(`
	mutation Create($cluster: String!, $provider: Provider!, $attributes: DnsRecordAttributes!) {
		createDnsRecord(cluster: $cluster, provider: $provider, attributes: $attributes) {
			...DnsRecord
		}
	}
	%s
`, DnsRecordFragment)

var deleteDnsRecord = fmt.Sprintf(`
	mutation Delete($name: String!, $type: DnsRecordType!) {
		deleteDnsRecord(name: $name, type: $type) {
			...DnsRecord
		}
	}
	%s
`, DnsRecordFragment)

func NewClient(conf *Config) *Client {
	base := conf.BaseUrl()
	return &Client{graphql.NewClient(base + "/gql"), conf}
}

func (c *Config) BaseUrl() string {
	host := "https://app.plural.sh"
	if c.Endpoint != "" {
		host = fmt.Sprintf("https://%s", c.Endpoint)
	}
	return host
}

func (client *Client) Build(doc string) *graphql.Request {
	req := graphql.NewRequest(doc)
	req.Header.Set("Authorization", "Bearer "+client.config.Token)
	return req
}

func (client *Client) Run(req *graphql.Request, resp interface{}) error {
	return client.gqlClient.Run(context.Background(), req, &resp)
}

func (client *Client) DnsRecords() (records []*DnsRecord, err error) {
	var resp struct {
		DnsRecords struct {
			Edges []struct {
				Node *DnsRecord
			}
		}
	}
	req := client.Build(dnsRecordsQuery)
	req.Var("cluster", client.config.Cluster)
	req.Var("provider", strings.ToUpper(client.config.Provider))
	err = client.Run(req, &resp)
	if err != nil {
		return
	}

	records = make([]*DnsRecord, len(resp.DnsRecords.Edges))
	for i, edge := range resp.DnsRecords.Edges {
		records[i] = edge.Node
	}
	return
}

func (client *Client) CreateRecord(record *DnsRecord) (result *DnsRecord, err error) {
	var resp struct {
		CreateDnsRecord *DnsRecord
	}
	req := client.Build(createDnsRecord)
	req.Var("cluster", client.config.Cluster)
	req.Var("provider", strings.ToUpper(client.config.Provider))
	req.Var("attributes", record)
	err = client.Run(req, &resp)
	result = resp.CreateDnsRecord
	return
}

func (client *Client) DeleteRecord(name, ttype string) error {
	var resp struct {
		DeleteDnsRecord *DnsRecord
	}
	req := client.Build(deleteDnsRecord)
	req.Var("type", ttype)
	req.Var("name", name)
	return client.Run(req, &resp)
}
