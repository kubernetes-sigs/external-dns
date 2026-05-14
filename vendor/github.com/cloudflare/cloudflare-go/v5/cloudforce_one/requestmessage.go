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

// RequestMessageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestMessageService] method instead.
type RequestMessageService struct {
	Options []option.RequestOption
}

// NewRequestMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestMessageService(opts ...option.RequestOption) (r *RequestMessageService) {
	r = &RequestMessageService{}
	r.Options = opts
	return
}

// Create a New Request Message
func (r *RequestMessageService) New(ctx context.Context, requestID string, params RequestMessageNewParams, opts ...option.RequestOption) (res *Message, err error) {
	var env RequestMessageNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/new", params.AccountID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Request Message
func (r *RequestMessageService) Update(ctx context.Context, requestID string, messageID int64, params RequestMessageUpdateParams, opts ...option.RequestOption) (res *Message, err error) {
	var env RequestMessageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", params.AccountID, requestID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Request Message
func (r *RequestMessageService) Delete(ctx context.Context, requestID string, messageID int64, body RequestMessageDeleteParams, opts ...option.RequestOption) (res *RequestMessageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", body.AccountID, requestID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List Request Messages
func (r *RequestMessageService) Get(ctx context.Context, requestID string, params RequestMessageGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[Message], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message", params.AccountID, requestID)
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

// List Request Messages
func (r *RequestMessageService) GetAutoPaging(ctx context.Context, requestID string, params RequestMessageGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Message] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, requestID, params, opts...))
}

type Message struct {
	// Message ID.
	ID int64 `json:"id,required"`
	// Author of message.
	Author string `json:"author,required"`
	// Content of message.
	Content string `json:"content,required"`
	// Whether the message is a follow-on request.
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Defines the message last updated time.
	Updated time.Time `json:"updated,required" format:"date-time"`
	// Defines the message creation time.
	Created time.Time   `json:"created" format:"date-time"`
	JSON    messageJSON `json:"-"`
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponse struct {
	Errors   []RequestMessageDeleteResponseError   `json:"errors,required"`
	Messages []RequestMessageDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestMessageDeleteResponseSuccess `json:"success,required"`
	JSON    requestMessageDeleteResponseJSON    `json:"-"`
}

// requestMessageDeleteResponseJSON contains the JSON metadata for the struct
// [RequestMessageDeleteResponse]
type requestMessageDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           RequestMessageDeleteResponseErrorsSource `json:"source"`
	JSON             requestMessageDeleteResponseErrorJSON    `json:"-"`
}

// requestMessageDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RequestMessageDeleteResponseError]
type requestMessageDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    requestMessageDeleteResponseErrorsSourceJSON `json:"-"`
}

// requestMessageDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [RequestMessageDeleteResponseErrorsSource]
type requestMessageDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           RequestMessageDeleteResponseMessagesSource `json:"source"`
	JSON             requestMessageDeleteResponseMessageJSON    `json:"-"`
}

// requestMessageDeleteResponseMessageJSON contains the JSON metadata for the
// struct [RequestMessageDeleteResponseMessage]
type requestMessageDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RequestMessageDeleteResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    requestMessageDeleteResponseMessagesSourceJSON `json:"-"`
}

// requestMessageDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [RequestMessageDeleteResponseMessagesSource]
type requestMessageDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestMessageDeleteResponseSuccess bool

const (
	RequestMessageDeleteResponseSuccessTrue RequestMessageDeleteResponseSuccess = true
)

func (r RequestMessageDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestMessageDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestMessageNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Content of message.
	Content param.Field[string] `json:"content"`
}

func (r RequestMessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestMessageNewResponseEnvelope struct {
	Errors   []RequestMessageNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestMessageNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Message                                  `json:"result"`
	JSON    requestMessageNewResponseEnvelopeJSON    `json:"-"`
}

// requestMessageNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestMessageNewResponseEnvelope]
type requestMessageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           RequestMessageNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestMessageNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestMessageNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestMessageNewResponseEnvelopeErrors]
type requestMessageNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    requestMessageNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestMessageNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RequestMessageNewResponseEnvelopeErrorsSource]
type requestMessageNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           RequestMessageNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestMessageNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestMessageNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RequestMessageNewResponseEnvelopeMessages]
type requestMessageNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestMessageNewResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    requestMessageNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestMessageNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RequestMessageNewResponseEnvelopeMessagesSource]
type requestMessageNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestMessageNewResponseEnvelopeSuccess bool

const (
	RequestMessageNewResponseEnvelopeSuccessTrue RequestMessageNewResponseEnvelopeSuccess = true
)

func (r RequestMessageNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestMessageNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestMessageUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Content of message.
	Content param.Field[string] `json:"content"`
}

func (r RequestMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestMessageUpdateResponseEnvelope struct {
	Errors   []RequestMessageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestMessageUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestMessageUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Message                                     `json:"result"`
	JSON    requestMessageUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestMessageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestMessageUpdateResponseEnvelope]
type requestMessageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RequestMessageUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestMessageUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestMessageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RequestMessageUpdateResponseEnvelopeErrors]
type requestMessageUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    requestMessageUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestMessageUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestMessageUpdateResponseEnvelopeErrorsSource]
type requestMessageUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           RequestMessageUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestMessageUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestMessageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestMessageUpdateResponseEnvelopeMessages]
type requestMessageUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestMessageUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    requestMessageUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestMessageUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RequestMessageUpdateResponseEnvelopeMessagesSource]
type requestMessageUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestMessageUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestMessageUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestMessageUpdateResponseEnvelopeSuccess bool

const (
	RequestMessageUpdateResponseEnvelopeSuccessTrue RequestMessageUpdateResponseEnvelopeSuccess = true
)

func (r RequestMessageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestMessageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestMessageDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestMessageGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of results.
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page.
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve mes ges created after this time.
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Retrieve messages created before this time.
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Field to sort results by.
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc).
	SortOrder param.Field[RequestMessageGetParamsSortOrder] `json:"sort_order"`
}

func (r RequestMessageGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc).
type RequestMessageGetParamsSortOrder string

const (
	RequestMessageGetParamsSortOrderAsc  RequestMessageGetParamsSortOrder = "asc"
	RequestMessageGetParamsSortOrderDesc RequestMessageGetParamsSortOrder = "desc"
)

func (r RequestMessageGetParamsSortOrder) IsKnown() bool {
	switch r {
	case RequestMessageGetParamsSortOrderAsc, RequestMessageGetParamsSortOrderDesc:
		return true
	}
	return false
}
