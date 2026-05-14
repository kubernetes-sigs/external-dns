// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// CloudIntegrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudIntegrationService] method instead.
type CloudIntegrationService struct {
	Options []option.RequestOption
}

// NewCloudIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudIntegrationService(opts ...option.RequestOption) (r *CloudIntegrationService) {
	r = &CloudIntegrationService{}
	r.Options = opts
	return
}

// Create a new Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) New(ctx context.Context, params CloudIntegrationNewParams, opts ...option.RequestOption) (res *CloudIntegrationNewResponse, err error) {
	var env CloudIntegrationNewResponseEnvelope
	if params.Forwarded.Present {
		opts = append(opts, option.WithHeader("forwarded", fmt.Sprintf("%s", params.Forwarded)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) Update(ctx context.Context, providerID string, params CloudIntegrationUpdateParams, opts ...option.RequestOption) (res *CloudIntegrationUpdateResponse, err error) {
	var env CloudIntegrationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", params.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Cloud Integrations (Closed Beta).
func (r *CloudIntegrationService) List(ctx context.Context, params CloudIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CloudIntegrationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers", params.AccountID)
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

// List Cloud Integrations (Closed Beta).
func (r *CloudIntegrationService) ListAutoPaging(ctx context.Context, params CloudIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CloudIntegrationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) Delete(ctx context.Context, providerID string, body CloudIntegrationDeleteParams, opts ...option.RequestOption) (res *CloudIntegrationDeleteResponse, err error) {
	var env CloudIntegrationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", body.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Run discovery for a Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) Discover(ctx context.Context, providerID string, params CloudIntegrationDiscoverParams, opts ...option.RequestOption) (res *CloudIntegrationDiscoverResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s/discover", params.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Run discovery for all Cloud Integrations in an account (Closed Beta).
func (r *CloudIntegrationService) DiscoverAll(ctx context.Context, body CloudIntegrationDiscoverAllParams, opts ...option.RequestOption) (res *CloudIntegrationDiscoverAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/discover", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update a Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) Edit(ctx context.Context, providerID string, params CloudIntegrationEditParams, opts ...option.RequestOption) (res *CloudIntegrationEditResponse, err error) {
	var env CloudIntegrationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", params.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read a Cloud Integration (Closed Beta).
func (r *CloudIntegrationService) Get(ctx context.Context, providerID string, params CloudIntegrationGetParams, opts ...option.RequestOption) (res *CloudIntegrationGetResponse, err error) {
	var env CloudIntegrationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", params.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get initial configuration to complete Cloud Integration setup (Closed Beta).
func (r *CloudIntegrationService) InitialSetup(ctx context.Context, providerID string, query CloudIntegrationInitialSetupParams, opts ...option.RequestOption) (res *CloudIntegrationInitialSetupResponse, err error) {
	var env CloudIntegrationInitialSetupResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s/initial_setup", query.AccountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudIntegrationNewResponse struct {
	ID                     string                                    `json:"id,required" format:"uuid"`
	CloudType              CloudIntegrationNewResponseCloudType      `json:"cloud_type,required"`
	FriendlyName           string                                    `json:"friendly_name,required"`
	LastUpdated            string                                    `json:"last_updated,required"`
	LifecycleState         CloudIntegrationNewResponseLifecycleState `json:"lifecycle_state,required"`
	State                  CloudIntegrationNewResponseState          `json:"state,required"`
	StateV2                CloudIntegrationNewResponseStateV2        `json:"state_v2,required"`
	AwsArn                 string                                    `json:"aws_arn"`
	AzureSubscriptionID    string                                    `json:"azure_subscription_id"`
	AzureTenantID          string                                    `json:"azure_tenant_id"`
	Description            string                                    `json:"description"`
	GcpProjectID           string                                    `json:"gcp_project_id"`
	GcpServiceAccountEmail string                                    `json:"gcp_service_account_email"`
	Status                 CloudIntegrationNewResponseStatus         `json:"status"`
	JSON                   cloudIntegrationNewResponseJSON           `json:"-"`
}

// cloudIntegrationNewResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationNewResponse]
type cloudIntegrationNewResponseJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseCloudType string

const (
	CloudIntegrationNewResponseCloudTypeAws        CloudIntegrationNewResponseCloudType = "AWS"
	CloudIntegrationNewResponseCloudTypeAzure      CloudIntegrationNewResponseCloudType = "AZURE"
	CloudIntegrationNewResponseCloudTypeGoogle     CloudIntegrationNewResponseCloudType = "GOOGLE"
	CloudIntegrationNewResponseCloudTypeCloudflare CloudIntegrationNewResponseCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationNewResponseCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseCloudTypeAws, CloudIntegrationNewResponseCloudTypeAzure, CloudIntegrationNewResponseCloudTypeGoogle, CloudIntegrationNewResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationNewResponseLifecycleState string

const (
	CloudIntegrationNewResponseLifecycleStateActive       CloudIntegrationNewResponseLifecycleState = "ACTIVE"
	CloudIntegrationNewResponseLifecycleStatePendingSetup CloudIntegrationNewResponseLifecycleState = "PENDING_SETUP"
	CloudIntegrationNewResponseLifecycleStateRetired      CloudIntegrationNewResponseLifecycleState = "RETIRED"
)

func (r CloudIntegrationNewResponseLifecycleState) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseLifecycleStateActive, CloudIntegrationNewResponseLifecycleStatePendingSetup, CloudIntegrationNewResponseLifecycleStateRetired:
		return true
	}
	return false
}

type CloudIntegrationNewResponseState string

const (
	CloudIntegrationNewResponseStateUnspecified CloudIntegrationNewResponseState = "UNSPECIFIED"
	CloudIntegrationNewResponseStatePending     CloudIntegrationNewResponseState = "PENDING"
	CloudIntegrationNewResponseStateDiscovering CloudIntegrationNewResponseState = "DISCOVERING"
	CloudIntegrationNewResponseStateFailed      CloudIntegrationNewResponseState = "FAILED"
	CloudIntegrationNewResponseStateSucceeded   CloudIntegrationNewResponseState = "SUCCEEDED"
)

func (r CloudIntegrationNewResponseState) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseStateUnspecified, CloudIntegrationNewResponseStatePending, CloudIntegrationNewResponseStateDiscovering, CloudIntegrationNewResponseStateFailed, CloudIntegrationNewResponseStateSucceeded:
		return true
	}
	return false
}

type CloudIntegrationNewResponseStateV2 string

const (
	CloudIntegrationNewResponseStateV2Unspecified CloudIntegrationNewResponseStateV2 = "UNSPECIFIED"
	CloudIntegrationNewResponseStateV2Pending     CloudIntegrationNewResponseStateV2 = "PENDING"
	CloudIntegrationNewResponseStateV2Discovering CloudIntegrationNewResponseStateV2 = "DISCOVERING"
	CloudIntegrationNewResponseStateV2Failed      CloudIntegrationNewResponseStateV2 = "FAILED"
	CloudIntegrationNewResponseStateV2Succeeded   CloudIntegrationNewResponseStateV2 = "SUCCEEDED"
)

func (r CloudIntegrationNewResponseStateV2) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseStateV2Unspecified, CloudIntegrationNewResponseStateV2Pending, CloudIntegrationNewResponseStateV2Discovering, CloudIntegrationNewResponseStateV2Failed, CloudIntegrationNewResponseStateV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationNewResponseStatus struct {
	DiscoveryProgress          CloudIntegrationNewResponseStatusDiscoveryProgress     `json:"discovery_progress,required"`
	DiscoveryProgressV2        CloudIntegrationNewResponseStatusDiscoveryProgressV2   `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        CloudIntegrationNewResponseStatusLastDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 `json:"last_discovery_status_v2,required"`
	Regions                    []string                                               `json:"regions,required"`
	CredentialsGoodSince       string                                                 `json:"credentials_good_since"`
	CredentialsMissingSince    string                                                 `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                                                 `json:"credentials_rejected_since"`
	DiscoveryMessage           string                                                 `json:"discovery_message"`
	DiscoveryMessageV2         string                                                 `json:"discovery_message_v2"`
	InUseBy                    []CloudIntegrationNewResponseStatusInUseBy             `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                                                 `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                                                 `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                                                 `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                                                 `json:"last_discovery_started_at_v2"`
	LastUpdated                string                                                 `json:"last_updated"`
	JSON                       cloudIntegrationNewResponseStatusJSON                  `json:"-"`
}

// cloudIntegrationNewResponseStatusJSON contains the JSON metadata for the struct
// [CloudIntegrationNewResponseStatus]
type cloudIntegrationNewResponseStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseStatusDiscoveryProgress struct {
	Done  int64                                                  `json:"done,required"`
	Total int64                                                  `json:"total,required"`
	Unit  string                                                 `json:"unit,required"`
	JSON  cloudIntegrationNewResponseStatusDiscoveryProgressJSON `json:"-"`
}

// cloudIntegrationNewResponseStatusDiscoveryProgressJSON contains the JSON
// metadata for the struct [CloudIntegrationNewResponseStatusDiscoveryProgress]
type cloudIntegrationNewResponseStatusDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseStatusDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseStatusDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseStatusDiscoveryProgressV2 struct {
	Done  int64                                                    `json:"done,required"`
	Total int64                                                    `json:"total,required"`
	Unit  string                                                   `json:"unit,required"`
	JSON  cloudIntegrationNewResponseStatusDiscoveryProgressV2JSON `json:"-"`
}

// cloudIntegrationNewResponseStatusDiscoveryProgressV2JSON contains the JSON
// metadata for the struct [CloudIntegrationNewResponseStatusDiscoveryProgressV2]
type cloudIntegrationNewResponseStatusDiscoveryProgressV2JSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseStatusDiscoveryProgressV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseStatusDiscoveryProgressV2JSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseStatusLastDiscoveryStatus string

const (
	CloudIntegrationNewResponseStatusLastDiscoveryStatusUnspecified CloudIntegrationNewResponseStatusLastDiscoveryStatus = "UNSPECIFIED"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusPending     CloudIntegrationNewResponseStatusLastDiscoveryStatus = "PENDING"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusDiscovering CloudIntegrationNewResponseStatusLastDiscoveryStatus = "DISCOVERING"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusFailed      CloudIntegrationNewResponseStatusLastDiscoveryStatus = "FAILED"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusSucceeded   CloudIntegrationNewResponseStatusLastDiscoveryStatus = "SUCCEEDED"
)

func (r CloudIntegrationNewResponseStatusLastDiscoveryStatus) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseStatusLastDiscoveryStatusUnspecified, CloudIntegrationNewResponseStatusLastDiscoveryStatusPending, CloudIntegrationNewResponseStatusLastDiscoveryStatusDiscovering, CloudIntegrationNewResponseStatusLastDiscoveryStatusFailed, CloudIntegrationNewResponseStatusLastDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 string

const (
	CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Unspecified CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 = "UNSPECIFIED"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Pending     CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 = "PENDING"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Discovering CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 = "DISCOVERING"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Failed      CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 = "FAILED"
	CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Succeeded   CloudIntegrationNewResponseStatusLastDiscoveryStatusV2 = "SUCCEEDED"
)

func (r CloudIntegrationNewResponseStatusLastDiscoveryStatusV2) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Unspecified, CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Pending, CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Discovering, CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Failed, CloudIntegrationNewResponseStatusLastDiscoveryStatusV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationNewResponseStatusInUseBy struct {
	ID         string                                             `json:"id,required" format:"uuid"`
	ClientType CloudIntegrationNewResponseStatusInUseByClientType `json:"client_type,required"`
	Name       string                                             `json:"name,required"`
	JSON       cloudIntegrationNewResponseStatusInUseByJSON       `json:"-"`
}

// cloudIntegrationNewResponseStatusInUseByJSON contains the JSON metadata for the
// struct [CloudIntegrationNewResponseStatusInUseBy]
type cloudIntegrationNewResponseStatusInUseByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseStatusInUseBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseStatusInUseByJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseStatusInUseByClientType string

const (
	CloudIntegrationNewResponseStatusInUseByClientTypeMagicWANCloudOnramp CloudIntegrationNewResponseStatusInUseByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r CloudIntegrationNewResponseStatusInUseByClientType) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseStatusInUseByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponse struct {
	ID                     string                                       `json:"id,required" format:"uuid"`
	CloudType              CloudIntegrationUpdateResponseCloudType      `json:"cloud_type,required"`
	FriendlyName           string                                       `json:"friendly_name,required"`
	LastUpdated            string                                       `json:"last_updated,required"`
	LifecycleState         CloudIntegrationUpdateResponseLifecycleState `json:"lifecycle_state,required"`
	State                  CloudIntegrationUpdateResponseState          `json:"state,required"`
	StateV2                CloudIntegrationUpdateResponseStateV2        `json:"state_v2,required"`
	AwsArn                 string                                       `json:"aws_arn"`
	AzureSubscriptionID    string                                       `json:"azure_subscription_id"`
	AzureTenantID          string                                       `json:"azure_tenant_id"`
	Description            string                                       `json:"description"`
	GcpProjectID           string                                       `json:"gcp_project_id"`
	GcpServiceAccountEmail string                                       `json:"gcp_service_account_email"`
	Status                 CloudIntegrationUpdateResponseStatus         `json:"status"`
	JSON                   cloudIntegrationUpdateResponseJSON           `json:"-"`
}

// cloudIntegrationUpdateResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationUpdateResponse]
type cloudIntegrationUpdateResponseJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseCloudType string

const (
	CloudIntegrationUpdateResponseCloudTypeAws        CloudIntegrationUpdateResponseCloudType = "AWS"
	CloudIntegrationUpdateResponseCloudTypeAzure      CloudIntegrationUpdateResponseCloudType = "AZURE"
	CloudIntegrationUpdateResponseCloudTypeGoogle     CloudIntegrationUpdateResponseCloudType = "GOOGLE"
	CloudIntegrationUpdateResponseCloudTypeCloudflare CloudIntegrationUpdateResponseCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationUpdateResponseCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseCloudTypeAws, CloudIntegrationUpdateResponseCloudTypeAzure, CloudIntegrationUpdateResponseCloudTypeGoogle, CloudIntegrationUpdateResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseLifecycleState string

const (
	CloudIntegrationUpdateResponseLifecycleStateActive       CloudIntegrationUpdateResponseLifecycleState = "ACTIVE"
	CloudIntegrationUpdateResponseLifecycleStatePendingSetup CloudIntegrationUpdateResponseLifecycleState = "PENDING_SETUP"
	CloudIntegrationUpdateResponseLifecycleStateRetired      CloudIntegrationUpdateResponseLifecycleState = "RETIRED"
)

func (r CloudIntegrationUpdateResponseLifecycleState) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseLifecycleStateActive, CloudIntegrationUpdateResponseLifecycleStatePendingSetup, CloudIntegrationUpdateResponseLifecycleStateRetired:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseState string

const (
	CloudIntegrationUpdateResponseStateUnspecified CloudIntegrationUpdateResponseState = "UNSPECIFIED"
	CloudIntegrationUpdateResponseStatePending     CloudIntegrationUpdateResponseState = "PENDING"
	CloudIntegrationUpdateResponseStateDiscovering CloudIntegrationUpdateResponseState = "DISCOVERING"
	CloudIntegrationUpdateResponseStateFailed      CloudIntegrationUpdateResponseState = "FAILED"
	CloudIntegrationUpdateResponseStateSucceeded   CloudIntegrationUpdateResponseState = "SUCCEEDED"
)

func (r CloudIntegrationUpdateResponseState) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseStateUnspecified, CloudIntegrationUpdateResponseStatePending, CloudIntegrationUpdateResponseStateDiscovering, CloudIntegrationUpdateResponseStateFailed, CloudIntegrationUpdateResponseStateSucceeded:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseStateV2 string

const (
	CloudIntegrationUpdateResponseStateV2Unspecified CloudIntegrationUpdateResponseStateV2 = "UNSPECIFIED"
	CloudIntegrationUpdateResponseStateV2Pending     CloudIntegrationUpdateResponseStateV2 = "PENDING"
	CloudIntegrationUpdateResponseStateV2Discovering CloudIntegrationUpdateResponseStateV2 = "DISCOVERING"
	CloudIntegrationUpdateResponseStateV2Failed      CloudIntegrationUpdateResponseStateV2 = "FAILED"
	CloudIntegrationUpdateResponseStateV2Succeeded   CloudIntegrationUpdateResponseStateV2 = "SUCCEEDED"
)

func (r CloudIntegrationUpdateResponseStateV2) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseStateV2Unspecified, CloudIntegrationUpdateResponseStateV2Pending, CloudIntegrationUpdateResponseStateV2Discovering, CloudIntegrationUpdateResponseStateV2Failed, CloudIntegrationUpdateResponseStateV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseStatus struct {
	DiscoveryProgress          CloudIntegrationUpdateResponseStatusDiscoveryProgress     `json:"discovery_progress,required"`
	DiscoveryProgressV2        CloudIntegrationUpdateResponseStatusDiscoveryProgressV2   `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        CloudIntegrationUpdateResponseStatusLastDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 `json:"last_discovery_status_v2,required"`
	Regions                    []string                                                  `json:"regions,required"`
	CredentialsGoodSince       string                                                    `json:"credentials_good_since"`
	CredentialsMissingSince    string                                                    `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                                                    `json:"credentials_rejected_since"`
	DiscoveryMessage           string                                                    `json:"discovery_message"`
	DiscoveryMessageV2         string                                                    `json:"discovery_message_v2"`
	InUseBy                    []CloudIntegrationUpdateResponseStatusInUseBy             `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                                                    `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                                                    `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                                                    `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                                                    `json:"last_discovery_started_at_v2"`
	LastUpdated                string                                                    `json:"last_updated"`
	JSON                       cloudIntegrationUpdateResponseStatusJSON                  `json:"-"`
}

// cloudIntegrationUpdateResponseStatusJSON contains the JSON metadata for the
// struct [CloudIntegrationUpdateResponseStatus]
type cloudIntegrationUpdateResponseStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseStatusDiscoveryProgress struct {
	Done  int64                                                     `json:"done,required"`
	Total int64                                                     `json:"total,required"`
	Unit  string                                                    `json:"unit,required"`
	JSON  cloudIntegrationUpdateResponseStatusDiscoveryProgressJSON `json:"-"`
}

// cloudIntegrationUpdateResponseStatusDiscoveryProgressJSON contains the JSON
// metadata for the struct [CloudIntegrationUpdateResponseStatusDiscoveryProgress]
type cloudIntegrationUpdateResponseStatusDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseStatusDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseStatusDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseStatusDiscoveryProgressV2 struct {
	Done  int64                                                       `json:"done,required"`
	Total int64                                                       `json:"total,required"`
	Unit  string                                                      `json:"unit,required"`
	JSON  cloudIntegrationUpdateResponseStatusDiscoveryProgressV2JSON `json:"-"`
}

// cloudIntegrationUpdateResponseStatusDiscoveryProgressV2JSON contains the JSON
// metadata for the struct
// [CloudIntegrationUpdateResponseStatusDiscoveryProgressV2]
type cloudIntegrationUpdateResponseStatusDiscoveryProgressV2JSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseStatusDiscoveryProgressV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseStatusDiscoveryProgressV2JSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseStatusLastDiscoveryStatus string

const (
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusUnspecified CloudIntegrationUpdateResponseStatusLastDiscoveryStatus = "UNSPECIFIED"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusPending     CloudIntegrationUpdateResponseStatusLastDiscoveryStatus = "PENDING"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusDiscovering CloudIntegrationUpdateResponseStatusLastDiscoveryStatus = "DISCOVERING"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusFailed      CloudIntegrationUpdateResponseStatusLastDiscoveryStatus = "FAILED"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusSucceeded   CloudIntegrationUpdateResponseStatusLastDiscoveryStatus = "SUCCEEDED"
)

func (r CloudIntegrationUpdateResponseStatusLastDiscoveryStatus) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseStatusLastDiscoveryStatusUnspecified, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusPending, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusDiscovering, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusFailed, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 string

const (
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Unspecified CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 = "UNSPECIFIED"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Pending     CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 = "PENDING"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Discovering CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 = "DISCOVERING"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Failed      CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 = "FAILED"
	CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Succeeded   CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2 = "SUCCEEDED"
)

func (r CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Unspecified, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Pending, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Discovering, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Failed, CloudIntegrationUpdateResponseStatusLastDiscoveryStatusV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseStatusInUseBy struct {
	ID         string                                                `json:"id,required" format:"uuid"`
	ClientType CloudIntegrationUpdateResponseStatusInUseByClientType `json:"client_type,required"`
	Name       string                                                `json:"name,required"`
	JSON       cloudIntegrationUpdateResponseStatusInUseByJSON       `json:"-"`
}

// cloudIntegrationUpdateResponseStatusInUseByJSON contains the JSON metadata for
// the struct [CloudIntegrationUpdateResponseStatusInUseBy]
type cloudIntegrationUpdateResponseStatusInUseByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseStatusInUseBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseStatusInUseByJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseStatusInUseByClientType string

const (
	CloudIntegrationUpdateResponseStatusInUseByClientTypeMagicWANCloudOnramp CloudIntegrationUpdateResponseStatusInUseByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r CloudIntegrationUpdateResponseStatusInUseByClientType) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseStatusInUseByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type CloudIntegrationListResponse struct {
	ID                     string                                     `json:"id,required" format:"uuid"`
	CloudType              CloudIntegrationListResponseCloudType      `json:"cloud_type,required"`
	FriendlyName           string                                     `json:"friendly_name,required"`
	LastUpdated            string                                     `json:"last_updated,required"`
	LifecycleState         CloudIntegrationListResponseLifecycleState `json:"lifecycle_state,required"`
	State                  CloudIntegrationListResponseState          `json:"state,required"`
	StateV2                CloudIntegrationListResponseStateV2        `json:"state_v2,required"`
	AwsArn                 string                                     `json:"aws_arn"`
	AzureSubscriptionID    string                                     `json:"azure_subscription_id"`
	AzureTenantID          string                                     `json:"azure_tenant_id"`
	Description            string                                     `json:"description"`
	GcpProjectID           string                                     `json:"gcp_project_id"`
	GcpServiceAccountEmail string                                     `json:"gcp_service_account_email"`
	Status                 CloudIntegrationListResponseStatus         `json:"status"`
	JSON                   cloudIntegrationListResponseJSON           `json:"-"`
}

// cloudIntegrationListResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationListResponse]
type cloudIntegrationListResponseJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListResponseCloudType string

const (
	CloudIntegrationListResponseCloudTypeAws        CloudIntegrationListResponseCloudType = "AWS"
	CloudIntegrationListResponseCloudTypeAzure      CloudIntegrationListResponseCloudType = "AZURE"
	CloudIntegrationListResponseCloudTypeGoogle     CloudIntegrationListResponseCloudType = "GOOGLE"
	CloudIntegrationListResponseCloudTypeCloudflare CloudIntegrationListResponseCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationListResponseCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseCloudTypeAws, CloudIntegrationListResponseCloudTypeAzure, CloudIntegrationListResponseCloudTypeGoogle, CloudIntegrationListResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationListResponseLifecycleState string

const (
	CloudIntegrationListResponseLifecycleStateActive       CloudIntegrationListResponseLifecycleState = "ACTIVE"
	CloudIntegrationListResponseLifecycleStatePendingSetup CloudIntegrationListResponseLifecycleState = "PENDING_SETUP"
	CloudIntegrationListResponseLifecycleStateRetired      CloudIntegrationListResponseLifecycleState = "RETIRED"
)

func (r CloudIntegrationListResponseLifecycleState) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseLifecycleStateActive, CloudIntegrationListResponseLifecycleStatePendingSetup, CloudIntegrationListResponseLifecycleStateRetired:
		return true
	}
	return false
}

type CloudIntegrationListResponseState string

const (
	CloudIntegrationListResponseStateUnspecified CloudIntegrationListResponseState = "UNSPECIFIED"
	CloudIntegrationListResponseStatePending     CloudIntegrationListResponseState = "PENDING"
	CloudIntegrationListResponseStateDiscovering CloudIntegrationListResponseState = "DISCOVERING"
	CloudIntegrationListResponseStateFailed      CloudIntegrationListResponseState = "FAILED"
	CloudIntegrationListResponseStateSucceeded   CloudIntegrationListResponseState = "SUCCEEDED"
)

func (r CloudIntegrationListResponseState) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseStateUnspecified, CloudIntegrationListResponseStatePending, CloudIntegrationListResponseStateDiscovering, CloudIntegrationListResponseStateFailed, CloudIntegrationListResponseStateSucceeded:
		return true
	}
	return false
}

type CloudIntegrationListResponseStateV2 string

const (
	CloudIntegrationListResponseStateV2Unspecified CloudIntegrationListResponseStateV2 = "UNSPECIFIED"
	CloudIntegrationListResponseStateV2Pending     CloudIntegrationListResponseStateV2 = "PENDING"
	CloudIntegrationListResponseStateV2Discovering CloudIntegrationListResponseStateV2 = "DISCOVERING"
	CloudIntegrationListResponseStateV2Failed      CloudIntegrationListResponseStateV2 = "FAILED"
	CloudIntegrationListResponseStateV2Succeeded   CloudIntegrationListResponseStateV2 = "SUCCEEDED"
)

func (r CloudIntegrationListResponseStateV2) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseStateV2Unspecified, CloudIntegrationListResponseStateV2Pending, CloudIntegrationListResponseStateV2Discovering, CloudIntegrationListResponseStateV2Failed, CloudIntegrationListResponseStateV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationListResponseStatus struct {
	DiscoveryProgress          CloudIntegrationListResponseStatusDiscoveryProgress     `json:"discovery_progress,required"`
	DiscoveryProgressV2        CloudIntegrationListResponseStatusDiscoveryProgressV2   `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        CloudIntegrationListResponseStatusLastDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      CloudIntegrationListResponseStatusLastDiscoveryStatusV2 `json:"last_discovery_status_v2,required"`
	Regions                    []string                                                `json:"regions,required"`
	CredentialsGoodSince       string                                                  `json:"credentials_good_since"`
	CredentialsMissingSince    string                                                  `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                                                  `json:"credentials_rejected_since"`
	DiscoveryMessage           string                                                  `json:"discovery_message"`
	DiscoveryMessageV2         string                                                  `json:"discovery_message_v2"`
	InUseBy                    []CloudIntegrationListResponseStatusInUseBy             `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                                                  `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                                                  `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                                                  `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                                                  `json:"last_discovery_started_at_v2"`
	LastUpdated                string                                                  `json:"last_updated"`
	JSON                       cloudIntegrationListResponseStatusJSON                  `json:"-"`
}

// cloudIntegrationListResponseStatusJSON contains the JSON metadata for the struct
// [CloudIntegrationListResponseStatus]
type cloudIntegrationListResponseStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CloudIntegrationListResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationListResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListResponseStatusDiscoveryProgress struct {
	Done  int64                                                   `json:"done,required"`
	Total int64                                                   `json:"total,required"`
	Unit  string                                                  `json:"unit,required"`
	JSON  cloudIntegrationListResponseStatusDiscoveryProgressJSON `json:"-"`
}

// cloudIntegrationListResponseStatusDiscoveryProgressJSON contains the JSON
// metadata for the struct [CloudIntegrationListResponseStatusDiscoveryProgress]
type cloudIntegrationListResponseStatusDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationListResponseStatusDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationListResponseStatusDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListResponseStatusDiscoveryProgressV2 struct {
	Done  int64                                                     `json:"done,required"`
	Total int64                                                     `json:"total,required"`
	Unit  string                                                    `json:"unit,required"`
	JSON  cloudIntegrationListResponseStatusDiscoveryProgressV2JSON `json:"-"`
}

// cloudIntegrationListResponseStatusDiscoveryProgressV2JSON contains the JSON
// metadata for the struct [CloudIntegrationListResponseStatusDiscoveryProgressV2]
type cloudIntegrationListResponseStatusDiscoveryProgressV2JSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationListResponseStatusDiscoveryProgressV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationListResponseStatusDiscoveryProgressV2JSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListResponseStatusLastDiscoveryStatus string

const (
	CloudIntegrationListResponseStatusLastDiscoveryStatusUnspecified CloudIntegrationListResponseStatusLastDiscoveryStatus = "UNSPECIFIED"
	CloudIntegrationListResponseStatusLastDiscoveryStatusPending     CloudIntegrationListResponseStatusLastDiscoveryStatus = "PENDING"
	CloudIntegrationListResponseStatusLastDiscoveryStatusDiscovering CloudIntegrationListResponseStatusLastDiscoveryStatus = "DISCOVERING"
	CloudIntegrationListResponseStatusLastDiscoveryStatusFailed      CloudIntegrationListResponseStatusLastDiscoveryStatus = "FAILED"
	CloudIntegrationListResponseStatusLastDiscoveryStatusSucceeded   CloudIntegrationListResponseStatusLastDiscoveryStatus = "SUCCEEDED"
)

func (r CloudIntegrationListResponseStatusLastDiscoveryStatus) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseStatusLastDiscoveryStatusUnspecified, CloudIntegrationListResponseStatusLastDiscoveryStatusPending, CloudIntegrationListResponseStatusLastDiscoveryStatusDiscovering, CloudIntegrationListResponseStatusLastDiscoveryStatusFailed, CloudIntegrationListResponseStatusLastDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type CloudIntegrationListResponseStatusLastDiscoveryStatusV2 string

const (
	CloudIntegrationListResponseStatusLastDiscoveryStatusV2Unspecified CloudIntegrationListResponseStatusLastDiscoveryStatusV2 = "UNSPECIFIED"
	CloudIntegrationListResponseStatusLastDiscoveryStatusV2Pending     CloudIntegrationListResponseStatusLastDiscoveryStatusV2 = "PENDING"
	CloudIntegrationListResponseStatusLastDiscoveryStatusV2Discovering CloudIntegrationListResponseStatusLastDiscoveryStatusV2 = "DISCOVERING"
	CloudIntegrationListResponseStatusLastDiscoveryStatusV2Failed      CloudIntegrationListResponseStatusLastDiscoveryStatusV2 = "FAILED"
	CloudIntegrationListResponseStatusLastDiscoveryStatusV2Succeeded   CloudIntegrationListResponseStatusLastDiscoveryStatusV2 = "SUCCEEDED"
)

func (r CloudIntegrationListResponseStatusLastDiscoveryStatusV2) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseStatusLastDiscoveryStatusV2Unspecified, CloudIntegrationListResponseStatusLastDiscoveryStatusV2Pending, CloudIntegrationListResponseStatusLastDiscoveryStatusV2Discovering, CloudIntegrationListResponseStatusLastDiscoveryStatusV2Failed, CloudIntegrationListResponseStatusLastDiscoveryStatusV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationListResponseStatusInUseBy struct {
	ID         string                                              `json:"id,required" format:"uuid"`
	ClientType CloudIntegrationListResponseStatusInUseByClientType `json:"client_type,required"`
	Name       string                                              `json:"name,required"`
	JSON       cloudIntegrationListResponseStatusInUseByJSON       `json:"-"`
}

// cloudIntegrationListResponseStatusInUseByJSON contains the JSON metadata for the
// struct [CloudIntegrationListResponseStatusInUseBy]
type cloudIntegrationListResponseStatusInUseByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationListResponseStatusInUseBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationListResponseStatusInUseByJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListResponseStatusInUseByClientType string

const (
	CloudIntegrationListResponseStatusInUseByClientTypeMagicWANCloudOnramp CloudIntegrationListResponseStatusInUseByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r CloudIntegrationListResponseStatusInUseByClientType) IsKnown() bool {
	switch r {
	case CloudIntegrationListResponseStatusInUseByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type CloudIntegrationDeleteResponse struct {
	ID   string                             `json:"id,required" format:"uuid"`
	JSON cloudIntegrationDeleteResponseJSON `json:"-"`
}

// cloudIntegrationDeleteResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationDeleteResponse]
type cloudIntegrationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponse struct {
	Errors   []CloudIntegrationDiscoverResponseError   `json:"errors,required"`
	Messages []CloudIntegrationDiscoverResponseMessage `json:"messages,required"`
	Success  bool                                      `json:"success,required"`
	JSON     cloudIntegrationDiscoverResponseJSON      `json:"-"`
}

// cloudIntegrationDiscoverResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationDiscoverResponse]
type cloudIntegrationDiscoverResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseError struct {
	Code             CloudIntegrationDiscoverResponseErrorsCode   `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Meta             CloudIntegrationDiscoverResponseErrorsMeta   `json:"meta"`
	Source           CloudIntegrationDiscoverResponseErrorsSource `json:"source"`
	JSON             cloudIntegrationDiscoverResponseErrorJSON    `json:"-"`
}

// cloudIntegrationDiscoverResponseErrorJSON contains the JSON metadata for the
// struct [CloudIntegrationDiscoverResponseError]
type cloudIntegrationDiscoverResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseErrorsCode int64

const (
	CloudIntegrationDiscoverResponseErrorsCode1001   CloudIntegrationDiscoverResponseErrorsCode = 1001
	CloudIntegrationDiscoverResponseErrorsCode1002   CloudIntegrationDiscoverResponseErrorsCode = 1002
	CloudIntegrationDiscoverResponseErrorsCode1003   CloudIntegrationDiscoverResponseErrorsCode = 1003
	CloudIntegrationDiscoverResponseErrorsCode1004   CloudIntegrationDiscoverResponseErrorsCode = 1004
	CloudIntegrationDiscoverResponseErrorsCode1005   CloudIntegrationDiscoverResponseErrorsCode = 1005
	CloudIntegrationDiscoverResponseErrorsCode1006   CloudIntegrationDiscoverResponseErrorsCode = 1006
	CloudIntegrationDiscoverResponseErrorsCode1007   CloudIntegrationDiscoverResponseErrorsCode = 1007
	CloudIntegrationDiscoverResponseErrorsCode1008   CloudIntegrationDiscoverResponseErrorsCode = 1008
	CloudIntegrationDiscoverResponseErrorsCode1009   CloudIntegrationDiscoverResponseErrorsCode = 1009
	CloudIntegrationDiscoverResponseErrorsCode1010   CloudIntegrationDiscoverResponseErrorsCode = 1010
	CloudIntegrationDiscoverResponseErrorsCode1011   CloudIntegrationDiscoverResponseErrorsCode = 1011
	CloudIntegrationDiscoverResponseErrorsCode1012   CloudIntegrationDiscoverResponseErrorsCode = 1012
	CloudIntegrationDiscoverResponseErrorsCode1013   CloudIntegrationDiscoverResponseErrorsCode = 1013
	CloudIntegrationDiscoverResponseErrorsCode1014   CloudIntegrationDiscoverResponseErrorsCode = 1014
	CloudIntegrationDiscoverResponseErrorsCode1015   CloudIntegrationDiscoverResponseErrorsCode = 1015
	CloudIntegrationDiscoverResponseErrorsCode1016   CloudIntegrationDiscoverResponseErrorsCode = 1016
	CloudIntegrationDiscoverResponseErrorsCode1017   CloudIntegrationDiscoverResponseErrorsCode = 1017
	CloudIntegrationDiscoverResponseErrorsCode2001   CloudIntegrationDiscoverResponseErrorsCode = 2001
	CloudIntegrationDiscoverResponseErrorsCode2002   CloudIntegrationDiscoverResponseErrorsCode = 2002
	CloudIntegrationDiscoverResponseErrorsCode2003   CloudIntegrationDiscoverResponseErrorsCode = 2003
	CloudIntegrationDiscoverResponseErrorsCode2004   CloudIntegrationDiscoverResponseErrorsCode = 2004
	CloudIntegrationDiscoverResponseErrorsCode2005   CloudIntegrationDiscoverResponseErrorsCode = 2005
	CloudIntegrationDiscoverResponseErrorsCode2006   CloudIntegrationDiscoverResponseErrorsCode = 2006
	CloudIntegrationDiscoverResponseErrorsCode2007   CloudIntegrationDiscoverResponseErrorsCode = 2007
	CloudIntegrationDiscoverResponseErrorsCode2008   CloudIntegrationDiscoverResponseErrorsCode = 2008
	CloudIntegrationDiscoverResponseErrorsCode2009   CloudIntegrationDiscoverResponseErrorsCode = 2009
	CloudIntegrationDiscoverResponseErrorsCode2010   CloudIntegrationDiscoverResponseErrorsCode = 2010
	CloudIntegrationDiscoverResponseErrorsCode2011   CloudIntegrationDiscoverResponseErrorsCode = 2011
	CloudIntegrationDiscoverResponseErrorsCode2012   CloudIntegrationDiscoverResponseErrorsCode = 2012
	CloudIntegrationDiscoverResponseErrorsCode2013   CloudIntegrationDiscoverResponseErrorsCode = 2013
	CloudIntegrationDiscoverResponseErrorsCode2014   CloudIntegrationDiscoverResponseErrorsCode = 2014
	CloudIntegrationDiscoverResponseErrorsCode2015   CloudIntegrationDiscoverResponseErrorsCode = 2015
	CloudIntegrationDiscoverResponseErrorsCode2016   CloudIntegrationDiscoverResponseErrorsCode = 2016
	CloudIntegrationDiscoverResponseErrorsCode2017   CloudIntegrationDiscoverResponseErrorsCode = 2017
	CloudIntegrationDiscoverResponseErrorsCode2018   CloudIntegrationDiscoverResponseErrorsCode = 2018
	CloudIntegrationDiscoverResponseErrorsCode2019   CloudIntegrationDiscoverResponseErrorsCode = 2019
	CloudIntegrationDiscoverResponseErrorsCode2020   CloudIntegrationDiscoverResponseErrorsCode = 2020
	CloudIntegrationDiscoverResponseErrorsCode2021   CloudIntegrationDiscoverResponseErrorsCode = 2021
	CloudIntegrationDiscoverResponseErrorsCode2022   CloudIntegrationDiscoverResponseErrorsCode = 2022
	CloudIntegrationDiscoverResponseErrorsCode3001   CloudIntegrationDiscoverResponseErrorsCode = 3001
	CloudIntegrationDiscoverResponseErrorsCode3002   CloudIntegrationDiscoverResponseErrorsCode = 3002
	CloudIntegrationDiscoverResponseErrorsCode3003   CloudIntegrationDiscoverResponseErrorsCode = 3003
	CloudIntegrationDiscoverResponseErrorsCode3004   CloudIntegrationDiscoverResponseErrorsCode = 3004
	CloudIntegrationDiscoverResponseErrorsCode3005   CloudIntegrationDiscoverResponseErrorsCode = 3005
	CloudIntegrationDiscoverResponseErrorsCode3006   CloudIntegrationDiscoverResponseErrorsCode = 3006
	CloudIntegrationDiscoverResponseErrorsCode3007   CloudIntegrationDiscoverResponseErrorsCode = 3007
	CloudIntegrationDiscoverResponseErrorsCode4001   CloudIntegrationDiscoverResponseErrorsCode = 4001
	CloudIntegrationDiscoverResponseErrorsCode4002   CloudIntegrationDiscoverResponseErrorsCode = 4002
	CloudIntegrationDiscoverResponseErrorsCode4003   CloudIntegrationDiscoverResponseErrorsCode = 4003
	CloudIntegrationDiscoverResponseErrorsCode4004   CloudIntegrationDiscoverResponseErrorsCode = 4004
	CloudIntegrationDiscoverResponseErrorsCode4005   CloudIntegrationDiscoverResponseErrorsCode = 4005
	CloudIntegrationDiscoverResponseErrorsCode4006   CloudIntegrationDiscoverResponseErrorsCode = 4006
	CloudIntegrationDiscoverResponseErrorsCode4007   CloudIntegrationDiscoverResponseErrorsCode = 4007
	CloudIntegrationDiscoverResponseErrorsCode4008   CloudIntegrationDiscoverResponseErrorsCode = 4008
	CloudIntegrationDiscoverResponseErrorsCode4009   CloudIntegrationDiscoverResponseErrorsCode = 4009
	CloudIntegrationDiscoverResponseErrorsCode4010   CloudIntegrationDiscoverResponseErrorsCode = 4010
	CloudIntegrationDiscoverResponseErrorsCode4011   CloudIntegrationDiscoverResponseErrorsCode = 4011
	CloudIntegrationDiscoverResponseErrorsCode4012   CloudIntegrationDiscoverResponseErrorsCode = 4012
	CloudIntegrationDiscoverResponseErrorsCode4013   CloudIntegrationDiscoverResponseErrorsCode = 4013
	CloudIntegrationDiscoverResponseErrorsCode4014   CloudIntegrationDiscoverResponseErrorsCode = 4014
	CloudIntegrationDiscoverResponseErrorsCode4015   CloudIntegrationDiscoverResponseErrorsCode = 4015
	CloudIntegrationDiscoverResponseErrorsCode4016   CloudIntegrationDiscoverResponseErrorsCode = 4016
	CloudIntegrationDiscoverResponseErrorsCode4017   CloudIntegrationDiscoverResponseErrorsCode = 4017
	CloudIntegrationDiscoverResponseErrorsCode4018   CloudIntegrationDiscoverResponseErrorsCode = 4018
	CloudIntegrationDiscoverResponseErrorsCode4019   CloudIntegrationDiscoverResponseErrorsCode = 4019
	CloudIntegrationDiscoverResponseErrorsCode4020   CloudIntegrationDiscoverResponseErrorsCode = 4020
	CloudIntegrationDiscoverResponseErrorsCode4021   CloudIntegrationDiscoverResponseErrorsCode = 4021
	CloudIntegrationDiscoverResponseErrorsCode4022   CloudIntegrationDiscoverResponseErrorsCode = 4022
	CloudIntegrationDiscoverResponseErrorsCode4023   CloudIntegrationDiscoverResponseErrorsCode = 4023
	CloudIntegrationDiscoverResponseErrorsCode5001   CloudIntegrationDiscoverResponseErrorsCode = 5001
	CloudIntegrationDiscoverResponseErrorsCode5002   CloudIntegrationDiscoverResponseErrorsCode = 5002
	CloudIntegrationDiscoverResponseErrorsCode5003   CloudIntegrationDiscoverResponseErrorsCode = 5003
	CloudIntegrationDiscoverResponseErrorsCode5004   CloudIntegrationDiscoverResponseErrorsCode = 5004
	CloudIntegrationDiscoverResponseErrorsCode102000 CloudIntegrationDiscoverResponseErrorsCode = 102000
	CloudIntegrationDiscoverResponseErrorsCode102001 CloudIntegrationDiscoverResponseErrorsCode = 102001
	CloudIntegrationDiscoverResponseErrorsCode102002 CloudIntegrationDiscoverResponseErrorsCode = 102002
	CloudIntegrationDiscoverResponseErrorsCode102003 CloudIntegrationDiscoverResponseErrorsCode = 102003
	CloudIntegrationDiscoverResponseErrorsCode102004 CloudIntegrationDiscoverResponseErrorsCode = 102004
	CloudIntegrationDiscoverResponseErrorsCode102005 CloudIntegrationDiscoverResponseErrorsCode = 102005
	CloudIntegrationDiscoverResponseErrorsCode102006 CloudIntegrationDiscoverResponseErrorsCode = 102006
	CloudIntegrationDiscoverResponseErrorsCode102007 CloudIntegrationDiscoverResponseErrorsCode = 102007
	CloudIntegrationDiscoverResponseErrorsCode102008 CloudIntegrationDiscoverResponseErrorsCode = 102008
	CloudIntegrationDiscoverResponseErrorsCode102009 CloudIntegrationDiscoverResponseErrorsCode = 102009
	CloudIntegrationDiscoverResponseErrorsCode102010 CloudIntegrationDiscoverResponseErrorsCode = 102010
	CloudIntegrationDiscoverResponseErrorsCode102011 CloudIntegrationDiscoverResponseErrorsCode = 102011
	CloudIntegrationDiscoverResponseErrorsCode102012 CloudIntegrationDiscoverResponseErrorsCode = 102012
	CloudIntegrationDiscoverResponseErrorsCode102013 CloudIntegrationDiscoverResponseErrorsCode = 102013
	CloudIntegrationDiscoverResponseErrorsCode102014 CloudIntegrationDiscoverResponseErrorsCode = 102014
	CloudIntegrationDiscoverResponseErrorsCode102015 CloudIntegrationDiscoverResponseErrorsCode = 102015
	CloudIntegrationDiscoverResponseErrorsCode102016 CloudIntegrationDiscoverResponseErrorsCode = 102016
	CloudIntegrationDiscoverResponseErrorsCode102017 CloudIntegrationDiscoverResponseErrorsCode = 102017
	CloudIntegrationDiscoverResponseErrorsCode102018 CloudIntegrationDiscoverResponseErrorsCode = 102018
	CloudIntegrationDiscoverResponseErrorsCode102019 CloudIntegrationDiscoverResponseErrorsCode = 102019
	CloudIntegrationDiscoverResponseErrorsCode102020 CloudIntegrationDiscoverResponseErrorsCode = 102020
	CloudIntegrationDiscoverResponseErrorsCode102021 CloudIntegrationDiscoverResponseErrorsCode = 102021
	CloudIntegrationDiscoverResponseErrorsCode102022 CloudIntegrationDiscoverResponseErrorsCode = 102022
	CloudIntegrationDiscoverResponseErrorsCode102023 CloudIntegrationDiscoverResponseErrorsCode = 102023
	CloudIntegrationDiscoverResponseErrorsCode102024 CloudIntegrationDiscoverResponseErrorsCode = 102024
	CloudIntegrationDiscoverResponseErrorsCode102025 CloudIntegrationDiscoverResponseErrorsCode = 102025
	CloudIntegrationDiscoverResponseErrorsCode102026 CloudIntegrationDiscoverResponseErrorsCode = 102026
	CloudIntegrationDiscoverResponseErrorsCode102027 CloudIntegrationDiscoverResponseErrorsCode = 102027
	CloudIntegrationDiscoverResponseErrorsCode102028 CloudIntegrationDiscoverResponseErrorsCode = 102028
	CloudIntegrationDiscoverResponseErrorsCode102029 CloudIntegrationDiscoverResponseErrorsCode = 102029
	CloudIntegrationDiscoverResponseErrorsCode102030 CloudIntegrationDiscoverResponseErrorsCode = 102030
	CloudIntegrationDiscoverResponseErrorsCode102031 CloudIntegrationDiscoverResponseErrorsCode = 102031
	CloudIntegrationDiscoverResponseErrorsCode102032 CloudIntegrationDiscoverResponseErrorsCode = 102032
	CloudIntegrationDiscoverResponseErrorsCode102033 CloudIntegrationDiscoverResponseErrorsCode = 102033
	CloudIntegrationDiscoverResponseErrorsCode102034 CloudIntegrationDiscoverResponseErrorsCode = 102034
	CloudIntegrationDiscoverResponseErrorsCode102035 CloudIntegrationDiscoverResponseErrorsCode = 102035
	CloudIntegrationDiscoverResponseErrorsCode102036 CloudIntegrationDiscoverResponseErrorsCode = 102036
	CloudIntegrationDiscoverResponseErrorsCode102037 CloudIntegrationDiscoverResponseErrorsCode = 102037
	CloudIntegrationDiscoverResponseErrorsCode102038 CloudIntegrationDiscoverResponseErrorsCode = 102038
	CloudIntegrationDiscoverResponseErrorsCode102039 CloudIntegrationDiscoverResponseErrorsCode = 102039
	CloudIntegrationDiscoverResponseErrorsCode102040 CloudIntegrationDiscoverResponseErrorsCode = 102040
	CloudIntegrationDiscoverResponseErrorsCode102041 CloudIntegrationDiscoverResponseErrorsCode = 102041
	CloudIntegrationDiscoverResponseErrorsCode102042 CloudIntegrationDiscoverResponseErrorsCode = 102042
	CloudIntegrationDiscoverResponseErrorsCode102043 CloudIntegrationDiscoverResponseErrorsCode = 102043
	CloudIntegrationDiscoverResponseErrorsCode102044 CloudIntegrationDiscoverResponseErrorsCode = 102044
	CloudIntegrationDiscoverResponseErrorsCode102045 CloudIntegrationDiscoverResponseErrorsCode = 102045
	CloudIntegrationDiscoverResponseErrorsCode102046 CloudIntegrationDiscoverResponseErrorsCode = 102046
	CloudIntegrationDiscoverResponseErrorsCode102047 CloudIntegrationDiscoverResponseErrorsCode = 102047
	CloudIntegrationDiscoverResponseErrorsCode102048 CloudIntegrationDiscoverResponseErrorsCode = 102048
	CloudIntegrationDiscoverResponseErrorsCode102049 CloudIntegrationDiscoverResponseErrorsCode = 102049
	CloudIntegrationDiscoverResponseErrorsCode102050 CloudIntegrationDiscoverResponseErrorsCode = 102050
	CloudIntegrationDiscoverResponseErrorsCode102051 CloudIntegrationDiscoverResponseErrorsCode = 102051
	CloudIntegrationDiscoverResponseErrorsCode102052 CloudIntegrationDiscoverResponseErrorsCode = 102052
	CloudIntegrationDiscoverResponseErrorsCode102053 CloudIntegrationDiscoverResponseErrorsCode = 102053
	CloudIntegrationDiscoverResponseErrorsCode102054 CloudIntegrationDiscoverResponseErrorsCode = 102054
	CloudIntegrationDiscoverResponseErrorsCode102055 CloudIntegrationDiscoverResponseErrorsCode = 102055
	CloudIntegrationDiscoverResponseErrorsCode102056 CloudIntegrationDiscoverResponseErrorsCode = 102056
	CloudIntegrationDiscoverResponseErrorsCode102057 CloudIntegrationDiscoverResponseErrorsCode = 102057
	CloudIntegrationDiscoverResponseErrorsCode102058 CloudIntegrationDiscoverResponseErrorsCode = 102058
	CloudIntegrationDiscoverResponseErrorsCode102059 CloudIntegrationDiscoverResponseErrorsCode = 102059
	CloudIntegrationDiscoverResponseErrorsCode102060 CloudIntegrationDiscoverResponseErrorsCode = 102060
	CloudIntegrationDiscoverResponseErrorsCode102061 CloudIntegrationDiscoverResponseErrorsCode = 102061
	CloudIntegrationDiscoverResponseErrorsCode102062 CloudIntegrationDiscoverResponseErrorsCode = 102062
	CloudIntegrationDiscoverResponseErrorsCode102063 CloudIntegrationDiscoverResponseErrorsCode = 102063
	CloudIntegrationDiscoverResponseErrorsCode102064 CloudIntegrationDiscoverResponseErrorsCode = 102064
	CloudIntegrationDiscoverResponseErrorsCode102065 CloudIntegrationDiscoverResponseErrorsCode = 102065
	CloudIntegrationDiscoverResponseErrorsCode102066 CloudIntegrationDiscoverResponseErrorsCode = 102066
	CloudIntegrationDiscoverResponseErrorsCode103001 CloudIntegrationDiscoverResponseErrorsCode = 103001
	CloudIntegrationDiscoverResponseErrorsCode103002 CloudIntegrationDiscoverResponseErrorsCode = 103002
	CloudIntegrationDiscoverResponseErrorsCode103003 CloudIntegrationDiscoverResponseErrorsCode = 103003
	CloudIntegrationDiscoverResponseErrorsCode103004 CloudIntegrationDiscoverResponseErrorsCode = 103004
	CloudIntegrationDiscoverResponseErrorsCode103005 CloudIntegrationDiscoverResponseErrorsCode = 103005
	CloudIntegrationDiscoverResponseErrorsCode103006 CloudIntegrationDiscoverResponseErrorsCode = 103006
	CloudIntegrationDiscoverResponseErrorsCode103007 CloudIntegrationDiscoverResponseErrorsCode = 103007
	CloudIntegrationDiscoverResponseErrorsCode103008 CloudIntegrationDiscoverResponseErrorsCode = 103008
)

func (r CloudIntegrationDiscoverResponseErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDiscoverResponseErrorsCode1001, CloudIntegrationDiscoverResponseErrorsCode1002, CloudIntegrationDiscoverResponseErrorsCode1003, CloudIntegrationDiscoverResponseErrorsCode1004, CloudIntegrationDiscoverResponseErrorsCode1005, CloudIntegrationDiscoverResponseErrorsCode1006, CloudIntegrationDiscoverResponseErrorsCode1007, CloudIntegrationDiscoverResponseErrorsCode1008, CloudIntegrationDiscoverResponseErrorsCode1009, CloudIntegrationDiscoverResponseErrorsCode1010, CloudIntegrationDiscoverResponseErrorsCode1011, CloudIntegrationDiscoverResponseErrorsCode1012, CloudIntegrationDiscoverResponseErrorsCode1013, CloudIntegrationDiscoverResponseErrorsCode1014, CloudIntegrationDiscoverResponseErrorsCode1015, CloudIntegrationDiscoverResponseErrorsCode1016, CloudIntegrationDiscoverResponseErrorsCode1017, CloudIntegrationDiscoverResponseErrorsCode2001, CloudIntegrationDiscoverResponseErrorsCode2002, CloudIntegrationDiscoverResponseErrorsCode2003, CloudIntegrationDiscoverResponseErrorsCode2004, CloudIntegrationDiscoverResponseErrorsCode2005, CloudIntegrationDiscoverResponseErrorsCode2006, CloudIntegrationDiscoverResponseErrorsCode2007, CloudIntegrationDiscoverResponseErrorsCode2008, CloudIntegrationDiscoverResponseErrorsCode2009, CloudIntegrationDiscoverResponseErrorsCode2010, CloudIntegrationDiscoverResponseErrorsCode2011, CloudIntegrationDiscoverResponseErrorsCode2012, CloudIntegrationDiscoverResponseErrorsCode2013, CloudIntegrationDiscoverResponseErrorsCode2014, CloudIntegrationDiscoverResponseErrorsCode2015, CloudIntegrationDiscoverResponseErrorsCode2016, CloudIntegrationDiscoverResponseErrorsCode2017, CloudIntegrationDiscoverResponseErrorsCode2018, CloudIntegrationDiscoverResponseErrorsCode2019, CloudIntegrationDiscoverResponseErrorsCode2020, CloudIntegrationDiscoverResponseErrorsCode2021, CloudIntegrationDiscoverResponseErrorsCode2022, CloudIntegrationDiscoverResponseErrorsCode3001, CloudIntegrationDiscoverResponseErrorsCode3002, CloudIntegrationDiscoverResponseErrorsCode3003, CloudIntegrationDiscoverResponseErrorsCode3004, CloudIntegrationDiscoverResponseErrorsCode3005, CloudIntegrationDiscoverResponseErrorsCode3006, CloudIntegrationDiscoverResponseErrorsCode3007, CloudIntegrationDiscoverResponseErrorsCode4001, CloudIntegrationDiscoverResponseErrorsCode4002, CloudIntegrationDiscoverResponseErrorsCode4003, CloudIntegrationDiscoverResponseErrorsCode4004, CloudIntegrationDiscoverResponseErrorsCode4005, CloudIntegrationDiscoverResponseErrorsCode4006, CloudIntegrationDiscoverResponseErrorsCode4007, CloudIntegrationDiscoverResponseErrorsCode4008, CloudIntegrationDiscoverResponseErrorsCode4009, CloudIntegrationDiscoverResponseErrorsCode4010, CloudIntegrationDiscoverResponseErrorsCode4011, CloudIntegrationDiscoverResponseErrorsCode4012, CloudIntegrationDiscoverResponseErrorsCode4013, CloudIntegrationDiscoverResponseErrorsCode4014, CloudIntegrationDiscoverResponseErrorsCode4015, CloudIntegrationDiscoverResponseErrorsCode4016, CloudIntegrationDiscoverResponseErrorsCode4017, CloudIntegrationDiscoverResponseErrorsCode4018, CloudIntegrationDiscoverResponseErrorsCode4019, CloudIntegrationDiscoverResponseErrorsCode4020, CloudIntegrationDiscoverResponseErrorsCode4021, CloudIntegrationDiscoverResponseErrorsCode4022, CloudIntegrationDiscoverResponseErrorsCode4023, CloudIntegrationDiscoverResponseErrorsCode5001, CloudIntegrationDiscoverResponseErrorsCode5002, CloudIntegrationDiscoverResponseErrorsCode5003, CloudIntegrationDiscoverResponseErrorsCode5004, CloudIntegrationDiscoverResponseErrorsCode102000, CloudIntegrationDiscoverResponseErrorsCode102001, CloudIntegrationDiscoverResponseErrorsCode102002, CloudIntegrationDiscoverResponseErrorsCode102003, CloudIntegrationDiscoverResponseErrorsCode102004, CloudIntegrationDiscoverResponseErrorsCode102005, CloudIntegrationDiscoverResponseErrorsCode102006, CloudIntegrationDiscoverResponseErrorsCode102007, CloudIntegrationDiscoverResponseErrorsCode102008, CloudIntegrationDiscoverResponseErrorsCode102009, CloudIntegrationDiscoverResponseErrorsCode102010, CloudIntegrationDiscoverResponseErrorsCode102011, CloudIntegrationDiscoverResponseErrorsCode102012, CloudIntegrationDiscoverResponseErrorsCode102013, CloudIntegrationDiscoverResponseErrorsCode102014, CloudIntegrationDiscoverResponseErrorsCode102015, CloudIntegrationDiscoverResponseErrorsCode102016, CloudIntegrationDiscoverResponseErrorsCode102017, CloudIntegrationDiscoverResponseErrorsCode102018, CloudIntegrationDiscoverResponseErrorsCode102019, CloudIntegrationDiscoverResponseErrorsCode102020, CloudIntegrationDiscoverResponseErrorsCode102021, CloudIntegrationDiscoverResponseErrorsCode102022, CloudIntegrationDiscoverResponseErrorsCode102023, CloudIntegrationDiscoverResponseErrorsCode102024, CloudIntegrationDiscoverResponseErrorsCode102025, CloudIntegrationDiscoverResponseErrorsCode102026, CloudIntegrationDiscoverResponseErrorsCode102027, CloudIntegrationDiscoverResponseErrorsCode102028, CloudIntegrationDiscoverResponseErrorsCode102029, CloudIntegrationDiscoverResponseErrorsCode102030, CloudIntegrationDiscoverResponseErrorsCode102031, CloudIntegrationDiscoverResponseErrorsCode102032, CloudIntegrationDiscoverResponseErrorsCode102033, CloudIntegrationDiscoverResponseErrorsCode102034, CloudIntegrationDiscoverResponseErrorsCode102035, CloudIntegrationDiscoverResponseErrorsCode102036, CloudIntegrationDiscoverResponseErrorsCode102037, CloudIntegrationDiscoverResponseErrorsCode102038, CloudIntegrationDiscoverResponseErrorsCode102039, CloudIntegrationDiscoverResponseErrorsCode102040, CloudIntegrationDiscoverResponseErrorsCode102041, CloudIntegrationDiscoverResponseErrorsCode102042, CloudIntegrationDiscoverResponseErrorsCode102043, CloudIntegrationDiscoverResponseErrorsCode102044, CloudIntegrationDiscoverResponseErrorsCode102045, CloudIntegrationDiscoverResponseErrorsCode102046, CloudIntegrationDiscoverResponseErrorsCode102047, CloudIntegrationDiscoverResponseErrorsCode102048, CloudIntegrationDiscoverResponseErrorsCode102049, CloudIntegrationDiscoverResponseErrorsCode102050, CloudIntegrationDiscoverResponseErrorsCode102051, CloudIntegrationDiscoverResponseErrorsCode102052, CloudIntegrationDiscoverResponseErrorsCode102053, CloudIntegrationDiscoverResponseErrorsCode102054, CloudIntegrationDiscoverResponseErrorsCode102055, CloudIntegrationDiscoverResponseErrorsCode102056, CloudIntegrationDiscoverResponseErrorsCode102057, CloudIntegrationDiscoverResponseErrorsCode102058, CloudIntegrationDiscoverResponseErrorsCode102059, CloudIntegrationDiscoverResponseErrorsCode102060, CloudIntegrationDiscoverResponseErrorsCode102061, CloudIntegrationDiscoverResponseErrorsCode102062, CloudIntegrationDiscoverResponseErrorsCode102063, CloudIntegrationDiscoverResponseErrorsCode102064, CloudIntegrationDiscoverResponseErrorsCode102065, CloudIntegrationDiscoverResponseErrorsCode102066, CloudIntegrationDiscoverResponseErrorsCode103001, CloudIntegrationDiscoverResponseErrorsCode103002, CloudIntegrationDiscoverResponseErrorsCode103003, CloudIntegrationDiscoverResponseErrorsCode103004, CloudIntegrationDiscoverResponseErrorsCode103005, CloudIntegrationDiscoverResponseErrorsCode103006, CloudIntegrationDiscoverResponseErrorsCode103007, CloudIntegrationDiscoverResponseErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationDiscoverResponseErrorsMeta struct {
	L10nKey       string                                         `json:"l10n_key"`
	LoggableError string                                         `json:"loggable_error"`
	TemplateData  interface{}                                    `json:"template_data"`
	TraceID       string                                         `json:"trace_id"`
	JSON          cloudIntegrationDiscoverResponseErrorsMetaJSON `json:"-"`
}

// cloudIntegrationDiscoverResponseErrorsMetaJSON contains the JSON metadata for
// the struct [CloudIntegrationDiscoverResponseErrorsMeta]
type cloudIntegrationDiscoverResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseErrorsSource struct {
	Parameter           string                                           `json:"parameter"`
	ParameterValueIndex int64                                            `json:"parameter_value_index"`
	Pointer             string                                           `json:"pointer"`
	JSON                cloudIntegrationDiscoverResponseErrorsSourceJSON `json:"-"`
}

// cloudIntegrationDiscoverResponseErrorsSourceJSON contains the JSON metadata for
// the struct [CloudIntegrationDiscoverResponseErrorsSource]
type cloudIntegrationDiscoverResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseMessage struct {
	Code             CloudIntegrationDiscoverResponseMessagesCode   `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Meta             CloudIntegrationDiscoverResponseMessagesMeta   `json:"meta"`
	Source           CloudIntegrationDiscoverResponseMessagesSource `json:"source"`
	JSON             cloudIntegrationDiscoverResponseMessageJSON    `json:"-"`
}

// cloudIntegrationDiscoverResponseMessageJSON contains the JSON metadata for the
// struct [CloudIntegrationDiscoverResponseMessage]
type cloudIntegrationDiscoverResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseMessageJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseMessagesCode int64

const (
	CloudIntegrationDiscoverResponseMessagesCode1001   CloudIntegrationDiscoverResponseMessagesCode = 1001
	CloudIntegrationDiscoverResponseMessagesCode1002   CloudIntegrationDiscoverResponseMessagesCode = 1002
	CloudIntegrationDiscoverResponseMessagesCode1003   CloudIntegrationDiscoverResponseMessagesCode = 1003
	CloudIntegrationDiscoverResponseMessagesCode1004   CloudIntegrationDiscoverResponseMessagesCode = 1004
	CloudIntegrationDiscoverResponseMessagesCode1005   CloudIntegrationDiscoverResponseMessagesCode = 1005
	CloudIntegrationDiscoverResponseMessagesCode1006   CloudIntegrationDiscoverResponseMessagesCode = 1006
	CloudIntegrationDiscoverResponseMessagesCode1007   CloudIntegrationDiscoverResponseMessagesCode = 1007
	CloudIntegrationDiscoverResponseMessagesCode1008   CloudIntegrationDiscoverResponseMessagesCode = 1008
	CloudIntegrationDiscoverResponseMessagesCode1009   CloudIntegrationDiscoverResponseMessagesCode = 1009
	CloudIntegrationDiscoverResponseMessagesCode1010   CloudIntegrationDiscoverResponseMessagesCode = 1010
	CloudIntegrationDiscoverResponseMessagesCode1011   CloudIntegrationDiscoverResponseMessagesCode = 1011
	CloudIntegrationDiscoverResponseMessagesCode1012   CloudIntegrationDiscoverResponseMessagesCode = 1012
	CloudIntegrationDiscoverResponseMessagesCode1013   CloudIntegrationDiscoverResponseMessagesCode = 1013
	CloudIntegrationDiscoverResponseMessagesCode1014   CloudIntegrationDiscoverResponseMessagesCode = 1014
	CloudIntegrationDiscoverResponseMessagesCode1015   CloudIntegrationDiscoverResponseMessagesCode = 1015
	CloudIntegrationDiscoverResponseMessagesCode1016   CloudIntegrationDiscoverResponseMessagesCode = 1016
	CloudIntegrationDiscoverResponseMessagesCode1017   CloudIntegrationDiscoverResponseMessagesCode = 1017
	CloudIntegrationDiscoverResponseMessagesCode2001   CloudIntegrationDiscoverResponseMessagesCode = 2001
	CloudIntegrationDiscoverResponseMessagesCode2002   CloudIntegrationDiscoverResponseMessagesCode = 2002
	CloudIntegrationDiscoverResponseMessagesCode2003   CloudIntegrationDiscoverResponseMessagesCode = 2003
	CloudIntegrationDiscoverResponseMessagesCode2004   CloudIntegrationDiscoverResponseMessagesCode = 2004
	CloudIntegrationDiscoverResponseMessagesCode2005   CloudIntegrationDiscoverResponseMessagesCode = 2005
	CloudIntegrationDiscoverResponseMessagesCode2006   CloudIntegrationDiscoverResponseMessagesCode = 2006
	CloudIntegrationDiscoverResponseMessagesCode2007   CloudIntegrationDiscoverResponseMessagesCode = 2007
	CloudIntegrationDiscoverResponseMessagesCode2008   CloudIntegrationDiscoverResponseMessagesCode = 2008
	CloudIntegrationDiscoverResponseMessagesCode2009   CloudIntegrationDiscoverResponseMessagesCode = 2009
	CloudIntegrationDiscoverResponseMessagesCode2010   CloudIntegrationDiscoverResponseMessagesCode = 2010
	CloudIntegrationDiscoverResponseMessagesCode2011   CloudIntegrationDiscoverResponseMessagesCode = 2011
	CloudIntegrationDiscoverResponseMessagesCode2012   CloudIntegrationDiscoverResponseMessagesCode = 2012
	CloudIntegrationDiscoverResponseMessagesCode2013   CloudIntegrationDiscoverResponseMessagesCode = 2013
	CloudIntegrationDiscoverResponseMessagesCode2014   CloudIntegrationDiscoverResponseMessagesCode = 2014
	CloudIntegrationDiscoverResponseMessagesCode2015   CloudIntegrationDiscoverResponseMessagesCode = 2015
	CloudIntegrationDiscoverResponseMessagesCode2016   CloudIntegrationDiscoverResponseMessagesCode = 2016
	CloudIntegrationDiscoverResponseMessagesCode2017   CloudIntegrationDiscoverResponseMessagesCode = 2017
	CloudIntegrationDiscoverResponseMessagesCode2018   CloudIntegrationDiscoverResponseMessagesCode = 2018
	CloudIntegrationDiscoverResponseMessagesCode2019   CloudIntegrationDiscoverResponseMessagesCode = 2019
	CloudIntegrationDiscoverResponseMessagesCode2020   CloudIntegrationDiscoverResponseMessagesCode = 2020
	CloudIntegrationDiscoverResponseMessagesCode2021   CloudIntegrationDiscoverResponseMessagesCode = 2021
	CloudIntegrationDiscoverResponseMessagesCode2022   CloudIntegrationDiscoverResponseMessagesCode = 2022
	CloudIntegrationDiscoverResponseMessagesCode3001   CloudIntegrationDiscoverResponseMessagesCode = 3001
	CloudIntegrationDiscoverResponseMessagesCode3002   CloudIntegrationDiscoverResponseMessagesCode = 3002
	CloudIntegrationDiscoverResponseMessagesCode3003   CloudIntegrationDiscoverResponseMessagesCode = 3003
	CloudIntegrationDiscoverResponseMessagesCode3004   CloudIntegrationDiscoverResponseMessagesCode = 3004
	CloudIntegrationDiscoverResponseMessagesCode3005   CloudIntegrationDiscoverResponseMessagesCode = 3005
	CloudIntegrationDiscoverResponseMessagesCode3006   CloudIntegrationDiscoverResponseMessagesCode = 3006
	CloudIntegrationDiscoverResponseMessagesCode3007   CloudIntegrationDiscoverResponseMessagesCode = 3007
	CloudIntegrationDiscoverResponseMessagesCode4001   CloudIntegrationDiscoverResponseMessagesCode = 4001
	CloudIntegrationDiscoverResponseMessagesCode4002   CloudIntegrationDiscoverResponseMessagesCode = 4002
	CloudIntegrationDiscoverResponseMessagesCode4003   CloudIntegrationDiscoverResponseMessagesCode = 4003
	CloudIntegrationDiscoverResponseMessagesCode4004   CloudIntegrationDiscoverResponseMessagesCode = 4004
	CloudIntegrationDiscoverResponseMessagesCode4005   CloudIntegrationDiscoverResponseMessagesCode = 4005
	CloudIntegrationDiscoverResponseMessagesCode4006   CloudIntegrationDiscoverResponseMessagesCode = 4006
	CloudIntegrationDiscoverResponseMessagesCode4007   CloudIntegrationDiscoverResponseMessagesCode = 4007
	CloudIntegrationDiscoverResponseMessagesCode4008   CloudIntegrationDiscoverResponseMessagesCode = 4008
	CloudIntegrationDiscoverResponseMessagesCode4009   CloudIntegrationDiscoverResponseMessagesCode = 4009
	CloudIntegrationDiscoverResponseMessagesCode4010   CloudIntegrationDiscoverResponseMessagesCode = 4010
	CloudIntegrationDiscoverResponseMessagesCode4011   CloudIntegrationDiscoverResponseMessagesCode = 4011
	CloudIntegrationDiscoverResponseMessagesCode4012   CloudIntegrationDiscoverResponseMessagesCode = 4012
	CloudIntegrationDiscoverResponseMessagesCode4013   CloudIntegrationDiscoverResponseMessagesCode = 4013
	CloudIntegrationDiscoverResponseMessagesCode4014   CloudIntegrationDiscoverResponseMessagesCode = 4014
	CloudIntegrationDiscoverResponseMessagesCode4015   CloudIntegrationDiscoverResponseMessagesCode = 4015
	CloudIntegrationDiscoverResponseMessagesCode4016   CloudIntegrationDiscoverResponseMessagesCode = 4016
	CloudIntegrationDiscoverResponseMessagesCode4017   CloudIntegrationDiscoverResponseMessagesCode = 4017
	CloudIntegrationDiscoverResponseMessagesCode4018   CloudIntegrationDiscoverResponseMessagesCode = 4018
	CloudIntegrationDiscoverResponseMessagesCode4019   CloudIntegrationDiscoverResponseMessagesCode = 4019
	CloudIntegrationDiscoverResponseMessagesCode4020   CloudIntegrationDiscoverResponseMessagesCode = 4020
	CloudIntegrationDiscoverResponseMessagesCode4021   CloudIntegrationDiscoverResponseMessagesCode = 4021
	CloudIntegrationDiscoverResponseMessagesCode4022   CloudIntegrationDiscoverResponseMessagesCode = 4022
	CloudIntegrationDiscoverResponseMessagesCode4023   CloudIntegrationDiscoverResponseMessagesCode = 4023
	CloudIntegrationDiscoverResponseMessagesCode5001   CloudIntegrationDiscoverResponseMessagesCode = 5001
	CloudIntegrationDiscoverResponseMessagesCode5002   CloudIntegrationDiscoverResponseMessagesCode = 5002
	CloudIntegrationDiscoverResponseMessagesCode5003   CloudIntegrationDiscoverResponseMessagesCode = 5003
	CloudIntegrationDiscoverResponseMessagesCode5004   CloudIntegrationDiscoverResponseMessagesCode = 5004
	CloudIntegrationDiscoverResponseMessagesCode102000 CloudIntegrationDiscoverResponseMessagesCode = 102000
	CloudIntegrationDiscoverResponseMessagesCode102001 CloudIntegrationDiscoverResponseMessagesCode = 102001
	CloudIntegrationDiscoverResponseMessagesCode102002 CloudIntegrationDiscoverResponseMessagesCode = 102002
	CloudIntegrationDiscoverResponseMessagesCode102003 CloudIntegrationDiscoverResponseMessagesCode = 102003
	CloudIntegrationDiscoverResponseMessagesCode102004 CloudIntegrationDiscoverResponseMessagesCode = 102004
	CloudIntegrationDiscoverResponseMessagesCode102005 CloudIntegrationDiscoverResponseMessagesCode = 102005
	CloudIntegrationDiscoverResponseMessagesCode102006 CloudIntegrationDiscoverResponseMessagesCode = 102006
	CloudIntegrationDiscoverResponseMessagesCode102007 CloudIntegrationDiscoverResponseMessagesCode = 102007
	CloudIntegrationDiscoverResponseMessagesCode102008 CloudIntegrationDiscoverResponseMessagesCode = 102008
	CloudIntegrationDiscoverResponseMessagesCode102009 CloudIntegrationDiscoverResponseMessagesCode = 102009
	CloudIntegrationDiscoverResponseMessagesCode102010 CloudIntegrationDiscoverResponseMessagesCode = 102010
	CloudIntegrationDiscoverResponseMessagesCode102011 CloudIntegrationDiscoverResponseMessagesCode = 102011
	CloudIntegrationDiscoverResponseMessagesCode102012 CloudIntegrationDiscoverResponseMessagesCode = 102012
	CloudIntegrationDiscoverResponseMessagesCode102013 CloudIntegrationDiscoverResponseMessagesCode = 102013
	CloudIntegrationDiscoverResponseMessagesCode102014 CloudIntegrationDiscoverResponseMessagesCode = 102014
	CloudIntegrationDiscoverResponseMessagesCode102015 CloudIntegrationDiscoverResponseMessagesCode = 102015
	CloudIntegrationDiscoverResponseMessagesCode102016 CloudIntegrationDiscoverResponseMessagesCode = 102016
	CloudIntegrationDiscoverResponseMessagesCode102017 CloudIntegrationDiscoverResponseMessagesCode = 102017
	CloudIntegrationDiscoverResponseMessagesCode102018 CloudIntegrationDiscoverResponseMessagesCode = 102018
	CloudIntegrationDiscoverResponseMessagesCode102019 CloudIntegrationDiscoverResponseMessagesCode = 102019
	CloudIntegrationDiscoverResponseMessagesCode102020 CloudIntegrationDiscoverResponseMessagesCode = 102020
	CloudIntegrationDiscoverResponseMessagesCode102021 CloudIntegrationDiscoverResponseMessagesCode = 102021
	CloudIntegrationDiscoverResponseMessagesCode102022 CloudIntegrationDiscoverResponseMessagesCode = 102022
	CloudIntegrationDiscoverResponseMessagesCode102023 CloudIntegrationDiscoverResponseMessagesCode = 102023
	CloudIntegrationDiscoverResponseMessagesCode102024 CloudIntegrationDiscoverResponseMessagesCode = 102024
	CloudIntegrationDiscoverResponseMessagesCode102025 CloudIntegrationDiscoverResponseMessagesCode = 102025
	CloudIntegrationDiscoverResponseMessagesCode102026 CloudIntegrationDiscoverResponseMessagesCode = 102026
	CloudIntegrationDiscoverResponseMessagesCode102027 CloudIntegrationDiscoverResponseMessagesCode = 102027
	CloudIntegrationDiscoverResponseMessagesCode102028 CloudIntegrationDiscoverResponseMessagesCode = 102028
	CloudIntegrationDiscoverResponseMessagesCode102029 CloudIntegrationDiscoverResponseMessagesCode = 102029
	CloudIntegrationDiscoverResponseMessagesCode102030 CloudIntegrationDiscoverResponseMessagesCode = 102030
	CloudIntegrationDiscoverResponseMessagesCode102031 CloudIntegrationDiscoverResponseMessagesCode = 102031
	CloudIntegrationDiscoverResponseMessagesCode102032 CloudIntegrationDiscoverResponseMessagesCode = 102032
	CloudIntegrationDiscoverResponseMessagesCode102033 CloudIntegrationDiscoverResponseMessagesCode = 102033
	CloudIntegrationDiscoverResponseMessagesCode102034 CloudIntegrationDiscoverResponseMessagesCode = 102034
	CloudIntegrationDiscoverResponseMessagesCode102035 CloudIntegrationDiscoverResponseMessagesCode = 102035
	CloudIntegrationDiscoverResponseMessagesCode102036 CloudIntegrationDiscoverResponseMessagesCode = 102036
	CloudIntegrationDiscoverResponseMessagesCode102037 CloudIntegrationDiscoverResponseMessagesCode = 102037
	CloudIntegrationDiscoverResponseMessagesCode102038 CloudIntegrationDiscoverResponseMessagesCode = 102038
	CloudIntegrationDiscoverResponseMessagesCode102039 CloudIntegrationDiscoverResponseMessagesCode = 102039
	CloudIntegrationDiscoverResponseMessagesCode102040 CloudIntegrationDiscoverResponseMessagesCode = 102040
	CloudIntegrationDiscoverResponseMessagesCode102041 CloudIntegrationDiscoverResponseMessagesCode = 102041
	CloudIntegrationDiscoverResponseMessagesCode102042 CloudIntegrationDiscoverResponseMessagesCode = 102042
	CloudIntegrationDiscoverResponseMessagesCode102043 CloudIntegrationDiscoverResponseMessagesCode = 102043
	CloudIntegrationDiscoverResponseMessagesCode102044 CloudIntegrationDiscoverResponseMessagesCode = 102044
	CloudIntegrationDiscoverResponseMessagesCode102045 CloudIntegrationDiscoverResponseMessagesCode = 102045
	CloudIntegrationDiscoverResponseMessagesCode102046 CloudIntegrationDiscoverResponseMessagesCode = 102046
	CloudIntegrationDiscoverResponseMessagesCode102047 CloudIntegrationDiscoverResponseMessagesCode = 102047
	CloudIntegrationDiscoverResponseMessagesCode102048 CloudIntegrationDiscoverResponseMessagesCode = 102048
	CloudIntegrationDiscoverResponseMessagesCode102049 CloudIntegrationDiscoverResponseMessagesCode = 102049
	CloudIntegrationDiscoverResponseMessagesCode102050 CloudIntegrationDiscoverResponseMessagesCode = 102050
	CloudIntegrationDiscoverResponseMessagesCode102051 CloudIntegrationDiscoverResponseMessagesCode = 102051
	CloudIntegrationDiscoverResponseMessagesCode102052 CloudIntegrationDiscoverResponseMessagesCode = 102052
	CloudIntegrationDiscoverResponseMessagesCode102053 CloudIntegrationDiscoverResponseMessagesCode = 102053
	CloudIntegrationDiscoverResponseMessagesCode102054 CloudIntegrationDiscoverResponseMessagesCode = 102054
	CloudIntegrationDiscoverResponseMessagesCode102055 CloudIntegrationDiscoverResponseMessagesCode = 102055
	CloudIntegrationDiscoverResponseMessagesCode102056 CloudIntegrationDiscoverResponseMessagesCode = 102056
	CloudIntegrationDiscoverResponseMessagesCode102057 CloudIntegrationDiscoverResponseMessagesCode = 102057
	CloudIntegrationDiscoverResponseMessagesCode102058 CloudIntegrationDiscoverResponseMessagesCode = 102058
	CloudIntegrationDiscoverResponseMessagesCode102059 CloudIntegrationDiscoverResponseMessagesCode = 102059
	CloudIntegrationDiscoverResponseMessagesCode102060 CloudIntegrationDiscoverResponseMessagesCode = 102060
	CloudIntegrationDiscoverResponseMessagesCode102061 CloudIntegrationDiscoverResponseMessagesCode = 102061
	CloudIntegrationDiscoverResponseMessagesCode102062 CloudIntegrationDiscoverResponseMessagesCode = 102062
	CloudIntegrationDiscoverResponseMessagesCode102063 CloudIntegrationDiscoverResponseMessagesCode = 102063
	CloudIntegrationDiscoverResponseMessagesCode102064 CloudIntegrationDiscoverResponseMessagesCode = 102064
	CloudIntegrationDiscoverResponseMessagesCode102065 CloudIntegrationDiscoverResponseMessagesCode = 102065
	CloudIntegrationDiscoverResponseMessagesCode102066 CloudIntegrationDiscoverResponseMessagesCode = 102066
	CloudIntegrationDiscoverResponseMessagesCode103001 CloudIntegrationDiscoverResponseMessagesCode = 103001
	CloudIntegrationDiscoverResponseMessagesCode103002 CloudIntegrationDiscoverResponseMessagesCode = 103002
	CloudIntegrationDiscoverResponseMessagesCode103003 CloudIntegrationDiscoverResponseMessagesCode = 103003
	CloudIntegrationDiscoverResponseMessagesCode103004 CloudIntegrationDiscoverResponseMessagesCode = 103004
	CloudIntegrationDiscoverResponseMessagesCode103005 CloudIntegrationDiscoverResponseMessagesCode = 103005
	CloudIntegrationDiscoverResponseMessagesCode103006 CloudIntegrationDiscoverResponseMessagesCode = 103006
	CloudIntegrationDiscoverResponseMessagesCode103007 CloudIntegrationDiscoverResponseMessagesCode = 103007
	CloudIntegrationDiscoverResponseMessagesCode103008 CloudIntegrationDiscoverResponseMessagesCode = 103008
)

func (r CloudIntegrationDiscoverResponseMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDiscoverResponseMessagesCode1001, CloudIntegrationDiscoverResponseMessagesCode1002, CloudIntegrationDiscoverResponseMessagesCode1003, CloudIntegrationDiscoverResponseMessagesCode1004, CloudIntegrationDiscoverResponseMessagesCode1005, CloudIntegrationDiscoverResponseMessagesCode1006, CloudIntegrationDiscoverResponseMessagesCode1007, CloudIntegrationDiscoverResponseMessagesCode1008, CloudIntegrationDiscoverResponseMessagesCode1009, CloudIntegrationDiscoverResponseMessagesCode1010, CloudIntegrationDiscoverResponseMessagesCode1011, CloudIntegrationDiscoverResponseMessagesCode1012, CloudIntegrationDiscoverResponseMessagesCode1013, CloudIntegrationDiscoverResponseMessagesCode1014, CloudIntegrationDiscoverResponseMessagesCode1015, CloudIntegrationDiscoverResponseMessagesCode1016, CloudIntegrationDiscoverResponseMessagesCode1017, CloudIntegrationDiscoverResponseMessagesCode2001, CloudIntegrationDiscoverResponseMessagesCode2002, CloudIntegrationDiscoverResponseMessagesCode2003, CloudIntegrationDiscoverResponseMessagesCode2004, CloudIntegrationDiscoverResponseMessagesCode2005, CloudIntegrationDiscoverResponseMessagesCode2006, CloudIntegrationDiscoverResponseMessagesCode2007, CloudIntegrationDiscoverResponseMessagesCode2008, CloudIntegrationDiscoverResponseMessagesCode2009, CloudIntegrationDiscoverResponseMessagesCode2010, CloudIntegrationDiscoverResponseMessagesCode2011, CloudIntegrationDiscoverResponseMessagesCode2012, CloudIntegrationDiscoverResponseMessagesCode2013, CloudIntegrationDiscoverResponseMessagesCode2014, CloudIntegrationDiscoverResponseMessagesCode2015, CloudIntegrationDiscoverResponseMessagesCode2016, CloudIntegrationDiscoverResponseMessagesCode2017, CloudIntegrationDiscoverResponseMessagesCode2018, CloudIntegrationDiscoverResponseMessagesCode2019, CloudIntegrationDiscoverResponseMessagesCode2020, CloudIntegrationDiscoverResponseMessagesCode2021, CloudIntegrationDiscoverResponseMessagesCode2022, CloudIntegrationDiscoverResponseMessagesCode3001, CloudIntegrationDiscoverResponseMessagesCode3002, CloudIntegrationDiscoverResponseMessagesCode3003, CloudIntegrationDiscoverResponseMessagesCode3004, CloudIntegrationDiscoverResponseMessagesCode3005, CloudIntegrationDiscoverResponseMessagesCode3006, CloudIntegrationDiscoverResponseMessagesCode3007, CloudIntegrationDiscoverResponseMessagesCode4001, CloudIntegrationDiscoverResponseMessagesCode4002, CloudIntegrationDiscoverResponseMessagesCode4003, CloudIntegrationDiscoverResponseMessagesCode4004, CloudIntegrationDiscoverResponseMessagesCode4005, CloudIntegrationDiscoverResponseMessagesCode4006, CloudIntegrationDiscoverResponseMessagesCode4007, CloudIntegrationDiscoverResponseMessagesCode4008, CloudIntegrationDiscoverResponseMessagesCode4009, CloudIntegrationDiscoverResponseMessagesCode4010, CloudIntegrationDiscoverResponseMessagesCode4011, CloudIntegrationDiscoverResponseMessagesCode4012, CloudIntegrationDiscoverResponseMessagesCode4013, CloudIntegrationDiscoverResponseMessagesCode4014, CloudIntegrationDiscoverResponseMessagesCode4015, CloudIntegrationDiscoverResponseMessagesCode4016, CloudIntegrationDiscoverResponseMessagesCode4017, CloudIntegrationDiscoverResponseMessagesCode4018, CloudIntegrationDiscoverResponseMessagesCode4019, CloudIntegrationDiscoverResponseMessagesCode4020, CloudIntegrationDiscoverResponseMessagesCode4021, CloudIntegrationDiscoverResponseMessagesCode4022, CloudIntegrationDiscoverResponseMessagesCode4023, CloudIntegrationDiscoverResponseMessagesCode5001, CloudIntegrationDiscoverResponseMessagesCode5002, CloudIntegrationDiscoverResponseMessagesCode5003, CloudIntegrationDiscoverResponseMessagesCode5004, CloudIntegrationDiscoverResponseMessagesCode102000, CloudIntegrationDiscoverResponseMessagesCode102001, CloudIntegrationDiscoverResponseMessagesCode102002, CloudIntegrationDiscoverResponseMessagesCode102003, CloudIntegrationDiscoverResponseMessagesCode102004, CloudIntegrationDiscoverResponseMessagesCode102005, CloudIntegrationDiscoverResponseMessagesCode102006, CloudIntegrationDiscoverResponseMessagesCode102007, CloudIntegrationDiscoverResponseMessagesCode102008, CloudIntegrationDiscoverResponseMessagesCode102009, CloudIntegrationDiscoverResponseMessagesCode102010, CloudIntegrationDiscoverResponseMessagesCode102011, CloudIntegrationDiscoverResponseMessagesCode102012, CloudIntegrationDiscoverResponseMessagesCode102013, CloudIntegrationDiscoverResponseMessagesCode102014, CloudIntegrationDiscoverResponseMessagesCode102015, CloudIntegrationDiscoverResponseMessagesCode102016, CloudIntegrationDiscoverResponseMessagesCode102017, CloudIntegrationDiscoverResponseMessagesCode102018, CloudIntegrationDiscoverResponseMessagesCode102019, CloudIntegrationDiscoverResponseMessagesCode102020, CloudIntegrationDiscoverResponseMessagesCode102021, CloudIntegrationDiscoverResponseMessagesCode102022, CloudIntegrationDiscoverResponseMessagesCode102023, CloudIntegrationDiscoverResponseMessagesCode102024, CloudIntegrationDiscoverResponseMessagesCode102025, CloudIntegrationDiscoverResponseMessagesCode102026, CloudIntegrationDiscoverResponseMessagesCode102027, CloudIntegrationDiscoverResponseMessagesCode102028, CloudIntegrationDiscoverResponseMessagesCode102029, CloudIntegrationDiscoverResponseMessagesCode102030, CloudIntegrationDiscoverResponseMessagesCode102031, CloudIntegrationDiscoverResponseMessagesCode102032, CloudIntegrationDiscoverResponseMessagesCode102033, CloudIntegrationDiscoverResponseMessagesCode102034, CloudIntegrationDiscoverResponseMessagesCode102035, CloudIntegrationDiscoverResponseMessagesCode102036, CloudIntegrationDiscoverResponseMessagesCode102037, CloudIntegrationDiscoverResponseMessagesCode102038, CloudIntegrationDiscoverResponseMessagesCode102039, CloudIntegrationDiscoverResponseMessagesCode102040, CloudIntegrationDiscoverResponseMessagesCode102041, CloudIntegrationDiscoverResponseMessagesCode102042, CloudIntegrationDiscoverResponseMessagesCode102043, CloudIntegrationDiscoverResponseMessagesCode102044, CloudIntegrationDiscoverResponseMessagesCode102045, CloudIntegrationDiscoverResponseMessagesCode102046, CloudIntegrationDiscoverResponseMessagesCode102047, CloudIntegrationDiscoverResponseMessagesCode102048, CloudIntegrationDiscoverResponseMessagesCode102049, CloudIntegrationDiscoverResponseMessagesCode102050, CloudIntegrationDiscoverResponseMessagesCode102051, CloudIntegrationDiscoverResponseMessagesCode102052, CloudIntegrationDiscoverResponseMessagesCode102053, CloudIntegrationDiscoverResponseMessagesCode102054, CloudIntegrationDiscoverResponseMessagesCode102055, CloudIntegrationDiscoverResponseMessagesCode102056, CloudIntegrationDiscoverResponseMessagesCode102057, CloudIntegrationDiscoverResponseMessagesCode102058, CloudIntegrationDiscoverResponseMessagesCode102059, CloudIntegrationDiscoverResponseMessagesCode102060, CloudIntegrationDiscoverResponseMessagesCode102061, CloudIntegrationDiscoverResponseMessagesCode102062, CloudIntegrationDiscoverResponseMessagesCode102063, CloudIntegrationDiscoverResponseMessagesCode102064, CloudIntegrationDiscoverResponseMessagesCode102065, CloudIntegrationDiscoverResponseMessagesCode102066, CloudIntegrationDiscoverResponseMessagesCode103001, CloudIntegrationDiscoverResponseMessagesCode103002, CloudIntegrationDiscoverResponseMessagesCode103003, CloudIntegrationDiscoverResponseMessagesCode103004, CloudIntegrationDiscoverResponseMessagesCode103005, CloudIntegrationDiscoverResponseMessagesCode103006, CloudIntegrationDiscoverResponseMessagesCode103007, CloudIntegrationDiscoverResponseMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationDiscoverResponseMessagesMeta struct {
	L10nKey       string                                           `json:"l10n_key"`
	LoggableError string                                           `json:"loggable_error"`
	TemplateData  interface{}                                      `json:"template_data"`
	TraceID       string                                           `json:"trace_id"`
	JSON          cloudIntegrationDiscoverResponseMessagesMetaJSON `json:"-"`
}

// cloudIntegrationDiscoverResponseMessagesMetaJSON contains the JSON metadata for
// the struct [CloudIntegrationDiscoverResponseMessagesMeta]
type cloudIntegrationDiscoverResponseMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverResponseMessagesSource struct {
	Parameter           string                                             `json:"parameter"`
	ParameterValueIndex int64                                              `json:"parameter_value_index"`
	Pointer             string                                             `json:"pointer"`
	JSON                cloudIntegrationDiscoverResponseMessagesSourceJSON `json:"-"`
}

// cloudIntegrationDiscoverResponseMessagesSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationDiscoverResponseMessagesSource]
type cloudIntegrationDiscoverResponseMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponse struct {
	Errors   []CloudIntegrationDiscoverAllResponseError   `json:"errors,required"`
	Messages []CloudIntegrationDiscoverAllResponseMessage `json:"messages,required"`
	Success  bool                                         `json:"success,required"`
	JSON     cloudIntegrationDiscoverAllResponseJSON      `json:"-"`
}

// cloudIntegrationDiscoverAllResponseJSON contains the JSON metadata for the
// struct [CloudIntegrationDiscoverAllResponse]
type cloudIntegrationDiscoverAllResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseError struct {
	Code             CloudIntegrationDiscoverAllResponseErrorsCode   `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Meta             CloudIntegrationDiscoverAllResponseErrorsMeta   `json:"meta"`
	Source           CloudIntegrationDiscoverAllResponseErrorsSource `json:"source"`
	JSON             cloudIntegrationDiscoverAllResponseErrorJSON    `json:"-"`
}

// cloudIntegrationDiscoverAllResponseErrorJSON contains the JSON metadata for the
// struct [CloudIntegrationDiscoverAllResponseError]
type cloudIntegrationDiscoverAllResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseErrorsCode int64

const (
	CloudIntegrationDiscoverAllResponseErrorsCode1001   CloudIntegrationDiscoverAllResponseErrorsCode = 1001
	CloudIntegrationDiscoverAllResponseErrorsCode1002   CloudIntegrationDiscoverAllResponseErrorsCode = 1002
	CloudIntegrationDiscoverAllResponseErrorsCode1003   CloudIntegrationDiscoverAllResponseErrorsCode = 1003
	CloudIntegrationDiscoverAllResponseErrorsCode1004   CloudIntegrationDiscoverAllResponseErrorsCode = 1004
	CloudIntegrationDiscoverAllResponseErrorsCode1005   CloudIntegrationDiscoverAllResponseErrorsCode = 1005
	CloudIntegrationDiscoverAllResponseErrorsCode1006   CloudIntegrationDiscoverAllResponseErrorsCode = 1006
	CloudIntegrationDiscoverAllResponseErrorsCode1007   CloudIntegrationDiscoverAllResponseErrorsCode = 1007
	CloudIntegrationDiscoverAllResponseErrorsCode1008   CloudIntegrationDiscoverAllResponseErrorsCode = 1008
	CloudIntegrationDiscoverAllResponseErrorsCode1009   CloudIntegrationDiscoverAllResponseErrorsCode = 1009
	CloudIntegrationDiscoverAllResponseErrorsCode1010   CloudIntegrationDiscoverAllResponseErrorsCode = 1010
	CloudIntegrationDiscoverAllResponseErrorsCode1011   CloudIntegrationDiscoverAllResponseErrorsCode = 1011
	CloudIntegrationDiscoverAllResponseErrorsCode1012   CloudIntegrationDiscoverAllResponseErrorsCode = 1012
	CloudIntegrationDiscoverAllResponseErrorsCode1013   CloudIntegrationDiscoverAllResponseErrorsCode = 1013
	CloudIntegrationDiscoverAllResponseErrorsCode1014   CloudIntegrationDiscoverAllResponseErrorsCode = 1014
	CloudIntegrationDiscoverAllResponseErrorsCode1015   CloudIntegrationDiscoverAllResponseErrorsCode = 1015
	CloudIntegrationDiscoverAllResponseErrorsCode1016   CloudIntegrationDiscoverAllResponseErrorsCode = 1016
	CloudIntegrationDiscoverAllResponseErrorsCode1017   CloudIntegrationDiscoverAllResponseErrorsCode = 1017
	CloudIntegrationDiscoverAllResponseErrorsCode2001   CloudIntegrationDiscoverAllResponseErrorsCode = 2001
	CloudIntegrationDiscoverAllResponseErrorsCode2002   CloudIntegrationDiscoverAllResponseErrorsCode = 2002
	CloudIntegrationDiscoverAllResponseErrorsCode2003   CloudIntegrationDiscoverAllResponseErrorsCode = 2003
	CloudIntegrationDiscoverAllResponseErrorsCode2004   CloudIntegrationDiscoverAllResponseErrorsCode = 2004
	CloudIntegrationDiscoverAllResponseErrorsCode2005   CloudIntegrationDiscoverAllResponseErrorsCode = 2005
	CloudIntegrationDiscoverAllResponseErrorsCode2006   CloudIntegrationDiscoverAllResponseErrorsCode = 2006
	CloudIntegrationDiscoverAllResponseErrorsCode2007   CloudIntegrationDiscoverAllResponseErrorsCode = 2007
	CloudIntegrationDiscoverAllResponseErrorsCode2008   CloudIntegrationDiscoverAllResponseErrorsCode = 2008
	CloudIntegrationDiscoverAllResponseErrorsCode2009   CloudIntegrationDiscoverAllResponseErrorsCode = 2009
	CloudIntegrationDiscoverAllResponseErrorsCode2010   CloudIntegrationDiscoverAllResponseErrorsCode = 2010
	CloudIntegrationDiscoverAllResponseErrorsCode2011   CloudIntegrationDiscoverAllResponseErrorsCode = 2011
	CloudIntegrationDiscoverAllResponseErrorsCode2012   CloudIntegrationDiscoverAllResponseErrorsCode = 2012
	CloudIntegrationDiscoverAllResponseErrorsCode2013   CloudIntegrationDiscoverAllResponseErrorsCode = 2013
	CloudIntegrationDiscoverAllResponseErrorsCode2014   CloudIntegrationDiscoverAllResponseErrorsCode = 2014
	CloudIntegrationDiscoverAllResponseErrorsCode2015   CloudIntegrationDiscoverAllResponseErrorsCode = 2015
	CloudIntegrationDiscoverAllResponseErrorsCode2016   CloudIntegrationDiscoverAllResponseErrorsCode = 2016
	CloudIntegrationDiscoverAllResponseErrorsCode2017   CloudIntegrationDiscoverAllResponseErrorsCode = 2017
	CloudIntegrationDiscoverAllResponseErrorsCode2018   CloudIntegrationDiscoverAllResponseErrorsCode = 2018
	CloudIntegrationDiscoverAllResponseErrorsCode2019   CloudIntegrationDiscoverAllResponseErrorsCode = 2019
	CloudIntegrationDiscoverAllResponseErrorsCode2020   CloudIntegrationDiscoverAllResponseErrorsCode = 2020
	CloudIntegrationDiscoverAllResponseErrorsCode2021   CloudIntegrationDiscoverAllResponseErrorsCode = 2021
	CloudIntegrationDiscoverAllResponseErrorsCode2022   CloudIntegrationDiscoverAllResponseErrorsCode = 2022
	CloudIntegrationDiscoverAllResponseErrorsCode3001   CloudIntegrationDiscoverAllResponseErrorsCode = 3001
	CloudIntegrationDiscoverAllResponseErrorsCode3002   CloudIntegrationDiscoverAllResponseErrorsCode = 3002
	CloudIntegrationDiscoverAllResponseErrorsCode3003   CloudIntegrationDiscoverAllResponseErrorsCode = 3003
	CloudIntegrationDiscoverAllResponseErrorsCode3004   CloudIntegrationDiscoverAllResponseErrorsCode = 3004
	CloudIntegrationDiscoverAllResponseErrorsCode3005   CloudIntegrationDiscoverAllResponseErrorsCode = 3005
	CloudIntegrationDiscoverAllResponseErrorsCode3006   CloudIntegrationDiscoverAllResponseErrorsCode = 3006
	CloudIntegrationDiscoverAllResponseErrorsCode3007   CloudIntegrationDiscoverAllResponseErrorsCode = 3007
	CloudIntegrationDiscoverAllResponseErrorsCode4001   CloudIntegrationDiscoverAllResponseErrorsCode = 4001
	CloudIntegrationDiscoverAllResponseErrorsCode4002   CloudIntegrationDiscoverAllResponseErrorsCode = 4002
	CloudIntegrationDiscoverAllResponseErrorsCode4003   CloudIntegrationDiscoverAllResponseErrorsCode = 4003
	CloudIntegrationDiscoverAllResponseErrorsCode4004   CloudIntegrationDiscoverAllResponseErrorsCode = 4004
	CloudIntegrationDiscoverAllResponseErrorsCode4005   CloudIntegrationDiscoverAllResponseErrorsCode = 4005
	CloudIntegrationDiscoverAllResponseErrorsCode4006   CloudIntegrationDiscoverAllResponseErrorsCode = 4006
	CloudIntegrationDiscoverAllResponseErrorsCode4007   CloudIntegrationDiscoverAllResponseErrorsCode = 4007
	CloudIntegrationDiscoverAllResponseErrorsCode4008   CloudIntegrationDiscoverAllResponseErrorsCode = 4008
	CloudIntegrationDiscoverAllResponseErrorsCode4009   CloudIntegrationDiscoverAllResponseErrorsCode = 4009
	CloudIntegrationDiscoverAllResponseErrorsCode4010   CloudIntegrationDiscoverAllResponseErrorsCode = 4010
	CloudIntegrationDiscoverAllResponseErrorsCode4011   CloudIntegrationDiscoverAllResponseErrorsCode = 4011
	CloudIntegrationDiscoverAllResponseErrorsCode4012   CloudIntegrationDiscoverAllResponseErrorsCode = 4012
	CloudIntegrationDiscoverAllResponseErrorsCode4013   CloudIntegrationDiscoverAllResponseErrorsCode = 4013
	CloudIntegrationDiscoverAllResponseErrorsCode4014   CloudIntegrationDiscoverAllResponseErrorsCode = 4014
	CloudIntegrationDiscoverAllResponseErrorsCode4015   CloudIntegrationDiscoverAllResponseErrorsCode = 4015
	CloudIntegrationDiscoverAllResponseErrorsCode4016   CloudIntegrationDiscoverAllResponseErrorsCode = 4016
	CloudIntegrationDiscoverAllResponseErrorsCode4017   CloudIntegrationDiscoverAllResponseErrorsCode = 4017
	CloudIntegrationDiscoverAllResponseErrorsCode4018   CloudIntegrationDiscoverAllResponseErrorsCode = 4018
	CloudIntegrationDiscoverAllResponseErrorsCode4019   CloudIntegrationDiscoverAllResponseErrorsCode = 4019
	CloudIntegrationDiscoverAllResponseErrorsCode4020   CloudIntegrationDiscoverAllResponseErrorsCode = 4020
	CloudIntegrationDiscoverAllResponseErrorsCode4021   CloudIntegrationDiscoverAllResponseErrorsCode = 4021
	CloudIntegrationDiscoverAllResponseErrorsCode4022   CloudIntegrationDiscoverAllResponseErrorsCode = 4022
	CloudIntegrationDiscoverAllResponseErrorsCode4023   CloudIntegrationDiscoverAllResponseErrorsCode = 4023
	CloudIntegrationDiscoverAllResponseErrorsCode5001   CloudIntegrationDiscoverAllResponseErrorsCode = 5001
	CloudIntegrationDiscoverAllResponseErrorsCode5002   CloudIntegrationDiscoverAllResponseErrorsCode = 5002
	CloudIntegrationDiscoverAllResponseErrorsCode5003   CloudIntegrationDiscoverAllResponseErrorsCode = 5003
	CloudIntegrationDiscoverAllResponseErrorsCode5004   CloudIntegrationDiscoverAllResponseErrorsCode = 5004
	CloudIntegrationDiscoverAllResponseErrorsCode102000 CloudIntegrationDiscoverAllResponseErrorsCode = 102000
	CloudIntegrationDiscoverAllResponseErrorsCode102001 CloudIntegrationDiscoverAllResponseErrorsCode = 102001
	CloudIntegrationDiscoverAllResponseErrorsCode102002 CloudIntegrationDiscoverAllResponseErrorsCode = 102002
	CloudIntegrationDiscoverAllResponseErrorsCode102003 CloudIntegrationDiscoverAllResponseErrorsCode = 102003
	CloudIntegrationDiscoverAllResponseErrorsCode102004 CloudIntegrationDiscoverAllResponseErrorsCode = 102004
	CloudIntegrationDiscoverAllResponseErrorsCode102005 CloudIntegrationDiscoverAllResponseErrorsCode = 102005
	CloudIntegrationDiscoverAllResponseErrorsCode102006 CloudIntegrationDiscoverAllResponseErrorsCode = 102006
	CloudIntegrationDiscoverAllResponseErrorsCode102007 CloudIntegrationDiscoverAllResponseErrorsCode = 102007
	CloudIntegrationDiscoverAllResponseErrorsCode102008 CloudIntegrationDiscoverAllResponseErrorsCode = 102008
	CloudIntegrationDiscoverAllResponseErrorsCode102009 CloudIntegrationDiscoverAllResponseErrorsCode = 102009
	CloudIntegrationDiscoverAllResponseErrorsCode102010 CloudIntegrationDiscoverAllResponseErrorsCode = 102010
	CloudIntegrationDiscoverAllResponseErrorsCode102011 CloudIntegrationDiscoverAllResponseErrorsCode = 102011
	CloudIntegrationDiscoverAllResponseErrorsCode102012 CloudIntegrationDiscoverAllResponseErrorsCode = 102012
	CloudIntegrationDiscoverAllResponseErrorsCode102013 CloudIntegrationDiscoverAllResponseErrorsCode = 102013
	CloudIntegrationDiscoverAllResponseErrorsCode102014 CloudIntegrationDiscoverAllResponseErrorsCode = 102014
	CloudIntegrationDiscoverAllResponseErrorsCode102015 CloudIntegrationDiscoverAllResponseErrorsCode = 102015
	CloudIntegrationDiscoverAllResponseErrorsCode102016 CloudIntegrationDiscoverAllResponseErrorsCode = 102016
	CloudIntegrationDiscoverAllResponseErrorsCode102017 CloudIntegrationDiscoverAllResponseErrorsCode = 102017
	CloudIntegrationDiscoverAllResponseErrorsCode102018 CloudIntegrationDiscoverAllResponseErrorsCode = 102018
	CloudIntegrationDiscoverAllResponseErrorsCode102019 CloudIntegrationDiscoverAllResponseErrorsCode = 102019
	CloudIntegrationDiscoverAllResponseErrorsCode102020 CloudIntegrationDiscoverAllResponseErrorsCode = 102020
	CloudIntegrationDiscoverAllResponseErrorsCode102021 CloudIntegrationDiscoverAllResponseErrorsCode = 102021
	CloudIntegrationDiscoverAllResponseErrorsCode102022 CloudIntegrationDiscoverAllResponseErrorsCode = 102022
	CloudIntegrationDiscoverAllResponseErrorsCode102023 CloudIntegrationDiscoverAllResponseErrorsCode = 102023
	CloudIntegrationDiscoverAllResponseErrorsCode102024 CloudIntegrationDiscoverAllResponseErrorsCode = 102024
	CloudIntegrationDiscoverAllResponseErrorsCode102025 CloudIntegrationDiscoverAllResponseErrorsCode = 102025
	CloudIntegrationDiscoverAllResponseErrorsCode102026 CloudIntegrationDiscoverAllResponseErrorsCode = 102026
	CloudIntegrationDiscoverAllResponseErrorsCode102027 CloudIntegrationDiscoverAllResponseErrorsCode = 102027
	CloudIntegrationDiscoverAllResponseErrorsCode102028 CloudIntegrationDiscoverAllResponseErrorsCode = 102028
	CloudIntegrationDiscoverAllResponseErrorsCode102029 CloudIntegrationDiscoverAllResponseErrorsCode = 102029
	CloudIntegrationDiscoverAllResponseErrorsCode102030 CloudIntegrationDiscoverAllResponseErrorsCode = 102030
	CloudIntegrationDiscoverAllResponseErrorsCode102031 CloudIntegrationDiscoverAllResponseErrorsCode = 102031
	CloudIntegrationDiscoverAllResponseErrorsCode102032 CloudIntegrationDiscoverAllResponseErrorsCode = 102032
	CloudIntegrationDiscoverAllResponseErrorsCode102033 CloudIntegrationDiscoverAllResponseErrorsCode = 102033
	CloudIntegrationDiscoverAllResponseErrorsCode102034 CloudIntegrationDiscoverAllResponseErrorsCode = 102034
	CloudIntegrationDiscoverAllResponseErrorsCode102035 CloudIntegrationDiscoverAllResponseErrorsCode = 102035
	CloudIntegrationDiscoverAllResponseErrorsCode102036 CloudIntegrationDiscoverAllResponseErrorsCode = 102036
	CloudIntegrationDiscoverAllResponseErrorsCode102037 CloudIntegrationDiscoverAllResponseErrorsCode = 102037
	CloudIntegrationDiscoverAllResponseErrorsCode102038 CloudIntegrationDiscoverAllResponseErrorsCode = 102038
	CloudIntegrationDiscoverAllResponseErrorsCode102039 CloudIntegrationDiscoverAllResponseErrorsCode = 102039
	CloudIntegrationDiscoverAllResponseErrorsCode102040 CloudIntegrationDiscoverAllResponseErrorsCode = 102040
	CloudIntegrationDiscoverAllResponseErrorsCode102041 CloudIntegrationDiscoverAllResponseErrorsCode = 102041
	CloudIntegrationDiscoverAllResponseErrorsCode102042 CloudIntegrationDiscoverAllResponseErrorsCode = 102042
	CloudIntegrationDiscoverAllResponseErrorsCode102043 CloudIntegrationDiscoverAllResponseErrorsCode = 102043
	CloudIntegrationDiscoverAllResponseErrorsCode102044 CloudIntegrationDiscoverAllResponseErrorsCode = 102044
	CloudIntegrationDiscoverAllResponseErrorsCode102045 CloudIntegrationDiscoverAllResponseErrorsCode = 102045
	CloudIntegrationDiscoverAllResponseErrorsCode102046 CloudIntegrationDiscoverAllResponseErrorsCode = 102046
	CloudIntegrationDiscoverAllResponseErrorsCode102047 CloudIntegrationDiscoverAllResponseErrorsCode = 102047
	CloudIntegrationDiscoverAllResponseErrorsCode102048 CloudIntegrationDiscoverAllResponseErrorsCode = 102048
	CloudIntegrationDiscoverAllResponseErrorsCode102049 CloudIntegrationDiscoverAllResponseErrorsCode = 102049
	CloudIntegrationDiscoverAllResponseErrorsCode102050 CloudIntegrationDiscoverAllResponseErrorsCode = 102050
	CloudIntegrationDiscoverAllResponseErrorsCode102051 CloudIntegrationDiscoverAllResponseErrorsCode = 102051
	CloudIntegrationDiscoverAllResponseErrorsCode102052 CloudIntegrationDiscoverAllResponseErrorsCode = 102052
	CloudIntegrationDiscoverAllResponseErrorsCode102053 CloudIntegrationDiscoverAllResponseErrorsCode = 102053
	CloudIntegrationDiscoverAllResponseErrorsCode102054 CloudIntegrationDiscoverAllResponseErrorsCode = 102054
	CloudIntegrationDiscoverAllResponseErrorsCode102055 CloudIntegrationDiscoverAllResponseErrorsCode = 102055
	CloudIntegrationDiscoverAllResponseErrorsCode102056 CloudIntegrationDiscoverAllResponseErrorsCode = 102056
	CloudIntegrationDiscoverAllResponseErrorsCode102057 CloudIntegrationDiscoverAllResponseErrorsCode = 102057
	CloudIntegrationDiscoverAllResponseErrorsCode102058 CloudIntegrationDiscoverAllResponseErrorsCode = 102058
	CloudIntegrationDiscoverAllResponseErrorsCode102059 CloudIntegrationDiscoverAllResponseErrorsCode = 102059
	CloudIntegrationDiscoverAllResponseErrorsCode102060 CloudIntegrationDiscoverAllResponseErrorsCode = 102060
	CloudIntegrationDiscoverAllResponseErrorsCode102061 CloudIntegrationDiscoverAllResponseErrorsCode = 102061
	CloudIntegrationDiscoverAllResponseErrorsCode102062 CloudIntegrationDiscoverAllResponseErrorsCode = 102062
	CloudIntegrationDiscoverAllResponseErrorsCode102063 CloudIntegrationDiscoverAllResponseErrorsCode = 102063
	CloudIntegrationDiscoverAllResponseErrorsCode102064 CloudIntegrationDiscoverAllResponseErrorsCode = 102064
	CloudIntegrationDiscoverAllResponseErrorsCode102065 CloudIntegrationDiscoverAllResponseErrorsCode = 102065
	CloudIntegrationDiscoverAllResponseErrorsCode102066 CloudIntegrationDiscoverAllResponseErrorsCode = 102066
	CloudIntegrationDiscoverAllResponseErrorsCode103001 CloudIntegrationDiscoverAllResponseErrorsCode = 103001
	CloudIntegrationDiscoverAllResponseErrorsCode103002 CloudIntegrationDiscoverAllResponseErrorsCode = 103002
	CloudIntegrationDiscoverAllResponseErrorsCode103003 CloudIntegrationDiscoverAllResponseErrorsCode = 103003
	CloudIntegrationDiscoverAllResponseErrorsCode103004 CloudIntegrationDiscoverAllResponseErrorsCode = 103004
	CloudIntegrationDiscoverAllResponseErrorsCode103005 CloudIntegrationDiscoverAllResponseErrorsCode = 103005
	CloudIntegrationDiscoverAllResponseErrorsCode103006 CloudIntegrationDiscoverAllResponseErrorsCode = 103006
	CloudIntegrationDiscoverAllResponseErrorsCode103007 CloudIntegrationDiscoverAllResponseErrorsCode = 103007
	CloudIntegrationDiscoverAllResponseErrorsCode103008 CloudIntegrationDiscoverAllResponseErrorsCode = 103008
)

func (r CloudIntegrationDiscoverAllResponseErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDiscoverAllResponseErrorsCode1001, CloudIntegrationDiscoverAllResponseErrorsCode1002, CloudIntegrationDiscoverAllResponseErrorsCode1003, CloudIntegrationDiscoverAllResponseErrorsCode1004, CloudIntegrationDiscoverAllResponseErrorsCode1005, CloudIntegrationDiscoverAllResponseErrorsCode1006, CloudIntegrationDiscoverAllResponseErrorsCode1007, CloudIntegrationDiscoverAllResponseErrorsCode1008, CloudIntegrationDiscoverAllResponseErrorsCode1009, CloudIntegrationDiscoverAllResponseErrorsCode1010, CloudIntegrationDiscoverAllResponseErrorsCode1011, CloudIntegrationDiscoverAllResponseErrorsCode1012, CloudIntegrationDiscoverAllResponseErrorsCode1013, CloudIntegrationDiscoverAllResponseErrorsCode1014, CloudIntegrationDiscoverAllResponseErrorsCode1015, CloudIntegrationDiscoverAllResponseErrorsCode1016, CloudIntegrationDiscoverAllResponseErrorsCode1017, CloudIntegrationDiscoverAllResponseErrorsCode2001, CloudIntegrationDiscoverAllResponseErrorsCode2002, CloudIntegrationDiscoverAllResponseErrorsCode2003, CloudIntegrationDiscoverAllResponseErrorsCode2004, CloudIntegrationDiscoverAllResponseErrorsCode2005, CloudIntegrationDiscoverAllResponseErrorsCode2006, CloudIntegrationDiscoverAllResponseErrorsCode2007, CloudIntegrationDiscoverAllResponseErrorsCode2008, CloudIntegrationDiscoverAllResponseErrorsCode2009, CloudIntegrationDiscoverAllResponseErrorsCode2010, CloudIntegrationDiscoverAllResponseErrorsCode2011, CloudIntegrationDiscoverAllResponseErrorsCode2012, CloudIntegrationDiscoverAllResponseErrorsCode2013, CloudIntegrationDiscoverAllResponseErrorsCode2014, CloudIntegrationDiscoverAllResponseErrorsCode2015, CloudIntegrationDiscoverAllResponseErrorsCode2016, CloudIntegrationDiscoverAllResponseErrorsCode2017, CloudIntegrationDiscoverAllResponseErrorsCode2018, CloudIntegrationDiscoverAllResponseErrorsCode2019, CloudIntegrationDiscoverAllResponseErrorsCode2020, CloudIntegrationDiscoverAllResponseErrorsCode2021, CloudIntegrationDiscoverAllResponseErrorsCode2022, CloudIntegrationDiscoverAllResponseErrorsCode3001, CloudIntegrationDiscoverAllResponseErrorsCode3002, CloudIntegrationDiscoverAllResponseErrorsCode3003, CloudIntegrationDiscoverAllResponseErrorsCode3004, CloudIntegrationDiscoverAllResponseErrorsCode3005, CloudIntegrationDiscoverAllResponseErrorsCode3006, CloudIntegrationDiscoverAllResponseErrorsCode3007, CloudIntegrationDiscoverAllResponseErrorsCode4001, CloudIntegrationDiscoverAllResponseErrorsCode4002, CloudIntegrationDiscoverAllResponseErrorsCode4003, CloudIntegrationDiscoverAllResponseErrorsCode4004, CloudIntegrationDiscoverAllResponseErrorsCode4005, CloudIntegrationDiscoverAllResponseErrorsCode4006, CloudIntegrationDiscoverAllResponseErrorsCode4007, CloudIntegrationDiscoverAllResponseErrorsCode4008, CloudIntegrationDiscoverAllResponseErrorsCode4009, CloudIntegrationDiscoverAllResponseErrorsCode4010, CloudIntegrationDiscoverAllResponseErrorsCode4011, CloudIntegrationDiscoverAllResponseErrorsCode4012, CloudIntegrationDiscoverAllResponseErrorsCode4013, CloudIntegrationDiscoverAllResponseErrorsCode4014, CloudIntegrationDiscoverAllResponseErrorsCode4015, CloudIntegrationDiscoverAllResponseErrorsCode4016, CloudIntegrationDiscoverAllResponseErrorsCode4017, CloudIntegrationDiscoverAllResponseErrorsCode4018, CloudIntegrationDiscoverAllResponseErrorsCode4019, CloudIntegrationDiscoverAllResponseErrorsCode4020, CloudIntegrationDiscoverAllResponseErrorsCode4021, CloudIntegrationDiscoverAllResponseErrorsCode4022, CloudIntegrationDiscoverAllResponseErrorsCode4023, CloudIntegrationDiscoverAllResponseErrorsCode5001, CloudIntegrationDiscoverAllResponseErrorsCode5002, CloudIntegrationDiscoverAllResponseErrorsCode5003, CloudIntegrationDiscoverAllResponseErrorsCode5004, CloudIntegrationDiscoverAllResponseErrorsCode102000, CloudIntegrationDiscoverAllResponseErrorsCode102001, CloudIntegrationDiscoverAllResponseErrorsCode102002, CloudIntegrationDiscoverAllResponseErrorsCode102003, CloudIntegrationDiscoverAllResponseErrorsCode102004, CloudIntegrationDiscoverAllResponseErrorsCode102005, CloudIntegrationDiscoverAllResponseErrorsCode102006, CloudIntegrationDiscoverAllResponseErrorsCode102007, CloudIntegrationDiscoverAllResponseErrorsCode102008, CloudIntegrationDiscoverAllResponseErrorsCode102009, CloudIntegrationDiscoverAllResponseErrorsCode102010, CloudIntegrationDiscoverAllResponseErrorsCode102011, CloudIntegrationDiscoverAllResponseErrorsCode102012, CloudIntegrationDiscoverAllResponseErrorsCode102013, CloudIntegrationDiscoverAllResponseErrorsCode102014, CloudIntegrationDiscoverAllResponseErrorsCode102015, CloudIntegrationDiscoverAllResponseErrorsCode102016, CloudIntegrationDiscoverAllResponseErrorsCode102017, CloudIntegrationDiscoverAllResponseErrorsCode102018, CloudIntegrationDiscoverAllResponseErrorsCode102019, CloudIntegrationDiscoverAllResponseErrorsCode102020, CloudIntegrationDiscoverAllResponseErrorsCode102021, CloudIntegrationDiscoverAllResponseErrorsCode102022, CloudIntegrationDiscoverAllResponseErrorsCode102023, CloudIntegrationDiscoverAllResponseErrorsCode102024, CloudIntegrationDiscoverAllResponseErrorsCode102025, CloudIntegrationDiscoverAllResponseErrorsCode102026, CloudIntegrationDiscoverAllResponseErrorsCode102027, CloudIntegrationDiscoverAllResponseErrorsCode102028, CloudIntegrationDiscoverAllResponseErrorsCode102029, CloudIntegrationDiscoverAllResponseErrorsCode102030, CloudIntegrationDiscoverAllResponseErrorsCode102031, CloudIntegrationDiscoverAllResponseErrorsCode102032, CloudIntegrationDiscoverAllResponseErrorsCode102033, CloudIntegrationDiscoverAllResponseErrorsCode102034, CloudIntegrationDiscoverAllResponseErrorsCode102035, CloudIntegrationDiscoverAllResponseErrorsCode102036, CloudIntegrationDiscoverAllResponseErrorsCode102037, CloudIntegrationDiscoverAllResponseErrorsCode102038, CloudIntegrationDiscoverAllResponseErrorsCode102039, CloudIntegrationDiscoverAllResponseErrorsCode102040, CloudIntegrationDiscoverAllResponseErrorsCode102041, CloudIntegrationDiscoverAllResponseErrorsCode102042, CloudIntegrationDiscoverAllResponseErrorsCode102043, CloudIntegrationDiscoverAllResponseErrorsCode102044, CloudIntegrationDiscoverAllResponseErrorsCode102045, CloudIntegrationDiscoverAllResponseErrorsCode102046, CloudIntegrationDiscoverAllResponseErrorsCode102047, CloudIntegrationDiscoverAllResponseErrorsCode102048, CloudIntegrationDiscoverAllResponseErrorsCode102049, CloudIntegrationDiscoverAllResponseErrorsCode102050, CloudIntegrationDiscoverAllResponseErrorsCode102051, CloudIntegrationDiscoverAllResponseErrorsCode102052, CloudIntegrationDiscoverAllResponseErrorsCode102053, CloudIntegrationDiscoverAllResponseErrorsCode102054, CloudIntegrationDiscoverAllResponseErrorsCode102055, CloudIntegrationDiscoverAllResponseErrorsCode102056, CloudIntegrationDiscoverAllResponseErrorsCode102057, CloudIntegrationDiscoverAllResponseErrorsCode102058, CloudIntegrationDiscoverAllResponseErrorsCode102059, CloudIntegrationDiscoverAllResponseErrorsCode102060, CloudIntegrationDiscoverAllResponseErrorsCode102061, CloudIntegrationDiscoverAllResponseErrorsCode102062, CloudIntegrationDiscoverAllResponseErrorsCode102063, CloudIntegrationDiscoverAllResponseErrorsCode102064, CloudIntegrationDiscoverAllResponseErrorsCode102065, CloudIntegrationDiscoverAllResponseErrorsCode102066, CloudIntegrationDiscoverAllResponseErrorsCode103001, CloudIntegrationDiscoverAllResponseErrorsCode103002, CloudIntegrationDiscoverAllResponseErrorsCode103003, CloudIntegrationDiscoverAllResponseErrorsCode103004, CloudIntegrationDiscoverAllResponseErrorsCode103005, CloudIntegrationDiscoverAllResponseErrorsCode103006, CloudIntegrationDiscoverAllResponseErrorsCode103007, CloudIntegrationDiscoverAllResponseErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationDiscoverAllResponseErrorsMeta struct {
	L10nKey       string                                            `json:"l10n_key"`
	LoggableError string                                            `json:"loggable_error"`
	TemplateData  interface{}                                       `json:"template_data"`
	TraceID       string                                            `json:"trace_id"`
	JSON          cloudIntegrationDiscoverAllResponseErrorsMetaJSON `json:"-"`
}

// cloudIntegrationDiscoverAllResponseErrorsMetaJSON contains the JSON metadata for
// the struct [CloudIntegrationDiscoverAllResponseErrorsMeta]
type cloudIntegrationDiscoverAllResponseErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseErrorsSource struct {
	Parameter           string                                              `json:"parameter"`
	ParameterValueIndex int64                                               `json:"parameter_value_index"`
	Pointer             string                                              `json:"pointer"`
	JSON                cloudIntegrationDiscoverAllResponseErrorsSourceJSON `json:"-"`
}

// cloudIntegrationDiscoverAllResponseErrorsSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationDiscoverAllResponseErrorsSource]
type cloudIntegrationDiscoverAllResponseErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseMessage struct {
	Code             CloudIntegrationDiscoverAllResponseMessagesCode   `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Meta             CloudIntegrationDiscoverAllResponseMessagesMeta   `json:"meta"`
	Source           CloudIntegrationDiscoverAllResponseMessagesSource `json:"source"`
	JSON             cloudIntegrationDiscoverAllResponseMessageJSON    `json:"-"`
}

// cloudIntegrationDiscoverAllResponseMessageJSON contains the JSON metadata for
// the struct [CloudIntegrationDiscoverAllResponseMessage]
type cloudIntegrationDiscoverAllResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseMessageJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseMessagesCode int64

const (
	CloudIntegrationDiscoverAllResponseMessagesCode1001   CloudIntegrationDiscoverAllResponseMessagesCode = 1001
	CloudIntegrationDiscoverAllResponseMessagesCode1002   CloudIntegrationDiscoverAllResponseMessagesCode = 1002
	CloudIntegrationDiscoverAllResponseMessagesCode1003   CloudIntegrationDiscoverAllResponseMessagesCode = 1003
	CloudIntegrationDiscoverAllResponseMessagesCode1004   CloudIntegrationDiscoverAllResponseMessagesCode = 1004
	CloudIntegrationDiscoverAllResponseMessagesCode1005   CloudIntegrationDiscoverAllResponseMessagesCode = 1005
	CloudIntegrationDiscoverAllResponseMessagesCode1006   CloudIntegrationDiscoverAllResponseMessagesCode = 1006
	CloudIntegrationDiscoverAllResponseMessagesCode1007   CloudIntegrationDiscoverAllResponseMessagesCode = 1007
	CloudIntegrationDiscoverAllResponseMessagesCode1008   CloudIntegrationDiscoverAllResponseMessagesCode = 1008
	CloudIntegrationDiscoverAllResponseMessagesCode1009   CloudIntegrationDiscoverAllResponseMessagesCode = 1009
	CloudIntegrationDiscoverAllResponseMessagesCode1010   CloudIntegrationDiscoverAllResponseMessagesCode = 1010
	CloudIntegrationDiscoverAllResponseMessagesCode1011   CloudIntegrationDiscoverAllResponseMessagesCode = 1011
	CloudIntegrationDiscoverAllResponseMessagesCode1012   CloudIntegrationDiscoverAllResponseMessagesCode = 1012
	CloudIntegrationDiscoverAllResponseMessagesCode1013   CloudIntegrationDiscoverAllResponseMessagesCode = 1013
	CloudIntegrationDiscoverAllResponseMessagesCode1014   CloudIntegrationDiscoverAllResponseMessagesCode = 1014
	CloudIntegrationDiscoverAllResponseMessagesCode1015   CloudIntegrationDiscoverAllResponseMessagesCode = 1015
	CloudIntegrationDiscoverAllResponseMessagesCode1016   CloudIntegrationDiscoverAllResponseMessagesCode = 1016
	CloudIntegrationDiscoverAllResponseMessagesCode1017   CloudIntegrationDiscoverAllResponseMessagesCode = 1017
	CloudIntegrationDiscoverAllResponseMessagesCode2001   CloudIntegrationDiscoverAllResponseMessagesCode = 2001
	CloudIntegrationDiscoverAllResponseMessagesCode2002   CloudIntegrationDiscoverAllResponseMessagesCode = 2002
	CloudIntegrationDiscoverAllResponseMessagesCode2003   CloudIntegrationDiscoverAllResponseMessagesCode = 2003
	CloudIntegrationDiscoverAllResponseMessagesCode2004   CloudIntegrationDiscoverAllResponseMessagesCode = 2004
	CloudIntegrationDiscoverAllResponseMessagesCode2005   CloudIntegrationDiscoverAllResponseMessagesCode = 2005
	CloudIntegrationDiscoverAllResponseMessagesCode2006   CloudIntegrationDiscoverAllResponseMessagesCode = 2006
	CloudIntegrationDiscoverAllResponseMessagesCode2007   CloudIntegrationDiscoverAllResponseMessagesCode = 2007
	CloudIntegrationDiscoverAllResponseMessagesCode2008   CloudIntegrationDiscoverAllResponseMessagesCode = 2008
	CloudIntegrationDiscoverAllResponseMessagesCode2009   CloudIntegrationDiscoverAllResponseMessagesCode = 2009
	CloudIntegrationDiscoverAllResponseMessagesCode2010   CloudIntegrationDiscoverAllResponseMessagesCode = 2010
	CloudIntegrationDiscoverAllResponseMessagesCode2011   CloudIntegrationDiscoverAllResponseMessagesCode = 2011
	CloudIntegrationDiscoverAllResponseMessagesCode2012   CloudIntegrationDiscoverAllResponseMessagesCode = 2012
	CloudIntegrationDiscoverAllResponseMessagesCode2013   CloudIntegrationDiscoverAllResponseMessagesCode = 2013
	CloudIntegrationDiscoverAllResponseMessagesCode2014   CloudIntegrationDiscoverAllResponseMessagesCode = 2014
	CloudIntegrationDiscoverAllResponseMessagesCode2015   CloudIntegrationDiscoverAllResponseMessagesCode = 2015
	CloudIntegrationDiscoverAllResponseMessagesCode2016   CloudIntegrationDiscoverAllResponseMessagesCode = 2016
	CloudIntegrationDiscoverAllResponseMessagesCode2017   CloudIntegrationDiscoverAllResponseMessagesCode = 2017
	CloudIntegrationDiscoverAllResponseMessagesCode2018   CloudIntegrationDiscoverAllResponseMessagesCode = 2018
	CloudIntegrationDiscoverAllResponseMessagesCode2019   CloudIntegrationDiscoverAllResponseMessagesCode = 2019
	CloudIntegrationDiscoverAllResponseMessagesCode2020   CloudIntegrationDiscoverAllResponseMessagesCode = 2020
	CloudIntegrationDiscoverAllResponseMessagesCode2021   CloudIntegrationDiscoverAllResponseMessagesCode = 2021
	CloudIntegrationDiscoverAllResponseMessagesCode2022   CloudIntegrationDiscoverAllResponseMessagesCode = 2022
	CloudIntegrationDiscoverAllResponseMessagesCode3001   CloudIntegrationDiscoverAllResponseMessagesCode = 3001
	CloudIntegrationDiscoverAllResponseMessagesCode3002   CloudIntegrationDiscoverAllResponseMessagesCode = 3002
	CloudIntegrationDiscoverAllResponseMessagesCode3003   CloudIntegrationDiscoverAllResponseMessagesCode = 3003
	CloudIntegrationDiscoverAllResponseMessagesCode3004   CloudIntegrationDiscoverAllResponseMessagesCode = 3004
	CloudIntegrationDiscoverAllResponseMessagesCode3005   CloudIntegrationDiscoverAllResponseMessagesCode = 3005
	CloudIntegrationDiscoverAllResponseMessagesCode3006   CloudIntegrationDiscoverAllResponseMessagesCode = 3006
	CloudIntegrationDiscoverAllResponseMessagesCode3007   CloudIntegrationDiscoverAllResponseMessagesCode = 3007
	CloudIntegrationDiscoverAllResponseMessagesCode4001   CloudIntegrationDiscoverAllResponseMessagesCode = 4001
	CloudIntegrationDiscoverAllResponseMessagesCode4002   CloudIntegrationDiscoverAllResponseMessagesCode = 4002
	CloudIntegrationDiscoverAllResponseMessagesCode4003   CloudIntegrationDiscoverAllResponseMessagesCode = 4003
	CloudIntegrationDiscoverAllResponseMessagesCode4004   CloudIntegrationDiscoverAllResponseMessagesCode = 4004
	CloudIntegrationDiscoverAllResponseMessagesCode4005   CloudIntegrationDiscoverAllResponseMessagesCode = 4005
	CloudIntegrationDiscoverAllResponseMessagesCode4006   CloudIntegrationDiscoverAllResponseMessagesCode = 4006
	CloudIntegrationDiscoverAllResponseMessagesCode4007   CloudIntegrationDiscoverAllResponseMessagesCode = 4007
	CloudIntegrationDiscoverAllResponseMessagesCode4008   CloudIntegrationDiscoverAllResponseMessagesCode = 4008
	CloudIntegrationDiscoverAllResponseMessagesCode4009   CloudIntegrationDiscoverAllResponseMessagesCode = 4009
	CloudIntegrationDiscoverAllResponseMessagesCode4010   CloudIntegrationDiscoverAllResponseMessagesCode = 4010
	CloudIntegrationDiscoverAllResponseMessagesCode4011   CloudIntegrationDiscoverAllResponseMessagesCode = 4011
	CloudIntegrationDiscoverAllResponseMessagesCode4012   CloudIntegrationDiscoverAllResponseMessagesCode = 4012
	CloudIntegrationDiscoverAllResponseMessagesCode4013   CloudIntegrationDiscoverAllResponseMessagesCode = 4013
	CloudIntegrationDiscoverAllResponseMessagesCode4014   CloudIntegrationDiscoverAllResponseMessagesCode = 4014
	CloudIntegrationDiscoverAllResponseMessagesCode4015   CloudIntegrationDiscoverAllResponseMessagesCode = 4015
	CloudIntegrationDiscoverAllResponseMessagesCode4016   CloudIntegrationDiscoverAllResponseMessagesCode = 4016
	CloudIntegrationDiscoverAllResponseMessagesCode4017   CloudIntegrationDiscoverAllResponseMessagesCode = 4017
	CloudIntegrationDiscoverAllResponseMessagesCode4018   CloudIntegrationDiscoverAllResponseMessagesCode = 4018
	CloudIntegrationDiscoverAllResponseMessagesCode4019   CloudIntegrationDiscoverAllResponseMessagesCode = 4019
	CloudIntegrationDiscoverAllResponseMessagesCode4020   CloudIntegrationDiscoverAllResponseMessagesCode = 4020
	CloudIntegrationDiscoverAllResponseMessagesCode4021   CloudIntegrationDiscoverAllResponseMessagesCode = 4021
	CloudIntegrationDiscoverAllResponseMessagesCode4022   CloudIntegrationDiscoverAllResponseMessagesCode = 4022
	CloudIntegrationDiscoverAllResponseMessagesCode4023   CloudIntegrationDiscoverAllResponseMessagesCode = 4023
	CloudIntegrationDiscoverAllResponseMessagesCode5001   CloudIntegrationDiscoverAllResponseMessagesCode = 5001
	CloudIntegrationDiscoverAllResponseMessagesCode5002   CloudIntegrationDiscoverAllResponseMessagesCode = 5002
	CloudIntegrationDiscoverAllResponseMessagesCode5003   CloudIntegrationDiscoverAllResponseMessagesCode = 5003
	CloudIntegrationDiscoverAllResponseMessagesCode5004   CloudIntegrationDiscoverAllResponseMessagesCode = 5004
	CloudIntegrationDiscoverAllResponseMessagesCode102000 CloudIntegrationDiscoverAllResponseMessagesCode = 102000
	CloudIntegrationDiscoverAllResponseMessagesCode102001 CloudIntegrationDiscoverAllResponseMessagesCode = 102001
	CloudIntegrationDiscoverAllResponseMessagesCode102002 CloudIntegrationDiscoverAllResponseMessagesCode = 102002
	CloudIntegrationDiscoverAllResponseMessagesCode102003 CloudIntegrationDiscoverAllResponseMessagesCode = 102003
	CloudIntegrationDiscoverAllResponseMessagesCode102004 CloudIntegrationDiscoverAllResponseMessagesCode = 102004
	CloudIntegrationDiscoverAllResponseMessagesCode102005 CloudIntegrationDiscoverAllResponseMessagesCode = 102005
	CloudIntegrationDiscoverAllResponseMessagesCode102006 CloudIntegrationDiscoverAllResponseMessagesCode = 102006
	CloudIntegrationDiscoverAllResponseMessagesCode102007 CloudIntegrationDiscoverAllResponseMessagesCode = 102007
	CloudIntegrationDiscoverAllResponseMessagesCode102008 CloudIntegrationDiscoverAllResponseMessagesCode = 102008
	CloudIntegrationDiscoverAllResponseMessagesCode102009 CloudIntegrationDiscoverAllResponseMessagesCode = 102009
	CloudIntegrationDiscoverAllResponseMessagesCode102010 CloudIntegrationDiscoverAllResponseMessagesCode = 102010
	CloudIntegrationDiscoverAllResponseMessagesCode102011 CloudIntegrationDiscoverAllResponseMessagesCode = 102011
	CloudIntegrationDiscoverAllResponseMessagesCode102012 CloudIntegrationDiscoverAllResponseMessagesCode = 102012
	CloudIntegrationDiscoverAllResponseMessagesCode102013 CloudIntegrationDiscoverAllResponseMessagesCode = 102013
	CloudIntegrationDiscoverAllResponseMessagesCode102014 CloudIntegrationDiscoverAllResponseMessagesCode = 102014
	CloudIntegrationDiscoverAllResponseMessagesCode102015 CloudIntegrationDiscoverAllResponseMessagesCode = 102015
	CloudIntegrationDiscoverAllResponseMessagesCode102016 CloudIntegrationDiscoverAllResponseMessagesCode = 102016
	CloudIntegrationDiscoverAllResponseMessagesCode102017 CloudIntegrationDiscoverAllResponseMessagesCode = 102017
	CloudIntegrationDiscoverAllResponseMessagesCode102018 CloudIntegrationDiscoverAllResponseMessagesCode = 102018
	CloudIntegrationDiscoverAllResponseMessagesCode102019 CloudIntegrationDiscoverAllResponseMessagesCode = 102019
	CloudIntegrationDiscoverAllResponseMessagesCode102020 CloudIntegrationDiscoverAllResponseMessagesCode = 102020
	CloudIntegrationDiscoverAllResponseMessagesCode102021 CloudIntegrationDiscoverAllResponseMessagesCode = 102021
	CloudIntegrationDiscoverAllResponseMessagesCode102022 CloudIntegrationDiscoverAllResponseMessagesCode = 102022
	CloudIntegrationDiscoverAllResponseMessagesCode102023 CloudIntegrationDiscoverAllResponseMessagesCode = 102023
	CloudIntegrationDiscoverAllResponseMessagesCode102024 CloudIntegrationDiscoverAllResponseMessagesCode = 102024
	CloudIntegrationDiscoverAllResponseMessagesCode102025 CloudIntegrationDiscoverAllResponseMessagesCode = 102025
	CloudIntegrationDiscoverAllResponseMessagesCode102026 CloudIntegrationDiscoverAllResponseMessagesCode = 102026
	CloudIntegrationDiscoverAllResponseMessagesCode102027 CloudIntegrationDiscoverAllResponseMessagesCode = 102027
	CloudIntegrationDiscoverAllResponseMessagesCode102028 CloudIntegrationDiscoverAllResponseMessagesCode = 102028
	CloudIntegrationDiscoverAllResponseMessagesCode102029 CloudIntegrationDiscoverAllResponseMessagesCode = 102029
	CloudIntegrationDiscoverAllResponseMessagesCode102030 CloudIntegrationDiscoverAllResponseMessagesCode = 102030
	CloudIntegrationDiscoverAllResponseMessagesCode102031 CloudIntegrationDiscoverAllResponseMessagesCode = 102031
	CloudIntegrationDiscoverAllResponseMessagesCode102032 CloudIntegrationDiscoverAllResponseMessagesCode = 102032
	CloudIntegrationDiscoverAllResponseMessagesCode102033 CloudIntegrationDiscoverAllResponseMessagesCode = 102033
	CloudIntegrationDiscoverAllResponseMessagesCode102034 CloudIntegrationDiscoverAllResponseMessagesCode = 102034
	CloudIntegrationDiscoverAllResponseMessagesCode102035 CloudIntegrationDiscoverAllResponseMessagesCode = 102035
	CloudIntegrationDiscoverAllResponseMessagesCode102036 CloudIntegrationDiscoverAllResponseMessagesCode = 102036
	CloudIntegrationDiscoverAllResponseMessagesCode102037 CloudIntegrationDiscoverAllResponseMessagesCode = 102037
	CloudIntegrationDiscoverAllResponseMessagesCode102038 CloudIntegrationDiscoverAllResponseMessagesCode = 102038
	CloudIntegrationDiscoverAllResponseMessagesCode102039 CloudIntegrationDiscoverAllResponseMessagesCode = 102039
	CloudIntegrationDiscoverAllResponseMessagesCode102040 CloudIntegrationDiscoverAllResponseMessagesCode = 102040
	CloudIntegrationDiscoverAllResponseMessagesCode102041 CloudIntegrationDiscoverAllResponseMessagesCode = 102041
	CloudIntegrationDiscoverAllResponseMessagesCode102042 CloudIntegrationDiscoverAllResponseMessagesCode = 102042
	CloudIntegrationDiscoverAllResponseMessagesCode102043 CloudIntegrationDiscoverAllResponseMessagesCode = 102043
	CloudIntegrationDiscoverAllResponseMessagesCode102044 CloudIntegrationDiscoverAllResponseMessagesCode = 102044
	CloudIntegrationDiscoverAllResponseMessagesCode102045 CloudIntegrationDiscoverAllResponseMessagesCode = 102045
	CloudIntegrationDiscoverAllResponseMessagesCode102046 CloudIntegrationDiscoverAllResponseMessagesCode = 102046
	CloudIntegrationDiscoverAllResponseMessagesCode102047 CloudIntegrationDiscoverAllResponseMessagesCode = 102047
	CloudIntegrationDiscoverAllResponseMessagesCode102048 CloudIntegrationDiscoverAllResponseMessagesCode = 102048
	CloudIntegrationDiscoverAllResponseMessagesCode102049 CloudIntegrationDiscoverAllResponseMessagesCode = 102049
	CloudIntegrationDiscoverAllResponseMessagesCode102050 CloudIntegrationDiscoverAllResponseMessagesCode = 102050
	CloudIntegrationDiscoverAllResponseMessagesCode102051 CloudIntegrationDiscoverAllResponseMessagesCode = 102051
	CloudIntegrationDiscoverAllResponseMessagesCode102052 CloudIntegrationDiscoverAllResponseMessagesCode = 102052
	CloudIntegrationDiscoverAllResponseMessagesCode102053 CloudIntegrationDiscoverAllResponseMessagesCode = 102053
	CloudIntegrationDiscoverAllResponseMessagesCode102054 CloudIntegrationDiscoverAllResponseMessagesCode = 102054
	CloudIntegrationDiscoverAllResponseMessagesCode102055 CloudIntegrationDiscoverAllResponseMessagesCode = 102055
	CloudIntegrationDiscoverAllResponseMessagesCode102056 CloudIntegrationDiscoverAllResponseMessagesCode = 102056
	CloudIntegrationDiscoverAllResponseMessagesCode102057 CloudIntegrationDiscoverAllResponseMessagesCode = 102057
	CloudIntegrationDiscoverAllResponseMessagesCode102058 CloudIntegrationDiscoverAllResponseMessagesCode = 102058
	CloudIntegrationDiscoverAllResponseMessagesCode102059 CloudIntegrationDiscoverAllResponseMessagesCode = 102059
	CloudIntegrationDiscoverAllResponseMessagesCode102060 CloudIntegrationDiscoverAllResponseMessagesCode = 102060
	CloudIntegrationDiscoverAllResponseMessagesCode102061 CloudIntegrationDiscoverAllResponseMessagesCode = 102061
	CloudIntegrationDiscoverAllResponseMessagesCode102062 CloudIntegrationDiscoverAllResponseMessagesCode = 102062
	CloudIntegrationDiscoverAllResponseMessagesCode102063 CloudIntegrationDiscoverAllResponseMessagesCode = 102063
	CloudIntegrationDiscoverAllResponseMessagesCode102064 CloudIntegrationDiscoverAllResponseMessagesCode = 102064
	CloudIntegrationDiscoverAllResponseMessagesCode102065 CloudIntegrationDiscoverAllResponseMessagesCode = 102065
	CloudIntegrationDiscoverAllResponseMessagesCode102066 CloudIntegrationDiscoverAllResponseMessagesCode = 102066
	CloudIntegrationDiscoverAllResponseMessagesCode103001 CloudIntegrationDiscoverAllResponseMessagesCode = 103001
	CloudIntegrationDiscoverAllResponseMessagesCode103002 CloudIntegrationDiscoverAllResponseMessagesCode = 103002
	CloudIntegrationDiscoverAllResponseMessagesCode103003 CloudIntegrationDiscoverAllResponseMessagesCode = 103003
	CloudIntegrationDiscoverAllResponseMessagesCode103004 CloudIntegrationDiscoverAllResponseMessagesCode = 103004
	CloudIntegrationDiscoverAllResponseMessagesCode103005 CloudIntegrationDiscoverAllResponseMessagesCode = 103005
	CloudIntegrationDiscoverAllResponseMessagesCode103006 CloudIntegrationDiscoverAllResponseMessagesCode = 103006
	CloudIntegrationDiscoverAllResponseMessagesCode103007 CloudIntegrationDiscoverAllResponseMessagesCode = 103007
	CloudIntegrationDiscoverAllResponseMessagesCode103008 CloudIntegrationDiscoverAllResponseMessagesCode = 103008
)

func (r CloudIntegrationDiscoverAllResponseMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDiscoverAllResponseMessagesCode1001, CloudIntegrationDiscoverAllResponseMessagesCode1002, CloudIntegrationDiscoverAllResponseMessagesCode1003, CloudIntegrationDiscoverAllResponseMessagesCode1004, CloudIntegrationDiscoverAllResponseMessagesCode1005, CloudIntegrationDiscoverAllResponseMessagesCode1006, CloudIntegrationDiscoverAllResponseMessagesCode1007, CloudIntegrationDiscoverAllResponseMessagesCode1008, CloudIntegrationDiscoverAllResponseMessagesCode1009, CloudIntegrationDiscoverAllResponseMessagesCode1010, CloudIntegrationDiscoverAllResponseMessagesCode1011, CloudIntegrationDiscoverAllResponseMessagesCode1012, CloudIntegrationDiscoverAllResponseMessagesCode1013, CloudIntegrationDiscoverAllResponseMessagesCode1014, CloudIntegrationDiscoverAllResponseMessagesCode1015, CloudIntegrationDiscoverAllResponseMessagesCode1016, CloudIntegrationDiscoverAllResponseMessagesCode1017, CloudIntegrationDiscoverAllResponseMessagesCode2001, CloudIntegrationDiscoverAllResponseMessagesCode2002, CloudIntegrationDiscoverAllResponseMessagesCode2003, CloudIntegrationDiscoverAllResponseMessagesCode2004, CloudIntegrationDiscoverAllResponseMessagesCode2005, CloudIntegrationDiscoverAllResponseMessagesCode2006, CloudIntegrationDiscoverAllResponseMessagesCode2007, CloudIntegrationDiscoverAllResponseMessagesCode2008, CloudIntegrationDiscoverAllResponseMessagesCode2009, CloudIntegrationDiscoverAllResponseMessagesCode2010, CloudIntegrationDiscoverAllResponseMessagesCode2011, CloudIntegrationDiscoverAllResponseMessagesCode2012, CloudIntegrationDiscoverAllResponseMessagesCode2013, CloudIntegrationDiscoverAllResponseMessagesCode2014, CloudIntegrationDiscoverAllResponseMessagesCode2015, CloudIntegrationDiscoverAllResponseMessagesCode2016, CloudIntegrationDiscoverAllResponseMessagesCode2017, CloudIntegrationDiscoverAllResponseMessagesCode2018, CloudIntegrationDiscoverAllResponseMessagesCode2019, CloudIntegrationDiscoverAllResponseMessagesCode2020, CloudIntegrationDiscoverAllResponseMessagesCode2021, CloudIntegrationDiscoverAllResponseMessagesCode2022, CloudIntegrationDiscoverAllResponseMessagesCode3001, CloudIntegrationDiscoverAllResponseMessagesCode3002, CloudIntegrationDiscoverAllResponseMessagesCode3003, CloudIntegrationDiscoverAllResponseMessagesCode3004, CloudIntegrationDiscoverAllResponseMessagesCode3005, CloudIntegrationDiscoverAllResponseMessagesCode3006, CloudIntegrationDiscoverAllResponseMessagesCode3007, CloudIntegrationDiscoverAllResponseMessagesCode4001, CloudIntegrationDiscoverAllResponseMessagesCode4002, CloudIntegrationDiscoverAllResponseMessagesCode4003, CloudIntegrationDiscoverAllResponseMessagesCode4004, CloudIntegrationDiscoverAllResponseMessagesCode4005, CloudIntegrationDiscoverAllResponseMessagesCode4006, CloudIntegrationDiscoverAllResponseMessagesCode4007, CloudIntegrationDiscoverAllResponseMessagesCode4008, CloudIntegrationDiscoverAllResponseMessagesCode4009, CloudIntegrationDiscoverAllResponseMessagesCode4010, CloudIntegrationDiscoverAllResponseMessagesCode4011, CloudIntegrationDiscoverAllResponseMessagesCode4012, CloudIntegrationDiscoverAllResponseMessagesCode4013, CloudIntegrationDiscoverAllResponseMessagesCode4014, CloudIntegrationDiscoverAllResponseMessagesCode4015, CloudIntegrationDiscoverAllResponseMessagesCode4016, CloudIntegrationDiscoverAllResponseMessagesCode4017, CloudIntegrationDiscoverAllResponseMessagesCode4018, CloudIntegrationDiscoverAllResponseMessagesCode4019, CloudIntegrationDiscoverAllResponseMessagesCode4020, CloudIntegrationDiscoverAllResponseMessagesCode4021, CloudIntegrationDiscoverAllResponseMessagesCode4022, CloudIntegrationDiscoverAllResponseMessagesCode4023, CloudIntegrationDiscoverAllResponseMessagesCode5001, CloudIntegrationDiscoverAllResponseMessagesCode5002, CloudIntegrationDiscoverAllResponseMessagesCode5003, CloudIntegrationDiscoverAllResponseMessagesCode5004, CloudIntegrationDiscoverAllResponseMessagesCode102000, CloudIntegrationDiscoverAllResponseMessagesCode102001, CloudIntegrationDiscoverAllResponseMessagesCode102002, CloudIntegrationDiscoverAllResponseMessagesCode102003, CloudIntegrationDiscoverAllResponseMessagesCode102004, CloudIntegrationDiscoverAllResponseMessagesCode102005, CloudIntegrationDiscoverAllResponseMessagesCode102006, CloudIntegrationDiscoverAllResponseMessagesCode102007, CloudIntegrationDiscoverAllResponseMessagesCode102008, CloudIntegrationDiscoverAllResponseMessagesCode102009, CloudIntegrationDiscoverAllResponseMessagesCode102010, CloudIntegrationDiscoverAllResponseMessagesCode102011, CloudIntegrationDiscoverAllResponseMessagesCode102012, CloudIntegrationDiscoverAllResponseMessagesCode102013, CloudIntegrationDiscoverAllResponseMessagesCode102014, CloudIntegrationDiscoverAllResponseMessagesCode102015, CloudIntegrationDiscoverAllResponseMessagesCode102016, CloudIntegrationDiscoverAllResponseMessagesCode102017, CloudIntegrationDiscoverAllResponseMessagesCode102018, CloudIntegrationDiscoverAllResponseMessagesCode102019, CloudIntegrationDiscoverAllResponseMessagesCode102020, CloudIntegrationDiscoverAllResponseMessagesCode102021, CloudIntegrationDiscoverAllResponseMessagesCode102022, CloudIntegrationDiscoverAllResponseMessagesCode102023, CloudIntegrationDiscoverAllResponseMessagesCode102024, CloudIntegrationDiscoverAllResponseMessagesCode102025, CloudIntegrationDiscoverAllResponseMessagesCode102026, CloudIntegrationDiscoverAllResponseMessagesCode102027, CloudIntegrationDiscoverAllResponseMessagesCode102028, CloudIntegrationDiscoverAllResponseMessagesCode102029, CloudIntegrationDiscoverAllResponseMessagesCode102030, CloudIntegrationDiscoverAllResponseMessagesCode102031, CloudIntegrationDiscoverAllResponseMessagesCode102032, CloudIntegrationDiscoverAllResponseMessagesCode102033, CloudIntegrationDiscoverAllResponseMessagesCode102034, CloudIntegrationDiscoverAllResponseMessagesCode102035, CloudIntegrationDiscoverAllResponseMessagesCode102036, CloudIntegrationDiscoverAllResponseMessagesCode102037, CloudIntegrationDiscoverAllResponseMessagesCode102038, CloudIntegrationDiscoverAllResponseMessagesCode102039, CloudIntegrationDiscoverAllResponseMessagesCode102040, CloudIntegrationDiscoverAllResponseMessagesCode102041, CloudIntegrationDiscoverAllResponseMessagesCode102042, CloudIntegrationDiscoverAllResponseMessagesCode102043, CloudIntegrationDiscoverAllResponseMessagesCode102044, CloudIntegrationDiscoverAllResponseMessagesCode102045, CloudIntegrationDiscoverAllResponseMessagesCode102046, CloudIntegrationDiscoverAllResponseMessagesCode102047, CloudIntegrationDiscoverAllResponseMessagesCode102048, CloudIntegrationDiscoverAllResponseMessagesCode102049, CloudIntegrationDiscoverAllResponseMessagesCode102050, CloudIntegrationDiscoverAllResponseMessagesCode102051, CloudIntegrationDiscoverAllResponseMessagesCode102052, CloudIntegrationDiscoverAllResponseMessagesCode102053, CloudIntegrationDiscoverAllResponseMessagesCode102054, CloudIntegrationDiscoverAllResponseMessagesCode102055, CloudIntegrationDiscoverAllResponseMessagesCode102056, CloudIntegrationDiscoverAllResponseMessagesCode102057, CloudIntegrationDiscoverAllResponseMessagesCode102058, CloudIntegrationDiscoverAllResponseMessagesCode102059, CloudIntegrationDiscoverAllResponseMessagesCode102060, CloudIntegrationDiscoverAllResponseMessagesCode102061, CloudIntegrationDiscoverAllResponseMessagesCode102062, CloudIntegrationDiscoverAllResponseMessagesCode102063, CloudIntegrationDiscoverAllResponseMessagesCode102064, CloudIntegrationDiscoverAllResponseMessagesCode102065, CloudIntegrationDiscoverAllResponseMessagesCode102066, CloudIntegrationDiscoverAllResponseMessagesCode103001, CloudIntegrationDiscoverAllResponseMessagesCode103002, CloudIntegrationDiscoverAllResponseMessagesCode103003, CloudIntegrationDiscoverAllResponseMessagesCode103004, CloudIntegrationDiscoverAllResponseMessagesCode103005, CloudIntegrationDiscoverAllResponseMessagesCode103006, CloudIntegrationDiscoverAllResponseMessagesCode103007, CloudIntegrationDiscoverAllResponseMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationDiscoverAllResponseMessagesMeta struct {
	L10nKey       string                                              `json:"l10n_key"`
	LoggableError string                                              `json:"loggable_error"`
	TemplateData  interface{}                                         `json:"template_data"`
	TraceID       string                                              `json:"trace_id"`
	JSON          cloudIntegrationDiscoverAllResponseMessagesMetaJSON `json:"-"`
}

// cloudIntegrationDiscoverAllResponseMessagesMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationDiscoverAllResponseMessagesMeta]
type cloudIntegrationDiscoverAllResponseMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverAllResponseMessagesSource struct {
	Parameter           string                                                `json:"parameter"`
	ParameterValueIndex int64                                                 `json:"parameter_value_index"`
	Pointer             string                                                `json:"pointer"`
	JSON                cloudIntegrationDiscoverAllResponseMessagesSourceJSON `json:"-"`
}

// cloudIntegrationDiscoverAllResponseMessagesSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationDiscoverAllResponseMessagesSource]
type cloudIntegrationDiscoverAllResponseMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDiscoverAllResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDiscoverAllResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponse struct {
	ID                     string                                     `json:"id,required" format:"uuid"`
	CloudType              CloudIntegrationEditResponseCloudType      `json:"cloud_type,required"`
	FriendlyName           string                                     `json:"friendly_name,required"`
	LastUpdated            string                                     `json:"last_updated,required"`
	LifecycleState         CloudIntegrationEditResponseLifecycleState `json:"lifecycle_state,required"`
	State                  CloudIntegrationEditResponseState          `json:"state,required"`
	StateV2                CloudIntegrationEditResponseStateV2        `json:"state_v2,required"`
	AwsArn                 string                                     `json:"aws_arn"`
	AzureSubscriptionID    string                                     `json:"azure_subscription_id"`
	AzureTenantID          string                                     `json:"azure_tenant_id"`
	Description            string                                     `json:"description"`
	GcpProjectID           string                                     `json:"gcp_project_id"`
	GcpServiceAccountEmail string                                     `json:"gcp_service_account_email"`
	Status                 CloudIntegrationEditResponseStatus         `json:"status"`
	JSON                   cloudIntegrationEditResponseJSON           `json:"-"`
}

// cloudIntegrationEditResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationEditResponse]
type cloudIntegrationEditResponseJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseCloudType string

const (
	CloudIntegrationEditResponseCloudTypeAws        CloudIntegrationEditResponseCloudType = "AWS"
	CloudIntegrationEditResponseCloudTypeAzure      CloudIntegrationEditResponseCloudType = "AZURE"
	CloudIntegrationEditResponseCloudTypeGoogle     CloudIntegrationEditResponseCloudType = "GOOGLE"
	CloudIntegrationEditResponseCloudTypeCloudflare CloudIntegrationEditResponseCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationEditResponseCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseCloudTypeAws, CloudIntegrationEditResponseCloudTypeAzure, CloudIntegrationEditResponseCloudTypeGoogle, CloudIntegrationEditResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationEditResponseLifecycleState string

const (
	CloudIntegrationEditResponseLifecycleStateActive       CloudIntegrationEditResponseLifecycleState = "ACTIVE"
	CloudIntegrationEditResponseLifecycleStatePendingSetup CloudIntegrationEditResponseLifecycleState = "PENDING_SETUP"
	CloudIntegrationEditResponseLifecycleStateRetired      CloudIntegrationEditResponseLifecycleState = "RETIRED"
)

func (r CloudIntegrationEditResponseLifecycleState) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseLifecycleStateActive, CloudIntegrationEditResponseLifecycleStatePendingSetup, CloudIntegrationEditResponseLifecycleStateRetired:
		return true
	}
	return false
}

type CloudIntegrationEditResponseState string

const (
	CloudIntegrationEditResponseStateUnspecified CloudIntegrationEditResponseState = "UNSPECIFIED"
	CloudIntegrationEditResponseStatePending     CloudIntegrationEditResponseState = "PENDING"
	CloudIntegrationEditResponseStateDiscovering CloudIntegrationEditResponseState = "DISCOVERING"
	CloudIntegrationEditResponseStateFailed      CloudIntegrationEditResponseState = "FAILED"
	CloudIntegrationEditResponseStateSucceeded   CloudIntegrationEditResponseState = "SUCCEEDED"
)

func (r CloudIntegrationEditResponseState) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseStateUnspecified, CloudIntegrationEditResponseStatePending, CloudIntegrationEditResponseStateDiscovering, CloudIntegrationEditResponseStateFailed, CloudIntegrationEditResponseStateSucceeded:
		return true
	}
	return false
}

type CloudIntegrationEditResponseStateV2 string

const (
	CloudIntegrationEditResponseStateV2Unspecified CloudIntegrationEditResponseStateV2 = "UNSPECIFIED"
	CloudIntegrationEditResponseStateV2Pending     CloudIntegrationEditResponseStateV2 = "PENDING"
	CloudIntegrationEditResponseStateV2Discovering CloudIntegrationEditResponseStateV2 = "DISCOVERING"
	CloudIntegrationEditResponseStateV2Failed      CloudIntegrationEditResponseStateV2 = "FAILED"
	CloudIntegrationEditResponseStateV2Succeeded   CloudIntegrationEditResponseStateV2 = "SUCCEEDED"
)

func (r CloudIntegrationEditResponseStateV2) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseStateV2Unspecified, CloudIntegrationEditResponseStateV2Pending, CloudIntegrationEditResponseStateV2Discovering, CloudIntegrationEditResponseStateV2Failed, CloudIntegrationEditResponseStateV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationEditResponseStatus struct {
	DiscoveryProgress          CloudIntegrationEditResponseStatusDiscoveryProgress     `json:"discovery_progress,required"`
	DiscoveryProgressV2        CloudIntegrationEditResponseStatusDiscoveryProgressV2   `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        CloudIntegrationEditResponseStatusLastDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 `json:"last_discovery_status_v2,required"`
	Regions                    []string                                                `json:"regions,required"`
	CredentialsGoodSince       string                                                  `json:"credentials_good_since"`
	CredentialsMissingSince    string                                                  `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                                                  `json:"credentials_rejected_since"`
	DiscoveryMessage           string                                                  `json:"discovery_message"`
	DiscoveryMessageV2         string                                                  `json:"discovery_message_v2"`
	InUseBy                    []CloudIntegrationEditResponseStatusInUseBy             `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                                                  `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                                                  `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                                                  `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                                                  `json:"last_discovery_started_at_v2"`
	LastUpdated                string                                                  `json:"last_updated"`
	JSON                       cloudIntegrationEditResponseStatusJSON                  `json:"-"`
}

// cloudIntegrationEditResponseStatusJSON contains the JSON metadata for the struct
// [CloudIntegrationEditResponseStatus]
type cloudIntegrationEditResponseStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseStatusDiscoveryProgress struct {
	Done  int64                                                   `json:"done,required"`
	Total int64                                                   `json:"total,required"`
	Unit  string                                                  `json:"unit,required"`
	JSON  cloudIntegrationEditResponseStatusDiscoveryProgressJSON `json:"-"`
}

// cloudIntegrationEditResponseStatusDiscoveryProgressJSON contains the JSON
// metadata for the struct [CloudIntegrationEditResponseStatusDiscoveryProgress]
type cloudIntegrationEditResponseStatusDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseStatusDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseStatusDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseStatusDiscoveryProgressV2 struct {
	Done  int64                                                     `json:"done,required"`
	Total int64                                                     `json:"total,required"`
	Unit  string                                                    `json:"unit,required"`
	JSON  cloudIntegrationEditResponseStatusDiscoveryProgressV2JSON `json:"-"`
}

// cloudIntegrationEditResponseStatusDiscoveryProgressV2JSON contains the JSON
// metadata for the struct [CloudIntegrationEditResponseStatusDiscoveryProgressV2]
type cloudIntegrationEditResponseStatusDiscoveryProgressV2JSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseStatusDiscoveryProgressV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseStatusDiscoveryProgressV2JSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseStatusLastDiscoveryStatus string

const (
	CloudIntegrationEditResponseStatusLastDiscoveryStatusUnspecified CloudIntegrationEditResponseStatusLastDiscoveryStatus = "UNSPECIFIED"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusPending     CloudIntegrationEditResponseStatusLastDiscoveryStatus = "PENDING"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusDiscovering CloudIntegrationEditResponseStatusLastDiscoveryStatus = "DISCOVERING"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusFailed      CloudIntegrationEditResponseStatusLastDiscoveryStatus = "FAILED"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusSucceeded   CloudIntegrationEditResponseStatusLastDiscoveryStatus = "SUCCEEDED"
)

func (r CloudIntegrationEditResponseStatusLastDiscoveryStatus) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseStatusLastDiscoveryStatusUnspecified, CloudIntegrationEditResponseStatusLastDiscoveryStatusPending, CloudIntegrationEditResponseStatusLastDiscoveryStatusDiscovering, CloudIntegrationEditResponseStatusLastDiscoveryStatusFailed, CloudIntegrationEditResponseStatusLastDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 string

const (
	CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Unspecified CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 = "UNSPECIFIED"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Pending     CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 = "PENDING"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Discovering CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 = "DISCOVERING"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Failed      CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 = "FAILED"
	CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Succeeded   CloudIntegrationEditResponseStatusLastDiscoveryStatusV2 = "SUCCEEDED"
)

func (r CloudIntegrationEditResponseStatusLastDiscoveryStatusV2) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Unspecified, CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Pending, CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Discovering, CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Failed, CloudIntegrationEditResponseStatusLastDiscoveryStatusV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationEditResponseStatusInUseBy struct {
	ID         string                                              `json:"id,required" format:"uuid"`
	ClientType CloudIntegrationEditResponseStatusInUseByClientType `json:"client_type,required"`
	Name       string                                              `json:"name,required"`
	JSON       cloudIntegrationEditResponseStatusInUseByJSON       `json:"-"`
}

// cloudIntegrationEditResponseStatusInUseByJSON contains the JSON metadata for the
// struct [CloudIntegrationEditResponseStatusInUseBy]
type cloudIntegrationEditResponseStatusInUseByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseStatusInUseBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseStatusInUseByJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseStatusInUseByClientType string

const (
	CloudIntegrationEditResponseStatusInUseByClientTypeMagicWANCloudOnramp CloudIntegrationEditResponseStatusInUseByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r CloudIntegrationEditResponseStatusInUseByClientType) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseStatusInUseByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type CloudIntegrationGetResponse struct {
	ID                     string                                    `json:"id,required" format:"uuid"`
	CloudType              CloudIntegrationGetResponseCloudType      `json:"cloud_type,required"`
	FriendlyName           string                                    `json:"friendly_name,required"`
	LastUpdated            string                                    `json:"last_updated,required"`
	LifecycleState         CloudIntegrationGetResponseLifecycleState `json:"lifecycle_state,required"`
	State                  CloudIntegrationGetResponseState          `json:"state,required"`
	StateV2                CloudIntegrationGetResponseStateV2        `json:"state_v2,required"`
	AwsArn                 string                                    `json:"aws_arn"`
	AzureSubscriptionID    string                                    `json:"azure_subscription_id"`
	AzureTenantID          string                                    `json:"azure_tenant_id"`
	Description            string                                    `json:"description"`
	GcpProjectID           string                                    `json:"gcp_project_id"`
	GcpServiceAccountEmail string                                    `json:"gcp_service_account_email"`
	Status                 CloudIntegrationGetResponseStatus         `json:"status"`
	JSON                   cloudIntegrationGetResponseJSON           `json:"-"`
}

// cloudIntegrationGetResponseJSON contains the JSON metadata for the struct
// [CloudIntegrationGetResponse]
type cloudIntegrationGetResponseJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseCloudType string

const (
	CloudIntegrationGetResponseCloudTypeAws        CloudIntegrationGetResponseCloudType = "AWS"
	CloudIntegrationGetResponseCloudTypeAzure      CloudIntegrationGetResponseCloudType = "AZURE"
	CloudIntegrationGetResponseCloudTypeGoogle     CloudIntegrationGetResponseCloudType = "GOOGLE"
	CloudIntegrationGetResponseCloudTypeCloudflare CloudIntegrationGetResponseCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationGetResponseCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseCloudTypeAws, CloudIntegrationGetResponseCloudTypeAzure, CloudIntegrationGetResponseCloudTypeGoogle, CloudIntegrationGetResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationGetResponseLifecycleState string

const (
	CloudIntegrationGetResponseLifecycleStateActive       CloudIntegrationGetResponseLifecycleState = "ACTIVE"
	CloudIntegrationGetResponseLifecycleStatePendingSetup CloudIntegrationGetResponseLifecycleState = "PENDING_SETUP"
	CloudIntegrationGetResponseLifecycleStateRetired      CloudIntegrationGetResponseLifecycleState = "RETIRED"
)

func (r CloudIntegrationGetResponseLifecycleState) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseLifecycleStateActive, CloudIntegrationGetResponseLifecycleStatePendingSetup, CloudIntegrationGetResponseLifecycleStateRetired:
		return true
	}
	return false
}

type CloudIntegrationGetResponseState string

const (
	CloudIntegrationGetResponseStateUnspecified CloudIntegrationGetResponseState = "UNSPECIFIED"
	CloudIntegrationGetResponseStatePending     CloudIntegrationGetResponseState = "PENDING"
	CloudIntegrationGetResponseStateDiscovering CloudIntegrationGetResponseState = "DISCOVERING"
	CloudIntegrationGetResponseStateFailed      CloudIntegrationGetResponseState = "FAILED"
	CloudIntegrationGetResponseStateSucceeded   CloudIntegrationGetResponseState = "SUCCEEDED"
)

func (r CloudIntegrationGetResponseState) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseStateUnspecified, CloudIntegrationGetResponseStatePending, CloudIntegrationGetResponseStateDiscovering, CloudIntegrationGetResponseStateFailed, CloudIntegrationGetResponseStateSucceeded:
		return true
	}
	return false
}

type CloudIntegrationGetResponseStateV2 string

const (
	CloudIntegrationGetResponseStateV2Unspecified CloudIntegrationGetResponseStateV2 = "UNSPECIFIED"
	CloudIntegrationGetResponseStateV2Pending     CloudIntegrationGetResponseStateV2 = "PENDING"
	CloudIntegrationGetResponseStateV2Discovering CloudIntegrationGetResponseStateV2 = "DISCOVERING"
	CloudIntegrationGetResponseStateV2Failed      CloudIntegrationGetResponseStateV2 = "FAILED"
	CloudIntegrationGetResponseStateV2Succeeded   CloudIntegrationGetResponseStateV2 = "SUCCEEDED"
)

func (r CloudIntegrationGetResponseStateV2) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseStateV2Unspecified, CloudIntegrationGetResponseStateV2Pending, CloudIntegrationGetResponseStateV2Discovering, CloudIntegrationGetResponseStateV2Failed, CloudIntegrationGetResponseStateV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationGetResponseStatus struct {
	DiscoveryProgress          CloudIntegrationGetResponseStatusDiscoveryProgress     `json:"discovery_progress,required"`
	DiscoveryProgressV2        CloudIntegrationGetResponseStatusDiscoveryProgressV2   `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        CloudIntegrationGetResponseStatusLastDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 `json:"last_discovery_status_v2,required"`
	Regions                    []string                                               `json:"regions,required"`
	CredentialsGoodSince       string                                                 `json:"credentials_good_since"`
	CredentialsMissingSince    string                                                 `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                                                 `json:"credentials_rejected_since"`
	DiscoveryMessage           string                                                 `json:"discovery_message"`
	DiscoveryMessageV2         string                                                 `json:"discovery_message_v2"`
	InUseBy                    []CloudIntegrationGetResponseStatusInUseBy             `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                                                 `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                                                 `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                                                 `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                                                 `json:"last_discovery_started_at_v2"`
	LastUpdated                string                                                 `json:"last_updated"`
	JSON                       cloudIntegrationGetResponseStatusJSON                  `json:"-"`
}

// cloudIntegrationGetResponseStatusJSON contains the JSON metadata for the struct
// [CloudIntegrationGetResponseStatus]
type cloudIntegrationGetResponseStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseStatusJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseStatusDiscoveryProgress struct {
	Done  int64                                                  `json:"done,required"`
	Total int64                                                  `json:"total,required"`
	Unit  string                                                 `json:"unit,required"`
	JSON  cloudIntegrationGetResponseStatusDiscoveryProgressJSON `json:"-"`
}

// cloudIntegrationGetResponseStatusDiscoveryProgressJSON contains the JSON
// metadata for the struct [CloudIntegrationGetResponseStatusDiscoveryProgress]
type cloudIntegrationGetResponseStatusDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseStatusDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseStatusDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseStatusDiscoveryProgressV2 struct {
	Done  int64                                                    `json:"done,required"`
	Total int64                                                    `json:"total,required"`
	Unit  string                                                   `json:"unit,required"`
	JSON  cloudIntegrationGetResponseStatusDiscoveryProgressV2JSON `json:"-"`
}

// cloudIntegrationGetResponseStatusDiscoveryProgressV2JSON contains the JSON
// metadata for the struct [CloudIntegrationGetResponseStatusDiscoveryProgressV2]
type cloudIntegrationGetResponseStatusDiscoveryProgressV2JSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseStatusDiscoveryProgressV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseStatusDiscoveryProgressV2JSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseStatusLastDiscoveryStatus string

const (
	CloudIntegrationGetResponseStatusLastDiscoveryStatusUnspecified CloudIntegrationGetResponseStatusLastDiscoveryStatus = "UNSPECIFIED"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusPending     CloudIntegrationGetResponseStatusLastDiscoveryStatus = "PENDING"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusDiscovering CloudIntegrationGetResponseStatusLastDiscoveryStatus = "DISCOVERING"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusFailed      CloudIntegrationGetResponseStatusLastDiscoveryStatus = "FAILED"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusSucceeded   CloudIntegrationGetResponseStatusLastDiscoveryStatus = "SUCCEEDED"
)

func (r CloudIntegrationGetResponseStatusLastDiscoveryStatus) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseStatusLastDiscoveryStatusUnspecified, CloudIntegrationGetResponseStatusLastDiscoveryStatusPending, CloudIntegrationGetResponseStatusLastDiscoveryStatusDiscovering, CloudIntegrationGetResponseStatusLastDiscoveryStatusFailed, CloudIntegrationGetResponseStatusLastDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 string

const (
	CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Unspecified CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 = "UNSPECIFIED"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Pending     CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 = "PENDING"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Discovering CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 = "DISCOVERING"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Failed      CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 = "FAILED"
	CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Succeeded   CloudIntegrationGetResponseStatusLastDiscoveryStatusV2 = "SUCCEEDED"
)

func (r CloudIntegrationGetResponseStatusLastDiscoveryStatusV2) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Unspecified, CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Pending, CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Discovering, CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Failed, CloudIntegrationGetResponseStatusLastDiscoveryStatusV2Succeeded:
		return true
	}
	return false
}

type CloudIntegrationGetResponseStatusInUseBy struct {
	ID         string                                             `json:"id,required" format:"uuid"`
	ClientType CloudIntegrationGetResponseStatusInUseByClientType `json:"client_type,required"`
	Name       string                                             `json:"name,required"`
	JSON       cloudIntegrationGetResponseStatusInUseByJSON       `json:"-"`
}

// cloudIntegrationGetResponseStatusInUseByJSON contains the JSON metadata for the
// struct [CloudIntegrationGetResponseStatusInUseBy]
type cloudIntegrationGetResponseStatusInUseByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseStatusInUseBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseStatusInUseByJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseStatusInUseByClientType string

const (
	CloudIntegrationGetResponseStatusInUseByClientTypeMagicWANCloudOnramp CloudIntegrationGetResponseStatusInUseByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r CloudIntegrationGetResponseStatusInUseByClientType) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseStatusInUseByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type CloudIntegrationInitialSetupResponse struct {
	ItemType               string                                   `json:"item_type,required"`
	AwsTrustPolicy         string                                   `json:"aws_trust_policy"`
	AzureConsentURL        string                                   `json:"azure_consent_url"`
	IntegrationIdentityTag string                                   `json:"integration_identity_tag"`
	TagCliCommand          string                                   `json:"tag_cli_command"`
	JSON                   cloudIntegrationInitialSetupResponseJSON `json:"-"`
	union                  CloudIntegrationInitialSetupResponseUnion
}

// cloudIntegrationInitialSetupResponseJSON contains the JSON metadata for the
// struct [CloudIntegrationInitialSetupResponse]
type cloudIntegrationInitialSetupResponseJSON struct {
	ItemType               apijson.Field
	AwsTrustPolicy         apijson.Field
	AzureConsentURL        apijson.Field
	IntegrationIdentityTag apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r cloudIntegrationInitialSetupResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CloudIntegrationInitialSetupResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CloudIntegrationInitialSetupResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudIntegrationInitialSetupResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy],
// [CloudIntegrationInitialSetupResponseMcnAzureSetup],
// [CloudIntegrationInitialSetupResponseMcnGcpSetup].
func (r CloudIntegrationInitialSetupResponse) AsUnion() CloudIntegrationInitialSetupResponseUnion {
	return r.union
}

// Union satisfied by [CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy],
// [CloudIntegrationInitialSetupResponseMcnAzureSetup] or
// [CloudIntegrationInitialSetupResponseMcnGcpSetup].
type CloudIntegrationInitialSetupResponseUnion interface {
	implementsCloudIntegrationInitialSetupResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudIntegrationInitialSetupResponseUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudIntegrationInitialSetupResponseMcnAzureSetup{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudIntegrationInitialSetupResponseMcnGcpSetup{}),
		},
	)
}

type CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy struct {
	AwsTrustPolicy string                                                    `json:"aws_trust_policy,required"`
	ItemType       string                                                    `json:"item_type,required"`
	JSON           cloudIntegrationInitialSetupResponseMcnAwsTrustPolicyJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseMcnAwsTrustPolicyJSON contains the JSON
// metadata for the struct [CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy]
type cloudIntegrationInitialSetupResponseMcnAwsTrustPolicyJSON struct {
	AwsTrustPolicy apijson.Field
	ItemType       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseMcnAwsTrustPolicyJSON) RawJSON() string {
	return r.raw
}

func (r CloudIntegrationInitialSetupResponseMcnAwsTrustPolicy) implementsCloudIntegrationInitialSetupResponse() {
}

type CloudIntegrationInitialSetupResponseMcnAzureSetup struct {
	AzureConsentURL        string                                                `json:"azure_consent_url,required"`
	IntegrationIdentityTag string                                                `json:"integration_identity_tag,required"`
	ItemType               string                                                `json:"item_type,required"`
	TagCliCommand          string                                                `json:"tag_cli_command,required"`
	JSON                   cloudIntegrationInitialSetupResponseMcnAzureSetupJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseMcnAzureSetupJSON contains the JSON metadata
// for the struct [CloudIntegrationInitialSetupResponseMcnAzureSetup]
type cloudIntegrationInitialSetupResponseMcnAzureSetupJSON struct {
	AzureConsentURL        apijson.Field
	IntegrationIdentityTag apijson.Field
	ItemType               apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseMcnAzureSetup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseMcnAzureSetupJSON) RawJSON() string {
	return r.raw
}

func (r CloudIntegrationInitialSetupResponseMcnAzureSetup) implementsCloudIntegrationInitialSetupResponse() {
}

type CloudIntegrationInitialSetupResponseMcnGcpSetup struct {
	IntegrationIdentityTag string                                              `json:"integration_identity_tag,required"`
	ItemType               string                                              `json:"item_type,required"`
	TagCliCommand          string                                              `json:"tag_cli_command,required"`
	JSON                   cloudIntegrationInitialSetupResponseMcnGcpSetupJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseMcnGcpSetupJSON contains the JSON metadata
// for the struct [CloudIntegrationInitialSetupResponseMcnGcpSetup]
type cloudIntegrationInitialSetupResponseMcnGcpSetupJSON struct {
	IntegrationIdentityTag apijson.Field
	ItemType               apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseMcnGcpSetup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseMcnGcpSetupJSON) RawJSON() string {
	return r.raw
}

func (r CloudIntegrationInitialSetupResponseMcnGcpSetup) implementsCloudIntegrationInitialSetupResponse() {
}

type CloudIntegrationNewParams struct {
	AccountID    param.Field[string]                             `path:"account_id,required"`
	CloudType    param.Field[CloudIntegrationNewParamsCloudType] `json:"cloud_type,required"`
	FriendlyName param.Field[string]                             `json:"friendly_name,required"`
	Description  param.Field[string]                             `json:"description"`
	Forwarded    param.Field[string]                             `header:"forwarded"`
}

func (r CloudIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudIntegrationNewParamsCloudType string

const (
	CloudIntegrationNewParamsCloudTypeAws        CloudIntegrationNewParamsCloudType = "AWS"
	CloudIntegrationNewParamsCloudTypeAzure      CloudIntegrationNewParamsCloudType = "AZURE"
	CloudIntegrationNewParamsCloudTypeGoogle     CloudIntegrationNewParamsCloudType = "GOOGLE"
	CloudIntegrationNewParamsCloudTypeCloudflare CloudIntegrationNewParamsCloudType = "CLOUDFLARE"
)

func (r CloudIntegrationNewParamsCloudType) IsKnown() bool {
	switch r {
	case CloudIntegrationNewParamsCloudTypeAws, CloudIntegrationNewParamsCloudTypeAzure, CloudIntegrationNewParamsCloudTypeGoogle, CloudIntegrationNewParamsCloudTypeCloudflare:
		return true
	}
	return false
}

type CloudIntegrationNewResponseEnvelope struct {
	Errors   []CloudIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationNewResponse                   `json:"result,required"`
	Success  bool                                          `json:"success,required"`
	JSON     cloudIntegrationNewResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudIntegrationNewResponseEnvelope]
type cloudIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeErrors struct {
	Code             CloudIntegrationNewResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Meta             CloudIntegrationNewResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CloudIntegrationNewResponseEnvelopeErrors]
type cloudIntegrationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationNewResponseEnvelopeErrorsCode1001   CloudIntegrationNewResponseEnvelopeErrorsCode = 1001
	CloudIntegrationNewResponseEnvelopeErrorsCode1002   CloudIntegrationNewResponseEnvelopeErrorsCode = 1002
	CloudIntegrationNewResponseEnvelopeErrorsCode1003   CloudIntegrationNewResponseEnvelopeErrorsCode = 1003
	CloudIntegrationNewResponseEnvelopeErrorsCode1004   CloudIntegrationNewResponseEnvelopeErrorsCode = 1004
	CloudIntegrationNewResponseEnvelopeErrorsCode1005   CloudIntegrationNewResponseEnvelopeErrorsCode = 1005
	CloudIntegrationNewResponseEnvelopeErrorsCode1006   CloudIntegrationNewResponseEnvelopeErrorsCode = 1006
	CloudIntegrationNewResponseEnvelopeErrorsCode1007   CloudIntegrationNewResponseEnvelopeErrorsCode = 1007
	CloudIntegrationNewResponseEnvelopeErrorsCode1008   CloudIntegrationNewResponseEnvelopeErrorsCode = 1008
	CloudIntegrationNewResponseEnvelopeErrorsCode1009   CloudIntegrationNewResponseEnvelopeErrorsCode = 1009
	CloudIntegrationNewResponseEnvelopeErrorsCode1010   CloudIntegrationNewResponseEnvelopeErrorsCode = 1010
	CloudIntegrationNewResponseEnvelopeErrorsCode1011   CloudIntegrationNewResponseEnvelopeErrorsCode = 1011
	CloudIntegrationNewResponseEnvelopeErrorsCode1012   CloudIntegrationNewResponseEnvelopeErrorsCode = 1012
	CloudIntegrationNewResponseEnvelopeErrorsCode1013   CloudIntegrationNewResponseEnvelopeErrorsCode = 1013
	CloudIntegrationNewResponseEnvelopeErrorsCode1014   CloudIntegrationNewResponseEnvelopeErrorsCode = 1014
	CloudIntegrationNewResponseEnvelopeErrorsCode1015   CloudIntegrationNewResponseEnvelopeErrorsCode = 1015
	CloudIntegrationNewResponseEnvelopeErrorsCode1016   CloudIntegrationNewResponseEnvelopeErrorsCode = 1016
	CloudIntegrationNewResponseEnvelopeErrorsCode1017   CloudIntegrationNewResponseEnvelopeErrorsCode = 1017
	CloudIntegrationNewResponseEnvelopeErrorsCode2001   CloudIntegrationNewResponseEnvelopeErrorsCode = 2001
	CloudIntegrationNewResponseEnvelopeErrorsCode2002   CloudIntegrationNewResponseEnvelopeErrorsCode = 2002
	CloudIntegrationNewResponseEnvelopeErrorsCode2003   CloudIntegrationNewResponseEnvelopeErrorsCode = 2003
	CloudIntegrationNewResponseEnvelopeErrorsCode2004   CloudIntegrationNewResponseEnvelopeErrorsCode = 2004
	CloudIntegrationNewResponseEnvelopeErrorsCode2005   CloudIntegrationNewResponseEnvelopeErrorsCode = 2005
	CloudIntegrationNewResponseEnvelopeErrorsCode2006   CloudIntegrationNewResponseEnvelopeErrorsCode = 2006
	CloudIntegrationNewResponseEnvelopeErrorsCode2007   CloudIntegrationNewResponseEnvelopeErrorsCode = 2007
	CloudIntegrationNewResponseEnvelopeErrorsCode2008   CloudIntegrationNewResponseEnvelopeErrorsCode = 2008
	CloudIntegrationNewResponseEnvelopeErrorsCode2009   CloudIntegrationNewResponseEnvelopeErrorsCode = 2009
	CloudIntegrationNewResponseEnvelopeErrorsCode2010   CloudIntegrationNewResponseEnvelopeErrorsCode = 2010
	CloudIntegrationNewResponseEnvelopeErrorsCode2011   CloudIntegrationNewResponseEnvelopeErrorsCode = 2011
	CloudIntegrationNewResponseEnvelopeErrorsCode2012   CloudIntegrationNewResponseEnvelopeErrorsCode = 2012
	CloudIntegrationNewResponseEnvelopeErrorsCode2013   CloudIntegrationNewResponseEnvelopeErrorsCode = 2013
	CloudIntegrationNewResponseEnvelopeErrorsCode2014   CloudIntegrationNewResponseEnvelopeErrorsCode = 2014
	CloudIntegrationNewResponseEnvelopeErrorsCode2015   CloudIntegrationNewResponseEnvelopeErrorsCode = 2015
	CloudIntegrationNewResponseEnvelopeErrorsCode2016   CloudIntegrationNewResponseEnvelopeErrorsCode = 2016
	CloudIntegrationNewResponseEnvelopeErrorsCode2017   CloudIntegrationNewResponseEnvelopeErrorsCode = 2017
	CloudIntegrationNewResponseEnvelopeErrorsCode2018   CloudIntegrationNewResponseEnvelopeErrorsCode = 2018
	CloudIntegrationNewResponseEnvelopeErrorsCode2019   CloudIntegrationNewResponseEnvelopeErrorsCode = 2019
	CloudIntegrationNewResponseEnvelopeErrorsCode2020   CloudIntegrationNewResponseEnvelopeErrorsCode = 2020
	CloudIntegrationNewResponseEnvelopeErrorsCode2021   CloudIntegrationNewResponseEnvelopeErrorsCode = 2021
	CloudIntegrationNewResponseEnvelopeErrorsCode2022   CloudIntegrationNewResponseEnvelopeErrorsCode = 2022
	CloudIntegrationNewResponseEnvelopeErrorsCode3001   CloudIntegrationNewResponseEnvelopeErrorsCode = 3001
	CloudIntegrationNewResponseEnvelopeErrorsCode3002   CloudIntegrationNewResponseEnvelopeErrorsCode = 3002
	CloudIntegrationNewResponseEnvelopeErrorsCode3003   CloudIntegrationNewResponseEnvelopeErrorsCode = 3003
	CloudIntegrationNewResponseEnvelopeErrorsCode3004   CloudIntegrationNewResponseEnvelopeErrorsCode = 3004
	CloudIntegrationNewResponseEnvelopeErrorsCode3005   CloudIntegrationNewResponseEnvelopeErrorsCode = 3005
	CloudIntegrationNewResponseEnvelopeErrorsCode3006   CloudIntegrationNewResponseEnvelopeErrorsCode = 3006
	CloudIntegrationNewResponseEnvelopeErrorsCode3007   CloudIntegrationNewResponseEnvelopeErrorsCode = 3007
	CloudIntegrationNewResponseEnvelopeErrorsCode4001   CloudIntegrationNewResponseEnvelopeErrorsCode = 4001
	CloudIntegrationNewResponseEnvelopeErrorsCode4002   CloudIntegrationNewResponseEnvelopeErrorsCode = 4002
	CloudIntegrationNewResponseEnvelopeErrorsCode4003   CloudIntegrationNewResponseEnvelopeErrorsCode = 4003
	CloudIntegrationNewResponseEnvelopeErrorsCode4004   CloudIntegrationNewResponseEnvelopeErrorsCode = 4004
	CloudIntegrationNewResponseEnvelopeErrorsCode4005   CloudIntegrationNewResponseEnvelopeErrorsCode = 4005
	CloudIntegrationNewResponseEnvelopeErrorsCode4006   CloudIntegrationNewResponseEnvelopeErrorsCode = 4006
	CloudIntegrationNewResponseEnvelopeErrorsCode4007   CloudIntegrationNewResponseEnvelopeErrorsCode = 4007
	CloudIntegrationNewResponseEnvelopeErrorsCode4008   CloudIntegrationNewResponseEnvelopeErrorsCode = 4008
	CloudIntegrationNewResponseEnvelopeErrorsCode4009   CloudIntegrationNewResponseEnvelopeErrorsCode = 4009
	CloudIntegrationNewResponseEnvelopeErrorsCode4010   CloudIntegrationNewResponseEnvelopeErrorsCode = 4010
	CloudIntegrationNewResponseEnvelopeErrorsCode4011   CloudIntegrationNewResponseEnvelopeErrorsCode = 4011
	CloudIntegrationNewResponseEnvelopeErrorsCode4012   CloudIntegrationNewResponseEnvelopeErrorsCode = 4012
	CloudIntegrationNewResponseEnvelopeErrorsCode4013   CloudIntegrationNewResponseEnvelopeErrorsCode = 4013
	CloudIntegrationNewResponseEnvelopeErrorsCode4014   CloudIntegrationNewResponseEnvelopeErrorsCode = 4014
	CloudIntegrationNewResponseEnvelopeErrorsCode4015   CloudIntegrationNewResponseEnvelopeErrorsCode = 4015
	CloudIntegrationNewResponseEnvelopeErrorsCode4016   CloudIntegrationNewResponseEnvelopeErrorsCode = 4016
	CloudIntegrationNewResponseEnvelopeErrorsCode4017   CloudIntegrationNewResponseEnvelopeErrorsCode = 4017
	CloudIntegrationNewResponseEnvelopeErrorsCode4018   CloudIntegrationNewResponseEnvelopeErrorsCode = 4018
	CloudIntegrationNewResponseEnvelopeErrorsCode4019   CloudIntegrationNewResponseEnvelopeErrorsCode = 4019
	CloudIntegrationNewResponseEnvelopeErrorsCode4020   CloudIntegrationNewResponseEnvelopeErrorsCode = 4020
	CloudIntegrationNewResponseEnvelopeErrorsCode4021   CloudIntegrationNewResponseEnvelopeErrorsCode = 4021
	CloudIntegrationNewResponseEnvelopeErrorsCode4022   CloudIntegrationNewResponseEnvelopeErrorsCode = 4022
	CloudIntegrationNewResponseEnvelopeErrorsCode4023   CloudIntegrationNewResponseEnvelopeErrorsCode = 4023
	CloudIntegrationNewResponseEnvelopeErrorsCode5001   CloudIntegrationNewResponseEnvelopeErrorsCode = 5001
	CloudIntegrationNewResponseEnvelopeErrorsCode5002   CloudIntegrationNewResponseEnvelopeErrorsCode = 5002
	CloudIntegrationNewResponseEnvelopeErrorsCode5003   CloudIntegrationNewResponseEnvelopeErrorsCode = 5003
	CloudIntegrationNewResponseEnvelopeErrorsCode5004   CloudIntegrationNewResponseEnvelopeErrorsCode = 5004
	CloudIntegrationNewResponseEnvelopeErrorsCode102000 CloudIntegrationNewResponseEnvelopeErrorsCode = 102000
	CloudIntegrationNewResponseEnvelopeErrorsCode102001 CloudIntegrationNewResponseEnvelopeErrorsCode = 102001
	CloudIntegrationNewResponseEnvelopeErrorsCode102002 CloudIntegrationNewResponseEnvelopeErrorsCode = 102002
	CloudIntegrationNewResponseEnvelopeErrorsCode102003 CloudIntegrationNewResponseEnvelopeErrorsCode = 102003
	CloudIntegrationNewResponseEnvelopeErrorsCode102004 CloudIntegrationNewResponseEnvelopeErrorsCode = 102004
	CloudIntegrationNewResponseEnvelopeErrorsCode102005 CloudIntegrationNewResponseEnvelopeErrorsCode = 102005
	CloudIntegrationNewResponseEnvelopeErrorsCode102006 CloudIntegrationNewResponseEnvelopeErrorsCode = 102006
	CloudIntegrationNewResponseEnvelopeErrorsCode102007 CloudIntegrationNewResponseEnvelopeErrorsCode = 102007
	CloudIntegrationNewResponseEnvelopeErrorsCode102008 CloudIntegrationNewResponseEnvelopeErrorsCode = 102008
	CloudIntegrationNewResponseEnvelopeErrorsCode102009 CloudIntegrationNewResponseEnvelopeErrorsCode = 102009
	CloudIntegrationNewResponseEnvelopeErrorsCode102010 CloudIntegrationNewResponseEnvelopeErrorsCode = 102010
	CloudIntegrationNewResponseEnvelopeErrorsCode102011 CloudIntegrationNewResponseEnvelopeErrorsCode = 102011
	CloudIntegrationNewResponseEnvelopeErrorsCode102012 CloudIntegrationNewResponseEnvelopeErrorsCode = 102012
	CloudIntegrationNewResponseEnvelopeErrorsCode102013 CloudIntegrationNewResponseEnvelopeErrorsCode = 102013
	CloudIntegrationNewResponseEnvelopeErrorsCode102014 CloudIntegrationNewResponseEnvelopeErrorsCode = 102014
	CloudIntegrationNewResponseEnvelopeErrorsCode102015 CloudIntegrationNewResponseEnvelopeErrorsCode = 102015
	CloudIntegrationNewResponseEnvelopeErrorsCode102016 CloudIntegrationNewResponseEnvelopeErrorsCode = 102016
	CloudIntegrationNewResponseEnvelopeErrorsCode102017 CloudIntegrationNewResponseEnvelopeErrorsCode = 102017
	CloudIntegrationNewResponseEnvelopeErrorsCode102018 CloudIntegrationNewResponseEnvelopeErrorsCode = 102018
	CloudIntegrationNewResponseEnvelopeErrorsCode102019 CloudIntegrationNewResponseEnvelopeErrorsCode = 102019
	CloudIntegrationNewResponseEnvelopeErrorsCode102020 CloudIntegrationNewResponseEnvelopeErrorsCode = 102020
	CloudIntegrationNewResponseEnvelopeErrorsCode102021 CloudIntegrationNewResponseEnvelopeErrorsCode = 102021
	CloudIntegrationNewResponseEnvelopeErrorsCode102022 CloudIntegrationNewResponseEnvelopeErrorsCode = 102022
	CloudIntegrationNewResponseEnvelopeErrorsCode102023 CloudIntegrationNewResponseEnvelopeErrorsCode = 102023
	CloudIntegrationNewResponseEnvelopeErrorsCode102024 CloudIntegrationNewResponseEnvelopeErrorsCode = 102024
	CloudIntegrationNewResponseEnvelopeErrorsCode102025 CloudIntegrationNewResponseEnvelopeErrorsCode = 102025
	CloudIntegrationNewResponseEnvelopeErrorsCode102026 CloudIntegrationNewResponseEnvelopeErrorsCode = 102026
	CloudIntegrationNewResponseEnvelopeErrorsCode102027 CloudIntegrationNewResponseEnvelopeErrorsCode = 102027
	CloudIntegrationNewResponseEnvelopeErrorsCode102028 CloudIntegrationNewResponseEnvelopeErrorsCode = 102028
	CloudIntegrationNewResponseEnvelopeErrorsCode102029 CloudIntegrationNewResponseEnvelopeErrorsCode = 102029
	CloudIntegrationNewResponseEnvelopeErrorsCode102030 CloudIntegrationNewResponseEnvelopeErrorsCode = 102030
	CloudIntegrationNewResponseEnvelopeErrorsCode102031 CloudIntegrationNewResponseEnvelopeErrorsCode = 102031
	CloudIntegrationNewResponseEnvelopeErrorsCode102032 CloudIntegrationNewResponseEnvelopeErrorsCode = 102032
	CloudIntegrationNewResponseEnvelopeErrorsCode102033 CloudIntegrationNewResponseEnvelopeErrorsCode = 102033
	CloudIntegrationNewResponseEnvelopeErrorsCode102034 CloudIntegrationNewResponseEnvelopeErrorsCode = 102034
	CloudIntegrationNewResponseEnvelopeErrorsCode102035 CloudIntegrationNewResponseEnvelopeErrorsCode = 102035
	CloudIntegrationNewResponseEnvelopeErrorsCode102036 CloudIntegrationNewResponseEnvelopeErrorsCode = 102036
	CloudIntegrationNewResponseEnvelopeErrorsCode102037 CloudIntegrationNewResponseEnvelopeErrorsCode = 102037
	CloudIntegrationNewResponseEnvelopeErrorsCode102038 CloudIntegrationNewResponseEnvelopeErrorsCode = 102038
	CloudIntegrationNewResponseEnvelopeErrorsCode102039 CloudIntegrationNewResponseEnvelopeErrorsCode = 102039
	CloudIntegrationNewResponseEnvelopeErrorsCode102040 CloudIntegrationNewResponseEnvelopeErrorsCode = 102040
	CloudIntegrationNewResponseEnvelopeErrorsCode102041 CloudIntegrationNewResponseEnvelopeErrorsCode = 102041
	CloudIntegrationNewResponseEnvelopeErrorsCode102042 CloudIntegrationNewResponseEnvelopeErrorsCode = 102042
	CloudIntegrationNewResponseEnvelopeErrorsCode102043 CloudIntegrationNewResponseEnvelopeErrorsCode = 102043
	CloudIntegrationNewResponseEnvelopeErrorsCode102044 CloudIntegrationNewResponseEnvelopeErrorsCode = 102044
	CloudIntegrationNewResponseEnvelopeErrorsCode102045 CloudIntegrationNewResponseEnvelopeErrorsCode = 102045
	CloudIntegrationNewResponseEnvelopeErrorsCode102046 CloudIntegrationNewResponseEnvelopeErrorsCode = 102046
	CloudIntegrationNewResponseEnvelopeErrorsCode102047 CloudIntegrationNewResponseEnvelopeErrorsCode = 102047
	CloudIntegrationNewResponseEnvelopeErrorsCode102048 CloudIntegrationNewResponseEnvelopeErrorsCode = 102048
	CloudIntegrationNewResponseEnvelopeErrorsCode102049 CloudIntegrationNewResponseEnvelopeErrorsCode = 102049
	CloudIntegrationNewResponseEnvelopeErrorsCode102050 CloudIntegrationNewResponseEnvelopeErrorsCode = 102050
	CloudIntegrationNewResponseEnvelopeErrorsCode102051 CloudIntegrationNewResponseEnvelopeErrorsCode = 102051
	CloudIntegrationNewResponseEnvelopeErrorsCode102052 CloudIntegrationNewResponseEnvelopeErrorsCode = 102052
	CloudIntegrationNewResponseEnvelopeErrorsCode102053 CloudIntegrationNewResponseEnvelopeErrorsCode = 102053
	CloudIntegrationNewResponseEnvelopeErrorsCode102054 CloudIntegrationNewResponseEnvelopeErrorsCode = 102054
	CloudIntegrationNewResponseEnvelopeErrorsCode102055 CloudIntegrationNewResponseEnvelopeErrorsCode = 102055
	CloudIntegrationNewResponseEnvelopeErrorsCode102056 CloudIntegrationNewResponseEnvelopeErrorsCode = 102056
	CloudIntegrationNewResponseEnvelopeErrorsCode102057 CloudIntegrationNewResponseEnvelopeErrorsCode = 102057
	CloudIntegrationNewResponseEnvelopeErrorsCode102058 CloudIntegrationNewResponseEnvelopeErrorsCode = 102058
	CloudIntegrationNewResponseEnvelopeErrorsCode102059 CloudIntegrationNewResponseEnvelopeErrorsCode = 102059
	CloudIntegrationNewResponseEnvelopeErrorsCode102060 CloudIntegrationNewResponseEnvelopeErrorsCode = 102060
	CloudIntegrationNewResponseEnvelopeErrorsCode102061 CloudIntegrationNewResponseEnvelopeErrorsCode = 102061
	CloudIntegrationNewResponseEnvelopeErrorsCode102062 CloudIntegrationNewResponseEnvelopeErrorsCode = 102062
	CloudIntegrationNewResponseEnvelopeErrorsCode102063 CloudIntegrationNewResponseEnvelopeErrorsCode = 102063
	CloudIntegrationNewResponseEnvelopeErrorsCode102064 CloudIntegrationNewResponseEnvelopeErrorsCode = 102064
	CloudIntegrationNewResponseEnvelopeErrorsCode102065 CloudIntegrationNewResponseEnvelopeErrorsCode = 102065
	CloudIntegrationNewResponseEnvelopeErrorsCode102066 CloudIntegrationNewResponseEnvelopeErrorsCode = 102066
	CloudIntegrationNewResponseEnvelopeErrorsCode103001 CloudIntegrationNewResponseEnvelopeErrorsCode = 103001
	CloudIntegrationNewResponseEnvelopeErrorsCode103002 CloudIntegrationNewResponseEnvelopeErrorsCode = 103002
	CloudIntegrationNewResponseEnvelopeErrorsCode103003 CloudIntegrationNewResponseEnvelopeErrorsCode = 103003
	CloudIntegrationNewResponseEnvelopeErrorsCode103004 CloudIntegrationNewResponseEnvelopeErrorsCode = 103004
	CloudIntegrationNewResponseEnvelopeErrorsCode103005 CloudIntegrationNewResponseEnvelopeErrorsCode = 103005
	CloudIntegrationNewResponseEnvelopeErrorsCode103006 CloudIntegrationNewResponseEnvelopeErrorsCode = 103006
	CloudIntegrationNewResponseEnvelopeErrorsCode103007 CloudIntegrationNewResponseEnvelopeErrorsCode = 103007
	CloudIntegrationNewResponseEnvelopeErrorsCode103008 CloudIntegrationNewResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationNewResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseEnvelopeErrorsCode1001, CloudIntegrationNewResponseEnvelopeErrorsCode1002, CloudIntegrationNewResponseEnvelopeErrorsCode1003, CloudIntegrationNewResponseEnvelopeErrorsCode1004, CloudIntegrationNewResponseEnvelopeErrorsCode1005, CloudIntegrationNewResponseEnvelopeErrorsCode1006, CloudIntegrationNewResponseEnvelopeErrorsCode1007, CloudIntegrationNewResponseEnvelopeErrorsCode1008, CloudIntegrationNewResponseEnvelopeErrorsCode1009, CloudIntegrationNewResponseEnvelopeErrorsCode1010, CloudIntegrationNewResponseEnvelopeErrorsCode1011, CloudIntegrationNewResponseEnvelopeErrorsCode1012, CloudIntegrationNewResponseEnvelopeErrorsCode1013, CloudIntegrationNewResponseEnvelopeErrorsCode1014, CloudIntegrationNewResponseEnvelopeErrorsCode1015, CloudIntegrationNewResponseEnvelopeErrorsCode1016, CloudIntegrationNewResponseEnvelopeErrorsCode1017, CloudIntegrationNewResponseEnvelopeErrorsCode2001, CloudIntegrationNewResponseEnvelopeErrorsCode2002, CloudIntegrationNewResponseEnvelopeErrorsCode2003, CloudIntegrationNewResponseEnvelopeErrorsCode2004, CloudIntegrationNewResponseEnvelopeErrorsCode2005, CloudIntegrationNewResponseEnvelopeErrorsCode2006, CloudIntegrationNewResponseEnvelopeErrorsCode2007, CloudIntegrationNewResponseEnvelopeErrorsCode2008, CloudIntegrationNewResponseEnvelopeErrorsCode2009, CloudIntegrationNewResponseEnvelopeErrorsCode2010, CloudIntegrationNewResponseEnvelopeErrorsCode2011, CloudIntegrationNewResponseEnvelopeErrorsCode2012, CloudIntegrationNewResponseEnvelopeErrorsCode2013, CloudIntegrationNewResponseEnvelopeErrorsCode2014, CloudIntegrationNewResponseEnvelopeErrorsCode2015, CloudIntegrationNewResponseEnvelopeErrorsCode2016, CloudIntegrationNewResponseEnvelopeErrorsCode2017, CloudIntegrationNewResponseEnvelopeErrorsCode2018, CloudIntegrationNewResponseEnvelopeErrorsCode2019, CloudIntegrationNewResponseEnvelopeErrorsCode2020, CloudIntegrationNewResponseEnvelopeErrorsCode2021, CloudIntegrationNewResponseEnvelopeErrorsCode2022, CloudIntegrationNewResponseEnvelopeErrorsCode3001, CloudIntegrationNewResponseEnvelopeErrorsCode3002, CloudIntegrationNewResponseEnvelopeErrorsCode3003, CloudIntegrationNewResponseEnvelopeErrorsCode3004, CloudIntegrationNewResponseEnvelopeErrorsCode3005, CloudIntegrationNewResponseEnvelopeErrorsCode3006, CloudIntegrationNewResponseEnvelopeErrorsCode3007, CloudIntegrationNewResponseEnvelopeErrorsCode4001, CloudIntegrationNewResponseEnvelopeErrorsCode4002, CloudIntegrationNewResponseEnvelopeErrorsCode4003, CloudIntegrationNewResponseEnvelopeErrorsCode4004, CloudIntegrationNewResponseEnvelopeErrorsCode4005, CloudIntegrationNewResponseEnvelopeErrorsCode4006, CloudIntegrationNewResponseEnvelopeErrorsCode4007, CloudIntegrationNewResponseEnvelopeErrorsCode4008, CloudIntegrationNewResponseEnvelopeErrorsCode4009, CloudIntegrationNewResponseEnvelopeErrorsCode4010, CloudIntegrationNewResponseEnvelopeErrorsCode4011, CloudIntegrationNewResponseEnvelopeErrorsCode4012, CloudIntegrationNewResponseEnvelopeErrorsCode4013, CloudIntegrationNewResponseEnvelopeErrorsCode4014, CloudIntegrationNewResponseEnvelopeErrorsCode4015, CloudIntegrationNewResponseEnvelopeErrorsCode4016, CloudIntegrationNewResponseEnvelopeErrorsCode4017, CloudIntegrationNewResponseEnvelopeErrorsCode4018, CloudIntegrationNewResponseEnvelopeErrorsCode4019, CloudIntegrationNewResponseEnvelopeErrorsCode4020, CloudIntegrationNewResponseEnvelopeErrorsCode4021, CloudIntegrationNewResponseEnvelopeErrorsCode4022, CloudIntegrationNewResponseEnvelopeErrorsCode4023, CloudIntegrationNewResponseEnvelopeErrorsCode5001, CloudIntegrationNewResponseEnvelopeErrorsCode5002, CloudIntegrationNewResponseEnvelopeErrorsCode5003, CloudIntegrationNewResponseEnvelopeErrorsCode5004, CloudIntegrationNewResponseEnvelopeErrorsCode102000, CloudIntegrationNewResponseEnvelopeErrorsCode102001, CloudIntegrationNewResponseEnvelopeErrorsCode102002, CloudIntegrationNewResponseEnvelopeErrorsCode102003, CloudIntegrationNewResponseEnvelopeErrorsCode102004, CloudIntegrationNewResponseEnvelopeErrorsCode102005, CloudIntegrationNewResponseEnvelopeErrorsCode102006, CloudIntegrationNewResponseEnvelopeErrorsCode102007, CloudIntegrationNewResponseEnvelopeErrorsCode102008, CloudIntegrationNewResponseEnvelopeErrorsCode102009, CloudIntegrationNewResponseEnvelopeErrorsCode102010, CloudIntegrationNewResponseEnvelopeErrorsCode102011, CloudIntegrationNewResponseEnvelopeErrorsCode102012, CloudIntegrationNewResponseEnvelopeErrorsCode102013, CloudIntegrationNewResponseEnvelopeErrorsCode102014, CloudIntegrationNewResponseEnvelopeErrorsCode102015, CloudIntegrationNewResponseEnvelopeErrorsCode102016, CloudIntegrationNewResponseEnvelopeErrorsCode102017, CloudIntegrationNewResponseEnvelopeErrorsCode102018, CloudIntegrationNewResponseEnvelopeErrorsCode102019, CloudIntegrationNewResponseEnvelopeErrorsCode102020, CloudIntegrationNewResponseEnvelopeErrorsCode102021, CloudIntegrationNewResponseEnvelopeErrorsCode102022, CloudIntegrationNewResponseEnvelopeErrorsCode102023, CloudIntegrationNewResponseEnvelopeErrorsCode102024, CloudIntegrationNewResponseEnvelopeErrorsCode102025, CloudIntegrationNewResponseEnvelopeErrorsCode102026, CloudIntegrationNewResponseEnvelopeErrorsCode102027, CloudIntegrationNewResponseEnvelopeErrorsCode102028, CloudIntegrationNewResponseEnvelopeErrorsCode102029, CloudIntegrationNewResponseEnvelopeErrorsCode102030, CloudIntegrationNewResponseEnvelopeErrorsCode102031, CloudIntegrationNewResponseEnvelopeErrorsCode102032, CloudIntegrationNewResponseEnvelopeErrorsCode102033, CloudIntegrationNewResponseEnvelopeErrorsCode102034, CloudIntegrationNewResponseEnvelopeErrorsCode102035, CloudIntegrationNewResponseEnvelopeErrorsCode102036, CloudIntegrationNewResponseEnvelopeErrorsCode102037, CloudIntegrationNewResponseEnvelopeErrorsCode102038, CloudIntegrationNewResponseEnvelopeErrorsCode102039, CloudIntegrationNewResponseEnvelopeErrorsCode102040, CloudIntegrationNewResponseEnvelopeErrorsCode102041, CloudIntegrationNewResponseEnvelopeErrorsCode102042, CloudIntegrationNewResponseEnvelopeErrorsCode102043, CloudIntegrationNewResponseEnvelopeErrorsCode102044, CloudIntegrationNewResponseEnvelopeErrorsCode102045, CloudIntegrationNewResponseEnvelopeErrorsCode102046, CloudIntegrationNewResponseEnvelopeErrorsCode102047, CloudIntegrationNewResponseEnvelopeErrorsCode102048, CloudIntegrationNewResponseEnvelopeErrorsCode102049, CloudIntegrationNewResponseEnvelopeErrorsCode102050, CloudIntegrationNewResponseEnvelopeErrorsCode102051, CloudIntegrationNewResponseEnvelopeErrorsCode102052, CloudIntegrationNewResponseEnvelopeErrorsCode102053, CloudIntegrationNewResponseEnvelopeErrorsCode102054, CloudIntegrationNewResponseEnvelopeErrorsCode102055, CloudIntegrationNewResponseEnvelopeErrorsCode102056, CloudIntegrationNewResponseEnvelopeErrorsCode102057, CloudIntegrationNewResponseEnvelopeErrorsCode102058, CloudIntegrationNewResponseEnvelopeErrorsCode102059, CloudIntegrationNewResponseEnvelopeErrorsCode102060, CloudIntegrationNewResponseEnvelopeErrorsCode102061, CloudIntegrationNewResponseEnvelopeErrorsCode102062, CloudIntegrationNewResponseEnvelopeErrorsCode102063, CloudIntegrationNewResponseEnvelopeErrorsCode102064, CloudIntegrationNewResponseEnvelopeErrorsCode102065, CloudIntegrationNewResponseEnvelopeErrorsCode102066, CloudIntegrationNewResponseEnvelopeErrorsCode103001, CloudIntegrationNewResponseEnvelopeErrorsCode103002, CloudIntegrationNewResponseEnvelopeErrorsCode103003, CloudIntegrationNewResponseEnvelopeErrorsCode103004, CloudIntegrationNewResponseEnvelopeErrorsCode103005, CloudIntegrationNewResponseEnvelopeErrorsCode103006, CloudIntegrationNewResponseEnvelopeErrorsCode103007, CloudIntegrationNewResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationNewResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                            `json:"l10n_key"`
	LoggableError string                                            `json:"loggable_error"`
	TemplateData  interface{}                                       `json:"template_data"`
	TraceID       string                                            `json:"trace_id"`
	JSON          cloudIntegrationNewResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeErrorsMetaJSON contains the JSON metadata for
// the struct [CloudIntegrationNewResponseEnvelopeErrorsMeta]
type cloudIntegrationNewResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeErrorsSource struct {
	Parameter           string                                              `json:"parameter"`
	ParameterValueIndex int64                                               `json:"parameter_value_index"`
	Pointer             string                                              `json:"pointer"`
	JSON                cloudIntegrationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationNewResponseEnvelopeErrorsSource]
type cloudIntegrationNewResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeMessages struct {
	Code             CloudIntegrationNewResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Meta             CloudIntegrationNewResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CloudIntegrationNewResponseEnvelopeMessages]
type cloudIntegrationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationNewResponseEnvelopeMessagesCode1001   CloudIntegrationNewResponseEnvelopeMessagesCode = 1001
	CloudIntegrationNewResponseEnvelopeMessagesCode1002   CloudIntegrationNewResponseEnvelopeMessagesCode = 1002
	CloudIntegrationNewResponseEnvelopeMessagesCode1003   CloudIntegrationNewResponseEnvelopeMessagesCode = 1003
	CloudIntegrationNewResponseEnvelopeMessagesCode1004   CloudIntegrationNewResponseEnvelopeMessagesCode = 1004
	CloudIntegrationNewResponseEnvelopeMessagesCode1005   CloudIntegrationNewResponseEnvelopeMessagesCode = 1005
	CloudIntegrationNewResponseEnvelopeMessagesCode1006   CloudIntegrationNewResponseEnvelopeMessagesCode = 1006
	CloudIntegrationNewResponseEnvelopeMessagesCode1007   CloudIntegrationNewResponseEnvelopeMessagesCode = 1007
	CloudIntegrationNewResponseEnvelopeMessagesCode1008   CloudIntegrationNewResponseEnvelopeMessagesCode = 1008
	CloudIntegrationNewResponseEnvelopeMessagesCode1009   CloudIntegrationNewResponseEnvelopeMessagesCode = 1009
	CloudIntegrationNewResponseEnvelopeMessagesCode1010   CloudIntegrationNewResponseEnvelopeMessagesCode = 1010
	CloudIntegrationNewResponseEnvelopeMessagesCode1011   CloudIntegrationNewResponseEnvelopeMessagesCode = 1011
	CloudIntegrationNewResponseEnvelopeMessagesCode1012   CloudIntegrationNewResponseEnvelopeMessagesCode = 1012
	CloudIntegrationNewResponseEnvelopeMessagesCode1013   CloudIntegrationNewResponseEnvelopeMessagesCode = 1013
	CloudIntegrationNewResponseEnvelopeMessagesCode1014   CloudIntegrationNewResponseEnvelopeMessagesCode = 1014
	CloudIntegrationNewResponseEnvelopeMessagesCode1015   CloudIntegrationNewResponseEnvelopeMessagesCode = 1015
	CloudIntegrationNewResponseEnvelopeMessagesCode1016   CloudIntegrationNewResponseEnvelopeMessagesCode = 1016
	CloudIntegrationNewResponseEnvelopeMessagesCode1017   CloudIntegrationNewResponseEnvelopeMessagesCode = 1017
	CloudIntegrationNewResponseEnvelopeMessagesCode2001   CloudIntegrationNewResponseEnvelopeMessagesCode = 2001
	CloudIntegrationNewResponseEnvelopeMessagesCode2002   CloudIntegrationNewResponseEnvelopeMessagesCode = 2002
	CloudIntegrationNewResponseEnvelopeMessagesCode2003   CloudIntegrationNewResponseEnvelopeMessagesCode = 2003
	CloudIntegrationNewResponseEnvelopeMessagesCode2004   CloudIntegrationNewResponseEnvelopeMessagesCode = 2004
	CloudIntegrationNewResponseEnvelopeMessagesCode2005   CloudIntegrationNewResponseEnvelopeMessagesCode = 2005
	CloudIntegrationNewResponseEnvelopeMessagesCode2006   CloudIntegrationNewResponseEnvelopeMessagesCode = 2006
	CloudIntegrationNewResponseEnvelopeMessagesCode2007   CloudIntegrationNewResponseEnvelopeMessagesCode = 2007
	CloudIntegrationNewResponseEnvelopeMessagesCode2008   CloudIntegrationNewResponseEnvelopeMessagesCode = 2008
	CloudIntegrationNewResponseEnvelopeMessagesCode2009   CloudIntegrationNewResponseEnvelopeMessagesCode = 2009
	CloudIntegrationNewResponseEnvelopeMessagesCode2010   CloudIntegrationNewResponseEnvelopeMessagesCode = 2010
	CloudIntegrationNewResponseEnvelopeMessagesCode2011   CloudIntegrationNewResponseEnvelopeMessagesCode = 2011
	CloudIntegrationNewResponseEnvelopeMessagesCode2012   CloudIntegrationNewResponseEnvelopeMessagesCode = 2012
	CloudIntegrationNewResponseEnvelopeMessagesCode2013   CloudIntegrationNewResponseEnvelopeMessagesCode = 2013
	CloudIntegrationNewResponseEnvelopeMessagesCode2014   CloudIntegrationNewResponseEnvelopeMessagesCode = 2014
	CloudIntegrationNewResponseEnvelopeMessagesCode2015   CloudIntegrationNewResponseEnvelopeMessagesCode = 2015
	CloudIntegrationNewResponseEnvelopeMessagesCode2016   CloudIntegrationNewResponseEnvelopeMessagesCode = 2016
	CloudIntegrationNewResponseEnvelopeMessagesCode2017   CloudIntegrationNewResponseEnvelopeMessagesCode = 2017
	CloudIntegrationNewResponseEnvelopeMessagesCode2018   CloudIntegrationNewResponseEnvelopeMessagesCode = 2018
	CloudIntegrationNewResponseEnvelopeMessagesCode2019   CloudIntegrationNewResponseEnvelopeMessagesCode = 2019
	CloudIntegrationNewResponseEnvelopeMessagesCode2020   CloudIntegrationNewResponseEnvelopeMessagesCode = 2020
	CloudIntegrationNewResponseEnvelopeMessagesCode2021   CloudIntegrationNewResponseEnvelopeMessagesCode = 2021
	CloudIntegrationNewResponseEnvelopeMessagesCode2022   CloudIntegrationNewResponseEnvelopeMessagesCode = 2022
	CloudIntegrationNewResponseEnvelopeMessagesCode3001   CloudIntegrationNewResponseEnvelopeMessagesCode = 3001
	CloudIntegrationNewResponseEnvelopeMessagesCode3002   CloudIntegrationNewResponseEnvelopeMessagesCode = 3002
	CloudIntegrationNewResponseEnvelopeMessagesCode3003   CloudIntegrationNewResponseEnvelopeMessagesCode = 3003
	CloudIntegrationNewResponseEnvelopeMessagesCode3004   CloudIntegrationNewResponseEnvelopeMessagesCode = 3004
	CloudIntegrationNewResponseEnvelopeMessagesCode3005   CloudIntegrationNewResponseEnvelopeMessagesCode = 3005
	CloudIntegrationNewResponseEnvelopeMessagesCode3006   CloudIntegrationNewResponseEnvelopeMessagesCode = 3006
	CloudIntegrationNewResponseEnvelopeMessagesCode3007   CloudIntegrationNewResponseEnvelopeMessagesCode = 3007
	CloudIntegrationNewResponseEnvelopeMessagesCode4001   CloudIntegrationNewResponseEnvelopeMessagesCode = 4001
	CloudIntegrationNewResponseEnvelopeMessagesCode4002   CloudIntegrationNewResponseEnvelopeMessagesCode = 4002
	CloudIntegrationNewResponseEnvelopeMessagesCode4003   CloudIntegrationNewResponseEnvelopeMessagesCode = 4003
	CloudIntegrationNewResponseEnvelopeMessagesCode4004   CloudIntegrationNewResponseEnvelopeMessagesCode = 4004
	CloudIntegrationNewResponseEnvelopeMessagesCode4005   CloudIntegrationNewResponseEnvelopeMessagesCode = 4005
	CloudIntegrationNewResponseEnvelopeMessagesCode4006   CloudIntegrationNewResponseEnvelopeMessagesCode = 4006
	CloudIntegrationNewResponseEnvelopeMessagesCode4007   CloudIntegrationNewResponseEnvelopeMessagesCode = 4007
	CloudIntegrationNewResponseEnvelopeMessagesCode4008   CloudIntegrationNewResponseEnvelopeMessagesCode = 4008
	CloudIntegrationNewResponseEnvelopeMessagesCode4009   CloudIntegrationNewResponseEnvelopeMessagesCode = 4009
	CloudIntegrationNewResponseEnvelopeMessagesCode4010   CloudIntegrationNewResponseEnvelopeMessagesCode = 4010
	CloudIntegrationNewResponseEnvelopeMessagesCode4011   CloudIntegrationNewResponseEnvelopeMessagesCode = 4011
	CloudIntegrationNewResponseEnvelopeMessagesCode4012   CloudIntegrationNewResponseEnvelopeMessagesCode = 4012
	CloudIntegrationNewResponseEnvelopeMessagesCode4013   CloudIntegrationNewResponseEnvelopeMessagesCode = 4013
	CloudIntegrationNewResponseEnvelopeMessagesCode4014   CloudIntegrationNewResponseEnvelopeMessagesCode = 4014
	CloudIntegrationNewResponseEnvelopeMessagesCode4015   CloudIntegrationNewResponseEnvelopeMessagesCode = 4015
	CloudIntegrationNewResponseEnvelopeMessagesCode4016   CloudIntegrationNewResponseEnvelopeMessagesCode = 4016
	CloudIntegrationNewResponseEnvelopeMessagesCode4017   CloudIntegrationNewResponseEnvelopeMessagesCode = 4017
	CloudIntegrationNewResponseEnvelopeMessagesCode4018   CloudIntegrationNewResponseEnvelopeMessagesCode = 4018
	CloudIntegrationNewResponseEnvelopeMessagesCode4019   CloudIntegrationNewResponseEnvelopeMessagesCode = 4019
	CloudIntegrationNewResponseEnvelopeMessagesCode4020   CloudIntegrationNewResponseEnvelopeMessagesCode = 4020
	CloudIntegrationNewResponseEnvelopeMessagesCode4021   CloudIntegrationNewResponseEnvelopeMessagesCode = 4021
	CloudIntegrationNewResponseEnvelopeMessagesCode4022   CloudIntegrationNewResponseEnvelopeMessagesCode = 4022
	CloudIntegrationNewResponseEnvelopeMessagesCode4023   CloudIntegrationNewResponseEnvelopeMessagesCode = 4023
	CloudIntegrationNewResponseEnvelopeMessagesCode5001   CloudIntegrationNewResponseEnvelopeMessagesCode = 5001
	CloudIntegrationNewResponseEnvelopeMessagesCode5002   CloudIntegrationNewResponseEnvelopeMessagesCode = 5002
	CloudIntegrationNewResponseEnvelopeMessagesCode5003   CloudIntegrationNewResponseEnvelopeMessagesCode = 5003
	CloudIntegrationNewResponseEnvelopeMessagesCode5004   CloudIntegrationNewResponseEnvelopeMessagesCode = 5004
	CloudIntegrationNewResponseEnvelopeMessagesCode102000 CloudIntegrationNewResponseEnvelopeMessagesCode = 102000
	CloudIntegrationNewResponseEnvelopeMessagesCode102001 CloudIntegrationNewResponseEnvelopeMessagesCode = 102001
	CloudIntegrationNewResponseEnvelopeMessagesCode102002 CloudIntegrationNewResponseEnvelopeMessagesCode = 102002
	CloudIntegrationNewResponseEnvelopeMessagesCode102003 CloudIntegrationNewResponseEnvelopeMessagesCode = 102003
	CloudIntegrationNewResponseEnvelopeMessagesCode102004 CloudIntegrationNewResponseEnvelopeMessagesCode = 102004
	CloudIntegrationNewResponseEnvelopeMessagesCode102005 CloudIntegrationNewResponseEnvelopeMessagesCode = 102005
	CloudIntegrationNewResponseEnvelopeMessagesCode102006 CloudIntegrationNewResponseEnvelopeMessagesCode = 102006
	CloudIntegrationNewResponseEnvelopeMessagesCode102007 CloudIntegrationNewResponseEnvelopeMessagesCode = 102007
	CloudIntegrationNewResponseEnvelopeMessagesCode102008 CloudIntegrationNewResponseEnvelopeMessagesCode = 102008
	CloudIntegrationNewResponseEnvelopeMessagesCode102009 CloudIntegrationNewResponseEnvelopeMessagesCode = 102009
	CloudIntegrationNewResponseEnvelopeMessagesCode102010 CloudIntegrationNewResponseEnvelopeMessagesCode = 102010
	CloudIntegrationNewResponseEnvelopeMessagesCode102011 CloudIntegrationNewResponseEnvelopeMessagesCode = 102011
	CloudIntegrationNewResponseEnvelopeMessagesCode102012 CloudIntegrationNewResponseEnvelopeMessagesCode = 102012
	CloudIntegrationNewResponseEnvelopeMessagesCode102013 CloudIntegrationNewResponseEnvelopeMessagesCode = 102013
	CloudIntegrationNewResponseEnvelopeMessagesCode102014 CloudIntegrationNewResponseEnvelopeMessagesCode = 102014
	CloudIntegrationNewResponseEnvelopeMessagesCode102015 CloudIntegrationNewResponseEnvelopeMessagesCode = 102015
	CloudIntegrationNewResponseEnvelopeMessagesCode102016 CloudIntegrationNewResponseEnvelopeMessagesCode = 102016
	CloudIntegrationNewResponseEnvelopeMessagesCode102017 CloudIntegrationNewResponseEnvelopeMessagesCode = 102017
	CloudIntegrationNewResponseEnvelopeMessagesCode102018 CloudIntegrationNewResponseEnvelopeMessagesCode = 102018
	CloudIntegrationNewResponseEnvelopeMessagesCode102019 CloudIntegrationNewResponseEnvelopeMessagesCode = 102019
	CloudIntegrationNewResponseEnvelopeMessagesCode102020 CloudIntegrationNewResponseEnvelopeMessagesCode = 102020
	CloudIntegrationNewResponseEnvelopeMessagesCode102021 CloudIntegrationNewResponseEnvelopeMessagesCode = 102021
	CloudIntegrationNewResponseEnvelopeMessagesCode102022 CloudIntegrationNewResponseEnvelopeMessagesCode = 102022
	CloudIntegrationNewResponseEnvelopeMessagesCode102023 CloudIntegrationNewResponseEnvelopeMessagesCode = 102023
	CloudIntegrationNewResponseEnvelopeMessagesCode102024 CloudIntegrationNewResponseEnvelopeMessagesCode = 102024
	CloudIntegrationNewResponseEnvelopeMessagesCode102025 CloudIntegrationNewResponseEnvelopeMessagesCode = 102025
	CloudIntegrationNewResponseEnvelopeMessagesCode102026 CloudIntegrationNewResponseEnvelopeMessagesCode = 102026
	CloudIntegrationNewResponseEnvelopeMessagesCode102027 CloudIntegrationNewResponseEnvelopeMessagesCode = 102027
	CloudIntegrationNewResponseEnvelopeMessagesCode102028 CloudIntegrationNewResponseEnvelopeMessagesCode = 102028
	CloudIntegrationNewResponseEnvelopeMessagesCode102029 CloudIntegrationNewResponseEnvelopeMessagesCode = 102029
	CloudIntegrationNewResponseEnvelopeMessagesCode102030 CloudIntegrationNewResponseEnvelopeMessagesCode = 102030
	CloudIntegrationNewResponseEnvelopeMessagesCode102031 CloudIntegrationNewResponseEnvelopeMessagesCode = 102031
	CloudIntegrationNewResponseEnvelopeMessagesCode102032 CloudIntegrationNewResponseEnvelopeMessagesCode = 102032
	CloudIntegrationNewResponseEnvelopeMessagesCode102033 CloudIntegrationNewResponseEnvelopeMessagesCode = 102033
	CloudIntegrationNewResponseEnvelopeMessagesCode102034 CloudIntegrationNewResponseEnvelopeMessagesCode = 102034
	CloudIntegrationNewResponseEnvelopeMessagesCode102035 CloudIntegrationNewResponseEnvelopeMessagesCode = 102035
	CloudIntegrationNewResponseEnvelopeMessagesCode102036 CloudIntegrationNewResponseEnvelopeMessagesCode = 102036
	CloudIntegrationNewResponseEnvelopeMessagesCode102037 CloudIntegrationNewResponseEnvelopeMessagesCode = 102037
	CloudIntegrationNewResponseEnvelopeMessagesCode102038 CloudIntegrationNewResponseEnvelopeMessagesCode = 102038
	CloudIntegrationNewResponseEnvelopeMessagesCode102039 CloudIntegrationNewResponseEnvelopeMessagesCode = 102039
	CloudIntegrationNewResponseEnvelopeMessagesCode102040 CloudIntegrationNewResponseEnvelopeMessagesCode = 102040
	CloudIntegrationNewResponseEnvelopeMessagesCode102041 CloudIntegrationNewResponseEnvelopeMessagesCode = 102041
	CloudIntegrationNewResponseEnvelopeMessagesCode102042 CloudIntegrationNewResponseEnvelopeMessagesCode = 102042
	CloudIntegrationNewResponseEnvelopeMessagesCode102043 CloudIntegrationNewResponseEnvelopeMessagesCode = 102043
	CloudIntegrationNewResponseEnvelopeMessagesCode102044 CloudIntegrationNewResponseEnvelopeMessagesCode = 102044
	CloudIntegrationNewResponseEnvelopeMessagesCode102045 CloudIntegrationNewResponseEnvelopeMessagesCode = 102045
	CloudIntegrationNewResponseEnvelopeMessagesCode102046 CloudIntegrationNewResponseEnvelopeMessagesCode = 102046
	CloudIntegrationNewResponseEnvelopeMessagesCode102047 CloudIntegrationNewResponseEnvelopeMessagesCode = 102047
	CloudIntegrationNewResponseEnvelopeMessagesCode102048 CloudIntegrationNewResponseEnvelopeMessagesCode = 102048
	CloudIntegrationNewResponseEnvelopeMessagesCode102049 CloudIntegrationNewResponseEnvelopeMessagesCode = 102049
	CloudIntegrationNewResponseEnvelopeMessagesCode102050 CloudIntegrationNewResponseEnvelopeMessagesCode = 102050
	CloudIntegrationNewResponseEnvelopeMessagesCode102051 CloudIntegrationNewResponseEnvelopeMessagesCode = 102051
	CloudIntegrationNewResponseEnvelopeMessagesCode102052 CloudIntegrationNewResponseEnvelopeMessagesCode = 102052
	CloudIntegrationNewResponseEnvelopeMessagesCode102053 CloudIntegrationNewResponseEnvelopeMessagesCode = 102053
	CloudIntegrationNewResponseEnvelopeMessagesCode102054 CloudIntegrationNewResponseEnvelopeMessagesCode = 102054
	CloudIntegrationNewResponseEnvelopeMessagesCode102055 CloudIntegrationNewResponseEnvelopeMessagesCode = 102055
	CloudIntegrationNewResponseEnvelopeMessagesCode102056 CloudIntegrationNewResponseEnvelopeMessagesCode = 102056
	CloudIntegrationNewResponseEnvelopeMessagesCode102057 CloudIntegrationNewResponseEnvelopeMessagesCode = 102057
	CloudIntegrationNewResponseEnvelopeMessagesCode102058 CloudIntegrationNewResponseEnvelopeMessagesCode = 102058
	CloudIntegrationNewResponseEnvelopeMessagesCode102059 CloudIntegrationNewResponseEnvelopeMessagesCode = 102059
	CloudIntegrationNewResponseEnvelopeMessagesCode102060 CloudIntegrationNewResponseEnvelopeMessagesCode = 102060
	CloudIntegrationNewResponseEnvelopeMessagesCode102061 CloudIntegrationNewResponseEnvelopeMessagesCode = 102061
	CloudIntegrationNewResponseEnvelopeMessagesCode102062 CloudIntegrationNewResponseEnvelopeMessagesCode = 102062
	CloudIntegrationNewResponseEnvelopeMessagesCode102063 CloudIntegrationNewResponseEnvelopeMessagesCode = 102063
	CloudIntegrationNewResponseEnvelopeMessagesCode102064 CloudIntegrationNewResponseEnvelopeMessagesCode = 102064
	CloudIntegrationNewResponseEnvelopeMessagesCode102065 CloudIntegrationNewResponseEnvelopeMessagesCode = 102065
	CloudIntegrationNewResponseEnvelopeMessagesCode102066 CloudIntegrationNewResponseEnvelopeMessagesCode = 102066
	CloudIntegrationNewResponseEnvelopeMessagesCode103001 CloudIntegrationNewResponseEnvelopeMessagesCode = 103001
	CloudIntegrationNewResponseEnvelopeMessagesCode103002 CloudIntegrationNewResponseEnvelopeMessagesCode = 103002
	CloudIntegrationNewResponseEnvelopeMessagesCode103003 CloudIntegrationNewResponseEnvelopeMessagesCode = 103003
	CloudIntegrationNewResponseEnvelopeMessagesCode103004 CloudIntegrationNewResponseEnvelopeMessagesCode = 103004
	CloudIntegrationNewResponseEnvelopeMessagesCode103005 CloudIntegrationNewResponseEnvelopeMessagesCode = 103005
	CloudIntegrationNewResponseEnvelopeMessagesCode103006 CloudIntegrationNewResponseEnvelopeMessagesCode = 103006
	CloudIntegrationNewResponseEnvelopeMessagesCode103007 CloudIntegrationNewResponseEnvelopeMessagesCode = 103007
	CloudIntegrationNewResponseEnvelopeMessagesCode103008 CloudIntegrationNewResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationNewResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationNewResponseEnvelopeMessagesCode1001, CloudIntegrationNewResponseEnvelopeMessagesCode1002, CloudIntegrationNewResponseEnvelopeMessagesCode1003, CloudIntegrationNewResponseEnvelopeMessagesCode1004, CloudIntegrationNewResponseEnvelopeMessagesCode1005, CloudIntegrationNewResponseEnvelopeMessagesCode1006, CloudIntegrationNewResponseEnvelopeMessagesCode1007, CloudIntegrationNewResponseEnvelopeMessagesCode1008, CloudIntegrationNewResponseEnvelopeMessagesCode1009, CloudIntegrationNewResponseEnvelopeMessagesCode1010, CloudIntegrationNewResponseEnvelopeMessagesCode1011, CloudIntegrationNewResponseEnvelopeMessagesCode1012, CloudIntegrationNewResponseEnvelopeMessagesCode1013, CloudIntegrationNewResponseEnvelopeMessagesCode1014, CloudIntegrationNewResponseEnvelopeMessagesCode1015, CloudIntegrationNewResponseEnvelopeMessagesCode1016, CloudIntegrationNewResponseEnvelopeMessagesCode1017, CloudIntegrationNewResponseEnvelopeMessagesCode2001, CloudIntegrationNewResponseEnvelopeMessagesCode2002, CloudIntegrationNewResponseEnvelopeMessagesCode2003, CloudIntegrationNewResponseEnvelopeMessagesCode2004, CloudIntegrationNewResponseEnvelopeMessagesCode2005, CloudIntegrationNewResponseEnvelopeMessagesCode2006, CloudIntegrationNewResponseEnvelopeMessagesCode2007, CloudIntegrationNewResponseEnvelopeMessagesCode2008, CloudIntegrationNewResponseEnvelopeMessagesCode2009, CloudIntegrationNewResponseEnvelopeMessagesCode2010, CloudIntegrationNewResponseEnvelopeMessagesCode2011, CloudIntegrationNewResponseEnvelopeMessagesCode2012, CloudIntegrationNewResponseEnvelopeMessagesCode2013, CloudIntegrationNewResponseEnvelopeMessagesCode2014, CloudIntegrationNewResponseEnvelopeMessagesCode2015, CloudIntegrationNewResponseEnvelopeMessagesCode2016, CloudIntegrationNewResponseEnvelopeMessagesCode2017, CloudIntegrationNewResponseEnvelopeMessagesCode2018, CloudIntegrationNewResponseEnvelopeMessagesCode2019, CloudIntegrationNewResponseEnvelopeMessagesCode2020, CloudIntegrationNewResponseEnvelopeMessagesCode2021, CloudIntegrationNewResponseEnvelopeMessagesCode2022, CloudIntegrationNewResponseEnvelopeMessagesCode3001, CloudIntegrationNewResponseEnvelopeMessagesCode3002, CloudIntegrationNewResponseEnvelopeMessagesCode3003, CloudIntegrationNewResponseEnvelopeMessagesCode3004, CloudIntegrationNewResponseEnvelopeMessagesCode3005, CloudIntegrationNewResponseEnvelopeMessagesCode3006, CloudIntegrationNewResponseEnvelopeMessagesCode3007, CloudIntegrationNewResponseEnvelopeMessagesCode4001, CloudIntegrationNewResponseEnvelopeMessagesCode4002, CloudIntegrationNewResponseEnvelopeMessagesCode4003, CloudIntegrationNewResponseEnvelopeMessagesCode4004, CloudIntegrationNewResponseEnvelopeMessagesCode4005, CloudIntegrationNewResponseEnvelopeMessagesCode4006, CloudIntegrationNewResponseEnvelopeMessagesCode4007, CloudIntegrationNewResponseEnvelopeMessagesCode4008, CloudIntegrationNewResponseEnvelopeMessagesCode4009, CloudIntegrationNewResponseEnvelopeMessagesCode4010, CloudIntegrationNewResponseEnvelopeMessagesCode4011, CloudIntegrationNewResponseEnvelopeMessagesCode4012, CloudIntegrationNewResponseEnvelopeMessagesCode4013, CloudIntegrationNewResponseEnvelopeMessagesCode4014, CloudIntegrationNewResponseEnvelopeMessagesCode4015, CloudIntegrationNewResponseEnvelopeMessagesCode4016, CloudIntegrationNewResponseEnvelopeMessagesCode4017, CloudIntegrationNewResponseEnvelopeMessagesCode4018, CloudIntegrationNewResponseEnvelopeMessagesCode4019, CloudIntegrationNewResponseEnvelopeMessagesCode4020, CloudIntegrationNewResponseEnvelopeMessagesCode4021, CloudIntegrationNewResponseEnvelopeMessagesCode4022, CloudIntegrationNewResponseEnvelopeMessagesCode4023, CloudIntegrationNewResponseEnvelopeMessagesCode5001, CloudIntegrationNewResponseEnvelopeMessagesCode5002, CloudIntegrationNewResponseEnvelopeMessagesCode5003, CloudIntegrationNewResponseEnvelopeMessagesCode5004, CloudIntegrationNewResponseEnvelopeMessagesCode102000, CloudIntegrationNewResponseEnvelopeMessagesCode102001, CloudIntegrationNewResponseEnvelopeMessagesCode102002, CloudIntegrationNewResponseEnvelopeMessagesCode102003, CloudIntegrationNewResponseEnvelopeMessagesCode102004, CloudIntegrationNewResponseEnvelopeMessagesCode102005, CloudIntegrationNewResponseEnvelopeMessagesCode102006, CloudIntegrationNewResponseEnvelopeMessagesCode102007, CloudIntegrationNewResponseEnvelopeMessagesCode102008, CloudIntegrationNewResponseEnvelopeMessagesCode102009, CloudIntegrationNewResponseEnvelopeMessagesCode102010, CloudIntegrationNewResponseEnvelopeMessagesCode102011, CloudIntegrationNewResponseEnvelopeMessagesCode102012, CloudIntegrationNewResponseEnvelopeMessagesCode102013, CloudIntegrationNewResponseEnvelopeMessagesCode102014, CloudIntegrationNewResponseEnvelopeMessagesCode102015, CloudIntegrationNewResponseEnvelopeMessagesCode102016, CloudIntegrationNewResponseEnvelopeMessagesCode102017, CloudIntegrationNewResponseEnvelopeMessagesCode102018, CloudIntegrationNewResponseEnvelopeMessagesCode102019, CloudIntegrationNewResponseEnvelopeMessagesCode102020, CloudIntegrationNewResponseEnvelopeMessagesCode102021, CloudIntegrationNewResponseEnvelopeMessagesCode102022, CloudIntegrationNewResponseEnvelopeMessagesCode102023, CloudIntegrationNewResponseEnvelopeMessagesCode102024, CloudIntegrationNewResponseEnvelopeMessagesCode102025, CloudIntegrationNewResponseEnvelopeMessagesCode102026, CloudIntegrationNewResponseEnvelopeMessagesCode102027, CloudIntegrationNewResponseEnvelopeMessagesCode102028, CloudIntegrationNewResponseEnvelopeMessagesCode102029, CloudIntegrationNewResponseEnvelopeMessagesCode102030, CloudIntegrationNewResponseEnvelopeMessagesCode102031, CloudIntegrationNewResponseEnvelopeMessagesCode102032, CloudIntegrationNewResponseEnvelopeMessagesCode102033, CloudIntegrationNewResponseEnvelopeMessagesCode102034, CloudIntegrationNewResponseEnvelopeMessagesCode102035, CloudIntegrationNewResponseEnvelopeMessagesCode102036, CloudIntegrationNewResponseEnvelopeMessagesCode102037, CloudIntegrationNewResponseEnvelopeMessagesCode102038, CloudIntegrationNewResponseEnvelopeMessagesCode102039, CloudIntegrationNewResponseEnvelopeMessagesCode102040, CloudIntegrationNewResponseEnvelopeMessagesCode102041, CloudIntegrationNewResponseEnvelopeMessagesCode102042, CloudIntegrationNewResponseEnvelopeMessagesCode102043, CloudIntegrationNewResponseEnvelopeMessagesCode102044, CloudIntegrationNewResponseEnvelopeMessagesCode102045, CloudIntegrationNewResponseEnvelopeMessagesCode102046, CloudIntegrationNewResponseEnvelopeMessagesCode102047, CloudIntegrationNewResponseEnvelopeMessagesCode102048, CloudIntegrationNewResponseEnvelopeMessagesCode102049, CloudIntegrationNewResponseEnvelopeMessagesCode102050, CloudIntegrationNewResponseEnvelopeMessagesCode102051, CloudIntegrationNewResponseEnvelopeMessagesCode102052, CloudIntegrationNewResponseEnvelopeMessagesCode102053, CloudIntegrationNewResponseEnvelopeMessagesCode102054, CloudIntegrationNewResponseEnvelopeMessagesCode102055, CloudIntegrationNewResponseEnvelopeMessagesCode102056, CloudIntegrationNewResponseEnvelopeMessagesCode102057, CloudIntegrationNewResponseEnvelopeMessagesCode102058, CloudIntegrationNewResponseEnvelopeMessagesCode102059, CloudIntegrationNewResponseEnvelopeMessagesCode102060, CloudIntegrationNewResponseEnvelopeMessagesCode102061, CloudIntegrationNewResponseEnvelopeMessagesCode102062, CloudIntegrationNewResponseEnvelopeMessagesCode102063, CloudIntegrationNewResponseEnvelopeMessagesCode102064, CloudIntegrationNewResponseEnvelopeMessagesCode102065, CloudIntegrationNewResponseEnvelopeMessagesCode102066, CloudIntegrationNewResponseEnvelopeMessagesCode103001, CloudIntegrationNewResponseEnvelopeMessagesCode103002, CloudIntegrationNewResponseEnvelopeMessagesCode103003, CloudIntegrationNewResponseEnvelopeMessagesCode103004, CloudIntegrationNewResponseEnvelopeMessagesCode103005, CloudIntegrationNewResponseEnvelopeMessagesCode103006, CloudIntegrationNewResponseEnvelopeMessagesCode103007, CloudIntegrationNewResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationNewResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                              `json:"l10n_key"`
	LoggableError string                                              `json:"loggable_error"`
	TemplateData  interface{}                                         `json:"template_data"`
	TraceID       string                                              `json:"trace_id"`
	JSON          cloudIntegrationNewResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeMessagesMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationNewResponseEnvelopeMessagesMeta]
type cloudIntegrationNewResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationNewResponseEnvelopeMessagesSource struct {
	Parameter           string                                                `json:"parameter"`
	ParameterValueIndex int64                                                 `json:"parameter_value_index"`
	Pointer             string                                                `json:"pointer"`
	JSON                cloudIntegrationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationNewResponseEnvelopeMessagesSource]
type cloudIntegrationNewResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateParams struct {
	AccountID              param.Field[string] `path:"account_id,required"`
	AwsArn                 param.Field[string] `json:"aws_arn"`
	AzureSubscriptionID    param.Field[string] `json:"azure_subscription_id"`
	AzureTenantID          param.Field[string] `json:"azure_tenant_id"`
	Description            param.Field[string] `json:"description"`
	FriendlyName           param.Field[string] `json:"friendly_name"`
	GcpProjectID           param.Field[string] `json:"gcp_project_id"`
	GcpServiceAccountEmail param.Field[string] `json:"gcp_service_account_email"`
}

func (r CloudIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudIntegrationUpdateResponseEnvelope struct {
	Errors   []CloudIntegrationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationUpdateResponse                   `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     cloudIntegrationUpdateResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudIntegrationUpdateResponseEnvelope]
type cloudIntegrationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeErrors struct {
	Code             CloudIntegrationUpdateResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Meta             CloudIntegrationUpdateResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CloudIntegrationUpdateResponseEnvelopeErrors]
type cloudIntegrationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1001   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1002   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1003   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1004   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1005   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1006   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1007   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1008   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1008
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1009   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1009
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1010   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1010
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1011   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1011
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1012   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1012
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1013   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1013
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1014   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1014
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1015   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1015
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1016   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1016
	CloudIntegrationUpdateResponseEnvelopeErrorsCode1017   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 1017
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2001   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2002   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2003   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2004   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2005   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2006   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2007   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2008   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2008
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2009   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2009
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2010   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2010
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2011   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2011
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2012   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2012
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2013   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2013
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2014   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2014
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2015   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2015
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2016   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2016
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2017   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2017
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2018   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2018
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2019   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2019
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2020   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2020
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2021   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2021
	CloudIntegrationUpdateResponseEnvelopeErrorsCode2022   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 2022
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3001   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3002   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3003   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3004   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3005   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3006   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode3007   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 3007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4001   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4002   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4003   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4004   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4005   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4006   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4007   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4008   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4008
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4009   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4009
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4010   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4010
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4011   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4011
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4012   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4012
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4013   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4013
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4014   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4014
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4015   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4015
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4016   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4016
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4017   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4017
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4018   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4018
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4019   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4019
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4020   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4020
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4021   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4021
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4022   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4022
	CloudIntegrationUpdateResponseEnvelopeErrorsCode4023   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 4023
	CloudIntegrationUpdateResponseEnvelopeErrorsCode5001   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 5001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode5002   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 5002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode5003   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 5003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode5004   CloudIntegrationUpdateResponseEnvelopeErrorsCode = 5004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102000 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102000
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102001 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102002 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102003 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102004 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102005 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102006 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102007 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102008 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102008
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102009 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102009
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102010 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102010
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102011 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102011
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102012 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102012
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102013 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102013
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102014 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102014
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102015 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102015
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102016 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102016
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102017 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102017
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102018 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102018
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102019 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102019
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102020 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102020
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102021 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102021
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102022 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102022
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102023 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102023
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102024 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102024
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102025 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102025
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102026 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102026
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102027 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102027
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102028 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102028
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102029 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102029
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102030 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102030
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102031 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102031
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102032 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102032
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102033 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102033
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102034 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102034
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102035 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102035
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102036 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102036
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102037 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102037
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102038 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102038
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102039 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102039
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102040 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102040
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102041 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102041
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102042 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102042
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102043 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102043
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102044 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102044
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102045 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102045
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102046 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102046
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102047 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102047
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102048 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102048
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102049 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102049
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102050 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102050
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102051 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102051
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102052 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102052
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102053 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102053
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102054 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102054
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102055 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102055
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102056 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102056
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102057 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102057
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102058 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102058
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102059 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102059
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102060 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102060
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102061 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102061
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102062 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102062
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102063 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102063
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102064 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102064
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102065 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102065
	CloudIntegrationUpdateResponseEnvelopeErrorsCode102066 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 102066
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103001 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103001
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103002 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103002
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103003 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103003
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103004 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103004
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103005 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103005
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103006 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103006
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103007 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103007
	CloudIntegrationUpdateResponseEnvelopeErrorsCode103008 CloudIntegrationUpdateResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationUpdateResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseEnvelopeErrorsCode1001, CloudIntegrationUpdateResponseEnvelopeErrorsCode1002, CloudIntegrationUpdateResponseEnvelopeErrorsCode1003, CloudIntegrationUpdateResponseEnvelopeErrorsCode1004, CloudIntegrationUpdateResponseEnvelopeErrorsCode1005, CloudIntegrationUpdateResponseEnvelopeErrorsCode1006, CloudIntegrationUpdateResponseEnvelopeErrorsCode1007, CloudIntegrationUpdateResponseEnvelopeErrorsCode1008, CloudIntegrationUpdateResponseEnvelopeErrorsCode1009, CloudIntegrationUpdateResponseEnvelopeErrorsCode1010, CloudIntegrationUpdateResponseEnvelopeErrorsCode1011, CloudIntegrationUpdateResponseEnvelopeErrorsCode1012, CloudIntegrationUpdateResponseEnvelopeErrorsCode1013, CloudIntegrationUpdateResponseEnvelopeErrorsCode1014, CloudIntegrationUpdateResponseEnvelopeErrorsCode1015, CloudIntegrationUpdateResponseEnvelopeErrorsCode1016, CloudIntegrationUpdateResponseEnvelopeErrorsCode1017, CloudIntegrationUpdateResponseEnvelopeErrorsCode2001, CloudIntegrationUpdateResponseEnvelopeErrorsCode2002, CloudIntegrationUpdateResponseEnvelopeErrorsCode2003, CloudIntegrationUpdateResponseEnvelopeErrorsCode2004, CloudIntegrationUpdateResponseEnvelopeErrorsCode2005, CloudIntegrationUpdateResponseEnvelopeErrorsCode2006, CloudIntegrationUpdateResponseEnvelopeErrorsCode2007, CloudIntegrationUpdateResponseEnvelopeErrorsCode2008, CloudIntegrationUpdateResponseEnvelopeErrorsCode2009, CloudIntegrationUpdateResponseEnvelopeErrorsCode2010, CloudIntegrationUpdateResponseEnvelopeErrorsCode2011, CloudIntegrationUpdateResponseEnvelopeErrorsCode2012, CloudIntegrationUpdateResponseEnvelopeErrorsCode2013, CloudIntegrationUpdateResponseEnvelopeErrorsCode2014, CloudIntegrationUpdateResponseEnvelopeErrorsCode2015, CloudIntegrationUpdateResponseEnvelopeErrorsCode2016, CloudIntegrationUpdateResponseEnvelopeErrorsCode2017, CloudIntegrationUpdateResponseEnvelopeErrorsCode2018, CloudIntegrationUpdateResponseEnvelopeErrorsCode2019, CloudIntegrationUpdateResponseEnvelopeErrorsCode2020, CloudIntegrationUpdateResponseEnvelopeErrorsCode2021, CloudIntegrationUpdateResponseEnvelopeErrorsCode2022, CloudIntegrationUpdateResponseEnvelopeErrorsCode3001, CloudIntegrationUpdateResponseEnvelopeErrorsCode3002, CloudIntegrationUpdateResponseEnvelopeErrorsCode3003, CloudIntegrationUpdateResponseEnvelopeErrorsCode3004, CloudIntegrationUpdateResponseEnvelopeErrorsCode3005, CloudIntegrationUpdateResponseEnvelopeErrorsCode3006, CloudIntegrationUpdateResponseEnvelopeErrorsCode3007, CloudIntegrationUpdateResponseEnvelopeErrorsCode4001, CloudIntegrationUpdateResponseEnvelopeErrorsCode4002, CloudIntegrationUpdateResponseEnvelopeErrorsCode4003, CloudIntegrationUpdateResponseEnvelopeErrorsCode4004, CloudIntegrationUpdateResponseEnvelopeErrorsCode4005, CloudIntegrationUpdateResponseEnvelopeErrorsCode4006, CloudIntegrationUpdateResponseEnvelopeErrorsCode4007, CloudIntegrationUpdateResponseEnvelopeErrorsCode4008, CloudIntegrationUpdateResponseEnvelopeErrorsCode4009, CloudIntegrationUpdateResponseEnvelopeErrorsCode4010, CloudIntegrationUpdateResponseEnvelopeErrorsCode4011, CloudIntegrationUpdateResponseEnvelopeErrorsCode4012, CloudIntegrationUpdateResponseEnvelopeErrorsCode4013, CloudIntegrationUpdateResponseEnvelopeErrorsCode4014, CloudIntegrationUpdateResponseEnvelopeErrorsCode4015, CloudIntegrationUpdateResponseEnvelopeErrorsCode4016, CloudIntegrationUpdateResponseEnvelopeErrorsCode4017, CloudIntegrationUpdateResponseEnvelopeErrorsCode4018, CloudIntegrationUpdateResponseEnvelopeErrorsCode4019, CloudIntegrationUpdateResponseEnvelopeErrorsCode4020, CloudIntegrationUpdateResponseEnvelopeErrorsCode4021, CloudIntegrationUpdateResponseEnvelopeErrorsCode4022, CloudIntegrationUpdateResponseEnvelopeErrorsCode4023, CloudIntegrationUpdateResponseEnvelopeErrorsCode5001, CloudIntegrationUpdateResponseEnvelopeErrorsCode5002, CloudIntegrationUpdateResponseEnvelopeErrorsCode5003, CloudIntegrationUpdateResponseEnvelopeErrorsCode5004, CloudIntegrationUpdateResponseEnvelopeErrorsCode102000, CloudIntegrationUpdateResponseEnvelopeErrorsCode102001, CloudIntegrationUpdateResponseEnvelopeErrorsCode102002, CloudIntegrationUpdateResponseEnvelopeErrorsCode102003, CloudIntegrationUpdateResponseEnvelopeErrorsCode102004, CloudIntegrationUpdateResponseEnvelopeErrorsCode102005, CloudIntegrationUpdateResponseEnvelopeErrorsCode102006, CloudIntegrationUpdateResponseEnvelopeErrorsCode102007, CloudIntegrationUpdateResponseEnvelopeErrorsCode102008, CloudIntegrationUpdateResponseEnvelopeErrorsCode102009, CloudIntegrationUpdateResponseEnvelopeErrorsCode102010, CloudIntegrationUpdateResponseEnvelopeErrorsCode102011, CloudIntegrationUpdateResponseEnvelopeErrorsCode102012, CloudIntegrationUpdateResponseEnvelopeErrorsCode102013, CloudIntegrationUpdateResponseEnvelopeErrorsCode102014, CloudIntegrationUpdateResponseEnvelopeErrorsCode102015, CloudIntegrationUpdateResponseEnvelopeErrorsCode102016, CloudIntegrationUpdateResponseEnvelopeErrorsCode102017, CloudIntegrationUpdateResponseEnvelopeErrorsCode102018, CloudIntegrationUpdateResponseEnvelopeErrorsCode102019, CloudIntegrationUpdateResponseEnvelopeErrorsCode102020, CloudIntegrationUpdateResponseEnvelopeErrorsCode102021, CloudIntegrationUpdateResponseEnvelopeErrorsCode102022, CloudIntegrationUpdateResponseEnvelopeErrorsCode102023, CloudIntegrationUpdateResponseEnvelopeErrorsCode102024, CloudIntegrationUpdateResponseEnvelopeErrorsCode102025, CloudIntegrationUpdateResponseEnvelopeErrorsCode102026, CloudIntegrationUpdateResponseEnvelopeErrorsCode102027, CloudIntegrationUpdateResponseEnvelopeErrorsCode102028, CloudIntegrationUpdateResponseEnvelopeErrorsCode102029, CloudIntegrationUpdateResponseEnvelopeErrorsCode102030, CloudIntegrationUpdateResponseEnvelopeErrorsCode102031, CloudIntegrationUpdateResponseEnvelopeErrorsCode102032, CloudIntegrationUpdateResponseEnvelopeErrorsCode102033, CloudIntegrationUpdateResponseEnvelopeErrorsCode102034, CloudIntegrationUpdateResponseEnvelopeErrorsCode102035, CloudIntegrationUpdateResponseEnvelopeErrorsCode102036, CloudIntegrationUpdateResponseEnvelopeErrorsCode102037, CloudIntegrationUpdateResponseEnvelopeErrorsCode102038, CloudIntegrationUpdateResponseEnvelopeErrorsCode102039, CloudIntegrationUpdateResponseEnvelopeErrorsCode102040, CloudIntegrationUpdateResponseEnvelopeErrorsCode102041, CloudIntegrationUpdateResponseEnvelopeErrorsCode102042, CloudIntegrationUpdateResponseEnvelopeErrorsCode102043, CloudIntegrationUpdateResponseEnvelopeErrorsCode102044, CloudIntegrationUpdateResponseEnvelopeErrorsCode102045, CloudIntegrationUpdateResponseEnvelopeErrorsCode102046, CloudIntegrationUpdateResponseEnvelopeErrorsCode102047, CloudIntegrationUpdateResponseEnvelopeErrorsCode102048, CloudIntegrationUpdateResponseEnvelopeErrorsCode102049, CloudIntegrationUpdateResponseEnvelopeErrorsCode102050, CloudIntegrationUpdateResponseEnvelopeErrorsCode102051, CloudIntegrationUpdateResponseEnvelopeErrorsCode102052, CloudIntegrationUpdateResponseEnvelopeErrorsCode102053, CloudIntegrationUpdateResponseEnvelopeErrorsCode102054, CloudIntegrationUpdateResponseEnvelopeErrorsCode102055, CloudIntegrationUpdateResponseEnvelopeErrorsCode102056, CloudIntegrationUpdateResponseEnvelopeErrorsCode102057, CloudIntegrationUpdateResponseEnvelopeErrorsCode102058, CloudIntegrationUpdateResponseEnvelopeErrorsCode102059, CloudIntegrationUpdateResponseEnvelopeErrorsCode102060, CloudIntegrationUpdateResponseEnvelopeErrorsCode102061, CloudIntegrationUpdateResponseEnvelopeErrorsCode102062, CloudIntegrationUpdateResponseEnvelopeErrorsCode102063, CloudIntegrationUpdateResponseEnvelopeErrorsCode102064, CloudIntegrationUpdateResponseEnvelopeErrorsCode102065, CloudIntegrationUpdateResponseEnvelopeErrorsCode102066, CloudIntegrationUpdateResponseEnvelopeErrorsCode103001, CloudIntegrationUpdateResponseEnvelopeErrorsCode103002, CloudIntegrationUpdateResponseEnvelopeErrorsCode103003, CloudIntegrationUpdateResponseEnvelopeErrorsCode103004, CloudIntegrationUpdateResponseEnvelopeErrorsCode103005, CloudIntegrationUpdateResponseEnvelopeErrorsCode103006, CloudIntegrationUpdateResponseEnvelopeErrorsCode103007, CloudIntegrationUpdateResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                               `json:"l10n_key"`
	LoggableError string                                               `json:"loggable_error"`
	TemplateData  interface{}                                          `json:"template_data"`
	TraceID       string                                               `json:"trace_id"`
	JSON          cloudIntegrationUpdateResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationUpdateResponseEnvelopeErrorsMeta]
type cloudIntegrationUpdateResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeErrorsSource struct {
	Parameter           string                                                 `json:"parameter"`
	ParameterValueIndex int64                                                  `json:"parameter_value_index"`
	Pointer             string                                                 `json:"pointer"`
	JSON                cloudIntegrationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CloudIntegrationUpdateResponseEnvelopeErrorsSource]
type cloudIntegrationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeMessages struct {
	Code             CloudIntegrationUpdateResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Meta             CloudIntegrationUpdateResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudIntegrationUpdateResponseEnvelopeMessages]
type cloudIntegrationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1001   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1002   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1003   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1004   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1005   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1006   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1007   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1008   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1008
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1009   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1009
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1010   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1010
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1011   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1011
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1012   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1012
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1013   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1013
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1014   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1014
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1015   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1015
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1016   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1016
	CloudIntegrationUpdateResponseEnvelopeMessagesCode1017   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 1017
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2001   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2002   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2003   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2004   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2005   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2006   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2007   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2008   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2008
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2009   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2009
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2010   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2010
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2011   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2011
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2012   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2012
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2013   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2013
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2014   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2014
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2015   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2015
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2016   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2016
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2017   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2017
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2018   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2018
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2019   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2019
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2020   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2020
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2021   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2021
	CloudIntegrationUpdateResponseEnvelopeMessagesCode2022   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 2022
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3001   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3002   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3003   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3004   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3005   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3006   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode3007   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 3007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4001   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4002   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4003   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4004   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4005   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4006   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4007   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4008   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4008
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4009   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4009
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4010   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4010
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4011   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4011
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4012   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4012
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4013   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4013
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4014   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4014
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4015   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4015
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4016   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4016
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4017   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4017
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4018   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4018
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4019   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4019
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4020   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4020
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4021   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4021
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4022   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4022
	CloudIntegrationUpdateResponseEnvelopeMessagesCode4023   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 4023
	CloudIntegrationUpdateResponseEnvelopeMessagesCode5001   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 5001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode5002   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 5002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode5003   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 5003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode5004   CloudIntegrationUpdateResponseEnvelopeMessagesCode = 5004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102000 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102000
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102001 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102002 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102003 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102004 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102005 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102006 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102007 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102008 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102008
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102009 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102009
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102010 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102010
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102011 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102011
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102012 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102012
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102013 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102013
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102014 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102014
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102015 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102015
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102016 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102016
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102017 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102017
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102018 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102018
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102019 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102019
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102020 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102020
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102021 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102021
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102022 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102022
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102023 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102023
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102024 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102024
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102025 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102025
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102026 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102026
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102027 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102027
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102028 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102028
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102029 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102029
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102030 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102030
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102031 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102031
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102032 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102032
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102033 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102033
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102034 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102034
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102035 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102035
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102036 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102036
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102037 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102037
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102038 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102038
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102039 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102039
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102040 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102040
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102041 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102041
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102042 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102042
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102043 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102043
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102044 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102044
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102045 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102045
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102046 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102046
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102047 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102047
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102048 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102048
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102049 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102049
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102050 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102050
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102051 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102051
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102052 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102052
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102053 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102053
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102054 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102054
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102055 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102055
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102056 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102056
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102057 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102057
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102058 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102058
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102059 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102059
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102060 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102060
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102061 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102061
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102062 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102062
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102063 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102063
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102064 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102064
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102065 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102065
	CloudIntegrationUpdateResponseEnvelopeMessagesCode102066 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 102066
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103001 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103001
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103002 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103002
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103003 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103003
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103004 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103004
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103005 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103005
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103006 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103006
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103007 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103007
	CloudIntegrationUpdateResponseEnvelopeMessagesCode103008 CloudIntegrationUpdateResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationUpdateResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationUpdateResponseEnvelopeMessagesCode1001, CloudIntegrationUpdateResponseEnvelopeMessagesCode1002, CloudIntegrationUpdateResponseEnvelopeMessagesCode1003, CloudIntegrationUpdateResponseEnvelopeMessagesCode1004, CloudIntegrationUpdateResponseEnvelopeMessagesCode1005, CloudIntegrationUpdateResponseEnvelopeMessagesCode1006, CloudIntegrationUpdateResponseEnvelopeMessagesCode1007, CloudIntegrationUpdateResponseEnvelopeMessagesCode1008, CloudIntegrationUpdateResponseEnvelopeMessagesCode1009, CloudIntegrationUpdateResponseEnvelopeMessagesCode1010, CloudIntegrationUpdateResponseEnvelopeMessagesCode1011, CloudIntegrationUpdateResponseEnvelopeMessagesCode1012, CloudIntegrationUpdateResponseEnvelopeMessagesCode1013, CloudIntegrationUpdateResponseEnvelopeMessagesCode1014, CloudIntegrationUpdateResponseEnvelopeMessagesCode1015, CloudIntegrationUpdateResponseEnvelopeMessagesCode1016, CloudIntegrationUpdateResponseEnvelopeMessagesCode1017, CloudIntegrationUpdateResponseEnvelopeMessagesCode2001, CloudIntegrationUpdateResponseEnvelopeMessagesCode2002, CloudIntegrationUpdateResponseEnvelopeMessagesCode2003, CloudIntegrationUpdateResponseEnvelopeMessagesCode2004, CloudIntegrationUpdateResponseEnvelopeMessagesCode2005, CloudIntegrationUpdateResponseEnvelopeMessagesCode2006, CloudIntegrationUpdateResponseEnvelopeMessagesCode2007, CloudIntegrationUpdateResponseEnvelopeMessagesCode2008, CloudIntegrationUpdateResponseEnvelopeMessagesCode2009, CloudIntegrationUpdateResponseEnvelopeMessagesCode2010, CloudIntegrationUpdateResponseEnvelopeMessagesCode2011, CloudIntegrationUpdateResponseEnvelopeMessagesCode2012, CloudIntegrationUpdateResponseEnvelopeMessagesCode2013, CloudIntegrationUpdateResponseEnvelopeMessagesCode2014, CloudIntegrationUpdateResponseEnvelopeMessagesCode2015, CloudIntegrationUpdateResponseEnvelopeMessagesCode2016, CloudIntegrationUpdateResponseEnvelopeMessagesCode2017, CloudIntegrationUpdateResponseEnvelopeMessagesCode2018, CloudIntegrationUpdateResponseEnvelopeMessagesCode2019, CloudIntegrationUpdateResponseEnvelopeMessagesCode2020, CloudIntegrationUpdateResponseEnvelopeMessagesCode2021, CloudIntegrationUpdateResponseEnvelopeMessagesCode2022, CloudIntegrationUpdateResponseEnvelopeMessagesCode3001, CloudIntegrationUpdateResponseEnvelopeMessagesCode3002, CloudIntegrationUpdateResponseEnvelopeMessagesCode3003, CloudIntegrationUpdateResponseEnvelopeMessagesCode3004, CloudIntegrationUpdateResponseEnvelopeMessagesCode3005, CloudIntegrationUpdateResponseEnvelopeMessagesCode3006, CloudIntegrationUpdateResponseEnvelopeMessagesCode3007, CloudIntegrationUpdateResponseEnvelopeMessagesCode4001, CloudIntegrationUpdateResponseEnvelopeMessagesCode4002, CloudIntegrationUpdateResponseEnvelopeMessagesCode4003, CloudIntegrationUpdateResponseEnvelopeMessagesCode4004, CloudIntegrationUpdateResponseEnvelopeMessagesCode4005, CloudIntegrationUpdateResponseEnvelopeMessagesCode4006, CloudIntegrationUpdateResponseEnvelopeMessagesCode4007, CloudIntegrationUpdateResponseEnvelopeMessagesCode4008, CloudIntegrationUpdateResponseEnvelopeMessagesCode4009, CloudIntegrationUpdateResponseEnvelopeMessagesCode4010, CloudIntegrationUpdateResponseEnvelopeMessagesCode4011, CloudIntegrationUpdateResponseEnvelopeMessagesCode4012, CloudIntegrationUpdateResponseEnvelopeMessagesCode4013, CloudIntegrationUpdateResponseEnvelopeMessagesCode4014, CloudIntegrationUpdateResponseEnvelopeMessagesCode4015, CloudIntegrationUpdateResponseEnvelopeMessagesCode4016, CloudIntegrationUpdateResponseEnvelopeMessagesCode4017, CloudIntegrationUpdateResponseEnvelopeMessagesCode4018, CloudIntegrationUpdateResponseEnvelopeMessagesCode4019, CloudIntegrationUpdateResponseEnvelopeMessagesCode4020, CloudIntegrationUpdateResponseEnvelopeMessagesCode4021, CloudIntegrationUpdateResponseEnvelopeMessagesCode4022, CloudIntegrationUpdateResponseEnvelopeMessagesCode4023, CloudIntegrationUpdateResponseEnvelopeMessagesCode5001, CloudIntegrationUpdateResponseEnvelopeMessagesCode5002, CloudIntegrationUpdateResponseEnvelopeMessagesCode5003, CloudIntegrationUpdateResponseEnvelopeMessagesCode5004, CloudIntegrationUpdateResponseEnvelopeMessagesCode102000, CloudIntegrationUpdateResponseEnvelopeMessagesCode102001, CloudIntegrationUpdateResponseEnvelopeMessagesCode102002, CloudIntegrationUpdateResponseEnvelopeMessagesCode102003, CloudIntegrationUpdateResponseEnvelopeMessagesCode102004, CloudIntegrationUpdateResponseEnvelopeMessagesCode102005, CloudIntegrationUpdateResponseEnvelopeMessagesCode102006, CloudIntegrationUpdateResponseEnvelopeMessagesCode102007, CloudIntegrationUpdateResponseEnvelopeMessagesCode102008, CloudIntegrationUpdateResponseEnvelopeMessagesCode102009, CloudIntegrationUpdateResponseEnvelopeMessagesCode102010, CloudIntegrationUpdateResponseEnvelopeMessagesCode102011, CloudIntegrationUpdateResponseEnvelopeMessagesCode102012, CloudIntegrationUpdateResponseEnvelopeMessagesCode102013, CloudIntegrationUpdateResponseEnvelopeMessagesCode102014, CloudIntegrationUpdateResponseEnvelopeMessagesCode102015, CloudIntegrationUpdateResponseEnvelopeMessagesCode102016, CloudIntegrationUpdateResponseEnvelopeMessagesCode102017, CloudIntegrationUpdateResponseEnvelopeMessagesCode102018, CloudIntegrationUpdateResponseEnvelopeMessagesCode102019, CloudIntegrationUpdateResponseEnvelopeMessagesCode102020, CloudIntegrationUpdateResponseEnvelopeMessagesCode102021, CloudIntegrationUpdateResponseEnvelopeMessagesCode102022, CloudIntegrationUpdateResponseEnvelopeMessagesCode102023, CloudIntegrationUpdateResponseEnvelopeMessagesCode102024, CloudIntegrationUpdateResponseEnvelopeMessagesCode102025, CloudIntegrationUpdateResponseEnvelopeMessagesCode102026, CloudIntegrationUpdateResponseEnvelopeMessagesCode102027, CloudIntegrationUpdateResponseEnvelopeMessagesCode102028, CloudIntegrationUpdateResponseEnvelopeMessagesCode102029, CloudIntegrationUpdateResponseEnvelopeMessagesCode102030, CloudIntegrationUpdateResponseEnvelopeMessagesCode102031, CloudIntegrationUpdateResponseEnvelopeMessagesCode102032, CloudIntegrationUpdateResponseEnvelopeMessagesCode102033, CloudIntegrationUpdateResponseEnvelopeMessagesCode102034, CloudIntegrationUpdateResponseEnvelopeMessagesCode102035, CloudIntegrationUpdateResponseEnvelopeMessagesCode102036, CloudIntegrationUpdateResponseEnvelopeMessagesCode102037, CloudIntegrationUpdateResponseEnvelopeMessagesCode102038, CloudIntegrationUpdateResponseEnvelopeMessagesCode102039, CloudIntegrationUpdateResponseEnvelopeMessagesCode102040, CloudIntegrationUpdateResponseEnvelopeMessagesCode102041, CloudIntegrationUpdateResponseEnvelopeMessagesCode102042, CloudIntegrationUpdateResponseEnvelopeMessagesCode102043, CloudIntegrationUpdateResponseEnvelopeMessagesCode102044, CloudIntegrationUpdateResponseEnvelopeMessagesCode102045, CloudIntegrationUpdateResponseEnvelopeMessagesCode102046, CloudIntegrationUpdateResponseEnvelopeMessagesCode102047, CloudIntegrationUpdateResponseEnvelopeMessagesCode102048, CloudIntegrationUpdateResponseEnvelopeMessagesCode102049, CloudIntegrationUpdateResponseEnvelopeMessagesCode102050, CloudIntegrationUpdateResponseEnvelopeMessagesCode102051, CloudIntegrationUpdateResponseEnvelopeMessagesCode102052, CloudIntegrationUpdateResponseEnvelopeMessagesCode102053, CloudIntegrationUpdateResponseEnvelopeMessagesCode102054, CloudIntegrationUpdateResponseEnvelopeMessagesCode102055, CloudIntegrationUpdateResponseEnvelopeMessagesCode102056, CloudIntegrationUpdateResponseEnvelopeMessagesCode102057, CloudIntegrationUpdateResponseEnvelopeMessagesCode102058, CloudIntegrationUpdateResponseEnvelopeMessagesCode102059, CloudIntegrationUpdateResponseEnvelopeMessagesCode102060, CloudIntegrationUpdateResponseEnvelopeMessagesCode102061, CloudIntegrationUpdateResponseEnvelopeMessagesCode102062, CloudIntegrationUpdateResponseEnvelopeMessagesCode102063, CloudIntegrationUpdateResponseEnvelopeMessagesCode102064, CloudIntegrationUpdateResponseEnvelopeMessagesCode102065, CloudIntegrationUpdateResponseEnvelopeMessagesCode102066, CloudIntegrationUpdateResponseEnvelopeMessagesCode103001, CloudIntegrationUpdateResponseEnvelopeMessagesCode103002, CloudIntegrationUpdateResponseEnvelopeMessagesCode103003, CloudIntegrationUpdateResponseEnvelopeMessagesCode103004, CloudIntegrationUpdateResponseEnvelopeMessagesCode103005, CloudIntegrationUpdateResponseEnvelopeMessagesCode103006, CloudIntegrationUpdateResponseEnvelopeMessagesCode103007, CloudIntegrationUpdateResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationUpdateResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                 `json:"l10n_key"`
	LoggableError string                                                 `json:"loggable_error"`
	TemplateData  interface{}                                            `json:"template_data"`
	TraceID       string                                                 `json:"trace_id"`
	JSON          cloudIntegrationUpdateResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct [CloudIntegrationUpdateResponseEnvelopeMessagesMeta]
type cloudIntegrationUpdateResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationUpdateResponseEnvelopeMessagesSource struct {
	Parameter           string                                                   `json:"parameter"`
	ParameterValueIndex int64                                                    `json:"parameter_value_index"`
	Pointer             string                                                   `json:"pointer"`
	JSON                cloudIntegrationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CloudIntegrationUpdateResponseEnvelopeMessagesSource]
type cloudIntegrationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationListParams struct {
	AccountID  param.Field[string] `path:"account_id,required"`
	Cloudflare param.Field[bool]   `query:"cloudflare"`
	Desc       param.Field[bool]   `query:"desc"`
	// One of ["updated_at", "id", "cloud_type", "name"].
	OrderBy param.Field[string] `query:"order_by"`
	Status  param.Field[bool]   `query:"status"`
}

// URLQuery serializes [CloudIntegrationListParams]'s query parameters as
// `url.Values`.
func (r CloudIntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CloudIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CloudIntegrationDeleteResponseEnvelope struct {
	Errors   []CloudIntegrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationDeleteResponse                   `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     cloudIntegrationDeleteResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudIntegrationDeleteResponseEnvelope]
type cloudIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeErrors struct {
	Code             CloudIntegrationDeleteResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Meta             CloudIntegrationDeleteResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CloudIntegrationDeleteResponseEnvelopeErrors]
type cloudIntegrationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1001   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1002   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1003   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1004   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1005   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1006   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1007   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1008   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1008
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1009   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1009
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1010   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1010
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1011   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1011
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1012   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1012
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1013   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1013
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1014   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1014
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1015   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1015
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1016   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1016
	CloudIntegrationDeleteResponseEnvelopeErrorsCode1017   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 1017
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2001   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2002   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2003   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2004   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2005   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2006   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2007   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2008   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2008
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2009   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2009
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2010   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2010
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2011   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2011
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2012   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2012
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2013   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2013
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2014   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2014
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2015   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2015
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2016   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2016
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2017   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2017
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2018   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2018
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2019   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2019
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2020   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2020
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2021   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2021
	CloudIntegrationDeleteResponseEnvelopeErrorsCode2022   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 2022
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3001   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3002   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3003   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3004   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3005   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3006   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode3007   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 3007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4001   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4002   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4003   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4004   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4005   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4006   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4007   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4008   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4008
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4009   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4009
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4010   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4010
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4011   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4011
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4012   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4012
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4013   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4013
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4014   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4014
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4015   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4015
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4016   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4016
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4017   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4017
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4018   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4018
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4019   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4019
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4020   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4020
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4021   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4021
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4022   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4022
	CloudIntegrationDeleteResponseEnvelopeErrorsCode4023   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 4023
	CloudIntegrationDeleteResponseEnvelopeErrorsCode5001   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 5001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode5002   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 5002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode5003   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 5003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode5004   CloudIntegrationDeleteResponseEnvelopeErrorsCode = 5004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102000 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102000
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102001 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102002 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102003 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102004 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102005 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102006 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102007 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102008 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102008
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102009 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102009
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102010 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102010
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102011 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102011
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102012 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102012
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102013 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102013
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102014 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102014
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102015 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102015
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102016 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102016
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102017 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102017
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102018 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102018
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102019 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102019
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102020 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102020
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102021 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102021
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102022 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102022
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102023 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102023
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102024 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102024
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102025 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102025
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102026 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102026
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102027 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102027
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102028 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102028
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102029 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102029
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102030 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102030
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102031 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102031
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102032 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102032
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102033 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102033
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102034 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102034
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102035 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102035
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102036 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102036
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102037 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102037
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102038 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102038
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102039 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102039
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102040 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102040
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102041 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102041
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102042 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102042
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102043 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102043
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102044 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102044
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102045 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102045
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102046 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102046
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102047 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102047
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102048 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102048
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102049 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102049
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102050 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102050
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102051 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102051
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102052 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102052
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102053 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102053
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102054 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102054
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102055 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102055
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102056 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102056
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102057 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102057
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102058 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102058
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102059 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102059
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102060 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102060
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102061 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102061
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102062 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102062
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102063 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102063
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102064 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102064
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102065 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102065
	CloudIntegrationDeleteResponseEnvelopeErrorsCode102066 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 102066
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103001 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103001
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103002 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103002
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103003 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103003
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103004 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103004
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103005 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103005
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103006 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103006
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103007 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103007
	CloudIntegrationDeleteResponseEnvelopeErrorsCode103008 CloudIntegrationDeleteResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationDeleteResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDeleteResponseEnvelopeErrorsCode1001, CloudIntegrationDeleteResponseEnvelopeErrorsCode1002, CloudIntegrationDeleteResponseEnvelopeErrorsCode1003, CloudIntegrationDeleteResponseEnvelopeErrorsCode1004, CloudIntegrationDeleteResponseEnvelopeErrorsCode1005, CloudIntegrationDeleteResponseEnvelopeErrorsCode1006, CloudIntegrationDeleteResponseEnvelopeErrorsCode1007, CloudIntegrationDeleteResponseEnvelopeErrorsCode1008, CloudIntegrationDeleteResponseEnvelopeErrorsCode1009, CloudIntegrationDeleteResponseEnvelopeErrorsCode1010, CloudIntegrationDeleteResponseEnvelopeErrorsCode1011, CloudIntegrationDeleteResponseEnvelopeErrorsCode1012, CloudIntegrationDeleteResponseEnvelopeErrorsCode1013, CloudIntegrationDeleteResponseEnvelopeErrorsCode1014, CloudIntegrationDeleteResponseEnvelopeErrorsCode1015, CloudIntegrationDeleteResponseEnvelopeErrorsCode1016, CloudIntegrationDeleteResponseEnvelopeErrorsCode1017, CloudIntegrationDeleteResponseEnvelopeErrorsCode2001, CloudIntegrationDeleteResponseEnvelopeErrorsCode2002, CloudIntegrationDeleteResponseEnvelopeErrorsCode2003, CloudIntegrationDeleteResponseEnvelopeErrorsCode2004, CloudIntegrationDeleteResponseEnvelopeErrorsCode2005, CloudIntegrationDeleteResponseEnvelopeErrorsCode2006, CloudIntegrationDeleteResponseEnvelopeErrorsCode2007, CloudIntegrationDeleteResponseEnvelopeErrorsCode2008, CloudIntegrationDeleteResponseEnvelopeErrorsCode2009, CloudIntegrationDeleteResponseEnvelopeErrorsCode2010, CloudIntegrationDeleteResponseEnvelopeErrorsCode2011, CloudIntegrationDeleteResponseEnvelopeErrorsCode2012, CloudIntegrationDeleteResponseEnvelopeErrorsCode2013, CloudIntegrationDeleteResponseEnvelopeErrorsCode2014, CloudIntegrationDeleteResponseEnvelopeErrorsCode2015, CloudIntegrationDeleteResponseEnvelopeErrorsCode2016, CloudIntegrationDeleteResponseEnvelopeErrorsCode2017, CloudIntegrationDeleteResponseEnvelopeErrorsCode2018, CloudIntegrationDeleteResponseEnvelopeErrorsCode2019, CloudIntegrationDeleteResponseEnvelopeErrorsCode2020, CloudIntegrationDeleteResponseEnvelopeErrorsCode2021, CloudIntegrationDeleteResponseEnvelopeErrorsCode2022, CloudIntegrationDeleteResponseEnvelopeErrorsCode3001, CloudIntegrationDeleteResponseEnvelopeErrorsCode3002, CloudIntegrationDeleteResponseEnvelopeErrorsCode3003, CloudIntegrationDeleteResponseEnvelopeErrorsCode3004, CloudIntegrationDeleteResponseEnvelopeErrorsCode3005, CloudIntegrationDeleteResponseEnvelopeErrorsCode3006, CloudIntegrationDeleteResponseEnvelopeErrorsCode3007, CloudIntegrationDeleteResponseEnvelopeErrorsCode4001, CloudIntegrationDeleteResponseEnvelopeErrorsCode4002, CloudIntegrationDeleteResponseEnvelopeErrorsCode4003, CloudIntegrationDeleteResponseEnvelopeErrorsCode4004, CloudIntegrationDeleteResponseEnvelopeErrorsCode4005, CloudIntegrationDeleteResponseEnvelopeErrorsCode4006, CloudIntegrationDeleteResponseEnvelopeErrorsCode4007, CloudIntegrationDeleteResponseEnvelopeErrorsCode4008, CloudIntegrationDeleteResponseEnvelopeErrorsCode4009, CloudIntegrationDeleteResponseEnvelopeErrorsCode4010, CloudIntegrationDeleteResponseEnvelopeErrorsCode4011, CloudIntegrationDeleteResponseEnvelopeErrorsCode4012, CloudIntegrationDeleteResponseEnvelopeErrorsCode4013, CloudIntegrationDeleteResponseEnvelopeErrorsCode4014, CloudIntegrationDeleteResponseEnvelopeErrorsCode4015, CloudIntegrationDeleteResponseEnvelopeErrorsCode4016, CloudIntegrationDeleteResponseEnvelopeErrorsCode4017, CloudIntegrationDeleteResponseEnvelopeErrorsCode4018, CloudIntegrationDeleteResponseEnvelopeErrorsCode4019, CloudIntegrationDeleteResponseEnvelopeErrorsCode4020, CloudIntegrationDeleteResponseEnvelopeErrorsCode4021, CloudIntegrationDeleteResponseEnvelopeErrorsCode4022, CloudIntegrationDeleteResponseEnvelopeErrorsCode4023, CloudIntegrationDeleteResponseEnvelopeErrorsCode5001, CloudIntegrationDeleteResponseEnvelopeErrorsCode5002, CloudIntegrationDeleteResponseEnvelopeErrorsCode5003, CloudIntegrationDeleteResponseEnvelopeErrorsCode5004, CloudIntegrationDeleteResponseEnvelopeErrorsCode102000, CloudIntegrationDeleteResponseEnvelopeErrorsCode102001, CloudIntegrationDeleteResponseEnvelopeErrorsCode102002, CloudIntegrationDeleteResponseEnvelopeErrorsCode102003, CloudIntegrationDeleteResponseEnvelopeErrorsCode102004, CloudIntegrationDeleteResponseEnvelopeErrorsCode102005, CloudIntegrationDeleteResponseEnvelopeErrorsCode102006, CloudIntegrationDeleteResponseEnvelopeErrorsCode102007, CloudIntegrationDeleteResponseEnvelopeErrorsCode102008, CloudIntegrationDeleteResponseEnvelopeErrorsCode102009, CloudIntegrationDeleteResponseEnvelopeErrorsCode102010, CloudIntegrationDeleteResponseEnvelopeErrorsCode102011, CloudIntegrationDeleteResponseEnvelopeErrorsCode102012, CloudIntegrationDeleteResponseEnvelopeErrorsCode102013, CloudIntegrationDeleteResponseEnvelopeErrorsCode102014, CloudIntegrationDeleteResponseEnvelopeErrorsCode102015, CloudIntegrationDeleteResponseEnvelopeErrorsCode102016, CloudIntegrationDeleteResponseEnvelopeErrorsCode102017, CloudIntegrationDeleteResponseEnvelopeErrorsCode102018, CloudIntegrationDeleteResponseEnvelopeErrorsCode102019, CloudIntegrationDeleteResponseEnvelopeErrorsCode102020, CloudIntegrationDeleteResponseEnvelopeErrorsCode102021, CloudIntegrationDeleteResponseEnvelopeErrorsCode102022, CloudIntegrationDeleteResponseEnvelopeErrorsCode102023, CloudIntegrationDeleteResponseEnvelopeErrorsCode102024, CloudIntegrationDeleteResponseEnvelopeErrorsCode102025, CloudIntegrationDeleteResponseEnvelopeErrorsCode102026, CloudIntegrationDeleteResponseEnvelopeErrorsCode102027, CloudIntegrationDeleteResponseEnvelopeErrorsCode102028, CloudIntegrationDeleteResponseEnvelopeErrorsCode102029, CloudIntegrationDeleteResponseEnvelopeErrorsCode102030, CloudIntegrationDeleteResponseEnvelopeErrorsCode102031, CloudIntegrationDeleteResponseEnvelopeErrorsCode102032, CloudIntegrationDeleteResponseEnvelopeErrorsCode102033, CloudIntegrationDeleteResponseEnvelopeErrorsCode102034, CloudIntegrationDeleteResponseEnvelopeErrorsCode102035, CloudIntegrationDeleteResponseEnvelopeErrorsCode102036, CloudIntegrationDeleteResponseEnvelopeErrorsCode102037, CloudIntegrationDeleteResponseEnvelopeErrorsCode102038, CloudIntegrationDeleteResponseEnvelopeErrorsCode102039, CloudIntegrationDeleteResponseEnvelopeErrorsCode102040, CloudIntegrationDeleteResponseEnvelopeErrorsCode102041, CloudIntegrationDeleteResponseEnvelopeErrorsCode102042, CloudIntegrationDeleteResponseEnvelopeErrorsCode102043, CloudIntegrationDeleteResponseEnvelopeErrorsCode102044, CloudIntegrationDeleteResponseEnvelopeErrorsCode102045, CloudIntegrationDeleteResponseEnvelopeErrorsCode102046, CloudIntegrationDeleteResponseEnvelopeErrorsCode102047, CloudIntegrationDeleteResponseEnvelopeErrorsCode102048, CloudIntegrationDeleteResponseEnvelopeErrorsCode102049, CloudIntegrationDeleteResponseEnvelopeErrorsCode102050, CloudIntegrationDeleteResponseEnvelopeErrorsCode102051, CloudIntegrationDeleteResponseEnvelopeErrorsCode102052, CloudIntegrationDeleteResponseEnvelopeErrorsCode102053, CloudIntegrationDeleteResponseEnvelopeErrorsCode102054, CloudIntegrationDeleteResponseEnvelopeErrorsCode102055, CloudIntegrationDeleteResponseEnvelopeErrorsCode102056, CloudIntegrationDeleteResponseEnvelopeErrorsCode102057, CloudIntegrationDeleteResponseEnvelopeErrorsCode102058, CloudIntegrationDeleteResponseEnvelopeErrorsCode102059, CloudIntegrationDeleteResponseEnvelopeErrorsCode102060, CloudIntegrationDeleteResponseEnvelopeErrorsCode102061, CloudIntegrationDeleteResponseEnvelopeErrorsCode102062, CloudIntegrationDeleteResponseEnvelopeErrorsCode102063, CloudIntegrationDeleteResponseEnvelopeErrorsCode102064, CloudIntegrationDeleteResponseEnvelopeErrorsCode102065, CloudIntegrationDeleteResponseEnvelopeErrorsCode102066, CloudIntegrationDeleteResponseEnvelopeErrorsCode103001, CloudIntegrationDeleteResponseEnvelopeErrorsCode103002, CloudIntegrationDeleteResponseEnvelopeErrorsCode103003, CloudIntegrationDeleteResponseEnvelopeErrorsCode103004, CloudIntegrationDeleteResponseEnvelopeErrorsCode103005, CloudIntegrationDeleteResponseEnvelopeErrorsCode103006, CloudIntegrationDeleteResponseEnvelopeErrorsCode103007, CloudIntegrationDeleteResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationDeleteResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                               `json:"l10n_key"`
	LoggableError string                                               `json:"loggable_error"`
	TemplateData  interface{}                                          `json:"template_data"`
	TraceID       string                                               `json:"trace_id"`
	JSON          cloudIntegrationDeleteResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationDeleteResponseEnvelopeErrorsMeta]
type cloudIntegrationDeleteResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeErrorsSource struct {
	Parameter           string                                                 `json:"parameter"`
	ParameterValueIndex int64                                                  `json:"parameter_value_index"`
	Pointer             string                                                 `json:"pointer"`
	JSON                cloudIntegrationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CloudIntegrationDeleteResponseEnvelopeErrorsSource]
type cloudIntegrationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeMessages struct {
	Code             CloudIntegrationDeleteResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Meta             CloudIntegrationDeleteResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CloudIntegrationDeleteResponseEnvelopeMessages]
type cloudIntegrationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1001   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1002   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1003   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1004   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1005   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1006   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1007   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1008   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1008
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1009   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1009
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1010   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1010
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1011   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1011
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1012   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1012
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1013   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1013
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1014   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1014
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1015   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1015
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1016   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1016
	CloudIntegrationDeleteResponseEnvelopeMessagesCode1017   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 1017
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2001   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2002   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2003   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2004   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2005   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2006   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2007   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2008   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2008
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2009   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2009
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2010   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2010
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2011   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2011
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2012   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2012
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2013   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2013
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2014   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2014
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2015   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2015
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2016   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2016
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2017   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2017
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2018   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2018
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2019   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2019
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2020   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2020
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2021   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2021
	CloudIntegrationDeleteResponseEnvelopeMessagesCode2022   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 2022
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3001   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3002   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3003   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3004   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3005   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3006   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode3007   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 3007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4001   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4002   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4003   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4004   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4005   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4006   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4007   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4008   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4008
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4009   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4009
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4010   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4010
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4011   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4011
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4012   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4012
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4013   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4013
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4014   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4014
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4015   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4015
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4016   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4016
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4017   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4017
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4018   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4018
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4019   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4019
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4020   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4020
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4021   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4021
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4022   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4022
	CloudIntegrationDeleteResponseEnvelopeMessagesCode4023   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 4023
	CloudIntegrationDeleteResponseEnvelopeMessagesCode5001   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 5001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode5002   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 5002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode5003   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 5003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode5004   CloudIntegrationDeleteResponseEnvelopeMessagesCode = 5004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102000 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102000
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102001 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102002 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102003 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102004 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102005 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102006 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102007 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102008 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102008
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102009 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102009
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102010 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102010
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102011 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102011
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102012 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102012
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102013 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102013
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102014 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102014
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102015 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102015
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102016 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102016
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102017 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102017
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102018 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102018
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102019 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102019
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102020 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102020
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102021 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102021
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102022 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102022
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102023 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102023
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102024 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102024
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102025 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102025
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102026 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102026
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102027 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102027
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102028 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102028
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102029 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102029
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102030 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102030
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102031 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102031
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102032 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102032
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102033 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102033
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102034 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102034
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102035 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102035
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102036 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102036
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102037 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102037
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102038 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102038
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102039 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102039
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102040 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102040
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102041 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102041
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102042 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102042
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102043 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102043
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102044 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102044
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102045 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102045
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102046 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102046
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102047 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102047
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102048 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102048
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102049 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102049
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102050 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102050
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102051 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102051
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102052 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102052
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102053 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102053
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102054 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102054
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102055 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102055
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102056 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102056
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102057 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102057
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102058 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102058
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102059 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102059
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102060 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102060
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102061 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102061
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102062 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102062
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102063 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102063
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102064 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102064
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102065 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102065
	CloudIntegrationDeleteResponseEnvelopeMessagesCode102066 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 102066
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103001 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103001
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103002 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103002
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103003 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103003
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103004 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103004
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103005 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103005
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103006 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103006
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103007 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103007
	CloudIntegrationDeleteResponseEnvelopeMessagesCode103008 CloudIntegrationDeleteResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationDeleteResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationDeleteResponseEnvelopeMessagesCode1001, CloudIntegrationDeleteResponseEnvelopeMessagesCode1002, CloudIntegrationDeleteResponseEnvelopeMessagesCode1003, CloudIntegrationDeleteResponseEnvelopeMessagesCode1004, CloudIntegrationDeleteResponseEnvelopeMessagesCode1005, CloudIntegrationDeleteResponseEnvelopeMessagesCode1006, CloudIntegrationDeleteResponseEnvelopeMessagesCode1007, CloudIntegrationDeleteResponseEnvelopeMessagesCode1008, CloudIntegrationDeleteResponseEnvelopeMessagesCode1009, CloudIntegrationDeleteResponseEnvelopeMessagesCode1010, CloudIntegrationDeleteResponseEnvelopeMessagesCode1011, CloudIntegrationDeleteResponseEnvelopeMessagesCode1012, CloudIntegrationDeleteResponseEnvelopeMessagesCode1013, CloudIntegrationDeleteResponseEnvelopeMessagesCode1014, CloudIntegrationDeleteResponseEnvelopeMessagesCode1015, CloudIntegrationDeleteResponseEnvelopeMessagesCode1016, CloudIntegrationDeleteResponseEnvelopeMessagesCode1017, CloudIntegrationDeleteResponseEnvelopeMessagesCode2001, CloudIntegrationDeleteResponseEnvelopeMessagesCode2002, CloudIntegrationDeleteResponseEnvelopeMessagesCode2003, CloudIntegrationDeleteResponseEnvelopeMessagesCode2004, CloudIntegrationDeleteResponseEnvelopeMessagesCode2005, CloudIntegrationDeleteResponseEnvelopeMessagesCode2006, CloudIntegrationDeleteResponseEnvelopeMessagesCode2007, CloudIntegrationDeleteResponseEnvelopeMessagesCode2008, CloudIntegrationDeleteResponseEnvelopeMessagesCode2009, CloudIntegrationDeleteResponseEnvelopeMessagesCode2010, CloudIntegrationDeleteResponseEnvelopeMessagesCode2011, CloudIntegrationDeleteResponseEnvelopeMessagesCode2012, CloudIntegrationDeleteResponseEnvelopeMessagesCode2013, CloudIntegrationDeleteResponseEnvelopeMessagesCode2014, CloudIntegrationDeleteResponseEnvelopeMessagesCode2015, CloudIntegrationDeleteResponseEnvelopeMessagesCode2016, CloudIntegrationDeleteResponseEnvelopeMessagesCode2017, CloudIntegrationDeleteResponseEnvelopeMessagesCode2018, CloudIntegrationDeleteResponseEnvelopeMessagesCode2019, CloudIntegrationDeleteResponseEnvelopeMessagesCode2020, CloudIntegrationDeleteResponseEnvelopeMessagesCode2021, CloudIntegrationDeleteResponseEnvelopeMessagesCode2022, CloudIntegrationDeleteResponseEnvelopeMessagesCode3001, CloudIntegrationDeleteResponseEnvelopeMessagesCode3002, CloudIntegrationDeleteResponseEnvelopeMessagesCode3003, CloudIntegrationDeleteResponseEnvelopeMessagesCode3004, CloudIntegrationDeleteResponseEnvelopeMessagesCode3005, CloudIntegrationDeleteResponseEnvelopeMessagesCode3006, CloudIntegrationDeleteResponseEnvelopeMessagesCode3007, CloudIntegrationDeleteResponseEnvelopeMessagesCode4001, CloudIntegrationDeleteResponseEnvelopeMessagesCode4002, CloudIntegrationDeleteResponseEnvelopeMessagesCode4003, CloudIntegrationDeleteResponseEnvelopeMessagesCode4004, CloudIntegrationDeleteResponseEnvelopeMessagesCode4005, CloudIntegrationDeleteResponseEnvelopeMessagesCode4006, CloudIntegrationDeleteResponseEnvelopeMessagesCode4007, CloudIntegrationDeleteResponseEnvelopeMessagesCode4008, CloudIntegrationDeleteResponseEnvelopeMessagesCode4009, CloudIntegrationDeleteResponseEnvelopeMessagesCode4010, CloudIntegrationDeleteResponseEnvelopeMessagesCode4011, CloudIntegrationDeleteResponseEnvelopeMessagesCode4012, CloudIntegrationDeleteResponseEnvelopeMessagesCode4013, CloudIntegrationDeleteResponseEnvelopeMessagesCode4014, CloudIntegrationDeleteResponseEnvelopeMessagesCode4015, CloudIntegrationDeleteResponseEnvelopeMessagesCode4016, CloudIntegrationDeleteResponseEnvelopeMessagesCode4017, CloudIntegrationDeleteResponseEnvelopeMessagesCode4018, CloudIntegrationDeleteResponseEnvelopeMessagesCode4019, CloudIntegrationDeleteResponseEnvelopeMessagesCode4020, CloudIntegrationDeleteResponseEnvelopeMessagesCode4021, CloudIntegrationDeleteResponseEnvelopeMessagesCode4022, CloudIntegrationDeleteResponseEnvelopeMessagesCode4023, CloudIntegrationDeleteResponseEnvelopeMessagesCode5001, CloudIntegrationDeleteResponseEnvelopeMessagesCode5002, CloudIntegrationDeleteResponseEnvelopeMessagesCode5003, CloudIntegrationDeleteResponseEnvelopeMessagesCode5004, CloudIntegrationDeleteResponseEnvelopeMessagesCode102000, CloudIntegrationDeleteResponseEnvelopeMessagesCode102001, CloudIntegrationDeleteResponseEnvelopeMessagesCode102002, CloudIntegrationDeleteResponseEnvelopeMessagesCode102003, CloudIntegrationDeleteResponseEnvelopeMessagesCode102004, CloudIntegrationDeleteResponseEnvelopeMessagesCode102005, CloudIntegrationDeleteResponseEnvelopeMessagesCode102006, CloudIntegrationDeleteResponseEnvelopeMessagesCode102007, CloudIntegrationDeleteResponseEnvelopeMessagesCode102008, CloudIntegrationDeleteResponseEnvelopeMessagesCode102009, CloudIntegrationDeleteResponseEnvelopeMessagesCode102010, CloudIntegrationDeleteResponseEnvelopeMessagesCode102011, CloudIntegrationDeleteResponseEnvelopeMessagesCode102012, CloudIntegrationDeleteResponseEnvelopeMessagesCode102013, CloudIntegrationDeleteResponseEnvelopeMessagesCode102014, CloudIntegrationDeleteResponseEnvelopeMessagesCode102015, CloudIntegrationDeleteResponseEnvelopeMessagesCode102016, CloudIntegrationDeleteResponseEnvelopeMessagesCode102017, CloudIntegrationDeleteResponseEnvelopeMessagesCode102018, CloudIntegrationDeleteResponseEnvelopeMessagesCode102019, CloudIntegrationDeleteResponseEnvelopeMessagesCode102020, CloudIntegrationDeleteResponseEnvelopeMessagesCode102021, CloudIntegrationDeleteResponseEnvelopeMessagesCode102022, CloudIntegrationDeleteResponseEnvelopeMessagesCode102023, CloudIntegrationDeleteResponseEnvelopeMessagesCode102024, CloudIntegrationDeleteResponseEnvelopeMessagesCode102025, CloudIntegrationDeleteResponseEnvelopeMessagesCode102026, CloudIntegrationDeleteResponseEnvelopeMessagesCode102027, CloudIntegrationDeleteResponseEnvelopeMessagesCode102028, CloudIntegrationDeleteResponseEnvelopeMessagesCode102029, CloudIntegrationDeleteResponseEnvelopeMessagesCode102030, CloudIntegrationDeleteResponseEnvelopeMessagesCode102031, CloudIntegrationDeleteResponseEnvelopeMessagesCode102032, CloudIntegrationDeleteResponseEnvelopeMessagesCode102033, CloudIntegrationDeleteResponseEnvelopeMessagesCode102034, CloudIntegrationDeleteResponseEnvelopeMessagesCode102035, CloudIntegrationDeleteResponseEnvelopeMessagesCode102036, CloudIntegrationDeleteResponseEnvelopeMessagesCode102037, CloudIntegrationDeleteResponseEnvelopeMessagesCode102038, CloudIntegrationDeleteResponseEnvelopeMessagesCode102039, CloudIntegrationDeleteResponseEnvelopeMessagesCode102040, CloudIntegrationDeleteResponseEnvelopeMessagesCode102041, CloudIntegrationDeleteResponseEnvelopeMessagesCode102042, CloudIntegrationDeleteResponseEnvelopeMessagesCode102043, CloudIntegrationDeleteResponseEnvelopeMessagesCode102044, CloudIntegrationDeleteResponseEnvelopeMessagesCode102045, CloudIntegrationDeleteResponseEnvelopeMessagesCode102046, CloudIntegrationDeleteResponseEnvelopeMessagesCode102047, CloudIntegrationDeleteResponseEnvelopeMessagesCode102048, CloudIntegrationDeleteResponseEnvelopeMessagesCode102049, CloudIntegrationDeleteResponseEnvelopeMessagesCode102050, CloudIntegrationDeleteResponseEnvelopeMessagesCode102051, CloudIntegrationDeleteResponseEnvelopeMessagesCode102052, CloudIntegrationDeleteResponseEnvelopeMessagesCode102053, CloudIntegrationDeleteResponseEnvelopeMessagesCode102054, CloudIntegrationDeleteResponseEnvelopeMessagesCode102055, CloudIntegrationDeleteResponseEnvelopeMessagesCode102056, CloudIntegrationDeleteResponseEnvelopeMessagesCode102057, CloudIntegrationDeleteResponseEnvelopeMessagesCode102058, CloudIntegrationDeleteResponseEnvelopeMessagesCode102059, CloudIntegrationDeleteResponseEnvelopeMessagesCode102060, CloudIntegrationDeleteResponseEnvelopeMessagesCode102061, CloudIntegrationDeleteResponseEnvelopeMessagesCode102062, CloudIntegrationDeleteResponseEnvelopeMessagesCode102063, CloudIntegrationDeleteResponseEnvelopeMessagesCode102064, CloudIntegrationDeleteResponseEnvelopeMessagesCode102065, CloudIntegrationDeleteResponseEnvelopeMessagesCode102066, CloudIntegrationDeleteResponseEnvelopeMessagesCode103001, CloudIntegrationDeleteResponseEnvelopeMessagesCode103002, CloudIntegrationDeleteResponseEnvelopeMessagesCode103003, CloudIntegrationDeleteResponseEnvelopeMessagesCode103004, CloudIntegrationDeleteResponseEnvelopeMessagesCode103005, CloudIntegrationDeleteResponseEnvelopeMessagesCode103006, CloudIntegrationDeleteResponseEnvelopeMessagesCode103007, CloudIntegrationDeleteResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationDeleteResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                 `json:"l10n_key"`
	LoggableError string                                                 `json:"loggable_error"`
	TemplateData  interface{}                                            `json:"template_data"`
	TraceID       string                                                 `json:"trace_id"`
	JSON          cloudIntegrationDeleteResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct [CloudIntegrationDeleteResponseEnvelopeMessagesMeta]
type cloudIntegrationDeleteResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDeleteResponseEnvelopeMessagesSource struct {
	Parameter           string                                                   `json:"parameter"`
	ParameterValueIndex int64                                                    `json:"parameter_value_index"`
	Pointer             string                                                   `json:"pointer"`
	JSON                cloudIntegrationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CloudIntegrationDeleteResponseEnvelopeMessagesSource]
type cloudIntegrationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationDiscoverParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	V2        param.Field[bool]   `query:"v2"`
}

// URLQuery serializes [CloudIntegrationDiscoverParams]'s query parameters as
// `url.Values`.
func (r CloudIntegrationDiscoverParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CloudIntegrationDiscoverAllParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CloudIntegrationEditParams struct {
	AccountID              param.Field[string] `path:"account_id,required"`
	AwsArn                 param.Field[string] `json:"aws_arn"`
	AzureSubscriptionID    param.Field[string] `json:"azure_subscription_id"`
	AzureTenantID          param.Field[string] `json:"azure_tenant_id"`
	Description            param.Field[string] `json:"description"`
	FriendlyName           param.Field[string] `json:"friendly_name"`
	GcpProjectID           param.Field[string] `json:"gcp_project_id"`
	GcpServiceAccountEmail param.Field[string] `json:"gcp_service_account_email"`
}

func (r CloudIntegrationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudIntegrationEditResponseEnvelope struct {
	Errors   []CloudIntegrationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationEditResponse                   `json:"result,required"`
	Success  bool                                           `json:"success,required"`
	JSON     cloudIntegrationEditResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudIntegrationEditResponseEnvelope]
type cloudIntegrationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeErrors struct {
	Code             CloudIntegrationEditResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Meta             CloudIntegrationEditResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CloudIntegrationEditResponseEnvelopeErrors]
type cloudIntegrationEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationEditResponseEnvelopeErrorsCode1001   CloudIntegrationEditResponseEnvelopeErrorsCode = 1001
	CloudIntegrationEditResponseEnvelopeErrorsCode1002   CloudIntegrationEditResponseEnvelopeErrorsCode = 1002
	CloudIntegrationEditResponseEnvelopeErrorsCode1003   CloudIntegrationEditResponseEnvelopeErrorsCode = 1003
	CloudIntegrationEditResponseEnvelopeErrorsCode1004   CloudIntegrationEditResponseEnvelopeErrorsCode = 1004
	CloudIntegrationEditResponseEnvelopeErrorsCode1005   CloudIntegrationEditResponseEnvelopeErrorsCode = 1005
	CloudIntegrationEditResponseEnvelopeErrorsCode1006   CloudIntegrationEditResponseEnvelopeErrorsCode = 1006
	CloudIntegrationEditResponseEnvelopeErrorsCode1007   CloudIntegrationEditResponseEnvelopeErrorsCode = 1007
	CloudIntegrationEditResponseEnvelopeErrorsCode1008   CloudIntegrationEditResponseEnvelopeErrorsCode = 1008
	CloudIntegrationEditResponseEnvelopeErrorsCode1009   CloudIntegrationEditResponseEnvelopeErrorsCode = 1009
	CloudIntegrationEditResponseEnvelopeErrorsCode1010   CloudIntegrationEditResponseEnvelopeErrorsCode = 1010
	CloudIntegrationEditResponseEnvelopeErrorsCode1011   CloudIntegrationEditResponseEnvelopeErrorsCode = 1011
	CloudIntegrationEditResponseEnvelopeErrorsCode1012   CloudIntegrationEditResponseEnvelopeErrorsCode = 1012
	CloudIntegrationEditResponseEnvelopeErrorsCode1013   CloudIntegrationEditResponseEnvelopeErrorsCode = 1013
	CloudIntegrationEditResponseEnvelopeErrorsCode1014   CloudIntegrationEditResponseEnvelopeErrorsCode = 1014
	CloudIntegrationEditResponseEnvelopeErrorsCode1015   CloudIntegrationEditResponseEnvelopeErrorsCode = 1015
	CloudIntegrationEditResponseEnvelopeErrorsCode1016   CloudIntegrationEditResponseEnvelopeErrorsCode = 1016
	CloudIntegrationEditResponseEnvelopeErrorsCode1017   CloudIntegrationEditResponseEnvelopeErrorsCode = 1017
	CloudIntegrationEditResponseEnvelopeErrorsCode2001   CloudIntegrationEditResponseEnvelopeErrorsCode = 2001
	CloudIntegrationEditResponseEnvelopeErrorsCode2002   CloudIntegrationEditResponseEnvelopeErrorsCode = 2002
	CloudIntegrationEditResponseEnvelopeErrorsCode2003   CloudIntegrationEditResponseEnvelopeErrorsCode = 2003
	CloudIntegrationEditResponseEnvelopeErrorsCode2004   CloudIntegrationEditResponseEnvelopeErrorsCode = 2004
	CloudIntegrationEditResponseEnvelopeErrorsCode2005   CloudIntegrationEditResponseEnvelopeErrorsCode = 2005
	CloudIntegrationEditResponseEnvelopeErrorsCode2006   CloudIntegrationEditResponseEnvelopeErrorsCode = 2006
	CloudIntegrationEditResponseEnvelopeErrorsCode2007   CloudIntegrationEditResponseEnvelopeErrorsCode = 2007
	CloudIntegrationEditResponseEnvelopeErrorsCode2008   CloudIntegrationEditResponseEnvelopeErrorsCode = 2008
	CloudIntegrationEditResponseEnvelopeErrorsCode2009   CloudIntegrationEditResponseEnvelopeErrorsCode = 2009
	CloudIntegrationEditResponseEnvelopeErrorsCode2010   CloudIntegrationEditResponseEnvelopeErrorsCode = 2010
	CloudIntegrationEditResponseEnvelopeErrorsCode2011   CloudIntegrationEditResponseEnvelopeErrorsCode = 2011
	CloudIntegrationEditResponseEnvelopeErrorsCode2012   CloudIntegrationEditResponseEnvelopeErrorsCode = 2012
	CloudIntegrationEditResponseEnvelopeErrorsCode2013   CloudIntegrationEditResponseEnvelopeErrorsCode = 2013
	CloudIntegrationEditResponseEnvelopeErrorsCode2014   CloudIntegrationEditResponseEnvelopeErrorsCode = 2014
	CloudIntegrationEditResponseEnvelopeErrorsCode2015   CloudIntegrationEditResponseEnvelopeErrorsCode = 2015
	CloudIntegrationEditResponseEnvelopeErrorsCode2016   CloudIntegrationEditResponseEnvelopeErrorsCode = 2016
	CloudIntegrationEditResponseEnvelopeErrorsCode2017   CloudIntegrationEditResponseEnvelopeErrorsCode = 2017
	CloudIntegrationEditResponseEnvelopeErrorsCode2018   CloudIntegrationEditResponseEnvelopeErrorsCode = 2018
	CloudIntegrationEditResponseEnvelopeErrorsCode2019   CloudIntegrationEditResponseEnvelopeErrorsCode = 2019
	CloudIntegrationEditResponseEnvelopeErrorsCode2020   CloudIntegrationEditResponseEnvelopeErrorsCode = 2020
	CloudIntegrationEditResponseEnvelopeErrorsCode2021   CloudIntegrationEditResponseEnvelopeErrorsCode = 2021
	CloudIntegrationEditResponseEnvelopeErrorsCode2022   CloudIntegrationEditResponseEnvelopeErrorsCode = 2022
	CloudIntegrationEditResponseEnvelopeErrorsCode3001   CloudIntegrationEditResponseEnvelopeErrorsCode = 3001
	CloudIntegrationEditResponseEnvelopeErrorsCode3002   CloudIntegrationEditResponseEnvelopeErrorsCode = 3002
	CloudIntegrationEditResponseEnvelopeErrorsCode3003   CloudIntegrationEditResponseEnvelopeErrorsCode = 3003
	CloudIntegrationEditResponseEnvelopeErrorsCode3004   CloudIntegrationEditResponseEnvelopeErrorsCode = 3004
	CloudIntegrationEditResponseEnvelopeErrorsCode3005   CloudIntegrationEditResponseEnvelopeErrorsCode = 3005
	CloudIntegrationEditResponseEnvelopeErrorsCode3006   CloudIntegrationEditResponseEnvelopeErrorsCode = 3006
	CloudIntegrationEditResponseEnvelopeErrorsCode3007   CloudIntegrationEditResponseEnvelopeErrorsCode = 3007
	CloudIntegrationEditResponseEnvelopeErrorsCode4001   CloudIntegrationEditResponseEnvelopeErrorsCode = 4001
	CloudIntegrationEditResponseEnvelopeErrorsCode4002   CloudIntegrationEditResponseEnvelopeErrorsCode = 4002
	CloudIntegrationEditResponseEnvelopeErrorsCode4003   CloudIntegrationEditResponseEnvelopeErrorsCode = 4003
	CloudIntegrationEditResponseEnvelopeErrorsCode4004   CloudIntegrationEditResponseEnvelopeErrorsCode = 4004
	CloudIntegrationEditResponseEnvelopeErrorsCode4005   CloudIntegrationEditResponseEnvelopeErrorsCode = 4005
	CloudIntegrationEditResponseEnvelopeErrorsCode4006   CloudIntegrationEditResponseEnvelopeErrorsCode = 4006
	CloudIntegrationEditResponseEnvelopeErrorsCode4007   CloudIntegrationEditResponseEnvelopeErrorsCode = 4007
	CloudIntegrationEditResponseEnvelopeErrorsCode4008   CloudIntegrationEditResponseEnvelopeErrorsCode = 4008
	CloudIntegrationEditResponseEnvelopeErrorsCode4009   CloudIntegrationEditResponseEnvelopeErrorsCode = 4009
	CloudIntegrationEditResponseEnvelopeErrorsCode4010   CloudIntegrationEditResponseEnvelopeErrorsCode = 4010
	CloudIntegrationEditResponseEnvelopeErrorsCode4011   CloudIntegrationEditResponseEnvelopeErrorsCode = 4011
	CloudIntegrationEditResponseEnvelopeErrorsCode4012   CloudIntegrationEditResponseEnvelopeErrorsCode = 4012
	CloudIntegrationEditResponseEnvelopeErrorsCode4013   CloudIntegrationEditResponseEnvelopeErrorsCode = 4013
	CloudIntegrationEditResponseEnvelopeErrorsCode4014   CloudIntegrationEditResponseEnvelopeErrorsCode = 4014
	CloudIntegrationEditResponseEnvelopeErrorsCode4015   CloudIntegrationEditResponseEnvelopeErrorsCode = 4015
	CloudIntegrationEditResponseEnvelopeErrorsCode4016   CloudIntegrationEditResponseEnvelopeErrorsCode = 4016
	CloudIntegrationEditResponseEnvelopeErrorsCode4017   CloudIntegrationEditResponseEnvelopeErrorsCode = 4017
	CloudIntegrationEditResponseEnvelopeErrorsCode4018   CloudIntegrationEditResponseEnvelopeErrorsCode = 4018
	CloudIntegrationEditResponseEnvelopeErrorsCode4019   CloudIntegrationEditResponseEnvelopeErrorsCode = 4019
	CloudIntegrationEditResponseEnvelopeErrorsCode4020   CloudIntegrationEditResponseEnvelopeErrorsCode = 4020
	CloudIntegrationEditResponseEnvelopeErrorsCode4021   CloudIntegrationEditResponseEnvelopeErrorsCode = 4021
	CloudIntegrationEditResponseEnvelopeErrorsCode4022   CloudIntegrationEditResponseEnvelopeErrorsCode = 4022
	CloudIntegrationEditResponseEnvelopeErrorsCode4023   CloudIntegrationEditResponseEnvelopeErrorsCode = 4023
	CloudIntegrationEditResponseEnvelopeErrorsCode5001   CloudIntegrationEditResponseEnvelopeErrorsCode = 5001
	CloudIntegrationEditResponseEnvelopeErrorsCode5002   CloudIntegrationEditResponseEnvelopeErrorsCode = 5002
	CloudIntegrationEditResponseEnvelopeErrorsCode5003   CloudIntegrationEditResponseEnvelopeErrorsCode = 5003
	CloudIntegrationEditResponseEnvelopeErrorsCode5004   CloudIntegrationEditResponseEnvelopeErrorsCode = 5004
	CloudIntegrationEditResponseEnvelopeErrorsCode102000 CloudIntegrationEditResponseEnvelopeErrorsCode = 102000
	CloudIntegrationEditResponseEnvelopeErrorsCode102001 CloudIntegrationEditResponseEnvelopeErrorsCode = 102001
	CloudIntegrationEditResponseEnvelopeErrorsCode102002 CloudIntegrationEditResponseEnvelopeErrorsCode = 102002
	CloudIntegrationEditResponseEnvelopeErrorsCode102003 CloudIntegrationEditResponseEnvelopeErrorsCode = 102003
	CloudIntegrationEditResponseEnvelopeErrorsCode102004 CloudIntegrationEditResponseEnvelopeErrorsCode = 102004
	CloudIntegrationEditResponseEnvelopeErrorsCode102005 CloudIntegrationEditResponseEnvelopeErrorsCode = 102005
	CloudIntegrationEditResponseEnvelopeErrorsCode102006 CloudIntegrationEditResponseEnvelopeErrorsCode = 102006
	CloudIntegrationEditResponseEnvelopeErrorsCode102007 CloudIntegrationEditResponseEnvelopeErrorsCode = 102007
	CloudIntegrationEditResponseEnvelopeErrorsCode102008 CloudIntegrationEditResponseEnvelopeErrorsCode = 102008
	CloudIntegrationEditResponseEnvelopeErrorsCode102009 CloudIntegrationEditResponseEnvelopeErrorsCode = 102009
	CloudIntegrationEditResponseEnvelopeErrorsCode102010 CloudIntegrationEditResponseEnvelopeErrorsCode = 102010
	CloudIntegrationEditResponseEnvelopeErrorsCode102011 CloudIntegrationEditResponseEnvelopeErrorsCode = 102011
	CloudIntegrationEditResponseEnvelopeErrorsCode102012 CloudIntegrationEditResponseEnvelopeErrorsCode = 102012
	CloudIntegrationEditResponseEnvelopeErrorsCode102013 CloudIntegrationEditResponseEnvelopeErrorsCode = 102013
	CloudIntegrationEditResponseEnvelopeErrorsCode102014 CloudIntegrationEditResponseEnvelopeErrorsCode = 102014
	CloudIntegrationEditResponseEnvelopeErrorsCode102015 CloudIntegrationEditResponseEnvelopeErrorsCode = 102015
	CloudIntegrationEditResponseEnvelopeErrorsCode102016 CloudIntegrationEditResponseEnvelopeErrorsCode = 102016
	CloudIntegrationEditResponseEnvelopeErrorsCode102017 CloudIntegrationEditResponseEnvelopeErrorsCode = 102017
	CloudIntegrationEditResponseEnvelopeErrorsCode102018 CloudIntegrationEditResponseEnvelopeErrorsCode = 102018
	CloudIntegrationEditResponseEnvelopeErrorsCode102019 CloudIntegrationEditResponseEnvelopeErrorsCode = 102019
	CloudIntegrationEditResponseEnvelopeErrorsCode102020 CloudIntegrationEditResponseEnvelopeErrorsCode = 102020
	CloudIntegrationEditResponseEnvelopeErrorsCode102021 CloudIntegrationEditResponseEnvelopeErrorsCode = 102021
	CloudIntegrationEditResponseEnvelopeErrorsCode102022 CloudIntegrationEditResponseEnvelopeErrorsCode = 102022
	CloudIntegrationEditResponseEnvelopeErrorsCode102023 CloudIntegrationEditResponseEnvelopeErrorsCode = 102023
	CloudIntegrationEditResponseEnvelopeErrorsCode102024 CloudIntegrationEditResponseEnvelopeErrorsCode = 102024
	CloudIntegrationEditResponseEnvelopeErrorsCode102025 CloudIntegrationEditResponseEnvelopeErrorsCode = 102025
	CloudIntegrationEditResponseEnvelopeErrorsCode102026 CloudIntegrationEditResponseEnvelopeErrorsCode = 102026
	CloudIntegrationEditResponseEnvelopeErrorsCode102027 CloudIntegrationEditResponseEnvelopeErrorsCode = 102027
	CloudIntegrationEditResponseEnvelopeErrorsCode102028 CloudIntegrationEditResponseEnvelopeErrorsCode = 102028
	CloudIntegrationEditResponseEnvelopeErrorsCode102029 CloudIntegrationEditResponseEnvelopeErrorsCode = 102029
	CloudIntegrationEditResponseEnvelopeErrorsCode102030 CloudIntegrationEditResponseEnvelopeErrorsCode = 102030
	CloudIntegrationEditResponseEnvelopeErrorsCode102031 CloudIntegrationEditResponseEnvelopeErrorsCode = 102031
	CloudIntegrationEditResponseEnvelopeErrorsCode102032 CloudIntegrationEditResponseEnvelopeErrorsCode = 102032
	CloudIntegrationEditResponseEnvelopeErrorsCode102033 CloudIntegrationEditResponseEnvelopeErrorsCode = 102033
	CloudIntegrationEditResponseEnvelopeErrorsCode102034 CloudIntegrationEditResponseEnvelopeErrorsCode = 102034
	CloudIntegrationEditResponseEnvelopeErrorsCode102035 CloudIntegrationEditResponseEnvelopeErrorsCode = 102035
	CloudIntegrationEditResponseEnvelopeErrorsCode102036 CloudIntegrationEditResponseEnvelopeErrorsCode = 102036
	CloudIntegrationEditResponseEnvelopeErrorsCode102037 CloudIntegrationEditResponseEnvelopeErrorsCode = 102037
	CloudIntegrationEditResponseEnvelopeErrorsCode102038 CloudIntegrationEditResponseEnvelopeErrorsCode = 102038
	CloudIntegrationEditResponseEnvelopeErrorsCode102039 CloudIntegrationEditResponseEnvelopeErrorsCode = 102039
	CloudIntegrationEditResponseEnvelopeErrorsCode102040 CloudIntegrationEditResponseEnvelopeErrorsCode = 102040
	CloudIntegrationEditResponseEnvelopeErrorsCode102041 CloudIntegrationEditResponseEnvelopeErrorsCode = 102041
	CloudIntegrationEditResponseEnvelopeErrorsCode102042 CloudIntegrationEditResponseEnvelopeErrorsCode = 102042
	CloudIntegrationEditResponseEnvelopeErrorsCode102043 CloudIntegrationEditResponseEnvelopeErrorsCode = 102043
	CloudIntegrationEditResponseEnvelopeErrorsCode102044 CloudIntegrationEditResponseEnvelopeErrorsCode = 102044
	CloudIntegrationEditResponseEnvelopeErrorsCode102045 CloudIntegrationEditResponseEnvelopeErrorsCode = 102045
	CloudIntegrationEditResponseEnvelopeErrorsCode102046 CloudIntegrationEditResponseEnvelopeErrorsCode = 102046
	CloudIntegrationEditResponseEnvelopeErrorsCode102047 CloudIntegrationEditResponseEnvelopeErrorsCode = 102047
	CloudIntegrationEditResponseEnvelopeErrorsCode102048 CloudIntegrationEditResponseEnvelopeErrorsCode = 102048
	CloudIntegrationEditResponseEnvelopeErrorsCode102049 CloudIntegrationEditResponseEnvelopeErrorsCode = 102049
	CloudIntegrationEditResponseEnvelopeErrorsCode102050 CloudIntegrationEditResponseEnvelopeErrorsCode = 102050
	CloudIntegrationEditResponseEnvelopeErrorsCode102051 CloudIntegrationEditResponseEnvelopeErrorsCode = 102051
	CloudIntegrationEditResponseEnvelopeErrorsCode102052 CloudIntegrationEditResponseEnvelopeErrorsCode = 102052
	CloudIntegrationEditResponseEnvelopeErrorsCode102053 CloudIntegrationEditResponseEnvelopeErrorsCode = 102053
	CloudIntegrationEditResponseEnvelopeErrorsCode102054 CloudIntegrationEditResponseEnvelopeErrorsCode = 102054
	CloudIntegrationEditResponseEnvelopeErrorsCode102055 CloudIntegrationEditResponseEnvelopeErrorsCode = 102055
	CloudIntegrationEditResponseEnvelopeErrorsCode102056 CloudIntegrationEditResponseEnvelopeErrorsCode = 102056
	CloudIntegrationEditResponseEnvelopeErrorsCode102057 CloudIntegrationEditResponseEnvelopeErrorsCode = 102057
	CloudIntegrationEditResponseEnvelopeErrorsCode102058 CloudIntegrationEditResponseEnvelopeErrorsCode = 102058
	CloudIntegrationEditResponseEnvelopeErrorsCode102059 CloudIntegrationEditResponseEnvelopeErrorsCode = 102059
	CloudIntegrationEditResponseEnvelopeErrorsCode102060 CloudIntegrationEditResponseEnvelopeErrorsCode = 102060
	CloudIntegrationEditResponseEnvelopeErrorsCode102061 CloudIntegrationEditResponseEnvelopeErrorsCode = 102061
	CloudIntegrationEditResponseEnvelopeErrorsCode102062 CloudIntegrationEditResponseEnvelopeErrorsCode = 102062
	CloudIntegrationEditResponseEnvelopeErrorsCode102063 CloudIntegrationEditResponseEnvelopeErrorsCode = 102063
	CloudIntegrationEditResponseEnvelopeErrorsCode102064 CloudIntegrationEditResponseEnvelopeErrorsCode = 102064
	CloudIntegrationEditResponseEnvelopeErrorsCode102065 CloudIntegrationEditResponseEnvelopeErrorsCode = 102065
	CloudIntegrationEditResponseEnvelopeErrorsCode102066 CloudIntegrationEditResponseEnvelopeErrorsCode = 102066
	CloudIntegrationEditResponseEnvelopeErrorsCode103001 CloudIntegrationEditResponseEnvelopeErrorsCode = 103001
	CloudIntegrationEditResponseEnvelopeErrorsCode103002 CloudIntegrationEditResponseEnvelopeErrorsCode = 103002
	CloudIntegrationEditResponseEnvelopeErrorsCode103003 CloudIntegrationEditResponseEnvelopeErrorsCode = 103003
	CloudIntegrationEditResponseEnvelopeErrorsCode103004 CloudIntegrationEditResponseEnvelopeErrorsCode = 103004
	CloudIntegrationEditResponseEnvelopeErrorsCode103005 CloudIntegrationEditResponseEnvelopeErrorsCode = 103005
	CloudIntegrationEditResponseEnvelopeErrorsCode103006 CloudIntegrationEditResponseEnvelopeErrorsCode = 103006
	CloudIntegrationEditResponseEnvelopeErrorsCode103007 CloudIntegrationEditResponseEnvelopeErrorsCode = 103007
	CloudIntegrationEditResponseEnvelopeErrorsCode103008 CloudIntegrationEditResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationEditResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseEnvelopeErrorsCode1001, CloudIntegrationEditResponseEnvelopeErrorsCode1002, CloudIntegrationEditResponseEnvelopeErrorsCode1003, CloudIntegrationEditResponseEnvelopeErrorsCode1004, CloudIntegrationEditResponseEnvelopeErrorsCode1005, CloudIntegrationEditResponseEnvelopeErrorsCode1006, CloudIntegrationEditResponseEnvelopeErrorsCode1007, CloudIntegrationEditResponseEnvelopeErrorsCode1008, CloudIntegrationEditResponseEnvelopeErrorsCode1009, CloudIntegrationEditResponseEnvelopeErrorsCode1010, CloudIntegrationEditResponseEnvelopeErrorsCode1011, CloudIntegrationEditResponseEnvelopeErrorsCode1012, CloudIntegrationEditResponseEnvelopeErrorsCode1013, CloudIntegrationEditResponseEnvelopeErrorsCode1014, CloudIntegrationEditResponseEnvelopeErrorsCode1015, CloudIntegrationEditResponseEnvelopeErrorsCode1016, CloudIntegrationEditResponseEnvelopeErrorsCode1017, CloudIntegrationEditResponseEnvelopeErrorsCode2001, CloudIntegrationEditResponseEnvelopeErrorsCode2002, CloudIntegrationEditResponseEnvelopeErrorsCode2003, CloudIntegrationEditResponseEnvelopeErrorsCode2004, CloudIntegrationEditResponseEnvelopeErrorsCode2005, CloudIntegrationEditResponseEnvelopeErrorsCode2006, CloudIntegrationEditResponseEnvelopeErrorsCode2007, CloudIntegrationEditResponseEnvelopeErrorsCode2008, CloudIntegrationEditResponseEnvelopeErrorsCode2009, CloudIntegrationEditResponseEnvelopeErrorsCode2010, CloudIntegrationEditResponseEnvelopeErrorsCode2011, CloudIntegrationEditResponseEnvelopeErrorsCode2012, CloudIntegrationEditResponseEnvelopeErrorsCode2013, CloudIntegrationEditResponseEnvelopeErrorsCode2014, CloudIntegrationEditResponseEnvelopeErrorsCode2015, CloudIntegrationEditResponseEnvelopeErrorsCode2016, CloudIntegrationEditResponseEnvelopeErrorsCode2017, CloudIntegrationEditResponseEnvelopeErrorsCode2018, CloudIntegrationEditResponseEnvelopeErrorsCode2019, CloudIntegrationEditResponseEnvelopeErrorsCode2020, CloudIntegrationEditResponseEnvelopeErrorsCode2021, CloudIntegrationEditResponseEnvelopeErrorsCode2022, CloudIntegrationEditResponseEnvelopeErrorsCode3001, CloudIntegrationEditResponseEnvelopeErrorsCode3002, CloudIntegrationEditResponseEnvelopeErrorsCode3003, CloudIntegrationEditResponseEnvelopeErrorsCode3004, CloudIntegrationEditResponseEnvelopeErrorsCode3005, CloudIntegrationEditResponseEnvelopeErrorsCode3006, CloudIntegrationEditResponseEnvelopeErrorsCode3007, CloudIntegrationEditResponseEnvelopeErrorsCode4001, CloudIntegrationEditResponseEnvelopeErrorsCode4002, CloudIntegrationEditResponseEnvelopeErrorsCode4003, CloudIntegrationEditResponseEnvelopeErrorsCode4004, CloudIntegrationEditResponseEnvelopeErrorsCode4005, CloudIntegrationEditResponseEnvelopeErrorsCode4006, CloudIntegrationEditResponseEnvelopeErrorsCode4007, CloudIntegrationEditResponseEnvelopeErrorsCode4008, CloudIntegrationEditResponseEnvelopeErrorsCode4009, CloudIntegrationEditResponseEnvelopeErrorsCode4010, CloudIntegrationEditResponseEnvelopeErrorsCode4011, CloudIntegrationEditResponseEnvelopeErrorsCode4012, CloudIntegrationEditResponseEnvelopeErrorsCode4013, CloudIntegrationEditResponseEnvelopeErrorsCode4014, CloudIntegrationEditResponseEnvelopeErrorsCode4015, CloudIntegrationEditResponseEnvelopeErrorsCode4016, CloudIntegrationEditResponseEnvelopeErrorsCode4017, CloudIntegrationEditResponseEnvelopeErrorsCode4018, CloudIntegrationEditResponseEnvelopeErrorsCode4019, CloudIntegrationEditResponseEnvelopeErrorsCode4020, CloudIntegrationEditResponseEnvelopeErrorsCode4021, CloudIntegrationEditResponseEnvelopeErrorsCode4022, CloudIntegrationEditResponseEnvelopeErrorsCode4023, CloudIntegrationEditResponseEnvelopeErrorsCode5001, CloudIntegrationEditResponseEnvelopeErrorsCode5002, CloudIntegrationEditResponseEnvelopeErrorsCode5003, CloudIntegrationEditResponseEnvelopeErrorsCode5004, CloudIntegrationEditResponseEnvelopeErrorsCode102000, CloudIntegrationEditResponseEnvelopeErrorsCode102001, CloudIntegrationEditResponseEnvelopeErrorsCode102002, CloudIntegrationEditResponseEnvelopeErrorsCode102003, CloudIntegrationEditResponseEnvelopeErrorsCode102004, CloudIntegrationEditResponseEnvelopeErrorsCode102005, CloudIntegrationEditResponseEnvelopeErrorsCode102006, CloudIntegrationEditResponseEnvelopeErrorsCode102007, CloudIntegrationEditResponseEnvelopeErrorsCode102008, CloudIntegrationEditResponseEnvelopeErrorsCode102009, CloudIntegrationEditResponseEnvelopeErrorsCode102010, CloudIntegrationEditResponseEnvelopeErrorsCode102011, CloudIntegrationEditResponseEnvelopeErrorsCode102012, CloudIntegrationEditResponseEnvelopeErrorsCode102013, CloudIntegrationEditResponseEnvelopeErrorsCode102014, CloudIntegrationEditResponseEnvelopeErrorsCode102015, CloudIntegrationEditResponseEnvelopeErrorsCode102016, CloudIntegrationEditResponseEnvelopeErrorsCode102017, CloudIntegrationEditResponseEnvelopeErrorsCode102018, CloudIntegrationEditResponseEnvelopeErrorsCode102019, CloudIntegrationEditResponseEnvelopeErrorsCode102020, CloudIntegrationEditResponseEnvelopeErrorsCode102021, CloudIntegrationEditResponseEnvelopeErrorsCode102022, CloudIntegrationEditResponseEnvelopeErrorsCode102023, CloudIntegrationEditResponseEnvelopeErrorsCode102024, CloudIntegrationEditResponseEnvelopeErrorsCode102025, CloudIntegrationEditResponseEnvelopeErrorsCode102026, CloudIntegrationEditResponseEnvelopeErrorsCode102027, CloudIntegrationEditResponseEnvelopeErrorsCode102028, CloudIntegrationEditResponseEnvelopeErrorsCode102029, CloudIntegrationEditResponseEnvelopeErrorsCode102030, CloudIntegrationEditResponseEnvelopeErrorsCode102031, CloudIntegrationEditResponseEnvelopeErrorsCode102032, CloudIntegrationEditResponseEnvelopeErrorsCode102033, CloudIntegrationEditResponseEnvelopeErrorsCode102034, CloudIntegrationEditResponseEnvelopeErrorsCode102035, CloudIntegrationEditResponseEnvelopeErrorsCode102036, CloudIntegrationEditResponseEnvelopeErrorsCode102037, CloudIntegrationEditResponseEnvelopeErrorsCode102038, CloudIntegrationEditResponseEnvelopeErrorsCode102039, CloudIntegrationEditResponseEnvelopeErrorsCode102040, CloudIntegrationEditResponseEnvelopeErrorsCode102041, CloudIntegrationEditResponseEnvelopeErrorsCode102042, CloudIntegrationEditResponseEnvelopeErrorsCode102043, CloudIntegrationEditResponseEnvelopeErrorsCode102044, CloudIntegrationEditResponseEnvelopeErrorsCode102045, CloudIntegrationEditResponseEnvelopeErrorsCode102046, CloudIntegrationEditResponseEnvelopeErrorsCode102047, CloudIntegrationEditResponseEnvelopeErrorsCode102048, CloudIntegrationEditResponseEnvelopeErrorsCode102049, CloudIntegrationEditResponseEnvelopeErrorsCode102050, CloudIntegrationEditResponseEnvelopeErrorsCode102051, CloudIntegrationEditResponseEnvelopeErrorsCode102052, CloudIntegrationEditResponseEnvelopeErrorsCode102053, CloudIntegrationEditResponseEnvelopeErrorsCode102054, CloudIntegrationEditResponseEnvelopeErrorsCode102055, CloudIntegrationEditResponseEnvelopeErrorsCode102056, CloudIntegrationEditResponseEnvelopeErrorsCode102057, CloudIntegrationEditResponseEnvelopeErrorsCode102058, CloudIntegrationEditResponseEnvelopeErrorsCode102059, CloudIntegrationEditResponseEnvelopeErrorsCode102060, CloudIntegrationEditResponseEnvelopeErrorsCode102061, CloudIntegrationEditResponseEnvelopeErrorsCode102062, CloudIntegrationEditResponseEnvelopeErrorsCode102063, CloudIntegrationEditResponseEnvelopeErrorsCode102064, CloudIntegrationEditResponseEnvelopeErrorsCode102065, CloudIntegrationEditResponseEnvelopeErrorsCode102066, CloudIntegrationEditResponseEnvelopeErrorsCode103001, CloudIntegrationEditResponseEnvelopeErrorsCode103002, CloudIntegrationEditResponseEnvelopeErrorsCode103003, CloudIntegrationEditResponseEnvelopeErrorsCode103004, CloudIntegrationEditResponseEnvelopeErrorsCode103005, CloudIntegrationEditResponseEnvelopeErrorsCode103006, CloudIntegrationEditResponseEnvelopeErrorsCode103007, CloudIntegrationEditResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationEditResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                             `json:"l10n_key"`
	LoggableError string                                             `json:"loggable_error"`
	TemplateData  interface{}                                        `json:"template_data"`
	TraceID       string                                             `json:"trace_id"`
	JSON          cloudIntegrationEditResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationEditResponseEnvelopeErrorsMeta]
type cloudIntegrationEditResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeErrorsSource struct {
	Parameter           string                                               `json:"parameter"`
	ParameterValueIndex int64                                                `json:"parameter_value_index"`
	Pointer             string                                               `json:"pointer"`
	JSON                cloudIntegrationEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationEditResponseEnvelopeErrorsSource]
type cloudIntegrationEditResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeMessages struct {
	Code             CloudIntegrationEditResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Meta             CloudIntegrationEditResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CloudIntegrationEditResponseEnvelopeMessages]
type cloudIntegrationEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationEditResponseEnvelopeMessagesCode1001   CloudIntegrationEditResponseEnvelopeMessagesCode = 1001
	CloudIntegrationEditResponseEnvelopeMessagesCode1002   CloudIntegrationEditResponseEnvelopeMessagesCode = 1002
	CloudIntegrationEditResponseEnvelopeMessagesCode1003   CloudIntegrationEditResponseEnvelopeMessagesCode = 1003
	CloudIntegrationEditResponseEnvelopeMessagesCode1004   CloudIntegrationEditResponseEnvelopeMessagesCode = 1004
	CloudIntegrationEditResponseEnvelopeMessagesCode1005   CloudIntegrationEditResponseEnvelopeMessagesCode = 1005
	CloudIntegrationEditResponseEnvelopeMessagesCode1006   CloudIntegrationEditResponseEnvelopeMessagesCode = 1006
	CloudIntegrationEditResponseEnvelopeMessagesCode1007   CloudIntegrationEditResponseEnvelopeMessagesCode = 1007
	CloudIntegrationEditResponseEnvelopeMessagesCode1008   CloudIntegrationEditResponseEnvelopeMessagesCode = 1008
	CloudIntegrationEditResponseEnvelopeMessagesCode1009   CloudIntegrationEditResponseEnvelopeMessagesCode = 1009
	CloudIntegrationEditResponseEnvelopeMessagesCode1010   CloudIntegrationEditResponseEnvelopeMessagesCode = 1010
	CloudIntegrationEditResponseEnvelopeMessagesCode1011   CloudIntegrationEditResponseEnvelopeMessagesCode = 1011
	CloudIntegrationEditResponseEnvelopeMessagesCode1012   CloudIntegrationEditResponseEnvelopeMessagesCode = 1012
	CloudIntegrationEditResponseEnvelopeMessagesCode1013   CloudIntegrationEditResponseEnvelopeMessagesCode = 1013
	CloudIntegrationEditResponseEnvelopeMessagesCode1014   CloudIntegrationEditResponseEnvelopeMessagesCode = 1014
	CloudIntegrationEditResponseEnvelopeMessagesCode1015   CloudIntegrationEditResponseEnvelopeMessagesCode = 1015
	CloudIntegrationEditResponseEnvelopeMessagesCode1016   CloudIntegrationEditResponseEnvelopeMessagesCode = 1016
	CloudIntegrationEditResponseEnvelopeMessagesCode1017   CloudIntegrationEditResponseEnvelopeMessagesCode = 1017
	CloudIntegrationEditResponseEnvelopeMessagesCode2001   CloudIntegrationEditResponseEnvelopeMessagesCode = 2001
	CloudIntegrationEditResponseEnvelopeMessagesCode2002   CloudIntegrationEditResponseEnvelopeMessagesCode = 2002
	CloudIntegrationEditResponseEnvelopeMessagesCode2003   CloudIntegrationEditResponseEnvelopeMessagesCode = 2003
	CloudIntegrationEditResponseEnvelopeMessagesCode2004   CloudIntegrationEditResponseEnvelopeMessagesCode = 2004
	CloudIntegrationEditResponseEnvelopeMessagesCode2005   CloudIntegrationEditResponseEnvelopeMessagesCode = 2005
	CloudIntegrationEditResponseEnvelopeMessagesCode2006   CloudIntegrationEditResponseEnvelopeMessagesCode = 2006
	CloudIntegrationEditResponseEnvelopeMessagesCode2007   CloudIntegrationEditResponseEnvelopeMessagesCode = 2007
	CloudIntegrationEditResponseEnvelopeMessagesCode2008   CloudIntegrationEditResponseEnvelopeMessagesCode = 2008
	CloudIntegrationEditResponseEnvelopeMessagesCode2009   CloudIntegrationEditResponseEnvelopeMessagesCode = 2009
	CloudIntegrationEditResponseEnvelopeMessagesCode2010   CloudIntegrationEditResponseEnvelopeMessagesCode = 2010
	CloudIntegrationEditResponseEnvelopeMessagesCode2011   CloudIntegrationEditResponseEnvelopeMessagesCode = 2011
	CloudIntegrationEditResponseEnvelopeMessagesCode2012   CloudIntegrationEditResponseEnvelopeMessagesCode = 2012
	CloudIntegrationEditResponseEnvelopeMessagesCode2013   CloudIntegrationEditResponseEnvelopeMessagesCode = 2013
	CloudIntegrationEditResponseEnvelopeMessagesCode2014   CloudIntegrationEditResponseEnvelopeMessagesCode = 2014
	CloudIntegrationEditResponseEnvelopeMessagesCode2015   CloudIntegrationEditResponseEnvelopeMessagesCode = 2015
	CloudIntegrationEditResponseEnvelopeMessagesCode2016   CloudIntegrationEditResponseEnvelopeMessagesCode = 2016
	CloudIntegrationEditResponseEnvelopeMessagesCode2017   CloudIntegrationEditResponseEnvelopeMessagesCode = 2017
	CloudIntegrationEditResponseEnvelopeMessagesCode2018   CloudIntegrationEditResponseEnvelopeMessagesCode = 2018
	CloudIntegrationEditResponseEnvelopeMessagesCode2019   CloudIntegrationEditResponseEnvelopeMessagesCode = 2019
	CloudIntegrationEditResponseEnvelopeMessagesCode2020   CloudIntegrationEditResponseEnvelopeMessagesCode = 2020
	CloudIntegrationEditResponseEnvelopeMessagesCode2021   CloudIntegrationEditResponseEnvelopeMessagesCode = 2021
	CloudIntegrationEditResponseEnvelopeMessagesCode2022   CloudIntegrationEditResponseEnvelopeMessagesCode = 2022
	CloudIntegrationEditResponseEnvelopeMessagesCode3001   CloudIntegrationEditResponseEnvelopeMessagesCode = 3001
	CloudIntegrationEditResponseEnvelopeMessagesCode3002   CloudIntegrationEditResponseEnvelopeMessagesCode = 3002
	CloudIntegrationEditResponseEnvelopeMessagesCode3003   CloudIntegrationEditResponseEnvelopeMessagesCode = 3003
	CloudIntegrationEditResponseEnvelopeMessagesCode3004   CloudIntegrationEditResponseEnvelopeMessagesCode = 3004
	CloudIntegrationEditResponseEnvelopeMessagesCode3005   CloudIntegrationEditResponseEnvelopeMessagesCode = 3005
	CloudIntegrationEditResponseEnvelopeMessagesCode3006   CloudIntegrationEditResponseEnvelopeMessagesCode = 3006
	CloudIntegrationEditResponseEnvelopeMessagesCode3007   CloudIntegrationEditResponseEnvelopeMessagesCode = 3007
	CloudIntegrationEditResponseEnvelopeMessagesCode4001   CloudIntegrationEditResponseEnvelopeMessagesCode = 4001
	CloudIntegrationEditResponseEnvelopeMessagesCode4002   CloudIntegrationEditResponseEnvelopeMessagesCode = 4002
	CloudIntegrationEditResponseEnvelopeMessagesCode4003   CloudIntegrationEditResponseEnvelopeMessagesCode = 4003
	CloudIntegrationEditResponseEnvelopeMessagesCode4004   CloudIntegrationEditResponseEnvelopeMessagesCode = 4004
	CloudIntegrationEditResponseEnvelopeMessagesCode4005   CloudIntegrationEditResponseEnvelopeMessagesCode = 4005
	CloudIntegrationEditResponseEnvelopeMessagesCode4006   CloudIntegrationEditResponseEnvelopeMessagesCode = 4006
	CloudIntegrationEditResponseEnvelopeMessagesCode4007   CloudIntegrationEditResponseEnvelopeMessagesCode = 4007
	CloudIntegrationEditResponseEnvelopeMessagesCode4008   CloudIntegrationEditResponseEnvelopeMessagesCode = 4008
	CloudIntegrationEditResponseEnvelopeMessagesCode4009   CloudIntegrationEditResponseEnvelopeMessagesCode = 4009
	CloudIntegrationEditResponseEnvelopeMessagesCode4010   CloudIntegrationEditResponseEnvelopeMessagesCode = 4010
	CloudIntegrationEditResponseEnvelopeMessagesCode4011   CloudIntegrationEditResponseEnvelopeMessagesCode = 4011
	CloudIntegrationEditResponseEnvelopeMessagesCode4012   CloudIntegrationEditResponseEnvelopeMessagesCode = 4012
	CloudIntegrationEditResponseEnvelopeMessagesCode4013   CloudIntegrationEditResponseEnvelopeMessagesCode = 4013
	CloudIntegrationEditResponseEnvelopeMessagesCode4014   CloudIntegrationEditResponseEnvelopeMessagesCode = 4014
	CloudIntegrationEditResponseEnvelopeMessagesCode4015   CloudIntegrationEditResponseEnvelopeMessagesCode = 4015
	CloudIntegrationEditResponseEnvelopeMessagesCode4016   CloudIntegrationEditResponseEnvelopeMessagesCode = 4016
	CloudIntegrationEditResponseEnvelopeMessagesCode4017   CloudIntegrationEditResponseEnvelopeMessagesCode = 4017
	CloudIntegrationEditResponseEnvelopeMessagesCode4018   CloudIntegrationEditResponseEnvelopeMessagesCode = 4018
	CloudIntegrationEditResponseEnvelopeMessagesCode4019   CloudIntegrationEditResponseEnvelopeMessagesCode = 4019
	CloudIntegrationEditResponseEnvelopeMessagesCode4020   CloudIntegrationEditResponseEnvelopeMessagesCode = 4020
	CloudIntegrationEditResponseEnvelopeMessagesCode4021   CloudIntegrationEditResponseEnvelopeMessagesCode = 4021
	CloudIntegrationEditResponseEnvelopeMessagesCode4022   CloudIntegrationEditResponseEnvelopeMessagesCode = 4022
	CloudIntegrationEditResponseEnvelopeMessagesCode4023   CloudIntegrationEditResponseEnvelopeMessagesCode = 4023
	CloudIntegrationEditResponseEnvelopeMessagesCode5001   CloudIntegrationEditResponseEnvelopeMessagesCode = 5001
	CloudIntegrationEditResponseEnvelopeMessagesCode5002   CloudIntegrationEditResponseEnvelopeMessagesCode = 5002
	CloudIntegrationEditResponseEnvelopeMessagesCode5003   CloudIntegrationEditResponseEnvelopeMessagesCode = 5003
	CloudIntegrationEditResponseEnvelopeMessagesCode5004   CloudIntegrationEditResponseEnvelopeMessagesCode = 5004
	CloudIntegrationEditResponseEnvelopeMessagesCode102000 CloudIntegrationEditResponseEnvelopeMessagesCode = 102000
	CloudIntegrationEditResponseEnvelopeMessagesCode102001 CloudIntegrationEditResponseEnvelopeMessagesCode = 102001
	CloudIntegrationEditResponseEnvelopeMessagesCode102002 CloudIntegrationEditResponseEnvelopeMessagesCode = 102002
	CloudIntegrationEditResponseEnvelopeMessagesCode102003 CloudIntegrationEditResponseEnvelopeMessagesCode = 102003
	CloudIntegrationEditResponseEnvelopeMessagesCode102004 CloudIntegrationEditResponseEnvelopeMessagesCode = 102004
	CloudIntegrationEditResponseEnvelopeMessagesCode102005 CloudIntegrationEditResponseEnvelopeMessagesCode = 102005
	CloudIntegrationEditResponseEnvelopeMessagesCode102006 CloudIntegrationEditResponseEnvelopeMessagesCode = 102006
	CloudIntegrationEditResponseEnvelopeMessagesCode102007 CloudIntegrationEditResponseEnvelopeMessagesCode = 102007
	CloudIntegrationEditResponseEnvelopeMessagesCode102008 CloudIntegrationEditResponseEnvelopeMessagesCode = 102008
	CloudIntegrationEditResponseEnvelopeMessagesCode102009 CloudIntegrationEditResponseEnvelopeMessagesCode = 102009
	CloudIntegrationEditResponseEnvelopeMessagesCode102010 CloudIntegrationEditResponseEnvelopeMessagesCode = 102010
	CloudIntegrationEditResponseEnvelopeMessagesCode102011 CloudIntegrationEditResponseEnvelopeMessagesCode = 102011
	CloudIntegrationEditResponseEnvelopeMessagesCode102012 CloudIntegrationEditResponseEnvelopeMessagesCode = 102012
	CloudIntegrationEditResponseEnvelopeMessagesCode102013 CloudIntegrationEditResponseEnvelopeMessagesCode = 102013
	CloudIntegrationEditResponseEnvelopeMessagesCode102014 CloudIntegrationEditResponseEnvelopeMessagesCode = 102014
	CloudIntegrationEditResponseEnvelopeMessagesCode102015 CloudIntegrationEditResponseEnvelopeMessagesCode = 102015
	CloudIntegrationEditResponseEnvelopeMessagesCode102016 CloudIntegrationEditResponseEnvelopeMessagesCode = 102016
	CloudIntegrationEditResponseEnvelopeMessagesCode102017 CloudIntegrationEditResponseEnvelopeMessagesCode = 102017
	CloudIntegrationEditResponseEnvelopeMessagesCode102018 CloudIntegrationEditResponseEnvelopeMessagesCode = 102018
	CloudIntegrationEditResponseEnvelopeMessagesCode102019 CloudIntegrationEditResponseEnvelopeMessagesCode = 102019
	CloudIntegrationEditResponseEnvelopeMessagesCode102020 CloudIntegrationEditResponseEnvelopeMessagesCode = 102020
	CloudIntegrationEditResponseEnvelopeMessagesCode102021 CloudIntegrationEditResponseEnvelopeMessagesCode = 102021
	CloudIntegrationEditResponseEnvelopeMessagesCode102022 CloudIntegrationEditResponseEnvelopeMessagesCode = 102022
	CloudIntegrationEditResponseEnvelopeMessagesCode102023 CloudIntegrationEditResponseEnvelopeMessagesCode = 102023
	CloudIntegrationEditResponseEnvelopeMessagesCode102024 CloudIntegrationEditResponseEnvelopeMessagesCode = 102024
	CloudIntegrationEditResponseEnvelopeMessagesCode102025 CloudIntegrationEditResponseEnvelopeMessagesCode = 102025
	CloudIntegrationEditResponseEnvelopeMessagesCode102026 CloudIntegrationEditResponseEnvelopeMessagesCode = 102026
	CloudIntegrationEditResponseEnvelopeMessagesCode102027 CloudIntegrationEditResponseEnvelopeMessagesCode = 102027
	CloudIntegrationEditResponseEnvelopeMessagesCode102028 CloudIntegrationEditResponseEnvelopeMessagesCode = 102028
	CloudIntegrationEditResponseEnvelopeMessagesCode102029 CloudIntegrationEditResponseEnvelopeMessagesCode = 102029
	CloudIntegrationEditResponseEnvelopeMessagesCode102030 CloudIntegrationEditResponseEnvelopeMessagesCode = 102030
	CloudIntegrationEditResponseEnvelopeMessagesCode102031 CloudIntegrationEditResponseEnvelopeMessagesCode = 102031
	CloudIntegrationEditResponseEnvelopeMessagesCode102032 CloudIntegrationEditResponseEnvelopeMessagesCode = 102032
	CloudIntegrationEditResponseEnvelopeMessagesCode102033 CloudIntegrationEditResponseEnvelopeMessagesCode = 102033
	CloudIntegrationEditResponseEnvelopeMessagesCode102034 CloudIntegrationEditResponseEnvelopeMessagesCode = 102034
	CloudIntegrationEditResponseEnvelopeMessagesCode102035 CloudIntegrationEditResponseEnvelopeMessagesCode = 102035
	CloudIntegrationEditResponseEnvelopeMessagesCode102036 CloudIntegrationEditResponseEnvelopeMessagesCode = 102036
	CloudIntegrationEditResponseEnvelopeMessagesCode102037 CloudIntegrationEditResponseEnvelopeMessagesCode = 102037
	CloudIntegrationEditResponseEnvelopeMessagesCode102038 CloudIntegrationEditResponseEnvelopeMessagesCode = 102038
	CloudIntegrationEditResponseEnvelopeMessagesCode102039 CloudIntegrationEditResponseEnvelopeMessagesCode = 102039
	CloudIntegrationEditResponseEnvelopeMessagesCode102040 CloudIntegrationEditResponseEnvelopeMessagesCode = 102040
	CloudIntegrationEditResponseEnvelopeMessagesCode102041 CloudIntegrationEditResponseEnvelopeMessagesCode = 102041
	CloudIntegrationEditResponseEnvelopeMessagesCode102042 CloudIntegrationEditResponseEnvelopeMessagesCode = 102042
	CloudIntegrationEditResponseEnvelopeMessagesCode102043 CloudIntegrationEditResponseEnvelopeMessagesCode = 102043
	CloudIntegrationEditResponseEnvelopeMessagesCode102044 CloudIntegrationEditResponseEnvelopeMessagesCode = 102044
	CloudIntegrationEditResponseEnvelopeMessagesCode102045 CloudIntegrationEditResponseEnvelopeMessagesCode = 102045
	CloudIntegrationEditResponseEnvelopeMessagesCode102046 CloudIntegrationEditResponseEnvelopeMessagesCode = 102046
	CloudIntegrationEditResponseEnvelopeMessagesCode102047 CloudIntegrationEditResponseEnvelopeMessagesCode = 102047
	CloudIntegrationEditResponseEnvelopeMessagesCode102048 CloudIntegrationEditResponseEnvelopeMessagesCode = 102048
	CloudIntegrationEditResponseEnvelopeMessagesCode102049 CloudIntegrationEditResponseEnvelopeMessagesCode = 102049
	CloudIntegrationEditResponseEnvelopeMessagesCode102050 CloudIntegrationEditResponseEnvelopeMessagesCode = 102050
	CloudIntegrationEditResponseEnvelopeMessagesCode102051 CloudIntegrationEditResponseEnvelopeMessagesCode = 102051
	CloudIntegrationEditResponseEnvelopeMessagesCode102052 CloudIntegrationEditResponseEnvelopeMessagesCode = 102052
	CloudIntegrationEditResponseEnvelopeMessagesCode102053 CloudIntegrationEditResponseEnvelopeMessagesCode = 102053
	CloudIntegrationEditResponseEnvelopeMessagesCode102054 CloudIntegrationEditResponseEnvelopeMessagesCode = 102054
	CloudIntegrationEditResponseEnvelopeMessagesCode102055 CloudIntegrationEditResponseEnvelopeMessagesCode = 102055
	CloudIntegrationEditResponseEnvelopeMessagesCode102056 CloudIntegrationEditResponseEnvelopeMessagesCode = 102056
	CloudIntegrationEditResponseEnvelopeMessagesCode102057 CloudIntegrationEditResponseEnvelopeMessagesCode = 102057
	CloudIntegrationEditResponseEnvelopeMessagesCode102058 CloudIntegrationEditResponseEnvelopeMessagesCode = 102058
	CloudIntegrationEditResponseEnvelopeMessagesCode102059 CloudIntegrationEditResponseEnvelopeMessagesCode = 102059
	CloudIntegrationEditResponseEnvelopeMessagesCode102060 CloudIntegrationEditResponseEnvelopeMessagesCode = 102060
	CloudIntegrationEditResponseEnvelopeMessagesCode102061 CloudIntegrationEditResponseEnvelopeMessagesCode = 102061
	CloudIntegrationEditResponseEnvelopeMessagesCode102062 CloudIntegrationEditResponseEnvelopeMessagesCode = 102062
	CloudIntegrationEditResponseEnvelopeMessagesCode102063 CloudIntegrationEditResponseEnvelopeMessagesCode = 102063
	CloudIntegrationEditResponseEnvelopeMessagesCode102064 CloudIntegrationEditResponseEnvelopeMessagesCode = 102064
	CloudIntegrationEditResponseEnvelopeMessagesCode102065 CloudIntegrationEditResponseEnvelopeMessagesCode = 102065
	CloudIntegrationEditResponseEnvelopeMessagesCode102066 CloudIntegrationEditResponseEnvelopeMessagesCode = 102066
	CloudIntegrationEditResponseEnvelopeMessagesCode103001 CloudIntegrationEditResponseEnvelopeMessagesCode = 103001
	CloudIntegrationEditResponseEnvelopeMessagesCode103002 CloudIntegrationEditResponseEnvelopeMessagesCode = 103002
	CloudIntegrationEditResponseEnvelopeMessagesCode103003 CloudIntegrationEditResponseEnvelopeMessagesCode = 103003
	CloudIntegrationEditResponseEnvelopeMessagesCode103004 CloudIntegrationEditResponseEnvelopeMessagesCode = 103004
	CloudIntegrationEditResponseEnvelopeMessagesCode103005 CloudIntegrationEditResponseEnvelopeMessagesCode = 103005
	CloudIntegrationEditResponseEnvelopeMessagesCode103006 CloudIntegrationEditResponseEnvelopeMessagesCode = 103006
	CloudIntegrationEditResponseEnvelopeMessagesCode103007 CloudIntegrationEditResponseEnvelopeMessagesCode = 103007
	CloudIntegrationEditResponseEnvelopeMessagesCode103008 CloudIntegrationEditResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationEditResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationEditResponseEnvelopeMessagesCode1001, CloudIntegrationEditResponseEnvelopeMessagesCode1002, CloudIntegrationEditResponseEnvelopeMessagesCode1003, CloudIntegrationEditResponseEnvelopeMessagesCode1004, CloudIntegrationEditResponseEnvelopeMessagesCode1005, CloudIntegrationEditResponseEnvelopeMessagesCode1006, CloudIntegrationEditResponseEnvelopeMessagesCode1007, CloudIntegrationEditResponseEnvelopeMessagesCode1008, CloudIntegrationEditResponseEnvelopeMessagesCode1009, CloudIntegrationEditResponseEnvelopeMessagesCode1010, CloudIntegrationEditResponseEnvelopeMessagesCode1011, CloudIntegrationEditResponseEnvelopeMessagesCode1012, CloudIntegrationEditResponseEnvelopeMessagesCode1013, CloudIntegrationEditResponseEnvelopeMessagesCode1014, CloudIntegrationEditResponseEnvelopeMessagesCode1015, CloudIntegrationEditResponseEnvelopeMessagesCode1016, CloudIntegrationEditResponseEnvelopeMessagesCode1017, CloudIntegrationEditResponseEnvelopeMessagesCode2001, CloudIntegrationEditResponseEnvelopeMessagesCode2002, CloudIntegrationEditResponseEnvelopeMessagesCode2003, CloudIntegrationEditResponseEnvelopeMessagesCode2004, CloudIntegrationEditResponseEnvelopeMessagesCode2005, CloudIntegrationEditResponseEnvelopeMessagesCode2006, CloudIntegrationEditResponseEnvelopeMessagesCode2007, CloudIntegrationEditResponseEnvelopeMessagesCode2008, CloudIntegrationEditResponseEnvelopeMessagesCode2009, CloudIntegrationEditResponseEnvelopeMessagesCode2010, CloudIntegrationEditResponseEnvelopeMessagesCode2011, CloudIntegrationEditResponseEnvelopeMessagesCode2012, CloudIntegrationEditResponseEnvelopeMessagesCode2013, CloudIntegrationEditResponseEnvelopeMessagesCode2014, CloudIntegrationEditResponseEnvelopeMessagesCode2015, CloudIntegrationEditResponseEnvelopeMessagesCode2016, CloudIntegrationEditResponseEnvelopeMessagesCode2017, CloudIntegrationEditResponseEnvelopeMessagesCode2018, CloudIntegrationEditResponseEnvelopeMessagesCode2019, CloudIntegrationEditResponseEnvelopeMessagesCode2020, CloudIntegrationEditResponseEnvelopeMessagesCode2021, CloudIntegrationEditResponseEnvelopeMessagesCode2022, CloudIntegrationEditResponseEnvelopeMessagesCode3001, CloudIntegrationEditResponseEnvelopeMessagesCode3002, CloudIntegrationEditResponseEnvelopeMessagesCode3003, CloudIntegrationEditResponseEnvelopeMessagesCode3004, CloudIntegrationEditResponseEnvelopeMessagesCode3005, CloudIntegrationEditResponseEnvelopeMessagesCode3006, CloudIntegrationEditResponseEnvelopeMessagesCode3007, CloudIntegrationEditResponseEnvelopeMessagesCode4001, CloudIntegrationEditResponseEnvelopeMessagesCode4002, CloudIntegrationEditResponseEnvelopeMessagesCode4003, CloudIntegrationEditResponseEnvelopeMessagesCode4004, CloudIntegrationEditResponseEnvelopeMessagesCode4005, CloudIntegrationEditResponseEnvelopeMessagesCode4006, CloudIntegrationEditResponseEnvelopeMessagesCode4007, CloudIntegrationEditResponseEnvelopeMessagesCode4008, CloudIntegrationEditResponseEnvelopeMessagesCode4009, CloudIntegrationEditResponseEnvelopeMessagesCode4010, CloudIntegrationEditResponseEnvelopeMessagesCode4011, CloudIntegrationEditResponseEnvelopeMessagesCode4012, CloudIntegrationEditResponseEnvelopeMessagesCode4013, CloudIntegrationEditResponseEnvelopeMessagesCode4014, CloudIntegrationEditResponseEnvelopeMessagesCode4015, CloudIntegrationEditResponseEnvelopeMessagesCode4016, CloudIntegrationEditResponseEnvelopeMessagesCode4017, CloudIntegrationEditResponseEnvelopeMessagesCode4018, CloudIntegrationEditResponseEnvelopeMessagesCode4019, CloudIntegrationEditResponseEnvelopeMessagesCode4020, CloudIntegrationEditResponseEnvelopeMessagesCode4021, CloudIntegrationEditResponseEnvelopeMessagesCode4022, CloudIntegrationEditResponseEnvelopeMessagesCode4023, CloudIntegrationEditResponseEnvelopeMessagesCode5001, CloudIntegrationEditResponseEnvelopeMessagesCode5002, CloudIntegrationEditResponseEnvelopeMessagesCode5003, CloudIntegrationEditResponseEnvelopeMessagesCode5004, CloudIntegrationEditResponseEnvelopeMessagesCode102000, CloudIntegrationEditResponseEnvelopeMessagesCode102001, CloudIntegrationEditResponseEnvelopeMessagesCode102002, CloudIntegrationEditResponseEnvelopeMessagesCode102003, CloudIntegrationEditResponseEnvelopeMessagesCode102004, CloudIntegrationEditResponseEnvelopeMessagesCode102005, CloudIntegrationEditResponseEnvelopeMessagesCode102006, CloudIntegrationEditResponseEnvelopeMessagesCode102007, CloudIntegrationEditResponseEnvelopeMessagesCode102008, CloudIntegrationEditResponseEnvelopeMessagesCode102009, CloudIntegrationEditResponseEnvelopeMessagesCode102010, CloudIntegrationEditResponseEnvelopeMessagesCode102011, CloudIntegrationEditResponseEnvelopeMessagesCode102012, CloudIntegrationEditResponseEnvelopeMessagesCode102013, CloudIntegrationEditResponseEnvelopeMessagesCode102014, CloudIntegrationEditResponseEnvelopeMessagesCode102015, CloudIntegrationEditResponseEnvelopeMessagesCode102016, CloudIntegrationEditResponseEnvelopeMessagesCode102017, CloudIntegrationEditResponseEnvelopeMessagesCode102018, CloudIntegrationEditResponseEnvelopeMessagesCode102019, CloudIntegrationEditResponseEnvelopeMessagesCode102020, CloudIntegrationEditResponseEnvelopeMessagesCode102021, CloudIntegrationEditResponseEnvelopeMessagesCode102022, CloudIntegrationEditResponseEnvelopeMessagesCode102023, CloudIntegrationEditResponseEnvelopeMessagesCode102024, CloudIntegrationEditResponseEnvelopeMessagesCode102025, CloudIntegrationEditResponseEnvelopeMessagesCode102026, CloudIntegrationEditResponseEnvelopeMessagesCode102027, CloudIntegrationEditResponseEnvelopeMessagesCode102028, CloudIntegrationEditResponseEnvelopeMessagesCode102029, CloudIntegrationEditResponseEnvelopeMessagesCode102030, CloudIntegrationEditResponseEnvelopeMessagesCode102031, CloudIntegrationEditResponseEnvelopeMessagesCode102032, CloudIntegrationEditResponseEnvelopeMessagesCode102033, CloudIntegrationEditResponseEnvelopeMessagesCode102034, CloudIntegrationEditResponseEnvelopeMessagesCode102035, CloudIntegrationEditResponseEnvelopeMessagesCode102036, CloudIntegrationEditResponseEnvelopeMessagesCode102037, CloudIntegrationEditResponseEnvelopeMessagesCode102038, CloudIntegrationEditResponseEnvelopeMessagesCode102039, CloudIntegrationEditResponseEnvelopeMessagesCode102040, CloudIntegrationEditResponseEnvelopeMessagesCode102041, CloudIntegrationEditResponseEnvelopeMessagesCode102042, CloudIntegrationEditResponseEnvelopeMessagesCode102043, CloudIntegrationEditResponseEnvelopeMessagesCode102044, CloudIntegrationEditResponseEnvelopeMessagesCode102045, CloudIntegrationEditResponseEnvelopeMessagesCode102046, CloudIntegrationEditResponseEnvelopeMessagesCode102047, CloudIntegrationEditResponseEnvelopeMessagesCode102048, CloudIntegrationEditResponseEnvelopeMessagesCode102049, CloudIntegrationEditResponseEnvelopeMessagesCode102050, CloudIntegrationEditResponseEnvelopeMessagesCode102051, CloudIntegrationEditResponseEnvelopeMessagesCode102052, CloudIntegrationEditResponseEnvelopeMessagesCode102053, CloudIntegrationEditResponseEnvelopeMessagesCode102054, CloudIntegrationEditResponseEnvelopeMessagesCode102055, CloudIntegrationEditResponseEnvelopeMessagesCode102056, CloudIntegrationEditResponseEnvelopeMessagesCode102057, CloudIntegrationEditResponseEnvelopeMessagesCode102058, CloudIntegrationEditResponseEnvelopeMessagesCode102059, CloudIntegrationEditResponseEnvelopeMessagesCode102060, CloudIntegrationEditResponseEnvelopeMessagesCode102061, CloudIntegrationEditResponseEnvelopeMessagesCode102062, CloudIntegrationEditResponseEnvelopeMessagesCode102063, CloudIntegrationEditResponseEnvelopeMessagesCode102064, CloudIntegrationEditResponseEnvelopeMessagesCode102065, CloudIntegrationEditResponseEnvelopeMessagesCode102066, CloudIntegrationEditResponseEnvelopeMessagesCode103001, CloudIntegrationEditResponseEnvelopeMessagesCode103002, CloudIntegrationEditResponseEnvelopeMessagesCode103003, CloudIntegrationEditResponseEnvelopeMessagesCode103004, CloudIntegrationEditResponseEnvelopeMessagesCode103005, CloudIntegrationEditResponseEnvelopeMessagesCode103006, CloudIntegrationEditResponseEnvelopeMessagesCode103007, CloudIntegrationEditResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationEditResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                               `json:"l10n_key"`
	LoggableError string                                               `json:"loggable_error"`
	TemplateData  interface{}                                          `json:"template_data"`
	TraceID       string                                               `json:"trace_id"`
	JSON          cloudIntegrationEditResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeMessagesMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationEditResponseEnvelopeMessagesMeta]
type cloudIntegrationEditResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationEditResponseEnvelopeMessagesSource struct {
	Parameter           string                                                 `json:"parameter"`
	ParameterValueIndex int64                                                  `json:"parameter_value_index"`
	Pointer             string                                                 `json:"pointer"`
	JSON                cloudIntegrationEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CloudIntegrationEditResponseEnvelopeMessagesSource]
type cloudIntegrationEditResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Status    param.Field[bool]   `query:"status"`
}

// URLQuery serializes [CloudIntegrationGetParams]'s query parameters as
// `url.Values`.
func (r CloudIntegrationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CloudIntegrationGetResponseEnvelope struct {
	Errors   []CloudIntegrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationGetResponse                   `json:"result,required"`
	Success  bool                                          `json:"success,required"`
	JSON     cloudIntegrationGetResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CloudIntegrationGetResponseEnvelope]
type cloudIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeErrors struct {
	Code             CloudIntegrationGetResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Meta             CloudIntegrationGetResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CloudIntegrationGetResponseEnvelopeErrors]
type cloudIntegrationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationGetResponseEnvelopeErrorsCode1001   CloudIntegrationGetResponseEnvelopeErrorsCode = 1001
	CloudIntegrationGetResponseEnvelopeErrorsCode1002   CloudIntegrationGetResponseEnvelopeErrorsCode = 1002
	CloudIntegrationGetResponseEnvelopeErrorsCode1003   CloudIntegrationGetResponseEnvelopeErrorsCode = 1003
	CloudIntegrationGetResponseEnvelopeErrorsCode1004   CloudIntegrationGetResponseEnvelopeErrorsCode = 1004
	CloudIntegrationGetResponseEnvelopeErrorsCode1005   CloudIntegrationGetResponseEnvelopeErrorsCode = 1005
	CloudIntegrationGetResponseEnvelopeErrorsCode1006   CloudIntegrationGetResponseEnvelopeErrorsCode = 1006
	CloudIntegrationGetResponseEnvelopeErrorsCode1007   CloudIntegrationGetResponseEnvelopeErrorsCode = 1007
	CloudIntegrationGetResponseEnvelopeErrorsCode1008   CloudIntegrationGetResponseEnvelopeErrorsCode = 1008
	CloudIntegrationGetResponseEnvelopeErrorsCode1009   CloudIntegrationGetResponseEnvelopeErrorsCode = 1009
	CloudIntegrationGetResponseEnvelopeErrorsCode1010   CloudIntegrationGetResponseEnvelopeErrorsCode = 1010
	CloudIntegrationGetResponseEnvelopeErrorsCode1011   CloudIntegrationGetResponseEnvelopeErrorsCode = 1011
	CloudIntegrationGetResponseEnvelopeErrorsCode1012   CloudIntegrationGetResponseEnvelopeErrorsCode = 1012
	CloudIntegrationGetResponseEnvelopeErrorsCode1013   CloudIntegrationGetResponseEnvelopeErrorsCode = 1013
	CloudIntegrationGetResponseEnvelopeErrorsCode1014   CloudIntegrationGetResponseEnvelopeErrorsCode = 1014
	CloudIntegrationGetResponseEnvelopeErrorsCode1015   CloudIntegrationGetResponseEnvelopeErrorsCode = 1015
	CloudIntegrationGetResponseEnvelopeErrorsCode1016   CloudIntegrationGetResponseEnvelopeErrorsCode = 1016
	CloudIntegrationGetResponseEnvelopeErrorsCode1017   CloudIntegrationGetResponseEnvelopeErrorsCode = 1017
	CloudIntegrationGetResponseEnvelopeErrorsCode2001   CloudIntegrationGetResponseEnvelopeErrorsCode = 2001
	CloudIntegrationGetResponseEnvelopeErrorsCode2002   CloudIntegrationGetResponseEnvelopeErrorsCode = 2002
	CloudIntegrationGetResponseEnvelopeErrorsCode2003   CloudIntegrationGetResponseEnvelopeErrorsCode = 2003
	CloudIntegrationGetResponseEnvelopeErrorsCode2004   CloudIntegrationGetResponseEnvelopeErrorsCode = 2004
	CloudIntegrationGetResponseEnvelopeErrorsCode2005   CloudIntegrationGetResponseEnvelopeErrorsCode = 2005
	CloudIntegrationGetResponseEnvelopeErrorsCode2006   CloudIntegrationGetResponseEnvelopeErrorsCode = 2006
	CloudIntegrationGetResponseEnvelopeErrorsCode2007   CloudIntegrationGetResponseEnvelopeErrorsCode = 2007
	CloudIntegrationGetResponseEnvelopeErrorsCode2008   CloudIntegrationGetResponseEnvelopeErrorsCode = 2008
	CloudIntegrationGetResponseEnvelopeErrorsCode2009   CloudIntegrationGetResponseEnvelopeErrorsCode = 2009
	CloudIntegrationGetResponseEnvelopeErrorsCode2010   CloudIntegrationGetResponseEnvelopeErrorsCode = 2010
	CloudIntegrationGetResponseEnvelopeErrorsCode2011   CloudIntegrationGetResponseEnvelopeErrorsCode = 2011
	CloudIntegrationGetResponseEnvelopeErrorsCode2012   CloudIntegrationGetResponseEnvelopeErrorsCode = 2012
	CloudIntegrationGetResponseEnvelopeErrorsCode2013   CloudIntegrationGetResponseEnvelopeErrorsCode = 2013
	CloudIntegrationGetResponseEnvelopeErrorsCode2014   CloudIntegrationGetResponseEnvelopeErrorsCode = 2014
	CloudIntegrationGetResponseEnvelopeErrorsCode2015   CloudIntegrationGetResponseEnvelopeErrorsCode = 2015
	CloudIntegrationGetResponseEnvelopeErrorsCode2016   CloudIntegrationGetResponseEnvelopeErrorsCode = 2016
	CloudIntegrationGetResponseEnvelopeErrorsCode2017   CloudIntegrationGetResponseEnvelopeErrorsCode = 2017
	CloudIntegrationGetResponseEnvelopeErrorsCode2018   CloudIntegrationGetResponseEnvelopeErrorsCode = 2018
	CloudIntegrationGetResponseEnvelopeErrorsCode2019   CloudIntegrationGetResponseEnvelopeErrorsCode = 2019
	CloudIntegrationGetResponseEnvelopeErrorsCode2020   CloudIntegrationGetResponseEnvelopeErrorsCode = 2020
	CloudIntegrationGetResponseEnvelopeErrorsCode2021   CloudIntegrationGetResponseEnvelopeErrorsCode = 2021
	CloudIntegrationGetResponseEnvelopeErrorsCode2022   CloudIntegrationGetResponseEnvelopeErrorsCode = 2022
	CloudIntegrationGetResponseEnvelopeErrorsCode3001   CloudIntegrationGetResponseEnvelopeErrorsCode = 3001
	CloudIntegrationGetResponseEnvelopeErrorsCode3002   CloudIntegrationGetResponseEnvelopeErrorsCode = 3002
	CloudIntegrationGetResponseEnvelopeErrorsCode3003   CloudIntegrationGetResponseEnvelopeErrorsCode = 3003
	CloudIntegrationGetResponseEnvelopeErrorsCode3004   CloudIntegrationGetResponseEnvelopeErrorsCode = 3004
	CloudIntegrationGetResponseEnvelopeErrorsCode3005   CloudIntegrationGetResponseEnvelopeErrorsCode = 3005
	CloudIntegrationGetResponseEnvelopeErrorsCode3006   CloudIntegrationGetResponseEnvelopeErrorsCode = 3006
	CloudIntegrationGetResponseEnvelopeErrorsCode3007   CloudIntegrationGetResponseEnvelopeErrorsCode = 3007
	CloudIntegrationGetResponseEnvelopeErrorsCode4001   CloudIntegrationGetResponseEnvelopeErrorsCode = 4001
	CloudIntegrationGetResponseEnvelopeErrorsCode4002   CloudIntegrationGetResponseEnvelopeErrorsCode = 4002
	CloudIntegrationGetResponseEnvelopeErrorsCode4003   CloudIntegrationGetResponseEnvelopeErrorsCode = 4003
	CloudIntegrationGetResponseEnvelopeErrorsCode4004   CloudIntegrationGetResponseEnvelopeErrorsCode = 4004
	CloudIntegrationGetResponseEnvelopeErrorsCode4005   CloudIntegrationGetResponseEnvelopeErrorsCode = 4005
	CloudIntegrationGetResponseEnvelopeErrorsCode4006   CloudIntegrationGetResponseEnvelopeErrorsCode = 4006
	CloudIntegrationGetResponseEnvelopeErrorsCode4007   CloudIntegrationGetResponseEnvelopeErrorsCode = 4007
	CloudIntegrationGetResponseEnvelopeErrorsCode4008   CloudIntegrationGetResponseEnvelopeErrorsCode = 4008
	CloudIntegrationGetResponseEnvelopeErrorsCode4009   CloudIntegrationGetResponseEnvelopeErrorsCode = 4009
	CloudIntegrationGetResponseEnvelopeErrorsCode4010   CloudIntegrationGetResponseEnvelopeErrorsCode = 4010
	CloudIntegrationGetResponseEnvelopeErrorsCode4011   CloudIntegrationGetResponseEnvelopeErrorsCode = 4011
	CloudIntegrationGetResponseEnvelopeErrorsCode4012   CloudIntegrationGetResponseEnvelopeErrorsCode = 4012
	CloudIntegrationGetResponseEnvelopeErrorsCode4013   CloudIntegrationGetResponseEnvelopeErrorsCode = 4013
	CloudIntegrationGetResponseEnvelopeErrorsCode4014   CloudIntegrationGetResponseEnvelopeErrorsCode = 4014
	CloudIntegrationGetResponseEnvelopeErrorsCode4015   CloudIntegrationGetResponseEnvelopeErrorsCode = 4015
	CloudIntegrationGetResponseEnvelopeErrorsCode4016   CloudIntegrationGetResponseEnvelopeErrorsCode = 4016
	CloudIntegrationGetResponseEnvelopeErrorsCode4017   CloudIntegrationGetResponseEnvelopeErrorsCode = 4017
	CloudIntegrationGetResponseEnvelopeErrorsCode4018   CloudIntegrationGetResponseEnvelopeErrorsCode = 4018
	CloudIntegrationGetResponseEnvelopeErrorsCode4019   CloudIntegrationGetResponseEnvelopeErrorsCode = 4019
	CloudIntegrationGetResponseEnvelopeErrorsCode4020   CloudIntegrationGetResponseEnvelopeErrorsCode = 4020
	CloudIntegrationGetResponseEnvelopeErrorsCode4021   CloudIntegrationGetResponseEnvelopeErrorsCode = 4021
	CloudIntegrationGetResponseEnvelopeErrorsCode4022   CloudIntegrationGetResponseEnvelopeErrorsCode = 4022
	CloudIntegrationGetResponseEnvelopeErrorsCode4023   CloudIntegrationGetResponseEnvelopeErrorsCode = 4023
	CloudIntegrationGetResponseEnvelopeErrorsCode5001   CloudIntegrationGetResponseEnvelopeErrorsCode = 5001
	CloudIntegrationGetResponseEnvelopeErrorsCode5002   CloudIntegrationGetResponseEnvelopeErrorsCode = 5002
	CloudIntegrationGetResponseEnvelopeErrorsCode5003   CloudIntegrationGetResponseEnvelopeErrorsCode = 5003
	CloudIntegrationGetResponseEnvelopeErrorsCode5004   CloudIntegrationGetResponseEnvelopeErrorsCode = 5004
	CloudIntegrationGetResponseEnvelopeErrorsCode102000 CloudIntegrationGetResponseEnvelopeErrorsCode = 102000
	CloudIntegrationGetResponseEnvelopeErrorsCode102001 CloudIntegrationGetResponseEnvelopeErrorsCode = 102001
	CloudIntegrationGetResponseEnvelopeErrorsCode102002 CloudIntegrationGetResponseEnvelopeErrorsCode = 102002
	CloudIntegrationGetResponseEnvelopeErrorsCode102003 CloudIntegrationGetResponseEnvelopeErrorsCode = 102003
	CloudIntegrationGetResponseEnvelopeErrorsCode102004 CloudIntegrationGetResponseEnvelopeErrorsCode = 102004
	CloudIntegrationGetResponseEnvelopeErrorsCode102005 CloudIntegrationGetResponseEnvelopeErrorsCode = 102005
	CloudIntegrationGetResponseEnvelopeErrorsCode102006 CloudIntegrationGetResponseEnvelopeErrorsCode = 102006
	CloudIntegrationGetResponseEnvelopeErrorsCode102007 CloudIntegrationGetResponseEnvelopeErrorsCode = 102007
	CloudIntegrationGetResponseEnvelopeErrorsCode102008 CloudIntegrationGetResponseEnvelopeErrorsCode = 102008
	CloudIntegrationGetResponseEnvelopeErrorsCode102009 CloudIntegrationGetResponseEnvelopeErrorsCode = 102009
	CloudIntegrationGetResponseEnvelopeErrorsCode102010 CloudIntegrationGetResponseEnvelopeErrorsCode = 102010
	CloudIntegrationGetResponseEnvelopeErrorsCode102011 CloudIntegrationGetResponseEnvelopeErrorsCode = 102011
	CloudIntegrationGetResponseEnvelopeErrorsCode102012 CloudIntegrationGetResponseEnvelopeErrorsCode = 102012
	CloudIntegrationGetResponseEnvelopeErrorsCode102013 CloudIntegrationGetResponseEnvelopeErrorsCode = 102013
	CloudIntegrationGetResponseEnvelopeErrorsCode102014 CloudIntegrationGetResponseEnvelopeErrorsCode = 102014
	CloudIntegrationGetResponseEnvelopeErrorsCode102015 CloudIntegrationGetResponseEnvelopeErrorsCode = 102015
	CloudIntegrationGetResponseEnvelopeErrorsCode102016 CloudIntegrationGetResponseEnvelopeErrorsCode = 102016
	CloudIntegrationGetResponseEnvelopeErrorsCode102017 CloudIntegrationGetResponseEnvelopeErrorsCode = 102017
	CloudIntegrationGetResponseEnvelopeErrorsCode102018 CloudIntegrationGetResponseEnvelopeErrorsCode = 102018
	CloudIntegrationGetResponseEnvelopeErrorsCode102019 CloudIntegrationGetResponseEnvelopeErrorsCode = 102019
	CloudIntegrationGetResponseEnvelopeErrorsCode102020 CloudIntegrationGetResponseEnvelopeErrorsCode = 102020
	CloudIntegrationGetResponseEnvelopeErrorsCode102021 CloudIntegrationGetResponseEnvelopeErrorsCode = 102021
	CloudIntegrationGetResponseEnvelopeErrorsCode102022 CloudIntegrationGetResponseEnvelopeErrorsCode = 102022
	CloudIntegrationGetResponseEnvelopeErrorsCode102023 CloudIntegrationGetResponseEnvelopeErrorsCode = 102023
	CloudIntegrationGetResponseEnvelopeErrorsCode102024 CloudIntegrationGetResponseEnvelopeErrorsCode = 102024
	CloudIntegrationGetResponseEnvelopeErrorsCode102025 CloudIntegrationGetResponseEnvelopeErrorsCode = 102025
	CloudIntegrationGetResponseEnvelopeErrorsCode102026 CloudIntegrationGetResponseEnvelopeErrorsCode = 102026
	CloudIntegrationGetResponseEnvelopeErrorsCode102027 CloudIntegrationGetResponseEnvelopeErrorsCode = 102027
	CloudIntegrationGetResponseEnvelopeErrorsCode102028 CloudIntegrationGetResponseEnvelopeErrorsCode = 102028
	CloudIntegrationGetResponseEnvelopeErrorsCode102029 CloudIntegrationGetResponseEnvelopeErrorsCode = 102029
	CloudIntegrationGetResponseEnvelopeErrorsCode102030 CloudIntegrationGetResponseEnvelopeErrorsCode = 102030
	CloudIntegrationGetResponseEnvelopeErrorsCode102031 CloudIntegrationGetResponseEnvelopeErrorsCode = 102031
	CloudIntegrationGetResponseEnvelopeErrorsCode102032 CloudIntegrationGetResponseEnvelopeErrorsCode = 102032
	CloudIntegrationGetResponseEnvelopeErrorsCode102033 CloudIntegrationGetResponseEnvelopeErrorsCode = 102033
	CloudIntegrationGetResponseEnvelopeErrorsCode102034 CloudIntegrationGetResponseEnvelopeErrorsCode = 102034
	CloudIntegrationGetResponseEnvelopeErrorsCode102035 CloudIntegrationGetResponseEnvelopeErrorsCode = 102035
	CloudIntegrationGetResponseEnvelopeErrorsCode102036 CloudIntegrationGetResponseEnvelopeErrorsCode = 102036
	CloudIntegrationGetResponseEnvelopeErrorsCode102037 CloudIntegrationGetResponseEnvelopeErrorsCode = 102037
	CloudIntegrationGetResponseEnvelopeErrorsCode102038 CloudIntegrationGetResponseEnvelopeErrorsCode = 102038
	CloudIntegrationGetResponseEnvelopeErrorsCode102039 CloudIntegrationGetResponseEnvelopeErrorsCode = 102039
	CloudIntegrationGetResponseEnvelopeErrorsCode102040 CloudIntegrationGetResponseEnvelopeErrorsCode = 102040
	CloudIntegrationGetResponseEnvelopeErrorsCode102041 CloudIntegrationGetResponseEnvelopeErrorsCode = 102041
	CloudIntegrationGetResponseEnvelopeErrorsCode102042 CloudIntegrationGetResponseEnvelopeErrorsCode = 102042
	CloudIntegrationGetResponseEnvelopeErrorsCode102043 CloudIntegrationGetResponseEnvelopeErrorsCode = 102043
	CloudIntegrationGetResponseEnvelopeErrorsCode102044 CloudIntegrationGetResponseEnvelopeErrorsCode = 102044
	CloudIntegrationGetResponseEnvelopeErrorsCode102045 CloudIntegrationGetResponseEnvelopeErrorsCode = 102045
	CloudIntegrationGetResponseEnvelopeErrorsCode102046 CloudIntegrationGetResponseEnvelopeErrorsCode = 102046
	CloudIntegrationGetResponseEnvelopeErrorsCode102047 CloudIntegrationGetResponseEnvelopeErrorsCode = 102047
	CloudIntegrationGetResponseEnvelopeErrorsCode102048 CloudIntegrationGetResponseEnvelopeErrorsCode = 102048
	CloudIntegrationGetResponseEnvelopeErrorsCode102049 CloudIntegrationGetResponseEnvelopeErrorsCode = 102049
	CloudIntegrationGetResponseEnvelopeErrorsCode102050 CloudIntegrationGetResponseEnvelopeErrorsCode = 102050
	CloudIntegrationGetResponseEnvelopeErrorsCode102051 CloudIntegrationGetResponseEnvelopeErrorsCode = 102051
	CloudIntegrationGetResponseEnvelopeErrorsCode102052 CloudIntegrationGetResponseEnvelopeErrorsCode = 102052
	CloudIntegrationGetResponseEnvelopeErrorsCode102053 CloudIntegrationGetResponseEnvelopeErrorsCode = 102053
	CloudIntegrationGetResponseEnvelopeErrorsCode102054 CloudIntegrationGetResponseEnvelopeErrorsCode = 102054
	CloudIntegrationGetResponseEnvelopeErrorsCode102055 CloudIntegrationGetResponseEnvelopeErrorsCode = 102055
	CloudIntegrationGetResponseEnvelopeErrorsCode102056 CloudIntegrationGetResponseEnvelopeErrorsCode = 102056
	CloudIntegrationGetResponseEnvelopeErrorsCode102057 CloudIntegrationGetResponseEnvelopeErrorsCode = 102057
	CloudIntegrationGetResponseEnvelopeErrorsCode102058 CloudIntegrationGetResponseEnvelopeErrorsCode = 102058
	CloudIntegrationGetResponseEnvelopeErrorsCode102059 CloudIntegrationGetResponseEnvelopeErrorsCode = 102059
	CloudIntegrationGetResponseEnvelopeErrorsCode102060 CloudIntegrationGetResponseEnvelopeErrorsCode = 102060
	CloudIntegrationGetResponseEnvelopeErrorsCode102061 CloudIntegrationGetResponseEnvelopeErrorsCode = 102061
	CloudIntegrationGetResponseEnvelopeErrorsCode102062 CloudIntegrationGetResponseEnvelopeErrorsCode = 102062
	CloudIntegrationGetResponseEnvelopeErrorsCode102063 CloudIntegrationGetResponseEnvelopeErrorsCode = 102063
	CloudIntegrationGetResponseEnvelopeErrorsCode102064 CloudIntegrationGetResponseEnvelopeErrorsCode = 102064
	CloudIntegrationGetResponseEnvelopeErrorsCode102065 CloudIntegrationGetResponseEnvelopeErrorsCode = 102065
	CloudIntegrationGetResponseEnvelopeErrorsCode102066 CloudIntegrationGetResponseEnvelopeErrorsCode = 102066
	CloudIntegrationGetResponseEnvelopeErrorsCode103001 CloudIntegrationGetResponseEnvelopeErrorsCode = 103001
	CloudIntegrationGetResponseEnvelopeErrorsCode103002 CloudIntegrationGetResponseEnvelopeErrorsCode = 103002
	CloudIntegrationGetResponseEnvelopeErrorsCode103003 CloudIntegrationGetResponseEnvelopeErrorsCode = 103003
	CloudIntegrationGetResponseEnvelopeErrorsCode103004 CloudIntegrationGetResponseEnvelopeErrorsCode = 103004
	CloudIntegrationGetResponseEnvelopeErrorsCode103005 CloudIntegrationGetResponseEnvelopeErrorsCode = 103005
	CloudIntegrationGetResponseEnvelopeErrorsCode103006 CloudIntegrationGetResponseEnvelopeErrorsCode = 103006
	CloudIntegrationGetResponseEnvelopeErrorsCode103007 CloudIntegrationGetResponseEnvelopeErrorsCode = 103007
	CloudIntegrationGetResponseEnvelopeErrorsCode103008 CloudIntegrationGetResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationGetResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseEnvelopeErrorsCode1001, CloudIntegrationGetResponseEnvelopeErrorsCode1002, CloudIntegrationGetResponseEnvelopeErrorsCode1003, CloudIntegrationGetResponseEnvelopeErrorsCode1004, CloudIntegrationGetResponseEnvelopeErrorsCode1005, CloudIntegrationGetResponseEnvelopeErrorsCode1006, CloudIntegrationGetResponseEnvelopeErrorsCode1007, CloudIntegrationGetResponseEnvelopeErrorsCode1008, CloudIntegrationGetResponseEnvelopeErrorsCode1009, CloudIntegrationGetResponseEnvelopeErrorsCode1010, CloudIntegrationGetResponseEnvelopeErrorsCode1011, CloudIntegrationGetResponseEnvelopeErrorsCode1012, CloudIntegrationGetResponseEnvelopeErrorsCode1013, CloudIntegrationGetResponseEnvelopeErrorsCode1014, CloudIntegrationGetResponseEnvelopeErrorsCode1015, CloudIntegrationGetResponseEnvelopeErrorsCode1016, CloudIntegrationGetResponseEnvelopeErrorsCode1017, CloudIntegrationGetResponseEnvelopeErrorsCode2001, CloudIntegrationGetResponseEnvelopeErrorsCode2002, CloudIntegrationGetResponseEnvelopeErrorsCode2003, CloudIntegrationGetResponseEnvelopeErrorsCode2004, CloudIntegrationGetResponseEnvelopeErrorsCode2005, CloudIntegrationGetResponseEnvelopeErrorsCode2006, CloudIntegrationGetResponseEnvelopeErrorsCode2007, CloudIntegrationGetResponseEnvelopeErrorsCode2008, CloudIntegrationGetResponseEnvelopeErrorsCode2009, CloudIntegrationGetResponseEnvelopeErrorsCode2010, CloudIntegrationGetResponseEnvelopeErrorsCode2011, CloudIntegrationGetResponseEnvelopeErrorsCode2012, CloudIntegrationGetResponseEnvelopeErrorsCode2013, CloudIntegrationGetResponseEnvelopeErrorsCode2014, CloudIntegrationGetResponseEnvelopeErrorsCode2015, CloudIntegrationGetResponseEnvelopeErrorsCode2016, CloudIntegrationGetResponseEnvelopeErrorsCode2017, CloudIntegrationGetResponseEnvelopeErrorsCode2018, CloudIntegrationGetResponseEnvelopeErrorsCode2019, CloudIntegrationGetResponseEnvelopeErrorsCode2020, CloudIntegrationGetResponseEnvelopeErrorsCode2021, CloudIntegrationGetResponseEnvelopeErrorsCode2022, CloudIntegrationGetResponseEnvelopeErrorsCode3001, CloudIntegrationGetResponseEnvelopeErrorsCode3002, CloudIntegrationGetResponseEnvelopeErrorsCode3003, CloudIntegrationGetResponseEnvelopeErrorsCode3004, CloudIntegrationGetResponseEnvelopeErrorsCode3005, CloudIntegrationGetResponseEnvelopeErrorsCode3006, CloudIntegrationGetResponseEnvelopeErrorsCode3007, CloudIntegrationGetResponseEnvelopeErrorsCode4001, CloudIntegrationGetResponseEnvelopeErrorsCode4002, CloudIntegrationGetResponseEnvelopeErrorsCode4003, CloudIntegrationGetResponseEnvelopeErrorsCode4004, CloudIntegrationGetResponseEnvelopeErrorsCode4005, CloudIntegrationGetResponseEnvelopeErrorsCode4006, CloudIntegrationGetResponseEnvelopeErrorsCode4007, CloudIntegrationGetResponseEnvelopeErrorsCode4008, CloudIntegrationGetResponseEnvelopeErrorsCode4009, CloudIntegrationGetResponseEnvelopeErrorsCode4010, CloudIntegrationGetResponseEnvelopeErrorsCode4011, CloudIntegrationGetResponseEnvelopeErrorsCode4012, CloudIntegrationGetResponseEnvelopeErrorsCode4013, CloudIntegrationGetResponseEnvelopeErrorsCode4014, CloudIntegrationGetResponseEnvelopeErrorsCode4015, CloudIntegrationGetResponseEnvelopeErrorsCode4016, CloudIntegrationGetResponseEnvelopeErrorsCode4017, CloudIntegrationGetResponseEnvelopeErrorsCode4018, CloudIntegrationGetResponseEnvelopeErrorsCode4019, CloudIntegrationGetResponseEnvelopeErrorsCode4020, CloudIntegrationGetResponseEnvelopeErrorsCode4021, CloudIntegrationGetResponseEnvelopeErrorsCode4022, CloudIntegrationGetResponseEnvelopeErrorsCode4023, CloudIntegrationGetResponseEnvelopeErrorsCode5001, CloudIntegrationGetResponseEnvelopeErrorsCode5002, CloudIntegrationGetResponseEnvelopeErrorsCode5003, CloudIntegrationGetResponseEnvelopeErrorsCode5004, CloudIntegrationGetResponseEnvelopeErrorsCode102000, CloudIntegrationGetResponseEnvelopeErrorsCode102001, CloudIntegrationGetResponseEnvelopeErrorsCode102002, CloudIntegrationGetResponseEnvelopeErrorsCode102003, CloudIntegrationGetResponseEnvelopeErrorsCode102004, CloudIntegrationGetResponseEnvelopeErrorsCode102005, CloudIntegrationGetResponseEnvelopeErrorsCode102006, CloudIntegrationGetResponseEnvelopeErrorsCode102007, CloudIntegrationGetResponseEnvelopeErrorsCode102008, CloudIntegrationGetResponseEnvelopeErrorsCode102009, CloudIntegrationGetResponseEnvelopeErrorsCode102010, CloudIntegrationGetResponseEnvelopeErrorsCode102011, CloudIntegrationGetResponseEnvelopeErrorsCode102012, CloudIntegrationGetResponseEnvelopeErrorsCode102013, CloudIntegrationGetResponseEnvelopeErrorsCode102014, CloudIntegrationGetResponseEnvelopeErrorsCode102015, CloudIntegrationGetResponseEnvelopeErrorsCode102016, CloudIntegrationGetResponseEnvelopeErrorsCode102017, CloudIntegrationGetResponseEnvelopeErrorsCode102018, CloudIntegrationGetResponseEnvelopeErrorsCode102019, CloudIntegrationGetResponseEnvelopeErrorsCode102020, CloudIntegrationGetResponseEnvelopeErrorsCode102021, CloudIntegrationGetResponseEnvelopeErrorsCode102022, CloudIntegrationGetResponseEnvelopeErrorsCode102023, CloudIntegrationGetResponseEnvelopeErrorsCode102024, CloudIntegrationGetResponseEnvelopeErrorsCode102025, CloudIntegrationGetResponseEnvelopeErrorsCode102026, CloudIntegrationGetResponseEnvelopeErrorsCode102027, CloudIntegrationGetResponseEnvelopeErrorsCode102028, CloudIntegrationGetResponseEnvelopeErrorsCode102029, CloudIntegrationGetResponseEnvelopeErrorsCode102030, CloudIntegrationGetResponseEnvelopeErrorsCode102031, CloudIntegrationGetResponseEnvelopeErrorsCode102032, CloudIntegrationGetResponseEnvelopeErrorsCode102033, CloudIntegrationGetResponseEnvelopeErrorsCode102034, CloudIntegrationGetResponseEnvelopeErrorsCode102035, CloudIntegrationGetResponseEnvelopeErrorsCode102036, CloudIntegrationGetResponseEnvelopeErrorsCode102037, CloudIntegrationGetResponseEnvelopeErrorsCode102038, CloudIntegrationGetResponseEnvelopeErrorsCode102039, CloudIntegrationGetResponseEnvelopeErrorsCode102040, CloudIntegrationGetResponseEnvelopeErrorsCode102041, CloudIntegrationGetResponseEnvelopeErrorsCode102042, CloudIntegrationGetResponseEnvelopeErrorsCode102043, CloudIntegrationGetResponseEnvelopeErrorsCode102044, CloudIntegrationGetResponseEnvelopeErrorsCode102045, CloudIntegrationGetResponseEnvelopeErrorsCode102046, CloudIntegrationGetResponseEnvelopeErrorsCode102047, CloudIntegrationGetResponseEnvelopeErrorsCode102048, CloudIntegrationGetResponseEnvelopeErrorsCode102049, CloudIntegrationGetResponseEnvelopeErrorsCode102050, CloudIntegrationGetResponseEnvelopeErrorsCode102051, CloudIntegrationGetResponseEnvelopeErrorsCode102052, CloudIntegrationGetResponseEnvelopeErrorsCode102053, CloudIntegrationGetResponseEnvelopeErrorsCode102054, CloudIntegrationGetResponseEnvelopeErrorsCode102055, CloudIntegrationGetResponseEnvelopeErrorsCode102056, CloudIntegrationGetResponseEnvelopeErrorsCode102057, CloudIntegrationGetResponseEnvelopeErrorsCode102058, CloudIntegrationGetResponseEnvelopeErrorsCode102059, CloudIntegrationGetResponseEnvelopeErrorsCode102060, CloudIntegrationGetResponseEnvelopeErrorsCode102061, CloudIntegrationGetResponseEnvelopeErrorsCode102062, CloudIntegrationGetResponseEnvelopeErrorsCode102063, CloudIntegrationGetResponseEnvelopeErrorsCode102064, CloudIntegrationGetResponseEnvelopeErrorsCode102065, CloudIntegrationGetResponseEnvelopeErrorsCode102066, CloudIntegrationGetResponseEnvelopeErrorsCode103001, CloudIntegrationGetResponseEnvelopeErrorsCode103002, CloudIntegrationGetResponseEnvelopeErrorsCode103003, CloudIntegrationGetResponseEnvelopeErrorsCode103004, CloudIntegrationGetResponseEnvelopeErrorsCode103005, CloudIntegrationGetResponseEnvelopeErrorsCode103006, CloudIntegrationGetResponseEnvelopeErrorsCode103007, CloudIntegrationGetResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationGetResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                            `json:"l10n_key"`
	LoggableError string                                            `json:"loggable_error"`
	TemplateData  interface{}                                       `json:"template_data"`
	TraceID       string                                            `json:"trace_id"`
	JSON          cloudIntegrationGetResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeErrorsMetaJSON contains the JSON metadata for
// the struct [CloudIntegrationGetResponseEnvelopeErrorsMeta]
type cloudIntegrationGetResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeErrorsSource struct {
	Parameter           string                                              `json:"parameter"`
	ParameterValueIndex int64                                               `json:"parameter_value_index"`
	Pointer             string                                              `json:"pointer"`
	JSON                cloudIntegrationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationGetResponseEnvelopeErrorsSource]
type cloudIntegrationGetResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeMessages struct {
	Code             CloudIntegrationGetResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Meta             CloudIntegrationGetResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CloudIntegrationGetResponseEnvelopeMessages]
type cloudIntegrationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationGetResponseEnvelopeMessagesCode1001   CloudIntegrationGetResponseEnvelopeMessagesCode = 1001
	CloudIntegrationGetResponseEnvelopeMessagesCode1002   CloudIntegrationGetResponseEnvelopeMessagesCode = 1002
	CloudIntegrationGetResponseEnvelopeMessagesCode1003   CloudIntegrationGetResponseEnvelopeMessagesCode = 1003
	CloudIntegrationGetResponseEnvelopeMessagesCode1004   CloudIntegrationGetResponseEnvelopeMessagesCode = 1004
	CloudIntegrationGetResponseEnvelopeMessagesCode1005   CloudIntegrationGetResponseEnvelopeMessagesCode = 1005
	CloudIntegrationGetResponseEnvelopeMessagesCode1006   CloudIntegrationGetResponseEnvelopeMessagesCode = 1006
	CloudIntegrationGetResponseEnvelopeMessagesCode1007   CloudIntegrationGetResponseEnvelopeMessagesCode = 1007
	CloudIntegrationGetResponseEnvelopeMessagesCode1008   CloudIntegrationGetResponseEnvelopeMessagesCode = 1008
	CloudIntegrationGetResponseEnvelopeMessagesCode1009   CloudIntegrationGetResponseEnvelopeMessagesCode = 1009
	CloudIntegrationGetResponseEnvelopeMessagesCode1010   CloudIntegrationGetResponseEnvelopeMessagesCode = 1010
	CloudIntegrationGetResponseEnvelopeMessagesCode1011   CloudIntegrationGetResponseEnvelopeMessagesCode = 1011
	CloudIntegrationGetResponseEnvelopeMessagesCode1012   CloudIntegrationGetResponseEnvelopeMessagesCode = 1012
	CloudIntegrationGetResponseEnvelopeMessagesCode1013   CloudIntegrationGetResponseEnvelopeMessagesCode = 1013
	CloudIntegrationGetResponseEnvelopeMessagesCode1014   CloudIntegrationGetResponseEnvelopeMessagesCode = 1014
	CloudIntegrationGetResponseEnvelopeMessagesCode1015   CloudIntegrationGetResponseEnvelopeMessagesCode = 1015
	CloudIntegrationGetResponseEnvelopeMessagesCode1016   CloudIntegrationGetResponseEnvelopeMessagesCode = 1016
	CloudIntegrationGetResponseEnvelopeMessagesCode1017   CloudIntegrationGetResponseEnvelopeMessagesCode = 1017
	CloudIntegrationGetResponseEnvelopeMessagesCode2001   CloudIntegrationGetResponseEnvelopeMessagesCode = 2001
	CloudIntegrationGetResponseEnvelopeMessagesCode2002   CloudIntegrationGetResponseEnvelopeMessagesCode = 2002
	CloudIntegrationGetResponseEnvelopeMessagesCode2003   CloudIntegrationGetResponseEnvelopeMessagesCode = 2003
	CloudIntegrationGetResponseEnvelopeMessagesCode2004   CloudIntegrationGetResponseEnvelopeMessagesCode = 2004
	CloudIntegrationGetResponseEnvelopeMessagesCode2005   CloudIntegrationGetResponseEnvelopeMessagesCode = 2005
	CloudIntegrationGetResponseEnvelopeMessagesCode2006   CloudIntegrationGetResponseEnvelopeMessagesCode = 2006
	CloudIntegrationGetResponseEnvelopeMessagesCode2007   CloudIntegrationGetResponseEnvelopeMessagesCode = 2007
	CloudIntegrationGetResponseEnvelopeMessagesCode2008   CloudIntegrationGetResponseEnvelopeMessagesCode = 2008
	CloudIntegrationGetResponseEnvelopeMessagesCode2009   CloudIntegrationGetResponseEnvelopeMessagesCode = 2009
	CloudIntegrationGetResponseEnvelopeMessagesCode2010   CloudIntegrationGetResponseEnvelopeMessagesCode = 2010
	CloudIntegrationGetResponseEnvelopeMessagesCode2011   CloudIntegrationGetResponseEnvelopeMessagesCode = 2011
	CloudIntegrationGetResponseEnvelopeMessagesCode2012   CloudIntegrationGetResponseEnvelopeMessagesCode = 2012
	CloudIntegrationGetResponseEnvelopeMessagesCode2013   CloudIntegrationGetResponseEnvelopeMessagesCode = 2013
	CloudIntegrationGetResponseEnvelopeMessagesCode2014   CloudIntegrationGetResponseEnvelopeMessagesCode = 2014
	CloudIntegrationGetResponseEnvelopeMessagesCode2015   CloudIntegrationGetResponseEnvelopeMessagesCode = 2015
	CloudIntegrationGetResponseEnvelopeMessagesCode2016   CloudIntegrationGetResponseEnvelopeMessagesCode = 2016
	CloudIntegrationGetResponseEnvelopeMessagesCode2017   CloudIntegrationGetResponseEnvelopeMessagesCode = 2017
	CloudIntegrationGetResponseEnvelopeMessagesCode2018   CloudIntegrationGetResponseEnvelopeMessagesCode = 2018
	CloudIntegrationGetResponseEnvelopeMessagesCode2019   CloudIntegrationGetResponseEnvelopeMessagesCode = 2019
	CloudIntegrationGetResponseEnvelopeMessagesCode2020   CloudIntegrationGetResponseEnvelopeMessagesCode = 2020
	CloudIntegrationGetResponseEnvelopeMessagesCode2021   CloudIntegrationGetResponseEnvelopeMessagesCode = 2021
	CloudIntegrationGetResponseEnvelopeMessagesCode2022   CloudIntegrationGetResponseEnvelopeMessagesCode = 2022
	CloudIntegrationGetResponseEnvelopeMessagesCode3001   CloudIntegrationGetResponseEnvelopeMessagesCode = 3001
	CloudIntegrationGetResponseEnvelopeMessagesCode3002   CloudIntegrationGetResponseEnvelopeMessagesCode = 3002
	CloudIntegrationGetResponseEnvelopeMessagesCode3003   CloudIntegrationGetResponseEnvelopeMessagesCode = 3003
	CloudIntegrationGetResponseEnvelopeMessagesCode3004   CloudIntegrationGetResponseEnvelopeMessagesCode = 3004
	CloudIntegrationGetResponseEnvelopeMessagesCode3005   CloudIntegrationGetResponseEnvelopeMessagesCode = 3005
	CloudIntegrationGetResponseEnvelopeMessagesCode3006   CloudIntegrationGetResponseEnvelopeMessagesCode = 3006
	CloudIntegrationGetResponseEnvelopeMessagesCode3007   CloudIntegrationGetResponseEnvelopeMessagesCode = 3007
	CloudIntegrationGetResponseEnvelopeMessagesCode4001   CloudIntegrationGetResponseEnvelopeMessagesCode = 4001
	CloudIntegrationGetResponseEnvelopeMessagesCode4002   CloudIntegrationGetResponseEnvelopeMessagesCode = 4002
	CloudIntegrationGetResponseEnvelopeMessagesCode4003   CloudIntegrationGetResponseEnvelopeMessagesCode = 4003
	CloudIntegrationGetResponseEnvelopeMessagesCode4004   CloudIntegrationGetResponseEnvelopeMessagesCode = 4004
	CloudIntegrationGetResponseEnvelopeMessagesCode4005   CloudIntegrationGetResponseEnvelopeMessagesCode = 4005
	CloudIntegrationGetResponseEnvelopeMessagesCode4006   CloudIntegrationGetResponseEnvelopeMessagesCode = 4006
	CloudIntegrationGetResponseEnvelopeMessagesCode4007   CloudIntegrationGetResponseEnvelopeMessagesCode = 4007
	CloudIntegrationGetResponseEnvelopeMessagesCode4008   CloudIntegrationGetResponseEnvelopeMessagesCode = 4008
	CloudIntegrationGetResponseEnvelopeMessagesCode4009   CloudIntegrationGetResponseEnvelopeMessagesCode = 4009
	CloudIntegrationGetResponseEnvelopeMessagesCode4010   CloudIntegrationGetResponseEnvelopeMessagesCode = 4010
	CloudIntegrationGetResponseEnvelopeMessagesCode4011   CloudIntegrationGetResponseEnvelopeMessagesCode = 4011
	CloudIntegrationGetResponseEnvelopeMessagesCode4012   CloudIntegrationGetResponseEnvelopeMessagesCode = 4012
	CloudIntegrationGetResponseEnvelopeMessagesCode4013   CloudIntegrationGetResponseEnvelopeMessagesCode = 4013
	CloudIntegrationGetResponseEnvelopeMessagesCode4014   CloudIntegrationGetResponseEnvelopeMessagesCode = 4014
	CloudIntegrationGetResponseEnvelopeMessagesCode4015   CloudIntegrationGetResponseEnvelopeMessagesCode = 4015
	CloudIntegrationGetResponseEnvelopeMessagesCode4016   CloudIntegrationGetResponseEnvelopeMessagesCode = 4016
	CloudIntegrationGetResponseEnvelopeMessagesCode4017   CloudIntegrationGetResponseEnvelopeMessagesCode = 4017
	CloudIntegrationGetResponseEnvelopeMessagesCode4018   CloudIntegrationGetResponseEnvelopeMessagesCode = 4018
	CloudIntegrationGetResponseEnvelopeMessagesCode4019   CloudIntegrationGetResponseEnvelopeMessagesCode = 4019
	CloudIntegrationGetResponseEnvelopeMessagesCode4020   CloudIntegrationGetResponseEnvelopeMessagesCode = 4020
	CloudIntegrationGetResponseEnvelopeMessagesCode4021   CloudIntegrationGetResponseEnvelopeMessagesCode = 4021
	CloudIntegrationGetResponseEnvelopeMessagesCode4022   CloudIntegrationGetResponseEnvelopeMessagesCode = 4022
	CloudIntegrationGetResponseEnvelopeMessagesCode4023   CloudIntegrationGetResponseEnvelopeMessagesCode = 4023
	CloudIntegrationGetResponseEnvelopeMessagesCode5001   CloudIntegrationGetResponseEnvelopeMessagesCode = 5001
	CloudIntegrationGetResponseEnvelopeMessagesCode5002   CloudIntegrationGetResponseEnvelopeMessagesCode = 5002
	CloudIntegrationGetResponseEnvelopeMessagesCode5003   CloudIntegrationGetResponseEnvelopeMessagesCode = 5003
	CloudIntegrationGetResponseEnvelopeMessagesCode5004   CloudIntegrationGetResponseEnvelopeMessagesCode = 5004
	CloudIntegrationGetResponseEnvelopeMessagesCode102000 CloudIntegrationGetResponseEnvelopeMessagesCode = 102000
	CloudIntegrationGetResponseEnvelopeMessagesCode102001 CloudIntegrationGetResponseEnvelopeMessagesCode = 102001
	CloudIntegrationGetResponseEnvelopeMessagesCode102002 CloudIntegrationGetResponseEnvelopeMessagesCode = 102002
	CloudIntegrationGetResponseEnvelopeMessagesCode102003 CloudIntegrationGetResponseEnvelopeMessagesCode = 102003
	CloudIntegrationGetResponseEnvelopeMessagesCode102004 CloudIntegrationGetResponseEnvelopeMessagesCode = 102004
	CloudIntegrationGetResponseEnvelopeMessagesCode102005 CloudIntegrationGetResponseEnvelopeMessagesCode = 102005
	CloudIntegrationGetResponseEnvelopeMessagesCode102006 CloudIntegrationGetResponseEnvelopeMessagesCode = 102006
	CloudIntegrationGetResponseEnvelopeMessagesCode102007 CloudIntegrationGetResponseEnvelopeMessagesCode = 102007
	CloudIntegrationGetResponseEnvelopeMessagesCode102008 CloudIntegrationGetResponseEnvelopeMessagesCode = 102008
	CloudIntegrationGetResponseEnvelopeMessagesCode102009 CloudIntegrationGetResponseEnvelopeMessagesCode = 102009
	CloudIntegrationGetResponseEnvelopeMessagesCode102010 CloudIntegrationGetResponseEnvelopeMessagesCode = 102010
	CloudIntegrationGetResponseEnvelopeMessagesCode102011 CloudIntegrationGetResponseEnvelopeMessagesCode = 102011
	CloudIntegrationGetResponseEnvelopeMessagesCode102012 CloudIntegrationGetResponseEnvelopeMessagesCode = 102012
	CloudIntegrationGetResponseEnvelopeMessagesCode102013 CloudIntegrationGetResponseEnvelopeMessagesCode = 102013
	CloudIntegrationGetResponseEnvelopeMessagesCode102014 CloudIntegrationGetResponseEnvelopeMessagesCode = 102014
	CloudIntegrationGetResponseEnvelopeMessagesCode102015 CloudIntegrationGetResponseEnvelopeMessagesCode = 102015
	CloudIntegrationGetResponseEnvelopeMessagesCode102016 CloudIntegrationGetResponseEnvelopeMessagesCode = 102016
	CloudIntegrationGetResponseEnvelopeMessagesCode102017 CloudIntegrationGetResponseEnvelopeMessagesCode = 102017
	CloudIntegrationGetResponseEnvelopeMessagesCode102018 CloudIntegrationGetResponseEnvelopeMessagesCode = 102018
	CloudIntegrationGetResponseEnvelopeMessagesCode102019 CloudIntegrationGetResponseEnvelopeMessagesCode = 102019
	CloudIntegrationGetResponseEnvelopeMessagesCode102020 CloudIntegrationGetResponseEnvelopeMessagesCode = 102020
	CloudIntegrationGetResponseEnvelopeMessagesCode102021 CloudIntegrationGetResponseEnvelopeMessagesCode = 102021
	CloudIntegrationGetResponseEnvelopeMessagesCode102022 CloudIntegrationGetResponseEnvelopeMessagesCode = 102022
	CloudIntegrationGetResponseEnvelopeMessagesCode102023 CloudIntegrationGetResponseEnvelopeMessagesCode = 102023
	CloudIntegrationGetResponseEnvelopeMessagesCode102024 CloudIntegrationGetResponseEnvelopeMessagesCode = 102024
	CloudIntegrationGetResponseEnvelopeMessagesCode102025 CloudIntegrationGetResponseEnvelopeMessagesCode = 102025
	CloudIntegrationGetResponseEnvelopeMessagesCode102026 CloudIntegrationGetResponseEnvelopeMessagesCode = 102026
	CloudIntegrationGetResponseEnvelopeMessagesCode102027 CloudIntegrationGetResponseEnvelopeMessagesCode = 102027
	CloudIntegrationGetResponseEnvelopeMessagesCode102028 CloudIntegrationGetResponseEnvelopeMessagesCode = 102028
	CloudIntegrationGetResponseEnvelopeMessagesCode102029 CloudIntegrationGetResponseEnvelopeMessagesCode = 102029
	CloudIntegrationGetResponseEnvelopeMessagesCode102030 CloudIntegrationGetResponseEnvelopeMessagesCode = 102030
	CloudIntegrationGetResponseEnvelopeMessagesCode102031 CloudIntegrationGetResponseEnvelopeMessagesCode = 102031
	CloudIntegrationGetResponseEnvelopeMessagesCode102032 CloudIntegrationGetResponseEnvelopeMessagesCode = 102032
	CloudIntegrationGetResponseEnvelopeMessagesCode102033 CloudIntegrationGetResponseEnvelopeMessagesCode = 102033
	CloudIntegrationGetResponseEnvelopeMessagesCode102034 CloudIntegrationGetResponseEnvelopeMessagesCode = 102034
	CloudIntegrationGetResponseEnvelopeMessagesCode102035 CloudIntegrationGetResponseEnvelopeMessagesCode = 102035
	CloudIntegrationGetResponseEnvelopeMessagesCode102036 CloudIntegrationGetResponseEnvelopeMessagesCode = 102036
	CloudIntegrationGetResponseEnvelopeMessagesCode102037 CloudIntegrationGetResponseEnvelopeMessagesCode = 102037
	CloudIntegrationGetResponseEnvelopeMessagesCode102038 CloudIntegrationGetResponseEnvelopeMessagesCode = 102038
	CloudIntegrationGetResponseEnvelopeMessagesCode102039 CloudIntegrationGetResponseEnvelopeMessagesCode = 102039
	CloudIntegrationGetResponseEnvelopeMessagesCode102040 CloudIntegrationGetResponseEnvelopeMessagesCode = 102040
	CloudIntegrationGetResponseEnvelopeMessagesCode102041 CloudIntegrationGetResponseEnvelopeMessagesCode = 102041
	CloudIntegrationGetResponseEnvelopeMessagesCode102042 CloudIntegrationGetResponseEnvelopeMessagesCode = 102042
	CloudIntegrationGetResponseEnvelopeMessagesCode102043 CloudIntegrationGetResponseEnvelopeMessagesCode = 102043
	CloudIntegrationGetResponseEnvelopeMessagesCode102044 CloudIntegrationGetResponseEnvelopeMessagesCode = 102044
	CloudIntegrationGetResponseEnvelopeMessagesCode102045 CloudIntegrationGetResponseEnvelopeMessagesCode = 102045
	CloudIntegrationGetResponseEnvelopeMessagesCode102046 CloudIntegrationGetResponseEnvelopeMessagesCode = 102046
	CloudIntegrationGetResponseEnvelopeMessagesCode102047 CloudIntegrationGetResponseEnvelopeMessagesCode = 102047
	CloudIntegrationGetResponseEnvelopeMessagesCode102048 CloudIntegrationGetResponseEnvelopeMessagesCode = 102048
	CloudIntegrationGetResponseEnvelopeMessagesCode102049 CloudIntegrationGetResponseEnvelopeMessagesCode = 102049
	CloudIntegrationGetResponseEnvelopeMessagesCode102050 CloudIntegrationGetResponseEnvelopeMessagesCode = 102050
	CloudIntegrationGetResponseEnvelopeMessagesCode102051 CloudIntegrationGetResponseEnvelopeMessagesCode = 102051
	CloudIntegrationGetResponseEnvelopeMessagesCode102052 CloudIntegrationGetResponseEnvelopeMessagesCode = 102052
	CloudIntegrationGetResponseEnvelopeMessagesCode102053 CloudIntegrationGetResponseEnvelopeMessagesCode = 102053
	CloudIntegrationGetResponseEnvelopeMessagesCode102054 CloudIntegrationGetResponseEnvelopeMessagesCode = 102054
	CloudIntegrationGetResponseEnvelopeMessagesCode102055 CloudIntegrationGetResponseEnvelopeMessagesCode = 102055
	CloudIntegrationGetResponseEnvelopeMessagesCode102056 CloudIntegrationGetResponseEnvelopeMessagesCode = 102056
	CloudIntegrationGetResponseEnvelopeMessagesCode102057 CloudIntegrationGetResponseEnvelopeMessagesCode = 102057
	CloudIntegrationGetResponseEnvelopeMessagesCode102058 CloudIntegrationGetResponseEnvelopeMessagesCode = 102058
	CloudIntegrationGetResponseEnvelopeMessagesCode102059 CloudIntegrationGetResponseEnvelopeMessagesCode = 102059
	CloudIntegrationGetResponseEnvelopeMessagesCode102060 CloudIntegrationGetResponseEnvelopeMessagesCode = 102060
	CloudIntegrationGetResponseEnvelopeMessagesCode102061 CloudIntegrationGetResponseEnvelopeMessagesCode = 102061
	CloudIntegrationGetResponseEnvelopeMessagesCode102062 CloudIntegrationGetResponseEnvelopeMessagesCode = 102062
	CloudIntegrationGetResponseEnvelopeMessagesCode102063 CloudIntegrationGetResponseEnvelopeMessagesCode = 102063
	CloudIntegrationGetResponseEnvelopeMessagesCode102064 CloudIntegrationGetResponseEnvelopeMessagesCode = 102064
	CloudIntegrationGetResponseEnvelopeMessagesCode102065 CloudIntegrationGetResponseEnvelopeMessagesCode = 102065
	CloudIntegrationGetResponseEnvelopeMessagesCode102066 CloudIntegrationGetResponseEnvelopeMessagesCode = 102066
	CloudIntegrationGetResponseEnvelopeMessagesCode103001 CloudIntegrationGetResponseEnvelopeMessagesCode = 103001
	CloudIntegrationGetResponseEnvelopeMessagesCode103002 CloudIntegrationGetResponseEnvelopeMessagesCode = 103002
	CloudIntegrationGetResponseEnvelopeMessagesCode103003 CloudIntegrationGetResponseEnvelopeMessagesCode = 103003
	CloudIntegrationGetResponseEnvelopeMessagesCode103004 CloudIntegrationGetResponseEnvelopeMessagesCode = 103004
	CloudIntegrationGetResponseEnvelopeMessagesCode103005 CloudIntegrationGetResponseEnvelopeMessagesCode = 103005
	CloudIntegrationGetResponseEnvelopeMessagesCode103006 CloudIntegrationGetResponseEnvelopeMessagesCode = 103006
	CloudIntegrationGetResponseEnvelopeMessagesCode103007 CloudIntegrationGetResponseEnvelopeMessagesCode = 103007
	CloudIntegrationGetResponseEnvelopeMessagesCode103008 CloudIntegrationGetResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationGetResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationGetResponseEnvelopeMessagesCode1001, CloudIntegrationGetResponseEnvelopeMessagesCode1002, CloudIntegrationGetResponseEnvelopeMessagesCode1003, CloudIntegrationGetResponseEnvelopeMessagesCode1004, CloudIntegrationGetResponseEnvelopeMessagesCode1005, CloudIntegrationGetResponseEnvelopeMessagesCode1006, CloudIntegrationGetResponseEnvelopeMessagesCode1007, CloudIntegrationGetResponseEnvelopeMessagesCode1008, CloudIntegrationGetResponseEnvelopeMessagesCode1009, CloudIntegrationGetResponseEnvelopeMessagesCode1010, CloudIntegrationGetResponseEnvelopeMessagesCode1011, CloudIntegrationGetResponseEnvelopeMessagesCode1012, CloudIntegrationGetResponseEnvelopeMessagesCode1013, CloudIntegrationGetResponseEnvelopeMessagesCode1014, CloudIntegrationGetResponseEnvelopeMessagesCode1015, CloudIntegrationGetResponseEnvelopeMessagesCode1016, CloudIntegrationGetResponseEnvelopeMessagesCode1017, CloudIntegrationGetResponseEnvelopeMessagesCode2001, CloudIntegrationGetResponseEnvelopeMessagesCode2002, CloudIntegrationGetResponseEnvelopeMessagesCode2003, CloudIntegrationGetResponseEnvelopeMessagesCode2004, CloudIntegrationGetResponseEnvelopeMessagesCode2005, CloudIntegrationGetResponseEnvelopeMessagesCode2006, CloudIntegrationGetResponseEnvelopeMessagesCode2007, CloudIntegrationGetResponseEnvelopeMessagesCode2008, CloudIntegrationGetResponseEnvelopeMessagesCode2009, CloudIntegrationGetResponseEnvelopeMessagesCode2010, CloudIntegrationGetResponseEnvelopeMessagesCode2011, CloudIntegrationGetResponseEnvelopeMessagesCode2012, CloudIntegrationGetResponseEnvelopeMessagesCode2013, CloudIntegrationGetResponseEnvelopeMessagesCode2014, CloudIntegrationGetResponseEnvelopeMessagesCode2015, CloudIntegrationGetResponseEnvelopeMessagesCode2016, CloudIntegrationGetResponseEnvelopeMessagesCode2017, CloudIntegrationGetResponseEnvelopeMessagesCode2018, CloudIntegrationGetResponseEnvelopeMessagesCode2019, CloudIntegrationGetResponseEnvelopeMessagesCode2020, CloudIntegrationGetResponseEnvelopeMessagesCode2021, CloudIntegrationGetResponseEnvelopeMessagesCode2022, CloudIntegrationGetResponseEnvelopeMessagesCode3001, CloudIntegrationGetResponseEnvelopeMessagesCode3002, CloudIntegrationGetResponseEnvelopeMessagesCode3003, CloudIntegrationGetResponseEnvelopeMessagesCode3004, CloudIntegrationGetResponseEnvelopeMessagesCode3005, CloudIntegrationGetResponseEnvelopeMessagesCode3006, CloudIntegrationGetResponseEnvelopeMessagesCode3007, CloudIntegrationGetResponseEnvelopeMessagesCode4001, CloudIntegrationGetResponseEnvelopeMessagesCode4002, CloudIntegrationGetResponseEnvelopeMessagesCode4003, CloudIntegrationGetResponseEnvelopeMessagesCode4004, CloudIntegrationGetResponseEnvelopeMessagesCode4005, CloudIntegrationGetResponseEnvelopeMessagesCode4006, CloudIntegrationGetResponseEnvelopeMessagesCode4007, CloudIntegrationGetResponseEnvelopeMessagesCode4008, CloudIntegrationGetResponseEnvelopeMessagesCode4009, CloudIntegrationGetResponseEnvelopeMessagesCode4010, CloudIntegrationGetResponseEnvelopeMessagesCode4011, CloudIntegrationGetResponseEnvelopeMessagesCode4012, CloudIntegrationGetResponseEnvelopeMessagesCode4013, CloudIntegrationGetResponseEnvelopeMessagesCode4014, CloudIntegrationGetResponseEnvelopeMessagesCode4015, CloudIntegrationGetResponseEnvelopeMessagesCode4016, CloudIntegrationGetResponseEnvelopeMessagesCode4017, CloudIntegrationGetResponseEnvelopeMessagesCode4018, CloudIntegrationGetResponseEnvelopeMessagesCode4019, CloudIntegrationGetResponseEnvelopeMessagesCode4020, CloudIntegrationGetResponseEnvelopeMessagesCode4021, CloudIntegrationGetResponseEnvelopeMessagesCode4022, CloudIntegrationGetResponseEnvelopeMessagesCode4023, CloudIntegrationGetResponseEnvelopeMessagesCode5001, CloudIntegrationGetResponseEnvelopeMessagesCode5002, CloudIntegrationGetResponseEnvelopeMessagesCode5003, CloudIntegrationGetResponseEnvelopeMessagesCode5004, CloudIntegrationGetResponseEnvelopeMessagesCode102000, CloudIntegrationGetResponseEnvelopeMessagesCode102001, CloudIntegrationGetResponseEnvelopeMessagesCode102002, CloudIntegrationGetResponseEnvelopeMessagesCode102003, CloudIntegrationGetResponseEnvelopeMessagesCode102004, CloudIntegrationGetResponseEnvelopeMessagesCode102005, CloudIntegrationGetResponseEnvelopeMessagesCode102006, CloudIntegrationGetResponseEnvelopeMessagesCode102007, CloudIntegrationGetResponseEnvelopeMessagesCode102008, CloudIntegrationGetResponseEnvelopeMessagesCode102009, CloudIntegrationGetResponseEnvelopeMessagesCode102010, CloudIntegrationGetResponseEnvelopeMessagesCode102011, CloudIntegrationGetResponseEnvelopeMessagesCode102012, CloudIntegrationGetResponseEnvelopeMessagesCode102013, CloudIntegrationGetResponseEnvelopeMessagesCode102014, CloudIntegrationGetResponseEnvelopeMessagesCode102015, CloudIntegrationGetResponseEnvelopeMessagesCode102016, CloudIntegrationGetResponseEnvelopeMessagesCode102017, CloudIntegrationGetResponseEnvelopeMessagesCode102018, CloudIntegrationGetResponseEnvelopeMessagesCode102019, CloudIntegrationGetResponseEnvelopeMessagesCode102020, CloudIntegrationGetResponseEnvelopeMessagesCode102021, CloudIntegrationGetResponseEnvelopeMessagesCode102022, CloudIntegrationGetResponseEnvelopeMessagesCode102023, CloudIntegrationGetResponseEnvelopeMessagesCode102024, CloudIntegrationGetResponseEnvelopeMessagesCode102025, CloudIntegrationGetResponseEnvelopeMessagesCode102026, CloudIntegrationGetResponseEnvelopeMessagesCode102027, CloudIntegrationGetResponseEnvelopeMessagesCode102028, CloudIntegrationGetResponseEnvelopeMessagesCode102029, CloudIntegrationGetResponseEnvelopeMessagesCode102030, CloudIntegrationGetResponseEnvelopeMessagesCode102031, CloudIntegrationGetResponseEnvelopeMessagesCode102032, CloudIntegrationGetResponseEnvelopeMessagesCode102033, CloudIntegrationGetResponseEnvelopeMessagesCode102034, CloudIntegrationGetResponseEnvelopeMessagesCode102035, CloudIntegrationGetResponseEnvelopeMessagesCode102036, CloudIntegrationGetResponseEnvelopeMessagesCode102037, CloudIntegrationGetResponseEnvelopeMessagesCode102038, CloudIntegrationGetResponseEnvelopeMessagesCode102039, CloudIntegrationGetResponseEnvelopeMessagesCode102040, CloudIntegrationGetResponseEnvelopeMessagesCode102041, CloudIntegrationGetResponseEnvelopeMessagesCode102042, CloudIntegrationGetResponseEnvelopeMessagesCode102043, CloudIntegrationGetResponseEnvelopeMessagesCode102044, CloudIntegrationGetResponseEnvelopeMessagesCode102045, CloudIntegrationGetResponseEnvelopeMessagesCode102046, CloudIntegrationGetResponseEnvelopeMessagesCode102047, CloudIntegrationGetResponseEnvelopeMessagesCode102048, CloudIntegrationGetResponseEnvelopeMessagesCode102049, CloudIntegrationGetResponseEnvelopeMessagesCode102050, CloudIntegrationGetResponseEnvelopeMessagesCode102051, CloudIntegrationGetResponseEnvelopeMessagesCode102052, CloudIntegrationGetResponseEnvelopeMessagesCode102053, CloudIntegrationGetResponseEnvelopeMessagesCode102054, CloudIntegrationGetResponseEnvelopeMessagesCode102055, CloudIntegrationGetResponseEnvelopeMessagesCode102056, CloudIntegrationGetResponseEnvelopeMessagesCode102057, CloudIntegrationGetResponseEnvelopeMessagesCode102058, CloudIntegrationGetResponseEnvelopeMessagesCode102059, CloudIntegrationGetResponseEnvelopeMessagesCode102060, CloudIntegrationGetResponseEnvelopeMessagesCode102061, CloudIntegrationGetResponseEnvelopeMessagesCode102062, CloudIntegrationGetResponseEnvelopeMessagesCode102063, CloudIntegrationGetResponseEnvelopeMessagesCode102064, CloudIntegrationGetResponseEnvelopeMessagesCode102065, CloudIntegrationGetResponseEnvelopeMessagesCode102066, CloudIntegrationGetResponseEnvelopeMessagesCode103001, CloudIntegrationGetResponseEnvelopeMessagesCode103002, CloudIntegrationGetResponseEnvelopeMessagesCode103003, CloudIntegrationGetResponseEnvelopeMessagesCode103004, CloudIntegrationGetResponseEnvelopeMessagesCode103005, CloudIntegrationGetResponseEnvelopeMessagesCode103006, CloudIntegrationGetResponseEnvelopeMessagesCode103007, CloudIntegrationGetResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationGetResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                              `json:"l10n_key"`
	LoggableError string                                              `json:"loggable_error"`
	TemplateData  interface{}                                         `json:"template_data"`
	TraceID       string                                              `json:"trace_id"`
	JSON          cloudIntegrationGetResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeMessagesMetaJSON contains the JSON metadata
// for the struct [CloudIntegrationGetResponseEnvelopeMessagesMeta]
type cloudIntegrationGetResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationGetResponseEnvelopeMessagesSource struct {
	Parameter           string                                                `json:"parameter"`
	ParameterValueIndex int64                                                 `json:"parameter_value_index"`
	Pointer             string                                                `json:"pointer"`
	JSON                cloudIntegrationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CloudIntegrationGetResponseEnvelopeMessagesSource]
type cloudIntegrationGetResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type CloudIntegrationInitialSetupResponseEnvelope struct {
	Errors   []CloudIntegrationInitialSetupResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudIntegrationInitialSetupResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudIntegrationInitialSetupResponse                   `json:"result,required"`
	Success  bool                                                   `json:"success,required"`
	JSON     cloudIntegrationInitialSetupResponseEnvelopeJSON       `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeJSON contains the JSON metadata for
// the struct [CloudIntegrationInitialSetupResponseEnvelope]
type cloudIntegrationInitialSetupResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeErrors struct {
	Code             CloudIntegrationInitialSetupResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Meta             CloudIntegrationInitialSetupResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           CloudIntegrationInitialSetupResponseEnvelopeErrorsSource `json:"source"`
	JSON             cloudIntegrationInitialSetupResponseEnvelopeErrorsJSON   `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudIntegrationInitialSetupResponseEnvelopeErrors]
type cloudIntegrationInitialSetupResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeErrorsCode int64

const (
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1001   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1002   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1003   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1004   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1005   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1006   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1007   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1008   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1008
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1009   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1009
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1010   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1010
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1011   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1011
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1012   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1012
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1013   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1013
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1014   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1014
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1015   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1015
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1016   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1016
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1017   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 1017
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2001   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2002   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2003   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2004   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2005   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2006   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2007   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2008   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2008
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2009   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2009
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2010   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2010
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2011   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2011
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2012   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2012
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2013   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2013
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2014   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2014
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2015   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2015
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2016   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2016
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2017   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2017
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2018   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2018
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2019   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2019
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2020   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2020
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2021   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2021
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2022   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 2022
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3001   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3002   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3003   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3004   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3005   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3006   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3007   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 3007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4001   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4002   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4003   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4004   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4005   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4006   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4007   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4008   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4008
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4009   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4009
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4010   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4010
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4011   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4011
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4012   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4012
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4013   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4013
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4014   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4014
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4015   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4015
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4016   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4016
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4017   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4017
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4018   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4018
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4019   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4019
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4020   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4020
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4021   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4021
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4022   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4022
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4023   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 4023
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5001   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 5001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5002   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 5002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5003   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 5003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5004   CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 5004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102000 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102000
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102001 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102002 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102003 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102004 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102005 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102006 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102007 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102008 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102008
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102009 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102009
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102010 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102010
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102011 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102011
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102012 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102012
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102013 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102013
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102014 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102014
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102015 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102015
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102016 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102016
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102017 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102017
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102018 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102018
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102019 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102019
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102020 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102020
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102021 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102021
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102022 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102022
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102023 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102023
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102024 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102024
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102025 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102025
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102026 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102026
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102027 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102027
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102028 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102028
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102029 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102029
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102030 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102030
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102031 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102031
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102032 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102032
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102033 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102033
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102034 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102034
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102035 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102035
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102036 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102036
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102037 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102037
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102038 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102038
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102039 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102039
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102040 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102040
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102041 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102041
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102042 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102042
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102043 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102043
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102044 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102044
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102045 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102045
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102046 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102046
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102047 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102047
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102048 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102048
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102049 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102049
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102050 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102050
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102051 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102051
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102052 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102052
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102053 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102053
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102054 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102054
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102055 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102055
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102056 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102056
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102057 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102057
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102058 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102058
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102059 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102059
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102060 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102060
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102061 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102061
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102062 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102062
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102063 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102063
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102064 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102064
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102065 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102065
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102066 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 102066
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103001 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103001
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103002 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103002
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103003 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103003
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103004 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103004
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103005 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103005
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103006 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103006
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103007 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103007
	CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103008 CloudIntegrationInitialSetupResponseEnvelopeErrorsCode = 103008
)

func (r CloudIntegrationInitialSetupResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1008, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1009, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1010, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1011, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1012, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1013, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1014, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1015, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1016, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode1017, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2008, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2009, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2010, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2011, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2012, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2013, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2014, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2015, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2016, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2017, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2018, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2019, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2020, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2021, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode2022, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode3007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4008, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4009, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4010, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4011, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4012, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4013, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4014, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4015, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4016, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4017, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4018, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4019, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4020, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4021, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4022, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode4023, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode5004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102000, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102008, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102009, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102010, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102011, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102012, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102013, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102014, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102015, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102016, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102017, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102018, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102019, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102020, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102021, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102022, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102023, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102024, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102025, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102026, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102027, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102028, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102029, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102030, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102031, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102032, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102033, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102034, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102035, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102036, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102037, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102038, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102039, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102040, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102041, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102042, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102043, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102044, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102045, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102046, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102047, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102048, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102049, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102050, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102051, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102052, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102053, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102054, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102055, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102056, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102057, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102058, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102059, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102060, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102061, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102062, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102063, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102064, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102065, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode102066, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103001, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103002, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103003, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103004, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103005, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103006, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103007, CloudIntegrationInitialSetupResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type CloudIntegrationInitialSetupResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                                     `json:"l10n_key"`
	LoggableError string                                                     `json:"loggable_error"`
	TemplateData  interface{}                                                `json:"template_data"`
	TraceID       string                                                     `json:"trace_id"`
	JSON          cloudIntegrationInitialSetupResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeErrorsMetaJSON contains the JSON
// metadata for the struct [CloudIntegrationInitialSetupResponseEnvelopeErrorsMeta]
type cloudIntegrationInitialSetupResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeErrorsSource struct {
	Parameter           string                                                       `json:"parameter"`
	ParameterValueIndex int64                                                        `json:"parameter_value_index"`
	Pointer             string                                                       `json:"pointer"`
	JSON                cloudIntegrationInitialSetupResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CloudIntegrationInitialSetupResponseEnvelopeErrorsSource]
type cloudIntegrationInitialSetupResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeMessages struct {
	Code             CloudIntegrationInitialSetupResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Meta             CloudIntegrationInitialSetupResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           CloudIntegrationInitialSetupResponseEnvelopeMessagesSource `json:"source"`
	JSON             cloudIntegrationInitialSetupResponseEnvelopeMessagesJSON   `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CloudIntegrationInitialSetupResponseEnvelopeMessages]
type cloudIntegrationInitialSetupResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeMessagesCode int64

const (
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1001   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1002   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1003   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1004   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1005   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1006   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1007   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1008   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1008
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1009   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1009
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1010   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1010
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1011   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1011
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1012   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1012
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1013   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1013
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1014   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1014
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1015   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1015
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1016   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1016
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1017   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 1017
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2001   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2002   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2003   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2004   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2005   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2006   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2007   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2008   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2008
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2009   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2009
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2010   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2010
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2011   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2011
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2012   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2012
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2013   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2013
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2014   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2014
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2015   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2015
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2016   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2016
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2017   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2017
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2018   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2018
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2019   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2019
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2020   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2020
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2021   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2021
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2022   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 2022
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3001   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3002   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3003   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3004   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3005   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3006   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3007   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 3007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4001   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4002   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4003   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4004   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4005   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4006   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4007   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4008   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4008
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4009   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4009
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4010   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4010
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4011   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4011
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4012   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4012
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4013   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4013
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4014   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4014
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4015   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4015
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4016   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4016
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4017   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4017
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4018   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4018
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4019   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4019
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4020   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4020
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4021   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4021
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4022   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4022
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4023   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 4023
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5001   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 5001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5002   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 5002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5003   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 5003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5004   CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 5004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102000 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102000
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102001 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102002 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102003 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102004 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102005 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102006 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102007 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102008 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102008
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102009 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102009
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102010 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102010
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102011 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102011
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102012 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102012
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102013 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102013
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102014 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102014
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102015 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102015
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102016 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102016
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102017 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102017
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102018 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102018
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102019 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102019
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102020 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102020
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102021 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102021
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102022 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102022
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102023 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102023
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102024 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102024
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102025 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102025
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102026 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102026
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102027 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102027
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102028 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102028
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102029 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102029
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102030 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102030
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102031 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102031
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102032 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102032
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102033 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102033
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102034 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102034
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102035 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102035
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102036 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102036
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102037 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102037
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102038 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102038
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102039 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102039
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102040 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102040
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102041 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102041
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102042 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102042
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102043 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102043
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102044 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102044
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102045 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102045
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102046 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102046
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102047 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102047
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102048 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102048
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102049 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102049
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102050 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102050
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102051 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102051
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102052 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102052
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102053 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102053
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102054 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102054
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102055 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102055
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102056 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102056
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102057 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102057
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102058 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102058
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102059 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102059
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102060 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102060
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102061 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102061
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102062 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102062
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102063 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102063
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102064 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102064
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102065 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102065
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102066 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 102066
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103001 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103001
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103002 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103002
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103003 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103003
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103004 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103004
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103005 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103005
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103006 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103006
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103007 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103007
	CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103008 CloudIntegrationInitialSetupResponseEnvelopeMessagesCode = 103008
)

func (r CloudIntegrationInitialSetupResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1008, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1009, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1010, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1011, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1012, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1013, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1014, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1015, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1016, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode1017, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2008, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2009, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2010, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2011, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2012, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2013, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2014, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2015, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2016, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2017, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2018, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2019, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2020, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2021, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode2022, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode3007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4008, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4009, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4010, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4011, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4012, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4013, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4014, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4015, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4016, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4017, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4018, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4019, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4020, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4021, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4022, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode4023, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode5004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102000, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102008, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102009, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102010, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102011, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102012, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102013, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102014, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102015, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102016, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102017, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102018, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102019, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102020, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102021, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102022, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102023, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102024, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102025, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102026, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102027, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102028, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102029, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102030, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102031, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102032, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102033, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102034, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102035, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102036, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102037, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102038, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102039, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102040, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102041, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102042, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102043, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102044, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102045, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102046, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102047, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102048, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102049, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102050, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102051, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102052, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102053, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102054, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102055, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102056, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102057, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102058, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102059, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102060, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102061, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102062, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102063, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102064, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102065, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode102066, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103001, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103002, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103003, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103004, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103005, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103006, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103007, CloudIntegrationInitialSetupResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type CloudIntegrationInitialSetupResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                       `json:"l10n_key"`
	LoggableError string                                                       `json:"loggable_error"`
	TemplateData  interface{}                                                  `json:"template_data"`
	TraceID       string                                                       `json:"trace_id"`
	JSON          cloudIntegrationInitialSetupResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeMessagesMetaJSON contains the JSON
// metadata for the struct
// [CloudIntegrationInitialSetupResponseEnvelopeMessagesMeta]
type cloudIntegrationInitialSetupResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type CloudIntegrationInitialSetupResponseEnvelopeMessagesSource struct {
	Parameter           string                                                         `json:"parameter"`
	ParameterValueIndex int64                                                          `json:"parameter_value_index"`
	Pointer             string                                                         `json:"pointer"`
	JSON                cloudIntegrationInitialSetupResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cloudIntegrationInitialSetupResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CloudIntegrationInitialSetupResponseEnvelopeMessagesSource]
type cloudIntegrationInitialSetupResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudIntegrationInitialSetupResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudIntegrationInitialSetupResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
