package oapi

import (
	"encoding/json"
	"time"
)

// UnmarshalJSON unmarshals a DbaasService structure into a temporary structure whose "CreatedAt" field of type
// string to be able to parse the original timestamp (ISO 8601) into a time.Time object, since json.Unmarshal()
// only supports RFC 3339 format.
func (s *DbaasService) UnmarshalJSON(data []byte) error {
	raw := struct {
		Acl                   *[]DbaasServiceAcl             `json:"acl,omitempty"` // nolint:revive
		Backups               *[]DbaasServiceBackup          `json:"backups,omitempty"`
		Components            *[]DbaasServiceComponents      `json:"components,omitempty"`
		ConnectionInfo        *DbaasService_ConnectionInfo   `json:"connection-info,omitempty"`
		ConnectionPools       *[]DbaasServiceConnectionPools `json:"connection-pools,omitempty"`
		CreatedAt             *string                        `json:"created-at,omitempty"`
		DiskSize              *int64                         `json:"disk-size,omitempty"`
		Features              *DbaasService_Features         `json:"features,omitempty"`
		Integrations          *[]DbaasServiceIntegration     `json:"integrations,omitempty"`
		Maintenance           *DbaasServiceMaintenance       `json:"maintenance,omitempty"`
		Metadata              *DbaasService_Metadata         `json:"metadata,omitempty"`
		Name                  DbaasServiceName               `json:"name"`
		NodeCount             *int64                         `json:"node-count,omitempty"`
		NodeCpuCount          *int64                         `json:"node-cpu-count,omitempty"` // nolint:revive
		NodeMemory            *int64                         `json:"node-memory,omitempty"`
		NodeStates            *[]DbaasNodeState              `json:"node-states,omitempty"`
		Notifications         *[]DbaasServiceNotification    `json:"notifications,omitempty"`
		Plan                  string                         `json:"plan"`
		State                 *DbaasServiceState             `json:"state,omitempty"`
		TerminationProtection *bool                          `json:"termination-protection,omitempty"`
		Type                  DbaasServiceTypeName           `json:"type"`
		UpdatedAt             *string                        `json:"updated-at,omitempty"`
		Uri                   *string                        `json:"uri,omitempty"`        // nolint:revive
		UriParams             *DbaasService_UriParams        `json:"uri-params,omitempty"` // nolint:revive
		UserConfig            *DbaasService_UserConfig       `json:"user-config,omitempty"`
		Users                 *[]DbaasServiceUser            `json:"users,omitempty"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.CreatedAt != nil {
		createdAt, err := time.Parse(iso8601Format, *raw.CreatedAt)
		if err != nil {
			return err
		}
		s.CreatedAt = &createdAt
	}

	if raw.UpdatedAt != nil {
		updatedAt, err := time.Parse(iso8601Format, *raw.UpdatedAt)
		if err != nil {
			return err
		}
		s.UpdatedAt = &updatedAt
	}

	s.Acl = raw.Acl
	s.Backups = raw.Backups
	s.Components = raw.Components
	s.ConnectionInfo = raw.ConnectionInfo
	s.ConnectionPools = raw.ConnectionPools
	s.DiskSize = raw.DiskSize
	s.Features = raw.Features
	s.Integrations = raw.Integrations
	s.Maintenance = raw.Maintenance
	s.Metadata = raw.Metadata
	s.Name = raw.Name
	s.NodeCount = raw.NodeCount
	s.NodeCpuCount = raw.NodeCpuCount
	s.NodeMemory = raw.NodeMemory
	s.NodeStates = raw.NodeStates
	s.Notifications = raw.Notifications
	s.Plan = raw.Plan
	s.State = raw.State
	s.TerminationProtection = raw.TerminationProtection
	s.Type = raw.Type
	s.Uri = raw.Uri
	s.UriParams = raw.UriParams
	s.UserConfig = raw.UserConfig
	s.Users = raw.Users

	return nil
}

// MarshalJSON returns the JSON encoding of a DbaasService structure after having formatted the CreatedAt field
// in the original timestamp (ISO 8601), since time.MarshalJSON() only supports RFC 3339 format.
func (s *DbaasService) MarshalJSON() ([]byte, error) {
	raw := struct {
		Acl                   *[]DbaasServiceAcl             `json:"acl,omitempty"` // nolint:revive
		Backups               *[]DbaasServiceBackup          `json:"backups,omitempty"`
		Components            *[]DbaasServiceComponents      `json:"components,omitempty"`
		ConnectionInfo        *DbaasService_ConnectionInfo   `json:"connection-info,omitempty"`
		ConnectionPools       *[]DbaasServiceConnectionPools `json:"connection-pools,omitempty"`
		CreatedAt             *string                        `json:"created-at,omitempty"`
		DiskSize              *int64                         `json:"disk-size,omitempty"`
		Features              *DbaasService_Features         `json:"features,omitempty"`
		Integrations          *[]DbaasServiceIntegration     `json:"integrations,omitempty"`
		Maintenance           *DbaasServiceMaintenance       `json:"maintenance,omitempty"`
		Metadata              *DbaasService_Metadata         `json:"metadata,omitempty"`
		Name                  DbaasServiceName               `json:"name"`
		NodeCount             *int64                         `json:"node-count,omitempty"`
		NodeCpuCount          *int64                         `json:"node-cpu-count,omitempty"` // nolint:revive
		NodeMemory            *int64                         `json:"node-memory,omitempty"`
		NodeStates            *[]DbaasNodeState              `json:"node-states,omitempty"`
		Notifications         *[]DbaasServiceNotification    `json:"notifications,omitempty"`
		Plan                  string                         `json:"plan"`
		State                 *DbaasServiceState             `json:"state,omitempty"`
		TerminationProtection *bool                          `json:"termination-protection,omitempty"`
		Type                  DbaasServiceTypeName           `json:"type"`
		UpdatedAt             *string                        `json:"updated-at,omitempty"`
		Uri                   *string                        `json:"uri,omitempty"`        // nolint:revive
		UriParams             *DbaasService_UriParams        `json:"uri-params,omitempty"` // nolint:revive
		UserConfig            *DbaasService_UserConfig       `json:"user-config,omitempty"`
		Users                 *[]DbaasServiceUser            `json:"users,omitempty"`
	}{}

	if s.CreatedAt != nil {
		createdAt := s.CreatedAt.Format(iso8601Format)
		raw.CreatedAt = &createdAt
	}

	if s.UpdatedAt != nil {
		updatedAt := s.UpdatedAt.Format(iso8601Format)
		raw.UpdatedAt = &updatedAt
	}

	raw.Acl = s.Acl
	raw.Backups = s.Backups
	raw.Components = s.Components
	raw.ConnectionInfo = s.ConnectionInfo
	raw.ConnectionPools = s.ConnectionPools
	raw.DiskSize = s.DiskSize
	raw.Features = s.Features
	raw.Integrations = s.Integrations
	raw.Maintenance = s.Maintenance
	raw.Metadata = s.Metadata
	raw.Name = s.Name
	raw.NodeCount = s.NodeCount
	raw.NodeCpuCount = s.NodeCpuCount
	raw.NodeMemory = s.NodeMemory
	raw.NodeStates = s.NodeStates
	raw.Notifications = s.Notifications
	raw.Plan = s.Plan
	raw.State = s.State
	raw.TerminationProtection = s.TerminationProtection
	raw.Type = s.Type
	raw.Uri = s.Uri
	raw.UriParams = s.UriParams
	raw.UserConfig = s.UserConfig
	raw.Users = s.Users

	return json.Marshal(raw)
}
