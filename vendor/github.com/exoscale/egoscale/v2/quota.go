package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// Quota represents an Exoscale organization quota.
type Quota struct {
	Resource *string
	Usage    *int64
	Limit    *int64
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (q Quota) ToAPIMock() interface{} {
	return oapi.Quota{
		Limit:    q.Limit,
		Resource: q.Resource,
		Usage:    q.Usage,
	}
}

func quotaFromAPI(q *oapi.Quota) *Quota {
	return &Quota{
		Resource: q.Resource,
		Usage:    q.Usage,
		Limit:    q.Limit,
	}
}

// ListQuotas returns the list of Exoscale organization quotas.
func (c *Client) ListQuotas(ctx context.Context, zone string) ([]*Quota, error) {
	list := make([]*Quota, 0)

	resp, err := c.ListQuotasWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Quotas != nil {
		for i := range *resp.JSON200.Quotas {
			list = append(list, quotaFromAPI(&(*resp.JSON200.Quotas)[i]))
		}
	}

	return list, nil
}

// GetQuota returns the current Exoscale organization quota for the specified resource.
func (c *Client) GetQuota(ctx context.Context, zone, resource string) (*Quota, error) {
	resp, err := c.GetQuotaWithResponse(apiv2.WithZone(ctx, zone), resource)
	if err != nil {
		return nil, err
	}

	return quotaFromAPI(resp.JSON200), nil
}
