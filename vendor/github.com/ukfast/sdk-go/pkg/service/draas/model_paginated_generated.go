package draas

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedSolution represents a paginated collection of Solution
type PaginatedSolution struct {
	*connection.PaginatedBase
	Items []Solution
}

// NewPaginatedSolution returns a pointer to an initialized PaginatedSolution struct
func NewPaginatedSolution(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Solution) *PaginatedSolution {
	return &PaginatedSolution{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedBackupResource represents a paginated collection of BackupResource
type PaginatedBackupResource struct {
	*connection.PaginatedBase
	Items []BackupResource
}

// NewPaginatedBackupResource returns a pointer to an initialized PaginatedBackupResource struct
func NewPaginatedBackupResource(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []BackupResource) *PaginatedBackupResource {
	return &PaginatedBackupResource{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedIOPSTier represents a paginated collection of IOPSTier
type PaginatedIOPSTier struct {
	*connection.PaginatedBase
	Items []IOPSTier
}

// NewPaginatedIOPSTier returns a pointer to an initialized PaginatedIOPSTier struct
func NewPaginatedIOPSTier(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []IOPSTier) *PaginatedIOPSTier {
	return &PaginatedIOPSTier{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFailoverPlan represents a paginated collection of FailoverPlan
type PaginatedFailoverPlan struct {
	*connection.PaginatedBase
	Items []FailoverPlan
}

// NewPaginatedFailoverPlan returns a pointer to an initialized PaginatedFailoverPlan struct
func NewPaginatedFailoverPlan(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []FailoverPlan) *PaginatedFailoverPlan {
	return &PaginatedFailoverPlan{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedComputeResource represents a paginated collection of ComputeResource
type PaginatedComputeResource struct {
	*connection.PaginatedBase
	Items []ComputeResource
}

// NewPaginatedComputeResource returns a pointer to an initialized PaginatedComputeResource struct
func NewPaginatedComputeResource(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ComputeResource) *PaginatedComputeResource {
	return &PaginatedComputeResource{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHardwarePlan represents a paginated collection of HardwarePlan
type PaginatedHardwarePlan struct {
	*connection.PaginatedBase
	Items []HardwarePlan
}

// NewPaginatedHardwarePlan returns a pointer to an initialized PaginatedHardwarePlan struct
func NewPaginatedHardwarePlan(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []HardwarePlan) *PaginatedHardwarePlan {
	return &PaginatedHardwarePlan{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedReplica represents a paginated collection of Replica
type PaginatedReplica struct {
	*connection.PaginatedBase
	Items []Replica
}

// NewPaginatedReplica returns a pointer to an initialized PaginatedReplica struct
func NewPaginatedReplica(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Replica) *PaginatedReplica {
	return &PaginatedReplica{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedBillingType represents a paginated collection of BillingType
type PaginatedBillingType struct {
	*connection.PaginatedBase
	Items []BillingType
}

// NewPaginatedBillingType returns a pointer to an initialized PaginatedBillingType struct
func NewPaginatedBillingType(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []BillingType) *PaginatedBillingType {
	return &PaginatedBillingType{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
