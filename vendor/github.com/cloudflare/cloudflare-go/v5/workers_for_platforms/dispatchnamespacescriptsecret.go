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

// DispatchNamespaceScriptSecretService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptSecretService] method instead.
type DispatchNamespaceScriptSecretService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSecretService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSecretService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSecretService) {
	r = &DispatchNamespaceScriptSecretService{}
	r.Options = opts
	return
}

// Add a secret to a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSecretUpdateParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSecretUpdateResponse, err error) {
	var env DispatchNamespaceScriptSecretUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List secrets bound to a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) List(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceScriptSecretListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", query.AccountID, dispatchNamespace, scriptName)
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

// List secrets bound to a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) ListAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSecretListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceScriptSecretListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, dispatchNamespace, scriptName, query, opts...))
}

// Remove a secret from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, secretName string, body DispatchNamespaceScriptSecretDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSecretDeleteResponse, err error) {
	var env DispatchNamespaceScriptSecretDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", body.AccountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a given secret binding (value omitted) on a script uploaded to a Workers for
// Platforms namespace.
func (r *DispatchNamespaceScriptSecretService) Get(ctx context.Context, dispatchNamespace string, scriptName string, secretName string, query DispatchNamespaceScriptSecretGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSecretGetResponse, err error) {
	var env DispatchNamespaceScriptSecretGetResponseEnvelope
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", query.AccountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A secret value accessible through a binding.
type DispatchNamespaceScriptSecretUpdateResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretUpdateResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretUpdateResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                                     `json:"usages"`
	JSON   dispatchNamespaceScriptSecretUpdateResponseJSON `json:"-"`
	union  DispatchNamespaceScriptSecretUpdateResponseUnion
}

// dispatchNamespaceScriptSecretUpdateResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSecretUpdateResponse]
type dispatchNamespaceScriptSecretUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dispatchNamespaceScriptSecretUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSecretUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSecretUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSecretUpdateResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey].
func (r DispatchNamespaceScriptSecretUpdateResponse) AsUnion() DispatchNamespaceScriptSecretUpdateResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText] or
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey].
type DispatchNamespaceScriptSecretUpdateResponseUnion interface {
	implementsDispatchNamespaceScriptSecretUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSecretUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText]
type dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSecretUpdateResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey]
type dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSecretUpdateResponse() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateResponseType string

const (
	DispatchNamespaceScriptSecretUpdateResponseTypeSecretText DispatchNamespaceScriptSecretUpdateResponseType = "secret_text"
	DispatchNamespaceScriptSecretUpdateResponseTypeSecretKey  DispatchNamespaceScriptSecretUpdateResponseType = "secret_key"
)

func (r DispatchNamespaceScriptSecretUpdateResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseTypeSecretText, DispatchNamespaceScriptSecretUpdateResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretUpdateResponseFormat string

const (
	DispatchNamespaceScriptSecretUpdateResponseFormatRaw   DispatchNamespaceScriptSecretUpdateResponseFormat = "raw"
	DispatchNamespaceScriptSecretUpdateResponseFormatPkcs8 DispatchNamespaceScriptSecretUpdateResponseFormat = "pkcs8"
	DispatchNamespaceScriptSecretUpdateResponseFormatSpki  DispatchNamespaceScriptSecretUpdateResponseFormat = "spki"
	DispatchNamespaceScriptSecretUpdateResponseFormatJwk   DispatchNamespaceScriptSecretUpdateResponseFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretUpdateResponseFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseFormatRaw, DispatchNamespaceScriptSecretUpdateResponseFormatPkcs8, DispatchNamespaceScriptSecretUpdateResponseFormatSpki, DispatchNamespaceScriptSecretUpdateResponseFormatJwk:
		return true
	}
	return false
}

// A secret value accessible through a binding.
type DispatchNamespaceScriptSecretListResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretListResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretListResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                                   `json:"usages"`
	JSON   dispatchNamespaceScriptSecretListResponseJSON `json:"-"`
	union  DispatchNamespaceScriptSecretListResponseUnion
}

// dispatchNamespaceScriptSecretListResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSecretListResponse]
type dispatchNamespaceScriptSecretListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dispatchNamespaceScriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSecretListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSecretListResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey].
func (r DispatchNamespaceScriptSecretListResponse) AsUnion() DispatchNamespaceScriptSecretListResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText] or
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey].
type DispatchNamespaceScriptSecretListResponseUnion interface {
	implementsDispatchNamespaceScriptSecretListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSecretListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText]
type dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSecretListResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey]
type dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSecretListResponse() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSecretListResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretListResponseType string

const (
	DispatchNamespaceScriptSecretListResponseTypeSecretText DispatchNamespaceScriptSecretListResponseType = "secret_text"
	DispatchNamespaceScriptSecretListResponseTypeSecretKey  DispatchNamespaceScriptSecretListResponseType = "secret_key"
)

func (r DispatchNamespaceScriptSecretListResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseTypeSecretText, DispatchNamespaceScriptSecretListResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretListResponseFormat string

const (
	DispatchNamespaceScriptSecretListResponseFormatRaw   DispatchNamespaceScriptSecretListResponseFormat = "raw"
	DispatchNamespaceScriptSecretListResponseFormatPkcs8 DispatchNamespaceScriptSecretListResponseFormat = "pkcs8"
	DispatchNamespaceScriptSecretListResponseFormatSpki  DispatchNamespaceScriptSecretListResponseFormat = "spki"
	DispatchNamespaceScriptSecretListResponseFormatJwk   DispatchNamespaceScriptSecretListResponseFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretListResponseFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretListResponseFormatRaw, DispatchNamespaceScriptSecretListResponseFormatPkcs8, DispatchNamespaceScriptSecretListResponseFormatSpki, DispatchNamespaceScriptSecretListResponseFormatJwk:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretDeleteResponse = interface{}

// A secret value accessible through a binding.
type DispatchNamespaceScriptSecretGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretGetResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretGetResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                                  `json:"usages"`
	JSON   dispatchNamespaceScriptSecretGetResponseJSON `json:"-"`
	union  DispatchNamespaceScriptSecretGetResponseUnion
}

// dispatchNamespaceScriptSecretGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSecretGetResponse]
type dispatchNamespaceScriptSecretGetResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dispatchNamespaceScriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSecretGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSecretGetResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText],
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey].
func (r DispatchNamespaceScriptSecretGetResponse) AsUnion() DispatchNamespaceScriptSecretGetResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText] or
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey].
type DispatchNamespaceScriptSecretGetResponseUnion interface {
	implementsDispatchNamespaceScriptSecretGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSecretGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText]
type dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSecretGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey]
type dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSecretGetResponse() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSecretGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretGetResponseType string

const (
	DispatchNamespaceScriptSecretGetResponseTypeSecretText DispatchNamespaceScriptSecretGetResponseType = "secret_text"
	DispatchNamespaceScriptSecretGetResponseTypeSecretKey  DispatchNamespaceScriptSecretGetResponseType = "secret_key"
)

func (r DispatchNamespaceScriptSecretGetResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseTypeSecretText, DispatchNamespaceScriptSecretGetResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretGetResponseFormat string

const (
	DispatchNamespaceScriptSecretGetResponseFormatRaw   DispatchNamespaceScriptSecretGetResponseFormat = "raw"
	DispatchNamespaceScriptSecretGetResponseFormatPkcs8 DispatchNamespaceScriptSecretGetResponseFormat = "pkcs8"
	DispatchNamespaceScriptSecretGetResponseFormatSpki  DispatchNamespaceScriptSecretGetResponseFormat = "spki"
	DispatchNamespaceScriptSecretGetResponseFormatJwk   DispatchNamespaceScriptSecretGetResponseFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretGetResponseFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseFormatRaw, DispatchNamespaceScriptSecretGetResponseFormatPkcs8, DispatchNamespaceScriptSecretGetResponseFormatSpki, DispatchNamespaceScriptSecretGetResponseFormatJwk:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A secret value accessible through a binding.
	Body DispatchNamespaceScriptSecretUpdateParamsBodyUnion `json:"body,required"`
}

func (r DispatchNamespaceScriptSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// A secret value accessible through a binding.
type DispatchNamespaceScriptSecretUpdateParamsBody struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type      param.Field[DispatchNamespaceScriptSecretUpdateParamsBodyType] `json:"type,required"`
	Algorithm param.Field[interface{}]                                       `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptSecretUpdateParamsBodyFormat] `json:"format"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// The secret value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
}

func (r DispatchNamespaceScriptSecretUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSecretUpdateParamsBody) implementsDispatchNamespaceScriptSecretUpdateParamsBodyUnion() {
}

// A secret value accessible through a binding.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey],
// [DispatchNamespaceScriptSecretUpdateParamsBody].
type DispatchNamespaceScriptSecretUpdateParamsBodyUnion interface {
	implementsDispatchNamespaceScriptSecretUpdateParamsBodyUnion()
}

type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSecretUpdateParamsBodyUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptSecretUpdateParamsBodyUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSecretUpdateParamsBodyType string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyTypeSecretText DispatchNamespaceScriptSecretUpdateParamsBodyType = "secret_text"
	DispatchNamespaceScriptSecretUpdateParamsBodyTypeSecretKey  DispatchNamespaceScriptSecretUpdateParamsBodyType = "secret_key"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyTypeSecretText, DispatchNamespaceScriptSecretUpdateParamsBodyTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptSecretUpdateParamsBodyFormat string

const (
	DispatchNamespaceScriptSecretUpdateParamsBodyFormatRaw   DispatchNamespaceScriptSecretUpdateParamsBodyFormat = "raw"
	DispatchNamespaceScriptSecretUpdateParamsBodyFormatPkcs8 DispatchNamespaceScriptSecretUpdateParamsBodyFormat = "pkcs8"
	DispatchNamespaceScriptSecretUpdateParamsBodyFormatSpki  DispatchNamespaceScriptSecretUpdateParamsBodyFormat = "spki"
	DispatchNamespaceScriptSecretUpdateParamsBodyFormatJwk   DispatchNamespaceScriptSecretUpdateParamsBodyFormat = "jwk"
)

func (r DispatchNamespaceScriptSecretUpdateParamsBodyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateParamsBodyFormatRaw, DispatchNamespaceScriptSecretUpdateParamsBodyFormatPkcs8, DispatchNamespaceScriptSecretUpdateParamsBodyFormatSpki, DispatchNamespaceScriptSecretUpdateParamsBodyFormatJwk:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A secret value accessible through a binding.
	Result DispatchNamespaceScriptSecretUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSecretUpdateResponseEnvelope]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code,required"`
	Message          string                                                            `json:"message,required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSecretDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSecretDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSecretDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceScriptSecretDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSecretDeleteResponseEnvelope]
type dispatchNamespaceScriptSecretDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrors]
type dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code,required"`
	Message          string                                                            `json:"message,required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessages]
type dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccessTrue DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSecretGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSecretGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptSecretGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSecretGetResponseEnvelopeMessages `json:"messages,required"`
	// A secret value accessible through a binding.
	Result DispatchNamespaceScriptSecretGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSecretGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSecretGetResponseEnvelope]
type dispatchNamespaceScriptSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretGetResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSecretGetResponseEnvelopeErrors]
type dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretGetResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           DispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSecretGetResponseEnvelopeMessages]
type dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSecretGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSecretGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSecretGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSecretGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
