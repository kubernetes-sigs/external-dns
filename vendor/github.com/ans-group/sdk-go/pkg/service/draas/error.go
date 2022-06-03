package draas

import "fmt"

// SolutionNotFoundError indicates a solution was not found
type SolutionNotFoundError struct {
	ID string
}

func (e *SolutionNotFoundError) Error() string {
	return fmt.Sprintf("Solution not found with ID [%s]", e.ID)
}

// IOPSTierNotFoundError indicates an IOPS tier was not found
type IOPSTierNotFoundError struct {
	ID string
}

func (e *IOPSTierNotFoundError) Error() string {
	return fmt.Sprintf("IOPS tier not found with ID [%s]", e.ID)
}

// FailoverPlanNotFoundError indicates a failover plan was not found
type FailoverPlanNotFoundError struct {
	ID string
}

func (e *FailoverPlanNotFoundError) Error() string {
	return fmt.Sprintf("Failover plan not found with ID [%s]", e.ID)
}

// ComputeResourceNotFoundError indicates compute resources was not found
type ComputeResourceNotFoundError struct {
	ID string
}

func (e *ComputeResourceNotFoundError) Error() string {
	return fmt.Sprintf("Compute resources not found with ID [%s]", e.ID)
}

// HardwarePlanNotFoundError indicates hardware plan was not found
type HardwarePlanNotFoundError struct {
	ID string
}

func (e *HardwarePlanNotFoundError) Error() string {
	return fmt.Sprintf("Hardware plan not found with ID [%s]", e.ID)
}

// BillingTypeNotFoundError indicates billing type was not found
type BillingTypeNotFoundError struct {
	ID string
}

func (e *BillingTypeNotFoundError) Error() string {
	return fmt.Sprintf("Billing type not found with ID [%s]", e.ID)
}

// ReplicaNotFoundError indicates a replica was not found
type ReplicaNotFoundError struct {
	ID string
}

func (e *ReplicaNotFoundError) Error() string {
	return fmt.Sprintf("Replica not found with ID [%s]", e.ID)
}
