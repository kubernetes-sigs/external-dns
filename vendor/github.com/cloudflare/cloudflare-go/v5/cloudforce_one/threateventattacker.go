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

// ThreatEventAttackerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventAttackerService] method instead.
type ThreatEventAttackerService struct {
	Options []option.RequestOption
}

// NewThreatEventAttackerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventAttackerService(opts ...option.RequestOption) (r *ThreatEventAttackerService) {
	r = &ThreatEventAttackerService{}
	r.Options = opts
	return
}

// Lists attackers
func (r *ThreatEventAttackerService) List(ctx context.Context, query ThreatEventAttackerListParams, opts ...option.RequestOption) (res *ThreatEventAttackerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/attackers", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventAttackerListResponse struct {
	Items ThreatEventAttackerListResponseItems `json:"items,required"`
	Type  string                               `json:"type,required"`
	JSON  threatEventAttackerListResponseJSON  `json:"-"`
}

// threatEventAttackerListResponseJSON contains the JSON metadata for the struct
// [ThreatEventAttackerListResponse]
type threatEventAttackerListResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventAttackerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventAttackerListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventAttackerListResponseItems struct {
	Type string                                   `json:"type,required"`
	JSON threatEventAttackerListResponseItemsJSON `json:"-"`
}

// threatEventAttackerListResponseItemsJSON contains the JSON metadata for the
// struct [ThreatEventAttackerListResponseItems]
type threatEventAttackerListResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventAttackerListResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventAttackerListResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventAttackerListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
