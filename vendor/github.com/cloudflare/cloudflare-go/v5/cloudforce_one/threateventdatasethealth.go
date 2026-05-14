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

// ThreatEventDatasetHealthService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventDatasetHealthService] method instead.
type ThreatEventDatasetHealthService struct {
	Options []option.RequestOption
}

// NewThreatEventDatasetHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventDatasetHealthService(opts ...option.RequestOption) (r *ThreatEventDatasetHealthService) {
	r = &ThreatEventDatasetHealthService{}
	r.Options = opts
	return
}

// Benchmark Durable Object warmup
func (r *ThreatEventDatasetHealthService) Get(ctx context.Context, datasetID string, query ThreatEventDatasetHealthGetParams, opts ...option.RequestOption) (res *ThreatEventDatasetHealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/health", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventDatasetHealthGetResponse struct {
	Items ThreatEventDatasetHealthGetResponseItems `json:"items,required"`
	Type  string                                   `json:"type,required"`
	JSON  threatEventDatasetHealthGetResponseJSON  `json:"-"`
}

// threatEventDatasetHealthGetResponseJSON contains the JSON metadata for the
// struct [ThreatEventDatasetHealthGetResponse]
type threatEventDatasetHealthGetResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponseItems struct {
	Type string                                       `json:"type,required"`
	JSON threatEventDatasetHealthGetResponseItemsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponseItemsJSON contains the JSON metadata for the
// struct [ThreatEventDatasetHealthGetResponseItems]
type threatEventDatasetHealthGetResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
