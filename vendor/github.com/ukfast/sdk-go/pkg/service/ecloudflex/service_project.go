package ecloudflex

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetProjects retrieves a list of projects
func (s *Service) GetProjects(parameters connection.APIRequestParameters) ([]Project, error) {
	var projects []Project

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetProjectsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, project := range response.(*PaginatedProject).Items {
			projects = append(projects, project)
		}
	}

	return projects, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetProjectsPaginated retrieves a paginated list of projects
func (s *Service) GetProjectsPaginated(parameters connection.APIRequestParameters) (*PaginatedProject, error) {
	body, err := s.getProjectsPaginatedResponseBody(parameters)

	return NewPaginatedProject(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetProjectsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getProjectsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetProjectSliceResponseBody, error) {
	body := &GetProjectSliceResponseBody{}

	response, err := s.connection.Get("/ecloud-flex/v1/projects", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetProject retrieves a single project by id
func (s *Service) GetProject(projectID int) (Project, error) {
	body, err := s.getProjectResponseBody(projectID)

	return body.Data, err
}

func (s *Service) getProjectResponseBody(projectID int) (*GetProjectResponseBody, error) {
	body := &GetProjectResponseBody{}

	if projectID < 1 {
		return body, fmt.Errorf("invalid project id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud-flex/v1/projects/%d", projectID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ProjectNotFoundError{ID: projectID}
		}

		return nil
	})
}
