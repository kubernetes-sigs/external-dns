// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// PolicyService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r *PolicyService) {
	r = &PolicyService{}
	r.Options = opts
	return
}

// Creates a new Notification policy.
func (r *PolicyService) New(ctx context.Context, params PolicyNewParams, opts ...option.RequestOption) (res *PolicyNewResponse, err error) {
	var env PolicyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Notification policy.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *PolicyUpdateResponse, err error) {
	var env PolicyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all Notification policies.
func (r *PolicyService) List(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Policy], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get a list of all Notification policies.
func (r *PolicyService) ListAutoPaging(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Policy] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Notification policy.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (res *PolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details for a single policy.
func (r *PolicyService) Get(ctx context.Context, policyID string, query PolicyGetParams, opts ...option.RequestOption) (res *Policy, err error) {
	var env PolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of IDs that will be used when dispatching a notification. IDs for email
// type will be the email address.
type Mechanism struct {
	Email     []MechanismEmail     `json:"email"`
	Pagerduty []MechanismPagerduty `json:"pagerduty"`
	Webhooks  []MechanismWebhook   `json:"webhooks"`
	JSON      mechanismJSON        `json:"-"`
}

// mechanismJSON contains the JSON metadata for the struct [Mechanism]
type mechanismJSON struct {
	Email       apijson.Field
	Pagerduty   apijson.Field
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Mechanism) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismJSON) RawJSON() string {
	return r.raw
}

type MechanismEmail struct {
	// The email address
	ID   string             `json:"id"`
	JSON mechanismEmailJSON `json:"-"`
}

// mechanismEmailJSON contains the JSON metadata for the struct [MechanismEmail]
type mechanismEmailJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismEmailJSON) RawJSON() string {
	return r.raw
}

type MechanismPagerduty struct {
	// UUID
	ID   string                 `json:"id"`
	JSON mechanismPagerdutyJSON `json:"-"`
}

// mechanismPagerdutyJSON contains the JSON metadata for the struct
// [MechanismPagerduty]
type mechanismPagerdutyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismPagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismPagerdutyJSON) RawJSON() string {
	return r.raw
}

type MechanismWebhook struct {
	// UUID
	ID   string               `json:"id"`
	JSON mechanismWebhookJSON `json:"-"`
}

// mechanismWebhookJSON contains the JSON metadata for the struct
// [MechanismWebhook]
type mechanismWebhookJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismWebhookJSON) RawJSON() string {
	return r.raw
}

// List of IDs that will be used when dispatching a notification. IDs for email
// type will be the email address.
type MechanismParam struct {
	Email     param.Field[[]MechanismEmailParam]     `json:"email"`
	Pagerduty param.Field[[]MechanismPagerdutyParam] `json:"pagerduty"`
	Webhooks  param.Field[[]MechanismWebhookParam]   `json:"webhooks"`
}

func (r MechanismParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismEmailParam struct {
	// The email address
	ID param.Field[string] `json:"id"`
}

func (r MechanismEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismPagerdutyParam struct {
	// UUID
	ID param.Field[string] `json:"id"`
}

func (r MechanismPagerdutyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismWebhookParam struct {
	// UUID
	ID param.Field[string] `json:"id"`
}

func (r MechanismWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Policy struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval string `json:"alert_interval"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType PolicyAlertType `json:"alert_type"`
	Created   time.Time       `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters PolicyFilter `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms Mechanism `json:"mechanisms"`
	Modified   time.Time `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string     `json:"name"`
	JSON policyJSON `json:"-"`
}

// policyJSON contains the JSON metadata for the struct [Policy]
type policyJSON struct {
	ID            apijson.Field
	AlertInterval apijson.Field
	AlertType     apijson.Field
	Created       apijson.Field
	Description   apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Mechanisms    apijson.Field
	Modified      apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyJSON) RawJSON() string {
	return r.raw
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type PolicyAlertType string

const (
	PolicyAlertTypeAccessCustomCertificateExpirationType         PolicyAlertType = "access_custom_certificate_expiration_type"
	PolicyAlertTypeAdvancedDDoSAttackL4Alert                     PolicyAlertType = "advanced_ddos_attack_l4_alert"
	PolicyAlertTypeAdvancedDDoSAttackL7Alert                     PolicyAlertType = "advanced_ddos_attack_l7_alert"
	PolicyAlertTypeAdvancedHTTPAlertError                        PolicyAlertType = "advanced_http_alert_error"
	PolicyAlertTypeBGPHijackNotification                         PolicyAlertType = "bgp_hijack_notification"
	PolicyAlertTypeBillingUsageAlert                             PolicyAlertType = "billing_usage_alert"
	PolicyAlertTypeBlockNotificationBlockRemoved                 PolicyAlertType = "block_notification_block_removed"
	PolicyAlertTypeBlockNotificationNewBlock                     PolicyAlertType = "block_notification_new_block"
	PolicyAlertTypeBlockNotificationReviewRejected               PolicyAlertType = "block_notification_review_rejected"
	PolicyAlertTypeBotTrafficBasicAlert                          PolicyAlertType = "bot_traffic_basic_alert"
	PolicyAlertTypeBrandProtectionAlert                          PolicyAlertType = "brand_protection_alert"
	PolicyAlertTypeBrandProtectionDigest                         PolicyAlertType = "brand_protection_digest"
	PolicyAlertTypeClickhouseAlertFwAnomaly                      PolicyAlertType = "clickhouse_alert_fw_anomaly"
	PolicyAlertTypeClickhouseAlertFwEntAnomaly                   PolicyAlertType = "clickhouse_alert_fw_ent_anomaly"
	PolicyAlertTypeCloudforceOneRequestNotification              PolicyAlertType = "cloudforce_one_request_notification"
	PolicyAlertTypeCustomAnalytics                               PolicyAlertType = "custom_analytics"
	PolicyAlertTypeCustomBotDetectionAlert                       PolicyAlertType = "custom_bot_detection_alert"
	PolicyAlertTypeCustomSSLCertificateEventType                 PolicyAlertType = "custom_ssl_certificate_event_type"
	PolicyAlertTypeDedicatedSSLCertificateEventType              PolicyAlertType = "dedicated_ssl_certificate_event_type"
	PolicyAlertTypeDeviceConnectivityAnomalyAlert                PolicyAlertType = "device_connectivity_anomaly_alert"
	PolicyAlertTypeDosAttackL4                                   PolicyAlertType = "dos_attack_l4"
	PolicyAlertTypeDosAttackL7                                   PolicyAlertType = "dos_attack_l7"
	PolicyAlertTypeExpiringServiceTokenAlert                     PolicyAlertType = "expiring_service_token_alert"
	PolicyAlertTypeFailingLogpushJobDisabledAlert                PolicyAlertType = "failing_logpush_job_disabled_alert"
	PolicyAlertTypeFbmAutoAdvertisement                          PolicyAlertType = "fbm_auto_advertisement"
	PolicyAlertTypeFbmDosdAttack                                 PolicyAlertType = "fbm_dosd_attack"
	PolicyAlertTypeFbmVolumetricAttack                           PolicyAlertType = "fbm_volumetric_attack"
	PolicyAlertTypeHealthCheckStatusNotification                 PolicyAlertType = "health_check_status_notification"
	PolicyAlertTypeHostnameAopCustomCertificateExpirationType    PolicyAlertType = "hostname_aop_custom_certificate_expiration_type"
	PolicyAlertTypeHTTPAlertEdgeError                            PolicyAlertType = "http_alert_edge_error"
	PolicyAlertTypeHTTPAlertOriginError                          PolicyAlertType = "http_alert_origin_error"
	PolicyAlertTypeImageNotification                             PolicyAlertType = "image_notification"
	PolicyAlertTypeImageResizingNotification                     PolicyAlertType = "image_resizing_notification"
	PolicyAlertTypeIncidentAlert                                 PolicyAlertType = "incident_alert"
	PolicyAlertTypeLoadBalancingHealthAlert                      PolicyAlertType = "load_balancing_health_alert"
	PolicyAlertTypeLoadBalancingPoolEnablementAlert              PolicyAlertType = "load_balancing_pool_enablement_alert"
	PolicyAlertTypeLogoMatchAlert                                PolicyAlertType = "logo_match_alert"
	PolicyAlertTypeMagicTunnelHealthCheckEvent                   PolicyAlertType = "magic_tunnel_health_check_event"
	PolicyAlertTypeMagicWANTunnelHealth                          PolicyAlertType = "magic_wan_tunnel_health"
	PolicyAlertTypeMaintenanceEventNotification                  PolicyAlertType = "maintenance_event_notification"
	PolicyAlertTypeMTLSCertificateStoreCertificateExpirationType PolicyAlertType = "mtls_certificate_store_certificate_expiration_type"
	PolicyAlertTypePagesEventAlert                               PolicyAlertType = "pages_event_alert"
	PolicyAlertTypeRadarNotification                             PolicyAlertType = "radar_notification"
	PolicyAlertTypeRealOriginMonitoring                          PolicyAlertType = "real_origin_monitoring"
	PolicyAlertTypeScriptmonitorAlertNewCodeChangeDetections     PolicyAlertType = "scriptmonitor_alert_new_code_change_detections"
	PolicyAlertTypeScriptmonitorAlertNewHosts                    PolicyAlertType = "scriptmonitor_alert_new_hosts"
	PolicyAlertTypeScriptmonitorAlertNewMaliciousHosts           PolicyAlertType = "scriptmonitor_alert_new_malicious_hosts"
	PolicyAlertTypeScriptmonitorAlertNewMaliciousScripts         PolicyAlertType = "scriptmonitor_alert_new_malicious_scripts"
	PolicyAlertTypeScriptmonitorAlertNewMaliciousURL             PolicyAlertType = "scriptmonitor_alert_new_malicious_url"
	PolicyAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     PolicyAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	PolicyAlertTypeScriptmonitorAlertNewResources                PolicyAlertType = "scriptmonitor_alert_new_resources"
	PolicyAlertTypeSecondaryDNSAllPrimariesFailing               PolicyAlertType = "secondary_dns_all_primaries_failing"
	PolicyAlertTypeSecondaryDNSPrimariesFailing                  PolicyAlertType = "secondary_dns_primaries_failing"
	PolicyAlertTypeSecondaryDNSWarning                           PolicyAlertType = "secondary_dns_warning"
	PolicyAlertTypeSecondaryDNSZoneSuccessfullyUpdated           PolicyAlertType = "secondary_dns_zone_successfully_updated"
	PolicyAlertTypeSecondaryDNSZoneValidationWarning             PolicyAlertType = "secondary_dns_zone_validation_warning"
	PolicyAlertTypeSecurityInsightsAlert                         PolicyAlertType = "security_insights_alert"
	PolicyAlertTypeSentinelAlert                                 PolicyAlertType = "sentinel_alert"
	PolicyAlertTypeStreamLiveNotifications                       PolicyAlertType = "stream_live_notifications"
	PolicyAlertTypeSyntheticTestLatencyAlert                     PolicyAlertType = "synthetic_test_latency_alert"
	PolicyAlertTypeSyntheticTestLowAvailabilityAlert             PolicyAlertType = "synthetic_test_low_availability_alert"
	PolicyAlertTypeTrafficAnomaliesAlert                         PolicyAlertType = "traffic_anomalies_alert"
	PolicyAlertTypeTunnelHealthEvent                             PolicyAlertType = "tunnel_health_event"
	PolicyAlertTypeTunnelUpdateEvent                             PolicyAlertType = "tunnel_update_event"
	PolicyAlertTypeUniversalSSLEventType                         PolicyAlertType = "universal_ssl_event_type"
	PolicyAlertTypeWebAnalyticsMetricsUpdate                     PolicyAlertType = "web_analytics_metrics_update"
	PolicyAlertTypeZoneAopCustomCertificateExpirationType        PolicyAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r PolicyAlertType) IsKnown() bool {
	switch r {
	case PolicyAlertTypeAccessCustomCertificateExpirationType, PolicyAlertTypeAdvancedDDoSAttackL4Alert, PolicyAlertTypeAdvancedDDoSAttackL7Alert, PolicyAlertTypeAdvancedHTTPAlertError, PolicyAlertTypeBGPHijackNotification, PolicyAlertTypeBillingUsageAlert, PolicyAlertTypeBlockNotificationBlockRemoved, PolicyAlertTypeBlockNotificationNewBlock, PolicyAlertTypeBlockNotificationReviewRejected, PolicyAlertTypeBotTrafficBasicAlert, PolicyAlertTypeBrandProtectionAlert, PolicyAlertTypeBrandProtectionDigest, PolicyAlertTypeClickhouseAlertFwAnomaly, PolicyAlertTypeClickhouseAlertFwEntAnomaly, PolicyAlertTypeCloudforceOneRequestNotification, PolicyAlertTypeCustomAnalytics, PolicyAlertTypeCustomBotDetectionAlert, PolicyAlertTypeCustomSSLCertificateEventType, PolicyAlertTypeDedicatedSSLCertificateEventType, PolicyAlertTypeDeviceConnectivityAnomalyAlert, PolicyAlertTypeDosAttackL4, PolicyAlertTypeDosAttackL7, PolicyAlertTypeExpiringServiceTokenAlert, PolicyAlertTypeFailingLogpushJobDisabledAlert, PolicyAlertTypeFbmAutoAdvertisement, PolicyAlertTypeFbmDosdAttack, PolicyAlertTypeFbmVolumetricAttack, PolicyAlertTypeHealthCheckStatusNotification, PolicyAlertTypeHostnameAopCustomCertificateExpirationType, PolicyAlertTypeHTTPAlertEdgeError, PolicyAlertTypeHTTPAlertOriginError, PolicyAlertTypeImageNotification, PolicyAlertTypeImageResizingNotification, PolicyAlertTypeIncidentAlert, PolicyAlertTypeLoadBalancingHealthAlert, PolicyAlertTypeLoadBalancingPoolEnablementAlert, PolicyAlertTypeLogoMatchAlert, PolicyAlertTypeMagicTunnelHealthCheckEvent, PolicyAlertTypeMagicWANTunnelHealth, PolicyAlertTypeMaintenanceEventNotification, PolicyAlertTypeMTLSCertificateStoreCertificateExpirationType, PolicyAlertTypePagesEventAlert, PolicyAlertTypeRadarNotification, PolicyAlertTypeRealOriginMonitoring, PolicyAlertTypeScriptmonitorAlertNewCodeChangeDetections, PolicyAlertTypeScriptmonitorAlertNewHosts, PolicyAlertTypeScriptmonitorAlertNewMaliciousHosts, PolicyAlertTypeScriptmonitorAlertNewMaliciousScripts, PolicyAlertTypeScriptmonitorAlertNewMaliciousURL, PolicyAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, PolicyAlertTypeScriptmonitorAlertNewResources, PolicyAlertTypeSecondaryDNSAllPrimariesFailing, PolicyAlertTypeSecondaryDNSPrimariesFailing, PolicyAlertTypeSecondaryDNSWarning, PolicyAlertTypeSecondaryDNSZoneSuccessfullyUpdated, PolicyAlertTypeSecondaryDNSZoneValidationWarning, PolicyAlertTypeSecurityInsightsAlert, PolicyAlertTypeSentinelAlert, PolicyAlertTypeStreamLiveNotifications, PolicyAlertTypeSyntheticTestLatencyAlert, PolicyAlertTypeSyntheticTestLowAvailabilityAlert, PolicyAlertTypeTrafficAnomaliesAlert, PolicyAlertTypeTunnelHealthEvent, PolicyAlertTypeTunnelUpdateEvent, PolicyAlertTypeUniversalSSLEventType, PolicyAlertTypeWebAnalyticsMetricsUpdate, PolicyAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type PolicyFilter struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Usage depends on specific alert type
	AlertTriggerPreferencesValue []string `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled []string `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment []string `json:"environment"`
	// Used for configuring pages_event_alert
	Event []string `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource []string `json:"event_source"`
	// Usage depends on specific alert type
	EventType []string `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy []string `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID []string `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact []PolicyFilterIncidentImpact `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID []string `json:"input_id"`
	// Used for configuring security_insights_alert
	InsightClass []string `json:"insight_class"`
	// Used for configuring billing_usage_alert
	Limit []string `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag []string `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond []string `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth []string `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus []string `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond []string `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID []string `json:"pool_id"`
	// Usage depends on specific alert type
	POPNames []string `json:"pop_names"`
	// Used for configuring billing_usage_alert
	Product []string `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID []string `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol []string `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag []string `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond []string `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors []string `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services []string `json:"services"`
	// Usage depends on specific alert type
	Slo []string `json:"slo"`
	// Used for configuring health_check_status_notification
	Status []string `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname []string `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP []string `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName []string `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions []PolicyFilterTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Usage depends on specific alert type
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string         `json:"zones"`
	JSON  policyFilterJSON `json:"-"`
}

// policyFilterJSON contains the JSON metadata for the struct [PolicyFilter]
type policyFilterJSON struct {
	Actions                      apijson.Field
	AffectedASNs                 apijson.Field
	AffectedComponents           apijson.Field
	AffectedLocations            apijson.Field
	AirportCode                  apijson.Field
	AlertTriggerPreferences      apijson.Field
	AlertTriggerPreferencesValue apijson.Field
	Enabled                      apijson.Field
	Environment                  apijson.Field
	Event                        apijson.Field
	EventSource                  apijson.Field
	EventType                    apijson.Field
	GroupBy                      apijson.Field
	HealthCheckID                apijson.Field
	IncidentImpact               apijson.Field
	InputID                      apijson.Field
	InsightClass                 apijson.Field
	Limit                        apijson.Field
	LogoTag                      apijson.Field
	MegabitsPerSecond            apijson.Field
	NewHealth                    apijson.Field
	NewStatus                    apijson.Field
	PacketsPerSecond             apijson.Field
	PoolID                       apijson.Field
	POPNames                     apijson.Field
	Product                      apijson.Field
	ProjectID                    apijson.Field
	Protocol                     apijson.Field
	QueryTag                     apijson.Field
	RequestsPerSecond            apijson.Field
	Selectors                    apijson.Field
	Services                     apijson.Field
	Slo                          apijson.Field
	Status                       apijson.Field
	TargetHostname               apijson.Field
	TargetIP                     apijson.Field
	TargetZoneName               apijson.Field
	TrafficExclusions            apijson.Field
	TunnelID                     apijson.Field
	TunnelName                   apijson.Field
	Where                        apijson.Field
	Zones                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PolicyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyFilterJSON) RawJSON() string {
	return r.raw
}

type PolicyFilterIncidentImpact string

const (
	PolicyFilterIncidentImpactIncidentImpactNone     PolicyFilterIncidentImpact = "INCIDENT_IMPACT_NONE"
	PolicyFilterIncidentImpactIncidentImpactMinor    PolicyFilterIncidentImpact = "INCIDENT_IMPACT_MINOR"
	PolicyFilterIncidentImpactIncidentImpactMajor    PolicyFilterIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	PolicyFilterIncidentImpactIncidentImpactCritical PolicyFilterIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

func (r PolicyFilterIncidentImpact) IsKnown() bool {
	switch r {
	case PolicyFilterIncidentImpactIncidentImpactNone, PolicyFilterIncidentImpactIncidentImpactMinor, PolicyFilterIncidentImpactIncidentImpactMajor, PolicyFilterIncidentImpactIncidentImpactCritical:
		return true
	}
	return false
}

type PolicyFilterTrafficExclusion string

const (
	PolicyFilterTrafficExclusionSecurityEvents PolicyFilterTrafficExclusion = "security_events"
)

func (r PolicyFilterTrafficExclusion) IsKnown() bool {
	switch r {
	case PolicyFilterTrafficExclusionSecurityEvents:
		return true
	}
	return false
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type PolicyFilterParam struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedASNs param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Usage depends on specific alert type
	AlertTriggerPreferencesValue param.Field[[]string] `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled param.Field[[]string] `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment param.Field[[]string] `json:"environment"`
	// Used for configuring pages_event_alert
	Event param.Field[[]string] `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource param.Field[[]string] `json:"event_source"`
	// Usage depends on specific alert type
	EventType param.Field[[]string] `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy param.Field[[]string] `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID param.Field[[]string] `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact param.Field[[]PolicyFilterIncidentImpact] `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID param.Field[[]string] `json:"input_id"`
	// Used for configuring security_insights_alert
	InsightClass param.Field[[]string] `json:"insight_class"`
	// Used for configuring billing_usage_alert
	Limit param.Field[[]string] `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag param.Field[[]string] `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond param.Field[[]string] `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth param.Field[[]string] `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus param.Field[[]string] `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond param.Field[[]string] `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID param.Field[[]string] `json:"pool_id"`
	// Usage depends on specific alert type
	POPNames param.Field[[]string] `json:"pop_names"`
	// Used for configuring billing_usage_alert
	Product param.Field[[]string] `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID param.Field[[]string] `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol param.Field[[]string] `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag param.Field[[]string] `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond param.Field[[]string] `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors param.Field[[]string] `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services param.Field[[]string] `json:"services"`
	// Usage depends on specific alert type
	Slo param.Field[[]string] `json:"slo"`
	// Used for configuring health_check_status_notification
	Status param.Field[[]string] `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname param.Field[[]string] `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP param.Field[[]string] `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName param.Field[[]string] `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions param.Field[[]PolicyFilterTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Usage depends on specific alert type
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r PolicyFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PolicyNewResponse struct {
	// UUID
	ID   string                `json:"id"`
	JSON policyNewResponseJSON `json:"-"`
}

// policyNewResponseJSON contains the JSON metadata for the struct
// [PolicyNewResponse]
type policyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseJSON) RawJSON() string {
	return r.raw
}

type PolicyUpdateResponse struct {
	// UUID
	ID   string                   `json:"id"`
	JSON policyUpdateResponseJSON `json:"-"`
}

// policyUpdateResponseJSON contains the JSON metadata for the struct
// [PolicyUpdateResponse]
type policyUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PolicyDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    PolicyDeleteResponseSuccess    `json:"success,required"`
	ResultInfo PolicyDeleteResponseResultInfo `json:"result_info"`
	JSON       policyDeleteResponseJSON       `json:"-"`
}

// policyDeleteResponseJSON contains the JSON metadata for the struct
// [PolicyDeleteResponse]
type policyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyDeleteResponseSuccess bool

const (
	PolicyDeleteResponseSuccessTrue PolicyDeleteResponseSuccess = true
)

func (r PolicyDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case PolicyDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type PolicyDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                            `json:"total_count"`
	JSON       policyDeleteResponseResultInfoJSON `json:"-"`
}

// policyDeleteResponseResultInfoJSON contains the JSON metadata for the struct
// [PolicyDeleteResponseResultInfo]
type policyDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type PolicyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[PolicyNewParamsAlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[MechanismParam] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval param.Field[string] `json:"alert_interval"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[PolicyFilterParam] `json:"filters"`
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type PolicyNewParamsAlertType string

const (
	PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType         PolicyNewParamsAlertType = "access_custom_certificate_expiration_type"
	PolicyNewParamsAlertTypeAdvancedDDoSAttackL4Alert                     PolicyNewParamsAlertType = "advanced_ddos_attack_l4_alert"
	PolicyNewParamsAlertTypeAdvancedDDoSAttackL7Alert                     PolicyNewParamsAlertType = "advanced_ddos_attack_l7_alert"
	PolicyNewParamsAlertTypeAdvancedHTTPAlertError                        PolicyNewParamsAlertType = "advanced_http_alert_error"
	PolicyNewParamsAlertTypeBGPHijackNotification                         PolicyNewParamsAlertType = "bgp_hijack_notification"
	PolicyNewParamsAlertTypeBillingUsageAlert                             PolicyNewParamsAlertType = "billing_usage_alert"
	PolicyNewParamsAlertTypeBlockNotificationBlockRemoved                 PolicyNewParamsAlertType = "block_notification_block_removed"
	PolicyNewParamsAlertTypeBlockNotificationNewBlock                     PolicyNewParamsAlertType = "block_notification_new_block"
	PolicyNewParamsAlertTypeBlockNotificationReviewRejected               PolicyNewParamsAlertType = "block_notification_review_rejected"
	PolicyNewParamsAlertTypeBotTrafficBasicAlert                          PolicyNewParamsAlertType = "bot_traffic_basic_alert"
	PolicyNewParamsAlertTypeBrandProtectionAlert                          PolicyNewParamsAlertType = "brand_protection_alert"
	PolicyNewParamsAlertTypeBrandProtectionDigest                         PolicyNewParamsAlertType = "brand_protection_digest"
	PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly                      PolicyNewParamsAlertType = "clickhouse_alert_fw_anomaly"
	PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly                   PolicyNewParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	PolicyNewParamsAlertTypeCloudforceOneRequestNotification              PolicyNewParamsAlertType = "cloudforce_one_request_notification"
	PolicyNewParamsAlertTypeCustomAnalytics                               PolicyNewParamsAlertType = "custom_analytics"
	PolicyNewParamsAlertTypeCustomBotDetectionAlert                       PolicyNewParamsAlertType = "custom_bot_detection_alert"
	PolicyNewParamsAlertTypeCustomSSLCertificateEventType                 PolicyNewParamsAlertType = "custom_ssl_certificate_event_type"
	PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType              PolicyNewParamsAlertType = "dedicated_ssl_certificate_event_type"
	PolicyNewParamsAlertTypeDeviceConnectivityAnomalyAlert                PolicyNewParamsAlertType = "device_connectivity_anomaly_alert"
	PolicyNewParamsAlertTypeDosAttackL4                                   PolicyNewParamsAlertType = "dos_attack_l4"
	PolicyNewParamsAlertTypeDosAttackL7                                   PolicyNewParamsAlertType = "dos_attack_l7"
	PolicyNewParamsAlertTypeExpiringServiceTokenAlert                     PolicyNewParamsAlertType = "expiring_service_token_alert"
	PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert                PolicyNewParamsAlertType = "failing_logpush_job_disabled_alert"
	PolicyNewParamsAlertTypeFbmAutoAdvertisement                          PolicyNewParamsAlertType = "fbm_auto_advertisement"
	PolicyNewParamsAlertTypeFbmDosdAttack                                 PolicyNewParamsAlertType = "fbm_dosd_attack"
	PolicyNewParamsAlertTypeFbmVolumetricAttack                           PolicyNewParamsAlertType = "fbm_volumetric_attack"
	PolicyNewParamsAlertTypeHealthCheckStatusNotification                 PolicyNewParamsAlertType = "health_check_status_notification"
	PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType    PolicyNewParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	PolicyNewParamsAlertTypeHTTPAlertEdgeError                            PolicyNewParamsAlertType = "http_alert_edge_error"
	PolicyNewParamsAlertTypeHTTPAlertOriginError                          PolicyNewParamsAlertType = "http_alert_origin_error"
	PolicyNewParamsAlertTypeImageNotification                             PolicyNewParamsAlertType = "image_notification"
	PolicyNewParamsAlertTypeImageResizingNotification                     PolicyNewParamsAlertType = "image_resizing_notification"
	PolicyNewParamsAlertTypeIncidentAlert                                 PolicyNewParamsAlertType = "incident_alert"
	PolicyNewParamsAlertTypeLoadBalancingHealthAlert                      PolicyNewParamsAlertType = "load_balancing_health_alert"
	PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert              PolicyNewParamsAlertType = "load_balancing_pool_enablement_alert"
	PolicyNewParamsAlertTypeLogoMatchAlert                                PolicyNewParamsAlertType = "logo_match_alert"
	PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent                   PolicyNewParamsAlertType = "magic_tunnel_health_check_event"
	PolicyNewParamsAlertTypeMagicWANTunnelHealth                          PolicyNewParamsAlertType = "magic_wan_tunnel_health"
	PolicyNewParamsAlertTypeMaintenanceEventNotification                  PolicyNewParamsAlertType = "maintenance_event_notification"
	PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType PolicyNewParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	PolicyNewParamsAlertTypePagesEventAlert                               PolicyNewParamsAlertType = "pages_event_alert"
	PolicyNewParamsAlertTypeRadarNotification                             PolicyNewParamsAlertType = "radar_notification"
	PolicyNewParamsAlertTypeRealOriginMonitoring                          PolicyNewParamsAlertType = "real_origin_monitoring"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     PolicyNewParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts                    PolicyNewParamsAlertType = "scriptmonitor_alert_new_hosts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL             PolicyNewParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     PolicyNewParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	PolicyNewParamsAlertTypeScriptmonitorAlertNewResources                PolicyNewParamsAlertType = "scriptmonitor_alert_new_resources"
	PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing               PolicyNewParamsAlertType = "secondary_dns_all_primaries_failing"
	PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing                  PolicyNewParamsAlertType = "secondary_dns_primaries_failing"
	PolicyNewParamsAlertTypeSecondaryDNSWarning                           PolicyNewParamsAlertType = "secondary_dns_warning"
	PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           PolicyNewParamsAlertType = "secondary_dns_zone_successfully_updated"
	PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning             PolicyNewParamsAlertType = "secondary_dns_zone_validation_warning"
	PolicyNewParamsAlertTypeSecurityInsightsAlert                         PolicyNewParamsAlertType = "security_insights_alert"
	PolicyNewParamsAlertTypeSentinelAlert                                 PolicyNewParamsAlertType = "sentinel_alert"
	PolicyNewParamsAlertTypeStreamLiveNotifications                       PolicyNewParamsAlertType = "stream_live_notifications"
	PolicyNewParamsAlertTypeSyntheticTestLatencyAlert                     PolicyNewParamsAlertType = "synthetic_test_latency_alert"
	PolicyNewParamsAlertTypeSyntheticTestLowAvailabilityAlert             PolicyNewParamsAlertType = "synthetic_test_low_availability_alert"
	PolicyNewParamsAlertTypeTrafficAnomaliesAlert                         PolicyNewParamsAlertType = "traffic_anomalies_alert"
	PolicyNewParamsAlertTypeTunnelHealthEvent                             PolicyNewParamsAlertType = "tunnel_health_event"
	PolicyNewParamsAlertTypeTunnelUpdateEvent                             PolicyNewParamsAlertType = "tunnel_update_event"
	PolicyNewParamsAlertTypeUniversalSSLEventType                         PolicyNewParamsAlertType = "universal_ssl_event_type"
	PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate                     PolicyNewParamsAlertType = "web_analytics_metrics_update"
	PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType        PolicyNewParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r PolicyNewParamsAlertType) IsKnown() bool {
	switch r {
	case PolicyNewParamsAlertTypeAccessCustomCertificateExpirationType, PolicyNewParamsAlertTypeAdvancedDDoSAttackL4Alert, PolicyNewParamsAlertTypeAdvancedDDoSAttackL7Alert, PolicyNewParamsAlertTypeAdvancedHTTPAlertError, PolicyNewParamsAlertTypeBGPHijackNotification, PolicyNewParamsAlertTypeBillingUsageAlert, PolicyNewParamsAlertTypeBlockNotificationBlockRemoved, PolicyNewParamsAlertTypeBlockNotificationNewBlock, PolicyNewParamsAlertTypeBlockNotificationReviewRejected, PolicyNewParamsAlertTypeBotTrafficBasicAlert, PolicyNewParamsAlertTypeBrandProtectionAlert, PolicyNewParamsAlertTypeBrandProtectionDigest, PolicyNewParamsAlertTypeClickhouseAlertFwAnomaly, PolicyNewParamsAlertTypeClickhouseAlertFwEntAnomaly, PolicyNewParamsAlertTypeCloudforceOneRequestNotification, PolicyNewParamsAlertTypeCustomAnalytics, PolicyNewParamsAlertTypeCustomBotDetectionAlert, PolicyNewParamsAlertTypeCustomSSLCertificateEventType, PolicyNewParamsAlertTypeDedicatedSSLCertificateEventType, PolicyNewParamsAlertTypeDeviceConnectivityAnomalyAlert, PolicyNewParamsAlertTypeDosAttackL4, PolicyNewParamsAlertTypeDosAttackL7, PolicyNewParamsAlertTypeExpiringServiceTokenAlert, PolicyNewParamsAlertTypeFailingLogpushJobDisabledAlert, PolicyNewParamsAlertTypeFbmAutoAdvertisement, PolicyNewParamsAlertTypeFbmDosdAttack, PolicyNewParamsAlertTypeFbmVolumetricAttack, PolicyNewParamsAlertTypeHealthCheckStatusNotification, PolicyNewParamsAlertTypeHostnameAopCustomCertificateExpirationType, PolicyNewParamsAlertTypeHTTPAlertEdgeError, PolicyNewParamsAlertTypeHTTPAlertOriginError, PolicyNewParamsAlertTypeImageNotification, PolicyNewParamsAlertTypeImageResizingNotification, PolicyNewParamsAlertTypeIncidentAlert, PolicyNewParamsAlertTypeLoadBalancingHealthAlert, PolicyNewParamsAlertTypeLoadBalancingPoolEnablementAlert, PolicyNewParamsAlertTypeLogoMatchAlert, PolicyNewParamsAlertTypeMagicTunnelHealthCheckEvent, PolicyNewParamsAlertTypeMagicWANTunnelHealth, PolicyNewParamsAlertTypeMaintenanceEventNotification, PolicyNewParamsAlertTypeMTLSCertificateStoreCertificateExpirationType, PolicyNewParamsAlertTypePagesEventAlert, PolicyNewParamsAlertTypeRadarNotification, PolicyNewParamsAlertTypeRealOriginMonitoring, PolicyNewParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections, PolicyNewParamsAlertTypeScriptmonitorAlertNewHosts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousHosts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousScripts, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaliciousURL, PolicyNewParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, PolicyNewParamsAlertTypeScriptmonitorAlertNewResources, PolicyNewParamsAlertTypeSecondaryDNSAllPrimariesFailing, PolicyNewParamsAlertTypeSecondaryDNSPrimariesFailing, PolicyNewParamsAlertTypeSecondaryDNSWarning, PolicyNewParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated, PolicyNewParamsAlertTypeSecondaryDNSZoneValidationWarning, PolicyNewParamsAlertTypeSecurityInsightsAlert, PolicyNewParamsAlertTypeSentinelAlert, PolicyNewParamsAlertTypeStreamLiveNotifications, PolicyNewParamsAlertTypeSyntheticTestLatencyAlert, PolicyNewParamsAlertTypeSyntheticTestLowAvailabilityAlert, PolicyNewParamsAlertTypeTrafficAnomaliesAlert, PolicyNewParamsAlertTypeTunnelHealthEvent, PolicyNewParamsAlertTypeTunnelUpdateEvent, PolicyNewParamsAlertTypeUniversalSSLEventType, PolicyNewParamsAlertTypeWebAnalyticsMetricsUpdate, PolicyNewParamsAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

type PolicyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PolicyNewResponseEnvelopeSuccess `json:"success,required"`
	Result  PolicyNewResponse                `json:"result"`
	JSON    policyNewResponseEnvelopeJSON    `json:"-"`
}

// policyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyNewResponseEnvelope]
type policyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyNewResponseEnvelopeSuccess bool

const (
	PolicyNewResponseEnvelopeSuccessTrue PolicyNewResponseEnvelopeSuccess = true
)

func (r PolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval param.Field[string] `json:"alert_interval"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[PolicyUpdateParamsAlertType] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[PolicyFilterParam] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[MechanismParam] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type PolicyUpdateParamsAlertType string

const (
	PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType         PolicyUpdateParamsAlertType = "access_custom_certificate_expiration_type"
	PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL4Alert                     PolicyUpdateParamsAlertType = "advanced_ddos_attack_l4_alert"
	PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL7Alert                     PolicyUpdateParamsAlertType = "advanced_ddos_attack_l7_alert"
	PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError                        PolicyUpdateParamsAlertType = "advanced_http_alert_error"
	PolicyUpdateParamsAlertTypeBGPHijackNotification                         PolicyUpdateParamsAlertType = "bgp_hijack_notification"
	PolicyUpdateParamsAlertTypeBillingUsageAlert                             PolicyUpdateParamsAlertType = "billing_usage_alert"
	PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved                 PolicyUpdateParamsAlertType = "block_notification_block_removed"
	PolicyUpdateParamsAlertTypeBlockNotificationNewBlock                     PolicyUpdateParamsAlertType = "block_notification_new_block"
	PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected               PolicyUpdateParamsAlertType = "block_notification_review_rejected"
	PolicyUpdateParamsAlertTypeBotTrafficBasicAlert                          PolicyUpdateParamsAlertType = "bot_traffic_basic_alert"
	PolicyUpdateParamsAlertTypeBrandProtectionAlert                          PolicyUpdateParamsAlertType = "brand_protection_alert"
	PolicyUpdateParamsAlertTypeBrandProtectionDigest                         PolicyUpdateParamsAlertType = "brand_protection_digest"
	PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly                      PolicyUpdateParamsAlertType = "clickhouse_alert_fw_anomaly"
	PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly                   PolicyUpdateParamsAlertType = "clickhouse_alert_fw_ent_anomaly"
	PolicyUpdateParamsAlertTypeCloudforceOneRequestNotification              PolicyUpdateParamsAlertType = "cloudforce_one_request_notification"
	PolicyUpdateParamsAlertTypeCustomAnalytics                               PolicyUpdateParamsAlertType = "custom_analytics"
	PolicyUpdateParamsAlertTypeCustomBotDetectionAlert                       PolicyUpdateParamsAlertType = "custom_bot_detection_alert"
	PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType                 PolicyUpdateParamsAlertType = "custom_ssl_certificate_event_type"
	PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType              PolicyUpdateParamsAlertType = "dedicated_ssl_certificate_event_type"
	PolicyUpdateParamsAlertTypeDeviceConnectivityAnomalyAlert                PolicyUpdateParamsAlertType = "device_connectivity_anomaly_alert"
	PolicyUpdateParamsAlertTypeDosAttackL4                                   PolicyUpdateParamsAlertType = "dos_attack_l4"
	PolicyUpdateParamsAlertTypeDosAttackL7                                   PolicyUpdateParamsAlertType = "dos_attack_l7"
	PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert                     PolicyUpdateParamsAlertType = "expiring_service_token_alert"
	PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert                PolicyUpdateParamsAlertType = "failing_logpush_job_disabled_alert"
	PolicyUpdateParamsAlertTypeFbmAutoAdvertisement                          PolicyUpdateParamsAlertType = "fbm_auto_advertisement"
	PolicyUpdateParamsAlertTypeFbmDosdAttack                                 PolicyUpdateParamsAlertType = "fbm_dosd_attack"
	PolicyUpdateParamsAlertTypeFbmVolumetricAttack                           PolicyUpdateParamsAlertType = "fbm_volumetric_attack"
	PolicyUpdateParamsAlertTypeHealthCheckStatusNotification                 PolicyUpdateParamsAlertType = "health_check_status_notification"
	PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType    PolicyUpdateParamsAlertType = "hostname_aop_custom_certificate_expiration_type"
	PolicyUpdateParamsAlertTypeHTTPAlertEdgeError                            PolicyUpdateParamsAlertType = "http_alert_edge_error"
	PolicyUpdateParamsAlertTypeHTTPAlertOriginError                          PolicyUpdateParamsAlertType = "http_alert_origin_error"
	PolicyUpdateParamsAlertTypeImageNotification                             PolicyUpdateParamsAlertType = "image_notification"
	PolicyUpdateParamsAlertTypeImageResizingNotification                     PolicyUpdateParamsAlertType = "image_resizing_notification"
	PolicyUpdateParamsAlertTypeIncidentAlert                                 PolicyUpdateParamsAlertType = "incident_alert"
	PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert                      PolicyUpdateParamsAlertType = "load_balancing_health_alert"
	PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert              PolicyUpdateParamsAlertType = "load_balancing_pool_enablement_alert"
	PolicyUpdateParamsAlertTypeLogoMatchAlert                                PolicyUpdateParamsAlertType = "logo_match_alert"
	PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent                   PolicyUpdateParamsAlertType = "magic_tunnel_health_check_event"
	PolicyUpdateParamsAlertTypeMagicWANTunnelHealth                          PolicyUpdateParamsAlertType = "magic_wan_tunnel_health"
	PolicyUpdateParamsAlertTypeMaintenanceEventNotification                  PolicyUpdateParamsAlertType = "maintenance_event_notification"
	PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType PolicyUpdateParamsAlertType = "mtls_certificate_store_certificate_expiration_type"
	PolicyUpdateParamsAlertTypePagesEventAlert                               PolicyUpdateParamsAlertType = "pages_event_alert"
	PolicyUpdateParamsAlertTypeRadarNotification                             PolicyUpdateParamsAlertType = "radar_notification"
	PolicyUpdateParamsAlertTypeRealOriginMonitoring                          PolicyUpdateParamsAlertType = "real_origin_monitoring"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections     PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_code_change_detections"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts                    PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_hosts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts           PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_hosts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts         PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_scripts"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL             PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_malicious_url"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL     PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_max_length_resource_url"
	PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources                PolicyUpdateParamsAlertType = "scriptmonitor_alert_new_resources"
	PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing               PolicyUpdateParamsAlertType = "secondary_dns_all_primaries_failing"
	PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing                  PolicyUpdateParamsAlertType = "secondary_dns_primaries_failing"
	PolicyUpdateParamsAlertTypeSecondaryDNSWarning                           PolicyUpdateParamsAlertType = "secondary_dns_warning"
	PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated           PolicyUpdateParamsAlertType = "secondary_dns_zone_successfully_updated"
	PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning             PolicyUpdateParamsAlertType = "secondary_dns_zone_validation_warning"
	PolicyUpdateParamsAlertTypeSecurityInsightsAlert                         PolicyUpdateParamsAlertType = "security_insights_alert"
	PolicyUpdateParamsAlertTypeSentinelAlert                                 PolicyUpdateParamsAlertType = "sentinel_alert"
	PolicyUpdateParamsAlertTypeStreamLiveNotifications                       PolicyUpdateParamsAlertType = "stream_live_notifications"
	PolicyUpdateParamsAlertTypeSyntheticTestLatencyAlert                     PolicyUpdateParamsAlertType = "synthetic_test_latency_alert"
	PolicyUpdateParamsAlertTypeSyntheticTestLowAvailabilityAlert             PolicyUpdateParamsAlertType = "synthetic_test_low_availability_alert"
	PolicyUpdateParamsAlertTypeTrafficAnomaliesAlert                         PolicyUpdateParamsAlertType = "traffic_anomalies_alert"
	PolicyUpdateParamsAlertTypeTunnelHealthEvent                             PolicyUpdateParamsAlertType = "tunnel_health_event"
	PolicyUpdateParamsAlertTypeTunnelUpdateEvent                             PolicyUpdateParamsAlertType = "tunnel_update_event"
	PolicyUpdateParamsAlertTypeUniversalSSLEventType                         PolicyUpdateParamsAlertType = "universal_ssl_event_type"
	PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate                     PolicyUpdateParamsAlertType = "web_analytics_metrics_update"
	PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType        PolicyUpdateParamsAlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r PolicyUpdateParamsAlertType) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsAlertTypeAccessCustomCertificateExpirationType, PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL4Alert, PolicyUpdateParamsAlertTypeAdvancedDDoSAttackL7Alert, PolicyUpdateParamsAlertTypeAdvancedHTTPAlertError, PolicyUpdateParamsAlertTypeBGPHijackNotification, PolicyUpdateParamsAlertTypeBillingUsageAlert, PolicyUpdateParamsAlertTypeBlockNotificationBlockRemoved, PolicyUpdateParamsAlertTypeBlockNotificationNewBlock, PolicyUpdateParamsAlertTypeBlockNotificationReviewRejected, PolicyUpdateParamsAlertTypeBotTrafficBasicAlert, PolicyUpdateParamsAlertTypeBrandProtectionAlert, PolicyUpdateParamsAlertTypeBrandProtectionDigest, PolicyUpdateParamsAlertTypeClickhouseAlertFwAnomaly, PolicyUpdateParamsAlertTypeClickhouseAlertFwEntAnomaly, PolicyUpdateParamsAlertTypeCloudforceOneRequestNotification, PolicyUpdateParamsAlertTypeCustomAnalytics, PolicyUpdateParamsAlertTypeCustomBotDetectionAlert, PolicyUpdateParamsAlertTypeCustomSSLCertificateEventType, PolicyUpdateParamsAlertTypeDedicatedSSLCertificateEventType, PolicyUpdateParamsAlertTypeDeviceConnectivityAnomalyAlert, PolicyUpdateParamsAlertTypeDosAttackL4, PolicyUpdateParamsAlertTypeDosAttackL7, PolicyUpdateParamsAlertTypeExpiringServiceTokenAlert, PolicyUpdateParamsAlertTypeFailingLogpushJobDisabledAlert, PolicyUpdateParamsAlertTypeFbmAutoAdvertisement, PolicyUpdateParamsAlertTypeFbmDosdAttack, PolicyUpdateParamsAlertTypeFbmVolumetricAttack, PolicyUpdateParamsAlertTypeHealthCheckStatusNotification, PolicyUpdateParamsAlertTypeHostnameAopCustomCertificateExpirationType, PolicyUpdateParamsAlertTypeHTTPAlertEdgeError, PolicyUpdateParamsAlertTypeHTTPAlertOriginError, PolicyUpdateParamsAlertTypeImageNotification, PolicyUpdateParamsAlertTypeImageResizingNotification, PolicyUpdateParamsAlertTypeIncidentAlert, PolicyUpdateParamsAlertTypeLoadBalancingHealthAlert, PolicyUpdateParamsAlertTypeLoadBalancingPoolEnablementAlert, PolicyUpdateParamsAlertTypeLogoMatchAlert, PolicyUpdateParamsAlertTypeMagicTunnelHealthCheckEvent, PolicyUpdateParamsAlertTypeMagicWANTunnelHealth, PolicyUpdateParamsAlertTypeMaintenanceEventNotification, PolicyUpdateParamsAlertTypeMTLSCertificateStoreCertificateExpirationType, PolicyUpdateParamsAlertTypePagesEventAlert, PolicyUpdateParamsAlertTypeRadarNotification, PolicyUpdateParamsAlertTypeRealOriginMonitoring, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewCodeChangeDetections, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewHosts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousHosts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousScripts, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaliciousURL, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewMaxLengthResourceURL, PolicyUpdateParamsAlertTypeScriptmonitorAlertNewResources, PolicyUpdateParamsAlertTypeSecondaryDNSAllPrimariesFailing, PolicyUpdateParamsAlertTypeSecondaryDNSPrimariesFailing, PolicyUpdateParamsAlertTypeSecondaryDNSWarning, PolicyUpdateParamsAlertTypeSecondaryDNSZoneSuccessfullyUpdated, PolicyUpdateParamsAlertTypeSecondaryDNSZoneValidationWarning, PolicyUpdateParamsAlertTypeSecurityInsightsAlert, PolicyUpdateParamsAlertTypeSentinelAlert, PolicyUpdateParamsAlertTypeStreamLiveNotifications, PolicyUpdateParamsAlertTypeSyntheticTestLatencyAlert, PolicyUpdateParamsAlertTypeSyntheticTestLowAvailabilityAlert, PolicyUpdateParamsAlertTypeTrafficAnomaliesAlert, PolicyUpdateParamsAlertTypeTunnelHealthEvent, PolicyUpdateParamsAlertTypeTunnelUpdateEvent, PolicyUpdateParamsAlertTypeUniversalSSLEventType, PolicyUpdateParamsAlertTypeWebAnalyticsMetricsUpdate, PolicyUpdateParamsAlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

type PolicyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  PolicyUpdateResponse                `json:"result"`
	JSON    policyUpdateResponseEnvelopeJSON    `json:"-"`
}

// policyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyUpdateResponseEnvelope]
type policyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyUpdateResponseEnvelopeSuccess bool

const (
	PolicyUpdateResponseEnvelopeSuccessTrue PolicyUpdateResponseEnvelopeSuccess = true
)

func (r PolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type PolicyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PolicyGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Policy                           `json:"result"`
	JSON    policyGetResponseEnvelopeJSON    `json:"-"`
}

// policyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyGetResponseEnvelope]
type policyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyGetResponseEnvelopeSuccess bool

const (
	PolicyGetResponseEnvelopeSuccessTrue PolicyGetResponseEnvelopeSuccess = true
)

func (r PolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
