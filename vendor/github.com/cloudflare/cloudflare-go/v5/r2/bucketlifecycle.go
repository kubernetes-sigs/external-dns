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

// BucketLifecycleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketLifecycleService] method instead.
type BucketLifecycleService struct {
	Options []option.RequestOption
}

// NewBucketLifecycleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketLifecycleService(opts ...option.RequestOption) (r *BucketLifecycleService) {
	r = &BucketLifecycleService{}
	r.Options = opts
	return
}

// Set the object lifecycle rules for a bucket.
func (r *BucketLifecycleService) Update(ctx context.Context, bucketName string, params BucketLifecycleUpdateParams, opts ...option.RequestOption) (res *BucketLifecycleUpdateResponse, err error) {
	var env BucketLifecycleUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lifecycle", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get object lifecycle rules for a bucket.
func (r *BucketLifecycleService) Get(ctx context.Context, bucketName string, params BucketLifecycleGetParams, opts ...option.RequestOption) (res *BucketLifecycleGetResponse, err error) {
	var env BucketLifecycleGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lifecycle", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketLifecycleUpdateResponse = interface{}

type BucketLifecycleGetResponse struct {
	Rules []BucketLifecycleGetResponseRule `json:"rules"`
	JSON  bucketLifecycleGetResponseJSON   `json:"-"`
}

// bucketLifecycleGetResponseJSON contains the JSON metadata for the struct
// [BucketLifecycleGetResponse]
type bucketLifecycleGetResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketLifecycleGetResponseRule struct {
	// Unique identifier for this rule.
	ID string `json:"id,required"`
	// Conditions that apply to all transitions of this rule.
	Conditions BucketLifecycleGetResponseRulesConditions `json:"conditions,required"`
	// Whether or not this rule is in effect.
	Enabled bool `json:"enabled,required"`
	// Transition to abort ongoing multipart uploads.
	AbortMultipartUploadsTransition BucketLifecycleGetResponseRulesAbortMultipartUploadsTransition `json:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	DeleteObjectsTransition BucketLifecycleGetResponseRulesDeleteObjectsTransition `json:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	StorageClassTransitions []BucketLifecycleGetResponseRulesStorageClassTransition `json:"storageClassTransitions"`
	JSON                    bucketLifecycleGetResponseRuleJSON                      `json:"-"`
}

// bucketLifecycleGetResponseRuleJSON contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRule]
type bucketLifecycleGetResponseRuleJSON struct {
	ID                              apijson.Field
	Conditions                      apijson.Field
	Enabled                         apijson.Field
	AbortMultipartUploadsTransition apijson.Field
	DeleteObjectsTransition         apijson.Field
	StorageClassTransitions         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

// Conditions that apply to all transitions of this rule.
type BucketLifecycleGetResponseRulesConditions struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the
	// given prefix, an empty prefix can be provided to scope rule to all
	// objects/uploads.
	Prefix string                                        `json:"prefix,required"`
	JSON   bucketLifecycleGetResponseRulesConditionsJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesConditionsJSON contains the JSON metadata for the
// struct [BucketLifecycleGetResponseRulesConditions]
type bucketLifecycleGetResponseRulesConditionsJSON struct {
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesConditionsJSON) RawJSON() string {
	return r.raw
}

// Transition to abort ongoing multipart uploads.
type BucketLifecycleGetResponseRulesAbortMultipartUploadsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionCondition `json:"condition"`
	JSON      bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionJSON      `json:"-"`
}

// bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionJSON contains the
// JSON metadata for the struct
// [BucketLifecycleGetResponseRulesAbortMultipartUploadsTransition]
type bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesAbortMultipartUploadsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionCondition struct {
	MaxAge int64                                                                       `json:"maxAge,required"`
	Type   BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionType `json:"type,required"`
	JSON   bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionJSON
// contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionCondition]
type bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionJSON) RawJSON() string {
	return r.raw
}

type BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionType string

const (
	BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionTypeAge BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionType = "Age"
)

func (r BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesAbortMultipartUploadsTransitionConditionTypeAge:
		return true
	}
	return false
}

// Transition to delete objects.
type BucketLifecycleGetResponseRulesDeleteObjectsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition `json:"condition"`
	JSON      bucketLifecycleGetResponseRulesDeleteObjectsTransitionJSON      `json:"-"`
}

// bucketLifecycleGetResponseRulesDeleteObjectsTransitionJSON contains the JSON
// metadata for the struct [BucketLifecycleGetResponseRulesDeleteObjectsTransition]
type bucketLifecycleGetResponseRulesDeleteObjectsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesDeleteObjectsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesDeleteObjectsTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition struct {
	Type   BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionType `json:"type,required"`
	Date   time.Time                                                           `json:"date" format:"date"`
	MaxAge int64                                                               `json:"maxAge"`
	JSON   bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionJSON `json:"-"`
	union  BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionUnion
}

// bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionJSON contains the
// JSON metadata for the struct
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition]
type bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionJSON) RawJSON() string {
	return r.raw
}

func (r *BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition) UnmarshalJSON(data []byte) (err error) {
	*r = BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition],
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition].
func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition) AsUnion() BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition]
// or
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition].
type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionUnion interface {
	implementsBucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition{}),
		},
	)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition struct {
	MaxAge int64                                                                                      `json:"maxAge,required"`
	Type   BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType `json:"type,required"`
	JSON   bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON
// contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition]
type bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition) implementsBucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition() {
}

type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType string

const (
	BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition struct {
	Date time.Time                                                                                   `json:"date,required" format:"date"`
	Type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType `json:"type,required"`
	JSON bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON
// contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition]
type bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition) implementsBucketLifecycleGetResponseRulesDeleteObjectsTransitionCondition() {
}

type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType string

const (
	BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionType string

const (
	BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionTypeAge  BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionType = "Age"
	BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionTypeDate BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionType = "Date"
)

func (r BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionTypeAge, BucketLifecycleGetResponseRulesDeleteObjectsTransitionConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleGetResponseRulesStorageClassTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition    BucketLifecycleGetResponseRulesStorageClassTransitionsCondition    `json:"condition,required"`
	StorageClass BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClass `json:"storageClass,required"`
	JSON         bucketLifecycleGetResponseRulesStorageClassTransitionJSON          `json:"-"`
}

// bucketLifecycleGetResponseRulesStorageClassTransitionJSON contains the JSON
// metadata for the struct [BucketLifecycleGetResponseRulesStorageClassTransition]
type bucketLifecycleGetResponseRulesStorageClassTransitionJSON struct {
	Condition    apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesStorageClassTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesStorageClassTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleGetResponseRulesStorageClassTransitionsCondition struct {
	Type   BucketLifecycleGetResponseRulesStorageClassTransitionsConditionType `json:"type,required"`
	Date   time.Time                                                           `json:"date" format:"date"`
	MaxAge int64                                                               `json:"maxAge"`
	JSON   bucketLifecycleGetResponseRulesStorageClassTransitionsConditionJSON `json:"-"`
	union  BucketLifecycleGetResponseRulesStorageClassTransitionsConditionUnion
}

// bucketLifecycleGetResponseRulesStorageClassTransitionsConditionJSON contains the
// JSON metadata for the struct
// [BucketLifecycleGetResponseRulesStorageClassTransitionsCondition]
type bucketLifecycleGetResponseRulesStorageClassTransitionsConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bucketLifecycleGetResponseRulesStorageClassTransitionsConditionJSON) RawJSON() string {
	return r.raw
}

func (r *BucketLifecycleGetResponseRulesStorageClassTransitionsCondition) UnmarshalJSON(data []byte) (err error) {
	*r = BucketLifecycleGetResponseRulesStorageClassTransitionsCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition],
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition].
func (r BucketLifecycleGetResponseRulesStorageClassTransitionsCondition) AsUnion() BucketLifecycleGetResponseRulesStorageClassTransitionsConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition]
// or
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition].
type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionUnion interface {
	implementsBucketLifecycleGetResponseRulesStorageClassTransitionsCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketLifecycleGetResponseRulesStorageClassTransitionsConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition{}),
		},
	)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition struct {
	MaxAge int64                                                                                      `json:"maxAge,required"`
	Type   BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType `json:"type,required"`
	JSON   bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionJSON
// contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition]
type bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeCondition) implementsBucketLifecycleGetResponseRulesStorageClassTransitionsCondition() {
}

type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType string

const (
	BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition struct {
	Date time.Time                                                                                   `json:"date,required" format:"date"`
	Type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionType `json:"type,required"`
	JSON bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionJSON `json:"-"`
}

// bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionJSON
// contains the JSON metadata for the struct
// [BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition]
type bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateCondition) implementsBucketLifecycleGetResponseRulesStorageClassTransitionsCondition() {
}

type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionType string

const (
	BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleGetResponseRulesStorageClassTransitionsConditionType string

const (
	BucketLifecycleGetResponseRulesStorageClassTransitionsConditionTypeAge  BucketLifecycleGetResponseRulesStorageClassTransitionsConditionType = "Age"
	BucketLifecycleGetResponseRulesStorageClassTransitionsConditionTypeDate BucketLifecycleGetResponseRulesStorageClassTransitionsConditionType = "Date"
)

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesStorageClassTransitionsConditionTypeAge, BucketLifecycleGetResponseRulesStorageClassTransitionsConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClass string

const (
	BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClassInfrequentAccess BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClass = "InfrequentAccess"
)

func (r BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClass) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseRulesStorageClassTransitionsStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketLifecycleUpdateParams struct {
	// Account ID.
	AccountID param.Field[string]                            `path:"account_id,required"`
	Rules     param.Field[[]BucketLifecycleUpdateParamsRule] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketLifecycleUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketLifecycleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketLifecycleUpdateParamsRule struct {
	// Unique identifier for this rule.
	ID param.Field[string] `json:"id,required"`
	// Conditions that apply to all transitions of this rule.
	Conditions param.Field[BucketLifecycleUpdateParamsRulesConditions] `json:"conditions,required"`
	// Whether or not this rule is in effect.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Transition to abort ongoing multipart uploads.
	AbortMultipartUploadsTransition param.Field[BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransition] `json:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	DeleteObjectsTransition param.Field[BucketLifecycleUpdateParamsRulesDeleteObjectsTransition] `json:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	StorageClassTransitions param.Field[[]BucketLifecycleUpdateParamsRulesStorageClassTransition] `json:"storageClassTransitions"`
}

func (r BucketLifecycleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Conditions that apply to all transitions of this rule.
type BucketLifecycleUpdateParamsRulesConditions struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the
	// given prefix, an empty prefix can be provided to scope rule to all
	// objects/uploads.
	Prefix param.Field[string] `json:"prefix,required"`
}

func (r BucketLifecycleUpdateParamsRulesConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transition to abort ongoing multipart uploads.
type BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition param.Field[BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionCondition] `json:"condition"`
}

func (r BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionCondition struct {
	MaxAge param.Field[int64]                                                                        `json:"maxAge,required"`
	Type   param.Field[BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionType] `json:"type,required"`
}

func (r BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionType string

const (
	BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionTypeAge BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionType = "Age"
)

func (r BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesAbortMultipartUploadsTransitionConditionTypeAge:
		return true
	}
	return false
}

// Transition to delete objects.
type BucketLifecycleUpdateParamsRulesDeleteObjectsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition param.Field[BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion] `json:"condition"`
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionCondition struct {
	Type   param.Field[BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionType] `json:"type,required"`
	Date   param.Field[time.Time]                                                            `json:"date" format:"date"`
	MaxAge param.Field[int64]                                                                `json:"maxAge"`
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionCondition) implementsBucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion() {
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Satisfied by
// [r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition],
// [r2.BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition],
// [BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionCondition].
type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion interface {
	implementsBucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion()
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition struct {
	MaxAge param.Field[int64]                                                                                       `json:"maxAge,required"`
	Type   param.Field[BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType] `json:"type,required"`
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeCondition) implementsBucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion() {
}

type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType string

const (
	BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition struct {
	Date param.Field[time.Time]                                                                                    `json:"date,required" format:"date"`
	Type param.Field[BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType] `json:"type,required"`
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateCondition) implementsBucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionUnion() {
}

type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType string

const (
	BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionType string

const (
	BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionTypeAge  BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionType = "Age"
	BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionTypeDate BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionType = "Date"
)

func (r BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionTypeAge, BucketLifecycleUpdateParamsRulesDeleteObjectsTransitionConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleUpdateParamsRulesStorageClassTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition    param.Field[BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion] `json:"condition,required"`
	StorageClass param.Field[BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClass]   `json:"storageClass,required"`
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleUpdateParamsRulesStorageClassTransitionsCondition struct {
	Type   param.Field[BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionType] `json:"type,required"`
	Date   param.Field[time.Time]                                                            `json:"date" format:"date"`
	MaxAge param.Field[int64]                                                                `json:"maxAge"`
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsCondition) implementsBucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion() {
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Satisfied by
// [r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeCondition],
// [r2.BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateCondition],
// [BucketLifecycleUpdateParamsRulesStorageClassTransitionsCondition].
type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion interface {
	implementsBucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion()
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeCondition struct {
	MaxAge param.Field[int64]                                                                                       `json:"maxAge,required"`
	Type   param.Field[BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType] `json:"type,required"`
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeCondition) implementsBucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion() {
}

type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType string

const (
	BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateCondition struct {
	Date param.Field[time.Time]                                                                                    `json:"date,required" format:"date"`
	Type param.Field[BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionType] `json:"type,required"`
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateCondition) implementsBucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionUnion() {
}

type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionType string

const (
	BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionType string

const (
	BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionTypeAge  BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionType = "Age"
	BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionTypeDate BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionType = "Date"
)

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionType) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionTypeAge, BucketLifecycleUpdateParamsRulesStorageClassTransitionsConditionTypeDate:
		return true
	}
	return false
}

type BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClass string

const (
	BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClassInfrequentAccess BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClass = "InfrequentAccess"
)

func (r BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClass) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsRulesStorageClassTransitionsStorageClassInfrequentAccess:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketLifecycleUpdateParamsCfR2Jurisdiction string

const (
	BucketLifecycleUpdateParamsCfR2JurisdictionDefault BucketLifecycleUpdateParamsCfR2Jurisdiction = "default"
	BucketLifecycleUpdateParamsCfR2JurisdictionEu      BucketLifecycleUpdateParamsCfR2Jurisdiction = "eu"
	BucketLifecycleUpdateParamsCfR2JurisdictionFedramp BucketLifecycleUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketLifecycleUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateParamsCfR2JurisdictionDefault, BucketLifecycleUpdateParamsCfR2JurisdictionEu, BucketLifecycleUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketLifecycleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   BucketLifecycleUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketLifecycleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketLifecycleUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketLifecycleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketLifecycleUpdateResponseEnvelope]
type bucketLifecycleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketLifecycleUpdateResponseEnvelopeSuccess bool

const (
	BucketLifecycleUpdateResponseEnvelopeSuccessTrue BucketLifecycleUpdateResponseEnvelopeSuccess = true
)

func (r BucketLifecycleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketLifecycleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketLifecycleGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketLifecycleGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketLifecycleGetParamsCfR2Jurisdiction string

const (
	BucketLifecycleGetParamsCfR2JurisdictionDefault BucketLifecycleGetParamsCfR2Jurisdiction = "default"
	BucketLifecycleGetParamsCfR2JurisdictionEu      BucketLifecycleGetParamsCfR2Jurisdiction = "eu"
	BucketLifecycleGetParamsCfR2JurisdictionFedramp BucketLifecycleGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketLifecycleGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketLifecycleGetParamsCfR2JurisdictionDefault, BucketLifecycleGetParamsCfR2JurisdictionEu, BucketLifecycleGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketLifecycleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []string                   `json:"messages,required"`
	Result   BucketLifecycleGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketLifecycleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketLifecycleGetResponseEnvelopeJSON    `json:"-"`
}

// bucketLifecycleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketLifecycleGetResponseEnvelope]
type bucketLifecycleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketLifecycleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketLifecycleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketLifecycleGetResponseEnvelopeSuccess bool

const (
	BucketLifecycleGetResponseEnvelopeSuccessTrue BucketLifecycleGetResponseEnvelopeSuccess = true
)

func (r BucketLifecycleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketLifecycleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
