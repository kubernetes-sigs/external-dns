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

// ThreatEventCountryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventCountryService] method instead.
type ThreatEventCountryService struct {
	Options []option.RequestOption
}

// NewThreatEventCountryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventCountryService(opts ...option.RequestOption) (r *ThreatEventCountryService) {
	r = &ThreatEventCountryService{}
	r.Options = opts
	return
}

// Retrieves countries information for all countries
func (r *ThreatEventCountryService) List(ctx context.Context, query ThreatEventCountryListParams, opts ...option.RequestOption) (res *[]ThreatEventCountryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/countries", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventCountryListResponse struct {
	Result  []ThreatEventCountryListResponseResult `json:"result,required"`
	Success string                                 `json:"success,required"`
	JSON    threatEventCountryListResponseJSON     `json:"-"`
}

// threatEventCountryListResponseJSON contains the JSON metadata for the struct
// [ThreatEventCountryListResponse]
type threatEventCountryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCountryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCountryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCountryListResponseResult struct {
	Alpha3 string                                   `json:"alpha3,required"`
	Name   string                                   `json:"name,required"`
	JSON   threatEventCountryListResponseResultJSON `json:"-"`
}

// threatEventCountryListResponseResultJSON contains the JSON metadata for the
// struct [ThreatEventCountryListResponseResult]
type threatEventCountryListResponseResultJSON struct {
	Alpha3      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCountryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCountryListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCountryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
