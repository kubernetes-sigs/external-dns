package ibclient

import (
	"fmt"
	"regexp"
)

func (objMgr *ObjectManager) CreateNetwork(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*Network, error) {
	network := NewNetwork(netview, cidr, isIPv6, comment, eas)

	ref, err := objMgr.connector.CreateObject(network)
	if err != nil {
		return nil, err
	}
	network.Ref = ref

	return network, err
}

func (objMgr *ObjectManager) AllocateNetwork(
	netview string,
	cidr string,
	isIPv6 bool,
	prefixLen uint,
	comment string,
	eas EA) (network *Network, err error) {

	network = nil
	cidr = fmt.Sprintf("func:nextavailablenetwork:%s,%s,%d", cidr, netview, prefixLen)
	networkReq := NewNetwork(netview, cidr, isIPv6, comment, eas)

	ref, err := objMgr.connector.CreateObject(networkReq)
	if err == nil {
		if isIPv6 {
			network, err = BuildIPv6NetworkFromRef(ref)
		} else {
			network, err = BuildNetworkFromRef(ref)
		}
	}

	return
}

func (objMgr *ObjectManager) AllocateNextAvailableIp(
	name string,
	objectType string,
	objectParams map[string]string,
	params map[string][]string,
	useEaInheritance bool,
	ea EA,
	comment string,
	disable bool,
	n *int, ipAddrType string,
	enableDns bool, enableDhcp bool,
	macAddr string, duid string,
	networkView string, dnsView string,
	useTtl bool, ttl uint32, aliases []string) (interface{}, error) {

	networkIp := NewIpNextAvailable(name, objectType, objectParams, params, useEaInheritance, ea, comment, disable, n, ipAddrType,
		enableDns, enableDhcp, macAddr, duid, networkView, dnsView, useTtl, ttl, aliases)

	ref, err := objMgr.connector.CreateObject(networkIp)
	if err != nil {
		return nil, err
	}

	switch objectType {
	case "record:a":
		return objMgr.GetARecordByRef(ref)
	case "record:aaaa":
		return objMgr.GetAAAARecordByRef(ref)
	case "record:host":
		return objMgr.GetHostRecordByRef(ref)
	}

	return nil, err
}

func (objMgr *ObjectManager) AllocateNetworkByEA(
	netview string, isIPv6 bool, comment string, eas EA, eaMap map[string]string, prefixLen uint, object string) (network *Network, err error) {

	var (
		containerObject string
		objectType      string
	)

	objectType = getNetworkObjectType(isIPv6, "network", "ipv6network")

	if object == "network" {
		containerObject = getNetworkObjectType(isIPv6, "network", "ipv6network")
	} else {
		containerObject = getNetworkObjectType(isIPv6, "networkcontainer", "ipv6networkcontainer")
	}

	nextAvailableNetworkInfo := NetworkContainerNextAvailableInfo{
		Function:     "next_available_network",
		ResultField:  "networks",
		Object:       containerObject,
		ObjectParams: eaMap,
		Params:       map[string]uint{"cidr": prefixLen},
	}

	nextAvailableNetwork := NetworkContainerNextAvailable{
		Network:     &nextAvailableNetworkInfo,
		objectType:  objectType,
		Comment:     comment,
		Ea:          eas,
		NetviewName: netview,
	}

	ref, err := objMgr.connector.CreateObject(&nextAvailableNetwork)
	if err == nil {
		if isIPv6 {
			network, err = BuildIPv6NetworkFromRef(ref)
		} else {
			network, err = BuildNetworkFromRef(ref)
		}
	}
	return
}

func (objMgr *ObjectManager) GetNetwork(netview string, cidr string, isIPv6 bool, ea EA) (*Network, error) {
	if netview != "" && cidr != "" {
		var res []Network

		network := NewNetwork(netview, cidr, isIPv6, "", ea)

		network.Cidr = cidr

		if ea != nil && len(ea) > 0 {
			network.eaSearch = EASearch(ea)
		}

		sf := map[string]string{
			"network_view": netview,
			"network":      cidr,
		}
		queryParams := NewQueryParams(false, sf)
		err := objMgr.connector.GetObject(network, "", queryParams, &res)

		if err != nil {
			return nil, err
		} else if res == nil || len(res) == 0 {
			return nil, NewNotFoundError(
				fmt.Sprintf(
					"Network with cidr: %s in network view: %s is not found.",
					cidr, netview))
		}

		return &res[0], nil
	} else {
		err := fmt.Errorf("both network view and cidr values are required")
		return nil, err
	}
}

func (objMgr *ObjectManager) GetNetworkByRef(ref string) (*Network, error) {
	r := regexp.MustCompile("^ipv6network\\/.+")
	isIPv6 := r.MatchString(ref)

	network := NewNetwork("", "", isIPv6, "", nil)
	err := objMgr.connector.GetObject(network, ref, NewQueryParams(false, nil), network)
	return network, err
}

// UpdateNetwork updates comment and EA parameters.
// EAs which exist will be updated,
// those which do exist but not in setEas map, will be deleted,
// EAs which do not exist will be created as new.
func (objMgr *ObjectManager) UpdateNetwork(
	ref string,
	setEas EA,
	comment string) (*Network, error) {

	r := regexp.MustCompile("^ipv6network\\/.+")
	isIPv6 := r.MatchString(ref)

	nw := NewNetwork("", "", isIPv6, "", nil)
	err := objMgr.connector.GetObject(
		nw, ref, NewQueryParams(false, nil), nw)

	if err != nil {
		return nil, err
	}

	nw.Ea = setEas
	nw.Comment = comment

	// Network view is not allowed to be updated,
	// thus making its name empty (will not appear among data which we update).
	netViewSaved := nw.NetviewName
	nw.NetviewName = ""

	newRef, err := objMgr.connector.UpdateObject(nw, ref)
	if err != nil {
		return nil, err
	}

	nw.Ref = newRef
	nw.NetviewName = netViewSaved

	return nw, nil
}

func (objMgr *ObjectManager) DeleteNetwork(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
