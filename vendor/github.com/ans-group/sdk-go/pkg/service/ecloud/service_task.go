package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTasks retrieves a list of tasks
func (s *Service) GetTasks(parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(s.GetTasksPaginated, parameters)
}

// GetTasksPaginated retrieves a paginated list of tasks
func (s *Service) GetTasksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getTasksPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetTasksPaginated), err
}

func (s *Service) getTasksPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	response, err := s.connection.Get("/ecloud/v2/tasks", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTask retrieves a single task by id
func (s *Service) GetTask(taskID string) (Task, error) {
	body, err := s.getTaskResponseBody(taskID)

	return body.Data, err
}

func (s *Service) getTaskResponseBody(taskID string) (*connection.APIResponseBodyData[Task], error) {
	body := &connection.APIResponseBodyData[Task]{}

	if taskID == "" {
		return body, fmt.Errorf("invalid task id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/tasks/%s", taskID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TaskNotFoundError{ID: taskID}
		}

		return nil
	})
}
