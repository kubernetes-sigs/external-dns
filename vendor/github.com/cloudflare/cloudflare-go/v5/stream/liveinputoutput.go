// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// LiveInputOutputService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLiveInputOutputService] method instead.
type LiveInputOutputService struct {
	Options []option.RequestOption
}

// NewLiveInputOutputService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLiveInputOutputService(opts ...option.RequestOption) (r *LiveInputOutputService) {
	r = &LiveInputOutputService{}
	r.Options = opts
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *LiveInputOutputService) New(ctx context.Context, liveInputIdentifier string, params LiveInputOutputNewParams, opts ...option.RequestOption) (res *Output, err error) {
	var env LiveInputOutputNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", params.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the state of an output.
func (r *LiveInputOutputService) Update(ctx context.Context, liveInputIdentifier string, outputIdentifier string, params LiveInputOutputUpdateParams, opts ...option.RequestOption) (res *Output, err error) {
	var env LiveInputOutputUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	if outputIdentifier == "" {
		err = errors.New("missing required output_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", params.AccountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *LiveInputOutputService) List(ctx context.Context, liveInputIdentifier string, query LiveInputOutputListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Output], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", query.AccountID, liveInputIdentifier)
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

// Retrieves all outputs associated with a specified live input.
func (r *LiveInputOutputService) ListAutoPaging(ctx context.Context, liveInputIdentifier string, query LiveInputOutputListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Output] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, liveInputIdentifier, query, opts...))
}

// Deletes an output and removes it from the associated live input.
func (r *LiveInputOutputService) Delete(ctx context.Context, liveInputIdentifier string, outputIdentifier string, body LiveInputOutputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	if outputIdentifier == "" {
		err = errors.New("missing required output_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", body.AccountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Output struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled bool `json:"enabled"`
	// The streamKey used to authenticate against an output's target.
	StreamKey string `json:"streamKey"`
	// A unique identifier for the output.
	UID string `json:"uid"`
	// The URL an output uses to restream.
	URL  string     `json:"url"`
	JSON outputJSON `json:"-"`
}

// outputJSON contains the JSON metadata for the struct [Output]
type outputJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	UID         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Output) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The streamKey used to authenticate against an output's target.
	StreamKey param.Field[string] `json:"streamKey,required"`
	// The URL an output uses to restream.
	URL param.Field[string] `json:"url,required"`
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r LiveInputOutputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LiveInputOutputNewResponseEnvelope struct {
	Errors   []LiveInputOutputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputOutputNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputOutputNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Output                                    `json:"result"`
	JSON    liveInputOutputNewResponseEnvelopeJSON    `json:"-"`
}

// liveInputOutputNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputOutputNewResponseEnvelope]
type liveInputOutputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           LiveInputOutputNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputOutputNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputOutputNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputOutputNewResponseEnvelopeErrors]
type liveInputOutputNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    liveInputOutputNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputOutputNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [LiveInputOutputNewResponseEnvelopeErrorsSource]
type liveInputOutputNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           LiveInputOutputNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputOutputNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputOutputNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LiveInputOutputNewResponseEnvelopeMessages]
type liveInputOutputNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    liveInputOutputNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputOutputNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [LiveInputOutputNewResponseEnvelopeMessagesSource]
type liveInputOutputNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputOutputNewResponseEnvelopeSuccess bool

const (
	LiveInputOutputNewResponseEnvelopeSuccessTrue LiveInputOutputNewResponseEnvelopeSuccess = true
)

func (r LiveInputOutputNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputOutputNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputOutputUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r LiveInputOutputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LiveInputOutputUpdateResponseEnvelope struct {
	Errors   []LiveInputOutputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputOutputUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputOutputUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Output                                       `json:"result"`
	JSON    liveInputOutputUpdateResponseEnvelopeJSON    `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LiveInputOutputUpdateResponseEnvelope]
type liveInputOutputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           LiveInputOutputUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputOutputUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LiveInputOutputUpdateResponseEnvelopeErrors]
type liveInputOutputUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    liveInputOutputUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [LiveInputOutputUpdateResponseEnvelopeErrorsSource]
type liveInputOutputUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           LiveInputOutputUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputOutputUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LiveInputOutputUpdateResponseEnvelopeMessages]
type liveInputOutputUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputOutputUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    liveInputOutputUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputOutputUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [LiveInputOutputUpdateResponseEnvelopeMessagesSource]
type liveInputOutputUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputOutputUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputOutputUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputOutputUpdateResponseEnvelopeSuccess bool

const (
	LiveInputOutputUpdateResponseEnvelopeSuccessTrue LiveInputOutputUpdateResponseEnvelopeSuccess = true
)

func (r LiveInputOutputUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputOutputUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputOutputListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputOutputDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
