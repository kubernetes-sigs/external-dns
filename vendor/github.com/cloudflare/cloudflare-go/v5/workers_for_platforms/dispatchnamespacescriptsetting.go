// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/workers"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptSettingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptSettingService] method instead.
type DispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSettingService) {
	r = &DispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings.
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingEditResponse, err error) {
	var env DispatchNamespaceScriptSettingEditResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingGetResponse, err error) {
	var env DispatchNamespaceScriptSettingGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptSettingEditResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []DispatchNamespaceScriptSettingEditResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingEditResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability DispatchNamespaceScriptSettingEditResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement DispatchNamespaceScriptSettingEditResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers.
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptSettingEditResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingEditResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSettingEditResponse]
type dispatchNamespaceScriptSettingEditResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptSettingEditResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsType `json:"type,required"`
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
	Format DispatchNamespaceScriptSettingEditResponseBindingsFormat `json:"format"`
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
	// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	// [[]DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                                `json:"workflow_name"`
	JSON         dispatchNamespaceScriptSettingEditResponseBindingJSON `json:"-"`
	union        DispatchNamespaceScriptSettingEditResponseBindingsUnion
}

// dispatchNamespaceScriptSettingEditResponseBindingJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingEditResponseBinding]
type dispatchNamespaceScriptSettingEditResponseBindingJSON struct {
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

func (r dispatchNamespaceScriptSettingEditResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingEditResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingEditResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingEditResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow].
func (r DispatchNamespaceScriptSettingEditResponseBinding) AsUnion() DispatchNamespaceScriptSettingEditResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey]
// or
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow].
type DispatchNamespaceScriptSettingEditResponseBindingsUnion interface {
	implementsDispatchNamespaceScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                  `json:"service"`
	JSON    dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                         `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                           `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeBrowser                DispatchNamespaceScriptSettingEditResponseBindingsType = "browser"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsTypePipelines              DispatchNamespaceScriptSettingEditResponseBindingsType = "pipelines"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeService                DispatchNamespaceScriptSettingEditResponseBindingsType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsType = "version_metadata"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptSettingEditResponseBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretKey              DispatchNamespaceScriptSettingEditResponseBindingsType = "secret_key"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeWorkflow               DispatchNamespaceScriptSettingEditResponseBindingsType = "workflow"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsTypeBrowser, DispatchNamespaceScriptSettingEditResponseBindingsTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsTypePipelines, DispatchNamespaceScriptSettingEditResponseBindingsTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsTypeService, DispatchNamespaceScriptSettingEditResponseBindingsTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsTypeVersionMetadata, DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretKey, DispatchNamespaceScriptSettingEditResponseBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingEditResponseBindingsFormat string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsFormatRaw   DispatchNamespaceScriptSettingEditResponseBindingsFormat = "raw"
	DispatchNamespaceScriptSettingEditResponseBindingsFormatPkcs8 DispatchNamespaceScriptSettingEditResponseBindingsFormat = "pkcs8"
	DispatchNamespaceScriptSettingEditResponseBindingsFormatSpki  DispatchNamespaceScriptSettingEditResponseBindingsFormat = "spki"
	DispatchNamespaceScriptSettingEditResponseBindingsFormatJwk   DispatchNamespaceScriptSettingEditResponseBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsFormatRaw, DispatchNamespaceScriptSettingEditResponseBindingsFormatPkcs8, DispatchNamespaceScriptSettingEditResponseBindingsFormatSpki, DispatchNamespaceScriptSettingEditResponseBindingsFormatJwk:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                                `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingEditResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingEditResponseLimits]
type dispatchNamespaceScriptSettingEditResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditResponseMigrations struct {
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes"`
	// This field can have the runtime type of [[]string].
	NewSqliteClasses interface{} `json:"new_sqlite_classes"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                              `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingEditResponseMigrationsJSON `json:"-"`
	union              DispatchNamespaceScriptSettingEditResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingEditResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingEditResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingEditResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations].
func (r DispatchNamespaceScriptSettingEditResponseMigrations) AsUnion() DispatchNamespaceScriptSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations].
type DispatchNamespaceScriptSettingEditResponseMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations struct {
	JSON dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingEditResponseMigrations() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingEditResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate,nullable"`
	// Log settings for the Worker.
	Logs DispatchNamespaceScriptSettingEditResponseObservabilityLogs `json:"logs,nullable"`
	JSON dispatchNamespaceScriptSettingEditResponseObservabilityJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseObservabilityJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseObservability]
type dispatchNamespaceScriptSettingEditResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type DispatchNamespaceScriptSettingEditResponseObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64                                                         `json:"head_sampling_rate,nullable"`
	JSON             dispatchNamespaceScriptSettingEditResponseObservabilityLogsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseObservabilityLogsJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseObservabilityLogs]
type dispatchNamespaceScriptSettingEditResponseObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingEditResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponsePlacement]
type dispatchNamespaceScriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingEditResponsePlacementModeSmart DispatchNamespaceScriptSettingEditResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingEditResponseUsageModel string

const (
	DispatchNamespaceScriptSettingEditResponseUsageModelStandard DispatchNamespaceScriptSettingEditResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingEditResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []DispatchNamespaceScriptSettingGetResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingGetResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability DispatchNamespaceScriptSettingGetResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement DispatchNamespaceScriptSettingGetResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers.
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptSettingGetResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingGetResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSettingGetResponse]
type dispatchNamespaceScriptSettingGetResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptSettingGetResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsType `json:"type,required"`
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
	Format DispatchNamespaceScriptSettingGetResponseBindingsFormat `json:"format"`
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
	// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	// [[]DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                               `json:"workflow_name"`
	JSON         dispatchNamespaceScriptSettingGetResponseBindingJSON `json:"-"`
	union        DispatchNamespaceScriptSettingGetResponseBindingsUnion
}

// dispatchNamespaceScriptSettingGetResponseBindingJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseBinding]
type dispatchNamespaceScriptSettingGetResponseBindingJSON struct {
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

func (r dispatchNamespaceScriptSettingGetResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingGetResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingGetResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingGetResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow].
func (r DispatchNamespaceScriptSettingGetResponseBinding) AsUnion() DispatchNamespaceScriptSettingGetResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey]
// or
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow].
type DispatchNamespaceScriptSettingGetResponseBindingsUnion interface {
	implementsDispatchNamespaceScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                 `json:"service"`
	JSON    dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                        `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                          `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeBrowser                DispatchNamespaceScriptSettingGetResponseBindingsType = "browser"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsTypePipelines              DispatchNamespaceScriptSettingGetResponseBindingsType = "pipelines"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeService                DispatchNamespaceScriptSettingGetResponseBindingsType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsType = "version_metadata"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptSettingGetResponseBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretKey              DispatchNamespaceScriptSettingGetResponseBindingsType = "secret_key"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeWorkflow               DispatchNamespaceScriptSettingGetResponseBindingsType = "workflow"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsTypeBrowser, DispatchNamespaceScriptSettingGetResponseBindingsTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsTypePipelines, DispatchNamespaceScriptSettingGetResponseBindingsTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsTypeService, DispatchNamespaceScriptSettingGetResponseBindingsTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsTypeVersionMetadata, DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretKey, DispatchNamespaceScriptSettingGetResponseBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingGetResponseBindingsFormat string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsFormatRaw   DispatchNamespaceScriptSettingGetResponseBindingsFormat = "raw"
	DispatchNamespaceScriptSettingGetResponseBindingsFormatPkcs8 DispatchNamespaceScriptSettingGetResponseBindingsFormat = "pkcs8"
	DispatchNamespaceScriptSettingGetResponseBindingsFormatSpki  DispatchNamespaceScriptSettingGetResponseBindingsFormat = "spki"
	DispatchNamespaceScriptSettingGetResponseBindingsFormatJwk   DispatchNamespaceScriptSettingGetResponseBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsFormatRaw, DispatchNamespaceScriptSettingGetResponseBindingsFormatPkcs8, DispatchNamespaceScriptSettingGetResponseBindingsFormatSpki, DispatchNamespaceScriptSettingGetResponseBindingsFormatJwk:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingGetResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                               `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingGetResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseLimits]
type dispatchNamespaceScriptSettingGetResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingGetResponseMigrations struct {
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes"`
	// This field can have the runtime type of [[]string].
	NewSqliteClasses interface{} `json:"new_sqlite_classes"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                             `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingGetResponseMigrationsJSON `json:"-"`
	union              DispatchNamespaceScriptSettingGetResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingGetResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingGetResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingGetResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations].
func (r DispatchNamespaceScriptSettingGetResponseMigrations) AsUnion() DispatchNamespaceScriptSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations].
type DispatchNamespaceScriptSettingGetResponseMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations struct {
	JSON dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingGetResponseMigrations() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingGetResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate,nullable"`
	// Log settings for the Worker.
	Logs DispatchNamespaceScriptSettingGetResponseObservabilityLogs `json:"logs,nullable"`
	JSON dispatchNamespaceScriptSettingGetResponseObservabilityJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseObservabilityJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseObservability]
type dispatchNamespaceScriptSettingGetResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type DispatchNamespaceScriptSettingGetResponseObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64                                                        `json:"head_sampling_rate,nullable"`
	JSON             dispatchNamespaceScriptSettingGetResponseObservabilityLogsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseObservabilityLogsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseObservabilityLogs]
type dispatchNamespaceScriptSettingGetResponseObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingGetResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponsePlacement]
type dispatchNamespaceScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingGetResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingGetResponsePlacementModeSmart DispatchNamespaceScriptSettingGetResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingGetResponseUsageModel string

const (
	DispatchNamespaceScriptSettingGetResponseUsageModelStandard DispatchNamespaceScriptSettingGetResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingGetResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParams struct {
	// Identifier.
	AccountID param.Field[string]                                           `path:"account_id,required"`
	Settings  param.Field[DispatchNamespaceScriptSettingEditParamsSettings] `json:"settings"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type DispatchNamespaceScriptSettingEditParamsSettings struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptSettingEditParamsSettingsBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits param.Field[DispatchNamespaceScriptSettingEditParamsSettingsLimits] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptSettingEditParamsSettingsObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[DispatchNamespaceScriptSettingEditParamsSettingsPlacement] `json:"placement"`
	// Tags to help you manage your Workers.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptSettingEditParamsSettingsUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptSettingEditParamsSettingsBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsType] `json:"type,required"`
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
	Format param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat] `json:"format"`
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBinding) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowser],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelines],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow],
// [DispatchNamespaceScriptSettingEditParamsSettingsBinding].
type DispatchNamespaceScriptSettingEditParamsSettingsBindingUnion interface {
	implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion()
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeBrowser                DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "browser"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePipelines              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "pipelines"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "version_metadata"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretKey              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "secret_key"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeWorkflow               DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "workflow"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeBrowser, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePipelines, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVersionMetadata, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretKey, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatRaw   DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat = "raw"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatPkcs8 DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat = "pkcs8"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatSpki  DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat = "spki"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatJwk   DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatRaw, DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatPkcs8, DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatSpki, DispatchNamespaceScriptSettingEditParamsSettingsBindingsFormatJwk:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsMigrations struct {
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptSettingEditParamsSettingsMigrations].
type DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion()
}

type DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[DispatchNamespaceScriptSettingEditParamsSettingsObservabilityLogs] `json:"logs"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsPlacementModeSmart DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsPlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingEditParamsSettingsUsageModel string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsUsageModelStandard DispatchNamespaceScriptSettingEditParamsSettingsUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingEditResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseEnvelope]
type dispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           DispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeErrors]
type dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           DispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeMessages]
type dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingGetResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseEnvelope]
type dispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseEnvelopeErrors struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           DispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeErrors]
type dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseEnvelopeMessages struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           DispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeMessages]
type dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
