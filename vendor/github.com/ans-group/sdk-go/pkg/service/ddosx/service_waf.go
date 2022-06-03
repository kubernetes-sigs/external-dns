package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetWAFLogs retrieves a list of logs
func (s *Service) GetWAFLogs(parameters connection.APIRequestParameters) ([]WAFLog, error) {
	return connection.InvokeRequestAll(s.GetWAFLogsPaginated, parameters)
}

// GetWAFLogsPaginated retrieves a paginated list of logs
func (s *Service) GetWAFLogsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLog], error) {
	body, err := s.getWAFLogsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetWAFLogsPaginated), err
}

func (s *Service) getWAFLogsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFLog], error) {
	body := &connection.APIResponseBodyData[[]WAFLog]{}

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

func (s *Service) getWAFLogResponseBody(requestID string) (*connection.APIResponseBodyData[WAFLog], error) {
	body := &connection.APIResponseBodyData[WAFLog]{}

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
	return connection.InvokeRequestAll(s.GetWAFLogMatchesPaginated, parameters)
}

// GetWAFLogMatchesPaginated retrieves a paginated list of log matches
func (s *Service) GetWAFLogMatchesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
	body, err := s.getWAFLogMatchesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetWAFLogMatchesPaginated), err
}

func (s *Service) getWAFLogMatchesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFLogMatch], error) {
	body := &connection.APIResponseBodyData[[]WAFLogMatch]{}

	response, err := s.connection.Get("/ddosx/v1/waf/logs/matches", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetWAFLogRequestMatches retrieves a list of log matches for request
func (s *Service) GetWAFLogRequestMatches(requestID string, parameters connection.APIRequestParameters) ([]WAFLogMatch, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
		return s.GetWAFLogRequestMatchesPaginated(requestID, p)
	}, parameters)
}

// GetWAFLogRequestMatchesPaginated retrieves a paginated list of matches for request
func (s *Service) GetWAFLogRequestMatchesPaginated(requestID string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
	body, err := s.getWAFLogRequestMatchesPaginatedResponseBody(requestID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFLogMatch], error) {
		return s.GetWAFLogRequestMatchesPaginated(requestID, p)
	}), err
}

func (s *Service) getWAFLogRequestMatchesPaginatedResponseBody(requestID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFLogMatch], error) {
	body := &connection.APIResponseBodyData[[]WAFLogMatch]{}

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

func (s *Service) getWAFLogRequestMatchResponseBody(requestID string, matchID string) (*connection.APIResponseBodyData[WAFLogMatch], error) {
	body := &connection.APIResponseBodyData[WAFLogMatch]{}

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
