package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDatastores retrieves a list of datastores
func (s *Service) GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error) {
	return connection.InvokeRequestAll(s.GetDatastoresPaginated, parameters)
}

// GetDatastoresPaginated retrieves a paginated list of datastores
func (s *Service) GetDatastoresPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
	body, err := s.getDatastoresPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDatastoresPaginated), err
}

func (s *Service) getDatastoresPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Datastore], error) {
	body := &connection.APIResponseBodyData[[]Datastore]{}

	response, err := s.connection.Get("/ecloud/v1/datastores", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDatastore retrieves a single datastore by ID
func (s *Service) GetDatastore(datastoreID int) (Datastore, error) {
	body, err := s.getDatastoreResponseBody(datastoreID)

	return body.Data, err
}

func (s *Service) getDatastoreResponseBody(datastoreID int) (*connection.APIResponseBodyData[Datastore], error) {
	body := &connection.APIResponseBodyData[Datastore]{}

	if datastoreID < 1 {
		return body, fmt.Errorf("invalid datastore id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/datastores/%d", datastoreID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DatastoreNotFoundError{ID: datastoreID}
		}

		return nil
	})
}
