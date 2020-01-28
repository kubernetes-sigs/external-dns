package provider

import (
	"context"
	"sigs.k8s.io/external-dns/endpoint"
	"strings"
	"testing"
)

type mockGatewayClient struct {
	mockBluecatZones *[]BluecatZone
	mockBluecatHosts *[]BluecatHostRecord
	mockBluecatCNAMEs *[]BluecatCNAMERecord
}

func (g *mockGatewayClient) getBluecatZones() ([]BluecatZone, error) {
	return *g.mockBluecatZones, nil
}

func createMockBluecatZone(fqdn string) BluecatZone {
	return BluecatZone{
		Name:         strings.Split(fqdn, ".")[0],
	}
}

func createMockBluecatHostRecord(fqdn, target string) BluecatHostRecord {
	props := "absoluteName="+fqdn+"|addresses="+target+"|"
	nameParts := strings.Split(fqdn, ".")
	return BluecatHostRecord{
		Name:       nameParts[0],
		Properties: props,
	}
}

func createMockBluecatCNAME(alias, target string) BluecatCNAMERecord {
	props := "absoluteName="+alias+"|linkedRecordName="+target+"|"
	nameParts := strings.Split(alias, ".")
	return BluecatCNAMERecord{
		Name:       nameParts[0],
		Properties: props,
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
			createMockBluecatHostRecord("example.com", "123.123.123.122"),
			createMockBluecatHostRecord("nginx.example.com", "123.123.123.123"),
			createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124"),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com"),
		},
	}

	provider := newBluecatProvider(
		NewDomainFilter([]string{"example.com"}),
		NewZoneIDFilter([]string{""}), true, client)
	actual, err := provider.Records(context.Background())

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