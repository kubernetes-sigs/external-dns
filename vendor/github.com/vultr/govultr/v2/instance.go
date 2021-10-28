package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const instancePath = "/v2/instances"

// InstanceService is the interface to interact with the instance endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/instances
type InstanceService interface {
	Create(ctx context.Context, instanceReq *InstanceCreateReq) (*Instance, error)
	Get(ctx context.Context, instanceID string) (*Instance, error)
	Update(ctx context.Context, instanceID string, instanceReq *InstanceUpdateReq) error
	Delete(ctx context.Context, instanceID string) error
	List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, error)

	Start(ctx context.Context, instanceID string) error
	Halt(ctx context.Context, instanceID string) error
	Reboot(ctx context.Context, instanceID string) error
	Reinstall(ctx context.Context, instanceID string) error

	MassStart(ctx context.Context, instanceList []string) error
	MassHalt(ctx context.Context, instanceList []string) error
	MassReboot(ctx context.Context, instanceList []string) error

	Restore(ctx context.Context, instanceID string, restoreReq *RestoreReq) error

	GetBandwidth(ctx context.Context, instanceID string) (*Bandwidth, error)
	GetNeighbors(ctx context.Context, instanceID string) (*Neighbors, error)

	ListPrivateNetworks(ctx context.Context, instanceID string, options *ListOptions) ([]PrivateNetwork, *Meta, error)
	AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error
	DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error

	ISOStatus(ctx context.Context, instanceID string) (*Iso, error)
	AttachISO(ctx context.Context, instanceID, isoID string) error
	DetachISO(ctx context.Context, instanceID string) error

	GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, error)
	SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) error

	CreateIPv4(ctx context.Context, instanceID string, reboot *bool) (*IPv4, error)
	ListIPv4(ctx context.Context, instanceID string, option *ListOptions) ([]IPv4, *Meta, error)
	DeleteIPv4(ctx context.Context, instanceID, ip string) error
	ListIPv6(ctx context.Context, instanceID string, option *ListOptions) ([]IPv6, *Meta, error)

	CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error
	ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, error)
	DeleteReverseIPv6(ctx context.Context, instanceID, ip string) error

	CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error
	DefaultReverseIPv4(ctx context.Context, instanceID, ip string) error

	GetUserData(ctx context.Context, instanceID string) (*UserData, error)

	GetUpgrades(ctx context.Context, instanceID string) (*Upgrades, error)
}

// InstanceServiceHandler handles interaction with the server methods for the Vultr API
type InstanceServiceHandler struct {
	client *Client
}

// Instance represents a VPS
type Instance struct {
	ID               string   `json:"id"`
	Os               string   `json:"os"`
	RAM              int      `json:"ram"`
	Disk             int      `json:"disk"`
	Plan             string   `json:"plan"`
	MainIP           string   `json:"main_ip"`
	VCPUCount        int      `json:"vcpu_count"`
	Region           string   `json:"region"`
	DefaultPassword  string   `json:"default_password,omitempty"`
	DateCreated      string   `json:"date_created"`
	Status           string   `json:"status"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6NetworkSize    int      `json:"v6_network_size"`
	Label            string   `json:"label"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	Tag              string   `json:"tag"`
	OsID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	ImageID          string   `json:"image_id"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	Features         []string `json:"features"`
}

type instanceBase struct {
	Instance *Instance `json:"instance"`
}

type ipv4Base struct {
	IPv4 *IPv4 `json:"ipv4"`
}

type instancesBase struct {
	Instances []Instance `json:"instances"`
	Meta      *Meta      `json:"meta"`
}

// Neighbors that might exist on the same host.
type Neighbors struct {
	Neighbors []string `json:"neighbors"`
}

// Bandwidth used on a given instance.
type Bandwidth struct {
	Bandwidth map[string]struct {
		IncomingBytes int `json:"incoming_bytes"`
		OutgoingBytes int `json:"outgoing_bytes"`
	} `json:"bandwidth"`
}

type privateNetworksBase struct {
	PrivateNetworks []PrivateNetwork `json:"private_networks"`
	Meta            *Meta            `json:"meta"`
}

// PrivateNetwork information for a given instance.
type PrivateNetwork struct {
	NetworkID  string `json:"network_id"`
	MacAddress string `json:"mac_address"`
	IPAddress  string `json:"ip_address"`
}

type isoStatusBase struct {
	IsoStatus *Iso `json:"iso_status"`
}

// Iso information for a given instance.
type Iso struct {
	State string `json:"state"`
	IsoID string `json:"iso_id"`
}

type backupScheduleBase struct {
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
}

// BackupSchedule information for a given instance.
type BackupSchedule struct {
	Enabled             *bool  `json:"enabled,omitempty"`
	Type                string `json:"type,omitempty"`
	NextScheduleTimeUTC string `json:"next_scheduled_time_utc,omitempty"`
	Hour                int    `json:"hour,omitempty"`
	Dow                 int    `json:"dow,omitempty"`
	Dom                 int    `json:"dom,omitempty"`
}

// BackupScheduleReq struct used to create a backup schedule for an instance.
type BackupScheduleReq struct {
	Type string `json:"type"`
	Hour *int   `json:"hour,omitempty"`
	Dow  *int   `json:"dow,omitempty"`
	Dom  int    `json:"dom,omitempty"`
}

// RestoreReq struct used to supply whether a restore should be from a backup or snapshot.
type RestoreReq struct {
	BackupID   string `json:"backup_id,omitempty"`
	SnapshotID string `json:"snapshot_id,omitempty"`
}

// todo can we remove this list and return this data back in the list?
type reverseIPv6sBase struct {
	ReverseIPv6s []ReverseIP `json:"reverse_ipv6s"`
	// no meta?
}

// ReverseIP information for a given instance.
type ReverseIP struct {
	IP      string `json:"ip"`
	Reverse string `json:"reverse"`
}

type userDataBase struct {
	UserData *UserData `json:"user_data"`
}

// UserData information for a given struct.
type UserData struct {
	Data string `json:"data"`
}

type upgradeBase struct {
	Upgrades *Upgrades `json:"upgrades"`
}

// Upgrades that are available for a given Instance.
type Upgrades struct {
	Applications []Application `json:"applications,omitempty"`
	OS           []OS          `json:"os,omitempty"`
	Plans        []string      `json:"plans,omitempty"`
}

// InstanceCreateReq struct used to create an instance.
type InstanceCreateReq struct {
	Region               string   `json:"region,omitempty"`
	Plan                 string   `json:"plan,omitempty"`
	Label                string   `json:"label,omitempty"`
	Tag                  string   `json:"tag,omitempty"`
	OsID                 int      `json:"os_id,omitempty"`
	ISOID                string   `json:"iso_id,omitempty"`
	AppID                int      `json:"app_id,omitempty"`
	ImageID              string   `json:"image_id,omitempty"`
	FirewallGroupID      string   `json:"firewall_group_id,omitempty"`
	Hostname             string   `json:"hostname,omitempty"`
	IPXEChainURL         string   `json:"ipxe_chain_url,omitempty"`
	ScriptID             string   `json:"script_id,omitempty"`
	SnapshotID           string   `json:"snapshot_id,omitempty"`
	EnableIPv6           *bool    `json:"enable_ipv6,omitempty"`
	EnablePrivateNetwork *bool    `json:"enable_private_network,omitempty"`
	AttachPrivateNetwork []string `json:"attach_private_network,omitempty"`
	SSHKeys              []string `json:"sshkey_id,omitempty"`
	Backups              string   `json:"backups,omitempty"`
	DDOSProtection       *bool    `json:"ddos_protection,omitempty"`
	UserData             string   `json:"user_data,omitempty"`
	ReservedIPv4         string   `json:"reserved_ipv4,omitempty"`
	ActivationEmail      *bool    `json:"activation_email,omitempty"`
}

// InstanceUpdateReq struct used to update an instance.
type InstanceUpdateReq struct {
	Plan                 string   `json:"plan,omitempty"`
	Label                string   `json:"label,omitempty"`
	Tag                  string   `json:"tag,omitempty"`
	OsID                 int      `json:"os_id,omitempty"`
	AppID                int      `json:"app_id,omitempty"`
	ImageID              string   `json:"image_id,omitempty"`
	EnableIPv6           *bool    `json:"enable_ipv6,omitempty"`
	EnablePrivateNetwork *bool    `json:"enable_private_network,omitempty"`
	AttachPrivateNetwork []string `json:"attach_private_network,omitempty"`
	DetachPrivateNetwork []string `json:"detach_private_network,omitempty"`
	Backups              string   `json:"backups,omitempty"`
	DDOSProtection       *bool    `json:"ddos_protection"`
	UserData             string   `json:"user_data,omitempty"`
	FirewallGroupID      string   `json:"firewall_group_id,omitempty"`
}

// Create will create the server with the given parameters
func (i *InstanceServiceHandler) Create(ctx context.Context, instanceReq *InstanceCreateReq) (*Instance, error) {
	uri := fmt.Sprintf("%s", instancePath)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, instanceReq)
	if err != nil {
		return nil, err
	}

	instance := new(instanceBase)
	if err = i.client.DoWithContext(ctx, req, instance); err != nil {
		return nil, err
	}

	return instance.Instance, nil
}

// Get will get the server with the given instanceID
func (i *InstanceServiceHandler) Get(ctx context.Context, instanceID string) (*Instance, error) {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	instance := new(instanceBase)
	if err = i.client.DoWithContext(ctx, req, instance); err != nil {
		return nil, err
	}

	return instance.Instance, nil
}

// Update will update the server with the given parameters
func (i *InstanceServiceHandler) Update(ctx context.Context, instanceID string, instanceReq *InstanceUpdateReq) error {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPatch, uri, instanceReq)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// Delete an instance. All data will be permanently lost, and the IP address will be released
func (i *InstanceServiceHandler) Delete(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// List all instances on your account.
func (i *InstanceServiceHandler) List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, instancePath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	instances := new(instancesBase)
	if err = i.client.DoWithContext(ctx, req, instances); err != nil {
		return nil, nil, err
	}

	return instances.Instances, instances.Meta, nil
}

// Start will start a vps instance the machine is already running, it will be restarted.
func (i *InstanceServiceHandler) Start(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/start", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// Halt will pause an instance.
func (i *InstanceServiceHandler) Halt(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/halt", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// Reboot an instance.
func (i *InstanceServiceHandler) Reboot(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/reboot", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// Reinstall an instance.
func (i *InstanceServiceHandler) Reinstall(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/reinstall", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// MassStart will start a list of instances the machine is already running, it will be restarted.
func (i *InstanceServiceHandler) MassStart(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/start", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// MassHalt will pause a list of instances.
func (i *InstanceServiceHandler) MassHalt(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/halt", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// MassReboot reboots a list of instances.
func (i *InstanceServiceHandler) MassReboot(ctx context.Context, instanceList []string) error {
	uri := fmt.Sprintf("%s/reboot", instancePath)

	reqBody := RequestBody{"instance_ids": instanceList}
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// Restore an instance.
func (i *InstanceServiceHandler) Restore(ctx context.Context, instanceID string, restoreReq *RestoreReq) error {
	uri := fmt.Sprintf("%s/%s/restore", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// GetBandwidth for a given instance.
func (i *InstanceServiceHandler) GetBandwidth(ctx context.Context, instanceID string) (*Bandwidth, error) {
	uri := fmt.Sprintf("%s/%s/bandwidth", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	bandwidth := new(Bandwidth)
	if err = i.client.DoWithContext(ctx, req, bandwidth); err != nil {
		return nil, err
	}

	return bandwidth, nil
}

// GetNeighbors gets a list of other instances in the same location as this Instance.
func (i *InstanceServiceHandler) GetNeighbors(ctx context.Context, instanceID string) (*Neighbors, error) {
	uri := fmt.Sprintf("%s/%s/neighbors", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	neighbors := new(Neighbors)
	if err = i.client.DoWithContext(ctx, req, neighbors); err != nil {
		return nil, err
	}

	return neighbors, nil
}

// ListPrivateNetworks currently attached to an instance.
func (i *InstanceServiceHandler) ListPrivateNetworks(ctx context.Context, instanceID string, options *ListOptions) ([]PrivateNetwork, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/private-networks", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	networks := new(privateNetworksBase)
	if err = i.client.DoWithContext(ctx, req, networks); err != nil {
		return nil, nil, err
	}

	return networks.PrivateNetworks, networks.Meta, nil
}

// AttachPrivateNetwork to an instance
func (i *InstanceServiceHandler) AttachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-networks/attach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// DetachPrivateNetwork from an instance.
func (i *InstanceServiceHandler) DetachPrivateNetwork(ctx context.Context, instanceID, networkID string) error {
	uri := fmt.Sprintf("%s/%s/private-networks/detach", instancePath, instanceID)
	body := RequestBody{"network_id": networkID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// ISOStatus retrieves the current ISO state for a given VPS.
// The returned state may be one of: ready | isomounting | isomounted.
func (i *InstanceServiceHandler) ISOStatus(ctx context.Context, instanceID string) (*Iso, error) {
	uri := fmt.Sprintf("%s/%s/iso", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	iso := new(isoStatusBase)
	if err = i.client.DoWithContext(ctx, req, iso); err != nil {
		return nil, err
	}
	return iso.IsoStatus, nil
}

// AttachISO will attach an ISO to the given instance and reboot it
func (i *InstanceServiceHandler) AttachISO(ctx context.Context, instanceID, isoID string) error {
	uri := fmt.Sprintf("%s/%s/iso/attach", instancePath, instanceID)
	body := RequestBody{"iso_id": isoID}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// DetachISO will detach the currently mounted ISO and reboot the instance.
func (i *InstanceServiceHandler) DetachISO(ctx context.Context, instanceID string) error {
	uri := fmt.Sprintf("%s/%s/iso/detach", instancePath, instanceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// GetBackupSchedule retrieves the backup schedule for a given instance - all time values are in UTC
func (i *InstanceServiceHandler) GetBackupSchedule(ctx context.Context, instanceID string) (*BackupSchedule, error) {
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	backup := new(backupScheduleBase)
	if err = i.client.DoWithContext(ctx, req, backup); err != nil {
		return nil, err
	}

	return backup.BackupSchedule, nil
}

// SetBackupSchedule sets the backup schedule for a given instance - all time values are in UTC.
func (i *InstanceServiceHandler) SetBackupSchedule(ctx context.Context, instanceID string, backup *BackupScheduleReq) error {
	uri := fmt.Sprintf("%s/%s/backup-schedule", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, backup)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// CreateIPv4 an additional IPv4 address for given instance.
func (i *InstanceServiceHandler) CreateIPv4(ctx context.Context, instanceID string, reboot *bool) (*IPv4, error) {
	uri := fmt.Sprintf("%s/%s/ipv4", instancePath, instanceID)

	body := RequestBody{"reboot": reboot}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	ip := new(ipv4Base)
	if err = i.client.DoWithContext(ctx, req, ip); err != nil {
		return nil, err
	}

	return ip.IPv4, nil
}

// ListIPv4 addresses that are currently assigned to a given instance.
func (i *InstanceServiceHandler) ListIPv4(ctx context.Context, instanceID string, options *ListOptions) ([]IPv4, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv4", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	if err = i.client.DoWithContext(ctx, req, ips); err != nil {
		return nil, nil, err
	}

	return ips.IPv4s, ips.Meta, nil
}

// DeleteIPv4 address from a given instance.
func (i *InstanceServiceHandler) DeleteIPv4(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv4/%s", instancePath, instanceID, ip)
	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// ListIPv6 addresses that are currently assigned to a given instance.
func (i *InstanceServiceHandler) ListIPv6(ctx context.Context, instanceID string, options *ListOptions) ([]IPv6, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/ipv6", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	ips := new(ipBase)
	if err = i.client.DoWithContext(ctx, req, ips); err != nil {
		return nil, nil, err
	}

	return ips.IPv6s, ips.Meta, nil
}

// CreateReverseIPv6 for a given instance.
func (i *InstanceServiceHandler) CreateReverseIPv6(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// ListReverseIPv6 currently assigned to a given instance.
func (i *InstanceServiceHandler) ListReverseIPv6(ctx context.Context, instanceID string) ([]ReverseIP, error) {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	reverse := new(reverseIPv6sBase)
	if err = i.client.DoWithContext(ctx, req, reverse); err != nil {
		return nil, err
	}

	return reverse.ReverseIPv6s, nil
}

// DeleteReverseIPv6 a given reverse IPv6.
func (i *InstanceServiceHandler) DeleteReverseIPv6(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv6/reverse/%s", instancePath, instanceID, ip)
	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// CreateReverseIPv4 for a given IP on a given instance.
func (i *InstanceServiceHandler) CreateReverseIPv4(ctx context.Context, instanceID string, reverseReq *ReverseIP) error {
	uri := fmt.Sprintf("%s/%s/ipv4/reverse", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reverseReq)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// DefaultReverseIPv4 will set the IPs reverse setting back to the original one supplied by Vultr.
func (i *InstanceServiceHandler) DefaultReverseIPv4(ctx context.Context, instanceID, ip string) error {
	uri := fmt.Sprintf("%s/%s/ipv4/reverse/default", instancePath, instanceID)
	reqBody := RequestBody{"ip": ip}

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// GetUserData from given instance. The userdata returned will be in base64 encoding.
func (i *InstanceServiceHandler) GetUserData(ctx context.Context, instanceID string) (*UserData, error) {
	uri := fmt.Sprintf("%s/%s/user-data", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	userData := new(userDataBase)
	if err = i.client.DoWithContext(ctx, req, userData); err != nil {
		return nil, err
	}

	return userData.UserData, nil
}

// GetUpgrades that are available for a given instance.
func (i *InstanceServiceHandler) GetUpgrades(ctx context.Context, instanceID string) (*Upgrades, error) {
	uri := fmt.Sprintf("%s/%s/upgrades", instancePath, instanceID)
	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	upgrades := new(upgradeBase)
	if err = i.client.DoWithContext(ctx, req, upgrades); err != nil {
		return nil, err
	}

	return upgrades.Upgrades, nil
}
