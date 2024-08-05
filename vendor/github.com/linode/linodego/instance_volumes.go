package linodego

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// InstanceVolumesPagedResponse represents a paginated InstanceVolume API response
type InstanceVolumesPagedResponse struct {
	*PageOptions
	Data []Volume `json:"data"`
}

// endpoint gets the endpoint URL for InstanceVolume
<<<<<<< HEAD
func (InstanceVolumesPagedResponse) endpointWithID(c *Client, id int) string {
<<<<<<< HEAD
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
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
	if err != nil {
		panic(err)
	}
	return endpoint
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func (InstanceVolumesPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.InstanceVolumes.endpointWithID(id)
	if err != nil {
		panic(err)
	}
	return endpoint
=======
func (InstanceVolumesPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("linode/instances/%d/volumes", id)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

func (resp *InstanceVolumesPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(InstanceVolumesPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*InstanceVolumesPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListInstanceVolumes lists InstanceVolumes
func (c *Client) ListInstanceVolumes(ctx context.Context, linodeID int, opts *ListOptions) ([]Volume, error) {
	response := InstanceVolumesPagedResponse{}
<<<<<<< HEAD
	err := c.listHelperWithID(ctx, &response, linodeID, opts)

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	err := c.listHelperWithID(ctx, &response, linodeID, opts)

=======
	err := c.listHelper(ctx, &response, opts, linodeID)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
