package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// LongviewClient represents a LongviewClient object
type LongviewClient struct {
	ID          int        `json:"id"`
	APIKey      string     `json:"api_key"`
	Created     *time.Time `json:"-"`
	InstallCode string     `json:"install_code"`
	Label       string     `json:"label"`
	Updated     *time.Time `json:"-"`
	Apps        struct {
		Apache any `json:"apache"`
		MySQL  any `json:"mysql"`
		NginX  any `json:"nginx"`
	} `json:"apps"`
}

// LongviewClientCreateOptions is an options struct used when Creating a Longview Client
type LongviewClientCreateOptions struct {
	Label string `json:"label"`
}

// LongviewClientCreateOptions is an options struct used when Updating a Longview Client
type LongviewClientUpdateOptions struct {
	Label string `json:"label"`
}

// LongviewPlan represents a Longview Plan object
type LongviewPlan struct {
	ID              string `json:"id"`
	Label           string `json:"label"`
	ClientsIncluded int    `json:"clients_included"`
	Price           struct {
		Hourly  float64 `json:"hourly"`
		Monthly float64 `json:"monthly"`
	} `json:"price"`
}

// LongviewPlanUpdateOptions is an options struct used when Updating a Longview Plan
type LongviewPlanUpdateOptions struct {
	LongviewSubscription string `json:"longview_subscription"`
}

// LongviewClientsPagedResponse represents a paginated LongviewClient API response
type LongviewClientsPagedResponse struct {
	*PageOptions
	Data []LongviewClient `json:"data"`
}

// endpoint gets the endpoint URL for LongviewClient
func (LongviewClientsPagedResponse) endpoint(_ ...any) string {
	return "longview/clients"
}

func (resp *LongviewClientsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LongviewClientsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LongviewClientsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLongviewClient gets the template with the provided ID
func (c *Client) GetLongviewClient(ctx context.Context, clientID int) (*LongviewClient, error) {
	e := fmt.Sprintf("longview/clients/%d", clientID)
	r, err := c.R(ctx).SetResult(&LongviewClient{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient), nil
}

// CreateLongviewClient creates a Longview Client
func (c *Client) CreateLongviewClient(ctx context.Context, opts LongviewClientCreateOptions) (*LongviewClient, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "longview/clients"
	req := c.R(ctx).SetResult(&LongviewClient{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient), nil
}

// DeleteLongviewClient deletes a Longview Client
func (c *Client) DeleteLongviewClient(ctx context.Context, clientID int) error {
	e := fmt.Sprintf("longview/clients/%d", clientID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

// UpdateLongviewClient updates a Longview Client
func (c *Client) UpdateLongviewClient(ctx context.Context, clientID int, opts LongviewClientUpdateOptions) (*LongviewClient, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("longview/clients/%d", clientID)
	req := c.R(ctx).SetResult(&LongviewClient{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient), nil
}

// GetLongviewPlan gets the template with the provided ID
func (c *Client) GetLongviewPlan(ctx context.Context) (*LongviewPlan, error) {
	e := "longview/plan"
	r, err := c.R(ctx).SetResult(&LongviewPlan{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewPlan), nil
}

// UpdateLongviewPlan updates a Longview Plan
func (c *Client) UpdateLongviewPlan(ctx context.Context, opts LongviewPlanUpdateOptions) (*LongviewPlan, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "longview/plan"
	req := c.R(ctx).SetResult(&LongviewPlan{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewPlan), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LongviewClient) UnmarshalJSON(b []byte) error {
	type Mask LongviewClient

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
