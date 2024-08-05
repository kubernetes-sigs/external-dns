package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDeployments retrieves a list of deployments
func (s *Service) GetDeployments(parameters connection.APIRequestParameters) ([]Deployment, error) {
	return connection.InvokeRequestAll(s.GetDeploymentsPaginated, parameters)
}

// GetDeploymentsPaginated retrieves a paginated list of deployments
func (s *Service) GetDeploymentsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Deployment], error) {
	body, err := s.getDeploymentsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDeploymentsPaginated), err
}

func (s *Service) getDeploymentsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Deployment], error) {
	return connection.Get[[]Deployment](s.connection, "/loadbalancers/v2/deployments", parameters)
}

// GetDeployment retrieves a single deployment by id
func (s *Service) GetDeployment(deploymentID int) (Deployment, error) {
	body, err := s.getDeploymentResponseBody(deploymentID)

	return body.Data, err
}

func (s *Service) getDeploymentResponseBody(deploymentID int) (*connection.APIResponseBodyData[Deployment], error) {
	if deploymentID < 1 {
		return &connection.APIResponseBodyData[Deployment]{}, fmt.Errorf("invalid deployment id")
	}

	return connection.Get[Deployment](s.connection, fmt.Sprintf("/loadbalancers/v2/deployments/%d", deploymentID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DeploymentNotFoundError{ID: deploymentID}))
}
