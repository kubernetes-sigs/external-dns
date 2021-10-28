package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals an Instance structure into a temporary structure whose "CreatedAt" field of type
// string to be able to parse the original timestamp (ISO 8601) into a time.Time object, since json.Unmarshal()
// only supports RFC 3339 format.
func (i *Instance) UnmarshalJSON(data []byte) error {
	raw := struct {
		AntiAffinityGroups *[]AntiAffinityGroup `json:"anti-affinity-groups,omitempty"`
		CreatedAt          *string              `json:"created-at,omitempty"`
		DeployTarget       *DeployTarget        `json:"deploy-target,omitempty"`
		DiskSize           *int64               `json:"disk-size,omitempty"`
		ElasticIps         *[]ElasticIp         `json:"elastic-ips,omitempty"`
		Id                 *string              `json:"id,omitempty"` // nolint:revive
		InstanceType       *InstanceType        `json:"instance-type,omitempty"`
		Ipv6Address        *string              `json:"ipv6-address,omitempty"`
		Labels             *Labels              `json:"labels,omitempty"`
		Manager            *Manager             `json:"manager,omitempty"`
		Name               *string              `json:"name,omitempty"`
		PrivateNetworks    *[]PrivateNetwork    `json:"private-networks,omitempty"`
		PublicIp           *string              `json:"public-ip,omitempty"` // nolint:revive
		SecurityGroups     *[]SecurityGroup     `json:"security-groups,omitempty"`
		Snapshots          *[]Snapshot          `json:"snapshots,omitempty"`
		SshKey             *SshKey              `json:"ssh-key,omitempty"` // nolint:revive
		State              *InstanceState       `json:"state,omitempty"`
		Template           *Template            `json:"template,omitempty"`
		UserData           *string              `json:"user-data,omitempty"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.CreatedAt != nil {
		createdAt, err := time.Parse(iso8601Format, *raw.CreatedAt)
		if err != nil {
			return err
		}
		i.CreatedAt = &createdAt
	}

	i.AntiAffinityGroups = raw.AntiAffinityGroups
	i.DeployTarget = raw.DeployTarget
	i.DiskSize = raw.DiskSize
	i.ElasticIps = raw.ElasticIps
	i.Id = raw.Id
	i.InstanceType = raw.InstanceType
	i.Ipv6Address = raw.Ipv6Address
	i.Labels = raw.Labels
	i.Manager = raw.Manager
	i.Name = raw.Name
	i.PrivateNetworks = raw.PrivateNetworks
	i.PublicIp = raw.PublicIp
	i.SecurityGroups = raw.SecurityGroups
	i.Snapshots = raw.Snapshots
	i.SshKey = raw.SshKey
	i.State = raw.State
	i.Template = raw.Template
	i.UserData = raw.UserData

	return nil
}

// MarshalJSON returns the JSON encoding of an Instance structure after having formatted the CreatedAt field
// in the original timestamp (ISO 8601), since time.MarshalJSON() only supports RFC 3339 format.
func (i *Instance) MarshalJSON() ([]byte, error) {
	raw := struct {
		AntiAffinityGroups *[]AntiAffinityGroup `json:"anti-affinity-groups,omitempty"`
		CreatedAt          *string              `json:"created-at,omitempty"`
		DeployTarget       *DeployTarget        `json:"deploy-target,omitempty"`
		DiskSize           *int64               `json:"disk-size,omitempty"`
		ElasticIps         *[]ElasticIp         `json:"elastic-ips,omitempty"`
		Id                 *string              `json:"id,omitempty"` // nolint:revive
		InstanceType       *InstanceType        `json:"instance-type,omitempty"`
		Ipv6Address        *string              `json:"ipv6-address,omitempty"`
		Labels             *Labels              `json:"labels,omitempty"`
		Manager            *Manager             `json:"manager,omitempty"`
		Name               *string              `json:"name,omitempty"`
		PrivateNetworks    *[]PrivateNetwork    `json:"private-networks,omitempty"`
		PublicIp           *string              `json:"public-ip,omitempty"` // nolint:revive
		SecurityGroups     *[]SecurityGroup     `json:"security-groups,omitempty"`
		Snapshots          *[]Snapshot          `json:"snapshots,omitempty"`
		SshKey             *SshKey              `json:"ssh-key,omitempty"` // nolint:revive
		State              *InstanceState       `json:"state,omitempty"`
		Template           *Template            `json:"template,omitempty"`
		UserData           *string              `json:"user-data,omitempty"`
	}{}

	if i.CreatedAt != nil {
		createdAt := i.CreatedAt.Format(iso8601Format)
		raw.CreatedAt = &createdAt
	}

	raw.AntiAffinityGroups = i.AntiAffinityGroups
	raw.DeployTarget = i.DeployTarget
	raw.DiskSize = i.DiskSize
	raw.ElasticIps = i.ElasticIps
	raw.Id = i.Id
	raw.InstanceType = i.InstanceType
	raw.Ipv6Address = i.Ipv6Address
	raw.Labels = i.Labels
	raw.Manager = i.Manager
	raw.Name = i.Name
	raw.PrivateNetworks = i.PrivateNetworks
	raw.PublicIp = i.PublicIp
	raw.SecurityGroups = i.SecurityGroups
	raw.Snapshots = i.Snapshots
	raw.SshKey = i.SshKey
	raw.State = i.State
	raw.Template = i.Template
	raw.UserData = i.UserData

	return json.Marshal(raw)
}
