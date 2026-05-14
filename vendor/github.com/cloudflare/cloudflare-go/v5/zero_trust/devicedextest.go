// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DeviceDEXTestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceDEXTestService] method instead.
type DeviceDEXTestService struct {
	Options []option.RequestOption
}

// NewDeviceDEXTestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceDEXTestService(opts ...option.RequestOption) (r *DeviceDEXTestService) {
	r = &DeviceDEXTestService{}
	r.Options = opts
	return
}

// Create a DEX test.
func (r *DeviceDEXTestService) New(ctx context.Context, params DeviceDEXTestNewParams, opts ...option.RequestOption) (res *DeviceDEXTestNewResponse, err error) {
	var env DeviceDEXTestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX test.
func (r *DeviceDEXTestService) Update(ctx context.Context, dexTestID string, params DeviceDEXTestUpdateParams, opts ...option.RequestOption) (res *DeviceDEXTestUpdateResponse, err error) {
	var env DeviceDEXTestUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", params.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all DEX tests
func (r *DeviceDEXTestService) List(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DeviceDEXTestListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests", query.AccountID)
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

// Fetch all DEX tests
func (r *DeviceDEXTestService) ListAutoPaging(ctx context.Context, query DeviceDEXTestListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DeviceDEXTestListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *DeviceDEXTestService) Delete(ctx context.Context, dexTestID string, body DeviceDEXTestDeleteParams, opts ...option.RequestOption) (res *DeviceDEXTestDeleteResponse, err error) {
	var env DeviceDEXTestDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", body.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single DEX test.
func (r *DeviceDEXTestService) Get(ctx context.Context, dexTestID string, query DeviceDEXTestGetParams, opts ...option.RequestOption) (res *DeviceDEXTestGetResponse, err error) {
	var env DeviceDEXTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/dex_tests/%s", query.AccountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceDEXTestNewResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestNewResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []DeviceDEXTestNewResponseTargetPolicy `json:"target_policies"`
	Targeted       bool                                   `json:"targeted"`
	// The unique identifier for the test.
	TestID string                       `json:"test_id"`
	JSON   deviceDEXTestNewResponseJSON `json:"-"`
}

// deviceDEXTestNewResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponse]
type deviceDEXTestNewResponseJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestNewResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                           `json:"method"`
	JSON   deviceDEXTestNewResponseDataJSON `json:"-"`
}

// deviceDEXTestNewResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponseData]
type deviceDEXTestNewResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseTargetPolicy struct {
	// The id of the DEX rule
	ID string `json:"id"`
	// Whether the DEX rule is the account default
	Default bool `json:"default"`
	// The name of the DEX rule
	Name string                                   `json:"name"`
	JSON deviceDEXTestNewResponseTargetPolicyJSON `json:"-"`
}

// deviceDEXTestNewResponseTargetPolicyJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseTargetPolicy]
type deviceDEXTestNewResponseTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestUpdateResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []DeviceDEXTestUpdateResponseTargetPolicy `json:"target_policies"`
	Targeted       bool                                      `json:"targeted"`
	// The unique identifier for the test.
	TestID string                          `json:"test_id"`
	JSON   deviceDEXTestUpdateResponseJSON `json:"-"`
}

// deviceDEXTestUpdateResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestUpdateResponse]
type deviceDEXTestUpdateResponseJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestUpdateResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                              `json:"method"`
	JSON   deviceDEXTestUpdateResponseDataJSON `json:"-"`
}

// deviceDEXTestUpdateResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestUpdateResponseData]
type deviceDEXTestUpdateResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseTargetPolicy struct {
	// The id of the DEX rule
	ID string `json:"id"`
	// Whether the DEX rule is the account default
	Default bool `json:"default"`
	// The name of the DEX rule
	Name string                                      `json:"name"`
	JSON deviceDEXTestUpdateResponseTargetPolicyJSON `json:"-"`
}

// deviceDEXTestUpdateResponseTargetPolicyJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseTargetPolicy]
type deviceDEXTestUpdateResponseTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestListResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestListResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []DeviceDEXTestListResponseTargetPolicy `json:"target_policies"`
	Targeted       bool                                    `json:"targeted"`
	// The unique identifier for the test.
	TestID string                        `json:"test_id"`
	JSON   deviceDEXTestListResponseJSON `json:"-"`
}

// deviceDEXTestListResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestListResponse]
type deviceDEXTestListResponseJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDEXTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestListResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestListResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                            `json:"method"`
	JSON   deviceDEXTestListResponseDataJSON `json:"-"`
}

// deviceDEXTestListResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestListResponseData]
type deviceDEXTestListResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestListResponseDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestListResponseTargetPolicy struct {
	// The id of the DEX rule
	ID string `json:"id"`
	// Whether the DEX rule is the account default
	Default bool `json:"default"`
	// The name of the DEX rule
	Name string                                    `json:"name"`
	JSON deviceDEXTestListResponseTargetPolicyJSON `json:"-"`
}

// deviceDEXTestListResponseTargetPolicyJSON contains the JSON metadata for the
// struct [DeviceDEXTestListResponseTargetPolicy]
type deviceDEXTestListResponseTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestListResponseTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestListResponseTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponse struct {
	DEXTests []DeviceDEXTestDeleteResponseDEXTest `json:"dex_tests"`
	JSON     deviceDEXTestDeleteResponseJSON      `json:"-"`
}

// deviceDEXTestDeleteResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestDeleteResponse]
type deviceDEXTestDeleteResponseJSON struct {
	DEXTests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseDEXTest struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestDeleteResponseDEXTestsData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []DeviceDEXTestDeleteResponseDEXTestsTargetPolicy `json:"target_policies"`
	Targeted       bool                                              `json:"targeted"`
	// The unique identifier for the test.
	TestID string                                 `json:"test_id"`
	JSON   deviceDEXTestDeleteResponseDEXTestJSON `json:"-"`
}

// deviceDEXTestDeleteResponseDEXTestJSON contains the JSON metadata for the struct
// [DeviceDEXTestDeleteResponseDEXTest]
type deviceDEXTestDeleteResponseDEXTestJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseDEXTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseDEXTestJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestDeleteResponseDEXTestsData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                                      `json:"method"`
	JSON   deviceDEXTestDeleteResponseDEXTestsDataJSON `json:"-"`
}

// deviceDEXTestDeleteResponseDEXTestsDataJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseDEXTestsData]
type deviceDEXTestDeleteResponseDEXTestsDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseDEXTestsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseDEXTestsDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseDEXTestsTargetPolicy struct {
	// The id of the DEX rule
	ID string `json:"id"`
	// Whether the DEX rule is the account default
	Default bool `json:"default"`
	// The name of the DEX rule
	Name string                                              `json:"name"`
	JSON deviceDEXTestDeleteResponseDEXTestsTargetPolicyJSON `json:"-"`
}

// deviceDEXTestDeleteResponseDEXTestsTargetPolicyJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeleteResponseDEXTestsTargetPolicy]
type deviceDEXTestDeleteResponseDEXTestsTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseDEXTestsTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseDEXTestsTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponse struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDEXTestGetResponseData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies []DeviceDEXTestGetResponseTargetPolicy `json:"target_policies"`
	Targeted       bool                                   `json:"targeted"`
	// The unique identifier for the test.
	TestID string                       `json:"test_id"`
	JSON   deviceDEXTestGetResponseJSON `json:"-"`
}

// deviceDEXTestGetResponseJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponse]
type deviceDEXTestGetResponseJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestGetResponseData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                           `json:"method"`
	JSON   deviceDEXTestGetResponseDataJSON `json:"-"`
}

// deviceDEXTestGetResponseDataJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponseData]
type deviceDEXTestGetResponseDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseTargetPolicy struct {
	// The id of the DEX rule
	ID string `json:"id"`
	// Whether the DEX rule is the account default
	Default bool `json:"default"`
	// The name of the DEX rule
	Name string                                   `json:"name"`
	JSON deviceDEXTestGetResponseTargetPolicyJSON `json:"-"`
}

// deviceDEXTestGetResponseTargetPolicyJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseTargetPolicy]
type deviceDEXTestGetResponseTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestNewParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies param.Field[[]DeviceDEXTestNewParamsTargetPolicy] `json:"target_policies"`
	Targeted       param.Field[bool]                                 `json:"targeted"`
}

func (r DeviceDEXTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestNewParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestNewParamsTargetPolicy struct {
	// The id of the DEX rule
	ID param.Field[string] `json:"id"`
	// Whether the DEX rule is the account default
	Default param.Field[bool] `json:"default"`
	// The name of the DEX rule
	Name param.Field[string] `json:"name"`
}

func (r DeviceDEXTestNewParamsTargetPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestNewResponseEnvelope struct {
	Errors   []DeviceDEXTestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DeviceDEXTestNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DeviceDEXTestNewResponse                `json:"result"`
	JSON    deviceDEXTestNewResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestNewResponseEnvelope]
type deviceDEXTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DeviceDEXTestNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeErrors]
type deviceDEXTestNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    deviceDEXTestNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DeviceDEXTestNewResponseEnvelopeErrorsSource]
type deviceDEXTestNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DeviceDEXTestNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestNewResponseEnvelopeMessages]
type deviceDEXTestNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    deviceDEXTestNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestNewResponseEnvelopeMessagesSource]
type deviceDEXTestNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestNewResponseEnvelopeSuccess bool

const (
	DeviceDEXTestNewResponseEnvelopeSuccessTrue DeviceDEXTestNewResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDEXTestUpdateParamsData] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
	// DEX rules targeted by this test
	TargetPolicies param.Field[[]DeviceDEXTestUpdateParamsTargetPolicy] `json:"target_policies"`
	Targeted       param.Field[bool]                                    `json:"targeted"`
}

func (r DeviceDEXTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDEXTestUpdateParamsData struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDEXTestUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestUpdateParamsTargetPolicy struct {
	// The id of the DEX rule
	ID param.Field[string] `json:"id"`
	// Whether the DEX rule is the account default
	Default param.Field[bool] `json:"default"`
	// The name of the DEX rule
	Name param.Field[string] `json:"name"`
}

func (r DeviceDEXTestUpdateParamsTargetPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDEXTestUpdateResponseEnvelope struct {
	Errors   []DeviceDEXTestUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DeviceDEXTestUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DeviceDEXTestUpdateResponse                `json:"result"`
	JSON    deviceDEXTestUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelope]
type deviceDEXTestUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DeviceDEXTestUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestUpdateResponseEnvelopeErrors]
type deviceDEXTestUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestUpdateResponseEnvelopeErrorsSource]
type deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DeviceDEXTestUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestUpdateResponseEnvelopeMessages]
type deviceDEXTestUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestUpdateResponseEnvelopeMessagesSource]
type deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestUpdateResponseEnvelopeSuccess bool

const (
	DeviceDEXTestUpdateResponseEnvelopeSuccessTrue DeviceDEXTestUpdateResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestDeleteResponseEnvelope struct {
	Errors   []DeviceDEXTestDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DeviceDEXTestDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DeviceDEXTestDeleteResponse                `json:"result"`
	JSON    deviceDEXTestDeleteResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelope]
type deviceDEXTestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DeviceDEXTestDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestDeleteResponseEnvelopeErrors]
type deviceDEXTestDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeleteResponseEnvelopeErrorsSource]
type deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DeviceDEXTestDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDEXTestDeleteResponseEnvelopeMessages]
type deviceDEXTestDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestDeleteResponseEnvelopeMessagesSource]
type deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestDeleteResponseEnvelopeSuccess bool

const (
	DeviceDEXTestDeleteResponseEnvelopeSuccessTrue DeviceDEXTestDeleteResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceDEXTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDEXTestGetResponseEnvelope struct {
	Errors   []DeviceDEXTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDEXTestGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DeviceDEXTestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DeviceDEXTestGetResponse                `json:"result"`
	JSON    deviceDEXTestGetResponseEnvelopeJSON    `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDEXTestGetResponseEnvelope]
type deviceDEXTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DeviceDEXTestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             deviceDEXTestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeErrors]
type deviceDEXTestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    deviceDEXTestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DeviceDEXTestGetResponseEnvelopeErrorsSource]
type deviceDEXTestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DeviceDEXTestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             deviceDEXTestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDEXTestGetResponseEnvelopeMessages]
type deviceDEXTestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDEXTestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    deviceDEXTestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// deviceDEXTestGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceDEXTestGetResponseEnvelopeMessagesSource]
type deviceDEXTestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDEXTestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDEXTestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceDEXTestGetResponseEnvelopeSuccess bool

const (
	DeviceDEXTestGetResponseEnvelopeSuccessTrue DeviceDEXTestGetResponseEnvelopeSuccess = true
)

func (r DeviceDEXTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceDEXTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
