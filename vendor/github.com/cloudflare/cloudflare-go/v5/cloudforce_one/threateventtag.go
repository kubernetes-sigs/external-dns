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

// ThreatEventTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagService] method instead.
type ThreatEventTagService struct {
	Options []option.RequestOption
}

// NewThreatEventTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventTagService(opts ...option.RequestOption) (r *ThreatEventTagService) {
	r = &ThreatEventTagService{}
	r.Options = opts
	return
}

// Creates a new tag
func (r *ThreatEventTagService) New(ctx context.Context, params ThreatEventTagNewParams, opts ...option.RequestOption) (res *ThreatEventTagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ThreatEventTagNewResponse struct {
	Name string                        `json:"name,required"`
	UUID string                        `json:"uuid,required"`
	JSON threatEventTagNewResponseJSON `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r ThreatEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
