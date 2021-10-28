package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTasks retrieves a list of tasks
func (s *Service) GetTasks(parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTasksPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTasksPaginated retrieves a paginated list of tasks
func (s *Service) GetTasksPaginated(parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getTasksPaginatedResponseBody(parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTasksPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTasksPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

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

func (s *Service) getTaskResponseBody(taskID string) (*GetTaskResponseBody, error) {
	body := &GetTaskResponseBody{}

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
