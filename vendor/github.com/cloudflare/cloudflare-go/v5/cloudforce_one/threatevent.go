// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ThreatEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventService] method instead.
type ThreatEventService struct {
	Options          []option.RequestOption
	Attackers        *ThreatEventAttackerService
	Categories       *ThreatEventCategoryService
	Countries        *ThreatEventCountryService
	Crons            *ThreatEventCronService
	Datasets         *ThreatEventDatasetService
	IndicatorTypes   *ThreatEventIndicatorTypeService
	Raw              *ThreatEventRawService
	Relate           *ThreatEventRelateService
	Tags             *ThreatEventTagService
	EventTags        *ThreatEventEventTagService
	TargetIndustries *ThreatEventTargetIndustryService
	Insights         *ThreatEventInsightService
}

// NewThreatEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventService(opts ...option.RequestOption) (r *ThreatEventService) {
	r = &ThreatEventService{}
	r.Options = opts
	r.Attackers = NewThreatEventAttackerService(opts...)
	r.Categories = NewThreatEventCategoryService(opts...)
	r.Countries = NewThreatEventCountryService(opts...)
	r.Crons = NewThreatEventCronService(opts...)
	r.Datasets = NewThreatEventDatasetService(opts...)
	r.IndicatorTypes = NewThreatEventIndicatorTypeService(opts...)
	r.Raw = NewThreatEventRawService(opts...)
	r.Relate = NewThreatEventRelateService(opts...)
	r.Tags = NewThreatEventTagService(opts...)
	r.EventTags = NewThreatEventEventTagService(opts...)
	r.TargetIndustries = NewThreatEventTargetIndustryService(opts...)
	r.Insights = NewThreatEventInsightService(opts...)
	return
}

// To create a dataset, see the
// [`Create Dataset`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/create/)
// endpoint. When `datasetId` parameter is unspecified, it will be created in a
// default dataset named `Cloudforce One Threat Events`.
func (r *ThreatEventService) New(ctx context.Context, params ThreatEventNewParams, opts ...option.RequestOption) (res *ThreatEventNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathAccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create", params.PathAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// When `datasetId` is unspecified, events will be listed from the
// `Cloudforce One Threat Events` dataset. To list existing datasets (and their
// IDs), use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint). Also, must provide query parameters.
func (r *ThreatEventService) List(ctx context.Context, params ThreatEventListParams, opts ...option.RequestOption) (res *[]ThreatEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *ThreatEventService) Delete(ctx context.Context, eventID string, body ThreatEventDeleteParams, opts ...option.RequestOption) (res *ThreatEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", body.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *ThreatEventService) BulkNew(ctx context.Context, params ThreatEventBulkNewParams, opts ...option.RequestOption) (res *float64, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an event
func (r *ThreatEventService) Edit(ctx context.Context, eventID string, params ThreatEventEditParams, opts ...option.RequestOption) (res *ThreatEventEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", params.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Reads an event
func (r *ThreatEventService) Get(ctx context.Context, eventID string, query ThreatEventGetParams, opts ...option.RequestOption) (res *ThreatEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", query.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventNewResponse struct {
	ID              float64                    `json:"id,required"`
	AccountID       float64                    `json:"accountId,required"`
	Attacker        string                     `json:"attacker,required"`
	AttackerCountry string                     `json:"attackerCountry,required"`
	Category        string                     `json:"category,required"`
	Date            string                     `json:"date,required"`
	Event           string                     `json:"event,required"`
	Indicator       string                     `json:"indicator,required"`
	IndicatorType   string                     `json:"indicatorType,required"`
	IndicatorTypeID float64                    `json:"indicatorTypeId,required"`
	KillChain       float64                    `json:"killChain,required"`
	MitreAttack     []string                   `json:"mitreAttack,required"`
	NumReferenced   float64                    `json:"numReferenced,required"`
	NumReferences   float64                    `json:"numReferences,required"`
	RawID           string                     `json:"rawId,required"`
	Referenced      []string                   `json:"referenced,required"`
	ReferencedIDs   []float64                  `json:"referencedIds,required"`
	References      []string                   `json:"references,required"`
	ReferencesIDs   []float64                  `json:"referencesIds,required"`
	Tags            []string                   `json:"tags,required"`
	TargetCountry   string                     `json:"targetCountry,required"`
	TargetIndustry  string                     `json:"targetIndustry,required"`
	TLP             string                     `json:"tlp,required"`
	UUID            string                     `json:"uuid,required"`
	Insight         string                     `json:"insight"`
	ReleasabilityID string                     `json:"releasabilityId"`
	JSON            threatEventNewResponseJSON `json:"-"`
}

// threatEventNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventNewResponse]
type threatEventNewResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventListResponse struct {
	ID              float64                     `json:"id,required"`
	AccountID       float64                     `json:"accountId,required"`
	Attacker        string                      `json:"attacker,required"`
	AttackerCountry string                      `json:"attackerCountry,required"`
	Category        string                      `json:"category,required"`
	Date            string                      `json:"date,required"`
	Event           string                      `json:"event,required"`
	Indicator       string                      `json:"indicator,required"`
	IndicatorType   string                      `json:"indicatorType,required"`
	IndicatorTypeID float64                     `json:"indicatorTypeId,required"`
	KillChain       float64                     `json:"killChain,required"`
	MitreAttack     []string                    `json:"mitreAttack,required"`
	NumReferenced   float64                     `json:"numReferenced,required"`
	NumReferences   float64                     `json:"numReferences,required"`
	RawID           string                      `json:"rawId,required"`
	Referenced      []string                    `json:"referenced,required"`
	ReferencedIDs   []float64                   `json:"referencedIds,required"`
	References      []string                    `json:"references,required"`
	ReferencesIDs   []float64                   `json:"referencesIds,required"`
	Tags            []string                    `json:"tags,required"`
	TargetCountry   string                      `json:"targetCountry,required"`
	TargetIndustry  string                      `json:"targetIndustry,required"`
	TLP             string                      `json:"tlp,required"`
	UUID            string                      `json:"uuid,required"`
	Insight         string                      `json:"insight"`
	ReleasabilityID string                      `json:"releasabilityId"`
	JSON            threatEventListResponseJSON `json:"-"`
}

// threatEventListResponseJSON contains the JSON metadata for the struct
// [ThreatEventListResponse]
type threatEventListResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDeleteResponse struct {
	UUID string                        `json:"uuid,required"`
	JSON threatEventDeleteResponseJSON `json:"-"`
}

// threatEventDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventDeleteResponse]
type threatEventDeleteResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventEditResponse struct {
	ID              float64                     `json:"id,required"`
	AccountID       float64                     `json:"accountId,required"`
	Attacker        string                      `json:"attacker,required"`
	AttackerCountry string                      `json:"attackerCountry,required"`
	Category        string                      `json:"category,required"`
	Date            string                      `json:"date,required"`
	Event           string                      `json:"event,required"`
	Indicator       string                      `json:"indicator,required"`
	IndicatorType   string                      `json:"indicatorType,required"`
	IndicatorTypeID float64                     `json:"indicatorTypeId,required"`
	KillChain       float64                     `json:"killChain,required"`
	MitreAttack     []string                    `json:"mitreAttack,required"`
	NumReferenced   float64                     `json:"numReferenced,required"`
	NumReferences   float64                     `json:"numReferences,required"`
	RawID           string                      `json:"rawId,required"`
	Referenced      []string                    `json:"referenced,required"`
	ReferencedIDs   []float64                   `json:"referencedIds,required"`
	References      []string                    `json:"references,required"`
	ReferencesIDs   []float64                   `json:"referencesIds,required"`
	Tags            []string                    `json:"tags,required"`
	TargetCountry   string                      `json:"targetCountry,required"`
	TargetIndustry  string                      `json:"targetIndustry,required"`
	TLP             string                      `json:"tlp,required"`
	UUID            string                      `json:"uuid,required"`
	Insight         string                      `json:"insight"`
	ReleasabilityID string                      `json:"releasabilityId"`
	JSON            threatEventEditResponseJSON `json:"-"`
}

// threatEventEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventEditResponse]
type threatEventEditResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGetResponse struct {
	ID              float64                    `json:"id,required"`
	AccountID       float64                    `json:"accountId,required"`
	Attacker        string                     `json:"attacker,required"`
	AttackerCountry string                     `json:"attackerCountry,required"`
	Category        string                     `json:"category,required"`
	Date            string                     `json:"date,required"`
	Event           string                     `json:"event,required"`
	Indicator       string                     `json:"indicator,required"`
	IndicatorType   string                     `json:"indicatorType,required"`
	IndicatorTypeID float64                    `json:"indicatorTypeId,required"`
	KillChain       float64                    `json:"killChain,required"`
	MitreAttack     []string                   `json:"mitreAttack,required"`
	NumReferenced   float64                    `json:"numReferenced,required"`
	NumReferences   float64                    `json:"numReferences,required"`
	RawID           string                     `json:"rawId,required"`
	Referenced      []string                   `json:"referenced,required"`
	ReferencedIDs   []float64                  `json:"referencedIds,required"`
	References      []string                   `json:"references,required"`
	ReferencesIDs   []float64                  `json:"referencesIds,required"`
	Tags            []string                   `json:"tags,required"`
	TargetCountry   string                     `json:"targetCountry,required"`
	TargetIndustry  string                     `json:"targetIndustry,required"`
	TLP             string                     `json:"tlp,required"`
	UUID            string                     `json:"uuid,required"`
	Insight         string                     `json:"insight"`
	ReleasabilityID string                     `json:"releasabilityId"`
	JSON            threatEventGetResponseJSON `json:"-"`
}

// threatEventGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventGetResponse]
type threatEventGetResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventNewParams struct {
	// Account ID.
	PathAccountID   param.Field[string]                  `path:"account_id,required"`
	Category        param.Field[string]                  `json:"category,required"`
	Date            param.Field[time.Time]               `json:"date,required" format:"date-time"`
	Event           param.Field[string]                  `json:"event,required"`
	IndicatorType   param.Field[string]                  `json:"indicatorType,required"`
	Raw             param.Field[ThreatEventNewParamsRaw] `json:"raw,required"`
	TLP             param.Field[string]                  `json:"tlp,required"`
	BodyAccountID   param.Field[float64]                 `json:"accountId"`
	Attacker        param.Field[string]                  `json:"attacker"`
	AttackerCountry param.Field[string]                  `json:"attackerCountry"`
	DatasetID       param.Field[string]                  `json:"datasetId"`
	Indicator       param.Field[string]                  `json:"indicator"`
	Tags            param.Field[[]string]                `json:"tags"`
	TargetCountry   param.Field[string]                  `json:"targetCountry"`
	TargetIndustry  param.Field[string]                  `json:"targetIndustry"`
}

func (r ThreatEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventNewParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data,required"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventNewParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventListParams struct {
	// Account ID.
	AccountID    param.Field[string]                        `path:"account_id,required"`
	DatasetID    param.Field[[]string]                      `query:"datasetId"`
	ForceRefresh param.Field[bool]                          `query:"forceRefresh"`
	Order        param.Field[ThreatEventListParamsOrder]    `query:"order"`
	OrderBy      param.Field[string]                        `query:"orderBy"`
	Page         param.Field[float64]                       `query:"page"`
	PageSize     param.Field[float64]                       `query:"pageSize"`
	Search       param.Field[[]ThreatEventListParamsSearch] `query:"search"`
}

// URLQuery serializes [ThreatEventListParams]'s query parameters as `url.Values`.
func (r ThreatEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventListParamsOrder string

const (
	ThreatEventListParamsOrderAsc  ThreatEventListParamsOrder = "asc"
	ThreatEventListParamsOrderDesc ThreatEventListParamsOrder = "desc"
)

func (r ThreatEventListParamsOrder) IsKnown() bool {
	switch r {
	case ThreatEventListParamsOrderAsc, ThreatEventListParamsOrderDesc:
		return true
	}
	return false
}

type ThreatEventListParamsSearch struct {
	Field param.Field[string]                                `query:"field"`
	Op    param.Field[ThreatEventListParamsSearchOp]         `query:"op"`
	Value param.Field[ThreatEventListParamsSearchValueUnion] `query:"value"`
}

// URLQuery serializes [ThreatEventListParamsSearch]'s query parameters as
// `url.Values`.
func (r ThreatEventListParamsSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventListParamsSearchOp string

const (
	ThreatEventListParamsSearchOpEquals     ThreatEventListParamsSearchOp = "equals"
	ThreatEventListParamsSearchOpNot        ThreatEventListParamsSearchOp = "not"
	ThreatEventListParamsSearchOpGt         ThreatEventListParamsSearchOp = "gt"
	ThreatEventListParamsSearchOpGte        ThreatEventListParamsSearchOp = "gte"
	ThreatEventListParamsSearchOpLt         ThreatEventListParamsSearchOp = "lt"
	ThreatEventListParamsSearchOpLte        ThreatEventListParamsSearchOp = "lte"
	ThreatEventListParamsSearchOpLike       ThreatEventListParamsSearchOp = "like"
	ThreatEventListParamsSearchOpContains   ThreatEventListParamsSearchOp = "contains"
	ThreatEventListParamsSearchOpStartsWith ThreatEventListParamsSearchOp = "startsWith"
	ThreatEventListParamsSearchOpEndsWith   ThreatEventListParamsSearchOp = "endsWith"
	ThreatEventListParamsSearchOpIn         ThreatEventListParamsSearchOp = "in"
	ThreatEventListParamsSearchOpFind       ThreatEventListParamsSearchOp = "find"
)

func (r ThreatEventListParamsSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventListParamsSearchOpEquals, ThreatEventListParamsSearchOpNot, ThreatEventListParamsSearchOpGt, ThreatEventListParamsSearchOpGte, ThreatEventListParamsSearchOpLt, ThreatEventListParamsSearchOpLte, ThreatEventListParamsSearchOpLike, ThreatEventListParamsSearchOpContains, ThreatEventListParamsSearchOpStartsWith, ThreatEventListParamsSearchOpEndsWith, ThreatEventListParamsSearchOpIn, ThreatEventListParamsSearchOpFind:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [cloudforce_one.ThreatEventListParamsSearchValueArray].
type ThreatEventListParamsSearchValueUnion interface {
	ImplementsThreatEventListParamsSearchValueUnion()
}

type ThreatEventListParamsSearchValueArray []ThreatEventListParamsSearchValueArrayItemUnion

func (r ThreatEventListParamsSearchValueArray) ImplementsThreatEventListParamsSearchValueUnion() {}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ThreatEventListParamsSearchValueArrayItemUnion interface {
	ImplementsThreatEventListParamsSearchValueArrayItemUnion()
}

type ThreatEventDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventBulkNewParams struct {
	// Account ID.
	AccountID param.Field[string]                         `path:"account_id,required"`
	Data      param.Field[[]ThreatEventBulkNewParamsData] `json:"data,required"`
	DatasetID param.Field[string]                         `json:"datasetId,required"`
}

func (r ThreatEventBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventBulkNewParamsData struct {
	Category        param.Field[string]                          `json:"category,required"`
	Date            param.Field[time.Time]                       `json:"date,required" format:"date-time"`
	Event           param.Field[string]                          `json:"event,required"`
	IndicatorType   param.Field[string]                          `json:"indicatorType,required"`
	Raw             param.Field[ThreatEventBulkNewParamsDataRaw] `json:"raw,required"`
	TLP             param.Field[string]                          `json:"tlp,required"`
	AccountID       param.Field[float64]                         `json:"accountId"`
	Attacker        param.Field[string]                          `json:"attacker"`
	AttackerCountry param.Field[string]                          `json:"attackerCountry"`
	DatasetID       param.Field[string]                          `json:"datasetId"`
	Indicator       param.Field[string]                          `json:"indicator"`
	Tags            param.Field[[]string]                        `json:"tags"`
	TargetCountry   param.Field[string]                          `json:"targetCountry"`
	TargetIndustry  param.Field[string]                          `json:"targetIndustry"`
}

func (r ThreatEventBulkNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventBulkNewParamsDataRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data,required"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventBulkNewParamsDataRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventEditParams struct {
	// Account ID.
	AccountID       param.Field[string]                   `path:"account_id,required"`
	Attacker        param.Field[string]                   `json:"attacker"`
	AttackerCountry param.Field[string]                   `json:"attackerCountry"`
	Category        param.Field[string]                   `json:"category"`
	Date            param.Field[time.Time]                `json:"date" format:"date-time"`
	Event           param.Field[string]                   `json:"event"`
	Indicator       param.Field[string]                   `json:"indicator"`
	IndicatorType   param.Field[string]                   `json:"indicatorType"`
	Insight         param.Field[string]                   `json:"insight"`
	Raw             param.Field[ThreatEventEditParamsRaw] `json:"raw"`
	TargetCountry   param.Field[string]                   `json:"targetCountry"`
	TargetIndustry  param.Field[string]                   `json:"targetIndustry"`
	TLP             param.Field[string]                   `json:"tlp"`
}

func (r ThreatEventEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventEditParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventEditParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
