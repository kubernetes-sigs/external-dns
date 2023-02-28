package stackit

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	stackitdnsclient "github.com/stackitcloud/stackit-dns-api-client-go"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"testing"
)

func TestRecords(t *testing.T) {
	t.Parallel()

	stackitProvider := getStackitProvider(t)
	stackitProvider.Client = mockClient{
		getZoneError:     false,
		getRecordsError:  false,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	endpoints, err := stackitProvider.Records(context.Background())
	assert.NoError(t, err)

	// 1 zone, 1 record but with 2 pages therefore 4 endpoints
	assert.Equal(t, 4, len(endpoints))

	// test with filter
	stackitProvider.DomainFilter = endpoint.NewDomainFilter([]string{"test.com"})
	endpoints, err = stackitProvider.Records(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 4, len(endpoints))
}

func TestApplyChange(t *testing.T) {
	t.Parallel()

	stackitProvider := getStackitProvider(t)
	stackitProvider.Client = mockClient{
		getZoneError:     false,
		getRecordsError:  false,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	changes := getExampleChanges()

	err := stackitProvider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)

	// test with dry run
	stackitProvider.DryRun = true
	err = stackitProvider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)

	// test endpoint not found
	changes.Create[0].DNSName = "notfound.test.com"
	err = stackitProvider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)

	// test zone not found
	changes.Create[0].DNSName = "notfound.com"
	err = stackitProvider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)

	// test no endpoints
	changes = &plan.Changes{}
	err = stackitProvider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)
}

// TestClientInterfaceError tests the client interface error. Just for coverage.
func TestClientInterfaceError(t *testing.T) {
	t.Parallel()
	stackitProvider := getStackitProvider(t)

	_, err := stackitProvider.Client.GetRRSets(context.Background(), "", "", nil)
	assert.Error(t, err)

	_, err = stackitProvider.Client.GetZones(context.Background(), "", nil)
	assert.Error(t, err)

	_, err = stackitProvider.Client.CreateRRSet(
		context.Background(),
		stackitdnsclient.RrsetRrSetPost{},
		"",
		"",
	)
	assert.Error(t, err)

	_, err = stackitProvider.Client.UpdateRRSet(
		context.Background(),
		stackitdnsclient.RrsetRrSetPatch{},
		"",
		"",
		"",
	)
	assert.Error(t, err)

	_, err = stackitProvider.Client.DeleteRRSet(context.Background(), "", "", "")
	assert.Error(t, err)
}

func TestRecordsError(t *testing.T) {
	t.Parallel()

	// get zones returns error
	stackitProvider := getStackitProvider(t)
	stackitProvider.Client = mockClient{
		getZoneError:     true,
		getRecordsError:  false,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	_, err := stackitProvider.Records(context.Background())
	assert.Error(t, err)

	// get zones work but records fail.
	stackitProvider.Client = mockClient{
		getZoneError:     false,
		getRecordsError:  true,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	_, err = stackitProvider.Records(context.Background())
	assert.Error(t, err)
}

func TestApplyChangeError(t *testing.T) {
	t.Parallel()

	stackitProvider := getStackitProvider(t)

	// get zones error
	stackitProvider.Client = mockClient{
		getZoneError:     true,
		getRecordsError:  false,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	changes := getExampleChanges()

	err := stackitProvider.ApplyChanges(context.Background(), changes)
	assert.Error(t, err)

	// get zones work but records fail
	stackitProvider.Client = mockClient{
		getZoneError:     false,
		getRecordsError:  true,
		createRRSetError: false,
		updateRRSetError: false,
		deleteRRSetError: false,
	}

	err = stackitProvider.ApplyChanges(context.Background(), changes)
	// we ignore errors therefore check no error
	assert.NoError(t, err)

	// get zones and records work but create/update/delete fail
	stackitProvider.Client = mockClient{
		getZoneError:     false,
		getRecordsError:  false,
		createRRSetError: true,
		updateRRSetError: true,
		deleteRRSetError: true,
	}

	err = stackitProvider.ApplyChanges(context.Background(), changes)
	// we ignore errors therefore check no error
	assert.NoError(t, err)
}

func getExampleChanges() *plan.Changes {
	endpoints := getExampleEndpoints()

	return &plan.Changes{
		Create:    endpoints,
		UpdateOld: endpoints,
		UpdateNew: endpoints,
		Delete:    endpoints,
	}
}

func getExampleEndpoints() []*endpoint.Endpoint {
	return []*endpoint.Endpoint{
		{
			DNSName:          "test.test.com",
			Targets:          []string{"1.2.3.4"},
			RecordType:       "A",
			SetIdentifier:    "",
			RecordTTL:        0,
			Labels:           nil,
			ProviderSpecific: nil,
		},
	}
}

func getStackitProvider(t *testing.T) *StackitDNSProvider {
	provider, err := NewStackitDNSProvider(endpoint.NewDomainFilter([]string{}), false, Config{
		BasePath:  "test.com",
		Token:     "",
		ProjectId: uuid.NewString(),
	})
	assert.NoError(t, err)

	return provider
}

// mockClient to mock the api requests.
type mockClient struct {
	getZoneError     bool
	getRecordsError  bool
	createRRSetError bool
	updateRRSetError bool
	deleteRRSetError bool
}

var ErrResponse = fmt.Errorf("error")

func (m mockClient) GetZones(
	ctx context.Context,
	projectId string,
	localVarOptionals *stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts,
) (stackitdnsclient.ZoneResponseZoneAll, error) {
	if m.getZoneError {
		return stackitdnsclient.ZoneResponseZoneAll{}, ErrResponse
	}

	return stackitdnsclient.ZoneResponseZoneAll{
		ItemsPerPage: 10,
		Message:      "success",
		TotalItems:   1,
		TotalPages:   2,
		Zones: []stackitdnsclient.DomainZone{{
			Acl:               "0.0.0.0/0",
			ContactEmail:      "hostmaster.at.stackit.cloud",
			DefaultTTL:        300,
			Description:       "test",
			DnsName:           "test.com",
			Error_:            "",
			ExpireTime:        300,
			Id:                uuid.NewString(),
			IsReverseZone:     false,
			Name:              "test",
			NegativeCache:     300,
			Primaries:         nil,
			PrimaryNameServer: "",
			RecordCount:       1,
			RefreshTime:       300,
			RetryTime:         300,
			SerialNumber:      2023022300,
			State:             "CREATE_SUCCEEDED",
			Type_:             "primary",
			Visibility:        "public",
		}},
	}, nil
}

func (m mockClient) GetRRSets(
	ctx context.Context,
	projectId string,
	zoneId string,
	localVarOptionals *stackitdnsclient.RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts,
) (stackitdnsclient.RrsetResponseRrSetAll, error) {
	if m.getRecordsError {
		return stackitdnsclient.RrsetResponseRrSetAll{}, ErrResponse
	}

	return stackitdnsclient.RrsetResponseRrSetAll{
		ItemsPerPage: 10,
		Message:      "success",
		RrSets: []stackitdnsclient.DomainRrSet{{
			Active:  true,
			Comment: "test",
			Error_:  "",
			Id:      uuid.NewString(),
			Name:    "test.test.com.",
			Records: []stackitdnsclient.DomainRecord{{
				Content: "1.2.3.4",
				Id:      uuid.NewString(),
			}},
			State: "CREATE_SUCCEEDED",
			Ttl:   200,
			Type_: "A",
		}},
		TotalItems: 1,
		TotalPages: 2,
	}, nil
}

func (m mockClient) CreateRRSet(
	ctx context.Context,
	body stackitdnsclient.RrsetRrSetPost,
	projectId string,
	zoneId string,
) (stackitdnsclient.RrsetResponseRrSet, error) {
	if m.createRRSetError {
		return stackitdnsclient.RrsetResponseRrSet{}, ErrResponse
	}

	return stackitdnsclient.RrsetResponseRrSet{
		Message: "success",
		Rrset: &stackitdnsclient.DomainRrSet{
			Active:  true,
			Comment: "test",
			Error_:  "",
			Id:      uuid.NewString(),
			Name:    "test.test.com.",
			Records: []stackitdnsclient.DomainRecord{{
				Content: "1.2.3.4",
				Id:      uuid.NewString(),
			}},
			State: "CREATE_SUCCEEDED",
			Ttl:   200,
			Type_: "A",
		},
	}, nil
}

func (m mockClient) UpdateRRSet(
	ctx context.Context,
	body stackitdnsclient.RrsetRrSetPatch,
	projectId string,
	zoneId string,
	rrsetId string,
) (stackitdnsclient.RrsetResponseRrSet, error) {
	if m.updateRRSetError {
		return stackitdnsclient.RrsetResponseRrSet{}, ErrResponse
	}

	return stackitdnsclient.RrsetResponseRrSet{
		Message: "success",
		Rrset: &stackitdnsclient.DomainRrSet{
			Active:  true,
			Comment: "test",
			Error_:  "",
			Id:      uuid.NewString(),
			Name:    "test.test.com.",
			Records: []stackitdnsclient.DomainRecord{{
				Content: "1.2.3.4",
				Id:      uuid.NewString(),
			}},
			State: "CREATE_SUCCEEDED",
			Ttl:   200,
			Type_: "A",
		},
	}, nil
}

func (m mockClient) DeleteRRSet(
	ctx context.Context,
	projectId string,
	zoneId string,
	rrsetId string,
) (stackitdnsclient.SerializerMessage, error) {
	if m.deleteRRSetError {
		return stackitdnsclient.SerializerMessage{}, ErrResponse
	}

	return stackitdnsclient.SerializerMessage{Message: "success"}, nil
}
