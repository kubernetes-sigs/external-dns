// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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
)

// InvestigateMoveService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateMoveService] method instead.
type InvestigateMoveService struct {
	Options []option.RequestOption
}

// NewInvestigateMoveService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateMoveService(opts ...option.RequestOption) (r *InvestigateMoveService) {
	r = &InvestigateMoveService{}
	r.Options = opts
	return
}

// Move a message
func (r *InvestigateMoveService) New(ctx context.Context, postfixID string, params InvestigateMoveNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[InvestigateMoveNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/move", params.AccountID, postfixID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Move a message
func (r *InvestigateMoveService) NewAutoPaging(ctx context.Context, postfixID string, params InvestigateMoveNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InvestigateMoveNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, postfixID, params, opts...))
}

// Move multiple messages
func (r *InvestigateMoveService) Bulk(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) (res *pagination.SinglePage[InvestigateMoveBulkResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/move", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Move multiple messages
func (r *InvestigateMoveService) BulkAutoPaging(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InvestigateMoveBulkResponse] {
	return pagination.NewSinglePageAutoPager(r.Bulk(ctx, params, opts...))
}

type InvestigateMoveNewResponse struct {
	CompletedTimestamp time.Time                      `json:"completed_timestamp,required" format:"date-time"`
	ItemCount          int64                          `json:"item_count,required"`
	Destination        string                         `json:"destination,nullable"`
	MessageID          string                         `json:"message_id,nullable"`
	Operation          string                         `json:"operation,nullable"`
	Recipient          string                         `json:"recipient,nullable"`
	Status             string                         `json:"status,nullable"`
	JSON               investigateMoveNewResponseJSON `json:"-"`
}

// investigateMoveNewResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveNewResponse]
type investigateMoveNewResponseJSON struct {
	CompletedTimestamp apijson.Field
	ItemCount          apijson.Field
	Destination        apijson.Field
	MessageID          apijson.Field
	Operation          apijson.Field
	Recipient          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateMoveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveBulkResponse struct {
	CompletedTimestamp time.Time                       `json:"completed_timestamp,required" format:"date-time"`
	ItemCount          int64                           `json:"item_count,required"`
	Destination        string                          `json:"destination,nullable"`
	MessageID          string                          `json:"message_id,nullable"`
	Operation          string                          `json:"operation,nullable"`
	Recipient          string                          `json:"recipient,nullable"`
	Status             string                          `json:"status,nullable"`
	JSON               investigateMoveBulkResponseJSON `json:"-"`
}

// investigateMoveBulkResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveBulkResponse]
type investigateMoveBulkResponseJSON struct {
	CompletedTimestamp apijson.Field
	ItemCount          apijson.Field
	Destination        apijson.Field
	MessageID          apijson.Field
	Operation          apijson.Field
	Recipient          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateMoveBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveBulkResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveNewParams struct {
	// Account Identifier
	AccountID   param.Field[string]                              `path:"account_id,required"`
	Destination param.Field[InvestigateMoveNewParamsDestination] `json:"destination,required"`
}

func (r InvestigateMoveNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateMoveNewParamsDestination string

const (
	InvestigateMoveNewParamsDestinationInbox                     InvestigateMoveNewParamsDestination = "Inbox"
	InvestigateMoveNewParamsDestinationJunkEmail                 InvestigateMoveNewParamsDestination = "JunkEmail"
	InvestigateMoveNewParamsDestinationDeletedItems              InvestigateMoveNewParamsDestination = "DeletedItems"
	InvestigateMoveNewParamsDestinationRecoverableItemsDeletions InvestigateMoveNewParamsDestination = "RecoverableItemsDeletions"
	InvestigateMoveNewParamsDestinationRecoverableItemsPurges    InvestigateMoveNewParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateMoveNewParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateMoveNewParamsDestinationInbox, InvestigateMoveNewParamsDestinationJunkEmail, InvestigateMoveNewParamsDestinationDeletedItems, InvestigateMoveNewParamsDestinationRecoverableItemsDeletions, InvestigateMoveNewParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateMoveBulkParams struct {
	// Account Identifier
	AccountID   param.Field[string]                               `path:"account_id,required"`
	Destination param.Field[InvestigateMoveBulkParamsDestination] `json:"destination,required"`
	PostfixIDs  param.Field[[]string]                             `json:"postfix_ids,required"`
}

func (r InvestigateMoveBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateMoveBulkParamsDestination string

const (
	InvestigateMoveBulkParamsDestinationInbox                     InvestigateMoveBulkParamsDestination = "Inbox"
	InvestigateMoveBulkParamsDestinationJunkEmail                 InvestigateMoveBulkParamsDestination = "JunkEmail"
	InvestigateMoveBulkParamsDestinationDeletedItems              InvestigateMoveBulkParamsDestination = "DeletedItems"
	InvestigateMoveBulkParamsDestinationRecoverableItemsDeletions InvestigateMoveBulkParamsDestination = "RecoverableItemsDeletions"
	InvestigateMoveBulkParamsDestinationRecoverableItemsPurges    InvestigateMoveBulkParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateMoveBulkParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateMoveBulkParamsDestinationInbox, InvestigateMoveBulkParamsDestinationJunkEmail, InvestigateMoveBulkParamsDestinationDeletedItems, InvestigateMoveBulkParamsDestinationRecoverableItemsDeletions, InvestigateMoveBulkParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}
