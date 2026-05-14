// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// DLPEmailAccountMappingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEmailAccountMappingService] method instead.
type DLPEmailAccountMappingService struct {
	Options []option.RequestOption
}

// NewDLPEmailAccountMappingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPEmailAccountMappingService(opts ...option.RequestOption) (r *DLPEmailAccountMappingService) {
	r = &DLPEmailAccountMappingService{}
	r.Options = opts
	return
}

// Create mapping
func (r *DLPEmailAccountMappingService) New(ctx context.Context, params DLPEmailAccountMappingNewParams, opts ...option.RequestOption) (res *DLPEmailAccountMappingNewResponse, err error) {
	var env DLPEmailAccountMappingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/account_mapping", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get mapping
func (r *DLPEmailAccountMappingService) Get(ctx context.Context, query DLPEmailAccountMappingGetParams, opts ...option.RequestOption) (res *DLPEmailAccountMappingGetResponse, err error) {
	var env DLPEmailAccountMappingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/account_mapping", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEmailAccountMappingNewResponse struct {
	AddinIdentifierToken string                                            `json:"addin_identifier_token,required" format:"uuid"`
	AuthRequirements     DLPEmailAccountMappingNewResponseAuthRequirements `json:"auth_requirements,required"`
	JSON                 dlpEmailAccountMappingNewResponseJSON             `json:"-"`
}

// dlpEmailAccountMappingNewResponseJSON contains the JSON metadata for the struct
// [DLPEmailAccountMappingNewResponse]
type dlpEmailAccountMappingNewResponseJSON struct {
	AddinIdentifierToken apijson.Field
	AuthRequirements     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingNewResponseAuthRequirements struct {
	Type DLPEmailAccountMappingNewResponseAuthRequirementsType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	AllowedMicrosoftOrganizations interface{}                                           `json:"allowed_microsoft_organizations"`
	JSON                          dlpEmailAccountMappingNewResponseAuthRequirementsJSON `json:"-"`
	union                         DLPEmailAccountMappingNewResponseAuthRequirementsUnion
}

// dlpEmailAccountMappingNewResponseAuthRequirementsJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingNewResponseAuthRequirements]
type dlpEmailAccountMappingNewResponseAuthRequirementsJSON struct {
	Type                          apijson.Field
	AllowedMicrosoftOrganizations apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r dlpEmailAccountMappingNewResponseAuthRequirementsJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEmailAccountMappingNewResponseAuthRequirements) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEmailAccountMappingNewResponseAuthRequirements{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEmailAccountMappingNewResponseAuthRequirementsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPEmailAccountMappingNewResponseAuthRequirementsObject],
// [DLPEmailAccountMappingNewResponseAuthRequirementsType].
func (r DLPEmailAccountMappingNewResponseAuthRequirements) AsUnion() DLPEmailAccountMappingNewResponseAuthRequirementsUnion {
	return r.union
}

// Union satisfied by [DLPEmailAccountMappingNewResponseAuthRequirementsObject] or
// [DLPEmailAccountMappingNewResponseAuthRequirementsType].
type DLPEmailAccountMappingNewResponseAuthRequirementsUnion interface {
	implementsDLPEmailAccountMappingNewResponseAuthRequirements()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailAccountMappingNewResponseAuthRequirementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailAccountMappingNewResponseAuthRequirementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailAccountMappingNewResponseAuthRequirementsType{}),
		},
	)
}

type DLPEmailAccountMappingNewResponseAuthRequirementsObject struct {
	AllowedMicrosoftOrganizations []string                                                    `json:"allowed_microsoft_organizations,required"`
	Type                          DLPEmailAccountMappingNewResponseAuthRequirementsObjectType `json:"type,required"`
	JSON                          dlpEmailAccountMappingNewResponseAuthRequirementsObjectJSON `json:"-"`
}

// dlpEmailAccountMappingNewResponseAuthRequirementsObjectJSON contains the JSON
// metadata for the struct
// [DLPEmailAccountMappingNewResponseAuthRequirementsObject]
type dlpEmailAccountMappingNewResponseAuthRequirementsObjectJSON struct {
	AllowedMicrosoftOrganizations apijson.Field
	Type                          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseAuthRequirementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseAuthRequirementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEmailAccountMappingNewResponseAuthRequirementsObject) implementsDLPEmailAccountMappingNewResponseAuthRequirements() {
}

type DLPEmailAccountMappingNewResponseAuthRequirementsObjectType string

const (
	DLPEmailAccountMappingNewResponseAuthRequirementsObjectTypeOrg DLPEmailAccountMappingNewResponseAuthRequirementsObjectType = "Org"
)

func (r DLPEmailAccountMappingNewResponseAuthRequirementsObjectType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingNewResponseAuthRequirementsObjectTypeOrg:
		return true
	}
	return false
}

type DLPEmailAccountMappingNewResponseAuthRequirementsType struct {
	Type DLPEmailAccountMappingNewResponseAuthRequirementsTypeType `json:"type,required"`
	JSON dlpEmailAccountMappingNewResponseAuthRequirementsTypeJSON `json:"-"`
}

// dlpEmailAccountMappingNewResponseAuthRequirementsTypeJSON contains the JSON
// metadata for the struct [DLPEmailAccountMappingNewResponseAuthRequirementsType]
type dlpEmailAccountMappingNewResponseAuthRequirementsTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseAuthRequirementsType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseAuthRequirementsTypeJSON) RawJSON() string {
	return r.raw
}

func (r DLPEmailAccountMappingNewResponseAuthRequirementsType) implementsDLPEmailAccountMappingNewResponseAuthRequirements() {
}

type DLPEmailAccountMappingNewResponseAuthRequirementsTypeType string

const (
	DLPEmailAccountMappingNewResponseAuthRequirementsTypeTypeNoAuth DLPEmailAccountMappingNewResponseAuthRequirementsTypeType = "NoAuth"
)

func (r DLPEmailAccountMappingNewResponseAuthRequirementsTypeType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingNewResponseAuthRequirementsTypeTypeNoAuth:
		return true
	}
	return false
}

type DLPEmailAccountMappingGetResponse struct {
	AddinIdentifierToken string                                            `json:"addin_identifier_token,required" format:"uuid"`
	AuthRequirements     DLPEmailAccountMappingGetResponseAuthRequirements `json:"auth_requirements,required"`
	JSON                 dlpEmailAccountMappingGetResponseJSON             `json:"-"`
}

// dlpEmailAccountMappingGetResponseJSON contains the JSON metadata for the struct
// [DLPEmailAccountMappingGetResponse]
type dlpEmailAccountMappingGetResponseJSON struct {
	AddinIdentifierToken apijson.Field
	AuthRequirements     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingGetResponseAuthRequirements struct {
	Type DLPEmailAccountMappingGetResponseAuthRequirementsType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	AllowedMicrosoftOrganizations interface{}                                           `json:"allowed_microsoft_organizations"`
	JSON                          dlpEmailAccountMappingGetResponseAuthRequirementsJSON `json:"-"`
	union                         DLPEmailAccountMappingGetResponseAuthRequirementsUnion
}

// dlpEmailAccountMappingGetResponseAuthRequirementsJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingGetResponseAuthRequirements]
type dlpEmailAccountMappingGetResponseAuthRequirementsJSON struct {
	Type                          apijson.Field
	AllowedMicrosoftOrganizations apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r dlpEmailAccountMappingGetResponseAuthRequirementsJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEmailAccountMappingGetResponseAuthRequirements) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEmailAccountMappingGetResponseAuthRequirements{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEmailAccountMappingGetResponseAuthRequirementsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPEmailAccountMappingGetResponseAuthRequirementsObject],
// [DLPEmailAccountMappingGetResponseAuthRequirementsType].
func (r DLPEmailAccountMappingGetResponseAuthRequirements) AsUnion() DLPEmailAccountMappingGetResponseAuthRequirementsUnion {
	return r.union
}

// Union satisfied by [DLPEmailAccountMappingGetResponseAuthRequirementsObject] or
// [DLPEmailAccountMappingGetResponseAuthRequirementsType].
type DLPEmailAccountMappingGetResponseAuthRequirementsUnion interface {
	implementsDLPEmailAccountMappingGetResponseAuthRequirements()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEmailAccountMappingGetResponseAuthRequirementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailAccountMappingGetResponseAuthRequirementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEmailAccountMappingGetResponseAuthRequirementsType{}),
		},
	)
}

type DLPEmailAccountMappingGetResponseAuthRequirementsObject struct {
	AllowedMicrosoftOrganizations []string                                                    `json:"allowed_microsoft_organizations,required"`
	Type                          DLPEmailAccountMappingGetResponseAuthRequirementsObjectType `json:"type,required"`
	JSON                          dlpEmailAccountMappingGetResponseAuthRequirementsObjectJSON `json:"-"`
}

// dlpEmailAccountMappingGetResponseAuthRequirementsObjectJSON contains the JSON
// metadata for the struct
// [DLPEmailAccountMappingGetResponseAuthRequirementsObject]
type dlpEmailAccountMappingGetResponseAuthRequirementsObjectJSON struct {
	AllowedMicrosoftOrganizations apijson.Field
	Type                          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseAuthRequirementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseAuthRequirementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEmailAccountMappingGetResponseAuthRequirementsObject) implementsDLPEmailAccountMappingGetResponseAuthRequirements() {
}

type DLPEmailAccountMappingGetResponseAuthRequirementsObjectType string

const (
	DLPEmailAccountMappingGetResponseAuthRequirementsObjectTypeOrg DLPEmailAccountMappingGetResponseAuthRequirementsObjectType = "Org"
)

func (r DLPEmailAccountMappingGetResponseAuthRequirementsObjectType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingGetResponseAuthRequirementsObjectTypeOrg:
		return true
	}
	return false
}

type DLPEmailAccountMappingGetResponseAuthRequirementsType struct {
	Type DLPEmailAccountMappingGetResponseAuthRequirementsTypeType `json:"type,required"`
	JSON dlpEmailAccountMappingGetResponseAuthRequirementsTypeJSON `json:"-"`
}

// dlpEmailAccountMappingGetResponseAuthRequirementsTypeJSON contains the JSON
// metadata for the struct [DLPEmailAccountMappingGetResponseAuthRequirementsType]
type dlpEmailAccountMappingGetResponseAuthRequirementsTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseAuthRequirementsType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseAuthRequirementsTypeJSON) RawJSON() string {
	return r.raw
}

func (r DLPEmailAccountMappingGetResponseAuthRequirementsType) implementsDLPEmailAccountMappingGetResponseAuthRequirements() {
}

type DLPEmailAccountMappingGetResponseAuthRequirementsTypeType string

const (
	DLPEmailAccountMappingGetResponseAuthRequirementsTypeTypeNoAuth DLPEmailAccountMappingGetResponseAuthRequirementsTypeType = "NoAuth"
)

func (r DLPEmailAccountMappingGetResponseAuthRequirementsTypeType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingGetResponseAuthRequirementsTypeTypeNoAuth:
		return true
	}
	return false
}

type DLPEmailAccountMappingNewParams struct {
	AccountID        param.Field[string]                                               `path:"account_id,required"`
	AuthRequirements param.Field[DLPEmailAccountMappingNewParamsAuthRequirementsUnion] `json:"auth_requirements,required"`
}

func (r DLPEmailAccountMappingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEmailAccountMappingNewParamsAuthRequirements struct {
	Type                          param.Field[DLPEmailAccountMappingNewParamsAuthRequirementsType] `json:"type,required"`
	AllowedMicrosoftOrganizations param.Field[interface{}]                                         `json:"allowed_microsoft_organizations"`
}

func (r DLPEmailAccountMappingNewParamsAuthRequirements) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEmailAccountMappingNewParamsAuthRequirements) implementsDLPEmailAccountMappingNewParamsAuthRequirementsUnion() {
}

// Satisfied by [zero_trust.DLPEmailAccountMappingNewParamsAuthRequirementsObject],
// [zero_trust.DLPEmailAccountMappingNewParamsAuthRequirementsType],
// [DLPEmailAccountMappingNewParamsAuthRequirements].
type DLPEmailAccountMappingNewParamsAuthRequirementsUnion interface {
	implementsDLPEmailAccountMappingNewParamsAuthRequirementsUnion()
}

type DLPEmailAccountMappingNewParamsAuthRequirementsObject struct {
	AllowedMicrosoftOrganizations param.Field[[]string]                                                  `json:"allowed_microsoft_organizations,required"`
	Type                          param.Field[DLPEmailAccountMappingNewParamsAuthRequirementsObjectType] `json:"type,required"`
}

func (r DLPEmailAccountMappingNewParamsAuthRequirementsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEmailAccountMappingNewParamsAuthRequirementsObject) implementsDLPEmailAccountMappingNewParamsAuthRequirementsUnion() {
}

type DLPEmailAccountMappingNewParamsAuthRequirementsObjectType string

const (
	DLPEmailAccountMappingNewParamsAuthRequirementsObjectTypeOrg DLPEmailAccountMappingNewParamsAuthRequirementsObjectType = "Org"
)

func (r DLPEmailAccountMappingNewParamsAuthRequirementsObjectType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingNewParamsAuthRequirementsObjectTypeOrg:
		return true
	}
	return false
}

type DLPEmailAccountMappingNewParamsAuthRequirementsType struct {
	Type param.Field[DLPEmailAccountMappingNewParamsAuthRequirementsTypeType] `json:"type,required"`
}

func (r DLPEmailAccountMappingNewParamsAuthRequirementsType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEmailAccountMappingNewParamsAuthRequirementsType) implementsDLPEmailAccountMappingNewParamsAuthRequirementsUnion() {
}

type DLPEmailAccountMappingNewParamsAuthRequirementsTypeType string

const (
	DLPEmailAccountMappingNewParamsAuthRequirementsTypeTypeNoAuth DLPEmailAccountMappingNewParamsAuthRequirementsTypeType = "NoAuth"
)

func (r DLPEmailAccountMappingNewParamsAuthRequirementsTypeType) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingNewParamsAuthRequirementsTypeTypeNoAuth:
		return true
	}
	return false
}

type DLPEmailAccountMappingNewResponseEnvelope struct {
	Errors   []DLPEmailAccountMappingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailAccountMappingNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailAccountMappingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailAccountMappingNewResponse                `json:"result"`
	JSON    dlpEmailAccountMappingNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailAccountMappingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEmailAccountMappingNewResponseEnvelope]
type dlpEmailAccountMappingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingNewResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEmailAccountMappingNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailAccountMappingNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailAccountMappingNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingNewResponseEnvelopeErrors]
type dlpEmailAccountMappingNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEmailAccountMappingNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailAccountMappingNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEmailAccountMappingNewResponseEnvelopeErrorsSource]
type dlpEmailAccountMappingNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingNewResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEmailAccountMappingNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailAccountMappingNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailAccountMappingNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingNewResponseEnvelopeMessages]
type dlpEmailAccountMappingNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEmailAccountMappingNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailAccountMappingNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEmailAccountMappingNewResponseEnvelopeMessagesSource]
type dlpEmailAccountMappingNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailAccountMappingNewResponseEnvelopeSuccess bool

const (
	DLPEmailAccountMappingNewResponseEnvelopeSuccessTrue DLPEmailAccountMappingNewResponseEnvelopeSuccess = true
)

func (r DLPEmailAccountMappingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEmailAccountMappingGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEmailAccountMappingGetResponseEnvelope struct {
	Errors   []DLPEmailAccountMappingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEmailAccountMappingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEmailAccountMappingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEmailAccountMappingGetResponse                `json:"result"`
	JSON    dlpEmailAccountMappingGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEmailAccountMappingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEmailAccountMappingGetResponseEnvelope]
type dlpEmailAccountMappingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingGetResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEmailAccountMappingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEmailAccountMappingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEmailAccountMappingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingGetResponseEnvelopeErrors]
type dlpEmailAccountMappingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEmailAccountMappingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEmailAccountMappingGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEmailAccountMappingGetResponseEnvelopeErrorsSource]
type dlpEmailAccountMappingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingGetResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEmailAccountMappingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEmailAccountMappingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEmailAccountMappingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEmailAccountMappingGetResponseEnvelopeMessages]
type dlpEmailAccountMappingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEmailAccountMappingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEmailAccountMappingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEmailAccountMappingGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEmailAccountMappingGetResponseEnvelopeMessagesSource]
type dlpEmailAccountMappingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEmailAccountMappingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEmailAccountMappingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEmailAccountMappingGetResponseEnvelopeSuccess bool

const (
	DLPEmailAccountMappingGetResponseEnvelopeSuccessTrue DLPEmailAccountMappingGetResponseEnvelopeSuccess = true
)

func (r DLPEmailAccountMappingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEmailAccountMappingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
