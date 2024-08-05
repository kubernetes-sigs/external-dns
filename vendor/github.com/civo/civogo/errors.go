package civogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

// Errors raised by package civogo
var (
	ResponseDecodeFailedError = constError("ResponseDecodeFailed")
	DisabledServiceError      = constError("DisabledServiceError")
	NoAPIKeySuppliedError     = constError("NoAPIKeySuppliedError")
	MultipleMatchesError      = constError("MultipleMatchesError")
	ZeroMatchesError          = constError("ZeroMatchesError")
	IDisEmptyError            = constError("IDisEmptyError")
	TimeoutError              = constError("TimeoutError")
	RegionUnavailableError    = constError("RegionUnavailable")

	CivoStatsdRecordFailedError = constError("CivoStatsdRecordFailedError")
	AuthenticationFailedError   = constError("AuthenticationFailedError")
	CommonError                 = constError("Error")

	// Volume Error
	CannotRescueNewVolumeError              = constError("CannotRescueNewVolumeError")
	CannotRestoreNewVolumeError             = constError("CannotRestoreNewVolumeError")
	CannotScaleAlreadyRescalingClusterError = constError("CannotScaleAlreadyRescalingClusterError")
	VolumeInvalidSizeError                  = constError("VolumeInvalidSizeError")

	DatabaseAccountDestroyError      = constError("DatabaseAccountDestroyError")
	DatabaseAccountNotFoundError     = constError("DatabaseAccountNotFoundError")
	DatabaseAccountAccessDeniedError = constError("DatabaseAccountAccessDeniedError")
	DatabaseCreatingAccountError     = constError("DatabaseCreatingAccountError")

	DatabaseUpdatingAccountError = constError("DatabaseUpdatingAccountError")
	DatabaseAccountStatsError    = constError("DatabaseAccountStatsError")
	DatabaseActionListingError   = constError("DatabaseActionListingError")
	DatabaseActionCreateError    = constError("DatabaseActionCreateError")
	DatabaseAPIKeyCreateError    = constError("DatabaseApiKeyCreateError")
	DatabaseAPIKeyDuplicateError = constError("DatabaseApiKeyDuplicateError")
	DatabaseAPIKeyNotFoundError  = constError("DatabaseApiKeyNotFoundError")

	DatabaseAPIkeyDestroyError           = constError("DatabaseAPIkeyDestroyError")
	DatabaseAuditLogListingError         = constError("DatabaseAuditLogListingError")
	DatabaseBlueprintNotFoundError       = constError("DatabaseBlueprintNotFoundError")
	DatabaseBlueprintDeleteFailedError   = constError("DatabaseBlueprintDeleteFailedError")
	DatabaseBlueprintCreateError         = constError("DatabaseBlueprintCreateError")
	DatabaseBlueprintUpdateError         = constError("DatabaseBlueprintUpdateError")
	ParameterEmptyVolumeIDError          = constError("ParameterEmptyVolumeIDError")
	ParameterEmptyOpenstackVolumeIDError = constError("ParameterEmptyOpenstackVolumeIDError")
	DatabaseChangeAPIKeyError            = constError("DatabaseChangeAPIKeyError")
	DatabaseChargeListingError           = constError("DatabaseChargeListingError")
	DatabaseConnectionFailedError        = constError("DatabaseConnectionFailedError")

	DatabaseDNSDomainCreateError        = constError("DatabaseDnsDomainCreateError")
	DatabaseDNSDomainUpdateError        = constError("DatabaseDnsDomainUpdateError")
	DatabaseDNSDomainDuplicateNameError = constError("DatabaseDnsDomainDuplicateNameError")
	DatabaseDNSDomainNotFoundError      = constError("DatabaseDNSDomainNotFoundError")
	DatabaseDNSRecordCreateError        = constError("DatabaseDNSRecordCreateError")
	DatabaseDNSRecordNotFoundError      = constError("DatabaseDNSRecordNotFoundError")
	DatabaseDNSRecordUpdateError        = constError("DatabaseDNSRecordUpdateError")
	DatabaseListingDNSDomainsError      = constError("DatabaseListingDNSDomainsError")

	DatabaseFirewallCreateError           = constError("DatabaseFirewallCreateError")
	DatabaseFirewallRulesInvalidParams    = constError("DatabaseFirewallRulesInvalidParams")
	DatabaseFirewallDuplicateNameError    = constError("DatabaseFirewallDuplicateNameError")
	DatabaseFirewallMismatchError         = constError("DatabaseFirewallMismatchError")
	DatabaseFirewallNotFoundError         = constError("DatabaseFirewallNotFoundError")
	DatabaseFirewallSaveFailedError       = constError("DatabaseFirewallSaveFailedError")
	DatabaseFirewallDeleteFailedError     = constError("DatabaseFirewallDeleteFailedError")
	DatabaseFirewallRuleCreateError       = constError("DatabaseFirewallRuleCreateError")
	DatabaseFirewallRuleDeleteFailedError = constError("DatabaseFirewallRuleDeleteFailedError")
	DatabaseFirewallRulesFindError        = constError("DatabaseFirewallRulesFindError")
	DatabaseListingFirewallsError         = constError("DatabaseListingFirewallsError")
	FirewallDuplicateError                = constError("FirewallDuplicateError")

	// Instances Errors
	DatabaseInstanceAlreadyinRescueStateError              = constError("DatabaseInstanceAlreadyinRescueStateError")
	DatabaseInstanceBuildError                             = constError("DatabaseInstanceBuildError")
	DatabaseInstanceBuildMultipleWithExistingPublicIPError = constError("DatabaseInstanceBuildMultipleWithExistingPublicIPError")
	DatabaseInstanceCreateError                            = constError("DatabaseInstanceCreateError")
	DatabaseInstanceSnapshotTooBigError                    = constError("DatabaseInstanceSnapshotTooBigError")
	DatabaseInstanceDuplicateError                         = constError("DatabaseInstanceDuplicateError")
	DatabaseInstanceDuplicateNameError                     = constError("DatabaseInstanceDuplicateNameError")
	DatabaseInstanceListError                              = constError("DatabaseInstanceListError")
	DatabaseInstanceNotFoundError                          = constError("DatabaseInstanceNotFoundError")
	DatabaseInstanceNotInOpenStackError                    = constError("DatabaseInstanceNotInOpenStackError")
	DatabaseCannotManageClusterInstanceError               = constError("DatabaseCannotManageClusterInstanceError")
	DatabaseOldInstanceFindError                           = constError("DatabaseOldInstanceFindError")
	DatabaseCannotMoveIPError                              = constError("DatabaseCannotMoveIPError")
	DatabaseIPFindError                                    = constError("DatabaseIPFindError")

	// Kubernetes Errors
	DatabaseKubernetesClusterInvalidError         = constError("DatabaseKubernetesClusterInvalid")
	DatabaseKubernetesApplicationNotFoundError    = constError("DatabaseKubernetesApplicationNotFound")
	DatabaseKubernetesApplicationInvalidPlanError = constError("DatabaseKubernetesApplicationInvalidPlan")
	DatabaseKubernetesClusterDuplicateError       = constError("DatabaseKubernetesClusterDuplicate")
	DatabaseKubernetesClusterNotFoundError        = constError("DatabaseKubernetesClusterNotFound")
	DatabaseKubernetesNodeNotFoundError           = constError("DatabaseKubernetesNodeNotFound")

	DatabaseClusterPoolNotFoundError                       = constError("DatabaseClusterPoolNotFound")
	DatabaseClusterPoolInstanceNotFoundError               = constError("DatabaseClusterPoolInstanceNotFound")
	DatabaseClusterPoolInstanceDeleteFailedError           = constError("DatabaseClusterPoolInstanceDeleteFailed")
	DatabaseClusterPoolNoSufficientInstancesAvailableError = constError("DatabaseClusterPoolNoSufficientInstancesAvailable")

	DatabaseListingAccountsError              = constError("DatabaseListingAccountsError")
	DatabaseListingMembershipsError           = constError("DatabaseListingMembershipsError")
	DatabaseMembershipCannotDeleteError       = constError("DatabaseMembershipCannotDeleteError")
	DatabaseMembershipsGrantAccessError       = constError("DatabaseMembershipsGrantAccessError")
	DatabaseMembershipsInvalidInvitationError = constError("DatabaseMembershipsInvalidInvitationError")
	DatabaseMembershipsInvalidStatusError     = constError("DatabaseMembershipsInvalidStatusError")
	DatabaseMembershipsNotFoundError          = constError("DatabaseMembershipsNotFoundError")
	DatabaseMembershipsSuspendedError         = constError("DatabaseMembershipsSuspendedError")

	DatabaseLoadBalancerSaveError      = constError("DatabaseLoadBalancerSaveError")
	DatabaseLoadBalancerDeleteError    = constError("DatabaseLoadBalancerDeleteError")
	DatabaseLoadBalancerUpdateError    = constError("DatabaseLoadBalancerUpdateError")
	DatabaseLoadBalancerDuplicateError = constError("DatabaseLoadBalancerDuplicateError")
	DatabaseLoadBalancerExistsError    = constError("DatabaseLoadBalancerExistsError")
	DatabaseLoadBalancerNotFoundError  = constError("DatabaseLoadBalancerNotFoundError")

	DatabaseNetworksListError              = constError("DatabaseNetworksListError")
	DatabaseNetworkCreateError             = constError("DatabaseNetworkCreateError")
	DatabaseNetworkExistsError             = constError("DatabaseNetworkExistsError")
	DatabaseNetworkDeleteLastError         = constError("DatabaseNetworkDeleteLastError")
	DatabaseNetworkDeleteWithInstanceError = constError("DatabaseNetworkDeleteWithInstanceError")
	DatabaseNetworkInUseByVolumes          = constError("DatabaseNetworkInUseByVolumes")
	DatabaseNetworkDuplicateNameError      = constError("DatabaseNetworkDuplicateNameError")
	DatabaseNetworkLookupError             = constError("DatabaseNetworkLookupError")
	DatabaseNetworkNotFoundError           = constError("DatabaseNetworkNotFoundError")
	DatabaseNetworkSaveError               = constError("DatabaseNetworkSaveError")

	DatabasePrivateIPFromPublicIPError = constError("DatabasePrivateIPFromPublicIPError")

	DatabaseQuotaNotFoundError   = constError("DatabaseQuotaNotFoundError")
	DatabaseQuotaUpdateError     = constError("DatabaseQuotaUpdateError")
	QuotaLimitReachedError       = constError("QuotaLimitReachedError")
	DatabaseServiceNotFoundError = constError("DatabaseServiceNotFoundError")
	DatabaseSizeNotFoundError    = constError("DatabaseServiceNotFoundError")
	DatabaseSizesListError       = constError("DatabaseSizesListError")

	DatabaseSnapshotCannotDeleteInUseError      = constError("DatabaseSnapshotCannotDeleteInUseError")
	DatabaseSnapshotCannotReplaceError          = constError("DatabaseSnapshotCannotReplaceError")
	DatabaseSnapshotCreateError                 = constError("DatabaseSnapshotCreateError")
	DatabaseSnapshotCreateInstanceNotFoundError = constError("DatabaseSnapshotCreateInstanceNotFoundError")
	DatabaseSnapshotCreateAlreadyInProcessError = constError("DatabaseSnapshotCreateAlreadyInProcessError")
	DatabaseSnapshotNotFoundError               = constError("DatabaseSnapshotNotFoundError")
	DatabaseSnapshotsListError                  = constError("DatabaseSnapshotsListError")

	DatabaseSSHKeyDestroyError       = constError("DatabaseSSHKeyDestroyError")
	DatabaseSSHKeyCreateError        = constError("DatabaseSSHKeyCreateError")
	DatabaseSSHKeyUpdateError        = constError("DatabaseSSHKeyUpdateError")
	DatabaseSSHKeyDuplicateNameError = constError("DatabaseSSHKeyDuplicateNameError")
	DatabaseSSHKeyNotFoundError      = constError("DatabaseSSHKeyNotFoundError")
	SSHKeyDuplicateError             = constError("SSHKeyDuplicateError")

	DatabaseTeamCannotDeleteError      = constError("DatabaseTeamCannotDeleteError")
	DatabaseTeamCreateError            = constError("DatabaseTeamCreateError")
	DatabaseTeamListingError           = constError("DatabaseTeamListingError")
	DatabaseTeamMembershipCreateError  = constError("DatabaseTeamMembershipCreateError")
	DatabaseTeamNotFoundError          = constError("DatabaseTeamNotFoundError")
	DatabaseTemplateDestroyError       = constError("DatabaseTemplateDestroyError")
	DatabaseTemplateNotFoundError      = constError("DatabaseTemplateNotFoundError")
	DatabaseTemplateUpdateError        = constError("DatabaseTemplateUpdateError")
	DatabaseTemplateWouldConflictError = constError("DatabaseTemplateWouldConflictError")
	DatabaseImageIDInvalidError        = constError("DatabaseImageIDInvalidError")

	DatabaseUserAlreadyExistsError          = constError("DatabaseUserAlreadyExistsError")
	DatabaseUserNewError                    = constError("DatabaseUserNewError")
	DatabaseUserConfirmedError              = constError("DatabaseUserConfirmedError")
	DatabaseUserSuspendedError              = constError("DatabaseUserSuspendedError")
	DatabaseUserLoginFailedError            = constError("DatabaseUserLoginFailedError")
	DatabaseUserNoChangeStatusError         = constError("DatabaseUserNoChangeStatusError")
	DatabaseUserNotFoundError               = constError("DatabaseUserNotFoundError")
	DatabaseUserPasswordInvalidError        = constError("DatabaseUserPasswordInvalidError")
	DatabaseUserPasswordSecuringFailedError = constError("DatabaseUserPasswordSecuringFailedError")
	DatabaseUserUpdateError                 = constError("DatabaseUserUpdateError")
	DatabaseCreatingUserError               = constError("DatabaseCreatingUserError")

	DatabaseVolumeIDInvalidError                 = constError("DatabaseVolumeIDInvalidError")
	DatabaseVolumeDuplicateNameError             = constError("DatabaseVolumeDuplicateNameError")
	DatabaseVolumeCannotMultipleAttachError      = constError("DatabaseVolumeCannotMultipleAttachError")
	DatabaseVolumeStillAttachedCannotResizeError = constError("DatabaseVolumeStillAttachedCannotResizeError")
	DatabaseVolumeNotAttachedError               = constError("DatabaseVolumeNotAttachedError")
	DatabaseVolumeNotFoundError                  = constError("DatabaseVolumeNotFoundError")
	DatabaseVolumeDeleteFailedError              = constError("DatabaseVolumeDeleteFailedError")

	DatabaseWebhookDestroyError       = constError("DatabaseWebhookDestroyError")
	DatabaseWebhookNotFoundError      = constError("DatabaseWebhookNotFoundError")
	DatabaseWebhookUpdateError        = constError("DatabaseWebhookUpdateError")
	DatabaseWebhookWouldConflictError = constError("DatabaseWebhookWouldConflictError")

	OpenstackConnectionFailedError        = constError("OpenstackConnectionFailedError")
	OpenstackCreatingProjectError         = constError("OpenstackCreatingProjectError")
	OpenstackCreatingUserError            = constError("OpenstackCreatingUserError")
	OpenstackFirewallCreateError          = constError("OpenstackFirewallCreateError")
	OpenstackFirewallDestroyError         = constError("OpenstackFirewallDestroyError")
	OpenstackFirewallRuleDestroyError     = constError("OpenstackFirewallRuleDestroyError")
	OpenstackInstanceCreateError          = constError("OpenstackInstanceCreateError")
	OpenstackInstanceDestroyError         = constError("OpenstackInstanceDestroyError")
	OpenstackInstanceFindError            = constError("OpenstackInstanceFindError")
	OpenstackInstanceRebootError          = constError("OpenstackInstanceRebootError")
	OpenstackInstanceRebuildError         = constError("OpenstackInstanceRebuildError")
	OpenstackInstanceResizeError          = constError("OpenstackInstanceResizeError")
	OpenstackInstanceRestoreError         = constError("OpenstackInstanceRestoreError")
	OpenstackInstanceSetFirewallError     = constError("OpenstackInstanceSetFirewallError")
	OpenstackInstanceStartError           = constError("OpenstackInstanceStartError")
	OpenstackInstanceStopError            = constError("OpenstackInstanceStopError")
	OpenstackIPCreateError                = constError("OpenstackIPCreateError")
	OpenstackNetworkCreateFailedError     = constError("OpenstackNetworkCreateFailedError")
	OpenstackNnetworkDestroyFailedError   = constError("OpenstackNnetworkDestroyFailedError")
	OpenstackNetworkEnsureConfiguredError = constError("OpenstackNetworkEnsureConfiguredError")
	OpenstackPublicIPConnectError         = constError("OpenstackPublicIPConnectError")
	OpenstackQuotaApplyError              = constError("OpenstackQuotaApplyError")
	OpenstackSnapshotDestroyError         = constError("OpenstackSnapshotDestroyError")
	OpenstackSSHKeyUploadError            = constError("OpenstackSSHKeyUploadError")
	OpenstackProjectDestroyError          = constError("OpenstackProjectDestroyError")
	OpenstackProjectFindError             = constError("OpenstackProjectFindError")
	OpenstackUserDestroyError             = constError("OpenstackUserDestroyError")
	OpenstackURLGlanceError               = constError("OpenstackUrlGlanceError")
	OpenstackURLNovaError                 = constError("OpenstackURLNovaError")
	AuthenticationInvalidKeyError         = constError("AuthenticationInvalidKeyError")
	AuthenticationAccessDeniedError       = constError("AuthenticationAccessDeniedError")

	InstanceStateMustBeActiveOrShutoffError = constError("InstanceStateMustBeActiveOrShutoffError")
	MarshalingObjectsToJSONError            = constError("MarshalingObjectsToJsonError")
	NetworkCreateDefaultError               = constError("NetworkCreateDefaultError")
	NetworkDeleteDefaultError               = constError("NetworkDeleteDefaultError")
	ParameterTimeValueError                 = constError("ParameterTimeValueError")
	ParameterDateRangeTooLongError          = constError("ParameterDateRangeTooLongError")
	ParameterDNSRecordTypeError             = constError("ParameterDnsRecordTypeError")
	ParameterDNSRecordCnameApexError        = constError("ParameterDNSRecordCnameApexError")
	ParameterPublicKeyEmptyError            = constError("ParameterPublicKeyEmptyError")
	ParameterDateRangeError                 = constError("ParameterDateRangeError")
	ParameterIDMissingError                 = constError("ParameterIDMissingError")
	ParameterIDToIntegerError               = constError("ParameterIDToIntegerError")
	ParameterImageAndVolumeIDMissingError   = constError("ParameterImageAndVolumeIDMissingError")
	ParameterLabelInvalidError              = constError("ParameterLabelInvalidError")
	ParameterNameInvalidError               = constError("ParameterNameInvalidError")
	ParameterPrivateIPMissingError          = constError("ParameterPrivateIPMissingError")
	ParameterPublicIPMissingError           = constError("ParameterPublicIPMissingError")
	ParameterSizeMissingError               = constError("ParameterSizeMissingError")
	ParameterVolumeSizeIncorrectError       = constError("ParameterVolumeSizeIncorrectError")
	ParameterVolumeSizeMustIncreaseError    = constError("ParameterVolumeSizeMustIncreaseError")
	CannotResizeVolumeError                 = constError("CannotResizeVolumeError")
	ParameterSnapshotMissingError           = constError("ParameterSnapshotMissingError")
	ParameterSnapshotIncorrectFormatError   = constError("ParameterSnapshotIncorrectFormatError")
	ParameterStartPortMissingError          = constError("ParameterStartPortMissingError")
	DatabaseTemplateParseRequestError       = constError("DatabaseTemplateParseRequestError")
	ParameterValueMissingError              = constError("ParameterValueMissingError")

	OutOFCapacityError                           = constError("OutOFCapacityError")
	CannotGetConsoleError                        = constError("CannotGetConsoleError")
	DatabaseDNSDomainInvalidError                = constError("DatabaseDNSDomainInvalidError")
	DatabaseFirewallExistsError                  = constError("DatabaseFirewallExistsError")
	DatabaseKubernetesClusterNoPoolsError        = constError("DatabaseKubernetesClusterNoPoolsError")
	DatabaseKubernetesClusterInvalidVersionError = constError("DatabaseKubernetesClusterInvalidVersionError")
	DatabaseNamespacesListError                  = constError("DatabaseNamespacesListError")
	DatabaseNamespaceCreateError                 = constError("DatabaseNamespaceCreateError")
	DatabaseNamespaceExistsError                 = constError("DatabaseNamespaceExistsError")
	DatabaseNamespaceDeleteLastError             = constError("DatabaseNamespaceDeleteLastError")
	DatabaseNamespaceDeleteWithInstanceError     = constError("DatabaseNamespaceDeleteWithInstanceError")
	DatabaseNamespaceDuplicateNameError          = constError("DatabaseNamespaceDuplicateNameError")
	DatabaseNamespaceLookupError                 = constError("DatabaseNamespaceLookupError")
	DatabaseNamespaceNotFoundError               = constError("DatabaseNamespaceNotFoundError")
	DatabaseNamespaceSaveError                   = constError("DatabaseNamespaceSaveError")
	DatabaseQuotaLockFailedError                 = constError("DatabaseQuotaLockFailedError")
	DatabaseDiskImageNotFoundError               = constError("DatabaseDiskImageNotFoundError")
	DatabaseDiskImageNotImplementedError         = constError("DatabaseDiskImageNotImplementedError")
	DatabaseTemplateExistsError                  = constError("DatabaseTemplateExistsError")
	DatabaseTemplateSaveFailedError              = constError("DatabaseTemplateSaveFailedError")
	KubernetesClusterInvalidNameError            = constError("KubernetesClusterInvalidNameError")

	AccountNotEnabledIncCardError     = constError("AccountNotEnabledIncCardError")
	AccountNotEnabledWithoutCardError = constError("AccountNotEnabledWithoutCardError")

	UnknownError        = constError("UnknownError")
	AuthenticationError = constError("AuthenticationError")
	InternalServerError = constError("InternalServerError")
)

type constError string

func (err constError) Error() string {
	return string(err)
}

func (err constError) Is(target error) bool {
	ts := target.Error()
	es := string(err)
	return ts == es || strings.HasPrefix(ts, es+": ")
}

func (err constError) wrap(inner error) error {
	return wrapError{msg: string(err), err: inner}
}

type wrapError struct {
	err error
	msg string
}

func (err wrapError) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s: %v", err.msg, err.err)
	}
	return err.msg
}

func (err wrapError) Unwrap() error {
	return err.err
}

func (err wrapError) Is(target error) bool {
	return constError(err.msg).Is(target)
}

func decodeError(err error) error {
	var response map[string]interface{}
	var msg strings.Builder

	switch err := err.(type) {
	case *url.Error:
		if _, ok := err.Err.(net.Error); ok {
			err := fmt.Errorf("we found a problem connected against the api")
			return TimeoutError.wrap(err)
		}
	case net.Error:
		if err.Timeout() {
			err := fmt.Errorf("we found a network issue")
			return TimeoutError.wrap(err)
		}
		if _, ok := err.(*net.DNSError); ok {
			err := fmt.Errorf("we found a dns issue")
			return TimeoutError.wrap(err)
		}
	case wrapError:
		return err
	case HTTPError:
		errorData := err
		reason := []byte(errorData.Reason)

		if err := json.Unmarshal(reason, &response); err != nil {
			err := fmt.Errorf("failed to decode the response expected from the API - status: %s, code: %d, reason: %s", errorData.Status, errorData.Code, errorData.Reason)
			return ResponseDecodeFailedError.wrap(err)
		}

		if _, ok := response["status"].(float64); ok {
			if response["status"].(float64) == 500 {
				err := errors.New("internal Server Error")
				return InternalServerError.wrap(err)
			}
		}

		if response["result"] == "requires_authentication" {
			err := errors.New("authentication Error")
			return AuthenticationError.wrap(err)
		}

		if _, ok := response["reason"]; ok {
			msg.WriteString(response["reason"].(string))
			if _, ok := response["details"]; ok {
				msg.WriteString(", " + response["details"].(string))
			}
		}

		switch response["code"] {
		case "region_unavailable":
			err := errors.New(msg.String())
			return RegionUnavailableError.wrap(err)
		case "database_kubernetes_cluster_invalid":
			err := errors.New(msg.String())
			return DatabaseKubernetesClusterInvalidError.wrap(err)
		case "disabled_service":
			err := errors.New(msg.String())
			return DisabledServiceError.wrap(err)
		case "civostatsd_record_failed":
			err := errors.New(msg.String())
			return CivoStatsdRecordFailedError.wrap(err)
		case "authentication_failed":
			err := errors.New(msg.String())
			return AuthenticationFailedError.wrap(err)
		case "cannot_rescue_new_volume":
			err := errors.New(msg.String())
			return CannotRescueNewVolumeError.wrap(err)
		case "cannot_restore_new_volume":
			err := errors.New(msg.String())
			return CannotRestoreNewVolumeError.wrap(err)
		case "cannot_scale_already_rescaling_cluster":
			err := errors.New(msg.String())
			return CannotScaleAlreadyRescalingClusterError.wrap(err)
		case "database_account_destroy":
			err := errors.New(msg.String())
			return DatabaseAccountDestroyError.wrap(err)
		case "database_account_not_found":
			err := errors.New(msg.String())
			return DatabaseAccountNotFoundError.wrap(err)
		case "database_account_access_denied":
			err := errors.New(msg.String())
			return DatabaseAccountAccessDeniedError.wrap(err)
		case "database_creating_account":
			err := errors.New(msg.String())
			return DatabaseCreatingAccountError.wrap(err)
		case "database_updating_account":
			err := errors.New(msg.String())
			return DatabaseUpdatingAccountError.wrap(err)
		case "database_account_stats":
			err := errors.New(msg.String())
			return DatabaseAccountStatsError.wrap(err)
		case "database_action_listing":
			err := errors.New(msg.String())
			return DatabaseActionListingError.wrap(err)
		case "database_action_create":
			err := errors.New(msg.String())
			return DatabaseActionCreateError.wrap(err)
		case "database_api_key_create":
			err := errors.New(msg.String())
			return DatabaseAPIKeyCreateError.wrap(err)
		case "database_api_key_duplicate":
			err := errors.New(msg.String())
			return DatabaseAPIKeyDuplicateError.wrap(err)
		case "database_api_key_not_found":
			err := errors.New(msg.String())
			return DatabaseAPIKeyNotFoundError.wrap(err)
		case "database_api_key_destroy":
			err := errors.New(msg.String())
			return DatabaseAPIkeyDestroyError.wrap(err)
		case "database_audit_log_listing":
			err := errors.New(msg.String())
			return DatabaseAuditLogListingError.wrap(err)
		case "database_blueprint_not_found":
			err := errors.New(msg.String())
			return DatabaseBlueprintNotFoundError.wrap(err)
		case "database_blueprint_delete_failed":
			err := errors.New(msg.String())
			return DatabaseBlueprintDeleteFailedError.wrap(err)
		case "database_blueprint_create":
			err := errors.New(msg.String())
			return DatabaseBlueprintCreateError.wrap(err)
		case "database_blueprint_update":
			err := errors.New(msg.String())
			return DatabaseBlueprintUpdateError.wrap(err)
		case "parameter_empty_volume_id":
			err := errors.New(msg.String())
			return ParameterEmptyVolumeIDError.wrap(err)
		case "parameter_empty_openstack_volume_id":
			err := errors.New(msg.String())
			return ParameterEmptyOpenstackVolumeIDError.wrap(err)
		case "database_change_api_key":
			err := errors.New(msg.String())
			return DatabaseChangeAPIKeyError.wrap(err)
		case "database_charge_listing":
			err := errors.New(msg.String())
			return DatabaseChargeListingError.wrap(err)
		case "database_connection_failed":
			err := errors.New(msg.String())
			return DatabaseConnectionFailedError.wrap(err)
		case "database_dns_domain_create":
			err := errors.New(msg.String())
			return DatabaseDNSDomainCreateError.wrap(err)
		case "database_dns_domain_update":
			err := errors.New(msg.String())
			return DatabaseDNSDomainUpdateError.wrap(err)
		case "database_dns_domain_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseDNSDomainDuplicateNameError.wrap(err)
		case "database_dns_domain_not_found":
			err := errors.New(msg.String())
			return DatabaseDNSDomainNotFoundError.wrap(err)
		case "database_dns_record_create":
			err := errors.New(msg.String())
			return DatabaseDNSRecordCreateError.wrap(err)
		case "database_dns_record_not_found":
			err := errors.New(msg.String())
			return DatabaseDNSRecordNotFoundError.wrap(err)
		case "database_dns_record_update":
			err := errors.New(msg.String())
			return DatabaseDNSRecordUpdateError.wrap(err)
		case "database_firewall_create":
			err := errors.New(msg.String())
			return DatabaseFirewallCreateError.wrap(err)
		case "database_firewall_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseFirewallDuplicateNameError.wrap(err)
		case "database_firewall_rules_invalid_params":
			err := errors.New(msg.String())
			return DatabaseFirewallRulesInvalidParams.wrap(err)
		case "database_firewall_mismatch":
			err := errors.New(msg.String())
			return DatabaseFirewallMismatchError.wrap(err)
		case "database_firewall_not_found":
			err := errors.New(msg.String())
			return DatabaseFirewallNotFoundError.wrap(err)
		case "database_firewall_save_failed":
			err := errors.New(msg.String())
			return DatabaseFirewallSaveFailedError.wrap(err)
		case "database_firewall_delete_failed":
			err := errors.New(msg.String())
			return DatabaseFirewallDeleteFailedError.wrap(err)
		case "database_firewall_rule_create":
			err := errors.New(msg.String())
			return DatabaseFirewallRuleCreateError.wrap(err)
		case "database_firewall_rule_delete_failed":
			err := errors.New(msg.String())
			return DatabaseFirewallRuleDeleteFailedError.wrap(err)
		case "database_firewall_rules_find":
			err := errors.New(msg.String())
			return DatabaseFirewallRulesFindError.wrap(err)
		case "database_cannot_manage_cluster_instance":
			err := errors.New(msg.String())
			return DatabaseCannotManageClusterInstanceError.wrap(err)
		case "database_old_instance_find":
			err := errors.New(msg.String())
			return DatabaseOldInstanceFindError.wrap(err)
		case "database_cannot_move_ip":
			err := errors.New(msg.String())
			return DatabaseCannotMoveIPError.wrap(err)
		case "database_ip_find":
			err := errors.New(msg.String())
			return DatabaseIPFindError.wrap(err)
		case "database_listing_accounts":
			err := errors.New(msg.String())
			return DatabaseListingAccountsError.wrap(err)
		case "database_listing_firewalls":
			err := errors.New(msg.String())
			return DatabaseListingFirewallsError.wrap(err)
		case "database_listing_dns_domains":
			err := errors.New(msg.String())
			return DatabaseListingDNSDomainsError.wrap(err)
		case "database_listing_memberships":
			err := errors.New(msg.String())
			return DatabaseListingMembershipsError.wrap(err)
		case "database_loadbalancer_not_found":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerNotFoundError.wrap(err)
		case "database_loadbalancer_exists":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerExistsError.wrap(err)
		case "database_loadbalancer_save_failed":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerSaveError.wrap(err)
		case "database_loadbalancer_deleted_failed":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerDeleteError.wrap(err)
		case "database_loadbalancer_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerDuplicateError.wrap(err)
		case "database_loadbalancer_update_failed":
			err := errors.New(msg.String())
			return DatabaseLoadBalancerUpdateError.wrap(err)
		case "database_membership_cannot_delete":
			err := errors.New(msg.String())
			return DatabaseMembershipCannotDeleteError.wrap(err)
		case "database_memberships_grant_access":
			err := errors.New(msg.String())
			return DatabaseMembershipsGrantAccessError.wrap(err)
		case "database_memberships_invalid_invitation":
			err := errors.New(msg.String())
			return DatabaseMembershipsInvalidInvitationError.wrap(err)
		case "database_memberships_invalid_status":
			err := errors.New(msg.String())
			return DatabaseMembershipsInvalidStatusError.wrap(err)
		case "database_memberships_not_found":
			err := errors.New(msg.String())
			return DatabaseMembershipsNotFoundError.wrap(err)
		case "database_memberships_suspended":
			err := errors.New(msg.String())
			return DatabaseMembershipsSuspendedError.wrap(err)
		case "database_networks_list":
			err := errors.New(msg.String())
			return DatabaseNetworksListError.wrap(err)
		case "database_network_create":
			err := errors.New(msg.String())
			return DatabaseNetworkCreateError.wrap(err)
		case "database_network_exists":
			err := errors.New(msg.String())
			return DatabaseNetworkExistsError.wrap(err)
		case "database_network_delete_last":
			err := errors.New(msg.String())
			return DatabaseNetworkDeleteLastError.wrap(err)
		case "database_network_delete_with_instance":
			err := errors.New(msg.String())
			return DatabaseNetworkDeleteWithInstanceError.wrap(err)
		case "database_network_inuse_by_volumes":
			err := errors.New(msg.String())
			return DatabaseNetworkInUseByVolumes.wrap(err)
		case "database_network_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseNetworkDuplicateNameError.wrap(err)
		case "database_network_lookup":
			err := errors.New(msg.String())
			return DatabaseNetworkLookupError.wrap(err)
		case "database_network_not_found":
			err := errors.New(msg.String())
			return DatabaseNetworkNotFoundError.wrap(err)
		case "database_network_save":
			err := errors.New(msg.String())
			return DatabaseNetworkSaveError.wrap(err)
		case "database_private_ip_from_public_ip":
			err := errors.New(msg.String())
			return DatabasePrivateIPFromPublicIPError.wrap(err)
		case "database_quota_not_found":
			err := errors.New(msg.String())
			return DatabaseQuotaNotFoundError.wrap(err)
		case "database_quota_update":
			err := errors.New(msg.String())
			return DatabaseQuotaUpdateError.wrap(err)
		case "database_service_not_found":
			err := errors.New(msg.String())
			return DatabaseServiceNotFoundError.wrap(err)
		case "database_size_not_found":
			err := errors.New(msg.String())
			return DatabaseSizeNotFoundError.wrap(err)
		case "database_sizes_list":
			err := errors.New(msg.String())
			return DatabaseSizesListError.wrap(err)
		case "database_snapshot_cannot_delete_in_use":
			err := errors.New(msg.String())
			return DatabaseSnapshotCannotDeleteInUseError.wrap(err)
		case "database_snapshot_cannot_replace":
			err := errors.New(msg.String())
			return DatabaseSnapshotCannotReplaceError.wrap(err)
		case "database_snapshot_create":
			err := errors.New(msg.String())
			return DatabaseSnapshotCreateError.wrap(err)
		case "database_snapshot_create_instance_not_found":
			err := errors.New(msg.String())
			return DatabaseSnapshotCreateInstanceNotFoundError.wrap(err)
		case "database_snapshot_create_already_in_process":
			err := errors.New(msg.String())
			return DatabaseSnapshotCreateAlreadyInProcessError.wrap(err)
		case "database_snapshot_not_found":
			err := errors.New(msg.String())
			return DatabaseSnapshotNotFoundError.wrap(err)
		case "database_snapshots_list":
			err := errors.New(msg.String())
			return DatabaseSnapshotsListError.wrap(err)
		case "database_ssh_key_destroy":
			err := errors.New(msg.String())
			return DatabaseSSHKeyDestroyError.wrap(err)
		case "database_ssh_key_create":
			err := errors.New(msg.String())
			return DatabaseSSHKeyCreateError.wrap(err)
		case "database_ssh_key_update":
			err := errors.New(msg.String())
			return DatabaseSSHKeyUpdateError.wrap(err)
		case "database_ssh_key_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseSSHKeyDuplicateNameError.wrap(err)
		case "database_ssh_key_not_found":
			err := errors.New(msg.String())
			return DatabaseSSHKeyNotFoundError.wrap(err)
		case "database_team_cannot_delete":
			err := errors.New(msg.String())
			return DatabaseTeamCannotDeleteError.wrap(err)
		case "database_team_create":
			err := errors.New(msg.String())
			return DatabaseTeamCreateError.wrap(err)
		case "database_team_listing":
			err := errors.New(msg.String())
			return DatabaseTeamListingError.wrap(err)
		case "database_team_membership_create":
			err := errors.New(msg.String())
			return DatabaseTeamMembershipCreateError.wrap(err)
		case "database_team_not_found":
			err := errors.New(msg.String())
			return DatabaseTeamNotFoundError.wrap(err)
		case "database_template_destroy":
			err := errors.New(msg.String())
			return DatabaseTemplateDestroyError.wrap(err)
		case "database_template_not_found":
			err := errors.New(msg.String())
			return DatabaseTemplateNotFoundError.wrap(err)
		case "database_template_update":
			err := errors.New(msg.String())
			return DatabaseTemplateUpdateError.wrap(err)
		case "database_template_would_conflict":
			err := errors.New(msg.String())
			return DatabaseTemplateWouldConflictError.wrap(err)
		case "database_image_id_invalid":
			err := errors.New(msg.String())
			return DatabaseImageIDInvalidError.wrap(err)
		case "database_volume_id_invalid":
			err := errors.New(msg.String())
			return DatabaseVolumeIDInvalidError.wrap(err)
		case "database_user_already_exists":
			err := errors.New(msg.String())
			return DatabaseUserAlreadyExistsError.wrap(err)
		case "database_user_new":
			err := errors.New(msg.String())
			return DatabaseUserNewError.wrap(err)
		case "database_user_confirmed":
			err := errors.New(msg.String())
			return DatabaseUserConfirmedError.wrap(err)
		case "database_user_suspended":
			err := errors.New(msg.String())
			return DatabaseUserSuspendedError.wrap(err)
		case "database_user_login_failed":
			err := errors.New(msg.String())
			return DatabaseUserLoginFailedError.wrap(err)
		case "database_user_no_change_status":
			err := errors.New(msg.String())
			return DatabaseUserNoChangeStatusError.wrap(err)
		case "database_user_not_found":
			err := errors.New(msg.String())
			return DatabaseUserNotFoundError.wrap(err)
		case "database_user_password_invalid":
			err := errors.New(msg.String())
			return DatabaseUserPasswordInvalidError.wrap(err)
		case "database_user_password_securing_failed":
			err := errors.New(msg.String())
			return DatabaseUserPasswordSecuringFailedError.wrap(err)
		case "database_user_update":
			err := errors.New(msg.String())
			return DatabaseUserUpdateError.wrap(err)
		case "database_creating_user":
			err := errors.New(msg.String())
			return DatabaseCreatingUserError.wrap(err)
		case "database_volume_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseVolumeDuplicateNameError.wrap(err)
		case "database_volume_cannot_multiple_attach":
			err := errors.New(msg.String())
			return DatabaseVolumeCannotMultipleAttachError.wrap(err)
		case "database_volume_still_attached_cannot_resize":
			err := errors.New(msg.String())
			return DatabaseVolumeStillAttachedCannotResizeError.wrap(err)
		case "database_volume_not_attached":
			err := errors.New(msg.String())
			return DatabaseVolumeNotAttachedError.wrap(err)
		case "database_volume_not_found":
			err := errors.New(msg.String())
			return DatabaseVolumeNotFoundError.wrap(err)
		case "database_volume_delete_failed":
			err := errors.New(msg.String())
			return DatabaseVolumeDeleteFailedError.wrap(err)
		case "database_webhook_destroy":
			err := errors.New(msg.String())
			return DatabaseWebhookDestroyError.wrap(err)
		case "database_webhook_not_found":
			err := errors.New(msg.String())
			return DatabaseWebhookNotFoundError.wrap(err)
		case "database_webhook_update":
			err := errors.New(msg.String())
			return DatabaseWebhookUpdateError.wrap(err)
		case "database_webhook_would_conflict":
			err := errors.New(msg.String())
			return DatabaseWebhookWouldConflictError.wrap(err)
		case "openstack_connection_failed":
			err := errors.New(msg.String())
			return OpenstackConnectionFailedError.wrap(err)
		case "openstack_creating_project":
			err := errors.New(msg.String())
			return OpenstackCreatingProjectError.wrap(err)
		case "openstack_creating_user":
			err := errors.New(msg.String())
			return OpenstackCreatingUserError.wrap(err)
		case "openstack_firewall_create":
			err := errors.New(msg.String())
			return OpenstackFirewallCreateError.wrap(err)
		case "openstack_firewall_destroy":
			err := errors.New(msg.String())
			return OpenstackFirewallDestroyError.wrap(err)
		case "openstack_firewall_rule_destroy":
			err := errors.New(msg.String())
			return OpenstackFirewallRuleDestroyError.wrap(err)
		case "openstack_instance_create":
			err := errors.New(msg.String())
			return OpenstackInstanceCreateError.wrap(err)
		case "openstack_instance_destroy":
			err := errors.New(msg.String())
			return OpenstackInstanceDestroyError.wrap(err)
		case "openstack_instance_find":
			err := errors.New(msg.String())
			return OpenstackInstanceFindError.wrap(err)
		case "openstack_instance_reboot":
			err := errors.New(msg.String())
			return OpenstackInstanceRebootError.wrap(err)
		case "openstack_instance_rebuild":
			err := errors.New(msg.String())
			return OpenstackInstanceRebuildError.wrap(err)
		case "openstack_instance_resize":
			err := errors.New(msg.String())
			return OpenstackInstanceResizeError.wrap(err)
		case "openstack_instance_restore":
			err := errors.New(msg.String())
			return OpenstackInstanceRestoreError.wrap(err)
		case "openstack_instance_set_firewall":
			err := errors.New(msg.String())
			return OpenstackInstanceSetFirewallError.wrap(err)
		case "openstack_instance_start":
			err := errors.New(msg.String())
			return OpenstackInstanceStartError.wrap(err)
		case "openstack_instance_stop":
			err := errors.New(msg.String())
			return OpenstackInstanceStopError.wrap(err)
		case "openstack_ip_create":
			err := errors.New(msg.String())
			return OpenstackIPCreateError.wrap(err)
		case "openstack_network_create_failed":
			err := errors.New(msg.String())
			return OpenstackNetworkCreateFailedError.wrap(err)
		case "openstack_network_destroy_failed":
			err := errors.New(msg.String())
			return OpenstackNnetworkDestroyFailedError.wrap(err)
		case "openstack_network_ensure_configured":
			err := errors.New(msg.String())
			return OpenstackNetworkEnsureConfiguredError.wrap(err)
		case "openstack_public_ip_connect":
			err := errors.New(msg.String())
			return OpenstackPublicIPConnectError.wrap(err)
		case "openstack_quota_apply":
			err := errors.New(msg.String())
			return OpenstackQuotaApplyError.wrap(err)
		case "openstack_snapshot_destroy":
			err := errors.New(msg.String())
			return OpenstackSnapshotDestroyError.wrap(err)
		case "openstack_ssh_key_upload":
			err := errors.New(msg.String())
			return OpenstackSSHKeyUploadError.wrap(err)
		case "openstack_project_destroy":
			err := errors.New(msg.String())
			return OpenstackProjectDestroyError.wrap(err)
		case "openstack_project_find":
			err := errors.New(msg.String())
			return OpenstackProjectFindError.wrap(err)
		case "openstack_user_destroy":
			err := errors.New(msg.String())
			return OpenstackUserDestroyError.wrap(err)
		case "openstack_url_glance":
			err := errors.New(msg.String())
			return OpenstackURLGlanceError.wrap(err)
		case "openstack_url_nova":
			err := errors.New(msg.String())
			return OpenstackURLNovaError.wrap(err)
		case "authentication_invalid_key":
			err := errors.New(msg.String())
			return AuthenticationInvalidKeyError.wrap(err)
		case "authentication_access_denied":
			err := errors.New(msg.String())
			return AuthenticationAccessDeniedError.wrap(err)
		case "firewall_duplicate":
			err := errors.New(msg.String())
			return FirewallDuplicateError.wrap(err)
		case "instance_state_must_be_active_or_shutoff":
			err := errors.New(msg.String())
			return InstanceStateMustBeActiveOrShutoffError.wrap(err)
		case "marshaling_objects_to_json":
			err := errors.New(msg.String())
			return MarshalingObjectsToJSONError.wrap(err)
		case "network_create_default":
			err := errors.New(msg.String())
			return NetworkCreateDefaultError.wrap(err)
		case "network_delete_default":
			err := errors.New(msg.String())
			return NetworkDeleteDefaultError.wrap(err)
		case "parameter_time_value":
			err := errors.New(msg.String())
			return ParameterTimeValueError.wrap(err)
		case "parameter_date_range_too_long":
			err := errors.New(msg.String())
			return ParameterDateRangeTooLongError.wrap(err)
		case "parameter_dns_record_type":
			err := errors.New(msg.String())
			return ParameterDNSRecordTypeError.wrap(err)
		case "parameter_dns_record_cname_apex":
			err := errors.New(msg.String())
			return ParameterDNSRecordCnameApexError.wrap(err)
		case "parameter_public_key_empty":
			err := errors.New(msg.String())
			return ParameterPublicKeyEmptyError.wrap(err)
		case "parameter_date_range":
			err := errors.New(msg.String())
			return ParameterDateRangeError.wrap(err)
		case "parameter_id_missing":
			err := errors.New(msg.String())
			return ParameterIDMissingError.wrap(err)
		case "parameter_id_to_integer":
			err := errors.New(msg.String())
			return ParameterIDToIntegerError.wrap(err)
		case "parameter_image_and_volume_id_missing":
			err := errors.New(msg.String())
			return ParameterImageAndVolumeIDMissingError.wrap(err)
		case "parameter_label_invalid":
			err := errors.New(msg.String())
			return ParameterLabelInvalidError.wrap(err)
		case "parameter_name_invalid":
			err := errors.New(msg.String())
			return ParameterNameInvalidError.wrap(err)
		case "parameter_private_ip_missing":
			err := errors.New(msg.String())
			return ParameterPrivateIPMissingError.wrap(err)
		case "parameter_public_ip_missing":
			err := errors.New(msg.String())
			return ParameterPublicIPMissingError.wrap(err)
		case "parameter_size_missing":
			err := errors.New(msg.String())
			return ParameterSizeMissingError.wrap(err)
		case "parameter_volume_size_incorrect":
			err := errors.New(msg.String())
			return ParameterVolumeSizeIncorrectError.wrap(err)
		case "parameter_volume_size_must_increase":
			err := errors.New(msg.String())
			return ParameterVolumeSizeMustIncreaseError.wrap(err)
		case "parameter_snapshot_missing":
			err := errors.New(msg.String())
			return ParameterSnapshotMissingError.wrap(err)
		case "parameter_snapshot_incorrect_format":
			err := errors.New(msg.String())
			return ParameterSnapshotIncorrectFormatError.wrap(err)
		case "parameter_start_port_missing":
			err := errors.New(msg.String())
			return ParameterStartPortMissingError.wrap(err)
		case "database_template_parse_request":
			err := errors.New(msg.String())
			return DatabaseTemplateParseRequestError.wrap(err)
		case "parameter_value_missing":
			err := errors.New(msg.String())
			return ParameterValueMissingError.wrap(err)
		case "quota_limit_reached":
			err := errors.New(msg.String())
			return QuotaLimitReachedError.wrap(err)
		case "sshkey_duplicate":
			err := errors.New(msg.String())
			return SSHKeyDuplicateError.wrap(err)
		case "volume_invalid_size":
			err := errors.New(msg.String())
			return VolumeInvalidSizeError.wrap(err)
		case "cannot_resize_volume":
			err := errors.New(msg.String())
			return CannotResizeVolumeError.wrap(err)
		case "database_kubernetes_application_not_found":
			err := errors.New(msg.String())
			return DatabaseKubernetesApplicationNotFoundError.wrap(err)
		case "database_kubernetes_application_invalid_plan":
			err := errors.New(msg.String())
			return DatabaseKubernetesApplicationInvalidPlanError.wrap(err)
		case "database_kubernetes_cluster_duplicate":
			err := errors.New(msg.String())
			return DatabaseKubernetesClusterDuplicateError.wrap(err)
		case "database_kubernetes_cluster_not_found":
			err := errors.New(msg.String())
			return DatabaseKubernetesClusterNotFoundError.wrap(err)
		case "database_kubernetes_node_not_found":
			err := errors.New(msg.String())
			return DatabaseKubernetesNodeNotFoundError.wrap(err)
		case "database_cluster_pool_not_found":
			err := errors.New(msg.String())
			return DatabaseClusterPoolNotFoundError.wrap(err)
		case "database_cluster_pool_instance_not_found":
			err := errors.New(msg.String())
			return DatabaseClusterPoolInstanceNotFoundError.wrap(err)
		case "database_cluster_pool_instance_delete_failed":
			err := errors.New(msg.String())
			return DatabaseClusterPoolInstanceDeleteFailedError.wrap(err)
		case "database_cluster_pool_no_sufficient_instances_available":
			err := errors.New(msg.String())
			return DatabaseClusterPoolNoSufficientInstancesAvailableError.wrap(err)
		case "database_instance_already_in_rescue_state":
			err := errors.New(msg.String())
			return DatabaseInstanceAlreadyinRescueStateError.wrap(err)
		case "database_instance_build":
			err := errors.New(msg.String())
			return DatabaseInstanceBuildError.wrap(err)
		case "database_instance_build_multiple_with_existing_public_ip":
			err := errors.New(msg.String())
			return DatabaseInstanceBuildMultipleWithExistingPublicIPError.wrap(err)
		case "database_instance_create":
			err := errors.New(msg.String())
			return DatabaseInstanceCreateError.wrap(err)
		case "database_instance_snapshot_too_big":
			err := errors.New(msg.String())
			return DatabaseInstanceSnapshotTooBigError.wrap(err)
		case "instance_duplicate":
			err := errors.New(msg.String())
			return DatabaseInstanceDuplicateError.wrap(err)
		case "database_instance_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseInstanceDuplicateNameError.wrap(err)
		case "database_instance_list":
			err := errors.New(msg.String())
			return DatabaseInstanceListError.wrap(err)
		case "database_instance_find":
			err := errors.New(msg.String())
			return DatabaseInstanceNotFoundError.wrap(err)
		case "database_instance_not_in_openstack":
			err := errors.New(msg.String())
			return DatabaseInstanceNotInOpenStackError.wrap(err)
		case "account_not_enabled_inc_card":
			err := errors.New(msg.String())
			return AccountNotEnabledIncCardError.wrap(err)
		case "account_not_enabled_without_card":
			err := errors.New(msg.String())
			return AccountNotEnabledWithoutCardError.wrap(err)
		case "out_of_capacity":
			err := errors.New(msg.String())
			return OutOFCapacityError.wrap(err)
		case "cannot_get_console":
			err := errors.New(msg.String())
			return CannotGetConsoleError.wrap(err)
		case "database_dns_domain_invalid":
			err := errors.New(msg.String())
			return DatabaseDNSDomainInvalidError.wrap(err)

		case "database_firewall_exists":
			err := errors.New(msg.String())
			return DatabaseFirewallExistsError.wrap(err)
		case "database_kubernetes_cluster_no_pools":
			err := errors.New(msg.String())
			return DatabaseKubernetesClusterNoPoolsError.wrap(err)
		case "database_kubernetes_cluster_invalid_version":
			err := errors.New(msg.String())
			return DatabaseKubernetesClusterInvalidVersionError.wrap(err)
		case "database_namespaces_list":
			err := errors.New(msg.String())
			return DatabaseNamespacesListError.wrap(err)
		case "database_namespace_create":
			err := errors.New(msg.String())
			return DatabaseNamespaceCreateError.wrap(err)
		case "database_namespace_exists":
			err := errors.New(msg.String())
			return DatabaseNamespaceExistsError.wrap(err)
		case "database_namespace_delete_last":
			err := errors.New(msg.String())
			return DatabaseNamespaceDeleteLastError.wrap(err)
		case "database_namespace_delete_with_instance":
			err := errors.New(msg.String())
			return DatabaseNamespaceDeleteWithInstanceError.wrap(err)
		case "database_namespace_duplicate_name":
			err := errors.New(msg.String())
			return DatabaseNamespaceDuplicateNameError.wrap(err)
		case "database_namespace_lookup":
			err := errors.New(msg.String())
			return DatabaseNamespaceLookupError.wrap(err)
		case "database_namespace_not_found":
			err := errors.New(msg.String())
			return DatabaseNamespaceNotFoundError.wrap(err)
		case "database_namespace_save":
			err := errors.New(msg.String())
			return DatabaseNamespaceSaveError.wrap(err)
		case "database_quota_lock_failed":
			err := errors.New(msg.String())
			return DatabaseQuotaLockFailedError.wrap(err)
		case "database_disk_image_not_found":
			err := errors.New(msg.String())
			return DatabaseDiskImageNotFoundError.wrap(err)
		case "database_disk_image_not_implemented":
			err := errors.New(msg.String())
			return DatabaseDiskImageNotImplementedError.wrap(err)
		case "database_template_exists":
			err := errors.New(msg.String())
			return DatabaseTemplateExistsError.wrap(err)
		case "database_template_save_failed":
			err := errors.New(msg.String())
			return DatabaseTemplateSaveFailedError.wrap(err)
		case "kubernetes_cluster_invalid_name":
			err := errors.New(msg.String())
			return KubernetesClusterInvalidNameError.wrap(err)
		default:
			err := fmt.Errorf(fmt.Sprintf("Unknown error response - status: %s, code: %d, reason: %s", errorData.Status, errorData.Code, errorData.Reason))
			return CommonError.wrap(err)
		}
	}

	return UnknownError.wrap(err)
}
