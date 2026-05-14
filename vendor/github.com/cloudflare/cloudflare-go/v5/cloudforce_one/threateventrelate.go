// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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

// ThreatEventRelateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventRelateService] method instead.
type ThreatEventRelateService struct {
	Options []option.RequestOption
}

// NewThreatEventRelateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventRelateService(opts ...option.RequestOption) (r *ThreatEventRelateService) {
	r = &ThreatEventRelateService{}
	r.Options = opts
	return
}

// Removes an event reference
func (r *ThreatEventRelateService) Delete(ctx context.Context, eventID string, body ThreatEventRelateDeleteParams, opts ...option.RequestOption) (res *ThreatEventRelateDeleteResponse, err error) {
	var env ThreatEventRelateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/relate/%s", body.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventRelateDeleteResponse struct {
	Success bool                                `json:"success,required"`
	JSON    threatEventRelateDeleteResponseJSON `json:"-"`
}

// threatEventRelateDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventRelateDeleteResponse]
type threatEventRelateDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRelateDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventRelateDeleteResponseEnvelope struct {
	Result  ThreatEventRelateDeleteResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    threatEventRelateDeleteResponseEnvelopeJSON `json:"-"`
}

// threatEventRelateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventRelateDeleteResponseEnvelope]
type threatEventRelateDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
