//go:generate go run ../../gen/model_response/main.go -package draas -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package draas -source model.go -destination model_paginated_generated.go

package draas

// Solution represents a solution
// +genie:model_response
// +genie:model_paginated
type Solution struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IOPSTierID    string `json:"iops_tier_id"`
	BillingTypeID string `json:"billing_type_id"`
}

// BackupResource represents backup resources for a solution
// +genie:model_response
// +genie:model_paginated
type BackupResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Quota in DB
	Quota int `json:"quota"`
	// Used quota in DB
	UsedQuota float32 `json:"used_quota"`
}

// IOPSTier represents an IOPS tier
// +genie:model_response
// +genie:model_paginated
type IOPSTier struct {
	ID        string `json:"id"`
	IOPSLimit int    `json:"iops_limit"`
}

// BackupService represents the backup service for a solution
// +genie:model_response
type BackupService struct {
	Service     string `json:"service"`
	AccountName string `json:"account_name"`
	Gateway     string `json:"gateway"`
	Port        int    `json:"port"`
}

// FailoverPlan represents a failover plan
// +genie:model_response
// +genie:model_paginated
type FailoverPlan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	VMs         []struct {
		Name string `json:"name"`
	} `json:"vms"`
}

// ComputeResource represents compute resources for a solution
// +genie:model_response
// +genie:model_paginated
type ComputeResource struct {
	ID             string `json:"id"`
	HardwarePlanID string `json:"hardware_plan_id"`
	Memory         struct {
		// Used memory in GB
		Used float32 `json:"used"`
		// Memory limit in GB
		Limit float32 `json:"limit"`
	} `json:"memory"`
	CPU struct {
		Used int `json:"used"`
	} `json:"cpu"`
	Storage []struct {
		Name string `json:"name"`
		// Used storage in GB
		Used int `json:"used"`
		// Storage limit in GB
		Limit int `json:"limit"`
	} `json:"storage"`
}

// HardwarePlan represents a hardware plan
// +genie:model_response
// +genie:model_paginated
type HardwarePlan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Limits      struct {
		Processor int `json:"processor"`
		Memory    int `json:"memory"`
	} `json:"limits"`
	Networks struct {
		Public  int `json:"public"`
		Private int `json:"private"`
	} `json:"networks"`
	Storage []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"Type"`
		Quota int    `json:"quota"`
	} `json:"storage"`
}

// Replica represents a replica
// +genie:model_response
// +genie:model_paginated
type Replica struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	CPU      int    `json:"cpu"`
	RAM      int    `json:"ram"`
	Disk     int    `json:"disk"`
	IOPS     int    `json:"iops"`
	Power    bool   `json:"power"`
}

// BillingType represents a billing type
// +genie:model_response
// +genie:model_paginated
type BillingType struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
