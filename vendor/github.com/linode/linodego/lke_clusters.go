package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// LKEClusterStatus represents the status of an LKECluster
type LKEClusterStatus string

// LKEClusterStatus enums start with LKECluster
const (
	LKEClusterReady    LKEClusterStatus = "ready"
	LKEClusterNotReady LKEClusterStatus = "not_ready"
)

// LKECluster represents a LKECluster object
type LKECluster struct {
<<<<<<< HEAD
<<<<<<< HEAD
	ID           int                    `json:"id"`
	Created      *time.Time             `json:"-"`
	Updated      *time.Time             `json:"-"`
	Label        string                 `json:"label"`
	Region       string                 `json:"region"`
	Status       LKEClusterStatus       `json:"status"`
	K8sVersion   string                 `json:"k8s_version"`
	Tags         []string               `json:"tags"`
	ControlPlane LKEClusterControlPlane `json:"control_plane"`
}

// LKEClusterCreateOptions fields are those accepted by CreateLKECluster
type LKEClusterCreateOptions struct {
	NodePools    []LKENodePoolCreateOptions `json:"node_pools"`
	Label        string                     `json:"label"`
	Region       string                     `json:"region"`
	K8sVersion   string                     `json:"k8s_version"`
	Tags         []string                   `json:"tags,omitempty"`
	ControlPlane *LKEClusterControlPlane    `json:"control_plane,omitempty"`
}

// LKEClusterUpdateOptions fields are those accepted by UpdateLKECluster
type LKEClusterUpdateOptions struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	K8sVersion string    `json:"k8s_version,omitempty"`
	Label      string    `json:"label,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	K8sVersion string    `json:"k8s_version,omitempty"`
	Label      string    `json:"label,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
=======
	K8sVersion   string                  `json:"k8s_version,omitempty"`
	Label        string                  `json:"label,omitempty"`
	Tags         *[]string               `json:"tags,omitempty"`
	ControlPlane *LKEClusterControlPlane `json:"control_plane,omitempty"`
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
}

// LKEClusterAPIEndpoint fields are those returned by ListLKEClusterAPIEndpoints
type LKEClusterAPIEndpoint struct {
	Endpoint string `json:"endpoint"`
}

// LKEClusterKubeconfig fields are those returned by GetLKEClusterKubeconfig
type LKEClusterKubeconfig struct {
	KubeConfig string `json:"kubeconfig"`
}

// LKEClusterDashboard fields are those returned by GetLKEClusterDashboard
type LKEClusterDashboard struct {
	URL string `json:"url"`
}

// LKEClusterControlPlane fields contained within the `control_plane` attribute of an LKE cluster.
type LKEClusterControlPlane struct {
	HighAvailability bool `json:"high_availability"`
}

// LKEVersion fields are those returned by GetLKEVersion
type LKEVersion struct {
	ID string `json:"id"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LKECluster) UnmarshalJSON(b []byte) error {
	type Mask LKECluster

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetCreateOptions converts a LKECluster to LKEClusterCreateOptions for use in CreateLKECluster
func (i LKECluster) GetCreateOptions() (o LKEClusterCreateOptions) {
	o.Label = i.Label
	o.Region = i.Region
	o.K8sVersion = i.K8sVersion
	o.Tags = i.Tags
	o.ControlPlane = &i.ControlPlane
	// @TODO copy NodePools?
	return
}

// GetUpdateOptions converts a LKECluster to LKEClusterUpdateOptions for use in UpdateLKECluster
func (i LKECluster) GetUpdateOptions() (o LKEClusterUpdateOptions) {
	o.K8sVersion = i.K8sVersion
	o.Label = i.Label
	o.Tags = &i.Tags
	o.ControlPlane = &i.ControlPlane
	return
}

// LKEClustersPagedResponse represents a paginated LKECluster API response
type LKEClustersPagedResponse struct {
	*PageOptions
	Data []LKECluster `json:"data"`
}

// LKEClusterAPIEndpointsPagedResponse represents a paginated LKEClusterAPIEndpoints API response
type LKEClusterAPIEndpointsPagedResponse struct {
	*PageOptions
	Data []LKEClusterAPIEndpoint `json:"data"`
}

// LKEVersionsPagedResponse represents a paginated LKEVersion API response
type LKEVersionsPagedResponse struct {
	*PageOptions
	Data []LKEVersion `json:"data"`
}

// endpoint gets the endpoint URL for LKECluster
func (LKEClustersPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEClusters.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusters when processing paginated LKECluster responses
func (resp *LKEClustersPagedResponse) appendData(r *LKEClustersPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// endpoint gets the endpoint URL for LKEVersion
func (LKEVersionsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEVersions.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// endpoint gets the endpoint URL for LKEClusterAPIEndpoints
func (LKEClusterAPIEndpointsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterAPIEndpoints.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterAPIEndpoints when processing paginated LKEClusterAPIEndpoints responses
func (resp *LKEClusterAPIEndpointsPagedResponse) appendData(r *LKEClusterAPIEndpointsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// appendData appends LKEVersions when processing paginated LKEVersion responses
func (resp *LKEVersionsPagedResponse) appendData(r *LKEVersionsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusters lists LKEClusters
func (c *Client) ListLKEClusters(ctx context.Context, opts *ListOptions) ([]LKECluster, error) {
	response := LKEClustersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKECluster gets the lkeCluster with the provided ID
func (c *Client) GetLKECluster(ctx context.Context, id int) (*LKECluster, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKECluster{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// CreateLKECluster creates a LKECluster
func (c *Client) CreateLKECluster(ctx context.Context, createOpts LKEClusterCreateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// UpdateLKECluster updates the LKECluster with the specified id
func (c *Client) UpdateLKECluster(ctx context.Context, id int, updateOpts LKEClusterUpdateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKECluster deletes the LKECluster with the specified id
func (c *Client) DeleteLKECluster(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// ListLKEClusterAPIEndpoints gets the API Endpoint for the LKE Cluster specified
func (c *Client) ListLKEClusterAPIEndpoints(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterAPIEndpoint, error) {
	response := LKEClusterAPIEndpointsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKEClusterKubeconfig gets the Kubeconfig for the LKE Cluster specified
func (c *Client) GetLKEClusterKubeconfig(ctx context.Context, id int) (*LKEClusterKubeconfig, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/kubeconfig", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterKubeconfig{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterKubeconfig), nil
}

// GetLKEClusterDashboard gets information about the dashboard for an LKE cluster
func (c *Client) GetLKEClusterDashboard(ctx context.Context, id int) (*LKEClusterDashboard, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/dashboard", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterDashboard{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterDashboard), nil
}

// RecycleLKEClusterNodes recycles all nodes in all pools of the specified LKE Cluster.
func (c *Client) RecycleLKEClusterNodes(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/recycle", e, id)
	_, err = coupleAPIErrors(c.R(ctx).Post(e))
	return err
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
||||||| parent of 5ce8c7613 (update vendored files)
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
=======
	K8sVersion string    `json:"k8s_version,omitempty"`
	Label      string    `json:"label,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
>>>>>>> 5ce8c7613 (update vendored files)
}

// LKEClusterAPIEndpoint fields are those returned by ListLKEClusterAPIEndpoints
type LKEClusterAPIEndpoint struct {
	Endpoint string `json:"endpoint"`
}

// LKEClusterKubeconfig fields are those returned by GetLKEClusterKubeconfig
type LKEClusterKubeconfig struct {
	KubeConfig string `json:"kubeconfig"`
}

// LKEVersion fields are those returned by GetLKEVersion
type LKEVersion struct {
	ID string `json:"id"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LKECluster) UnmarshalJSON(b []byte) error {
	type Mask LKECluster

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetCreateOptions converts a LKECluster to LKEClusterCreateOptions for use in CreateLKECluster
func (i LKECluster) GetCreateOptions() (o LKEClusterCreateOptions) {
	o.Label = i.Label
	o.Region = i.Region
	o.K8sVersion = i.K8sVersion
	o.Tags = i.Tags
	// @TODO copy NodePools?
	return
}

// GetUpdateOptions converts a LKECluster to LKEClusterUpdateOptions for use in UpdateLKECluster
func (i LKECluster) GetUpdateOptions() (o LKEClusterUpdateOptions) {
	o.K8sVersion = i.K8sVersion
	o.Label = i.Label
	o.Tags = &i.Tags
	return
}

// LKEClustersPagedResponse represents a paginated LKECluster API response
type LKEClustersPagedResponse struct {
	*PageOptions
	Data []LKECluster `json:"data"`
}

// LKEClusterAPIEndpointsPagedResponse represents a paginated LKEClusterAPIEndpoints API response
type LKEClusterAPIEndpointsPagedResponse struct {
	*PageOptions
	Data []LKEClusterAPIEndpoint `json:"data"`
}

// LKEVersionsPagedResponse represents a paginated LKEVersion API response
type LKEVersionsPagedResponse struct {
	*PageOptions
	Data []LKEVersion `json:"data"`
}

// endpoint gets the endpoint URL for LKECluster
func (LKEClustersPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEClusters.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusters when processing paginated LKECluster responses
func (resp *LKEClustersPagedResponse) appendData(r *LKEClustersPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// endpoint gets the endpoint URL for LKEVersion
func (LKEVersionsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEVersions.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// endpoint gets the endpoint URL for LKEClusterAPIEndpoints
func (LKEClusterAPIEndpointsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterAPIEndpoints.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterAPIEndpoints when processing paginated LKEClusterAPIEndpoints responses
func (resp *LKEClusterAPIEndpointsPagedResponse) appendData(r *LKEClusterAPIEndpointsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// appendData appends LKEVersions when processing paginated LKEVersion responses
func (resp *LKEVersionsPagedResponse) appendData(r *LKEVersionsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusters lists LKEClusters
func (c *Client) ListLKEClusters(ctx context.Context, opts *ListOptions) ([]LKECluster, error) {
	response := LKEClustersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKECluster gets the lkeCluster with the provided ID
func (c *Client) GetLKECluster(ctx context.Context, id int) (*LKECluster, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKECluster{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// CreateLKECluster creates a LKECluster
func (c *Client) CreateLKECluster(ctx context.Context, createOpts LKEClusterCreateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// UpdateLKECluster updates the LKECluster with the specified id
func (c *Client) UpdateLKECluster(ctx context.Context, id int, updateOpts LKEClusterUpdateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKECluster deletes the LKECluster with the specified id
func (c *Client) DeleteLKECluster(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// ListLKEClusterAPIEndpoints gets the API Endpoint for the LKE Cluster specified
func (c *Client) ListLKEClusterAPIEndpoints(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterAPIEndpoint, error) {
	response := LKEClusterAPIEndpointsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKEClusterKubeconfig gets the Kubeconfig for the LKE Cluster specified
func (c *Client) GetLKEClusterKubeconfig(ctx context.Context, id int) (*LKEClusterKubeconfig, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/kubeconfig", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterKubeconfig{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterKubeconfig), nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// RecycleLKEClusterNodes recycles all nodes in all pools of the specified LKE Cluster.
func (c *Client) RecycleLKEClusterNodes(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/recycle", e, id)
	_, err = coupleAPIErrors(c.R(ctx).Post(e))
	return err
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
||||||| parent of 6b7ce455e (update vendored files)
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
=======
	K8sVersion string    `json:"k8s_version,omitempty"`
	Label      string    `json:"label,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
>>>>>>> 6b7ce455e (update vendored files)
}

// LKEClusterAPIEndpoint fields are those returned by ListLKEClusterAPIEndpoints
type LKEClusterAPIEndpoint struct {
	Endpoint string `json:"endpoint"`
}

// LKEClusterKubeconfig fields are those returned by GetLKEClusterKubeconfig
type LKEClusterKubeconfig struct {
	KubeConfig string `json:"kubeconfig"`
}

// LKEVersion fields are those returned by GetLKEVersion
type LKEVersion struct {
	ID string `json:"id"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LKECluster) UnmarshalJSON(b []byte) error {
	type Mask LKECluster

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetCreateOptions converts a LKECluster to LKEClusterCreateOptions for use in CreateLKECluster
func (i LKECluster) GetCreateOptions() (o LKEClusterCreateOptions) {
	o.Label = i.Label
	o.Region = i.Region
	o.K8sVersion = i.K8sVersion
	o.Tags = i.Tags
	// @TODO copy NodePools?
	return
}

// GetUpdateOptions converts a LKECluster to LKEClusterUpdateOptions for use in UpdateLKECluster
func (i LKECluster) GetUpdateOptions() (o LKEClusterUpdateOptions) {
	o.K8sVersion = i.K8sVersion
	o.Label = i.Label
	o.Tags = &i.Tags
	return
}

// LKEClustersPagedResponse represents a paginated LKECluster API response
type LKEClustersPagedResponse struct {
	*PageOptions
	Data []LKECluster `json:"data"`
}

// LKEClusterAPIEndpointsPagedResponse represents a paginated LKEClusterAPIEndpoints API response
type LKEClusterAPIEndpointsPagedResponse struct {
	*PageOptions
	Data []LKEClusterAPIEndpoint `json:"data"`
}

// LKEVersionsPagedResponse represents a paginated LKEVersion API response
type LKEVersionsPagedResponse struct {
	*PageOptions
	Data []LKEVersion `json:"data"`
}

// endpoint gets the endpoint URL for LKECluster
func (LKEClustersPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEClusters.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusters when processing paginated LKECluster responses
func (resp *LKEClustersPagedResponse) appendData(r *LKEClustersPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// endpoint gets the endpoint URL for LKEVersion
func (LKEVersionsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEVersions.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// endpoint gets the endpoint URL for LKEClusterAPIEndpoints
func (LKEClusterAPIEndpointsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterAPIEndpoints.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterAPIEndpoints when processing paginated LKEClusterAPIEndpoints responses
func (resp *LKEClusterAPIEndpointsPagedResponse) appendData(r *LKEClusterAPIEndpointsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// appendData appends LKEVersions when processing paginated LKEVersion responses
func (resp *LKEVersionsPagedResponse) appendData(r *LKEVersionsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusters lists LKEClusters
func (c *Client) ListLKEClusters(ctx context.Context, opts *ListOptions) ([]LKECluster, error) {
	response := LKEClustersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKECluster gets the lkeCluster with the provided ID
func (c *Client) GetLKECluster(ctx context.Context, id int) (*LKECluster, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKECluster{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// CreateLKECluster creates a LKECluster
func (c *Client) CreateLKECluster(ctx context.Context, createOpts LKEClusterCreateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// UpdateLKECluster updates the LKECluster with the specified id
func (c *Client) UpdateLKECluster(ctx context.Context, id int, updateOpts LKEClusterUpdateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKECluster deletes the LKECluster with the specified id
func (c *Client) DeleteLKECluster(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// ListLKEClusterAPIEndpoints gets the API Endpoint for the LKE Cluster specified
func (c *Client) ListLKEClusterAPIEndpoints(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterAPIEndpoint, error) {
	response := LKEClusterAPIEndpointsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKEClusterKubeconfig gets the Kubeconfig for the LKE Cluster specified
func (c *Client) GetLKEClusterKubeconfig(ctx context.Context, id int) (*LKEClusterKubeconfig, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/kubeconfig", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterKubeconfig{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterKubeconfig), nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// RecycleLKEClusterNodes recycles all nodes in all pools of the specified LKE Cluster.
func (c *Client) RecycleLKEClusterNodes(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/recycle", e, id)
	_, err = coupleAPIErrors(c.R(ctx).Post(e))
	return err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
||||||| parent of 4d7e5ad26 (update vendored files)
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
=======
	K8sVersion string    `json:"k8s_version,omitempty"`
	Label      string    `json:"label,omitempty"`
	Tags       *[]string `json:"tags,omitempty"`
>>>>>>> 4d7e5ad26 (update vendored files)
}

// LKEClusterAPIEndpoint fields are those returned by ListLKEClusterAPIEndpoints
type LKEClusterAPIEndpoint struct {
	Endpoint string `json:"endpoint"`
}

// LKEClusterKubeconfig fields are those returned by GetLKEClusterKubeconfig
type LKEClusterKubeconfig struct {
	KubeConfig string `json:"kubeconfig"`
}

// LKEVersion fields are those returned by GetLKEVersion
type LKEVersion struct {
	ID string `json:"id"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LKECluster) UnmarshalJSON(b []byte) error {
	type Mask LKECluster

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetCreateOptions converts a LKECluster to LKEClusterCreateOptions for use in CreateLKECluster
func (i LKECluster) GetCreateOptions() (o LKEClusterCreateOptions) {
	o.Label = i.Label
	o.Region = i.Region
	o.K8sVersion = i.K8sVersion
	o.Tags = i.Tags
	// @TODO copy NodePools?
	return
}

// GetUpdateOptions converts a LKECluster to LKEClusterUpdateOptions for use in UpdateLKECluster
func (i LKECluster) GetUpdateOptions() (o LKEClusterUpdateOptions) {
	o.K8sVersion = i.K8sVersion
	o.Label = i.Label
	o.Tags = &i.Tags
	return
}

// LKEClustersPagedResponse represents a paginated LKECluster API response
type LKEClustersPagedResponse struct {
	*PageOptions
	Data []LKECluster `json:"data"`
}

// LKEClusterAPIEndpointsPagedResponse represents a paginated LKEClusterAPIEndpoints API response
type LKEClusterAPIEndpointsPagedResponse struct {
	*PageOptions
	Data []LKEClusterAPIEndpoint `json:"data"`
}

// LKEVersionsPagedResponse represents a paginated LKEVersion API response
type LKEVersionsPagedResponse struct {
	*PageOptions
	Data []LKEVersion `json:"data"`
}

// endpoint gets the endpoint URL for LKECluster
func (LKEClustersPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEClusters.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusters when processing paginated LKECluster responses
func (resp *LKEClustersPagedResponse) appendData(r *LKEClustersPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// endpoint gets the endpoint URL for LKEVersion
func (LKEVersionsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LKEVersions.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// endpoint gets the endpoint URL for LKEClusterAPIEndpoints
func (LKEClusterAPIEndpointsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterAPIEndpoints.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterAPIEndpoints when processing paginated LKEClusterAPIEndpoints responses
func (resp *LKEClusterAPIEndpointsPagedResponse) appendData(r *LKEClusterAPIEndpointsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// appendData appends LKEVersions when processing paginated LKEVersion responses
func (resp *LKEVersionsPagedResponse) appendData(r *LKEVersionsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusters lists LKEClusters
func (c *Client) ListLKEClusters(ctx context.Context, opts *ListOptions) ([]LKECluster, error) {
	response := LKEClustersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKECluster gets the lkeCluster with the provided ID
func (c *Client) GetLKECluster(ctx context.Context, id int) (*LKECluster, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKECluster{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// CreateLKECluster creates a LKECluster
func (c *Client) CreateLKECluster(ctx context.Context, createOpts LKEClusterCreateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// UpdateLKECluster updates the LKECluster with the specified id
func (c *Client) UpdateLKECluster(ctx context.Context, id int, updateOpts LKEClusterUpdateOptions) (*LKECluster, error) {
	var body string
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKECluster{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKECluster deletes the LKECluster with the specified id
func (c *Client) DeleteLKECluster(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// ListLKEClusterAPIEndpoints gets the API Endpoint for the LKE Cluster specified
func (c *Client) ListLKEClusterAPIEndpoints(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterAPIEndpoint, error) {
	response := LKEClusterAPIEndpointsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKEClusterKubeconfig gets the Kubeconfig for the LKE Cluster specified
func (c *Client) GetLKEClusterKubeconfig(ctx context.Context, id int) (*LKEClusterKubeconfig, error) {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/kubeconfig", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterKubeconfig{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterKubeconfig), nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// RecycleLKEClusterNodes recycles all nodes in all pools of the specified LKE Cluster.
func (c *Client) RecycleLKEClusterNodes(ctx context.Context, id int) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/recycle", e, id)
	_, err = coupleAPIErrors(c.R(ctx).Post(e))
	return err
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID         int              `json:"id"`
	Created    *time.Time       `json:"-"`
	Updated    *time.Time       `json:"-"`
	Label      string           `json:"label"`
	Region     string           `json:"region"`
	Status     LKEClusterStatus `json:"status"`
	K8sVersion string           `json:"k8s_version"`
	Tags       []string         `json:"tags"`
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	ID         int              `json:"id"`
	Created    *time.Time       `json:"-"`
	Updated    *time.Time       `json:"-"`
	Label      string           `json:"label"`
	Region     string           `json:"region"`
	Status     LKEClusterStatus `json:"status"`
	K8sVersion string           `json:"k8s_version"`
	Tags       []string         `json:"tags"`
=======
	ID           int                    `json:"id"`
	Created      *time.Time             `json:"-"`
	Updated      *time.Time             `json:"-"`
	Label        string                 `json:"label"`
	Region       string                 `json:"region"`
	Status       LKEClusterStatus       `json:"status"`
	K8sVersion   string                 `json:"k8s_version"`
	Tags         []string               `json:"tags"`
	ControlPlane LKEClusterControlPlane `json:"control_plane"`
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

// LKEClusterCreateOptions fields are those accepted by CreateLKECluster
type LKEClusterCreateOptions struct {
	NodePools    []LKENodePoolCreateOptions `json:"node_pools"`
	Label        string                     `json:"label"`
	Region       string                     `json:"region"`
	K8sVersion   string                     `json:"k8s_version"`
	Tags         []string                   `json:"tags,omitempty"`
	ControlPlane *LKEClusterControlPlane    `json:"control_plane,omitempty"`
}

// LKEClusterUpdateOptions fields are those accepted by UpdateLKECluster
type LKEClusterUpdateOptions struct {
	K8sVersion   string                  `json:"k8s_version,omitempty"`
	Label        string                  `json:"label,omitempty"`
	Tags         *[]string               `json:"tags,omitempty"`
	ControlPlane *LKEClusterControlPlane `json:"control_plane,omitempty"`
}

// LKEClusterAPIEndpoint fields are those returned by ListLKEClusterAPIEndpoints
type LKEClusterAPIEndpoint struct {
	Endpoint string `json:"endpoint"`
}

// LKEClusterKubeconfig fields are those returned by GetLKEClusterKubeconfig
type LKEClusterKubeconfig struct {
	KubeConfig string `json:"kubeconfig"`
}

// LKEClusterDashboard fields are those returned by GetLKEClusterDashboard
type LKEClusterDashboard struct {
	URL string `json:"url"`
}

// LKEClusterControlPlane fields contained within the `control_plane` attribute of an LKE cluster.
type LKEClusterControlPlane struct {
	HighAvailability bool `json:"high_availability"`
}

// LKEVersion fields are those returned by GetLKEVersion
type LKEVersion struct {
	ID string `json:"id"`
}

// LKEClusterRegenerateOptions fields are those accepted by RegenerateLKECluster
type LKEClusterRegenerateOptions struct {
	KubeConfig   bool `json:"kubeconfig"`
	ServiceToken bool `json:"servicetoken"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LKECluster) UnmarshalJSON(b []byte) error {
	type Mask LKECluster

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetCreateOptions converts a LKECluster to LKEClusterCreateOptions for use in CreateLKECluster
func (i LKECluster) GetCreateOptions() (o LKEClusterCreateOptions) {
	o.Label = i.Label
	o.Region = i.Region
	o.K8sVersion = i.K8sVersion
	o.Tags = i.Tags
	o.ControlPlane = &i.ControlPlane
	// @TODO copy NodePools?
	return
}

// GetUpdateOptions converts a LKECluster to LKEClusterUpdateOptions for use in UpdateLKECluster
func (i LKECluster) GetUpdateOptions() (o LKEClusterUpdateOptions) {
	o.K8sVersion = i.K8sVersion
	o.Label = i.Label
	o.Tags = &i.Tags
	o.ControlPlane = &i.ControlPlane
	return
}

// LKEVersionsPagedResponse represents a paginated LKEVersion API response
type LKEVersionsPagedResponse struct {
	*PageOptions
	Data []LKEVersion `json:"data"`
}

// endpoint gets the endpoint URL for LKEVersion
func (LKEVersionsPagedResponse) endpoint(_ ...any) string {
	return "lke/versions"
}

func (resp *LKEVersionsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LKEVersionsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LKEVersionsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListLKEVersions lists the Kubernetes versions available through LKE. This endpoint is cached by default.
func (c *Client) ListLKEVersions(ctx context.Context, opts *ListOptions) ([]LKEVersion, error) {
	response := LKEVersionsPagedResponse{}

	endpoint, err := generateListCacheURL(response.endpoint(), opts)
	if err != nil {
		return nil, err
	}

	if result := c.getCachedResponse(endpoint); result != nil {
		return result.([]LKEVersion), nil
	}

	err = c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(endpoint, response.Data, &cacheExpiryTime)

	return response.Data, nil
}

// GetLKEVersion gets details about a specific LKE Version. This endpoint is cached by default.
func (c *Client) GetLKEVersion(ctx context.Context, version string) (*LKEVersion, error) {
	version = url.PathEscape(version)
	e := fmt.Sprintf("lke/versions/%s", version)

	if result := c.getCachedResponse(e); result != nil {
		result := result.(LKEVersion)
		return &result, nil
	}

	req := c.R(ctx).SetResult(&LKEVersion{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(e, r.Result(), &cacheExpiryTime)

	return r.Result().(*LKEVersion), nil
}

// LKEClusterAPIEndpointsPagedResponse represents a paginated LKEClusterAPIEndpoints API response
type LKEClusterAPIEndpointsPagedResponse struct {
	*PageOptions
	Data []LKEClusterAPIEndpoint `json:"data"`
}

// endpoint gets the endpoint URL for LKEClusterAPIEndpointsPagedResponse
func (LKEClusterAPIEndpointsPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("lke/clusters/%d/api-endpoints", id)
}

func (resp *LKEClusterAPIEndpointsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LKEClusterAPIEndpointsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LKEClusterAPIEndpointsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListLKEClusterAPIEndpoints gets the API Endpoint for the LKE Cluster specified
func (c *Client) ListLKEClusterAPIEndpoints(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterAPIEndpoint, error) {
	response := LKEClusterAPIEndpointsPagedResponse{}
	err := c.listHelper(ctx, &response, opts, clusterID)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// LKEClustersPagedResponse represents a paginated LKECluster API response
type LKEClustersPagedResponse struct {
	*PageOptions
	Data []LKECluster `json:"data"`
}

// endpoint gets the endpoint URL for LKECluster
func (LKEClustersPagedResponse) endpoint(_ ...any) string {
	return "lke/clusters"
}

func (resp *LKEClustersPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LKEClustersPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LKEClustersPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListLKEClusters lists LKEClusters
func (c *Client) ListLKEClusters(ctx context.Context, opts *ListOptions) ([]LKECluster, error) {
	response := LKEClustersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLKECluster gets the lkeCluster with the provided ID
func (c *Client) GetLKECluster(ctx context.Context, clusterID int) (*LKECluster, error) {
	e := fmt.Sprintf("lke/clusters/%d", clusterID)
	req := c.R(ctx).SetResult(&LKECluster{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// CreateLKECluster creates a LKECluster
func (c *Client) CreateLKECluster(ctx context.Context, opts LKEClusterCreateOptions) (*LKECluster, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "lke/clusters"
	req := c.R(ctx).SetResult(&LKECluster{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// UpdateLKECluster updates the LKECluster with the specified id
func (c *Client) UpdateLKECluster(ctx context.Context, clusterID int, opts LKEClusterUpdateOptions) (*LKECluster, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("lke/clusters/%d", clusterID)
	req := c.R(ctx).SetResult(&LKECluster{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKECluster deletes the LKECluster with the specified id
func (c *Client) DeleteLKECluster(ctx context.Context, clusterID int) error {
	e := fmt.Sprintf("lke/clusters/%d", clusterID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// GetLKEClusterKubeconfig gets the Kubeconfig for the LKE Cluster specified
func (c *Client) GetLKEClusterKubeconfig(ctx context.Context, clusterID int) (*LKEClusterKubeconfig, error) {
	e := fmt.Sprintf("lke/clusters/%d/kubeconfig", clusterID)
	req := c.R(ctx).SetResult(&LKEClusterKubeconfig{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterKubeconfig), nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// GetLKEClusterDashboard gets information about the dashboard for an LKE cluster
func (c *Client) GetLKEClusterDashboard(ctx context.Context, clusterID int) (*LKEClusterDashboard, error) {
	e := fmt.Sprintf("lke/clusters/%d/dashboard", clusterID)
	req := c.R(ctx).SetResult(&LKEClusterDashboard{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterDashboard), nil
}

// RecycleLKEClusterNodes recycles all nodes in all pools of the specified LKE Cluster.
func (c *Client) RecycleLKEClusterNodes(ctx context.Context, clusterID int) error {
	e := fmt.Sprintf("lke/clusters/%d/recycle", clusterID)
	_, err := coupleAPIErrors(c.R(ctx).Post(e))
	return err
}

// RegenerateLKECluster regenerates the Kubeconfig file and/or the service account token for the specified LKE Cluster.
func (c *Client) RegenerateLKECluster(ctx context.Context, clusterID int, opts LKEClusterRegenerateOptions) (*LKECluster, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("lke/clusters/%d/regenerate", clusterID)
	req := c.R(ctx).SetResult(&LKECluster{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKECluster), nil
}

// DeleteLKEClusterServiceToken deletes and regenerate the service account token for a Cluster.
func (c *Client) DeleteLKEClusterServiceToken(ctx context.Context, clusterID int) error {
	e := fmt.Sprintf("lke/clusters/%d/servicetoken", clusterID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
