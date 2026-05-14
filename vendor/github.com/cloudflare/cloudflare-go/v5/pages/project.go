// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// ProjectService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options     []option.RequestOption
	Deployments *ProjectDeploymentService
	Domains     *ProjectDomainService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Deployments = NewProjectDeploymentService(opts...)
	r.Domains = NewProjectDomainService(opts...)
	return
}

// Create a new project.
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all user projects.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Deployment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", query.AccountID)
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

// Fetch a list of all user projects.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Deployment] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a project by name.
func (r *ProjectService) Delete(ctx context.Context, projectName string, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	var env ProjectDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a project by name.
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Purge all cached build artifacts for a Pages project
func (r *ProjectService) PurgeBuildCache(ctx context.Context, projectName string, body ProjectPurgeBuildCacheParams, opts ...option.RequestOption) (res *ProjectPurgeBuildCacheResponse, err error) {
	var env ProjectPurgeBuildCacheResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Deployment struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,nullable"`
	// Configs for the project build process.
	BuildConfig DeploymentBuildConfig `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentDeploymentTrigger `json:"deployment_trigger"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]DeploymentEnvVar `json:"env_vars"`
	// Type of deploy.
	Environment DeploymentEnvironment `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string           `json:"short_id"`
	Source  DeploymentSource `json:"source"`
	// List of past stages.
	Stages []Stage `json:"stages"`
	// The live URL to view this deployment.
	URL  string         `json:"url"`
	JSON deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type DeploymentBuildConfig struct {
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string `json:"root_dir,nullable"`
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string                    `json:"web_analytics_token,nullable"`
	JSON              deploymentBuildConfigJSON `json:"-"`
}

// deploymentBuildConfigJSON contains the JSON metadata for the struct
// [DeploymentBuildConfig]
type deploymentBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type DeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type DeploymentDeploymentTriggerType `json:"type"`
	JSON deploymentDeploymentTriggerJSON `json:"-"`
}

// deploymentDeploymentTriggerJSON contains the JSON metadata for the struct
// [DeploymentDeploymentTrigger]
type deploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type DeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                  `json:"commit_message"`
	JSON          deploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentDeploymentTriggerMetadataJSON contains the JSON metadata for the
// struct [DeploymentDeploymentTriggerMetadata]
type deploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type DeploymentDeploymentTriggerType string

const (
	DeploymentDeploymentTriggerTypePush  DeploymentDeploymentTriggerType = "push"
	DeploymentDeploymentTriggerTypeADHoc DeploymentDeploymentTriggerType = "ad_hoc"
)

func (r DeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case DeploymentDeploymentTriggerTypePush, DeploymentDeploymentTriggerTypeADHoc:
		return true
	}
	return false
}

// A plaintext environment variable.
type DeploymentEnvVar struct {
	Type DeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string               `json:"value,required"`
	JSON  deploymentEnvVarJSON `json:"-"`
	union DeploymentEnvVarsUnion
}

// deploymentEnvVarJSON contains the JSON metadata for the struct
// [DeploymentEnvVar]
type deploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r deploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *DeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = DeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeploymentEnvVarsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DeploymentEnvVarsPagesPlainTextEnvVar],
// [DeploymentEnvVarsPagesSecretTextEnvVar].
func (r DeploymentEnvVar) AsUnion() DeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [DeploymentEnvVarsPagesPlainTextEnvVar] or
// [DeploymentEnvVarsPagesSecretTextEnvVar].
type DeploymentEnvVarsUnion interface {
	implementsDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type DeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type DeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                    `json:"value,required"`
	JSON  deploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// deploymentEnvVarsPagesPlainTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentEnvVarsPagesPlainTextEnvVar]
type deploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentEnvVarsPagesPlainTextEnvVar) implementsDeploymentEnvVar() {}

type DeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	DeploymentEnvVarsPagesPlainTextEnvVarTypePlainText DeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r DeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type DeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type DeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                     `json:"value,required"`
	JSON  deploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// deploymentEnvVarsPagesSecretTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentEnvVarsPagesSecretTextEnvVar]
type deploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentEnvVarsPagesSecretTextEnvVar) implementsDeploymentEnvVar() {}

type DeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	DeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText DeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r DeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type DeploymentEnvVarsType string

const (
	DeploymentEnvVarsTypePlainText  DeploymentEnvVarsType = "plain_text"
	DeploymentEnvVarsTypeSecretText DeploymentEnvVarsType = "secret_text"
)

func (r DeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsTypePlainText, DeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type DeploymentEnvironment string

const (
	DeploymentEnvironmentPreview    DeploymentEnvironment = "preview"
	DeploymentEnvironmentProduction DeploymentEnvironment = "production"
)

func (r DeploymentEnvironment) IsKnown() bool {
	switch r {
	case DeploymentEnvironmentPreview, DeploymentEnvironmentProduction:
		return true
	}
	return false
}

type DeploymentSource struct {
	Config DeploymentSourceConfig `json:"config"`
	Type   string                 `json:"type"`
	JSON   deploymentSourceJSON   `json:"-"`
}

// deploymentSourceJSON contains the JSON metadata for the struct
// [DeploymentSource]
type deploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceJSON) RawJSON() string {
	return r.raw
}

type DeploymentSourceConfig struct {
	DeploymentsEnabled           bool                                           `json:"deployments_enabled"`
	Owner                        string                                         `json:"owner"`
	PathExcludes                 []string                                       `json:"path_excludes"`
	PathIncludes                 []string                                       `json:"path_includes"`
	PrCommentsEnabled            bool                                           `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                       `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                       `json:"preview_branch_includes"`
	PreviewDeploymentSetting     DeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                         `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                           `json:"production_deployments_enabled"`
	RepoName                     string                                         `json:"repo_name"`
	JSON                         deploymentSourceConfigJSON                     `json:"-"`
}

// deploymentSourceConfigJSON contains the JSON metadata for the struct
// [DeploymentSourceConfig]
type deploymentSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

type DeploymentSourceConfigPreviewDeploymentSetting string

const (
	DeploymentSourceConfigPreviewDeploymentSettingAll    DeploymentSourceConfigPreviewDeploymentSetting = "all"
	DeploymentSourceConfigPreviewDeploymentSettingNone   DeploymentSourceConfigPreviewDeploymentSetting = "none"
	DeploymentSourceConfigPreviewDeploymentSettingCustom DeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r DeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case DeploymentSourceConfigPreviewDeploymentSettingAll, DeploymentSourceConfigPreviewDeploymentSettingNone, DeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type DeploymentParam struct {
	// Configs for the project build process.
	BuildConfig param.Field[DeploymentBuildConfigParam] `json:"build_config"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]DeploymentEnvVarsUnionParam] `json:"env_vars"`
	Source  param.Field[DeploymentSourceParam]                  `json:"source"`
}

func (r DeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type DeploymentBuildConfigParam struct {
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r DeploymentBuildConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type DeploymentDeploymentTriggerParam struct {
	// Additional info about the trigger.
	Metadata param.Field[DeploymentDeploymentTriggerMetadataParam] `json:"metadata"`
}

func (r DeploymentDeploymentTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type DeploymentDeploymentTriggerMetadataParam struct {
}

func (r DeploymentDeploymentTriggerMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type DeploymentEnvVarParam struct {
	Type param.Field[DeploymentEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentEnvVarParam) implementsDeploymentEnvVarsUnionParam() {}

// A plaintext environment variable.
//
// Satisfied by [pages.DeploymentEnvVarsPagesPlainTextEnvVarParam],
// [pages.DeploymentEnvVarsPagesSecretTextEnvVarParam], [DeploymentEnvVarParam].
type DeploymentEnvVarsUnionParam interface {
	implementsDeploymentEnvVarsUnionParam()
}

// A plaintext environment variable.
type DeploymentEnvVarsPagesPlainTextEnvVarParam struct {
	Type param.Field[DeploymentEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentEnvVarsPagesPlainTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentEnvVarsPagesPlainTextEnvVarParam) implementsDeploymentEnvVarsUnionParam() {}

// An encrypted environment variable.
type DeploymentEnvVarsPagesSecretTextEnvVarParam struct {
	Type param.Field[DeploymentEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentEnvVarsPagesSecretTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentEnvVarsPagesSecretTextEnvVarParam) implementsDeploymentEnvVarsUnionParam() {}

type DeploymentSourceParam struct {
	Config param.Field[DeploymentSourceConfigParam] `json:"config"`
	Type   param.Field[string]                      `json:"type"`
}

func (r DeploymentSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentSourceConfigParam struct {
	DeploymentsEnabled           param.Field[bool]                                           `json:"deployments_enabled"`
	Owner                        param.Field[string]                                         `json:"owner"`
	PathExcludes                 param.Field[[]string]                                       `json:"path_excludes"`
	PathIncludes                 param.Field[[]string]                                       `json:"path_includes"`
	PrCommentsEnabled            param.Field[bool]                                           `json:"pr_comments_enabled"`
	PreviewBranchExcludes        param.Field[[]string]                                       `json:"preview_branch_excludes"`
	PreviewBranchIncludes        param.Field[[]string]                                       `json:"preview_branch_includes"`
	PreviewDeploymentSetting     param.Field[DeploymentSourceConfigPreviewDeploymentSetting] `json:"preview_deployment_setting"`
	ProductionBranch             param.Field[string]                                         `json:"production_branch"`
	ProductionDeploymentsEnabled param.Field[bool]                                           `json:"production_deployments_enabled"`
	RepoName                     param.Field[string]                                         `json:"repo_name"`
}

func (r DeploymentSourceConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Project struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig ProjectBuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment Deployment `json:"canonical_deployment,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment Deployment `json:"latest_deployment,nullable"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string        `json:"production_branch"`
	Source           ProjectSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string      `json:"subdomain"`
	JSON      projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID                  apijson.Field
	BuildConfig         apijson.Field
	CanonicalDeployment apijson.Field
	CreatedOn           apijson.Field
	DeploymentConfigs   apijson.Field
	Domains             apijson.Field
	LatestDeployment    apijson.Field
	Name                apijson.Field
	ProductionBranch    apijson.Field
	Source              apijson.Field
	Subdomain           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectBuildConfig struct {
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string `json:"root_dir,nullable"`
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string                 `json:"web_analytics_token,nullable"`
	JSON              projectBuildConfigJSON `json:"-"`
}

// projectBuildConfigJSON contains the JSON metadata for the struct
// [ProjectBuildConfig]
type projectBuildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectDeploymentConfigsPreview `json:"preview"`
	// Configs for production deploys.
	Production ProjectDeploymentConfigsProduction `json:"production"`
	JSON       projectDeploymentConfigsJSON       `json:"-"`
}

// projectDeploymentConfigsJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigs]
type projectDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectDeploymentConfigsPreviewAIBinding `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectDeploymentConfigsPreviewBrowser `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectDeploymentConfigsPreviewD1Database `json:"d1_databases,nullable"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces,nullable"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentConfigsPreviewEnvVar `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces,nullable"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsPreviewPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectDeploymentConfigsPreviewQueueProducer `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectDeploymentConfigsPreviewR2Bucket `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services map[string]ProjectDeploymentConfigsPreviewService `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings,nullable"`
	JSON              projectDeploymentConfigsPreviewJSON                        `json:"-"`
}

// projectDeploymentConfigsPreviewJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsPreview]
type projectDeploymentConfigsPreviewJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	Browsers                apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	HyperdriveBindings      apijson.Field
	KVNamespaces            apijson.Field
	MTLSCertificates        apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	Services                apijson.Field
	VectorizeBindings       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                       `json:"project_id"`
	JSON      projectDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAIBindingJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewAIBinding]
type projectDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                    `json:"dataset"`
	JSON    projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectDeploymentConfigsPreviewBrowser struct {
	JSON projectDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectDeploymentConfigsPreviewBrowserJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewBrowser]
type projectDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                        `json:"id"`
	JSON projectDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectDeploymentConfigsPreviewD1DatabaseJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewD1Database]
type projectDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                    `json:"namespace_id"`
	JSON        projectDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewDurableObjectNamespaceJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewDurableObjectNamespace]
type projectDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                    `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectDeploymentConfigsPreviewEnvVarsUnion
}

// projectDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewEnvVar]
type projectDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentConfigsPreviewEnvVarsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentConfigsPreviewEnvVar) AsUnion() ProjectDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
// or [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                         `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentConfigsPreviewEnvVar() {
}

type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                          `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentConfigsPreviewEnvVar() {
}

type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                               `json:"id"`
	JSON projectDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewHyperdriveBindingJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewHyperdriveBinding]
type projectDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                         `json:"namespace_id"`
	JSON        projectDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewKVNamespaceJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewKVNamespace]
type projectDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                             `json:"certificate_id"`
	JSON          projectDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewMTLSCertificate]
type projectDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                       `json:"mode"`
	JSON projectDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectDeploymentConfigsPreviewPlacementJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewPlacement]
type projectDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                           `json:"name"`
	JSON projectDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectDeploymentConfigsPreviewQueueProducerJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewQueueProducer]
type projectDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsPreviewR2Bucket struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                      `json:"name"`
	JSON projectDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectDeploymentConfigsPreviewR2BucketJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewR2Bucket]
type projectDeploymentConfigsPreviewR2BucketJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsPreviewService struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                     `json:"service"`
	JSON    projectDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewServiceJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewService]
type projectDeploymentConfigsPreviewServiceJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                              `json:"index_name"`
	JSON      projectDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewVectorizeBindingJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewVectorizeBinding]
type projectDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectDeploymentConfigsProductionAIBinding `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectDeploymentConfigsProductionBrowser `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectDeploymentConfigsProductionD1Database `json:"d1_databases,nullable"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces,nullable"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentConfigsProductionEnvVar `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectDeploymentConfigsProductionKVNamespace `json:"kv_namespaces,nullable"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsProductionPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectDeploymentConfigsProductionQueueProducer `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectDeploymentConfigsProductionR2Bucket `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services map[string]ProjectDeploymentConfigsProductionService `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings,nullable"`
	JSON              projectDeploymentConfigsProductionJSON                        `json:"-"`
}

// projectDeploymentConfigsProductionJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsProduction]
type projectDeploymentConfigsProductionJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	Browsers                apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	HyperdriveBindings      apijson.Field
	KVNamespaces            apijson.Field
	MTLSCertificates        apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	Services                apijson.Field
	VectorizeBindings       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type ProjectDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                          `json:"project_id"`
	JSON      projectDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionAIBindingJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionAIBinding]
type projectDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                       `json:"dataset"`
	JSON    projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionAnalyticsEngineDataset]
type projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectDeploymentConfigsProductionBrowser struct {
	JSON projectDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectDeploymentConfigsProductionBrowserJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionBrowser]
type projectDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                           `json:"id"`
	JSON projectDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectDeploymentConfigsProductionD1DatabaseJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionD1Database]
type projectDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                       `json:"namespace_id"`
	JSON        projectDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsProductionDurableObjectNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionDurableObjectNamespace]
type projectDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                       `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectDeploymentConfigsProductionEnvVarsUnion
}

// projectDeploymentConfigsProductionEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionEnvVar]
type projectDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentConfigsProductionEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentConfigsProductionEnvVar) AsUnion() ProjectDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                            `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentConfigsProductionEnvVar() {
}

type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                             `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentConfigsProductionEnvVar() {
}

type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentConfigsProductionEnvVarsType string

const (
	ProjectDeploymentConfigsProductionEnvVarsTypePlainText  ProjectDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectDeploymentConfigsProductionEnvVarsTypeSecretText ProjectDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsTypePlainText, ProjectDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                  `json:"id"`
	JSON projectDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionHyperdriveBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionHyperdriveBinding]
type projectDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                            `json:"namespace_id"`
	JSON        projectDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsProductionKVNamespaceJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionKVNamespace]
type projectDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                `json:"certificate_id"`
	JSON          projectDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectDeploymentConfigsProductionMTLSCertificateJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionMTLSCertificate]
type projectDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                          `json:"mode"`
	JSON projectDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectDeploymentConfigsProductionPlacementJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionPlacement]
type projectDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                              `json:"name"`
	JSON projectDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectDeploymentConfigsProductionQueueProducerJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionQueueProducer]
type projectDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsProductionR2Bucket struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                                         `json:"name"`
	JSON projectDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectDeploymentConfigsProductionR2BucketJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionR2Bucket]
type projectDeploymentConfigsProductionR2BucketJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsProductionService struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                                        `json:"service"`
	JSON    projectDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectDeploymentConfigsProductionServiceJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionService]
type projectDeploymentConfigsProductionServiceJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                 `json:"index_name"`
	JSON      projectDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionVectorizeBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionVectorizeBinding]
type projectDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

type ProjectSource struct {
	Config ProjectSourceConfig `json:"config"`
	Type   string              `json:"type"`
	JSON   projectSourceJSON   `json:"-"`
}

// projectSourceJSON contains the JSON metadata for the struct [ProjectSource]
type projectSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectSourceConfig struct {
	DeploymentsEnabled           bool                                        `json:"deployments_enabled"`
	Owner                        string                                      `json:"owner"`
	PathExcludes                 []string                                    `json:"path_excludes"`
	PathIncludes                 []string                                    `json:"path_includes"`
	PrCommentsEnabled            bool                                        `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                                    `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                                    `json:"preview_branch_includes"`
	PreviewDeploymentSetting     ProjectSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                                      `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                        `json:"production_deployments_enabled"`
	RepoName                     string                                      `json:"repo_name"`
	JSON                         projectSourceConfigJSON                     `json:"-"`
}

// projectSourceConfigJSON contains the JSON metadata for the struct
// [ProjectSourceConfig]
type projectSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSourceConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectSourceConfigPreviewDeploymentSetting string

const (
	ProjectSourceConfigPreviewDeploymentSettingAll    ProjectSourceConfigPreviewDeploymentSetting = "all"
	ProjectSourceConfigPreviewDeploymentSettingNone   ProjectSourceConfigPreviewDeploymentSetting = "none"
	ProjectSourceConfigPreviewDeploymentSettingCustom ProjectSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectSourceConfigPreviewDeploymentSettingAll, ProjectSourceConfigPreviewDeploymentSettingNone, ProjectSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type ProjectParam struct {
	// Configs for the project build process.
	BuildConfig param.Field[ProjectBuildConfigParam] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectDeploymentConfigsParam] `json:"deployment_configs"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string]             `json:"production_branch"`
	Source           param.Field[ProjectSourceParam] `json:"source"`
}

func (r ProjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectBuildConfigParam struct {
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r ProjectBuildConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectDeploymentConfigsParam struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectDeploymentConfigsPreviewParam] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectDeploymentConfigsProductionParam] `json:"production"`
}

func (r ProjectDeploymentConfigsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectDeploymentConfigsPreviewParam struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectDeploymentConfigsPreviewAIBindingParam] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectDeploymentConfigsPreviewBrowserParam] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectDeploymentConfigsPreviewD1DatabaseParam] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectDeploymentConfigsPreviewEnvVarsUnionParam] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectDeploymentConfigsPreviewHyperdriveBindingParam] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectDeploymentConfigsPreviewKVNamespaceParam] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectDeploymentConfigsPreviewMTLSCertificateParam] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectDeploymentConfigsPreviewPlacementParam] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectDeploymentConfigsPreviewQueueProducerParam] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectDeploymentConfigsPreviewR2BucketParam] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectDeploymentConfigsPreviewServiceParam] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectDeploymentConfigsPreviewVectorizeBindingParam] `json:"vectorize_bindings"`
}

func (r ProjectDeploymentConfigsPreviewParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectDeploymentConfigsPreviewAIBindingParam struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectDeploymentConfigsPreviewAIBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectDeploymentConfigsPreviewAnalyticsEngineDatasetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectDeploymentConfigsPreviewBrowserParam struct {
}

func (r ProjectDeploymentConfigsPreviewBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectDeploymentConfigsPreviewD1DatabaseParam struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsPreviewD1DatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsPreviewDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsPreviewEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsPreviewEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsPreviewEnvVarParam) implementsProjectDeploymentConfigsPreviewEnvVarsUnionParam() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarParam],
// [pages.ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarParam],
// [ProjectDeploymentConfigsPreviewEnvVarParam].
type ProjectDeploymentConfigsPreviewEnvVarsUnionParam interface {
	implementsProjectDeploymentConfigsPreviewEnvVarsUnionParam()
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarParam) implementsProjectDeploymentConfigsPreviewEnvVarsUnionParam() {
}

// An encrypted environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarParam) implementsProjectDeploymentConfigsPreviewEnvVarsUnionParam() {
}

// Hyperdrive binding.
type ProjectDeploymentConfigsPreviewHyperdriveBindingParam struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsPreviewHyperdriveBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectDeploymentConfigsPreviewKVNamespaceParam struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsPreviewKVNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectDeploymentConfigsPreviewMTLSCertificateParam struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectDeploymentConfigsPreviewMTLSCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsPreviewPlacementParam struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectDeploymentConfigsPreviewPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectDeploymentConfigsPreviewQueueProducerParam struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsPreviewQueueProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectDeploymentConfigsPreviewR2BucketParam struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsPreviewR2BucketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectDeploymentConfigsPreviewServiceParam struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectDeploymentConfigsPreviewServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectDeploymentConfigsPreviewVectorizeBindingParam struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectDeploymentConfigsPreviewVectorizeBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectDeploymentConfigsProductionParam struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectDeploymentConfigsProductionAIBindingParam] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectDeploymentConfigsProductionBrowserParam] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectDeploymentConfigsProductionD1DatabaseParam] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectDeploymentConfigsProductionDurableObjectNamespaceParam] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectDeploymentConfigsProductionEnvVarsUnionParam] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectDeploymentConfigsProductionHyperdriveBindingParam] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectDeploymentConfigsProductionKVNamespaceParam] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectDeploymentConfigsProductionMTLSCertificateParam] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectDeploymentConfigsProductionPlacementParam] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectDeploymentConfigsProductionQueueProducerParam] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectDeploymentConfigsProductionR2BucketParam] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectDeploymentConfigsProductionServiceParam] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectDeploymentConfigsProductionVectorizeBindingParam] `json:"vectorize_bindings"`
}

func (r ProjectDeploymentConfigsProductionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectDeploymentConfigsProductionAIBindingParam struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r ProjectDeploymentConfigsProductionAIBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r ProjectDeploymentConfigsProductionAnalyticsEngineDatasetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectDeploymentConfigsProductionBrowserParam struct {
}

func (r ProjectDeploymentConfigsProductionBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectDeploymentConfigsProductionD1DatabaseParam struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsProductionD1DatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectDeploymentConfigsProductionDurableObjectNamespaceParam struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsProductionDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsProductionEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsProductionEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsProductionEnvVarParam) implementsProjectDeploymentConfigsProductionEnvVarsUnionParam() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarParam],
// [pages.ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarParam],
// [ProjectDeploymentConfigsProductionEnvVarParam].
type ProjectDeploymentConfigsProductionEnvVarsUnionParam interface {
	implementsProjectDeploymentConfigsProductionEnvVarsUnionParam()
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarParam) implementsProjectDeploymentConfigsProductionEnvVarsUnionParam() {
}

// An encrypted environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarParam struct {
	Type param.Field[ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarParam) implementsProjectDeploymentConfigsProductionEnvVarsUnionParam() {
}

// Hyperdrive binding.
type ProjectDeploymentConfigsProductionHyperdriveBindingParam struct {
	ID param.Field[string] `json:"id"`
}

func (r ProjectDeploymentConfigsProductionHyperdriveBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectDeploymentConfigsProductionKVNamespaceParam struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r ProjectDeploymentConfigsProductionKVNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectDeploymentConfigsProductionMTLSCertificateParam struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ProjectDeploymentConfigsProductionMTLSCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsProductionPlacementParam struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r ProjectDeploymentConfigsProductionPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectDeploymentConfigsProductionQueueProducerParam struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsProductionQueueProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectDeploymentConfigsProductionR2BucketParam struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r ProjectDeploymentConfigsProductionR2BucketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectDeploymentConfigsProductionServiceParam struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r ProjectDeploymentConfigsProductionServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type ProjectDeploymentConfigsProductionVectorizeBindingParam struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r ProjectDeploymentConfigsProductionVectorizeBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectSourceParam struct {
	Config param.Field[ProjectSourceConfigParam] `json:"config"`
	Type   param.Field[string]                   `json:"type"`
}

func (r ProjectSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectSourceConfigParam struct {
	DeploymentsEnabled           param.Field[bool]                                        `json:"deployments_enabled"`
	Owner                        param.Field[string]                                      `json:"owner"`
	PathExcludes                 param.Field[[]string]                                    `json:"path_excludes"`
	PathIncludes                 param.Field[[]string]                                    `json:"path_includes"`
	PrCommentsEnabled            param.Field[bool]                                        `json:"pr_comments_enabled"`
	PreviewBranchExcludes        param.Field[[]string]                                    `json:"preview_branch_excludes"`
	PreviewBranchIncludes        param.Field[[]string]                                    `json:"preview_branch_includes"`
	PreviewDeploymentSetting     param.Field[ProjectSourceConfigPreviewDeploymentSetting] `json:"preview_deployment_setting"`
	ProductionBranch             param.Field[string]                                      `json:"production_branch"`
	ProductionDeploymentsEnabled param.Field[bool]                                        `json:"production_deployments_enabled"`
	RepoName                     param.Field[string]                                      `json:"repo_name"`
}

func (r ProjectSourceConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type Stage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name StageName `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status StageStatus `json:"status"`
	JSON   stageJSON   `json:"-"`
}

// stageJSON contains the JSON metadata for the struct [Stage]
type stageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stageJSON) RawJSON() string {
	return r.raw
}

// The current build stage.
type StageName string

const (
	StageNameQueued     StageName = "queued"
	StageNameInitialize StageName = "initialize"
	StageNameCloneRepo  StageName = "clone_repo"
	StageNameBuild      StageName = "build"
	StageNameDeploy     StageName = "deploy"
)

func (r StageName) IsKnown() bool {
	switch r {
	case StageNameQueued, StageNameInitialize, StageNameCloneRepo, StageNameBuild, StageNameDeploy:
		return true
	}
	return false
}

// State of the current stage.
type StageStatus string

const (
	StageStatusSuccess  StageStatus = "success"
	StageStatusIdle     StageStatus = "idle"
	StageStatusActive   StageStatus = "active"
	StageStatusFailure  StageStatus = "failure"
	StageStatusCanceled StageStatus = "canceled"
)

func (r StageStatus) IsKnown() bool {
	switch r {
	case StageStatusSuccess, StageStatusIdle, StageStatusActive, StageStatusFailure, StageStatusCanceled:
		return true
	}
	return false
}

// The status of the deployment.
type StageParam struct {
	// The current build stage.
	Name param.Field[StageName] `json:"name"`
}

func (r StageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeleteResponse = interface{}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Project   ProjectParam        `json:"project,required"`
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Project)
}

type ProjectNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Project               `json:"result,required"`
	// Whether the API call was successful
	Success ProjectNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectNewResponseEnvelopeJSON    `json:"-"`
}

// projectNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelope]
type projectNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectNewResponseEnvelopeSuccess bool

const (
	ProjectNewResponseEnvelopeSuccessFalse ProjectNewResponseEnvelopeSuccess = false
	ProjectNewResponseEnvelopeSuccessTrue  ProjectNewResponseEnvelopeSuccess = true
)

func (r ProjectNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectNewResponseEnvelopeSuccessFalse, ProjectNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ProjectDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDeleteResponseEnvelope]
type projectDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeleteResponseEnvelopeSuccessFalse ProjectDeleteResponseEnvelopeSuccess = false
	ProjectDeleteResponseEnvelopeSuccessTrue  ProjectDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeleteResponseEnvelopeSuccessFalse, ProjectDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Project   ProjectParam        `json:"project,required"`
}

func (r ProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Project)
}

type ProjectEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Project               `json:"result,required"`
	// Whether the API call was successful
	Success ProjectEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectEditResponseEnvelopeJSON    `json:"-"`
}

// projectEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectEditResponseEnvelope]
type projectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectEditResponseEnvelopeSuccess bool

const (
	ProjectEditResponseEnvelopeSuccessFalse ProjectEditResponseEnvelopeSuccess = false
	ProjectEditResponseEnvelopeSuccessTrue  ProjectEditResponseEnvelopeSuccess = true
)

func (r ProjectEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectEditResponseEnvelopeSuccessFalse, ProjectEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Project               `json:"result,required"`
	// Whether the API call was successful
	Success ProjectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectGetResponseEnvelopeJSON    `json:"-"`
}

// projectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelope]
type projectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectGetResponseEnvelopeSuccess bool

const (
	ProjectGetResponseEnvelopeSuccessFalse ProjectGetResponseEnvelopeSuccess = false
	ProjectGetResponseEnvelopeSuccessTrue  ProjectGetResponseEnvelopeSuccess = true
)

func (r ProjectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectGetResponseEnvelopeSuccessFalse, ProjectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectPurgeBuildCacheParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectPurgeBuildCacheResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ProjectPurgeBuildCacheResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectPurgeBuildCacheResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectPurgeBuildCacheResponseEnvelopeJSON    `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectPurgeBuildCacheResponseEnvelope]
type projectPurgeBuildCacheResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectPurgeBuildCacheResponseEnvelopeSuccess bool

const (
	ProjectPurgeBuildCacheResponseEnvelopeSuccessFalse ProjectPurgeBuildCacheResponseEnvelopeSuccess = false
	ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue  ProjectPurgeBuildCacheResponseEnvelopeSuccess = true
)

func (r ProjectPurgeBuildCacheResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectPurgeBuildCacheResponseEnvelopeSuccessFalse, ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
