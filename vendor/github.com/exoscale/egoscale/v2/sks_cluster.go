package v2

import (
	"context"
	"errors"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SKSClusterOIDCConfig represents an SKS cluster OpenID Connect configuration.
type SKSClusterOIDCConfig struct {
	ClientID       *string `req-for:"create"`
	GroupsClaim    *string
	GroupsPrefix   *string
	IssuerURL      *string `req-for:"create"`
	RequiredClaim  *map[string]string
	UsernameClaim  *string
	UsernamePrefix *string
}

// CreateSKSClusterOpt represents a CreateSKSCluster operation option.
type CreateSKSClusterOpt func(body *oapi.CreateSksClusterJSONRequestBody) error

// CreateSKSClusterWithOIDC sets the OpenID Connect configuration to provide to the Kubernetes API Server.
func CreateSKSClusterWithOIDC(v *SKSClusterOIDCConfig) CreateSKSClusterOpt {
	return func(b *oapi.CreateSksClusterJSONRequestBody) error {
		if err := validateOperationParams(v, "create"); err != nil {
			return err
		}

		if v != nil {
			b.Oidc = &oapi.SksOidc{
				ClientId:     *v.ClientID,
				GroupsClaim:  v.GroupsClaim,
				GroupsPrefix: v.GroupsPrefix,
				IssuerUrl:    *v.IssuerURL,
				RequiredClaim: func() *oapi.SksOidc_RequiredClaim {
					if v.RequiredClaim != nil {
						return &oapi.SksOidc_RequiredClaim{AdditionalProperties: *v.RequiredClaim}
					}
					return nil
				}(),
				UsernameClaim:  v.UsernameClaim,
				UsernamePrefix: v.UsernamePrefix,
			}
		}

		return nil
	}
}

// ListSKSClusterVersionsOpt represents a ListSKSClusterVersions operation option.
type ListSKSClusterVersionsOpt func(params *oapi.ListSksClusterVersionsParams)

// ListSKSClusterVersionsWithDeprecated includes deprecated results when listing SKS Cluster versions
// nolint:gocritic
func ListSKSClusterVersionsWithDeprecated(v bool) ListSKSClusterVersionsOpt {
	return func(p *oapi.ListSksClusterVersionsParams) {
		if v {
			vs := "true"
			p.IncludeDeprecated = &vs
		}
	}
}

// SKSCluster represents an SKS cluster.
type SKSCluster struct {
	AddOns       *[]string
	AutoUpgrade  *bool
	CNI          *string
	CreatedAt    *time.Time
	Description  *string
	Endpoint     *string
	ID           *string `req-for:"update,delete"`
	Labels       *map[string]string
	Name         *string `req-for:"create"`
	Nodepools    []*SKSNodepool
	ServiceLevel *string `req-for:"create"`
	State        *string
	Version      *string `req-for:"create"`
	Zone         *string
}

func sksClusterFromAPI(c *oapi.SksCluster, zone string) *SKSCluster {
	return &SKSCluster{
		AddOns: func() (v *[]string) {
			if c.Addons != nil {
				addOns := make([]string, 0)
				for _, a := range *c.Addons {
					addOns = append(addOns, string(a))
				}
				v = &addOns
			}
			return
		}(),
		AutoUpgrade: c.AutoUpgrade,
		CNI:         (*string)(c.Cni),
		CreatedAt:   c.CreatedAt,
		Description: c.Description,
		Endpoint:    c.Endpoint,
		ID:          c.Id,
		Labels: func() (v *map[string]string) {
			if c.Labels != nil && len(c.Labels.AdditionalProperties) > 0 {
				v = &c.Labels.AdditionalProperties
			}
			return
		}(),
		Name: c.Name,
		Nodepools: func() []*SKSNodepool {
			nodepools := make([]*SKSNodepool, 0)
			if c.Nodepools != nil {
				for _, n := range *c.Nodepools {
					n := n
					nodepools = append(nodepools, sksNodepoolFromAPI(&n))
				}
			}
			return nodepools
		}(),
		ServiceLevel: (*string)(c.Level),
		State:        (*string)(c.State),
		Version:      c.Version,
		Zone:         &zone,
	}
}

// CreateSKSCluster creates an SKS cluster.
func (c *Client) CreateSKSCluster(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	opts ...CreateSKSClusterOpt,
) (*SKSCluster, error) {
	if err := validateOperationParams(cluster, "create"); err != nil {
		return nil, err
	}

	body := oapi.CreateSksClusterJSONRequestBody{
		Addons: func() (v *[]oapi.CreateSksClusterJSONBodyAddons) {
			if cluster.AddOns != nil {
				addOns := make([]oapi.CreateSksClusterJSONBodyAddons, len(*cluster.AddOns))
				for i, a := range *cluster.AddOns {
					addOns[i] = oapi.CreateSksClusterJSONBodyAddons(a)
				}
				v = &addOns
			}
			return
		}(),
		AutoUpgrade: cluster.AutoUpgrade,
		Cni:         (*oapi.CreateSksClusterJSONBodyCni)(cluster.CNI),
		Description: cluster.Description,
		Labels: func() (v *oapi.Labels) {
			if cluster.Labels != nil {
				v = &oapi.Labels{AdditionalProperties: *cluster.Labels}
			}
			return
		}(),
		Level:   oapi.CreateSksClusterJSONBodyLevel(*cluster.ServiceLevel),
		Name:    *cluster.Name,
		Version: *cluster.Version,
	}

	for _, opt := range opts {
		if err := opt(&body); err != nil {
			return nil, err
		}
	}

	resp, err := c.CreateSksClusterWithResponse(apiv2.WithZone(ctx, zone), body)
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

	return c.GetSKSCluster(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteSKSCluster deletes an SKS cluster.
func (c *Client) DeleteSKSCluster(ctx context.Context, zone string, cluster *SKSCluster) error {
	if err := validateOperationParams(cluster, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSksClusterWithResponse(apiv2.WithZone(ctx, zone), *cluster.ID)
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

// FindSKSCluster attempts to find an SKS cluster by name or ID.
func (c *Client) FindSKSCluster(ctx context.Context, zone, x string) (*SKSCluster, error) {
	res, err := c.ListSKSClusters(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetSKSCluster(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetSKSCluster returns the SKS cluster corresponding to the specified ID.
func (c *Client) GetSKSCluster(ctx context.Context, zone, id string) (*SKSCluster, error) {
	resp, err := c.GetSksClusterWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return sksClusterFromAPI(resp.JSON200, zone), nil
}

// GetSKSClusterAuthorityCert returns the SKS cluster base64-encoded certificate content for the specified authority.
func (c *Client) GetSKSClusterAuthorityCert(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	authority string,
) (string, error) {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return "", err
	}

	if authority == "" {
		return "", errors.New("authority not specified")
	}

	resp, err := c.GetSksClusterAuthorityCertWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		oapi.GetSksClusterAuthorityCertParamsAuthority(authority),
	)
	if err != nil {
		return "", err
	}

	return oapi.OptionalString(resp.JSON200.Cacert), nil
}

// SKSClusterDeprecatedResource represents an resources deployed in a cluster
// that will be removed in a future release of Kubernetes.
type SKSClusterDeprecatedResource struct {
	Group          *string
	RemovedRelease *string
	Resource       *string
	SubResource    *string
	Version        *string
	RawProperties  map[string]string
}

func sksClusterDeprecatedResourcesFromAPI(c *oapi.SksClusterDeprecatedResource, zone string) *SKSClusterDeprecatedResource {
	return &SKSClusterDeprecatedResource{
		Group:          (*string)(mapValueOrNil(c.AdditionalProperties, "group")),
		RemovedRelease: (*string)(mapValueOrNil(c.AdditionalProperties, "removed_release")),
		Resource:       (*string)(mapValueOrNil(c.AdditionalProperties, "resource")),
		SubResource:    (*string)(mapValueOrNil(c.AdditionalProperties, "subresource")),
		Version:        (*string)(mapValueOrNil(c.AdditionalProperties, "version")),
		RawProperties:  c.AdditionalProperties,
	}
}

func (c *Client) ListSKSClusterDeprecatedResources(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
) ([]*SKSClusterDeprecatedResource, error) {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return nil, err
	}

	resp, err := c.ListSksClusterDeprecatedResourcesWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
	)
	if err != nil {
		return nil, err
	}

	var list []*SKSClusterDeprecatedResource
	if resp.JSON200 != nil && len(*resp.JSON200) > 0 {
		for i := range *resp.JSON200 {
			list = append(list, sksClusterDeprecatedResourcesFromAPI(&(*resp.JSON200)[i], zone))
		}
	}

	return list, nil
}

// GetSKSClusterKubeconfig returns a base64-encoded kubeconfig content for the specified user name, optionally
// associated to specified groups for a duration d (default API-set TTL applies if not specified).
// Fore more information: https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
func (c *Client) GetSKSClusterKubeconfig(
	ctx context.Context,
	zone string,
	cluster *SKSCluster,
	user string,
	groups []string,
	d time.Duration,
) (string, error) {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return "", err
	}

	if user == "" {
		return "", errors.New("user not specified")
	}

	resp, err := c.GenerateSksClusterKubeconfigWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		oapi.GenerateSksClusterKubeconfigJSONRequestBody{
			User:   &user,
			Groups: &groups,
			Ttl: func() *int64 {
				ttl := int64(d.Seconds())
				if ttl > 0 {
					return &ttl
				}
				return nil
			}(),
		})
	if err != nil {
		return "", err
	}

	return oapi.OptionalString(resp.JSON200.Kubeconfig), nil
}

// ListSKSClusters returns the list of existing SKS clusters.
func (c *Client) ListSKSClusters(ctx context.Context, zone string) ([]*SKSCluster, error) {
	list := make([]*SKSCluster, 0)

	resp, err := c.ListSksClustersWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SksClusters != nil {
		for i := range *resp.JSON200.SksClusters {
			list = append(list, sksClusterFromAPI(&(*resp.JSON200.SksClusters)[i], zone))
		}
	}

	return list, nil
}

// ListSKSClusterVersions returns the list of Kubernetes versions supported during SKS cluster creation.
func (c *Client) ListSKSClusterVersions(ctx context.Context, opts ...ListSKSClusterVersionsOpt) ([]string, error) {
	list := make([]string, 0)

	params := oapi.ListSksClusterVersionsParams{}

	for _, opt := range opts {
		opt(&params)
	}

	resp, err := c.ListSksClusterVersionsWithResponse(ctx, &params)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SksClusterVersions != nil {
		for i := range *resp.JSON200.SksClusterVersions {
			version := &(*resp.JSON200.SksClusterVersions)[i]
			list = append(list, *version)
		}
	}

	return list, nil
}

// RotateSKSClusterCCMCredentials rotates the Exoscale IAM credentials managed by the SKS control plane for the
// Kubernetes Exoscale Cloud Controller Manager.
func (c *Client) RotateSKSClusterCCMCredentials(ctx context.Context, zone string, cluster *SKSCluster) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}

	resp, err := c.RotateSksCcmCredentialsWithResponse(apiv2.WithZone(ctx, zone), *cluster.ID)
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

// UpdateSKSCluster updates an SKS cluster.
func (c *Client) UpdateSKSCluster(ctx context.Context, zone string, cluster *SKSCluster) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateSksClusterWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		oapi.UpdateSksClusterJSONRequestBody{
			AutoUpgrade: cluster.AutoUpgrade,
			Description: oapi.NilableString(cluster.Description),
			Labels: func() (v *oapi.Labels) {
				if cluster.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *cluster.Labels}
				}
				return
			}(),
			Name: cluster.Name,
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

// UpgradeSKSCluster upgrades an SKS cluster to the requested Kubernetes version.
func (c *Client) UpgradeSKSCluster(ctx context.Context, zone string, cluster *SKSCluster, version string) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}

	resp, err := c.UpgradeSksClusterWithResponse(
		apiv2.WithZone(ctx, zone),
		*cluster.ID,
		oapi.UpgradeSksClusterJSONRequestBody{Version: version})
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

// UpgradeSKSClusterServiceLevel upgrades an SKS cluster to service level "pro".
func (c *Client) UpgradeSKSClusterServiceLevel(ctx context.Context, zone string, cluster *SKSCluster) error {
	if err := validateOperationParams(cluster, "update"); err != nil {
		return err
	}

	resp, err := c.UpgradeSksClusterServiceLevelWithResponse(apiv2.WithZone(ctx, zone), *cluster.ID)
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
