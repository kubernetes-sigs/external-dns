package linodego

import (
	"context"
)

// InstanceVolumesPagedResponse represents a paginated InstanceVolume API response
type InstanceVolumesPagedResponse struct {
	*PageOptions
	Data []Volume `json:"data"`
}

// endpoint gets the endpoint URL for InstanceVolume
func (InstanceVolumesPagedResponse) endpointWithID(c *Client, id int) string {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	endpoint, err := c.InstanceVolumes.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceVolumes when processing paginated InstanceVolume responses
func (resp *InstanceVolumesPagedResponse) appendData(r *InstanceVolumesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceVolumes lists InstanceVolumes
func (c *Client) ListInstanceVolumes(ctx context.Context, linodeID int, opts *ListOptions) ([]Volume, error) {
	response := InstanceVolumesPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
||||||| parent of 5ce8c7613 (update vendored files)
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
=======
	endpoint, err := c.InstanceVolumes.endpointWithParams(id)
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceVolumes when processing paginated InstanceVolume responses
func (resp *InstanceVolumesPagedResponse) appendData(r *InstanceVolumesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceVolumes lists InstanceVolumes
func (c *Client) ListInstanceVolumes(ctx context.Context, linodeID int, opts *ListOptions) ([]Volume, error) {
	response := InstanceVolumesPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
<<<<<<< HEAD

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
||||||| parent of 6b7ce455e (update vendored files)
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
=======
	endpoint, err := c.InstanceVolumes.endpointWithParams(id)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceVolumes when processing paginated InstanceVolume responses
func (resp *InstanceVolumesPagedResponse) appendData(r *InstanceVolumesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceVolumes lists InstanceVolumes
func (c *Client) ListInstanceVolumes(ctx context.Context, linodeID int, opts *ListOptions) ([]Volume, error) {
	response := InstanceVolumesPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
<<<<<<< HEAD

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
||||||| parent of 4d7e5ad26 (update vendored files)
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
=======
	endpoint, err := c.InstanceVolumes.endpointWithParams(id)
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceVolumes when processing paginated InstanceVolume responses
func (resp *InstanceVolumesPagedResponse) appendData(r *InstanceVolumesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceVolumes lists InstanceVolumes
func (c *Client) ListInstanceVolumes(ctx context.Context, linodeID int, opts *ListOptions) ([]Volume, error) {
	response := InstanceVolumesPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
<<<<<<< HEAD

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
