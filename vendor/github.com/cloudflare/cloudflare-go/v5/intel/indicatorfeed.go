// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// IndicatorFeedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedService] method instead.
type IndicatorFeedService struct {
	Options     []option.RequestOption
	Snapshots   *IndicatorFeedSnapshotService
	Permissions *IndicatorFeedPermissionService
	Downloads   *IndicatorFeedDownloadService
}

// NewIndicatorFeedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIndicatorFeedService(opts ...option.RequestOption) (r *IndicatorFeedService) {
	r = &IndicatorFeedService{}
	r.Options = opts
	r.Snapshots = NewIndicatorFeedSnapshotService(opts...)
	r.Permissions = NewIndicatorFeedPermissionService(opts...)
	r.Downloads = NewIndicatorFeedDownloadService(opts...)
	return
}

// Create new indicator feed
func (r *IndicatorFeedService) New(ctx context.Context, params IndicatorFeedNewParams, opts ...option.RequestOption) (res *IndicatorFeedNewResponse, err error) {
	var env IndicatorFeedNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update indicator feed metadata
func (r *IndicatorFeedService) Update(ctx context.Context, feedID int64, params IndicatorFeedUpdateParams, opts ...option.RequestOption) (res *IndicatorFeedUpdateResponse, err error) {
	var env IndicatorFeedUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get indicator feeds owned by this account
func (r *IndicatorFeedService) List(ctx context.Context, query IndicatorFeedListParams, opts ...option.RequestOption) (res *pagination.SinglePage[IndicatorFeedListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get indicator feeds owned by this account
func (r *IndicatorFeedService) ListAutoPaging(ctx context.Context, query IndicatorFeedListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IndicatorFeedListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Get indicator feed data
func (r *IndicatorFeedService) Data(ctx context.Context, feedID int64, query IndicatorFeedDataParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *IndicatorFeedService) Get(ctx context.Context, feedID int64, query IndicatorFeedGetParams, opts ...option.RequestOption) (res *IndicatorFeedGetResponse, err error) {
	var env IndicatorFeedGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedNewResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                       `json:"name"`
	JSON indicatorFeedNewResponseJSON `json:"-"`
}

// indicatorFeedNewResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedNewResponse]
type indicatorFeedNewResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IndicatorFeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                          `json:"name"`
	JSON indicatorFeedUpdateResponseJSON `json:"-"`
}

// indicatorFeedUpdateResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedUpdateResponse]
type indicatorFeedUpdateResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string                        `json:"name"`
	JSON indicatorFeedListResponseJSON `json:"-"`
}

// indicatorFeedListResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedListResponse]
type indicatorFeedListResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IndicatorFeedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedListResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// Status of the latest snapshot uploaded
	LatestUploadStatus IndicatorFeedGetResponseLatestUploadStatus `json:"latest_upload_status"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string `json:"name"`
	// The unique identifier for the provider
	ProviderID string `json:"provider_id"`
	// The provider of the indicator feed
	ProviderName string                       `json:"provider_name"`
	JSON         indicatorFeedGetResponseJSON `json:"-"`
}

// indicatorFeedGetResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedGetResponse]
type indicatorFeedGetResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	IsAttributable     apijson.Field
	IsDownloadable     apijson.Field
	IsPublic           apijson.Field
	LatestUploadStatus apijson.Field
	ModifiedOn         apijson.Field
	Name               apijson.Field
	ProviderID         apijson.Field
	ProviderName       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IndicatorFeedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the latest snapshot uploaded
type IndicatorFeedGetResponseLatestUploadStatus string

const (
	IndicatorFeedGetResponseLatestUploadStatusMirroring    IndicatorFeedGetResponseLatestUploadStatus = "Mirroring"
	IndicatorFeedGetResponseLatestUploadStatusUnifying     IndicatorFeedGetResponseLatestUploadStatus = "Unifying"
	IndicatorFeedGetResponseLatestUploadStatusLoading      IndicatorFeedGetResponseLatestUploadStatus = "Loading"
	IndicatorFeedGetResponseLatestUploadStatusProvisioning IndicatorFeedGetResponseLatestUploadStatus = "Provisioning"
	IndicatorFeedGetResponseLatestUploadStatusComplete     IndicatorFeedGetResponseLatestUploadStatus = "Complete"
	IndicatorFeedGetResponseLatestUploadStatusError        IndicatorFeedGetResponseLatestUploadStatus = "Error"
)

func (r IndicatorFeedGetResponseLatestUploadStatus) IsKnown() bool {
	switch r {
	case IndicatorFeedGetResponseLatestUploadStatusMirroring, IndicatorFeedGetResponseLatestUploadStatusUnifying, IndicatorFeedGetResponseLatestUploadStatusLoading, IndicatorFeedGetResponseLatestUploadStatusProvisioning, IndicatorFeedGetResponseLatestUploadStatusComplete, IndicatorFeedGetResponseLatestUploadStatusError:
		return true
	}
	return false
}

type IndicatorFeedNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The description of the example test
	Description param.Field[string] `json:"description"`
	// The name of the indicator feed
	Name param.Field[string] `json:"name"`
}

func (r IndicatorFeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedNewResponseEnvelope struct {
	Errors   []IndicatorFeedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedNewResponse                `json:"result"`
	JSON    indicatorFeedNewResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndicatorFeedNewResponseEnvelope]
type indicatorFeedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           IndicatorFeedNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedNewResponseEnvelopeErrors]
type indicatorFeedNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    indicatorFeedNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [IndicatorFeedNewResponseEnvelopeErrorsSource]
type indicatorFeedNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           IndicatorFeedNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndicatorFeedNewResponseEnvelopeMessages]
type indicatorFeedNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    indicatorFeedNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [IndicatorFeedNewResponseEnvelopeMessagesSource]
type indicatorFeedNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedNewResponseEnvelopeSuccessTrue IndicatorFeedNewResponseEnvelopeSuccess = true
)

func (r IndicatorFeedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The new description of the feed
	Description param.Field[string] `json:"description"`
	// The new is_attributable value of the feed
	IsAttributable param.Field[bool] `json:"is_attributable"`
	// The new is_downloadable value of the feed
	IsDownloadable param.Field[bool] `json:"is_downloadable"`
	// The new is_public value of the feed
	IsPublic param.Field[bool] `json:"is_public"`
	// The new name of the feed
	Name param.Field[string] `json:"name"`
}

func (r IndicatorFeedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedUpdateResponseEnvelope struct {
	Errors   []IndicatorFeedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedUpdateResponse                `json:"result"`
	JSON    indicatorFeedUpdateResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndicatorFeedUpdateResponseEnvelope]
type indicatorFeedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           IndicatorFeedUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedUpdateResponseEnvelopeErrors]
type indicatorFeedUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    indicatorFeedUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [IndicatorFeedUpdateResponseEnvelopeErrorsSource]
type indicatorFeedUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           IndicatorFeedUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IndicatorFeedUpdateResponseEnvelopeMessages]
type indicatorFeedUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    indicatorFeedUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [IndicatorFeedUpdateResponseEnvelopeMessagesSource]
type indicatorFeedUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedUpdateResponseEnvelopeSuccess bool

const (
	IndicatorFeedUpdateResponseEnvelopeSuccessTrue IndicatorFeedUpdateResponseEnvelopeSuccess = true
)

func (r IndicatorFeedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedDataParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedGetResponseEnvelope struct {
	Errors   []IndicatorFeedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedGetResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedGetResponse                `json:"result"`
	JSON    indicatorFeedGetResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndicatorFeedGetResponseEnvelope]
type indicatorFeedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           IndicatorFeedGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndicatorFeedGetResponseEnvelopeErrors]
type indicatorFeedGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    indicatorFeedGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [IndicatorFeedGetResponseEnvelopeErrorsSource]
type indicatorFeedGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           IndicatorFeedGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndicatorFeedGetResponseEnvelopeMessages]
type indicatorFeedGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    indicatorFeedGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [IndicatorFeedGetResponseEnvelopeMessagesSource]
type indicatorFeedGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedGetResponseEnvelopeSuccess bool

const (
	IndicatorFeedGetResponseEnvelopeSuccessTrue IndicatorFeedGetResponseEnvelopeSuccess = true
)

func (r IndicatorFeedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
