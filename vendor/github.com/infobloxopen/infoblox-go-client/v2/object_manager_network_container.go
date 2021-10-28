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

	reference, err := objMgr.connector.UpdateObject(nc, ref)
	if err != nil {
		return nil, err
	}

	nc.Ref = reference
	return nc, nil
}

func (objMgr *ObjectManager) DeleteNetworkContainer(ref string) (string, error) {
	ncRegExp := regexp.MustCompile("^(ipv6)?networkcontainer\\/.+")
	if !ncRegExp.MatchString(ref) {
		return "", fmt.Errorf("'ref' does not reference a network container")
	}

	return objMgr.connector.DeleteObject(ref)
}
