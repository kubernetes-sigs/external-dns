// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

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

// ReceivedFieldService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReceivedFieldService] method instead.
type ReceivedFieldService struct {
	Options []option.RequestOption
}

// NewReceivedFieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReceivedFieldService(opts ...option.RequestOption) (r *ReceivedFieldService) {
	r = &ReceivedFieldService{}
	r.Options = opts
	return
}

// Lists all fields available. The response is json object with key-value pairs,
// where keys are field names, and values are descriptions.
func (r *ReceivedFieldService) Get(ctx context.Context, query ReceivedFieldGetParams, opts ...option.RequestOption) (res *ReceivedFieldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/received/fields", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ReceivedFieldGetResponse struct {
	Key  string                       `json:"key"`
	JSON receivedFieldGetResponseJSON `json:"-"`
}

// receivedFieldGetResponseJSON contains the JSON metadata for the struct
// [ReceivedFieldGetResponse]
type receivedFieldGetResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReceivedFieldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r receivedFieldGetResponseJSON) RawJSON() string {
	return r.raw
}

type ReceivedFieldGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
