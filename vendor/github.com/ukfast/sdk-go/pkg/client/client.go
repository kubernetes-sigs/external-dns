package client

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/account"
	"github.com/ukfast/sdk-go/pkg/service/billing"
	"github.com/ukfast/sdk-go/pkg/service/ddosx"
	"github.com/ukfast/sdk-go/pkg/service/draas"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
	"github.com/ukfast/sdk-go/pkg/service/ecloudflex"
	"github.com/ukfast/sdk-go/pkg/service/loadbalancer"
	"github.com/ukfast/sdk-go/pkg/service/pss"
	"github.com/ukfast/sdk-go/pkg/service/registrar"
	"github.com/ukfast/sdk-go/pkg/service/safedns"
	"github.com/ukfast/sdk-go/pkg/service/sharedexchange"
	"github.com/ukfast/sdk-go/pkg/service/ssl"
	"github.com/ukfast/sdk-go/pkg/service/storage"
)

type Client interface {
	AccountService() account.AccountService
	BillingService() billing.BillingService
	DDoSXService() ddosx.DDoSXService
	DRaaSService() draas.DRaaSService
	ECloudService() ecloud.ECloudService
	ECloudFlexService() ecloudflex.ECloudFlexService
	LoadBalancerService() loadbalancer.LoadBalancerService
	PSSService() pss.PSSService
	RegistrarService() registrar.RegistrarService
	SafeDNSService() safedns.SafeDNSService
	SharedExchangeService() sharedexchange.SharedExchangeService
	SSLService() ssl.SSLService
	StorageService() storage.StorageService
}

type UKFastClient struct {
	connection connection.Connection
}

func NewClient(connection connection.Connection) *UKFastClient {
	return &UKFastClient{
		connection: connection,
	}
}

func (c *UKFastClient) AccountService() account.AccountService {
	return account.NewService(c.connection)
}

func (c *UKFastClient) BillingService() billing.BillingService {
	return billing.NewService(c.connection)
}

func (c *UKFastClient) DDoSXService() ddosx.DDoSXService {
	return ddosx.NewService(c.connection)
}

func (c *UKFastClient) DRaaSService() draas.DRaaSService {
	return draas.NewService(c.connection)
}

func (c *UKFastClient) ECloudService() ecloud.ECloudService {
	return ecloud.NewService(c.connection)
}

func (c *UKFastClient) ECloudFlexService() ecloudflex.ECloudFlexService {
	return ecloudflex.NewService(c.connection)
}

func (c *UKFastClient) LoadBalancerService() loadbalancer.LoadBalancerService {
	return loadbalancer.NewService(c.connection)
}

func (c *UKFastClient) PSSService() pss.PSSService {
	return pss.NewService(c.connection)
}

func (c *UKFastClient) RegistrarService() registrar.RegistrarService {
	return registrar.NewService(c.connection)
}

func (c *UKFastClient) SafeDNSService() safedns.SafeDNSService {
	return safedns.NewService(c.connection)
}

func (c *UKFastClient) SharedExchangeService() sharedexchange.SharedExchangeService {
	return sharedexchange.NewService(c.connection)
}

func (c *UKFastClient) SSLService() ssl.SSLService {
	return ssl.NewService(c.connection)
}

func (c *UKFastClient) StorageService() storage.StorageService {
	return storage.NewService(c.connection)
}
