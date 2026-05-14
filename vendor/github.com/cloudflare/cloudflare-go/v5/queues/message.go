// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

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

// MessageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r *MessageService) {
	r = &MessageService{}
	r.Options = opts
	return
}

// Acknowledge + Retry messages from a Queue
func (r *MessageService) Ack(ctx context.Context, queueID string, params MessageAckParams, opts ...option.RequestOption) (res *MessageAckResponse, err error) {
	var env MessageAckResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/ack", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Push a batch of message to a Queue
func (r *MessageService) BulkPush(ctx context.Context, queueID string, params MessageBulkPushParams, opts ...option.RequestOption) (res *MessageBulkPushResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/batch", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Pull a batch of messages from a Queue
func (r *MessageService) Pull(ctx context.Context, queueID string, params MessagePullParams, opts ...option.RequestOption) (res *MessagePullResponse, err error) {
	var env MessagePullResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/pull", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Push a message to a Queue
func (r *MessageService) Push(ctx context.Context, queueID string, params MessagePushParams, opts ...option.RequestOption) (res *MessagePushResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MessageAckResponse struct {
	// The number of messages that were succesfully acknowledged.
	AckCount float64 `json:"ackCount"`
	// The number of messages that were succesfully retried.
	RetryCount float64                `json:"retryCount"`
	Warnings   []string               `json:"warnings"`
	JSON       messageAckResponseJSON `json:"-"`
}

// messageAckResponseJSON contains the JSON metadata for the struct
// [MessageAckResponse]
type messageAckResponseJSON struct {
	AckCount    apijson.Field
	RetryCount  apijson.Field
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseJSON) RawJSON() string {
	return r.raw
}

type MessageBulkPushResponse struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success MessageBulkPushResponseSuccess `json:"success"`
	JSON    messageBulkPushResponseJSON    `json:"-"`
}

// messageBulkPushResponseJSON contains the JSON metadata for the struct
// [MessageBulkPushResponse]
type messageBulkPushResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBulkPushResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBulkPushResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessageBulkPushResponseSuccess bool

const (
	MessageBulkPushResponseSuccessTrue MessageBulkPushResponseSuccess = true
)

func (r MessageBulkPushResponseSuccess) IsKnown() bool {
	switch r {
	case MessageBulkPushResponseSuccessTrue:
		return true
	}
	return false
}

type MessagePullResponse struct {
	// The number of unacknowledged messages in the queue
	MessageBacklogCount float64                      `json:"message_backlog_count"`
	Messages            []MessagePullResponseMessage `json:"messages"`
	JSON                messagePullResponseJSON      `json:"-"`
}

// messagePullResponseJSON contains the JSON metadata for the struct
// [MessagePullResponse]
type messagePullResponseJSON struct {
	MessageBacklogCount apijson.Field
	Messages            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MessagePullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseJSON) RawJSON() string {
	return r.raw
}

type MessagePullResponseMessage struct {
	ID       string  `json:"id"`
	Attempts float64 `json:"attempts"`
	Body     string  `json:"body"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID     string                         `json:"lease_id"`
	Metadata    interface{}                    `json:"metadata"`
	TimestampMs float64                        `json:"timestamp_ms"`
	JSON        messagePullResponseMessageJSON `json:"-"`
}

// messagePullResponseMessageJSON contains the JSON metadata for the struct
// [MessagePullResponseMessage]
type messagePullResponseMessageJSON struct {
	ID          apijson.Field
	Attempts    apijson.Field
	Body        apijson.Field
	LeaseID     apijson.Field
	Metadata    apijson.Field
	TimestampMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MessagePushResponse struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success MessagePushResponseSuccess `json:"success"`
	JSON    messagePushResponseJSON    `json:"-"`
}

// messagePushResponseJSON contains the JSON metadata for the struct
// [MessagePushResponse]
type messagePushResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePushResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePushResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessagePushResponseSuccess bool

const (
	MessagePushResponseSuccessTrue MessagePushResponseSuccess = true
)

func (r MessagePushResponseSuccess) IsKnown() bool {
	switch r {
	case MessagePushResponseSuccessTrue:
		return true
	}
	return false
}

type MessageAckParams struct {
	// A Resource identifier.
	AccountID param.Field[string]                  `path:"account_id,required"`
	Acks      param.Field[[]MessageAckParamsAck]   `json:"acks"`
	Retries   param.Field[[]MessageAckParamsRetry] `json:"retries"`
}

func (r MessageAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsAck struct {
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsAck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckParamsRetry struct {
	// The number of seconds to delay before making the message available for another
	// attempt.
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r MessageAckParamsRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageAckResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   MessageAckResponse    `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessageAckResponseEnvelopeSuccess `json:"success"`
	JSON    messageAckResponseEnvelopeJSON    `json:"-"`
}

// messageAckResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessageAckResponseEnvelope]
type messageAckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAckResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessageAckResponseEnvelopeSuccess bool

const (
	MessageAckResponseEnvelopeSuccessTrue MessageAckResponseEnvelopeSuccess = true
)

func (r MessageAckResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessageAckResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MessageBulkPushParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The number of seconds to wait for attempting to deliver this batch to consumers
	DelaySeconds param.Field[float64]                             `json:"delay_seconds"`
	Messages     param.Field[[]MessageBulkPushParamsMessageUnion] `json:"messages"`
}

func (r MessageBulkPushParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBulkPushParamsMessage struct {
	Body        param.Field[interface{}]                              `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessage) implementsMessageBulkPushParamsMessageUnion() {}

// Satisfied by [queues.MessageBulkPushParamsMessagesMqQueueMessageText],
// [queues.MessageBulkPushParamsMessagesMqQueueMessageJson],
// [MessageBulkPushParamsMessage].
type MessageBulkPushParamsMessageUnion interface {
	implementsMessageBulkPushParamsMessageUnion()
}

type MessageBulkPushParamsMessagesMqQueueMessageText struct {
	Body        param.Field[string]                                                     `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesMqQueueMessageTextContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessagesMqQueueMessageText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessagesMqQueueMessageText) implementsMessageBulkPushParamsMessageUnion() {
}

type MessageBulkPushParamsMessagesMqQueueMessageTextContentType string

const (
	MessageBulkPushParamsMessagesMqQueueMessageTextContentTypeText MessageBulkPushParamsMessagesMqQueueMessageTextContentType = "text"
)

func (r MessageBulkPushParamsMessagesMqQueueMessageTextContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesMqQueueMessageTextContentTypeText:
		return true
	}
	return false
}

type MessageBulkPushParamsMessagesMqQueueMessageJson struct {
	Body        param.Field[interface{}]                                                `json:"body"`
	ContentType param.Field[MessageBulkPushParamsMessagesMqQueueMessageJsonContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessageBulkPushParamsMessagesMqQueueMessageJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBulkPushParamsMessagesMqQueueMessageJson) implementsMessageBulkPushParamsMessageUnion() {
}

type MessageBulkPushParamsMessagesMqQueueMessageJsonContentType string

const (
	MessageBulkPushParamsMessagesMqQueueMessageJsonContentTypeJson MessageBulkPushParamsMessagesMqQueueMessageJsonContentType = "json"
)

func (r MessageBulkPushParamsMessagesMqQueueMessageJsonContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesMqQueueMessageJsonContentTypeJson:
		return true
	}
	return false
}

type MessageBulkPushParamsMessagesContentType string

const (
	MessageBulkPushParamsMessagesContentTypeText MessageBulkPushParamsMessagesContentType = "text"
	MessageBulkPushParamsMessagesContentTypeJson MessageBulkPushParamsMessagesContentType = "json"
)

func (r MessageBulkPushParamsMessagesContentType) IsKnown() bool {
	switch r {
	case MessageBulkPushParamsMessagesContentTypeText, MessageBulkPushParamsMessagesContentTypeJson:
		return true
	}
	return false
}

type MessagePullParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r MessagePullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessagePullResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   MessagePullResponse   `json:"result"`
	// Indicates if the API call was successful or not.
	Success MessagePullResponseEnvelopeSuccess `json:"success"`
	JSON    messagePullResponseEnvelopeJSON    `json:"-"`
}

// messagePullResponseEnvelopeJSON contains the JSON metadata for the struct
// [MessagePullResponseEnvelope]
type messagePullResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagePullResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagePullResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MessagePullResponseEnvelopeSuccess bool

const (
	MessagePullResponseEnvelopeSuccessTrue MessagePullResponseEnvelopeSuccess = true
)

func (r MessagePullResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MessagePullResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MessagePushParams struct {
	// A Resource identifier.
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      MessagePushParamsBodyUnion `json:"body"`
}

func (r MessagePushParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MessagePushParamsBody struct {
	Body        param.Field[interface{}]                      `json:"body"`
	ContentType param.Field[MessagePushParamsBodyContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBody) implementsMessagePushParamsBodyUnion() {}

// Satisfied by [queues.MessagePushParamsBodyMqQueueMessageText],
// [queues.MessagePushParamsBodyMqQueueMessageJson], [MessagePushParamsBody].
type MessagePushParamsBodyUnion interface {
	implementsMessagePushParamsBodyUnion()
}

type MessagePushParamsBodyMqQueueMessageText struct {
	Body        param.Field[string]                                             `json:"body"`
	ContentType param.Field[MessagePushParamsBodyMqQueueMessageTextContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBodyMqQueueMessageText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBodyMqQueueMessageText) implementsMessagePushParamsBodyUnion() {}

type MessagePushParamsBodyMqQueueMessageTextContentType string

const (
	MessagePushParamsBodyMqQueueMessageTextContentTypeText MessagePushParamsBodyMqQueueMessageTextContentType = "text"
)

func (r MessagePushParamsBodyMqQueueMessageTextContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyMqQueueMessageTextContentTypeText:
		return true
	}
	return false
}

type MessagePushParamsBodyMqQueueMessageJson struct {
	Body        param.Field[interface{}]                                        `json:"body"`
	ContentType param.Field[MessagePushParamsBodyMqQueueMessageJsonContentType] `json:"content_type"`
	// The number of seconds to wait for attempting to deliver this message to
	// consumers
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
}

func (r MessagePushParamsBodyMqQueueMessageJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessagePushParamsBodyMqQueueMessageJson) implementsMessagePushParamsBodyUnion() {}

type MessagePushParamsBodyMqQueueMessageJsonContentType string

const (
	MessagePushParamsBodyMqQueueMessageJsonContentTypeJson MessagePushParamsBodyMqQueueMessageJsonContentType = "json"
)

func (r MessagePushParamsBodyMqQueueMessageJsonContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyMqQueueMessageJsonContentTypeJson:
		return true
	}
	return false
}

type MessagePushParamsBodyContentType string

const (
	MessagePushParamsBodyContentTypeText MessagePushParamsBodyContentType = "text"
	MessagePushParamsBodyContentTypeJson MessagePushParamsBodyContentType = "json"
)

func (r MessagePushParamsBodyContentType) IsKnown() bool {
	switch r {
	case MessagePushParamsBodyContentTypeText, MessagePushParamsBodyContentTypeJson:
		return true
	}
	return false
}
