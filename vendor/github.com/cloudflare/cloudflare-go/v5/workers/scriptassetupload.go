// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ScriptAssetUploadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptAssetUploadService] method instead.
type ScriptAssetUploadService struct {
	Options []option.RequestOption
}

// NewScriptAssetUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptAssetUploadService(opts ...option.RequestOption) (r *ScriptAssetUploadService) {
	r = &ScriptAssetUploadService{}
	r.Options = opts
	return
}

// Start uploading a collection of assets for use in a Worker version. To learn
// more about the direct uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/.
func (r *ScriptAssetUploadService) New(ctx context.Context, scriptName string, params ScriptAssetUploadNewParams, opts ...option.RequestOption) (res *ScriptAssetUploadNewResponse, err error) {
	var env ScriptAssetUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/assets-upload-session", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptAssetUploadNewResponse struct {
	// The requests to make to upload assets.
	Buckets [][]string `json:"buckets"`
	// A JWT to use as authentication for uploading assets.
	JWT  string                           `json:"jwt"`
	JSON scriptAssetUploadNewResponseJSON `json:"-"`
}

// scriptAssetUploadNewResponseJSON contains the JSON metadata for the struct
// [ScriptAssetUploadNewResponse]
type scriptAssetUploadNewResponseJSON struct {
	Buckets     apijson.Field
	JWT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptAssetUploadNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// A manifest ([path]: {hash, size}) map of files to upload. As an example,
	// `/blog/hello-world.html` would be a valid path key.
	Manifest param.Field[map[string]ScriptAssetUploadNewParamsManifest] `json:"manifest,required"`
}

func (r ScriptAssetUploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptAssetUploadNewParamsManifest struct {
	// The hash of the file.
	Hash param.Field[string] `json:"hash,required"`
	// The size of the file in bytes.
	Size param.Field[int64] `json:"size,required"`
}

func (r ScriptAssetUploadNewParamsManifest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptAssetUploadNewResponseEnvelope struct {
	Errors   []ScriptAssetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptAssetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptAssetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptAssetUploadNewResponse                `json:"result"`
	JSON    scriptAssetUploadNewResponseEnvelopeJSON    `json:"-"`
}

// scriptAssetUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptAssetUploadNewResponseEnvelope]
type scriptAssetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptAssetUploadNewResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptAssetUploadNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptAssetUploadNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptAssetUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptAssetUploadNewResponseEnvelopeErrors]
type scriptAssetUploadNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptAssetUploadNewResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptAssetUploadNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptAssetUploadNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptAssetUploadNewResponseEnvelopeErrorsSource]
type scriptAssetUploadNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptAssetUploadNewResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ScriptAssetUploadNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptAssetUploadNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptAssetUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptAssetUploadNewResponseEnvelopeMessages]
type scriptAssetUploadNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptAssetUploadNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    scriptAssetUploadNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptAssetUploadNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ScriptAssetUploadNewResponseEnvelopeMessagesSource]
type scriptAssetUploadNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptAssetUploadNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptAssetUploadNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptAssetUploadNewResponseEnvelopeSuccess bool

const (
	ScriptAssetUploadNewResponseEnvelopeSuccessTrue ScriptAssetUploadNewResponseEnvelopeSuccess = true
)

func (r ScriptAssetUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptAssetUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
