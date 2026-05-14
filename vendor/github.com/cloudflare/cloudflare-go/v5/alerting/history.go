// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// HistoryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoryService] method instead.
type HistoryService struct {
	Options []option.RequestOption
}

// NewHistoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHistoryService(opts ...option.RequestOption) (r *HistoryService) {
	r = &HistoryService{}
	r.Options = opts
	return
}

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *HistoryService) List(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[History], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/history", params.AccountID)
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

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *HistoryService) ListAutoPaging(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[History] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type History struct {
	// UUID
	ID string `json:"id"`
	// Message body included in the notification sent.
	AlertBody string `json:"alert_body"`
	// Type of notification that has been dispatched.
	AlertType string `json:"alert_type"`
	// Description of the notification policy (if present).
	Description string `json:"description"`
	// The mechanism to which the notification has been dispatched.
	Mechanism string `json:"mechanism"`
	// The type of mechanism to which the notification has been dispatched. This can be
	// email/pagerduty/webhook based on the mechanism configured.
	MechanismType HistoryMechanismType `json:"mechanism_type"`
	// Name of the policy.
	Name string `json:"name"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// Timestamp of when the notification was dispatched in ISO 8601 format.
	Sent time.Time   `json:"sent" format:"date-time"`
	JSON historyJSON `json:"-"`
}

// historyJSON contains the JSON metadata for the struct [History]
type historyJSON struct {
	ID            apijson.Field
	AlertBody     apijson.Field
	AlertType     apijson.Field
	Description   apijson.Field
	Mechanism     apijson.Field
	MechanismType apijson.Field
	Name          apijson.Field
	PolicyID      apijson.Field
	Sent          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *History) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyJSON) RawJSON() string {
	return r.raw
}

// The type of mechanism to which the notification has been dispatched. This can be
// email/pagerduty/webhook based on the mechanism configured.
type HistoryMechanismType string

const (
	HistoryMechanismTypeEmail     HistoryMechanismType = "email"
	HistoryMechanismTypePagerduty HistoryMechanismType = "pagerduty"
	HistoryMechanismTypeWebhook   HistoryMechanismType = "webhook"
)

func (r HistoryMechanismType) IsKnown() bool {
	switch r {
	case HistoryMechanismTypeEmail, HistoryMechanismTypePagerduty, HistoryMechanismTypeWebhook:
		return true
	}
	return false
}

type HistoryListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Limit the returned results to history records older than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limit the returned results to history records newer than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
}

// URLQuery serializes [HistoryListParams]'s query parameters as `url.Values`.
func (r HistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
