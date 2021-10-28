package draas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// DRaaSService is an interface for managing the DRaaS service
type DRaaSService interface {
	GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error)
	GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error)
	GetSolution(solutionID string) (Solution, error)
	PatchSolution(solutionID string, req PatchSolutionRequest) error

	GetSolutionBackupResources(solutionID string, parameters connection.APIRequestParameters) ([]BackupResource, error)
	GetSolutionBackupResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedBackupResource, error)

	GetSolutionBackupService(solutionID string) (BackupService, error)
	ResetSolutionBackupServiceCredentials(solutionID string, req ResetBackupServiceCredentialsRequest) error

	GetSolutionFailoverPlans(solutionID string, parameters connection.APIRequestParameters) ([]FailoverPlan, error)
	GetSolutionFailoverPlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedFailoverPlan, error)
	GetSolutionFailoverPlan(solutionID string, failoverPlanID string) (FailoverPlan, error)
	StartSolutionFailoverPlan(solutionID string, failoverPlanID string, req StartFailoverPlanRequest) error
	StopSolutionFailoverPlan(solutionID string, failoverPlanID string) error

	GetSolutionComputeResources(solutionID string, parameters connection.APIRequestParameters) ([]ComputeResource, error)
	GetSolutionComputeResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedComputeResource, error)
	GetSolutionComputeResource(solutionID string, computeResourcesID string) (ComputeResource, error)

	GetSolutionHardwarePlans(solutionID string, parameters connection.APIRequestParameters) ([]HardwarePlan, error)
	GetSolutionHardwarePlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedHardwarePlan, error)
	GetSolutionHardwarePlan(solutionID string, hardwarePlanID string) (HardwarePlan, error)
	GetSolutionHardwarePlanReplicas(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) ([]Replica, error)

	UpdateSolutionReplicaIOPS(solutionID string, replicaID string, req UpdateReplicaIOPSRequest) error

	GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error)
	GetIOPSTier(iopsTierID string) (IOPSTier, error)

	GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error)
	GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingType, error)
	GetBillingType(billingTypeID string) (BillingType, error)
}

// Service implements DRaaSService for managing
// DRaaS certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of DRaaSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
