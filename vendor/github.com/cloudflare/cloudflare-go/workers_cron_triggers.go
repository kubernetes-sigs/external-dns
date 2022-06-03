package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WorkerCronTriggerResponse represents the response from the Worker cron trigger
// API endpoint.
type WorkerCronTriggerResponse struct {
	Response
	Result WorkerCronTriggerSchedules `json:"result"`
}

// WorkerCronTriggerSchedules contains the schedule of Worker cron triggers.
type WorkerCronTriggerSchedules struct {
	Schedules []WorkerCronTrigger `json:"schedules"`
}

// WorkerCronTrigger holds an individual cron schedule for a worker.
type WorkerCronTrigger struct {
	Cron       string     `json:"cron"`
	CreatedOn  *time.Time `json:"created_on,omitempty"`
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
}

// ListWorkerCronTriggers fetches all available cron triggers for a single Worker
// script.
//
// API reference: https://api.cloudflare.com/#worker-cron-trigger-get-cron-triggers
func (api *API) ListWorkerCronTriggers(ctx context.Context, accountID, scriptName string) ([]WorkerCronTrigger, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WorkerCronTrigger{}, err
	}

	result := WorkerCronTriggerResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []WorkerCronTrigger{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result.Schedules, err
}

// UpdateWorkerCronTriggers updates a single schedule for a Worker cron trigger.
//
// API reference: https://api.cloudflare.com/#worker-cron-trigger-update-cron-triggers
func (api *API) UpdateWorkerCronTriggers(ctx context.Context, accountID, scriptName string, crons []WorkerCronTrigger) ([]WorkerCronTrigger, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, crons)
	if err != nil {
		return []WorkerCronTrigger{}, err
	}

	result := WorkerCronTriggerResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []WorkerCronTrigger{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result.Schedules, err
}
