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

// PublishService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPublishService] method instead.
type PublishService struct {
	Options []option.RequestOption
}

// NewPublishService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPublishService(opts ...option.RequestOption) (r *PublishService) {
	r = &PublishService{}
	r.Options = opts
	return
}

// Publish current Zaraz preview configuration for a zone.
func (r *PublishService) New(ctx context.Context, params PublishNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env PublishNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/publish", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PublishNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Zaraz configuration description.
	Body string `json:"body"`
}

func (r PublishNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PublishNewResponseEnvelope struct {
	Errors   []PublishNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PublishNewResponseEnvelopeMessages `json:"messages,required"`
	Result   string                               `json:"result,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	JSON    publishNewResponseEnvelopeJSON `json:"-"`
}

// publishNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PublishNewResponseEnvelope]
type publishNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PublishNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publishNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PublishNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           PublishNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             publishNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// publishNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PublishNewResponseEnvelopeErrors]
type publishNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PublishNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publishNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PublishNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    publishNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// publishNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PublishNewResponseEnvelopeErrorsSource]
type publishNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PublishNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publishNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PublishNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           PublishNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             publishNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// publishNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PublishNewResponseEnvelopeMessages]
type publishNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PublishNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publishNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PublishNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    publishNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// publishNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PublishNewResponseEnvelopeMessagesSource]
type publishNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PublishNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publishNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
