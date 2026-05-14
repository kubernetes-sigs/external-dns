// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// HyperdriveService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHyperdriveService] method instead.
type HyperdriveService struct {
	Options []option.RequestOption
	Configs *ConfigService
}

// NewHyperdriveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHyperdriveService(opts ...option.RequestOption) (r *HyperdriveService) {
	r = &HyperdriveService{}
	r.Options = opts
	r.Configs = NewConfigService(opts...)
	return
}

type Hyperdrive struct {
	// Define configurations using a unique string identifier.
	ID      string            `json:"id,required"`
	Name    string            `json:"name,required"`
	Origin  HyperdriveOrigin  `json:"origin,required"`
	Caching HyperdriveCaching `json:"caching"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time      `json:"modified_on" format:"date-time"`
	MTLS       HyperdriveMTLS `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	OriginConnectionLimit int64          `json:"origin_connection_limit"`
	JSON                  hyperdriveJSON `json:"-"`
}

// hyperdriveJSON contains the JSON metadata for the struct [Hyperdrive]
type hyperdriveJSON struct {
	ID                    apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	Caching               apijson.Field
	CreatedOn             apijson.Field
	ModifiedOn            apijson.Field
	MTLS                  apijson.Field
	OriginConnectionLimit apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Hyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveJSON) RawJSON() string {
	return r.raw
}

type HyperdriveOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string `json:"user,required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port  int64                `json:"port"`
	JSON  hyperdriveOriginJSON `json:"-"`
	union HyperdriveOriginUnion
}

// hyperdriveOriginJSON contains the JSON metadata for the struct
// [HyperdriveOrigin]
type hyperdriveOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r hyperdriveOriginJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveOriginUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [HyperdriveOriginPublicDatabase],
// [HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r HyperdriveOrigin) AsUnion() HyperdriveOriginUnion {
	return r.union
}

// Union satisfied by [HyperdriveOriginPublicDatabase] or
// [HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type HyperdriveOriginUnion interface {
	implementsHyperdriveOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
	)
}

type HyperdriveOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginPublicDatabaseScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string                             `json:"user,required"`
	JSON hyperdriveOriginPublicDatabaseJSON `json:"-"`
}

// hyperdriveOriginPublicDatabaseJSON contains the JSON metadata for the struct
// [HyperdriveOriginPublicDatabase]
type hyperdriveOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveOriginPublicDatabase) implementsHyperdriveOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginPublicDatabaseScheme string

const (
	HyperdriveOriginPublicDatabaseSchemePostgres   HyperdriveOriginPublicDatabaseScheme = "postgres"
	HyperdriveOriginPublicDatabaseSchemePostgresql HyperdriveOriginPublicDatabaseScheme = "postgresql"
	HyperdriveOriginPublicDatabaseSchemeMysql      HyperdriveOriginPublicDatabaseScheme = "mysql"
)

func (r HyperdriveOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginPublicDatabaseSchemePostgres, HyperdriveOriginPublicDatabaseSchemePostgresql, HyperdriveOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id,required"`
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string                                                            `json:"user,required"`
	JSON hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON contains the
// JSON metadata for the struct
// [HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveOrigin() {}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveOriginScheme string

const (
	HyperdriveOriginSchemePostgres   HyperdriveOriginScheme = "postgres"
	HyperdriveOriginSchemePostgresql HyperdriveOriginScheme = "postgresql"
	HyperdriveOriginSchemeMysql      HyperdriveOriginScheme = "mysql"
)

func (r HyperdriveOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveOriginSchemePostgres, HyperdriveOriginSchemePostgresql, HyperdriveOriginSchemeMysql:
		return true
	}
	return false
}

type HyperdriveCaching struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge int64 `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate int64                 `json:"stale_while_revalidate"`
	JSON                 hyperdriveCachingJSON `json:"-"`
	union                HyperdriveCachingUnion
}

// hyperdriveCachingJSON contains the JSON metadata for the struct
// [HyperdriveCaching]
type hyperdriveCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r hyperdriveCachingJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveCaching) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveCachingUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [HyperdriveCachingHyperdriveHyperdriveCachingCommon],
// [HyperdriveCachingHyperdriveHyperdriveCachingEnabled].
func (r HyperdriveCaching) AsUnion() HyperdriveCachingUnion {
	return r.union
}

// Union satisfied by [HyperdriveCachingHyperdriveHyperdriveCachingCommon] or
// [HyperdriveCachingHyperdriveHyperdriveCachingEnabled].
type HyperdriveCachingUnion interface {
	implementsHyperdriveCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveCachingHyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type HyperdriveCachingHyperdriveHyperdriveCachingCommon struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool                                                   `json:"disabled"`
	JSON     hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON contains the JSON
// metadata for the struct [HyperdriveCachingHyperdriveHyperdriveCachingCommon]
type hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveCachingHyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveCachingHyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommon) implementsHyperdriveCaching() {}

type HyperdriveCachingHyperdriveHyperdriveCachingEnabled struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge int64 `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate int64                                                   `json:"stale_while_revalidate"`
	JSON                 hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON contains the JSON
// metadata for the struct [HyperdriveCachingHyperdriveHyperdriveCachingEnabled]
type hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveCaching() {}

type HyperdriveMTLS struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID string `json:"mtls_certificate_id"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	Sslmode string             `json:"sslmode"`
	JSON    hyperdriveMTLSJSON `json:"-"`
}

// hyperdriveMTLSJSON contains the JSON metadata for the struct [HyperdriveMTLS]
type hyperdriveMTLSJSON struct {
	CACertificateID   apijson.Field
	MTLSCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HyperdriveMTLS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveMTLSJSON) RawJSON() string {
	return r.raw
}

type HyperdriveParam struct {
	Name    param.Field[string]                      `json:"name,required"`
	Origin  param.Field[HyperdriveOriginUnionParam]  `json:"origin,required"`
	Caching param.Field[HyperdriveCachingUnionParam] `json:"caching"`
	MTLS    param.Field[HyperdriveMTLSParam]         `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r HyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveOriginParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
}

func (r HyperdriveOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginParam) implementsHyperdriveOriginUnionParam() {}

// Satisfied by [hyperdrive.HyperdriveOriginPublicDatabaseParam],
// [hyperdrive.HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam],
// [HyperdriveOriginParam].
type HyperdriveOriginUnionParam interface {
	implementsHyperdriveOriginUnionParam()
}

type HyperdriveOriginPublicDatabaseParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginPublicDatabaseScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveOriginPublicDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginPublicDatabaseParam) implementsHyperdriveOriginUnionParam() {}

type HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) implementsHyperdriveOriginUnionParam() {
}

type HyperdriveCachingParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingParam) implementsHyperdriveCachingUnionParam() {}

// Satisfied by
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingCommonParam],
// [hyperdrive.HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam],
// [HyperdriveCachingParam].
type HyperdriveCachingUnionParam interface {
	implementsHyperdriveCachingUnionParam()
}

type HyperdriveCachingHyperdriveHyperdriveCachingCommonParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingCommonParam) implementsHyperdriveCachingUnionParam() {
}

type HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) implementsHyperdriveCachingUnionParam() {
}

type HyperdriveMTLSParam struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CACertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r HyperdriveMTLSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
