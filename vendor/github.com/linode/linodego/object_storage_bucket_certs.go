package linodego

import (
	"context"
	"encoding/json"
<<<<<<< HEAD
)

type ObjectStorageBucketCert struct {
	SSL bool `json:"ssl"`
}

type ObjectStorageBucketCertUploadOptions struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// UploadObjectStorageBucketCert uploads a TLS/SSL Cert to be used with an Object Storage Bucket.
func (c *Client) UploadObjectStorageBucketCert(ctx context.Context, clusterID, bucket string, uploadOpts ObjectStorageBucketCertUploadOptions) (*ObjectStorageBucketCert, error) {
	e, err := c.ObjectStorageBucketCerts.endpointWithParams(clusterID, bucket)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(uploadOpts)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucketCert{}).SetBody(string(body)).Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// GetObjectStorageBucketCert gets an ObjectStorageBucketCert
func (c *Client) GetObjectStorageBucketCert(ctx context.Context, clusterID, bucket string) (*ObjectStorageBucketCert, error) {
	e, err := c.ObjectStorageBucketCerts.endpointWithParams(clusterID, bucket)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&ObjectStorageBucketCert{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// DeleteObjectStorageBucketCert deletes an ObjectStorageBucketCert
func (c *Client) DeleteObjectStorageBucketCert(ctx context.Context, clusterID, bucket string) error {
	e, err := c.ObjectStorageBucketCerts.endpointWithParams(clusterID, bucket)
	if err != nil {
		return err
	}

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/url"
)

type ObjectStorageBucketCert struct {
	SSL bool `json:"ssl"`
}

type ObjectStorageBucketCertUploadOptions struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// UploadObjectStorageBucketCert uploads a TLS/SSL Cert to be used with an Object Storage Bucket.
func (c *Client) UploadObjectStorageBucketCert(ctx context.Context, clusterID, bucket string, opts ObjectStorageBucketCertUploadOptions) (*ObjectStorageBucketCert, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	clusterID = url.PathEscape(clusterID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterID, bucket)
	req := c.R(ctx).SetResult(&ObjectStorageBucketCert{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// GetObjectStorageBucketCert gets an ObjectStorageBucketCert
func (c *Client) GetObjectStorageBucketCert(ctx context.Context, clusterID, bucket string) (*ObjectStorageBucketCert, error) {
	clusterID = url.PathEscape(clusterID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterID, bucket)
	req := c.R(ctx).SetResult(&ObjectStorageBucketCert{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// DeleteObjectStorageBucketCert deletes an ObjectStorageBucketCert
func (c *Client) DeleteObjectStorageBucketCert(ctx context.Context, clusterID, bucket string) error {
	clusterID = url.PathEscape(clusterID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterID, bucket)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return err
}
