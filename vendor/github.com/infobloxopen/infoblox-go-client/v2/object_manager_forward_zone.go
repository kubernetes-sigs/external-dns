package ibclient

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Forwarding Server to be [] list
type NullableForwardingServers struct {
	Servers []*Forwardingmemberserver
	IsNull  bool // Indicates if the entire slice should be null
}

func (nfs NullableForwardingServers) MarshalJSON() ([]byte, error) {
	if reflect.DeepEqual(nfs.Servers, []*Forwardingmemberserver{}) {
		return []byte("[]"), nil
	}

	return json.Marshal(nfs.Servers)
}

func (nfs *NullableForwardingServers) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nfs.IsNull = true
		nfs.Servers = nil
		return nil
	}
	nfs.IsNull = false
	return json.Unmarshal(data, &nfs.Servers)
}

func (objMgr *ObjectManager) CreateZoneForward(
	comment string,
	disable bool,
	eas EA,
	forwardTo NullableNameServers,
	forwardersOnly bool,
	forwardingServers *NullableForwardingServers,
	fqdn string,
	nsGroup string,
	view string,
	zoneFormat string,
	externalNsGroup string) (*ZoneForward, error) {
	// check if required fields are present
	// Check for FQDN
	if fqdn == "" {
		return nil, fmt.Errorf("FQDN is required to create a forward zone")
	}
	if forwardingServers != nil && forwardingServers.Servers != nil {
		forwardingServers = &NullableForwardingServers{Servers: forwardingServers.Servers}
	} else {
		forwardingServers = nil
	}
	zoneForward := NewZoneForward(comment, disable, eas, forwardTo, forwardersOnly, forwardingServers, fqdn, nsGroup, view, zoneFormat, "", externalNsGroup)
	ref, err := objMgr.connector.CreateObject(zoneForward)
	if err != nil {
		return nil, err
	}
	zoneForward.Ref = ref
	return zoneForward, nil
}

func (objMgr *ObjectManager) DeleteZoneForward(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)

}

func (objMgr *ObjectManager) GetZoneForwardByRef(ref string) (*ZoneForward, error) {
	zoneForward := NewEmptyZoneForward()
	err := objMgr.connector.GetObject(zoneForward, ref, NewQueryParams(false, nil), &zoneForward)
	if err != nil {
		return nil, err
	}
	return zoneForward, nil
}

func (objMgr *ObjectManager) GetZoneForwardFilters(queryParams *QueryParams) ([]ZoneForward, error) {

	var res []ZoneForward
	zoneForward := NewEmptyZoneForward()

	err := objMgr.connector.GetObject(
		zoneForward, "", queryParams, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (objMgr *ObjectManager) UpdateZoneForward(
	ref string,
	comment string,
	disable bool,
	eas EA,
	forwardTo NullableNameServers,
	forwardersOnly bool,
	forwardingServers *NullableForwardingServers,
	nsGroup string,
	externlNsGroup string) (*ZoneForward, error) {

	zoneForward := NewEmptyZoneForward()

	zoneForward.Comment = &comment
	zoneForward.Disable = &disable
	zoneForward.Ea = eas
	zoneForward.ForwardTo = forwardTo
	zoneForward.ForwardersOnly = &forwardersOnly
	if forwardingServers != nil && forwardingServers.Servers != nil {
		zoneForward.ForwardingServers = &NullableForwardingServers{Servers: forwardingServers.Servers}
	} else {
		zoneForward.ForwardingServers = nil
	}
	if nsGroup != "" {
		zoneForward.NsGroup = &nsGroup
	} else {
		zoneForward.NsGroup = nil
	}
	if externlNsGroup != "" {
		zoneForward.ExternalNsGroup = &externlNsGroup
	} else {
		zoneForward.ExternalNsGroup = nil
	}

	new_ref, err := objMgr.connector.UpdateObject(zoneForward, ref)
	if err != nil {
		return nil, err
	}
	zoneForward.Ref = new_ref
	return zoneForward, nil

}

func NewEmptyZoneForward() *ZoneForward {
	zoneForward := &ZoneForward{}
	zoneForward.SetReturnFields(append(zoneForward.ReturnFields(), "zone_format", "ns_group", "external_ns_group", "comment", "disable", "extattrs", "forwarders_only", "forwarding_servers"))
	return zoneForward
}

func NewZoneForward(comment string,
	disable bool,
	eas EA,
	forwardTo NullableNameServers,
	forwardersOnly bool,
	forwardingServers *NullableForwardingServers,
	fqdn string,
	nsGroup string,
	view string,
	zoneFormat string,
	ref string,
	externalNsGroup string) *ZoneForward {

	zoneForward := NewEmptyZoneForward()

	zoneForward.Comment = &comment
	zoneForward.Disable = &disable
	zoneForward.Ea = eas
	zoneForward.ForwardTo = forwardTo
	zoneForward.ForwardersOnly = &forwardersOnly
	zoneForward.ForwardingServers = forwardingServers

	zoneForward.Fqdn = fqdn
	if nsGroup == "" {
		zoneForward.NsGroup = nil
	} else {
		zoneForward.NsGroup = &nsGroup
	}
	if externalNsGroup == "" {
		zoneForward.ExternalNsGroup = nil
	} else {
		zoneForward.ExternalNsGroup = &externalNsGroup
	}
	if view == "" {
		view = "default"
	}
	zoneForward.View = &view
	if zoneFormat == "" {
		zoneFormat = "FORWARD"
	}
	zoneForward.ZoneFormat = zoneFormat
	zoneForward.Ref = ref

	return zoneForward
}
