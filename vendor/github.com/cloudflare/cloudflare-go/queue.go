package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingQueueName         = errors.New("required queue name is missing")
	ErrMissingQueueConsumerName = errors.New("required queue consumer name is missing")
)

type Queue struct {
	ID                  string          `json:"queue_id,omitempty"`
	Name                string          `json:"queue_name,omitempty"`
	CreatedOn           *time.Time      `json:"created_on,omitempty"`
	ModifiedOn          *time.Time      `json:"modified_on,omitempty"`
	ProducersTotalCount int             `json:"producers_total_count,omitempty"`
	Producers           []QueueProducer `json:"producers,omitempty"`
	ConsumersTotalCount int             `json:"consumers_total_count,omitempty"`
	Consumers           []QueueConsumer `json:"consumers,omitempty"`
}

type QueueProducer struct {
	Service     string `json:"service,omitempty"`
	Environment string `json:"environment,omitempty"`
}

type QueueConsumer struct {
	Name            string                `json:"-"`
	Service         string                `json:"service,omitempty"`
	ScriptName      string                `json:"script_name,omitempty"`
	Environment     string                `json:"environment,omitempty"`
	Settings        QueueConsumerSettings `json:"settings,omitempty"`
	QueueName       string                `json:"queue_name,omitempty"`
	CreatedOn       *time.Time            `json:"created_on,omitempty"`
	DeadLetterQueue string                `json:"dead_letter_queue,omitempty"`
}

type QueueConsumerSettings struct {
	BatchSize   int `json:"batch_size,omitempty"`
	MaxRetires  int `json:"max_retries,omitempty"`
	MaxWaitTime int `json:"max_wait_time_ms,omitempty"`
}

type QueueListResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []Queue `json:"result"`
}

type CreateQueueParams struct {
	Name string `json:"queue_name"`
}

type QueueResponse struct {
	Response
	Result Queue `json:"result"`
}

type ListQueueConsumersResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []QueueConsumer `json:"result"`
}

type ListQueuesParams struct {
	ResultInfo
}

type QueueConsumerResponse struct {
	Response
	Result QueueConsumer `json:"result"`
}

type UpdateQueueParams struct {
	Name        string `json:"-"`
	UpdatedName string `json:"queue_name,omitempty"`
}

type ListQueueConsumersParams struct {
	QueueName string `url:"-"`
	ResultInfo
}

type CreateQueueConsumerParams struct {
	QueueName string `json:"-"`
	Consumer  QueueConsumer
}

type UpdateQueueConsumerParams struct {
	QueueName string `json:"-"`
	Consumer  QueueConsumer
}

type DeleteQueueConsumerParams struct {
	QueueName, ConsumerName string
}

// ListQueues returns the queues owned by an account.
//
// API reference: https://api.cloudflare.com/#queue-list-queues
func (api *API) ListQueues(ctx context.Context, rc *ResourceContainer, params ListQueuesParams) ([]Queue, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []Queue{}, &ResultInfo{}, ErrMissingAccountID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = 50
	}
	if params.Page < 1 {
		params.Page = 1
	}

	var queues []Queue
	var qResponse QueueListResponse
	for {
		qResponse = QueueListResponse{}
		uri := buildURI(fmt.Sprintf("/accounts/%s/workers/queues", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []Queue{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &qResponse)
		if err != nil {
			return []Queue{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
		}

		queues = append(queues, qResponse.Result...)
		params.ResultInfo = qResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return queues, &qResponse.ResultInfo, nil
}

// CreateQueue creates a new queue.
//
// API reference: https://api.cloudflare.com/#queue-create-queue
func (api *API) CreateQueue(ctx context.Context, rc *ResourceContainer, queue CreateQueueParams) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}

	if queue.Name == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, queue)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteQueue deletes a queue.
//
// API reference: https://api.cloudflare.com/#queue-delete-queue
func (api *API) DeleteQueue(ctx context.Context, rc *ResourceContainer, queueName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if queueName == "" {
		return ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, queueName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	return nil
}

// GetQueue returns a single queue based on the name.
//
// API reference: https://api.cloudflare.com/#queue-get-queue
func (api *API) GetQueue(ctx context.Context, rc *ResourceContainer, queueName string) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}

	if queueName == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, queueName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateQueue updates a queue.
//
// API reference: https://api.cloudflare.com/#queue-update-queue
func (api *API) UpdateQueue(ctx context.Context, rc *ResourceContainer, params UpdateQueueParams) (Queue, error) {
	if rc.Identifier == "" {
		return Queue{}, ErrMissingAccountID
	}

	if params.Name == "" || params.UpdatedName == "" {
		return Queue{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s", rc.Identifier, params.Name)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Queue{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListQueueConsumers returns the consumers of a queue.
//
// API reference: https://api.cloudflare.com/#queue-list-queue-consumers
func (api *API) ListQueueConsumers(ctx context.Context, rc *ResourceContainer, params ListQueueConsumersParams) ([]QueueConsumer, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []QueueConsumer{}, &ResultInfo{}, ErrMissingAccountID
	}

	if params.QueueName == "" {
		return []QueueConsumer{}, &ResultInfo{}, ErrMissingQueueName
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = 50
	}
	if params.Page < 1 {
		params.Page = 1
	}

	var queuesConsumers []QueueConsumer
	var qResponse ListQueueConsumersResponse
	for {
		qResponse = ListQueueConsumersResponse{}
		uri := buildURI(fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", rc.Identifier, params.QueueName), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []QueueConsumer{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &qResponse)
		if err != nil {
			return []QueueConsumer{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
		}

		queuesConsumers = append(queuesConsumers, qResponse.Result...)
		params.ResultInfo = qResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return queuesConsumers, &qResponse.ResultInfo, nil
}

// CreateQueueConsumer creates a new consumer for a queue.
//
// API reference: https://api.cloudflare.com/#queue-create-queue-consumer
func (api *API) CreateQueueConsumer(ctx context.Context, rc *ResourceContainer, params CreateQueueConsumerParams) (QueueConsumer, error) {
	if rc.Identifier == "" {
		return QueueConsumer{}, ErrMissingAccountID
	}

	if params.QueueName == "" {
		return QueueConsumer{}, ErrMissingQueueName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers", rc.Identifier, params.QueueName)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Consumer)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueConsumerResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteQueueConsumer deletes the consumer for a queue..
//
// API reference: https://api.cloudflare.com/#queue-delete-queue-consumer
func (api *API) DeleteQueueConsumer(ctx context.Context, rc *ResourceContainer, params DeleteQueueConsumerParams) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if params.QueueName == "" {
		return ErrMissingQueueName
	}

	if params.ConsumerName == "" {
		return ErrMissingQueueConsumerName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", rc.Identifier, params.QueueName, params.ConsumerName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}

// UpdateQueueConsumer updates the consumer for a queue, or creates one if it does not exist..
//
// API reference: https://api.cloudflare.com/#queue-update-queue-consumer
func (api *API) UpdateQueueConsumer(ctx context.Context, rc *ResourceContainer, params UpdateQueueConsumerParams) (QueueConsumer, error) {
	if rc.Identifier == "" {
		return QueueConsumer{}, ErrMissingAccountID
	}

	if params.QueueName == "" {
		return QueueConsumer{}, ErrMissingQueueName
	}

	if params.Consumer.Name == "" {
		return QueueConsumer{}, ErrMissingQueueConsumerName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/queues/%s/consumers/%s", rc.Identifier, params.QueueName, params.Consumer.Name)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Consumer)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueueConsumerResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return QueueConsumer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
