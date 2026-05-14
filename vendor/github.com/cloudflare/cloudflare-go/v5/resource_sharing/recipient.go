// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_sharing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// RecipientService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientService] method instead.
type RecipientService struct {
	Options []option.RequestOption
}

// NewRecipientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientService(opts ...option.RequestOption) (r *RecipientService) {
	r = &RecipientService{}
	r.Options = opts
	return
}

// Create a new share recipient
func (r *RecipientService) New(ctx context.Context, shareID string, params RecipientNewParams, opts ...option.RequestOption) (res *RecipientNewResponse, err error) {
	var env RecipientNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.PathAccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", params.PathAccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List share recipients by share ID.
func (r *RecipientService) List(ctx context.Context, shareID string, params RecipientListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RecipientListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", params.AccountID, shareID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List share recipients by share ID.
func (r *RecipientService) ListAutoPaging(ctx context.Context, shareID string, params RecipientListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RecipientListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, shareID, params, opts...))
}

// Deletion is not immediate, an updated share recipient object with a new status
// will be returned.
func (r *RecipientService) Delete(ctx context.Context, shareID string, recipientID string, body RecipientDeleteParams, opts ...option.RequestOption) (res *RecipientDeleteResponse, err error) {
	var env RecipientDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if recipientID == "" {
		err = errors.New("missing required recipient_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", body.AccountID, shareID, recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get share recipient by ID.
func (r *RecipientService) Get(ctx context.Context, shareID string, recipientID string, query RecipientGetParams, opts ...option.RequestOption) (res *RecipientGetResponse, err error) {
	var env RecipientGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if recipientID == "" {
		err = errors.New("missing required recipient_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", query.AccountID, shareID, recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RecipientNewResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus RecipientNewResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                   `json:"status_message,required"`
	JSON          recipientNewResponseJSON `json:"-"`
}

// recipientNewResponseJSON contains the JSON metadata for the struct
// [RecipientNewResponse]
type recipientNewResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecipientNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientNewResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type RecipientNewResponseAssociationStatus string

const (
	RecipientNewResponseAssociationStatusAssociating    RecipientNewResponseAssociationStatus = "associating"
	RecipientNewResponseAssociationStatusAssociated     RecipientNewResponseAssociationStatus = "associated"
	RecipientNewResponseAssociationStatusDisassociating RecipientNewResponseAssociationStatus = "disassociating"
	RecipientNewResponseAssociationStatusDisassociated  RecipientNewResponseAssociationStatus = "disassociated"
)

func (r RecipientNewResponseAssociationStatus) IsKnown() bool {
	switch r {
	case RecipientNewResponseAssociationStatusAssociating, RecipientNewResponseAssociationStatusAssociated, RecipientNewResponseAssociationStatusDisassociating, RecipientNewResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type RecipientListResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus RecipientListResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                    `json:"status_message,required"`
	JSON          recipientListResponseJSON `json:"-"`
}

// recipientListResponseJSON contains the JSON metadata for the struct
// [RecipientListResponse]
type recipientListResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecipientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientListResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type RecipientListResponseAssociationStatus string

const (
	RecipientListResponseAssociationStatusAssociating    RecipientListResponseAssociationStatus = "associating"
	RecipientListResponseAssociationStatusAssociated     RecipientListResponseAssociationStatus = "associated"
	RecipientListResponseAssociationStatusDisassociating RecipientListResponseAssociationStatus = "disassociating"
	RecipientListResponseAssociationStatusDisassociated  RecipientListResponseAssociationStatus = "disassociated"
)

func (r RecipientListResponseAssociationStatus) IsKnown() bool {
	switch r {
	case RecipientListResponseAssociationStatusAssociating, RecipientListResponseAssociationStatusAssociated, RecipientListResponseAssociationStatusDisassociating, RecipientListResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type RecipientDeleteResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus RecipientDeleteResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                      `json:"status_message,required"`
	JSON          recipientDeleteResponseJSON `json:"-"`
}

// recipientDeleteResponseJSON contains the JSON metadata for the struct
// [RecipientDeleteResponse]
type recipientDeleteResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecipientDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type RecipientDeleteResponseAssociationStatus string

const (
	RecipientDeleteResponseAssociationStatusAssociating    RecipientDeleteResponseAssociationStatus = "associating"
	RecipientDeleteResponseAssociationStatusAssociated     RecipientDeleteResponseAssociationStatus = "associated"
	RecipientDeleteResponseAssociationStatusDisassociating RecipientDeleteResponseAssociationStatus = "disassociating"
	RecipientDeleteResponseAssociationStatusDisassociated  RecipientDeleteResponseAssociationStatus = "disassociated"
)

func (r RecipientDeleteResponseAssociationStatus) IsKnown() bool {
	switch r {
	case RecipientDeleteResponseAssociationStatusAssociating, RecipientDeleteResponseAssociationStatusAssociated, RecipientDeleteResponseAssociationStatusDisassociating, RecipientDeleteResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type RecipientGetResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus RecipientGetResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                   `json:"status_message,required"`
	JSON          recipientGetResponseJSON `json:"-"`
}

// recipientGetResponseJSON contains the JSON metadata for the struct
// [RecipientGetResponse]
type recipientGetResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecipientGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientGetResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type RecipientGetResponseAssociationStatus string

const (
	RecipientGetResponseAssociationStatusAssociating    RecipientGetResponseAssociationStatus = "associating"
	RecipientGetResponseAssociationStatusAssociated     RecipientGetResponseAssociationStatus = "associated"
	RecipientGetResponseAssociationStatusDisassociating RecipientGetResponseAssociationStatus = "disassociating"
	RecipientGetResponseAssociationStatusDisassociated  RecipientGetResponseAssociationStatus = "disassociated"
)

func (r RecipientGetResponseAssociationStatus) IsKnown() bool {
	switch r {
	case RecipientGetResponseAssociationStatusAssociating, RecipientGetResponseAssociationStatusAssociated, RecipientGetResponseAssociationStatusDisassociating, RecipientGetResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type RecipientNewParams struct {
	// Account identifier.
	PathAccountID param.Field[string] `path:"account_id,required"`
	// Account identifier.
	BodyAccountID param.Field[string] `json:"account_id"`
	// Organization identifier.
	OrganizationID param.Field[string] `json:"organization_id"`
}

func (r RecipientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecipientNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                             `json:"success,required"`
	Result  RecipientNewResponse             `json:"result"`
	JSON    recipientNewResponseEnvelopeJSON `json:"-"`
}

// recipientNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecipientNewResponseEnvelope]
type recipientNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecipientNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RecipientListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [RecipientListParams]'s query parameters as `url.Values`.
func (r RecipientListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RecipientDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RecipientDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                `json:"success,required"`
	Result  RecipientDeleteResponse             `json:"result"`
	JSON    recipientDeleteResponseEnvelopeJSON `json:"-"`
}

// recipientDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecipientDeleteResponseEnvelope]
type recipientDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecipientDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RecipientGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RecipientGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                             `json:"success,required"`
	Result  RecipientGetResponse             `json:"result"`
	JSON    recipientGetResponseEnvelopeJSON `json:"-"`
}

// recipientGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecipientGetResponseEnvelope]
type recipientGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecipientGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
