// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// ProjectDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDomainService] method instead.
type ProjectDomainService struct {
	Options []option.RequestOption
}

// NewProjectDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectDomainService(opts ...option.RequestOption) (r *ProjectDomainService) {
	r = &ProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *ProjectDomainService) New(ctx context.Context, projectName string, params ProjectDomainNewParams, opts ...option.RequestOption) (res *ProjectDomainNewResponse, err error) {
	var env ProjectDomainNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) List(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProjectDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", query.AccountID, projectName)
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

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) ListAutoPaging(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProjectDomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, projectName, query, opts...))
}

// Delete a Pages project's domain.
func (r *ProjectDomainService) Delete(ctx context.Context, projectName string, domainName string, body ProjectDomainDeleteParams, opts ...option.RequestOption) (res *ProjectDomainDeleteResponse, err error) {
	var env ProjectDomainDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry the validation status of a single domain.
func (r *ProjectDomainService) Edit(ctx context.Context, projectName string, domainName string, params ProjectDomainEditParams, opts ...option.RequestOption) (res *ProjectDomainEditResponse, err error) {
	var env ProjectDomainEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", params.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *ProjectDomainService) Get(ctx context.Context, projectName string, domainName string, query ProjectDomainGetParams, opts ...option.RequestOption) (res *ProjectDomainGetResponse, err error) {
	var env ProjectDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", query.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDomainNewResponse struct {
	ID                   string                                       `json:"id"`
	CertificateAuthority ProjectDomainNewResponseCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                                       `json:"created_on"`
	DomainID             string                                       `json:"domain_id"`
	Name                 string                                       `json:"name"`
	Status               ProjectDomainNewResponseStatus               `json:"status"`
	ValidationData       ProjectDomainNewResponseValidationData       `json:"validation_data"`
	VerificationData     ProjectDomainNewResponseVerificationData     `json:"verification_data"`
	ZoneTag              string                                       `json:"zone_tag"`
	JSON                 projectDomainNewResponseJSON                 `json:"-"`
}

// projectDomainNewResponseJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponse]
type projectDomainNewResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseCertificateAuthority string

const (
	ProjectDomainNewResponseCertificateAuthorityGoogle      ProjectDomainNewResponseCertificateAuthority = "google"
	ProjectDomainNewResponseCertificateAuthorityLetsEncrypt ProjectDomainNewResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainNewResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseCertificateAuthorityGoogle, ProjectDomainNewResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainNewResponseStatus string

const (
	ProjectDomainNewResponseStatusInitializing ProjectDomainNewResponseStatus = "initializing"
	ProjectDomainNewResponseStatusPending      ProjectDomainNewResponseStatus = "pending"
	ProjectDomainNewResponseStatusActive       ProjectDomainNewResponseStatus = "active"
	ProjectDomainNewResponseStatusDeactivated  ProjectDomainNewResponseStatus = "deactivated"
	ProjectDomainNewResponseStatusBlocked      ProjectDomainNewResponseStatus = "blocked"
	ProjectDomainNewResponseStatusError        ProjectDomainNewResponseStatus = "error"
)

func (r ProjectDomainNewResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseStatusInitializing, ProjectDomainNewResponseStatusPending, ProjectDomainNewResponseStatusActive, ProjectDomainNewResponseStatusDeactivated, ProjectDomainNewResponseStatusBlocked, ProjectDomainNewResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainNewResponseValidationData struct {
	ErrorMessage string                                       `json:"error_message"`
	Method       ProjectDomainNewResponseValidationDataMethod `json:"method"`
	Status       ProjectDomainNewResponseValidationDataStatus `json:"status"`
	TXTName      string                                       `json:"txt_name"`
	TXTValue     string                                       `json:"txt_value"`
	JSON         projectDomainNewResponseValidationDataJSON   `json:"-"`
}

// projectDomainNewResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseValidationData]
type projectDomainNewResponseValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainNewResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseValidationDataMethod string

const (
	ProjectDomainNewResponseValidationDataMethodHTTP ProjectDomainNewResponseValidationDataMethod = "http"
	ProjectDomainNewResponseValidationDataMethodTXT  ProjectDomainNewResponseValidationDataMethod = "txt"
)

func (r ProjectDomainNewResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseValidationDataMethodHTTP, ProjectDomainNewResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainNewResponseValidationDataStatus string

const (
	ProjectDomainNewResponseValidationDataStatusInitializing ProjectDomainNewResponseValidationDataStatus = "initializing"
	ProjectDomainNewResponseValidationDataStatusPending      ProjectDomainNewResponseValidationDataStatus = "pending"
	ProjectDomainNewResponseValidationDataStatusActive       ProjectDomainNewResponseValidationDataStatus = "active"
	ProjectDomainNewResponseValidationDataStatusDeactivated  ProjectDomainNewResponseValidationDataStatus = "deactivated"
	ProjectDomainNewResponseValidationDataStatusError        ProjectDomainNewResponseValidationDataStatus = "error"
)

func (r ProjectDomainNewResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseValidationDataStatusInitializing, ProjectDomainNewResponseValidationDataStatusPending, ProjectDomainNewResponseValidationDataStatusActive, ProjectDomainNewResponseValidationDataStatusDeactivated, ProjectDomainNewResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainNewResponseVerificationData struct {
	ErrorMessage string                                         `json:"error_message"`
	Status       ProjectDomainNewResponseVerificationDataStatus `json:"status"`
	JSON         projectDomainNewResponseVerificationDataJSON   `json:"-"`
}

// projectDomainNewResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseVerificationData]
type projectDomainNewResponseVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainNewResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseVerificationDataStatus string

const (
	ProjectDomainNewResponseVerificationDataStatusPending     ProjectDomainNewResponseVerificationDataStatus = "pending"
	ProjectDomainNewResponseVerificationDataStatusActive      ProjectDomainNewResponseVerificationDataStatus = "active"
	ProjectDomainNewResponseVerificationDataStatusDeactivated ProjectDomainNewResponseVerificationDataStatus = "deactivated"
	ProjectDomainNewResponseVerificationDataStatusBlocked     ProjectDomainNewResponseVerificationDataStatus = "blocked"
	ProjectDomainNewResponseVerificationDataStatusError       ProjectDomainNewResponseVerificationDataStatus = "error"
)

func (r ProjectDomainNewResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseVerificationDataStatusPending, ProjectDomainNewResponseVerificationDataStatusActive, ProjectDomainNewResponseVerificationDataStatusDeactivated, ProjectDomainNewResponseVerificationDataStatusBlocked, ProjectDomainNewResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponse struct {
	ID                   string                                        `json:"id"`
	CertificateAuthority ProjectDomainListResponseCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                                        `json:"created_on"`
	DomainID             string                                        `json:"domain_id"`
	Name                 string                                        `json:"name"`
	Status               ProjectDomainListResponseStatus               `json:"status"`
	ValidationData       ProjectDomainListResponseValidationData       `json:"validation_data"`
	VerificationData     ProjectDomainListResponseVerificationData     `json:"verification_data"`
	ZoneTag              string                                        `json:"zone_tag"`
	JSON                 projectDomainListResponseJSON                 `json:"-"`
}

// projectDomainListResponseJSON contains the JSON metadata for the struct
// [ProjectDomainListResponse]
type projectDomainListResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseCertificateAuthority string

const (
	ProjectDomainListResponseCertificateAuthorityGoogle      ProjectDomainListResponseCertificateAuthority = "google"
	ProjectDomainListResponseCertificateAuthorityLetsEncrypt ProjectDomainListResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainListResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseCertificateAuthorityGoogle, ProjectDomainListResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainListResponseStatus string

const (
	ProjectDomainListResponseStatusInitializing ProjectDomainListResponseStatus = "initializing"
	ProjectDomainListResponseStatusPending      ProjectDomainListResponseStatus = "pending"
	ProjectDomainListResponseStatusActive       ProjectDomainListResponseStatus = "active"
	ProjectDomainListResponseStatusDeactivated  ProjectDomainListResponseStatus = "deactivated"
	ProjectDomainListResponseStatusBlocked      ProjectDomainListResponseStatus = "blocked"
	ProjectDomainListResponseStatusError        ProjectDomainListResponseStatus = "error"
)

func (r ProjectDomainListResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseStatusInitializing, ProjectDomainListResponseStatusPending, ProjectDomainListResponseStatusActive, ProjectDomainListResponseStatusDeactivated, ProjectDomainListResponseStatusBlocked, ProjectDomainListResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponseValidationData struct {
	ErrorMessage string                                        `json:"error_message"`
	Method       ProjectDomainListResponseValidationDataMethod `json:"method"`
	Status       ProjectDomainListResponseValidationDataStatus `json:"status"`
	TXTName      string                                        `json:"txt_name"`
	TXTValue     string                                        `json:"txt_value"`
	JSON         projectDomainListResponseValidationDataJSON   `json:"-"`
}

// projectDomainListResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseValidationData]
type projectDomainListResponseValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainListResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseValidationDataMethod string

const (
	ProjectDomainListResponseValidationDataMethodHTTP ProjectDomainListResponseValidationDataMethod = "http"
	ProjectDomainListResponseValidationDataMethodTXT  ProjectDomainListResponseValidationDataMethod = "txt"
)

func (r ProjectDomainListResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseValidationDataMethodHTTP, ProjectDomainListResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainListResponseValidationDataStatus string

const (
	ProjectDomainListResponseValidationDataStatusInitializing ProjectDomainListResponseValidationDataStatus = "initializing"
	ProjectDomainListResponseValidationDataStatusPending      ProjectDomainListResponseValidationDataStatus = "pending"
	ProjectDomainListResponseValidationDataStatusActive       ProjectDomainListResponseValidationDataStatus = "active"
	ProjectDomainListResponseValidationDataStatusDeactivated  ProjectDomainListResponseValidationDataStatus = "deactivated"
	ProjectDomainListResponseValidationDataStatusError        ProjectDomainListResponseValidationDataStatus = "error"
)

func (r ProjectDomainListResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseValidationDataStatusInitializing, ProjectDomainListResponseValidationDataStatusPending, ProjectDomainListResponseValidationDataStatusActive, ProjectDomainListResponseValidationDataStatusDeactivated, ProjectDomainListResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponseVerificationData struct {
	ErrorMessage string                                          `json:"error_message"`
	Status       ProjectDomainListResponseVerificationDataStatus `json:"status"`
	JSON         projectDomainListResponseVerificationDataJSON   `json:"-"`
}

// projectDomainListResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseVerificationData]
type projectDomainListResponseVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainListResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseVerificationDataStatus string

const (
	ProjectDomainListResponseVerificationDataStatusPending     ProjectDomainListResponseVerificationDataStatus = "pending"
	ProjectDomainListResponseVerificationDataStatusActive      ProjectDomainListResponseVerificationDataStatus = "active"
	ProjectDomainListResponseVerificationDataStatusDeactivated ProjectDomainListResponseVerificationDataStatus = "deactivated"
	ProjectDomainListResponseVerificationDataStatusBlocked     ProjectDomainListResponseVerificationDataStatus = "blocked"
	ProjectDomainListResponseVerificationDataStatusError       ProjectDomainListResponseVerificationDataStatus = "error"
)

func (r ProjectDomainListResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseVerificationDataStatusPending, ProjectDomainListResponseVerificationDataStatusActive, ProjectDomainListResponseVerificationDataStatusDeactivated, ProjectDomainListResponseVerificationDataStatusBlocked, ProjectDomainListResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainDeleteResponse = interface{}

type ProjectDomainEditResponse struct {
	ID                   string                                        `json:"id"`
	CertificateAuthority ProjectDomainEditResponseCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                                        `json:"created_on"`
	DomainID             string                                        `json:"domain_id"`
	Name                 string                                        `json:"name"`
	Status               ProjectDomainEditResponseStatus               `json:"status"`
	ValidationData       ProjectDomainEditResponseValidationData       `json:"validation_data"`
	VerificationData     ProjectDomainEditResponseVerificationData     `json:"verification_data"`
	ZoneTag              string                                        `json:"zone_tag"`
	JSON                 projectDomainEditResponseJSON                 `json:"-"`
}

// projectDomainEditResponseJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponse]
type projectDomainEditResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseCertificateAuthority string

const (
	ProjectDomainEditResponseCertificateAuthorityGoogle      ProjectDomainEditResponseCertificateAuthority = "google"
	ProjectDomainEditResponseCertificateAuthorityLetsEncrypt ProjectDomainEditResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainEditResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseCertificateAuthorityGoogle, ProjectDomainEditResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainEditResponseStatus string

const (
	ProjectDomainEditResponseStatusInitializing ProjectDomainEditResponseStatus = "initializing"
	ProjectDomainEditResponseStatusPending      ProjectDomainEditResponseStatus = "pending"
	ProjectDomainEditResponseStatusActive       ProjectDomainEditResponseStatus = "active"
	ProjectDomainEditResponseStatusDeactivated  ProjectDomainEditResponseStatus = "deactivated"
	ProjectDomainEditResponseStatusBlocked      ProjectDomainEditResponseStatus = "blocked"
	ProjectDomainEditResponseStatusError        ProjectDomainEditResponseStatus = "error"
)

func (r ProjectDomainEditResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseStatusInitializing, ProjectDomainEditResponseStatusPending, ProjectDomainEditResponseStatusActive, ProjectDomainEditResponseStatusDeactivated, ProjectDomainEditResponseStatusBlocked, ProjectDomainEditResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainEditResponseValidationData struct {
	ErrorMessage string                                        `json:"error_message"`
	Method       ProjectDomainEditResponseValidationDataMethod `json:"method"`
	Status       ProjectDomainEditResponseValidationDataStatus `json:"status"`
	TXTName      string                                        `json:"txt_name"`
	TXTValue     string                                        `json:"txt_value"`
	JSON         projectDomainEditResponseValidationDataJSON   `json:"-"`
}

// projectDomainEditResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseValidationData]
type projectDomainEditResponseValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainEditResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseValidationDataMethod string

const (
	ProjectDomainEditResponseValidationDataMethodHTTP ProjectDomainEditResponseValidationDataMethod = "http"
	ProjectDomainEditResponseValidationDataMethodTXT  ProjectDomainEditResponseValidationDataMethod = "txt"
)

func (r ProjectDomainEditResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseValidationDataMethodHTTP, ProjectDomainEditResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainEditResponseValidationDataStatus string

const (
	ProjectDomainEditResponseValidationDataStatusInitializing ProjectDomainEditResponseValidationDataStatus = "initializing"
	ProjectDomainEditResponseValidationDataStatusPending      ProjectDomainEditResponseValidationDataStatus = "pending"
	ProjectDomainEditResponseValidationDataStatusActive       ProjectDomainEditResponseValidationDataStatus = "active"
	ProjectDomainEditResponseValidationDataStatusDeactivated  ProjectDomainEditResponseValidationDataStatus = "deactivated"
	ProjectDomainEditResponseValidationDataStatusError        ProjectDomainEditResponseValidationDataStatus = "error"
)

func (r ProjectDomainEditResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseValidationDataStatusInitializing, ProjectDomainEditResponseValidationDataStatusPending, ProjectDomainEditResponseValidationDataStatusActive, ProjectDomainEditResponseValidationDataStatusDeactivated, ProjectDomainEditResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainEditResponseVerificationData struct {
	ErrorMessage string                                          `json:"error_message"`
	Status       ProjectDomainEditResponseVerificationDataStatus `json:"status"`
	JSON         projectDomainEditResponseVerificationDataJSON   `json:"-"`
}

// projectDomainEditResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseVerificationData]
type projectDomainEditResponseVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainEditResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseVerificationDataStatus string

const (
	ProjectDomainEditResponseVerificationDataStatusPending     ProjectDomainEditResponseVerificationDataStatus = "pending"
	ProjectDomainEditResponseVerificationDataStatusActive      ProjectDomainEditResponseVerificationDataStatus = "active"
	ProjectDomainEditResponseVerificationDataStatusDeactivated ProjectDomainEditResponseVerificationDataStatus = "deactivated"
	ProjectDomainEditResponseVerificationDataStatusBlocked     ProjectDomainEditResponseVerificationDataStatus = "blocked"
	ProjectDomainEditResponseVerificationDataStatusError       ProjectDomainEditResponseVerificationDataStatus = "error"
)

func (r ProjectDomainEditResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseVerificationDataStatusPending, ProjectDomainEditResponseVerificationDataStatusActive, ProjectDomainEditResponseVerificationDataStatusDeactivated, ProjectDomainEditResponseVerificationDataStatusBlocked, ProjectDomainEditResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponse struct {
	ID                   string                                       `json:"id"`
	CertificateAuthority ProjectDomainGetResponseCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                                       `json:"created_on"`
	DomainID             string                                       `json:"domain_id"`
	Name                 string                                       `json:"name"`
	Status               ProjectDomainGetResponseStatus               `json:"status"`
	ValidationData       ProjectDomainGetResponseValidationData       `json:"validation_data"`
	VerificationData     ProjectDomainGetResponseVerificationData     `json:"verification_data"`
	ZoneTag              string                                       `json:"zone_tag"`
	JSON                 projectDomainGetResponseJSON                 `json:"-"`
}

// projectDomainGetResponseJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponse]
type projectDomainGetResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseCertificateAuthority string

const (
	ProjectDomainGetResponseCertificateAuthorityGoogle      ProjectDomainGetResponseCertificateAuthority = "google"
	ProjectDomainGetResponseCertificateAuthorityLetsEncrypt ProjectDomainGetResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainGetResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseCertificateAuthorityGoogle, ProjectDomainGetResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainGetResponseStatus string

const (
	ProjectDomainGetResponseStatusInitializing ProjectDomainGetResponseStatus = "initializing"
	ProjectDomainGetResponseStatusPending      ProjectDomainGetResponseStatus = "pending"
	ProjectDomainGetResponseStatusActive       ProjectDomainGetResponseStatus = "active"
	ProjectDomainGetResponseStatusDeactivated  ProjectDomainGetResponseStatus = "deactivated"
	ProjectDomainGetResponseStatusBlocked      ProjectDomainGetResponseStatus = "blocked"
	ProjectDomainGetResponseStatusError        ProjectDomainGetResponseStatus = "error"
)

func (r ProjectDomainGetResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseStatusInitializing, ProjectDomainGetResponseStatusPending, ProjectDomainGetResponseStatusActive, ProjectDomainGetResponseStatusDeactivated, ProjectDomainGetResponseStatusBlocked, ProjectDomainGetResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponseValidationData struct {
	ErrorMessage string                                       `json:"error_message"`
	Method       ProjectDomainGetResponseValidationDataMethod `json:"method"`
	Status       ProjectDomainGetResponseValidationDataStatus `json:"status"`
	TXTName      string                                       `json:"txt_name"`
	TXTValue     string                                       `json:"txt_value"`
	JSON         projectDomainGetResponseValidationDataJSON   `json:"-"`
}

// projectDomainGetResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseValidationData]
type projectDomainGetResponseValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainGetResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseValidationDataMethod string

const (
	ProjectDomainGetResponseValidationDataMethodHTTP ProjectDomainGetResponseValidationDataMethod = "http"
	ProjectDomainGetResponseValidationDataMethodTXT  ProjectDomainGetResponseValidationDataMethod = "txt"
)

func (r ProjectDomainGetResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseValidationDataMethodHTTP, ProjectDomainGetResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainGetResponseValidationDataStatus string

const (
	ProjectDomainGetResponseValidationDataStatusInitializing ProjectDomainGetResponseValidationDataStatus = "initializing"
	ProjectDomainGetResponseValidationDataStatusPending      ProjectDomainGetResponseValidationDataStatus = "pending"
	ProjectDomainGetResponseValidationDataStatusActive       ProjectDomainGetResponseValidationDataStatus = "active"
	ProjectDomainGetResponseValidationDataStatusDeactivated  ProjectDomainGetResponseValidationDataStatus = "deactivated"
	ProjectDomainGetResponseValidationDataStatusError        ProjectDomainGetResponseValidationDataStatus = "error"
)

func (r ProjectDomainGetResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseValidationDataStatusInitializing, ProjectDomainGetResponseValidationDataStatusPending, ProjectDomainGetResponseValidationDataStatusActive, ProjectDomainGetResponseValidationDataStatusDeactivated, ProjectDomainGetResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponseVerificationData struct {
	ErrorMessage string                                         `json:"error_message"`
	Status       ProjectDomainGetResponseVerificationDataStatus `json:"status"`
	JSON         projectDomainGetResponseVerificationDataJSON   `json:"-"`
}

// projectDomainGetResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseVerificationData]
type projectDomainGetResponseVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainGetResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseVerificationDataStatus string

const (
	ProjectDomainGetResponseVerificationDataStatusPending     ProjectDomainGetResponseVerificationDataStatus = "pending"
	ProjectDomainGetResponseVerificationDataStatusActive      ProjectDomainGetResponseVerificationDataStatus = "active"
	ProjectDomainGetResponseVerificationDataStatusDeactivated ProjectDomainGetResponseVerificationDataStatus = "deactivated"
	ProjectDomainGetResponseVerificationDataStatusBlocked     ProjectDomainGetResponseVerificationDataStatus = "blocked"
	ProjectDomainGetResponseVerificationDataStatusError       ProjectDomainGetResponseVerificationDataStatus = "error"
)

func (r ProjectDomainGetResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseVerificationDataStatusPending, ProjectDomainGetResponseVerificationDataStatusActive, ProjectDomainGetResponseVerificationDataStatusDeactivated, ProjectDomainGetResponseVerificationDataStatusBlocked, ProjectDomainGetResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name"`
}

func (r ProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDomainNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ProjectDomainNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// projectDomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponseEnvelope]
type projectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainNewResponseEnvelopeSuccess bool

const (
	ProjectDomainNewResponseEnvelopeSuccessFalse ProjectDomainNewResponseEnvelopeSuccess = false
	ProjectDomainNewResponseEnvelopeSuccessTrue  ProjectDomainNewResponseEnvelopeSuccess = true
)

func (r ProjectDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseEnvelopeSuccessFalse, ProjectDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   ProjectDomainDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDomainDeleteResponseEnvelope]
type projectDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainDeleteResponseEnvelopeSuccess bool

const (
	ProjectDomainDeleteResponseEnvelopeSuccessFalse ProjectDomainDeleteResponseEnvelopeSuccess = false
	ProjectDomainDeleteResponseEnvelopeSuccessTrue  ProjectDomainDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDomainDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainDeleteResponseEnvelopeSuccessFalse, ProjectDomainDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ProjectDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   ProjectDomainEditResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainEditResponseEnvelopeJSON    `json:"-"`
}

// projectDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponseEnvelope]
type projectDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainEditResponseEnvelopeSuccess bool

const (
	ProjectDomainEditResponseEnvelopeSuccessFalse ProjectDomainEditResponseEnvelopeSuccess = false
	ProjectDomainEditResponseEnvelopeSuccessTrue  ProjectDomainEditResponseEnvelopeSuccess = true
)

func (r ProjectDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseEnvelopeSuccessFalse, ProjectDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ProjectDomainGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// projectDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponseEnvelope]
type projectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainGetResponseEnvelopeSuccess bool

const (
	ProjectDomainGetResponseEnvelopeSuccessFalse ProjectDomainGetResponseEnvelopeSuccess = false
	ProjectDomainGetResponseEnvelopeSuccessTrue  ProjectDomainGetResponseEnvelopeSuccess = true
)

func (r ProjectDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseEnvelopeSuccessFalse, ProjectDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
