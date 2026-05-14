// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// NamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceKeyService] method instead.
type NamespaceKeyService struct {
	Options []option.RequestOption
}

// NewNamespaceKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceKeyService(opts ...option.RequestOption) (r *NamespaceKeyService) {
	r = &NamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) List(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) (res *pagination.CursorLimitPagination[Key], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", params.AccountID, namespaceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) ListAutoPaging(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) *pagination.CursorLimitPaginationAutoPager[Key] {
	return pagination.NewCursorLimitPaginationAutoPager(r.List(ctx, namespaceID, params, opts...))
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
//
// Deprecated: Please use kv.namespaces.bulk_delete instead
func (r *NamespaceKeyService) BulkDelete(ctx context.Context, namespaceID string, params NamespaceKeyBulkDeleteParams, opts ...option.RequestOption) (res *NamespaceKeyBulkDeleteResponse, err error) {
	var env NamespaceKeyBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk/delete", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based
// values. JSON values can optionally be parsed instead of being returned as a
// string value. Metadata can be included if `withMetadata` is true.
//
// Deprecated: Please use kv.namespaces.bulk_get instead
func (r *NamespaceKeyService) BulkGet(ctx context.Context, namespaceID string, params NamespaceKeyBulkGetParams, opts ...option.RequestOption) (res *NamespaceKeyBulkGetResponse, err error) {
	var env NamespaceKeyBulkGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk/get", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
//
// Deprecated: Please use kv.namespaces.bulk_update instead
func (r *NamespaceKeyService) BulkUpdate(ctx context.Context, namespaceID string, params NamespaceKeyBulkUpdateParams, opts ...option.RequestOption) (res *NamespaceKeyBulkUpdateResponse, err error) {
	var env NamespaceKeyBulkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type Key struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata interface{} `json:"metadata"`
	JSON     keyJSON     `json:"-"`
}

// keyJSON contains the JSON metadata for the struct [Key]
type keyJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Key) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyBulkDeleteResponse struct {
	// Number of keys successfully updated.
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                           `json:"unsuccessful_keys"`
	JSON             namespaceKeyBulkDeleteResponseJSON `json:"-"`
}

// namespaceKeyBulkDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceKeyBulkDeleteResponse]
type namespaceKeyBulkDeleteResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceKeyBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyBulkGetResponse struct {
	// This field can have the runtime type of
	// [map[string]NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion],
	// [map[string]NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValue].
	Values interface{}                     `json:"values"`
	JSON   namespaceKeyBulkGetResponseJSON `json:"-"`
	union  NamespaceKeyBulkGetResponseUnion
}

// namespaceKeyBulkGetResponseJSON contains the JSON metadata for the struct
// [NamespaceKeyBulkGetResponse]
type namespaceKeyBulkGetResponseJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r namespaceKeyBulkGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *NamespaceKeyBulkGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = NamespaceKeyBulkGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [NamespaceKeyBulkGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [NamespaceKeyBulkGetResponseWorkersKVBulkGetResult],
// [NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata].
func (r NamespaceKeyBulkGetResponse) AsUnion() NamespaceKeyBulkGetResponseUnion {
	return r.union
}

// Union satisfied by [NamespaceKeyBulkGetResponseWorkersKVBulkGetResult] or
// [NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata].
type NamespaceKeyBulkGetResponseUnion interface {
	implementsNamespaceKeyBulkGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceKeyBulkGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceKeyBulkGetResponseWorkersKVBulkGetResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata{}),
		},
	)
}

type NamespaceKeyBulkGetResponseWorkersKVBulkGetResult struct {
	// Requested keys are paired with their values in an object.
	Values map[string]NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion `json:"values"`
	JSON   namespaceKeyBulkGetResponseWorkersKVBulkGetResultJSON                   `json:"-"`
}

// namespaceKeyBulkGetResponseWorkersKVBulkGetResultJSON contains the JSON metadata
// for the struct [NamespaceKeyBulkGetResponseWorkersKVBulkGetResult]
type namespaceKeyBulkGetResponseWorkersKVBulkGetResultJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkGetResponseWorkersKVBulkGetResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkGetResponseWorkersKVBulkGetResultJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceKeyBulkGetResponseWorkersKVBulkGetResult) implementsNamespaceKeyBulkGetResponse() {}

// The value associated with the key.
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesMap].
type NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion interface {
	ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesMap{}),
		},
	)
}

type NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesMap map[string]interface{}

func (r NamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesMap) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion() {
}

type NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata struct {
	// Requested keys are paired with their values and metadata in an object.
	Values map[string]NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValue `json:"values"`
	JSON   namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON             `json:"-"`
}

// namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON contains the
// JSON metadata for the struct
// [NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata]
type namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadata) implementsNamespaceKeyBulkGetResponse() {
}

type NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValue struct {
	// The metadata associated with the key.
	Metadata interface{} `json:"metadata,required"`
	// The value associated with the key.
	Value interface{} `json:"value,required"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration float64                                                                `json:"expiration"`
	JSON       namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON `json:"-"`
}

// namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON contains
// the JSON metadata for the struct
// [NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValue]
type namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON struct {
	Metadata    apijson.Field
	Value       apijson.Field
	Expiration  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyBulkUpdateResponse struct {
	// Number of keys successfully updated.
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                           `json:"unsuccessful_keys"`
	JSON             namespaceKeyBulkUpdateResponseJSON `json:"-"`
}

// namespaceKeyBulkUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceKeyBulkUpdateResponse]
type namespaceKeyBulkUpdateResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceKeyBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the `cursors`
	// object in the `result_info` structure.
	Cursor param.Field[string] `query:"cursor"`
	// Limits the number of keys returned in the response. The cursor attribute may be
	// used to iterate over the next batch of keys if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// Filters returned keys by a name prefix. Exact matches and any key names that
	// begin with the prefix will be returned.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [NamespaceKeyListParams]'s query parameters as `url.Values`.
func (r NamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceKeyBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      []string            `json:"body,required"`
}

func (r NamespaceKeyBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceKeyBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceKeyBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceKeyBulkDeleteResponse                `json:"result,nullable"`
	JSON    namespaceKeyBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceKeyBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceKeyBulkDeleteResponseEnvelope]
type namespaceKeyBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceKeyBulkDeleteResponseEnvelopeSuccess bool

const (
	NamespaceKeyBulkDeleteResponseEnvelopeSuccessTrue NamespaceKeyBulkDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceKeyBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceKeyBulkGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of keys to retrieve (maximum of 100).
	Keys param.Field[[]string] `json:"keys,required"`
	// Whether to parse JSON values in the response.
	Type param.Field[NamespaceKeyBulkGetParamsType] `json:"type"`
	// Whether to include metadata in the response.
	WithMetadata param.Field[bool] `json:"withMetadata"`
}

func (r NamespaceKeyBulkGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to parse JSON values in the response.
type NamespaceKeyBulkGetParamsType string

const (
	NamespaceKeyBulkGetParamsTypeText NamespaceKeyBulkGetParamsType = "text"
	NamespaceKeyBulkGetParamsTypeJson NamespaceKeyBulkGetParamsType = "json"
)

func (r NamespaceKeyBulkGetParamsType) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkGetParamsTypeText, NamespaceKeyBulkGetParamsTypeJson:
		return true
	}
	return false
}

type NamespaceKeyBulkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceKeyBulkGetResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceKeyBulkGetResponse                `json:"result,nullable"`
	JSON    namespaceKeyBulkGetResponseEnvelopeJSON    `json:"-"`
}

// namespaceKeyBulkGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceKeyBulkGetResponseEnvelope]
type namespaceKeyBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceKeyBulkGetResponseEnvelopeSuccess bool

const (
	NamespaceKeyBulkGetResponseEnvelopeSuccessTrue NamespaceKeyBulkGetResponseEnvelopeSuccess = true
)

func (r NamespaceKeyBulkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceKeyBulkUpdateParams struct {
	// Identifier.
	AccountID param.Field[string]                `path:"account_id,required"`
	Body      []NamespaceKeyBulkUpdateParamsBody `json:"body,required"`
}

func (r NamespaceKeyBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceKeyBulkUpdateParamsBody struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key,required"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
	// Indicates whether or not the server should base64 decode the value before
	// storing it. Useful for writing values that wouldn't otherwise be valid JSON
	// strings, such as images.
	Base64 param.Field[bool] `json:"base64"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration param.Field[float64] `json:"expiration"`
	// Expires the key after a number of seconds. Must be at least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r NamespaceKeyBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceKeyBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceKeyBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceKeyBulkUpdateResponse                `json:"result,nullable"`
	JSON    namespaceKeyBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceKeyBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceKeyBulkUpdateResponseEnvelope]
type namespaceKeyBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceKeyBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceKeyBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceKeyBulkUpdateResponseEnvelopeSuccess bool

const (
	NamespaceKeyBulkUpdateResponseEnvelopeSuccessTrue NamespaceKeyBulkUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceKeyBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceKeyBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
