package ecloudflex

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// ECloudFlexService is an interface for managing eCloud Flex
type ECloudFlexService interface {
	GetProjects(parameters connection.APIRequestParameters) ([]Project, error)
	GetProjectsPaginated(parameters connection.APIRequestParameters) (*PaginatedProject, error)
	GetProject(projectID int) (Project, error)
}

// Service implements ECloudFlexService for managing
// ECloudFlex certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of ECloudFlexService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
