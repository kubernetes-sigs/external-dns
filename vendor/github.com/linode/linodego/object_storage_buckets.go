package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/linode/linodego/internal/parseabletime"
)

// ObjectStorageBucket represents a ObjectStorage object
type ObjectStorageBucket struct {
	Label string `json:"label"`

	// Deprecated: The 'Cluster' field has been deprecated in favor of the 'Region' field.
	// For example, a Cluster value of `us-mia-1` will translate to a Region value of `us-mia`.
	//
	// This is necessary because there are now multiple Object Storage clusters to a region.
	//
	// NOTE: The 'Cluster' field will always return a value similar to `<REGION>-1` (e.g., `us-mia-1`)
	// for backward compatibility purposes.
	Cluster string `json:"cluster"`
	Region  string `json:"region"`

	S3Endpoint   string                    `json:"s3_endpoint"`
	EndpointType ObjectStorageEndpointType `json:"endpoint_type"`
	Created      *time.Time                `json:"-"`
	Hostname     string                    `json:"hostname"`
	Objects      int                       `json:"objects"`
	Size         int                       `json:"size"`
}

// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

type ObjectStorageBucketAccessV2 struct {
	ACL         ObjectStorageACL `json:"acl"`
	ACLXML      string           `json:"acl_xml"`
	CorsEnabled *bool            `json:"cors_enabled"`
	CorsXML     *string          `json:"cors_xml"`
}

// ObjectStorageBucketContent holds the content of an ObjectStorageBucket
type ObjectStorageBucketContent struct {
	Data        []ObjectStorageBucketContentData `json:"data"`
	IsTruncated bool                             `json:"is_truncated"`
	NextMarker  *string                          `json:"next_marker"`
}

// ObjectStorageBucketContentData holds the data of the content of an ObjectStorageBucket
type ObjectStorageBucketContentData struct {
	Etag         string     `json:"etag"`
	LastModified *time.Time `json:"last_modified"`
	Name         string     `json:"name"`
	Owner        string     `json:"owner"`
	Size         int        `json:"size"`
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
	// Deprecated: The 'Cluster' field has been deprecated.
	//
	// Going forward, the 'Region' field will be the supported way to designate where an
	// Object Storage Bucket should be created. For example, a 'Cluster' value of `us-mia-1`
	// will translate to a Region value of `us-mia`.
	Cluster string `json:"cluster,omitempty"`
	Region  string `json:"region,omitempty"`

	Label        string                    `json:"label"`
	S3Endpoint   string                    `json:"s3_endpoint,omitempty"`
	EndpointType ObjectStorageEndpointType `json:"endpoint_type,omitempty"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketListContentsParams fields are the query parameters for ListObjectStorageBucketContents
type ObjectStorageBucketListContentsParams struct {
	Marker    *string
	Delimiter *string
	Prefix    *string
	PageSize  *int
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

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	return getPaginatedResults[ObjectStorageBucket](ctx, c, "object-storage/buckets", opts)
}

// ListObjectStorageBucketsInCluster lists all ObjectStorageBuckets of a cluster
func (c *Client) ListObjectStorageBucketsInCluster(ctx context.Context, opts *ListOptions, clusterOrRegionID string) ([]ObjectStorageBucket, error) {
	return getPaginatedResults[ObjectStorageBucket](ctx, c, formatAPIPath("object-storage/buckets/%s", clusterOrRegionID), opts)
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterOrRegionID, label string) (*ObjectStorageBucket, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s", clusterOrRegionID, label)
	return doGETRequest[ObjectStorageBucket](ctx, c, e)
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, opts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	return doPOSTRequest[ObjectStorageBucket](ctx, c, "object-storage/buckets", opts)
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
//
// Deprecated: use GetObjectStorageBucketAccessV2 for new implementations
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterOrRegionID, label string) (*ObjectStorageBucketAccess, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/access", clusterOrRegionID, label)
	return doGETRequest[ObjectStorageBucketAccess](ctx, c, e)
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterOrRegionID, label string, opts ObjectStorageBucketUpdateAccessOptions) error {
	e := formatAPIPath("object-storage/buckets/%s/%s/access", clusterOrRegionID, label)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// GetObjectStorageBucketAccessV2 gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccessV2(ctx context.Context, clusterOrRegionID, label string) (*ObjectStorageBucketAccessV2, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/access", clusterOrRegionID, label)
	return doGETRequest[ObjectStorageBucketAccessV2](ctx, c, e)
}

// DeleteObjectStorageBucket deletes the ObjectStorageBucket with the specified label
func (c *Client) DeleteObjectStorageBucket(ctx context.Context, clusterOrRegionID, label string) error {
	e := formatAPIPath("object-storage/buckets/%s/%s", clusterOrRegionID, label)
	return doDELETERequest(ctx, c, e)
}

// ListObjectStorageBucketContents lists the contents of the specified ObjectStorageBucket
func (c *Client) ListObjectStorageBucketContents(
	ctx context.Context,
	clusterOrRegionID, label string,
	params *ObjectStorageBucketListContentsParams,
) (*ObjectStorageBucketContent, error) {
	basePath := formatAPIPath("object-storage/buckets/%s/%s/object-list", clusterOrRegionID, label)

	queryString := ""

	if params != nil {
		values, err := query.Values(params)
		if err != nil {
			return nil, fmt.Errorf("failed to encode query params: %w", err)
		}

		queryString = "?" + values.Encode()
	}

	e := basePath + queryString

	return doGETRequest[ObjectStorageBucketContent](ctx, c, e)
}
