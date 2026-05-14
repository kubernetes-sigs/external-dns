// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// V1StatService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1StatService] method instead.
type V1StatService struct {
	Options []option.RequestOption
}

// NewV1StatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1StatService(opts ...option.RequestOption) (r *V1StatService) {
	r = &V1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *V1StatService) Get(ctx context.Context, query V1StatGetParams, opts ...option.RequestOption) (res *Stat, err error) {
	var env V1StatGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/stats", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Stat struct {
	Count StatCount `json:"count"`
	JSON  statJSON  `json:"-"`
}

// statJSON contains the JSON metadata for the struct [Stat]
type statJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statJSON) RawJSON() string {
	return r.raw
}

type StatCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64       `json:"current"`
	JSON    statCountJSON `json:"-"`
}

// statCountJSON contains the JSON metadata for the struct [StatCount]
type statCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statCountJSON) RawJSON() string {
	return r.raw
}

type V1StatGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1StatGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Stat                  `json:"result,required"`
	// Whether the API call was successful
	Success V1StatGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1StatGetResponseEnvelopeJSON    `json:"-"`
}

// v1StatGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1StatGetResponseEnvelope]
type v1StatGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1StatGetResponseEnvelopeSuccess bool

const (
	V1StatGetResponseEnvelopeSuccessTrue V1StatGetResponseEnvelopeSuccess = true
)

func (r V1StatGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1StatGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
