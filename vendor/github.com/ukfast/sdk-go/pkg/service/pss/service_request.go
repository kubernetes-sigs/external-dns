package pss

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// CreateRequest creates a new request
func (s *Service) CreateRequest(req CreateRequestRequest) (int, error) {
	body, err := s.createRequestResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createRequestResponseBody(req CreateRequestRequest) (*GetRequestResponseBody, error) {
	body := &GetRequestResponseBody{}

	response, err := s.connection.Post("/pss/v1/requests", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRequests retrieves a list of requests
func (s *Service) GetRequests(parameters connection.APIRequestParameters) ([]Request, error) {
	var requests []Request

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRequestsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, request := range response.(*PaginatedRequest).Items {
			requests = append(requests, request)
		}
	}

	return requests, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRequestsPaginated retrieves a paginated list of requests
func (s *Service) GetRequestsPaginated(parameters connection.APIRequestParameters) (*PaginatedRequest, error) {
	body, err := s.getRequestsPaginatedResponseBody(parameters)

	return NewPaginatedRequest(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRequestsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRequestsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRequestSliceResponseBody, error) {
	body := &GetRequestSliceResponseBody{}

	response, err := s.connection.Get("/pss/v1/requests", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRequest retrieves a single request by id
func (s *Service) GetRequest(requestID int) (Request, error) {
	body, err := s.getRequestResponseBody(requestID)

	return body.Data, err
}

func (s *Service) getRequestResponseBody(requestID int) (*GetRequestResponseBody, error) {
	body := &GetRequestResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/requests/%d", requestID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestNotFoundError{ID: requestID}
		}

		return nil
	})
}

// PatchRequest patches a request
func (s *Service) PatchRequest(requestID int, req PatchRequestRequest) error {
	_, err := s.patchRequestResponseBody(requestID, req)

	return err
}

func (s *Service) patchRequestResponseBody(requestID int, req PatchRequestRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/pss/v1/requests/%d", requestID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestNotFoundError{ID: requestID}
		}

		return nil
	})
}

// CreateRequestReply creates a new request reply
func (s *Service) CreateRequestReply(requestID int, req CreateReplyRequest) (string, error) {
	body, err := s.createRequestReplyResponseBody(requestID, req)

	return body.Data.ID, err
}

func (s *Service) createRequestReplyResponseBody(requestID int, req CreateReplyRequest) (*GetReplyResponseBody, error) {
	body := &GetReplyResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/pss/v1/requests/%d/replies", requestID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestNotFoundError{ID: requestID}
		}

		return nil
	})
}

// GetRequestReplies is an alias for GetRequestConversation
func (s *Service) GetRequestReplies(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error) {
	return s.GetRequestConversation(solutionID, parameters)
}

// GetRequestRepliesPaginated is an alias for GetRequestConversationPaginated
func (s *Service) GetRequestRepliesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedReply, error) {
	return s.GetRequestConversationPaginated(solutionID, parameters)
}

// GetRequestConversation retrieves a list of replies
func (s *Service) GetRequestConversation(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error) {
	var replies []Reply

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRequestConversationPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, reply := range response.(*PaginatedReply).Items {
			replies = append(replies, reply)
		}
	}

	return replies, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRequestConversationPaginated retrieves a paginated list of domains
func (s *Service) GetRequestConversationPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedReply, error) {
	body, err := s.getRequestConversationPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedReply(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRequestConversationPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRequestConversationPaginatedResponseBody(requestID int, parameters connection.APIRequestParameters) (*GetReplySliceResponseBody, error) {
	body := &GetReplySliceResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/requests/%d/conversation", requestID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestNotFoundError{ID: requestID}
		}

		return nil
	})
}

// GetRequestFeedback retrieves feedback for a request
func (s *Service) GetRequestFeedback(requestID int) (Feedback, error) {
	body, err := s.getRequestFeedbackResponseBody(requestID)

	return body.Data, err
}

func (s *Service) getRequestFeedbackResponseBody(requestID int) (*GetFeedbackResponseBody, error) {
	body := &GetFeedbackResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/pss/v1/requests/%d/feedback", requestID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestFeedbackNotFoundError{RequestID: requestID}
		}

		return nil
	})
}

// CreateRequestFeedback creates a new request feedback
func (s *Service) CreateRequestFeedback(requestID int, req CreateFeedbackRequest) (int, error) {
	body, err := s.createRequestFeedbackResponseBody(requestID, req)

	return body.Data.ID, err
}

func (s *Service) createRequestFeedbackResponseBody(requestID int, req CreateFeedbackRequest) (*GetFeedbackResponseBody, error) {
	body := &GetFeedbackResponseBody{}

	if requestID < 1 {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/pss/v1/requests/%d/feedback", requestID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RequestNotFoundError{ID: requestID}
		}

		return nil
	})
}
