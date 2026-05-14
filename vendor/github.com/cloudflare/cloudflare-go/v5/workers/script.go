// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptService] method instead.
type ScriptService struct {
	Options                  []option.RequestOption
	Assets                   *ScriptAssetService
	Subdomain                *ScriptSubdomainService
	Schedules                *ScriptScheduleService
	Tail                     *ScriptTailService
	Content                  *ScriptContentService
	Settings                 *ScriptSettingService
	Deployments              *ScriptDeploymentService
	Versions                 *ScriptVersionService
	Secrets                  *ScriptSecretService
	ScriptAndVersionSettings *ScriptScriptAndVersionSettingService
}

// NewScriptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScriptService(opts ...option.RequestOption) (r *ScriptService) {
	r = &ScriptService{}
	r.Options = opts
	r.Assets = NewScriptAssetService(opts...)
	r.Subdomain = NewScriptSubdomainService(opts...)
	r.Schedules = NewScriptScheduleService(opts...)
	r.Tail = NewScriptTailService(opts...)
	r.Content = NewScriptContentService(opts...)
	r.Settings = NewScriptSettingService(opts...)
	r.Deployments = NewScriptDeploymentService(opts...)
	r.Versions = NewScriptVersionService(opts...)
	r.Secrets = NewScriptSecretService(opts...)
	r.ScriptAndVersionSettings = NewScriptScriptAndVersionSettingService(opts...)
	return
}

// Upload a worker module. You can find more about the multipart metadata on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *ScriptService) Update(ctx context.Context, scriptName string, params ScriptUpdateParams, opts ...option.RequestOption) (res *ScriptUpdateResponse, err error) {
	var env ScriptUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of uploaded workers.
func (r *ScriptService) List(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Script], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts", params.AccountID)
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

// Fetch a list of uploaded workers.
func (r *ScriptService) ListAutoPaging(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Script] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete your worker. This call has no response body on a successful delete.
func (r *ScriptService) Delete(ctx context.Context, scriptName string, params ScriptDeleteParams, opts ...option.RequestOption) (res *ScriptDeleteResponse, err error) {
	var env ScriptDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ScriptService) Get(ctx context.Context, scriptName string, query ScriptGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/javascript")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Script struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptPlacement `json:"placement"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementMode ScriptPlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementStatus ScriptPlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUsageModel `json:"usage_model"`
	JSON       scriptJSON       `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID              apijson.Field
	CreatedOn       apijson.Field
	Etag            apijson.Field
	HasAssets       apijson.Field
	HasModules      apijson.Field
	Logpush         apijson.Field
	ModifiedOn      apijson.Field
	Placement       apijson.Field
	PlacementMode   apijson.Field
	PlacementStatus apijson.Field
	TailConsumers   apijson.Field
	UsageModel      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptPlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status ScriptPlacementStatus `json:"status"`
	JSON   scriptPlacementJSON   `json:"-"`
}

// scriptPlacementJSON contains the JSON metadata for the struct [ScriptPlacement]
type scriptPlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScriptPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptPlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacementMode string

const (
	ScriptPlacementModeSmart ScriptPlacementMode = "smart"
)

func (r ScriptPlacementMode) IsKnown() bool {
	switch r {
	case ScriptPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacementStatus string

const (
	ScriptPlacementStatusSuccess                 ScriptPlacementStatus = "SUCCESS"
	ScriptPlacementStatusUnsupportedApplication  ScriptPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptPlacementStatusInsufficientInvocations ScriptPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptPlacementStatusSuccess, ScriptPlacementStatusUnsupportedApplication, ScriptPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUsageModel string

const (
	ScriptUsageModelStandard ScriptUsageModel = "standard"
)

func (r ScriptUsageModel) IsKnown() bool {
	switch r {
	case ScriptUsageModelStandard:
		return true
	}
	return false
}

type ScriptSetting struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability ScriptSettingObservability `json:"observability,nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript  `json:"tail_consumers,nullable"`
	JSON          scriptSettingJSON `json:"-"`
}

// scriptSettingJSON contains the JSON metadata for the struct [ScriptSetting]
type scriptSettingJSON struct {
	Logpush       apijson.Field
	Observability apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type ScriptSettingObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate,nullable"`
	// Log settings for the Worker.
	Logs ScriptSettingObservabilityLogs `json:"logs,nullable"`
	JSON scriptSettingObservabilityJSON `json:"-"`
}

// scriptSettingObservabilityJSON contains the JSON metadata for the struct
// [ScriptSettingObservability]
type scriptSettingObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type ScriptSettingObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64                            `json:"head_sampling_rate,nullable"`
	JSON             scriptSettingObservabilityLogsJSON `json:"-"`
}

// scriptSettingObservabilityLogsJSON contains the JSON metadata for the struct
// [ScriptSettingObservabilityLogs]
type scriptSettingObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingParam struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptSettingObservabilityParam] `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
}

func (r ScriptSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability settings for the Worker.
type ScriptSettingObservabilityParam struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ScriptSettingObservabilityLogsParam] `json:"logs"`
}

func (r ScriptSettingObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ScriptSettingObservabilityLogsParam struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptSettingObservabilityLogsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateResponse struct {
	StartupTimeMs int64 `json:"startup_time_ms,required"`
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptUpdateResponsePlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode ScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus ScriptUpdateResponsePlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       scriptUpdateResponseJSON       `json:"-"`
}

// scriptUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptUpdateResponse]
type scriptUpdateResponseJSON struct {
	StartupTimeMs   apijson.Field
	ID              apijson.Field
	CreatedOn       apijson.Field
	Etag            apijson.Field
	HasAssets       apijson.Field
	HasModules      apijson.Field
	Logpush         apijson.Field
	ModifiedOn      apijson.Field
	Placement       apijson.Field
	PlacementMode   apijson.Field
	PlacementStatus apijson.Field
	TailConsumers   apijson.Field
	UsageModel      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptUpdateResponsePlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status ScriptUpdateResponsePlacementStatus `json:"status"`
	JSON   scriptUpdateResponsePlacementJSON   `json:"-"`
}

// scriptUpdateResponsePlacementJSON contains the JSON metadata for the struct
// [ScriptUpdateResponsePlacement]
type scriptUpdateResponsePlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScriptUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacementMode string

const (
	ScriptUpdateResponsePlacementModeSmart ScriptUpdateResponsePlacementMode = "smart"
)

func (r ScriptUpdateResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptUpdateResponsePlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacementStatus string

const (
	ScriptUpdateResponsePlacementStatusSuccess                 ScriptUpdateResponsePlacementStatus = "SUCCESS"
	ScriptUpdateResponsePlacementStatusUnsupportedApplication  ScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateResponsePlacementStatusInsufficientInvocations ScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateResponsePlacementStatusSuccess, ScriptUpdateResponsePlacementStatusUnsupportedApplication, ScriptUpdateResponsePlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUpdateResponseUsageModel string

const (
	ScriptUpdateResponseUsageModelStandard ScriptUpdateResponseUsageModel = "standard"
)

func (r ScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseUsageModelStandard:
		return true
	}
	return false
}

type ScriptDeleteResponse = interface{}

type ScriptUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptUpdateParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r ScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// JSON-encoded metadata about the uploaded parts and Worker configuration.
type ScriptUpdateParamsMetadata struct {
	// Configuration for assets within a Worker.
	Assets param.Field[ScriptUpdateParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptUpdateParamsMetadataBindingUnion] `json:"bindings"`
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token.
	KeepAssets param.Field[bool] `json:"keep_assets"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptUpdateParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptUpdateParamsMetadataObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[ScriptUpdateParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type ScriptUpdateParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[ScriptUpdateParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r ScriptUpdateParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type ScriptUpdateParamsMetadataAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[ScriptUpdateParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r ScriptUpdateParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type ScriptUpdateParamsMetadataAssetsConfigHTMLHandling string

const (
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone               ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r ScriptUpdateParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling string

const (
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone                  ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "none"
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page               ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone, ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page, ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Satisfied by
// [workers.ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion interface {
	ImplementsScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()
}

type ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray []string

func (r ScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray) ImplementsScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}

// A binding to allow the Worker to communicate with resources.
type ScriptUpdateParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsType] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID        param.Field[string]      `json:"id"`
	Algorithm param.Field[interface{}] `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptUpdateParamsMetadataBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string]      `json:"namespace_id"`
	Outbound    param.Field[interface{}] `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id"`
	// The text value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r ScriptUpdateParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBinding) implementsScriptUpdateParamsMetadataBindingUnion() {}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow],
// [ScriptUpdateParamsMetadataBinding].
type ScriptUpdateParamsMetadataBindingUnion interface {
	implementsScriptUpdateParamsMetadataBindingUnion()
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1 ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign       ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsType string

const (
	ScriptUpdateParamsMetadataBindingsTypeAI                     ScriptUpdateParamsMetadataBindingsType = "ai"
	ScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsTypeAssets                 ScriptUpdateParamsMetadataBindingsType = "assets"
	ScriptUpdateParamsMetadataBindingsTypeBrowser                ScriptUpdateParamsMetadataBindingsType = "browser"
	ScriptUpdateParamsMetadataBindingsTypeD1                     ScriptUpdateParamsMetadataBindingsType = "d1"
	ScriptUpdateParamsMetadataBindingsTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsTypeHyperdrive             ScriptUpdateParamsMetadataBindingsType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsTypeJson                   ScriptUpdateParamsMetadataBindingsType = "json"
	ScriptUpdateParamsMetadataBindingsTypeKVNamespace            ScriptUpdateParamsMetadataBindingsType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsTypePlainText              ScriptUpdateParamsMetadataBindingsType = "plain_text"
	ScriptUpdateParamsMetadataBindingsTypePipelines              ScriptUpdateParamsMetadataBindingsType = "pipelines"
	ScriptUpdateParamsMetadataBindingsTypeQueue                  ScriptUpdateParamsMetadataBindingsType = "queue"
	ScriptUpdateParamsMetadataBindingsTypeR2Bucket               ScriptUpdateParamsMetadataBindingsType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsTypeSecretText             ScriptUpdateParamsMetadataBindingsType = "secret_text"
	ScriptUpdateParamsMetadataBindingsTypeService                ScriptUpdateParamsMetadataBindingsType = "service"
	ScriptUpdateParamsMetadataBindingsTypeTailConsumer           ScriptUpdateParamsMetadataBindingsType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsTypeVectorize              ScriptUpdateParamsMetadataBindingsType = "vectorize"
	ScriptUpdateParamsMetadataBindingsTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsType = "version_metadata"
	ScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret     ScriptUpdateParamsMetadataBindingsType = "secrets_store_secret"
	ScriptUpdateParamsMetadataBindingsTypeSecretKey              ScriptUpdateParamsMetadataBindingsType = "secret_key"
	ScriptUpdateParamsMetadataBindingsTypeWorkflow               ScriptUpdateParamsMetadataBindingsType = "workflow"
)

func (r ScriptUpdateParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsTypeAI, ScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsTypeAssets, ScriptUpdateParamsMetadataBindingsTypeBrowser, ScriptUpdateParamsMetadataBindingsTypeD1, ScriptUpdateParamsMetadataBindingsTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsTypeHyperdrive, ScriptUpdateParamsMetadataBindingsTypeJson, ScriptUpdateParamsMetadataBindingsTypeKVNamespace, ScriptUpdateParamsMetadataBindingsTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsTypePlainText, ScriptUpdateParamsMetadataBindingsTypePipelines, ScriptUpdateParamsMetadataBindingsTypeQueue, ScriptUpdateParamsMetadataBindingsTypeR2Bucket, ScriptUpdateParamsMetadataBindingsTypeSecretText, ScriptUpdateParamsMetadataBindingsTypeService, ScriptUpdateParamsMetadataBindingsTypeTailConsumer, ScriptUpdateParamsMetadataBindingsTypeVectorize, ScriptUpdateParamsMetadataBindingsTypeVersionMetadata, ScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret, ScriptUpdateParamsMetadataBindingsTypeSecretKey, ScriptUpdateParamsMetadataBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptUpdateParamsMetadataBindingsFormat string

const (
	ScriptUpdateParamsMetadataBindingsFormatRaw   ScriptUpdateParamsMetadataBindingsFormat = "raw"
	ScriptUpdateParamsMetadataBindingsFormatPkcs8 ScriptUpdateParamsMetadataBindingsFormat = "pkcs8"
	ScriptUpdateParamsMetadataBindingsFormatSpki  ScriptUpdateParamsMetadataBindingsFormat = "spki"
	ScriptUpdateParamsMetadataBindingsFormatJwk   ScriptUpdateParamsMetadataBindingsFormat = "jwk"
)

func (r ScriptUpdateParamsMetadataBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsFormatRaw, ScriptUpdateParamsMetadataBindingsFormatPkcs8, ScriptUpdateParamsMetadataBindingsFormatSpki, ScriptUpdateParamsMetadataBindingsFormatJwk:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptUpdateParamsMetadataMigrations struct {
	DeletedClasses   param.Field[interface{}] `json:"deleted_classes"`
	NewClasses       param.Field[interface{}] `json:"new_classes"`
	NewSqliteClasses param.Field[interface{}] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes"`
	Steps              param.Field[interface{}] `json:"steps"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes"`
}

func (r ScriptUpdateParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataMigrations) implementsScriptUpdateParamsMetadataMigrationsUnion() {}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [ScriptUpdateParamsMetadataMigrations].
type ScriptUpdateParamsMetadataMigrationsUnion interface {
	implementsScriptUpdateParamsMetadataMigrationsUnion()
}

type ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) implementsScriptUpdateParamsMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type ScriptUpdateParamsMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ScriptUpdateParamsMetadataObservabilityLogs] `json:"logs"`
}

func (r ScriptUpdateParamsMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ScriptUpdateParamsMetadataObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptUpdateParamsMetadataObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r ScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacementMode string

const (
	ScriptUpdateParamsMetadataPlacementModeSmart ScriptUpdateParamsMetadataPlacementMode = "smart"
)

func (r ScriptUpdateParamsMetadataPlacementMode) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacementStatus string

const (
	ScriptUpdateParamsMetadataPlacementStatusSuccess                 ScriptUpdateParamsMetadataPlacementStatus = "SUCCESS"
	ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  ScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations ScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataPlacementStatusSuccess, ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUpdateParamsMetadataUsageModel string

const (
	ScriptUpdateParamsMetadataUsageModelStandard ScriptUpdateParamsMetadataUsageModel = "standard"
)

func (r ScriptUpdateParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type ScriptUpdateResponseEnvelope struct {
	Errors   []ScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelope]
type scriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ScriptUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelopeErrors]
type scriptUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    scriptUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptUpdateResponseEnvelopeErrorsSource]
type scriptUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptUpdateResponseEnvelopeMessages]
type scriptUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptUpdateResponseEnvelopeMessagesSource]
type scriptUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptUpdateResponseEnvelopeSuccess bool

const (
	ScriptUpdateResponseEnvelopeSuccessTrue ScriptUpdateResponseEnvelopeSuccess = true
)

func (r ScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where
	// allowed is 'yes' or 'no'.
	Tags param.Field[string] `query:"tags"`
}

// URLQuery serializes [ScriptListParams]'s query parameters as `url.Values`.
func (r ScriptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [ScriptDeleteParams]'s query parameters as `url.Values`.
func (r ScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeleteResponseEnvelope struct {
	Errors   []ScriptDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptDeleteResponse                `json:"result,nullable"`
	JSON    scriptDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptDeleteResponseEnvelope]
type scriptDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ScriptDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScriptDeleteResponseEnvelopeErrors]
type scriptDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    scriptDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptDeleteResponseEnvelopeErrorsSource]
type scriptDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptDeleteResponseEnvelopeMessages]
type scriptDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptDeleteResponseEnvelopeMessagesSource]
type scriptDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeleteResponseEnvelopeSuccess bool

const (
	ScriptDeleteResponseEnvelopeSuccessTrue ScriptDeleteResponseEnvelopeSuccess = true
)

func (r ScriptDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
