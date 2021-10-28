package draas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	var solutions []Solution

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, solution := range response.(*PaginatedSolution).Items {
			solutions = append(solutions, solution)
		}
	}

	return solutions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error) {
	body, err := s.getSolutionsPaginatedResponseBody(parameters)

	return NewPaginatedSolution(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSolutionSliceResponseBody, error) {
	body := &GetSolutionSliceResponseBody{}

	response, err := s.connection.Get("/draas/v1/solutions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSolution retrieves a single solution by id
func (s *Service) GetSolution(solutionID string) (Solution, error) {
	body, err := s.getSolutionResponseBody(solutionID)

	return body.Data, err
}

func (s *Service) getSolutionResponseBody(solutionID string) (*GetSolutionResponseBody, error) {
	body := &GetSolutionResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s", solutionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// PatchSolution patches a solution by ID
func (s *Service) PatchSolution(solutionID string, req PatchSolutionRequest) error {
	_, err := s.patchSolutionResponseBody(solutionID, req)

	return err
}

func (s *Service) patchSolutionResponseBody(solutionID string, req PatchSolutionRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/draas/v1/solutions/%s", solutionID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionBackupResources retrieves a collection of backup resources for specified solution
func (s *Service) GetSolutionBackupResources(solutionID string, parameters connection.APIRequestParameters) ([]BackupResource, error) {
	var resources []BackupResource

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionBackupResourcesPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, resource := range response.(*PaginatedBackupResource).Items {
			resources = append(resources, resource)
		}
	}

	return resources, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionBackupResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedBackupResource, error) {
	body, err := s.getSolutionBackupResourcesPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedBackupResource(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionBackupResourcesPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionBackupResourcesPaginatedResponseBody(solutionID string, parameters connection.APIRequestParameters) (*GetBackupResourceSliceResponseBody, error) {
	body := &GetBackupResourceSliceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/backup-resources", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionBackupService retrieves the backup service for the specified solution
func (s *Service) GetSolutionBackupService(solutionID string) (BackupService, error) {
	body, err := s.getSolutionBackupServiceResponseBody(solutionID)

	return body.Data, err
}

func (s *Service) getSolutionBackupServiceResponseBody(solutionID string) (*GetBackupServiceResponseBody, error) {
	body := &GetBackupServiceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/backup-service", solutionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// ResetSolutionBackupServiceCredentials resets the credentials for the solution backup service
func (s *Service) ResetSolutionBackupServiceCredentials(solutionID string, req ResetBackupServiceCredentialsRequest) error {
	_, err := s.resetSolutionBackupServiceCredentialsResponseBody(solutionID, req)

	return err
}

func (s *Service) resetSolutionBackupServiceCredentialsResponseBody(solutionID string, req ResetBackupServiceCredentialsRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/draas/v1/solutions/%s/backup-service/reset-credentials", solutionID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionFailoverPlans retrieves a collection of failover plans for specified solution
func (s *Service) GetSolutionFailoverPlans(solutionID string, parameters connection.APIRequestParameters) ([]FailoverPlan, error) {
	var resources []FailoverPlan

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionFailoverPlansPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, resource := range response.(*PaginatedFailoverPlan).Items {
			resources = append(resources, resource)
		}
	}

	return resources, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solution failover plans
func (s *Service) GetSolutionFailoverPlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedFailoverPlan, error) {
	body, err := s.getSolutionFailoverPlansPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedFailoverPlan(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionFailoverPlansPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionFailoverPlansPaginatedResponseBody(solutionID string, parameters connection.APIRequestParameters) (*GetFailoverPlanSliceResponseBody, error) {
	body := &GetFailoverPlanSliceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/failover-plans", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionFailoverPlan retrieves a single solution failover plan by id
func (s *Service) GetSolutionFailoverPlan(solutionID string, failoverPlanID string) (FailoverPlan, error) {
	body, err := s.getSolutionFailoverPlanResponseBody(solutionID, failoverPlanID)

	return body.Data, err
}

func (s *Service) getSolutionFailoverPlanResponseBody(solutionID string, failoverPlanID string) (*GetFailoverPlanResponseBody, error) {
	body := &GetFailoverPlanResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return body, fmt.Errorf("invalid failover plan id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s", solutionID, failoverPlanID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FailoverPlanNotFoundError{ID: failoverPlanID}
		}

		return nil
	})
}

// StartSolutionFailoverPlan starts the specified failover plan
func (s *Service) StartSolutionFailoverPlan(solutionID string, failoverPlanID string, req StartFailoverPlanRequest) error {
	_, err := s.startSolutionFailoverPlanResponseBody(solutionID, failoverPlanID, req)

	return err
}

func (s *Service) startSolutionFailoverPlanResponseBody(solutionID string, failoverPlanID string, req StartFailoverPlanRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return body, fmt.Errorf("invalid failover plan id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s/start", solutionID, failoverPlanID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FailoverPlanNotFoundError{ID: failoverPlanID}
		}

		return nil
	})
}

// StopSolutionFailoverPlan stops the specified failover plan
func (s *Service) StopSolutionFailoverPlan(solutionID string, failoverPlanID string) error {
	_, err := s.stopSolutionFailoverPlanResponseBody(solutionID, failoverPlanID)

	return err
}

func (s *Service) stopSolutionFailoverPlanResponseBody(solutionID string, failoverPlanID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if failoverPlanID == "" {
		return body, fmt.Errorf("invalid failover plan id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/draas/v1/solutions/%s/failover-plans/%s/stop", solutionID, failoverPlanID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FailoverPlanNotFoundError{ID: failoverPlanID}
		}

		return nil
	})
}

// GetSolutionComputeResources retrieves a collection of compute resources for specified solution
func (s *Service) GetSolutionComputeResources(solutionID string, parameters connection.APIRequestParameters) ([]ComputeResource, error) {
	var resources []ComputeResource

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionComputeResourcesPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, resource := range response.(*PaginatedComputeResource).Items {
			resources = append(resources, resource)
		}
	}

	return resources, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionComputeResourcesPaginated retrieves a paginated list of solution compute resources
func (s *Service) GetSolutionComputeResourcesPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedComputeResource, error) {
	body, err := s.getSolutionComputeResourcesPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedComputeResource(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionComputeResourcesPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionComputeResourcesPaginatedResponseBody(solutionID string, parameters connection.APIRequestParameters) (*GetComputeResourceSliceResponseBody, error) {
	body := &GetComputeResourceSliceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/compute-resources", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionComputeResource retrieves compute resources by id
func (s *Service) GetSolutionComputeResource(solutionID string, computeResourceID string) (ComputeResource, error) {
	body, err := s.getSolutionComputeResourceResponseBody(solutionID, computeResourceID)

	return body.Data, err
}

func (s *Service) getSolutionComputeResourceResponseBody(solutionID string, computeResourceID string) (*GetComputeResourceResponseBody, error) {
	body := &GetComputeResourceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if computeResourceID == "" {
		return body, fmt.Errorf("invalid compute resource id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/compute-resources/%s", solutionID, computeResourceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ComputeResourceNotFoundError{ID: computeResourceID}
		}

		return nil
	})
}

// GetSolutionHardwarePlans retrieves a collection of hardware plans for specified solution
func (s *Service) GetSolutionHardwarePlans(solutionID string, parameters connection.APIRequestParameters) ([]HardwarePlan, error) {
	var resources []HardwarePlan

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHardwarePlansPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, resource := range response.(*PaginatedHardwarePlan).Items {
			resources = append(resources, resource)
		}
	}

	return resources, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionHardwarePlansPaginated retrieves a paginated list of solution hardware plans
func (s *Service) GetSolutionHardwarePlansPaginated(solutionID string, parameters connection.APIRequestParameters) (*PaginatedHardwarePlan, error) {
	body, err := s.getSolutionHardwarePlansPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedHardwarePlan(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHardwarePlansPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionHardwarePlansPaginatedResponseBody(solutionID string, parameters connection.APIRequestParameters) (*GetHardwarePlanSliceResponseBody, error) {
	body := &GetHardwarePlanSliceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionHardwarePlan retrieves hardware plans by id
func (s *Service) GetSolutionHardwarePlan(solutionID string, hardwarePlanID string) (HardwarePlan, error) {
	body, err := s.getSolutionHardwarePlanResponseBody(solutionID, hardwarePlanID)

	return body.Data, err
}

func (s *Service) getSolutionHardwarePlanResponseBody(solutionID string, hardwarePlanID string) (*GetHardwarePlanResponseBody, error) {
	body := &GetHardwarePlanResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if hardwarePlanID == "" {
		return body, fmt.Errorf("invalid hardware plan id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans/%s", solutionID, hardwarePlanID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HardwarePlanNotFoundError{ID: hardwarePlanID}
		}

		return nil
	})
}

// GetSolutionHardwarePlanReplicas retrieves a collection of hardware plans for specified solution
func (s *Service) GetSolutionHardwarePlanReplicas(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) ([]Replica, error) {
	var resources []Replica

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHardwarePlanReplicasPaginated(solutionID, hardwarePlanID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, resource := range response.(*PaginatedReplica).Items {
			resources = append(resources, resource)
		}
	}

	return resources, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionHardwarePlanReplicasPaginated retrieves a paginated list of solution hardware plans
func (s *Service) GetSolutionHardwarePlanReplicasPaginated(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) (*PaginatedReplica, error) {
	body, err := s.getSolutionHardwarePlanReplicasPaginatedResponseBody(solutionID, hardwarePlanID, parameters)

	return NewPaginatedReplica(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHardwarePlanReplicasPaginated(solutionID, hardwarePlanID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionHardwarePlanReplicasPaginatedResponseBody(solutionID string, hardwarePlanID string, parameters connection.APIRequestParameters) (*GetReplicaSliceResponseBody, error) {
	body := &GetReplicaSliceResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if hardwarePlanID == "" {
		return body, fmt.Errorf("invalid hardware plan id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/solutions/%s/hardware-plans/%s/replicas", solutionID, hardwarePlanID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// UpdateSolutionReplicaIOPS updates a solution replica by ID
func (s *Service) UpdateSolutionReplicaIOPS(solutionID string, replicaID string, req UpdateReplicaIOPSRequest) error {
	_, err := s.updateSolutionReplicaResponseBody(solutionID, replicaID, req)

	return err
}

func (s *Service) updateSolutionReplicaResponseBody(solutionID string, replicaID string, req UpdateReplicaIOPSRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID == "" {
		return body, fmt.Errorf("invalid solution id")
	}
	if replicaID == "" {
		return body, fmt.Errorf("invalid replica id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/draas/v1/solutions/%s/replicas/%s/iops", solutionID, replicaID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ReplicaNotFoundError{ID: replicaID}
		}

		return nil
	})
}
