// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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
)

// ScriptDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptDeploymentService] method instead.
type ScriptDeploymentService struct {
	Options []option.RequestOption
}

// NewScriptDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptDeploymentService(opts ...option.RequestOption) (r *ScriptDeploymentService) {
	r = &ScriptDeploymentService{}
	r.Options = opts
	return
}

// Deployments configure how
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions)
// are deployed to traffic. A deployment can consist of one or two versions of a
// Worker.
func (r *ScriptDeploymentService) New(ctx context.Context, scriptName string, params ScriptDeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ScriptDeploymentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of Worker Deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *ScriptDeploymentService) List(ctx context.Context, scriptName string, query ScriptDeploymentListParams, opts ...option.RequestOption) (res *ScriptDeploymentListResponse, err error) {
	var env ScriptDeploymentListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Worker Deployment. The latest deployment, which is actively serving
// traffic, cannot be deleted. All other deployments can be deleted.
func (r *ScriptDeploymentService) Delete(ctx context.Context, scriptName string, deploymentID string, body ScriptDeploymentDeleteParams, opts ...option.RequestOption) (res *ScriptDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments/%s", body.AccountID, scriptName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get information about a Worker Deployment.
func (r *ScriptDeploymentService) Get(ctx context.Context, scriptName string, deploymentID string, query ScriptDeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env ScriptDeploymentGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments/%s", query.AccountID, scriptName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Deployment struct {
	ID          string                `json:"id,required" format:"uuid"`
	CreatedOn   time.Time             `json:"created_on,required" format:"date-time"`
	Source      string                `json:"source,required"`
	Strategy    DeploymentStrategy    `json:"strategy,required"`
	Versions    []DeploymentVersion   `json:"versions,required"`
	Annotations DeploymentAnnotations `json:"annotations"`
	AuthorEmail string                `json:"author_email" format:"email"`
	JSON        deploymentJSON        `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

type DeploymentStrategy string

const (
	DeploymentStrategyPercentage DeploymentStrategy = "percentage"
)

func (r DeploymentStrategy) IsKnown() bool {
	switch r {
	case DeploymentStrategyPercentage:
		return true
	}
	return false
}

type DeploymentVersion struct {
	Percentage float64               `json:"percentage,required"`
	VersionID  string                `json:"version_id,required" format:"uuid"`
	JSON       deploymentVersionJSON `json:"-"`
}

// deploymentVersionJSON contains the JSON metadata for the struct
// [DeploymentVersion]
type deploymentVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentVersionJSON) RawJSON() string {
	return r.raw
}

type DeploymentAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage string `json:"workers/message"`
	// Operation that triggered the creation of the deployment.
	WorkersTriggeredBy string                    `json:"workers/triggered_by"`
	JSON               deploymentAnnotationsJSON `json:"-"`
}

// deploymentAnnotationsJSON contains the JSON metadata for the struct
// [DeploymentAnnotations]
type deploymentAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DeploymentAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentAnnotationsJSON) RawJSON() string {
	return r.raw
}

type DeploymentParam struct {
	Strategy    param.Field[DeploymentStrategy]         `json:"strategy,required"`
	Versions    param.Field[[]DeploymentVersionParam]   `json:"versions,required"`
	Annotations param.Field[DeploymentAnnotationsParam] `json:"annotations"`
}

func (r DeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentVersionParam struct {
	Percentage param.Field[float64] `json:"percentage,required"`
	VersionID  param.Field[string]  `json:"version_id,required" format:"uuid"`
}

func (r DeploymentVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentAnnotationsParam struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r DeploymentAnnotationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptDeploymentListResponse struct {
	Deployments []Deployment                     `json:"deployments,required"`
	JSON        scriptDeploymentListResponseJSON `json:"-"`
}

// scriptDeploymentListResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentListResponse]
type scriptDeploymentListResponseJSON struct {
	Deployments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponse struct {
	Errors   []ScriptDeploymentDeleteResponseError   `json:"errors,required"`
	Messages []ScriptDeploymentDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptDeploymentDeleteResponseSuccess `json:"success,required"`
	JSON    scriptDeploymentDeleteResponseJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseJSON contains the JSON metadata for the struct
// [ScriptDeploymentDeleteResponse]
type scriptDeploymentDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptDeploymentDeleteResponseErrorsSource `json:"source"`
	JSON             scriptDeploymentDeleteResponseErrorJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ScriptDeploymentDeleteResponseError]
type scriptDeploymentDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptDeploymentDeleteResponseErrorsSourceJSON `json:"-"`
}

// scriptDeploymentDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptDeploymentDeleteResponseErrorsSource]
type scriptDeploymentDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseMessage struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptDeploymentDeleteResponseMessagesSource `json:"source"`
	JSON             scriptDeploymentDeleteResponseMessageJSON    `json:"-"`
}

// scriptDeploymentDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ScriptDeploymentDeleteResponseMessage]
type scriptDeploymentDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentDeleteResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptDeploymentDeleteResponseMessagesSourceJSON `json:"-"`
}

// scriptDeploymentDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptDeploymentDeleteResponseMessagesSource]
type scriptDeploymentDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentDeleteResponseSuccess bool

const (
	ScriptDeploymentDeleteResponseSuccessTrue ScriptDeploymentDeleteResponseSuccess = true
)

func (r ScriptDeploymentDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentNewParams struct {
	// Identifier.
	AccountID  param.Field[string] `path:"account_id,required"`
	Deployment DeploymentParam     `json:"deployment,required"`
	// If set to true, the deployment will be created even if normally blocked by
	// something such rolling back to an older version when a secret has changed.
	Force param.Field[bool] `query:"force"`
}

func (r ScriptDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Deployment)
}

// URLQuery serializes [ScriptDeploymentNewParams]'s query parameters as
// `url.Values`.
func (r ScriptDeploymentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeploymentNewResponseEnvelope struct {
	Errors   []ScriptDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                    `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseEnvelope]
type scriptDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptDeploymentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptDeploymentNewResponseEnvelopeErrors]
type scriptDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptDeploymentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentNewResponseEnvelopeErrorsSource]
type scriptDeploymentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptDeploymentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentNewResponseEnvelopeMessages]
type scriptDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptDeploymentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentNewResponseEnvelopeMessagesSource]
type scriptDeploymentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentNewResponseEnvelopeSuccess bool

const (
	ScriptDeploymentNewResponseEnvelopeSuccessTrue ScriptDeploymentNewResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeploymentListResponseEnvelope struct {
	Errors   []ScriptDeploymentListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptDeploymentListResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptDeploymentListResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptDeploymentListResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptDeploymentListResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentListResponseEnvelope]
type scriptDeploymentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptDeploymentListResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentListResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptDeploymentListResponseEnvelopeErrors]
type scriptDeploymentListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptDeploymentListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentListResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentListResponseEnvelopeErrorsSource]
type scriptDeploymentListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ScriptDeploymentListResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentListResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentListResponseEnvelopeMessages]
type scriptDeploymentListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentListResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    scriptDeploymentListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ScriptDeploymentListResponseEnvelopeMessagesSource]
type scriptDeploymentListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentListResponseEnvelopeSuccess bool

const (
	ScriptDeploymentListResponseEnvelopeSuccessTrue ScriptDeploymentListResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptDeploymentDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeploymentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeploymentGetResponseEnvelope struct {
	Errors   []ScriptDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Deployment                                    `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseEnvelope]
type scriptDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptDeploymentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeploymentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptDeploymentGetResponseEnvelopeErrors]
type scriptDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptDeploymentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentGetResponseEnvelopeErrorsSource]
type scriptDeploymentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptDeploymentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeploymentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptDeploymentGetResponseEnvelopeMessages]
type scriptDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeploymentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptDeploymentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeploymentGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptDeploymentGetResponseEnvelopeMessagesSource]
type scriptDeploymentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeploymentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeploymentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeploymentGetResponseEnvelopeSuccess bool

const (
	ScriptDeploymentGetResponseEnvelopeSuccessTrue ScriptDeploymentGetResponseEnvelopeSuccess = true
)

func (r ScriptDeploymentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeploymentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
