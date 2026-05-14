// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking

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

// OnRampAddressSpaceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnRampAddressSpaceService] method instead.
type OnRampAddressSpaceService struct {
	Options []option.RequestOption
}

// NewOnRampAddressSpaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOnRampAddressSpaceService(opts ...option.RequestOption) (r *OnRampAddressSpaceService) {
	r = &OnRampAddressSpaceService{}
	r.Options = opts
	return
}

// Update the Magic WAN Address Space (Closed Beta).
func (r *OnRampAddressSpaceService) Update(ctx context.Context, params OnRampAddressSpaceUpdateParams, opts ...option.RequestOption) (res *OnRampAddressSpaceUpdateResponse, err error) {
	var env OnRampAddressSpaceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read the Magic WAN Address Space (Closed Beta).
func (r *OnRampAddressSpaceService) List(ctx context.Context, query OnRampAddressSpaceListParams, opts ...option.RequestOption) (res *OnRampAddressSpaceListResponse, err error) {
	var env OnRampAddressSpaceListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the Magic WAN Address Space (Closed Beta).
func (r *OnRampAddressSpaceService) Edit(ctx context.Context, params OnRampAddressSpaceEditParams, opts ...option.RequestOption) (res *OnRampAddressSpaceEditResponse, err error) {
	var env OnRampAddressSpaceEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OnRampAddressSpaceUpdateResponse struct {
	Prefixes []string                             `json:"prefixes,required"`
	JSON     onRampAddressSpaceUpdateResponseJSON `json:"-"`
}

// onRampAddressSpaceUpdateResponseJSON contains the JSON metadata for the struct
// [OnRampAddressSpaceUpdateResponse]
type onRampAddressSpaceUpdateResponseJSON struct {
	Prefixes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponse struct {
	Prefixes []string                           `json:"prefixes,required"`
	JSON     onRampAddressSpaceListResponseJSON `json:"-"`
}

// onRampAddressSpaceListResponseJSON contains the JSON metadata for the struct
// [OnRampAddressSpaceListResponse]
type onRampAddressSpaceListResponseJSON struct {
	Prefixes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponse struct {
	Prefixes []string                           `json:"prefixes,required"`
	JSON     onRampAddressSpaceEditResponseJSON `json:"-"`
}

// onRampAddressSpaceEditResponseJSON contains the JSON metadata for the struct
// [OnRampAddressSpaceEditResponse]
type onRampAddressSpaceEditResponseJSON struct {
	Prefixes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Prefixes  param.Field[[]string] `json:"prefixes,required"`
}

func (r OnRampAddressSpaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OnRampAddressSpaceUpdateResponseEnvelope struct {
	Errors   []OnRampAddressSpaceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OnRampAddressSpaceUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   OnRampAddressSpaceUpdateResponse                   `json:"result,required"`
	Success  bool                                               `json:"success,required"`
	JSON     onRampAddressSpaceUpdateResponseEnvelopeJSON       `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [OnRampAddressSpaceUpdateResponseEnvelope]
type onRampAddressSpaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeErrors struct {
	Code             OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Meta             OnRampAddressSpaceUpdateResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           OnRampAddressSpaceUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             onRampAddressSpaceUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceUpdateResponseEnvelopeErrors]
type onRampAddressSpaceUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode int64

const (
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1001   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1002   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1003   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1004   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1005   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1006   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1007   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1008   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1008
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1009   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1009
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1010   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1010
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1011   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1011
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1012   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1012
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1013   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1013
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1014   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1014
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1015   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1015
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1016   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1016
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1017   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 1017
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2001   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2002   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2003   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2004   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2005   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2006   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2007   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2008   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2008
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2009   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2009
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2010   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2010
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2011   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2011
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2012   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2012
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2013   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2013
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2014   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2014
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2015   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2015
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2016   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2016
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2017   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2017
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2018   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2018
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2019   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2019
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2020   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2020
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2021   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2021
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2022   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 2022
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3001   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3002   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3003   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3004   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3005   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3006   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3007   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 3007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4001   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4002   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4003   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4004   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4005   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4006   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4007   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4008   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4008
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4009   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4009
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4010   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4010
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4011   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4011
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4012   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4012
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4013   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4013
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4014   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4014
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4015   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4015
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4016   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4016
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4017   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4017
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4018   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4018
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4019   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4019
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4020   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4020
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4021   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4021
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4022   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4022
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4023   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 4023
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5001   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 5001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5002   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 5002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5003   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 5003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5004   OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 5004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102000 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102000
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102001 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102002 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102003 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102004 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102005 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102006 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102007 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102008 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102008
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102009 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102009
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102010 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102010
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102011 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102011
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102012 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102012
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102013 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102013
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102014 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102014
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102015 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102015
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102016 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102016
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102017 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102017
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102018 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102018
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102019 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102019
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102020 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102020
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102021 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102021
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102022 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102022
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102023 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102023
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102024 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102024
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102025 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102025
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102026 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102026
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102027 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102027
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102028 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102028
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102029 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102029
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102030 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102030
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102031 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102031
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102032 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102032
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102033 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102033
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102034 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102034
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102035 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102035
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102036 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102036
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102037 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102037
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102038 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102038
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102039 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102039
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102040 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102040
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102041 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102041
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102042 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102042
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102043 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102043
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102044 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102044
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102045 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102045
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102046 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102046
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102047 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102047
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102048 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102048
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102049 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102049
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102050 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102050
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102051 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102051
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102052 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102052
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102053 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102053
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102054 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102054
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102055 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102055
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102056 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102056
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102057 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102057
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102058 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102058
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102059 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102059
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102060 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102060
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102061 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102061
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102062 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102062
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102063 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102063
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102064 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102064
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102065 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102065
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102066 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 102066
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103001 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103001
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103002 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103002
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103003 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103003
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103004 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103004
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103005 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103005
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103006 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103006
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103007 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103007
	OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103008 OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode = 103008
)

func (r OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1008, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1009, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1010, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1011, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1012, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1013, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1014, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1015, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1016, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode1017, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2008, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2009, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2010, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2011, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2012, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2013, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2014, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2015, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2016, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2017, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2018, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2019, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2020, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2021, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode2022, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode3007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4008, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4009, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4010, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4011, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4012, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4013, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4014, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4015, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4016, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4017, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4018, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4019, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4020, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4021, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4022, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode4023, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode5004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102000, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102008, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102009, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102010, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102011, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102012, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102013, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102014, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102015, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102016, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102017, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102018, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102019, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102020, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102021, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102022, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102023, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102024, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102025, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102026, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102027, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102028, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102029, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102030, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102031, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102032, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102033, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102034, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102035, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102036, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102037, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102038, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102039, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102040, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102041, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102042, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102043, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102044, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102045, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102046, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102047, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102048, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102049, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102050, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102051, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102052, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102053, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102054, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102055, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102056, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102057, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102058, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102059, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102060, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102061, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102062, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102063, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102064, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102065, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode102066, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103001, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103002, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103003, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103004, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103005, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103006, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103007, OnRampAddressSpaceUpdateResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceUpdateResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                                 `json:"l10n_key"`
	LoggableError string                                                 `json:"loggable_error"`
	TemplateData  interface{}                                            `json:"template_data"`
	TraceID       string                                                 `json:"trace_id"`
	JSON          onRampAddressSpaceUpdateResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeErrorsMetaJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceUpdateResponseEnvelopeErrorsMeta]
type onRampAddressSpaceUpdateResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeErrorsSource struct {
	Parameter           string                                                   `json:"parameter"`
	ParameterValueIndex int64                                                    `json:"parameter_value_index"`
	Pointer             string                                                   `json:"pointer"`
	JSON                onRampAddressSpaceUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceUpdateResponseEnvelopeErrorsSource]
type onRampAddressSpaceUpdateResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeMessages struct {
	Code             OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Meta             OnRampAddressSpaceUpdateResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           OnRampAddressSpaceUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             onRampAddressSpaceUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceUpdateResponseEnvelopeMessages]
type onRampAddressSpaceUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode int64

const (
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1001   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1002   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1003   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1004   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1005   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1006   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1007   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1008   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1008
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1009   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1009
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1010   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1010
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1011   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1011
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1012   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1012
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1013   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1013
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1014   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1014
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1015   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1015
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1016   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1016
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1017   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 1017
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2001   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2002   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2003   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2004   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2005   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2006   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2007   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2008   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2008
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2009   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2009
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2010   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2010
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2011   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2011
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2012   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2012
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2013   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2013
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2014   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2014
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2015   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2015
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2016   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2016
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2017   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2017
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2018   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2018
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2019   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2019
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2020   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2020
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2021   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2021
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2022   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 2022
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3001   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3002   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3003   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3004   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3005   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3006   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3007   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 3007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4001   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4002   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4003   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4004   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4005   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4006   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4007   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4008   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4008
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4009   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4009
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4010   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4010
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4011   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4011
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4012   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4012
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4013   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4013
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4014   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4014
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4015   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4015
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4016   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4016
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4017   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4017
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4018   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4018
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4019   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4019
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4020   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4020
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4021   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4021
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4022   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4022
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4023   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 4023
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5001   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 5001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5002   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 5002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5003   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 5003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5004   OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 5004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102000 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102000
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102001 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102002 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102003 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102004 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102005 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102006 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102007 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102008 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102008
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102009 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102009
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102010 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102010
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102011 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102011
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102012 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102012
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102013 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102013
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102014 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102014
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102015 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102015
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102016 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102016
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102017 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102017
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102018 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102018
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102019 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102019
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102020 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102020
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102021 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102021
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102022 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102022
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102023 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102023
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102024 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102024
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102025 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102025
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102026 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102026
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102027 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102027
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102028 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102028
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102029 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102029
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102030 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102030
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102031 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102031
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102032 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102032
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102033 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102033
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102034 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102034
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102035 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102035
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102036 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102036
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102037 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102037
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102038 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102038
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102039 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102039
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102040 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102040
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102041 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102041
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102042 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102042
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102043 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102043
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102044 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102044
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102045 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102045
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102046 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102046
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102047 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102047
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102048 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102048
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102049 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102049
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102050 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102050
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102051 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102051
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102052 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102052
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102053 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102053
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102054 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102054
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102055 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102055
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102056 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102056
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102057 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102057
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102058 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102058
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102059 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102059
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102060 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102060
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102061 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102061
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102062 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102062
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102063 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102063
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102064 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102064
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102065 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102065
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102066 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 102066
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103001 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103001
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103002 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103002
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103003 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103003
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103004 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103004
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103005 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103005
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103006 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103006
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103007 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103007
	OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103008 OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode = 103008
)

func (r OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1008, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1009, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1010, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1011, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1012, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1013, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1014, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1015, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1016, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode1017, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2008, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2009, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2010, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2011, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2012, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2013, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2014, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2015, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2016, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2017, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2018, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2019, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2020, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2021, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode2022, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode3007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4008, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4009, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4010, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4011, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4012, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4013, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4014, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4015, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4016, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4017, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4018, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4019, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4020, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4021, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4022, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode4023, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode5004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102000, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102008, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102009, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102010, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102011, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102012, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102013, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102014, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102015, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102016, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102017, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102018, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102019, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102020, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102021, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102022, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102023, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102024, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102025, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102026, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102027, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102028, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102029, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102030, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102031, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102032, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102033, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102034, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102035, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102036, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102037, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102038, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102039, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102040, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102041, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102042, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102043, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102044, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102045, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102046, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102047, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102048, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102049, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102050, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102051, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102052, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102053, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102054, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102055, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102056, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102057, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102058, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102059, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102060, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102061, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102062, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102063, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102064, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102065, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode102066, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103001, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103002, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103003, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103004, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103005, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103006, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103007, OnRampAddressSpaceUpdateResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceUpdateResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                   `json:"l10n_key"`
	LoggableError string                                                   `json:"loggable_error"`
	TemplateData  interface{}                                              `json:"template_data"`
	TraceID       string                                                   `json:"trace_id"`
	JSON          onRampAddressSpaceUpdateResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceUpdateResponseEnvelopeMessagesMeta]
type onRampAddressSpaceUpdateResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceUpdateResponseEnvelopeMessagesSource struct {
	Parameter           string                                                     `json:"parameter"`
	ParameterValueIndex int64                                                      `json:"parameter_value_index"`
	Pointer             string                                                     `json:"pointer"`
	JSON                onRampAddressSpaceUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// onRampAddressSpaceUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceUpdateResponseEnvelopeMessagesSource]
type onRampAddressSpaceUpdateResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type OnRampAddressSpaceListResponseEnvelope struct {
	Errors   []OnRampAddressSpaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OnRampAddressSpaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   OnRampAddressSpaceListResponse                   `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     onRampAddressSpaceListResponseEnvelopeJSON       `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [OnRampAddressSpaceListResponseEnvelope]
type onRampAddressSpaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeErrors struct {
	Code             OnRampAddressSpaceListResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Meta             OnRampAddressSpaceListResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           OnRampAddressSpaceListResponseEnvelopeErrorsSource `json:"source"`
	JSON             onRampAddressSpaceListResponseEnvelopeErrorsJSON   `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OnRampAddressSpaceListResponseEnvelopeErrors]
type onRampAddressSpaceListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeErrorsCode int64

const (
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1001   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1002   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1003   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1004   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1005   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1006   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1007   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1008   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1008
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1009   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1009
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1010   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1010
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1011   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1011
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1012   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1012
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1013   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1013
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1014   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1014
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1015   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1015
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1016   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1016
	OnRampAddressSpaceListResponseEnvelopeErrorsCode1017   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 1017
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2001   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2002   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2003   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2004   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2005   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2006   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2007   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2008   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2008
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2009   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2009
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2010   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2010
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2011   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2011
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2012   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2012
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2013   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2013
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2014   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2014
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2015   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2015
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2016   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2016
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2017   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2017
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2018   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2018
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2019   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2019
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2020   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2020
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2021   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2021
	OnRampAddressSpaceListResponseEnvelopeErrorsCode2022   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 2022
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3001   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3002   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3003   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3004   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3005   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3006   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode3007   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 3007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4001   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4002   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4003   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4004   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4005   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4006   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4007   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4008   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4008
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4009   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4009
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4010   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4010
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4011   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4011
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4012   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4012
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4013   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4013
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4014   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4014
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4015   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4015
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4016   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4016
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4017   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4017
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4018   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4018
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4019   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4019
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4020   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4020
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4021   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4021
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4022   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4022
	OnRampAddressSpaceListResponseEnvelopeErrorsCode4023   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 4023
	OnRampAddressSpaceListResponseEnvelopeErrorsCode5001   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 5001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode5002   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 5002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode5003   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 5003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode5004   OnRampAddressSpaceListResponseEnvelopeErrorsCode = 5004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102000 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102000
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102001 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102002 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102003 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102004 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102005 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102006 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102007 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102008 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102008
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102009 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102009
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102010 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102010
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102011 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102011
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102012 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102012
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102013 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102013
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102014 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102014
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102015 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102015
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102016 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102016
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102017 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102017
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102018 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102018
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102019 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102019
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102020 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102020
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102021 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102021
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102022 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102022
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102023 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102023
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102024 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102024
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102025 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102025
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102026 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102026
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102027 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102027
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102028 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102028
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102029 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102029
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102030 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102030
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102031 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102031
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102032 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102032
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102033 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102033
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102034 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102034
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102035 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102035
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102036 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102036
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102037 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102037
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102038 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102038
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102039 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102039
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102040 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102040
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102041 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102041
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102042 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102042
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102043 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102043
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102044 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102044
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102045 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102045
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102046 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102046
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102047 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102047
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102048 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102048
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102049 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102049
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102050 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102050
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102051 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102051
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102052 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102052
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102053 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102053
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102054 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102054
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102055 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102055
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102056 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102056
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102057 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102057
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102058 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102058
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102059 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102059
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102060 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102060
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102061 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102061
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102062 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102062
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102063 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102063
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102064 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102064
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102065 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102065
	OnRampAddressSpaceListResponseEnvelopeErrorsCode102066 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 102066
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103001 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103001
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103002 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103002
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103003 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103003
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103004 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103004
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103005 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103005
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103006 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103006
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103007 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103007
	OnRampAddressSpaceListResponseEnvelopeErrorsCode103008 OnRampAddressSpaceListResponseEnvelopeErrorsCode = 103008
)

func (r OnRampAddressSpaceListResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceListResponseEnvelopeErrorsCode1001, OnRampAddressSpaceListResponseEnvelopeErrorsCode1002, OnRampAddressSpaceListResponseEnvelopeErrorsCode1003, OnRampAddressSpaceListResponseEnvelopeErrorsCode1004, OnRampAddressSpaceListResponseEnvelopeErrorsCode1005, OnRampAddressSpaceListResponseEnvelopeErrorsCode1006, OnRampAddressSpaceListResponseEnvelopeErrorsCode1007, OnRampAddressSpaceListResponseEnvelopeErrorsCode1008, OnRampAddressSpaceListResponseEnvelopeErrorsCode1009, OnRampAddressSpaceListResponseEnvelopeErrorsCode1010, OnRampAddressSpaceListResponseEnvelopeErrorsCode1011, OnRampAddressSpaceListResponseEnvelopeErrorsCode1012, OnRampAddressSpaceListResponseEnvelopeErrorsCode1013, OnRampAddressSpaceListResponseEnvelopeErrorsCode1014, OnRampAddressSpaceListResponseEnvelopeErrorsCode1015, OnRampAddressSpaceListResponseEnvelopeErrorsCode1016, OnRampAddressSpaceListResponseEnvelopeErrorsCode1017, OnRampAddressSpaceListResponseEnvelopeErrorsCode2001, OnRampAddressSpaceListResponseEnvelopeErrorsCode2002, OnRampAddressSpaceListResponseEnvelopeErrorsCode2003, OnRampAddressSpaceListResponseEnvelopeErrorsCode2004, OnRampAddressSpaceListResponseEnvelopeErrorsCode2005, OnRampAddressSpaceListResponseEnvelopeErrorsCode2006, OnRampAddressSpaceListResponseEnvelopeErrorsCode2007, OnRampAddressSpaceListResponseEnvelopeErrorsCode2008, OnRampAddressSpaceListResponseEnvelopeErrorsCode2009, OnRampAddressSpaceListResponseEnvelopeErrorsCode2010, OnRampAddressSpaceListResponseEnvelopeErrorsCode2011, OnRampAddressSpaceListResponseEnvelopeErrorsCode2012, OnRampAddressSpaceListResponseEnvelopeErrorsCode2013, OnRampAddressSpaceListResponseEnvelopeErrorsCode2014, OnRampAddressSpaceListResponseEnvelopeErrorsCode2015, OnRampAddressSpaceListResponseEnvelopeErrorsCode2016, OnRampAddressSpaceListResponseEnvelopeErrorsCode2017, OnRampAddressSpaceListResponseEnvelopeErrorsCode2018, OnRampAddressSpaceListResponseEnvelopeErrorsCode2019, OnRampAddressSpaceListResponseEnvelopeErrorsCode2020, OnRampAddressSpaceListResponseEnvelopeErrorsCode2021, OnRampAddressSpaceListResponseEnvelopeErrorsCode2022, OnRampAddressSpaceListResponseEnvelopeErrorsCode3001, OnRampAddressSpaceListResponseEnvelopeErrorsCode3002, OnRampAddressSpaceListResponseEnvelopeErrorsCode3003, OnRampAddressSpaceListResponseEnvelopeErrorsCode3004, OnRampAddressSpaceListResponseEnvelopeErrorsCode3005, OnRampAddressSpaceListResponseEnvelopeErrorsCode3006, OnRampAddressSpaceListResponseEnvelopeErrorsCode3007, OnRampAddressSpaceListResponseEnvelopeErrorsCode4001, OnRampAddressSpaceListResponseEnvelopeErrorsCode4002, OnRampAddressSpaceListResponseEnvelopeErrorsCode4003, OnRampAddressSpaceListResponseEnvelopeErrorsCode4004, OnRampAddressSpaceListResponseEnvelopeErrorsCode4005, OnRampAddressSpaceListResponseEnvelopeErrorsCode4006, OnRampAddressSpaceListResponseEnvelopeErrorsCode4007, OnRampAddressSpaceListResponseEnvelopeErrorsCode4008, OnRampAddressSpaceListResponseEnvelopeErrorsCode4009, OnRampAddressSpaceListResponseEnvelopeErrorsCode4010, OnRampAddressSpaceListResponseEnvelopeErrorsCode4011, OnRampAddressSpaceListResponseEnvelopeErrorsCode4012, OnRampAddressSpaceListResponseEnvelopeErrorsCode4013, OnRampAddressSpaceListResponseEnvelopeErrorsCode4014, OnRampAddressSpaceListResponseEnvelopeErrorsCode4015, OnRampAddressSpaceListResponseEnvelopeErrorsCode4016, OnRampAddressSpaceListResponseEnvelopeErrorsCode4017, OnRampAddressSpaceListResponseEnvelopeErrorsCode4018, OnRampAddressSpaceListResponseEnvelopeErrorsCode4019, OnRampAddressSpaceListResponseEnvelopeErrorsCode4020, OnRampAddressSpaceListResponseEnvelopeErrorsCode4021, OnRampAddressSpaceListResponseEnvelopeErrorsCode4022, OnRampAddressSpaceListResponseEnvelopeErrorsCode4023, OnRampAddressSpaceListResponseEnvelopeErrorsCode5001, OnRampAddressSpaceListResponseEnvelopeErrorsCode5002, OnRampAddressSpaceListResponseEnvelopeErrorsCode5003, OnRampAddressSpaceListResponseEnvelopeErrorsCode5004, OnRampAddressSpaceListResponseEnvelopeErrorsCode102000, OnRampAddressSpaceListResponseEnvelopeErrorsCode102001, OnRampAddressSpaceListResponseEnvelopeErrorsCode102002, OnRampAddressSpaceListResponseEnvelopeErrorsCode102003, OnRampAddressSpaceListResponseEnvelopeErrorsCode102004, OnRampAddressSpaceListResponseEnvelopeErrorsCode102005, OnRampAddressSpaceListResponseEnvelopeErrorsCode102006, OnRampAddressSpaceListResponseEnvelopeErrorsCode102007, OnRampAddressSpaceListResponseEnvelopeErrorsCode102008, OnRampAddressSpaceListResponseEnvelopeErrorsCode102009, OnRampAddressSpaceListResponseEnvelopeErrorsCode102010, OnRampAddressSpaceListResponseEnvelopeErrorsCode102011, OnRampAddressSpaceListResponseEnvelopeErrorsCode102012, OnRampAddressSpaceListResponseEnvelopeErrorsCode102013, OnRampAddressSpaceListResponseEnvelopeErrorsCode102014, OnRampAddressSpaceListResponseEnvelopeErrorsCode102015, OnRampAddressSpaceListResponseEnvelopeErrorsCode102016, OnRampAddressSpaceListResponseEnvelopeErrorsCode102017, OnRampAddressSpaceListResponseEnvelopeErrorsCode102018, OnRampAddressSpaceListResponseEnvelopeErrorsCode102019, OnRampAddressSpaceListResponseEnvelopeErrorsCode102020, OnRampAddressSpaceListResponseEnvelopeErrorsCode102021, OnRampAddressSpaceListResponseEnvelopeErrorsCode102022, OnRampAddressSpaceListResponseEnvelopeErrorsCode102023, OnRampAddressSpaceListResponseEnvelopeErrorsCode102024, OnRampAddressSpaceListResponseEnvelopeErrorsCode102025, OnRampAddressSpaceListResponseEnvelopeErrorsCode102026, OnRampAddressSpaceListResponseEnvelopeErrorsCode102027, OnRampAddressSpaceListResponseEnvelopeErrorsCode102028, OnRampAddressSpaceListResponseEnvelopeErrorsCode102029, OnRampAddressSpaceListResponseEnvelopeErrorsCode102030, OnRampAddressSpaceListResponseEnvelopeErrorsCode102031, OnRampAddressSpaceListResponseEnvelopeErrorsCode102032, OnRampAddressSpaceListResponseEnvelopeErrorsCode102033, OnRampAddressSpaceListResponseEnvelopeErrorsCode102034, OnRampAddressSpaceListResponseEnvelopeErrorsCode102035, OnRampAddressSpaceListResponseEnvelopeErrorsCode102036, OnRampAddressSpaceListResponseEnvelopeErrorsCode102037, OnRampAddressSpaceListResponseEnvelopeErrorsCode102038, OnRampAddressSpaceListResponseEnvelopeErrorsCode102039, OnRampAddressSpaceListResponseEnvelopeErrorsCode102040, OnRampAddressSpaceListResponseEnvelopeErrorsCode102041, OnRampAddressSpaceListResponseEnvelopeErrorsCode102042, OnRampAddressSpaceListResponseEnvelopeErrorsCode102043, OnRampAddressSpaceListResponseEnvelopeErrorsCode102044, OnRampAddressSpaceListResponseEnvelopeErrorsCode102045, OnRampAddressSpaceListResponseEnvelopeErrorsCode102046, OnRampAddressSpaceListResponseEnvelopeErrorsCode102047, OnRampAddressSpaceListResponseEnvelopeErrorsCode102048, OnRampAddressSpaceListResponseEnvelopeErrorsCode102049, OnRampAddressSpaceListResponseEnvelopeErrorsCode102050, OnRampAddressSpaceListResponseEnvelopeErrorsCode102051, OnRampAddressSpaceListResponseEnvelopeErrorsCode102052, OnRampAddressSpaceListResponseEnvelopeErrorsCode102053, OnRampAddressSpaceListResponseEnvelopeErrorsCode102054, OnRampAddressSpaceListResponseEnvelopeErrorsCode102055, OnRampAddressSpaceListResponseEnvelopeErrorsCode102056, OnRampAddressSpaceListResponseEnvelopeErrorsCode102057, OnRampAddressSpaceListResponseEnvelopeErrorsCode102058, OnRampAddressSpaceListResponseEnvelopeErrorsCode102059, OnRampAddressSpaceListResponseEnvelopeErrorsCode102060, OnRampAddressSpaceListResponseEnvelopeErrorsCode102061, OnRampAddressSpaceListResponseEnvelopeErrorsCode102062, OnRampAddressSpaceListResponseEnvelopeErrorsCode102063, OnRampAddressSpaceListResponseEnvelopeErrorsCode102064, OnRampAddressSpaceListResponseEnvelopeErrorsCode102065, OnRampAddressSpaceListResponseEnvelopeErrorsCode102066, OnRampAddressSpaceListResponseEnvelopeErrorsCode103001, OnRampAddressSpaceListResponseEnvelopeErrorsCode103002, OnRampAddressSpaceListResponseEnvelopeErrorsCode103003, OnRampAddressSpaceListResponseEnvelopeErrorsCode103004, OnRampAddressSpaceListResponseEnvelopeErrorsCode103005, OnRampAddressSpaceListResponseEnvelopeErrorsCode103006, OnRampAddressSpaceListResponseEnvelopeErrorsCode103007, OnRampAddressSpaceListResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceListResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                               `json:"l10n_key"`
	LoggableError string                                               `json:"loggable_error"`
	TemplateData  interface{}                                          `json:"template_data"`
	TraceID       string                                               `json:"trace_id"`
	JSON          onRampAddressSpaceListResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceListResponseEnvelopeErrorsMeta]
type onRampAddressSpaceListResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeErrorsSource struct {
	Parameter           string                                                 `json:"parameter"`
	ParameterValueIndex int64                                                  `json:"parameter_value_index"`
	Pointer             string                                                 `json:"pointer"`
	JSON                onRampAddressSpaceListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceListResponseEnvelopeErrorsSource]
type onRampAddressSpaceListResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeMessages struct {
	Code             OnRampAddressSpaceListResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Meta             OnRampAddressSpaceListResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           OnRampAddressSpaceListResponseEnvelopeMessagesSource `json:"source"`
	JSON             onRampAddressSpaceListResponseEnvelopeMessagesJSON   `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceListResponseEnvelopeMessages]
type onRampAddressSpaceListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeMessagesCode int64

const (
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1001   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1002   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1003   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1004   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1005   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1006   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1007   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1008   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1008
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1009   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1009
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1010   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1010
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1011   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1011
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1012   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1012
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1013   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1013
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1014   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1014
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1015   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1015
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1016   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1016
	OnRampAddressSpaceListResponseEnvelopeMessagesCode1017   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 1017
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2001   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2002   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2003   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2004   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2005   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2006   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2007   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2008   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2008
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2009   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2009
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2010   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2010
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2011   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2011
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2012   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2012
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2013   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2013
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2014   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2014
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2015   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2015
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2016   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2016
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2017   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2017
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2018   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2018
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2019   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2019
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2020   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2020
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2021   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2021
	OnRampAddressSpaceListResponseEnvelopeMessagesCode2022   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 2022
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3001   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3002   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3003   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3004   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3005   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3006   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode3007   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 3007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4001   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4002   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4003   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4004   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4005   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4006   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4007   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4008   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4008
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4009   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4009
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4010   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4010
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4011   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4011
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4012   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4012
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4013   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4013
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4014   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4014
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4015   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4015
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4016   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4016
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4017   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4017
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4018   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4018
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4019   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4019
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4020   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4020
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4021   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4021
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4022   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4022
	OnRampAddressSpaceListResponseEnvelopeMessagesCode4023   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 4023
	OnRampAddressSpaceListResponseEnvelopeMessagesCode5001   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 5001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode5002   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 5002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode5003   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 5003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode5004   OnRampAddressSpaceListResponseEnvelopeMessagesCode = 5004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102000 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102000
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102001 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102002 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102003 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102004 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102005 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102006 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102007 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102008 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102008
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102009 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102009
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102010 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102010
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102011 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102011
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102012 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102012
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102013 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102013
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102014 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102014
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102015 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102015
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102016 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102016
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102017 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102017
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102018 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102018
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102019 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102019
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102020 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102020
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102021 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102021
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102022 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102022
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102023 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102023
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102024 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102024
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102025 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102025
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102026 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102026
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102027 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102027
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102028 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102028
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102029 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102029
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102030 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102030
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102031 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102031
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102032 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102032
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102033 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102033
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102034 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102034
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102035 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102035
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102036 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102036
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102037 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102037
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102038 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102038
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102039 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102039
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102040 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102040
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102041 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102041
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102042 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102042
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102043 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102043
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102044 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102044
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102045 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102045
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102046 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102046
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102047 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102047
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102048 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102048
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102049 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102049
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102050 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102050
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102051 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102051
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102052 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102052
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102053 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102053
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102054 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102054
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102055 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102055
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102056 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102056
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102057 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102057
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102058 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102058
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102059 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102059
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102060 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102060
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102061 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102061
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102062 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102062
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102063 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102063
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102064 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102064
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102065 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102065
	OnRampAddressSpaceListResponseEnvelopeMessagesCode102066 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 102066
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103001 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103001
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103002 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103002
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103003 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103003
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103004 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103004
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103005 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103005
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103006 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103006
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103007 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103007
	OnRampAddressSpaceListResponseEnvelopeMessagesCode103008 OnRampAddressSpaceListResponseEnvelopeMessagesCode = 103008
)

func (r OnRampAddressSpaceListResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceListResponseEnvelopeMessagesCode1001, OnRampAddressSpaceListResponseEnvelopeMessagesCode1002, OnRampAddressSpaceListResponseEnvelopeMessagesCode1003, OnRampAddressSpaceListResponseEnvelopeMessagesCode1004, OnRampAddressSpaceListResponseEnvelopeMessagesCode1005, OnRampAddressSpaceListResponseEnvelopeMessagesCode1006, OnRampAddressSpaceListResponseEnvelopeMessagesCode1007, OnRampAddressSpaceListResponseEnvelopeMessagesCode1008, OnRampAddressSpaceListResponseEnvelopeMessagesCode1009, OnRampAddressSpaceListResponseEnvelopeMessagesCode1010, OnRampAddressSpaceListResponseEnvelopeMessagesCode1011, OnRampAddressSpaceListResponseEnvelopeMessagesCode1012, OnRampAddressSpaceListResponseEnvelopeMessagesCode1013, OnRampAddressSpaceListResponseEnvelopeMessagesCode1014, OnRampAddressSpaceListResponseEnvelopeMessagesCode1015, OnRampAddressSpaceListResponseEnvelopeMessagesCode1016, OnRampAddressSpaceListResponseEnvelopeMessagesCode1017, OnRampAddressSpaceListResponseEnvelopeMessagesCode2001, OnRampAddressSpaceListResponseEnvelopeMessagesCode2002, OnRampAddressSpaceListResponseEnvelopeMessagesCode2003, OnRampAddressSpaceListResponseEnvelopeMessagesCode2004, OnRampAddressSpaceListResponseEnvelopeMessagesCode2005, OnRampAddressSpaceListResponseEnvelopeMessagesCode2006, OnRampAddressSpaceListResponseEnvelopeMessagesCode2007, OnRampAddressSpaceListResponseEnvelopeMessagesCode2008, OnRampAddressSpaceListResponseEnvelopeMessagesCode2009, OnRampAddressSpaceListResponseEnvelopeMessagesCode2010, OnRampAddressSpaceListResponseEnvelopeMessagesCode2011, OnRampAddressSpaceListResponseEnvelopeMessagesCode2012, OnRampAddressSpaceListResponseEnvelopeMessagesCode2013, OnRampAddressSpaceListResponseEnvelopeMessagesCode2014, OnRampAddressSpaceListResponseEnvelopeMessagesCode2015, OnRampAddressSpaceListResponseEnvelopeMessagesCode2016, OnRampAddressSpaceListResponseEnvelopeMessagesCode2017, OnRampAddressSpaceListResponseEnvelopeMessagesCode2018, OnRampAddressSpaceListResponseEnvelopeMessagesCode2019, OnRampAddressSpaceListResponseEnvelopeMessagesCode2020, OnRampAddressSpaceListResponseEnvelopeMessagesCode2021, OnRampAddressSpaceListResponseEnvelopeMessagesCode2022, OnRampAddressSpaceListResponseEnvelopeMessagesCode3001, OnRampAddressSpaceListResponseEnvelopeMessagesCode3002, OnRampAddressSpaceListResponseEnvelopeMessagesCode3003, OnRampAddressSpaceListResponseEnvelopeMessagesCode3004, OnRampAddressSpaceListResponseEnvelopeMessagesCode3005, OnRampAddressSpaceListResponseEnvelopeMessagesCode3006, OnRampAddressSpaceListResponseEnvelopeMessagesCode3007, OnRampAddressSpaceListResponseEnvelopeMessagesCode4001, OnRampAddressSpaceListResponseEnvelopeMessagesCode4002, OnRampAddressSpaceListResponseEnvelopeMessagesCode4003, OnRampAddressSpaceListResponseEnvelopeMessagesCode4004, OnRampAddressSpaceListResponseEnvelopeMessagesCode4005, OnRampAddressSpaceListResponseEnvelopeMessagesCode4006, OnRampAddressSpaceListResponseEnvelopeMessagesCode4007, OnRampAddressSpaceListResponseEnvelopeMessagesCode4008, OnRampAddressSpaceListResponseEnvelopeMessagesCode4009, OnRampAddressSpaceListResponseEnvelopeMessagesCode4010, OnRampAddressSpaceListResponseEnvelopeMessagesCode4011, OnRampAddressSpaceListResponseEnvelopeMessagesCode4012, OnRampAddressSpaceListResponseEnvelopeMessagesCode4013, OnRampAddressSpaceListResponseEnvelopeMessagesCode4014, OnRampAddressSpaceListResponseEnvelopeMessagesCode4015, OnRampAddressSpaceListResponseEnvelopeMessagesCode4016, OnRampAddressSpaceListResponseEnvelopeMessagesCode4017, OnRampAddressSpaceListResponseEnvelopeMessagesCode4018, OnRampAddressSpaceListResponseEnvelopeMessagesCode4019, OnRampAddressSpaceListResponseEnvelopeMessagesCode4020, OnRampAddressSpaceListResponseEnvelopeMessagesCode4021, OnRampAddressSpaceListResponseEnvelopeMessagesCode4022, OnRampAddressSpaceListResponseEnvelopeMessagesCode4023, OnRampAddressSpaceListResponseEnvelopeMessagesCode5001, OnRampAddressSpaceListResponseEnvelopeMessagesCode5002, OnRampAddressSpaceListResponseEnvelopeMessagesCode5003, OnRampAddressSpaceListResponseEnvelopeMessagesCode5004, OnRampAddressSpaceListResponseEnvelopeMessagesCode102000, OnRampAddressSpaceListResponseEnvelopeMessagesCode102001, OnRampAddressSpaceListResponseEnvelopeMessagesCode102002, OnRampAddressSpaceListResponseEnvelopeMessagesCode102003, OnRampAddressSpaceListResponseEnvelopeMessagesCode102004, OnRampAddressSpaceListResponseEnvelopeMessagesCode102005, OnRampAddressSpaceListResponseEnvelopeMessagesCode102006, OnRampAddressSpaceListResponseEnvelopeMessagesCode102007, OnRampAddressSpaceListResponseEnvelopeMessagesCode102008, OnRampAddressSpaceListResponseEnvelopeMessagesCode102009, OnRampAddressSpaceListResponseEnvelopeMessagesCode102010, OnRampAddressSpaceListResponseEnvelopeMessagesCode102011, OnRampAddressSpaceListResponseEnvelopeMessagesCode102012, OnRampAddressSpaceListResponseEnvelopeMessagesCode102013, OnRampAddressSpaceListResponseEnvelopeMessagesCode102014, OnRampAddressSpaceListResponseEnvelopeMessagesCode102015, OnRampAddressSpaceListResponseEnvelopeMessagesCode102016, OnRampAddressSpaceListResponseEnvelopeMessagesCode102017, OnRampAddressSpaceListResponseEnvelopeMessagesCode102018, OnRampAddressSpaceListResponseEnvelopeMessagesCode102019, OnRampAddressSpaceListResponseEnvelopeMessagesCode102020, OnRampAddressSpaceListResponseEnvelopeMessagesCode102021, OnRampAddressSpaceListResponseEnvelopeMessagesCode102022, OnRampAddressSpaceListResponseEnvelopeMessagesCode102023, OnRampAddressSpaceListResponseEnvelopeMessagesCode102024, OnRampAddressSpaceListResponseEnvelopeMessagesCode102025, OnRampAddressSpaceListResponseEnvelopeMessagesCode102026, OnRampAddressSpaceListResponseEnvelopeMessagesCode102027, OnRampAddressSpaceListResponseEnvelopeMessagesCode102028, OnRampAddressSpaceListResponseEnvelopeMessagesCode102029, OnRampAddressSpaceListResponseEnvelopeMessagesCode102030, OnRampAddressSpaceListResponseEnvelopeMessagesCode102031, OnRampAddressSpaceListResponseEnvelopeMessagesCode102032, OnRampAddressSpaceListResponseEnvelopeMessagesCode102033, OnRampAddressSpaceListResponseEnvelopeMessagesCode102034, OnRampAddressSpaceListResponseEnvelopeMessagesCode102035, OnRampAddressSpaceListResponseEnvelopeMessagesCode102036, OnRampAddressSpaceListResponseEnvelopeMessagesCode102037, OnRampAddressSpaceListResponseEnvelopeMessagesCode102038, OnRampAddressSpaceListResponseEnvelopeMessagesCode102039, OnRampAddressSpaceListResponseEnvelopeMessagesCode102040, OnRampAddressSpaceListResponseEnvelopeMessagesCode102041, OnRampAddressSpaceListResponseEnvelopeMessagesCode102042, OnRampAddressSpaceListResponseEnvelopeMessagesCode102043, OnRampAddressSpaceListResponseEnvelopeMessagesCode102044, OnRampAddressSpaceListResponseEnvelopeMessagesCode102045, OnRampAddressSpaceListResponseEnvelopeMessagesCode102046, OnRampAddressSpaceListResponseEnvelopeMessagesCode102047, OnRampAddressSpaceListResponseEnvelopeMessagesCode102048, OnRampAddressSpaceListResponseEnvelopeMessagesCode102049, OnRampAddressSpaceListResponseEnvelopeMessagesCode102050, OnRampAddressSpaceListResponseEnvelopeMessagesCode102051, OnRampAddressSpaceListResponseEnvelopeMessagesCode102052, OnRampAddressSpaceListResponseEnvelopeMessagesCode102053, OnRampAddressSpaceListResponseEnvelopeMessagesCode102054, OnRampAddressSpaceListResponseEnvelopeMessagesCode102055, OnRampAddressSpaceListResponseEnvelopeMessagesCode102056, OnRampAddressSpaceListResponseEnvelopeMessagesCode102057, OnRampAddressSpaceListResponseEnvelopeMessagesCode102058, OnRampAddressSpaceListResponseEnvelopeMessagesCode102059, OnRampAddressSpaceListResponseEnvelopeMessagesCode102060, OnRampAddressSpaceListResponseEnvelopeMessagesCode102061, OnRampAddressSpaceListResponseEnvelopeMessagesCode102062, OnRampAddressSpaceListResponseEnvelopeMessagesCode102063, OnRampAddressSpaceListResponseEnvelopeMessagesCode102064, OnRampAddressSpaceListResponseEnvelopeMessagesCode102065, OnRampAddressSpaceListResponseEnvelopeMessagesCode102066, OnRampAddressSpaceListResponseEnvelopeMessagesCode103001, OnRampAddressSpaceListResponseEnvelopeMessagesCode103002, OnRampAddressSpaceListResponseEnvelopeMessagesCode103003, OnRampAddressSpaceListResponseEnvelopeMessagesCode103004, OnRampAddressSpaceListResponseEnvelopeMessagesCode103005, OnRampAddressSpaceListResponseEnvelopeMessagesCode103006, OnRampAddressSpaceListResponseEnvelopeMessagesCode103007, OnRampAddressSpaceListResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceListResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                 `json:"l10n_key"`
	LoggableError string                                                 `json:"loggable_error"`
	TemplateData  interface{}                                            `json:"template_data"`
	TraceID       string                                                 `json:"trace_id"`
	JSON          onRampAddressSpaceListResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceListResponseEnvelopeMessagesMeta]
type onRampAddressSpaceListResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceListResponseEnvelopeMessagesSource struct {
	Parameter           string                                                   `json:"parameter"`
	ParameterValueIndex int64                                                    `json:"parameter_value_index"`
	Pointer             string                                                   `json:"pointer"`
	JSON                onRampAddressSpaceListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// onRampAddressSpaceListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceListResponseEnvelopeMessagesSource]
type onRampAddressSpaceListResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Prefixes  param.Field[[]string] `json:"prefixes,required"`
}

func (r OnRampAddressSpaceEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OnRampAddressSpaceEditResponseEnvelope struct {
	Errors   []OnRampAddressSpaceEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OnRampAddressSpaceEditResponseEnvelopeMessages `json:"messages,required"`
	Result   OnRampAddressSpaceEditResponse                   `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     onRampAddressSpaceEditResponseEnvelopeJSON       `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [OnRampAddressSpaceEditResponseEnvelope]
type onRampAddressSpaceEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeErrors struct {
	Code             OnRampAddressSpaceEditResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Meta             OnRampAddressSpaceEditResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           OnRampAddressSpaceEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             onRampAddressSpaceEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OnRampAddressSpaceEditResponseEnvelopeErrors]
type onRampAddressSpaceEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeErrorsCode int64

const (
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1001   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1002   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1003   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1004   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1005   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1006   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1007   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1008   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1008
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1009   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1009
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1010   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1010
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1011   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1011
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1012   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1012
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1013   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1013
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1014   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1014
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1015   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1015
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1016   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1016
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode1017   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 1017
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2001   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2002   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2003   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2004   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2005   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2006   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2007   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2008   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2008
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2009   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2009
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2010   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2010
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2011   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2011
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2012   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2012
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2013   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2013
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2014   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2014
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2015   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2015
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2016   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2016
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2017   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2017
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2018   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2018
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2019   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2019
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2020   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2020
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2021   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2021
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode2022   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 2022
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3001   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3002   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3003   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3004   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3005   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3006   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode3007   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 3007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4001   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4002   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4003   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4004   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4005   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4006   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4007   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4008   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4008
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4009   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4009
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4010   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4010
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4011   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4011
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4012   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4012
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4013   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4013
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4014   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4014
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4015   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4015
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4016   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4016
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4017   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4017
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4018   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4018
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4019   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4019
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4020   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4020
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4021   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4021
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4022   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4022
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode4023   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 4023
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode5001   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 5001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode5002   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 5002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode5003   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 5003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode5004   OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 5004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102000 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102000
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102001 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102002 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102003 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102004 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102005 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102006 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102007 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102008 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102008
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102009 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102009
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102010 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102010
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102011 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102011
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102012 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102012
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102013 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102013
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102014 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102014
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102015 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102015
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102016 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102016
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102017 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102017
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102018 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102018
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102019 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102019
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102020 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102020
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102021 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102021
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102022 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102022
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102023 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102023
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102024 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102024
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102025 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102025
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102026 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102026
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102027 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102027
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102028 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102028
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102029 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102029
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102030 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102030
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102031 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102031
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102032 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102032
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102033 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102033
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102034 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102034
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102035 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102035
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102036 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102036
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102037 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102037
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102038 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102038
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102039 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102039
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102040 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102040
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102041 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102041
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102042 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102042
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102043 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102043
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102044 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102044
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102045 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102045
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102046 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102046
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102047 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102047
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102048 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102048
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102049 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102049
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102050 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102050
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102051 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102051
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102052 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102052
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102053 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102053
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102054 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102054
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102055 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102055
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102056 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102056
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102057 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102057
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102058 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102058
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102059 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102059
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102060 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102060
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102061 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102061
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102062 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102062
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102063 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102063
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102064 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102064
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102065 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102065
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode102066 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 102066
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103001 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103001
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103002 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103002
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103003 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103003
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103004 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103004
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103005 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103005
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103006 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103006
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103007 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103007
	OnRampAddressSpaceEditResponseEnvelopeErrorsCode103008 OnRampAddressSpaceEditResponseEnvelopeErrorsCode = 103008
)

func (r OnRampAddressSpaceEditResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceEditResponseEnvelopeErrorsCode1001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1008, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1009, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1010, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1011, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1012, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1013, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1014, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1015, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1016, OnRampAddressSpaceEditResponseEnvelopeErrorsCode1017, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2008, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2009, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2010, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2011, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2012, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2013, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2014, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2015, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2016, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2017, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2018, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2019, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2020, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2021, OnRampAddressSpaceEditResponseEnvelopeErrorsCode2022, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode3007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4008, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4009, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4010, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4011, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4012, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4013, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4014, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4015, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4016, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4017, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4018, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4019, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4020, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4021, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4022, OnRampAddressSpaceEditResponseEnvelopeErrorsCode4023, OnRampAddressSpaceEditResponseEnvelopeErrorsCode5001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode5002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode5003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode5004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102000, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102008, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102009, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102010, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102011, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102012, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102013, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102014, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102015, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102016, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102017, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102018, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102019, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102020, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102021, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102022, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102023, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102024, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102025, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102026, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102027, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102028, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102029, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102030, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102031, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102032, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102033, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102034, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102035, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102036, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102037, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102038, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102039, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102040, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102041, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102042, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102043, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102044, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102045, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102046, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102047, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102048, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102049, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102050, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102051, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102052, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102053, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102054, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102055, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102056, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102057, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102058, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102059, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102060, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102061, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102062, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102063, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102064, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102065, OnRampAddressSpaceEditResponseEnvelopeErrorsCode102066, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103001, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103002, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103003, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103004, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103005, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103006, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103007, OnRampAddressSpaceEditResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceEditResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                               `json:"l10n_key"`
	LoggableError string                                               `json:"loggable_error"`
	TemplateData  interface{}                                          `json:"template_data"`
	TraceID       string                                               `json:"trace_id"`
	JSON          onRampAddressSpaceEditResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceEditResponseEnvelopeErrorsMeta]
type onRampAddressSpaceEditResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeErrorsSource struct {
	Parameter           string                                                 `json:"parameter"`
	ParameterValueIndex int64                                                  `json:"parameter_value_index"`
	Pointer             string                                                 `json:"pointer"`
	JSON                onRampAddressSpaceEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceEditResponseEnvelopeErrorsSource]
type onRampAddressSpaceEditResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeMessages struct {
	Code             OnRampAddressSpaceEditResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Meta             OnRampAddressSpaceEditResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           OnRampAddressSpaceEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             onRampAddressSpaceEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OnRampAddressSpaceEditResponseEnvelopeMessages]
type onRampAddressSpaceEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeMessagesCode int64

const (
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1001   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1002   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1003   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1004   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1005   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1006   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1007   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1008   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1008
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1009   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1009
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1010   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1010
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1011   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1011
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1012   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1012
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1013   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1013
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1014   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1014
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1015   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1015
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1016   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1016
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode1017   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 1017
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2001   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2002   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2003   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2004   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2005   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2006   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2007   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2008   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2008
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2009   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2009
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2010   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2010
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2011   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2011
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2012   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2012
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2013   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2013
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2014   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2014
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2015   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2015
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2016   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2016
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2017   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2017
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2018   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2018
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2019   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2019
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2020   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2020
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2021   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2021
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode2022   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 2022
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3001   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3002   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3003   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3004   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3005   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3006   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode3007   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 3007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4001   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4002   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4003   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4004   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4005   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4006   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4007   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4008   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4008
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4009   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4009
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4010   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4010
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4011   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4011
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4012   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4012
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4013   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4013
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4014   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4014
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4015   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4015
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4016   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4016
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4017   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4017
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4018   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4018
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4019   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4019
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4020   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4020
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4021   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4021
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4022   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4022
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode4023   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 4023
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode5001   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 5001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode5002   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 5002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode5003   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 5003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode5004   OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 5004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102000 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102000
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102001 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102002 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102003 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102004 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102005 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102006 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102007 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102008 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102008
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102009 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102009
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102010 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102010
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102011 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102011
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102012 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102012
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102013 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102013
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102014 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102014
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102015 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102015
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102016 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102016
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102017 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102017
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102018 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102018
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102019 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102019
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102020 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102020
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102021 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102021
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102022 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102022
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102023 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102023
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102024 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102024
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102025 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102025
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102026 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102026
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102027 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102027
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102028 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102028
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102029 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102029
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102030 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102030
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102031 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102031
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102032 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102032
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102033 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102033
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102034 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102034
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102035 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102035
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102036 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102036
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102037 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102037
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102038 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102038
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102039 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102039
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102040 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102040
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102041 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102041
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102042 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102042
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102043 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102043
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102044 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102044
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102045 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102045
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102046 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102046
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102047 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102047
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102048 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102048
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102049 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102049
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102050 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102050
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102051 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102051
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102052 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102052
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102053 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102053
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102054 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102054
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102055 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102055
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102056 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102056
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102057 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102057
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102058 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102058
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102059 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102059
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102060 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102060
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102061 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102061
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102062 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102062
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102063 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102063
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102064 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102064
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102065 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102065
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode102066 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 102066
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103001 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103001
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103002 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103002
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103003 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103003
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103004 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103004
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103005 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103005
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103006 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103006
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103007 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103007
	OnRampAddressSpaceEditResponseEnvelopeMessagesCode103008 OnRampAddressSpaceEditResponseEnvelopeMessagesCode = 103008
)

func (r OnRampAddressSpaceEditResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case OnRampAddressSpaceEditResponseEnvelopeMessagesCode1001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1008, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1009, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1010, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1011, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1012, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1013, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1014, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1015, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1016, OnRampAddressSpaceEditResponseEnvelopeMessagesCode1017, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2008, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2009, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2010, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2011, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2012, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2013, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2014, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2015, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2016, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2017, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2018, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2019, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2020, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2021, OnRampAddressSpaceEditResponseEnvelopeMessagesCode2022, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode3007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4008, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4009, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4010, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4011, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4012, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4013, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4014, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4015, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4016, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4017, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4018, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4019, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4020, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4021, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4022, OnRampAddressSpaceEditResponseEnvelopeMessagesCode4023, OnRampAddressSpaceEditResponseEnvelopeMessagesCode5001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode5002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode5003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode5004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102000, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102008, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102009, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102010, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102011, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102012, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102013, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102014, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102015, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102016, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102017, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102018, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102019, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102020, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102021, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102022, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102023, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102024, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102025, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102026, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102027, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102028, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102029, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102030, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102031, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102032, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102033, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102034, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102035, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102036, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102037, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102038, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102039, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102040, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102041, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102042, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102043, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102044, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102045, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102046, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102047, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102048, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102049, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102050, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102051, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102052, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102053, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102054, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102055, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102056, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102057, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102058, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102059, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102060, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102061, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102062, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102063, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102064, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102065, OnRampAddressSpaceEditResponseEnvelopeMessagesCode102066, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103001, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103002, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103003, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103004, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103005, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103006, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103007, OnRampAddressSpaceEditResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type OnRampAddressSpaceEditResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                 `json:"l10n_key"`
	LoggableError string                                                 `json:"loggable_error"`
	TemplateData  interface{}                                            `json:"template_data"`
	TraceID       string                                                 `json:"trace_id"`
	JSON          onRampAddressSpaceEditResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceEditResponseEnvelopeMessagesMeta]
type onRampAddressSpaceEditResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type OnRampAddressSpaceEditResponseEnvelopeMessagesSource struct {
	Parameter           string                                                   `json:"parameter"`
	ParameterValueIndex int64                                                    `json:"parameter_value_index"`
	Pointer             string                                                   `json:"pointer"`
	JSON                onRampAddressSpaceEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// onRampAddressSpaceEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OnRampAddressSpaceEditResponseEnvelopeMessagesSource]
type onRampAddressSpaceEditResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OnRampAddressSpaceEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r onRampAddressSpaceEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
