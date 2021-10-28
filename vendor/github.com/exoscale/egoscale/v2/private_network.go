package v2

import (
	"context"
	"net"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// PrivateNetworkLease represents a managed Private Network lease.
type PrivateNetworkLease struct {
	InstanceID *string
	IPAddress  *net.IP
}

// PrivateNetwork represents a Private Network.
type PrivateNetwork struct {
	Description *string
	EndIP       *net.IP
	ID          *string `req-for:"update,delete"`
	Name        *string `req-for:"create"`
	Netmask     *net.IP
	StartIP     *net.IP
	Leases      []*PrivateNetworkLease
}

func privateNetworkFromAPI(p *oapi.PrivateNetwork) *PrivateNetwork {
	return &PrivateNetwork{
		Description: p.Description,
		EndIP: func() (v *net.IP) {
			if p.EndIp != nil {
				ip := net.ParseIP(*p.EndIp)
				v = &ip
			}
			return
		}(),
		ID:   p.Id,
		Name: p.Name,
		Netmask: func() (v *net.IP) {
			if p.Netmask != nil {
				ip := net.ParseIP(*p.Netmask)
				v = &ip
			}
			return
		}(),
		StartIP: func() (v *net.IP) {
			if p.StartIp != nil {
				ip := net.ParseIP(*p.StartIp)
				v = &ip
			}
			return
		}(),
		Leases: func() (v []*PrivateNetworkLease) {
			if p.Leases != nil {
				v = make([]*PrivateNetworkLease, len(*p.Leases))
				for i, lease := range *p.Leases {
					v[i] = &PrivateNetworkLease{
						InstanceID: lease.InstanceId,
						IPAddress:  func() *net.IP { ip := net.ParseIP(*lease.Ip); return &ip }(),
					}
				}
			}
			return
		}(),
	}
}

// CreatePrivateNetwork creates a Private Network.
func (c *Client) CreatePrivateNetwork(
	ctx context.Context,
	zone string,
	privateNetwork *PrivateNetwork,
) (*PrivateNetwork, error) {
	if err := validateOperationParams(privateNetwork, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreatePrivateNetworkJSONRequestBody{
			Description: privateNetwork.Description,
			EndIp: func() (ip *string) {
				if privateNetwork.EndIP != nil {
					v := privateNetwork.EndIP.String()
					return &v
				}
				return
			}(),
			Name: *privateNetwork.Name,
			Netmask: func() (ip *string) {
				if privateNetwork.Netmask != nil {
					v := privateNetwork.Netmask.String()
					return &v
				}
				return
			}(),
			StartIp: func() (ip *string) {
				if privateNetwork.StartIP != nil {
					v := privateNetwork.StartIP.String()
					return &v
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
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetPrivateNetwork(ctx, zone, *res.(*oapi.Reference).Id)
}

// DeletePrivateNetwork deletes a Private Network.
func (c *Client) DeletePrivateNetwork(ctx context.Context, zone string, privateNetwork *PrivateNetwork) error {
	if err := validateOperationParams(privateNetwork, "delete"); err != nil {
		return err
	}

	resp, err := c.DeletePrivateNetworkWithResponse(apiv2.WithZone(ctx, zone), *privateNetwork.ID)
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

// FindPrivateNetwork attempts to find a Private Network by name or ID.
// In case the identifier is a name and multiple resources match, an ErrTooManyFound error is returned.
func (c *Client) FindPrivateNetwork(ctx context.Context, zone, x string) (*PrivateNetwork, error) {
	res, err := c.ListPrivateNetworks(ctx, zone)
	if err != nil {
		return nil, err
	}

	var found *PrivateNetwork
	for _, r := range res {
		if *r.ID == x {
			return c.GetPrivateNetwork(ctx, zone, *r.ID)
		}

		// Historically, the Exoscale API allowed users to create multiple Private Networks sharing a common name.
		// This function being expected to return one resource at most, in case the specified identifier is a name
		// we have to check that there aren't more that one matching result before returning it.
		if *r.Name == x {
			if found != nil {
				return nil, apiv2.ErrTooManyFound
			}
			found = r
		}
	}

	if found != nil {
		return c.GetPrivateNetwork(ctx, zone, *found.ID)
	}

	return nil, apiv2.ErrNotFound
}

// GetPrivateNetwork returns the Private Network corresponding to the specified ID.
func (c *Client) GetPrivateNetwork(ctx context.Context, zone, id string) (*PrivateNetwork, error) {
	resp, err := c.GetPrivateNetworkWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return privateNetworkFromAPI(resp.JSON200), nil
}

// ListPrivateNetworks returns the list of existing Private Networks.
func (c *Client) ListPrivateNetworks(ctx context.Context, zone string) ([]*PrivateNetwork, error) {
	list := make([]*PrivateNetwork, 0)

	resp, err := c.ListPrivateNetworksWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.PrivateNetworks != nil {
		for i := range *resp.JSON200.PrivateNetworks {
			list = append(list, privateNetworkFromAPI(&(*resp.JSON200.PrivateNetworks)[i]))
		}
	}

	return list, nil
}

// UpdatePrivateNetwork updates a Private Network.
func (c *Client) UpdatePrivateNetwork(ctx context.Context, zone string, privateNetwork *PrivateNetwork) error {
	if err := validateOperationParams(privateNetwork, "update"); err != nil {
		return err
	}

	resp, err := c.UpdatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		*privateNetwork.ID,
		oapi.UpdatePrivateNetworkJSONRequestBody{
			Description: privateNetwork.Description,
			EndIp: func() (ip *string) {
				if privateNetwork.EndIP != nil {
					v := privateNetwork.EndIP.String()
					return &v
				}
				return
			}(),
			Name: privateNetwork.Name,
			Netmask: func() (ip *string) {
				if privateNetwork.Netmask != nil {
					v := privateNetwork.Netmask.String()
					return &v
				}
				return
			}(),
			StartIp: func() (ip *string) {
				if privateNetwork.StartIP != nil {
					v := privateNetwork.StartIP.String()
					return &v
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
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// UpdatePrivateNetworkInstanceIPAddress updates the IP address of a Compute instance attached to a managed
// Private Network.
func (c *Client) UpdatePrivateNetworkInstanceIPAddress(
	ctx context.Context,
	zone string,
	instance *Instance,
	privateNetwork *PrivateNetwork,
	ip net.IP,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(privateNetwork, "update"); err != nil {
		return err
	}

	resp, err := c.UpdatePrivateNetworkInstanceIpWithResponse(
		apiv2.WithZone(ctx, zone),
		*privateNetwork.ID,
		oapi.UpdatePrivateNetworkInstanceIpJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
			Ip: func() *string {
				s := ip.String()
				return &s
			}(),
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))

	return err
}
