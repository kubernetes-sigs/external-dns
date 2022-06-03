package storage

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// StorageService is an interface for managing UKFast Storage
type StorageService interface {
	GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error)
	GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error)
	GetSolution(certificateID int) (Solution, error)
	GetHosts(parameters connection.APIRequestParameters) ([]Host, error)
	GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error)
	GetHost(hostID int) (Host, error)
	GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error)
	GetVolumesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error)
	GetVolume(volumeID int) (Volume, error)
}

// Service implements StorageService for managing
// Storage certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of StorageService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
