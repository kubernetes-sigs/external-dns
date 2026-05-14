package ibclient

import (
	"fmt"
	"regexp"
)

func (objMgr *ObjectManager) CreateNetworkContainer(netview string, cidr string, isIPv6 bool, comment string, eas EA) (*NetworkContainer, error) {
	container := NewNetworkContainer(netview, cidr, isIPv6, comment, eas)

	ref, err := objMgr.connector.CreateObject(container)
	if err != nil {
		return nil, err
	}

	container.Ref = ref
	return container, nil
}

// TODO normalize IPv4 and IPv6 addresses
func (objMgr *ObjectManager) GetNetworkContainer(netview string, cidr string, isIPv6 bool, eaSearch EA) (*NetworkContainer, error) {
	var res []NetworkContainer

	nc := NewNetworkContainer(netview, cidr, isIPv6, "", nil)
	nc.eaSearch = EASearch(eaSearch)
	sf := map[string]string{
		"network_view": netview,
		"network":      cidr,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(nc, "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError("network container not found")
	}

	return &res[0], nil
}

func (objMgr *ObjectManager) GetNetworkContainerByRef(ref string) (*NetworkContainer, error) {
	nc := NewNetworkContainer("", "", false, "", nil)

	err := objMgr.connector.GetObject(
		nc, ref, NewQueryParams(false, nil), nc)
	if err != nil {
		return nil, err
	}

	return nc, nil
}

func (objMgr *ObjectManager) UpdateNetworkContainer(
	ref string,
	setEas EA,
	comment string) (*NetworkContainer, error) {

	nc := &NetworkContainer{}
	nc.returnFields = []string{"extattrs", "comment"}

	err := objMgr.connector.GetObject(
		nc, ref, NewQueryParams(false, nil), nc)
	if err != nil {
		return nil, err
	}

	nc.Ea = setEas
	nc.Comment = comment

	// Network view is not allowed to be updated,
	// thus making its name empty (will not appear among data which we update).
	netViewSaved := nc.NetviewName
	nc.NetviewName = ""

	reference, err := objMgr.connector.UpdateObject(nc, ref)
	if err != nil {
		return nil, err
	}

	nc.Ref = reference
	nc.NetviewName = netViewSaved

	return nc, nil
}

func (objMgr *ObjectManager) AllocateNetworkContainer(
	netview string,
	cidr string,
	isIPv6 bool,
	prefixLen uint,
	comment string,
	eas EA) (*NetworkContainer, error) {

	containerInfo := NewNetworkContainerNextAvailableInfo(netview, cidr, prefixLen, isIPv6)
	container := NewNetworkContainerNextAvailable(containerInfo, isIPv6, comment, eas)

	ref, err := objMgr.connector.CreateObject(container)

	if err != nil {
		return nil, err
	}

	if isIPv6 {
		return BuildIPv6NetworkContainerFromRef(ref)
	} else {
		return BuildNetworkContainerFromRef(ref)
	}
}

func (objMgr *ObjectManager) AllocateNetworkContainerByEA(
	netview string, isIPv6 bool, comment string, eas EA, eaMap map[string]string, prefixLen uint) (*NetworkContainer, error) {

	var object string
	object = getNetworkObjectType(isIPv6, "networkcontainer", "ipv6networkcontainer")

	nextAvailableNetworkInfo := NetworkContainerNextAvailableInfo{
		Function:     "next_available_network",
		ResultField:  "networks",
		Object:       object,
		ObjectParams: eaMap,
		Params:       map[string]uint{"cidr": prefixLen},
	}

	net := NetworkContainerNextAvailable{
		Network:     &nextAvailableNetworkInfo,
		objectType:  object,
		Comment:     comment,
		Ea:          eas,
		NetviewName: netview,
	}
	ref, err := objMgr.connector.CreateObject(&net)

	if err != nil {
		return nil, err
	}
	if isIPv6 {
		return BuildIPv6NetworkContainerFromRef(ref)
	} else {
		return BuildNetworkContainerFromRef(ref)
	}
}

func (objMgr *ObjectManager) DeleteNetworkContainer(ref string) (string, error) {
	ncRegExp := regexp.MustCompile("^(ipv6)?networkcontainer\\/.+")
	if !ncRegExp.MatchString(ref) {
		return "", fmt.Errorf("'ref' does not reference a network container")
	}

	return objMgr.connector.DeleteObject(ref)
}
