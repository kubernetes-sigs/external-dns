package ibclient

import (
	"fmt"
	"net"
	"regexp"
)

func (objMgr *ObjectManager) AllocateIP(
	netview string,
	cidr string,
	ipAddr string,
	isIPv6 bool,
	macOrDuid string,
	name string,
	comment string,
	eas EA,
	clients string,
	agentCircuitId string,
	agentRemoteId string,
	clientIdentifierPrependZero *bool,
	dhcpClientIdentifier string,
	disable bool,
	Options []*Dhcpoption,
	useOptions bool,
) (*FixedAddress, error) {

	if isIPv6 {
		if len(macOrDuid) == 0 {
			return nil, fmt.Errorf("the DUID field cannot be left empty")
		}
	} else {
		if len(macOrDuid) == 0 && clients == "" {
			macOrDuid = MACADDR_ZERO
		}
	}
	if ipAddr == "" && cidr != "" {
		if netview == "" {
			netview = "default"
		}
		ipAddr = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netview)
	}

	var clientsPointer *string
	if clients != "" {
		clientsPointer = &clients
	}
	var agentCircuitIdPointer *string
	if agentCircuitId != "" {
		agentCircuitIdPointer = &agentCircuitId
	}
	var agentRemoteIdPointer *string
	if agentRemoteId != "" {
		agentRemoteIdPointer = &agentRemoteId
	}
	var clientIdentifierPrependZeroPointer *bool
	if clientIdentifierPrependZero != nil {
		clientIdentifierPrependZeroPointer = clientIdentifierPrependZero
	}
	var dhcpClientIdentifierPointer *string
	if dhcpClientIdentifier != "" {
		dhcpClientIdentifierPointer = &dhcpClientIdentifier
	}
	fixedAddr := NewFixedAddress(
		netview, name, ipAddr, cidr, macOrDuid, clientsPointer, eas, "", isIPv6, comment, agentCircuitIdPointer, agentRemoteIdPointer, clientIdentifierPrependZeroPointer, dhcpClientIdentifierPointer, disable, Options, useOptions)
	ref, err := objMgr.connector.CreateObject(fixedAddr)
	if err != nil {
		return nil, err
	}

	fixedAddr.Ref = ref
	fixedAddr, err = objMgr.GetFixedAddressByRef(ref)

	return fixedAddr, err
}

func (objMgr *ObjectManager) GetFixedAddress(netview string, cidr string, ipAddr string, isIpv6 bool, macOrDuid string) (*FixedAddress, error) {
	var res []FixedAddress

	fixedAddr := NewEmptyFixedAddress(isIpv6)
	sf := map[string]string{
		"network_view": netview,
		"network":      cidr,
	}
	if isIpv6 {
		sf["ipv6addr"] = ipAddr
		if macOrDuid != "" {
			sf["duid"] = macOrDuid
		}
	} else {
		sf["ipv4addr"] = ipAddr
		if macOrDuid != "" {
			sf["mac"] = macOrDuid
		}
	}

	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(fixedAddr, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

func (objMgr *ObjectManager) GetFixedAddressByRef(ref string) (*FixedAddress, error) {
	r := regexp.MustCompile("^ipv6fixedaddress/.+")
	isIPv6 := r.MatchString(ref)

	fixedAddr := NewEmptyFixedAddress(isIPv6)
	err := objMgr.connector.GetObject(
		fixedAddr, ref, NewQueryParams(false, nil), &fixedAddr)
	return fixedAddr, err
}

func (objMgr *ObjectManager) UpdateFixedAddress(
	fixedAddrRef string,
	netview string,
	name string,
	cidr string,
	ipAddr string,
	matchClient string,
	macOrDuid string,
	comment string,
	eas EA,
	agentCircuitId string,
	agentRemoteId string,
	clientIdentifierPrependZero *bool,
	dhcpClientIdentifier string,
	disable bool,
	Options []*Dhcpoption,
	useOptions bool,
) (*FixedAddress, error) {

	r := regexp.MustCompile("^ipv6fixedaddress/.+")
	isIPv6 := r.MatchString(fixedAddrRef)
	if !isIPv6 {
		if !validateMatchClient(matchClient) {
			return nil, fmt.Errorf("wrong value for match_client passed %s \n ", matchClient)
		}
	}
	var clientsPointer *string
	if matchClient != "" {
		clientsPointer = &matchClient
	}
	var agentCircuitIdPointer *string
	if agentCircuitId != "" {
		agentCircuitIdPointer = &agentCircuitId
	}
	var agentRemoteIdPointer *string
	if agentRemoteId != "" {
		agentRemoteIdPointer = &agentRemoteId
	}
	var clientIdentifierPrependZeroPointer *bool
	if clientIdentifierPrependZero != nil {
		clientIdentifierPrependZeroPointer = clientIdentifierPrependZero
	}
	var dhcpClientIdentifierPointer *string
	if dhcpClientIdentifier != "" {
		dhcpClientIdentifierPointer = &dhcpClientIdentifier
	}
	updateFixedAddr := NewFixedAddress(
		"", name, "", "",
		macOrDuid, clientsPointer, eas, fixedAddrRef, isIPv6, comment, agentCircuitIdPointer, agentRemoteIdPointer, clientIdentifierPrependZeroPointer, dhcpClientIdentifierPointer, disable, Options, useOptions)

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
				updateFixedAddr.IPv6Address = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netview)
				updateFixedAddr.Cidr = cidr
				updateFixedAddr.NetviewName = netview
			} else {
				if ipAddress.To4() == nil {
					return nil, fmt.Errorf("CIDR value must be an IPv4 CIDR, not an IPv6 one")
				}
				updateFixedAddr.IPv4Address = fmt.Sprintf("func:nextavailableip:%s,%s", cidr, netview)
				updateFixedAddr.Cidr = cidr
				updateFixedAddr.NetviewName = netview
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
			updateFixedAddr.IPv6Address = ipAddr
			updateFixedAddr.Cidr = cidr
			updateFixedAddr.NetviewName = netview
		} else {
			if ipAddress.To4() == nil {
				return nil, fmt.Errorf("IP address must be an IPv4 address, not an IPv6 one")
			}
			updateFixedAddr.IPv4Address = ipAddr
			updateFixedAddr.Cidr = cidr
			updateFixedAddr.NetviewName = netview
		}
	}
	refResp, err := objMgr.connector.UpdateObject(updateFixedAddr, fixedAddrRef)
	if err != nil {
		return nil, err
	}
	updateFixedAddr.Ref = refResp
	updateFixedAddr, err = objMgr.GetFixedAddressByRef(refResp)
	if err != nil {
		return nil, err
	}
	return updateFixedAddr, nil
}

func (objMgr *ObjectManager) ReleaseIP(netview string, cidr string, ipAddr string, isIpv6 bool, macOrDuid string) (string, error) {
	fixAddress, _ := objMgr.GetFixedAddress(netview, cidr, ipAddr, isIpv6, macOrDuid)
	if fixAddress == nil {
		return "", nil
	}
	return objMgr.connector.DeleteObject(fixAddress.Ref)
}

func (objMgr *ObjectManager) DeleteFixedAddress(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) GetAllFixedAddress(queryParams *QueryParams, isIpv6 bool) ([]FixedAddress, error) {
	var res []FixedAddress
	fixedAddress := NewEmptyFixedAddress(isIpv6)
	err := objMgr.connector.GetObject(fixedAddress, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting fixed address: %s", err)
	}
	return res, nil
}
