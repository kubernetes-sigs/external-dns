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
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// ScriptVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptVersionService] method instead.
type ScriptVersionService struct {
	Options []option.RequestOption
}

// NewScriptVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptVersionService(opts ...option.RequestOption) (r *ScriptVersionService) {
	r = &ScriptVersionService{}
	r.Options = opts
	return
}

// Upload a Worker Version without deploying to Cloudflare's network. You can find
// more about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *ScriptVersionService) New(ctx context.Context, scriptName string, params ScriptVersionNewParams, opts ...option.RequestOption) (res *ScriptVersionNewResponse, err error) {
	var env ScriptVersionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) List(ctx context.Context, scriptName string, params ScriptVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ScriptVersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
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

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) ListAutoPaging(ctx context.Context, scriptName string, params ScriptVersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ScriptVersionListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, scriptName, params, opts...))
}

// Get Version Detail
func (r *ScriptVersionService) Get(ctx context.Context, scriptName string, versionID string, query ScriptVersionGetParams, opts ...option.RequestOption) (res *ScriptVersionGetResponse, err error) {
	var env ScriptVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", query.AccountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptVersionNewResponse struct {
	Resources     ScriptVersionNewResponseResources `json:"resources,required"`
	ID            string                            `json:"id"`
	Metadata      ScriptVersionNewResponseMetadata  `json:"metadata"`
	Number        float64                           `json:"number"`
	StartupTimeMs int64                             `json:"startup_time_ms"`
	JSON          scriptVersionNewResponseJSON      `json:"-"`
}

// scriptVersionNewResponseJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponse]
type scriptVersionNewResponseJSON struct {
	Resources     apijson.Field
	ID            apijson.Field
	Metadata      apijson.Field
	Number        apijson.Field
	StartupTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResources struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings      []ScriptVersionNewResponseResourcesBinding     `json:"bindings"`
	Script        ScriptVersionNewResponseResourcesScript        `json:"script"`
	ScriptRuntime ScriptVersionNewResponseResourcesScriptRuntime `json:"script_runtime"`
	JSON          scriptVersionNewResponseResourcesJSON          `json:"-"`
}

// scriptVersionNewResponseResourcesJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseResources]
type scriptVersionNewResponseResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewResponseResourcesBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// This field can have the runtime type of
	// [[]ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                       `json:"workflow_name"`
	JSON         scriptVersionNewResponseResourcesBindingJSON `json:"-"`
	union        ScriptVersionNewResponseResourcesBindingsUnion
}

// scriptVersionNewResponseResourcesBindingJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesBinding]
type scriptVersionNewResponseResourcesBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Algorithm     apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	Format        apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	KeyJwk        apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	Pipeline      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	SecretName    apijson.Field
	Service       apijson.Field
	StoreID       apijson.Field
	Text          apijson.Field
	Usages        apijson.Field
	WorkflowName  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionNewResponseResourcesBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionNewResponseResourcesBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionNewResponseResourcesBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionNewResponseResourcesBindingsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow].
func (r ScriptVersionNewResponseResourcesBinding) AsUnion() ScriptVersionNewResponseResourcesBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey] or
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow].
type ScriptVersionNewResponseResourcesBindingsUnion interface {
	implementsScriptVersionNewResponseResourcesBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionNewResponseResourcesBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAITypeAI ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1TypeD1 ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonTypeJson ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueTypeQueue ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceTypeService ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumer) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey) implementsScriptVersionNewResponseResourcesBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                  `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsType string

const (
	ScriptVersionNewResponseResourcesBindingsTypeAI                     ScriptVersionNewResponseResourcesBindingsType = "ai"
	ScriptVersionNewResponseResourcesBindingsTypeAnalyticsEngine        ScriptVersionNewResponseResourcesBindingsType = "analytics_engine"
	ScriptVersionNewResponseResourcesBindingsTypeAssets                 ScriptVersionNewResponseResourcesBindingsType = "assets"
	ScriptVersionNewResponseResourcesBindingsTypeBrowser                ScriptVersionNewResponseResourcesBindingsType = "browser"
	ScriptVersionNewResponseResourcesBindingsTypeD1                     ScriptVersionNewResponseResourcesBindingsType = "d1"
	ScriptVersionNewResponseResourcesBindingsTypeDispatchNamespace      ScriptVersionNewResponseResourcesBindingsType = "dispatch_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsType = "durable_object_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeHyperdrive             ScriptVersionNewResponseResourcesBindingsType = "hyperdrive"
	ScriptVersionNewResponseResourcesBindingsTypeJson                   ScriptVersionNewResponseResourcesBindingsType = "json"
	ScriptVersionNewResponseResourcesBindingsTypeKVNamespace            ScriptVersionNewResponseResourcesBindingsType = "kv_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeMTLSCertificate        ScriptVersionNewResponseResourcesBindingsType = "mtls_certificate"
	ScriptVersionNewResponseResourcesBindingsTypePlainText              ScriptVersionNewResponseResourcesBindingsType = "plain_text"
	ScriptVersionNewResponseResourcesBindingsTypePipelines              ScriptVersionNewResponseResourcesBindingsType = "pipelines"
	ScriptVersionNewResponseResourcesBindingsTypeQueue                  ScriptVersionNewResponseResourcesBindingsType = "queue"
	ScriptVersionNewResponseResourcesBindingsTypeR2Bucket               ScriptVersionNewResponseResourcesBindingsType = "r2_bucket"
	ScriptVersionNewResponseResourcesBindingsTypeSecretText             ScriptVersionNewResponseResourcesBindingsType = "secret_text"
	ScriptVersionNewResponseResourcesBindingsTypeService                ScriptVersionNewResponseResourcesBindingsType = "service"
	ScriptVersionNewResponseResourcesBindingsTypeTailConsumer           ScriptVersionNewResponseResourcesBindingsType = "tail_consumer"
	ScriptVersionNewResponseResourcesBindingsTypeVectorize              ScriptVersionNewResponseResourcesBindingsType = "vectorize"
	ScriptVersionNewResponseResourcesBindingsTypeVersionMetadata        ScriptVersionNewResponseResourcesBindingsType = "version_metadata"
	ScriptVersionNewResponseResourcesBindingsTypeSecretsStoreSecret     ScriptVersionNewResponseResourcesBindingsType = "secrets_store_secret"
	ScriptVersionNewResponseResourcesBindingsTypeSecretKey              ScriptVersionNewResponseResourcesBindingsType = "secret_key"
	ScriptVersionNewResponseResourcesBindingsTypeWorkflow               ScriptVersionNewResponseResourcesBindingsType = "workflow"
)

func (r ScriptVersionNewResponseResourcesBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsTypeAI, ScriptVersionNewResponseResourcesBindingsTypeAnalyticsEngine, ScriptVersionNewResponseResourcesBindingsTypeAssets, ScriptVersionNewResponseResourcesBindingsTypeBrowser, ScriptVersionNewResponseResourcesBindingsTypeD1, ScriptVersionNewResponseResourcesBindingsTypeDispatchNamespace, ScriptVersionNewResponseResourcesBindingsTypeDurableObjectNamespace, ScriptVersionNewResponseResourcesBindingsTypeHyperdrive, ScriptVersionNewResponseResourcesBindingsTypeJson, ScriptVersionNewResponseResourcesBindingsTypeKVNamespace, ScriptVersionNewResponseResourcesBindingsTypeMTLSCertificate, ScriptVersionNewResponseResourcesBindingsTypePlainText, ScriptVersionNewResponseResourcesBindingsTypePipelines, ScriptVersionNewResponseResourcesBindingsTypeQueue, ScriptVersionNewResponseResourcesBindingsTypeR2Bucket, ScriptVersionNewResponseResourcesBindingsTypeSecretText, ScriptVersionNewResponseResourcesBindingsTypeService, ScriptVersionNewResponseResourcesBindingsTypeTailConsumer, ScriptVersionNewResponseResourcesBindingsTypeVectorize, ScriptVersionNewResponseResourcesBindingsTypeVersionMetadata, ScriptVersionNewResponseResourcesBindingsTypeSecretsStoreSecret, ScriptVersionNewResponseResourcesBindingsTypeSecretKey, ScriptVersionNewResponseResourcesBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsFormat string

const (
	ScriptVersionNewResponseResourcesBindingsFormatRaw   ScriptVersionNewResponseResourcesBindingsFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsFormatPkcs8 ScriptVersionNewResponseResourcesBindingsFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsFormatSpki  ScriptVersionNewResponseResourcesBindingsFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsFormatJwk   ScriptVersionNewResponseResourcesBindingsFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsFormatRaw, ScriptVersionNewResponseResourcesBindingsFormatPkcs8, ScriptVersionNewResponseResourcesBindingsFormatSpki, ScriptVersionNewResponseResourcesBindingsFormatJwk:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesScript struct {
	Etag             string                                                `json:"etag"`
	Handlers         []string                                              `json:"handlers"`
	LastDeployedFrom string                                                `json:"last_deployed_from"`
	NamedHandlers    []ScriptVersionNewResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             scriptVersionNewResponseResourcesScriptJSON           `json:"-"`
}

// scriptVersionNewResponseResourcesScriptJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesScript]
type scriptVersionNewResponseResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptNamedHandler struct {
	Handlers []string                                                `json:"handlers"`
	Name     string                                                  `json:"name"`
	JSON     scriptVersionNewResponseResourcesScriptNamedHandlerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptNamedHandlerJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseResourcesScriptNamedHandler]
type scriptVersionNewResponseResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntime struct {
	CompatibilityDate  string                                                   `json:"compatibility_date"`
	CompatibilityFlags []string                                                 `json:"compatibility_flags"`
	Limits             ScriptVersionNewResponseResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                                   `json:"migration_tag"`
	UsageModel         ScriptVersionNewResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               scriptVersionNewResponseResourcesScriptRuntimeJSON       `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseResourcesScriptRuntime]
type scriptVersionNewResponseResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntimeLimits struct {
	CPUMs int64                                                    `json:"cpu_ms"`
	JSON  scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseResourcesScriptRuntimeLimits]
type scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntimeUsageModel string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelBundled  ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "bundled"
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelUnbound  ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "unbound"
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelStandard ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "standard"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeUsageModelBundled, ScriptVersionNewResponseResourcesScriptRuntimeUsageModelUnbound, ScriptVersionNewResponseResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionNewResponseMetadata struct {
	AuthorEmail string                                 `json:"author_email"`
	AuthorID    string                                 `json:"author_id"`
	CreatedOn   string                                 `json:"created_on"`
	HasPreview  bool                                   `json:"hasPreview"`
	ModifiedOn  string                                 `json:"modified_on"`
	Source      ScriptVersionNewResponseMetadataSource `json:"source"`
	JSON        scriptVersionNewResponseMetadataJSON   `json:"-"`
}

// scriptVersionNewResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseMetadata]
type scriptVersionNewResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseMetadataSource string

const (
	ScriptVersionNewResponseMetadataSourceUnknown      ScriptVersionNewResponseMetadataSource = "unknown"
	ScriptVersionNewResponseMetadataSourceAPI          ScriptVersionNewResponseMetadataSource = "api"
	ScriptVersionNewResponseMetadataSourceWrangler     ScriptVersionNewResponseMetadataSource = "wrangler"
	ScriptVersionNewResponseMetadataSourceTerraform    ScriptVersionNewResponseMetadataSource = "terraform"
	ScriptVersionNewResponseMetadataSourceDash         ScriptVersionNewResponseMetadataSource = "dash"
	ScriptVersionNewResponseMetadataSourceDashTemplate ScriptVersionNewResponseMetadataSource = "dash_template"
	ScriptVersionNewResponseMetadataSourceIntegration  ScriptVersionNewResponseMetadataSource = "integration"
	ScriptVersionNewResponseMetadataSourceQuickEditor  ScriptVersionNewResponseMetadataSource = "quick_editor"
	ScriptVersionNewResponseMetadataSourcePlayground   ScriptVersionNewResponseMetadataSource = "playground"
	ScriptVersionNewResponseMetadataSourceWorkersci    ScriptVersionNewResponseMetadataSource = "workersci"
)

func (r ScriptVersionNewResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseMetadataSourceUnknown, ScriptVersionNewResponseMetadataSourceAPI, ScriptVersionNewResponseMetadataSourceWrangler, ScriptVersionNewResponseMetadataSourceTerraform, ScriptVersionNewResponseMetadataSourceDash, ScriptVersionNewResponseMetadataSourceDashTemplate, ScriptVersionNewResponseMetadataSourceIntegration, ScriptVersionNewResponseMetadataSourceQuickEditor, ScriptVersionNewResponseMetadataSourcePlayground, ScriptVersionNewResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionListResponse struct {
	ID       string                            `json:"id"`
	Metadata ScriptVersionListResponseMetadata `json:"metadata"`
	Number   float64                           `json:"number"`
	JSON     scriptVersionListResponseJSON     `json:"-"`
}

// scriptVersionListResponseJSON contains the JSON metadata for the struct
// [ScriptVersionListResponse]
type scriptVersionListResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponseMetadata struct {
	AuthorEmail string                                  `json:"author_email"`
	AuthorID    string                                  `json:"author_id"`
	CreatedOn   string                                  `json:"created_on"`
	HasPreview  bool                                    `json:"hasPreview"`
	ModifiedOn  string                                  `json:"modified_on"`
	Source      ScriptVersionListResponseMetadataSource `json:"source"`
	JSON        scriptVersionListResponseMetadataJSON   `json:"-"`
}

// scriptVersionListResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionListResponseMetadata]
type scriptVersionListResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponseMetadataSource string

const (
	ScriptVersionListResponseMetadataSourceUnknown      ScriptVersionListResponseMetadataSource = "unknown"
	ScriptVersionListResponseMetadataSourceAPI          ScriptVersionListResponseMetadataSource = "api"
	ScriptVersionListResponseMetadataSourceWrangler     ScriptVersionListResponseMetadataSource = "wrangler"
	ScriptVersionListResponseMetadataSourceTerraform    ScriptVersionListResponseMetadataSource = "terraform"
	ScriptVersionListResponseMetadataSourceDash         ScriptVersionListResponseMetadataSource = "dash"
	ScriptVersionListResponseMetadataSourceDashTemplate ScriptVersionListResponseMetadataSource = "dash_template"
	ScriptVersionListResponseMetadataSourceIntegration  ScriptVersionListResponseMetadataSource = "integration"
	ScriptVersionListResponseMetadataSourceQuickEditor  ScriptVersionListResponseMetadataSource = "quick_editor"
	ScriptVersionListResponseMetadataSourcePlayground   ScriptVersionListResponseMetadataSource = "playground"
	ScriptVersionListResponseMetadataSourceWorkersci    ScriptVersionListResponseMetadataSource = "workersci"
)

func (r ScriptVersionListResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionListResponseMetadataSourceUnknown, ScriptVersionListResponseMetadataSourceAPI, ScriptVersionListResponseMetadataSourceWrangler, ScriptVersionListResponseMetadataSourceTerraform, ScriptVersionListResponseMetadataSourceDash, ScriptVersionListResponseMetadataSourceDashTemplate, ScriptVersionListResponseMetadataSourceIntegration, ScriptVersionListResponseMetadataSourceQuickEditor, ScriptVersionListResponseMetadataSourcePlayground, ScriptVersionListResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionGetResponse struct {
	Resources ScriptVersionGetResponseResources `json:"resources,required"`
	ID        string                            `json:"id"`
	Metadata  ScriptVersionGetResponseMetadata  `json:"metadata"`
	Number    float64                           `json:"number"`
	JSON      scriptVersionGetResponseJSON      `json:"-"`
}

// scriptVersionGetResponseJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponse]
type scriptVersionGetResponseJSON struct {
	Resources   apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResources struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings      []ScriptVersionGetResponseResourcesBinding     `json:"bindings"`
	Script        ScriptVersionGetResponseResourcesScript        `json:"script"`
	ScriptRuntime ScriptVersionGetResponseResourcesScriptRuntime `json:"script_runtime"`
	JSON          scriptVersionGetResponseResourcesJSON          `json:"-"`
}

// scriptVersionGetResponseResourcesJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseResources]
type scriptVersionGetResponseResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionGetResponseResourcesBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// This field can have the runtime type of
	// [[]ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                       `json:"workflow_name"`
	JSON         scriptVersionGetResponseResourcesBindingJSON `json:"-"`
	union        ScriptVersionGetResponseResourcesBindingsUnion
}

// scriptVersionGetResponseResourcesBindingJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesBinding]
type scriptVersionGetResponseResourcesBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Algorithm     apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	Format        apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	KeyJwk        apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	Pipeline      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	SecretName    apijson.Field
	Service       apijson.Field
	StoreID       apijson.Field
	Text          apijson.Field
	Usages        apijson.Field
	WorkflowName  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionGetResponseResourcesBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionGetResponseResourcesBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionGetResponseResourcesBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionGetResponseResourcesBindingsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow].
func (r ScriptVersionGetResponseResourcesBinding) AsUnion() ScriptVersionGetResponseResourcesBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey] or
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow].
type ScriptVersionGetResponseResourcesBindingsUnion interface {
	implementsScriptVersionGetResponseResourcesBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionGetResponseResourcesBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAITypeAI ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1TypeD1 ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonTypeJson ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueTypeQueue ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceTypeService ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumer) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey) implementsScriptVersionGetResponseResourcesBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                  `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsType string

const (
	ScriptVersionGetResponseResourcesBindingsTypeAI                     ScriptVersionGetResponseResourcesBindingsType = "ai"
	ScriptVersionGetResponseResourcesBindingsTypeAnalyticsEngine        ScriptVersionGetResponseResourcesBindingsType = "analytics_engine"
	ScriptVersionGetResponseResourcesBindingsTypeAssets                 ScriptVersionGetResponseResourcesBindingsType = "assets"
	ScriptVersionGetResponseResourcesBindingsTypeBrowser                ScriptVersionGetResponseResourcesBindingsType = "browser"
	ScriptVersionGetResponseResourcesBindingsTypeD1                     ScriptVersionGetResponseResourcesBindingsType = "d1"
	ScriptVersionGetResponseResourcesBindingsTypeDispatchNamespace      ScriptVersionGetResponseResourcesBindingsType = "dispatch_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsType = "durable_object_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeHyperdrive             ScriptVersionGetResponseResourcesBindingsType = "hyperdrive"
	ScriptVersionGetResponseResourcesBindingsTypeJson                   ScriptVersionGetResponseResourcesBindingsType = "json"
	ScriptVersionGetResponseResourcesBindingsTypeKVNamespace            ScriptVersionGetResponseResourcesBindingsType = "kv_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeMTLSCertificate        ScriptVersionGetResponseResourcesBindingsType = "mtls_certificate"
	ScriptVersionGetResponseResourcesBindingsTypePlainText              ScriptVersionGetResponseResourcesBindingsType = "plain_text"
	ScriptVersionGetResponseResourcesBindingsTypePipelines              ScriptVersionGetResponseResourcesBindingsType = "pipelines"
	ScriptVersionGetResponseResourcesBindingsTypeQueue                  ScriptVersionGetResponseResourcesBindingsType = "queue"
	ScriptVersionGetResponseResourcesBindingsTypeR2Bucket               ScriptVersionGetResponseResourcesBindingsType = "r2_bucket"
	ScriptVersionGetResponseResourcesBindingsTypeSecretText             ScriptVersionGetResponseResourcesBindingsType = "secret_text"
	ScriptVersionGetResponseResourcesBindingsTypeService                ScriptVersionGetResponseResourcesBindingsType = "service"
	ScriptVersionGetResponseResourcesBindingsTypeTailConsumer           ScriptVersionGetResponseResourcesBindingsType = "tail_consumer"
	ScriptVersionGetResponseResourcesBindingsTypeVectorize              ScriptVersionGetResponseResourcesBindingsType = "vectorize"
	ScriptVersionGetResponseResourcesBindingsTypeVersionMetadata        ScriptVersionGetResponseResourcesBindingsType = "version_metadata"
	ScriptVersionGetResponseResourcesBindingsTypeSecretsStoreSecret     ScriptVersionGetResponseResourcesBindingsType = "secrets_store_secret"
	ScriptVersionGetResponseResourcesBindingsTypeSecretKey              ScriptVersionGetResponseResourcesBindingsType = "secret_key"
	ScriptVersionGetResponseResourcesBindingsTypeWorkflow               ScriptVersionGetResponseResourcesBindingsType = "workflow"
)

func (r ScriptVersionGetResponseResourcesBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsTypeAI, ScriptVersionGetResponseResourcesBindingsTypeAnalyticsEngine, ScriptVersionGetResponseResourcesBindingsTypeAssets, ScriptVersionGetResponseResourcesBindingsTypeBrowser, ScriptVersionGetResponseResourcesBindingsTypeD1, ScriptVersionGetResponseResourcesBindingsTypeDispatchNamespace, ScriptVersionGetResponseResourcesBindingsTypeDurableObjectNamespace, ScriptVersionGetResponseResourcesBindingsTypeHyperdrive, ScriptVersionGetResponseResourcesBindingsTypeJson, ScriptVersionGetResponseResourcesBindingsTypeKVNamespace, ScriptVersionGetResponseResourcesBindingsTypeMTLSCertificate, ScriptVersionGetResponseResourcesBindingsTypePlainText, ScriptVersionGetResponseResourcesBindingsTypePipelines, ScriptVersionGetResponseResourcesBindingsTypeQueue, ScriptVersionGetResponseResourcesBindingsTypeR2Bucket, ScriptVersionGetResponseResourcesBindingsTypeSecretText, ScriptVersionGetResponseResourcesBindingsTypeService, ScriptVersionGetResponseResourcesBindingsTypeTailConsumer, ScriptVersionGetResponseResourcesBindingsTypeVectorize, ScriptVersionGetResponseResourcesBindingsTypeVersionMetadata, ScriptVersionGetResponseResourcesBindingsTypeSecretsStoreSecret, ScriptVersionGetResponseResourcesBindingsTypeSecretKey, ScriptVersionGetResponseResourcesBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsFormat string

const (
	ScriptVersionGetResponseResourcesBindingsFormatRaw   ScriptVersionGetResponseResourcesBindingsFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsFormatPkcs8 ScriptVersionGetResponseResourcesBindingsFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsFormatSpki  ScriptVersionGetResponseResourcesBindingsFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsFormatJwk   ScriptVersionGetResponseResourcesBindingsFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsFormatRaw, ScriptVersionGetResponseResourcesBindingsFormatPkcs8, ScriptVersionGetResponseResourcesBindingsFormatSpki, ScriptVersionGetResponseResourcesBindingsFormatJwk:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesScript struct {
	Etag             string                                                `json:"etag"`
	Handlers         []string                                              `json:"handlers"`
	LastDeployedFrom string                                                `json:"last_deployed_from"`
	NamedHandlers    []ScriptVersionGetResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             scriptVersionGetResponseResourcesScriptJSON           `json:"-"`
}

// scriptVersionGetResponseResourcesScriptJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesScript]
type scriptVersionGetResponseResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptNamedHandler struct {
	Handlers []string                                                `json:"handlers"`
	Name     string                                                  `json:"name"`
	JSON     scriptVersionGetResponseResourcesScriptNamedHandlerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptNamedHandlerJSON contains the JSON
// metadata for the struct [ScriptVersionGetResponseResourcesScriptNamedHandler]
type scriptVersionGetResponseResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntime struct {
	CompatibilityDate  string                                                   `json:"compatibility_date"`
	CompatibilityFlags []string                                                 `json:"compatibility_flags"`
	Limits             ScriptVersionGetResponseResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                                   `json:"migration_tag"`
	UsageModel         ScriptVersionGetResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               scriptVersionGetResponseResourcesScriptRuntimeJSON       `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeJSON contains the JSON metadata
// for the struct [ScriptVersionGetResponseResourcesScriptRuntime]
type scriptVersionGetResponseResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntimeLimits struct {
	CPUMs int64                                                    `json:"cpu_ms"`
	JSON  scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON contains the JSON
// metadata for the struct [ScriptVersionGetResponseResourcesScriptRuntimeLimits]
type scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntimeUsageModel string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelBundled  ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "bundled"
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelUnbound  ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "unbound"
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelStandard ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "standard"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeUsageModelBundled, ScriptVersionGetResponseResourcesScriptRuntimeUsageModelUnbound, ScriptVersionGetResponseResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionGetResponseMetadata struct {
	AuthorEmail string                                 `json:"author_email"`
	AuthorID    string                                 `json:"author_id"`
	CreatedOn   string                                 `json:"created_on"`
	HasPreview  bool                                   `json:"hasPreview"`
	ModifiedOn  string                                 `json:"modified_on"`
	Source      ScriptVersionGetResponseMetadataSource `json:"source"`
	JSON        scriptVersionGetResponseMetadataJSON   `json:"-"`
}

// scriptVersionGetResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseMetadata]
type scriptVersionGetResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseMetadataSource string

const (
	ScriptVersionGetResponseMetadataSourceUnknown      ScriptVersionGetResponseMetadataSource = "unknown"
	ScriptVersionGetResponseMetadataSourceAPI          ScriptVersionGetResponseMetadataSource = "api"
	ScriptVersionGetResponseMetadataSourceWrangler     ScriptVersionGetResponseMetadataSource = "wrangler"
	ScriptVersionGetResponseMetadataSourceTerraform    ScriptVersionGetResponseMetadataSource = "terraform"
	ScriptVersionGetResponseMetadataSourceDash         ScriptVersionGetResponseMetadataSource = "dash"
	ScriptVersionGetResponseMetadataSourceDashTemplate ScriptVersionGetResponseMetadataSource = "dash_template"
	ScriptVersionGetResponseMetadataSourceIntegration  ScriptVersionGetResponseMetadataSource = "integration"
	ScriptVersionGetResponseMetadataSourceQuickEditor  ScriptVersionGetResponseMetadataSource = "quick_editor"
	ScriptVersionGetResponseMetadataSourcePlayground   ScriptVersionGetResponseMetadataSource = "playground"
	ScriptVersionGetResponseMetadataSourceWorkersci    ScriptVersionGetResponseMetadataSource = "workersci"
)

func (r ScriptVersionGetResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseMetadataSourceUnknown, ScriptVersionGetResponseMetadataSourceAPI, ScriptVersionGetResponseMetadataSourceWrangler, ScriptVersionGetResponseMetadataSourceTerraform, ScriptVersionGetResponseMetadataSourceDash, ScriptVersionGetResponseMetadataSourceDashTemplate, ScriptVersionGetResponseMetadataSourceIntegration, ScriptVersionGetResponseMetadataSourceQuickEditor, ScriptVersionGetResponseMetadataSourcePlayground, ScriptVersionGetResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptVersionNewParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r ScriptVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ScriptVersionNewParamsMetadata struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker, which is required for
	// Version Upload.
	MainModule  param.Field[string]                                    `json:"main_module,required"`
	Annotations param.Field[ScriptVersionNewParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptVersionNewParamsMetadataBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptVersionNewParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptVersionNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataAnnotations struct {
	// Human-readable message about the version. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r ScriptVersionNewParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsType] `json:"type,required"`
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
	Format param.Field[ScriptVersionNewParamsMetadataBindingsFormat] `json:"format"`
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

func (r ScriptVersionNewParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBinding) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow],
// [ScriptVersionNewParamsMetadataBinding].
type ScriptVersionNewParamsMetadataBindingUnion interface {
	implementsScriptVersionNewParamsMetadataBindingUnion()
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsType string

const (
	ScriptVersionNewParamsMetadataBindingsTypeAI                     ScriptVersionNewParamsMetadataBindingsType = "ai"
	ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsTypeAssets                 ScriptVersionNewParamsMetadataBindingsType = "assets"
	ScriptVersionNewParamsMetadataBindingsTypeBrowser                ScriptVersionNewParamsMetadataBindingsType = "browser"
	ScriptVersionNewParamsMetadataBindingsTypeD1                     ScriptVersionNewParamsMetadataBindingsType = "d1"
	ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsTypeJson                   ScriptVersionNewParamsMetadataBindingsType = "json"
	ScriptVersionNewParamsMetadataBindingsTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsTypePlainText              ScriptVersionNewParamsMetadataBindingsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsTypePipelines              ScriptVersionNewParamsMetadataBindingsType = "pipelines"
	ScriptVersionNewParamsMetadataBindingsTypeQueue                  ScriptVersionNewParamsMetadataBindingsType = "queue"
	ScriptVersionNewParamsMetadataBindingsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsTypeSecretText             ScriptVersionNewParamsMetadataBindingsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsTypeService                ScriptVersionNewParamsMetadataBindingsType = "service"
	ScriptVersionNewParamsMetadataBindingsTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsTypeVectorize              ScriptVersionNewParamsMetadataBindingsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsType = "version_metadata"
	ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret     ScriptVersionNewParamsMetadataBindingsType = "secrets_store_secret"
	ScriptVersionNewParamsMetadataBindingsTypeSecretKey              ScriptVersionNewParamsMetadataBindingsType = "secret_key"
	ScriptVersionNewParamsMetadataBindingsTypeWorkflow               ScriptVersionNewParamsMetadataBindingsType = "workflow"
)

func (r ScriptVersionNewParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsTypeAI, ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsTypeAssets, ScriptVersionNewParamsMetadataBindingsTypeBrowser, ScriptVersionNewParamsMetadataBindingsTypeD1, ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsTypeJson, ScriptVersionNewParamsMetadataBindingsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsTypePlainText, ScriptVersionNewParamsMetadataBindingsTypePipelines, ScriptVersionNewParamsMetadataBindingsTypeQueue, ScriptVersionNewParamsMetadataBindingsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsTypeSecretText, ScriptVersionNewParamsMetadataBindingsTypeService, ScriptVersionNewParamsMetadataBindingsTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsTypeVectorize, ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata, ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret, ScriptVersionNewParamsMetadataBindingsTypeSecretKey, ScriptVersionNewParamsMetadataBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewParamsMetadataBindingsFormat string

const (
	ScriptVersionNewParamsMetadataBindingsFormatRaw   ScriptVersionNewParamsMetadataBindingsFormat = "raw"
	ScriptVersionNewParamsMetadataBindingsFormatPkcs8 ScriptVersionNewParamsMetadataBindingsFormat = "pkcs8"
	ScriptVersionNewParamsMetadataBindingsFormatSpki  ScriptVersionNewParamsMetadataBindingsFormat = "spki"
	ScriptVersionNewParamsMetadataBindingsFormatJwk   ScriptVersionNewParamsMetadataBindingsFormat = "jwk"
)

func (r ScriptVersionNewParamsMetadataBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsFormatRaw, ScriptVersionNewParamsMetadataBindingsFormatPkcs8, ScriptVersionNewParamsMetadataBindingsFormatSpki, ScriptVersionNewParamsMetadataBindingsFormatJwk:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptVersionNewParamsMetadataUsageModel string

const (
	ScriptVersionNewParamsMetadataUsageModelStandard ScriptVersionNewParamsMetadataUsageModel = "standard"
)

func (r ScriptVersionNewParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionNewResponseEnvelope struct {
	Errors   []ScriptVersionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptVersionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptVersionNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptVersionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionNewResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseEnvelope]
type scriptVersionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptVersionNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptVersionNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptVersionNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseEnvelopeErrors]
type scriptVersionNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptVersionNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptVersionNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptVersionNewResponseEnvelopeErrorsSource]
type scriptVersionNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptVersionNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptVersionNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptVersionNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseEnvelopeMessages]
type scriptVersionNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptVersionNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptVersionNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseEnvelopeMessagesSource]
type scriptVersionNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptVersionNewResponseEnvelopeSuccess bool

const (
	ScriptVersionNewResponseEnvelopeSuccessTrue ScriptVersionNewResponseEnvelopeSuccess = true
)

func (r ScriptVersionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Only return versions that can be used in a deployment. Ignores pagination.
	Deployable param.Field[bool] `query:"deployable"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ScriptVersionListParams]'s query parameters as
// `url.Values`.
func (r ScriptVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptVersionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionGetResponseEnvelope struct {
	Errors   []ScriptVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptVersionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptVersionGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseEnvelope]
type scriptVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseEnvelopeErrors]
type scriptVersionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptVersionGetResponseEnvelopeErrorsSource]
type scriptVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseEnvelopeMessages]
type scriptVersionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptVersionGetResponseEnvelopeMessagesSource]
type scriptVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptVersionGetResponseEnvelopeSuccess bool

const (
	ScriptVersionGetResponseEnvelopeSuccessTrue ScriptVersionGetResponseEnvelopeSuccess = true
)

func (r ScriptVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
