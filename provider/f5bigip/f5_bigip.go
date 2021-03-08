/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package f5bigip

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/f5devcentral/go-bigip"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	contentType                  = "application/json"
	providerSpecificPropertyName = "bigip/virtual-server"
)

var (
	partition = "Common"
)

// Config for F5Bigip DNS Provider
type F5BigipDNSConfig struct {
	DryRun       bool
	DomainFilter endpoint.DomainFilter
	Auth         AuthConfig
}

// AuthConfig to login to F5 DNS Service
type AuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// BigIp wideIPS
type WideIPs struct {
	Items []WideIP `json:"items"`
}

// BigIp wideIP
type WideIP struct {
	Name           string        `json:"name"`
	Partition      string        `json:"partition"`
	Description    string        `json:"description"`
	Enabled        bool          `json:"enabled"`
	LastResortPool string        `json:"lastResortPool"`
	PoolAs         []WideIPPoolA `json:"pools"`
}

// BigIp wideIP Pool A
type WideIPPoolA struct {
	Name      string       `json:"name"`
	Partition string       `json:"partition"`
	Ttl       endpoint.TTL `json:"ttl"`
	Ratio     int          `json:"ratio"`
}

// BigIp Pool A
type PoolA struct {
	Name                     string       `json:"name"`
	AlternateMode            string       `json:"alternateMode"`
	Partition                string       `json:"partition"`
	Enabled                  bool         `json:"enabled"`
	FallbackIp               string       `json:"fallbackIp"`
	FallbackMode             string       `json:"fallbackMode"`
	LoadBalancingMode        string       `json:"loadBalancingMode"`
	Ttl                      endpoint.TTL `json:"ttl"`
	VerifyMemberAvailability string       `json:"verifyMemberAvailability"`
}

// BigIp wideIP Pool A Members
type PoolAMembers struct {
	Items []PoolAMember `json:"items"`
}

// BigIp wideIP Pool A Member
type PoolAMember struct {
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	Enabled     bool   `json:"enabled"`
	MemberOrder int    `json:"memberOrder"`
	Monitor     string `json:"monitor"`
	Ratio       int    `json:"ratio"`
}

type F5BigipClient interface {
	APICall(options *bigip.APIRequest) ([]byte, error)
	GetGtmserver(name string) (*bigip.Server, error)
	UpdateGtmserver(name string, p *bigip.Server) error
}

// F5BigipDNSProvider is the bigip DNS provider
type F5BigipProvider struct {
	provider.BaseProvider
	client F5BigipClient
	config *F5BigipDNSConfig
}

//NewF5BigipProvider initializes a new F5Bigip client based Provider.
func newF5BigipClient() (*bigip.BigIP, error) {
	host := os.Getenv("F5BIGIP_HOST")
	port := os.Getenv("F5BIGIP_PORT")
	user := os.Getenv("F5BIGIP_USER")
	passwd := os.Getenv("F5BIGIP_PASSWD")

	configOptions := &bigip.ConfigOptions{
		APICallTimeout: 60 * time.Second,
	}

	return bigip.NewSession(host, port, user, passwd, configOptions), nil
}

// NewF5DNSProvider to instantiate F5 provider
func NewF5bigipProvider(cfg F5BigipDNSConfig) (*F5BigipProvider, error) {
	client, err := newF5BigipClient()
	if err != nil {
		return nil, err
	}
	// Create a server for datacenter
	err = createServer(client)
	if err != nil {
		return nil, err
	}
	return &F5BigipProvider{
		client: client,
		config: &cfg,
	}, nil
}

// First find the pool of wideip
// The second step is to find the host of the virtual server in the pool
// The third step is to add the corresponding host to the result
func (p F5BigipProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	// Get bigip A record wide IP
	wideIPs, err := p.getWideIPs()
	if err != nil {
		return nil, err
	}
	for _, wideIP := range wideIPs {
		dnsName := wideIP.Name
		for _, poolA := range wideIP.PoolAs {
			// Get bigip pool A members and ttl
			poolAMembers, ttl, err := p.getPoolAMemberAndTtl(poolA)
			if err != nil {
				return nil, err
			}
			// Get vs address via pool A member
			target, err := p.getPoolAMemberVSRecords(poolAMembers)
			if err != nil {
				return nil, err
			}
			ep := endpoint.NewEndpointWithTTL(
				dnsName,
				endpoint.RecordTypeA,
				endpoint.TTL(ttl),
				target...,
			)
			ep.WithProviderSpecific(providerSpecificPropertyName, wideIP.Description)
			result = append(result, ep)
		}
	}
	return filter(result, p.config.DomainFilter), nil
}

// ApplyChanges stores changes back to bigip converting them to bigip format and aggregating A and TXT records.
func (p *F5BigipProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Create
	if len(changes.Create) > 0 {
		// "Replacing" non-existent records creates them
		err := p.createRecords(changes.Create)
		if err != nil {
			return err
		}
	}

	// Update Old
	if len(changes.UpdateOld) > 0 {
		for _, change := range changes.UpdateOld {
			err := p.deleteServer(change)
			if err != nil {
				log.Errorf("Delete Server err: %v", err)
			}
		}
	}

	// Update New
	if len(changes.UpdateNew) > 0 {
		var allVsRecords []bigip.VSrecord
		// The key is vsRecord.Name values is virtual servers
		mapVsRecords := make(map[string][]bigip.VSrecord)
		for _, change := range changes.UpdateNew {
			for _, item := range change.ProviderSpecific {
				if item.Name == providerSpecificPropertyName {
					var vsRecordsWithCluster map[string][]bigip.VSrecord
					err := json.Unmarshal([]byte(item.Value), &vsRecordsWithCluster)
					if err != nil {
						continue
					}
					vsRecords := filterVsrecords(vsRecordsWithCluster)
					for _, vsRecord := range vsRecords {
						mapVsRecords[change.DNSName] = append(mapVsRecords[change.DNSName], vsRecord)
						if isContain(allVsRecords, vsRecord, false, false) {
							continue
						}
						allVsRecords = append(allVsRecords, vsRecord)
					}
				}
			}
		}
		err := p.updateServer(allVsRecords)
		if err != nil {
			return err
		}
		for _, change := range changes.UpdateNew {
			err := p.updatePoolA(change)
			if err != nil {
				log.Errorf("Update pool A err: %v", err)
			}
			err = p.updatePoolAMember(mapVsRecords[change.DNSName], change.DNSName)
			if err != nil {
				log.Errorf("Update pool member err: %v", err)
			}
			err = p.updateWideIP(change)
			if err != nil {
				log.Errorf("Update wide IP err: %v", err)
			}
		}
	}

	// Delete
	if len(changes.Delete) > 0 {
		for _, change := range changes.Delete {
			err := p.deleteServer(change)
			if err != nil {
				log.Errorf("Delete Server err: %v", err)
			}
			deleteWideipAOptions := &bigip.APIRequest{
				"delete",
				fmt.Sprintf("/mgmt/tm/gtm/wideip/a/~%s~%s", partition, change.DNSName),
				"",
				contentType,
			}
			err = p.delete(deleteWideipAOptions, change.DNSName)
			if err != nil {
				log.Errorf("Delete wideip a err: %v", err)
			}
			deletePoolAOptions := &bigip.APIRequest{
				"delete",
				fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s", partition, change.DNSName),
				"",
				contentType,
			}
			err = p.delete(deletePoolAOptions, change.DNSName)
			if err != nil {
				log.Errorf("Delete pool a err: %v", err)
			}
		}
	}
	return nil
}

func (p F5BigipProvider) getWideIPs() ([]WideIP, error) {
	var wideIps WideIPs
	wideipOptions := &bigip.APIRequest{
		"get",
		"mgmt/tm/gtm/wideip/a",
		"",
		contentType,
	}
	resp, err := p.client.APICall(wideipOptions)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &wideIps)
	if err != nil {
		return nil, err
	}

	return wideIps.Items, nil
}

func (p F5BigipProvider) getWideIP(name string) (WideIP, error) {
	var wideIP WideIP
	wideIPs, err := p.getWideIPs()
	if err != nil {
		return wideIP, err
	}
	for _, item := range wideIPs {
		if item.Name == name {
			wideIP = item
			break
		}
	}
	return wideIP, nil
}

func (p F5BigipProvider) getPoolAMemberAndTtl(poolA WideIPPoolA) ([]PoolAMember, endpoint.TTL, error) {
	var poolMember PoolAMembers
	poolOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s", poolA.Partition, poolA.Name),
		"",
		contentType,
	}
	respPool, err := p.client.APICall(poolOptions)
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(respPool, &poolA)
	if err != nil {
		return nil, 0, err
	}
	// pool member is virtual server
	poolMemberOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s/members", poolA.Partition, poolA.Name),
		"",
		"application/json",
	}
	respPoolMember, err := p.client.APICall(poolMemberOptions)
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(respPoolMember, &poolMember)
	if err != nil {
		return nil, 0, err
	}
	return poolMember.Items, poolA.Ttl, nil
}

func (p F5BigipProvider) getPoolAMemberVSRecords(poolAMembers []PoolAMember) ([]string, error) {
	var target []string
	for _, member := range poolAMembers {
		nameArr := strings.Split(member.Name, ":")
		if len(nameArr) == 2 {
			var vsRecord bigip.VSrecord
			serverName := os.Getenv("F5BIGIP_SERVER_NAME")
			virtualServerOptions := &bigip.APIRequest{
				"get",
				fmt.Sprintf("/mgmt/tm/gtm/server/~%s~%s/virtual-servers/%s", member.Partition, serverName, nameArr[1]),
				"",
				"application/json",
			}
			respPoolvirtualServer, err := p.client.APICall(virtualServerOptions)
			if err != nil {
				log.Errorf("virtual Server error: %s", err)
			}

			err = json.Unmarshal(respPoolvirtualServer, &vsRecord)
			if err != nil {
				return nil, err
			}
			// Obtain the address of the ipv4
			r := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?:\/\d{2})?`)
			host := r.FindString(vsRecord.Destination)
			target = append(target, host)
		} else {
			err := fmt.Errorf("%s unable to extract the virtual server name", member.Name)
			return nil, err
		}
	}
	return target, nil
}

func (p F5BigipProvider) createRecords(endpoints []*endpoint.Endpoint) error {
	// To create a server need to associate all virtual servers
	var allVsRecords []bigip.VSrecord
	// The key is vsRecord.Name values is virtual servers
	mapVsRecords := make(map[string][]bigip.VSrecord)
	for _, endpoint := range endpoints {
		if !p.config.DomainFilter.Match(endpoint.DNSName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", endpoint.DNSName)
			continue
		}
		// Get all virtual servers on ingress
		if endpoint.RecordType == "A" {
			for _, item := range endpoint.ProviderSpecific {
				if item.Name == providerSpecificPropertyName {
					var vsRecordsWithCluster map[string][]bigip.VSrecord
					err := json.Unmarshal([]byte(item.Value), &vsRecordsWithCluster)
					if err != nil {
						continue
					}
					vsRecords := filterVsrecords(vsRecordsWithCluster)
					for _, vsRecord := range vsRecords {
						mapVsRecords[endpoint.DNSName] = append(mapVsRecords[endpoint.DNSName], vsRecord)
						if isContain(allVsRecords, vsRecord, false, false) {
							continue
						}
						allVsRecords = append(allVsRecords, vsRecord)
					}
				}
			}
		}
	}
	err := p.updateServer(allVsRecords)
	if err != nil {
		log.Infof("Create/Update datacenter server error: %v", err)
		return err
	}
	for _, endpoint := range endpoints {
		if endpoint.RecordType == "A" && len(mapVsRecords[endpoint.DNSName]) > 0 {
			err := p.createPoolA(endpoint.DNSName, endpoint.RecordTTL)
			if err != nil {
				log.Errorf("create pool A err: %v", err)
			}
			err = p.addPoolAMember(mapVsRecords[endpoint.DNSName], endpoint.DNSName)
			if err != nil {
				log.Errorf("create pool A member err: %v", err)
			}
			err = p.createWideIP(endpoint)
			if err != nil {
				log.Errorf("create wide ip err: %v", err)
			}
		}
	}
	return nil
}

func (p F5BigipProvider) updateServer(vsRecords []bigip.VSrecord) error {
	var serverVRecords bigip.VSrecords
	serverName := os.Getenv("F5BIGIP_SERVER_NAME")
	serverVRecordsOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/server/~%s~%s/virtual-servers", partition, serverName),
		"",
		"application/json",
	}
	respServerVirtualServer, err := p.client.APICall(serverVRecordsOptions)
	if err != nil {
		log.Errorf("Get virtual Server error: %s", err)
	}
	if respServerVirtualServer != nil {
		err = json.Unmarshal(respServerVirtualServer, &serverVRecords)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	if len(serverVRecords.Items) > 0 {
		for _, vs := range serverVRecords.Items {
			if !isContain(vsRecords, vs, false, true) {
				vsRecords = append(vsRecords, vs)
			}
		}
	}
	server, err := p.client.GetGtmserver(serverName)
	if err != nil {
		return err
	}
	if server == nil {
		log.Errorf("Get pool Virtual Server error: %s", serverName)
	}
	updateServer := &bigip.Server{
		serverName,
		server.Datacenter,
		"",
		false,
		"",
		server.Addresses,
		vsRecords,
	}
	err = p.client.UpdateGtmserver(updateServer.Name, updateServer)
	if err != nil {
		return err
	}
	return nil
}

func (p F5BigipProvider) deleteServer(endpoint *endpoint.Endpoint) error {
	wideip, err := p.getWideIP(endpoint.DNSName)
	if err != nil {
		return err
	}
	for _, poolA := range wideip.PoolAs {
		var vsName []string
		poolAMembers, _, err := p.getPoolAMemberAndTtl(poolA)
		if err != nil {
			return err
		}
		for _, poolAMember := range poolAMembers {
			nameArr := strings.Split(poolAMember.Name, ":")
			if len(nameArr) == 2 {
				vsName = append(vsName, nameArr[1])
			}
		}

		if len(vsName) > 0 {
			err = p.deleteVirtualServers(vsName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p F5BigipProvider) deleteVirtualServers(vsName []string) error {
	var serverVRecords bigip.VSrecords
	var vsRecords []bigip.VSrecord
	serverName := os.Getenv("F5BIGIP_SERVER_NAME")
	serverVRecordsOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/server/~%s~%s/virtual-servers", partition, serverName),
		"",
		"application/json",
	}
	respServerVirtualServer, err := p.client.APICall(serverVRecordsOptions)
	if err != nil {
		log.Errorf("Get virtual Server error: %s", err)
	}
	if respServerVirtualServer != nil {
		err = json.Unmarshal(respServerVirtualServer, &serverVRecords)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	if len(serverVRecords.Items) > 0 {
		for _, vs := range serverVRecords.Items {
			if !isContainVsname(vsName, vs.Name) {
				vsRecords = append(vsRecords, vs)
			}
		}
	}
	log.Infof("vsRecords: %v", vsRecords)
	server, err := p.client.GetGtmserver(serverName)
	if err != nil {
		return err
	}
	if server == nil {
		log.Errorf("Get pool Virtual Server error: %s", serverName)
	}
	updateServer := &bigip.Server{
		serverName,
		server.Datacenter,
		"",
		false,
		"",
		server.Addresses,
		vsRecords,
	}
	err = p.client.UpdateGtmserver(updateServer.Name, updateServer)
	if err != nil {
		return err
	}
	return nil
}

func (p F5BigipProvider) createPoolA(poolAName string, ttl endpoint.TTL) error {
	createPoolA := PoolA{
		poolAName,
		"round-robin",
		partition,
		true,
		"any",
		"return-to-dns",
		"round-robin",
		ttl,
		"enabled",
	}
	reqBody, err := json.Marshal(createPoolA)
	if err != nil {
		return err
	}
	createPoolAOptions := &bigip.APIRequest{
		"post",
		"mgmt/tm/gtm/pool/a",
		string(reqBody),
		contentType,
	}
	_, err = p.client.APICall(createPoolAOptions)
	if err != nil {
		return err
	}
	return nil
}

func (p F5BigipProvider) deletePoolAMember(poolAMembers []PoolAMember, poolAName string) error {
	for _, poolAMember := range poolAMembers {
		deletePoolAMemberOptions := &bigip.APIRequest{
			"delete",
			fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s/members/~%s~%s", partition, poolAName, partition, poolAMember.Name),
			"",
			contentType,
		}
		_, err := p.client.APICall(deletePoolAMemberOptions)
		if err != nil {
			log.Errorf("delete pool member err: %v", err)
			continue
		}
	}
	return nil
}

func (p F5BigipProvider) addPoolAMember(vsRecords []bigip.VSrecord, poolAName string) error {
	serverName := os.Getenv("F5BIGIP_SERVER_NAME")
	for _, vsRecord := range vsRecords {
		poolAMember := PoolAMember{
			fmt.Sprintf("%s:%s", serverName, vsRecord.Name),
			partition,
			true,
			0,
			"default",
			1,
		}
		reqBody, err := json.Marshal(poolAMember)
		if err != nil {
			return err
		}
		addPoolAMemberOptions := &bigip.APIRequest{
			"post",
			fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s/members", partition, poolAName),
			string(reqBody),
			contentType,
		}
		_, err = p.client.APICall(addPoolAMemberOptions)
		if err != nil {
			log.Errorf("Add pool member err: %v", err)
			continue
		}
	}
	return nil
}

func (p F5BigipProvider) updatePoolA(endpoint *endpoint.Endpoint) error {
	var poolA PoolA
	poolOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s", partition, endpoint.DNSName),
		"",
		contentType,
	}
	respPool, err := p.client.APICall(poolOptions)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respPool, &poolA)
	if err != nil {
		return err
	}
	if poolA.Ttl != endpoint.RecordTTL {
		poolA.Ttl = endpoint.RecordTTL
		reqBody, err := json.Marshal(poolA)
		if err != nil {
			return err
		}
		updatePoolAOptions := &bigip.APIRequest{
			"put",
			fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s", partition, endpoint.DNSName),
			string(reqBody),
			contentType,
		}
		_, err = p.client.APICall(updatePoolAOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p F5BigipProvider) updatePoolAMember(vsRecords []bigip.VSrecord, poolAName string) error {
	var poolMember PoolAMembers
	poolMemberOptions := &bigip.APIRequest{
		"get",
		fmt.Sprintf("/mgmt/tm/gtm/pool/a/~%s~%s/members", partition, poolAName),
		"",
		"application/json",
	}
	respPoolMember, err := p.client.APICall(poolMemberOptions)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respPoolMember, &poolMember)
	if err != nil {
		return err
	}
	err = p.deletePoolAMember(poolMember.Items, poolAName)
	if err != nil {
		log.Errorf("Delete pool member err: %v", err)
	}
	err = p.addPoolAMember(vsRecords, poolAName)
	if err != nil {
		log.Errorf("Add pool member err: %v", err)
	}
	return nil
}

func (p F5BigipProvider) updateWideIP(endpoint *endpoint.Endpoint) error {
	wideIP, err := p.getWideIP(endpoint.DNSName)
	if err != nil {
		return err
	}
	description := ""
	for _, item := range endpoint.ProviderSpecific {
		if item.Name == providerSpecificPropertyName {
			description = formartDescription(item.Value)
		}
	}
	wideIP.Description = description
	wideIP.LastResortPool = "none"
	log.Infof("wideIP update: %v", wideIP)
	reqBody, err := json.Marshal(wideIP)
	if err != nil {
		return err
	}
	updateWideIpOptions := &bigip.APIRequest{
		"put",
		fmt.Sprintf("/mgmt/tm/gtm/wideip/a/~%s~%s", wideIP.Partition, wideIP.Name),
		string(reqBody),
		contentType,
	}
	_, err = p.client.APICall(updateWideIpOptions)
	if err != nil {
		return err
	}
	return nil
}

func (p F5BigipProvider) createWideIP(endpoint *endpoint.Endpoint) error {
	var wideIpPoolAs []WideIPPoolA
	createPoolA := WideIPPoolA{
		endpoint.DNSName,
		partition,
		endpoint.RecordTTL,
		1,
	}
	wideIpPoolAs = append(wideIpPoolAs, createPoolA)
	description := ""
	for _, item := range endpoint.ProviderSpecific {
		if item.Name == providerSpecificPropertyName {
			description = formartDescription(item.Value)
		}
	}
	createWideIp := WideIP{
		endpoint.DNSName,
		partition,
		description,
		true,
		"none",
		wideIpPoolAs,
	}
	reqBody, err := json.Marshal(createWideIp)
	if err != nil {
		return err
	}
	createWideIpOptions := &bigip.APIRequest{
		"post",
		"/mgmt/tm/gtm/wideip/a",
		string(reqBody),
		contentType,
	}
	_, err = p.client.APICall(createWideIpOptions)
	if err != nil {
		return err
	}
	return nil
}

func (p F5BigipProvider) delete(req *bigip.APIRequest, name string) error {
	_, err := p.client.APICall(req)
	if err != nil {
		if strings.Contains(err.Error(), "not Found") {
			log.Infof("%s already deleted", name)
		} else {
			return err
		}
	}
	return nil
}

func (p *F5BigipProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	if name != providerSpecificPropertyName {
		return true
	}
	return previous == strings.Replace(formartDescription(current), "\"", "", -1)
}

func createServer(client *bigip.BigIP) error {
	var serverAddresses []bigip.ServerAddresses
	var vsRecords []bigip.VSrecord
	dataCenterName := os.Getenv("F5BIGIP_DATACENTER_NAME")
	serverName := os.Getenv("F5BIGIP_SERVER_NAME")
	deviceIps := os.Getenv("F5BIGIP_DEVICE_IPS")
	server, err := client.GetGtmserver(serverName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(deviceIps), &serverAddresses)
	if err != nil {
		log.Errorf("DeviceIps format error: %s", err.Error())
		return err
	}
	if server == nil {
		createServer := &bigip.Server{
			serverName,
			dataCenterName,
			"",
			false,
			"",
			serverAddresses,
			vsRecords,
		}
		err = client.CreateGtmserver(createServer)
		if err != nil {
			return err
		}
		_, err = client.GetGtmserver(serverName)
		if err != nil {
			return err
		}
		return nil
	}
	updateServer := &bigip.Server{
		serverName,
		dataCenterName,
		"",
		false,
		"",
		serverAddresses,
		server.GTMVirtual_Server,
	}
	err = client.UpdateGtmserver(updateServer.Name, updateServer)
	if err != nil {
		return err
	}
	_, err = client.GetGtmserver(serverName)
	if err != nil {
		return err
	}
	return nil
}

func isContain(vsrecords []bigip.VSrecord, vsrecord bigip.VSrecord, onlyDestination bool, onlyName bool) bool {
	for _, item := range vsrecords {
		if onlyDestination {
			if vsrecord.Destination == item.Destination {
				return true
			}
		} else if onlyName {
			if vsrecord.Name == item.Name {
				return true
			}
		} else {
			if vsrecord.Destination == item.Destination && vsrecord.Name == item.Name {
				return true
			}
		}
	}
	return false
}

func isContainVsname(items []string, vsName string) bool {
	for _, eachItem := range items {
		if eachItem == vsName {
			return true
		}
	}
	return false
}
func filter(result []*endpoint.Endpoint, domainFilter endpoint.DomainFilter) []*endpoint.Endpoint {
	ret := make([]*endpoint.Endpoint, 0, len(result))
	for _, item := range result {
		if !domainFilter.Match(item.DNSName) {
			continue
		}
		ret = append(ret, item)
	}
	return ret
}

func filterVsrecords(vsRecordsWithCluster map[string][]bigip.VSrecord) []bigip.VSrecord {
	var vsRecords []bigip.VSrecord
	for _, value := range vsRecordsWithCluster {
		for _, vsRecord := range value {
			if isContain(vsRecords, vsRecord, true, false) {
				continue
			}
			if vsRecord.Name == "" {
				vsRecord.Name = strings.Replace(vsRecord.Destination, ":", "rancher", -1)
			}
			vsRecords = append(vsRecords, vsRecord)
		}
	}
	return vsRecords
}

func formartDescription(providerSpecifi string) string {
	var vsRecordsWithCluster map[string][]bigip.VSrecord
	err := json.Unmarshal([]byte(providerSpecifi), &vsRecordsWithCluster)
	if err != nil {
		return ""
	}
	for key, value := range vsRecordsWithCluster {
		var vsRecords []bigip.VSrecord
		for _, vsRecord := range value {
			if vsRecord.Name == "" {
				vsRecord.Name = strings.Replace(vsRecord.Destination, ":", "rancher", -1)
			}
			vsRecords = append(vsRecords, vsRecord)
		}
		vsRecordsWithCluster[key] = vsRecords
	}
	description, err := json.Marshal(vsRecordsWithCluster)
	if err != nil {
		return ""
	}
	return string(description)
}
