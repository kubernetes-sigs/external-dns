// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// GatewayAppTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayAppTypeService] method instead.
type GatewayAppTypeService struct {
	Options []option.RequestOption
}

// NewGatewayAppTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayAppTypeService(opts ...option.RequestOption) (r *GatewayAppTypeService) {
	r = &GatewayAppTypeService{}
	r.Options = opts
	return
}

// Fetches all application and application type mappings.
func (r *GatewayAppTypeService) List(ctx context.Context, query GatewayAppTypeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AppType], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/app_types", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetches all application and application type mappings.
func (r *GatewayAppTypeService) ListAutoPaging(ctx context.Context, query GatewayAppTypeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AppType] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type AppType struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name  string      `json:"name"`
	JSON  appTypeJSON `json:"-"`
	union AppTypeUnion
}

// appTypeJSON contains the JSON metadata for the struct [AppType]
type appTypeJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r appTypeJSON) RawJSON() string {
	return r.raw
}

func (r *AppType) UnmarshalJSON(data []byte) (err error) {
	*r = AppType{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppTypeUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [AppTypeZeroTrustGatewayApplication],
// [AppTypeZeroTrustGatewayApplicationType].
func (r AppType) AsUnion() AppTypeUnion {
	return r.union
}

// Union satisfied by [AppTypeZeroTrustGatewayApplication] or
// [AppTypeZeroTrustGatewayApplicationType].
type AppTypeUnion interface {
	implementsAppType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppTypeZeroTrustGatewayApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppTypeZeroTrustGatewayApplicationType{}),
		},
	)
}

type AppTypeZeroTrustGatewayApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                 `json:"name"`
	JSON appTypeZeroTrustGatewayApplicationJSON `json:"-"`
}

// appTypeZeroTrustGatewayApplicationJSON contains the JSON metadata for the struct
// [AppTypeZeroTrustGatewayApplication]
type appTypeZeroTrustGatewayApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AppTypeZeroTrustGatewayApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appTypeZeroTrustGatewayApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppTypeZeroTrustGatewayApplication) implementsAppType() {}

type AppTypeZeroTrustGatewayApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                     `json:"name"`
	JSON appTypeZeroTrustGatewayApplicationTypeJSON `json:"-"`
}

// appTypeZeroTrustGatewayApplicationTypeJSON contains the JSON metadata for the
// struct [AppTypeZeroTrustGatewayApplicationType]
type appTypeZeroTrustGatewayApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppTypeZeroTrustGatewayApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appTypeZeroTrustGatewayApplicationTypeJSON) RawJSON() string {
	return r.raw
}

func (r AppTypeZeroTrustGatewayApplicationType) implementsAppType() {}

type GatewayAppTypeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
