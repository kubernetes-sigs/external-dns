package ibclient

import (
	"encoding/json"
	"errors"
	"fmt"
<<<<<<< HEAD
)

// Compile-time interface checks
var _ IBObjectManager = new(ObjectManager)

type IBObjectManager interface {
	AllocateIP(netview string, cidr string, ipAddr string, isIPv6 bool, macOrDuid string, name string, comment string, eas EA) (*FixedAddress, error)
	AllocateNetwork(netview string, cidr string, isIPv6 bool, prefixLen uint, comment string, eas EA) (network *Network, err error)
	CreateARecord(netView string, dnsView string, name string, cidr string, ipAddr string, ttl uint32, useTTL bool, comment string, ea EA) (*RecordA, error)
	CreateAAAARecord(netView string, dnsView string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordAAAA, error)
	CreateZoneAuth(fqdn string, ea EA) (*ZoneAuth, error)
	CreateCNAMERecord(dnsview string, canonical string, recordname string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordCNAME, error)
	CreateDefaultNetviews(globalNetview string, localNetview string) (globalNetviewRef string, localNetviewRef string, err error)
	CreateEADefinition(eadef EADefinition) (*EADefinition, error)
	CreateHostRecord(enabledns bool, enabledhcp bool, recordName string, netview string, dnsview string, ipv4cidr string, ipv6cidr string, ipv4Addr string, ipv6Addr string, macAddr string, duid string, useTtl bool, ttl uint32, comment string, eas EA, aliases []string) (*HostRecord, error)
	CreateNetwork(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*Network, error)
	CreateNetworkContainer(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*NetworkContainer, error)
	CreateNetworkView(name string, comment string, setEas EA) (*NetworkView, error)
	CreatePTRRecord(networkView string, dnsView string, ptrdname string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordPTR, error)
	CreateTXTRecord(recordname string, text string, ttl uint, dnsview string) (*RecordTXT, error)
	CreateZoneDelegated(fqdn string, delegate_to []NameServer) (*ZoneDelegated, error)
	DeleteARecord(ref string) (string, error)
	DeleteAAAARecord(ref string) (string, error)
	DeleteZoneAuth(ref string) (string, error)
	DeleteCNAMERecord(ref string) (string, error)
	DeleteFixedAddress(ref string) (string, error)
	DeleteHostRecord(ref string) (string, error)
	DeleteNetwork(ref string) (string, error)
	DeleteNetworkContainer(ref string) (string, error)
	DeleteNetworkView(ref string) (string, error)
	DeletePTRRecord(ref string) (string, error)
	DeleteTXTRecord(ref string) (string, error)
	DeleteZoneDelegated(ref string) (string, error)
	GetARecordByRef(ref string) (*RecordA, error)
	GetARecord(dnsview string, recordName string, ipAddr string) (*RecordA, error)
	GetAAAARecord(dnsview string, recordName string, ipAddr string) (*RecordAAAA, error)
	GetAAAARecordByRef(ref string) (*RecordAAAA, error)
	GetCNAMERecord(dnsview string, canonical string, recordName string) (*RecordCNAME, error)
	GetCNAMERecordByRef(ref string) (*RecordCNAME, error)
	GetEADefinition(name string) (*EADefinition, error)
	GetFixedAddress(netview string, cidr string, ipAddr string, isIPv6 bool, macOrDuid string) (*FixedAddress, error)
	GetFixedAddressByRef(ref string) (*FixedAddress, error)
	GetHostRecord(netview string, dnsview string, recordName string, ipv4addr string, ipv6addr string) (*HostRecord, error)
	SearchHostRecordByAltId(internalId string, ref string, eaNameForInternalId string) (*HostRecord, error)
	GetHostRecordByRef(ref string) (*HostRecord, error)
	GetIpAddressFromHostRecord(host HostRecord) (string, error)
	GetNetwork(netview string, cidr string, isIPv6 bool, ea EA) (*Network, error)
	GetNetworkByRef(ref string) (*Network, error)
	GetNetworkContainer(netview string, cidr string, isIPv6 bool, eaSearch EA) (*NetworkContainer, error)
	GetNetworkContainerByRef(ref string) (*NetworkContainer, error)
	GetNetworkView(name string) (*NetworkView, error)
	GetNetworkViewByRef(ref string) (*NetworkView, error)
	GetPTRRecord(dnsview string, ptrdname string, recordName string, ipAddr string) (*RecordPTR, error)
	GetPTRRecordByRef(ref string) (*RecordPTR, error)
	GetZoneAuthByRef(ref string) (*ZoneAuth, error)
	GetZoneDelegated(fqdn string) (*ZoneDelegated, error)
	GetCapacityReport(name string) ([]CapacityReport, error)
	GetUpgradeStatus(statusType string) ([]UpgradeStatus, error)
	GetAllMembers() ([]Member, error)
	GetGridInfo() ([]Grid, error)
	GetGridLicense() ([]License, error)
	ReleaseIP(netview string, cidr string, ipAddr string, isIPv6 bool, macAddr string) (string, error)
	UpdateAAAARecord(ref string, netView string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordAAAA, error)
	UpdateCNAMERecord(ref string, canonical string, recordName string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordCNAME, error)
	UpdateFixedAddress(fixedAddrRef string, netview string, name string, cidr string, ipAddr string, matchclient string, macOrDuid string, comment string, eas EA) (*FixedAddress, error)
	UpdateHostRecord(hostRref string, enabledns bool, enabledhcp bool, name string, netview string, ipv4cidr string, ipv6cidr string, ipv4Addr string, ipv6Addr string, macAddress string, duid string, useTtl bool, ttl uint32, comment string, eas EA, aliases []string) (*HostRecord, error)
	UpdateNetwork(ref string, setEas EA, comment string) (*Network, error)
	UpdateNetworkContainer(ref string, setEas EA, comment string) (*NetworkContainer, error)
	UpdateNetworkView(ref string, name string, comment string, setEas EA) (*NetworkView, error)
	UpdatePTRRecord(ref string, netview string, ptrdname string, name string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordPTR, error)
	UpdateARecord(ref string, name string, ipAddr string, cidr string, netview string, ttl uint32, useTTL bool, comment string, eas EA) (*RecordA, error)
	UpdateZoneDelegated(ref string, delegate_to []NameServer) (*ZoneDelegated, error)
}

type ObjectManager struct {
	connector IBConnector
	cmpType   string
	tenantID  string
}

func NewObjectManager(connector IBConnector, cmpType string, tenantID string) IBObjectManager {
	objMgr := &ObjectManager{}

	objMgr.connector = connector
	objMgr.cmpType = cmpType
	objMgr.tenantID = tenantID

	return objMgr
}

// CreateMultiObject unmarshals the result into slice of maps
func (objMgr *ObjectManager) CreateMultiObject(req *MultiRequest) ([]map[string]interface{}, error) {

	conn := objMgr.connector.(*Connector)
	queryParams := NewQueryParams(false, nil)
	res, err := conn.makeRequest(CREATE, req, "", queryParams)

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(res, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUpgradeStatus returns the grid upgrade information
func (objMgr *ObjectManager) GetUpgradeStatus(statusType string) ([]UpgradeStatus, error) {
	var res []UpgradeStatus

	if statusType == "" {
		// TODO option may vary according to the WAPI version, need to
		// throw relevant  error.
		msg := fmt.Sprintf("Status type can not be nil")
		return res, errors.New(msg)
	}
	upgradestatus := NewUpgradeStatus(UpgradeStatus{})

	sf := map[string]string{
		"type": statusType,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(upgradestatus, "", queryParams, &res)

	return res, err
}

// GetAllMembers returns all members information
func (objMgr *ObjectManager) GetAllMembers() ([]Member, error) {
	var res []Member

	memberObj := NewMember(Member{})
	err := objMgr.connector.GetObject(
		memberObj, "", NewQueryParams(false, nil), &res)
	return res, err
}

// GetCapacityReport returns all capacity for members
func (objMgr *ObjectManager) GetCapacityReport(name string) ([]CapacityReport, error) {
	var res []CapacityReport

	capacityReport := NewCapcityReport(CapacityReport{})

	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(capacityReport, "", queryParams, &res)
	return res, err
}

// GetLicense returns the license details for member
func (objMgr *ObjectManager) GetLicense() ([]License, error) {
	var res []License

	licenseObj := NewLicense(License{})
	err := objMgr.connector.GetObject(
		licenseObj, "", NewQueryParams(false, nil), &res)
	return res, err
}

// GetLicense returns the license details for grid
func (objMgr *ObjectManager) GetGridLicense() ([]License, error) {
	var res []License

	licenseObj := NewGridLicense(License{})
	err := objMgr.connector.GetObject(
		licenseObj, "", NewQueryParams(false, nil), &res)
	return res, err
}

// GetGridInfo returns the details for grid
func (objMgr *ObjectManager) GetGridInfo() ([]Grid, error) {
	var res []Grid

	gridObj := NewGrid(Grid{})
	err := objMgr.connector.GetObject(
		gridObj, "", NewQueryParams(false, nil), &res)
	return res, err
}

// CreateZoneAuth creates zones and subs by passing fqdn
func (objMgr *ObjectManager) CreateZoneAuth(
	fqdn string,
	eas EA) (*ZoneAuth, error) {

	zoneAuth := NewZoneAuth(ZoneAuth{
		Fqdn: fqdn,
		Ea:   eas})

	ref, err := objMgr.connector.CreateObject(zoneAuth)
	zoneAuth.Ref = ref
	return zoneAuth, err
}

// Retreive a authortative zone by ref
func (objMgr *ObjectManager) GetZoneAuthByRef(ref string) (*ZoneAuth, error) {
	res := NewZoneAuth(ZoneAuth{})

	if ref == "" {
		return nil, fmt.Errorf("empty reference to an object is not allowed")
	}

	err := objMgr.connector.GetObject(
		res, ref, NewQueryParams(false, nil), res)
	return res, err
}

// DeleteZoneAuth deletes an auth zone
func (objMgr *ObjectManager) DeleteZoneAuth(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

// GetZoneAuth returns the authoritatives zones
func (objMgr *ObjectManager) GetZoneAuth() ([]ZoneAuth, error) {
	var res []ZoneAuth

	zoneAuth := NewZoneAuth(ZoneAuth{})
	err := objMgr.connector.GetObject(
		zoneAuth, "", NewQueryParams(false, nil), &res)

	return res, err
}

// GetZoneDelegated returns the delegated zone
func (objMgr *ObjectManager) GetZoneDelegated(fqdn string) (*ZoneDelegated, error) {
	if len(fqdn) == 0 {
		return nil, nil
	}
	var res []ZoneDelegated

	zoneDelegated := NewZoneDelegated(ZoneDelegated{})

	sf := map[string]string{
		"fqdn": fqdn,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(zoneDelegated, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

// CreateZoneDelegated creates delegated zone
func (objMgr *ObjectManager) CreateZoneDelegated(fqdn string, delegate_to []NameServer) (*ZoneDelegated, error) {
	zoneDelegated := NewZoneDelegated(ZoneDelegated{
		Fqdn:       fqdn,
		DelegateTo: delegate_to})

	ref, err := objMgr.connector.CreateObject(zoneDelegated)
	zoneDelegated.Ref = ref

	return zoneDelegated, err
}

// UpdateZoneDelegated updates delegated zone
func (objMgr *ObjectManager) UpdateZoneDelegated(ref string, delegate_to []NameServer) (*ZoneDelegated, error) {
	zoneDelegated := NewZoneDelegated(ZoneDelegated{
		Ref:        ref,
		DelegateTo: delegate_to})

	refResp, err := objMgr.connector.UpdateObject(zoneDelegated, ref)
	zoneDelegated.Ref = refResp
	return zoneDelegated, err
}

// DeleteZoneDelegated deletes delegated zone
func (objMgr *ObjectManager) DeleteZoneDelegated(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"regexp"
	"strings"
)

// Compile-time interface checks
var _ IBObjectManager = new(ObjectManager)

type IBObjectManager interface {
	GetDNSView(name string) (*View, error)
	AllocateIP(netview string, cidr string, ipAddr string, isIPv6 bool, macOrDuid string, name string, comment string, eas EA) (*FixedAddress, error)
	AllocateNetwork(netview string, cidr string, isIPv6 bool, prefixLen uint, comment string, eas EA) (network *Network, err error)
	AllocateNetworkContainer(netview string, cidr string, isIPv6 bool, prefixLen uint, comment string, eas EA) (netContainer *NetworkContainer, err error)
	CreateARecord(netView string, dnsView string, name string, cidr string, ipAddr string, ttl uint32, useTTL bool, comment string, ea EA) (*RecordA, error)
	CreateAAAARecord(netView string, dnsView string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordAAAA, error)
	CreateZoneAuth(fqdn string, ea EA) (*ZoneAuth, error)
	CreateCNAMERecord(dnsview string, canonical string, recordname string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordCNAME, error)
	CreateDefaultNetviews(globalNetview string, localNetview string) (globalNetviewRef string, localNetviewRef string, err error)
	CreateEADefinition(eadef EADefinition) (*EADefinition, error)
	CreateHostRecord(enabledns bool, enabledhcp bool, recordName string, netview string, dnsview string, ipv4cidr string, ipv6cidr string, ipv4Addr string, ipv6Addr string, macAddr string, duid string, useTtl bool, ttl uint32, comment string, eas EA, aliases []string) (*HostRecord, error)
	CreateMXRecord(dnsView string, fqdn string, mx string, preference uint32, ttl uint32, useTtl bool, comment string, eas EA) (*RecordMX, error)
	CreateNetwork(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*Network, error)
	CreateNetworkContainer(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*NetworkContainer, error)
	CreateNetworkView(name string, comment string, setEas EA) (*NetworkView, error)
	CreatePTRRecord(networkView string, dnsView string, ptrdname string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, eas EA) (*RecordPTR, error)
	CreateSRVRecord(dnsView string, name string, priority uint32, weight uint32, port uint32, target string, ttl uint32, useTtl bool, comment string, eas EA) (*RecordSRV, error)
	CreateTXTRecord(dnsView string, recordName string, text string, ttl uint32, useTtl bool, comment string, eas EA) (*RecordTXT, error)
	CreateZoneDelegated(fqdn string, delegate_to []NameServer) (*ZoneDelegated, error)
	DeleteARecord(ref string) (string, error)
	DeleteAAAARecord(ref string) (string, error)
	DeleteZoneAuth(ref string) (string, error)
	DeleteCNAMERecord(ref string) (string, error)
	DeleteFixedAddress(ref string) (string, error)
	DeleteHostRecord(ref string) (string, error)
	DeleteMXRecord(ref string) (string, error)
	DeleteNetwork(ref string) (string, error)
	DeleteNetworkContainer(ref string) (string, error)
	DeleteNetworkView(ref string) (string, error)
	DeletePTRRecord(ref string) (string, error)
	DeleteSRVRecord(ref string) (string, error)
	DeleteTXTRecord(ref string) (string, error)
	DeleteZoneDelegated(ref string) (string, error)
	GetARecordByRef(ref string) (*RecordA, error)
	GetARecord(dnsview string, recordName string, ipAddr string) (*RecordA, error)
	GetAAAARecord(dnsview string, recordName string, ipAddr string) (*RecordAAAA, error)
	GetAAAARecordByRef(ref string) (*RecordAAAA, error)
	GetCNAMERecord(dnsview string, canonical string, recordName string) (*RecordCNAME, error)
	GetCNAMERecordByRef(ref string) (*RecordCNAME, error)
	GetEADefinition(name string) (*EADefinition, error)
	GetFixedAddress(netview string, cidr string, ipAddr string, isIPv6 bool, macOrDuid string) (*FixedAddress, error)
	GetFixedAddressByRef(ref string) (*FixedAddress, error)
	GetHostRecord(netview string, dnsview string, recordName string, ipv4addr string, ipv6addr string) (*HostRecord, error)
	SearchHostRecordByAltId(internalId string, ref string, eaNameForInternalId string) (*HostRecord, error)
	GetHostRecordByRef(ref string) (*HostRecord, error)
	GetIpAddressFromHostRecord(host HostRecord) (string, error)
	GetMXRecord(dnsView string, fqdn string, mx string, preference uint32) (*RecordMX, error)
	GetMXRecordByRef(ref string) (*RecordMX, error)
	GetNetwork(netview string, cidr string, isIPv6 bool, ea EA) (*Network, error)
	GetNetworkByRef(ref string) (*Network, error)
	GetNetworkContainer(netview string, cidr string, isIPv6 bool, eaSearch EA) (*NetworkContainer, error)
	GetNetworkContainerByRef(ref string) (*NetworkContainer, error)
	GetNetworkView(name string) (*NetworkView, error)
	GetNetworkViewByRef(ref string) (*NetworkView, error)
	GetPTRRecord(dnsview string, ptrdname string, recordName string, ipAddr string) (*RecordPTR, error)
	GetPTRRecordByRef(ref string) (*RecordPTR, error)
	GetSRVRecord(dnsView string, name string, target string, port uint32) (*RecordSRV, error)
	GetSRVRecordByRef(ref string) (*RecordSRV, error)
	GetTXTRecord(dnsview string, name string) (*RecordTXT, error)
	GetTXTRecordByRef(ref string) (*RecordTXT, error)
	GetZoneAuthByRef(ref string) (*ZoneAuth, error)
	GetZoneDelegated(fqdn string) (*ZoneDelegated, error)
	GetCapacityReport(name string) ([]CapacityReport, error)
	GetUpgradeStatus(statusType string) ([]UpgradeStatus, error)
	GetAllMembers() ([]Member, error)
	GetGridInfo() ([]Grid, error)
	GetGridLicense() ([]License, error)
	SearchObjectByAltId(objType string, internalId string, ref string, eaNameForInternalId string) (interface{}, error)
	ReleaseIP(netview string, cidr string, ipAddr string, isIPv6 bool, macAddr string) (string, error)
	UpdateAAAARecord(ref string, netView string, recordName string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordAAAA, error)
	UpdateCNAMERecord(ref string, canonical string, recordName string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordCNAME, error)
	UpdateFixedAddress(fixedAddrRef string, netview string, name string, cidr string, ipAddr string, matchclient string, macOrDuid string, comment string, eas EA) (*FixedAddress, error)
	UpdateHostRecord(hostRref string, enabledns bool, enabledhcp bool, name string, netview string, dnsView string, ipv4cidr string, ipv6cidr string, ipv4Addr string, ipv6Addr string, macAddress string, duid string, useTtl bool, ttl uint32, comment string, eas EA, aliases []string) (*HostRecord, error)
	UpdateMXRecord(ref string, dnsView string, fqdn string, mx string, preference uint32, ttl uint32, useTtl bool, comment string, eas EA) (*RecordMX, error)
	UpdateNetwork(ref string, setEas EA, comment string) (*Network, error)
	UpdateNetworkContainer(ref string, setEas EA, comment string) (*NetworkContainer, error)
	UpdateNetworkView(ref string, name string, comment string, setEas EA) (*NetworkView, error)
	UpdatePTRRecord(ref string, netview string, ptrdname string, name string, cidr string, ipAddr string, useTtl bool, ttl uint32, comment string, setEas EA) (*RecordPTR, error)
	UpdateSRVRecord(ref string, name string, priority uint32, weight uint32, port uint32, target string, ttl uint32, useTtl bool, comment string, eas EA) (*RecordSRV, error)
	UpdateTXTRecord(ref string, recordName string, text string, ttl uint32, useTtl bool, comment string, eas EA) (*RecordTXT, error)
	UpdateARecord(ref string, name string, ipAddr string, cidr string, netview string, ttl uint32, useTTL bool, comment string, eas EA) (*RecordA, error)
	UpdateZoneDelegated(ref string, delegate_to []NameServer) (*ZoneDelegated, error)
	GetDnsMember(ref string) ([]Dns, error)
	UpdateDnsStatus(ref string, status bool) (Dns, error)
	GetDhcpMember(ref string) ([]Dhcp, error)
	UpdateDhcpStatus(ref string, status bool) (Dhcp, error)
}

const (
	ARecord               = "A"
	AaaaRecord            = "AAAA"
	CnameRecord           = "CNAME"
	MxRecord              = "MX"
	SrvRecord             = "SRV"
	TxtRecord             = "TXT"
	PtrRecord             = "PTR"
	HostRecordConst       = "Host"
	DnsViewConst          = "DNSView"
	ZoneAuthConst         = "ZoneAuth"
	NetworkViewConst      = "NetworkView"
	NetworkConst          = "Network"
	NetworkContainerConst = "NetworkContainer"
)

// Map of record type to its corresponding object
var getRecordTypeMap = map[string]func(ref string) IBObject{
	ARecord: func(ref string) IBObject {
		return NewEmptyRecordA()
	},
	AaaaRecord: func(ref string) IBObject {
		return NewEmptyRecordAAAA()
	},
	CnameRecord: func(ref string) IBObject {
		return NewEmptyRecordCNAME()
	},
	MxRecord: func(ref string) IBObject {
		return NewEmptyRecordMX()
	},
	SrvRecord: func(ref string) IBObject {
		return NewEmptyRecordSRV()
	},
	TxtRecord: func(ref string) IBObject {
		return NewEmptyRecordTXT()
	},
	PtrRecord: func(ref string) IBObject {
		return NewEmptyRecordPTR()
	},
	HostRecordConst: func(ref string) IBObject {
		return NewEmptyHostRecord()
	},
	DnsViewConst: func(ref string) IBObject {
		return NewEmptyDNSView()
	},
	ZoneAuthConst: func(ref string) IBObject {
		zone := &ZoneAuth{}
		zone.SetReturnFields(append(
			zone.ReturnFields(),
			"comment",
			"ns_group",
			"soa_default_ttl",
			"soa_expire",
			"soa_negative_ttl",
			"soa_refresh",
			"soa_retry",
			"view",
			"zone_format",
			"extattrs",
		))
		return zone
	},
	NetworkViewConst: func(ref string) IBObject {
		return NewEmptyNetworkView()
	},
	NetworkContainerConst: func(ref string) IBObject {
		return NewNetworkContainer("", "", false, "", nil)
	},
	NetworkConst: func(ref string) IBObject {
		r := regexp.MustCompile("^ipv6network\\/.+")
		isIPv6 := r.MatchString(ref)
		return NewNetwork("", "", isIPv6, "", nil)
	},
}

// Map returns the object with search fields with the given record type
var getObjectWithSearchFieldsMap = map[string]func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error){
	ARecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordA).Ref != "" {
			return res, nil
		}
		var recordAList []*RecordA
		err := objMgr.connector.GetObject(NewEmptyRecordA(), "", NewQueryParams(false, sf), &recordAList)
		if err == nil && len(recordAList) > 0 {
			res = recordAList[0]
		}
		return res, err
	},
	AaaaRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordAAAA).Ref != "" {
			return res, nil
		}
		var recordAaaList []*RecordAAAA
		err := objMgr.connector.GetObject(NewEmptyRecordAAAA(), "", NewQueryParams(false, sf), &recordAaaList)
		if err == nil && len(recordAaaList) > 0 {
			res = recordAaaList[0]
		}
		return res, err
	},
	CnameRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordCNAME).Ref != "" {
			return res, nil
		}
		var cNameList []*RecordCNAME
		err := objMgr.connector.GetObject(NewEmptyRecordCNAME(), "", NewQueryParams(false, sf), &cNameList)
		if err == nil && len(cNameList) > 0 {
			res = cNameList[0]
		}
		return res, err
	},
	MxRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordMX).Ref != "" {
			return res, nil
		}
		var mxList []*RecordMX
		err := objMgr.connector.GetObject(NewEmptyRecordMX(), "", NewQueryParams(false, sf), &mxList)
		if err == nil && len(mxList) > 0 {
			res = mxList[0]
		}
		return res, err

	},
	SrvRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordSRV).Ref != "" {
			return res, nil
		}
		var srvList []*RecordSRV
		err := objMgr.connector.GetObject(NewEmptyRecordSRV(), "", NewQueryParams(false, sf), &srvList)
		if err == nil && len(srvList) > 0 {
			res = srvList[0]
		}
		return res, err
	},
	TxtRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordTXT).Ref != "" {
			return res, nil
		}
		var txtList []*RecordTXT
		err := objMgr.connector.GetObject(NewEmptyRecordTXT(), "", NewQueryParams(false, sf), &txtList)
		if err == nil && len(txtList) > 0 {
			res = txtList[0]
		}
		return res, err
	},
	PtrRecord: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*RecordPTR).Ref != "" {
			return res, nil
		}
		var ptrList []*RecordPTR
		err := objMgr.connector.GetObject(NewEmptyRecordPTR(), "", NewQueryParams(false, sf), &ptrList)
		if err == nil && len(ptrList) > 0 {
			res = ptrList[0]
		}
		return res, err
	},
	HostRecordConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*HostRecord).Ref != "" {
			return res, nil
		}
		var hostRecordList []*HostRecord
		err := objMgr.connector.GetObject(NewEmptyHostRecord(), "", NewQueryParams(false, sf), &hostRecordList)
		if err == nil && len(hostRecordList) > 0 {
			res = hostRecordList[0]
		}
		return res, err
	},
	DnsViewConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*View).Ref != "" {
			return res, nil
		}
		var dnsViewList []*View
		err := objMgr.connector.GetObject(NewEmptyDNSView(), "", NewQueryParams(false, sf), &dnsViewList)
		if err == nil && len(dnsViewList) > 0 {
			res = dnsViewList[0]
		}
		return res, err
	},
	ZoneAuthConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*ZoneAuth).Ref != "" {
			return res, nil
		}
		zoneAuth := recordType.(*ZoneAuth)
		var zoneAuthList []*ZoneAuth
		err := objMgr.connector.GetObject(zoneAuth, "", NewQueryParams(false, sf), &zoneAuthList)
		if err == nil && len(zoneAuthList) > 0 {
			res = zoneAuthList[0]
		}
		return res, err
	},
	NetworkViewConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*NetworkView).Ref != "" {
			return res, nil
		}
		var networkViewList []*NetworkView
		err := objMgr.connector.GetObject(NewEmptyNetworkView(), "", NewQueryParams(false, sf), &networkViewList)
		if err == nil && len(networkViewList) > 0 {
			res = networkViewList[0]
		}
		return res, err
	},
	// TODO: Do we need to add netview string, cidr string, isIPv6 bool, ea EA to create network container
	NetworkContainerConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*NetworkContainer).Ref != "" {
			return res, nil
		}
		var networkContainerList []*NetworkContainer
		err := objMgr.connector.GetObject(NewNetworkContainer("", "", false, "", nil), "", NewQueryParams(false, sf), &networkContainerList)
		if err == nil && len(networkContainerList) > 0 {
			res = networkContainerList[0]
		}
		return res, err

	},
	//TODO: Do we need to add netview string, cidr string, isIPv6 bool, ea EA to create network
	NetworkConst: func(recordType IBObject, objMgr *ObjectManager, sf map[string]string) (interface{}, error) {
		var res interface{}
		if recordType.(*Network).Ref != "" {
			return res, nil
		}
		var networkList []*Network
		err := objMgr.connector.GetObject(NewNetwork("", "", false, "", nil), "", NewQueryParams(false, sf), &networkList)
		if err == nil && len(networkList) > 0 {
			res = networkList[0]
		}
		return res, err
	},
}

type ObjectManager struct {
	connector IBConnector
	cmpType   string
	tenantID  string
}

func NewObjectManager(connector IBConnector, cmpType string, tenantID string) IBObjectManager {
	objMgr := &ObjectManager{}

	objMgr.connector = connector
	objMgr.cmpType = cmpType
	objMgr.tenantID = tenantID

	return objMgr
}

// CreateMultiObject unmarshals the result into slice of maps
func (objMgr *ObjectManager) CreateMultiObject(req *MultiRequest) ([]map[string]interface{}, error) {

	conn := objMgr.connector.(*Connector)
	queryParams := NewQueryParams(false, nil)
	res, err := conn.makeRequest(CREATE, req, "", queryParams)

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(res, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUpgradeStatus returns the grid upgrade information
func (objMgr *ObjectManager) GetUpgradeStatus(statusType string) ([]UpgradeStatus, error) {
	var res []UpgradeStatus

	if statusType == "" {
		// TODO option may vary according to the WAPI version, need to
		// throw relevant  error.
		msg := fmt.Sprintf("Status type can not be nil")
		return res, errors.New(msg)
	}
	upgradestatus := NewUpgradeStatus(UpgradeStatus{})

	sf := map[string]string{
		"type": statusType,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(upgradestatus, "", queryParams, &res)

	return res, err
}

// GetAllMembers returns all members information
func (objMgr *ObjectManager) GetAllMembers() ([]Member, error) {
	var res []Member

	memberObj := NewMember(Member{})
	err := objMgr.connector.GetObject(
		memberObj, "", NewQueryParams(false, nil), &res,
	)
	return res, err
}

// GetCapacityReport returns all capacity for members
func (objMgr *ObjectManager) GetCapacityReport(name string) ([]CapacityReport, error) {
	var res []CapacityReport

	capacityReport := NewCapcityReport(CapacityReport{})

	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(capacityReport, "", queryParams, &res)
	return res, err
}

// GetLicense returns the license details for member
func (objMgr *ObjectManager) GetLicense() ([]License, error) {
	var res []License

	licenseObj := NewLicense(License{})
	err := objMgr.connector.GetObject(
		licenseObj, "", NewQueryParams(false, nil), &res,
	)
	return res, err
}

// GetLicense returns the license details for grid
func (objMgr *ObjectManager) GetGridLicense() ([]License, error) {
	var res []License

	licenseObj := NewGridLicense(License{})
	err := objMgr.connector.GetObject(
		licenseObj, "", NewQueryParams(false, nil), &res,
	)
	return res, err
}

// GetGridInfo returns the details for grid
func (objMgr *ObjectManager) GetGridInfo() ([]Grid, error) {
	var res []Grid

	gridObj := NewGrid(Grid{})
	err := objMgr.connector.GetObject(
		gridObj, "", NewQueryParams(false, nil), &res,
	)
	return res, err
}

// CreateZoneAuth creates zones and subs by passing fqdn
func (objMgr *ObjectManager) CreateZoneAuth(
	fqdn string,
	eas EA) (*ZoneAuth, error) {

	zoneAuth := NewZoneAuth(
		ZoneAuth{
			Fqdn: fqdn,
			Ea:   eas},
	)

	ref, err := objMgr.connector.CreateObject(zoneAuth)
	zoneAuth.Ref = ref
	return zoneAuth, err
}

// Retreive a authortative zone by ref
func (objMgr *ObjectManager) GetZoneAuthByRef(ref string) (*ZoneAuth, error) {
	res := NewZoneAuth(ZoneAuth{})

	if ref == "" {
		return nil, fmt.Errorf("empty reference to an object is not allowed")
	}

	err := objMgr.connector.GetObject(
		res, ref, NewQueryParams(false, nil), res,
	)
	return res, err
}

// DeleteZoneAuth deletes an auth zone
func (objMgr *ObjectManager) DeleteZoneAuth(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

// GetZoneAuth returns the authoritatives zones
func (objMgr *ObjectManager) GetZoneAuth() ([]ZoneAuth, error) {
	var res []ZoneAuth

	zoneAuth := NewZoneAuth(ZoneAuth{})
	err := objMgr.connector.GetObject(
		zoneAuth, "", NewQueryParams(false, nil), &res,
	)

	return res, err
}

// GetZoneDelegated returns the delegated zone
func (objMgr *ObjectManager) GetZoneDelegated(fqdn string) (*ZoneDelegated, error) {
	if len(fqdn) == 0 {
		return nil, nil
	}
	var res []ZoneDelegated

	zoneDelegated := NewZoneDelegated(ZoneDelegated{})

	sf := map[string]string{
		"fqdn": fqdn,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(zoneDelegated, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

// CreateZoneDelegated creates delegated zone
func (objMgr *ObjectManager) CreateZoneDelegated(fqdn string, delegate_to []NameServer) (*ZoneDelegated, error) {
	zoneDelegated := NewZoneDelegated(
		ZoneDelegated{
			Fqdn:       fqdn,
			DelegateTo: delegate_to},
	)

	ref, err := objMgr.connector.CreateObject(zoneDelegated)
	zoneDelegated.Ref = ref

	return zoneDelegated, err
}

// UpdateZoneDelegated updates delegated zone
func (objMgr *ObjectManager) UpdateZoneDelegated(ref string, delegate_to []NameServer) (*ZoneDelegated, error) {
	zoneDelegated := NewZoneDelegated(
		ZoneDelegated{
			Ref:        ref,
			DelegateTo: delegate_to},
	)

	refResp, err := objMgr.connector.UpdateObject(zoneDelegated, ref)
	zoneDelegated.Ref = refResp
	return zoneDelegated, err
}

// DeleteZoneDelegated deletes delegated zone
func (objMgr *ObjectManager) DeleteZoneDelegated(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

// SearchObjectByAltId is a generic function to search object by alternate id
func (objMgr *ObjectManager) SearchObjectByAltId(
	objType string, ref string, internalId string, eaNameForInternalId string) (interface{}, error) {
	var (
		err        error
		recordType IBObject
		res        interface{}
	)
	val, ok := getRecordTypeMap[objType]
	if !ok {
		return nil, fmt.Errorf("unknown record type")
	}
	recordType = val(ref)

	if ref != "" {
		// Fetching object by reference
		if err := objMgr.connector.GetObject(recordType, ref, NewQueryParams(false, nil), &res); err != nil {
			if _, ok := err.(*NotFoundError); !ok {
				return nil, err
			}
		}
		success, err := validateObjByInternalId(res, internalId, eaNameForInternalId)
		if err != nil {
			return nil, err
		} else if success {
			return res, nil
		}
	}

	sf := map[string]string{
		fmt.Sprintf("*%s", eaNameForInternalId): internalId,
	}

	// Fetch the object by search fields
	getObjectWithSearchFields, ok := getObjectWithSearchFieldsMap[objType]
	if !ok {
		return nil, fmt.Errorf("unknown record type")
	}
	res, err = getObjectWithSearchFields(recordType, objMgr, sf)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, NewNotFoundError("record not found")
	}

	return &res, nil
}

// validateObjByInternalId validates the object by comparing the given internal with the object's internal id
func validateObjByInternalId(res interface{}, internalId, eaNameForInternalId string) (bool, error) {
	var success bool
	if res == nil {
		return success, nil
	} else if strings.TrimSpace(internalId) == "" {
		// return object if internal id is empty
		success = true
		return success, nil
	}
	byteObj, err := json.Marshal(res)
	if err != nil {
		return success, fmt.Errorf("error marshaling JSON: %v", err)
	}
	obj := make(map[string]interface{})
	err = json.Unmarshal(byteObj, &obj)
	if err != nil {
		return success, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	extAttrs, err := getInterfaceValueFromMap(obj, "extattrs")
	if err == nil {
		resInternalId, err := getInterfaceValueFromMap(extAttrs, eaNameForInternalId)
		if err == nil && resInternalId["value"] != nil && resInternalId["value"].(string) == internalId {
			// return object if object's internal id matches with the given internal id
			success = true
			return success, nil
		}
		return success, err
	}
	return success, err
}

// getInterfaceValueFromMap returns the value, after converting it into a map[string]interface{}, of the given key from the map
func getInterfaceValueFromMap(m map[string]interface{}, key string) (map[string]interface{}, error) {
	if val, ok := m[key]; ok && val != nil {
		res := val.(map[string]interface{})
		return res, nil
	}
	return nil, fmt.Errorf("key %s not found in map", key)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
