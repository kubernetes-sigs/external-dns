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

// ObjectStorageBucket represents a ObjectStorage object
type ObjectStorageBucket struct {
	Label   string `json:"label"`
	Cluster string `json:"cluster"`

	Created  *time.Time `json:"-"`
	Hostname string     `json:"hostname"`
	Objects  int        `json:"objects"`
	Size     int        `json:"size"`
}

// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	Cluster string `json:"cluster"`
	Label   string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends ObjectStorageBuckets when processing paginated ObjectStorageBucket responses
func (resp *ObjectStorageBucketsPagedResponse) appendData(r *ObjectStorageBucketsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterID, label string) (*ObjectStorageBucket, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s/%s", e, clusterID, label)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucket{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, createOpts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageBucket{})

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
	return r.Result().(*ObjectStorageBucket), nil
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterID, label string) (*ObjectStorageBucketAccess, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})

	r, err := coupleAPIErrors(
		req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterID, label string, access ObjectStorageBucketUpdateAccessOptions) error {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	bodyData, err := json.Marshal(access)
	if err != nil {
		return err
	}

	body := string(bodyData)

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))
	if err != nil {
		return err
	}

	return nil
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

>>>>>>> 5ce8c7613 (update vendored files)
// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	Cluster string `json:"cluster"`
	Label   string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends ObjectStorageBuckets when processing paginated ObjectStorageBucket responses
func (resp *ObjectStorageBucketsPagedResponse) appendData(r *ObjectStorageBucketsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterID, label string) (*ObjectStorageBucket, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s/%s", e, clusterID, label)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucket{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, createOpts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageBucket{})

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
	return r.Result().(*ObjectStorageBucket), nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterID, label string) (*ObjectStorageBucketAccess, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})

	r, err := coupleAPIErrors(
		req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterID, label string, access ObjectStorageBucketUpdateAccessOptions) error {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	bodyData, err := json.Marshal(access)
	if err != nil {
		return err
	}

	body := string(bodyData)

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))
	if err != nil {
		return err
	}

	return nil
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

>>>>>>> 6b7ce455e (update vendored files)
// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	Cluster string `json:"cluster"`
	Label   string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends ObjectStorageBuckets when processing paginated ObjectStorageBucket responses
func (resp *ObjectStorageBucketsPagedResponse) appendData(r *ObjectStorageBucketsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterID, label string) (*ObjectStorageBucket, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s/%s", e, clusterID, label)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucket{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, createOpts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageBucket{})

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
	return r.Result().(*ObjectStorageBucket), nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterID, label string) (*ObjectStorageBucketAccess, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})

	r, err := coupleAPIErrors(
		req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterID, label string, access ObjectStorageBucketUpdateAccessOptions) error {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	bodyData, err := json.Marshal(access)
	if err != nil {
		return err
	}

	body := string(bodyData)

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))
	if err != nil {
		return err
	}

	return nil
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

>>>>>>> 4d7e5ad26 (update vendored files)
// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	Cluster string `json:"cluster"`
	Label   string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends ObjectStorageBuckets when processing paginated ObjectStorageBucket responses
func (resp *ObjectStorageBucketsPagedResponse) appendData(r *ObjectStorageBucketsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterID, label string) (*ObjectStorageBucket, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s/%s", e, clusterID, label)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucket{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, createOpts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageBucket{})

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
	return r.Result().(*ObjectStorageBucket), nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterID, label string) (*ObjectStorageBucketAccess, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})

	r, err := coupleAPIErrors(
		req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterID, label string, access ObjectStorageBucketUpdateAccessOptions) error {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s/%s/access", e, clusterID, label)

	bodyData, err := json.Marshal(access)
	if err != nil {
		return err
	}

	body := string(bodyData)

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))
	if err != nil {
		return err
	}

	return nil
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	Cluster string `json:"cluster"`
	Label   string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(args ...any) string {
	endpoint := "object-storage/buckets"
	if len(args) > 0 {
		endpoint = fmt.Sprintf(endpoint+"/%s", url.PathEscape(args[0].(string)))
	}
	return endpoint
}

func (resp *ObjectStorageBucketsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(ObjectStorageBucketsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*ObjectStorageBucketsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// ListObjectStorageBucketsInCluster lists all ObjectStorageBuckets of a cluster
func (c *Client) ListObjectStorageBucketsInCluster(ctx context.Context, opts *ListOptions, clusterID string) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts, clusterID)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterID, label string) (*ObjectStorageBucket, error) {
	label = url.PathEscape(label)
	clusterID = url.PathEscape(clusterID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s", clusterID, label)
	req := c.R(ctx).SetResult(&ObjectStorageBucket{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, opts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "object-storage/buckets"
	req := c.R(ctx).SetResult(&ObjectStorageBucket{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterID, label string) (*ObjectStorageBucketAccess, error) {
	label = url.PathEscape(label)
	clusterID = url.PathEscape(clusterID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/access", clusterID, label)
	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterID, label string, opts ObjectStorageBucketUpdateAccessOptions) error {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	label = url.PathEscape(label)
	clusterID = url.PathEscape(clusterID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/access", clusterID, label)
	_, err = coupleAPIErrors(c.R(ctx).SetBody(string(body)).Post(e))
	if err != nil {
		return err
	}

	return nil
}

// DeleteObjectStorageBucket deletes the ObjectStorageBucket with the specified label
func (c *Client) DeleteObjectStorageBucket(ctx context.Context, clusterID, label string) error {
	label = url.PathEscape(label)
	e := fmt.Sprintf("object-storage/buckets/%s/%s", clusterID, label)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
