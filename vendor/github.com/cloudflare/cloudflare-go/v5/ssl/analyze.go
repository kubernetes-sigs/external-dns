// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AnalyzeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyzeService] method instead.
type AnalyzeService struct {
	Options []option.RequestOption
}

// NewAnalyzeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAnalyzeService(opts ...option.RequestOption) (r *AnalyzeService) {
	r = &AnalyzeService{}
	r.Options = opts
	return
}

// Returns the set of hostnames, the signature algorithm, and the expiration date
// of the certificate.
func (r *AnalyzeService) New(ctx context.Context, params AnalyzeNewParams, opts ...option.RequestOption) (res *AnalyzeNewResponse, err error) {
	var env AnalyzeNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/analyze", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyzeNewResponse = interface{}

type AnalyzeNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[custom_hostnames.BundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r AnalyzeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AnalyzeNewResponseEnvelope struct {
	Errors   []AnalyzeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyzeNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyzeNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AnalyzeNewResponse                `json:"result"`
	JSON    analyzeNewResponseEnvelopeJSON    `json:"-"`
}

// analyzeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyzeNewResponseEnvelope]
type analyzeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyzeNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AnalyzeNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyzeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyzeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AnalyzeNewResponseEnvelopeErrors]
type analyzeNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyzeNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    analyzeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyzeNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AnalyzeNewResponseEnvelopeErrorsSource]
type analyzeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyzeNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AnalyzeNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyzeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyzeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AnalyzeNewResponseEnvelopeMessages]
type analyzeNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyzeNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    analyzeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyzeNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AnalyzeNewResponseEnvelopeMessagesSource]
type analyzeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyzeNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyzeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyzeNewResponseEnvelopeSuccess bool

const (
	AnalyzeNewResponseEnvelopeSuccessTrue AnalyzeNewResponseEnvelopeSuccess = true
)

func (r AnalyzeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyzeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
