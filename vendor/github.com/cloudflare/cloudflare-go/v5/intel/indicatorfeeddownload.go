// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// IndicatorFeedDownloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedDownloadService] method instead.
type IndicatorFeedDownloadService struct {
	Options []option.RequestOption
}

// NewIndicatorFeedDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndicatorFeedDownloadService(opts ...option.RequestOption) (r *IndicatorFeedDownloadService) {
	r = &IndicatorFeedDownloadService{}
	r.Options = opts
	return
}

// Download indicator feed data
func (r *IndicatorFeedDownloadService) Get(ctx context.Context, feedID int64, query IndicatorFeedDownloadGetParams, opts ...option.RequestOption) (res *IndicatorFeedDownloadGetResponse, err error) {
	var env IndicatorFeedDownloadGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator_feeds/%v/download", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedDownloadGetResponse struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                               `json:"status"`
	JSON   indicatorFeedDownloadGetResponseJSON `json:"-"`
}

// indicatorFeedDownloadGetResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedDownloadGetResponse]
type indicatorFeedDownloadGetResponseJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedDownloadGetResponseEnvelope struct {
	Errors   []IndicatorFeedDownloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedDownloadGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedDownloadGetResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedDownloadGetResponse                `json:"result"`
	JSON    indicatorFeedDownloadGetResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndicatorFeedDownloadGetResponseEnvelope]
type indicatorFeedDownloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           IndicatorFeedDownloadGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedDownloadGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IndicatorFeedDownloadGetResponseEnvelopeErrors]
type indicatorFeedDownloadGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    indicatorFeedDownloadGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [IndicatorFeedDownloadGetResponseEnvelopeErrorsSource]
type indicatorFeedDownloadGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           IndicatorFeedDownloadGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedDownloadGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IndicatorFeedDownloadGetResponseEnvelopeMessages]
type indicatorFeedDownloadGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    indicatorFeedDownloadGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [IndicatorFeedDownloadGetResponseEnvelopeMessagesSource]
type indicatorFeedDownloadGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedDownloadGetResponseEnvelopeSuccess bool

const (
	IndicatorFeedDownloadGetResponseEnvelopeSuccessTrue IndicatorFeedDownloadGetResponseEnvelopeSuccess = true
)

func (r IndicatorFeedDownloadGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedDownloadGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
