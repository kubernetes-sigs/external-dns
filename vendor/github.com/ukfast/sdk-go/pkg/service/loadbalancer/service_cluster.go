package loadbalancer

import (
	"errors"
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetClusters retrieves a list of clusters
func (s *Service) GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error) {
	var clusters []Cluster

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetClustersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, cluster := range response.(*PaginatedCluster).Items {
			clusters = append(clusters, cluster)
		}
	}

	return clusters, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetClustersPaginated retrieves a paginated list of clusters
func (s *Service) GetClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedCluster, error) {
	body, err := s.getClustersPaginatedResponseBody(parameters)

	return NewPaginatedCluster(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetClustersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getClustersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetClusterSliceResponseBody, error) {
	body := &GetClusterSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/clusters", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCluster retrieves a single cluster by id
func (s *Service) GetCluster(clusterID int) (Cluster, error) {
	body, err := s.getClusterResponseBody(clusterID)

	return body.Data, err
}

func (s *Service) getClusterResponseBody(clusterID int) (*GetClusterResponseBody, error) {
	body := &GetClusterResponseBody{}

	if clusterID < 1 {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/clusters/%d", clusterID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

// PatchCluster patches a Cluster
func (s *Service) PatchCluster(clusterID int, req PatchClusterRequest) error {
	_, err := s.patchClusterResponseBody(clusterID, req)

	return err
}

func (s *Service) patchClusterResponseBody(clusterID int, req PatchClusterRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clusterID < 1 {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/clusters/%d", clusterID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

// DeployCluster deploys a Cluster
func (s *Service) DeployCluster(clusterID int) error {
	_, err := s.deployClusterResponseBody(clusterID)

	return err
}

func (s *Service) deployClusterResponseBody(clusterID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clusterID < 1 {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/loadbalancers/v2/clusters/%d/deploy", clusterID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

// ValidateCluster validates a cluster
func (s *Service) ValidateCluster(clusterID int) error {
	response := &connection.APIResponse{}

	if clusterID < 1 {
		return fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/clusters/%d/validate", clusterID), connection.APIRequestParameters{})
	if err != nil {
		return err
	}

	if response.StatusCode == 422 {
		body := &validateClusterResponseBody{}
		err := response.DeserializeResponseBody(body)
		if err != nil {
			return err
		}

		return errors.New(body.ErrorString())
	}

	return response.HandleResponse(&connection.APIResponseBody{}, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}

type validateClusterResponseBody struct {
	Errors interface{} `json:"errors"`
}

func (r *validateClusterResponseBody) ErrorString() string {
	return fmt.Sprintf("%+v", r.Errors)
}
func (r *validateClusterResponseBody) Pagination() connection.APIResponseMetadataPagination {
	return connection.APIResponseMetadataPagination{}
}

// GetCluster retrieves a single cluster by id
func (s *Service) GetClusterACLTemplates(clusterID int) (ACLTemplates, error) {
	body, err := s.getClusterACLTemplatesResponseBody(clusterID)

	return body.Data, err
}

func (s *Service) getClusterACLTemplatesResponseBody(clusterID int) (*GetACLTemplatesResponseBody, error) {
	body := &GetACLTemplatesResponseBody{}

	if clusterID < 1 {
		return body, fmt.Errorf("invalid cluster id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/clusters/%d/acl-templates", clusterID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClusterNotFoundError{ID: clusterID}
		}

		return nil
	})
}
