package ecloud

import "github.com/ukfast/sdk-go/pkg/connection"

// GetVirtualMachineSliceResponseBody represents an API response body containing []VirtualMachine data
type GetVirtualMachineSliceResponseBody struct {
	connection.APIResponseBody
	Data []VirtualMachine `json:"data"`
}

// GetVirtualMachineResponseBody represents an API response body containing VirtualMachine data
type GetVirtualMachineResponseBody struct {
	connection.APIResponseBody
	Data VirtualMachine `json:"data"`
}

// GetTagSliceResponseBody represents an API response body containing []Tag data
type GetTagSliceResponseBody struct {
	connection.APIResponseBody
	Data []Tag `json:"data"`
}

// GetTagResponseBody represents an API response body containing Tag data
type GetTagResponseBody struct {
	connection.APIResponseBody
	Data Tag `json:"data"`
}

// GetSolutionSliceResponseBody represents an API response body containing []Solution data
type GetSolutionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents an API response body containing Solution data
type GetSolutionResponseBody struct {
	connection.APIResponseBody
	Data Solution `json:"data"`
}

// GetSiteSliceResponseBody represents an API response body containing []Site data
type GetSiteSliceResponseBody struct {
	connection.APIResponseBody
	Data []Site `json:"data"`
}

// GetSiteResponseBody represents an API response body containing Site data
type GetSiteResponseBody struct {
	connection.APIResponseBody
	Data Site `json:"data"`
}

// GetV1NetworkSliceResponseBody represents an API response body containing []V1Network data
type GetV1NetworkSliceResponseBody struct {
	connection.APIResponseBody
	Data []V1Network `json:"data"`
}

// GetV1NetworkResponseBody represents an API response body containing V1Network data
type GetV1NetworkResponseBody struct {
	connection.APIResponseBody
	Data V1Network `json:"data"`
}

// GetV1HostSliceResponseBody represents an API response body containing []V1Host data
type GetV1HostSliceResponseBody struct {
	connection.APIResponseBody
	Data []V1Host `json:"data"`
}

// GetV1HostResponseBody represents an API response body containing V1Host data
type GetV1HostResponseBody struct {
	connection.APIResponseBody
	Data V1Host `json:"data"`
}

// GetDatastoreSliceResponseBody represents an API response body containing []Datastore data
type GetDatastoreSliceResponseBody struct {
	connection.APIResponseBody
	Data []Datastore `json:"data"`
}

// GetDatastoreResponseBody represents an API response body containing Datastore data
type GetDatastoreResponseBody struct {
	connection.APIResponseBody
	Data Datastore `json:"data"`
}

// GetFirewallSliceResponseBody represents an API response body containing []Firewall data
type GetFirewallSliceResponseBody struct {
	connection.APIResponseBody
	Data []Firewall `json:"data"`
}

// GetFirewallResponseBody represents an API response body containing Firewall data
type GetFirewallResponseBody struct {
	connection.APIResponseBody
	Data Firewall `json:"data"`
}

// GetFirewallConfigSliceResponseBody represents an API response body containing []FirewallConfig data
type GetFirewallConfigSliceResponseBody struct {
	connection.APIResponseBody
	Data []FirewallConfig `json:"data"`
}

// GetFirewallConfigResponseBody represents an API response body containing FirewallConfig data
type GetFirewallConfigResponseBody struct {
	connection.APIResponseBody
	Data FirewallConfig `json:"data"`
}

// GetTemplateSliceResponseBody represents an API response body containing []Template data
type GetTemplateSliceResponseBody struct {
	connection.APIResponseBody
	Data []Template `json:"data"`
}

// GetTemplateResponseBody represents an API response body containing Template data
type GetTemplateResponseBody struct {
	connection.APIResponseBody
	Data Template `json:"data"`
}

// GetPodSliceResponseBody represents an API response body containing []Pod data
type GetPodSliceResponseBody struct {
	connection.APIResponseBody
	Data []Pod `json:"data"`
}

// GetPodResponseBody represents an API response body containing Pod data
type GetPodResponseBody struct {
	connection.APIResponseBody
	Data Pod `json:"data"`
}

// GetApplianceSliceResponseBody represents an API response body containing []Appliance data
type GetApplianceSliceResponseBody struct {
	connection.APIResponseBody
	Data []Appliance `json:"data"`
}

// GetApplianceResponseBody represents an API response body containing Appliance data
type GetApplianceResponseBody struct {
	connection.APIResponseBody
	Data Appliance `json:"data"`
}

// GetApplianceParameterSliceResponseBody represents an API response body containing []ApplianceParameter data
type GetApplianceParameterSliceResponseBody struct {
	connection.APIResponseBody
	Data []ApplianceParameter `json:"data"`
}

// GetApplianceParameterResponseBody represents an API response body containing ApplianceParameter data
type GetApplianceParameterResponseBody struct {
	connection.APIResponseBody
	Data ApplianceParameter `json:"data"`
}

// GetActiveDirectoryDomainSliceResponseBody represents an API response body containing []ActiveDirectoryDomain data
type GetActiveDirectoryDomainSliceResponseBody struct {
	connection.APIResponseBody
	Data []ActiveDirectoryDomain `json:"data"`
}

// GetActiveDirectoryDomainResponseBody represents an API response body containing ActiveDirectoryDomain data
type GetActiveDirectoryDomainResponseBody struct {
	connection.APIResponseBody
	Data ActiveDirectoryDomain `json:"data"`
}

// GetConsoleSessionSliceResponseBody represents an API response body containing []ConsoleSession data
type GetConsoleSessionSliceResponseBody struct {
	connection.APIResponseBody
	Data []ConsoleSession `json:"data"`
}

// GetConsoleSessionResponseBody represents an API response body containing ConsoleSession data
type GetConsoleSessionResponseBody struct {
	connection.APIResponseBody
	Data ConsoleSession `json:"data"`
}

// GetVPCSliceResponseBody represents an API response body containing []VPC data
type GetVPCSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPC `json:"data"`
}

// GetVPCResponseBody represents an API response body containing VPC data
type GetVPCResponseBody struct {
	connection.APIResponseBody
	Data VPC `json:"data"`
}

// GetAvailabilityZoneSliceResponseBody represents an API response body containing []AvailabilityZone data
type GetAvailabilityZoneSliceResponseBody struct {
	connection.APIResponseBody
	Data []AvailabilityZone `json:"data"`
}

// GetAvailabilityZoneResponseBody represents an API response body containing AvailabilityZone data
type GetAvailabilityZoneResponseBody struct {
	connection.APIResponseBody
	Data AvailabilityZone `json:"data"`
}

// GetNetworkSliceResponseBody represents an API response body containing []Network data
type GetNetworkSliceResponseBody struct {
	connection.APIResponseBody
	Data []Network `json:"data"`
}

// GetNetworkResponseBody represents an API response body containing Network data
type GetNetworkResponseBody struct {
	connection.APIResponseBody
	Data Network `json:"data"`
}

// GetDHCPSliceResponseBody represents an API response body containing []DHCP data
type GetDHCPSliceResponseBody struct {
	connection.APIResponseBody
	Data []DHCP `json:"data"`
}

// GetDHCPResponseBody represents an API response body containing DHCP data
type GetDHCPResponseBody struct {
	connection.APIResponseBody
	Data DHCP `json:"data"`
}

// GetVPNSliceResponseBody represents an API response body containing []VPN data
type GetVPNSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPN `json:"data"`
}

// GetVPNResponseBody represents an API response body containing VPN data
type GetVPNResponseBody struct {
	connection.APIResponseBody
	Data VPN `json:"data"`
}

// GetInstanceSliceResponseBody represents an API response body containing []Instance data
type GetInstanceSliceResponseBody struct {
	connection.APIResponseBody
	Data []Instance `json:"data"`
}

// GetInstanceResponseBody represents an API response body containing Instance data
type GetInstanceResponseBody struct {
	connection.APIResponseBody
	Data Instance `json:"data"`
}

// GetFloatingIPSliceResponseBody represents an API response body containing []FloatingIP data
type GetFloatingIPSliceResponseBody struct {
	connection.APIResponseBody
	Data []FloatingIP `json:"data"`
}

// GetFloatingIPResponseBody represents an API response body containing FloatingIP data
type GetFloatingIPResponseBody struct {
	connection.APIResponseBody
	Data FloatingIP `json:"data"`
}

// GetFirewallPolicySliceResponseBody represents an API response body containing []FirewallPolicy data
type GetFirewallPolicySliceResponseBody struct {
	connection.APIResponseBody
	Data []FirewallPolicy `json:"data"`
}

// GetFirewallPolicyResponseBody represents an API response body containing FirewallPolicy data
type GetFirewallPolicyResponseBody struct {
	connection.APIResponseBody
	Data FirewallPolicy `json:"data"`
}

// GetFirewallRuleSliceResponseBody represents an API response body containing []FirewallRule data
type GetFirewallRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []FirewallRule `json:"data"`
}

// GetFirewallRuleResponseBody represents an API response body containing FirewallRule data
type GetFirewallRuleResponseBody struct {
	connection.APIResponseBody
	Data FirewallRule `json:"data"`
}

// GetFirewallRulePortSliceResponseBody represents an API response body containing []FirewallRulePort data
type GetFirewallRulePortSliceResponseBody struct {
	connection.APIResponseBody
	Data []FirewallRulePort `json:"data"`
}

// GetFirewallRulePortResponseBody represents an API response body containing FirewallRulePort data
type GetFirewallRulePortResponseBody struct {
	connection.APIResponseBody
	Data FirewallRulePort `json:"data"`
}

// GetRegionSliceResponseBody represents an API response body containing []Region data
type GetRegionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Region `json:"data"`
}

// GetRegionResponseBody represents an API response body containing Region data
type GetRegionResponseBody struct {
	connection.APIResponseBody
	Data Region `json:"data"`
}

// GetRouterSliceResponseBody represents an API response body containing []Router data
type GetRouterSliceResponseBody struct {
	connection.APIResponseBody
	Data []Router `json:"data"`
}

// GetRouterResponseBody represents an API response body containing Router data
type GetRouterResponseBody struct {
	connection.APIResponseBody
	Data Router `json:"data"`
}

// GetCredentialSliceResponseBody represents an API response body containing []Credential data
type GetCredentialSliceResponseBody struct {
	connection.APIResponseBody
	Data []Credential `json:"data"`
}

// GetCredentialResponseBody represents an API response body containing Credential data
type GetCredentialResponseBody struct {
	connection.APIResponseBody
	Data Credential `json:"data"`
}

// GetVolumeSliceResponseBody represents an API response body containing []Volume data
type GetVolumeSliceResponseBody struct {
	connection.APIResponseBody
	Data []Volume `json:"data"`
}

// GetVolumeResponseBody represents an API response body containing Volume data
type GetVolumeResponseBody struct {
	connection.APIResponseBody
	Data Volume `json:"data"`
}

// GetNICSliceResponseBody represents an API response body containing []NIC data
type GetNICSliceResponseBody struct {
	connection.APIResponseBody
	Data []NIC `json:"data"`
}

// GetNICResponseBody represents an API response body containing NIC data
type GetNICResponseBody struct {
	connection.APIResponseBody
	Data NIC `json:"data"`
}

// GetRouterThroughputSliceResponseBody represents an API response body containing []RouterThroughput data
type GetRouterThroughputSliceResponseBody struct {
	connection.APIResponseBody
	Data []RouterThroughput `json:"data"`
}

// GetRouterThroughputResponseBody represents an API response body containing RouterThroughput data
type GetRouterThroughputResponseBody struct {
	connection.APIResponseBody
	Data RouterThroughput `json:"data"`
}

// GetDiscountPlanSliceResponseBody represents an API response body containing []DiscountPlan data
type GetDiscountPlanSliceResponseBody struct {
	connection.APIResponseBody
	Data []DiscountPlan `json:"data"`
}

// GetDiscountPlanResponseBody represents an API response body containing DiscountPlan data
type GetDiscountPlanResponseBody struct {
	connection.APIResponseBody
	Data DiscountPlan `json:"data"`
}

// GetBillingMetricSliceResponseBody represents an API response body containing []BillingMetric data
type GetBillingMetricSliceResponseBody struct {
	connection.APIResponseBody
	Data []BillingMetric `json:"data"`
}

// GetBillingMetricResponseBody represents an API response body containing BillingMetric data
type GetBillingMetricResponseBody struct {
	connection.APIResponseBody
	Data BillingMetric `json:"data"`
}

// GetImageSliceResponseBody represents an API response body containing []Image data
type GetImageSliceResponseBody struct {
	connection.APIResponseBody
	Data []Image `json:"data"`
}

// GetImageResponseBody represents an API response body containing Image data
type GetImageResponseBody struct {
	connection.APIResponseBody
	Data Image `json:"data"`
}

// GetImageParameterSliceResponseBody represents an API response body containing []ImageParameter data
type GetImageParameterSliceResponseBody struct {
	connection.APIResponseBody
	Data []ImageParameter `json:"data"`
}

// GetImageParameterResponseBody represents an API response body containing ImageParameter data
type GetImageParameterResponseBody struct {
	connection.APIResponseBody
	Data ImageParameter `json:"data"`
}

// GetImageMetadataSliceResponseBody represents an API response body containing []ImageMetadata data
type GetImageMetadataSliceResponseBody struct {
	connection.APIResponseBody
	Data []ImageMetadata `json:"data"`
}

// GetImageMetadataResponseBody represents an API response body containing ImageMetadata data
type GetImageMetadataResponseBody struct {
	connection.APIResponseBody
	Data ImageMetadata `json:"data"`
}

// GetHostGroupSliceResponseBody represents an API response body containing []HostGroup data
type GetHostGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []HostGroup `json:"data"`
}

// GetHostGroupResponseBody represents an API response body containing HostGroup data
type GetHostGroupResponseBody struct {
	connection.APIResponseBody
	Data HostGroup `json:"data"`
}

// GetHostSpecSliceResponseBody represents an API response body containing []HostSpec data
type GetHostSpecSliceResponseBody struct {
	connection.APIResponseBody
	Data []HostSpec `json:"data"`
}

// GetHostSpecResponseBody represents an API response body containing HostSpec data
type GetHostSpecResponseBody struct {
	connection.APIResponseBody
	Data HostSpec `json:"data"`
}

// GetHostSliceResponseBody represents an API response body containing []Host data
type GetHostSliceResponseBody struct {
	connection.APIResponseBody
	Data []Host `json:"data"`
}

// GetHostResponseBody represents an API response body containing Host data
type GetHostResponseBody struct {
	connection.APIResponseBody
	Data Host `json:"data"`
}

// GetSSHKeyPairSliceResponseBody represents an API response body containing []SSHKeyPair data
type GetSSHKeyPairSliceResponseBody struct {
	connection.APIResponseBody
	Data []SSHKeyPair `json:"data"`
}

// GetSSHKeyPairResponseBody represents an API response body containing SSHKeyPair data
type GetSSHKeyPairResponseBody struct {
	connection.APIResponseBody
	Data SSHKeyPair `json:"data"`
}

// GetTaskSliceResponseBody represents an API response body containing []Task data
type GetTaskSliceResponseBody struct {
	connection.APIResponseBody
	Data []Task `json:"data"`
}

// GetTaskResponseBody represents an API response body containing Task data
type GetTaskResponseBody struct {
	connection.APIResponseBody
	Data Task `json:"data"`
}

// GetTaskReferenceSliceResponseBody represents an API response body containing []TaskReference data
type GetTaskReferenceSliceResponseBody struct {
	connection.APIResponseBody
	Data []TaskReference `json:"data"`
}

// GetTaskReferenceResponseBody represents an API response body containing TaskReference data
type GetTaskReferenceResponseBody struct {
	connection.APIResponseBody
	Data TaskReference `json:"data"`
}

// GetNetworkPolicySliceResponseBody represents an API response body containing []NetworkPolicy data
type GetNetworkPolicySliceResponseBody struct {
	connection.APIResponseBody
	Data []NetworkPolicy `json:"data"`
}

// GetNetworkPolicyResponseBody represents an API response body containing NetworkPolicy data
type GetNetworkPolicyResponseBody struct {
	connection.APIResponseBody
	Data NetworkPolicy `json:"data"`
}

// GetNetworkRuleSliceResponseBody represents an API response body containing []NetworkRule data
type GetNetworkRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []NetworkRule `json:"data"`
}

// GetNetworkRuleResponseBody represents an API response body containing NetworkRule data
type GetNetworkRuleResponseBody struct {
	connection.APIResponseBody
	Data NetworkRule `json:"data"`
}

// GetNetworkRulePortSliceResponseBody represents an API response body containing []NetworkRulePort data
type GetNetworkRulePortSliceResponseBody struct {
	connection.APIResponseBody
	Data []NetworkRulePort `json:"data"`
}

// GetNetworkRulePortResponseBody represents an API response body containing NetworkRulePort data
type GetNetworkRulePortResponseBody struct {
	connection.APIResponseBody
	Data NetworkRulePort `json:"data"`
}

// GetVolumeGroupSliceResponseBody represents an API response body containing []VolumeGroup data
type GetVolumeGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []VolumeGroup `json:"data"`
}

// GetVolumeGroupResponseBody represents an API response body containing VolumeGroup data
type GetVolumeGroupResponseBody struct {
	connection.APIResponseBody
	Data VolumeGroup `json:"data"`
}

// GetVPNProfileGroupSliceResponseBody represents an API response body containing []VPNProfileGroup data
type GetVPNProfileGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPNProfileGroup `json:"data"`
}

// GetVPNProfileGroupResponseBody represents an API response body containing VPNProfileGroup data
type GetVPNProfileGroupResponseBody struct {
	connection.APIResponseBody
	Data VPNProfileGroup `json:"data"`
}

// GetVPNServiceSliceResponseBody represents an API response body containing []VPNService data
type GetVPNServiceSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPNService `json:"data"`
}

// GetVPNServiceResponseBody represents an API response body containing VPNService data
type GetVPNServiceResponseBody struct {
	connection.APIResponseBody
	Data VPNService `json:"data"`
}

// GetVPNEndpointSliceResponseBody represents an API response body containing []VPNEndpoint data
type GetVPNEndpointSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPNEndpoint `json:"data"`
}

// GetVPNEndpointResponseBody represents an API response body containing VPNEndpoint data
type GetVPNEndpointResponseBody struct {
	connection.APIResponseBody
	Data VPNEndpoint `json:"data"`
}

// GetVPNSessionSliceResponseBody represents an API response body containing []VPNSession data
type GetVPNSessionSliceResponseBody struct {
	connection.APIResponseBody
	Data []VPNSession `json:"data"`
}

// GetVPNSessionResponseBody represents an API response body containing VPNSession data
type GetVPNSessionResponseBody struct {
	connection.APIResponseBody
	Data VPNSession `json:"data"`
}

// GetVPNSessionPreSharedKeySliceResponseBody represents an API response body containing []VPNSessionPreSharedKey data
type GetVPNSessionPreSharedKeySliceResponseBody struct {
	connection.APIResponseBody
	Data []VPNSessionPreSharedKey `json:"data"`
}

// GetVPNSessionPreSharedKeyResponseBody represents an API response body containing VPNSessionPreSharedKey data
type GetVPNSessionPreSharedKeyResponseBody struct {
	connection.APIResponseBody
	Data VPNSessionPreSharedKey `json:"data"`
}

// GetLoadBalancerSliceResponseBody represents an API response body containing []LoadBalancer data
type GetLoadBalancerSliceResponseBody struct {
	connection.APIResponseBody
	Data []LoadBalancer `json:"data"`
}

// GetLoadBalancerResponseBody represents an API response body containing LoadBalancer data
type GetLoadBalancerResponseBody struct {
	connection.APIResponseBody
	Data LoadBalancer `json:"data"`
}

// GetLoadBalancerSpecSliceResponseBody represents an API response body containing []LoadBalancerSpec data
type GetLoadBalancerSpecSliceResponseBody struct {
	connection.APIResponseBody
	Data []LoadBalancerSpec `json:"data"`
}

// GetLoadBalancerSpecResponseBody represents an API response body containing LoadBalancerSpec data
type GetLoadBalancerSpecResponseBody struct {
	connection.APIResponseBody
	Data LoadBalancerSpec `json:"data"`
}

// GetVIPSliceResponseBody represents an API response body containing []VIP data
type GetVIPSliceResponseBody struct {
	connection.APIResponseBody
	Data []VIP `json:"data"`
}

// GetVIPResponseBody represents an API response body containing VIP data
type GetVIPResponseBody struct {
	connection.APIResponseBody
	Data VIP `json:"data"`
}
