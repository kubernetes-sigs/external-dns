package hetzner

import (
	"context"
	"fmt"
	"os"
	"testing"

	hclouddns "git.blindage.org/21h/hcloud-dns"

	"sigs.k8s.io/external-dns/endpoint"
)

type mockHetznerProvider struct {
	HetznerProvider
	Client mockHetznerClient
}

type mockHetznerClient struct {
	hclouddns.HCloudDNS
}

func (m *mockHetznerClient) GetZones(params hclouddns.HCloudGetZonesParams) (hclouddns.HCloudAnswerGetZones, error) {
	return hclouddns.HCloudAnswerGetZones{
		Zones: []hclouddns.HCloudZone{
			{
				ID:           "HetznerZoneID",
				Name:         "blindage.org",
				TTL:          666,
				RecordsCount: 1,
			},
		},
	}, nil
}

func (m *mockHetznerClient) GetRecords(params hclouddns.HCloudGetRecordsParams) (hclouddns.HCloudAnswerGetRecords, error) {
	return hclouddns.HCloudAnswerGetRecords{
		Records: []hclouddns.HCloudRecord{
			{
				ID:         "HetznerRecordID",
				RecordType: "A",
				Name:       "local",
				Value:      "127.0.0.1",
				TTL:        666,
				ZoneID:     "HetznerZoneImocked.Client.D",
			},
		},
	}, nil
}

func TestNewHetznerProvider(t *testing.T) {
	_ = os.Setenv("HETZNER_TOKEN", "myHetznerToken")
	_, err := NewHetznerProvider(context.Background(), endpoint.NewDomainFilter([]string{"blindage.org"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv("HETZNER_TOKEN")
	_, err = NewHetznerProvider(context.Background(), endpoint.NewDomainFilter([]string{"blindage.org"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestHetznerProvider_Records(t *testing.T) {
	mockedClient := mockHetznerClient{}
	mockedProvider := mockHetznerProvider{
		Client: mockedClient,
	}

	expectedZonesAnswer, err := mockedClient.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("expectedZonesAnswer:", expectedZonesAnswer)

	endpoints, err := mockedProvider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("endpoints:", endpoints)

	// if !reflect.DeepEqual(expectedZonesAnswer.Zones, endpoints) {
	// 	t.Fatal(err)
	// }
}
