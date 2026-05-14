// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// UserSchemaHostService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSchemaHostService] method instead.
type UserSchemaHostService struct {
	Options []option.RequestOption
}

// NewUserSchemaHostService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserSchemaHostService(opts ...option.RequestOption) (r *UserSchemaHostService) {
	r = &UserSchemaHostService{}
	r.Options = opts
	return
}

// Retrieve schema hosts in a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaHostService) List(ctx context.Context, params UserSchemaHostListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserSchemaHostListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/hosts", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve schema hosts in a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaHostService) ListAutoPaging(ctx context.Context, params UserSchemaHostListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserSchemaHostListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type UserSchemaHostListResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Hosts serving the schema, e.g zone.host.com
	Hosts []string `json:"hosts,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID.
	SchemaID string                         `json:"schema_id,required"`
	JSON     userSchemaHostListResponseJSON `json:"-"`
}

// userSchemaHostListResponseJSON contains the JSON metadata for the struct
// [UserSchemaHostListResponse]
type userSchemaHostListResponseJSON struct {
	CreatedAt   apijson.Field
	Hosts       apijson.Field
	Name        apijson.Field
	SchemaID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaHostListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaHostListResponseJSON) RawJSON() string {
	return r.raw
}

type UserSchemaHostListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [UserSchemaHostListParams]'s query parameters as
// `url.Values`.
func (r UserSchemaHostListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
