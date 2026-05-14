// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ThreatEventCronService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventCronService] method instead.
type ThreatEventCronService struct {
	Options []option.RequestOption
}

// NewThreatEventCronService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventCronService(opts ...option.RequestOption) (r *ThreatEventCronService) {
	r = &ThreatEventCronService{}
	r.Options = opts
	return
}

// Reads the last cron update time
func (r *ThreatEventCronService) List(ctx context.Context, query ThreatEventCronListParams, opts ...option.RequestOption) (res *ThreatEventCronListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/cron", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reads the last cron update time
func (r *ThreatEventCronService) Edit(ctx context.Context, body ThreatEventCronEditParams, opts ...option.RequestOption) (res *ThreatEventCronEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/cron", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type ThreatEventCronListResponse struct {
	Update string                          `json:"update,required"`
	JSON   threatEventCronListResponseJSON `json:"-"`
}

// threatEventCronListResponseJSON contains the JSON metadata for the struct
// [ThreatEventCronListResponse]
type threatEventCronListResponseJSON struct {
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCronListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCronListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCronEditResponse struct {
	ID     float64                         `json:"id,required"`
	Update string                          `json:"update,required"`
	JSON   threatEventCronEditResponseJSON `json:"-"`
}

// threatEventCronEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventCronEditResponse]
type threatEventCronEditResponseJSON struct {
	ID          apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCronEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCronEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCronListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventCronEditParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
