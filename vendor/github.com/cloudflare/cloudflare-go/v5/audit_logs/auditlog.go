// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// AuditLogService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r *AuditLogService) {
	r = &AuditLogService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for an account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *AuditLogService) List(ctx context.Context, params AuditLogListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.AuditLog], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/audit_logs", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Gets a list of audit logs for an account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *AuditLogService) ListAutoPaging(ctx context.Context, params AuditLogListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.AuditLog] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type AuditLogListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Finds a specific log by its ID.
	ID     param.Field[string]                   `query:"id"`
	Action param.Field[AuditLogListParamsAction] `query:"action"`
	Actor  param.Field[AuditLogListParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. A `full-date`
	// that conforms to RFC3339.
	Before param.Field[AuditLogListParamsBeforeUnion] `query:"before" format:"date"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[AuditLogListParamsDirection] `query:"direction"`
	// Indicates that this request is an export of logs in CSV format.
	Export param.Field[bool] `query:"export"`
	// Indicates whether or not to hide user level audit logs.
	HideUserLogs param.Field[bool] `query:"hide_user_logs"`
	// Defines which page of results to return.
	Page param.Field[float64] `query:"page"`
	// Sets the number of results to return per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limits the returned results to logs newer than the specified date. A `full-date`
	// that conforms to RFC3339.
	Since param.Field[AuditLogListParamsSinceUnion] `query:"since" format:"date"`
	Zone  param.Field[AuditLogListParamsZone]       `query:"zone"`
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AuditLogListParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AuditLogListParamsAction]'s query parameters as
// `url.Values`.
func (r AuditLogListParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AuditLogListParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [AuditLogListParamsActor]'s query parameters as
// `url.Values`.
func (r AuditLogListParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Limits the returned results to logs older than the specified date. A `full-date`
// that conforms to RFC3339.
//
// Satisfied by [shared.UnionTime], [shared.UnionTime].
type AuditLogListParamsBeforeUnion interface {
	ImplementsAuditLogListParamsBeforeUnion()
}

// Changes the direction of the chronological sorting.
type AuditLogListParamsDirection string

const (
	AuditLogListParamsDirectionDesc AuditLogListParamsDirection = "desc"
	AuditLogListParamsDirectionAsc  AuditLogListParamsDirection = "asc"
)

func (r AuditLogListParamsDirection) IsKnown() bool {
	switch r {
	case AuditLogListParamsDirectionDesc, AuditLogListParamsDirectionAsc:
		return true
	}
	return false
}

// Limits the returned results to logs newer than the specified date. A `full-date`
// that conforms to RFC3339.
//
// Satisfied by [shared.UnionTime], [shared.UnionTime].
type AuditLogListParamsSinceUnion interface {
	ImplementsAuditLogListParamsSinceUnion()
}

type AuditLogListParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [AuditLogListParamsZone]'s query parameters as `url.Values`.
func (r AuditLogListParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
