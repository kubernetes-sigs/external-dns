package ibclient

import (
	"fmt"
	"net"
	"strings"
)

func (objMgr *ObjectManager) CreateARecord(
	netView string,
	dnsView string,
	name string,
	cidr string,
	ipAddr string,
	ttl uint32,
	useTTL bool,
	comment string,
	eas EA) (*RecordA, error) {

	cleanName := strings.TrimSpace(name)
	if cleanName == "" || cleanName != name {
		return nil, fmt.Errorf(
			"'name' argument is expected to be non-empty and it must NOT contain leading/trailing spaces")
	}

	recordA := NewRecordA(dnsView, "", name, "", ttl, useTTL, comment, eas, "")

	if ipAddr == "" {
		if cidr == "" {
			return nil, fmt.Errorf("CIDR must not be empty")
		}
		ip, _, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
		}
		if ip.To4() == nil {
			return nil, fmt.Errorf("CIDR value must be an IPv4 CIDR, not an IPv6 one")
		}
		if netView == "" {
			recordA.Ipv4Addr = fmt.Sprintf("func:nextavailableip:%s", cidr)
		} else {
			recordA.Ipv4Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netView)
		}
	} else {
		ip := net.ParseIP(ipAddr)
		if ip == nil {
			return nil, fmt.Errorf("'IP address for the record is not valid")
		}
		if ip.To4() == nil {
			return nil, fmt.Errorf("IP address must be an IPv4 address, not an IPv6 one")
		}
		recordA.Ipv4Addr = ipAddr
	}

	ref, err := objMgr.connector.CreateObject(recordA)
	if err != nil {
		return nil, err
	}

	newRec, err := objMgr.GetARecordByRef(ref)
	if err != nil {
		return nil, err
	}

	return newRec, nil
}

func (objMgr *ObjectManager) UpdateARecord(
	ref string,
	name string,
	ipAddr string,
	cidr string,
	netView string,
	ttl uint32,
	useTTL bool,
	comment string,
	eas EA) (*RecordA, error) {

	cleanName := strings.ToLower(strings.TrimSpace(name))
	if cleanName == "" || cleanName != name {
		return nil, fmt.Errorf(
			"'name' argument is expected to be non-empty and it must NOT contain leading/trailing spaces")
	}

	rec, err := objMgr.GetARecordByRef(ref)
	if err != nil {
		return nil, err
	}
	newIpAddr := rec.Ipv4Addr
	if ipAddr == "" {
		if cidr != "" {
			ip, _, err := net.ParseCIDR(cidr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
			}
			if ip.To4() == nil {
				return nil, fmt.Errorf("CIDR value must be an IPv4 CIDR, not an IPv6 one")
			}
			if netView == "" {
				newIpAddr = fmt.Sprintf("func:nextavailableip:%s", cidr)
			} else {
				newIpAddr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netView)
			}
		}
		// else: leaving ipv4addr field untouched
	} else {
		ip := net.ParseIP(ipAddr)
		if ip == nil {
			return nil, fmt.Errorf("'IP address for the record is not valid")
		}
		if ip.To4() == nil {
			return nil, fmt.Errorf("IP address must be an IPv4 address, not an IPv6 one")
		}
		newIpAddr = ipAddr
	}
	rec = NewRecordA(
		"", "", name, newIpAddr, ttl, useTTL, comment, eas, ref)
	ref, err = objMgr.connector.UpdateObject(rec, ref)
	if err != nil {
		return nil, err
	}
	rec.Ref = ref

	rec, err = objMgr.GetARecordByRef(ref)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

func (objMgr *ObjectManager) GetARecord(dnsview string, recordName string, ipAddr string) (*RecordA, error) {
	var res []RecordA
	recordA := NewEmptyRecordA()
	if dnsview == "" || recordName == "" || ipAddr == "" {
		return nil, fmt.Errorf("DNS view, IPv4 address and record name of the record are required to retreive a unique A record")
	}
	sf := map[string]string{
		"view":     dnsview,
		"name":     recordName,
		"ipv4addr": ipAddr,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordA, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"A record with name '%s' and IPv4 address '%s' in DNS view '%s' is not found",
				recordName, ipAddr, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) GetARecordByRef(ref string) (*RecordA, error) {
	recordA := NewEmptyRecordA()
	err := objMgr.connector.GetObject(
		recordA, ref, NewQueryParams(false, nil), &recordA)
	return recordA, err
}

func (objMgr *ObjectManager) DeleteARecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
