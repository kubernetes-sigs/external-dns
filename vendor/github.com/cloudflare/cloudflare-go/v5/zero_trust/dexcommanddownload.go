// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DEXCommandDownloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandDownloadService] method instead.
type DEXCommandDownloadService struct {
	Options []option.RequestOption
}

// NewDEXCommandDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXCommandDownloadService(opts ...option.RequestOption) (r *DEXCommandDownloadService) {
	r = &DEXCommandDownloadService{}
	r.Options = opts
	return
}

// Downloads artifacts for an executed command. Bulk downloads are not supported
func (r *DEXCommandDownloadService) Get(ctx context.Context, commandID string, filename string, query DEXCommandDownloadGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if commandID == "" {
		err = errors.New("missing required command_id parameter")
		return
	}
	if filename == "" {
		err = errors.New("missing required filename parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/%s/downloads/%s", query.AccountID, commandID, filename)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DEXCommandDownloadGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
