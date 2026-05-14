// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// FinetuneService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinetuneService] method instead.
type FinetuneService struct {
	Options []option.RequestOption
	Assets  *FinetuneAssetService
	Public  *FinetunePublicService
}

// NewFinetuneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFinetuneService(opts ...option.RequestOption) (r *FinetuneService) {
	r = &FinetuneService{}
	r.Options = opts
	r.Assets = NewFinetuneAssetService(opts...)
	r.Public = NewFinetunePublicService(opts...)
	return
}

// Create a new Finetune
func (r *FinetuneService) New(ctx context.Context, params FinetuneNewParams, opts ...option.RequestOption) (res *FinetuneNewResponse, err error) {
	var env FinetuneNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Finetunes
func (r *FinetuneService) List(ctx context.Context, query FinetuneListParams, opts ...option.RequestOption) (res *FinetuneListResponse, err error) {
	var env FinetuneListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FinetuneNewResponse struct {
	ID          string                  `json:"id,required" format:"uuid"`
	CreatedAt   time.Time               `json:"created_at,required" format:"date-time"`
	Model       string                  `json:"model,required"`
	ModifiedAt  time.Time               `json:"modified_at,required" format:"date-time"`
	Name        string                  `json:"name,required"`
	Public      bool                    `json:"public,required"`
	Description string                  `json:"description"`
	JSON        finetuneNewResponseJSON `json:"-"`
}

// finetuneNewResponseJSON contains the JSON metadata for the struct
// [FinetuneNewResponse]
type finetuneNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneNewResponseJSON) RawJSON() string {
	return r.raw
}

type FinetuneListResponse struct {
	ID          string                   `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                `json:"created_at,required" format:"date-time"`
	Model       string                   `json:"model,required"`
	ModifiedAt  time.Time                `json:"modified_at,required" format:"date-time"`
	Name        string                   `json:"name,required"`
	Description string                   `json:"description"`
	JSON        finetuneListResponseJSON `json:"-"`
}

// finetuneListResponseJSON contains the JSON metadata for the struct
// [FinetuneListResponse]
type finetuneListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneListResponseJSON) RawJSON() string {
	return r.raw
}

type FinetuneNewParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Model       param.Field[string] `json:"model,required"`
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Public      param.Field[bool]   `json:"public"`
}

func (r FinetuneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FinetuneNewResponseEnvelope struct {
	Result  FinetuneNewResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    finetuneNewResponseEnvelopeJSON `json:"-"`
}

// finetuneNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FinetuneNewResponseEnvelope]
type finetuneNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FinetuneListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type FinetuneListResponseEnvelope struct {
	Result  FinetuneListResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    finetuneListResponseEnvelopeJSON `json:"-"`
}

// finetuneListResponseEnvelopeJSON contains the JSON metadata for the struct
// [FinetuneListResponseEnvelope]
type finetuneListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
