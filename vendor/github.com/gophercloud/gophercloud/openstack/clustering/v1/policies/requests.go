package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	// Limit limits the number of Policies to return.
	Limit int `q:"limit"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker"`

	// Sorts the response by one or more attribute and optional sort direction combinations.
	Sort string `q:"sort"`

	// GlobalProject indicates whether to include resources for all projects or resources for the current project
	GlobalProject *bool `q:"global_project"`

	// Name to filter the response by the specified name property of the object
	Name string `q:"name"`

	// Filter the response by the specified type property of the object
	Type string `q:"type"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to retrieve a list of policies.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := policyListURL(client)
	if opts != nil {
		query, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := PolicyPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	})
}
