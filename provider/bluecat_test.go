package provider

import (
	"strings"
	"testing"
)

type mockGatewayClient struct {
	mockBluecatZones *[]BluecatZone
	mockBluecatHosts *[]BluecatHostRecord
	mockBluecatCNAMEs *[]BluecatCNAMERecord
}

func (c *mockGatewayClient)

func createMockBluecatZone(fqdn string) BluecatZone {
	return BluecatZone{
		AbsoluteName: fqdn,
		Name:         strings.Split(fqdn, ".")[0],
		Deployable:   true,
	}
}

func newBluecatProvider(domainFilter DomainFilter, zoneIDFilter ZoneIDFilter, dryRun bool, client GatewayClient) *BluecatProvider {
	return &BluecatProvider{
		domainFilter: domainFilter,
		zoneIdFilter: zoneIDFilter,
		dryRun: dryRun,
		gatewayClient: client,
	}
}

func TestBluecatRecords(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones:  &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatHosts:  &[]BluecatHostRecord{},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{},
	}

	provider := newBluecatProvider(NewDomainFilter([]string{"example.com"}), NewZoneIDFilter([]string("")), true, client)

	actual, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}
}