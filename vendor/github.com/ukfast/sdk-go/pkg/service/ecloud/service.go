package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/account"
)

// ECloudService is an interface for managing eCloud
type ECloudService interface {
	// Virtual Machine
	GetVirtualMachines(parameters connection.APIRequestParameters) ([]VirtualMachine, error)
	GetVirtualMachinesPaginated(parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error)
	GetVirtualMachine(vmID int) (VirtualMachine, error)
	CreateVirtualMachine(req CreateVirtualMachineRequest) (int, error)
	PatchVirtualMachine(vmID int, patch PatchVirtualMachineRequest) error
	CloneVirtualMachine(vmID int, req CloneVirtualMachineRequest) (int, error)
	DeleteVirtualMachine(vmID int) error
	PowerOnVirtualMachine(vmID int) error
	PowerOffVirtualMachine(vmID int) error
	PowerResetVirtualMachine(vmID int) error
	PowerShutdownVirtualMachine(vmID int) error
	PowerRestartVirtualMachine(vmID int) error
	CreateVirtualMachineTemplate(vmID int, req CreateVirtualMachineTemplateRequest) error
	GetVirtualMachineTags(vmID int, parameters connection.APIRequestParameters) ([]Tag, error)
	GetVirtualMachineTagsPaginated(vmID int, parameters connection.APIRequestParameters) (*PaginatedTag, error)
	GetVirtualMachineTag(vmID int, tagKey string) (Tag, error)
	CreateVirtualMachineTag(vmID int, req CreateTagRequest) error
	PatchVirtualMachineTag(vmID int, tagKey string, patch PatchTagRequest) error
	DeleteVirtualMachineTag(vmID int, tagKey string) error
	CreateVirtualMachineConsoleSession(vmID int) (ConsoleSession, error)

	// Solution
	GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error)
	GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error)
	GetSolution(solutionID int) (Solution, error)
	PatchSolution(solutionID int, patch PatchSolutionRequest) (int, error)
	GetSolutionVirtualMachines(solutionID int, parameters connection.APIRequestParameters) ([]VirtualMachine, error)
	GetSolutionVirtualMachinesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error)
	GetSolutionSites(solutionID int, parameters connection.APIRequestParameters) ([]Site, error)
	GetSolutionSitesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedSite, error)
	GetSolutionDatastores(solutionID int, parameters connection.APIRequestParameters) ([]Datastore, error)
	GetSolutionDatastoresPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedDatastore, error)
	GetSolutionHosts(solutionID int, parameters connection.APIRequestParameters) ([]V1Host, error)
	GetSolutionHostsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedV1Host, error)
	GetSolutionNetworks(solutionID int, parameters connection.APIRequestParameters) ([]V1Network, error)
	GetSolutionNetworksPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedV1Network, error)
	GetSolutionFirewalls(solutionID int, parameters connection.APIRequestParameters) ([]Firewall, error)
	GetSolutionFirewallsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedFirewall, error)
	GetSolutionTemplates(solutionID int, parameters connection.APIRequestParameters) ([]Template, error)
	GetSolutionTemplatesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTemplate, error)
	GetSolutionTemplate(solutionID int, templateName string) (Template, error)
	DeleteSolutionTemplate(solutionID int, templateName string) error
	RenameSolutionTemplate(solutionID int, templateName string, req RenameTemplateRequest) error
	GetSolutionTags(solutionID int, parameters connection.APIRequestParameters) ([]Tag, error)
	GetSolutionTagsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTag, error)
	GetSolutionTag(solutionID int, tagKey string) (Tag, error)
	CreateSolutionTag(solutionID int, req CreateTagRequest) error
	PatchSolutionTag(solutionID int, tagKey string, patch PatchTagRequest) error
	DeleteSolutionTag(solutionID int, tagKey string) error

	// Site
	GetSites(parameters connection.APIRequestParameters) ([]Site, error)
	GetSitesPaginated(parameters connection.APIRequestParameters) (*PaginatedSite, error)
	GetSite(siteID int) (Site, error)

	// Host
	GetV1Hosts(parameters connection.APIRequestParameters) ([]V1Host, error)
	GetV1HostsPaginated(parameters connection.APIRequestParameters) (*PaginatedV1Host, error)
	GetV1Host(hostID int) (V1Host, error)

	// Datastore
	GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error)
	GetDatastoresPaginated(parameters connection.APIRequestParameters) (*PaginatedDatastore, error)
	GetDatastore(datastoreID int) (Datastore, error)

	// Firewall
	GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error)
	GetFirewallsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewall, error)
	GetFirewall(firewallID int) (Firewall, error)
	GetFirewallConfig(firewallID int) (FirewallConfig, error)

	// Pod
	GetPods(parameters connection.APIRequestParameters) ([]Pod, error)
	GetPodsPaginated(parameters connection.APIRequestParameters) (*PaginatedPod, error)
	GetPod(podID int) (Pod, error)
	GetPodTemplates(podID int, parameters connection.APIRequestParameters) ([]Template, error)
	GetPodTemplatesPaginated(podID int, parameters connection.APIRequestParameters) (*PaginatedTemplate, error)
	GetPodTemplate(podID int, templateName string) (Template, error)
	RenamePodTemplate(podID int, templateName string, req RenameTemplateRequest) error
	DeletePodTemplate(podID int, templateName string) error
	GetPodAppliances(podID int, parameters connection.APIRequestParameters) ([]Appliance, error)
	GetPodAppliancesPaginated(podID int, parameters connection.APIRequestParameters) (*PaginatedAppliance, error)
	PodConsoleAvailable(podID int) (bool, error)

	// Appliance
	GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error)
	GetAppliancesPaginated(parameters connection.APIRequestParameters) (*PaginatedAppliance, error)
	GetAppliance(applianceID string) (Appliance, error)
	GetApplianceParameters(applianceID string, reqParameters connection.APIRequestParameters) ([]ApplianceParameter, error)
	GetApplianceParametersPaginated(applianceID string, parameters connection.APIRequestParameters) (*PaginatedApplianceParameter, error)

	// Credit
	GetCredits(parameters connection.APIRequestParameters) ([]account.Credit, error)

	GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error)
	GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedActiveDirectoryDomain, error)
	GetActiveDirectoryDomain(domainID int) (ActiveDirectoryDomain, error)

	// V2

	// VPC
	GetVPCs(parameters connection.APIRequestParameters) ([]VPC, error)
	GetVPCsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPC, error)
	GetVPC(vpcID string) (VPC, error)
	CreateVPC(req CreateVPCRequest) (string, error)
	PatchVPC(vpcID string, patch PatchVPCRequest) error
	DeleteVPC(vpcID string) error
	DeployVPCDefaults(vpcID string) error
	GetVPCVolumes(vpcID string, parameters connection.APIRequestParameters) ([]Volume, error)
	GetVPCVolumesPaginated(vpcID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetVPCInstances(vpcID string, parameters connection.APIRequestParameters) ([]Instance, error)
	GetVPCInstancesPaginated(vpcID string, parameters connection.APIRequestParameters) (*PaginatedInstance, error)
	GetVPCTasks(vpcID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetVPCTasksPaginated(vpcID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Availability zone
	GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error)
	GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*PaginatedAvailabilityZone, error)
	GetAvailabilityZone(azID string) (AvailabilityZone, error)

	// Network
	GetNetworks(parameters connection.APIRequestParameters) ([]Network, error)
	GetNetworksPaginated(parameters connection.APIRequestParameters) (*PaginatedNetwork, error)
	GetNetwork(networkID string) (Network, error)
	CreateNetwork(req CreateNetworkRequest) (string, error)
	PatchNetwork(networkID string, patch PatchNetworkRequest) error
	DeleteNetwork(networkID string) error
	GetNetworkNICs(networkID string, parameters connection.APIRequestParameters) ([]NIC, error)
	GetNetworkNICsPaginated(networkID string, parameters connection.APIRequestParameters) (*PaginatedNIC, error)
	GetNetworkTasks(networkID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetNetworkTasksPaginated(networkID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// DHCP
	GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error)
	GetDHCPsPaginated(parameters connection.APIRequestParameters) (*PaginatedDHCP, error)
	GetDHCP(dhcpID string) (DHCP, error)
	GetDHCPTasks(dhcpID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetDHCPTasksPaginated(dhcpID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Instance
	GetInstances(parameters connection.APIRequestParameters) ([]Instance, error)
	GetInstancesPaginated(parameters connection.APIRequestParameters) (*PaginatedInstance, error)
	GetInstance(instanceID string) (Instance, error)
	CreateInstance(req CreateInstanceRequest) (string, error)
	PatchInstance(instanceID string, req PatchInstanceRequest) error
	DeleteInstance(instanceID string) error
	LockInstance(instanceID string) error
	UnlockInstance(instanceID string) error
	PowerOnInstance(instanceID string) (string, error)
	PowerOffInstance(instanceID string) (string, error)
	PowerResetInstance(instanceID string) (string, error)
	PowerShutdownInstance(instanceID string) (string, error)
	PowerRestartInstance(instanceID string) (string, error)
	MigrateInstance(instanceID string, req MigrateInstanceRequest) (string, error)
	CreateInstanceConsoleSession(instanceID string) (ConsoleSession, error)
	GetInstanceVolumes(instanceID string, parameters connection.APIRequestParameters) ([]Volume, error)
	GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetInstanceCredentials(instanceID string, parameters connection.APIRequestParameters) ([]Credential, error)
	GetInstanceCredentialsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedCredential, error)
	GetInstanceNICs(instanceID string, parameters connection.APIRequestParameters) ([]NIC, error)
	GetInstanceNICsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedNIC, error)
	GetInstanceTasks(instanceID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetInstanceTasksPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)
	AttachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error)
	DetachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error)
	GetInstanceFloatingIPs(instanceID string, parameters connection.APIRequestParameters) ([]FloatingIP, error)
	GetInstanceFloatingIPsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedFloatingIP, error)

	// Floating IP
	GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error)
	GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedFloatingIP, error)
	GetFloatingIP(fipID string) (FloatingIP, error)
	CreateFloatingIP(req CreateFloatingIPRequest) (TaskReference, error)
	PatchFloatingIP(fipID string, req PatchFloatingIPRequest) (TaskReference, error)
	DeleteFloatingIP(fipID string) (string, error)
	AssignFloatingIP(fipID string, req AssignFloatingIPRequest) (string, error)
	UnassignFloatingIP(fipID string) (string, error)
	GetFloatingIPTasks(fipID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetFloatingIPTasksPaginated(fipID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Firewall Policy
	GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error)
	GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error)
	GetFirewallPolicy(policyID string) (FirewallPolicy, error)
	CreateFirewallPolicy(req CreateFirewallPolicyRequest) (TaskReference, error)
	PatchFirewallPolicy(policyID string, req PatchFirewallPolicyRequest) (TaskReference, error)
	DeleteFirewallPolicy(policyID string) (string, error)
	GetFirewallPolicyFirewallRules(policyID string, parameters connection.APIRequestParameters) ([]FirewallRule, error)
	GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error)
	GetFirewallPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetFirewallPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Firewall Rule
	GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error)
	GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error)
	GetFirewallRule(ruleID string) (FirewallRule, error)
	CreateFirewallRule(req CreateFirewallRuleRequest) (TaskReference, error)
	PatchFirewallRule(ruleID string, req PatchFirewallRuleRequest) (TaskReference, error)
	DeleteFirewallRule(ruleID string) (string, error)
	GetFirewallRuleFirewallRulePorts(firewallRuleID string, parameters connection.APIRequestParameters) ([]FirewallRulePort, error)
	GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error)

	// Firewall Rule Ports
	GetFirewallRulePorts(parameters connection.APIRequestParameters) ([]FirewallRulePort, error)
	GetFirewallRulePortsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error)
	GetFirewallRulePort(ruleID string) (FirewallRulePort, error)
	CreateFirewallRulePort(req CreateFirewallRulePortRequest) (TaskReference, error)
	PatchFirewallRulePort(ruleID string, req PatchFirewallRulePortRequest) (TaskReference, error)
	DeleteFirewallRulePort(ruleID string) (string, error)

	// Router
	GetRouters(parameters connection.APIRequestParameters) ([]Router, error)
	GetRoutersPaginated(parameters connection.APIRequestParameters) (*PaginatedRouter, error)
	GetRouter(routerID string) (Router, error)
	CreateRouter(req CreateRouterRequest) (string, error)
	PatchRouter(routerID string, patch PatchRouterRequest) error
	DeleteRouter(routerID string) error
	GetRouterFirewallPolicies(routerID string, parameters connection.APIRequestParameters) ([]FirewallPolicy, error)
	GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error)
	GetRouterNetworks(routerID string, parameters connection.APIRequestParameters) ([]Network, error)
	GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedNetwork, error)
	GetRouterVPNs(routerID string, parameters connection.APIRequestParameters) ([]VPN, error)
	GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedVPN, error)
	DeployRouterDefaultFirewallPolicies(routerID string) error
	GetRouterTasks(routerID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetRouterTasksPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Region
	GetRegions(parameters connection.APIRequestParameters) ([]Region, error)
	GetRegionsPaginated(parameters connection.APIRequestParameters) (*PaginatedRegion, error)
	GetRegion(regionID string) (Region, error)

	// Volumes
	GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error)
	GetVolumesPaginated(parameters connection.APIRequestParameters) (*PaginatedVolume, error)
	GetVolume(volumeID string) (Volume, error)
	CreateVolume(req CreateVolumeRequest) (TaskReference, error)
	PatchVolume(volumeID string, patch PatchVolumeRequest) (TaskReference, error)
	DeleteVolume(volumeID string) (string, error)
	GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error)
	GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*PaginatedInstance, error)
	GetVolumeTasks(volumeID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetVolumeTasksPaginated(volumeID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// NICs
	GetNICs(parameters connection.APIRequestParameters) ([]NIC, error)
	GetNICsPaginated(parameters connection.APIRequestParameters) (*PaginatedNIC, error)
	GetNIC(nicID string) (NIC, error)
	GetNICTasks(nicID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetNICTasksPaginated(nicID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Billing metrics
	GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error)
	GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingMetric, error)
	GetBillingMetric(metricID string) (BillingMetric, error)

	// Router throughputs
	GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error)
	GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*PaginatedRouterThroughput, error)
	GetRouterThroughput(metricID string) (RouterThroughput, error)

	// Image
	GetImages(parameters connection.APIRequestParameters) ([]Image, error)
	GetImagesPaginated(parameters connection.APIRequestParameters) (*PaginatedImage, error)
	GetImage(imageID string) (Image, error)
	GetImageParameters(imageID string, parameters connection.APIRequestParameters) ([]ImageParameter, error)
	GetImageParametersPaginated(imageID string, parameters connection.APIRequestParameters) (*PaginatedImageParameter, error)
	GetImageMetadata(imageID string, parameters connection.APIRequestParameters) ([]ImageMetadata, error)
	GetImageMetadataPaginated(imageID string, parameters connection.APIRequestParameters) (*PaginatedImageMetadata, error)

	// HostSpecs
	GetHostSpecs(parameters connection.APIRequestParameters) ([]HostSpec, error)
	GetHostSpecsPaginated(parameters connection.APIRequestParameters) (*PaginatedHostSpec, error)
	GetHostSpec(specID string) (HostSpec, error)

	// HostGroups
	GetHostGroups(parameters connection.APIRequestParameters) ([]HostGroup, error)
	GetHostGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedHostGroup, error)
	GetHostGroup(hostGroupID string) (HostGroup, error)
	CreateHostGroup(req CreateHostGroupRequest) (TaskReference, error)
	PatchHostGroup(hostGroupID string, patch PatchHostGroupRequest) (TaskReference, error)
	DeleteHostGroup(hostGroupID string) (string, error)
	GetHostGroupTasks(hostGroupID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetHostGroupTasksPaginated(hostGroupID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Hosts
	GetHosts(parameters connection.APIRequestParameters) ([]Host, error)
	GetHostsPaginated(parameters connection.APIRequestParameters) (*PaginatedHost, error)
	GetHost(hostID string) (Host, error)
	CreateHost(req CreateHostRequest) (TaskReference, error)
	PatchHost(hostID string, patch PatchHostRequest) (TaskReference, error)
	DeleteHost(hostID string) (string, error)
	GetHostTasks(hostID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetHostTasksPaginated(hostID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// SSHKeyPairs
	GetSSHKeyPairs(parameters connection.APIRequestParameters) ([]SSHKeyPair, error)
	GetSSHKeyPairsPaginated(parameters connection.APIRequestParameters) (*PaginatedSSHKeyPair, error)
	GetSSHKeyPair(keypairID string) (SSHKeyPair, error)
	CreateSSHKeyPair(req CreateSSHKeyPairRequest) (string, error)
	PatchSSHKeyPair(keypairID string, patch PatchSSHKeyPairRequest) error
	DeleteSSHKeyPair(keypairID string) error

	// Tasks
	GetTasks(parameters connection.APIRequestParameters) ([]Task, error)
	GetTasksPaginated(parameters connection.APIRequestParameters) (*PaginatedTask, error)
	GetTask(taskID string) (Task, error)

	// Network Policy
	GetNetworkPolicies(parameters connection.APIRequestParameters) ([]NetworkPolicy, error)
	GetNetworkPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkPolicy, error)
	GetNetworkPolicy(policyID string) (NetworkPolicy, error)
	CreateNetworkPolicy(req CreateNetworkPolicyRequest) (TaskReference, error)
	PatchNetworkPolicy(policyID string, req PatchNetworkPolicyRequest) (TaskReference, error)
	DeleteNetworkPolicy(policyID string) (string, error)
	GetNetworkPolicyNetworkRules(policyID string, parameters connection.APIRequestParameters) ([]NetworkRule, error)
	GetNetworkPolicyNetworkRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedNetworkRule, error)
	GetNetworkPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error)
	GetNetworkPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedTask, error)

	// Network Rule
	GetNetworkRules(parameters connection.APIRequestParameters) ([]NetworkRule, error)
	GetNetworkRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkRule, error)
	GetNetworkRule(ruleID string) (NetworkRule, error)
	CreateNetworkRule(req CreateNetworkRuleRequest) (TaskReference, error)
	PatchNetworkRule(ruleID string, req PatchNetworkRuleRequest) (TaskReference, error)
	DeleteNetworkRule(ruleID string) (string, error)
	GetNetworkRuleNetworkRulePorts(networkRuleID string, parameters connection.APIRequestParameters) ([]NetworkRulePort, error)
	GetNetworkRuleNetworkRulePortsPaginated(networkRuleID string, parameters connection.APIRequestParameters) (*PaginatedNetworkRulePort, error)

	// Network Rule Ports
	GetNetworkRulePorts(parameters connection.APIRequestParameters) ([]NetworkRulePort, error)
	GetNetworkRulePortsPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkRulePort, error)
	GetNetworkRulePort(ruleID string) (NetworkRulePort, error)
	CreateNetworkRulePort(req CreateNetworkRulePortRequest) (TaskReference, error)
	PatchNetworkRulePort(ruleID string, req PatchNetworkRulePortRequest) (TaskReference, error)
	DeleteNetworkRulePort(ruleID string) (string, error)

	//Volume Groups
	GetVolumeGroups(parameters connection.APIRequestParameters) ([]VolumeGroup, error)
	GetVolumeGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedVolumeGroup, error)
	GetVolumeGroup(groupID string) (VolumeGroup, error)
	CreateVolumeGroup(req CreateVolumeGroupRequest) (TaskReference, error)
	PatchVolumeGroup(groupID string, patch PatchVolumeGroupRequest) (TaskReference, error)
	DeleteVolumeGroup(groupID string) (string, error)
	GetVolumeGroupVolumes(groupID string, parameters connection.APIRequestParameters) ([]Volume, error)
	GetVolumeGroupVolumesPaginated(groupID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error)

	// VPN Endpoint
	GetVPNEndpoints(parameters connection.APIRequestParameters) ([]VPNEndpoint, error)
	GetVPNEndpointsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNEndpoint, error)
	GetVPNEndpoint(endpointID string) (VPNEndpoint, error)
	CreateVPNEndpoint(req CreateVPNEndpointRequest) (TaskReference, error)
	PatchVPNEndpoint(endpointID string, req PatchVPNEndpointRequest) (TaskReference, error)
	DeleteVPNEndpoint(endpointID string) (string, error)

	// VPN Service
	GetVPNServices(parameters connection.APIRequestParameters) ([]VPNService, error)
	GetVPNServicesPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNService, error)
	GetVPNService(serviceID string) (VPNService, error)
	CreateVPNService(req CreateVPNServiceRequest) (TaskReference, error)
	PatchVPNService(serviceID string, req PatchVPNServiceRequest) (TaskReference, error)
	DeleteVPNService(serviceID string) (string, error)

	// VPN Session
	GetVPNSessions(parameters connection.APIRequestParameters) ([]VPNSession, error)
	GetVPNSessionsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNSession, error)
	GetVPNSession(sessionID string) (VPNSession, error)
	CreateVPNSession(req CreateVPNSessionRequest) (TaskReference, error)
	PatchVPNSession(sessionID string, req PatchVPNSessionRequest) (TaskReference, error)
	DeleteVPNSession(sessionID string) (string, error)
	GetVPNSessionPreSharedKey(sessionID string) (VPNSessionPreSharedKey, error)

	// VPN Profile Group
	GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error)
	GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNProfileGroup, error)
	GetVPNProfileGroup(groupID string) (VPNProfileGroup, error)

	// Load Balancer
	GetLoadBalancers(parameters connection.APIRequestParameters) ([]LoadBalancer, error)
	GetLoadBalancersPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancer, error)
	GetLoadBalancer(loadbalancerID string) (LoadBalancer, error)
	CreateLoadBalancer(req CreateLoadBalancerRequest) (TaskReference, error)
	PatchLoadBalancer(loadbalancerID string, req PatchLoadBalancerRequest) (TaskReference, error)
	DeleteLoadBalancer(loadbalancerID string) (string, error)

	// Load Balancer Spec
	GetLoadBalancerSpecs(parameters connection.APIRequestParameters) ([]LoadBalancerSpec, error)
	GetLoadBalancerSpecsPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancerSpec, error)
	GetLoadBalancerSpec(lbSpecID string) (LoadBalancerSpec, error)

	// VIP
	GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error)
	GetVIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedVIP, error)
	GetVIP(vipID string) (VIP, error)
	CreateVIP(req CreateVIPRequest) (TaskReference, error)
	PatchVIP(vipID string, patch PatchVIPRequest) (TaskReference, error)
	DeleteVIP(vipID string) (string, error)
}

// Service implements ECloudService for managing
// eCloud via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of eCloud Service
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
