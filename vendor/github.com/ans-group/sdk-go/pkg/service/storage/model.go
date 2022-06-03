package storage

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// Solution represents a solution
type Solution struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	SanID     int                 `json:"san_id"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Host represents a host
type Host struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	OSType     string              `json:"os_type"`
	IQN        string              `json:"iqn"`
	ServerID   int                 `json:"server_id"`
	Status     string              `json:"status"`
	SolutionID int                 `json:"solution_id"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// HostSet represents a host set
type HostSet struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	SolutionID int                 `json:"solution_id"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// Volume represents a volume
type Volume struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	WWN        string              `json:"wwn"`
	SizeGB     int                 `json:"size_gb"`
	Status     string              `json:"status"`
	SolutionID int                 `json:"solution_id"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// VolumeSet represents a volume set
type VolumeSet struct {
	ID         int                 `json:"id"`
	SolutionID int                 `json:"solution_id"`
	MaxIOPS    int                 `json:"max_iops"`
	CreatedAt  connection.DateTime `json:"created_at"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
}

// IOPS represents an IOPS tier
type IOPS struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}
