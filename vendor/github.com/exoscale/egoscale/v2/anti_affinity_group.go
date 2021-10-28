package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// AntiAffinityGroup represents an Anti-Affinity Group.
type AntiAffinityGroup struct {
	Description *string
	ID          *string `req-for:"delete"`
	InstanceIDs *[]string
	Name        *string `req-for:"create"`
}

func antiAffinityGroupFromAPI(a *oapi.AntiAffinityGroup) *AntiAffinityGroup {
	return &AntiAffinityGroup{
		Description: a.Description,
		ID:          a.Id,
		InstanceIDs: func() (v *[]string) {
			if a.Instances != nil && len(*a.Instances) > 0 {
				ids := make([]string, len(*a.Instances))
				for i, item := range *a.Instances {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		Name: a.Name,
	}
}

// CreateAntiAffinityGroup creates an Anti-Affinity Group.
func (c *Client) CreateAntiAffinityGroup(
	ctx context.Context,
	zone string,
	antiAffinityGroup *AntiAffinityGroup,
) (*AntiAffinityGroup, error) {
	if err := validateOperationParams(antiAffinityGroup, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateAntiAffinityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreateAntiAffinityGroupJSONRequestBody{
			Description: antiAffinityGroup.Description,
			Name:        *antiAffinityGroup.Name,
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetAntiAffinityGroup(ctx, zone, *res.(*oapi.Reference).Id)
}

// DeleteAntiAffinityGroup deletes an Anti-Affinity Group.
func (c *Client) DeleteAntiAffinityGroup(
	ctx context.Context,
	zone string,
	antiAffinityGroup *AntiAffinityGroup,
) error {
	if err := validateOperationParams(antiAffinityGroup, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteAntiAffinityGroupWithResponse(apiv2.WithZone(ctx, zone), *antiAffinityGroup.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// FindAntiAffinityGroup attempts to find an Anti-Affinity Group by name or ID.
func (c *Client) FindAntiAffinityGroup(ctx context.Context, zone, x string) (*AntiAffinityGroup, error) {
	res, err := c.ListAntiAffinityGroups(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetAntiAffinityGroup(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetAntiAffinityGroup returns the Anti-Affinity Group corresponding to the specified ID.
func (c *Client) GetAntiAffinityGroup(ctx context.Context, zone, id string) (*AntiAffinityGroup, error) {
	resp, err := c.GetAntiAffinityGroupWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return antiAffinityGroupFromAPI(resp.JSON200), nil
}

// ListAntiAffinityGroups returns the list of existing Anti-Affinity Groups.
func (c *Client) ListAntiAffinityGroups(ctx context.Context, zone string) ([]*AntiAffinityGroup, error) {
	list := make([]*AntiAffinityGroup, 0)

	resp, err := c.ListAntiAffinityGroupsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.AntiAffinityGroups != nil {
		for i := range *resp.JSON200.AntiAffinityGroups {
			list = append(list, antiAffinityGroupFromAPI(&(*resp.JSON200.AntiAffinityGroups)[i]))
		}
	}

	return list, nil
}
