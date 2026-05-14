// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SchemaService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchemaService] method instead.
type SchemaService struct {
	Options []option.RequestOption
}

// NewSchemaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSchemaService(opts ...option.RequestOption) (r *SchemaService) {
	r = &SchemaService{}
	r.Options = opts
	return
}

// Retrieve operations and features as OpenAPI schemas
func (r *SchemaService) List(ctx context.Context, params SchemaListParams, opts ...option.RequestOption) (res *SchemaListResponse, err error) {
	var env SchemaListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/schemas", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SchemaListResponse struct {
	Schemas   []interface{}          `json:"schemas"`
	Timestamp string                 `json:"timestamp"`
	JSON      schemaListResponseJSON `json:"-"`
}

// schemaListResponseJSON contains the JSON metadata for the struct
// [SchemaListResponse]
type schemaListResponseJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaListResponseJSON) RawJSON() string {
	return r.raw
}

type SchemaListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]SchemaListParamsFeature] `query:"feature"`
	// Receive schema only for the given host(s).
	Host param.Field[[]string] `query:"host"`
}

// URLQuery serializes [SchemaListParams]'s query parameters as `url.Values`.
func (r SchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SchemaListParamsFeature string

const (
	SchemaListParamsFeatureThresholds       SchemaListParamsFeature = "thresholds"
	SchemaListParamsFeatureParameterSchemas SchemaListParamsFeature = "parameter_schemas"
	SchemaListParamsFeatureSchemaInfo       SchemaListParamsFeature = "schema_info"
)

func (r SchemaListParamsFeature) IsKnown() bool {
	switch r {
	case SchemaListParamsFeatureThresholds, SchemaListParamsFeatureParameterSchemas, SchemaListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

type SchemaListResponseEnvelope struct {
	Errors   Message            `json:"errors,required"`
	Messages Message            `json:"messages,required"`
	Result   SchemaListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SchemaListResponseEnvelopeSuccess `json:"success,required"`
	JSON    schemaListResponseEnvelopeJSON    `json:"-"`
}

// schemaListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SchemaListResponseEnvelope]
type schemaListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemaListResponseEnvelopeSuccess bool

const (
	SchemaListResponseEnvelopeSuccessTrue SchemaListResponseEnvelopeSuccess = true
)

func (r SchemaListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SchemaListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
