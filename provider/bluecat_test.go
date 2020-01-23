package provider

import (
	"sigs.k8s.io/external-dns/endpoint"
	"strings"
	"testing"
)

type mockGatewayClient struct {
	mockBluecatZones *[]BluecatZone
	mockBluecatHosts *[]BluecatHostRecord
	mockBluecatCNAMEs *[]BluecatCNAMERecord
}

func createMockBluecatZone(fqdn string) BluecatZone {
	return BluecatZone{
		Name:         strings.Split(fqdn, ".")[0],
	}
}

func createMockBluecatHost(fqdn, target string) BluecatHostRecord {
	return BluecatHostRecord{
		Name:          fqdn,
		Addresses:     []string{target},
	}
}

func createMockBluecatCNAME(alias, target string) BluecatCNAMERecord {
	return BluecatCNAMERecord{
		Name:             alias,
		LinkedRecordName: target,
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
		mockBluecatHosts:  &[]BluecatHostRecord{
			createMockBluecatHost("example.com", "123.123.123.122"),
			createMockBluecatHost("nginx.example.com", "123.123.123.123"),
			createMockBluecatHost("whitespace.example.com", "123.123.123.124"),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com"),
		},
	}

	provider := newBluecatProvider(NewDomainFilter([]string{"example.com"}), NewZoneIDFilter([]string{""}), true, client)
	actual, err := provider.Records()

	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122"),
		endpoint.NewEndpoint("nginx.example.com", endpoint.RecordTypeA, "123.123.123.123"),
		endpoint.NewEndpoint("whitespace.example.com", endpoint.RecordTypeA, "123.123.123.124"),
		endpoint.NewEndpoint("hack.example.com", endpoint.RecordTypeCNAME, "bluecatnetworks.com"),
	}
	validateEndpoints(t, actual, expected)
}