// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

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
)

// DiscoveryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscoveryService] method instead.
type DiscoveryService struct {
	Options    []option.RequestOption
	Operations *DiscoveryOperationService
}

// NewDiscoveryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiscoveryService(opts ...option.RequestOption) (r *DiscoveryService) {
	r = &DiscoveryService{}
	r.Options = opts
	r.Operations = NewDiscoveryOperationService(opts...)
	return
}

// Retrieve the most up to date view of discovered operations, rendered as OpenAPI
// schemas
func (r *DiscoveryService) Get(ctx context.Context, query DiscoveryGetParams, opts ...option.RequestOption) (res *DiscoveryGetResponse, err error) {
	var env DiscoveryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiscoveryOperation struct {
	// UUID.
	ID string `json:"id,required"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method DiscoveryOperationMethod `json:"method,required"`
	// API discovery engine(s) that discovered this operation
	Origin []DiscoveryOperationOrigin `json:"origin,required"`
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State    DiscoveryOperationState    `json:"state,required"`
	Features DiscoveryOperationFeatures `json:"features"`
	JSON     discoveryOperationJSON     `json:"-"`
}

// discoveryOperationJSON contains the JSON metadata for the struct
// [DiscoveryOperation]
type discoveryOperationJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	Origin      apijson.Field
	State       apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type DiscoveryOperationMethod string

const (
	DiscoveryOperationMethodGet     DiscoveryOperationMethod = "GET"
	DiscoveryOperationMethodPost    DiscoveryOperationMethod = "POST"
	DiscoveryOperationMethodHead    DiscoveryOperationMethod = "HEAD"
	DiscoveryOperationMethodOptions DiscoveryOperationMethod = "OPTIONS"
	DiscoveryOperationMethodPut     DiscoveryOperationMethod = "PUT"
	DiscoveryOperationMethodDelete  DiscoveryOperationMethod = "DELETE"
	DiscoveryOperationMethodConnect DiscoveryOperationMethod = "CONNECT"
	DiscoveryOperationMethodPatch   DiscoveryOperationMethod = "PATCH"
	DiscoveryOperationMethodTrace   DiscoveryOperationMethod = "TRACE"
)

func (r DiscoveryOperationMethod) IsKnown() bool {
	switch r {
	case DiscoveryOperationMethodGet, DiscoveryOperationMethodPost, DiscoveryOperationMethodHead, DiscoveryOperationMethodOptions, DiscoveryOperationMethodPut, DiscoveryOperationMethodDelete, DiscoveryOperationMethodConnect, DiscoveryOperationMethodPatch, DiscoveryOperationMethodTrace:
		return true
	}
	return false
}

//   - `ML` - Discovered operation was sourced using ML API Discovery _
//     `SessionIdentifier` - Discovered operation was sourced using Session
//     Identifier API Discovery _ `LabelDiscovery` - Discovered operation was
//     identified to have a specific label
type DiscoveryOperationOrigin string

const (
	DiscoveryOperationOriginMl                DiscoveryOperationOrigin = "ML"
	DiscoveryOperationOriginSessionIdentifier DiscoveryOperationOrigin = "SessionIdentifier"
	DiscoveryOperationOriginLabelDiscovery    DiscoveryOperationOrigin = "LabelDiscovery"
)

func (r DiscoveryOperationOrigin) IsKnown() bool {
	switch r {
	case DiscoveryOperationOriginMl, DiscoveryOperationOriginSessionIdentifier, DiscoveryOperationOriginLabelDiscovery:
		return true
	}
	return false
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type DiscoveryOperationState string

const (
	DiscoveryOperationStateReview  DiscoveryOperationState = "review"
	DiscoveryOperationStateSaved   DiscoveryOperationState = "saved"
	DiscoveryOperationStateIgnored DiscoveryOperationState = "ignored"
)

func (r DiscoveryOperationState) IsKnown() bool {
	switch r {
	case DiscoveryOperationStateReview, DiscoveryOperationStateSaved, DiscoveryOperationStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationFeatures struct {
	TrafficStats DiscoveryOperationFeaturesTrafficStats `json:"traffic_stats"`
	JSON         discoveryOperationFeaturesJSON         `json:"-"`
}

// discoveryOperationFeaturesJSON contains the JSON metadata for the struct
// [DiscoveryOperationFeatures]
type discoveryOperationFeaturesJSON struct {
	TrafficStats apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DiscoveryOperationFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationFeaturesJSON) RawJSON() string {
	return r.raw
}

type DiscoveryOperationFeaturesTrafficStats struct {
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The period in seconds these statistics were computed over
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The average number of requests seen during this period
	Requests float64                                    `json:"requests,required"`
	JSON     discoveryOperationFeaturesTrafficStatsJSON `json:"-"`
}

// discoveryOperationFeaturesTrafficStatsJSON contains the JSON metadata for the
// struct [DiscoveryOperationFeaturesTrafficStats]
type discoveryOperationFeaturesTrafficStatsJSON struct {
	LastUpdated   apijson.Field
	PeriodSeconds apijson.Field
	Requests      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DiscoveryOperationFeaturesTrafficStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationFeaturesTrafficStatsJSON) RawJSON() string {
	return r.raw
}

type DiscoveryGetResponse struct {
	Schemas   []interface{}            `json:"schemas,required"`
	Timestamp time.Time                `json:"timestamp,required" format:"date-time"`
	JSON      discoveryGetResponseJSON `json:"-"`
}

// discoveryGetResponseJSON contains the JSON metadata for the struct
// [DiscoveryGetResponse]
type discoveryGetResponseJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryGetResponseJSON) RawJSON() string {
	return r.raw
}

type DiscoveryGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DiscoveryGetResponseEnvelope struct {
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   DiscoveryGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DiscoveryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryGetResponseEnvelopeJSON    `json:"-"`
}

// discoveryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DiscoveryGetResponseEnvelope]
type discoveryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DiscoveryGetResponseEnvelopeSuccess bool

const (
	DiscoveryGetResponseEnvelopeSuccessTrue DiscoveryGetResponseEnvelopeSuccess = true
)

func (r DiscoveryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
