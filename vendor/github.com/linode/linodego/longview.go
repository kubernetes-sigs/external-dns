package linodego

import (
	"context"
	"fmt"
)

// LongviewClient represents a LongviewClient object
type LongviewClient struct {
	ID int `json:"id"`
	// UpdatedStr string `json:"updated"`
	// Updated *time.Time `json:"-"`
}

// LongviewClientsPagedResponse represents a paginated LongviewClient API response
type LongviewClientsPagedResponse struct {
	*PageOptions
	Data []LongviewClient `json:"data"`
}

// endpoint gets the endpoint URL for LongviewClient
func (LongviewClientsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LongviewClients.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LongviewClients when processing paginated LongviewClient responses
func (resp *LongviewClientsPagedResponse) appendData(r *LongviewClientsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListLongviewClients lists LongviewClients
func (c *Client) ListLongviewClients(ctx context.Context, opts *ListOptions) ([]LongviewClient, error) {
	response := LongviewClientsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLongviewClient gets the template with the provided ID
func (c *Client) GetLongviewClient(ctx context.Context, id string) (*LongviewClient, error) {
	e, err := c.LongviewClients.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := c.R(ctx).SetResult(&LongviewClient{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient), nil
}
