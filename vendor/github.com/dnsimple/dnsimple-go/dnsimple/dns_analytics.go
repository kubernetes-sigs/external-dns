package dnsimple

import (
	"context"
	"fmt"
)

// DnsAnalyticsService handles communication with the DNS Analytics related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/dns-analytics/
type DnsAnalyticsService struct {
	client *Client
}

// DnsAnalytics represents DNS Analytics data.
type DnsAnalytics struct {
	Volume   int64
	ZoneName string
	Date     string
}

type DnsAnalyticsQueryParameters struct {
	AccountId interface{} `json:"account_id"`
	StartDate string      `json:"start_date"`
	EndDate   string      `json:"end_date"`
	Sort      string      `json:"sort"`
	Page      int64       `json:"page"`
	PerPage   int64       `json:"per_page"`
	Groupings string      `json:"groupings"`
}

// RowAndHeaderData represents the special payload of `data` when it includes lists of `rows` and `headers`.
type RowAndHeaderData struct {
	Rows    [][]interface{} `json:"rows"`
	Headers []string        `json:"headers"`
}

// DnsAnalyticsResponse represents a response from an API method that returns DnsAnalytics data.
type DnsAnalyticsResponse struct {
	Response
	Data             []DnsAnalytics
	RowAndHeaderData RowAndHeaderData            `json:"data"`
	Query            DnsAnalyticsQueryParameters `json:"query"`
}

func (r *DnsAnalyticsResponse) marshalData() {
	list := make([]DnsAnalytics, len(r.RowAndHeaderData.Rows))

	for i, row := range r.RowAndHeaderData.Rows {
		var dataEntry DnsAnalytics
		for j, header := range r.RowAndHeaderData.Headers {
			switch header {
			case "volume":
				dataEntry.Volume = int64(row[j].(float64))
			case "zone_name":
				dataEntry.ZoneName = row[j].(string)
			case "date":
				dataEntry.Date = row[j].(string)
			}
		}

		list[i] = dataEntry
	}
	r.Data = list
}

// DnsAnalyticsOptions specifies the optional parameters you can provide
// to customize the DnsAnalyticsService.Query method.
type DnsAnalyticsOptions struct {
	// Group results by the provided list of attributes separated by a comma
	Groupings *string `url:"groupings,omitempty"`

	// Only include results starting from the provided date in ISO8601 format
	StartDate *string `url:"start_date,omitempty"`

	// Only include results up to the provided date in ISO8601 format
	EndDate *string `url:"end_date,omitempty"`

	ListOptions
}

// Query gets DNS Analytics data for an account
//
// See https://developer.dnsimple.com/v2/dns-analytics/#query
func (s *DnsAnalyticsService) Query(ctx context.Context, accountID int64, options *DnsAnalyticsOptions) (*DnsAnalyticsResponse, error) {
	path := versioned(fmt.Sprintf("/%v/dns_analytics", accountID))
	dnsAnalyticsResponse := &DnsAnalyticsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, dnsAnalyticsResponse)
	if err != nil {
		return dnsAnalyticsResponse, err
	}

	dnsAnalyticsResponse.HTTPResponse = resp
	dnsAnalyticsResponse.marshalData()
	return dnsAnalyticsResponse, nil
}
