package ibclient

import "fmt"

func (objMgr *ObjectManager) GetDnsMember(ref string) ([]Dns, error) {
	var res []Dns
	queryParams := NewQueryParams(false, nil)
	dns := NewDns(Dns{})
	err := objMgr.connector.GetObject(dns, ref, queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"could not find any node"))
	}
	return res, nil
}

func (objMgr *ObjectManager) GetDhcpMember(ref string) ([]Dhcp, error) {
	var res []Dhcp
	queryParams := NewQueryParams(false, nil)
	dhcp := NewDhcp(Dhcp{})
	err := objMgr.connector.GetObject(dhcp, ref, queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"could not find ny node"))
	}
	return res, nil
}

func (objMgr *ObjectManager) UpdateDnsStatus(ref string, status bool) (Dns, error) {
	dns := NewDns(Dns{})
	dns.EnableDns = status
	resp, err := objMgr.connector.UpdateObject(dns, ref)
	if err != nil {
		return *dns, err
	}
	dns.Ref = resp
	return *dns, nil
}

func (objMgr *ObjectManager) UpdateDhcpStatus(ref string, status bool) (Dhcp, error) {
	dhcp := NewDhcp(Dhcp{})
	dhcp.EnableDhcp = status
	resp, err := objMgr.connector.UpdateObject(dhcp, ref)
	if err != nil {
		return *dhcp, err
	}
	dhcp.Ref = resp
	return *dhcp, nil
}
