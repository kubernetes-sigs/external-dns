package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"strconv"

	"errors"
)

// WorkersKVNamespaceRequest provides parameters for creating and updating storage namespaces.
type WorkersKVNamespaceRequest struct {
	Title string `json:"title"`
}

// WorkersKVPair is used in an array in the request to the bulk KV api.
type WorkersKVPair struct {
	Key           string      `json:"key"`
	Value         string      `json:"value"`
	Expiration    int         `json:"expiration,omitempty"`
	ExpirationTTL int         `json:"expiration_ttl,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Base64        bool        `json:"base64,omitempty"`
}

// WorkersKVBulkWriteRequest is the request to the bulk KV api.
type WorkersKVBulkWriteRequest []*WorkersKVPair

// WorkersKVNamespaceResponse is the response received when creating storage namespaces.
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespace contains the unique identifier and title of a storage namespace.
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListWorkersKVNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct.
type ListWorkersKVNamespacesResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value.
type StorageKey struct {
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	Metadata   interface{} `json:"metadata"`
}

// ListWorkersKVsOptions contains optional parameters for listing a namespace's keys.
type ListWorkersKVsOptions struct {
	Limit  *int
	Cursor *string
	Prefix *string
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct.
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, req *WorkersKVNamespaceRequest) (WorkersKVNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return WorkersKVNamespaceResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVNamespaces lists storage namespaces
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context) ([]WorkersKVNamespace, error) {
	v := url.Values{}
	v.Set("per_page", "100")

	var namespaces []WorkersKVNamespace
	page := 1

	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces?%s", api.AccountID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []WorkersKVNamespace{}, errors.Wrap(err, errMakeRequestError)
		}

		var p ListWorkersKVNamespacesResponse
		if err := json.Unmarshal(res, &p); err != nil {
			return []WorkersKVNamespace{}, errors.Wrap(err, errUnmarshalError)
		}

		if !p.Success {
			return []WorkersKVNamespace{}, errors.New(errRequestNotSuccessful)
		}

		namespaces = append(namespaces, p.Result...)
		if p.ResultInfo.Page >= p.ResultInfo.TotalPages {
			break
		}

		page++
	}

	return namespaces, nil
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a namespace's title
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, namespaceID string, req *WorkersKVNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKV writes a value identified by a key.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, value, http.Header{"Content-Type": []string{"application/octet-stream"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKVBulk writes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-multiple-key-value-pairs
func (api *API) WriteWorkersKVBulk(ctx context.Context, namespaceID string, kvs WorkersKVBulkWriteRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, kvs, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ReadWorkersKV returns the value associated with the given key in the given namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api API) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	return res, nil
}

// DeleteWorkersKV deletes a key and value for a provided storage namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api API) DeleteWorkersKV(ctx context.Context, namespaceID, key string) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// DeleteWorkersKVBulk deletes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-multiple-key-value-pairs
func (api *API) DeleteWorkersKVBulk(ctx context.Context, namespaceID string, keys []string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodDelete, uri, keys, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVs lists a namespace's keys
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVs(ctx context.Context, namespaceID string) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// encode encodes non-nil fields into URL encoded form.
func (o ListWorkersKVsOptions) encode() string {
	v := url.Values{}
	if o.Limit != nil {
		v.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Cursor != nil {
		v.Set("cursor", *o.Cursor)
	}
	if o.Prefix != nil {
		v.Set("prefix", *o.Prefix)
	}
	return v.Encode()
}

// ListWorkersKVsWithOptions lists a namespace's keys with optional parameters
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVsWithOptions(ctx context.Context, namespaceID string, o ListWorkersKVsOptions) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys?%s", api.AccountID, namespaceID, o.encode())
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"strconv"
>>>>>>> 5ce8c7613 (update vendored files)

	"github.com/pkg/errors"
)

// WorkersKVNamespaceRequest provides parameters for creating and updating storage namespaces
type WorkersKVNamespaceRequest struct {
	Title string `json:"title"`
}

// WorkersKVPair is used in an array in the request to the bulk KV api
type WorkersKVPair struct {
	Key           string      `json:"key"`
	Value         string      `json:"value"`
	Expiration    int         `json:"expiration,omitempty"`
	ExpirationTTL int         `json:"expiration_ttl,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Base64        bool        `json:"base64,omitempty"`
}

// WorkersKVBulkWriteRequest is the request to the bulk KV api
type WorkersKVBulkWriteRequest []*WorkersKVPair

// WorkersKVNamespaceResponse is the response received when creating storage namespaces
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespace contains the unique identifier and title of a storage namespace
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListWorkersKVNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct
type ListWorkersKVNamespacesResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value
type StorageKey struct {
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	Metadata   interface{} `json:"metadata"`
}

// ListWorkersKVsOptions contains optional parameters for listing a namespace's keys
type ListWorkersKVsOptions struct {
	Limit  *int
	Cursor *string
	Prefix *string
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, req *WorkersKVNamespaceRequest) (WorkersKVNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return WorkersKVNamespaceResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVNamespaces lists storage namespaces
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context) ([]WorkersKVNamespace, error) {
	v := url.Values{}
	v.Set("per_page", "100")

	var namespaces []WorkersKVNamespace
	page := 1

	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces?%s", api.AccountID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []WorkersKVNamespace{}, errors.Wrap(err, errMakeRequestError)
		}

		var p ListWorkersKVNamespacesResponse
		if err := json.Unmarshal(res, &p); err != nil {
			return []WorkersKVNamespace{}, errors.Wrap(err, errUnmarshalError)
		}

		if !p.Success {
			return []WorkersKVNamespace{}, errors.New(errRequestNotSuccessful)
		}

		namespaces = append(namespaces, p.Result...)
		if p.ResultInfo.Page >= p.ResultInfo.TotalPages {
			break
		}

		page++
	}

	return namespaces, nil
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a namespace's title
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, namespaceID string, req *WorkersKVNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKV writes a value identified by a key.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, value, http.Header{"Content-Type": []string{"application/octet-stream"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKVBulk writes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-multiple-key-value-pairs
func (api *API) WriteWorkersKVBulk(ctx context.Context, namespaceID string, kvs WorkersKVBulkWriteRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, kvs, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ReadWorkersKV returns the value associated with the given key in the given namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api API) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	return res, nil
}

// DeleteWorkersKV deletes a key and value for a provided storage namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api API) DeleteWorkersKV(ctx context.Context, namespaceID, key string) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// DeleteWorkersKVBulk deletes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-multiple-key-value-pairs
func (api *API) DeleteWorkersKVBulk(ctx context.Context, namespaceID string, keys []string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodDelete, uri, keys, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVs lists a namespace's keys
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVs(ctx context.Context, namespaceID string) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", api.AccountID, namespaceID)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// encode encodes non-nil fields into URL encoded form.
func (o ListWorkersKVsOptions) encode() string {
	v := url.Values{}
	if o.Limit != nil {
		v.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Cursor != nil {
		v.Set("cursor", *o.Cursor)
	}
	if o.Prefix != nil {
		v.Set("prefix", *o.Prefix)
	}
	return v.Encode()
}

// ListWorkersKVsWithOptions lists a namespace's keys with optional parameters
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVsWithOptions(ctx context.Context, namespaceID string, o ListWorkersKVsOptions) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys?%s", api.AccountID, namespaceID, o.encode())
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"strconv"
>>>>>>> 6b7ce455e (update vendored files)

	"github.com/pkg/errors"
)

// WorkersKVNamespaceRequest provides parameters for creating and updating storage namespaces
type WorkersKVNamespaceRequest struct {
	Title string `json:"title"`
}

// WorkersKVPair is used in an array in the request to the bulk KV api
type WorkersKVPair struct {
	Key           string      `json:"key"`
	Value         string      `json:"value"`
	Expiration    int         `json:"expiration,omitempty"`
	ExpirationTTL int         `json:"expiration_ttl,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Base64        bool        `json:"base64,omitempty"`
}

// WorkersKVBulkWriteRequest is the request to the bulk KV api
type WorkersKVBulkWriteRequest []*WorkersKVPair

// WorkersKVNamespaceResponse is the response received when creating storage namespaces
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespace contains the unique identifier and title of a storage namespace
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListWorkersKVNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct
type ListWorkersKVNamespacesResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value
type StorageKey struct {
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	Metadata   interface{} `json:"metadata"`
}

// ListWorkersKVsOptions contains optional parameters for listing a namespace's keys
type ListWorkersKVsOptions struct {
	Limit  *int
	Cursor *string
	Prefix *string
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, req *WorkersKVNamespaceRequest) (WorkersKVNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return WorkersKVNamespaceResponse{}, err
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVNamespaces lists storage namespaces
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context) ([]WorkersKVNamespace, error) {
	v := url.Values{}
	v.Set("per_page", "100")

	var namespaces []WorkersKVNamespace
	page := 1

	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces?%s", api.AccountID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []WorkersKVNamespace{}, err
		}

		var p ListWorkersKVNamespacesResponse
		if err := json.Unmarshal(res, &p); err != nil {
			return []WorkersKVNamespace{}, errors.Wrap(err, errUnmarshalError)
		}

		if !p.Success {
			return []WorkersKVNamespace{}, errors.New(errRequestNotSuccessful)
		}

		namespaces = append(namespaces, p.Result...)
		if p.ResultInfo.Page >= p.ResultInfo.TotalPages {
			break
		}

		page++
	}

	return namespaces, nil
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a namespace's title
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, namespaceID string, req *WorkersKVNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKV writes a value identified by a key.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, value, http.Header{"Content-Type": []string{"application/octet-stream"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKVBulk writes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-multiple-key-value-pairs
func (api *API) WriteWorkersKVBulk(ctx context.Context, namespaceID string, kvs WorkersKVBulkWriteRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodPut, uri, kvs, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ReadWorkersKV returns the value associated with the given key in the given namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api API) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteWorkersKV deletes a key and value for a provided storage namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api API) DeleteWorkersKV(ctx context.Context, namespaceID, key string) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// DeleteWorkersKVBulk deletes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-multiple-key-value-pairs
func (api *API) DeleteWorkersKVBulk(ctx context.Context, namespaceID string, keys []string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestWithHeaders(
		http.MethodDelete, uri, keys, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVs lists a namespace's keys
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVs(ctx context.Context, namespaceID string) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", api.AccountID, namespaceID)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, err
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// encode encodes non-nil fields into URL encoded form.
func (o ListWorkersKVsOptions) encode() string {
	v := url.Values{}
	if o.Limit != nil {
		v.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Cursor != nil {
		v.Set("cursor", *o.Cursor)
	}
	if o.Prefix != nil {
		v.Set("prefix", *o.Prefix)
	}
	return v.Encode()
}

// ListWorkersKVsWithOptions lists a namespace's keys with optional parameters
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVsWithOptions(ctx context.Context, namespaceID string, o ListWorkersKVsOptions) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys?%s", api.AccountID, namespaceID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"strconv"
>>>>>>> 4d7e5ad26 (update vendored files)

	"github.com/pkg/errors"
)

// WorkersKVNamespaceRequest provides parameters for creating and updating storage namespaces
type WorkersKVNamespaceRequest struct {
	Title string `json:"title"`
}

// WorkersKVPair is used in an array in the request to the bulk KV api
type WorkersKVPair struct {
	Key           string      `json:"key"`
	Value         string      `json:"value"`
	Expiration    int         `json:"expiration,omitempty"`
	ExpirationTTL int         `json:"expiration_ttl,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Base64        bool        `json:"base64,omitempty"`
}

// WorkersKVBulkWriteRequest is the request to the bulk KV api
type WorkersKVBulkWriteRequest []*WorkersKVPair

// WorkersKVNamespaceResponse is the response received when creating storage namespaces
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespace contains the unique identifier and title of a storage namespace
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListWorkersKVNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct
type ListWorkersKVNamespacesResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value
type StorageKey struct {
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	Metadata   interface{} `json:"metadata"`
}

// ListWorkersKVsOptions contains optional parameters for listing a namespace's keys
type ListWorkersKVsOptions struct {
	Limit  *int
	Cursor *string
	Prefix *string
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, req *WorkersKVNamespaceRequest) (WorkersKVNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return WorkersKVNamespaceResponse{}, err
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersKVNamespaces lists storage namespaces
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context) ([]WorkersKVNamespace, error) {
	v := url.Values{}
	v.Set("per_page", "100")

	var namespaces []WorkersKVNamespace
	page := 1

	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces?%s", api.AccountID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []WorkersKVNamespace{}, err
		}

		var p ListWorkersKVNamespacesResponse
		if err := json.Unmarshal(res, &p); err != nil {
			return []WorkersKVNamespace{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}

		if !p.Success {
			return []WorkersKVNamespace{}, errors.New(errRequestNotSuccessful)
		}

		namespaces = append(namespaces, p.Result...)
		if p.ResultInfo.Page >= p.ResultInfo.TotalPages {
			break
		}

		page++
	}

	return namespaces, nil
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a namespace's title
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, namespaceID string, req *WorkersKVNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// WriteWorkersKV writes a value identified by a key.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodPut, uri, value, http.Header{"Content-Type": []string{"application/octet-stream"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// WriteWorkersKVBulk writes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-multiple-key-value-pairs
func (api *API) WriteWorkersKVBulk(ctx context.Context, namespaceID string, kvs WorkersKVBulkWriteRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodPut, uri, kvs, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ReadWorkersKV returns the value associated with the given key in the given namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api API) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteWorkersKV deletes a key and value for a provided storage namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api API) DeleteWorkersKV(ctx context.Context, namespaceID, key string) (Response, error) {
	key = url.PathEscape(key)
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.AccountID, namespaceID, key)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, err
}

// DeleteWorkersKVBulk deletes multiple KVs at once.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-multiple-key-value-pairs
func (api *API) DeleteWorkersKVBulk(ctx context.Context, namespaceID string, keys []string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.AccountID, namespaceID)
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodDelete, uri, keys, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersKVs lists a namespace's keys
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVs(ctx context.Context, namespaceID string) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", api.AccountID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return ListStorageKeysResponse{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return ListStorageKeysResponse{}, errors.Wrap(err, errMakeRequestError)
=======
		return ListStorageKeysResponse{}, err
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, err
}

// encode encodes non-nil fields into URL encoded form.
func (o ListWorkersKVsOptions) encode() string {
	v := url.Values{}
	if o.Limit != nil {
		v.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Cursor != nil {
		v.Set("cursor", *o.Cursor)
	}
	if o.Prefix != nil {
		v.Set("prefix", *o.Prefix)
	}
	return v.Encode()
}

// ListWorkersKVsWithOptions lists a namespace's keys with optional parameters
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListWorkersKVsWithOptions(ctx context.Context, namespaceID string, o ListWorkersKVsOptions) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys?%s", api.AccountID, namespaceID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

	"github.com/goccy/go-json"
)

// CreateWorkersKVNamespaceParams provides parameters for creating and updating storage namespaces.
type CreateWorkersKVNamespaceParams struct {
	Title string `json:"title"`
}

type UpdateWorkersKVNamespaceParams struct {
	NamespaceID string `json:"-"`
	Title       string `json:"title"`
}

// WorkersKVPair is used in an array in the request to the bulk KV api.
type WorkersKVPair struct {
	Key           string      `json:"key"`
	Value         string      `json:"value"`
	Expiration    int         `json:"expiration,omitempty"`
	ExpirationTTL int         `json:"expiration_ttl,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Base64        bool        `json:"base64,omitempty"`
}

// WorkersKVNamespaceResponse is the response received when creating storage namespaces.
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespace contains the unique identifier and title of a storage namespace.
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListWorkersKVNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct.
type ListWorkersKVNamespacesResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value.
type StorageKey struct {
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	Metadata   interface{} `json:"metadata"`
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct.
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

type ListWorkersKVNamespacesParams struct {
	ResultInfo
}

type WriteWorkersKVEntryParams struct {
	NamespaceID string
	Key         string
	Value       []byte
}

type WriteWorkersKVEntriesParams struct {
	NamespaceID string
	KVs         []*WorkersKVPair
}

type GetWorkersKVParams struct {
	NamespaceID string
	Key         string
}

type DeleteWorkersKVEntryParams struct {
	NamespaceID string
	Key         string
}

type DeleteWorkersKVEntriesParams struct {
	NamespaceID string
	Keys        []string
}

type ListWorkersKVsParams struct {
	NamespaceID string `url:"-"`
	Limit       int    `url:"limit,omitempty"`
	Cursor      string `url:"cursor,omitempty"`
	Prefix      string `url:"prefix,omitempty"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, rc *ResourceContainer, params CreateWorkersKVNamespaceParams) (WorkersKVNamespaceResponse, error) {
	if rc.Level != AccountRouteLevel {
		return WorkersKVNamespaceResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkersKVNamespaceResponse{}, ErrMissingIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return WorkersKVNamespaceResponse{}, err
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersKVNamespaces lists storage namespaces.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context, rc *ResourceContainer, params ListWorkersKVNamespacesParams) ([]WorkersKVNamespace, *ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []WorkersKVNamespace{}, &ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []WorkersKVNamespace{}, &ResultInfo{}, ErrMissingIdentifier
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = 50
	}
	if params.Page < 1 {
		params.Page = 1
	}

	var namespaces []WorkersKVNamespace
	var nsResponse ListWorkersKVNamespacesResponse
	for {
		nsResponse = ListWorkersKVNamespacesResponse{}
		uri := buildURI(fmt.Sprintf("/accounts/%s/storage/kv/namespaces", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []WorkersKVNamespace{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &nsResponse)
		if err != nil {
			return []WorkersKVNamespace{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal workers KV namespaces JSON data: %w", err)
		}

		namespaces = append(namespaces, nsResponse.Result...)
		params.ResultInfo = nsResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return namespaces, &nsResponse.ResultInfo, nil
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, rc *ResourceContainer, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", rc.Identifier, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a KV namespace based on the ID.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, rc *ResourceContainer, params UpdateWorkersKVNamespaceParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", rc.Identifier, params.NamespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// WriteWorkersKVEntry writes a single KV value based on the key.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-write-key-value-pair-with-metadata
func (api *API) WriteWorkersKVEntry(ctx context.Context, rc *ResourceContainer, params WriteWorkersKVEntryParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", rc.Identifier, params.NamespaceID, url.PathEscape(params.Key))
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodPut, uri, params.Value, http.Header{"Content-Type": []string{"application/octet-stream"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// WriteWorkersKVEntries writes multiple KVs at once.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-write-multiple-key-value-pairs
func (api *API) WriteWorkersKVEntries(ctx context.Context, rc *ResourceContainer, params WriteWorkersKVEntriesParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", rc.Identifier, params.NamespaceID)
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodPut, uri, params.KVs, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// GetWorkersKV returns the value associated with the given key in the
// given namespace.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-read-key-value-pair
func (api API) GetWorkersKV(ctx context.Context, rc *ResourceContainer, params GetWorkersKVParams) ([]byte, error) {
	if rc.Level != AccountRouteLevel {
		return []byte(``), ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []byte(``), ErrMissingIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", rc.Identifier, params.NamespaceID, url.PathEscape(params.Key))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteWorkersKVEntry deletes a key and value for a provided storage namespace.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-delete-key-value-pair
func (api API) DeleteWorkersKVEntry(ctx context.Context, rc *ResourceContainer, params DeleteWorkersKVEntryParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", rc.Identifier, params.NamespaceID, url.PathEscape(params.Key))
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, err
}

// DeleteWorkersKVEntries deletes multiple KVs at once.
//
// API reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-delete-multiple-key-value-pairs
func (api *API) DeleteWorkersKVEntries(ctx context.Context, rc *ResourceContainer, params DeleteWorkersKVEntriesParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", rc.Identifier, params.NamespaceID)
	res, err := api.makeRequestContextWithHeaders(
		ctx, http.MethodDelete, uri, params.Keys, http.Header{"Content-Type": []string{"application/json"}},
	)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersKVKeys lists a namespace's keys.
//
// API Reference: https://developers.cloudflare.com/api/operations/workers-kv-namespace-list-a-namespace'-s-keys
func (api API) ListWorkersKVKeys(ctx context.Context, rc *ResourceContainer, params ListWorkersKVsParams) (ListStorageKeysResponse, error) {
	if rc.Level != AccountRouteLevel {
		return ListStorageKeysResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ListStorageKeysResponse{}, ErrMissingIdentifier
	}

	uri := buildURI(
		fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", rc.Identifier, params.NamespaceID),
		params,
	)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, err
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
<<<<<<< HEAD
		return result, errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return result, errors.Wrap(err, errUnmarshalError)
=======
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	return result, err
}
