// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AttackSurfaceReportIssueTypeService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackSurfaceReportIssueTypeService] method instead.
type AttackSurfaceReportIssueTypeService struct {
	Options []option.RequestOption
}

// NewAttackSurfaceReportIssueTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackSurfaceReportIssueTypeService(opts ...option.RequestOption) (r *AttackSurfaceReportIssueTypeService) {
	r = &AttackSurfaceReportIssueTypeService{}
	r.Options = opts
	return
}

// Get Security Center Issues Types
func (r *AttackSurfaceReportIssueTypeService) Get(ctx context.Context, query AttackSurfaceReportIssueTypeGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issue-types", query.AccountID)
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

// Get Security Center Issues Types
func (r *AttackSurfaceReportIssueTypeService) GetAutoPaging(ctx context.Context, query AttackSurfaceReportIssueTypeGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type AttackSurfaceReportIssueTypeGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
