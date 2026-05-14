// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Creates an account subscription.
func (r *SubscriptionService) New(ctx context.Context, params SubscriptionNewParams, opts ...option.RequestOption) (res *shared.Subscription, err error) {
	var env SubscriptionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an account subscription.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionIdentifier string, params SubscriptionUpdateParams, opts ...option.RequestOption) (res *shared.Subscription, err error) {
	var env SubscriptionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", params.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an account's subscription.
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionIdentifier string, body SubscriptionDeleteParams, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	var env SubscriptionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", body.AccountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all of an account's subscriptions.
func (r *SubscriptionService) Get(ctx context.Context, query SubscriptionGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[shared.Subscription], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions", query.AccountID)
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

// Lists all of an account's subscriptions.
func (r *SubscriptionService) GetAutoPaging(ctx context.Context, query SubscriptionGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[shared.Subscription] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type SubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                         `json:"subscription_id"`
	JSON           subscriptionDeleteResponseJSON `json:"-"`
}

// subscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponse]
type subscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	// Identifier
	AccountID    param.Field[string]      `path:"account_id,required"`
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   shared.Subscription   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseEnvelope]
type subscriptionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionNewResponseEnvelopeSuccess bool

const (
	SubscriptionNewResponseEnvelopeSuccessTrue SubscriptionNewResponseEnvelopeSuccess = true
)

func (r SubscriptionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	// Identifier
	AccountID    param.Field[string]      `path:"account_id,required"`
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   shared.Subscription   `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// subscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseEnvelope]
type subscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

func (r SubscriptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SubscriptionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   SubscriptionDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionDeleteResponseEnvelopeJSON    `json:"-"`
}

// subscriptionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponseEnvelope]
type subscriptionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionDeleteResponseEnvelopeSuccess bool

const (
	SubscriptionDeleteResponseEnvelopeSuccessTrue SubscriptionDeleteResponseEnvelopeSuccess = true
)

func (r SubscriptionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
