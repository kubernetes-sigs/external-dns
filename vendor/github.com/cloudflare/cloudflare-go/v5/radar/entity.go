// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EntityService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options   []option.RequestOption
	ASNs      *EntityASNService
	Locations *EntityLocationService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.ASNs = NewEntityASNService(opts...)
	r.Locations = NewEntityLocationService(opts...)
	return
}

// Retrieves IP address information.
func (r *EntityService) Get(ctx context.Context, query EntityGetParams, opts ...option.RequestOption) (res *EntityGetResponse, err error) {
	var env EntityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/entities/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EntityGetResponse struct {
	IP   EntityGetResponseIP   `json:"ip,required"`
	JSON entityGetResponseJSON `json:"-"`
}

// entityGetResponseJSON contains the JSON metadata for the struct
// [EntityGetResponse]
type entityGetResponseJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGetResponseJSON) RawJSON() string {
	return r.raw
}

type EntityGetResponseIP struct {
	ASN          string                  `json:"asn,required"`
	ASNLocation  string                  `json:"asnLocation,required"`
	ASNName      string                  `json:"asnName,required"`
	ASNOrgName   string                  `json:"asnOrgName,required"`
	IP           string                  `json:"ip,required"`
	IPVersion    string                  `json:"ipVersion,required"`
	Location     string                  `json:"location,required"`
	LocationName string                  `json:"locationName,required"`
	JSON         entityGetResponseIPJSON `json:"-"`
}

// entityGetResponseIPJSON contains the JSON metadata for the struct
// [EntityGetResponseIP]
type entityGetResponseIPJSON struct {
	ASN          apijson.Field
	ASNLocation  apijson.Field
	ASNName      apijson.Field
	ASNOrgName   apijson.Field
	IP           apijson.Field
	IPVersion    apijson.Field
	Location     apijson.Field
	LocationName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGetResponseIPJSON) RawJSON() string {
	return r.raw
}

type EntityGetParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required" format:"ip"`
	// Format in which results will be returned.
	Format param.Field[EntityGetParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityGetParams]'s query parameters as `url.Values`.
func (r EntityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityGetParamsFormat string

const (
	EntityGetParamsFormatJson EntityGetParamsFormat = "JSON"
	EntityGetParamsFormatCsv  EntityGetParamsFormat = "CSV"
)

func (r EntityGetParamsFormat) IsKnown() bool {
	switch r {
	case EntityGetParamsFormatJson, EntityGetParamsFormatCsv:
		return true
	}
	return false
}

type EntityGetResponseEnvelope struct {
	Result  EntityGetResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    entityGetResponseEnvelopeJSON `json:"-"`
}

// entityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityGetResponseEnvelope]
type entityGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
