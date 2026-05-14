// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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
)

// InvestigateReleaseService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateReleaseService] method instead.
type InvestigateReleaseService struct {
	Options []option.RequestOption
}

// NewInvestigateReleaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateReleaseService(opts ...option.RequestOption) (r *InvestigateReleaseService) {
	r = &InvestigateReleaseService{}
	r.Options = opts
	return
}

// Release messages from quarantine
func (r *InvestigateReleaseService) Bulk(ctx context.Context, params InvestigateReleaseBulkParams, opts ...option.RequestOption) (res *pagination.SinglePage[InvestigateReleaseBulkResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/release", params.AccountID)
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

// Release messages from quarantine
func (r *InvestigateReleaseService) BulkAutoPaging(ctx context.Context, params InvestigateReleaseBulkParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InvestigateReleaseBulkResponse] {
	return pagination.NewSinglePageAutoPager(r.Bulk(ctx, params, opts...))
}

type InvestigateReleaseBulkResponse struct {
	// The identifier of the message.
	PostfixID   string                             `json:"postfix_id,required"`
	Delivered   []string                           `json:"delivered,nullable"`
	Failed      []string                           `json:"failed,nullable"`
	Undelivered []string                           `json:"undelivered,nullable"`
	JSON        investigateReleaseBulkResponseJSON `json:"-"`
}

// investigateReleaseBulkResponseJSON contains the JSON metadata for the struct
// [InvestigateReleaseBulkResponse]
type investigateReleaseBulkResponseJSON struct {
	PostfixID   apijson.Field
	Delivered   apijson.Field
	Failed      apijson.Field
	Undelivered apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateReleaseBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReleaseBulkResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateReleaseBulkParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of messages identfied by their `postfix_id`s that should be released.
	Body []string `json:"body,required"`
}

func (r InvestigateReleaseBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
