package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LoadBalancerService is an interface for managing the LoadBalancer service
type LoadBalancerService interface {
	// Cluster
	GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error)
	GetClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedCluster, error)
	GetCluster(clusterID int) (Cluster, error)
	PatchCluster(clusterID int, req PatchClusterRequest) error
	DeployCluster(clusterID int) error
	ValidateCluster(clusterID int) error

	// Cluster ACL Templates
	GetClusterACLTemplates(clusterID int) (ACLTemplates, error)

	// Target Group
	GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error)
	GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedTargetGroup, error)
	GetTargetGroup(groupID int) (TargetGroup, error)
	CreateTargetGroup(req CreateTargetGroupRequest) (int, error)
	PatchTargetGroup(groupID int, req PatchTargetGroupRequest) error
	DeleteTargetGroup(groupID int) error

	// Target Group ACL
	GetTargetGroupACLs(targetGroupID int, parameters connection.APIRequestParameters) ([]ACL, error)
	GetTargetGroupACLsPaginated(targetGroupID int, parameters connection.APIRequestParameters) (*PaginatedACL, error)

	// Target Group Target
	GetTargetGroupTargets(groupID int, parameters connection.APIRequestParameters) ([]Target, error)
	GetTargetGroupTargetsPaginated(groupID int, parameters connection.APIRequestParameters) (*PaginatedTarget, error)
	GetTargetGroupTarget(groupID int, targetID int) (Target, error)
	CreateTargetGroupTarget(groupID int, req CreateTargetRequest) (int, error)
	PatchTargetGroupTarget(groupID int, targetID int, req PatchTargetRequest) error
	DeleteTargetGroupTarget(groupID int, targetID int) error

	// VIP
	GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error)
	GetVIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedVIP, error)
	GetVIP(vipID int) (VIP, error)

	// Listener
	GetListeners(parameters connection.APIRequestParameters) ([]Listener, error)
	GetListenersPaginated(parameters connection.APIRequestParameters) (*PaginatedListener, error)
	GetListener(listenerID int) (Listener, error)
	CreateListener(req CreateListenerRequest) (int, error)
	PatchListener(listenerID int, req PatchListenerRequest) error
	DeleteListener(listenerID int) error

	// Listener ACL
	GetListenerACLs(listenerID int, parameters connection.APIRequestParameters) ([]ACL, error)
	GetListenerACLsPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedACL, error)

	// Listener Access IP
	GetListenerAccessIPs(listenerID int, parameters connection.APIRequestParameters) ([]AccessIP, error)
	GetListenerAccessIPsPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedAccessIP, error)
	CreateListenerAccessIP(listenerID int, req CreateAccessIPRequest) (int, error)

	// Listener Bind
	GetListenerBinds(listenerID int, parameters connection.APIRequestParameters) ([]Bind, error)
	GetListenerBindsPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedBind, error)
	GetListenerBind(listenerID int, bindID int) (Bind, error)
	CreateListenerBind(listenerID int, req CreateBindRequest) (int, error)
	PatchListenerBind(listenerID int, bindID int, req PatchBindRequest) error
	DeleteListenerBind(listenerID int, bindID int) error

	// Access IP
	GetAccessIP(accessIP int) (AccessIP, error)
	PatchAccessIP(accessIP int, req PatchAccessIPRequest) error
	DeleteAccessIP(accessIP int) error

	// Listener Certificate
	GetListenerCertificates(listenerID int, parameters connection.APIRequestParameters) ([]Certificate, error)
	GetListenerCertificatesPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedCertificate, error)
	GetListenerCertificate(listenerID int, certificateID int) (Certificate, error)
	CreateListenerCertificate(listenerID int, req CreateCertificateRequest) (int, error)
	PatchListenerCertificate(listenerID int, certificateID int, req PatchCertificateRequest) error
	DeleteListenerCertificate(listenerID int, certificateID int) error

	// Bind
	GetBinds(parameters connection.APIRequestParameters) ([]Bind, error)
	GetBindsPaginated(parameters connection.APIRequestParameters) (*PaginatedBind, error)

	// Header
	GetHeaders(parameters connection.APIRequestParameters) ([]Header, error)
	GetHeadersPaginated(parameters connection.APIRequestParameters) (*PaginatedHeader, error)

	// ACL
	GetACLs(parameters connection.APIRequestParameters) ([]ACL, error)
	GetACLsPaginated(parameters connection.APIRequestParameters) (*PaginatedACL, error)
	GetACL(aclID int) (ACL, error)
	CreateACL(req CreateACLRequest) (int, error)
	PatchACL(aclID int, req PatchACLRequest) error
	DeleteACL(aclID int) error
}

// Service implements LoadBalancerService for managing
// LoadBalancer certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of LoadBalancerService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
