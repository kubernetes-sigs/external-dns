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

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ScriptContentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptContentService] method instead.
type ScriptContentService struct {
	Options []option.RequestOption
}

// NewScriptContentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptContentService(opts ...option.RequestOption) (r *ScriptContentService) {
	r = &ScriptContentService{}
	r.Options = opts
	return
}

// Put script content without touching config or metadata.
func (r *ScriptContentService) Update(ctx context.Context, scriptName string, params ScriptContentUpdateParams, opts ...option.RequestOption) (res *Script, err error) {
	var env ScriptContentUpdateResponseEnvelope
	if params.CfWorkerBodyPart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-BODY-PART", fmt.Sprintf("%s", params.CfWorkerBodyPart)))
	}
	if params.CfWorkerMainModulePart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-MAIN-MODULE-PART", fmt.Sprintf("%s", params.CfWorkerMainModulePart)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch script content only.
func (r *ScriptContentService) Get(ctx context.Context, scriptName string, query ScriptContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content/v2", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScriptContentUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptContentUpdateParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files                  param.Field[[]io.Reader] `json:"files" format:"binary"`
	CfWorkerBodyPart       param.Field[string]      `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]      `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r ScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ScriptContentUpdateParamsMetadata struct {
	// Name of the uploaded file that contains the Worker script (e.g. the file adding
	// a listener to the `fetch` event). Indicates a `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r ScriptContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptContentUpdateResponseEnvelope struct {
	Errors   []ScriptContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   Script                                        `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptContentUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptContentUpdateResponseEnvelope]
type scriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptContentUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptContentUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptContentUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptContentUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptContentUpdateResponseEnvelopeErrors]
type scriptContentUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptContentUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptContentUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptContentUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptContentUpdateResponseEnvelopeErrorsSource]
type scriptContentUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptContentUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptContentUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptContentUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptContentUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptContentUpdateResponseEnvelopeMessages]
type scriptContentUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptContentUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptContentUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptContentUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptContentUpdateResponseEnvelopeMessagesSource]
type scriptContentUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptContentUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptContentUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptContentUpdateResponseEnvelopeSuccess bool

const (
	ScriptContentUpdateResponseEnvelopeSuccessTrue ScriptContentUpdateResponseEnvelopeSuccess = true
)

func (r ScriptContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptContentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
