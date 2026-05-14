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
)

// RequestPriorityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestPriorityService] method instead.
type RequestPriorityService struct {
	Options []option.RequestOption
}

// NewRequestPriorityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestPriorityService(opts ...option.RequestOption) (r *RequestPriorityService) {
	r = &RequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Intelligence Requirement
func (r *RequestPriorityService) New(ctx context.Context, params RequestPriorityNewParams, opts ...option.RequestOption) (res *Priority, err error) {
	var env RequestPriorityNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *RequestPriorityService) Update(ctx context.Context, priorityID string, params RequestPriorityUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	var env RequestPriorityUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if priorityID == "" {
		err = errors.New("missing required priority_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", params.AccountID, priorityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Requirement
func (r *RequestPriorityService) Delete(ctx context.Context, priorityID string, body RequestPriorityDeleteParams, opts ...option.RequestOption) (res *RequestPriorityDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if priorityID == "" {
		err = errors.New("missing required priority_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", body.AccountID, priorityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a Priority Intelligence Requirement
func (r *RequestPriorityService) Get(ctx context.Context, priorityID string, query RequestPriorityGetParams, opts ...option.RequestOption) (res *Item, err error) {
	var env RequestPriorityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if priorityID == "" {
		err = errors.New("missing required priority_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", query.AccountID, priorityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *RequestPriorityService) Quota(ctx context.Context, query RequestPriorityQuotaParams, opts ...option.RequestOption) (res *Quota, err error) {
	var env RequestPriorityQuotaResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Label = string

type LabelParam = string

type Priority struct {
	// UUID.
	ID string `json:"id,required"`
	// Priority creation time.
	Created time.Time `json:"created,required" format:"date-time"`
	// List of labels.
	Labels []Label `json:"labels,required"`
	// Priority.
	Priority int64 `json:"priority,required"`
	// Requirement.
	Requirement string `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP PriorityTLP `json:"tlp,required"`
	// Priority last updated time.
	Updated time.Time    `json:"updated,required" format:"date-time"`
	JSON    priorityJSON `json:"-"`
}

// priorityJSON contains the JSON metadata for the struct [Priority]
type priorityJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	TLP         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Priority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priorityJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP).
type PriorityTLP string

const (
	PriorityTLPClear       PriorityTLP = "clear"
	PriorityTLPAmber       PriorityTLP = "amber"
	PriorityTLPAmberStrict PriorityTLP = "amber-strict"
	PriorityTLPGreen       PriorityTLP = "green"
	PriorityTLPRed         PriorityTLP = "red"
)

func (r PriorityTLP) IsKnown() bool {
	switch r {
	case PriorityTLPClear, PriorityTLPAmber, PriorityTLPAmberStrict, PriorityTLPGreen, PriorityTLPRed:
		return true
	}
	return false
}

type PriorityEditParam struct {
	// List of labels.
	Labels param.Field[[]LabelParam] `json:"labels,required"`
	// Priority.
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement.
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP).
	TLP param.Field[PriorityEditTLP] `json:"tlp,required"`
}

func (r PriorityEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP).
type PriorityEditTLP string

const (
	PriorityEditTLPClear       PriorityEditTLP = "clear"
	PriorityEditTLPAmber       PriorityEditTLP = "amber"
	PriorityEditTLPAmberStrict PriorityEditTLP = "amber-strict"
	PriorityEditTLPGreen       PriorityEditTLP = "green"
	PriorityEditTLPRed         PriorityEditTLP = "red"
)

func (r PriorityEditTLP) IsKnown() bool {
	switch r {
	case PriorityEditTLPClear, PriorityEditTLPAmber, PriorityEditTLPAmberStrict, PriorityEditTLPGreen, PriorityEditTLPRed:
		return true
	}
	return false
}

type RequestPriorityDeleteResponse struct {
	Errors   []RequestPriorityDeleteResponseError   `json:"errors,required"`
	Messages []RequestPriorityDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestPriorityDeleteResponseSuccess `json:"success,required"`
	JSON    requestPriorityDeleteResponseJSON    `json:"-"`
}

// requestPriorityDeleteResponseJSON contains the JSON metadata for the struct
// [RequestPriorityDeleteResponse]
type requestPriorityDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           RequestPriorityDeleteResponseErrorsSource `json:"source"`
	JSON             requestPriorityDeleteResponseErrorJSON    `json:"-"`
}

// requestPriorityDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RequestPriorityDeleteResponseError]
type requestPriorityDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    requestPriorityDeleteResponseErrorsSourceJSON `json:"-"`
}

// requestPriorityDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [RequestPriorityDeleteResponseErrorsSource]
type requestPriorityDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           RequestPriorityDeleteResponseMessagesSource `json:"source"`
	JSON             requestPriorityDeleteResponseMessageJSON    `json:"-"`
}

// requestPriorityDeleteResponseMessageJSON contains the JSON metadata for the
// struct [RequestPriorityDeleteResponseMessage]
type requestPriorityDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityDeleteResponseMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    requestPriorityDeleteResponseMessagesSourceJSON `json:"-"`
}

// requestPriorityDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [RequestPriorityDeleteResponseMessagesSource]
type requestPriorityDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestPriorityDeleteResponseSuccess bool

const (
	RequestPriorityDeleteResponseSuccessTrue RequestPriorityDeleteResponseSuccess = true
)

func (r RequestPriorityDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityNewParams struct {
	// Identifier.
	AccountID    param.Field[string] `path:"account_id,required"`
	PriorityEdit PriorityEditParam   `json:"priority_edit,required"`
}

func (r RequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type RequestPriorityNewResponseEnvelope struct {
	Errors   []RequestPriorityNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Priority                                  `json:"result"`
	JSON    requestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityNewResponseEnvelope]
type requestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RequestPriorityNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestPriorityNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestPriorityNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestPriorityNewResponseEnvelopeErrors]
type requestPriorityNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    requestPriorityNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestPriorityNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestPriorityNewResponseEnvelopeErrorsSource]
type requestPriorityNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RequestPriorityNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestPriorityNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestPriorityNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityNewResponseEnvelopeMessages]
type requestPriorityNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    requestPriorityNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestPriorityNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RequestPriorityNewResponseEnvelopeMessagesSource]
type requestPriorityNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestPriorityNewResponseEnvelopeSuccess bool

const (
	RequestPriorityNewResponseEnvelopeSuccessTrue RequestPriorityNewResponseEnvelopeSuccess = true
)

func (r RequestPriorityNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityUpdateParams struct {
	// Identifier.
	AccountID    param.Field[string] `path:"account_id,required"`
	PriorityEdit PriorityEditParam   `json:"priority_edit,required"`
}

func (r RequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type RequestPriorityUpdateResponseEnvelope struct {
	Errors   []RequestPriorityUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Item                                         `json:"result"`
	JSON    requestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityUpdateResponseEnvelope]
type requestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           RequestPriorityUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestPriorityUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestPriorityUpdateResponseEnvelopeErrors]
type requestPriorityUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    requestPriorityUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestPriorityUpdateResponseEnvelopeErrorsSource]
type requestPriorityUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           RequestPriorityUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestPriorityUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityUpdateResponseEnvelopeMessages]
type requestPriorityUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    requestPriorityUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RequestPriorityUpdateResponseEnvelopeMessagesSource]
type requestPriorityUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	RequestPriorityUpdateResponseEnvelopeSuccessTrue RequestPriorityUpdateResponseEnvelopeSuccess = true
)

func (r RequestPriorityUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestPriorityGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestPriorityGetResponseEnvelope struct {
	Errors   []RequestPriorityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Item                                      `json:"result"`
	JSON    requestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityGetResponseEnvelope]
type requestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RequestPriorityGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestPriorityGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestPriorityGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestPriorityGetResponseEnvelopeErrors]
type requestPriorityGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    requestPriorityGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestPriorityGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestPriorityGetResponseEnvelopeErrorsSource]
type requestPriorityGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RequestPriorityGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestPriorityGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestPriorityGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityGetResponseEnvelopeMessages]
type requestPriorityGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    requestPriorityGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestPriorityGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RequestPriorityGetResponseEnvelopeMessagesSource]
type requestPriorityGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestPriorityGetResponseEnvelopeSuccess bool

const (
	RequestPriorityGetResponseEnvelopeSuccessTrue RequestPriorityGetResponseEnvelopeSuccess = true
)

func (r RequestPriorityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityQuotaParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestPriorityQuotaResponseEnvelope struct {
	Errors   []RequestPriorityQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestPriorityQuotaResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	Result  Quota                                       `json:"result"`
	JSON    requestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityQuotaResponseEnvelope]
type requestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RequestPriorityQuotaResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestPriorityQuotaResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestPriorityQuotaResponseEnvelopeErrors]
type requestPriorityQuotaResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    requestPriorityQuotaResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestPriorityQuotaResponseEnvelopeErrorsSource]
type requestPriorityQuotaResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           RequestPriorityQuotaResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestPriorityQuotaResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestPriorityQuotaResponseEnvelopeMessages]
type requestPriorityQuotaResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestPriorityQuotaResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    requestPriorityQuotaResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RequestPriorityQuotaResponseEnvelopeMessagesSource]
type requestPriorityQuotaResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	RequestPriorityQuotaResponseEnvelopeSuccessTrue RequestPriorityQuotaResponseEnvelopeSuccess = true
)

func (r RequestPriorityQuotaResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityQuotaResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
