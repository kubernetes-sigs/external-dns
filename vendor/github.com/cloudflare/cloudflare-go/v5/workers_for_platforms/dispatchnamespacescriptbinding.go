// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptBindingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptBindingService] method instead.
type DispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptBindingService) {
	r = &DispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceScriptBindingGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
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

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) GetAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceScriptBindingGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, dispatchNamespace, scriptName, query, opts...))
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptBindingGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseType `json:"type,required"`
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
	Format DispatchNamespaceScriptBindingGetResponseFormat `json:"format"`
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
	// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound].
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
	// [[]DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                        `json:"workflow_name"`
	JSON         dispatchNamespaceScriptBindingGetResponseJSON `json:"-"`
	union        DispatchNamespaceScriptBindingGetResponseUnion
}

// dispatchNamespaceScriptBindingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptBindingGetResponse]
type dispatchNamespaceScriptBindingGetResponseJSON struct {
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

func (r dispatchNamespaceScriptBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptBindingGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptBindingGetResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow].
func (r DispatchNamespaceScriptBindingGetResponse) AsUnion() DispatchNamespaceScriptBindingGetResponseUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret],
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey] or
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow].
type DispatchNamespaceScriptBindingGetResponseUnion interface {
	implementsDispatchNamespaceScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptBindingGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowser) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelines) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                  `json:"script_name"`
	JSON       dispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseType string

const (
	DispatchNamespaceScriptBindingGetResponseTypeAI                     DispatchNamespaceScriptBindingGetResponseType = "ai"
	DispatchNamespaceScriptBindingGetResponseTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseTypeAssets                 DispatchNamespaceScriptBindingGetResponseType = "assets"
	DispatchNamespaceScriptBindingGetResponseTypeBrowser                DispatchNamespaceScriptBindingGetResponseType = "browser"
	DispatchNamespaceScriptBindingGetResponseTypeD1                     DispatchNamespaceScriptBindingGetResponseType = "d1"
	DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseTypeJson                   DispatchNamespaceScriptBindingGetResponseType = "json"
	DispatchNamespaceScriptBindingGetResponseTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseTypePlainText              DispatchNamespaceScriptBindingGetResponseType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseTypePipelines              DispatchNamespaceScriptBindingGetResponseType = "pipelines"
	DispatchNamespaceScriptBindingGetResponseTypeQueue                  DispatchNamespaceScriptBindingGetResponseType = "queue"
	DispatchNamespaceScriptBindingGetResponseTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseTypeSecretText             DispatchNamespaceScriptBindingGetResponseType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseTypeService                DispatchNamespaceScriptBindingGetResponseType = "service"
	DispatchNamespaceScriptBindingGetResponseTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseTypeVectorize              DispatchNamespaceScriptBindingGetResponseType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseType = "version_metadata"
	DispatchNamespaceScriptBindingGetResponseTypeSecretsStoreSecret     DispatchNamespaceScriptBindingGetResponseType = "secrets_store_secret"
	DispatchNamespaceScriptBindingGetResponseTypeSecretKey              DispatchNamespaceScriptBindingGetResponseType = "secret_key"
	DispatchNamespaceScriptBindingGetResponseTypeWorkflow               DispatchNamespaceScriptBindingGetResponseType = "workflow"
)

func (r DispatchNamespaceScriptBindingGetResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseTypeAI, DispatchNamespaceScriptBindingGetResponseTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseTypeAssets, DispatchNamespaceScriptBindingGetResponseTypeBrowser, DispatchNamespaceScriptBindingGetResponseTypeD1, DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseTypeJson, DispatchNamespaceScriptBindingGetResponseTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseTypePlainText, DispatchNamespaceScriptBindingGetResponseTypePipelines, DispatchNamespaceScriptBindingGetResponseTypeQueue, DispatchNamespaceScriptBindingGetResponseTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseTypeSecretText, DispatchNamespaceScriptBindingGetResponseTypeService, DispatchNamespaceScriptBindingGetResponseTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseTypeVectorize, DispatchNamespaceScriptBindingGetResponseTypeVersionMetadata, DispatchNamespaceScriptBindingGetResponseTypeSecretsStoreSecret, DispatchNamespaceScriptBindingGetResponseTypeSecretKey, DispatchNamespaceScriptBindingGetResponseTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptBindingGetResponseFormat string

const (
	DispatchNamespaceScriptBindingGetResponseFormatRaw   DispatchNamespaceScriptBindingGetResponseFormat = "raw"
	DispatchNamespaceScriptBindingGetResponseFormatPkcs8 DispatchNamespaceScriptBindingGetResponseFormat = "pkcs8"
	DispatchNamespaceScriptBindingGetResponseFormatSpki  DispatchNamespaceScriptBindingGetResponseFormat = "spki"
	DispatchNamespaceScriptBindingGetResponseFormatJwk   DispatchNamespaceScriptBindingGetResponseFormat = "jwk"
)

func (r DispatchNamespaceScriptBindingGetResponseFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseFormatRaw, DispatchNamespaceScriptBindingGetResponseFormatPkcs8, DispatchNamespaceScriptBindingGetResponseFormatSpki, DispatchNamespaceScriptBindingGetResponseFormatJwk:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
