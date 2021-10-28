package ibclient

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// TODO check if the respective zone exists before creation of the record
func (objMgr *ObjectManager) CreatePTRRecord(
	networkView string,
	dnsView string,
	ptrdname string,
	recordName string,
	cidr string,
	ipAddr string,
	useTtl bool,
	ttl uint32,
	comment string,
	eas EA) (*RecordPTR, error) {

	if ptrdname == "" {
		return nil, fmt.Errorf("ptrdname is a required field to create a PTR record")
	}
	recordPTR := NewRecordPTR(dnsView, ptrdname, useTtl, ttl, comment, eas)

	if recordName != "" {
		recordPTR.Name = recordName
	} else if ipAddr == "" && cidr != "" {
		if networkView == "" {
			networkView = "default"
		}
		ipAddress, net, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
		if ipAddress.To4() != nil {
			if net.String() != cidr {
				return nil, fmt.Errorf("%s is an invalid CIDR. Note: leading zeros should be removed if exists", cidr)
			}
			recordPTR.Ipv4Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, networkView)
		} else {
			recordPTR.Ipv6Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, networkView)
		}
	} else if ipAddr != "" {
		ipAddress := net.ParseIP(ipAddr)
		if ipAddress == nil {
			return nil, fmt.Errorf("%s is an invalid IP address", ipAddr)
		}
		if ipAddress.To4() != nil {
			recordPTR.Ipv4Addr = ipAddr
		} else {
			recordPTR.Ipv6Addr = ipAddr
		}
	} else {
		return nil, fmt.Errorf("CIDR and network view are required to allocate a next available IP address\n" +
			"IP address is required to create PTR record in reverse mapping zone\n" +
			"record name is required to create a record in forwarrd mapping zone")
	}
	ref, err := objMgr.connector.CreateObject(recordPTR)
	if err != nil {
		return nil, err
	}
	recordPTR, err = objMgr.GetPTRRecordByRef(ref)
	return recordPTR, err
}

func (objMgr *ObjectManager) GetPTRRecord(dnsview string, ptrdname string, recordName string, ipAddr string) (*RecordPTR, error) {
	var res []RecordPTR
	recordPtr := NewEmptyRecordPTR()
	sf := map[string]string{
		"view":     dnsview,
		"ptrdname": ptrdname,
	}
	cleanName := strings.TrimSpace(recordName)
	if ipAddr != "" {
		ipAddress := net.ParseIP(ipAddr)
		if ipAddress == nil {
			return nil, fmt.Errorf("%s is an invalid IP address", ipAddr)
		}
		if ipAddress.To4() != nil {
			sf["ipv4addr"] = ipAddr
		} else {
			sf["ipv6addr"] = ipAddr
		}
	} else if cleanName != "" {
		sf["name"] = cleanName
	} else {
		return nil, fmt.Errorf("record name or IP Address of the record has to be passed to get a unique record")
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordPtr, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"PTR record with name/IP '%v' and ptrdname '%s' in DNS view '%s' is not found",
				[]string{recordName, ipAddr}, ptrdname, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) GetPTRRecordByRef(ref string) (*RecordPTR, error) {
	recordPTR := NewEmptyRecordPTR()
	err := objMgr.connector.GetObject(
		recordPTR, ref, NewQueryParams(false, nil), &recordPTR)
	return recordPTR, err
}

func (objMgr *ObjectManager) DeletePTRRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) UpdatePTRRecord(
	ref string,
	netview string,
	ptrdname string,
	name string,
	cidr string,
	ipAddr string,
	useTtl bool,
	ttl uint32,
	comment string,
	setEas EA) (*RecordPTR, error) {

	recordPTR := NewRecordPTR("", ptrdname, useTtl, ttl, comment, setEas)
	recordPTR.Ref = ref
	recordPTR.Name = name
	isIPv6, _ := regexp.MatchString(`^record:ptr/.+.ip6.arpa/.+`, ref)

	if name == "" {
		if ipAddr == "" {
			if cidr != "" {
				ipAddress, _, err := net.ParseCIDR(cidr)
				if err != nil {
					return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
				}
				if netview == "" {
					netview = "default"
				}
				if isIPv6 {
					if ipAddress.To4() != nil || ipAddress.To16() == nil {
						return nil, fmt.Errorf("CIDR value must be an IPv6 CIDR, not an IPv4 one")
					}
					recordPTR.Ipv6Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netview)
				} else {
					if ipAddress.To4() == nil {
						return nil, fmt.Errorf("CIDR value must be an IPv4 CIDR, not an IPv6 one")
					}
					recordPTR.Ipv4Addr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netview)
				}
			}
		} else {
			ipAddress := net.ParseIP(ipAddr)
			if ipAddress == nil {
				return nil, fmt.Errorf("IP address for the record is not valid")
			}
			if isIPv6 {
				if ipAddress.To4() != nil || ipAddress.To16() == nil {
					return nil, fmt.Errorf("IP address must be an IPv6 address, not an IPv4 one")
				}
				recordPTR.Ipv6Addr = ipAddr
			} else {
				if ipAddress.To4() == nil {
					return nil, fmt.Errorf("IP address must be an IPv4 address, not an IPv6 one")
				}
				recordPTR.Ipv4Addr = ipAddr
			}
		}
	}
	reference, err := objMgr.connector.UpdateObject(recordPTR, ref)
	if err != nil {
		return nil, err
	}

	recordPTR, err = objMgr.GetPTRRecordByRef(reference)
	return recordPTR, err
}
