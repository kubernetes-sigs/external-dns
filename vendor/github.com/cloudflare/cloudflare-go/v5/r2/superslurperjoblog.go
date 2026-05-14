// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// SuperSlurperJobLogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperSlurperJobLogService] method instead.
type SuperSlurperJobLogService struct {
	Options []option.RequestOption
}

// NewSuperSlurperJobLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperSlurperJobLogService(opts ...option.RequestOption) (r *SuperSlurperJobLogService) {
	r = &SuperSlurperJobLogService{}
	r.Options = opts
	return
}

// Get job logs
func (r *SuperSlurperJobLogService) List(ctx context.Context, jobID string, params SuperSlurperJobLogListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SuperSlurperJobLogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/logs", params.AccountID, jobID)
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

// Get job logs
func (r *SuperSlurperJobLogService) ListAutoPaging(ctx context.Context, jobID string, params SuperSlurperJobLogListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SuperSlurperJobLogListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, jobID, params, opts...))
}

type SuperSlurperJobLogListResponse struct {
	CreatedAt string                                `json:"createdAt"`
	Job       string                                `json:"job"`
	LogType   SuperSlurperJobLogListResponseLogType `json:"logType"`
	Message   string                                `json:"message,nullable"`
	ObjectKey string                                `json:"objectKey,nullable"`
	JSON      superSlurperJobLogListResponseJSON    `json:"-"`
}

// superSlurperJobLogListResponseJSON contains the JSON metadata for the struct
// [SuperSlurperJobLogListResponse]
type superSlurperJobLogListResponseJSON struct {
	CreatedAt   apijson.Field
	Job         apijson.Field
	LogType     apijson.Field
	Message     apijson.Field
	ObjectKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobLogListResponseJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperJobLogListResponseLogType string

const (
	SuperSlurperJobLogListResponseLogTypeMigrationStart                      SuperSlurperJobLogListResponseLogType = "migrationStart"
	SuperSlurperJobLogListResponseLogTypeMigrationComplete                   SuperSlurperJobLogListResponseLogType = "migrationComplete"
	SuperSlurperJobLogListResponseLogTypeMigrationAbort                      SuperSlurperJobLogListResponseLogType = "migrationAbort"
	SuperSlurperJobLogListResponseLogTypeMigrationError                      SuperSlurperJobLogListResponseLogType = "migrationError"
	SuperSlurperJobLogListResponseLogTypeMigrationPause                      SuperSlurperJobLogListResponseLogType = "migrationPause"
	SuperSlurperJobLogListResponseLogTypeMigrationResume                     SuperSlurperJobLogListResponseLogType = "migrationResume"
	SuperSlurperJobLogListResponseLogTypeMigrationErrorFailedContinuation    SuperSlurperJobLogListResponseLogType = "migrationErrorFailedContinuation"
	SuperSlurperJobLogListResponseLogTypeImportErrorRetryExhaustion          SuperSlurperJobLogListResponseLogType = "importErrorRetryExhaustion"
	SuperSlurperJobLogListResponseLogTypeImportSkippedStorageClass           SuperSlurperJobLogListResponseLogType = "importSkippedStorageClass"
	SuperSlurperJobLogListResponseLogTypeImportSkippedOversized              SuperSlurperJobLogListResponseLogType = "importSkippedOversized"
	SuperSlurperJobLogListResponseLogTypeImportSkippedEmptyObject            SuperSlurperJobLogListResponseLogType = "importSkippedEmptyObject"
	SuperSlurperJobLogListResponseLogTypeImportSkippedUnsupportedContentType SuperSlurperJobLogListResponseLogType = "importSkippedUnsupportedContentType"
	SuperSlurperJobLogListResponseLogTypeImportSkippedExcludedContentType    SuperSlurperJobLogListResponseLogType = "importSkippedExcludedContentType"
	SuperSlurperJobLogListResponseLogTypeImportSkippedInvalidMedia           SuperSlurperJobLogListResponseLogType = "importSkippedInvalidMedia"
	SuperSlurperJobLogListResponseLogTypeImportSkippedRequiresRetrieval      SuperSlurperJobLogListResponseLogType = "importSkippedRequiresRetrieval"
)

func (r SuperSlurperJobLogListResponseLogType) IsKnown() bool {
	switch r {
	case SuperSlurperJobLogListResponseLogTypeMigrationStart, SuperSlurperJobLogListResponseLogTypeMigrationComplete, SuperSlurperJobLogListResponseLogTypeMigrationAbort, SuperSlurperJobLogListResponseLogTypeMigrationError, SuperSlurperJobLogListResponseLogTypeMigrationPause, SuperSlurperJobLogListResponseLogTypeMigrationResume, SuperSlurperJobLogListResponseLogTypeMigrationErrorFailedContinuation, SuperSlurperJobLogListResponseLogTypeImportErrorRetryExhaustion, SuperSlurperJobLogListResponseLogTypeImportSkippedStorageClass, SuperSlurperJobLogListResponseLogTypeImportSkippedOversized, SuperSlurperJobLogListResponseLogTypeImportSkippedEmptyObject, SuperSlurperJobLogListResponseLogTypeImportSkippedUnsupportedContentType, SuperSlurperJobLogListResponseLogTypeImportSkippedExcludedContentType, SuperSlurperJobLogListResponseLogTypeImportSkippedInvalidMedia, SuperSlurperJobLogListResponseLogTypeImportSkippedRequiresRetrieval:
		return true
	}
	return false
}

type SuperSlurperJobLogListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Limit     param.Field[int64]  `query:"limit"`
	Offset    param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [SuperSlurperJobLogListParams]'s query parameters as
// `url.Values`.
func (r SuperSlurperJobLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
