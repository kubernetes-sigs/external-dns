// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

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

// ModelSchemaService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelSchemaService] method instead.
type ModelSchemaService struct {
	Options []option.RequestOption
}

// NewModelSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelSchemaService(opts ...option.RequestOption) (r *ModelSchemaService) {
	r = &ModelSchemaService{}
	r.Options = opts
	return
}

// Get Model Schema
func (r *ModelSchemaService) Get(ctx context.Context, params ModelSchemaGetParams, opts ...option.RequestOption) (res *ModelSchemaGetResponse, err error) {
	var env ModelSchemaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/schema", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ModelSchemaGetResponse = interface{}

type ModelSchemaGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Model Name
	Model param.Field[string] `query:"model,required"`
}

// URLQuery serializes [ModelSchemaGetParams]'s query parameters as `url.Values`.
func (r ModelSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ModelSchemaGetResponseEnvelope struct {
	Result  ModelSchemaGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    modelSchemaGetResponseEnvelopeJSON `json:"-"`
}

// modelSchemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ModelSchemaGetResponseEnvelope]
type modelSchemaGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
