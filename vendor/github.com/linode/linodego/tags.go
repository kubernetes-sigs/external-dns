package linodego

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Tag represents a Tag object
type Tag struct {
	Label string `json:"label"`
}

// TaggedObject represents a Tagged Object object
type TaggedObject struct {
	Type    string          `json:"type"`
	RawData json.RawMessage `json:"data"`
	Data    any             `json:"-"`
}

// SortedObjects currently only includes Instances
type SortedObjects struct {
	Instances     []Instance
	LKEClusters   []LKECluster
	Domains       []Domain
	Volumes       []Volume
	NodeBalancers []NodeBalancer
	/*
		StackScripts  []Stackscript
	*/
}

// TaggedObjectList are a list of TaggedObjects, as returning by ListTaggedObjects
type TaggedObjectList []TaggedObject

// TagCreateOptions fields are those accepted by CreateTag
type TagCreateOptions struct {
	Label   string `json:"label"`
	Linodes []int  `json:"linodes,omitempty"`
	// @TODO is this implemented?
	LKEClusters   []int `json:"lke_clusters,omitempty"`
	Domains       []int `json:"domains,omitempty"`
	Volumes       []int `json:"volumes,omitempty"`
	NodeBalancers []int `json:"nodebalancers,omitempty"`
}

// GetCreateOptions converts a Tag to TagCreateOptions for use in CreateTag
func (i Tag) GetCreateOptions() (o TagCreateOptions) {
	o.Label = i.Label
	return
}

// TagsPagedResponse represents a paginated Tag API response
type TagsPagedResponse struct {
	*PageOptions
	Data []Tag `json:"data"`
}

// endpoint gets the endpoint URL for Tag
func (TagsPagedResponse) endpoint(_ ...any) string {
	return "tags"
}

func (resp *TagsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(TagsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*TagsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// TaggedObjectsPagedResponse represents a paginated Tag API response
type TaggedObjectsPagedResponse struct {
	*PageOptions
	Data []TaggedObject `json:"data"`
}

// endpoint gets the endpoint URL for Tag
func (TaggedObjectsPagedResponse) endpoint(ids ...any) string {
	id := url.PathEscape(ids[0].(string))
	return fmt.Sprintf("tags/%s", id)
}

func (resp *TaggedObjectsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(TaggedObjectsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*TaggedObjectsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListTags lists Tags
func (c *Client) ListTags(ctx context.Context, opts *ListOptions) ([]Tag, error) {
	response := TagsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixData stores an object of the type defined by Type in Data using RawData
func (i *TaggedObject) fixData() (*TaggedObject, error) {
	switch i.Type {
	case "linode":
		obj := Instance{}
		if err := json.Unmarshal(i.RawData, &obj); err != nil {
			return nil, err
		}
		i.Data = obj
	case "lke_cluster":
		obj := LKECluster{}
		if err := json.Unmarshal(i.RawData, &obj); err != nil {
			return nil, err
		}
		i.Data = obj
	case "nodebalancer":
		obj := NodeBalancer{}
		if err := json.Unmarshal(i.RawData, &obj); err != nil {
			return nil, err
		}
		i.Data = obj
	case "domain":
		obj := Domain{}
		if err := json.Unmarshal(i.RawData, &obj); err != nil {
			return nil, err
		}
		i.Data = obj
	case "volume":
		obj := Volume{}
		if err := json.Unmarshal(i.RawData, &obj); err != nil {
			return nil, err
		}
		i.Data = obj
	}

	return i, nil
}

// ListTaggedObjects lists Tagged Objects
func (c *Client) ListTaggedObjects(ctx context.Context, label string, opts *ListOptions) (TaggedObjectList, error) {
	response := TaggedObjectsPagedResponse{}
<<<<<<< HEAD
	err := c.listHelperWithID(ctx, &response, label, opts)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if err != nil {
		return nil, err
	}

	for i := range response.Data {
		if _, err := response.Data[i].fixData(); err != nil {
			return nil, err
		}
	}
	return response.Data, nil
}

// SortedObjects converts a list of TaggedObjects into a Sorted Objects struct, for easier access
func (t TaggedObjectList) SortedObjects() (SortedObjects, error) {
	so := SortedObjects{}

	for _, o := range t {
		switch o.Type {
		case "linode":
			if instance, ok := o.Data.(Instance); ok {
				so.Instances = append(so.Instances, instance)
			} else {
				return so, errors.New("expected an Instance when Type was \"linode\"")
			}
		case "lke_cluster":
			if lkeCluster, ok := o.Data.(LKECluster); ok {
				so.LKEClusters = append(so.LKEClusters, lkeCluster)
			} else {
				return so, errors.New("expected an LKECluster when Type was \"lke_cluster\"")
			}
		case "domain":
			if domain, ok := o.Data.(Domain); ok {
				so.Domains = append(so.Domains, domain)
			} else {
				return so, errors.New("expected a Domain when Type was \"domain\"")
			}
		case "volume":
			if volume, ok := o.Data.(Volume); ok {
				so.Volumes = append(so.Volumes, volume)
			} else {
				return so, errors.New("expected an Volume when Type was \"volume\"")
			}
		case "nodebalancer":
			if nodebalancer, ok := o.Data.(NodeBalancer); ok {
				so.NodeBalancers = append(so.NodeBalancers, nodebalancer)
			} else {
				return so, errors.New("expected an NodeBalancer when Type was \"nodebalancer\"")
			}
		}
	}
	return so, nil
}

// CreateTag creates a Tag
func (c *Client) CreateTag(ctx context.Context, createOpts TagCreateOptions) (*Tag, error) {
	var body string
	e, err := c.Tags.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		return nil, err
	}

	for i := range response.Data {
		if _, err := response.Data[i].fixData(); err != nil {
			return nil, err
		}
	}
	return response.Data, nil
}

// SortedObjects converts a list of TaggedObjects into a Sorted Objects struct, for easier access
func (t TaggedObjectList) SortedObjects() (SortedObjects, error) {
	so := SortedObjects{}

	for _, o := range t {
		switch o.Type {
		case "linode":
			if instance, ok := o.Data.(Instance); ok {
				so.Instances = append(so.Instances, instance)
			} else {
				return so, errors.New("expected an Instance when Type was \"linode\"")
			}
		case "lke_cluster":
			if lkeCluster, ok := o.Data.(LKECluster); ok {
				so.LKEClusters = append(so.LKEClusters, lkeCluster)
			} else {
				return so, errors.New("expected an LKECluster when Type was \"lke_cluster\"")
			}
		case "domain":
			if domain, ok := o.Data.(Domain); ok {
				so.Domains = append(so.Domains, domain)
			} else {
				return so, errors.New("expected a Domain when Type was \"domain\"")
			}
		case "volume":
			if volume, ok := o.Data.(Volume); ok {
				so.Volumes = append(so.Volumes, volume)
			} else {
				return so, errors.New("expected an Volume when Type was \"volume\"")
			}
		case "nodebalancer":
			if nodebalancer, ok := o.Data.(NodeBalancer); ok {
				so.NodeBalancers = append(so.NodeBalancers, nodebalancer)
			} else {
				return so, errors.New("expected an NodeBalancer when Type was \"nodebalancer\"")
			}
		}
	}
	return so, nil
}

// CreateTag creates a Tag
func (c *Client) CreateTag(ctx context.Context, createOpts TagCreateOptions) (*Tag, error) {
	var body string
	e, err := c.Tags.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
<<<<<<< HEAD

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return nil, err
	}

	for i := range response.Data {
		if _, err := response.Data[i].fixData(); err != nil {
			return nil, err
		}
	}
	return response.Data, nil
}

// SortedObjects converts a list of TaggedObjects into a Sorted Objects struct, for easier access
func (t TaggedObjectList) SortedObjects() (SortedObjects, error) {
	so := SortedObjects{}

	for _, o := range t {
		switch o.Type {
		case "linode":
			if instance, ok := o.Data.(Instance); ok {
				so.Instances = append(so.Instances, instance)
			} else {
				return so, errors.New("expected an Instance when Type was \"linode\"")
			}
		case "lke_cluster":
			if lkeCluster, ok := o.Data.(LKECluster); ok {
				so.LKEClusters = append(so.LKEClusters, lkeCluster)
			} else {
				return so, errors.New("expected an LKECluster when Type was \"lke_cluster\"")
			}
		case "domain":
			if domain, ok := o.Data.(Domain); ok {
				so.Domains = append(so.Domains, domain)
			} else {
				return so, errors.New("expected a Domain when Type was \"domain\"")
			}
		case "volume":
			if volume, ok := o.Data.(Volume); ok {
				so.Volumes = append(so.Volumes, volume)
			} else {
				return so, errors.New("expected an Volume when Type was \"volume\"")
			}
		case "nodebalancer":
			if nodebalancer, ok := o.Data.(NodeBalancer); ok {
				so.NodeBalancers = append(so.NodeBalancers, nodebalancer)
			} else {
				return so, errors.New("expected an NodeBalancer when Type was \"nodebalancer\"")
			}
		}
	}
	return so, nil
}

// CreateTag creates a Tag
func (c *Client) CreateTag(ctx context.Context, createOpts TagCreateOptions) (*Tag, error) {
	var body string
	e, err := c.Tags.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
<<<<<<< HEAD

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}

	for i := range response.Data {
		if _, err := response.Data[i].fixData(); err != nil {
			return nil, err
		}
	}
	return response.Data, nil
}

// SortedObjects converts a list of TaggedObjects into a Sorted Objects struct, for easier access
func (t TaggedObjectList) SortedObjects() (SortedObjects, error) {
	so := SortedObjects{}

	for _, o := range t {
		switch o.Type {
		case "linode":
			if instance, ok := o.Data.(Instance); ok {
				so.Instances = append(so.Instances, instance)
			} else {
				return so, errors.New("expected an Instance when Type was \"linode\"")
			}
		case "lke_cluster":
			if lkeCluster, ok := o.Data.(LKECluster); ok {
				so.LKEClusters = append(so.LKEClusters, lkeCluster)
			} else {
				return so, errors.New("expected an LKECluster when Type was \"lke_cluster\"")
			}
		case "domain":
			if domain, ok := o.Data.(Domain); ok {
				so.Domains = append(so.Domains, domain)
			} else {
				return so, errors.New("expected a Domain when Type was \"domain\"")
			}
		case "volume":
			if volume, ok := o.Data.(Volume); ok {
				so.Volumes = append(so.Volumes, volume)
			} else {
				return so, errors.New("expected an Volume when Type was \"volume\"")
			}
		case "nodebalancer":
			if nodebalancer, ok := o.Data.(NodeBalancer); ok {
				so.NodeBalancers = append(so.NodeBalancers, nodebalancer)
			} else {
				return so, errors.New("expected an NodeBalancer when Type was \"nodebalancer\"")
			}
		}
	}
	return so, nil
}

// CreateTag creates a Tag
func (c *Client) CreateTag(ctx context.Context, createOpts TagCreateOptions) (*Tag, error) {
	var body string
	e, err := c.Tags.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
<<<<<<< HEAD

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	err := c.listHelperWithID(ctx, &response, label, opts)

=======
	label = url.PathEscape(label)
	err := c.listHelper(ctx, &response, opts, label)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}

	for i := range response.Data {
		if _, err := response.Data[i].fixData(); err != nil {
			return nil, err
		}
	}
	return response.Data, nil
}

// SortedObjects converts a list of TaggedObjects into a Sorted Objects struct, for easier access
func (t TaggedObjectList) SortedObjects() (SortedObjects, error) {
	so := SortedObjects{}

	for _, o := range t {
		switch o.Type {
		case "linode":
			if instance, ok := o.Data.(Instance); ok {
				so.Instances = append(so.Instances, instance)
			} else {
				return so, errors.New("expected an Instance when Type was \"linode\"")
			}
		case "lke_cluster":
			if lkeCluster, ok := o.Data.(LKECluster); ok {
				so.LKEClusters = append(so.LKEClusters, lkeCluster)
			} else {
				return so, errors.New("expected an LKECluster when Type was \"lke_cluster\"")
			}
		case "domain":
			if domain, ok := o.Data.(Domain); ok {
				so.Domains = append(so.Domains, domain)
			} else {
				return so, errors.New("expected a Domain when Type was \"domain\"")
			}
		case "volume":
			if volume, ok := o.Data.(Volume); ok {
				so.Volumes = append(so.Volumes, volume)
			} else {
				return so, errors.New("expected an Volume when Type was \"volume\"")
			}
		case "nodebalancer":
			if nodebalancer, ok := o.Data.(NodeBalancer); ok {
				so.NodeBalancers = append(so.NodeBalancers, nodebalancer)
			} else {
				return so, errors.New("expected an NodeBalancer when Type was \"nodebalancer\"")
			}
		}
	}
	return so, nil
}

// CreateTag creates a Tag
func (c *Client) CreateTag(ctx context.Context, opts TagCreateOptions) (*Tag, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	req := c.R(ctx).SetResult(&Tag{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))

=======
	e := "tags"
	req := c.R(ctx).SetResult(&Tag{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}
	return r.Result().(*Tag), nil
}

// DeleteTag deletes the Tag with the specified id
func (c *Client) DeleteTag(ctx context.Context, label string) error {
	label = url.PathEscape(label)
	e := fmt.Sprintf("tags/%s", label)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
