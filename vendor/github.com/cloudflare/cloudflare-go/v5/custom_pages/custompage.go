// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

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

// CustomPageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomPageService] method instead.
type CustomPageService struct {
	Options []option.RequestOption
}

// NewCustomPageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomPageService(opts ...option.RequestOption) (r *CustomPageService) {
	r = &CustomPageService{}
	r.Options = opts
	return
}

// Updates the configuration of an existing custom page.
func (r *CustomPageService) Update(ctx context.Context, identifier CustomPageUpdateParamsIdentifier, params CustomPageUpdateParams, opts ...option.RequestOption) (res *CustomPageUpdateResponse, err error) {
	var env CustomPageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the custom pages.
func (r *CustomPageService) List(ctx context.Context, query CustomPageListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomPageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages", accountOrZone, accountOrZoneID)
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

// Fetches all the custom pages.
func (r *CustomPageService) ListAutoPaging(ctx context.Context, query CustomPageListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomPageListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches the details of a custom page.
func (r *CustomPageService) Get(ctx context.Context, identifier CustomPageGetParamsIdentifier, query CustomPageGetParams, opts ...option.RequestOption) (res *CustomPageGetResponse, err error) {
	var env CustomPageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages/%v", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomPageUpdateResponse struct {
	ID             string    `json:"id"`
	CreatedOn      time.Time `json:"created_on" format:"date-time"`
	Description    string    `json:"description"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	PreviewTarget  string    `json:"preview_target"`
	RequiredTokens []string  `json:"required_tokens"`
	// The custom page state.
	State CustomPageUpdateResponseState `json:"state"`
	// The URL associated with the custom page.
	URL  string                       `json:"url" format:"uri"`
	JSON customPageUpdateResponseJSON `json:"-"`
}

// customPageUpdateResponseJSON contains the JSON metadata for the struct
// [CustomPageUpdateResponse]
type customPageUpdateResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	PreviewTarget  apijson.Field
	RequiredTokens apijson.Field
	State          apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The custom page state.
type CustomPageUpdateResponseState string

const (
	CustomPageUpdateResponseStateDefault    CustomPageUpdateResponseState = "default"
	CustomPageUpdateResponseStateCustomized CustomPageUpdateResponseState = "customized"
)

func (r CustomPageUpdateResponseState) IsKnown() bool {
	switch r {
	case CustomPageUpdateResponseStateDefault, CustomPageUpdateResponseStateCustomized:
		return true
	}
	return false
}

type CustomPageListResponse struct {
	ID             string    `json:"id"`
	CreatedOn      time.Time `json:"created_on" format:"date-time"`
	Description    string    `json:"description"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	PreviewTarget  string    `json:"preview_target"`
	RequiredTokens []string  `json:"required_tokens"`
	// The custom page state.
	State CustomPageListResponseState `json:"state"`
	// The URL associated with the custom page.
	URL  string                     `json:"url" format:"uri"`
	JSON customPageListResponseJSON `json:"-"`
}

// customPageListResponseJSON contains the JSON metadata for the struct
// [CustomPageListResponse]
type customPageListResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	PreviewTarget  apijson.Field
	RequiredTokens apijson.Field
	State          apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageListResponseJSON) RawJSON() string {
	return r.raw
}

// The custom page state.
type CustomPageListResponseState string

const (
	CustomPageListResponseStateDefault    CustomPageListResponseState = "default"
	CustomPageListResponseStateCustomized CustomPageListResponseState = "customized"
)

func (r CustomPageListResponseState) IsKnown() bool {
	switch r {
	case CustomPageListResponseStateDefault, CustomPageListResponseStateCustomized:
		return true
	}
	return false
}

type CustomPageGetResponse struct {
	ID             string    `json:"id"`
	CreatedOn      time.Time `json:"created_on" format:"date-time"`
	Description    string    `json:"description"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	PreviewTarget  string    `json:"preview_target"`
	RequiredTokens []string  `json:"required_tokens"`
	// The custom page state.
	State CustomPageGetResponseState `json:"state"`
	// The URL associated with the custom page.
	URL  string                    `json:"url" format:"uri"`
	JSON customPageGetResponseJSON `json:"-"`
}

// customPageGetResponseJSON contains the JSON metadata for the struct
// [CustomPageGetResponse]
type customPageGetResponseJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	PreviewTarget  apijson.Field
	RequiredTokens apijson.Field
	State          apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseJSON) RawJSON() string {
	return r.raw
}

// The custom page state.
type CustomPageGetResponseState string

const (
	CustomPageGetResponseStateDefault    CustomPageGetResponseState = "default"
	CustomPageGetResponseStateCustomized CustomPageGetResponseState = "customized"
)

func (r CustomPageGetResponseState) IsKnown() bool {
	switch r {
	case CustomPageGetResponseStateDefault, CustomPageGetResponseStateCustomized:
		return true
	}
	return false
}

type CustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[CustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r CustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Error Page Types
type CustomPageUpdateParamsIdentifier string

const (
	CustomPageUpdateParamsIdentifierWAFBlock         CustomPageUpdateParamsIdentifier = "waf_block"
	CustomPageUpdateParamsIdentifierIPBlock          CustomPageUpdateParamsIdentifier = "ip_block"
	CustomPageUpdateParamsIdentifierCountryChallenge CustomPageUpdateParamsIdentifier = "country_challenge"
	CustomPageUpdateParamsIdentifier500Errors        CustomPageUpdateParamsIdentifier = "500_errors"
	CustomPageUpdateParamsIdentifier1000Errors       CustomPageUpdateParamsIdentifier = "1000_errors"
	CustomPageUpdateParamsIdentifierManagedChallenge CustomPageUpdateParamsIdentifier = "managed_challenge"
	CustomPageUpdateParamsIdentifierRatelimitBlock   CustomPageUpdateParamsIdentifier = "ratelimit_block"
)

func (r CustomPageUpdateParamsIdentifier) IsKnown() bool {
	switch r {
	case CustomPageUpdateParamsIdentifierWAFBlock, CustomPageUpdateParamsIdentifierIPBlock, CustomPageUpdateParamsIdentifierCountryChallenge, CustomPageUpdateParamsIdentifier500Errors, CustomPageUpdateParamsIdentifier1000Errors, CustomPageUpdateParamsIdentifierManagedChallenge, CustomPageUpdateParamsIdentifierRatelimitBlock:
		return true
	}
	return false
}

// The custom page state.
type CustomPageUpdateParamsState string

const (
	CustomPageUpdateParamsStateDefault    CustomPageUpdateParamsState = "default"
	CustomPageUpdateParamsStateCustomized CustomPageUpdateParamsState = "customized"
)

func (r CustomPageUpdateParamsState) IsKnown() bool {
	switch r {
	case CustomPageUpdateParamsStateDefault, CustomPageUpdateParamsStateCustomized:
		return true
	}
	return false
}

type CustomPageUpdateResponseEnvelope struct {
	Errors   []CustomPageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomPageUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomPageUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomPageUpdateResponse                `json:"result"`
	JSON    customPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// customPageUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomPageUpdateResponseEnvelope]
type customPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomPageUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           CustomPageUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             customPageUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// customPageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomPageUpdateResponseEnvelopeErrors]
type customPageUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomPageUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    customPageUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customPageUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CustomPageUpdateResponseEnvelopeErrorsSource]
type customPageUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomPageUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CustomPageUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             customPageUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// customPageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomPageUpdateResponseEnvelopeMessages]
type customPageUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomPageUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    customPageUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customPageUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomPageUpdateResponseEnvelopeMessagesSource]
type customPageUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomPageUpdateResponseEnvelopeSuccess bool

const (
	CustomPageUpdateResponseEnvelopeSuccessTrue CustomPageUpdateResponseEnvelopeSuccess = true
)

func (r CustomPageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomPageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomPageListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type CustomPageGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Error Page Types
type CustomPageGetParamsIdentifier string

const (
	CustomPageGetParamsIdentifierWAFBlock         CustomPageGetParamsIdentifier = "waf_block"
	CustomPageGetParamsIdentifierIPBlock          CustomPageGetParamsIdentifier = "ip_block"
	CustomPageGetParamsIdentifierCountryChallenge CustomPageGetParamsIdentifier = "country_challenge"
	CustomPageGetParamsIdentifier500Errors        CustomPageGetParamsIdentifier = "500_errors"
	CustomPageGetParamsIdentifier1000Errors       CustomPageGetParamsIdentifier = "1000_errors"
	CustomPageGetParamsIdentifierManagedChallenge CustomPageGetParamsIdentifier = "managed_challenge"
	CustomPageGetParamsIdentifierRatelimitBlock   CustomPageGetParamsIdentifier = "ratelimit_block"
)

func (r CustomPageGetParamsIdentifier) IsKnown() bool {
	switch r {
	case CustomPageGetParamsIdentifierWAFBlock, CustomPageGetParamsIdentifierIPBlock, CustomPageGetParamsIdentifierCountryChallenge, CustomPageGetParamsIdentifier500Errors, CustomPageGetParamsIdentifier1000Errors, CustomPageGetParamsIdentifierManagedChallenge, CustomPageGetParamsIdentifierRatelimitBlock:
		return true
	}
	return false
}

type CustomPageGetResponseEnvelope struct {
	Errors   []CustomPageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomPageGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomPageGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomPageGetResponse                `json:"result"`
	JSON    customPageGetResponseEnvelopeJSON    `json:"-"`
}

// customPageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomPageGetResponseEnvelope]
type customPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomPageGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           CustomPageGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             customPageGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// customPageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomPageGetResponseEnvelopeErrors]
type customPageGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomPageGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    customPageGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customPageGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CustomPageGetResponseEnvelopeErrorsSource]
type customPageGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomPageGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           CustomPageGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             customPageGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// customPageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomPageGetResponseEnvelopeMessages]
type customPageGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomPageGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    customPageGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customPageGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CustomPageGetResponseEnvelopeMessagesSource]
type customPageGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomPageGetResponseEnvelopeSuccess bool

const (
	CustomPageGetResponseEnvelopeSuccessTrue CustomPageGetResponseEnvelopeSuccess = true
)

func (r CustomPageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomPageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
