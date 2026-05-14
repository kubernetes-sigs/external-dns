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

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/workers"
)

// DispatchNamespaceScriptContentService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptContentService] method instead.
type DispatchNamespaceScriptContentService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptContentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentService) {
	r = &DispatchNamespaceScriptContentService{}
	r.Options = opts
	return
}

// Put script content for a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptContentService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptContentUpdateParams, opts ...option.RequestOption) (res *workers.Script, err error) {
	var env DispatchNamespaceScriptContentUpdateResponseEnvelope
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
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch script content from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptContentService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DispatchNamespaceScriptContentUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[workers.WorkerMetadataParam] `json:"metadata,required"`
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

func (r DispatchNamespaceScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type DispatchNamespaceScriptContentUpdateResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   workers.Script                                                 `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptContentUpdateResponseEnvelope]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           DispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages struct {
	Code             int64                                                              `json:"code,required"`
	Message          string                                                             `json:"message,required"`
	DocumentationURL string                                                             `json:"documentation_url"`
	Source           DispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                                 `json:"pointer"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
