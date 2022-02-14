package source

import (
	"context"
	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/external-dns/endpoint"
	"testing"
)

type mockSource struct {
	endpoints []*endpoint.Endpoint
}

func (m *mockSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return m.endpoints, nil
}

func (m mockSource) AddEventHandler(ctx context.Context, f func()) {
	// NOP
}

func MockSource(endpoints []*endpoint.Endpoint) Source {
	return &mockSource{
		endpoints: endpoints,
	}
}

type SuppressIPv6TestSuite struct {
	suite.Suite
}

func TestSuppressIPv6(t *testing.T) {
	suite.Run(t, new(SuppressIPv6TestSuite))
}

func (s *SuppressIPv6TestSuite) TestGetIPv4Targets() {
	s.Len(getIp4Targets([]string{"xxx.biz", "::1"}), 0)

	s.Equal([]string{"192.168.113.1"}, []string(
		getIp4Targets([]string{
			"192.168.113.1",
			"2001:470:1f09:370:b73b:1902:5ab9:e7d",
		},
		),
	))
}

func (s *SuppressIPv6TestSuite) TestEndpointsRemoved() {
	endpoints := []*endpoint.Endpoint{
		&endpoint.Endpoint{
			DNSName:    "notargets.tailify.com",
			Targets:    []string{"2001::1", "2002::2"},
			RecordType: "A",
			RecordTTL:  300,
		},
		&endpoint.Endpoint{
			DNSName:    "targets.tailify.com",
			Targets:    []string{"2001::1", "8.8.8.8"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	ms := MockSource(endpoints)

	ss := NewSuppressedSource(ms)

	expected := []*endpoint.Endpoint{
		&endpoint.Endpoint{
			DNSName:    "targets.tailify.com",
			Targets:    []string{"8.8.8.8"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}

	ctx := context.TODO()

	endps, err := ss.Endpoints(ctx)
	s.Nil(err)
	s.Equal(expected, endps)

}
