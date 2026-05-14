// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"github.com/cloudflare/cloudflare-go/v5/workers"
)

// DispatchNamespaceScriptService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptService] method instead.
type DispatchNamespaceScriptService struct {
	Options     []option.RequestOption
	AssetUpload *DispatchNamespaceScriptAssetUploadService
	Content     *DispatchNamespaceScriptContentService
	Settings    *DispatchNamespaceScriptSettingService
	Bindings    *DispatchNamespaceScriptBindingService
	Secrets     *DispatchNamespaceScriptSecretService
	Tags        *DispatchNamespaceScriptTagService
}

// NewDispatchNamespaceScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptService(opts ...option.RequestOption) (r *DispatchNamespaceScriptService) {
	r = &DispatchNamespaceScriptService{}
	r.Options = opts
	r.AssetUpload = NewDispatchNamespaceScriptAssetUploadService(opts...)
	r.Content = NewDispatchNamespaceScriptContentService(opts...)
	r.Settings = NewDispatchNamespaceScriptSettingService(opts...)
	r.Bindings = NewDispatchNamespaceScriptBindingService(opts...)
	r.Secrets = NewDispatchNamespaceScriptSecretService(opts...)
	r.Tags = NewDispatchNamespaceScriptTagService(opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace. You can find more
// about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *DispatchNamespaceScriptService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptUpdateParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptUpdateResponse, err error) {
	var env DispatchNamespaceScriptUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *DispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptDeleteResponse, err error) {
	var env DispatchNamespaceScriptDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *Script, err error) {
	var env DispatchNamespaceScriptGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type Script struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time      `json:"modified_on" format:"date-time"`
	Script     workers.Script `json:"script"`
	JSON       scriptJSON     `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponse struct {
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
	Placement DispatchNamespaceScriptUpdateResponsePlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptUpdateResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptUpdateResponse]
type dispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *DispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"status"`
	JSON   dispatchNamespaceScriptUpdateResponsePlacementJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponsePlacementJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptUpdateResponsePlacement]
type dispatchNamespaceScriptUpdateResponsePlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementMode string

const (
	DispatchNamespaceScriptUpdateResponsePlacementModeSmart DispatchNamespaceScriptUpdateResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementStatus string

const (
	DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess                 DispatchNamespaceScriptUpdateResponsePlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess, DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateResponseUsageModel string

const (
	DispatchNamespaceScriptUpdateResponseUsageModelStandard DispatchNamespaceScriptUpdateResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptDeleteResponse = interface{}

type DispatchNamespaceScriptUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[DispatchNamespaceScriptUpdateParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r DispatchNamespaceScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type DispatchNamespaceScriptUpdateParamsMetadata struct {
	// Configuration for assets within a Worker.
	Assets param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingUnion] `json:"bindings"`
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
	Migrations param.Field[DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone                  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "none"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
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
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray []string

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray) ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptUpdateParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsType] `json:"type,required"`
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
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat] `json:"format"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow],
// [DispatchNamespaceScriptUpdateParamsMetadataBinding].
type DispatchNamespaceScriptUpdateParamsMetadataBindingUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowser                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "browser"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePipelines              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "pipelines"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "version_metadata"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretKey              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secret_key"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWorkflow               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "workflow"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowser, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePipelines, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatJwk:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptUpdateParamsMetadataMigrations struct {
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

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptUpdateParamsMetadataMigrations].
type DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs] `json:"logs"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementMode string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart DispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "smart"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess                 DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateParamsMetadataUsageModel string

const (
	DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard DispatchNamespaceScriptUpdateParamsMetadataUsageModel = "standard"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseEnvelope]
type dispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [DispatchNamespaceScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r DispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DispatchNamespaceScriptDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptDeleteResponseEnvelope]
type dispatchNamespaceScriptDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeErrors]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeMessages]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue DispatchNamespaceScriptDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result Script `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptGetResponseEnvelope]
type dispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrors]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeMessages]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptGetResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
