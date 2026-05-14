// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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

// RequestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestService] method instead.
type RequestService struct {
	Options  []option.RequestOption
	Message  *RequestMessageService
	Priority *RequestPriorityService
	Assets   *RequestAssetService
}

// NewRequestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRequestService(opts ...option.RequestOption) (r *RequestService) {
	r = &RequestService{}
	r.Options = opts
	r.Message = NewRequestMessageService(opts...)
	r.Priority = NewRequestPriorityService(opts...)
	r.Assets = NewRequestAssetService(opts...)
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided, a default will be assigned.
func (r *RequestService) New(ctx context.Context, params RequestNewParams, opts ...option.RequestOption) (res *Item, err error) {
	var env RequestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/new", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a request alters the request in the Cloudforce One queue. This API may
// be used to update any attributes of the request after the initial submission.
// Only fields that you choose to update need to be add to the request body.
func (r *RequestService) Update(ctx context.Context, requestID string, params RequestUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	var env RequestUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", params.AccountID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Requests
func (r *RequestService) List(ctx context.Context, params RequestListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ListItem], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List Requests
func (r *RequestService) ListAutoPaging(ctx context.Context, params RequestListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ListItem] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a Request
func (r *RequestService) Delete(ctx context.Context, requestID string, body RequestDeleteParams, opts ...option.RequestOption) (res *RequestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", body.AccountID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Request Priority, Status, and TLP constants
func (r *RequestService) Constants(ctx context.Context, query RequestConstantsParams, opts ...option.RequestOption) (res *RequestConstants, err error) {
	var env RequestConstantsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/constants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Request
func (r *RequestService) Get(ctx context.Context, requestID string, query RequestGetParams, opts ...option.RequestOption) (res *Item, err error) {
	var env RequestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", query.AccountID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Quota
func (r *RequestService) Quota(ctx context.Context, query RequestQuotaParams, opts ...option.RequestOption) (res *Quota, err error) {
	var env RequestQuotaResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/quota", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Request Types
func (r *RequestService) Types(ctx context.Context, query RequestTypesParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/types", query.AccountID)
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

// Get Request Types
func (r *RequestService) TypesAutoPaging(ctx context.Context, query RequestTypesParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.Types(ctx, query, opts...))
}

type Item struct {
	// UUID.
	ID string `json:"id,required"`
	// Request content.
	Content  string    `json:"content,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	Priority time.Time `json:"priority,required" format:"date-time"`
	// Requested information from request.
	Request string `json:"request,required"`
	// Brief description of the request.
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP       ItemTLP   `json:"tlp,required"`
	Updated   time.Time `json:"updated,required" format:"date-time"`
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages.
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID.
	ReadableID string `json:"readable_id"`
	// Request Status.
	Status ItemStatus `json:"status"`
	// Tokens for the request.
	Tokens int64    `json:"tokens"`
	JSON   itemJSON `json:"-"`
}

// itemJSON contains the JSON metadata for the struct [Item]
type itemJSON struct {
	ID            apijson.Field
	Content       apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	TLP           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Item) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP).
type ItemTLP string

const (
	ItemTLPClear       ItemTLP = "clear"
	ItemTLPAmber       ItemTLP = "amber"
	ItemTLPAmberStrict ItemTLP = "amber-strict"
	ItemTLPGreen       ItemTLP = "green"
	ItemTLPRed         ItemTLP = "red"
)

func (r ItemTLP) IsKnown() bool {
	switch r {
	case ItemTLPClear, ItemTLPAmber, ItemTLPAmberStrict, ItemTLPGreen, ItemTLPRed:
		return true
	}
	return false
}

// Request Status.
type ItemStatus string

const (
	ItemStatusOpen      ItemStatus = "open"
	ItemStatusAccepted  ItemStatus = "accepted"
	ItemStatusReported  ItemStatus = "reported"
	ItemStatusApproved  ItemStatus = "approved"
	ItemStatusCompleted ItemStatus = "completed"
	ItemStatusDeclined  ItemStatus = "declined"
)

func (r ItemStatus) IsKnown() bool {
	switch r {
	case ItemStatusOpen, ItemStatusAccepted, ItemStatusReported, ItemStatusApproved, ItemStatusCompleted, ItemStatusDeclined:
		return true
	}
	return false
}

type ListItem struct {
	// UUID.
	ID string `json:"id,required"`
	// Request creation time.
	Created  time.Time        `json:"created,required" format:"date-time"`
	Priority ListItemPriority `json:"priority,required"`
	// Requested information from request.
	Request string `json:"request,required"`
	// Brief description of the request.
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP ListItemTLP `json:"tlp,required"`
	// Request last updated time.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Request completion time.
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages.
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID.
	ReadableID string `json:"readable_id"`
	// Request Status.
	Status ListItemStatus `json:"status"`
	// Tokens for the request.
	Tokens int64        `json:"tokens"`
	JSON   listItemJSON `json:"-"`
}

// listItemJSON contains the JSON metadata for the struct [ListItem]
type listItemJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	TLP           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemJSON) RawJSON() string {
	return r.raw
}

type ListItemPriority string

const (
	ListItemPriorityRoutine ListItemPriority = "routine"
	ListItemPriorityHigh    ListItemPriority = "high"
	ListItemPriorityUrgent  ListItemPriority = "urgent"
)

func (r ListItemPriority) IsKnown() bool {
	switch r {
	case ListItemPriorityRoutine, ListItemPriorityHigh, ListItemPriorityUrgent:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP).
type ListItemTLP string

const (
	ListItemTLPClear       ListItemTLP = "clear"
	ListItemTLPAmber       ListItemTLP = "amber"
	ListItemTLPAmberStrict ListItemTLP = "amber-strict"
	ListItemTLPGreen       ListItemTLP = "green"
	ListItemTLPRed         ListItemTLP = "red"
)

func (r ListItemTLP) IsKnown() bool {
	switch r {
	case ListItemTLPClear, ListItemTLPAmber, ListItemTLPAmberStrict, ListItemTLPGreen, ListItemTLPRed:
		return true
	}
	return false
}

// Request Status.
type ListItemStatus string

const (
	ListItemStatusOpen      ListItemStatus = "open"
	ListItemStatusAccepted  ListItemStatus = "accepted"
	ListItemStatusReported  ListItemStatus = "reported"
	ListItemStatusApproved  ListItemStatus = "approved"
	ListItemStatusCompleted ListItemStatus = "completed"
	ListItemStatusDeclined  ListItemStatus = "declined"
)

func (r ListItemStatus) IsKnown() bool {
	switch r {
	case ListItemStatusOpen, ListItemStatusAccepted, ListItemStatusReported, ListItemStatusApproved, ListItemStatusCompleted, ListItemStatusDeclined:
		return true
	}
	return false
}

type Quota struct {
	// Anniversary date is when annual quota limit is refreshed.
	AnniversaryDate time.Time `json:"anniversary_date" format:"date-time"`
	// Quarter anniversary date is when quota limit is refreshed each quarter.
	QuarterAnniversaryDate time.Time `json:"quarter_anniversary_date" format:"date-time"`
	// Tokens for the quarter.
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter.
	Remaining int64     `json:"remaining"`
	JSON      quotaJSON `json:"-"`
}

// quotaJSON contains the JSON metadata for the struct [Quota]
type quotaJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Quota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaJSON) RawJSON() string {
	return r.raw
}

type RequestConstants struct {
	Priority []RequestConstantsPriority `json:"priority"`
	Status   []RequestConstantsStatus   `json:"status"`
	TLP      []RequestConstantsTLP      `json:"tlp"`
	JSON     requestConstantsJSON       `json:"-"`
}

// requestConstantsJSON contains the JSON metadata for the struct
// [RequestConstants]
type requestConstantsJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsPriority string

const (
	RequestConstantsPriorityRoutine RequestConstantsPriority = "routine"
	RequestConstantsPriorityHigh    RequestConstantsPriority = "high"
	RequestConstantsPriorityUrgent  RequestConstantsPriority = "urgent"
)

func (r RequestConstantsPriority) IsKnown() bool {
	switch r {
	case RequestConstantsPriorityRoutine, RequestConstantsPriorityHigh, RequestConstantsPriorityUrgent:
		return true
	}
	return false
}

// Request Status.
type RequestConstantsStatus string

const (
	RequestConstantsStatusOpen      RequestConstantsStatus = "open"
	RequestConstantsStatusAccepted  RequestConstantsStatus = "accepted"
	RequestConstantsStatusReported  RequestConstantsStatus = "reported"
	RequestConstantsStatusApproved  RequestConstantsStatus = "approved"
	RequestConstantsStatusCompleted RequestConstantsStatus = "completed"
	RequestConstantsStatusDeclined  RequestConstantsStatus = "declined"
)

func (r RequestConstantsStatus) IsKnown() bool {
	switch r {
	case RequestConstantsStatusOpen, RequestConstantsStatusAccepted, RequestConstantsStatusReported, RequestConstantsStatusApproved, RequestConstantsStatusCompleted, RequestConstantsStatusDeclined:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP).
type RequestConstantsTLP string

const (
	RequestConstantsTLPClear       RequestConstantsTLP = "clear"
	RequestConstantsTLPAmber       RequestConstantsTLP = "amber"
	RequestConstantsTLPAmberStrict RequestConstantsTLP = "amber-strict"
	RequestConstantsTLPGreen       RequestConstantsTLP = "green"
	RequestConstantsTLPRed         RequestConstantsTLP = "red"
)

func (r RequestConstantsTLP) IsKnown() bool {
	switch r {
	case RequestConstantsTLPClear, RequestConstantsTLPAmber, RequestConstantsTLPAmberStrict, RequestConstantsTLPGreen, RequestConstantsTLPRed:
		return true
	}
	return false
}

type RequestTypes []string

type RequestDeleteResponse struct {
	Errors   []RequestDeleteResponseError   `json:"errors,required"`
	Messages []RequestDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestDeleteResponseSuccess `json:"success,required"`
	JSON    requestDeleteResponseJSON    `json:"-"`
}

// requestDeleteResponseJSON contains the JSON metadata for the struct
// [RequestDeleteResponse]
type requestDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           RequestDeleteResponseErrorsSource `json:"source"`
	JSON             requestDeleteResponseErrorJSON    `json:"-"`
}

// requestDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RequestDeleteResponseError]
type requestDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    requestDeleteResponseErrorsSourceJSON `json:"-"`
}

// requestDeleteResponseErrorsSourceJSON contains the JSON metadata for the struct
// [RequestDeleteResponseErrorsSource]
type requestDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           RequestDeleteResponseMessagesSource `json:"source"`
	JSON             requestDeleteResponseMessageJSON    `json:"-"`
}

// requestDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RequestDeleteResponseMessage]
type requestDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RequestDeleteResponseMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    requestDeleteResponseMessagesSourceJSON `json:"-"`
}

// requestDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [RequestDeleteResponseMessagesSource]
type requestDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestDeleteResponseSuccess bool

const (
	RequestDeleteResponseSuccessTrue RequestDeleteResponseSuccess = true
)

func (r RequestDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Request content.
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request.
	Priority param.Field[string] `json:"priority"`
	// Requested information from request.
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request.
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP param.Field[RequestNewParamsTLP] `json:"tlp"`
}

func (r RequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP).
type RequestNewParamsTLP string

const (
	RequestNewParamsTLPClear       RequestNewParamsTLP = "clear"
	RequestNewParamsTLPAmber       RequestNewParamsTLP = "amber"
	RequestNewParamsTLPAmberStrict RequestNewParamsTLP = "amber-strict"
	RequestNewParamsTLPGreen       RequestNewParamsTLP = "green"
	RequestNewParamsTLPRed         RequestNewParamsTLP = "red"
)

func (r RequestNewParamsTLP) IsKnown() bool {
	switch r {
	case RequestNewParamsTLPClear, RequestNewParamsTLPAmber, RequestNewParamsTLPAmberStrict, RequestNewParamsTLPGreen, RequestNewParamsTLPRed:
		return true
	}
	return false
}

type RequestNewResponseEnvelope struct {
	Errors   []RequestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Item                              `json:"result"`
	JSON    requestNewResponseEnvelopeJSON    `json:"-"`
}

// requestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelope]
type requestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           RequestNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelopeErrors]
type requestNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    requestNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RequestNewResponseEnvelopeErrorsSource]
type requestNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           RequestNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RequestNewResponseEnvelopeMessages]
type requestNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    requestNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RequestNewResponseEnvelopeMessagesSource]
type requestNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestNewResponseEnvelopeSuccess bool

const (
	RequestNewResponseEnvelopeSuccessTrue RequestNewResponseEnvelopeSuccess = true
)

func (r RequestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Request content.
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request.
	Priority param.Field[string] `json:"priority"`
	// Requested information from request.
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request.
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP param.Field[RequestUpdateParamsTLP] `json:"tlp"`
}

func (r RequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP).
type RequestUpdateParamsTLP string

const (
	RequestUpdateParamsTLPClear       RequestUpdateParamsTLP = "clear"
	RequestUpdateParamsTLPAmber       RequestUpdateParamsTLP = "amber"
	RequestUpdateParamsTLPAmberStrict RequestUpdateParamsTLP = "amber-strict"
	RequestUpdateParamsTLPGreen       RequestUpdateParamsTLP = "green"
	RequestUpdateParamsTLPRed         RequestUpdateParamsTLP = "red"
)

func (r RequestUpdateParamsTLP) IsKnown() bool {
	switch r {
	case RequestUpdateParamsTLPClear, RequestUpdateParamsTLPAmber, RequestUpdateParamsTLPAmberStrict, RequestUpdateParamsTLPGreen, RequestUpdateParamsTLPRed:
		return true
	}
	return false
}

type RequestUpdateResponseEnvelope struct {
	Errors   []RequestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Item                                 `json:"result"`
	JSON    requestUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestUpdateResponseEnvelope]
type requestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           RequestUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestUpdateResponseEnvelopeErrors]
type requestUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    requestUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RequestUpdateResponseEnvelopeErrorsSource]
type requestUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           RequestUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestUpdateResponseEnvelopeMessages]
type requestUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    requestUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RequestUpdateResponseEnvelopeMessagesSource]
type requestUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestUpdateResponseEnvelopeSuccess bool

const (
	RequestUpdateResponseEnvelopeSuccessTrue RequestUpdateResponseEnvelopeSuccess = true
)

func (r RequestUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of results.
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page.
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve requests completed after this time.
	CompletedAfter param.Field[time.Time] `json:"completed_after" format:"date-time"`
	// Retrieve requests completed before this time.
	CompletedBefore param.Field[time.Time] `json:"completed_before" format:"date-time"`
	// Retrieve requests created after this time.
	CreatedAfter param.Field[time.Time] `json:"created_after" format:"date-time"`
	// Retrieve requests created before this time.
	CreatedBefore param.Field[time.Time] `json:"created_before" format:"date-time"`
	// Requested information from request.
	RequestType param.Field[string] `json:"request_type"`
	// Field to sort results by.
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc).
	SortOrder param.Field[RequestListParamsSortOrder] `json:"sort_order"`
	// Request Status.
	Status param.Field[RequestListParamsStatus] `json:"status"`
}

func (r RequestListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc).
type RequestListParamsSortOrder string

const (
	RequestListParamsSortOrderAsc  RequestListParamsSortOrder = "asc"
	RequestListParamsSortOrderDesc RequestListParamsSortOrder = "desc"
)

func (r RequestListParamsSortOrder) IsKnown() bool {
	switch r {
	case RequestListParamsSortOrderAsc, RequestListParamsSortOrderDesc:
		return true
	}
	return false
}

// Request Status.
type RequestListParamsStatus string

const (
	RequestListParamsStatusOpen      RequestListParamsStatus = "open"
	RequestListParamsStatusAccepted  RequestListParamsStatus = "accepted"
	RequestListParamsStatusReported  RequestListParamsStatus = "reported"
	RequestListParamsStatusApproved  RequestListParamsStatus = "approved"
	RequestListParamsStatusCompleted RequestListParamsStatus = "completed"
	RequestListParamsStatusDeclined  RequestListParamsStatus = "declined"
)

func (r RequestListParamsStatus) IsKnown() bool {
	switch r {
	case RequestListParamsStatusOpen, RequestListParamsStatusAccepted, RequestListParamsStatusReported, RequestListParamsStatusApproved, RequestListParamsStatusCompleted, RequestListParamsStatusDeclined:
		return true
	}
	return false
}

type RequestDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestConstantsParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestConstantsResponseEnvelope struct {
	Errors   []RequestConstantsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestConstantsResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestConstantsResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestConstants                        `json:"result"`
	JSON    requestConstantsResponseEnvelopeJSON    `json:"-"`
}

// requestConstantsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestConstantsResponseEnvelope]
type requestConstantsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RequestConstantsResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestConstantsResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestConstantsResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestConstantsResponseEnvelopeErrors]
type requestConstantsResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    requestConstantsResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestConstantsResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RequestConstantsResponseEnvelopeErrorsSource]
type requestConstantsResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RequestConstantsResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestConstantsResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestConstantsResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestConstantsResponseEnvelopeMessages]
type requestConstantsResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestConstantsResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    requestConstantsResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestConstantsResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RequestConstantsResponseEnvelopeMessagesSource]
type requestConstantsResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConstantsResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConstantsResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestConstantsResponseEnvelopeSuccess bool

const (
	RequestConstantsResponseEnvelopeSuccessTrue RequestConstantsResponseEnvelopeSuccess = true
)

func (r RequestConstantsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestConstantsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestGetResponseEnvelope struct {
	Errors   []RequestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Item                              `json:"result"`
	JSON    requestGetResponseEnvelopeJSON    `json:"-"`
}

// requestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelope]
type requestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           RequestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelopeErrors]
type requestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    requestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RequestGetResponseEnvelopeErrorsSource]
type requestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           RequestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RequestGetResponseEnvelopeMessages]
type requestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    requestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RequestGetResponseEnvelopeMessagesSource]
type requestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestGetResponseEnvelopeSuccess bool

const (
	RequestGetResponseEnvelopeSuccessTrue RequestGetResponseEnvelopeSuccess = true
)

func (r RequestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestQuotaParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestQuotaResponseEnvelope struct {
	Errors   []RequestQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestQuotaResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestQuotaResponseEnvelopeSuccess `json:"success,required"`
	Result  Quota                               `json:"result"`
	JSON    requestQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestQuotaResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestQuotaResponseEnvelope]
type requestQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           RequestQuotaResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestQuotaResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestQuotaResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RequestQuotaResponseEnvelopeErrors]
type requestQuotaResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    requestQuotaResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestQuotaResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RequestQuotaResponseEnvelopeErrorsSource]
type requestQuotaResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           RequestQuotaResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestQuotaResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestQuotaResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestQuotaResponseEnvelopeMessages]
type requestQuotaResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestQuotaResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    requestQuotaResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestQuotaResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RequestQuotaResponseEnvelopeMessagesSource]
type requestQuotaResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestQuotaResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestQuotaResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestQuotaResponseEnvelopeSuccess bool

const (
	RequestQuotaResponseEnvelopeSuccessTrue RequestQuotaResponseEnvelopeSuccess = true
)

func (r RequestQuotaResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestQuotaResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestTypesParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
