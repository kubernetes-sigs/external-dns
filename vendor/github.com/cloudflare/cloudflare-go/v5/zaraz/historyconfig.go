// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// HistoryConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoryConfigService] method instead.
type HistoryConfigService struct {
	Options []option.RequestOption
}

// NewHistoryConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHistoryConfigService(opts ...option.RequestOption) (r *HistoryConfigService) {
	r = &HistoryConfigService{}
	r.Options = opts
	return
}

// Gets a history of published Zaraz configurations by ID(s) for a zone.
func (r *HistoryConfigService) Get(ctx context.Context, params HistoryConfigGetParams, opts ...option.RequestOption) (res *HistoryConfigGetResponse, err error) {
	var env HistoryConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history/configs", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HistoryConfigGetResponse map[string]HistoryConfigGetResponseItem

type HistoryConfigGetResponseItem struct {
	// ID of the configuration
	ID int64 `json:"id,required"`
	// Zaraz configuration
	Config Configuration `json:"config,required"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                           `json:"userId,required"`
	JSON   historyConfigGetResponseItemJSON `json:"-"`
}

// historyConfigGetResponseItemJSON contains the JSON metadata for the struct
// [HistoryConfigGetResponseItem]
type historyConfigGetResponseItemJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryConfigGetResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseItemJSON) RawJSON() string {
	return r.raw
}

type HistoryConfigGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Comma separated list of Zaraz configuration IDs
	IDs param.Field[[]int64] `query:"ids,required"`
}

// URLQuery serializes [HistoryConfigGetParams]'s query parameters as `url.Values`.
func (r HistoryConfigGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HistoryConfigGetResponseEnvelope struct {
	Errors   []HistoryConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HistoryConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Object where keys are numericc onfiguration IDs
	Result HistoryConfigGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                 `json:"success,required"`
	JSON    historyConfigGetResponseEnvelopeJSON `json:"-"`
}

// historyConfigGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HistoryConfigGetResponseEnvelope]
type historyConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HistoryConfigGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           HistoryConfigGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             historyConfigGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// historyConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HistoryConfigGetResponseEnvelopeErrors]
type historyConfigGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HistoryConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HistoryConfigGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    historyConfigGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// historyConfigGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [HistoryConfigGetResponseEnvelopeErrorsSource]
type historyConfigGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryConfigGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HistoryConfigGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           HistoryConfigGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             historyConfigGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// historyConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HistoryConfigGetResponseEnvelopeMessages]
type historyConfigGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HistoryConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HistoryConfigGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    historyConfigGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// historyConfigGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [HistoryConfigGetResponseEnvelopeMessagesSource]
type historyConfigGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryConfigGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyConfigGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
