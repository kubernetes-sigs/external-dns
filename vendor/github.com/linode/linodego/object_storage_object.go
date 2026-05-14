package linodego

import (
	"context"
)

type ObjectStorageObjectURLCreateOptions struct {
	Name               string `json:"name"`
	Method             string `json:"method"`
	ContentType        string `json:"content_type,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	ExpiresIn          *int   `json:"expires_in,omitempty"`
}

type ObjectStorageObjectURL struct {
	URL    string `json:"url"`
	Exists bool   `json:"exists"`
}

// Deprecated: Please use ObjectStorageObjectACLConfigV2 for all new implementations.
type ObjectStorageObjectACLConfig struct {
	ACL    string `json:"acl"`
	ACLXML string `json:"acl_xml"`
}

type ObjectStorageObjectACLConfigV2 struct {
	ACL    *string `json:"acl"`
	ACLXML *string `json:"acl_xml"`
}

type ObjectStorageObjectACLConfigUpdateOptions struct {
	Name string `json:"name"`
	ACL  string `json:"acl"`
}

func (c *Client) CreateObjectStorageObjectURL(
	ctx context.Context,
	objectID, label string,
	opts ObjectStorageObjectURLCreateOptions,
) (*ObjectStorageObjectURL, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/object-url", objectID, label)
	return doPOSTRequest[ObjectStorageObjectURL](ctx, c, e, opts)
}

// Deprecated: use GetObjectStorageObjectACLConfigV2 for new implementations
func (c *Client) GetObjectStorageObjectACLConfig(ctx context.Context, objectID, label, object string) (*ObjectStorageObjectACLConfig, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/object-acl?name=%s", objectID, label, object)
	return doGETRequest[ObjectStorageObjectACLConfig](ctx, c, e)
}

// Deprecated: use UpdateObjectStorageObjectACLConfigV2 for new implementations
func (c *Client) UpdateObjectStorageObjectACLConfig(
	ctx context.Context,
	objectID, label string,
	opts ObjectStorageObjectACLConfigUpdateOptions,
) (*ObjectStorageObjectACLConfig, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/object-acl", objectID, label)
	return doPUTRequest[ObjectStorageObjectACLConfig](ctx, c, e, opts)
}

func (c *Client) GetObjectStorageObjectACLConfigV2(ctx context.Context, objectID, label, object string) (*ObjectStorageObjectACLConfigV2, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/object-acl?name=%s", objectID, label, object)
	return doGETRequest[ObjectStorageObjectACLConfigV2](ctx, c, e)
}

func (c *Client) UpdateObjectStorageObjectACLConfigV2(
	ctx context.Context,
	objectID, label string,
	opts ObjectStorageObjectACLConfigUpdateOptions,
) (*ObjectStorageObjectACLConfigV2, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/object-acl", objectID, label)
	return doPUTRequest[ObjectStorageObjectACLConfigV2](ctx, c, e, opts)
}
