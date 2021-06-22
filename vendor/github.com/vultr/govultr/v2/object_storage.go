package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ObjectStorageService is the interface to interact with the object storage endpoints on the Vultr API.
// Link : https://www.vultr.com/api/#tag/s3
type ObjectStorageService interface {
	Create(ctx context.Context, clusterID int, label string) (*ObjectStorage, error)
	Get(ctx context.Context, id string) (*ObjectStorage, error)
	Update(ctx context.Context, id, label string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, options *ListOptions) ([]ObjectStorage, *Meta, error)

	ListCluster(ctx context.Context, options *ListOptions) ([]ObjectStorageCluster, *Meta, error)
	RegenerateKeys(ctx context.Context, id string) (*S3Keys, error)
}

// ObjectStorageServiceHandler handles interaction with the firewall rule methods for the Vultr API.
type ObjectStorageServiceHandler struct {
	client *Client
}

// ObjectStorage represents a Vultr Object Storage subscription.
type ObjectStorage struct {
	ID                   string `json:"id"`
	DateCreated          string `json:"date_created"`
	ObjectStoreClusterID int    `json:"cluster_id"`
	Region               string `json:"region"`
	Location             string `json:"location"`
	Label                string `json:"label"`
	Status               string `json:"status"`
	S3Keys
}

// S3Keys define your api access to your cluster
type S3Keys struct {
	S3Hostname  string `json:"s3_hostname"`
	S3AccessKey string `json:"s3_access_key"`
	S3SecretKey string `json:"s3_secret_key"`
}

// ObjectStorageCluster represents a Vultr Object Storage cluster.
type ObjectStorageCluster struct {
	ID       int    `json:"id"`
	Region   string `json:"region"`
	Hostname string `json:"hostname"`
	Deploy   string `json:"deploy"`
}

type objectStoragesBase struct {
	ObjectStorages []ObjectStorage `json:"object_storages"`
	Meta           *Meta           `json:"meta"`
}

type objectStorageBase struct {
	ObjectStorage *ObjectStorage `json:"object_storage"`
}

type objectStorageClustersBase struct {
	Clusters []ObjectStorageCluster `json:"clusters"`
	Meta     *Meta                  `json:"meta"`
}

type s3KeysBase struct {
	S3Credentials *S3Keys `json:"s3_credentials"`
}

//// ObjectListOptions are your optional params you have available to list data.
//type ObjectListOptions struct {
//	IncludeS3 bool
//	Label     string
//}

// Create an object storage subscription
func (o *ObjectStorageServiceHandler) Create(ctx context.Context, clusterID int, label string) (*ObjectStorage, error) {
	uri := "/v2/object-storage"

	values := RequestBody{"cluster_id": clusterID, "label": label}
	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, values)
	if err != nil {
		return nil, err
	}

	objectStorage := new(objectStorageBase)
	err = o.client.DoWithContext(ctx, req, objectStorage)
	if err != nil {
		return nil, err
	}

	return objectStorage.ObjectStorage, nil
}

// Get returns a specified object storage by the provided ID
func (o *ObjectStorageServiceHandler) Get(ctx context.Context, id string) (*ObjectStorage, error) {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	objectStorage := new(objectStorageBase)
	if err = o.client.DoWithContext(ctx, req, objectStorage); err != nil {
		return nil, err
	}

	return objectStorage.ObjectStorage, nil
}

// Update a Object Storage Subscription.
func (o *ObjectStorageServiceHandler) Update(ctx context.Context, id, label string) error {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	value := RequestBody{"label": label}
	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, value)
	if err != nil {
		return err
	}

	return o.client.DoWithContext(ctx, req, nil)
}

// Delete a object storage subscription.
func (o *ObjectStorageServiceHandler) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return o.client.DoWithContext(ctx, req, nil)
}

// List all object storage subscriptions on the current account. This includes both pending and active subscriptions.
func (o *ObjectStorageServiceHandler) List(ctx context.Context, options *ListOptions) ([]ObjectStorage, *Meta, error) {
	uri := "/v2/object-storage"

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	objectStorage := new(objectStoragesBase)
	if err = o.client.DoWithContext(ctx, req, objectStorage); err != nil {
		return nil, nil, err
	}

	return objectStorage.ObjectStorages, objectStorage.Meta, nil
}

// ListCluster returns back your object storage clusters.
// Clusters may be removed over time. The "deploy" field can be used to determine whether or not new deployments are allowed in the cluster.
func (o *ObjectStorageServiceHandler) ListCluster(ctx context.Context, options *ListOptions) ([]ObjectStorageCluster, *Meta, error) {
	uri := "/v2/object-storage/clusters"
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	clusters := new(objectStorageClustersBase)
	if err = o.client.DoWithContext(ctx, req, clusters); err != nil {
		return nil, nil, err
	}

	return clusters.Clusters, clusters.Meta, nil
}

// RegenerateKeys of the S3 API Keys for an object storage subscription
func (o *ObjectStorageServiceHandler) RegenerateKeys(ctx context.Context, id string) (*S3Keys, error) {
	uri := fmt.Sprintf("/v2/object-storage/%s/regenerate-keys", id)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	s3Keys := new(s3KeysBase)
	if err = o.client.DoWithContext(ctx, req, s3Keys); err != nil {
		return nil, err
	}

	return s3Keys.S3Credentials, nil
}
