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

// DefaultService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDefaultService] method instead.
type DefaultService struct {
	Options []option.RequestOption
}

// NewDefaultService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDefaultService(opts ...option.RequestOption) (r *DefaultService) {
	r = &DefaultService{}
	r.Options = opts
	return
}

// Gets default Zaraz configuration for a zone.
func (r *DefaultService) Get(ctx context.Context, query DefaultGetParams, opts ...option.RequestOption) (res *Configuration, err error) {
	var env DefaultGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/default", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DefaultGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DefaultGetResponseEnvelope struct {
	Errors   []DefaultGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DefaultGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result Configuration `json:"result,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	JSON    defaultGetResponseEnvelopeJSON `json:"-"`
}

// defaultGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DefaultGetResponseEnvelope]
type defaultGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DefaultGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           DefaultGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             defaultGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// defaultGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DefaultGetResponseEnvelopeErrors]
type defaultGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DefaultGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DefaultGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    defaultGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// defaultGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DefaultGetResponseEnvelopeErrorsSource]
type defaultGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DefaultGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DefaultGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             defaultGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// defaultGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DefaultGetResponseEnvelopeMessages]
type defaultGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DefaultGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DefaultGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    defaultGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// defaultGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DefaultGetResponseEnvelopeMessagesSource]
type defaultGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
