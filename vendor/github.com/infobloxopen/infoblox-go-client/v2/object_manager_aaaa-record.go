package ibclient

import (
	"fmt"
	"net"
	"strings"
)

func (objMgr *ObjectManager) CreateAAAARecord(
	netView string,
	dnsView string,
	recordName string,
	cidr string,
	ipAddr string,
	useTtl bool,
	ttl uint32,
	comment string,
	eas EA) (*RecordAAAA, error) {

	cleanName := strings.ToLower(strings.TrimSpace(recordName))
	if cleanName == "" || cleanName != recordName {
		return nil, fmt.Errorf(
			"'name' argument is expected to be non-empty and it must NOT contain leading/trailing spaces")
	}
	recordAAAA := NewRecordAAAA(dnsView, recordName, "", useTtl, ttl, comment, eas, "")

	if ipAddr == "" {
		if cidr == "" {
			return nil, fmt.Errorf("CIDR must not be empty")
		}
		ipAddress, _, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
		}
		if ipAddress.To4() != nil || ipAddress.To16() == nil {
			return nil, fmt.Errorf("CIDR value must be an IPv6 CIDR, not an IPv4 one")
		}
		if netView == "" {
			netView = "default"
		}
		recordAAAA.Ipv6Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netView)
	} else {
		ipAddress := net.ParseIP(ipAddr)
		if ipAddress == nil {
			return nil, fmt.Errorf("IP address for the record is not valid")
		}
		if ipAddress.To4() != nil || ipAddress.To16() == nil {
			return nil, fmt.Errorf("IP address must be an IPv6 address, not an IPv4 one")
		}
		recordAAAA.Ipv6Addr = ipAddr
	}
	ref, err := objMgr.connector.CreateObject(recordAAAA)
	if err != nil {
		return nil, err
	}
	recordAAAA, err = objMgr.GetAAAARecordByRef(ref)
	if err != nil {
		return nil, err
	}
	return recordAAAA, nil
}

func (objMgr *ObjectManager) GetAAAARecord(dnsview string, recordName string, ipAddr string) (*RecordAAAA, error) {
	var res []RecordAAAA
	recordAAAA := NewEmptyRecordAAAA()
	if dnsview == "" || recordName == "" || ipAddr == "" {
		return nil, fmt.Errorf("DNS view, IPv6 address and record name of the record are required to retreive a unique AAAA record")
	}
	sf := map[string]string{
		"view":     dnsview,
		"name":     recordName,
		"ipv6addr": ipAddr,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordAAAA, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"AAAA record with name '%s' and IPv6 address '%s' in DNS view '%s' is not found",
				recordName, ipAddr, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) GetAAAARecordByRef(ref string) (*RecordAAAA, error) {
	recordAAAA := NewEmptyRecordAAAA()
	err := objMgr.connector.GetObject(
		recordAAAA, ref, NewQueryParams(false, nil), &recordAAAA)
	return recordAAAA, err
}

func (objMgr *ObjectManager) DeleteAAAARecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) UpdateAAAARecord(
	ref string,
	netView string,
	recordName string,
	ipAddr string,
	cidr string,
	useTtl bool,
	ttl uint32,
	comment string,
	setEas EA) (*RecordAAAA, error) {

	cleanName := strings.ToLower(strings.TrimSpace(recordName))
	if cleanName == "" || cleanName != recordName {
		return nil, fmt.Errorf(
			"'name' argument is expected to be non-empty and it must NOT contain leading/trailing spaces")
	}

	rec, err := objMgr.GetAAAARecordByRef(ref)
	if err != nil {
		return nil, err
	}
	newIpAddr := rec.Ipv6Addr
	if ipAddr == "" {
		if cidr != "" {
			ipAddress, _, err := net.ParseCIDR(cidr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
			}
			if ipAddress.To4() != nil || ipAddress.To16() == nil {
				return nil, fmt.Errorf("CIDR value must be an IPv6 CIDR, not an IPv4 one")
			}
			if netView == "" {
				netView = "default"
			}
			newIpAddr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netView)
		}
	} else {
		ipAddress := net.ParseIP(ipAddr)
		if ipAddress == nil {
			return nil, fmt.Errorf("IP address for the record is not valid")
		}
		if ipAddress.To4() != nil || ipAddress.To16() == nil {
			return nil, fmt.Errorf("IP address must be an IPv6 address, not an IPv4 one")
		}
		newIpAddr = ipAddr
	}
	recordAAAA := NewRecordAAAA("", recordName, newIpAddr, useTtl, ttl, comment, setEas, ref)
	reference, err := objMgr.connector.UpdateObject(recordAAAA, ref)
	if err != nil {
		return nil, err
	}

	recordAAAA, err = objMgr.GetAAAARecordByRef(reference)
	if err != nil {
		return nil, err
	}
	return recordAAAA, nil
}
