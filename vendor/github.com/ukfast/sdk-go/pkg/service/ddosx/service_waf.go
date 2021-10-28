package ddosx

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetWAFLogs retrieves a list of logs
func (s *Service) GetWAFLogs(parameters connection.APIRequestParameters) ([]WAFLog, error) {
	var matches []WAFLog

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, match := range response.(*PaginatedWAFLog).Items {
			matches = append(matches, match)
		}
	}

	return matches, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetWAFLogsPaginated retrieves a paginated list of logs
func (s *Service) GetWAFLogsPaginated(parameters connection.APIRequestParameters) (*PaginatedWAFLog, error) {
	body, err := s.getWAFLogsPaginatedResponseBody(parameters)

	return NewPaginatedWAFLog(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getWAFLogsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetWAFLogSliceResponseBody, error) {
	body := &GetWAFLogSliceResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/waf/logs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetWAFLog retrieves a single log by id
func (s *Service) GetWAFLog(requestID string) (WAFLog, error) {
	body, err := s.getWAFLogResponseBody(requestID)

	return body.Data, err
}

func (s *Service) getWAFLogResponseBody(requestID string) (*GetWAFLogResponseBody, error) {
	body := &GetWAFLogResponseBody{}

	if requestID == "" {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/waf/logs/%s", requestID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFLogNotFoundError{ID: requestID}
		}

		return nil
	})
}

// GetWAFLogMatches retrieves a list of log matches
func (s *Service) GetWAFLogMatches(parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	var matches []WAFLogMatch

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogMatchesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, match := range response.(*PaginatedWAFLogMatch).Items {
			matches = append(matches, match)
		}
	}

	return matches, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetWAFLogMatchesPaginated retrieves a paginated list of log matches
func (s *Service) GetWAFLogMatchesPaginated(parameters connection.APIRequestParameters) (*PaginatedWAFLogMatch, error) {
	body, err := s.getWAFLogMatchesPaginatedResponseBody(parameters)

	return NewPaginatedWAFLogMatch(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogMatchesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getWAFLogMatchesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetWAFLogMatchSliceResponseBody, error) {
	body := &GetWAFLogMatchSliceResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/waf/logs/matches", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetWAFLogRequestMatches retrieves a list of log matches for request
func (s *Service) GetWAFLogRequestMatches(requestID string, parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	var matches []WAFLogMatch

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogRequestMatchesPaginated(requestID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, match := range response.(*PaginatedWAFLogMatch).Items {
			matches = append(matches, match)
		}
	}

	return matches, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetWAFLogRequestMatchesPaginated retrieves a paginated list of matches for request
func (s *Service) GetWAFLogRequestMatchesPaginated(requestID string, parameters connection.APIRequestParameters) (*PaginatedWAFLogMatch, error) {
	body, err := s.getWAFLogRequestMatchesPaginatedResponseBody(requestID, parameters)

	return NewPaginatedWAFLogMatch(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetWAFLogMatchesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getWAFLogRequestMatchesPaginatedResponseBody(requestID string, parameters connection.APIRequestParameters) (*GetWAFLogMatchSliceResponseBody, error) {
	body := &GetWAFLogMatchSliceResponseBody{}

	if requestID == "" {
		return body, fmt.Errorf("invalid request id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/waf/logs/%s/matches", requestID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFLogNotFoundError{ID: requestID}
		}

		return nil
	})
}

// GetWAFLogRequestMatch retrieves a single waf log request match
func (s *Service) GetWAFLogRequestMatch(requestID string, matchID string) (WAFLogMatch, error) {
	body, err := s.getWAFLogRequestMatchResponseBody(requestID, matchID)

	return body.Data, err
}

func (s *Service) getWAFLogRequestMatchResponseBody(requestID string, matchID string) (*GetWAFLogMatchResponseBody, error) {
	body := &GetWAFLogMatchResponseBody{}

	if requestID == "" {
		return body, fmt.Errorf("invalid request id")
	}

	if matchID == "" {
		return body, fmt.Errorf("invalid match id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/waf/logs/%s/matches/%s", requestID, matchID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFLogMatchNotFoundError{ID: requestID}
		}

		return nil
	})
}
