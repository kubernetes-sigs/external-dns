package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PagesPreviewDeploymentSetting string

const (
	PagesPreviewAllBranches    PagesPreviewDeploymentSetting = "all"
	PagesPreviewNoBranches     PagesPreviewDeploymentSetting = "none"
	PagesPreviewCustomBranches PagesPreviewDeploymentSetting = "custom"
)

// PagesProject represents a Pages project.
type PagesProject struct {
	Name                string                        `json:"name,omitempty"`
	ID                  string                        `json:"id"`
	CreatedOn           *time.Time                    `json:"created_on"`
	SubDomain           string                        `json:"subdomain"`
	Domains             []string                      `json:"domains,omitempty"`
	Source              *PagesProjectSource           `json:"source,omitempty"`
	BuildConfig         PagesProjectBuildConfig       `json:"build_config"`
	DeploymentConfigs   PagesProjectDeploymentConfigs `json:"deployment_configs"`
	LatestDeployment    PagesProjectDeployment        `json:"latest_deployment"`
	CanonicalDeployment PagesProjectDeployment        `json:"canonical_deployment"`
	ProductionBranch    string                        `json:"production_branch,omitempty"`
}

// PagesProjectSource represents the configuration of a Pages project source.
type PagesProjectSource struct {
	Type   string                    `json:"type"`
	Config *PagesProjectSourceConfig `json:"config"`
}

// PagesProjectSourceConfig represents the properties use to configure a Pages project source.
type PagesProjectSourceConfig struct {
	Owner                        string                        `json:"owner"`
	RepoName                     string                        `json:"repo_name"`
	ProductionBranch             string                        `json:"production_branch"`
	PRCommentsEnabled            bool                          `json:"pr_comments_enabled"`
	DeploymentsEnabled           bool                          `json:"deployments_enabled"`
	ProductionDeploymentsEnabled bool                          `json:"production_deployments_enabled"`
	PreviewDeploymentSetting     PagesPreviewDeploymentSetting `json:"preview_deployment_setting"`
	PreviewBranchIncludes        []string                      `json:"preview_branch_includes"`
	PreviewBranchExcludes        []string                      `json:"preview_branch_excludes"`
}

// PagesProjectBuildConfig represents the configuration of a Pages project build process.
type PagesProjectBuildConfig struct {
	BuildCommand      string `json:"build_command"`
	DestinationDir    string `json:"destination_dir"`
	RootDir           string `json:"root_dir"`
	WebAnalyticsTag   string `json:"web_analytics_tag"`
	WebAnalyticsToken string `json:"web_analytics_token"`
}

// PagesProjectDeploymentConfigs represents the configuration for deployments in a Pages project.
type PagesProjectDeploymentConfigs struct {
	Preview    PagesProjectDeploymentConfigEnvironment `json:"preview"`
	Production PagesProjectDeploymentConfigEnvironment `json:"production"`
}

// PagesProjectDeploymentConfigEnvironment represents the configuration for preview or production deploys.
type PagesProjectDeploymentConfigEnvironment struct {
	EnvVars            map[string]PagesProjectDeploymentVar `json:"env_vars"`
	CompatibilityDate  string                               `json:"compatibility_date,omitempty"`
	CompatibilityFlags []string                             `json:"compatibility_flags,omitempty"`
	KvNamespaces       NamespaceBindingMap                  `json:"kv_namespaces,omitempty"`
	DoNamespaces       NamespaceBindingMap                  `json:"durable_object_namespaces,omitempty"`
	D1Databases        D1BindingMap                         `json:"d1_databases,omitempty"`
	R2Bindings         R2BindingMap                         `json:"r2_buckets,omitempty"`
}

// PagesProjectDeploymentVar represents a deployment environment variable.
type PagesProjectDeploymentVar struct {
	Value string `json:"value"`
}

// PagesProjectDeployment represents a deployment to a Pages project.
type PagesProjectDeployment struct {
	ID                 string                        `json:"id"`
	ShortID            string                        `json:"short_id"`
	ProjectID          string                        `json:"project_id"`
	ProjectName        string                        `json:"project_name"`
	Environment        string                        `json:"environment"`
	URL                string                        `json:"url"`
	CreatedOn          *time.Time                    `json:"created_on"`
	ModifiedOn         *time.Time                    `json:"modified_on"`
	Aliases            []string                      `json:"aliases,omitempty"`
	LatestStage        PagesProjectDeploymentStage   `json:"latest_stage"`
	EnvVars            map[string]map[string]string  `json:"env_vars"`
	DeploymentTrigger  PagesProjectDeploymentTrigger `json:"deployment_trigger"`
	Stages             []PagesProjectDeploymentStage `json:"stages"`
	BuildConfig        PagesProjectBuildConfig       `json:"build_config"`
	Source             PagesProjectSource            `json:"source"`
	CompatibilityDate  string                        `json:"compatibility_date,omitempty"`
	CompatibilityFlags []string                      `json:"compatibility_flags,omitempty"`
	ProductionBranch   string                        `json:"production_branch,omitempty"`
}

// PagesProjectDeploymentStage represents an individual stage in a Pages project deployment.
type PagesProjectDeploymentStage struct {
	Name      string     `json:"name"`
	StartedOn *time.Time `json:"started_on,omitempty"`
	EndedOn   *time.Time `json:"ended_on,omitempty"`
	Status    string     `json:"status"`
}

// PagesProjectDeploymentTrigger represents information about what caused a deployment.
type PagesProjectDeploymentTrigger struct {
	Type     string                                 `json:"type"`
	Metadata *PagesProjectDeploymentTriggerMetadata `json:"metadata"`
}

// PagesProjectDeploymentTriggerMetadata represents additional information about the cause of a deployment.
type PagesProjectDeploymentTriggerMetadata struct {
	Branch        string `json:"branch"`
	CommitHash    string `json:"commit_hash"`
	CommitMessage string `json:"commit_message"`
}

type pagesProjectResponse struct {
	Response
	Result PagesProject `json:"result"`
}

type pagesProjectListResponse struct {
	Response
	Result     []PagesProject `json:"result"`
	ResultInfo `json:"result_info"`
}

type NamespaceBindingMap map[string]*NamespaceBindingValue

type NamespaceBindingValue struct {
	Value string `json:"namespace_id"`
}

type R2BindingMap map[string]*R2BindingValue

type R2BindingValue struct {
	Name string `json:"name"`
}

type D1BindingMap map[string]*D1Binding

type D1Binding struct {
	ID string `json:"id"`
}

// ListPagesProjects returns all Pages projects for an account.
//
// API reference: https://api.cloudflare.com/#pages-project-get-projects
func (api *API) ListPagesProjects(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]PagesProject, ResultInfo, error) {
	uri := buildURI(fmt.Sprintf("/accounts/%s/pages/projects", accountID), pageOpts)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PagesProject{}, ResultInfo{}, err
	}
	var r pagesProjectListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []PagesProject{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, nil
}

// PagesProject returns a single Pages project by name.
//
// API reference: https://api.cloudflare.com/#pages-project-get-project
func (api *API) PagesProject(ctx context.Context, accountID, projectName string) (PagesProject, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s", accountID, projectName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesProject{}, err
	}
	var r pagesProjectResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProject{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreatePagesProject creates a new Pages project in an account.
//
// API reference: https://api.cloudflare.com/#pages-project-create-project
func (api *API) CreatePagesProject(ctx context.Context, accountID string, pagesProject PagesProject) (PagesProject, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, pagesProject)
	if err != nil {
		return PagesProject{}, err
	}
	var r pagesProjectResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProject{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdatePagesProject updates an existing Pages project.
//
// API reference: https://api.cloudflare.com/#pages-project-update-project
func (api *API) UpdatePagesProject(ctx context.Context, accountID, projectName string, pagesProject PagesProject) (PagesProject, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s", accountID, projectName)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, pagesProject)
	if err != nil {
		return PagesProject{}, err
	}
	var r pagesProjectResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProject{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeletePagesProject deletes a Pages project by name.
//
// API reference: https://api.cloudflare.com/#pages-project-delete-project
func (api *API) DeletePagesProject(ctx context.Context, accountID, projectName string) error {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s", accountID, projectName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r pagesProjectResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}
