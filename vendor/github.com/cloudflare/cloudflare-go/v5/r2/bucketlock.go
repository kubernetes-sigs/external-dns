// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// BucketLockService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketLockService] method instead.
type BucketLockService struct {
	Options []option.RequestOption
}

// NewBucketLockService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketLockService(opts ...option.RequestOption) (r *BucketLockService) {
	r = &BucketLockService{}
	r.Options = opts
	return
}

// Set lock rules for a bucket.
func (r *BucketLockService) Update(ctx context.Context, bucketName string, params BucketLockUpdateParams, opts ...option.RequestOption) (res *BucketLockUpdateResponse, err error) {
	var env BucketLockUpdateResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lock", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get lock rules for a bucket.
func (r *BucketLockService) Get(ctx context.Context, bucketName string, params BucketLockGetParams, opts ...option.RequestOption) (res *BucketLockGetResponse, err error) {
	var env BucketLockGetResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lock", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketLockUpdateResponse = interface{}

type BucketLockGetResponse struct {
	Rules []BucketLockGetResponseRule `json:"rules"`
	JSON  bucketLockGetResponseJSON   `json:"-"`
}

// bucketLockGetResponseJSON contains the JSON metadata for the struct
// [BucketLockGetResponse]
type bucketLockGetResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketLockGetResponseRule struct {
	// Unique identifier for this rule.
	ID string `json:"id,required"`
	// Condition to apply a lock rule to an object for how long in seconds.
	Condition BucketLockGetResponseRulesCondition `json:"condition,required"`
	// Whether or not this rule is in effect.
	Enabled bool `json:"enabled,required"`
	// Rule will only apply to objects/uploads in the bucket that start with the given
	// prefix, an empty prefix can be provided to scope rule to all objects/uploads.
	Prefix string                        `json:"prefix"`
	JSON   bucketLockGetResponseRuleJSON `json:"-"`
}

// bucketLockGetResponseRuleJSON contains the JSON metadata for the struct
// [BucketLockGetResponseRule]
type bucketLockGetResponseRuleJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	Enabled     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

// Condition to apply a lock rule to an object for how long in seconds.
type BucketLockGetResponseRulesCondition struct {
	Type          BucketLockGetResponseRulesConditionType `json:"type,required"`
	Date          time.Time                               `json:"date" format:"date"`
	MaxAgeSeconds int64                                   `json:"maxAgeSeconds"`
	JSON          bucketLockGetResponseRulesConditionJSON `json:"-"`
	union         BucketLockGetResponseRulesConditionUnion
}

// bucketLockGetResponseRulesConditionJSON contains the JSON metadata for the
// struct [BucketLockGetResponseRulesCondition]
type bucketLockGetResponseRulesConditionJSON struct {
	Type          apijson.Field
	Date          apijson.Field
	MaxAgeSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r bucketLockGetResponseRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *BucketLockGetResponseRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = BucketLockGetResponseRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BucketLockGetResponseRulesConditionUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketLockGetResponseRulesConditionR2LockRuleAgeCondition],
// [BucketLockGetResponseRulesConditionR2LockRuleDateCondition],
// [BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition].
func (r BucketLockGetResponseRulesCondition) AsUnion() BucketLockGetResponseRulesConditionUnion {
	return r.union
}

// Condition to apply a lock rule to an object for how long in seconds.
//
// Union satisfied by [BucketLockGetResponseRulesConditionR2LockRuleAgeCondition],
// [BucketLockGetResponseRulesConditionR2LockRuleDateCondition] or
// [BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition].
type BucketLockGetResponseRulesConditionUnion interface {
	implementsBucketLockGetResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketLockGetResponseRulesConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLockGetResponseRulesConditionR2LockRuleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLockGetResponseRulesConditionR2LockRuleDateCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition{}),
		},
	)
}

// Condition to apply a lock rule to an object for how long in seconds.
type BucketLockGetResponseRulesConditionR2LockRuleAgeCondition struct {
	MaxAgeSeconds int64                                                         `json:"maxAgeSeconds,required"`
	Type          BucketLockGetResponseRulesConditionR2LockRuleAgeConditionType `json:"type,required"`
	JSON          bucketLockGetResponseRulesConditionR2LockRuleAgeConditionJSON `json:"-"`
}

// bucketLockGetResponseRulesConditionR2LockRuleAgeConditionJSON contains the JSON
// metadata for the struct
// [BucketLockGetResponseRulesConditionR2LockRuleAgeCondition]
type bucketLockGetResponseRulesConditionR2LockRuleAgeConditionJSON struct {
	MaxAgeSeconds apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BucketLockGetResponseRulesConditionR2LockRuleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseRulesConditionR2LockRuleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLockGetResponseRulesConditionR2LockRuleAgeCondition) implementsBucketLockGetResponseRulesCondition() {
}

type BucketLockGetResponseRulesConditionR2LockRuleAgeConditionType string

const (
	BucketLockGetResponseRulesConditionR2LockRuleAgeConditionTypeAge BucketLockGetResponseRulesConditionR2LockRuleAgeConditionType = "Age"
)

func (r BucketLockGetResponseRulesConditionR2LockRuleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLockGetResponseRulesConditionR2LockRuleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition to apply a lock rule to an object until a specific date.
type BucketLockGetResponseRulesConditionR2LockRuleDateCondition struct {
	Date time.Time                                                      `json:"date,required" format:"date"`
	Type BucketLockGetResponseRulesConditionR2LockRuleDateConditionType `json:"type,required"`
	JSON bucketLockGetResponseRulesConditionR2LockRuleDateConditionJSON `json:"-"`
}

// bucketLockGetResponseRulesConditionR2LockRuleDateConditionJSON contains the JSON
// metadata for the struct
// [BucketLockGetResponseRulesConditionR2LockRuleDateCondition]
type bucketLockGetResponseRulesConditionR2LockRuleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockGetResponseRulesConditionR2LockRuleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseRulesConditionR2LockRuleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLockGetResponseRulesConditionR2LockRuleDateCondition) implementsBucketLockGetResponseRulesCondition() {
}

type BucketLockGetResponseRulesConditionR2LockRuleDateConditionType string

const (
	BucketLockGetResponseRulesConditionR2LockRuleDateConditionTypeDate BucketLockGetResponseRulesConditionR2LockRuleDateConditionType = "Date"
)

func (r BucketLockGetResponseRulesConditionR2LockRuleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLockGetResponseRulesConditionR2LockRuleDateConditionTypeDate:
		return true
	}
	return false
}

// Condition to apply a lock rule indefinitely.
type BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition struct {
	Type BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionType `json:"type,required"`
	JSON bucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionJSON `json:"-"`
}

// bucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionJSON contains
// the JSON metadata for the struct
// [BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition]
type bucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLockGetResponseRulesConditionR2LockRuleIndefiniteCondition) implementsBucketLockGetResponseRulesCondition() {
}

type BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionType string

const (
	BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionTypeIndefinite BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionType = "Indefinite"
)

func (r BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionType) IsKnown() bool {
	switch r {
	case BucketLockGetResponseRulesConditionR2LockRuleIndefiniteConditionTypeIndefinite:
		return true
	}
	return false
}

type BucketLockGetResponseRulesConditionType string

const (
	BucketLockGetResponseRulesConditionTypeAge        BucketLockGetResponseRulesConditionType = "Age"
	BucketLockGetResponseRulesConditionTypeDate       BucketLockGetResponseRulesConditionType = "Date"
	BucketLockGetResponseRulesConditionTypeIndefinite BucketLockGetResponseRulesConditionType = "Indefinite"
)

func (r BucketLockGetResponseRulesConditionType) IsKnown() bool {
	switch r {
	case BucketLockGetResponseRulesConditionTypeAge, BucketLockGetResponseRulesConditionTypeDate, BucketLockGetResponseRulesConditionTypeIndefinite:
		return true
	}
	return false
}

type BucketLockUpdateParams struct {
	// Account ID.
	AccountID param.Field[string]                       `path:"account_id,required"`
	Rules     param.Field[[]BucketLockUpdateParamsRule] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketLockUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketLockUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketLockUpdateParamsRule struct {
	// Unique identifier for this rule.
	ID param.Field[string] `json:"id,required"`
	// Condition to apply a lock rule to an object for how long in seconds.
	Condition param.Field[BucketLockUpdateParamsRulesConditionUnion] `json:"condition,required"`
	// Whether or not this rule is in effect.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Rule will only apply to objects/uploads in the bucket that start with the given
	// prefix, an empty prefix can be provided to scope rule to all objects/uploads.
	Prefix param.Field[string] `json:"prefix"`
}

func (r BucketLockUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition to apply a lock rule to an object for how long in seconds.
type BucketLockUpdateParamsRulesCondition struct {
	Type          param.Field[BucketLockUpdateParamsRulesConditionType] `json:"type,required"`
	Date          param.Field[time.Time]                                `json:"date" format:"date"`
	MaxAgeSeconds param.Field[int64]                                    `json:"maxAgeSeconds"`
}

func (r BucketLockUpdateParamsRulesCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLockUpdateParamsRulesCondition) implementsBucketLockUpdateParamsRulesConditionUnion() {}

// Condition to apply a lock rule to an object for how long in seconds.
//
// Satisfied by [r2.BucketLockUpdateParamsRulesConditionR2LockRuleAgeCondition],
// [r2.BucketLockUpdateParamsRulesConditionR2LockRuleDateCondition],
// [r2.BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteCondition],
// [BucketLockUpdateParamsRulesCondition].
type BucketLockUpdateParamsRulesConditionUnion interface {
	implementsBucketLockUpdateParamsRulesConditionUnion()
}

// Condition to apply a lock rule to an object for how long in seconds.
type BucketLockUpdateParamsRulesConditionR2LockRuleAgeCondition struct {
	MaxAgeSeconds param.Field[int64]                                                          `json:"maxAgeSeconds,required"`
	Type          param.Field[BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionType] `json:"type,required"`
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleAgeCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleAgeCondition) implementsBucketLockUpdateParamsRulesConditionUnion() {
}

type BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionType string

const (
	BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionTypeAge BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionType = "Age"
)

func (r BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLockUpdateParamsRulesConditionR2LockRuleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition to apply a lock rule to an object until a specific date.
type BucketLockUpdateParamsRulesConditionR2LockRuleDateCondition struct {
	Date param.Field[time.Time]                                                       `json:"date,required" format:"date"`
	Type param.Field[BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionType] `json:"type,required"`
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleDateCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleDateCondition) implementsBucketLockUpdateParamsRulesConditionUnion() {
}

type BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionType string

const (
	BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionTypeDate BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionType = "Date"
)

func (r BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLockUpdateParamsRulesConditionR2LockRuleDateConditionTypeDate:
		return true
	}
	return false
}

// Condition to apply a lock rule indefinitely.
type BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteCondition struct {
	Type param.Field[BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionType] `json:"type,required"`
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteCondition) implementsBucketLockUpdateParamsRulesConditionUnion() {
}

type BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionType string

const (
	BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionTypeIndefinite BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionType = "Indefinite"
)

func (r BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionType) IsKnown() bool {
	switch r {
	case BucketLockUpdateParamsRulesConditionR2LockRuleIndefiniteConditionTypeIndefinite:
		return true
	}
	return false
}

type BucketLockUpdateParamsRulesConditionType string

const (
	BucketLockUpdateParamsRulesConditionTypeAge        BucketLockUpdateParamsRulesConditionType = "Age"
	BucketLockUpdateParamsRulesConditionTypeDate       BucketLockUpdateParamsRulesConditionType = "Date"
	BucketLockUpdateParamsRulesConditionTypeIndefinite BucketLockUpdateParamsRulesConditionType = "Indefinite"
)

func (r BucketLockUpdateParamsRulesConditionType) IsKnown() bool {
	switch r {
	case BucketLockUpdateParamsRulesConditionTypeAge, BucketLockUpdateParamsRulesConditionTypeDate, BucketLockUpdateParamsRulesConditionTypeIndefinite:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketLockUpdateParamsCfR2Jurisdiction string

const (
	BucketLockUpdateParamsCfR2JurisdictionDefault BucketLockUpdateParamsCfR2Jurisdiction = "default"
	BucketLockUpdateParamsCfR2JurisdictionEu      BucketLockUpdateParamsCfR2Jurisdiction = "eu"
	BucketLockUpdateParamsCfR2JurisdictionFedramp BucketLockUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketLockUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketLockUpdateParamsCfR2JurisdictionDefault, BucketLockUpdateParamsCfR2JurisdictionEu, BucketLockUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketLockUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []string                 `json:"messages,required"`
	Result   BucketLockUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketLockUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketLockUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketLockUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketLockUpdateResponseEnvelope]
type bucketLockUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketLockUpdateResponseEnvelopeSuccess bool

const (
	BucketLockUpdateResponseEnvelopeSuccessTrue BucketLockUpdateResponseEnvelopeSuccess = true
)

func (r BucketLockUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketLockUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketLockGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketLockGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketLockGetParamsCfR2Jurisdiction string

const (
	BucketLockGetParamsCfR2JurisdictionDefault BucketLockGetParamsCfR2Jurisdiction = "default"
	BucketLockGetParamsCfR2JurisdictionEu      BucketLockGetParamsCfR2Jurisdiction = "eu"
	BucketLockGetParamsCfR2JurisdictionFedramp BucketLockGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketLockGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketLockGetParamsCfR2JurisdictionDefault, BucketLockGetParamsCfR2JurisdictionEu, BucketLockGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketLockGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   BucketLockGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketLockGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketLockGetResponseEnvelopeJSON    `json:"-"`
}

// bucketLockGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketLockGetResponseEnvelope]
type bucketLockGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLockGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLockGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketLockGetResponseEnvelopeSuccess bool

const (
	BucketLockGetResponseEnvelopeSuccessTrue BucketLockGetResponseEnvelopeSuccess = true
)

func (r BucketLockGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketLockGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
