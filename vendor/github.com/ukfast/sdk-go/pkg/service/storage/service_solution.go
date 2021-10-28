package storage

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

func (s *Service) getSolutionResponseBody(solutionID int) (*GetSolutionResponseBody, error) {
	body := &GetSolutionResponseBody{}

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
