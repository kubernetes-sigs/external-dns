// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// ConsumerService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConsumerService] method instead.
type ConsumerService struct {
	Options []option.RequestOption
}

// NewConsumerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConsumerService(opts ...option.RequestOption) (r *ConsumerService) {
	r = &ConsumerService{}
	r.Options = opts
	return
}

// Creates a new consumer for a Queue
func (r *ConsumerService) New(ctx context.Context, queueID string, params ConsumerNewParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *ConsumerService) Update(ctx context.Context, queueID string, consumerID string, params ConsumerUpdateParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", params.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the consumers for a Queue
func (r *ConsumerService) List(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Consumer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", query.AccountID, queueID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns the consumers for a Queue
func (r *ConsumerService) ListAutoPaging(ctx context.Context, queueID string, query ConsumerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Consumer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, queueID, query, opts...))
}

// Deletes the consumer for a queue.
func (r *ConsumerService) Delete(ctx context.Context, queueID string, consumerID string, body ConsumerDeleteParams, opts ...option.RequestOption) (res *ConsumerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", body.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches the consumer for a queue by consumer id
func (r *ConsumerService) Get(ctx context.Context, queueID string, consumerID string, query ConsumerGetParams, opts ...option.RequestOption) (res *Consumer, err error) {
	var env ConsumerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", query.AccountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Consumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID string `json:"queue_id"`
	// Name of a Worker
	Script string `json:"script"`
	// This field can have the runtime type of [ConsumerMqWorkerConsumerSettings],
	// [ConsumerMqHTTPConsumerSettings].
	Settings interface{}  `json:"settings"`
	Type     ConsumerType `json:"type"`
	JSON     consumerJSON `json:"-"`
	union    ConsumerUnion
}

// consumerJSON contains the JSON metadata for the struct [Consumer]
type consumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Script      apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r consumerJSON) RawJSON() string {
	return r.raw
}

func (r *Consumer) UnmarshalJSON(data []byte) (err error) {
	*r = Consumer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConsumerUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ConsumerMqWorkerConsumer],
// [ConsumerMqHTTPConsumer].
func (r Consumer) AsUnion() ConsumerUnion {
	return r.union
}

// Union satisfied by [ConsumerMqWorkerConsumer] or [ConsumerMqHTTPConsumer].
type ConsumerUnion interface {
	implementsConsumer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConsumerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerMqWorkerConsumer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConsumerMqHTTPConsumer{}),
		},
	)
}

type ConsumerMqWorkerConsumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID string `json:"queue_id"`
	// Name of a Worker
	Script   string                           `json:"script"`
	Settings ConsumerMqWorkerConsumerSettings `json:"settings"`
	Type     ConsumerMqWorkerConsumerType     `json:"type"`
	JSON     consumerMqWorkerConsumerJSON     `json:"-"`
}

// consumerMqWorkerConsumerJSON contains the JSON metadata for the struct
// [ConsumerMqWorkerConsumer]
type consumerMqWorkerConsumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Script      apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqWorkerConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqWorkerConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqWorkerConsumer) implementsConsumer() {}

type ConsumerMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency float64 `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs float64 `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64                              `json:"retry_delay"`
	JSON       consumerMqWorkerConsumerSettingsJSON `json:"-"`
}

// consumerMqWorkerConsumerSettingsJSON contains the JSON metadata for the struct
// [ConsumerMqWorkerConsumerSettings]
type consumerMqWorkerConsumerSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConsumerMqWorkerConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqWorkerConsumerSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqWorkerConsumerType string

const (
	ConsumerMqWorkerConsumerTypeWorker ConsumerMqWorkerConsumerType = "worker"
)

func (r ConsumerMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case ConsumerMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type ConsumerMqHTTPConsumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID  string                         `json:"queue_id"`
	Settings ConsumerMqHTTPConsumerSettings `json:"settings"`
	Type     ConsumerMqHTTPConsumerType     `json:"type"`
	JSON     consumerMqHTTPConsumerJSON     `json:"-"`
}

// consumerMqHTTPConsumerJSON contains the JSON metadata for the struct
// [ConsumerMqHTTPConsumer]
type consumerMqHTTPConsumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerMqHTTPConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqHTTPConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ConsumerMqHTTPConsumer) implementsConsumer() {}

type ConsumerMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                            `json:"visibility_timeout_ms"`
	JSON                consumerMqHTTPConsumerSettingsJSON `json:"-"`
}

// consumerMqHTTPConsumerSettingsJSON contains the JSON metadata for the struct
// [ConsumerMqHTTPConsumerSettings]
type consumerMqHTTPConsumerSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConsumerMqHTTPConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerMqHTTPConsumerSettingsJSON) RawJSON() string {
	return r.raw
}

type ConsumerMqHTTPConsumerType string

const (
	ConsumerMqHTTPConsumerTypeHTTPPull ConsumerMqHTTPConsumerType = "http_pull"
)

func (r ConsumerMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case ConsumerMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerType string

const (
	ConsumerTypeWorker   ConsumerType = "worker"
	ConsumerTypeHTTPPull ConsumerType = "http_pull"
)

func (r ConsumerType) IsKnown() bool {
	switch r {
	case ConsumerTypeWorker, ConsumerTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerParam struct {
	// Name of a Worker
	ScriptName param.Field[string]       `json:"script_name"`
	Settings   param.Field[interface{}]  `json:"settings"`
	Type       param.Field[ConsumerType] `json:"type"`
}

func (r ConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerParam) implementsConsumerUnionParam() {}

// Satisfied by [queues.ConsumerMqWorkerConsumerParam],
// [queues.ConsumerMqHTTPConsumerParam], [ConsumerParam].
type ConsumerUnionParam interface {
	implementsConsumerUnionParam()
}

type ConsumerMqWorkerConsumerParam struct {
	// Name of a Worker
	ScriptName param.Field[string]                                `json:"script_name"`
	Settings   param.Field[ConsumerMqWorkerConsumerSettingsParam] `json:"settings"`
	Type       param.Field[ConsumerMqWorkerConsumerType]          `json:"type"`
}

func (r ConsumerMqWorkerConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqWorkerConsumerParam) implementsConsumerUnionParam() {}

type ConsumerMqWorkerConsumerSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r ConsumerMqWorkerConsumerSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerMqHTTPConsumerParam struct {
	Settings param.Field[ConsumerMqHTTPConsumerSettingsParam] `json:"settings"`
	Type     param.Field[ConsumerMqHTTPConsumerType]          `json:"type"`
}

func (r ConsumerMqHTTPConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerMqHTTPConsumerParam) implementsConsumerUnionParam() {}

type ConsumerMqHTTPConsumerSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r ConsumerMqHTTPConsumerSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success ConsumerDeleteResponseSuccess `json:"success"`
	JSON    consumerDeleteResponseJSON    `json:"-"`
}

// consumerDeleteResponseJSON contains the JSON metadata for the struct
// [ConsumerDeleteResponse]
type consumerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerDeleteResponseSuccess bool

const (
	ConsumerDeleteResponseSuccessTrue ConsumerDeleteResponseSuccess = true
)

func (r ConsumerDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ConsumerDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ConsumerNewParams struct {
	// A Resource identifier.
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      ConsumerNewParamsBodyUnion `json:"body"`
}

func (r ConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConsumerNewParamsBody struct {
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]                    `json:"script_name"`
	Settings   param.Field[interface{}]               `json:"settings"`
	Type       param.Field[ConsumerNewParamsBodyType] `json:"type"`
}

func (r ConsumerNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBody) implementsConsumerNewParamsBodyUnion() {}

// Satisfied by [queues.ConsumerNewParamsBodyMqWorkerConsumer],
// [queues.ConsumerNewParamsBodyMqHTTPConsumer], [ConsumerNewParamsBody].
type ConsumerNewParamsBodyUnion interface {
	implementsConsumerNewParamsBodyUnion()
}

type ConsumerNewParamsBodyMqWorkerConsumer struct {
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]                                        `json:"script_name"`
	Settings   param.Field[ConsumerNewParamsBodyMqWorkerConsumerSettings] `json:"settings"`
	Type       param.Field[ConsumerNewParamsBodyMqWorkerConsumerType]     `json:"type"`
}

func (r ConsumerNewParamsBodyMqWorkerConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqWorkerConsumer) implementsConsumerNewParamsBodyUnion() {}

type ConsumerNewParamsBodyMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r ConsumerNewParamsBodyMqWorkerConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqWorkerConsumerType string

const (
	ConsumerNewParamsBodyMqWorkerConsumerTypeWorker ConsumerNewParamsBodyMqWorkerConsumerType = "worker"
)

func (r ConsumerNewParamsBodyMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type ConsumerNewParamsBodyMqHTTPConsumer struct {
	DeadLetterQueue param.Field[string]                                      `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerNewParamsBodyMqHTTPConsumerSettings] `json:"settings"`
	Type            param.Field[ConsumerNewParamsBodyMqHTTPConsumerType]     `json:"type"`
}

func (r ConsumerNewParamsBodyMqHTTPConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerNewParamsBodyMqHTTPConsumer) implementsConsumerNewParamsBodyUnion() {}

type ConsumerNewParamsBodyMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r ConsumerNewParamsBodyMqHTTPConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerNewParamsBodyMqHTTPConsumerType string

const (
	ConsumerNewParamsBodyMqHTTPConsumerTypeHTTPPull ConsumerNewParamsBodyMqHTTPConsumerType = "http_pull"
)

func (r ConsumerNewParamsBodyMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewParamsBodyType string

const (
	ConsumerNewParamsBodyTypeWorker   ConsumerNewParamsBodyType = "worker"
	ConsumerNewParamsBodyTypeHTTPPull ConsumerNewParamsBodyType = "http_pull"
)

func (r ConsumerNewParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerNewParamsBodyTypeWorker, ConsumerNewParamsBodyTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Consumer              `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerNewResponseEnvelopeSuccess `json:"success"`
	JSON    consumerNewResponseEnvelopeJSON    `json:"-"`
}

// consumerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerNewResponseEnvelope]
type consumerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerNewResponseEnvelopeSuccess bool

const (
	ConsumerNewResponseEnvelopeSuccessTrue ConsumerNewResponseEnvelopeSuccess = true
)

func (r ConsumerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConsumerUpdateParams struct {
	// A Resource identifier.
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      ConsumerUpdateParamsBodyUnion `json:"body,required"`
}

func (r ConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConsumerUpdateParamsBody struct {
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]                       `json:"script_name"`
	Settings   param.Field[interface{}]                  `json:"settings"`
	Type       param.Field[ConsumerUpdateParamsBodyType] `json:"type"`
}

func (r ConsumerUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBody) implementsConsumerUpdateParamsBodyUnion() {}

// Satisfied by [queues.ConsumerUpdateParamsBodyMqWorkerConsumer],
// [queues.ConsumerUpdateParamsBodyMqHTTPConsumer], [ConsumerUpdateParamsBody].
type ConsumerUpdateParamsBodyUnion interface {
	implementsConsumerUpdateParamsBodyUnion()
}

type ConsumerUpdateParamsBodyMqWorkerConsumer struct {
	DeadLetterQueue param.Field[string] `json:"dead_letter_queue"`
	// Name of a Worker
	ScriptName param.Field[string]                                           `json:"script_name"`
	Settings   param.Field[ConsumerUpdateParamsBodyMqWorkerConsumerSettings] `json:"settings"`
	Type       param.Field[ConsumerUpdateParamsBodyMqWorkerConsumerType]     `json:"type"`
}

func (r ConsumerUpdateParamsBodyMqWorkerConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqWorkerConsumer) implementsConsumerUpdateParamsBodyUnion() {}

type ConsumerUpdateParamsBodyMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r ConsumerUpdateParamsBodyMqWorkerConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqWorkerConsumerType string

const (
	ConsumerUpdateParamsBodyMqWorkerConsumerTypeWorker ConsumerUpdateParamsBodyMqWorkerConsumerType = "worker"
)

func (r ConsumerUpdateParamsBodyMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type ConsumerUpdateParamsBodyMqHTTPConsumer struct {
	DeadLetterQueue param.Field[string]                                         `json:"dead_letter_queue"`
	Settings        param.Field[ConsumerUpdateParamsBodyMqHTTPConsumerSettings] `json:"settings"`
	Type            param.Field[ConsumerUpdateParamsBodyMqHTTPConsumerType]     `json:"type"`
}

func (r ConsumerUpdateParamsBodyMqHTTPConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsumerUpdateParamsBodyMqHTTPConsumer) implementsConsumerUpdateParamsBodyUnion() {}

type ConsumerUpdateParamsBodyMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r ConsumerUpdateParamsBodyMqHTTPConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConsumerUpdateParamsBodyMqHTTPConsumerType string

const (
	ConsumerUpdateParamsBodyMqHTTPConsumerTypeHTTPPull ConsumerUpdateParamsBodyMqHTTPConsumerType = "http_pull"
)

func (r ConsumerUpdateParamsBodyMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerUpdateParamsBodyType string

const (
	ConsumerUpdateParamsBodyTypeWorker   ConsumerUpdateParamsBodyType = "worker"
	ConsumerUpdateParamsBodyTypeHTTPPull ConsumerUpdateParamsBodyType = "http_pull"
)

func (r ConsumerUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case ConsumerUpdateParamsBodyTypeWorker, ConsumerUpdateParamsBodyTypeHTTPPull:
		return true
	}
	return false
}

type ConsumerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Consumer              `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    consumerUpdateResponseEnvelopeJSON    `json:"-"`
}

// consumerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerUpdateResponseEnvelope]
type consumerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerUpdateResponseEnvelopeSuccess bool

const (
	ConsumerUpdateResponseEnvelopeSuccessTrue ConsumerUpdateResponseEnvelopeSuccess = true
)

func (r ConsumerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConsumerListParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerDeleteParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerGetParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConsumerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Consumer              `json:"result"`
	// Indicates if the API call was successful or not.
	Success ConsumerGetResponseEnvelopeSuccess `json:"success"`
	JSON    consumerGetResponseEnvelopeJSON    `json:"-"`
}

// consumerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConsumerGetResponseEnvelope]
type consumerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type ConsumerGetResponseEnvelopeSuccess bool

const (
	ConsumerGetResponseEnvelopeSuccessTrue ConsumerGetResponseEnvelopeSuccess = true
)

func (r ConsumerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConsumerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
