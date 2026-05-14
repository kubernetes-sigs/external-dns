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

// ThreatEventRawService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventRawService] method instead.
type ThreatEventRawService struct {
	Options []option.RequestOption
}

// NewThreatEventRawService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventRawService(opts ...option.RequestOption) (r *ThreatEventRawService) {
	r = &ThreatEventRawService{}
	r.Options = opts
	return
}

// Updates a raw event
func (r *ThreatEventRawService) Edit(ctx context.Context, eventID string, rawID string, params ThreatEventRawEditParams, opts ...option.RequestOption) (res *ThreatEventRawEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if rawID == "" {
		err = errors.New("missing required raw_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/raw/%s", params.AccountID, eventID, rawID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Reads data for a raw event
func (r *ThreatEventRawService) Get(ctx context.Context, eventID string, rawID string, query ThreatEventRawGetParams, opts ...option.RequestOption) (res *ThreatEventRawGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if rawID == "" {
		err = errors.New("missing required raw_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/raw/%s", query.AccountID, eventID, rawID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventRawEditResponse struct {
	ID   string                         `json:"id,required"`
	Data interface{}                    `json:"data,required"`
	JSON threatEventRawEditResponseJSON `json:"-"`
}

// threatEventRawEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventRawEditResponse]
type threatEventRawEditResponseJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRawEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRawEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRawGetResponse struct {
	ID        string                        `json:"id,required"`
	AccountID float64                       `json:"accountId,required"`
	Created   string                        `json:"created,required"`
	Data      interface{}                   `json:"data,required"`
	Source    string                        `json:"source,required"`
	TLP       string                        `json:"tlp,required"`
	JSON      threatEventRawGetResponseJSON `json:"-"`
}

// threatEventRawGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventRawGetResponse]
type threatEventRawGetResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Source      apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRawGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRawGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRawEditParams struct {
	// Account ID.
	AccountID param.Field[string]      `path:"account_id,required"`
	Data      param.Field[interface{}] `json:"data"`
	Source    param.Field[string]      `json:"source"`
	TLP       param.Field[string]      `json:"tlp"`
}

func (r ThreatEventRawEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventRawGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
