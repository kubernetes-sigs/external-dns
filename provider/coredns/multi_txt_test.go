package coredns

import (
	"context"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// mockClient implements coreDNSClient for testing
type mockClient struct {
	services map[string]*Service
}

func (m *mockClient) GetServices(prefix string) ([]*Service, error) {
	var result []*Service
	for key, service := range m.services {
		if key[:len(prefix)] == prefix {
			svc := *service
			svc.Key = key
			result = append(result, &svc)
		}
	}
	return result, nil
}

func (m *mockClient) SaveService(service *Service) error {
	if m.services == nil {
		m.services = make(map[string]*Service)
	}
	m.services[service.Key] = service
	return nil
}

func (m *mockClient) DeleteService(key string) error {
	delete(m.services, key)
	return nil
}

func TestCoreDNSProviderMultiTXT(t *testing.T) {
	client := &mockClient{}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/dev/test/",
		domainFilter:  nil,
	}

	// Test multi-target TXT record creation
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "example.test.dev",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  30,
			Targets:    []string{"v=1", "key=value", "third=string"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{
		Create: desired,
	}

	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Fatalf("ApplyChanges failed: %v", err)
	}

	// Verify three separate etcd keys were created
	if len(client.services) != 3 {
		t.Errorf("Expected 3 services, got %d", len(client.services))
		for k, v := range client.services {
			t.Logf("Service key: %s, text: %s", k, v.Text)
		}
	}

	// Verify each target has its own service
	expectedTexts := map[string]bool{"v=1": false, "key=value": false, "third=string": false}
	for _, service := range client.services {
		if _, exists := expectedTexts[service.Text]; exists {
			expectedTexts[service.Text] = true
		}
	}

	for text, found := range expectedTexts {
		if !found {
			t.Errorf("Expected TXT value %q not found in services", text)
		}
	}
}
