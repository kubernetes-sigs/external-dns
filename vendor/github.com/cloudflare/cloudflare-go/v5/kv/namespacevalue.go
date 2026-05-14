// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// NamespaceValueService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceValueService] method instead.
type NamespaceValueService struct {
	Options []option.RequestOption
}

// NewNamespaceValueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceValueService(opts ...option.RequestOption) (r *NamespaceValueService) {
	r = &NamespaceValueService{}
	r.Options = opts
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored. If JSON metadata to be associated with the key/value pair is needed, use
// `multipart/form-data` content type for your PUT request (see dropdown below in
// `REQUEST BODY SCHEMA`). Existing values, expirations, and metadata will be
// overwritten. If neither `expiration` nor `expiration_ttl` is specified, the
// key-value pair will never expire. If both are set, `expiration_ttl` is used and
// `expiration` is ignored.
func (r *NamespaceValueService) Update(ctx context.Context, namespaceID string, keyName string, params NamespaceValueUpdateParams, opts ...option.RequestOption) (res *NamespaceValueUpdateResponse, err error) {
	var env NamespaceValueUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", params.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *NamespaceValueService) Delete(ctx context.Context, namespaceID string, keyName string, body NamespaceValueDeleteParams, opts ...option.RequestOption) (res *NamespaceValueDeleteResponse, err error) {
	var env NamespaceValueDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", body.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the value associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name. If the KV-pair is set to expire at some point, the expiration time as
// measured in seconds since the UNIX epoch will be returned in the `expiration`
// response header.
func (r *NamespaceValueService) Get(ctx context.Context, namespaceID string, keyName string, query NamespaceValueGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", query.AccountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NamespaceValueUpdateResponse struct {
	JSON namespaceValueUpdateResponseJSON `json:"-"`
}

// namespaceValueUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceValueUpdateResponse]
type namespaceValueUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueDeleteResponse struct {
	JSON namespaceValueDeleteResponseJSON `json:"-"`
}

// namespaceValueDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceValueDeleteResponse]
type namespaceValueDeleteResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceValueUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration param.Field[float64] `query:"expiration"`
	// Expires the key after a number of seconds. Must be at least 60.
	ExpirationTTL param.Field[float64] `query:"expiration_ttl"`
	// Associates arbitrary JSON data with a key/value pair.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r NamespaceValueUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [NamespaceValueUpdateParams]'s query parameters as
// `url.Values`.
func (r NamespaceValueUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceValueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceValueUpdateResponse                `json:"result,nullable"`
	JSON    namespaceValueUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceValueUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceValueUpdateResponseEnvelope]
type namespaceValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceValueUpdateResponseEnvelopeSuccess bool

const (
	NamespaceValueUpdateResponseEnvelopeSuccessTrue NamespaceValueUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceValueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceValueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceValueDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceValueDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceValueDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceValueDeleteResponse                `json:"result,nullable"`
	JSON    namespaceValueDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceValueDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceValueDeleteResponseEnvelope]
type namespaceValueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceValueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceValueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceValueDeleteResponseEnvelopeSuccess bool

const (
	NamespaceValueDeleteResponseEnvelopeSuccessTrue NamespaceValueDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceValueDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceValueDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceValueGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
