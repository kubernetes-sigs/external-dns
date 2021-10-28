package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// DeployTarget represents a Deploy Target.
type DeployTarget struct {
	Description *string
	ID          *string
	Name        *string
	Type        *string
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (d DeployTarget) ToAPIMock() interface{} {
	return oapi.DeployTarget{
		Description: d.Description,
		Id:          d.ID,
		Name:        d.Name,
		Type:        (*oapi.DeployTargetType)(d.Type),
	}
}

func deployTargetFromAPI(d *oapi.DeployTarget) *DeployTarget {
	return &DeployTarget{
		Description: d.Description,
		ID:          d.Id,
		Name:        d.Name,
		Type:        (*string)(d.Type),
	}
}

// FindDeployTarget attempts to find a Deploy Target by name or ID.
func (c *Client) FindDeployTarget(ctx context.Context, zone, x string) (*DeployTarget, error) {
	res, err := c.ListDeployTargets(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetDeployTarget(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetDeployTarget returns the Deploy Target corresponding to the specified ID.
func (c *Client) GetDeployTarget(ctx context.Context, zone, id string) (*DeployTarget, error) {
	resp, err := c.GetDeployTargetWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return deployTargetFromAPI(resp.JSON200), nil
}

// ListDeployTargets returns the list of existing Deploy Targets.
func (c *Client) ListDeployTargets(ctx context.Context, zone string) ([]*DeployTarget, error) {
	list := make([]*DeployTarget, 0)

	resp, err := c.ListDeployTargetsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DeployTargets != nil {
		for i := range *resp.JSON200.DeployTargets {
			list = append(list, deployTargetFromAPI(&(*resp.JSON200.DeployTargets)[i]))
		}
	}

	return list, nil
}
