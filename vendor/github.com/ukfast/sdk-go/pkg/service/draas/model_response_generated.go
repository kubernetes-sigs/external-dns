package draas

import "github.com/ukfast/sdk-go/pkg/connection"

// GetSolutionSliceResponseBody represents an API response body containing []Solution data
type GetSolutionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents an API response body containing Solution data
type GetSolutionResponseBody struct {
	connection.APIResponseBody
	Data Solution `json:"data"`
}

// GetBackupResourceSliceResponseBody represents an API response body containing []BackupResource data
type GetBackupResourceSliceResponseBody struct {
	connection.APIResponseBody
	Data []BackupResource `json:"data"`
}

// GetBackupResourceResponseBody represents an API response body containing BackupResource data
type GetBackupResourceResponseBody struct {
	connection.APIResponseBody
	Data BackupResource `json:"data"`
}

// GetIOPSTierSliceResponseBody represents an API response body containing []IOPSTier data
type GetIOPSTierSliceResponseBody struct {
	connection.APIResponseBody
	Data []IOPSTier `json:"data"`
}

// GetIOPSTierResponseBody represents an API response body containing IOPSTier data
type GetIOPSTierResponseBody struct {
	connection.APIResponseBody
	Data IOPSTier `json:"data"`
}

// GetBackupServiceSliceResponseBody represents an API response body containing []BackupService data
type GetBackupServiceSliceResponseBody struct {
	connection.APIResponseBody
	Data []BackupService `json:"data"`
}

// GetBackupServiceResponseBody represents an API response body containing BackupService data
type GetBackupServiceResponseBody struct {
	connection.APIResponseBody
	Data BackupService `json:"data"`
}

// GetFailoverPlanSliceResponseBody represents an API response body containing []FailoverPlan data
type GetFailoverPlanSliceResponseBody struct {
	connection.APIResponseBody
	Data []FailoverPlan `json:"data"`
}

// GetFailoverPlanResponseBody represents an API response body containing FailoverPlan data
type GetFailoverPlanResponseBody struct {
	connection.APIResponseBody
	Data FailoverPlan `json:"data"`
}

// GetComputeResourceSliceResponseBody represents an API response body containing []ComputeResource data
type GetComputeResourceSliceResponseBody struct {
	connection.APIResponseBody
	Data []ComputeResource `json:"data"`
}

// GetComputeResourceResponseBody represents an API response body containing ComputeResource data
type GetComputeResourceResponseBody struct {
	connection.APIResponseBody
	Data ComputeResource `json:"data"`
}

// GetHardwarePlanSliceResponseBody represents an API response body containing []HardwarePlan data
type GetHardwarePlanSliceResponseBody struct {
	connection.APIResponseBody
	Data []HardwarePlan `json:"data"`
}

// GetHardwarePlanResponseBody represents an API response body containing HardwarePlan data
type GetHardwarePlanResponseBody struct {
	connection.APIResponseBody
	Data HardwarePlan `json:"data"`
}

// GetReplicaSliceResponseBody represents an API response body containing []Replica data
type GetReplicaSliceResponseBody struct {
	connection.APIResponseBody
	Data []Replica `json:"data"`
}

// GetReplicaResponseBody represents an API response body containing Replica data
type GetReplicaResponseBody struct {
	connection.APIResponseBody
	Data Replica `json:"data"`
}

// GetBillingTypeSliceResponseBody represents an API response body containing []BillingType data
type GetBillingTypeSliceResponseBody struct {
	connection.APIResponseBody
	Data []BillingType `json:"data"`
}

// GetBillingTypeResponseBody represents an API response body containing BillingType data
type GetBillingTypeResponseBody struct {
	connection.APIResponseBody
	Data BillingType `json:"data"`
}
