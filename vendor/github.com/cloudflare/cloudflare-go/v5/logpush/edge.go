// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

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

// EdgeService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEdgeService] method instead.
type EdgeService struct {
	Options []option.RequestOption
}

// NewEdgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEdgeService(opts ...option.RequestOption) (r *EdgeService) {
	r = &EdgeService{}
	r.Options = opts
	return
}

// Creates a new Instant Logs job for a zone.
func (r *EdgeService) New(ctx context.Context, params EdgeNewParams, opts ...option.RequestOption) (res *InstantLogpushJob, err error) {
	var env EdgeNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/edge", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Instant Logs jobs for a zone.
func (r *EdgeService) Get(ctx context.Context, query EdgeGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[InstantLogpushJob], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/edge", query.ZoneID)
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

// Lists Instant Logs jobs for a zone.
func (r *EdgeService) GetAutoPaging(ctx context.Context, query EdgeGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InstantLogpushJob] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type InstantLogpushJob struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string                `json:"session_id"`
	JSON      instantLogpushJobJSON `json:"-"`
}

// instantLogpushJobJSON contains the JSON metadata for the struct
// [InstantLogpushJob]
type instantLogpushJobJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstantLogpushJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instantLogpushJobJSON) RawJSON() string {
	return r.raw
}

type EdgeNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Comma-separated list of fields.
	Fields param.Field[string] `json:"fields"`
	// Filters to drill down into specific events.
	Filter param.Field[string] `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample param.Field[int64] `json:"sample"`
}

func (r EdgeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EdgeNewResponseEnvelope struct {
	Errors   []EdgeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EdgeNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EdgeNewResponseEnvelopeSuccess `json:"success,required"`
	Result  InstantLogpushJob              `json:"result,nullable"`
	JSON    edgeNewResponseEnvelopeJSON    `json:"-"`
}

// edgeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelope]
type edgeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeErrors struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           EdgeNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             edgeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// edgeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelopeErrors]
type edgeNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    edgeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// edgeNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [EdgeNewResponseEnvelopeErrorsSource]
type edgeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeMessages struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           EdgeNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             edgeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// edgeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [EdgeNewResponseEnvelopeMessages]
type edgeNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EdgeNewResponseEnvelopeMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    edgeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// edgeNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [EdgeNewResponseEnvelopeMessagesSource]
type edgeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EdgeNewResponseEnvelopeSuccess bool

const (
	EdgeNewResponseEnvelopeSuccessTrue EdgeNewResponseEnvelopeSuccess = true
)

func (r EdgeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EdgeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EdgeGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
