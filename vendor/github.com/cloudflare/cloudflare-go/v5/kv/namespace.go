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

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options  []option.RequestOption
	Keys     *NamespaceKeyService
	Metadata *NamespaceMetadataService
	Values   *NamespaceValueService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Keys = NewNamespaceKeyService(opts...)
	r.Metadata = NewNamespaceMetadataService(opts...)
	r.Values = NewNamespaceValueService(opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *NamespaceService) New(ctx context.Context, params NamespaceNewParams, opts ...option.RequestOption) (res *Namespace, err error) {
	var env NamespaceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modifies a namespace's title.
func (r *NamespaceService) Update(ctx context.Context, namespaceID string, params NamespaceUpdateParams, opts ...option.RequestOption) (res *Namespace, err error) {
	var env NamespaceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the namespaces owned by an account.
func (r *NamespaceService) List(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Namespace], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
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

// Returns the namespaces owned by an account.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Namespace] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the namespace corresponding to the given ID.
func (r *NamespaceService) Delete(ctx context.Context, namespaceID string, body NamespaceDeleteParams, opts ...option.RequestOption) (res *NamespaceDeleteResponse, err error) {
	var env NamespaceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", body.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *NamespaceService) BulkDelete(ctx context.Context, namespaceID string, params NamespaceBulkDeleteParams, opts ...option.RequestOption) (res *NamespaceBulkDeleteResponse, err error) {
	var env NamespaceBulkDeleteResponseEnvelope
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
func (r *NamespaceService) BulkGet(ctx context.Context, namespaceID string, params NamespaceBulkGetParams, opts ...option.RequestOption) (res *NamespaceBulkGetResponse, err error) {
	var env NamespaceBulkGetResponseEnvelope
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
func (r *NamespaceService) BulkUpdate(ctx context.Context, namespaceID string, params NamespaceBulkUpdateParams, opts ...option.RequestOption) (res *NamespaceBulkUpdateResponse, err error) {
	var env NamespaceBulkUpdateResponseEnvelope
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

// Get the namespace corresponding to the given ID.
func (r *NamespaceService) Get(ctx context.Context, namespaceID string, query NamespaceGetParams, opts ...option.RequestOption) (res *Namespace, err error) {
	var env NamespaceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", query.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Namespace struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if new beta namespace, with additional preview features.
	Beta bool `json:"beta"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool          `json:"supports_url_encoding"`
	JSON                namespaceJSON `json:"-"`
}

// namespaceJSON contains the JSON metadata for the struct [Namespace]
type namespaceJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	Beta                apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Namespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceJSON) RawJSON() string {
	return r.raw
}

type NamespaceDeleteResponse struct {
	JSON namespaceDeleteResponseJSON `json:"-"`
}

// namespaceDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceDeleteResponse]
type namespaceDeleteResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkDeleteResponse struct {
	// Number of keys successfully updated.
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                        `json:"unsuccessful_keys"`
	JSON             namespaceBulkDeleteResponseJSON `json:"-"`
}

// namespaceBulkDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceBulkDeleteResponse]
type namespaceBulkDeleteResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkGetResponse struct {
	// This field can have the runtime type of
	// [map[string]NamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion],
	// [map[string]NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValue].
	Values interface{}                  `json:"values"`
	JSON   namespaceBulkGetResponseJSON `json:"-"`
	union  NamespaceBulkGetResponseUnion
}

// namespaceBulkGetResponseJSON contains the JSON metadata for the struct
// [NamespaceBulkGetResponse]
type namespaceBulkGetResponseJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r namespaceBulkGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *NamespaceBulkGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = NamespaceBulkGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [NamespaceBulkGetResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [NamespaceBulkGetResponseWorkersKVBulkGetResult],
// [NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata].
func (r NamespaceBulkGetResponse) AsUnion() NamespaceBulkGetResponseUnion {
	return r.union
}

// Union satisfied by [NamespaceBulkGetResponseWorkersKVBulkGetResult] or
// [NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata].
type NamespaceBulkGetResponseUnion interface {
	implementsNamespaceBulkGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceBulkGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceBulkGetResponseWorkersKVBulkGetResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata{}),
		},
	)
}

type NamespaceBulkGetResponseWorkersKVBulkGetResult struct {
	// Requested keys are paired with their values in an object.
	Values map[string]NamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion `json:"values"`
	JSON   namespaceBulkGetResponseWorkersKVBulkGetResultJSON                   `json:"-"`
}

// namespaceBulkGetResponseWorkersKVBulkGetResultJSON contains the JSON metadata
// for the struct [NamespaceBulkGetResponseWorkersKVBulkGetResult]
type namespaceBulkGetResponseWorkersKVBulkGetResultJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkGetResponseWorkersKVBulkGetResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkGetResponseWorkersKVBulkGetResultJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceBulkGetResponseWorkersKVBulkGetResult) implementsNamespaceBulkGetResponse() {}

// The value associated with the key.
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [NamespaceBulkGetResponseWorkersKVBulkGetResultValuesMap].
type NamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion interface {
	ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(NamespaceBulkGetResponseWorkersKVBulkGetResultValuesMap{}),
		},
	)
}

type NamespaceBulkGetResponseWorkersKVBulkGetResultValuesMap map[string]interface{}

func (r NamespaceBulkGetResponseWorkersKVBulkGetResultValuesMap) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion() {
}

type NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata struct {
	// Requested keys are paired with their values and metadata in an object.
	Values map[string]NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValue `json:"values"`
	JSON   namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON             `json:"-"`
}

// namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON contains the JSON
// metadata for the struct
// [NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata]
type namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadata) implementsNamespaceBulkGetResponse() {
}

type NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValue struct {
	// The metadata associated with the key.
	Metadata interface{} `json:"metadata,required"`
	// The value associated with the key.
	Value interface{} `json:"value,required"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration float64                                                             `json:"expiration"`
	JSON       namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON `json:"-"`
}

// namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON contains the
// JSON metadata for the struct
// [NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValue]
type namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON struct {
	Metadata    apijson.Field
	Value       apijson.Field
	Expiration  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkGetResponseWorkersKVBulkGetResultWithMetadataValueJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkUpdateResponse struct {
	// Number of keys successfully updated.
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string                        `json:"unsuccessful_keys"`
	JSON             namespaceBulkUpdateResponseJSON `json:"-"`
}

// namespaceBulkUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceBulkUpdateResponse]
type namespaceBulkUpdateResponseJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r NamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Namespace                           `json:"result"`
	JSON    namespaceNewResponseEnvelopeJSON    `json:"-"`
}

// namespaceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceNewResponseEnvelope]
type namespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceNewResponseEnvelopeSuccess bool

const (
	NamespaceNewResponseEnvelopeSuccessTrue NamespaceNewResponseEnvelopeSuccess = true
)

func (r NamespaceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r NamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Namespace             `json:"result,required"`
	// Whether the API call was successful.
	Success NamespaceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceUpdateResponseEnvelope]
type namespaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceUpdateResponseEnvelopeSuccess bool

const (
	NamespaceUpdateResponseEnvelopeSuccessTrue NamespaceUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order namespaces.
	Direction param.Field[NamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[NamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order namespaces.
type NamespaceListParamsDirection string

const (
	NamespaceListParamsDirectionAsc  NamespaceListParamsDirection = "asc"
	NamespaceListParamsDirectionDesc NamespaceListParamsDirection = "desc"
)

func (r NamespaceListParamsDirection) IsKnown() bool {
	switch r {
	case NamespaceListParamsDirectionAsc, NamespaceListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type NamespaceListParamsOrder string

const (
	NamespaceListParamsOrderID    NamespaceListParamsOrder = "id"
	NamespaceListParamsOrderTitle NamespaceListParamsOrder = "title"
)

func (r NamespaceListParamsOrder) IsKnown() bool {
	switch r {
	case NamespaceListParamsOrderID, NamespaceListParamsOrderTitle:
		return true
	}
	return false
}

type NamespaceDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceDeleteResponse                `json:"result,nullable"`
	JSON    namespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceDeleteResponseEnvelope]
type namespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceDeleteResponseEnvelopeSuccess bool

const (
	NamespaceDeleteResponseEnvelopeSuccessTrue NamespaceDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      []string            `json:"body,required"`
}

func (r NamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceBulkDeleteResponse                `json:"result,nullable"`
	JSON    namespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkDeleteResponseEnvelope]
type namespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	NamespaceBulkDeleteResponseEnvelopeSuccessTrue NamespaceBulkDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceBulkGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of keys to retrieve (maximum of 100).
	Keys param.Field[[]string] `json:"keys,required"`
	// Whether to parse JSON values in the response.
	Type param.Field[NamespaceBulkGetParamsType] `json:"type"`
	// Whether to include metadata in the response.
	WithMetadata param.Field[bool] `json:"withMetadata"`
}

func (r NamespaceBulkGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to parse JSON values in the response.
type NamespaceBulkGetParamsType string

const (
	NamespaceBulkGetParamsTypeText NamespaceBulkGetParamsType = "text"
	NamespaceBulkGetParamsTypeJson NamespaceBulkGetParamsType = "json"
)

func (r NamespaceBulkGetParamsType) IsKnown() bool {
	switch r {
	case NamespaceBulkGetParamsTypeText, NamespaceBulkGetParamsTypeJson:
		return true
	}
	return false
}

type NamespaceBulkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceBulkGetResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceBulkGetResponse                `json:"result,nullable"`
	JSON    namespaceBulkGetResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceBulkGetResponseEnvelope]
type namespaceBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceBulkGetResponseEnvelopeSuccess bool

const (
	NamespaceBulkGetResponseEnvelopeSuccessTrue NamespaceBulkGetResponseEnvelopeSuccess = true
)

func (r NamespaceBulkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceBulkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceBulkUpdateParams struct {
	// Identifier.
	AccountID param.Field[string]             `path:"account_id,required"`
	Body      []NamespaceBulkUpdateParamsBody `json:"body,required"`
}

func (r NamespaceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceBulkUpdateParamsBody struct {
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

func (r NamespaceBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceBulkUpdateResponse                `json:"result,nullable"`
	JSON    namespaceBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkUpdateResponseEnvelope]
type namespaceBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceBulkUpdateResponseEnvelopeSuccess bool

const (
	NamespaceBulkUpdateResponseEnvelopeSuccessTrue NamespaceBulkUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success NamespaceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Namespace                           `json:"result"`
	JSON    namespaceGetResponseEnvelopeJSON    `json:"-"`
}

// namespaceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceGetResponseEnvelope]
type namespaceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type NamespaceGetResponseEnvelopeSuccess bool

const (
	NamespaceGetResponseEnvelopeSuccessTrue NamespaceGetResponseEnvelopeSuccess = true
)

func (r NamespaceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
