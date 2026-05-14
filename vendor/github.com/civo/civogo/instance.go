package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/civo/civogo/utils"
)

const (
	// DefaultInstanceUser is the default username used in newly created instances.
	DefaultInstanceUser string = "civo"
)

// Instance represents a virtual server within Civo's infrastructure
type Instance struct {
	ID                       string           `json:"id,omitempty"`
	OpenstackServerID        string           `json:"openstack_server_id,omitempty"`
	Hostname                 string           `json:"hostname,omitempty"`
	ReverseDNS               string           `json:"reverse_dns,omitempty"`
	Size                     string           `json:"size,omitempty"`
	Region                   string           `json:"region,omitempty"`
	NetworkID                string           `json:"network_id,omitempty"`
	PrivateIP                string           `json:"private_ip,omitempty"`
	PublicIP                 string           `json:"public_ip,omitempty"`
	IPv6                     string           `json:"ipv6,omitempty"`
	PseudoIP                 string           `json:"pseudo_ip,omitempty"`
	TemplateID               string           `json:"template_id,omitempty"`
	SourceType               string           `json:"source_type,omitempty"`
	SourceID                 string           `json:"source_id,omitempty"`
	SnapshotID               string           `json:"snapshot_id,omitempty"`
	InitialUser              string           `json:"initial_user,omitempty"`
	InitialPassword          string           `json:"initial_password,omitempty"`
	SSHKey                   string           `json:"ssh_key,omitempty"`
	SSHKeyID                 string           `json:"ssh_key_id,omitempty"`
	Status                   string           `json:"status,omitempty"`
	Notes                    string           `json:"notes,omitempty"`
	FirewallID               string           `json:"firewall_id,omitempty"`
	Tags                     []string         `json:"tags,omitempty"`
	CivostatsdToken          string           `json:"civostatsd_token,omitempty"`
	CivostatsdStats          string           `json:"civostatsd_stats,omitempty"`
	CivostatsdStatsPerMinute []string         `json:"civostatsd_stats_per_minute,omitempty"`
	CivostatsdStatsPerHour   []string         `json:"civostatsd_stats_per_hour,omitempty"`
	OpenstackImageID         string           `json:"openstack_image_id,omitempty"`
	RescuePassword           string           `json:"rescue_password,omitempty"`
	VolumeBacked             bool             `json:"volume_backed,omitempty"`
	CPUCores                 int              `json:"cpu_cores,omitempty"`
	RAMMegabytes             int              `json:"ram_mb,omitempty"`
	DiskGigabytes            int              `json:"disk_gb,omitempty"`
	GPUCount                 int              `json:"gpu_count,omitempty"`
	GPUType                  string           `json:"gpu_type,omitempty"`
	Script                   string           `json:"script,omitempty"`
	CreatedAt                time.Time        `json:"created_at,omitempty"`
	ReservedIPID             string           `json:"reserved_ip_id,omitempty"`
	ReservedIPName           string           `json:"reserved_ip_name,omitempty"`
	ReservedIP               string           `json:"reserved_ip,omitempty"`
	VolumeType               string           `json:"volume_type,omitempty"`
	Subnets                  []Subnet         `json:"subnets,omitempty"`
	AttachedVolumes          []AttachedVolume `json:"attached_volumes,omitempty"`
	PlacementRule            PlacementRule    `json:"placement_rule,omitempty"`
	NetworkBandwidthLimit    int              `json:"network_bandwidth_limit,omitempty"`
	AllowedIPs               []string         `json:"allowed_ips,omitempty"`
}

// InstanceVnc represents VNC information for an instances
type InstanceVnc struct {
	URI        string `json:"uri,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}

// CreateInstanceVncResp represents VNC information for a new instance console
type CreateInstanceVncResp struct {
	URI      string `json:"uri,omitempty"`
	Duration string `json:"duration,omitempty"`
}

// PaginatedInstanceList returns a paginated list of Instance object
type PaginatedInstanceList struct {
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
	Pages   int        `json:"pages"`
	Items   []Instance `json:"items"`
}

// AttachedVolume disk information
type AttachedVolume struct {
	// ID of the volume to attach
	ID string `json:"id"`
}

// InstanceConfig describes the parameters for a new instance
// none of the fields are mandatory and will be automatically
// set with default values
type InstanceConfig struct {
	Count                 int              `json:"count"`
	Hostname              string           `json:"hostname"`
	ReverseDNS            string           `json:"reverse_dns"`
	Size                  string           `json:"size"`
	Region                string           `json:"region"`
	PublicIPRequired      string           `json:"public_ip"`
	ReservedIPv4          string           `json:"reserved_ipv4"`
	PrivateIPv4           string           `json:"private_ipv4"`
	NetworkID             string           `json:"network_id"`
	TemplateID            string           `json:"template_id"`
	SourceType            string           `json:"source_type"`
	SourceID              string           `json:"source_id"`
	SnapshotID            string           `json:"snapshot_id"`
	Subnets               []string         `json:"subnets,omitempty"`
	InitialUser           string           `json:"initial_user"`
	SSHKeyID              string           `json:"ssh_key_id"`
	Script                string           `json:"script"`
	Tags                  []string         `json:"-"`
	TagsList              string           `json:"tags"`
	FirewallID            string           `json:"firewall_id"`
	VolumeType            string           `json:"volume_type,omitempty"`
	AttachedVolumes       []AttachedVolume `json:"attached_volumes"`
	PlacementRule         PlacementRule    `json:"placement_rule"`
	NetworkBandwidthLimit int              `json:"network_bandwidth_limit,omitempty"`
	AllowedIPs            []string         `json:"allowed_ips,omitempty"`
}

// AffinityRule represents a affinity rule
type AffinityRule struct {
	Type      string   `json:"type"`
	Exclusive bool     `json:"exclusive"`
	Tags      []string `json:"tags"`
}

// PlacementRule represents a placement rule
type PlacementRule struct {
	AffinityRules []AffinityRule    `json:"affinity_rules,omitempty"`
	NodeSelector  map[string]string `json:"node_selector,omitempty"`
}

// ListInstances returns a page of Instances owned by the calling API account
func (c *Client) ListInstances(page int, perPage int) (*PaginatedInstanceList, error) {
	url := "/v2/instances"
	if page != 0 && perPage != 0 {
		url = url + fmt.Sprintf("?page=%d&per_page=%d", page, perPage)
	}

	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	PaginatedInstances := PaginatedInstanceList{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&PaginatedInstances)
	return &PaginatedInstances, err
}

// ListAllInstances returns all (well, upto 99,999,999 instances) Instances owned by the calling API account
func (c *Client) ListAllInstances() ([]Instance, error) {
	instances, err := c.ListInstances(1, 99999999)
	if err != nil {
		return []Instance{}, decodeError(err)
	}

	return instances.Items, nil
}

// FindInstance finds a instance by either part of the ID or part of the hostname
func (c *Client) FindInstance(search string) (*Instance, error) {
	instances, err := c.ListAllInstances()
	if err != nil {
		return nil, decodeError(err)
	}

	partialMatchesCount := 0
	result := Instance{}

	for _, value := range instances {
		if value.Hostname == search || value.ID == search {
			return &value, nil
		} else if strings.Contains(value.Hostname, search) || strings.Contains(value.ID, search) {
			partialMatchesCount++
			result = value
		}
	}

	if partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// GetInstance returns a single Instance by its full ID
func (c *Client) GetInstance(id string) (*Instance, error) {
	resp, err := c.SendGetRequest("/v2/instances/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	instance := Instance{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&instance)
	return &instance, err
}

// NewInstanceConfig returns an initialized config for a new instance
func (c *Client) NewInstanceConfig() (*InstanceConfig, error) {
	network, err := c.GetDefaultNetwork()
	if err != nil {
		return nil, decodeError(err)
	}

	return &InstanceConfig{
		Count:            1,
		Hostname:         utils.RandomName(),
		ReverseDNS:       "",
		Region:           c.Region,
		PublicIPRequired: "true",
		NetworkID:        network.ID,
		SnapshotID:       "",
		InitialUser:      DefaultInstanceUser,
		SSHKeyID:         "",
		Script:           "",
		Tags:             []string{""},
		FirewallID:       "",
	}, nil
}

// CreateInstance creates a new instance in the account
func (c *Client) CreateInstance(config *InstanceConfig) (*Instance, error) {
	config.TagsList = strings.Join(config.Tags, " ")
	body, err := c.SendPostRequest("/v2/instances", config)
	if err != nil {
		return nil, decodeError(err)
	}

	var instance Instance
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}

// SetInstanceTags sets the tags for the specified instance
func (c *Client) SetInstanceTags(i *Instance, tags string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/tags", i.ID), map[string]string{
		"tags":   tags,
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// UpdateInstance updates an Instance's hostname, reverse DNS or notes
func (c *Client) UpdateInstance(i *Instance) (*SimpleResponse, error) {
	params := map[string]interface{}{
		"hostname":    i.Hostname,
		"reverse_dns": i.ReverseDNS,
		"notes":       i.Notes,
		"region":      c.Region,
		"public_ip":   i.PublicIP,
		"subnets":     i.Subnets,
	}

	if i.Notes == "" {
		params["notes_delete"] = "true"
	}

	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s", i.ID), params)
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// GetInstanceVnc enables and gets the VNC information for an instance
// duration is optional and follows Go's duration string format (e.g. "30m", "1h", "24h")
func (c *Client) GetInstanceVnc(id string, duration ...string) (CreateInstanceVncResp, error) {
	url := fmt.Sprintf("/v2/instances/%s/vnc", id)
	if len(duration) > 0 && duration[0] != "" {
		url = fmt.Sprintf("%s?duration=%s", url, duration[0])
	}

	resp, err := c.SendPutRequest(url, map[string]string{
		"region": c.Region,
	})
	vnc := CreateInstanceVncResp{}

	if err != nil {
		return vnc, decodeError(err)
	}

	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&vnc)
	return vnc, err
}

// GetInstanceVncStatus returns the VNC status for an instance
func (c *Client) GetInstanceVncStatus(id string) (*InstanceVnc, error) {
	url := fmt.Sprintf("/v2/instances/%s/vnc", id)
	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	vnc := InstanceVnc{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&vnc)
	return &vnc, err

}

// DeleteInstanceVncSession terminates the VNC session for an instance.
func (c *Client) DeleteInstanceVncSession(id string) (*SimpleResponse, error) {
	url := fmt.Sprintf("/v2/instances/%s/vnc", id)
	resp, err := c.SendDeleteRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}
	return c.DecodeSimpleResponse(resp)
}

// DeleteInstance deletes an instance and frees its resources
func (c *Client) DeleteInstance(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/instances/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// RebootInstance reboots an instance (short version of HardRebootInstance)
func (c *Client) RebootInstance(id string) (*SimpleResponse, error) {
	return c.HardRebootInstance(id)
}

// HardRebootInstance harshly reboots an instance (like shutting the power off and booting it again)
func (c *Client) HardRebootInstance(id string) (*SimpleResponse, error) {
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/instances/%s/hard_reboots", id), map[string]string{
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// SoftRebootInstance requests the VM to shut down nicely
func (c *Client) SoftRebootInstance(id string) (*SimpleResponse, error) {
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/instances/%s/soft_reboots", id), map[string]string{
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// StopInstance shuts the power down to the instance
func (c *Client) StopInstance(id string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/stop", id), map[string]string{
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// StartInstance starts the instance booting from the shutdown state
func (c *Client) StartInstance(id string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/start", id), map[string]string{
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// UpgradeInstance resizes the instance up to the new specification
// it's not possible to resize the instance to a smaller size
func (c *Client) UpgradeInstance(id, newSize string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/resize", id), map[string]string{
		"size":   newSize,
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// MovePublicIPToInstance moves a public IP to the specified instance
func (c *Client) MovePublicIPToInstance(id, ipAddress string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/ip/%s", id, ipAddress), "")
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// SetInstanceFirewall changes the current firewall for an instance
func (c *Client) SetInstanceFirewall(id, firewallID string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/firewall", id), map[string]string{
		"firewall_id": firewallID,
		"region":      c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// EnableRecoveryMode enables recovery mode for the specified instance
func (c *Client) EnableRecoveryMode(id string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/recovery?region=%s", id, c.Region), nil)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DisableRecoveryMode disables recovery mode for the specified instance
func (c *Client) DisableRecoveryMode(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/instances/%s/recovery?region=%s", id, c.Region))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// GetRecoveryStatus gets the recovery status for the specified instance
func (c *Client) GetRecoveryStatus(id string) (*SimpleResponse, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/instances/%s/recovery", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateInstanceAllowedIPs sets the list of IP addresses that an instance is allowed to use
func (c *Client) UpdateInstanceAllowedIPs(id string, allowedIPs []string) (*SimpleResponse, error) {
	// Create a map to match the expected JSON structure
	payload := map[string][]string{
		"allowed_ips": allowedIPs,
	}
	// Send the payload map instead of the raw allowedIPs slice
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/allowed_ips", id), payload)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateInstanceBandwidth sets the list of IP addresses that an instance is allowed to use
func (c *Client) UpdateInstanceBandwidth(id string, bandwidthLimit int) (*SimpleResponse, error) {
	payload := map[string]int{
		"network_bandwidth_limit": bandwidthLimit,
	}
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/instances/%s/network_bandwidth_limit", id), payload)
	if err != nil {
		return nil, decodeError(err)
	}
	return c.DecodeSimpleResponse(resp)
}
