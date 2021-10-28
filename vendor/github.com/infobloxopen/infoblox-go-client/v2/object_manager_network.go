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

	newRef, err := objMgr.connector.UpdateObject(nw, ref)
	if err != nil {
		return nil, err
	}

	nw.Ref = newRef
	return nw, nil
}

func (objMgr *ObjectManager) DeleteNetwork(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
