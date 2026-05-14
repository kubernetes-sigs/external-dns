// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// SettingDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingDomainService] method instead.
type SettingDomainService struct {
	Options []option.RequestOption
}

// NewSettingDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingDomainService(opts ...option.RequestOption) (r *SettingDomainService) {
	r = &SettingDomainService{}
	r.Options = opts
	return
}

// Lists, searches, and sorts an account’s email domains.
func (r *SettingDomainService) List(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", params.AccountID)
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

// Lists, searches, and sorts an account’s email domains.
func (r *SettingDomainService) ListAutoPaging(ctx context.Context, params SettingDomainListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingDomainListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Unprotect an email domain
func (r *SettingDomainService) Delete(ctx context.Context, domainID int64, body SettingDomainDeleteParams, opts ...option.RequestOption) (res *SettingDomainDeleteResponse, err error) {
	var env SettingDomainDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Unprotect multiple email domains
func (r *SettingDomainService) BulkDelete(ctx context.Context, body SettingDomainBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[SettingDomainBulkDeleteResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", body.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Unprotect multiple email domains
func (r *SettingDomainService) BulkDeleteAutoPaging(ctx context.Context, body SettingDomainBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SettingDomainBulkDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, body, opts...))
}

// Update an email domain
func (r *SettingDomainService) Edit(ctx context.Context, domainID int64, params SettingDomainEditParams, opts ...option.RequestOption) (res *SettingDomainEditResponse, err error) {
	var env SettingDomainEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", params.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an email domain
func (r *SettingDomainService) Get(ctx context.Context, domainID int64, query SettingDomainGetParams, opts ...option.RequestOption) (res *SettingDomainGetResponse, err error) {
	var env SettingDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingDomainListResponse struct {
	// The unique identifier for the domain.
	ID                   int64                                          `json:"id,required"`
	AllowedDeliveryModes []SettingDomainListResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                      `json:"created_at,required" format:"date-time"`
	Domain               string                                         `json:"domain,required"`
	DropDispositions     []SettingDomainListResponseDropDisposition     `json:"drop_dispositions,required"`
	IPRestrictions       []string                                       `json:"ip_restrictions,required"`
	LastModified         time.Time                                      `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                          `json:"lookback_hops,required"`
	Transport            string                                         `json:"transport,required"`
	Authorization        SettingDomainListResponseAuthorization         `json:"authorization,nullable"`
	EmailsProcessed      SettingDomainListResponseEmailsProcessed       `json:"emails_processed,nullable"`
	Folder               SettingDomainListResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainListResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                         `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                         `json:"o365_tenant_id,nullable"`
	RequireTLSInbound    bool                                           `json:"require_tls_inbound,nullable"`
	RequireTLSOutbound   bool                                           `json:"require_tls_outbound,nullable"`
	JSON                 settingDomainListResponseJSON                  `json:"-"`
}

// settingDomainListResponseJSON contains the JSON metadata for the struct
// [SettingDomainListResponse]
type settingDomainListResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseAllowedDeliveryMode string

const (
	SettingDomainListResponseAllowedDeliveryModeDirect    SettingDomainListResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainListResponseAllowedDeliveryModeBcc       SettingDomainListResponseAllowedDeliveryMode = "BCC"
	SettingDomainListResponseAllowedDeliveryModeJournal   SettingDomainListResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainListResponseAllowedDeliveryModeAPI       SettingDomainListResponseAllowedDeliveryMode = "API"
	SettingDomainListResponseAllowedDeliveryModeRetroScan SettingDomainListResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListResponseAllowedDeliveryModeDirect, SettingDomainListResponseAllowedDeliveryModeBcc, SettingDomainListResponseAllowedDeliveryModeJournal, SettingDomainListResponseAllowedDeliveryModeAPI, SettingDomainListResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainListResponseDropDisposition string

const (
	SettingDomainListResponseDropDispositionMalicious    SettingDomainListResponseDropDisposition = "MALICIOUS"
	SettingDomainListResponseDropDispositionMaliciousBec SettingDomainListResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainListResponseDropDispositionSuspicious   SettingDomainListResponseDropDisposition = "SUSPICIOUS"
	SettingDomainListResponseDropDispositionSpoof        SettingDomainListResponseDropDisposition = "SPOOF"
	SettingDomainListResponseDropDispositionSpam         SettingDomainListResponseDropDisposition = "SPAM"
	SettingDomainListResponseDropDispositionBulk         SettingDomainListResponseDropDisposition = "BULK"
	SettingDomainListResponseDropDispositionEncrypted    SettingDomainListResponseDropDisposition = "ENCRYPTED"
	SettingDomainListResponseDropDispositionExternal     SettingDomainListResponseDropDisposition = "EXTERNAL"
	SettingDomainListResponseDropDispositionUnknown      SettingDomainListResponseDropDisposition = "UNKNOWN"
	SettingDomainListResponseDropDispositionNone         SettingDomainListResponseDropDisposition = "NONE"
)

func (r SettingDomainListResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainListResponseDropDispositionMalicious, SettingDomainListResponseDropDispositionMaliciousBec, SettingDomainListResponseDropDispositionSuspicious, SettingDomainListResponseDropDispositionSpoof, SettingDomainListResponseDropDispositionSpam, SettingDomainListResponseDropDispositionBulk, SettingDomainListResponseDropDispositionEncrypted, SettingDomainListResponseDropDispositionExternal, SettingDomainListResponseDropDispositionUnknown, SettingDomainListResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainListResponseAuthorization struct {
	Authorized    bool                                       `json:"authorized,required"`
	Timestamp     time.Time                                  `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                     `json:"status_message,nullable"`
	JSON          settingDomainListResponseAuthorizationJSON `json:"-"`
}

// settingDomainListResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainListResponseAuthorization]
type settingDomainListResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainListResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseEmailsProcessed struct {
	Timestamp                    time.Time                                    `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                        `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                        `json:"total_emails_processed_previous,required"`
	JSON                         settingDomainListResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainListResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainListResponseEmailsProcessed]
type settingDomainListResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainListResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainListResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainListResponseFolder string

const (
	SettingDomainListResponseFolderAllItems SettingDomainListResponseFolder = "AllItems"
	SettingDomainListResponseFolderInbox    SettingDomainListResponseFolder = "Inbox"
)

func (r SettingDomainListResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainListResponseFolderAllItems, SettingDomainListResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainListResponseInboxProvider string

const (
	SettingDomainListResponseInboxProviderMicrosoft SettingDomainListResponseInboxProvider = "Microsoft"
	SettingDomainListResponseInboxProviderGoogle    SettingDomainListResponseInboxProvider = "Google"
)

func (r SettingDomainListResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainListResponseInboxProviderMicrosoft, SettingDomainListResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainDeleteResponse struct {
	// The unique identifier for the domain.
	ID   int64                           `json:"id,required"`
	JSON settingDomainDeleteResponseJSON `json:"-"`
}

// settingDomainDeleteResponseJSON contains the JSON metadata for the struct
// [SettingDomainDeleteResponse]
type settingDomainDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBulkDeleteResponse struct {
	// The unique identifier for the domain.
	ID   int64                               `json:"id,required"`
	JSON settingDomainBulkDeleteResponseJSON `json:"-"`
}

// settingDomainBulkDeleteResponseJSON contains the JSON metadata for the struct
// [SettingDomainBulkDeleteResponse]
type settingDomainBulkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponse struct {
	// The unique identifier for the domain.
	ID                   int64                                          `json:"id,required"`
	AllowedDeliveryModes []SettingDomainEditResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                      `json:"created_at,required" format:"date-time"`
	Domain               string                                         `json:"domain,required"`
	DropDispositions     []SettingDomainEditResponseDropDisposition     `json:"drop_dispositions,required"`
	IPRestrictions       []string                                       `json:"ip_restrictions,required"`
	LastModified         time.Time                                      `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                          `json:"lookback_hops,required"`
	Transport            string                                         `json:"transport,required"`
	Authorization        SettingDomainEditResponseAuthorization         `json:"authorization,nullable"`
	EmailsProcessed      SettingDomainEditResponseEmailsProcessed       `json:"emails_processed,nullable"`
	Folder               SettingDomainEditResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainEditResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                         `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                         `json:"o365_tenant_id,nullable"`
	RequireTLSInbound    bool                                           `json:"require_tls_inbound,nullable"`
	RequireTLSOutbound   bool                                           `json:"require_tls_outbound,nullable"`
	JSON                 settingDomainEditResponseJSON                  `json:"-"`
}

// settingDomainEditResponseJSON contains the JSON metadata for the struct
// [SettingDomainEditResponse]
type settingDomainEditResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseAllowedDeliveryMode string

const (
	SettingDomainEditResponseAllowedDeliveryModeDirect    SettingDomainEditResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainEditResponseAllowedDeliveryModeBcc       SettingDomainEditResponseAllowedDeliveryMode = "BCC"
	SettingDomainEditResponseAllowedDeliveryModeJournal   SettingDomainEditResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainEditResponseAllowedDeliveryModeAPI       SettingDomainEditResponseAllowedDeliveryMode = "API"
	SettingDomainEditResponseAllowedDeliveryModeRetroScan SettingDomainEditResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainEditResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseAllowedDeliveryModeDirect, SettingDomainEditResponseAllowedDeliveryModeBcc, SettingDomainEditResponseAllowedDeliveryModeJournal, SettingDomainEditResponseAllowedDeliveryModeAPI, SettingDomainEditResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainEditResponseDropDisposition string

const (
	SettingDomainEditResponseDropDispositionMalicious    SettingDomainEditResponseDropDisposition = "MALICIOUS"
	SettingDomainEditResponseDropDispositionMaliciousBec SettingDomainEditResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainEditResponseDropDispositionSuspicious   SettingDomainEditResponseDropDisposition = "SUSPICIOUS"
	SettingDomainEditResponseDropDispositionSpoof        SettingDomainEditResponseDropDisposition = "SPOOF"
	SettingDomainEditResponseDropDispositionSpam         SettingDomainEditResponseDropDisposition = "SPAM"
	SettingDomainEditResponseDropDispositionBulk         SettingDomainEditResponseDropDisposition = "BULK"
	SettingDomainEditResponseDropDispositionEncrypted    SettingDomainEditResponseDropDisposition = "ENCRYPTED"
	SettingDomainEditResponseDropDispositionExternal     SettingDomainEditResponseDropDisposition = "EXTERNAL"
	SettingDomainEditResponseDropDispositionUnknown      SettingDomainEditResponseDropDisposition = "UNKNOWN"
	SettingDomainEditResponseDropDispositionNone         SettingDomainEditResponseDropDisposition = "NONE"
)

func (r SettingDomainEditResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseDropDispositionMalicious, SettingDomainEditResponseDropDispositionMaliciousBec, SettingDomainEditResponseDropDispositionSuspicious, SettingDomainEditResponseDropDispositionSpoof, SettingDomainEditResponseDropDispositionSpam, SettingDomainEditResponseDropDispositionBulk, SettingDomainEditResponseDropDispositionEncrypted, SettingDomainEditResponseDropDispositionExternal, SettingDomainEditResponseDropDispositionUnknown, SettingDomainEditResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainEditResponseAuthorization struct {
	Authorized    bool                                       `json:"authorized,required"`
	Timestamp     time.Time                                  `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                     `json:"status_message,nullable"`
	JSON          settingDomainEditResponseAuthorizationJSON `json:"-"`
}

// settingDomainEditResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseAuthorization]
type settingDomainEditResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainEditResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseEmailsProcessed struct {
	Timestamp                    time.Time                                    `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                        `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                        `json:"total_emails_processed_previous,required"`
	JSON                         settingDomainEditResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainEditResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainEditResponseEmailsProcessed]
type settingDomainEditResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainEditResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainEditResponseFolder string

const (
	SettingDomainEditResponseFolderAllItems SettingDomainEditResponseFolder = "AllItems"
	SettingDomainEditResponseFolderInbox    SettingDomainEditResponseFolder = "Inbox"
)

func (r SettingDomainEditResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseFolderAllItems, SettingDomainEditResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainEditResponseInboxProvider string

const (
	SettingDomainEditResponseInboxProviderMicrosoft SettingDomainEditResponseInboxProvider = "Microsoft"
	SettingDomainEditResponseInboxProviderGoogle    SettingDomainEditResponseInboxProvider = "Google"
)

func (r SettingDomainEditResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainEditResponseInboxProviderMicrosoft, SettingDomainEditResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainGetResponse struct {
	// The unique identifier for the domain.
	ID                   int64                                         `json:"id,required"`
	AllowedDeliveryModes []SettingDomainGetResponseAllowedDeliveryMode `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                     `json:"created_at,required" format:"date-time"`
	Domain               string                                        `json:"domain,required"`
	DropDispositions     []SettingDomainGetResponseDropDisposition     `json:"drop_dispositions,required"`
	IPRestrictions       []string                                      `json:"ip_restrictions,required"`
	LastModified         time.Time                                     `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                         `json:"lookback_hops,required"`
	Transport            string                                        `json:"transport,required"`
	Authorization        SettingDomainGetResponseAuthorization         `json:"authorization,nullable"`
	EmailsProcessed      SettingDomainGetResponseEmailsProcessed       `json:"emails_processed,nullable"`
	Folder               SettingDomainGetResponseFolder                `json:"folder,nullable"`
	InboxProvider        SettingDomainGetResponseInboxProvider         `json:"inbox_provider,nullable"`
	IntegrationID        string                                        `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                        `json:"o365_tenant_id,nullable"`
	RequireTLSInbound    bool                                          `json:"require_tls_inbound,nullable"`
	RequireTLSOutbound   bool                                          `json:"require_tls_outbound,nullable"`
	JSON                 settingDomainGetResponseJSON                  `json:"-"`
}

// settingDomainGetResponseJSON contains the JSON metadata for the struct
// [SettingDomainGetResponse]
type settingDomainGetResponseJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTLSInbound    apijson.Field
	RequireTLSOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseAllowedDeliveryMode string

const (
	SettingDomainGetResponseAllowedDeliveryModeDirect    SettingDomainGetResponseAllowedDeliveryMode = "DIRECT"
	SettingDomainGetResponseAllowedDeliveryModeBcc       SettingDomainGetResponseAllowedDeliveryMode = "BCC"
	SettingDomainGetResponseAllowedDeliveryModeJournal   SettingDomainGetResponseAllowedDeliveryMode = "JOURNAL"
	SettingDomainGetResponseAllowedDeliveryModeAPI       SettingDomainGetResponseAllowedDeliveryMode = "API"
	SettingDomainGetResponseAllowedDeliveryModeRetroScan SettingDomainGetResponseAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainGetResponseAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseAllowedDeliveryModeDirect, SettingDomainGetResponseAllowedDeliveryModeBcc, SettingDomainGetResponseAllowedDeliveryModeJournal, SettingDomainGetResponseAllowedDeliveryModeAPI, SettingDomainGetResponseAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainGetResponseDropDisposition string

const (
	SettingDomainGetResponseDropDispositionMalicious    SettingDomainGetResponseDropDisposition = "MALICIOUS"
	SettingDomainGetResponseDropDispositionMaliciousBec SettingDomainGetResponseDropDisposition = "MALICIOUS-BEC"
	SettingDomainGetResponseDropDispositionSuspicious   SettingDomainGetResponseDropDisposition = "SUSPICIOUS"
	SettingDomainGetResponseDropDispositionSpoof        SettingDomainGetResponseDropDisposition = "SPOOF"
	SettingDomainGetResponseDropDispositionSpam         SettingDomainGetResponseDropDisposition = "SPAM"
	SettingDomainGetResponseDropDispositionBulk         SettingDomainGetResponseDropDisposition = "BULK"
	SettingDomainGetResponseDropDispositionEncrypted    SettingDomainGetResponseDropDisposition = "ENCRYPTED"
	SettingDomainGetResponseDropDispositionExternal     SettingDomainGetResponseDropDisposition = "EXTERNAL"
	SettingDomainGetResponseDropDispositionUnknown      SettingDomainGetResponseDropDisposition = "UNKNOWN"
	SettingDomainGetResponseDropDispositionNone         SettingDomainGetResponseDropDisposition = "NONE"
)

func (r SettingDomainGetResponseDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseDropDispositionMalicious, SettingDomainGetResponseDropDispositionMaliciousBec, SettingDomainGetResponseDropDispositionSuspicious, SettingDomainGetResponseDropDispositionSpoof, SettingDomainGetResponseDropDispositionSpam, SettingDomainGetResponseDropDispositionBulk, SettingDomainGetResponseDropDispositionEncrypted, SettingDomainGetResponseDropDispositionExternal, SettingDomainGetResponseDropDispositionUnknown, SettingDomainGetResponseDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainGetResponseAuthorization struct {
	Authorized    bool                                      `json:"authorized,required"`
	Timestamp     time.Time                                 `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                    `json:"status_message,nullable"`
	JSON          settingDomainGetResponseAuthorizationJSON `json:"-"`
}

// settingDomainGetResponseAuthorizationJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseAuthorization]
type settingDomainGetResponseAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDomainGetResponseAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseAuthorizationJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseEmailsProcessed struct {
	Timestamp                    time.Time                                   `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                       `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                       `json:"total_emails_processed_previous,required"`
	JSON                         settingDomainGetResponseEmailsProcessedJSON `json:"-"`
}

// settingDomainGetResponseEmailsProcessedJSON contains the JSON metadata for the
// struct [SettingDomainGetResponseEmailsProcessed]
type settingDomainGetResponseEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SettingDomainGetResponseEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetResponseFolder string

const (
	SettingDomainGetResponseFolderAllItems SettingDomainGetResponseFolder = "AllItems"
	SettingDomainGetResponseFolderInbox    SettingDomainGetResponseFolder = "Inbox"
)

func (r SettingDomainGetResponseFolder) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseFolderAllItems, SettingDomainGetResponseFolderInbox:
		return true
	}
	return false
}

type SettingDomainGetResponseInboxProvider string

const (
	SettingDomainGetResponseInboxProviderMicrosoft SettingDomainGetResponseInboxProvider = "Microsoft"
	SettingDomainGetResponseInboxProviderGoogle    SettingDomainGetResponseInboxProvider = "Google"
)

func (r SettingDomainGetResponseInboxProvider) IsKnown() bool {
	switch r {
	case SettingDomainGetResponseInboxProviderMicrosoft, SettingDomainGetResponseInboxProviderGoogle:
		return true
	}
	return false
}

type SettingDomainListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Filters response to domains with the currently active delivery mode.
	ActiveDeliveryMode param.Field[SettingDomainListParamsActiveDeliveryMode] `query:"active_delivery_mode"`
	// Filters response to domains with the provided delivery mode.
	AllowedDeliveryMode param.Field[SettingDomainListParamsAllowedDeliveryMode] `query:"allowed_delivery_mode"`
	// The sorting direction.
	Direction param.Field[SettingDomainListParamsDirection] `query:"direction"`
	// Filters results by the provided domains, allowing for multiple occurrences.
	Domain param.Field[[]string] `query:"domain"`
	// The field to sort by.
	Order param.Field[SettingDomainListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingDomainListParams]'s query parameters as
// `url.Values`.
func (r SettingDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters response to domains with the currently active delivery mode.
type SettingDomainListParamsActiveDeliveryMode string

const (
	SettingDomainListParamsActiveDeliveryModeDirect    SettingDomainListParamsActiveDeliveryMode = "DIRECT"
	SettingDomainListParamsActiveDeliveryModeBcc       SettingDomainListParamsActiveDeliveryMode = "BCC"
	SettingDomainListParamsActiveDeliveryModeJournal   SettingDomainListParamsActiveDeliveryMode = "JOURNAL"
	SettingDomainListParamsActiveDeliveryModeAPI       SettingDomainListParamsActiveDeliveryMode = "API"
	SettingDomainListParamsActiveDeliveryModeRetroScan SettingDomainListParamsActiveDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListParamsActiveDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListParamsActiveDeliveryModeDirect, SettingDomainListParamsActiveDeliveryModeBcc, SettingDomainListParamsActiveDeliveryModeJournal, SettingDomainListParamsActiveDeliveryModeAPI, SettingDomainListParamsActiveDeliveryModeRetroScan:
		return true
	}
	return false
}

// Filters response to domains with the provided delivery mode.
type SettingDomainListParamsAllowedDeliveryMode string

const (
	SettingDomainListParamsAllowedDeliveryModeDirect    SettingDomainListParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainListParamsAllowedDeliveryModeBcc       SettingDomainListParamsAllowedDeliveryMode = "BCC"
	SettingDomainListParamsAllowedDeliveryModeJournal   SettingDomainListParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainListParamsAllowedDeliveryModeAPI       SettingDomainListParamsAllowedDeliveryMode = "API"
	SettingDomainListParamsAllowedDeliveryModeRetroScan SettingDomainListParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainListParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainListParamsAllowedDeliveryModeDirect, SettingDomainListParamsAllowedDeliveryModeBcc, SettingDomainListParamsAllowedDeliveryModeJournal, SettingDomainListParamsAllowedDeliveryModeAPI, SettingDomainListParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

// The sorting direction.
type SettingDomainListParamsDirection string

const (
	SettingDomainListParamsDirectionAsc  SettingDomainListParamsDirection = "asc"
	SettingDomainListParamsDirectionDesc SettingDomainListParamsDirection = "desc"
)

func (r SettingDomainListParamsDirection) IsKnown() bool {
	switch r {
	case SettingDomainListParamsDirectionAsc, SettingDomainListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingDomainListParamsOrder string

const (
	SettingDomainListParamsOrderDomain    SettingDomainListParamsOrder = "domain"
	SettingDomainListParamsOrderCreatedAt SettingDomainListParamsOrder = "created_at"
)

func (r SettingDomainListParamsOrder) IsKnown() bool {
	switch r {
	case SettingDomainListParamsOrderDomain, SettingDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingDomainDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingDomainDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   SettingDomainDeleteResponse             `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     settingDomainDeleteResponseEnvelopeJSON `json:"-"`
}

// settingDomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDomainDeleteResponseEnvelope]
type settingDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainBulkDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingDomainEditParams struct {
	// Account Identifier
	AccountID            param.Field[string]                                       `path:"account_id,required"`
	IPRestrictions       param.Field[[]string]                                     `json:"ip_restrictions,required"`
	AllowedDeliveryModes param.Field[[]SettingDomainEditParamsAllowedDeliveryMode] `json:"allowed_delivery_modes"`
	Domain               param.Field[string]                                       `json:"domain"`
	DropDispositions     param.Field[[]SettingDomainEditParamsDropDisposition]     `json:"drop_dispositions"`
	Folder               param.Field[SettingDomainEditParamsFolder]                `json:"folder"`
	IntegrationID        param.Field[string]                                       `json:"integration_id" format:"uuid"`
	LookbackHops         param.Field[int64]                                        `json:"lookback_hops"`
	RequireTLSInbound    param.Field[bool]                                         `json:"require_tls_inbound"`
	RequireTLSOutbound   param.Field[bool]                                         `json:"require_tls_outbound"`
	Transport            param.Field[string]                                       `json:"transport"`
}

func (r SettingDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingDomainEditParamsAllowedDeliveryMode string

const (
	SettingDomainEditParamsAllowedDeliveryModeDirect    SettingDomainEditParamsAllowedDeliveryMode = "DIRECT"
	SettingDomainEditParamsAllowedDeliveryModeBcc       SettingDomainEditParamsAllowedDeliveryMode = "BCC"
	SettingDomainEditParamsAllowedDeliveryModeJournal   SettingDomainEditParamsAllowedDeliveryMode = "JOURNAL"
	SettingDomainEditParamsAllowedDeliveryModeAPI       SettingDomainEditParamsAllowedDeliveryMode = "API"
	SettingDomainEditParamsAllowedDeliveryModeRetroScan SettingDomainEditParamsAllowedDeliveryMode = "RETRO_SCAN"
)

func (r SettingDomainEditParamsAllowedDeliveryMode) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsAllowedDeliveryModeDirect, SettingDomainEditParamsAllowedDeliveryModeBcc, SettingDomainEditParamsAllowedDeliveryModeJournal, SettingDomainEditParamsAllowedDeliveryModeAPI, SettingDomainEditParamsAllowedDeliveryModeRetroScan:
		return true
	}
	return false
}

type SettingDomainEditParamsDropDisposition string

const (
	SettingDomainEditParamsDropDispositionMalicious    SettingDomainEditParamsDropDisposition = "MALICIOUS"
	SettingDomainEditParamsDropDispositionMaliciousBec SettingDomainEditParamsDropDisposition = "MALICIOUS-BEC"
	SettingDomainEditParamsDropDispositionSuspicious   SettingDomainEditParamsDropDisposition = "SUSPICIOUS"
	SettingDomainEditParamsDropDispositionSpoof        SettingDomainEditParamsDropDisposition = "SPOOF"
	SettingDomainEditParamsDropDispositionSpam         SettingDomainEditParamsDropDisposition = "SPAM"
	SettingDomainEditParamsDropDispositionBulk         SettingDomainEditParamsDropDisposition = "BULK"
	SettingDomainEditParamsDropDispositionEncrypted    SettingDomainEditParamsDropDisposition = "ENCRYPTED"
	SettingDomainEditParamsDropDispositionExternal     SettingDomainEditParamsDropDisposition = "EXTERNAL"
	SettingDomainEditParamsDropDispositionUnknown      SettingDomainEditParamsDropDisposition = "UNKNOWN"
	SettingDomainEditParamsDropDispositionNone         SettingDomainEditParamsDropDisposition = "NONE"
)

func (r SettingDomainEditParamsDropDisposition) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsDropDispositionMalicious, SettingDomainEditParamsDropDispositionMaliciousBec, SettingDomainEditParamsDropDispositionSuspicious, SettingDomainEditParamsDropDispositionSpoof, SettingDomainEditParamsDropDispositionSpam, SettingDomainEditParamsDropDispositionBulk, SettingDomainEditParamsDropDispositionEncrypted, SettingDomainEditParamsDropDispositionExternal, SettingDomainEditParamsDropDispositionUnknown, SettingDomainEditParamsDropDispositionNone:
		return true
	}
	return false
}

type SettingDomainEditParamsFolder string

const (
	SettingDomainEditParamsFolderAllItems SettingDomainEditParamsFolder = "AllItems"
	SettingDomainEditParamsFolderInbox    SettingDomainEditParamsFolder = "Inbox"
)

func (r SettingDomainEditParamsFolder) IsKnown() bool {
	switch r {
	case SettingDomainEditParamsFolderAllItems, SettingDomainEditParamsFolderInbox:
		return true
	}
	return false
}

type SettingDomainEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   SettingDomainEditResponse             `json:"result,required"`
	Success  bool                                  `json:"success,required"`
	JSON     settingDomainEditResponseEnvelopeJSON `json:"-"`
}

// settingDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainEditResponseEnvelope]
type settingDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDomainGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                `json:"errors,required"`
	Messages []shared.ResponseInfo                `json:"messages,required"`
	Result   SettingDomainGetResponse             `json:"result,required"`
	Success  bool                                 `json:"success,required"`
	JSON     settingDomainGetResponseEnvelopeJSON `json:"-"`
}

// settingDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingDomainGetResponseEnvelope]
type settingDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
