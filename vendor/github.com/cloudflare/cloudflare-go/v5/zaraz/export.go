// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ExportService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExportService] method instead.
type ExportService struct {
	Options []option.RequestOption
}

// NewExportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExportService(opts ...option.RequestOption) (r *ExportService) {
	r = &ExportService{}
	r.Options = opts
	return
}

// Exports full current published Zaraz configuration for a zone, secret variables
// included.
func (r *ExportService) Get(ctx context.Context, query ExportGetParams, opts ...option.RequestOption) (res *Configuration, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/export", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExportGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
