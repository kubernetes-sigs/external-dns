package linodego

import (
	"context"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
)

type ObjectStorageObjectURLCreateOptions struct {
	Name               string `json:"name"`
	Method             string `json:"method"`
	ContentType        string `json:"content_type,omit_empty"`
	ContentDisposition string `json:"content_disposition,omit_empty"`
	ExpiresIn          *int   `json:"expires_in,omit_empty"`
}

type ObjectStorageObjectURL struct {
	URL    string `json:"url"`
	Exists bool   `json:"exists"`
}

type ObjectStorageObjectACLConfig struct {
	ACL    string `json:"acl"`
	ACLXML string `json:"acl_xml"`
}

type ObjectStorageObjectACLConfigUpdateOptions struct {
	Name string `json:"name"`
	ACL  string `json:"acl"`
}

func (c *Client) CreateObjectStorageObjectURL(ctx context.Context, clusterID, label string, options ObjectStorageObjectURLCreateOptions) (*ObjectStorageObjectURL, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageObjectURL{})
	e = fmt.Sprintf("%s/%s/%s/object-url", e, clusterID, label)

	if bodyData, err := json.Marshal(options); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.SetBody(body).Post(e))
	return r.Result().(*ObjectStorageObjectURL), err
}

func (c *Client) GetObjectStorageObjectACLConfig(ctx context.Context, clusterID, label, object string) (*ObjectStorageObjectACLConfig, error) {
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageObjectACLConfig{})
	e = fmt.Sprintf("%s/%s/%s/object-acl?name=%s", e, clusterID, label, object)

	r, err := coupleAPIErrors(req.Get(e))
	return r.Result().(*ObjectStorageObjectACLConfig), err
}

func (c *Client) UpdateObjectStorageObjectACLConfig(ctx context.Context, clusterID, label string, options ObjectStorageObjectACLConfigUpdateOptions) (*ObjectStorageObjectACLConfig, error) {
	var body string
	e, err := c.ObjectStorageBuckets.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&ObjectStorageObjectACLConfig{})
	e = fmt.Sprintf("%s/%s/%s/object-acl", e, clusterID, label)

	if bodyData, err := json.Marshal(options); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.SetBody(body).Put(e))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"net/url"
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

type ObjectStorageObjectACLConfig struct {
	ACL    string `json:"acl"`
	ACLXML string `json:"acl_xml"`
}

type ObjectStorageObjectACLConfigUpdateOptions struct {
	Name string `json:"name"`
	ACL  string `json:"acl"`
}

func (c *Client) CreateObjectStorageObjectURL(ctx context.Context, objectID, label string, opts ObjectStorageObjectURLCreateOptions) (*ObjectStorageObjectURL, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	label = url.PathEscape(label)
	objectID = url.PathEscape(objectID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/object-url", objectID, label)
	req := c.R(ctx).SetResult(&ObjectStorageObjectURL{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	return r.Result().(*ObjectStorageObjectURL), err
}

func (c *Client) GetObjectStorageObjectACLConfig(ctx context.Context, objectID, label, object string) (*ObjectStorageObjectACLConfig, error) {
	label = url.PathEscape(label)
	object = url.QueryEscape(object)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/object-acl?name=%s", objectID, label, object)
	req := c.R(ctx).SetResult(&ObjectStorageObjectACLConfig{})
	r, err := coupleAPIErrors(req.Get(e))
	return r.Result().(*ObjectStorageObjectACLConfig), err
}

func (c *Client) UpdateObjectStorageObjectACLConfig(ctx context.Context, objectID, label string, opts ObjectStorageObjectACLConfigUpdateOptions) (*ObjectStorageObjectACLConfig, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	label = url.PathEscape(label)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/object-acl", objectID, label)
	req := c.R(ctx).SetResult(&ObjectStorageObjectACLConfig{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}

>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return r.Result().(*ObjectStorageObjectACLConfig), err
}
