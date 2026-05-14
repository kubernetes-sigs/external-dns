// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// PCAPDownloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPCAPDownloadService] method instead.
type PCAPDownloadService struct {
	Options []option.RequestOption
}

// NewPCAPDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPCAPDownloadService(opts ...option.RequestOption) (r *PCAPDownloadService) {
	r = &PCAPDownloadService{}
	r.Options = opts
	return
}

// Download PCAP information into a file. Response is a binary PCAP file.
func (r *PCAPDownloadService) Get(ctx context.Context, pcapID string, query PCAPDownloadGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.tcpdump.pcap")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pcapID == "" {
		err = errors.New("missing required pcap_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/%s/download", query.AccountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PCAPDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
