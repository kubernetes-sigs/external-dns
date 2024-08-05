package linodego

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
)

// Deprecated: LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk = LKENodePoolDisk

// Deprecated: LKEClusterPoolAutoscaler represents an AutoScaler configuration
type LKEClusterPoolAutoscaler = LKENodePoolAutoscaler

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}
=======
// Deprecated: LKEClusterPoolLinode represents a LKEClusterPoolLinode object
type LKEClusterPoolLinode = LKENodePoolLinode
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)

// Deprecated: LKEClusterPool represents a LKEClusterPool object
type LKEClusterPool = LKENodePool

// Deprecated: LKEClusterPoolCreateOptions fields are those accepted by CreateLKEClusterPool
type LKEClusterPoolCreateOptions = LKENodePoolCreateOptions

// Deprecated: LKEClusterPoolUpdateOptions fields are those accepted by UpdateLKEClusterPool
type LKEClusterPoolUpdateOptions = LKENodePoolUpdateOptions

// Deprecated: LKEClusterPoolsPagedResponse represents a paginated LKEClusterPool API response
type LKEClusterPoolsPagedResponse = LKENodePoolsPagedResponse

// Deprecated: ListLKEClusterPools lists LKEClusterPools
func (c *Client) ListLKEClusterPools(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterPool, error) {
	return c.ListLKENodePools(ctx, clusterID, opts)
}

// Deprecated: GetLKEClusterPool gets the lkeClusterPool with the provided ID
func (c *Client) GetLKEClusterPool(ctx context.Context, clusterID, id int) (*LKEClusterPool, error) {
	return c.GetLKENodePool(ctx, clusterID, id)
}

// Deprecated: CreateLKEClusterPool creates a LKEClusterPool
func (c *Client) CreateLKEClusterPool(ctx context.Context, clusterID int, createOpts LKEClusterPoolCreateOptions) (*LKEClusterPool, error) {
	return c.CreateLKENodePool(ctx, clusterID, createOpts)
}

// Deprecated: UpdateLKEClusterPool updates the LKEClusterPool with the specified id
func (c *Client) UpdateLKEClusterPool(ctx context.Context, clusterID, id int, updateOpts LKEClusterPoolUpdateOptions) (*LKEClusterPool, error) {
	return c.UpdateLKENodePool(ctx, clusterID, id, updateOpts)
}

// Deprecated: DeleteLKEClusterPool deletes the LKEClusterPool with the specified id
func (c *Client) DeleteLKEClusterPool(ctx context.Context, clusterID, id int) error {
	return c.DeleteLKENodePool(ctx, clusterID, id)
}

// Deprecated: DeleteLKEClusterPoolNode deletes a given node from a cluster pool
func (c *Client) DeleteLKEClusterPoolNode(ctx context.Context, clusterID int, id string) error {
<<<<<<< HEAD
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/nodes/%s", e, clusterID, id)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
// LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}

type LKEClusterPoolAutoscaler struct {
	Enabled bool `json:"enabled"`
	Min     int  `json:"min"`
	Max     int  `json:"max"`
}

>>>>>>> 5ce8c7613 (update vendored files)
// LKEClusterPoolLinode represents a LKEClusterPoolLinode object
type LKEClusterPoolLinode struct {
	ID         string          `json:"id"`
	InstanceID int             `json:"instance_id"`
	Status     LKELinodeStatus `json:"status"`
}

// LKEClusterPool represents a LKEClusterPool object
type LKEClusterPool struct {
	ID      int                    `json:"id"`
	Count   int                    `json:"count"`
	Type    string                 `json:"type"`
	Disks   []LKEClusterPoolDisk   `json:"disks"`
	Linodes []LKEClusterPoolLinode `json:"nodes"`
	Tags    []string               `json:"tags"`

	Autoscaler LKEClusterPoolAutoscaler `json:"autoscaler"`
}

// LKEClusterPoolCreateOptions fields are those accepted by CreateLKEClusterPool
type LKEClusterPoolCreateOptions struct {
	Count int                  `json:"count"`
	Type  string               `json:"type"`
	Disks []LKEClusterPoolDisk `json:"disks"`
	Tags  []string             `json:"tags"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// LKEClusterPoolUpdateOptions fields are those accepted by UpdateLKEClusterPool
type LKEClusterPoolUpdateOptions struct {
	Count int       `json:"count,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// GetCreateOptions converts a LKEClusterPool to LKEClusterPoolCreateOptions for
// use in CreateLKEClusterPool
func (l LKEClusterPool) GetCreateOptions() (o LKEClusterPoolCreateOptions) {
	o.Count = l.Count
	o.Disks = l.Disks
	o.Tags = l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// GetUpdateOptions converts a LKEClusterPool to LKEClusterPoolUpdateOptions for use in UpdateLKEClusterPool
func (l LKEClusterPool) GetUpdateOptions() (o LKEClusterPoolUpdateOptions) {
	o.Count = l.Count
	o.Tags = &l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// LKEClusterPoolsPagedResponse represents a paginated LKEClusterPool API response
type LKEClusterPoolsPagedResponse struct {
	*PageOptions
	Data []LKEClusterPool `json:"data"`
}

// endpointWithID gets the endpoint URL for InstanceConfigs of a given Instance
func (LKEClusterPoolsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterPools.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterPools when processing paginated LKEClusterPool responses
func (resp *LKEClusterPoolsPagedResponse) appendData(r *LKEClusterPoolsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusterPools lists LKEClusterPools
func (c *Client) ListLKEClusterPools(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterPool, error) {
	response := LKEClusterPoolsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetLKEClusterPool gets the lkeClusterPool with the provided ID
func (c *Client) GetLKEClusterPool(ctx context.Context, clusterID, id int) (*LKEClusterPool, error) {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterPool{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterPool), nil
}

// CreateLKEClusterPool creates a LKEClusterPool
func (c *Client) CreateLKEClusterPool(ctx context.Context, clusterID int, createOpts LKEClusterPoolCreateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// UpdateLKEClusterPool updates the LKEClusterPool with the specified id
func (c *Client) UpdateLKEClusterPool(ctx context.Context, clusterID, id int, updateOpts LKEClusterPoolUpdateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// DeleteLKEClusterPool deletes the LKEClusterPool with the specified id
func (c *Client) DeleteLKEClusterPool(ctx context.Context,
	clusterID, id int) error {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// DeleteLKEClusterPoolNode deletes a given node from a cluster pool
func (c *Client) DeleteLKEClusterPoolNode(ctx context.Context, clusterID int, id string) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/nodes/%s", e, clusterID, id)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
// LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}

type LKEClusterPoolAutoscaler struct {
	Enabled bool `json:"enabled"`
	Min     int  `json:"min"`
	Max     int  `json:"max"`
}

>>>>>>> 6b7ce455e (update vendored files)
// LKEClusterPoolLinode represents a LKEClusterPoolLinode object
type LKEClusterPoolLinode struct {
	ID         string          `json:"id"`
	InstanceID int             `json:"instance_id"`
	Status     LKELinodeStatus `json:"status"`
}

// LKEClusterPool represents a LKEClusterPool object
type LKEClusterPool struct {
	ID      int                    `json:"id"`
	Count   int                    `json:"count"`
	Type    string                 `json:"type"`
	Disks   []LKEClusterPoolDisk   `json:"disks"`
	Linodes []LKEClusterPoolLinode `json:"nodes"`
	Tags    []string               `json:"tags"`

	Autoscaler LKEClusterPoolAutoscaler `json:"autoscaler"`
}

// LKEClusterPoolCreateOptions fields are those accepted by CreateLKEClusterPool
type LKEClusterPoolCreateOptions struct {
	Count int                  `json:"count"`
	Type  string               `json:"type"`
	Disks []LKEClusterPoolDisk `json:"disks"`
	Tags  []string             `json:"tags"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// LKEClusterPoolUpdateOptions fields are those accepted by UpdateLKEClusterPool
type LKEClusterPoolUpdateOptions struct {
	Count int       `json:"count,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// GetCreateOptions converts a LKEClusterPool to LKEClusterPoolCreateOptions for
// use in CreateLKEClusterPool
func (l LKEClusterPool) GetCreateOptions() (o LKEClusterPoolCreateOptions) {
	o.Count = l.Count
	o.Disks = l.Disks
	o.Tags = l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// GetUpdateOptions converts a LKEClusterPool to LKEClusterPoolUpdateOptions for use in UpdateLKEClusterPool
func (l LKEClusterPool) GetUpdateOptions() (o LKEClusterPoolUpdateOptions) {
	o.Count = l.Count
	o.Tags = &l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// LKEClusterPoolsPagedResponse represents a paginated LKEClusterPool API response
type LKEClusterPoolsPagedResponse struct {
	*PageOptions
	Data []LKEClusterPool `json:"data"`
}

// endpointWithID gets the endpoint URL for InstanceConfigs of a given Instance
func (LKEClusterPoolsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterPools.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterPools when processing paginated LKEClusterPool responses
func (resp *LKEClusterPoolsPagedResponse) appendData(r *LKEClusterPoolsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusterPools lists LKEClusterPools
func (c *Client) ListLKEClusterPools(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterPool, error) {
	response := LKEClusterPoolsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetLKEClusterPool gets the lkeClusterPool with the provided ID
func (c *Client) GetLKEClusterPool(ctx context.Context, clusterID, id int) (*LKEClusterPool, error) {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterPool{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterPool), nil
}

// CreateLKEClusterPool creates a LKEClusterPool
func (c *Client) CreateLKEClusterPool(ctx context.Context, clusterID int, createOpts LKEClusterPoolCreateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// UpdateLKEClusterPool updates the LKEClusterPool with the specified id
func (c *Client) UpdateLKEClusterPool(ctx context.Context, clusterID, id int, updateOpts LKEClusterPoolUpdateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// DeleteLKEClusterPool deletes the LKEClusterPool with the specified id
func (c *Client) DeleteLKEClusterPool(ctx context.Context,
	clusterID, id int) error {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// DeleteLKEClusterPoolNode deletes a given node from a cluster pool
func (c *Client) DeleteLKEClusterPoolNode(ctx context.Context, clusterID int, id string) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/nodes/%s", e, clusterID, id)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}

type LKEClusterPoolAutoscaler struct {
	Enabled bool `json:"enabled"`
	Min     int  `json:"min"`
	Max     int  `json:"max"`
}

>>>>>>> 4d7e5ad26 (update vendored files)
// LKEClusterPoolLinode represents a LKEClusterPoolLinode object
type LKEClusterPoolLinode struct {
	ID         string          `json:"id"`
	InstanceID int             `json:"instance_id"`
	Status     LKELinodeStatus `json:"status"`
}

// LKEClusterPool represents a LKEClusterPool object
type LKEClusterPool struct {
	ID      int                    `json:"id"`
	Count   int                    `json:"count"`
	Type    string                 `json:"type"`
	Disks   []LKEClusterPoolDisk   `json:"disks"`
	Linodes []LKEClusterPoolLinode `json:"nodes"`
	Tags    []string               `json:"tags"`

	Autoscaler LKEClusterPoolAutoscaler `json:"autoscaler"`
}

// LKEClusterPoolCreateOptions fields are those accepted by CreateLKEClusterPool
type LKEClusterPoolCreateOptions struct {
	Count int                  `json:"count"`
	Type  string               `json:"type"`
	Disks []LKEClusterPoolDisk `json:"disks"`
	Tags  []string             `json:"tags"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// LKEClusterPoolUpdateOptions fields are those accepted by UpdateLKEClusterPool
type LKEClusterPoolUpdateOptions struct {
	Count int       `json:"count,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`

	Autoscaler *LKEClusterPoolAutoscaler `json:"autoscaler,omitempty"`
}

// GetCreateOptions converts a LKEClusterPool to LKEClusterPoolCreateOptions for
// use in CreateLKEClusterPool
func (l LKEClusterPool) GetCreateOptions() (o LKEClusterPoolCreateOptions) {
	o.Count = l.Count
	o.Disks = l.Disks
	o.Tags = l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// GetUpdateOptions converts a LKEClusterPool to LKEClusterPoolUpdateOptions for use in UpdateLKEClusterPool
func (l LKEClusterPool) GetUpdateOptions() (o LKEClusterPoolUpdateOptions) {
	o.Count = l.Count
	o.Tags = &l.Tags
	o.Autoscaler = &l.Autoscaler
	return
}

// LKEClusterPoolsPagedResponse represents a paginated LKEClusterPool API response
type LKEClusterPoolsPagedResponse struct {
	*PageOptions
	Data []LKEClusterPool `json:"data"`
}

// endpointWithID gets the endpoint URL for InstanceConfigs of a given Instance
func (LKEClusterPoolsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.LKEClusterPools.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LKEClusterPools when processing paginated LKEClusterPool responses
func (resp *LKEClusterPoolsPagedResponse) appendData(r *LKEClusterPoolsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLKEClusterPools lists LKEClusterPools
func (c *Client) ListLKEClusterPools(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterPool, error) {
	response := LKEClusterPoolsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, clusterID, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetLKEClusterPool gets the lkeClusterPool with the provided ID
func (c *Client) GetLKEClusterPool(ctx context.Context, clusterID, id int) (*LKEClusterPool, error) {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&LKEClusterPool{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LKEClusterPool), nil
}

// CreateLKEClusterPool creates a LKEClusterPool
func (c *Client) CreateLKEClusterPool(ctx context.Context, clusterID int, createOpts LKEClusterPoolCreateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// UpdateLKEClusterPool updates the LKEClusterPool with the specified id
func (c *Client) UpdateLKEClusterPool(ctx context.Context, clusterID, id int, updateOpts LKEClusterPoolUpdateOptions) (*LKEClusterPool, error) {
	var body string
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&LKEClusterPool{})

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
	return r.Result().(*LKEClusterPool), nil
}

// DeleteLKEClusterPool deletes the LKEClusterPool with the specified id
func (c *Client) DeleteLKEClusterPool(ctx context.Context,
	clusterID, id int) error {
	e, err := c.LKEClusterPools.endpointWithParams(clusterID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// DeleteLKEClusterPoolNode deletes a given node from a cluster pool
func (c *Client) DeleteLKEClusterPoolNode(ctx context.Context, clusterID int, id string) error {
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/nodes/%s", e, clusterID, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	e, err := c.LKEClusters.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/nodes/%s", e, clusterID, id)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
=======
	return c.DeleteLKENodePoolNode(ctx, clusterID, id)
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"fmt"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
	"fmt"
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)

// Deprecated: LKEClusterPoolDisk represents a Node disk in an LKEClusterPool object
type LKEClusterPoolDisk = LKENodePoolDisk

// Deprecated: LKEClusterPoolAutoscaler represents an AutoScaler configuration
type LKEClusterPoolAutoscaler = LKENodePoolAutoscaler

// Deprecated: LKEClusterPoolLinode represents a LKEClusterPoolLinode object
type LKEClusterPoolLinode = LKENodePoolLinode

// Deprecated: LKEClusterPool represents a LKEClusterPool object
type LKEClusterPool = LKENodePool

// Deprecated: LKEClusterPoolCreateOptions fields are those accepted by CreateLKEClusterPool
type LKEClusterPoolCreateOptions = LKENodePoolCreateOptions

// Deprecated: LKEClusterPoolUpdateOptions fields are those accepted by UpdateLKEClusterPool
type LKEClusterPoolUpdateOptions = LKENodePoolUpdateOptions

// Deprecated: LKEClusterPoolsPagedResponse represents a paginated LKEClusterPool API response
type LKEClusterPoolsPagedResponse = LKENodePoolsPagedResponse

// Deprecated: ListLKEClusterPools lists LKEClusterPools
func (c *Client) ListLKEClusterPools(ctx context.Context, clusterID int, opts *ListOptions) ([]LKEClusterPool, error) {
	return c.ListLKENodePools(ctx, clusterID, opts)
}

// Deprecated: GetLKEClusterPool gets the lkeClusterPool with the provided ID
func (c *Client) GetLKEClusterPool(ctx context.Context, clusterID, id int) (*LKEClusterPool, error) {
	return c.GetLKENodePool(ctx, clusterID, id)
}

// Deprecated: CreateLKEClusterPool creates a LKEClusterPool
func (c *Client) CreateLKEClusterPool(ctx context.Context, clusterID int, createOpts LKEClusterPoolCreateOptions) (*LKEClusterPool, error) {
	return c.CreateLKENodePool(ctx, clusterID, createOpts)
}

// Deprecated: UpdateLKEClusterPool updates the LKEClusterPool with the specified id
func (c *Client) UpdateLKEClusterPool(ctx context.Context, clusterID, id int, updateOpts LKEClusterPoolUpdateOptions) (*LKEClusterPool, error) {
	return c.UpdateLKENodePool(ctx, clusterID, id, updateOpts)
}

// Deprecated: DeleteLKEClusterPool deletes the LKEClusterPool with the specified id
func (c *Client) DeleteLKEClusterPool(ctx context.Context, clusterID, id int) error {
	return c.DeleteLKENodePool(ctx, clusterID, id)
}

<<<<<<< HEAD
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
=======
// Deprecated: DeleteLKEClusterPoolNode deletes a given node from a cluster pool
func (c *Client) DeleteLKEClusterPoolNode(ctx context.Context, clusterID int, id string) error {
	return c.DeleteLKENodePoolNode(ctx, clusterID, id)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
