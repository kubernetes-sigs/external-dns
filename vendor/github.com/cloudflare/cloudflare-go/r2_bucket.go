package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
<<<<<<< HEAD
)

var (
	ErrMissingBucketName = errors.New("require bucket name missing")
)

type CreateR2BucketParameters struct {
	Name string `json:"name,omitempty"`
}

// CreateR2Bucket Creates a new R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-create-bucket
func (api *API) CreateR2Bucket(ctx context.Context, rc *ResourceContainer, params CreateR2BucketParameters) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if params.Name == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets", rc.Identifier)
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)

	return err
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingBucketName = errors.New("require bucket name missing")
)

// R2Bucket defines a container for objects stored in R2 Storage.
type R2Bucket struct {
	Name         string     `json:"name"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	Location     string     `json:"location,omitempty"`
}

// R2Buckets represents the map of buckets response from
// the R2 buckets endpoint.
type R2Buckets struct {
	Buckets []R2Bucket `json:"buckets"`
}

// R2BucketListResponse represents the response from the list
// R2 buckets endpoint.
type R2BucketListResponse struct {
	Result R2Buckets `json:"result"`
	Response
}

type ListR2BucketsParams struct {
	Name       string `url:"name_contains,omitempty"`
	StartAfter string `url:"start_after,omitempty"`
	PerPage    int64  `url:"per_page,omitempty"`
	Order      string `url:"order,omitempty"`
	Direction  string `url:"direction,omitempty"`
	Cursor     string `url:"cursor,omitempty"`
}

type CreateR2BucketParameters struct {
	Name         string `json:"name,omitempty"`
	LocationHint string `json:"locationHint,omitempty"`
}

type R2BucketResponse struct {
	Result R2Bucket `json:"result"`
	Response
}

// ListR2Buckets Lists R2 buckets.
func (api *API) ListR2Buckets(ctx context.Context, rc *ResourceContainer, params ListR2BucketsParams) ([]R2Bucket, error) {
	if rc.Identifier == "" {
		return []R2Bucket{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/r2/buckets", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []R2Bucket{}, err
	}

	var r2BucketListResponse R2BucketListResponse
	err = json.Unmarshal(res, &r2BucketListResponse)
	if err != nil {
		return []R2Bucket{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r2BucketListResponse.Result.Buckets, nil
}

// CreateR2Bucket Creates a new R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-create-bucket
func (api *API) CreateR2Bucket(ctx context.Context, rc *ResourceContainer, params CreateR2BucketParameters) (R2Bucket, error) {
	if rc.Identifier == "" {
		return R2Bucket{}, ErrMissingAccountID
	}

	if params.Name == "" {
		return R2Bucket{}, ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return R2Bucket{}, err
	}

	var r2BucketResponse R2BucketResponse
	err = json.Unmarshal(res, &r2BucketResponse)
	if err != nil {
		return R2Bucket{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r2BucketResponse.Result, nil
}

// GetR2Bucket Gets an existing R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-get-bucket
func (api *API) GetR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) (R2Bucket, error) {
	if rc.Identifier == "" {
		return R2Bucket{}, ErrMissingAccountID
	}

	if bucketName == "" {
		return R2Bucket{}, ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return R2Bucket{}, err
	}

	var r2BucketResponse R2BucketResponse
	err = json.Unmarshal(res, &r2BucketResponse)
	if err != nil {
		return R2Bucket{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r2BucketResponse.Result, nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

// DeleteR2Bucket Deletes an existing R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-delete-bucket
func (api *API) DeleteR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if bucketName == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	return err
}
