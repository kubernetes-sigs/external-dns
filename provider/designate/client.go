package designate

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	"github.com/gophercloud/gophercloud/pagination"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/pkg/tlsutils"
)

// interface between provider and OpenStack DNS API
type designateClientInterface interface {
	// ForEachZone calls handler for each zone managed by the Designate
	ForEachZone(handler func(zone *zones.Zone) error) error

	// ForEachRecordSet calls handler for each recordset in the given DNS zone
	ForEachRecordSet(zoneID string, handler func(recordSet *recordsets.RecordSet) error) error

	// CreateRecordSet creates recordset in the given DNS zone
	CreateRecordSet(zoneID string, opts recordsets.CreateOpts) (string, error)

	// UpdateRecordSet updates recordset in the given DNS zone
	UpdateRecordSet(zoneID, recordSetID string, opts recordsets.UpdateOpts) error

	// DeleteRecordSet deletes recordset in the given DNS zone
	DeleteRecordSet(zoneID, recordSetID string) error
}

// implementation of the designateClientInterface
type designateClient struct {
	serviceClient *gophercloud.ServiceClient
}

// factory function for the designateClientInterface
func newDesignateClient() (designateClientInterface, error) {
	serviceClient, err := createDesignateServiceClient()
	if err != nil {
		return nil, err
	}
	return &designateClient{serviceClient}, nil
}

// copies environment variables to new names without overwriting existing values
func remapEnv(mapping map[string]string) {
	for k, v := range mapping {
		currentVal := os.Getenv(k)
		newVal := os.Getenv(v)
		if currentVal == "" && newVal != "" {
			os.Setenv(k, newVal)
		}
	}
}

// returns OpenStack Keystone authentication settings by obtaining values from standard environment variables.
// also fixes incompatibilities between gophercloud implementation and *-stackrc files that can be downloaded
// from OpenStack dashboard in latest versions
func getAuthSettings() (gophercloud.AuthOptions, error) {
	remapEnv(map[string]string{
		"OS_TENANT_NAME": "OS_PROJECT_NAME",
		"OS_TENANT_ID":   "OS_PROJECT_ID",
		"OS_DOMAIN_NAME": "OS_USER_DOMAIN_NAME",
		"OS_DOMAIN_ID":   "OS_USER_DOMAIN_ID",
	})

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return gophercloud.AuthOptions{}, err
	}
	opts.AllowReauth = true
	if !strings.HasSuffix(opts.IdentityEndpoint, "/") {
		opts.IdentityEndpoint += "/"
	}
	if !strings.HasSuffix(opts.IdentityEndpoint, "/v2.0/") && !strings.HasSuffix(opts.IdentityEndpoint, "/v3/") {
		opts.IdentityEndpoint += "v2.0/"
	}
	return opts, nil
}

// authenticate in OpenStack and obtain Designate service endpoint
func createDesignateServiceClient() (*gophercloud.ServiceClient, error) {
	opts, err := getAuthSettings()
	if err != nil {
		return nil, err
	}
	log.Infof("Using OpenStack Keystone at %s", opts.IdentityEndpoint)
	authProvider, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := tlsutils.CreateTLSConfig("OPENSTACK")
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
	authProvider.HTTPClient.Transport = transport

	if err = openstack.Authenticate(authProvider, opts); err != nil {
		return nil, err
	}

	eo := gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	}

	client, err := openstack.NewDNSV2(authProvider, eo)
	if err != nil {
		return nil, err
	}
	log.Infof("Found OpenStack Designate service at %s", client.Endpoint)
	return client, nil
}

// ForEachZone calls handler for each zone managed by the Designate
func (c designateClient) ForEachZone(handler func(zone *zones.Zone) error) error {
	pager := zones.List(c.serviceClient, zones.ListOpts{})
	return pager.EachPage(
		func(page pagination.Page) (bool, error) {
			list, err := zones.ExtractZones(page)
			if err != nil {
				return false, err
			}
			for _, zone := range list {
				err := handler(&zone)
				if err != nil {
					return false, err
				}
			}
			return true, nil
		},
	)
}

// ForEachRecordSet calls handler for each recordset in the given DNS zone
func (c designateClient) ForEachRecordSet(zoneID string, handler func(recordSet *recordsets.RecordSet) error) error {
	pager := recordsets.ListByZone(c.serviceClient, zoneID, recordsets.ListOpts{})
	return pager.EachPage(
		func(page pagination.Page) (bool, error) {
			list, err := recordsets.ExtractRecordSets(page)
			if err != nil {
				return false, err
			}
			for _, recordSet := range list {
				err := handler(&recordSet)
				if err != nil {
					return false, err
				}
			}
			return true, nil
		},
	)
}

// CreateRecordSet creates recordset in the given DNS zone
func (c designateClient) CreateRecordSet(zoneID string, opts recordsets.CreateOpts) (string, error) {
	r, err := recordsets.Create(c.serviceClient, zoneID, opts).Extract()
	if err != nil {
		return "", err
	}
	return r.ID, nil
}

// UpdateRecordSet updates recordset in the given DNS zone
func (c designateClient) UpdateRecordSet(zoneID, recordSetID string, opts recordsets.UpdateOpts) error {
	_, err := recordsets.Update(c.serviceClient, zoneID, recordSetID, opts).Extract()
	return err
}

// DeleteRecordSet deletes recordset in the given DNS zone
func (c designateClient) DeleteRecordSet(zoneID, recordSetID string) error {
	return recordsets.Delete(c.serviceClient, zoneID, recordSetID).ExtractErr()
}
