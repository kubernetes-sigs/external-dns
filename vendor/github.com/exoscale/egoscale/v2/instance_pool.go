package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// InstancePoolManager represents an Instance Pool manager.
type InstancePoolManager struct {
	ID   string
	Type string
}

// InstancePool represents an Instance Pool.
type InstancePool struct {
	AntiAffinityGroupIDs *[]string
	DeployTargetID       *string
	Description          *string
	DiskSize             *int64 `req-for:"create"`
	ElasticIPIDs         *[]string
	ID                   *string `req-for:"update,delete"`
	IPv6Enabled          *bool
	InstanceIDs          *[]string
	InstancePrefix       *string
	InstanceTypeID       *string `req-for:"create"`
	Labels               *map[string]string
	Manager              *InstancePoolManager
	Name                 *string `req-for:"create"`
	PrivateNetworkIDs    *[]string
	SSHKey               *string
	SecurityGroupIDs     *[]string
	Size                 *int64 `req-for:"create"`
	State                *string
	TemplateID           *string `req-for:"create"`
	UserData             *string
}

func instancePoolFromAPI(i *oapi.InstancePool) *InstancePool {
	return &InstancePool{
		AntiAffinityGroupIDs: func() (v *[]string) {
			if i.AntiAffinityGroups != nil && len(*i.AntiAffinityGroups) > 0 {
				ids := make([]string, len(*i.AntiAffinityGroups))
				for i, item := range *i.AntiAffinityGroups {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		DeployTargetID: func() (v *string) {
			if i.DeployTarget != nil {
				v = i.DeployTarget.Id
			}
			return
		}(),
		Description: i.Description,
		DiskSize:    i.DiskSize,
		ElasticIPIDs: func() (v *[]string) {
			if i.ElasticIps != nil && len(*i.ElasticIps) > 0 {
				ids := make([]string, len(*i.ElasticIps))
				for i, item := range *i.ElasticIps {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		ID:          i.Id,
		IPv6Enabled: i.Ipv6Enabled,
		InstanceIDs: func() (v *[]string) {
			if i.Instances != nil && len(*i.Instances) > 0 {
				ids := make([]string, len(*i.Instances))
				for i, item := range *i.Instances {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		InstancePrefix: i.InstancePrefix,
		InstanceTypeID: i.InstanceType.Id,
		Labels: func() (v *map[string]string) {
			if i.Labels != nil && len(i.Labels.AdditionalProperties) > 0 {
				v = &i.Labels.AdditionalProperties
			}
			return
		}(),
		Manager: func() *InstancePoolManager {
			if i.Manager != nil {
				return &InstancePoolManager{
					ID:   *i.Manager.Id,
					Type: string(*i.Manager.Type),
				}
			}
			return nil
		}(),
		Name: i.Name,
		PrivateNetworkIDs: func() (v *[]string) {
			if i.PrivateNetworks != nil && len(*i.PrivateNetworks) > 0 {
				ids := make([]string, len(*i.PrivateNetworks))
				for i, item := range *i.PrivateNetworks {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		SSHKey: func() (v *string) {
			if i.SshKey != nil {
				v = i.SshKey.Name
			}
			return
		}(),
		SecurityGroupIDs: func() (v *[]string) {
			if i.SecurityGroups != nil && len(*i.SecurityGroups) > 0 {
				ids := make([]string, len(*i.SecurityGroups))
				for i, item := range *i.SecurityGroups {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		Size:       i.Size,
		State:      (*string)(i.State),
		TemplateID: i.Template.Id,
		UserData:   i.UserData,
	}
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (i InstancePool) ToAPIMock() interface{} {
	return oapi.InstancePool{
		AntiAffinityGroups: func() *[]oapi.AntiAffinityGroup {
			if i.AntiAffinityGroupIDs != nil {
				list := make([]oapi.AntiAffinityGroup, len(*i.AntiAffinityGroupIDs))
				for j, id := range *i.AntiAffinityGroupIDs {
					id := id
					list[j] = oapi.AntiAffinityGroup{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		DeployTarget: &oapi.DeployTarget{Id: i.DeployTargetID},
		Description:  i.Description,
		DiskSize:     i.DiskSize,
		ElasticIps: func() *[]oapi.ElasticIp {
			if i.ElasticIPIDs != nil {
				list := make([]oapi.ElasticIp, len(*i.ElasticIPIDs))
				for j, id := range *i.ElasticIPIDs {
					id := id
					list[j] = oapi.ElasticIp{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		Id:             i.ID,
		InstancePrefix: i.InstancePrefix,
		InstanceType:   &oapi.InstanceType{Id: i.InstanceTypeID},
		Instances: func() *[]oapi.Instance {
			if i.InstanceIDs != nil {
				list := make([]oapi.Instance, len(*i.InstanceIDs))
				for j, id := range *i.InstanceIDs {
					id := id
					list[j] = oapi.Instance{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		Ipv6Enabled: i.IPv6Enabled,
		Labels: func() *oapi.Labels {
			if i.Labels != nil {
				return &oapi.Labels{AdditionalProperties: *i.Labels}
			}
			return nil
		}(),
		Manager: func() *oapi.Manager {
			if i.Manager != nil {
				return &oapi.Manager{
					Id:   &i.Manager.ID,
					Type: (*oapi.ManagerType)(&i.Manager.Type),
				}
			}
			return nil
		}(),
		Name: i.Name,
		PrivateNetworks: func() *[]oapi.PrivateNetwork {
			if i.PrivateNetworkIDs != nil {
				list := make([]oapi.PrivateNetwork, len(*i.PrivateNetworkIDs))
				for j, id := range *i.PrivateNetworkIDs {
					id := id
					list[j] = oapi.PrivateNetwork{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		SecurityGroups: func() *[]oapi.SecurityGroup {
			if i.SecurityGroupIDs != nil {
				list := make([]oapi.SecurityGroup, len(*i.SecurityGroupIDs))
				for j, id := range *i.SecurityGroupIDs {
					id := id
					list[j] = oapi.SecurityGroup{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		Size: i.Size,
		SshKey: func() *oapi.SshKey {
			if i.SSHKey != nil {
				return &oapi.SshKey{Name: i.SSHKey}
			}
			return nil
		}(),
		State:    (*oapi.InstancePoolState)(i.State),
		Template: &oapi.Template{Id: i.TemplateID},
		UserData: i.UserData,
	}
}

// CreateInstancePool creates an Instance Pool.
func (c *Client) CreateInstancePool(
	ctx context.Context,
	zone string,
	instancePool *InstancePool,
) (*InstancePool, error) {
	if err := validateOperationParams(instancePool, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateInstancePoolWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreateInstancePoolJSONRequestBody{
			AntiAffinityGroups: func() (v *[]oapi.AntiAffinityGroup) {
				if instancePool.AntiAffinityGroupIDs != nil {
					ids := make([]oapi.AntiAffinityGroup, len(*instancePool.AntiAffinityGroupIDs))
					for i, item := range *instancePool.AntiAffinityGroupIDs {
						item := item
						ids[i] = oapi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *oapi.DeployTarget) {
				if instancePool.DeployTargetID != nil {
					v = &oapi.DeployTarget{Id: instancePool.DeployTargetID}
				}
				return
			}(),
			Description: instancePool.Description,
			DiskSize:    *instancePool.DiskSize,
			ElasticIps: func() (v *[]oapi.ElasticIp) {
				if instancePool.ElasticIPIDs != nil {
					ids := make([]oapi.ElasticIp, len(*instancePool.ElasticIPIDs))
					for i, item := range *instancePool.ElasticIPIDs {
						item := item
						ids[i] = oapi.ElasticIp{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			InstancePrefix: instancePool.InstancePrefix,
			InstanceType:   oapi.InstanceType{Id: instancePool.InstanceTypeID},
			Ipv6Enabled:    instancePool.IPv6Enabled,
			Labels: func() (v *oapi.Labels) {
				if instancePool.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *instancePool.Labels}
				}
				return
			}(),
			Name: *instancePool.Name,
			PrivateNetworks: func() (v *[]oapi.PrivateNetwork) {
				if instancePool.PrivateNetworkIDs != nil {
					ids := make([]oapi.PrivateNetwork, len(*instancePool.PrivateNetworkIDs))
					for i, item := range *instancePool.PrivateNetworkIDs {
						item := item
						ids[i] = oapi.PrivateNetwork{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SecurityGroups: func() (v *[]oapi.SecurityGroup) {
				if instancePool.SecurityGroupIDs != nil {
					ids := make([]oapi.SecurityGroup, len(*instancePool.SecurityGroupIDs))
					for i, item := range *instancePool.SecurityGroupIDs {
						item := item
						ids[i] = oapi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			Size: *instancePool.Size,
			SshKey: func() (v *oapi.SshKey) {
				if instancePool.SSHKey != nil {
					v = &oapi.SshKey{Name: instancePool.SSHKey}
				}
				return
			}(),
			Template: oapi.Template{Id: instancePool.TemplateID},
			UserData: instancePool.UserData,
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

	return c.GetInstancePool(ctx, zone, *res.(*oapi.Reference).Id)
}

// DeleteInstancePool deletes an Instance Pool.
func (c *Client) DeleteInstancePool(ctx context.Context, zone string, instancePool *InstancePool) error {
	if err := validateOperationParams(instancePool, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteInstancePoolWithResponse(apiv2.WithZone(ctx, zone), *instancePool.ID)
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

// EvictInstancePoolMembers evicts the specified members (identified by their Compute instance ID) from the
// Instance Pool corresponding to the specified ID.
func (c *Client) EvictInstancePoolMembers(
	ctx context.Context,
	zone string,
	instancePool *InstancePool,
	members []string,
) error {
	if err := validateOperationParams(instancePool, "update"); err != nil {
		return err
	}

	resp, err := c.EvictInstancePoolMembersWithResponse(
		apiv2.WithZone(ctx, zone),
		*instancePool.ID,
		oapi.EvictInstancePoolMembersJSONRequestBody{Instances: &members},
	)
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

// FindInstancePool attempts to find an Instance Pool by name or ID.
func (c *Client) FindInstancePool(ctx context.Context, zone, x string) (*InstancePool, error) {
	res, err := c.ListInstancePools(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetInstancePool(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetInstancePool returns the Instance Pool corresponding to the specified ID.
func (c *Client) GetInstancePool(ctx context.Context, zone, id string) (*InstancePool, error) {
	resp, err := c.GetInstancePoolWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return instancePoolFromAPI(resp.JSON200), nil
}

// ListInstancePools returns the list of existing Instance Pools.
func (c *Client) ListInstancePools(ctx context.Context, zone string) ([]*InstancePool, error) {
	list := make([]*InstancePool, 0)

	resp, err := c.ListInstancePoolsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.InstancePools != nil {
		for i := range *resp.JSON200.InstancePools {
			list = append(list, instancePoolFromAPI(&(*resp.JSON200.InstancePools)[i]))
		}
	}

	return list, nil
}

// ScaleInstancePool scales an Instance Pool.
func (c *Client) ScaleInstancePool(ctx context.Context, zone string, instancePool *InstancePool, size int64) error {
	if err := validateOperationParams(instancePool, "update"); err != nil {
		return err
	}

	resp, err := c.ScaleInstancePoolWithResponse(
		apiv2.WithZone(ctx, zone),
		*instancePool.ID,
		oapi.ScaleInstancePoolJSONRequestBody{Size: size},
	)
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

// UpdateInstancePool updates an Instance Pool.
func (c *Client) UpdateInstancePool(ctx context.Context, zone string, instancePool *InstancePool) error {
	if err := validateOperationParams(instancePool, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateInstancePoolWithResponse(
		apiv2.WithZone(ctx, zone),
		*instancePool.ID,
		oapi.UpdateInstancePoolJSONRequestBody{
			AntiAffinityGroups: func() (v *[]oapi.AntiAffinityGroup) {
				if instancePool.AntiAffinityGroupIDs != nil {
					ids := make([]oapi.AntiAffinityGroup, len(*instancePool.AntiAffinityGroupIDs))
					for i, item := range *instancePool.AntiAffinityGroupIDs {
						item := item
						ids[i] = oapi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *oapi.DeployTarget) {
				if instancePool.DeployTargetID != nil {
					v = &oapi.DeployTarget{Id: instancePool.DeployTargetID}
				}
				return
			}(),
			Description: instancePool.Description,
			DiskSize:    instancePool.DiskSize,
			ElasticIps: func() (v *[]oapi.ElasticIp) {
				if instancePool.ElasticIPIDs != nil {
					ids := make([]oapi.ElasticIp, len(*instancePool.ElasticIPIDs))
					for i, item := range *instancePool.ElasticIPIDs {
						item := item
						ids[i] = oapi.ElasticIp{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			InstancePrefix: instancePool.InstancePrefix,
			InstanceType: func() (v *oapi.InstanceType) {
				if instancePool.InstanceTypeID != nil {
					v = &oapi.InstanceType{Id: instancePool.InstanceTypeID}
				}
				return
			}(),
			Ipv6Enabled: instancePool.IPv6Enabled,
			Labels: func() (v *oapi.Labels) {
				if instancePool.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *instancePool.Labels}
				}
				return
			}(),
			Name: instancePool.Name,
			PrivateNetworks: func() (v *[]oapi.PrivateNetwork) {
				if instancePool.PrivateNetworkIDs != nil {
					ids := make([]oapi.PrivateNetwork, len(*instancePool.PrivateNetworkIDs))
					for i, item := range *instancePool.PrivateNetworkIDs {
						item := item
						ids[i] = oapi.PrivateNetwork{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SecurityGroups: func() (v *[]oapi.SecurityGroup) {
				if instancePool.SecurityGroupIDs != nil {
					ids := make([]oapi.SecurityGroup, len(*instancePool.SecurityGroupIDs))
					for i, item := range *instancePool.SecurityGroupIDs {
						item := item
						ids[i] = oapi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SshKey: func() (v *oapi.SshKey) {
				if instancePool.SSHKey != nil {
					v = &oapi.SshKey{Name: instancePool.SSHKey}
				}
				return
			}(),
			Template: func() (v *oapi.Template) {
				if instancePool.TemplateID != nil {
					v = &oapi.Template{Id: instancePool.TemplateID}
				}
				return
			}(),
			UserData: instancePool.UserData,
		})
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
