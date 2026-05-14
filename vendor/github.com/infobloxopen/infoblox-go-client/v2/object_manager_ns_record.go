package ibclient

import (
	"encoding/json"
	"fmt"
)

func (z ZoneNameServer) MarshalJSON() ([]byte, error) {
	type Alias ZoneNameServer
	return json.Marshal(&struct {
		*Alias
		AutoCreatePtr bool `json:"auto_create_ptr"`
	}{
		Alias:         (*Alias)(&z),
		AutoCreatePtr: z.AutoCreatePtr,
	})
}
func NewRecordNS(name string, nameServer string, dnsView string, addresses []*ZoneNameServer, msDelegationName string) *RecordNS {
	res := NewEmptyRecordNS()
	res.Name = name
	res.Nameserver = &nameServer
	res.View = dnsView
	res.Addresses = addresses
	res.MsDelegationName = &msDelegationName
	return res
}

func NewEmptyRecordNS() *RecordNS {
	res := &RecordNS{}
	res.SetReturnFields(append(res.returnFields, "addresses", "cloud_info", "creator", "dns_name", "last_queried", "ms_delegation_name", "name", "nameserver", "policy", "view", "zone"))
	return res
}

func (objMgr *ObjectManager) CreateNSRecord(name string, nameServer string, dnsView string, addresses []*ZoneNameServer, msDelegationName string) (*RecordNS, error) {
	if name == "" || nameServer == "" || len(addresses) == 0 {
		return nil, fmt.Errorf("name, nameserver and addresses are required for NS record creation")
	}
	if dnsView == "" {
		dnsView = "default"
	}
	nsRecord := NewRecordNS(name, nameServer, dnsView, addresses, msDelegationName)
	ref, err := objMgr.connector.CreateObject(nsRecord)
	if err != nil {
		return nil, err
	}
	nsRecord.Ref = ref
	return nsRecord, nil
}

func (objMgr *ObjectManager) GetNSRecordByRef(ref string) (*RecordNS, error) {
	recordNS := NewEmptyRecordNS()
	err := objMgr.connector.GetObject(
		recordNS, ref, NewQueryParams(false, nil), &recordNS)
	return recordNS, err
}

func (objMgr *ObjectManager) DeleteNSRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) UpdateNSRecord(ref string, name string, nameServer string, dnsView string, addresses []*ZoneNameServer, msDelegationName string) (*RecordNS, error) {
	nsRecord := NewRecordNS(name, nameServer, dnsView, addresses, msDelegationName)
	nsRecord.Ref = ref

	ref, err := objMgr.connector.UpdateObject(nsRecord, ref)
	if err != nil {
		return nil, err
	}
	newRec, err := objMgr.GetNSRecordByRef(ref)
	if err != nil {
		return nil, err
	}
	return newRec, nil
}

func (objMgr *ObjectManager) GetAllRecordNS(queryParams *QueryParams) ([]RecordNS, error) {
	var res []RecordNS
	recordNS := NewEmptyRecordNS()
	err := objMgr.connector.GetObject(recordNS, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting NS Record: %s", err)
	}
	return res, nil
}
