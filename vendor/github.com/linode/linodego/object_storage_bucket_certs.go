package linodego

import (
	"context"
)

// Deprecated: Please use ObjectStorageBucketCertV2 for all new implementations.
type ObjectStorageBucketCert struct {
	SSL bool `json:"ssl"`
}

type ObjectStorageBucketCertV2 struct {
	SSL *bool `json:"ssl"`
}

type ObjectStorageBucketCertUploadOptions struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// UploadObjectStorageBucketCert uploads a TLS/SSL Cert to be used with an Object Storage Bucket.
//
// Deprecated: Please use UploadObjectStorageBucketCertV2 for all new implementations.
func (c *Client) UploadObjectStorageBucketCert(
	ctx context.Context,
	clusterOrRegionID, bucket string,
	opts ObjectStorageBucketCertUploadOptions,
) (*ObjectStorageBucketCert, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	return doPOSTRequest[ObjectStorageBucketCert](ctx, c, e, opts)
}

// GetObjectStorageBucketCert gets an ObjectStorageBucketCert
//
// Deprecated: Please use GetObjectStorageBucketCertV2 for all new implementations.
func (c *Client) GetObjectStorageBucketCert(ctx context.Context, clusterOrRegionID, bucket string) (*ObjectStorageBucketCert, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	return doGETRequest[ObjectStorageBucketCert](ctx, c, e)
}

// UploadObjectStorageBucketCertV2 uploads a TLS/SSL Cert to be used with an Object Storage Bucket.
func (c *Client) UploadObjectStorageBucketCertV2(
	ctx context.Context,
	clusterOrRegionID, bucket string,
	opts ObjectStorageBucketCertUploadOptions,
) (*ObjectStorageBucketCertV2, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	return doPOSTRequest[ObjectStorageBucketCertV2](ctx, c, e, opts)
}

// GetObjectStorageBucketCertV2 gets an ObjectStorageBucketCert
func (c *Client) GetObjectStorageBucketCertV2(ctx context.Context, clusterOrRegionID, bucket string) (*ObjectStorageBucketCertV2, error) {
	e := formatAPIPath("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	return doGETRequest[ObjectStorageBucketCertV2](ctx, c, e)
}

// DeleteObjectStorageBucketCert deletes an ObjectStorageBucketCert
func (c *Client) DeleteObjectStorageBucketCert(ctx context.Context, clusterOrRegionID, bucket string) error {
	e := formatAPIPath("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	return doDELETERequest(ctx, c, e)
}
