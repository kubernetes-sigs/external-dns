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

// ScanResultService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanResultService] method instead.
type ScanResultService struct {
	Options []option.RequestOption
}

// NewScanResultService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScanResultService(opts ...option.RequestOption) (r *ScanResultService) {
	r = &ScanResultService{}
	r.Options = opts
	return
}

// Get the Latest Scan Result
func (r *ScanResultService) Get(ctx context.Context, configID string, query ScanResultGetParams, opts ...option.RequestOption) (res *ScanResultGetResponse, err error) {
	var env ScanResultGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/results/%s", query.AccountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScanResult struct {
	Number float64        `json:"number"`
	Proto  string         `json:"proto"`
	Status string         `json:"status"`
	JSON   scanResultJSON `json:"-"`
}

// scanResultJSON contains the JSON metadata for the struct [ScanResult]
type scanResultJSON struct {
	Number      apijson.Field
	Proto       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultJSON) RawJSON() string {
	return r.raw
}

type ScanResultGetResponse struct {
	OneOneOneOne []ScanResult              `json:"1.1.1.1,required"`
	JSON         scanResultGetResponseJSON `json:"-"`
}

// scanResultGetResponseJSON contains the JSON metadata for the struct
// [ScanResultGetResponse]
type scanResultGetResponseJSON struct {
	OneOneOneOne apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanResultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScanResultGetParams struct {
	// Defines the Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanResultGetResponseEnvelope struct {
	Errors   []string                          `json:"errors,required"`
	Messages []string                          `json:"messages,required"`
	Result   ScanResultGetResponse             `json:"result,required"`
	Success  bool                              `json:"success,required"`
	JSON     scanResultGetResponseEnvelopeJSON `json:"-"`
}

// scanResultGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanResultGetResponseEnvelope]
type scanResultGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanResultGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
