// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

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

// ZarazService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZarazService] method instead.
type ZarazService struct {
	Options  []option.RequestOption
	Config   *ConfigService
	Default  *DefaultService
	Export   *ExportService
	History  *HistoryService
	Publish  *PublishService
	Workflow *WorkflowService
}

// NewZarazService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZarazService(opts ...option.RequestOption) (r *ZarazService) {
	r = &ZarazService{}
	r.Options = opts
	r.Config = NewConfigService(opts...)
	r.Default = NewDefaultService(opts...)
	r.Export = NewExportService(opts...)
	r.History = NewHistoryService(opts...)
	r.Publish = NewPublishService(opts...)
	r.Workflow = NewWorkflowService(opts...)
	return
}

// Updates Zaraz workflow for a zone.
func (r *ZarazService) Update(ctx context.Context, params ZarazUpdateParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env ZarazUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ButtonTextTranslation struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string         `json:"reject_all,required"`
	JSON      buttonTextTranslationJSON `json:"-"`
}

// buttonTextTranslationJSON contains the JSON metadata for the struct
// [ButtonTextTranslation]
type buttonTextTranslationJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ButtonTextTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buttonTextTranslationJSON) RawJSON() string {
	return r.raw
}

type ButtonTextTranslationParam struct {
	// Object where keys are language codes
	AcceptAll param.Field[map[string]string] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[map[string]string] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[map[string]string] `json:"reject_all,required"`
}

func (r ButtonTextTranslationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string     `json:"firingTriggers,required"`
	JSON           neoEventJSON `json:"-"`
}

// neoEventJSON contains the JSON metadata for the struct [NeoEvent]
type neoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r neoEventJSON) RawJSON() string {
	return r.raw
}

type NeoEventParam struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r NeoEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Zaraz workflow
	Workflow Workflow `json:"workflow,required"`
}

func (r ZarazUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Workflow)
}

type ZarazUpdateResponseEnvelope struct {
	Errors   []ZarazUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result Workflow `json:"result,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	JSON    zarazUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazUpdateResponseEnvelope]
type zarazUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZarazUpdateResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           ZarazUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             zarazUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// zarazUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZarazUpdateResponseEnvelopeErrors]
type zarazUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZarazUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    zarazUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ZarazUpdateResponseEnvelopeErrorsSource]
type zarazUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZarazUpdateResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ZarazUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             zarazUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// zarazUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazUpdateResponseEnvelopeMessages]
type zarazUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZarazUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    zarazUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ZarazUpdateResponseEnvelopeMessagesSource]
type zarazUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
