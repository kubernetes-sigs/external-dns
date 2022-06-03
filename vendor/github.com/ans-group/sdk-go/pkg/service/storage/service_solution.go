package storage

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return connection.InvokeRequestAll(s.GetSolutionsPaginated, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	body, err := s.getSolutionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSolutionsPaginated), err
}

func (s *Service) getSolutionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Solution], error) {
	body := &connection.APIResponseBodyData[[]Solution]{}

	response, err := s.connection.Get("/ukfast-storage/v1/solutions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSolution retrieves a single solution by id
func (s *Service) GetSolution(solutionID int) (Solution, error) {
	body, err := s.getSolutionResponseBody(solutionID)

	return body.Data, err
}

func (s *Service) getSolutionResponseBody(solutionID int) (*connection.APIResponseBodyData[Solution], error) {
	body := &connection.APIResponseBodyData[Solution]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ukfast-storage/v1/solutions/%d", solutionID), connection.APIRequestParameters{})
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
