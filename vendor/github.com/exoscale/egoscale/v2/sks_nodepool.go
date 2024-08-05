package v2

import (
	"context"
	"fmt"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SKSNodepoolTaint represents an SKS Nodepool Kubernetes Node taint.
type SKSNodepoolTaint struct {
	Effect string
	Value  string
}

func sksNodepoolTaintFromAPI(t *oapi.SksNodepoolTaint) *SKSNodepoolTaint {
	return &SKSNodepoolTaint{
		Effect: string(t.Effect),
		Value:  t.Value,
	}
}

// SKSNodepool represents an SKS Nodepool.
type SKSNodepool struct {
	AddOns               *[]string
	AntiAffinityGroupIDs *[]string
	CreatedAt            *time.Time
	DeployTargetID       *string
	Description          *string
	DiskSize             *int64  `req-for:"create"`
	ID                   *string `req-for:"update,delete"`
	InstancePoolID       *string
	InstancePrefix       *string
	InstanceTypeID       *string `req-for:"create"`
	Labels               *map[string]string
	Name                 *string `req-for:"create"`
	PrivateNetworkIDs    *[]string
	SecurityGroupIDs     *[]string
	Size                 *int64 `req-for:"create"`
	State                *string
	Taints               *map[string]*SKSNodepoolTaint
	TemplateID           *string
	Version              *string
}

func sksNodepoolFromAPI(n *oapi.SksNodepool) *SKSNodepool {
	return &SKSNodepool{
		AddOns: func() (v *[]string) {
			if n.Addons != nil {
				addOns := make([]string, 0)
				for _, a := range *n.Addons {
					addOns = append(addOns, string(a))
				}
				v = &addOns
			}
			return
		}(),
		AntiAffinityGroupIDs: func() (v *[]string) {
			if n.AntiAffinityGroups != nil && len(*n.AntiAffinityGroups) > 0 {
				ids := make([]string, 0)
				for _, item := range *n.AntiAffinityGroups {
					item := item
					ids = append(ids, *item.Id)
				}
				v = &ids
			}
			return
		}(),
		CreatedAt: n.CreatedAt,
		DeployTargetID: func() (v *string) {
			if n.DeployTarget != nil {
				v = &n.DeployTarget.Id
			}
			return
		}(),
		Description:    n.Description,
		DiskSize:       n.DiskSize,
		ID:             n.Id,
		InstancePoolID: n.InstancePool.Id,
		InstancePrefix: n.InstancePrefix,
		InstanceTypeID: n.InstanceType.Id,
		Labels: func() (v *map[string]string) {
			if n.Labels != nil && len(n.Labels.AdditionalProperties) > 0 {
				v = &n.Labels.AdditionalProperties
			}
			return
		}(),
		Name: n.Name,
		PrivateNetworkIDs: func() (v *[]string) {
			if n.PrivateNetworks != nil && len(*n.PrivateNetworks) > 0 {
				ids := make([]string, 0)
				for _, item := range *n.PrivateNetworks {
					item := item
					ids = append(ids, *item.Id)
				}
				v = &ids
			}
			return
		}(),
		SecurityGroupIDs: func() (v *[]string) {
			if n.SecurityGroups != nil && len(*n.SecurityGroups) > 0 {
				ids := make([]string, 0)
				for _, item := range *n.SecurityGroups {
					item := item
					ids = append(ids, *item.Id)
				}
				v = &ids
			}
			return
		}(),
		Size:  n.Size,
		State: (*string)(n.State),
		Taints: func() (v *map[string]*SKSNodepoolTaint) {
			if n.Taints != nil && len(n.Taints.AdditionalProperties) > 0 {
				taints := make(map[string]*SKSNodepoolTaint)
				for k, t := range n.Taints.AdditionalProperties {
					taints[k] = sksNodepoolTaintFromAPI(&t)
				}
				v = &taints
			}
			return
		}(),
		TemplateID: n.Template.Id,
		Version:    n.Version,
	}
}

// CreateSKSNodepool create an SKS Nodepool.
func (c *Client) CreateSKSNodepool(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	nodepool *SKSNodepool,
) (*SKSNodepool, error) {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return nil, err
	}
	if err := validateOperationParams(nodepool, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateSksNodepoolWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		oapi.CreateSksNodepoolJSONRequestBody{
			Addons: func() (v *[]oapi.CreateSksNodepoolJSONBodyAddons) {
				if nodepool.AddOns != nil {
					addOns := make([]oapi.CreateSksNodepoolJSONBodyAddons, len(*nodepool.AddOns))
					for i, a := range *nodepool.AddOns {
						addOns[i] = oapi.CreateSksNodepoolJSONBodyAddons(a)
					}
					v = &addOns
				}
				return
			}(),
			AntiAffinityGroups: func() (v *[]oapi.AntiAffinityGroup) {
				if nodepool.AntiAffinityGroupIDs != nil {
					ids := make([]oapi.AntiAffinityGroup, len(*nodepool.AntiAffinityGroupIDs))
					for i, item := range *nodepool.AntiAffinityGroupIDs {
						item := item
						ids[i] = oapi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *oapi.DeployTarget) {
				if nodepool.DeployTargetID != nil {
					v = &oapi.DeployTarget{Id: *nodepool.DeployTargetID}
				}
				return
			}(),
			Description:    nodepool.Description,
			DiskSize:       *nodepool.DiskSize,
			InstancePrefix: nodepool.InstancePrefix,
			InstanceType:   oapi.InstanceType{Id: nodepool.InstanceTypeID},
			Labels: func() (v *oapi.Labels) {
				if nodepool.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *nodepool.Labels}
				}
				return
			}(),
			Name: *nodepool.Name,
			PrivateNetworks: func() (v *[]oapi.PrivateNetwork) {
				if nodepool.PrivateNetworkIDs != nil {
					ids := make([]oapi.PrivateNetwork, len(*nodepool.PrivateNetworkIDs))
					for i, item := range *nodepool.PrivateNetworkIDs {
						item := item
						ids[i] = oapi.PrivateNetwork{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SecurityGroups: func() (v *[]oapi.SecurityGroup) {
				if nodepool.SecurityGroupIDs != nil {
					ids := make([]oapi.SecurityGroup, len(*nodepool.SecurityGroupIDs))
					for i, item := range *nodepool.SecurityGroupIDs {
						item := item
						ids[i] = oapi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			Size: *nodepool.Size,
			Taints: func() (v *oapi.SksNodepoolTaints) {
				if nodepool.Taints != nil {
					taints := oapi.SksNodepoolTaints{AdditionalProperties: map[string]oapi.SksNodepoolTaint{}}
					for k, t := range *nodepool.Taints {
						taints.AdditionalProperties[k] = oapi.SksNodepoolTaint{
							Effect: (oapi.SksNodepoolTaintEffect)(t.Effect),
							Value:  t.Value,
						}
					}
					v = &taints
				}
				return
			}(),
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	nodepoolRes, err := c.GetSksNodepoolWithResponse(ctx, *cluster.ID, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Nodepool: %s", err)
	}

	return sksNodepoolFromAPI(nodepoolRes.JSON200), nil
}

// DeleteSKSNodepool deletes an SKS Nodepool.
func (c *Client) DeleteSKSNodepool(ctx context.Context, zone string, cluster *SKSCluster, nodepool *SKSNodepool) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(nodepool, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSksNodepoolWithResponse(apiv2.WithZone(ctx, zone), *cluster.ID, *nodepool.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// EvictSKSNodepoolMembers evicts the specified members (identified by their Compute instance ID) from the
// SKS cluster Nodepool.
func (c *Client) EvictSKSNodepoolMembers(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	nodepool *SKSNodepool,
	members []string,
) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(nodepool, "update"); err != nil {
		return err
	}

	resp, err := c.EvictSksNodepoolMembersWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		*nodepool.ID,
		oapi.EvictSksNodepoolMembersJSONRequestBody{Instances: &members},
	)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ScaleSKSNodepool scales the SKS cluster Nodepool to the specified number of Kubernetes Nodes.
func (c *Client) ScaleSKSNodepool(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	nodepool *SKSNodepool,
	size int64,
) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(nodepool, "update"); err != nil {
		return err
	}

	resp, err := c.ScaleSksNodepoolWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		*nodepool.ID,
		oapi.ScaleSksNodepoolJSONRequestBody{Size: size},
	)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// UpdateSKSNodepool updates an SKS Nodepool.
func (c *Client) UpdateSKSNodepool(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	nodepool *SKSNodepool,
) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(nodepool, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateSksNodepoolWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		*nodepool.ID,
		oapi.UpdateSksNodepoolJSONRequestBody{
			AntiAffinityGroups: func() (v *[]oapi.AntiAffinityGroup) {
				if nodepool.AntiAffinityGroupIDs != nil {
					ids := make([]oapi.AntiAffinityGroup, len(*nodepool.AntiAffinityGroupIDs))
					for i, item := range *nodepool.AntiAffinityGroupIDs {
						item := item
						ids[i] = oapi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *oapi.DeployTarget) {
				if nodepool.DeployTargetID != nil {
					v = &oapi.DeployTarget{Id: *nodepool.DeployTargetID}
				}
				return
			}(),
			Description:    oapi.NilableString(nodepool.Description),
			DiskSize:       nodepool.DiskSize,
			InstancePrefix: nodepool.InstancePrefix,
			InstanceType: func() (v *oapi.InstanceType) {
				if nodepool.InstanceTypeID != nil {
					v = &oapi.InstanceType{Id: nodepool.InstanceTypeID}
				}
				return
			}(),
			Labels: func() (v *oapi.Labels) {
				if nodepool.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *nodepool.Labels}
				}
				return
			}(),
			Name: nodepool.Name,
			PrivateNetworks: func() (v *[]oapi.PrivateNetwork) {
				if nodepool.PrivateNetworkIDs != nil {
					ids := make([]oapi.PrivateNetwork, len(*nodepool.PrivateNetworkIDs))
					for i, item := range *nodepool.PrivateNetworkIDs {
						item := item
						ids[i] = oapi.PrivateNetwork{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SecurityGroups: func() (v *[]oapi.SecurityGroup) {
				if nodepool.SecurityGroupIDs != nil {
					ids := make([]oapi.SecurityGroup, len(*nodepool.SecurityGroupIDs))
					for i, item := range *nodepool.SecurityGroupIDs {
						item := item
						ids[i] = oapi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			Taints: func() (v *oapi.SksNodepoolTaints) {
				if nodepool.Taints != nil {
					taints := oapi.SksNodepoolTaints{AdditionalProperties: map[string]oapi.SksNodepoolTaint{}}
					for k, t := range *nodepool.Taints {
						taints.AdditionalProperties[k] = oapi.SksNodepoolTaint{
							Effect: (oapi.SksNodepoolTaintEffect)(t.Effect),
							Value:  t.Value,
						}
					}
					v = &taints
				}
				return
			}(),
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
