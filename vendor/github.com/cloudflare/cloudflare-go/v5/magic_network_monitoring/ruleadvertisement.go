// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

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

// RuleAdvertisementService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleAdvertisementService] method instead.
type RuleAdvertisementService struct {
	Options []option.RequestOption
}

// NewRuleAdvertisementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRuleAdvertisementService(opts ...option.RequestOption) (r *RuleAdvertisementService) {
	r = &RuleAdvertisementService{}
	r.Options = opts
	return
}

// Update advertisement for rule.
func (r *RuleAdvertisementService) Edit(ctx context.Context, ruleID string, params RuleAdvertisementEditParams, opts ...option.RequestOption) (res *Advertisement, err error) {
	var env RuleAdvertisementEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s/advertisement", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Advertisement struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool              `json:"automatic_advertisement,required,nullable"`
	JSON                   advertisementJSON `json:"-"`
}

// advertisementJSON contains the JSON metadata for the struct [Advertisement]
type advertisementJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Advertisement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisementJSON) RawJSON() string {
	return r.raw
}

type RuleAdvertisementEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r RuleAdvertisementEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleAdvertisementEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Advertisement         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleAdvertisementEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleAdvertisementEditResponseEnvelopeJSON    `json:"-"`
}

// ruleAdvertisementEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [RuleAdvertisementEditResponseEnvelope]
type ruleAdvertisementEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleAdvertisementEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleAdvertisementEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleAdvertisementEditResponseEnvelopeSuccess bool

const (
	RuleAdvertisementEditResponseEnvelopeSuccessTrue RuleAdvertisementEditResponseEnvelopeSuccess = true
)

func (r RuleAdvertisementEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleAdvertisementEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
