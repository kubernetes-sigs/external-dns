// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DestinationWebhookService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationWebhookService] method instead.
type DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewDestinationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationWebhookService(opts ...option.RequestOption) (r *DestinationWebhookService) {
	r = &DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *DestinationWebhookService) New(ctx context.Context, params DestinationWebhookNewParams, opts ...option.RequestOption) (res *DestinationWebhookNewResponse, err error) {
	var env DestinationWebhookNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a webhook destination.
func (r *DestinationWebhookService) Update(ctx context.Context, webhookID string, params DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *DestinationWebhookUpdateResponse, err error) {
	var env DestinationWebhookUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all configured webhook destinations.
func (r *DestinationWebhookService) List(ctx context.Context, query DestinationWebhookListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Webhooks], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", query.AccountID)
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

// Gets a list of all configured webhook destinations.
func (r *DestinationWebhookService) ListAutoPaging(ctx context.Context, query DestinationWebhookListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Webhooks] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured webhook destination.
func (r *DestinationWebhookService) Delete(ctx context.Context, webhookID string, body DestinationWebhookDeleteParams, opts ...option.RequestOption) (res *DestinationWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details for a single webhooks destination.
func (r *DestinationWebhookService) Get(ctx context.Context, webhookID string, query DestinationWebhookGetParams, opts ...option.RequestOption) (res *Webhooks, err error) {
	var env DestinationWebhookGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Webhooks struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Type of webhook endpoint.
	Type WebhooksType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string       `json:"url"`
	JSON webhooksJSON `json:"-"`
}

// webhooksJSON contains the JSON metadata for the struct [Webhooks]
type webhooksJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Webhooks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhooksJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type WebhooksType string

const (
	WebhooksTypeSlack   WebhooksType = "slack"
	WebhooksTypeGeneric WebhooksType = "generic"
	WebhooksTypeGchat   WebhooksType = "gchat"
)

func (r WebhooksType) IsKnown() bool {
	switch r {
	case WebhooksTypeSlack, WebhooksTypeGeneric, WebhooksTypeGchat:
		return true
	}
	return false
}

type DestinationWebhookNewResponse struct {
	// UUID
	ID   string                            `json:"id"`
	JSON destinationWebhookNewResponseJSON `json:"-"`
}

// destinationWebhookNewResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookNewResponse]
type destinationWebhookNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookUpdateResponse struct {
	// UUID
	ID   string                               `json:"id"`
	JSON destinationWebhookUpdateResponseJSON `json:"-"`
}

// destinationWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookUpdateResponse]
type destinationWebhookUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DestinationWebhookDeleteResponseSuccess    `json:"success,required"`
	ResultInfo DestinationWebhookDeleteResponseResultInfo `json:"result_info"`
	JSON       destinationWebhookDeleteResponseJSON       `json:"-"`
}

// destinationWebhookDeleteResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookDeleteResponse]
type destinationWebhookDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookDeleteResponseSuccess bool

const (
	DestinationWebhookDeleteResponseSuccessTrue DestinationWebhookDeleteResponseSuccess = true
)

func (r DestinationWebhookDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       destinationWebhookDeleteResponseResultInfoJSON `json:"-"`
}

// destinationWebhookDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [DestinationWebhookDeleteResponseResultInfo]
type destinationWebhookDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationWebhookNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DestinationWebhookNewResponse                `json:"result"`
	JSON    destinationWebhookNewResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookNewResponseEnvelope]
type destinationWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookNewResponseEnvelopeSuccess bool

const (
	DestinationWebhookNewResponseEnvelopeSuccessTrue DestinationWebhookNewResponseEnvelopeSuccess = true
)

func (r DestinationWebhookNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DestinationWebhookUpdateResponse                `json:"result"`
	JSON    destinationWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookUpdateResponseEnvelope]
type destinationWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookUpdateResponseEnvelopeSuccess bool

const (
	DestinationWebhookUpdateResponseEnvelopeSuccessTrue DestinationWebhookUpdateResponseEnvelopeSuccess = true
)

func (r DestinationWebhookUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Webhooks                                     `json:"result"`
	JSON    destinationWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookGetResponseEnvelope]
type destinationWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookGetResponseEnvelopeSuccess bool

const (
	DestinationWebhookGetResponseEnvelopeSuccessTrue DestinationWebhookGetResponseEnvelopeSuccess = true
)

func (r DestinationWebhookGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
