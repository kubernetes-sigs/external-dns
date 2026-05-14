// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// ScriptSecretService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSecretService] method instead.
type ScriptSecretService struct {
	Options []option.RequestOption
}

// NewScriptSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSecretService(opts ...option.RequestOption) (r *ScriptSecretService) {
	r = &ScriptSecretService{}
	r.Options = opts
	return
}

// Add a secret to a script.
func (r *ScriptSecretService) Update(ctx context.Context, scriptName string, params ScriptSecretUpdateParams, opts ...option.RequestOption) (res *ScriptSecretUpdateResponse, err error) {
	var env ScriptSecretUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List secrets bound to a script.
func (r *ScriptSecretService) List(ctx context.Context, scriptName string, query ScriptSecretListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ScriptSecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", query.AccountID, scriptName)
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

// List secrets bound to a script.
func (r *ScriptSecretService) ListAutoPaging(ctx context.Context, scriptName string, query ScriptSecretListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ScriptSecretListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, scriptName, query, opts...))
}

// Remove a secret from a script.
func (r *ScriptSecretService) Delete(ctx context.Context, scriptName string, secretName string, body ScriptSecretDeleteParams, opts ...option.RequestOption) (res *ScriptSecretDeleteResponse, err error) {
	var env ScriptSecretDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", body.AccountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a given secret binding (value omitted) on a script.
func (r *ScriptSecretService) Get(ctx context.Context, scriptName string, secretName string, query ScriptSecretGetParams, opts ...option.RequestOption) (res *ScriptSecretGetResponse, err error) {
	var env ScriptSecretGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", query.AccountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A secret value accessible through a binding.
type ScriptSecretUpdateResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretUpdateResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretUpdateResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                    `json:"usages"`
	JSON   scriptSecretUpdateResponseJSON `json:"-"`
	union  ScriptSecretUpdateResponseUnion
}

// scriptSecretUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptSecretUpdateResponse]
type scriptSecretUpdateResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptSecretUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptSecretUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptSecretUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptSecretUpdateResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptSecretUpdateResponseWorkersBindingKindSecretText],
// [ScriptSecretUpdateResponseWorkersBindingKindSecretKey].
func (r ScriptSecretUpdateResponse) AsUnion() ScriptSecretUpdateResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by [ScriptSecretUpdateResponseWorkersBindingKindSecretText] or
// [ScriptSecretUpdateResponseWorkersBindingKindSecretKey].
type ScriptSecretUpdateResponseUnion interface {
	implementsScriptSecretUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSecretUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretUpdateResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretUpdateResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type ScriptSecretUpdateResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretUpdateResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptSecretUpdateResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptSecretUpdateResponseWorkersBindingKindSecretTextJSON contains the JSON
// metadata for the struct [ScriptSecretUpdateResponseWorkersBindingKindSecretText]
type scriptSecretUpdateResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretText) implementsScriptSecretUpdateResponse() {
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateResponseWorkersBindingKindSecretTextType string

const (
	ScriptSecretUpdateResponseWorkersBindingKindSecretTextTypeSecretText ScriptSecretUpdateResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretUpdateResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretUpdateResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON contains the JSON
// metadata for the struct [ScriptSecretUpdateResponseWorkersBindingKindSecretKey]
type scriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretKey) implementsScriptSecretUpdateResponse() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat string

const (
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatRaw   ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "raw"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatPkcs8 ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatSpki  ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "spki"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatJwk   ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatRaw, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatPkcs8, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatSpki, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateResponseWorkersBindingKindSecretKeyType string

const (
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyTypeSecretKey ScriptSecretUpdateResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage string

const (
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageEncrypt    ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDecrypt    ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageSign       ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "sign"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageVerify     ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "verify"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveKey  ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveBits ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageWrapKey    ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageEncrypt, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDecrypt, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageSign, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageVerify, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveKey, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageDeriveBits, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageWrapKey, ScriptSecretUpdateResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateResponseType string

const (
	ScriptSecretUpdateResponseTypeSecretText ScriptSecretUpdateResponseType = "secret_text"
	ScriptSecretUpdateResponseTypeSecretKey  ScriptSecretUpdateResponseType = "secret_key"
)

func (r ScriptSecretUpdateResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseTypeSecretText, ScriptSecretUpdateResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretUpdateResponseFormat string

const (
	ScriptSecretUpdateResponseFormatRaw   ScriptSecretUpdateResponseFormat = "raw"
	ScriptSecretUpdateResponseFormatPkcs8 ScriptSecretUpdateResponseFormat = "pkcs8"
	ScriptSecretUpdateResponseFormatSpki  ScriptSecretUpdateResponseFormat = "spki"
	ScriptSecretUpdateResponseFormatJwk   ScriptSecretUpdateResponseFormat = "jwk"
)

func (r ScriptSecretUpdateResponseFormat) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseFormatRaw, ScriptSecretUpdateResponseFormatPkcs8, ScriptSecretUpdateResponseFormatSpki, ScriptSecretUpdateResponseFormatJwk:
		return true
	}
	return false
}

// A secret value accessible through a binding.
type ScriptSecretListResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretListResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretListResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]ScriptSecretListResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                  `json:"usages"`
	JSON   scriptSecretListResponseJSON `json:"-"`
	union  ScriptSecretListResponseUnion
}

// scriptSecretListResponseJSON contains the JSON metadata for the struct
// [ScriptSecretListResponse]
type scriptSecretListResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptSecretListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptSecretListResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptSecretListResponseWorkersBindingKindSecretText],
// [ScriptSecretListResponseWorkersBindingKindSecretKey].
func (r ScriptSecretListResponse) AsUnion() ScriptSecretListResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by [ScriptSecretListResponseWorkersBindingKindSecretText] or
// [ScriptSecretListResponseWorkersBindingKindSecretKey].
type ScriptSecretListResponseUnion interface {
	implementsScriptSecretListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSecretListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretListResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretListResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type ScriptSecretListResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretListResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptSecretListResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptSecretListResponseWorkersBindingKindSecretTextJSON contains the JSON
// metadata for the struct [ScriptSecretListResponseWorkersBindingKindSecretText]
type scriptSecretListResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretListResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretListResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretListResponseWorkersBindingKindSecretText) implementsScriptSecretListResponse() {}

// The kind of resource that the binding provides.
type ScriptSecretListResponseWorkersBindingKindSecretTextType string

const (
	ScriptSecretListResponseWorkersBindingKindSecretTextTypeSecretText ScriptSecretListResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptSecretListResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretListResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretListResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretListResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptSecretListResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptSecretListResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptSecretListResponseWorkersBindingKindSecretKeyJSON contains the JSON
// metadata for the struct [ScriptSecretListResponseWorkersBindingKindSecretKey]
type scriptSecretListResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretListResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretListResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretListResponseWorkersBindingKindSecretKey) implementsScriptSecretListResponse() {}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretListResponseWorkersBindingKindSecretKeyFormat string

const (
	ScriptSecretListResponseWorkersBindingKindSecretKeyFormatRaw   ScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "raw"
	ScriptSecretListResponseWorkersBindingKindSecretKeyFormatPkcs8 ScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptSecretListResponseWorkersBindingKindSecretKeyFormatSpki  ScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "spki"
	ScriptSecretListResponseWorkersBindingKindSecretKeyFormatJwk   ScriptSecretListResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptSecretListResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseWorkersBindingKindSecretKeyFormatRaw, ScriptSecretListResponseWorkersBindingKindSecretKeyFormatPkcs8, ScriptSecretListResponseWorkersBindingKindSecretKeyFormatSpki, ScriptSecretListResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretListResponseWorkersBindingKindSecretKeyType string

const (
	ScriptSecretListResponseWorkersBindingKindSecretKeyTypeSecretKey ScriptSecretListResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptSecretListResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptSecretListResponseWorkersBindingKindSecretKeyUsage string

const (
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageEncrypt    ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDecrypt    ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageSign       ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "sign"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageVerify     ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "verify"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveKey  ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveBits ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageWrapKey    ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptSecretListResponseWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptSecretListResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptSecretListResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseWorkersBindingKindSecretKeyUsageEncrypt, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDecrypt, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageSign, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageVerify, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveKey, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageDeriveBits, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageWrapKey, ScriptSecretListResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretListResponseType string

const (
	ScriptSecretListResponseTypeSecretText ScriptSecretListResponseType = "secret_text"
	ScriptSecretListResponseTypeSecretKey  ScriptSecretListResponseType = "secret_key"
)

func (r ScriptSecretListResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseTypeSecretText, ScriptSecretListResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretListResponseFormat string

const (
	ScriptSecretListResponseFormatRaw   ScriptSecretListResponseFormat = "raw"
	ScriptSecretListResponseFormatPkcs8 ScriptSecretListResponseFormat = "pkcs8"
	ScriptSecretListResponseFormatSpki  ScriptSecretListResponseFormat = "spki"
	ScriptSecretListResponseFormatJwk   ScriptSecretListResponseFormat = "jwk"
)

func (r ScriptSecretListResponseFormat) IsKnown() bool {
	switch r {
	case ScriptSecretListResponseFormatRaw, ScriptSecretListResponseFormatPkcs8, ScriptSecretListResponseFormatSpki, ScriptSecretListResponseFormatJwk:
		return true
	}
	return false
}

type ScriptSecretDeleteResponse = interface{}

// A secret value accessible through a binding.
type ScriptSecretGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretGetResponseType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretGetResponseFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage].
	Usages interface{}                 `json:"usages"`
	JSON   scriptSecretGetResponseJSON `json:"-"`
	union  ScriptSecretGetResponseUnion
}

// scriptSecretGetResponseJSON contains the JSON metadata for the struct
// [ScriptSecretGetResponse]
type scriptSecretGetResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptSecretGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptSecretGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptSecretGetResponseWorkersBindingKindSecretText],
// [ScriptSecretGetResponseWorkersBindingKindSecretKey].
func (r ScriptSecretGetResponse) AsUnion() ScriptSecretGetResponseUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by [ScriptSecretGetResponseWorkersBindingKindSecretText] or
// [ScriptSecretGetResponseWorkersBindingKindSecretKey].
type ScriptSecretGetResponseUnion interface {
	implementsScriptSecretGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSecretGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptSecretGetResponseWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type ScriptSecretGetResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretGetResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptSecretGetResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptSecretGetResponseWorkersBindingKindSecretTextJSON contains the JSON
// metadata for the struct [ScriptSecretGetResponseWorkersBindingKindSecretText]
type scriptSecretGetResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretGetResponseWorkersBindingKindSecretText) implementsScriptSecretGetResponse() {}

// The kind of resource that the binding provides.
type ScriptSecretGetResponseWorkersBindingKindSecretTextType string

const (
	ScriptSecretGetResponseWorkersBindingKindSecretTextTypeSecretText ScriptSecretGetResponseWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptSecretGetResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretGetResponseWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptSecretGetResponseWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptSecretGetResponseWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptSecretGetResponseWorkersBindingKindSecretKeyJSON contains the JSON
// metadata for the struct [ScriptSecretGetResponseWorkersBindingKindSecretKey]
type scriptSecretGetResponseWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSecretGetResponseWorkersBindingKindSecretKey) implementsScriptSecretGetResponse() {}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat string

const (
	ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatRaw   ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "raw"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatPkcs8 ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatSpki  ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "spki"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatJwk   ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptSecretGetResponseWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatRaw, ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatPkcs8, ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatSpki, ScriptSecretGetResponseWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretGetResponseWorkersBindingKindSecretKeyType string

const (
	ScriptSecretGetResponseWorkersBindingKindSecretKeyTypeSecretKey ScriptSecretGetResponseWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptSecretGetResponseWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage string

const (
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageEncrypt    ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDecrypt    ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageSign       ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "sign"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageVerify     ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "verify"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveKey  ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveBits ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageWrapKey    ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptSecretGetResponseWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageEncrypt, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDecrypt, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageSign, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageVerify, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveKey, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageDeriveBits, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageWrapKey, ScriptSecretGetResponseWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretGetResponseType string

const (
	ScriptSecretGetResponseTypeSecretText ScriptSecretGetResponseType = "secret_text"
	ScriptSecretGetResponseTypeSecretKey  ScriptSecretGetResponseType = "secret_key"
)

func (r ScriptSecretGetResponseType) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseTypeSecretText, ScriptSecretGetResponseTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretGetResponseFormat string

const (
	ScriptSecretGetResponseFormatRaw   ScriptSecretGetResponseFormat = "raw"
	ScriptSecretGetResponseFormatPkcs8 ScriptSecretGetResponseFormat = "pkcs8"
	ScriptSecretGetResponseFormatSpki  ScriptSecretGetResponseFormat = "spki"
	ScriptSecretGetResponseFormatJwk   ScriptSecretGetResponseFormat = "jwk"
)

func (r ScriptSecretGetResponseFormat) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseFormatRaw, ScriptSecretGetResponseFormatPkcs8, ScriptSecretGetResponseFormatSpki, ScriptSecretGetResponseFormatJwk:
		return true
	}
	return false
}

type ScriptSecretUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A secret value accessible through a binding.
	Body ScriptSecretUpdateParamsBodyUnion `json:"body,required"`
}

func (r ScriptSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// A secret value accessible through a binding.
type ScriptSecretUpdateParamsBody struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type      param.Field[ScriptSecretUpdateParamsBodyType] `json:"type,required"`
	Algorithm param.Field[interface{}]                      `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptSecretUpdateParamsBodyFormat] `json:"format"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// The secret value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
}

func (r ScriptSecretUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSecretUpdateParamsBody) implementsScriptSecretUpdateParamsBodyUnion() {}

// A secret value accessible through a binding.
//
// Satisfied by [workers.ScriptSecretUpdateParamsBodyWorkersBindingKindSecretText],
// [workers.ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey],
// [ScriptSecretUpdateParamsBody].
type ScriptSecretUpdateParamsBodyUnion interface {
	implementsScriptSecretUpdateParamsBodyUnion()
}

type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretText) implementsScriptSecretUpdateParamsBodyUnion() {
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType string

const (
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextTypeSecretText ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKey) implementsScriptSecretUpdateParamsBodyUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat string

const (
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatRaw   ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "raw"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatPkcs8 ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatSpki  ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "spki"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatJwk   ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatRaw, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatPkcs8, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatSpki, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType string

const (
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyTypeSecretKey ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage string

const (
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageEncrypt    ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDecrypt    ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageSign       ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "sign"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageVerify     ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "verify"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveKey  ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveBits ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageWrapKey    ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageEncrypt, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDecrypt, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageSign, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageVerify, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveKey, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageDeriveBits, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageWrapKey, ScriptSecretUpdateParamsBodyWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptSecretUpdateParamsBodyType string

const (
	ScriptSecretUpdateParamsBodyTypeSecretText ScriptSecretUpdateParamsBodyType = "secret_text"
	ScriptSecretUpdateParamsBodyTypeSecretKey  ScriptSecretUpdateParamsBodyType = "secret_key"
)

func (r ScriptSecretUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyTypeSecretText, ScriptSecretUpdateParamsBodyTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptSecretUpdateParamsBodyFormat string

const (
	ScriptSecretUpdateParamsBodyFormatRaw   ScriptSecretUpdateParamsBodyFormat = "raw"
	ScriptSecretUpdateParamsBodyFormatPkcs8 ScriptSecretUpdateParamsBodyFormat = "pkcs8"
	ScriptSecretUpdateParamsBodyFormatSpki  ScriptSecretUpdateParamsBodyFormat = "spki"
	ScriptSecretUpdateParamsBodyFormatJwk   ScriptSecretUpdateParamsBodyFormat = "jwk"
)

func (r ScriptSecretUpdateParamsBodyFormat) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateParamsBodyFormatRaw, ScriptSecretUpdateParamsBodyFormatPkcs8, ScriptSecretUpdateParamsBodyFormatSpki, ScriptSecretUpdateParamsBodyFormatJwk:
		return true
	}
	return false
}

type ScriptSecretUpdateResponseEnvelope struct {
	Errors   []ScriptSecretUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSecretUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretUpdateResponseEnvelopeSuccess `json:"success,required"`
	// A secret value accessible through a binding.
	Result ScriptSecretUpdateResponse             `json:"result"`
	JSON   scriptSecretUpdateResponseEnvelopeJSON `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretUpdateResponseEnvelope]
type scriptSecretUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptSecretUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSecretUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSecretUpdateResponseEnvelopeErrors]
type scriptSecretUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptSecretUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptSecretUpdateResponseEnvelopeErrorsSource]
type scriptSecretUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptSecretUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSecretUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptSecretUpdateResponseEnvelopeMessages]
type scriptSecretUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptSecretUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSecretUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSecretUpdateResponseEnvelopeMessagesSource]
type scriptSecretUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretUpdateResponseEnvelopeSuccess bool

const (
	ScriptSecretUpdateResponseEnvelopeSuccessTrue ScriptSecretUpdateResponseEnvelopeSuccess = true
)

func (r ScriptSecretUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSecretListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretDeleteResponseEnvelope struct {
	Errors   []ScriptSecretDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSecretDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptSecretDeleteResponse                `json:"result,nullable"`
	JSON    scriptSecretDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretDeleteResponseEnvelope]
type scriptSecretDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretDeleteResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptSecretDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSecretDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSecretDeleteResponseEnvelopeErrors]
type scriptSecretDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptSecretDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptSecretDeleteResponseEnvelopeErrorsSource]
type scriptSecretDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretDeleteResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptSecretDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSecretDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptSecretDeleteResponseEnvelopeMessages]
type scriptSecretDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptSecretDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSecretDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSecretDeleteResponseEnvelopeMessagesSource]
type scriptSecretDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretDeleteResponseEnvelopeSuccess bool

const (
	ScriptSecretDeleteResponseEnvelopeSuccessTrue ScriptSecretDeleteResponseEnvelopeSuccess = true
)

func (r ScriptSecretDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSecretGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSecretGetResponseEnvelope struct {
	Errors   []ScriptSecretGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSecretGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptSecretGetResponseEnvelopeSuccess `json:"success,required"`
	// A secret value accessible through a binding.
	Result ScriptSecretGetResponse             `json:"result"`
	JSON   scriptSecretGetResponseEnvelopeJSON `json:"-"`
}

// scriptSecretGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSecretGetResponseEnvelope]
type scriptSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ScriptSecretGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSecretGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSecretGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSecretGetResponseEnvelopeErrors]
type scriptSecretGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    scriptSecretGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSecretGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptSecretGetResponseEnvelopeErrorsSource]
type scriptSecretGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ScriptSecretGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSecretGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSecretGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptSecretGetResponseEnvelopeMessages]
type scriptSecretGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSecretGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    scriptSecretGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSecretGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptSecretGetResponseEnvelopeMessagesSource]
type scriptSecretGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSecretGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSecretGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSecretGetResponseEnvelopeSuccess bool

const (
	ScriptSecretGetResponseEnvelopeSuccessTrue ScriptSecretGetResponseEnvelopeSuccess = true
)

func (r ScriptSecretGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSecretGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
