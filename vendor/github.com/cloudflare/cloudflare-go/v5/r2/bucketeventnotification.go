// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

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

// BucketEventNotificationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketEventNotificationService] method instead.
type BucketEventNotificationService struct {
	Options []option.RequestOption
}

// NewBucketEventNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketEventNotificationService(opts ...option.RequestOption) (r *BucketEventNotificationService) {
	r = &BucketEventNotificationService{}
	r.Options = opts
	return
}

// Create event notification rule.
func (r *BucketEventNotificationService) Update(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationUpdateParams, opts ...option.RequestOption) (res *BucketEventNotificationUpdateResponse, err error) {
	var env BucketEventNotificationUpdateResponseEnvelope
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all event notification rules for a bucket.
func (r *BucketEventNotificationService) List(ctx context.Context, bucketName string, params BucketEventNotificationListParams, opts ...option.RequestOption) (res *BucketEventNotificationListResponse, err error) {
	var env BucketEventNotificationListResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an event notification rule. **If no body is provided, all rules for
// specified queue will be deleted**.
func (r *BucketEventNotificationService) Delete(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationDeleteParams, opts ...option.RequestOption) (res *BucketEventNotificationDeleteResponse, err error) {
	var env BucketEventNotificationDeleteResponseEnvelope
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single event notification rule.
func (r *BucketEventNotificationService) Get(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationGetParams, opts ...option.RequestOption) (res *BucketEventNotificationGetResponse, err error) {
	var env BucketEventNotificationGetResponseEnvelope
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketEventNotificationUpdateResponse = interface{}

type BucketEventNotificationListResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []BucketEventNotificationListResponseQueue `json:"queues"`
	JSON   bucketEventNotificationListResponseJSON    `json:"-"`
}

// bucketEventNotificationListResponseJSON contains the JSON metadata for the
// struct [BucketEventNotificationListResponse]
type bucketEventNotificationListResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueue struct {
	// Queue ID.
	QueueID string `json:"queueId"`
	// Name of the queue.
	QueueName string                                          `json:"queueName"`
	Rules     []BucketEventNotificationListResponseQueuesRule `json:"rules"`
	JSON      bucketEventNotificationListResponseQueueJSON    `json:"-"`
}

// bucketEventNotificationListResponseQueueJSON contains the JSON metadata for the
// struct [BucketEventNotificationListResponseQueue]
type bucketEventNotificationListResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseQueueJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions []BucketEventNotificationListResponseQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created.
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description string `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix string `json:"prefix"`
	// Rule ID.
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix.
	Suffix string                                            `json:"suffix"`
	JSON   bucketEventNotificationListResponseQueuesRuleJSON `json:"-"`
}

// bucketEventNotificationListResponseQueuesRuleJSON contains the JSON metadata for
// the struct [BucketEventNotificationListResponseQueuesRule]
type bucketEventNotificationListResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueuesRulesAction string

const (
	BucketEventNotificationListResponseQueuesRulesActionPutObject               BucketEventNotificationListResponseQueuesRulesAction = "PutObject"
	BucketEventNotificationListResponseQueuesRulesActionCopyObject              BucketEventNotificationListResponseQueuesRulesAction = "CopyObject"
	BucketEventNotificationListResponseQueuesRulesActionDeleteObject            BucketEventNotificationListResponseQueuesRulesAction = "DeleteObject"
	BucketEventNotificationListResponseQueuesRulesActionCompleteMultipartUpload BucketEventNotificationListResponseQueuesRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationListResponseQueuesRulesActionLifecycleDeletion       BucketEventNotificationListResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationListResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationListResponseQueuesRulesActionPutObject, BucketEventNotificationListResponseQueuesRulesActionCopyObject, BucketEventNotificationListResponseQueuesRulesActionDeleteObject, BucketEventNotificationListResponseQueuesRulesActionCompleteMultipartUpload, BucketEventNotificationListResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type BucketEventNotificationDeleteResponse = interface{}

type BucketEventNotificationGetResponse struct {
	// Queue ID.
	QueueID string `json:"queueId"`
	// Name of the queue.
	QueueName string                                   `json:"queueName"`
	Rules     []BucketEventNotificationGetResponseRule `json:"rules"`
	JSON      bucketEventNotificationGetResponseJSON   `json:"-"`
}

// bucketEventNotificationGetResponseJSON contains the JSON metadata for the struct
// [BucketEventNotificationGetResponse]
type bucketEventNotificationGetResponseJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions []BucketEventNotificationGetResponseRulesAction `json:"actions,required"`
	// Timestamp when the rule was created.
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description string `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix string `json:"prefix"`
	// Rule ID.
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix.
	Suffix string                                     `json:"suffix"`
	JSON   bucketEventNotificationGetResponseRuleJSON `json:"-"`
}

// bucketEventNotificationGetResponseRuleJSON contains the JSON metadata for the
// struct [BucketEventNotificationGetResponseRule]
type bucketEventNotificationGetResponseRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseRulesAction string

const (
	BucketEventNotificationGetResponseRulesActionPutObject               BucketEventNotificationGetResponseRulesAction = "PutObject"
	BucketEventNotificationGetResponseRulesActionCopyObject              BucketEventNotificationGetResponseRulesAction = "CopyObject"
	BucketEventNotificationGetResponseRulesActionDeleteObject            BucketEventNotificationGetResponseRulesAction = "DeleteObject"
	BucketEventNotificationGetResponseRulesActionCompleteMultipartUpload BucketEventNotificationGetResponseRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationGetResponseRulesActionLifecycleDeletion       BucketEventNotificationGetResponseRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationGetResponseRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseRulesActionPutObject, BucketEventNotificationGetResponseRulesActionCopyObject, BucketEventNotificationGetResponseRulesActionDeleteObject, BucketEventNotificationGetResponseRulesActionCompleteMultipartUpload, BucketEventNotificationGetResponseRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type BucketEventNotificationUpdateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications.
	Rules param.Field[[]BucketEventNotificationUpdateParamsRule] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEventNotificationUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketEventNotificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationUpdateParamsRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions param.Field[[]BucketEventNotificationUpdateParamsRulesAction] `json:"actions,required"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description param.Field[string] `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix.
	Suffix param.Field[string] `json:"suffix"`
}

func (r BucketEventNotificationUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationUpdateParamsRulesAction string

const (
	BucketEventNotificationUpdateParamsRulesActionPutObject               BucketEventNotificationUpdateParamsRulesAction = "PutObject"
	BucketEventNotificationUpdateParamsRulesActionCopyObject              BucketEventNotificationUpdateParamsRulesAction = "CopyObject"
	BucketEventNotificationUpdateParamsRulesActionDeleteObject            BucketEventNotificationUpdateParamsRulesAction = "DeleteObject"
	BucketEventNotificationUpdateParamsRulesActionCompleteMultipartUpload BucketEventNotificationUpdateParamsRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationUpdateParamsRulesActionLifecycleDeletion       BucketEventNotificationUpdateParamsRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateParamsRulesActionPutObject, BucketEventNotificationUpdateParamsRulesActionCopyObject, BucketEventNotificationUpdateParamsRulesActionDeleteObject, BucketEventNotificationUpdateParamsRulesActionCompleteMultipartUpload, BucketEventNotificationUpdateParamsRulesActionLifecycleDeletion:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketEventNotificationUpdateParamsCfR2Jurisdiction string

const (
	BucketEventNotificationUpdateParamsCfR2JurisdictionDefault BucketEventNotificationUpdateParamsCfR2Jurisdiction = "default"
	BucketEventNotificationUpdateParamsCfR2JurisdictionEu      BucketEventNotificationUpdateParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationUpdateParamsCfR2JurisdictionFedramp BucketEventNotificationUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateParamsCfR2JurisdictionDefault, BucketEventNotificationUpdateParamsCfR2JurisdictionEu, BucketEventNotificationUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   BucketEventNotificationUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationUpdateResponseEnvelope]
type bucketEventNotificationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationUpdateResponseEnvelopeSuccess bool

const (
	BucketEventNotificationUpdateResponseEnvelopeSuccessTrue BucketEventNotificationUpdateResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEventNotificationListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketEventNotificationListParamsCfR2Jurisdiction string

const (
	BucketEventNotificationListParamsCfR2JurisdictionDefault BucketEventNotificationListParamsCfR2Jurisdiction = "default"
	BucketEventNotificationListParamsCfR2JurisdictionEu      BucketEventNotificationListParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationListParamsCfR2JurisdictionFedramp BucketEventNotificationListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationListParamsCfR2JurisdictionDefault, BucketEventNotificationListParamsCfR2JurisdictionEu, BucketEventNotificationListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationListResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	Result   BucketEventNotificationListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationListResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationListResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationListResponseEnvelope]
type bucketEventNotificationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationListResponseEnvelopeSuccess bool

const (
	BucketEventNotificationListResponseEnvelopeSuccessTrue BucketEventNotificationListResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEventNotificationDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketEventNotificationDeleteParamsCfR2Jurisdiction string

const (
	BucketEventNotificationDeleteParamsCfR2JurisdictionDefault BucketEventNotificationDeleteParamsCfR2Jurisdiction = "default"
	BucketEventNotificationDeleteParamsCfR2JurisdictionEu      BucketEventNotificationDeleteParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationDeleteParamsCfR2JurisdictionFedramp BucketEventNotificationDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationDeleteParamsCfR2JurisdictionDefault, BucketEventNotificationDeleteParamsCfR2JurisdictionEu, BucketEventNotificationDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   BucketEventNotificationDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationDeleteResponseEnvelope]
type bucketEventNotificationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationDeleteResponseEnvelopeSuccess bool

const (
	BucketEventNotificationDeleteResponseEnvelopeSuccessTrue BucketEventNotificationDeleteResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction.
	Jurisdiction param.Field[BucketEventNotificationGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction.
type BucketEventNotificationGetParamsCfR2Jurisdiction string

const (
	BucketEventNotificationGetParamsCfR2JurisdictionDefault BucketEventNotificationGetParamsCfR2Jurisdiction = "default"
	BucketEventNotificationGetParamsCfR2JurisdictionEu      BucketEventNotificationGetParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationGetParamsCfR2JurisdictionFedramp BucketEventNotificationGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetParamsCfR2JurisdictionDefault, BucketEventNotificationGetParamsCfR2JurisdictionEu, BucketEventNotificationGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   BucketEventNotificationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationGetResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationGetResponseEnvelope]
type bucketEventNotificationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationGetResponseEnvelopeSuccess bool

const (
	BucketEventNotificationGetResponseEnvelopeSuccessTrue BucketEventNotificationGetResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
