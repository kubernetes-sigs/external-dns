// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// CatalogSyncService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogSyncService] method instead.
type CatalogSyncService struct {
	Options          []option.RequestOption
	PrebuiltPolicies *CatalogSyncPrebuiltPolicyService
}

// NewCatalogSyncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogSyncService(opts ...option.RequestOption) (r *CatalogSyncService) {
	r = &CatalogSyncService{}
	r.Options = opts
	r.PrebuiltPolicies = NewCatalogSyncPrebuiltPolicyService(opts...)
	return
}

// Create a new Catalog Sync (Closed Beta).
func (r *CatalogSyncService) New(ctx context.Context, params CatalogSyncNewParams, opts ...option.RequestOption) (res *CatalogSyncNewResponse, err error) {
	var env CatalogSyncNewResponseEnvelope
	if params.Forwarded.Present {
		opts = append(opts, option.WithHeader("forwarded", fmt.Sprintf("%s", params.Forwarded)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Catalog Sync (Closed Beta).
func (r *CatalogSyncService) Update(ctx context.Context, syncID string, params CatalogSyncUpdateParams, opts ...option.RequestOption) (res *CatalogSyncUpdateResponse, err error) {
	var env CatalogSyncUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", params.AccountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Catalog Syncs (Closed Beta).
func (r *CatalogSyncService) List(ctx context.Context, query CatalogSyncListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CatalogSyncListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs", query.AccountID)
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

// List Catalog Syncs (Closed Beta).
func (r *CatalogSyncService) ListAutoPaging(ctx context.Context, query CatalogSyncListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CatalogSyncListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Catalog Sync (Closed Beta).
func (r *CatalogSyncService) Delete(ctx context.Context, syncID string, params CatalogSyncDeleteParams, opts ...option.RequestOption) (res *CatalogSyncDeleteResponse, err error) {
	var env CatalogSyncDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", params.AccountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Catalog Sync (Closed Beta).
func (r *CatalogSyncService) Edit(ctx context.Context, syncID string, params CatalogSyncEditParams, opts ...option.RequestOption) (res *CatalogSyncEditResponse, err error) {
	var env CatalogSyncEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", params.AccountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read a Catalog Sync (Closed Beta).
func (r *CatalogSyncService) Get(ctx context.Context, syncID string, query CatalogSyncGetParams, opts ...option.RequestOption) (res *CatalogSyncGetResponse, err error) {
	var env CatalogSyncGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", query.AccountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Refresh a Catalog Sync's destination by running the sync policy against latest
// resource catalog (Closed Beta).
func (r *CatalogSyncService) Refresh(ctx context.Context, syncID string, body CatalogSyncRefreshParams, opts ...option.RequestOption) (res *string, err error) {
	var env CatalogSyncRefreshResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s/refresh", body.AccountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CatalogSyncNewResponse struct {
	ID                       string                                 `json:"id,required" format:"uuid"`
	Description              string                                 `json:"description,required"`
	DestinationID            string                                 `json:"destination_id,required" format:"uuid"`
	DestinationType          CatalogSyncNewResponseDestinationType  `json:"destination_type,required"`
	LastUserUpdateAt         string                                 `json:"last_user_update_at,required"`
	Name                     string                                 `json:"name,required"`
	Policy                   string                                 `json:"policy,required"`
	UpdateMode               CatalogSyncNewResponseUpdateMode       `json:"update_mode,required"`
	Errors                   map[string]CatalogSyncNewResponseError `json:"errors"`
	IncludesDiscoveriesUntil string                                 `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                                 `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                                 `json:"last_successful_update_at"`
	JSON                     catalogSyncNewResponseJSON             `json:"-"`
}

// catalogSyncNewResponseJSON contains the JSON metadata for the struct
// [CatalogSyncNewResponse]
type catalogSyncNewResponseJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CatalogSyncNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseDestinationType string

const (
	CatalogSyncNewResponseDestinationTypeNone          CatalogSyncNewResponseDestinationType = "NONE"
	CatalogSyncNewResponseDestinationTypeZeroTrustList CatalogSyncNewResponseDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncNewResponseDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncNewResponseDestinationTypeNone, CatalogSyncNewResponseDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncNewResponseUpdateMode string

const (
	CatalogSyncNewResponseUpdateModeAuto   CatalogSyncNewResponseUpdateMode = "AUTO"
	CatalogSyncNewResponseUpdateModeManual CatalogSyncNewResponseUpdateMode = "MANUAL"
)

func (r CatalogSyncNewResponseUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncNewResponseUpdateModeAuto, CatalogSyncNewResponseUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncNewResponseError struct {
	Code             CatalogSyncNewResponseErrorsCode   `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Meta             CatalogSyncNewResponseErrorsMeta   `json:"meta"`
	Source           CatalogSyncNewResponseErrorsSource `json:"source"`
	JSON             catalogSyncNewResponseErrorJSON    `json:"-"`
}

// catalogSyncNewResponseErrorJSON contains the JSON metadata for the struct
// [CatalogSyncNewResponseError]
type catalogSyncNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseErrorsCode int64

const (
	CatalogSyncNewResponseErrorsCode1001   CatalogSyncNewResponseErrorsCode = 1001
	CatalogSyncNewResponseErrorsCode1002   CatalogSyncNewResponseErrorsCode = 1002
	CatalogSyncNewResponseErrorsCode1003   CatalogSyncNewResponseErrorsCode = 1003
	CatalogSyncNewResponseErrorsCode1004   CatalogSyncNewResponseErrorsCode = 1004
	CatalogSyncNewResponseErrorsCode1005   CatalogSyncNewResponseErrorsCode = 1005
	CatalogSyncNewResponseErrorsCode1006   CatalogSyncNewResponseErrorsCode = 1006
	CatalogSyncNewResponseErrorsCode1007   CatalogSyncNewResponseErrorsCode = 1007
	CatalogSyncNewResponseErrorsCode1008   CatalogSyncNewResponseErrorsCode = 1008
	CatalogSyncNewResponseErrorsCode1009   CatalogSyncNewResponseErrorsCode = 1009
	CatalogSyncNewResponseErrorsCode1010   CatalogSyncNewResponseErrorsCode = 1010
	CatalogSyncNewResponseErrorsCode1011   CatalogSyncNewResponseErrorsCode = 1011
	CatalogSyncNewResponseErrorsCode1012   CatalogSyncNewResponseErrorsCode = 1012
	CatalogSyncNewResponseErrorsCode1013   CatalogSyncNewResponseErrorsCode = 1013
	CatalogSyncNewResponseErrorsCode1014   CatalogSyncNewResponseErrorsCode = 1014
	CatalogSyncNewResponseErrorsCode1015   CatalogSyncNewResponseErrorsCode = 1015
	CatalogSyncNewResponseErrorsCode1016   CatalogSyncNewResponseErrorsCode = 1016
	CatalogSyncNewResponseErrorsCode1017   CatalogSyncNewResponseErrorsCode = 1017
	CatalogSyncNewResponseErrorsCode2001   CatalogSyncNewResponseErrorsCode = 2001
	CatalogSyncNewResponseErrorsCode2002   CatalogSyncNewResponseErrorsCode = 2002
	CatalogSyncNewResponseErrorsCode2003   CatalogSyncNewResponseErrorsCode = 2003
	CatalogSyncNewResponseErrorsCode2004   CatalogSyncNewResponseErrorsCode = 2004
	CatalogSyncNewResponseErrorsCode2005   CatalogSyncNewResponseErrorsCode = 2005
	CatalogSyncNewResponseErrorsCode2006   CatalogSyncNewResponseErrorsCode = 2006
	CatalogSyncNewResponseErrorsCode2007   CatalogSyncNewResponseErrorsCode = 2007
	CatalogSyncNewResponseErrorsCode2008   CatalogSyncNewResponseErrorsCode = 2008
	CatalogSyncNewResponseErrorsCode2009   CatalogSyncNewResponseErrorsCode = 2009
	CatalogSyncNewResponseErrorsCode2010   CatalogSyncNewResponseErrorsCode = 2010
	CatalogSyncNewResponseErrorsCode2011   CatalogSyncNewResponseErrorsCode = 2011
	CatalogSyncNewResponseErrorsCode2012   CatalogSyncNewResponseErrorsCode = 2012
	CatalogSyncNewResponseErrorsCode2013   CatalogSyncNewResponseErrorsCode = 2013
	CatalogSyncNewResponseErrorsCode2014   CatalogSyncNewResponseErrorsCode = 2014
	CatalogSyncNewResponseErrorsCode2015   CatalogSyncNewResponseErrorsCode = 2015
	CatalogSyncNewResponseErrorsCode2016   CatalogSyncNewResponseErrorsCode = 2016
	CatalogSyncNewResponseErrorsCode2017   CatalogSyncNewResponseErrorsCode = 2017
	CatalogSyncNewResponseErrorsCode2018   CatalogSyncNewResponseErrorsCode = 2018
	CatalogSyncNewResponseErrorsCode2019   CatalogSyncNewResponseErrorsCode = 2019
	CatalogSyncNewResponseErrorsCode2020   CatalogSyncNewResponseErrorsCode = 2020
	CatalogSyncNewResponseErrorsCode2021   CatalogSyncNewResponseErrorsCode = 2021
	CatalogSyncNewResponseErrorsCode2022   CatalogSyncNewResponseErrorsCode = 2022
	CatalogSyncNewResponseErrorsCode3001   CatalogSyncNewResponseErrorsCode = 3001
	CatalogSyncNewResponseErrorsCode3002   CatalogSyncNewResponseErrorsCode = 3002
	CatalogSyncNewResponseErrorsCode3003   CatalogSyncNewResponseErrorsCode = 3003
	CatalogSyncNewResponseErrorsCode3004   CatalogSyncNewResponseErrorsCode = 3004
	CatalogSyncNewResponseErrorsCode3005   CatalogSyncNewResponseErrorsCode = 3005
	CatalogSyncNewResponseErrorsCode3006   CatalogSyncNewResponseErrorsCode = 3006
	CatalogSyncNewResponseErrorsCode3007   CatalogSyncNewResponseErrorsCode = 3007
	CatalogSyncNewResponseErrorsCode4001   CatalogSyncNewResponseErrorsCode = 4001
	CatalogSyncNewResponseErrorsCode4002   CatalogSyncNewResponseErrorsCode = 4002
	CatalogSyncNewResponseErrorsCode4003   CatalogSyncNewResponseErrorsCode = 4003
	CatalogSyncNewResponseErrorsCode4004   CatalogSyncNewResponseErrorsCode = 4004
	CatalogSyncNewResponseErrorsCode4005   CatalogSyncNewResponseErrorsCode = 4005
	CatalogSyncNewResponseErrorsCode4006   CatalogSyncNewResponseErrorsCode = 4006
	CatalogSyncNewResponseErrorsCode4007   CatalogSyncNewResponseErrorsCode = 4007
	CatalogSyncNewResponseErrorsCode4008   CatalogSyncNewResponseErrorsCode = 4008
	CatalogSyncNewResponseErrorsCode4009   CatalogSyncNewResponseErrorsCode = 4009
	CatalogSyncNewResponseErrorsCode4010   CatalogSyncNewResponseErrorsCode = 4010
	CatalogSyncNewResponseErrorsCode4011   CatalogSyncNewResponseErrorsCode = 4011
	CatalogSyncNewResponseErrorsCode4012   CatalogSyncNewResponseErrorsCode = 4012
	CatalogSyncNewResponseErrorsCode4013   CatalogSyncNewResponseErrorsCode = 4013
	CatalogSyncNewResponseErrorsCode4014   CatalogSyncNewResponseErrorsCode = 4014
	CatalogSyncNewResponseErrorsCode4015   CatalogSyncNewResponseErrorsCode = 4015
	CatalogSyncNewResponseErrorsCode4016   CatalogSyncNewResponseErrorsCode = 4016
	CatalogSyncNewResponseErrorsCode4017   CatalogSyncNewResponseErrorsCode = 4017
	CatalogSyncNewResponseErrorsCode4018   CatalogSyncNewResponseErrorsCode = 4018
	CatalogSyncNewResponseErrorsCode4019   CatalogSyncNewResponseErrorsCode = 4019
	CatalogSyncNewResponseErrorsCode4020   CatalogSyncNewResponseErrorsCode = 4020
	CatalogSyncNewResponseErrorsCode4021   CatalogSyncNewResponseErrorsCode = 4021
	CatalogSyncNewResponseErrorsCode4022   CatalogSyncNewResponseErrorsCode = 4022
	CatalogSyncNewResponseErrorsCode4023   CatalogSyncNewResponseErrorsCode = 4023
	CatalogSyncNewResponseErrorsCode5001   CatalogSyncNewResponseErrorsCode = 5001
	CatalogSyncNewResponseErrorsCode5002   CatalogSyncNewResponseErrorsCode = 5002
	CatalogSyncNewResponseErrorsCode5003   CatalogSyncNewResponseErrorsCode = 5003
	CatalogSyncNewResponseErrorsCode5004   CatalogSyncNewResponseErrorsCode = 5004
	CatalogSyncNewResponseErrorsCode102000 CatalogSyncNewResponseErrorsCode = 102000
	CatalogSyncNewResponseErrorsCode102001 CatalogSyncNewResponseErrorsCode = 102001
	CatalogSyncNewResponseErrorsCode102002 CatalogSyncNewResponseErrorsCode = 102002
	CatalogSyncNewResponseErrorsCode102003 CatalogSyncNewResponseErrorsCode = 102003
	CatalogSyncNewResponseErrorsCode102004 CatalogSyncNewResponseErrorsCode = 102004
	CatalogSyncNewResponseErrorsCode102005 CatalogSyncNewResponseErrorsCode = 102005
	CatalogSyncNewResponseErrorsCode102006 CatalogSyncNewResponseErrorsCode = 102006
	CatalogSyncNewResponseErrorsCode102007 CatalogSyncNewResponseErrorsCode = 102007
	CatalogSyncNewResponseErrorsCode102008 CatalogSyncNewResponseErrorsCode = 102008
	CatalogSyncNewResponseErrorsCode102009 CatalogSyncNewResponseErrorsCode = 102009
	CatalogSyncNewResponseErrorsCode102010 CatalogSyncNewResponseErrorsCode = 102010
	CatalogSyncNewResponseErrorsCode102011 CatalogSyncNewResponseErrorsCode = 102011
	CatalogSyncNewResponseErrorsCode102012 CatalogSyncNewResponseErrorsCode = 102012
	CatalogSyncNewResponseErrorsCode102013 CatalogSyncNewResponseErrorsCode = 102013
	CatalogSyncNewResponseErrorsCode102014 CatalogSyncNewResponseErrorsCode = 102014
	CatalogSyncNewResponseErrorsCode102015 CatalogSyncNewResponseErrorsCode = 102015
	CatalogSyncNewResponseErrorsCode102016 CatalogSyncNewResponseErrorsCode = 102016
	CatalogSyncNewResponseErrorsCode102017 CatalogSyncNewResponseErrorsCode = 102017
	CatalogSyncNewResponseErrorsCode102018 CatalogSyncNewResponseErrorsCode = 102018
	CatalogSyncNewResponseErrorsCode102019 CatalogSyncNewResponseErrorsCode = 102019
	CatalogSyncNewResponseErrorsCode102020 CatalogSyncNewResponseErrorsCode = 102020
	CatalogSyncNewResponseErrorsCode102021 CatalogSyncNewResponseErrorsCode = 102021
	CatalogSyncNewResponseErrorsCode102022 CatalogSyncNewResponseErrorsCode = 102022
	CatalogSyncNewResponseErrorsCode102023 CatalogSyncNewResponseErrorsCode = 102023
	CatalogSyncNewResponseErrorsCode102024 CatalogSyncNewResponseErrorsCode = 102024
	CatalogSyncNewResponseErrorsCode102025 CatalogSyncNewResponseErrorsCode = 102025
	CatalogSyncNewResponseErrorsCode102026 CatalogSyncNewResponseErrorsCode = 102026
	CatalogSyncNewResponseErrorsCode102027 CatalogSyncNewResponseErrorsCode = 102027
	CatalogSyncNewResponseErrorsCode102028 CatalogSyncNewResponseErrorsCode = 102028
	CatalogSyncNewResponseErrorsCode102029 CatalogSyncNewResponseErrorsCode = 102029
	CatalogSyncNewResponseErrorsCode102030 CatalogSyncNewResponseErrorsCode = 102030
	CatalogSyncNewResponseErrorsCode102031 CatalogSyncNewResponseErrorsCode = 102031
	CatalogSyncNewResponseErrorsCode102032 CatalogSyncNewResponseErrorsCode = 102032
	CatalogSyncNewResponseErrorsCode102033 CatalogSyncNewResponseErrorsCode = 102033
	CatalogSyncNewResponseErrorsCode102034 CatalogSyncNewResponseErrorsCode = 102034
	CatalogSyncNewResponseErrorsCode102035 CatalogSyncNewResponseErrorsCode = 102035
	CatalogSyncNewResponseErrorsCode102036 CatalogSyncNewResponseErrorsCode = 102036
	CatalogSyncNewResponseErrorsCode102037 CatalogSyncNewResponseErrorsCode = 102037
	CatalogSyncNewResponseErrorsCode102038 CatalogSyncNewResponseErrorsCode = 102038
	CatalogSyncNewResponseErrorsCode102039 CatalogSyncNewResponseErrorsCode = 102039
	CatalogSyncNewResponseErrorsCode102040 CatalogSyncNewResponseErrorsCode = 102040
	CatalogSyncNewResponseErrorsCode102041 CatalogSyncNewResponseErrorsCode = 102041
	CatalogSyncNewResponseErrorsCode102042 CatalogSyncNewResponseErrorsCode = 102042
	CatalogSyncNewResponseErrorsCode102043 CatalogSyncNewResponseErrorsCode = 102043
	CatalogSyncNewResponseErrorsCode102044 CatalogSyncNewResponseErrorsCode = 102044
	CatalogSyncNewResponseErrorsCode102045 CatalogSyncNewResponseErrorsCode = 102045
	CatalogSyncNewResponseErrorsCode102046 CatalogSyncNewResponseErrorsCode = 102046
	CatalogSyncNewResponseErrorsCode102047 CatalogSyncNewResponseErrorsCode = 102047
	CatalogSyncNewResponseErrorsCode102048 CatalogSyncNewResponseErrorsCode = 102048
	CatalogSyncNewResponseErrorsCode102049 CatalogSyncNewResponseErrorsCode = 102049
	CatalogSyncNewResponseErrorsCode102050 CatalogSyncNewResponseErrorsCode = 102050
	CatalogSyncNewResponseErrorsCode102051 CatalogSyncNewResponseErrorsCode = 102051
	CatalogSyncNewResponseErrorsCode102052 CatalogSyncNewResponseErrorsCode = 102052
	CatalogSyncNewResponseErrorsCode102053 CatalogSyncNewResponseErrorsCode = 102053
	CatalogSyncNewResponseErrorsCode102054 CatalogSyncNewResponseErrorsCode = 102054
	CatalogSyncNewResponseErrorsCode102055 CatalogSyncNewResponseErrorsCode = 102055
	CatalogSyncNewResponseErrorsCode102056 CatalogSyncNewResponseErrorsCode = 102056
	CatalogSyncNewResponseErrorsCode102057 CatalogSyncNewResponseErrorsCode = 102057
	CatalogSyncNewResponseErrorsCode102058 CatalogSyncNewResponseErrorsCode = 102058
	CatalogSyncNewResponseErrorsCode102059 CatalogSyncNewResponseErrorsCode = 102059
	CatalogSyncNewResponseErrorsCode102060 CatalogSyncNewResponseErrorsCode = 102060
	CatalogSyncNewResponseErrorsCode102061 CatalogSyncNewResponseErrorsCode = 102061
	CatalogSyncNewResponseErrorsCode102062 CatalogSyncNewResponseErrorsCode = 102062
	CatalogSyncNewResponseErrorsCode102063 CatalogSyncNewResponseErrorsCode = 102063
	CatalogSyncNewResponseErrorsCode102064 CatalogSyncNewResponseErrorsCode = 102064
	CatalogSyncNewResponseErrorsCode102065 CatalogSyncNewResponseErrorsCode = 102065
	CatalogSyncNewResponseErrorsCode102066 CatalogSyncNewResponseErrorsCode = 102066
	CatalogSyncNewResponseErrorsCode103001 CatalogSyncNewResponseErrorsCode = 103001
	CatalogSyncNewResponseErrorsCode103002 CatalogSyncNewResponseErrorsCode = 103002
	CatalogSyncNewResponseErrorsCode103003 CatalogSyncNewResponseErrorsCode = 103003
	CatalogSyncNewResponseErrorsCode103004 CatalogSyncNewResponseErrorsCode = 103004
	CatalogSyncNewResponseErrorsCode103005 CatalogSyncNewResponseErrorsCode = 103005
	CatalogSyncNewResponseErrorsCode103006 CatalogSyncNewResponseErrorsCode = 103006
	CatalogSyncNewResponseErrorsCode103007 CatalogSyncNewResponseErrorsCode = 103007
	CatalogSyncNewResponseErrorsCode103008 CatalogSyncNewResponseErrorsCode = 103008
)

func (r CatalogSyncNewResponseErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncNewResponseErrorsCode1001, CatalogSyncNewResponseErrorsCode1002, CatalogSyncNewResponseErrorsCode1003, CatalogSyncNewResponseErrorsCode1004, CatalogSyncNewResponseErrorsCode1005, CatalogSyncNewResponseErrorsCode1006, CatalogSyncNewResponseErrorsCode1007, CatalogSyncNewResponseErrorsCode1008, CatalogSyncNewResponseErrorsCode1009, CatalogSyncNewResponseErrorsCode1010, CatalogSyncNewResponseErrorsCode1011, CatalogSyncNewResponseErrorsCode1012, CatalogSyncNewResponseErrorsCode1013, CatalogSyncNewResponseErrorsCode1014, CatalogSyncNewResponseErrorsCode1015, CatalogSyncNewResponseErrorsCode1016, CatalogSyncNewResponseErrorsCode1017, CatalogSyncNewResponseErrorsCode2001, CatalogSyncNewResponseErrorsCode2002, CatalogSyncNewResponseErrorsCode2003, CatalogSyncNewResponseErrorsCode2004, CatalogSyncNewResponseErrorsCode2005, CatalogSyncNewResponseErrorsCode2006, CatalogSyncNewResponseErrorsCode2007, CatalogSyncNewResponseErrorsCode2008, CatalogSyncNewResponseErrorsCode2009, CatalogSyncNewResponseErrorsCode2010, CatalogSyncNewResponseErrorsCode2011, CatalogSyncNewResponseErrorsCode2012, CatalogSyncNewResponseErrorsCode2013, CatalogSyncNewResponseErrorsCode2014, CatalogSyncNewResponseErrorsCode2015, CatalogSyncNewResponseErrorsCode2016, CatalogSyncNewResponseErrorsCode2017, CatalogSyncNewResponseErrorsCode2018, CatalogSyncNewResponseErrorsCode2019, CatalogSyncNewResponseErrorsCode2020, CatalogSyncNewResponseErrorsCode2021, CatalogSyncNewResponseErrorsCode2022, CatalogSyncNewResponseErrorsCode3001, CatalogSyncNewResponseErrorsCode3002, CatalogSyncNewResponseErrorsCode3003, CatalogSyncNewResponseErrorsCode3004, CatalogSyncNewResponseErrorsCode3005, CatalogSyncNewResponseErrorsCode3006, CatalogSyncNewResponseErrorsCode3007, CatalogSyncNewResponseErrorsCode4001, CatalogSyncNewResponseErrorsCode4002, CatalogSyncNewResponseErrorsCode4003, CatalogSyncNewResponseErrorsCode4004, CatalogSyncNewResponseErrorsCode4005, CatalogSyncNewResponseErrorsCode4006, CatalogSyncNewResponseErrorsCode4007, CatalogSyncNewResponseErrorsCode4008, CatalogSyncNewResponseErrorsCode4009, CatalogSyncNewResponseErrorsCode4010, CatalogSyncNewResponseErrorsCode4011, CatalogSyncNewResponseErrorsCode4012, CatalogSyncNewResponseErrorsCode4013, CatalogSyncNewResponseErrorsCode4014, CatalogSyncNewResponseErrorsCode4015, CatalogSyncNewResponseErrorsCode4016, CatalogSyncNewResponseErrorsCode4017, CatalogSyncNewResponseErrorsCode4018, CatalogSyncNewResponseErrorsCode4019, CatalogSyncNewResponseErrorsCode4020, CatalogSyncNewResponseErrorsCode4021, CatalogSyncNewResponseErrorsCode4022, CatalogSyncNewResponseErrorsCode4023, CatalogSyncNewResponseErrorsCode5001, CatalogSyncNewResponseErrorsCode5002, CatalogSyncNewResponseErrorsCode5003, CatalogSyncNewResponseErrorsCode5004, CatalogSyncNewResponseErrorsCode102000, CatalogSyncNewResponseErrorsCode102001, CatalogSyncNewResponseErrorsCode102002, CatalogSyncNewResponseErrorsCode102003, CatalogSyncNewResponseErrorsCode102004, CatalogSyncNewResponseErrorsCode102005, CatalogSyncNewResponseErrorsCode102006, CatalogSyncNewResponseErrorsCode102007, CatalogSyncNewResponseErrorsCode102008, CatalogSyncNewResponseErrorsCode102009, CatalogSyncNewResponseErrorsCode102010, CatalogSyncNewResponseErrorsCode102011, CatalogSyncNewResponseErrorsCode102012, CatalogSyncNewResponseErrorsCode102013, CatalogSyncNewResponseErrorsCode102014, CatalogSyncNewResponseErrorsCode102015, CatalogSyncNewResponseErrorsCode102016, CatalogSyncNewResponseErrorsCode102017, CatalogSyncNewResponseErrorsCode102018, CatalogSyncNewResponseErrorsCode102019, CatalogSyncNewResponseErrorsCode102020, CatalogSyncNewResponseErrorsCode102021, CatalogSyncNewResponseErrorsCode102022, CatalogSyncNewResponseErrorsCode102023, CatalogSyncNewResponseErrorsCode102024, CatalogSyncNewResponseErrorsCode102025, CatalogSyncNewResponseErrorsCode102026, CatalogSyncNewResponseErrorsCode102027, CatalogSyncNewResponseErrorsCode102028, CatalogSyncNewResponseErrorsCode102029, CatalogSyncNewResponseErrorsCode102030, CatalogSyncNewResponseErrorsCode102031, CatalogSyncNewResponseErrorsCode102032, CatalogSyncNewResponseErrorsCode102033, CatalogSyncNewResponseErrorsCode102034, CatalogSyncNewResponseErrorsCode102035, CatalogSyncNewResponseErrorsCode102036, CatalogSyncNewResponseErrorsCode102037, CatalogSyncNewResponseErrorsCode102038, CatalogSyncNewResponseErrorsCode102039, CatalogSyncNewResponseErrorsCode102040, CatalogSyncNewResponseErrorsCode102041, CatalogSyncNewResponseErrorsCode102042, CatalogSyncNewResponseErrorsCode102043, CatalogSyncNewResponseErrorsCode102044, CatalogSyncNewResponseErrorsCode102045, CatalogSyncNewResponseErrorsCode102046, CatalogSyncNewResponseErrorsCode102047, CatalogSyncNewResponseErrorsCode102048, CatalogSyncNewResponseErrorsCode102049, CatalogSyncNewResponseErrorsCode102050, CatalogSyncNewResponseErrorsCode102051, CatalogSyncNewResponseErrorsCode102052, CatalogSyncNewResponseErrorsCode102053, CatalogSyncNewResponseErrorsCode102054, CatalogSyncNewResponseErrorsCode102055, CatalogSyncNewResponseErrorsCode102056, CatalogSyncNewResponseErrorsCode102057, CatalogSyncNewResponseErrorsCode102058, CatalogSyncNewResponseErrorsCode102059, CatalogSyncNewResponseErrorsCode102060, CatalogSyncNewResponseErrorsCode102061, CatalogSyncNewResponseErrorsCode102062, CatalogSyncNewResponseErrorsCode102063, CatalogSyncNewResponseErrorsCode102064, CatalogSyncNewResponseErrorsCode102065, CatalogSyncNewResponseErrorsCode102066, CatalogSyncNewResponseErrorsCode103001, CatalogSyncNewResponseErrorsCode103002, CatalogSyncNewResponseErrorsCode103003, CatalogSyncNewResponseErrorsCode103004, CatalogSyncNewResponseErrorsCode103005, CatalogSyncNewResponseErrorsCode103006, CatalogSyncNewResponseErrorsCode103007, CatalogSyncNewResponseErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncNewResponseErrorsMeta struct {
	L10nKey       string                               `json:"l10n_key"`
	LoggableError string                               `json:"loggable_error"`
	TemplateData  interface{}                          `json:"template_data"`
	TraceID       string                               `json:"trace_id"`
	JSON          catalogSyncNewResponseErrorsMetaJSON `json:"-"`
}

// catalogSyncNewResponseErrorsMetaJSON contains the JSON metadata for the struct
// [CatalogSyncNewResponseErrorsMeta]
type catalogSyncNewResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncNewResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseErrorsSource struct {
	Parameter           string                                 `json:"parameter"`
	ParameterValueIndex int64                                  `json:"parameter_value_index"`
	Pointer             string                                 `json:"pointer"`
	JSON                catalogSyncNewResponseErrorsSourceJSON `json:"-"`
}

// catalogSyncNewResponseErrorsSourceJSON contains the JSON metadata for the struct
// [CatalogSyncNewResponseErrorsSource]
type catalogSyncNewResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponse struct {
	ID                       string                                    `json:"id,required" format:"uuid"`
	Description              string                                    `json:"description,required"`
	DestinationID            string                                    `json:"destination_id,required" format:"uuid"`
	DestinationType          CatalogSyncUpdateResponseDestinationType  `json:"destination_type,required"`
	LastUserUpdateAt         string                                    `json:"last_user_update_at,required"`
	Name                     string                                    `json:"name,required"`
	Policy                   string                                    `json:"policy,required"`
	UpdateMode               CatalogSyncUpdateResponseUpdateMode       `json:"update_mode,required"`
	Errors                   map[string]CatalogSyncUpdateResponseError `json:"errors"`
	IncludesDiscoveriesUntil string                                    `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                                    `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                                    `json:"last_successful_update_at"`
	JSON                     catalogSyncUpdateResponseJSON             `json:"-"`
}

// catalogSyncUpdateResponseJSON contains the JSON metadata for the struct
// [CatalogSyncUpdateResponse]
type catalogSyncUpdateResponseJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseDestinationType string

const (
	CatalogSyncUpdateResponseDestinationTypeNone          CatalogSyncUpdateResponseDestinationType = "NONE"
	CatalogSyncUpdateResponseDestinationTypeZeroTrustList CatalogSyncUpdateResponseDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncUpdateResponseDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateResponseDestinationTypeNone, CatalogSyncUpdateResponseDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseUpdateMode string

const (
	CatalogSyncUpdateResponseUpdateModeAuto   CatalogSyncUpdateResponseUpdateMode = "AUTO"
	CatalogSyncUpdateResponseUpdateModeManual CatalogSyncUpdateResponseUpdateMode = "MANUAL"
)

func (r CatalogSyncUpdateResponseUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateResponseUpdateModeAuto, CatalogSyncUpdateResponseUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseError struct {
	Code             CatalogSyncUpdateResponseErrorsCode   `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Meta             CatalogSyncUpdateResponseErrorsMeta   `json:"meta"`
	Source           CatalogSyncUpdateResponseErrorsSource `json:"source"`
	JSON             catalogSyncUpdateResponseErrorJSON    `json:"-"`
}

// catalogSyncUpdateResponseErrorJSON contains the JSON metadata for the struct
// [CatalogSyncUpdateResponseError]
type catalogSyncUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseErrorsCode int64

const (
	CatalogSyncUpdateResponseErrorsCode1001   CatalogSyncUpdateResponseErrorsCode = 1001
	CatalogSyncUpdateResponseErrorsCode1002   CatalogSyncUpdateResponseErrorsCode = 1002
	CatalogSyncUpdateResponseErrorsCode1003   CatalogSyncUpdateResponseErrorsCode = 1003
	CatalogSyncUpdateResponseErrorsCode1004   CatalogSyncUpdateResponseErrorsCode = 1004
	CatalogSyncUpdateResponseErrorsCode1005   CatalogSyncUpdateResponseErrorsCode = 1005
	CatalogSyncUpdateResponseErrorsCode1006   CatalogSyncUpdateResponseErrorsCode = 1006
	CatalogSyncUpdateResponseErrorsCode1007   CatalogSyncUpdateResponseErrorsCode = 1007
	CatalogSyncUpdateResponseErrorsCode1008   CatalogSyncUpdateResponseErrorsCode = 1008
	CatalogSyncUpdateResponseErrorsCode1009   CatalogSyncUpdateResponseErrorsCode = 1009
	CatalogSyncUpdateResponseErrorsCode1010   CatalogSyncUpdateResponseErrorsCode = 1010
	CatalogSyncUpdateResponseErrorsCode1011   CatalogSyncUpdateResponseErrorsCode = 1011
	CatalogSyncUpdateResponseErrorsCode1012   CatalogSyncUpdateResponseErrorsCode = 1012
	CatalogSyncUpdateResponseErrorsCode1013   CatalogSyncUpdateResponseErrorsCode = 1013
	CatalogSyncUpdateResponseErrorsCode1014   CatalogSyncUpdateResponseErrorsCode = 1014
	CatalogSyncUpdateResponseErrorsCode1015   CatalogSyncUpdateResponseErrorsCode = 1015
	CatalogSyncUpdateResponseErrorsCode1016   CatalogSyncUpdateResponseErrorsCode = 1016
	CatalogSyncUpdateResponseErrorsCode1017   CatalogSyncUpdateResponseErrorsCode = 1017
	CatalogSyncUpdateResponseErrorsCode2001   CatalogSyncUpdateResponseErrorsCode = 2001
	CatalogSyncUpdateResponseErrorsCode2002   CatalogSyncUpdateResponseErrorsCode = 2002
	CatalogSyncUpdateResponseErrorsCode2003   CatalogSyncUpdateResponseErrorsCode = 2003
	CatalogSyncUpdateResponseErrorsCode2004   CatalogSyncUpdateResponseErrorsCode = 2004
	CatalogSyncUpdateResponseErrorsCode2005   CatalogSyncUpdateResponseErrorsCode = 2005
	CatalogSyncUpdateResponseErrorsCode2006   CatalogSyncUpdateResponseErrorsCode = 2006
	CatalogSyncUpdateResponseErrorsCode2007   CatalogSyncUpdateResponseErrorsCode = 2007
	CatalogSyncUpdateResponseErrorsCode2008   CatalogSyncUpdateResponseErrorsCode = 2008
	CatalogSyncUpdateResponseErrorsCode2009   CatalogSyncUpdateResponseErrorsCode = 2009
	CatalogSyncUpdateResponseErrorsCode2010   CatalogSyncUpdateResponseErrorsCode = 2010
	CatalogSyncUpdateResponseErrorsCode2011   CatalogSyncUpdateResponseErrorsCode = 2011
	CatalogSyncUpdateResponseErrorsCode2012   CatalogSyncUpdateResponseErrorsCode = 2012
	CatalogSyncUpdateResponseErrorsCode2013   CatalogSyncUpdateResponseErrorsCode = 2013
	CatalogSyncUpdateResponseErrorsCode2014   CatalogSyncUpdateResponseErrorsCode = 2014
	CatalogSyncUpdateResponseErrorsCode2015   CatalogSyncUpdateResponseErrorsCode = 2015
	CatalogSyncUpdateResponseErrorsCode2016   CatalogSyncUpdateResponseErrorsCode = 2016
	CatalogSyncUpdateResponseErrorsCode2017   CatalogSyncUpdateResponseErrorsCode = 2017
	CatalogSyncUpdateResponseErrorsCode2018   CatalogSyncUpdateResponseErrorsCode = 2018
	CatalogSyncUpdateResponseErrorsCode2019   CatalogSyncUpdateResponseErrorsCode = 2019
	CatalogSyncUpdateResponseErrorsCode2020   CatalogSyncUpdateResponseErrorsCode = 2020
	CatalogSyncUpdateResponseErrorsCode2021   CatalogSyncUpdateResponseErrorsCode = 2021
	CatalogSyncUpdateResponseErrorsCode2022   CatalogSyncUpdateResponseErrorsCode = 2022
	CatalogSyncUpdateResponseErrorsCode3001   CatalogSyncUpdateResponseErrorsCode = 3001
	CatalogSyncUpdateResponseErrorsCode3002   CatalogSyncUpdateResponseErrorsCode = 3002
	CatalogSyncUpdateResponseErrorsCode3003   CatalogSyncUpdateResponseErrorsCode = 3003
	CatalogSyncUpdateResponseErrorsCode3004   CatalogSyncUpdateResponseErrorsCode = 3004
	CatalogSyncUpdateResponseErrorsCode3005   CatalogSyncUpdateResponseErrorsCode = 3005
	CatalogSyncUpdateResponseErrorsCode3006   CatalogSyncUpdateResponseErrorsCode = 3006
	CatalogSyncUpdateResponseErrorsCode3007   CatalogSyncUpdateResponseErrorsCode = 3007
	CatalogSyncUpdateResponseErrorsCode4001   CatalogSyncUpdateResponseErrorsCode = 4001
	CatalogSyncUpdateResponseErrorsCode4002   CatalogSyncUpdateResponseErrorsCode = 4002
	CatalogSyncUpdateResponseErrorsCode4003   CatalogSyncUpdateResponseErrorsCode = 4003
	CatalogSyncUpdateResponseErrorsCode4004   CatalogSyncUpdateResponseErrorsCode = 4004
	CatalogSyncUpdateResponseErrorsCode4005   CatalogSyncUpdateResponseErrorsCode = 4005
	CatalogSyncUpdateResponseErrorsCode4006   CatalogSyncUpdateResponseErrorsCode = 4006
	CatalogSyncUpdateResponseErrorsCode4007   CatalogSyncUpdateResponseErrorsCode = 4007
	CatalogSyncUpdateResponseErrorsCode4008   CatalogSyncUpdateResponseErrorsCode = 4008
	CatalogSyncUpdateResponseErrorsCode4009   CatalogSyncUpdateResponseErrorsCode = 4009
	CatalogSyncUpdateResponseErrorsCode4010   CatalogSyncUpdateResponseErrorsCode = 4010
	CatalogSyncUpdateResponseErrorsCode4011   CatalogSyncUpdateResponseErrorsCode = 4011
	CatalogSyncUpdateResponseErrorsCode4012   CatalogSyncUpdateResponseErrorsCode = 4012
	CatalogSyncUpdateResponseErrorsCode4013   CatalogSyncUpdateResponseErrorsCode = 4013
	CatalogSyncUpdateResponseErrorsCode4014   CatalogSyncUpdateResponseErrorsCode = 4014
	CatalogSyncUpdateResponseErrorsCode4015   CatalogSyncUpdateResponseErrorsCode = 4015
	CatalogSyncUpdateResponseErrorsCode4016   CatalogSyncUpdateResponseErrorsCode = 4016
	CatalogSyncUpdateResponseErrorsCode4017   CatalogSyncUpdateResponseErrorsCode = 4017
	CatalogSyncUpdateResponseErrorsCode4018   CatalogSyncUpdateResponseErrorsCode = 4018
	CatalogSyncUpdateResponseErrorsCode4019   CatalogSyncUpdateResponseErrorsCode = 4019
	CatalogSyncUpdateResponseErrorsCode4020   CatalogSyncUpdateResponseErrorsCode = 4020
	CatalogSyncUpdateResponseErrorsCode4021   CatalogSyncUpdateResponseErrorsCode = 4021
	CatalogSyncUpdateResponseErrorsCode4022   CatalogSyncUpdateResponseErrorsCode = 4022
	CatalogSyncUpdateResponseErrorsCode4023   CatalogSyncUpdateResponseErrorsCode = 4023
	CatalogSyncUpdateResponseErrorsCode5001   CatalogSyncUpdateResponseErrorsCode = 5001
	CatalogSyncUpdateResponseErrorsCode5002   CatalogSyncUpdateResponseErrorsCode = 5002
	CatalogSyncUpdateResponseErrorsCode5003   CatalogSyncUpdateResponseErrorsCode = 5003
	CatalogSyncUpdateResponseErrorsCode5004   CatalogSyncUpdateResponseErrorsCode = 5004
	CatalogSyncUpdateResponseErrorsCode102000 CatalogSyncUpdateResponseErrorsCode = 102000
	CatalogSyncUpdateResponseErrorsCode102001 CatalogSyncUpdateResponseErrorsCode = 102001
	CatalogSyncUpdateResponseErrorsCode102002 CatalogSyncUpdateResponseErrorsCode = 102002
	CatalogSyncUpdateResponseErrorsCode102003 CatalogSyncUpdateResponseErrorsCode = 102003
	CatalogSyncUpdateResponseErrorsCode102004 CatalogSyncUpdateResponseErrorsCode = 102004
	CatalogSyncUpdateResponseErrorsCode102005 CatalogSyncUpdateResponseErrorsCode = 102005
	CatalogSyncUpdateResponseErrorsCode102006 CatalogSyncUpdateResponseErrorsCode = 102006
	CatalogSyncUpdateResponseErrorsCode102007 CatalogSyncUpdateResponseErrorsCode = 102007
	CatalogSyncUpdateResponseErrorsCode102008 CatalogSyncUpdateResponseErrorsCode = 102008
	CatalogSyncUpdateResponseErrorsCode102009 CatalogSyncUpdateResponseErrorsCode = 102009
	CatalogSyncUpdateResponseErrorsCode102010 CatalogSyncUpdateResponseErrorsCode = 102010
	CatalogSyncUpdateResponseErrorsCode102011 CatalogSyncUpdateResponseErrorsCode = 102011
	CatalogSyncUpdateResponseErrorsCode102012 CatalogSyncUpdateResponseErrorsCode = 102012
	CatalogSyncUpdateResponseErrorsCode102013 CatalogSyncUpdateResponseErrorsCode = 102013
	CatalogSyncUpdateResponseErrorsCode102014 CatalogSyncUpdateResponseErrorsCode = 102014
	CatalogSyncUpdateResponseErrorsCode102015 CatalogSyncUpdateResponseErrorsCode = 102015
	CatalogSyncUpdateResponseErrorsCode102016 CatalogSyncUpdateResponseErrorsCode = 102016
	CatalogSyncUpdateResponseErrorsCode102017 CatalogSyncUpdateResponseErrorsCode = 102017
	CatalogSyncUpdateResponseErrorsCode102018 CatalogSyncUpdateResponseErrorsCode = 102018
	CatalogSyncUpdateResponseErrorsCode102019 CatalogSyncUpdateResponseErrorsCode = 102019
	CatalogSyncUpdateResponseErrorsCode102020 CatalogSyncUpdateResponseErrorsCode = 102020
	CatalogSyncUpdateResponseErrorsCode102021 CatalogSyncUpdateResponseErrorsCode = 102021
	CatalogSyncUpdateResponseErrorsCode102022 CatalogSyncUpdateResponseErrorsCode = 102022
	CatalogSyncUpdateResponseErrorsCode102023 CatalogSyncUpdateResponseErrorsCode = 102023
	CatalogSyncUpdateResponseErrorsCode102024 CatalogSyncUpdateResponseErrorsCode = 102024
	CatalogSyncUpdateResponseErrorsCode102025 CatalogSyncUpdateResponseErrorsCode = 102025
	CatalogSyncUpdateResponseErrorsCode102026 CatalogSyncUpdateResponseErrorsCode = 102026
	CatalogSyncUpdateResponseErrorsCode102027 CatalogSyncUpdateResponseErrorsCode = 102027
	CatalogSyncUpdateResponseErrorsCode102028 CatalogSyncUpdateResponseErrorsCode = 102028
	CatalogSyncUpdateResponseErrorsCode102029 CatalogSyncUpdateResponseErrorsCode = 102029
	CatalogSyncUpdateResponseErrorsCode102030 CatalogSyncUpdateResponseErrorsCode = 102030
	CatalogSyncUpdateResponseErrorsCode102031 CatalogSyncUpdateResponseErrorsCode = 102031
	CatalogSyncUpdateResponseErrorsCode102032 CatalogSyncUpdateResponseErrorsCode = 102032
	CatalogSyncUpdateResponseErrorsCode102033 CatalogSyncUpdateResponseErrorsCode = 102033
	CatalogSyncUpdateResponseErrorsCode102034 CatalogSyncUpdateResponseErrorsCode = 102034
	CatalogSyncUpdateResponseErrorsCode102035 CatalogSyncUpdateResponseErrorsCode = 102035
	CatalogSyncUpdateResponseErrorsCode102036 CatalogSyncUpdateResponseErrorsCode = 102036
	CatalogSyncUpdateResponseErrorsCode102037 CatalogSyncUpdateResponseErrorsCode = 102037
	CatalogSyncUpdateResponseErrorsCode102038 CatalogSyncUpdateResponseErrorsCode = 102038
	CatalogSyncUpdateResponseErrorsCode102039 CatalogSyncUpdateResponseErrorsCode = 102039
	CatalogSyncUpdateResponseErrorsCode102040 CatalogSyncUpdateResponseErrorsCode = 102040
	CatalogSyncUpdateResponseErrorsCode102041 CatalogSyncUpdateResponseErrorsCode = 102041
	CatalogSyncUpdateResponseErrorsCode102042 CatalogSyncUpdateResponseErrorsCode = 102042
	CatalogSyncUpdateResponseErrorsCode102043 CatalogSyncUpdateResponseErrorsCode = 102043
	CatalogSyncUpdateResponseErrorsCode102044 CatalogSyncUpdateResponseErrorsCode = 102044
	CatalogSyncUpdateResponseErrorsCode102045 CatalogSyncUpdateResponseErrorsCode = 102045
	CatalogSyncUpdateResponseErrorsCode102046 CatalogSyncUpdateResponseErrorsCode = 102046
	CatalogSyncUpdateResponseErrorsCode102047 CatalogSyncUpdateResponseErrorsCode = 102047
	CatalogSyncUpdateResponseErrorsCode102048 CatalogSyncUpdateResponseErrorsCode = 102048
	CatalogSyncUpdateResponseErrorsCode102049 CatalogSyncUpdateResponseErrorsCode = 102049
	CatalogSyncUpdateResponseErrorsCode102050 CatalogSyncUpdateResponseErrorsCode = 102050
	CatalogSyncUpdateResponseErrorsCode102051 CatalogSyncUpdateResponseErrorsCode = 102051
	CatalogSyncUpdateResponseErrorsCode102052 CatalogSyncUpdateResponseErrorsCode = 102052
	CatalogSyncUpdateResponseErrorsCode102053 CatalogSyncUpdateResponseErrorsCode = 102053
	CatalogSyncUpdateResponseErrorsCode102054 CatalogSyncUpdateResponseErrorsCode = 102054
	CatalogSyncUpdateResponseErrorsCode102055 CatalogSyncUpdateResponseErrorsCode = 102055
	CatalogSyncUpdateResponseErrorsCode102056 CatalogSyncUpdateResponseErrorsCode = 102056
	CatalogSyncUpdateResponseErrorsCode102057 CatalogSyncUpdateResponseErrorsCode = 102057
	CatalogSyncUpdateResponseErrorsCode102058 CatalogSyncUpdateResponseErrorsCode = 102058
	CatalogSyncUpdateResponseErrorsCode102059 CatalogSyncUpdateResponseErrorsCode = 102059
	CatalogSyncUpdateResponseErrorsCode102060 CatalogSyncUpdateResponseErrorsCode = 102060
	CatalogSyncUpdateResponseErrorsCode102061 CatalogSyncUpdateResponseErrorsCode = 102061
	CatalogSyncUpdateResponseErrorsCode102062 CatalogSyncUpdateResponseErrorsCode = 102062
	CatalogSyncUpdateResponseErrorsCode102063 CatalogSyncUpdateResponseErrorsCode = 102063
	CatalogSyncUpdateResponseErrorsCode102064 CatalogSyncUpdateResponseErrorsCode = 102064
	CatalogSyncUpdateResponseErrorsCode102065 CatalogSyncUpdateResponseErrorsCode = 102065
	CatalogSyncUpdateResponseErrorsCode102066 CatalogSyncUpdateResponseErrorsCode = 102066
	CatalogSyncUpdateResponseErrorsCode103001 CatalogSyncUpdateResponseErrorsCode = 103001
	CatalogSyncUpdateResponseErrorsCode103002 CatalogSyncUpdateResponseErrorsCode = 103002
	CatalogSyncUpdateResponseErrorsCode103003 CatalogSyncUpdateResponseErrorsCode = 103003
	CatalogSyncUpdateResponseErrorsCode103004 CatalogSyncUpdateResponseErrorsCode = 103004
	CatalogSyncUpdateResponseErrorsCode103005 CatalogSyncUpdateResponseErrorsCode = 103005
	CatalogSyncUpdateResponseErrorsCode103006 CatalogSyncUpdateResponseErrorsCode = 103006
	CatalogSyncUpdateResponseErrorsCode103007 CatalogSyncUpdateResponseErrorsCode = 103007
	CatalogSyncUpdateResponseErrorsCode103008 CatalogSyncUpdateResponseErrorsCode = 103008
)

func (r CatalogSyncUpdateResponseErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateResponseErrorsCode1001, CatalogSyncUpdateResponseErrorsCode1002, CatalogSyncUpdateResponseErrorsCode1003, CatalogSyncUpdateResponseErrorsCode1004, CatalogSyncUpdateResponseErrorsCode1005, CatalogSyncUpdateResponseErrorsCode1006, CatalogSyncUpdateResponseErrorsCode1007, CatalogSyncUpdateResponseErrorsCode1008, CatalogSyncUpdateResponseErrorsCode1009, CatalogSyncUpdateResponseErrorsCode1010, CatalogSyncUpdateResponseErrorsCode1011, CatalogSyncUpdateResponseErrorsCode1012, CatalogSyncUpdateResponseErrorsCode1013, CatalogSyncUpdateResponseErrorsCode1014, CatalogSyncUpdateResponseErrorsCode1015, CatalogSyncUpdateResponseErrorsCode1016, CatalogSyncUpdateResponseErrorsCode1017, CatalogSyncUpdateResponseErrorsCode2001, CatalogSyncUpdateResponseErrorsCode2002, CatalogSyncUpdateResponseErrorsCode2003, CatalogSyncUpdateResponseErrorsCode2004, CatalogSyncUpdateResponseErrorsCode2005, CatalogSyncUpdateResponseErrorsCode2006, CatalogSyncUpdateResponseErrorsCode2007, CatalogSyncUpdateResponseErrorsCode2008, CatalogSyncUpdateResponseErrorsCode2009, CatalogSyncUpdateResponseErrorsCode2010, CatalogSyncUpdateResponseErrorsCode2011, CatalogSyncUpdateResponseErrorsCode2012, CatalogSyncUpdateResponseErrorsCode2013, CatalogSyncUpdateResponseErrorsCode2014, CatalogSyncUpdateResponseErrorsCode2015, CatalogSyncUpdateResponseErrorsCode2016, CatalogSyncUpdateResponseErrorsCode2017, CatalogSyncUpdateResponseErrorsCode2018, CatalogSyncUpdateResponseErrorsCode2019, CatalogSyncUpdateResponseErrorsCode2020, CatalogSyncUpdateResponseErrorsCode2021, CatalogSyncUpdateResponseErrorsCode2022, CatalogSyncUpdateResponseErrorsCode3001, CatalogSyncUpdateResponseErrorsCode3002, CatalogSyncUpdateResponseErrorsCode3003, CatalogSyncUpdateResponseErrorsCode3004, CatalogSyncUpdateResponseErrorsCode3005, CatalogSyncUpdateResponseErrorsCode3006, CatalogSyncUpdateResponseErrorsCode3007, CatalogSyncUpdateResponseErrorsCode4001, CatalogSyncUpdateResponseErrorsCode4002, CatalogSyncUpdateResponseErrorsCode4003, CatalogSyncUpdateResponseErrorsCode4004, CatalogSyncUpdateResponseErrorsCode4005, CatalogSyncUpdateResponseErrorsCode4006, CatalogSyncUpdateResponseErrorsCode4007, CatalogSyncUpdateResponseErrorsCode4008, CatalogSyncUpdateResponseErrorsCode4009, CatalogSyncUpdateResponseErrorsCode4010, CatalogSyncUpdateResponseErrorsCode4011, CatalogSyncUpdateResponseErrorsCode4012, CatalogSyncUpdateResponseErrorsCode4013, CatalogSyncUpdateResponseErrorsCode4014, CatalogSyncUpdateResponseErrorsCode4015, CatalogSyncUpdateResponseErrorsCode4016, CatalogSyncUpdateResponseErrorsCode4017, CatalogSyncUpdateResponseErrorsCode4018, CatalogSyncUpdateResponseErrorsCode4019, CatalogSyncUpdateResponseErrorsCode4020, CatalogSyncUpdateResponseErrorsCode4021, CatalogSyncUpdateResponseErrorsCode4022, CatalogSyncUpdateResponseErrorsCode4023, CatalogSyncUpdateResponseErrorsCode5001, CatalogSyncUpdateResponseErrorsCode5002, CatalogSyncUpdateResponseErrorsCode5003, CatalogSyncUpdateResponseErrorsCode5004, CatalogSyncUpdateResponseErrorsCode102000, CatalogSyncUpdateResponseErrorsCode102001, CatalogSyncUpdateResponseErrorsCode102002, CatalogSyncUpdateResponseErrorsCode102003, CatalogSyncUpdateResponseErrorsCode102004, CatalogSyncUpdateResponseErrorsCode102005, CatalogSyncUpdateResponseErrorsCode102006, CatalogSyncUpdateResponseErrorsCode102007, CatalogSyncUpdateResponseErrorsCode102008, CatalogSyncUpdateResponseErrorsCode102009, CatalogSyncUpdateResponseErrorsCode102010, CatalogSyncUpdateResponseErrorsCode102011, CatalogSyncUpdateResponseErrorsCode102012, CatalogSyncUpdateResponseErrorsCode102013, CatalogSyncUpdateResponseErrorsCode102014, CatalogSyncUpdateResponseErrorsCode102015, CatalogSyncUpdateResponseErrorsCode102016, CatalogSyncUpdateResponseErrorsCode102017, CatalogSyncUpdateResponseErrorsCode102018, CatalogSyncUpdateResponseErrorsCode102019, CatalogSyncUpdateResponseErrorsCode102020, CatalogSyncUpdateResponseErrorsCode102021, CatalogSyncUpdateResponseErrorsCode102022, CatalogSyncUpdateResponseErrorsCode102023, CatalogSyncUpdateResponseErrorsCode102024, CatalogSyncUpdateResponseErrorsCode102025, CatalogSyncUpdateResponseErrorsCode102026, CatalogSyncUpdateResponseErrorsCode102027, CatalogSyncUpdateResponseErrorsCode102028, CatalogSyncUpdateResponseErrorsCode102029, CatalogSyncUpdateResponseErrorsCode102030, CatalogSyncUpdateResponseErrorsCode102031, CatalogSyncUpdateResponseErrorsCode102032, CatalogSyncUpdateResponseErrorsCode102033, CatalogSyncUpdateResponseErrorsCode102034, CatalogSyncUpdateResponseErrorsCode102035, CatalogSyncUpdateResponseErrorsCode102036, CatalogSyncUpdateResponseErrorsCode102037, CatalogSyncUpdateResponseErrorsCode102038, CatalogSyncUpdateResponseErrorsCode102039, CatalogSyncUpdateResponseErrorsCode102040, CatalogSyncUpdateResponseErrorsCode102041, CatalogSyncUpdateResponseErrorsCode102042, CatalogSyncUpdateResponseErrorsCode102043, CatalogSyncUpdateResponseErrorsCode102044, CatalogSyncUpdateResponseErrorsCode102045, CatalogSyncUpdateResponseErrorsCode102046, CatalogSyncUpdateResponseErrorsCode102047, CatalogSyncUpdateResponseErrorsCode102048, CatalogSyncUpdateResponseErrorsCode102049, CatalogSyncUpdateResponseErrorsCode102050, CatalogSyncUpdateResponseErrorsCode102051, CatalogSyncUpdateResponseErrorsCode102052, CatalogSyncUpdateResponseErrorsCode102053, CatalogSyncUpdateResponseErrorsCode102054, CatalogSyncUpdateResponseErrorsCode102055, CatalogSyncUpdateResponseErrorsCode102056, CatalogSyncUpdateResponseErrorsCode102057, CatalogSyncUpdateResponseErrorsCode102058, CatalogSyncUpdateResponseErrorsCode102059, CatalogSyncUpdateResponseErrorsCode102060, CatalogSyncUpdateResponseErrorsCode102061, CatalogSyncUpdateResponseErrorsCode102062, CatalogSyncUpdateResponseErrorsCode102063, CatalogSyncUpdateResponseErrorsCode102064, CatalogSyncUpdateResponseErrorsCode102065, CatalogSyncUpdateResponseErrorsCode102066, CatalogSyncUpdateResponseErrorsCode103001, CatalogSyncUpdateResponseErrorsCode103002, CatalogSyncUpdateResponseErrorsCode103003, CatalogSyncUpdateResponseErrorsCode103004, CatalogSyncUpdateResponseErrorsCode103005, CatalogSyncUpdateResponseErrorsCode103006, CatalogSyncUpdateResponseErrorsCode103007, CatalogSyncUpdateResponseErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseErrorsMeta struct {
	L10nKey       string                                  `json:"l10n_key"`
	LoggableError string                                  `json:"loggable_error"`
	TemplateData  interface{}                             `json:"template_data"`
	TraceID       string                                  `json:"trace_id"`
	JSON          catalogSyncUpdateResponseErrorsMetaJSON `json:"-"`
}

// catalogSyncUpdateResponseErrorsMetaJSON contains the JSON metadata for the
// struct [CatalogSyncUpdateResponseErrorsMeta]
type catalogSyncUpdateResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseErrorsSource struct {
	Parameter           string                                    `json:"parameter"`
	ParameterValueIndex int64                                     `json:"parameter_value_index"`
	Pointer             string                                    `json:"pointer"`
	JSON                catalogSyncUpdateResponseErrorsSourceJSON `json:"-"`
}

// catalogSyncUpdateResponseErrorsSourceJSON contains the JSON metadata for the
// struct [CatalogSyncUpdateResponseErrorsSource]
type catalogSyncUpdateResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncListResponse struct {
	ID                       string                                  `json:"id,required" format:"uuid"`
	Description              string                                  `json:"description,required"`
	DestinationID            string                                  `json:"destination_id,required" format:"uuid"`
	DestinationType          CatalogSyncListResponseDestinationType  `json:"destination_type,required"`
	LastUserUpdateAt         string                                  `json:"last_user_update_at,required"`
	Name                     string                                  `json:"name,required"`
	Policy                   string                                  `json:"policy,required"`
	UpdateMode               CatalogSyncListResponseUpdateMode       `json:"update_mode,required"`
	Errors                   map[string]CatalogSyncListResponseError `json:"errors"`
	IncludesDiscoveriesUntil string                                  `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                                  `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                                  `json:"last_successful_update_at"`
	JSON                     catalogSyncListResponseJSON             `json:"-"`
}

// catalogSyncListResponseJSON contains the JSON metadata for the struct
// [CatalogSyncListResponse]
type catalogSyncListResponseJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CatalogSyncListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncListResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncListResponseDestinationType string

const (
	CatalogSyncListResponseDestinationTypeNone          CatalogSyncListResponseDestinationType = "NONE"
	CatalogSyncListResponseDestinationTypeZeroTrustList CatalogSyncListResponseDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncListResponseDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncListResponseDestinationTypeNone, CatalogSyncListResponseDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncListResponseUpdateMode string

const (
	CatalogSyncListResponseUpdateModeAuto   CatalogSyncListResponseUpdateMode = "AUTO"
	CatalogSyncListResponseUpdateModeManual CatalogSyncListResponseUpdateMode = "MANUAL"
)

func (r CatalogSyncListResponseUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncListResponseUpdateModeAuto, CatalogSyncListResponseUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncListResponseError struct {
	Code             CatalogSyncListResponseErrorsCode   `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Meta             CatalogSyncListResponseErrorsMeta   `json:"meta"`
	Source           CatalogSyncListResponseErrorsSource `json:"source"`
	JSON             catalogSyncListResponseErrorJSON    `json:"-"`
}

// catalogSyncListResponseErrorJSON contains the JSON metadata for the struct
// [CatalogSyncListResponseError]
type catalogSyncListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncListResponseErrorsCode int64

const (
	CatalogSyncListResponseErrorsCode1001   CatalogSyncListResponseErrorsCode = 1001
	CatalogSyncListResponseErrorsCode1002   CatalogSyncListResponseErrorsCode = 1002
	CatalogSyncListResponseErrorsCode1003   CatalogSyncListResponseErrorsCode = 1003
	CatalogSyncListResponseErrorsCode1004   CatalogSyncListResponseErrorsCode = 1004
	CatalogSyncListResponseErrorsCode1005   CatalogSyncListResponseErrorsCode = 1005
	CatalogSyncListResponseErrorsCode1006   CatalogSyncListResponseErrorsCode = 1006
	CatalogSyncListResponseErrorsCode1007   CatalogSyncListResponseErrorsCode = 1007
	CatalogSyncListResponseErrorsCode1008   CatalogSyncListResponseErrorsCode = 1008
	CatalogSyncListResponseErrorsCode1009   CatalogSyncListResponseErrorsCode = 1009
	CatalogSyncListResponseErrorsCode1010   CatalogSyncListResponseErrorsCode = 1010
	CatalogSyncListResponseErrorsCode1011   CatalogSyncListResponseErrorsCode = 1011
	CatalogSyncListResponseErrorsCode1012   CatalogSyncListResponseErrorsCode = 1012
	CatalogSyncListResponseErrorsCode1013   CatalogSyncListResponseErrorsCode = 1013
	CatalogSyncListResponseErrorsCode1014   CatalogSyncListResponseErrorsCode = 1014
	CatalogSyncListResponseErrorsCode1015   CatalogSyncListResponseErrorsCode = 1015
	CatalogSyncListResponseErrorsCode1016   CatalogSyncListResponseErrorsCode = 1016
	CatalogSyncListResponseErrorsCode1017   CatalogSyncListResponseErrorsCode = 1017
	CatalogSyncListResponseErrorsCode2001   CatalogSyncListResponseErrorsCode = 2001
	CatalogSyncListResponseErrorsCode2002   CatalogSyncListResponseErrorsCode = 2002
	CatalogSyncListResponseErrorsCode2003   CatalogSyncListResponseErrorsCode = 2003
	CatalogSyncListResponseErrorsCode2004   CatalogSyncListResponseErrorsCode = 2004
	CatalogSyncListResponseErrorsCode2005   CatalogSyncListResponseErrorsCode = 2005
	CatalogSyncListResponseErrorsCode2006   CatalogSyncListResponseErrorsCode = 2006
	CatalogSyncListResponseErrorsCode2007   CatalogSyncListResponseErrorsCode = 2007
	CatalogSyncListResponseErrorsCode2008   CatalogSyncListResponseErrorsCode = 2008
	CatalogSyncListResponseErrorsCode2009   CatalogSyncListResponseErrorsCode = 2009
	CatalogSyncListResponseErrorsCode2010   CatalogSyncListResponseErrorsCode = 2010
	CatalogSyncListResponseErrorsCode2011   CatalogSyncListResponseErrorsCode = 2011
	CatalogSyncListResponseErrorsCode2012   CatalogSyncListResponseErrorsCode = 2012
	CatalogSyncListResponseErrorsCode2013   CatalogSyncListResponseErrorsCode = 2013
	CatalogSyncListResponseErrorsCode2014   CatalogSyncListResponseErrorsCode = 2014
	CatalogSyncListResponseErrorsCode2015   CatalogSyncListResponseErrorsCode = 2015
	CatalogSyncListResponseErrorsCode2016   CatalogSyncListResponseErrorsCode = 2016
	CatalogSyncListResponseErrorsCode2017   CatalogSyncListResponseErrorsCode = 2017
	CatalogSyncListResponseErrorsCode2018   CatalogSyncListResponseErrorsCode = 2018
	CatalogSyncListResponseErrorsCode2019   CatalogSyncListResponseErrorsCode = 2019
	CatalogSyncListResponseErrorsCode2020   CatalogSyncListResponseErrorsCode = 2020
	CatalogSyncListResponseErrorsCode2021   CatalogSyncListResponseErrorsCode = 2021
	CatalogSyncListResponseErrorsCode2022   CatalogSyncListResponseErrorsCode = 2022
	CatalogSyncListResponseErrorsCode3001   CatalogSyncListResponseErrorsCode = 3001
	CatalogSyncListResponseErrorsCode3002   CatalogSyncListResponseErrorsCode = 3002
	CatalogSyncListResponseErrorsCode3003   CatalogSyncListResponseErrorsCode = 3003
	CatalogSyncListResponseErrorsCode3004   CatalogSyncListResponseErrorsCode = 3004
	CatalogSyncListResponseErrorsCode3005   CatalogSyncListResponseErrorsCode = 3005
	CatalogSyncListResponseErrorsCode3006   CatalogSyncListResponseErrorsCode = 3006
	CatalogSyncListResponseErrorsCode3007   CatalogSyncListResponseErrorsCode = 3007
	CatalogSyncListResponseErrorsCode4001   CatalogSyncListResponseErrorsCode = 4001
	CatalogSyncListResponseErrorsCode4002   CatalogSyncListResponseErrorsCode = 4002
	CatalogSyncListResponseErrorsCode4003   CatalogSyncListResponseErrorsCode = 4003
	CatalogSyncListResponseErrorsCode4004   CatalogSyncListResponseErrorsCode = 4004
	CatalogSyncListResponseErrorsCode4005   CatalogSyncListResponseErrorsCode = 4005
	CatalogSyncListResponseErrorsCode4006   CatalogSyncListResponseErrorsCode = 4006
	CatalogSyncListResponseErrorsCode4007   CatalogSyncListResponseErrorsCode = 4007
	CatalogSyncListResponseErrorsCode4008   CatalogSyncListResponseErrorsCode = 4008
	CatalogSyncListResponseErrorsCode4009   CatalogSyncListResponseErrorsCode = 4009
	CatalogSyncListResponseErrorsCode4010   CatalogSyncListResponseErrorsCode = 4010
	CatalogSyncListResponseErrorsCode4011   CatalogSyncListResponseErrorsCode = 4011
	CatalogSyncListResponseErrorsCode4012   CatalogSyncListResponseErrorsCode = 4012
	CatalogSyncListResponseErrorsCode4013   CatalogSyncListResponseErrorsCode = 4013
	CatalogSyncListResponseErrorsCode4014   CatalogSyncListResponseErrorsCode = 4014
	CatalogSyncListResponseErrorsCode4015   CatalogSyncListResponseErrorsCode = 4015
	CatalogSyncListResponseErrorsCode4016   CatalogSyncListResponseErrorsCode = 4016
	CatalogSyncListResponseErrorsCode4017   CatalogSyncListResponseErrorsCode = 4017
	CatalogSyncListResponseErrorsCode4018   CatalogSyncListResponseErrorsCode = 4018
	CatalogSyncListResponseErrorsCode4019   CatalogSyncListResponseErrorsCode = 4019
	CatalogSyncListResponseErrorsCode4020   CatalogSyncListResponseErrorsCode = 4020
	CatalogSyncListResponseErrorsCode4021   CatalogSyncListResponseErrorsCode = 4021
	CatalogSyncListResponseErrorsCode4022   CatalogSyncListResponseErrorsCode = 4022
	CatalogSyncListResponseErrorsCode4023   CatalogSyncListResponseErrorsCode = 4023
	CatalogSyncListResponseErrorsCode5001   CatalogSyncListResponseErrorsCode = 5001
	CatalogSyncListResponseErrorsCode5002   CatalogSyncListResponseErrorsCode = 5002
	CatalogSyncListResponseErrorsCode5003   CatalogSyncListResponseErrorsCode = 5003
	CatalogSyncListResponseErrorsCode5004   CatalogSyncListResponseErrorsCode = 5004
	CatalogSyncListResponseErrorsCode102000 CatalogSyncListResponseErrorsCode = 102000
	CatalogSyncListResponseErrorsCode102001 CatalogSyncListResponseErrorsCode = 102001
	CatalogSyncListResponseErrorsCode102002 CatalogSyncListResponseErrorsCode = 102002
	CatalogSyncListResponseErrorsCode102003 CatalogSyncListResponseErrorsCode = 102003
	CatalogSyncListResponseErrorsCode102004 CatalogSyncListResponseErrorsCode = 102004
	CatalogSyncListResponseErrorsCode102005 CatalogSyncListResponseErrorsCode = 102005
	CatalogSyncListResponseErrorsCode102006 CatalogSyncListResponseErrorsCode = 102006
	CatalogSyncListResponseErrorsCode102007 CatalogSyncListResponseErrorsCode = 102007
	CatalogSyncListResponseErrorsCode102008 CatalogSyncListResponseErrorsCode = 102008
	CatalogSyncListResponseErrorsCode102009 CatalogSyncListResponseErrorsCode = 102009
	CatalogSyncListResponseErrorsCode102010 CatalogSyncListResponseErrorsCode = 102010
	CatalogSyncListResponseErrorsCode102011 CatalogSyncListResponseErrorsCode = 102011
	CatalogSyncListResponseErrorsCode102012 CatalogSyncListResponseErrorsCode = 102012
	CatalogSyncListResponseErrorsCode102013 CatalogSyncListResponseErrorsCode = 102013
	CatalogSyncListResponseErrorsCode102014 CatalogSyncListResponseErrorsCode = 102014
	CatalogSyncListResponseErrorsCode102015 CatalogSyncListResponseErrorsCode = 102015
	CatalogSyncListResponseErrorsCode102016 CatalogSyncListResponseErrorsCode = 102016
	CatalogSyncListResponseErrorsCode102017 CatalogSyncListResponseErrorsCode = 102017
	CatalogSyncListResponseErrorsCode102018 CatalogSyncListResponseErrorsCode = 102018
	CatalogSyncListResponseErrorsCode102019 CatalogSyncListResponseErrorsCode = 102019
	CatalogSyncListResponseErrorsCode102020 CatalogSyncListResponseErrorsCode = 102020
	CatalogSyncListResponseErrorsCode102021 CatalogSyncListResponseErrorsCode = 102021
	CatalogSyncListResponseErrorsCode102022 CatalogSyncListResponseErrorsCode = 102022
	CatalogSyncListResponseErrorsCode102023 CatalogSyncListResponseErrorsCode = 102023
	CatalogSyncListResponseErrorsCode102024 CatalogSyncListResponseErrorsCode = 102024
	CatalogSyncListResponseErrorsCode102025 CatalogSyncListResponseErrorsCode = 102025
	CatalogSyncListResponseErrorsCode102026 CatalogSyncListResponseErrorsCode = 102026
	CatalogSyncListResponseErrorsCode102027 CatalogSyncListResponseErrorsCode = 102027
	CatalogSyncListResponseErrorsCode102028 CatalogSyncListResponseErrorsCode = 102028
	CatalogSyncListResponseErrorsCode102029 CatalogSyncListResponseErrorsCode = 102029
	CatalogSyncListResponseErrorsCode102030 CatalogSyncListResponseErrorsCode = 102030
	CatalogSyncListResponseErrorsCode102031 CatalogSyncListResponseErrorsCode = 102031
	CatalogSyncListResponseErrorsCode102032 CatalogSyncListResponseErrorsCode = 102032
	CatalogSyncListResponseErrorsCode102033 CatalogSyncListResponseErrorsCode = 102033
	CatalogSyncListResponseErrorsCode102034 CatalogSyncListResponseErrorsCode = 102034
	CatalogSyncListResponseErrorsCode102035 CatalogSyncListResponseErrorsCode = 102035
	CatalogSyncListResponseErrorsCode102036 CatalogSyncListResponseErrorsCode = 102036
	CatalogSyncListResponseErrorsCode102037 CatalogSyncListResponseErrorsCode = 102037
	CatalogSyncListResponseErrorsCode102038 CatalogSyncListResponseErrorsCode = 102038
	CatalogSyncListResponseErrorsCode102039 CatalogSyncListResponseErrorsCode = 102039
	CatalogSyncListResponseErrorsCode102040 CatalogSyncListResponseErrorsCode = 102040
	CatalogSyncListResponseErrorsCode102041 CatalogSyncListResponseErrorsCode = 102041
	CatalogSyncListResponseErrorsCode102042 CatalogSyncListResponseErrorsCode = 102042
	CatalogSyncListResponseErrorsCode102043 CatalogSyncListResponseErrorsCode = 102043
	CatalogSyncListResponseErrorsCode102044 CatalogSyncListResponseErrorsCode = 102044
	CatalogSyncListResponseErrorsCode102045 CatalogSyncListResponseErrorsCode = 102045
	CatalogSyncListResponseErrorsCode102046 CatalogSyncListResponseErrorsCode = 102046
	CatalogSyncListResponseErrorsCode102047 CatalogSyncListResponseErrorsCode = 102047
	CatalogSyncListResponseErrorsCode102048 CatalogSyncListResponseErrorsCode = 102048
	CatalogSyncListResponseErrorsCode102049 CatalogSyncListResponseErrorsCode = 102049
	CatalogSyncListResponseErrorsCode102050 CatalogSyncListResponseErrorsCode = 102050
	CatalogSyncListResponseErrorsCode102051 CatalogSyncListResponseErrorsCode = 102051
	CatalogSyncListResponseErrorsCode102052 CatalogSyncListResponseErrorsCode = 102052
	CatalogSyncListResponseErrorsCode102053 CatalogSyncListResponseErrorsCode = 102053
	CatalogSyncListResponseErrorsCode102054 CatalogSyncListResponseErrorsCode = 102054
	CatalogSyncListResponseErrorsCode102055 CatalogSyncListResponseErrorsCode = 102055
	CatalogSyncListResponseErrorsCode102056 CatalogSyncListResponseErrorsCode = 102056
	CatalogSyncListResponseErrorsCode102057 CatalogSyncListResponseErrorsCode = 102057
	CatalogSyncListResponseErrorsCode102058 CatalogSyncListResponseErrorsCode = 102058
	CatalogSyncListResponseErrorsCode102059 CatalogSyncListResponseErrorsCode = 102059
	CatalogSyncListResponseErrorsCode102060 CatalogSyncListResponseErrorsCode = 102060
	CatalogSyncListResponseErrorsCode102061 CatalogSyncListResponseErrorsCode = 102061
	CatalogSyncListResponseErrorsCode102062 CatalogSyncListResponseErrorsCode = 102062
	CatalogSyncListResponseErrorsCode102063 CatalogSyncListResponseErrorsCode = 102063
	CatalogSyncListResponseErrorsCode102064 CatalogSyncListResponseErrorsCode = 102064
	CatalogSyncListResponseErrorsCode102065 CatalogSyncListResponseErrorsCode = 102065
	CatalogSyncListResponseErrorsCode102066 CatalogSyncListResponseErrorsCode = 102066
	CatalogSyncListResponseErrorsCode103001 CatalogSyncListResponseErrorsCode = 103001
	CatalogSyncListResponseErrorsCode103002 CatalogSyncListResponseErrorsCode = 103002
	CatalogSyncListResponseErrorsCode103003 CatalogSyncListResponseErrorsCode = 103003
	CatalogSyncListResponseErrorsCode103004 CatalogSyncListResponseErrorsCode = 103004
	CatalogSyncListResponseErrorsCode103005 CatalogSyncListResponseErrorsCode = 103005
	CatalogSyncListResponseErrorsCode103006 CatalogSyncListResponseErrorsCode = 103006
	CatalogSyncListResponseErrorsCode103007 CatalogSyncListResponseErrorsCode = 103007
	CatalogSyncListResponseErrorsCode103008 CatalogSyncListResponseErrorsCode = 103008
)

func (r CatalogSyncListResponseErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncListResponseErrorsCode1001, CatalogSyncListResponseErrorsCode1002, CatalogSyncListResponseErrorsCode1003, CatalogSyncListResponseErrorsCode1004, CatalogSyncListResponseErrorsCode1005, CatalogSyncListResponseErrorsCode1006, CatalogSyncListResponseErrorsCode1007, CatalogSyncListResponseErrorsCode1008, CatalogSyncListResponseErrorsCode1009, CatalogSyncListResponseErrorsCode1010, CatalogSyncListResponseErrorsCode1011, CatalogSyncListResponseErrorsCode1012, CatalogSyncListResponseErrorsCode1013, CatalogSyncListResponseErrorsCode1014, CatalogSyncListResponseErrorsCode1015, CatalogSyncListResponseErrorsCode1016, CatalogSyncListResponseErrorsCode1017, CatalogSyncListResponseErrorsCode2001, CatalogSyncListResponseErrorsCode2002, CatalogSyncListResponseErrorsCode2003, CatalogSyncListResponseErrorsCode2004, CatalogSyncListResponseErrorsCode2005, CatalogSyncListResponseErrorsCode2006, CatalogSyncListResponseErrorsCode2007, CatalogSyncListResponseErrorsCode2008, CatalogSyncListResponseErrorsCode2009, CatalogSyncListResponseErrorsCode2010, CatalogSyncListResponseErrorsCode2011, CatalogSyncListResponseErrorsCode2012, CatalogSyncListResponseErrorsCode2013, CatalogSyncListResponseErrorsCode2014, CatalogSyncListResponseErrorsCode2015, CatalogSyncListResponseErrorsCode2016, CatalogSyncListResponseErrorsCode2017, CatalogSyncListResponseErrorsCode2018, CatalogSyncListResponseErrorsCode2019, CatalogSyncListResponseErrorsCode2020, CatalogSyncListResponseErrorsCode2021, CatalogSyncListResponseErrorsCode2022, CatalogSyncListResponseErrorsCode3001, CatalogSyncListResponseErrorsCode3002, CatalogSyncListResponseErrorsCode3003, CatalogSyncListResponseErrorsCode3004, CatalogSyncListResponseErrorsCode3005, CatalogSyncListResponseErrorsCode3006, CatalogSyncListResponseErrorsCode3007, CatalogSyncListResponseErrorsCode4001, CatalogSyncListResponseErrorsCode4002, CatalogSyncListResponseErrorsCode4003, CatalogSyncListResponseErrorsCode4004, CatalogSyncListResponseErrorsCode4005, CatalogSyncListResponseErrorsCode4006, CatalogSyncListResponseErrorsCode4007, CatalogSyncListResponseErrorsCode4008, CatalogSyncListResponseErrorsCode4009, CatalogSyncListResponseErrorsCode4010, CatalogSyncListResponseErrorsCode4011, CatalogSyncListResponseErrorsCode4012, CatalogSyncListResponseErrorsCode4013, CatalogSyncListResponseErrorsCode4014, CatalogSyncListResponseErrorsCode4015, CatalogSyncListResponseErrorsCode4016, CatalogSyncListResponseErrorsCode4017, CatalogSyncListResponseErrorsCode4018, CatalogSyncListResponseErrorsCode4019, CatalogSyncListResponseErrorsCode4020, CatalogSyncListResponseErrorsCode4021, CatalogSyncListResponseErrorsCode4022, CatalogSyncListResponseErrorsCode4023, CatalogSyncListResponseErrorsCode5001, CatalogSyncListResponseErrorsCode5002, CatalogSyncListResponseErrorsCode5003, CatalogSyncListResponseErrorsCode5004, CatalogSyncListResponseErrorsCode102000, CatalogSyncListResponseErrorsCode102001, CatalogSyncListResponseErrorsCode102002, CatalogSyncListResponseErrorsCode102003, CatalogSyncListResponseErrorsCode102004, CatalogSyncListResponseErrorsCode102005, CatalogSyncListResponseErrorsCode102006, CatalogSyncListResponseErrorsCode102007, CatalogSyncListResponseErrorsCode102008, CatalogSyncListResponseErrorsCode102009, CatalogSyncListResponseErrorsCode102010, CatalogSyncListResponseErrorsCode102011, CatalogSyncListResponseErrorsCode102012, CatalogSyncListResponseErrorsCode102013, CatalogSyncListResponseErrorsCode102014, CatalogSyncListResponseErrorsCode102015, CatalogSyncListResponseErrorsCode102016, CatalogSyncListResponseErrorsCode102017, CatalogSyncListResponseErrorsCode102018, CatalogSyncListResponseErrorsCode102019, CatalogSyncListResponseErrorsCode102020, CatalogSyncListResponseErrorsCode102021, CatalogSyncListResponseErrorsCode102022, CatalogSyncListResponseErrorsCode102023, CatalogSyncListResponseErrorsCode102024, CatalogSyncListResponseErrorsCode102025, CatalogSyncListResponseErrorsCode102026, CatalogSyncListResponseErrorsCode102027, CatalogSyncListResponseErrorsCode102028, CatalogSyncListResponseErrorsCode102029, CatalogSyncListResponseErrorsCode102030, CatalogSyncListResponseErrorsCode102031, CatalogSyncListResponseErrorsCode102032, CatalogSyncListResponseErrorsCode102033, CatalogSyncListResponseErrorsCode102034, CatalogSyncListResponseErrorsCode102035, CatalogSyncListResponseErrorsCode102036, CatalogSyncListResponseErrorsCode102037, CatalogSyncListResponseErrorsCode102038, CatalogSyncListResponseErrorsCode102039, CatalogSyncListResponseErrorsCode102040, CatalogSyncListResponseErrorsCode102041, CatalogSyncListResponseErrorsCode102042, CatalogSyncListResponseErrorsCode102043, CatalogSyncListResponseErrorsCode102044, CatalogSyncListResponseErrorsCode102045, CatalogSyncListResponseErrorsCode102046, CatalogSyncListResponseErrorsCode102047, CatalogSyncListResponseErrorsCode102048, CatalogSyncListResponseErrorsCode102049, CatalogSyncListResponseErrorsCode102050, CatalogSyncListResponseErrorsCode102051, CatalogSyncListResponseErrorsCode102052, CatalogSyncListResponseErrorsCode102053, CatalogSyncListResponseErrorsCode102054, CatalogSyncListResponseErrorsCode102055, CatalogSyncListResponseErrorsCode102056, CatalogSyncListResponseErrorsCode102057, CatalogSyncListResponseErrorsCode102058, CatalogSyncListResponseErrorsCode102059, CatalogSyncListResponseErrorsCode102060, CatalogSyncListResponseErrorsCode102061, CatalogSyncListResponseErrorsCode102062, CatalogSyncListResponseErrorsCode102063, CatalogSyncListResponseErrorsCode102064, CatalogSyncListResponseErrorsCode102065, CatalogSyncListResponseErrorsCode102066, CatalogSyncListResponseErrorsCode103001, CatalogSyncListResponseErrorsCode103002, CatalogSyncListResponseErrorsCode103003, CatalogSyncListResponseErrorsCode103004, CatalogSyncListResponseErrorsCode103005, CatalogSyncListResponseErrorsCode103006, CatalogSyncListResponseErrorsCode103007, CatalogSyncListResponseErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncListResponseErrorsMeta struct {
	L10nKey       string                                `json:"l10n_key"`
	LoggableError string                                `json:"loggable_error"`
	TemplateData  interface{}                           `json:"template_data"`
	TraceID       string                                `json:"trace_id"`
	JSON          catalogSyncListResponseErrorsMetaJSON `json:"-"`
}

// catalogSyncListResponseErrorsMetaJSON contains the JSON metadata for the struct
// [CatalogSyncListResponseErrorsMeta]
type catalogSyncListResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncListResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncListResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncListResponseErrorsSource struct {
	Parameter           string                                  `json:"parameter"`
	ParameterValueIndex int64                                   `json:"parameter_value_index"`
	Pointer             string                                  `json:"pointer"`
	JSON                catalogSyncListResponseErrorsSourceJSON `json:"-"`
}

// catalogSyncListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [CatalogSyncListResponseErrorsSource]
type catalogSyncListResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponse struct {
	ID   string                        `json:"id,required" format:"uuid"`
	JSON catalogSyncDeleteResponseJSON `json:"-"`
}

// catalogSyncDeleteResponseJSON contains the JSON metadata for the struct
// [CatalogSyncDeleteResponse]
type catalogSyncDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponse struct {
	ID                       string                                  `json:"id,required" format:"uuid"`
	Description              string                                  `json:"description,required"`
	DestinationID            string                                  `json:"destination_id,required" format:"uuid"`
	DestinationType          CatalogSyncEditResponseDestinationType  `json:"destination_type,required"`
	LastUserUpdateAt         string                                  `json:"last_user_update_at,required"`
	Name                     string                                  `json:"name,required"`
	Policy                   string                                  `json:"policy,required"`
	UpdateMode               CatalogSyncEditResponseUpdateMode       `json:"update_mode,required"`
	Errors                   map[string]CatalogSyncEditResponseError `json:"errors"`
	IncludesDiscoveriesUntil string                                  `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                                  `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                                  `json:"last_successful_update_at"`
	JSON                     catalogSyncEditResponseJSON             `json:"-"`
}

// catalogSyncEditResponseJSON contains the JSON metadata for the struct
// [CatalogSyncEditResponse]
type catalogSyncEditResponseJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CatalogSyncEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseDestinationType string

const (
	CatalogSyncEditResponseDestinationTypeNone          CatalogSyncEditResponseDestinationType = "NONE"
	CatalogSyncEditResponseDestinationTypeZeroTrustList CatalogSyncEditResponseDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncEditResponseDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncEditResponseDestinationTypeNone, CatalogSyncEditResponseDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncEditResponseUpdateMode string

const (
	CatalogSyncEditResponseUpdateModeAuto   CatalogSyncEditResponseUpdateMode = "AUTO"
	CatalogSyncEditResponseUpdateModeManual CatalogSyncEditResponseUpdateMode = "MANUAL"
)

func (r CatalogSyncEditResponseUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncEditResponseUpdateModeAuto, CatalogSyncEditResponseUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncEditResponseError struct {
	Code             CatalogSyncEditResponseErrorsCode   `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Meta             CatalogSyncEditResponseErrorsMeta   `json:"meta"`
	Source           CatalogSyncEditResponseErrorsSource `json:"source"`
	JSON             catalogSyncEditResponseErrorJSON    `json:"-"`
}

// catalogSyncEditResponseErrorJSON contains the JSON metadata for the struct
// [CatalogSyncEditResponseError]
type catalogSyncEditResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncEditResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseErrorsCode int64

const (
	CatalogSyncEditResponseErrorsCode1001   CatalogSyncEditResponseErrorsCode = 1001
	CatalogSyncEditResponseErrorsCode1002   CatalogSyncEditResponseErrorsCode = 1002
	CatalogSyncEditResponseErrorsCode1003   CatalogSyncEditResponseErrorsCode = 1003
	CatalogSyncEditResponseErrorsCode1004   CatalogSyncEditResponseErrorsCode = 1004
	CatalogSyncEditResponseErrorsCode1005   CatalogSyncEditResponseErrorsCode = 1005
	CatalogSyncEditResponseErrorsCode1006   CatalogSyncEditResponseErrorsCode = 1006
	CatalogSyncEditResponseErrorsCode1007   CatalogSyncEditResponseErrorsCode = 1007
	CatalogSyncEditResponseErrorsCode1008   CatalogSyncEditResponseErrorsCode = 1008
	CatalogSyncEditResponseErrorsCode1009   CatalogSyncEditResponseErrorsCode = 1009
	CatalogSyncEditResponseErrorsCode1010   CatalogSyncEditResponseErrorsCode = 1010
	CatalogSyncEditResponseErrorsCode1011   CatalogSyncEditResponseErrorsCode = 1011
	CatalogSyncEditResponseErrorsCode1012   CatalogSyncEditResponseErrorsCode = 1012
	CatalogSyncEditResponseErrorsCode1013   CatalogSyncEditResponseErrorsCode = 1013
	CatalogSyncEditResponseErrorsCode1014   CatalogSyncEditResponseErrorsCode = 1014
	CatalogSyncEditResponseErrorsCode1015   CatalogSyncEditResponseErrorsCode = 1015
	CatalogSyncEditResponseErrorsCode1016   CatalogSyncEditResponseErrorsCode = 1016
	CatalogSyncEditResponseErrorsCode1017   CatalogSyncEditResponseErrorsCode = 1017
	CatalogSyncEditResponseErrorsCode2001   CatalogSyncEditResponseErrorsCode = 2001
	CatalogSyncEditResponseErrorsCode2002   CatalogSyncEditResponseErrorsCode = 2002
	CatalogSyncEditResponseErrorsCode2003   CatalogSyncEditResponseErrorsCode = 2003
	CatalogSyncEditResponseErrorsCode2004   CatalogSyncEditResponseErrorsCode = 2004
	CatalogSyncEditResponseErrorsCode2005   CatalogSyncEditResponseErrorsCode = 2005
	CatalogSyncEditResponseErrorsCode2006   CatalogSyncEditResponseErrorsCode = 2006
	CatalogSyncEditResponseErrorsCode2007   CatalogSyncEditResponseErrorsCode = 2007
	CatalogSyncEditResponseErrorsCode2008   CatalogSyncEditResponseErrorsCode = 2008
	CatalogSyncEditResponseErrorsCode2009   CatalogSyncEditResponseErrorsCode = 2009
	CatalogSyncEditResponseErrorsCode2010   CatalogSyncEditResponseErrorsCode = 2010
	CatalogSyncEditResponseErrorsCode2011   CatalogSyncEditResponseErrorsCode = 2011
	CatalogSyncEditResponseErrorsCode2012   CatalogSyncEditResponseErrorsCode = 2012
	CatalogSyncEditResponseErrorsCode2013   CatalogSyncEditResponseErrorsCode = 2013
	CatalogSyncEditResponseErrorsCode2014   CatalogSyncEditResponseErrorsCode = 2014
	CatalogSyncEditResponseErrorsCode2015   CatalogSyncEditResponseErrorsCode = 2015
	CatalogSyncEditResponseErrorsCode2016   CatalogSyncEditResponseErrorsCode = 2016
	CatalogSyncEditResponseErrorsCode2017   CatalogSyncEditResponseErrorsCode = 2017
	CatalogSyncEditResponseErrorsCode2018   CatalogSyncEditResponseErrorsCode = 2018
	CatalogSyncEditResponseErrorsCode2019   CatalogSyncEditResponseErrorsCode = 2019
	CatalogSyncEditResponseErrorsCode2020   CatalogSyncEditResponseErrorsCode = 2020
	CatalogSyncEditResponseErrorsCode2021   CatalogSyncEditResponseErrorsCode = 2021
	CatalogSyncEditResponseErrorsCode2022   CatalogSyncEditResponseErrorsCode = 2022
	CatalogSyncEditResponseErrorsCode3001   CatalogSyncEditResponseErrorsCode = 3001
	CatalogSyncEditResponseErrorsCode3002   CatalogSyncEditResponseErrorsCode = 3002
	CatalogSyncEditResponseErrorsCode3003   CatalogSyncEditResponseErrorsCode = 3003
	CatalogSyncEditResponseErrorsCode3004   CatalogSyncEditResponseErrorsCode = 3004
	CatalogSyncEditResponseErrorsCode3005   CatalogSyncEditResponseErrorsCode = 3005
	CatalogSyncEditResponseErrorsCode3006   CatalogSyncEditResponseErrorsCode = 3006
	CatalogSyncEditResponseErrorsCode3007   CatalogSyncEditResponseErrorsCode = 3007
	CatalogSyncEditResponseErrorsCode4001   CatalogSyncEditResponseErrorsCode = 4001
	CatalogSyncEditResponseErrorsCode4002   CatalogSyncEditResponseErrorsCode = 4002
	CatalogSyncEditResponseErrorsCode4003   CatalogSyncEditResponseErrorsCode = 4003
	CatalogSyncEditResponseErrorsCode4004   CatalogSyncEditResponseErrorsCode = 4004
	CatalogSyncEditResponseErrorsCode4005   CatalogSyncEditResponseErrorsCode = 4005
	CatalogSyncEditResponseErrorsCode4006   CatalogSyncEditResponseErrorsCode = 4006
	CatalogSyncEditResponseErrorsCode4007   CatalogSyncEditResponseErrorsCode = 4007
	CatalogSyncEditResponseErrorsCode4008   CatalogSyncEditResponseErrorsCode = 4008
	CatalogSyncEditResponseErrorsCode4009   CatalogSyncEditResponseErrorsCode = 4009
	CatalogSyncEditResponseErrorsCode4010   CatalogSyncEditResponseErrorsCode = 4010
	CatalogSyncEditResponseErrorsCode4011   CatalogSyncEditResponseErrorsCode = 4011
	CatalogSyncEditResponseErrorsCode4012   CatalogSyncEditResponseErrorsCode = 4012
	CatalogSyncEditResponseErrorsCode4013   CatalogSyncEditResponseErrorsCode = 4013
	CatalogSyncEditResponseErrorsCode4014   CatalogSyncEditResponseErrorsCode = 4014
	CatalogSyncEditResponseErrorsCode4015   CatalogSyncEditResponseErrorsCode = 4015
	CatalogSyncEditResponseErrorsCode4016   CatalogSyncEditResponseErrorsCode = 4016
	CatalogSyncEditResponseErrorsCode4017   CatalogSyncEditResponseErrorsCode = 4017
	CatalogSyncEditResponseErrorsCode4018   CatalogSyncEditResponseErrorsCode = 4018
	CatalogSyncEditResponseErrorsCode4019   CatalogSyncEditResponseErrorsCode = 4019
	CatalogSyncEditResponseErrorsCode4020   CatalogSyncEditResponseErrorsCode = 4020
	CatalogSyncEditResponseErrorsCode4021   CatalogSyncEditResponseErrorsCode = 4021
	CatalogSyncEditResponseErrorsCode4022   CatalogSyncEditResponseErrorsCode = 4022
	CatalogSyncEditResponseErrorsCode4023   CatalogSyncEditResponseErrorsCode = 4023
	CatalogSyncEditResponseErrorsCode5001   CatalogSyncEditResponseErrorsCode = 5001
	CatalogSyncEditResponseErrorsCode5002   CatalogSyncEditResponseErrorsCode = 5002
	CatalogSyncEditResponseErrorsCode5003   CatalogSyncEditResponseErrorsCode = 5003
	CatalogSyncEditResponseErrorsCode5004   CatalogSyncEditResponseErrorsCode = 5004
	CatalogSyncEditResponseErrorsCode102000 CatalogSyncEditResponseErrorsCode = 102000
	CatalogSyncEditResponseErrorsCode102001 CatalogSyncEditResponseErrorsCode = 102001
	CatalogSyncEditResponseErrorsCode102002 CatalogSyncEditResponseErrorsCode = 102002
	CatalogSyncEditResponseErrorsCode102003 CatalogSyncEditResponseErrorsCode = 102003
	CatalogSyncEditResponseErrorsCode102004 CatalogSyncEditResponseErrorsCode = 102004
	CatalogSyncEditResponseErrorsCode102005 CatalogSyncEditResponseErrorsCode = 102005
	CatalogSyncEditResponseErrorsCode102006 CatalogSyncEditResponseErrorsCode = 102006
	CatalogSyncEditResponseErrorsCode102007 CatalogSyncEditResponseErrorsCode = 102007
	CatalogSyncEditResponseErrorsCode102008 CatalogSyncEditResponseErrorsCode = 102008
	CatalogSyncEditResponseErrorsCode102009 CatalogSyncEditResponseErrorsCode = 102009
	CatalogSyncEditResponseErrorsCode102010 CatalogSyncEditResponseErrorsCode = 102010
	CatalogSyncEditResponseErrorsCode102011 CatalogSyncEditResponseErrorsCode = 102011
	CatalogSyncEditResponseErrorsCode102012 CatalogSyncEditResponseErrorsCode = 102012
	CatalogSyncEditResponseErrorsCode102013 CatalogSyncEditResponseErrorsCode = 102013
	CatalogSyncEditResponseErrorsCode102014 CatalogSyncEditResponseErrorsCode = 102014
	CatalogSyncEditResponseErrorsCode102015 CatalogSyncEditResponseErrorsCode = 102015
	CatalogSyncEditResponseErrorsCode102016 CatalogSyncEditResponseErrorsCode = 102016
	CatalogSyncEditResponseErrorsCode102017 CatalogSyncEditResponseErrorsCode = 102017
	CatalogSyncEditResponseErrorsCode102018 CatalogSyncEditResponseErrorsCode = 102018
	CatalogSyncEditResponseErrorsCode102019 CatalogSyncEditResponseErrorsCode = 102019
	CatalogSyncEditResponseErrorsCode102020 CatalogSyncEditResponseErrorsCode = 102020
	CatalogSyncEditResponseErrorsCode102021 CatalogSyncEditResponseErrorsCode = 102021
	CatalogSyncEditResponseErrorsCode102022 CatalogSyncEditResponseErrorsCode = 102022
	CatalogSyncEditResponseErrorsCode102023 CatalogSyncEditResponseErrorsCode = 102023
	CatalogSyncEditResponseErrorsCode102024 CatalogSyncEditResponseErrorsCode = 102024
	CatalogSyncEditResponseErrorsCode102025 CatalogSyncEditResponseErrorsCode = 102025
	CatalogSyncEditResponseErrorsCode102026 CatalogSyncEditResponseErrorsCode = 102026
	CatalogSyncEditResponseErrorsCode102027 CatalogSyncEditResponseErrorsCode = 102027
	CatalogSyncEditResponseErrorsCode102028 CatalogSyncEditResponseErrorsCode = 102028
	CatalogSyncEditResponseErrorsCode102029 CatalogSyncEditResponseErrorsCode = 102029
	CatalogSyncEditResponseErrorsCode102030 CatalogSyncEditResponseErrorsCode = 102030
	CatalogSyncEditResponseErrorsCode102031 CatalogSyncEditResponseErrorsCode = 102031
	CatalogSyncEditResponseErrorsCode102032 CatalogSyncEditResponseErrorsCode = 102032
	CatalogSyncEditResponseErrorsCode102033 CatalogSyncEditResponseErrorsCode = 102033
	CatalogSyncEditResponseErrorsCode102034 CatalogSyncEditResponseErrorsCode = 102034
	CatalogSyncEditResponseErrorsCode102035 CatalogSyncEditResponseErrorsCode = 102035
	CatalogSyncEditResponseErrorsCode102036 CatalogSyncEditResponseErrorsCode = 102036
	CatalogSyncEditResponseErrorsCode102037 CatalogSyncEditResponseErrorsCode = 102037
	CatalogSyncEditResponseErrorsCode102038 CatalogSyncEditResponseErrorsCode = 102038
	CatalogSyncEditResponseErrorsCode102039 CatalogSyncEditResponseErrorsCode = 102039
	CatalogSyncEditResponseErrorsCode102040 CatalogSyncEditResponseErrorsCode = 102040
	CatalogSyncEditResponseErrorsCode102041 CatalogSyncEditResponseErrorsCode = 102041
	CatalogSyncEditResponseErrorsCode102042 CatalogSyncEditResponseErrorsCode = 102042
	CatalogSyncEditResponseErrorsCode102043 CatalogSyncEditResponseErrorsCode = 102043
	CatalogSyncEditResponseErrorsCode102044 CatalogSyncEditResponseErrorsCode = 102044
	CatalogSyncEditResponseErrorsCode102045 CatalogSyncEditResponseErrorsCode = 102045
	CatalogSyncEditResponseErrorsCode102046 CatalogSyncEditResponseErrorsCode = 102046
	CatalogSyncEditResponseErrorsCode102047 CatalogSyncEditResponseErrorsCode = 102047
	CatalogSyncEditResponseErrorsCode102048 CatalogSyncEditResponseErrorsCode = 102048
	CatalogSyncEditResponseErrorsCode102049 CatalogSyncEditResponseErrorsCode = 102049
	CatalogSyncEditResponseErrorsCode102050 CatalogSyncEditResponseErrorsCode = 102050
	CatalogSyncEditResponseErrorsCode102051 CatalogSyncEditResponseErrorsCode = 102051
	CatalogSyncEditResponseErrorsCode102052 CatalogSyncEditResponseErrorsCode = 102052
	CatalogSyncEditResponseErrorsCode102053 CatalogSyncEditResponseErrorsCode = 102053
	CatalogSyncEditResponseErrorsCode102054 CatalogSyncEditResponseErrorsCode = 102054
	CatalogSyncEditResponseErrorsCode102055 CatalogSyncEditResponseErrorsCode = 102055
	CatalogSyncEditResponseErrorsCode102056 CatalogSyncEditResponseErrorsCode = 102056
	CatalogSyncEditResponseErrorsCode102057 CatalogSyncEditResponseErrorsCode = 102057
	CatalogSyncEditResponseErrorsCode102058 CatalogSyncEditResponseErrorsCode = 102058
	CatalogSyncEditResponseErrorsCode102059 CatalogSyncEditResponseErrorsCode = 102059
	CatalogSyncEditResponseErrorsCode102060 CatalogSyncEditResponseErrorsCode = 102060
	CatalogSyncEditResponseErrorsCode102061 CatalogSyncEditResponseErrorsCode = 102061
	CatalogSyncEditResponseErrorsCode102062 CatalogSyncEditResponseErrorsCode = 102062
	CatalogSyncEditResponseErrorsCode102063 CatalogSyncEditResponseErrorsCode = 102063
	CatalogSyncEditResponseErrorsCode102064 CatalogSyncEditResponseErrorsCode = 102064
	CatalogSyncEditResponseErrorsCode102065 CatalogSyncEditResponseErrorsCode = 102065
	CatalogSyncEditResponseErrorsCode102066 CatalogSyncEditResponseErrorsCode = 102066
	CatalogSyncEditResponseErrorsCode103001 CatalogSyncEditResponseErrorsCode = 103001
	CatalogSyncEditResponseErrorsCode103002 CatalogSyncEditResponseErrorsCode = 103002
	CatalogSyncEditResponseErrorsCode103003 CatalogSyncEditResponseErrorsCode = 103003
	CatalogSyncEditResponseErrorsCode103004 CatalogSyncEditResponseErrorsCode = 103004
	CatalogSyncEditResponseErrorsCode103005 CatalogSyncEditResponseErrorsCode = 103005
	CatalogSyncEditResponseErrorsCode103006 CatalogSyncEditResponseErrorsCode = 103006
	CatalogSyncEditResponseErrorsCode103007 CatalogSyncEditResponseErrorsCode = 103007
	CatalogSyncEditResponseErrorsCode103008 CatalogSyncEditResponseErrorsCode = 103008
)

func (r CatalogSyncEditResponseErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncEditResponseErrorsCode1001, CatalogSyncEditResponseErrorsCode1002, CatalogSyncEditResponseErrorsCode1003, CatalogSyncEditResponseErrorsCode1004, CatalogSyncEditResponseErrorsCode1005, CatalogSyncEditResponseErrorsCode1006, CatalogSyncEditResponseErrorsCode1007, CatalogSyncEditResponseErrorsCode1008, CatalogSyncEditResponseErrorsCode1009, CatalogSyncEditResponseErrorsCode1010, CatalogSyncEditResponseErrorsCode1011, CatalogSyncEditResponseErrorsCode1012, CatalogSyncEditResponseErrorsCode1013, CatalogSyncEditResponseErrorsCode1014, CatalogSyncEditResponseErrorsCode1015, CatalogSyncEditResponseErrorsCode1016, CatalogSyncEditResponseErrorsCode1017, CatalogSyncEditResponseErrorsCode2001, CatalogSyncEditResponseErrorsCode2002, CatalogSyncEditResponseErrorsCode2003, CatalogSyncEditResponseErrorsCode2004, CatalogSyncEditResponseErrorsCode2005, CatalogSyncEditResponseErrorsCode2006, CatalogSyncEditResponseErrorsCode2007, CatalogSyncEditResponseErrorsCode2008, CatalogSyncEditResponseErrorsCode2009, CatalogSyncEditResponseErrorsCode2010, CatalogSyncEditResponseErrorsCode2011, CatalogSyncEditResponseErrorsCode2012, CatalogSyncEditResponseErrorsCode2013, CatalogSyncEditResponseErrorsCode2014, CatalogSyncEditResponseErrorsCode2015, CatalogSyncEditResponseErrorsCode2016, CatalogSyncEditResponseErrorsCode2017, CatalogSyncEditResponseErrorsCode2018, CatalogSyncEditResponseErrorsCode2019, CatalogSyncEditResponseErrorsCode2020, CatalogSyncEditResponseErrorsCode2021, CatalogSyncEditResponseErrorsCode2022, CatalogSyncEditResponseErrorsCode3001, CatalogSyncEditResponseErrorsCode3002, CatalogSyncEditResponseErrorsCode3003, CatalogSyncEditResponseErrorsCode3004, CatalogSyncEditResponseErrorsCode3005, CatalogSyncEditResponseErrorsCode3006, CatalogSyncEditResponseErrorsCode3007, CatalogSyncEditResponseErrorsCode4001, CatalogSyncEditResponseErrorsCode4002, CatalogSyncEditResponseErrorsCode4003, CatalogSyncEditResponseErrorsCode4004, CatalogSyncEditResponseErrorsCode4005, CatalogSyncEditResponseErrorsCode4006, CatalogSyncEditResponseErrorsCode4007, CatalogSyncEditResponseErrorsCode4008, CatalogSyncEditResponseErrorsCode4009, CatalogSyncEditResponseErrorsCode4010, CatalogSyncEditResponseErrorsCode4011, CatalogSyncEditResponseErrorsCode4012, CatalogSyncEditResponseErrorsCode4013, CatalogSyncEditResponseErrorsCode4014, CatalogSyncEditResponseErrorsCode4015, CatalogSyncEditResponseErrorsCode4016, CatalogSyncEditResponseErrorsCode4017, CatalogSyncEditResponseErrorsCode4018, CatalogSyncEditResponseErrorsCode4019, CatalogSyncEditResponseErrorsCode4020, CatalogSyncEditResponseErrorsCode4021, CatalogSyncEditResponseErrorsCode4022, CatalogSyncEditResponseErrorsCode4023, CatalogSyncEditResponseErrorsCode5001, CatalogSyncEditResponseErrorsCode5002, CatalogSyncEditResponseErrorsCode5003, CatalogSyncEditResponseErrorsCode5004, CatalogSyncEditResponseErrorsCode102000, CatalogSyncEditResponseErrorsCode102001, CatalogSyncEditResponseErrorsCode102002, CatalogSyncEditResponseErrorsCode102003, CatalogSyncEditResponseErrorsCode102004, CatalogSyncEditResponseErrorsCode102005, CatalogSyncEditResponseErrorsCode102006, CatalogSyncEditResponseErrorsCode102007, CatalogSyncEditResponseErrorsCode102008, CatalogSyncEditResponseErrorsCode102009, CatalogSyncEditResponseErrorsCode102010, CatalogSyncEditResponseErrorsCode102011, CatalogSyncEditResponseErrorsCode102012, CatalogSyncEditResponseErrorsCode102013, CatalogSyncEditResponseErrorsCode102014, CatalogSyncEditResponseErrorsCode102015, CatalogSyncEditResponseErrorsCode102016, CatalogSyncEditResponseErrorsCode102017, CatalogSyncEditResponseErrorsCode102018, CatalogSyncEditResponseErrorsCode102019, CatalogSyncEditResponseErrorsCode102020, CatalogSyncEditResponseErrorsCode102021, CatalogSyncEditResponseErrorsCode102022, CatalogSyncEditResponseErrorsCode102023, CatalogSyncEditResponseErrorsCode102024, CatalogSyncEditResponseErrorsCode102025, CatalogSyncEditResponseErrorsCode102026, CatalogSyncEditResponseErrorsCode102027, CatalogSyncEditResponseErrorsCode102028, CatalogSyncEditResponseErrorsCode102029, CatalogSyncEditResponseErrorsCode102030, CatalogSyncEditResponseErrorsCode102031, CatalogSyncEditResponseErrorsCode102032, CatalogSyncEditResponseErrorsCode102033, CatalogSyncEditResponseErrorsCode102034, CatalogSyncEditResponseErrorsCode102035, CatalogSyncEditResponseErrorsCode102036, CatalogSyncEditResponseErrorsCode102037, CatalogSyncEditResponseErrorsCode102038, CatalogSyncEditResponseErrorsCode102039, CatalogSyncEditResponseErrorsCode102040, CatalogSyncEditResponseErrorsCode102041, CatalogSyncEditResponseErrorsCode102042, CatalogSyncEditResponseErrorsCode102043, CatalogSyncEditResponseErrorsCode102044, CatalogSyncEditResponseErrorsCode102045, CatalogSyncEditResponseErrorsCode102046, CatalogSyncEditResponseErrorsCode102047, CatalogSyncEditResponseErrorsCode102048, CatalogSyncEditResponseErrorsCode102049, CatalogSyncEditResponseErrorsCode102050, CatalogSyncEditResponseErrorsCode102051, CatalogSyncEditResponseErrorsCode102052, CatalogSyncEditResponseErrorsCode102053, CatalogSyncEditResponseErrorsCode102054, CatalogSyncEditResponseErrorsCode102055, CatalogSyncEditResponseErrorsCode102056, CatalogSyncEditResponseErrorsCode102057, CatalogSyncEditResponseErrorsCode102058, CatalogSyncEditResponseErrorsCode102059, CatalogSyncEditResponseErrorsCode102060, CatalogSyncEditResponseErrorsCode102061, CatalogSyncEditResponseErrorsCode102062, CatalogSyncEditResponseErrorsCode102063, CatalogSyncEditResponseErrorsCode102064, CatalogSyncEditResponseErrorsCode102065, CatalogSyncEditResponseErrorsCode102066, CatalogSyncEditResponseErrorsCode103001, CatalogSyncEditResponseErrorsCode103002, CatalogSyncEditResponseErrorsCode103003, CatalogSyncEditResponseErrorsCode103004, CatalogSyncEditResponseErrorsCode103005, CatalogSyncEditResponseErrorsCode103006, CatalogSyncEditResponseErrorsCode103007, CatalogSyncEditResponseErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncEditResponseErrorsMeta struct {
	L10nKey       string                                `json:"l10n_key"`
	LoggableError string                                `json:"loggable_error"`
	TemplateData  interface{}                           `json:"template_data"`
	TraceID       string                                `json:"trace_id"`
	JSON          catalogSyncEditResponseErrorsMetaJSON `json:"-"`
}

// catalogSyncEditResponseErrorsMetaJSON contains the JSON metadata for the struct
// [CatalogSyncEditResponseErrorsMeta]
type catalogSyncEditResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncEditResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseErrorsSource struct {
	Parameter           string                                  `json:"parameter"`
	ParameterValueIndex int64                                   `json:"parameter_value_index"`
	Pointer             string                                  `json:"pointer"`
	JSON                catalogSyncEditResponseErrorsSourceJSON `json:"-"`
}

// catalogSyncEditResponseErrorsSourceJSON contains the JSON metadata for the
// struct [CatalogSyncEditResponseErrorsSource]
type catalogSyncEditResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncEditResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponse struct {
	ID                       string                                 `json:"id,required" format:"uuid"`
	Description              string                                 `json:"description,required"`
	DestinationID            string                                 `json:"destination_id,required" format:"uuid"`
	DestinationType          CatalogSyncGetResponseDestinationType  `json:"destination_type,required"`
	LastUserUpdateAt         string                                 `json:"last_user_update_at,required"`
	Name                     string                                 `json:"name,required"`
	Policy                   string                                 `json:"policy,required"`
	UpdateMode               CatalogSyncGetResponseUpdateMode       `json:"update_mode,required"`
	Errors                   map[string]CatalogSyncGetResponseError `json:"errors"`
	IncludesDiscoveriesUntil string                                 `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                                 `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                                 `json:"last_successful_update_at"`
	JSON                     catalogSyncGetResponseJSON             `json:"-"`
}

// catalogSyncGetResponseJSON contains the JSON metadata for the struct
// [CatalogSyncGetResponse]
type catalogSyncGetResponseJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CatalogSyncGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseDestinationType string

const (
	CatalogSyncGetResponseDestinationTypeNone          CatalogSyncGetResponseDestinationType = "NONE"
	CatalogSyncGetResponseDestinationTypeZeroTrustList CatalogSyncGetResponseDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncGetResponseDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncGetResponseDestinationTypeNone, CatalogSyncGetResponseDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncGetResponseUpdateMode string

const (
	CatalogSyncGetResponseUpdateModeAuto   CatalogSyncGetResponseUpdateMode = "AUTO"
	CatalogSyncGetResponseUpdateModeManual CatalogSyncGetResponseUpdateMode = "MANUAL"
)

func (r CatalogSyncGetResponseUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncGetResponseUpdateModeAuto, CatalogSyncGetResponseUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncGetResponseError struct {
	Code             CatalogSyncGetResponseErrorsCode   `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Meta             CatalogSyncGetResponseErrorsMeta   `json:"meta"`
	Source           CatalogSyncGetResponseErrorsSource `json:"source"`
	JSON             catalogSyncGetResponseErrorJSON    `json:"-"`
}

// catalogSyncGetResponseErrorJSON contains the JSON metadata for the struct
// [CatalogSyncGetResponseError]
type catalogSyncGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseErrorsCode int64

const (
	CatalogSyncGetResponseErrorsCode1001   CatalogSyncGetResponseErrorsCode = 1001
	CatalogSyncGetResponseErrorsCode1002   CatalogSyncGetResponseErrorsCode = 1002
	CatalogSyncGetResponseErrorsCode1003   CatalogSyncGetResponseErrorsCode = 1003
	CatalogSyncGetResponseErrorsCode1004   CatalogSyncGetResponseErrorsCode = 1004
	CatalogSyncGetResponseErrorsCode1005   CatalogSyncGetResponseErrorsCode = 1005
	CatalogSyncGetResponseErrorsCode1006   CatalogSyncGetResponseErrorsCode = 1006
	CatalogSyncGetResponseErrorsCode1007   CatalogSyncGetResponseErrorsCode = 1007
	CatalogSyncGetResponseErrorsCode1008   CatalogSyncGetResponseErrorsCode = 1008
	CatalogSyncGetResponseErrorsCode1009   CatalogSyncGetResponseErrorsCode = 1009
	CatalogSyncGetResponseErrorsCode1010   CatalogSyncGetResponseErrorsCode = 1010
	CatalogSyncGetResponseErrorsCode1011   CatalogSyncGetResponseErrorsCode = 1011
	CatalogSyncGetResponseErrorsCode1012   CatalogSyncGetResponseErrorsCode = 1012
	CatalogSyncGetResponseErrorsCode1013   CatalogSyncGetResponseErrorsCode = 1013
	CatalogSyncGetResponseErrorsCode1014   CatalogSyncGetResponseErrorsCode = 1014
	CatalogSyncGetResponseErrorsCode1015   CatalogSyncGetResponseErrorsCode = 1015
	CatalogSyncGetResponseErrorsCode1016   CatalogSyncGetResponseErrorsCode = 1016
	CatalogSyncGetResponseErrorsCode1017   CatalogSyncGetResponseErrorsCode = 1017
	CatalogSyncGetResponseErrorsCode2001   CatalogSyncGetResponseErrorsCode = 2001
	CatalogSyncGetResponseErrorsCode2002   CatalogSyncGetResponseErrorsCode = 2002
	CatalogSyncGetResponseErrorsCode2003   CatalogSyncGetResponseErrorsCode = 2003
	CatalogSyncGetResponseErrorsCode2004   CatalogSyncGetResponseErrorsCode = 2004
	CatalogSyncGetResponseErrorsCode2005   CatalogSyncGetResponseErrorsCode = 2005
	CatalogSyncGetResponseErrorsCode2006   CatalogSyncGetResponseErrorsCode = 2006
	CatalogSyncGetResponseErrorsCode2007   CatalogSyncGetResponseErrorsCode = 2007
	CatalogSyncGetResponseErrorsCode2008   CatalogSyncGetResponseErrorsCode = 2008
	CatalogSyncGetResponseErrorsCode2009   CatalogSyncGetResponseErrorsCode = 2009
	CatalogSyncGetResponseErrorsCode2010   CatalogSyncGetResponseErrorsCode = 2010
	CatalogSyncGetResponseErrorsCode2011   CatalogSyncGetResponseErrorsCode = 2011
	CatalogSyncGetResponseErrorsCode2012   CatalogSyncGetResponseErrorsCode = 2012
	CatalogSyncGetResponseErrorsCode2013   CatalogSyncGetResponseErrorsCode = 2013
	CatalogSyncGetResponseErrorsCode2014   CatalogSyncGetResponseErrorsCode = 2014
	CatalogSyncGetResponseErrorsCode2015   CatalogSyncGetResponseErrorsCode = 2015
	CatalogSyncGetResponseErrorsCode2016   CatalogSyncGetResponseErrorsCode = 2016
	CatalogSyncGetResponseErrorsCode2017   CatalogSyncGetResponseErrorsCode = 2017
	CatalogSyncGetResponseErrorsCode2018   CatalogSyncGetResponseErrorsCode = 2018
	CatalogSyncGetResponseErrorsCode2019   CatalogSyncGetResponseErrorsCode = 2019
	CatalogSyncGetResponseErrorsCode2020   CatalogSyncGetResponseErrorsCode = 2020
	CatalogSyncGetResponseErrorsCode2021   CatalogSyncGetResponseErrorsCode = 2021
	CatalogSyncGetResponseErrorsCode2022   CatalogSyncGetResponseErrorsCode = 2022
	CatalogSyncGetResponseErrorsCode3001   CatalogSyncGetResponseErrorsCode = 3001
	CatalogSyncGetResponseErrorsCode3002   CatalogSyncGetResponseErrorsCode = 3002
	CatalogSyncGetResponseErrorsCode3003   CatalogSyncGetResponseErrorsCode = 3003
	CatalogSyncGetResponseErrorsCode3004   CatalogSyncGetResponseErrorsCode = 3004
	CatalogSyncGetResponseErrorsCode3005   CatalogSyncGetResponseErrorsCode = 3005
	CatalogSyncGetResponseErrorsCode3006   CatalogSyncGetResponseErrorsCode = 3006
	CatalogSyncGetResponseErrorsCode3007   CatalogSyncGetResponseErrorsCode = 3007
	CatalogSyncGetResponseErrorsCode4001   CatalogSyncGetResponseErrorsCode = 4001
	CatalogSyncGetResponseErrorsCode4002   CatalogSyncGetResponseErrorsCode = 4002
	CatalogSyncGetResponseErrorsCode4003   CatalogSyncGetResponseErrorsCode = 4003
	CatalogSyncGetResponseErrorsCode4004   CatalogSyncGetResponseErrorsCode = 4004
	CatalogSyncGetResponseErrorsCode4005   CatalogSyncGetResponseErrorsCode = 4005
	CatalogSyncGetResponseErrorsCode4006   CatalogSyncGetResponseErrorsCode = 4006
	CatalogSyncGetResponseErrorsCode4007   CatalogSyncGetResponseErrorsCode = 4007
	CatalogSyncGetResponseErrorsCode4008   CatalogSyncGetResponseErrorsCode = 4008
	CatalogSyncGetResponseErrorsCode4009   CatalogSyncGetResponseErrorsCode = 4009
	CatalogSyncGetResponseErrorsCode4010   CatalogSyncGetResponseErrorsCode = 4010
	CatalogSyncGetResponseErrorsCode4011   CatalogSyncGetResponseErrorsCode = 4011
	CatalogSyncGetResponseErrorsCode4012   CatalogSyncGetResponseErrorsCode = 4012
	CatalogSyncGetResponseErrorsCode4013   CatalogSyncGetResponseErrorsCode = 4013
	CatalogSyncGetResponseErrorsCode4014   CatalogSyncGetResponseErrorsCode = 4014
	CatalogSyncGetResponseErrorsCode4015   CatalogSyncGetResponseErrorsCode = 4015
	CatalogSyncGetResponseErrorsCode4016   CatalogSyncGetResponseErrorsCode = 4016
	CatalogSyncGetResponseErrorsCode4017   CatalogSyncGetResponseErrorsCode = 4017
	CatalogSyncGetResponseErrorsCode4018   CatalogSyncGetResponseErrorsCode = 4018
	CatalogSyncGetResponseErrorsCode4019   CatalogSyncGetResponseErrorsCode = 4019
	CatalogSyncGetResponseErrorsCode4020   CatalogSyncGetResponseErrorsCode = 4020
	CatalogSyncGetResponseErrorsCode4021   CatalogSyncGetResponseErrorsCode = 4021
	CatalogSyncGetResponseErrorsCode4022   CatalogSyncGetResponseErrorsCode = 4022
	CatalogSyncGetResponseErrorsCode4023   CatalogSyncGetResponseErrorsCode = 4023
	CatalogSyncGetResponseErrorsCode5001   CatalogSyncGetResponseErrorsCode = 5001
	CatalogSyncGetResponseErrorsCode5002   CatalogSyncGetResponseErrorsCode = 5002
	CatalogSyncGetResponseErrorsCode5003   CatalogSyncGetResponseErrorsCode = 5003
	CatalogSyncGetResponseErrorsCode5004   CatalogSyncGetResponseErrorsCode = 5004
	CatalogSyncGetResponseErrorsCode102000 CatalogSyncGetResponseErrorsCode = 102000
	CatalogSyncGetResponseErrorsCode102001 CatalogSyncGetResponseErrorsCode = 102001
	CatalogSyncGetResponseErrorsCode102002 CatalogSyncGetResponseErrorsCode = 102002
	CatalogSyncGetResponseErrorsCode102003 CatalogSyncGetResponseErrorsCode = 102003
	CatalogSyncGetResponseErrorsCode102004 CatalogSyncGetResponseErrorsCode = 102004
	CatalogSyncGetResponseErrorsCode102005 CatalogSyncGetResponseErrorsCode = 102005
	CatalogSyncGetResponseErrorsCode102006 CatalogSyncGetResponseErrorsCode = 102006
	CatalogSyncGetResponseErrorsCode102007 CatalogSyncGetResponseErrorsCode = 102007
	CatalogSyncGetResponseErrorsCode102008 CatalogSyncGetResponseErrorsCode = 102008
	CatalogSyncGetResponseErrorsCode102009 CatalogSyncGetResponseErrorsCode = 102009
	CatalogSyncGetResponseErrorsCode102010 CatalogSyncGetResponseErrorsCode = 102010
	CatalogSyncGetResponseErrorsCode102011 CatalogSyncGetResponseErrorsCode = 102011
	CatalogSyncGetResponseErrorsCode102012 CatalogSyncGetResponseErrorsCode = 102012
	CatalogSyncGetResponseErrorsCode102013 CatalogSyncGetResponseErrorsCode = 102013
	CatalogSyncGetResponseErrorsCode102014 CatalogSyncGetResponseErrorsCode = 102014
	CatalogSyncGetResponseErrorsCode102015 CatalogSyncGetResponseErrorsCode = 102015
	CatalogSyncGetResponseErrorsCode102016 CatalogSyncGetResponseErrorsCode = 102016
	CatalogSyncGetResponseErrorsCode102017 CatalogSyncGetResponseErrorsCode = 102017
	CatalogSyncGetResponseErrorsCode102018 CatalogSyncGetResponseErrorsCode = 102018
	CatalogSyncGetResponseErrorsCode102019 CatalogSyncGetResponseErrorsCode = 102019
	CatalogSyncGetResponseErrorsCode102020 CatalogSyncGetResponseErrorsCode = 102020
	CatalogSyncGetResponseErrorsCode102021 CatalogSyncGetResponseErrorsCode = 102021
	CatalogSyncGetResponseErrorsCode102022 CatalogSyncGetResponseErrorsCode = 102022
	CatalogSyncGetResponseErrorsCode102023 CatalogSyncGetResponseErrorsCode = 102023
	CatalogSyncGetResponseErrorsCode102024 CatalogSyncGetResponseErrorsCode = 102024
	CatalogSyncGetResponseErrorsCode102025 CatalogSyncGetResponseErrorsCode = 102025
	CatalogSyncGetResponseErrorsCode102026 CatalogSyncGetResponseErrorsCode = 102026
	CatalogSyncGetResponseErrorsCode102027 CatalogSyncGetResponseErrorsCode = 102027
	CatalogSyncGetResponseErrorsCode102028 CatalogSyncGetResponseErrorsCode = 102028
	CatalogSyncGetResponseErrorsCode102029 CatalogSyncGetResponseErrorsCode = 102029
	CatalogSyncGetResponseErrorsCode102030 CatalogSyncGetResponseErrorsCode = 102030
	CatalogSyncGetResponseErrorsCode102031 CatalogSyncGetResponseErrorsCode = 102031
	CatalogSyncGetResponseErrorsCode102032 CatalogSyncGetResponseErrorsCode = 102032
	CatalogSyncGetResponseErrorsCode102033 CatalogSyncGetResponseErrorsCode = 102033
	CatalogSyncGetResponseErrorsCode102034 CatalogSyncGetResponseErrorsCode = 102034
	CatalogSyncGetResponseErrorsCode102035 CatalogSyncGetResponseErrorsCode = 102035
	CatalogSyncGetResponseErrorsCode102036 CatalogSyncGetResponseErrorsCode = 102036
	CatalogSyncGetResponseErrorsCode102037 CatalogSyncGetResponseErrorsCode = 102037
	CatalogSyncGetResponseErrorsCode102038 CatalogSyncGetResponseErrorsCode = 102038
	CatalogSyncGetResponseErrorsCode102039 CatalogSyncGetResponseErrorsCode = 102039
	CatalogSyncGetResponseErrorsCode102040 CatalogSyncGetResponseErrorsCode = 102040
	CatalogSyncGetResponseErrorsCode102041 CatalogSyncGetResponseErrorsCode = 102041
	CatalogSyncGetResponseErrorsCode102042 CatalogSyncGetResponseErrorsCode = 102042
	CatalogSyncGetResponseErrorsCode102043 CatalogSyncGetResponseErrorsCode = 102043
	CatalogSyncGetResponseErrorsCode102044 CatalogSyncGetResponseErrorsCode = 102044
	CatalogSyncGetResponseErrorsCode102045 CatalogSyncGetResponseErrorsCode = 102045
	CatalogSyncGetResponseErrorsCode102046 CatalogSyncGetResponseErrorsCode = 102046
	CatalogSyncGetResponseErrorsCode102047 CatalogSyncGetResponseErrorsCode = 102047
	CatalogSyncGetResponseErrorsCode102048 CatalogSyncGetResponseErrorsCode = 102048
	CatalogSyncGetResponseErrorsCode102049 CatalogSyncGetResponseErrorsCode = 102049
	CatalogSyncGetResponseErrorsCode102050 CatalogSyncGetResponseErrorsCode = 102050
	CatalogSyncGetResponseErrorsCode102051 CatalogSyncGetResponseErrorsCode = 102051
	CatalogSyncGetResponseErrorsCode102052 CatalogSyncGetResponseErrorsCode = 102052
	CatalogSyncGetResponseErrorsCode102053 CatalogSyncGetResponseErrorsCode = 102053
	CatalogSyncGetResponseErrorsCode102054 CatalogSyncGetResponseErrorsCode = 102054
	CatalogSyncGetResponseErrorsCode102055 CatalogSyncGetResponseErrorsCode = 102055
	CatalogSyncGetResponseErrorsCode102056 CatalogSyncGetResponseErrorsCode = 102056
	CatalogSyncGetResponseErrorsCode102057 CatalogSyncGetResponseErrorsCode = 102057
	CatalogSyncGetResponseErrorsCode102058 CatalogSyncGetResponseErrorsCode = 102058
	CatalogSyncGetResponseErrorsCode102059 CatalogSyncGetResponseErrorsCode = 102059
	CatalogSyncGetResponseErrorsCode102060 CatalogSyncGetResponseErrorsCode = 102060
	CatalogSyncGetResponseErrorsCode102061 CatalogSyncGetResponseErrorsCode = 102061
	CatalogSyncGetResponseErrorsCode102062 CatalogSyncGetResponseErrorsCode = 102062
	CatalogSyncGetResponseErrorsCode102063 CatalogSyncGetResponseErrorsCode = 102063
	CatalogSyncGetResponseErrorsCode102064 CatalogSyncGetResponseErrorsCode = 102064
	CatalogSyncGetResponseErrorsCode102065 CatalogSyncGetResponseErrorsCode = 102065
	CatalogSyncGetResponseErrorsCode102066 CatalogSyncGetResponseErrorsCode = 102066
	CatalogSyncGetResponseErrorsCode103001 CatalogSyncGetResponseErrorsCode = 103001
	CatalogSyncGetResponseErrorsCode103002 CatalogSyncGetResponseErrorsCode = 103002
	CatalogSyncGetResponseErrorsCode103003 CatalogSyncGetResponseErrorsCode = 103003
	CatalogSyncGetResponseErrorsCode103004 CatalogSyncGetResponseErrorsCode = 103004
	CatalogSyncGetResponseErrorsCode103005 CatalogSyncGetResponseErrorsCode = 103005
	CatalogSyncGetResponseErrorsCode103006 CatalogSyncGetResponseErrorsCode = 103006
	CatalogSyncGetResponseErrorsCode103007 CatalogSyncGetResponseErrorsCode = 103007
	CatalogSyncGetResponseErrorsCode103008 CatalogSyncGetResponseErrorsCode = 103008
)

func (r CatalogSyncGetResponseErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncGetResponseErrorsCode1001, CatalogSyncGetResponseErrorsCode1002, CatalogSyncGetResponseErrorsCode1003, CatalogSyncGetResponseErrorsCode1004, CatalogSyncGetResponseErrorsCode1005, CatalogSyncGetResponseErrorsCode1006, CatalogSyncGetResponseErrorsCode1007, CatalogSyncGetResponseErrorsCode1008, CatalogSyncGetResponseErrorsCode1009, CatalogSyncGetResponseErrorsCode1010, CatalogSyncGetResponseErrorsCode1011, CatalogSyncGetResponseErrorsCode1012, CatalogSyncGetResponseErrorsCode1013, CatalogSyncGetResponseErrorsCode1014, CatalogSyncGetResponseErrorsCode1015, CatalogSyncGetResponseErrorsCode1016, CatalogSyncGetResponseErrorsCode1017, CatalogSyncGetResponseErrorsCode2001, CatalogSyncGetResponseErrorsCode2002, CatalogSyncGetResponseErrorsCode2003, CatalogSyncGetResponseErrorsCode2004, CatalogSyncGetResponseErrorsCode2005, CatalogSyncGetResponseErrorsCode2006, CatalogSyncGetResponseErrorsCode2007, CatalogSyncGetResponseErrorsCode2008, CatalogSyncGetResponseErrorsCode2009, CatalogSyncGetResponseErrorsCode2010, CatalogSyncGetResponseErrorsCode2011, CatalogSyncGetResponseErrorsCode2012, CatalogSyncGetResponseErrorsCode2013, CatalogSyncGetResponseErrorsCode2014, CatalogSyncGetResponseErrorsCode2015, CatalogSyncGetResponseErrorsCode2016, CatalogSyncGetResponseErrorsCode2017, CatalogSyncGetResponseErrorsCode2018, CatalogSyncGetResponseErrorsCode2019, CatalogSyncGetResponseErrorsCode2020, CatalogSyncGetResponseErrorsCode2021, CatalogSyncGetResponseErrorsCode2022, CatalogSyncGetResponseErrorsCode3001, CatalogSyncGetResponseErrorsCode3002, CatalogSyncGetResponseErrorsCode3003, CatalogSyncGetResponseErrorsCode3004, CatalogSyncGetResponseErrorsCode3005, CatalogSyncGetResponseErrorsCode3006, CatalogSyncGetResponseErrorsCode3007, CatalogSyncGetResponseErrorsCode4001, CatalogSyncGetResponseErrorsCode4002, CatalogSyncGetResponseErrorsCode4003, CatalogSyncGetResponseErrorsCode4004, CatalogSyncGetResponseErrorsCode4005, CatalogSyncGetResponseErrorsCode4006, CatalogSyncGetResponseErrorsCode4007, CatalogSyncGetResponseErrorsCode4008, CatalogSyncGetResponseErrorsCode4009, CatalogSyncGetResponseErrorsCode4010, CatalogSyncGetResponseErrorsCode4011, CatalogSyncGetResponseErrorsCode4012, CatalogSyncGetResponseErrorsCode4013, CatalogSyncGetResponseErrorsCode4014, CatalogSyncGetResponseErrorsCode4015, CatalogSyncGetResponseErrorsCode4016, CatalogSyncGetResponseErrorsCode4017, CatalogSyncGetResponseErrorsCode4018, CatalogSyncGetResponseErrorsCode4019, CatalogSyncGetResponseErrorsCode4020, CatalogSyncGetResponseErrorsCode4021, CatalogSyncGetResponseErrorsCode4022, CatalogSyncGetResponseErrorsCode4023, CatalogSyncGetResponseErrorsCode5001, CatalogSyncGetResponseErrorsCode5002, CatalogSyncGetResponseErrorsCode5003, CatalogSyncGetResponseErrorsCode5004, CatalogSyncGetResponseErrorsCode102000, CatalogSyncGetResponseErrorsCode102001, CatalogSyncGetResponseErrorsCode102002, CatalogSyncGetResponseErrorsCode102003, CatalogSyncGetResponseErrorsCode102004, CatalogSyncGetResponseErrorsCode102005, CatalogSyncGetResponseErrorsCode102006, CatalogSyncGetResponseErrorsCode102007, CatalogSyncGetResponseErrorsCode102008, CatalogSyncGetResponseErrorsCode102009, CatalogSyncGetResponseErrorsCode102010, CatalogSyncGetResponseErrorsCode102011, CatalogSyncGetResponseErrorsCode102012, CatalogSyncGetResponseErrorsCode102013, CatalogSyncGetResponseErrorsCode102014, CatalogSyncGetResponseErrorsCode102015, CatalogSyncGetResponseErrorsCode102016, CatalogSyncGetResponseErrorsCode102017, CatalogSyncGetResponseErrorsCode102018, CatalogSyncGetResponseErrorsCode102019, CatalogSyncGetResponseErrorsCode102020, CatalogSyncGetResponseErrorsCode102021, CatalogSyncGetResponseErrorsCode102022, CatalogSyncGetResponseErrorsCode102023, CatalogSyncGetResponseErrorsCode102024, CatalogSyncGetResponseErrorsCode102025, CatalogSyncGetResponseErrorsCode102026, CatalogSyncGetResponseErrorsCode102027, CatalogSyncGetResponseErrorsCode102028, CatalogSyncGetResponseErrorsCode102029, CatalogSyncGetResponseErrorsCode102030, CatalogSyncGetResponseErrorsCode102031, CatalogSyncGetResponseErrorsCode102032, CatalogSyncGetResponseErrorsCode102033, CatalogSyncGetResponseErrorsCode102034, CatalogSyncGetResponseErrorsCode102035, CatalogSyncGetResponseErrorsCode102036, CatalogSyncGetResponseErrorsCode102037, CatalogSyncGetResponseErrorsCode102038, CatalogSyncGetResponseErrorsCode102039, CatalogSyncGetResponseErrorsCode102040, CatalogSyncGetResponseErrorsCode102041, CatalogSyncGetResponseErrorsCode102042, CatalogSyncGetResponseErrorsCode102043, CatalogSyncGetResponseErrorsCode102044, CatalogSyncGetResponseErrorsCode102045, CatalogSyncGetResponseErrorsCode102046, CatalogSyncGetResponseErrorsCode102047, CatalogSyncGetResponseErrorsCode102048, CatalogSyncGetResponseErrorsCode102049, CatalogSyncGetResponseErrorsCode102050, CatalogSyncGetResponseErrorsCode102051, CatalogSyncGetResponseErrorsCode102052, CatalogSyncGetResponseErrorsCode102053, CatalogSyncGetResponseErrorsCode102054, CatalogSyncGetResponseErrorsCode102055, CatalogSyncGetResponseErrorsCode102056, CatalogSyncGetResponseErrorsCode102057, CatalogSyncGetResponseErrorsCode102058, CatalogSyncGetResponseErrorsCode102059, CatalogSyncGetResponseErrorsCode102060, CatalogSyncGetResponseErrorsCode102061, CatalogSyncGetResponseErrorsCode102062, CatalogSyncGetResponseErrorsCode102063, CatalogSyncGetResponseErrorsCode102064, CatalogSyncGetResponseErrorsCode102065, CatalogSyncGetResponseErrorsCode102066, CatalogSyncGetResponseErrorsCode103001, CatalogSyncGetResponseErrorsCode103002, CatalogSyncGetResponseErrorsCode103003, CatalogSyncGetResponseErrorsCode103004, CatalogSyncGetResponseErrorsCode103005, CatalogSyncGetResponseErrorsCode103006, CatalogSyncGetResponseErrorsCode103007, CatalogSyncGetResponseErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncGetResponseErrorsMeta struct {
	L10nKey       string                               `json:"l10n_key"`
	LoggableError string                               `json:"loggable_error"`
	TemplateData  interface{}                          `json:"template_data"`
	TraceID       string                               `json:"trace_id"`
	JSON          catalogSyncGetResponseErrorsMetaJSON `json:"-"`
}

// catalogSyncGetResponseErrorsMetaJSON contains the JSON metadata for the struct
// [CatalogSyncGetResponseErrorsMeta]
type catalogSyncGetResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncGetResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseErrorsSource struct {
	Parameter           string                                 `json:"parameter"`
	ParameterValueIndex int64                                  `json:"parameter_value_index"`
	Pointer             string                                 `json:"pointer"`
	JSON                catalogSyncGetResponseErrorsSourceJSON `json:"-"`
}

// catalogSyncGetResponseErrorsSourceJSON contains the JSON metadata for the struct
// [CatalogSyncGetResponseErrorsSource]
type catalogSyncGetResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewParams struct {
	AccountID       param.Field[string]                              `path:"account_id,required"`
	DestinationType param.Field[CatalogSyncNewParamsDestinationType] `json:"destination_type,required"`
	Name            param.Field[string]                              `json:"name,required"`
	UpdateMode      param.Field[CatalogSyncNewParamsUpdateMode]      `json:"update_mode,required"`
	Description     param.Field[string]                              `json:"description"`
	Policy          param.Field[string]                              `json:"policy"`
	Forwarded       param.Field[string]                              `header:"forwarded"`
}

func (r CatalogSyncNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CatalogSyncNewParamsDestinationType string

const (
	CatalogSyncNewParamsDestinationTypeNone          CatalogSyncNewParamsDestinationType = "NONE"
	CatalogSyncNewParamsDestinationTypeZeroTrustList CatalogSyncNewParamsDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncNewParamsDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncNewParamsDestinationTypeNone, CatalogSyncNewParamsDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncNewParamsUpdateMode string

const (
	CatalogSyncNewParamsUpdateModeAuto   CatalogSyncNewParamsUpdateMode = "AUTO"
	CatalogSyncNewParamsUpdateModeManual CatalogSyncNewParamsUpdateMode = "MANUAL"
)

func (r CatalogSyncNewParamsUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncNewParamsUpdateModeAuto, CatalogSyncNewParamsUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncNewResponseEnvelope struct {
	Errors   []CatalogSyncNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CatalogSyncNewResponse                   `json:"result,required"`
	Success  bool                                     `json:"success,required"`
	JSON     catalogSyncNewResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncNewResponseEnvelope]
type catalogSyncNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeErrors struct {
	Code             CatalogSyncNewResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Meta             CatalogSyncNewResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncNewResponseEnvelopeErrors]
type catalogSyncNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeErrorsCode int64

const (
	CatalogSyncNewResponseEnvelopeErrorsCode1001   CatalogSyncNewResponseEnvelopeErrorsCode = 1001
	CatalogSyncNewResponseEnvelopeErrorsCode1002   CatalogSyncNewResponseEnvelopeErrorsCode = 1002
	CatalogSyncNewResponseEnvelopeErrorsCode1003   CatalogSyncNewResponseEnvelopeErrorsCode = 1003
	CatalogSyncNewResponseEnvelopeErrorsCode1004   CatalogSyncNewResponseEnvelopeErrorsCode = 1004
	CatalogSyncNewResponseEnvelopeErrorsCode1005   CatalogSyncNewResponseEnvelopeErrorsCode = 1005
	CatalogSyncNewResponseEnvelopeErrorsCode1006   CatalogSyncNewResponseEnvelopeErrorsCode = 1006
	CatalogSyncNewResponseEnvelopeErrorsCode1007   CatalogSyncNewResponseEnvelopeErrorsCode = 1007
	CatalogSyncNewResponseEnvelopeErrorsCode1008   CatalogSyncNewResponseEnvelopeErrorsCode = 1008
	CatalogSyncNewResponseEnvelopeErrorsCode1009   CatalogSyncNewResponseEnvelopeErrorsCode = 1009
	CatalogSyncNewResponseEnvelopeErrorsCode1010   CatalogSyncNewResponseEnvelopeErrorsCode = 1010
	CatalogSyncNewResponseEnvelopeErrorsCode1011   CatalogSyncNewResponseEnvelopeErrorsCode = 1011
	CatalogSyncNewResponseEnvelopeErrorsCode1012   CatalogSyncNewResponseEnvelopeErrorsCode = 1012
	CatalogSyncNewResponseEnvelopeErrorsCode1013   CatalogSyncNewResponseEnvelopeErrorsCode = 1013
	CatalogSyncNewResponseEnvelopeErrorsCode1014   CatalogSyncNewResponseEnvelopeErrorsCode = 1014
	CatalogSyncNewResponseEnvelopeErrorsCode1015   CatalogSyncNewResponseEnvelopeErrorsCode = 1015
	CatalogSyncNewResponseEnvelopeErrorsCode1016   CatalogSyncNewResponseEnvelopeErrorsCode = 1016
	CatalogSyncNewResponseEnvelopeErrorsCode1017   CatalogSyncNewResponseEnvelopeErrorsCode = 1017
	CatalogSyncNewResponseEnvelopeErrorsCode2001   CatalogSyncNewResponseEnvelopeErrorsCode = 2001
	CatalogSyncNewResponseEnvelopeErrorsCode2002   CatalogSyncNewResponseEnvelopeErrorsCode = 2002
	CatalogSyncNewResponseEnvelopeErrorsCode2003   CatalogSyncNewResponseEnvelopeErrorsCode = 2003
	CatalogSyncNewResponseEnvelopeErrorsCode2004   CatalogSyncNewResponseEnvelopeErrorsCode = 2004
	CatalogSyncNewResponseEnvelopeErrorsCode2005   CatalogSyncNewResponseEnvelopeErrorsCode = 2005
	CatalogSyncNewResponseEnvelopeErrorsCode2006   CatalogSyncNewResponseEnvelopeErrorsCode = 2006
	CatalogSyncNewResponseEnvelopeErrorsCode2007   CatalogSyncNewResponseEnvelopeErrorsCode = 2007
	CatalogSyncNewResponseEnvelopeErrorsCode2008   CatalogSyncNewResponseEnvelopeErrorsCode = 2008
	CatalogSyncNewResponseEnvelopeErrorsCode2009   CatalogSyncNewResponseEnvelopeErrorsCode = 2009
	CatalogSyncNewResponseEnvelopeErrorsCode2010   CatalogSyncNewResponseEnvelopeErrorsCode = 2010
	CatalogSyncNewResponseEnvelopeErrorsCode2011   CatalogSyncNewResponseEnvelopeErrorsCode = 2011
	CatalogSyncNewResponseEnvelopeErrorsCode2012   CatalogSyncNewResponseEnvelopeErrorsCode = 2012
	CatalogSyncNewResponseEnvelopeErrorsCode2013   CatalogSyncNewResponseEnvelopeErrorsCode = 2013
	CatalogSyncNewResponseEnvelopeErrorsCode2014   CatalogSyncNewResponseEnvelopeErrorsCode = 2014
	CatalogSyncNewResponseEnvelopeErrorsCode2015   CatalogSyncNewResponseEnvelopeErrorsCode = 2015
	CatalogSyncNewResponseEnvelopeErrorsCode2016   CatalogSyncNewResponseEnvelopeErrorsCode = 2016
	CatalogSyncNewResponseEnvelopeErrorsCode2017   CatalogSyncNewResponseEnvelopeErrorsCode = 2017
	CatalogSyncNewResponseEnvelopeErrorsCode2018   CatalogSyncNewResponseEnvelopeErrorsCode = 2018
	CatalogSyncNewResponseEnvelopeErrorsCode2019   CatalogSyncNewResponseEnvelopeErrorsCode = 2019
	CatalogSyncNewResponseEnvelopeErrorsCode2020   CatalogSyncNewResponseEnvelopeErrorsCode = 2020
	CatalogSyncNewResponseEnvelopeErrorsCode2021   CatalogSyncNewResponseEnvelopeErrorsCode = 2021
	CatalogSyncNewResponseEnvelopeErrorsCode2022   CatalogSyncNewResponseEnvelopeErrorsCode = 2022
	CatalogSyncNewResponseEnvelopeErrorsCode3001   CatalogSyncNewResponseEnvelopeErrorsCode = 3001
	CatalogSyncNewResponseEnvelopeErrorsCode3002   CatalogSyncNewResponseEnvelopeErrorsCode = 3002
	CatalogSyncNewResponseEnvelopeErrorsCode3003   CatalogSyncNewResponseEnvelopeErrorsCode = 3003
	CatalogSyncNewResponseEnvelopeErrorsCode3004   CatalogSyncNewResponseEnvelopeErrorsCode = 3004
	CatalogSyncNewResponseEnvelopeErrorsCode3005   CatalogSyncNewResponseEnvelopeErrorsCode = 3005
	CatalogSyncNewResponseEnvelopeErrorsCode3006   CatalogSyncNewResponseEnvelopeErrorsCode = 3006
	CatalogSyncNewResponseEnvelopeErrorsCode3007   CatalogSyncNewResponseEnvelopeErrorsCode = 3007
	CatalogSyncNewResponseEnvelopeErrorsCode4001   CatalogSyncNewResponseEnvelopeErrorsCode = 4001
	CatalogSyncNewResponseEnvelopeErrorsCode4002   CatalogSyncNewResponseEnvelopeErrorsCode = 4002
	CatalogSyncNewResponseEnvelopeErrorsCode4003   CatalogSyncNewResponseEnvelopeErrorsCode = 4003
	CatalogSyncNewResponseEnvelopeErrorsCode4004   CatalogSyncNewResponseEnvelopeErrorsCode = 4004
	CatalogSyncNewResponseEnvelopeErrorsCode4005   CatalogSyncNewResponseEnvelopeErrorsCode = 4005
	CatalogSyncNewResponseEnvelopeErrorsCode4006   CatalogSyncNewResponseEnvelopeErrorsCode = 4006
	CatalogSyncNewResponseEnvelopeErrorsCode4007   CatalogSyncNewResponseEnvelopeErrorsCode = 4007
	CatalogSyncNewResponseEnvelopeErrorsCode4008   CatalogSyncNewResponseEnvelopeErrorsCode = 4008
	CatalogSyncNewResponseEnvelopeErrorsCode4009   CatalogSyncNewResponseEnvelopeErrorsCode = 4009
	CatalogSyncNewResponseEnvelopeErrorsCode4010   CatalogSyncNewResponseEnvelopeErrorsCode = 4010
	CatalogSyncNewResponseEnvelopeErrorsCode4011   CatalogSyncNewResponseEnvelopeErrorsCode = 4011
	CatalogSyncNewResponseEnvelopeErrorsCode4012   CatalogSyncNewResponseEnvelopeErrorsCode = 4012
	CatalogSyncNewResponseEnvelopeErrorsCode4013   CatalogSyncNewResponseEnvelopeErrorsCode = 4013
	CatalogSyncNewResponseEnvelopeErrorsCode4014   CatalogSyncNewResponseEnvelopeErrorsCode = 4014
	CatalogSyncNewResponseEnvelopeErrorsCode4015   CatalogSyncNewResponseEnvelopeErrorsCode = 4015
	CatalogSyncNewResponseEnvelopeErrorsCode4016   CatalogSyncNewResponseEnvelopeErrorsCode = 4016
	CatalogSyncNewResponseEnvelopeErrorsCode4017   CatalogSyncNewResponseEnvelopeErrorsCode = 4017
	CatalogSyncNewResponseEnvelopeErrorsCode4018   CatalogSyncNewResponseEnvelopeErrorsCode = 4018
	CatalogSyncNewResponseEnvelopeErrorsCode4019   CatalogSyncNewResponseEnvelopeErrorsCode = 4019
	CatalogSyncNewResponseEnvelopeErrorsCode4020   CatalogSyncNewResponseEnvelopeErrorsCode = 4020
	CatalogSyncNewResponseEnvelopeErrorsCode4021   CatalogSyncNewResponseEnvelopeErrorsCode = 4021
	CatalogSyncNewResponseEnvelopeErrorsCode4022   CatalogSyncNewResponseEnvelopeErrorsCode = 4022
	CatalogSyncNewResponseEnvelopeErrorsCode4023   CatalogSyncNewResponseEnvelopeErrorsCode = 4023
	CatalogSyncNewResponseEnvelopeErrorsCode5001   CatalogSyncNewResponseEnvelopeErrorsCode = 5001
	CatalogSyncNewResponseEnvelopeErrorsCode5002   CatalogSyncNewResponseEnvelopeErrorsCode = 5002
	CatalogSyncNewResponseEnvelopeErrorsCode5003   CatalogSyncNewResponseEnvelopeErrorsCode = 5003
	CatalogSyncNewResponseEnvelopeErrorsCode5004   CatalogSyncNewResponseEnvelopeErrorsCode = 5004
	CatalogSyncNewResponseEnvelopeErrorsCode102000 CatalogSyncNewResponseEnvelopeErrorsCode = 102000
	CatalogSyncNewResponseEnvelopeErrorsCode102001 CatalogSyncNewResponseEnvelopeErrorsCode = 102001
	CatalogSyncNewResponseEnvelopeErrorsCode102002 CatalogSyncNewResponseEnvelopeErrorsCode = 102002
	CatalogSyncNewResponseEnvelopeErrorsCode102003 CatalogSyncNewResponseEnvelopeErrorsCode = 102003
	CatalogSyncNewResponseEnvelopeErrorsCode102004 CatalogSyncNewResponseEnvelopeErrorsCode = 102004
	CatalogSyncNewResponseEnvelopeErrorsCode102005 CatalogSyncNewResponseEnvelopeErrorsCode = 102005
	CatalogSyncNewResponseEnvelopeErrorsCode102006 CatalogSyncNewResponseEnvelopeErrorsCode = 102006
	CatalogSyncNewResponseEnvelopeErrorsCode102007 CatalogSyncNewResponseEnvelopeErrorsCode = 102007
	CatalogSyncNewResponseEnvelopeErrorsCode102008 CatalogSyncNewResponseEnvelopeErrorsCode = 102008
	CatalogSyncNewResponseEnvelopeErrorsCode102009 CatalogSyncNewResponseEnvelopeErrorsCode = 102009
	CatalogSyncNewResponseEnvelopeErrorsCode102010 CatalogSyncNewResponseEnvelopeErrorsCode = 102010
	CatalogSyncNewResponseEnvelopeErrorsCode102011 CatalogSyncNewResponseEnvelopeErrorsCode = 102011
	CatalogSyncNewResponseEnvelopeErrorsCode102012 CatalogSyncNewResponseEnvelopeErrorsCode = 102012
	CatalogSyncNewResponseEnvelopeErrorsCode102013 CatalogSyncNewResponseEnvelopeErrorsCode = 102013
	CatalogSyncNewResponseEnvelopeErrorsCode102014 CatalogSyncNewResponseEnvelopeErrorsCode = 102014
	CatalogSyncNewResponseEnvelopeErrorsCode102015 CatalogSyncNewResponseEnvelopeErrorsCode = 102015
	CatalogSyncNewResponseEnvelopeErrorsCode102016 CatalogSyncNewResponseEnvelopeErrorsCode = 102016
	CatalogSyncNewResponseEnvelopeErrorsCode102017 CatalogSyncNewResponseEnvelopeErrorsCode = 102017
	CatalogSyncNewResponseEnvelopeErrorsCode102018 CatalogSyncNewResponseEnvelopeErrorsCode = 102018
	CatalogSyncNewResponseEnvelopeErrorsCode102019 CatalogSyncNewResponseEnvelopeErrorsCode = 102019
	CatalogSyncNewResponseEnvelopeErrorsCode102020 CatalogSyncNewResponseEnvelopeErrorsCode = 102020
	CatalogSyncNewResponseEnvelopeErrorsCode102021 CatalogSyncNewResponseEnvelopeErrorsCode = 102021
	CatalogSyncNewResponseEnvelopeErrorsCode102022 CatalogSyncNewResponseEnvelopeErrorsCode = 102022
	CatalogSyncNewResponseEnvelopeErrorsCode102023 CatalogSyncNewResponseEnvelopeErrorsCode = 102023
	CatalogSyncNewResponseEnvelopeErrorsCode102024 CatalogSyncNewResponseEnvelopeErrorsCode = 102024
	CatalogSyncNewResponseEnvelopeErrorsCode102025 CatalogSyncNewResponseEnvelopeErrorsCode = 102025
	CatalogSyncNewResponseEnvelopeErrorsCode102026 CatalogSyncNewResponseEnvelopeErrorsCode = 102026
	CatalogSyncNewResponseEnvelopeErrorsCode102027 CatalogSyncNewResponseEnvelopeErrorsCode = 102027
	CatalogSyncNewResponseEnvelopeErrorsCode102028 CatalogSyncNewResponseEnvelopeErrorsCode = 102028
	CatalogSyncNewResponseEnvelopeErrorsCode102029 CatalogSyncNewResponseEnvelopeErrorsCode = 102029
	CatalogSyncNewResponseEnvelopeErrorsCode102030 CatalogSyncNewResponseEnvelopeErrorsCode = 102030
	CatalogSyncNewResponseEnvelopeErrorsCode102031 CatalogSyncNewResponseEnvelopeErrorsCode = 102031
	CatalogSyncNewResponseEnvelopeErrorsCode102032 CatalogSyncNewResponseEnvelopeErrorsCode = 102032
	CatalogSyncNewResponseEnvelopeErrorsCode102033 CatalogSyncNewResponseEnvelopeErrorsCode = 102033
	CatalogSyncNewResponseEnvelopeErrorsCode102034 CatalogSyncNewResponseEnvelopeErrorsCode = 102034
	CatalogSyncNewResponseEnvelopeErrorsCode102035 CatalogSyncNewResponseEnvelopeErrorsCode = 102035
	CatalogSyncNewResponseEnvelopeErrorsCode102036 CatalogSyncNewResponseEnvelopeErrorsCode = 102036
	CatalogSyncNewResponseEnvelopeErrorsCode102037 CatalogSyncNewResponseEnvelopeErrorsCode = 102037
	CatalogSyncNewResponseEnvelopeErrorsCode102038 CatalogSyncNewResponseEnvelopeErrorsCode = 102038
	CatalogSyncNewResponseEnvelopeErrorsCode102039 CatalogSyncNewResponseEnvelopeErrorsCode = 102039
	CatalogSyncNewResponseEnvelopeErrorsCode102040 CatalogSyncNewResponseEnvelopeErrorsCode = 102040
	CatalogSyncNewResponseEnvelopeErrorsCode102041 CatalogSyncNewResponseEnvelopeErrorsCode = 102041
	CatalogSyncNewResponseEnvelopeErrorsCode102042 CatalogSyncNewResponseEnvelopeErrorsCode = 102042
	CatalogSyncNewResponseEnvelopeErrorsCode102043 CatalogSyncNewResponseEnvelopeErrorsCode = 102043
	CatalogSyncNewResponseEnvelopeErrorsCode102044 CatalogSyncNewResponseEnvelopeErrorsCode = 102044
	CatalogSyncNewResponseEnvelopeErrorsCode102045 CatalogSyncNewResponseEnvelopeErrorsCode = 102045
	CatalogSyncNewResponseEnvelopeErrorsCode102046 CatalogSyncNewResponseEnvelopeErrorsCode = 102046
	CatalogSyncNewResponseEnvelopeErrorsCode102047 CatalogSyncNewResponseEnvelopeErrorsCode = 102047
	CatalogSyncNewResponseEnvelopeErrorsCode102048 CatalogSyncNewResponseEnvelopeErrorsCode = 102048
	CatalogSyncNewResponseEnvelopeErrorsCode102049 CatalogSyncNewResponseEnvelopeErrorsCode = 102049
	CatalogSyncNewResponseEnvelopeErrorsCode102050 CatalogSyncNewResponseEnvelopeErrorsCode = 102050
	CatalogSyncNewResponseEnvelopeErrorsCode102051 CatalogSyncNewResponseEnvelopeErrorsCode = 102051
	CatalogSyncNewResponseEnvelopeErrorsCode102052 CatalogSyncNewResponseEnvelopeErrorsCode = 102052
	CatalogSyncNewResponseEnvelopeErrorsCode102053 CatalogSyncNewResponseEnvelopeErrorsCode = 102053
	CatalogSyncNewResponseEnvelopeErrorsCode102054 CatalogSyncNewResponseEnvelopeErrorsCode = 102054
	CatalogSyncNewResponseEnvelopeErrorsCode102055 CatalogSyncNewResponseEnvelopeErrorsCode = 102055
	CatalogSyncNewResponseEnvelopeErrorsCode102056 CatalogSyncNewResponseEnvelopeErrorsCode = 102056
	CatalogSyncNewResponseEnvelopeErrorsCode102057 CatalogSyncNewResponseEnvelopeErrorsCode = 102057
	CatalogSyncNewResponseEnvelopeErrorsCode102058 CatalogSyncNewResponseEnvelopeErrorsCode = 102058
	CatalogSyncNewResponseEnvelopeErrorsCode102059 CatalogSyncNewResponseEnvelopeErrorsCode = 102059
	CatalogSyncNewResponseEnvelopeErrorsCode102060 CatalogSyncNewResponseEnvelopeErrorsCode = 102060
	CatalogSyncNewResponseEnvelopeErrorsCode102061 CatalogSyncNewResponseEnvelopeErrorsCode = 102061
	CatalogSyncNewResponseEnvelopeErrorsCode102062 CatalogSyncNewResponseEnvelopeErrorsCode = 102062
	CatalogSyncNewResponseEnvelopeErrorsCode102063 CatalogSyncNewResponseEnvelopeErrorsCode = 102063
	CatalogSyncNewResponseEnvelopeErrorsCode102064 CatalogSyncNewResponseEnvelopeErrorsCode = 102064
	CatalogSyncNewResponseEnvelopeErrorsCode102065 CatalogSyncNewResponseEnvelopeErrorsCode = 102065
	CatalogSyncNewResponseEnvelopeErrorsCode102066 CatalogSyncNewResponseEnvelopeErrorsCode = 102066
	CatalogSyncNewResponseEnvelopeErrorsCode103001 CatalogSyncNewResponseEnvelopeErrorsCode = 103001
	CatalogSyncNewResponseEnvelopeErrorsCode103002 CatalogSyncNewResponseEnvelopeErrorsCode = 103002
	CatalogSyncNewResponseEnvelopeErrorsCode103003 CatalogSyncNewResponseEnvelopeErrorsCode = 103003
	CatalogSyncNewResponseEnvelopeErrorsCode103004 CatalogSyncNewResponseEnvelopeErrorsCode = 103004
	CatalogSyncNewResponseEnvelopeErrorsCode103005 CatalogSyncNewResponseEnvelopeErrorsCode = 103005
	CatalogSyncNewResponseEnvelopeErrorsCode103006 CatalogSyncNewResponseEnvelopeErrorsCode = 103006
	CatalogSyncNewResponseEnvelopeErrorsCode103007 CatalogSyncNewResponseEnvelopeErrorsCode = 103007
	CatalogSyncNewResponseEnvelopeErrorsCode103008 CatalogSyncNewResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncNewResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncNewResponseEnvelopeErrorsCode1001, CatalogSyncNewResponseEnvelopeErrorsCode1002, CatalogSyncNewResponseEnvelopeErrorsCode1003, CatalogSyncNewResponseEnvelopeErrorsCode1004, CatalogSyncNewResponseEnvelopeErrorsCode1005, CatalogSyncNewResponseEnvelopeErrorsCode1006, CatalogSyncNewResponseEnvelopeErrorsCode1007, CatalogSyncNewResponseEnvelopeErrorsCode1008, CatalogSyncNewResponseEnvelopeErrorsCode1009, CatalogSyncNewResponseEnvelopeErrorsCode1010, CatalogSyncNewResponseEnvelopeErrorsCode1011, CatalogSyncNewResponseEnvelopeErrorsCode1012, CatalogSyncNewResponseEnvelopeErrorsCode1013, CatalogSyncNewResponseEnvelopeErrorsCode1014, CatalogSyncNewResponseEnvelopeErrorsCode1015, CatalogSyncNewResponseEnvelopeErrorsCode1016, CatalogSyncNewResponseEnvelopeErrorsCode1017, CatalogSyncNewResponseEnvelopeErrorsCode2001, CatalogSyncNewResponseEnvelopeErrorsCode2002, CatalogSyncNewResponseEnvelopeErrorsCode2003, CatalogSyncNewResponseEnvelopeErrorsCode2004, CatalogSyncNewResponseEnvelopeErrorsCode2005, CatalogSyncNewResponseEnvelopeErrorsCode2006, CatalogSyncNewResponseEnvelopeErrorsCode2007, CatalogSyncNewResponseEnvelopeErrorsCode2008, CatalogSyncNewResponseEnvelopeErrorsCode2009, CatalogSyncNewResponseEnvelopeErrorsCode2010, CatalogSyncNewResponseEnvelopeErrorsCode2011, CatalogSyncNewResponseEnvelopeErrorsCode2012, CatalogSyncNewResponseEnvelopeErrorsCode2013, CatalogSyncNewResponseEnvelopeErrorsCode2014, CatalogSyncNewResponseEnvelopeErrorsCode2015, CatalogSyncNewResponseEnvelopeErrorsCode2016, CatalogSyncNewResponseEnvelopeErrorsCode2017, CatalogSyncNewResponseEnvelopeErrorsCode2018, CatalogSyncNewResponseEnvelopeErrorsCode2019, CatalogSyncNewResponseEnvelopeErrorsCode2020, CatalogSyncNewResponseEnvelopeErrorsCode2021, CatalogSyncNewResponseEnvelopeErrorsCode2022, CatalogSyncNewResponseEnvelopeErrorsCode3001, CatalogSyncNewResponseEnvelopeErrorsCode3002, CatalogSyncNewResponseEnvelopeErrorsCode3003, CatalogSyncNewResponseEnvelopeErrorsCode3004, CatalogSyncNewResponseEnvelopeErrorsCode3005, CatalogSyncNewResponseEnvelopeErrorsCode3006, CatalogSyncNewResponseEnvelopeErrorsCode3007, CatalogSyncNewResponseEnvelopeErrorsCode4001, CatalogSyncNewResponseEnvelopeErrorsCode4002, CatalogSyncNewResponseEnvelopeErrorsCode4003, CatalogSyncNewResponseEnvelopeErrorsCode4004, CatalogSyncNewResponseEnvelopeErrorsCode4005, CatalogSyncNewResponseEnvelopeErrorsCode4006, CatalogSyncNewResponseEnvelopeErrorsCode4007, CatalogSyncNewResponseEnvelopeErrorsCode4008, CatalogSyncNewResponseEnvelopeErrorsCode4009, CatalogSyncNewResponseEnvelopeErrorsCode4010, CatalogSyncNewResponseEnvelopeErrorsCode4011, CatalogSyncNewResponseEnvelopeErrorsCode4012, CatalogSyncNewResponseEnvelopeErrorsCode4013, CatalogSyncNewResponseEnvelopeErrorsCode4014, CatalogSyncNewResponseEnvelopeErrorsCode4015, CatalogSyncNewResponseEnvelopeErrorsCode4016, CatalogSyncNewResponseEnvelopeErrorsCode4017, CatalogSyncNewResponseEnvelopeErrorsCode4018, CatalogSyncNewResponseEnvelopeErrorsCode4019, CatalogSyncNewResponseEnvelopeErrorsCode4020, CatalogSyncNewResponseEnvelopeErrorsCode4021, CatalogSyncNewResponseEnvelopeErrorsCode4022, CatalogSyncNewResponseEnvelopeErrorsCode4023, CatalogSyncNewResponseEnvelopeErrorsCode5001, CatalogSyncNewResponseEnvelopeErrorsCode5002, CatalogSyncNewResponseEnvelopeErrorsCode5003, CatalogSyncNewResponseEnvelopeErrorsCode5004, CatalogSyncNewResponseEnvelopeErrorsCode102000, CatalogSyncNewResponseEnvelopeErrorsCode102001, CatalogSyncNewResponseEnvelopeErrorsCode102002, CatalogSyncNewResponseEnvelopeErrorsCode102003, CatalogSyncNewResponseEnvelopeErrorsCode102004, CatalogSyncNewResponseEnvelopeErrorsCode102005, CatalogSyncNewResponseEnvelopeErrorsCode102006, CatalogSyncNewResponseEnvelopeErrorsCode102007, CatalogSyncNewResponseEnvelopeErrorsCode102008, CatalogSyncNewResponseEnvelopeErrorsCode102009, CatalogSyncNewResponseEnvelopeErrorsCode102010, CatalogSyncNewResponseEnvelopeErrorsCode102011, CatalogSyncNewResponseEnvelopeErrorsCode102012, CatalogSyncNewResponseEnvelopeErrorsCode102013, CatalogSyncNewResponseEnvelopeErrorsCode102014, CatalogSyncNewResponseEnvelopeErrorsCode102015, CatalogSyncNewResponseEnvelopeErrorsCode102016, CatalogSyncNewResponseEnvelopeErrorsCode102017, CatalogSyncNewResponseEnvelopeErrorsCode102018, CatalogSyncNewResponseEnvelopeErrorsCode102019, CatalogSyncNewResponseEnvelopeErrorsCode102020, CatalogSyncNewResponseEnvelopeErrorsCode102021, CatalogSyncNewResponseEnvelopeErrorsCode102022, CatalogSyncNewResponseEnvelopeErrorsCode102023, CatalogSyncNewResponseEnvelopeErrorsCode102024, CatalogSyncNewResponseEnvelopeErrorsCode102025, CatalogSyncNewResponseEnvelopeErrorsCode102026, CatalogSyncNewResponseEnvelopeErrorsCode102027, CatalogSyncNewResponseEnvelopeErrorsCode102028, CatalogSyncNewResponseEnvelopeErrorsCode102029, CatalogSyncNewResponseEnvelopeErrorsCode102030, CatalogSyncNewResponseEnvelopeErrorsCode102031, CatalogSyncNewResponseEnvelopeErrorsCode102032, CatalogSyncNewResponseEnvelopeErrorsCode102033, CatalogSyncNewResponseEnvelopeErrorsCode102034, CatalogSyncNewResponseEnvelopeErrorsCode102035, CatalogSyncNewResponseEnvelopeErrorsCode102036, CatalogSyncNewResponseEnvelopeErrorsCode102037, CatalogSyncNewResponseEnvelopeErrorsCode102038, CatalogSyncNewResponseEnvelopeErrorsCode102039, CatalogSyncNewResponseEnvelopeErrorsCode102040, CatalogSyncNewResponseEnvelopeErrorsCode102041, CatalogSyncNewResponseEnvelopeErrorsCode102042, CatalogSyncNewResponseEnvelopeErrorsCode102043, CatalogSyncNewResponseEnvelopeErrorsCode102044, CatalogSyncNewResponseEnvelopeErrorsCode102045, CatalogSyncNewResponseEnvelopeErrorsCode102046, CatalogSyncNewResponseEnvelopeErrorsCode102047, CatalogSyncNewResponseEnvelopeErrorsCode102048, CatalogSyncNewResponseEnvelopeErrorsCode102049, CatalogSyncNewResponseEnvelopeErrorsCode102050, CatalogSyncNewResponseEnvelopeErrorsCode102051, CatalogSyncNewResponseEnvelopeErrorsCode102052, CatalogSyncNewResponseEnvelopeErrorsCode102053, CatalogSyncNewResponseEnvelopeErrorsCode102054, CatalogSyncNewResponseEnvelopeErrorsCode102055, CatalogSyncNewResponseEnvelopeErrorsCode102056, CatalogSyncNewResponseEnvelopeErrorsCode102057, CatalogSyncNewResponseEnvelopeErrorsCode102058, CatalogSyncNewResponseEnvelopeErrorsCode102059, CatalogSyncNewResponseEnvelopeErrorsCode102060, CatalogSyncNewResponseEnvelopeErrorsCode102061, CatalogSyncNewResponseEnvelopeErrorsCode102062, CatalogSyncNewResponseEnvelopeErrorsCode102063, CatalogSyncNewResponseEnvelopeErrorsCode102064, CatalogSyncNewResponseEnvelopeErrorsCode102065, CatalogSyncNewResponseEnvelopeErrorsCode102066, CatalogSyncNewResponseEnvelopeErrorsCode103001, CatalogSyncNewResponseEnvelopeErrorsCode103002, CatalogSyncNewResponseEnvelopeErrorsCode103003, CatalogSyncNewResponseEnvelopeErrorsCode103004, CatalogSyncNewResponseEnvelopeErrorsCode103005, CatalogSyncNewResponseEnvelopeErrorsCode103006, CatalogSyncNewResponseEnvelopeErrorsCode103007, CatalogSyncNewResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncNewResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                       `json:"l10n_key"`
	LoggableError string                                       `json:"loggable_error"`
	TemplateData  interface{}                                  `json:"template_data"`
	TraceID       string                                       `json:"trace_id"`
	JSON          catalogSyncNewResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncNewResponseEnvelopeErrorsMetaJSON contains the JSON metadata for the
// struct [CatalogSyncNewResponseEnvelopeErrorsMeta]
type catalogSyncNewResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeErrorsSource struct {
	Parameter           string                                         `json:"parameter"`
	ParameterValueIndex int64                                          `json:"parameter_value_index"`
	Pointer             string                                         `json:"pointer"`
	JSON                catalogSyncNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CatalogSyncNewResponseEnvelopeErrorsSource]
type catalogSyncNewResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeMessages struct {
	Code             CatalogSyncNewResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Meta             CatalogSyncNewResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CatalogSyncNewResponseEnvelopeMessages]
type catalogSyncNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeMessagesCode int64

const (
	CatalogSyncNewResponseEnvelopeMessagesCode1001   CatalogSyncNewResponseEnvelopeMessagesCode = 1001
	CatalogSyncNewResponseEnvelopeMessagesCode1002   CatalogSyncNewResponseEnvelopeMessagesCode = 1002
	CatalogSyncNewResponseEnvelopeMessagesCode1003   CatalogSyncNewResponseEnvelopeMessagesCode = 1003
	CatalogSyncNewResponseEnvelopeMessagesCode1004   CatalogSyncNewResponseEnvelopeMessagesCode = 1004
	CatalogSyncNewResponseEnvelopeMessagesCode1005   CatalogSyncNewResponseEnvelopeMessagesCode = 1005
	CatalogSyncNewResponseEnvelopeMessagesCode1006   CatalogSyncNewResponseEnvelopeMessagesCode = 1006
	CatalogSyncNewResponseEnvelopeMessagesCode1007   CatalogSyncNewResponseEnvelopeMessagesCode = 1007
	CatalogSyncNewResponseEnvelopeMessagesCode1008   CatalogSyncNewResponseEnvelopeMessagesCode = 1008
	CatalogSyncNewResponseEnvelopeMessagesCode1009   CatalogSyncNewResponseEnvelopeMessagesCode = 1009
	CatalogSyncNewResponseEnvelopeMessagesCode1010   CatalogSyncNewResponseEnvelopeMessagesCode = 1010
	CatalogSyncNewResponseEnvelopeMessagesCode1011   CatalogSyncNewResponseEnvelopeMessagesCode = 1011
	CatalogSyncNewResponseEnvelopeMessagesCode1012   CatalogSyncNewResponseEnvelopeMessagesCode = 1012
	CatalogSyncNewResponseEnvelopeMessagesCode1013   CatalogSyncNewResponseEnvelopeMessagesCode = 1013
	CatalogSyncNewResponseEnvelopeMessagesCode1014   CatalogSyncNewResponseEnvelopeMessagesCode = 1014
	CatalogSyncNewResponseEnvelopeMessagesCode1015   CatalogSyncNewResponseEnvelopeMessagesCode = 1015
	CatalogSyncNewResponseEnvelopeMessagesCode1016   CatalogSyncNewResponseEnvelopeMessagesCode = 1016
	CatalogSyncNewResponseEnvelopeMessagesCode1017   CatalogSyncNewResponseEnvelopeMessagesCode = 1017
	CatalogSyncNewResponseEnvelopeMessagesCode2001   CatalogSyncNewResponseEnvelopeMessagesCode = 2001
	CatalogSyncNewResponseEnvelopeMessagesCode2002   CatalogSyncNewResponseEnvelopeMessagesCode = 2002
	CatalogSyncNewResponseEnvelopeMessagesCode2003   CatalogSyncNewResponseEnvelopeMessagesCode = 2003
	CatalogSyncNewResponseEnvelopeMessagesCode2004   CatalogSyncNewResponseEnvelopeMessagesCode = 2004
	CatalogSyncNewResponseEnvelopeMessagesCode2005   CatalogSyncNewResponseEnvelopeMessagesCode = 2005
	CatalogSyncNewResponseEnvelopeMessagesCode2006   CatalogSyncNewResponseEnvelopeMessagesCode = 2006
	CatalogSyncNewResponseEnvelopeMessagesCode2007   CatalogSyncNewResponseEnvelopeMessagesCode = 2007
	CatalogSyncNewResponseEnvelopeMessagesCode2008   CatalogSyncNewResponseEnvelopeMessagesCode = 2008
	CatalogSyncNewResponseEnvelopeMessagesCode2009   CatalogSyncNewResponseEnvelopeMessagesCode = 2009
	CatalogSyncNewResponseEnvelopeMessagesCode2010   CatalogSyncNewResponseEnvelopeMessagesCode = 2010
	CatalogSyncNewResponseEnvelopeMessagesCode2011   CatalogSyncNewResponseEnvelopeMessagesCode = 2011
	CatalogSyncNewResponseEnvelopeMessagesCode2012   CatalogSyncNewResponseEnvelopeMessagesCode = 2012
	CatalogSyncNewResponseEnvelopeMessagesCode2013   CatalogSyncNewResponseEnvelopeMessagesCode = 2013
	CatalogSyncNewResponseEnvelopeMessagesCode2014   CatalogSyncNewResponseEnvelopeMessagesCode = 2014
	CatalogSyncNewResponseEnvelopeMessagesCode2015   CatalogSyncNewResponseEnvelopeMessagesCode = 2015
	CatalogSyncNewResponseEnvelopeMessagesCode2016   CatalogSyncNewResponseEnvelopeMessagesCode = 2016
	CatalogSyncNewResponseEnvelopeMessagesCode2017   CatalogSyncNewResponseEnvelopeMessagesCode = 2017
	CatalogSyncNewResponseEnvelopeMessagesCode2018   CatalogSyncNewResponseEnvelopeMessagesCode = 2018
	CatalogSyncNewResponseEnvelopeMessagesCode2019   CatalogSyncNewResponseEnvelopeMessagesCode = 2019
	CatalogSyncNewResponseEnvelopeMessagesCode2020   CatalogSyncNewResponseEnvelopeMessagesCode = 2020
	CatalogSyncNewResponseEnvelopeMessagesCode2021   CatalogSyncNewResponseEnvelopeMessagesCode = 2021
	CatalogSyncNewResponseEnvelopeMessagesCode2022   CatalogSyncNewResponseEnvelopeMessagesCode = 2022
	CatalogSyncNewResponseEnvelopeMessagesCode3001   CatalogSyncNewResponseEnvelopeMessagesCode = 3001
	CatalogSyncNewResponseEnvelopeMessagesCode3002   CatalogSyncNewResponseEnvelopeMessagesCode = 3002
	CatalogSyncNewResponseEnvelopeMessagesCode3003   CatalogSyncNewResponseEnvelopeMessagesCode = 3003
	CatalogSyncNewResponseEnvelopeMessagesCode3004   CatalogSyncNewResponseEnvelopeMessagesCode = 3004
	CatalogSyncNewResponseEnvelopeMessagesCode3005   CatalogSyncNewResponseEnvelopeMessagesCode = 3005
	CatalogSyncNewResponseEnvelopeMessagesCode3006   CatalogSyncNewResponseEnvelopeMessagesCode = 3006
	CatalogSyncNewResponseEnvelopeMessagesCode3007   CatalogSyncNewResponseEnvelopeMessagesCode = 3007
	CatalogSyncNewResponseEnvelopeMessagesCode4001   CatalogSyncNewResponseEnvelopeMessagesCode = 4001
	CatalogSyncNewResponseEnvelopeMessagesCode4002   CatalogSyncNewResponseEnvelopeMessagesCode = 4002
	CatalogSyncNewResponseEnvelopeMessagesCode4003   CatalogSyncNewResponseEnvelopeMessagesCode = 4003
	CatalogSyncNewResponseEnvelopeMessagesCode4004   CatalogSyncNewResponseEnvelopeMessagesCode = 4004
	CatalogSyncNewResponseEnvelopeMessagesCode4005   CatalogSyncNewResponseEnvelopeMessagesCode = 4005
	CatalogSyncNewResponseEnvelopeMessagesCode4006   CatalogSyncNewResponseEnvelopeMessagesCode = 4006
	CatalogSyncNewResponseEnvelopeMessagesCode4007   CatalogSyncNewResponseEnvelopeMessagesCode = 4007
	CatalogSyncNewResponseEnvelopeMessagesCode4008   CatalogSyncNewResponseEnvelopeMessagesCode = 4008
	CatalogSyncNewResponseEnvelopeMessagesCode4009   CatalogSyncNewResponseEnvelopeMessagesCode = 4009
	CatalogSyncNewResponseEnvelopeMessagesCode4010   CatalogSyncNewResponseEnvelopeMessagesCode = 4010
	CatalogSyncNewResponseEnvelopeMessagesCode4011   CatalogSyncNewResponseEnvelopeMessagesCode = 4011
	CatalogSyncNewResponseEnvelopeMessagesCode4012   CatalogSyncNewResponseEnvelopeMessagesCode = 4012
	CatalogSyncNewResponseEnvelopeMessagesCode4013   CatalogSyncNewResponseEnvelopeMessagesCode = 4013
	CatalogSyncNewResponseEnvelopeMessagesCode4014   CatalogSyncNewResponseEnvelopeMessagesCode = 4014
	CatalogSyncNewResponseEnvelopeMessagesCode4015   CatalogSyncNewResponseEnvelopeMessagesCode = 4015
	CatalogSyncNewResponseEnvelopeMessagesCode4016   CatalogSyncNewResponseEnvelopeMessagesCode = 4016
	CatalogSyncNewResponseEnvelopeMessagesCode4017   CatalogSyncNewResponseEnvelopeMessagesCode = 4017
	CatalogSyncNewResponseEnvelopeMessagesCode4018   CatalogSyncNewResponseEnvelopeMessagesCode = 4018
	CatalogSyncNewResponseEnvelopeMessagesCode4019   CatalogSyncNewResponseEnvelopeMessagesCode = 4019
	CatalogSyncNewResponseEnvelopeMessagesCode4020   CatalogSyncNewResponseEnvelopeMessagesCode = 4020
	CatalogSyncNewResponseEnvelopeMessagesCode4021   CatalogSyncNewResponseEnvelopeMessagesCode = 4021
	CatalogSyncNewResponseEnvelopeMessagesCode4022   CatalogSyncNewResponseEnvelopeMessagesCode = 4022
	CatalogSyncNewResponseEnvelopeMessagesCode4023   CatalogSyncNewResponseEnvelopeMessagesCode = 4023
	CatalogSyncNewResponseEnvelopeMessagesCode5001   CatalogSyncNewResponseEnvelopeMessagesCode = 5001
	CatalogSyncNewResponseEnvelopeMessagesCode5002   CatalogSyncNewResponseEnvelopeMessagesCode = 5002
	CatalogSyncNewResponseEnvelopeMessagesCode5003   CatalogSyncNewResponseEnvelopeMessagesCode = 5003
	CatalogSyncNewResponseEnvelopeMessagesCode5004   CatalogSyncNewResponseEnvelopeMessagesCode = 5004
	CatalogSyncNewResponseEnvelopeMessagesCode102000 CatalogSyncNewResponseEnvelopeMessagesCode = 102000
	CatalogSyncNewResponseEnvelopeMessagesCode102001 CatalogSyncNewResponseEnvelopeMessagesCode = 102001
	CatalogSyncNewResponseEnvelopeMessagesCode102002 CatalogSyncNewResponseEnvelopeMessagesCode = 102002
	CatalogSyncNewResponseEnvelopeMessagesCode102003 CatalogSyncNewResponseEnvelopeMessagesCode = 102003
	CatalogSyncNewResponseEnvelopeMessagesCode102004 CatalogSyncNewResponseEnvelopeMessagesCode = 102004
	CatalogSyncNewResponseEnvelopeMessagesCode102005 CatalogSyncNewResponseEnvelopeMessagesCode = 102005
	CatalogSyncNewResponseEnvelopeMessagesCode102006 CatalogSyncNewResponseEnvelopeMessagesCode = 102006
	CatalogSyncNewResponseEnvelopeMessagesCode102007 CatalogSyncNewResponseEnvelopeMessagesCode = 102007
	CatalogSyncNewResponseEnvelopeMessagesCode102008 CatalogSyncNewResponseEnvelopeMessagesCode = 102008
	CatalogSyncNewResponseEnvelopeMessagesCode102009 CatalogSyncNewResponseEnvelopeMessagesCode = 102009
	CatalogSyncNewResponseEnvelopeMessagesCode102010 CatalogSyncNewResponseEnvelopeMessagesCode = 102010
	CatalogSyncNewResponseEnvelopeMessagesCode102011 CatalogSyncNewResponseEnvelopeMessagesCode = 102011
	CatalogSyncNewResponseEnvelopeMessagesCode102012 CatalogSyncNewResponseEnvelopeMessagesCode = 102012
	CatalogSyncNewResponseEnvelopeMessagesCode102013 CatalogSyncNewResponseEnvelopeMessagesCode = 102013
	CatalogSyncNewResponseEnvelopeMessagesCode102014 CatalogSyncNewResponseEnvelopeMessagesCode = 102014
	CatalogSyncNewResponseEnvelopeMessagesCode102015 CatalogSyncNewResponseEnvelopeMessagesCode = 102015
	CatalogSyncNewResponseEnvelopeMessagesCode102016 CatalogSyncNewResponseEnvelopeMessagesCode = 102016
	CatalogSyncNewResponseEnvelopeMessagesCode102017 CatalogSyncNewResponseEnvelopeMessagesCode = 102017
	CatalogSyncNewResponseEnvelopeMessagesCode102018 CatalogSyncNewResponseEnvelopeMessagesCode = 102018
	CatalogSyncNewResponseEnvelopeMessagesCode102019 CatalogSyncNewResponseEnvelopeMessagesCode = 102019
	CatalogSyncNewResponseEnvelopeMessagesCode102020 CatalogSyncNewResponseEnvelopeMessagesCode = 102020
	CatalogSyncNewResponseEnvelopeMessagesCode102021 CatalogSyncNewResponseEnvelopeMessagesCode = 102021
	CatalogSyncNewResponseEnvelopeMessagesCode102022 CatalogSyncNewResponseEnvelopeMessagesCode = 102022
	CatalogSyncNewResponseEnvelopeMessagesCode102023 CatalogSyncNewResponseEnvelopeMessagesCode = 102023
	CatalogSyncNewResponseEnvelopeMessagesCode102024 CatalogSyncNewResponseEnvelopeMessagesCode = 102024
	CatalogSyncNewResponseEnvelopeMessagesCode102025 CatalogSyncNewResponseEnvelopeMessagesCode = 102025
	CatalogSyncNewResponseEnvelopeMessagesCode102026 CatalogSyncNewResponseEnvelopeMessagesCode = 102026
	CatalogSyncNewResponseEnvelopeMessagesCode102027 CatalogSyncNewResponseEnvelopeMessagesCode = 102027
	CatalogSyncNewResponseEnvelopeMessagesCode102028 CatalogSyncNewResponseEnvelopeMessagesCode = 102028
	CatalogSyncNewResponseEnvelopeMessagesCode102029 CatalogSyncNewResponseEnvelopeMessagesCode = 102029
	CatalogSyncNewResponseEnvelopeMessagesCode102030 CatalogSyncNewResponseEnvelopeMessagesCode = 102030
	CatalogSyncNewResponseEnvelopeMessagesCode102031 CatalogSyncNewResponseEnvelopeMessagesCode = 102031
	CatalogSyncNewResponseEnvelopeMessagesCode102032 CatalogSyncNewResponseEnvelopeMessagesCode = 102032
	CatalogSyncNewResponseEnvelopeMessagesCode102033 CatalogSyncNewResponseEnvelopeMessagesCode = 102033
	CatalogSyncNewResponseEnvelopeMessagesCode102034 CatalogSyncNewResponseEnvelopeMessagesCode = 102034
	CatalogSyncNewResponseEnvelopeMessagesCode102035 CatalogSyncNewResponseEnvelopeMessagesCode = 102035
	CatalogSyncNewResponseEnvelopeMessagesCode102036 CatalogSyncNewResponseEnvelopeMessagesCode = 102036
	CatalogSyncNewResponseEnvelopeMessagesCode102037 CatalogSyncNewResponseEnvelopeMessagesCode = 102037
	CatalogSyncNewResponseEnvelopeMessagesCode102038 CatalogSyncNewResponseEnvelopeMessagesCode = 102038
	CatalogSyncNewResponseEnvelopeMessagesCode102039 CatalogSyncNewResponseEnvelopeMessagesCode = 102039
	CatalogSyncNewResponseEnvelopeMessagesCode102040 CatalogSyncNewResponseEnvelopeMessagesCode = 102040
	CatalogSyncNewResponseEnvelopeMessagesCode102041 CatalogSyncNewResponseEnvelopeMessagesCode = 102041
	CatalogSyncNewResponseEnvelopeMessagesCode102042 CatalogSyncNewResponseEnvelopeMessagesCode = 102042
	CatalogSyncNewResponseEnvelopeMessagesCode102043 CatalogSyncNewResponseEnvelopeMessagesCode = 102043
	CatalogSyncNewResponseEnvelopeMessagesCode102044 CatalogSyncNewResponseEnvelopeMessagesCode = 102044
	CatalogSyncNewResponseEnvelopeMessagesCode102045 CatalogSyncNewResponseEnvelopeMessagesCode = 102045
	CatalogSyncNewResponseEnvelopeMessagesCode102046 CatalogSyncNewResponseEnvelopeMessagesCode = 102046
	CatalogSyncNewResponseEnvelopeMessagesCode102047 CatalogSyncNewResponseEnvelopeMessagesCode = 102047
	CatalogSyncNewResponseEnvelopeMessagesCode102048 CatalogSyncNewResponseEnvelopeMessagesCode = 102048
	CatalogSyncNewResponseEnvelopeMessagesCode102049 CatalogSyncNewResponseEnvelopeMessagesCode = 102049
	CatalogSyncNewResponseEnvelopeMessagesCode102050 CatalogSyncNewResponseEnvelopeMessagesCode = 102050
	CatalogSyncNewResponseEnvelopeMessagesCode102051 CatalogSyncNewResponseEnvelopeMessagesCode = 102051
	CatalogSyncNewResponseEnvelopeMessagesCode102052 CatalogSyncNewResponseEnvelopeMessagesCode = 102052
	CatalogSyncNewResponseEnvelopeMessagesCode102053 CatalogSyncNewResponseEnvelopeMessagesCode = 102053
	CatalogSyncNewResponseEnvelopeMessagesCode102054 CatalogSyncNewResponseEnvelopeMessagesCode = 102054
	CatalogSyncNewResponseEnvelopeMessagesCode102055 CatalogSyncNewResponseEnvelopeMessagesCode = 102055
	CatalogSyncNewResponseEnvelopeMessagesCode102056 CatalogSyncNewResponseEnvelopeMessagesCode = 102056
	CatalogSyncNewResponseEnvelopeMessagesCode102057 CatalogSyncNewResponseEnvelopeMessagesCode = 102057
	CatalogSyncNewResponseEnvelopeMessagesCode102058 CatalogSyncNewResponseEnvelopeMessagesCode = 102058
	CatalogSyncNewResponseEnvelopeMessagesCode102059 CatalogSyncNewResponseEnvelopeMessagesCode = 102059
	CatalogSyncNewResponseEnvelopeMessagesCode102060 CatalogSyncNewResponseEnvelopeMessagesCode = 102060
	CatalogSyncNewResponseEnvelopeMessagesCode102061 CatalogSyncNewResponseEnvelopeMessagesCode = 102061
	CatalogSyncNewResponseEnvelopeMessagesCode102062 CatalogSyncNewResponseEnvelopeMessagesCode = 102062
	CatalogSyncNewResponseEnvelopeMessagesCode102063 CatalogSyncNewResponseEnvelopeMessagesCode = 102063
	CatalogSyncNewResponseEnvelopeMessagesCode102064 CatalogSyncNewResponseEnvelopeMessagesCode = 102064
	CatalogSyncNewResponseEnvelopeMessagesCode102065 CatalogSyncNewResponseEnvelopeMessagesCode = 102065
	CatalogSyncNewResponseEnvelopeMessagesCode102066 CatalogSyncNewResponseEnvelopeMessagesCode = 102066
	CatalogSyncNewResponseEnvelopeMessagesCode103001 CatalogSyncNewResponseEnvelopeMessagesCode = 103001
	CatalogSyncNewResponseEnvelopeMessagesCode103002 CatalogSyncNewResponseEnvelopeMessagesCode = 103002
	CatalogSyncNewResponseEnvelopeMessagesCode103003 CatalogSyncNewResponseEnvelopeMessagesCode = 103003
	CatalogSyncNewResponseEnvelopeMessagesCode103004 CatalogSyncNewResponseEnvelopeMessagesCode = 103004
	CatalogSyncNewResponseEnvelopeMessagesCode103005 CatalogSyncNewResponseEnvelopeMessagesCode = 103005
	CatalogSyncNewResponseEnvelopeMessagesCode103006 CatalogSyncNewResponseEnvelopeMessagesCode = 103006
	CatalogSyncNewResponseEnvelopeMessagesCode103007 CatalogSyncNewResponseEnvelopeMessagesCode = 103007
	CatalogSyncNewResponseEnvelopeMessagesCode103008 CatalogSyncNewResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncNewResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncNewResponseEnvelopeMessagesCode1001, CatalogSyncNewResponseEnvelopeMessagesCode1002, CatalogSyncNewResponseEnvelopeMessagesCode1003, CatalogSyncNewResponseEnvelopeMessagesCode1004, CatalogSyncNewResponseEnvelopeMessagesCode1005, CatalogSyncNewResponseEnvelopeMessagesCode1006, CatalogSyncNewResponseEnvelopeMessagesCode1007, CatalogSyncNewResponseEnvelopeMessagesCode1008, CatalogSyncNewResponseEnvelopeMessagesCode1009, CatalogSyncNewResponseEnvelopeMessagesCode1010, CatalogSyncNewResponseEnvelopeMessagesCode1011, CatalogSyncNewResponseEnvelopeMessagesCode1012, CatalogSyncNewResponseEnvelopeMessagesCode1013, CatalogSyncNewResponseEnvelopeMessagesCode1014, CatalogSyncNewResponseEnvelopeMessagesCode1015, CatalogSyncNewResponseEnvelopeMessagesCode1016, CatalogSyncNewResponseEnvelopeMessagesCode1017, CatalogSyncNewResponseEnvelopeMessagesCode2001, CatalogSyncNewResponseEnvelopeMessagesCode2002, CatalogSyncNewResponseEnvelopeMessagesCode2003, CatalogSyncNewResponseEnvelopeMessagesCode2004, CatalogSyncNewResponseEnvelopeMessagesCode2005, CatalogSyncNewResponseEnvelopeMessagesCode2006, CatalogSyncNewResponseEnvelopeMessagesCode2007, CatalogSyncNewResponseEnvelopeMessagesCode2008, CatalogSyncNewResponseEnvelopeMessagesCode2009, CatalogSyncNewResponseEnvelopeMessagesCode2010, CatalogSyncNewResponseEnvelopeMessagesCode2011, CatalogSyncNewResponseEnvelopeMessagesCode2012, CatalogSyncNewResponseEnvelopeMessagesCode2013, CatalogSyncNewResponseEnvelopeMessagesCode2014, CatalogSyncNewResponseEnvelopeMessagesCode2015, CatalogSyncNewResponseEnvelopeMessagesCode2016, CatalogSyncNewResponseEnvelopeMessagesCode2017, CatalogSyncNewResponseEnvelopeMessagesCode2018, CatalogSyncNewResponseEnvelopeMessagesCode2019, CatalogSyncNewResponseEnvelopeMessagesCode2020, CatalogSyncNewResponseEnvelopeMessagesCode2021, CatalogSyncNewResponseEnvelopeMessagesCode2022, CatalogSyncNewResponseEnvelopeMessagesCode3001, CatalogSyncNewResponseEnvelopeMessagesCode3002, CatalogSyncNewResponseEnvelopeMessagesCode3003, CatalogSyncNewResponseEnvelopeMessagesCode3004, CatalogSyncNewResponseEnvelopeMessagesCode3005, CatalogSyncNewResponseEnvelopeMessagesCode3006, CatalogSyncNewResponseEnvelopeMessagesCode3007, CatalogSyncNewResponseEnvelopeMessagesCode4001, CatalogSyncNewResponseEnvelopeMessagesCode4002, CatalogSyncNewResponseEnvelopeMessagesCode4003, CatalogSyncNewResponseEnvelopeMessagesCode4004, CatalogSyncNewResponseEnvelopeMessagesCode4005, CatalogSyncNewResponseEnvelopeMessagesCode4006, CatalogSyncNewResponseEnvelopeMessagesCode4007, CatalogSyncNewResponseEnvelopeMessagesCode4008, CatalogSyncNewResponseEnvelopeMessagesCode4009, CatalogSyncNewResponseEnvelopeMessagesCode4010, CatalogSyncNewResponseEnvelopeMessagesCode4011, CatalogSyncNewResponseEnvelopeMessagesCode4012, CatalogSyncNewResponseEnvelopeMessagesCode4013, CatalogSyncNewResponseEnvelopeMessagesCode4014, CatalogSyncNewResponseEnvelopeMessagesCode4015, CatalogSyncNewResponseEnvelopeMessagesCode4016, CatalogSyncNewResponseEnvelopeMessagesCode4017, CatalogSyncNewResponseEnvelopeMessagesCode4018, CatalogSyncNewResponseEnvelopeMessagesCode4019, CatalogSyncNewResponseEnvelopeMessagesCode4020, CatalogSyncNewResponseEnvelopeMessagesCode4021, CatalogSyncNewResponseEnvelopeMessagesCode4022, CatalogSyncNewResponseEnvelopeMessagesCode4023, CatalogSyncNewResponseEnvelopeMessagesCode5001, CatalogSyncNewResponseEnvelopeMessagesCode5002, CatalogSyncNewResponseEnvelopeMessagesCode5003, CatalogSyncNewResponseEnvelopeMessagesCode5004, CatalogSyncNewResponseEnvelopeMessagesCode102000, CatalogSyncNewResponseEnvelopeMessagesCode102001, CatalogSyncNewResponseEnvelopeMessagesCode102002, CatalogSyncNewResponseEnvelopeMessagesCode102003, CatalogSyncNewResponseEnvelopeMessagesCode102004, CatalogSyncNewResponseEnvelopeMessagesCode102005, CatalogSyncNewResponseEnvelopeMessagesCode102006, CatalogSyncNewResponseEnvelopeMessagesCode102007, CatalogSyncNewResponseEnvelopeMessagesCode102008, CatalogSyncNewResponseEnvelopeMessagesCode102009, CatalogSyncNewResponseEnvelopeMessagesCode102010, CatalogSyncNewResponseEnvelopeMessagesCode102011, CatalogSyncNewResponseEnvelopeMessagesCode102012, CatalogSyncNewResponseEnvelopeMessagesCode102013, CatalogSyncNewResponseEnvelopeMessagesCode102014, CatalogSyncNewResponseEnvelopeMessagesCode102015, CatalogSyncNewResponseEnvelopeMessagesCode102016, CatalogSyncNewResponseEnvelopeMessagesCode102017, CatalogSyncNewResponseEnvelopeMessagesCode102018, CatalogSyncNewResponseEnvelopeMessagesCode102019, CatalogSyncNewResponseEnvelopeMessagesCode102020, CatalogSyncNewResponseEnvelopeMessagesCode102021, CatalogSyncNewResponseEnvelopeMessagesCode102022, CatalogSyncNewResponseEnvelopeMessagesCode102023, CatalogSyncNewResponseEnvelopeMessagesCode102024, CatalogSyncNewResponseEnvelopeMessagesCode102025, CatalogSyncNewResponseEnvelopeMessagesCode102026, CatalogSyncNewResponseEnvelopeMessagesCode102027, CatalogSyncNewResponseEnvelopeMessagesCode102028, CatalogSyncNewResponseEnvelopeMessagesCode102029, CatalogSyncNewResponseEnvelopeMessagesCode102030, CatalogSyncNewResponseEnvelopeMessagesCode102031, CatalogSyncNewResponseEnvelopeMessagesCode102032, CatalogSyncNewResponseEnvelopeMessagesCode102033, CatalogSyncNewResponseEnvelopeMessagesCode102034, CatalogSyncNewResponseEnvelopeMessagesCode102035, CatalogSyncNewResponseEnvelopeMessagesCode102036, CatalogSyncNewResponseEnvelopeMessagesCode102037, CatalogSyncNewResponseEnvelopeMessagesCode102038, CatalogSyncNewResponseEnvelopeMessagesCode102039, CatalogSyncNewResponseEnvelopeMessagesCode102040, CatalogSyncNewResponseEnvelopeMessagesCode102041, CatalogSyncNewResponseEnvelopeMessagesCode102042, CatalogSyncNewResponseEnvelopeMessagesCode102043, CatalogSyncNewResponseEnvelopeMessagesCode102044, CatalogSyncNewResponseEnvelopeMessagesCode102045, CatalogSyncNewResponseEnvelopeMessagesCode102046, CatalogSyncNewResponseEnvelopeMessagesCode102047, CatalogSyncNewResponseEnvelopeMessagesCode102048, CatalogSyncNewResponseEnvelopeMessagesCode102049, CatalogSyncNewResponseEnvelopeMessagesCode102050, CatalogSyncNewResponseEnvelopeMessagesCode102051, CatalogSyncNewResponseEnvelopeMessagesCode102052, CatalogSyncNewResponseEnvelopeMessagesCode102053, CatalogSyncNewResponseEnvelopeMessagesCode102054, CatalogSyncNewResponseEnvelopeMessagesCode102055, CatalogSyncNewResponseEnvelopeMessagesCode102056, CatalogSyncNewResponseEnvelopeMessagesCode102057, CatalogSyncNewResponseEnvelopeMessagesCode102058, CatalogSyncNewResponseEnvelopeMessagesCode102059, CatalogSyncNewResponseEnvelopeMessagesCode102060, CatalogSyncNewResponseEnvelopeMessagesCode102061, CatalogSyncNewResponseEnvelopeMessagesCode102062, CatalogSyncNewResponseEnvelopeMessagesCode102063, CatalogSyncNewResponseEnvelopeMessagesCode102064, CatalogSyncNewResponseEnvelopeMessagesCode102065, CatalogSyncNewResponseEnvelopeMessagesCode102066, CatalogSyncNewResponseEnvelopeMessagesCode103001, CatalogSyncNewResponseEnvelopeMessagesCode103002, CatalogSyncNewResponseEnvelopeMessagesCode103003, CatalogSyncNewResponseEnvelopeMessagesCode103004, CatalogSyncNewResponseEnvelopeMessagesCode103005, CatalogSyncNewResponseEnvelopeMessagesCode103006, CatalogSyncNewResponseEnvelopeMessagesCode103007, CatalogSyncNewResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncNewResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                         `json:"l10n_key"`
	LoggableError string                                         `json:"loggable_error"`
	TemplateData  interface{}                                    `json:"template_data"`
	TraceID       string                                         `json:"trace_id"`
	JSON          catalogSyncNewResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncNewResponseEnvelopeMessagesMetaJSON contains the JSON metadata for
// the struct [CatalogSyncNewResponseEnvelopeMessagesMeta]
type catalogSyncNewResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncNewResponseEnvelopeMessagesSource struct {
	Parameter           string                                           `json:"parameter"`
	ParameterValueIndex int64                                            `json:"parameter_value_index"`
	Pointer             string                                           `json:"pointer"`
	JSON                catalogSyncNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CatalogSyncNewResponseEnvelopeMessagesSource]
type catalogSyncNewResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateParams struct {
	AccountID   param.Field[string]                            `path:"account_id,required"`
	Description param.Field[string]                            `json:"description"`
	Name        param.Field[string]                            `json:"name"`
	Policy      param.Field[string]                            `json:"policy"`
	UpdateMode  param.Field[CatalogSyncUpdateParamsUpdateMode] `json:"update_mode"`
}

func (r CatalogSyncUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CatalogSyncUpdateParamsUpdateMode string

const (
	CatalogSyncUpdateParamsUpdateModeAuto   CatalogSyncUpdateParamsUpdateMode = "AUTO"
	CatalogSyncUpdateParamsUpdateModeManual CatalogSyncUpdateParamsUpdateMode = "MANUAL"
)

func (r CatalogSyncUpdateParamsUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateParamsUpdateModeAuto, CatalogSyncUpdateParamsUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseEnvelope struct {
	Errors   []CatalogSyncUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CatalogSyncUpdateResponse                   `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     catalogSyncUpdateResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncUpdateResponseEnvelope]
type catalogSyncUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeErrors struct {
	Code             CatalogSyncUpdateResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Meta             CatalogSyncUpdateResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncUpdateResponseEnvelopeErrors]
type catalogSyncUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeErrorsCode int64

const (
	CatalogSyncUpdateResponseEnvelopeErrorsCode1001   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1001
	CatalogSyncUpdateResponseEnvelopeErrorsCode1002   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1002
	CatalogSyncUpdateResponseEnvelopeErrorsCode1003   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1003
	CatalogSyncUpdateResponseEnvelopeErrorsCode1004   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1004
	CatalogSyncUpdateResponseEnvelopeErrorsCode1005   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1005
	CatalogSyncUpdateResponseEnvelopeErrorsCode1006   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1006
	CatalogSyncUpdateResponseEnvelopeErrorsCode1007   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1007
	CatalogSyncUpdateResponseEnvelopeErrorsCode1008   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1008
	CatalogSyncUpdateResponseEnvelopeErrorsCode1009   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1009
	CatalogSyncUpdateResponseEnvelopeErrorsCode1010   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1010
	CatalogSyncUpdateResponseEnvelopeErrorsCode1011   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1011
	CatalogSyncUpdateResponseEnvelopeErrorsCode1012   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1012
	CatalogSyncUpdateResponseEnvelopeErrorsCode1013   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1013
	CatalogSyncUpdateResponseEnvelopeErrorsCode1014   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1014
	CatalogSyncUpdateResponseEnvelopeErrorsCode1015   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1015
	CatalogSyncUpdateResponseEnvelopeErrorsCode1016   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1016
	CatalogSyncUpdateResponseEnvelopeErrorsCode1017   CatalogSyncUpdateResponseEnvelopeErrorsCode = 1017
	CatalogSyncUpdateResponseEnvelopeErrorsCode2001   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2001
	CatalogSyncUpdateResponseEnvelopeErrorsCode2002   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2002
	CatalogSyncUpdateResponseEnvelopeErrorsCode2003   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2003
	CatalogSyncUpdateResponseEnvelopeErrorsCode2004   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2004
	CatalogSyncUpdateResponseEnvelopeErrorsCode2005   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2005
	CatalogSyncUpdateResponseEnvelopeErrorsCode2006   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2006
	CatalogSyncUpdateResponseEnvelopeErrorsCode2007   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2007
	CatalogSyncUpdateResponseEnvelopeErrorsCode2008   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2008
	CatalogSyncUpdateResponseEnvelopeErrorsCode2009   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2009
	CatalogSyncUpdateResponseEnvelopeErrorsCode2010   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2010
	CatalogSyncUpdateResponseEnvelopeErrorsCode2011   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2011
	CatalogSyncUpdateResponseEnvelopeErrorsCode2012   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2012
	CatalogSyncUpdateResponseEnvelopeErrorsCode2013   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2013
	CatalogSyncUpdateResponseEnvelopeErrorsCode2014   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2014
	CatalogSyncUpdateResponseEnvelopeErrorsCode2015   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2015
	CatalogSyncUpdateResponseEnvelopeErrorsCode2016   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2016
	CatalogSyncUpdateResponseEnvelopeErrorsCode2017   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2017
	CatalogSyncUpdateResponseEnvelopeErrorsCode2018   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2018
	CatalogSyncUpdateResponseEnvelopeErrorsCode2019   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2019
	CatalogSyncUpdateResponseEnvelopeErrorsCode2020   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2020
	CatalogSyncUpdateResponseEnvelopeErrorsCode2021   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2021
	CatalogSyncUpdateResponseEnvelopeErrorsCode2022   CatalogSyncUpdateResponseEnvelopeErrorsCode = 2022
	CatalogSyncUpdateResponseEnvelopeErrorsCode3001   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3001
	CatalogSyncUpdateResponseEnvelopeErrorsCode3002   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3002
	CatalogSyncUpdateResponseEnvelopeErrorsCode3003   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3003
	CatalogSyncUpdateResponseEnvelopeErrorsCode3004   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3004
	CatalogSyncUpdateResponseEnvelopeErrorsCode3005   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3005
	CatalogSyncUpdateResponseEnvelopeErrorsCode3006   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3006
	CatalogSyncUpdateResponseEnvelopeErrorsCode3007   CatalogSyncUpdateResponseEnvelopeErrorsCode = 3007
	CatalogSyncUpdateResponseEnvelopeErrorsCode4001   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4001
	CatalogSyncUpdateResponseEnvelopeErrorsCode4002   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4002
	CatalogSyncUpdateResponseEnvelopeErrorsCode4003   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4003
	CatalogSyncUpdateResponseEnvelopeErrorsCode4004   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4004
	CatalogSyncUpdateResponseEnvelopeErrorsCode4005   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4005
	CatalogSyncUpdateResponseEnvelopeErrorsCode4006   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4006
	CatalogSyncUpdateResponseEnvelopeErrorsCode4007   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4007
	CatalogSyncUpdateResponseEnvelopeErrorsCode4008   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4008
	CatalogSyncUpdateResponseEnvelopeErrorsCode4009   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4009
	CatalogSyncUpdateResponseEnvelopeErrorsCode4010   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4010
	CatalogSyncUpdateResponseEnvelopeErrorsCode4011   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4011
	CatalogSyncUpdateResponseEnvelopeErrorsCode4012   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4012
	CatalogSyncUpdateResponseEnvelopeErrorsCode4013   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4013
	CatalogSyncUpdateResponseEnvelopeErrorsCode4014   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4014
	CatalogSyncUpdateResponseEnvelopeErrorsCode4015   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4015
	CatalogSyncUpdateResponseEnvelopeErrorsCode4016   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4016
	CatalogSyncUpdateResponseEnvelopeErrorsCode4017   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4017
	CatalogSyncUpdateResponseEnvelopeErrorsCode4018   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4018
	CatalogSyncUpdateResponseEnvelopeErrorsCode4019   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4019
	CatalogSyncUpdateResponseEnvelopeErrorsCode4020   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4020
	CatalogSyncUpdateResponseEnvelopeErrorsCode4021   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4021
	CatalogSyncUpdateResponseEnvelopeErrorsCode4022   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4022
	CatalogSyncUpdateResponseEnvelopeErrorsCode4023   CatalogSyncUpdateResponseEnvelopeErrorsCode = 4023
	CatalogSyncUpdateResponseEnvelopeErrorsCode5001   CatalogSyncUpdateResponseEnvelopeErrorsCode = 5001
	CatalogSyncUpdateResponseEnvelopeErrorsCode5002   CatalogSyncUpdateResponseEnvelopeErrorsCode = 5002
	CatalogSyncUpdateResponseEnvelopeErrorsCode5003   CatalogSyncUpdateResponseEnvelopeErrorsCode = 5003
	CatalogSyncUpdateResponseEnvelopeErrorsCode5004   CatalogSyncUpdateResponseEnvelopeErrorsCode = 5004
	CatalogSyncUpdateResponseEnvelopeErrorsCode102000 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102000
	CatalogSyncUpdateResponseEnvelopeErrorsCode102001 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102001
	CatalogSyncUpdateResponseEnvelopeErrorsCode102002 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102002
	CatalogSyncUpdateResponseEnvelopeErrorsCode102003 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102003
	CatalogSyncUpdateResponseEnvelopeErrorsCode102004 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102004
	CatalogSyncUpdateResponseEnvelopeErrorsCode102005 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102005
	CatalogSyncUpdateResponseEnvelopeErrorsCode102006 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102006
	CatalogSyncUpdateResponseEnvelopeErrorsCode102007 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102007
	CatalogSyncUpdateResponseEnvelopeErrorsCode102008 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102008
	CatalogSyncUpdateResponseEnvelopeErrorsCode102009 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102009
	CatalogSyncUpdateResponseEnvelopeErrorsCode102010 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102010
	CatalogSyncUpdateResponseEnvelopeErrorsCode102011 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102011
	CatalogSyncUpdateResponseEnvelopeErrorsCode102012 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102012
	CatalogSyncUpdateResponseEnvelopeErrorsCode102013 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102013
	CatalogSyncUpdateResponseEnvelopeErrorsCode102014 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102014
	CatalogSyncUpdateResponseEnvelopeErrorsCode102015 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102015
	CatalogSyncUpdateResponseEnvelopeErrorsCode102016 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102016
	CatalogSyncUpdateResponseEnvelopeErrorsCode102017 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102017
	CatalogSyncUpdateResponseEnvelopeErrorsCode102018 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102018
	CatalogSyncUpdateResponseEnvelopeErrorsCode102019 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102019
	CatalogSyncUpdateResponseEnvelopeErrorsCode102020 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102020
	CatalogSyncUpdateResponseEnvelopeErrorsCode102021 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102021
	CatalogSyncUpdateResponseEnvelopeErrorsCode102022 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102022
	CatalogSyncUpdateResponseEnvelopeErrorsCode102023 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102023
	CatalogSyncUpdateResponseEnvelopeErrorsCode102024 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102024
	CatalogSyncUpdateResponseEnvelopeErrorsCode102025 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102025
	CatalogSyncUpdateResponseEnvelopeErrorsCode102026 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102026
	CatalogSyncUpdateResponseEnvelopeErrorsCode102027 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102027
	CatalogSyncUpdateResponseEnvelopeErrorsCode102028 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102028
	CatalogSyncUpdateResponseEnvelopeErrorsCode102029 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102029
	CatalogSyncUpdateResponseEnvelopeErrorsCode102030 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102030
	CatalogSyncUpdateResponseEnvelopeErrorsCode102031 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102031
	CatalogSyncUpdateResponseEnvelopeErrorsCode102032 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102032
	CatalogSyncUpdateResponseEnvelopeErrorsCode102033 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102033
	CatalogSyncUpdateResponseEnvelopeErrorsCode102034 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102034
	CatalogSyncUpdateResponseEnvelopeErrorsCode102035 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102035
	CatalogSyncUpdateResponseEnvelopeErrorsCode102036 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102036
	CatalogSyncUpdateResponseEnvelopeErrorsCode102037 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102037
	CatalogSyncUpdateResponseEnvelopeErrorsCode102038 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102038
	CatalogSyncUpdateResponseEnvelopeErrorsCode102039 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102039
	CatalogSyncUpdateResponseEnvelopeErrorsCode102040 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102040
	CatalogSyncUpdateResponseEnvelopeErrorsCode102041 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102041
	CatalogSyncUpdateResponseEnvelopeErrorsCode102042 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102042
	CatalogSyncUpdateResponseEnvelopeErrorsCode102043 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102043
	CatalogSyncUpdateResponseEnvelopeErrorsCode102044 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102044
	CatalogSyncUpdateResponseEnvelopeErrorsCode102045 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102045
	CatalogSyncUpdateResponseEnvelopeErrorsCode102046 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102046
	CatalogSyncUpdateResponseEnvelopeErrorsCode102047 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102047
	CatalogSyncUpdateResponseEnvelopeErrorsCode102048 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102048
	CatalogSyncUpdateResponseEnvelopeErrorsCode102049 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102049
	CatalogSyncUpdateResponseEnvelopeErrorsCode102050 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102050
	CatalogSyncUpdateResponseEnvelopeErrorsCode102051 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102051
	CatalogSyncUpdateResponseEnvelopeErrorsCode102052 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102052
	CatalogSyncUpdateResponseEnvelopeErrorsCode102053 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102053
	CatalogSyncUpdateResponseEnvelopeErrorsCode102054 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102054
	CatalogSyncUpdateResponseEnvelopeErrorsCode102055 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102055
	CatalogSyncUpdateResponseEnvelopeErrorsCode102056 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102056
	CatalogSyncUpdateResponseEnvelopeErrorsCode102057 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102057
	CatalogSyncUpdateResponseEnvelopeErrorsCode102058 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102058
	CatalogSyncUpdateResponseEnvelopeErrorsCode102059 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102059
	CatalogSyncUpdateResponseEnvelopeErrorsCode102060 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102060
	CatalogSyncUpdateResponseEnvelopeErrorsCode102061 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102061
	CatalogSyncUpdateResponseEnvelopeErrorsCode102062 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102062
	CatalogSyncUpdateResponseEnvelopeErrorsCode102063 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102063
	CatalogSyncUpdateResponseEnvelopeErrorsCode102064 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102064
	CatalogSyncUpdateResponseEnvelopeErrorsCode102065 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102065
	CatalogSyncUpdateResponseEnvelopeErrorsCode102066 CatalogSyncUpdateResponseEnvelopeErrorsCode = 102066
	CatalogSyncUpdateResponseEnvelopeErrorsCode103001 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103001
	CatalogSyncUpdateResponseEnvelopeErrorsCode103002 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103002
	CatalogSyncUpdateResponseEnvelopeErrorsCode103003 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103003
	CatalogSyncUpdateResponseEnvelopeErrorsCode103004 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103004
	CatalogSyncUpdateResponseEnvelopeErrorsCode103005 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103005
	CatalogSyncUpdateResponseEnvelopeErrorsCode103006 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103006
	CatalogSyncUpdateResponseEnvelopeErrorsCode103007 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103007
	CatalogSyncUpdateResponseEnvelopeErrorsCode103008 CatalogSyncUpdateResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncUpdateResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateResponseEnvelopeErrorsCode1001, CatalogSyncUpdateResponseEnvelopeErrorsCode1002, CatalogSyncUpdateResponseEnvelopeErrorsCode1003, CatalogSyncUpdateResponseEnvelopeErrorsCode1004, CatalogSyncUpdateResponseEnvelopeErrorsCode1005, CatalogSyncUpdateResponseEnvelopeErrorsCode1006, CatalogSyncUpdateResponseEnvelopeErrorsCode1007, CatalogSyncUpdateResponseEnvelopeErrorsCode1008, CatalogSyncUpdateResponseEnvelopeErrorsCode1009, CatalogSyncUpdateResponseEnvelopeErrorsCode1010, CatalogSyncUpdateResponseEnvelopeErrorsCode1011, CatalogSyncUpdateResponseEnvelopeErrorsCode1012, CatalogSyncUpdateResponseEnvelopeErrorsCode1013, CatalogSyncUpdateResponseEnvelopeErrorsCode1014, CatalogSyncUpdateResponseEnvelopeErrorsCode1015, CatalogSyncUpdateResponseEnvelopeErrorsCode1016, CatalogSyncUpdateResponseEnvelopeErrorsCode1017, CatalogSyncUpdateResponseEnvelopeErrorsCode2001, CatalogSyncUpdateResponseEnvelopeErrorsCode2002, CatalogSyncUpdateResponseEnvelopeErrorsCode2003, CatalogSyncUpdateResponseEnvelopeErrorsCode2004, CatalogSyncUpdateResponseEnvelopeErrorsCode2005, CatalogSyncUpdateResponseEnvelopeErrorsCode2006, CatalogSyncUpdateResponseEnvelopeErrorsCode2007, CatalogSyncUpdateResponseEnvelopeErrorsCode2008, CatalogSyncUpdateResponseEnvelopeErrorsCode2009, CatalogSyncUpdateResponseEnvelopeErrorsCode2010, CatalogSyncUpdateResponseEnvelopeErrorsCode2011, CatalogSyncUpdateResponseEnvelopeErrorsCode2012, CatalogSyncUpdateResponseEnvelopeErrorsCode2013, CatalogSyncUpdateResponseEnvelopeErrorsCode2014, CatalogSyncUpdateResponseEnvelopeErrorsCode2015, CatalogSyncUpdateResponseEnvelopeErrorsCode2016, CatalogSyncUpdateResponseEnvelopeErrorsCode2017, CatalogSyncUpdateResponseEnvelopeErrorsCode2018, CatalogSyncUpdateResponseEnvelopeErrorsCode2019, CatalogSyncUpdateResponseEnvelopeErrorsCode2020, CatalogSyncUpdateResponseEnvelopeErrorsCode2021, CatalogSyncUpdateResponseEnvelopeErrorsCode2022, CatalogSyncUpdateResponseEnvelopeErrorsCode3001, CatalogSyncUpdateResponseEnvelopeErrorsCode3002, CatalogSyncUpdateResponseEnvelopeErrorsCode3003, CatalogSyncUpdateResponseEnvelopeErrorsCode3004, CatalogSyncUpdateResponseEnvelopeErrorsCode3005, CatalogSyncUpdateResponseEnvelopeErrorsCode3006, CatalogSyncUpdateResponseEnvelopeErrorsCode3007, CatalogSyncUpdateResponseEnvelopeErrorsCode4001, CatalogSyncUpdateResponseEnvelopeErrorsCode4002, CatalogSyncUpdateResponseEnvelopeErrorsCode4003, CatalogSyncUpdateResponseEnvelopeErrorsCode4004, CatalogSyncUpdateResponseEnvelopeErrorsCode4005, CatalogSyncUpdateResponseEnvelopeErrorsCode4006, CatalogSyncUpdateResponseEnvelopeErrorsCode4007, CatalogSyncUpdateResponseEnvelopeErrorsCode4008, CatalogSyncUpdateResponseEnvelopeErrorsCode4009, CatalogSyncUpdateResponseEnvelopeErrorsCode4010, CatalogSyncUpdateResponseEnvelopeErrorsCode4011, CatalogSyncUpdateResponseEnvelopeErrorsCode4012, CatalogSyncUpdateResponseEnvelopeErrorsCode4013, CatalogSyncUpdateResponseEnvelopeErrorsCode4014, CatalogSyncUpdateResponseEnvelopeErrorsCode4015, CatalogSyncUpdateResponseEnvelopeErrorsCode4016, CatalogSyncUpdateResponseEnvelopeErrorsCode4017, CatalogSyncUpdateResponseEnvelopeErrorsCode4018, CatalogSyncUpdateResponseEnvelopeErrorsCode4019, CatalogSyncUpdateResponseEnvelopeErrorsCode4020, CatalogSyncUpdateResponseEnvelopeErrorsCode4021, CatalogSyncUpdateResponseEnvelopeErrorsCode4022, CatalogSyncUpdateResponseEnvelopeErrorsCode4023, CatalogSyncUpdateResponseEnvelopeErrorsCode5001, CatalogSyncUpdateResponseEnvelopeErrorsCode5002, CatalogSyncUpdateResponseEnvelopeErrorsCode5003, CatalogSyncUpdateResponseEnvelopeErrorsCode5004, CatalogSyncUpdateResponseEnvelopeErrorsCode102000, CatalogSyncUpdateResponseEnvelopeErrorsCode102001, CatalogSyncUpdateResponseEnvelopeErrorsCode102002, CatalogSyncUpdateResponseEnvelopeErrorsCode102003, CatalogSyncUpdateResponseEnvelopeErrorsCode102004, CatalogSyncUpdateResponseEnvelopeErrorsCode102005, CatalogSyncUpdateResponseEnvelopeErrorsCode102006, CatalogSyncUpdateResponseEnvelopeErrorsCode102007, CatalogSyncUpdateResponseEnvelopeErrorsCode102008, CatalogSyncUpdateResponseEnvelopeErrorsCode102009, CatalogSyncUpdateResponseEnvelopeErrorsCode102010, CatalogSyncUpdateResponseEnvelopeErrorsCode102011, CatalogSyncUpdateResponseEnvelopeErrorsCode102012, CatalogSyncUpdateResponseEnvelopeErrorsCode102013, CatalogSyncUpdateResponseEnvelopeErrorsCode102014, CatalogSyncUpdateResponseEnvelopeErrorsCode102015, CatalogSyncUpdateResponseEnvelopeErrorsCode102016, CatalogSyncUpdateResponseEnvelopeErrorsCode102017, CatalogSyncUpdateResponseEnvelopeErrorsCode102018, CatalogSyncUpdateResponseEnvelopeErrorsCode102019, CatalogSyncUpdateResponseEnvelopeErrorsCode102020, CatalogSyncUpdateResponseEnvelopeErrorsCode102021, CatalogSyncUpdateResponseEnvelopeErrorsCode102022, CatalogSyncUpdateResponseEnvelopeErrorsCode102023, CatalogSyncUpdateResponseEnvelopeErrorsCode102024, CatalogSyncUpdateResponseEnvelopeErrorsCode102025, CatalogSyncUpdateResponseEnvelopeErrorsCode102026, CatalogSyncUpdateResponseEnvelopeErrorsCode102027, CatalogSyncUpdateResponseEnvelopeErrorsCode102028, CatalogSyncUpdateResponseEnvelopeErrorsCode102029, CatalogSyncUpdateResponseEnvelopeErrorsCode102030, CatalogSyncUpdateResponseEnvelopeErrorsCode102031, CatalogSyncUpdateResponseEnvelopeErrorsCode102032, CatalogSyncUpdateResponseEnvelopeErrorsCode102033, CatalogSyncUpdateResponseEnvelopeErrorsCode102034, CatalogSyncUpdateResponseEnvelopeErrorsCode102035, CatalogSyncUpdateResponseEnvelopeErrorsCode102036, CatalogSyncUpdateResponseEnvelopeErrorsCode102037, CatalogSyncUpdateResponseEnvelopeErrorsCode102038, CatalogSyncUpdateResponseEnvelopeErrorsCode102039, CatalogSyncUpdateResponseEnvelopeErrorsCode102040, CatalogSyncUpdateResponseEnvelopeErrorsCode102041, CatalogSyncUpdateResponseEnvelopeErrorsCode102042, CatalogSyncUpdateResponseEnvelopeErrorsCode102043, CatalogSyncUpdateResponseEnvelopeErrorsCode102044, CatalogSyncUpdateResponseEnvelopeErrorsCode102045, CatalogSyncUpdateResponseEnvelopeErrorsCode102046, CatalogSyncUpdateResponseEnvelopeErrorsCode102047, CatalogSyncUpdateResponseEnvelopeErrorsCode102048, CatalogSyncUpdateResponseEnvelopeErrorsCode102049, CatalogSyncUpdateResponseEnvelopeErrorsCode102050, CatalogSyncUpdateResponseEnvelopeErrorsCode102051, CatalogSyncUpdateResponseEnvelopeErrorsCode102052, CatalogSyncUpdateResponseEnvelopeErrorsCode102053, CatalogSyncUpdateResponseEnvelopeErrorsCode102054, CatalogSyncUpdateResponseEnvelopeErrorsCode102055, CatalogSyncUpdateResponseEnvelopeErrorsCode102056, CatalogSyncUpdateResponseEnvelopeErrorsCode102057, CatalogSyncUpdateResponseEnvelopeErrorsCode102058, CatalogSyncUpdateResponseEnvelopeErrorsCode102059, CatalogSyncUpdateResponseEnvelopeErrorsCode102060, CatalogSyncUpdateResponseEnvelopeErrorsCode102061, CatalogSyncUpdateResponseEnvelopeErrorsCode102062, CatalogSyncUpdateResponseEnvelopeErrorsCode102063, CatalogSyncUpdateResponseEnvelopeErrorsCode102064, CatalogSyncUpdateResponseEnvelopeErrorsCode102065, CatalogSyncUpdateResponseEnvelopeErrorsCode102066, CatalogSyncUpdateResponseEnvelopeErrorsCode103001, CatalogSyncUpdateResponseEnvelopeErrorsCode103002, CatalogSyncUpdateResponseEnvelopeErrorsCode103003, CatalogSyncUpdateResponseEnvelopeErrorsCode103004, CatalogSyncUpdateResponseEnvelopeErrorsCode103005, CatalogSyncUpdateResponseEnvelopeErrorsCode103006, CatalogSyncUpdateResponseEnvelopeErrorsCode103007, CatalogSyncUpdateResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                          `json:"l10n_key"`
	LoggableError string                                          `json:"loggable_error"`
	TemplateData  interface{}                                     `json:"template_data"`
	TraceID       string                                          `json:"trace_id"`
	JSON          catalogSyncUpdateResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeErrorsMetaJSON contains the JSON metadata for
// the struct [CatalogSyncUpdateResponseEnvelopeErrorsMeta]
type catalogSyncUpdateResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeErrorsSource struct {
	Parameter           string                                            `json:"parameter"`
	ParameterValueIndex int64                                             `json:"parameter_value_index"`
	Pointer             string                                            `json:"pointer"`
	JSON                catalogSyncUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CatalogSyncUpdateResponseEnvelopeErrorsSource]
type catalogSyncUpdateResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeMessages struct {
	Code             CatalogSyncUpdateResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Meta             CatalogSyncUpdateResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CatalogSyncUpdateResponseEnvelopeMessages]
type catalogSyncUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeMessagesCode int64

const (
	CatalogSyncUpdateResponseEnvelopeMessagesCode1001   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1001
	CatalogSyncUpdateResponseEnvelopeMessagesCode1002   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1002
	CatalogSyncUpdateResponseEnvelopeMessagesCode1003   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1003
	CatalogSyncUpdateResponseEnvelopeMessagesCode1004   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1004
	CatalogSyncUpdateResponseEnvelopeMessagesCode1005   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1005
	CatalogSyncUpdateResponseEnvelopeMessagesCode1006   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1006
	CatalogSyncUpdateResponseEnvelopeMessagesCode1007   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1007
	CatalogSyncUpdateResponseEnvelopeMessagesCode1008   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1008
	CatalogSyncUpdateResponseEnvelopeMessagesCode1009   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1009
	CatalogSyncUpdateResponseEnvelopeMessagesCode1010   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1010
	CatalogSyncUpdateResponseEnvelopeMessagesCode1011   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1011
	CatalogSyncUpdateResponseEnvelopeMessagesCode1012   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1012
	CatalogSyncUpdateResponseEnvelopeMessagesCode1013   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1013
	CatalogSyncUpdateResponseEnvelopeMessagesCode1014   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1014
	CatalogSyncUpdateResponseEnvelopeMessagesCode1015   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1015
	CatalogSyncUpdateResponseEnvelopeMessagesCode1016   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1016
	CatalogSyncUpdateResponseEnvelopeMessagesCode1017   CatalogSyncUpdateResponseEnvelopeMessagesCode = 1017
	CatalogSyncUpdateResponseEnvelopeMessagesCode2001   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2001
	CatalogSyncUpdateResponseEnvelopeMessagesCode2002   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2002
	CatalogSyncUpdateResponseEnvelopeMessagesCode2003   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2003
	CatalogSyncUpdateResponseEnvelopeMessagesCode2004   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2004
	CatalogSyncUpdateResponseEnvelopeMessagesCode2005   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2005
	CatalogSyncUpdateResponseEnvelopeMessagesCode2006   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2006
	CatalogSyncUpdateResponseEnvelopeMessagesCode2007   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2007
	CatalogSyncUpdateResponseEnvelopeMessagesCode2008   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2008
	CatalogSyncUpdateResponseEnvelopeMessagesCode2009   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2009
	CatalogSyncUpdateResponseEnvelopeMessagesCode2010   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2010
	CatalogSyncUpdateResponseEnvelopeMessagesCode2011   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2011
	CatalogSyncUpdateResponseEnvelopeMessagesCode2012   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2012
	CatalogSyncUpdateResponseEnvelopeMessagesCode2013   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2013
	CatalogSyncUpdateResponseEnvelopeMessagesCode2014   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2014
	CatalogSyncUpdateResponseEnvelopeMessagesCode2015   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2015
	CatalogSyncUpdateResponseEnvelopeMessagesCode2016   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2016
	CatalogSyncUpdateResponseEnvelopeMessagesCode2017   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2017
	CatalogSyncUpdateResponseEnvelopeMessagesCode2018   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2018
	CatalogSyncUpdateResponseEnvelopeMessagesCode2019   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2019
	CatalogSyncUpdateResponseEnvelopeMessagesCode2020   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2020
	CatalogSyncUpdateResponseEnvelopeMessagesCode2021   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2021
	CatalogSyncUpdateResponseEnvelopeMessagesCode2022   CatalogSyncUpdateResponseEnvelopeMessagesCode = 2022
	CatalogSyncUpdateResponseEnvelopeMessagesCode3001   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3001
	CatalogSyncUpdateResponseEnvelopeMessagesCode3002   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3002
	CatalogSyncUpdateResponseEnvelopeMessagesCode3003   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3003
	CatalogSyncUpdateResponseEnvelopeMessagesCode3004   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3004
	CatalogSyncUpdateResponseEnvelopeMessagesCode3005   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3005
	CatalogSyncUpdateResponseEnvelopeMessagesCode3006   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3006
	CatalogSyncUpdateResponseEnvelopeMessagesCode3007   CatalogSyncUpdateResponseEnvelopeMessagesCode = 3007
	CatalogSyncUpdateResponseEnvelopeMessagesCode4001   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4001
	CatalogSyncUpdateResponseEnvelopeMessagesCode4002   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4002
	CatalogSyncUpdateResponseEnvelopeMessagesCode4003   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4003
	CatalogSyncUpdateResponseEnvelopeMessagesCode4004   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4004
	CatalogSyncUpdateResponseEnvelopeMessagesCode4005   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4005
	CatalogSyncUpdateResponseEnvelopeMessagesCode4006   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4006
	CatalogSyncUpdateResponseEnvelopeMessagesCode4007   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4007
	CatalogSyncUpdateResponseEnvelopeMessagesCode4008   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4008
	CatalogSyncUpdateResponseEnvelopeMessagesCode4009   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4009
	CatalogSyncUpdateResponseEnvelopeMessagesCode4010   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4010
	CatalogSyncUpdateResponseEnvelopeMessagesCode4011   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4011
	CatalogSyncUpdateResponseEnvelopeMessagesCode4012   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4012
	CatalogSyncUpdateResponseEnvelopeMessagesCode4013   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4013
	CatalogSyncUpdateResponseEnvelopeMessagesCode4014   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4014
	CatalogSyncUpdateResponseEnvelopeMessagesCode4015   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4015
	CatalogSyncUpdateResponseEnvelopeMessagesCode4016   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4016
	CatalogSyncUpdateResponseEnvelopeMessagesCode4017   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4017
	CatalogSyncUpdateResponseEnvelopeMessagesCode4018   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4018
	CatalogSyncUpdateResponseEnvelopeMessagesCode4019   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4019
	CatalogSyncUpdateResponseEnvelopeMessagesCode4020   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4020
	CatalogSyncUpdateResponseEnvelopeMessagesCode4021   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4021
	CatalogSyncUpdateResponseEnvelopeMessagesCode4022   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4022
	CatalogSyncUpdateResponseEnvelopeMessagesCode4023   CatalogSyncUpdateResponseEnvelopeMessagesCode = 4023
	CatalogSyncUpdateResponseEnvelopeMessagesCode5001   CatalogSyncUpdateResponseEnvelopeMessagesCode = 5001
	CatalogSyncUpdateResponseEnvelopeMessagesCode5002   CatalogSyncUpdateResponseEnvelopeMessagesCode = 5002
	CatalogSyncUpdateResponseEnvelopeMessagesCode5003   CatalogSyncUpdateResponseEnvelopeMessagesCode = 5003
	CatalogSyncUpdateResponseEnvelopeMessagesCode5004   CatalogSyncUpdateResponseEnvelopeMessagesCode = 5004
	CatalogSyncUpdateResponseEnvelopeMessagesCode102000 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102000
	CatalogSyncUpdateResponseEnvelopeMessagesCode102001 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102001
	CatalogSyncUpdateResponseEnvelopeMessagesCode102002 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102002
	CatalogSyncUpdateResponseEnvelopeMessagesCode102003 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102003
	CatalogSyncUpdateResponseEnvelopeMessagesCode102004 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102004
	CatalogSyncUpdateResponseEnvelopeMessagesCode102005 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102005
	CatalogSyncUpdateResponseEnvelopeMessagesCode102006 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102006
	CatalogSyncUpdateResponseEnvelopeMessagesCode102007 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102007
	CatalogSyncUpdateResponseEnvelopeMessagesCode102008 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102008
	CatalogSyncUpdateResponseEnvelopeMessagesCode102009 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102009
	CatalogSyncUpdateResponseEnvelopeMessagesCode102010 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102010
	CatalogSyncUpdateResponseEnvelopeMessagesCode102011 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102011
	CatalogSyncUpdateResponseEnvelopeMessagesCode102012 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102012
	CatalogSyncUpdateResponseEnvelopeMessagesCode102013 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102013
	CatalogSyncUpdateResponseEnvelopeMessagesCode102014 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102014
	CatalogSyncUpdateResponseEnvelopeMessagesCode102015 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102015
	CatalogSyncUpdateResponseEnvelopeMessagesCode102016 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102016
	CatalogSyncUpdateResponseEnvelopeMessagesCode102017 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102017
	CatalogSyncUpdateResponseEnvelopeMessagesCode102018 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102018
	CatalogSyncUpdateResponseEnvelopeMessagesCode102019 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102019
	CatalogSyncUpdateResponseEnvelopeMessagesCode102020 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102020
	CatalogSyncUpdateResponseEnvelopeMessagesCode102021 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102021
	CatalogSyncUpdateResponseEnvelopeMessagesCode102022 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102022
	CatalogSyncUpdateResponseEnvelopeMessagesCode102023 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102023
	CatalogSyncUpdateResponseEnvelopeMessagesCode102024 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102024
	CatalogSyncUpdateResponseEnvelopeMessagesCode102025 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102025
	CatalogSyncUpdateResponseEnvelopeMessagesCode102026 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102026
	CatalogSyncUpdateResponseEnvelopeMessagesCode102027 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102027
	CatalogSyncUpdateResponseEnvelopeMessagesCode102028 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102028
	CatalogSyncUpdateResponseEnvelopeMessagesCode102029 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102029
	CatalogSyncUpdateResponseEnvelopeMessagesCode102030 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102030
	CatalogSyncUpdateResponseEnvelopeMessagesCode102031 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102031
	CatalogSyncUpdateResponseEnvelopeMessagesCode102032 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102032
	CatalogSyncUpdateResponseEnvelopeMessagesCode102033 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102033
	CatalogSyncUpdateResponseEnvelopeMessagesCode102034 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102034
	CatalogSyncUpdateResponseEnvelopeMessagesCode102035 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102035
	CatalogSyncUpdateResponseEnvelopeMessagesCode102036 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102036
	CatalogSyncUpdateResponseEnvelopeMessagesCode102037 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102037
	CatalogSyncUpdateResponseEnvelopeMessagesCode102038 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102038
	CatalogSyncUpdateResponseEnvelopeMessagesCode102039 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102039
	CatalogSyncUpdateResponseEnvelopeMessagesCode102040 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102040
	CatalogSyncUpdateResponseEnvelopeMessagesCode102041 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102041
	CatalogSyncUpdateResponseEnvelopeMessagesCode102042 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102042
	CatalogSyncUpdateResponseEnvelopeMessagesCode102043 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102043
	CatalogSyncUpdateResponseEnvelopeMessagesCode102044 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102044
	CatalogSyncUpdateResponseEnvelopeMessagesCode102045 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102045
	CatalogSyncUpdateResponseEnvelopeMessagesCode102046 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102046
	CatalogSyncUpdateResponseEnvelopeMessagesCode102047 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102047
	CatalogSyncUpdateResponseEnvelopeMessagesCode102048 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102048
	CatalogSyncUpdateResponseEnvelopeMessagesCode102049 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102049
	CatalogSyncUpdateResponseEnvelopeMessagesCode102050 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102050
	CatalogSyncUpdateResponseEnvelopeMessagesCode102051 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102051
	CatalogSyncUpdateResponseEnvelopeMessagesCode102052 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102052
	CatalogSyncUpdateResponseEnvelopeMessagesCode102053 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102053
	CatalogSyncUpdateResponseEnvelopeMessagesCode102054 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102054
	CatalogSyncUpdateResponseEnvelopeMessagesCode102055 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102055
	CatalogSyncUpdateResponseEnvelopeMessagesCode102056 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102056
	CatalogSyncUpdateResponseEnvelopeMessagesCode102057 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102057
	CatalogSyncUpdateResponseEnvelopeMessagesCode102058 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102058
	CatalogSyncUpdateResponseEnvelopeMessagesCode102059 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102059
	CatalogSyncUpdateResponseEnvelopeMessagesCode102060 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102060
	CatalogSyncUpdateResponseEnvelopeMessagesCode102061 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102061
	CatalogSyncUpdateResponseEnvelopeMessagesCode102062 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102062
	CatalogSyncUpdateResponseEnvelopeMessagesCode102063 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102063
	CatalogSyncUpdateResponseEnvelopeMessagesCode102064 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102064
	CatalogSyncUpdateResponseEnvelopeMessagesCode102065 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102065
	CatalogSyncUpdateResponseEnvelopeMessagesCode102066 CatalogSyncUpdateResponseEnvelopeMessagesCode = 102066
	CatalogSyncUpdateResponseEnvelopeMessagesCode103001 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103001
	CatalogSyncUpdateResponseEnvelopeMessagesCode103002 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103002
	CatalogSyncUpdateResponseEnvelopeMessagesCode103003 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103003
	CatalogSyncUpdateResponseEnvelopeMessagesCode103004 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103004
	CatalogSyncUpdateResponseEnvelopeMessagesCode103005 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103005
	CatalogSyncUpdateResponseEnvelopeMessagesCode103006 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103006
	CatalogSyncUpdateResponseEnvelopeMessagesCode103007 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103007
	CatalogSyncUpdateResponseEnvelopeMessagesCode103008 CatalogSyncUpdateResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncUpdateResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncUpdateResponseEnvelopeMessagesCode1001, CatalogSyncUpdateResponseEnvelopeMessagesCode1002, CatalogSyncUpdateResponseEnvelopeMessagesCode1003, CatalogSyncUpdateResponseEnvelopeMessagesCode1004, CatalogSyncUpdateResponseEnvelopeMessagesCode1005, CatalogSyncUpdateResponseEnvelopeMessagesCode1006, CatalogSyncUpdateResponseEnvelopeMessagesCode1007, CatalogSyncUpdateResponseEnvelopeMessagesCode1008, CatalogSyncUpdateResponseEnvelopeMessagesCode1009, CatalogSyncUpdateResponseEnvelopeMessagesCode1010, CatalogSyncUpdateResponseEnvelopeMessagesCode1011, CatalogSyncUpdateResponseEnvelopeMessagesCode1012, CatalogSyncUpdateResponseEnvelopeMessagesCode1013, CatalogSyncUpdateResponseEnvelopeMessagesCode1014, CatalogSyncUpdateResponseEnvelopeMessagesCode1015, CatalogSyncUpdateResponseEnvelopeMessagesCode1016, CatalogSyncUpdateResponseEnvelopeMessagesCode1017, CatalogSyncUpdateResponseEnvelopeMessagesCode2001, CatalogSyncUpdateResponseEnvelopeMessagesCode2002, CatalogSyncUpdateResponseEnvelopeMessagesCode2003, CatalogSyncUpdateResponseEnvelopeMessagesCode2004, CatalogSyncUpdateResponseEnvelopeMessagesCode2005, CatalogSyncUpdateResponseEnvelopeMessagesCode2006, CatalogSyncUpdateResponseEnvelopeMessagesCode2007, CatalogSyncUpdateResponseEnvelopeMessagesCode2008, CatalogSyncUpdateResponseEnvelopeMessagesCode2009, CatalogSyncUpdateResponseEnvelopeMessagesCode2010, CatalogSyncUpdateResponseEnvelopeMessagesCode2011, CatalogSyncUpdateResponseEnvelopeMessagesCode2012, CatalogSyncUpdateResponseEnvelopeMessagesCode2013, CatalogSyncUpdateResponseEnvelopeMessagesCode2014, CatalogSyncUpdateResponseEnvelopeMessagesCode2015, CatalogSyncUpdateResponseEnvelopeMessagesCode2016, CatalogSyncUpdateResponseEnvelopeMessagesCode2017, CatalogSyncUpdateResponseEnvelopeMessagesCode2018, CatalogSyncUpdateResponseEnvelopeMessagesCode2019, CatalogSyncUpdateResponseEnvelopeMessagesCode2020, CatalogSyncUpdateResponseEnvelopeMessagesCode2021, CatalogSyncUpdateResponseEnvelopeMessagesCode2022, CatalogSyncUpdateResponseEnvelopeMessagesCode3001, CatalogSyncUpdateResponseEnvelopeMessagesCode3002, CatalogSyncUpdateResponseEnvelopeMessagesCode3003, CatalogSyncUpdateResponseEnvelopeMessagesCode3004, CatalogSyncUpdateResponseEnvelopeMessagesCode3005, CatalogSyncUpdateResponseEnvelopeMessagesCode3006, CatalogSyncUpdateResponseEnvelopeMessagesCode3007, CatalogSyncUpdateResponseEnvelopeMessagesCode4001, CatalogSyncUpdateResponseEnvelopeMessagesCode4002, CatalogSyncUpdateResponseEnvelopeMessagesCode4003, CatalogSyncUpdateResponseEnvelopeMessagesCode4004, CatalogSyncUpdateResponseEnvelopeMessagesCode4005, CatalogSyncUpdateResponseEnvelopeMessagesCode4006, CatalogSyncUpdateResponseEnvelopeMessagesCode4007, CatalogSyncUpdateResponseEnvelopeMessagesCode4008, CatalogSyncUpdateResponseEnvelopeMessagesCode4009, CatalogSyncUpdateResponseEnvelopeMessagesCode4010, CatalogSyncUpdateResponseEnvelopeMessagesCode4011, CatalogSyncUpdateResponseEnvelopeMessagesCode4012, CatalogSyncUpdateResponseEnvelopeMessagesCode4013, CatalogSyncUpdateResponseEnvelopeMessagesCode4014, CatalogSyncUpdateResponseEnvelopeMessagesCode4015, CatalogSyncUpdateResponseEnvelopeMessagesCode4016, CatalogSyncUpdateResponseEnvelopeMessagesCode4017, CatalogSyncUpdateResponseEnvelopeMessagesCode4018, CatalogSyncUpdateResponseEnvelopeMessagesCode4019, CatalogSyncUpdateResponseEnvelopeMessagesCode4020, CatalogSyncUpdateResponseEnvelopeMessagesCode4021, CatalogSyncUpdateResponseEnvelopeMessagesCode4022, CatalogSyncUpdateResponseEnvelopeMessagesCode4023, CatalogSyncUpdateResponseEnvelopeMessagesCode5001, CatalogSyncUpdateResponseEnvelopeMessagesCode5002, CatalogSyncUpdateResponseEnvelopeMessagesCode5003, CatalogSyncUpdateResponseEnvelopeMessagesCode5004, CatalogSyncUpdateResponseEnvelopeMessagesCode102000, CatalogSyncUpdateResponseEnvelopeMessagesCode102001, CatalogSyncUpdateResponseEnvelopeMessagesCode102002, CatalogSyncUpdateResponseEnvelopeMessagesCode102003, CatalogSyncUpdateResponseEnvelopeMessagesCode102004, CatalogSyncUpdateResponseEnvelopeMessagesCode102005, CatalogSyncUpdateResponseEnvelopeMessagesCode102006, CatalogSyncUpdateResponseEnvelopeMessagesCode102007, CatalogSyncUpdateResponseEnvelopeMessagesCode102008, CatalogSyncUpdateResponseEnvelopeMessagesCode102009, CatalogSyncUpdateResponseEnvelopeMessagesCode102010, CatalogSyncUpdateResponseEnvelopeMessagesCode102011, CatalogSyncUpdateResponseEnvelopeMessagesCode102012, CatalogSyncUpdateResponseEnvelopeMessagesCode102013, CatalogSyncUpdateResponseEnvelopeMessagesCode102014, CatalogSyncUpdateResponseEnvelopeMessagesCode102015, CatalogSyncUpdateResponseEnvelopeMessagesCode102016, CatalogSyncUpdateResponseEnvelopeMessagesCode102017, CatalogSyncUpdateResponseEnvelopeMessagesCode102018, CatalogSyncUpdateResponseEnvelopeMessagesCode102019, CatalogSyncUpdateResponseEnvelopeMessagesCode102020, CatalogSyncUpdateResponseEnvelopeMessagesCode102021, CatalogSyncUpdateResponseEnvelopeMessagesCode102022, CatalogSyncUpdateResponseEnvelopeMessagesCode102023, CatalogSyncUpdateResponseEnvelopeMessagesCode102024, CatalogSyncUpdateResponseEnvelopeMessagesCode102025, CatalogSyncUpdateResponseEnvelopeMessagesCode102026, CatalogSyncUpdateResponseEnvelopeMessagesCode102027, CatalogSyncUpdateResponseEnvelopeMessagesCode102028, CatalogSyncUpdateResponseEnvelopeMessagesCode102029, CatalogSyncUpdateResponseEnvelopeMessagesCode102030, CatalogSyncUpdateResponseEnvelopeMessagesCode102031, CatalogSyncUpdateResponseEnvelopeMessagesCode102032, CatalogSyncUpdateResponseEnvelopeMessagesCode102033, CatalogSyncUpdateResponseEnvelopeMessagesCode102034, CatalogSyncUpdateResponseEnvelopeMessagesCode102035, CatalogSyncUpdateResponseEnvelopeMessagesCode102036, CatalogSyncUpdateResponseEnvelopeMessagesCode102037, CatalogSyncUpdateResponseEnvelopeMessagesCode102038, CatalogSyncUpdateResponseEnvelopeMessagesCode102039, CatalogSyncUpdateResponseEnvelopeMessagesCode102040, CatalogSyncUpdateResponseEnvelopeMessagesCode102041, CatalogSyncUpdateResponseEnvelopeMessagesCode102042, CatalogSyncUpdateResponseEnvelopeMessagesCode102043, CatalogSyncUpdateResponseEnvelopeMessagesCode102044, CatalogSyncUpdateResponseEnvelopeMessagesCode102045, CatalogSyncUpdateResponseEnvelopeMessagesCode102046, CatalogSyncUpdateResponseEnvelopeMessagesCode102047, CatalogSyncUpdateResponseEnvelopeMessagesCode102048, CatalogSyncUpdateResponseEnvelopeMessagesCode102049, CatalogSyncUpdateResponseEnvelopeMessagesCode102050, CatalogSyncUpdateResponseEnvelopeMessagesCode102051, CatalogSyncUpdateResponseEnvelopeMessagesCode102052, CatalogSyncUpdateResponseEnvelopeMessagesCode102053, CatalogSyncUpdateResponseEnvelopeMessagesCode102054, CatalogSyncUpdateResponseEnvelopeMessagesCode102055, CatalogSyncUpdateResponseEnvelopeMessagesCode102056, CatalogSyncUpdateResponseEnvelopeMessagesCode102057, CatalogSyncUpdateResponseEnvelopeMessagesCode102058, CatalogSyncUpdateResponseEnvelopeMessagesCode102059, CatalogSyncUpdateResponseEnvelopeMessagesCode102060, CatalogSyncUpdateResponseEnvelopeMessagesCode102061, CatalogSyncUpdateResponseEnvelopeMessagesCode102062, CatalogSyncUpdateResponseEnvelopeMessagesCode102063, CatalogSyncUpdateResponseEnvelopeMessagesCode102064, CatalogSyncUpdateResponseEnvelopeMessagesCode102065, CatalogSyncUpdateResponseEnvelopeMessagesCode102066, CatalogSyncUpdateResponseEnvelopeMessagesCode103001, CatalogSyncUpdateResponseEnvelopeMessagesCode103002, CatalogSyncUpdateResponseEnvelopeMessagesCode103003, CatalogSyncUpdateResponseEnvelopeMessagesCode103004, CatalogSyncUpdateResponseEnvelopeMessagesCode103005, CatalogSyncUpdateResponseEnvelopeMessagesCode103006, CatalogSyncUpdateResponseEnvelopeMessagesCode103007, CatalogSyncUpdateResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncUpdateResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                            `json:"l10n_key"`
	LoggableError string                                            `json:"loggable_error"`
	TemplateData  interface{}                                       `json:"template_data"`
	TraceID       string                                            `json:"trace_id"`
	JSON          catalogSyncUpdateResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeMessagesMetaJSON contains the JSON metadata for
// the struct [CatalogSyncUpdateResponseEnvelopeMessagesMeta]
type catalogSyncUpdateResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncUpdateResponseEnvelopeMessagesSource struct {
	Parameter           string                                              `json:"parameter"`
	ParameterValueIndex int64                                               `json:"parameter_value_index"`
	Pointer             string                                              `json:"pointer"`
	JSON                catalogSyncUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CatalogSyncUpdateResponseEnvelopeMessagesSource]
type catalogSyncUpdateResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CatalogSyncDeleteParams struct {
	AccountID         param.Field[string] `path:"account_id,required"`
	DeleteDestination param.Field[bool]   `query:"delete_destination"`
}

// URLQuery serializes [CatalogSyncDeleteParams]'s query parameters as
// `url.Values`.
func (r CatalogSyncDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CatalogSyncDeleteResponseEnvelope struct {
	Errors   []CatalogSyncDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CatalogSyncDeleteResponse                   `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     catalogSyncDeleteResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncDeleteResponseEnvelope]
type catalogSyncDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeErrors struct {
	Code             CatalogSyncDeleteResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Meta             CatalogSyncDeleteResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncDeleteResponseEnvelopeErrors]
type catalogSyncDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeErrorsCode int64

const (
	CatalogSyncDeleteResponseEnvelopeErrorsCode1001   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1001
	CatalogSyncDeleteResponseEnvelopeErrorsCode1002   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1002
	CatalogSyncDeleteResponseEnvelopeErrorsCode1003   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1003
	CatalogSyncDeleteResponseEnvelopeErrorsCode1004   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1004
	CatalogSyncDeleteResponseEnvelopeErrorsCode1005   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1005
	CatalogSyncDeleteResponseEnvelopeErrorsCode1006   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1006
	CatalogSyncDeleteResponseEnvelopeErrorsCode1007   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1007
	CatalogSyncDeleteResponseEnvelopeErrorsCode1008   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1008
	CatalogSyncDeleteResponseEnvelopeErrorsCode1009   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1009
	CatalogSyncDeleteResponseEnvelopeErrorsCode1010   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1010
	CatalogSyncDeleteResponseEnvelopeErrorsCode1011   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1011
	CatalogSyncDeleteResponseEnvelopeErrorsCode1012   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1012
	CatalogSyncDeleteResponseEnvelopeErrorsCode1013   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1013
	CatalogSyncDeleteResponseEnvelopeErrorsCode1014   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1014
	CatalogSyncDeleteResponseEnvelopeErrorsCode1015   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1015
	CatalogSyncDeleteResponseEnvelopeErrorsCode1016   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1016
	CatalogSyncDeleteResponseEnvelopeErrorsCode1017   CatalogSyncDeleteResponseEnvelopeErrorsCode = 1017
	CatalogSyncDeleteResponseEnvelopeErrorsCode2001   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2001
	CatalogSyncDeleteResponseEnvelopeErrorsCode2002   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2002
	CatalogSyncDeleteResponseEnvelopeErrorsCode2003   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2003
	CatalogSyncDeleteResponseEnvelopeErrorsCode2004   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2004
	CatalogSyncDeleteResponseEnvelopeErrorsCode2005   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2005
	CatalogSyncDeleteResponseEnvelopeErrorsCode2006   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2006
	CatalogSyncDeleteResponseEnvelopeErrorsCode2007   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2007
	CatalogSyncDeleteResponseEnvelopeErrorsCode2008   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2008
	CatalogSyncDeleteResponseEnvelopeErrorsCode2009   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2009
	CatalogSyncDeleteResponseEnvelopeErrorsCode2010   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2010
	CatalogSyncDeleteResponseEnvelopeErrorsCode2011   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2011
	CatalogSyncDeleteResponseEnvelopeErrorsCode2012   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2012
	CatalogSyncDeleteResponseEnvelopeErrorsCode2013   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2013
	CatalogSyncDeleteResponseEnvelopeErrorsCode2014   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2014
	CatalogSyncDeleteResponseEnvelopeErrorsCode2015   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2015
	CatalogSyncDeleteResponseEnvelopeErrorsCode2016   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2016
	CatalogSyncDeleteResponseEnvelopeErrorsCode2017   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2017
	CatalogSyncDeleteResponseEnvelopeErrorsCode2018   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2018
	CatalogSyncDeleteResponseEnvelopeErrorsCode2019   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2019
	CatalogSyncDeleteResponseEnvelopeErrorsCode2020   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2020
	CatalogSyncDeleteResponseEnvelopeErrorsCode2021   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2021
	CatalogSyncDeleteResponseEnvelopeErrorsCode2022   CatalogSyncDeleteResponseEnvelopeErrorsCode = 2022
	CatalogSyncDeleteResponseEnvelopeErrorsCode3001   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3001
	CatalogSyncDeleteResponseEnvelopeErrorsCode3002   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3002
	CatalogSyncDeleteResponseEnvelopeErrorsCode3003   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3003
	CatalogSyncDeleteResponseEnvelopeErrorsCode3004   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3004
	CatalogSyncDeleteResponseEnvelopeErrorsCode3005   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3005
	CatalogSyncDeleteResponseEnvelopeErrorsCode3006   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3006
	CatalogSyncDeleteResponseEnvelopeErrorsCode3007   CatalogSyncDeleteResponseEnvelopeErrorsCode = 3007
	CatalogSyncDeleteResponseEnvelopeErrorsCode4001   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4001
	CatalogSyncDeleteResponseEnvelopeErrorsCode4002   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4002
	CatalogSyncDeleteResponseEnvelopeErrorsCode4003   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4003
	CatalogSyncDeleteResponseEnvelopeErrorsCode4004   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4004
	CatalogSyncDeleteResponseEnvelopeErrorsCode4005   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4005
	CatalogSyncDeleteResponseEnvelopeErrorsCode4006   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4006
	CatalogSyncDeleteResponseEnvelopeErrorsCode4007   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4007
	CatalogSyncDeleteResponseEnvelopeErrorsCode4008   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4008
	CatalogSyncDeleteResponseEnvelopeErrorsCode4009   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4009
	CatalogSyncDeleteResponseEnvelopeErrorsCode4010   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4010
	CatalogSyncDeleteResponseEnvelopeErrorsCode4011   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4011
	CatalogSyncDeleteResponseEnvelopeErrorsCode4012   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4012
	CatalogSyncDeleteResponseEnvelopeErrorsCode4013   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4013
	CatalogSyncDeleteResponseEnvelopeErrorsCode4014   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4014
	CatalogSyncDeleteResponseEnvelopeErrorsCode4015   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4015
	CatalogSyncDeleteResponseEnvelopeErrorsCode4016   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4016
	CatalogSyncDeleteResponseEnvelopeErrorsCode4017   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4017
	CatalogSyncDeleteResponseEnvelopeErrorsCode4018   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4018
	CatalogSyncDeleteResponseEnvelopeErrorsCode4019   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4019
	CatalogSyncDeleteResponseEnvelopeErrorsCode4020   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4020
	CatalogSyncDeleteResponseEnvelopeErrorsCode4021   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4021
	CatalogSyncDeleteResponseEnvelopeErrorsCode4022   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4022
	CatalogSyncDeleteResponseEnvelopeErrorsCode4023   CatalogSyncDeleteResponseEnvelopeErrorsCode = 4023
	CatalogSyncDeleteResponseEnvelopeErrorsCode5001   CatalogSyncDeleteResponseEnvelopeErrorsCode = 5001
	CatalogSyncDeleteResponseEnvelopeErrorsCode5002   CatalogSyncDeleteResponseEnvelopeErrorsCode = 5002
	CatalogSyncDeleteResponseEnvelopeErrorsCode5003   CatalogSyncDeleteResponseEnvelopeErrorsCode = 5003
	CatalogSyncDeleteResponseEnvelopeErrorsCode5004   CatalogSyncDeleteResponseEnvelopeErrorsCode = 5004
	CatalogSyncDeleteResponseEnvelopeErrorsCode102000 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102000
	CatalogSyncDeleteResponseEnvelopeErrorsCode102001 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102001
	CatalogSyncDeleteResponseEnvelopeErrorsCode102002 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102002
	CatalogSyncDeleteResponseEnvelopeErrorsCode102003 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102003
	CatalogSyncDeleteResponseEnvelopeErrorsCode102004 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102004
	CatalogSyncDeleteResponseEnvelopeErrorsCode102005 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102005
	CatalogSyncDeleteResponseEnvelopeErrorsCode102006 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102006
	CatalogSyncDeleteResponseEnvelopeErrorsCode102007 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102007
	CatalogSyncDeleteResponseEnvelopeErrorsCode102008 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102008
	CatalogSyncDeleteResponseEnvelopeErrorsCode102009 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102009
	CatalogSyncDeleteResponseEnvelopeErrorsCode102010 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102010
	CatalogSyncDeleteResponseEnvelopeErrorsCode102011 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102011
	CatalogSyncDeleteResponseEnvelopeErrorsCode102012 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102012
	CatalogSyncDeleteResponseEnvelopeErrorsCode102013 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102013
	CatalogSyncDeleteResponseEnvelopeErrorsCode102014 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102014
	CatalogSyncDeleteResponseEnvelopeErrorsCode102015 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102015
	CatalogSyncDeleteResponseEnvelopeErrorsCode102016 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102016
	CatalogSyncDeleteResponseEnvelopeErrorsCode102017 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102017
	CatalogSyncDeleteResponseEnvelopeErrorsCode102018 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102018
	CatalogSyncDeleteResponseEnvelopeErrorsCode102019 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102019
	CatalogSyncDeleteResponseEnvelopeErrorsCode102020 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102020
	CatalogSyncDeleteResponseEnvelopeErrorsCode102021 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102021
	CatalogSyncDeleteResponseEnvelopeErrorsCode102022 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102022
	CatalogSyncDeleteResponseEnvelopeErrorsCode102023 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102023
	CatalogSyncDeleteResponseEnvelopeErrorsCode102024 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102024
	CatalogSyncDeleteResponseEnvelopeErrorsCode102025 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102025
	CatalogSyncDeleteResponseEnvelopeErrorsCode102026 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102026
	CatalogSyncDeleteResponseEnvelopeErrorsCode102027 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102027
	CatalogSyncDeleteResponseEnvelopeErrorsCode102028 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102028
	CatalogSyncDeleteResponseEnvelopeErrorsCode102029 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102029
	CatalogSyncDeleteResponseEnvelopeErrorsCode102030 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102030
	CatalogSyncDeleteResponseEnvelopeErrorsCode102031 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102031
	CatalogSyncDeleteResponseEnvelopeErrorsCode102032 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102032
	CatalogSyncDeleteResponseEnvelopeErrorsCode102033 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102033
	CatalogSyncDeleteResponseEnvelopeErrorsCode102034 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102034
	CatalogSyncDeleteResponseEnvelopeErrorsCode102035 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102035
	CatalogSyncDeleteResponseEnvelopeErrorsCode102036 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102036
	CatalogSyncDeleteResponseEnvelopeErrorsCode102037 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102037
	CatalogSyncDeleteResponseEnvelopeErrorsCode102038 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102038
	CatalogSyncDeleteResponseEnvelopeErrorsCode102039 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102039
	CatalogSyncDeleteResponseEnvelopeErrorsCode102040 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102040
	CatalogSyncDeleteResponseEnvelopeErrorsCode102041 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102041
	CatalogSyncDeleteResponseEnvelopeErrorsCode102042 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102042
	CatalogSyncDeleteResponseEnvelopeErrorsCode102043 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102043
	CatalogSyncDeleteResponseEnvelopeErrorsCode102044 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102044
	CatalogSyncDeleteResponseEnvelopeErrorsCode102045 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102045
	CatalogSyncDeleteResponseEnvelopeErrorsCode102046 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102046
	CatalogSyncDeleteResponseEnvelopeErrorsCode102047 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102047
	CatalogSyncDeleteResponseEnvelopeErrorsCode102048 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102048
	CatalogSyncDeleteResponseEnvelopeErrorsCode102049 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102049
	CatalogSyncDeleteResponseEnvelopeErrorsCode102050 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102050
	CatalogSyncDeleteResponseEnvelopeErrorsCode102051 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102051
	CatalogSyncDeleteResponseEnvelopeErrorsCode102052 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102052
	CatalogSyncDeleteResponseEnvelopeErrorsCode102053 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102053
	CatalogSyncDeleteResponseEnvelopeErrorsCode102054 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102054
	CatalogSyncDeleteResponseEnvelopeErrorsCode102055 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102055
	CatalogSyncDeleteResponseEnvelopeErrorsCode102056 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102056
	CatalogSyncDeleteResponseEnvelopeErrorsCode102057 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102057
	CatalogSyncDeleteResponseEnvelopeErrorsCode102058 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102058
	CatalogSyncDeleteResponseEnvelopeErrorsCode102059 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102059
	CatalogSyncDeleteResponseEnvelopeErrorsCode102060 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102060
	CatalogSyncDeleteResponseEnvelopeErrorsCode102061 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102061
	CatalogSyncDeleteResponseEnvelopeErrorsCode102062 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102062
	CatalogSyncDeleteResponseEnvelopeErrorsCode102063 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102063
	CatalogSyncDeleteResponseEnvelopeErrorsCode102064 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102064
	CatalogSyncDeleteResponseEnvelopeErrorsCode102065 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102065
	CatalogSyncDeleteResponseEnvelopeErrorsCode102066 CatalogSyncDeleteResponseEnvelopeErrorsCode = 102066
	CatalogSyncDeleteResponseEnvelopeErrorsCode103001 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103001
	CatalogSyncDeleteResponseEnvelopeErrorsCode103002 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103002
	CatalogSyncDeleteResponseEnvelopeErrorsCode103003 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103003
	CatalogSyncDeleteResponseEnvelopeErrorsCode103004 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103004
	CatalogSyncDeleteResponseEnvelopeErrorsCode103005 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103005
	CatalogSyncDeleteResponseEnvelopeErrorsCode103006 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103006
	CatalogSyncDeleteResponseEnvelopeErrorsCode103007 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103007
	CatalogSyncDeleteResponseEnvelopeErrorsCode103008 CatalogSyncDeleteResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncDeleteResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncDeleteResponseEnvelopeErrorsCode1001, CatalogSyncDeleteResponseEnvelopeErrorsCode1002, CatalogSyncDeleteResponseEnvelopeErrorsCode1003, CatalogSyncDeleteResponseEnvelopeErrorsCode1004, CatalogSyncDeleteResponseEnvelopeErrorsCode1005, CatalogSyncDeleteResponseEnvelopeErrorsCode1006, CatalogSyncDeleteResponseEnvelopeErrorsCode1007, CatalogSyncDeleteResponseEnvelopeErrorsCode1008, CatalogSyncDeleteResponseEnvelopeErrorsCode1009, CatalogSyncDeleteResponseEnvelopeErrorsCode1010, CatalogSyncDeleteResponseEnvelopeErrorsCode1011, CatalogSyncDeleteResponseEnvelopeErrorsCode1012, CatalogSyncDeleteResponseEnvelopeErrorsCode1013, CatalogSyncDeleteResponseEnvelopeErrorsCode1014, CatalogSyncDeleteResponseEnvelopeErrorsCode1015, CatalogSyncDeleteResponseEnvelopeErrorsCode1016, CatalogSyncDeleteResponseEnvelopeErrorsCode1017, CatalogSyncDeleteResponseEnvelopeErrorsCode2001, CatalogSyncDeleteResponseEnvelopeErrorsCode2002, CatalogSyncDeleteResponseEnvelopeErrorsCode2003, CatalogSyncDeleteResponseEnvelopeErrorsCode2004, CatalogSyncDeleteResponseEnvelopeErrorsCode2005, CatalogSyncDeleteResponseEnvelopeErrorsCode2006, CatalogSyncDeleteResponseEnvelopeErrorsCode2007, CatalogSyncDeleteResponseEnvelopeErrorsCode2008, CatalogSyncDeleteResponseEnvelopeErrorsCode2009, CatalogSyncDeleteResponseEnvelopeErrorsCode2010, CatalogSyncDeleteResponseEnvelopeErrorsCode2011, CatalogSyncDeleteResponseEnvelopeErrorsCode2012, CatalogSyncDeleteResponseEnvelopeErrorsCode2013, CatalogSyncDeleteResponseEnvelopeErrorsCode2014, CatalogSyncDeleteResponseEnvelopeErrorsCode2015, CatalogSyncDeleteResponseEnvelopeErrorsCode2016, CatalogSyncDeleteResponseEnvelopeErrorsCode2017, CatalogSyncDeleteResponseEnvelopeErrorsCode2018, CatalogSyncDeleteResponseEnvelopeErrorsCode2019, CatalogSyncDeleteResponseEnvelopeErrorsCode2020, CatalogSyncDeleteResponseEnvelopeErrorsCode2021, CatalogSyncDeleteResponseEnvelopeErrorsCode2022, CatalogSyncDeleteResponseEnvelopeErrorsCode3001, CatalogSyncDeleteResponseEnvelopeErrorsCode3002, CatalogSyncDeleteResponseEnvelopeErrorsCode3003, CatalogSyncDeleteResponseEnvelopeErrorsCode3004, CatalogSyncDeleteResponseEnvelopeErrorsCode3005, CatalogSyncDeleteResponseEnvelopeErrorsCode3006, CatalogSyncDeleteResponseEnvelopeErrorsCode3007, CatalogSyncDeleteResponseEnvelopeErrorsCode4001, CatalogSyncDeleteResponseEnvelopeErrorsCode4002, CatalogSyncDeleteResponseEnvelopeErrorsCode4003, CatalogSyncDeleteResponseEnvelopeErrorsCode4004, CatalogSyncDeleteResponseEnvelopeErrorsCode4005, CatalogSyncDeleteResponseEnvelopeErrorsCode4006, CatalogSyncDeleteResponseEnvelopeErrorsCode4007, CatalogSyncDeleteResponseEnvelopeErrorsCode4008, CatalogSyncDeleteResponseEnvelopeErrorsCode4009, CatalogSyncDeleteResponseEnvelopeErrorsCode4010, CatalogSyncDeleteResponseEnvelopeErrorsCode4011, CatalogSyncDeleteResponseEnvelopeErrorsCode4012, CatalogSyncDeleteResponseEnvelopeErrorsCode4013, CatalogSyncDeleteResponseEnvelopeErrorsCode4014, CatalogSyncDeleteResponseEnvelopeErrorsCode4015, CatalogSyncDeleteResponseEnvelopeErrorsCode4016, CatalogSyncDeleteResponseEnvelopeErrorsCode4017, CatalogSyncDeleteResponseEnvelopeErrorsCode4018, CatalogSyncDeleteResponseEnvelopeErrorsCode4019, CatalogSyncDeleteResponseEnvelopeErrorsCode4020, CatalogSyncDeleteResponseEnvelopeErrorsCode4021, CatalogSyncDeleteResponseEnvelopeErrorsCode4022, CatalogSyncDeleteResponseEnvelopeErrorsCode4023, CatalogSyncDeleteResponseEnvelopeErrorsCode5001, CatalogSyncDeleteResponseEnvelopeErrorsCode5002, CatalogSyncDeleteResponseEnvelopeErrorsCode5003, CatalogSyncDeleteResponseEnvelopeErrorsCode5004, CatalogSyncDeleteResponseEnvelopeErrorsCode102000, CatalogSyncDeleteResponseEnvelopeErrorsCode102001, CatalogSyncDeleteResponseEnvelopeErrorsCode102002, CatalogSyncDeleteResponseEnvelopeErrorsCode102003, CatalogSyncDeleteResponseEnvelopeErrorsCode102004, CatalogSyncDeleteResponseEnvelopeErrorsCode102005, CatalogSyncDeleteResponseEnvelopeErrorsCode102006, CatalogSyncDeleteResponseEnvelopeErrorsCode102007, CatalogSyncDeleteResponseEnvelopeErrorsCode102008, CatalogSyncDeleteResponseEnvelopeErrorsCode102009, CatalogSyncDeleteResponseEnvelopeErrorsCode102010, CatalogSyncDeleteResponseEnvelopeErrorsCode102011, CatalogSyncDeleteResponseEnvelopeErrorsCode102012, CatalogSyncDeleteResponseEnvelopeErrorsCode102013, CatalogSyncDeleteResponseEnvelopeErrorsCode102014, CatalogSyncDeleteResponseEnvelopeErrorsCode102015, CatalogSyncDeleteResponseEnvelopeErrorsCode102016, CatalogSyncDeleteResponseEnvelopeErrorsCode102017, CatalogSyncDeleteResponseEnvelopeErrorsCode102018, CatalogSyncDeleteResponseEnvelopeErrorsCode102019, CatalogSyncDeleteResponseEnvelopeErrorsCode102020, CatalogSyncDeleteResponseEnvelopeErrorsCode102021, CatalogSyncDeleteResponseEnvelopeErrorsCode102022, CatalogSyncDeleteResponseEnvelopeErrorsCode102023, CatalogSyncDeleteResponseEnvelopeErrorsCode102024, CatalogSyncDeleteResponseEnvelopeErrorsCode102025, CatalogSyncDeleteResponseEnvelopeErrorsCode102026, CatalogSyncDeleteResponseEnvelopeErrorsCode102027, CatalogSyncDeleteResponseEnvelopeErrorsCode102028, CatalogSyncDeleteResponseEnvelopeErrorsCode102029, CatalogSyncDeleteResponseEnvelopeErrorsCode102030, CatalogSyncDeleteResponseEnvelopeErrorsCode102031, CatalogSyncDeleteResponseEnvelopeErrorsCode102032, CatalogSyncDeleteResponseEnvelopeErrorsCode102033, CatalogSyncDeleteResponseEnvelopeErrorsCode102034, CatalogSyncDeleteResponseEnvelopeErrorsCode102035, CatalogSyncDeleteResponseEnvelopeErrorsCode102036, CatalogSyncDeleteResponseEnvelopeErrorsCode102037, CatalogSyncDeleteResponseEnvelopeErrorsCode102038, CatalogSyncDeleteResponseEnvelopeErrorsCode102039, CatalogSyncDeleteResponseEnvelopeErrorsCode102040, CatalogSyncDeleteResponseEnvelopeErrorsCode102041, CatalogSyncDeleteResponseEnvelopeErrorsCode102042, CatalogSyncDeleteResponseEnvelopeErrorsCode102043, CatalogSyncDeleteResponseEnvelopeErrorsCode102044, CatalogSyncDeleteResponseEnvelopeErrorsCode102045, CatalogSyncDeleteResponseEnvelopeErrorsCode102046, CatalogSyncDeleteResponseEnvelopeErrorsCode102047, CatalogSyncDeleteResponseEnvelopeErrorsCode102048, CatalogSyncDeleteResponseEnvelopeErrorsCode102049, CatalogSyncDeleteResponseEnvelopeErrorsCode102050, CatalogSyncDeleteResponseEnvelopeErrorsCode102051, CatalogSyncDeleteResponseEnvelopeErrorsCode102052, CatalogSyncDeleteResponseEnvelopeErrorsCode102053, CatalogSyncDeleteResponseEnvelopeErrorsCode102054, CatalogSyncDeleteResponseEnvelopeErrorsCode102055, CatalogSyncDeleteResponseEnvelopeErrorsCode102056, CatalogSyncDeleteResponseEnvelopeErrorsCode102057, CatalogSyncDeleteResponseEnvelopeErrorsCode102058, CatalogSyncDeleteResponseEnvelopeErrorsCode102059, CatalogSyncDeleteResponseEnvelopeErrorsCode102060, CatalogSyncDeleteResponseEnvelopeErrorsCode102061, CatalogSyncDeleteResponseEnvelopeErrorsCode102062, CatalogSyncDeleteResponseEnvelopeErrorsCode102063, CatalogSyncDeleteResponseEnvelopeErrorsCode102064, CatalogSyncDeleteResponseEnvelopeErrorsCode102065, CatalogSyncDeleteResponseEnvelopeErrorsCode102066, CatalogSyncDeleteResponseEnvelopeErrorsCode103001, CatalogSyncDeleteResponseEnvelopeErrorsCode103002, CatalogSyncDeleteResponseEnvelopeErrorsCode103003, CatalogSyncDeleteResponseEnvelopeErrorsCode103004, CatalogSyncDeleteResponseEnvelopeErrorsCode103005, CatalogSyncDeleteResponseEnvelopeErrorsCode103006, CatalogSyncDeleteResponseEnvelopeErrorsCode103007, CatalogSyncDeleteResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncDeleteResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                          `json:"l10n_key"`
	LoggableError string                                          `json:"loggable_error"`
	TemplateData  interface{}                                     `json:"template_data"`
	TraceID       string                                          `json:"trace_id"`
	JSON          catalogSyncDeleteResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeErrorsMetaJSON contains the JSON metadata for
// the struct [CatalogSyncDeleteResponseEnvelopeErrorsMeta]
type catalogSyncDeleteResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeErrorsSource struct {
	Parameter           string                                            `json:"parameter"`
	ParameterValueIndex int64                                             `json:"parameter_value_index"`
	Pointer             string                                            `json:"pointer"`
	JSON                catalogSyncDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CatalogSyncDeleteResponseEnvelopeErrorsSource]
type catalogSyncDeleteResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeMessages struct {
	Code             CatalogSyncDeleteResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Meta             CatalogSyncDeleteResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CatalogSyncDeleteResponseEnvelopeMessages]
type catalogSyncDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeMessagesCode int64

const (
	CatalogSyncDeleteResponseEnvelopeMessagesCode1001   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1001
	CatalogSyncDeleteResponseEnvelopeMessagesCode1002   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1002
	CatalogSyncDeleteResponseEnvelopeMessagesCode1003   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1003
	CatalogSyncDeleteResponseEnvelopeMessagesCode1004   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1004
	CatalogSyncDeleteResponseEnvelopeMessagesCode1005   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1005
	CatalogSyncDeleteResponseEnvelopeMessagesCode1006   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1006
	CatalogSyncDeleteResponseEnvelopeMessagesCode1007   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1007
	CatalogSyncDeleteResponseEnvelopeMessagesCode1008   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1008
	CatalogSyncDeleteResponseEnvelopeMessagesCode1009   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1009
	CatalogSyncDeleteResponseEnvelopeMessagesCode1010   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1010
	CatalogSyncDeleteResponseEnvelopeMessagesCode1011   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1011
	CatalogSyncDeleteResponseEnvelopeMessagesCode1012   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1012
	CatalogSyncDeleteResponseEnvelopeMessagesCode1013   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1013
	CatalogSyncDeleteResponseEnvelopeMessagesCode1014   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1014
	CatalogSyncDeleteResponseEnvelopeMessagesCode1015   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1015
	CatalogSyncDeleteResponseEnvelopeMessagesCode1016   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1016
	CatalogSyncDeleteResponseEnvelopeMessagesCode1017   CatalogSyncDeleteResponseEnvelopeMessagesCode = 1017
	CatalogSyncDeleteResponseEnvelopeMessagesCode2001   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2001
	CatalogSyncDeleteResponseEnvelopeMessagesCode2002   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2002
	CatalogSyncDeleteResponseEnvelopeMessagesCode2003   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2003
	CatalogSyncDeleteResponseEnvelopeMessagesCode2004   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2004
	CatalogSyncDeleteResponseEnvelopeMessagesCode2005   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2005
	CatalogSyncDeleteResponseEnvelopeMessagesCode2006   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2006
	CatalogSyncDeleteResponseEnvelopeMessagesCode2007   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2007
	CatalogSyncDeleteResponseEnvelopeMessagesCode2008   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2008
	CatalogSyncDeleteResponseEnvelopeMessagesCode2009   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2009
	CatalogSyncDeleteResponseEnvelopeMessagesCode2010   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2010
	CatalogSyncDeleteResponseEnvelopeMessagesCode2011   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2011
	CatalogSyncDeleteResponseEnvelopeMessagesCode2012   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2012
	CatalogSyncDeleteResponseEnvelopeMessagesCode2013   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2013
	CatalogSyncDeleteResponseEnvelopeMessagesCode2014   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2014
	CatalogSyncDeleteResponseEnvelopeMessagesCode2015   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2015
	CatalogSyncDeleteResponseEnvelopeMessagesCode2016   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2016
	CatalogSyncDeleteResponseEnvelopeMessagesCode2017   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2017
	CatalogSyncDeleteResponseEnvelopeMessagesCode2018   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2018
	CatalogSyncDeleteResponseEnvelopeMessagesCode2019   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2019
	CatalogSyncDeleteResponseEnvelopeMessagesCode2020   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2020
	CatalogSyncDeleteResponseEnvelopeMessagesCode2021   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2021
	CatalogSyncDeleteResponseEnvelopeMessagesCode2022   CatalogSyncDeleteResponseEnvelopeMessagesCode = 2022
	CatalogSyncDeleteResponseEnvelopeMessagesCode3001   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3001
	CatalogSyncDeleteResponseEnvelopeMessagesCode3002   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3002
	CatalogSyncDeleteResponseEnvelopeMessagesCode3003   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3003
	CatalogSyncDeleteResponseEnvelopeMessagesCode3004   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3004
	CatalogSyncDeleteResponseEnvelopeMessagesCode3005   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3005
	CatalogSyncDeleteResponseEnvelopeMessagesCode3006   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3006
	CatalogSyncDeleteResponseEnvelopeMessagesCode3007   CatalogSyncDeleteResponseEnvelopeMessagesCode = 3007
	CatalogSyncDeleteResponseEnvelopeMessagesCode4001   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4001
	CatalogSyncDeleteResponseEnvelopeMessagesCode4002   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4002
	CatalogSyncDeleteResponseEnvelopeMessagesCode4003   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4003
	CatalogSyncDeleteResponseEnvelopeMessagesCode4004   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4004
	CatalogSyncDeleteResponseEnvelopeMessagesCode4005   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4005
	CatalogSyncDeleteResponseEnvelopeMessagesCode4006   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4006
	CatalogSyncDeleteResponseEnvelopeMessagesCode4007   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4007
	CatalogSyncDeleteResponseEnvelopeMessagesCode4008   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4008
	CatalogSyncDeleteResponseEnvelopeMessagesCode4009   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4009
	CatalogSyncDeleteResponseEnvelopeMessagesCode4010   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4010
	CatalogSyncDeleteResponseEnvelopeMessagesCode4011   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4011
	CatalogSyncDeleteResponseEnvelopeMessagesCode4012   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4012
	CatalogSyncDeleteResponseEnvelopeMessagesCode4013   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4013
	CatalogSyncDeleteResponseEnvelopeMessagesCode4014   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4014
	CatalogSyncDeleteResponseEnvelopeMessagesCode4015   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4015
	CatalogSyncDeleteResponseEnvelopeMessagesCode4016   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4016
	CatalogSyncDeleteResponseEnvelopeMessagesCode4017   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4017
	CatalogSyncDeleteResponseEnvelopeMessagesCode4018   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4018
	CatalogSyncDeleteResponseEnvelopeMessagesCode4019   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4019
	CatalogSyncDeleteResponseEnvelopeMessagesCode4020   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4020
	CatalogSyncDeleteResponseEnvelopeMessagesCode4021   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4021
	CatalogSyncDeleteResponseEnvelopeMessagesCode4022   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4022
	CatalogSyncDeleteResponseEnvelopeMessagesCode4023   CatalogSyncDeleteResponseEnvelopeMessagesCode = 4023
	CatalogSyncDeleteResponseEnvelopeMessagesCode5001   CatalogSyncDeleteResponseEnvelopeMessagesCode = 5001
	CatalogSyncDeleteResponseEnvelopeMessagesCode5002   CatalogSyncDeleteResponseEnvelopeMessagesCode = 5002
	CatalogSyncDeleteResponseEnvelopeMessagesCode5003   CatalogSyncDeleteResponseEnvelopeMessagesCode = 5003
	CatalogSyncDeleteResponseEnvelopeMessagesCode5004   CatalogSyncDeleteResponseEnvelopeMessagesCode = 5004
	CatalogSyncDeleteResponseEnvelopeMessagesCode102000 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102000
	CatalogSyncDeleteResponseEnvelopeMessagesCode102001 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102001
	CatalogSyncDeleteResponseEnvelopeMessagesCode102002 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102002
	CatalogSyncDeleteResponseEnvelopeMessagesCode102003 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102003
	CatalogSyncDeleteResponseEnvelopeMessagesCode102004 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102004
	CatalogSyncDeleteResponseEnvelopeMessagesCode102005 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102005
	CatalogSyncDeleteResponseEnvelopeMessagesCode102006 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102006
	CatalogSyncDeleteResponseEnvelopeMessagesCode102007 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102007
	CatalogSyncDeleteResponseEnvelopeMessagesCode102008 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102008
	CatalogSyncDeleteResponseEnvelopeMessagesCode102009 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102009
	CatalogSyncDeleteResponseEnvelopeMessagesCode102010 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102010
	CatalogSyncDeleteResponseEnvelopeMessagesCode102011 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102011
	CatalogSyncDeleteResponseEnvelopeMessagesCode102012 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102012
	CatalogSyncDeleteResponseEnvelopeMessagesCode102013 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102013
	CatalogSyncDeleteResponseEnvelopeMessagesCode102014 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102014
	CatalogSyncDeleteResponseEnvelopeMessagesCode102015 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102015
	CatalogSyncDeleteResponseEnvelopeMessagesCode102016 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102016
	CatalogSyncDeleteResponseEnvelopeMessagesCode102017 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102017
	CatalogSyncDeleteResponseEnvelopeMessagesCode102018 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102018
	CatalogSyncDeleteResponseEnvelopeMessagesCode102019 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102019
	CatalogSyncDeleteResponseEnvelopeMessagesCode102020 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102020
	CatalogSyncDeleteResponseEnvelopeMessagesCode102021 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102021
	CatalogSyncDeleteResponseEnvelopeMessagesCode102022 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102022
	CatalogSyncDeleteResponseEnvelopeMessagesCode102023 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102023
	CatalogSyncDeleteResponseEnvelopeMessagesCode102024 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102024
	CatalogSyncDeleteResponseEnvelopeMessagesCode102025 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102025
	CatalogSyncDeleteResponseEnvelopeMessagesCode102026 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102026
	CatalogSyncDeleteResponseEnvelopeMessagesCode102027 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102027
	CatalogSyncDeleteResponseEnvelopeMessagesCode102028 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102028
	CatalogSyncDeleteResponseEnvelopeMessagesCode102029 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102029
	CatalogSyncDeleteResponseEnvelopeMessagesCode102030 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102030
	CatalogSyncDeleteResponseEnvelopeMessagesCode102031 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102031
	CatalogSyncDeleteResponseEnvelopeMessagesCode102032 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102032
	CatalogSyncDeleteResponseEnvelopeMessagesCode102033 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102033
	CatalogSyncDeleteResponseEnvelopeMessagesCode102034 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102034
	CatalogSyncDeleteResponseEnvelopeMessagesCode102035 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102035
	CatalogSyncDeleteResponseEnvelopeMessagesCode102036 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102036
	CatalogSyncDeleteResponseEnvelopeMessagesCode102037 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102037
	CatalogSyncDeleteResponseEnvelopeMessagesCode102038 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102038
	CatalogSyncDeleteResponseEnvelopeMessagesCode102039 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102039
	CatalogSyncDeleteResponseEnvelopeMessagesCode102040 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102040
	CatalogSyncDeleteResponseEnvelopeMessagesCode102041 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102041
	CatalogSyncDeleteResponseEnvelopeMessagesCode102042 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102042
	CatalogSyncDeleteResponseEnvelopeMessagesCode102043 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102043
	CatalogSyncDeleteResponseEnvelopeMessagesCode102044 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102044
	CatalogSyncDeleteResponseEnvelopeMessagesCode102045 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102045
	CatalogSyncDeleteResponseEnvelopeMessagesCode102046 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102046
	CatalogSyncDeleteResponseEnvelopeMessagesCode102047 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102047
	CatalogSyncDeleteResponseEnvelopeMessagesCode102048 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102048
	CatalogSyncDeleteResponseEnvelopeMessagesCode102049 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102049
	CatalogSyncDeleteResponseEnvelopeMessagesCode102050 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102050
	CatalogSyncDeleteResponseEnvelopeMessagesCode102051 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102051
	CatalogSyncDeleteResponseEnvelopeMessagesCode102052 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102052
	CatalogSyncDeleteResponseEnvelopeMessagesCode102053 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102053
	CatalogSyncDeleteResponseEnvelopeMessagesCode102054 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102054
	CatalogSyncDeleteResponseEnvelopeMessagesCode102055 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102055
	CatalogSyncDeleteResponseEnvelopeMessagesCode102056 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102056
	CatalogSyncDeleteResponseEnvelopeMessagesCode102057 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102057
	CatalogSyncDeleteResponseEnvelopeMessagesCode102058 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102058
	CatalogSyncDeleteResponseEnvelopeMessagesCode102059 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102059
	CatalogSyncDeleteResponseEnvelopeMessagesCode102060 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102060
	CatalogSyncDeleteResponseEnvelopeMessagesCode102061 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102061
	CatalogSyncDeleteResponseEnvelopeMessagesCode102062 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102062
	CatalogSyncDeleteResponseEnvelopeMessagesCode102063 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102063
	CatalogSyncDeleteResponseEnvelopeMessagesCode102064 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102064
	CatalogSyncDeleteResponseEnvelopeMessagesCode102065 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102065
	CatalogSyncDeleteResponseEnvelopeMessagesCode102066 CatalogSyncDeleteResponseEnvelopeMessagesCode = 102066
	CatalogSyncDeleteResponseEnvelopeMessagesCode103001 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103001
	CatalogSyncDeleteResponseEnvelopeMessagesCode103002 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103002
	CatalogSyncDeleteResponseEnvelopeMessagesCode103003 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103003
	CatalogSyncDeleteResponseEnvelopeMessagesCode103004 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103004
	CatalogSyncDeleteResponseEnvelopeMessagesCode103005 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103005
	CatalogSyncDeleteResponseEnvelopeMessagesCode103006 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103006
	CatalogSyncDeleteResponseEnvelopeMessagesCode103007 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103007
	CatalogSyncDeleteResponseEnvelopeMessagesCode103008 CatalogSyncDeleteResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncDeleteResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncDeleteResponseEnvelopeMessagesCode1001, CatalogSyncDeleteResponseEnvelopeMessagesCode1002, CatalogSyncDeleteResponseEnvelopeMessagesCode1003, CatalogSyncDeleteResponseEnvelopeMessagesCode1004, CatalogSyncDeleteResponseEnvelopeMessagesCode1005, CatalogSyncDeleteResponseEnvelopeMessagesCode1006, CatalogSyncDeleteResponseEnvelopeMessagesCode1007, CatalogSyncDeleteResponseEnvelopeMessagesCode1008, CatalogSyncDeleteResponseEnvelopeMessagesCode1009, CatalogSyncDeleteResponseEnvelopeMessagesCode1010, CatalogSyncDeleteResponseEnvelopeMessagesCode1011, CatalogSyncDeleteResponseEnvelopeMessagesCode1012, CatalogSyncDeleteResponseEnvelopeMessagesCode1013, CatalogSyncDeleteResponseEnvelopeMessagesCode1014, CatalogSyncDeleteResponseEnvelopeMessagesCode1015, CatalogSyncDeleteResponseEnvelopeMessagesCode1016, CatalogSyncDeleteResponseEnvelopeMessagesCode1017, CatalogSyncDeleteResponseEnvelopeMessagesCode2001, CatalogSyncDeleteResponseEnvelopeMessagesCode2002, CatalogSyncDeleteResponseEnvelopeMessagesCode2003, CatalogSyncDeleteResponseEnvelopeMessagesCode2004, CatalogSyncDeleteResponseEnvelopeMessagesCode2005, CatalogSyncDeleteResponseEnvelopeMessagesCode2006, CatalogSyncDeleteResponseEnvelopeMessagesCode2007, CatalogSyncDeleteResponseEnvelopeMessagesCode2008, CatalogSyncDeleteResponseEnvelopeMessagesCode2009, CatalogSyncDeleteResponseEnvelopeMessagesCode2010, CatalogSyncDeleteResponseEnvelopeMessagesCode2011, CatalogSyncDeleteResponseEnvelopeMessagesCode2012, CatalogSyncDeleteResponseEnvelopeMessagesCode2013, CatalogSyncDeleteResponseEnvelopeMessagesCode2014, CatalogSyncDeleteResponseEnvelopeMessagesCode2015, CatalogSyncDeleteResponseEnvelopeMessagesCode2016, CatalogSyncDeleteResponseEnvelopeMessagesCode2017, CatalogSyncDeleteResponseEnvelopeMessagesCode2018, CatalogSyncDeleteResponseEnvelopeMessagesCode2019, CatalogSyncDeleteResponseEnvelopeMessagesCode2020, CatalogSyncDeleteResponseEnvelopeMessagesCode2021, CatalogSyncDeleteResponseEnvelopeMessagesCode2022, CatalogSyncDeleteResponseEnvelopeMessagesCode3001, CatalogSyncDeleteResponseEnvelopeMessagesCode3002, CatalogSyncDeleteResponseEnvelopeMessagesCode3003, CatalogSyncDeleteResponseEnvelopeMessagesCode3004, CatalogSyncDeleteResponseEnvelopeMessagesCode3005, CatalogSyncDeleteResponseEnvelopeMessagesCode3006, CatalogSyncDeleteResponseEnvelopeMessagesCode3007, CatalogSyncDeleteResponseEnvelopeMessagesCode4001, CatalogSyncDeleteResponseEnvelopeMessagesCode4002, CatalogSyncDeleteResponseEnvelopeMessagesCode4003, CatalogSyncDeleteResponseEnvelopeMessagesCode4004, CatalogSyncDeleteResponseEnvelopeMessagesCode4005, CatalogSyncDeleteResponseEnvelopeMessagesCode4006, CatalogSyncDeleteResponseEnvelopeMessagesCode4007, CatalogSyncDeleteResponseEnvelopeMessagesCode4008, CatalogSyncDeleteResponseEnvelopeMessagesCode4009, CatalogSyncDeleteResponseEnvelopeMessagesCode4010, CatalogSyncDeleteResponseEnvelopeMessagesCode4011, CatalogSyncDeleteResponseEnvelopeMessagesCode4012, CatalogSyncDeleteResponseEnvelopeMessagesCode4013, CatalogSyncDeleteResponseEnvelopeMessagesCode4014, CatalogSyncDeleteResponseEnvelopeMessagesCode4015, CatalogSyncDeleteResponseEnvelopeMessagesCode4016, CatalogSyncDeleteResponseEnvelopeMessagesCode4017, CatalogSyncDeleteResponseEnvelopeMessagesCode4018, CatalogSyncDeleteResponseEnvelopeMessagesCode4019, CatalogSyncDeleteResponseEnvelopeMessagesCode4020, CatalogSyncDeleteResponseEnvelopeMessagesCode4021, CatalogSyncDeleteResponseEnvelopeMessagesCode4022, CatalogSyncDeleteResponseEnvelopeMessagesCode4023, CatalogSyncDeleteResponseEnvelopeMessagesCode5001, CatalogSyncDeleteResponseEnvelopeMessagesCode5002, CatalogSyncDeleteResponseEnvelopeMessagesCode5003, CatalogSyncDeleteResponseEnvelopeMessagesCode5004, CatalogSyncDeleteResponseEnvelopeMessagesCode102000, CatalogSyncDeleteResponseEnvelopeMessagesCode102001, CatalogSyncDeleteResponseEnvelopeMessagesCode102002, CatalogSyncDeleteResponseEnvelopeMessagesCode102003, CatalogSyncDeleteResponseEnvelopeMessagesCode102004, CatalogSyncDeleteResponseEnvelopeMessagesCode102005, CatalogSyncDeleteResponseEnvelopeMessagesCode102006, CatalogSyncDeleteResponseEnvelopeMessagesCode102007, CatalogSyncDeleteResponseEnvelopeMessagesCode102008, CatalogSyncDeleteResponseEnvelopeMessagesCode102009, CatalogSyncDeleteResponseEnvelopeMessagesCode102010, CatalogSyncDeleteResponseEnvelopeMessagesCode102011, CatalogSyncDeleteResponseEnvelopeMessagesCode102012, CatalogSyncDeleteResponseEnvelopeMessagesCode102013, CatalogSyncDeleteResponseEnvelopeMessagesCode102014, CatalogSyncDeleteResponseEnvelopeMessagesCode102015, CatalogSyncDeleteResponseEnvelopeMessagesCode102016, CatalogSyncDeleteResponseEnvelopeMessagesCode102017, CatalogSyncDeleteResponseEnvelopeMessagesCode102018, CatalogSyncDeleteResponseEnvelopeMessagesCode102019, CatalogSyncDeleteResponseEnvelopeMessagesCode102020, CatalogSyncDeleteResponseEnvelopeMessagesCode102021, CatalogSyncDeleteResponseEnvelopeMessagesCode102022, CatalogSyncDeleteResponseEnvelopeMessagesCode102023, CatalogSyncDeleteResponseEnvelopeMessagesCode102024, CatalogSyncDeleteResponseEnvelopeMessagesCode102025, CatalogSyncDeleteResponseEnvelopeMessagesCode102026, CatalogSyncDeleteResponseEnvelopeMessagesCode102027, CatalogSyncDeleteResponseEnvelopeMessagesCode102028, CatalogSyncDeleteResponseEnvelopeMessagesCode102029, CatalogSyncDeleteResponseEnvelopeMessagesCode102030, CatalogSyncDeleteResponseEnvelopeMessagesCode102031, CatalogSyncDeleteResponseEnvelopeMessagesCode102032, CatalogSyncDeleteResponseEnvelopeMessagesCode102033, CatalogSyncDeleteResponseEnvelopeMessagesCode102034, CatalogSyncDeleteResponseEnvelopeMessagesCode102035, CatalogSyncDeleteResponseEnvelopeMessagesCode102036, CatalogSyncDeleteResponseEnvelopeMessagesCode102037, CatalogSyncDeleteResponseEnvelopeMessagesCode102038, CatalogSyncDeleteResponseEnvelopeMessagesCode102039, CatalogSyncDeleteResponseEnvelopeMessagesCode102040, CatalogSyncDeleteResponseEnvelopeMessagesCode102041, CatalogSyncDeleteResponseEnvelopeMessagesCode102042, CatalogSyncDeleteResponseEnvelopeMessagesCode102043, CatalogSyncDeleteResponseEnvelopeMessagesCode102044, CatalogSyncDeleteResponseEnvelopeMessagesCode102045, CatalogSyncDeleteResponseEnvelopeMessagesCode102046, CatalogSyncDeleteResponseEnvelopeMessagesCode102047, CatalogSyncDeleteResponseEnvelopeMessagesCode102048, CatalogSyncDeleteResponseEnvelopeMessagesCode102049, CatalogSyncDeleteResponseEnvelopeMessagesCode102050, CatalogSyncDeleteResponseEnvelopeMessagesCode102051, CatalogSyncDeleteResponseEnvelopeMessagesCode102052, CatalogSyncDeleteResponseEnvelopeMessagesCode102053, CatalogSyncDeleteResponseEnvelopeMessagesCode102054, CatalogSyncDeleteResponseEnvelopeMessagesCode102055, CatalogSyncDeleteResponseEnvelopeMessagesCode102056, CatalogSyncDeleteResponseEnvelopeMessagesCode102057, CatalogSyncDeleteResponseEnvelopeMessagesCode102058, CatalogSyncDeleteResponseEnvelopeMessagesCode102059, CatalogSyncDeleteResponseEnvelopeMessagesCode102060, CatalogSyncDeleteResponseEnvelopeMessagesCode102061, CatalogSyncDeleteResponseEnvelopeMessagesCode102062, CatalogSyncDeleteResponseEnvelopeMessagesCode102063, CatalogSyncDeleteResponseEnvelopeMessagesCode102064, CatalogSyncDeleteResponseEnvelopeMessagesCode102065, CatalogSyncDeleteResponseEnvelopeMessagesCode102066, CatalogSyncDeleteResponseEnvelopeMessagesCode103001, CatalogSyncDeleteResponseEnvelopeMessagesCode103002, CatalogSyncDeleteResponseEnvelopeMessagesCode103003, CatalogSyncDeleteResponseEnvelopeMessagesCode103004, CatalogSyncDeleteResponseEnvelopeMessagesCode103005, CatalogSyncDeleteResponseEnvelopeMessagesCode103006, CatalogSyncDeleteResponseEnvelopeMessagesCode103007, CatalogSyncDeleteResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncDeleteResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                            `json:"l10n_key"`
	LoggableError string                                            `json:"loggable_error"`
	TemplateData  interface{}                                       `json:"template_data"`
	TraceID       string                                            `json:"trace_id"`
	JSON          catalogSyncDeleteResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeMessagesMetaJSON contains the JSON metadata for
// the struct [CatalogSyncDeleteResponseEnvelopeMessagesMeta]
type catalogSyncDeleteResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncDeleteResponseEnvelopeMessagesSource struct {
	Parameter           string                                              `json:"parameter"`
	ParameterValueIndex int64                                               `json:"parameter_value_index"`
	Pointer             string                                              `json:"pointer"`
	JSON                catalogSyncDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CatalogSyncDeleteResponseEnvelopeMessagesSource]
type catalogSyncDeleteResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditParams struct {
	AccountID   param.Field[string]                          `path:"account_id,required"`
	Description param.Field[string]                          `json:"description"`
	Name        param.Field[string]                          `json:"name"`
	Policy      param.Field[string]                          `json:"policy"`
	UpdateMode  param.Field[CatalogSyncEditParamsUpdateMode] `json:"update_mode"`
}

func (r CatalogSyncEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CatalogSyncEditParamsUpdateMode string

const (
	CatalogSyncEditParamsUpdateModeAuto   CatalogSyncEditParamsUpdateMode = "AUTO"
	CatalogSyncEditParamsUpdateModeManual CatalogSyncEditParamsUpdateMode = "MANUAL"
)

func (r CatalogSyncEditParamsUpdateMode) IsKnown() bool {
	switch r {
	case CatalogSyncEditParamsUpdateModeAuto, CatalogSyncEditParamsUpdateModeManual:
		return true
	}
	return false
}

type CatalogSyncEditResponseEnvelope struct {
	Errors   []CatalogSyncEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CatalogSyncEditResponse                   `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     catalogSyncEditResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncEditResponseEnvelope]
type catalogSyncEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeErrors struct {
	Code             CatalogSyncEditResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Meta             CatalogSyncEditResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncEditResponseEnvelopeErrors]
type catalogSyncEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeErrorsCode int64

const (
	CatalogSyncEditResponseEnvelopeErrorsCode1001   CatalogSyncEditResponseEnvelopeErrorsCode = 1001
	CatalogSyncEditResponseEnvelopeErrorsCode1002   CatalogSyncEditResponseEnvelopeErrorsCode = 1002
	CatalogSyncEditResponseEnvelopeErrorsCode1003   CatalogSyncEditResponseEnvelopeErrorsCode = 1003
	CatalogSyncEditResponseEnvelopeErrorsCode1004   CatalogSyncEditResponseEnvelopeErrorsCode = 1004
	CatalogSyncEditResponseEnvelopeErrorsCode1005   CatalogSyncEditResponseEnvelopeErrorsCode = 1005
	CatalogSyncEditResponseEnvelopeErrorsCode1006   CatalogSyncEditResponseEnvelopeErrorsCode = 1006
	CatalogSyncEditResponseEnvelopeErrorsCode1007   CatalogSyncEditResponseEnvelopeErrorsCode = 1007
	CatalogSyncEditResponseEnvelopeErrorsCode1008   CatalogSyncEditResponseEnvelopeErrorsCode = 1008
	CatalogSyncEditResponseEnvelopeErrorsCode1009   CatalogSyncEditResponseEnvelopeErrorsCode = 1009
	CatalogSyncEditResponseEnvelopeErrorsCode1010   CatalogSyncEditResponseEnvelopeErrorsCode = 1010
	CatalogSyncEditResponseEnvelopeErrorsCode1011   CatalogSyncEditResponseEnvelopeErrorsCode = 1011
	CatalogSyncEditResponseEnvelopeErrorsCode1012   CatalogSyncEditResponseEnvelopeErrorsCode = 1012
	CatalogSyncEditResponseEnvelopeErrorsCode1013   CatalogSyncEditResponseEnvelopeErrorsCode = 1013
	CatalogSyncEditResponseEnvelopeErrorsCode1014   CatalogSyncEditResponseEnvelopeErrorsCode = 1014
	CatalogSyncEditResponseEnvelopeErrorsCode1015   CatalogSyncEditResponseEnvelopeErrorsCode = 1015
	CatalogSyncEditResponseEnvelopeErrorsCode1016   CatalogSyncEditResponseEnvelopeErrorsCode = 1016
	CatalogSyncEditResponseEnvelopeErrorsCode1017   CatalogSyncEditResponseEnvelopeErrorsCode = 1017
	CatalogSyncEditResponseEnvelopeErrorsCode2001   CatalogSyncEditResponseEnvelopeErrorsCode = 2001
	CatalogSyncEditResponseEnvelopeErrorsCode2002   CatalogSyncEditResponseEnvelopeErrorsCode = 2002
	CatalogSyncEditResponseEnvelopeErrorsCode2003   CatalogSyncEditResponseEnvelopeErrorsCode = 2003
	CatalogSyncEditResponseEnvelopeErrorsCode2004   CatalogSyncEditResponseEnvelopeErrorsCode = 2004
	CatalogSyncEditResponseEnvelopeErrorsCode2005   CatalogSyncEditResponseEnvelopeErrorsCode = 2005
	CatalogSyncEditResponseEnvelopeErrorsCode2006   CatalogSyncEditResponseEnvelopeErrorsCode = 2006
	CatalogSyncEditResponseEnvelopeErrorsCode2007   CatalogSyncEditResponseEnvelopeErrorsCode = 2007
	CatalogSyncEditResponseEnvelopeErrorsCode2008   CatalogSyncEditResponseEnvelopeErrorsCode = 2008
	CatalogSyncEditResponseEnvelopeErrorsCode2009   CatalogSyncEditResponseEnvelopeErrorsCode = 2009
	CatalogSyncEditResponseEnvelopeErrorsCode2010   CatalogSyncEditResponseEnvelopeErrorsCode = 2010
	CatalogSyncEditResponseEnvelopeErrorsCode2011   CatalogSyncEditResponseEnvelopeErrorsCode = 2011
	CatalogSyncEditResponseEnvelopeErrorsCode2012   CatalogSyncEditResponseEnvelopeErrorsCode = 2012
	CatalogSyncEditResponseEnvelopeErrorsCode2013   CatalogSyncEditResponseEnvelopeErrorsCode = 2013
	CatalogSyncEditResponseEnvelopeErrorsCode2014   CatalogSyncEditResponseEnvelopeErrorsCode = 2014
	CatalogSyncEditResponseEnvelopeErrorsCode2015   CatalogSyncEditResponseEnvelopeErrorsCode = 2015
	CatalogSyncEditResponseEnvelopeErrorsCode2016   CatalogSyncEditResponseEnvelopeErrorsCode = 2016
	CatalogSyncEditResponseEnvelopeErrorsCode2017   CatalogSyncEditResponseEnvelopeErrorsCode = 2017
	CatalogSyncEditResponseEnvelopeErrorsCode2018   CatalogSyncEditResponseEnvelopeErrorsCode = 2018
	CatalogSyncEditResponseEnvelopeErrorsCode2019   CatalogSyncEditResponseEnvelopeErrorsCode = 2019
	CatalogSyncEditResponseEnvelopeErrorsCode2020   CatalogSyncEditResponseEnvelopeErrorsCode = 2020
	CatalogSyncEditResponseEnvelopeErrorsCode2021   CatalogSyncEditResponseEnvelopeErrorsCode = 2021
	CatalogSyncEditResponseEnvelopeErrorsCode2022   CatalogSyncEditResponseEnvelopeErrorsCode = 2022
	CatalogSyncEditResponseEnvelopeErrorsCode3001   CatalogSyncEditResponseEnvelopeErrorsCode = 3001
	CatalogSyncEditResponseEnvelopeErrorsCode3002   CatalogSyncEditResponseEnvelopeErrorsCode = 3002
	CatalogSyncEditResponseEnvelopeErrorsCode3003   CatalogSyncEditResponseEnvelopeErrorsCode = 3003
	CatalogSyncEditResponseEnvelopeErrorsCode3004   CatalogSyncEditResponseEnvelopeErrorsCode = 3004
	CatalogSyncEditResponseEnvelopeErrorsCode3005   CatalogSyncEditResponseEnvelopeErrorsCode = 3005
	CatalogSyncEditResponseEnvelopeErrorsCode3006   CatalogSyncEditResponseEnvelopeErrorsCode = 3006
	CatalogSyncEditResponseEnvelopeErrorsCode3007   CatalogSyncEditResponseEnvelopeErrorsCode = 3007
	CatalogSyncEditResponseEnvelopeErrorsCode4001   CatalogSyncEditResponseEnvelopeErrorsCode = 4001
	CatalogSyncEditResponseEnvelopeErrorsCode4002   CatalogSyncEditResponseEnvelopeErrorsCode = 4002
	CatalogSyncEditResponseEnvelopeErrorsCode4003   CatalogSyncEditResponseEnvelopeErrorsCode = 4003
	CatalogSyncEditResponseEnvelopeErrorsCode4004   CatalogSyncEditResponseEnvelopeErrorsCode = 4004
	CatalogSyncEditResponseEnvelopeErrorsCode4005   CatalogSyncEditResponseEnvelopeErrorsCode = 4005
	CatalogSyncEditResponseEnvelopeErrorsCode4006   CatalogSyncEditResponseEnvelopeErrorsCode = 4006
	CatalogSyncEditResponseEnvelopeErrorsCode4007   CatalogSyncEditResponseEnvelopeErrorsCode = 4007
	CatalogSyncEditResponseEnvelopeErrorsCode4008   CatalogSyncEditResponseEnvelopeErrorsCode = 4008
	CatalogSyncEditResponseEnvelopeErrorsCode4009   CatalogSyncEditResponseEnvelopeErrorsCode = 4009
	CatalogSyncEditResponseEnvelopeErrorsCode4010   CatalogSyncEditResponseEnvelopeErrorsCode = 4010
	CatalogSyncEditResponseEnvelopeErrorsCode4011   CatalogSyncEditResponseEnvelopeErrorsCode = 4011
	CatalogSyncEditResponseEnvelopeErrorsCode4012   CatalogSyncEditResponseEnvelopeErrorsCode = 4012
	CatalogSyncEditResponseEnvelopeErrorsCode4013   CatalogSyncEditResponseEnvelopeErrorsCode = 4013
	CatalogSyncEditResponseEnvelopeErrorsCode4014   CatalogSyncEditResponseEnvelopeErrorsCode = 4014
	CatalogSyncEditResponseEnvelopeErrorsCode4015   CatalogSyncEditResponseEnvelopeErrorsCode = 4015
	CatalogSyncEditResponseEnvelopeErrorsCode4016   CatalogSyncEditResponseEnvelopeErrorsCode = 4016
	CatalogSyncEditResponseEnvelopeErrorsCode4017   CatalogSyncEditResponseEnvelopeErrorsCode = 4017
	CatalogSyncEditResponseEnvelopeErrorsCode4018   CatalogSyncEditResponseEnvelopeErrorsCode = 4018
	CatalogSyncEditResponseEnvelopeErrorsCode4019   CatalogSyncEditResponseEnvelopeErrorsCode = 4019
	CatalogSyncEditResponseEnvelopeErrorsCode4020   CatalogSyncEditResponseEnvelopeErrorsCode = 4020
	CatalogSyncEditResponseEnvelopeErrorsCode4021   CatalogSyncEditResponseEnvelopeErrorsCode = 4021
	CatalogSyncEditResponseEnvelopeErrorsCode4022   CatalogSyncEditResponseEnvelopeErrorsCode = 4022
	CatalogSyncEditResponseEnvelopeErrorsCode4023   CatalogSyncEditResponseEnvelopeErrorsCode = 4023
	CatalogSyncEditResponseEnvelopeErrorsCode5001   CatalogSyncEditResponseEnvelopeErrorsCode = 5001
	CatalogSyncEditResponseEnvelopeErrorsCode5002   CatalogSyncEditResponseEnvelopeErrorsCode = 5002
	CatalogSyncEditResponseEnvelopeErrorsCode5003   CatalogSyncEditResponseEnvelopeErrorsCode = 5003
	CatalogSyncEditResponseEnvelopeErrorsCode5004   CatalogSyncEditResponseEnvelopeErrorsCode = 5004
	CatalogSyncEditResponseEnvelopeErrorsCode102000 CatalogSyncEditResponseEnvelopeErrorsCode = 102000
	CatalogSyncEditResponseEnvelopeErrorsCode102001 CatalogSyncEditResponseEnvelopeErrorsCode = 102001
	CatalogSyncEditResponseEnvelopeErrorsCode102002 CatalogSyncEditResponseEnvelopeErrorsCode = 102002
	CatalogSyncEditResponseEnvelopeErrorsCode102003 CatalogSyncEditResponseEnvelopeErrorsCode = 102003
	CatalogSyncEditResponseEnvelopeErrorsCode102004 CatalogSyncEditResponseEnvelopeErrorsCode = 102004
	CatalogSyncEditResponseEnvelopeErrorsCode102005 CatalogSyncEditResponseEnvelopeErrorsCode = 102005
	CatalogSyncEditResponseEnvelopeErrorsCode102006 CatalogSyncEditResponseEnvelopeErrorsCode = 102006
	CatalogSyncEditResponseEnvelopeErrorsCode102007 CatalogSyncEditResponseEnvelopeErrorsCode = 102007
	CatalogSyncEditResponseEnvelopeErrorsCode102008 CatalogSyncEditResponseEnvelopeErrorsCode = 102008
	CatalogSyncEditResponseEnvelopeErrorsCode102009 CatalogSyncEditResponseEnvelopeErrorsCode = 102009
	CatalogSyncEditResponseEnvelopeErrorsCode102010 CatalogSyncEditResponseEnvelopeErrorsCode = 102010
	CatalogSyncEditResponseEnvelopeErrorsCode102011 CatalogSyncEditResponseEnvelopeErrorsCode = 102011
	CatalogSyncEditResponseEnvelopeErrorsCode102012 CatalogSyncEditResponseEnvelopeErrorsCode = 102012
	CatalogSyncEditResponseEnvelopeErrorsCode102013 CatalogSyncEditResponseEnvelopeErrorsCode = 102013
	CatalogSyncEditResponseEnvelopeErrorsCode102014 CatalogSyncEditResponseEnvelopeErrorsCode = 102014
	CatalogSyncEditResponseEnvelopeErrorsCode102015 CatalogSyncEditResponseEnvelopeErrorsCode = 102015
	CatalogSyncEditResponseEnvelopeErrorsCode102016 CatalogSyncEditResponseEnvelopeErrorsCode = 102016
	CatalogSyncEditResponseEnvelopeErrorsCode102017 CatalogSyncEditResponseEnvelopeErrorsCode = 102017
	CatalogSyncEditResponseEnvelopeErrorsCode102018 CatalogSyncEditResponseEnvelopeErrorsCode = 102018
	CatalogSyncEditResponseEnvelopeErrorsCode102019 CatalogSyncEditResponseEnvelopeErrorsCode = 102019
	CatalogSyncEditResponseEnvelopeErrorsCode102020 CatalogSyncEditResponseEnvelopeErrorsCode = 102020
	CatalogSyncEditResponseEnvelopeErrorsCode102021 CatalogSyncEditResponseEnvelopeErrorsCode = 102021
	CatalogSyncEditResponseEnvelopeErrorsCode102022 CatalogSyncEditResponseEnvelopeErrorsCode = 102022
	CatalogSyncEditResponseEnvelopeErrorsCode102023 CatalogSyncEditResponseEnvelopeErrorsCode = 102023
	CatalogSyncEditResponseEnvelopeErrorsCode102024 CatalogSyncEditResponseEnvelopeErrorsCode = 102024
	CatalogSyncEditResponseEnvelopeErrorsCode102025 CatalogSyncEditResponseEnvelopeErrorsCode = 102025
	CatalogSyncEditResponseEnvelopeErrorsCode102026 CatalogSyncEditResponseEnvelopeErrorsCode = 102026
	CatalogSyncEditResponseEnvelopeErrorsCode102027 CatalogSyncEditResponseEnvelopeErrorsCode = 102027
	CatalogSyncEditResponseEnvelopeErrorsCode102028 CatalogSyncEditResponseEnvelopeErrorsCode = 102028
	CatalogSyncEditResponseEnvelopeErrorsCode102029 CatalogSyncEditResponseEnvelopeErrorsCode = 102029
	CatalogSyncEditResponseEnvelopeErrorsCode102030 CatalogSyncEditResponseEnvelopeErrorsCode = 102030
	CatalogSyncEditResponseEnvelopeErrorsCode102031 CatalogSyncEditResponseEnvelopeErrorsCode = 102031
	CatalogSyncEditResponseEnvelopeErrorsCode102032 CatalogSyncEditResponseEnvelopeErrorsCode = 102032
	CatalogSyncEditResponseEnvelopeErrorsCode102033 CatalogSyncEditResponseEnvelopeErrorsCode = 102033
	CatalogSyncEditResponseEnvelopeErrorsCode102034 CatalogSyncEditResponseEnvelopeErrorsCode = 102034
	CatalogSyncEditResponseEnvelopeErrorsCode102035 CatalogSyncEditResponseEnvelopeErrorsCode = 102035
	CatalogSyncEditResponseEnvelopeErrorsCode102036 CatalogSyncEditResponseEnvelopeErrorsCode = 102036
	CatalogSyncEditResponseEnvelopeErrorsCode102037 CatalogSyncEditResponseEnvelopeErrorsCode = 102037
	CatalogSyncEditResponseEnvelopeErrorsCode102038 CatalogSyncEditResponseEnvelopeErrorsCode = 102038
	CatalogSyncEditResponseEnvelopeErrorsCode102039 CatalogSyncEditResponseEnvelopeErrorsCode = 102039
	CatalogSyncEditResponseEnvelopeErrorsCode102040 CatalogSyncEditResponseEnvelopeErrorsCode = 102040
	CatalogSyncEditResponseEnvelopeErrorsCode102041 CatalogSyncEditResponseEnvelopeErrorsCode = 102041
	CatalogSyncEditResponseEnvelopeErrorsCode102042 CatalogSyncEditResponseEnvelopeErrorsCode = 102042
	CatalogSyncEditResponseEnvelopeErrorsCode102043 CatalogSyncEditResponseEnvelopeErrorsCode = 102043
	CatalogSyncEditResponseEnvelopeErrorsCode102044 CatalogSyncEditResponseEnvelopeErrorsCode = 102044
	CatalogSyncEditResponseEnvelopeErrorsCode102045 CatalogSyncEditResponseEnvelopeErrorsCode = 102045
	CatalogSyncEditResponseEnvelopeErrorsCode102046 CatalogSyncEditResponseEnvelopeErrorsCode = 102046
	CatalogSyncEditResponseEnvelopeErrorsCode102047 CatalogSyncEditResponseEnvelopeErrorsCode = 102047
	CatalogSyncEditResponseEnvelopeErrorsCode102048 CatalogSyncEditResponseEnvelopeErrorsCode = 102048
	CatalogSyncEditResponseEnvelopeErrorsCode102049 CatalogSyncEditResponseEnvelopeErrorsCode = 102049
	CatalogSyncEditResponseEnvelopeErrorsCode102050 CatalogSyncEditResponseEnvelopeErrorsCode = 102050
	CatalogSyncEditResponseEnvelopeErrorsCode102051 CatalogSyncEditResponseEnvelopeErrorsCode = 102051
	CatalogSyncEditResponseEnvelopeErrorsCode102052 CatalogSyncEditResponseEnvelopeErrorsCode = 102052
	CatalogSyncEditResponseEnvelopeErrorsCode102053 CatalogSyncEditResponseEnvelopeErrorsCode = 102053
	CatalogSyncEditResponseEnvelopeErrorsCode102054 CatalogSyncEditResponseEnvelopeErrorsCode = 102054
	CatalogSyncEditResponseEnvelopeErrorsCode102055 CatalogSyncEditResponseEnvelopeErrorsCode = 102055
	CatalogSyncEditResponseEnvelopeErrorsCode102056 CatalogSyncEditResponseEnvelopeErrorsCode = 102056
	CatalogSyncEditResponseEnvelopeErrorsCode102057 CatalogSyncEditResponseEnvelopeErrorsCode = 102057
	CatalogSyncEditResponseEnvelopeErrorsCode102058 CatalogSyncEditResponseEnvelopeErrorsCode = 102058
	CatalogSyncEditResponseEnvelopeErrorsCode102059 CatalogSyncEditResponseEnvelopeErrorsCode = 102059
	CatalogSyncEditResponseEnvelopeErrorsCode102060 CatalogSyncEditResponseEnvelopeErrorsCode = 102060
	CatalogSyncEditResponseEnvelopeErrorsCode102061 CatalogSyncEditResponseEnvelopeErrorsCode = 102061
	CatalogSyncEditResponseEnvelopeErrorsCode102062 CatalogSyncEditResponseEnvelopeErrorsCode = 102062
	CatalogSyncEditResponseEnvelopeErrorsCode102063 CatalogSyncEditResponseEnvelopeErrorsCode = 102063
	CatalogSyncEditResponseEnvelopeErrorsCode102064 CatalogSyncEditResponseEnvelopeErrorsCode = 102064
	CatalogSyncEditResponseEnvelopeErrorsCode102065 CatalogSyncEditResponseEnvelopeErrorsCode = 102065
	CatalogSyncEditResponseEnvelopeErrorsCode102066 CatalogSyncEditResponseEnvelopeErrorsCode = 102066
	CatalogSyncEditResponseEnvelopeErrorsCode103001 CatalogSyncEditResponseEnvelopeErrorsCode = 103001
	CatalogSyncEditResponseEnvelopeErrorsCode103002 CatalogSyncEditResponseEnvelopeErrorsCode = 103002
	CatalogSyncEditResponseEnvelopeErrorsCode103003 CatalogSyncEditResponseEnvelopeErrorsCode = 103003
	CatalogSyncEditResponseEnvelopeErrorsCode103004 CatalogSyncEditResponseEnvelopeErrorsCode = 103004
	CatalogSyncEditResponseEnvelopeErrorsCode103005 CatalogSyncEditResponseEnvelopeErrorsCode = 103005
	CatalogSyncEditResponseEnvelopeErrorsCode103006 CatalogSyncEditResponseEnvelopeErrorsCode = 103006
	CatalogSyncEditResponseEnvelopeErrorsCode103007 CatalogSyncEditResponseEnvelopeErrorsCode = 103007
	CatalogSyncEditResponseEnvelopeErrorsCode103008 CatalogSyncEditResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncEditResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncEditResponseEnvelopeErrorsCode1001, CatalogSyncEditResponseEnvelopeErrorsCode1002, CatalogSyncEditResponseEnvelopeErrorsCode1003, CatalogSyncEditResponseEnvelopeErrorsCode1004, CatalogSyncEditResponseEnvelopeErrorsCode1005, CatalogSyncEditResponseEnvelopeErrorsCode1006, CatalogSyncEditResponseEnvelopeErrorsCode1007, CatalogSyncEditResponseEnvelopeErrorsCode1008, CatalogSyncEditResponseEnvelopeErrorsCode1009, CatalogSyncEditResponseEnvelopeErrorsCode1010, CatalogSyncEditResponseEnvelopeErrorsCode1011, CatalogSyncEditResponseEnvelopeErrorsCode1012, CatalogSyncEditResponseEnvelopeErrorsCode1013, CatalogSyncEditResponseEnvelopeErrorsCode1014, CatalogSyncEditResponseEnvelopeErrorsCode1015, CatalogSyncEditResponseEnvelopeErrorsCode1016, CatalogSyncEditResponseEnvelopeErrorsCode1017, CatalogSyncEditResponseEnvelopeErrorsCode2001, CatalogSyncEditResponseEnvelopeErrorsCode2002, CatalogSyncEditResponseEnvelopeErrorsCode2003, CatalogSyncEditResponseEnvelopeErrorsCode2004, CatalogSyncEditResponseEnvelopeErrorsCode2005, CatalogSyncEditResponseEnvelopeErrorsCode2006, CatalogSyncEditResponseEnvelopeErrorsCode2007, CatalogSyncEditResponseEnvelopeErrorsCode2008, CatalogSyncEditResponseEnvelopeErrorsCode2009, CatalogSyncEditResponseEnvelopeErrorsCode2010, CatalogSyncEditResponseEnvelopeErrorsCode2011, CatalogSyncEditResponseEnvelopeErrorsCode2012, CatalogSyncEditResponseEnvelopeErrorsCode2013, CatalogSyncEditResponseEnvelopeErrorsCode2014, CatalogSyncEditResponseEnvelopeErrorsCode2015, CatalogSyncEditResponseEnvelopeErrorsCode2016, CatalogSyncEditResponseEnvelopeErrorsCode2017, CatalogSyncEditResponseEnvelopeErrorsCode2018, CatalogSyncEditResponseEnvelopeErrorsCode2019, CatalogSyncEditResponseEnvelopeErrorsCode2020, CatalogSyncEditResponseEnvelopeErrorsCode2021, CatalogSyncEditResponseEnvelopeErrorsCode2022, CatalogSyncEditResponseEnvelopeErrorsCode3001, CatalogSyncEditResponseEnvelopeErrorsCode3002, CatalogSyncEditResponseEnvelopeErrorsCode3003, CatalogSyncEditResponseEnvelopeErrorsCode3004, CatalogSyncEditResponseEnvelopeErrorsCode3005, CatalogSyncEditResponseEnvelopeErrorsCode3006, CatalogSyncEditResponseEnvelopeErrorsCode3007, CatalogSyncEditResponseEnvelopeErrorsCode4001, CatalogSyncEditResponseEnvelopeErrorsCode4002, CatalogSyncEditResponseEnvelopeErrorsCode4003, CatalogSyncEditResponseEnvelopeErrorsCode4004, CatalogSyncEditResponseEnvelopeErrorsCode4005, CatalogSyncEditResponseEnvelopeErrorsCode4006, CatalogSyncEditResponseEnvelopeErrorsCode4007, CatalogSyncEditResponseEnvelopeErrorsCode4008, CatalogSyncEditResponseEnvelopeErrorsCode4009, CatalogSyncEditResponseEnvelopeErrorsCode4010, CatalogSyncEditResponseEnvelopeErrorsCode4011, CatalogSyncEditResponseEnvelopeErrorsCode4012, CatalogSyncEditResponseEnvelopeErrorsCode4013, CatalogSyncEditResponseEnvelopeErrorsCode4014, CatalogSyncEditResponseEnvelopeErrorsCode4015, CatalogSyncEditResponseEnvelopeErrorsCode4016, CatalogSyncEditResponseEnvelopeErrorsCode4017, CatalogSyncEditResponseEnvelopeErrorsCode4018, CatalogSyncEditResponseEnvelopeErrorsCode4019, CatalogSyncEditResponseEnvelopeErrorsCode4020, CatalogSyncEditResponseEnvelopeErrorsCode4021, CatalogSyncEditResponseEnvelopeErrorsCode4022, CatalogSyncEditResponseEnvelopeErrorsCode4023, CatalogSyncEditResponseEnvelopeErrorsCode5001, CatalogSyncEditResponseEnvelopeErrorsCode5002, CatalogSyncEditResponseEnvelopeErrorsCode5003, CatalogSyncEditResponseEnvelopeErrorsCode5004, CatalogSyncEditResponseEnvelopeErrorsCode102000, CatalogSyncEditResponseEnvelopeErrorsCode102001, CatalogSyncEditResponseEnvelopeErrorsCode102002, CatalogSyncEditResponseEnvelopeErrorsCode102003, CatalogSyncEditResponseEnvelopeErrorsCode102004, CatalogSyncEditResponseEnvelopeErrorsCode102005, CatalogSyncEditResponseEnvelopeErrorsCode102006, CatalogSyncEditResponseEnvelopeErrorsCode102007, CatalogSyncEditResponseEnvelopeErrorsCode102008, CatalogSyncEditResponseEnvelopeErrorsCode102009, CatalogSyncEditResponseEnvelopeErrorsCode102010, CatalogSyncEditResponseEnvelopeErrorsCode102011, CatalogSyncEditResponseEnvelopeErrorsCode102012, CatalogSyncEditResponseEnvelopeErrorsCode102013, CatalogSyncEditResponseEnvelopeErrorsCode102014, CatalogSyncEditResponseEnvelopeErrorsCode102015, CatalogSyncEditResponseEnvelopeErrorsCode102016, CatalogSyncEditResponseEnvelopeErrorsCode102017, CatalogSyncEditResponseEnvelopeErrorsCode102018, CatalogSyncEditResponseEnvelopeErrorsCode102019, CatalogSyncEditResponseEnvelopeErrorsCode102020, CatalogSyncEditResponseEnvelopeErrorsCode102021, CatalogSyncEditResponseEnvelopeErrorsCode102022, CatalogSyncEditResponseEnvelopeErrorsCode102023, CatalogSyncEditResponseEnvelopeErrorsCode102024, CatalogSyncEditResponseEnvelopeErrorsCode102025, CatalogSyncEditResponseEnvelopeErrorsCode102026, CatalogSyncEditResponseEnvelopeErrorsCode102027, CatalogSyncEditResponseEnvelopeErrorsCode102028, CatalogSyncEditResponseEnvelopeErrorsCode102029, CatalogSyncEditResponseEnvelopeErrorsCode102030, CatalogSyncEditResponseEnvelopeErrorsCode102031, CatalogSyncEditResponseEnvelopeErrorsCode102032, CatalogSyncEditResponseEnvelopeErrorsCode102033, CatalogSyncEditResponseEnvelopeErrorsCode102034, CatalogSyncEditResponseEnvelopeErrorsCode102035, CatalogSyncEditResponseEnvelopeErrorsCode102036, CatalogSyncEditResponseEnvelopeErrorsCode102037, CatalogSyncEditResponseEnvelopeErrorsCode102038, CatalogSyncEditResponseEnvelopeErrorsCode102039, CatalogSyncEditResponseEnvelopeErrorsCode102040, CatalogSyncEditResponseEnvelopeErrorsCode102041, CatalogSyncEditResponseEnvelopeErrorsCode102042, CatalogSyncEditResponseEnvelopeErrorsCode102043, CatalogSyncEditResponseEnvelopeErrorsCode102044, CatalogSyncEditResponseEnvelopeErrorsCode102045, CatalogSyncEditResponseEnvelopeErrorsCode102046, CatalogSyncEditResponseEnvelopeErrorsCode102047, CatalogSyncEditResponseEnvelopeErrorsCode102048, CatalogSyncEditResponseEnvelopeErrorsCode102049, CatalogSyncEditResponseEnvelopeErrorsCode102050, CatalogSyncEditResponseEnvelopeErrorsCode102051, CatalogSyncEditResponseEnvelopeErrorsCode102052, CatalogSyncEditResponseEnvelopeErrorsCode102053, CatalogSyncEditResponseEnvelopeErrorsCode102054, CatalogSyncEditResponseEnvelopeErrorsCode102055, CatalogSyncEditResponseEnvelopeErrorsCode102056, CatalogSyncEditResponseEnvelopeErrorsCode102057, CatalogSyncEditResponseEnvelopeErrorsCode102058, CatalogSyncEditResponseEnvelopeErrorsCode102059, CatalogSyncEditResponseEnvelopeErrorsCode102060, CatalogSyncEditResponseEnvelopeErrorsCode102061, CatalogSyncEditResponseEnvelopeErrorsCode102062, CatalogSyncEditResponseEnvelopeErrorsCode102063, CatalogSyncEditResponseEnvelopeErrorsCode102064, CatalogSyncEditResponseEnvelopeErrorsCode102065, CatalogSyncEditResponseEnvelopeErrorsCode102066, CatalogSyncEditResponseEnvelopeErrorsCode103001, CatalogSyncEditResponseEnvelopeErrorsCode103002, CatalogSyncEditResponseEnvelopeErrorsCode103003, CatalogSyncEditResponseEnvelopeErrorsCode103004, CatalogSyncEditResponseEnvelopeErrorsCode103005, CatalogSyncEditResponseEnvelopeErrorsCode103006, CatalogSyncEditResponseEnvelopeErrorsCode103007, CatalogSyncEditResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncEditResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                        `json:"l10n_key"`
	LoggableError string                                        `json:"loggable_error"`
	TemplateData  interface{}                                   `json:"template_data"`
	TraceID       string                                        `json:"trace_id"`
	JSON          catalogSyncEditResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncEditResponseEnvelopeErrorsMetaJSON contains the JSON metadata for the
// struct [CatalogSyncEditResponseEnvelopeErrorsMeta]
type catalogSyncEditResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeErrorsSource struct {
	Parameter           string                                          `json:"parameter"`
	ParameterValueIndex int64                                           `json:"parameter_value_index"`
	Pointer             string                                          `json:"pointer"`
	JSON                catalogSyncEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CatalogSyncEditResponseEnvelopeErrorsSource]
type catalogSyncEditResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeMessages struct {
	Code             CatalogSyncEditResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Meta             CatalogSyncEditResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CatalogSyncEditResponseEnvelopeMessages]
type catalogSyncEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeMessagesCode int64

const (
	CatalogSyncEditResponseEnvelopeMessagesCode1001   CatalogSyncEditResponseEnvelopeMessagesCode = 1001
	CatalogSyncEditResponseEnvelopeMessagesCode1002   CatalogSyncEditResponseEnvelopeMessagesCode = 1002
	CatalogSyncEditResponseEnvelopeMessagesCode1003   CatalogSyncEditResponseEnvelopeMessagesCode = 1003
	CatalogSyncEditResponseEnvelopeMessagesCode1004   CatalogSyncEditResponseEnvelopeMessagesCode = 1004
	CatalogSyncEditResponseEnvelopeMessagesCode1005   CatalogSyncEditResponseEnvelopeMessagesCode = 1005
	CatalogSyncEditResponseEnvelopeMessagesCode1006   CatalogSyncEditResponseEnvelopeMessagesCode = 1006
	CatalogSyncEditResponseEnvelopeMessagesCode1007   CatalogSyncEditResponseEnvelopeMessagesCode = 1007
	CatalogSyncEditResponseEnvelopeMessagesCode1008   CatalogSyncEditResponseEnvelopeMessagesCode = 1008
	CatalogSyncEditResponseEnvelopeMessagesCode1009   CatalogSyncEditResponseEnvelopeMessagesCode = 1009
	CatalogSyncEditResponseEnvelopeMessagesCode1010   CatalogSyncEditResponseEnvelopeMessagesCode = 1010
	CatalogSyncEditResponseEnvelopeMessagesCode1011   CatalogSyncEditResponseEnvelopeMessagesCode = 1011
	CatalogSyncEditResponseEnvelopeMessagesCode1012   CatalogSyncEditResponseEnvelopeMessagesCode = 1012
	CatalogSyncEditResponseEnvelopeMessagesCode1013   CatalogSyncEditResponseEnvelopeMessagesCode = 1013
	CatalogSyncEditResponseEnvelopeMessagesCode1014   CatalogSyncEditResponseEnvelopeMessagesCode = 1014
	CatalogSyncEditResponseEnvelopeMessagesCode1015   CatalogSyncEditResponseEnvelopeMessagesCode = 1015
	CatalogSyncEditResponseEnvelopeMessagesCode1016   CatalogSyncEditResponseEnvelopeMessagesCode = 1016
	CatalogSyncEditResponseEnvelopeMessagesCode1017   CatalogSyncEditResponseEnvelopeMessagesCode = 1017
	CatalogSyncEditResponseEnvelopeMessagesCode2001   CatalogSyncEditResponseEnvelopeMessagesCode = 2001
	CatalogSyncEditResponseEnvelopeMessagesCode2002   CatalogSyncEditResponseEnvelopeMessagesCode = 2002
	CatalogSyncEditResponseEnvelopeMessagesCode2003   CatalogSyncEditResponseEnvelopeMessagesCode = 2003
	CatalogSyncEditResponseEnvelopeMessagesCode2004   CatalogSyncEditResponseEnvelopeMessagesCode = 2004
	CatalogSyncEditResponseEnvelopeMessagesCode2005   CatalogSyncEditResponseEnvelopeMessagesCode = 2005
	CatalogSyncEditResponseEnvelopeMessagesCode2006   CatalogSyncEditResponseEnvelopeMessagesCode = 2006
	CatalogSyncEditResponseEnvelopeMessagesCode2007   CatalogSyncEditResponseEnvelopeMessagesCode = 2007
	CatalogSyncEditResponseEnvelopeMessagesCode2008   CatalogSyncEditResponseEnvelopeMessagesCode = 2008
	CatalogSyncEditResponseEnvelopeMessagesCode2009   CatalogSyncEditResponseEnvelopeMessagesCode = 2009
	CatalogSyncEditResponseEnvelopeMessagesCode2010   CatalogSyncEditResponseEnvelopeMessagesCode = 2010
	CatalogSyncEditResponseEnvelopeMessagesCode2011   CatalogSyncEditResponseEnvelopeMessagesCode = 2011
	CatalogSyncEditResponseEnvelopeMessagesCode2012   CatalogSyncEditResponseEnvelopeMessagesCode = 2012
	CatalogSyncEditResponseEnvelopeMessagesCode2013   CatalogSyncEditResponseEnvelopeMessagesCode = 2013
	CatalogSyncEditResponseEnvelopeMessagesCode2014   CatalogSyncEditResponseEnvelopeMessagesCode = 2014
	CatalogSyncEditResponseEnvelopeMessagesCode2015   CatalogSyncEditResponseEnvelopeMessagesCode = 2015
	CatalogSyncEditResponseEnvelopeMessagesCode2016   CatalogSyncEditResponseEnvelopeMessagesCode = 2016
	CatalogSyncEditResponseEnvelopeMessagesCode2017   CatalogSyncEditResponseEnvelopeMessagesCode = 2017
	CatalogSyncEditResponseEnvelopeMessagesCode2018   CatalogSyncEditResponseEnvelopeMessagesCode = 2018
	CatalogSyncEditResponseEnvelopeMessagesCode2019   CatalogSyncEditResponseEnvelopeMessagesCode = 2019
	CatalogSyncEditResponseEnvelopeMessagesCode2020   CatalogSyncEditResponseEnvelopeMessagesCode = 2020
	CatalogSyncEditResponseEnvelopeMessagesCode2021   CatalogSyncEditResponseEnvelopeMessagesCode = 2021
	CatalogSyncEditResponseEnvelopeMessagesCode2022   CatalogSyncEditResponseEnvelopeMessagesCode = 2022
	CatalogSyncEditResponseEnvelopeMessagesCode3001   CatalogSyncEditResponseEnvelopeMessagesCode = 3001
	CatalogSyncEditResponseEnvelopeMessagesCode3002   CatalogSyncEditResponseEnvelopeMessagesCode = 3002
	CatalogSyncEditResponseEnvelopeMessagesCode3003   CatalogSyncEditResponseEnvelopeMessagesCode = 3003
	CatalogSyncEditResponseEnvelopeMessagesCode3004   CatalogSyncEditResponseEnvelopeMessagesCode = 3004
	CatalogSyncEditResponseEnvelopeMessagesCode3005   CatalogSyncEditResponseEnvelopeMessagesCode = 3005
	CatalogSyncEditResponseEnvelopeMessagesCode3006   CatalogSyncEditResponseEnvelopeMessagesCode = 3006
	CatalogSyncEditResponseEnvelopeMessagesCode3007   CatalogSyncEditResponseEnvelopeMessagesCode = 3007
	CatalogSyncEditResponseEnvelopeMessagesCode4001   CatalogSyncEditResponseEnvelopeMessagesCode = 4001
	CatalogSyncEditResponseEnvelopeMessagesCode4002   CatalogSyncEditResponseEnvelopeMessagesCode = 4002
	CatalogSyncEditResponseEnvelopeMessagesCode4003   CatalogSyncEditResponseEnvelopeMessagesCode = 4003
	CatalogSyncEditResponseEnvelopeMessagesCode4004   CatalogSyncEditResponseEnvelopeMessagesCode = 4004
	CatalogSyncEditResponseEnvelopeMessagesCode4005   CatalogSyncEditResponseEnvelopeMessagesCode = 4005
	CatalogSyncEditResponseEnvelopeMessagesCode4006   CatalogSyncEditResponseEnvelopeMessagesCode = 4006
	CatalogSyncEditResponseEnvelopeMessagesCode4007   CatalogSyncEditResponseEnvelopeMessagesCode = 4007
	CatalogSyncEditResponseEnvelopeMessagesCode4008   CatalogSyncEditResponseEnvelopeMessagesCode = 4008
	CatalogSyncEditResponseEnvelopeMessagesCode4009   CatalogSyncEditResponseEnvelopeMessagesCode = 4009
	CatalogSyncEditResponseEnvelopeMessagesCode4010   CatalogSyncEditResponseEnvelopeMessagesCode = 4010
	CatalogSyncEditResponseEnvelopeMessagesCode4011   CatalogSyncEditResponseEnvelopeMessagesCode = 4011
	CatalogSyncEditResponseEnvelopeMessagesCode4012   CatalogSyncEditResponseEnvelopeMessagesCode = 4012
	CatalogSyncEditResponseEnvelopeMessagesCode4013   CatalogSyncEditResponseEnvelopeMessagesCode = 4013
	CatalogSyncEditResponseEnvelopeMessagesCode4014   CatalogSyncEditResponseEnvelopeMessagesCode = 4014
	CatalogSyncEditResponseEnvelopeMessagesCode4015   CatalogSyncEditResponseEnvelopeMessagesCode = 4015
	CatalogSyncEditResponseEnvelopeMessagesCode4016   CatalogSyncEditResponseEnvelopeMessagesCode = 4016
	CatalogSyncEditResponseEnvelopeMessagesCode4017   CatalogSyncEditResponseEnvelopeMessagesCode = 4017
	CatalogSyncEditResponseEnvelopeMessagesCode4018   CatalogSyncEditResponseEnvelopeMessagesCode = 4018
	CatalogSyncEditResponseEnvelopeMessagesCode4019   CatalogSyncEditResponseEnvelopeMessagesCode = 4019
	CatalogSyncEditResponseEnvelopeMessagesCode4020   CatalogSyncEditResponseEnvelopeMessagesCode = 4020
	CatalogSyncEditResponseEnvelopeMessagesCode4021   CatalogSyncEditResponseEnvelopeMessagesCode = 4021
	CatalogSyncEditResponseEnvelopeMessagesCode4022   CatalogSyncEditResponseEnvelopeMessagesCode = 4022
	CatalogSyncEditResponseEnvelopeMessagesCode4023   CatalogSyncEditResponseEnvelopeMessagesCode = 4023
	CatalogSyncEditResponseEnvelopeMessagesCode5001   CatalogSyncEditResponseEnvelopeMessagesCode = 5001
	CatalogSyncEditResponseEnvelopeMessagesCode5002   CatalogSyncEditResponseEnvelopeMessagesCode = 5002
	CatalogSyncEditResponseEnvelopeMessagesCode5003   CatalogSyncEditResponseEnvelopeMessagesCode = 5003
	CatalogSyncEditResponseEnvelopeMessagesCode5004   CatalogSyncEditResponseEnvelopeMessagesCode = 5004
	CatalogSyncEditResponseEnvelopeMessagesCode102000 CatalogSyncEditResponseEnvelopeMessagesCode = 102000
	CatalogSyncEditResponseEnvelopeMessagesCode102001 CatalogSyncEditResponseEnvelopeMessagesCode = 102001
	CatalogSyncEditResponseEnvelopeMessagesCode102002 CatalogSyncEditResponseEnvelopeMessagesCode = 102002
	CatalogSyncEditResponseEnvelopeMessagesCode102003 CatalogSyncEditResponseEnvelopeMessagesCode = 102003
	CatalogSyncEditResponseEnvelopeMessagesCode102004 CatalogSyncEditResponseEnvelopeMessagesCode = 102004
	CatalogSyncEditResponseEnvelopeMessagesCode102005 CatalogSyncEditResponseEnvelopeMessagesCode = 102005
	CatalogSyncEditResponseEnvelopeMessagesCode102006 CatalogSyncEditResponseEnvelopeMessagesCode = 102006
	CatalogSyncEditResponseEnvelopeMessagesCode102007 CatalogSyncEditResponseEnvelopeMessagesCode = 102007
	CatalogSyncEditResponseEnvelopeMessagesCode102008 CatalogSyncEditResponseEnvelopeMessagesCode = 102008
	CatalogSyncEditResponseEnvelopeMessagesCode102009 CatalogSyncEditResponseEnvelopeMessagesCode = 102009
	CatalogSyncEditResponseEnvelopeMessagesCode102010 CatalogSyncEditResponseEnvelopeMessagesCode = 102010
	CatalogSyncEditResponseEnvelopeMessagesCode102011 CatalogSyncEditResponseEnvelopeMessagesCode = 102011
	CatalogSyncEditResponseEnvelopeMessagesCode102012 CatalogSyncEditResponseEnvelopeMessagesCode = 102012
	CatalogSyncEditResponseEnvelopeMessagesCode102013 CatalogSyncEditResponseEnvelopeMessagesCode = 102013
	CatalogSyncEditResponseEnvelopeMessagesCode102014 CatalogSyncEditResponseEnvelopeMessagesCode = 102014
	CatalogSyncEditResponseEnvelopeMessagesCode102015 CatalogSyncEditResponseEnvelopeMessagesCode = 102015
	CatalogSyncEditResponseEnvelopeMessagesCode102016 CatalogSyncEditResponseEnvelopeMessagesCode = 102016
	CatalogSyncEditResponseEnvelopeMessagesCode102017 CatalogSyncEditResponseEnvelopeMessagesCode = 102017
	CatalogSyncEditResponseEnvelopeMessagesCode102018 CatalogSyncEditResponseEnvelopeMessagesCode = 102018
	CatalogSyncEditResponseEnvelopeMessagesCode102019 CatalogSyncEditResponseEnvelopeMessagesCode = 102019
	CatalogSyncEditResponseEnvelopeMessagesCode102020 CatalogSyncEditResponseEnvelopeMessagesCode = 102020
	CatalogSyncEditResponseEnvelopeMessagesCode102021 CatalogSyncEditResponseEnvelopeMessagesCode = 102021
	CatalogSyncEditResponseEnvelopeMessagesCode102022 CatalogSyncEditResponseEnvelopeMessagesCode = 102022
	CatalogSyncEditResponseEnvelopeMessagesCode102023 CatalogSyncEditResponseEnvelopeMessagesCode = 102023
	CatalogSyncEditResponseEnvelopeMessagesCode102024 CatalogSyncEditResponseEnvelopeMessagesCode = 102024
	CatalogSyncEditResponseEnvelopeMessagesCode102025 CatalogSyncEditResponseEnvelopeMessagesCode = 102025
	CatalogSyncEditResponseEnvelopeMessagesCode102026 CatalogSyncEditResponseEnvelopeMessagesCode = 102026
	CatalogSyncEditResponseEnvelopeMessagesCode102027 CatalogSyncEditResponseEnvelopeMessagesCode = 102027
	CatalogSyncEditResponseEnvelopeMessagesCode102028 CatalogSyncEditResponseEnvelopeMessagesCode = 102028
	CatalogSyncEditResponseEnvelopeMessagesCode102029 CatalogSyncEditResponseEnvelopeMessagesCode = 102029
	CatalogSyncEditResponseEnvelopeMessagesCode102030 CatalogSyncEditResponseEnvelopeMessagesCode = 102030
	CatalogSyncEditResponseEnvelopeMessagesCode102031 CatalogSyncEditResponseEnvelopeMessagesCode = 102031
	CatalogSyncEditResponseEnvelopeMessagesCode102032 CatalogSyncEditResponseEnvelopeMessagesCode = 102032
	CatalogSyncEditResponseEnvelopeMessagesCode102033 CatalogSyncEditResponseEnvelopeMessagesCode = 102033
	CatalogSyncEditResponseEnvelopeMessagesCode102034 CatalogSyncEditResponseEnvelopeMessagesCode = 102034
	CatalogSyncEditResponseEnvelopeMessagesCode102035 CatalogSyncEditResponseEnvelopeMessagesCode = 102035
	CatalogSyncEditResponseEnvelopeMessagesCode102036 CatalogSyncEditResponseEnvelopeMessagesCode = 102036
	CatalogSyncEditResponseEnvelopeMessagesCode102037 CatalogSyncEditResponseEnvelopeMessagesCode = 102037
	CatalogSyncEditResponseEnvelopeMessagesCode102038 CatalogSyncEditResponseEnvelopeMessagesCode = 102038
	CatalogSyncEditResponseEnvelopeMessagesCode102039 CatalogSyncEditResponseEnvelopeMessagesCode = 102039
	CatalogSyncEditResponseEnvelopeMessagesCode102040 CatalogSyncEditResponseEnvelopeMessagesCode = 102040
	CatalogSyncEditResponseEnvelopeMessagesCode102041 CatalogSyncEditResponseEnvelopeMessagesCode = 102041
	CatalogSyncEditResponseEnvelopeMessagesCode102042 CatalogSyncEditResponseEnvelopeMessagesCode = 102042
	CatalogSyncEditResponseEnvelopeMessagesCode102043 CatalogSyncEditResponseEnvelopeMessagesCode = 102043
	CatalogSyncEditResponseEnvelopeMessagesCode102044 CatalogSyncEditResponseEnvelopeMessagesCode = 102044
	CatalogSyncEditResponseEnvelopeMessagesCode102045 CatalogSyncEditResponseEnvelopeMessagesCode = 102045
	CatalogSyncEditResponseEnvelopeMessagesCode102046 CatalogSyncEditResponseEnvelopeMessagesCode = 102046
	CatalogSyncEditResponseEnvelopeMessagesCode102047 CatalogSyncEditResponseEnvelopeMessagesCode = 102047
	CatalogSyncEditResponseEnvelopeMessagesCode102048 CatalogSyncEditResponseEnvelopeMessagesCode = 102048
	CatalogSyncEditResponseEnvelopeMessagesCode102049 CatalogSyncEditResponseEnvelopeMessagesCode = 102049
	CatalogSyncEditResponseEnvelopeMessagesCode102050 CatalogSyncEditResponseEnvelopeMessagesCode = 102050
	CatalogSyncEditResponseEnvelopeMessagesCode102051 CatalogSyncEditResponseEnvelopeMessagesCode = 102051
	CatalogSyncEditResponseEnvelopeMessagesCode102052 CatalogSyncEditResponseEnvelopeMessagesCode = 102052
	CatalogSyncEditResponseEnvelopeMessagesCode102053 CatalogSyncEditResponseEnvelopeMessagesCode = 102053
	CatalogSyncEditResponseEnvelopeMessagesCode102054 CatalogSyncEditResponseEnvelopeMessagesCode = 102054
	CatalogSyncEditResponseEnvelopeMessagesCode102055 CatalogSyncEditResponseEnvelopeMessagesCode = 102055
	CatalogSyncEditResponseEnvelopeMessagesCode102056 CatalogSyncEditResponseEnvelopeMessagesCode = 102056
	CatalogSyncEditResponseEnvelopeMessagesCode102057 CatalogSyncEditResponseEnvelopeMessagesCode = 102057
	CatalogSyncEditResponseEnvelopeMessagesCode102058 CatalogSyncEditResponseEnvelopeMessagesCode = 102058
	CatalogSyncEditResponseEnvelopeMessagesCode102059 CatalogSyncEditResponseEnvelopeMessagesCode = 102059
	CatalogSyncEditResponseEnvelopeMessagesCode102060 CatalogSyncEditResponseEnvelopeMessagesCode = 102060
	CatalogSyncEditResponseEnvelopeMessagesCode102061 CatalogSyncEditResponseEnvelopeMessagesCode = 102061
	CatalogSyncEditResponseEnvelopeMessagesCode102062 CatalogSyncEditResponseEnvelopeMessagesCode = 102062
	CatalogSyncEditResponseEnvelopeMessagesCode102063 CatalogSyncEditResponseEnvelopeMessagesCode = 102063
	CatalogSyncEditResponseEnvelopeMessagesCode102064 CatalogSyncEditResponseEnvelopeMessagesCode = 102064
	CatalogSyncEditResponseEnvelopeMessagesCode102065 CatalogSyncEditResponseEnvelopeMessagesCode = 102065
	CatalogSyncEditResponseEnvelopeMessagesCode102066 CatalogSyncEditResponseEnvelopeMessagesCode = 102066
	CatalogSyncEditResponseEnvelopeMessagesCode103001 CatalogSyncEditResponseEnvelopeMessagesCode = 103001
	CatalogSyncEditResponseEnvelopeMessagesCode103002 CatalogSyncEditResponseEnvelopeMessagesCode = 103002
	CatalogSyncEditResponseEnvelopeMessagesCode103003 CatalogSyncEditResponseEnvelopeMessagesCode = 103003
	CatalogSyncEditResponseEnvelopeMessagesCode103004 CatalogSyncEditResponseEnvelopeMessagesCode = 103004
	CatalogSyncEditResponseEnvelopeMessagesCode103005 CatalogSyncEditResponseEnvelopeMessagesCode = 103005
	CatalogSyncEditResponseEnvelopeMessagesCode103006 CatalogSyncEditResponseEnvelopeMessagesCode = 103006
	CatalogSyncEditResponseEnvelopeMessagesCode103007 CatalogSyncEditResponseEnvelopeMessagesCode = 103007
	CatalogSyncEditResponseEnvelopeMessagesCode103008 CatalogSyncEditResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncEditResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncEditResponseEnvelopeMessagesCode1001, CatalogSyncEditResponseEnvelopeMessagesCode1002, CatalogSyncEditResponseEnvelopeMessagesCode1003, CatalogSyncEditResponseEnvelopeMessagesCode1004, CatalogSyncEditResponseEnvelopeMessagesCode1005, CatalogSyncEditResponseEnvelopeMessagesCode1006, CatalogSyncEditResponseEnvelopeMessagesCode1007, CatalogSyncEditResponseEnvelopeMessagesCode1008, CatalogSyncEditResponseEnvelopeMessagesCode1009, CatalogSyncEditResponseEnvelopeMessagesCode1010, CatalogSyncEditResponseEnvelopeMessagesCode1011, CatalogSyncEditResponseEnvelopeMessagesCode1012, CatalogSyncEditResponseEnvelopeMessagesCode1013, CatalogSyncEditResponseEnvelopeMessagesCode1014, CatalogSyncEditResponseEnvelopeMessagesCode1015, CatalogSyncEditResponseEnvelopeMessagesCode1016, CatalogSyncEditResponseEnvelopeMessagesCode1017, CatalogSyncEditResponseEnvelopeMessagesCode2001, CatalogSyncEditResponseEnvelopeMessagesCode2002, CatalogSyncEditResponseEnvelopeMessagesCode2003, CatalogSyncEditResponseEnvelopeMessagesCode2004, CatalogSyncEditResponseEnvelopeMessagesCode2005, CatalogSyncEditResponseEnvelopeMessagesCode2006, CatalogSyncEditResponseEnvelopeMessagesCode2007, CatalogSyncEditResponseEnvelopeMessagesCode2008, CatalogSyncEditResponseEnvelopeMessagesCode2009, CatalogSyncEditResponseEnvelopeMessagesCode2010, CatalogSyncEditResponseEnvelopeMessagesCode2011, CatalogSyncEditResponseEnvelopeMessagesCode2012, CatalogSyncEditResponseEnvelopeMessagesCode2013, CatalogSyncEditResponseEnvelopeMessagesCode2014, CatalogSyncEditResponseEnvelopeMessagesCode2015, CatalogSyncEditResponseEnvelopeMessagesCode2016, CatalogSyncEditResponseEnvelopeMessagesCode2017, CatalogSyncEditResponseEnvelopeMessagesCode2018, CatalogSyncEditResponseEnvelopeMessagesCode2019, CatalogSyncEditResponseEnvelopeMessagesCode2020, CatalogSyncEditResponseEnvelopeMessagesCode2021, CatalogSyncEditResponseEnvelopeMessagesCode2022, CatalogSyncEditResponseEnvelopeMessagesCode3001, CatalogSyncEditResponseEnvelopeMessagesCode3002, CatalogSyncEditResponseEnvelopeMessagesCode3003, CatalogSyncEditResponseEnvelopeMessagesCode3004, CatalogSyncEditResponseEnvelopeMessagesCode3005, CatalogSyncEditResponseEnvelopeMessagesCode3006, CatalogSyncEditResponseEnvelopeMessagesCode3007, CatalogSyncEditResponseEnvelopeMessagesCode4001, CatalogSyncEditResponseEnvelopeMessagesCode4002, CatalogSyncEditResponseEnvelopeMessagesCode4003, CatalogSyncEditResponseEnvelopeMessagesCode4004, CatalogSyncEditResponseEnvelopeMessagesCode4005, CatalogSyncEditResponseEnvelopeMessagesCode4006, CatalogSyncEditResponseEnvelopeMessagesCode4007, CatalogSyncEditResponseEnvelopeMessagesCode4008, CatalogSyncEditResponseEnvelopeMessagesCode4009, CatalogSyncEditResponseEnvelopeMessagesCode4010, CatalogSyncEditResponseEnvelopeMessagesCode4011, CatalogSyncEditResponseEnvelopeMessagesCode4012, CatalogSyncEditResponseEnvelopeMessagesCode4013, CatalogSyncEditResponseEnvelopeMessagesCode4014, CatalogSyncEditResponseEnvelopeMessagesCode4015, CatalogSyncEditResponseEnvelopeMessagesCode4016, CatalogSyncEditResponseEnvelopeMessagesCode4017, CatalogSyncEditResponseEnvelopeMessagesCode4018, CatalogSyncEditResponseEnvelopeMessagesCode4019, CatalogSyncEditResponseEnvelopeMessagesCode4020, CatalogSyncEditResponseEnvelopeMessagesCode4021, CatalogSyncEditResponseEnvelopeMessagesCode4022, CatalogSyncEditResponseEnvelopeMessagesCode4023, CatalogSyncEditResponseEnvelopeMessagesCode5001, CatalogSyncEditResponseEnvelopeMessagesCode5002, CatalogSyncEditResponseEnvelopeMessagesCode5003, CatalogSyncEditResponseEnvelopeMessagesCode5004, CatalogSyncEditResponseEnvelopeMessagesCode102000, CatalogSyncEditResponseEnvelopeMessagesCode102001, CatalogSyncEditResponseEnvelopeMessagesCode102002, CatalogSyncEditResponseEnvelopeMessagesCode102003, CatalogSyncEditResponseEnvelopeMessagesCode102004, CatalogSyncEditResponseEnvelopeMessagesCode102005, CatalogSyncEditResponseEnvelopeMessagesCode102006, CatalogSyncEditResponseEnvelopeMessagesCode102007, CatalogSyncEditResponseEnvelopeMessagesCode102008, CatalogSyncEditResponseEnvelopeMessagesCode102009, CatalogSyncEditResponseEnvelopeMessagesCode102010, CatalogSyncEditResponseEnvelopeMessagesCode102011, CatalogSyncEditResponseEnvelopeMessagesCode102012, CatalogSyncEditResponseEnvelopeMessagesCode102013, CatalogSyncEditResponseEnvelopeMessagesCode102014, CatalogSyncEditResponseEnvelopeMessagesCode102015, CatalogSyncEditResponseEnvelopeMessagesCode102016, CatalogSyncEditResponseEnvelopeMessagesCode102017, CatalogSyncEditResponseEnvelopeMessagesCode102018, CatalogSyncEditResponseEnvelopeMessagesCode102019, CatalogSyncEditResponseEnvelopeMessagesCode102020, CatalogSyncEditResponseEnvelopeMessagesCode102021, CatalogSyncEditResponseEnvelopeMessagesCode102022, CatalogSyncEditResponseEnvelopeMessagesCode102023, CatalogSyncEditResponseEnvelopeMessagesCode102024, CatalogSyncEditResponseEnvelopeMessagesCode102025, CatalogSyncEditResponseEnvelopeMessagesCode102026, CatalogSyncEditResponseEnvelopeMessagesCode102027, CatalogSyncEditResponseEnvelopeMessagesCode102028, CatalogSyncEditResponseEnvelopeMessagesCode102029, CatalogSyncEditResponseEnvelopeMessagesCode102030, CatalogSyncEditResponseEnvelopeMessagesCode102031, CatalogSyncEditResponseEnvelopeMessagesCode102032, CatalogSyncEditResponseEnvelopeMessagesCode102033, CatalogSyncEditResponseEnvelopeMessagesCode102034, CatalogSyncEditResponseEnvelopeMessagesCode102035, CatalogSyncEditResponseEnvelopeMessagesCode102036, CatalogSyncEditResponseEnvelopeMessagesCode102037, CatalogSyncEditResponseEnvelopeMessagesCode102038, CatalogSyncEditResponseEnvelopeMessagesCode102039, CatalogSyncEditResponseEnvelopeMessagesCode102040, CatalogSyncEditResponseEnvelopeMessagesCode102041, CatalogSyncEditResponseEnvelopeMessagesCode102042, CatalogSyncEditResponseEnvelopeMessagesCode102043, CatalogSyncEditResponseEnvelopeMessagesCode102044, CatalogSyncEditResponseEnvelopeMessagesCode102045, CatalogSyncEditResponseEnvelopeMessagesCode102046, CatalogSyncEditResponseEnvelopeMessagesCode102047, CatalogSyncEditResponseEnvelopeMessagesCode102048, CatalogSyncEditResponseEnvelopeMessagesCode102049, CatalogSyncEditResponseEnvelopeMessagesCode102050, CatalogSyncEditResponseEnvelopeMessagesCode102051, CatalogSyncEditResponseEnvelopeMessagesCode102052, CatalogSyncEditResponseEnvelopeMessagesCode102053, CatalogSyncEditResponseEnvelopeMessagesCode102054, CatalogSyncEditResponseEnvelopeMessagesCode102055, CatalogSyncEditResponseEnvelopeMessagesCode102056, CatalogSyncEditResponseEnvelopeMessagesCode102057, CatalogSyncEditResponseEnvelopeMessagesCode102058, CatalogSyncEditResponseEnvelopeMessagesCode102059, CatalogSyncEditResponseEnvelopeMessagesCode102060, CatalogSyncEditResponseEnvelopeMessagesCode102061, CatalogSyncEditResponseEnvelopeMessagesCode102062, CatalogSyncEditResponseEnvelopeMessagesCode102063, CatalogSyncEditResponseEnvelopeMessagesCode102064, CatalogSyncEditResponseEnvelopeMessagesCode102065, CatalogSyncEditResponseEnvelopeMessagesCode102066, CatalogSyncEditResponseEnvelopeMessagesCode103001, CatalogSyncEditResponseEnvelopeMessagesCode103002, CatalogSyncEditResponseEnvelopeMessagesCode103003, CatalogSyncEditResponseEnvelopeMessagesCode103004, CatalogSyncEditResponseEnvelopeMessagesCode103005, CatalogSyncEditResponseEnvelopeMessagesCode103006, CatalogSyncEditResponseEnvelopeMessagesCode103007, CatalogSyncEditResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncEditResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                          `json:"l10n_key"`
	LoggableError string                                          `json:"loggable_error"`
	TemplateData  interface{}                                     `json:"template_data"`
	TraceID       string                                          `json:"trace_id"`
	JSON          catalogSyncEditResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncEditResponseEnvelopeMessagesMetaJSON contains the JSON metadata for
// the struct [CatalogSyncEditResponseEnvelopeMessagesMeta]
type catalogSyncEditResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncEditResponseEnvelopeMessagesSource struct {
	Parameter           string                                            `json:"parameter"`
	ParameterValueIndex int64                                             `json:"parameter_value_index"`
	Pointer             string                                            `json:"pointer"`
	JSON                catalogSyncEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CatalogSyncEditResponseEnvelopeMessagesSource]
type catalogSyncEditResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CatalogSyncGetResponseEnvelope struct {
	Errors   []CatalogSyncGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CatalogSyncGetResponse                   `json:"result,required"`
	Success  bool                                     `json:"success,required"`
	JSON     catalogSyncGetResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncGetResponseEnvelope]
type catalogSyncGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeErrors struct {
	Code             CatalogSyncGetResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Meta             CatalogSyncGetResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncGetResponseEnvelopeErrors]
type catalogSyncGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeErrorsCode int64

const (
	CatalogSyncGetResponseEnvelopeErrorsCode1001   CatalogSyncGetResponseEnvelopeErrorsCode = 1001
	CatalogSyncGetResponseEnvelopeErrorsCode1002   CatalogSyncGetResponseEnvelopeErrorsCode = 1002
	CatalogSyncGetResponseEnvelopeErrorsCode1003   CatalogSyncGetResponseEnvelopeErrorsCode = 1003
	CatalogSyncGetResponseEnvelopeErrorsCode1004   CatalogSyncGetResponseEnvelopeErrorsCode = 1004
	CatalogSyncGetResponseEnvelopeErrorsCode1005   CatalogSyncGetResponseEnvelopeErrorsCode = 1005
	CatalogSyncGetResponseEnvelopeErrorsCode1006   CatalogSyncGetResponseEnvelopeErrorsCode = 1006
	CatalogSyncGetResponseEnvelopeErrorsCode1007   CatalogSyncGetResponseEnvelopeErrorsCode = 1007
	CatalogSyncGetResponseEnvelopeErrorsCode1008   CatalogSyncGetResponseEnvelopeErrorsCode = 1008
	CatalogSyncGetResponseEnvelopeErrorsCode1009   CatalogSyncGetResponseEnvelopeErrorsCode = 1009
	CatalogSyncGetResponseEnvelopeErrorsCode1010   CatalogSyncGetResponseEnvelopeErrorsCode = 1010
	CatalogSyncGetResponseEnvelopeErrorsCode1011   CatalogSyncGetResponseEnvelopeErrorsCode = 1011
	CatalogSyncGetResponseEnvelopeErrorsCode1012   CatalogSyncGetResponseEnvelopeErrorsCode = 1012
	CatalogSyncGetResponseEnvelopeErrorsCode1013   CatalogSyncGetResponseEnvelopeErrorsCode = 1013
	CatalogSyncGetResponseEnvelopeErrorsCode1014   CatalogSyncGetResponseEnvelopeErrorsCode = 1014
	CatalogSyncGetResponseEnvelopeErrorsCode1015   CatalogSyncGetResponseEnvelopeErrorsCode = 1015
	CatalogSyncGetResponseEnvelopeErrorsCode1016   CatalogSyncGetResponseEnvelopeErrorsCode = 1016
	CatalogSyncGetResponseEnvelopeErrorsCode1017   CatalogSyncGetResponseEnvelopeErrorsCode = 1017
	CatalogSyncGetResponseEnvelopeErrorsCode2001   CatalogSyncGetResponseEnvelopeErrorsCode = 2001
	CatalogSyncGetResponseEnvelopeErrorsCode2002   CatalogSyncGetResponseEnvelopeErrorsCode = 2002
	CatalogSyncGetResponseEnvelopeErrorsCode2003   CatalogSyncGetResponseEnvelopeErrorsCode = 2003
	CatalogSyncGetResponseEnvelopeErrorsCode2004   CatalogSyncGetResponseEnvelopeErrorsCode = 2004
	CatalogSyncGetResponseEnvelopeErrorsCode2005   CatalogSyncGetResponseEnvelopeErrorsCode = 2005
	CatalogSyncGetResponseEnvelopeErrorsCode2006   CatalogSyncGetResponseEnvelopeErrorsCode = 2006
	CatalogSyncGetResponseEnvelopeErrorsCode2007   CatalogSyncGetResponseEnvelopeErrorsCode = 2007
	CatalogSyncGetResponseEnvelopeErrorsCode2008   CatalogSyncGetResponseEnvelopeErrorsCode = 2008
	CatalogSyncGetResponseEnvelopeErrorsCode2009   CatalogSyncGetResponseEnvelopeErrorsCode = 2009
	CatalogSyncGetResponseEnvelopeErrorsCode2010   CatalogSyncGetResponseEnvelopeErrorsCode = 2010
	CatalogSyncGetResponseEnvelopeErrorsCode2011   CatalogSyncGetResponseEnvelopeErrorsCode = 2011
	CatalogSyncGetResponseEnvelopeErrorsCode2012   CatalogSyncGetResponseEnvelopeErrorsCode = 2012
	CatalogSyncGetResponseEnvelopeErrorsCode2013   CatalogSyncGetResponseEnvelopeErrorsCode = 2013
	CatalogSyncGetResponseEnvelopeErrorsCode2014   CatalogSyncGetResponseEnvelopeErrorsCode = 2014
	CatalogSyncGetResponseEnvelopeErrorsCode2015   CatalogSyncGetResponseEnvelopeErrorsCode = 2015
	CatalogSyncGetResponseEnvelopeErrorsCode2016   CatalogSyncGetResponseEnvelopeErrorsCode = 2016
	CatalogSyncGetResponseEnvelopeErrorsCode2017   CatalogSyncGetResponseEnvelopeErrorsCode = 2017
	CatalogSyncGetResponseEnvelopeErrorsCode2018   CatalogSyncGetResponseEnvelopeErrorsCode = 2018
	CatalogSyncGetResponseEnvelopeErrorsCode2019   CatalogSyncGetResponseEnvelopeErrorsCode = 2019
	CatalogSyncGetResponseEnvelopeErrorsCode2020   CatalogSyncGetResponseEnvelopeErrorsCode = 2020
	CatalogSyncGetResponseEnvelopeErrorsCode2021   CatalogSyncGetResponseEnvelopeErrorsCode = 2021
	CatalogSyncGetResponseEnvelopeErrorsCode2022   CatalogSyncGetResponseEnvelopeErrorsCode = 2022
	CatalogSyncGetResponseEnvelopeErrorsCode3001   CatalogSyncGetResponseEnvelopeErrorsCode = 3001
	CatalogSyncGetResponseEnvelopeErrorsCode3002   CatalogSyncGetResponseEnvelopeErrorsCode = 3002
	CatalogSyncGetResponseEnvelopeErrorsCode3003   CatalogSyncGetResponseEnvelopeErrorsCode = 3003
	CatalogSyncGetResponseEnvelopeErrorsCode3004   CatalogSyncGetResponseEnvelopeErrorsCode = 3004
	CatalogSyncGetResponseEnvelopeErrorsCode3005   CatalogSyncGetResponseEnvelopeErrorsCode = 3005
	CatalogSyncGetResponseEnvelopeErrorsCode3006   CatalogSyncGetResponseEnvelopeErrorsCode = 3006
	CatalogSyncGetResponseEnvelopeErrorsCode3007   CatalogSyncGetResponseEnvelopeErrorsCode = 3007
	CatalogSyncGetResponseEnvelopeErrorsCode4001   CatalogSyncGetResponseEnvelopeErrorsCode = 4001
	CatalogSyncGetResponseEnvelopeErrorsCode4002   CatalogSyncGetResponseEnvelopeErrorsCode = 4002
	CatalogSyncGetResponseEnvelopeErrorsCode4003   CatalogSyncGetResponseEnvelopeErrorsCode = 4003
	CatalogSyncGetResponseEnvelopeErrorsCode4004   CatalogSyncGetResponseEnvelopeErrorsCode = 4004
	CatalogSyncGetResponseEnvelopeErrorsCode4005   CatalogSyncGetResponseEnvelopeErrorsCode = 4005
	CatalogSyncGetResponseEnvelopeErrorsCode4006   CatalogSyncGetResponseEnvelopeErrorsCode = 4006
	CatalogSyncGetResponseEnvelopeErrorsCode4007   CatalogSyncGetResponseEnvelopeErrorsCode = 4007
	CatalogSyncGetResponseEnvelopeErrorsCode4008   CatalogSyncGetResponseEnvelopeErrorsCode = 4008
	CatalogSyncGetResponseEnvelopeErrorsCode4009   CatalogSyncGetResponseEnvelopeErrorsCode = 4009
	CatalogSyncGetResponseEnvelopeErrorsCode4010   CatalogSyncGetResponseEnvelopeErrorsCode = 4010
	CatalogSyncGetResponseEnvelopeErrorsCode4011   CatalogSyncGetResponseEnvelopeErrorsCode = 4011
	CatalogSyncGetResponseEnvelopeErrorsCode4012   CatalogSyncGetResponseEnvelopeErrorsCode = 4012
	CatalogSyncGetResponseEnvelopeErrorsCode4013   CatalogSyncGetResponseEnvelopeErrorsCode = 4013
	CatalogSyncGetResponseEnvelopeErrorsCode4014   CatalogSyncGetResponseEnvelopeErrorsCode = 4014
	CatalogSyncGetResponseEnvelopeErrorsCode4015   CatalogSyncGetResponseEnvelopeErrorsCode = 4015
	CatalogSyncGetResponseEnvelopeErrorsCode4016   CatalogSyncGetResponseEnvelopeErrorsCode = 4016
	CatalogSyncGetResponseEnvelopeErrorsCode4017   CatalogSyncGetResponseEnvelopeErrorsCode = 4017
	CatalogSyncGetResponseEnvelopeErrorsCode4018   CatalogSyncGetResponseEnvelopeErrorsCode = 4018
	CatalogSyncGetResponseEnvelopeErrorsCode4019   CatalogSyncGetResponseEnvelopeErrorsCode = 4019
	CatalogSyncGetResponseEnvelopeErrorsCode4020   CatalogSyncGetResponseEnvelopeErrorsCode = 4020
	CatalogSyncGetResponseEnvelopeErrorsCode4021   CatalogSyncGetResponseEnvelopeErrorsCode = 4021
	CatalogSyncGetResponseEnvelopeErrorsCode4022   CatalogSyncGetResponseEnvelopeErrorsCode = 4022
	CatalogSyncGetResponseEnvelopeErrorsCode4023   CatalogSyncGetResponseEnvelopeErrorsCode = 4023
	CatalogSyncGetResponseEnvelopeErrorsCode5001   CatalogSyncGetResponseEnvelopeErrorsCode = 5001
	CatalogSyncGetResponseEnvelopeErrorsCode5002   CatalogSyncGetResponseEnvelopeErrorsCode = 5002
	CatalogSyncGetResponseEnvelopeErrorsCode5003   CatalogSyncGetResponseEnvelopeErrorsCode = 5003
	CatalogSyncGetResponseEnvelopeErrorsCode5004   CatalogSyncGetResponseEnvelopeErrorsCode = 5004
	CatalogSyncGetResponseEnvelopeErrorsCode102000 CatalogSyncGetResponseEnvelopeErrorsCode = 102000
	CatalogSyncGetResponseEnvelopeErrorsCode102001 CatalogSyncGetResponseEnvelopeErrorsCode = 102001
	CatalogSyncGetResponseEnvelopeErrorsCode102002 CatalogSyncGetResponseEnvelopeErrorsCode = 102002
	CatalogSyncGetResponseEnvelopeErrorsCode102003 CatalogSyncGetResponseEnvelopeErrorsCode = 102003
	CatalogSyncGetResponseEnvelopeErrorsCode102004 CatalogSyncGetResponseEnvelopeErrorsCode = 102004
	CatalogSyncGetResponseEnvelopeErrorsCode102005 CatalogSyncGetResponseEnvelopeErrorsCode = 102005
	CatalogSyncGetResponseEnvelopeErrorsCode102006 CatalogSyncGetResponseEnvelopeErrorsCode = 102006
	CatalogSyncGetResponseEnvelopeErrorsCode102007 CatalogSyncGetResponseEnvelopeErrorsCode = 102007
	CatalogSyncGetResponseEnvelopeErrorsCode102008 CatalogSyncGetResponseEnvelopeErrorsCode = 102008
	CatalogSyncGetResponseEnvelopeErrorsCode102009 CatalogSyncGetResponseEnvelopeErrorsCode = 102009
	CatalogSyncGetResponseEnvelopeErrorsCode102010 CatalogSyncGetResponseEnvelopeErrorsCode = 102010
	CatalogSyncGetResponseEnvelopeErrorsCode102011 CatalogSyncGetResponseEnvelopeErrorsCode = 102011
	CatalogSyncGetResponseEnvelopeErrorsCode102012 CatalogSyncGetResponseEnvelopeErrorsCode = 102012
	CatalogSyncGetResponseEnvelopeErrorsCode102013 CatalogSyncGetResponseEnvelopeErrorsCode = 102013
	CatalogSyncGetResponseEnvelopeErrorsCode102014 CatalogSyncGetResponseEnvelopeErrorsCode = 102014
	CatalogSyncGetResponseEnvelopeErrorsCode102015 CatalogSyncGetResponseEnvelopeErrorsCode = 102015
	CatalogSyncGetResponseEnvelopeErrorsCode102016 CatalogSyncGetResponseEnvelopeErrorsCode = 102016
	CatalogSyncGetResponseEnvelopeErrorsCode102017 CatalogSyncGetResponseEnvelopeErrorsCode = 102017
	CatalogSyncGetResponseEnvelopeErrorsCode102018 CatalogSyncGetResponseEnvelopeErrorsCode = 102018
	CatalogSyncGetResponseEnvelopeErrorsCode102019 CatalogSyncGetResponseEnvelopeErrorsCode = 102019
	CatalogSyncGetResponseEnvelopeErrorsCode102020 CatalogSyncGetResponseEnvelopeErrorsCode = 102020
	CatalogSyncGetResponseEnvelopeErrorsCode102021 CatalogSyncGetResponseEnvelopeErrorsCode = 102021
	CatalogSyncGetResponseEnvelopeErrorsCode102022 CatalogSyncGetResponseEnvelopeErrorsCode = 102022
	CatalogSyncGetResponseEnvelopeErrorsCode102023 CatalogSyncGetResponseEnvelopeErrorsCode = 102023
	CatalogSyncGetResponseEnvelopeErrorsCode102024 CatalogSyncGetResponseEnvelopeErrorsCode = 102024
	CatalogSyncGetResponseEnvelopeErrorsCode102025 CatalogSyncGetResponseEnvelopeErrorsCode = 102025
	CatalogSyncGetResponseEnvelopeErrorsCode102026 CatalogSyncGetResponseEnvelopeErrorsCode = 102026
	CatalogSyncGetResponseEnvelopeErrorsCode102027 CatalogSyncGetResponseEnvelopeErrorsCode = 102027
	CatalogSyncGetResponseEnvelopeErrorsCode102028 CatalogSyncGetResponseEnvelopeErrorsCode = 102028
	CatalogSyncGetResponseEnvelopeErrorsCode102029 CatalogSyncGetResponseEnvelopeErrorsCode = 102029
	CatalogSyncGetResponseEnvelopeErrorsCode102030 CatalogSyncGetResponseEnvelopeErrorsCode = 102030
	CatalogSyncGetResponseEnvelopeErrorsCode102031 CatalogSyncGetResponseEnvelopeErrorsCode = 102031
	CatalogSyncGetResponseEnvelopeErrorsCode102032 CatalogSyncGetResponseEnvelopeErrorsCode = 102032
	CatalogSyncGetResponseEnvelopeErrorsCode102033 CatalogSyncGetResponseEnvelopeErrorsCode = 102033
	CatalogSyncGetResponseEnvelopeErrorsCode102034 CatalogSyncGetResponseEnvelopeErrorsCode = 102034
	CatalogSyncGetResponseEnvelopeErrorsCode102035 CatalogSyncGetResponseEnvelopeErrorsCode = 102035
	CatalogSyncGetResponseEnvelopeErrorsCode102036 CatalogSyncGetResponseEnvelopeErrorsCode = 102036
	CatalogSyncGetResponseEnvelopeErrorsCode102037 CatalogSyncGetResponseEnvelopeErrorsCode = 102037
	CatalogSyncGetResponseEnvelopeErrorsCode102038 CatalogSyncGetResponseEnvelopeErrorsCode = 102038
	CatalogSyncGetResponseEnvelopeErrorsCode102039 CatalogSyncGetResponseEnvelopeErrorsCode = 102039
	CatalogSyncGetResponseEnvelopeErrorsCode102040 CatalogSyncGetResponseEnvelopeErrorsCode = 102040
	CatalogSyncGetResponseEnvelopeErrorsCode102041 CatalogSyncGetResponseEnvelopeErrorsCode = 102041
	CatalogSyncGetResponseEnvelopeErrorsCode102042 CatalogSyncGetResponseEnvelopeErrorsCode = 102042
	CatalogSyncGetResponseEnvelopeErrorsCode102043 CatalogSyncGetResponseEnvelopeErrorsCode = 102043
	CatalogSyncGetResponseEnvelopeErrorsCode102044 CatalogSyncGetResponseEnvelopeErrorsCode = 102044
	CatalogSyncGetResponseEnvelopeErrorsCode102045 CatalogSyncGetResponseEnvelopeErrorsCode = 102045
	CatalogSyncGetResponseEnvelopeErrorsCode102046 CatalogSyncGetResponseEnvelopeErrorsCode = 102046
	CatalogSyncGetResponseEnvelopeErrorsCode102047 CatalogSyncGetResponseEnvelopeErrorsCode = 102047
	CatalogSyncGetResponseEnvelopeErrorsCode102048 CatalogSyncGetResponseEnvelopeErrorsCode = 102048
	CatalogSyncGetResponseEnvelopeErrorsCode102049 CatalogSyncGetResponseEnvelopeErrorsCode = 102049
	CatalogSyncGetResponseEnvelopeErrorsCode102050 CatalogSyncGetResponseEnvelopeErrorsCode = 102050
	CatalogSyncGetResponseEnvelopeErrorsCode102051 CatalogSyncGetResponseEnvelopeErrorsCode = 102051
	CatalogSyncGetResponseEnvelopeErrorsCode102052 CatalogSyncGetResponseEnvelopeErrorsCode = 102052
	CatalogSyncGetResponseEnvelopeErrorsCode102053 CatalogSyncGetResponseEnvelopeErrorsCode = 102053
	CatalogSyncGetResponseEnvelopeErrorsCode102054 CatalogSyncGetResponseEnvelopeErrorsCode = 102054
	CatalogSyncGetResponseEnvelopeErrorsCode102055 CatalogSyncGetResponseEnvelopeErrorsCode = 102055
	CatalogSyncGetResponseEnvelopeErrorsCode102056 CatalogSyncGetResponseEnvelopeErrorsCode = 102056
	CatalogSyncGetResponseEnvelopeErrorsCode102057 CatalogSyncGetResponseEnvelopeErrorsCode = 102057
	CatalogSyncGetResponseEnvelopeErrorsCode102058 CatalogSyncGetResponseEnvelopeErrorsCode = 102058
	CatalogSyncGetResponseEnvelopeErrorsCode102059 CatalogSyncGetResponseEnvelopeErrorsCode = 102059
	CatalogSyncGetResponseEnvelopeErrorsCode102060 CatalogSyncGetResponseEnvelopeErrorsCode = 102060
	CatalogSyncGetResponseEnvelopeErrorsCode102061 CatalogSyncGetResponseEnvelopeErrorsCode = 102061
	CatalogSyncGetResponseEnvelopeErrorsCode102062 CatalogSyncGetResponseEnvelopeErrorsCode = 102062
	CatalogSyncGetResponseEnvelopeErrorsCode102063 CatalogSyncGetResponseEnvelopeErrorsCode = 102063
	CatalogSyncGetResponseEnvelopeErrorsCode102064 CatalogSyncGetResponseEnvelopeErrorsCode = 102064
	CatalogSyncGetResponseEnvelopeErrorsCode102065 CatalogSyncGetResponseEnvelopeErrorsCode = 102065
	CatalogSyncGetResponseEnvelopeErrorsCode102066 CatalogSyncGetResponseEnvelopeErrorsCode = 102066
	CatalogSyncGetResponseEnvelopeErrorsCode103001 CatalogSyncGetResponseEnvelopeErrorsCode = 103001
	CatalogSyncGetResponseEnvelopeErrorsCode103002 CatalogSyncGetResponseEnvelopeErrorsCode = 103002
	CatalogSyncGetResponseEnvelopeErrorsCode103003 CatalogSyncGetResponseEnvelopeErrorsCode = 103003
	CatalogSyncGetResponseEnvelopeErrorsCode103004 CatalogSyncGetResponseEnvelopeErrorsCode = 103004
	CatalogSyncGetResponseEnvelopeErrorsCode103005 CatalogSyncGetResponseEnvelopeErrorsCode = 103005
	CatalogSyncGetResponseEnvelopeErrorsCode103006 CatalogSyncGetResponseEnvelopeErrorsCode = 103006
	CatalogSyncGetResponseEnvelopeErrorsCode103007 CatalogSyncGetResponseEnvelopeErrorsCode = 103007
	CatalogSyncGetResponseEnvelopeErrorsCode103008 CatalogSyncGetResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncGetResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncGetResponseEnvelopeErrorsCode1001, CatalogSyncGetResponseEnvelopeErrorsCode1002, CatalogSyncGetResponseEnvelopeErrorsCode1003, CatalogSyncGetResponseEnvelopeErrorsCode1004, CatalogSyncGetResponseEnvelopeErrorsCode1005, CatalogSyncGetResponseEnvelopeErrorsCode1006, CatalogSyncGetResponseEnvelopeErrorsCode1007, CatalogSyncGetResponseEnvelopeErrorsCode1008, CatalogSyncGetResponseEnvelopeErrorsCode1009, CatalogSyncGetResponseEnvelopeErrorsCode1010, CatalogSyncGetResponseEnvelopeErrorsCode1011, CatalogSyncGetResponseEnvelopeErrorsCode1012, CatalogSyncGetResponseEnvelopeErrorsCode1013, CatalogSyncGetResponseEnvelopeErrorsCode1014, CatalogSyncGetResponseEnvelopeErrorsCode1015, CatalogSyncGetResponseEnvelopeErrorsCode1016, CatalogSyncGetResponseEnvelopeErrorsCode1017, CatalogSyncGetResponseEnvelopeErrorsCode2001, CatalogSyncGetResponseEnvelopeErrorsCode2002, CatalogSyncGetResponseEnvelopeErrorsCode2003, CatalogSyncGetResponseEnvelopeErrorsCode2004, CatalogSyncGetResponseEnvelopeErrorsCode2005, CatalogSyncGetResponseEnvelopeErrorsCode2006, CatalogSyncGetResponseEnvelopeErrorsCode2007, CatalogSyncGetResponseEnvelopeErrorsCode2008, CatalogSyncGetResponseEnvelopeErrorsCode2009, CatalogSyncGetResponseEnvelopeErrorsCode2010, CatalogSyncGetResponseEnvelopeErrorsCode2011, CatalogSyncGetResponseEnvelopeErrorsCode2012, CatalogSyncGetResponseEnvelopeErrorsCode2013, CatalogSyncGetResponseEnvelopeErrorsCode2014, CatalogSyncGetResponseEnvelopeErrorsCode2015, CatalogSyncGetResponseEnvelopeErrorsCode2016, CatalogSyncGetResponseEnvelopeErrorsCode2017, CatalogSyncGetResponseEnvelopeErrorsCode2018, CatalogSyncGetResponseEnvelopeErrorsCode2019, CatalogSyncGetResponseEnvelopeErrorsCode2020, CatalogSyncGetResponseEnvelopeErrorsCode2021, CatalogSyncGetResponseEnvelopeErrorsCode2022, CatalogSyncGetResponseEnvelopeErrorsCode3001, CatalogSyncGetResponseEnvelopeErrorsCode3002, CatalogSyncGetResponseEnvelopeErrorsCode3003, CatalogSyncGetResponseEnvelopeErrorsCode3004, CatalogSyncGetResponseEnvelopeErrorsCode3005, CatalogSyncGetResponseEnvelopeErrorsCode3006, CatalogSyncGetResponseEnvelopeErrorsCode3007, CatalogSyncGetResponseEnvelopeErrorsCode4001, CatalogSyncGetResponseEnvelopeErrorsCode4002, CatalogSyncGetResponseEnvelopeErrorsCode4003, CatalogSyncGetResponseEnvelopeErrorsCode4004, CatalogSyncGetResponseEnvelopeErrorsCode4005, CatalogSyncGetResponseEnvelopeErrorsCode4006, CatalogSyncGetResponseEnvelopeErrorsCode4007, CatalogSyncGetResponseEnvelopeErrorsCode4008, CatalogSyncGetResponseEnvelopeErrorsCode4009, CatalogSyncGetResponseEnvelopeErrorsCode4010, CatalogSyncGetResponseEnvelopeErrorsCode4011, CatalogSyncGetResponseEnvelopeErrorsCode4012, CatalogSyncGetResponseEnvelopeErrorsCode4013, CatalogSyncGetResponseEnvelopeErrorsCode4014, CatalogSyncGetResponseEnvelopeErrorsCode4015, CatalogSyncGetResponseEnvelopeErrorsCode4016, CatalogSyncGetResponseEnvelopeErrorsCode4017, CatalogSyncGetResponseEnvelopeErrorsCode4018, CatalogSyncGetResponseEnvelopeErrorsCode4019, CatalogSyncGetResponseEnvelopeErrorsCode4020, CatalogSyncGetResponseEnvelopeErrorsCode4021, CatalogSyncGetResponseEnvelopeErrorsCode4022, CatalogSyncGetResponseEnvelopeErrorsCode4023, CatalogSyncGetResponseEnvelopeErrorsCode5001, CatalogSyncGetResponseEnvelopeErrorsCode5002, CatalogSyncGetResponseEnvelopeErrorsCode5003, CatalogSyncGetResponseEnvelopeErrorsCode5004, CatalogSyncGetResponseEnvelopeErrorsCode102000, CatalogSyncGetResponseEnvelopeErrorsCode102001, CatalogSyncGetResponseEnvelopeErrorsCode102002, CatalogSyncGetResponseEnvelopeErrorsCode102003, CatalogSyncGetResponseEnvelopeErrorsCode102004, CatalogSyncGetResponseEnvelopeErrorsCode102005, CatalogSyncGetResponseEnvelopeErrorsCode102006, CatalogSyncGetResponseEnvelopeErrorsCode102007, CatalogSyncGetResponseEnvelopeErrorsCode102008, CatalogSyncGetResponseEnvelopeErrorsCode102009, CatalogSyncGetResponseEnvelopeErrorsCode102010, CatalogSyncGetResponseEnvelopeErrorsCode102011, CatalogSyncGetResponseEnvelopeErrorsCode102012, CatalogSyncGetResponseEnvelopeErrorsCode102013, CatalogSyncGetResponseEnvelopeErrorsCode102014, CatalogSyncGetResponseEnvelopeErrorsCode102015, CatalogSyncGetResponseEnvelopeErrorsCode102016, CatalogSyncGetResponseEnvelopeErrorsCode102017, CatalogSyncGetResponseEnvelopeErrorsCode102018, CatalogSyncGetResponseEnvelopeErrorsCode102019, CatalogSyncGetResponseEnvelopeErrorsCode102020, CatalogSyncGetResponseEnvelopeErrorsCode102021, CatalogSyncGetResponseEnvelopeErrorsCode102022, CatalogSyncGetResponseEnvelopeErrorsCode102023, CatalogSyncGetResponseEnvelopeErrorsCode102024, CatalogSyncGetResponseEnvelopeErrorsCode102025, CatalogSyncGetResponseEnvelopeErrorsCode102026, CatalogSyncGetResponseEnvelopeErrorsCode102027, CatalogSyncGetResponseEnvelopeErrorsCode102028, CatalogSyncGetResponseEnvelopeErrorsCode102029, CatalogSyncGetResponseEnvelopeErrorsCode102030, CatalogSyncGetResponseEnvelopeErrorsCode102031, CatalogSyncGetResponseEnvelopeErrorsCode102032, CatalogSyncGetResponseEnvelopeErrorsCode102033, CatalogSyncGetResponseEnvelopeErrorsCode102034, CatalogSyncGetResponseEnvelopeErrorsCode102035, CatalogSyncGetResponseEnvelopeErrorsCode102036, CatalogSyncGetResponseEnvelopeErrorsCode102037, CatalogSyncGetResponseEnvelopeErrorsCode102038, CatalogSyncGetResponseEnvelopeErrorsCode102039, CatalogSyncGetResponseEnvelopeErrorsCode102040, CatalogSyncGetResponseEnvelopeErrorsCode102041, CatalogSyncGetResponseEnvelopeErrorsCode102042, CatalogSyncGetResponseEnvelopeErrorsCode102043, CatalogSyncGetResponseEnvelopeErrorsCode102044, CatalogSyncGetResponseEnvelopeErrorsCode102045, CatalogSyncGetResponseEnvelopeErrorsCode102046, CatalogSyncGetResponseEnvelopeErrorsCode102047, CatalogSyncGetResponseEnvelopeErrorsCode102048, CatalogSyncGetResponseEnvelopeErrorsCode102049, CatalogSyncGetResponseEnvelopeErrorsCode102050, CatalogSyncGetResponseEnvelopeErrorsCode102051, CatalogSyncGetResponseEnvelopeErrorsCode102052, CatalogSyncGetResponseEnvelopeErrorsCode102053, CatalogSyncGetResponseEnvelopeErrorsCode102054, CatalogSyncGetResponseEnvelopeErrorsCode102055, CatalogSyncGetResponseEnvelopeErrorsCode102056, CatalogSyncGetResponseEnvelopeErrorsCode102057, CatalogSyncGetResponseEnvelopeErrorsCode102058, CatalogSyncGetResponseEnvelopeErrorsCode102059, CatalogSyncGetResponseEnvelopeErrorsCode102060, CatalogSyncGetResponseEnvelopeErrorsCode102061, CatalogSyncGetResponseEnvelopeErrorsCode102062, CatalogSyncGetResponseEnvelopeErrorsCode102063, CatalogSyncGetResponseEnvelopeErrorsCode102064, CatalogSyncGetResponseEnvelopeErrorsCode102065, CatalogSyncGetResponseEnvelopeErrorsCode102066, CatalogSyncGetResponseEnvelopeErrorsCode103001, CatalogSyncGetResponseEnvelopeErrorsCode103002, CatalogSyncGetResponseEnvelopeErrorsCode103003, CatalogSyncGetResponseEnvelopeErrorsCode103004, CatalogSyncGetResponseEnvelopeErrorsCode103005, CatalogSyncGetResponseEnvelopeErrorsCode103006, CatalogSyncGetResponseEnvelopeErrorsCode103007, CatalogSyncGetResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncGetResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                       `json:"l10n_key"`
	LoggableError string                                       `json:"loggable_error"`
	TemplateData  interface{}                                  `json:"template_data"`
	TraceID       string                                       `json:"trace_id"`
	JSON          catalogSyncGetResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncGetResponseEnvelopeErrorsMetaJSON contains the JSON metadata for the
// struct [CatalogSyncGetResponseEnvelopeErrorsMeta]
type catalogSyncGetResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeErrorsSource struct {
	Parameter           string                                         `json:"parameter"`
	ParameterValueIndex int64                                          `json:"parameter_value_index"`
	Pointer             string                                         `json:"pointer"`
	JSON                catalogSyncGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CatalogSyncGetResponseEnvelopeErrorsSource]
type catalogSyncGetResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeMessages struct {
	Code             CatalogSyncGetResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Meta             CatalogSyncGetResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CatalogSyncGetResponseEnvelopeMessages]
type catalogSyncGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeMessagesCode int64

const (
	CatalogSyncGetResponseEnvelopeMessagesCode1001   CatalogSyncGetResponseEnvelopeMessagesCode = 1001
	CatalogSyncGetResponseEnvelopeMessagesCode1002   CatalogSyncGetResponseEnvelopeMessagesCode = 1002
	CatalogSyncGetResponseEnvelopeMessagesCode1003   CatalogSyncGetResponseEnvelopeMessagesCode = 1003
	CatalogSyncGetResponseEnvelopeMessagesCode1004   CatalogSyncGetResponseEnvelopeMessagesCode = 1004
	CatalogSyncGetResponseEnvelopeMessagesCode1005   CatalogSyncGetResponseEnvelopeMessagesCode = 1005
	CatalogSyncGetResponseEnvelopeMessagesCode1006   CatalogSyncGetResponseEnvelopeMessagesCode = 1006
	CatalogSyncGetResponseEnvelopeMessagesCode1007   CatalogSyncGetResponseEnvelopeMessagesCode = 1007
	CatalogSyncGetResponseEnvelopeMessagesCode1008   CatalogSyncGetResponseEnvelopeMessagesCode = 1008
	CatalogSyncGetResponseEnvelopeMessagesCode1009   CatalogSyncGetResponseEnvelopeMessagesCode = 1009
	CatalogSyncGetResponseEnvelopeMessagesCode1010   CatalogSyncGetResponseEnvelopeMessagesCode = 1010
	CatalogSyncGetResponseEnvelopeMessagesCode1011   CatalogSyncGetResponseEnvelopeMessagesCode = 1011
	CatalogSyncGetResponseEnvelopeMessagesCode1012   CatalogSyncGetResponseEnvelopeMessagesCode = 1012
	CatalogSyncGetResponseEnvelopeMessagesCode1013   CatalogSyncGetResponseEnvelopeMessagesCode = 1013
	CatalogSyncGetResponseEnvelopeMessagesCode1014   CatalogSyncGetResponseEnvelopeMessagesCode = 1014
	CatalogSyncGetResponseEnvelopeMessagesCode1015   CatalogSyncGetResponseEnvelopeMessagesCode = 1015
	CatalogSyncGetResponseEnvelopeMessagesCode1016   CatalogSyncGetResponseEnvelopeMessagesCode = 1016
	CatalogSyncGetResponseEnvelopeMessagesCode1017   CatalogSyncGetResponseEnvelopeMessagesCode = 1017
	CatalogSyncGetResponseEnvelopeMessagesCode2001   CatalogSyncGetResponseEnvelopeMessagesCode = 2001
	CatalogSyncGetResponseEnvelopeMessagesCode2002   CatalogSyncGetResponseEnvelopeMessagesCode = 2002
	CatalogSyncGetResponseEnvelopeMessagesCode2003   CatalogSyncGetResponseEnvelopeMessagesCode = 2003
	CatalogSyncGetResponseEnvelopeMessagesCode2004   CatalogSyncGetResponseEnvelopeMessagesCode = 2004
	CatalogSyncGetResponseEnvelopeMessagesCode2005   CatalogSyncGetResponseEnvelopeMessagesCode = 2005
	CatalogSyncGetResponseEnvelopeMessagesCode2006   CatalogSyncGetResponseEnvelopeMessagesCode = 2006
	CatalogSyncGetResponseEnvelopeMessagesCode2007   CatalogSyncGetResponseEnvelopeMessagesCode = 2007
	CatalogSyncGetResponseEnvelopeMessagesCode2008   CatalogSyncGetResponseEnvelopeMessagesCode = 2008
	CatalogSyncGetResponseEnvelopeMessagesCode2009   CatalogSyncGetResponseEnvelopeMessagesCode = 2009
	CatalogSyncGetResponseEnvelopeMessagesCode2010   CatalogSyncGetResponseEnvelopeMessagesCode = 2010
	CatalogSyncGetResponseEnvelopeMessagesCode2011   CatalogSyncGetResponseEnvelopeMessagesCode = 2011
	CatalogSyncGetResponseEnvelopeMessagesCode2012   CatalogSyncGetResponseEnvelopeMessagesCode = 2012
	CatalogSyncGetResponseEnvelopeMessagesCode2013   CatalogSyncGetResponseEnvelopeMessagesCode = 2013
	CatalogSyncGetResponseEnvelopeMessagesCode2014   CatalogSyncGetResponseEnvelopeMessagesCode = 2014
	CatalogSyncGetResponseEnvelopeMessagesCode2015   CatalogSyncGetResponseEnvelopeMessagesCode = 2015
	CatalogSyncGetResponseEnvelopeMessagesCode2016   CatalogSyncGetResponseEnvelopeMessagesCode = 2016
	CatalogSyncGetResponseEnvelopeMessagesCode2017   CatalogSyncGetResponseEnvelopeMessagesCode = 2017
	CatalogSyncGetResponseEnvelopeMessagesCode2018   CatalogSyncGetResponseEnvelopeMessagesCode = 2018
	CatalogSyncGetResponseEnvelopeMessagesCode2019   CatalogSyncGetResponseEnvelopeMessagesCode = 2019
	CatalogSyncGetResponseEnvelopeMessagesCode2020   CatalogSyncGetResponseEnvelopeMessagesCode = 2020
	CatalogSyncGetResponseEnvelopeMessagesCode2021   CatalogSyncGetResponseEnvelopeMessagesCode = 2021
	CatalogSyncGetResponseEnvelopeMessagesCode2022   CatalogSyncGetResponseEnvelopeMessagesCode = 2022
	CatalogSyncGetResponseEnvelopeMessagesCode3001   CatalogSyncGetResponseEnvelopeMessagesCode = 3001
	CatalogSyncGetResponseEnvelopeMessagesCode3002   CatalogSyncGetResponseEnvelopeMessagesCode = 3002
	CatalogSyncGetResponseEnvelopeMessagesCode3003   CatalogSyncGetResponseEnvelopeMessagesCode = 3003
	CatalogSyncGetResponseEnvelopeMessagesCode3004   CatalogSyncGetResponseEnvelopeMessagesCode = 3004
	CatalogSyncGetResponseEnvelopeMessagesCode3005   CatalogSyncGetResponseEnvelopeMessagesCode = 3005
	CatalogSyncGetResponseEnvelopeMessagesCode3006   CatalogSyncGetResponseEnvelopeMessagesCode = 3006
	CatalogSyncGetResponseEnvelopeMessagesCode3007   CatalogSyncGetResponseEnvelopeMessagesCode = 3007
	CatalogSyncGetResponseEnvelopeMessagesCode4001   CatalogSyncGetResponseEnvelopeMessagesCode = 4001
	CatalogSyncGetResponseEnvelopeMessagesCode4002   CatalogSyncGetResponseEnvelopeMessagesCode = 4002
	CatalogSyncGetResponseEnvelopeMessagesCode4003   CatalogSyncGetResponseEnvelopeMessagesCode = 4003
	CatalogSyncGetResponseEnvelopeMessagesCode4004   CatalogSyncGetResponseEnvelopeMessagesCode = 4004
	CatalogSyncGetResponseEnvelopeMessagesCode4005   CatalogSyncGetResponseEnvelopeMessagesCode = 4005
	CatalogSyncGetResponseEnvelopeMessagesCode4006   CatalogSyncGetResponseEnvelopeMessagesCode = 4006
	CatalogSyncGetResponseEnvelopeMessagesCode4007   CatalogSyncGetResponseEnvelopeMessagesCode = 4007
	CatalogSyncGetResponseEnvelopeMessagesCode4008   CatalogSyncGetResponseEnvelopeMessagesCode = 4008
	CatalogSyncGetResponseEnvelopeMessagesCode4009   CatalogSyncGetResponseEnvelopeMessagesCode = 4009
	CatalogSyncGetResponseEnvelopeMessagesCode4010   CatalogSyncGetResponseEnvelopeMessagesCode = 4010
	CatalogSyncGetResponseEnvelopeMessagesCode4011   CatalogSyncGetResponseEnvelopeMessagesCode = 4011
	CatalogSyncGetResponseEnvelopeMessagesCode4012   CatalogSyncGetResponseEnvelopeMessagesCode = 4012
	CatalogSyncGetResponseEnvelopeMessagesCode4013   CatalogSyncGetResponseEnvelopeMessagesCode = 4013
	CatalogSyncGetResponseEnvelopeMessagesCode4014   CatalogSyncGetResponseEnvelopeMessagesCode = 4014
	CatalogSyncGetResponseEnvelopeMessagesCode4015   CatalogSyncGetResponseEnvelopeMessagesCode = 4015
	CatalogSyncGetResponseEnvelopeMessagesCode4016   CatalogSyncGetResponseEnvelopeMessagesCode = 4016
	CatalogSyncGetResponseEnvelopeMessagesCode4017   CatalogSyncGetResponseEnvelopeMessagesCode = 4017
	CatalogSyncGetResponseEnvelopeMessagesCode4018   CatalogSyncGetResponseEnvelopeMessagesCode = 4018
	CatalogSyncGetResponseEnvelopeMessagesCode4019   CatalogSyncGetResponseEnvelopeMessagesCode = 4019
	CatalogSyncGetResponseEnvelopeMessagesCode4020   CatalogSyncGetResponseEnvelopeMessagesCode = 4020
	CatalogSyncGetResponseEnvelopeMessagesCode4021   CatalogSyncGetResponseEnvelopeMessagesCode = 4021
	CatalogSyncGetResponseEnvelopeMessagesCode4022   CatalogSyncGetResponseEnvelopeMessagesCode = 4022
	CatalogSyncGetResponseEnvelopeMessagesCode4023   CatalogSyncGetResponseEnvelopeMessagesCode = 4023
	CatalogSyncGetResponseEnvelopeMessagesCode5001   CatalogSyncGetResponseEnvelopeMessagesCode = 5001
	CatalogSyncGetResponseEnvelopeMessagesCode5002   CatalogSyncGetResponseEnvelopeMessagesCode = 5002
	CatalogSyncGetResponseEnvelopeMessagesCode5003   CatalogSyncGetResponseEnvelopeMessagesCode = 5003
	CatalogSyncGetResponseEnvelopeMessagesCode5004   CatalogSyncGetResponseEnvelopeMessagesCode = 5004
	CatalogSyncGetResponseEnvelopeMessagesCode102000 CatalogSyncGetResponseEnvelopeMessagesCode = 102000
	CatalogSyncGetResponseEnvelopeMessagesCode102001 CatalogSyncGetResponseEnvelopeMessagesCode = 102001
	CatalogSyncGetResponseEnvelopeMessagesCode102002 CatalogSyncGetResponseEnvelopeMessagesCode = 102002
	CatalogSyncGetResponseEnvelopeMessagesCode102003 CatalogSyncGetResponseEnvelopeMessagesCode = 102003
	CatalogSyncGetResponseEnvelopeMessagesCode102004 CatalogSyncGetResponseEnvelopeMessagesCode = 102004
	CatalogSyncGetResponseEnvelopeMessagesCode102005 CatalogSyncGetResponseEnvelopeMessagesCode = 102005
	CatalogSyncGetResponseEnvelopeMessagesCode102006 CatalogSyncGetResponseEnvelopeMessagesCode = 102006
	CatalogSyncGetResponseEnvelopeMessagesCode102007 CatalogSyncGetResponseEnvelopeMessagesCode = 102007
	CatalogSyncGetResponseEnvelopeMessagesCode102008 CatalogSyncGetResponseEnvelopeMessagesCode = 102008
	CatalogSyncGetResponseEnvelopeMessagesCode102009 CatalogSyncGetResponseEnvelopeMessagesCode = 102009
	CatalogSyncGetResponseEnvelopeMessagesCode102010 CatalogSyncGetResponseEnvelopeMessagesCode = 102010
	CatalogSyncGetResponseEnvelopeMessagesCode102011 CatalogSyncGetResponseEnvelopeMessagesCode = 102011
	CatalogSyncGetResponseEnvelopeMessagesCode102012 CatalogSyncGetResponseEnvelopeMessagesCode = 102012
	CatalogSyncGetResponseEnvelopeMessagesCode102013 CatalogSyncGetResponseEnvelopeMessagesCode = 102013
	CatalogSyncGetResponseEnvelopeMessagesCode102014 CatalogSyncGetResponseEnvelopeMessagesCode = 102014
	CatalogSyncGetResponseEnvelopeMessagesCode102015 CatalogSyncGetResponseEnvelopeMessagesCode = 102015
	CatalogSyncGetResponseEnvelopeMessagesCode102016 CatalogSyncGetResponseEnvelopeMessagesCode = 102016
	CatalogSyncGetResponseEnvelopeMessagesCode102017 CatalogSyncGetResponseEnvelopeMessagesCode = 102017
	CatalogSyncGetResponseEnvelopeMessagesCode102018 CatalogSyncGetResponseEnvelopeMessagesCode = 102018
	CatalogSyncGetResponseEnvelopeMessagesCode102019 CatalogSyncGetResponseEnvelopeMessagesCode = 102019
	CatalogSyncGetResponseEnvelopeMessagesCode102020 CatalogSyncGetResponseEnvelopeMessagesCode = 102020
	CatalogSyncGetResponseEnvelopeMessagesCode102021 CatalogSyncGetResponseEnvelopeMessagesCode = 102021
	CatalogSyncGetResponseEnvelopeMessagesCode102022 CatalogSyncGetResponseEnvelopeMessagesCode = 102022
	CatalogSyncGetResponseEnvelopeMessagesCode102023 CatalogSyncGetResponseEnvelopeMessagesCode = 102023
	CatalogSyncGetResponseEnvelopeMessagesCode102024 CatalogSyncGetResponseEnvelopeMessagesCode = 102024
	CatalogSyncGetResponseEnvelopeMessagesCode102025 CatalogSyncGetResponseEnvelopeMessagesCode = 102025
	CatalogSyncGetResponseEnvelopeMessagesCode102026 CatalogSyncGetResponseEnvelopeMessagesCode = 102026
	CatalogSyncGetResponseEnvelopeMessagesCode102027 CatalogSyncGetResponseEnvelopeMessagesCode = 102027
	CatalogSyncGetResponseEnvelopeMessagesCode102028 CatalogSyncGetResponseEnvelopeMessagesCode = 102028
	CatalogSyncGetResponseEnvelopeMessagesCode102029 CatalogSyncGetResponseEnvelopeMessagesCode = 102029
	CatalogSyncGetResponseEnvelopeMessagesCode102030 CatalogSyncGetResponseEnvelopeMessagesCode = 102030
	CatalogSyncGetResponseEnvelopeMessagesCode102031 CatalogSyncGetResponseEnvelopeMessagesCode = 102031
	CatalogSyncGetResponseEnvelopeMessagesCode102032 CatalogSyncGetResponseEnvelopeMessagesCode = 102032
	CatalogSyncGetResponseEnvelopeMessagesCode102033 CatalogSyncGetResponseEnvelopeMessagesCode = 102033
	CatalogSyncGetResponseEnvelopeMessagesCode102034 CatalogSyncGetResponseEnvelopeMessagesCode = 102034
	CatalogSyncGetResponseEnvelopeMessagesCode102035 CatalogSyncGetResponseEnvelopeMessagesCode = 102035
	CatalogSyncGetResponseEnvelopeMessagesCode102036 CatalogSyncGetResponseEnvelopeMessagesCode = 102036
	CatalogSyncGetResponseEnvelopeMessagesCode102037 CatalogSyncGetResponseEnvelopeMessagesCode = 102037
	CatalogSyncGetResponseEnvelopeMessagesCode102038 CatalogSyncGetResponseEnvelopeMessagesCode = 102038
	CatalogSyncGetResponseEnvelopeMessagesCode102039 CatalogSyncGetResponseEnvelopeMessagesCode = 102039
	CatalogSyncGetResponseEnvelopeMessagesCode102040 CatalogSyncGetResponseEnvelopeMessagesCode = 102040
	CatalogSyncGetResponseEnvelopeMessagesCode102041 CatalogSyncGetResponseEnvelopeMessagesCode = 102041
	CatalogSyncGetResponseEnvelopeMessagesCode102042 CatalogSyncGetResponseEnvelopeMessagesCode = 102042
	CatalogSyncGetResponseEnvelopeMessagesCode102043 CatalogSyncGetResponseEnvelopeMessagesCode = 102043
	CatalogSyncGetResponseEnvelopeMessagesCode102044 CatalogSyncGetResponseEnvelopeMessagesCode = 102044
	CatalogSyncGetResponseEnvelopeMessagesCode102045 CatalogSyncGetResponseEnvelopeMessagesCode = 102045
	CatalogSyncGetResponseEnvelopeMessagesCode102046 CatalogSyncGetResponseEnvelopeMessagesCode = 102046
	CatalogSyncGetResponseEnvelopeMessagesCode102047 CatalogSyncGetResponseEnvelopeMessagesCode = 102047
	CatalogSyncGetResponseEnvelopeMessagesCode102048 CatalogSyncGetResponseEnvelopeMessagesCode = 102048
	CatalogSyncGetResponseEnvelopeMessagesCode102049 CatalogSyncGetResponseEnvelopeMessagesCode = 102049
	CatalogSyncGetResponseEnvelopeMessagesCode102050 CatalogSyncGetResponseEnvelopeMessagesCode = 102050
	CatalogSyncGetResponseEnvelopeMessagesCode102051 CatalogSyncGetResponseEnvelopeMessagesCode = 102051
	CatalogSyncGetResponseEnvelopeMessagesCode102052 CatalogSyncGetResponseEnvelopeMessagesCode = 102052
	CatalogSyncGetResponseEnvelopeMessagesCode102053 CatalogSyncGetResponseEnvelopeMessagesCode = 102053
	CatalogSyncGetResponseEnvelopeMessagesCode102054 CatalogSyncGetResponseEnvelopeMessagesCode = 102054
	CatalogSyncGetResponseEnvelopeMessagesCode102055 CatalogSyncGetResponseEnvelopeMessagesCode = 102055
	CatalogSyncGetResponseEnvelopeMessagesCode102056 CatalogSyncGetResponseEnvelopeMessagesCode = 102056
	CatalogSyncGetResponseEnvelopeMessagesCode102057 CatalogSyncGetResponseEnvelopeMessagesCode = 102057
	CatalogSyncGetResponseEnvelopeMessagesCode102058 CatalogSyncGetResponseEnvelopeMessagesCode = 102058
	CatalogSyncGetResponseEnvelopeMessagesCode102059 CatalogSyncGetResponseEnvelopeMessagesCode = 102059
	CatalogSyncGetResponseEnvelopeMessagesCode102060 CatalogSyncGetResponseEnvelopeMessagesCode = 102060
	CatalogSyncGetResponseEnvelopeMessagesCode102061 CatalogSyncGetResponseEnvelopeMessagesCode = 102061
	CatalogSyncGetResponseEnvelopeMessagesCode102062 CatalogSyncGetResponseEnvelopeMessagesCode = 102062
	CatalogSyncGetResponseEnvelopeMessagesCode102063 CatalogSyncGetResponseEnvelopeMessagesCode = 102063
	CatalogSyncGetResponseEnvelopeMessagesCode102064 CatalogSyncGetResponseEnvelopeMessagesCode = 102064
	CatalogSyncGetResponseEnvelopeMessagesCode102065 CatalogSyncGetResponseEnvelopeMessagesCode = 102065
	CatalogSyncGetResponseEnvelopeMessagesCode102066 CatalogSyncGetResponseEnvelopeMessagesCode = 102066
	CatalogSyncGetResponseEnvelopeMessagesCode103001 CatalogSyncGetResponseEnvelopeMessagesCode = 103001
	CatalogSyncGetResponseEnvelopeMessagesCode103002 CatalogSyncGetResponseEnvelopeMessagesCode = 103002
	CatalogSyncGetResponseEnvelopeMessagesCode103003 CatalogSyncGetResponseEnvelopeMessagesCode = 103003
	CatalogSyncGetResponseEnvelopeMessagesCode103004 CatalogSyncGetResponseEnvelopeMessagesCode = 103004
	CatalogSyncGetResponseEnvelopeMessagesCode103005 CatalogSyncGetResponseEnvelopeMessagesCode = 103005
	CatalogSyncGetResponseEnvelopeMessagesCode103006 CatalogSyncGetResponseEnvelopeMessagesCode = 103006
	CatalogSyncGetResponseEnvelopeMessagesCode103007 CatalogSyncGetResponseEnvelopeMessagesCode = 103007
	CatalogSyncGetResponseEnvelopeMessagesCode103008 CatalogSyncGetResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncGetResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncGetResponseEnvelopeMessagesCode1001, CatalogSyncGetResponseEnvelopeMessagesCode1002, CatalogSyncGetResponseEnvelopeMessagesCode1003, CatalogSyncGetResponseEnvelopeMessagesCode1004, CatalogSyncGetResponseEnvelopeMessagesCode1005, CatalogSyncGetResponseEnvelopeMessagesCode1006, CatalogSyncGetResponseEnvelopeMessagesCode1007, CatalogSyncGetResponseEnvelopeMessagesCode1008, CatalogSyncGetResponseEnvelopeMessagesCode1009, CatalogSyncGetResponseEnvelopeMessagesCode1010, CatalogSyncGetResponseEnvelopeMessagesCode1011, CatalogSyncGetResponseEnvelopeMessagesCode1012, CatalogSyncGetResponseEnvelopeMessagesCode1013, CatalogSyncGetResponseEnvelopeMessagesCode1014, CatalogSyncGetResponseEnvelopeMessagesCode1015, CatalogSyncGetResponseEnvelopeMessagesCode1016, CatalogSyncGetResponseEnvelopeMessagesCode1017, CatalogSyncGetResponseEnvelopeMessagesCode2001, CatalogSyncGetResponseEnvelopeMessagesCode2002, CatalogSyncGetResponseEnvelopeMessagesCode2003, CatalogSyncGetResponseEnvelopeMessagesCode2004, CatalogSyncGetResponseEnvelopeMessagesCode2005, CatalogSyncGetResponseEnvelopeMessagesCode2006, CatalogSyncGetResponseEnvelopeMessagesCode2007, CatalogSyncGetResponseEnvelopeMessagesCode2008, CatalogSyncGetResponseEnvelopeMessagesCode2009, CatalogSyncGetResponseEnvelopeMessagesCode2010, CatalogSyncGetResponseEnvelopeMessagesCode2011, CatalogSyncGetResponseEnvelopeMessagesCode2012, CatalogSyncGetResponseEnvelopeMessagesCode2013, CatalogSyncGetResponseEnvelopeMessagesCode2014, CatalogSyncGetResponseEnvelopeMessagesCode2015, CatalogSyncGetResponseEnvelopeMessagesCode2016, CatalogSyncGetResponseEnvelopeMessagesCode2017, CatalogSyncGetResponseEnvelopeMessagesCode2018, CatalogSyncGetResponseEnvelopeMessagesCode2019, CatalogSyncGetResponseEnvelopeMessagesCode2020, CatalogSyncGetResponseEnvelopeMessagesCode2021, CatalogSyncGetResponseEnvelopeMessagesCode2022, CatalogSyncGetResponseEnvelopeMessagesCode3001, CatalogSyncGetResponseEnvelopeMessagesCode3002, CatalogSyncGetResponseEnvelopeMessagesCode3003, CatalogSyncGetResponseEnvelopeMessagesCode3004, CatalogSyncGetResponseEnvelopeMessagesCode3005, CatalogSyncGetResponseEnvelopeMessagesCode3006, CatalogSyncGetResponseEnvelopeMessagesCode3007, CatalogSyncGetResponseEnvelopeMessagesCode4001, CatalogSyncGetResponseEnvelopeMessagesCode4002, CatalogSyncGetResponseEnvelopeMessagesCode4003, CatalogSyncGetResponseEnvelopeMessagesCode4004, CatalogSyncGetResponseEnvelopeMessagesCode4005, CatalogSyncGetResponseEnvelopeMessagesCode4006, CatalogSyncGetResponseEnvelopeMessagesCode4007, CatalogSyncGetResponseEnvelopeMessagesCode4008, CatalogSyncGetResponseEnvelopeMessagesCode4009, CatalogSyncGetResponseEnvelopeMessagesCode4010, CatalogSyncGetResponseEnvelopeMessagesCode4011, CatalogSyncGetResponseEnvelopeMessagesCode4012, CatalogSyncGetResponseEnvelopeMessagesCode4013, CatalogSyncGetResponseEnvelopeMessagesCode4014, CatalogSyncGetResponseEnvelopeMessagesCode4015, CatalogSyncGetResponseEnvelopeMessagesCode4016, CatalogSyncGetResponseEnvelopeMessagesCode4017, CatalogSyncGetResponseEnvelopeMessagesCode4018, CatalogSyncGetResponseEnvelopeMessagesCode4019, CatalogSyncGetResponseEnvelopeMessagesCode4020, CatalogSyncGetResponseEnvelopeMessagesCode4021, CatalogSyncGetResponseEnvelopeMessagesCode4022, CatalogSyncGetResponseEnvelopeMessagesCode4023, CatalogSyncGetResponseEnvelopeMessagesCode5001, CatalogSyncGetResponseEnvelopeMessagesCode5002, CatalogSyncGetResponseEnvelopeMessagesCode5003, CatalogSyncGetResponseEnvelopeMessagesCode5004, CatalogSyncGetResponseEnvelopeMessagesCode102000, CatalogSyncGetResponseEnvelopeMessagesCode102001, CatalogSyncGetResponseEnvelopeMessagesCode102002, CatalogSyncGetResponseEnvelopeMessagesCode102003, CatalogSyncGetResponseEnvelopeMessagesCode102004, CatalogSyncGetResponseEnvelopeMessagesCode102005, CatalogSyncGetResponseEnvelopeMessagesCode102006, CatalogSyncGetResponseEnvelopeMessagesCode102007, CatalogSyncGetResponseEnvelopeMessagesCode102008, CatalogSyncGetResponseEnvelopeMessagesCode102009, CatalogSyncGetResponseEnvelopeMessagesCode102010, CatalogSyncGetResponseEnvelopeMessagesCode102011, CatalogSyncGetResponseEnvelopeMessagesCode102012, CatalogSyncGetResponseEnvelopeMessagesCode102013, CatalogSyncGetResponseEnvelopeMessagesCode102014, CatalogSyncGetResponseEnvelopeMessagesCode102015, CatalogSyncGetResponseEnvelopeMessagesCode102016, CatalogSyncGetResponseEnvelopeMessagesCode102017, CatalogSyncGetResponseEnvelopeMessagesCode102018, CatalogSyncGetResponseEnvelopeMessagesCode102019, CatalogSyncGetResponseEnvelopeMessagesCode102020, CatalogSyncGetResponseEnvelopeMessagesCode102021, CatalogSyncGetResponseEnvelopeMessagesCode102022, CatalogSyncGetResponseEnvelopeMessagesCode102023, CatalogSyncGetResponseEnvelopeMessagesCode102024, CatalogSyncGetResponseEnvelopeMessagesCode102025, CatalogSyncGetResponseEnvelopeMessagesCode102026, CatalogSyncGetResponseEnvelopeMessagesCode102027, CatalogSyncGetResponseEnvelopeMessagesCode102028, CatalogSyncGetResponseEnvelopeMessagesCode102029, CatalogSyncGetResponseEnvelopeMessagesCode102030, CatalogSyncGetResponseEnvelopeMessagesCode102031, CatalogSyncGetResponseEnvelopeMessagesCode102032, CatalogSyncGetResponseEnvelopeMessagesCode102033, CatalogSyncGetResponseEnvelopeMessagesCode102034, CatalogSyncGetResponseEnvelopeMessagesCode102035, CatalogSyncGetResponseEnvelopeMessagesCode102036, CatalogSyncGetResponseEnvelopeMessagesCode102037, CatalogSyncGetResponseEnvelopeMessagesCode102038, CatalogSyncGetResponseEnvelopeMessagesCode102039, CatalogSyncGetResponseEnvelopeMessagesCode102040, CatalogSyncGetResponseEnvelopeMessagesCode102041, CatalogSyncGetResponseEnvelopeMessagesCode102042, CatalogSyncGetResponseEnvelopeMessagesCode102043, CatalogSyncGetResponseEnvelopeMessagesCode102044, CatalogSyncGetResponseEnvelopeMessagesCode102045, CatalogSyncGetResponseEnvelopeMessagesCode102046, CatalogSyncGetResponseEnvelopeMessagesCode102047, CatalogSyncGetResponseEnvelopeMessagesCode102048, CatalogSyncGetResponseEnvelopeMessagesCode102049, CatalogSyncGetResponseEnvelopeMessagesCode102050, CatalogSyncGetResponseEnvelopeMessagesCode102051, CatalogSyncGetResponseEnvelopeMessagesCode102052, CatalogSyncGetResponseEnvelopeMessagesCode102053, CatalogSyncGetResponseEnvelopeMessagesCode102054, CatalogSyncGetResponseEnvelopeMessagesCode102055, CatalogSyncGetResponseEnvelopeMessagesCode102056, CatalogSyncGetResponseEnvelopeMessagesCode102057, CatalogSyncGetResponseEnvelopeMessagesCode102058, CatalogSyncGetResponseEnvelopeMessagesCode102059, CatalogSyncGetResponseEnvelopeMessagesCode102060, CatalogSyncGetResponseEnvelopeMessagesCode102061, CatalogSyncGetResponseEnvelopeMessagesCode102062, CatalogSyncGetResponseEnvelopeMessagesCode102063, CatalogSyncGetResponseEnvelopeMessagesCode102064, CatalogSyncGetResponseEnvelopeMessagesCode102065, CatalogSyncGetResponseEnvelopeMessagesCode102066, CatalogSyncGetResponseEnvelopeMessagesCode103001, CatalogSyncGetResponseEnvelopeMessagesCode103002, CatalogSyncGetResponseEnvelopeMessagesCode103003, CatalogSyncGetResponseEnvelopeMessagesCode103004, CatalogSyncGetResponseEnvelopeMessagesCode103005, CatalogSyncGetResponseEnvelopeMessagesCode103006, CatalogSyncGetResponseEnvelopeMessagesCode103007, CatalogSyncGetResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncGetResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                         `json:"l10n_key"`
	LoggableError string                                         `json:"loggable_error"`
	TemplateData  interface{}                                    `json:"template_data"`
	TraceID       string                                         `json:"trace_id"`
	JSON          catalogSyncGetResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncGetResponseEnvelopeMessagesMetaJSON contains the JSON metadata for
// the struct [CatalogSyncGetResponseEnvelopeMessagesMeta]
type catalogSyncGetResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncGetResponseEnvelopeMessagesSource struct {
	Parameter           string                                           `json:"parameter"`
	ParameterValueIndex int64                                            `json:"parameter_value_index"`
	Pointer             string                                           `json:"pointer"`
	JSON                catalogSyncGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CatalogSyncGetResponseEnvelopeMessagesSource]
type catalogSyncGetResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CatalogSyncRefreshResponseEnvelope struct {
	Errors   []CatalogSyncRefreshResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CatalogSyncRefreshResponseEnvelopeMessages `json:"messages,required"`
	Result   string                                       `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     catalogSyncRefreshResponseEnvelopeJSON       `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeJSON contains the JSON metadata for the struct
// [CatalogSyncRefreshResponseEnvelope]
type catalogSyncRefreshResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeErrors struct {
	Code             CatalogSyncRefreshResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Meta             CatalogSyncRefreshResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CatalogSyncRefreshResponseEnvelopeErrorsSource `json:"source"`
	JSON             catalogSyncRefreshResponseEnvelopeErrorsJSON   `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CatalogSyncRefreshResponseEnvelopeErrors]
type catalogSyncRefreshResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeErrorsCode int64

const (
	CatalogSyncRefreshResponseEnvelopeErrorsCode1001   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1001
	CatalogSyncRefreshResponseEnvelopeErrorsCode1002   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1002
	CatalogSyncRefreshResponseEnvelopeErrorsCode1003   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1003
	CatalogSyncRefreshResponseEnvelopeErrorsCode1004   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1004
	CatalogSyncRefreshResponseEnvelopeErrorsCode1005   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1005
	CatalogSyncRefreshResponseEnvelopeErrorsCode1006   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1006
	CatalogSyncRefreshResponseEnvelopeErrorsCode1007   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1007
	CatalogSyncRefreshResponseEnvelopeErrorsCode1008   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1008
	CatalogSyncRefreshResponseEnvelopeErrorsCode1009   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1009
	CatalogSyncRefreshResponseEnvelopeErrorsCode1010   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1010
	CatalogSyncRefreshResponseEnvelopeErrorsCode1011   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1011
	CatalogSyncRefreshResponseEnvelopeErrorsCode1012   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1012
	CatalogSyncRefreshResponseEnvelopeErrorsCode1013   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1013
	CatalogSyncRefreshResponseEnvelopeErrorsCode1014   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1014
	CatalogSyncRefreshResponseEnvelopeErrorsCode1015   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1015
	CatalogSyncRefreshResponseEnvelopeErrorsCode1016   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1016
	CatalogSyncRefreshResponseEnvelopeErrorsCode1017   CatalogSyncRefreshResponseEnvelopeErrorsCode = 1017
	CatalogSyncRefreshResponseEnvelopeErrorsCode2001   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2001
	CatalogSyncRefreshResponseEnvelopeErrorsCode2002   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2002
	CatalogSyncRefreshResponseEnvelopeErrorsCode2003   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2003
	CatalogSyncRefreshResponseEnvelopeErrorsCode2004   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2004
	CatalogSyncRefreshResponseEnvelopeErrorsCode2005   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2005
	CatalogSyncRefreshResponseEnvelopeErrorsCode2006   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2006
	CatalogSyncRefreshResponseEnvelopeErrorsCode2007   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2007
	CatalogSyncRefreshResponseEnvelopeErrorsCode2008   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2008
	CatalogSyncRefreshResponseEnvelopeErrorsCode2009   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2009
	CatalogSyncRefreshResponseEnvelopeErrorsCode2010   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2010
	CatalogSyncRefreshResponseEnvelopeErrorsCode2011   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2011
	CatalogSyncRefreshResponseEnvelopeErrorsCode2012   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2012
	CatalogSyncRefreshResponseEnvelopeErrorsCode2013   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2013
	CatalogSyncRefreshResponseEnvelopeErrorsCode2014   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2014
	CatalogSyncRefreshResponseEnvelopeErrorsCode2015   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2015
	CatalogSyncRefreshResponseEnvelopeErrorsCode2016   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2016
	CatalogSyncRefreshResponseEnvelopeErrorsCode2017   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2017
	CatalogSyncRefreshResponseEnvelopeErrorsCode2018   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2018
	CatalogSyncRefreshResponseEnvelopeErrorsCode2019   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2019
	CatalogSyncRefreshResponseEnvelopeErrorsCode2020   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2020
	CatalogSyncRefreshResponseEnvelopeErrorsCode2021   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2021
	CatalogSyncRefreshResponseEnvelopeErrorsCode2022   CatalogSyncRefreshResponseEnvelopeErrorsCode = 2022
	CatalogSyncRefreshResponseEnvelopeErrorsCode3001   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3001
	CatalogSyncRefreshResponseEnvelopeErrorsCode3002   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3002
	CatalogSyncRefreshResponseEnvelopeErrorsCode3003   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3003
	CatalogSyncRefreshResponseEnvelopeErrorsCode3004   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3004
	CatalogSyncRefreshResponseEnvelopeErrorsCode3005   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3005
	CatalogSyncRefreshResponseEnvelopeErrorsCode3006   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3006
	CatalogSyncRefreshResponseEnvelopeErrorsCode3007   CatalogSyncRefreshResponseEnvelopeErrorsCode = 3007
	CatalogSyncRefreshResponseEnvelopeErrorsCode4001   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4001
	CatalogSyncRefreshResponseEnvelopeErrorsCode4002   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4002
	CatalogSyncRefreshResponseEnvelopeErrorsCode4003   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4003
	CatalogSyncRefreshResponseEnvelopeErrorsCode4004   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4004
	CatalogSyncRefreshResponseEnvelopeErrorsCode4005   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4005
	CatalogSyncRefreshResponseEnvelopeErrorsCode4006   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4006
	CatalogSyncRefreshResponseEnvelopeErrorsCode4007   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4007
	CatalogSyncRefreshResponseEnvelopeErrorsCode4008   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4008
	CatalogSyncRefreshResponseEnvelopeErrorsCode4009   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4009
	CatalogSyncRefreshResponseEnvelopeErrorsCode4010   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4010
	CatalogSyncRefreshResponseEnvelopeErrorsCode4011   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4011
	CatalogSyncRefreshResponseEnvelopeErrorsCode4012   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4012
	CatalogSyncRefreshResponseEnvelopeErrorsCode4013   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4013
	CatalogSyncRefreshResponseEnvelopeErrorsCode4014   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4014
	CatalogSyncRefreshResponseEnvelopeErrorsCode4015   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4015
	CatalogSyncRefreshResponseEnvelopeErrorsCode4016   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4016
	CatalogSyncRefreshResponseEnvelopeErrorsCode4017   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4017
	CatalogSyncRefreshResponseEnvelopeErrorsCode4018   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4018
	CatalogSyncRefreshResponseEnvelopeErrorsCode4019   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4019
	CatalogSyncRefreshResponseEnvelopeErrorsCode4020   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4020
	CatalogSyncRefreshResponseEnvelopeErrorsCode4021   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4021
	CatalogSyncRefreshResponseEnvelopeErrorsCode4022   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4022
	CatalogSyncRefreshResponseEnvelopeErrorsCode4023   CatalogSyncRefreshResponseEnvelopeErrorsCode = 4023
	CatalogSyncRefreshResponseEnvelopeErrorsCode5001   CatalogSyncRefreshResponseEnvelopeErrorsCode = 5001
	CatalogSyncRefreshResponseEnvelopeErrorsCode5002   CatalogSyncRefreshResponseEnvelopeErrorsCode = 5002
	CatalogSyncRefreshResponseEnvelopeErrorsCode5003   CatalogSyncRefreshResponseEnvelopeErrorsCode = 5003
	CatalogSyncRefreshResponseEnvelopeErrorsCode5004   CatalogSyncRefreshResponseEnvelopeErrorsCode = 5004
	CatalogSyncRefreshResponseEnvelopeErrorsCode102000 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102000
	CatalogSyncRefreshResponseEnvelopeErrorsCode102001 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102001
	CatalogSyncRefreshResponseEnvelopeErrorsCode102002 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102002
	CatalogSyncRefreshResponseEnvelopeErrorsCode102003 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102003
	CatalogSyncRefreshResponseEnvelopeErrorsCode102004 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102004
	CatalogSyncRefreshResponseEnvelopeErrorsCode102005 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102005
	CatalogSyncRefreshResponseEnvelopeErrorsCode102006 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102006
	CatalogSyncRefreshResponseEnvelopeErrorsCode102007 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102007
	CatalogSyncRefreshResponseEnvelopeErrorsCode102008 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102008
	CatalogSyncRefreshResponseEnvelopeErrorsCode102009 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102009
	CatalogSyncRefreshResponseEnvelopeErrorsCode102010 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102010
	CatalogSyncRefreshResponseEnvelopeErrorsCode102011 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102011
	CatalogSyncRefreshResponseEnvelopeErrorsCode102012 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102012
	CatalogSyncRefreshResponseEnvelopeErrorsCode102013 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102013
	CatalogSyncRefreshResponseEnvelopeErrorsCode102014 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102014
	CatalogSyncRefreshResponseEnvelopeErrorsCode102015 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102015
	CatalogSyncRefreshResponseEnvelopeErrorsCode102016 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102016
	CatalogSyncRefreshResponseEnvelopeErrorsCode102017 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102017
	CatalogSyncRefreshResponseEnvelopeErrorsCode102018 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102018
	CatalogSyncRefreshResponseEnvelopeErrorsCode102019 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102019
	CatalogSyncRefreshResponseEnvelopeErrorsCode102020 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102020
	CatalogSyncRefreshResponseEnvelopeErrorsCode102021 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102021
	CatalogSyncRefreshResponseEnvelopeErrorsCode102022 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102022
	CatalogSyncRefreshResponseEnvelopeErrorsCode102023 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102023
	CatalogSyncRefreshResponseEnvelopeErrorsCode102024 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102024
	CatalogSyncRefreshResponseEnvelopeErrorsCode102025 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102025
	CatalogSyncRefreshResponseEnvelopeErrorsCode102026 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102026
	CatalogSyncRefreshResponseEnvelopeErrorsCode102027 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102027
	CatalogSyncRefreshResponseEnvelopeErrorsCode102028 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102028
	CatalogSyncRefreshResponseEnvelopeErrorsCode102029 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102029
	CatalogSyncRefreshResponseEnvelopeErrorsCode102030 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102030
	CatalogSyncRefreshResponseEnvelopeErrorsCode102031 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102031
	CatalogSyncRefreshResponseEnvelopeErrorsCode102032 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102032
	CatalogSyncRefreshResponseEnvelopeErrorsCode102033 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102033
	CatalogSyncRefreshResponseEnvelopeErrorsCode102034 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102034
	CatalogSyncRefreshResponseEnvelopeErrorsCode102035 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102035
	CatalogSyncRefreshResponseEnvelopeErrorsCode102036 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102036
	CatalogSyncRefreshResponseEnvelopeErrorsCode102037 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102037
	CatalogSyncRefreshResponseEnvelopeErrorsCode102038 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102038
	CatalogSyncRefreshResponseEnvelopeErrorsCode102039 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102039
	CatalogSyncRefreshResponseEnvelopeErrorsCode102040 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102040
	CatalogSyncRefreshResponseEnvelopeErrorsCode102041 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102041
	CatalogSyncRefreshResponseEnvelopeErrorsCode102042 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102042
	CatalogSyncRefreshResponseEnvelopeErrorsCode102043 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102043
	CatalogSyncRefreshResponseEnvelopeErrorsCode102044 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102044
	CatalogSyncRefreshResponseEnvelopeErrorsCode102045 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102045
	CatalogSyncRefreshResponseEnvelopeErrorsCode102046 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102046
	CatalogSyncRefreshResponseEnvelopeErrorsCode102047 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102047
	CatalogSyncRefreshResponseEnvelopeErrorsCode102048 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102048
	CatalogSyncRefreshResponseEnvelopeErrorsCode102049 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102049
	CatalogSyncRefreshResponseEnvelopeErrorsCode102050 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102050
	CatalogSyncRefreshResponseEnvelopeErrorsCode102051 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102051
	CatalogSyncRefreshResponseEnvelopeErrorsCode102052 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102052
	CatalogSyncRefreshResponseEnvelopeErrorsCode102053 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102053
	CatalogSyncRefreshResponseEnvelopeErrorsCode102054 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102054
	CatalogSyncRefreshResponseEnvelopeErrorsCode102055 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102055
	CatalogSyncRefreshResponseEnvelopeErrorsCode102056 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102056
	CatalogSyncRefreshResponseEnvelopeErrorsCode102057 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102057
	CatalogSyncRefreshResponseEnvelopeErrorsCode102058 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102058
	CatalogSyncRefreshResponseEnvelopeErrorsCode102059 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102059
	CatalogSyncRefreshResponseEnvelopeErrorsCode102060 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102060
	CatalogSyncRefreshResponseEnvelopeErrorsCode102061 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102061
	CatalogSyncRefreshResponseEnvelopeErrorsCode102062 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102062
	CatalogSyncRefreshResponseEnvelopeErrorsCode102063 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102063
	CatalogSyncRefreshResponseEnvelopeErrorsCode102064 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102064
	CatalogSyncRefreshResponseEnvelopeErrorsCode102065 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102065
	CatalogSyncRefreshResponseEnvelopeErrorsCode102066 CatalogSyncRefreshResponseEnvelopeErrorsCode = 102066
	CatalogSyncRefreshResponseEnvelopeErrorsCode103001 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103001
	CatalogSyncRefreshResponseEnvelopeErrorsCode103002 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103002
	CatalogSyncRefreshResponseEnvelopeErrorsCode103003 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103003
	CatalogSyncRefreshResponseEnvelopeErrorsCode103004 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103004
	CatalogSyncRefreshResponseEnvelopeErrorsCode103005 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103005
	CatalogSyncRefreshResponseEnvelopeErrorsCode103006 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103006
	CatalogSyncRefreshResponseEnvelopeErrorsCode103007 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103007
	CatalogSyncRefreshResponseEnvelopeErrorsCode103008 CatalogSyncRefreshResponseEnvelopeErrorsCode = 103008
)

func (r CatalogSyncRefreshResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CatalogSyncRefreshResponseEnvelopeErrorsCode1001, CatalogSyncRefreshResponseEnvelopeErrorsCode1002, CatalogSyncRefreshResponseEnvelopeErrorsCode1003, CatalogSyncRefreshResponseEnvelopeErrorsCode1004, CatalogSyncRefreshResponseEnvelopeErrorsCode1005, CatalogSyncRefreshResponseEnvelopeErrorsCode1006, CatalogSyncRefreshResponseEnvelopeErrorsCode1007, CatalogSyncRefreshResponseEnvelopeErrorsCode1008, CatalogSyncRefreshResponseEnvelopeErrorsCode1009, CatalogSyncRefreshResponseEnvelopeErrorsCode1010, CatalogSyncRefreshResponseEnvelopeErrorsCode1011, CatalogSyncRefreshResponseEnvelopeErrorsCode1012, CatalogSyncRefreshResponseEnvelopeErrorsCode1013, CatalogSyncRefreshResponseEnvelopeErrorsCode1014, CatalogSyncRefreshResponseEnvelopeErrorsCode1015, CatalogSyncRefreshResponseEnvelopeErrorsCode1016, CatalogSyncRefreshResponseEnvelopeErrorsCode1017, CatalogSyncRefreshResponseEnvelopeErrorsCode2001, CatalogSyncRefreshResponseEnvelopeErrorsCode2002, CatalogSyncRefreshResponseEnvelopeErrorsCode2003, CatalogSyncRefreshResponseEnvelopeErrorsCode2004, CatalogSyncRefreshResponseEnvelopeErrorsCode2005, CatalogSyncRefreshResponseEnvelopeErrorsCode2006, CatalogSyncRefreshResponseEnvelopeErrorsCode2007, CatalogSyncRefreshResponseEnvelopeErrorsCode2008, CatalogSyncRefreshResponseEnvelopeErrorsCode2009, CatalogSyncRefreshResponseEnvelopeErrorsCode2010, CatalogSyncRefreshResponseEnvelopeErrorsCode2011, CatalogSyncRefreshResponseEnvelopeErrorsCode2012, CatalogSyncRefreshResponseEnvelopeErrorsCode2013, CatalogSyncRefreshResponseEnvelopeErrorsCode2014, CatalogSyncRefreshResponseEnvelopeErrorsCode2015, CatalogSyncRefreshResponseEnvelopeErrorsCode2016, CatalogSyncRefreshResponseEnvelopeErrorsCode2017, CatalogSyncRefreshResponseEnvelopeErrorsCode2018, CatalogSyncRefreshResponseEnvelopeErrorsCode2019, CatalogSyncRefreshResponseEnvelopeErrorsCode2020, CatalogSyncRefreshResponseEnvelopeErrorsCode2021, CatalogSyncRefreshResponseEnvelopeErrorsCode2022, CatalogSyncRefreshResponseEnvelopeErrorsCode3001, CatalogSyncRefreshResponseEnvelopeErrorsCode3002, CatalogSyncRefreshResponseEnvelopeErrorsCode3003, CatalogSyncRefreshResponseEnvelopeErrorsCode3004, CatalogSyncRefreshResponseEnvelopeErrorsCode3005, CatalogSyncRefreshResponseEnvelopeErrorsCode3006, CatalogSyncRefreshResponseEnvelopeErrorsCode3007, CatalogSyncRefreshResponseEnvelopeErrorsCode4001, CatalogSyncRefreshResponseEnvelopeErrorsCode4002, CatalogSyncRefreshResponseEnvelopeErrorsCode4003, CatalogSyncRefreshResponseEnvelopeErrorsCode4004, CatalogSyncRefreshResponseEnvelopeErrorsCode4005, CatalogSyncRefreshResponseEnvelopeErrorsCode4006, CatalogSyncRefreshResponseEnvelopeErrorsCode4007, CatalogSyncRefreshResponseEnvelopeErrorsCode4008, CatalogSyncRefreshResponseEnvelopeErrorsCode4009, CatalogSyncRefreshResponseEnvelopeErrorsCode4010, CatalogSyncRefreshResponseEnvelopeErrorsCode4011, CatalogSyncRefreshResponseEnvelopeErrorsCode4012, CatalogSyncRefreshResponseEnvelopeErrorsCode4013, CatalogSyncRefreshResponseEnvelopeErrorsCode4014, CatalogSyncRefreshResponseEnvelopeErrorsCode4015, CatalogSyncRefreshResponseEnvelopeErrorsCode4016, CatalogSyncRefreshResponseEnvelopeErrorsCode4017, CatalogSyncRefreshResponseEnvelopeErrorsCode4018, CatalogSyncRefreshResponseEnvelopeErrorsCode4019, CatalogSyncRefreshResponseEnvelopeErrorsCode4020, CatalogSyncRefreshResponseEnvelopeErrorsCode4021, CatalogSyncRefreshResponseEnvelopeErrorsCode4022, CatalogSyncRefreshResponseEnvelopeErrorsCode4023, CatalogSyncRefreshResponseEnvelopeErrorsCode5001, CatalogSyncRefreshResponseEnvelopeErrorsCode5002, CatalogSyncRefreshResponseEnvelopeErrorsCode5003, CatalogSyncRefreshResponseEnvelopeErrorsCode5004, CatalogSyncRefreshResponseEnvelopeErrorsCode102000, CatalogSyncRefreshResponseEnvelopeErrorsCode102001, CatalogSyncRefreshResponseEnvelopeErrorsCode102002, CatalogSyncRefreshResponseEnvelopeErrorsCode102003, CatalogSyncRefreshResponseEnvelopeErrorsCode102004, CatalogSyncRefreshResponseEnvelopeErrorsCode102005, CatalogSyncRefreshResponseEnvelopeErrorsCode102006, CatalogSyncRefreshResponseEnvelopeErrorsCode102007, CatalogSyncRefreshResponseEnvelopeErrorsCode102008, CatalogSyncRefreshResponseEnvelopeErrorsCode102009, CatalogSyncRefreshResponseEnvelopeErrorsCode102010, CatalogSyncRefreshResponseEnvelopeErrorsCode102011, CatalogSyncRefreshResponseEnvelopeErrorsCode102012, CatalogSyncRefreshResponseEnvelopeErrorsCode102013, CatalogSyncRefreshResponseEnvelopeErrorsCode102014, CatalogSyncRefreshResponseEnvelopeErrorsCode102015, CatalogSyncRefreshResponseEnvelopeErrorsCode102016, CatalogSyncRefreshResponseEnvelopeErrorsCode102017, CatalogSyncRefreshResponseEnvelopeErrorsCode102018, CatalogSyncRefreshResponseEnvelopeErrorsCode102019, CatalogSyncRefreshResponseEnvelopeErrorsCode102020, CatalogSyncRefreshResponseEnvelopeErrorsCode102021, CatalogSyncRefreshResponseEnvelopeErrorsCode102022, CatalogSyncRefreshResponseEnvelopeErrorsCode102023, CatalogSyncRefreshResponseEnvelopeErrorsCode102024, CatalogSyncRefreshResponseEnvelopeErrorsCode102025, CatalogSyncRefreshResponseEnvelopeErrorsCode102026, CatalogSyncRefreshResponseEnvelopeErrorsCode102027, CatalogSyncRefreshResponseEnvelopeErrorsCode102028, CatalogSyncRefreshResponseEnvelopeErrorsCode102029, CatalogSyncRefreshResponseEnvelopeErrorsCode102030, CatalogSyncRefreshResponseEnvelopeErrorsCode102031, CatalogSyncRefreshResponseEnvelopeErrorsCode102032, CatalogSyncRefreshResponseEnvelopeErrorsCode102033, CatalogSyncRefreshResponseEnvelopeErrorsCode102034, CatalogSyncRefreshResponseEnvelopeErrorsCode102035, CatalogSyncRefreshResponseEnvelopeErrorsCode102036, CatalogSyncRefreshResponseEnvelopeErrorsCode102037, CatalogSyncRefreshResponseEnvelopeErrorsCode102038, CatalogSyncRefreshResponseEnvelopeErrorsCode102039, CatalogSyncRefreshResponseEnvelopeErrorsCode102040, CatalogSyncRefreshResponseEnvelopeErrorsCode102041, CatalogSyncRefreshResponseEnvelopeErrorsCode102042, CatalogSyncRefreshResponseEnvelopeErrorsCode102043, CatalogSyncRefreshResponseEnvelopeErrorsCode102044, CatalogSyncRefreshResponseEnvelopeErrorsCode102045, CatalogSyncRefreshResponseEnvelopeErrorsCode102046, CatalogSyncRefreshResponseEnvelopeErrorsCode102047, CatalogSyncRefreshResponseEnvelopeErrorsCode102048, CatalogSyncRefreshResponseEnvelopeErrorsCode102049, CatalogSyncRefreshResponseEnvelopeErrorsCode102050, CatalogSyncRefreshResponseEnvelopeErrorsCode102051, CatalogSyncRefreshResponseEnvelopeErrorsCode102052, CatalogSyncRefreshResponseEnvelopeErrorsCode102053, CatalogSyncRefreshResponseEnvelopeErrorsCode102054, CatalogSyncRefreshResponseEnvelopeErrorsCode102055, CatalogSyncRefreshResponseEnvelopeErrorsCode102056, CatalogSyncRefreshResponseEnvelopeErrorsCode102057, CatalogSyncRefreshResponseEnvelopeErrorsCode102058, CatalogSyncRefreshResponseEnvelopeErrorsCode102059, CatalogSyncRefreshResponseEnvelopeErrorsCode102060, CatalogSyncRefreshResponseEnvelopeErrorsCode102061, CatalogSyncRefreshResponseEnvelopeErrorsCode102062, CatalogSyncRefreshResponseEnvelopeErrorsCode102063, CatalogSyncRefreshResponseEnvelopeErrorsCode102064, CatalogSyncRefreshResponseEnvelopeErrorsCode102065, CatalogSyncRefreshResponseEnvelopeErrorsCode102066, CatalogSyncRefreshResponseEnvelopeErrorsCode103001, CatalogSyncRefreshResponseEnvelopeErrorsCode103002, CatalogSyncRefreshResponseEnvelopeErrorsCode103003, CatalogSyncRefreshResponseEnvelopeErrorsCode103004, CatalogSyncRefreshResponseEnvelopeErrorsCode103005, CatalogSyncRefreshResponseEnvelopeErrorsCode103006, CatalogSyncRefreshResponseEnvelopeErrorsCode103007, CatalogSyncRefreshResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CatalogSyncRefreshResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                           `json:"l10n_key"`
	LoggableError string                                           `json:"loggable_error"`
	TemplateData  interface{}                                      `json:"template_data"`
	TraceID       string                                           `json:"trace_id"`
	JSON          catalogSyncRefreshResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeErrorsMetaJSON contains the JSON metadata for
// the struct [CatalogSyncRefreshResponseEnvelopeErrorsMeta]
type catalogSyncRefreshResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeErrorsSource struct {
	Parameter           string                                             `json:"parameter"`
	ParameterValueIndex int64                                              `json:"parameter_value_index"`
	Pointer             string                                             `json:"pointer"`
	JSON                catalogSyncRefreshResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CatalogSyncRefreshResponseEnvelopeErrorsSource]
type catalogSyncRefreshResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeMessages struct {
	Code             CatalogSyncRefreshResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Meta             CatalogSyncRefreshResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CatalogSyncRefreshResponseEnvelopeMessagesSource `json:"source"`
	JSON             catalogSyncRefreshResponseEnvelopeMessagesJSON   `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CatalogSyncRefreshResponseEnvelopeMessages]
type catalogSyncRefreshResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeMessagesCode int64

const (
	CatalogSyncRefreshResponseEnvelopeMessagesCode1001   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1001
	CatalogSyncRefreshResponseEnvelopeMessagesCode1002   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1002
	CatalogSyncRefreshResponseEnvelopeMessagesCode1003   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1003
	CatalogSyncRefreshResponseEnvelopeMessagesCode1004   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1004
	CatalogSyncRefreshResponseEnvelopeMessagesCode1005   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1005
	CatalogSyncRefreshResponseEnvelopeMessagesCode1006   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1006
	CatalogSyncRefreshResponseEnvelopeMessagesCode1007   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1007
	CatalogSyncRefreshResponseEnvelopeMessagesCode1008   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1008
	CatalogSyncRefreshResponseEnvelopeMessagesCode1009   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1009
	CatalogSyncRefreshResponseEnvelopeMessagesCode1010   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1010
	CatalogSyncRefreshResponseEnvelopeMessagesCode1011   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1011
	CatalogSyncRefreshResponseEnvelopeMessagesCode1012   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1012
	CatalogSyncRefreshResponseEnvelopeMessagesCode1013   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1013
	CatalogSyncRefreshResponseEnvelopeMessagesCode1014   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1014
	CatalogSyncRefreshResponseEnvelopeMessagesCode1015   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1015
	CatalogSyncRefreshResponseEnvelopeMessagesCode1016   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1016
	CatalogSyncRefreshResponseEnvelopeMessagesCode1017   CatalogSyncRefreshResponseEnvelopeMessagesCode = 1017
	CatalogSyncRefreshResponseEnvelopeMessagesCode2001   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2001
	CatalogSyncRefreshResponseEnvelopeMessagesCode2002   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2002
	CatalogSyncRefreshResponseEnvelopeMessagesCode2003   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2003
	CatalogSyncRefreshResponseEnvelopeMessagesCode2004   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2004
	CatalogSyncRefreshResponseEnvelopeMessagesCode2005   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2005
	CatalogSyncRefreshResponseEnvelopeMessagesCode2006   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2006
	CatalogSyncRefreshResponseEnvelopeMessagesCode2007   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2007
	CatalogSyncRefreshResponseEnvelopeMessagesCode2008   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2008
	CatalogSyncRefreshResponseEnvelopeMessagesCode2009   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2009
	CatalogSyncRefreshResponseEnvelopeMessagesCode2010   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2010
	CatalogSyncRefreshResponseEnvelopeMessagesCode2011   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2011
	CatalogSyncRefreshResponseEnvelopeMessagesCode2012   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2012
	CatalogSyncRefreshResponseEnvelopeMessagesCode2013   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2013
	CatalogSyncRefreshResponseEnvelopeMessagesCode2014   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2014
	CatalogSyncRefreshResponseEnvelopeMessagesCode2015   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2015
	CatalogSyncRefreshResponseEnvelopeMessagesCode2016   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2016
	CatalogSyncRefreshResponseEnvelopeMessagesCode2017   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2017
	CatalogSyncRefreshResponseEnvelopeMessagesCode2018   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2018
	CatalogSyncRefreshResponseEnvelopeMessagesCode2019   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2019
	CatalogSyncRefreshResponseEnvelopeMessagesCode2020   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2020
	CatalogSyncRefreshResponseEnvelopeMessagesCode2021   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2021
	CatalogSyncRefreshResponseEnvelopeMessagesCode2022   CatalogSyncRefreshResponseEnvelopeMessagesCode = 2022
	CatalogSyncRefreshResponseEnvelopeMessagesCode3001   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3001
	CatalogSyncRefreshResponseEnvelopeMessagesCode3002   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3002
	CatalogSyncRefreshResponseEnvelopeMessagesCode3003   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3003
	CatalogSyncRefreshResponseEnvelopeMessagesCode3004   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3004
	CatalogSyncRefreshResponseEnvelopeMessagesCode3005   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3005
	CatalogSyncRefreshResponseEnvelopeMessagesCode3006   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3006
	CatalogSyncRefreshResponseEnvelopeMessagesCode3007   CatalogSyncRefreshResponseEnvelopeMessagesCode = 3007
	CatalogSyncRefreshResponseEnvelopeMessagesCode4001   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4001
	CatalogSyncRefreshResponseEnvelopeMessagesCode4002   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4002
	CatalogSyncRefreshResponseEnvelopeMessagesCode4003   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4003
	CatalogSyncRefreshResponseEnvelopeMessagesCode4004   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4004
	CatalogSyncRefreshResponseEnvelopeMessagesCode4005   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4005
	CatalogSyncRefreshResponseEnvelopeMessagesCode4006   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4006
	CatalogSyncRefreshResponseEnvelopeMessagesCode4007   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4007
	CatalogSyncRefreshResponseEnvelopeMessagesCode4008   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4008
	CatalogSyncRefreshResponseEnvelopeMessagesCode4009   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4009
	CatalogSyncRefreshResponseEnvelopeMessagesCode4010   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4010
	CatalogSyncRefreshResponseEnvelopeMessagesCode4011   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4011
	CatalogSyncRefreshResponseEnvelopeMessagesCode4012   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4012
	CatalogSyncRefreshResponseEnvelopeMessagesCode4013   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4013
	CatalogSyncRefreshResponseEnvelopeMessagesCode4014   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4014
	CatalogSyncRefreshResponseEnvelopeMessagesCode4015   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4015
	CatalogSyncRefreshResponseEnvelopeMessagesCode4016   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4016
	CatalogSyncRefreshResponseEnvelopeMessagesCode4017   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4017
	CatalogSyncRefreshResponseEnvelopeMessagesCode4018   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4018
	CatalogSyncRefreshResponseEnvelopeMessagesCode4019   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4019
	CatalogSyncRefreshResponseEnvelopeMessagesCode4020   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4020
	CatalogSyncRefreshResponseEnvelopeMessagesCode4021   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4021
	CatalogSyncRefreshResponseEnvelopeMessagesCode4022   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4022
	CatalogSyncRefreshResponseEnvelopeMessagesCode4023   CatalogSyncRefreshResponseEnvelopeMessagesCode = 4023
	CatalogSyncRefreshResponseEnvelopeMessagesCode5001   CatalogSyncRefreshResponseEnvelopeMessagesCode = 5001
	CatalogSyncRefreshResponseEnvelopeMessagesCode5002   CatalogSyncRefreshResponseEnvelopeMessagesCode = 5002
	CatalogSyncRefreshResponseEnvelopeMessagesCode5003   CatalogSyncRefreshResponseEnvelopeMessagesCode = 5003
	CatalogSyncRefreshResponseEnvelopeMessagesCode5004   CatalogSyncRefreshResponseEnvelopeMessagesCode = 5004
	CatalogSyncRefreshResponseEnvelopeMessagesCode102000 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102000
	CatalogSyncRefreshResponseEnvelopeMessagesCode102001 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102001
	CatalogSyncRefreshResponseEnvelopeMessagesCode102002 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102002
	CatalogSyncRefreshResponseEnvelopeMessagesCode102003 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102003
	CatalogSyncRefreshResponseEnvelopeMessagesCode102004 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102004
	CatalogSyncRefreshResponseEnvelopeMessagesCode102005 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102005
	CatalogSyncRefreshResponseEnvelopeMessagesCode102006 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102006
	CatalogSyncRefreshResponseEnvelopeMessagesCode102007 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102007
	CatalogSyncRefreshResponseEnvelopeMessagesCode102008 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102008
	CatalogSyncRefreshResponseEnvelopeMessagesCode102009 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102009
	CatalogSyncRefreshResponseEnvelopeMessagesCode102010 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102010
	CatalogSyncRefreshResponseEnvelopeMessagesCode102011 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102011
	CatalogSyncRefreshResponseEnvelopeMessagesCode102012 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102012
	CatalogSyncRefreshResponseEnvelopeMessagesCode102013 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102013
	CatalogSyncRefreshResponseEnvelopeMessagesCode102014 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102014
	CatalogSyncRefreshResponseEnvelopeMessagesCode102015 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102015
	CatalogSyncRefreshResponseEnvelopeMessagesCode102016 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102016
	CatalogSyncRefreshResponseEnvelopeMessagesCode102017 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102017
	CatalogSyncRefreshResponseEnvelopeMessagesCode102018 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102018
	CatalogSyncRefreshResponseEnvelopeMessagesCode102019 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102019
	CatalogSyncRefreshResponseEnvelopeMessagesCode102020 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102020
	CatalogSyncRefreshResponseEnvelopeMessagesCode102021 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102021
	CatalogSyncRefreshResponseEnvelopeMessagesCode102022 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102022
	CatalogSyncRefreshResponseEnvelopeMessagesCode102023 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102023
	CatalogSyncRefreshResponseEnvelopeMessagesCode102024 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102024
	CatalogSyncRefreshResponseEnvelopeMessagesCode102025 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102025
	CatalogSyncRefreshResponseEnvelopeMessagesCode102026 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102026
	CatalogSyncRefreshResponseEnvelopeMessagesCode102027 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102027
	CatalogSyncRefreshResponseEnvelopeMessagesCode102028 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102028
	CatalogSyncRefreshResponseEnvelopeMessagesCode102029 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102029
	CatalogSyncRefreshResponseEnvelopeMessagesCode102030 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102030
	CatalogSyncRefreshResponseEnvelopeMessagesCode102031 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102031
	CatalogSyncRefreshResponseEnvelopeMessagesCode102032 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102032
	CatalogSyncRefreshResponseEnvelopeMessagesCode102033 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102033
	CatalogSyncRefreshResponseEnvelopeMessagesCode102034 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102034
	CatalogSyncRefreshResponseEnvelopeMessagesCode102035 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102035
	CatalogSyncRefreshResponseEnvelopeMessagesCode102036 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102036
	CatalogSyncRefreshResponseEnvelopeMessagesCode102037 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102037
	CatalogSyncRefreshResponseEnvelopeMessagesCode102038 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102038
	CatalogSyncRefreshResponseEnvelopeMessagesCode102039 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102039
	CatalogSyncRefreshResponseEnvelopeMessagesCode102040 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102040
	CatalogSyncRefreshResponseEnvelopeMessagesCode102041 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102041
	CatalogSyncRefreshResponseEnvelopeMessagesCode102042 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102042
	CatalogSyncRefreshResponseEnvelopeMessagesCode102043 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102043
	CatalogSyncRefreshResponseEnvelopeMessagesCode102044 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102044
	CatalogSyncRefreshResponseEnvelopeMessagesCode102045 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102045
	CatalogSyncRefreshResponseEnvelopeMessagesCode102046 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102046
	CatalogSyncRefreshResponseEnvelopeMessagesCode102047 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102047
	CatalogSyncRefreshResponseEnvelopeMessagesCode102048 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102048
	CatalogSyncRefreshResponseEnvelopeMessagesCode102049 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102049
	CatalogSyncRefreshResponseEnvelopeMessagesCode102050 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102050
	CatalogSyncRefreshResponseEnvelopeMessagesCode102051 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102051
	CatalogSyncRefreshResponseEnvelopeMessagesCode102052 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102052
	CatalogSyncRefreshResponseEnvelopeMessagesCode102053 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102053
	CatalogSyncRefreshResponseEnvelopeMessagesCode102054 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102054
	CatalogSyncRefreshResponseEnvelopeMessagesCode102055 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102055
	CatalogSyncRefreshResponseEnvelopeMessagesCode102056 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102056
	CatalogSyncRefreshResponseEnvelopeMessagesCode102057 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102057
	CatalogSyncRefreshResponseEnvelopeMessagesCode102058 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102058
	CatalogSyncRefreshResponseEnvelopeMessagesCode102059 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102059
	CatalogSyncRefreshResponseEnvelopeMessagesCode102060 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102060
	CatalogSyncRefreshResponseEnvelopeMessagesCode102061 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102061
	CatalogSyncRefreshResponseEnvelopeMessagesCode102062 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102062
	CatalogSyncRefreshResponseEnvelopeMessagesCode102063 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102063
	CatalogSyncRefreshResponseEnvelopeMessagesCode102064 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102064
	CatalogSyncRefreshResponseEnvelopeMessagesCode102065 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102065
	CatalogSyncRefreshResponseEnvelopeMessagesCode102066 CatalogSyncRefreshResponseEnvelopeMessagesCode = 102066
	CatalogSyncRefreshResponseEnvelopeMessagesCode103001 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103001
	CatalogSyncRefreshResponseEnvelopeMessagesCode103002 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103002
	CatalogSyncRefreshResponseEnvelopeMessagesCode103003 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103003
	CatalogSyncRefreshResponseEnvelopeMessagesCode103004 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103004
	CatalogSyncRefreshResponseEnvelopeMessagesCode103005 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103005
	CatalogSyncRefreshResponseEnvelopeMessagesCode103006 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103006
	CatalogSyncRefreshResponseEnvelopeMessagesCode103007 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103007
	CatalogSyncRefreshResponseEnvelopeMessagesCode103008 CatalogSyncRefreshResponseEnvelopeMessagesCode = 103008
)

func (r CatalogSyncRefreshResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CatalogSyncRefreshResponseEnvelopeMessagesCode1001, CatalogSyncRefreshResponseEnvelopeMessagesCode1002, CatalogSyncRefreshResponseEnvelopeMessagesCode1003, CatalogSyncRefreshResponseEnvelopeMessagesCode1004, CatalogSyncRefreshResponseEnvelopeMessagesCode1005, CatalogSyncRefreshResponseEnvelopeMessagesCode1006, CatalogSyncRefreshResponseEnvelopeMessagesCode1007, CatalogSyncRefreshResponseEnvelopeMessagesCode1008, CatalogSyncRefreshResponseEnvelopeMessagesCode1009, CatalogSyncRefreshResponseEnvelopeMessagesCode1010, CatalogSyncRefreshResponseEnvelopeMessagesCode1011, CatalogSyncRefreshResponseEnvelopeMessagesCode1012, CatalogSyncRefreshResponseEnvelopeMessagesCode1013, CatalogSyncRefreshResponseEnvelopeMessagesCode1014, CatalogSyncRefreshResponseEnvelopeMessagesCode1015, CatalogSyncRefreshResponseEnvelopeMessagesCode1016, CatalogSyncRefreshResponseEnvelopeMessagesCode1017, CatalogSyncRefreshResponseEnvelopeMessagesCode2001, CatalogSyncRefreshResponseEnvelopeMessagesCode2002, CatalogSyncRefreshResponseEnvelopeMessagesCode2003, CatalogSyncRefreshResponseEnvelopeMessagesCode2004, CatalogSyncRefreshResponseEnvelopeMessagesCode2005, CatalogSyncRefreshResponseEnvelopeMessagesCode2006, CatalogSyncRefreshResponseEnvelopeMessagesCode2007, CatalogSyncRefreshResponseEnvelopeMessagesCode2008, CatalogSyncRefreshResponseEnvelopeMessagesCode2009, CatalogSyncRefreshResponseEnvelopeMessagesCode2010, CatalogSyncRefreshResponseEnvelopeMessagesCode2011, CatalogSyncRefreshResponseEnvelopeMessagesCode2012, CatalogSyncRefreshResponseEnvelopeMessagesCode2013, CatalogSyncRefreshResponseEnvelopeMessagesCode2014, CatalogSyncRefreshResponseEnvelopeMessagesCode2015, CatalogSyncRefreshResponseEnvelopeMessagesCode2016, CatalogSyncRefreshResponseEnvelopeMessagesCode2017, CatalogSyncRefreshResponseEnvelopeMessagesCode2018, CatalogSyncRefreshResponseEnvelopeMessagesCode2019, CatalogSyncRefreshResponseEnvelopeMessagesCode2020, CatalogSyncRefreshResponseEnvelopeMessagesCode2021, CatalogSyncRefreshResponseEnvelopeMessagesCode2022, CatalogSyncRefreshResponseEnvelopeMessagesCode3001, CatalogSyncRefreshResponseEnvelopeMessagesCode3002, CatalogSyncRefreshResponseEnvelopeMessagesCode3003, CatalogSyncRefreshResponseEnvelopeMessagesCode3004, CatalogSyncRefreshResponseEnvelopeMessagesCode3005, CatalogSyncRefreshResponseEnvelopeMessagesCode3006, CatalogSyncRefreshResponseEnvelopeMessagesCode3007, CatalogSyncRefreshResponseEnvelopeMessagesCode4001, CatalogSyncRefreshResponseEnvelopeMessagesCode4002, CatalogSyncRefreshResponseEnvelopeMessagesCode4003, CatalogSyncRefreshResponseEnvelopeMessagesCode4004, CatalogSyncRefreshResponseEnvelopeMessagesCode4005, CatalogSyncRefreshResponseEnvelopeMessagesCode4006, CatalogSyncRefreshResponseEnvelopeMessagesCode4007, CatalogSyncRefreshResponseEnvelopeMessagesCode4008, CatalogSyncRefreshResponseEnvelopeMessagesCode4009, CatalogSyncRefreshResponseEnvelopeMessagesCode4010, CatalogSyncRefreshResponseEnvelopeMessagesCode4011, CatalogSyncRefreshResponseEnvelopeMessagesCode4012, CatalogSyncRefreshResponseEnvelopeMessagesCode4013, CatalogSyncRefreshResponseEnvelopeMessagesCode4014, CatalogSyncRefreshResponseEnvelopeMessagesCode4015, CatalogSyncRefreshResponseEnvelopeMessagesCode4016, CatalogSyncRefreshResponseEnvelopeMessagesCode4017, CatalogSyncRefreshResponseEnvelopeMessagesCode4018, CatalogSyncRefreshResponseEnvelopeMessagesCode4019, CatalogSyncRefreshResponseEnvelopeMessagesCode4020, CatalogSyncRefreshResponseEnvelopeMessagesCode4021, CatalogSyncRefreshResponseEnvelopeMessagesCode4022, CatalogSyncRefreshResponseEnvelopeMessagesCode4023, CatalogSyncRefreshResponseEnvelopeMessagesCode5001, CatalogSyncRefreshResponseEnvelopeMessagesCode5002, CatalogSyncRefreshResponseEnvelopeMessagesCode5003, CatalogSyncRefreshResponseEnvelopeMessagesCode5004, CatalogSyncRefreshResponseEnvelopeMessagesCode102000, CatalogSyncRefreshResponseEnvelopeMessagesCode102001, CatalogSyncRefreshResponseEnvelopeMessagesCode102002, CatalogSyncRefreshResponseEnvelopeMessagesCode102003, CatalogSyncRefreshResponseEnvelopeMessagesCode102004, CatalogSyncRefreshResponseEnvelopeMessagesCode102005, CatalogSyncRefreshResponseEnvelopeMessagesCode102006, CatalogSyncRefreshResponseEnvelopeMessagesCode102007, CatalogSyncRefreshResponseEnvelopeMessagesCode102008, CatalogSyncRefreshResponseEnvelopeMessagesCode102009, CatalogSyncRefreshResponseEnvelopeMessagesCode102010, CatalogSyncRefreshResponseEnvelopeMessagesCode102011, CatalogSyncRefreshResponseEnvelopeMessagesCode102012, CatalogSyncRefreshResponseEnvelopeMessagesCode102013, CatalogSyncRefreshResponseEnvelopeMessagesCode102014, CatalogSyncRefreshResponseEnvelopeMessagesCode102015, CatalogSyncRefreshResponseEnvelopeMessagesCode102016, CatalogSyncRefreshResponseEnvelopeMessagesCode102017, CatalogSyncRefreshResponseEnvelopeMessagesCode102018, CatalogSyncRefreshResponseEnvelopeMessagesCode102019, CatalogSyncRefreshResponseEnvelopeMessagesCode102020, CatalogSyncRefreshResponseEnvelopeMessagesCode102021, CatalogSyncRefreshResponseEnvelopeMessagesCode102022, CatalogSyncRefreshResponseEnvelopeMessagesCode102023, CatalogSyncRefreshResponseEnvelopeMessagesCode102024, CatalogSyncRefreshResponseEnvelopeMessagesCode102025, CatalogSyncRefreshResponseEnvelopeMessagesCode102026, CatalogSyncRefreshResponseEnvelopeMessagesCode102027, CatalogSyncRefreshResponseEnvelopeMessagesCode102028, CatalogSyncRefreshResponseEnvelopeMessagesCode102029, CatalogSyncRefreshResponseEnvelopeMessagesCode102030, CatalogSyncRefreshResponseEnvelopeMessagesCode102031, CatalogSyncRefreshResponseEnvelopeMessagesCode102032, CatalogSyncRefreshResponseEnvelopeMessagesCode102033, CatalogSyncRefreshResponseEnvelopeMessagesCode102034, CatalogSyncRefreshResponseEnvelopeMessagesCode102035, CatalogSyncRefreshResponseEnvelopeMessagesCode102036, CatalogSyncRefreshResponseEnvelopeMessagesCode102037, CatalogSyncRefreshResponseEnvelopeMessagesCode102038, CatalogSyncRefreshResponseEnvelopeMessagesCode102039, CatalogSyncRefreshResponseEnvelopeMessagesCode102040, CatalogSyncRefreshResponseEnvelopeMessagesCode102041, CatalogSyncRefreshResponseEnvelopeMessagesCode102042, CatalogSyncRefreshResponseEnvelopeMessagesCode102043, CatalogSyncRefreshResponseEnvelopeMessagesCode102044, CatalogSyncRefreshResponseEnvelopeMessagesCode102045, CatalogSyncRefreshResponseEnvelopeMessagesCode102046, CatalogSyncRefreshResponseEnvelopeMessagesCode102047, CatalogSyncRefreshResponseEnvelopeMessagesCode102048, CatalogSyncRefreshResponseEnvelopeMessagesCode102049, CatalogSyncRefreshResponseEnvelopeMessagesCode102050, CatalogSyncRefreshResponseEnvelopeMessagesCode102051, CatalogSyncRefreshResponseEnvelopeMessagesCode102052, CatalogSyncRefreshResponseEnvelopeMessagesCode102053, CatalogSyncRefreshResponseEnvelopeMessagesCode102054, CatalogSyncRefreshResponseEnvelopeMessagesCode102055, CatalogSyncRefreshResponseEnvelopeMessagesCode102056, CatalogSyncRefreshResponseEnvelopeMessagesCode102057, CatalogSyncRefreshResponseEnvelopeMessagesCode102058, CatalogSyncRefreshResponseEnvelopeMessagesCode102059, CatalogSyncRefreshResponseEnvelopeMessagesCode102060, CatalogSyncRefreshResponseEnvelopeMessagesCode102061, CatalogSyncRefreshResponseEnvelopeMessagesCode102062, CatalogSyncRefreshResponseEnvelopeMessagesCode102063, CatalogSyncRefreshResponseEnvelopeMessagesCode102064, CatalogSyncRefreshResponseEnvelopeMessagesCode102065, CatalogSyncRefreshResponseEnvelopeMessagesCode102066, CatalogSyncRefreshResponseEnvelopeMessagesCode103001, CatalogSyncRefreshResponseEnvelopeMessagesCode103002, CatalogSyncRefreshResponseEnvelopeMessagesCode103003, CatalogSyncRefreshResponseEnvelopeMessagesCode103004, CatalogSyncRefreshResponseEnvelopeMessagesCode103005, CatalogSyncRefreshResponseEnvelopeMessagesCode103006, CatalogSyncRefreshResponseEnvelopeMessagesCode103007, CatalogSyncRefreshResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CatalogSyncRefreshResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                             `json:"l10n_key"`
	LoggableError string                                             `json:"loggable_error"`
	TemplateData  interface{}                                        `json:"template_data"`
	TraceID       string                                             `json:"trace_id"`
	JSON          catalogSyncRefreshResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeMessagesMetaJSON contains the JSON metadata
// for the struct [CatalogSyncRefreshResponseEnvelopeMessagesMeta]
type catalogSyncRefreshResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncRefreshResponseEnvelopeMessagesSource struct {
	Parameter           string                                               `json:"parameter"`
	ParameterValueIndex int64                                                `json:"parameter_value_index"`
	Pointer             string                                               `json:"pointer"`
	JSON                catalogSyncRefreshResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// catalogSyncRefreshResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CatalogSyncRefreshResponseEnvelopeMessagesSource]
type catalogSyncRefreshResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CatalogSyncRefreshResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncRefreshResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
