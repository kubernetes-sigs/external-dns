package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDatastores retrieves a list of datastores
func (s *Service) GetDatastores(parameters connection.APIRequestParameters) ([]Datastore, error) {
	var datastores []Datastore

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDatastoresPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, datastore := range response.(*PaginatedDatastore).Items {
			datastores = append(datastores, datastore)
		}
	}

	return datastores, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetDatastoresPaginated retrieves a paginated list of datastores
func (s *Service) GetDatastoresPaginated(parameters connection.APIRequestParameters) (*PaginatedDatastore, error) {
	body, err := s.getDatastoresPaginatedResponseBody(parameters)

	return NewPaginatedDatastore(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDatastoresPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getDatastoresPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetDatastoreSliceResponseBody, error) {
	body := &GetDatastoreSliceResponseBody{}

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

func (s *Service) getDatastoreResponseBody(datastoreID int) (*GetDatastoreResponseBody, error) {
	body := &GetDatastoreResponseBody{}

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
