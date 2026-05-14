// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// RecommendationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecommendationService] method instead.
type RecommendationService struct {
	Options []option.RequestOption
}

// NewRecommendationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecommendationService(opts ...option.RequestOption) (r *RecommendationService) {
	r = &RecommendationService{}
	r.Options = opts
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
//
// Deprecated: SSL/TLS Recommender has been decommissioned in favor of Automatic
// SSL/TLS
func (r *RecommendationService) Get(ctx context.Context, query RecommendationGetParams, opts ...option.RequestOption) (res *RecommendationGetResponse, err error) {
	var env RecommendationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/recommendation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RecommendationGetResponse struct {
	ID string `json:"id,required"`
	// Whether this setting can be updated or not.
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Current setting of the automatic SSL/TLS.
	Value RecommendationGetResponseValue `json:"value,required"`
	// Next time this zone will be scanned by the Automatic SSL/TLS.
	NextScheduledScan time.Time                     `json:"next_scheduled_scan,nullable" format:"date-time"`
	JSON              recommendationGetResponseJSON `json:"-"`
}

// recommendationGetResponseJSON contains the JSON metadata for the struct
// [RecommendationGetResponse]
type recommendationGetResponseJSON struct {
	ID                apijson.Field
	Editable          apijson.Field
	ModifiedOn        apijson.Field
	Value             apijson.Field
	NextScheduledScan apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecommendationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current setting of the automatic SSL/TLS.
type RecommendationGetResponseValue string

const (
	RecommendationGetResponseValueAuto   RecommendationGetResponseValue = "auto"
	RecommendationGetResponseValueCustom RecommendationGetResponseValue = "custom"
)

func (r RecommendationGetResponseValue) IsKnown() bool {
	switch r {
	case RecommendationGetResponseValueAuto, RecommendationGetResponseValueCustom:
		return true
	}
	return false
}

type RecommendationGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RecommendationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   RecommendationGetResponse `json:"result,required"`
	// Indicates the API call's success or failure.
	Success bool                                  `json:"success,required"`
	JSON    recommendationGetResponseEnvelopeJSON `json:"-"`
}

// recommendationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecommendationGetResponseEnvelope]
type recommendationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecommendationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
