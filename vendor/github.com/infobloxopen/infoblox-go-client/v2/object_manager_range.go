package ibclient

import (
	"encoding/json"
	"fmt"
)

func (d Range) MarshalJSON() ([]byte, error) {
	type Alias Range
	aux := &struct {
		Member *Dhcpmember `json:"member"`
		*Alias
	}{
		Member: d.Member,
		Alias:  (*Alias)(&d),
	}
	return json.Marshal(aux)
}

func (d *Range) UnmarshalJSON(data []byte) error {
	type Alias Range
	aux := &struct {
		Member *Dhcpmember `json:"member"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Member = aux.Member
	return nil
}
func NewEmptyRange() *Range {
	newRange := &Range{}
	newRange.SetReturnFields(append(newRange.ReturnFields(), "extattrs", "name", "disable", "options", "use_options", "cloud_info", "failover_association", "member", "server_association_type", "ms_server"))
	return newRange
}
func NewRange(comment string,
	name string,
	network *string,
	startAddr string,
	eas EA,
	disable bool,
	options []*Dhcpoption,
	useOptions bool,
	endAddr string,
	failOverAssociation string,
	member *Dhcpmember,
	ServerAssociationType string,
	template string,
	msServer string,
) *Range {
	newRange := NewEmptyRange()
	newRange.Comment = &comment
	newRange.Name = &name
	newRange.Network = network
	newRange.StartAddr = &startAddr
	newRange.Ea = eas
	newRange.Disable = &disable
	if options != nil {
		newRange.Options = options
	}
	newRange.UseOptions = &useOptions
	newRange.EndAddr = &endAddr
	newRange.FailoverAssociation = &failOverAssociation
	newRange.Member = member
	newRange.ServerAssociationType = ServerAssociationType
	newRange.Template = template
	newRange.MsServer = &Msdhcpserver{Ipv4Addr: msServer}
	return newRange
}
func (objMgr *ObjectManager) CreateNetworkRange(comment string, name string, network string, networkView string, startAddr string, endAddr string, disable bool, eas EA, member *Dhcpmember, failOverAssociation string, options []*Dhcpoption, useOptions bool, serverAssociation string, template string, msServer string) (*Range, error) {

	if startAddr == "" || endAddr == "" {
		return nil, fmt.Errorf("start address and end address fields are required to create a range within a Network")
	}
	if networkView == "" {
		networkView = "default"
	}
	var networkPointer *string
	if network != "" {
		networkPointer = &network
	}
	newRangeCreate := NewRange(comment, name, networkPointer, startAddr, eas, disable, options, useOptions, endAddr, failOverAssociation, member, serverAssociation, template, msServer)
	newRangeCreate.NetworkView = &networkView
	ref, err := objMgr.connector.CreateObject(newRangeCreate)
	if err != nil {
		return nil, err
	}
	newRangeCreate.Ref = ref
	return newRangeCreate, nil
}
func (objMgr *ObjectManager) GetNetworkRangeByRef(ref string) (*Range, error) {
	networkRange := NewEmptyRange()
	err := objMgr.connector.GetObject(
		networkRange, ref, NewQueryParams(false, nil), &networkRange)

	return networkRange, err
}
func (objMgr *ObjectManager) GetNetworkRange(queryParams *QueryParams) ([]Range, error) {
	var res []Range
	networkRange := NewEmptyRange()
	err := objMgr.connector.GetObject(
		networkRange, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting DHCP IPv4 Range: %s", err)
	}
	return res, nil
}
func (objMgr *ObjectManager) UpdateNetworkRange(ref string, comment string, name string, network string, startAddr string, endAddr string, disable bool, eas EA, member *Dhcpmember, failOverAssociation string, options []*Dhcpoption, useOptions bool, serverAssociationType string, NetworkView string, msServer string) (*Range, error) {
	if startAddr == "" || endAddr == "" {
		return nil, fmt.Errorf("start address and end address fields cannot be empty for a range within a Network")
	}
	var networkPointer *string
	if network != "" {
		networkPointer = &network
	}
	networkRange := NewRange(comment, name, networkPointer, startAddr, eas, disable, options, useOptions, endAddr, failOverAssociation, member, serverAssociationType, "", msServer)
	networkRange.NetworkView = &NetworkView
	networkRange.Ref = ref
	reference, err := objMgr.connector.UpdateObject(networkRange, ref)
	if err != nil {
		return nil, err
	}
	networkRange.Ref = reference

	networkRange, err = objMgr.GetNetworkRangeByRef(reference)
	if err != nil {
		return nil, err
	}

	return networkRange, nil
}
func (objMgr *ObjectManager) DeleteNetworkRange(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
