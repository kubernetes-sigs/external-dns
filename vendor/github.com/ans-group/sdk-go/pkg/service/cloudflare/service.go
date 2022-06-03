package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// CloudflareService is an interface for managing Cloudflare services
type CloudflareService interface {
	// Account
	GetAccounts(parameters connection.APIRequestParameters) ([]Account, error)
	GetAccountsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Account], error)
	GetAccount(accountID string) (Account, error)
	CreateAccount(req CreateAccountRequest) (string, error)
	PatchAccount(accountID string, req PatchAccountRequest) error
	CreateAccountMember(accountID string, req CreateAccountMemberRequest) error

	// Orchestration
	CreateOrchestration(req CreateOrchestrationRequest) error

	// Spend plan
	GetSpendPlans(parameters connection.APIRequestParameters) ([]SpendPlan, error)
	GetSpendPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SpendPlan], error)

	// Subscription
	GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error)
	GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Subscription], error)

	// Zone
	GetZones(parameters connection.APIRequestParameters) ([]Zone, error)
	GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error)
	GetZone(zoneID string) (Zone, error)
	CreateZone(req CreateZoneRequest) (string, error)
	PatchZone(zoneID string, req PatchZoneRequest) error
	DeleteZone(zoneID string) error

	// Spend
	GetTotalSpendMonthToDate() (TotalSpend, error)
}

// Service implements CloudflareService for managing the Shared Exchange service
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of CloudflareService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
