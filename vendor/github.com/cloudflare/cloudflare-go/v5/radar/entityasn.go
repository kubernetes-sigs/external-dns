// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EntityASNService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityASNService] method instead.
type EntityASNService struct {
	Options []option.RequestOption
}

// NewEntityASNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntityASNService(opts ...option.RequestOption) (r *EntityASNService) {
	r = &EntityASNService{}
	r.Options = opts
	return
}

// Retrieves a list of autonomous systems.
func (r *EntityASNService) List(ctx context.Context, query EntityASNListParams, opts ...option.RequestOption) (res *EntityASNListResponse, err error) {
	var env EntityASNListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested autonomous system information. (A confidence level below
// `5` indicates a low level of confidence in the traffic data - normally this
// happens because Cloudflare has a small amount of traffic from/to this AS).
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *EntityASNService) Get(ctx context.Context, asn int64, query EntityASNGetParams, opts ...option.RequestOption) (res *EntityASNGetResponse, err error) {
	var env EntityASNGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested autonomous system information based on IP address.
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *EntityASNService) IP(ctx context.Context, query EntityASNIPParams, opts ...option.RequestOption) (res *EntityAsnipResponse, err error) {
	var env EntityAsnipResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves AS-level relationship for given networks.
func (r *EntityASNService) Rel(ctx context.Context, asn int64, query EntityASNRelParams, opts ...option.RequestOption) (res *EntityASNRelResponse, err error) {
	var env EntityASNRelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EntityASNListResponse struct {
	ASNs []EntityASNListResponseASN `json:"asns,required"`
	JSON entityASNListResponseJSON  `json:"-"`
}

// entityASNListResponseJSON contains the JSON metadata for the struct
// [EntityASNListResponse]
type entityASNListResponseJSON struct {
	ASNs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNListResponseASN struct {
	ASN         int64                        `json:"asn,required"`
	Country     string                       `json:"country,required"`
	CountryName string                       `json:"countryName,required"`
	Name        string                       `json:"name,required"`
	Aka         string                       `json:"aka"`
	OrgName     string                       `json:"orgName"`
	Website     string                       `json:"website"`
	JSON        entityASNListResponseASNJSON `json:"-"`
}

// entityASNListResponseASNJSON contains the JSON metadata for the struct
// [EntityASNListResponseASN]
type entityASNListResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Name        apijson.Field
	Aka         apijson.Field
	OrgName     apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponse struct {
	ASN  EntityASNGetResponseASN  `json:"asn,required"`
	JSON entityASNGetResponseJSON `json:"-"`
}

// entityASNGetResponseJSON contains the JSON metadata for the struct
// [EntityASNGetResponse]
type entityASNGetResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASN struct {
	ASN             int64                                 `json:"asn,required"`
	ConfidenceLevel int64                                 `json:"confidenceLevel,required"`
	Country         string                                `json:"country,required"`
	CountryName     string                                `json:"countryName,required"`
	EstimatedUsers  EntityASNGetResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                `json:"name,required"`
	OrgName         string                                `json:"orgName,required"`
	Related         []EntityASNGetResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry.
	Source  string                      `json:"source,required"`
	Website string                      `json:"website,required"`
	Aka     string                      `json:"aka"`
	JSON    entityASNGetResponseASNJSON `json:"-"`
}

// entityASNGetResponseASNJSON contains the JSON metadata for the struct
// [EntityASNGetResponseASN]
type entityASNGetResponseASNJSON struct {
	ASN             apijson.Field
	ConfidenceLevel apijson.Field
	Country         apijson.Field
	CountryName     apijson.Field
	EstimatedUsers  apijson.Field
	Name            apijson.Field
	OrgName         apijson.Field
	Related         apijson.Field
	Source          apijson.Field
	Website         apijson.Field
	Aka             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EntityASNGetResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNEstimatedUsers struct {
	Locations []EntityASNGetResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users.
	EstimatedUsers int64                                     `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNEstimatedUsersJSON `json:"-"`
}

// entityASNGetResponseASNEstimatedUsersJSON contains the JSON metadata for the
// struct [EntityASNGetResponseASNEstimatedUsers]
type entityASNGetResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location.
	EstimatedUsers int64                                             `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// entityASNGetResponseASNEstimatedUsersLocationJSON contains the JSON metadata for
// the struct [EntityASNGetResponseASNEstimatedUsersLocation]
type entityASNGetResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users.
	EstimatedUsers int64                              `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNRelatedJSON `json:"-"`
}

// entityASNGetResponseASNRelatedJSON contains the JSON metadata for the struct
// [EntityASNGetResponseASNRelated]
type entityASNGetResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponse struct {
	ASN  EntityAsnipResponseASN  `json:"asn,required"`
	JSON entityAsnipResponseJSON `json:"-"`
}

// entityAsnipResponseJSON contains the JSON metadata for the struct
// [EntityAsnipResponse]
type entityAsnipResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityAsnipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASN struct {
	ASN            int64                                `json:"asn,required"`
	Country        string                               `json:"country,required"`
	CountryName    string                               `json:"countryName,required"`
	EstimatedUsers EntityAsnipResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                               `json:"name,required"`
	OrgName        string                               `json:"orgName,required"`
	Related        []EntityAsnipResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry.
	Source  string                     `json:"source,required"`
	Website string                     `json:"website,required"`
	Aka     string                     `json:"aka"`
	JSON    entityAsnipResponseASNJSON `json:"-"`
}

// entityAsnipResponseASNJSON contains the JSON metadata for the struct
// [EntityAsnipResponseASN]
type entityAsnipResponseASNJSON struct {
	ASN            apijson.Field
	Country        apijson.Field
	CountryName    apijson.Field
	EstimatedUsers apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	Related        apijson.Field
	Source         apijson.Field
	Website        apijson.Field
	Aka            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNEstimatedUsers struct {
	Locations []EntityAsnipResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users.
	EstimatedUsers int64                                    `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNEstimatedUsersJSON `json:"-"`
}

// entityAsnipResponseASNEstimatedUsersJSON contains the JSON metadata for the
// struct [EntityAsnipResponseASNEstimatedUsers]
type entityAsnipResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location.
	EstimatedUsers int64                                            `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// entityAsnipResponseASNEstimatedUsersLocationJSON contains the JSON metadata for
// the struct [EntityAsnipResponseASNEstimatedUsersLocation]
type entityAsnipResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users.
	EstimatedUsers int64                             `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNRelatedJSON `json:"-"`
}

// entityAsnipResponseASNRelatedJSON contains the JSON metadata for the struct
// [EntityAsnipResponseASNRelated]
type entityAsnipResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponse struct {
	Meta EntityASNRelResponseMeta  `json:"meta,required"`
	Rels []EntityASNRelResponseRel `json:"rels,required"`
	JSON entityASNRelResponseJSON  `json:"-"`
}

// entityASNRelResponseJSON contains the JSON metadata for the struct
// [EntityASNRelResponse]
type entityASNRelResponseJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponseMeta struct {
	DataTime   string                       `json:"data_time,required"`
	QueryTime  string                       `json:"query_time,required"`
	TotalPeers int64                        `json:"total_peers,required"`
	JSON       entityASNRelResponseMetaJSON `json:"-"`
}

// entityASNRelResponseMetaJSON contains the JSON metadata for the struct
// [EntityASNRelResponseMeta]
type entityASNRelResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponseRel struct {
	Asn1        int64                       `json:"asn1,required"`
	Asn1Country string                      `json:"asn1_country,required"`
	Asn1Name    string                      `json:"asn1_name,required"`
	Asn2        int64                       `json:"asn2,required"`
	Asn2Country string                      `json:"asn2_country,required"`
	Asn2Name    string                      `json:"asn2_name,required"`
	Rel         string                      `json:"rel,required"`
	JSON        entityASNRelResponseRelJSON `json:"-"`
}

// entityASNRelResponseRelJSON contains the JSON metadata for the struct
// [EntityASNRelResponseRel]
type entityASNRelResponseRelJSON struct {
	Asn1        apijson.Field
	Asn1Country apijson.Field
	Asn1Name    apijson.Field
	Asn2        apijson.Field
	Asn2Country apijson.Field
	Asn2Name    apijson.Field
	Rel         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseRelJSON) RawJSON() string {
	return r.raw
}

type EntityASNListParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list.
	ASN param.Field[string] `query:"asn"`
	// Format in which results will be returned.
	Format param.Field[EntityASNListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Specifies the metric to order the ASNs by.
	OrderBy param.Field[EntityASNListParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [EntityASNListParams]'s query parameters as `url.Values`.
func (r EntityASNListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNListParamsFormat string

const (
	EntityASNListParamsFormatJson EntityASNListParamsFormat = "JSON"
	EntityASNListParamsFormatCsv  EntityASNListParamsFormat = "CSV"
)

func (r EntityASNListParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNListParamsFormatJson, EntityASNListParamsFormatCsv:
		return true
	}
	return false
}

// Specifies the metric to order the ASNs by.
type EntityASNListParamsOrderBy string

const (
	EntityASNListParamsOrderByASN        EntityASNListParamsOrderBy = "ASN"
	EntityASNListParamsOrderByPopulation EntityASNListParamsOrderBy = "POPULATION"
)

func (r EntityASNListParamsOrderBy) IsKnown() bool {
	switch r {
	case EntityASNListParamsOrderByASN, EntityASNListParamsOrderByPopulation:
		return true
	}
	return false
}

type EntityASNListResponseEnvelope struct {
	Result  EntityASNListResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    entityASNListResponseEnvelopeJSON `json:"-"`
}

// entityASNListResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNListResponseEnvelope]
type entityASNListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetParams struct {
	// Format in which results will be returned.
	Format param.Field[EntityASNGetParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNGetParams]'s query parameters as `url.Values`.
func (r EntityASNGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNGetParamsFormat string

const (
	EntityASNGetParamsFormatJson EntityASNGetParamsFormat = "JSON"
	EntityASNGetParamsFormatCsv  EntityASNGetParamsFormat = "CSV"
)

func (r EntityASNGetParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNGetParamsFormatJson, EntityASNGetParamsFormatCsv:
		return true
	}
	return false
}

type EntityASNGetResponseEnvelope struct {
	Result  EntityASNGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    entityASNGetResponseEnvelopeJSON `json:"-"`
}

// entityASNGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNGetResponseEnvelope]
type entityASNGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNIPParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required" format:"ip"`
	// Format in which results will be returned.
	Format param.Field[EntityAsnipParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNIPParams]'s query parameters as `url.Values`.
func (r EntityASNIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityAsnipParamsFormat string

const (
	EntityAsnipParamsFormatJson EntityAsnipParamsFormat = "JSON"
	EntityAsnipParamsFormatCsv  EntityAsnipParamsFormat = "CSV"
)

func (r EntityAsnipParamsFormat) IsKnown() bool {
	switch r {
	case EntityAsnipParamsFormatJson, EntityAsnipParamsFormatCsv:
		return true
	}
	return false
}

type EntityAsnipResponseEnvelope struct {
	Result  EntityAsnipResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    entityAsnipResponseEnvelopeJSON `json:"-"`
}

// entityAsnipResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityAsnipResponseEnvelope]
type entityAsnipResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityAsnipResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelParams struct {
	// Retrieves the AS relationship of ASN2 with respect to the given ASN.
	Asn2 param.Field[int64] `query:"asn2"`
	// Format in which results will be returned.
	Format param.Field[EntityASNRelParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNRelParams]'s query parameters as `url.Values`.
func (r EntityASNRelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNRelParamsFormat string

const (
	EntityASNRelParamsFormatJson EntityASNRelParamsFormat = "JSON"
	EntityASNRelParamsFormatCsv  EntityASNRelParamsFormat = "CSV"
)

func (r EntityASNRelParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNRelParamsFormatJson, EntityASNRelParamsFormatCsv:
		return true
	}
	return false
}

type EntityASNRelResponseEnvelope struct {
	Result  EntityASNRelResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    entityASNRelResponseEnvelopeJSON `json:"-"`
}

// entityASNRelResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNRelResponseEnvelope]
type entityASNRelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
