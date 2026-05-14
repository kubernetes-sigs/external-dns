// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// ProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentService] method instead.
type ProjectDeploymentService struct {
	Options []option.RequestOption
	History *ProjectDeploymentHistoryService
}

// NewProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectDeploymentService(opts ...option.RequestOption) (r *ProjectDeploymentService) {
	r = &ProjectDeploymentService{}
	r.Options = opts
	r.History = NewProjectDeploymentHistoryService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *ProjectDeploymentService) New(ctx context.Context, projectName string, params ProjectDeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) List(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Deployment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
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

// Fetch a list of project deployments.
func (r *ProjectDeploymentService) ListAutoPaging(ctx context.Context, projectName string, params ProjectDeploymentListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Deployment] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, projectName, params, opts...))
}

// Delete a deployment.
func (r *ProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, body ProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *ProjectDeploymentDeleteResponse, err error) {
	var env ProjectDeploymentDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch information about a deployment.
func (r *ProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *ProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentRetryParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentRetryResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *ProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, params ProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ProjectDeploymentRollbackResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", params.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDeploymentDeleteResponse = interface{}

type ProjectDeploymentNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r ProjectDeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ProjectDeploymentNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Deployment            `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentNewResponseEnvelope]
type projectDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentNewResponseEnvelopeSuccess bool

const (
	ProjectDeploymentNewResponseEnvelopeSuccessFalse ProjectDeploymentNewResponseEnvelopeSuccess = false
	ProjectDeploymentNewResponseEnvelopeSuccessTrue  ProjectDeploymentNewResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewResponseEnvelopeSuccessFalse, ProjectDeploymentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// What type of deployments to fetch.
	Env param.Field[ProjectDeploymentListParamsEnv] `query:"env"`
}

// URLQuery serializes [ProjectDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// What type of deployments to fetch.
type ProjectDeploymentListParamsEnv string

const (
	ProjectDeploymentListParamsEnvProduction ProjectDeploymentListParamsEnv = "production"
	ProjectDeploymentListParamsEnvPreview    ProjectDeploymentListParamsEnv = "preview"
)

func (r ProjectDeploymentListParamsEnv) IsKnown() bool {
	switch r {
	case ProjectDeploymentListParamsEnvProduction, ProjectDeploymentListParamsEnvPreview:
		return true
	}
	return false
}

type ProjectDeploymentDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   ProjectDeploymentDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDeploymentDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentDeleteResponseEnvelope]
type projectDeploymentDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeploymentDeleteResponseEnvelopeSuccessFalse ProjectDeploymentDeleteResponseEnvelopeSuccess = false
	ProjectDeploymentDeleteResponseEnvelopeSuccessTrue  ProjectDeploymentDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentDeleteResponseEnvelopeSuccessFalse, ProjectDeploymentDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Deployment            `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentGetResponseEnvelope]
type projectDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentGetResponseEnvelopeSuccessFalse ProjectDeploymentGetResponseEnvelopeSuccess = false
	ProjectDeploymentGetResponseEnvelopeSuccessTrue  ProjectDeploymentGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentGetResponseEnvelopeSuccessFalse, ProjectDeploymentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRetryParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ProjectDeploymentRetryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDeploymentRetryResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Deployment            `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentRetryResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRetryResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRetryResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRetryResponseEnvelope]
type projectDeploymentRetryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRetryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRetryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentRetryResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRetryResponseEnvelopeSuccessFalse ProjectDeploymentRetryResponseEnvelopeSuccess = false
	ProjectDeploymentRetryResponseEnvelopeSuccessTrue  ProjectDeploymentRetryResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRetryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRetryResponseEnvelopeSuccessFalse, ProjectDeploymentRetryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDeploymentRollbackParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ProjectDeploymentRollbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Deployment            `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentRollbackResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentRollbackResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentRollbackResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDeploymentRollbackResponseEnvelope]
type projectDeploymentRollbackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentRollbackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentRollbackResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentRollbackResponseEnvelopeSuccess bool

const (
	ProjectDeploymentRollbackResponseEnvelopeSuccessFalse ProjectDeploymentRollbackResponseEnvelopeSuccess = false
	ProjectDeploymentRollbackResponseEnvelopeSuccessTrue  ProjectDeploymentRollbackResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentRollbackResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentRollbackResponseEnvelopeSuccessFalse, ProjectDeploymentRollbackResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
