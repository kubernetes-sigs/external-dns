//go:generate go run ../../gen/model_response/main.go -package ecloud -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package ecloud -source model.go -destination model_paginated_generated.go

package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

type VirtualMachineStatus string

const (
	VirtualMachineStatusComplete   VirtualMachineStatus = "Complete"
	VirtualMachineStatusFailed     VirtualMachineStatus = "Failed"
	VirtualMachineStatusBeingBuilt VirtualMachineStatus = "Being Built"
)

func (s VirtualMachineStatus) String() string {
	return string(s)
}

type VirtualMachineDiskType string

func (e VirtualMachineDiskType) String() string {
	return string(e)
}

const (
	VirtualMachineDiskTypeStandard VirtualMachineDiskType = "Standard"
	VirtualMachineDiskTypeCluster  VirtualMachineDiskType = "Cluster"
)

type VirtualMachinePowerStatus string

func (s VirtualMachinePowerStatus) String() string {
	return string(s)
}

const (
	VirtualMachinePowerStatusOnline  VirtualMachinePowerStatus = "Online"
	VirtualMachinePowerStatusOffline VirtualMachinePowerStatus = "Offline"
)

var VirtualMachinePowerStatusEnum connection.EnumSlice = []connection.Enum{
	VirtualMachinePowerStatusOnline,
	VirtualMachinePowerStatusOffline,
}

// ParseVirtualMachinePowerStatus attempts to parse a VirtualMachinePowerStatus from string
func ParseVirtualMachinePowerStatus(s string) (VirtualMachinePowerStatus, error) {
	e, err := connection.ParseEnum(s, VirtualMachinePowerStatusEnum)
	if err != nil {
		return "", err
	}

	return e.(VirtualMachinePowerStatus), err
}

type DatastoreStatus string

func (s DatastoreStatus) String() string {
	return string(s)
}

const (
	DatastoreStatusCompleted DatastoreStatus = "Completed"
	DatastoreStatusFailed    DatastoreStatus = "Failed"
	DatastoreStatusExpanding DatastoreStatus = "Expanding"
	DatastoreStatusQueued    DatastoreStatus = "Queued"
)

type SolutionEnvironment string

const (
	SolutionEnvironmentHybrid  SolutionEnvironment = "Hybrid"
	SolutionEnvironmentPrivate SolutionEnvironment = "Private"
)

func (s SolutionEnvironment) String() string {
	return string(s)
}

type FirewallRole string

func (r FirewallRole) String() string {
	return string(r)
}

const (
	FirewallRoleNA     FirewallRole = "N/A"
	FirewallRoleMaster FirewallRole = "Master"
	FirewallRoleSlave  FirewallRole = "Slave"
)

// VirtualMachine represents an eCloud Virtual Machine
// +genie:model_response
// +genie:model_paginated
type VirtualMachine struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Hostname     string `json:"hostname"`
	ComputerName string `json:"computername"`
	// Count in Cores
	CPU int `json:"cpu"`
	// Size in GB
	RAM int `json:"ram"`
	// Size in GB
	HDD         int                  `json:"hdd"`
	IPInternal  connection.IPAddress `json:"ip_internal"`
	IPExternal  connection.IPAddress `json:"ip_external"`
	Platform    string               `json:"platform"`
	Template    string               `json:"template"`
	Backup      bool                 `json:"backup"`
	Support     bool                 `json:"support"`
	Environment string               `json:"environment"`
	SolutionID  int                  `json:"solution_id"`
	Status      VirtualMachineStatus `json:"status"`
	PowerStatus string               `json:"power_status"`
	ToolsStatus string               `json:"tools_status"`
	Disks       []VirtualMachineDisk `json:"hdd_disks"`
	Encrypted   bool                 `json:"encrypted"`
	Role        string               `json:"role"`
	GPUProfile  string               `json:"gpu_profile"`
}

// VirtualMachineDisk represents an eCloud Virtual Machine disk
type VirtualMachineDisk struct {
	UUID string                 `json:"uuid"`
	Name string                 `json:"name"`
	Type VirtualMachineDiskType `json:"type"`
	Key  int                    `json:"key"`

	// Size in GB
	Capacity int `json:"capacity"`
}

// Tag represents an eCloud tag
// +genie:model_response
// +genie:model_paginated
type Tag struct {
	Key       string              `json:"key"`
	Value     string              `json:"value"`
	CreatedAt connection.DateTime `json:"created_at"`
}

// Solution represents an eCloud solution
// +genie:model_response
// +genie:model_paginated
type Solution struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Environment       SolutionEnvironment `json:"environment"`
	PodID             int                 `json:"pod_id"`
	EncryptionEnabled bool                `json:"encryption_enabled"`
	EncryptionDefault bool                `json:"encryption_default"`
}

// Site represents an eCloud site
// +genie:model_response
// +genie:model_paginated
type Site struct {
	ID         int    `json:"id"`
	State      string `json:"state"`
	SolutionID int    `json:"solution_id"`
	PodID      int    `json:"pod_id"`
}

// V1Network represents an eCloud v1 network
// +genie:model_response
// +genie:model_paginated
type V1Network struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// V1Host represents an eCloud v1 host
// +genie:model_response
// +genie:model_paginated
type V1Host struct {
	ID         int     `json:"id"`
	SolutionID int     `json:"solution_id"`
	PodID      int     `json:"pod_id"`
	Name       string  `json:"name"`
	CPU        HostCPU `json:"cpu"`
	RAM        HostRAM `json:"ram"`
}

// HostCPU represents an eCloud host's CPU resources
type HostCPU struct {
	Quantity int    `json:"qty"`
	Cores    int    `json:"cores"`
	Speed    string `json:"speed"`
}

// HostRAM represents an eCloud host's RAM resources
type HostRAM struct {
	// Size in GB
	Capacity int `json:"capacity"`
	// Size in GB
	Reserved int `json:"reserved"`
	// Size in GB
	Allocated int `json:"allocated"`
	// Size in GB
	Available int `json:"available"`
}

// Datastore represents an eCloud datastore
// +genie:model_response
// +genie:model_paginated
type Datastore struct {
	ID         int             `json:"id"`
	SolutionID int             `json:"solution_id"`
	SiteID     int             `json:"site_id"`
	Name       string          `json:"name"`
	Status     DatastoreStatus `json:"status"`
	// Size in GB
	Capacity int `json:"capacity"`
	// Size in GB
	Allocated int `json:"allocated"`
	// Size in GB
	Available int `json:"available"`
}

// Firewall represents an eCloud firewall
// +genie:model_response
// +genie:model_paginated
type Firewall struct {
	ID       int                  `json:"id"`
	Name     string               `json:"name"`
	Hostname string               `json:"hostname"`
	IP       connection.IPAddress `json:"ip"`
	Role     FirewallRole         `json:"role"`
}

// FirewallConfig represents an eCloud firewall config
// +genie:model_response
type FirewallConfig struct {
	Config string `json:"config"`
}

// Template represents an eCloud template
// +genie:model_response
// +genie:model_paginated
type Template struct {
	Name string `json:"name"`
	// Count in Cores
	CPU int `json:"cpu"`
	// Size in GB
	RAM int `json:"ram"`
	// Size in GB
	HDD             int                  `json:"hdd"`
	Disks           []VirtualMachineDisk `json:"hdd_disks"`
	Platform        string               `json:"platform"`
	OperatingSystem string               `json:"operating_system"`
	SolutionID      int                  `json:"solution_id"`
}

// Pod represents an eCloud pod
// +genie:model_response
// +genie:model_paginated
type Pod struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Services struct {
		Public     bool `json:"public"`
		Burst      bool `json:"burst"`
		Appliances bool `json:"appliances"`
	} `json:"services"`
}

// Appliance represents an eCloud appliance
// +genie:model_response
// +genie:model_paginated
type Appliance struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	LogoURI          string              `json:"logo_uri"`
	Description      string              `json:"description"`
	DocumentationURI string              `json:"documentation_uri"`
	Publisher        string              `json:"publisher"`
	CreatedAt        connection.DateTime `json:"created_at"`
}

// ApplianceParameter represents an eCloud appliance parameter
// +genie:model_response
// +genie:model_paginated
type ApplianceParameter struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Key            string `json:"key"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Required       bool   `json:"required"`
	ValidationRule string `json:"validation_rule"`
}

// ActiveDirectoryDomain represents an eCloud active directory domain
// +genie:model_response
// +genie:model_paginated
type ActiveDirectoryDomain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TemplateType string

const (
	TemplateTypeSolution TemplateType = "solution"
	TemplateTypePod      TemplateType = "pod"
)

var TemplateTypeEnum connection.EnumSlice = []connection.Enum{
	TemplateTypeSolution,
	TemplateTypePod,
}

// ParseTemplateType attempts to parse a TemplateType from string
func ParseTemplateType(s string) (TemplateType, error) {
	e, err := connection.ParseEnum(s, TemplateTypeEnum)
	if err != nil {
		return "", err
	}

	return e.(TemplateType), err
}

func (s TemplateType) String() string {
	return string(s)
}

// ConsoleSession represents an eCloud Virtual Machine console session
// +genie:model_response
type ConsoleSession struct {
	URL string `json:"url"`
}

type SyncStatus string

const (
	SyncStatusComplete   SyncStatus = "complete"
	SyncStatusFailed     SyncStatus = "failed"
	SyncStatusInProgress SyncStatus = "in-progress"
)

func (s SyncStatus) String() string {
	return string(s)
}

type SyncType string

const (
	SyncTypeUpdate SyncType = "update"
	SyncTypeDelete SyncType = "delete"
)

func (s SyncType) String() string {
	return string(s)
}

// ResourceSync represents the sync status of a resource
type ResourceSync struct {
	Status SyncStatus `json:"status"`
	Type   SyncType   `json:"type"`
}

type TaskStatus string

const (
	TaskStatusComplete   TaskStatus = "complete"
	TaskStatusFailed     TaskStatus = "failed"
	TaskStatusInProgress TaskStatus = "in-progress"
)

var TaskStatusEnum connection.EnumSlice = []connection.Enum{
	TaskStatusComplete,
	TaskStatusFailed,
	TaskStatusInProgress,
}

// ParseTaskStatus attempts to parse a TaskStatus from string
func ParseTaskStatus(s string) (TaskStatus, error) {
	e, err := connection.ParseEnum(s, TaskStatusEnum)
	if err != nil {
		return "", err
	}

	return e.(TaskStatus), err
}

func (s TaskStatus) String() string {
	return string(s)
}

// VPC represents an eCloud VPC
// +genie:model_response
// +genie:model_paginated
type VPC struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	RegionID           string              `json:"region_id"`
	Sync               ResourceSync        `json:"sync"`
	SupportEnabled     bool                `json:"support_enabled"`
	ConsoleEnabled     bool                `json:"console_enabled"`
	AdvancedNetworking bool                `json:"advanced_networking"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// AvailabilityZone represents an eCloud availability zone
// +genie:model_response
// +genie:model_paginated
type AvailabilityZone struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	DatacentreSiteID int    `json:"datacentre_site_id"`
	RegionID         string `json:"region_id"`
}

// Network represents an eCloud network
// +genie:model_response
// +genie:model_paginated
type Network struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	Subnet    string              `json:"subnet"`
	Sync      ResourceSync        `json:"sync"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// DHCP represents an eCloud DHCP server/policy
// +genie:model_response
// +genie:model_paginated
type DHCP struct {
	ID                 string              `json:"id"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// VPN represents an eCloud VPN
// +genie:model_response
// +genie:model_paginated
type VPN struct {
	ID        string              `json:"id"`
	RouterID  string              `json:"router_id"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Instance represents an eCloud instance
// +genie:model_response
// +genie:model_paginated
type Instance struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	ImageID            string              `json:"image_id"`
	VCPUCores          int                 `json:"vcpu_cores"`
	RAMCapacity        int                 `json:"ram_capacity"`
	Locked             bool                `json:"locked"`
	BackupEnabled      bool                `json:"backup_enabled"`
	Platform           string              `json:"platform"`
	VolumeCapacity     int                 `json:"volume_capacity"`
	VolumeGroupID      string              `json:"volume_group_id"`
	HostGroupID        string              `json:"host_group_id"`
	Sync               ResourceSync        `json:"sync"`
	Online             *bool               `json:"online"`
	AgentRunning       *bool               `json:"agent_running"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// FloatingIP represents an eCloud floating IP address
// +genie:model_response
// +genie:model_paginated
type FloatingIP struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	IPAddress          string              `json:"ip_address"`
	ResourceID         string              `json:"resource_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Sync               ResourceSync        `json:"sync"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// FirewallPolicy represents an eCloud firewall policy
// +genie:model_response
// +genie:model_paginated
type FirewallPolicy struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	Sequence  int                 `json:"sequence"`
	Sync      ResourceSync        `json:"sync"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

type FirewallRuleAction string

const (
	FirewallRuleActionAllow  FirewallRuleAction = "ALLOW"
	FirewallRuleActionDrop   FirewallRuleAction = "DROP"
	FirewallRuleActionReject FirewallRuleAction = "REJECT"
)

var FirewallRuleActionEnum connection.EnumSlice = []connection.Enum{
	FirewallRuleActionAllow,
	FirewallRuleActionDrop,
	FirewallRuleActionReject,
}

// ParseFirewallRuleAction attempts to parse a FirewallRuleAction from string
func ParseFirewallRuleAction(s string) (FirewallRuleAction, error) {
	e, err := connection.ParseEnum(s, FirewallRuleActionEnum)
	if err != nil {
		return "", err
	}

	return e.(FirewallRuleAction), err
}

func (s FirewallRuleAction) String() string {
	return string(s)
}

type FirewallRuleDirection string

const (
	FirewallRuleDirectionIn    FirewallRuleDirection = "IN"
	FirewallRuleDirectionOut   FirewallRuleDirection = "OUT"
	FirewallRuleDirectionInOut FirewallRuleDirection = "IN_OUT"
)

var FirewallRuleDirectionEnum connection.EnumSlice = []connection.Enum{
	FirewallRuleDirectionIn,
	FirewallRuleDirectionOut,
	FirewallRuleDirectionInOut,
}

// ParseFirewallRuleDirection attempts to parse a FirewallRuleDirection from string
func ParseFirewallRuleDirection(s string) (FirewallRuleDirection, error) {
	e, err := connection.ParseEnum(s, FirewallRuleDirectionEnum)
	if err != nil {
		return "", err
	}

	return e.(FirewallRuleDirection), err
}

func (s FirewallRuleDirection) String() string {
	return string(s)
}

// FirewallRule represents an eCloud firewall rule
// +genie:model_response
// +genie:model_paginated
type FirewallRule struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	FirewallPolicyID string                `json:"firewall_policy_id"`
	Sequence         int                   `json:"sequence"`
	Source           string                `json:"source"`
	Destination      string                `json:"destination"`
	Action           FirewallRuleAction    `json:"action"`
	Direction        FirewallRuleDirection `json:"direction"`
	Enabled          bool                  `json:"enabled"`
	CreatedAt        connection.DateTime   `json:"created_at"`
	UpdatedAt        connection.DateTime   `json:"updated_at"`
}

type FirewallRulePortProtocol string

const (
	FirewallRulePortProtocolTCP    FirewallRulePortProtocol = "TCP"
	FirewallRulePortProtocolUDP    FirewallRulePortProtocol = "UDP"
	FirewallRulePortProtocolICMPv4 FirewallRulePortProtocol = "ICMPv4"
)

var FirewallRulePortProtocolEnum connection.EnumSlice = []connection.Enum{
	FirewallRulePortProtocolTCP,
	FirewallRulePortProtocolUDP,
	FirewallRulePortProtocolICMPv4,
}

// ParseFirewallRulePortProtocol attempts to parse a FirewallRulePortProtocol from string
func ParseFirewallRulePortProtocol(s string) (FirewallRulePortProtocol, error) {
	e, err := connection.ParseEnum(s, FirewallRulePortProtocolEnum)
	if err != nil {
		return "", err
	}

	return e.(FirewallRulePortProtocol), err
}

func (s FirewallRulePortProtocol) String() string {
	return string(s)
}

// FirewallRulePort represents an eCloud firewall rule port
// +genie:model_response
// +genie:model_paginated
type FirewallRulePort struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	FirewallRuleID string                   `json:"firewall_rule_id"`
	Protocol       FirewallRulePortProtocol `json:"protocol"`
	Source         string                   `json:"source"`
	Destination    string                   `json:"destination"`
	CreatedAt      connection.DateTime      `json:"created_at"`
	UpdatedAt      connection.DateTime      `json:"updated_at"`
}

// Region represents an eCloud region
// +genie:model_response
// +genie:model_paginated
type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Router represents an eCloud router
// +genie:model_response
// +genie:model_paginated
type Router struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	RouterThroughputID string              `json:"router_throughput_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Sync               ResourceSync        `json:"sync"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// Credential represents an eCloud credential
// +genie:model_response
// +genie:model_paginated
type Credential struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	ResourceID string              `json:"resource_id"`
	Host       string              `json:"host"`
	Username   string              `json:"username"`
	Password   string              `json:"password"`
	Port       int                 `json:"port"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

type VolumeType string

const (
	VolumeTypeOS   VolumeType = "os"
	VolumeTypeData VolumeType = "data"
)

func (s VolumeType) String() string {
	return string(s)
}

// Volume represents an eCloud volume
// +genie:model_response
// +genie:model_paginated
type Volume struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Capacity           int                 `json:"capacity"`
	IOPS               int                 `json:"iops"`
	Attached           bool                `json:"attached"`
	Type               VolumeType          `json:"type"`
	VolumeGroupID      string              `json:"volume_group_id"`
	IsShared           bool                `json:"is_shared"`
	Sync               ResourceSync        `json:"sync"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// NIC represents an eCloud network interface card
// +genie:model_response
// +genie:model_paginated
type NIC struct {
	ID         string              `json:"id"`
	MACAddress string              `json:"mac_address"`
	InstanceID string              `json:"instance_id"`
	NetworkID  string              `json:"network_id"`
	IPAddress  string              `json:"ip_address"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// RouterThroughput represents an eCloud router throughput
// +genie:model_response
// +genie:model_paginated
type RouterThroughput struct {
	ID                 string              `json:"id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Name               string              `json:"name"`
	CommittedBandwidth int                 `json:"committed_bandwidth"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// DiscountPlan represents an eCloud discount plan
// +genie:model_response
// +genie:model_paginated
type DiscountPlan struct {
	ID                       string              `json:"id"`
	ResellerID               int                 `json:"reseller_id"`
	ContactID                int                 `json:"contact_id"`
	Name                     string              `json:"name"`
	CommitmentAmount         float32             `json:"commitment_amount"`
	CommitmentBeforeDiscount float32             `json:"commitment_before_discount"`
	DiscountRate             float32             `json:"discount_rate"`
	TermLength               int                 `json:"term_length"`
	TermStartDate            connection.DateTime `json:"term_start_date"`
	TermEndDate              connection.DateTime `json:"term_end_date"`
	Status                   string              `json:"status"`
	ResponseDate             connection.DateTime `json:"response_date"`
	CreatedAt                connection.DateTime `json:"created_at"`
	UpdatedAt                connection.DateTime `json:"updated_at"`
}

// BillingMetric represents an eCloud billing metric
// +genie:model_response
// +genie:model_paginated
type BillingMetric struct {
	ID         string              `json:"id"`
	ResourceID string              `json:"resource_id"`
	VPCID      string              `json:"vpc_id"`
	Key        string              `json:"key"`
	Value      string              `json:"value"`
	Start      connection.DateTime `json:"start"`
	End        connection.DateTime `json:"end"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Image represents an eCloud image
// +genie:model_response
// +genie:model_paginated
type Image struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	LogoURI          string              `json:"logo_uri"`
	Description      string              `json:"description"`
	DocumentationURI string              `json:"documentation_uri"`
	Publisher        string              `json:"publisher"`
	CreatedAt        connection.DateTime `json:"created_at"`
	UpdatedAt        connection.DateTime `json:"updated_at"`
}

// ImageParameter represents an eCloud image parameter
// +genie:model_response
// +genie:model_paginated
type ImageParameter struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Key            string              `json:"key"`
	Type           string              `json:"type"`
	Description    string              `json:"description"`
	Required       bool                `json:"required"`
	ValidationRule string              `json:"validation_rule"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// ImageMetadata represents eCloud image metadata
// +genie:model_response
// +genie:model_paginated
type ImageMetadata struct {
	Key       string              `json:"key"`
	Value     string              `json:"value"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// HostGroup represents an eCloud host group
// +genie:model_response
// +genie:model_paginated
type HostGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	WindowsEnabled     bool                `json:"windows_enabled"`
	HostSpecID         string              `json:"host_spec_id"`
	Sync               ResourceSync        `json:"sync"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// HostSpec represents an eCloud host spec
// +genie:model_response
// +genie:model_paginated
type HostSpec struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CPUSockets    int    `json:"cpu_sockets"`
	CPUType       string `json:"cpu_type"`
	CPUCores      int    `json:"cpu_cores"`
	CPUClockSpeed int    `json:"cpu_clock_speed"`
	RAMCapacity   int    `json:"ram_capacity"`
}

// Host represents an eCloud v2 host
// +genie:model_response
// +genie:model_paginated
type Host struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	HostGroupID string              `json:"host_group_id"`
	Sync        ResourceSync        `json:"sync"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// SSHKeyPair represents an eCloud SSH key pair
// +genie:model_response
// +genie:model_paginated
type SSHKeyPair struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

// Task represents a task against an eCloud resource
// +genie:model_response
// +genie:model_paginated
type Task struct {
	ID         string              `json:"id"`
	ResourceID string              `json:"resource_id"`
	Name       string              `json:"name"`
	Status     TaskStatus          `json:"status"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// TaskReference represents a reference to an on-going task
// +genie:model_response
type TaskReference struct {
	TaskID     string `json:"task_id"`
	ResourceID string `json:"id"`
}

// NetworkPolicy represents an eCloud network policy
// +genie:model_response
// +genie:model_paginated
type NetworkPolicy struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	NetworkID string              `json:"network_id"`
	VPCID     string              `json:"vpc_id"`
	Sync      ResourceSync        `json:"sync"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

type NetworkRuleAction string

const (
	NetworkRuleActionAllow  NetworkRuleAction = "ALLOW"
	NetworkRuleActionDrop   NetworkRuleAction = "DROP"
	NetworkRuleActionReject NetworkRuleAction = "REJECT"
)

var NetworkRuleActionEnum connection.EnumSlice = []connection.Enum{
	NetworkRuleActionAllow,
	NetworkRuleActionDrop,
	NetworkRuleActionReject,
}

// ParseNetworkRuleAction attempts to parse a NetworkRuleAction from string
func ParseNetworkRuleAction(s string) (NetworkRuleAction, error) {
	e, err := connection.ParseEnum(s, NetworkRuleActionEnum)
	if err != nil {
		return "", err
	}

	return e.(NetworkRuleAction), err
}

func (s NetworkRuleAction) String() string {
	return string(s)
}

type NetworkRuleDirection string

const (
	NetworkRuleDirectionIn    NetworkRuleDirection = "IN"
	NetworkRuleDirectionOut   NetworkRuleDirection = "OUT"
	NetworkRuleDirectionInOut NetworkRuleDirection = "IN_OUT"
)

var NetworkRuleDirectionEnum connection.EnumSlice = []connection.Enum{
	NetworkRuleDirectionIn,
	NetworkRuleDirectionOut,
	NetworkRuleDirectionInOut,
}

// ParseNetworkRuleDirection attempts to parse a NetworkRuleDirection from string
func ParseNetworkRuleDirection(s string) (NetworkRuleDirection, error) {
	e, err := connection.ParseEnum(s, NetworkRuleDirectionEnum)
	if err != nil {
		return "", err
	}

	return e.(NetworkRuleDirection), err
}

func (s NetworkRuleDirection) String() string {
	return string(s)
}

// NetworkRule represents an eCloud network rule
// +genie:model_response
// +genie:model_paginated
type NetworkRule struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	NetworkPolicyID string               `json:"network_policy_id"`
	Sequence        int                  `json:"sequence"`
	Source          string               `json:"source"`
	Destination     string               `json:"destination"`
	Type            string               `json:"type"`
	Action          NetworkRuleAction    `json:"action"`
	Direction       NetworkRuleDirection `json:"direction"`
	Enabled         bool                 `json:"enabled"`
	CreatedAt       connection.DateTime  `json:"created_at"`
	UpdatedAt       connection.DateTime  `json:"updated_at"`
}

type NetworkRulePortProtocol string

const (
	NetworkRulePortProtocolTCP    NetworkRulePortProtocol = "TCP"
	NetworkRulePortProtocolUDP    NetworkRulePortProtocol = "UDP"
	NetworkRulePortProtocolICMPv4 NetworkRulePortProtocol = "ICMPv4"
)

var NetworkRulePortProtocolEnum connection.EnumSlice = []connection.Enum{
	NetworkRulePortProtocolTCP,
	NetworkRulePortProtocolUDP,
	NetworkRulePortProtocolICMPv4,
}

// ParseNetworkRulePortProtocol attempts to parse a NetworkRulePortProtocol from string
func ParseNetworkRulePortProtocol(s string) (NetworkRulePortProtocol, error) {
	e, err := connection.ParseEnum(s, NetworkRulePortProtocolEnum)
	if err != nil {
		return "", err
	}

	return e.(NetworkRulePortProtocol), err
}

func (s NetworkRulePortProtocol) String() string {
	return string(s)
}

// NetworkRulePort represents an eCloud network rule port
// +genie:model_response
// +genie:model_paginated
type NetworkRulePort struct {
	ID            string                  `json:"id"`
	Name          string                  `json:"name"`
	NetworkRuleID string                  `json:"network_rule_id"`
	Protocol      NetworkRulePortProtocol `json:"protocol"`
	Source        string                  `json:"source"`
	Destination   string                  `json:"destination"`
	CreatedAt     connection.DateTime     `json:"created_at"`
	UpdatedAt     connection.DateTime     `json:"updated_at"`
}

// VolumeGroup represents an eCloud volume group resource
// +genie:model_response
// +genie:model_paginated
type VolumeGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	VPCID              string              `json:"vpc_id"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	Usage              VolumeGroupUsage    `json:"usage"`
	Sync               ResourceSync        `json:"sync"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

type VolumeGroupUsage struct {
	Volumes int `json:"volumes"`
}

// VPNProfileGroup represents an eCloud VPN profile group
// +genie:model_response
// +genie:model_paginated
type VPNProfileGroup struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// VPNService represents an eCloud VPN service
// +genie:model_response
// +genie:model_paginated
type VPNService struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	RouterID  string              `json:"router_id"`
	VPCID     string              `json:"vpc_id"`
	Sync      ResourceSync        `json:"sync"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// VPNEndpoint represents an eCloud VPN endpoint
// +genie:model_response
// +genie:model_paginated
type VPNEndpoint struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	VPNServiceID string              `json:"vpn_service_id"`
	FloatingIPID string              `json:"floating_ip_id"`
	Sync         ResourceSync        `json:"sync"`
	CreatedAt    connection.DateTime `json:"created_at"`
	UpdatedAt    connection.DateTime `json:"updated_at"`
}

// VPNSession represents an eCloud VPN session
// +genie:model_response
// +genie:model_paginated
type VPNSession struct {
	ID                string                   `json:"id"`
	Name              string                   `json:"name"`
	VPNProfileGroupID string                   `json:"vpn_profile_group_id"`
	VPNServiceID      string                   `json:"vpn_service_id"`
	VPNEndpointID     string                   `json:"vpn_endpoint_id"`
	RemoteIP          connection.IPAddress     `json:"remote_ip"`
	RemoteNetworks    string                   `json:"remote_networks"`
	LocalNetworks     string                   `json:"local_networks"`
	TunnelDetails     *VPNSessionTunnelDetails `json:"tunnel_details"`
	Sync              ResourceSync             `json:"sync"`
	CreatedAt         connection.DateTime      `json:"created_at"`
	UpdatedAt         connection.DateTime      `json:"updated_at"`
}

type VPNSessionTunnelDetails struct {
	SessionState     string                       `json:"session_state"`
	TunnelStatistics []VPNSessionTunnelStatistics `json:"tunnel_statistics"`
}

type VPNSessionTunnelStatistics struct {
	TunnelStatus     string `json:"tunnel_status"`
	TunnelDownReason string `json:"tunnel_down_reason"`
	LocalSubnet      string `json:"local_subnet"`
	PeerSubnet       string `json:"peer_subnet"`
}

// VPNSessionPreSharedKey represents an eCloud VPN session pre-shared key
// +genie:model_response
type VPNSessionPreSharedKey struct {
	PSK string `json:"psk"`
}

// LoadBalancer represents an eCloud loadbalancer
// +genie:model_response
// +genie:model_paginated
type LoadBalancer struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	AvailabilityZoneID string              `json:"availability_zone_id"`
	VPCID              string              `json:"vpc_id"`
	LoadBalancerSpecID string              `json:"load_balancer_spec_id"`
	Sync               ResourceSync        `json:"sync"`
	ConfigID           int                 `json:"config_id"`
	Nodes              int                 `json:"nodes"`
	NetworkID          string              `json:"network_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// LoadBalancerSpec represents an eCloud loadbalancer specification
// +genie:model_response
// +genie:model_paginated
type LoadBalancerSpec struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// VIP represents an eCloud load balancer VIP
// +genie:model_response
// +genie:model_paginated
type VIP struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	LoadBalancerID string              `json:"load_balancer_id"`
	IPAddressID    string              `json:"ip_address_id"`
	ConfigID       int                 `json:"config_id"`
	Sync           ResourceSync        `json:"sync"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}
