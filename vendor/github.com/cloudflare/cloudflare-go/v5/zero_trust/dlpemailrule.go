// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// DLPEmailRuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEmailRuleService] method instead.
type DLPEmailRuleService struct {
	Options []option.RequestOption
}

// NewDLPEmailRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPEmailRuleService(opts ...option.RequestOption) (r *DLPEmailRuleService) {
	r = &DLPEmailRuleService{}
	r.Options = opts
	return
}

// Create email scanner rule
func (r *DLPEmailRuleService) New(ctx context.Context, params DLPEmailRuleNewParams, opts ...option.RequestOption) (res *DLPEmailRuleNewResponse, err error) {
	var env DLPEmailRuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update email scanner rule
func (r *DLPEmailRuleService) Update(ctx context.Context, ruleID string, params DLPEmailRuleUpdateParams, opts ...option.RequestOption) (res *DLPEmailRuleUpdateResponse, err error) {
	var env DLPEmailRuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all email scanner rules for an account.
func (r *DLPEmailRuleService) List(ctx context.Context, query DLPEmailRuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEmailRuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", query.AccountID)
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

// Lists all email scanner rules for an account.
func (r *DLPEmailRuleService) ListAutoPaging(ctx context.Context, query DLPEmailRuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEmailRuleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete email scanner rule
func (r *DLPEmailRuleService) Delete(ctx context.Context, ruleID string, body DLPEmailRuleDeleteParams, opts ...option.RequestOption) (res *DLPEmailRuleDeleteResponse, err error) {
	var env DLPEmailRuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update email scanner rule priorities
func (r *DLPEmailRuleService) BulkEdit(ctx context.Context, params DLPEmailRuleBulkEditParams, opts ...option.RequestOption) (res *DLPEmailRuleBulkEditResponse, err error) {
	var env DLPEmailRuleBulkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an email scanner rule
func (r *DLPEmailRuleService) Get(ctx context.Context, ruleID string, query DLPEmailRuleGetParams, opts ...option.RequestOption) (res *DLPEmailRuleGetResponse, err error) {
	var env DLPEmailRuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEmailRuleNewResponse struct {
	Action DLPEmailRuleNewResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleNewResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled     bool                               `json:"enabled,required"`
	Name        string                             `json:"name,required"`
	Priority    int64                              `json:"priority,required"`
	RuleID      string                             `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                          `json:"updated_at,required" format:"date-time"`
	Description string                             `json:"description,nullable"`
	JSON        dlpEmailRuleNewResponseJSON        `json:"-"`
}

// dlpEmailRuleNewResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleNewResponse]
type dlpEmailRuleNewResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseAction struct {
	Action  DLPEmailRuleNewResponseActionAction `json:"action,required"`
	Message string                              `json:"message,nullable"`
	JSON    dlpEmailRuleNewResponseActionJSON   `json:"-"`
}

// dlpEmailRuleNewResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleNewResponseAction]
type dlpEmailRuleNewResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseActionAction string

const (
	DLPEmailRuleNewResponseActionActionBlock DLPEmailRuleNewResponseActionAction = "Block"
)

func (r DLPEmailRuleNewResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleNewResponseCondition struct {
	Operator DLPEmailRuleNewResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleNewResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleNewResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleNewResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleNewResponseConditionJSON contains the JSON metadata for the struct
// [DLPEmailRuleNewResponseCondition]
type dlpEmailRuleNewResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseConditionsOperator string

const (
	DLPEmailRuleNewResponseConditionsOperatorInList        DLPEmailRuleNewResponseConditionsOperator = "InList"
	DLPEmailRuleNewResponseConditionsOperatorNotInList     DLPEmailRuleNewResponseConditionsOperator = "NotInList"
	DLPEmailRuleNewResponseConditionsOperatorMatchRegex    DLPEmailRuleNewResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleNewResponseConditionsOperatorNotMatchRegex DLPEmailRuleNewResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleNewResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewResponseConditionsOperatorInList, DLPEmailRuleNewResponseConditionsOperatorNotInList, DLPEmailRuleNewResponseConditionsOperatorMatchRegex, DLPEmailRuleNewResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleNewResponseConditionsSelector string

const (
	DLPEmailRuleNewResponseConditionsSelectorRecipients  DLPEmailRuleNewResponseConditionsSelector = "Recipients"
	DLPEmailRuleNewResponseConditionsSelectorSender      DLPEmailRuleNewResponseConditionsSelector = "Sender"
	DLPEmailRuleNewResponseConditionsSelectorDLPProfiles DLPEmailRuleNewResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleNewResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewResponseConditionsSelectorRecipients, DLPEmailRuleNewResponseConditionsSelectorSender, DLPEmailRuleNewResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleNewResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleNewResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleNewResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleNewResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleNewResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleNewResponseConditionsValueArray []string

func (r DLPEmailRuleNewResponseConditionsValueArray) ImplementsDLPEmailRuleNewResponseConditionsValueUnion() {
}

type DLPEmailRuleUpdateResponse struct {
	Action DLPEmailRuleUpdateResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleUpdateResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled     bool                                  `json:"enabled,required"`
	Name        string                                `json:"name,required"`
	Priority    int64                                 `json:"priority,required"`
	RuleID      string                                `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	Description string                                `json:"description,nullable"`
	JSON        dlpEmailRuleUpdateResponseJSON        `json:"-"`
}

// dlpEmailRuleUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleUpdateResponse]
type dlpEmailRuleUpdateResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseAction struct {
	Action  DLPEmailRuleUpdateResponseActionAction `json:"action,required"`
	Message string                                 `json:"message,nullable"`
	JSON    dlpEmailRuleUpdateResponseActionJSON   `json:"-"`
}

// dlpEmailRuleUpdateResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleUpdateResponseAction]
type dlpEmailRuleUpdateResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseActionAction string

const (
	DLPEmailRuleUpdateResponseActionActionBlock DLPEmailRuleUpdateResponseActionAction = "Block"
)

func (r DLPEmailRuleUpdateResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleUpdateResponseCondition struct {
	Operator DLPEmailRuleUpdateResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleUpdateResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleUpdateResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleUpdateResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleUpdateResponseConditionJSON contains the JSON metadata for the
// struct [DLPEmailRuleUpdateResponseCondition]
type dlpEmailRuleUpdateResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseConditionsOperator string

const (
	DLPEmailRuleUpdateResponseConditionsOperatorInList        DLPEmailRuleUpdateResponseConditionsOperator = "InList"
	DLPEmailRuleUpdateResponseConditionsOperatorNotInList     DLPEmailRuleUpdateResponseConditionsOperator = "NotInList"
	DLPEmailRuleUpdateResponseConditionsOperatorMatchRegex    DLPEmailRuleUpdateResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleUpdateResponseConditionsOperatorNotMatchRegex DLPEmailRuleUpdateResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleUpdateResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateResponseConditionsOperatorInList, DLPEmailRuleUpdateResponseConditionsOperatorNotInList, DLPEmailRuleUpdateResponseConditionsOperatorMatchRegex, DLPEmailRuleUpdateResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleUpdateResponseConditionsSelector string

const (
	DLPEmailRuleUpdateResponseConditionsSelectorRecipients  DLPEmailRuleUpdateResponseConditionsSelector = "Recipients"
	DLPEmailRuleUpdateResponseConditionsSelectorSender      DLPEmailRuleUpdateResponseConditionsSelector = "Sender"
	DLPEmailRuleUpdateResponseConditionsSelectorDLPProfiles DLPEmailRuleUpdateResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleUpdateResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateResponseConditionsSelectorRecipients, DLPEmailRuleUpdateResponseConditionsSelectorSender, DLPEmailRuleUpdateResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleUpdateResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleUpdateResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleUpdateResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleUpdateResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleUpdateResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleUpdateResponseConditionsValueArray []string

func (r DLPEmailRuleUpdateResponseConditionsValueArray) ImplementsDLPEmailRuleUpdateResponseConditionsValueUnion() {
}

type DLPEmailRuleListResponse struct {
	Action DLPEmailRuleListResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleListResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled     bool                                `json:"enabled,required"`
	Name        string                              `json:"name,required"`
	Priority    int64                               `json:"priority,required"`
	RuleID      string                              `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Description string                              `json:"description,nullable"`
	JSON        dlpEmailRuleListResponseJSON        `json:"-"`
}

// dlpEmailRuleListResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleListResponse]
type dlpEmailRuleListResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleListResponseAction struct {
	Action  DLPEmailRuleListResponseActionAction `json:"action,required"`
	Message string                               `json:"message,nullable"`
	JSON    dlpEmailRuleListResponseActionJSON   `json:"-"`
}

// dlpEmailRuleListResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleListResponseAction]
type dlpEmailRuleListResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleListResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleListResponseActionAction string

const (
	DLPEmailRuleListResponseActionActionBlock DLPEmailRuleListResponseActionAction = "Block"
)

func (r DLPEmailRuleListResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleListResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleListResponseCondition struct {
	Operator DLPEmailRuleListResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleListResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleListResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleListResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleListResponseConditionJSON contains the JSON metadata for the struct
// [DLPEmailRuleListResponseCondition]
type dlpEmailRuleListResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleListResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleListResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleListResponseConditionsOperator string

const (
	DLPEmailRuleListResponseConditionsOperatorInList        DLPEmailRuleListResponseConditionsOperator = "InList"
	DLPEmailRuleListResponseConditionsOperatorNotInList     DLPEmailRuleListResponseConditionsOperator = "NotInList"
	DLPEmailRuleListResponseConditionsOperatorMatchRegex    DLPEmailRuleListResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleListResponseConditionsOperatorNotMatchRegex DLPEmailRuleListResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleListResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleListResponseConditionsOperatorInList, DLPEmailRuleListResponseConditionsOperatorNotInList, DLPEmailRuleListResponseConditionsOperatorMatchRegex, DLPEmailRuleListResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleListResponseConditionsSelector string

const (
	DLPEmailRuleListResponseConditionsSelectorRecipients  DLPEmailRuleListResponseConditionsSelector = "Recipients"
	DLPEmailRuleListResponseConditionsSelectorSender      DLPEmailRuleListResponseConditionsSelector = "Sender"
	DLPEmailRuleListResponseConditionsSelectorDLPProfiles DLPEmailRuleListResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleListResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleListResponseConditionsSelectorRecipients, DLPEmailRuleListResponseConditionsSelectorSender, DLPEmailRuleListResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleListResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleListResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleListResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleListResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleListResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleListResponseConditionsValueArray []string

func (r DLPEmailRuleListResponseConditionsValueArray) ImplementsDLPEmailRuleListResponseConditionsValueUnion() {
}

type DLPEmailRuleDeleteResponse struct {
	Action DLPEmailRuleDeleteResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleDeleteResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled     bool                                  `json:"enabled,required"`
	Name        string                                `json:"name,required"`
	Priority    int64                                 `json:"priority,required"`
	RuleID      string                                `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	Description string                                `json:"description,nullable"`
	JSON        dlpEmailRuleDeleteResponseJSON        `json:"-"`
}

// dlpEmailRuleDeleteResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleDeleteResponse]
type dlpEmailRuleDeleteResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseAction struct {
	Action  DLPEmailRuleDeleteResponseActionAction `json:"action,required"`
	Message string                                 `json:"message,nullable"`
	JSON    dlpEmailRuleDeleteResponseActionJSON   `json:"-"`
}

// dlpEmailRuleDeleteResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleDeleteResponseAction]
type dlpEmailRuleDeleteResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseActionAction string

const (
	DLPEmailRuleDeleteResponseActionActionBlock DLPEmailRuleDeleteResponseActionAction = "Block"
)

func (r DLPEmailRuleDeleteResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleDeleteResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleDeleteResponseCondition struct {
	Operator DLPEmailRuleDeleteResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleDeleteResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleDeleteResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleDeleteResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleDeleteResponseConditionJSON contains the JSON metadata for the
// struct [DLPEmailRuleDeleteResponseCondition]
type dlpEmailRuleDeleteResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseConditionsOperator string

const (
	DLPEmailRuleDeleteResponseConditionsOperatorInList        DLPEmailRuleDeleteResponseConditionsOperator = "InList"
	DLPEmailRuleDeleteResponseConditionsOperatorNotInList     DLPEmailRuleDeleteResponseConditionsOperator = "NotInList"
	DLPEmailRuleDeleteResponseConditionsOperatorMatchRegex    DLPEmailRuleDeleteResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleDeleteResponseConditionsOperatorNotMatchRegex DLPEmailRuleDeleteResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleDeleteResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleDeleteResponseConditionsOperatorInList, DLPEmailRuleDeleteResponseConditionsOperatorNotInList, DLPEmailRuleDeleteResponseConditionsOperatorMatchRegex, DLPEmailRuleDeleteResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleDeleteResponseConditionsSelector string

const (
	DLPEmailRuleDeleteResponseConditionsSelectorRecipients  DLPEmailRuleDeleteResponseConditionsSelector = "Recipients"
	DLPEmailRuleDeleteResponseConditionsSelectorSender      DLPEmailRuleDeleteResponseConditionsSelector = "Sender"
	DLPEmailRuleDeleteResponseConditionsSelectorDLPProfiles DLPEmailRuleDeleteResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleDeleteResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleDeleteResponseConditionsSelectorRecipients, DLPEmailRuleDeleteResponseConditionsSelectorSender, DLPEmailRuleDeleteResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleDeleteResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleDeleteResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleDeleteResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleDeleteResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleDeleteResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleDeleteResponseConditionsValueArray []string

func (r DLPEmailRuleDeleteResponseConditionsValueArray) ImplementsDLPEmailRuleDeleteResponseConditionsValueUnion() {
}

type DLPEmailRuleBulkEditResponse struct {
	Action DLPEmailRuleBulkEditResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleBulkEditResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled     bool                                    `json:"enabled,required"`
	Name        string                                  `json:"name,required"`
	Priority    int64                                   `json:"priority,required"`
	RuleID      string                                  `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                               `json:"updated_at,required" format:"date-time"`
	Description string                                  `json:"description,nullable"`
	JSON        dlpEmailRuleBulkEditResponseJSON        `json:"-"`
}

// dlpEmailRuleBulkEditResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleBulkEditResponse]
type dlpEmailRuleBulkEditResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseAction struct {
	Action  DLPEmailRuleBulkEditResponseActionAction `json:"action,required"`
	Message string                                   `json:"message,nullable"`
	JSON    dlpEmailRuleBulkEditResponseActionJSON   `json:"-"`
}

// dlpEmailRuleBulkEditResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleBulkEditResponseAction]
type dlpEmailRuleBulkEditResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseActionAction string

const (
	DLPEmailRuleBulkEditResponseActionActionBlock DLPEmailRuleBulkEditResponseActionAction = "Block"
)

func (r DLPEmailRuleBulkEditResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleBulkEditResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleBulkEditResponseCondition struct {
	Operator DLPEmailRuleBulkEditResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleBulkEditResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleBulkEditResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleBulkEditResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleBulkEditResponseConditionJSON contains the JSON metadata for the
// struct [DLPEmailRuleBulkEditResponseCondition]
type dlpEmailRuleBulkEditResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseConditionsOperator string

const (
	DLPEmailRuleBulkEditResponseConditionsOperatorInList        DLPEmailRuleBulkEditResponseConditionsOperator = "InList"
	DLPEmailRuleBulkEditResponseConditionsOperatorNotInList     DLPEmailRuleBulkEditResponseConditionsOperator = "NotInList"
	DLPEmailRuleBulkEditResponseConditionsOperatorMatchRegex    DLPEmailRuleBulkEditResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleBulkEditResponseConditionsOperatorNotMatchRegex DLPEmailRuleBulkEditResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleBulkEditResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleBulkEditResponseConditionsOperatorInList, DLPEmailRuleBulkEditResponseConditionsOperatorNotInList, DLPEmailRuleBulkEditResponseConditionsOperatorMatchRegex, DLPEmailRuleBulkEditResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleBulkEditResponseConditionsSelector string

const (
	DLPEmailRuleBulkEditResponseConditionsSelectorRecipients  DLPEmailRuleBulkEditResponseConditionsSelector = "Recipients"
	DLPEmailRuleBulkEditResponseConditionsSelectorSender      DLPEmailRuleBulkEditResponseConditionsSelector = "Sender"
	DLPEmailRuleBulkEditResponseConditionsSelectorDLPProfiles DLPEmailRuleBulkEditResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleBulkEditResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleBulkEditResponseConditionsSelectorRecipients, DLPEmailRuleBulkEditResponseConditionsSelectorSender, DLPEmailRuleBulkEditResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleBulkEditResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleBulkEditResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleBulkEditResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleBulkEditResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleBulkEditResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleBulkEditResponseConditionsValueArray []string

func (r DLPEmailRuleBulkEditResponseConditionsValueArray) ImplementsDLPEmailRuleBulkEditResponseConditionsValueUnion() {
}

type DLPEmailRuleGetResponse struct {
	Action DLPEmailRuleGetResponseAction `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  []DLPEmailRuleGetResponseCondition `json:"conditions,required"`
	CreatedAt   time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled     bool                               `json:"enabled,required"`
	Name        string                             `json:"name,required"`
	Priority    int64                              `json:"priority,required"`
	RuleID      string                             `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time                          `json:"updated_at,required" format:"date-time"`
	Description string                             `json:"description,nullable"`
	JSON        dlpEmailRuleGetResponseJSON        `json:"-"`
}

// dlpEmailRuleGetResponseJSON contains the JSON metadata for the struct
// [DLPEmailRuleGetResponse]
type dlpEmailRuleGetResponseJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseAction struct {
	Action  DLPEmailRuleGetResponseActionAction `json:"action,required"`
	Message string                              `json:"message,nullable"`
	JSON    dlpEmailRuleGetResponseActionJSON   `json:"-"`
}

// dlpEmailRuleGetResponseActionJSON contains the JSON metadata for the struct
// [DLPEmailRuleGetResponseAction]
type dlpEmailRuleGetResponseActionJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseActionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseActionAction string

const (
	DLPEmailRuleGetResponseActionActionBlock DLPEmailRuleGetResponseActionAction = "Block"
)

func (r DLPEmailRuleGetResponseActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleGetResponseActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleGetResponseCondition struct {
	Operator DLPEmailRuleGetResponseConditionsOperator   `json:"operator,required"`
	Selector DLPEmailRuleGetResponseConditionsSelector   `json:"selector,required"`
	Value    DLPEmailRuleGetResponseConditionsValueUnion `json:"value,required"`
	JSON     dlpEmailRuleGetResponseConditionJSON        `json:"-"`
}

// dlpEmailRuleGetResponseConditionJSON contains the JSON metadata for the struct
// [DLPEmailRuleGetResponseCondition]
type dlpEmailRuleGetResponseConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseConditionJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseConditionsOperator string

const (
	DLPEmailRuleGetResponseConditionsOperatorInList        DLPEmailRuleGetResponseConditionsOperator = "InList"
	DLPEmailRuleGetResponseConditionsOperatorNotInList     DLPEmailRuleGetResponseConditionsOperator = "NotInList"
	DLPEmailRuleGetResponseConditionsOperatorMatchRegex    DLPEmailRuleGetResponseConditionsOperator = "MatchRegex"
	DLPEmailRuleGetResponseConditionsOperatorNotMatchRegex DLPEmailRuleGetResponseConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleGetResponseConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleGetResponseConditionsOperatorInList, DLPEmailRuleGetResponseConditionsOperatorNotInList, DLPEmailRuleGetResponseConditionsOperatorMatchRegex, DLPEmailRuleGetResponseConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleGetResponseConditionsSelector string

const (
	DLPEmailRuleGetResponseConditionsSelectorRecipients  DLPEmailRuleGetResponseConditionsSelector = "Recipients"
	DLPEmailRuleGetResponseConditionsSelectorSender      DLPEmailRuleGetResponseConditionsSelector = "Sender"
	DLPEmailRuleGetResponseConditionsSelectorDLPProfiles DLPEmailRuleGetResponseConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleGetResponseConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleGetResponseConditionsSelectorRecipients, DLPEmailRuleGetResponseConditionsSelectorSender, DLPEmailRuleGetResponseConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Union satisfied by [DLPEmailRuleGetResponseConditionsValueArray] or
// [shared.UnionString].
type DLPEmailRuleGetResponseConditionsValueUnion interface {
	ImplementsDLPEmailRuleGetResponseConditionsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailRuleGetResponseConditionsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailRuleGetResponseConditionsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPEmailRuleGetResponseConditionsValueArray []string

func (r DLPEmailRuleGetResponseConditionsValueArray) ImplementsDLPEmailRuleGetResponseConditionsValueUnion() {
}

type DLPEmailRuleNewParams struct {
	AccountID param.Field[string]                      `path:"account_id,required"`
	Action    param.Field[DLPEmailRuleNewParamsAction] `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  param.Field[[]DLPEmailRuleNewParamsCondition] `json:"conditions,required"`
	Enabled     param.Field[bool]                             `json:"enabled,required"`
	Name        param.Field[string]                           `json:"name,required"`
	Description param.Field[string]                           `json:"description"`
}

func (r DLPEmailRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleNewParamsAction struct {
	Action  param.Field[DLPEmailRuleNewParamsActionAction] `json:"action,required"`
	Message param.Field[string]                            `json:"message"`
}

func (r DLPEmailRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleNewParamsActionAction string

const (
	DLPEmailRuleNewParamsActionActionBlock DLPEmailRuleNewParamsActionAction = "Block"
)

func (r DLPEmailRuleNewParamsActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewParamsActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleNewParamsCondition struct {
	Operator param.Field[DLPEmailRuleNewParamsConditionsOperator]   `json:"operator,required"`
	Selector param.Field[DLPEmailRuleNewParamsConditionsSelector]   `json:"selector,required"`
	Value    param.Field[DLPEmailRuleNewParamsConditionsValueUnion] `json:"value,required"`
}

func (r DLPEmailRuleNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleNewParamsConditionsOperator string

const (
	DLPEmailRuleNewParamsConditionsOperatorInList        DLPEmailRuleNewParamsConditionsOperator = "InList"
	DLPEmailRuleNewParamsConditionsOperatorNotInList     DLPEmailRuleNewParamsConditionsOperator = "NotInList"
	DLPEmailRuleNewParamsConditionsOperatorMatchRegex    DLPEmailRuleNewParamsConditionsOperator = "MatchRegex"
	DLPEmailRuleNewParamsConditionsOperatorNotMatchRegex DLPEmailRuleNewParamsConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleNewParamsConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewParamsConditionsOperatorInList, DLPEmailRuleNewParamsConditionsOperatorNotInList, DLPEmailRuleNewParamsConditionsOperatorMatchRegex, DLPEmailRuleNewParamsConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleNewParamsConditionsSelector string

const (
	DLPEmailRuleNewParamsConditionsSelectorRecipients  DLPEmailRuleNewParamsConditionsSelector = "Recipients"
	DLPEmailRuleNewParamsConditionsSelectorSender      DLPEmailRuleNewParamsConditionsSelector = "Sender"
	DLPEmailRuleNewParamsConditionsSelectorDLPProfiles DLPEmailRuleNewParamsConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleNewParamsConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewParamsConditionsSelectorRecipients, DLPEmailRuleNewParamsConditionsSelectorSender, DLPEmailRuleNewParamsConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Satisfied by [zero_trust.DLPEmailRuleNewParamsConditionsValueArray],
// [shared.UnionString].
type DLPEmailRuleNewParamsConditionsValueUnion interface {
	ImplementsDLPEmailRuleNewParamsConditionsValueUnion()
}

type DLPEmailRuleNewParamsConditionsValueArray []string

func (r DLPEmailRuleNewParamsConditionsValueArray) ImplementsDLPEmailRuleNewParamsConditionsValueUnion() {
}

type DLPEmailRuleNewResponseEnvelope struct {
	Errors   []DLPEmailRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailRuleNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailRuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailRuleNewResponse                `json:"result"`
	JSON    dlpEmailRuleNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEmailRuleNewResponseEnvelope]
type dlpEmailRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPEmailRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEmailRuleNewResponseEnvelopeErrors]
type dlpEmailRuleNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpEmailRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailRuleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEmailRuleNewResponseEnvelopeErrorsSource]
type dlpEmailRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPEmailRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEmailRuleNewResponseEnvelopeMessages]
type dlpEmailRuleNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpEmailRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailRuleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEmailRuleNewResponseEnvelopeMessagesSource]
type dlpEmailRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailRuleNewResponseEnvelopeSuccess bool

const (
	DLPEmailRuleNewResponseEnvelopeSuccessTrue DLPEmailRuleNewResponseEnvelopeSuccess = true
)

func (r DLPEmailRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEmailRuleUpdateParams struct {
	AccountID param.Field[string]                         `path:"account_id,required"`
	Action    param.Field[DLPEmailRuleUpdateParamsAction] `json:"action,required"`
	// Rule is triggered if all conditions match.
	Conditions  param.Field[[]DLPEmailRuleUpdateParamsCondition] `json:"conditions,required"`
	Enabled     param.Field[bool]                                `json:"enabled,required"`
	Name        param.Field[string]                              `json:"name,required"`
	Description param.Field[string]                              `json:"description"`
}

func (r DLPEmailRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleUpdateParamsAction struct {
	Action  param.Field[DLPEmailRuleUpdateParamsActionAction] `json:"action,required"`
	Message param.Field[string]                               `json:"message"`
}

func (r DLPEmailRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleUpdateParamsActionAction string

const (
	DLPEmailRuleUpdateParamsActionActionBlock DLPEmailRuleUpdateParamsActionAction = "Block"
)

func (r DLPEmailRuleUpdateParamsActionAction) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateParamsActionActionBlock:
		return true
	}
	return false
}

type DLPEmailRuleUpdateParamsCondition struct {
	Operator param.Field[DLPEmailRuleUpdateParamsConditionsOperator]   `json:"operator,required"`
	Selector param.Field[DLPEmailRuleUpdateParamsConditionsSelector]   `json:"selector,required"`
	Value    param.Field[DLPEmailRuleUpdateParamsConditionsValueUnion] `json:"value,required"`
}

func (r DLPEmailRuleUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleUpdateParamsConditionsOperator string

const (
	DLPEmailRuleUpdateParamsConditionsOperatorInList        DLPEmailRuleUpdateParamsConditionsOperator = "InList"
	DLPEmailRuleUpdateParamsConditionsOperatorNotInList     DLPEmailRuleUpdateParamsConditionsOperator = "NotInList"
	DLPEmailRuleUpdateParamsConditionsOperatorMatchRegex    DLPEmailRuleUpdateParamsConditionsOperator = "MatchRegex"
	DLPEmailRuleUpdateParamsConditionsOperatorNotMatchRegex DLPEmailRuleUpdateParamsConditionsOperator = "NotMatchRegex"
)

func (r DLPEmailRuleUpdateParamsConditionsOperator) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateParamsConditionsOperatorInList, DLPEmailRuleUpdateParamsConditionsOperatorNotInList, DLPEmailRuleUpdateParamsConditionsOperatorMatchRegex, DLPEmailRuleUpdateParamsConditionsOperatorNotMatchRegex:
		return true
	}
	return false
}

type DLPEmailRuleUpdateParamsConditionsSelector string

const (
	DLPEmailRuleUpdateParamsConditionsSelectorRecipients  DLPEmailRuleUpdateParamsConditionsSelector = "Recipients"
	DLPEmailRuleUpdateParamsConditionsSelectorSender      DLPEmailRuleUpdateParamsConditionsSelector = "Sender"
	DLPEmailRuleUpdateParamsConditionsSelectorDLPProfiles DLPEmailRuleUpdateParamsConditionsSelector = "DLPProfiles"
)

func (r DLPEmailRuleUpdateParamsConditionsSelector) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateParamsConditionsSelectorRecipients, DLPEmailRuleUpdateParamsConditionsSelectorSender, DLPEmailRuleUpdateParamsConditionsSelectorDLPProfiles:
		return true
	}
	return false
}

// Satisfied by [zero_trust.DLPEmailRuleUpdateParamsConditionsValueArray],
// [shared.UnionString].
type DLPEmailRuleUpdateParamsConditionsValueUnion interface {
	ImplementsDLPEmailRuleUpdateParamsConditionsValueUnion()
}

type DLPEmailRuleUpdateParamsConditionsValueArray []string

func (r DLPEmailRuleUpdateParamsConditionsValueArray) ImplementsDLPEmailRuleUpdateParamsConditionsValueUnion() {
}

type DLPEmailRuleUpdateResponseEnvelope struct {
	Errors   []DLPEmailRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailRuleUpdateResponse                `json:"result"`
	JSON    dlpEmailRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEmailRuleUpdateResponseEnvelope]
type dlpEmailRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPEmailRuleUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailRuleUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEmailRuleUpdateResponseEnvelopeErrors]
type dlpEmailRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpEmailRuleUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailRuleUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEmailRuleUpdateResponseEnvelopeErrorsSource]
type dlpEmailRuleUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPEmailRuleUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailRuleUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEmailRuleUpdateResponseEnvelopeMessages]
type dlpEmailRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpEmailRuleUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailRuleUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPEmailRuleUpdateResponseEnvelopeMessagesSource]
type dlpEmailRuleUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailRuleUpdateResponseEnvelopeSuccess bool

const (
	DLPEmailRuleUpdateResponseEnvelopeSuccessTrue DLPEmailRuleUpdateResponseEnvelopeSuccess = true
)

func (r DLPEmailRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEmailRuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEmailRuleDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEmailRuleDeleteResponseEnvelope struct {
	Errors   []DLPEmailRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailRuleDeleteResponse                `json:"result"`
	JSON    dlpEmailRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEmailRuleDeleteResponseEnvelope]
type dlpEmailRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPEmailRuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailRuleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEmailRuleDeleteResponseEnvelopeErrors]
type dlpEmailRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpEmailRuleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailRuleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEmailRuleDeleteResponseEnvelopeErrorsSource]
type dlpEmailRuleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPEmailRuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailRuleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEmailRuleDeleteResponseEnvelopeMessages]
type dlpEmailRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpEmailRuleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailRuleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPEmailRuleDeleteResponseEnvelopeMessagesSource]
type dlpEmailRuleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailRuleDeleteResponseEnvelopeSuccess bool

const (
	DLPEmailRuleDeleteResponseEnvelopeSuccessTrue DLPEmailRuleDeleteResponseEnvelopeSuccess = true
)

func (r DLPEmailRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEmailRuleBulkEditParams struct {
	AccountID     param.Field[string]           `path:"account_id,required"`
	NewPriorities param.Field[map[string]int64] `json:"new_priorities,required"`
}

func (r DLPEmailRuleBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailRuleBulkEditResponseEnvelope struct {
	Errors   []DLPEmailRuleBulkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailRuleBulkEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailRuleBulkEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailRuleBulkEditResponse                `json:"result"`
	JSON    dlpEmailRuleBulkEditResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailRuleBulkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEmailRuleBulkEditResponseEnvelope]
type dlpEmailRuleBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPEmailRuleBulkEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailRuleBulkEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailRuleBulkEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEmailRuleBulkEditResponseEnvelopeErrors]
type dlpEmailRuleBulkEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpEmailRuleBulkEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailRuleBulkEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEmailRuleBulkEditResponseEnvelopeErrorsSource]
type dlpEmailRuleBulkEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEmailRuleBulkEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailRuleBulkEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailRuleBulkEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEmailRuleBulkEditResponseEnvelopeMessages]
type dlpEmailRuleBulkEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleBulkEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEmailRuleBulkEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailRuleBulkEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEmailRuleBulkEditResponseEnvelopeMessagesSource]
type dlpEmailRuleBulkEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleBulkEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleBulkEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailRuleBulkEditResponseEnvelopeSuccess bool

const (
	DLPEmailRuleBulkEditResponseEnvelopeSuccessTrue DLPEmailRuleBulkEditResponseEnvelopeSuccess = true
)

func (r DLPEmailRuleBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailRuleBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEmailRuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEmailRuleGetResponseEnvelope struct {
	Errors   []DLPEmailRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailRuleGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailRuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailRuleGetResponse                `json:"result"`
	JSON    dlpEmailRuleGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEmailRuleGetResponseEnvelope]
type dlpEmailRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPEmailRuleGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailRuleGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEmailRuleGetResponseEnvelopeErrors]
type dlpEmailRuleGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpEmailRuleGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailRuleGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEmailRuleGetResponseEnvelopeErrorsSource]
type dlpEmailRuleGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPEmailRuleGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailRuleGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEmailRuleGetResponseEnvelopeMessages]
type dlpEmailRuleGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailRuleGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpEmailRuleGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailRuleGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEmailRuleGetResponseEnvelopeMessagesSource]
type dlpEmailRuleGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailRuleGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailRuleGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailRuleGetResponseEnvelopeSuccess bool

const (
	DLPEmailRuleGetResponseEnvelopeSuccessTrue DLPEmailRuleGetResponseEnvelopeSuccess = true
)

func (r DLPEmailRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
