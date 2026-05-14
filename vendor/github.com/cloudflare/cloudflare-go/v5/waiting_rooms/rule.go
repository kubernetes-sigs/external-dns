// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

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

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *RuleService) New(ctx context.Context, waitingRoomID string, params RuleNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[WaitingRoomRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", params.ZoneID, waitingRoomID)
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

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *RuleService) NewAutoPaging(ctx context.Context, waitingRoomID string, params RuleNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WaitingRoomRule] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, waitingRoomID, params, opts...))
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *RuleService) Update(ctx context.Context, waitingRoomID string, params RuleUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[WaitingRoomRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", params.ZoneID, waitingRoomID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *RuleService) UpdateAutoPaging(ctx context.Context, waitingRoomID string, params RuleUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WaitingRoomRule] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, waitingRoomID, params, opts...))
}

// Deletes a rule for a waiting room.
func (r *RuleService) Delete(ctx context.Context, waitingRoomID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[WaitingRoomRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", body.ZoneID, waitingRoomID, ruleID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes a rule for a waiting room.
func (r *RuleService) DeleteAutoPaging(ctx context.Context, waitingRoomID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WaitingRoomRule] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, waitingRoomID, ruleID, body, opts...))
}

// Patches a rule for a waiting room.
func (r *RuleService) Edit(ctx context.Context, waitingRoomID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *pagination.SinglePage[WaitingRoomRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", params.ZoneID, waitingRoomID, ruleID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPatch, path, params, &res, opts...)
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

// Patches a rule for a waiting room.
func (r *RuleService) EditAutoPaging(ctx context.Context, waitingRoomID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WaitingRoomRule] {
	return pagination.NewSinglePageAutoPager(r.Edit(ctx, waitingRoomID, ruleID, params, opts...))
}

// Lists rules for a waiting room.
func (r *RuleService) Get(ctx context.Context, waitingRoomID string, query RuleGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[WaitingRoomRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", query.ZoneID, waitingRoomID)
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

// Lists rules for a waiting room.
func (r *RuleService) GetAutoPaging(ctx context.Context, waitingRoomID string, query RuleGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WaitingRoomRule] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, waitingRoomID, query, opts...))
}

type WaitingRoomRule struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string              `json:"version"`
	JSON    waitingRoomRuleJSON `json:"-"`
}

// waitingRoomRuleJSON contains the JSON metadata for the struct [WaitingRoomRule]
type waitingRoomRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomRuleJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleAction string

const (
	WaitingRoomRuleActionBypassWaitingRoom WaitingRoomRuleAction = "bypass_waiting_room"
)

func (r WaitingRoomRuleAction) IsKnown() bool {
	switch r {
	case WaitingRoomRuleActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Rules  RuleNewParamsRules  `json:"rules,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Rules)
}

type RuleNewParamsRules struct {
	// The action to take when the expression matches.
	Action param.Field[RuleNewParamsRulesAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleNewParamsRules) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type RuleNewParamsRulesAction string

const (
	RuleNewParamsRulesActionBypassWaitingRoom RuleNewParamsRulesAction = "bypass_waiting_room"
)

func (r RuleNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RuleNewParamsRulesActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]    `path:"zone_id,required"`
	Rules  []RuleUpdateParamsRule `json:"rules,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Rules)
}

type RuleUpdateParamsRule struct {
	// The action to take when the expression matches.
	Action param.Field[RuleUpdateParamsRulesAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type RuleUpdateParamsRulesAction string

const (
	RuleUpdateParamsRulesActionBypassWaitingRoom RuleUpdateParamsRulesAction = "bypass_waiting_room"
)

func (r RuleUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RuleUpdateParamsRulesActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take when the expression matches.
	Action param.Field[RuleEditParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[RuleEditParamsPositionUnion] `json:"position"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type RuleEditParamsAction string

const (
	RuleEditParamsActionBypassWaitingRoom RuleEditParamsAction = "bypass_waiting_room"
)

func (r RuleEditParamsAction) IsKnown() bool {
	switch r {
	case RuleEditParamsActionBypassWaitingRoom:
		return true
	}
	return false
}

// Reorder the position of a rule
type RuleEditParamsPosition struct {
	// Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the last rule in the ruleset.
	After param.Field[string] `json:"after"`
	// Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the first rule in the ruleset.
	Before param.Field[string] `json:"before"`
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r RuleEditParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPosition) implementsRuleEditParamsPositionUnion() {}

// Reorder the position of a rule
//
// Satisfied by [waiting_rooms.RuleEditParamsPositionIndex],
// [waiting_rooms.RuleEditParamsPositionBefore],
// [waiting_rooms.RuleEditParamsPositionAfter], [RuleEditParamsPosition].
type RuleEditParamsPositionUnion interface {
	implementsRuleEditParamsPositionUnion()
}

type RuleEditParamsPositionIndex struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r RuleEditParamsPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionIndex) implementsRuleEditParamsPositionUnion() {}

type RuleEditParamsPositionBefore struct {
	// Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the first rule in the ruleset.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionBefore) implementsRuleEditParamsPositionUnion() {}

type RuleEditParamsPositionAfter struct {
	// Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the last rule in the ruleset.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAfter) implementsRuleEditParamsPositionUnion() {}

type RuleGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
