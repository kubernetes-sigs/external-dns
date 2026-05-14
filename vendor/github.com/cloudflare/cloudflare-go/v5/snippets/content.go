// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ContentService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentService] method instead.
type ContentService struct {
	Options []option.RequestOption
}

// NewContentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContentService(opts ...option.RequestOption) (r *ContentService) {
	r = &ContentService{}
	r.Options = opts
	return
}

// Fetches the content of a snippet belonging to the zone.
func (r *ContentService) Get(ctx context.Context, snippetName string, query ContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/form-data")}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s/content", query.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ContentGetParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
