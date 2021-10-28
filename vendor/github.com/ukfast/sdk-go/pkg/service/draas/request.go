package draas

import "github.com/ukfast/sdk-go/pkg/connection"

// PatchSolutionRequest represents a request to patch a solution
type PatchSolutionRequest struct {
	Name       string `json:"name,omitempty"`
	IOPSTierID string `json:"iops_tier_id,omitempty"`
}

// ResetBackupServiceCredentialsRequest represents a request to reset backup service credentials
type ResetBackupServiceCredentialsRequest struct {
	Password string `json:"password"`
}

// StartFailoverPlanRequest represents a request to start a failover plan
type StartFailoverPlanRequest struct {
	StartDate connection.DateTime `json:"start_date,omitempty"`
}

// UpdateReplicaIOPSRequest represents a request to update the IOPS for a replica
type UpdateReplicaIOPSRequest struct {
	IOPSTierID string `json:"iops_tier_id"`
}
