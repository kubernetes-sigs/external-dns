package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals a SksNodepool structure into a temporary structure whose "CreatedAt" field of type
// string to be able to parse the original timestamp (ISO 8601) into a time.Time object, since json.Unmarshal()
// only supports RFC 3339 format.
func (n *SksNodepool) UnmarshalJSON(data []byte) error {
	raw := struct {
		Addons             *[]SksNodepoolAddons `json:"addons,omitempty"`
		AntiAffinityGroups *[]AntiAffinityGroup `json:"anti-affinity-groups,omitempty"`
		CreatedAt          *string              `json:"created-at,omitempty"`
		DeployTarget       *DeployTarget        `json:"deploy-target,omitempty"`
		Description        *string              `json:"description,omitempty"`
		DiskSize           *int64               `json:"disk-size,omitempty"`
		Id                 *string              `json:"id,omitempty"` // nolint:revive
		InstancePool       *InstancePool        `json:"instance-pool,omitempty"`
		InstancePrefix     *string              `json:"instance-prefix,omitempty"`
		InstanceType       *InstanceType        `json:"instance-type,omitempty"`
		Labels             *Labels              `json:"labels,omitempty"`
		Name               *string              `json:"name,omitempty"`
		PrivateNetworks    *[]PrivateNetwork    `json:"private-networks,omitempty"`
		SecurityGroups     *[]SecurityGroup     `json:"security-groups,omitempty"`
		Size               *int64               `json:"size,omitempty"`
		State              *SksNodepoolState    `json:"state,omitempty"`
		Template           *Template            `json:"template,omitempty"`
		Version            *string              `json:"version,omitempty"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.CreatedAt != nil {
		createdAt, err := time.Parse(iso8601Format, *raw.CreatedAt)
		if err != nil {
			return err
		}
		n.CreatedAt = &createdAt
	}

	n.Addons = raw.Addons
	n.AntiAffinityGroups = raw.AntiAffinityGroups
	n.DeployTarget = raw.DeployTarget
	n.Description = raw.Description
	n.DiskSize = raw.DiskSize
	n.Id = raw.Id
	n.InstancePool = raw.InstancePool
	n.InstancePrefix = raw.InstancePrefix
	n.InstanceType = raw.InstanceType
	n.Labels = raw.Labels
	n.Name = raw.Name
	n.PrivateNetworks = raw.PrivateNetworks
	n.SecurityGroups = raw.SecurityGroups
	n.Size = raw.Size
	n.State = raw.State
	n.Template = raw.Template
	n.Version = raw.Version

	return nil
}

// MarshalJSON returns the JSON encoding of a SksNodepool structure after having formatted the CreatedAt field
// in the original timestamp (ISO 8601), since time.MarshalJSON() only supports RFC 3339 format.
func (n *SksNodepool) MarshalJSON() ([]byte, error) {
	raw := struct {
		Addons             *[]SksNodepoolAddons `json:"addons,omitempty"`
		AntiAffinityGroups *[]AntiAffinityGroup `json:"anti-affinity-groups,omitempty"`
		CreatedAt          *string              `json:"created-at,omitempty"`
		DeployTarget       *DeployTarget        `json:"deploy-target,omitempty"`
		Description        *string              `json:"description,omitempty"`
		DiskSize           *int64               `json:"disk-size,omitempty"`
		Id                 *string              `json:"id,omitempty"` // nolint:revive
		InstancePool       *InstancePool        `json:"instance-pool,omitempty"`
		InstancePrefix     *string              `json:"instance-prefix,omitempty"`
		InstanceType       *InstanceType        `json:"instance-type,omitempty"`
		Labels             *Labels              `json:"labels,omitempty"`
		Name               *string              `json:"name,omitempty"`
		PrivateNetworks    *[]PrivateNetwork    `json:"private-networks,omitempty"`
		SecurityGroups     *[]SecurityGroup     `json:"security-groups,omitempty"`
		Size               *int64               `json:"size,omitempty"`
		State              *SksNodepoolState    `json:"state,omitempty"`
		Template           *Template            `json:"template,omitempty"`
		Version            *string              `json:"version,omitempty"`
	}{}

	if n.CreatedAt != nil {
		createdAt := n.CreatedAt.Format(iso8601Format)
		raw.CreatedAt = &createdAt
	}

	raw.Addons = n.Addons
	raw.AntiAffinityGroups = n.AntiAffinityGroups
	raw.DeployTarget = n.DeployTarget
	raw.Description = n.Description
	raw.DiskSize = n.DiskSize
	raw.Id = n.Id
	raw.InstancePool = n.InstancePool
	raw.InstancePrefix = n.InstancePrefix
	raw.InstanceType = n.InstanceType
	raw.Labels = n.Labels
	raw.Name = n.Name
	raw.PrivateNetworks = n.PrivateNetworks
	raw.SecurityGroups = n.SecurityGroups
	raw.Size = n.Size
	raw.State = n.State
	raw.Template = n.Template
	raw.Version = n.Version

	return json.Marshal(raw)
}
