// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

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

// PrefixDelegationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixDelegationService] method instead.
type PrefixDelegationService struct {
	Options []option.RequestOption
}

// NewPrefixDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrefixDelegationService(opts ...option.RequestOption) (r *PrefixDelegationService) {
	r = &PrefixDelegationService{}
	r.Options = opts
	return
}

// Create a new account delegation for a given IP prefix.
func (r *PrefixDelegationService) New(ctx context.Context, prefixID string, params PrefixDelegationNewParams, opts ...option.RequestOption) (res *Delegations, err error) {
	var env PrefixDelegationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) List(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Delegations], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", query.AccountID, prefixID)
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

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Delegations] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete an account delegation for a given IP prefix.
func (r *PrefixDelegationService) Delete(ctx context.Context, prefixID string, delegationID string, body PrefixDelegationDeleteParams, opts ...option.RequestOption) (res *PrefixDelegationDeleteResponse, err error) {
	var env PrefixDelegationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if delegationID == "" {
		err = errors.New("missing required delegation_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", body.AccountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Delegations struct {
	// Identifier of a Delegation.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier of an IP Prefix.
	ParentPrefixID string          `json:"parent_prefix_id"`
	JSON           delegationsJSON `json:"-"`
}

// delegationsJSON contains the JSON metadata for the struct [Delegations]
type delegationsJSON struct {
	ID                 apijson.Field
	CIDR               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Delegations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegationsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponse struct {
	// Identifier of a Delegation.
	ID   string                             `json:"id"`
	JSON prefixDelegationDeleteResponseJSON `json:"-"`
}

// prefixDelegationDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixDelegationDeleteResponse]
type prefixDelegationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r PrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixDelegationNewResponseEnvelope struct {
	Errors   []PrefixDelegationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixDelegationNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixDelegationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Delegations                                `json:"result"`
	JSON    prefixDelegationNewResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationNewResponseEnvelope]
type prefixDelegationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           PrefixDelegationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixDelegationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixDelegationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixDelegationNewResponseEnvelopeErrors]
type prefixDelegationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    prefixDelegationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixDelegationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [PrefixDelegationNewResponseEnvelopeErrorsSource]
type prefixDelegationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           PrefixDelegationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixDelegationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixDelegationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixDelegationNewResponseEnvelopeMessages]
type prefixDelegationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    prefixDelegationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixDelegationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [PrefixDelegationNewResponseEnvelopeMessagesSource]
type prefixDelegationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixDelegationNewResponseEnvelopeSuccess bool

const (
	PrefixDelegationNewResponseEnvelopeSuccessTrue PrefixDelegationNewResponseEnvelopeSuccess = true
)

func (r PrefixDelegationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixDelegationListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteResponseEnvelope struct {
	Errors   []PrefixDelegationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixDelegationDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixDelegationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixDelegationDeleteResponse                `json:"result"`
	JSON    prefixDelegationDeleteResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationDeleteResponseEnvelope]
type prefixDelegationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           PrefixDelegationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixDelegationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PrefixDelegationDeleteResponseEnvelopeErrors]
type prefixDelegationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    prefixDelegationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [PrefixDelegationDeleteResponseEnvelopeErrorsSource]
type prefixDelegationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           PrefixDelegationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixDelegationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PrefixDelegationDeleteResponseEnvelopeMessages]
type prefixDelegationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    prefixDelegationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [PrefixDelegationDeleteResponseEnvelopeMessagesSource]
type prefixDelegationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixDelegationDeleteResponseEnvelopeSuccess bool

const (
	PrefixDelegationDeleteResponseEnvelopeSuccessTrue PrefixDelegationDeleteResponseEnvelopeSuccess = true
)

func (r PrefixDelegationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
