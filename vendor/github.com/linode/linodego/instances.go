package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type InterfaceGeneration string

const (
	GenerationLegacyConfig InterfaceGeneration = "legacy_config"
	GenerationLinode       InterfaceGeneration = "linode"
)

/*
 * https://techdocs.akamai.com/linode-api/reference/post-linode-instance
 */

// InstanceStatus constants start with Instance and include Linode API Instance Status values
type InstanceStatus string

// InstanceStatus constants reflect the current status of an Instance
const (
	InstanceBooting      InstanceStatus = "booting"
	InstanceRunning      InstanceStatus = "running"
	InstanceOffline      InstanceStatus = "offline"
	InstanceShuttingDown InstanceStatus = "shutting_down"
	InstanceRebooting    InstanceStatus = "rebooting"
	InstanceProvisioning InstanceStatus = "provisioning"
	InstanceDeleting     InstanceStatus = "deleting"
	InstanceMigrating    InstanceStatus = "migrating"
	InstanceRebuilding   InstanceStatus = "rebuilding"
	InstanceCloning      InstanceStatus = "cloning"
	InstanceRestoring    InstanceStatus = "restoring"
	InstanceResizing     InstanceStatus = "resizing"
)

type InstanceMigrationType string

const (
	WarmMigration InstanceMigrationType = "warm"
	ColdMigration InstanceMigrationType = "cold"
)

// Instance represents a linode object
type Instance struct {
	ID                  int                     `json:"id"`
	Created             *time.Time              `json:"-"`
	Updated             *time.Time              `json:"-"`
	Region              string                  `json:"region"`
	Alerts              *InstanceAlert          `json:"alerts"`
	Backups             *InstanceBackup         `json:"backups"`
	Image               string                  `json:"image"`
	Group               string                  `json:"group"`
	IPv4                []*net.IP               `json:"ipv4"`
	IPv6                string                  `json:"ipv6"`
	Label               string                  `json:"label"`
	Type                string                  `json:"type"`
	Status              InstanceStatus          `json:"status"`
	HasUserData         bool                    `json:"has_user_data"`
	Hypervisor          string                  `json:"hypervisor"`
	HostUUID            string                  `json:"host_uuid"`
	Specs               *InstanceSpec           `json:"specs"`
	WatchdogEnabled     bool                    `json:"watchdog_enabled"`
	Tags                []string                `json:"tags"`
	PlacementGroup      *InstancePlacementGroup `json:"placement_group"`
	DiskEncryption      InstanceDiskEncryption  `json:"disk_encryption"`
	LKEClusterID        int                     `json:"lke_cluster_id"`
	Capabilities        []string                `json:"capabilities"`
	InterfaceGeneration InterfaceGeneration     `json:"interface_generation"`
	MaintenancePolicy   string                  `json:"maintenance_policy"`

	// NOTE: Locks can only be used with v4beta.
	Locks []LockType `json:"locks"`
}

// InstanceSpec represents a linode spec
type InstanceSpec struct {
	Disk               int `json:"disk"`
	Memory             int `json:"memory"`
	VCPUs              int `json:"vcpus"`
	Transfer           int `json:"transfer"`
	GPUs               int `json:"gpus"`
	AcceleratedDevices int `json:"accelerated_devices"`
}

// InstanceAlert represents a metric alert
type InstanceAlert struct {
	CPU           int `json:"cpu"`
	IO            int `json:"io"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
}

// InstanceBackup represents backup settings for an instance
type InstanceBackup struct {
	Available      bool       `json:"available,omitempty"` // read-only
	Enabled        bool       `json:"enabled,omitempty"`   // read-only
	LastSuccessful *time.Time `json:"-"`                   // read-only
	Schedule       struct {
		Day    string `json:"day,omitempty"`
		Window string `json:"window,omitempty"`
	} `json:"schedule"`
}

type InstanceDiskEncryption string

const (
	InstanceDiskEncryptionEnabled  InstanceDiskEncryption = "enabled"
	InstanceDiskEncryptionDisabled InstanceDiskEncryption = "disabled"
)

// InstanceTransfer pool stats for a Linode Instance during the current billing month
type InstanceTransfer struct {
	// Bytes of transfer this instance has consumed
	Used int `json:"used"`

	// GB of billable transfer this instance has consumed
	Billable int `json:"billable"`

	// GB of transfer this instance adds to the Transfer pool
	Quota int `json:"quota"`
}

// Deprecated: use MonthlyInstanceTransferStatsV2 for new implementations
//
// MonthlyInstanceTransferStats pool stats for a Linode Instance network transfer statistics for a specific month
type MonthlyInstanceTransferStats struct {
	// The amount of inbound public network traffic received by this Linode, in bytes, for a specific year/month.
	BytesIn int `json:"bytes_in"`

	// The amount of outbound public network traffic sent by this Linode, in bytes, for a specific year/month.
	BytesOut int `json:"bytes_out"`

	// The total amount of public network traffic sent and received by this Linode, in bytes, for a specific year/month.
	BytesTotal int `json:"bytes_total"`
}

// MonthlyInstanceTransferStatsV2 pool stats for a Linode Instance network transfer statistics for a specific month
type MonthlyInstanceTransferStatsV2 struct {
	// The amount of inbound public network traffic received by this Linode, in bytes, for a specific year/month.
	BytesIn uint64 `json:"bytes_in"`

	// The amount of outbound public network traffic sent by this Linode, in bytes, for a specific year/month.
	BytesOut uint64 `json:"bytes_out"`

	// The total amount of public network traffic sent and received by this Linode, in bytes, for a specific year/month.
	BytesTotal uint64 `json:"bytes_total"`
}

// InstancePlacementGroup represents information about the placement group
// this Linode is a part of.
type InstancePlacementGroup struct {
	ID                   int                  `json:"id"`
	Label                string               `json:"label"`
	PlacementGroupType   PlacementGroupType   `json:"placement_group_type"`
	PlacementGroupPolicy PlacementGroupPolicy `json:"placement_group_policy"`
	MigratingTo          *int                 `json:"migrating_to"` // read-only
}

// InstanceMetadataOptions specifies various Instance creation fields
// that relate to the Linode Metadata service.
type InstanceMetadataOptions struct {
	// UserData expects a Base64-encoded string
	UserData string `json:"user_data,omitempty"`
}

// InstancePasswordResetOptions specifies the new password for the Linode
type InstancePasswordResetOptions struct {
	RootPass string `json:"root_pass"`
}

// InstanceCreateOptions require only Region and Type
type InstanceCreateOptions struct {
	Region string `json:"region"`
	Type   string `json:"type"`
	Label  string `json:"label,omitempty"`

	StackScriptID   int               `json:"stackscript_id,omitempty"`
	StackScriptData map[string]string `json:"stackscript_data,omitempty"`
	BackupID        int               `json:"backup_id,omitempty"`

	// Image is optional, but requires at least one of RootPass, AuthorizedUsers, or AuthorizedKeys if provided
	Image string `json:"image,omitempty"`

	RootPass        string   `json:"root_pass,omitempty"`
	AuthorizedKeys  []string `json:"authorized_keys,omitempty"`
	AuthorizedUsers []string `json:"authorized_users,omitempty"`

	BackupsEnabled      bool                                 `json:"backups_enabled,omitempty"`
	PrivateIP           bool                                 `json:"private_ip,omitempty"`
	NetworkHelper       *bool                                `json:"network_helper,omitempty"`
	Tags                []string                             `json:"tags,omitempty"`
	Metadata            *InstanceMetadataOptions             `json:"metadata,omitempty"`
	FirewallID          int                                  `json:"firewall_id,omitempty"`
	InterfaceGeneration InterfaceGeneration                  `json:"interface_generation,omitempty"`
	DiskEncryption      InstanceDiskEncryption               `json:"disk_encryption,omitempty"`
	PlacementGroup      *InstanceCreatePlacementGroupOptions `json:"placement_group,omitempty"`

	// Linode Interfaces to create the new instance with.
	// Conflicts with Interfaces.
	LinodeInterfaces []LinodeInterfaceCreateOptions `json:"-"`

	// Legacy (config) Interfaces to create the new instance with.
	// Conflicts with LinodeInterfaces.
	Interfaces []InstanceConfigInterfaceCreateOptions `json:"-"`

	// Creation fields that need to be set explicitly false, "", or 0 use pointers
	SwapSize *int  `json:"swap_size,omitempty"`
	Booted   *bool `json:"booted,omitempty"`

	// Deprecated: group is a deprecated property denoting a group label for the Linode.
	Group string `json:"group,omitempty"`

	IPv4 []string `json:"ipv4,omitempty"`

	MaintenancePolicy *string `json:"maintenance_policy,omitempty"`
	Kernel            *string `json:"kernel,omitzero"`
	BootSize          *int    `json:"boot_size,omitzero"`
}

// InstanceCreatePlacementGroupOptions represents the placement group
// to create this Linode under.
type InstanceCreatePlacementGroupOptions struct {
	ID            int   `json:"id"`
	CompliantOnly *bool `json:"compliant_only,omitempty"`
}

// InstanceUpdateOptions is an options struct used when Updating an Instance
type InstanceUpdateOptions struct {
	Label           string          `json:"label,omitempty"`
	Backups         *InstanceBackup `json:"backups,omitempty"`
	Alerts          *InstanceAlert  `json:"alerts,omitempty"`
	WatchdogEnabled *bool           `json:"watchdog_enabled,omitempty"`
	Tags            *[]string       `json:"tags,omitempty"`

	// Deprecated: group is a deprecated property denoting a group label for the Linode.
	Group *string `json:"group,omitempty"`

	MaintenancePolicy *string `json:"maintenance_policy,omitempty"`
}

// MarshalJSON contains logic necessary to populate the `interfaces` field of
// InstanceCreateOptions depending on whether Interfaces or LinodeInterfaces
// is specified.
func (i InstanceCreateOptions) MarshalJSON() ([]byte, error) {
	type Mask InstanceCreateOptions

	resultData := struct {
		*Mask

		Interfaces any `json:"interfaces,omitempty"`
	}{
		Mask:       (*Mask)(&i),
		Interfaces: nil,
	}

	if i.Interfaces != nil && i.LinodeInterfaces != nil {
		return nil, fmt.Errorf("fields Interfaces and LinodeInterfaces cannot be specified together")
	}

	if i.Interfaces != nil {
		resultData.Interfaces = i.Interfaces
	}

	if i.LinodeInterfaces != nil {
		resultData.Interfaces = i.LinodeInterfaces
	}

	return json.Marshal(resultData)
}

// UnmarshalJSON contains logic necessary to populate the Interfaces field
// depending on the value of interface_generation.
func (i *InstanceCreateOptions) UnmarshalJSON(b []byte) error {
	type Mask InstanceCreateOptions

	p := struct {
		*Mask

		GenericInterfaces any `json:"interfaces,omitempty"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	if p.GenericInterfaces == nil {
		// No interfaces were given - nothing to do here.
		return nil
	}

	if i.InterfaceGeneration == GenerationLinode {
		data := struct {
			Interfaces []LinodeInterfaceCreateOptions `json:"interfaces"`
		}{}

		err := json.Unmarshal(b, &data)
		i.LinodeInterfaces = data.Interfaces

		return err
	}

	if i.InterfaceGeneration == GenerationLegacyConfig {
		data := struct {
			Interfaces []InstanceConfigInterfaceCreateOptions `json:"interfaces"`
		}{}

		err := json.Unmarshal(b, &data)
		i.Interfaces = data.Interfaces

		return err
	}

	return fmt.Errorf("cannot unmarshal interfaces without valid value for interface_generation")
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *Instance) UnmarshalJSON(b []byte) error {
	type Mask Instance

	p := struct {
		*Mask

		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (backup *InstanceBackup) UnmarshalJSON(b []byte) error {
	type Mask InstanceBackup

	p := struct {
		*Mask

		LastSuccessful *parseabletime.ParseableTime `json:"last_successful"`
	}{
		Mask: (*Mask)(backup),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	backup.LastSuccessful = (*time.Time)(p.LastSuccessful)

	return nil
}

// GetUpdateOptions converts an Instance to InstanceUpdateOptions for use in UpdateInstance
func (i *Instance) GetUpdateOptions() InstanceUpdateOptions {
	return InstanceUpdateOptions{
		Label:             i.Label,
		Group:             &i.Group,
		Backups:           i.Backups,
		Alerts:            i.Alerts,
		WatchdogEnabled:   &i.WatchdogEnabled,
		Tags:              &i.Tags,
		MaintenancePolicy: &i.MaintenancePolicy,
	}
}

// InstanceCloneOptions is an options struct sent when Cloning an Instance
type InstanceCloneOptions struct {
	Region string `json:"region,omitempty"`
	Type   string `json:"type,omitempty"`

	// LinodeID is an optional existing instance to use as the target of the clone
	LinodeID       int                                  `json:"linode_id,omitempty"`
	Label          string                               `json:"label,omitempty"`
	BackupsEnabled bool                                 `json:"backups_enabled"`
	Disks          []int                                `json:"disks,omitempty"`
	Configs        []int                                `json:"configs,omitempty"`
	PrivateIP      bool                                 `json:"private_ip,omitempty"`
	Metadata       *InstanceMetadataOptions             `json:"metadata,omitempty"`
	PlacementGroup *InstanceCreatePlacementGroupOptions `json:"placement_group,omitempty"`

	// Deprecated: group is a deprecated property denoting a group label for the Linode.
	Group string `json:"group,omitempty"`
}

// InstanceResizeOptions is an options struct used when resizing an instance
type InstanceResizeOptions struct {
	Type          string                `json:"type"`
	MigrationType InstanceMigrationType `json:"migration_type,omitempty"`

	// When enabled, an instance resize will also resize a data disk if the instance has no more than one data disk and one swap disk
	AllowAutoDiskResize *bool `json:"allow_auto_disk_resize,omitempty"`
}

// InstanceMigrateOptions is an options struct used when migrating an instance
type InstanceMigrateOptions struct {
	Type    InstanceMigrationType `json:"type,omitempty"`
	Region  string                `json:"region,omitempty"`
	Upgrade *bool                 `json:"upgrade,omitempty"`

	PlacementGroup *InstanceCreatePlacementGroupOptions `json:"placement_group,omitempty"`
}

// ListInstances lists linode instances
func (c *Client) ListInstances(ctx context.Context, opts *ListOptions) ([]Instance, error) {
	return getPaginatedResults[Instance](ctx, c, "linode/instances", opts)
}

// GetInstance gets the instance with the provided ID
func (c *Client) GetInstance(ctx context.Context, linodeID int) (*Instance, error) {
	e := formatAPIPath("linode/instances/%d", linodeID)
	return doGETRequest[Instance](ctx, c, e)
}

// GetInstanceTransfer gets the instance's network transfer pool statistics for the current month.
func (c *Client) GetInstanceTransfer(ctx context.Context, linodeID int) (*InstanceTransfer, error) {
	e := formatAPIPath("linode/instances/%d/transfer", linodeID)
	return doGETRequest[InstanceTransfer](ctx, c, e)
}

// GetInstanceTransferMonthly gets the instance's network transfer pool statistics for a specific month.
func (c *Client) GetInstanceTransferMonthly(ctx context.Context, linodeID, year, month int) (*MonthlyInstanceTransferStats, error) {
	e := formatAPIPath("linode/instances/%d/transfer/%d/%d", linodeID, year, month)
	return doGETRequest[MonthlyInstanceTransferStats](ctx, c, e)
}

// GetInstanceTransferMonthlyV2 gets the instance's network transfer pool statistics for a specific month.
func (c *Client) GetInstanceTransferMonthlyV2(ctx context.Context, linodeID, year, month int) (*MonthlyInstanceTransferStatsV2, error) {
	e := formatAPIPath("linode/instances/%d/transfer/%d/%d", linodeID, year, month)
	return doGETRequest[MonthlyInstanceTransferStatsV2](ctx, c, e)
}

// CreateInstance creates a Linode instance
func (c *Client) CreateInstance(ctx context.Context, opts InstanceCreateOptions) (*Instance, error) {
	return doPOSTRequest[Instance](ctx, c, "linode/instances", opts)
}

// UpdateInstance updates a Linode instance
func (c *Client) UpdateInstance(ctx context.Context, linodeID int, opts InstanceUpdateOptions) (*Instance, error) {
	e := formatAPIPath("linode/instances/%d", linodeID)
	return doPUTRequest[Instance](ctx, c, e, opts)
}

// RenameInstance renames an Instance
func (c *Client) RenameInstance(ctx context.Context, linodeID int, label string) (*Instance, error) {
	return c.UpdateInstance(ctx, linodeID, InstanceUpdateOptions{Label: label})
}

// DeleteInstance deletes a Linode instance
func (c *Client) DeleteInstance(ctx context.Context, linodeID int) error {
	e := formatAPIPath("linode/instances/%d", linodeID)
	return doDELETERequest(ctx, c, e)
}

// BootInstance will boot a Linode instance
// A configID of 0 will cause Linode to choose the last/best config
func (c *Client) BootInstance(ctx context.Context, linodeID int, configID int) error {
	opts := make(map[string]int)

	if configID != 0 {
		opts = map[string]int{"config_id": configID}
	}

	e := formatAPIPath("linode/instances/%d/boot", linodeID)

	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// CloneInstance clone an existing Instances Disks and Configuration profiles to another Linode Instance
func (c *Client) CloneInstance(ctx context.Context, linodeID int, opts InstanceCloneOptions) (*Instance, error) {
	e := formatAPIPath("linode/instances/%d/clone", linodeID)
	return doPOSTRequest[Instance](ctx, c, e, opts)
}

// ResetInstancePassword resets a Linode instance's root password
func (c *Client) ResetInstancePassword(ctx context.Context, linodeID int, opts InstancePasswordResetOptions) error {
	e := formatAPIPath("linode/instances/%d/password", linodeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// RebootInstance reboots a Linode instance
// A configID of 0 will cause Linode to choose the last/best config
func (c *Client) RebootInstance(ctx context.Context, linodeID int, configID int) error {
	opts := make(map[string]int)

	if configID != 0 {
		opts = map[string]int{"config_id": configID}
	}

	e := formatAPIPath("linode/instances/%d/reboot", linodeID)

	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// InstanceRebuildOptions is a struct representing the options to send to the rebuild linode endpoint
type InstanceRebuildOptions struct {
	// Image is the image to deploy when rebuilding the Linode
	Image string `json:"image,omitempty"`

	RootPass        string   `json:"root_pass,omitempty"`
	AuthorizedKeys  []string `json:"authorized_keys,omitempty"`
	AuthorizedUsers []string `json:"authorized_users,omitempty"`

	StackScriptID   int                      `json:"stackscript_id,omitempty"`
	StackScriptData map[string]string        `json:"stackscript_data,omitempty"`
	Booted          *bool                    `json:"booted,omitempty"`
	Metadata        *InstanceMetadataOptions `json:"metadata,omitempty"`
	Type            string                   `json:"type,omitempty"`
	DiskEncryption  InstanceDiskEncryption   `json:"disk_encryption,omitempty"`
}

// RebuildInstance Deletes all Disks and Configs on this Linode,
// then deploys a new Image to this Linode with the given attributes.
func (c *Client) RebuildInstance(ctx context.Context, linodeID int, opts InstanceRebuildOptions) (*Instance, error) {
	e := formatAPIPath("linode/instances/%d/rebuild", linodeID)
	return doPOSTRequest[Instance](ctx, c, e, opts)
}

// InstanceRescueOptions fields are those accepted by RescueInstance
type InstanceRescueOptions struct {
	Devices InstanceConfigDeviceMap `json:"devices"`
}

// RescueInstance reboots an instance into a safe environment for performing many system recovery and disk management tasks.
// Rescue Mode is based on the Finnix recovery distribution, a self-contained and bootable Linux distribution.
// You can also use Rescue Mode for tasks other than disaster recovery, such as formatting disks to use different filesystems,
// copying data between disks, and downloading files from a disk via SSH and SFTP.
func (c *Client) RescueInstance(ctx context.Context, linodeID int, opts InstanceRescueOptions) error {
	e := formatAPIPath("linode/instances/%d/rescue", linodeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// ResizeInstance resizes an instance to new Linode type
func (c *Client) ResizeInstance(ctx context.Context, linodeID int, opts InstanceResizeOptions) error {
	e := formatAPIPath("linode/instances/%d/resize", linodeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// ShutdownInstance - Shutdown an instance
func (c *Client) ShutdownInstance(ctx context.Context, id int) error {
	return c.simpleInstanceAction(ctx, "shutdown", id)
}

// Deprecated: Please use UpgradeInstance instead.
// MutateInstance Upgrades a Linode to its next generation.
func (c *Client) MutateInstance(ctx context.Context, id int) error {
	return c.simpleInstanceAction(ctx, "mutate", id)
}

// InstanceUpgradeOptions is a struct representing the options for upgrading a Linode
type InstanceUpgradeOptions struct {
	// Automatically resize disks when resizing a Linode.
	// When resizing down to a smaller plan your Linode's data must fit within the smaller disk size.
	AllowAutoDiskResize bool `json:"allow_auto_disk_resize"`
}

// UpgradeInstance upgrades a Linode to its next generation.
func (c *Client) UpgradeInstance(ctx context.Context, linodeID int, opts InstanceUpgradeOptions) error {
	e := formatAPIPath("linode/instances/%d/mutate", linodeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// MigrateInstance - Migrate an instance
func (c *Client) MigrateInstance(ctx context.Context, linodeID int, opts InstanceMigrateOptions) error {
	e := formatAPIPath("linode/instances/%d/migrate", linodeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// simpleInstanceAction is a helper for Instance actions that take no parameters
// and return empty responses `{}` unless they return a standard error
func (c *Client) simpleInstanceAction(ctx context.Context, action string, linodeID int) error {
	e := formatAPIPath("linode/instances/%d/%s", linodeID, action)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}
