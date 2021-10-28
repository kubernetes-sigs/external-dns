package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// PatchTagRequest represents an eCloud tag patch request
type PatchTagRequest struct {
	Value string `json:"value,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchTagRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateTagRequest represents a request to create an eCloud tag
type CreateTagRequest struct {
	connection.APIRequestBodyDefaultValidator

	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateTagRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateVirtualMachineRequest represents a request to create an eCloud virtual machine
type CreateVirtualMachineRequest struct {
	connection.APIRequestBodyDefaultValidator

	Environment      string `json:"environment" validate:"required"`
	Template         string `json:"template,omitempty"`
	ApplianceID      string `json:"appliance_id,omitempty"`
	TemplatePassword string `json:"template_password,omitempty"`
	// Count in Cores
	CPU int `json:"cpu" validate:"required"`
	// Size in GB
	RAM int `json:"ram" validate:"required"`
	// Size in GB
	HDD                     int                                    `json:"hdd,omitempty"`
	Disks                   []CreateVirtualMachineRequestDisk      `json:"hdd_disks,omitempty"`
	Name                    string                                 `json:"name,omitempty"`
	ComputerName            string                                 `json:"computername,omitempty"`
	Tags                    []CreateTagRequest                     `json:"tags,omitempty"`
	Backup                  bool                                   `json:"backup"`
	Support                 bool                                   `json:"support"`
	Monitoring              bool                                   `json:"monitoring"`
	MonitoringContacts      []int                                  `json:"monitoring_contacts"`
	SolutionID              int                                    `json:"solution_id,omitempty"`
	DatastoreID             int                                    `json:"datastore_id,omitempty"`
	SiteID                  int                                    `json:"site_id,omitempty"`
	NetworkID               int                                    `json:"network_id,omitempty"`
	ExternalIPRequired      bool                                   `json:"external_ip_required"`
	SSHKeys                 []string                               `json:"ssh_keys,omitempty"`
	Parameters              []CreateVirtualMachineRequestParameter `json:"parameters,omitempty"`
	Encrypt                 *bool                                  `json:"encrypt,omitempty"`
	Role                    string                                 `json:"role,omitempty"`
	BootstrapScript         string                                 `json:"bootstrap_script,omitempty"`
	ActiveDirectoryDomainID int                                    `json:"ad_domain_id,omitempty"`
	PodID                   int                                    `json:"pod_id,omitempty"`
}

// CreateVirtualMachineRequestDisk represents a request to create an eCloud virtual machine disk
type CreateVirtualMachineRequestDisk struct {
	Name string `json:"name,omitempty"`
	// Size in GB
	Capacity int `json:"capacity" validate:"required"`
}

// CreateVirtualMachineRequestParameter represents an eCloud virtual machine parameter
type CreateVirtualMachineRequestParameter struct {
	connection.APIRequestBodyDefaultValidator

	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateVirtualMachineRequest) Validate() *connection.ValidationError {
	if c.HDD == 0 && (c.Disks == nil || len(c.Disks) < 1) {
		return connection.NewValidationError("HDD or Disks must be provided")
	}

	if c.Template == "" && c.ApplianceID == "" {
		return connection.NewValidationError("Template or ApplianceID must be provided")
	}

	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchSolutionRequest represents an eCloud solution patch request
type PatchSolutionRequest struct {
	Name              *string `json:"name,omitempty"`
	EncryptionDefault *bool   `json:"encryption_default,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchSolutionRequest) Validate() *connection.ValidationError {
	return nil
}

// RenameTemplateRequest represents an eCloud template rename request
type RenameTemplateRequest struct {
	connection.APIRequestBodyDefaultValidator

	Destination string `json:"destination" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *RenameTemplateRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CloneVirtualMachineRequest represents a request to clone an eCloud virtual machine
type CloneVirtualMachineRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name string `json:"name" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CloneVirtualMachineRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchVirtualMachineRequest represents an eCloud virtual machine patch request
type PatchVirtualMachineRequest struct {
	Name *string `json:"name,omitempty"`
	// Count in Cores
	CPU int `json:"cpu,omitempty"`
	// Size in GB
	RAM int `json:"ram,omitempty"`
	// KV map of hard disks, key being hard disk name, value being size in GB
	Disks []PatchVirtualMachineRequestDisk `json:"hdd_disks,omitempty"`
	Role  string                           `json:"role,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchVirtualMachineRequest) Validate() *connection.ValidationError {
	return nil
}

type PatchVirtualMachineRequestDiskState string

const (
	PatchVirtualMachineRequestDiskStatePresent PatchVirtualMachineRequestDiskState = "present"
	PatchVirtualMachineRequestDiskStateAbsent  PatchVirtualMachineRequestDiskState = "absent"
)

func (s PatchVirtualMachineRequestDiskState) String() string {
	return string(s)
}

// PatchVirtualMachineRequestDisk represents an eCloud virtual machine patch request disk
type PatchVirtualMachineRequestDisk struct {
	UUID string `json:"uuid,omitempty"`
	// Size in GB
	Capacity int                                 `json:"capacity,omitempty"`
	State    PatchVirtualMachineRequestDiskState `json:"state,omitempty"`
}

// CreateVirtualMachineTemplateRequest represents a request to clone an eCloud virtual machine template
type CreateVirtualMachineTemplateRequest struct {
	connection.APIRequestBodyDefaultValidator

	TemplateName string       `json:"template_name" validate:"required"`
	TemplateType TemplateType `json:"template_type"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateVirtualMachineTemplateRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateVPCRequest represents a request to create a VPC
type CreateVPCRequest struct {
	Name               string `json:"name,omitempty"`
	RegionID           string `json:"region_id"`
	AdvancedNetworking *bool  `json:"advanced_networking,omitempty"`
}

// PatchVPCRequest represents a request to patch a VPC
type PatchVPCRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateNetworkRequest represents a request to create a network
type CreateNetworkRequest struct {
	Name     string `json:"name,omitempty"`
	RouterID string `json:"router_id"`
	Subnet   string `json:"subnet"`
}

// PatchNetworkRequest represents a request to patch a network
type PatchNetworkRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateRouterRequest represents a request to create a router
type CreateRouterRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	AvailabilityZoneID string `json:"availability_zone_id"`
	RouterThroughputID string `json:"router_throughput_id,omitempty"`
}

// PatchRouterRequest represents a request to patch a router
type PatchRouterRequest struct {
	Name               string `json:"name,omitempty"`
	RouterThroughputID string `json:"router_throughput_id,omitempty"`
}

// CreateVPNRequest represents a request to create a VPN
type CreateVPNRequest struct {
	RouterID string `json:"router_id"`
}

// CreateLoadBalancerClusterRequest represents a request to create a load balancer cluster
type CreateLoadBalancerClusterRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	AvailabilityZoneID string `json:"availability_zone_id"`
	Nodes              int    `json:"nodes"`
}

// PatchLoadBalancerClusterRequest represents a request to patch a load balancer cluster
type PatchLoadBalancerClusterRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateInstanceRequest represents a request to create an instance
type CreateInstanceRequest struct {
	Name               string                 `json:"name,omitempty"`
	VPCID              string                 `json:"vpc_id"`
	ImageID            string                 `json:"image_id"`
	ImageData          map[string]interface{} `json:"image_data"`
	VCPUCores          int                    `json:"vcpu_cores"`
	RAMCapacity        int                    `json:"ram_capacity"`
	Locked             bool                   `json:"locked"`
	VolumeCapacity     int                    `json:"volume_capacity"`
	VolumeIOPS         int                    `json:"volume_iops,omitempty"`
	BackupEnabled      bool                   `json:"backup_enabled"`
	NetworkID          string                 `json:"network_id,omitempty"`
	FloatingIPID       string                 `json:"floating_ip_id,omitempty"`
	RequiresFloatingIP bool                   `json:"requires_floating_ip"`
	UserScript         string                 `json:"user_script,omitempty"`
	SSHKeyPairIDs      []string               `json:"ssh_key_pair_ids,omitempty"`
	HostGroupID        string                 `json:"host_group_id,omitempty"`
}

// PatchInstanceRequest represents a request to patch an instance
type PatchInstanceRequest struct {
	Name          string  `json:"name,omitempty"`
	VCPUCores     int     `json:"vcpu_cores,omitempty"`
	RAMCapacity   int     `json:"ram_capacity,omitempty"`
	VolumeGroupID *string `json:"volume_group_id,omitempty"`
}

// CreateFirewallPolicyRequest represents a request to create a firewall policy
type CreateFirewallPolicyRequest struct {
	RouterID string `json:"router_id"`
	Name     string `json:"name,omitempty"`
	Sequence int    `json:"sequence"`
}

// PatchFirewallPolicyRequest represents a request to patch a firewall policy
type PatchFirewallPolicyRequest struct {
	Name     string `json:"name,omitempty"`
	Sequence *int   `json:"sequence,omitempty"`
}

// CreateVolumeRequest represents a request to create a volume
type CreateVolumeRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	Capacity           int    `json:"capacity"`
	IOPS               int    `json:"iops,omitempty"`
	AvailabilityZoneID string `json:"availability_zone_id"`
	VolumeGroupID      string `json:"volume_group_id,omitempty"`
}

// PatchVolumeRequest represents a request to patch a volume
type PatchVolumeRequest struct {
	Name          string  `json:"name,omitempty"`
	Capacity      int     `json:"capacity,omitempty"`
	IOPS          int     `json:"iops,omitempty"`
	VolumeGroupID *string `json:"volume_group_id,omitempty"`
}

// CreateFirewallRuleRequest represents a request to create a firewall rule
type CreateFirewallRuleRequest struct {
	Name             string                          `json:"name,omitempty"`
	FirewallPolicyID string                          `json:"firewall_policy_id"`
	Sequence         int                             `json:"sequence"`
	Source           string                          `json:"source"`
	Destination      string                          `json:"destination"`
	Ports            []CreateFirewallRulePortRequest `json:"ports,omitempty"`
	Action           FirewallRuleAction              `json:"action"`
	Direction        FirewallRuleDirection           `json:"direction"`
	Enabled          bool                            `json:"enabled"`
}

// PatchFirewallRuleRequest represents a request to patch a firewall rule
type PatchFirewallRuleRequest struct {
	Name        string                         `json:"name,omitempty"`
	Sequence    *int                           `json:"sequence,omitempty"`
	Source      string                         `json:"source,omitempty"`
	Destination string                         `json:"destination,omitempty"`
	Ports       []PatchFirewallRulePortRequest `json:"ports,omitempty"`
	Action      FirewallRuleAction             `json:"action,omitempty"`
	Direction   FirewallRuleDirection          `json:"direction,omitempty"`
	Enabled     *bool                          `json:"enabled,omitempty"`
}

// CreateFirewallRulePortRequest represents a request to create a firewall rule port
type CreateFirewallRulePortRequest struct {
	Name           string                   `json:"name,omitempty"`
	FirewallRuleID string                   `json:"firewall_rule_id"`
	Protocol       FirewallRulePortProtocol `json:"protocol"`
	Source         string                   `json:"source,omitempty"`
	Destination    string                   `json:"destination,omitempty"`
}

// PatchFirewallRulePortRequest represents a request to patch a firewall rule port
type PatchFirewallRulePortRequest struct {
	Name        string                   `json:"name,omitempty"`
	Protocol    FirewallRulePortProtocol `json:"protocol,omitempty"`
	Source      string                   `json:"source,omitempty"`
	Destination string                   `json:"destination,omitempty"`
}

// CreateFloatingIPRequest represents a request to create a floating IP
type CreateFloatingIPRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	AvailabilityZoneID string `json:"availability_zone_id"`
}

// PatchFloatingIPRequest represents a request to patch a floating IP
type PatchFloatingIPRequest struct {
	Name string `json:"name,omitempty"`
}

// AssignFloatingIPRequest represents a request to assign a floating IP to a resource
type AssignFloatingIPRequest struct {
	ResourceID string `json:"resource_id"`
}

// CreateHostGroupRequest represents a request to create a host group
type CreateHostGroupRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	AvailabilityZoneID string `json:"availability_zone_id"`
	HostSpecID         string `json:"host_spec_id"`
	WindowsEnabled     bool   `json:"windows_enabled"`
}

// PatchHostGroupRequest represents a request to patch a host group
type PatchHostGroupRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateSSHKeyPairRequest represents a request to create a SSH key pair
type CreateSSHKeyPairRequest struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key"`
}

// PatchSSHKeyPairRequest represents a request to patch a SSH key pair
type PatchSSHKeyPairRequest struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// AttachDetachInstanceVolumeRequest represents a request to attach / detach a volume to/from an instance
type AttachDetachInstanceVolumeRequest struct {
	VolumeID string `json:"volume_id"`
}

// CreateHostRequest represents a request to create a host
type CreateHostRequest struct {
	Name        string `json:"name,omitempty"`
	HostGroupID string `json:"host_group_id"`
}

// PatchHostRequest represents a request to patch a host
type PatchHostRequest struct {
	Name string `json:"name,omitempty"`
}

type NetworkPolicyCatchallRuleAction string

const (
	NetworkPolicyCatchallRuleActionAllow  NetworkPolicyCatchallRuleAction = "ALLOW"
	NetworkPolicyCatchallRuleActionDrop   NetworkPolicyCatchallRuleAction = "DROP"
	NetworkPolicyCatchallRuleActionReject NetworkPolicyCatchallRuleAction = "REJECT"
)

var NetworkPolicyCatchallRuleActionEnum connection.EnumSlice = []connection.Enum{
	NetworkPolicyCatchallRuleActionAllow,
	NetworkPolicyCatchallRuleActionDrop,
	NetworkPolicyCatchallRuleActionReject,
}

// ParseNetworkPolicyCatchallRuleAction attempts to parse a NetworkPolicyCatchallRuleAction from string
func ParseNetworkPolicyCatchallRuleAction(s string) (NetworkPolicyCatchallRuleAction, error) {
	e, err := connection.ParseEnum(s, NetworkPolicyCatchallRuleActionEnum)
	if err != nil {
		return "", err
	}

	return e.(NetworkPolicyCatchallRuleAction), err
}

func (s NetworkPolicyCatchallRuleAction) String() string {
	return string(s)
}

// CreateNetworkPolicyRequest represents a request to create a network policy
type CreateNetworkPolicyRequest struct {
	Name               string                          `json:"name,omitempty"`
	NetworkID          string                          `json:"network_id"`
	CatchallRuleAction NetworkPolicyCatchallRuleAction `json:"catchall_rule_action,omitempty"`
}

// PatchNetworkPolicyRequest represents a request to patch a network policy
type PatchNetworkPolicyRequest struct {
	Name               string                          `json:"name,omitempty"`
	CatchallRuleAction NetworkPolicyCatchallRuleAction `json:"catchall_rule_action,omitempty"`
}

// CreateNetworkRuleRequest represents a request to create a network rule
type CreateNetworkRuleRequest struct {
	Name            string                         `json:"name,omitempty"`
	NetworkPolicyID string                         `json:"network_policy_id"`
	Sequence        int                            `json:"sequence"`
	Source          string                         `json:"source"`
	Destination     string                         `json:"destination"`
	Ports           []CreateNetworkRulePortRequest `json:"ports,omitempty"`
	Action          NetworkRuleAction              `json:"action"`
	Direction       NetworkRuleDirection           `json:"direction"`
	Enabled         bool                           `json:"enabled"`
}

// PatchNetworkRuleRequest represents a request to patch a network rule
type PatchNetworkRuleRequest struct {
	Name        string                        `json:"name,omitempty"`
	Sequence    *int                          `json:"sequence,omitempty"`
	Source      string                        `json:"source,omitempty"`
	Destination string                        `json:"destination,omitempty"`
	Ports       []PatchNetworkRulePortRequest `json:"ports,omitempty"`
	Action      NetworkRuleAction             `json:"action,omitempty"`
	Direction   NetworkRuleDirection          `json:"direction,omitempty"`
	Enabled     *bool                         `json:"enabled,omitempty"`
}

// CreateNetworkRulePortRequest represents a request to create a network rule port
type CreateNetworkRulePortRequest struct {
	Name          string                  `json:"name,omitempty"`
	NetworkRuleID string                  `json:"network_rule_id"`
	Protocol      NetworkRulePortProtocol `json:"protocol"`
	Source        string                  `json:"source,omitempty"`
	Destination   string                  `json:"destination,omitempty"`
}

// PatchNetworkRulePortRequest represents a request to patch a network rule port
type PatchNetworkRulePortRequest struct {
	Name        string                  `json:"name,omitempty"`
	Protocol    NetworkRulePortProtocol `json:"protocol,omitempty"`
	Source      string                  `json:"source,omitempty"`
	Destination string                  `json:"destination,omitempty"`
}

// CreateVPNEndpointRequest represents a request to create a VPN endpoint
type CreateVPNEndpointRequest struct {
	Name         string `json:"name,omitempty"`
	VPNServiceID string `json:"vpn_service_id"`
	FloatingIPID string `json:"floating_ip_id,omitempty"`
}

// PatchVPNEndpointRequest represents a request to patch a VPN endpoint
type PatchVPNEndpointRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateVPNSessionRequest represents a request to create a VPN session
type CreateVPNSessionRequest struct {
	Name              string               `json:"name,omitempty"`
	VPNProfileGroupID string               `json:"vpn_profile_group_id"`
	VPNServiceID      string               `json:"vpn_service_id"`
	VPNEndpointID     string               `json:"vpn_endpoint_id"`
	RemoteIP          connection.IPAddress `json:"remote_ip"`
	RemoteNetworks    string               `json:"remote_networks,omitempty"`
	LocalNetworks     string               `json:"local_networks,omitempty"`
}

// PatchVPNSessionRequest represents a request to patch a VPN session
type PatchVPNSessionRequest struct {
	Name              string               `json:"name,omitempty"`
	VPNProfileGroupID string               `json:"vpn_profile_group_id,omitempty"`
	RemoteIP          connection.IPAddress `json:"remote_ip,omitempty"`
	RemoteNetworks    string               `json:"remote_networks,omitempty"`
	LocalNetworks     string               `json:"local_networks,omitempty"`
}

// CreateVPNServiceRequest represents a request to create a VPN service
type CreateVPNServiceRequest struct {
	Name     string `json:"name,omitempty"`
	RouterID string `json:"router_id"`
}

// PatchVPNServiceRequest represents a request to patch a VPN service
type PatchVPNServiceRequest struct {
	Name string `json:"name,omitempty"`
}

// MigrateInstanceRequest represents a request to migrate an instance
type MigrateInstanceRequest struct {
	HostGroupID string `json:"host_group_id,omitempty"`
}

// CreateVolumeGroupRequest represents a request to create a volume group
type CreateVolumeGroupRequest struct {
	Name               string `json:"name,omitempty"`
	VPCID              string `json:"vpc_id"`
	AvailabilityZoneID string `json:"availability_zone_id"`
}

// PatchVolumeGroupRequest represents a request to patch a volume group
type PatchVolumeGroupRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateLoadBalancerRequest represents a request to create a load balancer
type CreateLoadBalancerRequest struct {
	Name               string `json:"name,omitempty"`
	AvailabilityZoneID string `json:"availability_zone_id"`
	VPCID              string `json:"vpc_id"`
	LoadBalancerSpecID string `json:"load_balancer_spec_id"`
	NetworkID          string `json:"network_id"`
}

// CreateLoadBalancerRequest represents a request to patch a load balancer
type PatchLoadBalancerRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateVIPRequest represents a request to create a load balancer VIP
type CreateVIPRequest struct {
	Name               string `json:"name,omitempty"`
	LoadBalancerID     string `json:"load_balancer_id"`
	AllocateFloatingIP bool   `json:"allocate_floating_ip"`
}

// PatchVIPRequest represents a request to patch a load balancer VIP
type PatchVIPRequest struct {
	Name string `json:"name,omitempty"`
}
