// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

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

// ExpressionTemplateFallthroughService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExpressionTemplateFallthroughService] method instead.
type ExpressionTemplateFallthroughService struct {
	Options []option.RequestOption
}

// NewExpressionTemplateFallthroughService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExpressionTemplateFallthroughService(opts ...option.RequestOption) (r *ExpressionTemplateFallthroughService) {
	r = &ExpressionTemplateFallthroughService{}
	r.Options = opts
	return
}

// Generate fallthrough WAF expression template from a set of API hosts
func (r *ExpressionTemplateFallthroughService) New(ctx context.Context, params ExpressionTemplateFallthroughNewParams, opts ...option.RequestOption) (res *ExpressionTemplateFallthroughNewResponse, err error) {
	var env ExpressionTemplateFallthroughNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/expression-template/fallthrough", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ExpressionTemplateFallthroughNewResponse struct {
	// WAF Expression for fallthrough
	Expression string `json:"expression,required"`
	// Title for the expression
	Title string                                       `json:"title,required"`
	JSON  expressionTemplateFallthroughNewResponseJSON `json:"-"`
}

// expressionTemplateFallthroughNewResponseJSON contains the JSON metadata for the
// struct [ExpressionTemplateFallthroughNewResponse]
type expressionTemplateFallthroughNewResponseJSON struct {
	Expression  apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExpressionTemplateFallthroughNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r expressionTemplateFallthroughNewResponseJSON) RawJSON() string {
	return r.raw
}

type ExpressionTemplateFallthroughNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// List of hosts to be targeted in the expression
	Hosts param.Field[[]string] `json:"hosts,required"`
}

func (r ExpressionTemplateFallthroughNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExpressionTemplateFallthroughNewResponseEnvelope struct {
	Errors   Message                                  `json:"errors,required"`
	Messages Message                                  `json:"messages,required"`
	Result   ExpressionTemplateFallthroughNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ExpressionTemplateFallthroughNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    expressionTemplateFallthroughNewResponseEnvelopeJSON    `json:"-"`
}

// expressionTemplateFallthroughNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ExpressionTemplateFallthroughNewResponseEnvelope]
type expressionTemplateFallthroughNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExpressionTemplateFallthroughNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r expressionTemplateFallthroughNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ExpressionTemplateFallthroughNewResponseEnvelopeSuccess bool

const (
	ExpressionTemplateFallthroughNewResponseEnvelopeSuccessTrue ExpressionTemplateFallthroughNewResponseEnvelopeSuccess = true
)

func (r ExpressionTemplateFallthroughNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ExpressionTemplateFallthroughNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
