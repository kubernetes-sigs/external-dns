package ibclient

import "fmt"

// Creates TXT Record. Use TTL of 0 to inherit TTL from the Zone
func (objMgr *ObjectManager) CreateTXTRecord(
	dnsView string,
	recordName string,
	text string,
	ttl uint32,
	useTtl bool,
	comment string,
	eas EA) (*RecordTXT, error) {

	recordTXT := NewRecordTXT(dnsView, "", recordName, text, ttl, useTtl, comment, eas)

	ref, err := objMgr.connector.CreateObject(recordTXT)
	if err != nil {
		return nil, err
	}
	recordTXT, err = objMgr.GetTXTRecordByRef(ref)
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecordByRef(ref string) (*RecordTXT, error) {
	recordTXT := NewEmptyRecordTXT()
	err := objMgr.connector.GetObject(
		recordTXT, ref, NewQueryParams(false, nil), &recordTXT)
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecord(dnsview string, name string) (*RecordTXT, error) {
	if dnsview == "" || name == "" {
		return nil, fmt.Errorf("DNS view and name are required to retrieve a unique txt record")
	}
	var res []RecordTXT

	recordTXT := NewEmptyRecordTXT()

	sf := map[string]string{
		"view": dnsview,
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordTXT, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"TXT record with name '%s' in DNS view '%s' is not found",
				name, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) UpdateTXTRecord(
	ref string,
	recordName string,
	text string,
	ttl uint32,
	useTtl bool,
	comment string,
	eas EA) (*RecordTXT, error) {

	recordTXT := NewRecordTXT("", "", recordName, text, ttl, useTtl, comment, eas)
	recordTXT.Ref = ref

	reference, err := objMgr.connector.UpdateObject(recordTXT, ref)
	if err != nil {
		return nil, err
	}

	recordTXT, err = objMgr.GetTXTRecordByRef(reference)
	return recordTXT, err
}

func (objMgr *ObjectManager) DeleteTXTRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}
