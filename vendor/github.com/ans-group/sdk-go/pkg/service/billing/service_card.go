package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCards retrieves a list of cards
func (s *Service) GetCards(parameters connection.APIRequestParameters) ([]Card, error) {
	return connection.InvokeRequestAll(s.GetCardsPaginated, parameters)
}

// GetCardsPaginated retrieves a paginated list of cards
func (s *Service) GetCardsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Card], error) {
	body, err := s.getCardsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetCardsPaginated), err
}

func (s *Service) getCardsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Card], error) {
	body := &connection.APIResponseBodyData[[]Card]{}

	response, err := s.connection.Get("/billing/v1/cards", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCard retrieves a single card by id
func (s *Service) GetCard(cardID int) (Card, error) {
	body, err := s.getCardResponseBody(cardID)

	return body.Data, err
}

func (s *Service) getCardResponseBody(cardID int) (*connection.APIResponseBodyData[Card], error) {
	body := &connection.APIResponseBodyData[Card]{}

	if cardID < 1 {
		return body, fmt.Errorf("invalid card id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/cards/%d", cardID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CardNotFoundError{ID: cardID}
		}

		return nil
	})
}

// CreateCard creates a new card
func (s *Service) CreateCard(req CreateCardRequest) (int, error) {
	body, err := s.createCardResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createCardResponseBody(req CreateCardRequest) (*connection.APIResponseBodyData[Card], error) {
	body := &connection.APIResponseBodyData[Card]{}

	response, err := s.connection.Post("/billing/v1/cards", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchCard patches a card
func (s *Service) PatchCard(cardID int, patch PatchCardRequest) error {
	_, err := s.patchCardResponseBody(cardID, patch)

	return err
}

func (s *Service) patchCardResponseBody(cardID int, patch PatchCardRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if cardID < 1 {
		return body, fmt.Errorf("invalid card id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/billing/v1/cards/%d", cardID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CardNotFoundError{ID: cardID}
		}

		return nil
	})
}

// DeleteCard removes a card
func (s *Service) DeleteCard(cardID int) error {
	_, err := s.deleteCardResponseBody(cardID)

	return err
}

func (s *Service) deleteCardResponseBody(cardID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if cardID < 1 {
		return body, fmt.Errorf("invalid card id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/billing/v1/cards/%d", cardID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CardNotFoundError{ID: cardID}
		}

		return nil
	})
}
