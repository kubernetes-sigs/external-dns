package ecloud

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedVirtualMachine represents a paginated collection of VirtualMachine
type PaginatedVirtualMachine struct {
	*connection.PaginatedBase
	Items []VirtualMachine
}

// NewPaginatedVirtualMachine returns a pointer to an initialized PaginatedVirtualMachine struct
func NewPaginatedVirtualMachine(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VirtualMachine) *PaginatedVirtualMachine {
	return &PaginatedVirtualMachine{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTag represents a paginated collection of Tag
type PaginatedTag struct {
	*connection.PaginatedBase
	Items []Tag
}

// NewPaginatedTag returns a pointer to an initialized PaginatedTag struct
func NewPaginatedTag(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Tag) *PaginatedTag {
	return &PaginatedTag{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSolution represents a paginated collection of Solution
type PaginatedSolution struct {
	*connection.PaginatedBase
	Items []Solution
}

// NewPaginatedSolution returns a pointer to an initialized PaginatedSolution struct
func NewPaginatedSolution(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Solution) *PaginatedSolution {
	return &PaginatedSolution{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSite represents a paginated collection of Site
type PaginatedSite struct {
	*connection.PaginatedBase
	Items []Site
}

// NewPaginatedSite returns a pointer to an initialized PaginatedSite struct
func NewPaginatedSite(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Site) *PaginatedSite {
	return &PaginatedSite{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedV1Network represents a paginated collection of V1Network
type PaginatedV1Network struct {
	*connection.PaginatedBase
	Items []V1Network
}

// NewPaginatedV1Network returns a pointer to an initialized PaginatedV1Network struct
func NewPaginatedV1Network(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []V1Network) *PaginatedV1Network {
	return &PaginatedV1Network{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedV1Host represents a paginated collection of V1Host
type PaginatedV1Host struct {
	*connection.PaginatedBase
	Items []V1Host
}

// NewPaginatedV1Host returns a pointer to an initialized PaginatedV1Host struct
func NewPaginatedV1Host(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []V1Host) *PaginatedV1Host {
	return &PaginatedV1Host{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedDatastore represents a paginated collection of Datastore
type PaginatedDatastore struct {
	*connection.PaginatedBase
	Items []Datastore
}

// NewPaginatedDatastore returns a pointer to an initialized PaginatedDatastore struct
func NewPaginatedDatastore(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Datastore) *PaginatedDatastore {
	return &PaginatedDatastore{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFirewall represents a paginated collection of Firewall
type PaginatedFirewall struct {
	*connection.PaginatedBase
	Items []Firewall
}

// NewPaginatedFirewall returns a pointer to an initialized PaginatedFirewall struct
func NewPaginatedFirewall(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Firewall) *PaginatedFirewall {
	return &PaginatedFirewall{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTemplate represents a paginated collection of Template
type PaginatedTemplate struct {
	*connection.PaginatedBase
	Items []Template
}

// NewPaginatedTemplate returns a pointer to an initialized PaginatedTemplate struct
func NewPaginatedTemplate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Template) *PaginatedTemplate {
	return &PaginatedTemplate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedPod represents a paginated collection of Pod
type PaginatedPod struct {
	*connection.PaginatedBase
	Items []Pod
}

// NewPaginatedPod returns a pointer to an initialized PaginatedPod struct
func NewPaginatedPod(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Pod) *PaginatedPod {
	return &PaginatedPod{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAppliance represents a paginated collection of Appliance
type PaginatedAppliance struct {
	*connection.PaginatedBase
	Items []Appliance
}

// NewPaginatedAppliance returns a pointer to an initialized PaginatedAppliance struct
func NewPaginatedAppliance(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Appliance) *PaginatedAppliance {
	return &PaginatedAppliance{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedApplianceParameter represents a paginated collection of ApplianceParameter
type PaginatedApplianceParameter struct {
	*connection.PaginatedBase
	Items []ApplianceParameter
}

// NewPaginatedApplianceParameter returns a pointer to an initialized PaginatedApplianceParameter struct
func NewPaginatedApplianceParameter(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ApplianceParameter) *PaginatedApplianceParameter {
	return &PaginatedApplianceParameter{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedActiveDirectoryDomain represents a paginated collection of ActiveDirectoryDomain
type PaginatedActiveDirectoryDomain struct {
	*connection.PaginatedBase
	Items []ActiveDirectoryDomain
}

// NewPaginatedActiveDirectoryDomain returns a pointer to an initialized PaginatedActiveDirectoryDomain struct
func NewPaginatedActiveDirectoryDomain(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ActiveDirectoryDomain) *PaginatedActiveDirectoryDomain {
	return &PaginatedActiveDirectoryDomain{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPC represents a paginated collection of VPC
type PaginatedVPC struct {
	*connection.PaginatedBase
	Items []VPC
}

// NewPaginatedVPC returns a pointer to an initialized PaginatedVPC struct
func NewPaginatedVPC(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPC) *PaginatedVPC {
	return &PaginatedVPC{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAvailabilityZone represents a paginated collection of AvailabilityZone
type PaginatedAvailabilityZone struct {
	*connection.PaginatedBase
	Items []AvailabilityZone
}

// NewPaginatedAvailabilityZone returns a pointer to an initialized PaginatedAvailabilityZone struct
func NewPaginatedAvailabilityZone(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []AvailabilityZone) *PaginatedAvailabilityZone {
	return &PaginatedAvailabilityZone{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNetwork represents a paginated collection of Network
type PaginatedNetwork struct {
	*connection.PaginatedBase
	Items []Network
}

// NewPaginatedNetwork returns a pointer to an initialized PaginatedNetwork struct
func NewPaginatedNetwork(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Network) *PaginatedNetwork {
	return &PaginatedNetwork{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedDHCP represents a paginated collection of DHCP
type PaginatedDHCP struct {
	*connection.PaginatedBase
	Items []DHCP
}

// NewPaginatedDHCP returns a pointer to an initialized PaginatedDHCP struct
func NewPaginatedDHCP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []DHCP) *PaginatedDHCP {
	return &PaginatedDHCP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPN represents a paginated collection of VPN
type PaginatedVPN struct {
	*connection.PaginatedBase
	Items []VPN
}

// NewPaginatedVPN returns a pointer to an initialized PaginatedVPN struct
func NewPaginatedVPN(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPN) *PaginatedVPN {
	return &PaginatedVPN{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedInstance represents a paginated collection of Instance
type PaginatedInstance struct {
	*connection.PaginatedBase
	Items []Instance
}

// NewPaginatedInstance returns a pointer to an initialized PaginatedInstance struct
func NewPaginatedInstance(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Instance) *PaginatedInstance {
	return &PaginatedInstance{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFloatingIP represents a paginated collection of FloatingIP
type PaginatedFloatingIP struct {
	*connection.PaginatedBase
	Items []FloatingIP
}

// NewPaginatedFloatingIP returns a pointer to an initialized PaginatedFloatingIP struct
func NewPaginatedFloatingIP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []FloatingIP) *PaginatedFloatingIP {
	return &PaginatedFloatingIP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFirewallPolicy represents a paginated collection of FirewallPolicy
type PaginatedFirewallPolicy struct {
	*connection.PaginatedBase
	Items []FirewallPolicy
}

// NewPaginatedFirewallPolicy returns a pointer to an initialized PaginatedFirewallPolicy struct
func NewPaginatedFirewallPolicy(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []FirewallPolicy) *PaginatedFirewallPolicy {
	return &PaginatedFirewallPolicy{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFirewallRule represents a paginated collection of FirewallRule
type PaginatedFirewallRule struct {
	*connection.PaginatedBase
	Items []FirewallRule
}

// NewPaginatedFirewallRule returns a pointer to an initialized PaginatedFirewallRule struct
func NewPaginatedFirewallRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []FirewallRule) *PaginatedFirewallRule {
	return &PaginatedFirewallRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFirewallRulePort represents a paginated collection of FirewallRulePort
type PaginatedFirewallRulePort struct {
	*connection.PaginatedBase
	Items []FirewallRulePort
}

// NewPaginatedFirewallRulePort returns a pointer to an initialized PaginatedFirewallRulePort struct
func NewPaginatedFirewallRulePort(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []FirewallRulePort) *PaginatedFirewallRulePort {
	return &PaginatedFirewallRulePort{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRegion represents a paginated collection of Region
type PaginatedRegion struct {
	*connection.PaginatedBase
	Items []Region
}

// NewPaginatedRegion returns a pointer to an initialized PaginatedRegion struct
func NewPaginatedRegion(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Region) *PaginatedRegion {
	return &PaginatedRegion{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRouter represents a paginated collection of Router
type PaginatedRouter struct {
	*connection.PaginatedBase
	Items []Router
}

// NewPaginatedRouter returns a pointer to an initialized PaginatedRouter struct
func NewPaginatedRouter(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Router) *PaginatedRouter {
	return &PaginatedRouter{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCredential represents a paginated collection of Credential
type PaginatedCredential struct {
	*connection.PaginatedBase
	Items []Credential
}

// NewPaginatedCredential returns a pointer to an initialized PaginatedCredential struct
func NewPaginatedCredential(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Credential) *PaginatedCredential {
	return &PaginatedCredential{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVolume represents a paginated collection of Volume
type PaginatedVolume struct {
	*connection.PaginatedBase
	Items []Volume
}

// NewPaginatedVolume returns a pointer to an initialized PaginatedVolume struct
func NewPaginatedVolume(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Volume) *PaginatedVolume {
	return &PaginatedVolume{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNIC represents a paginated collection of NIC
type PaginatedNIC struct {
	*connection.PaginatedBase
	Items []NIC
}

// NewPaginatedNIC returns a pointer to an initialized PaginatedNIC struct
func NewPaginatedNIC(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []NIC) *PaginatedNIC {
	return &PaginatedNIC{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRouterThroughput represents a paginated collection of RouterThroughput
type PaginatedRouterThroughput struct {
	*connection.PaginatedBase
	Items []RouterThroughput
}

// NewPaginatedRouterThroughput returns a pointer to an initialized PaginatedRouterThroughput struct
func NewPaginatedRouterThroughput(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []RouterThroughput) *PaginatedRouterThroughput {
	return &PaginatedRouterThroughput{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedDiscountPlan represents a paginated collection of DiscountPlan
type PaginatedDiscountPlan struct {
	*connection.PaginatedBase
	Items []DiscountPlan
}

// NewPaginatedDiscountPlan returns a pointer to an initialized PaginatedDiscountPlan struct
func NewPaginatedDiscountPlan(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []DiscountPlan) *PaginatedDiscountPlan {
	return &PaginatedDiscountPlan{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedBillingMetric represents a paginated collection of BillingMetric
type PaginatedBillingMetric struct {
	*connection.PaginatedBase
	Items []BillingMetric
}

// NewPaginatedBillingMetric returns a pointer to an initialized PaginatedBillingMetric struct
func NewPaginatedBillingMetric(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []BillingMetric) *PaginatedBillingMetric {
	return &PaginatedBillingMetric{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedImage represents a paginated collection of Image
type PaginatedImage struct {
	*connection.PaginatedBase
	Items []Image
}

// NewPaginatedImage returns a pointer to an initialized PaginatedImage struct
func NewPaginatedImage(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Image) *PaginatedImage {
	return &PaginatedImage{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedImageParameter represents a paginated collection of ImageParameter
type PaginatedImageParameter struct {
	*connection.PaginatedBase
	Items []ImageParameter
}

// NewPaginatedImageParameter returns a pointer to an initialized PaginatedImageParameter struct
func NewPaginatedImageParameter(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ImageParameter) *PaginatedImageParameter {
	return &PaginatedImageParameter{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedImageMetadata represents a paginated collection of ImageMetadata
type PaginatedImageMetadata struct {
	*connection.PaginatedBase
	Items []ImageMetadata
}

// NewPaginatedImageMetadata returns a pointer to an initialized PaginatedImageMetadata struct
func NewPaginatedImageMetadata(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ImageMetadata) *PaginatedImageMetadata {
	return &PaginatedImageMetadata{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHostGroup represents a paginated collection of HostGroup
type PaginatedHostGroup struct {
	*connection.PaginatedBase
	Items []HostGroup
}

// NewPaginatedHostGroup returns a pointer to an initialized PaginatedHostGroup struct
func NewPaginatedHostGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []HostGroup) *PaginatedHostGroup {
	return &PaginatedHostGroup{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHostSpec represents a paginated collection of HostSpec
type PaginatedHostSpec struct {
	*connection.PaginatedBase
	Items []HostSpec
}

// NewPaginatedHostSpec returns a pointer to an initialized PaginatedHostSpec struct
func NewPaginatedHostSpec(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []HostSpec) *PaginatedHostSpec {
	return &PaginatedHostSpec{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHost represents a paginated collection of Host
type PaginatedHost struct {
	*connection.PaginatedBase
	Items []Host
}

// NewPaginatedHost returns a pointer to an initialized PaginatedHost struct
func NewPaginatedHost(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Host) *PaginatedHost {
	return &PaginatedHost{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSSHKeyPair represents a paginated collection of SSHKeyPair
type PaginatedSSHKeyPair struct {
	*connection.PaginatedBase
	Items []SSHKeyPair
}

// NewPaginatedSSHKeyPair returns a pointer to an initialized PaginatedSSHKeyPair struct
func NewPaginatedSSHKeyPair(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []SSHKeyPair) *PaginatedSSHKeyPair {
	return &PaginatedSSHKeyPair{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTask represents a paginated collection of Task
type PaginatedTask struct {
	*connection.PaginatedBase
	Items []Task
}

// NewPaginatedTask returns a pointer to an initialized PaginatedTask struct
func NewPaginatedTask(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Task) *PaginatedTask {
	return &PaginatedTask{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNetworkPolicy represents a paginated collection of NetworkPolicy
type PaginatedNetworkPolicy struct {
	*connection.PaginatedBase
	Items []NetworkPolicy
}

// NewPaginatedNetworkPolicy returns a pointer to an initialized PaginatedNetworkPolicy struct
func NewPaginatedNetworkPolicy(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []NetworkPolicy) *PaginatedNetworkPolicy {
	return &PaginatedNetworkPolicy{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNetworkRule represents a paginated collection of NetworkRule
type PaginatedNetworkRule struct {
	*connection.PaginatedBase
	Items []NetworkRule
}

// NewPaginatedNetworkRule returns a pointer to an initialized PaginatedNetworkRule struct
func NewPaginatedNetworkRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []NetworkRule) *PaginatedNetworkRule {
	return &PaginatedNetworkRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNetworkRulePort represents a paginated collection of NetworkRulePort
type PaginatedNetworkRulePort struct {
	*connection.PaginatedBase
	Items []NetworkRulePort
}

// NewPaginatedNetworkRulePort returns a pointer to an initialized PaginatedNetworkRulePort struct
func NewPaginatedNetworkRulePort(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []NetworkRulePort) *PaginatedNetworkRulePort {
	return &PaginatedNetworkRulePort{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVolumeGroup represents a paginated collection of VolumeGroup
type PaginatedVolumeGroup struct {
	*connection.PaginatedBase
	Items []VolumeGroup
}

// NewPaginatedVolumeGroup returns a pointer to an initialized PaginatedVolumeGroup struct
func NewPaginatedVolumeGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VolumeGroup) *PaginatedVolumeGroup {
	return &PaginatedVolumeGroup{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPNProfileGroup represents a paginated collection of VPNProfileGroup
type PaginatedVPNProfileGroup struct {
	*connection.PaginatedBase
	Items []VPNProfileGroup
}

// NewPaginatedVPNProfileGroup returns a pointer to an initialized PaginatedVPNProfileGroup struct
func NewPaginatedVPNProfileGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPNProfileGroup) *PaginatedVPNProfileGroup {
	return &PaginatedVPNProfileGroup{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPNService represents a paginated collection of VPNService
type PaginatedVPNService struct {
	*connection.PaginatedBase
	Items []VPNService
}

// NewPaginatedVPNService returns a pointer to an initialized PaginatedVPNService struct
func NewPaginatedVPNService(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPNService) *PaginatedVPNService {
	return &PaginatedVPNService{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPNEndpoint represents a paginated collection of VPNEndpoint
type PaginatedVPNEndpoint struct {
	*connection.PaginatedBase
	Items []VPNEndpoint
}

// NewPaginatedVPNEndpoint returns a pointer to an initialized PaginatedVPNEndpoint struct
func NewPaginatedVPNEndpoint(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPNEndpoint) *PaginatedVPNEndpoint {
	return &PaginatedVPNEndpoint{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPNSession represents a paginated collection of VPNSession
type PaginatedVPNSession struct {
	*connection.PaginatedBase
	Items []VPNSession
}

// NewPaginatedVPNSession returns a pointer to an initialized PaginatedVPNSession struct
func NewPaginatedVPNSession(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPNSession) *PaginatedVPNSession {
	return &PaginatedVPNSession{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedLoadBalancer represents a paginated collection of LoadBalancer
type PaginatedLoadBalancer struct {
	*connection.PaginatedBase
	Items []LoadBalancer
}

// NewPaginatedLoadBalancer returns a pointer to an initialized PaginatedLoadBalancer struct
func NewPaginatedLoadBalancer(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []LoadBalancer) *PaginatedLoadBalancer {
	return &PaginatedLoadBalancer{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedLoadBalancerSpec represents a paginated collection of LoadBalancerSpec
type PaginatedLoadBalancerSpec struct {
	*connection.PaginatedBase
	Items []LoadBalancerSpec
}

// NewPaginatedLoadBalancerSpec returns a pointer to an initialized PaginatedLoadBalancerSpec struct
func NewPaginatedLoadBalancerSpec(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []LoadBalancerSpec) *PaginatedLoadBalancerSpec {
	return &PaginatedLoadBalancerSpec{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVIP represents a paginated collection of VIP
type PaginatedVIP struct {
	*connection.PaginatedBase
	Items []VIP
}

// NewPaginatedVIP returns a pointer to an initialized PaginatedVIP struct
func NewPaginatedVIP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VIP) *PaginatedVIP {
	return &PaginatedVIP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
