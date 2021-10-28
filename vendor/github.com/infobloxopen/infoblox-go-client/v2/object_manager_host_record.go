package ibclient

import (
	"fmt"
	"net"
)

func (objMgr *ObjectManager) CreateHostRecord(
	enabledns bool,
	enableDhcp bool,
	recordName string,
	netview string,
	dnsview string,
	ipv4cidr string,
	ipv6cidr string,
	ipv4Addr string,
	ipv6Addr string,
	macAddr string,
	duid string,
	useTtl bool,
	ttl uint32,
	comment string,
	eas EA,
	aliases []string) (*HostRecord, error) {

	if ipv4Addr == "" && ipv4cidr != "" {
		if netview == "" {
			netview = "default"
		}
		ipv4Addr = fmt.Sprintf("func:nextavailableip:%s,%s", ipv4cidr, netview)
	}
	if ipv6Addr == "" && ipv6cidr != "" {
		if netview == "" {
			netview = "default"
		}
		ipv6Addr = fmt.Sprintf("func:nextavailableip:%s,%s", ipv6cidr, netview)
	}
	recordHost := NewEmptyHostRecord()
	recordHostIpv6AddrSlice := []HostRecordIpv6Addr{}
	recordHostIpv4AddrSlice := []HostRecordIpv4Addr{}
	if ipv6Addr != "" {
		enableDhcpv6 := false
		if enableDhcp && duid != "" {
			enableDhcpv6 = true
		}
		recordHostIpv6Addr := NewHostRecordIpv6Addr(ipv6Addr, duid, enableDhcpv6, "")
		recordHostIpv6AddrSlice = []HostRecordIpv6Addr{*recordHostIpv6Addr}
	}
	if ipv4Addr != "" {
		enableDhcpv4 := false
		if enableDhcp && macAddr != "" && macAddr != MACADDR_ZERO {
			enableDhcpv4 = true
		}
		recordHostIpAddr := NewHostRecordIpv4Addr(ipv4Addr, macAddr, enableDhcpv4, "")
		recordHostIpv4AddrSlice = []HostRecordIpv4Addr{*recordHostIpAddr}
	}
	recordHost = NewHostRecord(
		netview, recordName, "", "", recordHostIpv4AddrSlice, recordHostIpv6AddrSlice,
		eas, enabledns, dnsview, "", "", useTtl, ttl, comment, aliases)
	ref, err := objMgr.connector.CreateObject(recordHost)
	if err != nil {
		return nil, err
	}
	recordHost.Ref = ref
	err = objMgr.connector.GetObject(
		recordHost, ref, NewQueryParams(false, nil), &recordHost)
	return recordHost, err
}

func (objMgr *ObjectManager) GetHostRecordByRef(ref string) (*HostRecord, error) {
	recordHost := NewEmptyHostRecord()
	err := objMgr.connector.GetObject(
		recordHost, ref, NewQueryParams(false, nil), &recordHost)

	return recordHost, err
}

func (objMgr *ObjectManager) SearchHostRecordByAltId(
	internalId string, ref string, eaNameForInternalId string) (*HostRecord, error) {

	if internalId == "" {
		return nil, fmt.Errorf("internal ID must not be empty")
	}

	recordHost := NewEmptyHostRecord()
	if ref != "" {
		if err := objMgr.connector.GetObject(recordHost, ref, NewQueryParams(false, nil), &recordHost); err != nil {
			return nil, err
		}
	}
	if recordHost.Ref != "" {
		return recordHost, nil
	}
	sf := map[string]string{
		fmt.Sprintf("*%s", eaNameForInternalId): internalId,
	}

	res := make([]HostRecord, 0)
	err := objMgr.connector.GetObject(recordHost, "", NewQueryParams(false, sf), &res)

	if err != nil {
		return nil, err
	}

	if res == nil || len(res) == 0 {
		return nil, NewNotFoundError("host record not found")
	}

	result := res[0]

	return &result, nil
}

func (objMgr *ObjectManager) GetHostRecord(netview string, dnsview string, recordName string, ipv4addr string, ipv6addr string) (*HostRecord, error) {
	var res []HostRecord

	recordHost := NewEmptyHostRecord()

	sf := map[string]string{
		"name": recordName,
	}
	if netview != "" {
		sf["network_view"] = netview
	}
	if dnsview != "" {
		sf["view"] = dnsview
	}
	if ipv4addr != "" {
		sf["ipv4addr"] = ipv4addr
	}
	if ipv6addr != "" {
		sf["ipv6addr"] = ipv6addr
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordHost, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}
	return &res[0], err

}

func (objMgr *ObjectManager) GetIpAddressFromHostRecord(host HostRecord) (string, error) {
	err := objMgr.connector.GetObject(
		&host, host.Ref, NewQueryParams(false, nil), &host)
	return host.Ipv4Addrs[0].Ipv4Addr, err
}

func (objMgr *ObjectManager) UpdateHostRecord(
	hostRref string,
	enabledns bool,
	enableDhcp bool,
	name string,
	netView string,
	ipv4cidr string,
	ipv6cidr string,
	ipv4Addr string,
	ipv6Addr string,
	macAddr string,
	duid string,
	useTtl bool,
	ttl uint32,
	comment string,
	eas EA,
	aliases []string) (*HostRecord, error) {

	recordHostIpv4AddrSlice := []HostRecordIpv4Addr{}
	recordHostIpv6AddrSlice := []HostRecordIpv6Addr{}

	enableDhcpv4 := false
	enableDhcpv6 := false
	if ipv6Addr != "" {
		if enableDhcp && duid != "" {
			enableDhcpv6 = true
		}
	}
	if ipv4Addr != "" {
		if enableDhcp && macAddr != "" && macAddr != MACADDR_ZERO {
			enableDhcpv4 = true
		}
	}

	if ipv4Addr == "" {
		if ipv4cidr != "" {
			ip, _, err := net.ParseCIDR(ipv4cidr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
			}
			if ip.To4() == nil {
				return nil, fmt.Errorf("CIDR value must be an IPv4 CIDR, not an IPv6 one")
			}
			if netView == "" {
				netView = "default"
			}
			newIpAddr := fmt.Sprintf("func:nextavailableip:%s,%s", ipv4cidr, netView)
			recordHostIpAddr := NewHostRecordIpv4Addr(newIpAddr, macAddr, enableDhcpv4, "")
			recordHostIpv4AddrSlice = []HostRecordIpv4Addr{*recordHostIpAddr}
		}
	} else {
		ip := net.ParseIP(ipv4Addr)
		if ip == nil {
			return nil, fmt.Errorf("'IP address for the record is not valid")
		}
		if ip.To4() == nil {
			return nil, fmt.Errorf("IP address must be an IPv4 address, not an IPv6 one")
		}
		recordHostIpAddr := NewHostRecordIpv4Addr(ipv4Addr, macAddr, enableDhcpv4, "")
		recordHostIpv4AddrSlice = []HostRecordIpv4Addr{*recordHostIpAddr}
	}
	if ipv6Addr == "" {
		if ipv6cidr != "" {
			ip, _, err := net.ParseCIDR(ipv6cidr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse CIDR value: %s", err.Error())
			}
			if ip.To4() != nil || ip.To16() == nil {
				return nil, fmt.Errorf("CIDR value must be an IPv6 CIDR, not an IPv4 one")
			}
			if netView == "" {
				netView = "default"
			}
			newIpAddr := fmt.Sprintf("func:nextavailableip:%s,%s", ipv6cidr, netView)
			recordHostIpAddr := NewHostRecordIpv6Addr(newIpAddr, duid, enableDhcpv6, "")
			recordHostIpv6AddrSlice = []HostRecordIpv6Addr{*recordHostIpAddr}
		}
	} else {
		ip := net.ParseIP(ipv6Addr)
		if ip == nil {
			return nil, fmt.Errorf("IP address for the record is not valid")
		}
		if ip.To4() != nil || ip.To16() == nil {
			return nil, fmt.Errorf("IP address must be an IPv6 address, not an IPv4 one")
		}
		recordHostIpAddr := NewHostRecordIpv6Addr(ipv6Addr, duid, enableDhcpv6, "")
		recordHostIpv6AddrSlice = []HostRecordIpv6Addr{*recordHostIpAddr}
	}
	updateHostRecord := NewHostRecord(
		"", name, "", "", recordHostIpv4AddrSlice, recordHostIpv6AddrSlice,
		eas, enabledns, "", "", hostRref, useTtl, ttl, comment, aliases)
	ref, err := objMgr.connector.UpdateObject(updateHostRecord, hostRref)
	if err != nil {
		return nil, err
	}

	updateHostRecord, err = objMgr.GetHostRecordByRef(ref)
	if err != nil {
		return nil, err
	}
	return updateHostRecord, nil
}

func (objMgr *ObjectManager) DeleteHostRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
